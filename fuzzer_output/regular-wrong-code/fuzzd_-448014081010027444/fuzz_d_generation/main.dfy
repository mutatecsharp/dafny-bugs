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
datatype D0 = DC0(cf0: bool)
datatype D1 = DC2(cf2: int, cf3: bool, cf4: D0) | DC3 | DC4(cf5: bool, cf6: bool, cf7: array<bool>, cf8: bool) | DC1(cf1: int) | DC5(cf9: D1)
datatype D2 = DC7(cf11: bool, cf12: int) | DC6(cf10: map<bool, int>) | DC8(cf13: D2)
datatype D3 = DC10(cf15: int, cf16: seq<int>) | DC11(cf17: seq<int>, cf18: bool) | DC9(cf14: seq<int>) | DC12(cf19: D3)
datatype D4 = DC14(cf21: C0, cf22: C0) | DC13(cf20: string)
datatype D5 = DC16(cf24: int, cf25: int) | DC17(cf26: bool) | DC15(cf23: multiset<set<bool>>)
datatype D6 = DC19(cf28: char, cf29: bool, cf30: map<int, int>) | DC18(cf27: set<bool>)
datatype D7 = DC21(cf32: int, cf33: bool) | DC22(cf34: int, cf35: bool, cf36: multiset<int>) | DC20(cf31: array<int>)
datatype D8 = DC23(cf37: multiset<map<int, int>>)
datatype D9 = DC25(cf39: int) | DC26(cf40: bool) | DC27(cf41: bool) | DC24(cf38: seq<D2>)
class GlobalState {
	var f0 : int
	const f1 : seq<bool>
	const f2 : bool
	const f3 : string
	const f4 : int
	var f5 : string
	const f6 : int
	const f7 : char
	const f8 : int
	const f9 : string
	const f10 : int
	const f11 : bool
	const f12 : int
	const f13 : int
	var f14 : array<set<bool>>
	constructor (f0 : int, f1 : seq<bool>, f2 : bool, f3 : string, f4 : int, f5 : string, f6 : int, f7 : char, f8 : int, f9 : string, f10 : int, f11 : bool, f12 : int, f13 : int, f14 : array<set<bool>>) {
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
	}
	
}

class C0 {
	const f15 : array<bool>
	constructor (f15 : array<bool>) {
		this.f15 := f15;
	}
	
}

function fm0(p0: bool, globalState: GlobalState): seq<bool> {
	[true, true] + ([!true] + [!true, true, !true])
}
function fm1(p0: map<int, bool>, p1: string, p2: string, globalState: GlobalState): bool {
	if (false) then true else !!!!!!!false && true
}
function fm2(p0: bool, p1: bool, globalState: GlobalState): map<int, bool> {
	(map[-|{true, true, false}| := false] + map[463 := false]) + (map[|[308, |{0x3a2, 0x131}|, |"kie"|, 0x3ad, |multiset{DC2(|map[false := multiset{true}]|, true, DC0(true))}|]| := true] + map[-0x396 := true])
}
function fm3(p0: int, p1: int, p2: int, globalState: GlobalState): multiset<bool> {
	(multiset{true, true} - multiset{false}) * multiset{true}
}
function fm4(globalState: GlobalState): int {
	if (false) then -885 else |[true, !false, true, false]|
}
function fm5(p0: int, p1: bool, p2: string, p3: multiset<string>, globalState: GlobalState): set<bool> {
	(DC18({false}).cf27 + {false, true}) - ({false} * {!true, !!false})
}
function fm6(p0: string, p1: int, globalState: GlobalState): map<D1, bool> {
	match DC7(true, 0x1ce) {
		case DC7(cf11, cf12) => map[DC2(0x7f, true, DC0(cf11)) := cf11]
		case DC6(cf10) => map[DC2(|map[{|(map v0 : int | v0 in map[|"wdogpmi"| := |(seq(abs(0x2c9), i0  => ('e')))|] :: (safeModuloInt(v0, 0x359)) := (true))|} := 0x3c9]|, false, DC0(!false)) := false]
		case DC8(cf13) => map v1 : D1 | v1 in {DC2(|"ogqm"|, true, DC0(true)), DC2(-0x1e4, !false, DC0(true))} :: (v1) := (true)
	}
}
function fm7(p0: bool, p1: char, p2: int, globalState: GlobalState): seq<int> {
	([0x24d, -|{645}|, -|(seq(abs(-720), i0  => ('d')))|] + (seq(abs(-0x20d), i1  => (0x2c)))) + ([-|multiset{false, !!!true}|] + [|"mqtumvw"|])
}
function fm8(p0: bool, p1: char, p2: char, globalState: GlobalState): string {
	("mrikrvux" + DC13("ur").cf20) + "hytv"
}
function fm9(p0: bool, globalState: GlobalState): char {
	't'
}
function fm10(p0: char, p1: bool, p2: int, globalState: GlobalState): seq<D2> {
	DC24([DC6(map[false := 0x36e])]).cf38
}
function fm11(globalState: GlobalState): set<char> {
	{'k', 'x', 'n', 'x', 'l'}
}
function fm12(globalState: GlobalState): D4 {
	DC13("urylnfh" + "mimbr")
}
function fm13(p0: int, globalState: GlobalState): set<int> {
	set v0 : int | (-0xd0 <= v0) && (v0 < 0x12f) :: (safeDivisionInt(v0, 0x365))
}
method m0(p0: bool, p1: int, p2: bool, globalState: GlobalState) returns (r0: int, r1: int, r2: bool) {
	var v0: multiset<bool> := multiset{true};
	var v1: seq<multiset<bool>> := [v0];
	if ((multiset{p2, !!p2, p2, false} * v0) > (v1[safeIndex(-p1, |v1|)] + v0)) {
		if (safeDivisionInt(p1, p1) !in multiset{p1}) {
			var v2: set<bool> := {!p0, p0};
			r1 := safeDivisionInt(p1 + p1, |v2|);
			var v3: array<int> := new int[15];
			v3[safeIndex(1, v3.Length)] := p1;
			v3[safeIndex(1, v3.Length)] := p1;
			v3[safeIndex(1, v3.Length)] := p1;
			var v4: array<bool> := new bool[10](i0 => p0);
			var v5 := new C0(v4);
			r2 := p2;
		} else {
			var v6: seq<bool> := [p0, p0];
			var v7: seq<int> := [|multiset(v6)|];
			var v8 := DC11(seq(abs(0x343), i1  => (p1)), p2);
			var v9 := "dwnjkq";
			var v10: seq<seq<int>> := [v7];
			var v11 := 'r';
			var v12: array<seq<int>> := new seq<int>[18] [v7, v8.cf17, (seq(abs(-0x280), i2  => (p1))) + v7, v7, [p1, p1, p1, p1, p1], [p1, p1, p1, p1], v7, (seq(abs(880), i3  => (p1)))[safeIndex(p1, |(seq(abs(880), i3  => (p1)))|) := p1], v7[safeIndex(p1, |v7|) := |v9|], v10[safeIndex(p1, |v10|)], v7[safeIndex(fm4(globalState), |v7|) := -0xf6], seq(abs(374), i4  => (0x372)), v7, v7, v7, v7, v7 + v7, fm7(p2, v11, 349, globalState)];
			v12[safeIndex(938, v12.Length)] := v7;
			v12[safeIndex(938, v12.Length)] := v7;
			var v13 := DC9(v12[safeIndex(938, v12.Length)]);
			v13 := v13;
			var v14: array<char> := new char[8] [fm9(p2, globalState), v11, v11, if (p2) then v11 else 'v', v11, v11, v11, v11];
			v14[safeIndex(896, v14.Length)] := v11;
			var v15: array<int> := new int[19](i5 => safeModuloInt(i5, p1));
			r2, v14[safeIndex(896, v14.Length)], v15 := p0, v11, v15;
			var v16: array<multiset<bool>> := new multiset<bool>[29](i6 => v0);
			v16[safeIndex(952, v16.Length)] := v0;
			v16[safeIndex(952, v16.Length)] := v0;
			var v17: array<bool> := new bool[16];
			var v18 := new C0(v17);
		}
		
		var v19: map<int, bool> := map[-p1 := p0];
		var v20 := "fxdguiwrc";
		var v21 := DC13(v20);
		var v22: seq<bool> := [p0, false, p0];
		var v23: map<seq<bool>, bool> := map[v22 := p0];
		var v24 := DC11(seq(abs(0x18), i7  => (p1)), p0);
		var v25: array<bool> := new bool[29] [false, p0, fm1(v19, v21.cf20, "lenrsoqb", globalState), p2, p0, p0, p0, p2, p0, p0, p2, !!p2, true, !p0, false, p2, p2, p0, if (v22 in v23) then v23[v22] else !p0, true, p2, !p0, v24.cf18, p0, p2, p0, p2, !p2, p0];
		var v26 := new C0(v25);
		var v27: array<seq<array<int>>> := new seq<array<int>>[13];
		var v28: array<int> := new int[21](i8 => i8 + p1);
		var v29: seq<array<int>> := [v28];
		v27[safeIndex(305, v27.Length)] := v29;
		v27[safeIndex(305, v27.Length)] := v29;
		var v30 := new C0(v26.f15);
		globalState.f0 := safeModuloInt(p1, -p1);
	} else {
		if (!p0) {
			var v31: array<bool> := new bool[16];
			v31[safeIndex(359, v31.Length)] := p2;
			v31[safeIndex(359, v31.Length)] := p2;
			var v32: C0 := new C0(v31);
			v32 := v32;
			var v33 := new C0(v32.f15);
			var v34: set<bool> := {p0, p0, p2, p2, v31[safeIndex(359, v31.Length)]};
			var v35: multiset<set<bool>> := multiset{v34, v34, v34};
			var v36 := DC15(v35);
			v31[safeIndex(359, v31.Length)] := v36.cf23 !! multiset{v34};
			var v37: array<set<char>> := new set<char>[5];
			v37[safeIndex(540, v37.Length)] := fm11(globalState);
			var v38: set<char> := {'f'};
			var v39 := 'v';
			v37[safeIndex(540, v37.Length)] := (v38 * {'y', v39, v39}) + v38;
		} else {
			var v40 := DC17(p2);
			v40 := if (true) then v40.(cf26 := p0) else v40;
			var v41: array<int> := new int[10];
			v41[safeIndex(913, v41.Length)] := p1;
			v41[safeIndex(913, v41.Length)] := -p1 - -0x241;
			var v42: array<string> := new string[14];
			var v43 := "ctnu";
			v42[safeIndex(658, v42.Length)] := v43;
			v42[safeIndex(658, v42.Length)] := v43;
			var v44: array<bool> := new bool[24](i9 => p0);
			var v45 := new C0(v44);
			var v46 := 'r';
			globalState.f5, globalState.f0 := fm8(p0 <== p0, v46, v46, globalState), p1;
		}
		
		var v47: map<int, bool> := map[p1 := p2];
		var v48: array<bool> := new bool[4] [p0, p2, p2, if (p1 in v47) then v47[p1] else p2];
		var v49 := new C0(v48);
		var v50: map<int, map<int, bool>> := map[p1 := v47 + v47];
		var v51: multiset<C0> := multiset{v49, v49};
		v50 := v50[|v51| * p1 := if (0x46 in v50) then v50[0x46] else v47];
		var v52: C0 := new C0(v49.f15);
		v52 := v52;
		v52.f15[safeIndex(186, v52.f15.Length)] := p2;
		var v53 := "uuolsr";
		var v54: multiset<string> := multiset{v53};
		v52.f15[safeIndex(186, v52.f15.Length)] := v54 != v54;
	}
	
	if (true) {
		r1 := fm4(globalState);
		match (fm12(globalState)) {
			case DC14(cf21, cf22) =>
				r2 := p2;
				r2 := p2;
				var v55 := DC1(-p1);
				var v56 := DC7(p2, v55.cf1);
				v56 := v56;
				var v57: map<bool, bool> := map[p2 := p2];
				var v58: multiset<set<int>> := multiset{{0x209}};
				var v59: set<int> := {p1, -0x2e7};
				v57 := v57[true := v58 !! multiset{v59, v59}];
			case DC13(cf20) =>
				var v60: map<int, int> := map[p1 := |map[0x13c := ((multiset{false})[p0 := abs(p1)])[p0 := abs(884)]]|];
				var v61 := 'd';
				var v62: map<string, bool> := map["jyww" := p0];
				var v63: seq<bool> := [!p0, p2, fm1(map[|v60| := p0], cf20[safeIndex(-p1, |cf20|) := v61], "rswrhg", globalState), if (cf20[safeIndex(p1, |cf20|) := v61] in v62) then v62[cf20[safeIndex(p1, |cf20|) := v61]] else p0];
				var v64: map<bool, seq<bool>> := map[p0 := v63];
				v64 := v64;
				var v65: array<map<bool, bool>> := new map<bool, bool>[20];
				v65[safeIndex(313, v65.Length)] := map[p0 := p2];
				var v66: map<bool, bool> := map[!p0 := if (true) then p0 else p2];
				v65[safeIndex(313, v65.Length)] := v66;
				var v67: array<bool> := new bool[23];
				v67[safeIndex(603, v67.Length)] := p0;
				v67[safeIndex(603, v67.Length)] := 0x3dd != p1;
				var v68: map<bool, int> := map[v67[safeIndex(603, v67.Length)] := p1];
				var v69 := DC6(v68);
				var v70: seq<D2> := [v69, v69, DC6(v68)];
				v67[safeIndex(603, v67.Length)] := v69 in v70;
		}
		
		var v71: array<bool> := new bool[13](i10 => p0);
		var v72 := new C0(v71);
		if (p0) {
			r0 := p1;
			var v73: set<bool> := {true, !true};
			var v74 := DC18(v73);
			r1, v73 := -p1, v74.cf27;
			var v75: array<int> := new int[13];
			var v76 := DC20(v75);
			var v77: array<array<int>> := new array<int>[18] [v75, v75, v75, v75, v75, v75, v75, v75, v75, v75, v76.cf31, v76.cf31, v75, v75, v75, v75, v75, v75];
			var v78: map<bool, array<array<int>>> := map[p2 := v77];
			v78 := v78[p0 := v77];
			globalState.f0 := p1;
			var v79 := "qcyapl";
			r2 := !(if (v79 != v79) then !p2 else false);
		} else {
			var v80: map<int, int> := map[p1 := p1];
			var v81: seq<int> := [|v80|, |fm0(p0, globalState)|, p1];
			var v82 := DC11(v81, p2);
			v82 := v82;
			r1 := p1;
			var v83: array<int> := new int[6](i11 => i11 * p1);
			v83[safeIndex(926, v83.Length)] := -p1;
			v83[safeIndex(926, v83.Length)] := (if (true in v0) then v0[true] else p1) - |"eskgngpo"|;
			var v84 := new C0(v72.f15);
			v72.f15[safeIndex(112, v72.f15.Length)] := p0;
			var v85: set<int> := {v83[safeIndex(926, v83.Length)], p1};
			var v86: seq<bool> := [p0];
			var v87 := 'f';
			var v88: map<bool, char> := map[v86[safeIndex(|v81|, |v86|)] := v87];
			v72.f15[safeIndex(112, v72.f15.Length)] := v85 > ({p1, |[v88]|} - v85);
		}
		
		var v89: map<bool, array<bool>> := map[!p2 := v72.f15];
		v89 := v89[p2 := v71];
	} else {
		var v90: seq<int> := [p1 * p1, p1];
		var v91: array<bool> := new bool[23];
		v91[safeIndex(921, v91.Length)] := p0;
		var v92 := "j";
		var v93 := 'p';
		v90, v91[safeIndex(921, v91.Length)] := v90, (p1 + |v92[safeIndex(p1, |v92|) := v93]|) < v90[safeIndex(p1, |v90|)];
		var v94: C0 := new C0(v91);
		var v95: seq<C0> := [v94];
		var v96: multiset<C0> := multiset{v94, v94, v95[safeIndex(-0x20e, |v95|)], v94};
		v96 := v96;
		var v97: array<multiset<map<int, int>>> := new multiset<map<int, int>>[26];
		var v98: map<int, int> := map[-p1 := p1];
		var v99: multiset<map<int, int>> := multiset{v98};
		var v100 := DC23(v99);
		v97[safeIndex(596, v97.Length)] := v100.cf37;
		v97[safeIndex(596, v97.Length)] := v99;
		var v101: map<bool, bool> := map[p2 := p0];
		var v102: map<map<bool, bool>, bool> := map[v101 := v91[safeIndex(921, v91.Length)]];
		v102 := v102[v101 := |v92| == p1];
		r0 := p1;
	}
	
	var v103: multiset<int> := multiset{p1};
	var v105 := 'a';
	var v106: seq<map<char, int>> := [map v104 : char | v104 in map[v105 := p2] :: (v104) := (|"spyhcfqae"|)];
	var v107: map<char, int> := map[v105 := 0x34c];
	var v108: set<int> := {p1, -p1, |v103|, |v106[safeIndex(fm4(globalState), |v106|) := v107]|, p1};
	var v109: seq<set<int>> := [fm13(0x1f7, globalState)];
	var v110: set<bool> := {p0};
	var v111: multiset<set<bool>> := multiset{v110, {true, p2}, v110, v110};
	var v112 := DC15(v111);
	var v113: map<seq<set<int>>, D5> := map[[{p1}, v108, v109[safeIndex(p1, |v109|)], v108, v108] := v112];
	v113 := v113[v109 := v112];
	var v114: array<bool> := new bool[22](i13 => p2);
	forall i12 | 0 <= i12 < v114.Length {
		v114[i12] := if (p0) then !p2 else p1 >= |"ldv"|;
	}
	var i14 := 0;
	while (v0 < (v0[p2 := abs(-p1)])[false := abs(p1)])
		decreases 100 - i14
	{
		if (i14 >= 100) {
			break;
		}
		
		i14 := i14 + 1;
		r2 := if (p2) then p0 else !(if (p0) then p2 else p2);
		r2 := if (p2) then false else p2;
		v114[safeIndex(117, v114.Length)] := !p0;
		v114[safeIndex(117, v114.Length)] := false;
		var v115 := DC16(|map[p1 := p0]| * p1, p1);
		var v116 := "gsrufvirc";
		v115, globalState.f0, v114[safeIndex(117, v114.Length)], v114[safeIndex(117, v114.Length)] := DC16(p1, p1).(cf24 := 0x24e), |v116|, !p2, v114[safeIndex(117, v114.Length)];
	}
	var v117 := new C0(v114);
	r0 := p1;
	r1 := -(0x3a9 + fm4(globalState));
	r2 := !p0;
}
method Main() {
	var v0 := true;
	var v1 := "fyhqxeqer";
	var v2: array<set<bool>> := new set<bool>[28];
	var globalState := new GlobalState(-0x3d6, [v0], false, v1 + v1, 0x2e9, v1, 0x124, 'x', 0x1a4, v1 + v1, 835, false, -89, 0x1e8, v2);
	var v3: seq<bool> := [v0, v0];
	v3 := v3 + fm0(v0, globalState);
	var v4: seq<seq<bool>> := [v3];
	var v5: map<bool, bool> := map[DC0(true).cf0 := v0];
	var v6: map<bool, string> := map[v4[safeIndex(|v5|, |v4|)] != v3 := v1];
	v6 := v6[fm1(fm2(v0, v0, globalState), "xqfstui", seq(abs(0x334), i0  => ('v')), globalState) := v1 + v1];
	var i1 := 0;
	while (if (v0) then v0 else v0)
		decreases 100 - i1
	{
		if (i1 >= 100) {
			break;
		}
		
		i1 := i1 + 1;
		var v7 := 795;
		var v8, v9, v10 := m0(!v0, -v7, v0, globalState);
		var v11: array<array<bool>> := new array<bool>[29];
		var v12: array<bool> := new bool[20];
		v11 := new array<bool>[27] [v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12];
		var v13: array<D0> := new D0[15];
		var v14 := DC0(v0);
		v13[safeIndex(349, v13.Length)] := v14;
		v13[safeIndex(349, v13.Length)] := v14;
		v7, v10 := v7, v10;
	}
	var v15 := 0x105;
	globalState.f0 := v15 * v15;
	var v16 := DC0(!v0);
	v16 := v16;
	globalState.f0 := safeDivisionInt(--(if (v0) then v15 else v15), v15 - v15);
	var i2 := 0;
	while (DC0(v0).cf0)
		decreases 100 - i2
	{
		if (i2 >= 100) {
			break;
		}
		
		i2 := i2 + 1;
		var v17: array<bool> := new bool[28](i3 => v0);
		v17[safeIndex(902, v17.Length)] := v0;
		var v18: array<D0> := new D0[1] [v16];
		var v19: map<array<D0>, int> := map[v18 := v15];
		v17[safeIndex(902, v17.Length)] := v19 == v19;
		var v20 := 'b';
		globalState.f0, v15, v0, globalState.f5, v20 := v15, safeModuloInt(v15, v15), !v17[safeIndex(902, v17.Length)], v1, 'c';
		var v21: set<seq<bool>> := {[v3[safeIndex(|fm3(v15, v15, v15, globalState)|, |v3|)]], v3};
		var v22: map<int, int> := map[safeDivisionInt(|v21|, 0x1aa) := v15];
		v22 := (v22 + map[v15 := v15]) + v22;
		var v23: multiset<int> := multiset{v15};
		var v24: map<multiset<int>, int> := map[v23[v15 := abs(v15)] := v15];
		v3 := v3[safeIndex(|v24|, |v3|) := v0] + v3;
	}
	var v25: array<bool> := new bool[28];
	var v26: set<array<bool>> := {v25, v25};
	var v27: map<int, set<array<bool>>> := map[fm4(globalState) := v26];
	var v28: array<set<array<bool>>> := new set<array<bool>>[25] [v26, {v25, v25}, v26, {v25}, v26 + v26, {v25}, v26, v26, {v25, v25, v25}, v26, v26 * v26, v26, v26, v26, {v25}, v26, v26, v26, {v25, v25} + v26, v26, v26, (if (v15 in v27) then v27[v15] else v26) + v26, v26, if (v3[safeIndex(v15, |v3|)]) then v26 else v26, v26];
	v28[safeIndex(693, v28.Length)] := v26;
	v28[safeIndex(693, v28.Length)] := v26;
	match (v16) {
		case DC0(cf0) =>
			v6 := v6[!v0 <==> v0 := v1 + v1];
			var v29: set<bool> := {cf0};
			var v30: array<int> := new int[3];
			v30[safeIndex(733, v30.Length)] := v15;
			var v31: set<int> := {-v15};
			var v32: multiset<string> := multiset{v1};
			v29, cf0, v30[safeIndex(733, v30.Length)] := fm5(v15, |v31| < v15, v1 + v1, v32 + v32, globalState), v0, v15 + fm4(globalState);
			var v33: map<int, bool> := map[v30[safeIndex(733, v30.Length)] := v0];
			var v34 := 'f';
			globalState.f0, cf0, v0, v30[safeIndex(733, v30.Length)], cf0 := 373, v29 > ({cf0} * v29), fm1(v33, "mtng", v1, globalState), -(v30[safeIndex(733, v30.Length)] + (if (cf0) then |v1[safeIndex(v30[safeIndex(733, v30.Length)], |v1|) := v34]| else v30[safeIndex(733, v30.Length)])), v15 >= |v1|;
			if (!false) {
				var v35 := new C0(v25);
				v0 := fm4(globalState) >= (|v3| * -v30[safeIndex(733, v30.Length)]);
				var v36, v37, v38 := m0(cf0, v30[safeIndex(733, v30.Length)], cf0, globalState);
				v5 := v5 + (v5 + map[!cf0 := v0]);
				var v39, v40, v41 := m0(fm4(globalState) > v36, -(v37 + DC1(|(seq(abs(0x30f), i4  => (v34)))|).cf1), cf0, globalState);
			} else {
				var v42: map<int, int> := map[|v33| := v30[safeIndex(733, v30.Length)]];
				var v43: map<bool, int> := map[cf0 := |v42|];
				var v44: map<int, map<bool, int>> := map[v15 := v43];
				var v45: seq<int> := [v30[safeIndex(733, v30.Length)], v15, v30[safeIndex(733, v30.Length)], |(if (v15 in v44) then v44[v15] else v43[false := v15])|, |v31|];
				var v46, v47, v48 := m0(v0, if (cf0) then |v45| else v30[safeIndex(733, v30.Length)], !(|v5| != v30[safeIndex(733, v30.Length)]), globalState);
				v6 := v6[false := "yhphth"];
				var v49, v50, v51 := m0(true && cf0, 0x339, fm1(fm2(v48, v0, globalState), v1, v1, globalState), globalState);
				globalState.f5 := v1;
				globalState.f5 := v1 + (v1 + (seq(abs(0x48), i5  => (v34))));
			}
			
	}
	
	var v52 := new C0(v25);
	var v53: set<bool> := {v0};
	var v54: multiset<string> := multiset{v1, v1};
	var v55 := DC4(v53 >= fm5(867, !false, v1, v54, globalState), v0, v25, v0);
	match (v55) {
		case DC2(cf2, cf3, cf4) =>
			var v56 := DC2(cf2, v0, DC0(v0));
			var v57: map<D1, bool> := map[v56 := cf3];
			var v58: multiset<map<D1, bool>> := multiset{v57, map[DC2(cf2, false, v16) := cf3]};
			var v59: seq<int> := [357, if (fm6("yhshk", |v3|, globalState) in v58) then v58[fm6("yhshk", |v3|, globalState)] else cf2, cf2, cf2, v15];
			var v60 := 'd';
			var v61: array<C0> := new C0[28];
			var v62: map<int, array<C0>> := map[v15 := v61];
			var v63: array<seq<int>> := new seq<int>[19] [v59, v59, [v15] + v59, (seq(abs(0x11d), i6  => (|map[v15 := true]|))) + (seq(abs(-0x1ff), i7  => (v15))), seq(abs(363), i8  => (v15)), v59 + v59, v59, v59, v59 + v59, seq(abs(637), i9  => (-v15)), seq(abs(-0x115), i10  => (cf2)), v59 + v59, v59 + v59, v59, (if (true) then v59 else fm7(cf3, v60, cf2, globalState))[safeIndex(|v62[v15 := v61]|, |(if (true) then v59 else fm7(cf3, v60, cf2, globalState))|) := cf2], v59, v59, v59, seq(abs(0x2a2), i11  => (|v3|))];
			v63[safeIndex(598, v63.Length)] := v59;
			v63[safeIndex(598, v63.Length)] := (v59 + v59)[safeIndex(cf2 * 0x19, |(v59 + v59)|) := |v1|];
			cf3, v52, v0, v15, v15 := !!DC0(cf3).cf0, v52, v0, cf2, v15;
			var v64, v65, v66 := m0(v3[safeIndex(-cf2, |v3|)], v15, v0, globalState);
			var v67: map<bool, seq<bool>> := map[v66 := v3];
			v15 := |(v67 + v67)| * 0xd6;
		case DC3() =>
			globalState.f0 := -v15;
			v52.f15[safeIndex(441, v52.f15.Length)] := v0;
			v52.f15[safeIndex(441, v52.f15.Length)] := v0;
			v52.f15[safeIndex(441, v52.f15.Length)] := v0;
			var v69: array<int> := new int[11](i12 => safeModuloInt(i12, |(map v68 : int | v68 in [v15] :: (safeDivisionInt(v68, v15)) := (v52.f15[safeIndex(441, v52.f15.Length)]))|));
			v69 := v69;
		case DC4(cf5, cf6, cf7, cf8) =>
			var v70: array<string> := new string[25];
			v70[safeIndex(675, v70.Length)] := "cccpixh";
			v70[safeIndex(675, v70.Length)] := v1;
			if (cf8) {
				globalState.f0 := v15;
				var v71: array<int> := new int[22];
				v71[safeIndex(181, v71.Length)] := fm4(globalState);
				v71[safeIndex(181, v71.Length)] := fm4(globalState);
				var v72: multiset<bool> := multiset{true, cf5};
				v0, cf6, v1 := !(v55.cf5 <== (v71[safeIndex(181, v71.Length)] == |(seq(abs(0x226), i13  => (v15)))|)), v72 > v72, v70[safeIndex(675, v70.Length)];
				var v73: map<int, bool> := map[v71[safeIndex(181, v71.Length)] := cf6];
				var v74 := 'b';
				cf8 := fm1(v73, v1, fm8(cf8, v74, fm9(cf5, globalState), globalState), globalState);
				var v75: array<C0> := new C0[15];
				var v76: seq<array<C0>> := [v75, v75, v75, v75, v75];
				v76 := v76;
			} else {
				globalState.f0 := v15;
				globalState.f0 := v15 + v15;
				var v77, v78, v79 := m0(v15 != v15, -v15, cf8, globalState);
				var v80 := new C0(v52.f15);
				v52.f15[safeIndex(295, v52.f15.Length)] := v0;
				v52.f15[safeIndex(295, v52.f15.Length)] := cf5;
			}
			
			var v81: map<int, bool> := map[v15 := cf8];
			v81 := v81[v15 := v0];
			globalState.f5 := v1;
		case DC1(cf1) =>
			globalState.f0 := safeDivisionInt(cf1, cf1);
			var v82: map<int, bool> := map[v15 := v0];
			var v83: map<int, int> := map[v15 := |v82|];
			var v84 := DC1(cf1);
			globalState.f0 := if (true) then cf1 else |v83[-v84.cf1 := cf1]|;
			v0 := v0;
			globalState.f0 := v15;
		case DC5(cf9) =>
			var v85: map<int, int> := map[v15 := v15];
			globalState.f0 := -(safeModuloInt(v15, v15) - (if (v15 in v85) then v85[v15] else v15));
			var v86 := new C0(v25);
			var v87: map<bool, int> := map[if (v0 in v5) then v5[v0] else v0 := v15];
			var v88 := DC6(v87);
			v87 := (v88.cf10 + map[v0 := v15]) + v87;
			v0 := v0;
	}
	
	var v89: map<int, bool> := map[v15 := false];
	var v90: map<bool, seq<int>> := map[fm1(v89, v1, "vvidmhm", globalState) := fm7(!v0, 'p', v15, globalState)];
	var v91: seq<int> := [|v53|, fm4(globalState), v15];
	v90 := v90[v0 := v91];
	var v92: map<string, string> := map["xndmxbthy" := v1];
	var v93: map<map<string, string>, bool> := map[v92 := v0];
	v93 := v93[map v94 : string | v94 in v54 :: (v94) := (v1) := if (v0) then v0 else v0];
	var v95 := 'p';
	var v96: multiset<bool> := multiset{v0, v0, v0};
	var v97: seq<multiset<bool>> := [v96];
	var v98: map<seq<D2>, array<bool>> := map[fm10(v95, v0, |v97|, globalState) := v52.f15];
	var v99: map<bool, int> := map[v0 := v15];
	var v100 := DC6(v99);
	var v101: seq<D2> := [v100, v100];
	v98 := v98[v101 := v52.f15];
	var v102, v103, v104 := m0(v0, |v3|, if (v0) then v0 else false, globalState);
	v104 := !v0;
	print v0, "\n";
	print v1, "\n";
	print globalState.f0, "\n";
	print globalState.f1 == [true], "\n";
	print globalState.f2, "\n";
	print globalState.f3, "\n";
	print globalState.f4, "\n";
	print globalState.f5, "\n";
	print globalState.f6, "\n";
	print globalState.f7, "\n";
	print globalState.f8, "\n";
	print globalState.f9, "\n";
	print globalState.f10, "\n";
	print globalState.f11, "\n";
	print globalState.f12, "\n";
	print globalState.f13, "\n";
	print v3 == [true, false, true, true, false, false, true, false, true, true, true, true, false, false, true, false], "\n";
	print v4 == [[true, true, true, true, false, false, true, false]], "\n";
	print v5 == map[true := false], "\n";
	print v6 == map[false := "fyhqxeqerfyhqxeqer", true := "fyhqxeqerfyhqxeqer"], "\n";
	print i1, "\n";
	print v15, "\n";
	print v16.cf0, "\n";
	print i2, "\n";
	print v25[15], "\n";
	print |v26|, "\n";
	print |v27|, "\n";
	print |v28[0]|, "\n";
	print |v28[1]|, "\n";
	print |v28[2]|, "\n";
	print |v28[3]|, "\n";
	print |v28[4]|, "\n";
	print |v28[5]|, "\n";
	print |v28[6]|, "\n";
	print |v28[7]|, "\n";
	print |v28[8]|, "\n";
	print |v28[9]|, "\n";
	print |v28[10]|, "\n";
	print |v28[11]|, "\n";
	print |v28[12]|, "\n";
	print |v28[13]|, "\n";
	print |v28[14]|, "\n";
	print |v28[15]|, "\n";
	print |v28[16]|, "\n";
	print |v28[17]|, "\n";
	print |v28[18]|, "\n";
	print |v28[19]|, "\n";
	print |v28[20]|, "\n";
	print |v28[21]|, "\n";
	print |v28[22]|, "\n";
	print |v28[23]|, "\n";
	print |v28[24]|, "\n";
	print v52.f15[15], "\n";
	print v53 == {false}, "\n";
	print v54 == multiset{"fyhqxeqer", "fyhqxeqer"}, "\n";
	print v55.cf5, "\n";
	print v55.cf6, "\n";
	print v55.cf7[15], "\n";
	print v55.cf8, "\n";
	print v89 == map[0 := false], "\n";
	print v90 == map[true := [589, -1, -720, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, -2, 7], false := [1, 4, 0]], "\n";
	print v91 == [1, 4, 0], "\n";
	print v92 == map["xndmxbthy" := "fyhqxeqer"], "\n";
	print v93 == map[map["xndmxbthy" := "fyhqxeqer"] := false, map["fyhqxeqer" := "fyhqxeqer"] := false], "\n";
	print v95, "\n";
	print v96 == multiset{false, false, false}, "\n";
	print v97 == [multiset{false, false, false}], "\n";
	print |v98|, "\n";
	print v99 == map[false := 0], "\n";
	print v100.cf10 == map[false := 0], "\n";
	print v101 == [DC6(map[false := 0]), DC6(map[false := 0])], "\n";
	print v102, "\n";
	print v103, "\n";
	print v104, "\n";
}