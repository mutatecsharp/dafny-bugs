datatype D0 = DC1(cf1: bool, cf2: bool, cf3: int, cf4: bool) | DC0(cf0: bool) | DC2(cf5: D0)
datatype D1 = DC4(cf7: array<char>, cf8: int, cf9: int, cf10: string, cf11: string) | DC5(cf12: int) | DC3(cf6: seq<bool>)
datatype D2 = DC7(cf14: bool, cf15: bool, cf16: int, cf17: int) | DC8 | DC9(cf18: int, cf19: int, cf20: bool) | DC6(cf13: map<int, int>)
datatype D3 = DC11(cf22: char) | DC10(cf21: array<int>)
datatype D4 = DC12(cf23: set<bool>)
datatype D5 = DC14(cf25: int, cf26: int, cf27: bool) | DC13(cf24: map<bool, seq<bool>>) | DC15(cf28: D5)
datatype D6 = DC17(cf30: int, cf31: int, cf32: int, cf33: int, cf34: int) | DC16(cf29: string)
datatype D7 = DC19(cf36: int, cf37: multiset<bool>, cf38: bool, cf39: map<bool, int>) | DC18(cf35: T2)
datatype D8 = DC20(cf40: seq<int>)
datatype D9 = DC22(cf42: bool) | DC21(cf41: map<seq<bool>, array<bool>>)
datatype D10 = DC24(cf44: string, cf45: seq<map<bool, int>>, cf46: bool) | DC25 | DC23(cf43: map<seq<int>, seq<D2>>)
datatype D11 = DC26(cf47: array<bool>)
datatype D12 = DC28(cf49: int) | DC27(cf48: T0)
datatype D13 = DC30(cf51: bool, cf52: int, cf53: bool, cf54: array<bool>) | DC29(cf50: map<bool, bool>)
datatype D14 = DC32(cf56: multiset<bool>, cf57: bool, cf58: bool, cf59: bool, cf60: int) | DC31(cf55: seq<set<bool>>)
datatype D15 = DC34(cf62: int, cf63: bool, cf64: int) | DC35(cf65: bool) | DC36(cf66: bool, cf67: int) | DC33(cf61: map<C5, bool>)
datatype D16 = DC38(cf69: int, cf70: int, cf71: bool) | DC37(cf68: map<int, bool>) | DC39(cf72: D16)
datatype D17 = DC41(cf74: bool, cf75: bool, cf76: int) | DC42(cf77: int, cf78: int, cf79: int) | DC43(cf80: int, cf81: bool) | DC40(cf73: seq<seq<int>>)
datatype D18 = DC44(cf82: multiset<map<int, bool>>)
datatype D19 = DC45(cf83: map<map<int, bool>, int>)
datatype D20 = DC47 | DC46(cf84: T1) | DC48(cf85: D20)
datatype D21 = DC50(cf87: bool, cf88: int, cf89: bool, cf90: int) | DC51(cf91: int, cf92: T1, cf93: int, cf94: int) | DC52(cf95: int, cf96: seq<map<bool, int>>) | DC49(cf86: map<bool, map<int, int>>)
datatype D22 = DC54(cf98: char, cf99: D9, cf100: bool, cf101: bool, cf102: bool) | DC53(cf97: set<T0>)
datatype D23 = DC56(cf104: int, cf105: int, cf106: int, cf107: bool) | DC57 | DC55(cf103: C3) | DC58(cf108: D23)
datatype D24 = DC60(cf110: T1, cf111: int, cf112: int, cf113: bool) | DC59(cf109: multiset<D1>) | DC61(cf114: D24)
datatype D25 = DC63(cf116: string, cf117: bool, cf118: int) | DC62(cf115: C8) | DC64(cf119: D25)
datatype D26 = DC66(cf121: int, cf122: bool, cf123: C3) | DC65(cf120: multiset<array<int>>) | DC67(cf124: D26)
class GlobalState {
	const f0 : set<int>
	const f1 : multiset<seq<char>>
	var f2 : bool
	const f3 : bool
	const f4 : char
	var f5 : array<array<bool>>
	var f6 : string
	const f7 : map<string, array<int>>
	const f8 : bool
	const f9 : bool
	const f10 : int
	const f11 : bool
	var f12 : int
	const f13 : int
	const f14 : array<int>
	var f15 : bool
	var f16 : int
	const f17 : bool
	const f18 : bool
	var f19 : char
	const f20 : string
	const f21 : int
	var f22 : bool
	const f23 : int
	var f24 : bool
	const f25 : map<map<char, int>, int>
	constructor (f0 : set<int>, f1 : multiset<seq<char>>, f2 : bool, f3 : bool, f4 : char, f5 : array<array<bool>>, f6 : string, f7 : map<string, array<int>>, f8 : bool, f9 : bool, f10 : int, f11 : bool, f12 : int, f13 : int, f14 : array<int>, f15 : bool, f16 : int, f17 : bool, f18 : bool, f19 : char, f20 : string, f21 : int, f22 : bool, f23 : int, f24 : bool, f25 : map<map<char, int>, int>) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
		this.f5 := f5;
		this.f6 := f6;
		this.f7 := f7;
		this.f8 := f8;
		this.f9 := f9;
		this.f10 := f10;
		this.f11 := f11;
		this.f12 := f12;
		this.f13 := f13;
		this.f14 := f14;
		this.f15 := f15;
		this.f16 := f16;
		this.f17 := f17;
		this.f18 := f18;
		this.f19 := f19;
		this.f20 := f20;
		this.f21 := f21;
		this.f22 := f22;
		this.f23 := f23;
		this.f24 := f24;
		this.f25 := f25;
	}
	
}

function fm0(p0: bool, globalState: GlobalState): multiset<string> {
	multiset{"cfbosyfj", seq(0x3de, i0  => ('u')), "w", "u" + (seq(0x37a, i1  => ('y')))}
}
function fm8(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): map<int, int> {
	map[|multiset{|(seq(0x1a8, i0  => ('t')))|, |map[|map[--0x3d5 := !false]| := DC3([!true])]|}| * |(set v3 : int | v3 in map[|(map v0 : set<int> | v0 in {set v2 : int | v2 in map[|multiset{multiset{|(map v1 : int | (0x118 <= v1) && (v1 < 0x78) :: (v1 * -118) := (|{map[|"cth"| := |map[0x154 := 0x5a]|], map[|{true, true}| := 0x26e]}|))|, |[|{true}|, |{-625}|]|}}| := |multiset{|[-0xea, |"qpmel"|]|}|] :: (v2 * |(seq(751, i1  => (-|(seq(0x3b1, i2  => (|[false]|)))|)))|)} :: (v0) := (0x5a))| := !true] :: (v3 + |{-0x398}|))| := |map[|(set v4 : int | v4 in map[|(seq(0x2e3, i3  => ('f')))| := |{false}|] :: (v4 * 0x1b6))| := |"jiwdrv"|]| - 0x276]
}
function fm9(p0: bool, p1: int, p2: bool, globalState: GlobalState): multiset<bool> {
	match DC12({false}) {
		case DC12(cf23) => multiset{false} * multiset([false])
	}
}
function fm10(p0: int, p1: int, p2: D0, globalState: GlobalState): multiset<int> {
	multiset([--0x26d]) - (multiset{545} * multiset([352, |[true]|]))
}
function fm11(p0: bool, p1: bool, p2: char, globalState: GlobalState): seq<bool> {
	match DC32(multiset{!true}, true, false, !true, 0x21e) {
		case DC32(cf56, cf57, cf58, cf59, cf60) => [cf59] + [cf58]
		case DC31(cf55) => [!true, true] + [true]
	}
}
function fm14(p0: bool, p1: int, p2: int, globalState: GlobalState): char {
	'd'
}
function fm15(globalState: GlobalState): seq<int> {
	((seq(835, i0  => (|{false}|))) + [0xd7]) + (seq(0xb, i1  => (0x258)))
}
function fm18(p0: set<bool>, globalState: GlobalState): set<bool> {
	{false ==> !false}
}
function fm19(p0: int, p1: int, globalState: GlobalState): D0 {
	DC0(0x196 == -0x82)
}
function fm20(p0: int, p1: bool, p2: int, p3: int, globalState: GlobalState): D2 {
	DC9(-0x24, 441, false && false)
}
function fm21(p0: int, globalState: GlobalState): multiset<bool> {
	multiset{false, false, !true} * multiset{false}
}
function fm22(p0: int, p1: bool, globalState: GlobalState): string {
	"vvtodqg"
}
function fm23(p0: char, globalState: GlobalState): seq<int> {
	match DC5(0x31c) {
		case DC4(cf7, cf8, cf9, cf10, cf11) => if (!true) then [cf8] else [|(map v0 : int | (-0x213 <= v0) && (v0 < 0x2ce) :: (v0 * cf9) := (true))|]
		case DC5(cf12) => [cf12, 584, -cf12, cf12] + [cf12]
		case DC3(cf6) => [0x6b]
	}
}
function fm24(p0: string, p1: bool, p2: int, globalState: GlobalState): char {
	'v'
}
function fm25(p0: bool, p1: bool, p2: char, globalState: GlobalState): seq<bool> {
	[[0xf4, 0x244] <= [-0x296, |[!false, true, true, true]|], true]
}
function fm26(p0: bool, p1: bool, p2: bool, globalState: GlobalState): multiset<map<int, int>> {
	multiset{map[|['o']| := |map[327 := true]|], map[|"jsiwka"| := |"vjgqe"|]} * (multiset([map v0 : int | v0 in map[0x14e := |map[true := |{false, true}|]|] :: (v0 % |{true}|) := (0x3c9), map[0xf3 := |multiset{DC54('u', DC22(false), true, true, true).cf98, 'x'}|], map[|map[true := true]| := |[|[|(map v1 : int | (-0x2a <= v1) && (v1 < 0x257) :: (v1 % 0x0) := (335))|]|]|], map[0x30 := 0x315]]) + multiset{map[0xa2 := |"bwfgndon"|]})
}
function fm27(p0: int, p1: int, globalState: GlobalState): map<string, bool> {
	match DC54('r', DC22(true), false, true, !false) {
		case DC54(cf98, cf99, cf100, cf101, cf102) => map["rqm" := cf100]
		case DC53(cf97) => map["smnurn" := !false] + map[seq(0x275, i0  => ('j')) := !true]
	}
}
function fm28(globalState: GlobalState): set<D0> {
	(set v0 : D0 | v0 in [DC1(true, true, |map[false := |multiset{true, false}|]|, !true), DC1(!false, !false, 0x129, true), DC1(false, !false, 106, true)] :: (v0)) - {DC1(false, !false, |map[false := false]|, false)}
}
function fm29(p0: bool, p1: int, globalState: GlobalState): map<int, int> {
	(map[|map[false := false]| := |[false]|] + map[0x29b := -|"jp"|]) + map[|['e']| := |DC24("lwtjp", [map[false := 0xe3]], true).cf44|]
}
function fm30(p0: int, p1: bool, globalState: GlobalState): map<bool, int> {
	(map[!true := |"ojauypo"|] + map[true := |[true]|]) + map[false := -0x3d3]
}
function fm31(p0: int, globalState: GlobalState): map<bool, map<int, int>> {
	DC49(map[true := map[|{0x329, 0x2e9, |"mf"|}| := 326]]).cf86
}
function fm32(p0: int, p1: int, globalState: GlobalState): map<bool, seq<bool>> {
	if (!false && !false) then map[true := [false]] + map[false := [true, true]] else map[false := [true]]
}
function fm33(p0: int, p1: int, globalState: GlobalState): D2 {
	DC8()
}
function fm34(globalState: GlobalState): map<seq<int>, seq<D2>> {
	(map[[|"ovj"|, 0x39] := [DC6(map[|multiset{false}| := 362])]] + map[[|[true]|] := [DC6(map[|"vwg"| := 725])]]) + (map[seq(412, i0  => (|"uxbg"|)) := seq(0x16c, i1  => (DC6(map[|"jlhmrsu"| := |[|[false, false]|]|])))] + map[[|map[{|map[|map[|map[|map[false := |map[|map[false := 0x3e0]| := -660]|]| := true]| := false]| := |(seq(-63, i2  => ('j')))|]|} := |[true]|]|] := [DC6(map v0 : int | v0 in {0x11b} :: (v0 * -|(seq(980, i3  => ('g')))|) := (---0xfe))]])
}
function fm35(p0: int, globalState: GlobalState): seq<D5> {
	if (true) then seq(0x44, i0  => (DC14(-|map[false := {false, true, true}]|, |map[true := |(seq(161, i1  => (|map[true := true]|)))|]|, true))) else [DC14(219, |{0x112}|, false)]
}
function fm36(p0: int, p1: bool, p2: bool, p3: bool, globalState: GlobalState): multiset<seq<D2>> {
	multiset{[DC6(map[-0x44 := -0x1e0])], seq(-0xff, i0  => (DC6(map[--0xf8 := 0x1b4]))), [DC6(map[282 := 0x37f]), DC6(map v0 : int | (691 <= v0) && (v0 < 0x3d5) :: (v0 - 0xca) := (-438)), DC6(map v1 : int | (-193 <= v1) && (v1 < 0x26a) :: (v1 / 560) := (|"ao"|))] + [DC6(map[|multiset{-0x289}| := 384])]}
}
function fm37(p0: bool, p1: int, p2: int, p3: D10, globalState: GlobalState): map<D6, char> {
	map[DC16("sqdccsvha") := 'h'] + ((map v0 : D6 | v0 in map[DC16("nrjkdbhoe") := false] :: (v0) := ('n')) + map[DC16("mxdmgbo") := 'd'])
}
function fm38(p0: bool, p1: int, p2: string, globalState: GlobalState): D6 {
	DC16((seq(0x1a7, i0  => ('x'))) + (seq(0x87, i1  => ('h'))))
}
function fm39(globalState: GlobalState): int {
	0x1df
}
function fm40(p0: int, p1: bool, p2: bool, globalState: GlobalState): char {
	match DC31([{!true}, {false}, {!true}, {DC24("hy", seq(846, i0  => (map[true := 0x357])), true).cf46}]) {
		case DC32(cf56, cf57, cf58, cf59, cf60) => if (cf58) then 'v' else 'i'
		case DC31(cf55) => if (true) then 'r' else 'q'
	}
}
function fm41(p0: bool, p1: D3, p2: int, p3: bool, globalState: GlobalState): seq<set<bool>> {
	[{true}, {false, true}, {true} - {!!true, false, true, false}, {true, false, false} * {DC7(true, true, -255, |multiset([true, true])|).cf15, true, true, !false, false}, if (true) then {true, false, true, false} else {!!true}]
}
function fm42(p0: int, globalState: GlobalState): multiset<int> {
	multiset{|"dagoucki"|} - (multiset([-0x1ea]) * multiset([|{map[false := 'q']}|, 0x2f8]))
}
function fm43(p0: int, p1: int, p2: int, p3: map<bool, char>, globalState: GlobalState): seq<int> {
	[-(|(seq(0x398, i0  => ('s')))| / 0x181)]
}
function fm45(p0: bool, p1: int, p2: int, p3: bool, globalState: GlobalState): seq<int> {
	[-0xfa - |[!true]|]
}
function fm46(p0: bool, p1: int, p2: int, p3: bool, globalState: GlobalState): map<int, map<int, bool>> {
	(map[528 := map[|[true]| := false]] + map[117 := map[|map[|multiset{false}| := |multiset([|"chdvwyj"|, |(map v0 : int | v0 in [|DC19(|{'v'}|, multiset{true}, true, map[true := 794]).cf37|, |(seq(0x2a8, i0  => ('t')))|] :: (v0 - 0x1a7) := (0x38d))|, |map[false := 612]|])|]| := !!true]]) + ((map v1 : int | (0x1ba <= v1) && (v1 < 367) :: (v1 + -0x160) := (map[0x1ec := false])) + map[|map[|(seq(0x1bf, i1  => (0x1d1)))| := false]| := map v2 : int | (812 <= v2) && (v2 < 557) :: (v2 * 0x37e) := (true)])
}
function fm47(p0: bool, p1: bool, p2: bool, p3: bool, globalState: GlobalState): map<int, bool> {
	match DC37(map[610 := true]) {
		case DC38(cf69, cf70, cf71) => map[|[|[cf71]|]| := cf71]
		case DC37(cf68) => map[|map["vtfmeu" := false]| := true] + (map v0 : int | v0 in cf68 :: (v0 % -0x304) := (true))
		case DC39(cf72) => map[|"tqifsibhn"| := true]
	}
}
function fm48(p0: int, globalState: GlobalState): multiset<char> {
	multiset{'o'} - multiset(seq(460, i0  => (DC11('t').cf22)))
}
function fm49(p0: multiset<int>, p1: int, p2: bool, globalState: GlobalState): multiset<multiset<int>> {
	(if (true) then multiset{multiset{|(set v0 : char | v0 in map['w' := 536] :: (v0))|, |{923}|}} else multiset{multiset([575, 0x131, 0x2e]), multiset{-264, 0x2b5}, multiset{0x278}, multiset{0xaa}}) - multiset{multiset([|map[0xb2 := !false]|]), multiset{|map[false := |[false]|]|}}
}
function fm50(globalState: GlobalState): multiset<set<int>> {
	multiset{{-0x180}} - (multiset{set v0 : int | v0 in [0x153, 982] :: (v0 - |[|(map v1 : int | (-0x170 <= v1) && (v1 < 0x64) :: (v1 / |map[|(seq(-0xaf, i0  => ('i')))| := |{|map[0x3a4 := multiset{733}]|}|]|) := (0x3dc))|, |"xnb"|]|)} * multiset{{|map[true := 0xf6]|}})
}
function fm51(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): set<int> {
	{-(if (false) then 0x1bb else 851), 0x23b, 0x33d}
}
function fm52(p0: bool, p1: int, globalState: GlobalState): set<bool> {
	{!false} + {false}
}
function fm53(p0: int, p1: int, p2: set<int>, p3: map<int, int>, globalState: GlobalState): D6 {
	match DC23(map[[|(map v0 : set<int> | v0 in [{|multiset{-916}|, 671}, {|map[false := true]|}, set v1 : int | v1 in [|map[true := true]|] :: (v1 % |{!false}|), {|[true, true]|}] :: (v0) := (-|(seq(64, i0  => ('s')))|))|, 660] := [DC6(map v2 : int | (0x36a <= v2) && (v2 < 0x1ef) :: (v2 * -0x74) := (823)), DC6(map[461 := 0x183])]]) {
		case DC24(cf44, cf45, cf46) => DC17(-489, |map[cf46 := false]|, 0x3cd, 0x11f, 0x110)
		case DC25() => DC17(|"tu"|, |map[false := false]|, -503, 0x1a, -|"grxj"|)
		case DC23(cf43) => if (false) then DC17(|[true]|, -0xb2, -617, 0x374, |(map v3 : map<int, bool> | v3 in [map v4 : int | (0x290 <= v4) && (v4 < 0x2e8) :: (v4 + -|[true]|) := (false)] :: (v3) := (map[|map[-0x34b := !false]| := 0x21a]))|) else DC17(|(seq(0x108, i1  => ('w')))|, --|"vghjbcr"|, -0x3db, |"rfqfeded"|, |['g']|)
	}
}
function fm54(p0: bool, p1: int, p2: int, globalState: GlobalState): D16 {
	DC38(-|(map v0 : int | (-0x162 <= v0) && (v0 < 574) :: (v0 % -0xad) := (false))| - 0x1bb, if (true) then |map[map[0x1b0 := -|{614, -|(map v1 : int | v1 in multiset{0x2ea, -|"wwtt"|} :: (v1 - |[0x2da]|) := ([true]))|, |"g"|}|] := [true, true, true]]| else --0x310, false)
}
function fm55(globalState: GlobalState): seq<set<int>> {
	[(set v0 : int | (284 <= v0) && (v0 < -0x62) :: (v0 % 210)) + (set v1 : int | (0x13a <= v1) && (v1 < 0x241) :: (v1 / -0x11f))]
}
function fm56(globalState: GlobalState): D3 {
	DC11(if (!true) then 'j' else 'd')
}
method m0(p0: char, globalState: GlobalState) returns (r0: char, r1: seq<bool>, r2: bool, r3: map<multiset<bool>, bool>) {
	var v0 := true;
	var v1: seq<bool> := [v0];
	var v2 := 762;
	if (v1[v2]) {
		var v3: multiset<int> := multiset{fm39(globalState)};
		var v4: map<multiset<int>, bool> := map[v3 := v0];
		var v5 := DC0(v0);
		var v6: seq<D0> := [v5, v5];
		var v7 := DC2(v6[v2]);
		v4 := v4[fm10(v2, v2, v7, globalState) := v0];
		globalState.f15 := !v0;
		var v8: array<int> := new int[29];
		v8[61] := v2 / v2;
		v8[61] := fm39(globalState);
		globalState.f22 := v0;
		globalState.f12 := |fm22(v8[61], v0, globalState)|;
	} else {
		var v9: map<int, bool> := map[v2 := true];
		var v10: map<bool, bool> := map[v0 := !v0];
		var v11: map<bool, int> := map[v0 := v2];
		var v12: seq<int> := [if (true in v11) then v11[true] else v2];
		var v13: seq<seq<int>> := [v12, v12, [|v1|, v2], v12, v12];
		var v14: C2 := new C2(v13, v0, v0);
		var v15: map<bool, C2> := map[v0 <==> (if (|v10| in v9) then v9[|v10|] else v0) := v14];
		var v16: multiset<bool> := multiset{v0, v0};
		var v17: C6 := new C6(|v16|);
		var v18: map<bool, C6> := map[v0 := v17];
		var v19 := DC1(true, v0, v12[|v18|], !v0);
		v15 := v15[v19.cf1 := v14];
		globalState.f2 := v0;
		var v20: array<bool> := new bool[28](i0 => true);
		v20[414] := v0;
		v20[414] := false;
		var v21: array<multiset<bool>> := new multiset<bool>[15] [v16, v16, v16 + v16, multiset(v1), multiset{false} + v16[v0 := -v2], v16, multiset{v20[414]}, fm9(true, v17.f40, v0, globalState), fm21(v17.f40, globalState), v16, v16, v16 - multiset([v20[414]]), multiset{v0, v20[414]}, v16 - fm9(v0, v2, v20[414], globalState), v16];
		v21[416] := v16;
		v21[416] := v16;
		var v22: seq<char> := [p0];
		if (|(v22 + [p0])| == 365) {
			globalState.f16 := 0xc8 + v2;
			var v23 := DC25();
			v23 := DC25();
			var v24: array<int> := new int[21](i1 => i1 / |{v17.f40, -v2, |map[v20[414] := v0]|}|);
			v24[201] := v17.f40;
			v24[338] := fm39(globalState);
			var v25 := DC22(v20[414]);
			var v26: map<seq<bool>, array<bool>> := map[v1 := v20];
			var v27 := DC21(v26);
			var v28: seq<D9> := [v27];
			v17.f40, v24[201], v0, v24[338] := v17.f40, |(v12 + v12)|, v25.cf42, |(if (v20[414]) then v28 else v28)|;
			globalState.f2 := v20[414] !in v1;
			v9 := v9 + map[v17.f40 := v20[414]];
		} else {
			var v29: set<int> := {|v11|};
			var v30 := DC22(if (v14.fm16(|v22|, [DC9(v2, |v29|, v0)], v17.f40, v17.f40, globalState)) then v20[414] else v0);
			v30 := v30;
			r1 := v1;
			var v31 := DC9(v17.f40, v2, if (v17.f40 in v9) then v9[v17.f40] else v20[414]);
			var v32: seq<D2> := [v31];
			var v33: map<int, int> := map[v2 := v17.f40];
			globalState.f12 := if (if (v0) then v0 else v14.fm16(v17.f40, v32, |v33|, v17.f40, globalState)) then 0x26e else v14.fm1(globalState) % |multiset{v17.fm2(v0, v2, globalState)}|;
			globalState.f16 := v17.fm1(globalState);
			globalState.f22 := p0 in v22;
		}
		
	}
	
	var v34: array<set<bool>> := new set<bool>[8](i2 => {true, v1[v2]} * {true, v0});
	var v35: set<bool> := {v0};
	v34[683] := v35 + v35;
	var v36: seq<map<bool, int>> := [map[v0 := v2]];
	var v37 := DC52(|[v0]|, v36);
	var v38: map<bool, D21> := map[v0 := v37];
	var v39: multiset<map<bool, D21>> := multiset{v38, v38, v38, v38[v0 := DC52(v2, v36)]};
	var v40: map<bool, set<bool>> := map[v0 := v35];
	var v41: seq<map<bool, D21>> := [v38];
	v34[683], v39 := fm18(if (v0 in v40) then v40[v0] else v35, globalState), v39 * multiset(v41);
	var v42: array<char> := new char[12];
	var v43: multiset<int> := multiset{v2};
	var v44: multiset<multiset<int>> := multiset{v43};
	globalState.f2, globalState.f12, v2, v42, r0 := v0, -(0x32c - ((if (v43 in v44) then v44[v43] else v2) + |fm18(v35, globalState)|)), v2 + v2, v42, p0;
	globalState.f16 := -v2;
	v34[683] := fm52(true <== v0, -(|v34[683]| - v2), globalState);
	for i3 := v2 to v2 {
		globalState.f2 := v0;
		v2 := v2 % -(|multiset{i3}| / v2);
		var v45: C8 := new C8(v43);
		var v46 := DC62(v45);
		match (DC64(v46)) {
			case DC63(cf116, cf117, cf118) =>
				var v47 := DC35(v0);
				var v48: C5 := new C5(v47.cf65, cf117);
				var v49: seq<C5> := [v48];
				var v50: map<seq<C5>, bool> := map[v49 := v0];
				v50 := v50[v49[v2 := v48] := v0];
				globalState.f24 := v0 && v0;
				var v51: multiset<bool> := multiset{true};
				var v52: map<int, multiset<bool>> := map[cf118 + i3 := v51];
				v52 := v52[cf118 := multiset{cf117, v0}];
				var v53: map<multiset<int>, bool> := map[v43 := v0];
				globalState.f15 := v53 == v53;
			case DC62(cf115) =>
				var v54: array<int> := new int[10];
				var v55 := DC10(v54);
				v54 := v55.cf21;
				globalState.f6 := fm22(i3, v0, globalState) + (seq(219, i4  => ('v')));
				var v56: array<bool> := new bool[7](i5 => v0);
				v56[919] := v0;
				v56[919], globalState.f2, globalState.f2, globalState.f15 := v0, v0, v0, v0;
				var v57: multiset<array<int>> := multiset{v54, v54};
				var v58: map<bool, multiset<array<int>>> := map[true := v57];
				var v59 := DC65(v57);
				var v60: map<int, multiset<array<int>>> := map[|"hmiq"| := v57[v54 := |(if (v56[919] in v58) then v58[v56[919]] else v59.cf120)|] - multiset{v54, v54, v54}];
				var v61: map<D3, int> := map[v55 := v2];
				var v62: C4 := new C4(v61, v56[919], v0);
				var v63 := DC38(-v2, i3, v0);
				v60 := v60[|(map[v62 := v0])[v62 := v62.fm2(true, v63.cf69, globalState)]| := v57];
			case DC64(cf119) =>
				var v64, v65 := v45.m4(globalState);
				globalState.f12 := |fm22(430, v65, globalState)|;
				globalState.f16 := i3;
				v64 := v64;
		}
		
		var v66: array<int> := new int[12](i6 => i6 / v2);
		var v67: map<bool, int> := map[v0 := 0x29c];
		v66[78] := v2 + (if (v0 in v67) then v67[v0] else v2);
		v66[78] := -fm39(globalState) + (fm39(globalState) % v2);
	}
	r0 := p0;
	r1 := fm25(v0, v0 ==> v0, p0, globalState);
	r2 := v2 > v2;
	var v68: multiset<bool> := multiset{true, v0};
	r3 := map[v68 + v68 := !(v1[v2 := v0] < v1)];
}
trait T0 {
	function fm1(globalState: GlobalState): int 
	function fm2(p0: bool, p1: int, globalState: GlobalState): bool 
}

trait T1 extends T0 {
	const f26 : bool
	const f27 : bool
	method m1(p0: bool, p1: seq<string>, p2: seq<bool>, globalState: GlobalState) 
}

trait T2 extends T1 {
	function fm3(p0: string, p1: seq<bool>, p2: int, globalState: GlobalState): seq<bool> 
	function fm4(p0: D0, p1: int, p2: int, globalState: GlobalState): int 
	method m2(p0: bool, p1: T0, p2: bool, globalState: GlobalState) returns (r0: char, r1: int) 
}

trait T3 {
	const f29 : multiset<int>
	method m3(p0: int, p1: int, p2: string, p3: int, globalState: GlobalState) returns (r0: bool, r1: bool, r2: array<int>, r3: T0) 
	method m4(globalState: GlobalState) returns (r0: bool, r1: bool) 
}

class C0 {
	const f32 : bool
	constructor (f32 : bool) {
		this.f32 := f32;
	}
	
	function fm6(p0: char, p1: bool, globalState: GlobalState): bool {
		f32
	}
	function fm7(p0: map<bool, int>, p1: seq<int>, p2: bool, globalState: GlobalState): bool {
		true
	}
}

class C1 extends T2 {
	const f36 : int
	var f37 : int
	constructor (f36 : int, f37 : int, f26 : bool, f27 : bool) {
		this.f36 := f36;
		this.f37 := f37;
		this.f26 := f26;
		this.f27 := f27;
	}
	
	function fm3(p0: string, p1: seq<bool>, p2: int, globalState: GlobalState): seq<bool> {
		([f26] + [f26]) + ([f27] + [f27, f26])
	}
	function fm4(p0: D0, p1: int, p2: int, globalState: GlobalState): int {
		f36
	}
	function fm1(globalState: GlobalState): int {
		f37
	}
	function fm2(p0: bool, p1: int, globalState: GlobalState): bool {
		f26
	}
	method m2(p0: bool, p1: T0, p2: bool, globalState: GlobalState) returns (r0: char, r1: int) {
		var i0 := 0;
		while (true)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			globalState.f12 := f37;
			globalState.f24 := p2 ==> fm2(f26, f37, globalState);
			var v0: array<seq<int>> := new seq<int>[10](i1 => [f36, f37]);
			var v1: seq<int> := [f37, f36];
			v0[661] := v1;
			v0[661] := v1;
			globalState.f2 := p2;
		}
		var v2: map<bool, bool> := map[!f27 := p2 ==> p2];
		var v3: array<C0> := new C0[5];
		var v4: C0 := new C0(f27);
		v3[845] := v4;
		var v5: set<int> := {f37, 0x3cf};
		var v6: array<bool> := new bool[9];
		var v7: multiset<int> := multiset{f36};
		var v8: map<array<bool>, int> := map[v6 := if (f36 in v7) then v7[f36] else f37];
		f37, v2, v3[845] := (-0x152 - |v5|) + fm4(fm19(f36, |v8|, globalState), -0x90, f36, globalState), v2, v4;
		var v9 := DC7(p2, f27, -0x364, f36);
		var v10: map<int, int> := map[f37 := v9.cf17];
		var v11 := DC6(v10);
		for i2 := f37 to |{v11}| {
			v5 := v5;
			var v12: array<int> := new int[6];
			v12[100] := 0x31c % f36;
			v12[100] := i2 % i2;
			var v13 := 'a';
			var v14: array<char> := new char[3] [v13, v13, v13];
			var v15: map<int, array<char>> := map[v12[100] := v14];
			v15 := v15[i2 := if (p0) then v14 else v14];
			v6[955] := f26;
			var v16 := "tniumeoec";
			var v17: map<int, multiset<int>> := map[-|v16| := v7];
			v6[955] := v12[100] > -|((if (f36 in v17) then v17[f36] else v7) * v7)|;
		}
		f37 := f37;
		var v18: seq<bool> := [f26];
		var v19: seq<int> := [|v18|];
		var v20 := DC20(v19);
		var v21: map<bool, seq<int>> := map[f26 := v20.cf40 + [|v18|]];
		var v22: array<seq<multiset<D2>>> := new seq<multiset<D2>>[1](i3 => [multiset{DC9(f36, f37, v4.f32)}, multiset{DC9(f36, f36, true), DC9(f36, f36, f27), DC9(f37, f37, v4.f32)}] + [multiset{DC9(f36, |[|map[multiset(v18) := v4.f32]|]|, v4.f32), DC9(|v18|, f37, p0)}]);
		var v23 := DC9(f36, f37, f27);
		var v24: multiset<D2> := multiset{v23};
		var v25: multiset<bool> := multiset{p2};
		var v26: seq<multiset<D2>> := [v24, multiset(([v23])[-|v25| := fm20(f36, p2, f36, v9.cf17, globalState)])];
		v22[193] := v26 + (seq(0x3e, i4  => (v24)));
		var v27: map<int, multiset<D2>> := map[f37 := v24];
		var v28 := DC0(f26);
		v21, v22[193] := v21, v26[-0x2cc := if (f37 in v27) then v27[f37] else v24] + ((v26[fm4(v28, f37, f36, globalState) := v24])[f36 := v24] + [v24, v24]);
		var v29: array<multiset<bool>> := new multiset<bool>[29](i5 => v25);
		v29 := v29;
		var v30: map<bool, char> := map[p2 := 'c'];
		var v31 := 'm';
		r0 := if (v4.f32 in v30) then v30[v4.f32] else v31;
		var v32: seq<char> := [v31, v31];
		var v33: map<int, map<int, int>> := map[|v32| := v10[f37 := f36]];
		r1 := |v33|;
	}
	method m1(p0: bool, p1: seq<string>, p2: seq<bool>, globalState: GlobalState) {
		var v0: array<bool> := new bool[3];
		v0[254] := f26;
		var v1: array<int> := new int[11];
		var v2: set<bool> := {f26};
		var v3: seq<set<bool>> := [v2];
		var v4: map<int, bool> := map[|v3[0x352]| := f26];
		var v5: map<bool, int> := map[f26 := -|v4|];
		var v6: multiset<int> := multiset{if (p0 in v5) then v5[p0] else f37, f37};
		var v7 := "k";
		var v8 := 't';
		v1[592] := if (|v7[561 := v8]| in v6) then v6[|v7[561 := v8]|] else fm1(globalState);
		var v9: array<char> := new char[24](i0 => v8);
		v0[254], v1[592], v9 := p0, f37, v9;
		var v10: multiset<bool> := multiset{p0, f27};
		var v11: seq<int> := [f37, |v7|];
		v1[592], v1[592], v1[592], globalState.f22, v0[254] := f37, |((v10 + v10) + v10)|, -0x2ce, if (if (p0) then p0 else !p0) then v11 != v11 else (if (f26 in v5) then v5[f26] else |v5|) < 349, f26;
		for i1 := f37 to v1[592] {
			var v12 := new C0(fm21(f37, globalState) >= v10);
			var v13: set<int> := {f37, v1[592], f36, v1[592]};
			var v14: map<set<int>, array<bool>> := map[v13 := v0];
			v14 := v14;
			v1[592] := i1;
			globalState.f15 := p2[f37 / |v2|];
		}
		var v15: map<set<bool>, bool> := map[v2 := !f26];
		v15 := if (p0) then map[v2 := v0[254]] else v15;
		for i2 := v1[592] to f37 * -v1[592] {
			var v16 := new C0(f27);
			globalState.f12 := v1[592] + i2;
			var v17 := DC4(v9, v1[592], |multiset{v8, v8}|, fm22(|(seq(0x158, i3  => (v8)))|, f27, globalState), seq(0xd1, i4  => (v8)));
			v17 := if (f26) then v17 else DC4(v9, |v17.cf11|, f37, v7, v7);
			globalState.f15 := false;
		}
		v1[592] := f37;
	}
}

class C2 extends T2 {
	const f35 : seq<seq<int>>
	constructor (f35 : seq<seq<int>>, f26 : bool, f27 : bool) {
		this.f35 := f35;
		this.f26 := f26;
		this.f27 := f27;
	}
	
	function fm3(p0: string, p1: seq<bool>, p2: int, globalState: GlobalState): seq<bool> {
		if ({|map[{|map[|multiset{|(set v0 : int | v0 in (seq(0x225, i0  => (-|map[0x24d := true]|))) :: (v0 / |[|{false, false}|]|))|}| := f27]|, |{true}|} := |[f26, f27, f26]|]|} == {DC1(false, f26, 0x340, f26).cf3}) then [f26, f27] else [f26, f26]
	}
	function fm4(p0: D0, p1: int, p2: int, globalState: GlobalState): int {
		|(match DC8() {
			case DC7(cf14, cf15, cf16, cf17) => map[f27 := multiset{'t', 'g'}]
			case DC8() => map[f26 := multiset{'g'}]
			case DC9(cf18, cf19, cf20) => map[false := multiset{'e', 'b'}] + map[false := multiset{'s'}]
			case DC6(cf13) => map[f26 := multiset{'x', 'b', 'r'}]
		})|
	}
	function fm1(globalState: GlobalState): int {
		|multiset{[DC19(|[f26]|, multiset{f27}, f27, map[f27 := |(seq(-0x2ec, i0  => (DC12({f26}))))|]).cf36, -153]}|
	}
	function fm2(p0: bool, p1: int, globalState: GlobalState): bool {
		false
	}
	function fm16(p0: int, p1: seq<D2>, p2: int, p3: int, globalState: GlobalState): bool {
		match DC17(|"uot"|, 884, |multiset([f26])|, -0x17e, |{f26, f27, true}|) {
			case DC17(cf30, cf31, cf32, cf33, cf34) => f26
			case DC16(cf29) => !f26
		}
	}
	function fm17(p0: int, p1: int, p2: int, p3: multiset<set<int>>, globalState: GlobalState): bool {
		f26
	}
	method m2(p0: bool, p1: T0, p2: bool, globalState: GlobalState) returns (r0: char, r1: int) {
		r1 := fm1(globalState);
		var v0: set<bool> := {p0};
		var v1 := -0x151;
		var v2 := "ajxqysp";
		globalState.f16, v0, r1 := -((v1 * |v2|) - v1), fm18(v0, globalState) - v0, v1;
		var v3: seq<bool> := [p0];
		var v4: seq<seq<bool>> := [v3[|v3| := p2], v3 + v3];
		v4 := (v4 + v4) + (seq(0x237, i0  => (v3)))[v1 := v3];
		var v5: array<bool> := new bool[18];
		v5[800] := true;
		var v6: map<multiset<int>, bool> := map[multiset{v1, |v2|, v1, v1} := f26];
		var v7: multiset<int> := multiset{v1, |v2|};
		var v8: map<bool, int> := map[false := |v2|];
		var v9: map<int, bool> := map[|v8| := !false];
		v5[800] := if ((v7 + multiset{|v9|}) in v6) then v6[v7 + multiset{|v9|}] else f26;
		var v10 := 'p';
		var v11: map<char, char> := map[v2[v1] := v10];
		var v12: seq<map<char, char>> := [v11];
		v12 := v12;
		var v13: seq<int> := [-0x120, |v7|, 0xb4, v1, |"flhymnjo"|];
		var v15: array<seq<int>> := new seq<int>[18] [v13, seq(220, i1  => (v1)), v13 + v13, v13, v13, (if (!v5[800]) then v13 else v13)[v1 := v1], v13 + v13, seq(0x29f, i2  => (|v7|)), v13, ([|v2|])[|(set v14 : char | v14 in v11 :: (v14))| := v1], v13, v13, v13, seq(0x1bc, i3  => (v1)), v13, v13[v1 := v1], v13, v13];
		var v16: set<int> := {|v7|};
		v15 := if (-v1 in v16) then v15 else v15;
		r0 := v10;
		r1 := -v1 * -v1;
	}
	method m1(p0: bool, p1: seq<string>, p2: seq<bool>, globalState: GlobalState) {
		var v0: array<int> := new int[1];
		var v1: multiset<array<int>> := multiset{v0};
		var v2 := 0xf0;
		var v3: map<int, array<int>> := map[v2 := v0];
		var v4 := DC0(p0);
		var v5 := 'a';
		var v6: multiset<char> := multiset{v5, v5, v5, v5};
		var v7 := "cudjrcpt";
		var v8: map<int, multiset<array<int>>> := map[|v7| := multiset{v0, v0, v0, v0, v0}];
		var v9: array<multiset<array<int>>> := new multiset<array<int>>[28] [v1, multiset{if (fm4(v4, v2, |v6|, globalState) in v3) then v3[fm4(v4, v2, |v6|, globalState)] else v0}, v1, v1, v1, v1, v1, multiset{v0, v0}, v1, v1, v1, v1[v0 := 892] - v1, v1 * v1, v1, v1, v1, v1, v1 + v1, v1, v1, v1 + multiset{v0, v0, v0, v0, v0}, v1, v1 - v1, v1, v1, v1, if (--v2 in v8) then v8[--v2] else multiset{v0, v0, v0}, v1];
		v9[532] := v1;
		v9[532] := v1;
		var i0 := 0;
		while (!(v6 < (multiset{v5} + multiset{v5})))
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v10: map<bool, int> := map[f26 := v2];
			v10 := v10[f27 := v2];
			var v11 := new C0(p0);
			v0[433] := -(-fm1(globalState) / v2);
			v0[433] := v2;
			var v12: array<int> := new int[23](i1 => i1 % v0[433]);
			var v13 := DC10(v12);
			v13 := DC10(v12);
		}
		var v14: set<int> := {v2};
		var v15: multiset<set<int>> := multiset{v14};
		var v16: multiset<bool> := multiset{fm17(|v7|, -v2, v2, v15, globalState), f27, f26, !!false};
		globalState.f12 := (|v7| * (if (!p0 in v16) then v16[!p0] else v2)) - fm1(globalState);
		var v17: map<bool, bool> := map[p0 := f26];
		m12(v2, |v17|, globalState);
		var v18: T2 := new C1(v2, |v7|, p0, f26);
		var v19 := DC18(v18);
		match (v19) {
			case DC19(cf36, cf37, cf38, cf39) =>
				globalState.f15 := v18.f26;
				globalState.f12 := v2 + (v2 - 648);
				var v20: map<int, int> := map[|[f26]| := cf36];
				var v21: array<array<bool>> := new array<bool>[28];
				var v22: map<array<array<bool>>, bool> := map[v21 := v18.f27];
				var v23: map<bool, map<bool, bool>> := map[if (v21 in v22) then v22[v21] else v18.f27 := v17];
				cf36 := v18.fm4(v4, |v20| * v2, |v23|, globalState);
				v5 := v5;
			case DC18(cf35) =>
				var v24: array<bool> := new bool[18](i2 => DC14(v2, |v7|, v18.f27).cf27);
				globalState.f6, v24 := "htwwcultw" + "jlkne", v24;
				globalState.f12 := v2 * 970;
				var v25: array<set<int>> := new set<int>[19](i3 => v14);
				v25[648] := set v26 : int | v26 in multiset(fm23(v5, globalState)) :: (v26 / |{false}|);
				v25[648] := v14;
				globalState.f16 := -v2;
		}
		
		v5 := if (v18.fm2(p0, v2, globalState)) then v5 else if (p0) then v5 else v5;
	}
	method m11(p0: seq<D1>, p1: array<bool>, globalState: GlobalState) returns (r0: string) {
		var v0 := DC0(f26);
		var v1 := "euqbmr";
		var v2: map<string, bool> := map[v1 := f27];
		var v3 := 671;
		var v4: array<int> := new int[14];
		var v5: map<int, array<int>> := map[fm4(v0.(cf0 := f26), |v2|, v3, globalState) := v4];
		v5 := v5[v3 := v4];
		var v6: map<string, int> := map[v1 := v3];
		for i0 := |v6| to v3 {
			var v7: map<char, bool> := map[fm24(v1, f26, v3, globalState) := f27];
			v7 := v7 + v7;
			var v8 := 'n';
			var v9, v10, v11, v12 := m0(v8, globalState);
			v4[234] := v3;
			var v13 := DC16(v1 + v1);
			var v14 := DC9(i0, i0, f27);
			var v15: multiset<set<int>> := multiset{{v3, 201}};
			globalState.f12, v4[234], v13, globalState.f2 := i0, i0, if (v14.cf19 != v3) then v13 else DC16(v1), fm17(i0, -i0, v3, v15, globalState);
			var v16: array<bool> := new bool[13](i1 => f27);
			var v17: map<bool, bool> := map[!f27 := f27];
			v16 := if (if (v11 in v17) then v17[v11] else true) then v16 else p1;
		}
		var i2 := 0;
		while (false <==> f26)
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			var v18 := 'g';
			var v19: set<bool> := {f27};
			var v20: map<int, set<bool>> := map[|([v18, v18] + v1)| := v19 * v19];
			v20 := v20;
			globalState.f22 := f26;
			globalState.f12 := v3;
			globalState.f16 := 308;
		}
		var v21: multiset<int> := multiset{v3};
		var v22: set<bool> := {f27, false};
		p1[156] := v21 < multiset{|v22|, v3};
		var v23: map<bool, bool> := map[f27 := false];
		var v24: multiset<bool> := multiset{if (f27 in v23) then v23[f27] else false};
		var v25: array<set<int>> := new set<int>[6](i3 => {|[f27]|});
		var v26: map<int, multiset<bool>> := map[v3 := v24];
		var v27: map<set<bool>, array<set<int>>> := map[v22 := v25];
		p1[156], v24, globalState.f15, v25, v3 := f26, if (v3 in v26) then v26[v3] else multiset{true, f27}, f27 && f26, if (true) then v25 else if ({f27} in v27) then v27[{f27}] else v25, -v3;
		var v28: set<int> := {v3};
		var i4 := 0;
		while ((v3 + v3) !in v28)
			decreases 100 - i4
		{
			if (i4 >= 100) {
				break;
			}
			
			i4 := i4 + 1;
			v4[960] := v3;
			var v29 := DC12(v22);
			var v30 := 'c';
			var v31: seq<bool> := [f26, false, true];
			var v32: map<bool, seq<bool>> := map[f26 := v31];
			globalState.f16, v4[960], v29 := fm4(v0, |map[p1[156] := v30]|, |v32|, globalState) - v3, v3, v29;
			var v33: C1 := new C1(v3, v3, f27, p1[156]);
			var v34: seq<C1> := [v33, v33, v33];
			globalState.f22, v33, globalState.f22, v30 := fm22(v3, f26, globalState) <= (v1 + v1), v34[v3], v30 in v1, v30;
			var v35 := new C0(p1[156]);
			var v36 := new C1(v33.f36, v3, v35.f32, p1[156]);
		}
		for i5 := v3 to fm1(globalState) {
			globalState.f19 := v1[if (f27) then v3 else i5];
			var v37 := 'n';
			if (-0x11f <= |(("lbthv")[i5 := v37] + v1)|) {
				var v38: map<int, bool> := map[i5 := f27];
				var v39 := DC9(|v38|, v3, false);
				var v40: seq<D2> := [v39, v39];
				var v41 := new C1(0x85, v3, fm16(i5, v40, |{i5}|, |v1|, globalState), f26);
				var v42: map<int, seq<bool>> := map[v41.f36 := [true, p1[156]]];
				var v43: seq<bool> := [true];
				var v44: map<seq<bool>, array<bool>> := map[if (v41.f36 in v42) then v42[v41.f36] else v43 := p1];
				v44, globalState.f12 := DC21(v44).cf41 + v44, (if (f26) then i5 else v3) * (v41.f36 / i5);
				var v45 := new C1(v41.f37, v41.f37 - i5, !f27, false && f26);
				var v46: map<bool, int> := map[f26 := v41.f36];
				v23 := v23[|v46| != v41.f37 := p1[156]];
				var v47: array<D0> := new D0[21];
				var v48: map<char, array<D0>> := map[v37 := v47];
				v48 := v48[v37 := v47];
			} else {
				var v49: multiset<set<int>> := multiset{v28};
				var v50: seq<bool> := [fm17(v3, i5, v3, v49, globalState)];
				globalState.f16 := fm4(v0, |(fm3(v1, v50, |map[v3 := p1[156]]|, globalState) + v50)|, 0x360, globalState);
				globalState.f16 := -fm4(v0, v3, 400, globalState);
				globalState.f12 := -(if (p1[156]) then v3 else v3) % |(v28 * v28)|;
				globalState.f12 := i5;
				globalState.f24 := p1[156];
			}
			
			var v51: seq<bool> := [!false, p1[156], true];
			var v52: map<seq<bool>, bool> := map[v51 + [p1[156], p1[156]] := f26];
			v52 := v52;
			v51 := [f27] + (v51 + [!f26]);
		}
		r0 := "bmiwf";
	}
	method m12(p0: int, p1: int, globalState: GlobalState) {
		var v0: multiset<int> := multiset{p1};
		var v1 := new C1(p1, p1, if (f27) then f27 else f27, multiset{p1, p1} !! v0);
		var v2: array<bool> := new bool[17](i0 => f26);
		var v3: seq<bool> := [f26, f26, f27];
		var v4: map<bool, seq<bool>> := map[f26 := v3];
		var v5 := DC13(v4);
		globalState.f24, v2, globalState.f22, v5 := ([f27, f27, f27, f27] != [f27, f26]) ==> f26, v2, f27, v5;
		var v6 := "adlk";
		for i1 := |v6| to v1.f37 / -|v3| {
			globalState.f12 := i1;
			var v7: array<int> := new int[12](i2 => i2 / v1.f36);
			v7[313] := v1.f37;
			var v8 := 's';
			var v9: seq<array<bool>> := [v2, v2];
			globalState.f6, v7[313], globalState.f22, globalState.f24, globalState.f2 := v6[|v3| := v8], --|v9|, f26, f27, !v1.fm2(f26, if (true) then v1.f37 else v1.f37, globalState);
			var v10: set<bool> := {!f27, f27, !false};
			v10 := (v10 * v10) * (v10 + v10);
			var v11: array<array<bool>> := new array<bool>[27] [v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2];
			var v12: seq<array<array<bool>>> := [v11, v11, v11];
			globalState.f22, globalState.f5 := f27 <==> f26, v12[v1.f36];
		}
		v2[575] := f27;
		v2[575] := ((seq(0x32f, i3  => ('r'))) + v6) < v6;
		globalState.f16 := -p0;
		v1.f37 := |(v6 + v6)|;
	}
}

class C3 extends T1 {
	const f34 : multiset<seq<int>>
	constructor (f34 : multiset<seq<int>>, f26 : bool, f27 : bool) {
		this.f34 := f34;
		this.f26 := f26;
		this.f27 := f27;
	}
	
	function fm1(globalState: GlobalState): int {
		if (f26) then |DC16("yoqe").cf29| + |"jmfpgf"| else |[f27, f26, f26]|
	}
	function fm2(p0: bool, p1: int, globalState: GlobalState): bool {
		false
	}
	function fm12(p0: int, globalState: GlobalState): int {
		(if (f27) then 84 else 719) + (|[f26]| % 676)
	}
	function fm13(p0: seq<multiset<char>>, p1: char, p2: int, globalState: GlobalState): int {
		|(seq(0x6d, i0  => (|DC16("r").cf29|)))| - (|map[f27 := map[956 := f26]]| * -|map[false := 'k']|)
	}
	method m1(p0: bool, p1: seq<string>, p2: seq<bool>, globalState: GlobalState) {
		var v0: array<int> := new int[21];
		var v1 := 0x119;
		var v2: map<int, array<int>> := map[v1 := v0];
		var v3: array<array<int>> := new array<int>[20] [v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, if (v1 in v2) then v2[v1] else v0, v0, v0, v0, v0, v0, v0];
		v3[142] := v0;
		var v4: seq<array<int>> := [v0, v0];
		v3[142] := v4[0x396];
		var v5: map<bool, bool> := map[f27 := f26];
		var v6: seq<map<bool, bool>> := [v5, map[p0 := f27], v5, v5, map[f27 := f26]];
		var v7: seq<int> := [v1, v1, v1];
		var v8: map<seq<map<bool, bool>>, seq<int>> := map[v6 := v7];
		var v9 := "qgla";
		var v10: map<string, seq<map<bool, bool>>> := map[v9 := v6];
		v8 := v8[(if (v9 in v10) then v10[v9] else v6) + [v5] := v7 + v7];
		var v11: seq<bool> := [f27, f26, p0, v1 == |v9|, p0];
		v11 := v11 + v11;
		var v12: seq<multiset<char>> := [multiset{'y'}];
		v1 := -fm13(v12 + v12, v9[v1], v1 + v1, globalState);
		v9 := v9 + v9;
		var v13: set<bool> := {f27};
		var v14: map<int, int> := map[-|v13| := v1];
		for i0 := -(if (v1 in v14) then v14[v1] else v1) to v1 {
			globalState.f15 := i0 < (i0 / |v11|);
			var v15: map<int, bool> := map[i0 := f27];
			var v16: multiset<int> := multiset{v1};
			var v17: map<int, string> := map[|v15| + |v16| := v9];
			var v18 := 'e';
			v17 := v17[|[fm14(f26, i0, v1, globalState), v18, v18, v18]| := (seq(934, i1  => ('x'))) + v9];
			v1 := |((if (p0) then [i0] else [i0]) + (fm15(globalState) + v7))|;
			var v19: T2 := new C1(|"ylv"|, |v9|, p0, f27);
			var v20: map<T2, int> := map[DC18(v19).cf35 := i0];
			var v21: array<map<T2, int>> := new map<T2, int>[12] [v20, v20 + map[v19 := |v4|], v20, v20 + map[v19 := v1], v20, v20, map[v19 := |v16|] + v20, v20, v20, v20, v20, v20];
			v21[318] := map[v19 := 196];
			v21[318] := map[v19 := i0] + (v20 + v20);
		}
	}
}

class C4 extends T1 {
	const f33 : map<D3, int>
	constructor (f33 : map<D3, int>, f26 : bool, f27 : bool) {
		this.f33 := f33;
		this.f26 := f26;
		this.f27 := f27;
	}
	
	function fm1(globalState: GlobalState): int {
		match DC12({f27}) {
			case DC12(cf23) => |[!f26, f26, f26]|
		}
	}
	function fm2(p0: bool, p1: int, globalState: GlobalState): bool {
		f27
	}
	method m1(p0: bool, p1: seq<string>, p2: seq<bool>, globalState: GlobalState) {
		var v0 := 0x20;
		globalState.f12 := v0;
		for i0 := v0 to v0 {
			if (true) {
				var v1: multiset<int> := multiset{-i0};
				var v2: multiset<bool> := multiset{v1 > multiset{|p2|}, f26 ==> p0};
				globalState.f12 := if (f26 in v2) then v2[f26] else i0 + v0;
				var v3: array<D2> := new D2[13](i1 => DC7(f26, f27, v0, v0));
				v3 := v3;
				var v4: map<int, int> := map[v0 := 0x88];
				globalState.f12 := i0 + (if (i0 in v4) then v4[i0] else |v2|);
				var v5 := new C0(f27);
				globalState.f15 := !!(v0 >= i0);
			} else {
				var v6: map<bool, seq<bool>> := map[f27 := p2];
				var v7 := DC13(v6);
				v6 := v7.cf24 + v6;
				var v8 := DC3(p2);
				var v9: C0 := new C0(p0);
				var v10: seq<C0> := [v9];
				var v11: map<D1, bool> := map[v8 := v9 in v10];
				var v12: map<bool, bool> := map[f26 := f27];
				var v13: seq<map<D1, bool>> := [v11];
				var v15: multiset<D1> := multiset{v8, DC3(p2)};
				v11 := map[v8 := if (v9.f32 in v12) then v12[v9.f32] else f27] + (v13[i0] + (map v14 : D1 | v14 in v15 :: (v14) := (v9.f32)));
				globalState.f12, globalState.f12 := v0, v0;
				var v16 := DC14(|p2|, v0, true);
				var v17: array<char> := new char[9];
				var v18: seq<array<char>> := [v17, v17];
				globalState.f12 := (v16.cf26 + |v18|) - v0;
				var v19 := "ldus";
				globalState.f15 := i0 < (|v19| * i0);
			}
			
			var v20: array<int> := new int[9](i2 => i2 + DC14(v0, v0, true).cf26);
			var v21 := 'o';
			var v22: set<char> := {v21};
			var v23: map<set<char>, int> := map[v22 := v0];
			var v24 := "cdyf";
			var v25: seq<int> := [v0, i0];
			var v26: multiset<seq<int>> := multiset{v25, v25};
			var v27: T1 := new C3(v26, p0, p0);
			var v28: array<int> := new int[14] [i0 * |"yghsv"|, |(map[v20 := |v23|] + map[v20 := i0])|, v0, v0, -v0, -i0, i0, |(v24 + v24)|, |map[seq(546, i3  => (v0)) := v27]|, i0, v0, i0, i0, v0];
			var v29 := DC10(v28);
			v20 := v29.cf21;
			var v30 := DC19(i0, multiset{!p0, f26}, false, map[false := v0]);
			var v31: multiset<bool> := multiset{v27.f27};
			var v32: array<D7> := new D7[12] [if (f27) then v30 else v30, v30.(cf37 := v31, cf38 := f26, cf36 := |(seq(0xe6, i4  => (v21)))|), v30, v30, v30, v30, v30, v30, v30, v30, if (p0) then v30 else v30, v30];
			var v33: map<bool, int> := map[v27.f26 := 0x27d];
			v32[896] := DC19(v0, v31, p0, v33);
			v32[896] := v30;
			var v34: array<bool> := new bool[4] [v27.f26, v27.f27, false, v27.f26];
			v34[510] := if (true) then v27.fm2(v27.f26, i0, globalState) else v27.f27;
			var v35: array<D1> := new D1[29];
			var v36: map<bool, array<D1>> := map[true := v35];
			var v37: map<int, int> := map[i0 := i0];
			var v38: map<map<bool, array<D1>>, int> := map[v36 := |v37|];
			v34[510], globalState.f6, globalState.f12, globalState.f22 := (v0 != i0) ==> f26, "fxf", v0 % v0, (v0 != (if (v36 in v38) then v38[v36] else v25[i0])) || (i0 <= v0);
		}
		var v39 := 'q';
		var v40: set<bool> := {f26};
		var v41: array<char> := new char[21] [v39, 'f', 'l', fm14(p0, v0, |v40|, globalState), v39, v39, 'v', v39, v39, v39, v39, v39, v39, v39, 'u', v39, 'o', v39, v39, 'm', v39];
		var v42: map<int, bool> := map[v0 := false];
		var v43: seq<map<int, bool>> := [v42, v42];
		var v44 := DC7(true, false, v0, |multiset(v43)|);
		var v45: map<seq<D2>, int> := map[[v44, v44] := v0];
		var v46: seq<D2> := [v44, v44];
		var v47 := "baa";
		var v48 := DC4(v41, v0, if (v46 in v45) then v45[v46] else v0, v47, v47[-|map[v0 := v0]| := v39]);
		match (v48) {
			case DC4(cf7, cf8, cf9, cf10, cf11) =>
				var v49: seq<int> := [v0, cf8, cf9];
				var v50: seq<seq<int>> := [v49, v49, v49[|v49| := cf8], [v0], v49];
				var v51: C2 := new C2(v50, false, f26);
				var v52: seq<C2> := [v51];
				globalState.f12 := if (f26) then |p2[-fm1(globalState) := !f26]| / cf9 else |(v52 + v52)|;
				var v53: map<int, int> := map[|p2| := cf8];
				var v54 := DC0(f26);
				v53 := v53[cf8 := v51.fm4(v54, v0, cf9, globalState)];
				globalState.f2 := !f26;
				if (!f27) {
					var v55: array<seq<int>> := new seq<int>[2];
					v55[59] := fm15(globalState) + (seq(0xbf, i5  => (cf8)));
					cf9, globalState.f15, globalState.f6, v55[59] := -v51.fm4(v54, cf9, v0, globalState), p2[v0], cf10, fm23(v39, globalState);
					globalState.f19 := v39;
					globalState.f6 := v47[-v0 := v39] + (seq(472, i6  => (v39)));
					var v56: set<string> := {"riqaxgwa", "jqtgis"};
					globalState.f24 := v51.fm2(v56 > v56, cf9, globalState);
					globalState.f15 := f26;
				} else {
					var v57: array<map<D3, int>> := new map<D3, int>[6];
					var v58: array<D3> := new D3[11];
					var v59: map<bool, array<char>> := map[!f26 := cf7];
					var v60: seq<array<char>> := [cf7];
					var v61: array<array<char>> := new array<char>[15] [v41, cf7, cf7, cf7, cf7, v41, v41, cf7, v41, cf7, if (true) then cf7 else v41, cf7, if (true in v59) then v59[true] else cf7, cf7, v60[-0x217]];
					v61[966] := v41;
					v57, v58, v61[966] := if (p0) then v57 else v57, v58, v41;
					cf9 := if (f27) then cf9 else |cf10| * cf9;
					var v62: set<int> := {-0x3bb};
					var v63 := DC9(|v62|, v0, false);
					var v64: multiset<bool> := multiset{!f27, f26, f26};
					var v65: seq<seq<bool>> := [p2];
					var v66: array<int> := new int[12] [cf8, v63.cf18, |v64|, -cf9, |multiset{f27, p0, p0, false}|, v0, cf8, cf8, |cf11|, v0, |(p2 + v65[v51.fm1(globalState)])|, v0];
					v66[232] := v0 * cf9;
					globalState.f16, v66, globalState.f15, globalState.f2, v66[232] := cf9, v66, !(if (f26) then p0 else if (cf8 in v42) then v42[cf8] else p0) ==> (-cf9 < cf8), p0, cf9 % v0;
					var v67: seq<bool> := [p0];
					v67 := v67;
					var v68 := DC19(v0, v64[f27 := cf9], f26, map[p0 := cf9]);
					var v69: multiset<int> := multiset{cf8, -cf9};
					var v70: map<bool, bool> := map[f26 := v51.fm2(p0, v0, globalState)];
					var v71: array<bool> := new bool[26] [false, false, |fm22(if (cf8 in v53) then v53[cf8] else cf9, f27, globalState)| > cf8, f27, f27, v68.cf38, f27, p0, f27, v47 <= v47, multiset(v49) !! v69, f27, f27, true, cf10 < "h", f26 && p0, f26, f27, f26, if (f26 in v70) then v70[f26] else f27, true, false, f27, f27, !(map[f26 := f27] == v70), !p0];
					v71[0] := false;
					var v72: map<bool, int> := map[f27 := v66[232]];
					globalState.f22, globalState.f16, globalState.f19, v71[0] := p0, fm1(globalState), v39, (if (v51.fm2(!f27, v66[232], globalState)) then p0 else f27) || (|v72| != v0);
				}
				
			case DC5(cf12) =>
				if (f26) {
					var v73: map<int, char> := map[v0 := v39];
					var v74: seq<map<int, char>> := [v73];
					globalState.f15 := |(v74 + v74)| >= v0;
					var v75: multiset<seq<int>> := multiset{[cf12]};
					var v76: C3 := new C3(v75, p2[cf12], f27);
					var v77: seq<int> := [cf12, v0];
					var v78: map<int, int> := map[|(seq(-0x3cb, i7  => (|[v0, v0, -v0]|)))| := v77[v76.fm12(|v47|, globalState)]];
					var v79: map<C3, int> := map[v76 := |v78|];
					var v80: map<map<C3, int>, char> := map[v79 := v39];
					var v81: map<int, map<int, int>> := map[v0 := v78];
					var v82: array<seq<bool>> := new seq<bool>[19] [p2 + p2[cf12 := f26], p2, [!p0], p2, fm25(true, f26, v39, globalState), p2, p2, p2 + p2, fm25(f27, f26, if (v79 in v80) then v80[v79] else v39, globalState), p2, p2, p2, (if (f26) then p2 else [false])[|map[v76.fm12(cf12, globalState) := if (|multiset(fm15(globalState))| in v81) then v81[|multiset(fm15(globalState))|] else map[cf12 := v0]]| := f27], p2, p2, p2[cf12 := f27], p2, p2, fm25(f26, f27, v39, globalState)];
					v82[171] := p2;
					v82[171] := p2;
					var v83: array<int> := new int[20];
					var v84 := DC10(v83);
					var v85: seq<array<int>> := [v84.cf21];
					var v86 := DC9(0x79, v0, false);
					var v87: array<array<int>> := new array<int>[27] [v83, v85[v0], v83, v83, v83, v83, v83, v83, v83, v83, v83, if (v86.cf20) then v83 else v83, v83, v83, v83, v83, if (f26) then v83 else v83, v83, v83, v83, v83, v83, v83, v84.cf21, v83, v83, v83];
					v87 := v87;
					var v88: multiset<bool> := multiset{f27, f27};
					var v89: array<bool> := new bool[25] [!p0 <== !!p0, f27, f26, p0, !p0, p0, f27, f26, p0, false, f27, f27, f27, !(v88 !! multiset{fm2(p0, v0, globalState)}), p0, f26, p0, p0, !f26, f27, f27, f27, p0 && false, f27, |v82[171]| < v0];
					v89[422] := fm2(p0, -cf12, globalState);
					v89[422] := p0;
					var v90: seq<seq<bool>> := [[fm2(p0, v0, globalState), v89[422]], p2, v82[171]];
					var v91: array<map<int, char>> := new map<int, char>[9];
					v91[651] := v73;
					v90, v89[422], globalState.f12, v91[651] := v90, true, if (cf12 in v78) then v78[cf12] else 953, v73;
				} else {
					var v92: map<bool, int> := map[true := v0];
					var v93: seq<map<bool, int>> := [v92];
					var v94: array<bool> := new bool[26] [f27, p0, v47[cf12 := 'e'] < "urnk", p0 ==> fm2(f27, v0, globalState), f27, fm2(p0, cf12, globalState), f27, !f27, !DC14(v0, cf12, f26).cf27, p0, -cf12 >= |p2|, p0, f26, f27, p0, p0, f26, f27, p2[v0], f26, p0, f26, true, p0, f27, v93 != v93];
					v94[26] := p2[v0];
					v94[26] := f27;
					var v95: array<int> := new int[8];
					var v96: map<array<int>, bool> := map[v95 := v94[26]];
					globalState.f22 := |(v96 + map[v95 := p0])| < cf12;
					var v97 := new C1(|v47[v0 := v39]| % v0, fm1(globalState), v94[26], f27);
					var v98 := DC11(v39);
					var v99: seq<int> := [|v40|, v97.f36];
					var v100: map<array<bool>, bool> := map[v94 := fm23(v98.cf22, globalState) <= v99];
					v100 := v100[v94 := !false];
					var v101: T2 := new C1(|[false, f26]|, v97.f37, v94[26], p0);
					var v102 := DC18(v101);
					v95[63] := v0;
					v102, v95[63], v94[26], globalState.f2, v101 := v102, v97.f37, f26, [v0] < v99, v101;
				}
				
				if (!(true ==> p0)) {
					globalState.f24 := f27 !in p2;
					globalState.f22 := f26;
					globalState.f16, v39, globalState.f16 := --v0, 'y', 687;
					globalState.f15 := f26 || p0;
					var v103: array<multiset<map<int, int>>> := new multiset<map<int, int>>[29](i8 => multiset{map[cf12 := v0]});
					v103[369] := fm26(fm2(f26, cf12, globalState), p0, p0, globalState);
					var v104: map<int, int> := map[v0 := cf12];
					var v105: multiset<map<int, int>> := multiset{map[cf12 := v0], v104 + map[-0x105 := cf12], v104};
					v103[369] := v105;
				} else {
					var v106: multiset<bool> := multiset{!fm2(p0, fm1(globalState), globalState)};
					globalState.f2, v106 := (v0 == cf12) <== !!f27, v106;
					var v107: array<seq<bool>> := new seq<bool>[4](i9 => p2);
					v107[74] := p2;
					v0, globalState.f24, v107[74] := cf12, cf12 != v0, (if (true) then p2 else p2) + p2;
					var v108 := DC8();
					var v109: set<D2> := {v108};
					v109 := v109;
					globalState.f15 := f26;
					var v110: array<bool> := new bool[17](i10 => p0);
					var v111: map<seq<bool>, array<bool>> := map[v107[74] := v110];
					var v112 := DC21(v111);
					var v113: seq<D9> := [v112];
					globalState.f24 := v112 in v113;
				}
				
				var v115: array<set<int>> := new set<int>[18](i11 => (set v114 : int | v114 in {v0} :: (v114 + --0x137)) * {|map[v0 := map[v0 := cf12]]|});
				var v116 := DC3(DC3([p0]).cf6);
				var v117: multiset<int> := multiset{|fm18(v40, globalState)|};
				var v118: seq<int> := [0x150];
				globalState.f15, v115, globalState.f12, v116 := f27 && f26, v115, -((v44.(cf16 := |v117|, cf15 := p0)).cf17 / v118[-0x114]), v116;
				var v119: map<bool, bool> := map[f26 := f26];
				v119 := v119[p0 := p0];
			case DC3(cf6) =>
				globalState.f19 := v39;
				var v120: map<string, int> := map["q" := v0];
				globalState.f12 := |v120|;
				var v121, v122, v123, v124 := m0(v39, globalState);
				if (if (!f26) then if (p0) then false else fm2(f26, v0, globalState) else p0) {
					var v125: map<bool, int> := map[p0 := fm1(globalState)];
					globalState.f15 := 954 >= (v0 * (if (v123 in v125) then v125[v123] else 869));
					var v126: map<int, int> := map[v0 := v0];
					v126 := v126[v0 := |v122|];
					v123 := v123;
					var v127: array<bool> := new bool[6](i12 => f26);
					var v128: map<int, array<bool>> := map[0xf3 := v127];
					var v129: array<array<bool>> := new array<bool>[23] [v127, if (v0 in v128) then v128[v0] else v127, v127, v127, v127, v127, v127, v127, v127, v127, v127, v127, v127, v127, v127, v127, v127, v127, v127, v127, v127, v127, v127];
					globalState.f5 := v129;
					v127[492] := p0;
					globalState.f19, v127[492], globalState.f12, v122, globalState.f15 := v121, !p0, v0 * -v0, v122 + (p2 + v122), f26;
				} else {
					var v130: set<int> := {v0, v0};
					var v131: map<int, int> := map[v0 := |v130|];
					var v132: array<int> := new int[13] [v0, v0, v0, v0, v0, 250, v0, |v131|, 492, -v0, v0, v0, v0];
					v132[305] := v0;
					v132[305] := -(v0 + v0);
					globalState.f24 := false;
					var v133 := DC19(v0, fm21(v132[305], globalState), true, map[f27 := v0]);
					globalState.f22 := !(if (f26) then |v47| < |v47| else v133.cf38);
					globalState.f2 := false || f27;
					var v134: map<string, bool> := map[v47 := f26];
					var v135: map<bool, int> := map[f27 := |v40|];
					v134 := fm27(|v135|, v0, globalState) + v134["bd" := v123];
				}
				
		}
		
		var v136: map<int, string> := map[v0 := "ga"];
		var v137 := DC1(f26, f26, |v136|, p0);
		var i13 := 0;
		while (fm28(globalState) == {v137})
			decreases 100 - i13
		{
			if (i13 >= 100) {
				break;
			}
			
			i13 := i13 + 1;
			globalState.f12 := v0;
			globalState.f12 := (v0 % v0) / |(if (f27) then p2 else p2)|;
			var v138: array<int> := new int[18];
			var v139: seq<array<int>> := [v138];
			globalState.f12, globalState.f15 := v0 / -0x158, f27 && (v139 == v139);
			var v140 := DC14(0x323, v0, f27);
			var v141: seq<int> := [v140.cf26, fm1(globalState)];
			v0 := v141[0x123];
		}
		var v142: array<bool> := new bool[14];
		v142[866] := f27;
		v142[866] := v0 >= v0;
		globalState.f6, globalState.f16, v142[866], globalState.f24 := v47, v0 - (-402 / |v40|), p0, !f26;
	}
	method m9(p0: bool, p1: array<bool>, p2: map<bool, bool>, globalState: GlobalState) returns (r0: bool, r1: bool, r2: int, r3: bool) {
		var v0: array<map<int, map<int, int>>> := new map<int, map<int, int>>[28];
		var v1 := -0x3c4;
		var v2: map<int, int> := map[v1 := v1];
		var v3: map<int, map<int, int>> := map[v1 := v2];
		v0[584] := v3;
		var v4: seq<int> := [v1, v1];
		var v5: multiset<bool> := multiset{fm2(p0, |v4|, globalState), f26};
		var v6: seq<bool> := [v5 >= v5];
		var v7: array<array<bool>> := new array<bool>[29];
		v7[644] := p1;
		v0[584], v6, v7[644] := v3, v6, p1;
		v1 := v1;
		globalState.f2 := p0;
		r0 := f26;
		var v8: array<array<string>> := new array<string>[28];
		var v9: array<string> := new string[22];
		var v10: map<int, array<string>> := map[-v1 := v9];
		v8[34] := if (v1 in v10) then v10[v1] else v9;
		var v11: array<int> := new int[23];
		p1[427] := f27;
		v8[34], v11, p1[427], v7[644] := v9, v11, p0, p1;
		p1[427] := !f27;
		var v12 := 'e';
		var v13: map<int, char> := map[v1 := v12];
		var v14: set<map<int, char>> := {v13};
		r0 := !(v14 >= (v14 - (set v15 : map<int, char> | v15 in v14 :: (v15))));
		r1 := !((seq(0x212, i0  => (-v1))) != v4);
		var v16 := DC22(f27);
		r2 := if (v16.cf42) then -v1 else |fm15(globalState)|;
		var v17 := DC8();
		r3 := match v17 {
			case DC7(cf14, cf15, cf16, cf17) => p1[427]
			case DC8() => p1[427]
			case DC9(cf18, cf19, cf20) => if (v6[cf18] in p2) then p2[v6[cf18]] else DC14(0x2f0, 0x1cf, cf20).cf27
			case DC6(cf13) => {v1, v1} == (set v18 : int | (-0x6f <= v18) && (v18 < -734) :: (v18 + v1))
		};
	}
	method m10(p0: bool, p1: int, p2: map<int, seq<int>>, globalState: GlobalState) returns (r0: int, r1: int, r2: int, r3: char) {
		var v0: map<seq<int>, bool> := map[[p1] := !p0];
		var v1: seq<bool> := [p0, fm2(f27, 771, globalState), p0];
		for i0 := |v0| to |(v1 + v1)| {
			var v2: array<seq<seq<bool>>> := new seq<seq<bool>>[19](i1 => [v1, v1, v1]);
			v2[649] := [v1] + [v1];
			var v3: seq<seq<bool>> := [v1, v1];
			v2[649] := v3 + [[p0], v1];
			var v4 := "usmgaxt";
			var v5: multiset<int> := multiset{|(seq(0x202, i2  => (p1)))|, 123, p1, i0, |(seq(0x376, i3  => ('j')))|};
			var v6: set<int> := {p1, |v5|};
			var v7: map<set<int>, bool> := map[v6 := f27];
			var v8: seq<int> := [-0xfc, |v6|];
			var v9: multiset<bool> := multiset{f26};
			var v10: array<int> := new int[20] [|fm29(false, i0, globalState)|, |fm22(|v4|, v1[p1], globalState)| / i0, |v6|, |(v6 - v6)|, i0, p1 - p1, -|(v7 + v7)|, p1 / i0, i0 + p1, |fm30(i0, p0, globalState)| / |v1|, i0, fm1(globalState), i0, i0, -(if (f27) then p1 else i0), p1, i0 % v8[i0], p1, |v4|, |v9|];
			v10[625] := p1;
			v10[625] := p1 - i0;
			var v11 := new C1(fm1(globalState), |v1|, f26, p0);
			v11.f37 := |(v9 - v9)|;
		}
		var v12: multiset<bool> := multiset{f26, f27, fm2(f27, p1, globalState)};
		globalState.f22 := v1[|v12|];
		var v13: array<seq<int>> := new seq<int>[13](i4 => [-DC14(p1, |[-0x156, |(seq(671, i5  => ('k')))|]|, p0).cf26]);
		v13 := v13;
		var v14 := "vgkdg";
		var v16: map<int, set<bool>> := map[602 := {f26}];
		if (|v14| !in (if (p0) then map v15 : int | (-648 <= v15) && (v15 < 0x207) :: (v15 * p1) := ({f26}) else v16)) {
			var v17: array<bool> := new bool[11];
			v17[104] := f26;
			v17[104] := !true;
			var v18: seq<seq<bool>> := [v1];
			v18 := v18;
			var v19: C0 := new C0(f27);
			var v20: map<string, C0> := map["ettqukfnq" := v19];
			var v21: array<C0> := new C0[16] [v19, v19, v19, v19, v19, v19, v19, v19, v19, if (v17[104]) then v19 else v19, v19, v19, if (v14 in v20) then v20[v14] else v19, v19, v19, v19];
			v21[78] := v19;
			v21[78] := v19;
			var v22: array<C3> := new C3[4];
			var v23: seq<int> := [p1, p1];
			var v24: multiset<seq<int>> := multiset{v23, v23, v23};
			var v25: C3 := new C3(v24, v17[104], f26);
			v22[905] := v25;
			v22[905] := v25;
			var v26: map<bool, bool> := map[p0 := p0];
			v26 := v26[[v1] != v18 := f26];
		} else {
			var v27: map<bool, int> := map[!f27 := p1];
			var v28: map<int, int> := map[|v27| := p1];
			var v29: seq<array<seq<int>>> := [v13, v13];
			r2, r2, globalState.f22, r2, v13 := p1, |(v1 + (v1 + v1))|, f26, if (false) then p1 else if (p1 in v28) then v28[p1] else p1, v29[fm1(globalState)];
			var v30: array<bool> := new bool[5];
			var v31 := 's';
			var v32: map<array<bool>, char> := map[v30 := v31];
			var v33, v34, v35, v36 := m0(if (v30 in v32) then v32[v30] else v31, globalState);
			var v37: map<int, seq<bool>> := map[p1 := v34];
			var v38 := new C1(p1, |(if (p1 in v37) then v37[p1] else v34)|, f26, f26);
			var v39, v40, v41, v42 := m0('l', globalState);
			if (DC14(v38.f37, |v12|, v41).cf27) {
				var v43: seq<int> := [-v38.f37, v38.f37];
				var v44: seq<seq<int>> := [v43, v43];
				var v45 := DC0(f26);
				var v46: set<int> := {v38.fm4(v45, |"uc"|, |v40|, globalState)};
				var v47: T2 := new C2(v44, if (v35) then p0 else v41, v46 < v46);
				v47 := new C2(seq(-0x78, i6  => (v43)), v14 <= v14, v47.f26);
				v30 := v30;
				globalState.f2 := v47.f26;
				var v48: set<bool> := {fm2(!v47.f27, v38.f37, globalState)};
				var v49: map<char, int> := map[v33 := v38.f37];
				var v50: map<int, C1> := map[v38.f36 := v38];
				var v51: map<int, string> := map[v38.f36 := v14];
				var v52: array<int> := new int[17] [|v48|, |v49|, v38.f36, |v50|, |(if (v38.f36 in v51) then v51[v38.f36] else v14)|, v38.f37, v38.f36, |"umscmdh"|, v38.f36, v38.f36, v38.f37, |fm31(p1, globalState)|, -v38.f36, v38.f36, v38.f37, v38.f37, v38.f36];
				var v53: map<array<int>, char> := map[v52 := v39];
				v53 := v53[v52 := v33];
				r0 := (if (v41) then v38.f36 else |v47.fm3(v14, v1, v38.f36, globalState)|) / v38.f36;
			} else {
				var v54: array<int> := new int[7] [v38.f36 / v38.f37, v38.f36, v38.f37, p1, v38.f37, v38.f36, v38.f37 + v38.f36];
				var v55: array<array<int>> := new array<int>[4] [v54, v54, v54, v54];
				globalState.f2, v54, v38.f37, v55 := v35, v54, 0x59, v55;
				var v56 := DC0(f26);
				globalState.f12 := v38.fm4(v56, -(v38.f36 * v38.f37), v38.f37 / v38.f36, globalState);
				globalState.f22 := !true;
				var v57: array<char> := new char[3];
				var v58: array<C1> := new C1[5] [v38, v38, v38, v38, v38];
				v58[832] := v38;
				v57, globalState.f22, globalState.f16, globalState.f16, v58[832] := v57, f27, 826 * v38.f36, v38.f36, v38;
				globalState.f22 := f26;
			}
			
		}
		
		var v59: array<bool> := new bool[7](i7 => f27);
		var v60: map<bool, bool> := map[f26 := f27];
		var v61, v62, v63, v64 := m9(f27, v59, if (f26) then v60 else v60, globalState);
		var v65: seq<seq<bool>> := [v1];
		v65 := ([v1] + v65) + v65;
		r0 := p1;
		r1 := v63;
		r2 := -p1;
		r3 := fm24(v14, f26, p1, globalState);
	}
}

class C5 extends T1 {
	constructor (f26 : bool, f27 : bool) {
		this.f26 := f26;
		this.f27 := f27;
	}
	
	function fm1(globalState: GlobalState): int {
		-0x164
	}
	function fm2(p0: bool, p1: int, globalState: GlobalState): bool {
		match if (f27) then DC2(DC1(f27, f27, 0x2f4, !true)) else DC2(DC0(f26)) {
			case DC1(cf1, cf2, cf3, cf4) => f27
			case DC0(cf0) => |"wrscocei"| == 0x1d4
			case DC2(cf5) => f26
		}
	}
	function fm5(p0: seq<bool>, p1: seq<bool>, p2: D0, p3: bool, globalState: GlobalState): char {
		'g'
	}
	method m1(p0: bool, p1: seq<string>, p2: seq<bool>, globalState: GlobalState) {
		var v0 := new C0(f26);
		var v2: map<seq<bool>, bool> := map[p2 := p0];
		var v3: map<map<seq<bool>, int>, bool> := map[map v1 : seq<bool> | v1 in v2 :: (v1) := (|"fmtrj"|) := p0];
		var v4 := -0x1ee;
		var v5: map<seq<bool>, int> := map[p2 := v4];
		var v7 := DC6(map v6 : int | (0x33c <= v6) && (v6 < 0x3d6) :: (v6 - v4) := (v4));
		v3 := v3[v5 := !(|v7.cf13| > v4)];
		globalState.f6 := "nbjowp";
		var v8 := new C0(p0 ==> f27);
		var v9 := DC9(0x212, |p2|, f26);
		match (v9) {
			case DC7(cf14, cf15, cf16, cf17) =>
				var v10: array<string> := new string[9];
				var v11 := "m";
				v10[397] := v11;
				v10[397] := v11 + v11;
				globalState.f22 := fm2(cf14, -v4, globalState);
				var v12: array<bool> := new bool[20];
				v12[461] := v10[397] <= v10[397];
				v12[461] := false;
				globalState.f2 := f27;
			case DC8() =>
				globalState.f15 := v0.f32;
				var v13 := DC1(v0.f32, v8.f32, v4, v8.f32);
				var v14: map<int, int> := map[v4 := v13.cf3];
				var v15: multiset<map<int, int>> := multiset{fm8(0x28, v4, v0.f32, v4, globalState), v14, v14, v14, v14};
				var v16 := 'o';
				var v17: map<bool, char> := map[p0 := v16];
				var v18: map<int, multiset<map<int, int>>> := map[|(v17 + map[p0 := v16])| := multiset{v14}];
				v15 := if (v4 in v18) then v18[v4] else v15 + v15;
				var v19 := "sksvlyg";
				globalState.f6 := v19[v4 := v16] + v19;
				var v20: multiset<bool> := multiset{v0.f32, v0.f32};
				v20 := fm9(!v8.f32, v4, v0.f32, globalState);
			case DC9(cf18, cf19, cf20) =>
				var v21 := "bfsobu";
				var v22: multiset<bool> := multiset{true};
				var v23: map<bool, int> := map[f26 := cf19];
				var v24: multiset<int> := multiset{|v22|, |v23|};
				var v25 := 'l';
				var v26: map<string, char> := map[v21[if (fm1(globalState) in v24) then v24[fm1(globalState)] else cf19 := v25] := v25];
				v26 := v26["ja" := v25];
				globalState.f16 := v4;
				cf18 := fm1(globalState);
				var v27: map<int, bool> := map[cf18 := v0.f32];
				var v28: map<map<int, bool>, int> := map[v27[cf18 := f26] := v4];
				cf18 := |v28|;
			case DC6(cf13) =>
				globalState.f16 := |(seq(0x217, i0  => ('g')))|;
				var v29 := "cc";
				var v30: map<bool, int> := map[v8.f32 := v4];
				var v31: seq<int> := [v4, v4, v4];
				var v32: map<int, bool> := map[v4 := v8.f32];
				var v33 := DC5(v4);
				var v34: array<int> := new int[21] [v4, 41, v4, v4, v4, v4, v4, 0x377, v4, v4, |v29|, v4, v4, v4, if (v0.f32 in v30) then v30[v0.f32] else v4, v4, v31[v4], |v29|, v4, |v32|, v33.cf12];
				var v35 := DC10(v34);
				var v36: array<array<int>> := new array<int>[25] [v34, v34, v34, v34, v35.cf21, v34, v34, v34, v34, v34, v34, v34, v34, v34, v34, v34, v34, v34, v34, v34, v34, v34, v34, v34, v34];
				v36 := v36;
				var v37: map<string, int> := map[v29 := v4];
				var v38: multiset<bool> := multiset{p0, v0.fm6('j', v8.f32, globalState), v0.f32, p2[|v37[v29 := v4]|], v8.f32};
				var v39: multiset<int> := multiset{|v38| + v4, v4};
				var v40: array<char> := new char[1];
				var v41 := DC4(v40, v4, v4, v29, v29);
				var v42: map<bool, bool> := map[p0 := v0.f32];
				var v43: map<bool, string> := map[v0.f32 := (v41.(cf8 := |v42|)).cf10];
				var v44 := DC1(v0.f32, f26, v4, !f26);
				v39 := fm10(v4, |v43|, DC2(v44), globalState) + v39;
				if (f26) {
					globalState.f16 := |((multiset{f26} + v38) - v38)|;
					globalState.f16 := v4;
					var v45: array<bool> := new bool[3] [p0, v0.f32, v0.f32];
					v45[752] := f26;
					var v46 := DC0(v0.f32);
					v45[752], globalState.f24 := !p0, if (v46.cf0) then v0.f32 else !f27;
					v34 := v34;
					globalState.f12 := -v4 % v4;
				} else {
					var v47: array<bool> := new bool[14];
					v47[358] := v39 < v39;
					var v48 := 'd';
					v47[358] := v8.fm6(v48, v8.f32, globalState);
					var v49 := DC8();
					v49, globalState.f24 := v49, true;
					var v50 := DC0(true);
					var v51 := m8(v50, f27, (multiset{v0.f32})[v0.f32 := v4], v4, globalState);
					v34[616] := v4;
					v34[616] := (if (v47[358]) then v4 else v4) - |v29|;
					v48 := v48;
				}
				
		}
		
		var v52 := 'e';
		var v53 := DC3(fm11(v8.f32, v0.f32, v52, globalState));
		match (v53.(cf6 := p2[v4 := v0.f32])) {
			case DC4(cf7, cf8, cf9, cf10, cf11) =>
				var v54: array<array<int>> := new array<int>[4];
				var v55: map<array<array<int>>, int> := map[v54 := fm1(globalState)];
				var v56: map<bool, set<bool>> := map[false := {true, v0.f32, f27}];
				v55 := v55[v54 := |v56| + cf8];
				var v57: set<bool> := {!v8.f32, f27};
				var v58 := DC12(v57);
				globalState.f2 := v58.cf23 > v57;
				v4 := cf9;
				var v59 := DC8();
				match (v59) {
					case DC7(cf14, cf15, cf16, cf17) =>
						var v60: array<int> := new int[18](i1 => i1 % |p2|);
						var v61: map<array<int>, int> := map[v60 := fm1(globalState)];
						v61 := v61[v60 := cf17 - |cf10|];
						globalState.f6 := (seq(45, i2  => (v52))) + cf11;
						v8 := v8;
						globalState.f6 := seq(0x33e, i3  => (v52));
					case DC8() =>
						var v62: seq<int> := [|(multiset{v52})[v52 := 0xcc]|];
						var v63: array<bool> := new bool[23];
						var v64: map<array<bool>, seq<int>> := map[v63 := v62[v4 := cf9]];
						var v65: map<seq<int>, seq<int>> := map[v62 := if (v63 in v64) then v64[v63] else v62];
						globalState.f15 := !(v65 != v65);
						var v66: map<int, seq<int>> := map[cf9 := v62];
						v66 := v66[v4 := ((seq(0x32e, i4  => (cf9))) + v62)[v4 := v4]];
						var v67: map<int, int> := map[v4 := cf8];
						v67 := v67[-0x3b0 := v4];
						globalState.f2 := v0.f32;
					case DC9(cf18, cf19, cf20) =>
						globalState.f2 := (cf18 < cf8) <== cf20;
						cf7[114] := if (v8.f32) then v52 else v52;
						var v68: map<bool, bool> := map[cf20 := v0.f32];
						var v69: multiset<int> := multiset{v4, -(cf8 - v4), -(if (false) then |v68| else cf18), cf9 * 194};
						cf7[114], v69, globalState.f12 := v52, (v69 * v69) + (v69 * multiset{|map[v4 := p0]|}), cf18;
						var v70: array<int> := new int[1](i5 => i5 / v4);
						var v71 := DC10(v70);
						var v72: map<D3, int> := map[v71 := fm1(globalState)];
						var v73: T1 := new C4(v72, v8.f32, f27);
						var v74: map<int, T1> := map[cf8 := v73];
						var v75: array<int> := new int[15] [cf18, fm1(globalState), |v57|, cf19, 507, |cf10|, cf19, -cf8, cf9, DC1(v0.f32, v8.f32, |cf10|, v8.f32).cf3, cf19, |v74|, cf19, -cf8, cf19];
						var v76: multiset<array<int>> := multiset{v75, v75, v70};
						globalState.f16 := |v76|;
						var v77 := new C1(cf8 + cf9, cf9 - cf9, cf8 >= v4, !v8.f32 || cf20);
					case DC6(cf13) =>
						var v78: C1 := new C1(|p2|, cf9, p0, f27);
						var v79: set<C1> := {v78, v78, v78};
						v79 := v79 - v79;
						globalState.f22 := true && v8.f32;
						globalState.f12 := v78.f36;
						globalState.f6 := cf11 + (if (v0.f32) then cf10 else cf11);
				}
				
			case DC5(cf12) =>
				var v80: map<bool, seq<bool>> := map[v8.f32 := [v0.f32, v0.f32]];
				var v81 := DC13(if (true) then fm32(fm1(globalState), cf12, globalState) else v80[p0 := p2]);
				match (v81) {
					case DC14(cf25, cf26, cf27) =>
						globalState.f16 := -cf26;
						globalState.f16, globalState.f12, globalState.f24 := cf26, 294, v0.f32;
						var v82: array<D2> := new D2[1](i6 => DC8());
						v82[247] := DC8();
						var v83: seq<char> := [v52];
						var v84: set<bool> := {p0, false};
						var v85: map<seq<char>, int> := map[v83 := |v84|];
						v82[247] := fm33(-(if (f26) then 743 else cf25), |(v85 + v85[v83 := cf12])|, globalState);
						globalState.f16 := -cf25;
					case DC13(cf24) =>
						var v86: array<seq<map<int, int>>> := new seq<map<int, int>>[3];
						var v87: map<int, int> := map[-0x2e := cf12];
						var v88: seq<map<int, int>> := [v87];
						v86[522] := v88;
						v86[522] := v88;
						var v89: array<int> := new int[12];
						v89 := v89;
						cf12 := v4;
						var v90 := "priuoxrge";
						var v91: multiset<string> := multiset{p1[v4], v90};
						v91 := v91;
					case DC15(cf28) =>
						var v92: multiset<seq<bool>> := multiset{[v8.f32]};
						var v93: array<bool> := new bool[13](i7 => v0.f32);
						var v94: multiset<array<bool>> := multiset{v93, v93, v93};
						var v95: map<bool, int> := map[f27 := v4];
						var v96 := "dq";
						var v97: seq<int> := [|v96|];
						var v98: set<bool> := {v0.f32, f26, false, 609 == |v92[p2 := if (v93 in v94) then v94[v93] else v4]|, v0.fm7(v95, v97, fm2(f26, -|v96|, globalState), globalState)};
						v98, v9, v97 := v98, v9, v97;
						var v99: multiset<seq<int>> := multiset{v97, v97, v97};
						var v100 := new C3(v99[v97 := v4], v4 <= cf12, v8.f32 <==> !v8.f32);
						v97 := (v97 + [cf12, cf12]) + v97;
						globalState.f22 := v8.f32;
				}
				
				match (DC3([true, f26, v0.f32])) {
					case DC4(cf7, cf8, cf9, cf10, cf11) =>
						var v101: array<int> := new int[11] [cf8, -0x9b, cf9, cf8, 0x110, 0x34a, v4, v4, cf9, |cf11|, v4];
						var v102 := DC10(v101);
						var v103: map<D3, int> := map[v102 := v4];
						var v104: C4 := new C4(v103, v0.f32, v8.f32);
						var v105: seq<C4> := [v104];
						var v106: set<bool> := {v0.f32};
						var v107: set<int> := {|v106|};
						var v108: multiset<seq<int>> := multiset{seq(231, i8  => (v4))};
						var v109: C3 := new C3(v108, v8.f32, !p0);
						var v110: map<bool, C3> := map[{|v105|, cf8} <= v107 := v109];
						v110 := v110;
						globalState.f22 := !(|(cf11 + cf11)| == (if (v8.f32) then -|(seq(0x4d, i9  => (v52)))| else v4));
						var v111: array<seq<bool>> := new seq<bool>[17](i10 => p2 + [v0.f32, f27]);
						v111[267] := p2 + p2;
						v111[267] := p2;
						var v112: map<string, int> := map[seq(0x78, i11  => (v52)) := cf8];
						v112 := v112[cf11 + cf10 := v4];
					case DC5(cf12) =>
						var v113: seq<int> := [v4, cf12];
						var v114 := "nhtyej";
						var v115: multiset<string> := multiset{v114, v114};
						var v116: seq<D2> := [v7, v7, v7];
						var v117: map<seq<int>, seq<D2>> := map[v113[|v115| := cf12] := v116];
						var v118: map<bool, map<seq<int>, seq<D2>>> := map[v8.f32 := map[v113 := v116]];
						var v119 := DC23(map[v113 := v116]);
						var v120: array<map<seq<int>, seq<D2>>> := new map<seq<int>, seq<D2>>[23] [v117, if (true in v118) then v118[true] else v117, v117, v117, v117, v117, v117[fm15(globalState) := [v7]], v117 + v117, (fm34(globalState))[[v4] := v116], v117 + v117, v117, map[v113 := v116], v117, v117[v113 := seq(0x2d1, i12  => (v7))], v117, v119.cf43, v117, v117, map[seq(0x37c, i13  => (cf12)) := v116], v117 + v117, v117[[cf12] := v116], v117, v117 + v117];
						v120[891] := map[v113 := v116];
						var v122: map<seq<int>, int> := map[v113 := v4];
						v120[891] := (v117 + (map v121 : seq<int> | v121 in v122 :: (v121) := (v116))[[cf12] := v116]) + v117;
						globalState.f16, globalState.f16 := |"ertaynatq"| - 893, cf12;
						globalState.f15 := v0.f32;
						var v123: array<C1> := new C1[12];
						var v124: multiset<array<C1>> := multiset{v123, v123};
						v124, globalState.f12 := v124, v9.cf19;
					case DC3(cf6) =>
						var v125 := "wsdmh";
						var v126: seq<int> := [fm1(globalState)];
						var v127: map<bool, int> := map[v0.f32 := v4];
						var v128: map<bool, bool> := map[p0 := false];
						var v129: array<int> := new int[15] [v126[cf12], -cf12, -v4, fm1(globalState), cf12, |v127|, 0x2ba, v4, |v128|, v4, cf12, -v4, cf12, v4, cf12];
						var v130: map<array<int>, string> := map[v129 := v125];
						var v131 := new C1(|v125|, -|v126|, v129 in v130, v8.f32);
						var v132 := DC0(f26);
						globalState.f22 := v0.f32 ==> fm2(f26, v131.fm4(v132, v131.f36, v4, globalState), globalState);
						var v133: map<string, char> := map[v125 + "uxenwl" := v52];
						v133 := v133[v125 := v52];
						v127 := v127[v8.fm6(v52, v0.f32, globalState) := 0x1e7];
				}
				
				if (v8.f32) {
					var v134 := "ck";
					var v135: multiset<int> := multiset{|v134|};
					var v136: map<multiset<int>, bool> := map[v135 := v0.f32];
					var v137: set<bool> := {v8.f32, v0.f32};
					var v138: map<int, set<bool>> := map[v4 := v137];
					var v140: array<bool> := new bool[15](i15 => f26);
					var v141: multiset<array<bool>> := multiset{v140};
					var v142: map<int, int> := map[v4 := if (v140 in v141) then v141[v140] else cf12];
					var v143: map<map<int, int>, int> := map[v142 := 53];
					var v144: multiset<bool> := multiset{v0.f32};
					var v145: array<int> := new int[24] [cf12, |v136|, v4 * v4, |(if (255 in v138) then v138[255] else {v8.f32})|, |(set v139 : int | (158 <= v139) && (v139 < 0x2ea) :: (v139 % |(seq(0x35a, i14  => (-|"cuhflyu"|)))|))|, -v4, cf12, v4, cf12, cf12, v4, v4 - v4, v4, cf12 / cf12, cf12, cf12, cf12 - |v143|, cf12, |v144|, v4, |v142|, v4, fm1(globalState) * v4, -v4 * cf12];
					v145[854] := cf12;
					v145[854] := if (!(f27 <==> f27)) then |v134| else v4;
					globalState.f16 := v4;
					var v146: map<bool, int> := map[v0.f32 := cf12];
					var v147 := DC14(v145[854], cf12, !v0.f32);
					var v148: seq<D5> := [v147, v147];
					var v149: map<bool, bool> := map[v0.f32 := v8.f32];
					var v150: seq<bool> := [(fm35(if (v0.f32 in v146) then v146[v0.f32] else v145[854], globalState))[v4 := v147] == v148, v149 != map[f27 := v0.f32], v8.f32, v8.f32];
					v150 := p2 + ([v0.f32, v8.f32] + v53.cf6);
					var v151: multiset<string> := multiset{v134};
					v145[854] := if (v134 in v151) then v151[v134] else v145[854];
					var v152 := new C1(-v145[854], cf12, v0.f32, f27);
				} else {
					cf12 := -(cf12 / cf12);
					var v153: multiset<int> := multiset{0x36e, v4};
					var v154 := m8(DC0(v0.f32), v8.f32, multiset{v0.f32, v0.f32, p0}, cf12 % |v153[cf12 := v4]|, globalState);
					var v155: array<int> := new int[29](i16 => i16 + 0x79);
					var v156 := DC10(v155);
					var v157: map<D3, int> := map[v156 := --v4];
					var v158 := new C4(if (v0.f32) then v157 else v157, v0.f32, f27);
					globalState.f2 := v8.f32;
					var v159: array<map<int, string>> := new map<int, string>[21];
					var v160 := "hnfdufwb";
					var v161: map<int, string> := map[cf12 := v160];
					var v162: seq<map<int, string>> := [v161, v161, v161, v161];
					v159[780] := v162[cf12] + v161;
					var v163 := DC13(v80[f26 := [v0.f32]]);
					var v164 := DC15(v163);
					var v165 := DC15(DC15(v164));
					var v166 := DC15(v164);
					var v167: array<D5> := new D5[6] [DC15(v163), v166, v166, v166, v166, v166];
					v159[780], v154, v167, globalState.f2 := (map v168 : int | v168 in v7.cf13 :: (v168 + 0x2d5) := (v160)) + (v161 + map[-v4 := v160]), p0, v167, !v0.f32;
				}
				
				if (f26) {
					var v169: map<int, bool> := map[-cf12 + cf12 := v8.f32];
					var v170: map<bool, int> := map[v8.f32 := v4];
					v169 := v169[|v170| := v0.f32];
					var v171: array<D0> := new D0[16];
					var v172 := DC0(f27);
					v171[95] := v172;
					var v173: map<bool, bool> := map[f27 := v8.f32];
					v171[95], globalState.f16 := DC0(if (v8.f32 in v173) then v173[v8.f32] else v0.f32), v4;
					var v174: seq<int> := [cf12, v4];
					var v175 := DC20(v174);
					var v176: seq<D8> := [v175, v175];
					var v177 := "iywdtfdjg";
					var v178 := DC17(cf12, cf12, v4, |v177[-0x4b := v52]|, v4);
					var v179: map<seq<D8>, D6> := map[v176 := v178];
					v179 := v179;
					var v180 := DC1(v0.f32, v8.f32, |v177|, p0);
					var v181: C1 := new C1(v4, v4, f26, f27);
					var v182: multiset<C1> := multiset{v181};
					var v183: set<bool> := {f26, v8.f32};
					var v184: map<bool, char> := map[true := v52];
					var v185 := DC14(v181.f37, |v177|, v8.f32);
					var v186: array<bool> := new bool[20] [!!(-0x3e2 != v4), v8.f32, fm2(true, fm1(globalState), globalState), v8.f32, v8.f32, v0.fm7(v170, v174, f27, globalState) ==> true, v0.f32 ==> v180.cf1, f26, f27, v8.f32, v0.fm7(v170, v174, !p0, globalState) && f26, (if (v181 in v182) then v182[v181] else |v174|) < |v183|, v0.fm6(v52, !v0.fm7(v170, [|v184|], v0.f32, globalState), globalState), true, v0.f32 <== !f27, v185.cf27, if (p0) then v8.f32 else p0, !!p0 <==> v0.f32, v8.f32, p2 == [p0, p0, v8.f32]];
					v186[947] := v8.f32;
					v186[947] := !v0.f32;
					var v187: array<D1> := new D1[15];
					var v188: seq<array<D1>> := [v187];
					globalState.f22 := v188 != v188;
				} else {
					globalState.f12 := cf12;
					var v189: seq<int> := [0x38f, |p2|];
					var v190: seq<seq<int>> := [v189];
					var v191: T2 := new C2(v190, v0.f32, p2[v4]);
					var v192 := DC18(v191);
					var v193: multiset<D7> := multiset{v192};
					var v194: seq<multiset<D7>> := [v193];
					var v195: seq<D7> := [v192, v192];
					globalState.f24 := !(v194[38] >= (multiset(v195))[v192 := cf12]);
					var v196: array<bool> := new bool[5];
					v196, globalState.f2, globalState.f16 := v196, v8.f32 <== f27, fm1(globalState);
					var v197: array<int> := new int[2];
					v197[581] := cf12;
					v197[581] := cf12;
					var v198: set<array<int>> := {v197};
					v198 := {v197, v197, v197, v197} * {v197, v197, v197};
				}
				
			case DC3(cf6) =>
				var v199: map<bool, bool> := map[v8.f32 := p0];
				globalState.f15 := if (f27 in v199) then v199[f27] else p2[0x1c2];
				var v200: map<bool, int> := map[v8.f32 := v4];
				var v201: seq<int> := [v4];
				var v202: map<bool, map<int, int>> := map[v0.fm7(v200, v201, p2[v4], globalState) := map[v4 := -0x2d]];
				var v203: map<int, int> := map[v4 := -0x15d];
				var v204: map<int, int> := map[if (v4 in v203) then v203[v4] else 0x390 := 0x2];
				var v205 := "gwqodhw";
				var v206: map<int, bool> := map[|"rphevkuq"| := v0.fm6(v52, v0.f32, globalState)];
				var v207: multiset<seq<int>> := multiset{v201, v201};
				var v208: C3 := new C3(v207, f27, v0.f32);
				var v209: array<int> := new int[22] [v4, v4, |v199|, v4, v4, |v206|, |v205|, |[v0.f32]|, v4, v4, v4, v4, v4, 138, v4, v4, v4, 0x3d, 0x21c, v4, v4, |[v208]|];
				var v210 := DC10(v209);
				var v211: C4 := new C4(map[v210 := |v205|], f27, v8.f32);
				var v212: map<C4, bool> := map[v211 := f27];
				var v213: map<seq<int>, int> := map[v201 := v4];
				var v214: array<int> := new int[29] [v4, -0x368, v4, |(v202 + map[v0.f32 := v203])|, v4, -0x22d, -248 / v4, v4, v4, -|(if (f26) then v205[|v199| := v52] else v205)|, -v4 - v4, 0x162 - v201[v4], -(v4 / |v206|), v4, |v205[v4 := v52]|, 0x352, 0x3ae, v4, v4, -|(v212 + v212)|, v4, |v213|, if (v8.f32) then 0xc4 else 0x298, v4 % v4, v4, v4, v4, v4, 0x1e0];
				v214 := v209;
				v212 := v212;
				var v215: array<seq<bool>> := new seq<bool>[8](i17 => [f26, f27]);
				v215[769] := cf6;
				v215[769] := p2 + cf6;
		}
		
	}
	method m7(p0: bool, p1: bool, p2: int, p3: multiset<int>, globalState: GlobalState) {
		var i0 := 0;
		while (f27)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v0: C1 := new C1(p2, p2, f26, p1);
			var v1: array<C1> := new C1[3] [v0, v0, v0];
			var v2: seq<array<C1>> := [v1];
			globalState.f2 := v2 <= v2;
			var v3 := "jn";
			var v4: set<string> := {(seq(0x339, i1  => ('c'))) + "yknlqwi", v3};
			v0.f37 := |v4|;
			var v5: seq<string> := [v3];
			v5 := ((v5 + (seq(0x14a, i2  => (v3)))) + v5)[v0.f37 := v3];
			var v6: array<array<map<T0, bool>>> := new array<map<T0, bool>>[5];
			var v7: array<map<T0, bool>> := new map<T0, bool>[26];
			v6[961] := v7;
			v6[961] := v7;
		}
		var v8: array<bool> := new bool[6](i3 => f26);
		v8[477] := f26;
		var v9: array<int> := new int[4];
		v8[234] := f26;
		var v10: seq<int> := [p2];
		var v11: multiset<seq<int>> := multiset{v10, (seq(-0x18c, i4  => (p2)))[p2 := p2]};
		globalState.f15, v8[477], v9, v8[234], globalState.f2 := false, f27, v9, p1, v11 >= (v11 + v11);
		for i5 := -fm1(globalState) / p2 to p2 {
			globalState.f12 := p2;
			var v12: map<bool, int> := map[fm2(v8[477], fm1(globalState), globalState) := p2];
			globalState.f16 := if (i5 <= p2) then |(seq(0x1cf, i6  => ('i')))| + |v12| else 0x377;
			var v13: seq<bool> := [!p1, f26];
			globalState.f15 := if (!!true) then v13[p2] else p0;
			if (p1) {
				var v14: map<bool, bool> := map[p0 := p0];
				var v15: map<int, int> := map[p2 := |v14|];
				var v16 := DC6(v15);
				var v17: seq<D2> := [v16, DC6(v15)];
				var v18: multiset<seq<D2>> := multiset{v17, v17, v17, v17, v17};
				v8[477] := (if (f26) then multiset{v17[0x2bc := v16]} else fm36(p2, v8[477], !fm2(p1, p2, globalState), f26, globalState)) >= v18;
				var v19 := DC10(v9);
				var v20: map<int, array<bool>> := map[p2 := v8];
				var v21: map<D3, int> := map[v19 := |v20|];
				var v22 := new C4(v21, p1, p1);
				var v23 := "gegvu";
				globalState.f6 := v23 + v23;
				var v24: seq<seq<int>> := [seq(0x75, i7  => (p2)), [p2], [-|v15|, i5], v10, [i5, i5]];
				var v25: C2 := new C2(v24, v8[477], f27);
				var v26: map<seq<int>, C2> := map[v10 := v25];
				var v27: array<map<seq<int>, C2>> := new map<seq<int>, C2>[3] [v26 + v26, v26, v26];
				v27[450] := v26;
				v27[450] := v26;
				var v28 := DC0(f26);
				v9[821] := v25.fm4(v28, i5, p2, globalState) - i5;
				globalState.f12, v9[821] := -(if (v13[i5]) then 606 else i5 - p2), i5;
			} else {
				var v29: array<array<int>> := new array<int>[6] [v9, v9, v9, v9, v9, v9];
				v29[699] := v9;
				v29[699] := v9;
				var v30 := "ctfdse";
				var v31: seq<array<int>> := [v29[699], v9, v9];
				v9, globalState.f6, globalState.f16, v9 := v29[699], v30, 0x32c, v31[p2];
				globalState.f16 := fm1(globalState);
				var v32: array<seq<array<D1>>> := new seq<array<D1>>[4];
				var v33: array<D1> := new D1[4];
				var v34: seq<array<D1>> := [v33, v33];
				v32[349] := v34;
				var v35: multiset<multiset<int>> := multiset{p3, p3};
				v32[349], globalState.f16, globalState.f15 := v34 + v34, i5, v35 >= v35;
				var v36 := new C1(p2, i5, --0x7d == |p3|, v8[477]);
			}
			
		}
		v8[477] := f26;
		var v37: map<int, bool> := map[0x3b2 := true];
		globalState.f24 := |v37| == p2;
		if (f26) {
			var v38: map<int, int> := map[|p3| := p2];
			var v39: multiset<int> := multiset{p2 / |v38|, p2};
			v39 := v39 - (v39 - v39);
			globalState.f16 := |v10|;
			var v40: seq<bool> := [f26];
			v10 := [|v40|, p2, 0x257 - p2];
			globalState.f22 := p1;
			var v41: array<array<int>> := new array<int>[21];
			v41 := v41;
		} else {
			v9[478] := p2;
			v9[478] := p2;
			var v42 := DC0(p0);
			var v43: seq<bool> := [true];
			var v44: multiset<bool> := multiset{p1};
			var v45 := m8(v42, v43[fm1(globalState)], v44, v9[478], globalState);
			var v46: multiset<array<bool>> := multiset{v8, v8};
			var v47 := "nhfpskf";
			var v48: map<int, string> := map[|(v46 * v46)| := if (f27) then v47 else v47];
			v48 := v48[v9[478] := "t" + fm22(p2, fm2(p0, v9[478], globalState), globalState)];
			var v49: map<bool, int> := map[p1 := v9[478]];
			var v50 := DC19(v9[478], v44, f27, v49);
			if (v50.cf38) {
				v9[478] := v9[478];
				globalState.f12 := p2;
				var v51: map<int, seq<bool>> := map[v9[478] := [v50.cf38]];
				var v52 := 'h';
				v51 := v51[fm1(globalState) := fm11(v8[477], v8[477], v52, globalState)];
				var v53: array<array<string>> := new array<string>[23];
				var v54: array<string> := new string[25] ["fbnprm", v47, v47, v47, v47, v47, "c", v47, v47, v47, v47, fm22(p2, v45, globalState), "hluwry", v47, v47, v47, v47, v47, v47, v47, v47, ("pkxwgoox")[v9[478] := v52], v47, seq(773, i8  => (v52)), v47];
				var v55: seq<array<string>> := [v54, v54];
				v53[192] := if (fm2(true, v9[478], globalState)) then v54 else v55[p2];
				v53[192] := v54;
				var v56: map<bool, seq<bool>> := map[f26 := [f27, v45]];
				var v57 := DC13(v56);
				v57 := v57;
			} else {
				var v58 := 'a';
				var v59: array<int> := new int[25] [p2, v9[478], p2, v9[478], p2, v9[478], v9[478], -v9[478], |v10|, p2, p2, p2, 0x25e, p2, v9[478], v9[478], p2, |[v58]|, p2, p2, v9[478], p2, |v47|, p2, p2];
				var v60: map<array<int>, int> := map[v59 := |v47|];
				var v61: map<multiset<bool>, set<bool>> := map[v44 := {f27}];
				var v62 := DC14(|(if (v44 in v61) then v61[v44] else {f26, f26})|, p2, v43[v9[478]]);
				globalState.f12, v9[478], globalState.f24 := if (f27) then |(v60 + map[v59 := v9[478]])| else p2, v62.cf26, if (!(p2 <= p2)) then v45 else v8[477];
				var v63: C0 := new C0(f26);
				var v64: set<bool> := {f26};
				var v65: map<seq<D3>, C0> := map[seq(0x106, i9  => (DC11(v58))) := v63];
				var v66 := DC11(v58);
				globalState.f2, v63 := v63.f32 in (if (f27) then DC12(v64).cf23 else v64), if ([v66, DC11(v58), v66, v66] in v65) then v65[[v66, DC11(v58), v66, v66]] else v63;
				globalState.f6 := v47;
				var v67: array<string> := new string[18] ["mxwll" + v47, v47, v47, v47[v9[478] := v58], v47, v47, v47 + v47, v47, "lawmo", v47, fm22(v9[478], p0, globalState), seq(0x2cc, i10  => (v58)), v47, seq(715, i11  => ('u')), v47, v47, v47, v47];
				v67 := new seq<char>[3](i12 => seq(0x2de, i13  => (v58)));
				var v68 := DC26(v8);
				var v69: seq<array<bool>> := [v8, v68.cf47, v8, v8];
				var v70: map<string, seq<array<bool>>> := map["hnkb" := [v8]];
				var v71: array<seq<array<bool>>> := new seq<array<bool>>[5] [v69, v69, v69 + v69, v69 + v69, v69 + (if (fm22(v9[478], v63.fm6(v58, v63.f32, globalState), globalState) in v70) then v70[fm22(v9[478], v63.fm6(v58, v63.f32, globalState), globalState)] else v69)[p2 := v8]];
				v71[573] := v69;
				v71[573] := v69 + (v69 + v69);
			}
			
			if (p1) {
				var v72: map<string, array<bool>> := map[v47 := v8];
				v72 := v72[seq(0xa2, i14  => ('l')) := v8];
				var v73 := new C3(v11, p1 <==> v45, v45 ==> v45);
				var v74 := m8(v42, f26, multiset{v8[477]}, 0x229 / |v47|, globalState);
				globalState.f12 := -v9[478] / p2;
				var v75 := 'h';
				globalState.f19 := v75;
			} else {
				var v76: map<map<int, bool>, bool> := map[v37 := p1];
				var v77: C1 := new C1(|v76|, p2, p0, false);
				var v78: map<C1, multiset<int>> := map[v77 := p3];
				var v79: map<int, multiset<int>> := map[p2 := if (v77 in v78) then v78[v77] else multiset{v77.f36}];
				var v80: C0 := new C0(true);
				var v81: seq<C0> := [v80, v80];
				var v82: map<map<int, multiset<int>>, seq<C0>> := map[v79 := v81 + v81];
				v82 := v82[v79 := v81];
				v10 := v10 + v10;
				var v83: array<int> := new int[29];
				var v84 := DC16("drgpsi");
				var v85 := 'c';
				var v86: map<D6, char> := map[v84 := v85];
				var v87: set<map<D6, char>> := {v86};
				var v89 := DC10(v83);
				var v90: C4 := new C4(map[v89 := -v9[478]], p0, true);
				var v91: seq<C4> := [v90, v90, v90];
				var v92: map<int, int> := map[|v91| := v77.f37];
				var v93 := DC6(map v88 : int | v88 in v92 :: (v88 * -0x320) := (p2));
				var v94: map<seq<int>, seq<D2>> := map[v10 := (seq(0x2cf, i15  => (DC6(map[v77.f37 := |[[v9[478]]]|]))))[|v47| := v93]];
				var v95 := DC23(v94);
				globalState.f16, v83, v8, v87, globalState.f16 := v77.f37, v83, v8, {fm37(v45, v77.fm1(globalState), p2, v95, globalState) + v86}, 569;
				v8[477] := v90.fm1(globalState) <= (|[v8[477]]| - v77.f36);
				var v96: set<int> := {|v47|};
				v8[477] := (|v96| - p2) < p2;
			}
			
		}
		
	}
	method m8(p0: D0, p1: bool, p2: multiset<bool>, p3: int, globalState: GlobalState) returns (r0: bool) {
		var v0: array<string> := new string[29];
		var v1 := 'h';
		v0[371] := ("qsnm")[p3 := v1];
		var v2 := "yxuu";
		v0[371] := v2;
		var v3: seq<bool> := [fm2(f26, p3, globalState), p1];
		globalState.f12 := |([f26] + v3)|;
		globalState.f16 := --p3;
		var v4: map<int, int> := map[|[p3, p3]| := 0x1ab];
		var v5 := DC6(v4);
		var v7: set<int> := {p3};
		var v8: array<D2> := new D2[18] [v5, v5, v5, v5, v5.(cf13 := map v6 : int | v6 in v7 :: (v6 / 0x10e) := (p3)), DC6(map[p3 := p3]), DC6(v4), v5, v5, v5, v5, DC6(v4), v5, v5, v5, DC6(v4), v5, DC6(map[p3 := p3])];
		var v9: map<bool, array<D2>> := map[fm5(v3, v3, p0, !false, globalState) !in (seq(0x3e7, i0  => ('c'))) := v8];
		v9 := v9[f26 := v8];
		for i1 := p3 to -p3 {
			var v10: array<bool> := new bool[28];
			var v11: map<bool, set<int>> := map[p1 := v7];
			v10[653] := !((if (f27 in v11) then v11[f27] else v7) >= {p3, i1});
			var v12: map<int, bool> := map[p3 := p1];
			var v13: multiset<array<bool>> := multiset{v10};
			globalState.f22, v10[653] := (|v12| % i1) != 0x129, v13 > (v13 - multiset{v10, v10});
			globalState.f16 := p3;
			var v14: multiset<seq<bool>> := multiset{fm25(v10[653], v10[653], 'q', globalState) + v3};
			v14 := v14 + multiset{v3, v3, [v10[653], fm2(p1, p3, globalState)], [f26, v10[653], f26]};
			globalState.f16 := i1;
		}
		var v15: seq<int> := [fm1(globalState)];
		var i2 := 0;
		while (if (!v3[|v15[p3 := 0x3c]|]) then f26 else false)
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			var v16 := DC1(f27, f26, |v3|, p1);
			var v17 := DC2(v16);
			var v18 := DC2(v17);
			match (DC2(v17)) {
				case DC1(cf1, cf2, cf3, cf4) =>
					var v19: array<char> := new char[10];
					v19[419] := v1;
					v19[419] := v1;
					var v20: map<int, bool> := map[cf3 := true <== cf2];
					v20 := v20[cf3 := cf3 <= p3];
					globalState.f24 := if (fm2(p1, p3, globalState)) then cf4 else true;
					var v21: array<bool> := new bool[28](i3 => cf2);
					v21[404] := cf1;
					v21[404] := !cf1;
				case DC0(cf0) =>
					var v22: array<int> := new int[6](i4 => i4 % p3);
					var v23: map<bool, array<int>> := map[f26 := v22];
					var v24: array<array<int>> := new array<int>[4] [v22, v22, if (!true in v23) then v23[!true] else v22, v22];
					v24[175] := v22;
					v24[175] := v22;
					var v25: multiset<int> := multiset{p3, p3, p3, p3, p3};
					var v26: array<bool> := new bool[11] [0x2d8 > p3, cf0, cf0, cf0, !(v0[371] != v2), !(v25 >= v25), fm2(false, p3, globalState), f26, f27, p1, true];
					v26[717] := p1;
					var v27 := DC16("oqncb");
					var v28: array<D6> := new D6[28] [DC16(seq(0x31a, i5  => ('b'))), v27, v27, fm38(true, p3, v0[371], globalState), v27, v27, v27, v27, v27, v27, DC16(v2).(cf29 := v2), v27, DC16(v2), v27.(cf29 := v2), v27, v27, v27, v27, DC16(v2), v27, v27, v27, v27, v27, v27, DC16(v0[371]), v27, v27];
					v28[485] := v27;
					v26[717], v28[485] := !f27, v27;
					var v29: map<D1, char> := map[DC5(p3) := v1];
					var v30: seq<map<D1, char>> := [v29];
					var v31: map<seq<map<D1, char>>, seq<bool>> := map[v30 + (seq(68, i6  => (map[DC5(p3) := v1]))) := v3];
					v31 := v31[[map v32 : D1 | v32 in map[DC5(p3) := p3] :: (v32) := (v1)] + v30 := v3];
					globalState.f6 := (v27.(cf29 := v0[371])).cf29;
				case DC2(cf5) =>
					var v33: array<bool> := new bool[6](i7 => f26);
					var v34: map<array<bool>, int> := map[v33 := p3];
					v34 := v34[v33 := p3];
					var v35: map<bool, int> := map[f27 := |v0[371]|];
					var v36: multiset<map<bool, int>> := multiset{v35, v35};
					globalState.f6, globalState.f24, globalState.f15, globalState.f24, globalState.f19 := v0[371], p1, |v36| > p3, f27, v1;
					globalState.f2 := p1;
					globalState.f16 := p3;
			}
			
			var v37, v38, v39, v40 := m0(v1, globalState);
			var v41: array<array<int>> := new array<int>[10];
			var v42: array<int> := new int[5];
			v41[894] := v42;
			var v43: T2 := new C1(p3, p3, true, false);
			var v44: seq<T2> := [v43];
			v41[894], v44, globalState.f12, globalState.f24 := if (f27 || p1) then v42 else v42, v44, -(p3 + p3), v43.f26;
			var v45 := DC10(v41[894]);
			var v46: map<bool, bool> := map[v43.f26 := f26];
			var v47: map<D3, int> := map[v45 := |v46|];
			var v48: T1 := new C4(v47, v43.f26, f26);
			var v49: map<T1, bool> := map[v48 := !false];
			v49 := v49[v48 := true];
		}
		r0 := v3[p3] ==> f27;
	}
}

class C6 extends T0 {
	var f40 : int
	constructor (f40 : int) {
		this.f40 := f40;
	}
	
	function fm1(globalState: GlobalState): int {
		(f40 - |(seq(0x71, i0  => (f40)))|) - f40
	}
	function fm2(p0: bool, p1: int, globalState: GlobalState): bool {
		((set v0 : int | v0 in multiset([f40, |map[false := !false]|, -f40]) :: (v0 % |multiset([!true, false, false, !false, true])|)) - {-f40, |multiset{|multiset{f40}|}|, |multiset{true, !true}|, f40}) == ({|(seq(0x190, i0  => ('x')))|, -f40, f40} - {f40, |[set v1 : int | (0xb2 <= v1) && (v1 < 0x3c7) :: (v1 * f40)]|})
	}
	function fm44(globalState: GlobalState): bool {
		!!false && !(!false <==> true)
	}
	method m15(p0: int, p1: D14, p2: bool, p3: bool, globalState: GlobalState) returns (r0: D1, r1: map<multiset<bool>, bool>, r2: int) {
		globalState.f12 := fm1(globalState);
		var v0: multiset<bool> := multiset{p2};
		var v1: map<int, bool> := map[|v0| + 0x362 := p2];
		var v2 := DC37(v1);
		v1 := v1 + v2.cf68;
		globalState.f15 := p3 && p2;
		var v3 := "ymvn";
		var v4: map<string, bool> := map[v3 := p2];
		globalState.f12 := -(if ((map[v3 := p2] != v4) in v0) then v0[map[v3 := p2] != v4] else p0);
		var i0 := 0;
		while (p2)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v5 := DC35(p2);
			v5 := v5;
			var v6: array<int> := new int[25](i1 => i1 / 0x29d);
			v6[182] := f40;
			v6[182] := 0x86;
			var v7: T2 := new C1(p0, f40, p2, p2);
			var v8 := DC18(v7);
			var v9: array<D7> := new D7[18] [v8, v8.(cf35 := v7), v8, v8, v8, v8, v8, DC18(v7).(cf35 := v7), v8, v8, v8, v8, v8, v8, v8, DC18(v7), v8, v8];
			v9[743] := v8;
			v9[743] := v8;
			var v10: array<bool> := new bool[7];
			v10[263] := v7.f26;
			v10[263] := p2;
		}
		var i2 := 0;
		while (p2)
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			var v11: seq<bool> := [p3];
			var v12: set<int> := {f40, f40, p0, p0, p0};
			var v13: array<bool> := new bool[8] [p3, v11[|v12|], true, fm2(p2, p0, globalState), true, p3, p2, p2];
			v13[147] := fm44(globalState);
			v13[147] := f40 > (fm1(globalState) + -p0);
			r2 := p0;
			globalState.f16, globalState.f12, globalState.f12, globalState.f6 := fm1(globalState), fm1(globalState), |[-0xc5, -fm1(globalState), |v12|, -f40 + f40, 0x1c4]|, v3;
			var v14: seq<int> := [|v3|];
			var v15: multiset<seq<int>> := multiset{[p0, p0], v14, v14};
			var v16 := new C3(v15, if (v13[147]) then p3 else true, v13[147]);
		}
		r0 := DC5(p0 % p0);
		var v17: seq<bool> := [p3];
		var v18: map<multiset<bool>, bool> := map[v0 := v17[f40]];
		r1 := v18;
		r2 := f40 * f40;
	}
	method m16(p0: int, globalState: GlobalState) returns (r0: seq<bool>) {
		var v0: seq<int> := [f40];
		var v1: seq<seq<int>> := [v0];
		var v2 := false;
		globalState.f16, v1 := p0, [fm45(v2, p0, -164, v2, globalState), v0, v0] + v1;
		globalState.f16, globalState.f12 := p0, f40;
		var v3: array<int> := new int[24];
		v3[946] := p0;
		var v4: map<int, map<int, bool>> := map[f40 := fm47(v2, v2, v2, !v2, globalState)];
		var v5 := "mswelsh";
		var v6: map<int, bool> := map[|v5| := v2];
		var v7: multiset<int> := multiset{-p0};
		var v8: map<multiset<int>, int> := map[v7 := f40];
		var v9: map<map<bool, int>, int> := map[map[v2 := f40] := |v8|];
		var v10: array<map<int, map<int, bool>>> := new map<int, map<int, bool>>[4] [fm46(v2, p0, p0, v2, globalState), v4, v4 + map[p0 := v6], fm46(v2, if (map[v2 := |"yxbdjm"|] in v9) then v9[map[v2 := |"yxbdjm"|]] else f40, p0, v2, globalState)];
		v10[385] := v4;
		var v11: array<bool> := new bool[27];
		var v12: seq<array<bool>> := [v11, v11];
		var v13: seq<seq<array<bool>>> := [v12];
		var v14: array<seq<array<bool>>> := new seq<array<bool>>[25] [v12 + v12, v12, v12, v12, v12, v12, v12 + v13[490], v12, [v11, v11], ([v11])[|[|"ajrpp"|]| := v11] + v12, v12, v12, v12, v12, v12, [v11, v11, v11], v12, v12[p0 := v11], v12 + v12[f40 := v11], v12, v12, v12, v12 + v12, v12, v12];
		v14[119] := v12;
		var v15 := 'k';
		var v16: map<multiset<int>, char> := map[multiset{p0} := v15];
		var v17: map<char, bool> := map[if (v7 in v16) then v16[v7] else v5[-p0] := v2];
		v3[946], globalState.f24, v10[385], v14[119], v17 := p0, if (v2) then v2 else if (!false) then v2 else v2, v4, v12[p0 := v11], v17;
		globalState.f24 := !!v2;
		var v18 := DC16(seq(0x3d1, i0  => (v15)));
		if (match if (v2) then v18 else v18 {
			case DC17(cf30, cf31, cf32, cf33, cf34) => v2
			case DC16(cf29) => v2
		}) {
			var v19: seq<bool> := [v2];
			v3[946] := |(v19 + (v19[f40 := v2] + v19))|;
			globalState.f24 := v2;
			v3 := v3;
			v3 := v3;
			var v20 := DC5(-f40);
			var v21: set<D1> := {v20};
			var v22: map<bool, int> := map[false := |["yeue", v5]|];
			v21, v22, v11 := {v20, v20, v20} - v21, v22, v11;
		} else {
			var v23: seq<string> := [v5];
			var v24: array<char> := new char[2] [v15, v15];
			var v25 := DC4(v24, p0, p0, v5, "cjntvhbdw");
			var v26: array<string> := new string[20] [v5, v5, v5, v5, v23[f40], "p", v5, v5, v5, v5, "agk", v25.cf10, v5 + v5, seq(566, i1  => (v15)), v5, "lwu", v5 + v5, "gdid", v5 + v5, v5];
			v26[141] := v5;
			var v27: set<bool> := {v2};
			var v28: multiset<bool> := multiset{true};
			v26[141] := v5 + (if (true) then v5 else v5[|v27| := v15])[if (false in v28) then v28[false] else |map[|v7| := v2]| := v15];
			var v29: map<bool, bool> := map[v2 := v2];
			globalState.f22 := !(v2 || (f40 <= |v29|));
			var v30 := DC14(p0, f40, v2);
			var v31: array<seq<int>> := new seq<int>[17];
			var v32: map<array<seq<int>>, array<int>> := map[v31 := if (v2) then v3 else v3];
			var v33: map<bool, int> := map[v2 := -f40];
			var v34 := DC19(-v3[946], v28, v2, map[v2 := v3[946]]);
			var v35 := DC19(f40, v28, v2, v34.cf39);
			v3, globalState.f24, v30 := if (v31 in v32) then v32[v31] else v3, (if (v2 in v33) then v33[v2] else |v33|) <= v35.cf36, v30.(cf25 := 176, cf26 := fm1(globalState));
			v12 := v12 + v14[119];
			var v36: array<array<int>> := new array<int>[2] [v3, v3];
			v36[542] := v3;
			v36[542] := if (true) then v3 else v3;
		}
		
		var v37: set<int> := {v3[946], |v5|, p0, v3[946]};
		var v38: map<bool, set<int>> := map[v2 := v37];
		v38 := v38[v5 <= "krmbocy" := v37];
		var v39: seq<bool> := [v2, v2, v2, v2, v2];
		r0 := v39 + [v2, v2, false, v2];
	}
}

class C7 extends T0 {
	var f38 : bool
	var f39 : map<int, map<bool, int>>
	constructor (f38 : bool, f39 : map<int, map<bool, int>>) {
		this.f38 := f38;
		this.f39 := f39;
	}
	
	function fm1(globalState: GlobalState): int {
		|(match DC20([|map[|map[true := -0x2ea]| := f38]|]) {
			case DC20(cf40) => map[|[f38, false, f38, f38, true]| := 0x1ac] + map[818 := |map[false := DC9(795, -873, f38).cf20]|]
		})|
	}
	function fm2(p0: bool, p1: int, globalState: GlobalState): bool {
		({multiset([699, -0x22, -0x2b, |"vqgjwd"|, |map[!f38 := true]|]), multiset{|"jhuvf"|}} - {multiset{|"ihdf"|, |[|(seq(-0x237, i0  => ("ro")))|, 0x367]|}}) > (set v0 : multiset<int> | v0 in [multiset{--0xc9, |"imnioove"|, 0x70}, multiset([373]), multiset([|multiset{f38, f38}|]), multiset{987}] :: (v0))
	}
	method m13(globalState: GlobalState) returns (r0: bool, r1: int, r2: bool) {
		var v0 := 0x268;
		var v1 := DC28(v0);
		var v2: map<bool, int> := map[fm2(!f38, v0, globalState) := v0];
		for i0 := |[v1, v1, v1, v1.(cf49 := v0).(cf49 := v0), v1]| to if (f38 in v2) then v2[f38] else v0 {
			var v3: map<int, int> := map[v0 := i0];
			var v4 := DC1(f38, f38, |v3|, f38);
			v4 := v4;
			var v5 := DC15(DC14(i0, fm1(globalState), fm2(f38, v0, globalState)));
			var v6 := DC15(v5);
			var v7: array<D5> := new D5[8] [v6, DC15(v5), v6, v6, v6, v6, v6, v6];
			v7[53] := if (!!f38) then v6 else v6;
			v7[53] := DC15(DC14(-v0, i0, f38)).(cf28 := v5).(cf28 := v5);
			var v8: seq<int> := [v0];
			var v9: seq<seq<int>> := [v8[i0 := v0]];
			var v10: C2 := new C2(v9[v0 := [v0, 250, v0, i0]], !f38, f38);
			var v11: multiset<int> := multiset{-248};
			var v12: seq<bool> := [f38, f38];
			var v13: array<int> := new int[19] [i0, |v11|, i0, 0x32c, 0x295, i0, i0, i0, |v12|, i0, i0, v0, v0, v8[i0], v0, i0, i0, i0, i0];
			var v14: map<array<int>, bool> := map[v13 := false];
			var v15 := DC14(i0, -|v12|, f38);
			var v16 := DC7(f38, f38, |v14|, v15.cf25);
			var v17: array<bool> := new bool[7];
			v17[687] := f38;
			var v18 := "mhjfmmfl";
			var v19 := DC11(v18[i0]);
			var v20: map<bool, D3> := map[false := v19];
			var v21: map<seq<int>, seq<D2>> := map[v8 := seq(726, i1  => (DC6(v3)))];
			var v22 := DC23(v21);
			var v23: map<D10, bool> := map[v22 := f38];
			v10, v16, globalState.f16, globalState.f12, v17[687] := if (f38) then v10 else v10, v16, |v20|, i0, if (v22 in v23) then v23[v22] else f38;
			v17[687] := v17[687];
		}
		var v24 := 'l';
		var v25, v26, v27, v28 := m0(v24, globalState);
		var v29: map<char, int> := map[v25 := 0x171];
		globalState.f24, r1 := !((v0 / fm1(globalState)) >= v0), (0x329 % 846) * (if (fm40(v0, fm2(v27, v0, globalState), v27, globalState) in v29) then v29[fm40(v0, fm2(v27, v0, globalState), v27, globalState)] else v0);
		var v30: multiset<int> := multiset{v0, v0};
		var v31: map<multiset<int>, bool> := map[v30 := f38];
		var v32: map<int, int> := map[v0 := |v31|];
		v32 := v32[-v0 := v0];
		f38 := v27;
		var v33 := DC25();
		v30 := match v33 {
			case DC24(cf44, cf45, cf46) => v30
			case DC25() => v30[|"iowvgu"| := v0]
			case DC23(cf43) => v30
		};
		r0 := v27;
		r1 := v0 * v0;
		r2 := v27;
	}
	method m14(p0: bool, p1: bool, p2: char, globalState: GlobalState) returns (r0: bool, r1: multiset<map<bool, int>>, r2: map<seq<bool>, int>, r3: bool) {
		var v0 := -0x1f1;
		var v1: seq<int> := [v0];
		var v2 := new C2([v1], !p0, f38);
		var v3: seq<bool> := [v2.fm2(p1, v0, globalState), p0];
		var v4 := DC3([p1] + v3);
		match (v4) {
			case DC4(cf7, cf8, cf9, cf10, cf11) =>
				var v5: map<bool, bool> := map[p1 := fm2(p1, cf8, globalState)];
				var v6: set<bool> := {f38, if (f38 in v5) then v5[f38] else p1, !f38, f38, p0};
				var v7: map<D4, bool> := map[if (p1) then DC12(v6) else DC12(v6) := p1];
				var v8 := DC12(v6);
				globalState.f22, v3, v3, cf9, globalState.f22 := (v6 * {true}) >= (v6 - v6), v3, v3, -(|v6| % cf8), if (v8 in v7) then v7[v8] else f38;
				var v9: seq<set<bool>> := [v6];
				var v10 := DC31([v6]);
				var v11: map<bool, seq<set<bool>>> := map[p1 := [v6]];
				var v12: array<seq<set<bool>>> := new seq<set<bool>>[18] [v9, v9 + v9[0x2ff := {p0, true}], v9, v9[cf8 := v6] + v9, v9, v9 + v9, [v6, v6, v6] + v9, [v6], v9 + (seq(0x31f, i0  => (v6))), v9, v9, v9[v0 := v6], v10.cf55 + v9, if (p0 in v11) then v11[p0] else v9, v10.cf55, seq(785, i1  => (v6)), if (f38) then v9 else seq(104, i2  => (v6)), v9 + v9];
				v12[645] := v9;
				var v13 := DC11('a');
				var v14: seq<seq<set<bool>>> := [fm41(!f38, v13, cf9, p0, globalState)];
				var v15: C5 := new C5(v3[cf9], true);
				var v16: map<C5, bool> := map[v15 := p1];
				var v17 := DC33(v16);
				globalState.f16, v12[645] := cf8, v9 + v14[|cf10[v2.fm4(DC0(p1), v0, |(v17.(cf61 := v16)).cf61|, globalState) := p2]|];
				var v18: map<int, int> := map[v0 := cf9];
				if (-(if (p0) then cf8 else v0) == |v18[v0 := cf8]|) {
					v18 := v18[0x237 := cf8];
					var v19: array<int> := new int[21](i3 => i3 / -cf9);
					v19[658] := v15.fm1(globalState);
					v19[658] := v1[v15.fm1(globalState)];
					var v20: C0 := new C0(cf8 <= v19[658]);
					v20, v19, v1 := v20, v19, v2.f35[v0];
					var v21 := DC34(cf8, p0, |cf11|);
					v21 := DC34(v19[658], false, v0 / cf9);
					var v22 := new C5(true, true || true);
				} else {
					var v23: array<bool> := new bool[1];
					v23 := v23;
					var v24: array<int> := new int[9] [cf9 / cf8, if (p1) then cf9 else cf9, cf9, -(|v3| % cf9), 0x344, cf8, v0, -cf8 * |cf10|, cf9];
					v24[563] := |v3[v0 := f38]|;
					v24[563] := v0;
					var v25: map<array<bool>, int> := map[v23 := cf9];
					var v26: map<bool, int> := map[p0 := |v25|];
					globalState.f12 := |(v26 + (if (p1) then v26 else map[p0 := v24[563]]))|;
					var v27: multiset<bool> := multiset{p1, !p1, f38};
					var v28: set<int> := {v24[563]};
					var v29 := DC9(v0 - -v0, -v0, {|v27|} >= v28);
					v29 := v29;
					globalState.f16 := cf9;
				}
				
				globalState.f22 := p0;
			case DC5(cf12) =>
				r3 := p1;
				var v30: set<int> := {cf12, v0};
				var v31: array<char> := new char[2] [p2, p2];
				var v32 := "ryndpso";
				var v33 := DC4(v31, cf12, v0, v32, v32);
				var v34: map<bool, bool> := map[p0 := p0];
				var v35: map<int, int> := map[cf12 := -|"vi"|];
				var v36: set<bool> := {f38, p0, true, p1};
				var v37: multiset<int> := multiset{v0, cf12, cf12, cf12};
				var v38: array<int> := new int[24] [-v0, cf12, v33.cf8, cf12, -v0, v0, |(v34 + v34)|, cf12, cf12, cf12, v2.fm1(globalState), cf12 % cf12, v0, |fm42(v0, globalState)|, cf12, v0, |"j"|, cf12, |v35| % |v36|, |multiset{cf12, v0}|, v0, v0 * -|v37|, v0, v0];
				v38[662] := |(v1 + v1)|;
				v30, v1, globalState.f15, globalState.f16, v38[662] := v30 - v30, fm43(v0, 0x38d, |(v3 + v3)|, map[p1 := p2], globalState), p0 && !p0, v0, cf12 + cf12;
				var v39: array<multiset<D11>> := new multiset<D11>[8];
				v39 := v39;
				var v40 := DC0(f38);
				cf12 := v2.fm4(v40.(cf0 := p0), v0 % cf12, cf12 - cf12, globalState);
			case DC3(cf6) =>
				var v41: T0 := new C6(v0);
				var v42, v43 := v2.m2(v0 != v0, v41, f38, globalState);
				if (p1) {
					globalState.f15 := p1;
					globalState.f22 := v41.fm2(p1, v0, globalState);
					var v44: map<bool, bool> := map[p0 := p1];
					var v45 := DC29(v44);
					var v46: set<D13> := {v45, v45};
					globalState.f15 := v46 !! v46;
					var v47: multiset<bool> := multiset{p0};
					var v48: seq<map<int, bool>> := [map[|v44| := false]];
					globalState.f24 := (-|v47| == |v48|) <==> (|cf6| == 0x1f9);
					v0 := v43;
				} else {
					var v49: set<int> := {v43};
					var v50: map<int, bool> := map[v0 := p1];
					var v51 := "bam";
					var v52 := DC13(map[p0 := v3]);
					var v53 := DC15(v52);
					var v54: map<D5, int> := map[v53 := v43];
					var v55: array<int> := new int[25] [v43, |v49|, v0, v0, -0x179, v0, v43, v0, |map[multiset{if (v43 in v50) then v50[v43] else f38} := v43]|, v41.fm1(globalState), -v0, v43, v0, v0, -v43, v43, |v51|, v41.fm1(globalState), |v54|, v43, -v0, |v51|, v0, v43, v43];
					var v56: multiset<bool> := multiset{f38, false, false};
					var v57: map<array<int>, multiset<bool>> := map[v55 := v56];
					v57 := v57 + v57;
					globalState.f22 := p0;
					globalState.f6 := "qx";
					v55[868] := v0;
					v55[868] := DC5(-v43).cf12;
					globalState.f15 := v51 <= (seq(0x291, i4  => ('u')));
				}
				
				var v58: T2 := new C1(v0, v0, p1, p0);
				var v59: multiset<T2> := multiset{v58, v58, v58, v58};
				var v60 := DC34(-0xea, p1, v0);
				var v61: map<bool, int> := map[p0 := |map[944 := v60.cf63]|];
				var v62: set<bool> := {true, p1, !v58.f27, v58.f26, v58.f26};
				var v63: multiset<char> := multiset{p2};
				var v64: array<int> := new int[8] [if (p0) then |v1| else v43, v43, |v59|, v43, |cf6|, 405, (if (v58.f26 in v61) then v61[v58.f26] else |v62|) % v0, |(fm48(|[v0]|, globalState) + v63)|];
				v64[347] := v43;
				var v65 := DC0(cf6[|(seq(0x16a, i5  => (p2)))|]);
				v64[347], v43 := -v2.fm4(v65, v43, v43, globalState), |cf6|;
				var v66: multiset<int> := multiset{-0x292};
				var v67: multiset<multiset<int>> := multiset{v66 - v66};
				v67 := fm49(v66[v43 := -0x3b4], |cf6[v0 := p1]|, v58.f27, globalState);
		}
		
		for i6 := |map[0xce := v0]| to v0 {
			var v68: multiset<int> := multiset{v0};
			var v69: set<bool> := {p1, f38};
			var v70: map<int, bool> := map[-v0 := p1];
			var v71: multiset<map<int, bool>> := multiset{map[|v68| := p1], map[|v69| := p0], v70, v70};
			var v72: map<bool, multiset<map<int, bool>>> := map[!f38 := v71];
			r3 := !(multiset{map[v0 := p1]} > (v71 * (if (p0 in v72) then v72[p0] else v71)));
			var v73: multiset<bool> := multiset{true};
			var v74: map<bool, bool> := map[f38 := p0];
			var v75: array<bool> := new bool[21] [p0, p0, p1, true, p0, p0, v3[if (f38 in v73) then v73[f38] else v0], f38, f38, p0, !true, p1, f38, f38, v2.fm2(if (f38 in v74) then v74[f38] else f38, |(seq(-0x1ff, i7  => (p2)))|, globalState), p1, p1, p1, p0, true, !DC7(p1, true, v0, v0).cf14];
			var v76 := DC32(v73, p1, p0, p1, |v69|);
			var v77: map<D14, bool> := map[v76 := f38];
			v75[329] := v2.fm17(v0, i6, |v77|, fm50(globalState), globalState);
			v75[329], globalState.f16 := p1, i6;
			var v78: T0 := new C6(|v68|);
			var v79, v80 := v2.m2(p0, v78, !(if (p1) then f38 else p0), globalState);
			var v81: map<bool, int> := map[p0 := v0];
			var v82 := DC19(v0, multiset{p1, p0}, v75[329], v81);
			var v83: map<map<bool, bool>, int> := map[map[p1 := f38] := v82.cf36];
			v75[329] := if ((0x281 % -|v83|) in v70) then v70[0x281 % -|v83|] else p0;
		}
		globalState.f16 := v0;
		var v84: array<bool> := new bool[6];
		v84 := v84;
		var v85 := "hcjrgaf";
		if ("tixwcfv" <= v85) {
			var v86: array<int> := new int[27];
			var v87 := DC10(v86);
			var v88: map<D3, int> := map[v87 := |[v0]|];
			var v89 := new C4(v88, p1, f38);
			var v90: seq<string> := [v85, v85, v85, v85];
			v2.m1(if (!!false) then true else f38, v90, v3 + v3, globalState);
			var v91: map<char, string> := map['w' := seq(0x3cb, i8  => ('q'))];
			v91, f38, globalState.f12, globalState.f6 := v91, true, v0, v85 + v85;
			r0 := false;
			var v92: array<string> := new string[9] ["pnu", v85 + "uxsmk", v85 + v85, v85, "m", "fws", v85, "cobamhx", v85];
			v92 := v92;
		} else {
			var v93: map<int, int> := map[v0 := 0xdf];
			var v94: map<array<bool>, int> := map[v84 := |multiset{false, f38, f38, p1, p1}|];
			var v95: set<bool> := {f38, !true, p0, true, false};
			var v96: array<int> := new int[18] [v0, |v93| - 647, --|v85|, v0, v0, v0 / v0, 834, v0, v0 + v0, if (v84 in v94) then v94[v84] else v0, v0, DC38(v0, v0, f38).cf70 - -v0, 0xae, v0, |v95|, v0, v0, if (!f38) then v0 else v0];
			v96[907] := -v0 / v0;
			var v97 := DC0(p0);
			var v98 := DC2(v97);
			var v99: seq<D0> := [v98];
			v0, globalState.f12, v96[907], v99, globalState.f12 := v0, (-v0 / -14) % v0, -|(v85 + v85)| * v2.fm1(globalState), seq(63, i9  => (v98)), v0;
			globalState.f22 := p1;
			var v100: array<array<int>> := new array<int>[15];
			v100[144] := v96;
			v100[144] := v96;
			v95 := v95 + {p1};
			var v101: map<bool, int> := map[false := v96[907]];
			if (if (true) then v101 == v101 else f38) {
				globalState.f24 := (f38 && false) || p1;
				globalState.f15 := f38;
				var v102: map<bool, bool> := map[p1 := p0];
				var v103: map<array<int>, bool> := map[v100[144] := p0];
				v102 := v102[if (v96 in v103) then v103[v96] else p0 := f38];
				v84 := v84;
				v96[907] := if (true) then v96[907] else v0 - v0;
			} else {
				var v104: set<int> := {v96[907] * v96[907], v0, 177, 0x3da, v96[907]};
				var v105: seq<seq<char>> := [v85, v85, v85];
				var v106: map<multiset<seq<char>>, set<int>> := map[multiset(v105 + v105) := {v96[907], v0}];
				var v107: multiset<seq<char>> := multiset{v85};
				v104 := if (v107[v85 := v0] in v106) then v106[v107[v85 := v0]] else v104;
				var v108 := DC40(v2.f35);
				var v109: map<char, seq<int>> := map[p2 := [v96[907], v0]];
				var v110: multiset<seq<int>> := multiset{v1};
				var v111: seq<multiset<seq<int>>> := [multiset(v108.cf73[v96[907] := [|(if (p2 in v109) then v109[p2] else v1)|]]), v110, multiset(seq(777, i10  => (v2.f35[v0]))), v110, v110];
				var v112 := new C3(v111[v0], f38, p1);
				var v113: seq<map<int, int>> := [v93, v93];
				v113 := v113;
				globalState.f16 := v0 / v96[907];
				var v114 := new C2(v2.f35, false, false);
			}
			
		}
		
		r0 := p1;
		var v115: map<bool, int> := map[f38 := v0];
		var v116: multiset<map<bool, int>> := multiset{v115, v115};
		var v117 := DC5(226);
		r1 := v116[v115 := v117.cf12] * (multiset{v115} - multiset{v115, v115});
		var v118: map<bool, char> := map[f38 := p2];
		var v119: map<seq<bool>, int> := map[v3 + v3 := |(if (p1) then map[p1 := p2] else v118)|];
		r2 := v119;
		r3 := p1;
	}
}

class C8 extends T3 {
	constructor (f29 : multiset<int>) {
		this.f29 := f29;
	}
	
	method m3(p0: int, p1: int, p2: string, p3: int, globalState: GlobalState) returns (r0: bool, r1: bool, r2: array<int>, r3: T0) {
		globalState.f16 := p0;
		for i0 := p0 to p0 {
			var v0 := 'y';
			if ((p3 / -589) == -|p2[p1 := v0]|) {
				var v1 := false;
				var v3: map<int, int> := map[p0 := i0];
				var v4: set<int> := {-207, |map[p0 := |(map v2 : int | v2 in v3 :: (v2 / i0) := (i0))|]|, p3};
				var v5: map<bool, set<int>> := map[v1 := v4];
				var v6: map<int, map<bool, set<int>>> := map[i0 * p0 := v5 + v5];
				v6 := v6[|p2| := v5 + v5];
				v3 := v3[fm39(globalState) := p1];
				var v7: array<int> := new int[11](i1 => i1 + p1);
				v7[905] := -p3;
				var v8: array<bool> := new bool[7](i2 => v1);
				v8[931] := !(v1 && v1);
				var v9: map<bool, bool> := map[false := true];
				var v10 := DC29(v9);
				var v11: seq<int> := [i0, |p2|, i0, p3, p0];
				v7[905], v8[931], globalState.f16, globalState.f15, globalState.f12 := 109 * (|v10.cf50| * i0), v1, p0 * p0, if (!v1) then p3 in v11 else v1 && v1, (i0 + p1) + (if (v1) then 0x242 else i0);
				globalState.f22 := v8[931];
				var v12 := new C0(!v8[931]);
			} else {
				var v13 := true;
				globalState.f22 := !(v13 && v13);
				globalState.f16 := p1;
				var v14: set<char> := {v0};
				var v15 := new C1(p3, i0, v13, |p2| == |v14|);
				var v16: array<bool> := new bool[21];
				v16[829] := true;
				v16[829] := !v13;
				globalState.f16 := v15.f37;
			}
			
			globalState.f12 := i0;
			var v17: seq<int> := [i0, fm39(globalState)];
			globalState.f16 := |v17|;
			globalState.f19 := v0;
		}
		var v18: array<int> := new int[27](i3 => i3 * p1);
		var v19: seq<array<int>> := [v18, v18, v18];
		var v20 := DC10(v19[p0]);
		var v21: map<D3, int> := map[v20.(cf21 := v18) := p1];
		var v22: set<array<int>> := {v18};
		var v23 := true;
		var v24 := new C4(v21, v22 !! {v18, v18, v18, v18}, v23);
		var v25 := 'f';
		var v26, v27, v28, v29 := m0(v25, globalState);
		var v30: map<bool, int> := map[v23 := p1];
		for i4 := |p2| - p1 to |v30| {
			globalState.f16 := i4;
			globalState.f16 := v24.fm1(globalState);
			v27 := ([v23] + v27) + v27;
			var v31: seq<int> := [p3];
			var v32: multiset<seq<int>> := multiset{v31};
			var v33: T1 := new C3(v32, v28, v23);
			var v34: map<T1, multiset<int>> := map[v33 := f29];
			var v35: map<map<T1, multiset<int>>, bool> := map[v34 := p1 < 0x4c];
			v35 := v35[v34[v33 := f29] := v27[v33.fm1(globalState)]];
		}
		var v36: set<bool> := {v28};
		r0, globalState.f16 := (v36 * v36) < v36, p0;
		r0 := v23 ==> v23;
		r1 := !(v23 || v23);
		r2 := v19[p0];
		var v38: T0 := new C7(v23, map v37 : int | (-0x149 <= v37) && (v37 < 0x131) :: (v37 + |multiset(v27)|) := (v30));
		r3 := v38;
	}
	method m4(globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0 := 0x122;
		var v1: seq<int> := [v0];
		var v2: C6 := new C6(v0);
		var v3: map<C6, seq<int>> := map[v2 := v1];
		var v4: multiset<seq<int>> := multiset{v1, if (v2 in v3) then v3[v2] else v1};
		var v5 := false;
		var v6: C3 := new C3(v4, false, v5);
		var v7: multiset<C3> := multiset{v6, v6};
		var v8: seq<bool> := [v5, false];
		var v9 := DC17(if (v6 in v7) then v7[v6] else v2.f40, v2.f40, 51, |multiset(v8[v0 := v5])|, v0);
		var v10 := 'v';
		var v11: multiset<char> := multiset{v10};
		var v12: seq<multiset<char>> := [v11];
		var v13: array<D6> := new D6[23] [v9, v9, v9, v9, v9, v9, DC17(v0, v0, -v2.f40, v2.f40, v2.f40), v9, v9, DC17(0xc3, v0, v2.f40, v0, v0), v9, v9, v9.(cf31 := 0x57, cf30 := v6.fm13(v12, v10, |f29|, globalState)), v9, v9, v9, DC17(v0, |fm51(v0, v0, v5, 597, globalState)|, v2.f40, v0, v2.f40).(cf33 := -v2.f40, cf34 := v0), v9, DC17(v2.f40, v6.fm13(v12, v10, v2.f40, globalState), v0, |f29|, v0), v9, v9.(cf30 := v0, cf33 := v2.f40), v9, v9];
		v13[913] := DC17(v0, v0, v0, 0x42, v0);
		v13[913] := DC17(|[v10, v10]| + v2.f40, v2.f40, v2.f40, v0, -v2.f40);
		globalState.f12 := v0;
		var v14: map<bool, int> := map[!v5 := v0];
		var v15: map<bool, map<bool, int>> := map[!v5 := v14];
		v15 := v15[v5 := map[v5 := v0] + map[v5 := v0]];
		if (!!(v8 <= [v5])) {
			var v16 := "kyvdj";
			var v17: multiset<bool> := multiset{v5};
			var v18: array<int> := new int[19] [v2.f40, v0, 0x9e, v2.f40, v2.fm1(globalState), fm39(globalState), |v16|, v2.f40, v2.f40 % |v17|, v2.f40 + v2.f40, v2.fm1(globalState) * v0, v0, v2.f40, v2.f40 + |v14|, v0 / v0, -336, |(v17 + multiset{v5})|, 416, v1[|f29|] / v0];
			v18 := v18;
			var v19: map<int, string> := map[v6.fm13(v12, v10, v2.f40, globalState) := v16];
			var v20: set<int> := {v0, v2.f40};
			var v21: map<string, int> := map[if (|v20| in v19) then v19[|v20|] else "xicujy" := if (v5) then |multiset(v16)| else --v2.f40];
			v21 := v21[v16 := |("akwbgonr" + v16)|];
			var v22 := DC10(v18);
			var v23: map<bool, array<int>> := map[v5 := v22.cf21];
			var v24: map<array<int>, bool> := map[if (v5 in v23) then v23[v5] else v18 := fm40(0x175, v5, true, globalState) !in v16[v2.f40 := v16[v2.f40]]];
			v24 := v24[v18 := v2.fm44(globalState)];
			var v25 := DC32(v17, v8[v2.f40], v5, v5, v2.f40);
			var v26 := DC11(v10);
			var v27: array<char> := new char[27] [v10, v10, v10, v10, v10, v10, v10, v10, v26.cf22, v10, v10, v10, v10, v10, v10, v10, 'v', v10, v10, v10, v10, v10, v10, v10, v10, v10, 'c'];
			var v28 := DC4(v27, v2.f40, v2.f40, v16, v16);
			var v29: map<int, int> := map[v0 := |v20|];
			var v30: seq<set<int>> := [{if (v2.f40 in v29) then v29[v2.f40] else |v29|, v2.f40}];
			var v31: array<bool> := new bool[22] [v5, v5, v5, v5, false, !v5, if (v25.cf58) then v5 else false, false, v28.cf10 < "kyjjbux", v5, v5, v5, v5, v5 || v5, v5, v20 <= v30[v2.f40], !true && v5, v5, v5 && v5, false, false, !v5];
			v31[536] := v5;
			v31[536] := !(v2.f40 != (if (v5) then -v0 else v0));
			v18[663] := v2.f40;
			v18[663] := v2.f40;
		} else {
			r1 := (if (v6 in v7) then v7[v6] else v0) < v0;
			var v32 := "vaiuecm";
			globalState.f6 := v32;
			var v33: array<int> := new int[18];
			v33[192] := -728;
			var v34: map<int, map<bool, int>> := map[0x1e0 := map[v5 := v0]];
			var v35: T0 := new C7(v5, v34);
			var v36: set<T0> := {v35};
			v33[192] := |(if (!(v5 && v5)) then v36 + v36 else v36 - v36)|;
			var v37 := DC10(v33);
			var v38: map<D3, int> := map[v37 := v2.f40];
			var v39: multiset<bool> := multiset{v5, v5};
			var v40 := DC32(v39, v5, !v5, v5, v2.f40);
			var v41: C1 := new C1(|f29|, |fm52(v5, v33[192], globalState)|, v40.cf57, v5);
			var v42: seq<C1> := [v41];
			var v43 := new C4(v38, v5, v42[v41.f36 := v41] != [v41, v41]);
			var v44: seq<string> := [v32];
			var v45: array<string> := new string[25] [v32, v32, v32, v44[v33[192]], v32, v32, v32, "hiqyfxr", v32, v32[-0x190 := v10] + "mfc", "bsbqonwkb", v32, v32, v32, v32, v32[v2.f40 := v10], v32, v32, v32, v32, "kasxjomfh", v32, v44[v2.f40], v32, v32];
			v45[935] := v32 + (seq(0x11b, i0  => ('c')));
			v45[935] := v32;
		}
		
		var v46 := "tgaaxgqsq";
		var v47 := DC16(v46 + v46);
		v47 := if (v5) then v47 else v47.(cf29 := v46);
		for i1 := -597 * v0 to v2.f40 - v0 {
			v10 := v10;
			var v48 := new C0(true);
			globalState.f22 := !(if (v5) then v5 else v48.f32);
			var v49 := DC1(true, v48.f32, v2.f40, v48.f32);
			var v50: set<D0> := {v49};
			v50 := v50 - (v50 + {v49});
		}
		r0 := v5;
		r1 := v8[v2.f40];
	}
}

class C9 extends T2, T3 {
	var f30 : bool
	const f31 : int
	constructor (f30 : bool, f31 : int, f26 : bool, f27 : bool, f29 : multiset<int>) {
		this.f30 := f30;
		this.f31 := f31;
		this.f26 := f26;
		this.f27 := f27;
		this.f29 := f29;
	}
	
	function fm3(p0: string, p1: seq<bool>, p2: int, globalState: GlobalState): seq<bool> {
		([f26, f26] + [f27]) + ([f26] + [f27])
	}
	function fm4(p0: D0, p1: int, p2: int, globalState: GlobalState): int {
		f31 - f31
	}
	function fm1(globalState: GlobalState): int {
		f31
	}
	function fm2(p0: bool, p1: int, globalState: GlobalState): bool {
		f27 <==> f27
	}
	method m2(p0: bool, p1: T0, p2: bool, globalState: GlobalState) returns (r0: char, r1: int) {
		var v0: array<map<bool, bool>> := new map<bool, bool>[3];
		var v1: map<bool, bool> := map[f27 := f30];
		v0[333] := v1;
		globalState.f12, v0[333] := if (f30) then -f31 else f31 / f31, v1 + (v1 + v1);
		var v2: seq<bool> := [if (f27 in v0[333]) then v0[333][f27] else p0, !f30, fm2(true, |map[f27 := f31]|, globalState)];
		var v3 := 'v';
		var v4: map<int, char> := map[|v2| := v3];
		var v5: map<int, bool> := map[f31 := f30];
		var v6: set<char> := {if (|v5| in v4) then v4[|v5|] else v3, v3};
		var i0 := 0;
		while (v6 !! v6)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v7 := DC0(p2);
			var v8: map<int, int> := map[f31 := f31];
			match (v7.(cf0 := map[fm1(globalState) := f31] != v8)) {
				case DC1(cf1, cf2, cf3, cf4) =>
					globalState.f22 := !v7.cf0;
					v4 := v4;
					var v9: array<bool> := new bool[14](i1 => p2);
					v9[924] := if (true) then f26 else f30;
					v9[924] := cf4;
					globalState.f12 := if (cf3 in f29) then f29[cf3] else f31;
				case DC0(cf0) =>
					var v10: array<bool> := new bool[25](i2 => p0);
					var v11: seq<array<bool>> := [v10, v10, v10, v10, v10];
					globalState.f12 := (if (p0) then 0x2b4 else f31) - |(v11 + v11)|;
					var v12 := "bex";
					var v13: array<int> := new int[24] [f31, -773, |v12|, f31, f31, f31, f31, f31, f31, f31, f31, f31, f31, f31, f31, f31, f31, |v2|, f31, f31, |v12|, f31, f31, |[v2, v2]|];
					var v14 := DC10(v13);
					var v15: map<D3, int> := map[v14 := f31];
					var v16 := new C4(v15, cf0, !f26);
					globalState.f2 := true;
					v13[971] := f31;
					v13[971] := f31;
				case DC2(cf5) =>
					var v17: array<string> := new string[15];
					var v18 := "dbooou";
					v17[893] := v18;
					var v19: set<bool> := {f30};
					v17[893] := if (v19 > v19) then v18 else v18 + v18;
					var v20: set<int> := {f31, f31};
					var v21: array<char> := new char[20];
					var v22 := DC4(v21, 0x2db, 0x117, seq(60, i3  => (v3)), "l");
					globalState.f12 := -(|v18| % |v20|) % v22.cf8;
					var v23: map<bool, int> := map[p0 := f31];
					globalState.f2 := if (f27) then false else p1.fm2(fm2(f30, |v23|, globalState), f31, globalState);
					globalState.f2 := fm2(f31 >= f31, f31 + 0x6f, globalState);
			}
			
			var v24: seq<char> := [v3, v3];
			var v25: seq<seq<char>> := [v24, seq(0xa5, i4  => (v3))];
			v25 := v25;
			if (f30) {
				var v26: array<bool> := new bool[4];
				v26[357] := !p1.fm2(f26, f31, globalState);
				v26[357] := fm2(f26, f31, globalState);
				var v27: seq<array<bool>> := [v26];
				var v28: array<int> := new int[17](i5 => i5 + f31);
				v27, globalState.f24, v28, globalState.f16 := v27, p1.fm2(v26[357], f31, globalState), v28, f31 % f31;
				var v29: seq<int> := [|[f31]|];
				v28[292] := if (p0) then f31 else |v29|;
				var v30: seq<seq<array<bool>>> := [v27, v27];
				globalState.f12, f30, v28[292], globalState.f15 := -f31, p1.fm2(v26 in v30[0xae], 0x16, globalState), f31, f30;
				v26 := v26;
				var v31: map<D3, int> := map[DC10(v28) := v28[292]];
				var v32 := new C4(v31, p2, !p0);
			} else {
				var v33: map<bool, int> := map[p2 := f31];
				var v34, v35, v36 := m6(v33, -f31, v3, f31, globalState);
				v3 := 'o';
				var v37: array<bool> := new bool[3](i6 => f26);
				v37 := v37;
				var v38: map<T0, bool> := map[p1 := p0];
				var v39 := DC27(p1);
				globalState.f24 := if (v39.cf48 in v38) then v38[v39.cf48] else p0;
				globalState.f24 := f26;
			}
			
			var v40: set<map<bool, bool>> := {v1, v0[333][p0 := f26], v1};
			v40 := v40 + v40;
		}
		var v41 := DC28(0xe5);
		match (v41) {
			case DC28(cf49) =>
				globalState.f22 := if (f30) then if (f30) then f30 else p2 else f30;
				globalState.f2 := v2[0xed];
				var v42 := new C5(false, false);
				globalState.f12 := f31;
			case DC27(cf48) =>
				var v43: array<int> := new int[26](i7 => i7 + f31);
				var v44 := DC10(v43);
				var v45: map<D3, int> := map[v44 := |[f30]|];
				var v46 := new C4(map[v44 := f31] + v45, p2, if (p2) then p0 else f27);
				if (f26) {
					var v47: set<int> := {v46.fm1(globalState), f31};
					v47 := v47;
					var v48 := "vsjg";
					globalState.f6 := v48;
					var v49: set<string> := {v48, v48, v48};
					var v50: multiset<char> := multiset{v3};
					var v51 := DC0(p2);
					var v52: set<bool> := {v2[f31]};
					var v53: T2 := new C1(|v48|, 0x320, f27, p0);
					var v54: seq<T2> := [v53, v53];
					var v55: map<seq<T2>, int> := map[v54 := f31];
					var v56: array<bool> := new bool[25] [v49 > v49, f27, f26, v50 >= v50, v51.cf0, f27 || f30, f31 <= f31, p0, !f26, !({f30, p2} < v52), cf48.fm2(f27, f31, globalState), !!(f30 || f27), p0, p0, f30, fm2(f27, |v55|, globalState), v53.f26, !(f27 <==> f27), !({379} !! v47), p2, f27 <== fm2(v53.f26, f31, globalState), f27, f30, v53.f26, v2[f31]];
					v56[56] := !v46.fm2(v53.f27, f31, globalState);
					v56[56] := v53.f26;
					globalState.f16 := |(v48 + (v48 + (seq(-0x364, i8  => (v3)))))|;
					var v57: seq<int> := [f31];
					globalState.f2 := v2[|v57| + f31];
				} else {
					var v58: map<int, int> := map[f31 := f31];
					var v59: seq<int> := [|v58|, f31];
					var v60: T3 := new C8(multiset(v59));
					var v61 := "qtfvyyvhs";
					var v62: map<string, seq<bool>> := map[v61 := v2];
					var v63: map<T3, map<string, seq<bool>>> := map[if (f27) then v60 else v60 := v62];
					v63 := v63[v60 := v62];
					v43[435] := f31;
					v43[435], globalState.f22, globalState.f12, globalState.f12, globalState.f22 := f31, true, 0x303 - f31, f31, f27;
					var v64: multiset<int> := multiset{v43[435]};
					v64 := v64;
					var v65: multiset<map<int, bool>> := multiset{v5};
					var v66 := DC44(v65);
					globalState.f12, v43[435], globalState.f15, v64 := v46.fm1(globalState), |v66.cf82| * (if (f30) then f31 else |v59|), p0, v60.f29 * (v60.f29 - v64);
					var v67: multiset<char> := multiset{v3, v3, 'u', v3};
					var v68: array<bool> := new bool[7] [v67 >= multiset{v3}, p0, p2 <== true, !!f26, p2, f30, p0 || f27];
					v68[355] := f26;
					v68[355] := f27;
				}
				
				var v69 := DC7(p2, false, f31, |f29| - f31);
				globalState.f15, v69 := if (!p2) then p2 else p0, v69;
				var v70: array<bool> := new bool[23](i9 => f26);
				v70[795] := cf48.fm2(f26, f31, globalState);
				v70[795] := p0;
		}
		
		var v71: array<bool> := new bool[17];
		forall i10 | 0 <= i10 < v71.Length {
			v71[i10] := DC38(f31, 928, f30).cf71;
		}
		for i11 := f31 to f31 {
			var v72 := "eobm";
			globalState.f6 := (v72[f31 := 'e'] + v72)[i11 := v3] + v72;
			globalState.f6 := v72;
			var v73: map<int, string> := map[f31 := "yxokab"];
			var v74: seq<int> := [i11, f31];
			var v75: seq<seq<int>> := [v74];
			var v76: T2 := new C2(v75, f30, p2);
			var v77: array<T2> := new T2[3] [v76, v76, v76];
			var v78: seq<array<T2>> := [v77, v77];
			var v79: set<int> := {f31, i11, |v72|, |v73|, |(v78 + [v77, v77])|};
			v79 := (set v80 : int | (0x329 <= v80) && (v80 < 0x2b1) :: (v80 + f31)) * v79;
			globalState.f12 := i11;
		}
		globalState.f12 := f31;
		r0 := v3;
		r1 := f31;
	}
	method m1(p0: bool, p1: seq<string>, p2: seq<bool>, globalState: GlobalState) {
		var v0 := DC0(f27);
		var v1: seq<int> := [fm4(v0, f31, f31, globalState)];
		var v2: multiset<seq<int>> := multiset{v1, v1};
		var v3: C3 := new C3(v2, f30, f30);
		v3 := v3;
		var v4 := 's';
		var v5: map<bool, seq<int>> := map[p0 := v1];
		var v6: map<char, int> := map[v4 := |(if (f27 in v5) then v5[f27] else v1)|];
		var v7: array<char> := new char[15](i1 => v4);
		var v8 := "kbutbaoar";
		var v9 := DC4(v7, f31, f31, v8, v8);
		for i0 := v3.fm12(|v6|, globalState) + f31 to v9.cf8 {
			v8 := v8 + v8;
			var v10: array<int> := new int[3](i2 => i2 * |{f27, f27, p0}|);
			v10[745] := f31 * |"sw"|;
			v10[745] := f31;
			var v11: map<bool, seq<bool>> := map[f30 := fm3(v8, p2, f31, globalState)];
			var v12 := DC13(v11 + v11);
			match (v12) {
				case DC14(cf25, cf26, cf27) =>
					var v13: map<int, string> := map[|p2| := "hhvqkm"];
					var v14: map<bool, int> := map[f27 := cf25];
					var v15: seq<map<bool, int>> := [v14, v14];
					var v16 := DC24(if (i0 in v13) then v13[i0] else v8, v15, f26);
					v16 := v16;
					globalState.f12 := f31 % cf26;
					var v17: seq<seq<int>> := [v1, v1];
					var v18 := new C2(v17, true, f30);
					var v19: map<int, bool> := map[v3.fm1(globalState) := p0];
					var v20: multiset<map<int, bool>> := multiset{v19};
					var v21 := DC44(v20);
					v21 := DC44(v20);
				case DC13(cf24) =>
					globalState.f15 := f27;
					var v22: map<seq<int>, int> := map[v1 := -v10[745]];
					v22 := v22[v1 := v10[745]];
					v4 := v4;
					var v23 := DC16(v8);
					var v24: map<bool, int> := map[f26 := v10[745]];
					var v25: map<seq<int>, bool> := map[([i0] + v1)[v10[745] := |v23.cf29|] := v24 != v24];
					v10[745] := |v25|;
				case DC15(cf28) =>
					var v26: map<int, bool> := map[i0 := fm2(p2[f31], i0, globalState)];
					var v27 := DC37(v26[f31 := p0]);
					var v28: map<int, int> := map[-f31 := v10[745]];
					var v29: map<map<int, bool>, map<int, int>> := map[v27.cf68 := v28];
					var v31: map<map<int, bool>, int> := map[map[0x67 := f26] := v10[745]];
					var v32 := DC45(v31);
					v29 := v29 + (map v30 : map<int, bool> | v30 in v32.cf83 :: (v30) := (map[|map[false := f26]| := v10[745]]));
					var v33: map<int, seq<bool>> := map[|v8| := p2];
					var v34: map<bool, int> := map[p0 := |v8|];
					var v35 := DC34(v10[745], f26, f31);
					var v36: multiset<D15> := multiset{v35};
					var v37: map<seq<bool>, multiset<D15>> := map[if (|v34| in v33) then v33[|v34|] else p2 := v36 + v36];
					v37 := map[p2[v10[745] := f27] := v36];
					var v38: C0 := new C0(true);
					var v39: map<bool, multiset<int>> := map[v38.f32 := f29];
					var v40: multiset<bool> := multiset{p0};
					globalState.f19, v38, globalState.f16, globalState.f6, globalState.f15 := if (f30) then v4 else v4, v38, |(v39 + map[fm2(p0, v10[745], globalState) := f29])| * |[f26]|, v8, !(v40 >= v40);
					var v41: array<map<bool, bool>> := new map<bool, bool>[2];
					var v42: map<bool, bool> := map[f27 := f26];
					var v43 := DC29(v42);
					v41[893] := v43.cf50 + v42;
					v41[893] := v42;
			}
			
			var v44: map<array<int>, bool> := map[v10 := p2[|p2|]];
			v44 := if (f30) then v44 else v44;
		}
		var v45: multiset<bool> := multiset{f26};
		var v46: set<bool> := {false, f27};
		var v47: map<multiset<bool>, int> := map[v45 := |v46|];
		var v48 := DC1(p0, false, |v8|, p2[|v47|]);
		var v49 := DC2(v48);
		globalState.f12 := -match v49 {
			case DC1(cf1, cf2, cf3, cf4) => cf3
			case DC0(cf0) => 0x2b3
			case DC2(cf5) => f31
		};
		var v50: map<bool, bool> := map[(DC19(f31, v45, v3.fm2(false, f31, globalState), map[f26 := |(seq(0x167, i4  => (v4)))|]).(cf36 := f31, cf37 := v45)).cf38 := f27];
		var i3 := 0;
		while (if (p0 in v50) then v50[p0] else f27)
			decreases 100 - i3
		{
			if (i3 >= 100) {
				break;
			}
			
			i3 := i3 + 1;
			var v51: seq<multiset<int>> := [f29];
			var v52: map<bool, int> := map[f26 := |v8|];
			globalState.f2 := (|(seq(0x264, i5  => (f31)))| % -f31) > -(if (f27) then |{f29, v51[f31]}| else if (f26 in v52) then v52[f26] else f31);
			var v53: T1 := new C3(multiset{[f31], v1[f31 := -0x1d9]}, false, p0);
			var v54 := DC46(v53);
			var v55: array<T1> := new T1[28] [v53, v53, v53, v53, v53, v53, v53, v53, v53, v53, v53, v53, v53, v53, v53, v54.cf84, v53, v53, v53, v53, v53, v53, v53, v53, v53, v53, v53, v53];
			v55[309] := v53;
			v55[309] := v53;
			var v56: map<int, bool> := map[f31 := v53.f27];
			var v57: array<bool> := new bool[3] [if (f31 in v56) then v56[f31] else v53.f27, f26, f27];
			var v58 := DC21(map[p2 := v57]);
			var v59: array<map<bool, int>> := new map<bool, int>[20](i6 => v52);
			var v60: map<seq<bool>, array<bool>> := map[[p2[f31]] := v57];
			v3, v58, v59, globalState.f12 := v3, DC21(v60), v59, 110;
			var v61 := DC5(f31);
			v61 := v61;
		}
		var v62: array<int> := new int[15](i7 => i7 / f31);
		v62[335] := f31;
		var v63: array<array<int>> := new array<int>[3];
		v63[174] := v62;
		var v64: seq<set<bool>> := [{f27}];
		v62[335], globalState.f15, v63[174], v62 := |v64[f31]| - f31, f26, v62, v62;
		globalState.f15 := f27;
	}
	method m3(p0: int, p1: int, p2: string, p3: int, globalState: GlobalState) returns (r0: bool, r1: bool, r2: array<int>, r3: T0) {
		var v0: seq<int> := [p3];
		globalState.f16 := v0[f31];
		for i0 := p3 to -0x16c {
			var v1: array<map<bool, bool>> := new map<bool, bool>[29];
			v1 := v1;
			var v2: map<int, bool> := map[-f31 := f27];
			v2 := v2[p0 := f30];
			var v3: array<bool> := new bool[9](i1 => false);
			v3[913] := f27;
			globalState.f6, globalState.f2, v3[913] := (seq(-0x1d, i2  => ('a'))) + p2, (p2 + fm22(-f31, f30, globalState)) == p2, f26;
			globalState.f16 := 763;
		}
		var i3 := 0;
		while (!(p3 <= -0x18f))
			decreases 100 - i3
		{
			if (i3 >= 100) {
				break;
			}
			
			i3 := i3 + 1;
			var v4: map<int, bool> := map[p1 := f30];
			var v5: map<bool, int> := map[f30 := p1];
			var v6: map<int, map<bool, int>> := map[|v4| := v5];
			var v7: multiset<bool> := multiset{f26};
			var v8 := 'o';
			var v9, v10, v11 := m6(if (p1 in v6) then v6[p1] else fm30(|v7|, true, globalState), fm39(globalState) * |map[f26 := p0]|, v8, p1, globalState);
			var v12: array<int> := new int[5](i4 => i4 * |multiset{f31, p0}|);
			var v13: T3 := new C8(f29);
			var v14: multiset<T3> := multiset{v13};
			var v15: array<bool> := new bool[10] [f30, f26, f26, true, fm2(f26, |v14|, globalState), v11, v11, f30, f30, f26];
			var v16: seq<array<bool>> := [v15, v15];
			v12[919] := if (v11 in v7) then v7[v11] else |v16|;
			var v17: map<int, int> := map[|v5| := |p2|];
			var v18: map<bool, map<int, int>> := map[v11 := v17];
			var v19 := DC49(v18);
			v12[919] := (v9 % p0) - (v10 / |v19.cf86|);
			v15[586] := v11;
			var v20: map<bool, char> := map[f26 := v8];
			v15[586] := |v20| > p1;
			v15[586] := f27 !in multiset{f27, fm2(f27, v12[919], globalState), false};
		}
		if (f27) {
			var v21: array<bool> := new bool[1] [f26];
			var v22 := DC30(f30, f31, f30, v21);
			var v23: seq<bool> := [f30, f30];
			var v24: T0 := new C6(|map[true := |v23|]|);
			var v25: map<bool, T0> := map[f27 := v24];
			var v26: set<T0> := {if (f26 in v25) then v25[f26] else v24};
			var v27 := DC53(v26);
			var v28: multiset<bool> := multiset{f27};
			var v29: array<bool> := new bool[18] [p3 >= p0, f27, f27 <== f27, !v22.cf51, f30, f27, v27.cf97 != v26, f27, f30 <== f27, f30, f26, f26, f26 && f27, !(v28 !! multiset(v23)), f30, (if (f30 in v28) then v28[f30] else p3) < -0xfd, fm2(f26, |p2|, globalState), f27];
			v29 := v21;
			var v30: array<int> := new int[6](i5 => i5 - f31);
			var v31: map<int, bool> := map[p0 := f30];
			v30[247] := 952 + |v31|;
			v30[247] := p0;
			var v32: array<map<int, int>> := new map<int, int>[8];
			var v33: map<int, int> := map[p1 := v30[247]];
			v32[462] := v33 + v33;
			v21[441] := !f30;
			var v34 := DC1(v23[|[f26]|], true, p1, f26);
			var v35 := DC2(v34);
			v32[462], r0, v21[441], globalState.f12, globalState.f22 := v33, f27, v30[247] <= (if (f26 in v28) then v28[f26] else |v33|), f31, !(fm10(f31, p3, DC2(v35), globalState) >= multiset{p3, p3});
			var v36 := DC10(v30);
			var v37: map<D3, int> := map[v36 := p0];
			var v38: C4 := new C4(v37, v21[441], f27);
			var v39: multiset<array<int>> := multiset{v30};
			f30, globalState.f12, v38, globalState.f12 := v39 <= v39, v0[v24.fm1(globalState)], if (f27) then v38 else v38, p1;
			var v40: set<int> := {v24.fm1(globalState)};
			var v41: set<bool> := {!f26, v21[441]};
			var v42: map<int, set<bool>> := map[p3 := v41];
			var v43 := DC50(f26, |v40|, v21[441], |v42|);
			match (v43) {
				case DC50(cf87, cf88, cf89, cf90) =>
					var v44: array<array<bool>> := new array<bool>[12];
					v44[562] := v21;
					v21[441], v44[562], v30[247], globalState.f6, globalState.f6 := cf89, v29, cf88, p2 + p2, p2;
					var v45: array<C4> := new C4[1];
					v45[519] := v38;
					r0, v30[247], globalState.f6, v45[519] := !(p2 < p2), f31, "icsw", v38;
					var v46: array<seq<int>> := new seq<int>[15](i6 => [p3]);
					v46[401] := seq(0x33b, i7  => (-|p2|));
					var v47 := DC28(p3);
					v46[401] := v0 + ([p0])[v0[v47.cf49] := p3];
					var v48: map<bool, int> := map[f26 := |[true]|];
					var v50 := DC41(!({|v48|, p3} > (set v49 : int | v49 in v40 :: (v49 - |"xxuvuil"|))), f26, p0);
					v50 := v50;
				case DC51(cf91, cf92, cf93, cf94) =>
					var v51: map<bool, int> := map[f27 := -449];
					var v52: C7 := new C7(cf92.f27, map[cf91 := v51]);
					var v53: map<C7, int> := map[v52 := p3];
					v30[247] := DC42(|v53|, p1, p3).cf79;
					globalState.f6 := if (p2 <= p2) then if (!f30) then p2 else "hsr" else p2;
					v30[247] := p1 + cf94;
					var v54: C1 := new C1(p3 % f31, cf93, v21[441], p0 == cf93);
					v54 := new C1(v54.f37 * p3, p0, !v21[441], cf92.f26);
				case DC52(cf95, cf96) =>
					globalState.f12 := p1;
					v0 := v0 + (v0 + v0[-p0 := p0])[p1 := f31];
					var v55: multiset<seq<int>> := multiset{v0};
					var v56 := new C3(v55, true, f27);
					var v57: array<array<multiset<bool>>> := new array<multiset<bool>>[7];
					var v58: array<multiset<bool>> := new multiset<bool>[9](i8 => v28);
					v57[36] := v58;
					globalState.f16, v57[36], v30, globalState.f16 := (v0[p1] + p0) - -v0[-797], v58, v30, |({|v28|} * {p1, f31, p3})|;
				case DC49(cf86) =>
					var v59 := new C4(v38.f33 + v38.f33, p2 != p2, f30);
					var v60: map<bool, bool> := map[v21[441] := v21[441]];
					var v61 := new C4(if (f30) then v38.f33 else v59.f33, v21[441], if (!v21[441] in v60) then v60[!v21[441]] else f26);
					var v62: T1 := new C4(v59.f33, f27, f26);
					var v63 := DC51(f31, v62, 875, f31);
					v30[247] := v63.cf91;
					var v64 := 'r';
					var v65: set<char> := {v64, v64};
					var v66: C6 := new C6(|v65|);
					var v67: multiset<C6> := multiset{v66};
					var v68: map<multiset<C6>, int> := map[v67 := p3];
					var v69 := DC37(v31);
					var v71: map<bool, char> := map[!false := v64];
					var v72 := DC32(multiset{v62.f27}, f26, f27, v21[441], |v71[v62.f27 := v64]|);
					var v73: map<D14, bool> := map[v72 := v62.fm2(f30, p1, globalState)];
					var v74: map<int, map<D14, bool>> := map[f31 := v73];
					var v75: array<map<int, bool>> := new map<int, bool>[6] [v31[|v68| := f26], v31, v31, v31 + v69.cf68, map[v59.fm1(globalState) := false], map v70 : int | v70 in v74 :: (v70 + p1) := (v62.f26)];
					v75[341] := v31;
					v75[341] := if (v21[441]) then v31 else v31 + v31;
			}
			
		} else {
			var v76: set<bool> := {f30};
			var v77 := DC17(|v76|, p0, p3, p1, -p0);
			var v78: map<bool, bool> := map[false := f26];
			var v79 := DC38(|v78|, p1, f30);
			var v80: map<bool, int> := map[f30 := |{v79}|];
			var v81 := DC0(f30);
			var v82: map<int, string> := map[fm4(v81, p0, -p3, globalState) := "k"];
			var v83: multiset<bool> := multiset{f27};
			var v84: set<int> := {f31, if (fm2((DC19(|"smafhihgt"|, v83, f26, v80).(cf36 := f31)).cf38, p0, globalState) in v83) then v83[fm2((DC19(|"smafhihgt"|, v83, f26, v80).(cf36 := f31)).cf38, p0, globalState)] else p0, p0};
			var v85: map<int, int> := map[f31 := f31];
			v77 := fm53(if (f27) then |multiset{p3, p3, |v80|}| else |(if (p0 in v82) then v82[p0] else p2)|, p1, v84, v85, globalState);
			globalState.f24 := (p2 + p2) < (p2 + p2);
			globalState.f15 := f27;
			globalState.f16 := p3 + p1;
			var v86: seq<seq<int>> := [seq(0x2c0, i9  => (p1))];
			var v87 := new C2(v86[-f31 := [p0]], f27, f27);
		}
		
		var v88: array<string> := new string[20];
		forall i10 | 0 <= i10 < v88.Length {
			v88[i10] := seq(0x2a8, i11  => (if (f30) then 'u' else 'e'));
		}
		var v89 := DC0(f30);
		var v90: seq<bool> := [!!false];
		var v91: map<bool, int> := map[f30 := f31];
		var v92: array<bool> := new bool[26] [true, fm2(f26, f31, globalState), f30, -0x21f > p1, f30 <== f27, v89.cf0, v90[|v91|], v0 <= v0, fm1(globalState) <= -p1, f27, f27, f26, "o" <= (seq(0xe, i13  => ('y'))), f30, true, f27, f27, f27, f26, f30, f26, f30 <== f27, f30, f30, f30, false];
		forall i12 | 0 <= i12 < v92.Length {
			v92[i12] := !(if (f27) then f26 else false);
		}
		r0 := fm2(f27, p0, globalState) <==> (multiset(v0) == f29);
		var v93: C0 := new C0(f27);
		var v94: map<bool, C0> := map[f30 := v93];
		r1 := |v94| >= f31;
		var v95: array<int> := new int[5] [p3, f31, f31, if (f30) then f31 else f31, f31];
		r2 := v95;
		var v96: T0 := new C6(p0);
		r3 := v96;
	}
	method m4(globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0 := DC16(seq(-0xac, i1  => ('m')));
		var i0 := 0;
		while (match v0 {
			case DC17(cf30, cf31, cf32, cf33, cf34) => true
			case DC16(cf29) => !true
		})
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v1: map<bool, bool> := map[f26 := f27];
			var v2: seq<map<bool, bool>> := [v1 + v1, v1];
			v2 := v2;
			var v3: set<int> := {f31};
			var v4: seq<int> := [-0x1d5, |map[|v3| := f26]|, f31];
			var v5: seq<seq<int>> := [v4, v4, v4];
			var v6 := new C2(v5, f30, f26);
			globalState.f2 := (f29 + f29) >= multiset(v4 + v4);
			var v7: seq<bool> := [f30];
			var v8: map<int, int> := map[f31 := v6.fm1(globalState)];
			var v9: map<bool, int> := map[f27 := |v8|];
			var v10: map<int, map<bool, int>> := map[|v7| := v9];
			var v11 := new C7(f30, v10);
		}
		var v12 := new C6(f31);
		var v13: set<bool> := {f30};
		v13 := v13;
		if (f27) {
			var v14 := DC39(fm54(true, f31, v12.f40, globalState));
			match (v14) {
				case DC38(cf69, cf70, cf71) =>
					var v15 := "nq";
					var v16 := DC11(fm24(v15, f26, cf70, globalState));
					var v17: array<bool> := new bool[27](i2 => !false);
					v17[77] := f30;
					var v18 := 'f';
					var v19: map<int, bool> := map[cf69 := true];
					var v20: array<int> := new int[19] [v12.f40, -v12.f40, v12.f40, |v19|, f31, v12.f40, -0x131, 0x27b, cf69, cf70, v12.f40, cf70, cf70, v12.f40, cf69, v12.f40, f31, |v15|, f31];
					var v21: multiset<array<int>> := multiset{v20, v20};
					var v22: set<int> := {cf69, if (v20 in v21) then v21[v20] else v12.f40, v12.f40};
					v16, v17[77], cf69, globalState.f16 := DC11(v18), v13 < (v13 - {f30, f30, f26}), v12.f40, |v22|;
					v12.f40 := v12.f40;
					globalState.f16 := cf69;
					var v23: seq<int> := [f31];
					globalState.f15 := cf69 in [|v23|];
				case DC37(cf68) =>
					var v24: array<bool> := new bool[18](i3 => f26);
					var v25: set<int> := {f31};
					var v26 := DC28(v12.f40);
					var v27: array<int> := new int[2] [|v25|, v26.cf49];
					v24[1] := !(|[v27, v27]| <= f31);
					v24[1] := f27;
					globalState.f2 := !v12.fm44(globalState);
					v27[117] := f31 % -0x301;
					v27[117] := v12.f40;
					globalState.f2 := v24[1] <==> !f30;
				case DC39(cf72) =>
					var v28 := "veta";
					var v29: map<bool, int> := map[f26 := |v28|];
					var v30 := 't';
					var v31: multiset<char> := multiset{'q', v30};
					v29 := v29[multiset{v30, v30, v30} >= v31 := v12.f40 % -v12.f40];
					r1 := !v12.fm44(globalState);
					var v32: array<int> := new int[21](i4 => i4 + v12.f40);
					var v33: seq<bool> := [true, f26, f30, f26];
					var v34: seq<int> := [|v33|, v12.f40];
					var v35: set<int> := {v12.f40, fm1(globalState)};
					v32 := new int[4] [-v34[fm39(globalState)], -|v35| - f31, |v28|, f31];
					var v36: array<bool> := new bool[10];
					v36[902] := f30;
					v36[902] := if (f26) then !f27 else f27;
			}
			
			var v37: seq<int> := [v12.f40, v12.f40 % |f29|];
			var v38 := "ivqetbs";
			globalState.f12 := v37[|v38|];
			var v39: array<bool> := new bool[2](i5 => |v38| <= f31);
			v39[191] := v12.f40 <= 0x314;
			v39[191] := if (f31 > |f29|) then f30 else f26;
			var v40: map<int, bool> := map[v12.f40 := !f30];
			v40 := v40;
			globalState.f15 := v12.fm44(globalState);
		} else {
			var v41: C5 := new C5(f27, !f26);
			v41 := v41;
			globalState.f16 := v12.f40 / (v12.f40 + 844);
			var v42 := DC19(v12.f40, multiset([f30, f26]), DC1(f27, f26, 0x158, f30).cf4, fm30(v12.f40, f27, globalState));
			var v43: map<bool, int> := map[v12.fm44(globalState) := f31];
			var v44: map<map<bool, int>, bool> := map[(v42.(cf39 := v43)).cf39 := f27 || f30];
			v44 := v44[v43[f30 := 990] := !true];
			if (f26) {
				var v45: array<int> := new int[25];
				v45[166] := f31;
				v45[166] := 0x219;
				var v46 := "ohnwk";
				var v47: map<int, string> := map[-558 := v46];
				var v48: C1 := new C1(|(if (v12.f40 in v47) then v47[v12.f40] else v46)|, v12.f40, f27, f27);
				var v49: map<int, C1> := map[v48.f36 := v48];
				var v50: seq<C1> := [v48, v48, if (f31 in v49) then v49[f31] else v48, v48];
				var v51: map<seq<C1>, int> := map[v50 := v48.f37];
				v51 := v51;
				var v52: seq<int> := [|v46|];
				var v53: seq<int> := [|v52|];
				var v54 := DC0(!true);
				var v55: multiset<seq<int>> := multiset{v52[fm4(v54, 0x2c9, 0x1c5, globalState) := v12.f40], [-0x38f], v52};
				var v56: C3 := new C3(v55, f27, f26);
				var v57: map<bool, C3> := map[f26 := v56];
				var v58 := 's';
				var v59 := DC55(v56);
				v56 := if ((v58 !in v46[v12.f40 := v58]) in v57) then v57[v58 !in v46[v12.f40 := v58]] else if (f26 in v57) then v57[f26] else v59.cf103;
				globalState.f2 := f27 && f30;
				globalState.f12 := (v45[166] / v12.f40) % v12.f40;
			} else {
				var v60: multiset<int> := multiset{v12.f40};
				v60 := v60;
				var v61: multiset<D1> := multiset{DC3([f26])};
				var v62: seq<bool> := [f26, f27];
				var v63 := DC3(v62);
				var v64: seq<D1> := [v63.(cf6 := v62), v63];
				var v65: set<int> := {|v62|};
				var v66: map<set<int>, int> := map[v65 := 0x3d3];
				var v67 := DC59(v61);
				var v68: array<multiset<D1>> := new multiset<D1>[27] [v61, v61 + v61, v61, v61[v63 := |v62|], v61[v63 := v12.f40] * v61, v61, multiset([v63, v63, v63]) + v61, v61, multiset{v63.(cf6 := v62)} + v61, multiset(v64), v61 + v61, multiset{v63, v63.(cf6 := v62)}, multiset(v64) + v61, v61, v61[v63 := |v66|], v61 - v67.cf109, v61, v61, v61, (multiset{v63})[v63 := f31], multiset{v63, v63, v63, v63}, v61 * v61, v61, multiset{v63, DC3(v62[f31 := f26])}, v61, v61, v61];
				v68[641] := multiset(v64);
				var v69: map<int, bool> := map[|map[v12.f40 := f30]| := f26];
				v68[641], globalState.f12, globalState.f16, globalState.f16, v12.f40 := v61[DC3(v62) := if (v41.fm2(f27, |v69|, globalState)) then v12.f40 else 120], -(v12.f40 % v12.f40), v12.f40, v12.f40 + v12.f40, if (f27) then fm39(globalState) else v12.f40;
				var v70: array<array<bool>> := new array<bool>[12];
				globalState.f5 := if (f27) then v70 else v70;
				v62 := if (f27) then v62 else [f30];
				var v71: multiset<bool> := multiset{f26};
				var v72: seq<multiset<bool>> := [multiset{f27, true}, v71, multiset{false}];
				var v73: map<bool, multiset<bool>> := map[f26 := v71];
				var v74 := DC32(v71, f27, f26, f30, 0x197);
				var v75: array<multiset<bool>> := new multiset<bool>[15] [v71, v71, v71, fm21(f31, globalState), v72[v12.f40], v71 - v71, if (f30 in v73) then v73[f30] else v71, v74.cf56, v71, v71, multiset(if (f27) then v62 else v62), v71, v72[v12.f40], multiset([fm2(f27, 0x31, globalState)]), v71];
				var v76: array<int> := new int[19](i6 => i6 * -v12.f40);
				var v77: seq<array<int>> := [v76];
				v75[434] := fm21(|v77|, globalState);
				var v78: seq<int> := [v12.f40 + v12.f40];
				var v79: map<int, int> := map[v12.f40 := v12.f40];
				globalState.f16, v12.f40, v75[434], globalState.f22 := if (f30) then v12.f40 else f31 % v12.f40, v78[fm1(globalState)], v71, fm2(v12.fm2(f27, v12.f40, globalState), v12.f40 / |v79|, globalState);
			}
			
			var v80: seq<seq<int>> := [[f31, f31]];
			var v81: seq<seq<seq<int>>> := [v80];
			var v82: T2 := new C2(if (f30) then v81[v12.f40] else v80, !f26 <== true, f27);
			v82 := v82;
		}
		
		var v83 := "uuuanc";
		var v84: seq<int> := [0x374 + |v83|, v12.f40];
		var v85 := 'i';
		var v86: map<bool, char> := map[f30 := v85];
		v84 := v84 + (v84 + fm43(f31, v12.f40, |"yjmphn"|, v86, globalState));
		var v87: array<int> := new int[14];
		v87[537] := f31;
		var v88: array<bool> := new bool[12](i7 => f27);
		var v89: set<array<bool>> := {v88};
		var v90: seq<set<array<bool>>> := [v89, v89, v89];
		var v91: map<int, set<array<bool>>> := map[v12.f40 := v89];
		v87[537] := |(v90[-v12.f40] + (if (v12.f40 in v91) then v91[v12.f40] else v89))|;
		var v92: seq<bool> := [f26];
		var v93 := DC38(0x2ce, |v92|, f26);
		r0 := match v93 {
			case DC38(cf69, cf70, cf71) => f26
			case DC37(cf68) => !(v87[537] < |v92|)
			case DC39(cf72) => !(v13 != DC12({f30, true}).cf23)
		};
		r1 := f27;
	}
	method m5(p0: int, p1: int, p2: bool, globalState: GlobalState) returns (r0: string, r1: int, r2: bool) {
		var v0 := "flsqn";
		var v1: seq<int> := [|v0|, p1];
		var v2: array<int> := new int[14] [p1, |[f30]|, f31, f31, f31, |v1|, |map[f31 := p1]|, f31, p1, f31, p0, p1, p1, p0];
		var v3: map<array<int>, int> := map[DC10(v2).cf21 := p1];
		v3 := v3[v2 := -p0];
		r1 := fm39(globalState);
		var v4: multiset<bool> := multiset{p2};
		var v5: set<int> := {|v1|, p1, p0, |"ujlp"|, -|v4|};
		var v6: map<bool, set<int>> := map[p2 := v5];
		var v7: seq<set<int>> := [{p1, -51, |map[f31 := f27]|}, {p0}, if (f27 in v6) then v6[f27] else v5];
		var v8: map<seq<set<int>>, bool> := map[v7 := f26];
		var i0 := 0;
		while (!(if (v7 in v8) then v8[v7] else f30))
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v9: seq<bool> := [f30, true];
			var v10: map<seq<bool>, int> := map[v9[p1 := true] := p0];
			var v11: C5 := new C5(f26, f30);
			var v12: map<int, C5> := map[f31 := v11];
			var v13: set<bool> := {f30};
			var v14: set<set<bool>> := {v13};
			var v15: map<int, bool> := map[|v14| := f26];
			var v16 := DC38(|v12|, |v15|, f27);
			var v17: seq<D16> := [v16, v16];
			r1 := fm4(DC0(f26), if ([f27, f30] in v10) then v10[[f27, f30]] else p1, |v9| % |v17|, globalState);
			globalState.f16 := 0x9a;
			var v18: map<bool, bool> := map[if (!f27) then f26 else p2 := f30 <== f26];
			v18 := v18[p2 := f30];
			var v19: multiset<seq<int>> := multiset{v1};
			var v20 := new C3(v19, true, p2 in multiset{f30});
		}
		r1 := f31;
		globalState.f15 := false;
		if (fm2(p2, p1, globalState)) {
			var v21: array<char> := new char[6];
			var v22: array<D20> := new D20[16](i1 => if (f30) then DC47() else DC47());
			var v23 := DC47();
			v22[120] := v23;
			globalState.f6, v21, v2, v22[120] := v0 + (seq(-0x18b, i2  => ('k'))), v21, v2, v23;
			globalState.f16 := |((v4 + v4) - v4)|;
			r1 := -p0;
			var v24: map<int, bool> := map[|v5| := f26];
			var v25 := DC0(f27);
			var v26: seq<bool> := [f26, f26];
			if (|(if (p2) then v24 else map[p1 := f26])| > fm4(v25, |v26|, p0, globalState)) {
				var v27: map<int, int> := map[p0 := p0];
				globalState.f12 := |(v27[p0 := p0] + v27)|;
				var v28: array<multiset<int>> := new multiset<int>[15];
				v28[654] := f29;
				var v29 := 'p';
				var v30 := DC16(v0);
				var v31: array<string> := new string[16] [v0, v0, v0, v0, v0, (seq(0x371, i3  => ('h'))) + v0, "gixiulo", seq(0x209, i4  => ('f')), "gxxceu", if (f30) then "qkrwt" else v0, v0[|v0| := v29], v0, v0, v0 + v0, v0, v30.cf29];
				v31[172] := v0 + fm22(|"aj"|, f30, globalState);
				var v32: seq<multiset<bool>> := [multiset(v26), (v4[true := v1[0x2f5]])[f27 := p1]];
				var v33: seq<string> := [fm22(p0, f26, globalState)];
				v4, v28[654], v31[172] := v32[if (f26 in v4) then v4[f26] else p1], (f29 - fm42(|f29|, globalState)) - f29, "pyblpod" + (v33[p1] + v0);
				globalState.f16 := p1;
				var v34: C0 := new C0(f30);
				v34 := v34;
				var v35: array<bool> := new bool[5](i5 => f27);
				var v36 := DC14(p1, p0, v34.f32);
				var v37: C1 := new C1(p0, p0, v36.cf27, v27 == v27);
				var v38: map<bool, int> := map[v34.f32 := -0x2b1];
				v35, v37, v26, globalState.f15, globalState.f15 := v35, v37, [v34.fm7(v38, v1, p2, globalState)], !f30, f30 ==> (false || p2);
			} else {
				var v39: map<int, int> := map[f31 := p1];
				var v40: map<map<bool, bool>, map<int, int>> := map[map[true := f26] := v39];
				v40 := v40;
				var v41 := 't';
				var v42, v43, v44, v45 := m0(v41, globalState);
				var v46 := new C0(p0 < f31);
				var v47 := new C2(seq(0x1b5, i6  => (v1)), f26, v46.f32);
				var v48: C5 := new C5(f26, f27);
				var v49: map<int, C5> := map[|v0| + p0 := v48];
				var v50: map<bool, bool> := map[v46.f32 := f27];
				v49 := v49[|(v50 + v50)| := v48];
			}
			
			var v51: set<bool> := {f26, f27};
			globalState.f16 := |v51|;
		} else {
			var v52: C5 := new C5(126 <= p0, if (p2) then true else f27);
			var v53: array<D6> := new D6[13](i7 => DC16(v0));
			var v54: array<char> := new char[13](i8 => 'r');
			var v55 := DC4(v54, p1, |v4|, v0, seq(-896, i9  => ('k')));
			v53[67] := DC16(v55.cf11);
			var v56: array<string> := new string[17](i10 => v0);
			v56[547] := seq(0x2c0, i11  => ('u'));
			var v57 := DC50(f30, -f31, f30, f31);
			v52, globalState.f2, globalState.f22, v53[67], v56[547] := v52, p1 == (p1 + ---0x382), f30, fm38(v57.cf89, -fm1(globalState), v0, globalState), seq(0x278, i12  => (if (false) then 'f' else 'e'));
			var v58: set<bool> := {f26};
			var v59: seq<set<bool>> := [v58, v58];
			var v60: map<seq<set<bool>>, string> := map[v59 := v0];
			globalState.f6 := if (v59 in v60) then v60[v59] else v0;
			r0 := "ofrmiv" + v56[547];
			globalState.f16 := f31;
			var v61: set<string> := {fm22(p0, f30, globalState), v56[547], if (p2) then v56[547] else "x"};
			var v62: map<string, int> := map[v56[547] := p0];
			v61, globalState.f2 := {v56[547]} + (set v63 : string | v63 in v62 :: (v63)), f26;
		}
		
		r0 := fm22(p1, !(p2 <==> f30), globalState);
		r1 := -fm4(DC0(f27), p0 + p1, p1, globalState);
		r2 := p2;
	}
	method m6(p0: map<bool, int>, p1: int, p2: char, p3: int, globalState: GlobalState) returns (r0: int, r1: int, r2: bool) {
		var v0: seq<bool> := [f27, f26];
		for i0 := p1 to |(v0 + v0)| {
			var v1: array<bool> := new bool[7](i1 => [|v0|, i0] == [p1, f31, |(seq(734, i2  => (p3)))|]);
			v1[908] := if (!f26) then f30 else f26;
			var v2: map<int, bool> := map[p1 := !f26];
			v1[908] := if (p3 in v2) then v2[p3] else true;
			var v3 := DC3([!f30]);
			var v5: seq<set<int>> := [set v4 : int | (435 <= v4) && (v4 < 0xb3) :: (v4 % |p0|)];
			var v6: set<int> := {0x221, -0x21b, fm1(globalState), 0x2eb, p3};
			var v7: set<bool> := {f26, true, fm2(v1[908], i0, globalState)};
			var v8 := DC17(|map[[v1[908], f27] := f31]|, f31, p3, -624, p1);
			globalState.f24, globalState.f15, v1[908], globalState.f12, v3 := fm55(globalState) != v5[p1 := v6], (|v7| + |[v1[908]]|) < f31, f30, v8.cf34, v3;
			globalState.f12 := fm39(globalState);
			globalState.f12 := i0;
		}
		var v9: map<int, bool> := map[p1 := !fm2(f30, f31, globalState)];
		var v10: map<int, bool> := map[|v9| := !f26];
		var v11 := "tvcq";
		var v12: map<bool, string> := map[(if (-0x321 in v9) then v9[-0x321] else fm2(f30, p1, globalState)) || f30 := v11];
		globalState.f6 := if (f27 in v12) then v12[f27] else "qntndfgm" + "jkyhfrdrq";
		var v13 := DC3(v0);
		if (match v13 {
			case DC4(cf7, cf8, cf9, cf10, cf11) => f27 <==> f30
			case DC5(cf12) => f27
			case DC3(cf6) => f26
		}) {
			if (!((p3 * 737) > -p1)) {
				var v14: array<bool> := new bool[25](i3 => f30);
				var v15 := DC26(v14);
				v15 := v15;
				globalState.f15 := (f27 && f30) <== (f26 && f30);
				var v16: map<int, char> := map[-0x43 / p3 := p2];
				v16 := v16[p1 := p2];
				var v17: multiset<int> := multiset{p3 % f31, p1};
				v17 := f29;
				globalState.f6, globalState.f6, globalState.f16, globalState.f19, globalState.f12 := v11, v11, f31, p2, p3;
			} else {
				var v18: multiset<seq<int>> := multiset{[p1]};
				var v19: C3 := new C3(v18, f26, f30);
				var v20: map<int, C3> := map[p3 := v19];
				var v21: set<map<int, C3>> := {v20, v20, map[p3 := v19] + v20, v20};
				var v22: array<array<bool>> := new array<bool>[6];
				var v23: array<bool> := new bool[7];
				v22[780] := v23;
				var v24: seq<map<bool, int>> := [p0, p0];
				var v25 := DC24(v11, v24, f30);
				v21, globalState.f12, globalState.f15, globalState.f15, v22[780] := v21, (f31 / f31) - (p3 * p1), p2 in v25.cf44, f26, v23;
				var v26: map<char, bool> := map[p2 := f27];
				v26 := v26[p2 := f30];
				globalState.f12, globalState.f2, v22[780] := f31, f26, v23;
				globalState.f15 := !false;
				globalState.f16 := p3;
			}
			
			var v27: array<bool> := new bool[25];
			v27 := v27;
			var v28: map<string, string> := map[v11 := seq(0x3a4, i4  => (p2))];
			globalState.f12 := |(v28 + v28[v11 := v11])|;
			v27[926] := false;
			v27[926] := "ioovj" != v11;
			var v29: seq<int> := [-f31, f31, p1];
			var v30: map<seq<int>, int> := map[v29 := p3];
			var v32 := DC20(seq(0x28, i5  => (p1)));
			var v33: seq<seq<int>> := [v29, v32.cf40];
			v30 := map v31 : seq<int> | v31 in v33 :: (v31) := (p3);
		} else {
			globalState.f24 := f27;
			var v34: multiset<bool> := multiset{f30};
			var v35: seq<int> := [p1, 72];
			var v36: map<int, map<bool, int>> := map[v35[p3] := p0];
			var v37: T0 := new C7(true, v36);
			var v38: set<int> := {p1};
			v34, globalState.f22, v37 := v34[f26 := |v38|], v35 <= v35, v37;
			var v39: seq<map<bool, int>> := [p0];
			var v40: map<bool, int> := map[f30 := -p3 - DC52(f31, v39).cf95];
			var v41 := DC57();
			var v42: map<D23, bool> := map[DC58(v41) := f27];
			var v43 := DC58(v41);
			var v44 := DC24(v11, v39, false);
			v40 := v40[if (f26) then f27 else if (v43 in v42) then v42[v43] else f26 := |multiset{|v11|, |v11|, |v44.cf44|, p1, |v0|}|];
			globalState.f2 := {f31, -0xe, f31, -0x1f7, f31} > v38;
			globalState.f16 := f31;
		}
		
		var v45 := new C8(multiset{p3});
		var v46: array<bool> := new bool[13](i7 => f30);
		forall i6 | 0 <= i6 < v46.Length {
			v46[i6] := f30 && (if (|v11| in v10) then v10[|v11|] else f26);
		}
		var v47: seq<int> := [p3];
		var v48: set<seq<int>> := {[922, fm1(globalState)], v47 + v47};
		v48 := v48;
		r0 := p3;
		r1 := fm1(globalState);
		r2 := !(p1 == p3);
	}
}

method Main() {
	var v0 := 0x3d3;
	var v1 := "ovdhmu";
	var v2: set<int> := {v0, -|v1|};
	var v3: multiset<seq<char>> := multiset{v1};
	var v4: array<array<bool>> := new array<bool>[8];
	var v5: array<int> := new int[17](i1 => i1 - |v2|);
	var v6: map<string, array<int>> := map[seq(0x1b2, i0  => ('j')) := v5];
	var v8 := 'v';
	var v9: set<char> := {v8};
	var v10: map<int, map<char, int>> := map[v0 := map v7 : char | v7 in v9 :: (v7) := (-0x1b9)];
	var v11: map<char, int> := map[v8 := v0];
	var v12: map<map<char, int>, int> := map[if (v0 in v10) then v10[v0] else v11 := v0];
	var globalState := new GlobalState(v2, v3, true, true, 'l', v4, v1, v6, false, true, 139, true, 0x42, -610, v5, false, 0x31c, true, true, 'w', v1, 0x3ae, true, 735, false, v12);
	var v14 := false;
	var v15: multiset<map<bool, int>> := multiset{map[v14 := v0]};
	var v16: map<bool, int> := map[false := v0];
	var v17: map<map<bool, int>, bool> := map[v16 := v14];
	globalState.f12 := |((map v13 : map<bool, int> | v13 in v15 :: (v13) := (DC1(v14, v14, v0, v14).cf2)) + v17)|;
	if (fm0(false, globalState) > v3) {
		var v18, v19, v20, v21 := m0(v8, globalState);
		var v22 := DC0(v20);
		v22 := v22;
		globalState.f12 := v0;
		var v23: multiset<int> := multiset{882};
		var v24: array<multiset<int>> := new multiset<int>[2] [multiset{v0}, v23];
		var v25 := DC3(v19);
		globalState.f2, v19, v24, v18 := v14, v25.cf6, v24, v18;
		var v26, v27, v28, v29 := m0(v8, globalState);
	} else {
		globalState.f15 := v2 !! v2;
		globalState.f24 := v14 || v14;
		var v30 := new C3(multiset{[v0]}, true, v14 && v14);
		var v31: map<bool, bool> := map[v14 := true];
		var v32: C1 := new C1(v0, |v31|, v14, false);
		var v33: map<C1, bool> := map[v32 := v14];
		var v34: seq<map<C1, bool>> := [v33];
		var v35: multiset<int> := multiset{|v1|};
		var v36: T3 := new C8(v35);
		var v37: multiset<T3> := multiset{v36, v36, v36, v36};
		var v38: map<int, int> := map[v0 + |v34[v32.f36]| := |v37|];
		v38 := v38[v0 := v0 + v0];
		var v39 := DC11('b');
		var v40: map<D3, bool> := map[v39 := v14];
		v40 := v40[fm56(globalState) := v14];
	}
	
	var v41: array<bool> := new bool[4](i2 => v14);
	v41[434] := v14;
	v41[434] := v14;
	var v42: seq<int> := [v0, v0, v0, v0];
	var v43: map<array<bool>, bool> := map[v41 := v14];
	var v44: map<bool, bool> := map[v41[434] := v41[434]];
	var v45: seq<bool> := [v14];
	var v46: multiset<int> := multiset{|v45|, v0, v0};
	var v47: T2 := new C9(v14, |v44|, !v41[434], !v14, v46);
	var v48: seq<T2> := [v47];
	var v49: multiset<int> := multiset{|v43|, v0, |v48|};
	var v50: C8 := new C8(multiset{|v42[|v49| := -790]|});
	var v51: array<C8> := new C8[2] [v50, v50];
	v51[758] := v50;
	v51[758] := v50;
	var v52: map<int, int> := map[v0 := v0];
	var v55: array<map<int, int>> := new map<int, int>[23] [v52, v52, v52, v52, map[|map[v14 := v14]| := v0], v52, v52, map v53 : int | v53 in v42 :: (v53 + (if (0xa5 in v52) then v52[0xa5] else v0)) := (v0), v52[|v1| := |[v47.f27, v14]|], v52, v52[v0 := |v1|], map v54 : int | v54 in v49 :: (v54 * |v16|) := (v0), v52, v52, v52, v52, map[|"etqlwqsyc"| := |v45|], v52, v52, v52, v52, v52, v52];
	var v56 := DC16(v1);
	var v57: map<array<map<int, int>>, int> := map[v55 := |(if (v47.f27) then v56.cf29 else v1)|];
	v0 := if (v55 in v57) then v57[v55] else v0;
	var v58: T0 := new C6(v0);
	var v59, v60 := v47.m2(v41[434], v58, v41[434], globalState);
	var v61, v62, v63, v64 := m0(v59, globalState);
	var v65: multiset<bool> := multiset{v41[434]};
	if ((multiset{v41[434]} - v65) !! v65[v63 := v60]) {
		var v66, v67 := v47.m2(v47.f27, v58, v47.f27, globalState);
		v5[49] := v0;
		var v68: multiset<seq<int>> := multiset{v42};
		var v69: C3 := new C3(if (true) then multiset{v42, [v67]} else v68, v58.fm2(v47.f27, v0, globalState), v41[434]);
		v5[49], v69 := v67 * (v0 + |"nhcidbx"|), v69;
		v44 := v44[v42[v0] == -v5[49] := v47.f26];
		var v70: seq<string> := ["oujjykne", v1];
		v69.m1(v47.f26, v70[-v0 := v1], fm25(v47.f26, v47.f26, v66, globalState), globalState);
		var v71: map<int, array<bool>> := map[370 := v41];
		v71 := v71[v5[49] := v41];
	} else {
		globalState.f2 := v47.f26 ==> v41[434];
		var v72: map<int, bool> := map[v60 := v14];
		var v73: map<bool, map<int, bool>> := map[v47.f26 := v72];
		var v75: set<bool> := {v47.f26, v47.f26};
		var v76: map<set<bool>, bool> := map[v75 := v47.f26];
		var v78: array<map<int, bool>> := new map<int, bool>[12] [v72, v72, map[v0 := v47.f27], map[|v1| := v47.f26] + map[|v42| := v47.fm2(v41[434], v60, globalState)], if (v63 in v73) then v73[v63] else v72, map v74 : int | (0x64 <= v74) && (v74 < 0x8) :: (v74 + v60) := (v47.f27), v72, v72, v72, v72, (map[v0 := v47.f26])[|v76| := v41[434]], map v77 : int | v77 in v49 :: (v77 % v60) := (v41[434])];
		v78[490] := v72;
		v78[490] := v72 + v72;
		v47.m1(v47.f27, [v1, v1], v62, globalState);
		v5[742] := v0 + |"rodtn"|;
		v5[742] := v0 + v0;
		var v79: C5 := new C5(v47.f27, v58.fm2(true, 0x90, globalState));
		var v80 := DC62(v51[758]);
		var v81: array<int> := new int[18];
		var v82 := DC10(v81);
		var v83: map<D3, int> := map[v82 := v0];
		var v84: T1 := new C4(v83, v47.f27, v47.f27);
		var v85 := DC60(v84, |v1|, |{|v62|}|, v84.f27);
		v5[742], v79, v51[758], v62, v0 := -755 - 0x8, v79, v80.cf115, v62, v0 / v85.cf112;
	}
	
	globalState.f2 := v41[434];
	var v86: array<C5> := new C5[27];
	var v87: C5 := new C5(v63, v14);
	v86[461] := v87;
	var v88: seq<T0> := [v58, v58];
	globalState.f19, globalState.f12, v86[461], globalState.f15, globalState.f22 := v59, v0, v87, !v14, [v58, v58, v58, v58, v58] <= (v88 + v88);
	var v89: seq<array<int>> := [v5, v5, v5];
	v5 := v89[|v16|];
	match (DC18(v47)) {
		case DC19(cf36, cf37, cf38, cf39) =>
			v41[434] := v58.fm2(v47.f26, cf36, globalState);
			v87.m7(multiset(v42) >= v46, v47.f26, -(cf36 / v60), v46, globalState);
			v5 := v5;
			var v90: map<array<int>, array<int>> := map[v5 := v5];
			var v91: map<map<array<int>, array<int>>, int> := map[v90 := v60];
			v91 := v91[v90 := v0];
		case DC18(cf35) =>
			var v92: map<string, int> := map[v1 := |v1|];
			var v93: map<int, bool> := map[|v52| := v41[434]];
			globalState.f16 := if ((seq(-0x73, i3  => (v8))) in v92) then v92[seq(-0x73, i3  => (v8))] else |v93|;
			if (multiset(v1 + v1) != multiset([v61])) {
				var v94 := DC0(cf35.f26);
				globalState.f16 := -((0x2f1 / v47.fm4(v94, v0, v0, globalState)) / 0x7b);
				globalState.f24 := v41[434];
				v41[434] := if (cf35.f26) then v41[434] else 's' !in v1;
				v41[434] := cf35.f27;
				var v95 := v87.m8(v94.(cf0 := v47.f26), cf35.f27, v65, v0 * v60, globalState);
			} else {
				var v96: map<bool, seq<bool>> := map[cf35.f27 := v45];
				v96 := v96[v47.f27 := v45];
				var v97: seq<seq<char>> := [v1, v1, v1, v1];
				var v98 := new C9(false, v0, v47.f27, v97[v0] != v1, v49);
				var v99: T1 := new C5(!cf35.f26, v47.f27);
				v99 := v99;
				globalState.f22 := v58.fm2(v14 ==> !cf35.f26, v60, globalState);
				var v100 := DC0(false);
				globalState.f16 := cf35.fm4(v100, |v45| % v42[v98.f31], if (v41[434]) then v0 else v0, globalState);
			}
			
			var v101 := new C6(v60 * v60);
			var v102 := DC0(false);
			var v103 := DC14(v0, v101.f40, v102.cf0);
			var v104 := DC15(v103);
			v104 := if (if (v41[434]) then cf35.f27 else true) then if (v63) then v104.(cf28 := v103) else v104 else v104;
	}
	
	var i4 := 0;
	while ((v65 - v65) >= v65)
		decreases 100 - i4
	{
		if (i4 >= 100) {
			break;
		}
		
		i4 := i4 + 1;
		var v105: map<int, bool> := map[563 := v41[434]];
		var v106 := DC37(v105);
		var v107: multiset<map<int, bool>> := multiset{v106.cf68, v105};
		var v108 := DC44(v107);
		match (v108) {
			case DC44(cf82) =>
				globalState.f2 := v47.f27;
				v16 := v16;
				v60 := v0 / v58.fm1(globalState);
				var v109: map<array<int>, bool> := map[v5 := v41[434]];
				v63, globalState.f12 := v47.f27, |((v109 + map[v5 := v47.f27]) + v109)|;
		}
		
		if (v41[434]) {
			globalState.f22 := true;
			var v110: array<string> := new string[1](i5 => "glxoh" + v1);
			var v111: map<int, string> := map[v60 := v1];
			v110[626] := (seq(-0x46, i6  => ('b'))) + (if (v0 in v111) then v111[v0] else v1);
			v110[626] := v1 + v1;
			var v112 := DC10(v5);
			var v113: map<D3, int> := map[v112 := v60];
			var v114: C4 := new C4(v113, v47.f27, v47.f26);
			var v115: map<C4, int> := map[v114 := v60];
			v115 := v115;
			globalState.f6 := v1;
			globalState.f6 := v110[626];
		} else {
			globalState.f24 := v47.f27;
			var v116 := new C1(|v42|, v60, v0 in v2, v47.f26);
			var v117: array<char> := new char[8](i7 => v61);
			v117[346] := v59;
			v117[346] := v61;
			v41[434] := v14;
			var v118: map<int, char> := map[v42[-v116.f37] := v8];
			v118 := v118[v60 := v117[346]];
		}
		
		var v119: map<seq<int>, int> := map[v42 := v0];
		var v120: set<C5> := {v86[461]};
		var v121: seq<set<C5>> := [v120, v120, v120, v120];
		v119 := v119[v42 + v42 := |v121[v0]|];
		var v122: array<T1> := new T1[8];
		v122 := v122;
	}
	for i8 := |{false, v41[434], v47.f26, v14}| to v60 {
		v41 := v41;
		var v123: set<bool> := {v47.f26, v47.f26, v41[434]};
		var v124 := DC5(-|v123|);
		globalState.f16, globalState.f12, globalState.f12, globalState.f24 := i8 - v124.cf12, v0, 855, v2 !! {i8};
		globalState.f12, v41[434] := -0x1c7, v47.fm2(false, |map[v60 := --0x3a4]|, globalState);
		var v125: multiset<seq<int>> := multiset{[v0, |"yrpw"|, i8, -i8, 0x7b]};
		var v126: C3 := new C3(v125, false, !v47.f27);
		var v127 := DC55(v126);
		v126 := v127.cf103;
	}
	var v128, v129 := v50.m4(globalState);
	v41[434] := !!v14;
}