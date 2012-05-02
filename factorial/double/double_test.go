package double_test

import (
	"math/big"
	"testing"

	"github.com/soniakeys/integer/swing"
	"github.com/soniakeys/integer/factorial/double"
)

func TestDoubleFactorial(t *testing.T) {
	d := new(big.Int)
	answer := new(big.Int)
	for _, tc := range tcs {
		if _, ok := answer.SetString(tc.s, 10); !ok {
			t.Error("invalid test case", tc)
			continue
		}
		switch _ = double.DoubleFactorial(d, tc.n); {
		case d == nil:
			t.Error("nil result. test case", tc)
		case d.Cmp(answer) != 0:
			t.Log("wrong answer, test case", tc)
			t.Error("got", answer)
		}
	}
}

func TestDoubleFactorialS(t *testing.T) {
	p := swing.New(tcs[len(tcs)-1].n + 1)
	d := new(big.Int)
	answer := new(big.Int)
	for _, tc := range tcs {
		if _, ok := answer.SetString(tc.s, 10); !ok {
			t.Error("invalid test case", tc)
			continue
		}
		switch _ = double.DoubleFactorialS(d, p, tc.n); {
		case d == nil:
			t.Error("nil result. test case", tc)
		case d.Cmp(answer) != 0:
			t.Log("wrong answer, test case", tc)
			t.Error("got", answer)
		}
	}
}

var tcs = []struct {
	n uint
	s string
}{
	{0, "1"},
	{1, "1"},
	{2, "2"},
	{3, "3"},
	{4, "8"},
	{5, "15"},
	{6, "48"},
	{7, "105"},
	{8, "384"},
	{9, "945"},
	{10, "3840"},
	{34, "46620662575398912000"},
	{37, "8200794532637891559375"},
	{63, "112275575285571389562324404930670903477890625"},
	{100, "34243224702511976248246432895208185975118675053719198827915654463488000000000000"},
	{103, "28352254429826839019508359891905756542124154226667992913078222750278633810791015625"},
	{104, "363252127644247044041398160152368436824058904969853261166529262548680704000000000000"},
	{500, "5849049697728183931901573966636399185893290101863305204136019757220414567257738129869679070426230366367652451980197858002263561449805551771020901113739313626336705563563705788360503630094403488675854668161534760788195420015279377621729517620792668944963981391489926671539372938481001173031117052763221491420281727661731208544954134335107331812412321791962113178938189516786683915122565052376248782141535507632768973188905459515532298174562947984906490257552942386774824261588679054048717674760963003462451200000000000000000000000000000000000000000000000000000000000000"},
	{501, "104510746430992941887672903221701774687693958020937523410364654058131809527854728379249021300764967518560976934211305561401337053955984964137051421361312300363167356445715307331726922796959128808412410939594196995562016845467597371284040921057059540932782811388408034712339663789273366959360793923894782445455391719381938160231368520328663730478625125685672007820874831426096933254104094661493009619303330424832414231055479364575767375357177905153440132816398992407261547919121408710339242666003272288642044702058302205588347587837672048038939465186558663845062255859375"},
	{1000, "3993984426547508861305390944632658769491465612732950884149352264463669653783918786673855847931586297648724730100712687233287867670033234347147721372160052970468852277435367085064790343283345982785772859940510918528070095138370429504769382461652184670792668564749747182728979673007824853940403174783379461981377594688094308995249697607727579137178606555069217194222363439778224206578431466544196150784271265561123088538027593033553551529401457853649535775425729310701035075686547366124355050660879155316288666973553745148520593257875601748532607163260780959394611457161826005390022163500384517756609464749567995251111398216507373580124815177400601722590893150586590359116008988749851761693190689228238008607755225687337923116872708482175000550889761043932312692198817965575874482087355854957667093854984963161313376186755900436597903368883575958968256738241094038338021793478788394592854525976460113138953893460202621461412466767406873753554433304755169025189150527313391959692691048456575394801454926197688818611863300874422358915700504682504830680617901494815422107753552276905563018887609770664552859008696057852779541484186292035509861625299662585190825328640000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
	{1001, "100849078093514604858042956309695841788573742002133966120718791290848182778795769931238362072611605430738671303158405922628623779705709875508286507881085570912184252140978492399288243766055452279190633553920991605565123686761347923894997207602460497394948626512140144240529165580243443608574307800024801466068095382874779101467296843546378688154339608903661050806863388534752621140691894596069672092368738706686458488964417636767119622835575185944553850974213261986345305203847488799629608101465848146363138762167180653056251006205563112031443860585897034397301298933754582457436080124441856183797115696583020394154991605786017411015430374818535553456362737377906985432778991103164065509107420268289234945537195124298972426095870588387306324700328785020361207669201115077041899148597103269701947621700634580709710853341206731241973992179503113789399883016283981715245005827871765155147853051640970219484088275333649293391870081482096390856650419636052213653410958282285786208628865361658196064720095614768521629999735498307012084078296034580060845031941815475649038649725375489312770050456989102358538379742333862513601611507823761965139985270584602153617130262906650646765124072268186054313817572598765698916028893868646350390930900454631086366787118269172651707776822149753570556640625"},
}

func BenchmarkDouble1e2(b *testing.B) {
	var z big.Int
	for i := 0; i < b.N; i++ {
		double.DoubleFactorial(&z, 1e2)
	}
}

func BenchmarkDouble1e3(b *testing.B) {
	var z big.Int
	for i := 0; i < b.N; i++ {
		double.DoubleFactorial(&z, 1e3)
	}
}

func BenchmarkDouble1e4(b *testing.B) {
	var z big.Int
	for i := 0; i < b.N; i++ {
		double.DoubleFactorial(&z, 1e4)
	}
}

func BenchmarkDouble1e5(b *testing.B) {
	var z big.Int
	for i := 0; i < b.N; i++ {
		double.DoubleFactorial(&z, 1e5)
	}
}

func BenchmarkDouble2e5(b *testing.B) {
	var z big.Int
	for i := 0; i < b.N; i++ {
		double.DoubleFactorial(&z, 2e5)
	}
}

var s *swing.Swing

func init() {
	s = swing.New(2e5)
}

func BenchmarkDoubleS1e2(b *testing.B) {
	var z big.Int
	for i := 0; i < b.N; i++ {
		double.DoubleFactorialS(&z, s, 1e2)
	}
}

func BenchmarkDoubleS1e3(b *testing.B) {
	var z big.Int
	for i := 0; i < b.N; i++ {
		double.DoubleFactorialS(&z, s, 1e3)
	}
}

func BenchmarkDoubleS1e4(b *testing.B) {
	var z big.Int
	for i := 0; i < b.N; i++ {
		double.DoubleFactorialS(&z, s, 1e4)
	}
}

func BenchmarkDoubleS1e5(b *testing.B) {
	var z big.Int
	for i := 0; i < b.N; i++ {
		double.DoubleFactorialS(&z, s, 1e5)
	}
}

func BenchmarkDoubleS2e5(b *testing.B) {
	var z big.Int
	for i := 0; i < b.N; i++ {
		double.DoubleFactorialS(&z, s, 2e5)
	}
}
