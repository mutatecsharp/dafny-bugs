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
datatype D0 = DC1(cf1: int) | DC2(cf2: bool, cf3: bool, cf4: int) | DC0(cf0: seq<bool>)
datatype D1 = DC4(cf6: int, cf7: bool, cf8: bool, cf9: bool) | DC3(cf5: map<set<bool>, string>)
datatype D2 = DC6(cf11: string, cf12: int, cf13: bool, cf14: int, cf15: bool) | DC5(cf10: string)
datatype D3 = DC8(cf17: int, cf18: int, cf19: bool) | DC7(cf16: array<int>)
datatype D4 = DC10(cf21: seq<int>, cf22: bool) | DC11(cf23: C0, cf24: int, cf25: int) | DC12(cf26: bool, cf27: seq<map<int, int>>, cf28: int, cf29: array<bool>, cf30: bool) | DC9(cf20: map<D3, seq<char>>)
datatype D5 = DC13(cf31: char)
datatype D6 = DC14(cf32: map<int, int>)
class GlobalState {
	var f0 : int
	const f1 : set<bool>
	const f2 : array<seq<bool>>
	var f3 : array<map<bool, int>>
	var f4 : char
	const f5 : string
	const f6 : int
	const f7 : int
	const f8 : bool
	var f9 : char
	const f10 : bool
	const f11 : int
	const f12 : bool
	const f13 : bool
	constructor (f0 : int, f1 : set<bool>, f2 : array<seq<bool>>, f3 : array<map<bool, int>>, f4 : char, f5 : string, f6 : int, f7 : int, f8 : bool, f9 : char, f10 : bool, f11 : int, f12 : bool, f13 : bool) {
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
	}
	
}

class C0 {
	const f14 : D3
	constructor (f14 : D3) {
		this.f14 := f14;
	}
	
	function fm6(p0: bool, p1: int, p2: bool, globalState: GlobalState): int {
		safeDivisionInt(|[-847, |"uqfinko"|]|, 15) - 0x267
	}
}

class C1 {
	constructor () {
	}
	
	function fm4(globalState: GlobalState): map<set<bool>, string> {
		DC3(map[{true, false, false} := "saf"]).cf5 + (map v0 : set<bool> | v0 in map[{true} := false] :: (v0) := ("byiwjxmen"))
	}
	function fm5(globalState: GlobalState): map<string, string> {
		(map[seq(abs(209), i0  => ('a')) := "kipi"] + map["vbvdrrtth" := "n"]) + map[DC5("rt").cf10 := "goeperti"]
	}
	method m1(p0: map<set<int>, bool>, p1: string, p2: bool, globalState: GlobalState) returns (r0: int, r1: bool, r2: set<int>, r3: int) {
		var v0 := 0x34;
		var v1 := DC1(v0);
		r3 := (v1.(cf1 := v0)).cf1;
		for i0 := |p1| to v0 + v0 {
			r1 := !p2;
			r1, v0, r1 := p2, i0, p2;
			r1 := p2;
			var v2: seq<int> := [|(p1 + p1)|];
			globalState.f0, r1 := v2[safeIndex(v0, |v2|)], fm2(globalState);
		}
		var v3: array<bool> := new bool[7];
		v3[safeIndex(729, v3.Length)] := p2;
		v3[safeIndex(729, v3.Length)] := p2 <== !p2;
		r1, v3[safeIndex(729, v3.Length)] := !(v0 >= |p1|), p2;
		for i1 := v0 - v0 to v0 {
			var v4 := DC2(v3[safeIndex(729, v3.Length)], v3[safeIndex(729, v3.Length)], i1);
			var v5: seq<D0> := [v4, v4];
			var v6: map<int, int> := map[v0 := 0xb2];
			var v7: multiset<map<int, int>> := multiset{v6, v6};
			var v8: seq<bool> := [p2, v3[safeIndex(729, v3.Length)]];
			var v9: set<seq<bool>> := {v8};
			var v10: seq<int> := [|v5|, -|v7|, i1, |v9|];
			r3, v3[safeIndex(729, v3.Length)] := v0, v10 == (v10 + v10);
			if (fm2(globalState)) {
				var v11: map<bool, bool> := map[p2 := fm2(globalState)];
				var v12: multiset<bool> := multiset{fm2(globalState), p2};
				v11 := v11[v12 > multiset{v3[safeIndex(729, v3.Length)]} := v3[safeIndex(729, v3.Length)]];
				v11 := v11[p2 := p2];
				var v13: array<int> := new int[23];
				var v14 := DC7(v13);
				var v15: map<bool, array<int>> := map[v3[safeIndex(729, v3.Length)] := v14.cf16];
				var v16: seq<array<int>> := [v13, v13, v13, v13];
				v15 := v15[v3[safeIndex(729, v3.Length)] := v16[safeIndex(i1, |v16|)]];
				var v17: array<seq<array<int>>> := new seq<array<int>>[29];
				v17[safeIndex(222, v17.Length)] := v16 + v16;
				var v18: map<int, bool> := map[v0 := p2];
				var v19: seq<map<int, bool>> := [v18, v18];
				v17[safeIndex(222, v17.Length)] := ((v16 + v16) + v16)[safeIndex(safeDivisionInt(|v19|, |p1|), |((v16 + v16) + v16)|) := if (p2 in v15) then v15[p2] else v13];
				v11 := v11[p2 := p2];
			} else {
				v0 := |v8|;
				var v20 := 'q';
				globalState.f4 := v20;
				v6 := v6[i1 := v0];
				v3[safeIndex(729, v3.Length)] := fm2(globalState);
				var v21: array<int> := new int[1];
				var v22: seq<array<int>> := [v21, v21];
				var v23: map<D3, int> := map[DC7(v22[safeIndex(0x2cc, |v22|)]) := v0];
				var v24: array<map<D3, int>> := new map<D3, int>[1] [v23];
				v24[safeIndex(445, v24.Length)] := v23 + v23;
				v24[safeIndex(445, v24.Length)] := v23;
			}
			
			globalState.f0 := v0;
			var v25 := DC8(i1, v0, v3[safeIndex(729, v3.Length)]);
			var v26 := new C0(v25);
		}
		globalState.f0 := safeDivisionInt(v0, v0 * |p1|);
		r0 := -|p1|;
		r1 := !p2;
		var v27: set<int> := {v0};
		r2 := v27 + v27;
		r3 := safeDivisionInt(-v0 - v0, v0);
	}
}

function fm0(globalState: GlobalState): int {
	DC2(false, false, -0x3ad).cf4
}
function fm1(p0: bool, p1: int, globalState: GlobalState): string {
	("csbalwk" + (seq(abs(-0x55), i0  => ('t')))) + (seq(abs(394), i1  => ('a')))
}
function fm2(globalState: GlobalState): bool {
	if (!true ==> false) then {-0x47} !! {-0x233} else false
}
function fm3(p0: int, p1: int, p2: int, p3: map<int, bool>, globalState: GlobalState): seq<bool> {
	([false, !!!false] + [true]) + [!!true, true]
}
function fm7(p0: bool, globalState: GlobalState): map<bool, bool> {
	(map[false := false] + map[true := false]) + (map[true := true] + map[false := !true])
}
function fm8(p0: int, globalState: GlobalState): map<int, bool> {
	map[450 := false] + map[-0x1e8 := false]
}
function fm9(globalState: GlobalState): set<bool> {
	({false, !true} - {false, true}) * {false}
}
function fm10(p0: bool, p1: string, globalState: GlobalState): D0 {
	if (!!false) then if (true) then DC2(false, true, -|map[981 := |(seq(abs(0x2f6), i0  => ('a')))|]|) else DC2(!false, DC8(0x1af, -0x133, true).cf19, |multiset([true])|) else DC2(DC2(!true, true, |[0x3af]|).cf3, true, 0x237)
}
function fm11(p0: char, globalState: GlobalState): D2 {
	DC6("ft", |{true}| + DC8(|{|{true, true}|}|, 0x31e, !!true).cf17, (set v0 : map<bool, bool> | v0 in {map[false := false], map[true := true]} :: (v0)) > {map[true := true]}, |(if (true) then multiset{true, true} else multiset{DC8(0xab, |"otasatp"|, false).cf19, false, false})|, multiset{|map[--|multiset{false}| := [false]]|, 0x76} !! multiset{587})
}
function fm12(p0: char, p1: bool, p2: int, globalState: GlobalState): multiset<int> {
	multiset{0xef}
}
function fm13(p0: int, p1: string, p2: bool, globalState: GlobalState): D4 {
	if (!(false <==> !!true)) then DC9(map v0 : D3 | v0 in multiset{DC8(0x2fc, 0xa8, !false)} :: (v0) := (['u'])) else if (true) then DC9(map[DC8(-0x298, |[0x5a]|, true) := ['x']]) else DC9(map[DC8(|map[-0x345 := true]|, |map[false := |['i']|]|, true) := seq(abs(82), i0  => ('y'))])
}
function fm14(p0: char, p1: int, globalState: GlobalState): set<int> {
	{0x131}
}
function fm15(p0: int, p1: int, globalState: GlobalState): D0 {
	DC0([!false, true])
}
function fm16(p0: int, p1: int, globalState: GlobalState): multiset<bool> {
	multiset{!true, !(!!DC2(true, false, 0x2df).cf2 && false)}
}
function fm17(p0: int, p1: bool, p2: bool, globalState: GlobalState): seq<map<int, int>> {
	match if (!false) then DC6("mjffy", |multiset{false}|, false, |map[true := [false, false, !false]]|, false) else DC6("l", 139, !true, |"dop"|, false) {
		case DC6(cf11, cf12, cf13, cf14, cf15) => (seq(abs(0x30c), i0  => (map v0 : int | (0x316 <= v0) && (v0 < 0x3ca) :: (safeDivisionInt(v0, cf14)) := (857)))) + (seq(abs(0x2bd), i1  => (map[-290 := cf12])))
		case DC5(cf10) => seq(abs(0x3d4), i2  => (DC14(map v1 : int | (-768 <= v1) && (v1 < 0x18a) :: (v1 * -880) := (586)).cf32))
	}
}
method m0(p0: bool, p1: int, p2: bool, p3: int, globalState: GlobalState) {
	if (p0) {
		var v0 := true;
		var v1: seq<int> := [p3];
		v0 := p1 !in v1;
		var v2 := DC8(-0x49, p3, v0);
		var v3: C0 := new C0(v2);
		var v4: seq<bool> := [v0];
		var v5: map<int, int> := map[|v4| := |v4|];
		var v6: map<bool, int> := map[p2 := |(multiset{p0, p2})[p0 := abs(p3)]|];
		var v7: array<bool> := new bool[13];
		var v8 := DC12(v0, [v5, v5, v5, map[p3 := |v6|], v5], p1, v7, p0);
		var v9: multiset<map<bool, int>> := multiset{map[v0 := v8.cf28], v6, v6};
		var v10 := DC11(v3, if (map[v0 := v3.fm6(p0, p1, !true, globalState)] in v9) then v9[map[v0 := v3.fm6(p0, p1, !true, globalState)]] else p3, p3);
		v10 := v10;
		v7[safeIndex(74, v7.Length)] := false;
		v7[safeIndex(74, v7.Length)] := fm2(globalState);
		v7[safeIndex(74, v7.Length)] := p0;
		if (p2 && fm2(globalState)) {
			var v11 := new C1();
			var v12: array<array<seq<bool>>> := new array<seq<bool>>[2];
			var v13 := 'j';
			var v14: map<char, seq<bool>> := map[v13 := v4];
			var v15 := DC0([!p2]);
			var v16: array<seq<bool>> := new seq<bool>[16] [[p0], v4, v4, [p2, v7[safeIndex(74, v7.Length)]], v4, v4, [false, true, v0], v4, [true], v4, v4, [v0, v7[safeIndex(74, v7.Length)]], v4, v4, v4, if (v13 in v14) then v14[v13] else v15.cf0];
			v12[safeIndex(557, v12.Length)] := if (!v7[safeIndex(74, v7.Length)]) then v16 else v16;
			v12[safeIndex(557, v12.Length)] := v16;
			var v17 := "d";
			var v18: map<bool, char> := map[p1 > |v17| := 'i'];
			var v19 := DC13(v13);
			v18 := v18[p2 := v19.cf31];
			v2 := v2;
			var v20 := new C0(v2);
		} else {
			var v21 := new C1();
			var v22 := new C1();
			globalState.f0 := p3;
			var v23 := "mmk";
			v23 := v23[safeIndex(p1, |v23|) := v23[safeIndex(p1, |v23|)]];
			var v24: seq<C0> := [v3];
			var v25 := DC4(p3, !p0, true, false);
			var v26: array<C0> := new C0[19] [v3, v3, v3, v3, v24[safeIndex(v25.cf6, |v24|)], v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3];
			var v27: map<array<C0>, bool> := map[v26 := v0];
			v27 := v27[v26 := p2];
		}
		
	} else {
		var v28: map<int, int> := map[DC8(p3, p3, p0).cf18 := |"slgt"|];
		var v29: multiset<bool> := multiset{p0, map[0x14e := p3] in map[v28 := DC4(p1, p2, !p2, false)], p2, p0, p0};
		v29 := v29 - v29;
		if (p2) {
			globalState.f0 := 0x46;
			var v30 := new C0(DC8(p3, fm0(globalState), p2));
			var v31 := 'c';
			var v32: array<bool> := new bool[15];
			var v33: map<char, array<bool>> := map[v31 := v32];
			v33 := v33[v31 := v32];
			var v34 := true;
			v34 := !(v34 || p2);
			var v35: seq<bool> := [!p0, p0];
			var v36 := DC9(map[v30.f14 := [v31, v31]]);
			var v37 := "osydgy";
			var v38: map<D3, seq<char>> := map[v30.f14 := v37];
			var v39: multiset<int> := multiset{p1, |fm1(p0, -p1, globalState)|, p1};
			var v40: array<D4> := new D4[20] [v36, v36, fm13(p3, v37, v34, globalState), v36, v36, v36, DC9(v38), v36, v36, DC9(v38), if (true) then v36 else v36, v36.(cf20 := v38).(cf20 := v38), v36, v36.(cf20 := v38).(cf20 := v38), if (p0) then v36.(cf20 := v38) else v36, v36, v36, DC9(v38), DC9(v38), fm13(|v39|, v37, v34, globalState)];
			v40[safeIndex(103, v40.Length)] := v36;
			var v41: map<int, string> := map[p3 := fm1(p0, |v37|, globalState)];
			var v42: C1 := new C1();
			var v43: map<bool, C1> := map[p2 := v42];
			var v44: map<int, bool> := map[|v43| := v34];
			v35, v40[safeIndex(103, v40.Length)] := fm3(p1, safeModuloInt(p1, p3), |v41| - p1, if (p0) then v44 else v44, globalState), if (p0) then v36.(cf20 := v38) else v36;
		} else {
			var v45 := true;
			var v46: array<array<int>> := new array<int>[11];
			var v47: array<int> := new int[8](i0 => i0 - -p3);
			v46[safeIndex(803, v46.Length)] := v47;
			var v48 := DC8(p1, p3, p0);
			var v49: C0 := new C0(v48);
			var v50: seq<C0> := [v49];
			var v51: seq<bool> := [p2, p2, p2];
			var v52: set<int> := {p1, |v51|};
			var v53: seq<set<int>> := [v52];
			var v54 := DC8(p3, |v50[safeIndex(|v53|, |v50|) := v49]|, p0);
			var v55: array<bool> := new bool[12] [v45, p2, v45, p2, v45, p0, v45, v45, v48.cf19, !p2, p0, p0];
			var v56: seq<array<bool>> := [v55];
			var v57: map<array<bool>, array<bool>> := map[v55 := v55];
			var v58: C1 := new C1();
			var v59: map<C1, array<bool>> := map[v58 := v55];
			var v60: array<array<bool>> := new array<bool>[12] [v56[safeIndex(p3, |v56|)], v55, v55, v55, v55, if (v55 in v57) then v57[v55] else v55, v55, v55, v55, v55, if (true) then v55 else v55, if (v58 in v59) then v59[v58] else v55];
			v60[safeIndex(247, v60.Length)] := v55;
			v45, v46[safeIndex(803, v46.Length)], v60[safeIndex(247, v60.Length)], v45 := v45, v47, v55, p2;
			var v61: array<set<int>> := new set<int>[26](i1 => v52);
			v61[safeIndex(590, v61.Length)] := v52;
			var v62 := 'e';
			v61[safeIndex(590, v61.Length)] := (fm14(v62, -p1, globalState) * v52) + v52;
			var v63: multiset<int> := multiset{p3};
			var v64: multiset<D0> := multiset{fm15(-|v63|, p1, globalState), DC0(v51)};
			var v65: map<bool, multiset<D0>> := map[p0 <== v45 := v64];
			v65 := v65[!p0 := v64];
			v45 := v51[safeIndex(p3, |v51|)];
			var v66: seq<multiset<bool>> := [v29];
			v66 := if (p0) then v66 else [multiset([p0]), v29, v29, fm16(p3, p1, globalState), v29];
		}
		
		var v67 := "iwvnrgqkx";
		var v68: map<int, bool> := map[|"jgktoekf"| := true];
		globalState.f0 := -safeModuloInt(|v67|, safeModuloInt(489, |v68|));
		var v69: map<int, multiset<bool>> := map[p1 := v29];
		var v70: seq<bool> := [p0];
		v69 := v69[p3 := multiset(v70)];
		var v71 := true;
		v71 := safeModuloInt(p1, p3) >= p3;
	}
	
	globalState.f0 := p1;
	var v72 := "oncu";
	var v73: seq<int> := [0x3, 0xec];
	for i2 := -|v72| to -v73[safeIndex(p1, |v73|)] {
		var v74: map<int, string> := map[p3 := v72];
		var v75 := 'i';
		var v76: map<map<int, string>, char> := map[v74 := v75];
		v76 := v76[v74 := v75];
		v72 := v72;
		globalState.f0 := p3;
		var v77: array<int> := new int[7](i3 => i3 - p1);
		v77[safeIndex(883, v77.Length)] := i2;
		var v78: multiset<int> := multiset{p3, p3};
		var v79: seq<multiset<int>> := [v78];
		var v80: C1 := new C1();
		var v81: map<C1, C1> := map[v80 := v80];
		v77[safeIndex(883, v77.Length)] := -(p3 - (if (p0) then |v79| else -|v81|));
	}
	globalState.f0 := p1;
	var v82: array<seq<int>> := new seq<int>[28];
	forall i4 | 0 <= i4 < v82.Length {
		v82[i4] := (v73 + v73) + v73;
	}
	var v83: array<bool> := new bool[27] [false, p0, p2, !p0, p2, p0, p0, p2, p2, p0, p0, p0, p0, p2, p2, p0, p2, fm2(globalState), p0, p2, p0, p0, p2, p0, !p0, p0, p0];
	var v84 := DC12(p2, fm17(p3, p2, false, globalState), p1, v83, p2);
	var v85: map<int, int> := map[p3 := -fm0(globalState)];
	var v86: seq<map<int, int>> := [v85];
	var v87: array<D4> := new D4[1] [v84.(cf27 := v86, cf28 := |v73|, cf29 := v83)];
	v87[safeIndex(156, v87.Length)] := v84;
	v87[safeIndex(156, v87.Length)] := v84;
}
method Main() {
	var v0 := true;
	var v1: set<bool> := {v0, v0, v0};
	var v2: array<seq<bool>> := new seq<bool>[1](i0 => DC0([v0, v0]).cf0);
	var v3 := -730;
	var v4: map<bool, int> := map[v0 := v3];
	var v5: array<map<bool, int>> := new map<bool, int>[21] [v4, map[v0 := v3], v4, v4, v4, map[v0 := v3], v4, v4, v4, v4, map[v0 := v3], v4, v4, v4, v4, v4, v4, v4, v4, v4, map[v0 := v3]];
	var globalState := new GlobalState(0x13d, {v0} - v1, v2, v5, 'u', "tinqkjdyt", 892, 437, true, 'o', false, -0x22, false, false);
	var v6 := "whhg";
	for i1 := |v1| to |v6| {
		var v7: array<int> := new int[15];
		var v8: seq<array<int>> := [v7];
		var v9: map<int, seq<array<int>>> := map[v3 + |v6| := v8];
		v9 := v9[-fm0(globalState) := v8];
		v3 := -v3;
		var v10: seq<int> := [v3, i1];
		if (safeDivisionInt(v3, |(multiset(v10))[v3 := abs(i1)]|) > v3) {
			var v11: seq<bool> := [v0];
			var v12 := DC2(v0, v0, -|v11|);
			v3 := v12.cf4;
			var v13: array<bool> := new bool[24];
			var v14: seq<array<bool>> := [v13];
			v13[safeIndex(49, v13.Length)] := v14 != v14;
			v13[safeIndex(49, v13.Length)] := v0 && v0;
			var v15: map<int, bool> := map[|v6| := v0];
			m0(v3 < i1, |v15|, v13[safeIndex(49, v13.Length)], v3, globalState);
			var v16: multiset<int> := multiset{i1};
			var v17: map<int, int> := map[i1 := safeModuloInt(if (v3 in v16) then v16[v3] else i1, i1)];
			var v18 := 'i';
			v17 := v17[|v6| := |(seq(abs(357), i2  => ('s')))[safeIndex(v3, |(seq(abs(357), i2  => ('s')))|) := v18]|];
			globalState.f0 := safeModuloInt(i1, i1);
		} else {
			v3 := v3;
			m0(!v0, v3, v3 > v3, i1, globalState);
			m0(v0, i1, false, v3 * v3, globalState);
			var v19: map<int, string> := map[v3 := v6];
			var v20: map<int, bool> := map[v3 := v6 < (if (i1 in v19) then v19[i1] else v6)];
			var v21: map<bool, bool> := map[v0 := v0];
			v0 := if (i1 in v20) then v20[i1] else if (v0 in v21) then v21[v0] else v0;
			v20 := v20;
		}
		
		var v22 := DC2(v0, v0, i1);
		var v23: map<bool, bool> := map[v0 := v22.cf2];
		m0(if (v0 in v23) then v23[v0] else v0, v3, v0, v3 + -v3, globalState);
	}
	m0(v3 <= v3, v3 - v3, "bffs" <= v6, v3, globalState);
	for i3 := v3 to -v3 {
		v0 := v0;
		m0(v0, i3 - fm0(globalState), false, v3, globalState);
		var v24: seq<bool> := [v0, v0];
		var v25: map<int, int> := map[i3 := v3];
		var v26: seq<map<int, int>> := [v25];
		var v27 := 'g';
		var v28: seq<int> := [|(fm1(v0, v3, globalState))[safeIndex(|v26|, |fm1(v0, v3, globalState)|) := v27]|];
		var v29: map<int, bool> := map[i3 := !v0];
		m0(v0 in v24, v28[safeIndex(v3, |v28|)], if (i3 in v29) then v29[i3] else v0, |v24[safeIndex(i3, |v24|) := v0]|, globalState);
		var v30: map<bool, bool> := map[v0 := fm2(globalState)];
		v3 := |v30| * i3;
	}
	var v31 := DC1(v3 + v3);
	match (v31) {
		case DC1(cf1) =>
			v0 := fm2(globalState);
			var v32: seq<bool> := [false];
			v32 := v32;
			var v33: map<int, bool> := map[cf1 := v0];
			var v34: seq<int> := [v3];
			v32 := v32 + (v32 + fm3(-cf1, -0x330, cf1, v33[|v34| := v0], globalState));
			var v35 := DC2(v0, v0, v3);
			var v36: array<bool> := new bool[12] [v0, !v0, v35.cf2, v0, v0, v0, v0, v0, false, v0, !!v35.cf2, v0];
			var v37: multiset<array<bool>> := multiset{v36, v36, v36};
			v0 := v0 <== !(v36 in v37);
		case DC2(cf2, cf3, cf4) =>
			var v38: seq<bool> := [v0, true, v0];
			m0(v38 <= v38, v3, cf3 in map[cf3 := -|v6|], v3, globalState);
			var v39 := new C1();
			var v40: array<map<bool, bool>> := new map<bool, bool>[5](i4 => map[v0 := cf3] + map[DC2(v0, cf2, |multiset{|multiset{'y'}|}|).cf3 := cf3]);
			v40[safeIndex(78, v40.Length)] := fm7(true, globalState);
			var v41: map<bool, bool> := map[cf3 := v0];
			v40[safeIndex(78, v40.Length)] := v41 + fm7(v0, globalState);
			var v42: array<bool> := new bool[6](i5 => cf3);
			v42[safeIndex(193, v42.Length)] := cf3;
			v42[safeIndex(903, v42.Length)] := v3 > cf4;
			v42[safeIndex(417, v42.Length)] := cf4 < v3;
			v42[safeIndex(193, v42.Length)], cf2, v42[safeIndex(903, v42.Length)], v42[safeIndex(417, v42.Length)] := cf4 <= cf4, cf3, false, cf3;
		case DC0(cf0) =>
			m0(v0, v3, v0, v3 + v3, globalState);
			var v43: multiset<map<bool, int>> := multiset{v4};
			m0(v0, |v43| - v3, v0, |v6|, globalState);
			var v44: array<bool> := new bool[15] [v0, v0, v0, v0, v0, v0, v0, fm2(globalState), v0, v0, v0, v0, v0, v0, v0];
			var v45: map<array<bool>, int> := map[v44 := v3];
			var v46: map<map<array<bool>, int>, string> := map[v45 + map[v44 := v3] := v6 + v6];
			v46 := v46[v45 + v45 := "jbs"];
			var v47 := 'j';
			globalState.f9 := v47;
	}
	
	v0 := v0;
	var v48 := DC4(423, v0, v0, v0);
	if (match v48 {
		case DC4(cf6, cf7, cf8, cf9) => !({cf6} >= {v3, |v4|})
		case DC3(cf5) => false
	}) {
		var v49: seq<int> := [v3, if (v0 in v4) then v4[v0] else v3, v3, 390, v3];
		globalState.f0, globalState.f0, v3 := safeDivisionInt(v3, |v49|), v3, v3;
		var v50: array<int> := new int[12];
		var v51: seq<array<int>> := [v50, v50, v50];
		v51 := v51;
		v50[safeIndex(338, v50.Length)] := v3;
		var v52: multiset<bool> := multiset{v0};
		var v53: map<int, int> := map[fm0(globalState) := v3];
		v50[safeIndex(338, v50.Length)] := if ((if (v0) then v0 else v0) in v52) then v52[if (v0) then v0 else v0] else |v53|;
		if (fm2(globalState)) {
			var v54: map<int, bool> := map[v3 := v0];
			v54 := fm8(v3, globalState) + fm8(0x1d9, globalState);
			var v55: array<char> := new char[9](i6 => 'w');
			var v56 := 'k';
			v55[safeIndex(726, v55.Length)] := v56;
			v55[safeIndex(726, v55.Length)] := 'g';
			globalState.f0 := safeDivisionInt(safeModuloInt(v50[safeIndex(338, v50.Length)], v3), v50[safeIndex(338, v50.Length)]);
			globalState.f0 := v3;
			var v57: array<seq<map<int, int>>> := new seq<map<int, int>>[2];
			v57[safeIndex(403, v57.Length)] := [v53, v53, v53];
			var v58: seq<map<int, int>> := [v53, v53];
			v57[safeIndex(403, v57.Length)] := v58 + v58;
		} else {
			var v59: map<bool, string> := map[!!v0 := "jlxmbhbd"];
			v59 := v59 + (v59 + map[v0 := v6]);
			var v60: set<int> := {v3};
			v60 := v60;
			var v61: map<D3, bool> := map[DC8(v3, v3, v0) := v0];
			var v63: map<set<bool>, string> := map[fm9(globalState) := v6];
			var v64: set<D1> := {DC3(v63)};
			var v65 := DC8(v50[safeIndex(338, v50.Length)], |v64|, v0);
			var v66: array<map<D3, bool>> := new map<D3, bool>[2] [v61, map v62 : D3 | v62 in [v65] :: (v62) := (v0)];
			v66[safeIndex(714, v66.Length)] := map[v65 := true];
			v66[safeIndex(714, v66.Length)] := v61 + map[v65 := fm2(globalState)];
			v0 := true;
			var v67 := DC2(v0, v0, v50[safeIndex(338, v50.Length)]);
			var v68: array<D0> := new D0[11] [v67, v67, v67, v67, DC2(v0, v0, fm0(globalState)), v67, v67, v67.(cf2 := v0), v67, v67, fm10(v0, "oysattfu", globalState).(cf3 := false)];
			v68[safeIndex(408, v68.Length)] := v67;
			globalState.f0, globalState.f4, v68[safeIndex(408, v68.Length)], v3 := |v1| + v3, 'e', fm10(v0, v6, globalState), -v50[safeIndex(338, v50.Length)];
		}
		
		globalState.f0 := v3;
	} else {
		globalState.f0 := 660;
		var v69: array<bool> := new bool[22];
		v69[safeIndex(779, v69.Length)] := !v0;
		v69[safeIndex(779, v69.Length)] := v3 == v3;
		var v70: seq<string> := [v6];
		v0, v69[safeIndex(779, v69.Length)] := v3 != v3, !(v3 != safeDivisionInt(fm0(globalState), |v70|));
		v0 := v3 <= |(seq(abs(0x1e6), i7  => (map[v3 := |multiset{v0, v69[safeIndex(779, v69.Length)]}|])))|;
		var v71: seq<map<seq<char>, array<bool>>> := [map[v6 := v69]];
		globalState.f0 := safeDivisionInt(-v3, |v71[safeIndex(|(seq(abs(0x240), i8  => ('f')))|, |v71|)]|);
	}
	
	var v72: map<int, bool> := map[v3 := v0];
	v72 := v72 + v72;
	var v73 := DC2(v0, v0, v3);
	v3 := match v73 {
		case DC1(cf1) => cf1
		case DC2(cf2, cf3, cf4) => |multiset{v0}|
		case DC0(cf0) => 667 * v3
	};
	v0 := v0;
	var v74: map<D1, string> := map[v48 := v6];
	var v75 := DC6(v6 + (if (DC4(-|v6|, v0, v0, v0) in v74) then v74[DC4(-|v6|, v0, v0, v0)] else v6), v3, v0, v3, v0);
	v75 := v75;
	var v76: map<int, int> := map[v3 := DC1(v3).cf1];
	var v77: seq<int> := [v3];
	var v78: seq<map<int, int>> := [v76, map[|v6| := |v77|], v76, v76, v76];
	var v79: seq<bool> := [v0];
	var v80 := DC8(v3, v3, false);
	var v81: C0 := new C0(v80);
	var v82: seq<C0> := [v81, v81];
	var v83: multiset<C0> := multiset{v82[safeIndex(v3, |v82|)]};
	var v84: multiset<int> := multiset{v3, |v83|, -v3};
	var v85: array<bool> := new bool[21] [v0, v0, true, v78 < v78, v0, !(v0 in v1), v0, !v0, v0, v0, v0, |multiset{"ryigcb", v6, v6}| == v3, multiset{|v79|} < v84, v0 || v0, v0, v0, v0, v0 || v0, v0, !v0, !v0];
	v85[safeIndex(938, v85.Length)] := v0;
	var v86 := 'd';
	var v87: map<char, int> := map[v86 := -v81.fm6(v0, DC6(v6, v3, if (v3 in v72) then v72[v3] else v0, v3, if (0x21c in v72) then v72[0x21c] else v0).cf12, !v0, globalState)];
	var v88: multiset<bool> := multiset{v0};
	v85[safeIndex(938, v85.Length)] := |"lthcc"| > (if (v86 in v87) then v87[v86] else if (v0 in v88) then v88[v0] else v3);
	m0(v0, v3 * v3, v0, v3, globalState);
	var v89: map<D3, seq<char>> := map[v80 := v6];
	var v90 := DC9(map[v81.f14 := v6]);
	v89 := v90.cf20;
	var v91: multiset<seq<bool>> := multiset{v79};
	globalState.f0, v77, globalState.f0 := |(v91 + multiset{[v0], v79})|, v77, (fm11(v86, globalState)).cf14;
	var i9 := 0;
	while (v0)
		decreases 100 - i9
	{
		if (i9 >= 100) {
			break;
		}
		
		i9 := i9 + 1;
		m0(v85[safeIndex(938, v85.Length)], v3, v85[safeIndex(938, v85.Length)], v3, globalState);
		var v92 := DC0([v85[safeIndex(938, v85.Length)], v0]);
		match (v92) {
			case DC1(cf1) =>
				var v93: array<multiset<int>> := new multiset<int>[21];
				v93[safeIndex(335, v93.Length)] := v84;
				v3, v85[safeIndex(938, v85.Length)], globalState.f9, v93[safeIndex(335, v93.Length)], v82 := cf1, v0, 'o', multiset(v77 + v77) * fm12(v86, v85[safeIndex(938, v85.Length)], cf1, globalState), v82;
				var v94 := new C1();
				var v95 := new C0(v81.f14);
				var v96 := new C0(v81.f14);
			case DC2(cf2, cf3, cf4) =>
				var v97: array<array<bool>> := new array<bool>[1];
				v97[safeIndex(336, v97.Length)] := v85;
				var v98: array<string> := new string[2] [v6, v6];
				v98[safeIndex(580, v98.Length)] := v6 + v6;
				var v99: array<int> := new int[24](i10 => i10 - cf4);
				v99[safeIndex(40, v99.Length)] := -0x159;
				v97[safeIndex(336, v97.Length)], v98[safeIndex(580, v98.Length)], v99[safeIndex(40, v99.Length)], cf2 := v85, v6, cf4, safeModuloInt(v3, v3) > cf4;
				v99[safeIndex(40, v99.Length)], cf3 := fm0(globalState), !DC12(v0, v78, v31.cf1, v85, v0).cf26;
				var v100: array<C0> := new C0[16];
				var v101: map<bool, array<C0>> := map[cf3 := v100];
				v101 := map[v85[safeIndex(938, v85.Length)] := v100];
				v3 := cf4;
			case DC0(cf0) =>
				var v102: array<int> := new int[27];
				var v103: map<bool, array<int>> := map[v0 := v102];
				var v104 := DC7(if (v79[safeIndex(v3, |v79|)] in v103) then v103[v79[safeIndex(v3, |v79|)]] else v102);
				v104 := v104;
				var v105: C1 := new C1();
				var v106: seq<C1> := [v105];
				v85[safeIndex(938, v85.Length)] := safeDivisionInt(|cf0|, v3) < |(v106 + [v105, v105])|;
				globalState.f0 := safeModuloInt(v3, v3);
				globalState.f0 := v3;
		}
		
		var v107: C1 := new C1();
		var v108: seq<C1> := [v107, v107, v107, v107];
		globalState.f0 := v3 + |(v108 + v108)|;
		globalState.f0 := v3 * v3;
	}
	var v109: map<string, bool> := map[v6 := v85[safeIndex(938, v85.Length)]];
	var v110: map<set<bool>, int> := map[v1 := if (v3 in v76) then v76[v3] else v3];
	v109 := v109[v6 := {v0} !in v110];
	print v0, "\n";
	print v1 == {true}, "\n";
	print v2[0] == [true, true], "\n";
	print v3, "\n";
	print v4 == map[true := -730], "\n";
	print v5[0] == map[true := -730], "\n";
	print v5[1] == map[true := -730], "\n";
	print v5[2] == map[true := -730], "\n";
	print v5[3] == map[true := -730], "\n";
	print v5[4] == map[true := -730], "\n";
	print v5[5] == map[true := -730], "\n";
	print v5[6] == map[true := -730], "\n";
	print v5[7] == map[true := -730], "\n";
	print v5[8] == map[true := -730], "\n";
	print v5[9] == map[true := -730], "\n";
	print v5[10] == map[true := -730], "\n";
	print v5[11] == map[true := -730], "\n";
	print v5[12] == map[true := -730], "\n";
	print v5[13] == map[true := -730], "\n";
	print v5[14] == map[true := -730], "\n";
	print v5[15] == map[true := -730], "\n";
	print v5[16] == map[true := -730], "\n";
	print v5[17] == map[true := -730], "\n";
	print v5[18] == map[true := -730], "\n";
	print v5[19] == map[true := -730], "\n";
	print v5[20] == map[true := -730], "\n";
	print globalState.f0, "\n";
	print globalState.f1 == {}, "\n";
	print globalState.f2[0] == [true, true], "\n";
	print globalState.f3[0] == map[true := -730], "\n";
	print globalState.f3[1] == map[true := -730], "\n";
	print globalState.f3[2] == map[true := -730], "\n";
	print globalState.f3[3] == map[true := -730], "\n";
	print globalState.f3[4] == map[true := -730], "\n";
	print globalState.f3[5] == map[true := -730], "\n";
	print globalState.f3[6] == map[true := -730], "\n";
	print globalState.f3[7] == map[true := -730], "\n";
	print globalState.f3[8] == map[true := -730], "\n";
	print globalState.f3[9] == map[true := -730], "\n";
	print globalState.f3[10] == map[true := -730], "\n";
	print globalState.f3[11] == map[true := -730], "\n";
	print globalState.f3[12] == map[true := -730], "\n";
	print globalState.f3[13] == map[true := -730], "\n";
	print globalState.f3[14] == map[true := -730], "\n";
	print globalState.f3[15] == map[true := -730], "\n";
	print globalState.f3[16] == map[true := -730], "\n";
	print globalState.f3[17] == map[true := -730], "\n";
	print globalState.f3[18] == map[true := -730], "\n";
	print globalState.f3[19] == map[true := -730], "\n";
	print globalState.f3[20] == map[true := -730], "\n";
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
	print v6, "\n";
	print v31.cf1, "\n";
	print v48.cf6, "\n";
	print v48.cf7, "\n";
	print v48.cf8, "\n";
	print v48.cf9, "\n";
	print v72 == map[1 := true], "\n";
	print v73.cf2, "\n";
	print v73.cf3, "\n";
	print v73.cf4, "\n";
	print v74 == map[DC4(423, true, true, true) := "whhg"], "\n";
	print v75.cf11, "\n";
	print v75.cf12, "\n";
	print v75.cf13, "\n";
	print v75.cf14, "\n";
	print v75.cf15, "\n";
	print v76 == map[1 := 1], "\n";
	print v77 == [1], "\n";
	print v78 == [map[1 := 1], map[4 := 1], map[1 := 1], map[1 := 1], map[1 := 1]], "\n";
	print v79 == [true], "\n";
	print v80.cf17, "\n";
	print v80.cf18, "\n";
	print v80.cf19, "\n";
	print v81.f14.cf17, "\n";
	print v81.f14.cf18, "\n";
	print v81.f14.cf19, "\n";
	print |v82|, "\n";
	print |v83|, "\n";
	print v84 == multiset{1, 1, -1}, "\n";
	print v85[0], "\n";
	print v85[1], "\n";
	print v85[2], "\n";
	print v85[3], "\n";
	print v85[4], "\n";
	print v85[5], "\n";
	print v85[6], "\n";
	print v85[7], "\n";
	print v85[8], "\n";
	print v85[9], "\n";
	print v85[10], "\n";
	print v85[11], "\n";
	print v85[12], "\n";
	print v85[13], "\n";
	print v85[14], "\n";
	print v85[15], "\n";
	print v85[16], "\n";
	print v85[17], "\n";
	print v85[18], "\n";
	print v85[19], "\n";
	print v85[20], "\n";
	print v86, "\n";
	print v87 == map['d' := 615], "\n";
	print v88 == multiset{true}, "\n";
	print v89 == map[DC8(1, 1, false) := "whhg"], "\n";
	print v90.cf20 == map[DC8(1, 1, false) := "whhg"], "\n";
	print v91 == multiset{[true]}, "\n";
	print i9, "\n";
	print v109 == map["whhg" := false], "\n";
	print v110 == map[{true} := 1], "\n";
}