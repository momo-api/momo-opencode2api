package protocol

import "testing"

func TestNormalizeResponsesProviderRequest(t *testing.T) {
	input := map[string]any{
		"model":             "muse-spark-1.3-contributor-free",
		"max_output_tokens": float64(1),
		"input": []any{
			map[string]any{
				"role":    "assistant",
				"content": []any{map[string]any{"type": "output_text", "text": "prior"}},
			},
		},
	}
	out, err := PrepareRequest(Responses, Responses, input, "https://opencode.ai/zen")
	if err != nil {
		t.Fatal(err)
	}
	if got := out["max_output_tokens"]; got != 16 {
		t.Fatalf("max_output_tokens = %#v, want 16", got)
	}
	items := out["input"].([]any)
	item := items[0].(map[string]any)
	if got := item["type"]; got != "message" {
		t.Fatalf("input[0].type = %#v, want message", got)
	}
	part := item["content"].([]any)[0].(map[string]any)
	if got := part["type"]; got != "input_text" {
		t.Fatalf("assistant history content type = %#v, want input_text", got)
	}
}

func TestNormalizeResponsesProviderRequestKeepsValidLimit(t *testing.T) {
	input := map[string]any{"model": "muse-spark-1.3-contributor-free", "max_output_tokens": float64(32), "input": "hello"}
	out, err := PrepareRequest(Responses, Responses, input, "https://opencode.ai/zen")
	if err != nil {
		t.Fatal(err)
	}
	if got := out["max_output_tokens"]; got != float64(32) {
		t.Fatalf("valid max_output_tokens changed to %#v", got)
	}
}
