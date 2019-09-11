package models

import "testing"

func TestHand_Evaluate(t *testing.T) {
	type fields struct {
		Cards        Cards
		BestHand     Cards
		BestHandType HandType
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		"Test_1", fields{Cards{}}, false,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hand{
				Cards:        tt.fields.Cards,
				BestHand:     tt.fields.BestHand,
				BestHandType: tt.fields.BestHandType,
			}
			if err := h.Parse(); (err != nil) != tt.wantErr {
				t.Errorf("Hand.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
