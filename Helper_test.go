package main

import (
	"strings"
	"testing"
)

func Test_parseAndRun(t *testing.T) {
	type args struct {
		prog string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example1", args{"print(1234)"}, "1234"},
		{"simple while", args{"i:=0;while(i<10){i=i+1;print(i)}"}, "12345678910"},
		{"simple while 2", args{"i:=0;while(i<10){i=i+1;print(i)}"}, "12345678910"},
		{"simple if else", args{"if(1<2){print(true)}else{print(false)}"}, "true"},
		{"else block", args{"if(2<1){print(true)}else{print(false)}"}, "false"},
		{"print number boolean", args{"print(-23456);print(true);print(false)"}, "-23456truefalse"},

		{"simple Declaration", args{"i:=0;print(i)"}, "0"},
		{"simple Assign", args{"i:=1;j:=5;j=i+j;print(j);print(i)"}, "61"},
		{"Declaration long variable name", args{"istDasHier_nichteinTollerVariblen--Name123:=0;print(istDasHier_nichteinTollerVariblen--Name123)"}, "0"},
		{"simple long variable name", args{"istDasHier_nichteinTollerVariblen--Name123:=1;istDasHier_nichteinTollerVariblen--Name123=3;print(istDasHier_nichteinTollerVariblen--Name123)"}, "3"},

		{"Assign Fail worng Type", args{"i:=0;i=true;print(i)"}, "Variable assignement Failed. Var: i0"},
		{"Assign Variable not Declared", args{"i=true;print(i)"}, "Variable unknown. Var: iError Evaluating Print Statement Illtyped Value"},
		{"Declaration Override Variable", args{"i:=0;i:=true;print(i)"}, "true"},

		{"Punkt vor strich", args{"print(2+3*4)"}, "14"},
		{"not first", args{"print(true||!true);print(false&&!false)"}, "truefalse"},
		{"und vor oder", args{"print(true||false&&false)"}, "true"},
		{"equal vor und", args{"print(false==false&&false)"}, "false"},
		{"Klammern zuerst", args{"print((2+3)*4)"}, "20"},

		{"simple ==", args{"print(2==2);print(true==true);print(true==false);print(2==-2)"}, "truetruefalsefalse"},
		{"simple +", args{"print(2+5)"}, "7"},
		{"simple *", args{"print(1*12345)"}, "12345"},
		{"simple ||", args{"print(true||true);print(true||false);print(false||true);print(false||false)"}, "truetruetruefalse"},
		{"simple &&", args{"print(true&&true);print(true&&false);print(false&&true);print(false&&false)"}, "truefalsefalsefalse"},
		{"simple <", args{"print(2<5);print(5<2);print(2<2)"}, "truefalsefalse"},
		{"simple !", args{"print(!true);print(!false);i:=true;print(!i)"}, "falsetruefalse"},

		{"|| Illtyped eval success ", args{"print(true||1)"}, "true"},
		{"|| Illtyped eval fail", args{"print(false||1)"}, "Error Evaluating Print Statement Illtyped Value"},
		{"&& Illtyped eval sucess", args{"print(false&&1)"}, "false"},
		{"&& Illtyped eval fail", args{"print(true&&1)"}, "Error Evaluating Print Statement Illtyped Value"},
		{"< Illtyped eval fail", args{"print(false<1)"}, "Error Evaluating Print Statement Illtyped Value"},
		{"+ Illtyped eval fail", args{"print(false+1)"}, "Error Evaluating Print Statement Illtyped Value"},
		{"+ Illtyped eval fail_2", args{"print(false+false)"}, "Error Evaluating Print Statement Illtyped Value"},
		{"* Illtyped eval fail", args{"print(false*1)"}, "Error Evaluating Print Statement Illtyped Value"},
		{"* Illtyped eval fail_2", args{"print(false*false)"}, "Error Evaluating Print Statement Illtyped Value"},

		{"while override Variable", args{"i:=0;while(i<5){i=i+1;print(i);i:=i+5;print(i)}"}, "16273849510"},
		{"If override Variable", args{"i:=true;if(i){i:=5;print(i)}else{print(123)};print(i)"}, "5true"},
		{"Else override Variable", args{"i:=true;if(!i){print(123)}else{i:=5;print(i)};print(i)"}, "5true"},

		{"Komplex Expression 1", args{"print(2+5 * (15+-12))"}, "17"},
		{"Komplex Expression 2", args{"print((2+5)==0*9+3*((4+5)*1*0+2)+1)"}, "true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.TrimSpace(parseAndRun(tt.args.prog)); got != tt.want {
				t.Errorf("parseAndRun() = %v, want %v", got, tt.want)
			}
		})
	}
}
