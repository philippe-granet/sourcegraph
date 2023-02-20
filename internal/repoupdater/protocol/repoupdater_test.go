package protocol

import (
	"testing"
	"testing/quick"
)

func TestProtoRoundtrip(t *testing.T) {
	t.Run("RepoUpdateSchedulerInfoResult", func(t *testing.T) {
		err := quick.Check(func(input1 RepoUpdateSchedulerInfoResult) bool {
			output1 := input1.ToProto()
			input2 := RepoUpdateSchedulerInfoResultFromProto(output1)
			return input1 == *input2
		}, nil)
		if err != nil {
			t.Fatal(err)
		}
	})
}
