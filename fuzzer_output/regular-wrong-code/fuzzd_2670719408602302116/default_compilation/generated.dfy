datatype D0 = DC1(cf1: int, cf2: int) | DC2 | DC0(cf0: int) | DC3(cf3: D0)
datatype D1 = DC4(cf4: bool)
datatype D2 = DC6(cf6: int, cf7: bool, cf8: bool) | DC7(cf9: char, cf10: bool, cf11: int) | DC5(cf5: array<bool>)
datatype D3 = DC9 | DC8(cf12: map<int, int>)
datatype D4 = DC11(cf14: bool, cf15: bool, cf16: int, cf17: string) | DC10(cf13: seq<bool>) | DC12(cf18: D4)
datatype D5 = DC14(cf20: int, cf21: bool) | DC15 | DC13(cf19: map<bool, bool>)
datatype D6 = DC16(cf22: multiset<int>)
datatype D7 = DC18 | DC17(cf23: map<int, bool>)
datatype D8 = DC19(cf24: seq<int>)
datatype D9 = DC21(cf26: bool, cf27: seq<bool>, cf28: bool) | DC22(cf29: seq<bool>) | DC20(cf25: set<bool>)
datatype D10 = DC24(cf31: int, cf32: bool, cf33: seq<char>) | DC23(cf30: multiset<string>)
datatype D11 = DC25(cf34: multiset<bool>)
datatype D12 = DC27(cf36: int, cf37: int, cf38: int, cf39: int) | DC26(cf35: map<bool, int>) | DC28(cf40: D12)
datatype D13 = DC30(cf42: bool, cf43: bool, cf44: set<int>, cf45: C0) | DC29(cf41: set<string>) | DC31(cf46: D13)
datatype D14 = DC32(cf47: seq<array<bool>>)
datatype D15 = DC33(cf48: multiset<array<D2>>)
datatype D16 = DC35(cf50: int, cf51: char, cf52: int, cf53: int, cf54: int) | DC34(cf49: C2)
datatype D17 = DC36(cf55: array<int>)
datatype D18 = DC37(cf56: C3)
datatype D19 = DC39(cf58: bool, cf59: int) | DC38(cf57: map<int, char>) | DC40(cf60: D19)
datatype D20 = DC41(cf61: seq<T0>)
datatype D21 = DC43(cf63: seq<int>, cf64: bool) | DC42(cf62: map<array<int>, int>)
datatype D22 = DC44(cf65: multiset<char>)
datatype D23 = DC46 | DC47(cf67: int, cf68: bool, cf69: bool, cf70: char) | DC45(cf66: array<char>)
datatype D24 = DC49 | DC50 | DC48(cf71: set<D4>) | DC51(cf72: D24)
datatype D25 = DC53(cf74: map<bool, int>, cf75: int) | DC52(cf73: array<array<bool>>) | DC54(cf76: D25)
datatype D26 = DC56(cf78: int, cf79: int) | DC57(cf80: int, cf81: C9) | DC55(cf77: T0)
datatype D27 = DC59(cf83: bool, cf84: array<bool>, cf85: bool) | DC58(cf82: map<char, bool>)
datatype D28 = DC60(cf86: map<array<int>, multiset<char>>)
class GlobalState {
	var f0 : seq<int>
	const f1 : string
	var f2 : multiset<int>
	var f3 : int
	var f4 : bool
	const f5 : bool
	const f6 : bool
	const f7 : int
	var f8 : int
	constructor (f0 : seq<int>, f1 : string, f2 : multiset<int>, f3 : int, f4 : bool, f5 : bool, f6 : bool, f7 : int, f8 : int) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
		this.f5 := f5;
		this.f6 := f6;
		this.f7 := f7;
		this.f8 := f8;
	}
	
}

function fm0(p0: int, globalState: GlobalState): multiset<bool> {
	multiset{DC6(-|map[true := |(seq(0x31d, i0  => ('i')))|]|, true, !true).cf8 <== !!!!true, "fqkb" < "ocpjiygo", -0x17c != -|['e', 'n']|}
}
function fm1(p0: int, p1: char, p2: bool, p3: map<int, int>, globalState: GlobalState): seq<int> {
	seq(284, i0  => (|((seq(-0x182, i1  => ('y'))) + "xk")|))
}
function fm5(globalState: GlobalState): D0 {
	DC3(if (true) then DC1(|{false}|, 0xbd) else DC3(DC2()))
}
function fm6(p0: int, globalState: GlobalState): int {
	84
}
function fm7(p0: char, globalState: GlobalState): map<bool, char> {
	map['y' in (seq(0x3, i0  => ('b'))) := 'l']
}
function fm11(p0: bool, p1: bool, globalState: GlobalState): map<bool, char> {
	map[[false] <= [!false, true, true, true] := 'k']
}
function fm14(p0: int, globalState: GlobalState): string {
	"g"
}
function fm17(p0: char, globalState: GlobalState): string {
	("elk" + "f") + ("f" + "jflao")
}
function fm18(p0: bool, p1: bool, globalState: GlobalState): char {
	'e'
}
function fm20(p0: bool, globalState: GlobalState): int {
	|("jos" + ((seq(0x1f4, i0  => ('u'))) + "ilgan"))|
}
function fm21(p0: bool, globalState: GlobalState): map<bool, bool> {
	map[false := false]
}
function fm22(p0: int, globalState: GlobalState): multiset<int> {
	multiset{376 + |map["akm" := !true]|}
}
function fm23(p0: string, globalState: GlobalState): seq<bool> {
	match DC19([|DC13(DC13(map[true := false]).cf19).cf19|, 382]) {
		case DC19(cf24) => [true] + DC21(false, [true], true).cf27
	}
}
function fm24(globalState: GlobalState): D3 {
	DC8(map[0xd8 := 0x174])
}
function fm25(globalState: GlobalState): string {
	(seq(-0x2cb, i0  => ('j'))) + ("unujxwuqx" + (seq(0x25, i1  => ('w'))))
}
function fm26(p0: D3, p1: int, p2: multiset<seq<bool>>, globalState: GlobalState): int {
	(-978 / -|(seq(0x125, i0  => ('l')))|) * (|"k"| / 0x138)
}
function fm27(p0: bool, globalState: GlobalState): multiset<int> {
	match DC39(false, 0x6a) {
		case DC39(cf58, cf59) => multiset([|multiset{506, |multiset{cf59}|, cf59, |multiset{cf59, |map[cf58 := cf59]|}|, cf59}|, cf59, -cf59])
		case DC38(cf57) => multiset{|multiset{true, true, true, !false, true}|, |"kh"|, -0x36e, |map[true := -273]|} - multiset{|multiset{939}|, |[-0x78]|}
		case DC40(cf60) => multiset([0x2af]) - multiset{DC0(|{--72}|).cf0, 919}
	}
}
function fm28(p0: int, globalState: GlobalState): set<int> {
	{|[true]|} + {564, |[|[|(set v0 : int | (0x2 <= v0) && (v0 < 0x34c) :: (v0 % |{|[true]|}|))|, 0x115]|, |map[seq(0x3b5, i0  => (|"qthfycdk"|)) := !!false]|, |"heoi"|]|, 0x183, -|multiset{!false}|}
}
function fm29(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): D11 {
	DC25(multiset{false, false, false})
}
function fm30(p0: int, p1: D9, p2: bool, p3: int, globalState: GlobalState): D4 {
	DC11(!(multiset{-0x197} < multiset{708, 0x315, -937}), 0xd2 != -759, 842, DC24(|map[true := -0x2a3]|, true, ['f', 'y']).cf33 + (seq(0x2cb, i0  => ('e'))))
}
function fm31(p0: bool, p1: bool, p2: bool, globalState: GlobalState): seq<bool> {
	[false] + ([true, false] + [true, false])
}
function fm32(p0: int, p1: seq<string>, p2: char, p3: bool, globalState: GlobalState): map<bool, map<bool, int>> {
	match DC23(multiset{"sisutg"}) {
		case DC24(cf31, cf32, cf33) => map[cf32 := map[cf32 := cf31]]
		case DC23(cf30) => map[DC7('n', true, -255).cf10 := map[true := 0x113]] + map[true := map[false := -0x268]]
	}
}
function fm33(p0: bool, p1: bool, globalState: GlobalState): map<seq<bool>, bool> {
	map[[!true, !false] := false] + map[[!true] := false]
}
function fm34(p0: bool, p1: bool, globalState: GlobalState): D0 {
	DC0(0x3e5)
}
function fm35(p0: bool, p1: int, p2: bool, p3: bool, globalState: GlobalState): set<bool> {
	({false, true} - {!true, !true, true, false, DC39(false, 831).cf58}) + ({false} * {!true, !false})
}
function fm36(p0: int, p1: set<bool>, p2: bool, globalState: GlobalState): D10 {
	DC23(multiset{"hx", "kueav"})
}
function fm37(p0: seq<D0>, p1: string, p2: int, p3: int, globalState: GlobalState): map<multiset<bool>, int> {
	((map v0 : multiset<bool> | v0 in multiset{multiset{false, false}, multiset{false}} :: (v0) := (0x2e1)) + map[multiset{true, true} := 0x316]) + (map v1 : multiset<bool> | v1 in map[multiset{true} := |[false]|] :: (v1) := (|map[true := true]|))
}
function fm38(p0: bool, globalState: GlobalState): D0 {
	DC3(DC3(DC3(DC1(-0x2d5, 483))))
}
function fm39(p0: int, p1: bool, p2: bool, globalState: GlobalState): D12 {
	match DC19([|map[!false := -707]|]) {
		case DC19(cf24) => DC26(map[false := 0x30])
	}
}
function fm40(p0: bool, globalState: GlobalState): D2 {
	match DC17(map[|"gndcte"| := !true]) {
		case DC18() => if (true) then DC7('j', !true, |"mhiixh"|) else DC7('p', true, -0x2dd)
		case DC17(cf23) => DC7('v', false, |[-|multiset{true, true}|]|)
	}
}
function fm41(p0: seq<bool>, p1: char, p2: int, globalState: GlobalState): set<string> {
	{"fvh", "sfclqlf" + "vxt"}
}
function fm42(p0: int, p1: int, p2: int, p3: bool, globalState: GlobalState): map<string, bool> {
	match if (true) then DC26(map[false := |[!true]|]) else DC26(map[true := -0x14c]) {
		case DC27(cf36, cf37, cf38, cf39) => (map v0 : string | v0 in (map v1 : string | v1 in multiset{"fup", "xmp"} :: (v1) := (true)) :: (v0) := (true)) + map["mbdnvowfp" := false]
		case DC26(cf35) => map v2 : string | v2 in ["c", "xegsurtcy"] :: (v2) := (true)
		case DC28(cf40) => map["s" := false] + (map v3 : string | v3 in map["rasmegb" := true] :: (v3) := (true))
	}
}
function fm43(p0: bool, p1: int, p2: char, globalState: GlobalState): map<int, char> {
	if ({|"s"|} > {|"k"|}) then map[0x13b := 'o'] + map[|[0x6]| := 's'] else (map v0 : int | v0 in (map v1 : int | v1 in map[-0x217 := |(seq(0x1a9, i0  => (|map[false := true]|)))|] :: (v1 * |"nkdcjuouw"|) := (--371)) :: (v0 * |[!true]|) := ('e')) + map[0x33b := 'b']
}
function fm44(p0: int, globalState: GlobalState): multiset<seq<bool>> {
	multiset{[true], [false, true], [!true]}
}
function fm45(p0: string, p1: int, p2: int, globalState: GlobalState): set<map<D12, int>> {
	(if (!!true) then {map[DC26(map[false := 754]) := 0x299], map[DC26(map[false := 0xca]) := |map[!!true := 198]|]} else {map[DC26(map[false := -0xef]) := 0x268], map v0 : D12 | v0 in {DC26(map[true := |map[false := true]|])} :: (v0) := (|(seq(-0xaa, i0  => ('m')))|)}) - ({map v1 : D12 | v1 in map[DC26(map[true := -|map[-|map[true := 0x14d]| := |[|multiset{DC11(true, false, 0xfb, "uid")}|]|]|]) := map v2 : int | v2 in map[|map[true := 't']| := 't'] :: (v2 + 0x34a) := ('s')] :: (v1) := (0xe6)} * (set v8 : map<D12, int> | v8 in map[map v3 : D12 | v3 in (map v4 : D12 | v4 in (map v5 : D12 | v5 in map[DC26(map[false := |(map v6 : int | (-428 <= v6) && (v6 < 0x212) :: (v6 / -0xce) := (|"daydimu"|))|]) := 'd'] :: (v5) := ("e")) :: (v4) := (|(map v7 : char | v7 in multiset(['j']) :: (v7) := (true))|)) :: (v3) := (0x2f2) := 'm'] :: (v8)))
}
function fm46(globalState: GlobalState): map<bool, int> {
	(map[false := |multiset{!true}|] + map[true := 0x1a8]) + (map[true := -|{!!!true}|] + map[false := 470])
}
function fm47(p0: int, p1: bool, p2: bool, globalState: GlobalState): map<set<int>, string> {
	map[{|['y', 'x']|} := (seq(879, i0  => ('g'))) + "yjv"]
}
function fm48(p0: seq<int>, p1: map<int, bool>, globalState: GlobalState): map<bool, D7> {
	map[if (false) then true else !false := DC18()]
}
function fm49(p0: bool, globalState: GlobalState): bool {
	|multiset{-0x2bd}| > |(seq(959, i0  => (0x39e)))|
}
function fm50(globalState: GlobalState): D3 {
	if (true) then DC9() else DC9()
}
function fm51(p0: int, globalState: GlobalState): D1 {
	DC4(false)
}
function fm52(globalState: GlobalState): map<int, int> {
	match DC25(multiset([true])) {
		case DC25(cf34) => map[0x171 := |[|(set v0 : int | (944 <= v0) && (v0 < 0x2b2) :: (v0 / |map[!false := |"kgpvcbka"|]|))|]|]
	}
}
function fm53(p0: int, p1: int, globalState: GlobalState): set<map<int, bool>> {
	{map[|"kc"| := false], map v0 : int | (-0x2cd <= v0) && (v0 < 0x82) :: (v0 - |(seq(0x2d8, i0  => ('p')))|) := (true)} * {map[|(set v1 : int | (0x27a <= v1) && (v1 < 0x2a6) :: (v1 % |map[false := DC0(|[-|"oghwvum"|]|)]|))| := true]}
}
function fm54(p0: int, globalState: GlobalState): multiset<char> {
	multiset((if (true) then ['x', 'd'] else ['m', 'i']) + (seq(-0x196, i0  => ('p'))))
}
method m15(globalState: GlobalState) returns (r0: int, r1: map<array<int>, multiset<char>>, r2: bool, r3: C0) {
	var v0 := "dpud";
	var v1 := 0x24;
	for i0 := |("bheo" + v0)| to v1 {
		var v2 := true;
		var v3: array<bool> := new bool[1] [v2];
		v3[925] := v2;
		v3[925] := v2;
		globalState.f8 := i0;
		var v4: seq<bool> := [v2, v2];
		var v5: multiset<bool> := multiset{true, v2};
		var v6: seq<int> := [v1, |[v1]|];
		globalState.f3 := (|v4| * |multiset{v1, v1, |v5|, |v5|}|) - v6[v1];
		v2 := !v2;
	}
	globalState.f3 := v1;
	globalState.f3 := |v0|;
	var v7 := false;
	var v9: map<bool, set<int>> := map[v7 := set v8 : int | (-0x388 <= v8) && (v8 < 973) :: (v8 * v1)];
	var v10: seq<int> := [v1, v1, v1, 0x291];
	globalState.f3 := -(if (true) then |v9| * v10[v1] else v1);
	var v11: array<bool> := new bool[21];
	var v12: seq<array<bool>> := [v11];
	v11 := v12[-v1];
	var v13: array<int> := new int[11];
	forall i1 | 0 <= i1 < v13.Length {
		v13[i1] := i1 % (v1 % v1);
	}
	r0 := v1;
	var v14: map<array<int>, multiset<char>> := map[v13 := fm54(v1, globalState)];
	var v15 := DC60(v14);
	r1 := v15.cf86;
	r2 := v7;
	var v16: C0 := new C0();
	r3 := if (v7) then v16 else v16;
}
trait T0 {
	function fm2(globalState: GlobalState): map<int, int> 
	function fm3(p0: seq<int>, p1: multiset<int>, p2: bool, globalState: GlobalState): bool 
	method m0(globalState: GlobalState) returns (r0: bool, r1: bool) 
}

class C0 {
	constructor () {
	}
	
	function fm19(globalState: GlobalState): int {
		(-0x2ae - |(seq(0x238, i0  => ('o')))|) - |multiset{0x211, |multiset{true}|}|
	}
}

class C1 extends T0 {
	constructor () {
	}
	
	function fm2(globalState: GlobalState): map<int, int> {
		map v0 : int | (-491 <= v0) && (v0 < -220) :: (v0 * 0x2b9) := (|([!false] + [!true, true])|)
	}
	function fm3(p0: seq<int>, p1: multiset<int>, p2: bool, globalState: GlobalState): bool {
		(!true <== true) && ([true, !false] == [true])
	}
	method m0(globalState: GlobalState) returns (r0: bool, r1: bool) {
		if (false) {
			var v0 := 0x34f;
			var v1: map<int, int> := map[v0 := v0];
			var v2 := DC8(v1);
			var v3 := true;
			var v4: array<D3> := new D3[13] [v2, if (!v3) then v2 else v2, v2, DC8(v1), v2, v2, v2, v2, fm24(globalState), v2, if (false) then v2.(cf12 := v1) else v2, v2, v2];
			v4[504] := v2;
			v4[504] := v2;
			var v5 := new C0();
			var v6: set<int> := {v0, v0, v0 / v0};
			var v7: map<bool, set<int>> := map[v3 := v6];
			v6 := if (v3) then if (v3 in v7) then v7[v3] else v6 else v6 - v6;
			var v8 := DC9();
			match (v8) {
				case DC9() =>
					var v9: set<bool> := {v3};
					var v10: array<multiset<int>> := new multiset<int>[15](i0 => multiset([v0, v0, -|[v3]|, v0]) + multiset([v0]));
					var v11: seq<int> := [v0, v0, v0, v0, v0];
					v10[577] := multiset(v11);
					var v12: seq<bool> := [v3];
					var v13: multiset<int> := multiset{v0, v0, |v9|, v0};
					var v14 := DC10(v12);
					var v15: multiset<D4> := multiset{v14};
					var v16: map<multiset<D4>, int> := map[v15 := v0];
					var v17: map<int, map<multiset<D4>, int>> := map[v0 := v16];
					globalState.f8, globalState.f8, v9, v10[577], globalState.f8 := |v12|, v0, v9, (v13 - v13) * (v13 * v13), v0 % |(if (-v0 in v17) then v17[-v0] else v16)|;
					var v18 := "uou";
					var v19 := 'r';
					v18, globalState.f4 := v18, (if (v3) then 'j' else v19) !in v18;
					globalState.f3 := v5.fm19(globalState);
					var v20 := DC7(v19, v3, v0);
					var v21 := DC11(v3, v3, v11[v0], ("e")[v0 := v19]);
					var v22 := DC19(v11);
					var v23: map<int, set<int>> := map[v0 := v6];
					var v24: map<int, multiset<int>> := map[-v0 := v13];
					var v25: seq<map<int, multiset<int>>> := [v24];
					var v26: array<int> := new int[28] [v0 / v0, v0, v0, |map[v0 := v0]|, v0 - |(seq(0x126, i1  => (v19)))|, v0, if (v3) then v20.cf11 else v21.cf16, v0, v0, v0, v0, v0, -245, if (-(if (v0 in v1) then v1[v0] else v0) in v1) then v1[-(if (v0 in v1) then v1[v0] else v0)] else v0, -(if (v0 in v10[577]) then v10[577][v0] else v0) - v0, v0 + v0, v0, |v22.cf24|, v0, v0, v0, |v23| / v0, v0 * |"nalyv"|, if (true) then v5.fm19(globalState) else v0, v0, 715, |(seq(-0xef, i2  => (v19)))| / |v25[512]|, -v0];
					v26[376] := v0 / v0;
					v26[376] := -(0x361 - |v18|);
				case DC8(cf12) =>
					var v27: array<int> := new int[15];
					var v28 := "pqitorcr";
					v27[257] := |v28| + v0;
					v27[257] := v0;
					var v29: set<bool> := {v3, v3, v3};
					v3 := {!v3} >= v29;
					globalState.f8 := v27[257];
					var v30: array<map<char, int>> := new map<char, int>[17](i3 => map['a' := v0] + map['n' := -0x134]);
					v30 := v30;
			}
			
			var v31: seq<bool> := [false, v3];
			if (v31[v0]) {
				var v32 := "prrqlk";
				var v33: array<bool> := new bool[16];
				var v34: map<bool, array<bool>> := map[v3 := v33];
				var v35 := 'v';
				globalState.f8 := |((fm1(v0, 'e', v3, v1, globalState))[-7 := |v32|] + fm1(|v34|, v35, false, v1, globalState))|;
				var v36: map<bool, bool> := map[v3 := v3];
				var v37: map<map<bool, bool>, bool> := map[v36 := v3];
				var v38: array<char> := new char[15];
				v38[347] := v35;
				v32, globalState.f3, v37, v38[347], globalState.f8 := fm25(globalState), if (v3) then v0 * (if (v0 in v1) then v1[v0] else v0) else v0, v37 + v37, v35, v0 / (v0 + -v0);
				var v39: multiset<int> := multiset{v0};
				var v40: multiset<bool> := multiset{!v3};
				var v41: seq<int> := [-v0, |v39|, v0, v0, if (v3 in v40) then v40[v3] else -v0];
				globalState.f8 := v41[v0 * v0];
				v3 := v3;
				globalState.f8 := -|map[v35 := v32 == v32]|;
			} else {
				var v42: array<array<array<int>>> := new array<array<int>>[24];
				var v43: array<array<int>> := new array<int>[11];
				v42[183] := v43;
				var v44 := DC14(v0, !v3);
				var v45: seq<D5> := [v44];
				var v46: set<bool> := {v3, v3, v3, v3};
				var v47 := DC20(v46);
				var v48 := "fw";
				var v49: seq<int> := [-|v48|];
				var v50: array<bool> := new bool[19] [v3, v3, v3, v3, v3 || v3, v3, v3, v3, !v3, !v3, v3, v3, v3, v3, v3, !(v46 > (v47.(cf25 := v46)).cf25), v3, (v44.(cf20 := |v49|)).cf21, v3];
				v50[436] := !v3;
				v3, r0, v42[183], v45, v50[436] := v3, "hqv" <= v48, v43, v45 + v45, false;
				var v51 := new C0();
				var v52: array<int> := new int[2];
				v52[931] := v0;
				globalState.f8, globalState.f3, v52[931], v50[436] := 901 * (v0 + v0), v0, v5.fm19(globalState), v50[436];
				var v53: multiset<bool> := multiset{v3};
				v53 := multiset{v3, v50[436], !v3, if (v31[v52[931]]) then v50[436] else v3};
				globalState.f4, r0, v52 := v3, false, v52;
			}
			
		} else {
			var v54 := false;
			if (v54) {
				var v55 := -0x38d;
				var v56: map<int, seq<bool>> := map[v55 := [!true]];
				var v57 := DC10([false, v54]);
				var v58: seq<D4> := [DC10(if (v55 in v56) then v56[v55] else [v54]), v57, v57, v57, v57];
				v58 := v58;
				var v59: seq<int> := [v55];
				var v60: map<bool, bool> := map[v54 := v54];
				var v61: multiset<int> := multiset{-0x32d, v55, |v60|};
				var v62: map<bool, int> := map[fm3(v59, v61, v54, globalState) := v55];
				var v63 := 'y';
				var v64: map<int, char> := map[v55 - |v62| := if (v54) then v63 else v63];
				v64 := v64[v55 * v55 := v63];
				var v65: array<char> := new char[26](i4 => v63);
				v65[51] := v63;
				var v66: array<bool> := new bool[12];
				globalState.f3, v65[51], v66 := v55 + v55, v63, v66;
				var v67 := new C0();
				globalState.f8 := v59[-v55];
			} else {
				var v68: set<bool> := {true};
				var v69: seq<bool> := [v54, {true, v54} !! v68];
				var v70 := 70;
				globalState.f4 := v69[v70];
				r0 := v54;
				var v71: array<char> := new char[28];
				var v72 := 'm';
				v71[358] := v72;
				v71[358] := v72;
				v71[358] := v71[358];
				var v73 := "jcrxtqw";
				v73 := v73;
			}
			
			var v74 := "ilxgf";
			var v75 := 857;
			var v76: map<int, int> := map[|v74| := v75];
			var v78 := 'u';
			v76 := v76[|(map v77 : int | (0x201 <= v77) && (v77 < 938) :: (v77 % |DC23(multiset([v74])).cf30|) := (v54))| / v75 := |(if (v54) then map[v78 := v78] else map[v78 := v78])|];
			var v79: seq<bool> := [v54, v54];
			var v80: set<bool> := {v54};
			var v81: seq<set<bool>> := [if (v54) then {v54} else v80];
			globalState.f8, globalState.f3 := -((v75 / 0x29d) / (|v79| * v75)), |v81|;
			var v82 := DC8(v76);
			var v83: multiset<seq<bool>> := multiset{v79};
			globalState.f3, v75 := v75, fm26(v82, v75, v83, globalState);
			var v84: seq<int> := [v75];
			var v85 := DC15();
			match (if (0x2dd !in v84) then v85 else v85) {
				case DC14(cf20, cf21) =>
					globalState.f4 := cf21;
					var v86: array<int> := new int[23] [cf20, 0x154, cf20, cf20, v75, v75, cf20, v75, fm26(v82, cf20, v83, globalState), -cf20, cf20, 0x4e, v75, v75, cf20, v75, |v74|, cf20, v75, cf20, v75, cf20, -0x23d];
					var v87: map<int, array<int>> := map[v75 := v86];
					var v88: map<multiset<array<int>>, bool> := map[multiset{if (cf20 in v87) then v87[cf20] else v86, v86} := !v54];
					var v89: multiset<array<int>> := multiset{v86};
					cf21 := if ((v89 + v89) in v88) then v88[v89 + v89] else v54;
					r0 := v54;
					globalState.f8 := cf20;
				case DC15() =>
					globalState.f3 := v75;
					globalState.f3 := v75;
					var v90 := DC7(v78, v54, v75);
					v90 := DC7(v78, v54, v75);
					globalState.f3 := --(|v74| + v75);
				case DC13(cf19) =>
					var v91: map<set<bool>, int> := map[v80 := fm26(v82, |v74|, v83, globalState)];
					r0 := v91 == v91;
					var v92: multiset<bool> := multiset{v54};
					r0 := (v75 + |v74|) > |v92|;
					r0 := v54;
					var v93: multiset<int> := multiset{627};
					m13(v75, v93 * v93, !({false} !! v80), globalState);
			}
			
		}
		
		var v94 := true;
		globalState.f4 := v94;
		var v95: array<bool> := new bool[17](i5 => v94);
		v95[782] := if (v94) then v94 else v94;
		var v96: array<char> := new char[17];
		var v97: array<int> := new int[18];
		var v98: map<array<char>, array<int>> := map[v96 := v97];
		v95[782] := v98 == v98;
		var i6 := 0;
		while (v95[782])
			decreases 100 - i6
		{
			if (i6 >= 100) {
				break;
			}
			
			i6 := i6 + 1;
			var v99 := -0xac;
			var v100 := DC19([v99, v99, v99]);
			var v101: multiset<D8> := multiset{DC19([v99]), v100};
			var v102: map<int, multiset<D8>> := map[-v99 := v101];
			var v103 := DC8(map[v99 := v99]);
			var v104: multiset<seq<bool>> := multiset{[true]};
			v102 := v102[fm26(v103, v99, v104, globalState) := v101];
			v99 := v99;
			var v105 := new C0();
			globalState.f8 := (v99 + v99) + v99;
		}
		var v106 := 0xa0;
		var v107: map<int, int> := map[v106 := v106];
		globalState.f8 := (v106 % v106) * (602 * |v107|);
		v97[991] := v106;
		var v108 := "p";
		var v109 := DC6(v106, true, v95[782]);
		v97[991], r1, globalState.f4, v108, r1 := -(v106 % (v106 * -v106)), v95[782], v109.cf8, v108, false;
		r0 := v94;
		r1 := match DC15() {
			case DC14(cf20, cf21) => DC6(v97[991], true, cf21).cf8
			case DC15() => {v95[782]} >= {v94}
			case DC13(cf19) => v94
		};
	}
	method m13(p0: int, p1: multiset<int>, p2: bool, globalState: GlobalState) {
		var v0 := "eahalakud";
		if ((|v0| - 0x96) in p1) {
			var v1: array<int> := new int[4];
			var v2: seq<array<int>> := [v1];
			v2 := v2;
			v1[958] := p0;
			v1[958] := p0;
			v1[958] := v1[958];
			v1 := new int[24](i0 => i0 + v1[958]);
			var v3: seq<int> := [p0, --256];
			var v4 := DC21(false, [p2, p2, fm3(v3, p1, p2, globalState), p2], p2);
			var v5: seq<bool> := [!p2, !fm3(seq(586, i1  => (v1[958])), p1, !false, globalState), p2];
			if ((v4.(cf26 := p2)).cf27 == (v5 + v5)) {
				globalState.f3 := if (p2) then p0 else 655;
				v1 := v1;
				globalState.f4 := p2;
				var v6 := DC6(p0, p0 == p0, p2);
				v6 := v6;
				globalState.f8 := v1[958];
			} else {
				var v7 := new C0();
				var v9: array<map<D1, bool>> := new map<D1, bool>[20](i2 => map[DC4(p2) := p2] + (map v8 : D1 | v8 in {DC4(p2), DC4(false)} :: (v8) := (p2)));
				var v10 := DC4(p2);
				var v11 := DC6(|map[p0 := !p2]|, p2, p2);
				var v12: map<D1, bool> := map[v10 := !v11.cf8];
				v9[378] := v12;
				var v14: set<D1> := {v10, v10};
				v9[378] := map v13 : D1 | v13 in (set v15 : D1 | v15 in v14 :: (v15)) :: (v13) := (p2);
				globalState.f8 := v7.fm19(globalState);
				var v16 := 'l';
				var v17: map<C0, char> := map[v7 := v16];
				var v18: map<int, int> := map[v1[958] := |(seq(-0x331, i5  => (v16)))|];
				var v19: map<bool, bool> := map[p2 := p2];
				var v20: map<map<bool, bool>, char> := map[v19 := v16];
				var v21: array<seq<int>> := new seq<int>[15] [v3, v3, v3, v3, if (!fm3(v3, p1, p2, globalState)) then v3 else v3, v3, [|v17|, v1[958]], seq(0x1b3, i3  => (v1[958])), (seq(817, i4  => (0x30e))) + v3, v3, fm1(v1[958], 'j', p2, v18, globalState) + [v1[958], v1[958]], v3, [-v1[958], |v20|], v3 + v3, v3];
				v21[361] := seq(868, i6  => (v1[958]));
				v21[361] := v3;
				var v22: array<string> := new string[11];
				v22[639] := "fusvcouy";
				v22[639] := "k";
			}
			
		} else {
			var v23: array<int> := new int[2](i7 => i7 - p0);
			var v24 := DC3(DC1(-p0, p0));
			var v25: seq<D0> := [v24];
			v23[119] := |v25|;
			var v26: seq<bool> := [p2, p2 || true, true];
			globalState.f3, v23[119] := (if (p2) then p0 else p0) + p0, |v26|;
			globalState.f4 := p2;
			var v27: set<bool> := {p2};
			var v28 := DC20(v27);
			var v29: seq<D9> := [v28];
			globalState.f8, v29 := -v23[119], v29;
			var v30: seq<int> := [v23[119], p0, v23[119], p0, p0];
			var v31: map<int, int> := map[|{p0}| := p0];
			var v32 := DC8(v31);
			var v33: multiset<seq<bool>> := multiset{v26};
			var v34: array<char> := new char[9](i8 => 't');
			var v35: map<array<char>, int> := map[v34 := -0x25f];
			if (fm3(v30, p1 + multiset{p0, fm26(v32, fm26(v32, v23[119], multiset([v26]), globalState), v33, globalState), |v35|, p0, p0}, p2, globalState)) {
				globalState.f4 := p2;
				var v36: map<bool, int> := map[p2 := v23[119]];
				globalState.f3 := if (p2 in v36) then v36[p2] else p0;
				v0 := v0;
				v0 := v0;
				var v37: set<int> := {|p1|};
				v37 := {v23[119], |map[p0 := |(set v38 : int | (331 <= v38) && (v38 < 658) :: (v38 - |map[|v0| := p0]|))|]|, v23[119]} + v37;
			} else {
				v24 := v24;
				var v39: map<int, seq<int>> := map[0x1d5 := [p0]];
				var v40: array<seq<int>> := new seq<int>[9] [[v23[119]], v30, v30, v30[v23[119] := 0x285], v30, (if (0xb6 in v39) then v39[0xb6] else v30) + v30[p0 := |v0|], v30, v30 + v30, v30];
				v40[744] := v30;
				var v41 := 'n';
				var v42: array<string> := new string[4] ["eypc", (v0[p0 := v41])[v23[119] := v41] + v0, v0 + v0, v0 + v0];
				var v43 := DC7(v41, p2, v23[119]);
				v40[744], v23[119], v42, globalState.f4 := v30, -p0 / (--v23[119] + p0), v42, v43.cf11 !in v31;
				var v44 := m14(p2, map[p2 := p2] == map[true := p2], p2, |v30|, globalState);
				var v45: seq<seq<bool>> := [v26];
				var v46 := DC14(fm26(v32, v23[119], multiset([v26]), globalState), v44);
				var v47: map<D5, int> := map[v46 := p0];
				var v48: map<seq<seq<bool>>, map<D5, int>> := map[v45 := v47];
				v48 := v48[v45 := v47];
				globalState.f4 := 628 != fm26(v32, v40[744][v23[119]], v33, globalState);
			}
			
			var v49 := DC0(v23[119]);
			v23[119], globalState.f8, globalState.f4, v49, globalState.f0 := p0, |(v0 + v0)|, p2, if (p2) then DC0(v23[119]) else v49, v30 + (v30 + v30);
		}
		
		var v50: map<int, bool> := map[p0 := p2];
		v50 := v50[p0 := if (p2) then p2 else p2];
		globalState.f4 := p2;
		var v51: seq<int> := [0x48, p0, p0, p0];
		var v52: map<seq<int>, bool> := map[v51[p0 := p0] := p2];
		var v53: map<map<seq<int>, bool>, bool> := map[v52 := p2];
		if ((if (v52 in v53) then v53[v52] else p2) && p2) {
			var v54 := new C0();
			globalState.f8 := p0;
			if (fm3(v51, p1, p2, globalState)) {
				var v55 := new C0();
				v0, globalState.f3 := "meyetabe", -p0;
				var v56 := new C0();
				globalState.f4 := fm3(v51, p1, if (p2) then p2 else p2, globalState);
				var v57: array<bool> := new bool[3](i9 => p2);
				var v58 := DC14(p0, p2);
				v57[263] := v58.cf21 <==> p2;
				v57[986] := p2;
				var v59: map<D3, multiset<bool>> := map[fm24(globalState) := multiset{p2, p2, p2, p2}];
				var v60: multiset<bool> := multiset{p2};
				var v61: seq<multiset<bool>> := [multiset{p2, p2}, v60];
				var v62 := DC25(v60);
				var v63: map<int, multiset<bool>> := map[p0 := v60];
				var v64: array<multiset<bool>> := new multiset<bool>[27] [v60[p2 := p0] * fm0(p0, globalState), v61[p0], v60 - multiset{true, p2, p2, p2}, if (p2) then v60 else multiset([p2, !p2]), v60, v62.cf34, multiset{p2}, multiset{p2, p2} - v60, v60 * v60, v60, v60, v60, v60, if (p2) then v61[|v51|] else v60, v60, v60, v60, v60 * v60, v60, v60, v60[p2 := p0], v60, if (p0 in v63) then v63[p0] else v60, v60, v60, v60 * v60, v62.cf34];
				v64[392] := fm0(|v0|, globalState);
				globalState.f2, v57[263], v57[986], v59, v64[392] := fm27(true, globalState), p2, p2 <== p2, v59, v60;
			} else {
				var v65: array<bool> := new bool[3](i10 => p0 != p0);
				var v66 := 'h';
				v65[578] := v66 in "ijcbji";
				v65[578] := p2;
				var v67: multiset<bool> := multiset{v65[578], v65[578], p2, v65[578]};
				var v68 := DC6(p0, v65[578], p2);
				var v69: seq<bool> := [p2];
				var v70: array<multiset<bool>> := new multiset<bool>[20] [v67, v67, multiset{v65[578]} - v67, v67, multiset{v65[578]} + multiset{v65[578]}, v67, multiset{p2, v65[578], false, !p2}, v67, v67, if (v65[578]) then v67 else v67, multiset{v65[578]}, v67 - multiset{v68.cf7, p2, v65[578]}, multiset{p2}, v67, multiset{DC21(!!v65[578], v69, !p2).cf28}, v67, v67, multiset(v69), v67, v67];
				v70[293] := v67;
				v70[293] := v67 + multiset{p2};
				var v71: array<int> := new int[3];
				v71[169] := p0;
				var v72: map<int, int> := map[-0x31e := p0];
				v71[169] := |(v72 + v72)| + p0;
				globalState.f8 := p0 * 0x3cf;
				v0 := v0;
			}
			
			v54 := v54;
			var v73: map<bool, int> := map[p2 := |v0|];
			globalState.f3 := |DC26(v73).cf35| + |(v50 + (map v74 : int | (0x3be <= v74) && (v74 < -919) :: (v74 + p0) := (p2)))|;
		} else {
			var v75: set<int> := {p0, -p0};
			if ((v75 * fm28(p0, globalState)) > v75) {
				globalState.f3 := -p0;
				var v76: set<bool> := {p2, p2, p2};
				var v77: multiset<bool> := multiset{p2, true, p2, p2};
				var v78: seq<seq<int>> := [v51];
				var v79: map<bool, bool> := map[p2 := !p2];
				var v80: array<bool> := new bool[16] [p0 > -p0, p2, v76 >= v76, p2, p2, v77 !! v77, p2, true, fm3(v51, multiset(v78[p0]), p2, globalState), p2, p2, p1 !! p1, if (p2 in v79) then v79[p2] else p2, p2, false, !(p2 ==> p2)];
				v80[463] := p2;
				v80[463] := !false;
				var v81: array<map<bool, bool>> := new map<bool, bool>[23];
				v81[820] := v79;
				v81[820] := v79;
				var v82: array<char> := new char[10];
				var v83 := 'x';
				v82[110] := v83;
				v82[110] := v83;
				globalState.f4 := (v77 + DC25(multiset{p2}).cf34) == v77;
			} else {
				globalState.f4 := |v75| !in v50;
				var v84: multiset<bool> := multiset{p2 <== p2, p2};
				v84 := v84;
				var v85 := 'v';
				v85 := v85;
				globalState.f3 := p0;
				var v86: array<int> := new int[16](i11 => i11 / p0);
				v86[852] := -0x167 + p0;
				var v87: seq<bool> := [!!true, p2, p2, p2, p2];
				var v88: multiset<seq<bool>> := multiset{v87, v87};
				var v89: map<bool, bool> := map[p2 := p2];
				var v90: map<int, seq<bool>> := map[|v89| := v87];
				v86[852] := if ((if (p0 in v90) then v90[p0] else v87) in v88) then v88[if (p0 in v90) then v90[p0] else v87] else p0;
			}
			
			globalState.f4 := p2;
			globalState.f0 := v51 + v51;
			globalState.f4 := (if (p2) then p2 else p2) || p2;
			var v91: array<string> := new string[2];
			v91[24] := seq(0x340, i12  => ('x'));
			var v92: array<bool> := new bool[27](i13 => p0 == p0);
			v92[212] := true;
			v91[24], v92, v92[212] := "miiuxq", v92, false;
		}
		
		var i14 := 0;
		while (v0 == (seq(0x42, i15  => ('o'))))
			decreases 100 - i14
		{
			if (i14 >= 100) {
				break;
			}
			
			i14 := i14 + 1;
			var v93: seq<bool> := [!true];
			var v94: multiset<bool> := multiset{v93[813], false};
			var v95 := DC25(v94);
			var v96: array<D11> := new D11[13] [v95, v95.(cf34 := v94), v95, fm29(p0, p0, p2, p0, globalState), if (!p2) then v95 else v95, v95, v95, v95, v95, DC25(v94), v95, v95, v95.(cf34 := multiset{p2})];
			v96[541] := v95;
			v96[541] := v95;
			var v97: seq<multiset<int>> := [p1];
			var v98: array<bool> := new bool[13] [p2, p2, p2, p2, v93[|v50|], p2, p1 <= v97[p0], p2, p2, p2, p2, p2, true];
			globalState.f4, globalState.f0, v98 := p2, v51, v98;
			var v99: map<bool, seq<int>> := map[p2 := v51];
			var v100: map<int, seq<int>> := map[p0 := (if (p2 in v99) then v99[p2] else v51)[p0 := 288]];
			globalState.f0 := if (p0 in v100) then v100[p0] else seq(0x26e, i16  => (p0));
			globalState.f4 := if (p2) then p2 else p2;
		}
		var v101: array<int> := new int[5] [if (p2) then -p0 else p0, p0, |(map[p2 := p2])[p2 := p2]| + 0x2d0, p0, p0];
		v101 := v101;
	}
	method m14(p0: bool, p1: bool, p2: bool, p3: int, globalState: GlobalState) returns (r0: bool) {
		globalState.f8 := p3 - p3;
		for i0 := p3 to p3 {
			var v0 := new C0();
			globalState.f4 := p0;
			globalState.f4 := p2;
			globalState.f3 := -i0;
		}
		if (p1) {
			var v1: array<int> := new int[1](i1 => i1 / p3);
			v1[394] := 199;
			v1[656] := p3;
			var v2: set<int> := {p3};
			var v3: seq<set<int>> := [v2, v2];
			var v4: seq<bool> := [p0, !p1];
			var v5 := "auprui";
			var v6: map<int, int> := map[|v5| := p3];
			var v7 := DC8(v6);
			var v8: map<seq<bool>, D3> := map[v4 := v7];
			var v9 := DC10(v4);
			var v10 := DC12(v9);
			var v11: map<D4, set<int>> := map[v10 := v2];
			var v12 := DC21(p1, v4, !p0);
			var v13: seq<int> := [p3, p3];
			var v14: multiset<int> := multiset{p3, p3, p3, p3, p3};
			var v15: seq<seq<bool>> := [fm31(p1, p1, p1, globalState)];
			var v16: array<bool> := new bool[29] [v3[|v8|] <= (if (DC12(v9).(cf18 := fm30(|v5|, v12, p1, |v5|, globalState)) in v11) then v11[DC12(v9).(cf18 := fm30(|v5|, v12, p1, |v5|, globalState))] else v2), p0, p0, p0, p0, p2, !p1, p0, p1, p2, !p1, !false, fm3(v13, v14, fm3(v13, v14, p2, globalState), globalState), p1, p1, p2, p1 <==> p1, !p1, p1, p3 == |v15|, p1, p2, p1, p1 && p1, p0, !p2, p0, p3 in v13, p2];
			v16[355] := v4[p3];
			globalState.f8, v1[394], v1[656], v16[355], globalState.f4 := p3, |{p3, p3}| / p3, p3, !p2, p1;
			var v17: seq<string> := [v5, v5, v5, v5, v5];
			var v18 := 'y';
			var v19: map<bool, map<bool, int>> := map[p1 := map[!p0 := v1[394]]];
			v16[355] := fm32(p3, v17, v18, p0, globalState) != v19;
			var v20: set<bool> := {p2};
			if ((v20 < v20) || (p2 in v20)) {
				v5 := "mosxmc";
				v1[394] := -(-(p3 / p3) / 0x345);
				var v21: map<array<bool>, char> := map[v16 := v18];
				globalState.f8 := |v21|;
				var v22: map<int, bool> := map[|v17| := p1];
				var v25: set<set<int>> := {v2, set v23 : int | v23 in v22 :: (v23 % |{false, !false, true, false, !true}|), set v24 : int | (280 <= v24) && (v24 < 0x1c3) :: (v24 * p3)};
				v25 := v25;
				var v26: array<map<bool, bool>> := new map<bool, bool>[15];
				v26[333] := map[fm3(v13, v14, !v16[355], globalState) := false];
				v26[333] := map[340 !in v6 := p2];
			} else {
				v1[394] := p3;
				var v27: multiset<seq<bool>> := multiset{v4};
				var v28: seq<multiset<seq<bool>>> := [v27, v27, v27];
				var v29: map<int, bool> := map[fm26(v7, 390, v28[v1[394]], globalState) := p1];
				globalState.f3 := -|v29| - 126;
				var v30: array<map<seq<bool>, bool>> := new map<seq<bool>, bool>[12];
				var v31: map<seq<bool>, bool> := map[[false, v16[355]] := p1];
				v30[154] := v31;
				globalState.f3, v2, v30[154], v18 := 0x189, v2 - {|v5|}, fm33(p1, p0, globalState) + ((map v32 : seq<bool> | v32 in v15 :: (v32) := (p1)) + v31), v18;
				var v33: map<array<bool>, int> := map[v16 := -v1[394]];
				v33 := v33[v16 := v1[394] % p3];
				var v34 := new C0();
			}
			
			var v35: map<int, char> := map[|v20| := v18];
			v1[394] := |v35| / v1[394];
			var v37 := DC16(multiset{p3, |v5|});
			var v38: multiset<seq<bool>> := multiset{v4};
			var v39: map<bool, int> := map[p0 := fm26(v7, |(map v36 : D6 | v36 in [v37, DC16(v14), v37.(cf22 := fm27(p2, globalState))] :: (v36) := (true))|, v38, globalState)];
			v39 := v39[p2 := p3 % v1[394]];
		} else {
			var v40: seq<int> := [p3];
			var v41: multiset<int> := multiset{p3, 0x1ec, p3, p3, p3};
			r0 := fm3(v40, v41, p1, globalState);
			var v42: set<bool> := {p0, p2, p1, p1 && p1};
			var v43: array<seq<bool>> := new seq<bool>[16];
			var v44: map<bool, bool> := map[p0 := p1];
			v42, v43, globalState.f3 := v42 * v42, v43, if ((|v44| % --p3) in v41) then v41[|v44| % --p3] else |map[p3 := p3]|;
			globalState.f8 := p3;
			var v45: array<char> := new char[28];
			v45[777] := 'h';
			var v46 := 's';
			v45[777] := v46;
			globalState.f8 := p3;
		}
		
		r0 := p0;
		var v47 := "mjj";
		var v48: array<bool> := new bool[13] [false, p1, !p0, p0, 'x' !in v47, p1, p2, p3 < (fm34(p0, p1, globalState)).cf0, p2, p1, p0, p0, "vvo" < (seq(0x2be, i2  => ('x')))];
		v48[184] := p1;
		v48[184] := p2;
		var v49: seq<bool> := [v47 < v47, v48[184], v48[184], p2];
		var v50 := 'f';
		var v51: multiset<int> := multiset{-p3};
		var v52: seq<int> := [|v51|];
		var v53 := DC7(v50, p1, |v52|);
		var v54 := DC27(p3, p3, p3, p3);
		v48[184] := v49[(v53.(cf11 := v54.cf37).(cf10 := true)).cf11];
		r0 := p1;
	}
}

class C2 extends T0 {
	constructor () {
	}
	
	function fm2(globalState: GlobalState): map<int, int> {
		((map v0 : int | v0 in map[920 := false] :: (v0 - 0x353) := (|"jf"|)) + map[441 := 0x254]) + map[|map[-204 := |"ee"|]| := 0x3da]
	}
	function fm3(p0: seq<int>, p1: multiset<int>, p2: bool, globalState: GlobalState): bool {
		match DC11(true, true, -|"rqonprou"|, "svtwbn") {
			case DC11(cf14, cf15, cf16, cf17) => false
			case DC10(cf13) => 0x17a != 0x2c4
			case DC12(cf18) => false
		}
	}
	method m0(globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0 := 0x22e;
		globalState.f3 := v0;
		var v2: multiset<int> := multiset{v0};
		if ((multiset{|(map v1 : int | (0x2fc <= v1) && (v1 < 0x206) :: (v1 - v0) := ('u'))|} + v2) > multiset{v0, v0, v0}) {
			var v3 := "xxhq";
			var v4: seq<string> := [v3, seq(-0x18b, i0  => ('c'))];
			var v5: map<string, string> := map["ilkioe" := v4[0x2ea]];
			var v6 := true;
			v3 := if (v3 in v5) then v5[v3] else fm17(fm18(v6, v6, globalState), globalState);
			var v7: seq<int> := [v0, v0];
			var v8: seq<bool> := [fm3(v7, v2, v6, globalState)];
			var v9 := DC6(if (v8[|v3|]) then v0 else 0x1be, v6, v6);
			v9 := v9;
			var v10: set<bool> := {v6};
			globalState.f8 := (v0 + v0) / |v10|;
			if (v6) {
				var v11 := new C0();
				var v12 := new C0();
				var v13: array<bool> := new bool[14];
				v13[665] := v6;
				var v14 := 'j';
				var v15: array<seq<array<bool>>> := new seq<array<bool>>[22];
				var v16: seq<array<bool>> := [v13, v13];
				v15[378] := v16;
				globalState.f8, globalState.f3, v13[665], v14, v15[378] := v0, 0x27b, v6, v14, v16;
				v13[665] := v13[665];
				var v17 := new C0();
			} else {
				var v18 := 'a';
				v3 := (v3[v0 := v18])[fm20(!false, globalState) := v18];
				var v19: map<bool, bool> := map[true := v6];
				var v20: map<bool, map<bool, bool>> := map[v6 := map[v6 := v6]];
				var v21 := DC13(map[v6 := true]);
				var v22: seq<map<bool, bool>> := [v19];
				var v23: array<map<bool, bool>> := new map<bool, bool>[14] [v19, v19, if (v6 in v20) then v20[v6] else map[v6 := v6], v19, map[v6 := v6], v19, v19, v21.cf19, v19, fm21(v6, globalState), v19, v22[848], v19, v19];
				var v24: map<array<map<bool, bool>>, bool> := map[v23 := v6];
				v24 := v24[v23 := v0 == |v3|];
				v0 := v0;
				globalState.f3 := v0;
				v10 := v10;
			}
			
			r0 := !!(|v3| <= v0);
		} else {
			var v25 := false;
			globalState.f3 := if (v25) then v0 else v0;
			var v26: array<bool> := new bool[20];
			v26[319] := v25;
			v26[319] := v25;
			var v27 := "ijeb";
			var v28: seq<int> := [-0x111];
			var v29: seq<seq<int>> := [[|v27|, v0], if (v26[319]) then v28 else v28];
			globalState.f0 := v29[v0];
			v0 := v0;
			var v30: array<seq<bool>> := new seq<bool>[27];
			var v31: map<int, int> := map[v0 := v0];
			var v32: array<int> := new int[11] [v0, v0, |([v26[319], v25] + [v26[319]])|, v0 / fm20(v26[319], globalState), v0 % v0, v0, |v31| + 0x20b, v0, -(v0 % v0), v0 % v0, v0];
			v32[813] := fm20(v26[319], globalState);
			v30, globalState.f3, v32[813], v25, v26 := v30, v0 / v0, v0, v25, if (v0 < v0) then v26 else v26;
		}
		
		var v33 := false;
		var v34 := DC16(v2);
		var v35: map<int, multiset<int>> := map[v0 := multiset{v0, v0}];
		var v36: map<bool, bool> := map[false := v33];
		var v37 := "xebnstm";
		var v38 := DC11(v33, !true, v0, v37);
		var v39: seq<bool> := [!v33];
		var v40: array<bool> := new bool[25] [!v33, v33, v33, v33, v33, !true, v34.cf22 !! (if (0x2a7 in v35) then v35[0x2a7] else (multiset(seq(-318, i1  => (|(seq(0x192, i2  => ('b')))|))))[|v36| := v0]), false, v33, v38.cf14, [v33, v33] <= v39, v33, v33, v33, v33, v33, v33, !v33, v33, v33, v33, v0 == v0, true, v33, v33];
		v40[696] := false;
		var v41: seq<int> := [v0];
		v40[696] := (|v41| - v0) >= v0;
		var v42: array<int> := new int[19](i3 => i3 / -v0);
		v42[923] := |"l"|;
		v42[923] := v0;
		var i4 := 0;
		while (v40[696])
			decreases 100 - i4
		{
			if (i4 >= 100) {
				break;
			}
			
			i4 := i4 + 1;
			r0 := v33;
			globalState.f8 := v0 - v42[923];
			var v43: map<array<int>, seq<int>> := map[v42 := v41];
			var v44: multiset<bool> := multiset{v40[696], v33, v33};
			v41 := if (v42 in v43) then v43[v42] else (v41 + v41)[if (v40[696] in v44) then v44[v40[696]] else 0x2b2 := v42[923]];
			v42[923] := v42[923] * v42[923];
		}
		var v45 := 'h';
		var v46: map<int, array<int>> := map[v0 := v42];
		v42, v45, v45, v42[923] := if (v0 in v46) then v46[v0] else v42, 'b', v45, v42[923] * v0;
		r0 := v33;
		r1 := match v38 {
			case DC11(cf14, cf15, cf16, cf17) => cf17 < cf17[cf16 := v45]
			case DC10(cf13) => v33
			case DC12(cf18) => v40[696]
		};
	}
	method m11(p0: map<int, bool>, globalState: GlobalState) returns (r0: int, r1: bool, r2: int) {
		var v0 := true;
		globalState.f4 := !v0;
		var v1 := 0x9f;
		globalState.f8, r2 := v1, v1;
		if (v0) {
			var v2: set<bool> := {v0};
			v2 := v2 - (if (!v0) then {v0} else v2);
			var v3 := "cooouputi";
			var v4: multiset<int> := multiset{v1};
			var v5: map<bool, int> := map[v0 := |v4|];
			var v6: multiset<int> := multiset{|("ligquqw" + v3)|, v1 * (if (v1 in v4) then v4[v1] else |v5|), v1 / v1};
			globalState.f2 := v4;
			var v7 := 'n';
			var v8: array<char> := new char[9] [v7, if (v0) then v7 else v7, v7, v7, v7, v7, fm18(v0, v0, globalState), v7, if (v0) then 't' else v7];
			v8[437] := 'o';
			v8[437] := v7;
			globalState.f4 := v1 == v1;
			var v9: array<int> := new int[26];
			v9[815] := v1;
			v9[815] := v1 % 0xd9;
		} else {
			var v10: map<int, int> := map[0x15a := v1];
			var v11 := DC8(v10);
			var v12: seq<D3> := [if (v0) then v11 else v11, v11];
			v12 := v12 + v12;
			var v13: array<seq<int>> := new seq<int>[16](i0 => [v1, v1]);
			var v14: seq<int> := [v1];
			v13[828] := v14 + v14;
			var v15: multiset<bool> := multiset{v0};
			v13[828], v0 := [v1, v1, v1, if (v0 in v15) then v15[v0] else v1], v0 && v0;
			var v16: map<bool, int> := map[v0 := v1];
			v16 := v16[v0 := |(p0 + p0)|];
			var v17: map<int, bool> := map[v1 := fm3(v13[828], (multiset{v1})[v1 := |v13[828]|], v0, globalState)];
			var v18 := DC17(map[v1 := v0]);
			v17 := v18.cf23;
			var v19 := "rd";
			var v20 := DC6(v1 + |v19|, !v0, !(v1 < v1));
			match (v20) {
				case DC6(cf6, cf7, cf8) =>
					var v21: seq<bool> := [v0, v0, true, v0, cf8];
					var v22: multiset<seq<bool>> := multiset{v21 + v21};
					v22 := v22[v21 := v1];
					var v23: array<D5> := new D5[12];
					v23[885] := DC15();
					var v24 := DC15();
					v23[885] := v24;
					var v25, v26, v27, v28 := m12(globalState);
					globalState.f8 := v26;
				case DC7(cf9, cf10, cf11) =>
					var v29: multiset<int> := multiset{0x2da};
					var v30: map<int, char> := map[|v29| := cf9];
					var v31: multiset<int> := multiset{0x302, |v30|};
					var v32: map<bool, bool> := map[!v0 := v0];
					var v33: seq<multiset<int>> := [v31, fm22(v1, globalState)];
					var v34: seq<bool> := [v29[0x1fe := |v32|] < v33[|v19|]];
					var v36: map<int, map<int, bool>> := map[v1 := (map v35 : int | v35 in v10 :: (v35 - v1) := (false)) + (map[cf11 := cf10])[v1 := false]];
					v34, globalState.f4, v36 := fm23(v19, globalState), v0, map[cf11 := v17] + v36;
					cf9 := 'a';
					v19 := fm17(cf9, globalState);
					var v37: map<int, seq<int>> := map[cf11 := (seq(-851, i1  => (v1)))[(v20.(cf6 := 0x1c7, cf7 := v0)).cf6 := v1] + v14];
					v37 := v37[-(v1 + v1) := v13[828]];
				case DC5(cf5) =>
					var v38: multiset<int> := multiset{|[|(seq(-644, i2  => ('g')))|]|};
					var v39 := 'h';
					var v40: set<char> := {'y', v39};
					var v41: map<multiset<int>, int> := map[v38 := |v40| + v1];
					v41 := v41[(fm22(v1, globalState))[v1 := v1] := 0x352];
					globalState.f8 := v1 + v1;
					cf5[338] := v0;
					cf5[338] := true;
					r1, globalState.f4, globalState.f0, v0 := !cf5[338], v0, v13[828] + v14, v0;
			}
			
		}
		
		globalState.f4 := true;
		var v42: array<int> := new int[11];
		var v43: set<array<int>> := {v42, v42, v42, v42};
		var i3 := 0;
		while (v42 !in (v43 - v43))
			decreases 100 - i3
		{
			if (i3 >= 100) {
				break;
			}
			
			i3 := i3 + 1;
			globalState.f4 := v0;
			var v44: seq<bool> := [v0];
			var v45: seq<bool> := [v0, v44 <= v44, v0];
			v45 := v44 + v44[v1 := v0];
			var v46: array<string> := new string[26];
			var v47: seq<int> := [v1, |v45|];
			var v48: multiset<int> := multiset{-v1};
			var v49: array<bool> := new bool[25] [!v0, v0, v0, DC4(v0).cf4, false, v0, !v0, v0, !v0, v0, true, v0, v0, false, !v0, v0, v0, !true, v0, v0, v0, v0, v0, v0, fm3(v47, v48, v0, globalState)];
			var v50 := DC5(v49);
			var v51: multiset<bool> := multiset{v0, v0};
			var v52: map<int, multiset<bool>> := map[v1 := v51];
			var v53: set<int> := {v1};
			v46, v0, v50, globalState.f8, r1 := v46, fm3(seq(0x37f, i4  => (v1)), v48, fm3(v47[v1 := v1], v48, false, globalState), globalState), v50, (101 - 0x30e) / v1, fm3(v47[|v52| := v1], multiset{|v53|, v1}, false, globalState);
			if (v0) {
				globalState.f8 := |map[!v0 := v0]| / 945;
				var v54: T0 := new C1();
				var v55: seq<T0> := [v54];
				var v56: seq<seq<T0>> := [v55];
				var v57: seq<seq<T0>> := [v56[v1]];
				globalState.f8 := v1 - --|(v56 + v57)|;
				var v58 := 't';
				var v59: multiset<char> := multiset{v58};
				v59 := (v59 * v59) + v59;
				var v60 := new C1();
				v42[526] := |v45|;
				v42[526], v49, v0 := v1, v49, !v0;
			} else {
				globalState.f4 := {v0} > {!!fm3([v1], multiset(v47), v0, globalState), v0};
				var v61 := new C1();
				var v62: map<int, int> := map[v1 := v1];
				v62 := map[v1 := |v53| % v1];
				globalState.f8 := if (v1 in v48) then v48[v1] else if (v0) then v1 else |v62|;
				v62 := v62 + v62;
			}
			
		}
		var v63: array<bool> := new bool[14](i5 => v0);
		var v64 := "wdjl";
		v63[775] := v64 <= v64;
		v0, v0, r1, v1, v63[775] := v0, v0, v0, (v1 + v1) % v1, v0;
		var v65: map<bool, bool> := map[v0 := v63[775]];
		r0 := if (if (v0 in v65) then v65[v0] else v0) then 704 else v1;
		r1 := v0;
		r2 := v1;
	}
	method m12(globalState: GlobalState) returns (r0: int, r1: int, r2: int, r3: bool) {
		var v0 := 0x390;
		var v1 := true;
		globalState.f4 := v0 != fm20(v1, globalState);
		var v2: array<bool> := new bool[6](i0 => v1);
		v2[787] := v1;
		v2[787] := v1;
		var v3 := "pncslplhn";
		var v4: map<string, int> := map[v3 := v0];
		v4 := v4["momthrphd" := 0x5c];
		var v5: seq<int> := [-0x8b];
		globalState.f0 := v5 + v5;
		var v6: set<string> := {v3};
		var v7 := DC29(v6);
		if (v6 == (v7.cf41 - v6)) {
			globalState.f4 := v2[787];
			v2 := new bool[5];
			var v8: map<int, int> := map[v0 := v0];
			var v9 := DC8(v8);
			var v10: seq<bool> := [false];
			var v11: multiset<seq<bool>> := multiset{v10, v10, v10, v10};
			var v12: multiset<int> := multiset{fm26(v9, v0, v11, globalState)};
			var v13 := DC16(v12);
			globalState.f4 := (v12 > v13.cf22) <== v2[787];
			var v14: multiset<string> := multiset{v3};
			var v15: multiset<bool> := multiset{v2[787], v1, v2[787], v2[787], v1};
			var v16: seq<multiset<int>> := [v12, multiset{|v14|, if (v1 in v15) then v15[v1] else v0}];
			globalState.f8 := |v16|;
			var v17: array<int> := new int[1];
			v17[346] := if (v1) then v0 else v0;
			var v18: set<bool> := {v1};
			var v19 := DC20(v18);
			var v20: seq<set<bool>> := [v18, v18, v18];
			var v21: array<D9> := new D9[26] [v19, DC20(v18), DC20({v1, v1, v1, v1, v2[787]}), v19, v19, v19, v19, DC20(v18), v19, v19, DC20(v18), v19, v19, v19, v19, DC20(fm35(v1, v0, v1, v2[787], globalState)), v19, v19.(cf25 := v18), v19, DC20(v20[v0]), v19, v19, v19, v19, v19, v19];
			v21[366] := v19;
			globalState.f8, v2[787], v17[346], v21[366] := fm26(DC8(v8), v0, multiset{v10, [v1], v10}, globalState) * (v0 - 0x35b), v1, v0, DC20(v18);
		} else {
			var v22: multiset<bool> := multiset{v1};
			var v23: C1 := new C1();
			var v24: map<C1, bool> := map[v23 := v2[787]];
			var v25: map<map<C1, bool>, multiset<bool>> := map[v24[v23 := v2[787]] := multiset{!v2[787], v2[787]}];
			r0, r1, r3, v3 := |v3|, 0x10, (v22 - (if (v24 in v25) then v25[v24] else v22)) <= v22, v3 + (seq(0x133, i1  => ('h')));
			globalState.f8 := v0;
			var v26 := 'n';
			var v27: map<bool, bool> := map[fm17(v26, globalState) <= "isrckjn" := v2[787]];
			v27 := v27[v1 := v2[787]];
			var v28: array<set<map<bool, int>>> := new set<map<bool, int>>[21];
			var v29: map<bool, int> := map[v2[787] := v0];
			var v30: set<map<bool, int>> := {v29};
			v28[779] := v30;
			v28[779] := (v30 + v30) - v30;
			var v31: map<int, bool> := map[v0 := v1];
			var v32 := DC17(v31[-0x67 := v1]);
			match (v32) {
				case DC18() =>
					globalState.f8 := v0 / v0;
					var v33 := new C1();
					var v35: seq<bool> := [v2[787], v2[787]];
					var v36: multiset<seq<bool>> := multiset{v35};
					var v37: set<int> := {|(seq(0x6b, i2  => ('n')))|};
					var v38: set<int> := {|v37|, v0, v0, |v6|, v0};
					var v39: multiset<int> := multiset{v0};
					var v40: array<int> := new int[23] [v0, -(v0 + v0), v0, v0, v0, v0 % v0, |(map v34 : seq<bool> | v34 in v36 :: (v34) := (v0))|, |v37|, v0 / 0x173, v0, v0, v0, |v29|, v0, |map[v0 := v1]|, -0x1a2, v0 + |[v2[787], true, v2[787], v1, v1]|, v0, v0, v5[v0], if (v0 in v39) then v39[v0] else v0, v0, if (v2[787] in v29) then v29[v2[787]] else |v35|];
					v40[473] := v0;
					var v41 := DC14(|fm17(v26, globalState)|, v1);
					v40[473] := -(-(v0 - |v3[v41.cf20 := v26]|) % v0);
					var v42 := new C1();
				case DC17(cf23) =>
					v26 := v26;
					r1 := -0x15a;
					v31 := v31;
					var v43: array<set<int>> := new set<int>[11];
					var v44 := DC24(-90, v1, ['q']);
					var v45: set<int> := {0xee, |multiset{v0}|};
					var v46: set<int> := {v0, |v45|};
					var v47: C0 := new C0();
					var v48 := DC30(v1, v44.cf32, v46, v47);
					v43[529] := (v48.(cf42 := v1)).cf44;
					v43[529] := v46;
			}
			
		}
		
		var v49: seq<bool> := [v2[787], true, v2[787]];
		var v50 := DC10(v49);
		v50 := v50.(cf13 := v49);
		r0 := v0;
		r1 := |(if (v2[787]) then v49 else [v2[787], v2[787], v2[787], v2[787]] + [v2[787]])|;
		var v51 := DC0(v0);
		r2 := match v51.(cf0 := v0) {
			case DC1(cf1, cf2) => |(map[v2[787] := cf2] + map[v2[787] := |map[cf2 := |multiset{cf2, cf2}|]|])|
			case DC2() => v0
			case DC0(cf0) => cf0 * --0x1b9
			case DC3(cf3) => v0
		};
		var v52: array<C0> := new C0[18];
		var v53: seq<array<C0>> := [v52];
		r3 := v53[v0] in v53;
	}
}

class C3 extends T0 {
	constructor () {
	}
	
	function fm2(globalState: GlobalState): map<int, int> {
		(map[-0x36e := 0x2e4] + map[0x297 := |[!true, false]|]) + (map v0 : int | v0 in map[|(map v1 : seq<D0> | v1 in {[DC3(DC2())], [DC3(DC3(DC2())), DC3(DC0(|"t"|)), DC3(DC0(0x25b))]} :: (v1) := (false))| := true] :: (v0 / |"idrigq"|) := (|(seq(-517, i0  => ('e')))|))
	}
	function fm3(p0: seq<int>, p1: multiset<int>, p2: bool, globalState: GlobalState): bool {
		!!false
	}
	function fm15(p0: int, globalState: GlobalState): int {
		|multiset{map[0x191 := 0x32b] + map[-|"ifkhv"| := |map[|multiset{false, false}| := false]|]}|
	}
	function fm16(p0: bool, globalState: GlobalState): bool {
		if (true) then true <==> !true else |map[true := ['p', 'k', 'j', 'u', 'x']]| != |[false]|
	}
	method m0(globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0: array<bool> := new bool[29](i0 => !false);
		var v1 := true;
		v0[239] := v1;
		var v2 := 0x172;
		v0[239] := v2 > -v2;
		for i1 := v2 to |(seq(-0x39e, i2  => ('n')))| {
			var v3 := "fm";
			var v4: seq<string> := [v3];
			if ((v4[v2] + v3) <= v3) {
				var v5 := DC1(i1, i1);
				globalState.f8 := -v5.cf1;
				var v6 := 'w';
				v6 := v6;
				v3 := (v3 + v3)[v2 := v6];
				v0[239] := !v1;
				var v7 := new C0();
			} else {
				var v8: map<int, int> := map[i1 := v2];
				var v9 := DC24(if (i1 in v8) then v8[i1] else v2, v0[239], v3);
				r0 := v9.cf32;
				globalState.f8 := v2;
				globalState.f8 := v2;
				var v10: set<int> := {i1, 0x387, -v2};
				var v11: C0 := new C0();
				var v12 := DC30(true, v1, v10, v11);
				r1 := v12.cf42;
				var v13: map<bool, int> := map[v2 > v2 := v11.fm19(globalState)];
				v13 := (v13 + v13) + map[v1 := i1];
			}
			
			var v14: map<int, bool> := map[v2 := v1];
			var v15: seq<map<int, bool>> := [v14];
			var v16: map<bool, int> := map[v15 < v15 := i1];
			v16 := v16[false := v2 - v2];
			v0[239], globalState.f8, globalState.f4, r0, globalState.f4 := !(i1 > i1), i1, v1, v1, v0[239];
			if (v0[239]) {
				var v17: set<string> := {v3, "dfqsa", v3 + v3, v3, v3};
				v17 := {v3} + v17;
				v2, v1, globalState.f8 := i1, v0[239], v2;
				var v18: map<int, int> := map[i1 := i1];
				var v20: array<seq<bool>> := new seq<bool>[16];
				var v21: map<array<seq<bool>>, bool> := map[v20 := false];
				var v22: map<set<int>, bool> := map[set v19 : int | v19 in v18 :: (v19 * 0x298) := if (v20 in v21) then v21[v20] else v1];
				v22 := v22;
				var v23 := m10(i1, v14 + v14, 0x149, globalState);
				var v24: seq<int> := [-v2, --v2, |v3|];
				var v25: array<int> := new int[6] [v2, v24[i1], i1, i1, v2, v2];
				var v26 := DC8(v18);
				var v27: seq<bool> := [v0[239]];
				var v28: multiset<seq<bool>> := multiset{v27};
				globalState.f3, v0[239], v25, v3, r0 := (v2 * fm26(v26, v2, v28, globalState)) / 853, v0[239], v25, (v3 + v3) + v3, false;
			} else {
				v1 := v1;
				v2 := fm20(v0[239], globalState);
				var v29: array<int> := new int[11];
				v29[389] := i1;
				v29[389] := |v3|;
				v29 := v29;
				var v30 := 'r';
				v30 := v30;
			}
			
		}
		if (v0[239]) {
			globalState.f8 := (v2 - v2) - --(0x271 % 0x3e3);
			var v31: multiset<int> := multiset{v2, v2};
			var v32: set<int> := {v2};
			var v33: C0 := new C0();
			var v34 := DC30(v1, true, v32, v33);
			v2, globalState.f4, v0[239], v1 := --(if (v0[239]) then v2 else 0x379), v31[v2 := v2] >= v31, v34.cf42, v2 == v2;
			var v35: array<array<int>> := new array<int>[18];
			var v36: array<int> := new int[28];
			v35[243] := v36;
			v35[243] := v36;
			var v37: set<bool> := {v1};
			v37 := v37 - {v0[239], v0[239]};
			if (v1 || v1) {
				globalState.f4 := !v0[239];
				r1 := fm16(v1, globalState);
				var v38 := new C1();
				var v39 := DC2();
				var v40 := DC21(v0[239], [v0[239], v0[239], v1, v0[239]], v1);
				v39, globalState.f3, r0, r1 := v39, -v2, if (v40.cf27 < [v1, v0[239], false]) then v2 <= v2 else v0[239] && false, v1;
				var v41: array<char> := new char[23](i3 => 'b');
				var v42: seq<array<char>> := [if (v0[239]) then v41 else v41, v41, v41];
				v42, globalState.f8 := (v42 + v42) + v42, v2 % 776;
			} else {
				var v43 := "ae";
				var v44: map<string, int> := map[v43 := v2];
				v44 := v44;
				var v45 := 'g';
				var v46 := DC9();
				var v47: map<char, D3> := map[v45 := v46];
				v47 := v47[v45 := v46];
				var v48: map<bool, int> := map[!v1 := v2];
				var v49: map<int, int> := map[v2 := |v48|];
				var v50: seq<int> := [|v48[fm16(v0[239], globalState) := |v43|]|, |v49|];
				globalState.f4 := fm3(v50, v31 - v31, v1, globalState);
				v36[522] := |v49|;
				v36[522] := v2 % v50[v2];
				var v51: seq<array<int>> := [v35[243], if (v1) then v35[243] else v35[243], v35[243], v35[243]];
				v51 := if (true) then v51 + v51 else v51;
			}
			
		} else {
			var v52: array<map<int, bool>> := new map<int, bool>[5];
			var v53: map<int, array<map<int, bool>>> := map[v2 := v52];
			v53 := v53[v2 := v52];
			var v54 := "aeeyibtfx";
			var v55: multiset<string> := multiset{v54};
			var v56 := DC23(v55);
			v56 := (if (v1) then fm36(|v54|, {v0[239]}, v1, globalState) else DC23(v55)).(cf30 := v55);
			globalState.f2 := fm27(v1, globalState);
			var v57: seq<int> := [-0x24d];
			var v58: multiset<int> := multiset{|"wc"|, |v57|};
			var v59: seq<bool> := [v0[239]];
			var v60 := DC24(v2, v1, v54);
			var v61 := DC21(v1, v59, v60.cf32);
			v0[239] := (if (v0[239]) then fm3(v57, v58, v1, globalState) else v61.cf26) || false;
			var v62 := 'o';
			globalState.f3 := v2 + |v54[v2 := v62]|;
		}
		
		globalState.f3 := |"h"|;
		for i4 := v2 to v2 {
			globalState.f3 := i4 % |multiset{v1}|;
			r0 := v0[239];
			var v63 := new C2();
			globalState.f3 := i4;
		}
		var v64 := DC7('l', v0[239], fm20(v0[239], globalState));
		match (v64) {
			case DC6(cf6, cf7, cf8) =>
				var v65: map<bool, bool> := map[v1 := v1];
				var v66: multiset<bool> := multiset{!v1, cf7, cf7, if (cf7 in v65) then v65[cf7] else !false};
				var v67: map<multiset<bool>, int> := map[v66 + v66 := cf6];
				var v68 := DC1(cf6, cf6);
				var v69 := DC3(v68);
				var v70: seq<D0> := [v69, v69, DC3(v68), fm38(v1, globalState), DC3(v68)];
				var v71 := 'o';
				v67 := fm37(v70, fm17(v71, globalState), v2, v2 + cf6, globalState);
				var v73: map<int, int> := map[v2 := cf6];
				var v74: set<map<int, int>> := {(map v72 : int | (0x8f <= v72) && (v72 < -189) :: (v72 * |[v0[239], !cf7, !v1, cf8]|) := (v2)) + v73};
				var v75: seq<set<map<int, int>>> := [v74];
				v74 := v75[cf6];
				var v76: array<int> := new int[25];
				v76 := v76;
				var v77: seq<bool> := [if (v0[239]) then cf8 else v0[239]];
				var v78: array<char> := new char[4];
				var v79: set<array<char>> := {v78};
				var v80: set<bool> := {cf7, v0[239]};
				v77 := fm31({v78, v78, v78} > v79, v80 >= {v77[-cf6], cf8}, !(if (cf7) then v0[239] else v1), globalState);
			case DC7(cf9, cf10, cf11) =>
				var v81: map<bool, char> := map[cf10 := cf9];
				v81 := v81[fm16(cf10, globalState) := cf9];
				globalState.f4 := false;
				var v82 := new C1();
				var v83 := new C0();
			case DC5(cf5) =>
				var v84: map<int, int> := map[v2 := v2];
				var v85 := DC8(v84);
				v85 := DC8(v84);
				var v86 := 't';
				v86 := v86;
				var v87: seq<string> := ["exys"];
				var v88: map<int, string> := map[0x3a6 := v87[v2]];
				var v89: set<char> := {v86};
				var v90 := "kjjgia";
				v88 := v88[v2 + |v89| := v90];
				globalState.f8 := v2;
		}
		
		var v91: set<bool> := {v1};
		r0 := v91 !! v91;
		r1 := v1;
	}
	method m10(p0: int, p1: map<int, bool>, p2: int, globalState: GlobalState) returns (r0: D3) {
		var v0 := false;
		var v1 := 'w';
		var v2: seq<char> := [v1, v1];
		var v3 := DC24(p2, v0, v2);
		if (v3.cf32) {
			if (v0 || false) {
				var v4: array<bool> := new bool[5];
				v4[423] := v0;
				v4[423] := v0;
				var v5: array<array<bool>> := new array<bool>[15];
				var v6: array<int> := new int[3];
				v6[458] := p0;
				var v7: seq<int> := [p0];
				var v8: multiset<int> := multiset{-0x117};
				v4[423], v5, v6[458], globalState.f3 := fm3(v7 + v7, v8, v0, globalState), v5, fm15(|(v2 + "x")|, globalState), p0;
				globalState.f0 := v7;
				v4[423] := v4[423];
				var v9: map<int, int> := map[p2 := p2];
				var v10: map<map<int, int>, bool> := map[v9 := v4[423]];
				globalState.f0 := v7 + (v7 + v7)[-|v10| := p0];
			} else {
				globalState.f8 := (p2 - p0) / p0;
				globalState.f8 := p2;
				var v11: seq<bool> := [v0, v0, !v0];
				var v12 := DC10(v11);
				var v13: seq<D4> := [DC12(v12)];
				var v15 := DC22(v11);
				var v16: map<int, int> := map[|[v15]| := |v2|];
				var v17: map<map<int, int>, bool> := map[v16 := fm16(v0, globalState)];
				var v18: map<seq<D4>, map<map<int, int>, int>> := map[v13 := (map v14 : map<int, int> | v14 in v17 :: (v14) := (p2))[v16 := p0]];
				var v19: map<map<int, int>, int> := map[v16 := p0];
				v18 := v18[v13 := v19];
				var v20: multiset<int> := multiset{fm20(v0, globalState)};
				globalState.f2 := v20 - v20;
				var v21: map<bool, int> := map[v0 := |v2|];
				var v22 := DC26(v21);
				var v23: array<D12> := new D12[18] [v22, DC26(map[v0 := p2]), v22, v22, v22, v22, v22, if (v0) then v22 else v22.(cf35 := map[v0 := p2]), DC26(v21).(cf35 := v21), DC26(map[v0 := p2]), fm39(p2, v0, v0, globalState), v22, v22, v22, v22, fm39(0x268, v0, v0, globalState), v22, v22];
				v23[54] := v22;
				globalState.f3, v23[54], globalState.f8 := p0 + |v2|, DC26(map[v0 := p2]), |v2|;
			}
			
			v0 := v0;
			var v24: multiset<bool> := multiset{false, v0, v0, v0};
			var v25: map<set<int>, bool> := map[fm28(p0, globalState) := v24 !! v24];
			var v26: map<bool, int> := map[v0 := 0x175];
			var v27: array<int> := new int[26];
			var v28: set<array<int>> := {v27, v27};
			var v29: set<int> := {0x365, 0x3b7, p2, |v26|, |v28|};
			if (if (v29 in v25) then v25[v29] else v0) {
				var v30 := new C2();
				var v31: seq<int> := [659, p0];
				var v32: multiset<int> := multiset{p0};
				v0 := fm3(v31 + v31, v32, v0, globalState);
				var v33: set<bool> := {v0, v0};
				var v34 := DC20(v33);
				v0, globalState.f4 := v0 || v0, v33 > v34.cf25;
				var v35 := DC16(v32);
				var v36: array<bool> := new bool[22];
				v36[637] := v29 >= v29;
				v35, v0, v0, v36[637] := v35, false, v0, v0;
				var v37: seq<bool> := [v36[637]];
				var v38: map<int, int> := map[p0 := p2];
				v37 := [v0] + (v37[|fm1(p0, v1, v36[637], v38, globalState)| := v36[637]] + fm31(true, false, v0, globalState));
			} else {
				v0 := v24 !! (v24 - v24);
				var v39: array<bool> := new bool[18];
				var v40: map<bool, array<bool>> := map[true := v39];
				var v41: seq<array<bool>> := [v39];
				v40 := v40[v0 := v41[p2]];
				globalState.f4 := v0;
				v0 := 66 != p2;
				var v42: set<bool> := {v0};
				var v43: array<set<bool>> := new set<bool>[3] [v42, v42, {v0}];
				var v44 := DC11(true, v0, -888, v2);
				var v45: map<array<set<bool>>, int> := map[v43 := |[false, v44.cf14, v0, v0]|];
				v45 := v45[v43 := p2];
			}
			
			globalState.f3 := p2;
			var v46 := DC15();
			match (v46) {
				case DC14(cf20, cf21) =>
					var v47: map<bool, bool> := map[cf21 := v0];
					var v48: seq<int> := [p2, -p2];
					var v50 := DC27(|v2|, |v47|, p0, |(set v49 : int | v49 in v48 :: (v49 + |multiset([false])|))|);
					var v51 := DC28(v50);
					var v52: set<bool> := {cf21};
					globalState.f3, globalState.f8, v24, v2, v51 := (p0 % p0) + cf20, |v48| % (|v52| / |v24|), fm0(p2, globalState), if (-(if (v0 in v26) then v26[v0] else 0x1cd) > 15) then seq(0x2e9, i0  => (v1)) else v2, v51;
					v47 := v47[cf21 ==> cf21 := v0];
					globalState.f8 := cf20 % (cf20 % -|"y"|);
					var v53: array<bool> := new bool[7];
					var v54: map<bool, array<bool>> := map[v0 := v53];
					v54 := v54[v0 := v53];
				case DC15() =>
					v27[113] := p2;
					var v55: array<bool> := new bool[4];
					var v56: seq<array<bool>> := [v55];
					var v57 := DC32([v55, v55]);
					var v58 := DC7(v1, v0 || v0, |(v56 + v57.cf47)|);
					v27[113], v2, v58 := p0, v2 + (v2 + v2), if (v0) then fm40(!v0, globalState) else v58;
					var v59: seq<int> := [v27[113]];
					var v60: array<seq<int>> := new seq<int>[3] [v59, seq(0x372, i1  => (p0)), v59];
					v60[195] := v59;
					v60[195] := v59;
					var v61: map<bool, seq<int>> := map[v0 := v59];
					globalState.f4, v1, globalState.f3, v60[195] := p0 >= v27[113], v1, p0, if (v0 in v61) then v61[v0] else v60[195];
					var v62 := new C1();
				case DC13(cf19) =>
					var v63: array<D12> := new D12[13];
					var v64 := DC26(map[false := p0]);
					v63[145] := v64;
					v63[145] := v64;
					v0 := v0;
					var v65: array<bool> := new bool[20];
					v65[561] := v0;
					v65[561] := v0 && v0;
					globalState.f3, v0, v0, globalState.f8 := p0, v65[561], |p1| <= p0, -p0;
			}
			
		} else {
			var v66: map<bool, int> := map[!v0 := p2];
			v66 := v66[v0 := p0];
			var v67: set<bool> := {v0, !(if (p0 in p1) then p1[p0] else v0), false, v0};
			var v68: map<set<bool>, bool> := map[{v0, v0, v0, v0} := v67 > v67];
			var v69 := DC20(v67);
			v68 := v68[v69.cf25 := v0];
			var v70: map<int, int> := map[|v2| := fm15(p0, globalState)];
			var v71: seq<int> := [p0];
			var v72: array<D2> := new D2[14];
			var v73: set<array<D2>> := {v72};
			var v74: array<int> := new int[25] [|{v0}|, p0, p0, |multiset{|map[true := p2]|}| + p2, -628 + p2, |[-p2, p2]|, p0, p0, p0, -p2, |{v1, v1, 'r'}|, if (p0 in v70) then v70[p0] else p2, -0x84, p0, p2, p0 * p2, p0, if (v0) then p2 else 0x13f, 0x25f, -(p0 + p0), p2, p0, |v71|, |v73|, p2];
			v74 := v74;
			var v75: array<bool> := new bool[17];
			v75[779] := v0;
			var v76: multiset<array<D2>> := multiset{v72};
			var v77: set<int> := {p0};
			v75[779] := |(v76 * DC33(v76).cf48)| !in v77;
			if (v0) {
				var v78 := new C0();
				v70 := v70[|v67| + p2 := p0];
				var v79: seq<bool> := [false];
				v0 := v79[p2];
				globalState.f4 := v79[p0];
				var v80: map<bool, bool> := map[v75[779] := v75[779]];
				globalState.f3, globalState.f4, v77 := p2, v75[779], DC30(if (v0 in v80) then v80[v0] else v75[779], v0, v77, v78).cf44;
			} else {
				v70 := v70 + v70;
				var v81 := DC25(fm0(-p2, globalState));
				var v82: multiset<bool> := multiset{v75[779]};
				v81 := DC25(v82 - v82);
				var v83 := new C2();
				var v84: seq<bool> := [fm16(false, globalState)];
				var v85: seq<bool> := [v0, if (v84[fm20(v75[779], globalState)]) then true else v75[779], v75[779]];
				v0 := v84[p2];
				v2 := v2;
			}
			
		}
		
		var v86: seq<bool> := [v0];
		match (DC11(v0, v0, p0, v2).(cf16 := |v86|, cf17 := v2)) {
			case DC11(cf14, cf15, cf16, cf17) =>
				var v87: array<seq<char>> := new string[5](i2 => cf17);
				v87[374] := v2;
				v87[374] := fm25(globalState);
				var v88 := DC22(v86);
				var v89: seq<int> := [-p0];
				var v90: map<int, int> := map[cf16 := p0];
				var v91: seq<seq<bool>> := [[cf15]];
				var v92: multiset<seq<bool>> := multiset{v86, v91[fm20(cf14, globalState)], v86};
				globalState.f0, globalState.f8, v88 := v89, fm26(DC8(v90).(cf12 := v90), cf16 + p0, v92, globalState), v88;
				var v93: array<bool> := new bool[9];
				var v94 := DC11(cf14, cf15, |cf17|, v2);
				v93[634] := v94.cf14;
				v93[634] := if (v0) then !cf14 else !(v86 <= v86);
				var v95: multiset<bool> := multiset{v0};
				var v96: C2 := new C2();
				var v97 := DC34(v96);
				var v98: map<C2, bool> := map[v97.cf49 := true];
				var v99: T0 := new C1();
				var v100: multiset<int> := multiset{p0};
				var v101: map<T0, multiset<int>> := map[v99 := v100];
				var v102: array<int> := new int[27] [-p0, p2, if (false in v95) then v95[false] else |v98|, cf16, |[p0, p0]|, 0x8d, cf16, p2 * |multiset{cf14}|, p2 + p2, p0, p0, cf16 + cf16, -0x353, |(if (v99 in v101) then v101[v99] else v100)|, --p0, cf16, -421 - p2, p0, |v90| + -0x18, |v89|, -(cf16 * p0), p0, p0, cf16, p2, p2, cf16];
				v102[367] := v89[p0];
				v102[367] := p2 + p0;
			case DC10(cf13) =>
				var v103 := new C0();
				var v104 := DC2();
				match (v104) {
					case DC1(cf1, cf2) =>
						var v105: array<D13> := new D13[13];
						var v106: set<int> := {0x112, cf2, p2, 0x34d, cf2};
						var v107 := DC30(false, v0, v106, v103);
						v105[562] := v107;
						var v108: seq<C0> := [v103, v103];
						v105[562] := DC30(v0, v0, {-0x104, p2}, v108[cf2]);
						var v109: array<bool> := new bool[17];
						v109[964] := v0;
						v109[964] := v1 !in v2;
						var v110: map<int, int> := map[p2 := 812];
						var v111 := DC8(v110);
						var v112: multiset<seq<bool>> := multiset{v86};
						globalState.f8 := fm26(v111.(cf12 := v110), |v2|, v112, globalState);
						var v113 := DC4(!v109[964]);
						var v114: map<D1, bool> := map[v113 := v109[964]];
						var v115: seq<int> := [|v2|, p2, cf1];
						v114 := v114[v113 := v115 <= [|cf13|, p2]];
					case DC2() =>
						var v116: array<int> := new int[8];
						var v117: array<array<int>> := new array<int>[5] [v116, v116, v116, v116, v116];
						v117[143] := v116;
						var v118: seq<int> := [p0];
						var v119: map<int, int> := map[p0 := p0 * v118[|v2|]];
						var v120: array<seq<char>> := new seq<char>[24](i3 => (seq(0x357, i4  => (v1))) + ['m', 'b', v1, v1, 'f']);
						v120[206] := v2;
						var v121 := DC36(v116);
						globalState.f3, globalState.f8, v117[143], v119, v120[206] := -|(fm1(p0, v1, v0, v119, globalState))[p0 := 0x3cd]|, |v2| - |v86|, v121.cf55, v119, v2;
						var v122: array<bool> := new bool[13] [v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, fm16(v0, globalState), v0, v0];
						var v123: map<int, array<bool>> := map[p0 := v122];
						v123 := v123 + v123;
						v2 := v120[206];
						globalState.f4 := v0;
					case DC0(cf0) =>
						globalState.f4 := v0;
						var v124: array<string> := new seq<char>[16](i5 => v2);
						v124[380] := "rwhy";
						var v125: map<bool, char> := map[v0 := v1];
						var v126: map<string, bool> := map[v2 := v0 !in v125];
						var v127: array<int> := new int[8];
						v127[530] := p0;
						var v128 := DC36(v127);
						var v129: seq<D17> := [v128];
						v124[380], v126, v127[530], globalState.f4 := v2, v126, |(v86 + (cf13 + fm31(v0, v0, v0, globalState)))[|(if (v0) then v129 else v129)| := v0]|, v0;
						globalState.f4 := v0;
						var v130: seq<int> := [cf0, v127[530]];
						var v131: multiset<int> := multiset{cf0};
						var v132: multiset<multiset<int>> := multiset{v131};
						var v133: map<seq<int>, int> := map[seq(-174, i6  => (p0)) := -cf0];
						var v134: map<bool, bool> := map[false := v0];
						var v135: map<bool, bool> := map[v0 := if (v0 in v134) then v134[v0] else fm3(v130, v131, v0, globalState)];
						var v136: map<int, int> := map[p0 := |v135|];
						var v137: set<map<int, int>> := {(map[cf0 := p2])[|v133| := p0], v136};
						var v139: array<bool> := new bool[20] [v0, !v0, v130 < [v127[530]], v132 !! v132, v0, v0, v0, v137 >= {map v138 : int | (-0x2fd <= v138) && (v138 < 0x1b1) :: (v138 - v127[530]) := (p0), v136}, true, !!v0, v0, v0, |[cf0, p0]| <= v127[530], v0, true, !!v0, v0, v0, {cf0} !! {p2, p2, |v124[380]|}, v0];
						v139[698] := v0;
						var v140 := DC27(p0, cf0, |map[v0 := {p0, cf0, p2}]|, p0);
						var v141: seq<D12> := [v140];
						v139[698] := v141 != (v141 + v141);
					case DC3(cf3) =>
						var v142: array<bool> := new bool[10];
						v142[866] := DC14(p0, v0).cf21;
						v142[866] := v0;
						v0 := v142[866];
						var v143: set<bool> := {v142[866], v142[866]};
						var v144: array<int> := new int[10] [p0, p2, |(seq(0x304, i7  => (v1)))|, p2, |{p0}|, p2, p2, |v2|, |cf13|, |v143|];
						var v145: map<array<int>, array<bool>> := map[v144 := v142];
						var v146: seq<map<array<int>, array<bool>>> := [v145, v145];
						var v147: array<map<array<int>, array<bool>>> := new map<array<int>, array<bool>>[26] [if (v0) then map[v144 := v142] else v145, v145, v145 + v146[p0], v145, if (v142[866]) then v145 else v145, v145, v145, map[v144 := v142] + v145, map[v144 := v142], v145, v145[v144 := v142] + v145, v145[v144 := v142] + map[v144 := v142], v145 + v145, v145[v144 := v142] + v145, v145, v145, v145, map[v144 := v142], v145, v145, v145, v145 + v145, v145 + map[v144 := v142], v145, v145, v145];
						v147[257] := map[v144 := v142];
						var v148: multiset<bool> := multiset{v142[866]};
						var v149: multiset<int> := multiset{p2, p2};
						var v150 := DC16(v149);
						var v151: set<D6> := {v150};
						v142[866], v147[257] := -(|v148| / |v151|) > 632, v145 + (map[v144 := v142] + map[v144 := v142]);
						v144[420] := p0 % p0;
						v144[420] := |((v2 + v2) + v2)|;
				}
				
				var v152: multiset<int> := multiset{|fm27(v0, globalState)|, p2 % p2, fm15(|p1|, globalState) * p2, p2 - 0x25b, p2 * -108};
				globalState.f2 := v152;
				var v153: multiset<multiset<int>> := multiset{v152 - v152};
				var v154: array<int> := new int[25];
				v154[182] := p2;
				globalState.f3, v2, v153, v154[182] := -(if (v0) then p2 else fm15(p2, globalState)), v2 + v2, v153[v152 := p2], p2;
			case DC12(cf18) =>
				globalState.f3 := p0;
				var v155: map<int, seq<bool>> := map[p0 % -p0 := v86 + [v0, v0]];
				v155 := v155[p2 := v86];
				globalState.f4 := v0;
				var v156 := new C2();
		}
		
		var v157: array<bool> := new bool[9];
		v157 := v157;
		var i8 := 0;
		while (if (p0 in p1) then p1[p0] else v86[p2])
			decreases 100 - i8
		{
			if (i8 >= 100) {
				break;
			}
			
			i8 := i8 + 1;
			var v158: map<string, int> := map[v2 := |v2|];
			v158 := v158[v2 := if (v0) then |("bly")[p0 := v1]| else -p2];
			var v159: set<int> := {p2, -0x6, p0};
			var v160: map<char, set<int>> := map[v1 := v159];
			v160 := v160[v1 := v159 * {p2}];
			var v161 := new C0();
			globalState.f4 := |p1| <= p2;
		}
		var v163 := DC29(set v162 : string | v162 in fm41(v86, v1, p0, globalState) :: (v162));
		match (match v163 {
			case DC30(cf42, cf43, cf44, cf45) => DC14(|(seq(145, i9  => (v1)))|, !cf42)
			case DC29(cf41) => DC14(p2, v0)
			case DC31(cf46) => DC14(p0, v0)
		}) {
			case DC14(cf20, cf21) =>
				v0 := cf21;
				var v164: map<int, bool> := map[909 := false ==> cf21];
				v164 := v164[cf20 := false];
				var v165: multiset<int> := multiset{-720, p2, p0, cf20};
				globalState.f2 := v165;
				var v166: set<bool> := {fm16(v0, globalState)};
				var v167: seq<int> := [p2, p2, |v166|];
				var v168: multiset<bool> := multiset{if (fm3(v167, v165, v0, globalState)) then v0 else cf21, cf21, cf21, v86[p0], v0};
				var v169 := DC7(v1, false, p0);
				v168, globalState.f3, cf20, v166 := v168 - v168, v169.cf11, fm15(p2 % 824, globalState), if (v0) then v166 else v166;
			case DC15() =>
				var v170 := new C0();
				var v171: seq<int> := [p0, p0];
				var v172: map<int, seq<bool>> := map[p0 := v86];
				globalState.f4 := (|v171| / p2) != --(p0 * |(if (v170.fm19(globalState) in v172) then v172[v170.fm19(globalState)] else v86)|);
				var v173 := DC8(fm2(globalState));
				var v174: multiset<seq<bool>> := multiset{v86, v86};
				globalState.f3 := -fm26(v173, p0, v174, globalState);
				var v175 := new C2();
			case DC13(cf19) =>
				globalState.f4 := v2 == "r";
				var v176: array<D9> := new D9[3];
				var v177: map<array<D9>, int> := map[v176 := p0];
				v177 := (v177 + v177) + v177;
				var v178: array<int> := new int[17](i10 => i10 + p2);
				v178[165] := 0x383;
				globalState.f4, v178, v178[165] := v0, v178, --(if (true) then p0 else p0);
				var v179: multiset<bool> := multiset{v0, v0};
				var v180: map<char, bool> := map[v1 := v0];
				var v181: seq<map<char, bool>> := [v180];
				var v182: multiset<bool> := multiset{v0, v0, v0, ([map[v1 := v0]])[if (v0 in v179) then v179[v0] else v178[165] := v180] < v181};
				v179 := fm0(|"ldambf"| % v178[165], globalState);
		}
		
		var i11 := 0;
		while (!(387 in fm27(v0, globalState)))
			decreases 100 - i11
		{
			if (i11 >= 100) {
				break;
			}
			
			i11 := i11 + 1;
			globalState.f3 := p2;
			var v183: multiset<map<int, bool>> := multiset{p1};
			globalState.f8 := |v183| + p0;
			v157[483] := true;
			var v184 := DC21(v0, v86, v0);
			v157[483] := v184.cf28;
			globalState.f8 := p2;
		}
		var v185 := DC9();
		r0 := v185;
	}
}

class C4 extends T0 {
	const f12 : array<map<bool, int>>
	constructor (f12 : array<map<bool, int>>) {
		this.f12 := f12;
	}
	
	function fm2(globalState: GlobalState): map<int, int> {
		if ({false} != {false, true, true}) then map[|(seq(794, i0  => (0xce)))| := |map[true := 'l']|] + (map v0 : int | (0x164 <= v0) && (v0 < 2) :: (v0 + |(seq(0x279, i1  => ('p')))|) := (|[|map[false := |multiset{|[true]|, 0x2e6}|]|, |{0x2f7}|]|)) else map[-686 := -|map[0x3bf := |map[|map[|map[false := true]| := !true]| := false]|]|] + map[|multiset([true, false])| := |{|[-0xa7]|, 0x2cd}|]
	}
	function fm3(p0: seq<int>, p1: multiset<int>, p2: bool, globalState: GlobalState): bool {
		(if (true) then false else true) || (|map[0x108 := true]| < -0x197)
	}
	function fm12(p0: bool, p1: int, p2: string, p3: D1, globalState: GlobalState): int {
		|([0x339] + [--0x3a0, -|"gpglxhpdq"|, -|"bw"|, |map[!!true := true]|])| % |"trm"|
	}
	function fm13(p0: bool, p1: string, p2: bool, globalState: GlobalState): bool {
		match DC3(DC0(|multiset{!false, true}|)) {
			case DC1(cf1, cf2) => cf1 >= |{!true, true, false, !false}|
			case DC2() => |"kammnnur"| < |map[{true, !!true, !false} := false]|
			case DC0(cf0) => !(multiset{[true]} !! multiset{[DC6(cf0, true, true).cf8, false], [false, true, true, true]})
			case DC3(cf3) => false
		}
	}
	method m0(globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0: array<bool> := new bool[23](i0 => multiset{true} == multiset{true, true, false});
		var v1 := false;
		v0[660] := v1 <==> !!!v1;
		var v2 := "upyxghc";
		var v3 := -263;
		var v4: seq<string> := [fm14(v3, globalState), v2, "vjftgkl"];
		var v5: multiset<string> := multiset{v2};
		v0[660] := (multiset{v2, v2} + multiset(v4)) < v5;
		v2 := v2;
		for i1 := v3 - -637 to v3 {
			globalState.f4 := v1;
			var v6 := new C1();
			globalState.f4 := v0[660] || v1;
			v0 := v0;
		}
		var v7 := new C3();
		if (v3 < |multiset{v3, 663, -v3}|) {
			var v8 := new C1();
			globalState.f4 := v1;
			var v9: map<int, bool> := map[v3 / fm20(v1, globalState) := v1];
			v0[660] := if (v3 in v9) then v9[v3] else true;
			var v10: multiset<bool> := multiset{v0[660]};
			var v11: map<array<bool>, int> := map[v0 := |v10| / v3];
			var v12 := DC5(v0);
			var v13: map<bool, int> := map[!v0[660] := v3];
			var v14: array<int> := new int[19](i2 => i2 / v3);
			var v15: map<int, array<int>> := map[if (v0[660] in v13) then v13[v0[660]] else v3 := v14];
			v11 := v11[v12.cf5 := |v15|];
			globalState.f3 := |v2| % v3;
		} else {
			var v16: array<int> := new int[6];
			var v17: map<int, bool> := map[v3 := v0[660]];
			v16[625] := |v17|;
			v16[625] := 975 + (v3 * v3);
			globalState.f4 := v0[660] <==> false;
			globalState.f3 := v3;
			var v18: set<int> := {207};
			var v19: C0 := new C0();
			var v20 := DC30(v0[660], v0[660], v18, v19);
			var v21 := DC31(v20);
			var v23 := DC1(v3, |v2|);
			v21 := if ((set v22 : int | (756 <= v22) && (v22 < 0x2b2) :: (v22 * v3)) !! {v23.cf1}) then if (v1) then v21 else v21 else v21;
			if (true) {
				var v24: seq<bool> := [v0[660]];
				globalState.f4, globalState.f3, v24 := v0[660], v3, (v24 + v24) + v24;
				v1 := |v24| < |(v2 + fm25(globalState))|;
				v1 := v1;
				var v25: seq<int> := [v3, v16[625]];
				var v26: map<array<bool>, seq<int>> := map[v0 := v25];
				v26 := v26[v0 := (seq(0x275, i3  => (v16[625]))) + v25];
				var v27: array<char> := new char[19](i4 => 'g');
				var v28 := 'w';
				v27[872] := v28;
				v27[872] := v28;
			} else {
				v16[625] := (v16[625] / v16[625]) - v16[625];
				var v29: seq<bool> := [!v1];
				v0, globalState.f4, v2 := v0, (v29 + v29[v3 := v1]) < (v29 + v29), v2;
				var v30 := DC4(v1);
				v16[625] := fm12(|v2| <= -v16[625], v3, v2, v30, globalState);
				var v31: seq<int> := [v16[625], v3];
				var v32: map<array<int>, int> := map[v16 := |v31|];
				var v33: map<int, int> := map[v3 := |v32|];
				globalState.f8 := v16[625] / |(map[v3 := v3] + v33)|;
				var v34 := new C2();
			}
			
		}
		
		var v35 := DC4(v1);
		globalState.f4, v2, v0[660], r1, globalState.f4 := v1, if (v0[660]) then v2 else "vtakwmgb", v0[660], !v35.cf4, v0[660] || v1;
		r0 := v1;
		r1 := v3 < -0x391;
	}
	method m9(p0: int, p1: int, globalState: GlobalState) {
		var v0: array<bool> := new bool[2](i0 => false !in map[false := "fjxvcpd"]);
		v0[364] := !false;
		var v1 := true;
		v0[364] := v1;
		globalState.f8 := |fm31(v1, if (v0[364]) then v0[364] else !v1, v1, globalState)|;
		for i1 := p0 to p1 {
			var v2 := "koegpi";
			var v3 := 'e';
			var v4: seq<string> := ["smk"];
			var v5: C0 := new C0();
			var v6: seq<C0> := [v5];
			var v7: set<string> := {v2, v4[|multiset(v6)|]};
			var v8: map<bool, set<string>> := map[(v2[p0 := v3])[|{v1}| := v3] != v2 := v7];
			var v9: multiset<int> := multiset{|v2|, i1, p0};
			var v10: set<multiset<int>> := {v9};
			v8 := v8[v10 >= v10 := v7 - v7];
			var v11: map<bool, bool> := map[v0[364] := v0[364]];
			var v12: seq<map<bool, bool>> := [map[!v1 := v0[364]], v11];
			globalState.f4, v2, v0[364], v0[364], v0[364] := !(v11 == v12[i1]), v2[i1 := v3], v1, v9 !! (v9 * v9), p1 > |v2|;
			globalState.f4 := v0[364];
			v3 := v3;
		}
		var v13 := DC28(DC27(p0, p1, 0x1, -p1));
		var v14 := DC28(v13);
		match (DC28(v13)) {
			case DC27(cf36, cf37, cf38, cf39) =>
				cf39 := -cf36;
				globalState.f8 := fm20(v0[364], globalState);
				var v15 := "io";
				cf38 := -cf39 / |(v15 + v15)|;
				var v16 := new C3();
			case DC26(cf35) =>
				var v17: array<int> := new int[26](i2 => i2 / |multiset{v0[364], v0[364]}|);
				v17 := v17;
				var v18: array<D0> := new D0[20](i3 => DC0(p1));
				var v19 := "ag";
				var v20: seq<int> := [p0, |v19|];
				var v21: seq<int> := [p0, p0, p0, v20[|(seq(0x22b, i4  => ('n')))|]];
				var v22 := DC0(|v21|);
				v18[867] := v22;
				v18[867] := v22;
				var v23: multiset<bool> := multiset{v0[364], v1};
				var v24: map<int, multiset<bool>> := map[p0 := v23];
				globalState.f8 := -(if ((v23 !! (if (p1 in v24) then v24[p1] else v23)) in v23) then v23[v23 !! (if (p1 in v24) then v24[p1] else v23)] else p0);
				v17 := v17;
			case DC28(cf40) =>
				globalState.f3 := p0;
				globalState.f8 := -323;
				globalState.f8 := p0;
				var v25 := "sdikjo";
				var v26 := 'w';
				var v27: seq<string> := [v25];
				var v28: seq<int> := [p0];
				var v29 := DC24(p1, !v1, v25);
				var v30: array<string> := new string[20] [v25, v25, "i", "ceextbwbk", v25 + (seq(0x1c3, i5  => ('t'))), v25, v25, v25, v25, "lncxgmpd" + v25, v25 + v25, v25, v25, fm17(v26, globalState), v27[v28[p0]], v25, seq(-408, i6  => (v26)), v29.cf33, v25, v25];
				v30[666] := v25 + v25;
				v30[666] := v25;
		}
		
		var v31: array<map<int, bool>> := new map<int, bool>[5](i7 => map[p1 := v1]);
		var v32: multiset<int> := multiset{p1, p0};
		v31[222] := map[p0 := fm3([p0], v32, false, globalState)] + (map v33 : int | (0x3a <= v33) && (v33 < 0x13e) :: (v33 / p1) := (v1));
		var v35: seq<bool> := [v1];
		var v36: map<int, int> := map[-0x1ee := |multiset(v35)|];
		var v37: C3 := new C3();
		var v38 := DC1(p0, |v35|);
		var v39: map<C3, int> := map[v37 := v38.cf1];
		var v40 := DC37(v37);
		var v41 := 'k';
		v31[222] := (map v34 : int | v34 in v36 :: (v34 % p0) := (v1))[if (v40.cf56 in v39) then v39[v40.cf56] else --|fm1(p0, v41, v0[364], v36, globalState)| := v1];
		var v42: seq<char> := [fm18(v0[364], v0[364], globalState), 'a', v41, v41];
		var v43 := DC24(p0, v1, v42);
		var v44: map<bool, string> := map[v0[364] := v43.cf33];
		var v45 := DC7(v41, v0[364], p1);
		v44 := v44[v45.cf10 := v42];
	}
}

class C5 extends T0 {
	constructor () {
	}
	
	function fm2(globalState: GlobalState): map<int, int> {
		match DC1(|[|multiset(DC10([!false, true]).cf13)|, |(set v0 : set<bool> | v0 in map[{true} := true] :: (v0))|, |{true, true}|]|, 0x115) {
			case DC1(cf1, cf2) => if (!!true) then map[cf2 := cf1] else map[|[true, false]| := cf1]
			case DC2() => (map v1 : int | (532 <= v1) && (v1 < 298) :: (v1 + -0xa1) := (-0x38a)) + map[477 := |(seq(-0x3b4, i0  => (-0x16e)))|]
			case DC0(cf0) => if (true) then map[cf0 := cf0] else map v2 : int | v2 in {cf0, 771, cf0} :: (v2 - |map[false := 'k']|) := (cf0)
			case DC3(cf3) => map[|map[false := true]| := -365] + map[|"c"| := 0x231]
		}
	}
	function fm3(p0: seq<int>, p1: multiset<int>, p2: bool, globalState: GlobalState): bool {
		([|"qvoets"|, |"p"|] + (seq(-0x2b2, i0  => (-0x1a3)))) == [994, 0x203, |(seq(453, i1  => (0x1ee)))|, -|(seq(0x381, i2  => ('r')))|]
	}
	method m0(globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0 := new C1();
		var v1: array<D4> := new D4[22];
		var v2: map<array<D4>, bool> := map[v1 := false];
		var v3 := "dyu";
		var v4: set<string> := {v3};
		var v5 := 0x70;
		var v6 := 'o';
		var v7 := true;
		v2 := v2[v1 := v4 != (set v8 : string | v8 in fm42(v5, v5, |map[v6 := v5]|, v7, globalState) :: (v8))];
		var v9: multiset<bool> := multiset{v7};
		globalState.f3 := (-0x44 + v5) * |(v9 - v9)|;
		v7 := "lxdhgpiq" < v3;
		var v10: map<bool, bool> := map[v7 := v7];
		var v11: map<int, map<int, char>> := map[87 := map[|v10| := v6]];
		var v12: map<int, char> := map[v5 := 'q'];
		var v13: seq<map<int, char>> := [v12, v12, v12, v12];
		var v15: map<int, int> := map[-0x101 := 0x365];
		var v16: seq<bool> := [v7, !v7];
		var v17: map<bool, int> := map[v7 := --fm26(DC8(v15), |v3|, multiset{([v7, v7])[v5 := false], v16}, globalState)];
		var v18: seq<int> := [|v3|, v5, -|v17|, v5, 0x154];
		var v19 := DC38(v12);
		var v21: map<int, bool> := map[|v16| := v7];
		var v22: array<map<int, char>> := new map<int, char>[20] [if (0x8 in v11) then v11[0x8] else map[v5 := v6], v12, v12, v12 + v12, v13[v5] + v12, v12, v12 + v12, v12, v12, v12, fm43(!v7, v5, v6, globalState), v12, map v14 : int | v14 in v18 :: (v14 * v5) := (v6), v12, v12, v12, if (v7) then map[v5 := v6] else v19.cf57, v12 + v12, map[v5 := v6], (map v20 : int | v20 in v21 :: (v20 + |v12|) := ('a'))[v5 := v6] + v12];
		v22 := v22;
		var v23 := DC1(v5, v5);
		var v24: set<D0> := {v23, v23};
		var v25: multiset<int> := multiset{v5};
		var v26 := DC39(v7, v5);
		v0.m13(|v24|, if (v7) then v25[|v3| := v5] else fm27(v26.cf58, globalState), v7, globalState);
		r0 := !true;
		r1 := !(if (v5 <= v5) then v7 else v0.fm3(v18, v25, v7, globalState));
	}
}

class C6 extends T0 {
	const f10 : int
	const f11 : string
	constructor (f10 : int, f11 : string) {
		this.f10 := f10;
		this.f11 := f11;
	}
	
	function fm2(globalState: GlobalState): map<int, int> {
		map[DC1(f10, f10).cf2 / 0x223 := f10]
	}
	function fm3(p0: seq<int>, p1: multiset<int>, p2: bool, globalState: GlobalState): bool {
		[!true] <= [true, false]
	}
	method m0(globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0: array<int> := new int[9];
		v0[359] := f10;
		var v1 := false;
		var v2: map<bool, int> := map[v1 := f10];
		v0[359] := if (v1) then |f11| else if (v1 in v2) then v2[v1] else 0x2f5;
		for i0 := v0[359] to 0x95 {
			v0 := if (v1) then v0 else v0;
			var v3: T0 := new C3();
			var v4: set<T0> := {v3};
			v4 := v4 * v4;
			var v5 := 'd';
			var v6: map<int, int> := map[f10 := i0];
			var v7: multiset<int> := multiset{f10};
			var v8: seq<bool> := [v1, v1, v1, v1, v1];
			var v9: seq<int> := [|f11|, -i0];
			var v10: map<bool, bool> := map[!v1 := v1];
			var v11: array<bool> := new bool[23] [v1, v1 <== v1, true, v1, if (v1) then v1 else v3.fm3(fm1(|map[0xf2 := v0[359]]|, v5, v1, v6, globalState), v7, v1, globalState), !v1, v1, v8[f10], v1, f10 <= f10, v1, v1, v3.fm3(v9, v7, v1, globalState), v1, v1 <==> true, v1, v1, v10 != v10, if (v1) then v1 else v1, v1, v1, v1, v1];
			v11[304] := v1;
			v11[304] := !((f10 !in v7) && v1);
			if ((0x203 - v0[359]) == -0x26) {
				var v12: map<bool, char> := map[v1 := v5];
				v12 := v12[v1 := v5];
				globalState.f4 := v11[304];
				var v13: multiset<bool> := multiset{v1, false};
				var v14: set<int> := {|v13|, v0[359] / f10};
				var v15: seq<array<int>> := [v0];
				v14, v11[304] := v14, !(v15 == v15);
				v11[304] := f10 < i0;
				var v16: multiset<seq<bool>> := multiset{v8};
				r1 := v16[v8 := -f10] < (v16 * multiset{v8, v8});
			} else {
				var v17: map<char, map<bool, bool>> := map[v5 := v10];
				var v18: map<map<bool, bool>, array<int>> := map[if (v5 in v17) then v17[v5] else v10 := v0];
				globalState.f4 := v10 in v18;
				r0 := !v1;
				var v19: map<bool, seq<int>> := map[v11[304] := [|v9|, -f10]];
				globalState.f4, v19 := v1 && v1, v19 + v19;
				var v20: set<int> := {v0[359], -|f11|, v0[359], i0 - v0[359]};
				var v21: multiset<bool> := multiset{v1};
				var v22: array<string> := new seq<char>[21](i1 => DC24(v0[359], v1, f11).cf33 + (seq(834, i2  => (v5))));
				v22[193] := f11;
				var v23 := DC39(true, f10);
				var v24 := DC40(v23);
				var v25: set<D19> := {v24, v24, DC40(v23), v24};
				var v26: map<set<D19>, multiset<bool>> := map[v25 := multiset{v11[304], v1, v11[304], false} + multiset(fm31(v1, v11[304], v1, globalState))];
				var v27: map<bool, seq<bool>> := map[v11[304] := v8];
				v20, v21, globalState.f3, globalState.f8, v22[193] := v20, if (v25 in v26) then v26[v25] else v21 + v21, |(if (v11[304]) then if (v11[304] in v27) then v27[v11[304]] else v8 else v8)|, |v7| % i0, (seq(-0x139, i3  => (v5))) + f11;
				var v28: set<bool> := {v11[304], true};
				v28 := (v28 * v28) * {true};
			}
			
		}
		var v29: array<bool> := new bool[17] [false, v1, v1, v1, v1, v1, true, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1];
		var v30: set<array<bool>> := {v29};
		var v31 := DC5(v29);
		var v32: set<bool> := {v1, v1};
		var v34: set<int> := {f10, f10};
		var v35: seq<bool> := [v1];
		var v36: seq<seq<int>> := [[|v35|]];
		var v37 := 'y';
		var v38: map<string, bool> := map[("wvptjdmb")[|v36[v0[359]]| := v37] := false];
		var v39: array<bool> := new bool[9] [v0[359] <= -f10, v30 !! {v29, v31.cf5}, v32 !! {v1}, v1, v1, v1, v1, !((set v33 : int | (0x3c4 <= v33) && (v33 < -0x3a5) :: (v33 + -|[v1]|)) > v34), if (v1) then !false else if (f11 in v38) then v38[f11] else v1];
		v29[443] := true;
		v29[443] := false;
		forall i4 | 0 <= i4 < v0.Length {
			v0[i4] := i4 * f10;
		}
		for i5 := f10 to v0[359] {
			r0 := v1;
			var v40: seq<array<int>> := [v0];
			v0 := if (v1) then v40[f10] else v0;
			var v41 := new C0();
			var v42 := DC18();
			match (v42) {
				case DC18() =>
					globalState.f3 := fm20(v29[443], globalState);
					var v43: seq<int> := [f10, i5];
					var v44: map<bool, bool> := map[v1 := v29[443]];
					v29[443] := (-i5 % |v43[|v35| := |v44|]|) <= v0[359];
					var v45 := DC20(v32);
					v45 := v45;
					var v46 := "mn";
					v46 := v46;
				case DC17(cf23) =>
					var v47: multiset<int> := multiset{v0[359], 0x216};
					var v48: multiset<bool> := multiset{fm3(seq(0x10b, i6  => (634)), v47, v29[443], globalState), v1, f10 <= 0x316, v29[443], v1};
					v0[359] := |v48|;
					v0[359] := -(if (v29[443]) then 817 else f10 / i5);
					globalState.f8 := i5;
					globalState.f3 := -v0[359];
			}
			
		}
		v37 := v37;
		r0 := f11 != f11;
		r1 := v1;
	}
	method m7(p0: bool, p1: seq<D2>, p2: set<char>, globalState: GlobalState) {
		var i0 := 0;
		while (p0)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v0: seq<bool> := [p0, if (p0) then p0 else p0, p0, !p0];
			var v1: array<string> := new string[5](i1 => f11[f10 := 'l']);
			v1[464] := fm14(f10, globalState) + f11;
			var v2: C2 := new C2();
			var v3 := DC34(v2);
			var v4 := 'x';
			var v5: multiset<int> := multiset{f10};
			var v6: map<bool, char> := map[v5 > fm27(p0, globalState) := v4];
			v0, v1[464], v3, v4, globalState.f4 := v0 + v0, f11 + f11, v3, if ((if (p0) then p0 else p0) in v6) then v6[if (p0) then p0 else p0] else v4, p0;
			globalState.f4 := p0;
			var v7: seq<int> := [f10];
			if (v2.fm3(v7, v5, p0, globalState)) {
				var v8, v9, v10, v11 := v2.m12(globalState);
				globalState.f4 := v11;
				var v12: set<int> := {f10};
				v12 := v12 * v12;
				globalState.f8 := -v9 * (v7[v10] + -|v5|);
				v1[464] := f11;
			} else {
				v4 := v4;
				globalState.f4 := p0;
				var v13: multiset<bool> := multiset{!p0, p0};
				var v14 := DC22(v0[|v13| := p0]);
				var v15: set<bool> := {p0};
				var v16: map<D9, int> := map[v14 := f10 * |v15|];
				var v18: map<int, int> := map[114 := |(map v17 : int | (-921 <= v17) && (v17 < 719) :: (v17 % |v15|) := (v4))|];
				var v19: multiset<seq<bool>> := multiset{v0};
				v16 := v16[v14 := fm26(DC8(v18), f10, v19, globalState)];
				var v20: set<int> := {f10, f10, f10};
				globalState.f8 := f10 - |v20|;
				globalState.f4 := f10 <= (if (f10 in v18) then v18[f10] else f10);
			}
			
			var v21 := new C3();
		}
		var v22: array<string> := new string[14](i3 => f11[f10 := 'd'] + "kjr");
		forall i2 | 0 <= i2 < v22.Length {
			v22[i2] := (seq(-239, i4  => ('t'))) + ("hxtmbhm" + f11);
		}
		var v23: array<bool> := new bool[12](i5 => {f10} >= {|map[|multiset{p0, p0}| := seq(0x21e, i6  => (f10))]|});
		v23[981] := p0;
		var v24: seq<int> := [f10, f10];
		var v25: seq<int> := [|v24|];
		globalState.f8, v23[981], globalState.f4, globalState.f8 := f10, if (p0) then p0 else fm3(v25, multiset{f10}, p0, globalState), p0 && p0, f10 % f10;
		var v26: array<int> := new int[21];
		var v27: array<array<int>> := new array<int>[18] [v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v26];
		var v28: map<array<array<int>>, int> := map[v27 := 179];
		for i7 := if (v27 in v28) then v28[v27] else fm20(v23[981], globalState) to -(f10 + f10) {
			var v29: multiset<bool> := multiset{p0};
			var v30: map<bool, bool> := map[v23[981] := v29 >= v29];
			var v31: multiset<int> := multiset{f10};
			globalState.f3, globalState.f8, v30, globalState.f4, v23[981] := (i7 + f10) - (0x72 / |v31|), f10, (v30 + v30) + v30, v23[981], v23[981];
			var v32 := new C5();
			var v33: map<map<bool, bool>, int> := map[v30 := i7];
			var v34: map<int, bool> := map[|['v']| := v23[981]];
			var v35: map<multiset<int>, bool> := map[v31 := !(map[f10 := p0] == v34)];
			var v36: C0 := new C0();
			var v37: set<C0> := {v36, v36};
			v33, v35, v26, globalState.f0 := v33 + (v33 + map[v30 := |map[f10 := |v37|]|]), map[multiset{i7} * v31 := v23[981]], v26, v24 + v25;
			var v38: seq<array<int>> := [v26];
			var v39: map<seq<array<int>>, bool> := map[v38 + v38 := v23[981]];
			var v40 := 'l';
			globalState.f4 := if (v38 in v39) then v39[v38] else f10 > |f11[i7 := v40]|;
		}
		if (true) {
			var v41: seq<bool> := [v23[981]];
			var v42 := DC10(v41[|[-f10, f10]| := p0]);
			var v43: map<int, D4> := map[f10 * f10 := v42];
			v43 := v43[f10 := v42];
			var v44 := "mjqmhkj";
			v44 := v44;
			var v45: multiset<int> := multiset{f10, -|v25|, f10};
			v23[981] := f10 == |v45|;
			globalState.f4 := p0;
			globalState.f3 := f10;
		} else {
			var v46 := 'i';
			v46 := if (v23[981]) then 'q' else v46;
			v22[682] := f11;
			v22[682] := f11;
			var v47: set<int> := {f10};
			var v48 := DC27(f10, |v47|, f10, f10);
			var v49: map<int, int> := map[|multiset{f10}| := v48.cf39];
			var v50 := DC8(v49);
			globalState.f8 := fm26(v50, f10, fm44(f10, globalState), globalState);
			globalState.f4 := if (p0) then v23[981] else v23[981];
			var v51: seq<bool> := [v23[981], p0];
			var v52: seq<seq<bool>> := [v51, v51, v51];
			var v53: set<seq<bool>> := {v52[|{false, p0, p0, v23[981], v23[981]}|]};
			globalState.f3 := |v53| * |f11|;
		}
		
		var v54 := 'k';
		var v55 := DC7(v54, true, f10);
		var v56: set<int> := {f10, v55.cf11};
		globalState.f8 := |v56|;
	}
	method m8(p0: int, globalState: GlobalState) returns (r0: multiset<int>) {
		var v0 := true;
		var v1: T0 := new C1();
		var v2: seq<T0> := [v1];
		var v3: map<bool, seq<T0>> := map[v0 := v2];
		globalState.f8 := |f11| % |(if (v0 in v3) then v3[v0] else DC41(v2).cf61)|;
		var v4: seq<int> := [p0];
		var v5: array<char> := new char[14];
		var v6: seq<array<char>> := [v5, v5];
		var v7 := DC8(map[|v6| := |f11|]);
		var v8: multiset<seq<bool>> := multiset{fm31(v0, v0, v0, globalState)};
		globalState.f3 := -v4[fm26(v7, f10, v8, globalState)];
		var v9 := new C3();
		var i0 := 0;
		while (f10 == -(0x185 * f10))
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v10: array<map<bool, int>> := new map<bool, int>[29](i1 => map[!false := p0]);
			var v11 := new C4(v10);
			var v12: map<bool, int> := map[v0 := f10];
			globalState.f3 := p0 / (if (v0 in v12) then v12[v0] else f10);
			var v13 := 'f';
			var v14: map<bool, char> := map[v11.fm13(v0, f11, v0, globalState) := v13];
			v14 := v14[v0 := v13];
			var v15: multiset<int> := multiset{f10};
			var v16: seq<bool> := [!v11.fm3(v4, v15, v0, globalState)];
			var v17 := DC4(v0);
			var v18 := DC1(v9.fm15(|v16|, globalState), -(v11.fm12(v0, p0, f11, v17, globalState) - p0));
			var v19: map<int, int> := map[f10 := p0];
			var v20: set<int> := {p0};
			var v21: multiset<set<int>> := multiset{{if (p0 in v15) then v15[p0] else p0, f10, |v19|}, v20, v20};
			var v22: map<D2, int> := map[DC7(v13, v0, f10) := |v21|];
			var v23: set<int> := {f10 - f10, |v22|};
			v0, globalState.f4, v18, v20 := v0, !(v13 !in "wcumpsjw"), v18.(cf1 := -(if (f10 in v15) then v15[f10] else p0)), v23;
		}
		var v24: set<bool> := {v0, v0};
		var v25: seq<bool> := [!(v24 < v24)];
		var v26: map<bool, bool> := map[v0 := v0];
		var i2 := 0;
		while (v25[|v26|])
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			var v27: array<bool> := new bool[18](i3 => v0);
			v27[450] := v0;
			var v28: multiset<int> := multiset{f10};
			v27[450] := (v28 * v28) !! v28;
			globalState.f8 := 463 % |f11|;
			var v29 := DC19([-p0]);
			match (v29) {
				case DC19(cf24) =>
					var v30: array<D2> := new D2[9];
					var v31: multiset<array<D2>> := multiset{v30};
					var v32 := DC33(v31);
					v32 := v32;
					r0 := v28;
					var v33 := "uoxvu";
					var v34 := 'e';
					var v35 := DC24(p0, v27[450], [v34]);
					v33 := v35.cf33;
					globalState.f8 := -(f10 * (|f11| % |v33|));
			}
			
			var v36: array<array<int>> := new array<int>[18];
			var v37: array<int> := new int[25];
			v36[506] := v37;
			v36[506] := v37;
		}
		var v38 := DC41(v2);
		match (v38) {
			case DC41(cf61) =>
				if (v0) {
					var v39 := 'n';
					v39 := v39;
					globalState.f3 := p0;
					v4 := v4;
					globalState.f8 := -f10 + f10;
					var v40 := "nujca";
					globalState.f3, v40 := -550 % |f11[-|v40| := v39]|, f11[fm26(v7, p0, fm44(p0, globalState), globalState) := v39];
				} else {
					var v41 := 'a';
					var v42: map<char, bool> := map[v41 := v0];
					globalState.f0 := (seq(958, i4  => (|v24|))) + [f10, |v42|, |v25|];
					var v43 := DC13(v26 + v26);
					v43 := v43;
					var v44: map<int, bool> := map[|f11| := v0];
					var v45: set<int> := {p0};
					var v46: multiset<char> := multiset{'j'};
					v44 := v44[|(v45 * v45)| := v46 == multiset{v41}];
					var v47: map<int, string> := map[p0 := f11];
					var v48: map<string, int> := map[if (f10 in v47) then v47[f10] else ("t")[|{!true, true}| := v41] := -(f10 + f10)];
					v48 := v48;
					var v49: multiset<int> := multiset{f10, f10, p0};
					var v50: array<bool> := new bool[2] [v0, v1.fm3(v4[p0 := f10], v49, v0, globalState)];
					v50[879] := v9.fm16(v0, globalState);
					v50[879] := (f11 < f11) !in v25;
				}
				
				v0 := !v0;
				var v51: array<bool> := new bool[17](i5 => v0);
				v51[708] := v0;
				v51[708] := v0;
				var v52 := new C1();
		}
		
		var v53: array<int> := new int[20];
		var v54: seq<array<int>> := [v53, v53];
		var v55: multiset<int> := multiset{v9.fm15(p0, globalState) % |v54|};
		r0 := v55;
	}
}

class C7 extends T0 {
	constructor () {
	}
	
	function fm2(globalState: GlobalState): map<int, int> {
		(map[|(seq(0x312, i0  => ('m')))| := 0x211] + map[|"ctfhl"| := |"w"|]) + DC8(map[-0x3bb := -|"awvtjnrl"|]).cf12
	}
	function fm3(p0: seq<int>, p1: multiset<int>, p2: bool, globalState: GlobalState): bool {
		(multiset{|map[|"no"| := false]|, DC0(|(set v0 : int | (385 <= v0) && (v0 < 0x2b5) :: (v0 * -0x2fa))|).cf0, |multiset{|(seq(501, i0  => (-|(seq(0xad, i1  => ('i')))|)))|}|, 0x12c, 0x20c} * multiset([|map[false := 0x31b]|])) !! multiset{|map[470 := true]|, -|map[!false := false]|, -0x34, |[|map[0x32d := DC8(map v1 : int | v1 in map[|map[0x286 := |"fsmppmgyx"|]| := 0x5f] :: (v1 - -781) := (0x31d))]|]|}
	}
	function fm10(globalState: GlobalState): int {
		|(((seq(15, i0  => (-0x71))) + [|[!!true]|]) + [993])|
	}
	method m0(globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0 := 0x389;
		var v1 := DC6(v0, false, true);
		var v2 := false;
		var v3: multiset<D2> := multiset{v1, v1, v1, v1, DC6(v0, v2, v2)};
		v3 := v3;
		v2 := v2;
		globalState.f8 := v0;
		var v4 := DC4(v2);
		if (match v4 {
			case DC4(cf4) => cf4
		}) {
			if (true && v2) {
				globalState.f3 := -v0;
				var v5: map<bool, int> := map[!v2 := v0];
				var v6: seq<int> := [v0, v0, v0];
				var v7: multiset<int> := multiset{v0};
				var v8: map<map<bool, int>, bool> := map[v5 := fm3(v6, v7, v2, globalState)];
				var v9: map<map<map<bool, int>, bool>, int> := map[v8 := v0];
				v0 := if (if (v2) then !!v2 else !v2) then |fm11(v2, v2, globalState)| else -(if (v8 in v9) then v9[v8] else -971);
				var v10: array<map<bool, int>> := new map<bool, int>[5] [v5, v5, map[v2 := 0x206], v5, map[v2 := v0]];
				var v11 := new C4(v10);
				var v12: T0 := new C4(v10);
				var v13 := "nph";
				var v14: map<bool, bool> := map[v2 := v11.fm3(seq(0x206, i0  => (|[v13, v13, "fmigiwsi", seq(211, i1  => ('t')), v13]|)), v7, v2, globalState)];
				v12, globalState.f4, globalState.f8, globalState.f8 := v12, !v2, -((v0 + |v13|) - |(v14 + fm21(v2, globalState))|), v0;
				var v16: array<set<set<int>>> := new set<set<int>>[7](i2 => {set v15 : int | v15 in map[v0 := -v0] :: (v15 % 0x264), {v0}, {v0, |[v2, v2]|}});
				var v17: set<int> := {|{true}|, -v0, |v13|};
				var v18: set<set<int>> := {v17};
				v16[883] := v18;
				var v19: array<bool> := new bool[7];
				v19[858] := v12.fm3(v6, multiset{v0, v0}, v2, globalState);
				r0, v12, v16[883], v19[858] := if (v2 in v14) then v14[v2] else !v2, v12, v18, v2;
			} else {
				var v20: set<int> := {v0};
				var v21: multiset<int> := multiset{|multiset{v2, v2}|};
				var v22: seq<bool> := [v20 >= v20, !!(v21 >= multiset([v0, v0])), !v2];
				v22 := v22;
				var v23: C1 := new C1();
				var v24: map<bool, bool> := map[v2 := v2];
				var v25: map<C1, bool> := map[v23 := if (v2 in v24) then v24[v2] else v2];
				var v26: seq<C1> := [v23, v23];
				v25 := v25[v26[|v20|] := !v2];
				var v27: map<bool, int> := map[v2 := v0];
				var v28 := DC26(v27);
				var v29: seq<map<bool, bool>> := [v24, v24, v24];
				var v30: map<D12, int> := map[v28 := fm20(v2, globalState)];
				var v31: set<map<D12, int>> := {(map[v28 := |v29[v0]|])[v28 := v0], if (v2) then v30 else v30};
				var v32 := "xhgx";
				var v33: seq<int> := [v0, 841];
				var v34 := DC19(v33);
				var v35: map<seq<int>, int> := map[seq(0x257, i3  => (0x3ac)) := v0];
				var v36: map<int, bool> := map[v0 := v2];
				v31 := fm45(v32, --|v34.cf24| - |v35|, |v36|, globalState);
				var v37: array<bool> := new bool[4] [false, v2, v2, v2];
				var v38 := DC5(v37);
				var v39: map<array<bool>, bool> := map[v38.cf5 := v2];
				var v40: array<bool> := new bool[10] [v2, !v2, v2, v2, v2, if (v37 in v39) then v39[v37] else v2, v2, v2, v2, false];
				var v41: map<array<bool>, int> := map[v40 := v0];
				v41 := v41[v37 := -116];
				var v42 := new C5();
			}
			
			var v43, v44, v45 := m5(v2, v0, |fm46(globalState)|, v0, globalState);
			v45 := v2;
			var v46: seq<bool> := [v45, v2, v45];
			var v47 := DC10(v46);
			var v48: seq<D4> := [v47];
			r1 := |v48| <= v0;
			var v49 := DC10(v46);
			var v50 := DC12(v49);
			var v51 := DC12(v49);
			match (if (v2) then DC12(DC12(v49)) else v51) {
				case DC11(cf14, cf15, cf16, cf17) =>
					var v52 := new C5();
					globalState.f3 := v0 + v0;
					globalState.f4 := v2;
					var v53: map<bool, int> := map[v2 := cf16 + cf16];
					var v54: multiset<int> := multiset{v0};
					v53 := v53[v52.fm3(seq(0x103, i4  => (|map[v45 := cf17]|)), v54, v2, globalState) := v0];
				case DC10(cf13) =>
					var v55: array<seq<array<int>>> := new seq<array<int>>[29];
					var v56: array<int> := new int[15];
					var v57: seq<array<int>> := [v56, v56, v56];
					v55[700] := v57;
					v55[700] := v57[-642 := v56];
					var v58: array<array<int>> := new array<int>[9];
					v58[275] := v56;
					v58[275] := v56;
					var v59 := "tkabsjr";
					var v60: map<bool, bool> := map[v0 < |v59| := v45];
					v60 := v60[v2 := v2];
					var v61: array<bool> := new bool[24] [v2, v45, v2, v2, v45, false, true, v45, v45, v45, false, true, v45, !v45, v45, v2, v2, v2, !true, !v45, v45, v45, v2, false];
					m6(fm10(globalState), v61, v0 * 454, globalState);
				case DC12(cf18) =>
					var v62 := "fwbkil";
					var v63 := DC11(v45, v45, v0, v62);
					var v64: multiset<int> := multiset{v0, v0};
					var v65: set<bool> := {v2, false};
					var v66: map<multiset<int>, set<bool>> := map[v64 := v65 - v65];
					var v67 := DC21(v45, v46, v45);
					v63, v66, v62, globalState.f8, globalState.f8 := fm30(fm10(globalState), v67, v2, v0, globalState), v66, fm25(globalState), (v0 - -v0) + |v63.cf17[v0 := v43]|, v0;
					var v68 := new C2();
					globalState.f4, v45, globalState.f8, v65, globalState.f4 := v46[v0], v2, |v62| / v0, v65, if (false) then v45 else v45;
					globalState.f4 := v62 in {v62, v62};
			}
			
		} else {
			var v69: map<bool, bool> := map[v2 := v2];
			var v70: seq<map<bool, bool>> := [v69, v69];
			var v71: array<seq<map<bool, bool>>> := new seq<map<bool, bool>>[3] [[v69, v69, v69], v70, v70];
			v71[932] := v70;
			v71[932] := (v70 + v70) + v70;
			globalState.f8 := v0;
			var v72 := new C1();
			var v73: seq<bool> := [v2, v2, false];
			var v74: seq<int> := [v0];
			var v75 := "i";
			var v76: multiset<int> := multiset{|v75|};
			var v77: set<bool> := {!v73[v0]};
			var v78 := 'v';
			var v79: multiset<char> := multiset{v78};
			var v80: array<bool> := new bool[23] [v73[v0], v2, !v2, v2, v2, v2, v0 >= fm10(globalState), v2, v72.fm3(v74, v76, v2, globalState), v2, v2, v77 >= v77, true, v2, v79 != v79, !v4.cf4, v2, v2, !v2, v2, v2, v2, v2];
			v80[957] := v2 <==> v72.fm3(seq(-0x219, i5  => (v0)), v76, v2, globalState);
			v80[957] := (v75 + v75[v0 := 'i']) == (fm17(v78, globalState) + v75);
			v2 := multiset(v74 + v74) < multiset(v74[-0x35c := v0]);
		}
		
		if (if (v2) then v2 else true) {
			var v81 := "cfusokk";
			var v82 := 'y';
			var v83: map<char, bool> := map[v82 := v2];
			var v84: seq<int> := [|v83|];
			var v85: map<bool, bool> := map[v2 := v2];
			var v86: multiset<int> := multiset{|v85|};
			var v87: array<bool> := new bool[18] [v2, true, v2, v2, v2, v2, v2, v2, v2, v2, false, v2, v2, false, fm3(v84, v86, v2, globalState), v2, v2, v2];
			var v88: map<int, array<bool>> := map[v0 := v87];
			var v89: map<string, map<int, array<bool>>> := map[v81 + (seq(565, i6  => ('b'))) := v88];
			v89 := v89[v81 := v88];
			if (v2) {
				v81 := v81 + v81;
				globalState.f3 := v0;
				v81 := v81;
				var v90: array<D10> := new D10[23](i7 => DC23(multiset{seq(0x5d, i8  => (v82)), seq(-0x329, i9  => (v82))}));
				var v91: multiset<string> := multiset{v81};
				var v92 := DC23(v91);
				v90[5] := v92;
				v90[5] := v92;
				globalState.f4 := v0 == (v0 % |fm25(globalState)|);
			} else {
				globalState.f3 := v0;
				globalState.f8 := v0;
				var v93: set<int> := {v0, v0, |v81| + v0, v0};
				globalState.f8 := -|v93|;
				var v94: seq<bool> := [v2];
				v87[903] := !(v94[v0] in v94);
				var v95: set<bool> := {v2, v2};
				v87[903], v95 := if (v2) then v2 else v2, v95;
				v85 := v85;
			}
			
			v2, globalState.f8 := true, -(v0 % (v0 - -v0));
			var v96: array<int> := new int[19];
			v96[825] := v0;
			v96[825] := v0;
			v96[825], v96 := v96[825], v96;
		} else {
			var v97: array<map<set<int>, string>> := new map<set<int>, string>[7];
			var v98: set<int> := {v0};
			var v99 := "xxcadue";
			var v100: map<set<int>, string> := map[v98 := v99];
			v97[594] := v100;
			var v101: array<D7> := new D7[17](i10 => DC18());
			var v102 := DC18();
			v101[331] := if (v2) then v102 else v102;
			v97[594], r0, v101[331] := (v100 + fm47(|v99|, v2, v2, globalState)) + map[v98 := seq(0x35, i11  => ('y'))], v2 <== v2, v102;
			var v103 := new C1();
			r1 := true;
			var v104: map<bool, int> := map[v2 ==> v2 := v0];
			v104 := v104;
			var v105: seq<map<bool, int>> := [v104];
			var v106: seq<int> := [v0, v0, v0, v0, v0];
			var v107: seq<map<bool, int>> := [v105[v0], v104, v104, map[v2 := |v106|] + v104];
			var v108: map<int, bool> := map[v0 := v2];
			v107 := v105[|v108[0x1fb := v2]| := fm46(globalState)];
		}
		
		var v109: array<char> := new char[5](i13 => 'h');
		forall i12 | 0 <= i12 < v109.Length {
			v109[i12] := 'h';
		}
		r0 := !v2;
		var v110: map<D2, int> := map[v1 := v0];
		var v111: set<map<D2, int>> := {v110};
		r1 := v111 > v111;
	}
	method m5(p0: bool, p1: int, p2: int, p3: int, globalState: GlobalState) returns (r0: char, r1: set<int>, r2: bool) {
		globalState.f8 := |(seq(217, i0  => ('n')))|;
		var v0: set<int> := {p2, p2};
		var v1: array<bool> := new bool[5](i2 => true);
		var v2: map<array<bool>, int> := map[v1 := -p2];
		var i1 := 0;
		while ((v0 - v0) >= {|v2|})
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			var v3: map<int, int> := map[0x2ed := p3];
			var v4 := DC8(v3);
			var v5: seq<bool> := [p0, p0];
			var v6: multiset<seq<bool>> := multiset{v5};
			globalState.f8 := fm26(v4, p1, v6, globalState) - p1;
			match (DC4(p0)) {
				case DC4(cf4) =>
					var v7 := "cwtkmlf";
					var v8: multiset<string> := multiset{v7};
					globalState.f0 := [p1 / |v8|, if (p1 in v3) then v3[p1] else p2, fm20(p0, globalState), p3];
					globalState.f4 := p0;
					v1 := new bool[14];
					v7 := fm14(p3, globalState);
			}
			
			v1[29] := p0;
			v1[29] := p0;
			var v9 := 'e';
			var v10 := "jrrugumfh";
			var v11: map<char, string> := map[v9 := v10];
			var v12: seq<int> := [p2, p2, |(if (v9 in v11) then v11[v9] else v10)|, p2, |v10|];
			globalState.f0 := v12 + v12;
		}
		var v13 := "uba";
		var v14: set<string> := {v13, v13};
		var v15 := DC29(v14);
		var v16: map<bool, D13> := map[p0 := v15];
		var v17 := DC31(if (p0 in v16) then v16[p0] else v15);
		match (v17) {
			case DC30(cf42, cf43, cf44, cf45) =>
				globalState.f3 := (p3 * p3) + p1;
				v13 := v13;
				if (p0) {
					globalState.f4 := p0 ==> p0;
					var v18: multiset<int> := multiset{0x27e};
					var v19: map<int, bool> := map[p1 * p3 := v18 >= v18];
					var v20: seq<int> := [p1];
					var v21: map<bool, multiset<int>> := map[cf42 := multiset{|v20|}];
					v19 := v19[p3 / p3 := v18 > (if (cf43 in v21) then v21[cf43] else multiset(v20))];
					var v22: map<char, bool> := map['y' := cf42];
					v22 := v22['n' := cf42];
					var v23: T0 := new C2();
					var v24: map<bool, char> := map[!fm3(v20, multiset(v20), cf42, globalState) := 's'];
					var v25: map<bool, string> := map[cf42 := v13];
					var v26 := DC39(cf42, p1);
					var v27: array<int> := new int[3];
					var v28: map<int, array<int>> := map[p1 := v27];
					var v29: map<bool, int> := map[false := p3];
					var v30: seq<array<int>> := [if ((if (cf43 in v29) then v29[cf43] else p3) in v28) then v28[if (cf43 in v29) then v29[cf43] else p3] else v27];
					var v31: map<bool, array<int>> := map[v26.cf58 := v30[|map[p3 := cf42]|]];
					var v32: map<map<bool, array<int>>, T0> := map[v31 := if (p0) then v23 else v23];
					var v33 := 'p';
					var v34: map<int, char> := map[p3 := v33];
					r1, v23, globalState.f4, v24, r2 := if (p2 != |(if (!p0 in v25) then v25[!p0] else v13)|) then cf44 else cf44, if (v31 in v32) then v32[v31] else if (cf43) then v23 else v23, fm43(cf43, p3, v33, globalState) == v34, fm11(cf42, cf43, globalState) + v24, p0;
					v27[783] := 462;
					v27[783] := if (false) then 0x2a8 - p3 else p2;
				} else {
					var v35: array<int> := new int[15];
					var v36 := DC36(v35);
					v36 := v36;
					var v37: C2 := new C2();
					var v38: map<C2, string> := map[v37 := v13];
					var v39: multiset<bool> := multiset{false};
					var v40: map<bool, bool> := map[p1 > p1 := v39 > v39];
					v35[205] := p2 * -0x1b2;
					v38, globalState.f4, v40, v35[205] := v38, p0, (v40 + v40)[!cf43 := p1 > p2], p2;
					cf42 := cf43;
					var v41: seq<int> := [0xcb, p2, p3];
					var v42: multiset<int> := multiset{p2};
					v35[205] := v41[if (v35[205] in v42) then v42[v35[205]] else v35[205]] * (-0x164 + p2);
					var v43 := DC16(v42);
					var v44: multiset<D6> := multiset{v43, v43, v43};
					var v45: map<multiset<bool>, int> := map[v39 := |v44|];
					var v46 := DC27(p1, p2, 826, |v45|);
					var v47: map<bool, D12> := map[!p0 := v46];
					var v48 := 'r';
					var v49: map<int, int> := map[|fm31(cf42, p0, cf43, globalState)| := p3];
					var v50 := DC8(v49);
					v47 := v47[p2 in fm1(v35[205], v48, cf43, v50.cf12, globalState) := v46];
				}
				
				globalState.f8 := p2;
			case DC29(cf41) =>
				globalState.f8 := fm10(globalState);
				var v51: array<array<map<int, bool>>> := new array<map<int, bool>>[7];
				var v52: array<map<int, bool>> := new map<int, bool>[9];
				v51[990] := v52;
				v51[990] := v52;
				var v53: array<int> := new int[23](i3 => i3 + |map[true := p0]|);
				var v54 := DC42(map[v53 := 731]);
				m6(724 / |v13|, v1, |v54.cf62|, globalState);
				globalState.f8 := p1;
			case DC31(cf46) =>
				v13 := v13;
				var v55 := new C1();
				var v56 := DC11(p0, true, -0x11c, v13);
				var v57 := DC12(v56);
				var v58 := DC12(v56);
				var v59 := DC12(v56);
				match (if (if (p0) then false else p0) then v59 else v59) {
					case DC11(cf14, cf15, cf16, cf17) =>
						v1[268] := false;
						v1[268] := true;
						var v60: array<map<bool, int>> := new map<bool, int>[5](i4 => map[cf14 := |map[p3 := |multiset{|[p2]|}|]|]);
						var v61 := new C4(v60);
						v60 := if (v1[268]) then v60 else v61.f12;
						var v62 := new C3();
					case DC10(cf13) =>
						var v63 := 'k';
						globalState.f4 := |v13[|v13| := v63]| != p1;
						var v64: C3 := new C3();
						var v65: multiset<C3> := multiset{v64};
						var v66: multiset<bool> := multiset{p0};
						var v67 := DC25(v66);
						var v68: map<int, D11> := map[p2 := v67];
						v13, v65, globalState.f8 := seq(78, i5  => ('s')), v65, -p3 * |v68|;
						var v69: map<int, int> := map[p3 := p2];
						var v70: multiset<seq<bool>> := multiset{[p0, p0]};
						var v71: multiset<int> := multiset{fm26(DC8(v69), p2, v70, globalState), p1, p1};
						v55.m13(0x300, v71, p0, globalState);
						globalState.f4, globalState.f8, r0, globalState.f3 := !p0, p1 / p3, v63, p2 / p3;
					case DC12(cf18) =>
						var v72: T0 := new C1();
						var v73 := DC6(|v0|, p0, p0);
						var v74: map<bool, T0> := map[v73.cf7 := v72];
						v72 := if (p0) then v72 else if (p0 in v74) then v74[p0] else v72;
						var v75 := new C3();
						globalState.f4 := if (p0) then p0 else p0;
						var v76: map<C1, int> := map[v55 := p1];
						var v77: map<bool, bool> := map[p0 := true];
						var v78: multiset<int> := multiset{p3, -p2, p2, 934, |v77|};
						var v79: array<int> := new int[8] [|v76| + p2, p1, p2, p1 - p3, 206, 0x341 - |v13|, |v78|, p3];
						v79[557] := p1;
						var v80: seq<bool> := [p0];
						var v81: map<int, int> := map[p2 := |v80|];
						var v82: seq<int> := [v75.fm15(p2, globalState)];
						var v83: seq<int> := [v82[0x30e], p3, -0x397, -p1, p1];
						v79[557] := |v81[|v82| := -0x282]| - (p1 + 0xa0);
				}
				
				var v84: map<bool, bool> := map[p0 := !p0];
				v84 := v84[true := p0];
		}
		
		var v85: seq<int> := [p1, -829, p2, |v13|];
		globalState.f3 := v85[0xa2];
		globalState.f3 := |(fm28(p1, globalState) + v0)|;
		var v86: map<int, int> := map[p3 := p1];
		var v87 := DC8(v86);
		var v89: seq<bool> := [true];
		var v90: multiset<int> := multiset{p1};
		var v91: multiset<seq<bool>> := multiset{v89[p3 := true], fm31(p0, p0, fm3(v85, v90, p0, globalState), globalState)};
		var v92 := 'v';
		var v93 := new C6(p1, v13[fm26(v87.(cf12 := v86), |(map v88 : int | (0x33 <= v88) && (v88 < -278) :: (v88 % p2) := (true))|, v91, globalState) := v92]);
		r0 := v92;
		r1 := v0;
		r2 := p3 < p1;
	}
	method m6(p0: int, p1: array<bool>, p2: int, globalState: GlobalState) {
		var v0 := true;
		p1[108] := v0;
		p1[108] := false;
		var v1 := 'm';
		var v2: multiset<char> := multiset{v1, v1, v1, v1};
		var v3: seq<char> := [v1];
		var v4 := DC44(v2);
		var v5: map<int, int> := map[|v3| := p2];
		var v6 := DC8(v5);
		var v7: seq<bool> := [v0, p1[108], v0];
		var v8: multiset<seq<bool>> := multiset{v7, v7};
		var v9 := DC7(v1, false, p2);
		v2, v0, p1[108] := v2 - (multiset(v3) - v4.cf65), if (fm26(v6, p0, v8, globalState) == fm20(p1[108], globalState)) then (v9.(cf9 := v1, cf10 := p1[108])).cf10 else p1[108], v0;
		p1[108] := !p1[108];
		var v10: C5 := new C5();
		var v11: map<bool, C5> := map[p1[108] := v10];
		for i0 := |v11| to p2 {
			var v12: array<int> := new int[18](i1 => i1 - 0x136);
			v12[424] := 632;
			var v13 := DC24(p0, v0, v3);
			v12[424] := fm26(v6, v13.cf31, v8, globalState) % p2;
			var v15: array<map<int, bool>> := new map<int, bool>[23](i2 => (map v14 : int | (-0x2d0 <= v14) && (v14 < 0x68) :: (v14 / |(seq(0x35a, i3  => (v1)))|) := (p1[108])) + map[0x1b := p1[108]]);
			var v16: map<int, bool> := map[v12[424] := v0];
			var v17: map<int, bool> := map[-0x3b1 := if (|v7| in v16) then v16[|v7|] else v0];
			v15[802] := v17;
			v15[802], p1[108], v3 := v16, v0, (seq(0x8b, i4  => (v1))) + (if (v0) then "fpl" else v3);
			v0 := v0;
			var v18 := new C3();
		}
		var v19: set<string> := {v3};
		if (v19 !! v19) {
			var v20: seq<int> := [p0];
			var v21: multiset<int> := multiset{p2};
			var v22: set<int> := {p2};
			var v23: C0 := new C0();
			var v24 := DC30(p1[108], v10.fm3(v20, v21, p1[108], globalState), v22 - v22, v23);
			v24 := DC30(v22 <= v22, p1[108], v22, v23);
			var v25: multiset<string> := multiset{v3, v3};
			var v26 := DC23(v25);
			match (v26) {
				case DC24(cf31, cf32, cf33) =>
					var v27: array<int> := new int[2](i5 => i5 * p2);
					v27[117] := -106;
					v27[117] := cf31;
					p1[108] := p1[108];
					globalState.f8 := p2;
					var v28: array<char> := new char[19](i6 => v1);
					var v29: map<array<char>, int> := map[v28 := 0x1cb];
					v29 := v29;
				case DC23(cf30) =>
					var v30 := new C0();
					v0 := p1[108];
					globalState.f8 := p0;
					var v31 := DC5(p1);
					v31 := v31;
			}
			
			var v32 := new C1();
			var v33 := DC18();
			var v34: map<bool, D7> := map[v24.cf43 := v33];
			var v35: map<int, bool> := map[|v22| := p1[108]];
			if (!(v34 == fm48(v20, v35, globalState))) {
				globalState.f8 := p2;
				var v36: map<seq<int>, multiset<int>> := map[[p0, p0, p2, p2, |v7|] := v21 * v21];
				var v37: seq<multiset<int>> := [v21];
				v36 := v36[v20 := v37[p0]];
				p1[108] := (DC4(p1[108]).(cf4 := v32.fm3(v20, v21, p1[108], globalState))).cf4;
				var v38: array<bool> := new bool[27];
				v38 := v38;
				var v39: array<seq<bool>> := new seq<bool>[3];
				v39[872] := v7;
				v39[872] := v7;
			} else {
				globalState.f4 := true;
				var v40 := DC39(false, |v3|);
				var v41: map<char, int> := map[v1 := p2];
				globalState.f8 := -v40.cf59 * (if (v1 in v41) then v41[v1] else p2);
				p1[108] := true;
				globalState.f3 := 0x27 * fm20(p1[108], globalState);
				p1[108] := p2 !in v35;
			}
			
			var v42: array<map<bool, int>> := new map<bool, int>[14](i7 => map[v0 := p2]);
			var v43: map<bool, int> := map[false := |v3|];
			v42[240] := v43 + v43;
			v42[240] := v43;
		} else {
			var v44: set<bool> := {v0, !true, p1[108]};
			globalState.f3 := |(v44 + v44)|;
			var v45: array<array<int>> := new array<int>[28];
			var v46: multiset<int> := multiset{0xf4};
			var v47: array<array<array<int>>> := new array<array<int>>[21] [v45, v45, v45, if (fm3([p2, p0, p0], v46, p1[108], globalState)) then v45 else v45, v45, if (!v0) then v45 else v45, v45, v45, v45, v45, v45, v45, v45, v45, v45, v45, v45, v45, v45, v45, v45];
			v47[320] := v45;
			var v48: array<char> := new char[11] [v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1];
			var v49: set<array<bool>> := {p1};
			var v50: seq<set<array<bool>>> := [{p1, p1}, v49];
			var v51: seq<int> := [|v50[|multiset(v7)|]|];
			v47[320], v3, v48 := v45, (if (v10.fm3(v51, multiset{436}, p1[108], globalState)) then seq(0x304, i8  => (v1)) else v3 + v3)[|v51| := v1], DC45(v48).cf66;
			var v52: array<int> := new int[7];
			var v53 := DC36(v52);
			match (v53) {
				case DC36(cf55) =>
					globalState.f4 := p1[108];
					var v54: seq<array<bool>> := [p1];
					var v55: array<array<bool>> := new array<bool>[16] [p1, p1, p1, p1, p1, p1, p1, p1, v54[190], p1, p1, p1, p1, p1, p1, p1];
					var v56: array<D13> := new D13[13];
					var v57: set<array<D13>> := {v56, v56};
					v3, v55, globalState.f8 := v3, v55, |v57| - (if (p2 in v5) then v5[p2] else p2);
					globalState.f8 := p0;
					var v58: map<string, string> := map["sl" := seq(0x2c7, i9  => (v1))];
					v0 := v7[|v58|] || v0;
			}
			
			globalState.f4 := v46 !! v46;
			var v59: map<int, set<bool>> := map[p0 := v44];
			var v60: seq<array<bool>> := [p1];
			var v61: array<array<bool>> := new array<bool>[22] [p1, p1, p1, p1, p1, p1, v60[67], p1, v60[|v51|], p1, p1, p1, p1, p1, p1, p1, p1, p1, p1, p1, p1, p1];
			v61[770] := p1;
			var v63 := DC27(|(map v62 : int | (-0xd8 <= v62) && (v62 < 152) :: (v62 - p2) := (p1[108]))|, p2, |v3|, -p0);
			globalState.f3, v59, globalState.f4, v1, v61[770] := v63.cf39, v59[p0 := v44] + v59, false, v1, p1;
		}
		
		var v65: map<char, map<char, bool>> := map[v1 := map v64 : char | v64 in v3 :: (v64) := (p1[108])];
		var v66: seq<int> := [p2, p2, p2, |v65|];
		if (v66 in (seq(0x2d4, i10  => (v66)))) {
			globalState.f8 := -p2;
			var v67: map<bool, int> := map[p1[108] := p0];
			var v68: map<seq<int>, C5> := map[v66 := v10];
			globalState.f8 := ((if (v0 in v67) then v67[v0] else fm26(v6, p2, v8, globalState)) / |v7|) + (fm10(globalState) / |v68|);
			if (p1[108]) {
				var v69 := DC27(0x354, p0, 0x376, p0);
				var v70: array<D12> := new D12[18] [v69, v69, v69, v69, v69, DC27(|map[0xc9 := p2]|, |v3|, fm20(true, globalState), p0), DC27(|v5|, |v3|, p2, p2), v69, v69, v69, v69, v69, v69, DC27(p2, |v67|, p2, p0), v69, v69, v69, v69.(cf39 := p0, cf36 := 183)];
				var v71: array<array<D12>> := new array<D12>[14] [if (v0) then v70 else v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, if (p1[108]) then v70 else v70, v70];
				v71[102] := v70;
				v71[102] := v70;
				v67 := v67[v0 := p0];
				v0 := p1[108];
				globalState.f8 := |({v1} + (set v72 : char | v72 in v3 :: (v72)))|;
				var v73: map<int, map<bool, int>> := map[p2 * |v66| := fm46(globalState)];
				var v74: multiset<bool> := multiset{v0};
				v73 := v73[0x258 / |v74| := v67[v0 := p0]];
			} else {
				var v75: multiset<int> := multiset{p2};
				var v76 := DC10([p1[108], fm3(v66, v75, false, globalState), p1[108]]);
				var v77: map<string, int> := map[v3 + v3 := p0];
				v3, v76, globalState.f8, globalState.f8 := v3, DC10(v7[p0 := !v0]), if (v3 in v77) then v77[v3] else p2, |(if (v0) then v3 + (seq(230, i11  => (v1))) else seq(0x38c, i12  => (v1)))|;
				v9 := fm40(v0, globalState);
				var v78: map<array<bool>, int> := map[p1 := fm10(globalState)];
				v78 := v78;
				var v79: array<char> := new char[1];
				v79[208] := v1;
				var v80: array<array<bool>> := new array<bool>[8];
				var v81: set<int> := {p0, |v3|, p0};
				v1, v0, globalState.f8, v79[208], v80 := 't', p1[108], p0, fm18(|v3| <= |v81|, v0, globalState), v80;
				var v82: C1 := new C1();
				v82 := v82;
			}
			
			globalState.f4 := !(("dgljestqn" + v3) != (v3 + v3));
			var v83: set<bool> := {!p1[108], v0};
			var v84: map<array<bool>, bool> := map[p1 := v0];
			var v85: map<int, bool> := map[p0 := p1[108]];
			var v86: set<int> := {p2};
			var v87: C0 := new C0();
			var v88 := DC30(v0, v0, v86, v87);
			var v89 := DC31(v88);
			var v90: set<D13> := {v89};
			var v92: array<bool> := new bool[20] [v0 && v0, -957 <= p2, true, |v83| <= fm20(v0, globalState), if (p1 in v84) then v84[p1] else p1[108], p1[108], v0, 0x2b7 <= p0, true, if (-0x3ba in v85) then v85[-0x3ba] else v0, false, p2 <= p0, v90 >= v90, p1[108], true, v0, true, |(set v91 : int | (0x141 <= v91) && (v91 < 692) :: (v91 + p2))| <= v66[p2], p1[108], p1[108]];
			var v93: map<string, bool> := map[v3 := p1[108]];
			var v94: map<map<string, bool>, bool> := map[v93 := v0];
			p1[108], v92 := if (v93 in v94) then v94[v93] else p1[108], v92;
		} else {
			v0 := v0;
			var v95 := new C6(0x3c6, v3);
			globalState.f8 := --v95.f10;
			var v96: set<int> := {p0};
			globalState.f3 := p0 - (|v96| / p2);
			var v97: array<int> := new int[3];
			var v98: multiset<array<int>> := multiset{v97};
			v98, globalState.f3, globalState.f3, globalState.f8 := multiset{v97, v97, v97}, fm20(if (true) then p1[108] else p1[108], globalState), -0x10c, p2;
		}
		
	}
}

class C8 {
	constructor () {
	}
	
	function fm8(p0: int, p1: bool, p2: int, p3: seq<string>, globalState: GlobalState): string {
		"wsbmpw" + (seq(-960, i0  => ('u')))
	}
	function fm9(p0: bool, p1: set<bool>, globalState: GlobalState): int {
		|map[|multiset{0x243, |"tspj"|}| := |(seq(0x214, i0  => ("coabyxfb")))|]| + |(map[false := true] + map[true := true])|
	}
	method m4(p0: bool, p1: map<bool, string>, p2: int, globalState: GlobalState) {
		for i0 := p2 + p2 to p2 - p2 {
			var v0 := DC2();
			match (v0) {
				case DC1(cf1, cf2) =>
					var v1: array<D0> := new D0[4];
					var v2 := DC0(0xb3);
					v1[571] := v2;
					var v3 := 'n';
					v1[571], v3 := v2, v3;
					var v4: array<string> := new string[16];
					var v5 := "vvute";
					v4[49] := v5;
					v4[49] := "rjcj";
					var v6: array<char> := new char[12];
					v6[492] := v3;
					v6[492] := v3;
					var v7: map<int, bool> := map[cf2 := p0];
					var v8: map<map<int, bool>, bool> := map[v7 := p0];
					globalState.f4 := (|(seq(49, i1  => (v3)))| * |v8|) >= p2;
				case DC2() =>
					globalState.f4 := p0 || p0;
					var v9: map<int, bool> := map[p2 := p0];
					var v10: multiset<map<int, bool>> := multiset{v9};
					var v11: set<int> := {i0, i0};
					var v12: multiset<bool> := multiset{p0};
					var v13: seq<int> := [p2];
					var v14: array<bool> := new bool[18] [true, false && p0, p0, p0, true, p0, v9 in v10, p0, p0, v11 > v11, |v11| == i0, (if (false in v12) then v12[false] else |v13|) <= i0, p0, p0 ==> !p0, p0 <==> p0, p0, p0, true];
					v14[679] := p0;
					var v15 := DC4(p0);
					var v16: seq<bool> := [p0, v15.cf4];
					var v17 := DC5(v14);
					var v18: multiset<int> := multiset{|v13|, |fm8(p2, p0, i0, seq(0x1d8, i2  => ("iutevq")), globalState)|, |(map[v14 := v17.cf5])[v14 := v14]|};
					v14[679] := v16[-|v18|];
					v14[679] := p0;
					globalState.f3 := ((if (p2 in v18) then v18[p2] else |v13|) * i0) / i0;
				case DC0(cf0) =>
					var v19 := "dlnuduto";
					var v20 := DC6(|v19|, p0, p0);
					var v21 := DC1(v20.cf6, -464);
					var v22: map<bool, D0> := map[p0 := v21];
					v22 := v22;
					globalState.f3 := i0;
					var v23: array<string> := new string[8];
					v23 := new string[29](i3 => "irpd");
					var v24: set<bool> := {p0, p0};
					var v25: map<set<bool>, int> := map[v24 := fm9(p0, v24, globalState)];
					v25 := v25[v24 := cf0];
				case DC3(cf3) =>
					var v26: array<map<bool, int>> := new map<bool, int>[10];
					var v27 := new C4(v26);
					var v28: multiset<int> := multiset{i0, 0x1ec};
					var v29 := "sjohkdo";
					var v30: seq<bool> := [p0];
					var v31: seq<int> := [p2];
					var v32: map<int, string> := map[|v28| := v29];
					var v33: array<bool> := new bool[17] [i0 in v28, p0, v27.fm13(p0, v29, p0, globalState), !!!p0, p0 <== v30[|v31|], p2 == i0, v27.fm13(p0, v29, p0, globalState), true, p0, p0, if (p0) then p0 else p0, p0, p0, p0, p0, p0, v27.fm13(p0, if (0x337 in v32) then v32[0x337] else v29, p0, globalState)];
					var v34: set<bool> := {p0, true};
					globalState.f8, globalState.f4, v33, globalState.f8 := -0x3d, {p0, false, !p0, p0, p0} <= (v34 * {v27.fm3(v31, multiset{p2}, p0, globalState)}), v33, i0;
					var v35 := DC10(v30);
					v35 := v35;
					var v36: map<int, bool> := map[p2 % i0 := p0];
					v36 := v36[i0 := p2 != i0];
			}
			
			var v37 := 's';
			v37 := v37;
			v37 := v37;
			var v38: array<seq<bool>> := new seq<bool>[26];
			v38[175] := fm31(false, !p0, p0, globalState);
			var v39: seq<bool> := [p0];
			var v40: map<bool, bool> := map[p0 := p0];
			v38[175] := (v39 + (v39[-0x29e := if (p0 in v40) then v40[p0] else p0])[330 := p0]) + (v39 + v39);
		}
		for i4 := p2 to p2 * p2 {
			var v41: seq<bool> := [p0];
			v41 := v41;
			globalState.f8 := p2;
			var v42 := 'a';
			var v43 := "vstrg";
			globalState.f4 := v42 !in v43;
			var v44 := new C6(i4, v43 + v43);
		}
		globalState.f4 := !true;
		var v45: seq<bool> := [p0, p0, p0, p0];
		var v46 := "am";
		var v47 := 'b';
		v45 := ([p0, p0] + v45)[|((if (p0 in p1) then p1[p0] else v46) + v46)[p2 := v47]| := false];
		if (fm49(p0, globalState)) {
			var v48: array<bool> := new bool[29](i5 => false);
			v48[334] := p0;
			v48[283] := p0;
			var v49: set<array<bool>> := {v48};
			globalState.f8, globalState.f3, v48[334], v48[283] := |({v48} + v49)|, if (p0) then p2 else p2 - p2, p0, p0;
			var v50: array<int> := new int[5](i6 => i6 % p2);
			var v51 := DC36(v50);
			var v52: array<D17> := new D17[9] [v51, v51, if (p0) then DC36(v50) else v51, v51, v51.(cf55 := v50), DC36(v50), v51, v51, v51];
			v52[306] := v51;
			v52[306] := v51.(cf55 := v50).(cf55 := v50);
			var v53: map<int, int> := map[p2 := p2];
			v53 := v53[|v45| := p2];
			var v54: set<int> := {-0x31f, p2};
			var v55 := DC9();
			var v56: set<D3> := {fm50(globalState), v55, v55};
			var v57: seq<set<D3>> := [v56, v56];
			globalState.f8 := |(v54 + v54)| - (if (p0) then |v57[p2]| else |v45|);
			var v58: array<string> := new string[4];
			v58[384] := v46 + v46;
			v58[384] := v46[p2 := v47];
		} else {
			var v59: map<string, int> := map[v46 := p2];
			var v60: map<bool, int> := map[p0 := p2];
			v59 := v59[v46 := -(if (p0 in v60) then v60[p0] else -|([p2, p2, |(seq(800, i7  => (v47)))|])[|v46| := p2]|)];
			var v61: T0 := new C2();
			var v62: seq<T0> := [v61];
			v61 := v62[p2];
			v46 := ("tvyfq" + v46) + v46;
			var v63 := new C3();
			var v64: array<string> := new string[24];
			v64[474] := v46;
			v64[474] := v46;
		}
		
		var v65: array<set<array<char>>> := new set<array<char>>[26];
		var v66: array<char> := new char[13] [v47, v47, v47, v47, v47, v47, v47, v47, v47, v47, v47, v47, v47];
		var v67: set<array<char>> := {v66, v66};
		v65[546] := v67;
		var v68 := DC39(p0, 0x2fa);
		var v69: array<bool> := new bool[6] [!("oxey" <= "brbh"), p0, v68.cf58, p0, fm14(p2, globalState) != v46, false];
		v69[712] := p0;
		var v70: map<int, set<array<char>>> := map[p2 := v67];
		v65[546], globalState.f4, v69[712] := v67 - (if (p2 in v70) then v70[p2] else v67), p0, (fm51(p2, globalState)).cf4;
	}
}

class C9 {
	const f9 : bool
	constructor (f9 : bool) {
		this.f9 := f9;
	}
	
	method m2(globalState: GlobalState) returns (r0: bool, r1: int, r2: set<int>) {
		match (fm5(globalState)) {
			case DC1(cf1, cf2) =>
				var v0: array<bool> := new bool[10];
				v0 := v0;
				var v1 := 'h';
				var v2: map<bool, char> := map[f9 := v1];
				var v3: map<int, map<bool, char>> := map[cf1 + cf1 := v2];
				var v4: map<bool, int> := map[f9 := cf1];
				var v5: seq<bool> := [f9];
				v3 := map[fm6(if (f9 in v4) then v4[f9] else |v5|, globalState) := fm7(v1, globalState)];
				r1 := cf2;
				if (f9) {
					var v6 := new C7();
					var v7: array<C0> := new C0[23];
					var v8: C0 := new C0();
					v7[100] := v8;
					v7[100] := v8;
					var v9: seq<int> := [cf1];
					globalState.f4 := v6.fm3(v9, multiset(v9), f9, globalState);
					var v10 := new C8();
					var v11 := DC5(v0);
					var v12: array<array<bool>> := new array<bool>[11] [v0, v0, v0, v0, v0, v0, v0, v0, v0, v11.cf5, v0];
					v12[812] := v0;
					v12[812] := v0;
				} else {
					globalState.f8 := 0x8e % -cf2;
					var v13: array<int> := new int[23](i0 => i0 * cf1);
					var v14: map<bool, bool> := map[f9 := false];
					var v15: map<array<int>, map<bool, bool>> := map[v13 := v14];
					var v16: map<int, bool> := map[cf1 - --|v15| := f9];
					var v18: set<int> := {-0x19};
					v16 := v16 + (map v17 : int | v17 in v18 :: (v17 - cf1) := (f9))[cf1 := f9];
					var v19 := "fml";
					var v20 := new C6(cf2, v19 + v19);
					var v21: set<bool> := {f9, f9, f9, v5[v20.f10], f9};
					globalState.f3 := |v21| / (-0x2e5 * fm6(cf2, globalState));
					var v22 := DC26(map[if (f9 in v14) then v14[f9] else f9 := cf2]);
					var v23: map<D12, bool> := map[v22 := true];
					var v24: map<int, map<D12, bool>> := map[cf2 := v23];
					globalState.f3, r0 := |(v24 + map[|v20.f11| := v23])|, f9;
				}
				
			case DC2() =>
				var v25: seq<bool> := [f9, !f9];
				var v26: map<bool, int> := map[true := |v25|];
				var v27 := -901;
				var v29 := "ivrfg";
				var v30: seq<int> := [|v29|];
				var v31: array<map<bool, int>> := new map<bool, int>[18] [v26, v26, v26, v26, fm46(globalState), v26, map[false := v27], v26, v26, v26, map[f9 := v27], map[f9 := |(map[v27 := f9])[|(map v28 : int | (993 <= v28) && (v28 < 0x21b) :: (v28 - |"jkmh"|) := (f9))| := f9]|], v26, v26, v26, map[true := |[v27, |multiset(v30)|, v27, v27]|], v26, v26];
				var v32 := new C4(v31);
				var v33: multiset<bool> := multiset{f9};
				var v34: map<bool, string> := map[f9 := "fal"];
				var v35 := DC11(0x1a1 < |(v33[true := 0x344])[f9 := v27]|, --|(if (f9 in v34) then v34[f9] else v29)| >= v27, v27, v29);
				var v36: map<seq<bool>, bool> := map[v25 := f9];
				var v37: map<bool, bool> := map[f9 := f9];
				var v38 := DC27(0x10, v27, v27, |v37|);
				var v39: array<bool> := new bool[20] [f9, f9, f9, f9, f9, f9, f9, f9, f9, f9, f9, f9, f9, false, false, f9, f9, f9, f9, f9];
				var v40: multiset<array<bool>> := multiset{v39, v39, v39};
				var v41: map<int, int> := map[v27 := v27];
				var v42 := DC8(v41);
				var v43: multiset<seq<bool>> := multiset{v25};
				var v44: C6 := new C6(v27, if (f9 in v34) then v34[f9] else seq(706, i1  => (v29[v27])));
				var v45: map<bool, C6> := map[f9 := v44];
				var v46: array<int> := new int[23] [-|v36|, v38.cf39 - v27, v27, |v40[v39 := v27]|, if (f9) then |v25| else fm6(v27, globalState), |v41| / v27, v27, -926, v27, if (f9 in v33) then v33[f9] else 0x343, v27 % v27, fm26(v42, v27, v43, globalState) / v27, --|(v45[f9 := v44])[DC21(f9, v25, true).cf26 := v44]|, v27, v44.f10, -v27, -v27, v27 - 0x34, v27, v27, -v27, |[v44]|, v27];
				var v47: seq<array<int>> := [v46, v46, v46, v46];
				var v48: seq<array<int>> := [v46, v46, v46, v47[v44.f10]];
				v35, v46, r0 := DC11(f9, true, -501, v44.f11 + v44.f11), v47[|v25|], f9;
				var v49 := new C4(v32.f12);
				r1 := fm20(f9, globalState);
			case DC0(cf0) =>
				var v50: map<bool, int> := map[f9 := cf0 * 944];
				var v51: seq<int> := [cf0];
				v50 := v50[f9 := cf0 + |v51|];
				globalState.f8 := cf0;
				globalState.f4 := cf0 <= -cf0;
				var v52: array<bool> := new bool[10](i2 => f9);
				var v53 := "jwsx";
				var v54: map<int, array<bool>> := map[|v53| := v52];
				var v55: multiset<int> := multiset{cf0, |[-0x3d8, -0x210]|, cf0};
				var v56: array<array<bool>> := new array<bool>[14] [v52, v52, if (|v55| in v54) then v54[|v55|] else v52, v52, v52, v52, v52, v52, v52, v52, if (true) then v52 else v52, DC5(v52).cf5, v52, v52];
				v56[707] := v52;
				v56[707] := v52;
			case DC3(cf3) =>
				if (f9) {
					var v57: array<seq<bool>> := new seq<bool>[1];
					v57 := v57;
					var v58 := "ljehk";
					var v59: seq<string> := [v58, "pbuyasjcd"];
					var v60: seq<bool> := [[v58, v58, v58, seq(0x2a5, i3  => ('d'))] != v59];
					v60 := v60;
					var v61 := new C5();
					v58 := "mrlflpdvu";
					var v62 := 'm';
					var v63: map<bool, char> := map[f9 := v62];
					var v64: map<map<bool, char>, string> := map[v63 := v58];
					v64 := v64[v63 + v63 := v58];
				} else {
					var v65 := "kypi";
					var v66: map<int, bool> := map[|v65| := false];
					var v67 := 0x33f;
					v66 := v66[fm6(v67, globalState) := false];
					globalState.f3 := -v67;
					globalState.f3 := v67;
					var v68: array<multiset<bool>> := new multiset<bool>[24];
					var v69: seq<bool> := [f9];
					v68[716] := multiset(v69);
					v68[716] := (multiset{f9, f9, !!f9})[f9 := v67];
					var v70: set<bool> := {!f9};
					r2 := {v67, v67 % v67, |({true, f9, f9, f9, f9} * v70)|, -(v67 * 293), v67};
				}
				
				var v71: seq<int> := [0x186];
				var v72 := 0x25b;
				var v73: set<seq<int>> := {v71 + v71, (seq(0x18b, i4  => (-0x37f)))[v72 := 206], v71};
				v73 := v73;
				var v74 := 'f';
				var v75: map<char, char> := map[v74 := v74];
				var v76: array<int> := new int[18];
				v76[282] := -|map[v76 := f9]|;
				var v77: array<bool> := new bool[3](i5 => f9);
				var v78: map<array<bool>, map<char, char>> := map[v77 := v75];
				var v79: map<int, int> := map[v72 := v72];
				var v80: map<int, bool> := map[|v79| := true];
				var v81: multiset<bool> := multiset{f9};
				v75, v76[282] := if (v77 in v78) then v78[v77] else map[v74 := fm18(if (v72 in v80) then v80[v72] else f9, f9, globalState)], if (f9) then v72 * |v81| else v71[0x173];
				var v82 := "ktuiokx";
				v82 := fm25(globalState);
		}
		
		var v83: array<int> := new int[27];
		v83[232] := fm20(f9, globalState);
		var v84 := 382;
		var v85: seq<bool> := [f9];
		var v86: seq<int> := [v84, v84, v84, |v85|];
		var v87: map<bool, bool> := map[f9 := f9];
		v83[232] := fm6(v84 % |multiset(v86[fm6(|v87|, globalState) := v84])|, globalState);
		var v88 := "ojpl";
		var v89 := DC11(f9, f9, |[v84, v83[232]]|, v88);
		for i6 := v89.cf16 to |(seq(722, i7  => ('d')))| {
			var v90 := DC10([f9]);
			var v91: map<bool, seq<bool>> := map[fm49(f9, globalState) := [f9] + v90.cf13];
			r1 := |v91|;
			r0 := f9;
			globalState.f8 := i6;
			v88 := v88;
		}
		globalState.f8, v83[232] := v83[232] + v83[232], |((v85 + v85[0x1a5 := f9]) + [f9])|;
		var v92: T0 := new C7();
		var v93: seq<T0> := [v92];
		var v94 := DC41(v93);
		v94 := DC41(v93);
		globalState.f3 := v83[232];
		var v95 := DC6(v83[232], f9, f9);
		r0 := match v95 {
			case DC6(cf6, cf7, cf8) => if (f9) then cf8 else !cf7
			case DC7(cf9, cf10, cf11) => f9
			case DC5(cf5) => multiset{0x211} !! multiset{v83[232]}
		};
		r1 := v84;
		var v96: set<int> := {v84};
		r2 := v96;
	}
	method m3(p0: int, p1: int, globalState: GlobalState) returns (r0: array<int>, r1: int, r2: int, r3: bool) {
		var v0: seq<bool> := [f9];
		var v1: seq<seq<bool>> := [v0, [false], v0, v0, v0];
		var v2: map<seq<bool>, int> := map[v0 := -|v1[p1]|];
		var v3 := "julgoho";
		var v4 := 'o';
		var v5: multiset<int> := multiset{(if (v0 in v2) then v2[v0] else p0) + p0, -539, if (f9) then |v3| else -|fm17(v4, globalState)|, -p0, p0};
		globalState.f3 := if (p0 in v5) then v5[p0] else p0;
		var v6: set<int> := {p0};
		if (f9 ==> (v6 !! v6)) {
			var v7: seq<int> := [989, p0];
			var v8: map<int, int> := map[p0 := 0x349];
			var v9 := DC8(v8);
			var v10: multiset<seq<bool>> := multiset{v0};
			globalState.f0 := (v7 + v7)[p1 * p1 := fm26(v9, 0x2af, v10, globalState)];
			var v11: array<int> := new int[17];
			var v12: set<array<int>> := {v11, v11};
			var v13: array<map<bool, int>> := new map<bool, int>[19];
			var v14: T0 := new C4(v13);
			var v15: seq<array<int>> := [v11];
			r3, globalState.f4, v12, globalState.f3, v14 := f9, f9, {v11, if (f9) then v11 else v11, v15[p1]}, --0x309, v14;
			if ((p0 + p0) < p0) {
				globalState.f3 := p1 % (fm6(fm20(f9, globalState), globalState) - p0);
				var v16: set<char> := {v4, fm18(f9, f9, globalState), fm18(f9, false, globalState), v4};
				v16 := v16 * ({v4} - {DC35(p0, v4, |(seq(-0x265, i0  => (v4)))[p1 := v4]|, p0, p1).cf51});
				globalState.f4 := if (f9) then f9 else f9;
				var v17: array<bool> := new bool[22](i1 => f9);
				var v18: seq<array<bool>> := [v17, v17, v17];
				v18 := if (f9) then v18 else v18;
				v0 := ((v0[p0 := f9])[p0 := f9])[fm20(f9, globalState) % p0 := f9];
			} else {
				var v19: array<multiset<array<bool>>> := new multiset<array<bool>>[11];
				var v20 := DC43(v7, f9);
				var v21: array<bool> := new bool[21] [f9, v14.fm3(v20.cf63, v5, f9, globalState), false, f9, f9, f9, f9, f9, f9, f9, false, f9, f9, f9, f9, f9, f9, f9, f9, f9, f9];
				var v22: multiset<array<bool>> := multiset{v21};
				v19[722] := v22;
				v19[722] := multiset{v21, v21, v21};
				var v23: map<seq<bool>, seq<int>> := map[v0 := v7];
				var v24: seq<map<seq<bool>, seq<int>>> := [v23];
				v23 := v24[-0x164];
				v21[664] := fm49(fm49(f9, globalState), globalState);
				v21[664] := if (false) then f9 else f9;
				globalState.f3 := p1;
				globalState.f4 := f9;
			}
			
			var v25: set<bool> := {f9, f9};
			var v26: map<int, bool> := map[p0 := v25 !! v25];
			v26 := v26[|map[v7 := v11]| := f9];
			v11[8] := p0;
			var v27: seq<seq<char>> := [v3];
			var v28 := DC24(p1, !f9, v27[|(seq(0x2bb, i2  => (v3)))|]);
			v11[8] := v28.cf31;
		} else {
			v3 := v3;
			var v29: array<bool> := new bool[5];
			v29[645] := f9;
			v29[645] := v6 != {p0, p1, p1};
			if (v29[645]) {
				var v30: map<int, bool> := map[p1 := !v29[645]];
				globalState.f8 := p1 * -|v30|;
				globalState.f4 := v29[645];
				var v31: map<int, string> := map[p0 := v3 + v3];
				v31 := v31[|(if (f9) then v5 else v5)| := v3];
				var v32: array<int> := new int[25];
				v32[30] := p0 * p1;
				v32[30] := p0;
				var v33: C6 := new C6(p1, v3);
				var v34: map<C6, bool> := map[v33 := !v29[645]];
				v34 := v34[v33 := f9];
			} else {
				var v35: map<int, int> := map[p0 := p1];
				var v36: seq<int> := [-p1, p0];
				var v37 := DC43(v36, v29[645]);
				var v38 := DC19(v36);
				v35, globalState.f4, v29, v37, v29[645] := map[p1 := p1] + map[p1 := p1], f9, v29, DC43(v36[p1 := p0], [v38.(cf24 := seq(0x336, i3  => (p0))), v38, v38] <= (seq(0x37f, i4  => (v38)))), v29[645];
				var v39: array<map<bool, int>> := new map<bool, int>[18];
				var v40: map<bool, int> := map[v29[645] := p1];
				v39[277] := v40;
				var v41: array<multiset<int>> := new multiset<int>[26];
				var v42: multiset<bool> := multiset{v29[645], f9};
				var v43: map<bool, bool> := map[v29[645] := f9];
				globalState.f4, globalState.f0, v39[277], v41 := v42 > (v42 * v42), (v36 + v36[p1 := p1])[|v43| := p1], v40 + (v40 + v40), v41;
				var v45: array<map<int, int>> := new map<int, int>[23](i5 => map v44 : int | v44 in v6 :: (v44 - p0) := (p0));
				v45[856] := v35;
				var v46 := DC21(v29[645], [v29[645]], f9);
				var v47: set<D4> := {fm30(p1, v46, f9, p0, globalState)};
				var v50 := DC8(map v48 : int | v48 in (set v49 : int | (0x39c <= v49) && (v49 < 0xa1) :: (v49 % |v36|)) :: (v48 % p1) := (p1));
				var v51: multiset<seq<bool>> := multiset{[!f9]};
				var v52 := DC48(v47);
				v45[856], globalState.f3, v29[645], globalState.f3, v47 := v35, fm26(v50, p1 / p1, v51 - v51, globalState), fm49(f9, globalState), p0 / -p1, (v47 - v52.cf71) + (v47 + v47);
				globalState.f4 := v0[p1];
				globalState.f3, r2 := |"dhkfosxv"| / (-|[f9]| + p1), -629;
			}
			
			var v53: map<bool, set<int>> := map[v29[645] := v6];
			var v54: map<int, int> := map[p1 := p0];
			r3, globalState.f3 := v29[645], |(map[p0 := |v53|] + v54)| % fm20(fm49(true, globalState), globalState);
			var v55: map<int, multiset<int>> := map[DC1(0x276, p1).cf2 := v5];
			var v56: map<map<int, multiset<int>>, int> := map[v55 := -p1];
			v56 := v56[v55 + v55 := 0x28c];
		}
		
		var v57: array<bool> := new bool[23];
		v57[570] := if (f9) then f9 else f9;
		v57[570] := false;
		var v58: array<char> := new char[4] [v4, 'p', v4, v4];
		v58[802] := v4;
		v58[802] := if (v57[570]) then v4 else if (!v57[570]) then v4 else 'r';
		var v59: seq<array<bool>> := [v57];
		var v60: array<array<bool>> := new array<bool>[22] [v57, v57, v57, v57, v57, v57, v57, v57, v57, if (v57[570]) then v57 else v57, v57, v57, v57, v57, v57, v57, v57, v57, v59[|v3|], v57, v57, v57];
		v60 := v60;
		var v61: array<D19> := new D19[10];
		var v62: set<array<D19>> := {v61};
		var v63: map<bool, char> := map[f9 := v58[802]];
		var v64: array<int> := new int[20] [p0, -|[f9, f9]|, p0, p0, 0x379, p1, p0, p1, fm20(v57[570], globalState), p0, p1, -p0, p1, -p0, p1, -p1, p1, |v63|, 240, p0];
		var v65: map<set<array<D19>>, array<int>> := map[v62 := v64];
		r0 := if ((v62 - v62) in v65) then v65[v62 - v62] else v64;
		r0 := v64;
		r1 := p0;
		r2 := p1;
		r3 := false;
	}
}

class C10 extends T0 {
	constructor () {
	}
	
	function fm2(globalState: GlobalState): map<int, int> {
		(map[645 := |[|[0x297]|, |{|multiset{true, true}|}|, |map[0x51 := false]|, |multiset{|multiset{"jsdxerq"}|, -|map[-0x170 := {|multiset{false, false}|}]|, |(map v0 : int | v0 in multiset{-85} :: (v0 - 0x382) := ('t'))|}|]|] + map[|{'u', 'r'}| := |(seq(0xaf, i0  => (23)))|]) + map[-0x1d := |{491}|]
	}
	function fm3(p0: seq<int>, p1: multiset<int>, p2: bool, globalState: GlobalState): bool {
		false
	}
	function fm4(p0: int, globalState: GlobalState): map<D0, bool> {
		(map v0 : D0 | v0 in [DC0(|{true, true}|)] :: (v0) := (true)) + (map v1 : D0 | v1 in map[DC0(-0x2a0) := 'r'] :: (v1) := (false))
	}
	method m0(globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0 := new C0();
		var v1 := true;
		var v2: multiset<bool> := multiset{v1};
		v2 := v2;
		var v3 := 0x1ac;
		for i0 := v3 + v3 to v3 * --v3 {
			v3 := -0x237;
			var v4: seq<bool> := [v1, v1];
			var v6: map<int, bool> := map[--i0 := v1];
			var v7: set<int> := {|(map v5 : int | v5 in v6 :: (v5 - |(seq(0x13f, i1  => ('b')))|) := (v1))|};
			var v8 := DC30(v1, v4[0x188], v7, v0);
			var v9: array<D13> := new D13[3] [v8, v8, v8];
			v9[563] := v8;
			var v10: array<bool> := new bool[13];
			v9[563], globalState.f8, v10, globalState.f8 := v8, v3, if (v4 == v4) then v10 else v10, v3;
			var v11: array<C1> := new C1[17];
			var v12: C1 := new C1();
			v11[458] := v12;
			globalState.f3, v11[458] := i0, v12;
			var v13 := "qajy";
			var v14: seq<int> := [0x3f];
			var v15: map<int, int> := map[|v13| / i0 := v14[v3]];
			v15 := v15[v3 := i0];
		}
		if (v2 > (if (v1) then v2 else v2)) {
			globalState.f8 := v3;
			if (v1) {
				var v16: array<map<bool, char>> := new map<bool, char>[21](i2 => map[v1 := 'h'] + map[v1 := 'b']);
				v16 := v16;
				var v17: C1 := new C1();
				var v18: map<C1, int> := map[v17 := 0x2f9];
				globalState.f3 := v3 * (if (v17 in v18) then v18[v17] else v3);
				r0, globalState.f4 := v1, v1;
				var v19: array<map<bool, int>> := new map<bool, int>[14];
				var v20: T0 := new C4(v19);
				var v21: set<T0> := {v20};
				var v22: map<int, set<T0>> := map[0x382 := v21];
				var v23: seq<set<T0>> := [v21];
				var v24: array<set<T0>> := new set<T0>[22] [v21 - {v20, v20, v20}, v21, v21 - v21, v21 * v21, v21, v21 + v21, v21 - v21, v21, if (v3 in v22) then v22[v3] else v21, v21 - {v20, v20, v20}, v21, v23[v3], v21, v21, v21, v21, {v20}, v21, v21, v21, {v20}, v21 - {v20}];
				v24[499] := v21;
				v24[499] := {v20};
				var v25 := 'h';
				v25 := v25;
			} else {
				var v26: array<array<string>> := new array<string>[10];
				var v27 := "sdej";
				var v28 := 'w';
				var v29 := DC24(v3, v1, v27);
				var v30: array<string> := new string[28] [v27, "uvvaguomr", "vv", v27, (seq(162, i3  => ('k')))[v3 := v28], v27, v27, v27, "r", fm25(globalState), v27, seq(0x13f, i4  => (v28)), "gwsellrim", v27[v3 := v28], v27, (seq(-0xb6, i5  => (v28)))[0x160 := v28], v27, v27, seq(0x120, i6  => (v28)), v29.cf33, v27, "p", v27, v27, "fuxk", v27, "lmxp", "gsdb"];
				v26[429] := v30;
				v26[429] := new string[28];
				var v31: map<int, bool> := map[v3 := true];
				var v32: seq<string> := [v27, "tgmejkpo", v27, v27, v27];
				var v33: set<bool> := {v1, v1};
				var v34: array<int> := new int[26] [|v2|, v3, -|v31|, v3, -|v32|, v3, v3, v3, v3, -v3, v3, |v33|, v3, |(seq(0x25b, i7  => (v3)))|, v3, v3, v3, v3, v3, v3, v3, v0.fm19(globalState), v3, v3, v3, |v27|];
				var v35: array<array<int>> := new array<int>[10] [v34, v34, v34, v34, v34, v34, v34, v34, v34, v34];
				v35[394] := v34;
				v35[394], r0, v1 := v34, v1, v3 < v3;
				var v36: array<D23> := new D23[12];
				var v37: array<char> := new char[4] [v28, v28, v28, v28];
				var v38 := DC45(v37);
				v36[282] := v38;
				v36[282] := v38.(cf66 := v37);
				var v39: map<int, int> := map[v3 := v3];
				var v40 := DC7(v28, v1, 0x312);
				var v41: seq<set<bool>> := [v33];
				v39 := v39[v40.cf11 := |(v33 * v41[|"guqrkvlrc"|])|];
				globalState.f3 := -v3;
			}
			
			var v42: array<bool> := new bool[12](i8 => v1);
			v42 := v42;
			var v43: array<int> := new int[3];
			v43[857] := --180;
			var v44: map<int, int> := map[v3 := v3];
			var v45: map<map<int, int>, int> := map[v44 := v3];
			v43[857] := if (v44 in v45) then v45[v44] else v3;
			var v46 := 'd';
			var v47: seq<char> := [v46, v46, v46];
			var v48: seq<char> := [v46, v46, v46, v47[v3]];
			v47 := v47;
		} else {
			if (false) {
				var v49: array<int> := new int[21](i9 => i9 % v3);
				v49[192] := v3;
				v49[192] := v3;
				var v50: array<bool> := new bool[10](i10 => v1);
				v50[256] := v1;
				v50[136] := v1;
				var v51: map<int, bool> := map[v49[192] := v1];
				var v52: seq<bool> := [if (v3 in v51) then v51[v3] else v1];
				var v53: map<bool, int> := map[v1 := v0.fm19(globalState)];
				var v54: multiset<int> := multiset{v49[192], v3, v49[192], |v53|, |"youkw"|};
				var v55: multiset<int> := multiset{v49[192], |v54|, v3};
				var v56: multiset<multiset<int>> := multiset{v55, v54};
				var v57: map<int, multiset<multiset<int>>> := map[v49[192] := v56];
				var v58: seq<string> := [seq(0x1c7, i11  => ('f'))];
				var v59: seq<int> := [|(if (v49[192] in v57) then v57[v49[192]] else multiset([v54, v54]))|, v49[192], |v58| + v49[192]];
				v3, v50[256], v50[136], v1, globalState.f3 := ---617, v1, v52[v3], v1, v59[v49[192]];
				var v60 := 'n';
				var v61: map<bool, char> := map[v1 := v60];
				var v62: set<int> := {v3, |v61|};
				var v63: set<int> := {|v62|, v49[192]};
				var v64: set<set<int>> := {v62, v63 + {|v52|}, v62};
				var v65: array<map<int, bool>> := new map<int, bool>[20](i12 => v51);
				v65[200] := map v66 : int | (137 <= v66) && (v66 < 0xef) :: (v66 + 0x208) := (v1);
				v64, v65[200], globalState.f8 := v64, v51 + v51, v49[192];
				var v67 := "catkgef";
				var v68: map<int, string> := map[v49[192] := v67];
				v49[192] := |("ds" + (if (v49[192] in v68) then v68[v49[192]] else v67))|;
				var v69 := DC5(v50);
				var v70: array<array<bool>> := new array<bool>[13] [v50, if (v50[256]) then v50 else v50, v50, v50, v50, v50, v50, v50, v50, v69.cf5, v50, v50, v50];
				v70[183] := v50;
				v70[183] := v50;
			} else {
				var v71 := 't';
				var v72 := "gi";
				var v73: array<char> := new char[12] [v71, v71, v71, v71, v71, v71, v71, 'p', v71, v72[650], v71, if (v1) then v71 else v71];
				v73[195] := 'x';
				var v74: map<int, char> := map[v3 := v72[v3]];
				v73[195] := if (v3 in v74) then v74[v3] else v71;
				var v75: array<int> := new int[26](i13 => i13 / |multiset{-v3, 966, |v2|, v3}|);
				v75[106] := v3;
				v75[106] := fm6(v3, globalState) % -|v72|;
				var v76: C3 := new C3();
				var v77 := DC37(v76);
				var v78: seq<D18> := [v77, v77.(cf56 := v76), v77, v77];
				v78 := v78;
				globalState.f3 := |(seq(0x1b4, i14  => (v73[195])))| - v3;
				globalState.f4 := v1;
			}
			
			globalState.f3 := v3 * (v3 * |map[v3 := v1]|);
			var v79: map<bool, bool> := map[fm49(v1, globalState) := v1];
			globalState.f4 := v1 !in v79;
			var v80: array<map<bool, seq<int>>> := new map<bool, seq<int>>[22];
			var v81: seq<bool> := [v1];
			var v82: seq<int> := [-|v81|, v3];
			var v83: map<bool, seq<int>> := map[v1 := v82];
			v80[207] := v83[v1 := v82];
			v80[207] := map[v1 && v1 := seq(-0x3b0, i15  => (v3))];
			var v84 := 'j';
			var v85 := "pudx";
			r1 := v84 in v85;
		}
		
		var v86 := 'b';
		var v87: map<int, char> := map[-713 := v86];
		var v88 := DC38(v87);
		var v89 := DC40(v88);
		var v90 := DC40(v88);
		match (v90) {
			case DC39(cf58, cf59) =>
				var v91 := "imuchiwj";
				var v92: array<int> := new int[18];
				var v93: map<int, array<int>> := map[v3 := v92];
				var v94: map<int, int> := map[v3 := v3];
				v3 := |v91[|v93| := v86]| * ((if (fm20(cf58, globalState) in v94) then v94[fm20(cf58, globalState)] else v3) - 0x37a);
				r0 := cf58;
				globalState.f4 := v1;
				v92 := v92;
			case DC38(cf57) =>
				var v95: multiset<int> := multiset{-v3};
				var v96 := DC16(v95);
				v96 := v96;
				var v97: array<int> := new int[19];
				var v98: map<array<int>, int> := map[v97 := v3];
				var v99 := DC42(v98);
				var v100: set<bool> := {v1};
				var v101: map<bool, bool> := map[v1 := false];
				var v102 := DC0(|v101|);
				globalState.f2, v99, globalState.f3 := (if (v1) then v95 else v95) - (v95 - multiset{if (v1 in v2) then v2[v1] else v3, v3, |v100|, v3, 309}), v99, v102.cf0;
				var v103 := "jo";
				var v104: map<bool, int> := map[|v103| != v3 := v3];
				globalState.f3 := |v104|;
				globalState.f8 := -((v3 / |v101|) / |v2[v1 := v3]|);
			case DC40(cf60) =>
				var v105 := "fgp";
				v1, globalState.f4, r0, globalState.f4 := v1, v1, v105 < (v105[v3 := v86] + (seq(0xe6, i16  => (v86)))), v1;
				var v106: array<int> := new int[8];
				v106 := v106;
				var v107: set<char> := {v86};
				var v108: C3 := new C3();
				var v109: map<set<char>, C3> := map[v107 * v107 := v108];
				v109 := v109[v107 := v108];
				globalState.f4 := v1;
		}
		
		var v110: map<bool, char> := map[v1 := v86];
		globalState.f4 := !v1 !in (v110 + v110);
		r0 := v1;
		r1 := v1;
	}
	method m1(p0: int, p1: bool, globalState: GlobalState) returns (r0: int) {
		var v0: map<int, bool> := map[p0 := p1];
		var v1: multiset<bool> := multiset{if (p0 in v0) then v0[p0] else p1, p1, p1};
		match (fm40(v1 <= multiset{p1}, globalState)) {
			case DC6(cf6, cf7, cf8) =>
				var v2: C3 := new C3();
				var v3: multiset<C3> := multiset{v2};
				var v4: set<int> := {cf6 - p0, cf6 + |v3|, p0, 0x125 / cf6, cf6};
				var v5: array<bool> := new bool[23];
				v5[739] := !p1;
				var v6: array<string> := new string[5];
				var v7 := "wbp";
				v6[235] := v7;
				v4, v5[739], v6[235], globalState.f4 := v4 - v4, (p0 - p0) != -(p0 - cf6), v7 + v7, cf8;
				var v8 := 'y';
				var v9: map<bool, bool> := map[p1 := v8 in v6[235]];
				v9 := fm21(!v5[739], globalState);
				v5[739] := !v5[739];
				var v11: seq<int> := [|(map v10 : int | (-0x1cd <= v10) && (v10 < 451) :: (v10 - p0) := (v6[235][p0 := v8]))|];
				globalState.f0 := v11[0x2d0 := p0];
			case DC7(cf9, cf10, cf11) =>
				var v12 := new C5();
				var v13 := new C5();
				globalState.f8 := --857;
				if (v1 != v1) {
					var v14: array<int> := new int[21](i0 => i0 / cf11);
					v14 := v14;
					var v15: array<C1> := new C1[11];
					var v16: C1 := new C1();
					v15[53] := v16;
					var v17: map<bool, array<int>> := map[p1 := v14];
					var v18: T0 := new C3();
					var v19: seq<map<bool, array<int>>> := [v17, if (cf10) then v17 else v17, v17];
					var v20: seq<array<int>> := [v14, if (p1 in v17) then v17[p1] else v14];
					v15[53], globalState.f3, v17, v18, cf10 := v16, cf11, v19[p0 % cf11], v18, v14 !in v20;
					var v21: map<bool, int> := map[cf10 || p1 := -p0];
					v21 := v21[cf10 := p0];
					v14[734] := -p0;
					v14[734] := fm6(p0, globalState);
					r0 := 0x141;
				} else {
					var v22 := "k";
					v22 := v22;
					globalState.f3 := p0;
					globalState.f8 := fm20(cf10, globalState) + (0xd4 - cf11);
					globalState.f4 := cf10;
					var v23 := DC6(cf11, false, p1);
					v23 := v23;
				}
				
			case DC5(cf5) =>
				var v24: array<map<bool, int>> := new map<bool, int>[15];
				var v25: map<bool, array<map<bool, int>>> := map[p1 := v24];
				var v26 := new C4(if (p1 in v25) then v25[p1] else v24);
				var v27: array<int> := new int[26](i1 => i1 % p0);
				v27[578] := p0 - p0;
				var v28: array<array<int>> := new array<int>[11] [v27, v27, v27, v27, v27, v27, v27, v27, v27, v27, v27];
				v28[789] := v27;
				v27[578], globalState.f3, globalState.f8, v28[789] := p0, -p0 / -0x1cf, p0 - p0, v27;
				globalState.f4 := p1;
				var v29, v30 := v26.m0(globalState);
		}
		
		for i2 := p0 to p0 + fm6(p0, globalState) {
			var v31: array<array<bool>> := new array<bool>[3];
			v31 := new array<bool>[1];
			var v32: array<bool> := new bool[14](i3 => true);
			v32 := if (!p1) then v32 else v32;
			if (p1) {
				globalState.f4 := p1;
				globalState.f3 := p0;
				var v33: array<int> := new int[26](i4 => i4 / |multiset(seq(0x2e9, i5  => (i2)))|);
				var v34 := "cqbcks";
				var v35: seq<string> := [v34];
				var v36 := DC36(v33);
				globalState.f4, v33 := [v34] <= v35, v36.cf55;
				var v37: array<multiset<bool>> := new multiset<bool>[13];
				v37 := v37;
				globalState.f4 := (seq(0x24a, i6  => (v34[p0]))) <= (seq(0x2e2, i7  => ('a')));
			} else {
				var v38 := 'l';
				var v39: seq<char> := [v38, v38];
				var v40 := DC24(p0, p1, v39);
				var v41: seq<seq<char>> := [v39, v40.cf33[i2 := v38]];
				var v42: map<int, seq<char>> := map[fm20(p1, globalState) := [v38, v38, v38] + v41[i2]];
				v42 := v42[-(p0 % p0) := v39];
				globalState.f4 := p1;
				v38 := if (p1) then v38 else v38;
				var v43 := DC52(v31);
				var v44: array<array<array<bool>>> := new array<array<bool>>[27] [v43.cf73, v31, v31, v31, v31, v31, v31, v31, if (p1) then v31 else v31, v31, v31, v31, v31, v31, if (p1) then v31 else v31, v31, v31, v31, v31, v31, v31, v31, v31, v31, v43.cf73, v31, v31];
				v44[26] := v31;
				v44[26] := v31;
				globalState.f8 := p0 + i2;
			}
			
			var v45: set<bool> := {p1};
			v32[292] := v45 !! {p1, p1, false};
			v32[292] := p1;
		}
		r0 := 0x165 / (0x297 * p0);
		var v46 := 'h';
		var v47: map<int, char> := map[p0 + p0 := v46];
		v47 := v47[p0 := if (p1) then 'a' else v46];
		var v48: array<C9> := new C9[26];
		var v49 := "msd";
		var v50: seq<string> := [v49];
		var v51 := DC11(p1, p1, p0, v49);
		var v52 := DC12(v51);
		var v53: seq<D4> := [v52, DC12(v51)];
		var v54: map<array<C9>, string> := map[v48 := v50[|v53|]];
		v54 := v54 + (v54 + v54);
		globalState.f4 := fm49(p1, globalState);
		var v55: set<int> := {0xa7, p0};
		var v56: map<bool, int> := map[fm49(p1, globalState) := |v55|];
		var v57: multiset<map<bool, int>> := multiset{v56, (map[p1 := p0])[p1 := p0]};
		var v58: seq<int> := [|v57|];
		r0 := p0 + v58[p0];
	}
}

method Main() {
	var v0 := 0x221;
	var v1: seq<int> := [v0];
	var v2 := false;
	var v3: map<bool, int> := map[v2 := v0];
	var v4 := DC0(|v3|);
	var v5 := "cj";
	var v6: set<string> := {v5};
	var v7: multiset<int> := multiset{v0};
	var globalState := new GlobalState((v1[v4.cf0 := v0])[v0 := |map[v2 := |v6|]|], "w" + v5, v7, 0x2e6, false, false, true, 0x153, 0x158);
	var v8: array<int> := new int[16](i0 => i0 * |{v2, v2}|);
	var v9 := DC1(v0, v0);
	v8, globalState.f3 := v8, v9.cf1;
	var i1 := 0;
	while (v2)
		decreases 100 - i1
	{
		if (i1 >= 100) {
			break;
		}
		
		i1 := i1 + 1;
		v8[982] := v0;
		v8[982] := v0;
		match (v4) {
			case DC1(cf1, cf2) =>
				var v10: array<int> := new int[24];
				v10 := new int[14](i2 => i2 - cf2);
				globalState.f3 := 0x2cc;
				globalState.f2 := v7;
				globalState.f4, v8[982], globalState.f4 := v2, -(cf2 % v0), v0 != (cf1 * cf1);
			case DC2() =>
				var v11: array<int> := new int[3];
				v11 := v11;
				v3 := v3[!(v2 <== v2) := v8[982]];
				var v12: map<int, int> := map[v8[982] := v0];
				globalState.f0 := if (v2) then v1 else [v0, |v12|];
				v3 := v3;
			case DC0(cf0) =>
				globalState.f8 := (cf0 / v8[982]) - DC0(|fm0(cf0, globalState)|).cf0;
				globalState.f4 := v2;
				var v13: array<seq<int>> := new seq<int>[1] [v1];
				var v14: seq<bool> := [v2, v2, v2, false, false];
				var v15: set<bool> := {v2};
				var v16: map<int, int> := map[|v15| := cf0];
				v13[529] := fm1(-|[cf0, -|v14|, |v5|, 0x1c6, |v5|]|, 'f', !v2, v16, globalState);
				v13[529] := v1;
				v0 := -(0x2a9 % v8[982]);
			case DC3(cf3) =>
				globalState.f8 := -v8[982];
				var v17 := new C5();
				var v18: set<int> := {v8[982], v0};
				v2, v5, globalState.f4, v5 := v5[v0] in "ethjt", v5, |v18| >= (-v8[982] / v8[982]), v5 + v5;
				globalState.f4 := v2;
		}
		
		var v19 := new C8();
		var v20: array<bool> := new bool[22](i3 => v2);
		v20[998] := !v2;
		var v21: seq<bool> := [v2, v2, false, v2];
		var v22: multiset<bool> := multiset{v2};
		var v23: multiset<multiset<bool>> := multiset{(multiset(v21))[v2 := v0], v22, v22};
		v20[998] := (v23 > v23) <== fm49(fm49(true, globalState), globalState);
	}
	var v24: set<int> := {|v5|, v0};
	if (!(v24 >= v24)) {
		if (v2) {
			var v25: array<bool> := new bool[2] [v2, !v2];
			var v26: map<int, array<bool>> := map[v0 := v25];
			v26 := v26;
			var v27: C5 := new C5();
			var v28: array<C5> := new C5[5] [v27, v27, v27, v27, v27];
			v28[780] := v27;
			v28[780] := v27;
			var v29: array<char> := new char[26](i4 => 'l');
			var v30 := 'd';
			v29[346] := v30;
			v29[346] := 'h';
			var v31 := new C5();
			var v32 := new C9(v2);
		} else {
			var v33 := DC46();
			var v34: multiset<D23> := multiset{v33};
			var v35: map<bool, multiset<D23>> := map[true := v34];
			var v36: map<int, int> := map[|v35| := -v0];
			v36 := fm52(globalState);
			globalState.f3 := v0;
			var v37: array<bool> := new bool[6](i5 => v2);
			v37[438] := !v2;
			v37[438] := v0 >= v0;
			var v38: map<array<int>, array<bool>> := map[v8 := v37];
			v38 := v38;
			v37[438] := |v5| < v0;
		}
		
		var v39, v40, v41, v42 := m15(globalState);
		globalState.f4 := v41;
		var v43 := DC24(|fm23(v5, globalState)|, v41, v5);
		var v44: multiset<bool> := multiset{v41, v43.cf32, v5 < v5};
		v44 := v44[!v2 := |fm35(v41, v0, v2, v41, globalState)|] - v44;
		var v45, v46, v47, v48 := m15(globalState);
	} else {
		if (fm18(v2, v2, globalState) in "q") {
			globalState.f8 := v1[v0];
			globalState.f8 := v0 + v0;
			v8[919] := v0;
			v8[919] := -v0 % |[|v3|]|;
			globalState.f0 := v1;
			v8[919] := v0;
		} else {
			var v49 := new C2();
			var v50: C7 := new C7();
			var v51: map<C7, bool> := map[v50 := v50.fm3(v1, v7, v2, globalState)];
			var v52: map<bool, map<C7, bool>> := map[v2 := v51];
			globalState.f4 := (v51 != (if (false in v52) then v52[false] else v51)) <==> v2;
			var v53: C9 := new C9(v2);
			var v54: map<int, C9> := map[v0 := v53];
			var v55: T0 := new C5();
			var v56: map<T0, int> := map[v55 := v0];
			var v57 := DC55(v55);
			var v58: seq<bool> := [v2, v2];
			var v59: array<map<int, C9>> := new map<int, C9>[5] [v54, map[if (v57.cf77 in v56) then v56[v57.cf77] else |v58| := v53], map[v0 := v53], v54, v54];
			v59 := v59;
			v5 := v5;
			globalState.f8 := v0;
		}
		
		var v60 := new C8();
		var v61: map<bool, string> := map[v2 := v5];
		v60.m4(-v0 == v0, map[v2 := v5] + v61, 0x2c5, globalState);
		if (v2) {
			v8 := if (!v2) then v8 else v8;
			globalState.f4 := !v2;
			v60.m4(v2, v61, v0, globalState);
			var v62: seq<bool> := [v2];
			v62 := v62 + [!v2, false, v2];
			v8 := v8;
		} else {
			var v63: map<bool, bool> := map[v2 := false];
			var v64: array<map<bool, int>> := new map<bool, int>[12] [v3, (map[v2 := -0x2ae])[v2 := v0], (fm46(globalState))[if (!v2 in v63) then v63[!v2] else v2 := 0x17b], v3, v3, v3, v3, map[false := v0], map[v2 := v0], v3, v3, v3];
			var v65: T0 := new C4(v64);
			v65 := v65;
			var v66: array<seq<int>> := new seq<int>[6] [if (true) then v1 else v1, v1, v1, [v0, v0, v0], v1, seq(0x32d, i6  => (v0))];
			v66[459] := v1;
			v66[459] := v1 + v1;
			var v67: array<bool> := new bool[21](i7 => v3 != v3);
			v67[229] := |v5| <= v0;
			v67[229] := !v2;
			globalState.f4 := !(v67[229] <== (v0 <= |[if (v2 in v63) then v63[v2] else v67[229]]|));
			var v68: multiset<bool> := multiset{true};
			var v69: seq<multiset<bool>> := [v68];
			var v70: seq<seq<int>> := [v66[459][v0 := v0]];
			v1, v69, globalState.f3 := ((v1 + [v66[459][v0]]) + v70[v0])[v0 := v0], v69, v0 + v0;
		}
		
		if (if (v2) then v2 else v2) {
			var v71 := 'r';
			var v72 := DC24(if (-v0 in v7) then v7[-v0] else fm6(68, globalState), v2, (v5 + [v71])[|fm7(v71, globalState)| := v71]);
			v72 := v72.(cf31 := |[v0, v0]|, cf32 := v2).(cf32 := v2).(cf31 := v0);
			v8[657] := v0;
			v8[657] := v0 / v0;
			var v73: array<D2> := new D2[5];
			var v74: array<bool> := new bool[25];
			var v75: seq<array<bool>> := [v74];
			var v76 := DC5(v75[v8[657]]);
			v73[747] := v76;
			var v77: C1 := new C1();
			var v78: array<C1> := new C1[4] [if (v2) then v77 else v77, v77, v77, v77];
			v78[751] := v77;
			var v79: multiset<string> := multiset{seq(0x299, i8  => (v71)), v5};
			var v80: C7 := new C7();
			var v81: map<int, C7> := map[v0 := v80];
			v73[747], v78[751], v2, globalState.f4, globalState.f8 := v76, v77, v2, v2, (if (v5 in v79) then v79[v5] else v0) - (|v81| % v0);
			var v82: map<bool, C8> := map[v2 := v60];
			var v83 := DC2();
			var v84: map<map<bool, C8>, D0> := map[v82 + v82 := v83];
			globalState.f4, v2, v8[657], globalState.f4, globalState.f3 := v2, fm49(v2, globalState), v8[657] / (0x68 / v0), !v2, |v84|;
			var v85: multiset<bool> := multiset{fm49(v2, globalState)};
			var v86: seq<string> := [v5];
			var v87: map<int, int> := map[if (v2 in v85) then v85[v2] else v0 := -0xd4 * |v86[|"jjqbikwfu"|]|];
			globalState.f4, v87 := v2, (if (v2) then v87 else v87)[v0 := v0];
		} else {
			var v88 := DC36(v8);
			var v89: seq<array<int>> := [v8];
			var v90: multiset<bool> := multiset{v2};
			var v91: array<array<int>> := new array<int>[8] [if (v2) then v8 else v8, v88.cf55, v8, v8, v8, v8, v8, v89[|v90|]];
			v91[630] := v8;
			v91[630] := new int[25];
			var v93: map<set<int>, bool> := map[{v0, v0, -|(set v92 : int | v92 in v24 :: (v92 + |map[false := [-0x342]]|))|, 0x25d} + v24 := v2];
			v93 := v93 + v93;
			var v94: seq<bool> := [v2, v2];
			globalState.f8 := v0 - (if (v94[v0] in v90) then v90[v94[v0]] else v0);
			var v95 := new C9(false && v2);
			var v96: C8 := new C8();
			v96 := v96;
		}
		
	}
	
	var v97 := new C0();
	var v98: map<int, int> := map[v0 := v0];
	var v99 := DC8(v98);
	var v100: seq<bool> := [v2];
	var v101: set<bool> := {v2};
	var i9 := 0;
	while (fm49(fm35(v2, fm26(v99, v0, multiset{v100, v100, v100}, globalState), v2, true, globalState) >= v101, globalState))
		decreases 100 - i9
	{
		if (i9 >= 100) {
			break;
		}
		
		i9 := i9 + 1;
		var v102: array<map<bool, int>> := new map<bool, int>[23];
		var v103: C4 := new C4(v102);
		var v104: seq<C4> := [v103];
		v103 := v104[208];
		globalState.f8 := if (!v2) then -0x109 / v0 else v0 + v0;
		var v105: T0 := new C7();
		var v106 := 'w';
		var v107: set<map<T0, char>> := {map[v105 := v106]};
		var v108: map<T0, char> := map[v105 := 'q'];
		v107 := {v108, v108, map[v105 := 'y'], map[v105 := v106]};
		v103.m9(fm6(v0, globalState), v0, globalState);
	}
	var v109 := DC9();
	v109 := v109;
	var v110 := DC43(seq(0x2a6, i11  => (|map[false := v2]|)), false);
	var i10 := 0;
	while (match v110.(cf64 := fm49(v2, globalState)) {
		case DC43(cf63, cf64) => v0 >= v0
		case DC42(cf62) => v2
	})
		decreases 100 - i10
	{
		if (i10 >= 100) {
			break;
		}
		
		i10 := i10 + 1;
		if (v2) {
			var v111: array<map<bool, int>> := new map<bool, int>[27];
			var v112 := new C4(v111);
			globalState.f8 := if (v2) then v0 else -v0;
			var v113: map<int, bool> := map[v0 := v112.fm3(v1, multiset{v0}, !v2, globalState)];
			var v114: set<map<int, bool>> := {v113, v113, v113, v113 + v113, v113};
			var v115: T0 := new C2();
			var v116: seq<T0> := [v115, v115];
			var v117 := 'y';
			var v118: C1 := new C1();
			var v119: map<C1, int> := map[v118 := 0x6e];
			v2, v114, v115, v5 := v5 <= fm25(globalState), {v113} * fm53(v0, v0, globalState), v116[v0], (("ukahennov")[v0 := v117] + ("uvcx" + v5))[if (v118 in v119) then v119[v118] else -v0 := v117];
			var v120: map<string, bool> := map["gofjuqlqk" := v2];
			globalState.f3 := |v120|;
			v0 := -0x100;
		} else {
			var v121: array<bool> := new bool[10];
			v121[429] := v2;
			v121[429] := v2;
			var v122, v123, v124, v125 := m15(globalState);
			var v126: map<bool, bool> := map[v124 := v121[429]];
			globalState.f4 := v126 != v126;
			var v127, v128, v129, v130 := m15(globalState);
			v8[49] := v122;
			v8[49] := v127;
		}
		
		globalState.f8, globalState.f8 := -0x109, -v0;
		var v131: T0 := new C7();
		var v132: seq<T0> := [v131];
		var v133: map<bool, bool> := map[v2 := v2];
		var v134: multiset<bool> := multiset{v2};
		var v135: array<bool> := new bool[29] [v24 !! v24, v100 == [v2, v2, v2, v2, v2], v2, if (v2 in v133) then v133[v2] else v2, v2, v2, v100 == v100, false, v2, true, if (v2) then v2 else v2, v2, v2, v2, v131.fm3([|v5|], v7, v2, globalState), v2, v2, v2, v2, v2, v2, v2, v2 <== v2, v2, v2, false, multiset{v2} !! v134, v2, v2];
		v135[299] := !(v2 ==> v2);
		var v137: array<map<char, bool>> := new map<char, bool>[4](i12 => map v136 : char | v136 in multiset{'w'} :: (v136) := (v2));
		var v138 := 'q';
		var v139: map<char, bool> := map[v138 := v2];
		v137[625] := v139[v138 := v2];
		var v140 := DC58(v139);
		globalState.f8, v132, v2, v135[299], v137[625] := v0, [if (v2) then v131 else v131, v131, v131, v131, v131], false, v131.fm3(v1, v7, v2, globalState), v140.cf82;
		var v141 := DC5(v135);
		match (v141) {
			case DC6(cf6, cf7, cf8) =>
				v3 := v3[cf7 := cf6];
				globalState.f8 := v0 / v0;
				var v142: array<array<bool>> := new array<bool>[25];
				v142 := v142;
				v8[851] := v0;
				v8[851] := cf6;
			case DC7(cf9, cf10, cf11) =>
				cf10 := v135[299];
				globalState.f4 := v135[299];
				v133 := v133[cf10 <== v135[299] := v0 >= |"dmy"|];
				globalState.f8 := -(if (cf10 in v3) then v3[cf10] else v0);
			case DC5(cf5) =>
				v135[299] := if (v135[299]) then v135[299] else v135[299];
				v8 := v8;
				var v143 := DC59(v2, v135, v135[299]);
				globalState.f4 := v143.cf83;
				globalState.f3 := v0 / v0;
		}
		
	}
	var v144: array<bool> := new bool[12](i13 => v2);
	v144[300] := v2;
	v144[300] := v2;
	for i14 := v0 to if (v2) then v0 else |(seq(0xfa, i15  => ('y')))| {
		var v145: seq<array<bool>> := [v144];
		globalState.f3 := |[v145[v0], v144]|;
		var v146 := 'y';
		v146 := 'l';
		v7 := v7;
		if (!false !in fm23(v5[v0 := v146], globalState)) {
			var v147: array<array<int>> := new array<int>[23];
			v147[431] := v8;
			v147[431] := v8;
			var v148: C5 := new C5();
			var v149: array<C5> := new C5[1] [v148];
			v149[995] := if (v144[300]) then v148 else v148;
			v149[995] := v148;
			globalState.f4 := v144[300];
			var v150: multiset<bool> := multiset{v2};
			globalState.f3, v144[300], globalState.f3, v100, globalState.f3 := if (!(i14 < i14)) then v0 else i14, v150[v144[300] := i14] > (multiset(v100) - v150), v0 - v0, (v100 + v100) + v100, |v100[-v0 := v2]|;
			v144[300] := v150 !! v150;
		} else {
			var v151: C8 := new C8();
			var v152: seq<C8> := [v151, v151];
			v151 := v152[i14];
			globalState.f3 := v0;
			v144[300] := fm49(v2, globalState);
			var v153: C3 := new C3();
			v153 := v153;
			globalState.f8 := v0;
		}
		
	}
	var v154: C8 := new C8();
	var v155: array<C8> := new C8[3] [v154, v154, v154];
	var v156: array<array<C8>> := new array<C8>[17] [v155, v155, v155, v155, v155, v155, v155, v155, v155, v155, v155, v155, v155, v155, v155, v155, v155];
	v156[845] := v155;
	var v157: map<bool, seq<bool>> := map[v144[300] := v100];
	var v158: multiset<seq<bool>> := multiset{if (v2 in v157) then v157[v2] else v100};
	globalState.f4, v5, globalState.f3, v156[845] := false, "kmhjp", fm26(v99, |v100|, v158, globalState), v155;
	globalState.f4 := !v2;
	globalState.f3 := v0 * v0;
	globalState.f4 := true;
	var v159 := DC6(v0, v144[300], v144[300]);
	for i16 := |v100| to v0 - v159.cf6 {
		if (v2) {
			var v160 := 'x';
			var v161: map<char, string> := map[v160 := v5];
			var v162: map<string, bool> := map[if (v160 in v161) then v161[v160] else seq(-0x30c, i17  => (DC35(v0, v5[v0], |{v2, v2}|, i16, v0).cf51)) := v2];
			var v163: seq<set<int>> := [{v0}];
			v162 := fm42(i16, 0xc6, |v163[i16]|, true, globalState);
			var v164: set<map<int, int>> := {v98, v98 + v98};
			var v165: map<bool, set<map<int, int>>> := map[v144[300] := {v98}];
			v164 := if ((v2 ==> true) in v165) then v165[v2 ==> true] else v164;
			globalState.f3 := 0x167;
			v144[300] := v144[300];
			var v166: multiset<string> := multiset{v5, v5};
			globalState.f3 := |v166| % i16;
		} else {
			globalState.f4 := v144[300];
			v3 := v3[v2 := |v6|];
			globalState.f4 := v2;
			var v167 := DC30(v144[300], v2, v24, v97);
			globalState.f4 := !v2 ==> v167.cf43;
			var v168: map<bool, string> := map[v2 := v5];
			v154.m4(v2, v168, 0xef, globalState);
		}
		
		globalState.f4 := true;
		globalState.f4 := true;
		v8 := v8;
	}
	var v169 := new C0();
	var v170: multiset<string> := multiset{v5};
	globalState.f4 := (v170 + v170) > v170;
}