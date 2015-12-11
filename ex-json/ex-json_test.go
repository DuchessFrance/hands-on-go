package ex_json
import (
	"testing"
)


func TestExJson_01 (t *testing.T ){

	duchessFrance := LoadJsonData()

	if(duchessFrance.Organisation != "Duchess France"){
		t.Fatalf("L'organisation n'est pas bien parsé")
	}

	if(duchessFrance.Team != "Go"){
		t.Fatalf("La valeur de la team n'est pas correcte")
	}

	if(len(duchessFrance.Members) != 2){
		t.Fatalf("L'équipe n'est pas correctement parsé")
	}
}