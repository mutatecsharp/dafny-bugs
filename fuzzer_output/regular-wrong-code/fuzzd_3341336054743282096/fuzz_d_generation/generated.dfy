datatype D0 = DC0(cf0: array<set<bool>>, cf1: int) | DC1(cf2: int)
datatype D1 = DC2(cf3: bool)
datatype D2 = DC4(cf5: int, cf6: bool, cf7: array<array<int>>) | DC5(cf8: int, cf9: seq<int>, cf10: bool) | DC3(cf4: set<seq<char>>) | DC6(cf11: D2)
datatype D3 = DC7(cf12: map<char, int>)
datatype D4 = DC9(cf14: multiset<int>, cf15: char, cf16: int, cf17: set<C0>, cf18: int) | DC10(cf19: int, cf20: int, cf21: int) | DC11(cf22: int, cf23: bool, cf24: bool, cf25: bool, cf26: bool) | DC8(cf13: seq<D0>) | DC12(cf27: D4)
datatype D5 = DC13(cf28: array<D3>)
datatype D6 = DC15(cf30: int, cf31: int, cf32: bool) | DC14(cf29: string) | DC16(cf33: D6)
datatype D7 = DC18 | DC19(cf35: int) | DC17(cf34: seq<bool>)
datatype D8 = DC20(cf36: set<array<bool>>)
datatype D9 = DC22(cf38: int) | DC21(cf37: multiset<bool>)
datatype D10 = DC23(cf39: array<int>)
datatype D11 = DC25(cf41: char, cf42: int, cf43: bool) | DC24(cf40: map<string, char>) | DC26(cf44: D11)
datatype D12 = DC28(cf46: bool, cf47: int) | DC29(cf48: bool, cf49: int, cf50: bool, cf51: int, cf52: bool) | DC27(cf45: set<bool>)
datatype D13 = DC31(cf54: bool, cf55: int, cf56: int) | DC32(cf57: string, cf58: bool, cf59: int, cf60: array<int>) | DC30(cf53: array<seq<int>>) | DC33(cf61: D13)
datatype D14 = DC35(cf63: int, cf64: string, cf65: int, cf66: bool, cf67: bool) | DC36(cf68: map<int, int>, cf69: map<bool, int>) | DC34(cf62: map<int, string>)
datatype D15 = DC38(cf71: char) | DC39(cf72: string, cf73: D9) | DC37(cf70: set<seq<int>>) | DC40(cf74: D15)
datatype D16 = DC42(cf76: int, cf77: int, cf78: int) | DC43(cf79: string, cf80: char, cf81: int) | DC44(cf82: int) | DC41(cf75: array<bool>)
datatype D17 = DC46(cf84: int) | DC45(cf83: multiset<multiset<int>>)
datatype D18 = DC48 | DC49(cf86: D8) | DC50(cf87: bool, cf88: array<bool>, cf89: bool, cf90: bool, cf91: string) | DC47(cf85: map<map<bool, int>, string>)
datatype D19 = DC51(cf92: C5)
datatype D20 = DC53(cf94: bool) | DC54(cf95: string, cf96: int, cf97: bool) | DC52(cf93: array<char>)
datatype D21 = DC55(cf98: T0)
datatype D22 = DC56(cf99: C4)
datatype D23 = DC58 | DC57(cf100: set<set<int>>)
datatype D24 = DC60 | DC61(cf102: int, cf103: multiset<array<int>>) | DC62(cf104: int) | DC59(cf101: array<D16>)
datatype D25 = DC64(cf106: multiset<map<T1, int>>) | DC65(cf107: int, cf108: array<bool>, cf109: int) | DC63(cf105: map<int, multiset<bool>>) | DC66(cf110: D25)
datatype D26 = DC68(cf112: bool) | DC69(cf113: string, cf114: bool) | DC67(cf111: multiset<set<bool>>)
datatype D27 = DC70(cf115: multiset<string>)
datatype D28 = DC72(cf117: int, cf118: int, cf119: bool, cf120: bool) | DC71(cf116: map<int, bool>)
class GlobalState {
	constructor () {
	}
	
}

function fm1(p0: int, p1: bool, p2: int, p3: bool, globalState: GlobalState): map<bool, bool> {
	map[false := true] + (map[DC29(true, 0x26c, false, |[false]|, false).cf52 := !true] + map[false := true])
}
function fm2(p0: bool, p1: int, globalState: GlobalState): bool {
	DC22(|(map v0 : string | v0 in map[seq(0xc0, i0  => ('d')) := -0x5b] :: (v0) := (0x361))|).cf38 <= 103
}
function fm3(p0: char, p1: int, p2: bool, p3: bool, globalState: GlobalState): map<seq<bool>, char> {
	(map[[!true] := 'y'] + (map v0 : seq<bool> | v0 in (seq(0x37f, i0  => ([false]))) :: (v0) := ('e'))) + map[[true] := 'm']
}
function fm4(p0: string, p1: char, p2: bool, globalState: GlobalState): map<int, bool> {
	map[271 := "s" != (seq(-0xfe, i0  => ('o')))]
}
function fm10(p0: bool, globalState: GlobalState): set<int> {
	{810} * (if (true) then set v0 : int | (0xc0 <= v0) && (v0 < 0x284) :: (v0 - -0x10d) else {|[false]|, -526})
}
function fm13(p0: char, p1: bool, globalState: GlobalState): multiset<int> {
	multiset{|[|map[[|[true, true]|] := |(seq(0x2cc, i0  => ('d')))|]|]| % |map[|(seq(0x26, i1  => ('b')))| := |(seq(0x388, i2  => (|map[true := true]|)))|]|}
}
function fm16(p0: string, p1: int, globalState: GlobalState): string {
	"l"
}
function fm17(globalState: GlobalState): bool {
	!!!(if (true) then (seq(-780, i0  => ([true]))) < [[false, false]] else {false} > {false})
}
function fm18(globalState: GlobalState): multiset<set<bool>> {
	multiset{{!!true}} * (multiset{{true}, {false, true}} - DC67(multiset{{false, !false, true}, {true, false}}).cf111)
}
function fm19(p0: bool, p1: bool, p2: map<char, int>, p3: bool, globalState: GlobalState): seq<bool> {
	[false, true, false, false, !false] + ([true, false, false, false, false] + [false, false])
}
function fm20(p0: bool, globalState: GlobalState): int {
	|(map[!true := 0x123] + map[false := |[multiset{true}, multiset{true}, multiset{!true, true}]|])| + -|(multiset{"qeutyx", "dpui"} - DC70(multiset{"ts", "jkwix", seq(-0x379, i0  => ('k'))}).cf115)|
}
function fm21(globalState: GlobalState): map<int, bool> {
	if (false <==> true) then map[817 := true] + map[0x2ea := false] else map[494 := true] + map[|"lno"| := true]
}
function fm22(p0: bool, p1: int, p2: int, globalState: GlobalState): map<int, int> {
	map[|(seq(0x38e, i0  => ('h')))| / 0x389 := -0x2bf]
}
function fm23(p0: bool, globalState: GlobalState): seq<bool> {
	if (|"hlbbvk"| != -|map[|[true]| := false]|) then [false, !true, false, !!true, true] + [!true] else if (true) then [true, !false, false] else [true]
}
function fm24(p0: bool, globalState: GlobalState): multiset<int> {
	multiset{-72, -0x149}
}
function fm25(p0: D4, p1: string, globalState: GlobalState): string {
	("vy" + "y") + ("mbgtm" + "pacqqiw")
}
function fm26(p0: set<bool>, p1: seq<bool>, p2: D0, p3: char, globalState: GlobalState): map<string, char> {
	DC24(map["le" := 's']).cf40
}
function fm29(p0: bool, globalState: GlobalState): char {
	if (false) then 'm' else 't'
}
function fm30(p0: int, p1: bool, p2: bool, globalState: GlobalState): bool {
	true
}
function fm31(p0: int, p1: bool, p2: int, globalState: GlobalState): string {
	"roi"
}
function fm32(p0: int, globalState: GlobalState): int {
	|(map v0 : int | v0 in map[|(map v1 : int | (89 <= v1) && (v1 < 359) :: (v1 - 952) := (false))| := |[false]|] :: (v0 * -|"rqy"|) := (0x146))|
}
function fm33(p0: int, p1: seq<char>, globalState: GlobalState): D6 {
	DC14("f" + "csnarsqlx")
}
function fm34(p0: bool, p1: bool, p2: bool, p3: int, globalState: GlobalState): D12 {
	if (!(|{|map[true := 'i']|}| > |[0x226, 706]|)) then DC28(true, -0x18b) else DC28(!false, |[true]|)
}
function fm35(p0: bool, p1: bool, globalState: GlobalState): seq<int> {
	[0x145]
}
function fm36(p0: char, p1: int, p2: int, p3: set<bool>, globalState: GlobalState): map<char, int> {
	(map['q' := -0x41] + map['i' := |"xaxw"|]) + map['d' := |map[{!false} := |"utuis"|]|]
}
function fm37(p0: bool, globalState: GlobalState): D7 {
	DC19(236)
}
function fm38(p0: bool, p1: int, p2: bool, globalState: GlobalState): seq<bool> {
	([!false] + [true, true]) + [true, false]
}
function fm39(p0: bool, globalState: GlobalState): map<int, bool> {
	((map v0 : int | (0x378 <= v0) && (v0 < 0x30b) :: (v0 % 0x30a) := (!!false)) + DC71(map v1 : int | v1 in [0x178] :: (v1 - |[true]|) := (false)).cf116) + (map[0x374 := false] + map[-|multiset{true, false}| := false])
}
function fm40(p0: bool, p1: char, p2: bool, p3: D7, globalState: GlobalState): D12 {
	if (true) then DC27({true}) else DC27({false, true})
}
function fm43(p0: D11, p1: bool, p2: int, globalState: GlobalState): map<int, string> {
	map[|map[997 := !false]| + -|"vwotwnwen"| := "amno"]
}
function fm44(p0: bool, globalState: GlobalState): int {
	0x35e
}
function fm45(p0: bool, p1: bool, p2: int, globalState: GlobalState): D14 {
	DC35(656 * |[|[true, true, true]|]|, (seq(65, i0  => ('n'))) + (seq(0x248, i1  => ('j'))), |multiset{true, !true}|, multiset{!true} >= multiset{!true}, |[false, true]| >= |"edklwnr"|)
}
function fm46(p0: bool, p1: int, p2: bool, globalState: GlobalState): string {
	("x" + (seq(0x1ed, i0  => ('m')))) + "sdfbv"
}
function fm47(p0: int, p1: bool, globalState: GlobalState): map<bool, bool> {
	if (-0x315 in [|multiset{'i'}|]) then map[false := !true] + map[true := false] else map[!false := false]
}
function fm48(p0: D7, p1: int, p2: int, p3: int, globalState: GlobalState): multiset<int> {
	multiset(seq(0x18b, i0  => (|map[[!false, false, false] := -0x17f]|))) + (multiset{0x331} * multiset{|(map v0 : int | (0x3be <= v0) && (v0 < 0x331) :: (v0 / 0x2c9) := (0x357))|, |"kjajll"|})
}
function fm49(p0: char, globalState: GlobalState): D0 {
	DC1(if (false) then 814 else |[multiset{false}, multiset{!false}]|)
}
function fm50(p0: bool, p1: bool, p2: int, globalState: GlobalState): map<int, bool> {
	map v0 : int | (0x3ca <= v0) && (v0 < 0x159) :: (v0 % 0xa9) := (!false)
}
function fm51(p0: bool, p1: int, p2: int, globalState: GlobalState): char {
	'k'
}
function fm52(p0: bool, p1: D15, p2: bool, globalState: GlobalState): map<char, int> {
	if (if (false) then true else true) then map['d' := -|(map v0 : char | v0 in map['j' := DC53(false).cf94] :: (v0) := (935))|] else map['s' := |[0x38e]|]
}
function fm53(p0: map<bool, int>, globalState: GlobalState): set<bool> {
	{|map[{|"souhprl"|, |map[|map[false := |"ywax"|]| := false]|, 426, |"liy"|, -|[|[false, false]|, -0x20d]|} := false]| >= 981, !!false <==> true, !(true <== false), true}
}
function fm54(p0: bool, p1: seq<bool>, globalState: GlobalState): D3 {
	match DC25('t', 0x4b, true) {
		case DC25(cf41, cf42, cf43) => if (cf43) then DC7(map[cf41 := cf42]) else DC7(map v0 : char | v0 in [cf41] :: (v0) := (cf42))
		case DC24(cf40) => if (false) then DC7(map['t' := |multiset{false, true}|]) else DC7(map['q' := 997])
		case DC26(cf44) => DC7(map['n' := |map[true := !true]|])
	}
}
function fm55(p0: int, globalState: GlobalState): set<seq<bool>> {
	{[false], [false], [false, false, true, true], [true, false, true], [!true, true, false, false, true]} - ({[false], [!false, false]} + {[false, false, false]})
}
function fm56(p0: bool, globalState: GlobalState): map<seq<bool>, int> {
	((map v0 : seq<bool> | v0 in map[[!false] := set v1 : char | v1 in ['e', 'f', 'o', 'y', 'd'] :: (v1)] :: (v0) := (421)) + (map v2 : seq<bool> | v2 in {[!false]} :: (v2) := (0x2f1))) + (map v3 : seq<bool> | v3 in {[false, false]} :: (v3) := (-0x33d))
}
function fm57(p0: bool, p1: bool, p2: bool, globalState: GlobalState): map<bool, char> {
	map[!false := 'g'] + map[true := 'u']
}
function fm58(p0: D16, p1: bool, p2: set<seq<int>>, globalState: GlobalState): multiset<char> {
	multiset{'o'}
}
function fm59(globalState: GlobalState): set<string> {
	{"wsdwwpuv"}
}
function fm60(p0: int, p1: bool, globalState: GlobalState): map<D11, int> {
	map v0 : D11 | v0 in (if (!true) then {DC25('o', -0x34f, true), DC25('d', -0x31e, true)} else {DC25('h', -|map[true := |(seq(0x1f4, i0  => (|[true]|)))|]|, true)}) :: (v0) := (-489)
}
function fm61(p0: int, p1: int, globalState: GlobalState): D9 {
	if (!true) then DC21(multiset{false}) else DC21(multiset{true, true})
}
function fm62(p0: seq<map<char, int>>, p1: bool, p2: int, p3: int, globalState: GlobalState): D15 {
	DC39("liockxaf", DC21(multiset([false])))
}
function fm63(p0: map<bool, int>, p1: int, globalState: GlobalState): multiset<bool> {
	(multiset{true} * multiset{!true, true}) * (multiset{true, !!false} + multiset{true})
}
function fm64(globalState: GlobalState): D4 {
	DC10(0x1f3 % 0x298, -|[845, |multiset{|map[|[-0x358]| := true]|}|]| - 575, |("ga" + "bywhwnw")|)
}
function fm65(p0: int, p1: char, p2: int, globalState: GlobalState): map<bool, int> {
	map[true := 0x2c0] + (map[true := |[true, true, true]|] + map[true := |"mdjiex"|])
}
function fm66(p0: int, p1: bool, globalState: GlobalState): D16 {
	DC42(149 % |{-0x10}|, --|([true, false] + [true, false])|, 0x13e % -0x342)
}
function fm67(p0: set<map<bool, bool>>, globalState: GlobalState): D11 {
	DC25('a', -|((seq(0x2b3, i0  => (0x7f))) + [|{false}|, |{true}|, |"vhdvmehq"|])|, true)
}
function fm68(p0: bool, p1: string, globalState: GlobalState): D14 {
	DC34(map[-61 := "vj"])
}
function fm69(p0: bool, globalState: GlobalState): map<int, D15> {
	map[-(if (!DC54("irhb", |[!true]|, true).cf97) then -|"wfcecnt"| else |[|"cicqjbm"|]|) := DC37({seq(0x1e4, i0  => (0x139)), [-0x322]})]
}
function fm70(p0: bool, globalState: GlobalState): set<map<char, int>> {
	match DC43("eqqftomtx", 'v', -|(map v0 : int | (0x3d3 <= v0) && (v0 < 0x345) :: (v0 - -0xf3) := (|map[true := |map[true := 'g']|]|))|) {
		case DC42(cf76, cf77, cf78) => set v2 : map<char, int> | v2 in [map v1 : char | v1 in (seq(0x30e, i0  => ('d'))) :: (v1) := (|(seq(-0x171, i1  => ("ttkmplyt")))|)] :: (v2)
		case DC43(cf79, cf80, cf81) => {map[cf80 := |cf79|]} * {map[cf80 := cf81]}
		case DC44(cf82) => (set v4 : map<char, int> | v4 in {map['d' := cf82], map v3 : char | v3 in multiset{'u'} :: (v3) := (cf82)} :: (v4)) * (set v5 : map<char, int> | v5 in {map['a' := 815]} :: (v5))
		case DC41(cf75) => {map['k' := |"imekcx"|], map['h' := -0xf6], map v6 : char | v6 in map['p' := false] :: (v6) := (--0x129)} + {map['d' := 85], map v7 : char | v7 in ['q'] :: (v7) := (|map[false := |(seq(-771, i2  => ('d')))|]|)}
	}
}
function fm71(p0: bool, globalState: GlobalState): seq<seq<bool>> {
	[[false, !false], [!true], [true] + [true], [false, true, true] + [!true]]
}
function fm72(globalState: GlobalState): seq<string> {
	[(seq(0x18b, i0  => ('q'))) + "usta"]
}
trait T0 {
	function fm5(p0: int, p1: bool, p2: map<char, D0>, p3: bool, globalState: GlobalState): bool 
	method m1(p0: int, globalState: GlobalState) returns (r0: bool, r1: bool) 
	method m2(p0: int, globalState: GlobalState) returns (r0: map<bool, multiset<int>>) 
}

trait T1 {
	const f0 : seq<int>
	var f1 : D0
	function fm6(p0: int, p1: int, globalState: GlobalState): int 
}

class C0 {
	var f8 : string
	constructor (f8 : string) {
		this.f8 := f8;
	}
	
	function fm14(globalState: GlobalState): char {
		'w'
	}
	function fm15(p0: seq<int>, p1: bool, globalState: GlobalState): string {
		(seq(0x5, i0  => ('u'))) + f8
	}
}

class C1 extends T1 {
	var f9 : int
	const f10 : array<map<char, int>>
	constructor (f9 : int, f10 : array<map<char, int>>, f0 : seq<int>, f1 : D0) {
		this.f9 := f9;
		this.f10 := f10;
		this.f0 := f0;
		this.f1 := f1;
	}
	
	function fm6(p0: int, p1: int, globalState: GlobalState): int {
		|multiset{0x25b / f9, f9, f9, |((set v0 : int | v0 in multiset{|"drwqyvqsw"|} :: (v0 + 0x265)) - {-f9, f9, |multiset([true])|})|, |((seq(0x31d, i0  => ('o'))) + ['h', 'c'])|}|
	}
	function fm27(p0: map<int, seq<bool>>, p1: int, p2: bool, p3: int, globalState: GlobalState): multiset<bool> {
		multiset(([!!false] + [true]) + ([false] + [false]))
	}
	function fm28(p0: int, p1: bool, globalState: GlobalState): int {
		f9
	}
	method m13(p0: char, p1: bool, globalState: GlobalState) returns (r0: bool) {
		f9 := f9;
		r0 := p1;
		if (p1) {
			var v0 := "dmxvlol";
			var v1 := new C0(v0);
			var v2 := new C0(v0 + v1.f8);
			r0, r0, f9 := true, p1, (f9 - f9) / (f9 - f9);
			var v3: multiset<bool> := multiset{p1, p1};
			var v4: seq<bool> := [p1];
			var v5 := DC17(v4);
			var v6: multiset<bool> := multiset{v0 == v2.f8, p1, p1, !(v3 > multiset(v5.cf34))};
			v3 := (multiset{p1} + multiset([p1, p1])) - v3;
			var v7 := new C0(v0 + v1.f8);
		} else {
			var v8: array<bool> := new bool[29];
			v8[517] := p1;
			v8[517] := p1;
			var v9: array<int> := new int[8];
			v9[609] := f9;
			v9[609] := f9 % f9;
			v8 := v8;
			var v10: array<char> := new char[2](i0 => DC25(p0, |multiset{DC14("pqjlcv"), DC14("lcnl")}|, p1).cf41);
			v10[801] := fm29(v8[517], globalState);
			v10[801], r0 := p0, v8[517];
			v9[609] := -((-65 + f9) / (v9[609] * v9[609]));
		}
		
		var v11: array<char> := new char[23](i2 => p0);
		var v12: seq<array<char>> := [v11, v11, v11];
		for i1 := f9 to |v12| {
			var v13: array<array<int>> := new array<int>[13];
			var v14 := DC4(f9, p1, v13);
			var v15: multiset<int> := multiset{f9 * v14.cf5};
			v15 := v15;
			var v16: array<bool> := new bool[14](i3 => p1 && p1);
			v16[887] := p1 && p1;
			v16[887] := p1;
			var v17: array<D6> := new D6[17];
			var v18: seq<bool> := [p1];
			var v19 := DC15(i1, f9, fm30(|v18|, v16[887], v16[887], globalState));
			v17[194] := v19;
			var v20: set<int> := {f9};
			var v21: array<set<int>> := new set<int>[4] [{f9}, v20 - v20, v20, v20];
			var v22: array<int> := new int[6](i4 => i4 / i1);
			v22[256] := f9 % f9;
			v17[194], v21, v16[887], v22[256] := v19, v21, v16[887], if (p1) then -i1 else f9 - f9;
			var v23 := "rarjd";
			var v24 := DC14(v23 + "jmypydeh");
			v24, v18, v22[256] := DC14(v23).(cf29 := v23), [v16[887], p1, p1], f9;
		}
		for i5 := f9 to f9 {
			var v25: array<seq<bool>> := new seq<bool>[3];
			var v26: seq<bool> := [p1, p1];
			v25[908] := v26 + v26[f9 := p1];
			var v27 := DC17(v26);
			var v28: map<bool, seq<bool>> := map[p1 := v26];
			v25[908] := (if (p1) then v27 else DC17(if (p1 in v28) then v28[p1] else v26)).cf34[f9 := DC29(p1, i5, p1, 0x381, p1).cf50];
			var v29 := "nyg";
			v29 := v29 + v29;
			f9 := f9 / f9;
			var v30: map<int, bool> := map[|v29| := p1];
			var v31: array<bool> := new bool[12];
			var v32: multiset<array<bool>> := multiset{v31, v31};
			var v33: multiset<char> := multiset{p0};
			var v34: multiset<int> := multiset{i5, i5, |v33|};
			var v35: set<multiset<int>> := {multiset([i5]), v34, multiset(f0), multiset{|fm31(f9, p1, f9, globalState)|}};
			var v36: array<int> := new int[22] [-i5 - i5, |v30|, f9 + i5, i5, -i5 / i5, f9 / i5, i5, f9 * f9, |f0|, -fm28(-0x50, p1, globalState), if (v31 in v32) then v32[v31] else |v35|, f9, i5, i5, f9, |f0|, i5, i5 * -|fm31(f9, p1, -|v29|, globalState)|, f9 + 0x159, i5, f9, f9];
			v36[48] := 0x372;
			v36[48] := if (v31 in v32) then v32[v31] else f9;
		}
		var v37 := DC29(p1, fm28(f9, !!p1, globalState), p1, f9, !p1);
		var v38: map<int, int> := map[v37.cf49 := f9];
		var v39: map<int, bool> := map[f9 := p1];
		var v40 := DC25(p0, 390, if (f9 in v39) then v39[f9] else p1);
		var v41 := "bhh";
		for i6 := -(if (v40.cf42 in v38) then v38[v40.cf42] else |multiset{p0}|) to |v41| {
			r0 := if (p1) then !p1 else p1;
			var v42: multiset<int> := multiset{i6, i6, fm6(f9, f9, globalState), i6, f9};
			f9 := (|v42| + i6) % (fm6(f9, f9, globalState) + f9);
			var v43 := 'q';
			v43 := v43;
			var v44: array<string> := new string[25](i7 => "htrgdatxk");
			v44[115] := v41;
			v44[115] := if (p1) then v41 else v41;
		}
		r0 := if (v41 <= "ly") then p0 !in fm31(f9, p1, |multiset{p0, p0, p0}|, globalState) else p1;
	}
}

class C2 extends T0 {
	constructor () {
	}
	
	function fm5(p0: int, p1: bool, p2: map<char, D0>, p3: bool, globalState: GlobalState): bool {
		false
	}
	function fm41(p0: bool, p1: int, globalState: GlobalState): multiset<bool> {
		multiset{false}
	}
	function fm42(p0: bool, globalState: GlobalState): bool {
		|(map v0 : int | (-0x22e <= v0) && (v0 < 0x230) :: (v0 / |[false]|) := ([DC5(-0x340, seq(-0x65, i0  => (|map[!false := 'p']|)), false), DC5(|map[|"tam"| := 'k']|, [|[false, !false, !true]|, |{true}|], true)]))| < 0x119
	}
	method m1(p0: int, globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0: multiset<int> := multiset{p0, p0};
		var v1: map<int, int> := map[|v0| := p0];
		var v2: array<map<char, int>> := new map<char, int>[1];
		var v3: seq<int> := [|"kpmro"|];
		var v4: T1 := new C1(p0, v2, v3, DC1(p0));
		var v5: map<T1, int> := map[v4 := p0];
		var v6: C0 := new C0("us");
		var v7: map<int, C0> := map[p0 := v6];
		var v8: seq<map<int, int>> := [v1, map[p0 := p0], map[|v5| := |v7|]];
		var v9 := 0x9c;
		v8, v9, v6.f8 := v8 + v8, |(seq(-425, i0  => (v9)))|, "fvuiwig";
		var v10: array<bool> := new bool[17](i2 => false in map[!false := multiset{false}]);
		forall i1 | 0 <= i1 < v10.Length {
			v10[i1] := v9 == v4.f0[v9];
		}
		var v11 := true;
		var v12: map<bool, int> := map[v11 := v9];
		var v13: seq<string> := [v6.f8, v6.f8, v6.f8];
		var v14: map<map<bool, int>, string> := map[v12 := v13[0x167]];
		var v16: seq<map<bool, int>> := [v12, v12];
		var v17 := m14(v11, v11, v14 + (map v15 : map<bool, int> | v15 in v16 :: (v15) := (v6.f8)), p0, globalState);
		r0 := v11;
		var v18 := 'x';
		var i3 := 0;
		while ({p0} >= (set v19 : int | v19 in map[v9 := v18] :: (v19 % --404)))
			decreases 100 - i3
		{
			if (i3 >= 100) {
				break;
			}
			
			i3 := i3 + 1;
			var v20 := DC15(v4.fm6(v17, v17, globalState), p0, v11);
			var v21 := DC16(v20);
			var v22 := DC16(v20);
			v22 := v22;
			r1 := v11;
			var v24: map<string, char> := map[v6.f8 := v18];
			var v25 := DC24(v24);
			var v26 := DC26(v25);
			var v27: seq<bool> := [v11, v11];
			var v28 := DC34(fm43(v26, v11, |v27|, globalState));
			v11 := ((map v23 : int | (0x1ba <= v23) && (v23 < -0x3a0) :: (v23 + v17) := ("qkbtfupqy")) + map[p0 := v6.f8]) != v28.cf62;
			v10[653] := v11;
			v10[653] := v11;
		}
		r1 := v11;
		r0 := v11;
		r1 := v11;
	}
	method m2(p0: int, globalState: GlobalState) returns (r0: map<bool, multiset<int>>) {
		var v0 := "jgwejo";
		v0 := v0 + v0;
		var i0 := 0;
		while (true)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v1: array<string> := new string[4];
			v1[932] := seq(0x36, i1  => ('i'));
			v1[932] := "pekucbt";
			var v2 := true;
			var v3 := DC1(p0);
			var v4: seq<bool> := [v2];
			var v5: map<D0, int> := map[v3 := |v4|];
			var v6: seq<map<D0, int>> := [v5];
			var v7: seq<int> := [p0, p0, p0];
			var v8: map<int, bool> := map[|v7| := v2];
			var v9: map<int, bool> := map[-|multiset(v6)| := (if (0x1b0 in v8) then v8[0x1b0] else v2) && v2];
			var v10: set<int> := {p0};
			v2 := if (|(map[|v10| := false] + v8)| in v8) then v8[|(map[|v10| := false] + v8)|] else v2 <==> v2;
			var v11: array<bool> := new bool[2];
			var v12: set<array<bool>> := {v11};
			var v13 := DC20(v12);
			var v14: map<bool, D8> := map[v2 := v13];
			v14 := v14[v2 := DC20({v11}).(cf36 := v12)];
			var v15: array<int> := new int[24](i2 => i2 + p0);
			v15[249] := p0;
			var v16: array<map<bool, bool>> := new map<bool, bool>[10];
			v16[948] := map[v2 := v2];
			var v17: map<int, string> := map[p0 := v1[932]];
			var v18 := 'i';
			var v19: map<bool, bool> := map[!v4[p0] := v2];
			v15[249], v0, v2, v16[948], v7 := fm44(!(p0 < |{v2}|), globalState), (if (v2) then if (p0 in v17) then v17[p0] else (seq(0x78, i3  => ('s')))[p0 := v18] else "fhkvhgy")[p0 := v18] + v0, v2, v19, v7;
		}
		var v20 := true;
		var v21: seq<int> := [p0];
		match (fm45(v20, v20, |v21| % p0, globalState)) {
			case DC35(cf63, cf64, cf65, cf66, cf67) =>
				var v22: map<int, int> := map[cf63 := cf65];
				var v23 := DC29(v20, if (p0 in v22) then v22[p0] else cf65, v20, -(cf65 + cf63), cf66);
				match (v23) {
					case DC28(cf46, cf47) =>
						cf47 := p0 + cf47;
						var v24: seq<bool> := [v21[cf63] <= cf47];
						v24 := v24;
						cf63 := cf65;
						cf65 := cf47;
					case DC29(cf48, cf49, cf50, cf51, cf52) =>
						var v25: array<int> := new int[7];
						v25[982] := cf65 - p0;
						v25[982] := cf49;
						v25[982] := |v0| - cf63;
						var v26 := DC14(v0);
						var v27 := 'd';
						var v28: seq<string> := [cf64];
						var v29: array<string> := new string[28] [v26.cf29[fm44(cf48, globalState) := v27] + v0, v0, cf64 + fm46(true, |[cf50]|, cf48, globalState), "y", v0, cf64, "rc", cf64 + v0, v0, cf64, "ovqriexd", fm46(true, fm44(cf48, globalState), cf50, globalState), seq(0x86, i4  => (v27)), v0, v26.cf29, v0, cf64, seq(552, i5  => ('f')), v0, v0, seq(-0x2dd, i6  => (v27)), cf64 + v28[cf65], "crcuft" + (seq(0x257, i7  => (v27))), DC14(cf64).cf29, v0, v0, "nys", "jijhsve"];
						v29[595] := cf64;
						var v30: multiset<string> := multiset{("sfisdg")[cf63 := v27], cf64};
						var v31: map<bool, multiset<string>> := map[cf66 := v30];
						var v32: map<int, bool> := map[p0 := cf48];
						v25, v29[595], cf63, cf63 := v25, v0, -cf65, -0x1e0 % |(if ((if (v25[982] in v32) then v32[v25[982]] else cf48) in v31) then v31[if (v25[982] in v32) then v32[v25[982]] else cf48] else v30)|;
						cf67 := cf48;
					case DC27(cf45) =>
						var v33: array<int> := new int[29](i8 => i8 - |[true]|);
						var v34 := DC23(v33);
						v34 := v34;
						v33[921] := cf63;
						v33[921] := -cf65;
						var v35 := 'h';
						var v36 := DC1(|cf64|);
						var v37: map<char, D0> := map[v35 := v36];
						var v38: map<bool, bool> := map[!cf66 := cf67];
						var v39: array<bool> := new bool[28] [v20, cf66, false, cf67, cf66, cf67, cf67, cf67, cf66, cf66, cf67, cf65 < -cf63, !!!v20, fm42(cf66, globalState), v20, cf66, fm42(cf66, globalState), cf67, cf65 <= cf65, cf67, cf67, !cf66, fm5(v33[921], cf66, v37, fm5(p0, true, map[v35 := v36], cf66, globalState), globalState), false, cf66 <==> cf67, if (fm42(v20, globalState) in v38) then v38[fm42(v20, globalState)] else v20, v20, cf66];
						var v40 := DC5(cf65, v21, v20);
						v39[7] := if (cf66) then v40.cf10 else cf67;
						v39[7] := cf67 || cf66;
						v33 := v33;
				}
				
				var v41: array<bool> := new bool[27](i9 => v20);
				var v42: seq<bool> := [true, false];
				v41[737] := v42[p0];
				var v43: array<int> := new int[5](i10 => i10 * p0);
				v43[919] := cf65;
				var v44 := 't';
				var v45 := DC1(p0);
				var v46: map<char, D0> := map[v44 := v45];
				v41[737], v43[919] := v0 == v0, fm44(fm5(cf65, cf66, v46, cf66, globalState), globalState);
				if (v20) {
					var v47: map<int, bool> := map[p0 := cf66];
					var v48: set<map<int, bool>> := {v47[v43[919] := v41[737]], map[|fm47(|v0|, v41[737], globalState)| := v41[737]]};
					var v49 := DC31(false, |map[v20 := cf65]|, v43[919]);
					var v50: multiset<int> := multiset{cf63, |v48|, |v42| % v49.cf56};
					var v51 := DC18();
					v50 := multiset{-fm44(false, globalState), cf65} + fm48(v51, v43[919], fm44(true, globalState), -50, globalState);
					var v52: map<map<char, bool>, int> := map[map[v44 := !v41[737]] := cf65];
					v52 := (if (v20) then v52 else v52) + v52;
					var v53: array<map<char, int>> := new map<char, int>[24];
					var v54 := new C1(cf63, v53, v21 + v21, v45);
					var v55: map<string, bool> := map[fm46(v41[737], |v50|, false, globalState) := v41[737]];
					var v56: map<seq<int>, string> := map[v21 := cf64];
					var v57: seq<string> := [v0, if (v21 in v56) then v56[v21] else "mmjnwlr", cf64, ("l")[p0 := v44]];
					var v58: set<bool> := {v41[737], v20, true};
					v41[737] := if (v57[0x22e] in v55) then v55[v57[0x22e]] else v58 < {v41[737]};
					var v59: map<char, int> := map[v44 := v54.f9];
					var v60: map<map<char, int>, char> := map[v59 + v59 := v44];
					v60 := v60;
				} else {
					var v61 := new C0("acttpcc");
					var v62: array<seq<bool>> := new seq<bool>[1] [v42];
					v62[392] := v42;
					v62[392] := v42;
					v20 := cf67;
					cf65 := p0;
					var v63: array<map<char, int>> := new map<char, int>[14];
					var v64 := new C1(v43[919], v63, v21[cf65 := cf65], v45);
				}
				
				if (cf67) {
					var v65: map<int, bool> := map[p0 := v0 <= [v44]];
					v65 := v65[v43[919] := true];
					v20 := !v41[737];
					var v66: set<seq<int>> := {v21};
					var v67 := DC37(v66);
					v41[737] := v66 <= v67.cf70;
					v41[737] := cf66;
					v43[919] := p0 / v43[919];
				} else {
					v43[919] := cf65;
					var v68: set<bool> := {v41[737], fm5(p0, true, map[v0[|map[cf63 := v41[737]]|] := v45], false, globalState)};
					var v69: map<set<bool>, bool> := map[v68 := v20];
					cf65 := v43[919] - |v69|;
					cf67 := cf67;
					var v70: array<map<char, int>> := new map<char, int>[8];
					var v71 := new C1(-834, v70, seq(-0x5c, i11  => (-|multiset{v20}|)), v45.(cf2 := cf63));
					v41[737] := false;
				}
				
			case DC36(cf68, cf69) =>
				var v72: array<int> := new int[3](i12 => i12 - p0);
				var v73: map<bool, array<int>> := map[true := v72];
				var v74 := DC28(v73 == v73[v20 := v72], p0);
				v74 := v74;
				var v75: map<bool, bool> := map[!false := v20 ==> v20];
				v75 := v75[v20 !in cf69 := v20];
				var v76: map<int, bool> := map[p0 := v20];
				var v77: map<bool, map<int, bool>> := map[v20 := v76];
				cf69 := cf69[v20 !in v77 := p0];
				if (v20) {
					var v78 := 0x279;
					v78 := p0;
					var v79: set<bool> := {true, v20};
					var v80: multiset<set<bool>> := multiset{v79};
					var v81: map<bool, multiset<set<bool>>> := map[v20 := v80];
					v80 := v80[v79 := v78] + (if (v20 in v81) then v81[v20] else multiset{v79});
					v73 := v73[if (v20 in v75) then v75[v20] else v20 := v72];
					var v82: seq<bool> := [!v20, v20, v20];
					var v83: seq<seq<bool>> := [v82, v82];
					var v84: map<bool, seq<bool>> := map[v82[v78] := v83[v21[v78]]];
					v84 := v84;
					var v85 := new C0(v0);
				} else {
					var v86: array<bool> := new bool[19];
					v86[133] := v20;
					v21, v86[133] := v21, v20 <== v20;
					var v87 := DC15(p0, fm44(v86[133], globalState), !v20);
					v87 := v87;
					v86[133] := v0 <= v0;
					var v88: map<map<bool, int>, string> := map[cf69 := v0];
					var v89 := m14(false, v86[133] && v86[133], v88 + map[cf69 := v0], -|v0|, globalState);
					var v90 := 's';
					var v91: map<bool, char> := map[false := v90];
					v91 := v91[v86[133] := v90];
				}
				
			case DC34(cf62) =>
				var v92 := DC28(v20 <== v20, p0);
				v92 := v92;
				v20 := v20;
				var v93: array<set<bool>> := new set<bool>[9](i13 => {v20, v20} - {true});
				var v94: set<bool> := {v20};
				v93[60] := v94;
				v93[60] := v94;
				var v95 := DC2(v20);
				if (!v95.cf3) {
					var v96 := 'c';
					var v97: map<int, int> := map[fm44(v20, globalState) * p0 := p0];
					v96, v20, v97, v20 := if (v20) then v96 else v96, v20 && v20, v97, map[p0 := -p0] == (v97 + v97);
					var v98: map<int, bool> := map[p0 + p0 := v20 in v93[60]];
					v98 := v98[p0 := true];
					v20 := v20;
					var v99: multiset<bool> := multiset{true, v20, false};
					var v100: map<int, multiset<bool>> := map[p0 := v99];
					var v101: array<multiset<bool>> := new multiset<bool>[19] [v99, v99, v99 + v99, multiset{v20}, v99, v99 + v99, multiset{v20, v20}, v99, v99, v99, v99, v99, if (p0 in v100) then v100[p0] else v99, v99, v99, v99, (multiset{v20, v20})[v20 := p0], v99, v99];
					var v102: seq<multiset<bool>> := [multiset{v92.cf46}];
					v101[741] := v102[p0];
					v101[741] := v99;
					v20 := p0 >= p0;
				} else {
					var v103: array<bool> := new bool[24];
					var v104 := DC41(v103);
					v103 := v104.cf75;
					var v105 := 0x31b;
					var v106: map<int, bool> := map[v105 := v20];
					var v107 := DC1(v105);
					v105 := (p0 * fm44(if (p0 in v106) then v106[p0] else v20, globalState)) % (v107.cf2 - |[v20]|);
					v20 := !v20;
					v20 := p0 != -(|v0| / p0);
					var v108: array<D0> := new D0[13];
					v108[631] := v107;
					var v109 := 'l';
					v108[631] := fm49(v109, globalState);
				}
				
		}
		
		if (!v20) {
			var v110: multiset<bool> := multiset{v20};
			var v111: map<int, bool> := map[|v110| := v20];
			var v112: seq<map<int, bool>> := [v111];
			var v113: array<map<int, bool>> := new map<int, bool>[5] [v111, v111, fm50(v20, v20, p0, globalState), v112[p0], (map[p0 := v20])[-0x315 := v20]];
			var v114: map<bool, array<map<int, bool>>> := map[v20 := v113];
			var v115: C0 := new C0(v0);
			var v116: seq<C0> := [v115];
			v114 := v114[|v116| >= p0 := v113];
			var v117: array<int> := new int[12];
			v117[99] := fm44(!v20, globalState);
			v117[99] := p0;
			v21 := v21 + (seq(-0x131, i14  => (p0)));
			v20 := v20;
			var v118: set<int> := {p0, 0x97, v117[99]};
			v118 := v118 * v118;
		} else {
			v20 := v20;
			var v119: array<array<int>> := new array<int>[7];
			var v120: set<bool> := {v20};
			var v121: seq<set<bool>> := [v120, v120];
			var v122 := DC27(v120);
			var v123: array<set<bool>> := new set<bool>[18] [v120, v120, v121[p0], v120, v120, v120, v120, v120, v120, v120, v120, v122.cf45, v120, v120, v120, v120, v120, v120];
			var v124 := DC0(v123, 596);
			var v125 := DC38(fm51(v20, p0, p0, globalState));
			var v126: map<D0, char> := map[v124 := v125.cf71];
			var v127: seq<bool> := [v20];
			var v128: seq<seq<int>> := [v21];
			var v129: array<int> := new int[23] [p0, p0, 550, |v0|, fm44(v20, globalState), p0, p0, --p0, p0, p0, p0, |v0|, p0, |v126|, |v127|, |v128|, p0, p0, -p0, 0x103, p0, p0, -31];
			v119[170] := v129;
			var v130: array<bool> := new bool[14];
			v130[703] := v20;
			v129[483] := fm44(v20, globalState);
			v129[844] := p0;
			v119[170], v130[703], v129[483], v129[844] := v129, v20, p0, p0 % v21[p0];
			v20 := fm42(v20, globalState);
			var v131: set<array<bool>> := {v130};
			var v132 := DC20(v131);
			var v133 := DC20(v132.cf36);
			match (v133) {
				case DC20(cf36) =>
					var v134: map<string, int> := map[seq(0x206, i15  => ('f')) := -p0];
					var v135: map<int, bool> := map[|v134| := v130[703]];
					var v137 := 'f';
					v20 := fm5(|v135|, v20 <== v20, map v136 : char | v136 in multiset{v137, v137, v137, v137, v137} :: (v136) := (DC1(p0)), v20, globalState);
					v123[607] := v120;
					v123[607] := v120;
					var v139: array<multiset<map<bool, int>>> := new multiset<map<bool, int>>[6](i16 => multiset{map[v130[703] := v129[483]]} + multiset{map[v130[703] := |(set v138 : multiset<int> | v138 in {multiset{|map[v130[703] := v130[703]]|, p0}} :: (v138))|]});
					var v140: map<int, int> := map[p0 := v129[483]];
					var v141: multiset<map<bool, int>> := multiset{map[!v130[703] := |v140|]};
					v139[64] := v141;
					v139[64] := v141;
					var v142 := DC32(v0, v130[703], v129[483], v119[170]);
					var v143: map<bool, int> := map[!v142.cf58 := p0];
					var v144: multiset<D8> := multiset{v133};
					var v145: map<map<bool, int>, int> := map[v143 := if (v132.(cf36 := {v130, v130}) in v144) then v144[v132.(cf36 := {v130, v130})] else |v0|];
					var v146 := DC1(v129[483]);
					var v147: map<char, D0> := map[v137 := v146];
					v145 := if (v20 && fm5(630, v20, v147, v130[703], globalState)) then map[map[v20 := p0] := if (true in v143) then v143[true] else p0] else map[map[v130[703] := -0x1c3] := p0];
			}
			
			var v148: array<char> := new char[19](i17 => 's');
			var v149: seq<array<char>> := [v148];
			v149 := v149;
		}
		
		var v150: map<string, bool> := map[v0 := v20];
		v150 := v150[v0 + (seq(680, i18  => ('a'))) := v20];
		if (!false) {
			var v151: array<int> := new int[15];
			var v152 := DC23(v151);
			v152 := v152;
			var v153 := 404;
			v153 := |((seq(-0x5a, i19  => ('b'))) + v0)|;
			var v154 := 'g';
			var v155: map<char, int> := map[v154 := p0];
			v155 := v155[v154 := v153];
			v154 := v154;
			v153 := -p0;
		} else {
			var v156 := DC10(|v0|, -p0, |map[v20 := false]|);
			var v157 := 'a';
			var v158: map<bool, char> := map[v20 := v157];
			var v159: set<int> := {|v158|};
			var v160: array<int> := new int[24] [fm44(v20, globalState), p0, 0x2ca, p0, p0, p0, v156.cf20, p0, p0, p0, p0, p0, p0, |v159|, p0, p0, p0, p0, |v0|, p0, p0, p0, p0, p0];
			var v161: array<array<int>> := new array<int>[12] [v160, v160, v160, v160, v160, v160, v160, v160, v160, v160, v160, v160];
			var v162 := DC4(p0, v20, v161);
			match (v162) {
				case DC4(cf5, cf6, cf7) =>
					v20 := false;
					var v163: array<map<char, int>> := new map<char, int>[16];
					var v164: multiset<bool> := multiset{v20};
					var v165 := DC21(v164);
					var v166 := DC39(v0, v165);
					var v167: seq<string> := [v0, v166.cf72];
					var v168 := new C1(p0, v163, (seq(0x17a, i20  => (|v21|))) + v21, DC1(|v167|));
					v160 := v160;
					v0 := v0;
				case DC5(cf8, cf9, cf10) =>
					v160[240] := 711;
					var v169: multiset<bool> := multiset{true, v20};
					v160[240] := (cf8 % |v0|) / (if (v20 in v169) then v169[v20] else p0);
					var v170: seq<string> := [v0, v0];
					var v171: map<seq<string>, int> := map[v170 := p0];
					v160[240] := if ((seq(-0x123, i21  => (v0))) in v171) then v171[seq(-0x123, i21  => (v0))] else -cf8 - cf8;
					var v172: array<bool> := new bool[10];
					v172[535] := cf10 <==> true;
					v172[535] := v20;
					v160[240] := (v160[240] % -0x390) / (|v0| * p0);
				case DC3(cf4) =>
					var v173 := 0x25a;
					var v174 := DC38(if (v20) then fm51(v20, v173, v173, globalState) else v157);
					var v175: map<bool, seq<int>> := map[v20 := v21];
					v173, cf4, v174, v20 := |(if (v20) then if (v20 in v175) then v175[v20] else seq(0xa, i22  => (p0)) else seq(-0x334, i23  => (v173)))|, cf4 * cf4, v174, v20;
					v173 := (p0 + -p0) + v173;
					var v176: array<bool> := new bool[18];
					var v177: multiset<bool> := multiset{v20};
					v176[612] := multiset{v20, v20} > v177;
					v176[612] := v20;
					var v178: map<int, bool> := map[p0 := v20];
					v178 := v178;
				case DC6(cf11) =>
					var v179: seq<bool> := [v20, v20];
					var v180: map<char, int> := map[v157 := 68];
					var v181 := DC38(v157);
					var v182: multiset<int> := multiset{p0, 0x13d};
					var v183: array<map<char, int>> := new map<char, int>[9] [v180, fm52(true, v181, v20, globalState), v180, v180, v180, map[v157 := fm44(v20, globalState)], map[v157 := if (p0 in v182) then v182[p0] else p0], v180, v180];
					var v184: T1 := new C1(|v179|, v183, v21, DC1(p0));
					var v185: map<T1, char> := map[v184 := 'b'];
					v185 := v185[v184 := 'e'];
					v160[353] := p0 % |v179|;
					v160[353] := 0x25a - 0x62;
					var v186 := new C0(v0);
					var v187: array<bool> := new bool[20];
					v187[825] := if (v20) then v20 else v20;
					var v188: multiset<seq<bool>> := multiset{v179};
					v160[353], v187[825], v188 := |v182|, v20, v188;
			}
			
			v160[653] := p0;
			var v189: map<bool, int> := map[false := p0];
			var v190: seq<map<bool, int>> := [v189, v189, v189, v189];
			v160[653] := p0 + |fm53(v190[p0], globalState)|;
			var v191: map<int, D2> := map[p0 := v162];
			var v192: map<bool, map<int, D2>> := map[v20 := v191];
			var v193: map<int, int> := map[p0 := v160[653]];
			v160[653], v160[653], v20, v20, v160[653] := v160[653], if (v189 != v189) then 52 + 0x1f3 else p0, v160[653] > v160[653], v191 != (if (v20 in v192) then v192[v20] else map[v160[653] := DC4(0x1b6, v20, v161)]), p0 * ((if (|multiset(v21)| in v193) then v193[|multiset(v21)|] else p0) / v160[653]);
			var v194: map<int, bool> := map[fm44(!v20, globalState) := v20];
			var v195: set<char> := {v157};
			v20 := if (p0 in v194) then v194[p0] else !({v157, v157, v157, v157} < v195);
			var v196: seq<bool> := [v20];
			var v197: set<bool> := {v20};
			var v198: multiset<int> := multiset{fm44(v20, globalState)};
			var v199: seq<D13> := [DC32(v0, v20, v160[653], v160)];
			var v200: array<bool> := new bool[20] [if (v20) then v20 else v20, true, v196[-v160[653]], !v20, v20, false, v197 !! {v20, v20}, false, v20, v20, v198 > multiset{if (v20 in v189) then v189[v20] else p0, |v199|}, v20, p0 >= v160[653], if (v160[653] in v194) then v194[v160[653]] else v20, fm42(v20, globalState), if (v20) then v20 else v20, if (v20) then v20 else v20, !v20, v20, v157 in v0];
			v200[166] := v20;
			v200[166] := p0 == p0;
		}
		
		var v201 := DC10(p0, p0, p0);
		r0 := match if (v20) then v201.(cf21 := p0) else v201 {
			case DC9(cf14, cf15, cf16, cf17, cf18) => map[v20 := multiset(v21)]
			case DC10(cf19, cf20, cf21) => map[v20 := multiset(DC5(|(seq(842, i24  => ('c')))|, v21, v20).cf9)]
			case DC11(cf22, cf23, cf24, cf25, cf26) => map[cf26 := multiset{cf22}]
			case DC8(cf13) => map[v20 := multiset{|map[map[|{true}| := v20] := v20]|}] + map[v20 := multiset{p0, |[v20, v20]|}]
			case DC12(cf27) => map[false := multiset{p0}] + map[v20 := multiset{p0}]
		};
	}
	method m14(p0: bool, p1: bool, p2: map<map<bool, int>, string>, p3: int, globalState: GlobalState) returns (r0: int) {
		r0 := fm44(p1, globalState);
		var v0 := "wnm";
		var v1: map<int, bool> := map[p3 := true];
		for i0 := -|v0| to 0x125 + |v1| {
			var v2: array<bool> := new bool[8];
			v2[332] := p1;
			v2[332] := p1 || (if (p0) then false else p1);
			var v3: set<bool> := {p0};
			var v4: array<set<bool>> := new set<bool>[18] [v3, v3, v3, v3, v3, v3, v3, v3, {v2[332], p0}, {p0}, v3, v3, v3, {v2[332]}, v3, v3, v3, v3];
			var v5 := DC0(v4, 0x2);
			var v6: seq<D0> := [v5, v5];
			var v7 := DC8(v6);
			v7 := v7;
			var v8 := 'n';
			var v9: map<char, int> := map[v8 := p3];
			var v11 := DC7(if (p1) then v9 else map v10 : char | v10 in v0 :: (v10) := (p3));
			v11 := fm54(v2[332], [p1, v2[332]], globalState);
			r0 := -p3;
		}
		var v12 := false;
		v12 := p0;
		for i1 := p3 to if (p1) then p3 else p3 {
			var v13 := 'g';
			v13 := if (v12) then v13 else v13;
			v12 := v0 <= v0;
			v0 := if (p1) then "w" else v0;
			var v14: array<int> := new int[17];
			v14[783] := p3;
			var v15: seq<bool> := [false];
			v14[783], r0, v15 := |(v1 + v1)| - i1, i1, v15 + [DC35(p3, "sia", i1, v12, p0).cf67, p0, p1, v12];
		}
		var v16: set<int> := {p3, p3};
		var v17 := DC15(|v16|, p3, p0);
		var v18 := DC16(v17);
		var v19: multiset<D6> := multiset{v18, DC16(DC16(v17).cf33)};
		v12 := !((if (p0) then p3 else p3) !in (if (!p1) then seq(0xcc, i2  => (p3)) else [|v19|, p3, p3, p3]));
		v0 := v0;
		r0 := p3;
	}
}

class C3 extends T0 {
	constructor () {
	}
	
	function fm5(p0: int, p1: bool, p2: map<char, D0>, p3: bool, globalState: GlobalState): bool {
		!!((0x26e - ---0x30b) in ([|[|"hi"|, |['x', 'f']|]|] + [|"fqgtkbdg"|, |{-285, |map['w' := 'k']|, |map[-394 := 0x304]|}|]))
	}
	method m1(p0: int, globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0 := false;
		if (v0) {
			var v1, v2, v3, v4 := m12(v0 ==> v0, globalState);
			if (v0) {
				var v5: array<map<char, int>> := new map<char, int>[26];
				var v6 := 'g';
				var v7 := DC1(p0);
				var v8 := "rnnc";
				var v9: map<bool, string> := map[fm5(p0, false, map[v6 := v7], v4, globalState) := v8];
				var v10: set<int> := {p0, v3};
				var v11: multiset<int> := multiset{p0};
				var v12: seq<int> := [|v9|, |v10|, |v11|, fm32(p0, globalState), v3];
				var v13: T1 := new C1(--0x303, v5, v12, v7);
				var v14: seq<T1> := [v13];
				v14 := v14;
				var v15: array<int> := new int[9](i0 => i0 % |[!v0]|);
				v15[189] := v2;
				v15[189] := v2;
				var v16: set<bool> := {v4};
				var v17 := DC27(v16);
				var v18: map<bool, set<bool>> := map[v4 := v16];
				v17 := DC27(if (v0 in v18) then v18[v0] else {v0});
				v15 := v15;
				var v19 := DC14(v8);
				var v20: multiset<bool> := multiset{v1};
				var v21: map<int, int> := map[v3 := if (v4 in v20) then v20[v4] else 0x106];
				var v22: array<D6> := new D6[12] [v19, v19, fm33(if (v3 in v21) then v21[v3] else p0, seq(0x38f, i1  => ('s')), globalState).(cf29 := seq(0x11f, i2  => ('i'))), v19, v19, v19, v19, v19, DC14(v8), v19, DC14(v8), fm33(v15[189], seq(0xbe, i3  => (v6)), globalState)];
				v22[402] := v19.(cf29 := v8);
				v22[402] := v19;
			} else {
				var v23: array<map<D11, bool>> := new map<D11, bool>[4];
				v23 := v23;
				v3 := -p0;
				r1 := v4;
				var v24: array<string> := new string[13];
				var v25 := "gvm";
				v24[93] := v25;
				v24[93] := v25;
				var v26: map<bool, bool> := map[v4 <==> v0 := v0];
				v26 := v26 + v26;
			}
			
			var v27 := "opafnutak";
			var v28: seq<string> := [v27, seq(0x26c, i4  => ('x')), v27, v27, v27];
			var v29 := new C0(v28[0xab]);
			if (true) {
				var v30: seq<bool> := [v4];
				var v31: seq<seq<bool>> := [v30, v30];
				var v32: map<int, int> := map[p0 := v3];
				var v33: seq<int> := [|v27|, 0x229, if (-0x169 in v32) then v32[-0x169] else p0, -v3, v2];
				var v34: set<int> := {p0};
				var v35: array<int> := new int[28] [v3, v2, v3, v2, v2, p0, p0, p0, v2, p0, p0, v2, v2, 0x164, v2, v3, -0x2d3, |multiset(v31[v3])|, v33[v3], v3, p0, v3, v2, |v34|, v3, p0, v2, p0];
				var v36: map<int, array<int>> := map[|[true]| := v35];
				var v37 := DC28(v0, v2);
				var v39: array<map<char, int>> := new map<char, int>[28](i5 => map v38 : char | v38 in {'s'} :: (v38) := (p0));
				var v40 := DC1(p0);
				var v41: C1 := new C1(p0, v39, v33, v40);
				var v42: map<C1, bool> := map[v41 := false];
				var v43: set<bool> := {!v1};
				var v44: array<bool> := new bool[20] [v4, v3 < v3, v1, v2 < |v36|, false, !v0, !v1, v1, if (v4) then v37.cf46 else v1, v1, v42 !in map[v42 := v43], v1, v4, v4 <== false, v1, v3 >= v41.f9, false, v34 >= v34, !(if (v4) then v1 else v0), v1];
				v44[637] := v4;
				v44[637] := v0;
				v3 := v2;
				var v45: map<bool, int> := map[v0 := v41.f9 + (fm34(v1, v0, fm30(v41.f9, v44[637], v44[637], globalState), 0x171, globalState)).cf47];
				v3 := |v45|;
				v41.f9 := |v29.f8|;
				var v46 := new C0(seq(-0x8, i6  => (DC25('w', v2, v30[0x251]).cf41)));
			} else {
				var v47: multiset<bool> := multiset{v1};
				var v48: multiset<int> := multiset{|v47|};
				var v49: array<bool> := new bool[16](i7 => v1);
				var v50 := DC28(v1, 0x20c);
				var v51 := 'y';
				var v52: map<bool, char> := map[v0 := 'c'];
				var v53 := DC1(|v52|);
				var v54: map<char, D0> := map[v51 := v53];
				var v55: map<bool, bool> := map[v1 := fm30(--p0, fm5(v2, v4, v54, v0, globalState), v4, globalState)];
				var v56: map<int, bool> := map[-v2 := v1];
				var v57: map<bool, int> := map[v0 := |map[v3 := v1]|];
				var v58: array<int> := new int[22] [v2, v2, v3, v2, fm32(p0, globalState) + p0, v2 / -v2, v2, -p0, v3, |v48|, p0, |multiset{v49, v49, v49}|, p0, v50.cf47, |v55|, v2 + |v56|, p0, v2, |v57|, p0, -p0, v2 - v2];
				v58[289] := if (!v1) then v3 else v2;
				v58[289] := fm32(174, globalState) + p0;
				var v59: array<multiset<int>> := new multiset<int>[8];
				var v60: multiset<char> := multiset{v51};
				v3, v59, v60 := p0, v59, v60;
				var v61: seq<bool> := [v1, v4, v0];
				v61 := (if (v1) then v61 else v61)[p0 := |v56| >= v58[289]];
				var v62: map<int, char> := map[0x2b1 := v51];
				var v63: map<map<int, char>, int> := map[v62 := v2];
				v63 := v63[v62 := -v2];
				v27 := v27;
			}
			
			var v64 := new C0(v27);
		} else {
			var v65: seq<bool> := [v0];
			var v66: multiset<bool> := multiset{v0, true, v0, v0, v65[p0]};
			var v67: map<int, seq<int>> := map[|map[false := p0]| := fm35(v0, false, globalState)];
			var v69: seq<int> := [p0, p0, p0];
			v0 := (if (!v0 in v66) then v66[!v0] else p0) <= |(v67 + (map v68 : int | v68 in v69 :: (v68 - -p0) := (v69)))|;
			var v70 := "is";
			var v71 := 'v';
			var v72: seq<char> := [v70[p0], v70[p0], v71, v71];
			var v73: set<bool> := {v0};
			var v74 := DC25(v70[|v73|], p0, v0);
			var v75: map<bool, int> := map[v0 := p0];
			var v76 := DC29(v0, p0, false, --|v75|, v0);
			match (if (v0 <==> v74.cf43) then v76 else v76) {
				case DC28(cf46, cf47) =>
					var v77: map<int, char> := map[-p0 := v71];
					var v78: array<char> := new char[22] ['c', v71, v71, v71, v71, v71, v71, v71, 'i', if (0x1fd in v77) then v77[0x1fd] else v71, v71, v71, v71, v71, v71, v71, 'h', 'b', fm29(cf46, globalState), v71, v71, v71];
					var v79: map<array<char>, set<bool>> := map[v78 := v73];
					v79 := v79;
					var v80: array<map<map<bool, D4>, multiset<int>>> := new map<map<bool, D4>, multiset<int>>[2];
					var v81: multiset<int> := multiset{-804};
					var v82: C0 := new C0(v72);
					var v83: set<C0> := {v82, v82, v82, v82};
					var v84 := DC9(v81, v71, cf47, v83, 712);
					var v85: map<map<bool, D4>, multiset<int>> := map[(map[true := DC9(v81, v71, 0xcf, v83, p0)])[v0 := v84] := multiset(v69)];
					v80[682] := v85 + v85;
					v80[682] := v85;
					var v86: array<set<bool>> := new set<bool>[25];
					var v87 := DC0(v86, cf47);
					var v88: map<int, bool> := map[fm32(p0, globalState) := v0];
					var v89: map<char, int> := map[v71 := |v88|];
					var v91: array<map<char, int>> := new map<char, int>[17] [v89, v89[v71 := |v72|], fm36(v71, -|map[v75 := v71]|, fm32(|map[v0 := cf47]|, globalState), DC27(v73).cf45, globalState), v89, v89, v89, v89, v89, v89, map v90 : char | v90 in v70 :: (v90) := (cf47), v89, map[v71 := 0x306], v89, v89, map[v71 := p0], v89, v89[v71 := p0]];
					var v92 := DC1(|v88|);
					var v93 := new C1(v87.cf1, v91, [p0, -p0], v92);
					var v94: array<bool> := new bool[22] [v0, v0, cf46, v0, true, v0, v0, false, v0, cf46, cf46, !cf46, cf46, cf46, v0, !cf46, v0, cf46, v0, v0, false, v0];
					var v95, v96, v97, v98 := m11(v71, v94, v66, v93.f9 / v93.f9, globalState);
				case DC29(cf48, cf49, cf50, cf51, cf52) =>
					var v99: array<map<char, int>> := new map<char, int>[26];
					var v100 := DC1(cf49);
					var v101 := new C1(cf49, v99, v69, v100);
					cf51 := 0x354;
					cf49 := if (cf50) then 501 else cf51 / |(seq(0x1b4, i8  => (v71)))|;
					var v102 := new C0((v70 + v72)[p0 := 'p']);
				case DC27(cf45) =>
					r1 := p0 < |("oggyllotw")[|v75| := v71]|;
					var v103 := 0x159;
					v103 := p0;
					v71 := v71;
					v103 := |{v0, v0, v0, v0}| % 469;
			}
			
			if (false) {
				var v104: array<int> := new int[2];
				var v105: map<int, bool> := map[p0 := v0];
				v104[318] := |"uxuolybl"| % |v105|;
				var v106: multiset<seq<int>> := multiset{v69};
				var v107 := DC19(|v106|);
				v104[318] := (p0 + |(set v108 : D7 | v108 in [DC19(-p0), fm37(true, globalState), v107, v107, v107] :: (v108))|) / |(map[v0 := p0] + v75)|;
				var v109: multiset<int> := multiset{p0};
				var v110: C0 := new C0(v72);
				var v111: set<C0> := {v110};
				var v112 := DC9(multiset{p0}, v72[p0], 748, v111, 83);
				var v113: array<bool> := new bool[17] [if (-0x311 in v105) then v105[-0x311] else v0, v0, v0, v71 !in v70, v109 !! v112.cf14, v0, v0, v0 <==> !false, v0, v0 && v0, v0, v0, v0, false, v0, v0, v0];
				v113 := new bool[11];
				var v114: array<seq<string>> := new seq<string>[8];
				var v115: seq<string> := [v72, v72, v72];
				v114[171] := v115;
				v114[171] := v115;
				v110.f8 := ("kmpvievok" + v72) + v110.f8;
				var v116: map<int, map<int, string>> := map[|v110.fm15(v69, v65[980], globalState)| := map[if (!v0 in v75) then v75[!v0] else p0 := v72]];
				var v117: map<int, string> := map[v104[318] := v110.f8];
				v116 := v116[p0 := v117];
			} else {
				v70 := v70;
				v0 := v0;
				v66 := v66;
				r1 := fm35(v0, v0, globalState) < v69;
				r0 := v0;
			}
			
			var v118 := 0x80;
			v118 := p0;
			r1 := v0;
		}
		
		var v119 := -72;
		v119 := v119 % p0;
		var v120: seq<int> := [p0, v119, v119, p0, p0];
		var v121: multiset<int> := multiset{p0, 0x15c};
		var v122: C0 := new C0("jctntdft");
		var v123: set<C0> := {v122};
		for i9 := v120[DC9(v121, 'n', p0, v123, p0).cf16] to v119 {
			var v124: array<int> := new int[4];
			v124[200] := v119;
			var v125: array<seq<int>> := new seq<int>[17];
			var v126: map<int, int> := map[i9 := p0];
			var v127 := DC30(v125);
			v119, v124[200], v125 := if (-i9 in v126) then v126[-i9] else if (i9 in v121) then v121[i9] else -v119, p0, (v127.(cf53 := v125)).cf53;
			r1 := !(if (v0 <== v0) then v0 else v0);
			v122.f8 := v122.f8;
			var v128: seq<seq<int>> := [v120, v120];
			if (if (|v128[v119]| <= |v120|) then -986 != 792 else v0) {
				v124[200] := (v119 / v124[200]) * |[v0, v0]|;
				var v129: map<bool, array<int>> := map[v0 := v124];
				v129 := v129[v0 := v124];
				v122.f8 := v122.fm15(v120, !v0, globalState) + v122.f8;
				v127 := v127;
				v121, v122.f8, v124[200], v120 := v121, v122.f8, i9, if (v0) then v120 + v120 else v120;
			} else {
				v122.f8 := v122.f8;
				var v130: seq<string> := [v122.f8];
				var v131: seq<bool> := [!v0];
				v0, v0, v122.f8 := v0, v0, v130[|v131[i9 := v0]|];
				var v132: array<bool> := new bool[9];
				var v133: array<array<int>> := new array<int>[22];
				var v134 := DC4(289, v0, v133);
				var v135 := 'c';
				var v136: multiset<bool> := multiset{v0};
				var v137 := DC1(|v136|);
				v132[581] := fm5(v134.cf5, v0, map[v135 := v137], v0, globalState);
				v132[581] := if (v0) then v136 > v136[v0 := v124[200]] else if (v0) then !!v0 else v0;
				v124[200] := v124[200];
				var v138 := new C0(v122.f8);
			}
			
		}
		var v139 := new C0(v122.f8);
		var v140: map<bool, int> := map[v0 := v119];
		for i10 := |v140| to if (v0) then v119 else |v140| {
			r1 := v0;
			r1 := true;
			v139.f8 := v139.f8;
			var v141: array<bool> := new bool[25];
			v141[704] := !v0;
			v141[704] := v0;
		}
		var v142: array<int> := new int[19](i11 => i11 - |map[v139.f8 := p0]|);
		v142[605] := v119;
		v142[605] := v119;
		r0 := v0;
		r1 := v0;
	}
	method m2(p0: int, globalState: GlobalState) returns (r0: map<bool, multiset<int>>) {
		var v0 := true;
		var v1: multiset<bool> := multiset{v0, v0};
		var v2: map<int, multiset<bool>> := map[p0 := DC21(v1).cf37];
		v2 := v2[p0 := v1];
		var v3 := DC18();
		var v4: array<set<bool>> := new set<bool>[4];
		var v5: array<bool> := new bool[19] [v0, v0, v0, v0, v0, v0, v0, false, v0, v0, v0, v0, fm30(p0, v0, v0, globalState), true, v0, v0, v0, v0, v0];
		var v6: array<array<bool>> := new array<bool>[10] [v5, v5, v5, v5, v5, v5, v5, v5, v5, v5];
		v3, v4, v6 := v3, v4, v6;
		for i0 := p0 / -0x298 to p0 {
			var v7 := 373;
			var v8 := "fbkmmu";
			var v9 := 'a';
			v7 := |((seq(-424, i1  => ('g'))) + v8[|v8| := v9])|;
			v9 := v9;
			var v10 := new C0(v8 + v8);
			var v11: array<int> := new int[16](i2 => i2 + v7);
			v11[616] := |(seq(625, i3  => (p0)))| + i0;
			v11[616] := if (v0) then v7 else i0 / i0;
		}
		var v12: array<D2> := new D2[14];
		forall i4 | 0 <= i4 < v12.Length {
			v12[i4] := DC5(p0, [p0, p0], v0);
		}
		var v13, v14, v15, v16 := m12(v0, globalState);
		var v17: map<bool, bool> := map[v16 := false];
		for i5 := |(v17 + v17)| to p0 / v15 {
			var v18 := DC22(v15 - v14);
			match (v18) {
				case DC22(cf38) =>
					var v19: seq<bool> := [v16];
					var v20: multiset<seq<bool>> := multiset{v19, [v0] + v19, v19};
					v20 := (v20 + v20) + v20;
					cf38 := p0 + (v14 + v15);
					var v21: array<seq<bool>> := new seq<bool>[23];
					v21[675] := fm38(v16, v14, v16, globalState) + v19;
					v21[675] := v19;
					var v22 := DC5(cf38, seq(0x18, i6  => (-v14)), v16);
					var v23 := DC6(v22);
					var v24: map<bool, int> := map[v13 := p0];
					var v25: seq<map<bool, int>> := [v24];
					var v26: multiset<map<bool, bool>> := multiset{v17, map[true := v13], v17[false := v0], v17, v17};
					var v27: seq<int> := [i5, |(fm38(v13, |v25[if (v17 in v26) then v26[v17] else -fm32(v15, globalState)]|, v16, globalState))[cf38 := v0]|, v15, fm32(|v21[675]|, globalState), cf38];
					var v28: map<D2, bool> := map[v23 := v27 != v27];
					v0 := if (v23 in v28) then v28[v23] else v16;
				case DC21(cf37) =>
					var v29: map<int, bool> := map[v14 := v0];
					var v30 := "cmpc";
					var v31: set<bool> := {v16};
					var v32: map<int, set<bool>> := map[|v30| := v31];
					var v33 := DC27(if (v15 in v32) then v32[v15] else v31);
					var v34: map<map<int, bool>, D12> := map[v29[i5 := v0] := v33];
					v34 := v34[v29 := v33.(cf45 := v31)];
					var v35: set<array<bool>> := {v5};
					var v36 := DC20(v35);
					var v37: array<D8> := new D8[22] [v36, v36.(cf36 := v35).(cf36 := v35), DC20({v5}), v36, DC20(v35), v36, v36, v36, v36, DC20(v35).(cf36 := v35), v36, v36, v36, if (!v0) then v36 else DC20(v35), v36, v36, v36, DC20(v35), v36, v36, v36, v36];
					v37[397] := v36;
					v37[397] := v36;
					var v38 := 'f';
					var v39: seq<bool> := [v16];
					var v40: map<char, int> := map[v38 := v15];
					var v42: set<char> := {v38, v38, v38};
					var v43: seq<map<char, int>> := [v40, v40];
					var v44: array<map<char, int>> := new map<char, int>[25] [map[v38 := |v39|], v40, v40, v40, map v41 : char | v41 in v42 :: (v41) := (|(seq(-0x248, i7  => (v38)))|), map[fm29(v0, globalState) := v15], map[v38 := fm32(v15, globalState)], v40, v40, v40, v43[p0], v43[v14], DC7(v40).cf12, v40, v40, v40, v40, map[v38 := i5], v40, v40, v40, v40, v40, v40, v40];
					var v45: seq<int> := [v14];
					var v46 := DC1(p0);
					var v47 := new C1(i5 + i5, v44, v45, v46);
					v5[582] := v13;
					v5[582] := v15 in fm39(v13, globalState);
			}
			
			v15 := i5;
			var v48 := 't';
			var v50 := DC19(|(set v49 : int | (0xc3 <= v49) && (v49 < 0x17f) :: (v49 % i5))|);
			match (fm40(v13, v48, v16, v50, globalState)) {
				case DC28(cf46, cf47) =>
					v15 := -i5;
					var v51: array<int> := new int[17];
					v51[625] := p0;
					var v52: map<int, array<int>> := map[i5 := v51];
					v51[625] := |[|v52|]|;
					var v53 := "ruavycy";
					var v54 := DC31(cf46, |v53|, |(seq(-504, i8  => (v48)))|);
					var v55: array<D13> := new D13[1] [v54.(cf56 := i5)];
					var v56: map<bool, array<D13>> := map[fm30(i5, v13, true, globalState) := v55];
					var v57: array<array<D13>> := new array<D13>[3] [v55, if (!cf46 in v56) then v56[!cf46] else v55, v55];
					v57[332] := v55;
					v57[332] := v55;
					var v58: seq<int> := [cf47];
					var v59: map<bool, int> := map[v0 := -|(seq(0x3d1, i9  => (-v15)))|];
					cf46, v53, v58, v13 := v16, v53, if (226 == v14) then [v54.cf55] else [i5, cf47, fm32(i5, globalState), |v59|], v16;
				case DC29(cf48, cf49, cf50, cf51, cf52) =>
					var v60: seq<bool> := [cf48, cf52];
					v5[626] := v60[i5];
					v5[626] := v16 || (true ==> v13);
					var v61 := "wnk";
					var v62 := DC14(v61);
					v62 := v62;
					var v63: map<char, int> := map[v48 := cf51];
					var v65: seq<map<char, int>> := [v63[v48 := p0], map[v48 := v14]];
					var v66: array<map<char, int>> := new map<char, int>[18] [v63, map[v48 := -|v1|], v63, v63, v63, map v64 : char | v64 in v63[v48 := cf51] :: (v64) := (0x32d), v63, v63, v63, v63, v63, v63, v63, v63, v63['k' := i5], v63, v65[v15], map[v48 := if (cf50 in v1) then v1[cf50] else v14]];
					var v67: seq<int> := [p0];
					var v68 := new C1(-cf51, if (v16) then v66 else v66, v67[0x227 := i5] + v67, DC1(fm32(-714, globalState)));
					var v69: map<int, int> := map[|(v1 - v1)| := 111];
					v69 := v69[p0 := -(i5 - cf51)];
				case DC27(cf45) =>
					var v70: array<map<char, int>> := new map<char, int>[21];
					var v71: T0 := new C2();
					var v72: map<T0, bool> := map[v71 := true];
					var v73: seq<int> := [|(seq(0x1da, i10  => (v48)))|, |v72|];
					var v74 := DC1(v15);
					var v75 := new C1(v15, v70, v73, v74);
					var v76 := "dkjicxbn";
					var v77 := new C0(v76 + v76);
					var v78: array<int> := new int[18](i11 => i11 / |[[v75.f9]]|);
					v78[84] := |v77.f8|;
					v78[84] := i5;
					var v79: set<int> := {i5};
					v79 := v79;
			}
			
			v48 := v48;
		}
		var v80 := "ohddunlw";
		var v81: set<bool> := {v16, v0};
		var v82: seq<int> := [v15, v15, |v81|];
		var v83: map<bool, multiset<int>> := map[v80 < v80 := multiset(v82 + v82)];
		r0 := v83;
	}
	method m11(p0: char, p1: array<bool>, p2: multiset<bool>, p3: int, globalState: GlobalState) returns (r0: map<int, map<bool, bool>>, r1: bool, r2: bool, r3: set<map<bool, int>>) {
		var v0 := true;
		p1[659] := v0;
		var v1: seq<int> := [p3];
		var v2 := "h";
		var v3: set<multiset<int>> := {multiset{p3}};
		var v4: map<int, set<multiset<int>>> := map[|v2| := v3];
		var v5: array<int> := new int[3] [p3, p3, p3];
		var v6: map<array<int>, int> := map[v5 := p3];
		p1[659], v0, v1, r1 := true, !((if (-0x3a3 in v4) then v4[-0x3a3] else v3) !! v3), [|v6|], v0;
		if (v0) {
			v0 := v0;
			v2 := v2;
			p1[659] := p1[659];
			var v7: map<bool, seq<char>> := map[v0 := if (v0) then [p0, p0] else v2];
			var v8 := -0x3b7;
			var v9: array<char> := new char[11](i0 => p0);
			v9[623] := 'v';
			v7, v8, v9[623], v8 := v7 + v7, |(seq(803, i1  => (p0)))| % (0x1eb - v8), p0, v8;
			r1 := v8 == v8;
		} else {
			var v10 := 'c';
			var v11 := 0x184;
			v5[98] := p3 * fm32(v11, globalState);
			var v12 := DC18();
			var v13: map<bool, D7> := map[p1[659] := v12];
			var v14: seq<map<bool, D7>> := [v13, v13];
			var v15 := DC31(p1[659], v11, p3);
			var v16: multiset<int> := multiset{fm44(true, globalState), v11, v11, fm44(v0, globalState)};
			var v17: multiset<multiset<int>> := multiset{v16, v16, v16 * multiset(v1)};
			var v18: seq<bool> := [v0, v0];
			v10, r2, v11, r2, v5[98] := fm51(v0, p3, |v2| / v1[|v14|], globalState), (if (v0) then v15 else DC31(v0, p3, p3)).cf54, if ((v16[v11 := |v18|] * multiset{0x330}) in v17) then v17[v16[v11 := |v18|] * multiset{0x330}] else v11, !p1[659], |"ghsv"|;
			if (-v5[98] <= --0x2e1) {
				var v19: set<char> := {v10, p0, p0};
				var v20: set<seq<bool>> := {v18, [v0]};
				var v21: map<int, bool> := map[|v19| := fm55(p3, globalState) >= v20];
				var v22: set<int> := {0x33e};
				v21 := v21[|({v5[98], -v11, v5[98]} - v22)| := DC25(v10, v11, false).cf43];
				var v23: array<char> := new char[17];
				v23 := v23;
				var v24: array<seq<int>> := new seq<int>[11];
				v24[29] := v1;
				v24[29] := v1 + v1;
				var v25: multiset<bool> := multiset{true};
				v25 := v25[p1[659] := v11];
				var v26: seq<array<bool>> := [p1, p1, p1, if (p1[659]) then p1 else p1, p1];
				v26 := v26;
			} else {
				p1[659] := v2 < v2;
				var v27: map<bool, bool> := map[p1[659] := p1[659]];
				v27 := v27[true := [v0] != v18];
				r1 := v0;
				var v28: set<bool> := {v0};
				var v29: map<bool, int> := map[v0 := p3];
				p1[659] := v28 !! fm53(v29, globalState);
				var v30: map<int, bool> := map[p3 := v0];
				r1, v2, v11 := -v11 >= (|map[v11 := v0]| * |v30|), v2, fm44(p1[659], globalState) * p3;
			}
			
			v5[98] := v5[98];
			var v31 := DC29(false, p3, p1[659], v11, p1[659]);
			v5[98] := v31.cf49;
			var v32 := DC28(p1[659], -v11);
			var v33: set<int> := {-p3, v11};
			var v34: multiset<D12> := multiset{DC28(p1[659], v11), v32, v32, v32.(cf47 := |v33|), v32};
			v5[98] := |v34|;
		}
		
		v5[913] := |v2|;
		v5[913] := p3;
		v0 := (p1[659] <== !v0) ==> v0;
		var v35, v36, v37, v38 := m12(p1[659], globalState);
		v5[913] := v36;
		var v39: map<bool, bool> := map[v38 := p1[659]];
		var v40 := DC42(0xad, v36, v36);
		var v41: map<bool, int> := map[fm30(v37, p1[659], v35, globalState) := v5[913]];
		r0 := (map[v5[913] := v39])[(v40.(cf76 := |v41|)).cf78 := v39];
		r1 := p1[659];
		var v42 := DC1(v36);
		var v43: seq<bool> := [v35, p1[659]];
		r2 := fm5(|multiset{p1[659]}|, v35, (map[p0 := v42.(cf2 := v5[913])])[p0 := DC1(|v43|)], true, globalState);
		var v44: set<map<bool, int>> := {v41, v41, v41, v41};
		r3 := {map[p1[659] := v5[913]]} * v44;
	}
	method m12(p0: bool, globalState: GlobalState) returns (r0: bool, r1: int, r2: int, r3: bool) {
		var v0 := 0x376;
		var v1: seq<int> := [v0];
		r3, r1, v1 := false, v0, ([v0])[v0 := 0x29b];
		for i0 := v0 to v0 {
			var v2 := "sxccix";
			if (v0 != |v2|) {
				r3, r3 := !p0, p0;
				var v3: array<array<set<int>>> := new array<set<int>>[8];
				var v4: map<bool, bool> := map[!p0 := p0];
				var v5: set<int> := {|v4|, 0x19f};
				var v7: map<bool, int> := map[p0 := i0];
				var v8 := DC36(map v6 : int | (0x17d <= v6) && (v6 < 529) :: (v6 - |[false]|) := (i0), v7);
				var v9: map<int, int> := map[|v2| := i0];
				var v10: array<D14> := new D14[6] [v8, DC36(v9, v7), v8, v8, v8, v8];
				var v11: seq<array<D14>> := [v10, v10];
				var v14: array<set<int>> := new set<int>[25] [v5, v5, v5, v5, v5, {|v11|, i0}, {-i0, i0}, v5, v5, v5, v5, {fm44(p0, globalState)}, v5, v5, set v12 : int | (0x0 <= v12) && (v12 < 289) :: (v12 + |v1|), v5, v5, v5, v5, v5, v5, {v0, |(seq(-0x24f, i1  => ('h')))|, i0, i0}, v5, v5, set v13 : int | (0x242 <= v13) && (v13 < 846) :: (v13 - |v4|)];
				v3[733] := v14;
				v3[733] := v14;
				v14 := new set<int>[17](i2 => v5);
				var v15 := 'f';
				var v16 := new C0(v2 + (seq(0x316, i3  => ('s')))[|v2| := v15]);
				var v17 := DC1(|v7|);
				var v18: map<char, D0> := map[v15 := v17];
				v4 := v4[fm5(i0, p0, v18, p0, globalState) := v2 != v2];
			} else {
				var v19 := DC18();
				v19 := v19;
				r3 := p0;
				var v20: multiset<bool> := multiset{true, p0};
				var v21: map<int, bool> := map[i0 := p0];
				var v22: seq<bool> := [p0];
				var v23: map<multiset<bool>, bool> := map[multiset(v22) := p0];
				var v24: seq<bool> := [p0, fm30(i0, if (|v23| in v21) then v21[|v23|] else p0, true, globalState), p0, p0];
				var v25 := DC11(v0, p0, p0, p0, p0);
				r3 := !(if (p0) then v20 !! multiset(v22) else v25.cf26);
				var v26 := 'k';
				v26 := if (p0) then v26 else v2[v0];
				var v27: array<int> := new int[2] [v0, v0];
				var v28: set<array<int>> := {v27};
				var v29: multiset<int> := multiset{i0, |v28|};
				var v30: map<multiset<int>, bool> := map[v29 := p0];
				r3 := (if (v29 in v30) then v30[v29] else true) <==> p0;
			}
			
			r0 := p0;
			var v31 := 'a';
			var v32: map<bool, int> := map[p0 := i0];
			var v33: map<char, int> := map[v31 := |v32|];
			var v34 := DC7(v33);
			var v35 := DC25(v31, |(map[v34 := v0])[v34 := i0]|, p0);
			match (v35) {
				case DC25(cf41, cf42, cf43) =>
					var v36: map<int, int> := map[v0 := i0];
					var v37: map<int, bool> := map[i0 := p0];
					var v38: seq<bool> := [cf43];
					var v39: set<seq<bool>> := {v38};
					var v40: array<int> := new int[28] [i0, v0, i0, v0 * (if (i0 in v36) then v36[i0] else |(seq(0x19d, i4  => (v31)))|), i0, fm32(i0, globalState), |v2| + v0, i0, i0, v1[|"y"|] - |v37|, |(v39 * v39)|, cf42, i0, |(seq(0x319, i5  => (cf41)))|, i0, 487 % cf42, i0 * |v2|, cf42, -(cf42 + i0), cf42, cf42, -v0, cf42 * |v2|, i0, if (p0) then i0 else 940, v0 / i0, 0x2d7 - -|{cf41}|, i0];
					v40[518] := fm32(v0, globalState) % 0x338;
					var v41: map<bool, array<int>> := map[p0 := v40];
					var v42: map<map<bool, array<int>>, int> := map[v41 := cf42];
					v40[518] := if (map[p0 := v40] in v42) then v42[map[p0 := v40]] else v0;
					var v43: array<bool> := new bool[6](i6 => cf43);
					var v44: map<int, array<bool>> := map[|v36| := v43];
					var v45: seq<array<bool>> := [if (v40[518] in v44) then v44[v40[518]] else v43];
					var v46: multiset<array<bool>> := multiset{v43};
					var v47 := DC41(v45[if (v43 in v46) then v46[v43] else |v38|]);
					v47 := v47;
					var v48: set<int> := {cf42, |(seq(0x3bc, i7  => ('s')))|};
					v48, r1, r1, r3, v0 := v48, 411 % v0, v0 * cf42, if (p0) then false else cf43, if (!p0) then v40[518] else cf42 / i0;
					var v49: seq<seq<bool>> := [v38];
					var v50 := DC1(v40[518]);
					var v51: map<char, D0> := map[cf41 := v50];
					r3 := fm5(-(i0 + |v49|), p0, v51, p0, globalState);
				case DC24(cf40) =>
					r3 := p0;
					var v52: seq<bool> := [p0, p0];
					var v53 := DC1(|map[v31 := i0]|);
					var v54: map<seq<bool>, int> := map[[p0, v52[i0], fm5(i0, p0, map[v31 := v53], p0, globalState), p0, p0] := v0];
					var v55: set<bool> := {p0};
					var v56: seq<set<bool>> := [{p0, p0}];
					v54 := if (v55 > v56[i0]) then fm56(false, globalState) else if (p0) then v54 else v54;
					var v57: array<bool> := new bool[15](i8 => p0);
					var v58: multiset<bool> := multiset{p0};
					v57[820] := !(v58 > v58);
					v57[820] := !v52[v0 + |fm57(p0, p0, p0, globalState)|];
					var v59: array<array<int>> := new array<int>[21];
					v59 := v59;
				case DC26(cf44) =>
					var v60: map<int, bool> := map[i0 := v0 == i0];
					v60 := v60;
					var v61 := new C2();
					var v62: multiset<bool> := multiset{p0};
					r0 := (v62 !! v62) || p0;
					var v63: array<bool> := new bool[14](i9 => !p0);
					var v64: map<int, array<bool>> := map[v0 := v63];
					var v65: map<string, map<int, array<bool>>> := map[v2 + v2 := v64];
					v65 := v65[v2 := v64];
			}
			
			r3 := p0;
		}
		var v66: seq<bool> := [p0];
		var v67: set<int> := {0x3d4};
		var v69: array<bool> := new bool[8] [p0, p0, p0, true, false, p0, p0, p0];
		var v70: multiset<array<bool>> := multiset{v69};
		var v71 := DC11(v0, p0, p0, p0, !p0);
		r0, r0, r3, r3, v66 := !((v67 * (set v68 : int | (0x1bd <= v68) && (v68 < 0x75) :: (v68 + |"yjfgxccy"|))) <= v67), p0, v70 != (v70 * v70), (v71.cf23 <== p0) && p0, fm38(p0, v0, p0, globalState);
		var v72 := DC2(p0);
		var v73: multiset<bool> := multiset{p0};
		var v74: map<D1, multiset<bool>> := map[v72 := v73];
		v74 := v74[v72 := multiset{p0, p0} * v73];
		var v75: map<bool, int> := map[v0 == v0 := |v66|];
		v75 := v75;
		var v76: array<int> := new int[13](i10 => i10 % v0);
		v76 := v76;
		var v77 := "qtblq";
		var v78 := DC21(v73);
		var v79 := DC39(v77, v78);
		r0 := match v79.(cf73 := v78) {
			case DC38(cf71) => p0
			case DC39(cf72, cf73) => p0
			case DC37(cf70) => p0
			case DC40(cf74) => p0
		};
		r1 := v0;
		r2 := v0;
		r3 := p0;
	}
}

class C4 extends T0 {
	constructor () {
	}
	
	function fm5(p0: int, p1: bool, p2: map<char, D0>, p3: bool, globalState: GlobalState): bool {
		!!false
	}
	method m1(p0: int, globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0: array<bool> := new bool[24](i0 => true);
		var v1 := true;
		v0[663] := v1;
		v0[663] := true;
		if (v1) {
			var v2: set<bool> := {v0[663]};
			var v3: array<set<bool>> := new set<bool>[14] [v2, v2, {v0[663], true, v1}, v2, v2, v2, v2, v2, v2, {v1}, v2, {v0[663]}, v2, v2];
			var v4 := DC0(v3, p0);
			var v5: map<bool, D0> := map[v0[663] := v4];
			v5 := v5[v1 := DC0(v3, fm20(!v0[663], globalState))];
			v0[663] := fm20(v1, globalState) != -|map[v0[663] := v1]|;
			var v6: map<int, bool> := map[p0 := v0[663]];
			var v7: map<bool, map<int, bool>> := map[false := v6];
			v7 := v7[true <==> !v0[663] := v6 + fm21(globalState)];
			v0[663] := v1;
			var v8: array<int> := new int[14];
			var v9: multiset<array<int>> := multiset{v8, v8};
			v0[663] := if (!v0[663]) then v0[663] else !(p0 <= |v9|);
		} else {
			var v10: array<seq<int>> := new seq<int>[23];
			v10 := new seq<int>[8](i1 => [p0, p0]);
			var v11: array<array<int>> := new array<int>[5];
			var v12 := DC4(p0, v0[663], v11);
			v0[663], v1 := v1, v12.cf6;
			var v13: multiset<char> := multiset{'t'};
			var v14: map<multiset<char>, int> := map[v13 := p0];
			v14 := v14[v13 - v13 := 0x1da];
			var v15 := new C0("o");
			var v16: set<bool> := {v0[663]};
			var v17: map<array<bool>, set<bool>> := map[v0 := v16];
			var v18: array<set<bool>> := new set<bool>[16] [v16 - v16, v16, {true} * v16, v16, v16, v16, v16, v16, v16, {v0[663], true}, v16, v16 - v16, v16, if (true) then {true, v1, v1, v0[663], v0[663]} else if (v0 in v17) then v17[v0] else v16, v16 * v16, v16];
			v18 := v18;
		}
		
		var v19: array<set<bool>> := new set<bool>[26](i2 => {v1, v1, v0[663], v1});
		var v20 := DC0(v19, p0);
		var v21: seq<D0> := [v20];
		var v22 := DC8(v21);
		if (v22.cf13 <= v21) {
			var v23: map<int, bool> := map[0x2b1 := v1];
			var v24 := "wuqe";
			var v25: seq<int> := [|v24|];
			var v26: map<bool, int> := map[v1 := |map[|v25| := p0]|];
			var v27: multiset<bool> := multiset{v1};
			var v28: multiset<int> := multiset{p0};
			var v29: map<int, int> := map[-p0 := 0x9f];
			var v30: seq<map<int, int>> := [v29];
			var v31: multiset<array<bool>> := multiset{v0};
			var v32: array<int> := new int[22] [p0, |v23|, p0, |v26|, |v24|, p0, p0, |v24|, |v27|, p0, v25[p0], -17, v25[p0], p0, |v28|, p0, |v30[p0]|, if (v0 in v31) then v31[v0] else p0, p0, p0, |fm22(false, p0, |v27|, globalState)|, p0];
			var v33: seq<array<int>> := [v32, v32, v32, v32, v32];
			var v34: map<array<int>, int> := map[v33[p0] := p0];
			v34 := v34[if (v1) then v32 else v32 := p0];
			var v35 := new C0(v24);
			var v36: array<seq<array<int>>> := new seq<array<int>>[3] [v33, v33 + v33, v33];
			v36[233] := [v32];
			v36[233] := [v32];
			var v37: map<bool, bool> := map[v1 := v0[663]];
			match (DC11(|v25| + p0, !false || (if (true in v37) then v37[true] else false), v1, p0 >= -0x2bf, v1)) {
				case DC9(cf14, cf15, cf16, cf17, cf18) =>
					cf18 := (cf16 + -p0) + (if (false) then cf16 else cf18);
					var v38: array<C0> := new C0[9];
					v38[476] := v35;
					v38[476] := v35;
					var v39: map<char, D0> := map['c' := DC1(0x1e6)];
					r1 := fm5(cf16, v1, v39 + v39, if (v0[663]) then v0[663] else !v0[663], globalState);
					v0[663] := if (v0[663]) then !v0[663] else v0[663];
				case DC10(cf19, cf20, cf21) =>
					var v40 := 'k';
					v26 := v26[v40 !in v35.f8 := cf19 + |v25|];
					var v41: set<seq<char>> := {v24, [v40, v40]};
					var v42 := DC3(v41);
					var v43 := DC6(v42);
					var v44: map<D2, string> := map[v43 := v35.f8];
					v44 := v44;
					var v45 := DC1(|v27|);
					var v46: map<char, D0> := map[v40 := v45];
					cf21 := (|[v1, v0[663], if (v0[663] in v37) then v37[v0[663]] else v1]| % cf20) % (if (fm5(cf21, v1, v46, v1, globalState) in v26) then v26[fm5(cf21, v1, v46, v1, globalState)] else |{v0[663]}|);
					v26 := (v26 + v26) + (map[v0[663] := -0x215] + v26);
				case DC11(cf22, cf23, cf24, cf25, cf26) =>
					var v47: seq<bool> := [cf24];
					var v48: seq<seq<bool>> := [v47, fm23(v1, globalState), v47[cf22 := cf25]];
					v48 := v48;
					var v49: seq<array<bool>> := [v0];
					var v50, v51 := m10(cf22, cf22, !(v1 ==> v0[663]), v49[p0], globalState);
					cf22 := cf22;
					var v52: array<array<int>> := new array<int>[10];
					v52 := v52;
				case DC8(cf13) =>
					var v53 := -892;
					v53 := --|(seq(0x10a, i3  => ('m')))|;
					var v54 := 'f';
					v54 := if (v1) then v54 else v54;
					var v55 := DC6(DC5(-v53, v25, v0[663]));
					var v56: seq<D2> := [v55];
					var v57 := DC2(v55 in v56);
					var v58: map<char, array<bool>> := map[v54 := v0];
					var v59: set<C0> := {v35, v35};
					var v60 := DC9((fm24(v0[663], globalState))[--v53 := v53], 'd', |{v1}| + -|v23|, v59, v53 - v53);
					var v61 := DC1(p0);
					var v62: map<char, D0> := map[v54 := v61];
					v57, v0[663], v58, v60 := v57, true || fm5(v53, true, v62, !!v0[663], globalState), v58, v60;
					var v63: array<array<array<int>>> := new array<array<int>>[15];
					var v64: array<array<int>> := new array<int>[17] [v32, v32, v32, v32, v32, v32, v32, v32, v32, v32, v32, v32, v32, v32, v32, v32, v32];
					v63[432] := v64;
					v63[432] := v64;
				case DC12(cf27) =>
					v1 := if (v1) then v0[663] || v0[663] else v0[663];
					v32 := v32;
					v0[663] := v1;
					var v65 := 0xe2;
					var v66: array<map<bool, array<bool>>> := new map<bool, array<bool>>[28];
					var v67: map<bool, array<bool>> := map[v1 := v0];
					v66[390] := v67;
					var v68: array<D3> := new D3[2](i4 => DC7(map['f' := p0]));
					var v69: map<int, array<D3>> := map[v65 := v68];
					var v70 := DC13(v68);
					var v71: array<array<D3>> := new array<D3>[26] [v68, v68, v68, v68, v68, v68, v68, v68, if (|v25| in v69) then v69[|v25|] else v68, v70.cf28, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68];
					v71[0] := v68;
					var v72 := 'p';
					var v73: set<bool> := {v1, v1, v0[663]};
					var v74: seq<set<bool>> := [v73];
					var v75: seq<map<bool, array<bool>>> := [v67];
					r0, v65, v66[390], v71[0], v72 := !false, |v74|, v67 + v75[v65], v68, 'd';
			}
			
			var v76, v77, v78 := m9(globalState);
		} else {
			var v79, v80, v81 := m9(globalState);
			var v83: set<int> := {-990, p0};
			var v84: set<int> := {|v83|, p0};
			var v85: multiset<int> := multiset{p0, |v83|};
			var v86 := 'a';
			var v87: map<int, char> := map[|v85| := v86];
			var v88: seq<map<int, char>> := [(map v82 : int | (-0x250 <= v82) && (v82 < 863) :: (v82 - |"dmetdhseu"|) := ('y')) + v87];
			v88 := v88;
			var v89 := 0x300;
			var v90: seq<bool> := [!v1];
			var v91: seq<seq<bool>> := [v90];
			var v92: seq<int> := [0x2dc, |(v90 + v91[p0])|, p0, p0, v89];
			var v93: map<bool, int> := map[v81 := p0];
			var v94 := "wvfkivcdy";
			v89, v89, v81, v92 := 665, if (v1 in v93) then v93[v1] else v89, (v94 <= v94) ==> (890 == (if (p0 in v85) then v85[p0] else |map[p0 := p0]|)), v92;
			var v95: seq<string> := [seq(0x3c7, i5  => (v86)), DC14("erond").cf29[p0 := 'f'], seq(0x20b, i6  => (v86))];
			var v96 := new C0(v94 + v95[v89]);
			var v97: map<bool, bool> := map[v0[663] := v81];
			var v98 := DC5(|v97|, v92, v81);
			v1 := if (v81) then v98.cf10 else !v90[v89];
		}
		
		var v99: set<int> := {p0, p0, p0};
		var v100: map<bool, int> := map[v0[663] := |v99|];
		var v101 := 'h';
		var v102 := DC1(p0);
		var v103: map<char, D0> := map[v101 := v102];
		if (fm5(|v100|, v1, v103, v1, globalState) ==> (p0 != p0)) {
			var v104: array<int> := new int[22];
			var v105 := "awvnycoec";
			v104[758] := |v105|;
			var v106 := 0x7d;
			v104[155] := p0;
			v104[758], v106, v104[155] := v106, v106 - v106, p0;
			var v107: multiset<array<bool>> := multiset{v0, v0};
			var v108 := DC10(p0, v106, v106);
			v104[758], v105 := |((multiset{v0})[v0 := p0] + v107)| % v106, fm25(v108, v105, globalState);
			var v109: map<int, int> := map[p0 := v104[758]];
			var v110: array<array<int>> := new array<int>[14];
			var v111 := DC4(if (p0 in v109) then v109[p0] else -269, false, v110);
			if (multiset{p0, p0} >= fm24(v111.cf6, globalState)) {
				var v112: seq<bool> := [v0[663]];
				var v113 := DC17(v112);
				var v114: array<seq<bool>> := new seq<bool>[20] [[v111.cf6, false], v112, v112, v112, v112, [v0[663]], v112, [v0[663], v1, v1, v0[663]] + v113.cf34, fm23(v0[663], globalState), v112 + [v1], v112, v112, v112, v112, v112 + v112, [true, v1, false], v112[82 := v0[663]], [!!v0[663], v0[663]], [v1, v1, v1] + v112, v112 + v112];
				v114[198] := v112 + v112;
				var v115: seq<int> := [v104[758], |v112|, v106];
				var v116: map<seq<int>, int> := map[v115 := v104[758]];
				v0[663], v104[758], v114[198] := !(v106 == fm20(v1, globalState)), if (v115 in v116) then v116[v115] else v104[758], [v1, v0[663]] + [false];
				var v117: C0 := new C0(v105);
				var v118: map<C0, bool> := map[v117 := !v0[663]];
				v118 := v118;
				var v119 := new C0(fm25(v108, v105, globalState) + v105);
				var v120: set<bool> := {v1, !v0[663]};
				var v121: seq<seq<int>> := [[|v120|, 0x2ff]];
				v121 := (v121 + v121)[v106 := seq(0x316, i7  => (|v109|))];
				var v122: map<char, int> := map[v101 := fm20(v0[663], globalState)];
				v122 := v122[v101 := p0];
			} else {
				var v123: map<bool, array<bool>> := map[v1 := v0];
				var v124: set<array<bool>> := {v0, if (!v0[663] in v123) then v123[!v0[663]] else v0};
				var v125: array<set<array<bool>>> := new set<array<bool>>[4] [v124, v124, {v0, v0}, v124 + {v0}];
				var v126 := DC20(v124);
				v125[434] := if (v1) then v126.cf36 else v124;
				var v127: map<bool, set<array<bool>>> := map[false := v124 + v124];
				var v128: map<bool, bool> := map[v1 := v1];
				v125[434] := if ((if (v1) then if (v1 in v128) then v128[v1] else false else !v0[663]) in v127) then v127[if (v1) then if (v1 in v128) then v128[v1] else false else !v0[663]] else v124 - v124;
				var v129 := new C0(seq(0x1bd, i8  => (v101)));
				v104[758] := |multiset{v0[663] && v0[663], fm5(v106, v0[663], v103, false, globalState)}|;
				var v130: multiset<int> := multiset{v106, v106};
				var v131 := DC9(v130, v101, v106, {v129}, v106);
				var v132 := new C0(v105 + v129.f8[v104[758] := v131.cf15]);
				v104[758] := |map[!v0[663] := p0]|;
			}
			
			v104[758] := v106;
			v106 := p0;
		} else {
			r0 := !v1;
			v1 := if (v0[663]) then v0[663] else !v0[663];
			var v133 := 0x3a;
			v133 := p0;
			var v134 := "gv";
			var v135 := new C0(v134 + (seq(-0x364, i9  => (v101))));
			v0[663] := v1;
		}
		
		var v136 := "ng";
		var v137: seq<bool> := [v1, v0[663]];
		var v138: seq<seq<bool>> := [v137, v137];
		var v139: seq<seq<bool>> := [v137, v138[p0]];
		var v140 := DC15(p0, |v136|, [fm23(v1, globalState), v137] == v138);
		match (v140) {
			case DC15(cf30, cf31, cf32) =>
				v102 := DC1(cf31 - 894);
				v137 := v137;
				v0[663] := cf32;
				var v141: array<int> := new int[2];
				v141[541] := -cf30;
				v141[541] := 0x14f;
			case DC14(cf29) =>
				var v142 := 0x134;
				v142 := -p0;
				var v143: C0 := new C0(cf29);
				v143, v142, v0[663] := v143, v142, v0[663];
				v140, v142, v143 := v140, v142 + |v100|, v143;
				var v144: array<T0> := new T0[8];
				var v145: multiset<bool> := multiset{v1};
				var v146 := DC21(v145);
				var v147: array<int> := new int[23] [|map[v1 := p0]|, v102.cf2, p0, -p0, p0, v142, p0, p0, -p0, p0, v142, -|(v146.(cf37 := multiset(v137))).cf37|, v142, p0, v142, v142, p0, v142, p0, 70, p0, p0, v142];
				var v148: array<array<int>> := new array<int>[8] [v147, v147, v147, v147, v147, v147, v147, v147];
				var v149: map<array<T0>, array<array<int>>> := map[v144 := v148];
				var v150: seq<array<T0>> := [v144];
				var v151: seq<array<array<int>>> := [if (v150[v142] in v149) then v149[v150[v142]] else v148];
				var v152: array<array<array<int>>> := new array<array<int>>[28] [v151[p0], v148, v148, v148, if (v0[663]) then v148 else v148, v148, v148, v148, v148, v148, v148, v148, v148, v148, v148, v148, v148, v148, v148, v148, v148, v148, v148, if (true) then v148 else v148, if (v1) then v148 else v148, v148, v148, v148];
				v152[766] := v148;
				v147[334] := |(v100 + v100)|;
				var v153: multiset<int> := multiset{fm20(true, globalState)};
				var v154: multiset<multiset<int>> := multiset{v153, v153};
				v152[766], v0[663], v142, v147[334], r0 := v148, v154 >= multiset{v153, v153, fm24(v0[663], globalState)}, if (fm5(p0, v0[663], v103, v1, globalState)) then 0x1f6 else p0, v142, v1;
			case DC16(cf33) =>
				var v155: map<int, bool> := map[p0 := !v0[663]];
				var v156: seq<int> := [p0, p0, p0, |v155|];
				var v157: map<multiset<int>, int> := map[multiset(v156) := fm20(v0[663], globalState)];
				var v158: set<bool> := {v0[663], v137[994], v1};
				var v159: array<int> := new int[29] [p0, p0 * p0, -p0, p0, --p0, p0, -p0, fm20(v1, globalState), p0, |[v100, v100]|, fm20(!false, globalState), p0, p0, |[v1, false, v0[663]]| / |v157|, p0, |(v158 + v158)|, |v137| / p0, p0, p0, 0xc7 * p0, -p0 + p0, p0, p0, if (v0[663]) then |v137| else p0, -208, p0, fm20(fm5(p0, v1, v103, v1, globalState), globalState), -0x196, p0];
				v159[691] := -p0;
				v159[112] := p0;
				var v160: seq<string> := [v136];
				v159[691], v159[112] := |(v160[p0] + "wqfk")| + p0, p0;
				var v161: multiset<bool> := multiset{v0[663], v1};
				var v162: map<string, multiset<bool>> := map[v136 := v161];
				var v163: multiset<map<string, multiset<bool>>> := multiset{v162, v162};
				v159[691] := |[|v161|, (if (map[v136 := v161] in v163) then v163[map[v136 := v161]] else p0) + v159[691], p0]|;
				v159[691] := 0x1d4;
				var v164: map<bool, string> := map[v1 := v136];
				var v165, v166 := m10(|v164| + p0, v159[691], !v0[663], v0, globalState);
		}
		
		var v167: array<multiset<bool>> := new multiset<bool>[20](i10 => multiset{true});
		v167 := new multiset<bool>[20];
		r0 := v0[663] ==> fm5(p0, v0[663], v103, false, globalState);
		var v168: map<int, bool> := map[p0 := v0[663]];
		r1 := v168 == v168;
	}
	method m2(p0: int, globalState: GlobalState) returns (r0: map<bool, multiset<int>>) {
		var v0 := false;
		var v1 := DC14(seq(0x1eb, i0  => ('g')));
		var v2 := new C0(if (v0) then v1.cf29 else "gpmds");
		var v3: map<int, bool> := map[|(seq(751, i1  => (p0)))| := v0];
		var v4: multiset<bool> := multiset{v0};
		var v5: multiset<int> := multiset{p0, p0, p0};
		var v6: set<int> := {p0, p0};
		var v7: map<multiset<int>, int> := map[v5 := |v6|];
		var v8: map<int, int> := map[|v7| := 0x33c];
		var v9: map<bool, int> := map[false := |v8|];
		var v10: array<int> := new int[20] [p0, |v2.f8|, 0x3a0, p0, p0, p0, p0, p0, p0, |v4|, p0, p0, |v9|, p0, p0, p0, p0, -p0, p0, p0];
		var v11: map<array<int>, int> := map[v10 := p0];
		var v12: map<map<array<int>, int>, array<int>> := map[v11 := v10];
		var v13: array<array<int>> := new array<int>[15] [v10, v10, if (v11 in v12) then v12[v11] else v10, v10, v10, v10, v10, v10, v10, v10, v10, v10, v10, v10, v10];
		var v14 := DC4(p0, v0, v13);
		var v15: seq<int> := [if (v0 in v9) then v9[v0] else p0, |"uotirlywr"|];
		var v16 := DC23(v10);
		var v17: map<int, array<int>> := map[p0 := v16.cf39];
		var v18: seq<C0> := [v2];
		var v19: array<int> := new int[19] [|v2.f8|, p0, |v3|, p0 + p0, p0 + p0, -p0, p0, p0 * p0, |[fm20(v0, globalState)]|, -(85 + |v2.f8|), p0, p0, v14.cf5, p0, v15[p0], p0, -|v17|, |v18| / p0, p0];
		var v20: multiset<string> := multiset{v1.cf29, v2.f8};
		v19 := new int[12] [v14.cf5, v15[p0], p0, -p0 % p0, p0, p0 + |v20|, p0, p0 + 0x4b, -(p0 + p0), p0 + p0, p0, p0 - p0];
		v10[655] := |v2.f8| % p0;
		v10[655] := p0;
		var i2 := 0;
		while (!false)
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			var v21 := DC10(p0, p0, v10[655]);
			var v22 := 't';
			var v23: map<string, char> := map[fm25(v21, v2.f8, globalState) := v22];
			var v24: map<char, map<string, char>> := map[v22 := v23];
			var v25: seq<multiset<bool>> := [v4];
			var v27: set<seq<char>> := {v2.f8};
			var v28: map<int, D2> := map[-p0 := DC3(v27)];
			var v29: map<string, map<int, D2>> := map[v2.f8 := v28];
			var v30: set<bool> := {v0};
			var v31: map<int, seq<bool>> := map[p0 := [v0]];
			var v32: map<C0, int> := map[v2 := p0];
			var v33: seq<bool> := [v0, v0];
			var v34 := DC1(v10[655]);
			var v37: map<string, bool> := map[v2.f8 := v0];
			var v40: map<string, int> := map[fm25(v21, v2.f8, globalState) := v10[655]];
			var v41 := DC27(v30);
			var v42 := DC24(fm26(v41.cf45, v33, v34, v22, globalState));
			var v43: array<map<string, char>> := new map<string, char>[29] [v23, map[v2.f8 := v22], v23, v23, v23, v23[v2.f8 := v22], if (v22 in v24) then v24[v22] else v23, map[v2.fm15([|v25|], v0, globalState) := 's'], v23, map v26 : string | v26 in v29 :: (v26) := ('t'), v23, v23, v23, v23, fm26(v30, if ((if (v2 in v32) then v32[v2] else v10[655]) in v31) then v31[if (v2 in v32) then v32[v2] else v10[655]] else v33, v34, v22, globalState), map[v2.f8 := v2.fm14(globalState)] + (map v35 : string | v35 in v20 :: (v35) := (v22))["u" := v22], v23, v23 + map[v2.f8 := v2.f8[v10[655]]], map v36 : string | v36 in v37 :: (v36) := (v22), v23, v23, v23, map v38 : string | v38 in (map v39 : string | v39 in v40 :: (v39) := (v0)) :: (v38) := (v22), v23, v23, v42.cf40, map[seq(0x37a, i3  => (v22)) := v22], v23 + v23, v23];
			v43[622] := v23;
			v43[622] := map[seq(0x88, i4  => (v22)) := v22];
			var v44: array<map<multiset<D3>, map<bool, T0>>> := new map<multiset<D3>, map<bool, T0>>[21];
			var v45 := DC7(map[v22 := |v6|]);
			var v46: multiset<D3> := multiset{v45, v45};
			var v47: T0 := new C2();
			var v48: map<bool, T0> := map[true := v47];
			var v49: map<multiset<D3>, map<bool, T0>> := map[v46 := v48];
			v44[735] := if (v0) then v49 else v49;
			var v50: seq<string> := [v2.f8, v2.f8];
			var v51: map<bool, map<bool, T0>> := map[v0 := v48];
			v0, v0, v44[735], v10[655] := v0, "xone" < v50[p0], (if (v0) then v49 else map[multiset{v45} := v48]) + map[v46 := if (v0 in v51) then v51[v0] else v48], -(if (v3 != v3) then p0 + fm20(v0, globalState) else v10[655] + p0);
			var v52: map<char, D0> := map['c' := v34];
			var v53 := DC5(v10[655], v15, false);
			if (fm5(v10[655], !(v10[655] <= p0), v52, v53.cf10, globalState)) {
				var v54: array<bool> := new bool[12](i5 => true);
				v54 := v54;
				v13[975] := v19;
				v13[975] := v19;
				var v55, v56 := v47.m1(-(p0 % 5), globalState);
				var v57: map<array<bool>, bool> := map[v54 := true];
				v56 := v57 == v57;
				v10[655] := v10[655];
			} else {
				var v58 := DC17(v33);
				v10[655] := |v58.cf34|;
				v2.f8 := seq(0x144, i6  => (v22));
				var v59 := DC44(v10[655]);
				v59 := v59.(cf82 := v15[|v9|]);
				v37 := v37[v2.f8 := v0];
				v5 := multiset((seq(0x137, i7  => (-0x4c))) + v15) * (v5 + v5);
			}
			
			var v60 := DC21(DC21(v4).cf37);
			var v61: seq<D9> := [v60, v60, v60];
			v61 := v61 + v61;
		}
		var v62 := new C2();
		v10[655] := p0;
		var v63: map<bool, multiset<int>> := map[!v0 := v5];
		r0 := v63;
	}
	method m9(globalState: GlobalState) returns (r0: array<D0>, r1: array<array<int>>, r2: bool) {
		var v0 := 'x';
		var v1 := "jvypfkdr";
		var v2 := false;
		var v3 := -189;
		var v4: seq<bool> := [v2];
		var v5: multiset<bool> := multiset{true, v2};
		var v6 := DC1(v3);
		var v7: map<char, D0> := map[v0 := v6];
		var v8: set<bool> := {v2};
		var v9: array<bool> := new bool[28] [v0 in v1, v2, v0 !in (seq(0x157, i0  => (v0))), v1 == fm31(v3, v2, |v4|, globalState), v2, v5 !! multiset([fm5(|v1|, v2, v7, v2, globalState)]), v2, v2, v2, v3 >= v3, v2 && v2, v2, v2 && true, v2, v2, v2, v2, v8 !! {v2}, v2, if (v2) then !v2 else true, v2, !(if (v2) then v2 else true), false, v2, v3 < v3, v5 <= v5, v2, v2];
		v9[117] := v2;
		v9[117] := v2;
		var v10: seq<int> := [v3];
		var v11: seq<seq<int>> := [v10, seq(0xb9, i1  => (|v10|))];
		var v12: set<seq<int>> := {[v3], v11[v3]};
		var v13 := DC37(v12);
		v13 := v13;
		var v14: C2 := new C2();
		v14, v3, v3, v9[117] := v14, 0x2b1, v3, v9[117] ==> v9[117];
		var v15: map<int, bool> := map[v3 := v9[117]];
		v15 := v15[v3 := {v2, v2, v9[117]} == v8];
		for i2 := v3 - v3 to v3 {
			v1 := v1;
			r2, v2 := true, (v3 % v3) in v10;
			var v16: array<array<string>> := new array<string>[3];
			var v17: array<string> := new string[20] [v1, v1, v1, v1, v1[i2 := v0], "oe", v1, v1, v1, v1, v1, v1, v1, v1, "diut", v1, v1, v1, v1, v1];
			v16[356] := v17;
			var v18: seq<array<string>> := [v17];
			v16[356] := v18[v3];
			var v19: set<array<bool>> := {v9};
			v9[117] := v19 > v19;
		}
		var v20: set<array<bool>> := {v9, v9, v9};
		var v21 := DC20(v20);
		var v22 := DC20(v21.cf36 - v20);
		match (v21) {
			case DC20(cf36) =>
				v3 := -((v3 / v3) % (v3 / 317));
				var v23: array<int> := new int[3];
				v23 := v23;
				var v25: map<char, int> := map[v0 := v3];
				var v27: set<char> := {v0};
				var v28: map<char, seq<bool>> := map[v0 := [v9[117], v9[117]]];
				var v30 := DC14(v1);
				var v31: map<char, D6> := map[v0 := v30];
				var v32: map<int, map<char, seq<bool>>> := map[v3 := v28];
				var v33: map<int, int> := map[v3 := v3];
				var v35: array<map<char, seq<bool>>> := new map<char, seq<bool>>[23] [map v24 : char | v24 in v25 :: (v24) := (v4), map v26 : char | v26 in v27 :: (v26) := (v4), if (v9[117]) then map[v0 := v4] else v28, v28, map v29 : char | v29 in v31 :: (v29) := (v4), if ((if (v3 in v33) then v33[v3] else v3) in v32) then v32[if (v3 in v33) then v33[v3] else v3] else v28, v28, v28, v28, map[v0 := v4], v28, v28, v28, v28, v28, map[fm51(false, v3, v3, globalState) := [!v9[117]]], v28 + v28, v28 + v28, map[v0 := v4], v28, v28, map[v0 := v4], (map v34 : char | v34 in [v0] :: (v34) := (v4)) + v28[fm51(v9[117], v3, |[v3]|, globalState) := v4]];
				v35[60] := map v36 : char | v36 in v28 :: (v36) := ([v2, v9[117]]);
				v35[60] := (v28 + (map v37 : char | v37 in v1 :: (v37) := (v4))) + v28;
				var v38: array<map<char, int>> := new map<char, int>[29];
				var v39 := new C1(|(v1 + v1)|, v38, v10, v6);
		}
		
		var v40: array<D0> := new D0[3](i3 => v6);
		r0 := v40;
		var v41: array<int> := new int[3](i4 => i4 / -v3);
		var v42 := DC32(v1, v2, |multiset{v2, false}|, v41);
		var v43: seq<string> := [v1];
		var v44: map<int, array<int>> := map[|v43[v3]| := v41];
		var v45: array<array<int>> := new array<int>[27] [v41, if (v2) then v41 else v41, v41, v41, v41, v41, v42.cf60, v41, v41, v41, v41, v41, v41, v41, v41, v41, v41, v41, v41, v41, v41, if (v3 in v44) then v44[v3] else v41, v41, v41, v41, v41, v41];
		r1 := v45;
		r2 := |v1| > 0x18c;
	}
	method m10(p0: int, p1: int, p2: bool, p3: array<bool>, globalState: GlobalState) returns (r0: bool, r1: int) {
		var v0 := 'u';
		var v1: map<char, int> := map[v0 := p1];
		var v2 := DC7(v1);
		v2 := v2.(cf12 := v1);
		for i0 := -0x193 to p0 {
			var v3: map<bool, int> := map[p2 := p1];
			var v4: multiset<int> := multiset{i0};
			var v5 := DC1(p1);
			var v6: map<char, D0> := map[v0 := v5];
			var v7: seq<bool> := [!fm5(i0, p2, v6, p2, globalState)];
			var v8: seq<int> := [|v3|];
			var v9: array<int> := new int[27];
			var v10: map<int, int> := map[|[v9]| := p1];
			var v11: seq<multiset<int>> := [v4];
			var v12: array<int> := new int[23] [-|(if (p2) then v3 else v3)|, |v4|, p0 * |v7|, p0, |v8|, 0xcd, i0, -p0, p0, p1, p1, p1, p1, -0x33c, if (-i0 in v10) then v10[-i0] else fm20(false, globalState), p0 / i0, 0x37c, i0, p1 * i0, p1, p1, i0, -|map[0x30b := |v11|]|];
			v12[740] := p1;
			v12[740] := -p1;
			v0 := v0;
			v12[740] := -p0;
			var v13: multiset<bool> := multiset{p2};
			r0 := (v13 * v13) >= v13;
		}
		var v14: array<multiset<int>> := new multiset<int>[27](i1 => multiset{-p0});
		var v15: multiset<int> := multiset{p0};
		v14[667] := v15;
		v14[667] := if (p2) then v15 else v15 - v15;
		var v16: seq<char> := [v0];
		var v17: array<map<char, int>> := new map<char, int>[3](i2 => v1);
		var v18: set<array<bool>> := {p3, p3, p3};
		var v19: seq<int> := [p0, |v18|, p0, p1, p1];
		var v20 := new C1(p1 + |v16|, v17, v19 + v19[p1 := 661], DC1(p1));
		var v21 := new C2();
		var v22 := DC1(v20.f9);
		var v23: seq<D0> := [v22];
		var v24: map<int, bool> := map[320 := v23[v20.f9 := v22] < v23];
		v24 := v24;
		r0 := p2;
		r1 := v20.f9;
	}
}

class C5 extends T1 {
	const f7 : array<int>
	constructor (f7 : array<int>, f0 : seq<int>, f1 : D0) {
		this.f7 := f7;
		this.f0 := f0;
		this.f1 := f1;
	}
	
	function fm6(p0: int, p1: int, globalState: GlobalState): int {
		0x1e7
	}
	method m7(p0: int, p1: map<bool, int>, p2: array<array<bool>>, p3: bool, globalState: GlobalState) returns (r0: int, r1: set<string>, r2: bool) {
		var v0 := "ibpiei";
		var v1 := 'c';
		var v2 := new C0((fm16(v0, p0, globalState))[p0 := v1]);
		if (p3) {
			var v3: set<bool> := {p3};
			f7[80] := f0[|v3|];
			var v4: map<int, int> := map[p0 := p0];
			f7[80] := (if (|"wsieh"| in v4) then v4[|"wsieh"|] else p0) - 333;
			if (fm17(globalState)) {
				r2 := p3;
				var v5: map<int, bool> := map[f7[80] := p3];
				var v6: seq<bool> := [p3, if (f7[80] in v5) then v5[f7[80]] else p3];
				var v7: array<bool> := new bool[26](i0 => p3);
				var v8: map<array<bool>, seq<bool>> := map[v7 := v6];
				v6 := v6 + (if (v7 in v8) then v8[v7] else v6);
				r0 := (-0x5d + -0x312) / f7[80];
				f7[80], r0 := p0, f7[80] - p0;
				r0 := f7[80];
			} else {
				f7[80] := f7[80] + p0;
				var v9: set<set<bool>> := {{p3}, v3};
				var v10: array<int> := new int[11];
				r2, v9, v10, r2, r2 := fm17(globalState), set v11 : set<bool> | v11 in fm18(globalState) :: (v11), v10, p3, !p3;
				r2 := p3;
				var v12: map<int, bool> := map[p0 := p3];
				v12 := v12[fm6(f7[80], p0, globalState) := p3];
				v0 := v2.f8 + (v0 + v2.f8);
			}
			
			f7[80] := p0;
			var v13: seq<seq<int>> := [f0];
			r2 := f0 !in v13;
			if (!p3) {
				v0 := v2.f8;
				f1 := f1;
				r0 := 0xca / p0;
				var v14: map<char, int> := map[v1 := p0];
				var v15 := DC7(v14);
				var v16: map<bool, bool> := map[false := p3];
				var v17: map<seq<bool>, bool> := map[fm19(fm17(globalState), p3, v15.cf12, p3, globalState) := if (p3 in v16) then v16[p3] else !false];
				v17 := (v17 + v17) + v17;
				var v18: multiset<int> := multiset{f7[80]};
				var v19: multiset<bool> := multiset{p3, p3, p3, !p3, p3};
				var v20: seq<bool> := [f7[80] >= f7[80], p3, v19 >= v19];
				var v21: T0 := new C4();
				var v22: seq<seq<bool>> := [v20];
				var v23: seq<seq<bool>> := [v22[p0]];
				r0, v18, v20, v21, f7[80] := f7[80], v18 - v18, v22[fm32(f7[80], globalState)], v21, f7[80];
			} else {
				var v24 := new C2();
				var v25: map<int, bool> := map[f7[80] := p3];
				var v26: seq<map<int, bool>> := [v25];
				var v27: map<map<int, bool>, int> := map[v26[f7[80]] := p0];
				var v28: map<char, D0> := map[v1 := f1];
				var v29: map<bool, bool> := map[p3 := v24.fm5(f7[80], false, v28, p3, globalState)];
				var v30: set<seq<int>> := {f0, seq(-806, i1  => (f7[80]))};
				var v31: multiset<bool> := multiset{p3, true, p3, p3};
				f7[80], v0, r2, f7[80], r2 := (p0 + 0xae) - (if (map[f7[80] := p3] in v27) then v27[map[f7[80] := p3]] else |v29|), v2.f8, !(if (p3) then multiset{v1, v1} >= fm58(DC42(f7[80], p0, p0), true, v30, globalState) else p3), |v31| * p0, p3;
				f7[80] := f7[80];
				r2 := v0 != v2.f8;
				var v32: array<bool> := new bool[15](i2 => p3 || false);
				var v33: map<char, int> := map[v1 := 583];
				var v34: set<map<char, int>> := {v33, v33};
				var v36: map<char, bool> := map[v1 := true];
				v32[536] := v34 == {v33, (map v35 : char | v35 in v36 :: (v35) := (f7[80]))[v1 := fm6(f7[80], f7[80], globalState)]};
				v32[536] := p3;
			}
			
		} else {
			r2 := p0 >= p0;
			var v37: seq<string> := [v2.f8];
			var v38: seq<string> := [v37[p0]];
			var v39: map<int, seq<int>> := map[|multiset([p0, p0])| - |v37| := [f0[|f0|], p0, p0, p0, |p1|] + f0];
			v39 := v39[p0 := f0];
			r2 := p3;
			r2 := !(p0 != -p0);
			var v40: set<int> := {p0};
			v40 := v40;
		}
		
		var v41: seq<bool> := [p3];
		r2, r2, r0, r0 := p3, if (p3) then p3 else !(v41 < fm19(p3, !false, map[v1 := |v2.f8|], p3, globalState)), p0, -(p0 * p0);
		var v42: map<int, bool> := map[-p0 := p3];
		v42 := v42 + v42;
		var v43 := DC28(p3, p0);
		v43 := v43;
		for i3 := p0 to p0 {
			v2.f8 := v0;
			r0 := p0;
			var v44: map<bool, bool> := map[p3 := p3];
			v44 := v44[p0 < p0 := p3];
			var v45 := DC25(v1, |v2.f8|, true);
			var v46: array<bool> := new bool[21] [p3, p3, p3, p3, p3, p3, p3, p3, true, p3, p3, v45.cf43, p3, p3, p3, !p3, p3, p3, p3, true, p3];
			var v47: set<array<bool>> := {v46};
			var v48: multiset<int> := multiset{|v47|, i3, 0xef, p0, i3};
			var v49: set<bool> := {p3, v48 == v48};
			v49 := v49;
		}
		var v50: multiset<int> := multiset{|v0|, -p0};
		r0 := |(v50 + (v50 + v50))|;
		r1 := fm59(globalState);
		r2 := true;
	}
	method m8(p0: bool, p1: bool, p2: multiset<int>, p3: bool, globalState: GlobalState) returns (r0: bool, r1: map<string, bool>, r2: seq<int>) {
		var v0 := "fi";
		var v1 := new C0(v0);
		var v2 := 'j';
		var v3 := -0x97;
		var v4 := DC43("ekkqo", v2, v3);
		r0 := !match v4 {
			case DC42(cf76, cf77, cf78) => p1
			case DC43(cf79, cf80, cf81) => p1
			case DC44(cf82) => p1
			case DC41(cf75) => p3
		};
		for i0 := -|v0| to fm44(p0, globalState) {
			var v5 := new C0("m");
			v3 := |(if (p1) then v5.f8 else v1.f8 + (seq(0x182, i1  => (v2))))|;
			r0 := false;
			var v6: seq<bool> := [true];
			var v7 := DC15(|(seq(-0x344, i2  => (v2)))|, |v6|, p3);
			var v8: map<bool, D6> := map[p1 := v7];
			match (DC16(if (p3 in v8) then v8[p3] else v7)) {
				case DC15(cf30, cf31, cf32) =>
					r0, r0, v3, cf30 := (cf30 * cf30) > (cf31 * 263), true || p0, 0xbf - v3, -cf30;
					var v9: array<int> := new int[6];
					var v13: array<bool> := new bool[16](i3 => (set v10 : int | (0x37c <= v10) && (v10 < 0x334) :: (v10 - i0)) != (set v12 : int | v12 in (map v11 : int | v11 in {i0} :: (v11 / cf31) := (49)) :: (v12 / 0x319)));
					v13[311] := fm17(globalState);
					var v14: set<int> := {v3, -459};
					cf32, v3, v9, v13[311] := 0x1cb == |([cf32, !p0, cf32])[-|v14| := p3]|, v3, v9, p3;
					cf32 := (f0 + f0) !in ((([f0])[fm20(v13[311], globalState) := f0])[cf31 := [i0]] + [[v3], f0[v3 := |(seq(906, i4  => (v2)))|]]);
					r0 := v2 in (seq(-0x374, i5  => (v2)));
				case DC14(cf29) =>
					r2 := f0 + (f0 + f0)[f0[fm44(fm17(globalState), globalState)] := v3];
					var v15 := new C4();
					v2 := v2;
					var v16 := new C0(v0);
				case DC16(cf33) =>
					var v17: array<bool> := new bool[15];
					v17[971] := p1;
					v17[971] := p3;
					v3 := -0x218 + v3;
					v17[971] := -fm44(p1, globalState) in {-289};
					v17 := v17;
			}
			
		}
		var i6 := 0;
		while (false)
			decreases 100 - i6
		{
			if (i6 >= 100) {
				break;
			}
			
			i6 := i6 + 1;
			r0 := p0;
			if (!p0) {
				var v18: array<int> := new int[25];
				v18 := f7;
				var v19: set<int> := {v3};
				r0 := v19 > v19;
				v3 := -(fm44(p1, globalState) + v3);
				v18[12] := v3;
				v18[12] := fm20(p1, globalState);
				var v20: array<map<char, int>> := new map<char, int>[16](i8 => map[v2 := v3]);
				var v21 := new C1(|((seq(0x3c6, i7  => (v2))) + "skvawdoq")|, v20, f0, f1);
			} else {
				var v22: array<array<string>> := new array<string>[5];
				var v23: array<string> := new string[18];
				v22[170] := v23;
				v22[170] := v23;
				var v24: set<int> := {v3};
				var v25: map<set<int>, seq<int>> := map[v24 := (seq(0x198, i9  => (v3))) + f0];
				var v26: map<int, set<int>> := map[v3 := v24];
				var v27: map<int, bool> := map[v3 := p1];
				v25 := v25[if (|v27| in v26) then v26[|v27|] else set v28 : int | (-23 <= v28) && (v28 < -0x166) :: (v28 % v3) := f0];
				var v29: array<set<bool>> := new set<bool>[22];
				var v30: set<bool> := {p1, p0};
				v29[877] := v30;
				v29[877] := v30;
				var v31: map<bool, int> := map[false := v3];
				var v32: set<map<bool, int>> := {v31, v31};
				v32 := if (p3) then v32 else v32;
				f7[125] := 0x88;
				var v33 := DC25(v2, v3, !p1);
				var v34: map<D11, int> := map[v33 := -0x31b];
				var v36: map<D11, bool> := map[v33 := p1];
				var v37: multiset<map<D11, int>> := multiset{v34 + v34, v34 + (map v35 : D11 | v35 in v36 :: (v35) := (v3))};
				f7[125] := if ((fm60(v3, p1, globalState) + map[v33 := v3]) in v37) then v37[fm60(v3, p1, globalState) + map[v33 := v3]] else |v0|;
			}
			
			v3 := v3 - (v3 / v3);
			var v38: array<bool> := new bool[8];
			v38[153] := p0;
			v38[153] := if (true) then p1 else p3;
		}
		var v39: seq<bool> := [fm30(v3, p1, p1, globalState)];
		var v40: map<int, char> := map[|v39| := v2];
		v40 := v40[v3 + -v3 := if (true) then v2 else v2];
		f7[934] := v3 * 446;
		f7[934] := v3;
		var v41 := DC5(0x258, f0, false);
		r0 := v41.cf10;
		r1 := map["fhwdeveiq" := p1];
		r2 := f0;
	}
}

class C6 {
	const f6 : int
	constructor (f6 : int) {
		this.f6 := f6;
	}
	
	function fm11(p0: bool, globalState: GlobalState): int {
		261
	}
	function fm12(p0: bool, p1: bool, globalState: GlobalState): int {
		0x199
	}
	method m6(p0: D0, p1: string, p2: D0, p3: int, globalState: GlobalState) returns (r0: bool, r1: bool, r2: bool, r3: int) {
		match (DC1(p3).(cf2 := |{f6, |p1|}|)) {
			case DC0(cf0, cf1) =>
				var v0 := false;
				if (v0) {
					var v1: seq<bool> := [if (v0) then v0 else v0, v0, p1 <= p1, v0];
					v1 := v1;
					var v2: seq<int> := [cf1];
					r3 := v2[fm12(v0, v0, globalState)];
					cf1 := p3;
					var v3: multiset<int> := multiset{f6 * |v1|};
					var v4 := 'r';
					var v5: map<int, bool> := map[-f6 := v0];
					var v6: multiset<bool> := multiset{false};
					var v7: map<int, int> := map[|v6| := f6];
					var v8: seq<map<int, int>> := [v7, v7];
					v3 := (fm13(v4, !(if (p3 in v5) then v5[p3] else v0), globalState))[111 := p3 - |v8|];
					var v9 := new C2();
				} else {
					var v10: array<map<bool, bool>> := new map<bool, bool>[28];
					var v11: map<bool, bool> := map[!v0 := v0];
					v10[55] := v11 + v11;
					var v12: seq<bool> := [v0, v0];
					v10[55] := v11 + map[v12[f6] := v0];
					var v13: map<bool, int> := map[v0 := f6];
					v13 := v13[v12[0x396] := cf1];
					var v14: array<map<char, int>> := new map<char, int>[5];
					var v15 := 's';
					var v16: multiset<char> := multiset{'t', v15};
					var v17: seq<int> := [p3, -|v16|, f6];
					var v18: C1 := new C1(|v13|, v14, v17, p0);
					var v19: map<C1, int> := map[v18 := v18.f9];
					cf1 := fm12(v0, !v0, globalState) + (if (v18 in v19) then v19[v18] else f6);
					v18.f9 := v18.fm6(p3, cf1, globalState);
					var v20: array<bool> := new bool[3](i0 => |multiset{v18.f9, f6}| < DC11(|p1|, false, v0, v0, !v0).cf22);
					v20[276] := v0;
					v20[276] := v0;
				}
				
				var v21: map<D0, int> := map[p0 := f6];
				v21 := v21[p0 := 0x1fd];
				var v22 := 'i';
				var v23 := DC43(p1, 'u', -0x290);
				var v24: seq<string> := [p1, p1, v23.cf79];
				var v25: seq<string> := [v24[p3]];
				var v26: map<string, bool> := map[p1 := v0];
				var v27: map<seq<string>, map<string, bool>> := map[[p1, p1, ("nydykbd")[-926 := v22], p1, p1] + v25 := v26];
				v27 := v27[v24 := v26];
				var v28 := new C3();
			case DC1(cf2) =>
				var v29 := new C0("jvlu");
				var v30: map<int, bool> := map[0x1cd := true];
				var v31: map<int, bool> := map[f6 := !(if (207 in v30) then v30[207] else true)];
				var v32 := true;
				r1 := (v30 == v31) <==> v32;
				r1 := (v32 && v32) ==> v32;
				r2 := v32;
		}
		
		var v33 := true;
		var v34: multiset<bool> := multiset{v33};
		var v35: seq<multiset<bool>> := [v34, v34];
		var v36: map<int, bool> := map[p3 := v33];
		var v37: multiset<int> := multiset{f6};
		var v38: set<bool> := {v33};
		var v39: array<int> := new int[29] [|v35|, p3, f6, -178, p3 - p3, -|v36|, -f6 - -f6, f6 % 0x2fa, p3 - f6, f6, 0x2eb, -0x2cb, f6, f6 + |(v34[v33 := f6])[v33 := |v37|]|, --f6, p3, p3, f6, |v37|, f6, p3 + f6, |v38|, 0x61, -(p3 * p3), f6, f6, 0x128, f6, f6];
		forall i1 | 0 <= i1 < v39.Length {
			v39[i1] := i1 / f6;
		}
		v39[574] := f6;
		v39[574] := -((p3 + p3) / (f6 + -f6));
		var v40 := 'l';
		var v41 := DC38(v40);
		var v42: set<char> := {v41.cf71};
		var v43: seq<int> := [f6, |v42|];
		var i2 := 0;
		while (v43 <= v43)
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			v33 := v33 && (if (v33) then v33 else v33);
			if (v33) {
				var v44 := DC10(p3, f6, p3);
				var v45 := DC39(fm25(v44, p1, globalState) + p1, fm61(-243, f6, globalState));
				var v46: map<char, int> := map[v40 := f6];
				var v47: C0 := new C0(fm16(fm46(v33, p3, v33, globalState), v39[574], globalState));
				var v48: set<C0> := {v47};
				var v49 := DC9(v37, v40, p3, v48, v43[f6]);
				var v50: seq<map<char, int>> := [v46[v49.cf15 := p3]];
				var v51 := DC19(|v47.f8|);
				var v52: map<D7, int> := map[v51 := p3];
				var v53: map<map<D7, int>, bool> := map[v52 := v33];
				v45 := fm62(v50, if (v52 in v53) then v53[v52] else v33, f6, v39[574], globalState);
				var v54 := DC32(p1, true, f6, v39);
				v39 := v54.cf60;
				var v55: seq<bool> := [true];
				v33 := v55[|v55|];
				r2 := v33;
				r3 := (|fm50(v33, v33, |p1|, globalState)| * |v47.f8[v39[574] := v40]|) + v39[574];
			} else {
				var v56: array<array<bool>> := new array<bool>[6];
				var v57: seq<bool> := [v33, v33];
				var v58: array<bool> := new bool[18] [v33, v33, v33, true, fm17(globalState), !!v33, fm17(globalState), true, !false, true, v33, v33, v33, false, !v57[p3], v33, v33, !v33];
				v56[692] := v58;
				v56[692] := v58;
				v36 := v36[-p3 := v33 && fm30(v39[574], v57[-f6], !v33, globalState)];
				r0 := v33;
				var v59: map<char, int> := map[v40 := |v36|];
				v59 := v59[if (v33) then v40 else fm29(!v33, globalState) := |"yyicitvk"|];
				v39[574] := v39[574];
			}
			
			var v60: seq<bool> := [v33, v33, fm17(globalState), v33, v33];
			var v61 := DC28(v60[p3], p3);
			var v62: map<int, int> := map[f6 := p3];
			v39[574] := |([p3, p3] + [|p1|, |v60|, v61.cf47, |v62|])|;
			v36 := v36[p3 := v33];
		}
		if ((|p1| - v39[574]) < f6) {
			var v63 := DC29(!true, f6, v33, f6 / v39[574], !v33);
			v63 := DC29(v33, p3, v33, f6 * v39[574], !v33);
			var v64: seq<bool> := [v38 != v38];
			v64 := v64;
			var v65 := DC8([p2, p2]);
			var v66: map<bool, D4> := map[v33 := v65];
			var v67: seq<D0> := [p2, p2];
			v66 := v66[!v33 := DC8(v67)];
			var v68 := new C3();
			var v69 := "knsogq";
			v69 := p1;
		} else {
			var v70: array<bool> := new bool[12] [v33, fm17(globalState), fm30(v39[574], true, !v33, globalState), v33, true, false, f6 > -(if (0x110 in v37) then v37[0x110] else p3), v33, v33, v33, v33, v33];
			v70[61] := v33;
			v70[61] := v33;
			var v71 := "ydabfju";
			var v72: multiset<string> := multiset{p1[|v43| := v40], "ayhpyhe"};
			var v73: set<int> := {f6, f6};
			r2, r3, r3, v71, v70[61] := (if (v33) then v72 else v72) <= v72, f6, if (v33) then f6 else v39[574] % v39[574], "su" + v71, v73 > v73;
			var v74 := DC28(v33, p3);
			v39[574] := v74.cf47;
			r3 := -f6 % fm12(v70[61], v70[61], globalState);
			v33 := v70[61];
		}
		
		var v75: seq<string> := [("fm" + p1)[-|map[p3 := v33]| := fm51(v33, v39[574], p3, globalState)], p1, p1, p1, p1];
		v75 := v75;
		r0 := v40 in (seq(-0x2ea, i3  => (v40)));
		r1 := !fm17(globalState);
		r2 := v33;
		var v76: array<bool> := new bool[24](i4 => if (|[v33]| in v36) then v36[|[v33]|] else if (v39[574] in v36) then v36[v39[574]] else v33);
		var v77: set<array<bool>> := {v76};
		var v78 := DC20(v77);
		var v79: map<bool, D8> := map[v33 := v78];
		var v80 := DC2(v33);
		r3 := |(v79 + v79[v80.cf3 := v78.(cf36 := v77)])| / 93;
	}
}

class C7 extends T1 {
	const f4 : multiset<int>
	const f5 : D0
	constructor (f4 : multiset<int>, f5 : D0, f0 : seq<int>, f1 : D0) {
		this.f4 := f4;
		this.f5 := f5;
		this.f0 := f0;
		this.f1 := f1;
	}
	
	function fm6(p0: int, p1: int, globalState: GlobalState): int {
		-0x4c - (|"af"| - -575)
	}
	function fm8(p0: map<bool, int>, globalState: GlobalState): bool {
		if (!false && false) then !(|[false]| != |(seq(440, i0  => ('n')))|) else !!!!!!false <==> false
	}
	function fm9(p0: seq<D0>, p1: int, p2: map<int, set<bool>>, p3: seq<bool>, globalState: GlobalState): map<int, int> {
		(map[606 := -633] + (map v0 : int | (834 <= v0) && (v0 < 0x2c4) :: (v0 * 0x255) := (0x34))) + ((map v1 : int | v1 in map[0x149 := true] :: (v1 % 809) := (0x2e7)) + map[|(seq(312, i0  => ('j')))| := -|"dx"|])
	}
	method m5(p0: int, p1: array<int>, globalState: GlobalState) {
		var v0 := true;
		var i0 := 0;
		while (v0)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			v0 := v0;
			var v1: array<char> := new char[4];
			v1 := new char[12](i1 => 'c');
			var v2: array<seq<int>> := new seq<int>[8];
			v2[910] := f0;
			v2[910] := f0 + f0;
			var v3: array<string> := new seq<char>[25](i2 => seq(0x2e1, i3  => ('u')));
			v3 := v3;
		}
		var v4: array<set<bool>> := new set<bool>[28];
		var v5 := DC0(v4, 0x258);
		var v6: multiset<D0> := multiset{v5};
		var v7: set<int> := {fm6(|v6[v5 := p0]|, p0, globalState)};
		var v8: seq<set<int>> := [v7, v7, v7, v7, v7];
		if (!(v8[p0] >= fm10(true, globalState))) {
			v7 := {if (v0) then p0 else p0, p0};
			var v9: array<array<bool>> := new array<bool>[7];
			v9 := v9;
			var v10 := 'c';
			var v11: seq<char> := [v10, v10, v10, v10];
			var v12: set<seq<char>> := {v11 + v11};
			var v13 := DC3(v12);
			var v14: seq<seq<char>> := [v11];
			v12 := (v13.(cf4 := set v15 : seq<char> | v15 in v14 :: (v15))).cf4;
			var v16: array<array<int>> := new array<int>[3];
			var v17 := DC4(p0, v0, v16);
			if (v17.cf6) {
				var v18: map<char, string> := map[v10 := v11];
				var v19: multiset<string> := multiset{if (v10 in v18) then v18[v10] else v11, v11, v11 + v11, v11};
				v19 := if (v0) then v19 + v19 else v19;
				var v20 := 221;
				v20 := v20;
				var v21: array<map<char, int>> := new map<char, int>[19];
				var v22 := new C1(fm32(v20, globalState), v21, f0, fm49(v10, globalState));
				var v23 := new C3();
				var v24: map<int, int> := map[v20 := v20];
				var v25: map<int, int> := map[if (v20 in v24) then v24[v20] else -v20 := -v20];
				var v26: seq<bool> := [v0, v0];
				v25 := v25[p0 := |v26|];
			} else {
				var v27: map<bool, int> := map[v0 := p0];
				var v28: map<int, map<bool, int>> := map[p0 := v27];
				var v29: multiset<bool> := multiset{v0};
				v0 := fm63(if (p0 in v28) then v28[p0] else map[v0 := |map[!v0 := p0]|], p0, globalState) != v29;
				var v30 := 838;
				v30 := v30;
				v30 := v30;
				var v31 := DC28(false, v30);
				var v32: map<int, seq<int>> := map[v31.cf47 := f0];
				var v33: map<bool, bool> := map[v0 := false];
				v11 := (if (v0) then v11[|v32| := fm51(v0, |v33|, v30, globalState)] else v11 + [v10, v10, v10])[v30 % p0 := v10];
				var v34: seq<bool> := [false, v0, p0 != 0x119];
				v34 := v34 + v34;
			}
			
			var v35: array<map<char, int>> := new map<char, int>[1](i4 => map[v10 := p0]);
			var v36: C1 := new C1(p0, v35, f0, f5);
			var v37: set<C1> := {v36, v36, v36};
			var v38: map<int, int> := map[-p0 := |v37|];
			var v39: map<bool, int> := map[v0 := |v7|];
			var v40: map<char, int> := map[v10 := v36.f9];
			var v41: map<seq<bool>, int> := map[fm19(v0, v0, v40, false, globalState) := p0];
			var v42: seq<map<int, int>> := [v38, v38[|v39| := |v41|], v38, v38 + v38];
			var v43: map<seq<int>, seq<map<int, int>>> := map[seq(-0x397, i5  => (p0)) := v42];
			var v44: multiset<bool> := multiset{v0};
			var v45: seq<D0> := [f5, DC1(p0), f1, DC1(v36.f9), f5];
			v42 := ((if (p0 == v36.f9) then if (f0[|v44| := v36.f9] in v43) then v43[f0[|v44| := v36.f9]] else v42 else v42 + v42)[v36.fm6(p0, v36.f9, globalState) := v38])[0x245 := fm9(v45, p0, map[p0 := {true}], [true], globalState)];
		} else {
			var v46 := 0x34a;
			v46 := -p0 - v46;
			var v47 := DC10(v46 + p0, v46, v46);
			v47 := v47;
			var v48 := new C2();
			var v49 := DC44(-p0);
			var v50: T0 := new C2();
			v49, v50 := v49, if (v0) then v50 else v50;
			var v51 := "ru";
			var v52 := DC35(p0 / v46, if (true) then v51 else seq(-598, i6  => ('w')), if (v46 in f4) then f4[v46] else v46, true, p0 >= p0);
			match (v52) {
				case DC35(cf63, cf64, cf65, cf66, cf67) =>
					v0 := !(cf67 ==> (true <== cf66));
					var v53 := new C0(cf64);
					cf63 := f0[cf63];
					var v54 := 'm';
					v54 := v54;
				case DC36(cf68, cf69) =>
					v0 := v0;
					var v55: set<bool> := {v0};
					v0 := v55 !! {v0};
					var v56 := new C4();
					var v57: seq<bool> := [true, v0];
					var v58: map<set<int>, bool> := map[if (v57[if (p0 in cf68) then cf68[p0] else p0]) then v7 else {v46} := v0];
					var v59: map<bool, bool> := map[v0 := false];
					v58 := v58[{0x25a, fm20(!v0, globalState), v46, |v59|, p0} := v0];
				case DC34(cf62) =>
					p1[885] := v46;
					var v60 := 'y';
					var v61: array<bool> := new bool[15](i7 => v0);
					var v62 := DC18();
					var v63: map<bool, bool> := map[true := v0];
					v0, p1[885], v0, v60, v61 := !!!(f4 >= fm48(v62, |v63|, p0, p0, globalState)), v46, v0, v60, v61;
					var v64 := new C6(v46);
					v61[747] := v0;
					var v65: multiset<bool> := multiset{!v0};
					var v66: seq<string> := [v51];
					v61[747], p1[885], v61, v51 := (v65 + v65) == v65, v64.f6, v61, v66[p0];
					p1[885] := -0x2d3;
			}
			
		}
		
		for i8 := p0 + p0 to p0 {
			var v67: map<int, bool> := map[i8 := v0];
			var v68: map<int, map<int, bool>> := map[p0 := v67];
			var v69 := 611;
			v68, v69 := v68, p0 - (i8 + -p0);
			if (v0) {
				var v70 := "ashcuj";
				var v71 := new C0(v70);
				v69 := i8;
				var v72: array<bool> := new bool[8];
				v72[891] := v0;
				v72[891] := v0;
				v69 := p0 % v69;
				var v73: map<bool, int> := map[false := p0];
				var v74: seq<int> := [|(v73 + map[v72[891] := i8])|, v69];
				v74 := (seq(11, i9  => (i8)))[v69 := v69];
			} else {
				var v75 := 'm';
				var v76 := "unqvelq";
				var v77: C0 := new C0(v76);
				var v78: set<C0> := {v77, v77, v77};
				var v79: array<array<int>> := new array<int>[8];
				var v80 := DC4(|v77.f8|, v0, v79);
				var v81 := DC6(v80);
				var v82: map<int, D2> := map[|v76| := v81];
				var v83 := DC9(f4, v75, v69, v78, |v82|);
				var v84: map<bool, string> := map[v0 := v77.f8];
				var v85 := DC29(v0, p0, v0, |(if (v0 in v84) then v84[v0] else v76)|, v0);
				var v86: multiset<D12> := multiset{v85, v85, v85.(cf49 := |v77.f8|)};
				v0, v69, v83, v0 := v0, p0, DC9(f4, v75, if (v85 in v86) then v86[v85] else p0, v78, v69 / -i8), v0;
				v69 := v69;
				var v90 := DC45(multiset{f4});
				v76, v69, v69 := v76, |((map v87 : multiset<int> | v87 in (seq(-905, i10  => (f4))) :: (v87) := (v0)) + (map v88 : multiset<int> | v88 in (map v89 : multiset<int> | v89 in v90.cf83 :: (v89) := ({v0, v0})) :: (v88) := (v0)))| / p0, p0;
				var v91: seq<bool> := [if (|f0| in v67) then v67[|f0|] else v0, v0, v0];
				v0 := (if (v0) then fm6(v69, |fm24(v0, globalState)|, globalState) else |v91|) != p0;
				var v92: array<bool> := new bool[5](i11 => v77.f8[|v76|] !in "tvmhkbqyi");
				v92[953] := v0;
				var v93: array<array<bool>> := new array<bool>[3];
				v92[953], v0, v77.f8, v0, v93 := false, v0, "yyco", fm30(p0 % i8, v0, true && v0, globalState), v93;
			}
			
			var v94 := DC8([v5]);
			v94 := v94;
			v0 := v0;
		}
		var v95: map<bool, int> := map[v0 := p0];
		var v96: map<map<bool, int>, string> := map[v95 := "fakxfloyk"];
		var v97: array<map<map<bool, int>, string>> := new map<map<bool, int>, string>[3] [v96, v96 + v96, v96 + v96];
		var v98 := DC47(v96);
		v97[894] := v98.cf85;
		v97[894] := v96;
		var v99: map<bool, bool> := map[false := v0];
		var v100: map<int, int> := map[0x3d6 := p0];
		var v101: set<map<int, int>> := {map[p0 := |v99|], v100, v100};
		var v102: map<int, bool> := map[816 := v0];
		var v103: map<int, bool> := map[|{if (p0 in v102) then v102[p0] else v0}| := v0];
		v0, v0, v0 := v101 > v101, (p0 - -p0) in v103, if (v0 && v0) then v0 else true;
		var i12 := 0;
		while (v0)
			decreases 100 - i12
		{
			if (i12 >= 100) {
				break;
			}
			
			i12 := i12 + 1;
			var v104: seq<bool> := [v0];
			var v105 := DC17(v104);
			var v106 := DC11(p0, v105.cf34 != v104, true, !v0, v0);
			v106 := v106;
			var v107 := 0x18f;
			v107 := v107;
			var v108 := DC5(v107, f0, v0);
			var v109: T1 := new C5(p1, v108.cf9, f1);
			var v110 := DC29(v0, v107, v0, 331, false);
			var v111: seq<T1> := [v109, v109, v109, if (v110.cf52) then v109 else v109];
			v109 := v111[fm6(v107, v107, globalState) * v107];
			var v112 := new C2();
		}
	}
}

class C8 extends T0, T1 {
	const f2 : int
	const f3 : D0
	constructor (f2 : int, f3 : D0, f0 : seq<int>, f1 : D0) {
		this.f2 := f2;
		this.f3 := f3;
		this.f0 := f0;
		this.f1 := f1;
	}
	
	function fm5(p0: int, p1: bool, p2: map<char, D0>, p3: bool, globalState: GlobalState): bool {
		!(f0[f2] in multiset(seq(191, i0  => (f2))))
	}
	function fm6(p0: int, p1: int, globalState: GlobalState): int {
		if (|f0| >= f2) then f2 else |[map[false := f2], map[!false := f2], map[true := f2], map[false := f2]]|
	}
	function fm7(p0: int, globalState: GlobalState): string {
		"uyousyj" + "b"
	}
	method m1(p0: int, globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0 := 'd';
		var v1: map<char, int> := map[v0 := p0];
		var v2 := true;
		var v3 := "smiohn";
		var v5: map<char, bool> := map[v0 := v2];
		var v6: array<map<char, int>> := new map<char, int>[25] [v1[v0 := f2], v1, v1, map[v0 := 0x199], fm52(v2, DC38(fm51(true, |v3|, 0x217, globalState)), v2, globalState), v1, v1, map[v0 := 0xed], v1, v1, map[v0 := f2], v1, (map['g' := f2])[v0 := -p0], v1, v1, v1, v1, map[v0 := p0], v1, map v4 : char | v4 in v5 :: (v4) := (p0), v1, map[fm29(v2, globalState) := |(seq(565, i0  => (v0)))[p0 := v0]|], v1[v0 := f2], v1, v1];
		var v7 := new C1(f2, v6, if (v2) then f0 else f0, f1);
		var v8: array<int> := new int[7](i1 => i1 / p0);
		var v9: map<string, array<int>> := map[v3 := v8];
		var v10: map<array<int>, int> := map[if (v3 in v9) then v9[v3] else v8 := p0 + p0];
		v10 := v10[v8 := |v5|];
		var v11: seq<bool> := [v2, false];
		var v12: seq<seq<bool>> := [v11];
		v11 := (v11 + v12[p0]) + v11;
		var v13: map<int, bool> := map[p0 := v2];
		var v14: map<int, int> := map[|v13| := v7.f9];
		var v15: set<int> := {f2};
		var v16: map<bool, multiset<int>> := map[v2 := multiset{f2, |v14|, -|v15|}];
		var v17: multiset<int> := multiset{fm44(v2, globalState)};
		for i2 := |(if (v2 in v16) then v16[v2] else v17)| to fm44(v2, globalState) + p0 {
			v2 := false;
			if (v2) {
				var v18: array<string> := new string[3] [v3, v3, v3];
				v18[389] := v3;
				v18[389] := v3 + v3;
				var v19: map<bool, bool> := map[v2 := !v2];
				v7.f9 := -|multiset{map[false := true], map[v2 := true], (map[!v2 := v2])[true := v2] + v19, v19, v19}|;
				var v20: array<map<int, bool>> := new map<int, bool>[6] [v13, v13[v7.fm28(v7.f9, v2, globalState) := v2], v13, v13, v13, fm50(false, v2, i2, globalState)];
				var v21: array<array<map<int, bool>>> := new array<map<int, bool>>[8] [v20, v20, v20, v20, v20, v20, v20, v20];
				v21[775] := v20;
				v21[775] := v20;
				r0 := v2;
				var v22: set<bool> := {v2};
				var v23: map<bool, int> := map[v2 := |v22|];
				r0 := v2 !in v23;
			} else {
				var v24: array<array<bool>> := new array<bool>[24];
				var v25: array<bool> := new bool[23](i3 => v2);
				v24[748] := v25;
				var v26: array<D14> := new D14[2];
				var v27 := DC35(|(seq(0x276, i4  => (-v7.f9)))|, v3, i2, true, v2);
				v26[606] := v27.(cf67 := if (v7.f9 in v13) then v13[v7.f9] else v2, cf64 := seq(0x37c, i5  => (v0)));
				v25[581] := v2;
				v24[748], v26[606], v25[581], v7.f9, r1 := v25, v27, v2, f2, fm17(globalState);
				v7.f9 := if (v2) then |(fm13(v0, true, globalState) * v17)| else v7.f9;
				var v28, v29, v30 := m3(i2, f3, fm30(p0, v25[581], DC50(v25[581], v25, v2, v25[581], v3).cf90, globalState), globalState);
				v7.f9 := (0x32b - f0[i2]) % (if (v28 in v17) then v17[v28] else v7.f9);
				v7.f9 := fm6(i2, -(if (v2) then p0 else 0x36d), globalState);
			}
			
			var v31: array<char> := new char[26](i6 => v0);
			v31[149] := v0;
			v31[149] := v0;
			var v32: map<bool, string> := map[v2 := v3];
			var v33: seq<string> := [(if (v2 in v32) then v32[v2] else seq(0x20e, i7  => ('u')))[v7.f9 := v0]];
			v33 := [v3 + "chsubpa"];
		}
		var v34: map<bool, char> := map[v2 := v0];
		v34 := v34[false := v0];
		var v36: map<bool, int> := map[v2 := p0];
		var v37 := DC47(map v35 : map<bool, int> | v35 in [v36, v36] :: (v35) := ("w"));
		r1, r1 := match v37 {
			case DC48() => v2
			case DC49(cf86) => p0 > f2
			case DC50(cf87, cf88, cf89, cf90, cf91) => cf90
			case DC47(cf85) => v15 > v15
		}, !(p0 > f2);
		r0 := if (v2) then false else v2;
		r1 := fm17(globalState);
	}
	method m2(p0: int, globalState: GlobalState) returns (r0: map<bool, multiset<int>>) {
		var v0 := true;
		v0 := v0;
		var v1: array<D1> := new D1[20];
		v1[807] := DC2(v0);
		var v2 := DC2(v0);
		v1[807] := v2;
		if (p0 <= p0) {
			var v3 := 0xe;
			var v4: array<bool> := new bool[10];
			var v5: multiset<array<bool>> := multiset{v4};
			v4[23] := !!v0;
			v3, v5, v4[23], v0 := f2, v5, !(v3 == v3), !!v0;
			var v6: map<int, bool> := map[p0 := v4[23]];
			var v9: map<bool, bool> := map[v0 := v4[23]];
			var v10: set<map<int, bool>> := {v6, v6, map[|v9| := v0]};
			v4[23] := {v6, map v7 : int | (0x237 <= v7) && (v7 < -0xbb) :: (v7 - f2) := (v0), map v8 : int | (-0x38f <= v8) && (v8 < 126) :: (v8 - f2) := (!v4[23])} !! v10;
			var v11: map<int, int> := map[-|(v5 + v5)| := p0];
			v11 := v11[-793 + |{f2}| := f2];
			var v12 := "i";
			var v13: map<bool, string> := map[v3 >= v3 := v12];
			var v14 := 'u';
			v13 := v13[v0 := v12[f2 := v14] + v12];
			var v15: seq<bool> := [v0];
			var v16: multiset<int> := multiset{v3};
			var v17: set<int> := {|(seq(-438, i0  => (|(seq(0x1ec, i1  => (v14)))|)))|, v3, f2};
			var v18: set<int> := {|v17|, f2};
			var v19: array<int> := new int[27] [p0, v3, p0, f2, |[v3, p0]|, |v15|, v3, 0x124, f2, -p0, p0, f2, p0, p0, v3, f2, f2, p0, |v12|, |multiset(v15)|, v3, v3, |v11|, |v16|, |v17|, v3, v3];
			var v20: C5 := new C5(v19, DC5(p0, f0, v4[23]).cf9, f1.(cf2 := --v3));
			var v21: multiset<C5> := multiset{v20, v20};
			v21 := v21;
		} else {
			var v22 := 0x96;
			v22 := if (!true) then v22 else f2;
			var v23 := "uefbomho";
			var v24 := DC14(v23);
			match (v24) {
				case DC15(cf30, cf31, cf32) =>
					var v25: array<set<bool>> := new set<bool>[20](i2 => {!cf32, cf32});
					var v26: map<int, array<set<bool>>> := map[if (false) then v22 else v22 := v25];
					v26 := v26[v22 := v25];
					cf32 := true;
					var v27: array<bool> := new bool[6];
					v27[669] := v0;
					v27[9] := cf32;
					var v28: seq<bool> := [cf32, v0, cf32, cf32];
					v27[669], v27[9], v0, v0, v0 := if (f0[|v23|] >= p0) then v0 else v28[f2], v0, cf32, true, v0;
					var v29: multiset<bool> := multiset{v0};
					m4(v0, 105, v29, p0, globalState);
				case DC14(cf29) =>
					var v30: array<bool> := new bool[19];
					var v31 := DC50(!v0, v30, false, false, cf29);
					v30[206] := v31.cf87;
					v30[206] := v0;
					var v32 := new C3();
					var v33 := 'm';
					var v34: multiset<char> := multiset{v33, v33};
					var v35: seq<multiset<char>> := [v34];
					v30[206], v30[206], v0, v22 := v35[0x11f := v34] < (seq(0x7f, i3  => (v34))), v30[206], !((f2 > |cf29|) && !v30[206]), --((-fm6(f2, f2, globalState) / 0x341) - f2);
					var v36: set<char> := {'s'};
					var v37: set<int> := {v22, f2, 0x19};
					var v38: set<bool> := {v0};
					var v39: map<int, int> := map[-v22 := -f2];
					var v40: map<char, char> := map[v33 := v33];
					var v41: seq<map<char, char>> := [v40, v40];
					var v42: multiset<bool> := multiset{v0};
					var v43: seq<bool> := [false];
					var v44: array<int> := new int[25] [994, v22, f2, v22, 0x13a, |v38|, f2, v22, if (p0 in v39) then v39[p0] else fm44(v0, globalState), p0, |f0|, -|f0|, f2, |v41[-|v42|]|, |cf29|, v22, f2, |f0|, p0, v22, 0x10a, v22, |multiset(v43)|, f2, p0];
					var v45: seq<array<int>> := [v44, v44, v44];
					var v46: map<bool, int> := map[v30[206] := p0];
					var v47: map<array<int>, int> := map[v45[|v46|] := f2];
					var v48: multiset<int> := multiset{|fm25(fm64(globalState), (fm16(v23, |multiset{f2, |v43|, v22, f2, |v37|}|, globalState))[f2 := v33], globalState)|, v22};
					var v49: map<char, int> := map[v33 := p0];
					var v50: map<int, bool> := map[p0 := v30[206]];
					v36, v30[206], v37, v47, v22 := {v33, v33, v33, v33, v33}, v0, {|v48|, 12, 0x217, |v49|} - ((set v51 : int | v51 in v50 :: (v51 / |[|DC5(|[true, false, true, true]|, [|map[0x29f := -0x9]|], true).cf9|]|)) * v37), v47 + v47, -(p0 / (if (v30[206]) then v22 else f2));
				case DC16(cf33) =>
					var v52 := new C4();
					v0 := v0;
					var v53: array<seq<bool>> := new seq<bool>[20];
					var v54: seq<bool> := [true];
					v53[956] := v54;
					var v55 := 'g';
					var v56: map<char, seq<bool>> := map[v55 := v54];
					var v57 := DC38(v55);
					v53[956] := if (v57.cf71 in v56) then v56[v57.cf71] else (v54 + v54)[v22 := v0];
					var v58: array<int> := new int[4];
					v58[442] := f2;
					var v59: multiset<char> := multiset{'i'};
					v58[708] := |v59|;
					var v60: seq<array<int>> := [v58, v58, v58, v58];
					var v61: multiset<bool> := multiset{false};
					v58[442], v58[708] := |v60|, |v61|;
			}
			
			var v62, v63, v64 := m3(|map[|v23| := v0]|, f3.(cf1 := v22), true, globalState);
			var v65 := DC19(f2);
			match (v65) {
				case DC18() =>
					var v66: map<char, D0> := map['s' := f1];
					v0 := fm44(!fm5(f2, !v0, v66, v0, globalState), globalState) == -859;
					var v67: seq<map<int, bool>> := [map[v62 := v0]];
					var v68: seq<int> := [|v67[v62]|, -681];
					var v69: seq<bool> := [v0];
					var v70: array<int> := new int[20];
					var v71 := DC32(v23, v69[0xa], v22, v70);
					v0, v68 := v71.cf58, (v68 + v68) + v68;
					var v72: array<bool> := new bool[13];
					v72 := v72;
					v0 := v0;
				case DC19(cf35) =>
					v0 := (v23 + v23) <= ((seq(588, i4  => ('y'))) + v23);
					v64 := fm44(true, globalState);
					v0 := v0;
					v64 := f2;
				case DC17(cf34) =>
					var v73: multiset<int> := multiset{v62, v22, fm6(fm32(0x3de, globalState), p0, globalState), p0, v64};
					v0 := v73 <= (v73 - multiset{f2});
					var v74: array<int> := new int[13];
					v74[230] := 341;
					v74[230] := -|f0| - -0xcd;
					var v75 := 'x';
					var v76: multiset<char> := multiset{v75, v75};
					var v77, v78, v79 := m3(0x93, f3, v76 <= v76, globalState);
					var v80: map<bool, bool> := map[v0 := cf34[|v23|]];
					v0 := if (v0 in v80) then v80[v0] else v0;
			}
			
			var v81: seq<bool> := [v0, v0, fm17(globalState)];
			var v82: map<string, bool> := map[fm7(v64, globalState) := v0];
			var v83: map<int, bool> := map[v22 := v0];
			v81, v62 := [if (v23 in v82) then v82[v23] else if (v64 in v83) then v83[v64] else v0], v64;
		}
		
		var v84: array<bool> := new bool[5](i5 => v0);
		v84 := if (v0) then v84 else v84;
		var v85 := 844;
		v85 := -p0;
		var v86: seq<bool> := [v0, v0 ==> v0];
		var i6 := 0;
		while (v86[if (v0) then |v86| else |(seq(0x49, i7  => ('q')))|])
			decreases 100 - i6
		{
			if (i6 >= 100) {
				break;
			}
			
			i6 := i6 + 1;
			v85 := 0x2d4;
			if (!DC15(p0, fm20(v0, globalState), !v0).cf32) {
				var v87: seq<array<bool>> := [if (v0) then v84 else v84, v84, v84];
				v87 := v87;
				var v88: array<seq<bool>> := new seq<bool>[18](i8 => v86);
				v88[89] := v86;
				var v89 := DC44(f2);
				var v90: map<int, seq<bool>> := map[v89.cf82 := v86];
				var v91: multiset<seq<bool>> := multiset{v86};
				v88[89] := if (|v91[[v0, v0, false, !v0, fm17(globalState)] := v85]| in v90) then v90[|v91[[v0, v0, false, !v0, fm17(globalState)] := v85]|] else v86;
				v85 := p0;
				var v92: array<int> := new int[2](i9 => i9 / p0);
				var v93: T1 := new C5(v92, f0, f1);
				var v94: seq<T1> := [v93, v93];
				var v95: array<T1> := new T1[13] [v93, if (v0) then v93 else v93, v93, v93, v93, v93, v94[v85], v93, v93, v93, v93, v93, v93];
				v95[13] := v93;
				v95[13] := v93;
				var v96 := 'r';
				var v97: map<char, int> := map[v96 := v85];
				var v98: seq<seq<bool>> := [v86];
				var v99: set<seq<bool>> := {fm19(v0, v0, v97, !v0, globalState), v98[-f2], fm19(v0, v0, v97, v0, globalState)};
				v85 := |v99|;
			} else {
				v85 := p0;
				var v100: array<int> := new int[17];
				var v101 := new C5(v100, f0 + f0, DC1(v85));
				var v102 := "v";
				v102 := fm16(v102, -|v102|, globalState) + v102;
				var v103: map<bool, bool> := map[v0 := v0];
				var v104 := 'y';
				var v105: map<array<int>, map<bool, map<bool, int>>> := map[v101.f7 := map[v0 := fm65(|v103|, v104, f2, globalState)]];
				var v106: map<bool, int> := map[v0 := p0];
				var v107: map<bool, map<bool, int>> := map[true := v106];
				v105 := v105[v101.f7 := v107];
				var v108: set<int> := {f2};
				var v109: map<array<bool>, int> := map[v84 := v85];
				v108, v0, v85, v109, v85 := v108 + v108, fm17(globalState), 0x16b, v109 + v109, --(f2 - v101.fm6(f2, --447, globalState));
			}
			
			v85 := p0;
			v0 := !!v0;
		}
		var v110: array<C6> := new C6[7];
		var v111: map<bool, multiset<int>> := map[v0 := multiset{f2}];
		var v112: map<array<C6>, map<bool, multiset<int>>> := map[v110 := v111];
		var v113: multiset<int> := multiset{v85};
		r0 := (if (v110 in v112) then v112[v110] else map[v0 := v113]) + v111;
	}
	method m3(p0: int, p1: D0, p2: bool, globalState: GlobalState) returns (r0: int, r1: map<int, char>, r2: int) {
		var i0 := 0;
		while (!p2)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v0: multiset<bool> := multiset{p2};
			var v1: multiset<int> := multiset{f2, 0x3b5, f2, 0x12f, f2};
			var v2: map<bool, int> := map[!p2 := f2];
			var v3: map<bool, int> := map[p2 := -|v2|];
			var v4 := 'm';
			var v5: map<int, int> := map[-0x3ab := p0];
			var v6: map<int, bool> := map[981 := p2];
			var v7: map<bool, char> := map[p2 := 'm'];
			var v8: set<map<bool, char>> := {v7};
			var v9: array<bool> := new bool[2] [false, if (0x214 in v6) then v6[0x214] else p2];
			var v10: map<array<bool>, int> := map[v9 := p0];
			var v11: array<int> := new int[22] [|(seq(810, i1  => ('p')))|, |v0|, f2, -f2 * |v1|, if (p2) then |fm31(-0x12d, p2, p0, globalState)| else p0, p0, |v2|, f2 + f2, 0x17, |[p2, fm5(f2, p2, map[v4 := f1], p2, globalState)]|, f2, -|[f2, -(if (-f2 in v5) then v5[-f2] else f2), f2, |v6|]|, f2, p0 + |map[|v8| := f2]|, 0x1f6, fm6(-f2, f2, globalState), f2, |v10| - fm44(p2, globalState), p0, if (!p2 in v2) then v2[!p2] else 0x320, 0x4, p0];
			v11[854] := p0;
			v11[854] := p0 + p0;
			var v12: seq<bool> := [p2, p2];
			var v13: map<seq<bool>, bool> := map[DC17(v12).cf34 := p2];
			r0 := f0[|v13|];
			var v14 := true;
			var v15 := "hivvd";
			v14 := fm5(|(v15 + v15)|, if (true) then v14 else p2, map[v4 := f1], |v15| >= v11[854], globalState);
			var v16 := new C6(|f0| % p0);
		}
		var i2 := 0;
		while (642 <= (f2 % p0))
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			if (p2) {
				var v17: array<array<int>> := new array<int>[23];
				v17 := v17;
				var v18 := "xkbca";
				var v19 := 'h';
				var v20: map<string, int> := map[v18 + ("snuobpymy")[f2 := v19] := f2];
				v20 := v20[v18 := 0x1cb / -p0];
				var v21: map<int, int> := map[f2 := f2];
				var v22 := DC31(true, p0, f2);
				var v23: map<bool, bool> := map[false := true];
				var v24: array<int> := new int[15] [f2, |(if (p2) then v18[p0 := v19] else "b")|, |v18|, |{p0, p0}| / p0, f2 % (fm66(|v21|, p2, globalState)).cf78, if (p2) then p0 else p0, v22.cf55, f2, p0, |v23|, p0, p0, |"gk"|, f2 - fm32(p0, globalState), p0];
				var v25: multiset<int> := multiset{f2};
				v24[102] := |v25|;
				var v26: array<bool> := new bool[25](i3 => p2);
				var v27: set<array<bool>> := {if (p2) then v26 else v26, v26, v26};
				var v28 := false;
				r2, v24[102], v27, v28 := f2 % p0, p0, v27, v28;
				v26[228] := p2 <== v28;
				v28, v26[228], v24[102] := !v28, v28, -v24[102];
				v24[102] := v24[102] + |(seq(141, i4  => (v19)))|;
			} else {
				var v29: multiset<bool> := multiset{p2, p2, p2, p2};
				var v30: multiset<int> := multiset{p0};
				var v31: set<int> := {f2, if (p2 in v29) then v29[p2] else if (f2 in v30) then v30[f2] else f2};
				var v32 := DC22(p0);
				var v33: map<set<int>, D9> := map[v31 := v32];
				v33 := v33;
				var v34: array<int> := new int[6];
				v34[39] := p0;
				v34[39] := -(p0 - f2);
				var v35 := "ovve";
				v35 := v35;
				var v36 := false;
				v36 := v36;
				v36 := (p0 / v34[39]) != |(if (v36) then f0 else seq(0x184, i5  => (|map[-0xe1 := |f0|]|)))|;
			}
			
			var v37 := 'l';
			var v38 := DC25(v37, p0, p0 != 0x320);
			var v39: set<int> := {f2};
			var v40: map<int, bool> := map[f2 := p2];
			var v41: map<int, bool> := map[|v39| := if (p0 in v40) then v40[p0] else p2];
			var v42 := DC44(p0);
			var v43: set<int> := {|(v40 + map[p0 := !p2])|, f2 - f2, v42.cf82, p0};
			var v44 := "ntyajnp";
			var v45 := true;
			var v46: map<bool, bool> := map[!v45 := v45];
			var v47: seq<bool> := [p2, p2, p2, p2, p2];
			var v48: set<map<bool, bool>> := {map[p2 := p2], map[p2 := v45], map[v47[f2] := p2]};
			v38, v39, r0, v44, v45 := fm67({v46} * v48, globalState), {p0} + {|multiset{f2}|}, f2, v44, p2;
			var v49: set<string> := {v44, seq(-0x2e9, i6  => (v37))};
			var v50: map<seq<bool>, set<string>> := map[v47 := v49 - v49];
			v50 := v50[v47 := {v44}];
			var v51 := new C2();
		}
		var v52 := DC44(f2);
		r2 := match v52 {
			case DC42(cf76, cf77, cf78) => -|[p2]|
			case DC43(cf79, cf80, cf81) => |{map[true := p2], map[!p2 := p2], map[p2 := p2], map[p2 := true], map[p2 := p2]}|
			case DC44(cf82) => f2 / f2
			case DC41(cf75) => f2
		};
		var v53: array<bool> := new bool[7];
		v53[674] := p2;
		var v54 := true;
		v53[674], v54 := v54, !p2;
		for i7 := 0x3b1 - p0 to if (v53[674]) then --773 else p0 {
			v54 := if (p2) then p0 >= f2 else v53[674];
			var v55 := DC29(true, p0, p2, i7, v53[674]);
			v53[674] := if (if (p2) then v55.cf50 else v53[674]) then p2 else v53[674];
			var v56: array<seq<char>> := new seq<char>[25];
			v56 := v56;
			var v57 := new C2();
		}
		var v58 := DC28(true, f2);
		var v59: map<bool, bool> := map[!v53[674] := v58.cf46];
		var v61: seq<bool> := [v54];
		var v62: map<seq<bool>, bool> := map[v61 := p2];
		var v63: map<bool, map<seq<bool>, bool>> := map[v53[674] := map v60 : seq<bool> | v60 in v62 :: (v60) := (v53[674])];
		var v64: set<bool> := {v54};
		var v65: map<int, int> := map[p0 := p0];
		var v66: array<int> := new int[18](i8 => i8 / p0);
		var v67 := 'i';
		var v68: map<array<int>, map<bool, int>> := map[v66 := fm65(-f2, v67, p0, globalState)];
		var v69 := "haw";
		var v70: C0 := new C0(v69);
		var v71: seq<C0> := [v70];
		var v72: array<int> := new int[27] [f2 - p0, 0x3ad * f2, p0, f2, |f0|, f2 + f2, p0, f2, |[if (v53[674] in v59) then v59[v53[674]] else p2]| + 0x143, f2, |(if (!p2 in v63) then v63[!p2] else map[v61 := v54])| + f2, |(v64 - v64)|, f2, |v65| / |(if (v66 in v68) then v68[v66] else map[p2 := f2])|, f2 / 0x163, p0 * |v69|, p0 - |{v53[674], !v53[674], p2, p2}|, -f2, p0, if (v53[674]) then |v71| else f2, p0 * fm32(0x1f6, globalState), -(|v65| + p0), p0, f2, |f0|, |[f2]| - -0x1d5, |[p0]| / f2];
		var v73 := DC29(v54, |v61|, v53[674], 0xc4, p2);
		v72[644] := v73.cf49;
		v72[644] := (-p0 + 0x1bf) + 0x17f;
		var v74: map<int, array<bool>> := map[-v72[644] / f2 := v53];
		r0 := |v74|;
		var v75: map<int, char> := map[f0[v72[644]] := v67];
		r1 := v75 + map[|v69| := v67];
		r2 := v72[644] % p0;
	}
	method m4(p0: bool, p1: int, p2: multiset<bool>, p3: int, globalState: GlobalState) {
		var v0: array<bool> := new bool[23](i0 => false);
		var v1: array<array<bool>> := new array<bool>[7] [v0, v0, v0, v0, v0, v0, v0];
		var v2 := 'a';
		var v3: map<char, D0> := map[v2 := f1];
		var v4: map<bool, array<array<bool>>> := map[fm5(p3, !p0, v3, p0, globalState) := v1];
		v1 := if ((|f0| == p3) in v4) then v4[|f0| == p3] else v1;
		var v5: seq<bool> := [!p0, p0, p0];
		var v6: seq<seq<bool>> := [v5];
		if ([p0] !in ([v5] + v6)) {
			var v7: set<array<bool>> := {v0};
			var v8 := DC20(v7);
			v8 := v8;
			var v9 := false;
			v9 := v9;
			if (p0) {
				var v10 := 0xf9;
				v10 := p1;
				var v11: map<int, bool> := map[868 := !p0];
				var v12: map<array<bool>, bool> := map[v0 := !v9];
				var v13: array<map<array<bool>, bool>> := new map<array<bool>, bool>[26] [map[v0 := if (0x259 in v11) then v11[0x259] else v9], map[v0 := p0] + v12, v12, map[v0 := p0] + v12, if (p0) then v12 else map[v0 := p0], v12, v12 + map[v0 := v9], v12 + v12, map[v0 := p0] + v12, v12, v12 + v12, v12, v12, v12, v12 + v12, v12, v12, v12, v12, v12, v12, map[v0 := p0] + v12, v12, v12, v12, v12];
				v13[351] := v12;
				var v14: map<int, int> := map[p1 := -f2];
				var v15: array<int> := new int[28](i1 => i1 * v10);
				var v16: C5 := new C5(v15, [f2, p1], f1);
				var v17: array<C5> := new C5[21] [v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16];
				var v18: map<int, array<C5>> := map[p1 := v17];
				var v19: seq<array<C5>> := [if (p1 in v18) then v18[p1] else v17];
				var v20 := DC29(v9, p3, !p0, |v19|, fm30(v10, false, p0, globalState));
				v10, v9, v2, v9, v13[351] := (|v14| % 0x340) - v20.cf49, p0, v2, p0, (v12 + v12) + map[v0 := p0];
				var v21: seq<int> := [f2 - p3, p1, p1];
				v21 := f0 + v21;
				v9 := v9;
				v10 := |((v21 + v21) + f0)|;
			} else {
				var v22 := "tqhfxuca";
				var v23: array<int> := new int[4] [p1, f2, f2, |v22|];
				var v24: C5 := new C5(v23, seq(0x38e, i2  => (p3)), f1);
				var v25 := DC51(v24);
				var v26: set<C5> := {v25.cf92, v24};
				var v27: map<bool, bool> := map[p0 := v5[f2]];
				var v28: map<set<C5>, int> := map[v26 + v26 := |(v27 + v27)|];
				v28 := v28[v26 := f2];
				v9, v9 := if (v9) then p0 else v9, p0;
				var v29: map<int, string> := map[-f0[p3] := v22[f2 := v2]];
				var v30 := DC34(v29);
				var v32: map<int, bool> := map[f2 := p0];
				var v33: map<int, char> := map[p1 := v2];
				var v35: set<int> := {p3, p1};
				var v36: array<D14> := new D14[23] [v30, DC34(map v31 : int | v31 in v32 :: (v31 * p3) := (v22)), DC34(v29).(cf62 := map[p1 := v22]), v30, v30, v30, fm68(p0, v22, globalState), fm68(!p0, v22, globalState), v30, DC34(map[-|v33| := v22]), DC34(map v34 : int | v34 in v35 :: (v34 % p3) := (seq(-0x247, i3  => (v2)))), v30, v30, fm68(p0, v22, globalState), v30, v30, v30, v30.(cf62 := v29), v30, v30, v30, v30, v30.(cf62 := v29)];
				v36[751] := v30;
				v36[751] := v30;
				v32 := (fm21(globalState))[p1 := v9];
				v9 := p0;
			}
			
			var v37 := new C3();
			var v38: array<D2> := new D2[26];
			var v39: seq<char> := [v2];
			var v40: multiset<seq<char>> := multiset{v39};
			var v42 := DC3(set v41 : seq<char> | v41 in v40 :: (v41));
			v38[580] := v42.(cf4 := {v39});
			v38[580] := v42;
		} else {
			v0[116] := p0;
			var v43 := 0x3b6;
			var v44 := "wvwvfkb";
			var v45: array<int> := new int[1];
			v45[315] := |map[f0 := f2]| % 232;
			var v46: set<bool> := {false, p0};
			var v47: map<int, D10> := map[|v46| := DC23(v45)];
			v0[116], v43, v44, v45[315] := p0, fm20(!p0, globalState), v44[p3 := if (p0) then 'u' else 'q'], if (f2 in v47) then fm44(p0, globalState) / v43 else p3;
			if (v0[116]) {
				v45[315] := f2;
				v0[116] := !v0[116];
				v45[315] := -f0[f2];
				var v48: T0 := new C3();
				v48 := v48;
				v44 := v44 + (v44 + "oo");
			} else {
				v0[116] := -(|v44| / v43) == (if (true) then f2 else fm20(fm17(globalState), globalState));
				var v49: map<int, seq<int>> := map[v43 := seq(-0x365, i4  => (0x21a))];
				v45[315] := if (v0[116]) then f2 else |(if (p1 in v49) then v49[p1] else f0)| / |{f2}|;
				var v50: map<int, int> := map[|v5| := 0x3a2];
				var v51: multiset<int> := multiset{|v50|};
				var v52: multiset<int> := multiset{v45[315], if (-461 in v50) then v50[-461] else f2, p3, f2, |v51| + p3};
				v45[315] := if ((0x3a8 * p3) in v51) then v51[0x3a8 * p3] else fm44(p0, globalState);
				v45[315] := f0[0x194];
				var v53: C4 := new C4();
				var v54: map<bool, C4> := map[p0 := v53];
				v54 := v54[v0[116] := v53];
			}
			
			var v55: map<int, char> := map[p3 := v2];
			v2 := if (p1 in v55) then v55[p1] else v2;
			var v56: map<bool, int> := map[p0 := v45[315]];
			var v57: map<int, map<bool, int>> := map[0x12f := v56];
			if (|v57[p3 := v56]| != p3) {
				var v58 := new C2();
				var v59: set<C2> := {v58};
				var v60: array<array<int>> := new array<int>[15];
				var v61: map<set<C2>, array<array<int>>> := map[v59 + v59 := v60];
				v61 := v61[{v58} := v60];
				v43 := v43;
				var v62: seq<int> := [f2];
				v62 := DC5(v43, v62, false).cf9;
				var v63: array<char> := new char[5](i5 => v2);
				var v64: seq<array<char>> := [v63];
				v64 := v64;
			} else {
				v45[315] := if (p0) then v43 else v45[315];
				v44 := v44;
				var v65: seq<int> := [v45[315]];
				v65 := v65;
				v2 := v44[p3 % |multiset{v44}|];
				var v66: array<array<map<int, int>>> := new array<map<int, int>>[25];
				var v67: array<map<int, int>> := new map<int, int>[29];
				v66[29] := if (v5[p3]) then v67 else v67;
				var v68: set<int> := {|v44|};
				var v69: map<int, bool> := map[|[v45[315]]| := v0[116]];
				var v70: map<set<int>, map<int, bool>> := map[{|v44|, v45[315], v45[315], v45[315], p1} := v69];
				v66[29], v43, v0[116] := v67, f2, v68 !in (v70 + v70);
			}
			
			var v71 := DC32(v44, p0, v43, v45);
			var v72: multiset<string> := multiset{v44, v71.cf57};
			v0[116] := fm5(p3, p0, v3, v45[315] < (if ("cu" in v72) then v72["cu"] else v45[315]), globalState);
		}
		
		var v73: array<int> := new int[28];
		v73[772] := p3;
		v73[772] := p1 * f2;
		v0[846] := p0;
		var v74: set<seq<int>> := {f0, [p3]};
		var v75 := DC37(v74);
		var v76: map<int, D15> := map[v73[772] := v75];
		var v77: array<char> := new char[12];
		var v78 := DC52(v77);
		var v79: array<array<char>> := new array<char>[22] [v77, v77, v77, v77, v77, v77, v77, v77, v77, v77, v77, v78.cf93, v77, v77, v77, v77, v77, v78.cf93, v77, v77, v77, v77];
		v79[419] := v77;
		var v81: map<set<int>, map<int, D15>> := map[{f2, v73[772]} := fm69(false, globalState)];
		var v82: set<int> := {--p3};
		var v83: map<bool, int> := map[p0 := f2];
		v0[846], v76, v79[419], v73[772] := p0, (map v80 : int | (0x137 <= v80) && (v80 < 702) :: (v80 - -843) := (v75)) + (if (v82 in v81) then v81[v82] else v76), v77, f0[if (false in v83) then v83[false] else f2];
		v2 := 'x';
		var v84 := new C4();
	}
}

class C9 {
	constructor () {
	}
	
	function fm0(p0: bool, p1: int, globalState: GlobalState): int {
		0x3f
	}
	method m0(p0: int, globalState: GlobalState) returns (r0: int, r1: int, r2: seq<bool>) {
		var v0 := false;
		var v1: multiset<bool> := multiset{v0};
		r0 := fm0(v0, |v1|, globalState);
		var v2 := 'q';
		var v3: seq<char> := ['w', v2];
		var v4: seq<seq<char>> := [v3, v3, seq(-0xf9, i0  => ('q')), (seq(0x18e, i1  => (v2))) + v3, [v2]];
		v4 := v4;
		var v5: multiset<char> := multiset{v2};
		var v6: array<multiset<char>> := new multiset<char>[1] [v5];
		forall i2 | 0 <= i2 < v6.Length {
			v6[i2] := v5;
		}
		var v7: set<bool> := {v0};
		var i3 := 0;
		while (v7 > v7)
			decreases 100 - i3
		{
			if (i3 >= 100) {
				break;
			}
			
			i3 := i3 + 1;
			var v8: map<bool, int> := map[v7 !! v7 := p0];
			var v9: seq<bool> := [true];
			var v10: array<bool> := new bool[13] [v0, v0, false, v0, v0, !!v0, v0, v0, v0, v0, v9[|v3|], v0, v0];
			var v11: set<array<bool>> := {v10};
			v8 := v8[v0 := -|(v11 - v11)|];
			r1 := -p0;
			var v12: seq<int> := [p0 + p0];
			v12 := ([p0, |v3|, p0])[|fm1(p0, v0, p0, v0, globalState)| - fm0(false, p0, globalState) := p0];
			v0 := v0;
		}
		var v13: map<int, bool> := map[p0 := v0];
		v13 := v13;
		var v14: array<set<bool>> := new set<bool>[3](i4 => v7);
		var v15 := DC0(if (v0) then v14 else v14, p0);
		match (v15) {
			case DC0(cf0, cf1) =>
				r0 := -(|"moityagb"| / p0);
				var v16: seq<bool> := [v0, v0, true, v0];
				var v17: array<bool> := new bool[9] [!!v0 <== v0, v16[fm0(v0, p0, globalState)], v16[cf1], v16 != v16, if (v0) then v0 else v0, p0 >= -cf1, if (v0) then !v0 else v0, v0, !v0];
				v17[505] := v0;
				var v18: map<bool, bool> := map[v0 && v0 := v0];
				v17[505] := !(if (!v0 in v18) then v18[!v0] else v0);
				var v19: seq<set<bool>> := [v7, {v0}, v7, v7, {v17[505]}];
				var v20: set<int> := {p0, cf1, |v19|};
				r0, v3, v17[505], v0 := 0x238 + p0, v3, !v0, v20 < v20;
				var v21: array<array<string>> := new array<string>[10];
				var v22: array<string> := new string[18];
				v21[401] := v22;
				v21[401], v17[505], v4 := v22, v16[cf1], v4 + [seq(0x16b, i5  => (v2)), v3];
			case DC1(cf2) =>
				if (!!(if (v0) then fm2(DC2(v0).cf3, cf2, globalState) else v0)) {
					cf2 := -cf2 + p0;
					v0 := v0;
					var v23: seq<bool> := [v0];
					v0 := fm2(v23[p0], p0, globalState);
					v0 := fm2(v0, p0 % 0x33d, globalState);
					var v24: map<seq<bool>, char> := map[v23 := v2];
					v24 := fm3(v2, cf2 * -cf2, v0, v0, globalState);
				} else {
					var v25: array<int> := new int[8];
					var v26: map<bool, int> := map[v0 := cf2];
					v25[785] := if (true in v26) then v26[true] else |v3|;
					v25[785] := fm0(v0 && v0, |fm4(v3, v2, v0, globalState)| - p0, globalState);
					v3 := v3 + "rbslwhc";
					var v27: map<D1, char> := map[DC2(v0) := v2];
					var v28 := DC2(false);
					v27 := v27[v28 := 'o'];
					var v29: array<bool> := new bool[5] [v0 <== v0, true, v0, v0, true];
					v29[550] := false;
					var v30 := DC1(-v25[785]);
					var v31: map<D0, set<bool>> := map[v30 := v7];
					v29[958] := v0;
					var v32: map<bool, bool> := map[false := v0];
					r0, v29[550], r0, v31, v29[958] := v25[785] % p0, !v0, cf2 % v25[785], if (v0 in v32) then v31 else v31[v30 := v7], true;
					v29[550] := v25[785] <= 0x1d6;
				}
				
				var v33 := new C3();
				cf2 := p0 * p0;
				v0 := !(v0 <==> !(v7 >= {true}));
		}
		
		r0 := p0;
		var v34: T0 := new C3();
		var v35: seq<bool> := [v0, v0, v0, v0];
		var v36: map<T0, seq<bool>> := map[v34 := v35];
		r1 := |(if (v34 in v36) then v36[v34] else v35)|;
		r2 := v35 + v35;
	}
}

method Main() {
	var globalState := new GlobalState();
	var v0 := 936;
	var v1: C4 := new C4();
	var v3 := "dtn";
	var v4: multiset<string> := multiset{v3, v3};
	var v5: array<set<bool>> := new set<bool>[5](i0 => {true, true});
	var v6 := DC0(v5, v0);
	var v7: seq<int> := [v0];
	var v8 := DC1(v0);
	var v9 := new C8(|map[|[v0]| := v1]| / |(map v2 : string | v2 in v4 :: (v2) := (-0x2a4))|, v6, v7 + v7, v8);
	var v10 := false;
	var v11, v12, v13 := v9.m3(0x1ea, DC0(v5, |v3|).(cf1 := v0), v10, globalState);
	var v14: map<bool, int> := map[v10 := 0x224];
	var v15: map<map<bool, int>, string> := map[v14 := v3];
	var v16 := DC47(v15);
	match (if (v10) then v16.(cf85 := v16.cf85) else v16) {
		case DC48() =>
			if (v10) {
				var v17: array<map<int, bool>> := new map<int, bool>[8];
				v17 := v17;
				var v18: array<bool> := new bool[14];
				v18[213] := v10;
				var v19: multiset<int> := multiset{v11};
				var v20: T1 := new C7(v19, v8.(cf2 := v9.f2), v7 + v7, DC1(v0));
				var v21: array<int> := new int[17];
				v21[721] := |fm70(v10, globalState)|;
				var v22 := 'e';
				var v23 := DC43(seq(0x193, i1  => ('c')), v22, v11);
				var v24: C0 := new C0(v23.cf79);
				var v25: map<C0, bool> := map[v24 := v10];
				v18[213], v20, v18, v21[721], v10 := v10, if (if (v24 in v25) then v25[v24] else v10) then v20 else v20, v18, v9.f2, (fm44(true, globalState) < v13) && v10;
				v18[213] := (v10 !in v14) ==> true;
				var v26: map<bool, bool> := map[v18[213] := fm17(globalState)];
				var v27: seq<bool> := [v18[213], false, v10, v10, v18[213]];
				var v28: multiset<bool> := multiset{v27[v13]};
				var v29: map<int, int> := map[v9.f2 := |(multiset{!(if (v10 in v26) then v26[v10] else v18[213])} - v28)|];
				v29 := v29[v13 := v9.f2];
				var v30, v31 := v9.m1(if (v21[721] in v29) then v29[v21[721]] else v23.cf81, globalState);
			} else {
				var v32 := 'c';
				var v33: map<char, int> := map[v32 := -fm32(v9.f2, globalState)];
				var v34: array<int> := new int[27];
				var v35: T1 := new C5(v34, v7, v8);
				var v36: C0 := new C0(v3);
				var v37: map<T1, C0> := map[v35 := v36];
				var v38: array<bool> := new bool[29] [false, v10, v10, v10, v10, v10, v10, v10, v10, v10, v10, v10, v10, true, v10, !v10, fm17(globalState), true, true, v10, v10, false, v10, v10, v10, true, v10, v10, v10];
				var v39, v40 := v1.m10(v9.f2 + v9.f2, |v33[fm51(v10, v9.f2, |v37|, globalState) := v0]|, v10, v38, globalState);
				var v41: seq<bool> := [v10];
				v41 := [v39 <==> v10, v39, false, v10];
				v11 := -0x1fd;
				var v42: seq<array<bool>> := [v38, v38, v38];
				v10 := [v38, v38] <= v42;
				var v43: map<bool, bool> := map[v10 := false];
				var v44, v45, v46 := v9.m3(--v13, v6, if (false in v43) then v43[false] else !v10, globalState);
			}
			
			var v47: seq<bool> := [!v10, v10];
			var v48: set<char> := {'f'};
			var v49: map<seq<bool>, seq<bool>> := map[v47 := v47[|v48| := v10]];
			v49 := map[v47 := v47] + (v49 + v49);
			v10 := v10;
			var v50: T0 := new C4();
			var v51 := DC53(v10);
			var v52: map<T0, bool> := map[v50 := v51.cf94];
			var v53 := DC55(v50);
			v52 := v52[v53.cf98 := true];
		case DC49(cf86) =>
			if (v10) {
				var v54 := 'n';
				var v55: map<char, D0> := map[v54 := v8];
				var v56: map<bool, seq<seq<bool>>> := map[true := fm71(v9.fm5(v9.f2, v10, v55, v10, globalState), globalState)];
				var v57: seq<bool> := [v10, v10, v10];
				var v58: multiset<bool> := multiset{v10, v10};
				var v59: seq<seq<bool>> := [[v10, !v10]];
				v11 := |(if ((multiset(v57) > v58) in v56) then v56[multiset(v57) > v58] else v59 + v59)|;
				var v60: set<bool> := {v10};
				v14 := v14[v60 > v60 := v9.f2];
				var v62: multiset<int> := multiset{|(set v61 : int | v61 in v7 :: (v61 - 0xb9))|, 12};
				var v63 := new C7(((v62[v9.f2 := v0])[v0 := v13])[v0 := v0], if (v10) then fm49(v54, globalState) else v8, v7, v8);
				var v64: array<map<char, int>> := new map<char, int>[18];
				var v65: C1 := new C1(v11, v64, v7[v9.f2 := v0], DC1(v0));
				var v66: map<C1, char> := map[v65 := v54];
				v66 := v66[v65 := v54];
				var v67: T0 := new C8(v65.f9, v6, seq(169, i2  => (v9.f2)), DC1(|v57|));
				var v68: map<T0, int> := map[v67 := v9.f2];
				var v69: array<bool> := new bool[24];
				v69[414] := v10;
				v68, v69[414] := v68 + v68, v10;
			} else {
				var v70: array<D17> := new D17[1];
				var v71: multiset<int> := multiset{v13, v9.f2};
				var v72 := DC45(multiset{v71, v71});
				v70[395] := v72;
				v70[395] := v72;
				var v73 := DC56(v1);
				var v74: array<C4> := new C4[28] [v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v73.cf99, v1, v1, v1, v1, v1, v1];
				var v75: seq<array<C4>> := [v74, v74, v74, v74];
				v11, v3, v75 := v9.f2, v9.fm7(v13, globalState), if (v10) then [v74, v74] + [v74] else v75;
				var v76: array<array<int>> := new array<int>[21];
				var v77 := DC10(-(v11 + v9.f2), v13, |v3|);
				var v78: seq<array<array<int>>> := [v76, v76, v76, v76];
				var v79: map<bool, array<array<int>>> := map[v10 := v78[fm44(v10, globalState)]];
				v11, v76, v77, v11 := |((seq(335, i3  => ('d'))) + (seq(235, i4  => ('h'))))|, if (true in v79) then v79[true] else v76, v77, if (v10) then v9.f2 else fm32(-v0, globalState);
				v11 := (v0 + -0x26a) - -|v3|;
				v11 := if (v3 in v4) then v4[v3] else v11;
			}
			
			v10 := v10;
			var v80 := DC18();
			v80 := v80;
			v10 := v10 <==> (v11 == v9.f2);
		case DC50(cf87, cf88, cf89, cf90, cf91) =>
			var v81: array<map<char, int>> := new map<char, int>[7];
			var v82: seq<seq<int>> := [v7];
			var v83: T1 := new C1(-0x114 + v9.f2, if (cf89) then v81 else v81, [v0, |multiset{|{v82[v9.f2]}|}|], v8);
			v83 := v83;
			cf89 := !!v10;
			var v84, v85 := v9.m1(v13, globalState);
			v11 := v9.f2;
		case DC47(cf85) =>
			v10 := v10;
			v10 := !(0x30c != (if (v10) then v11 else 770));
			v11 := v13 * (v11 / 0x2c3);
			var v86: set<int> := {v0};
			var v87: set<set<int>> := {v86};
			var v88: map<bool, set<set<int>>> := map[v10 := DC57(v87).cf100];
			v88 := v88[v10 := v87];
	}
	
	var v89 := DC53(!v10);
	if (v89.cf94) {
		v11 := v9.f2;
		var v90: array<int> := new int[23](i5 => i5 % v9.f2);
		v90[898] := |(v3 + "bnkl")|;
		var v91: set<bool> := {v10};
		v90[898] := |({v10} - v91)|;
		var v92: T0 := new C2();
		var v93: multiset<int> := multiset{0x18a};
		var v94 := DC55(v92);
		var v95: map<bool, T0> := map[!(v93 == v93) := v94.cf98];
		var v96: map<char, D0> := map['m' := v8];
		var v97: map<bool, bool> := map[!!v10 := v92.fm5(v0, v10, v96, v10, globalState)];
		v92 := if ((if (v10 in v97) then v97[v10] else v10) in v95) then v95[if (v10 in v97) then v97[v10] else v10] else v92;
		var v98: array<array<D16>> := new array<D16>[24];
		var v99: array<D16> := new D16[15](i6 => DC42(v9.f2, |v14|, v0));
		v98[548] := v99;
		var v100 := DC59(v99);
		v98[548] := v100.cf101;
		var v101 := 'a';
		v101 := v101;
	} else {
		var v102: array<bool> := new bool[19];
		v102[189] := v10;
		var v103 := DC28(v10, |v3|);
		v102[189] := (v103.cf47 % 0x35c) == 0x2fd;
		v10 := v10 <== v10;
		var v105: set<int> := {-v0, v11, |(map v104 : map<int, char> | v104 in (seq(0xa3, i7  => (v12))) :: (v104) := (v13))|, v0};
		var v106: seq<bool> := [v102[189]];
		var v107 := 'm';
		var v108: map<bool, char> := map[v102[189] := v107];
		var v109: multiset<int> := multiset{v9.f2, |v108|, v11};
		v102 := new bool[29] [v102[189], v10, v102[189], v102[189], !v102[189], v102[189], v102[189], v10, true, v102[189], if (v10) then v102[189] else false, v10, v10, v10 <==> v10, v102[189], v102[189], v102[189], v102[189], v102[189], v13 >= v0, {v0} >= v105, v102[189], v102[189], v3 < (seq(-976, i8  => ('q'))), v102[189], !v106[|v109|], v102[189] <==> true, v10, v102[189] && v10];
		var v110: map<bool, bool> := map[v10 := if (v10) then v102[189] else v102[189]];
		v110 := v110[v0 == v13 := v10];
		var v111: map<int, int> := map[v13 := v9.fm6(v0, |v7|, globalState)];
		v0 := if (v11 in v111) then v111[v11] else v0;
	}
	
	var v112: map<int, int> := map[v0 := v9.f2];
	v0 := v9.f2 * -|v112|;
	var i9 := 0;
	while (!match v16 {
		case DC48() => v10
		case DC49(cf86) => v10
		case DC50(cf87, cf88, cf89, cf90, cf91) => cf87
		case DC47(cf85) => v10
	})
		decreases 100 - i9
	{
		if (i9 >= 100) {
			break;
		}
		
		i9 := i9 + 1;
		var v113: map<bool, bool> := map[v10 := v10];
		v113, v10 := fm47(v9.f2 * v0, v9.f2 <= v0, globalState), v11 > v0;
		var v114: array<bool> := new bool[5];
		v114[792] := v10;
		var v115: T0 := new C2();
		v1, v114[792], v114, v0 := v1, !!v10, v114, v13 + |[v115]|;
		var v116: C3 := new C3();
		var v117: seq<C3> := [v116];
		v116 := if (v10) then v116 else v117[v9.f2];
		v10 := v114[792];
	}
	var v118 := new C4();
	var v119: array<int> := new int[15];
	var v120: map<array<int>, array<int>> := map[v119 := v119];
	var v121: map<bool, map<array<int>, array<int>>> := map[v10 := map[v119 := v119]];
	var v122: array<map<array<int>, array<int>>> := new map<array<int>, array<int>>[16] [if (v10) then v120 else v120, v120 + v120, map[v119 := v119] + map[v119 := v119], v120 + v120, v120 + map[v119 := v119], v120[v119 := v119], map[v119 := v119] + v120, if (v10 in v121) then v121[v10] else v120, v120, v120, v120[v119 := v119], v120, v120, v120, v120, v120];
	v122[343] := v120;
	v122[343] := v120;
	v10 := false;
	var v123: set<bool> := {false, v10, true, v10, true};
	var v124: map<bool, set<bool>> := map[v10 := v123];
	var v125: map<int, map<bool, set<bool>>> := map[v0 := v124];
	var v126: seq<bool> := [false, v10, false];
	var v127: seq<map<bool, set<bool>>> := [v124, v124, v124];
	var v128: array<map<bool, set<bool>>> := new map<bool, set<bool>>[28] [if (|v126| in v125) then v125[|v126|] else v124, v124, v124, v124, v124, v124, v124 + map[false := v123], v127[v0], v124, v124, map[true := v123] + v124, v124, v124, v124, v124, v124, v124, v124, if (v10) then v124 else v124, v124 + v124, v124, if (v9.f2 in v125) then v125[v9.f2] else v124, v124[true := v123], v124, v124, v124, v124, map[false := v123] + v124];
	v128[312] := v124;
	v128[312] := v124;
	v10 := v10;
	var v129 := 'a';
	var v130: array<char> := new char[3] [v129, v129, v129];
	var v131: seq<array<char>> := [v130];
	var v132: map<bool, array<char>> := map[v89.cf94 := v131[v11]];
	v132 := v132[v10 := v130];
	if (multiset{"vdgmtsaje", "lhgie"} == v4) {
		var v133: C0 := new C0(v3);
		var v134: map<C0, bool> := map[v133 := v10];
		v10 := if (if (v133 in v134) then v134[v133] else v10) then if (v126[v13]) then v10 else v10 else v10;
		var v135 := DC62(v9.f2);
		var v136: map<D24, array<int>> := map[v135.(cf104 := v11) := v119];
		var v137: array<array<int>> := new array<int>[27] [v119, v119, v119, v119, v119, v119, v119, v119, v119, v119, v119, v119, v119, if (v135 in v136) then v136[v135] else v119, v119, v119, v119, if (true) then v119 else v119, v119, v119, v119, if (v10) then v119 else v119, v119, v119, v119, v119, v119];
		v137[439] := v119;
		v137[439] := v119;
		var v138 := DC38(v129);
		match (v138) {
			case DC38(cf71) =>
				var v140: set<int> := {|v3|, v11};
				var v141: T0 := new C8(|(if (v10) then set v139 : int | (0x6d <= v139) && (v139 < 874) :: (v139 + v9.f2) else v140)|, DC0(v5, v9.f2), v7 + v7, v8);
				v141 := new C2();
				var v142 := v9.m2(v11, globalState);
				var v143 := DC42(v9.f2, v9.f2, v13);
				var v145: multiset<bool> := multiset{v10};
				var v146: seq<multiset<bool>> := [v145, v145, v145];
				var v147: seq<string> := [v133.f8, v133.f8, v133.fm15([v13, |(map v144 : multiset<bool> | v144 in multiset(v146) :: (v144) := (cf71))|, v9.f2, v11], v10, globalState), v133.f8, "s"];
				var v148: array<seq<string>> := new seq<string>[7] [(seq(0xe0, i10  => (DC39(v3, DC21(multiset{false, v10})).cf72))) + (seq(-0x33b, i11  => (v133.f8)))[v143.cf77 := v3], v147, v147, fm72(globalState), v147, [v3, v133.f8, "ea", v133.f8, v133.f8], v147[549 := v133.f8]];
				v148[350] := v147[v13 := v3];
				v148[350] := v147 + (v147 + v147);
				cf71 := cf71;
			case DC39(cf72, cf73) =>
				var v149: map<multiset<string>, bool> := map[v4 := |v112| >= v9.f2];
				v149 := v149[multiset{v133.f8, v133.f8, v133.f8} := v10];
				var v150: T0 := new C8(v9.f2, v6, v7, v8);
				var v151: multiset<T0> := multiset{v150};
				var v152: map<int, string> := map[fm32(-(if (v150 in v151) then v151[v150] else v9.f2), globalState) := v133.f8];
				v152 := v152[v0 := v133.f8 + (seq(0x393, i12  => (v129)))];
				var v153 := new C9();
				var v154 := DC23(v137[439]);
				var v155: multiset<D10> := multiset{DC23(v119), v154, v154};
				v155 := v155;
			case DC37(cf70) =>
				var v156, v157, v158 := v1.m9(globalState);
				v11 := v9.f2 * fm20(v158, globalState);
				var v159: set<int> := {v9.f2};
				var v160: multiset<bool> := multiset{v158};
				var v161: array<bool> := new bool[22] [v158, v133.f8 < "jtsgpyod", v158, v159 == v159, v10, v158, !(!v158 !in v160), v10, v158, v158, v10, v10, v158 ==> v158, v10, v10, v126[|multiset{v11}|], !v158, v158, v10, v129 in v133.f8, v158, v10 ==> v10];
				v161[491] := v10 <== v10;
				v161[491], v11, v158 := v13 <= (v11 + v9.f2), v9.f2, !v10;
				var v162: multiset<int> := multiset{v13};
				v126, v161[491], v161[491], v161[491] := v126, true, (v162 + fm24(v10, globalState)) >= v162, v161[491];
			case DC40(cf74) =>
				v13 := (|[v118]| % v9.f2) % v9.f2;
				var v163 := new C4();
				var v164 := v118.m2(v9.f2, globalState);
				var v165: map<int, bool> := map[v9.f2 := v10];
				var v166: map<char, D0> := map[v129 := v8.(cf2 := v9.f2)];
				v10 := v9.fm5(0xc8 / v11, 993 < |v165|, v166, v10, globalState);
		}
		
		v10 := "sdjic" == ("q" + v133.f8)[|v126| := v129];
		var v167: map<int, array<int>> := map[v9.f2 % v9.f2 := v119];
		v167 := v167;
	} else {
		var v168: array<array<int>> := new array<int>[14] [v119, v119, v119, v119, v119, v119, v119, v119, v119, v119, v119, v119, v119, v119];
		var v169: multiset<bool> := multiset{v10, v10, v10};
		var v170: map<int, multiset<bool>> := map[if (v10 in v169) then v169[v10] else 0xca := multiset(v126)];
		var v171 := DC63(v170);
		var v172: map<array<array<int>>, map<int, multiset<bool>>> := map[v168 := v171.cf105[879 := if (v0 in v170) then v170[v0] else v169]];
		v172 := v172[v168 := v170];
		var v173 := new C4();
		var v174 := DC61(v11, multiset{v119});
		v174 := v174;
		var v175 := new C5(v119, [v13], v8);
		v10 := v3[|v3|] !in "xg";
	}
	
	var i13 := 0;
	while (if (v10) then v10 else v10)
		decreases 100 - i13
	{
		if (i13 >= 100) {
			break;
		}
		
		i13 := i13 + 1;
		var v176 := DC44(v11);
		var v177: set<D16> := {DC44(v9.f2), v176, v176, v176};
		var v178 := DC25(v129, v9.f2, v10);
		v10, v10, v0, v10 := (v177 * v177) !! v177, v178.cf43, 559, v10;
		var v179: map<seq<bool>, int> := map[v126 := v9.f2];
		v179 := v179[v126 := v13];
		v11 := if (v10) then v11 else -v9.f2;
		var v180, v181 := v118.m1(v0 + v11, globalState);
	}
	v13 := 670;
	v3 := v3;
}