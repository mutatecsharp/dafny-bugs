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
datatype D0 = DC1(cf1: int, cf2: int, cf3: int, cf4: int, cf5: array<bool>) | DC2(cf6: bool) | DC3(cf7: int, cf8: bool, cf9: int, cf10: int) | DC0(cf0: int)
datatype D1 = DC5(cf12: bool, cf13: int) | DC4(cf11: char)
datatype D2 = DC6(cf14: seq<char>)
datatype D3 = DC8(cf16: seq<int>, cf17: bool, cf18: bool, cf19: string, cf20: bool) | DC7(cf15: map<C0, int>)
datatype D4 = DC10(cf22: char, cf23: array<seq<int>>, cf24: int) | DC9(cf21: C0) | DC11(cf25: D4)
datatype D5 = DC12(cf26: array<set<bool>>)
datatype D6 = DC14(cf28: bool, cf29: char, cf30: array<char>) | DC13(cf27: map<int, array<int>>)
datatype D7 = DC15(cf31: seq<bool>)
datatype D8 = DC17(cf33: bool, cf34: int) | DC18(cf35: seq<char>, cf36: seq<int>, cf37: bool) | DC16(cf32: map<int, map<int, int>>) | DC19(cf38: D8)
class GlobalState {
	const f0 : int
	const f1 : bool
	const f2 : int
	var f3 : string
	const f4 : string
	const f5 : int
	const f6 : int
	const f7 : bool
	const f8 : bool
	const f9 : int
	var f10 : int
	var f11 : int
	var f12 : bool
	const f13 : bool
	const f14 : multiset<set<char>>
	var f15 : char
	const f16 : int
	const f17 : int
	constructor (f0 : int, f1 : bool, f2 : int, f3 : string, f4 : string, f5 : int, f6 : int, f7 : bool, f8 : bool, f9 : int, f10 : int, f11 : int, f12 : bool, f13 : bool, f14 : multiset<set<char>>, f15 : char, f16 : int, f17 : int) {
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
	}
	
}

class C0 {
	constructor () {
	}
	
	function fm6(p0: int, p1: bool, globalState: GlobalState): multiset<int> {
		multiset{-(-0x1ef * |{!false, false}|), if (false) then -0x283 else 0x3e1}
	}
	function fm7(p0: bool, p1: int, p2: map<bool, int>, p3: string, globalState: GlobalState): string {
		"cdmixbs" + ("blliiliu" + "vthendu")
	}
}

class C1 {
	var f18 : int
	constructor (f18 : int) {
		this.f18 := f18;
	}
	
	method m1(p0: int, p1: bool, p2: array<D0>, p3: int, globalState: GlobalState) returns (r0: array<bool>, r1: int) {
		var v0: array<char> := new char[22];
		var v1 := 'a';
		v0[safeIndex(96, v0.Length)] := v1;
		v0[safeIndex(96, v0.Length)] := 'n';
		var v2: array<bool> := new bool[27];
		forall i0 | 0 <= i0 < v2.Length {
			v2[i0] := p1;
		}
		f18 := p0;
		var v3 := m0(p0, globalState);
		for i1 := f18 to fm0(globalState) {
			v2[safeIndex(862, v2.Length)] := true;
			var v4: multiset<int> := multiset{i1};
			v2[safeIndex(862, v2.Length)] := v4 > v4;
			var v5 := new C0();
			var v6 := "bfkwvkpem";
			var v7: map<string, int> := map[v6 := 938];
			v7 := v7;
			v3[safeIndex(941, v3.Length)] := -p3;
			v3[safeIndex(941, v3.Length)] := -(-(f18 + i1) + f18);
		}
		if (fm1(p1, "rrme" < fm8(globalState), p3, globalState)) {
			var v8: map<array<int>, D0> := map[v3 := DC0(p3)];
			v8 := v8;
			var v9: multiset<bool> := multiset{true, p1};
			var v10: map<array<char>, bool> := map[v0 := multiset{p1, p1, p1, p1} > v9];
			var v11: map<int, array<char>> := map[fm0(globalState) := v0];
			v10 := v10[if (p0 in v11) then v11[p0] else v0 := p1];
			globalState.f11 := p0;
			globalState.f12 := p1;
			var v12 := DC0(|(seq(abs(-0x39f), i2  => ('j')))|);
			var v13: set<D0> := {DC0(p3), v12, DC0(537)};
			v2[safeIndex(916, v2.Length)] := v13 > v13;
			globalState.f12, v2[safeIndex(916, v2.Length)] := false, p1;
		} else {
			var v14: seq<char> := [v0[safeIndex(96, v0.Length)], v0[safeIndex(96, v0.Length)]];
			var v15 := DC6(v14);
			var v16: map<D2, bool> := map[v15 := p1];
			v16 := map[DC6(v14) := true];
			var v17: seq<bool> := [p1];
			var v18: map<string, seq<bool>> := map[v14 := v17];
			v18 := map[v14 + v14 := v17];
			var v19: map<array<bool>, bool> := map[v2 := p1];
			var v20 := DC4(v1);
			v19, globalState.f11, globalState.f10, v20 := v19 + v19, -p0, p3, v20;
			r0 := v2;
			var v21: map<bool, bool> := map[p1 := p1];
			var v22: set<map<bool, bool>> := {v21, map[p1 := p1], v21};
			globalState.f12 := v22 !! v22;
		}
		
		r0 := v2;
		r1 := p3;
	}
	method m2(p0: array<bool>, p1: seq<char>, p2: int, globalState: GlobalState) {
		var v0: array<int> := new int[2] [-p2, f18 - -0x3c2];
		var v1 := false;
		v0 := if (v1) then v0 else v0;
		var i0 := 0;
		while (if (true) then f18 >= 0xb5 else p2 >= p2)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v2: array<seq<int>> := new seq<int>[25](i1 => if (v1) then seq(abs(145), i2  => (|multiset([v1, v1, v1])|)) else seq(abs(0x9a), i3  => (p2)));
			var v3: seq<int> := [p2];
			v2[safeIndex(869, v2.Length)] := v3;
			v2[safeIndex(869, v2.Length)] := v3;
			globalState.f12 := if (|p1| >= |p1|) then !v1 else v1;
			var v4: C0 := new C0();
			var v5: map<C0, int> := map[v4 := f18 - p2];
			p0[safeIndex(864, p0.Length)] := !v1;
			var v6 := DC7(v5);
			var v7 := DC1(|(seq(abs(0x38c), i4  => (f18)))|, f18, p2, f18, p0);
			var v8: seq<D0> := [DC1(0xfa, |"mocct"|, -f18, f18, p0)];
			v5, p0[safeIndex(864, p0.Length)], v1, globalState.f10 := v6.cf15, v1, v7 !in v8, |(p1 + (seq(abs(0x19), i5  => ('x'))))|;
			var v9 := DC2(v1);
			var v10: map<bool, D0> := map[p0[safeIndex(864, p0.Length)] := v9];
			var v11: map<map<bool, D0>, int> := map[v10 := f18];
			var v12 := 'c';
			var v13: multiset<int> := multiset{|p1[safeIndex(p2, |p1|) := v12]|};
			var v15: set<map<bool, D0>> := {v10, v10};
			globalState.f12 := v11 == (map[v10 := |v13|] + (map v14 : map<bool, D0> | v14 in v15 :: (v14) := (0x2)));
		}
		var v16: array<D3> := new D3[17](i7 => DC8([0x380, |multiset([|p1[safeIndex(505, |p1|) := 'v']|, p2])|], v1, v1, p1, v1));
		forall i6 | 0 <= i6 < v16.Length {
			v16[i6] := if (false) then DC8(seq(abs(0x3e6), i8  => (|map[v1 := f18]|)), v1, v1, p1, v1) else DC8(seq(abs(0x31), i9  => (p2)), !v1, v1, p1, v1);
		}
		if (false) {
			if (p2 >= p2) {
				var v17 := new C0();
				globalState.f12, globalState.f12, globalState.f10, globalState.f11 := v1, true, p2, p2;
				var v18: map<bool, string> := map[!v1 := p1];
				var v19 := DC6(p1);
				var v20: seq<int> := [f18];
				var v21 := DC8(v20, v1, v1, p1, v1);
				var v22: array<string> := new string[29] ["lglnyxquc", p1, if (v1 in v18) then v18[v1] else p1, seq(abs(-0x9c), i10  => ('u')), p1, p1, (v19.(cf14 := p1)).cf14, p1, seq(abs(0x2e2), i11  => (DC4('l').cf11)), p1, p1, p1 + v21.cf19, p1, "hfqwixbp", p1, p1, p1, p1 + p1, p1, "kyqwj", v21.cf19, p1, "aqevck", p1, p1 + p1, p1, p1, seq(abs(307), i12  => ('c')), "mgyrwllh"];
				v22[safeIndex(767, v22.Length)] := p1;
				v22[safeIndex(767, v22.Length)] := seq(abs(360), i13  => ('k'));
				var v23: seq<C0> := [v17];
				v17 := v23[safeIndex(f18 - -f18, |v23|)];
				var v24: array<seq<array<int>>> := new seq<array<int>>[7];
				var v25: seq<array<int>> := [v0, v0, v0, v0];
				v24[safeIndex(798, v24.Length)] := ([v0, v0, v0])[safeIndex(p2, |[v0, v0, v0]|) := v0] + v25;
				v24[safeIndex(798, v24.Length)] := v25 + [v0, v0];
			} else {
				v0 := v0;
				globalState.f15 := fm9(v1, v1, globalState);
				var v26: array<char> := new char[20](i14 => 'l');
				v26 := v26;
				v1 := if (v1) then v1 else v1;
				var v27 := new C0();
			}
			
			var v28 := 'h';
			var v29: map<string, char> := map[seq(abs(0x15e), i15  => ('f')) := v28];
			v29 := v29[p1 + p1 := v28];
			var v30: C0 := new C0();
			var v31: map<C0, int> := map[v30 := -p2];
			v31 := v31[v30 := f18];
			var v32 := DC1(|p1|, p2, p2, -safeDivisionInt(-f18, p2), p0);
			match (v32) {
				case DC1(cf1, cf2, cf3, cf4, cf5) =>
					globalState.f12 := safeModuloInt(cf4, f18) <= cf3;
					cf5[safeIndex(843, cf5.Length)] := v1;
					var v33: multiset<multiset<int>> := multiset{multiset(seq(abs(-0x134), i16  => (cf1)))};
					var v34: multiset<int> := multiset{0x355, f18};
					var v35: seq<bool> := [v1, v1];
					var v36: map<char, bool> := map[v28 := false];
					var v37: multiset<bool> := multiset{if (v28 in v36) then v36[v28] else v1};
					var v38: map<bool, int> := map[v1 := f18];
					cf5[safeIndex(843, cf5.Length)], v1, globalState.f3, globalState.f10 := v33 == (v33[v34[p2 := abs(|map[!v1 := v1]|)] := abs(p2)] + v33), multiset(v35 + v35) >= v37, v30.fm7(v1, f18, v38, p1, globalState) + p1, -safeDivisionInt(cf4, cf1);
					globalState.f3 := p1;
					var v39 := DC0(|v38|);
					var v40: seq<D0> := [v39];
					v40 := fm10(v1, v28, globalState) + v40;
				case DC2(cf6) =>
					v1, globalState.f11, globalState.f10 := false, fm0(globalState) - f18, safeModuloInt(safeDivisionInt(p2, |p1|), f18);
					var v41: map<bool, bool> := map[cf6 := cf6];
					v41 := v41[v1 := cf6] + v41;
					p0[safeIndex(647, p0.Length)] := cf6;
					p0[safeIndex(647, p0.Length)] := false;
					var v42: seq<bool> := [p0[safeIndex(647, p0.Length)], p0[safeIndex(647, p0.Length)], cf6];
					v42 := v42;
				case DC3(cf7, cf8, cf9, cf10) =>
					p0[safeIndex(301, p0.Length)] := v1;
					p0[safeIndex(301, p0.Length)] := v1 <== v1;
					cf8 := multiset(fm4(globalState)) !! multiset{cf8};
					var v43 := DC9(v30);
					var v44: map<bool, C0> := map[fm1(cf8, cf8, f18, globalState) := v30];
					var v45: array<C0> := new C0[17] [v30, v30, v43.cf21, v30, v30, v30, v30, v30, v30, v30, v30, v30, if (v1) then v30 else v30, v30, v30, v30, if (p0[safeIndex(301, p0.Length)] in v44) then v44[p0[safeIndex(301, p0.Length)]] else v30];
					v45[safeIndex(135, v45.Length)] := v30;
					p0[safeIndex(301, p0.Length)], v45[safeIndex(135, v45.Length)] := cf8 || cf8, v30;
					globalState.f10 := cf7;
				case DC0(cf0) =>
					var v46: map<int, bool> := map[f18 := v1];
					var v48: map<int, map<int, bool>> := map[305 := map v47 : int | (306 <= v47) && (v47 < 0x201) :: (v47 - cf0) := (v1)];
					var v49 := DC3(p2, v1, f18, f18);
					var v50: map<map<int, bool>, D0> := map[v46 + (if (cf0 in v48) then v48[cf0] else v46) := v49];
					v50 := v50[v46 := v49];
					var v51: map<bool, int> := map[v1 := f18];
					var v52 := DC6(p1);
					globalState.f3 := v30.fm7(v1, f18, if (v1) then v51 else map[fm1(v1, true, cf0, globalState) := 175], v52.cf14, globalState);
					globalState.f12 := v1;
					f18 := 0xfc - p2;
			}
			
			globalState.f12 := false;
		} else {
			var v53 := DC5(v1, f18);
			match (v53) {
				case DC5(cf12, cf13) =>
					var v54: map<int, bool> := map[0x210 := cf12];
					var v55: seq<bool> := [cf12, cf12];
					var v56 := 'a';
					var v57: map<bool, map<int, char>> := map[(if (cf13 in v54) then v54[cf13] else cf12) !in v55 := map[p2 := v56]];
					v57 := v57 + v57;
					var v58: array<set<bool>> := new set<bool>[21];
					var v59 := DC12(v58);
					v58, f18 := v59.cf26, p2;
					var v60: map<int, int> := map[f18 := -|fm11(v1, globalState)| + p2];
					v60 := v60[-0x3d7 := p2];
					globalState.f10 := -0x1ca;
				case DC4(cf11) =>
					var v61: C0 := new C0();
					v61 := v61;
					v1 := v1;
					v1 := v1;
					p0[safeIndex(317, p0.Length)] := f18 >= f18;
					p0[safeIndex(317, p0.Length)] := v1;
			}
			
			var v63: set<bool> := {v1};
			var v64: C0 := new C0();
			var v65: seq<int> := [f18, 0x37];
			var v66: seq<int> := [|(map v62 : int | (0x389 <= v62) && (v62 < 0xf3) :: (v62 * p2) := (v1))|, |v63|, |[v64]|, |v65|];
			globalState.f12 := fm1(v1, false, |v66|, globalState);
			globalState.f12 := false;
			var v67: array<array<multiset<int>>> := new array<multiset<int>>[25];
			var v68: array<multiset<int>> := new multiset<int>[2](i17 => multiset{f18});
			v67[safeIndex(589, v67.Length)] := v68;
			var v69: multiset<bool> := multiset{v1};
			var v70: multiset<int> := multiset{f18, p2, -p2, |v69|};
			var v71: seq<multiset<int>> := [v70];
			v67[safeIndex(589, v67.Length)] := new multiset<int>[15] [v71[safeIndex(p2, |v71|)], v64.fm6(|v66|, v1, globalState) - v70, v70, multiset(v66[safeIndex(-p2, |v66|) := 0x2c2] + v66), v70, v70, v70 * v70, v70, v70[p2 := abs(p2)], v70, v70, v70 - v70, v70, multiset{p2, f18}, v70 + (multiset{p2, f18})[p2 := abs(f18)]];
			var v72: seq<array<bool>> := [p0, p0];
			v72 := v72;
		}
		
		if (v1 || (v1 <== v1)) {
			var v73: set<bool> := {true};
			var v74 := 'i';
			var v75: seq<set<bool>> := [v73, fm12(v74, f18, globalState)];
			var v76: map<int, array<int>> := map[|v75[safeIndex(p2, |v75|)]| := v0];
			var v77 := DC13(v76);
			match (DC3(p2, fm1(v1, v1, |v77.cf27|, globalState), p2, p2)) {
				case DC1(cf1, cf2, cf3, cf4, cf5) =>
					v0[safeIndex(463, v0.Length)] := p2;
					var v78: multiset<bool> := multiset{v1, v1};
					var v79: seq<bool> := [v1, v1, v1, v1, v1];
					v0[safeIndex(463, v0.Length)] := (cf2 - (if (v1 in v78) then v78[v1] else p2)) * |fm13(v1, v79, globalState)|;
					var v80 := new C0();
					var v81 := new C0();
					var v82: map<C0, int> := map[v80 := cf2];
					var v83 := DC7(v82);
					v83 := (if (v1) then v83 else v83).(cf15 := v82);
				case DC2(cf6) =>
					var v84: C0 := new C0();
					var v85: map<bool, C0> := map[!cf6 := v84];
					v85 := v85[v1 := v84];
					p0[safeIndex(849, p0.Length)] := cf6;
					var v86: map<bool, bool> := map[cf6 := v1];
					var v87: seq<int> := [p2, |v86|];
					var v88: multiset<int> := multiset{p2};
					p0[safeIndex(849, p0.Length)] := (multiset(v87) * v88) < v88;
					var v89: map<seq<int>, int> := map[seq(abs(0x4f), i18  => (p2)) := f18];
					v89 := v89[v87 := p2 - p2];
					var v90: array<bool> := new bool[18] [p0[safeIndex(849, p0.Length)], !v1, p0[safeIndex(849, p0.Length)], false, p0[safeIndex(849, p0.Length)], false, false, cf6, cf6, p0[safeIndex(849, p0.Length)], p0[safeIndex(849, p0.Length)], cf6, p0[safeIndex(849, p0.Length)], cf6, v1, cf6, cf6, true];
					var v91: multiset<array<bool>> := multiset{v90};
					f18 := |(v87 + [p2, |v91|])|;
				case DC3(cf7, cf8, cf9, cf10) =>
					var v92: seq<bool> := [v1, v1];
					var v93: multiset<bool> := multiset{fm1(v1, cf8, f18, globalState), v1, v1, v1, v1};
					var v94: seq<int> := [p2 + |v92|, if (cf8 in v93) then v93[cf8] else p2, cf10];
					v94 := if (v1) then seq(abs(0x1ad), i19  => (|{cf7}|)) else v94;
					var v95: multiset<char> := multiset{v74};
					var v96: map<int, multiset<char>> := map[safeDivisionInt(fm0(globalState), p2) := v95];
					v96 := v96[cf9 := v95];
					v0 := new int[12];
					globalState.f10 := f18;
				case DC0(cf0) =>
					var v97: map<bool, bool> := map[v1 := v1];
					globalState.f11, globalState.f12, v0 := fm0(globalState), cf0 >= |(v97 + v97)|, v0;
					var v98: multiset<bool> := multiset{v1};
					v98 := (v98 - v98) + fm14(f18, DC5(v1, cf0), globalState);
					globalState.f12 := v1;
					var v99 := new C0();
			}
			
			v0 := v0;
			var v100: set<int> := {p2};
			v100 := v100;
			v1 := p1 == p1;
			var v101: map<bool, int> := map[v1 := f18];
			var v102 := DC3(0x1ad, v1, |v101|, 0x16e);
			var v103: map<string, bool> := map[p1 := v102.cf8];
			v103 := v103;
		} else {
			globalState.f10 := f18;
			var v104: array<char> := new char[11](i20 => 't');
			v104 := v104;
			var v105 := 'l';
			var v106: set<char> := {v105, fm9(v1, !false, globalState), v105, v105};
			var v107 := DC2(p2 >= |v106|);
			match (v107) {
				case DC1(cf1, cf2, cf3, cf4, cf5) =>
					cf5[safeIndex(173, cf5.Length)] := fm1(v1, v1, fm0(globalState), globalState) && v1;
					var v108: seq<int> := [cf1];
					var v109: map<array<int>, bool> := map[v0 := v1];
					var v110: multiset<int> := multiset{0x312};
					cf5[safeIndex(173, cf5.Length)] := (multiset(v108))[cf3 := abs(|v109|)] !! v110;
					var v111: seq<bool> := [cf5[safeIndex(173, cf5.Length)]];
					v111 := v111;
					var v112 := new C0();
					v1 := v111[safeIndex(fm0(globalState), |v111|)];
				case DC2(cf6) =>
					var v113 := DC14(false, v105, v104);
					v113 := v113;
					var v114: array<seq<int>> := new seq<int>[18];
					var v115 := DC10(v105, v114, f18);
					var v116: set<D4> := {v115};
					var v117: map<bool, bool> := map[v116 >= {v115} := v1];
					var v118: array<D0> := new D0[7](i21 => v107);
					v118[safeIndex(94, v118.Length)] := v107;
					v117, globalState.f3, v118[safeIndex(94, v118.Length)], globalState.f12 := v117, [v105], v107, cf6;
					var v119 := new C0();
					var v120 := new C0();
				case DC3(cf7, cf8, cf9, cf10) =>
					var v121 := new C0();
					var v122: map<char, int> := map[v105 := cf7];
					globalState.f11 := |(v122 + v122)|;
					var v123 := new C0();
					var v124: array<C0> := new C0[28];
					v124[safeIndex(371, v124.Length)] := v121;
					v124[safeIndex(371, v124.Length)] := v123;
				case DC0(cf0) =>
					var v125: set<int> := {f18};
					globalState.f12 := {cf0, cf0} <= v125;
					var v126: map<int, int> := map[f18 := p2];
					var v127: seq<bool> := [v1];
					var v131: C0 := new C0();
					var v132: set<C0> := {v131};
					var v133: array<map<int, int>> := new map<int, int>[21] [v126 + v126, map[f18 := f18], v126, if (v127[safeIndex(f18, |v127|)]) then v126 else v126, v126, v126, map v128 : int | (0xf1 <= v128) && (v128 < 378) :: (safeModuloInt(v128, -|v127|)) := (f18), v126, (map v129 : int | v129 in v126 :: (v129 + -p2) := (f18)) + (map v130 : int | (0x2df <= v130) && (v130 < 0x1fb) :: (safeDivisionInt(v130, cf0)) := (|multiset{cf0}|)), fm15(|v132|, v1, cf0, globalState), v126, v126, map[cf0 := 650], v126, v126 + v126, v126, map[f18 := cf0], v126, v126, v126, v126];
					v133 := new map<int, int>[20];
					var v134: array<D4> := new D4[14];
					var v135: array<seq<int>> := new seq<int>[26];
					var v136 := DC10(v105, v135, -f18);
					v134[safeIndex(875, v134.Length)] := v136;
					v134[safeIndex(875, v134.Length)] := v136;
					var v137: multiset<string> := multiset{p1};
					v126 := v126[|v137| := p2];
			}
			
			var v138: map<bool, bool> := map[v1 := v1];
			var v139: array<set<bool>> := new set<bool>[9](i22 => {v1});
			var v140 := DC12(v139);
			match (if (if (v1 in v138) then v138[v1] else v1) then DC12(v139) else v140) {
				case DC12(cf26) =>
					p0[safeIndex(125, p0.Length)] := f18 != p2;
					p0[safeIndex(125, p0.Length)] := v1;
					var v141 := new C0();
					globalState.f11 := f18;
					var v142: seq<bool> := [v1, v1, p0[safeIndex(125, p0.Length)]];
					var v143: seq<array<char>> := [v104];
					var v144 := DC14(false, v105, v143[safeIndex(p2, |v143|)]);
					var v145: map<seq<bool>, bool> := map[v142 + v142 := if (p0[safeIndex(125, p0.Length)]) then true else v144.cf28];
					v145 := v145[[v1] := !v1];
			}
			
			if (v1 ==> !(if (v1) then v1 else v1)) {
				var v146 := new C0();
				var v147 := m0(p2, globalState);
				v0[safeIndex(146, v0.Length)] := if (fm1(!v1, true, f18, globalState)) then f18 else f18;
				v0[safeIndex(146, v0.Length)] := (-f18 - f18) * -(p2 * f18);
				var v148: map<bool, int> := map[v1 := -|"cudbn"|];
				globalState.f3 := v146.fm7(v1, |(seq(abs(0x22a), i23  => (v105)))|, v148, p1, globalState);
				v105, globalState.f10 := v105, f18;
			} else {
				var v149 := new C0();
				var v151: array<seq<int>> := new seq<int>[4](i24 => [f18, |(map v150 : int | (640 <= v150) && (v150 < -60) :: (v150 * p2) := (f18))|]);
				var v152 := DC3(f18, false, f18, |v106|);
				var v153 := DC10(fm9(v1, v1, globalState), v151, |{v152}|);
				v153 := v153;
				var v154: seq<array<int>> := [v0, v0];
				v1 := !fm1(v1, v0 in v154, p2 + p2, globalState);
				globalState.f10 := if (v1) then p2 else f18 - p2;
				globalState.f10 := f18;
			}
			
		}
		
		var v155: map<int, bool> := map[fm0(globalState) := !v1];
		var v156: map<int, int> := map[|(if (v1) then v155 else v155)| := p2];
		v156 := v156[-f18 := -|p1|];
	}
}

function fm0(globalState: GlobalState): int {
	0x211 + (0x40 - |{|multiset{|{|"s"|}|, 0x2b7}|, 0x300}|)
}
function fm1(p0: bool, p1: bool, p2: int, globalState: GlobalState): bool {
	(if (false) then true else true) || true
}
function fm2(globalState: GlobalState): map<map<int, D0>, int> {
	((map v0 : map<int, D0> | v0 in map[map[|{multiset{true, true}, multiset(DC15([true]).cf31), multiset{true}}| := DC2(true)] := |(seq(abs(0x17c), i0  => ('i')))|] :: (v0) := (329)) + map[map[|[true]| := DC2(false)] := |[false, true, false, false, false]|]) + (map[map[-0xde := DC2(true)] := |"swxgyamao"|] + map[map[|multiset{679}| := DC2(true)] := |[0x111, |[|[0x1ff]|]|, |map[127 := !false]|]|])
}
function fm3(p0: int, globalState: GlobalState): D0 {
	DC2(!true <== false)
}
function fm4(globalState: GlobalState): seq<bool> {
	DC15([!!false]).cf31 + [true, true]
}
function fm5(p0: int, globalState: GlobalState): seq<int> {
	[|map[-0x360 := 0x3a9]|] + [|map[[|map[false := false]|] := 0x137]|]
}
function fm8(globalState: GlobalState): string {
	"hxltqisx" + "tbiom"
}
function fm9(p0: bool, p1: bool, globalState: GlobalState): char {
	if (false) then 'h' else 'v'
}
function fm10(p0: bool, p1: char, globalState: GlobalState): seq<D0> {
	[DC0(0x236), DC0(607)] + [DC0(|map[false := 'g']|), DC0(|{true, false, false, false, !false}|)]
}
function fm11(p0: bool, globalState: GlobalState): map<string, int> {
	map v0 : string | v0 in ["bb", "mwmmnrp", "obby", seq(abs(247), i0  => ('o'))] :: (v0) := (0x195 + 0x399)
}
function fm12(p0: char, p1: int, globalState: GlobalState): set<bool> {
	({false} * {true}) + {!true, false, false}
}
function fm13(p0: bool, p1: seq<bool>, globalState: GlobalState): map<int, bool> {
	map[-0x83 * |"bqfasi"| := true ==> !true]
}
function fm14(p0: int, p1: D1, globalState: GlobalState): multiset<bool> {
	multiset{true <==> false}
}
function fm15(p0: int, p1: bool, p2: int, globalState: GlobalState): map<int, int> {
	if (0x213 !in map[0x303 := false]) then map v0 : int | v0 in (seq(abs(0x3ae), i0  => (-0x36a))) :: (v0 * 0x367) := (|(map v1 : int | v1 in DC16(map[0x2f8 := map v2 : int | (0x3d0 <= v2) && (v2 < 0x2e9) :: (safeModuloInt(v2, 0x2c6)) := (|[!false, false]|)]).cf32 :: (v1 - |map[|(set v3 : char | v3 in multiset{'y', 'x'} :: (v3))| := |(seq(abs(228), i1  => (|"mamhl"|)))|]|) := (map[|multiset{false}| := true]))|) else map[0x5f := -232] + map[|{967}| := -0x2b9]
}
function fm16(p0: int, globalState: GlobalState): map<seq<bool>, bool> {
	map v0 : seq<bool> | v0 in {[false], [!false]} :: (v0) := (|{true, true}| == -608)
}
function fm17(p0: bool, globalState: GlobalState): map<bool, char> {
	map[multiset{true} !! multiset{true, false} := 'd']
}
method m0(p0: int, globalState: GlobalState) returns (r0: array<int>) {
	for i0 := |"r"| to p0 {
		var v0: set<int> := {|{i0}|, p0, p0};
		var v1 := false;
		var v2: C0 := new C0();
		var v3: map<C0, int> := map[v2 := 803];
		var v4 := DC7(v3);
		var v5: map<bool, D3> := map[v1 := v4];
		v0, v5 := set v6 : int | v6 in (v0 + v0) :: (v6 + (201 + |multiset(seq(abs(-511), i1  => (-128)))|)), v5;
		globalState.f11 := -i0;
		var v7: array<bool> := new bool[21](i2 => v1);
		var v8 := 't';
		var v9: map<char, array<bool>> := map[v8 := v7];
		var v10 := DC1(0x1bb, 0x13b, p0, p0, if (v8 in v9) then v9[v8] else v7);
		var v11: map<D0, array<bool>> := map[v10 := v7];
		var v12: array<array<bool>> := new array<bool>[17] [v7, v7, v7, v7, v7, if (v10 in v11) then v11[v10] else v7, v7, v7, v7, v7, v7, v7, v7, v7, v7, v10.cf5, v7];
		v12[safeIndex(241, v12.Length)] := v10.cf5;
		v12[safeIndex(241, v12.Length)] := v7;
		var v13: array<int> := new int[29];
		r0 := v13;
	}
	var v14: map<int, bool> := map[p0 := true];
	var v15 := false;
	var v16: seq<bool> := [false, !v15];
	var v17: set<int> := {p0};
	var v18: array<bool> := new bool[16] [v15, false, v15, false, v15, v15, v15, if (p0 in v14) then v14[p0] else v15, !v15, v15, v15, !v15, v15, !v15, !v15, v16[safeIndex(-|v17|, |v16|)]];
	var v19 := DC1(fm0(globalState), |(if (false) then v14 else v14)|, p0, p0, v18);
	match (v19) {
		case DC1(cf1, cf2, cf3, cf4, cf5) =>
			var v20: set<bool> := {true};
			var v21 := new C1(|v20| - cf2);
			var v22: array<int> := new int[26](i3 => i3 * |map[v15 := |v16|]|);
			r0 := v22;
			cf4 := v21.f18;
			globalState.f12 := v15;
		case DC2(cf6) =>
			var v23: C1 := new C1(-998);
			v23 := v23;
			v18 := v18;
			cf6 := if (v15 ==> v15) then cf6 else p0 <= p0;
			var v24 := new C0();
		case DC3(cf7, cf8, cf9, cf10) =>
			var v25 := "xbjalwtvy";
			var v26 := DC0(|v25|);
			v26 := v26;
			globalState.f12 := v25 <= ((fm8(globalState))[safeIndex(cf7, |fm8(globalState)|) := 'l'] + v25);
			var v27: set<bool> := {v15, v16[safeIndex(cf7, |v16|)], cf8, v15, v15};
			var v28 := 's';
			v27 := fm12(v28, p0, globalState);
			globalState.f12 := cf8;
		case DC0(cf0) =>
			v15 := v15;
			var v29 := "dw";
			var v30: seq<string> := ["igmkken", v29];
			if (v16[safeIndex(105, |v16|)] <==> (v29 in v30)) {
				var v31: seq<int> := [p0, cf0];
				var v32: map<bool, seq<int>> := map[v15 := v31];
				var v33 := DC8(v31, true, v15, v29, v15);
				v32 := v32[fm1(v15, v15, p0, globalState) := v33.cf16];
				var v34: multiset<int> := multiset{|(v29 + v29)|, if (v15) then cf0 else p0, p0 * p0, p0, p0};
				var v35: map<bool, bool> := map[fm1(v15, v15, p0, globalState) := cf0 < fm0(globalState)];
				var v36: set<bool> := {v15};
				var v37: multiset<seq<bool>> := multiset{v16};
				var v38: map<bool, multiset<seq<bool>>> := map[v15 := v37];
				var v39: map<set<bool>, bool> := map[v36 - v36 := (if (v15 in v38) then v38[v15] else v37) !! v37];
				var v40 := 'b';
				var v41: set<char> := {v40};
				var v42 := DC2(v15);
				globalState.f12, v34, v35, globalState.f12, globalState.f11 := if (v36 in v39) then v39[v36] else v15, v34, map[v15 := !v15], v41 == ({v40, v40, v40, v40, v40} - v41), if (v42.cf6 && v15) then cf0 else cf0;
				var v43: C0 := new C0();
				var v44 := DC9(v43);
				var v45 := DC11(v44);
				var v46 := DC11(v44);
				var v47 := DC11(v46);
				globalState.f11, v47, v40 := |v29|, v47, v40;
				var v48: array<seq<set<char>>> := new seq<set<char>>[29];
				var v49: seq<set<char>> := [v41];
				v48[safeIndex(291, v48.Length)] := v49;
				var v50: multiset<bool> := multiset{v15};
				v18, v48[safeIndex(291, v48.Length)], v50 := v18, v49, v50 * v50;
				var v51: array<int> := new int[21](i4 => i4 + p0);
				v51[safeIndex(4, v51.Length)] := safeModuloInt(cf0, cf0);
				globalState.f11, cf0, globalState.f11, v51[safeIndex(4, v51.Length)], globalState.f10 := p0, |v16|, cf0, (if (v15) then -722 else cf0) + fm0(globalState), safeDivisionInt(v31[safeIndex(p0, |v31|)], v31[safeIndex(|v29|, |v31|)]);
			} else {
				globalState.f12 := DC8([p0], fm1(v15, !v15, cf0, globalState), v15, (seq(abs(0x110), i5  => ('m')))[safeIndex(cf0, |(seq(abs(0x110), i5  => ('m')))|) := v29[safeIndex(p0, |v29|)]], true).cf20;
				globalState.f12 := v15;
				var v52: map<bool, bool> := map[v15 := cf0 == 0x1eb];
				v52 := v52[v29 != v29 := v15];
				v18 := v18;
				var v53: array<int> := new int[11](i6 => i6 + |map[cf0 := p0]|);
				v53[safeIndex(43, v53.Length)] := p0;
				v53[safeIndex(43, v53.Length)] := p0;
			}
			
			cf0 := p0 - cf0;
			var v54: multiset<bool> := multiset{v15, false};
			globalState.f12 := |v54| == cf0;
	}
	
	if (true && v15) {
		globalState.f11 := p0;
		globalState.f12 := !v15;
		var v55 := 'v';
		var v56: map<bool, char> := map[v15 := v55];
		var v57: map<map<bool, char>, bool> := map[v56 := v15];
		v18[safeIndex(708, v18.Length)] := if (fm17(v15, globalState) in v57) then v57[fm17(v15, globalState)] else v15;
		var v58: multiset<int> := multiset{p0};
		var v59: seq<int> := [p0];
		var v60: map<bool, bool> := map[v15 := v15];
		v18[safeIndex(708, v18.Length)] := (if (p0 in v58) then v58[p0] else p0) <= v59[safeIndex(|v60|, |v59|)];
		var v61: multiset<bool> := multiset{v18[safeIndex(708, v18.Length)]};
		var v62 := "fjuprfrne";
		var v63: set<multiset<bool>> := {v61[true := abs(|v62|)], v61, v61};
		var v64: map<int, int> := map[p0 := |v63|];
		var v65 := new C1(|v64|);
		var v66: array<char> := new char[3](i7 => v55);
		var v67 := DC14(v15, v55, v66);
		match (v67.(cf28 := v15 && v15, cf29 := v55)) {
			case DC14(cf28, cf29, cf30) =>
				globalState.f10 := v65.f18;
				v60 := v60[!v15 := v18[safeIndex(708, v18.Length)]];
				var v68: array<string> := new string[11];
				v68[safeIndex(394, v68.Length)] := v62;
				cf29, v68[safeIndex(394, v68.Length)] := v55, ("jgxwsdkj")[safeIndex(v65.f18, |"jgxwsdkj"|) := cf29];
				var v69: array<int> := new int[28];
				v69[safeIndex(195, v69.Length)] := v65.f18;
				v69[safeIndex(195, v69.Length)] := -p0;
			case DC13(cf27) =>
				globalState.f15 := v55;
				var v70: array<int> := new int[29];
				var v71: array<array<int>> := new array<int>[28] [v70, v70, v70, v70, v70, v70, v70, v70, if (v15) then v70 else v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, v70, v70];
				v71[safeIndex(425, v71.Length)] := v70;
				var v72: C0 := new C0();
				var v73: array<array<map<int, bool>>> := new array<map<int, bool>>[9];
				var v74: array<map<int, bool>> := new map<int, bool>[5](i8 => v14);
				v73[safeIndex(371, v73.Length)] := v74;
				var v75: set<array<int>> := {v70, v70};
				v71[safeIndex(425, v71.Length)], v72, v73[safeIndex(371, v73.Length)] := v70, v72, if (v75 > v75) then v74 else v74;
				var v76: array<map<string, int>> := new map<string, int>[5];
				var v77: map<string, int> := map[v62 := v65.f18];
				v76[safeIndex(313, v76.Length)] := v77;
				v76[safeIndex(313, v76.Length)] := v77;
				var v78: multiset<string> := multiset{v62};
				var v79: map<bool, int> := map[v18[safeIndex(708, v18.Length)] := p0];
				v70[safeIndex(205, v70.Length)] := if (v72.fm7(v15, |[v55, v55]|, v79, v62, globalState) in v78) then v78[v72.fm7(v15, |[v55, v55]|, v79, v62, globalState)] else -|v61|;
				var v81: map<set<int>, int> := map[set v80 : int | (-0x3ba <= v80) && (v80 < 990) :: (v80 * v65.f18) := p0];
				var v83 := DC13(cf27);
				var v84: multiset<D6> := multiset{v83};
				var v85: map<multiset<D6>, bool> := map[v84 := v18[safeIndex(708, v18.Length)]];
				v59, v18[safeIndex(708, v18.Length)], globalState.f10, v70[safeIndex(205, v70.Length)], v62 := v59, !v15, if ((set v82 : int | v82 in v14 :: (safeModuloInt(v82, |map[|[false]| := 420]|))) in v81) then v81[set v82 : int | v82 in v14 :: (safeModuloInt(v82, |map[|[false]| := 420]|))] else p0, v65.f18 - safeModuloInt(|v85|, |{v15}|), v62 + v62;
		}
		
	} else {
		var v86: multiset<bool> := multiset{!v15};
		var v87: map<array<bool>, array<bool>> := map[v18 := v18];
		globalState.f12, v86, globalState.f11, v87 := v15, v86, |(v17 + v17)|, if (p0 > p0) then v87 + v87 else v87;
		var v88: set<bool> := {true};
		globalState.f10 := p0 + |v88|;
		var v89: map<bool, bool> := map[v15 := v15];
		var v90: multiset<map<bool, bool>> := multiset{v89, map[fm1(v15, v15, p0, globalState) := v15]};
		v90 := if (v16[safeIndex(p0, |v16|)]) then v90[v89 := abs(p0)] else v90 + v90;
		var v91: array<array<string>> := new array<string>[19];
		var v92: array<string> := new string[4];
		v91[safeIndex(130, v91.Length)] := v92;
		var v93 := "hpyho";
		var v94 := DC6(v93);
		var v95 := 'c';
		v91[safeIndex(130, v91.Length)] := new string[5] [("lksj" + v93)[safeIndex(p0, |("lksj" + v93)|) := v93[safeIndex(p0, |v93|)]], v94.cf14, v93 + v93, seq(abs(0x74), i9  => ('p')), (v93 + v93)[safeIndex(fm0(globalState), |(v93 + v93)|) := v95]];
		globalState.f12, v95, globalState.f10, globalState.f3 := v15, v95, p0, fm8(globalState) + v93[safeIndex(p0, |v93|) := v95];
	}
	
	var i10 := 0;
	while (!v15)
		decreases 100 - i10
	{
		if (i10 >= 100) {
			break;
		}
		
		i10 := i10 + 1;
		var v96: C0 := new C0();
		var v97: map<C0, bool> := map[v96 := v15];
		var v98: map<bool, map<C0, bool>> := map[v15 := v97];
		v98 := v98[v15 := v97 + v97];
		globalState.f11 := |fm15(p0, v15, 0x31a, globalState)| - p0;
		var v99: seq<int> := [p0, p0];
		var v100: C1 := new C1(0x1c3);
		var v102 := 'u';
		var v103: seq<char> := [v102, v102];
		var v104: map<C1, map<char, int>> := map[v100 := map v101 : char | v101 in v103 :: (v101) := (v100.f18)];
		var v105: array<int> := new int[23] [if (v15) then p0 else 0x383, p0 * p0, 586, 120, p0, p0, p0, p0, p0, p0, fm0(globalState), |v99[safeIndex(p0, |v99|) := -p0]|, p0, p0, |v99[safeIndex(p0, |v99|) := 0x1d8]| * |v104|, v100.f18, v100.f18, p0, |v103[safeIndex(p0, |v103|) := v103[safeIndex(p0, |v103|)]]|, v100.f18, if (fm1(v15, v15, |v103|, globalState)) then fm0(globalState) else v99[safeIndex(v100.f18, |v99|)], v100.f18, p0];
		r0 := v105;
		globalState.f12 := !v15;
	}
	v18 := v18;
	var v106: array<char> := new char[20];
	v106 := v106;
	var v107: array<int> := new int[1] [p0];
	r0 := v107;
}
method Main() {
	var v0 := "yo";
	var v2 := 'k';
	var v3: set<char> := {v2};
	var v4: multiset<set<char>> := multiset{set v1 : char | v1 in ['n'] :: (v1), v3};
	var globalState := new GlobalState(0x86, true, 0x6b, v0, "ci", 380, -147, false, false, 0x2d7, 0x21c, 0x240, false, false, v4, 'c', 0x193, -68);
	var v5 := 0x3d4;
	for i0 := -0x2df to v5 {
		var v6: array<bool> := new bool[7];
		var v7: map<seq<array<bool>>, bool> := map[[v6, v6, v6] := false];
		var v8: seq<array<bool>> := [v6, v6, v6, v6, v6];
		v7 := v7[v8 + v8 := false];
		var v9: array<int> := new int[2] [-0xd, DC0(i0).cf0];
		var v10 := true;
		var v11: seq<bool> := [!v10, v10];
		var v12: map<int, bool> := map[v5 * i0 := v10];
		v9, globalState.f12, globalState.f12, globalState.f12, v5 := if (-0x1b6 > 880) then v9 else v9, !(v0 != v0), if (v10) then v10 else v11[safeIndex(v5, |v11|)], if (716 in v12) then v12[716] else v10, -i0 - (|v0| * --fm0(globalState));
		var v13: multiset<int> := multiset{fm0(globalState)};
		v13 := v13 + v13;
		var v14: multiset<bool> := multiset{v10, true, true};
		var v15: map<bool, int> := map[v10 := |v14|];
		var v16: map<int, map<bool, int>> := map[v5 := v15];
		v16 := v16[i0 := v15];
	}
	var v17 := false;
	var v18: map<bool, int> := map[v17 := v5];
	var v19: map<int, bool> := map[|{v5, v5, |v18|}| := !v17];
	if (if (v5 in v19) then v19[v5] else v5 == v5) {
		var v20: array<bool> := new bool[22];
		var v21 := DC3(v5, v17, v5, |v0|);
		v20[safeIndex(972, v20.Length)] := if (if (v5 in v19) then v19[v5] else v21.cf8) then fm1(v17, v17, v5, globalState) else v17;
		var v22: set<bool> := {v17, !v17};
		v20[safeIndex(972, v20.Length)] := fm1(false, v22 !! v22, safeModuloInt(v5, v5), globalState);
		var v23: array<map<map<int, D0>, int>> := new map<map<int, D0>, int>[21];
		var v25: map<int, int> := map[v5 := v5];
		var v26: map<map<int, D0>, int> := map[map v24 : int | v24 in v25 :: (v24 * v5) := (DC2(v20[safeIndex(972, v20.Length)])) := v5];
		v23[safeIndex(397, v23.Length)] := v26;
		v23[safeIndex(397, v23.Length)] := fm2(globalState);
		match (fm3(|v0|, globalState)) {
			case DC1(cf1, cf2, cf3, cf4, cf5) =>
				var v27 := m0(if (v20[safeIndex(972, v20.Length)]) then cf2 else cf4, globalState);
				globalState.f10 := --cf4;
				var v28: array<set<seq<int>>> := new set<seq<int>>[10];
				var v29: seq<int> := [cf2];
				var v30: set<seq<int>> := {v29};
				v28[safeIndex(2, v28.Length)] := v30;
				v28[safeIndex(2, v28.Length)] := v30;
				v20[safeIndex(972, v20.Length)] := !(v2 in v0);
			case DC2(cf6) =>
				var v31: seq<bool> := [v20[safeIndex(972, v20.Length)]];
				v31 := v31 + v31;
				globalState.f10, globalState.f10 := safeModuloInt(v5, v5), -v5 - v5;
				globalState.f12 := false;
				var v32: array<int> := new int[4](i1 => i1 + v5);
				v32[safeIndex(302, v32.Length)] := v5;
				var v33: map<multiset<bool>, bool> := map[multiset{cf6} := v20[safeIndex(972, v20.Length)]];
				v32[safeIndex(302, v32.Length)], v0 := if (cf6) then v5 else safeModuloInt(v5, |v33|), v0 + (seq(abs(0x191), i2  => (v2)));
			case DC3(cf7, cf8, cf9, cf10) =>
				var v34: seq<map<int, bool>> := [v19, v19[0x3bf := cf8], v19];
				var v35: multiset<char> := multiset{v2};
				var v36: seq<bool> := [cf8];
				var v37: multiset<int> := multiset{v5, cf10};
				var v38: array<int> := new int[29] [safeModuloInt(v5, cf7), -(if (cf8) then v5 else cf10), 0xb, cf7 + cf9, cf9, fm0(globalState) * |v34|, v5, v5, |(seq(abs(0x2b5), i3  => (cf9)))|, cf9 * 0x1a3, 0x1b5, |v35[v2 := abs(v5)]| - cf7, cf7, |(fm4(globalState) + v36)|, cf7, -0x21c, cf10, if (cf10 in v37) then v37[cf10] else cf10, v5, cf10, v5, v5, -(-|v0| - |fm5(cf10, globalState)|), cf10, v5 * -0x212, |[v5, -|v0|]|, cf10, v5, cf10];
				v38[safeIndex(882, v38.Length)] := cf9 + cf7;
				v38[safeIndex(882, v38.Length)] := |v0|;
				var v39 := DC4(v2);
				globalState.f15 := (v39.(cf11 := v2)).cf11;
				var v40: array<D0> := new D0[25];
				var v41: set<array<D0>> := {v40};
				v38[safeIndex(882, v38.Length)] := |(v41 + v41)|;
				cf8 := v20[safeIndex(972, v20.Length)] && v17;
			case DC0(cf0) =>
				globalState.f12 := v20[safeIndex(972, v20.Length)] <== v17;
				globalState.f3 := v0;
				globalState.f15 := v2;
				var v42 := m0(fm0(globalState), globalState);
		}
		
		var v43: multiset<int> := multiset{|v0|};
		v20[safeIndex(972, v20.Length)] := !((v43 + v43) >= v43);
		var v44: map<bool, bool> := map[v17 := !v17];
		var v45 := m0(if (|v44| in v43) then v43[|v44|] else v5, globalState);
	} else {
		var v46 := m0(v5, globalState);
		var v47 := DC3(v5, v17, -0x1ac, v5);
		if (v47.cf8) {
			var v48 := DC6(v0);
			globalState.f3 := v48.cf14;
			v46[safeIndex(421, v46.Length)] := v5;
			v5, v17, v46[safeIndex(421, v46.Length)], globalState.f12 := v5, v17, |v0| - v5, (v0 + v0) < v0;
			var v49 := m0(v46[safeIndex(421, v46.Length)], globalState);
			var v50 := DC2(v17);
			var v51: seq<D0> := [v50.(cf6 := !v17), v50];
			v51 := seq(abs(0x158), i4  => (v50));
			var v52: map<bool, bool> := map[v17 := v17];
			var v53: set<bool> := {if (v17 in v52) then v52[v17] else v17};
			var v54 := m0(if (!!v17) then 0x61 else |v53|, globalState);
		} else {
			globalState.f11, globalState.f3 := v5 + v5, "hfnyb";
			var v55: set<bool> := {v17, v17};
			var v56: map<bool, set<bool>> := map[v17 := v55];
			globalState.f10 := v5 - |v56|;
			var v57 := m0(v5, globalState);
			var v58 := m0(v5, globalState);
			globalState.f12 := v17;
		}
		
		var v59: seq<bool> := [v17, true, v17];
		var v60: array<bool> := new bool[8] [v17, v5 > v5, v17, v17, v17, !(if (v5 in v19) then v19[v5] else v17), !v59[safeIndex(v5, |v59|)], v17];
		v60[safeIndex(694, v60.Length)] := v17;
		var v61: multiset<bool> := multiset{v17};
		v60[safeIndex(694, v60.Length)] := fm1(v17, multiset{|v61|, v5, v5} != (multiset{v5, v5})[v5 := abs(v5)], -safeModuloInt(|"wbxk"|, v5), globalState);
		var v62 := m0(0x5, globalState);
		v60[safeIndex(694, v60.Length)] := (v60[safeIndex(694, v60.Length)] <== v60[safeIndex(694, v60.Length)]) <== !v60[safeIndex(694, v60.Length)];
	}
	
	var v63: array<bool> := new bool[7];
	forall i5 | 0 <= i5 < v63.Length {
		v63[i5] := |(multiset{DC6(v0), DC6(v0), DC6(seq(abs(0x369), i6  => (v2)))} * multiset{DC6(v0)})| !in multiset{v5};
	}
	var v64: set<bool> := {v17, fm1(v17, v17, |"yhaahdov"|, globalState), v17};
	v17 := v64 > v64;
	var v65: map<array<bool>, int> := map[v63 := v5];
	var v66: seq<bool> := [!fm1(!!v17, v17, |v65|, globalState), v17, v17, v17, v17];
	var v67: set<int> := {v5, 706, v5};
	var v68 := DC3(v5, |v66| < v5, v5, |v67| * v5);
	v68 := v68;
	var v69: map<map<bool, int>, bool> := map[v18 := v17];
	for i7 := -|v69| to 0x42 {
		var v70 := new C1(i7);
		globalState.f12 := (v17 ==> v17) <==> v17;
		v70.m2(v63, v0 + v0, |[v70.f18]|, globalState);
		if (v17 ==> v17) {
			v70 := new C1(v70.f18);
			var v71: map<int, C1> := map[0x326 := v70];
			var v72: seq<map<int, C1>> := [v71];
			var v73: seq<map<int, C1>> := [v72[safeIndex(v70.f18, |v72|)]];
			globalState.f12 := v73[safeIndex(i7, |v73|)] == v71;
			var v74: array<multiset<int>> := new multiset<int>[29];
			var v75: map<bool, bool> := map[false := v17];
			var v76: map<int, int> := map[i7 := 0x22c];
			var v77: array<int> := new int[22] [v5, |v75|, i7, v70.f18, i7, v5, i7, --670, v5, 0x129, v70.f18, v70.f18, i7, v70.f18, v70.f18, v70.f18, v70.f18, v5, fm0(globalState), |v76|, i7, v5];
			var v78: set<array<int>> := {v77};
			var v79: map<set<array<int>>, int> := map[v78 + v78 := |v18| * v5];
			globalState.f10, v74, v5, globalState.f12 := if (v78 in v79) then v79[v78] else 0x6e, v74, safeDivisionInt(v70.f18, v5), v17;
			var v80: map<bool, C1> := map[v17 := v70];
			var v81 := new C1(|(v80 + v80)|);
			globalState.f11 := v70.f18;
		} else {
			var v82: C0 := new C0();
			var v83: map<C0, int> := map[v82 := v70.f18];
			var v84 := DC7(v83);
			var v85: map<D3, int> := map[v84 := i7];
			v70.f18 := safeDivisionInt(|(map[v84 := v70.f18] + v85)|, |v0|);
			var v86: multiset<C0> := multiset{v82, v82};
			var v87 := new C1(|v86|);
			v63[safeIndex(948, v63.Length)] := v17;
			v63[safeIndex(948, v63.Length)] := v70.f18 >= v87.f18;
			var v88 := new C0();
			globalState.f10 := -v70.f18;
		}
		
	}
	var v89: seq<int> := [v5, -0xdd, v5, v5, |v0|];
	for i8 := |v89| to v5 {
		var v90 := new C0();
		v17 := v17;
		var v91: map<bool, string> := map[true := v0];
		var v92 := new C1(safeDivisionInt(0x2e5, |v91|));
		var v93: map<bool, array<bool>> := map[v17 := v63];
		v93 := v93;
	}
	var v94: array<string> := new string[7](i9 => v0 + v0[safeIndex(|v0|, |v0|) := v2]);
	v94 := v94;
	globalState.f12 := false;
	var v95: array<array<int>> := new array<int>[23];
	v95 := if (v17) then if (v17) then v95 else v95 else v95;
	globalState.f10 := v5;
	var v96: map<int, int> := map[v5 := v5];
	globalState.f12 := v5 != (-135 + |(fm4(globalState))[safeIndex(|v96|, |fm4(globalState)|) := v17]|);
	v94[safeIndex(551, v94.Length)] := seq(abs(0xee), i10  => (v2));
	var v97: seq<string> := [v0];
	v94[safeIndex(551, v94.Length)] := v97[safeIndex(v5, |v97|)];
	if (if (v17) then v17 else v17) {
		var v98 := m0(v5, globalState);
		globalState.f12 := safeDivisionInt(v5, v5) >= v5;
		globalState.f10 := if (0x1d1 in v96) then v96[0x1d1] else |v89|;
		globalState.f10 := v5;
		var v99: array<seq<int>> := new seq<int>[17];
		v99[safeIndex(514, v99.Length)] := [v5, fm0(globalState)];
		v99[safeIndex(514, v99.Length)] := [-0x2f7] + [v5, v5, 0x36d];
	} else {
		v5 := v5 - |v66|;
		var v100: seq<array<bool>> := [v63, v63, v63];
		var v101: C1 := new C1(|(if (v17) then v100 else v100)|);
		v101 := v101;
		v101 := v101;
		if (!(v5 >= v5)) {
			var v102: array<set<D5>> := new set<D5>[8];
			var v103: array<set<bool>> := new set<bool>[20];
			var v104 := DC12(v103);
			var v105: set<D5> := {v104, v104, v104.(cf26 := v103), v104};
			v102[safeIndex(230, v102.Length)] := if (v17) then v105 else v105;
			v102[safeIndex(230, v102.Length)] := v105;
			var v106 := m0(v5, globalState);
			v17 := v17;
			v106[safeIndex(553, v106.Length)] := v101.f18;
			v106[safeIndex(553, v106.Length)], v63 := v101.f18, v63;
			var v107: C0 := new C0();
			var v108: map<bool, C0> := map[!v17 := v107];
			globalState.f12 := fm1(!(|v108[v17 := v107]| == |v67|), v17, v5, globalState);
		} else {
			v101.m2(v63, v0, safeModuloInt(|multiset{seq(abs(0xed), i11  => (v2))}|, 0xf3), globalState);
			v66 := [fm1(v17, v17, v5, globalState)] + (v66 + v66);
			globalState.f11 := -|v94[safeIndex(551, v94.Length)]|;
			globalState.f12 := -v5 == (|v18| - v5);
			var v109 := new C1(v101.f18);
		}
		
		var v110: seq<C1> := [v101, v101, v101, v101];
		v101 := v110[safeIndex(v5, |v110|)];
	}
	
	var v111: multiset<bool> := multiset{v17, v17, fm1(v17, v17, v5, globalState), v17};
	var v112: map<bool, bool> := map[v17 := v17];
	var v113: map<seq<bool>, bool> := map[v66 := !(if (v17 in v112) then v112[v17] else v17)];
	globalState.f12 := fm16(|v111|, globalState) != v113;
	var v114: array<C0> := new C0[3];
	var v115: C0 := new C0();
	v114[safeIndex(204, v114.Length)] := v115;
	v114[safeIndex(204, v114.Length)] := v115;
	print v0, "\n";
	print v2, "\n";
	print v3 == {'k'}, "\n";
	print v4 == multiset{{'n'}, {'k'}}, "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
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
	print globalState.f14 == multiset{{'n'}, {'k'}}, "\n";
	print globalState.f15, "\n";
	print globalState.f16, "\n";
	print globalState.f17, "\n";
	print v5, "\n";
	print v17, "\n";
	print v18 == map[false := -2161], "\n";
	print v19 == map[2 := true], "\n";
	print v63[0], "\n";
	print v63[1], "\n";
	print v63[2], "\n";
	print v63[3], "\n";
	print v63[4], "\n";
	print v63[5], "\n";
	print v63[6], "\n";
	print v64 == {false, true}, "\n";
	print |v65|, "\n";
	print v66 == [true, false, false, false, false, false, false, false, false, false, false], "\n";
	print v67 == {-2161, 706}, "\n";
	print v68.cf7, "\n";
	print v68.cf8, "\n";
	print v68.cf9, "\n";
	print v68.cf10, "\n";
	print v69 == map[map[false := -2161] := false], "\n";
	print v89 == [1, -221, 1, 1, 403], "\n";
	print v94[0], "\n";
	print v94[1], "\n";
	print v94[2], "\n";
	print v94[3], "\n";
	print v94[4], "\n";
	print v94[5], "\n";
	print v94[6], "\n";
	print v96 == map[1 := 1], "\n";
	print v97 == ["yokkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk"], "\n";
	print v111 == multiset{false, false, false, true}, "\n";
	print v112 == map[false := false], "\n";
	print v113 == map[[true, false, false, false, false, false, false, false, false, false, false] := true], "\n";
}