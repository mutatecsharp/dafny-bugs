function abs(x: int): int {
	if (x < 0) then -1 * x else x
}
function safeIndex(x: int, length: int): int 
	requires length > 0
{
	if (x < 0) then 0 else if (x >= length) then x % length else x
}
function safeDivisionInt(x1: int, x2: int): int {
	if (x2 == 0) then x1 else x1 / x2
}
function safeModuloInt(x1: int, x2: int): int {
	if (x2 == 0) then x1 else x1 % x2
}
datatype D0 = DC0(cf0: int, cf1: bool, cf2: string, cf3: map<int, map<int, bool>>)
datatype D1 = DC2(cf5: char) | DC3(cf6: map<int, int>, cf7: int, cf8: D0) | DC4(cf9: int, cf10: bool, cf11: map<int, multiset<int>>) | DC1(cf4: seq<bool>) | DC5(cf12: D1)
datatype D2 = DC7(cf14: int, cf15: char) | DC8(cf16: array<bool>) | DC9(cf17: int) | DC6(cf13: array<int>)
datatype D3 = DC11 | DC10(cf18: seq<C0>)
datatype D4 = DC13(cf20: bool, cf21: int, cf22: array<bool>) | DC14(cf23: int, cf24: int, cf25: set<string>) | DC15(cf26: set<array<char>>, cf27: char) | DC12(cf19: map<bool, int>) | DC16(cf28: D4)
datatype D5 = DC18(cf30: int, cf31: int) | DC17(cf29: set<int>) | DC19(cf32: D5)
datatype D6 = DC21(cf34: bool, cf35: char, cf36: int, cf37: int) | DC20(cf33: map<multiset<bool>, int>)
class GlobalState {
	const f0 : int
	const f1 : int
	const f2 : map<char, int>
	const f3 : bool
	const f4 : int
	const f5 : int
	var f6 : int
	const f7 : multiset<int>
	const f8 : bool
	const f9 : int
	const f10 : int
	var f11 : int
	const f12 : int
	const f13 : bool
	var f14 : array<int>
	const f15 : int
	var f16 : int
	constructor (f0 : int, f1 : int, f2 : map<char, int>, f3 : bool, f4 : int, f5 : int, f6 : int, f7 : multiset<int>, f8 : bool, f9 : int, f10 : int, f11 : int, f12 : int, f13 : bool, f14 : array<int>, f15 : int, f16 : int) {
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
	}
	
}

class C0 {
	const f17 : bool
	constructor (f17 : bool) {
		this.f17 := f17;
	}
	
	function fm2(p0: bool, p1: bool, p2: bool, p3: string, globalState: GlobalState): string {
		(seq(abs(0x18c), i0  => ('e'))) + ("jcohubbi" + "bfaki")
	}
}

function fm0(p0: int, p1: map<int, int>, globalState: GlobalState): int {
	if (false) then 34 else safeDivisionInt(419, |[true, !false, false, false, true]|)
}
function fm1(p0: map<int, int>, p1: int, p2: int, globalState: GlobalState): map<int, int> {
	map[0x1bf := |map[true := false]|] + (map[0x371 := 850] + (map v0 : int | (-0x398 <= v0) && (v0 < 0x27c) :: (v0 - |"ff"|) := (0x60)))
}
function fm3(p0: bool, globalState: GlobalState): char {
	match DC11() {
		case DC11() => if (true) then 'h' else 'v'
		case DC10(cf18) => 't'
	}
}
function fm4(p0: int, p1: bool, globalState: GlobalState): bool {
	safeDivisionInt(|(seq(abs(-0x59), i0  => ('e')))|, -0x222) < safeDivisionInt(-556, 756)
}
function fm5(globalState: GlobalState): map<int, bool> {
	map[safeModuloInt(|(seq(abs(733), i0  => ('s')))|, |multiset{!false}|) := false]
}
function fm6(p0: bool, p1: bool, p2: bool, globalState: GlobalState): multiset<bool> {
	multiset{true, -765 < |(map v0 : int | v0 in [DC4(|map[false := |multiset{false}|]|, false, map[-523 := multiset{|"ylexru"|}]).cf9, |[|[true]|]|] :: (v0 + |{|"djl"|, |[false, !true, false]|}|) := (0x211))|, true || !true}
}
function fm7(p0: char, globalState: GlobalState): seq<bool> {
	[!(multiset{true} < multiset{false, true})]
}
function fm8(p0: bool, p1: int, globalState: GlobalState): D1 {
	DC3((map v0 : int | (-0x199 <= v0) && (v0 < 0x1d4) :: (v0 * -|"uicbdh"|) := (|multiset{0x2af, |(seq(abs(0x226), i0  => ('k')))|}|)) + map[|{true}| := -962], -|([-0xc4, 427] + [0xdf])|, if (false) then DC0(|"tlkawrtbe"|, false, seq(abs(0x7c), i1  => ('t')), map[0x2a3 := map[-0x114 := false]]) else DC0(0xcd, false, "xhtbrf", map[|multiset{|["bq"]|}| := map[-0x13e := true]]))
}
function fm9(globalState: GlobalState): seq<D1> {
	[DC3(map[0x327 := |map[false := true]|], |multiset{|"wdqdlcew"|, 266, 0x87}|, DC0(|"f"|, false, seq(abs(0x347), i0  => ('k')), map[0x2cd := map v0 : int | (0x16c <= v0) && (v0 < 0x353) :: (v0 - -|[DC0(-0x382, true, seq(abs(0x239), i1  => ('v')), map v1 : int | (0x26b <= v1) && (v1 < 0x395) :: (v1 - |["fcscl"]|) := (map v2 : int | (0x2c8 <= v2) && (v2 < 0xec) :: (safeDivisionInt(v2, |"qokpwk"|)) := (false)))]|) := (false)]))] + [DC3(map[-0x3c2 := |DC17(set v3 : int | (-634 <= v3) && (v3 < 0x9) :: (safeModuloInt(v3, |"bk"|))).cf29|], 0x296, DC0(0xa4, false, "newt", map[|map[-|"dxly"| := false]| := map[447 := true]])), DC3(map v4 : int | (620 <= v4) && (v4 < 0x1dd) :: (safeModuloInt(v4, 0x4d)) := (486), -0xd7, DC0(0x28b, false, "jhxxrjsnl", map[295 := map[-0x251 := true]])), DC3(map[414 := |multiset{false}|], |(seq(abs(356), i2  => ('x')))|, DC0(|map[!true := 0x363]|, true, "ouh", map[|"h"| := map[0xcf := false]]))]
}
function fm10(p0: bool, globalState: GlobalState): map<multiset<bool>, int> {
	map[multiset{true, false, true, false} := |(seq(abs(0x25b), i0  => ('u')))|] + (map[multiset{false} := -742] + DC20(map[multiset([!true, true]) := 0x394]).cf33)
}
function fm11(p0: int, p1: bool, globalState: GlobalState): map<int, map<int, bool>> {
	map[0x2f6 := map v0 : int | v0 in map[|[!true]| := true] :: (safeDivisionInt(v0, |map[0x26f := false]|)) := (true)] + (map[|"bre"| := map[-0x31e := true]] + map[|[!true]| := map[-0x2b1 := false]])
}
function fm12(p0: D0, p1: int, p2: char, p3: int, globalState: GlobalState): string {
	"mcyqkio" + ("egwe" + "sewnuy")
}
function fm13(globalState: GlobalState): D1 {
	DC4(0x31b, !!(|(map v0 : int | (0x165 <= v0) && (v0 < 0x238) :: (safeModuloInt(v0, 0x319)) := (-|(seq(abs(0xd), i0  => ('w')))|))| == 907), map[|(seq(abs(-463), i1  => ('k')))| := multiset{|map[true := |map[true := 0x157]|]|, -382, -|(set v1 : char | v1 in (seq(abs(0x245), i2  => ('r'))) :: (v1))|}] + map[0x202 := multiset{0x2d6}])
}
function fm14(globalState: GlobalState): seq<multiset<int>> {
	[multiset{-0x311}, multiset{-|"sxmprhotv"|}, multiset([|{false, true, false}|, 0x17a, --951, -|multiset{false}|])] + ((seq(abs(-648), i0  => (multiset([-950])))) + (seq(abs(0xf2), i1  => (multiset{-132, -|(seq(abs(914), i2  => ('w')))|, -0x60}))))
}
function fm15(p0: bool, p1: set<bool>, p2: bool, globalState: GlobalState): map<char, bool> {
	map['d' := !false && false]
}
function fm16(p0: char, p1: seq<int>, p2: map<D3, int>, globalState: GlobalState): D0 {
	DC0(-10, 0x2ab > |{0x13c}|, if (true) then "w" else "jpsbupx", map v0 : int | (0x29f <= v0) && (v0 < 265) :: (safeModuloInt(v0, -|(seq(abs(-0x1fa), i0  => ('l')))|)) := (map v1 : int | (0x35a <= v1) && (v1 < 123) :: (safeDivisionInt(v1, 0x19a)) := (!false)))
}
function fm17(p0: int, globalState: GlobalState): seq<int> {
	(seq(abs(0x282), i0  => (0x336))) + ([219, |"agbh"|, |"xvgfxoaqc"|, |{-0xd3, 0x9d, 330}|, 31] + [0x321, 0x390])
}
function fm18(p0: int, p1: int, p2: int, p3: char, globalState: GlobalState): set<string> {
	{"vvxhidyxj"} - {seq(abs(0x92), i0  => ('e'))}
}
function fm19(p0: multiset<bool>, globalState: GlobalState): map<bool, map<bool, int>> {
	(map[true := map[false := |map[false := false]|]] + map[true := map[true := 878]]) + (map[true := map[DC21(true, 'g', |"rjujka"|, |"lxlnkch"|).cf34 := 317]] + map[!false := map[true := |[|"xpd"|]|]])
}
function fm20(p0: bool, globalState: GlobalState): multiset<int> {
	multiset{safeModuloInt(0x3a3, 846), -0x254}
}
method m0(p0: array<multiset<int>>, p1: int, p2: bool, p3: bool, globalState: GlobalState) returns (r0: int, r1: array<int>, r2: set<bool>) {
	var v0 := true;
	var v1 := 'c';
	var v2: set<char> := {v1};
	v0 := p1 > safeDivisionInt(-0x2a5, |v2|);
	globalState.f6 := p1;
	r0 := (-0x1c5 + p1) + -(p1 * -p1);
	var v3 := "mhpactow";
	var v4 := DC0(p1, p3, v3, fm11(p1, v0, globalState));
	match (match v4.(cf2 := fm12(v4, p1, v1, p1, globalState), cf1 := p3) {
		case DC0(cf0, cf1, cf2, cf3) => DC2(v1)
	}) {
		case DC2(cf5) =>
			var v5: array<bool> := new bool[9];
			v0, v0, v3, v5, v3 := p2, v0, v3 + (v3 + "y"), v5, v3;
			var v6: map<int, int> := map[safeDivisionInt(fm0(p1, map[|v3| := |v3[safeIndex(p1, |v3|) := cf5]|], globalState), p1) := p1 - p1];
			var v7: set<string> := {v3, v3};
			v6 := v6[p1 := |v7|];
			v5[safeIndex(124, v5.Length)] := true;
			globalState.f11, v5[safeIndex(124, v5.Length)], globalState.f16 := p1, v0, p1;
			var v8: array<int> := new int[26](i0 => i0 + p1);
			v8[safeIndex(544, v8.Length)] := p1;
			v1, v8[safeIndex(544, v8.Length)], v5[safeIndex(124, v5.Length)] := v1, p1, v0;
		case DC3(cf6, cf7, cf8) =>
			var v9: C0 := new C0(v0);
			var v10: multiset<C0> := multiset{v9};
			var v11: map<C0, int> := map[v9 := if (v9 in v10) then v10[v9] else p1];
			var v12: set<map<C0, int>> := {v11, v11 + v11, v11, v11 + v11};
			v12 := v12;
			var v14: seq<int> := [cf7];
			var v15: seq<seq<bool>> := [[v9.f17]];
			var v16: map<int, bool> := map[cf7 := p3];
			var v17: set<int> := {0xbf};
			var v18: map<bool, set<int>> := map[if (0x382 in v16) then v16[0x382] else p2 := v17];
			var v19: set<int> := {cf7, -|(if (p2 in v18) then v18[p2] else v17)|, p1, cf7};
			var v20: array<int> := new int[28] [-p1, 763, p1 - fm0(p1, cf6, globalState), p1, 947, p1, cf7 - |(map v13 : int | v13 in cf6 :: (v13 * cf7) := (p1))|, |(v3[safeIndex(p1, |v3|) := v1] + v3)|, p1, if (v9 in v10) then v10[v9] else cf7, p1, cf7, cf7, cf7, cf7, p1, cf7, v14[safeIndex(p1, |v14|)], v14[safeIndex(cf7, |v14|)], cf7, |v15[safeIndex(p1, |v15|)]|, safeModuloInt(cf7, -0x374), |v19|, p1, cf7, cf7, p1, |v3|];
			v20[safeIndex(705, v20.Length)] := if (!p3) then 183 else p1;
			v20[safeIndex(705, v20.Length)], globalState.f16, v0, v0, v0 := fm0(p1, cf6, globalState), 0x38d, true, true, v0;
			var v21: multiset<int> := multiset{p1};
			var v22: set<bool> := {v9.f17, v9.f17};
			var v23: map<multiset<int>, set<bool>> := map[v21 - v21 := v22];
			v23 := v23;
			var v24: seq<array<int>> := [v20, v20, v20, v20];
			v14, v9, v24 := v14, v9, v24;
		case DC4(cf9, cf10, cf11) =>
			var v25: seq<bool> := [cf10, p2, p2, p3, p2];
			var v26 := DC1(v25);
			v26 := DC1([p2, p2]);
			match (DC2(v1)) {
				case DC2(cf5) =>
					var v27 := new C0(p3);
					var v28: seq<C0> := [v27];
					var v29: array<bool> := new bool[7];
					var v30: map<int, int> := map[cf9 := |map[v29 := cf5]|];
					var v31: map<bool, D1> := map[cf10 := DC4(cf9, fm4(|v28|, v27.f17, globalState), map[cf9 := multiset{cf9, fm0(|v3|, v30, globalState)}])];
					var v33: map<int, bool> := map[|v28| := v0];
					var v34 := DC4(cf9, (fm13(globalState)).cf10, map v32 : int | v32 in v33 :: (v32 + p1) := (multiset{cf9, p1}));
					cf10 := (v31 + v31) != map[v0 := v34];
					var v37: array<set<map<char, int>>> := new set<map<char, int>>[18](i1 => {map v35 : char | v35 in [v3[safeIndex(p1, |v3|)], v1] :: (v35) := (p1), map v36 : char | v36 in multiset([v1]) :: (v36) := (-p1)} + {map[cf5 := cf9]});
					var v39: multiset<char> := multiset{'k'};
					var v40: map<char, int> := map[cf5 := fm0(p1, v30, globalState)];
					var v41: multiset<map<char, int>> := multiset{map v38 : char | v38 in v39 :: (v38) := (|multiset{cf10}|), v40};
					var v43: set<map<char, int>> := {v40};
					v37[safeIndex(87, v37.Length)] := (set v42 : map<char, int> | v42 in v41 :: (v42)) - v43;
					v37[safeIndex(87, v37.Length)] := v43 * v43;
					v30 := v30 + v30;
				case DC3(cf6, cf7, cf8) =>
					var v44: set<int> := {|cf6|};
					var v45: array<char> := new char[16](i3 => v1);
					cf10 := |(seq(abs(0x25f), i2  => (|v3|)))| <= safeModuloInt(|v44|, |multiset{v45}|);
					var v46: array<bool> := new bool[8];
					var v47: C0 := new C0(cf10);
					var v48: multiset<bool> := multiset{p2};
					var v49: map<int, bool> := map[p1 := v47.f17];
					var v50: map<int, map<int, bool>> := map[cf9 := v49];
					var v51: map<C0, bool> := map[v47 := fm4(cf9, DC0(if (fm4(-968, p3, globalState) in v48) then v48[fm4(-968, p3, globalState)] else cf9, fm4(cf9, !v0, globalState), "suaxnms", v50).cf1, globalState)];
					v46[safeIndex(780, v46.Length)] := if (v47 in v51) then v51[v47] else v0;
					v46[safeIndex(780, v46.Length)] := v47.f17;
					v46[safeIndex(780, v46.Length)] := (set v52 : int | (0x1fd <= v52) && (v52 < 0x3d0) :: (v52 + cf9)) >= v44;
					var v53: seq<array<bool>> := [v46, v46, v46, v46, v46];
					v53 := v53;
				case DC4(cf9, cf10, cf11) =>
					v1 := v1;
					var v54: multiset<bool> := multiset{v0};
					globalState.f6 := if (v0 in v54) then v54[v0] else |[!p3]| + cf9;
					var v55: array<int> := new int[14];
					var v56: map<bool, array<int>> := map[cf10 := v55];
					var v57: seq<int> := [cf9];
					var v58: array<bool> := new bool[3] [p3, p3, cf10];
					var v59: map<int, array<bool>> := map[p1 + v57[safeIndex(p1, |v57|)] := v58];
					var v60: seq<map<bool, array<int>>> := [v56, v56 + v56];
					var v61: map<string, int> := map[seq(abs(0x337), i4  => (v1)) := p1];
					var v62: set<int> := {cf9};
					var v63: C0 := new C0(fm4(|{|map[v0 := p3]|, cf9}|, p3, globalState));
					var v64: multiset<C0> := multiset{v63, v63, v63};
					var v65: map<int, int> := map[|v62| := -|v64|];
					var v66: map<bool, map<int, int>> := map[p2 := v65];
					globalState.f11, v56, globalState.f11, v59 := --safeModuloInt(cf9, cf9), v60[safeIndex(cf9 - p1, |v60|)], fm0(p1 * |v61|, if (true in v66) then v66[true] else v65, globalState), v59;
					var v67: seq<string> := [seq(abs(0x3d6), i5  => ('e')), (seq(abs(-0x12c), i6  => (v1))) + v3];
					v3 := (v67[safeIndex(safeModuloInt(p1, p1), |v67|)])[safeIndex(cf9, |v67[safeIndex(safeModuloInt(p1, p1), |v67|)]|) := 'j'];
				case DC1(cf4) =>
					globalState.f6 := cf9;
					var v69: C0 := new C0(!cf10);
					var v70: map<C0, int> := map[v69 := cf9];
					globalState.f6 := |(map v68 : char | v68 in v3 :: (v68) := (0x35c))| - (|v70| + p1);
					v69 := v69;
					var v71 := DC2(fm3(cf10, globalState));
					v71 := v71.(cf5 := 'q');
				case DC5(cf12) =>
					var v72 := DC4(p1, true, cf11);
					var v73: map<bool, seq<int>> := map[v72.cf10 := [cf9]];
					var v74: array<int> := new int[4];
					var v75: C0 := new C0(p2);
					var v76: array<C0> := new C0[8] [v75, v75, v75, v75, v75, v75, v75, v75];
					var v77: map<array<int>, array<C0>> := map[v74 := v76];
					var v78: set<int> := {|multiset{cf9}|};
					var v79: map<int, int> := map[cf9 := |v78|];
					r0, globalState.f16, v73, globalState.f6, v1 := cf9, |(v77 + v77)|, v73 + v73, safeDivisionInt(safeDivisionInt(p1, cf9), |v79|), v1;
					var v80: array<bool> := new bool[28];
					v74[safeIndex(77, v74.Length)] := if (p1 in v79) then v79[p1] else v4.cf0;
					var v81: seq<array<bool>> := [v80, v80];
					v80, cf9, v74[safeIndex(77, v74.Length)], globalState.f6 := v80, safeDivisionInt(cf9, p1), |v81|, safeDivisionInt(0x102, p1);
					var v82: multiset<bool> := multiset{p3, if (v75.f17) then cf10 else false, v0};
					v82 := v82;
					v76[safeIndex(72, v76.Length)] := v75;
					var v83: array<array<bool>> := new array<bool>[18];
					v83[safeIndex(278, v83.Length)] := v81[safeIndex(-0x175, |v81|)];
					v80[safeIndex(562, v80.Length)] := p3;
					v76[safeIndex(72, v76.Length)], cf10, v83[safeIndex(278, v83.Length)], v80[safeIndex(562, v80.Length)], v0 := v75, safeModuloInt(p1, fm0(v74[safeIndex(77, v74.Length)], v79, globalState)) > v74[safeIndex(77, v74.Length)], v80, p2, p2;
			}
			
			var v84: array<map<int, int>> := new map<int, int>[9];
			v84 := v84;
			var v85: map<int, int> := map[p1 := cf9];
			var v86: multiset<bool> := multiset{v25[safeIndex(p1, |v25|)]};
			match (DC3(v85, p1, v4).(cf7 := |v86|)) {
				case DC2(cf5) =>
					cf10, v3 := cf10 || (!v0 || v0), "pufjq";
					var v87: array<set<array<int>>> := new set<array<int>>[5];
					var v88: array<int> := new int[9];
					v87[safeIndex(634, v87.Length)] := {v88, v88, v88};
					var v89: set<array<int>> := {v88, v88, v88, v88};
					var v90: seq<set<array<int>>> := [v89, v89, {v88, v88}, v89];
					v87[safeIndex(634, v87.Length)] := (v89 * v90[safeIndex(p1, |v90|)]) * v89;
					cf10 := (multiset(v25))[p3 := abs(cf9)] > v86;
					var v91 := new C0(p3);
				case DC3(cf6, cf7, cf8) =>
					v1 := v1;
					var v92: array<int> := new int[19];
					var v93: map<int, string> := map[p1 := v3[safeIndex(7, |v3|) := 't']];
					var v94: map<bool, int> := map[v0 := p1];
					v92[safeIndex(57, v92.Length)] := |(if ((if (cf10 in v94) then v94[cf10] else cf7) in v93) then v93[if (cf10 in v94) then v94[cf10] else cf7] else seq(abs(587), i7  => (v1)))|;
					var v95: array<bool> := new bool[18] [cf10, cf10, p2, p2, p3, p2, -0x1f4 <= p1, cf10, !v25[safeIndex(cf7, |v25|)], v0, (fm6(p2, p3, p2, globalState))[p2 := abs(cf9)] !! v86, p3, p3, false, p1 == cf7, p2, p2, true];
					v95[safeIndex(156, v95.Length)] := v0;
					v0, cf10, v92[safeIndex(57, v92.Length)], v95[safeIndex(156, v95.Length)], globalState.f16 := |("udu" + v3)[safeIndex(p1, |("udu" + v3)|) := v1]| >= 0x159, p3, safeModuloInt(cf7, cf9), cf10, cf7 - p1;
					var v96 := DC6(v92);
					var v97: seq<array<int>> := [v96.cf13, v92, v92];
					var v98: array<C0> := new C0[19];
					var v99: C0 := new C0(cf10);
					v98[safeIndex(687, v98.Length)] := v99;
					v95[safeIndex(156, v95.Length)], v97, v98[safeIndex(687, v98.Length)], v95 := v25[safeIndex(cf9, |v25|)], (if (cf9 > p1) then v97 + v97 else v97)[safeIndex(v92[safeIndex(57, v92.Length)], |(if (cf9 > p1) then v97 + v97 else v97)|) := v92], v99, v95;
					v95[safeIndex(156, v95.Length)] := cf9 <= p1;
				case DC4(cf9, cf10, cf11) =>
					var v100 := new C0(cf10);
					var v101: map<bool, bool> := map[v0 := cf10];
					globalState.f6 := |v101|;
					var v102: array<bool> := new bool[2](i8 => p3);
					v102[safeIndex(678, v102.Length)] := p2;
					v102[safeIndex(678, v102.Length)] := if (!(v3 <= v3) in v101) then v101[!(v3 <= v3)] else cf10;
					globalState.f16 := |v3|;
				case DC1(cf4) =>
					var v103 := new C0(p2 ==> !!p2);
					var v104: set<bool> := {v103.f17};
					var v105: map<int, bool> := map[cf9 := v25[safeIndex(|map[v103.f17 := v103.f17]|, |v25|)]];
					v0 := fm4(safeDivisionInt(p1, |v104|), if (cf9 in v105) then v105[cf9] else p2, globalState);
					var v106: array<bool> := new bool[2];
					v106 := v106;
					var v107: seq<C0> := [v103];
					var v108: seq<int> := [p1, p1, cf9, p1, p1];
					v103, v103, v0, v103 := v103, v103, !(v107 != v107), v107[safeIndex(p1 * v108[safeIndex(0x171, |v108|)], |v107|)];
				case DC5(cf12) =>
					globalState.f6 := p1;
					var v109: C0 := new C0(p3);
					var v110: map<int, C0> := map[cf9 := v109];
					v110 := v110[-cf9 := v109];
					var v111 := new C0(p2);
					globalState.f11, globalState.f6 := -0x1c5, (-p1 - p1) * |fm14(globalState)|;
			}
			
		case DC1(cf4) =>
			var v112: C0 := new C0(p2);
			v112 := v112;
			var v113: map<string, C0> := map[v3 := v112];
			v112 := if ((if (!!p2) then v3 else v3) in v113) then v113[if (!!p2) then v3 else v3] else if (false) then v112 else v112;
			r0 := p1;
			var v114: map<bool, int> := map[v112.f17 := p1];
			var v115: seq<C0> := [v112];
			var v116: map<int, int> := map[p1 := p1];
			var v117: map<int, int> := map[fm0(|v115|, v116, globalState) := p1];
			var v118: map<int, bool> := map[p1 := v0];
			var v119 := DC4(safeDivisionInt(p1, p1), cf4[safeIndex(p1, |cf4|)], map[p1 := multiset{if (cf4[safeIndex(|[446, if (p1 in v116) then v116[p1] else p1]|, |cf4|)] in v114) then v114[cf4[safeIndex(|[446, if (p1 in v116) then v116[p1] else p1]|, |cf4|)]] else |v118|}]);
			match (v119) {
				case DC2(cf5) =>
					var v120: set<seq<char>> := {seq(abs(0x7c), i9  => (cf5))};
					var v122: set<bool> := {!v112.f17, false, p3, v112.f17, (fm13(globalState)).cf10};
					var v124: seq<map<int, bool>> := [v118, map v123 : int | (92 <= v123) && (v123 < 86) :: (v123 * 2) := (p2)];
					var v125: array<bool> := new bool[15] [p2, p3, !p3, v112.f17, !v0, v120 > (set v121 : seq<char> | v121 in v120 :: (v121)), v112.f17, p3, |v122| > p1, p2, p2, p3, p3, v112.f17, map[p1 := true] in v124];
					var v126: multiset<int> := multiset{p1};
					var v127: map<int, multiset<int>> := map[p1 := v126];
					v125[safeIndex(328, v125.Length)] := if (v112.f17) then DC4(p1, p2, v127).cf10 else v112.f17;
					var v128: set<int> := {p1};
					v125[safeIndex(328, v125.Length)] := (v128 > v128) && p3;
					var v129 := new C0(p1 > p1);
					var v131: array<int> := new int[27] [p1, p1, p1, p1, p1, p1, p1, p1, p1, p1, -p1, |(map v130 : int | v130 in v118 :: (v130 + p1) := (v112.f17))|, p1, |v3|, p1, p1, p1, p1, |v126|, -p1, |v3|, p1, p1, p1, 0x3a2, -p1, p1];
					var v132: seq<array<int>> := [v131, v131];
					var v133: array<seq<array<int>>> := new seq<array<int>>[12] [v132, v132, [v131], v132, v132, v132, [v131] + v132, v132, v132 + v132, if (p2) then [v131] else v132, v132 + v132, v132];
					v133[safeIndex(807, v133.Length)] := v132;
					v133[safeIndex(807, v133.Length)], globalState.f11, v0, globalState.f11, v131 := v132, (|v122| + p1) * p1, !!(p2 || v112.f17), p1, v131;
					v0 := v125[safeIndex(328, v125.Length)];
				case DC3(cf6, cf7, cf8) =>
					globalState.f16 := if (v0) then p1 else cf7;
					var v134: map<char, bool> := map[v1 := !p3];
					var v135: set<bool> := {v112.f17, true};
					v134 := fm15(v0, v135, !p2, globalState);
					var v136 := DC10(v115);
					v115 := v136.cf18;
					v116 := v116[cf7 := cf7 * fm0(p1, cf6, globalState)];
				case DC4(cf9, cf10, cf11) =>
					var v137: array<seq<int>> := new seq<int>[4];
					var v138: set<C0> := {v112, v112};
					var v139: array<int> := new int[4] [-0x8e, -0x272, p1, |v138|];
					var v140: multiset<array<int>> := multiset{v139};
					var v141: seq<int> := [fm0(|v140|, v116, globalState)];
					v137[safeIndex(633, v137.Length)] := v141;
					var v142 := DC2(v1);
					var v143: seq<D1> := [v142, v142, v142];
					var v145: multiset<map<int, int>> := multiset{map v144 : int | (998 <= v144) && (v144 < 785) :: (v144 - p1) := (cf9), v116, fm1(v116, cf9, cf9, globalState), v116};
					var v146: seq<seq<int>> := [seq(abs(541), i10  => (cf9)), v141, [p1, |[0x1c6, |v143|, p1, -|v3|]|, p1], v141, ([if (v116 in v145) then v145[v116] else cf9])[safeIndex(cf9, |[if (v116 in v145) then v145[v116] else cf9]|) := p1]];
					v137[safeIndex(633, v137.Length)] := v146[safeIndex(p1, |v146|)];
					var v147 := new C0(cf10);
					var v148 := DC12(v114);
					v114 := v114 + v148.cf19;
					var v149: array<bool> := new bool[28];
					v149[safeIndex(609, v149.Length)] := fm4(fm0(|"ns"|, v116, globalState), p2, globalState);
					v149[safeIndex(609, v149.Length)] := cf10;
				case DC1(cf4) =>
					v116 := v116;
					var v150: multiset<int> := multiset{|"hhlfxth"|};
					v150 := v150;
					var v151: seq<int> := [p1, p1];
					var v152: set<int> := {p1, fm0(|[p2, p3]|, v116, globalState), p1, fm0(v151[safeIndex(p1, |v151|)], map[p1 := p1], globalState)};
					var v153: map<int, set<int>> := map[p1 := v152];
					var v154: set<set<int>> := {if (p1 in v153) then v153[p1] else {p1, p1, p1}, v152, v152, {|{v112.f17}|}};
					v0 := v152 in v154;
					v0, v0 := false, p2;
				case DC5(cf12) =>
					var v155: map<int, map<int, bool>> := map[p1 := v118];
					var v156: array<D0> := new D0[4] [v4.(cf3 := v155, cf0 := p1, cf2 := v3), v4, v4, v4];
					v156[safeIndex(210, v156.Length)] := v4;
					var v158: map<D3, bool> := map[DC11() := v112.f17];
					v156[safeIndex(210, v156.Length)] := fm16(v1, fm17(0x255, globalState), map v157 : D3 | v157 in v158 :: (v157) := (p1), globalState).(cf0 := p1, cf3 := fm11(|cf4|, p2, globalState));
					var v159 := new C0(v112.f17);
					var v160: array<int> := new int[4];
					var v161: seq<array<int>> := [if (p3) then v160 else v160, v160, v160, v160];
					v161 := [v160] + v161;
					v0 := v0;
			}
			
		case DC5(cf12) =>
			if (v0) {
				var v162: multiset<bool> := multiset{p2, !!p3};
				var v163: C0 := new C0(p2);
				var v164: map<multiset<bool>, C0> := map[v162 * v162[v0 := abs(p1)] := v163];
				v164 := v164[v162 := v163];
				v0 := p3;
				var v165: array<C0> := new C0[19];
				v165[safeIndex(738, v165.Length)] := v163;
				v165[safeIndex(738, v165.Length)] := v163;
				var v166: map<bool, bool> := map[v0 := true];
				v166 := if (v163.f17) then v166[v163.f17 := v163.f17] else v166 + v166;
				globalState.f11 := p1;
			} else {
				globalState.f6, v0 := p1, !p2;
				var v167 := new C0(p3);
				var v168: array<bool> := new bool[28](i11 => true);
				v168[safeIndex(524, v168.Length)] := if (v167.f17) then v167.f17 else v0;
				v168[safeIndex(524, v168.Length)] := p2;
				v0 := fm4(p1, v168[safeIndex(524, v168.Length)], globalState);
				var v169: map<bool, int> := map[v0 := p1];
				var v170: set<int> := {|[v168[safeIndex(524, v168.Length)], v168[safeIndex(524, v168.Length)]]|, if (p2 in v169) then v169[p2] else p1, -p1};
				v168[safeIndex(524, v168.Length)] := !({p1} !! v170);
			}
			
			var v171: multiset<bool> := multiset{v0, true};
			if (fm4(p1, |v171| != p1, globalState)) {
				var v172: array<string> := new string[6];
				v172[safeIndex(131, v172.Length)] := "kinwdfjx" + v3;
				v172[safeIndex(131, v172.Length)] := seq(abs(0x1d4), i12  => ('j'));
				var v173 := new C0(v171 < v171);
				var v174: map<int, int> := map[522 := p1];
				var v175: map<int, bool> := map[p1 := true];
				globalState.f16 := fm0(p1, v174, globalState) + |(if (p3) then v175 else fm5(globalState))|;
				v172[safeIndex(131, v172.Length)] := v172[safeIndex(131, v172.Length)];
				var v176: array<multiset<bool>> := new multiset<bool>[12];
				v176[safeIndex(489, v176.Length)] := v171[v0 := abs(p1)];
				var v177: array<D4> := new D4[20];
				var v178: array<char> := new char[7](i13 => v1);
				var v179: set<array<char>> := {v178};
				var v180 := DC15(v179, v1);
				v177[safeIndex(875, v177.Length)] := v180;
				var v181: array<bool> := new bool[27];
				v176[safeIndex(489, v176.Length)], v177[safeIndex(875, v177.Length)], v0 := if (p2) then fm6(p2, p3, p2, globalState) else v171, DC15(v179, v1), DC13(fm4(p1, p3, globalState), p1, v181).cf20;
			} else {
				var v182: seq<bool> := [!p2];
				v0 := v182[safeIndex(---0x358, |v182|)];
				var v183: map<int, int> := map[p1 := p1];
				var v184: C0 := new C0(p2);
				var v185: map<int, C0> := map[p1 := v184];
				var v186: map<int, map<int, C0>> := map[fm0(p1, v183, globalState) := v185 + v185];
				var v187: seq<map<int, map<int, C0>>> := [v186];
				v186 := v186 + v187[safeIndex(|map[!false := true]|, |v187|)];
				var v188: map<bool, char> := map[v184.f17 := v1];
				var v189: set<string> := {v3};
				var v190 := DC14(|v188|, |v183| - p1, v189);
				v190 := v190;
				var v191: array<bool> := new bool[26];
				v191 := v191;
				v0 := !(p3 ==> v0);
			}
			
			var v192: array<D4> := new D4[1](i14 => DC14(|v171|, 0x3e4, {v3, v3}));
			var v193: C0 := new C0(!p3);
			var v194 := DC14(|v3|, |[v193]|, fm18(p1, p1, p1, v1, globalState));
			var v195: set<string> := {"fpagmi"};
			v192[safeIndex(13, v192.Length)] := v194.(cf25 := v195);
			v192[safeIndex(13, v192.Length)] := DC14(p1, p1, {v3});
			var v196: seq<bool> := [v0];
			var v197: map<bool, int> := map[p3 := |v196|];
			var v198: map<bool, map<bool, int>> := map[v0 := v197];
			v198 := fm19(multiset{v0, p3, v193.f17, p2} - v171, globalState);
	}
	
	var v199 := new C0(p2);
	var i15 := 0;
	while (p3)
		decreases 100 - i15
	{
		if (i15 >= 100) {
			break;
		}
		
		i15 := i15 + 1;
		var v200: array<array<int>> := new array<int>[25];
		var v201: map<array<array<int>>, C0> := map[v200 := v199];
		v201 := v201[v200 := v199];
		globalState.f16 := p1;
		if (v199.f17) {
			v0 := p3;
			v1 := v1;
			var v202: array<set<bool>> := new set<bool>[1];
			v202[safeIndex(172, v202.Length)] := {p3, fm4(p1, v199.f17, globalState)};
			var v203: set<bool> := {!v199.f17};
			v202[safeIndex(172, v202.Length)] := v203;
			v199 := new C0(v199.f17);
			var v204: array<C0> := new C0[13];
			v204[safeIndex(107, v204.Length)] := v199;
			v0, v3, globalState.f16, v204[safeIndex(107, v204.Length)] := v0, v3, |v3|, if (v199.f17) then v199 else v199;
		} else {
			var v205 := DC7(p1, v1);
			v205 := v205;
			var v206: array<int> := new int[10];
			v200[safeIndex(610, v200.Length)] := v206;
			v200[safeIndex(610, v200.Length)] := v206;
			var v207: map<C0, char> := map[v199 := v1];
			v0 := fm4(-p1, v199 !in v207[v199 := v1], globalState);
			v0 := if (p1 == |fm20(p3, globalState)|) then v199.f17 else p2;
			var v208: multiset<int> := multiset{p1};
			p0[safeIndex(113, p0.Length)] := v208;
			v200[safeIndex(610, v200.Length)][safeIndex(887, v200[safeIndex(610, v200.Length)].Length)] := p1;
			var v209: array<bool> := new bool[1] [v199.f17];
			var v210: seq<bool> := [v199.f17, false, v199.f17];
			var v211: set<int> := {|v210|, p1};
			v209[safeIndex(649, v209.Length)] := {p1} < v211;
			var v212: seq<multiset<int>> := [multiset{p1} - v208, v208, v208, v208 - multiset{p1, p1}];
			var v213: map<int, C0> := map[p1 := v199];
			var v214: map<int, map<int, C0>> := map[p1 := v213];
			var v215: set<bool> := {false};
			var v216: map<int, int> := map[p1 := p1];
			v0, p0[safeIndex(113, p0.Length)], v200[safeIndex(610, v200.Length)][safeIndex(887, v200[safeIndex(610, v200.Length)].Length)], globalState.f6, v209[safeIndex(649, v209.Length)] := !!(v0 || v199.f17), v212[safeIndex(p1, |v212|)], p1, if (|v214| in v208) then v208[|v214|] else fm0(|v215|, v216, globalState), false;
		}
		
		var v217: array<int> := new int[12];
		var v218: seq<int> := [p1];
		var v219: seq<seq<int>> := [v218];
		var v220: map<bool, int> := map[v0 := p1];
		v217[safeIndex(549, v217.Length)] := |(v219[safeIndex(if (v199.f17 in v220) then v220[v199.f17] else p1, |v219|)] + [p1])|;
		v217[safeIndex(549, v217.Length)] := p1 + (p1 + p1);
	}
	var v221: map<int, int> := map[p1 := 0x19f];
	var v222: seq<bool> := [p3];
	r0 := (fm0(p1, v221, globalState) + 0x312) - |(v222 + [p3])|;
	var v223 := DC2(v1);
	var v224: map<D1, bool> := map[v223 := v0];
	var v225: multiset<bool> := multiset{v222[safeIndex(p1, |v222|)]};
	var v226: map<bool, int> := map[p3 := p1];
	var v227: array<int> := new int[27] [|[v0, v199.f17, v199.f17, p2, v0]|, p1 + -p1, if (p3) then --p1 else p1, p1, p1, safeDivisionInt(290, p1), p1, 0x2f0, if (p2) then p1 else |v224|, p1, -|v3|, p1, 0x40 + |(seq(abs(0x228), i16  => (v1)))[safeIndex(p1, |(seq(abs(0x228), i16  => (v1)))|) := v1]|, safeDivisionInt(|v225|, |map[if (p2 in v226) then v226[p2] else |v221| := v199]|), p1, p1, p1, p1, p1 + -p1, p1, p1, p1, p1, |(seq(abs(0x1d6), i17  => (p1)))|, p1, p1, p1];
	r1 := v227;
	var v228: set<bool> := {if (v0) then v199.f17 else v0, v199.f17, p3};
	r2 := v228;
}
method Main() {
	var v1 := 'd';
	var v2 := 514;
	var v4 := false;
	var v5: seq<bool> := [v4, true];
	var v6: map<int, bool> := map[|v5| := true];
	var v7: set<bool> := {if (v2 in v6) then v6[v2] else !v4, v4, v4, v4};
	var v8 := "frrp";
	var v9: array<int> := new int[7](i0 => safeDivisionInt(i0, 0x217));
	var globalState := new GlobalState(0x1c7, 0x77, (map v0 : char | v0 in [v1] :: (v0) := (0x44))[v1 := v2], true, 0x3a8, 0x115, 911, multiset{v2, -|(set v3 : int | v3 in [v2] :: (v3 * --0xe9))|, |v7|, |v8|}, true, 828, 0x26e, -572, 0x106, false, v9, 0x2d6, 237);
	var v10: map<int, int> := map[v2 := v2];
	for i1 := fm0(v2, v10[-923 := v2], globalState) to v2 {
		globalState.f16 := i1;
		var v11: multiset<bool> := multiset{!v4, v4, v4};
		var v12: multiset<multiset<bool>> := multiset{v11};
		var v13: seq<int> := [if (v11 in v12) then v12[v11] else i1];
		v13 := v13;
		var v14: map<bool, map<int, bool>> := map[v4 := v6 + v6];
		v14 := v14;
		var v16 := DC0(i1, v4, "yicqicecd", map v15 : int | (0x18c <= v15) && (v15 < -0x9c) :: (v15 - |v8|) := (v6));
		match (v16) {
			case DC0(cf0, cf1, cf2, cf3) =>
				v9[safeIndex(540, v9.Length)] := ---(fm0(cf0, v10, globalState) * i1);
				v9[safeIndex(540, v9.Length)] := v2;
				v10 := v10[-v9[safeIndex(540, v9.Length)] := v9[safeIndex(540, v9.Length)]];
				var v17 := DC1(v5);
				var v18: multiset<int> := multiset{cf0};
				var v19: map<bool, bool> := map[cf1 := v4];
				var v20: map<int, seq<bool>> := map[-v9[safeIndex(540, v9.Length)] := v5];
				var v21: array<seq<bool>> := new seq<bool>[25] [v5, v5, v17.cf4 + v5, v5, v5, v5, v5, v5[safeIndex(-cf0, |v5|) := cf1], (v5[safeIndex(|v18|, |v5|) := v4])[safeIndex(i1, |v5[safeIndex(|v18|, |v5|) := v4]|) := cf1] + [false, if (cf1 in v19) then v19[cf1] else v4], v5, v5, v5, v5, v5, if (-0x173 in v20) then v20[-0x173] else [cf1], v5, v5 + v5, [false], v5 + v5, v5, v5 + DC1(v5).cf4, v5, v5, v5, v5];
				v21[safeIndex(128, v21.Length)] := [v4, false] + v5;
				var v22: multiset<string> := multiset{v8};
				var v23: array<D0> := new D0[28](i2 => DC0(i1, cf1, v8, cf3));
				var v24: seq<set<bool>> := [{v4}];
				var v25: array<int> := new int[17];
				var v26: set<array<int>> := {v25};
				v21[safeIndex(128, v21.Length)], v22, v23, v24, globalState.f6 := v5 + v5, multiset{"jaeue"}, v23, seq(abs(203), i3  => (v7)), fm0(v2, fm1(v10, |v10|, |v26|, globalState) + v10, globalState);
				var v27: multiset<array<int>> := multiset{v25};
				cf2 := (("flvuwul")[safeIndex(cf0, |"flvuwul"|) := if (v4) then v1 else v8[safeIndex(cf0, |v8|)]])[safeIndex(safeModuloInt(240, |v27|), |("flvuwul")[safeIndex(cf0, |"flvuwul"|) := if (v4) then v1 else v8[safeIndex(cf0, |v8|)]]|) := v1];
		}
		
	}
	v9[safeIndex(12, v9.Length)] := v2;
	v9[safeIndex(12, v9.Length)] := safeModuloInt(v2, v2);
	var v28 := new C0(v4);
	var v29: multiset<int> := multiset{v2};
	var v30: map<int, multiset<int>> := map[v9[safeIndex(12, v9.Length)] := multiset{76}];
	var v31 := DC4(v2, v28.f17, v30);
	var v32 := DC5(v31);
	v1, v4, v4, v1, v1 := fm3(v4, globalState), v4 ==> (v29 < multiset{0xbb, |v10|}), match v32 {
		case DC2(cf5) => true
		case DC3(cf6, cf7, cf8) => v4
		case DC4(cf9, cf10, cf11) => v28.f17
		case DC1(cf4) => v4
		case DC5(cf12) => v28.f17 ==> v4
	}, if (v4) then v1 else v1, v1;
	var v33: multiset<bool> := multiset{false};
	var v34: array<bool> := new bool[6];
	v34[safeIndex(701, v34.Length)] := v4;
	v34[safeIndex(168, v34.Length)] := v28.f17;
	var v35: map<bool, int> := map[v28.f17 := v9[safeIndex(12, v9.Length)]];
	var v36: seq<map<int, int>> := [v10, v10];
	v33, v4, v34[safeIndex(701, v34.Length)], v34[safeIndex(168, v34.Length)] := v33, fm4(v9[safeIndex(12, v9.Length)] + v9[safeIndex(12, v9.Length)], v9[safeIndex(12, v9.Length)] >= v9[safeIndex(12, v9.Length)], globalState), fm4(|v8[safeIndex(if (fm4(v9[safeIndex(12, v9.Length)], v28.f17, globalState) in v35) then v35[fm4(v9[safeIndex(12, v9.Length)], v28.f17, globalState)] else fm0(0x10, v36[safeIndex(v9[safeIndex(12, v9.Length)], |v36|)], globalState), |v8|) := v1]|, v28.f17, globalState), fm4(v2, v28.f17, globalState);
	var v37 := new C0(v34[safeIndex(701, v34.Length)]);
	var v38: map<int, string> := map[v9[safeIndex(12, v9.Length)] := v8];
	v38 := v38[|v5| := v8];
	v34[safeIndex(701, v34.Length)] := (if (false) then v8 else v8[safeIndex(v2, |v8|) := v1]) <= v8;
	forall i4 | 0 <= i4 < v9.Length {
		v9[i4] := i4 * 0x29f;
	}
	var v39: array<seq<int>> := new seq<int>[14];
	forall i5 | 0 <= i5 < v39.Length {
		v39[i5] := [v9[safeIndex(12, v9.Length)]] + ([v2] + [0x81]);
	}
	globalState.f16 := v2;
	var i6 := 0;
	while (v37.f17)
		decreases 100 - i6
	{
		if (i6 >= 100) {
			break;
		}
		
		i6 := i6 + 1;
		var v40: map<D1, array<int>> := map[DC4(0x35d, v34[safeIndex(701, v34.Length)], v30) := v9];
		var v41 := DC4(|"ksjsyh"|, v28.f17, map[v9[safeIndex(12, v9.Length)] := v29]);
		var v42: multiset<array<int>> := multiset{v9, v9, if (v41 in v40) then v40[v41] else v9, v9, v9};
		v42 := v42;
		var v43: seq<array<bool>> := [v34, v34, v34];
		v34[safeIndex(701, v34.Length)] := !(v43 <= v43);
		var v44: array<string> := new string[8] ["fqsutwajj", v8, v8, v8, v8, v8, v8[safeIndex(|map[v34[safeIndex(701, v34.Length)] := v2]|, |v8|) := v1], "ibgqhjni"];
		var v45: map<bool, array<string>> := map[v28.f17 := v44];
		v45 := v45[true := v44];
		var v47 := new C0(v6 != (map v46 : int | v46 in fm5(globalState) :: (v46 - v2) := (v34[safeIndex(701, v34.Length)])));
	}
	v34[safeIndex(701, v34.Length)] := v37.f17;
	for i7 := 0x55 to v9[safeIndex(12, v9.Length)] {
		var v48 := new C0(v37.f17);
		var v49: seq<int> := [-|(seq(abs(0x36a), i8  => (v8)))|];
		v39[safeIndex(569, v39.Length)] := v49;
		v39[safeIndex(569, v39.Length)] := v49;
		var v50: map<int, map<int, bool>> := map[v2 := v6];
		var v51 := DC0(863, v37.f17, v8, v50);
		var v52 := DC3(v10, v2, v51);
		var v53: set<D1> := {v52, DC3(v10, i7, v51), v52};
		v9[safeIndex(12, v9.Length)] := safeModuloInt(i7, |(v53 * {v52})|);
		v35 := v35[v34[safeIndex(701, v34.Length)] := v9[safeIndex(12, v9.Length)]];
	}
	if (v8 <= v8) {
		v2 := v9[safeIndex(12, v9.Length)];
		var v54 := DC2(v1);
		var v55 := DC4(v2, v34[safeIndex(701, v34.Length)], v30);
		var v56: map<C0, multiset<bool>> := map[v28 := v33];
		var v58: set<int> := {|(map v57 : int | (-0x31a <= v57) && (v57 < 807) :: (safeDivisionInt(v57, v2)) := (v1))|};
		var v59: array<multiset<bool>> := new multiset<bool>[20] [fm6(false, v4, v28.f17, globalState), multiset{v55.cf10, v28.f17, false, v28.f17, true}, v33, v33, multiset(fm7(v54.cf5, globalState) + v5), v33, fm6(true, v34[safeIndex(701, v34.Length)], v37.f17, globalState), v33, v33, v33, v33, v33, v33, v33[v37.f17 := abs(v2)], if (v37 in v56) then v56[v37] else v33[v28.f17 := abs(|v58|)], multiset{true, v37.f17}, v33, multiset{v34[safeIndex(701, v34.Length)], v28.f17}, v33, v33];
		v59[safeIndex(406, v59.Length)] := v33;
		v54, v59[safeIndex(406, v59.Length)] := v54, multiset{v37.f17} + v33;
		var v60 := new C0(if (v4) then !v37.f17 else v37.f17);
		var v61: set<seq<bool>> := {v5};
		if (v61 > (v61 * v61)) {
			var v62: array<char> := new char[25](i9 => 'o');
			v62, v4 := v62, v60.f17;
			var v63: array<C0> := new C0[15];
			var v64: set<array<C0>> := {v63};
			v9, v34[safeIndex(701, v34.Length)], v4, v28 := v9, {v63} > v64, fm4(v9[safeIndex(12, v9.Length)], v37.f17, globalState), v28;
			var v65: map<int, map<int, bool>> := map[v2 := v6];
			var v66 := DC0(v2, v60.f17, v8, v65);
			var v67 := DC3(v10, v9[safeIndex(12, v9.Length)], v66);
			v67 := fm8(v28.f17, v9[safeIndex(12, v9.Length)], globalState);
			var v68: map<bool, string> := map[v37.f17 := "fyk"];
			v68 := v68[v37.f17 := v8];
			var v69: seq<D1> := [DC3(v10[v2 := v9[safeIndex(12, v9.Length)]], v9[safeIndex(12, v9.Length)], v66), v67, DC3(map[|[v58, v58]| := 214], v2, DC0(|v7|, v4, seq(abs(0x24e), i10  => (v1)), v65))];
			var v70: multiset<seq<D1>> := multiset{[v67, v67], v69, v69, v69, fm9(globalState)};
			v34[safeIndex(701, v34.Length)] := v70 > v70;
		} else {
			var v71: seq<int> := [safeDivisionInt(v9[safeIndex(12, v9.Length)], v2), v2];
			v9[safeIndex(12, v9.Length)] := -v71[safeIndex(v9[safeIndex(12, v9.Length)] + v9[safeIndex(12, v9.Length)], |v71|)];
			var v72 := new C0(v7 !! v7);
			var v73: map<int, map<int, bool>> := map[-0x179 := v6];
			var v74 := DC0(|v8|, v37.f17, v8, v73);
			var v75 := DC3(v10, v9[safeIndex(12, v9.Length)] * v2, v74);
			v75 := v75.(cf6 := map v76 : int | v76 in v71 :: (v76 * v9[safeIndex(12, v9.Length)]) := (|v10|));
			var v77: map<bool, bool> := map[v28.f17 := v37.f17];
			var v78: map<string, bool> := map["fh" := if (v37.f17 in v77) then v77[v37.f17] else fm4(v2, fm4(v2, v72.f17, globalState), globalState)];
			v34[safeIndex(701, v34.Length)] := if ((v8 + v8) in v78) then v78[v8 + v8] else v9[safeIndex(12, v9.Length)] <= |v5|;
			globalState.f16 := -v9[safeIndex(12, v9.Length)];
		}
		
		v9[safeIndex(12, v9.Length)] := if (true) then v9[safeIndex(12, v9.Length)] else v2;
	} else {
		v9 := v9;
		if (v7 > v7) {
			globalState.f16 := v9[safeIndex(12, v9.Length)];
			v4 := v28.f17;
			var v79: array<C0> := new C0[21];
			v79[safeIndex(643, v79.Length)] := v28;
			v79[safeIndex(643, v79.Length)] := v28;
			v2 := safeDivisionInt(v9[safeIndex(12, v9.Length)], if (fm4(v9[safeIndex(12, v9.Length)], v28.f17, globalState) in v35) then v35[fm4(v9[safeIndex(12, v9.Length)], v28.f17, globalState)] else fm0(v9[safeIndex(12, v9.Length)], v10, globalState));
			globalState.f6 := safeModuloInt(v9[safeIndex(12, v9.Length)], -|v5|);
		} else {
			v34[safeIndex(701, v34.Length)] := v28.f17;
			v1 := v1;
			var v80: array<multiset<int>> := new multiset<int>[7];
			v80[safeIndex(686, v80.Length)] := v29;
			v8, v2, v80[safeIndex(686, v80.Length)], v28, globalState.f14 := v8[safeIndex(-0x1ce - v2, |v8|) := v1], safeModuloInt(|map[v9 := v1]|, safeModuloInt(v2, v9[safeIndex(12, v9.Length)])), multiset{-v2}, v28, v9;
			v34[safeIndex(701, v34.Length)] := v37.f17;
			v9[safeIndex(12, v9.Length)] := v2;
		}
		
		if (v5[safeIndex(if (v9[safeIndex(12, v9.Length)] in v10) then v10[v9[safeIndex(12, v9.Length)]] else v2, |v5|)]) {
			var v81: map<string, multiset<int>> := map[v8 := v29];
			var v82: array<multiset<int>> := new multiset<int>[27] [v29, multiset{v2}, multiset{v9[safeIndex(12, v9.Length)]}, v29, v29, v29, v29, v29, v29, multiset([v9[safeIndex(12, v9.Length)]]), v29, v29, v29, v29, v29, if (v8 in v81) then v81[v8] else v29, v29, v29, v29, v29, v29, v29, v29, if (v9[safeIndex(12, v9.Length)] in v30) then v30[v9[safeIndex(12, v9.Length)]] else v29, v29, v29, v29];
			var v83: map<int, map<int, bool>> := map[v2 := map[v2 := v4]];
			var v84 := DC0(v9[safeIndex(12, v9.Length)], v28.f17, seq(abs(0x69), i11  => (v1)), v83);
			var v85: map<D1, C0> := map[DC5(DC3(v10, v9[safeIndex(12, v9.Length)], v84)) := v37];
			var v86: seq<int> := [v9[safeIndex(12, v9.Length)], v9[safeIndex(12, v9.Length)], |v85|, -v9[safeIndex(12, v9.Length)]];
			var v87: map<seq<int>, bool> := map[v86 := v28.f17];
			var v88, v89, v90 := m0(v82, v9[safeIndex(12, v9.Length)], |v87| > v2, v9[safeIndex(12, v9.Length)] != fm0(v2, map[-0x2dc := -0xad], globalState), globalState);
			v5 := v5;
			var v91, v92, v93 := m0(v82, v88, !(690 <= v9[safeIndex(12, v9.Length)]), v37.f17, globalState);
			var v94: set<int> := {v88};
			var v95: seq<string> := [seq(abs(0x263), i13  => (v1))];
			var v96: map<set<int>, seq<string>> := map[v94 * v94 := (seq(abs(0x115), i12  => (v84.cf2))) + v95];
			globalState.f16 := |(if ((v94 * v94) in v96) then v96[v94 * v94] else v95)|;
			v32 := v32.(cf12 := v31);
		} else {
			v2 := 0x30d;
			var v97: array<multiset<int>> := new multiset<int>[11];
			var v98: seq<C0> := [v28, v28];
			var v99, v100, v101 := m0(v97, -(if (!v28.f17) then |v98| else v2), v28.f17, fm4(v9[safeIndex(12, v9.Length)], v34[safeIndex(701, v34.Length)], globalState), globalState);
			v4 := v34[safeIndex(701, v34.Length)];
			v28 := if (v28.f17) then v37 else v37;
			var v102: set<C0> := {v28};
			var v103: map<set<C0>, bool> := map[v102 := v34[safeIndex(701, v34.Length)]];
			var v104: map<bool, array<bool>> := map[fm4(|v103|, v37.f17, globalState) := v34];
			var v105: map<seq<C0>, array<bool>> := map[v98 := v34];
			v104 := v104[fm4(v2, v28.f17, globalState) := if (v98 in v105) then v105[v98] else v34];
		}
		
		globalState.f11 := v2;
		if (true || (v8 <= v8)) {
			v4 := v5[safeIndex(--|v7|, |v5|)];
			v34[safeIndex(701, v34.Length)] := !v37.f17;
			var v106: map<int, C0> := map[if (v28.f17 in v35) then v35[v28.f17] else v2 := v37];
			var v107: seq<map<int, C0>> := [v106];
			v34[safeIndex(701, v34.Length)] := !((v106 + v106) in (v107 + v107));
			var v108: array<multiset<int>> := new multiset<int>[8];
			var v109, v110, v111 := m0(v108, |{v37.f17}|, v9[safeIndex(12, v9.Length)] < (if (v2 in v10) then v10[v2] else |[v9[safeIndex(12, v9.Length)]]|), v37.f17, globalState);
			var v112: map<multiset<bool>, int> := map[v33 := v9[safeIndex(12, v9.Length)]];
			var v115: seq<multiset<bool>> := [v33];
			var v116: array<map<multiset<bool>, int>> := new map<multiset<bool>, int>[8] [v112, v112, v112 + v112, map v113 : multiset<bool> | v113 in v112 :: (v113) := (560), fm10(v37.f17, globalState), map v114 : multiset<bool> | v114 in multiset(v115) :: (v114) := (v9[safeIndex(12, v9.Length)]), fm10(v28.f17, globalState) + v112, v112];
			v116[safeIndex(162, v116.Length)] := v112;
			v116[safeIndex(162, v116.Length)] := v112;
		} else {
			var v117: seq<array<bool>> := [v34];
			v4 := ([v34] + v117) != v117;
			v4 := v34[safeIndex(701, v34.Length)];
			v34[safeIndex(701, v34.Length)], v28, globalState.f6 := fm4(v9[safeIndex(12, v9.Length)], fm4(v2, v28.f17, globalState), globalState), v37, safeModuloInt(-v2, -(v9[safeIndex(12, v9.Length)] - v2));
			var v118: array<set<int>> := new set<int>[8];
			var v119: set<int> := {|v7|, v9[safeIndex(12, v9.Length)]};
			v118[safeIndex(596, v118.Length)] := v119;
			v118[safeIndex(596, v118.Length)] := v119 + v119;
			v34[safeIndex(701, v34.Length)] := (multiset(v5) - v33) <= v33;
		}
		
	}
	
	for i14 := v9[safeIndex(12, v9.Length)] to v2 {
		v4 := "to" <= (if (fm4(-i14, true, globalState)) then v8 else v8);
		var v120: array<multiset<int>> := new multiset<int>[16](i15 => multiset{v9[safeIndex(12, v9.Length)]});
		var v121, v122, v123 := m0(v120, v9[safeIndex(12, v9.Length)], v9[safeIndex(12, v9.Length)] != v9[safeIndex(12, v9.Length)], 't' in v8, globalState);
		var v124, v125, v126 := m0(v120, v121, v7 !! {v34[safeIndex(701, v34.Length)]}, v28.f17, globalState);
		var v127, v128, v129 := m0(v120, safeModuloInt(|v8|, i14), v28.f17, v28.f17, globalState);
	}
	print v1, "\n";
	print v2, "\n";
	print v4, "\n";
	print v5 == [false, true], "\n";
	print v6 == map[2 := true], "\n";
	print v7 == {true, false}, "\n";
	print v8, "\n";
	print v9[0], "\n";
	print v9[1], "\n";
	print v9[2], "\n";
	print v9[3], "\n";
	print v9[4], "\n";
	print v9[5], "\n";
	print v9[6], "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2 == map['d' := 514], "\n";
	print globalState.f3, "\n";
	print globalState.f4, "\n";
	print globalState.f5, "\n";
	print globalState.f6, "\n";
	print globalState.f7 == multiset{514, -1, 2, 4}, "\n";
	print globalState.f8, "\n";
	print globalState.f9, "\n";
	print globalState.f10, "\n";
	print globalState.f11, "\n";
	print globalState.f12, "\n";
	print globalState.f13, "\n";
	print globalState.f14[0], "\n";
	print globalState.f14[1], "\n";
	print globalState.f14[2], "\n";
	print globalState.f14[3], "\n";
	print globalState.f14[4], "\n";
	print globalState.f14[5], "\n";
	print globalState.f14[6], "\n";
	print globalState.f15, "\n";
	print globalState.f16, "\n";
	print v10 == map[514 := 514, -514 := 514], "\n";
	print v28.f17, "\n";
	print v29 == multiset{514}, "\n";
	print v30 == map[0 := multiset{76}], "\n";
	print v31.cf9, "\n";
	print v31.cf10, "\n";
	print v31.cf11 == map[0 := multiset{76}], "\n";
	print v32.cf12.cf9, "\n";
	print v32.cf12.cf10, "\n";
	print v32.cf12.cf11 == map[0 := multiset{76}], "\n";
	print v33 == multiset{false}, "\n";
	print v34[0], "\n";
	print v34[5], "\n";
	print v35 == map[false := 0], "\n";
	print v36 == [map[514 := 514, -514 := 514], map[514 := 514, -514 := 514]], "\n";
	print v37.f17, "\n";
	print v38 == map[0 := "frrp", 2 := "frrp"], "\n";
	print v39[0] == [3355, 514, 129], "\n";
	print v39[1] == [3355, 514, 129], "\n";
	print v39[2] == [3355, 514, 129], "\n";
	print v39[3] == [3355, 514, 129], "\n";
	print v39[4] == [3355, 514, 129], "\n";
	print v39[5] == [3355, 514, 129], "\n";
	print v39[6] == [3355, 514, 129], "\n";
	print v39[7] == [3355, 514, 129], "\n";
	print v39[8] == [3355, 514, 129], "\n";
	print v39[9] == [-874], "\n";
	print v39[10] == [3355, 514, 129], "\n";
	print v39[11] == [3355, 514, 129], "\n";
	print v39[12] == [3355, 514, 129], "\n";
	print v39[13] == [3355, 514, 129], "\n";
	print i6, "\n";
}