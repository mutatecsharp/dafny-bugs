datatype D0 = DC0(cf0: bool)
datatype D1 = DC2(cf2: int, cf3: int, cf4: int, cf5: map<int, int>) | DC3(cf6: int) | DC1(cf1: array<int>)
datatype D2 = DC5(cf8: array<int>, cf9: seq<array<bool>>, cf10: int) | DC6 | DC4(cf7: array<seq<D1>>)
datatype D3 = DC8 | DC9(cf12: int) | DC7(cf11: seq<bool>) | DC10(cf13: D3)
datatype D4 = DC12(cf15: array<multiset<bool>>, cf16: array<bool>, cf17: bool) | DC11(cf14: C0) | DC13(cf18: D4)
datatype D5 = DC15(cf20: C0) | DC14(cf19: map<int, bool>) | DC16(cf21: D5)
class GlobalState {
	const f0 : array<int>
	const f1 : seq<bool>
	const f2 : bool
	const f3 : set<int>
	const f4 : array<bool>
	var f5 : array<int>
	const f6 : map<int, char>
	const f7 : bool
	const f8 : bool
	const f9 : array<set<int>>
	const f10 : bool
	const f11 : char
	const f12 : array<bool>
	const f13 : int
	const f14 : bool
	var f15 : int
	const f16 : bool
	const f17 : int
	const f18 : int
	var f19 : int
	var f20 : multiset<map<bool, char>>
	const f21 : bool
	const f22 : array<int>
	const f23 : array<int>
	const f24 : int
	const f25 : int
	constructor (f0 : array<int>, f1 : seq<bool>, f2 : bool, f3 : set<int>, f4 : array<bool>, f5 : array<int>, f6 : map<int, char>, f7 : bool, f8 : bool, f9 : array<set<int>>, f10 : bool, f11 : char, f12 : array<bool>, f13 : int, f14 : bool, f15 : int, f16 : bool, f17 : int, f18 : int, f19 : int, f20 : multiset<map<bool, char>>, f21 : bool, f22 : array<int>, f23 : array<int>, f24 : int, f25 : int) {
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

function fm0(p0: seq<int>, p1: int, p2: int, globalState: GlobalState): bool {
	true
}
function fm2(p0: bool, globalState: GlobalState): seq<int> {
	[|multiset([DC8()] + [DC8(), DC8()])|, |"gyruwxl"|, 123 * 758]
}
function fm3(p0: int, p1: map<int, bool>, p2: int, globalState: GlobalState): int {
	|((map[false := false] + map[false := true]) + (map[!!false := true] + map[false := false]))|
}
function fm4(p0: int, p1: bool, p2: bool, p3: int, globalState: GlobalState): set<char> {
	set v2 : char | v2 in (map['o' := [|map[false := true]|]] + (map v0 : char | v0 in (map v1 : char | v1 in (seq(0x2c5, i0  => ('e'))) :: (v1) := (0x388)) :: (v0) := ([0x24b]))) :: (v2)
}
function fm5(p0: bool, p1: bool, p2: bool, globalState: GlobalState): multiset<int> {
	multiset{-|multiset{!false, false}|} - (if (false) then multiset{|[0x34c]|} else multiset{|map[true := false]|, |(seq(0x76, i0  => ('m')))|})
}
function fm6(globalState: GlobalState): char {
	'q'
}
function fm7(globalState: GlobalState): seq<bool> {
	[!true, false, true] + [true]
}
function fm8(globalState: GlobalState): multiset<string> {
	multiset{"b"} - multiset{"bl", "oj"}
}
class C0 {
	var f26 : bool
	constructor (f26 : bool) {
		this.f26 := f26;
	}
	
	function fm1(p0: bool, globalState: GlobalState): string {
		"yim"
	}
	method m0(p0: bool, p1: int, globalState: GlobalState) returns (r0: bool, r1: seq<multiset<int>>) {
		var v0: seq<int> := [p1, p1];
		var v1 := DC0(p0);
		var v2: seq<seq<int>> := [v0, v0, fm2(v1.cf0, globalState)];
		var v3: map<bool, seq<int>> := map[p1 < p1 := v2[-|(seq(0x214, i0  => ('a')))|]];
		var v4: array<string> := new string[15];
		var v5 := "ylwodeou";
		v4[306] := v5;
		var v6: seq<bool> := [f26, !f26, f26];
		globalState.f15, r0, v3, v4[306] := |(v0 + (v0 + v0))|, v6[p1], v3 + map[f26 := v0], v5;
		r0 := v1.cf0;
		for i1 := -p1 to p1 {
			var v7: array<int> := new int[1];
			v7[318] := i1;
			var v8: map<bool, int> := map[p0 := p1];
			var v9: map<int, bool> := map[i1 := f26];
			var v10: multiset<int> := multiset{i1};
			globalState.f5, r0, v7[318], v5, globalState.f19 := v7, !false, p1 / (p1 / |v8[true := p1]|), seq(848, i2  => ('t')), if ((p1 < fm3(i1, v9, i1, globalState)) in v8) then v8[p1 < fm3(i1, v9, i1, globalState)] else |v10|;
			f26 := f26;
			var v11: map<bool, bool> := map[if (p0) then p0 else p0 := f26];
			v11 := v11[if (p0) then p0 else !f26 := f26];
			var v12: map<int, array<int>> := map[i1 := v7];
			r0, r0, globalState.f15, globalState.f15, f26 := f26, v6[p1 * |v9|], i1 * |v12|, -|v0|, (|map[f26 := f26]| * v7[318]) >= v7[318];
		}
		var v13: map<bool, int> := map[f26 := p1];
		var v14: set<bool> := {p0, f26, p0};
		var v15: map<int, bool> := map[|(seq(0x1c8, i3  => ('j')))| := f26];
		v13 := (v13 + v13) + map[fm0(v0, 0x2b5, |v14|, globalState) := fm3(p1, v15, p1, globalState)];
		globalState.f19 := fm3(|fm4(-p1, f26, f26, |multiset{p1, p1}|, globalState)|, v15, -(p1 + p1), globalState);
		var v16: multiset<bool> := multiset{f26, false, f26, f26, f26};
		var v17: map<int, int> := map[p1 := p1];
		var v18 := DC2(|v4[306]|, p1, -594, v17);
		var v19: array<bool> := new bool[19] [f26, false, f26, p0, false, true, p0, true, f26, p0, f26, true, p0, DC0(p0).cf0, p0, f26, f26, p0, true];
		var v20: map<D1, array<bool>> := map[v18 := v19];
		var v21: map<int, array<bool>> := map[fm3(p1, v15, |v16|, globalState) := if (v18 in v20) then v20[v18] else v19];
		v21 := v21;
		var v22: multiset<int> := multiset{p1};
		var v23 := 'c';
		var v24: map<multiset<int>, char> := map[v22 := v23];
		r0 := (v24 + v24) != v24;
		var v25: seq<multiset<int>> := [v22, v22 - fm5(fm0(v0, p1, p1, globalState), f26, f26, globalState)];
		r1 := v25;
	}
}

method Main() {
	var v0: array<int> := new int[6];
	var v1 := true;
	var v2: seq<bool> := [!v1, v1];
	var v3 := 861;
	var v4: set<int> := {v3, v3};
	var v5 := DC0(false);
	var v6: array<bool> := new bool[12] [true, v5.cf0, false, v2[v3], v1, v1, !v1, v1, v1, v1, !v1, v1];
	var v7 := 'l';
	var v8: map<int, char> := map[v3 := v7];
	var v9: array<set<int>> := new set<int>[5];
	var v10: map<bool, char> := map[v1 := v7];
	var v11: multiset<map<bool, char>> := multiset{v10, v10};
	var v12 := DC1(v0);
	var globalState := new GlobalState(v0, v2 + v2, false, v4, v6, v0, v8, true, false, v9, false, 's', v6, 0xf2, true, 442, false, -101, 92, 0x3e7, v11 + v11, false, v12.cf1, v12.cf1, 0x117, 491);
	if ((v3 / v3) >= |[v1, v1, v1]|) {
		v6 := v6;
		var v13 := "shmbaf";
		var v14: multiset<string> := multiset{v13, v13, v13};
		var v15: map<bool, int> := map[fm0([v3], v3, -|v13|, globalState) := -|v4|];
		v3 := if ((seq(-0x34d, i0  => (v7))) in v14) then v14[seq(-0x34d, i0  => (v7))] else if (v1) then |v15| else -872;
		v1 := v7 !in v13;
		if (v4 > v4) {
			globalState.f15 := v3;
			v1 := v1;
			var v16 := new C0(v1);
			v16.f26 := !(v3 > (v3 / -0x16d));
			var v17, v18 := v16.m0(v16.f26, -|v13|, globalState);
		} else {
			var v19 := new C0(v3 != v3);
			var v20: map<int, int> := map[v3 := -v3];
			var v21 := DC2(v3, -v3, v3, v20 + v20);
			v21 := v21;
			var v22: array<array<int>> := new array<int>[12] [v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0];
			v22[333] := v0;
			v13, globalState.f19, v22[333], v19.f26, globalState.f19 := v13, v3, v0, v19.f26, 850;
			v13 := v13;
			var v23 := DC3(v3);
			v23 := v23;
		}
		
		var v24: C0 := new C0(v1);
		var v25: set<C0> := {v24};
		var v26: multiset<int> := multiset{|v25|, v3, 0x2d1, v3};
		v1 := |v26| >= v3;
	} else {
		v0 := v0;
		v7 := v7;
		var v27 := "u";
		var v28: set<bool> := {v1};
		globalState.f15 := (|v27| * v3) - (if (v1) then v3 else |v28|);
		var v29: array<seq<D1>> := new seq<D1>[3](i1 => [DC2(0x7b, v3, v3, map[v3 := |multiset(v2)|])]);
		var v30 := DC4(v29);
		v29 := v30.cf7;
		var v31: array<map<bool, bool>> := new map<bool, bool>[2];
		var v32: map<bool, bool> := map[v1 := v1];
		v31[434] := v32;
		var v33: multiset<string> := multiset{v27};
		globalState.f19, v27, v31[434] := |(v33 - v33)|, v27, v32;
	}
	
	var v34 := new C0(v1);
	var i2 := 0;
	while (v34.f26)
		decreases 100 - i2
	{
		if (i2 >= 100) {
			break;
		}
		
		i2 := i2 + 1;
		var v35: map<bool, bool> := map[v1 := false];
		var v36: map<bool, map<bool, bool>> := map[v1 := v35];
		var v37: map<map<bool, map<bool, bool>>, bool> := map[v36 := v1];
		var v38: seq<int> := [v3];
		v34.f26 := if (v36 in v37) then v37[v36] else fm0(v38, v3, v3, globalState);
		var v39: array<seq<D1>> := new seq<D1>[12];
		var v40 := DC4(if (v1) then v39 else v39);
		v40 := v40;
		var v41: set<bool> := {v1};
		var v42: array<set<bool>> := new set<bool>[4] [v41 * v41, {v1} * v41, v41, v41];
		v42[335] := v41;
		v42[335] := v41 * {v34.f26, v34.f26};
		if (v1) {
			v35 := v35[v34.f26 := v1];
			var v43: map<array<bool>, int> := map[v6 := v3];
			var v44: map<int, int> := map[v3 := |v38|];
			v43 := v43[v6 := (if (|v38| in v44) then v44[|v38|] else 0x3e7) % v3];
			var v45, v46 := v34.m0(v34.f26, v3, globalState);
			var v47: C0 := new C0(true);
			var v48 := DC3(|(v4 * v4)|);
			var v49: array<array<int>> := new array<int>[15];
			v49[501] := v0;
			var v50: seq<array<bool>> := [v6];
			var v51 := "qfhvebs";
			var v52: map<bool, string> := map[v34.f26 := v51];
			v47, v3, v48, v34.f26, v49[501] := v47, DC5(v0, v50, |v52|).cf10, DC3(v3), v34.f26, v0;
			var v53: array<char> := new char[14] [fm6(globalState), v7, v7, v7, v7, v7, v7, v7, v7, v7, v7, v7, v7, v7];
			v53[923] := v7;
			v53[923] := v7;
		} else {
			globalState.f5 := new int[21];
			globalState.f15 := v3;
			var v54: map<D2, char> := map[DC6() := v7];
			var v55: map<int, map<D2, char>> := map[v3 := v54];
			globalState.f19 := (|v55| - v3) * v3;
			v2 := v2[v3 := v34.f26];
			v7 := v7;
		}
		
	}
	var v56: array<D2> := new D2[7](i3 => DC6());
	var v57: array<array<D2>> := new array<D2>[25] [v56, v56, v56, if (v34.f26) then v56 else v56, v56, if (!v34.f26) then v56 else v56, v56, v56, v56, v56, v56, v56, v56, v56, if (v1) then v56 else v56, v56, v56, v56, v56, v56, v56, v56, v56, v56, v56];
	v57 := if (v34.f26) then v57 else v57;
	var v58 := "xseuq";
	v0[126] := |v58|;
	v0[126] := -v3;
	for i4 := 0xd2 to v3 {
		var v59, v60 := v34.m0(v34.f26, v0[126] * v3, globalState);
		var v61: map<char, set<int>> := map['v' := v4];
		v61 := v61[v7 := v4];
		var v62: seq<string> := [v58];
		if (v58 != (v62[v0[126]] + v58)) {
			globalState.f5 := v0;
			var v63: array<array<D1>> := new array<D1>[8];
			var v64: map<array<array<D1>>, array<int>> := map[v63 := v0];
			v64 := v64[v63 := v0];
			var v65: seq<array<bool>> := [v6];
			var v66 := DC5(v0, v65, -v0[126]);
			var v67: map<bool, D2> := map[v34.f26 := v66.(cf8 := v0, cf10 := 762)];
			v67 := v67;
			var v68, v69 := v34.m0(0x33a < i4, 0x8f, globalState);
			v34.f26 := v34.f26;
		} else {
			var v70: map<bool, bool> := map[v34.f26 || v1 := v58 < v58];
			v70 := v70[v34.f26 := v59 ==> v34.f26];
			v1 := v34.f26 <== v1;
			var v71, v72 := v34.m0(v59, |fm7(globalState)|, globalState);
			var v73: C0 := new C0(v71);
			v73 := v34;
			var v74, v75 := v73.m0(true, i4, globalState);
		}
		
		var v76: seq<seq<bool>> := [v2, v2];
		var v77 := DC7(v2);
		var v78: multiset<int> := multiset{v0[126]};
		var v79: map<multiset<int>, seq<bool>> := map[v78 := [v1]];
		var v80: array<seq<bool>> := new seq<bool>[28] [v2, v2, v2, v2 + v2, [v34.f26, v1, true, v59, !v34.f26], v2, if (true) then v2 else v2, v2 + v76[v3], v2 + v2, v2, [v34.f26] + [v34.f26], v2, ([v1, v34.f26])[|(seq(574, i5  => (v0[126])))| := v34.f26], v77.cf11, v2, v2[-|v34.fm1(v59, globalState)| := v34.f26], v2, v2, if (v1) then v2 else v2, v2, v2, v76[-260], v2 + v2, v2, v2 + v2, if (multiset{v3} in v79) then v79[multiset{v3}] else [v34.f26, v1], v2, v2];
		v80[414] := v2[v3 := v34.f26] + v2;
		v80[414] := [!false, v1, v34.f26];
	}
	if (v1) {
		if (v5.cf0) {
			v34.f26 := v34.f26;
			var v81, v82 := v34.m0(v34.f26, v0[126], globalState);
			globalState.f15 := v0[126];
			globalState.f15 := 0x18e;
			v34.f26 := v81;
		} else {
			var v83, v84 := v34.m0(v34.f26, v3, globalState);
			v6[320] := v7 !in "aqoumw";
			var v85: array<char> := new char[5] [v7, v7, v7, v7, v7];
			var v86: map<char, array<char>> := map[v7 := v85];
			var v87: array<array<char>> := new array<char>[8] [if (v34.f26) then v85 else v85, v85, v85, v85, v85, if (v7 in v86) then v86[v7] else v85, v85, v85];
			var v88 := DC9(v0[126]);
			var v89 := DC10(v88);
			v6[320], v87, v89, v0[126] := v5.cf0, v87, if (!v34.f26) then v89 else v89, v0[126];
			globalState.f19, v6[320] := v0[126], fm0(seq(-0x36f, i6  => (v3)), v0[126] / v0[126], v0[126], globalState);
			v0[126] := v0[126];
			var v90: seq<C0> := [v34, v34];
			var v91: seq<int> := [-|v90|];
			var v92: map<bool, int> := map[v34.f26 := v0[126]];
			var v93: map<int, map<bool, int>> := map[v3 := v92];
			var v94: multiset<bool> := multiset{fm0(v91, -v3, v3, globalState), v34.f26 !in (if (v0[126] in v93) then v93[v0[126]] else v92)};
			var v95: map<int, bool> := map[v0[126] := v6[320]];
			var v96: map<char, multiset<bool>> := map[v7 := multiset{v6[320], v6[320]}];
			v0[126], v7, v34.f26, v94 := fm3(82, v95, v0[126], globalState) + (v3 / v0[126]), v7, v6[320], if (v7 in v96) then v96[v7] else v94;
		}
		
		if (v1) {
			v0[126] := v3;
			var v97: seq<int> := [v0[126]];
			var v98: map<bool, bool> := map[v3 !in v97 := v34.f26];
			v98 := v98[v34.f26 := !v34.f26];
			var v99: set<bool> := {v34.f26};
			v99 := v99;
			var v100: C0 := new C0(false);
			var v101: seq<C0> := [v100, v100];
			var v102 := DC11(v101[v0[126]]);
			v100 := v102.cf14;
			var v103: array<char> := new char[19];
			var v104: map<array<int>, array<char>> := map[v0 := v103];
			var v105: array<array<char>> := new array<char>[14] [v103, v103, v103, v103, v103, v103, v103, v103, v103, if (v0 in v104) then v104[v0] else v103, v103, v103, v103, v103];
			var v106: map<seq<bool>, array<array<char>>> := map[v2 := v105];
			v106 := v106[v2 + v2 := v105];
		} else {
			var v107 := new C0(true);
			v34.f26 := true;
			v7 := v7;
			var v108: map<C0, char> := map[v107 := v7];
			var v109: seq<int> := [|v108|, 0xfe];
			var v110: map<int, bool> := map[-|map[v34.f26 := v34.f26]| := v34.f26];
			var v111: map<map<int, bool>, int> := map[v110 := |v58|];
			v109, globalState.f19 := v109, |v111| + (v0[126] - |multiset{v107.f26}|);
			var v112: seq<string> := [v58];
			var v113: multiset<string> := multiset{"d"};
			var v114: array<multiset<string>> := new multiset<string>[7] [multiset{v58, v58, "cbpcgr"} * multiset(v112), v113, v113 + fm8(globalState), (multiset{v58})["ihs" := v3], fm8(globalState), v113, multiset{v58, v58} + v113];
			v114[223] := v113;
			v114[223] := v113;
		}
		
		v6 := v6;
		globalState.f15 := 906 + 0x2f6;
		var v115: map<D2, string> := map[DC5(v0, [v6, v6, v6, v6], |v2|) := v58];
		v115 := v115;
	} else {
		var v116: seq<array<bool>> := [v6, v6, v6, v6];
		var v117 := DC5(v0, v116, v0[126]);
		var v118: seq<seq<array<bool>>> := [v116];
		v117 := DC5(v0, v118[|[v0, v0, v0]|], v0[126]);
		var v119: seq<int> := [0xc3];
		var v120: map<bool, int> := map[v2[v119[v3]] := v3];
		var v121: multiset<int> := multiset{v0[126]};
		var v122: map<int, int> := map[v0[126] := |v121|];
		var v123: multiset<map<int, int>> := multiset{v122};
		v120 := v120[v34.f26 := |v123|];
		var v124: map<bool, bool> := map[!v1 := v34.f26];
		v124 := v124[v34.f26 := v2[v3]];
		var v125 := new C0(v34.f26 <==> v1);
		var v126: set<bool> := {true};
		v126 := v126 + v126;
	}
	
	var v127: map<array<int>, int> := map[v0 := |v2|];
	v34.f26 := (|map[false := v1]| / v0[126]) != (v0[126] - -|v127|);
	globalState.f19 := v0[126];
	var v128: multiset<bool> := multiset{v34.f26, v1};
	var v129: map<int, bool> := map[|("stmah" + v58)| := !(v128 > v128)];
	v129 := v129[v3 := v1];
	var v130: array<multiset<bool>> := new multiset<bool>[11];
	var v131 := DC12(v130, v6, v1);
	var v132 := DC13(v131);
	v132 := DC13(v131);
	var v133 := new C0(v1);
	v0[126], v2, globalState.f19, v2 := 0x368, v2 + v2, (v3 * 165) * v0[126], [v1, v133.f26, v34.f26];
	if (v34.f26) {
		v34.f26 := v34.f26;
		var v134 := new C0(v34.f26);
		v0[126] := v0[126] % v3;
		var v135: map<bool, bool> := map[v34.f26 := v134.f26];
		var v136, v137 := v134.m0(v1 <==> v2[v3], |v135|, globalState);
		v3 := v3;
	} else {
		var v138, v139 := v133.m0(v34.f26, v3, globalState);
		v0[126] := v3;
		v133 := v133;
		var v140 := DC14(v129);
		var v141: multiset<seq<char>> := multiset{v58};
		var v142: map<string, int> := map[v58[fm3(v3, v140.cf19, v3, globalState) := v7] := |v141|];
		v142 := v142[v58 := v3 + v3];
		if (!v34.f26 <==> !v133.f26) {
			v0[126] := -v3;
			globalState.f5 := v0;
			v7 := if (v34.f26 in v10) then v10[v34.f26] else v7;
			var v143: map<char, array<int>> := map[v7 := v0];
			var v144: array<array<int>> := new array<int>[13] [if (v7 in v143) then v143[v7] else v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0];
			v144 := v144;
			var v145, v146 := v34.m0(v34.f26, v3 % v3, globalState);
		} else {
			v138 := v0[126] > v0[126];
			var v147 := DC6();
			v147 := v147;
			v133.f26 := v34.f26;
			var v148, v149 := v34.m0(v34.f26, v3, globalState);
			var v150: array<string> := new string[15](i7 => v58);
			v150[269] := v58;
			v150[269] := (if (v34.f26) then v58 else "faxxuf") + v58;
		}
		
	}
	
	forall i8 | 0 <= i8 < v9.Length {
		v9[i8] := {v0[126] - v0[126], |([[-379], seq(-899, i9  => (0xd4))] + [[v0[126], v3, v0[126]]])|, 0x3a6, -v3};
	}
	v1 := !v34.f26;
}