package offlineconversation.dto;

import lombok.Builder;

import java.util.List;
import java.util.UUID;

@Builder
public record ChatMessage(
    // uuids, jackson writes them as their strings, as the uuid.UUID of my-server's ChatMessage reads them
    UUID id,
    UUID fromId,
    String toIdType,
    UUID toId,
    String contentType,
    List<String> contents
) {
}