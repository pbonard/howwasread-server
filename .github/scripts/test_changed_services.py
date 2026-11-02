import json
import unittest
from pathlib import Path

from changed_services import select_services

SERVICES = json.loads((Path(__file__).parent.parent / "services.json").read_text())


def names(files):
    return [s["name"] for s in select_services(SERVICES, files)]


class SelectServicesTest(unittest.TestCase):
    def test_a_service_folder_selects_only_that_service(self):
        self.assertEqual(names(["backend/offlineconversation/src/Main.java"]), ["offlineconversation"])

    def test_common_selects_every_service_using_it(self):
        self.assertEqual(
            names(["backend/common/payload.go"]),
            ["onlineconversation"],
        )

    def test_go_module_files_select_the_go_services(self):
        self.assertEqual(names(["backend/go.mod"]), ["onlineconversation"])

    def test_files_outside_the_services_select_nothing(self):
        self.assertEqual(names(["README.md", "k8s/helm/values.yaml", "backend/.mockery.yml"]), [])

    def test_a_folder_sharing_a_prefix_is_not_the_service(self):
        self.assertEqual(names(["backend/onlineconversationx/main.go"]), [])

    def test_the_old_root_paths_select_nothing(self):
        self.assertEqual(names(["onlineconversation/main.go", "common/payload.go"]), [])

    def test_the_pipeline_itself_selects_every_service(self):
        for f in [".github/services.json", ".github/workflows/services.yml", ".github/scripts/deploy_tags.py"]:
            self.assertEqual(len(names([f])), len(SERVICES), f)

    def test_an_unknown_previous_commit_selects_every_service(self):
        self.assertEqual(len(names(None)), len(SERVICES))


if __name__ == "__main__":
    unittest.main()
