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
datatype D0 = DC1(cf1: bool) | DC0(cf0: bool) | DC2(cf2: D0)
datatype D1 = DC4(cf4: int, cf5: int) | DC5(cf6: bool, cf7: bool) | DC3(cf3: int) | DC6(cf8: D1)
datatype D2 = DC8(cf10: int, cf11: int, cf12: int) | DC9(cf13: bool, cf14: int, cf15: char) | DC7(cf9: seq<bool>)
datatype D3 = DC11(cf17: int, cf18: int, cf19: bool, cf20: int) | DC10(cf16: map<bool, bool>)
datatype D4 = DC13(cf22: string, cf23: int, cf24: bool, cf25: array<bool>) | DC14(cf26: bool, cf27: bool) | DC12(cf21: array<bool>) | DC15(cf28: D4)
datatype D5 = DC17(cf30: seq<seq<bool>>, cf31: bool, cf32: multiset<int>) | DC18(cf33: bool) | DC16(cf29: array<int>) | DC19(cf34: D5)
datatype D6 = DC20(cf35: map<int, char>)
datatype D7 = DC22(cf37: array<bool>, cf38: seq<array<D2>>, cf39: seq<array<bool>>, cf40: array<bool>, cf41: D5) | DC23(cf42: bool, cf43: int, cf44: bool) | DC21(cf36: map<map<int, bool>, int>)
datatype D8 = DC25 | DC26(cf46: bool, cf47: map<int, string>) | DC27 | DC24(cf45: set<bool>) | DC28(cf48: D8)
datatype D9 = DC30(cf50: char, cf51: bool, cf52: int, cf53: set<int>, cf54: int) | DC29(cf49: map<int, int>) | DC31(cf55: D9)
datatype D10 = DC32(cf56: seq<int>)
datatype D11 = DC34(cf58: multiset<bool>, cf59: multiset<int>, cf60: array<bool>, cf61: D1) | DC35 | DC33(cf57: set<multiset<bool>>) | DC36(cf62: D11)
datatype D12 = DC38(cf64: int) | DC39(cf65: int) | DC37(cf63: multiset<seq<bool>>) | DC40(cf66: D12)
datatype D13 = DC42(cf68: char, cf69: bool) | DC41(cf67: set<seq<int>>) | DC43(cf70: D13)
datatype D14 = DC45(cf72: map<bool, bool>, cf73: C0, cf74: int) | DC44(cf71: array<string>) | DC46(cf75: D14)
datatype D15 = DC48(cf77: int, cf78: C5, cf79: int) | DC49(cf80: map<array<bool>, char>, cf81: int) | DC47(cf76: map<C3, int>) | DC50(cf82: D15)
datatype D16 = DC52(cf84: int, cf85: bool, cf86: array<array<int>>, cf87: bool, cf88: int) | DC51(cf83: array<char>) | DC53(cf89: D16)
datatype D17 = DC54(cf90: array<array<bool>>)
datatype D18 = DC56(cf92: bool, cf93: bool, cf94: int) | DC57(cf95: C0) | DC55(cf91: C3)
datatype D19 = DC59(cf97: T0, cf98: map<bool, int>) | DC58(cf96: map<bool, int>)
datatype D20 = DC61 | DC60(cf99: C4) | DC62(cf100: D20)
datatype D21 = DC64(cf102: bool, cf103: C3) | DC65(cf104: int, cf105: int, cf106: C1) | DC63(cf101: C6) | DC66(cf107: D21)
datatype D22 = DC68(cf109: bool, cf110: int) | DC69(cf111: bool) | DC70(cf112: bool, cf113: bool, cf114: string, cf115: int, cf116: int) | DC67(cf108: string)
trait T0 {
	var f14 : bool
	const f15 : multiset<bool>
	function fm2(p0: bool, p1: bool, globalState: GlobalState): bool 
	function fm3(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): string 
	method m0(globalState: GlobalState) returns (r0: int, r1: bool, r2: map<map<int, bool>, int>, r3: bool) 
	method m1(p0: set<int>, p1: int, globalState: GlobalState) returns (r0: int, r1: int, r2: bool) 
}

class GlobalState {
	const f0 : int
	const f1 : multiset<int>
	var f2 : array<bool>
	var f3 : int
	const f4 : bool
	const f5 : string
	var f6 : set<bool>
	const f7 : bool
	const f8 : int
	var f9 : bool
	const f10 : int
	var f11 : array<set<int>>
	const f12 : array<string>
	var f13 : set<bool>
	constructor (f0 : int, f1 : multiset<int>, f2 : array<bool>, f3 : int, f4 : bool, f5 : string, f6 : set<bool>, f7 : bool, f8 : int, f9 : bool, f10 : int, f11 : array<set<int>>, f12 : array<string>, f13 : set<bool>) {
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
	const f23 : bool
	const f24 : bool
	constructor (f23 : bool, f24 : bool) {
		this.f23 := f23;
		this.f24 := f24;
	}
	
}

class C1 {
	constructor () {
	}
	
	method m10(p0: array<map<int, char>>, p1: array<int>, p2: bool, globalState: GlobalState) returns (r0: D1, r1: int, r2: bool, r3: int) {
		if (p2) {
			var v0: array<int> := new int[9];
			var v1 := DC16(p1);
			v0 := v1.cf29;
			var v2 := "dhcjdj";
			var v3: seq<string> := [v2];
			var v4 := -0x2e2;
			var v5 := new C0(v3[safeIndex(v4, |v3|)] <= "rylyonx", true);
			var v6 := new C0(!v5.f23, true || v5.f23);
			var v7: map<bool, int> := map[p2 := v4];
			r3 := -v4 - safeDivisionInt(-(if (p2 in v7) then v7[p2] else v4), v4);
			var v8: seq<int> := [v4];
			v8 := v8;
		} else {
			var v9: array<bool> := new bool[7];
			globalState.f2 := v9;
			var v10 := 0x1e9;
			var v11: seq<int> := [v10];
			var v12: map<bool, bool> := map[p2 := p2];
			var v13: multiset<bool> := multiset{p2, p2};
			var v14: array<int> := new int[27] [|v11|, v10, v10 - 620, v10, |v12|, fm0(p2, v10, v10, multiset{v10, v10, v10}, globalState), v10, v10, v10, v10, v10, v10, v10, v10, fm0(p2, v10, v10, multiset{0x226, v10}, globalState), v10, safeModuloInt(v10, v10), |(v13 * v13)|, v10, v10, v10, |v11|, v10, v10, 589, v10, v10];
			v14 := p1;
			var v15: seq<bool> := [p2, p2];
			p1[safeIndex(69, p1.Length)] := |v15|;
			p1[safeIndex(69, p1.Length)] := v10;
			var v16: set<bool> := {p2};
			var v17: multiset<int> := multiset{|v16|};
			var v18 := DC7(v15);
			var v19: seq<D2> := [v18, v18, v18, v18, v18];
			v11 := fm21(DC3(fm0(p2, v10, p1[safeIndex(69, p1.Length)], v17[v10 := abs(v10)], globalState)), v15[safeIndex(p1[safeIndex(69, p1.Length)], |v15|)], v19, globalState);
			var v20: set<seq<int>> := {v11, v11, v11, v11};
			var v21: map<bool, int> := map[!!p2 := |v20|];
			var v22 := "ye";
			var v23: set<int> := {v10};
			var v24: seq<map<int, bool>> := [map[if (0x28b in v17) then v17[0x28b] else |v23| := p2]];
			var v25: array<map<bool, int>> := new map<bool, int>[15] [v21, if (p2) then v21 else v21, v21, v21 + v21, (fm1(if (p2 in v12) then v12[p2] else p2, p1[safeIndex(69, p1.Length)], true, fm22(|v22|, v24, p2, p2, globalState), globalState))[p2 := v10], v21[false := p1[safeIndex(69, p1.Length)]], v21 + v21, v21, v21, v21, v21, v21, v21, v21, v21 + v21];
			v25[safeIndex(374, v25.Length)] := v21 + v21[p2 := p1[safeIndex(69, p1.Length)]];
			var v26 := DC5(true, p2);
			var v27: set<D1> := {DC5(!p2, p2), v26, v26, v26, v26};
			globalState.f3, globalState.f9, v25[safeIndex(374, v25.Length)], v11 := |(v27 + v27)|, p2, map[p2 := p1[safeIndex(69, p1.Length)]] + v21, v11 + (seq(abs(624), i0  => (|map['h' := 'o']|)));
		}
		
		var v28 := 0x10b;
		var v29: seq<int> := [v28, v28];
		var v30 := 'j';
		var v31: map<char, int> := map[v30 := v28];
		var v32: set<seq<int>> := {v29};
		if (fm23(safeDivisionInt(v28, v28), fm23(v28, p2, fm24(v28, p2, p2, v28, globalState), {v29, v29}, globalState), v31 + v31[v30 := v28], v32, globalState)) {
			var v33: seq<bool> := [p2, fm23(-0x1f1, false, v31, v32, globalState)];
			var v34 := DC7(v33);
			v33 := v34.cf9;
			var v35 := "yjf";
			var v36: map<int, int> := map[v28 := |v35|];
			var v37 := DC4(|v36|, 112);
			match (v37) {
				case DC4(cf4, cf5) =>
					var v38 := DC3(cf5);
					var v39: array<bool> := new bool[29] [p2, !!p2, p2, false, fm23(v28, true, v31, {fm21(v38, p2, seq(abs(-0x396), i1  => (v34)), globalState)}, globalState), p2, p2, p2, p2, p2, false, p2, p2, p2, !v33[safeIndex(cf5, |v33|)], false, p2, true, p2, true, p2, p2, p2, v33[safeIndex(cf4, |v33|)], !false, p2, p2, p2, true];
					var v40: map<bool, array<bool>> := map[v33[safeIndex(cf4, |v33|)] := v39];
					var v41: C0 := new C0(p2, !p2);
					var v42: seq<C0> := [v41, v41, v41];
					v39[safeIndex(170, v39.Length)] := v42 == v42;
					v35, v40, v39[safeIndex(170, v39.Length)] := v35, v40, if (cf4 < cf4) then v41.f23 else p2;
					var v43: array<D2> := new D2[28];
					v43[safeIndex(395, v43.Length)] := v34;
					var v44: map<bool, int> := map[v39[safeIndex(170, v39.Length)] := 0x11c];
					v30, v41, v30, cf5, v43[safeIndex(395, v43.Length)] := fm25(cf4, -(cf4 + cf4), v41.f23, v44, globalState), v41, v30, 0x3b4, v34;
					var v45: map<array<int>, bool> := map[p1 := v41.f23];
					v45 := v45[p1 := v41.f23];
					var v46: multiset<int> := multiset{-v28};
					var v47: map<array<int>, set<bool>> := map[p1 := {v39[safeIndex(170, v39.Length)]}];
					v46, globalState.f2, v35, v47, r3 := v46, v39, v35, v47, 79;
				case DC5(cf6, cf7) =>
					var v48 := DC14(cf7, false);
					var v49: map<D4, bool> := map[v48 := if (cf6) then cf7 else true];
					v49 := v49[v48 := v33[safeIndex(v28, |v33|)]];
					var v50: array<array<bool>> := new array<bool>[17];
					var v51: array<bool> := new bool[8];
					v50[safeIndex(641, v50.Length)] := v51;
					v50[safeIndex(641, v50.Length)] := v51;
					v50[safeIndex(641, v50.Length)][safeIndex(855, v50[safeIndex(641, v50.Length)].Length)] := p2;
					v50[safeIndex(641, v50.Length)][safeIndex(855, v50[safeIndex(641, v50.Length)].Length)] := p2;
					var v52: seq<string> := [v35, "tbocxqw", v35];
					var v53: map<seq<int>, string> := map[v29 := v52[safeIndex(|[v28, v28, v28]|, |v52|)]];
					globalState.f9 := v29 !in v53;
				case DC3(cf3) =>
					r3, r3 := v28, safeModuloInt(cf3, 0x144) * |v35|;
					var v54: array<bool> := new bool[24](i2 => p2);
					var v55 := DC13(v35, 0x3a2, p2, v54);
					globalState.f2 := v55.cf25;
					v54[safeIndex(457, v54.Length)] := p2;
					v54[safeIndex(457, v54.Length)] := p2;
					v28 := v28;
				case DC6(cf8) =>
					var v56: map<bool, bool> := map[p2 := true];
					var v57: array<bool> := new bool[27] [p2, p2, p2, p2, p2, p2, true, p2, p2, p2, if (p2 in v56) then v56[p2] else p2, p2, p2, p2, p2, p2, p2, v33[safeIndex(v28, |v33|)], p2, p2, !false, p2, v33[safeIndex(v28, |v33|)], true, p2, p2, false];
					var v58 := DC13(v35, |v35[safeIndex(0x2eb, |v35|) := v30]|, !false, v57);
					var v59 := DC15(v58);
					var v60: seq<D4> := [v59, v59];
					r1 := v28 - |multiset(v60 + v60)|;
					r2 := 434 != 0x16;
					var v61: multiset<int> := multiset{v28};
					var v62 := DC13(v35, |v61|, true, v57);
					globalState.f3 := v62.cf23;
					var v63: map<int, bool> := map[v28 := p2];
					var v64: seq<array<int>> := [p1];
					var v65: map<bool, int> := map[true := v28];
					var v66: map<map<int, bool>, bool> := map[v63 := |v64| != (if (p2 in v65) then v65[p2] else |v65|)];
					v66 := v66;
			}
			
			var v67 := DC5(p2, p2);
			v67 := v67;
			globalState.f3 := v28;
			var v68 := new C0(p2, p2);
		} else {
			var v69: array<char> := new char[18];
			var v70: multiset<array<char>> := multiset{v69, v69, v69, v69, v69};
			v70 := multiset{v69};
			var v71: multiset<int> := multiset{v28};
			var v72: map<bool, seq<char>> := map[p2 := (seq(abs(0x2c3), i3  => (v30)))[safeIndex(fm0(p2, v28, v28, v71, globalState), |(seq(abs(0x2c3), i3  => (v30)))|) := v30]];
			v72, v29 := (v72 + v72) + v72, seq(abs(59), i4  => (DC4(-0x346, v28).cf4));
			var v73: array<seq<seq<char>>> := new seq<seq<char>>[7](i5 => [['h']] + [[v30], [v30, v30]]);
			var v74: seq<char> := [v30, v30];
			var v75: seq<seq<char>> := [v74, [v30, v30], v74];
			v73[safeIndex(920, v73.Length)] := v75;
			v73[safeIndex(920, v73.Length)] := v75 + (v75 + fm26(v28, p2, |v74|, globalState));
			var v76: array<string> := new string[6](i6 => "tulofkoow");
			v76[safeIndex(573, v76.Length)] := (seq(abs(0x258), i7  => (v30))) + v74[safeIndex(v28, |v74|) := v30];
			v76[safeIndex(573, v76.Length)] := fm27(-v28, p2, globalState) + v74;
			var v77: seq<bool> := [true, p2, p2, p2, p2];
			var v78: map<int, int> := map[|v77| := v28];
			var v79: map<map<int, int>, int> := map[v78 := v28];
			globalState.f3 := |(if (p2) then v79 else v79)|;
		}
		
		p1[safeIndex(15, p1.Length)] := |v29|;
		p1[safeIndex(15, p1.Length)] := v28;
		r3 := v28 * p1[safeIndex(15, p1.Length)];
		for i8 := p1[safeIndex(15, p1.Length)] to v28 {
			var v80 := new C0(if (p2) then p2 else false, p2 && false);
			if (v28 <= i8) {
				r2 := v80.f24;
				var v81 := "lpi";
				v81 := v81;
				v81 := v81;
				var v82: map<C0, string> := map[v80 := fm27(-i8, p2, globalState)];
				var v83: array<bool> := new bool[14];
				var v84 := DC13(seq(abs(0x67), i9  => (v30)), 12, v80.f23, v83);
				var v85: seq<string> := [v81, v81];
				var v86: array<string> := new string[25] [v81, if (v80 in v82) then v82[v80] else v81, "pmsifhi" + v81, "sta", v81, v81 + fm27(|v81|, v80.f24, globalState), v81, fm27(v28, v80.f23, globalState), v81, v81, v81, v84.cf22, v81, v81, v81 + (seq(abs(0x1ce), i10  => ('o'))), v85[safeIndex(v28, |v85|)], (v81[safeIndex(i8, |v81|) := v30])[safeIndex(p1[safeIndex(15, p1.Length)], |v81[safeIndex(i8, |v81|) := v30]|) := v30], v81, v81, v84.cf22, v81[safeIndex(|v81|, |v81|) := v30], v81 + v81, v81, if (p2) then v81 else v81, "yhvna"];
				v86 := v86;
				var v87: array<seq<bool>> := new seq<bool>[21](i11 => [v80.f24, true, !!v80.f23]);
				var v88: seq<bool> := [v80.f24, false];
				v87[safeIndex(157, v87.Length)] := v88;
				v87[safeIndex(157, v87.Length)] := v88;
			} else {
				var v89: seq<char> := [v30];
				globalState.f3 := fm0(fm23(i8, v80.f23, v31, {v29}, globalState), |v89| + v28, |{p1[safeIndex(15, p1.Length)]}|, multiset{i8}, globalState);
				globalState.f9 := false;
				globalState.f9 := v28 > v28;
				var v90 := new C0(v80.f23, v80.f24);
				var v91: array<bool> := new bool[11];
				var v92: multiset<array<bool>> := multiset{v91, v91};
				globalState.f9 := !((v92[v91 := abs(911)] - v92) == v92);
			}
			
			r3 := p1[safeIndex(15, p1.Length)];
			var v93 := DC18(v80.f24);
			var v94: map<D5, int> := map[v93 := v28];
			globalState.f9 := v28 > (v28 * |v94|);
		}
		p1[safeIndex(15, p1.Length)] := p1[safeIndex(15, p1.Length)];
		r0 := fm28(p2, globalState);
		r1 := p1[safeIndex(15, p1.Length)];
		var v95: seq<bool> := [!p2];
		var v96: seq<seq<bool>> := [v95, v95[safeIndex(|map[p1[safeIndex(15, p1.Length)] := p1[safeIndex(15, p1.Length)]]|, |v95|) := p2]];
		var v97: multiset<int> := multiset{-v28};
		var v98: seq<D5> := [DC17(v96, p2, v97).(cf31 := p2)];
		r2 := fm23(p1[safeIndex(15, p1.Length)], p2, map[v30 := |v98|] + v31, v32, globalState);
		r3 := v28;
	}
}

class C2 {
	var f22 : map<int, bool>
	constructor (f22 : map<int, bool>) {
		this.f22 := f22;
	}
	
	method m8(p0: array<int>, globalState: GlobalState) returns (r0: char, r1: int, r2: D3) {
		var v0 := true;
		var i0 := 0;
		while (v0)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			globalState.f9 := v0;
			var v1 := 'y';
			r0 := v1;
			var v2 := DC1(v0);
			match (v2) {
				case DC1(cf1) =>
					var v3 := 0x8;
					globalState.f3 := v3;
					var v4: array<seq<D1>> := new seq<D1>[24](i1 => (seq(abs(981), i2  => (DC4(v3, v3)))) + [DC4(|f22|, |map[v3 := |map[cf1 := v0]|]|)]);
					var v5: map<bool, bool> := map[cf1 := cf1];
					var v6: map<char, int> := map[v1 := |v5|];
					var v7 := "pr";
					var v8: seq<int> := [if (v1 in v6) then v6[v1] else v3, v3, |v7|];
					var v9: array<bool> := new bool[17];
					var v10: map<array<bool>, char> := map[v9 := v1];
					var v11: seq<bool> := [v0, cf1];
					var v12: seq<seq<bool>> := [v11, v11];
					var v13: array<int> := new int[8] [|"cudqbqnau"|, -0x62, v3 + v3, v3, |(if (false) then [v3] else v8[safeIndex(v3, |v8|) := |v10|])|, v3 - |v8|, safeDivisionInt(fm0(cf1, v3, 0x1b9, multiset(v8), globalState), -|v12[safeIndex(v3, |v12|)]|), |v8| + 0x22e];
					var v14: array<D0> := new D0[13] [v2, v2, v2, DC1(if (v3 in f22) then f22[v3] else false), v2, v2, v2, v2, v2, v2, v2, v2, v2];
					v14[safeIndex(203, v14.Length)] := v2;
					var v15: seq<array<seq<D1>>> := [v4];
					var v16: map<int, array<int>> := map[v3 := p0];
					var v17: multiset<int> := multiset{|f22|};
					v4, v3, v13, v9, v14[safeIndex(203, v14.Length)] := v15[safeIndex(v3 * |v16|, |v15|)], fm0(false, v3 + 0x166, v3, v17, globalState), v13, DC12(v9).cf21, v2.(cf1 := cf1);
					var v18 := new C1();
					p0[safeIndex(418, p0.Length)] := v3;
					p0[safeIndex(418, p0.Length)] := (v3 * v3) * (v3 - v3);
				case DC0(cf0) =>
					var v19 := DC0(v0);
					v19 := v19;
					var v20: map<bool, string> := map[cf0 || v0 := seq(abs(0x211), i3  => (v1))];
					var v21 := "lksdhhcbj";
					v20 := v20[cf0 := v21];
					var v23 := m9(|(set v22 : int | (0x1f6 <= v22) && (v22 < 693) :: (safeDivisionInt(v22, |v21|)))|, globalState);
					var v24: array<bool> := new bool[3] [cf0, cf0, cf0];
					v24[safeIndex(461, v24.Length)] := v0;
					v24[safeIndex(461, v24.Length)] := v0;
				case DC2(cf2) =>
					var v25 := 0x399;
					r1 := v25;
					p0[safeIndex(63, p0.Length)] := 0x377;
					var v26 := "wfyhouxd";
					p0[safeIndex(63, p0.Length)] := safeDivisionInt(v25, -526 + |v26|);
					var v27: array<int> := new int[14];
					v27 := v27;
					var v28: array<D4> := new D4[14];
					var v29: map<int, array<D4>> := map[v25 := v28];
					var v30: array<array<D4>> := new array<D4>[10] [v28, v28, v28, v28, v28, if (true) then v28 else v28, v28, v28, if (p0[safeIndex(63, p0.Length)] in v29) then v29[p0[safeIndex(63, p0.Length)]] else v28, v28];
					v30[safeIndex(97, v30.Length)] := v28;
					v30[safeIndex(97, v30.Length)] := v28;
			}
			
			globalState.f9 := v0 && v0;
		}
		var v31 := 86;
		var v32: map<bool, bool> := map[v31 != -0x24a := true];
		var v33 := 'f';
		var v34: map<char, int> := map[v33 := v31];
		var v35: set<seq<int>> := {[v31, v31]};
		v32 := v32[v0 := fm23(v31, true, v34, v35, globalState)];
		r0 := v33;
		var v36 := new C1();
		var i4 := 0;
		while (v0)
			decreases 100 - i4
		{
			if (i4 >= 100) {
				break;
			}
			
			i4 := i4 + 1;
			var v37: set<int> := {0x16d, v31, v31};
			var v38: seq<set<int>> := [v37, v37, v37];
			var v39: map<int, set<int>> := map[v31 := v38[safeIndex(v31, |v38|)]];
			v39 := v39[v31 := v37];
			var v40 := "ncnudk";
			var v41: multiset<int> := multiset{|v40|, v31, v31, v31, v31};
			p0[safeIndex(215, p0.Length)] := |v41|;
			p0[safeIndex(215, p0.Length)] := |(seq(abs(0x322), i5  => (v31)))| + fm0(v0, v31, v31, v41, globalState);
			var v42 := new C0(v0, v0);
			var v43: map<int, int> := map[p0[safeIndex(215, p0.Length)] := v31];
			r1 := |[|v43|]| + p0[safeIndex(215, p0.Length)];
		}
		var v44: multiset<bool> := multiset{!v0};
		var v45: seq<int> := [v31, v31, v31, v31, 0x1ee];
		globalState.f3 := safeModuloInt(|v44|, |(if (v0) then v45 else v45)|);
		r0 := v33;
		r1 := v31;
		var v46: seq<map<char, int>> := [v34, v34, v34, map['v' := 0x2bd], v34];
		var v47 := "ymak";
		var v48 := DC10(v32);
		r2 := if (fm23(v31, v0, v46[safeIndex(|v47|, |v46|)], v35, globalState)) then v48 else v48;
	}
	method m9(p0: int, globalState: GlobalState) returns (r0: int) {
		var v0 := true;
		var v1: multiset<int> := multiset{p0, p0, 0x395, p0};
		var v2: seq<multiset<int>> := [v1];
		var v3 := "p";
		var v4 := DC19(DC17(fm29(globalState), v0, v2[safeIndex(-|v3|, |v2|)]));
		var v5: seq<D5> := [v4];
		var v6: seq<bool> := [v0];
		var v7 := DC7(v6);
		var v8: array<int> := new int[17] [p0, p0, -p0, p0, |([v4] + v5)|, p0, p0, -p0, p0, p0, p0 + p0, |v7.cf9| + p0, p0, p0, p0, p0, p0];
		v8[safeIndex(459, v8.Length)] := safeModuloInt(p0, p0);
		v8[safeIndex(459, v8.Length)] := -989;
		var v9: multiset<string> := multiset{"lpyalxhrf", v3};
		globalState.f9 := v6[safeIndex(|(v9 + v9)|, |v6|)];
		var v10 := new C1();
		var v11: array<bool> := new bool[23](i0 => DC18(true).cf33);
		v11[safeIndex(190, v11.Length)] := v0;
		v11[safeIndex(512, v11.Length)] := !v0;
		v11[safeIndex(190, v11.Length)], v0, v8[safeIndex(459, v8.Length)], v11[safeIndex(512, v11.Length)] := v0, !(safeDivisionInt(p0, v8[safeIndex(459, v8.Length)]) <= v8[safeIndex(459, v8.Length)]), safeModuloInt(v8[safeIndex(459, v8.Length)], fm0(v0, 0xb1, v8[safeIndex(459, v8.Length)], v1, globalState)), v0;
		v11[safeIndex(190, v11.Length)], v8[safeIndex(459, v8.Length)] := v11[safeIndex(190, v11.Length)], 0xae;
		var v12 := 'q';
		var v13: map<bool, bool> := map[!v0 := v11[safeIndex(190, v11.Length)]];
		var v14: seq<int> := [p0, |v13|, p0, v8[safeIndex(459, v8.Length)]];
		var v15: set<seq<int>> := {v14};
		var v16: set<bool> := {v0, fm23(v8[safeIndex(459, v8.Length)], false, map[v12 := v8[safeIndex(459, v8.Length)]], v15, globalState)};
		if ((v16 * v16) > v16) {
			if (v11[safeIndex(190, v11.Length)]) {
				var v17: map<bool, int> := map[v0 := p0];
				var v18: map<string, map<bool, int>> := map[v3 := v17];
				var v19 := DC11(p0, |v13|, v0, v14[safeIndex(p0, |v14|)]);
				v18 := v18[if (v0) then "byt" else "e" := fm1(v19.cf19, v8[safeIndex(459, v8.Length)], !v0, fm22(p0, [map[p0 := v0], map[p0 := v11[safeIndex(190, v11.Length)]]], v0, !v11[safeIndex(190, v11.Length)], globalState), globalState)];
				var v20: multiset<map<bool, int>> := multiset{v17[v11[safeIndex(190, v11.Length)] := p0], map[v0 := p0]};
				v8[safeIndex(459, v8.Length)] := |v20|;
				globalState.f3 := p0;
				globalState.f9 := !v0;
				globalState.f3 := fm0(true, 312, p0, v1, globalState);
			} else {
				v11[safeIndex(190, v11.Length)] := v11[safeIndex(190, v11.Length)];
				r0 := v14[safeIndex(0x2d6, |v14|)];
				var v21 := DC5(p0 !in v14, false);
				var v22: seq<seq<bool>> := [v6];
				var v23 := DC17(v22, v11[safeIndex(190, v11.Length)], v1);
				globalState.f9, v21, v8[safeIndex(459, v8.Length)], globalState.f9, v11[safeIndex(190, v11.Length)] := v11[safeIndex(190, v11.Length)], v21, -(safeDivisionInt(v8[safeIndex(459, v8.Length)], p0) - -(|v1| * v8[safeIndex(459, v8.Length)])), true, (v11[safeIndex(190, v11.Length)] || !v11[safeIndex(190, v11.Length)]) ==> v23.cf31;
				globalState.f3 := 0x1f;
				v11[safeIndex(190, v11.Length)] := 0x372 == (p0 * v8[safeIndex(459, v8.Length)]);
			}
			
			v8[safeIndex(459, v8.Length)] := p0 * (v8[safeIndex(459, v8.Length)] * p0);
			if (v11[safeIndex(190, v11.Length)]) {
				var v24: map<int, int> := map[v8[safeIndex(459, v8.Length)] := p0];
				v24 := (map[v8[safeIndex(459, v8.Length)] := v8[safeIndex(459, v8.Length)]])[v8[safeIndex(459, v8.Length)] := -p0] + v24;
				globalState.f9 := p0 >= p0;
				var v25 := DC14(v0, 446 <= p0);
				v25 := v25;
				v0 := v0;
				v8[safeIndex(459, v8.Length)] := -(if (v11[safeIndex(190, v11.Length)]) then v8[safeIndex(459, v8.Length)] else safeDivisionInt(-p0, |v3|));
			} else {
				globalState.f9 := !!v0;
				var v26: array<char> := new char[16](i1 => v12);
				v26[safeIndex(846, v26.Length)] := v12;
				v26[safeIndex(846, v26.Length)] := v12;
				var v27: map<char, int> := map[v26[safeIndex(846, v26.Length)] := 0x76];
				var v28 := DC5(v0, fm23(v8[safeIndex(459, v8.Length)], v11[safeIndex(190, v11.Length)], v27, v15, globalState));
				var v29 := DC6(v28);
				var v30 := DC6(v28);
				var v31: map<int, D1> := map[p0 := v30];
				v31 := v31[p0 := v30];
				globalState.f3 := |map[fm25(fm0(v0, |v1|, v8[safeIndex(459, v8.Length)], v1, globalState), v8[safeIndex(459, v8.Length)], v0, map[fm23(|v1|, v0, v27, v15, globalState) := v8[safeIndex(459, v8.Length)]], globalState) := v6]|;
				globalState.f3 := safeModuloInt(--0xcc, safeDivisionInt(-v8[safeIndex(459, v8.Length)], v8[safeIndex(459, v8.Length)]));
			}
			
			v0 := true;
			var v32 := DC9(v11[safeIndex(190, v11.Length)], p0, v12);
			var v33 := DC13(v3, v8[safeIndex(459, v8.Length)], v11[safeIndex(190, v11.Length)], v11);
			v8[safeIndex(459, v8.Length)], v3, v11[safeIndex(190, v11.Length)], globalState.f3 := v32.cf14, v33.cf22, false, safeModuloInt(p0, p0) + fm0(v11[safeIndex(190, v11.Length)], p0, p0, v1, globalState);
		} else {
			globalState.f3 := 0x343;
			globalState.f3 := --safeModuloInt(safeDivisionInt(v8[safeIndex(459, v8.Length)], v8[safeIndex(459, v8.Length)]), v8[safeIndex(459, v8.Length)]);
			var v34 := DC13(v3, 452, v0, v11);
			var v35: seq<D4> := [v34, v34, v34];
			v0 := v35 <= v35;
			var v36 := DC9(v0, -p0, v12);
			var v37: set<multiset<bool>> := {multiset{!true}, multiset{v11[safeIndex(190, v11.Length)]}};
			var v38: map<char, int> := map[if (false) then v36.cf15 else v12 := v8[safeIndex(459, v8.Length)] - |v37|];
			var v39: multiset<array<int>> := multiset{v8};
			v38 := v38['w' := |v39|];
			if (!(v8[safeIndex(459, v8.Length)] == v8[safeIndex(459, v8.Length)])) {
				globalState.f13 := v16;
				var v40: map<int, char> := map[v8[safeIndex(459, v8.Length)] := v12];
				var v43 := DC20(v40);
				var v44: array<map<int, char>> := new map<int, char>[28] [v40, v40[p0 := 'o'], v40, map v41 : int | v41 in fm30(v11[safeIndex(190, v11.Length)], globalState) :: (safeDivisionInt(v41, |v6|)) := (v12), v40, map v42 : int | (-39 <= v42) && (v42 < -511) :: (safeModuloInt(v42, v8[safeIndex(459, v8.Length)])) := ('v'), v40, fm31(v0, v0, globalState), v40, v40, v40, v40, v40, v40, v40, map[p0 := v12], v40, v40, v43.cf35, v40, v40, fm31(true, v0, globalState), v40, v40, v40, map[v8[safeIndex(459, v8.Length)] := v12], v40, v40];
				var v45 := DC1(v0);
				var v46, v47, v48, v49 := v10.m10(v44, v8, v11[safeIndex(190, v11.Length)] && v45.cf1, globalState);
				var v50: multiset<char> := multiset{v12, v12, v12, v12, v12};
				var v51 := DC14(|fm29(globalState)| != |v50|, v11[safeIndex(190, v11.Length)]);
				var v52: multiset<bool> := multiset{v48, v11[safeIndex(190, v11.Length)], v0, v11[safeIndex(190, v11.Length)], v48};
				v51 := DC14(if (v49 in f22) then f22[v49] else v0, v0).(cf26 := v52 > v52);
				globalState.f3 := v47;
				var v53 := new C0(v0, v0);
			} else {
				v8[safeIndex(459, v8.Length)] := 0x120;
				globalState.f9 := v11[safeIndex(190, v11.Length)];
				globalState.f9 := v0;
				var v54: map<int, string> := map[p0 := seq(abs(-811), i2  => (v12))];
				v3 := v3 + (if (-v8[safeIndex(459, v8.Length)] in v54) then v54[-v8[safeIndex(459, v8.Length)]] else "tecdvlbj");
				var v55: map<int, seq<bool>> := map[327 := v6];
				var v56: multiset<bool> := multiset{v0, v0, v0};
				r0, r0 := |{|v55| <= 970, v11[safeIndex(190, v11.Length)], v0, v11[safeIndex(190, v11.Length)]}|, if (if (v11[safeIndex(190, v11.Length)] in v13) then v13[v11[safeIndex(190, v11.Length)]] else false) then |v9| else if (v0 in v56) then v56[v0] else fm0(v11[safeIndex(190, v11.Length)], p0, p0, v1[554 := abs(|v14|)], globalState);
			}
			
		}
		
		var v58: set<char> := {v12};
		r0 := fm0(!v0, fm0(v11[safeIndex(190, v11.Length)], v8[safeIndex(459, v8.Length)], v8[safeIndex(459, v8.Length)], multiset{v8[safeIndex(459, v8.Length)]}, globalState), |(map v57 : char | v57 in v58 :: (v57) := (v11[safeIndex(190, v11.Length)]))|, if (false) then v1 else v1, globalState);
	}
}

class C3 extends T0 {
	const f20 : int
	const f21 : string
	constructor (f20 : int, f21 : string, f14 : bool, f15 : multiset<bool>) {
		this.f20 := f20;
		this.f21 := f21;
		this.f14 := f14;
		this.f15 := f15;
	}
	
	function fm2(p0: bool, p1: bool, globalState: GlobalState): bool {
		if (f14) then f14 else -0x31 != |f21|
	}
	function fm3(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): string {
		("dowcmd" + (f21[safeIndex(|(seq(abs(0x3b8), i0  => ('j')))|, |f21|) := 'y'] + f21))[safeIndex(f20, |("dowcmd" + (f21[safeIndex(|(seq(abs(0x3b8), i0  => ('j')))|, |f21|) := 'y'] + f21))|) := 'a']
	}
	method m0(globalState: GlobalState) returns (r0: int, r1: bool, r2: map<map<int, bool>, int>, r3: bool) {
		var v0: multiset<map<bool, bool>> := multiset{map[f14 := f14]};
		var v1: map<bool, bool> := map[f14 := f14];
		var v2: map<bool, bool> := map[if (!f14 in v1) then v1[!f14] else f14 := f14];
		var v3: multiset<int> := multiset{f20};
		v0 := v0 - (multiset{v2[false := f14], v2})[v1 := abs(fm0(f14, f20, f20, v3, globalState))];
		var v4 := DC7([f14]);
		var v5: multiset<D2> := multiset{v4};
		v5 := v5;
		globalState.f3 := safeDivisionInt(0x2c9, f20 * 115);
		var v6: seq<bool> := [f14, f14];
		for i0 := f20 to |v6| {
			var v7: array<int> := new int[1];
			v7[safeIndex(290, v7.Length)] := i0;
			var v8: array<bool> := new bool[11];
			var v9 := DC4(i0, |fm3(f20, |"fo"|, f14, f20, globalState)|);
			globalState.f2, globalState.f9, v7[safeIndex(290, v7.Length)], globalState.f3 := v8, !f14, if (fm2(f14, f14, globalState) in f15) then f15[fm2(f14, f14, globalState)] else |(seq(abs(0x36b), i1  => (f21)))|, v9.cf4;
			var v10: array<string> := new string[24];
			v10[safeIndex(632, v10.Length)] := seq(abs(-0x3e1), i2  => ('f'));
			var v11: array<D1> := new D1[8](i3 => v9.(cf4 := -v7[safeIndex(290, v7.Length)]));
			v11[safeIndex(817, v11.Length)] := v9;
			v7[safeIndex(290, v7.Length)], v10[safeIndex(632, v10.Length)], r1, globalState.f3, v11[safeIndex(817, v11.Length)] := i0, f21, (if (true) then false else f14) || f14, i0, v9;
			v7[safeIndex(290, v7.Length)] := 578;
			if (true) {
				var v12: map<int, bool> := map[f20 := f14];
				v12 := map[i0 := true] + (v12 + v12);
				var v13 := DC0(!f14);
				var v14: array<D0> := new D0[13] [v13, v13, fm20(-0x23f, v7[safeIndex(290, v7.Length)], v7[safeIndex(290, v7.Length)], globalState), if (f14) then v13 else v13, fm20(600, v7[safeIndex(290, v7.Length)], v7[safeIndex(290, v7.Length)], globalState), v13, if (f14) then v13 else v13, v13, v13, if (f14) then v13 else v13, v13, v13, v13];
				v14[safeIndex(768, v14.Length)] := v13;
				v14[safeIndex(768, v14.Length)] := v13;
				var v15 := 'e';
				var v16: multiset<char> := multiset{v15, 'p'};
				var v17: seq<int> := [f20];
				globalState.f3, v16, globalState.f3, globalState.f3, globalState.f3 := f20, v16, i0, f20, v17[safeIndex(f20, |v17|)];
				var v18: map<bool, multiset<bool>> := map[f14 := multiset{f14, f14, f14, f14, false}];
				v8[safeIndex(322, v8.Length)] := |(if (!f14 in v18) then v18[!f14] else f15)| != i0;
				v8[safeIndex(322, v8.Length)] := true;
				globalState.f3 := f20 - v7[safeIndex(290, v7.Length)];
			} else {
				var v19 := new C2(map[f20 := f14]);
				var v20: map<bool, int> := map[if (f14) then f14 else f14 := f20];
				v20 := v20[f14 := v7[safeIndex(290, v7.Length)] * i0];
				v7[safeIndex(290, v7.Length)] := safeDivisionInt(safeModuloInt(0x291, v7[safeIndex(290, v7.Length)]), 0x277);
				var v21: map<int, int> := map[v7[safeIndex(290, v7.Length)] := i0];
				var v22: map<array<int>, map<int, int>> := map[v7 := v21];
				var v23: seq<int> := [v7[safeIndex(290, v7.Length)]];
				var v24: array<map<array<int>, map<int, int>>> := new map<array<int>, map<int, int>>[25] [v22, map[v7 := v21] + v22, v22, map[v7 := v21], v22, map[v7 := v21], v22, v22 + v22, v22, v22, v22, v22, v22, v22, v22, v22, v22, v22 + v22, v22[v7 := v21], map[v7 := v21] + v22, v22 + v22, v22 + v22, map[v7 := v21], v22 + v22, map[v7 := map[135 := fm0(f14, i0, -v23[safeIndex(v7[safeIndex(290, v7.Length)], |v23|)], multiset(v23), globalState)]]];
				v24[safeIndex(349, v24.Length)] := v22 + v22;
				v24[safeIndex(349, v24.Length)] := v22[v7 := v21];
				f14 := if (v6[safeIndex(i0, |v6|)]) then f14 else f14;
			}
			
		}
		if (!f14) {
			var v25: seq<int> := [0x12d, f20 * -f20, f20];
			v25 := v25;
			var v26 := DC11(f20, f20 * f20, f14, 0x146 + |fm32(f21, v3, globalState)|);
			v26 := v26;
			var v27: array<int> := new int[8] [898, f20, f20, f20, f20, f20, |v6|, f20];
			var v28: seq<array<int>> := [v27];
			v28 := v28 + v28;
			r3 := (f20 > -f20) && f14;
			globalState.f9 := f14;
		} else {
			var v29: array<array<bool>> := new array<bool>[29];
			var v30: map<string, array<array<bool>>> := map["wnwodjwai" + f21 := v29];
			v30 := v30[f21 := v29];
			var v31 := new C0(f14, f14);
			var v32: array<bool> := new bool[24];
			v32[safeIndex(829, v32.Length)] := f14;
			v32[safeIndex(829, v32.Length)] := v31.f23;
			var v33: array<char> := new char[7];
			var v34 := 'p';
			var v35: map<array<char>, string> := map[v33 := f21[safeIndex(f20, |f21|) := v34]];
			v35 := v35;
			globalState.f3 := 0xff;
		}
		
		globalState.f3 := f20;
		r0 := if (f20 in v3) then v3[f20] else f20;
		r1 := f14;
		var v36: map<int, bool> := map[f20 := f14];
		var v37: map<map<int, bool>, int> := map[v36 := f20];
		r2 := DC21(v37).cf36;
		r3 := true;
	}
	method m1(p0: set<int>, p1: int, globalState: GlobalState) returns (r0: int, r1: int, r2: bool) {
		f14 := f14;
		var v0: array<bool> := new bool[14];
		v0[safeIndex(887, v0.Length)] := f14;
		v0[safeIndex(887, v0.Length)] := f14;
		var v1 := new C0(f14, f14 <==> f14);
		var v2: array<int> := new int[20];
		v2[safeIndex(232, v2.Length)] := |fm27(|p0|, v1.f24, globalState)|;
		v2[safeIndex(232, v2.Length)] := safeDivisionInt(p1, fm0(v0[safeIndex(887, v0.Length)], f20, p1, fm30(false, globalState), globalState));
		forall i0 | 0 <= i0 < v2.Length {
			v2[i0] := i0 + DC11(|map[f20 := -f20]|, |map[p0 := map[p1 := DC9(!v1.f24, p1, 'w')]]|, v1.f24, f20).cf20;
		}
		if (false) {
			r2 := if (v0[safeIndex(887, v0.Length)]) then p1 >= |f21| else f14;
			var v3: array<string> := new string[24];
			v3[safeIndex(644, v3.Length)] := f21;
			v3[safeIndex(644, v3.Length)] := f21;
			var v4: map<string, bool> := map[f21 := v1.f23];
			var v5: seq<bool> := [v1.f23];
			globalState.f6 := {true, if (v1.f23) then if (v3[safeIndex(644, v3.Length)] in v4) then v4[v3[safeIndex(644, v3.Length)]] else v5[safeIndex(p1, |v5|)] else v1.f24, !!false};
			f14 := !fm2(f14, !(v1.f23 <==> v1.f24), globalState);
			var v6 := 'f';
			f14, v6, v0[safeIndex(887, v0.Length)] := v1.f24 && v1.f24, v6, v1.f24;
		} else {
			var v7: array<array<int>> := new array<int>[20];
			v7[safeIndex(283, v7.Length)] := v2;
			v7[safeIndex(283, v7.Length)] := new int[8];
			globalState.f9 := v1.f24;
			var v8: map<bool, bool> := map[v0[safeIndex(887, v0.Length)] := !v0[safeIndex(887, v0.Length)]];
			var v9 := DC10(v8[v0[safeIndex(887, v0.Length)] := v1.f23]);
			var v10: map<D3, bool> := map[v9 := v2[safeIndex(232, v2.Length)] < -|[f14, v1.f24]|];
			v10 := v10[v9 := v1.f23];
			var v11: set<int> := {p1};
			v11 := v11 - (if (v0[safeIndex(887, v0.Length)]) then v11 else v11);
			v0[safeIndex(887, v0.Length)] := v1.f23;
		}
		
		var v12: seq<bool> := [f14, false];
		var v13 := 'h';
		var v14: map<int, char> := map[|v12| := v13];
		r0 := |(f15 - f15)| - |v14|;
		var v15: multiset<int> := multiset{v2[safeIndex(232, v2.Length)], p1};
		r1 := (|{p1}| * |v15|) - safeDivisionInt(f20, 546);
		r2 := !(!false ==> (p1 > |fm27(-|v12|, v0[safeIndex(887, v0.Length)], globalState)|));
	}
}

class C4 extends T0 {
	const f19 : bool
	constructor (f19 : bool, f14 : bool, f15 : multiset<bool>) {
		this.f19 := f19;
		this.f14 := f14;
		this.f15 := f15;
	}
	
	function fm2(p0: bool, p1: bool, globalState: GlobalState): bool {
		f19
	}
	function fm3(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): string {
		((seq(abs(-0xd), i0  => ('i'))) + "xlscb") + ("fu" + (seq(abs(866), i1  => ('x'))))
	}
	function fm17(p0: int, p1: bool, p2: multiset<seq<bool>>, p3: int, globalState: GlobalState): int {
		|(f15 * f15)|
	}
	function fm18(p0: bool, p1: int, p2: string, p3: seq<multiset<bool>>, globalState: GlobalState): map<bool, bool> {
		DC10(map[f14 := f19]).cf16 + map[f19 := f14]
	}
	method m0(globalState: GlobalState) returns (r0: int, r1: bool, r2: map<map<int, bool>, int>, r3: bool) {
		if (!f19) {
			var v0: array<bool> := new bool[10](i0 => false);
			v0[safeIndex(891, v0.Length)] := f19 <== f19;
			var v1 := -0x2b1;
			var v2: seq<bool> := [f14];
			var v3 := 'v';
			var v4: set<char> := {v3};
			var v5: map<char, bool> := map[v3 := f19];
			v0[safeIndex(891, v0.Length)] := (fm19(v1, v2[safeIndex(-v1, |v2|)], globalState) - v4) !! (set v6 : char | v6 in v5 :: (v6));
			globalState.f9 := f14;
			var v7: map<bool, bool> := map[v0[safeIndex(891, v0.Length)] := f14];
			var v8: seq<map<bool, bool>> := [v7];
			globalState.f9 := map[f14 := f14] in v8;
			var v9: seq<multiset<bool>> := [f15];
			var v10: array<int> := new int[13](i1 => i1 * v1);
			var v11: set<array<int>> := {v10, v10, v10};
			var v12: multiset<int> := multiset{v1};
			m7(v9, {v10} !! v11, if (f19) then 0x304 else v1, v1 >= fm0(f14, -0x3bb, v1, v12, globalState), globalState);
			var v13: map<int, char> := map[v1 := v3];
			var v14 := DC11(v1, v1, f19, |v13|);
			globalState.f3 := fm0(v1 >= |f15[f19 := abs(v1)]|, v1, v14.cf20, v12, globalState);
		} else {
			var v15 := DC5(f14, f14);
			var v16: array<D1> := new D1[18] [v15, v15, v15, v15, v15, v15, v15, v15, DC5(false, true), v15, v15, v15, v15, v15, v15, v15, DC5(f14, f14), v15];
			var v17: seq<array<D1>> := [v16, v16, v16];
			var v18 := -0x26a;
			var v19: array<seq<array<D1>>> := new seq<array<D1>>[12] [v17, v17 + v17, (v17 + v17)[safeIndex(v18, |(v17 + v17)|) := v16], v17, v17, v17, v17, [v16, v16], v17, v17 + v17, v17, v17];
			v19[safeIndex(930, v19.Length)] := v17;
			v19[safeIndex(930, v19.Length)] := (v17 + v17) + (v17 + v17);
			globalState.f9 := v18 <= -v18;
			var v20: array<bool> := new bool[26];
			var v21: array<array<bool>> := new array<bool>[8] [v20, v20, v20, v20, v20, v20, if (true) then v20 else v20, v20];
			v21[safeIndex(371, v21.Length)] := v20;
			v21[safeIndex(371, v21.Length)] := new bool[7];
			var v22 := "gkgn";
			var v23: T0 := new C3(v18, v22, !f19, fm33(f19, globalState));
			var v24: map<int, T0> := map[v18 + |[v18]| := v23];
			var v25: map<int, int> := map[v18 := -0x161];
			var v26: multiset<int> := multiset{|v25|};
			var v27 := DC4(v18, fm0(v23.f14, -|v22|, 548, v26, globalState));
			v24 := v24[v27.cf5 := v23];
			v18 := safeModuloInt(v18, v18 * v18);
		}
		
		var v28 := -847;
		var v29: seq<int> := [v28];
		var v30: multiset<int> := multiset{v29[safeIndex(|(seq(abs(-0x261), i2  => ('b')))|, |v29|)]};
		var v31: seq<int> := [v28, v28, |v30|];
		v29 := v31;
		var v32 := "hsfh";
		f14 := f14 <== ("eduiltxa" <= v32);
		var v33: map<int, map<bool, int>> := map[v28 := map[f14 := |[f14, false, f14]|]];
		var v34: map<map<int, map<bool, int>>, bool> := map[v33 := f19];
		var v35: seq<string> := [seq(abs(-0xc9), i4  => ('y'))];
		var v36: array<bool> := new bool[22] [if (v33 in v34) then v34[v33] else f14, 0x179 <= v28, f19, v28 <= -v28, f19, fm2(f19, f14, globalState), !f19, f19, f14, v32 != v35[safeIndex(v28, |v35|)], fm2(f19, f19, globalState), !f19, f19 <==> f14, f14 ==> f14, false, true || false, if (f14) then f14 else f19, f14, false, f14, multiset{v32} !! multiset{"lpfkgsqvw"}, !f14];
		forall i3 | 0 <= i3 < v36.Length {
			v36[i3] := ({false, f14} - {true, f14}) >= {f14, !true, f14, f14};
		}
		var v37 := new C1();
		var v38: map<bool, bool> := map[f19 := f19];
		var v39: set<int> := {v28, 0x260, |v38|, v28};
		r1 := !(v39 >= v39);
		r0 := |v32|;
		var v40 := 's';
		var v41: map<char, int> := map[v40 := v28];
		var v42: set<seq<int>> := {v29, v31};
		r1 := fm23(v28 * v28, f19, v41 + map[v40 := v28], v42, globalState);
		var v43: map<int, bool> := map[v28 := f14];
		r2 := map[v43 + map[v28 := fm2(f14, f14, globalState)] := v28];
		r3 := v28 == v28;
	}
	method m1(p0: set<int>, p1: int, globalState: GlobalState) returns (r0: int, r1: int, r2: bool) {
		r1 := p1;
		if (f19 <== !(f14 <== f19)) {
			var v0: array<string> := new string[3](i0 => "mpehub");
			var v2: map<int, int> := map[p1 := p1];
			r1, v0 := |(map v1 : int | v1 in v2 :: (v1 - |multiset([DC9(f14, p1, 'q').cf13])|) := (f14))|, v0;
			r1 := p1;
			r2 := true;
			var v3 := "iycqg";
			var v4: T0 := new C3(p1, v3, f14, f15);
			var v5: array<int> := new int[27](i1 => i1 * |[p1]|);
			var v6: map<array<int>, T0> := map[v5 := v4];
			if (v4 !in {v4, if (v5 in v6) then v6[v5] else v4, v4, v4}) {
				var v7 := 'v';
				var v8 := new C3(p1, v3[safeIndex(p1, |v3|) := v7], f14, v4.f15);
				v0[safeIndex(899, v0.Length)] := v8.f21;
				v0[safeIndex(899, v0.Length)] := if (!f19) then v8.f21 else v8.f21;
				var v9: set<seq<int>> := {seq(abs(0x2ed), i2  => (|v8.f21|))};
				var v10: seq<bool> := [f19];
				var v11: seq<bool> := [v10[safeIndex(p1, |v10|)]];
				var v12: set<string> := {v8.f21};
				r2 := if (!fm23(v8.f20, v4.f14, map[v7 := p1], v9, globalState)) then v11 == v10[safeIndex(v8.f20, |v10|) := f19] else fm34(globalState) > v12;
				var v13: array<bool> := new bool[8];
				globalState.f2 := v13;
				var v14: array<seq<bool>> := new seq<bool>[23](i3 => ([v4.f14])[safeIndex(|{f14}|, |[v4.f14]|) := true]);
				v14[safeIndex(603, v14.Length)] := [true, v4.f14, v4.f14];
				v5[safeIndex(7, v5.Length)] := v8.f20 * p1;
				v14[safeIndex(603, v14.Length)], v7, v5[safeIndex(7, v5.Length)], f14 := v10, v7, safeDivisionInt(-(|v3| + p1), |(v4.f15 - v4.f15)|), v4.f14;
			} else {
				r0 := p1 * -p1;
				r0 := -(0x26d - (if (v4.f14) then p1 else -p1));
				var v15: multiset<bool> := multiset{f19, p1 != p1, false, false, true};
				v15 := ((multiset{v4.f14})[v4.f14 := abs(p1)] - multiset{v4.f14, f14, f19, f19}) - (v4.f15 + v15);
				v5[safeIndex(591, v5.Length)] := p1;
				v5[safeIndex(591, v5.Length)] := p1;
				r2 := v4.f14;
			}
			
			var v16: array<bool> := new bool[21];
			var v17 := DC18(v4.f14);
			v16[safeIndex(334, v16.Length)] := v17.cf33;
			v16[safeIndex(334, v16.Length)] := !(!f19 && f14);
		} else {
			globalState.f9 := p1 > 819;
			var v18 := 'c';
			var v19: map<char, int> := map[v18 := p1];
			var v20: seq<int> := [p1, p1];
			var v21: multiset<int> := multiset{-0x2c7};
			var v22: map<int, int> := map[p1 := p1];
			if (fm23(p1, f14, v19, {v20}, globalState) <== (fm0(true, p1, fm0(f19, 0x36d, p1, v21, globalState), v21, globalState) !in v22)) {
				globalState.f3 := p1;
				var v23: C1 := new C1();
				v23 := v23;
				var v24: seq<bool> := [f14, f19, f14, f14];
				var v25: multiset<seq<bool>> := multiset{v24};
				var v26: seq<multiset<int>> := [v21];
				r1 := fm17(p1, f14, multiset{v24} - v25, |v26|, globalState);
				m6(globalState);
				var v27: array<D1> := new D1[6];
				var v28: seq<array<D1>> := [v27, v27, v27, v27];
				r1 := -|v28|;
			} else {
				var v29 := "erx";
				var v30: seq<bool> := [f19, f14];
				var v31 := new C3(p1, v29, !f14, multiset(v30));
				r2, v20, globalState.f3, globalState.f9 := -0x3be != p1, [p1], fm0(v31.fm2(f14, f19, globalState), p1, |fm35(p1, p1, globalState)|, v21, globalState), f19;
				r0 := p1;
				globalState.f3 := v31.f20;
				var v32 := new C3(v31.f20, "oek", f14, f15);
			}
			
			v18 := v18;
			var v33 := DC0(f19);
			var v34 := DC2(v33);
			var v35 := DC2(v34);
			match (v35) {
				case DC1(cf1) =>
					var v36 := "ukldksby";
					var v37: map<bool, int> := map[f19 := p1];
					v18 := fm25(p1, safeDivisionInt(p1, p1), v18 in v36, v37, globalState);
					globalState.f3 := p1;
					globalState.f9 := cf1;
					var v38: array<D3> := new D3[8](i4 => DC10(map[DC5(f14, f19).cf6 := cf1]));
					v38 := v38;
				case DC0(cf0) =>
					globalState.f3 := safeDivisionInt(p1, if (p1 in v22) then v22[p1] else p1);
					var v39: map<bool, int> := map[cf0 := p1];
					var v40: seq<map<bool, int>> := [v39];
					var v41: seq<seq<map<bool, int>>> := [fm36(v21, p1, cf0, f19, globalState)];
					v40 := v41[safeIndex(-p1, |v41|)] + v40;
					var v42: array<bool> := new bool[12];
					v42[safeIndex(853, v42.Length)] := f14;
					var v43: seq<bool> := [cf0];
					v42[safeIndex(853, v42.Length)] := (f19 <== f19) && v43[safeIndex(p1, |v43|)];
					m6(globalState);
				case DC2(cf2) =>
					var v44: array<array<int>> := new array<int>[3];
					var v45: array<int> := new int[25](i5 => safeDivisionInt(i5, 0x24c));
					v44[safeIndex(265, v44.Length)] := v45;
					v44[safeIndex(265, v44.Length)] := v45;
					var v46 := DC23(f14, |v20[safeIndex(-p1, |v20|) := p1]|, f14);
					var v47: map<D7, bool> := map[v46 := fm2(f19, f19, globalState)];
					v47 := v47[v46 := f19];
					var v48: seq<bool> := [f14];
					var v49: array<bool> := new bool[1](i6 => f19);
					var v50: set<bool> := {f19};
					var v51 := DC24(v50);
					var v52: C1 := new C1();
					var v53: map<int, C1> := map[p1 := v52];
					var v54: map<map<int, C1>, set<bool>> := map[v53 := {fm2(true, f19, globalState), !false}];
					var v55: multiset<seq<bool>> := multiset{v48, v48};
					var v56: array<bool> := new bool[25] [f19, f19, f14, !(if (f19) then f19 else f19), f14, f14, f14, !f19, true, p1 == |v48|, v49 !in map[v49 := v18], !f19, !f14, (if (p1 in v22) then v22[p1] else p1) <= |p0|, if (f19) then f19 else f19, f19, f19, f14, v51.cf45 !! (if (v53 in v54) then v54[v53] else v50), true, !(if (f19) then f19 else fm2(f19, f14, globalState)), f19 ==> f14, p1 == fm17(p1, f19, v55, 29, globalState), f14, f19];
					v49[safeIndex(870, v49.Length)] := f19;
					var v57: array<array<bool>> := new array<bool>[6];
					v57[safeIndex(25, v57.Length)] := v49;
					var v58: map<bool, array<bool>> := map[f14 := v49];
					v49[safeIndex(870, v49.Length)], v57[safeIndex(25, v57.Length)] := !!!f14, if (f19 in v58) then v58[f19] else v56;
					var v59: map<int, bool> := map[-0x107 := false];
					v44[safeIndex(265, v44.Length)][safeIndex(932, v44[safeIndex(265, v44.Length)].Length)] := if (p1 in v21) then v21[p1] else |v59|;
					var v60: map<bool, int> := map[v49[safeIndex(870, v49.Length)] := |(seq(abs(0x102), i7  => (p1)))|];
					globalState.f3, v44[safeIndex(265, v44.Length)][safeIndex(932, v44[safeIndex(265, v44.Length)].Length)] := if (f19 in v60) then v60[f19] else p1, p1 - p1;
			}
			
			var v61: array<int> := new int[6](i8 => i8 + p1);
			v61[safeIndex(327, v61.Length)] := p1;
			var v62: seq<bool> := [fm2(f14, false, globalState)];
			r0, v61[safeIndex(327, v61.Length)], globalState.f3, r2, globalState.f3 := p1, p1, p1 * |(v62 + v62)|, fm2(f14, f19, globalState), if (fm2(f14, f14, globalState)) then |(v62 + v62)| else p1;
		}
		
		var v63 := "iep";
		v63 := v63;
		if (f19) {
			globalState.f9 := f19;
			var v64 := DC14(false, f19);
			var v65: array<int> := new int[5];
			var v66: map<array<int>, bool> := map[v65 := !f19];
			var v67 := DC16(v65);
			var v68: set<bool> := {false};
			var v69: array<bool> := new bool[22] [!v64.cf26, f19, false, f19, if (true) then f14 else f19, fm2(f14, f19, globalState), f14, f14, f19, f14, fm2(f14, f14, globalState), if (f19) then f19 else f19, f19 <==> (if (v67.cf29 in v66) then v66[v67.cf29] else f14), !f14, f19, f19, v68 >= {!f19, f14, f19, f14}, f19, fm2(f14, f19, globalState), f14, p1 == p1, !f19];
			v69[safeIndex(949, v69.Length)] := f14;
			v69[safeIndex(949, v69.Length)] := f14;
			var v70: map<bool, array<bool>> := map[f14 := v69];
			v70 := v70[f19 := v69];
			var v71: seq<multiset<bool>> := [f15];
			m7(v71, v69[safeIndex(949, v69.Length)], p1, v69[safeIndex(949, v69.Length)], globalState);
			r2 := p1 == p1;
		} else {
			globalState.f3 := safeDivisionInt(-0x1ca, 0x2cb);
			var v72: set<bool> := {true};
			if (f14 !in v72) {
				var v73: array<set<bool>> := new set<bool>[8];
				v73 := v73;
				var v74: seq<bool> := [f14, f19];
				globalState.f9 := v74[safeIndex(p1, |v74|)];
				var v75 := DC1(!f14);
				var v76: map<int, set<bool>> := map[0x1f4 := {f19, f14}];
				var v77: array<D0> := new D0[25] [v75, v75, v75, v75, v75.(cf1 := f19), v75, v75, v75, v75, fm37(p0, v76, f19, DC1(!f14).cf1, globalState), v75, v75, fm37(p0, v76, f14, f19, globalState).(cf1 := f14), v75, v75, DC1(f14), v75, v75, v75, v75, v75, v75, v75, v75, v75.(cf1 := f14)];
				v77 := v77;
				globalState.f9 := false;
				v63 := v63;
			} else {
				globalState.f9 := if (f14) then f14 else f19;
				var v78: array<bool> := new bool[6](i9 => f15 !! f15);
				globalState.f2 := v78;
				r2 := f14;
				var v79: map<bool, bool> := map[f19 := f14];
				var v80: set<map<bool, bool>> := {v79, v79, v79};
				var v81: map<set<map<bool, bool>>, int> := map[v80 := -(p1 - 0x30)];
				v81 := v81[v80 := p1];
				var v82: array<int> := new int[25];
				v82[safeIndex(635, v82.Length)] := safeModuloInt(p1, p1);
				v82[safeIndex(635, v82.Length)] := -0x258;
			}
			
			var v83 := 'c';
			var v84: map<char, int> := map[v83 := p1];
			var v85: seq<int> := [p1];
			var v86: map<seq<int>, bool> := map[v85 := f19];
			var v88 := DC24({f14, fm23(p1, f14, v84, set v87 : seq<int> | v87 in v86 :: (v87), globalState)});
			v88 := v88;
			var v89: array<int> := new int[1](i10 => safeDivisionInt(i10, 380));
			var v90: map<int, string> := map[p1 := v63];
			var v91 := DC28(DC26(!f14, v90));
			var v92: multiset<D8> := multiset{v91};
			var v93: map<array<int>, int> := map[v89 := p1];
			var v94: map<bool, map<array<int>, int>> := map[f19 := v93];
			var v95: seq<bool> := [f14, !f14];
			var v96: seq<array<int>> := [v89, v89];
			var v97 := DC11(|v92|, |(if (v95[safeIndex(p1, |v95|)] in v94) then v94[v95[safeIndex(p1, |v95|)]] else map[v96[safeIndex(p1, |v96|)] := p1])|, f19, p1);
			v89, globalState.f9, v97, f14 := v89, f19, v97, !f19;
			globalState.f9 := v95[safeIndex(0x20d, |v95|)];
		}
		
		if (f19 && !f14) {
			var v98: seq<bool> := [f19];
			r1 := |((multiset{f14} * f15) + (multiset(v98) + f15))|;
			if (!fm2(p1 > p1, f19, globalState)) {
				globalState.f9 := f19;
				v63 := v63;
				r1 := p1 + p1;
				var v99: array<int> := new int[10](i11 => i11 + -p1);
				v99[safeIndex(808, v99.Length)] := p1 + p1;
				var v100: seq<int> := [|[f19, f14]|, p1];
				var v101: multiset<int> := multiset{p1};
				var v102: seq<int> := [p1 * fm0(f14, |v100|, p1, v101, globalState)];
				var v103: array<bool> := new bool[9] [!f19, f19, f14, f14, f19, f19, !f14, f14, f19];
				var v104 := 'n';
				var v105: map<char, int> := map[v104 := |v63|];
				var v106: set<seq<int>> := {v100};
				var v107 := DC23(fm23(|[v103]|, f19, v105, v106, globalState), p1, f14);
				v99[safeIndex(808, v99.Length)], r0 := v102[safeIndex(v107.cf43, |v102|)], p1;
				var v108: map<int, bool> := map[p1 := false];
				var v109: map<map<int, bool>, int> := map[v108 := v99[safeIndex(808, v99.Length)]];
				var v110 := DC21(v109);
				var v111: set<bool> := {f14, f14, f19, f14, f14};
				globalState.f3, r0, v99[safeIndex(808, v99.Length)], v110 := |(({f14, false, f14, f14} * {f19}) * v111)|, v99[safeIndex(808, v99.Length)], v99[safeIndex(808, v99.Length)], v110;
			} else {
				var v112: multiset<D8> := multiset{DC25()};
				var v113 := DC7(v98);
				globalState.f3, r1, v112, v63 := |v113.cf9|, safeModuloInt(p1 + p1, p1), v112 * v112, "e";
				globalState.f3 := p1;
				var v114: seq<int> := [p1];
				var v115: multiset<int> := multiset{p1};
				var v116: map<bool, multiset<int>> := map[f19 := multiset(v114) * v115];
				v116 := v116[f14 := v115 * v115];
				var v117 := DC11(p1, p1, f19, p1);
				var v118: C3 := new C3(p1, seq(abs(0x29d), i12  => ('b')), f19, f15);
				var v119: multiset<C3> := multiset{v118, v118};
				var v120 := 'q';
				var v121: set<char> := {v120, v120};
				var v122: multiset<set<char>> := multiset{v121};
				var v123: map<bool, int> := map[f14 := |v122|];
				var v124 := DC24({f19, false});
				var v125: map<D8, bool> := map[v124 := f14];
				var v126: map<bool, bool> := map[f14 := f19];
				var v127: array<D3> := new D3[20] [v117, v117, if (f19) then DC11(p1, p1, f19, 0x3c7) else v117.(cf19 := f14), DC11(p1, p1, f14, 0x1e5), v117, v117, v117, DC11(p1, p1, true, --|v119|), v117, v117, v117, v117, v117, v117, DC11(p1, fm0(true, fm0(f19, 0x217, p1, v115, globalState), v118.f20, multiset(v114), globalState), f14, p1).(cf20 := fm0(f19, fm0(f14, |{p1, |v123|}|, v118.f20, v115, globalState), v118.f20, v115, globalState), cf18 := |v125|, cf19 := f19), v117, v117, DC11(|v123|, |v126[f19 := f19]|, f14, v118.f20), if (f19) then v117 else v117.(cf20 := v118.f20), v117];
				v127[safeIndex(129, v127.Length)] := v117;
				var v128: array<multiset<bool>> := new multiset<bool>[1](i13 => f15);
				v128[safeIndex(891, v128.Length)] := f15;
				var v129: map<char, int> := map[fm25(p1, |v98|, f14, v123, globalState) := p1];
				var v130: set<seq<int>> := {v114};
				var v133: map<seq<bool>, bool> := map[v98 := false];
				v127[safeIndex(129, v127.Length)], v128[safeIndex(891, v128.Length)], globalState.f9 := v117.(cf18 := p1, cf19 := v120 !in v129), f15, ((map[[fm23(|v118.f21|, fm23(v118.f20, f19, v129, v130, globalState), map v131 : char | v131 in v118.f21 :: (v131) := (p1), set v132 : seq<int> | v132 in multiset{v114, [-0x222]} :: (v132), globalState), f14] := f19])[[f19] := f19])[v98 := f19] == v133;
				var v134: seq<map<bool, bool>> := [v126 + v126, v126, map[true := false], v126];
				v134 := v134 + v134;
			}
			
			var v135 := DC23(f19, p1, f14);
			var v136: seq<seq<bool>> := [v98, fm38(globalState), v98];
			globalState.f3 := v135.cf43 * |v136|;
			globalState.f3 := p1;
			var v137: array<int> := new int[1];
			v137[safeIndex(773, v137.Length)] := -|v98|;
			var v138: set<bool> := {f19, f14};
			globalState.f3, f14, v137[safeIndex(773, v137.Length)] := -p1, v138 > v138, safeDivisionInt(p1, p1);
		} else {
			var v139: multiset<bool> := multiset{f19 || f19};
			var v140: array<array<bool>> := new array<bool>[20];
			var v141: seq<bool> := [f19];
			var v142 := DC1(f14);
			var v143: array<bool> := new bool[12] [f14, f14, f14, f14, !f14, v141[safeIndex(p1, |v141|)], f19, f14, v142.cf1, f19, f19, f14];
			v140[safeIndex(985, v140.Length)] := v143;
			v139, v140[safeIndex(985, v140.Length)] := v139 + f15, if (false) then v143 else v143;
			v143[safeIndex(828, v143.Length)] := f19;
			var v144: set<bool> := {f19, f19, false};
			v143[safeIndex(828, v143.Length)] := (--|v144| >= 70) ==> !f14;
			var v145: map<int, int> := map[0x2c5 := p1];
			var v147: map<int, bool> := map[|v63| := f19];
			var v148: seq<map<int, int>> := [(map v146 : int | v146 in p0 :: (v146 + p1) := (p1))[0x379 := |v147|], v145];
			globalState.f3, v145, v63 := safeModuloInt(-|v63|, p1), v145 + v148[safeIndex(|(map v149 : int | (581 <= v149) && (v149 < 0x1b7) :: (safeDivisionInt(v149, |v144|)) := (f19))|, |v148|)], v63;
			var v150: multiset<int> := multiset{p1, p1};
			var v151 := DC4(fm0(v143[safeIndex(828, v143.Length)], p1, p1, v150, globalState), p1);
			var v153 := DC11(v151.cf5, |(map v152 : int | v152 in v150 :: (v152 * p1) := (v145))|, f14, p1);
			match (v153) {
				case DC11(cf17, cf18, cf19, cf20) =>
					var v154: map<string, map<int, int>> := map[v63 + v63 := v145];
					var v155 := 'd';
					var v156 := DC29(v145);
					v154 := v154[v63[safeIndex(|(seq(abs(345), i14  => ('v')))|, |v63|) := v155] := v156.cf49];
					r2 := if ((cf20 * cf17) in v147) then v147[cf20 * cf17] else cf19;
					var v157: set<char> := {v155, v155};
					var v158: seq<set<char>> := [v157, v157 - v157];
					v158 := v158;
					f14 := ((seq(abs(0x89), i15  => (v155))) != v63) <== cf19;
				case DC10(cf16) =>
					m6(globalState);
					var v159: array<int> := new int[1] [p1 - p1];
					v159[safeIndex(918, v159.Length)] := 0x3c6;
					v159[safeIndex(918, v159.Length)] := (p1 + fm0(f14, p1, p1, fm30(f19, globalState), globalState)) - p1;
					r2 := true;
					v143[safeIndex(828, v143.Length)], globalState.f2, r0, globalState.f3 := v159[safeIndex(918, v159.Length)] >= v151.cf4, v143, -942, p1 * v159[safeIndex(918, v159.Length)];
			}
			
			var v160: map<array<bool>, int> := map[v140[safeIndex(985, v140.Length)] := p1];
			var v161 := DC13(seq(abs(671), i16  => ('p')), 0x312, f19, v140[safeIndex(985, v140.Length)]);
			var v162: multiset<seq<bool>> := multiset{[f14, f19]};
			r1, globalState.f3, r1, r0, r1 := |fm30(v143[safeIndex(828, v143.Length)], globalState)| + p1, -(if (-|v160| != p1) then safeDivisionInt(p1, p1) else p1), (p1 * p1) - -p1, p1 + (--|v161.cf22| * -p1), safeDivisionInt(safeDivisionInt(fm17(p1, f14, v162, p1, globalState), p1), p1);
		}
		
		var v163: array<bool> := new bool[15];
		var v164: seq<array<bool>> := [v163];
		var v165 := DC13("uoqpauedn", p1, f14, v164[safeIndex(p1, |v164|)]);
		match (v165.(cf22 := seq(abs(0x1af), i17  => ('m')))) {
			case DC13(cf22, cf23, cf24, cf25) =>
				var v166: map<int, bool> := map[cf23 := cf24];
				var v167: map<int, map<int, bool>> := map[|v63| := v166];
				v167 := v167[safeDivisionInt(cf23, cf23) := v166];
				cf25[safeIndex(241, cf25.Length)] := cf24;
				cf25[safeIndex(241, cf25.Length)] := fm2(f19, f14, globalState);
				if (cf25[safeIndex(241, cf25.Length)]) {
					var v168: array<int> := new int[17](i18 => i18 + p1);
					var v169: map<int, array<int>> := map[p1 := v168];
					globalState.f3 := -|((v169 + v169) + v169)|;
					var v170: C3 := new C3(p1, v63[safeIndex(cf23, |v63|) := 't'], false, f15);
					v170 := v170;
					var v171 := new C3(p1 + v170.f20, v63 + cf22, false, f15);
					globalState.f3 := |f15|;
					var v172 := 'g';
					var v173 := DC30(v172, cf24, p1, p0, v171.f20);
					var v174: seq<int> := [p1];
					var v175 := DC9(cf25[safeIndex(241, cf25.Length)], v171.f20, v172);
					var v176: seq<bool> := [f14];
					var v177: set<string> := {v63};
					var v179: array<D9> := new D9[21] [DC30(v172, f19, cf23, p0, p1), v173, v173, v173, v173.(cf50 := v172), v173, v173, DC30(v172, cf24, v170.f20, p0, 746), DC30(v172, f19, v171.f20, p0, p1), DC30(v172, false, -|[v174, v174]|, {v171.f20, 0x5}, v171.f20), DC30(v175.cf15, cf25[safeIndex(241, cf25.Length)], v170.f20, p0, v170.f20), v173, fm39(|v176[safeIndex(p1, |v176|) := cf25[safeIndex(241, cf25.Length)]]|, v170.f20, p1, |v177|, globalState), v173, DC30(v172, f19, v170.f20, p0, p1), v173, DC30(v172, cf25[safeIndex(241, cf25.Length)], |(map v178 : char | v178 in (seq(abs(-0x244), i19  => (v172))) :: (v178) := (false))|, {|v170.f21|}, v171.f20), v173, v173, v173, v173];
					var v180: map<array<D9>, bool> := map[v179 := cf25[safeIndex(241, cf25.Length)]];
					r2, cf24 := if (v179 in v180) then v180[v179] else DC4(v170.f20, v170.f20).cf4 !in (seq(abs(588), i20  => (|map[cf24 := cf24]|))), v171.fm2(false, cf24, globalState);
				} else {
					f14 := (p1 - |"qd"|) != cf23;
					var v181 := DC14(f14, f19);
					var v182 := DC15(v181);
					var v183: map<D4, int> := map[v182 := cf23];
					var v184: set<bool> := {f14, f14};
					var v185: map<bool, set<bool>> := map[cf25[safeIndex(241, cf25.Length)] := v184];
					var v186: map<int, map<D4, int>> := map[p1 := v183[DC15(v181) := |v185|]];
					v186 := v186[cf23 := map[v182 := p1]];
					var v187: seq<bool> := [cf25[safeIndex(241, cf25.Length)]];
					var v188: multiset<string> := multiset{v63 + fm3(|v187|, cf23, !f19, cf23, globalState)};
					v188 := multiset{cf22};
					var v189: seq<int> := [|v63|, p1];
					var v190: map<array<bool>, int> := map[cf25 := |v189|];
					var v191 := 'y';
					var v192: map<char, int> := map[v191 := cf23];
					var v193: set<seq<int>> := {v189};
					var v194: seq<set<seq<int>>> := [v193];
					cf25[safeIndex(241, cf25.Length)] := cf25[safeIndex(241, cf25.Length)] <== fm23(p1, if (|v190| in v166) then v166[|v190|] else f14, v192, v194[safeIndex(940, |v194|)], globalState);
					var v195 := new C2(v166 + v166);
				}
				
				var v196: map<bool, multiset<bool>> := map[cf24 := f15];
				var v197 := new C3(p1, cf22 + cf22, cf24, if (cf25[safeIndex(241, cf25.Length)] in v196) then v196[cf25[safeIndex(241, cf25.Length)]] else f15);
			case DC14(cf26, cf27) =>
				var v198: multiset<int> := multiset{fm0(f14, p1, |v63|, multiset{|(seq(abs(-0x354), i21  => ('v')))|}, globalState), |v63|, 0x365};
				var v199: seq<multiset<int>> := [v198, multiset{p1}];
				r1 := fm0(cf26, |v63|, p1, v199[safeIndex(p1, |v199|)] * v198, globalState);
				var v200: seq<multiset<bool>> := [f15];
				var v201: seq<int> := [-|[p1, p1]|, p1];
				m7(v200 + (seq(abs(0x302), i22  => (f15))), p1 >= |v201|, -0x1c4, cf27, globalState);
				f14 := v63 < ("vlgqwvqgo" + v63);
				globalState.f3 := p1;
			case DC12(cf21) =>
				var v202: set<bool> := {f14};
				var v203: seq<int> := [|v202|, p1];
				var v204 := DC3(|v63|);
				var v205 := DC32(v203);
				var v206: map<int, bool> := map[p1 := f19];
				var v207: C2 := new C2(v206);
				var v208: map<C2, seq<int>> := map[v207 := v203];
				var v209: seq<bool> := [f19, f14];
				var v210: array<seq<int>> := new seq<int>[28] [seq(abs(955), i23  => (681)), v203, v203, [p1], fm21(v204, f14, fm40(globalState), globalState), v203, v203, v203, v203, v203, v203, v203, v203, v203, v203, [0xbd], v205.cf56, if (v207 in v208) then v208[v207] else v203, if (f19) then v203 else v203, v203, v203, v203 + v203, [p1, p1], seq(abs(-99), i24  => (|[!f19, f19, f14, f19, f19]|)), v203 + v203[safeIndex(|v209|, |v203|) := 790], v203, v203, v203];
				v210 := v210;
				var v211: map<int, set<bool>> := map[p1 := v202];
				var v212: array<set<bool>> := new set<bool>[9] [v202, v202, fm41(|v203|, p1, globalState), if (0x3e4 in v211) then v211[0x3e4] else v202, v202 * v202, v202, v202 + {false, false}, v202, v202];
				var v213: multiset<int> := multiset{-|v63|, p1};
				var v215: set<seq<int>> := {seq(abs(0x206), i25  => (|map[0x16e := |(map v214 : int | (210 <= v214) && (v214 < 0x3cc) :: (v214 * p1) := (true))|]|))};
				v212[safeIndex(632, v212.Length)] := {!true, fm23(fm0(false, 0x117, p1, multiset{if (p1 in v213) then v213[p1] else -p1, |v209|, p1}, globalState), f14, fm24(p1, f19, false, p1, globalState), v215, globalState), f19, f19, f19};
				globalState.f3, r2, v212[safeIndex(632, v212.Length)] := p1, !!((v63 + "nhxiqa") <= v63), if (f14) then {!f14, f14, f14, f19, f19} else v202 - v202;
				var v216: array<array<C0>> := new array<C0>[9];
				var v217: C3 := new C3(6, v63, f19, f15);
				var v218: map<array<array<C0>>, C3> := map[v216 := v217];
				v218 := v218[if (f19) then v216 else v216 := v217];
				var v219: map<set<int>, bool> := map[{v217.f20} := f19];
				var v220: array<int> := new int[14](i26 => i26 - DC4(-p1, --v217.f20).cf4);
				var v221: map<map<set<int>, bool>, array<int>> := map[v219 := v220];
				v221 := v221[v219 := v220];
			case DC15(cf28) =>
				if (p1 > 0x253) {
					v163[safeIndex(853, v163.Length)] := f19 <== f14;
					v163[safeIndex(853, v163.Length)] := f19;
					var v222 := 'r';
					var v223: map<char, bool> := map[v222 := v163[safeIndex(853, v163.Length)]];
					v223 := v223[v222 := f19];
					var v224: map<seq<bool>, int> := map[[!f19] := |v164|];
					var v226: seq<bool> := [v163[safeIndex(853, v163.Length)]];
					var v227: map<int, string> := map[p1 := v63];
					var v228 := DC26(v163[safeIndex(853, v163.Length)], v227);
					var v229: multiset<D8> := multiset{fm42(false, !v163[safeIndex(853, v163.Length)], 265, v163[safeIndex(853, v163.Length)], globalState), v228};
					var v230 := DC11(p1, p1, f14, -0x116);
					var v231: multiset<int> := multiset{v230.cf18, |v63|};
					var v232: set<multiset<int>> := {v231, v231};
					v224, globalState.f9, f14, v163[safeIndex(853, v163.Length)] := map v225 : seq<bool> | v225 in [v226] :: (v225) := (p1), !(v229 > v229), v163[safeIndex(853, v163.Length)], v232 !! v232;
					var v233: array<array<int>> := new array<int>[2];
					var v234: array<int> := new int[28](i27 => safeDivisionInt(i27, p1));
					v233[safeIndex(844, v233.Length)] := v234;
					var v235: map<int, char> := map[p1 := v222];
					var v236: map<bool, char> := map[f14 := v222];
					globalState.f3, globalState.f3, v233[safeIndex(844, v233.Length)], v63 := p1, p1, v234, (fm3(-p1, p1, !(if (v222 in v223) then v223[v222] else f14), safeModuloInt(|v235|, p1), globalState))[safeIndex(p1, |fm3(-p1, p1, !(if (v222 in v223) then v223[v222] else f14), safeModuloInt(|v235|, p1), globalState)|) := if (!f14 in v236) then v236[!f14] else v222];
					var v237: set<int> := {|{f19, f19, !v163[safeIndex(853, v163.Length)]}|, v230.cf18, p1};
					v237 := v237;
				} else {
					var v238: array<seq<int>> := new seq<int>[14](i28 => [p1]);
					var v239: map<char, int> := map['s' := p1];
					var v240: map<int, map<char, int>> := map[p1 := v239];
					var v241: map<bool, bool> := map[f19 := fm2(f19, f14, globalState)];
					var v242: seq<int> := [p1, |(if (|v241| in v240) then v240[|v241|] else v239)|, p1, p1, p1];
					v238[safeIndex(696, v238.Length)] := if (f14) then v242 else v242;
					v238[safeIndex(696, v238.Length)] := v242;
					var v243: seq<string> := [v63];
					r1 := if (f19) then safeDivisionInt(p1, |v63|) else |v243|;
					var v244 := DC12(v163);
					var v245: array<D4> := new D4[4] [v244, v244, v244, v244];
					var v246: map<array<D4>, bool> := map[v245 := true];
					globalState.f9 := v245 !in v246;
					var v247: map<int, bool> := map[|v63| := f14];
					f14 := (|v247| == p1) <==> f19;
					var v248: seq<bool> := [f19];
					f14 := f19 <== v248[safeIndex(364, |v248|)];
				}
				
				if (!true) {
					var v249: seq<int> := [p1, p1];
					var v250: multiset<int> := multiset{p1};
					var v251 := new C3(fm0(f19, |v249|, p1, v250, globalState), v63, p1 >= p1, f15);
					var v252: seq<multiset<int>> := [v250, v250, v250, v250];
					m7(fm43(v251.f20, p1, |[-fm0(!f19, v251.f20, p1, v250, globalState)]|, globalState), f14, |(v250 * v252[safeIndex(p1, |v252|)])|, f14, globalState);
					var v253 := DC8(|(seq(abs(-565), i29  => ('g')))|, safeDivisionInt(v251.f20, p1), p1);
					var v254: array<int> := new int[20];
					v254[safeIndex(298, v254.Length)] := -safeModuloInt(v251.f20, p1);
					v254[safeIndex(605, v254.Length)] := |v249| - p1;
					v163[safeIndex(902, v163.Length)] := true;
					var v255: set<array<int>> := {v254};
					r2, v253, v254[safeIndex(298, v254.Length)], v254[safeIndex(605, v254.Length)], v163[safeIndex(902, v163.Length)] := f19, fm44(f19, v251.f20, v251.f20, globalState), if (f14) then p1 else p1, v251.f20, (v255 + v255) == {v254};
					var v256: map<bool, int> := map[true := 560];
					v254[safeIndex(298, v254.Length)] := v251.f20 * (if (f14 in v256) then v256[f14] else v251.f20);
					globalState.f3 := |"owt"|;
				} else {
					var v257: array<array<int>> := new array<int>[15];
					v257 := v257;
					var v258 := DC26(f19, map[p1 := v63]);
					var v259: array<int> := new int[10];
					v259[safeIndex(559, v259.Length)] := if (f19) then p1 else p1;
					var v260 := DC14(f14, fm2(f14, !f14, globalState));
					var v261: map<seq<char>, array<int>> := map[fm3(p1, p1, v260.cf26, p1, globalState) := v259];
					var v262: seq<string> := [v63, v63, v63];
					var v263: map<int, string> := map[p1 := v262[safeIndex(p1, |v262|)]];
					r0, globalState.f3, v258, v259[safeIndex(559, v259.Length)] := p1, |v261|, v258.(cf47 := v263), safeModuloInt(p1 * |f15|, p1);
					var v264 := DC9(f19, v259[safeIndex(559, v259.Length)], 'n');
					r0 := v264.cf14;
					r2 := !f14;
					var v265: map<int, int> := map[v259[safeIndex(559, v259.Length)] := |p0|];
					var v266: map<map<int, int>, int> := map[v265 := v259[safeIndex(559, v259.Length)]];
					v266 := v266[v265 + v265 := v259[safeIndex(559, v259.Length)]];
				}
				
				f14 := !f14;
				if (f19) {
					var v267: map<int, bool> := map[p1 := f14];
					var v268 := new C2(v267[p1 := f19]);
					var v269: array<char> := new char[12];
					var v270 := 'a';
					v269[safeIndex(236, v269.Length)] := v270;
					var v271: set<multiset<bool>> := {f15, f15};
					globalState.f9, globalState.f9, v269[safeIndex(236, v269.Length)] := 0xf1 < --p1, v271 >= v271, v63[safeIndex(p1, |v63|)];
					var v272: multiset<int> := multiset{p1, p1};
					var v273: seq<int> := [|v272|];
					var v274: set<seq<int>> := {v273, [p1], v273};
					var v275: map<bool, bool> := map[false := f19];
					var v276: map<int, int> := map[-|"wwmccp"| := if (f14) then |{p1, p1}| else |(fm45(p1, fm23(p1, true, map[v270 := p1], v274, globalState), p1, globalState).(cf16 := v275)).cf16|];
					v276 := v276[p1 := |v63|];
					globalState.f3 := p1;
					globalState.f9 := v272 !! (v272 * v272);
				} else {
					globalState.f3 := p1;
					var v277: map<bool, int> := map[false := p1 * p1];
					v277 := v277;
					r1 := p1;
					v63 := v63 + (seq(abs(-0x394), i30  => ('w')));
					var v278: multiset<int> := multiset{p1, p1};
					globalState.f3 := fm0(f14 <== f19, p1, if (f19) then p1 else p1, v278, globalState);
				}
				
		}
		
		r0 := p1;
		r1 := safeDivisionInt(p1, p1);
		r2 := f14;
	}
	method m6(globalState: GlobalState) {
		if (f14) {
			var v0: array<int> := new int[22](i0 => safeDivisionInt(i0, |[f14]|));
			var v1 := 835;
			v0[safeIndex(612, v0.Length)] := v1 * 0x2ba;
			var v2 := "dchmesu";
			var v3: seq<int> := [v1];
			var v4: seq<bool> := [f14];
			var v5: set<int> := {v1};
			v0[safeIndex(612, v0.Length)], v2, globalState.f9, globalState.f3, globalState.f3 := --v3[safeIndex(v1, |v3|)], v2 + v2, {v1, ---v1, |v4|} > v5, 0x2c7, safeModuloInt(v1, v1);
			var v6: array<bool> := new bool[29](i1 => f19 in map[f14 := f14]);
			globalState.f2 := v6;
			v6[safeIndex(909, v6.Length)] := f14;
			v6[safeIndex(909, v6.Length)] := f14;
			var v7: map<int, bool> := map[v0[safeIndex(612, v0.Length)] := false];
			v5 := fm22(|"epm"|, [v7], f19, f19, globalState) * (v5 * v5);
			var v8 := 'h';
			v1, globalState.f3, v0, v2 := -|(v2[safeIndex(v1, |v2|) := v8] + "ob")|, if (v6[safeIndex(909, v6.Length)]) then v0[safeIndex(612, v0.Length)] else |[v6[safeIndex(909, v6.Length)], f19, f14]|, v0, v2;
		} else {
			var v9 := 0x333;
			globalState.f3 := v9;
			var v10: array<bool> := new bool[2](i2 => f19);
			var v11: seq<array<bool>> := [v10, v10, v10];
			var v12: multiset<int> := multiset{-v9};
			globalState.f2 := v11[safeIndex(fm0(f14, -v9, v9, v12, globalState), |v11|)];
			var v13: array<array<int>> := new array<int>[15];
			v13 := new array<int>[8];
			v10[safeIndex(630, v10.Length)] := f14;
			v10[safeIndex(630, v10.Length)] := f15 > f15;
			f14 := v10[safeIndex(630, v10.Length)];
		}
		
		var v14 := 0xca;
		for i3 := v14 to v14 {
			var v15 := "gmbu";
			var v16: multiset<bool> := multiset{f19};
			var v17 := DC25();
			v15, v16, globalState.f9, v17 := v15, v16, f14, v17;
			var v18: array<bool> := new bool[25](i4 => !(i3 != v14));
			v18[safeIndex(997, v18.Length)] := f19;
			v18[safeIndex(997, v18.Length)] := f14;
			var v19: seq<bool> := [v18[safeIndex(997, v18.Length)]];
			var v20: multiset<seq<bool>> := multiset{fm38(globalState)};
			v14 := fm17(i3, !v18[safeIndex(997, v18.Length)], multiset{v19} * v20, -899, globalState);
			globalState.f9 := v19[safeIndex(safeDivisionInt(i3, i3), |v19|)];
		}
		if (f19) {
			var v21: array<set<bool>> := new set<bool>[24](i5 => {f14, f14});
			var v22: set<bool> := {f14};
			v21[safeIndex(579, v21.Length)] := v22 - v22;
			v21[safeIndex(579, v21.Length)] := v22;
			if (f14) {
				var v23: array<seq<int>> := new seq<int>[2](i6 => [-v14, v14, v14]);
				var v24: array<D0> := new D0[24];
				var v25 := DC1(fm2(f14, true, globalState));
				var v26 := DC2(v25);
				v24[safeIndex(618, v24.Length)] := v26;
				var v27: array<D0> := new D0[25](i7 => DC1(f14));
				var v28: map<int, set<bool>> := map[v14 := {f14, f19}];
				v27[safeIndex(462, v27.Length)] := fm37({v14}, v28, f14, f19, globalState);
				var v29: array<bool> := new bool[7];
				v29[safeIndex(222, v29.Length)] := !f19;
				var v30: set<int> := {v14};
				var v31 := "pw";
				var v32: map<int, string> := map[|(seq(abs(0x157), i8  => ('v')))| := v31];
				var v33 := DC26(f14, v32);
				v23, v24[safeIndex(618, v24.Length)], globalState.f9, v27[safeIndex(462, v27.Length)], v29[safeIndex(222, v29.Length)] := v23, v26, v30 !! v30, DC1(fm2(!f19, f14, globalState)).(cf1 := v33.cf46).(cf1 := !f19), f14;
				globalState.f2 := v29;
				f14 := !v29[safeIndex(222, v29.Length)];
				var v34 := new C1();
				var v35: multiset<int> := multiset{|"uwcabk"|};
				var v36: seq<multiset<int>> := [v35, v35];
				v36 := v36;
			} else {
				var v37 := "nenovha";
				var v38: seq<bool> := [f14, f19, f14, f14, f14];
				var v39 := 'y';
				globalState.f3 := |((v37 + v37) + (if (v38[safeIndex(v14, |v38|)]) then v37 else v37)[safeIndex(-v14, |(if (v38[safeIndex(v14, |v38|)]) then v37 else v37)|) := v39])|;
				var v40: seq<int> := [571];
				var v41 := DC32(v40);
				var v42: map<D10, bool> := map[v41 := f14];
				var v43: seq<string> := [v37, v37, "qvqxuld"];
				var v44: array<string> := new string[10] [v37 + v37, "ebteocn", "lronodbd", v37, v37, fm27(v14, true, globalState), if (if (v41 in v42) then v42[v41] else f19) then v37 else v37, v43[safeIndex(v14, |v43|)], if (f19) then v37 else "evmfvq", v37];
				v44[safeIndex(523, v44.Length)] := v37;
				v44[safeIndex(523, v44.Length)] := v37;
				globalState.f3 := v14;
				var v45: seq<multiset<bool>> := [f15];
				m7([multiset{!f14}] + v45, v40 == v40, v14, f14, globalState);
				var v46: array<bool> := new bool[24];
				var v47: multiset<char> := multiset{v39, v39, v39, v39, v39};
				v46[safeIndex(750, v46.Length)] := v14 == |v47|;
				var v48: multiset<int> := multiset{|v40|, 0xb4};
				var v49: array<seq<map<int, bool>>> := new seq<map<int, bool>>[28](i9 => seq(abs(0x10f), i10  => (map[v14 := false])));
				var v50: map<int, bool> := map[v14 := f14];
				var v51: seq<map<int, bool>> := [v50];
				v49[safeIndex(598, v49.Length)] := v51;
				var v52: set<int> := {|v48|, v14};
				v46[safeIndex(750, v46.Length)], globalState.f9, v48, v49[safeIndex(598, v49.Length)] := f14, |(set v53 : int | v53 in v52 :: (v53 - 0xd8))| < (v14 - v14), v48, [v50, v50 + v50, v50, v50[0x12a := f19] + v50[v14 := f19]];
			}
			
			var v54: map<bool, int> := map[f14 := v14];
			var v58: seq<map<bool, int>> := [v54];
			var v59: seq<int> := [v14, v14];
			var v60: set<int> := {v14, v59[safeIndex(|f15|, |v59|)], v14};
			var v61 := "chvabiuu";
			var v62: array<map<bool, int>> := new map<bool, int>[23] [v54 + map[f14 := v14], v54, v54, v54 + fm1(false, v14, f14, set v55 : int | (-621 <= v55) && (v55 < 0x36f) :: (v55 + |(map v56 : int | v56 in (map v57 : int | v57 in map[-v14 := f14] :: (safeDivisionInt(v57, v14)) := (f19)) :: (safeDivisionInt(v56, 403)) := (v14))|), globalState), v58[safeIndex(v14, |v58|)] + v54, v54, if (f19) then v54 else v54, v54, map[f14 := v14], v54, v54, v54[false := v14], map[f19 := v14] + fm1(f19, v14, f14, v60, globalState), v54, v54 + v54, v54, v54, v54, v54, v54, map[f19 := v14], v54 + v54, v54 + map[!f19 := |v61|]];
			v62[safeIndex(860, v62.Length)] := v54;
			v62[safeIndex(860, v62.Length)] := v54;
			v14 := -(if (f19) then v14 else -v14);
			var v63 := new C1();
		} else {
			globalState.f9 := f14;
			if (f19) {
				var v64 := "frcf";
				v64 := v64 + v64;
				var v65: map<bool, multiset<bool>> := map[f14 := f15 * f15];
				v65 := v65[f19 ==> !f19 := f15];
				var v66: array<bool> := new bool[3];
				var v67 := DC13(v64, v14, f19, v66);
				var v68: multiset<int> := multiset{v14, v14, v14};
				v64 := fm3(fm0(f14, |v67.cf22|, v14, v68, globalState), |{v14}|, f14, v14, globalState);
				var v69 := 'm';
				var v70: map<char, bool> := map[v69 := f14];
				v66[safeIndex(156, v66.Length)] := false;
				var v71: seq<bool> := [f19];
				var v72: map<int, bool> := map[|v64| := true];
				var v73: map<set<int>, array<bool>> := map[fm22(v14, [v72], f14, f19, globalState) := v66];
				var v74: set<multiset<int>> := {v68};
				var v75: multiset<seq<bool>> := multiset{v71, v71, v71, v71};
				globalState.f9, v70, v66[safeIndex(156, v66.Length)], v14 := f19, if (v71[safeIndex(|v73|, |v71|)]) then v70 + v70 else fm46(v14, f19, globalState) + v70, true, fm17(v14, v74 !! v74, v75, v14, globalState);
				globalState.f3 := safeDivisionInt(v14, v14);
			} else {
				var v76: array<bool> := new bool[9];
				var v77: set<array<bool>> := {v76, v76};
				globalState.f3 := v14 * |v77|;
				var v78 := "wkk";
				var v79: set<bool> := {v78 < v78};
				globalState.f6 := v79;
				globalState.f3 := |v78| * v14;
				var v80: map<char, int> := map['i' := v14];
				var v81: seq<int> := [v14, |v78|];
				var v82: set<seq<int>> := {v81};
				var v83: map<bool, int> := map[fm23(v14, f14, v80, v82, globalState) || f14 := v14];
				v83 := v83[f14 := v14];
				globalState.f9 := f19;
			}
			
			var v84: array<seq<bool>> := new seq<bool>[9](i11 => [f19] + [!f19]);
			var v85: seq<bool> := [f14, f14, f19];
			v84[safeIndex(360, v84.Length)] := v85;
			var v86: map<bool, int> := map[!!f14 := v14];
			var v87 := "pw";
			var v88: seq<seq<bool>> := [v85];
			var v89: seq<int> := [v14, fm17(if (f14 in v86) then v86[f14] else v14, f14, multiset{v85, v85[safeIndex(|v87|, |v85|) := f19], v85, v85, v85}, -|v88[safeIndex(-v14, |v88|) := [f19]]|, globalState)];
			v84[safeIndex(360, v84.Length)], v89 := v85[safeIndex(v14, |v85|) := f19], seq(abs(0x190), i12  => (safeModuloInt(v14, |"ldnplt"|)));
			globalState.f9 := !!(f14 <==> !f14);
			globalState.f9 := f14;
		}
		
		var v90 := 'o';
		v90 := v90;
		var v91: array<bool> := new bool[7];
		var v92: seq<bool> := [f19, true, f19, f14, f19];
		var v93 := DC3(338);
		var v94: map<int, D1> := map[v14 := v93];
		v91[safeIndex(821, v91.Length)] := v92[safeIndex(|v94|, |v92|)];
		v91[safeIndex(821, v91.Length)] := f19;
		var v95: map<bool, int> := map[f14 := 0x65];
		var v96: map<int, string> := map[v14 := ("voqhvj")[safeIndex(v14, |"voqhvj"|) := v90]];
		var v97: multiset<int> := multiset{|v95|, v14, |v96|};
		v90 := fm25(v14, -v14 - |v97|, v91[safeIndex(821, v91.Length)] ==> f19, v95, globalState);
	}
	method m7(p0: seq<multiset<bool>>, p1: bool, p2: int, p3: bool, globalState: GlobalState) {
		var v0 := new C1();
		var v1: array<bool> := new bool[20];
		v1[safeIndex(886, v1.Length)] := p1;
		var v2: set<multiset<bool>> := {f15};
		var v3: array<set<multiset<bool>>> := new set<multiset<bool>>[7] [v2, DC33(v2).cf57, v2, v2, v2 - v2, v2, v2 - {f15, f15}];
		v3[safeIndex(43, v3.Length)] := fm47(globalState) - v2;
		var v4: set<int> := {p2};
		var v5 := "ovck";
		var v6: seq<int> := [-|v4|, --0x127, p2, |map[|v5| := v0]|];
		v1[safeIndex(886, v1.Length)], v3[safeIndex(43, v3.Length)], globalState.f9 := (p3 <== p3) && (v6 <= v6), v2, true;
		globalState.f3, globalState.f3 := p2, 538;
		var v7 := 'e';
		var v8: map<int, int> := map[p2 := -|v5|];
		var v9 := DC13("sikdur", p2, f19, v1);
		var v11: map<int, bool> := map[p2 := p3];
		var v12: seq<map<int, bool>> := [map v10 : int | (-314 <= v10) && (v10 < -878) :: (v10 + p2) := (f14), v11, v11, v11];
		var v13 := DC30(v7, p1, |v8|, fm22(v9.cf23, v12, f19, v1[safeIndex(886, v1.Length)], globalState), p2);
		var v14: map<D9, int> := map[v13 := p2];
		for i0 := |v14[v13 := p2]| to |(seq(abs(0x3a6), i1  => (v7)))| + p2 {
			var v15: multiset<int> := multiset{p2};
			var v16 := DC9(f19, |(seq(abs(41), i2  => ('k')))|, v7);
			var v17: array<D2> := new D2[6] [v16, v16, v16, v16, DC9(p3, p2, v7), DC9(p1, i0, v7)];
			var v18: seq<array<D2>> := [v17];
			var v19: array<int> := new int[22];
			var v20 := DC16(v19);
			var v21 := DC19(v20);
			var v22: map<bool, D5> := map[fm2(f14, p3, globalState) := v20];
			var v23 := DC19(if (f19 in v22) then v22[f19] else v21);
			var v24 := DC22(v1, v18, [v1, v1], v1, v23);
			var v25: seq<D7> := [v24];
			var v26: array<int> := new int[11] [p2, p2 + p2, p2, i0, (if (v1[safeIndex(886, v1.Length)] in f15) then f15[v1[safeIndex(886, v1.Length)]] else fm0(v1[safeIndex(886, v1.Length)], 24, i0, v15, globalState)) * p2, safeDivisionInt(-p2, |map[i0 := p2]|), |v5|, fm0(f14, p2, i0, v15, globalState), p2, i0 - |v5|, |v25|];
			var v27: seq<bool> := [f14, f14, f14, p3, f14];
			var v28: multiset<seq<bool>> := multiset{v27};
			v19[safeIndex(23, v19.Length)] := fm17(p2, !f19, v28, -fm0(v1[safeIndex(886, v1.Length)], i0, i0, v15, globalState), globalState);
			var v29 := DC37(v28);
			v19[safeIndex(23, v19.Length)] := safeModuloInt(p2 + 0x210, fm17(i0, v27[safeIndex(p2, |v27|)], v29.cf63, i0, globalState));
			var v30: map<bool, int> := map[f14 := v19[safeIndex(23, v19.Length)]];
			globalState.f9 := 0x80 > (i0 + (if (f19 in v30) then v30[f19] else p2));
			v7, v7, v19, globalState.f3 := v7, v7, v19, if (p1 in v30) then v30[p1] else p2 - p2;
			globalState.f9 := v27 == v27;
		}
		var v31: C2 := new C2(v11);
		v31 := v31;
		var v32: seq<bool> := [p1];
		var v33: map<bool, seq<bool>> := map[v1[safeIndex(886, v1.Length)] := v32 + v32];
		globalState.f3, globalState.f9, globalState.f9 := |(if ((f14 && p1) in v33) then v33[f14 && p1] else v32)|, f19 <== v1[safeIndex(886, v1.Length)], f19;
	}
}

class C5 extends T0 {
	const f18 : int
	constructor (f18 : int, f14 : bool, f15 : multiset<bool>) {
		this.f18 := f18;
		this.f14 := f14;
		this.f15 := f15;
	}
	
	function fm2(p0: bool, p1: bool, globalState: GlobalState): bool {
		f14
	}
	function fm3(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): string {
		"h" + "e"
	}
	function fm15(p0: bool, p1: bool, p2: int, globalState: GlobalState): string {
		("y" + "kfx") + ((seq(abs(0x28e), i0  => ('v'))) + (seq(abs(0x19), i1  => ('o'))))
	}
	function fm16(p0: int, p1: bool, p2: bool, globalState: GlobalState): bool {
		!f14
	}
	method m0(globalState: GlobalState) returns (r0: int, r1: bool, r2: map<map<int, bool>, int>, r3: bool) {
		var v0: seq<bool> := [f14];
		var v1 := DC7(v0);
		var v2: map<int, seq<bool>> := map[f18 := v1.cf9];
		var v3: seq<seq<bool>> := [v0, v0, v0, v0, v0];
		v2 := v2[f18 := v3[safeIndex(0x348, |v3|)]];
		var v4: seq<set<bool>> := [{true, f14, f14, true, f14}];
		var v5 := new C4(f14, |v3| > |v4|, f15);
		r0 := |v3[safeIndex(f18, |v3|)]|;
		r0 := f18;
		var v6: array<bool> := new bool[1];
		v6[safeIndex(819, v6.Length)] := f14;
		v6[safeIndex(819, v6.Length)] := if (f14) then f14 else v5.f19;
		var v7: map<int, int> := map[f18 := f18];
		var v8: map<bool, int> := map[v5.f19 := |v7|];
		var v9 := "lm";
		var v10: seq<int> := [f18];
		var v11: array<int> := new int[26] [f18, f18, f18, f18, f18, f18, f18, f18, f18, |v8|, |v9|, -0x25c, |v10|, 817, |[v6[safeIndex(819, v6.Length)], v5.f19]|, f18, f18, f18, f18, 0x3cf, f18, -0xbb, f18, f18, v10[safeIndex(f18, |v10|)], f18];
		var v12 := DC16(v11);
		var v13: map<string, array<int>> := map[v9 := v11];
		var v14: map<bool, array<int>> := map[v5.f19 := v11];
		var v15: array<array<int>> := new array<int>[26] [(v12.(cf29 := v11)).cf29, v11, v11, v11, v11, if (v5.f19) then v11 else v11, v11, v11, v11, v11, v11, if ((seq(abs(979), i0  => ('m'))) in v13) then v13[seq(abs(979), i0  => ('m'))] else v11, v11, v11, v11, v11, v11, v11, v11, v11, v11, v11, v11, v11, if (v6[safeIndex(819, v6.Length)] in v14) then v14[v6[safeIndex(819, v6.Length)]] else v11, v11];
		v15 := new array<int>[9];
		r0 := f18;
		r1 := f14 && f14;
		var v16: seq<array<bool>> := [v6];
		var v17 := DC13(v9, f18, v6[safeIndex(819, v6.Length)], v16[safeIndex(f18, |v16|)]);
		var v18: map<map<int, bool>, int> := map[((map[f18 := v6[safeIndex(819, v6.Length)]])[f18 := v5.f19])[v17.cf23 := f14] := f18];
		r2 := v18;
		r3 := v6[safeIndex(819, v6.Length)];
	}
	method m1(p0: set<int>, p1: int, globalState: GlobalState) returns (r0: int, r1: int, r2: bool) {
		var v0: array<int> := new int[27](i1 => i1 * |multiset{'b'}|);
		forall i0 | 0 <= i0 < v0.Length {
			v0[i0] := safeModuloInt(i0, p1);
		}
		var v1: map<bool, bool> := map[f14 := f14];
		var v2 := DC10(map[f14 := f14]);
		v1 := match v2 {
			case DC11(cf17, cf18, cf19, cf20) => map[f14 := cf19]
			case DC10(cf16) => cf16 + cf16
		};
		var v3: seq<int> := [f18];
		var v4: seq<bool> := [f14];
		var v5 := new C4(v3 <= v3, f14, multiset{v4[safeIndex(f18, |v4|)], f14});
		var v6 := 't';
		var v7 := "h";
		var v8: seq<string> := [v7, v7];
		var v9: array<bool> := new bool[21] [v5.f19, v5.f19, 720 >= f18, f14, f14, f15 >= f15, v5.f19, f14, f14, f14, true, f14, v6 in (seq(abs(815), i2  => (v6))), p1 <= |v7|, f18 < |v8|, p0 != p0, !v5.f19, p0 >= p0, v5.f19, fm2(v5.f19, v5.f19, globalState), f14];
		v9[safeIndex(928, v9.Length)] := f14;
		v0[safeIndex(46, v0.Length)] := -0x3a6;
		var v10: array<map<string, bool>> := new map<string, bool>[25];
		var v11: map<string, bool> := map[fm3(p1, p1, f14, p1, globalState) := f14];
		v10[safeIndex(245, v10.Length)] := v11;
		var v12: multiset<int> := multiset{-0x345};
		v9[safeIndex(928, v9.Length)], v0[safeIndex(46, v0.Length)], f14, v10[safeIndex(245, v10.Length)] := v12 !! (multiset{p1})[f18 := abs(f18)], -f18, v5.f19, v11;
		var v13: seq<multiset<int>> := [v12, multiset{f18}];
		var v14: set<multiset<int>> := {v13[safeIndex(p1, |v13|)]};
		for i3 := -v0[safeIndex(46, v0.Length)] to safeDivisionInt(v0[safeIndex(46, v0.Length)], |v14|) {
			var v15: set<seq<bool>> := {v4};
			var v17: set<bool> := {v5.fm2(v5.f19, f14, globalState)};
			var v18: map<bool, set<bool>> := map[v9[safeIndex(928, v9.Length)] := v17];
			var v19: map<seq<int>, int> := map[v3 := |(if (v5.f19 in v18) then v18[v5.f19] else v17)|];
			v7, v0[safeIndex(46, v0.Length)], v0[safeIndex(46, v0.Length)], v7, globalState.f9 := v7, -safeDivisionInt(-p1 * i3, |((set v16 : seq<bool> | v16 in v15 :: (v16)) * v15)|), f18, fm15(v5.f19, v0[safeIndex(46, v0.Length)] == p1, if (v3 in v19) then v19[v3] else p1, globalState), v9[safeIndex(928, v9.Length)];
			var v20: array<char> := new char[7](i4 => 'g');
			var v21: map<bool, int> := map[v9[safeIndex(928, v9.Length)] := 0x5f];
			v20[safeIndex(812, v20.Length)] := if (v5.f19) then fm25(p1, p1, false, v21, globalState) else v6;
			v20[safeIndex(812, v20.Length)] := 'w';
			var v22: map<map<bool, int>, bool> := map[v21 := fm2(v5.f19, v5.f19, globalState)];
			var v24: set<map<bool, int>> := {v21};
			if ((set v23 : map<bool, int> | v23 in v22 :: (v23)) <= v24) {
				r1 := i3;
				v20[safeIndex(812, v20.Length)] := 'h';
				globalState.f9, globalState.f9, v0[safeIndex(46, v0.Length)] := v5.f19, !(v5.f19 <==> f14), |(seq(abs(720), i5  => (|p0|)))|;
				globalState.f2 := new bool[29];
				var v25: array<array<bool>> := new array<bool>[13];
				var v26 := DC1(v5.f19);
				globalState.f3, r2, globalState.f3, v25 := v0[safeIndex(46, v0.Length)], if (v5.f19) then v26.cf1 !in v21 else v9[safeIndex(928, v9.Length)] || v9[safeIndex(928, v9.Length)], (0x17 * i3) - (0x3f * |"v"|), v25;
			} else {
				var v27: array<seq<int>> := new seq<int>[11](i6 => v3 + v3);
				v27 := v27;
				v9[safeIndex(928, v9.Length)] := !!!true;
				f14 := !(p1 <= f18);
				globalState.f3 := p1;
				v0 := v0;
			}
			
			var v28: seq<array<bool>> := [v9];
			var v29 := DC5(v9[safeIndex(928, v9.Length)], fm2(false, f14, globalState));
			var v30 := DC34(f15, v12, v9, v29);
			var v31: array<array<bool>> := new array<bool>[18] [v9, v9, v9, v9, v9, v28[safeIndex(p1, |v28|)], v9, v9, v9, v9, v9, v9, v9, v9, v9, v30.cf60, v9, v9];
			v31[safeIndex(777, v31.Length)] := v9;
			var v32: map<char, seq<bool>> := map[v20[safeIndex(812, v20.Length)] := v4];
			r2, v31[safeIndex(777, v31.Length)], v9[safeIndex(928, v9.Length)], v32 := v9[safeIndex(928, v9.Length)], v9, (f15 - f15) !! f15, v32;
		}
		var v33: map<int, array<bool>> := map[v0[safeIndex(46, v0.Length)] := v9];
		var v34: multiset<seq<bool>> := multiset{v4, v4};
		var v35: seq<array<bool>> := [v9];
		var v36: map<char, array<bool>> := map[v6 := v9];
		var v37: map<bool, int> := map[v9[safeIndex(928, v9.Length)] := f18];
		var v38: array<array<bool>> := new array<bool>[24] [v9, v9, v9, v9, v9, if (v5.fm17(f18, v5.f19, v34, v0[safeIndex(46, v0.Length)], globalState) in v33) then v33[v5.fm17(f18, v5.f19, v34, v0[safeIndex(46, v0.Length)], globalState)] else v9, v9, v9, v35[safeIndex(v0[safeIndex(46, v0.Length)], |v35|)], v9, v9, v9, v9, v9, v9, v9, v9, v9, if (fm25(p1, p1, v5.f19, v37, globalState) in v36) then v36[fm25(p1, p1, v5.f19, v37, globalState)] else v9, v9, v9, v9, v9, v9];
		var v39 := DC12(v9);
		var v40: map<D4, array<array<bool>>> := map[v39 := v38];
		var v41: seq<array<array<bool>>> := [v38];
		v38 := if (v39 in v40) then v40[v39] else v41[safeIndex(v3[safeIndex(p1, |v3|)], |v41|)];
		r0 := 0x231;
		r1 := p1;
		r2 := -p1 > v0[safeIndex(46, v0.Length)];
	}
}

class C6 extends T0 {
	constructor (f14 : bool, f15 : multiset<bool>) {
		this.f14 := f14;
		this.f15 := f15;
	}
	
	function fm2(p0: bool, p1: bool, globalState: GlobalState): bool {
		(set v0 : int | (0x1ea <= v0) && (v0 < 573) :: (safeDivisionInt(v0, 0x26d))) != ((set v1 : int | v1 in {-0xe0} :: (v1 * 544)) - (set v2 : int | (0x265 <= v2) && (v2 < 0x91) :: (v2 - 943)))
	}
	function fm3(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): string {
		"evj"
	}
	function fm11(p0: map<D0, bool>, p1: int, globalState: GlobalState): bool {
		f14 && ((set v0 : int | v0 in multiset{|multiset{|map[0x359 := |multiset{"bv", seq(abs(-494), i0  => ('v'))}|]|}|, 0x153} :: (v0 - 177)) <= {|[0x360, 0x1dc, 0xb4, DC3(0x286).cf3, 0x131]|, |"njb"|})
	}
	function fm12(p0: bool, p1: map<map<int, int>, int>, globalState: GlobalState): bool {
		(map[DC3(|[f14]|) := f15] + map[DC3(|"aksalss"|) := f15]) == (map v0 : D1 | v0 in (map v1 : D1 | v1 in [DC3(-509)] :: (v1) := (121)) :: (v0) := (f15))
	}
	method m0(globalState: GlobalState) returns (r0: int, r1: bool, r2: map<map<int, bool>, int>, r3: bool) {
		var v0 := -0x324;
		for i0 := v0 to v0 {
			var v1 := "li";
			var v2: map<int, string> := map[|(seq(abs(0x36e), i1  => ("pf")))[safeIndex(v0, |(seq(abs(0x36e), i1  => ("pf")))|) := seq(abs(0x157), i2  => ('t'))]| := v1];
			v2 := v2[v0 := if (f14) then v1 else v1];
			var v3 := DC0(f14);
			if ((if (v3.cf0) then v3 else v3).cf0) {
				var v4: seq<int> := [|v1|];
				var v5: multiset<int> := multiset{i0};
				var v6: multiset<seq<int>> := multiset{if (f14) then [fm0(f14, i0, |v4|, v5, globalState)] else [i0], v4, v4[safeIndex(v0, |v4|) := v0]};
				v6 := fm13(f15, 0x1e8 * i0, globalState);
				var v7: array<seq<array<bool>>> := new seq<array<bool>>[25];
				var v8: array<bool> := new bool[23];
				var v9: seq<array<bool>> := [v8, v8, v8];
				v7[safeIndex(777, v7.Length)] := v9;
				v7[safeIndex(777, v7.Length)] := [v8];
				var v10: map<bool, seq<int>> := map[multiset{false} == f15 := v4];
				v10 := v10[f14 := v4];
				var v11, v12, v13, v14 := m5(globalState);
				var v15: seq<bool> := [(seq(abs(0x200), i3  => ('g'))) == v1, true, v14, v14, v14];
				v15, v8, globalState.f9 := v15 + v15, v8, f14;
			} else {
				var v16: set<bool> := {f14};
				var v17: seq<set<bool>> := [v16];
				var v18: multiset<int> := multiset{i0};
				var v19: multiset<int> := multiset{fm0(f14, -|v18|, i0, v18, globalState), -0x14e, i0, -v0, i0};
				var v20: map<bool, bool> := map[v17 == [v16, v16] := v18 > v19];
				var v21: map<map<int, int>, int> := map[fm14(globalState) := v0];
				var v22: seq<map<map<int, int>, int>> := [v21];
				v20 := v20[!fm12(f14, v22[safeIndex(v0, |v22|)], globalState) := f14];
				var v23: map<string, D0> := map[v1 := DC0(f14)];
				var v24: seq<bool> := [f14, fm12(f14, fm48(i0, v0, f14, globalState), globalState), f14];
				var v25: T0 := new C5(v0, true, multiset(v24));
				var v26: map<T0, string> := map[v25 := v1];
				var v27: array<bool> := new bool[8];
				v23 := v23[(if (v25 in v26) then v26[v25] else DC13(v1, v0, f14, v27).cf22) + v1 := v3];
				var v28: map<string, int> := map["jvdqndys" := i0];
				r3 := (v1 + v1) !in v28;
				r0 := v0;
				var v29: seq<map<bool, bool>> := [v20];
				var v30: multiset<map<bool, bool>> := multiset{v20};
				var v31: seq<multiset<map<bool, bool>>> := [v30];
				var v32 := 'b';
				var v33: map<int, char> := map[i0 := v32];
				var v34: array<multiset<map<bool, bool>>> := new multiset<map<bool, bool>>[19] [multiset(v29) + v30, v30 * v30, v30, v30, v30, v31[safeIndex(|v33|, |v31|)], v30, multiset(v29[safeIndex(|v16|, |v29|) := v20]), v30, v30 - v30, v30, v30, v30, v30, v30, v30, v30, multiset{v20, v20}, v30[map[false := v25.f14] := abs(|v24|)]];
				v34[safeIndex(765, v34.Length)] := v30 + v30;
				v34[safeIndex(765, v34.Length)] := multiset(v29);
			}
			
			if (f14) {
				var v35: array<C0> := new C0[27];
				v35 := v35;
				globalState.f9 := f14;
				var v36: seq<int> := [v0];
				globalState.f3 := safeModuloInt(i0, |v36|);
				var v37: set<multiset<bool>> := {f15};
				var v38: map<D11, int> := map[DC33(v37) := v0];
				var v39 := DC33(v37);
				v38 := v38[v39 := -i0 * v0];
				var v40 := new C1();
			} else {
				var v41: array<int> := new int[2] [i0, v0];
				v41[safeIndex(4, v41.Length)] := v0;
				var v42: array<char> := new char[12];
				var v43 := 't';
				v42[safeIndex(347, v42.Length)] := v43;
				globalState.f3, v41[safeIndex(4, v41.Length)], v42[safeIndex(347, v42.Length)], v0 := v0 + i0, i0, v43, -i0;
				var v45: array<bool> := new bool[3](i4 => (set v44 : int | v44 in {i0} :: (v44 - 0x112)) !! {v0});
				v45[safeIndex(279, v45.Length)] := true;
				var v46: seq<bool> := [f14, f14 || true, f14];
				var v47: set<bool> := {f14};
				r3, r1, v45[safeIndex(279, v45.Length)] := f14, v46[safeIndex(|(if (f14) then {f14, true} else v47)|, |v46|)], f14;
				var v48: map<char, int> := map[v42[safeIndex(347, v42.Length)] := 0x37b];
				var v49: seq<int> := [v0, i0];
				var v50: set<seq<int>> := {v49, v49, v49[safeIndex(v41[safeIndex(4, v41.Length)], |v49|) := v41[safeIndex(4, v41.Length)]], v49};
				globalState.f9 := !(if (true) then fm23(-v0, f14, v48, v50, globalState) else false);
				var v51, v52, v53, v54 := m5(globalState);
				v1 := [v43] + v1;
			}
			
			var v55: array<int> := new int[1];
			var v56: map<int, int> := map[v0 := v0];
			v55[safeIndex(518, v55.Length)] := |(v56 + (map v57 : int | (964 <= v57) && (v57 < -0x363) :: (v57 - |v1|) := (i0))[-v0 := i0])|;
			var v58: map<string, int> := map[v1 := -v0];
			var v59: multiset<int> := multiset{|v1|};
			var v60: seq<int> := [v0];
			var v61: set<seq<int>> := {[fm0(f14, i0, i0, v59, globalState)], v60, v60};
			v55[safeIndex(518, v55.Length)] := safeDivisionInt(if (v1 in v58) then v58[v1] else |DC41(v61).cf67|, i0);
		}
		var v62: array<set<int>> := new set<int>[19];
		var v63: set<int> := {v0};
		v62[safeIndex(906, v62.Length)] := v63;
		v62[safeIndex(906, v62.Length)] := v63;
		var v64 := 'h';
		var v65: multiset<char> := multiset{v64};
		v65 := v65 * (v65 - v65);
		var v66: array<bool> := new bool[10];
		v66[safeIndex(542, v66.Length)] := !f14;
		v66[safeIndex(542, v66.Length)] := f14;
		forall i5 | 0 <= i5 < v66.Length {
			v66[i5] := (v0 + v0) != v0;
		}
		var v67, v68, v69, v70 := m5(globalState);
		r0 := (v0 * v68) + |(map v71 : int | v71 in (seq(abs(0x16b), i6  => (v0))) :: (safeDivisionInt(v71, v68)) := (v0))|;
		var v72 := DC3(-v0);
		var v73 := DC6(v72);
		r1 := match if (true) then v73 else v73 {
			case DC4(cf4, cf5) => DC42('r', f14).cf69
			case DC5(cf6, cf7) => f14
			case DC3(cf3) => f14
			case DC6(cf8) => {v64} <= {v64}
		};
		var v74: map<bool, int> := map[v66[safeIndex(542, v66.Length)] := v67];
		var v75: map<map<int, bool>, int> := map[map[if (v66[safeIndex(542, v66.Length)] in v74) then v74[v66[safeIndex(542, v66.Length)]] else -v68 := v66[safeIndex(542, v66.Length)]] := v0];
		r2 := v75;
		r3 := v70;
	}
	method m1(p0: set<int>, p1: int, globalState: GlobalState) returns (r0: int, r1: int, r2: bool) {
		var v0: array<bool> := new bool[5](i0 => f14);
		var v1: seq<char> := ['i'];
		v0[safeIndex(538, v0.Length)] := |v1| > p1;
		v0[safeIndex(538, v0.Length)] := fm2(f14, f14, globalState);
		globalState.f9 := v0[safeIndex(538, v0.Length)];
		var v2: array<string> := new string[27];
		var v3: array<array<string>> := new array<string>[17] [v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, DC44(v2).cf71];
		v3[safeIndex(2, v3.Length)] := v2;
		v3[safeIndex(2, v3.Length)] := if (DC5(f14, true).cf7) then v2 else if (v0[safeIndex(538, v0.Length)]) then v2 else v2;
		var v4 := 'y';
		var v5: set<char> := {v4};
		var v6: multiset<int> := multiset{p1};
		v5 := fm32(v1, v6, globalState);
		var v7 := DC12(v0);
		var v8: seq<D4> := [v7];
		var v9 := new C4(v8 != v8, v0[safeIndex(538, v0.Length)], f15);
		if (v9.f19) {
			if (v9.f19) {
				v0 := v0;
				var v10: seq<string> := [v1, v1, v1 + v1, v1];
				v10 := v10;
				globalState.f9 := v9.f19;
				var v11: map<int, int> := map[p1 := |f15|];
				var v12: map<int, map<int, int>> := map[|v1| := v11];
				var v13: map<map<int, bool>, int> := map[map[p1 := !!v9.f19] := p1];
				var v14 := DC21(v13);
				var v15: map<D7, set<char>> := map[v14 := v5];
				var v16: seq<map<int, map<int, int>>> := [v12, v12];
				v5, globalState.f3, r2, v12 := (if (v14 in v15) then v15[v14] else v5) * {'f'}, p1, !(p1 <= |[-fm0(v9.f19, -p1, p1, v6, globalState), p1]|), v16[safeIndex(safeModuloInt(p1, p1), |v16|)];
				var v17 := new C1();
			} else {
				var v18 := DC13(v1, p1, v0[safeIndex(538, v0.Length)], v0);
				v1, v1 := v18.cf22 + (seq(abs(0x124), i1  => (v4))), v1;
				var v19: map<bool, int> := map[v9.f19 := p1];
				var v20: array<char> := new char[19] [v4, v4, v4, v4, v4, v4, v4, v1[safeIndex(|v6|, |v1|)], v4, v4, fm25(p1, 0x81, f14, v19, globalState), v4, v4, 't', 'g', v4, v4, v4, v4];
				v20[safeIndex(774, v20.Length)] := v4;
				v20[safeIndex(774, v20.Length)] := v4;
				var v21: array<set<D8>> := new set<D8>[21];
				var v22: set<D8> := {DC25()};
				v21[safeIndex(949, v21.Length)] := v22;
				v21[safeIndex(949, v21.Length)] := v22;
				var v24: map<bool, set<int>> := map[v20[safeIndex(774, v20.Length)] !in v1 := p0 * (set v23 : int | (636 <= v23) && (v23 < 0xb5) :: (safeDivisionInt(v23, p1)))];
				v24 := v24;
				var v25: map<string, int> := map[v1 := |f15|];
				v25 := v25[v1 := |v1| - -p1];
			}
			
			var v26: map<bool, bool> := map[v9.f19 := !(p1 <= p1)];
			v26 := v26[f14 := p1 < p1];
			v1 := seq(abs(-0x322), i2  => (v4));
			var v27: array<set<int>> := new set<int>[5] [p0, p0, p0, p0, p0 + p0];
			var v28: map<int, bool> := map[p1 := v9.f19];
			var v29: map<int, int> := map[|v28| := p1];
			var v30: seq<int> := [p1];
			var v31: seq<int> := [|v29|, p1, fm0(f14, p1, p1, v6, globalState), |v30|, p1];
			v27[safeIndex(338, v27.Length)] := {p1, p1} + (set v32 : int | v32 in v30 :: (v32 + 876));
			var v33: array<int> := new int[14];
			v33[safeIndex(453, v33.Length)] := if (p1 in v29) then v29[p1] else p1;
			var v34: set<string> := {v1};
			v27[safeIndex(338, v27.Length)], r1, v33[safeIndex(453, v33.Length)], v1, v34 := {p1}, p1, p1, seq(abs(-551), i3  => (v4)), v34;
			var v35: map<map<bool, int>, array<int>> := map[map[v0[safeIndex(538, v0.Length)] := v33[safeIndex(453, v33.Length)]] := v33];
			var v36: map<bool, int> := map[f14 := v33[safeIndex(453, v33.Length)]];
			var v37: seq<array<int>> := [if (v36 in v35) then v35[v36] else v33, v33];
			var v38 := DC16(v33);
			v0[safeIndex(538, v0.Length)], v37, r1, globalState.f9 := DC0(v0[safeIndex(538, v0.Length)]).cf0, [v33, v33, v33, v38.cf29, v33], v33[safeIndex(453, v33.Length)], !true;
		} else {
			v4 := v4;
			var v39: array<array<C2>> := new array<C2>[26];
			v39 := new array<C2>[4];
			r1 := p1 - p1;
			if (v0[safeIndex(538, v0.Length)]) {
				var v40: array<int> := new int[3](i4 => i4 - p1);
				v40[safeIndex(831, v40.Length)] := p1;
				v40[safeIndex(831, v40.Length)] := p1;
				v0[safeIndex(538, v0.Length)] := v0[safeIndex(538, v0.Length)];
				var v41: map<int, string> := map[|(seq(abs(0x12f), i5  => (v4)))| := if (v9.f19) then v1 else v1];
				v41 := v41[|v1| - |v1| := if (v0[safeIndex(538, v0.Length)]) then "aupfiojbx" else v1];
				var v42: map<bool, int> := map[v0[safeIndex(538, v0.Length)] := v40[safeIndex(831, v40.Length)]];
				var v43: map<char, int> := map[v4 := 720];
				var v44: seq<int> := [|map[v9.f19 := v4]|, -p1];
				var v45: set<seq<int>> := {v44, v44};
				var v46 := DC41(v45);
				var v47: map<int, bool> := map[if (v9.f19 in v42) then v42[v9.f19] else v40[safeIndex(831, v40.Length)] := fm23(p1, v9.f19, v43, v46.cf67, globalState)];
				var v48 := new C2(v47);
				var v49 := new C5(v40[safeIndex(831, v40.Length)] * p1, v40[safeIndex(831, v40.Length)] > p1, f15 + f15);
			} else {
				v9.m6(globalState);
				var v50: map<int, bool> := map[|v6| := f14];
				var v51: seq<bool> := [v9.f19];
				var v52: C2 := new C2(v50 + v50[|[v51]| := v9.f19]);
				v52 := v52;
				var v53: seq<int> := [p1, p1];
				var v54: C0 := new C0(f14, v9.f19);
				var v55: set<C0> := {v54};
				var v56: seq<set<C0>> := [v55, v55];
				v0[safeIndex(538, v0.Length)] := p1 != safeDivisionInt(--v53[safeIndex(|v1|, |v53|)], |v56|);
				var v57: C3 := new C3(0x168, v1, v54.f23, f15);
				var v58: map<C3, int> := map[v57 := |"tfija"|];
				var v59 := DC47(v58);
				r1 := DC11(-p1, v9.fm17(|v59.cf76|, f14, multiset{v51}, v57.f20, globalState), v0[safeIndex(538, v0.Length)], v57.f20).cf17;
				v53 := v53;
			}
			
			v4 := v4;
		}
		
		var v60: set<bool> := {v9.f19, true, v9.f19, true, true};
		r0 := if (p1 >= (if (|v60| in v6) then v6[|v60|] else p1)) then p1 else p1;
		r1 := (p1 + p1) + |(v1 + v1)|;
		var v61: C0 := new C0(v0[safeIndex(538, v0.Length)], !true);
		var v62: map<C0, set<bool>> := map[v61 := v60];
		r2 := v62 == v62;
	}
	method m5(globalState: GlobalState) returns (r0: int, r1: int, r2: multiset<int>, r3: bool) {
		var v0: array<bool> := new bool[19](i1 => f14);
		forall i0 | 0 <= i0 < v0.Length {
			v0[i0] := safeModuloInt(-0x140, |"vuy"|) != |((seq(abs(-0x130), i2  => ('j'))) + "qdea")|;
		}
		if (f14) {
			r3 := f14;
			r3 := f14;
			var v1 := DC0(f14);
			match (v1) {
				case DC1(cf1) =>
					var v2 := -0x3d5;
					r0 := v2;
					var v3 := "ikpk";
					var v4: map<bool, string> := map[cf1 := v3];
					var v5: seq<string> := [v3, v3, v3, v3, v3];
					var v6 := 'f';
					var v7: set<int> := {v2};
					var v8 := DC30(v6, cf1, |(seq(abs(875), i3  => (v6)))|, v7, v2);
					var v9 := new C3(|(v4 + v4)|, v5[safeIndex(v2, |v5|)], if (cf1) then v8.cf51 else cf1, f15);
					var v10 := new C3(|[v9.f20, v9.f20, v9.f20]|, v3 + v3, true, multiset([f14, !cf1, f14]));
					var v11: seq<int> := [v10.f20, v2];
					r0 := |(v11 + (v11 + v11))|;
				case DC0(cf0) =>
					var v12 := 0x53;
					var v13 := new C5(v12 + v12, f14, fm33(f14, globalState));
					var v14 := 'r';
					var v15: map<char, int> := map[v14 := 0x42];
					var v16: seq<int> := [v13.f18];
					var v17: set<seq<int>> := {v16, seq(abs(0x125), i4  => (v13.f18)), v16, v16, v16};
					var v18 := new C4(cf0, !fm23(v12, cf0, v15, v17, globalState) <== cf0, f15);
					var v19: map<int, bool> := map[v12 := f14];
					var v20 := new C2(v19);
					globalState.f3 := v12;
				case DC2(cf2) =>
					var v21: set<bool> := {f14};
					r1 := |v21|;
					var v22 := 999;
					var v23 := DC38(v22);
					v23 := DC38(v22);
					globalState.f2 := v0;
					var v24: array<D8> := new D8[9](i5 => DC24(v21));
					var v25 := DC24(v21);
					v24[safeIndex(774, v24.Length)] := v25.(cf45 := {f14});
					v24[safeIndex(774, v24.Length)] := v25;
			}
			
			var v26: array<string> := new string[11](i6 => "tqux" + "mdsulisdp");
			var v27: set<multiset<bool>> := {f15};
			var v28 := DC33(v27);
			v26, v28 := v26, DC33(v27).(cf57 := v27);
			var v29 := new C1();
		} else {
			var v30 := 0x8c;
			var v31: map<int, int> := map[v30 := 0x29];
			v31 := v31[v30 := v30 * v30];
			var v32 := "shmrn";
			if (f14 ==> !(v32 <= v32)) {
				var v33 := new C0(f14, f14);
				v30 := 0x3;
				var v34: set<string> := {fm27(v30, f14, globalState)};
				v34 := (v34 + v34) - (if (v33.f24) then v34 else v34);
				var v35 := 'n';
				v35 := v35;
				var v36: multiset<int> := multiset{v30};
				var v37: seq<bool> := [v33.f23];
				var v38: set<multiset<bool>> := {(multiset{f14, v33.f24})[f14 := abs(|(seq(abs(0x270), i7  => (v35)))|)], f15, multiset{v33.f23}, multiset(v37)};
				globalState.f3, r0, globalState.f9 := v30, (|v32| - |v36|) + (v30 - |v32|), v38 !! v38;
			} else {
				var v39: map<bool, int> := map[false := v30];
				var v40 := 'w';
				var v41: set<char> := {fm25(v30, v30, f14, v39, globalState), v40};
				var v42: seq<set<char>> := [v41];
				var v44: map<char, bool> := map['e' := f14];
				var v45: seq<int> := [v30, v30];
				var v46: set<seq<int>> := {v45};
				globalState.f3, globalState.f9, v30, v42 := --(v30 + 0x292), !(fm23(v30, f14, map v43 : char | v43 in v44 :: (v43) := (|[v30, v30]|), v46, globalState) && (f14 <== f14)), if (|v32| in v31) then v31[|v32|] else v30, v42;
				var v47: set<bool> := {!f14};
				var v48: seq<bool> := [f14, fm25(v30, |v47|, f14, v39, globalState) in "dpidlr", f14, f14];
				v48 := v48;
				var v49 := DC1(f14);
				var v50 := DC2(v49);
				var v51: map<D0, bool> := map[v50 := f14];
				var v52: map<int, bool> := map[v30 := fm11(v51, v30, globalState)];
				var v53: seq<map<int, bool>> := [v52];
				var v54 := new C2(if (true) then v53[safeIndex(v30, |v53|)] else v52);
				var v55: multiset<int> := multiset{v30, v30};
				r0 := fm0(f14, v30, v30, v55, globalState);
				var v56: array<int> := new int[18];
				v56[safeIndex(407, v56.Length)] := v30;
				v56[safeIndex(407, v56.Length)] := v30 + v30;
			}
			
			var v57: seq<bool> := [f14];
			var v58: seq<seq<bool>> := [v57];
			if (([true, true, f14, f14] + v58[safeIndex(0x28e, |v58|)]) != v57[safeIndex(v30, |v57|) := f14]) {
				var v59: multiset<int> := multiset{v30, 943, v30, 97};
				r1 := if (safeModuloInt(v30, fm0(f14, |v32|, v30, v59, globalState)) in v59) then v59[safeModuloInt(v30, fm0(f14, |v32|, v30, v59, globalState))] else -v30;
				var v60 := new C3(safeModuloInt(v30, v30), v32, f14 ==> f14, f15[f14 := abs(v30)]);
				v32 := "srr";
				var v61: map<int, bool> := map[v60.f20 := if (f14) then f14 else f14];
				v61 := v61[if (f14) then 0x2e1 else v30 := true];
				var v62 := DC18(f14);
				globalState.f9 := v62.cf33;
			} else {
				globalState.f9 := -(v30 - -320) < v30;
				var v63 := DC5(f14, f14);
				var v64: C5 := new C5(|v57|, v63.cf6, f15);
				var v65 := DC48(v30, v64, 0x1eb);
				globalState.f3 := v65.cf79;
				var v66: map<array<bool>, bool> := map[v0 := f14];
				v0[safeIndex(935, v0.Length)] := !(if (v0 in v66) then v66[v0] else f14);
				v0[safeIndex(935, v0.Length)] := f14;
				var v67 := DC0(v0[safeIndex(935, v0.Length)]);
				globalState.f3 := if (v67.cf0) then v30 else -v64.f18;
				var v68 := DC23(f14, v30, f14);
				v68 := v68;
			}
			
			globalState.f2 := v0;
			var v69 := 'o';
			var v70 := DC9(f14, v30, 'v');
			globalState.f9 := (if (f14) then DC9(f14, v30, v69) else v70).cf13;
		}
		
		var v71: array<string> := new string[6];
		var v72 := "hau";
		v71[safeIndex(202, v71.Length)] := v72;
		var v73 := 0x1b5;
		v71[safeIndex(202, v71.Length)] := (v72 + (seq(abs(0x33), i8  => ('y')))) + (fm27(v73, f14, globalState) + v72);
		var v74: set<int> := {v73};
		for i9 := |v74| to v73 {
			var v75 := DC18(!true);
			var v76: map<int, D5> := map[i9 := v75];
			v76 := v76[|v74| - v73 := v75];
			v0[safeIndex(686, v0.Length)] := if (f14) then f14 else !f14;
			v0[safeIndex(129, v0.Length)] := f14;
			var v77 := 'h';
			v0[safeIndex(686, v0.Length)], v0[safeIndex(129, v0.Length)], v71[safeIndex(202, v71.Length)] := true, f14, v72[safeIndex(v73, |v72|) := v77] + v71[safeIndex(202, v71.Length)];
			var v78: seq<bool> := [true, f14];
			globalState.f3 := |((v78 + v78) + v78)|;
			var v79: map<int, int> := map[i9 := i9];
			var v80 := DC29(v79);
			var v81 := DC31(v80);
			var v82 := DC31(v81);
			var v83: seq<D9> := [v82];
			var v84: set<seq<D9>> := {v83};
			v84 := fm49(v73, globalState);
		}
		var v85: array<int> := new int[22](i10 => safeModuloInt(i10, v73));
		v85[safeIndex(239, v85.Length)] := v73;
		v85[safeIndex(239, v85.Length)] := -safeModuloInt(if (f14) then v73 else v73, if (f14) then 284 else 0x10d);
		var v86: seq<int> := [v85[safeIndex(239, v85.Length)]];
		var v87: map<int, bool> := map[|(v86 + [v85[safeIndex(239, v85.Length)]])| := f14];
		v87 := v87[v73 := v85[safeIndex(239, v85.Length)] <= 0x1fe];
		r0 := v85[safeIndex(239, v85.Length)];
		r1 := v73 * v73;
		var v88: multiset<int> := multiset{v73, |v72|, v73, safeModuloInt(v73, v85[safeIndex(239, v85.Length)])};
		r2 := v88;
		var v90: seq<map<int, bool>> := [v87];
		var v91 := DC21(map v89 : map<int, bool> | v89 in v90 :: (v89) := (v73));
		r3 := match v91 {
			case DC22(cf37, cf38, cf39, cf40, cf41) => !f14
			case DC23(cf42, cf43, cf44) => cf44
			case DC21(cf36) => f14
		};
	}
}

class C7 {
	constructor () {
	}
	
	function fm9(p0: int, globalState: GlobalState): int {
		|("kko" + "ddx")|
	}
	function fm10(p0: map<bool, int>, p1: multiset<int>, p2: bool, p3: string, globalState: GlobalState): bool {
		false
	}
	method m3(p0: int, globalState: GlobalState) {
		for i0 := p0 to p0 {
			var v0 := DC3(-181);
			var v1: map<D1, int> := map[v0 := i0];
			var v2: seq<int> := [|v1|, p0];
			var v3 := 'j';
			var v4: seq<char> := [v3];
			var v5: array<seq<int>> := new seq<int>[15] [v2, v2 + v2, v2, v2, v2, v2[safeIndex(i0, |v2|) := v0.cf3], [|v4|, p0], v2, v2, if (true) then v2 else v2, v2[safeIndex(i0, |v2|) := i0], v2 + [p0, p0, |v4|], v2, v2, v2];
			v5[safeIndex(789, v5.Length)] := v2[safeIndex(i0, |v2|) := p0];
			var v6: array<bool> := new bool[26];
			var v7 := true;
			v6[safeIndex(752, v6.Length)] := v7;
			globalState.f9, v5[safeIndex(789, v5.Length)], v6[safeIndex(752, v6.Length)] := v7, (v2 + v2) + (v2 + v2), !v7;
			var v8: map<int, char> := map[p0 := v3];
			var v9: multiset<map<int, char>> := multiset{v8};
			var v10: array<map<int, int>> := new map<int, int>[2];
			var v11: map<seq<int>, char> := map[seq(abs(414), i1  => (|[v6[safeIndex(752, v6.Length)]]|)) := v3];
			var v12: map<bool, array<bool>> := map[v7 := v6];
			v9, v7, globalState.f2, v10, v11 := v9 - v9, !v7, if (v6[safeIndex(752, v6.Length)] in v12) then v12[v6[safeIndex(752, v6.Length)]] else v6, v10, map[v5[safeIndex(789, v5.Length)] := v3];
			v4 := v4 + "reyaln";
			globalState.f9 := i0 != i0;
		}
		for i2 := -0x7e to p0 {
			var v13: array<D1> := new D1[2];
			var v14 := DC4(0x23d, p0);
			v13[safeIndex(84, v13.Length)] := v14.(cf4 := i2);
			v13[safeIndex(84, v13.Length)] := v14;
			var v15 := true;
			var v16 := new C0(v15, v15);
			var v17: map<int, int> := map[-0xbd := p0];
			globalState.f9 := v17 != fm14(globalState);
			var v18 := 'o';
			var v19: map<char, int> := map[v18 := -i2];
			var v20: seq<int> := [i2, i2];
			var v21: set<seq<int>> := {v20, v20};
			var v22: array<bool> := new bool[6](i3 => v16.f24);
			globalState.f2 := if (!(if (fm23(p0, v15, v19, v21, globalState)) then v15 else false)) then v22 else v22;
		}
		var v23: multiset<int> := multiset{-p0, 704};
		globalState.f9 := !!((v23 * v23) !! v23);
		var i4 := 0;
		while (p0 != p0)
			decreases 100 - i4
		{
			if (i4 >= 100) {
				break;
			}
			
			i4 := i4 + 1;
			var v24: array<bool> := new bool[14];
			var v25: seq<array<bool>> := [v24, v24, v24, v24, v24];
			var v26 := false;
			v25 := if (v26) then v25 else v25;
			v26 := v26;
			var v27: seq<bool> := [v26];
			var v28: seq<int> := [p0, |v27[safeIndex(p0, |v27|) := !v26]|, -573, |map[v26 := v26]|, -p0];
			var v29 := DC32((seq(abs(0x393), i5  => (0x1fa))) + v28);
			match (v29) {
				case DC32(cf56) =>
					var v30: array<char> := new char[8];
					var v31: seq<array<char>> := [v30];
					var v32 := DC51(v30);
					v31 := [v30, v32.cf83];
					var v33: map<D10, int> := map[v29 := p0];
					var v34: multiset<map<D10, int>> := multiset{v33};
					globalState.f9 := v33 in v34;
					var v35: set<bool> := {v26};
					globalState.f6 := v35 * fm41(p0, p0, globalState);
					var v36: map<int, bool> := map[p0 := !v26];
					var v37 := "o";
					var v38: multiset<bool> := multiset{v26};
					var v39: C3 := new C3(p0, v37, v26, v38);
					var v40: seq<C3> := [v39, v39, v39, v39, v39];
					var v41: map<seq<C3>, multiset<bool>> := map[v40 := v38];
					var v42: map<bool, map<seq<C3>, multiset<bool>>> := map[v26 := v41];
					var v43: map<int, map<seq<C3>, multiset<bool>>> := map[|v27| := v41];
					v36 := v36[-|(if (v26 in v42) then v42[v26] else if (p0 in v43) then v43[p0] else map[v40 := v38])[v40 := v38]| := v26];
			}
			
			var v44: set<seq<int>> := {[0x362, p0]};
			var v45 := DC41(v44);
			match (v45) {
				case DC42(cf68, cf69) =>
					var v46: set<bool> := {cf69, cf69, cf69};
					globalState.f6 := if (v26) then v46 else v46;
					var v47: array<int> := new int[4];
					v47[safeIndex(241, v47.Length)] := fm9(p0, globalState);
					v47[safeIndex(241, v47.Length)] := |(v25 + (v25 + v25))|;
					var v48: map<bool, seq<bool>> := map[cf69 := v27 + ([v26])[safeIndex(p0, |[v26]|) := cf69]];
					var v49: map<char, int> := map['h' := v47[safeIndex(241, v47.Length)]];
					var v50: map<bool, int> := map[cf69 := 0xca];
					var v51 := "jrpyg";
					v48 := v48[fm23(v47[safeIndex(241, v47.Length)], false, v49, {v28, v28, v28}, globalState) := if (v26 in v48) then v48[v26] else [v26, cf69, !fm10(v50, v23, false, v51, globalState)]];
					var v52: seq<string> := ["sct"];
					var v53 := DC13(seq(abs(326), i8  => (cf68)), |(seq(abs(-776), i9  => (DC42(cf68, DC30(cf68, cf69, p0, {v47[safeIndex(241, v47.Length)]}, v47[safeIndex(241, v47.Length)]).cf51).cf68)))|, v26, v24);
					var v54: array<string> := new string[23] [v52[safeIndex(-p0, |v52|)] + v51, v51, v51, v51 + v51, v51, v51 + v51, v51, v51 + v51, v51[safeIndex(v47[safeIndex(241, v47.Length)], |v51|) := cf68], seq(abs(0xda), i6  => (cf68)), v51, seq(abs(0x324), i7  => (cf68)), v51, v51, v53.cf22, seq(abs(325), i10  => (cf68)), v51, v51, "abhvu", v51 + "u", v51, if (cf69) then v51 else seq(abs(0x186), i11  => ('w')), v51];
					v54[safeIndex(243, v54.Length)] := "i";
					v54[safeIndex(243, v54.Length)] := v51 + "k";
				case DC41(cf67) =>
					var v55 := "gv";
					var v56: map<int, int> := map[p0 := |v55|];
					globalState.f3 := safeDivisionInt(if (v26) then p0 else p0, |(v56 + (map v57 : int | v57 in v23 :: (safeModuloInt(v57, |v55|)) := (p0))[p0 := p0])|);
					var v58 := 'f';
					var v59: array<int> := new int[13];
					var v60: map<array<int>, bool> := map[v59 := v26];
					var v61: set<int> := {fm0(v26, 418, |v60|, v23, globalState), |v55|};
					var v62: set<D9> := {DC30(v58, v26, p0, v61, p0)};
					var v63: map<char, int> := map[v58 := p0];
					var v64 := DC1(v26);
					v26, globalState.f2 := DC30(v58, v26, p0, v61, p0) in v62, if (fm23(p0, v26, v63, fm50(p0, v64.cf1, p0, p0, globalState), globalState)) then v24 else v24;
					var v65: array<array<bool>> := new array<bool>[13] [v24, v24, v24, v24, v24, v24, v24, v24, v24, v24, v24, v24, v24];
					var v66 := DC54(v65);
					var v67: array<array<array<bool>>> := new array<array<bool>>[3] [if (v26) then v65 else v65, v66.cf90, v65];
					v67[safeIndex(597, v67.Length)] := if (v26) then v65 else v65;
					v67[safeIndex(597, v67.Length)] := v65;
					v59[safeIndex(857, v59.Length)] := p0;
					var v68: map<int, string> := map[p0 := "itn"];
					v59[safeIndex(857, v59.Length)] := safeModuloInt(-0x304, |v68|);
				case DC43(cf70) =>
					var v69 := new C0(v26, !v26);
					globalState.f9 := v69.f24;
					globalState.f9 := (0x35 - 0xaf) > p0;
					var v70: array<string> := new string[10];
					var v71: array<set<int>> := new set<int>[10];
					var v72: map<bool, bool> := map[v26 := v69.f23];
					globalState.f11, v70 := v71, if (if (v26 in v72) then v72[v26] else v26) then if (false) then v70 else v70 else v70;
			}
			
		}
		var v73 := "g";
		var v74 := DC2(fm20(p0, |v73|, p0, globalState));
		var v75 := DC2(v74);
		if (match v75 {
			case DC1(cf1) => !!(0x3b8 in [|(seq(abs(234), i12  => ('v')))|, p0])
			case DC0(cf0) => cf0
			case DC2(cf2) => map[true := false] != map[false := !false]
		}) {
			var v76 := false;
			var v77: map<string, bool> := map[v73 := v76];
			v77 := v77;
			globalState.f9 := v76 || v76;
			var v78: multiset<bool> := multiset{v76, v76};
			var v79 := new C5(-p0, |v73| <= p0, v78);
			var v80: map<bool, multiset<bool>> := map[true := v78];
			var v81: map<multiset<bool>, bool> := map[if (!v76 in v80) then v80[!v76] else fm33(v76, globalState) := if (v76) then v76 else v76];
			v81 := v81[v78 := false];
			v76 := safeModuloInt(v79.f18, p0) < |v73|;
		} else {
			globalState.f3 := p0;
			var v82 := 'l';
			var v83 := DC9(true, p0, v82);
			globalState.f9 := v83.cf13;
			var v84 := true;
			var v85: map<string, bool> := map[seq(abs(0x199), i13  => (v82)) := v84];
			v85 := (v85 + v85) + v85;
			var v86: multiset<bool> := multiset{v84, v84};
			var v87 := new C4(v84, !v84, v86 - v86);
			globalState.f9 := v87.f19;
		}
		
		var v88 := true;
		if (v88) {
			var v89: multiset<bool> := multiset{false};
			var v90: map<bool, bool> := map[v88 := v88];
			var v91: array<bool> := new bool[15](i14 => v88);
			var v92 := DC34(v89, multiset{|v90|}, v91, DC5(v88, false));
			var v93: multiset<array<bool>> := multiset{v92.cf60};
			var v94 := 'x';
			var v95: map<int, char> := map[p0 - |v93| := v94];
			var v96: seq<int> := [p0, |v89|];
			v95 := v95[|v96| := v94];
			var v97: C0 := new C0(v88, v88);
			var v98: map<int, C0> := map[p0 := v97];
			v98 := v98;
			var v99 := DC45(v90, v97, p0);
			var v100 := DC46(v99);
			var v101: array<D14> := new D14[26] [v100, v100, v100.(cf75 := v99), v100, v100, v100, v100, v100, v100, v100, v100, v100, v100, v100, v100, v100, v100, DC46(v99), v100, v100, v100, if (v88) then v100 else v100, v100, v100, v100.(cf75 := v99), v100];
			v101[safeIndex(642, v101.Length)] := v100;
			globalState.f9, v101[safeIndex(642, v101.Length)], v73, globalState.f3 := v97.f23, if (v97.f23) then v100 else v100, fm27(p0, v88, globalState), p0;
			var v102: seq<bool> := [v97.f23];
			var v103 := DC7(v102);
			if ([true] == v103.cf9[safeIndex(p0, |v103.cf9|) := v97.f23]) {
				var v104: array<map<array<D11>, bool>> := new map<array<D11>, bool>[8];
				var v106: array<D11> := new D11[24](i15 => DC33(set v105 : multiset<bool> | v105 in [v89, v89, multiset(v102)] :: (v105)));
				var v107: map<char, int> := map[v94 := -p0];
				var v108: set<seq<int>> := {v96};
				var v109: map<array<D11>, bool> := map[v106 := fm23(p0, v97.f23, v107, v108, globalState)];
				v104[safeIndex(440, v104.Length)] := v109;
				v104[safeIndex(440, v104.Length)] := v109 + map[v106 := v97.f23];
				globalState.f9 := (|"yacasvvoi"| > p0) ==> v88;
				var v110: map<int, int> := map[safeModuloInt(p0, p0) := p0];
				v110 := v110 + v110;
				var v111 := DC42(v94, v97.f24);
				var v112 := DC43(v111);
				var v113 := DC43(DC43(v111));
				var v114 := DC43(v113);
				var v115 := DC43(v112);
				var v116: set<int> := {p0};
				var v118: map<int, set<int>> := map[p0 := {p0}];
				var v119: seq<string> := [v73];
				var v120 := DC45(v90, v97, |v119[safeIndex(p0, |v119|)]|);
				var v124: map<int, bool> := map[p0 := v97.f24];
				var v126: seq<map<int, bool>> := [v124, (map v125 : int | (0xce <= v125) && (v125 < 0x28) :: (safeDivisionInt(v125, -0x362)) := (v97.f23))[p0 := true]];
				var v128: array<int> := new int[18];
				var v130: array<set<int>> := new set<int>[19] [v116, v116, set v117 : int | (0x30a <= v117) && (v117 < 0x254) :: (v117 * p0), if (v97.f23) then if (|v120.cf72| in v118) then v118[|v120.cf72|] else v116 else v116, set v121 : int | (0x197 <= v121) && (v121 < -0x1b8) :: (safeModuloInt(v121, p0)), v116, v116, v116, set v122 : int | (-0xf3 <= v122) && (v122 < 0x2c9) :: (v122 * p0), (set v123 : int | (0x1b5 <= v123) && (v123 < -919) :: (v123 - |{v97.f23}|)) - v116, fm22(|v73|, v126, v97.f23, v97.f24, globalState), set v127 : int | (-200 <= v127) && (v127 < 0x68) :: (v127 + |v23|), v116, v116, v116, {p0, |v126|, fm0(v88, |multiset{v128}|, |[v97.f23, v97.f24]|, v23, globalState), p0, p0}, set v129 : int | (0x137 <= v129) && (v129 < 0x260) :: (v129 * |[v97.f24, v97.f24]|), v116, v116];
				globalState.f11, globalState.f9, globalState.f3, v115 := v130, v97.f23, p0, fm51(v88, v97.f23, v88, |v124|, globalState).(cf70 := v114);
				v73 := seq(abs(-0x143), i16  => (v73[safeIndex(|[v97.f24]|, |v73|)]));
			} else {
				var v131: array<int> := new int[22](i17 => i17 - p0);
				v131[safeIndex(927, v131.Length)] := p0;
				v131[safeIndex(927, v131.Length)] := p0;
				var v132 := new C3(safeModuloInt(845, p0), v73 + v73, false, v89 + v89);
				var v134 := DC30(v94, v97.f23, |v102|, set v133 : int | v133 in v96 :: (safeModuloInt(v133, 0x3d3)), v131[safeIndex(927, v131.Length)]);
				globalState.f9 := v134.cf51;
				v131[safeIndex(927, v131.Length)] := v131[safeIndex(927, v131.Length)];
				v91[safeIndex(28, v91.Length)] := v97.f23 && v97.f24;
				v91[safeIndex(28, v91.Length)] := v97.f24;
			}
			
			globalState.f9 := v97.f24;
		} else {
			globalState.f9 := v88;
			var v135: map<char, bool> := map['x' := !!v88];
			var v136: seq<map<char, bool>> := [v135, v135 + fm46(|v73|, v88, globalState), v135, v135, v135];
			globalState.f3 := |v136|;
			var v137: array<seq<char>> := new seq<char>[20];
			v137[safeIndex(984, v137.Length)] := v73;
			var v138 := 'e';
			var v139: array<bool> := new bool[14](i18 => v88);
			var v140: seq<array<bool>> := [v139];
			var v141: multiset<array<bool>> := multiset{v140[safeIndex(p0, |v140|)]};
			v137[safeIndex(984, v137.Length)] := (((fm27(p0, v88, globalState))[safeIndex(p0, |fm27(p0, v88, globalState)|) := v138])[safeIndex(-0x32, |(fm27(p0, v88, globalState))[safeIndex(p0, |fm27(p0, v88, globalState)|) := v138]|) := v138])[safeIndex(|v141|, |((fm27(p0, v88, globalState))[safeIndex(p0, |fm27(p0, v88, globalState)|) := v138])[safeIndex(-0x32, |(fm27(p0, v88, globalState))[safeIndex(p0, |fm27(p0, v88, globalState)|) := v138]|) := v138]|) := v138];
			var v142: array<int> := new int[8];
			v142[safeIndex(4, v142.Length)] := p0;
			v142[safeIndex(4, v142.Length)] := (|v137[safeIndex(984, v137.Length)]| + p0) - p0;
			v138 := v138;
		}
		
	}
	method m4(p0: bool, p1: int, p2: int, p3: set<array<char>>, globalState: GlobalState) returns (r0: bool) {
		globalState.f9 := p0;
		var v0: array<int> := new int[27](i1 => safeDivisionInt(i1, 0x96));
		forall i0 | 0 <= i0 < v0.Length {
			v0[i0] := i0 + DC11(0x1cf, p2, p0, |map[p0 := p2]|).cf18;
		}
		var v1: set<bool> := {p0, p0};
		var v2: map<set<bool>, bool> := map[v1 + fm41(p2, p1, globalState) := p0];
		globalState.f3 := |v2|;
		var v3: array<bool> := new bool[2] [false, p0];
		v3[safeIndex(262, v3.Length)] := p0;
		v3[safeIndex(262, v3.Length)] := p0;
		var v4: map<bool, bool> := map[p0 := v3[safeIndex(262, v3.Length)]];
		v0[safeIndex(107, v0.Length)] := -p1;
		v4, v3[safeIndex(262, v3.Length)], v0[safeIndex(107, v0.Length)], globalState.f3 := v4, p1 != -(p1 - p2), p1, p2;
		var v5 := 'c';
		var v6: multiset<int> := multiset{v0[safeIndex(107, v0.Length)], 0x1b9};
		var v7: multiset<bool> := multiset{false};
		var v8: T0 := new C5(v0[safeIndex(107, v0.Length)], p0, v7);
		var v9: multiset<T0> := multiset{v8, v8};
		var v10 := "nnbtloi";
		var v11: set<int> := {v0[safeIndex(107, v0.Length)], fm0(v3[safeIndex(262, v3.Length)], v0[safeIndex(107, v0.Length)], v0[safeIndex(107, v0.Length)], v6[|v9| := abs(p1)], globalState), |v10|, p1};
		globalState.f9, v0[safeIndex(107, v0.Length)], v5, v0[safeIndex(107, v0.Length)] := v3[safeIndex(262, v3.Length)], safeModuloInt(v0[safeIndex(107, v0.Length)], |v11|), v5, -p2;
		var v12: seq<bool> := [p0, true, true];
		r0 := [v3[safeIndex(262, v3.Length)]] < v12;
	}
}

class C8 extends T0 {
	const f16 : map<int, array<char>>
	const f17 : seq<map<bool, int>>
	constructor (f16 : map<int, array<char>>, f17 : seq<map<bool, int>>, f14 : bool, f15 : multiset<bool>) {
		this.f16 := f16;
		this.f17 := f17;
		this.f14 := f14;
		this.f15 := f15;
	}
	
	function fm2(p0: bool, p1: bool, globalState: GlobalState): bool {
		f14
	}
	function fm3(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): string {
		(if (f14) then seq(abs(-0x2e2), i0  => ('l')) else seq(abs(0x7b), i1  => ('l'))) + "jshowx"
	}
	function fm4(p0: map<bool, bool>, p1: bool, globalState: GlobalState): bool {
		f14
	}
	method m0(globalState: GlobalState) returns (r0: int, r1: bool, r2: map<map<int, bool>, int>, r3: bool) {
		var v0: array<bool> := new bool[8] [!f14, f14, f14, f14, f14, f14, f14, f14];
		var v1: multiset<array<bool>> := multiset{v0, v0};
		if (v0 !in v1) {
			var v2: set<bool> := {!f14, f14, f14, f14};
			globalState.f9 := v2 != v2;
			if (if (f14) then true else f14) {
				r3 := !f14;
				var v3 := "okwamfa";
				r3 := fm2(f14, v3 == v3, globalState);
				globalState.f13 := {f14};
				var v4 := 0x1af;
				globalState.f3 := -v4;
				var v5: seq<int> := [v4];
				var v6: seq<seq<int>> := [v5, v5, v5];
				v0[safeIndex(425, v0.Length)] := "ahnxhipsh" == v3;
				v6, v0[safeIndex(425, v0.Length)] := (if (f14) then v6 else v6) + v6, if (f15 == f15) then f14 else f14;
			} else {
				var v7: array<int> := new int[3];
				var v8 := 726;
				v7[safeIndex(68, v7.Length)] := fm0(f14, 0x145, v8, multiset{v8, 0x172, v8, v8, v8}, globalState);
				var v9: seq<int> := [-348];
				var v10: map<int, int> := map[v8 := v8];
				var v11: map<seq<int>, map<int, int>> := map[v9 := v10];
				var v12: multiset<int> := multiset{v8, |fm5(f14, globalState)|, |v11|};
				v7[safeIndex(68, v7.Length)] := fm0(f14, v8, v8, v12, globalState);
				var v13 := DC1(f14);
				var v14: seq<bool> := [v13.cf1];
				var v15: map<seq<int>, int> := map[v9 := |v14|];
				var v16: map<int, bool> := map[if (f14) then v7[safeIndex(68, v7.Length)] else if (v9 in v15) then v15[v9] else v7[safeIndex(68, v7.Length)] := f14 <==> f14];
				var v17 := DC0(f14);
				v16 := v16[safeModuloInt(v8, v8) := v17.cf0];
				var v18: map<bool, array<int>> := map[f14 := v7];
				var v19: set<int> := {v8, |v18|, v7[safeIndex(68, v7.Length)], v8};
				var v20 := "h";
				var v21 := m2(safeModuloInt(|map[f14 := |v19|]|, v8), v20 + "fmlv", f14, f14 ==> f14, globalState);
				var v23: map<string, bool> := map["mkl" := true];
				var v24: map<D0, bool> := map[DC0(false) := |(map v22 : string | v22 in v23 :: (v22) := (seq(abs(0x225), i0  => (v21))))| > |v20|];
				f14 := if (v17 in v24) then v24[v17] else |v20| > 0x108;
				var v25: map<bool, int> := map[f14 := 0x232];
				var v26: map<int, map<bool, int>> := map[v7[safeIndex(68, v7.Length)] := v25];
				v25 := if (-(0x3b1 - v7[safeIndex(68, v7.Length)]) in v26) then v26[-(0x3b1 - v7[safeIndex(68, v7.Length)])] else v25[f14 := |{f14, f14}|];
			}
			
			var v27 := -0x389;
			var v28: seq<int> := [if (true) then v27 else 0x3, v27];
			v28 := v28;
			r0 := v27;
			r3 := !(if (f14) then f14 else f14);
		} else {
			var v29 := "vo";
			var v30: array<D0> := new D0[17](i1 => DC0(f14));
			var v31: array<array<D0>> := new array<D0>[28] [v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30];
			var v32: map<string, array<array<D0>>> := map[v29 := v31];
			v32 := v32[v29 := v31];
			var v33 := 0x19b;
			var v34: seq<int> := [v33];
			var v35: set<int> := {|v34|, v33, v33};
			globalState.f3 := |(fm6(globalState) - v35)|;
			r0 := -607;
			var v36: array<string> := new string[8];
			var v37: map<bool, bool> := map[f14 := f14];
			var v38: map<map<bool, bool>, int> := map[if (!f14) then fm7(globalState) else v37 := v33];
			v36, v38, v29 := v36, map v39 : map<bool, bool> | v39 in (seq(abs(0x286), i2  => (v37[false := f14]))) :: (v39) := (v33), v29 + v29;
			var v40: array<int> := new int[15];
			var v41: map<int, array<int>> := map[v33 := v40];
			v41 := v41[v33 := v40];
		}
		
		var v42: set<bool> := {f14};
		var v43 := 0x2bf;
		var v44 := 'h';
		var v45: map<int, char> := map[v43 := v44];
		var v46: map<bool, map<int, char>> := map[{f14} !! v42 := v45];
		v46 := v46[f14 := map[v43 := v44]];
		var v47: seq<bool> := [f14, f14];
		var v48: map<seq<bool>, bool> := map[v47 := !f14];
		v48 := v48[v47 := fm2(f14, f14, globalState)];
		var v49: map<bool, bool> := map[f14 := !f14];
		var i3 := 0;
		while (fm4(v49, f14, globalState))
			decreases 100 - i3
		{
			if (i3 >= 100) {
				break;
			}
			
			i3 := i3 + 1;
			var v50: map<bool, int> := map[f14 := v43];
			var v51: multiset<int> := multiset{501};
			f14 := multiset{|v50|} > v51;
			var v52: map<multiset<int>, bool> := map[v51 - v51 := f14];
			v52 := v52;
			if (DC0(v47[safeIndex(v43, |v47|)]).cf0) {
				var v53 := DC4(v43, v43);
				globalState.f9 := -v43 == v53.cf4;
				var v54 := "wfaglrkli";
				var v55: map<bool, seq<int>> := map[f14 := [v43, |v54|]];
				var v56: seq<int> := [v43, v43];
				var v57: map<bool, seq<int>> := map[f14 := (if (f14 in v55) then v55[f14] else v56)[safeIndex(v43, |(if (f14 in v55) then v55[f14] else v56)|) := v43]];
				v55 := (map[f14 := [v43]] + fm8(f14, f14, |fm6(globalState)|, f14, globalState)) + (if (f14) then v57 else v55);
				var v58: map<bool, D1> := map[if (f14 in v49) then v49[f14] else f14 := v53];
				v58 := v58[v47[safeIndex(v43, |v47|)] := DC4(v43, v43)];
				globalState.f9 := f14;
				globalState.f9 := false;
			} else {
				var v59: array<D1> := new D1[15];
				var v60: array<array<D1>> := new array<D1>[28] [v59, v59, v59, v59, v59, v59, v59, v59, v59, v59, v59, if (f14) then v59 else v59, v59, v59, v59, v59, v59, v59, v59, v59, v59, v59, v59, v59, v59, v59, v59, v59];
				v60[safeIndex(33, v60.Length)] := v59;
				var v61 := DC4(v43, v43);
				v60[safeIndex(33, v60.Length)], v61, v43 := if (false && f14) then v59 else v59, v61, v43;
				var v62: seq<array<bool>> := [v0, v0, v0];
				v62 := (v62[safeIndex(-0x3de, |v62|) := v0])[safeIndex(v43, |v62[safeIndex(-0x3de, |v62|) := v0]|) := v0];
				var v63: seq<seq<bool>> := [v47];
				var v64 := new C6(f14, multiset(v63[safeIndex(v43, |v63|)]));
				var v65 := "ehllenung";
				v65 := v65;
				var v66: C7 := new C7();
				v66 := v66;
			}
			
			var v67: array<C5> := new C5[22];
			v67 := v67;
		}
		var v69: array<map<int, seq<int>>> := new map<int, seq<int>>[10](i4 => if (f14) then map v68 : int | (-0x229 <= v68) && (v68 < 0x3c7) :: (v68 * v43) := (seq(abs(0x27c), i5  => (v43))) else map[|multiset{v44, 'v'}| := DC32(seq(abs(281), i6  => (v43))).cf56]);
		var v70: C1 := new C1();
		var v71: map<bool, C1> := map[f14 := v70];
		var v72: seq<int> := [|v71|];
		var v73: map<int, seq<int>> := map[v43 := v72];
		v69[safeIndex(793, v69.Length)] := v73 + v73;
		var v74: set<int> := {v43};
		var v75 := DC30(v44, f14, v43, v74, v43);
		globalState.f9, v44, v69[safeIndex(793, v69.Length)] := true, (v75.(cf52 := v43, cf51 := f14, cf54 := v43)).cf50, fm52(-v43 * v43, globalState);
		var v76: array<int> := new int[2];
		var v77: seq<array<int>> := [v76, v76];
		v76 := if (if (f14) then true else f14) then v76 else v77[safeIndex(v43, |v77|)];
		r0 := |v74|;
		var v78: map<int, bool> := map[0x11f := f14];
		var v79: seq<map<int, bool>> := [v78];
		r1 := DC30(v44, f14, v43, fm22(v43, v79, f14, f14, globalState), v43).cf51;
		var v80: map<map<int, bool>, int> := map[v78 := v43];
		r2 := v80;
		r3 := !f14;
	}
	method m1(p0: set<int>, p1: int, globalState: GlobalState) returns (r0: int, r1: int, r2: bool) {
		var v0: map<int, int> := map[p1 := p1];
		v0 := v0;
		var v1: array<bool> := new bool[10] [f14, f14, f14, f14, f14, f14, false, f14, f14, f14];
		v1[safeIndex(984, v1.Length)] := f14;
		v1[safeIndex(984, v1.Length)] := !f14;
		var v2 := new C1();
		globalState.f9 := !(p1 >= p1);
		var v3: array<int> := new int[11];
		forall i0 | 0 <= i0 < v3.Length {
			v3[i0] := safeDivisionInt(i0, p1);
		}
		var v4: multiset<bool> := multiset{|v0| == p1};
		v4 := f15;
		var v5 := "cxh";
		r0 := |(v5 + "u")| * p1;
		r1 := p1;
		r2 := v1[safeIndex(984, v1.Length)];
	}
	method m2(p0: int, p1: string, p2: bool, p3: bool, globalState: GlobalState) returns (r0: char) {
		var v0 := DC27();
		v0 := v0;
		globalState.f3 := p0;
		var v1: array<int> := new int[20];
		forall i0 | 0 <= i0 < v1.Length {
			v1[i0] := safeDivisionInt(i0, p0);
		}
		var v2: seq<bool> := [f14, p3, p3];
		var v3: C3 := new C3(p0, p1, p3, f15);
		var v4: multiset<C3> := multiset{v3};
		var v5 := 'x';
		var v6: map<int, bool> := map[p0 := f14];
		var v7: map<char, int> := map[v5 := |v6[p0 := f14]|];
		var v8: set<seq<int>> := {seq(abs(14), i1  => (v3.f20))};
		var v9: map<bool, int> := map[fm23(p0, v2[safeIndex(|v4|, |v2|)], v7, v8, globalState) := v3.f20];
		var v10: multiset<int> := multiset{p0};
		globalState.f3 := if (!(f14 && f14) in v9) then v9[!(f14 && f14)] else -fm0(f14, |v3.f21|, p0, v10, globalState);
		var v11: array<bool> := new bool[14];
		globalState.f2 := v11;
		for i2 := p0 to v3.f20 {
			var v12: C5 := new C5(p0, p3, f15 * f15);
			var v13: map<int, string> := map[751 := p1];
			var v14 := DC26(p3, v13);
			var v15: map<bool, bool> := map[p2 := v12.fm2(p2, p3, globalState)];
			var v16: map<bool, bool> := map[p2 := if (p2 in v15) then v15[p2] else p3];
			v12, globalState.f3 := v12, -(if (v14.cf46) then v12.f18 else safeModuloInt(v12.f18, -|v16|));
			var v17 := DC12(v11);
			v17 := v17;
			var v18: T0 := new C5(p0, p3, f15);
			var v19: map<T0, bool> := map[v18 := p2];
			v19 := v19;
			var v20: array<array<int>> := new array<int>[9];
			v20[safeIndex(385, v20.Length)] := v1;
			v20[safeIndex(385, v20.Length)] := v1;
		}
		r0 := v5;
	}
}

function fm0(p0: bool, p1: int, p2: int, p3: multiset<int>, globalState: GlobalState): int {
	(|"fbkunt"| - |[|map[true := {-|multiset{|"ich"|, 910}|, 0xd0}]|]|) * |"cgadqenod"|
}
function fm1(p0: bool, p1: int, p2: bool, p3: set<int>, globalState: GlobalState): map<bool, int> {
	map[{true, true, !!false} > {!!false} := 0x286]
}
function fm5(p0: bool, globalState: GlobalState): seq<bool> {
	[false, false] + [!false]
}
function fm6(globalState: GlobalState): set<int> {
	((set v0 : int | (-0x362 <= v0) && (v0 < 9) :: (v0 - |multiset{DC7([false])}|)) + {|map["k" := 696]|}) - ({0x10b} * {305})
}
function fm7(globalState: GlobalState): map<bool, bool> {
	(map[true := true] + map[!true := true]) + (map[true := false] + map[false := true])
}
function fm8(p0: bool, p1: bool, p2: int, p3: bool, globalState: GlobalState): map<bool, seq<int>> {
	map[true := [DC30('q', false, -0xf8, {0x24c, 628, |"slllsh"|, |map[false := true]|, -0x2e4}, |"pbe"|).cf54]] + map[true := [-0x357]]
}
function fm13(p0: multiset<bool>, p1: int, globalState: GlobalState): multiset<seq<int>> {
	(multiset([[0x26c], [|"vsrbg"|, 0x269, |multiset{true}|]]) * multiset{seq(abs(481), i0  => (-|[false, !false]|)), [-578]}) * multiset([[|map[|{false}| := true]|], [-0x52, 0x173]] + [[-0x268]])
}
function fm14(globalState: GlobalState): map<int, int> {
	map[-0xbc := 0xca] + (map[0x14a := 681] + map[|(seq(abs(0x26b), i0  => (524)))| := |[false]|])
}
function fm19(p0: int, p1: bool, globalState: GlobalState): set<char> {
	{'a', 'd', 'a'} - {'g', 's', 'm', 'd', 'h'}
}
function fm20(p0: int, p1: int, p2: int, globalState: GlobalState): D0 {
	match if (false) then DC23(false, 842, false) else DC23(true, |[|[true]|, -0xa4]|, !true) {
		case DC22(cf37, cf38, cf39, cf40, cf41) => DC0(true)
		case DC23(cf42, cf43, cf44) => DC0(cf42)
		case DC21(cf36) => DC0(true)
	}
}
function fm21(p0: D1, p1: bool, p2: seq<D2>, globalState: GlobalState): seq<int> {
	match DC14(true, DC23(true, 0x30d, true).cf44) {
		case DC13(cf22, cf23, cf24, cf25) => seq(abs(0x153), i0  => (cf23))
		case DC14(cf26, cf27) => seq(abs(-0x327), i1  => (|[|[|{|multiset{|[cf26]|, |(set v0 : int | (752 <= v0) && (v0 < 428) :: (v0 - 708))|}|}|, 0x0]|]|))
		case DC12(cf21) => [0x6a, |{-0xa3, |map[0x3a9 := 0x29b]|, 0x31b, |{true, !true, !true}|}|]
		case DC15(cf28) => [|"fdqjqr"|, |multiset{141}|, |[!true, false]|, |"pvuwwxen"|, --|multiset{false, !true, true}|]
	}
}
function fm22(p0: int, p1: seq<map<int, bool>>, p2: bool, p3: bool, globalState: GlobalState): set<int> {
	if (!(true <== true)) then {-|{|map[true := |[|"nuqceuqr"|, -0x167]|]|, 0x366}|} + (set v0 : int | (0x300 <= v0) && (v0 < -0x2c0) :: (safeDivisionInt(v0, 0xec))) else {569, |multiset([false, !false])|, |(set v1 : int | v1 in multiset([|"l"|]) :: (safeModuloInt(v1, |[-|[|[|[688, 0x163, |[!true, true]|]|, |(seq(abs(0x19c), i0  => (384)))|]|, |"dp"|, -0x13e, -|map[true := true]|, 0x32d]|]|)))|, -0x388}
}
function fm23(p0: int, p1: bool, p2: map<char, int>, p3: set<seq<int>>, globalState: GlobalState): bool {
	true
}
function fm24(p0: int, p1: bool, p2: bool, p3: int, globalState: GlobalState): map<char, int> {
	map['s' := --0x30a] + map['y' := -|map[!!true := DC56(true, true, |{754, |"kk"|}|).cf92]|]
}
function fm25(p0: int, p1: int, p2: bool, p3: map<bool, int>, globalState: GlobalState): char {
	'g'
}
function fm26(p0: int, p1: bool, p2: int, globalState: GlobalState): seq<seq<char>> {
	match DC38(DC30('y', true, -|map[true := -0xaa]|, {-|multiset{--48, 0x106}|, |{true}|}, -0x60).cf52) {
		case DC38(cf64) => [['b', 'x']]
		case DC39(cf65) => seq(abs(881), i0  => (seq(abs(277), i1  => ('r'))))
		case DC37(cf63) => [seq(abs(114), i2  => ('c'))] + (seq(abs(87), i3  => ([DC30('j', true, |multiset([true])|, set v0 : int | (0x2ee <= v0) && (v0 < 0xd0) :: (safeModuloInt(v0, |(set v1 : int | v1 in map[|[|map[false := false]|]| := true] :: (safeModuloInt(v1, 0x371)))|)), -0x152).cf50])))
		case DC40(cf66) => [seq(abs(0x36a), i4  => ('j'))]
	}
}
function fm27(p0: int, p1: bool, globalState: GlobalState): string {
	"vkgblbh" + (seq(abs(122), i0  => ('t')))
}
function fm28(p0: bool, globalState: GlobalState): D1 {
	DC4(|map['a' := false]| * |"hyccnriii"|, -0x23d - 0x26a)
}
function fm29(globalState: GlobalState): seq<seq<bool>> {
	[[false] + [true, !true]]
}
function fm30(p0: bool, globalState: GlobalState): multiset<int> {
	multiset{-690} - multiset{-0x27f, |multiset{[false], [false]}|, -0x2e0, |(seq(abs(-596), i0  => (|[|"cmgn"|]|)))|, 0x1d1}
}
function fm31(p0: bool, p1: bool, globalState: GlobalState): map<int, char> {
	map[-|("ymdokrv" + (seq(abs(957), i0  => ('q'))))| := 'c']
}
function fm32(p0: string, p1: multiset<int>, globalState: GlobalState): set<char> {
	set v0 : char | v0 in (map['b' := multiset([false, true])] + map['p' := multiset{false}]) :: (v0)
}
function fm33(p0: bool, globalState: GlobalState): multiset<bool> {
	multiset(([!!false, !!true] + [true, true]) + ([!true] + [true, false, true, false]))
}
function fm34(globalState: GlobalState): set<string> {
	set v1 : string | v1 in ((map v0 : string | v0 in (seq(abs(0x393), i0  => (DC70(true, !false, "rgiu", |multiset{0x14b, |"ia"|}|, |map[-0x2f := |(seq(abs(0x35c), i1  => ('l')))|]|).cf114))) :: (v0) := (|[[true, !true]]|)) + map["elx" := 919]) :: (v1)
}
function fm35(p0: int, p1: int, globalState: GlobalState): map<int, D5> {
	map[|map[DC1(false).cf1 := false]| := DC18(false)] + ((map v0 : int | (0x20f <= v0) && (v0 < 0x13b) :: (v0 + 0x38d) := (DC18(!false))) + map[|"niflmco"| := DC18(true)])
}
function fm36(p0: multiset<int>, p1: int, p2: bool, p3: bool, globalState: GlobalState): seq<map<bool, int>> {
	([map[true := 0x1fe]] + [map[!true := |map[|map[true := !false]| := 0x3ac]|]]) + [map[true := |"xcail"|], map[false := |(map v0 : int | (0x3ba <= v0) && (v0 < 0x44) :: (v0 + |{false}|) := (|[|map[false := 'f']|]|))|]]
}
function fm37(p0: set<int>, p1: map<int, set<bool>>, p2: bool, p3: bool, globalState: GlobalState): D0 {
	DC1(!(|map[|[false]| := true]| < |"tvidhsf"|))
}
function fm38(globalState: GlobalState): seq<bool> {
	match DC17([[true, false, false]], false, multiset{372, |map[|(seq(abs(142), i0  => ('e')))| := |{'w', 'n', 'j', 'r'}|]|, |map[false := multiset{839}]|, |map[0x147 := false]|}) {
		case DC17(cf30, cf31, cf32) => [cf31] + [cf31]
		case DC18(cf33) => [cf33, true]
		case DC16(cf29) => [true, true]
		case DC19(cf34) => [false, true, !false, true, true] + [true]
	}
}
function fm39(p0: int, p1: int, p2: int, p3: int, globalState: GlobalState): D9 {
	DC30('w', if (true) then false else true, |map[|"sfdadptf"| := true]| - 0x34c, set v0 : int | (626 <= v0) && (v0 < 0x15d) :: (safeModuloInt(v0, 168)), |map[|{DC23(false, 232, !false).cf43, |"ixq"|, |map[-0xe9 := true]|}| := true]|)
}
function fm40(globalState: GlobalState): seq<D2> {
	[DC7([false, true])]
}
function fm41(p0: int, p1: int, globalState: GlobalState): set<bool> {
	({!false, !true, !true} * {!true}) + ({!!!false} + {false})
}
function fm42(p0: bool, p1: bool, p2: int, p3: bool, globalState: GlobalState): D8 {
	match DC33({multiset([false])}) {
		case DC34(cf58, cf59, cf60, cf61) => DC26(true, map v0 : int | v0 in DC32([-|"di"|]).cf56 :: (v0 - -0x28a) := ("pc"))
		case DC35() => DC26(false, map[0x42 := seq(abs(108), i0  => ('y'))])
		case DC33(cf57) => DC26(!true, map v1 : int | (0x210 <= v1) && (v1 < 0x317) :: (v1 - 586) := ("hsqevyqb"))
		case DC36(cf62) => if (!true) then DC26(true, map[-0x77 := seq(abs(0xa9), i1  => ('e'))]) else DC26(false, map[-0x3b9 := "m"])
	}
}
function fm43(p0: int, p1: int, p2: int, globalState: GlobalState): seq<multiset<bool>> {
	([multiset{DC0(true).cf0, !true, true, true}] + [multiset{false}, multiset([true, false]), multiset{false}, multiset([true]), multiset{!false}]) + ([multiset{true}] + (seq(abs(-251), i0  => (multiset([false, false])))))
}
function fm44(p0: bool, p1: int, p2: int, globalState: GlobalState): D2 {
	match DC41({[0x2b1]}) {
		case DC42(cf68, cf69) => DC8(-|[seq(abs(0x261), i0  => (cf68))]|, -0x93, -0xd6)
		case DC41(cf67) => DC8(|[|map[true := true]|, 0x75]|, 293, -852)
		case DC43(cf70) => DC8(|{188}|, |map[true := true]|, |[false, false, true]|)
	}
}
function fm45(p0: int, p1: bool, p2: int, globalState: GlobalState): D3 {
	match DC30('q', false, |{|map[--|map[false := true]| := 0]|, |multiset{false, true, false}|}|, {0x14f, 0xda, -0x3d3}, |map[-175 := |[!false]|]|) {
		case DC30(cf50, cf51, cf52, cf53, cf54) => DC10(map[cf51 := cf51])
		case DC29(cf49) => DC10(map[true := false])
		case DC31(cf55) => DC10(map[true := false])
	}
}
function fm46(p0: int, p1: bool, globalState: GlobalState): map<char, bool> {
	map['p' := !!true] + (map v0 : char | v0 in ['d', 'g', 'b', 'e', 'a'] :: (v0) := (true))
}
function fm47(globalState: GlobalState): set<multiset<bool>> {
	({multiset([false, false]), multiset{true}, multiset{false}} - {multiset([false, true, false, false]), multiset{!false}}) + (set v0 : multiset<bool> | v0 in {multiset([false, false, false, true]), multiset{true}, multiset{false, !false}, multiset([!true, true, true, true])} :: (v0))
}
function fm48(p0: int, p1: int, p2: bool, globalState: GlobalState): map<map<int, int>, int> {
	map[map[-|{!true}| := |map[[true] := true]|] := 0x367 + 758]
}
function fm49(p0: int, globalState: GlobalState): set<seq<D9>> {
	({[DC31(DC30('h', false, |(map v0 : seq<int> | v0 in map[[386] := map[!true := true]] :: (v0) := (false))|, set v1 : int | v1 in multiset{954} :: (safeDivisionInt(v1, |[|map[|map[-0xfd := |map[false := 0x370]|]| := DC4(|{-0xf}|, |multiset{0x3ce, |[!false, false]|, |{true}|, 0x6a, |"xpse"|}|)]|, 0xa2, |[false, false, !false]|, 0x3a1]|)), 0x115))], [DC31(DC31(DC29(map v2 : int | (0x8c <= v2) && (v2 < 0x284) :: (v2 + 0x125) := (|"phlpopw"|))))], seq(abs(-0x1bf), i0  => (DC31(DC31(DC30('i', true, |map[true := |map[false := true]|]|, {|[true, true, true, true]|}, |(seq(abs(-0x339), i1  => ('x')))|)))))} * {[DC31(DC31(DC31(DC30('w', false, |[|{false}|]|, {0x10f, 0x1bc}, 0xa8))).cf55), DC31(DC29(map[0x2dd := -0x1eb]))]}) * ((set v4 : seq<D9> | v4 in [[DC31(DC31(DC29(map v3 : int | (0x18a <= v3) && (v3 < -0xbb) :: (safeDivisionInt(v3, |(seq(abs(0x397), i2  => ('f')))|)) := (|multiset{-0x22f, 990}|)))), DC31(DC29(map[|multiset{false}| := |(seq(abs(395), i3  => ('g')))|]))]] :: (v4)) - {[DC31(DC31(DC31(DC31(DC31(DC29(map[-463 := 0x259])))))), DC31(DC31(DC31(DC29(map v5 : int | v5 in map[|multiset{false, false}| := false] :: (safeDivisionInt(v5, |map['j' := |map[0x7e := |{|"fjfiraa"|}|]|]|)) := (-|"to"|)))))]})
}
function fm50(p0: int, p1: bool, p2: int, p3: int, globalState: GlobalState): set<seq<int>> {
	{[|(map v0 : int | (0x266 <= v0) && (v0 < 0x1e4) :: (v0 + |(seq(abs(-399), i0  => ('b')))|) := (true))|]} * {[|[true, false]|]}
}
function fm51(p0: bool, p1: bool, p2: bool, p3: int, globalState: GlobalState): D13 {
	DC43(DC43(DC42('k', true)))
}
function fm52(p0: int, globalState: GlobalState): map<int, seq<int>> {
	(map[|{false}| := [|{|multiset{true}|}|, |(seq(abs(-0x135), i0  => (|"jlkqqbehd"|)))|, 589]] + map[0x242 := [---0x35b]]) + map[-0x21d := [83]]
}
function fm53(p0: bool, p1: bool, globalState: GlobalState): seq<set<int>> {
	[{-116, |[|[[|{!false}|], [50]]|, 0x3b1]|}, {|(map v0 : char | v0 in DC67("hac").cf108 :: (v0) := (|multiset(seq(abs(0x1aa), i0  => (-|(seq(abs(911), i1  => ('x')))|)))|))|}, {0x6a, -|[|multiset{|[false]|, 427}|]|}] + ([{|multiset{'q'}|, 0x146}] + [{0x377}])
}
function fm54(p0: int, p1: string, globalState: GlobalState): set<set<bool>> {
	if (!(|"pgi"| >= -|(seq(abs(371), i0  => ('l')))|)) then {DC24({DC5(true, false).cf7}).cf45} else (set v0 : set<bool> | v0 in {{true, false}} :: (v0)) * {{true, false}}
}
function fm55(p0: bool, globalState: GlobalState): seq<seq<int>> {
	((seq(abs(352), i0  => ([|"wvjfiqb"|]))) + [[-532, -0x1d6]]) + ([seq(abs(0x148), i1  => (|multiset{false}|))] + [[|(seq(abs(0x267), i2  => ('u')))|, |[false]|], seq(abs(-0x3e2), i3  => (|"xoxncefj"|)), [0x329]])
}
method m11(p0: D13, p1: int, p2: bool, p3: bool, globalState: GlobalState) returns (r0: int, r1: map<int, array<T0>>) {
	var v0 := "plwnbed";
	var v1: array<bool> := new bool[1];
	var v2 := DC13(v0, p1, false, v1);
	match ((if (false) then v2 else v2).(cf22 := "lkgqgh")) {
		case DC13(cf22, cf23, cf24, cf25) =>
			cf24 := if (true) then p3 else p2;
			globalState.f9 := p2;
			var v3: multiset<bool> := multiset{p3};
			var v4: T0 := new C6(cf24, v3 + v3[!p3 := abs(cf23)]);
			v4 := v4;
			var v5: array<char> := new char[22];
			var v6: map<int, array<char>> := map[-0x176 := v5];
			var v7 := DC56(p3, v4.f14, cf23);
			var v8: seq<bool> := [p2];
			var v9: C0 := new C0(v8[safeIndex(cf23, |v8|)], p2);
			var v10: seq<C0> := [v9];
			var v11: map<bool, int> := map[v7.cf93 := |v10|];
			var v12: seq<map<bool, int>> := [v11, map[p3 := cf23], v11[v9.f24 := p1]];
			var v13: C8 := new C8(v6, v12, p2, v3);
			var v14: array<C8> := new C8[5] [v13, v13, v13, v13, v13];
			v14 := if (cf24) then v14 else v14;
		case DC14(cf26, cf27) =>
			var v15: seq<bool> := [p2];
			var v16: C6 := new C6(true, multiset([cf26]));
			var v17: map<bool, C6> := map[p2 := v16];
			globalState.f3, globalState.f9 := safeDivisionInt(|(if (p3) then v15 else v15)|, |map[0x2b2 := cf26]| - |v17[v15[safeIndex(p1, |v15|)] := v16]|), v16.fm2(v15 <= v15, p2, globalState);
			var v18: multiset<bool> := multiset{cf26, cf27};
			var v19: C4 := new C4(cf27, cf26, v18);
			var v20: map<C4, bool> := map[v19 := p2];
			var v21: C5 := new C5(p1 + |v20|, p2, multiset{cf26, true});
			v21 := v21;
			var v22: array<int> := new int[24];
			v22[safeIndex(372, v22.Length)] := v21.f18 * -p1;
			v1[safeIndex(564, v1.Length)] := p1 != p1;
			var v23: array<C1> := new C1[9];
			var v24: set<array<C1>> := {v23, v23, v23};
			var v25: map<bool, int> := map[v16.fm2(p2, v19.f19, globalState) := safeDivisionInt(-p1, |v24|)];
			v0, globalState.f9, v22[safeIndex(372, v22.Length)], v1[safeIndex(564, v1.Length)] := v0 + (seq(abs(0x396), i0  => ('o'))), cf26, |v25|, p2;
			v1[safeIndex(564, v1.Length)] := v15[safeIndex(v22[safeIndex(372, v22.Length)], |v15|)];
		case DC12(cf21) =>
			v1[safeIndex(719, v1.Length)] := p1 < p1;
			var v26: set<bool> := {p3, p2, false};
			var v27: set<set<bool>> := {v26};
			globalState.f3, v1[safeIndex(719, v1.Length)] := -safeModuloInt(p1, p1), fm54(0x352, v0, globalState) >= (v27 + v27);
			if (v1[safeIndex(719, v1.Length)]) {
				globalState.f3 := p1;
				var v28: seq<int> := [p1];
				var v29: multiset<bool> := multiset{p3};
				var v30 := new C6(|v28| > p1, v29);
				globalState.f3 := p1 + p1;
				var v31: T0 := new C3(p1, v0, p3, v29);
				var v32: seq<T0> := [v31];
				var v33: array<int> := new int[17];
				var v34: set<array<int>> := {v33, v33, v33};
				var v35: set<int> := {|v34|};
				globalState.f9 := |([v32[safeIndex(p1, |v32|)], v31, v31, v31, v31] + v32)| !in v35;
				v1[safeIndex(719, v1.Length)] := !true;
			} else {
				v1[safeIndex(719, v1.Length)] := p1 >= p1;
				var v36: map<int, bool> := map[p1 := v1[safeIndex(719, v1.Length)]];
				var v37 := DC5(p3, true);
				var v38: map<bool, D1> := map[if (p1 in v36) then v36[p1] else true := v37];
				v38 := v38;
				var v39: set<int> := {p1, p1};
				var v40: array<C3> := new C3[28];
				var v41: seq<set<int>> := [v39, fm6(globalState)];
				var v42: seq<bool> := [p2, p3];
				var v43: map<bool, bool> := map[v1[safeIndex(719, v1.Length)] := v1[safeIndex(719, v1.Length)]];
				var v44 := 'c';
				var v45: map<char, int> := map[v44 := p1];
				var v46: seq<int> := [p1, p1, p1, p1, p1];
				var v47: set<seq<int>> := {v46, v46, v46, v46};
				v39, v1[safeIndex(719, v1.Length)], v40 := v41[safeIndex(|v42|, |v41|)] - v39, if (!v1[safeIndex(719, v1.Length)] in v43) then v43[!v1[safeIndex(719, v1.Length)]] else fm23(|v36|, if (p1 in v36) then v36[p1] else !p3, v45, v47, globalState), v40;
				var v48: array<array<bool>> := new array<bool>[8] [cf21, cf21, cf21, cf21, cf21, cf21, cf21, cf21];
				v48[safeIndex(246, v48.Length)] := cf21;
				v48[safeIndex(246, v48.Length)] := new bool[29];
				var v49: multiset<bool> := multiset{p3, p2};
				var v50 := new C5(p1, p3, v49 - v49[p3 := abs(|v49|)]);
			}
			
			var v51: map<bool, bool> := map[p2 := p3];
			var v52: multiset<bool> := multiset{v1[safeIndex(719, v1.Length)], v1[safeIndex(719, v1.Length)], p2};
			var v53: multiset<int> := multiset{-0x231, p1};
			var v54: seq<int> := [p1];
			var v55: seq<seq<int>> := [v54];
			var v56: map<bool, int> := map[false := 622];
			var v57: array<int> := new int[23] [safeModuloInt(p1, p1), fm0(p2, -p1, p1, (multiset{|v0|, -0x26e, p1, p1, -p1})[p1 := abs(-|v51|)], globalState), p1, p1 + -|v52|, p1, 978, fm0(false, p1, p1, v53, globalState), |v55|, p1, if (p2) then p1 else p1, if (v1[safeIndex(719, v1.Length)]) then v54[safeIndex(p1, |v54|)] else p1, safeDivisionInt(p1, p1), p1, 0x3a1, -|v54|, p1, -|v54|, p1, if (v1[safeIndex(719, v1.Length)] in v56) then v56[v1[safeIndex(719, v1.Length)]] else 0xab, p1, p1, p1, p1];
			v57 := v57;
			if (-safeDivisionInt(p1, p1) <= p1) {
				globalState.f9 := !(if (true) then v1[safeIndex(719, v1.Length)] else p2);
				var v58: map<array<bool>, bool> := map[cf21 := v1[safeIndex(719, v1.Length)]];
				var v59 := 'p';
				var v60: map<char, int> := map[v59 := |[p1, |{p1}|]|];
				var v61: set<seq<int>> := {v54};
				v58 := map[cf21 := fm23(p1, false, fm24(-773, p3, true, p1, globalState), set v62 : seq<int> | v62 in fm55(fm23(p1, p2, v60, v61, globalState), globalState) :: (v62), globalState)];
				var v64: map<seq<int>, bool> := map[v54 := p2];
				var v66: array<C6> := new C6[15];
				var v67: C3 := new C3(fm0(v1[safeIndex(719, v1.Length)], p1, p1, v53, globalState), v0[safeIndex(|map[fm0(fm23(p1, false, map v63 : char | v63 in v0 :: (v63) := (p1), set v65 : seq<int> | v65 in v64 :: (v65), globalState), p1, p1, v53, globalState) := v66]|, |v0|) := v59] + v0[safeIndex(p1, |v0|) := v59], v1[safeIndex(719, v1.Length)], v52);
				v67 := new C3(v67.f20, v67.f21, p3, v52);
				var v68: set<int> := {v67.f20, v67.f20};
				var v69, v70, v71 := v67.m1(v68, v67.f20, globalState);
				var v72 := DC18(false);
				var v73: map<bool, D5> := map[p3 := v72];
				v73 := v73[p3 := DC18(v1[safeIndex(719, v1.Length)])];
			} else {
				var v75 := 'f';
				var v77: seq<bool> := [p2];
				var v78: map<int, int> := map[|v0| := |(map v76 : seq<bool> | v76 in multiset{v77} :: (v76) := ([p1]))|];
				var v79: map<char, map<int, int>> := map[v75 := v78[|v77| := 0xa5]];
				var v80: set<int> := {p1, p1};
				var v81 := DC30(fm25(p1, p1, p3, fm1(p2, p1, true, v80, globalState), globalState), v1[safeIndex(719, v1.Length)], p1, v80, 0x3c1);
				var v82: map<char, int> := map[v81.cf50 := -|(seq(abs(0x210), i1  => (v75)))|];
				var v84: map<seq<int>, int> := map[[0x65] := p1];
				v1[safeIndex(719, v1.Length)] := fm23(p1, false, (map v74 : char | v74 in v79 :: (v74) := (0x30d)) + v82, set v85 : seq<int> | v85 in (map v83 : seq<int> | v83 in v84 :: (v83) := (0x21c)) :: (v85), globalState);
				var v86: array<string> := new string[1] [v0];
				v86[safeIndex(664, v86.Length)] := v0 + (seq(abs(0x365), i2  => (v75)));
				v86[safeIndex(664, v86.Length)] := v0;
				globalState.f3, cf21, r0 := 0x29b, cf21, p1 - |(v86[safeIndex(664, v86.Length)] + v86[safeIndex(664, v86.Length)][safeIndex(p1, |v86[safeIndex(664, v86.Length)]|) := v86[safeIndex(664, v86.Length)][safeIndex(|v56|, |v86[safeIndex(664, v86.Length)]|)]])|;
				v80 := (v80 * v80) + ({p1} - v80);
				var v87 := DC3(p1);
				v57[safeIndex(906, v57.Length)] := v87.cf3;
				v57[safeIndex(906, v57.Length)] := safeDivisionInt(p1, 0x1da) - p1;
			}
			
		case DC15(cf28) =>
			var v88 := 'i';
			var v89: map<char, int> := map[v88 := p1];
			var v90: multiset<bool> := multiset{p3, p3};
			var v91: seq<int> := [|v90|, 555];
			var v92: set<seq<int>> := {v91, [p1]};
			var v93: seq<bool> := [fm23(|{--0x320, p1}|, p2, v89, v92, globalState)];
			var v94: array<char> := new char[1];
			var v95: map<int, array<char>> := map[|multiset(v93)| := v94];
			var v96: map<bool, int> := map[p3 := -|v0|];
			var v97: seq<map<bool, int>> := [v96];
			var v98: T0 := new C8(v95, v97, v93[safeIndex(|v93|, |v93|)], v90[p3 := abs(-p1)]);
			var v99 := DC9(v98.f14, 0x13, v0[safeIndex(p1, |v0|)]);
			v98 := new C3(v99.cf14, fm27(p1, v98.f14, globalState), v98.f14, v98.f15 * multiset(v93));
			v98.f14 := false;
			var v100: array<map<bool, bool>> := new map<bool, bool>[20];
			var v101: map<bool, bool> := map[v98.f14 := p3];
			v100[safeIndex(988, v100.Length)] := v101;
			var v102: multiset<int> := multiset{|v0|};
			var v103: array<int> := new int[3] [-0x2e, fm0(v98.f14, p1, p1, v102, globalState), -p1];
			v103[safeIndex(718, v103.Length)] := p1;
			v1[safeIndex(672, v1.Length)] := v98.f14;
			v94[safeIndex(614, v94.Length)] := v88;
			var v104: seq<map<bool, bool>> := [map[true := v98.f14], v101, map[true := p2]];
			var v105: C8 := new C8(map[p1 := v94], v97, v98.f14, (multiset{p2})[v98.f14 := abs(p1)]);
			var v106: multiset<C8> := multiset{v105};
			var v107: map<int, bool> := map[p1 := v98.f14];
			v100[safeIndex(988, v100.Length)], r0, v103[safeIndex(718, v103.Length)], v1[safeIndex(672, v1.Length)], v94[safeIndex(614, v94.Length)] := v104[safeIndex(p1, |v104|)], -(-(|multiset{-|v106|}| * p1) + p1), p1, if (|fm27(0x1d3, v98.f14, globalState)| in v107) then v107[|fm27(0x1d3, v98.f14, globalState)|] else p2, v88;
			if (!v1[safeIndex(672, v1.Length)]) {
				var v108 := new C2(v107);
				v103[safeIndex(718, v103.Length)] := v91[safeIndex(v103[safeIndex(718, v103.Length)], |v91|)];
				var v109: array<array<array<int>>> := new array<array<int>>[12];
				var v110: map<int, array<int>> := map[v103[safeIndex(718, v103.Length)] := v103];
				var v111: map<bool, array<int>> := map[p2 := v103];
				var v112: array<array<int>> := new array<int>[25] [v103, v103, v103, v103, v103, v103, v103, v103, v103, v103, v103, v103, v103, v103, v103, v103, if (p1 in v110) then v110[p1] else v103, v103, v103, if (v98.f14 in v111) then v111[v98.f14] else v103, v103, v103, v103, v103, v103];
				v109[safeIndex(822, v109.Length)] := v112;
				v109[safeIndex(822, v109.Length)], globalState.f3, globalState.f3 := v112, v103[safeIndex(718, v103.Length)], safeDivisionInt(p1, v103[safeIndex(718, v103.Length)]);
				var v113 := new C1();
				var v114 := DC1(p3);
				v114 := v114;
			} else {
				var v115: array<bool> := new bool[12];
				var v116: map<int, array<bool>> := map[p1 := v115];
				globalState.f2 := if ((if (v98.f14) then v103[safeIndex(718, v103.Length)] else -369) in v116) then v116[if (v98.f14) then v103[safeIndex(718, v103.Length)] else -369] else v115;
				var v117: map<map<int, bool>, bool> := map[map[v103[safeIndex(718, v103.Length)] := v98.f14] := v98.f14];
				var v118: map<int, string> := map[|v117| := v0];
				var v119 := DC26(fm23(p1, v98.f14, v89, v92, globalState), v118);
				var v120 := DC3(0x328);
				var v121 := DC6(v120);
				var v122 := DC6(v120);
				var v123: map<D8, D1> := map[v119 := v121];
				var v124 := DC6(if (v119 in v123) then v123[v119] else v120);
				var v125: array<D1> := new D1[1] [v124];
				v125[safeIndex(26, v125.Length)] := v124;
				v125[safeIndex(26, v125.Length)] := v124;
				v103[safeIndex(718, v103.Length)], v98.f14 := safeModuloInt(0x29b, fm0(v98.f14, 0xad, p1, v102, globalState)), p3;
				globalState.f2 := v115;
				var v126: set<int> := {|v96|};
				v98, globalState.f3, v88, v126 := v98, -safeDivisionInt(|v107| + v103[safeIndex(718, v103.Length)], p1), v94[safeIndex(614, v94.Length)], v126 + (fm6(globalState) + v126);
			}
			
	}
	
	var v127: C7 := new C7();
	v127 := v127;
	r0 := -p1;
	v0 := v2.cf22;
	var v128: map<bool, int> := map[p3 := p1];
	var v129: C4 := new C4(p3, p3, multiset{p2});
	var v130 := DC60(v129);
	var v131: multiset<int> := multiset{p1, |fm7(globalState)|, p1, |v128|, |map[v130.cf99 := p3]|};
	var v132: map<int, bool> := map[-safeDivisionInt(p1, -p1) := v127.fm10(v128, v131, v129.f19, v0, globalState)];
	if (if (safeDivisionInt(p1, p1) in v132) then v132[safeDivisionInt(p1, p1)] else p3) {
		globalState.f3 := p1 + p1;
		var v133: multiset<bool> := multiset{p2, !p2, v129.f19};
		var v134: C5 := new C5(0x2c5, true, v133);
		v134 := if (v129.f19) then v134 else v134;
		var v135 := new C0(v129.f19, p2);
		globalState.f9 := v135.f23 <== ("wbhurb" < v2.cf22);
		var v136 := DC5(v129.f19, v135.f23);
		var v137 := DC6(DC6(v136));
		var v138: map<D1, bool> := map[v137.(cf8 := DC6(v136)) := p3];
		v138 := v138[v137 := p3];
	} else {
		var v139: array<int> := new int[4](i3 => safeDivisionInt(i3, p1));
		v139[safeIndex(520, v139.Length)] := p1;
		var v140: seq<int> := [421 + p1];
		v139[safeIndex(520, v139.Length)] := |v140|;
		var v141: set<bool> := {v127.fm10(v128, v131, v129.f19, v0, globalState), p2};
		globalState.f3 := safeModuloInt(p1, p1) - safeModuloInt(v139[safeIndex(520, v139.Length)], -|v141|);
		globalState.f9 := (if (v129.f19) then v129.f19 else true) || v129.f19;
		var v142 := DC0(p2);
		var v143: multiset<bool> := multiset{v142.cf0, v129.f19};
		var v144: T0 := new C3(v139[safeIndex(520, v139.Length)], v0, v129.f19, v143);
		var v145: C6 := new C6(p3, v144.f15);
		var v146 := DC63(v145);
		var v147: map<T0, C6> := map[v144 := v146.cf101];
		v147 := v147[v144 := v145];
		var v148 := 'o';
		var v149: set<multiset<bool>> := {v144.f15};
		globalState.f9, v148 := v144.f15 in v149, v148;
	}
	
	var v150 := DC18(v127.fm10(v128, v131, p2, v0, globalState));
	globalState.f9 := match v150 {
		case DC17(cf30, cf31, cf32) => p3
		case DC18(cf33) => p3
		case DC16(cf29) => 459 <= -|map[p1 := DC4(p1, p1).cf4]|
		case DC19(cf34) => [p2] != [v129.f19, v129.f19, v129.f19, p2]
	};
	var v151: seq<bool> := [v129.f19];
	r0 := safeModuloInt(|v151|, p1) * p1;
	var v152: array<T0> := new T0[13];
	var v153: map<int, array<T0>> := map[p1 := v152];
	r1 := v153;
}
method Main() {
	var v0 := 0xce;
	var v1: multiset<int> := multiset{v0};
	var v2: array<bool> := new bool[3];
	var v3 := "bwiunjo";
	var v4 := true;
	var v5: set<bool> := {v4, v4};
	var v6: array<set<int>> := new set<int>[12];
	var v7: array<string> := new string[11](i0 => v3);
	var globalState := new GlobalState(946, v1 - v1, v2, 0x298, true, v3, v5, false, 0x2fa, false, -0x1ab, v6, v7, {v4, v4, false});
	var v8: seq<bool> := [v4];
	v8 := v8 + v8;
	var v9: map<bool, set<bool>> := map[v0 != v0 := v5];
	v9 := v9[v4 := v5];
	var v10: seq<int> := [v0, v0];
	for i1 := fm0(v4, v0, v0, (multiset(v10))[v0 := abs(v0)], globalState) to v0 {
		var v11: map<int, array<string>> := map[v0 := v7];
		v7 := if ((v0 * -0x1de) in v11) then v11[v0 * -0x1de] else v7;
		var v12: map<bool, int> := map[v4 := safeModuloInt(i1, i1)];
		var v13: array<int> := new int[7];
		var v14: multiset<array<int>> := multiset{v13};
		var v15: map<int, array<int>> := map[v0 := v13];
		var v16: set<int> := {|v10|, v0, v0};
		v0, globalState.f3, globalState.f3, v12 := safeModuloInt(i1 - i1, safeDivisionInt(v0, v0)), |(v3 + (v3 + v3))|, --|(if (v4) then multiset{v13} else v14)[if (i1 in v15) then v15[i1] else v13 := abs(-v0)]|, v12 + fm1(false, |map[i1 := v4]|, v4, v16, globalState);
		var v17: map<int, bool> := map[v0 := v4];
		if (if (v0 in v17) then v17[v0] else true) {
			v13[safeIndex(0, v13.Length)] := v0;
			v13[safeIndex(0, v13.Length)] := 0x3cb;
			v3, globalState.f9 := v3, v4;
			globalState.f3 := --0x338;
			v4 := v4;
			var v18 := new C1();
		} else {
			var v19 := 'r';
			var v20: map<char, int> := map[v19 := v0];
			v20 := v20[v19 := v0];
			var v21 := DC16(v13);
			var v22: array<D5> := new D5[15] [v21, v21, v21, v21, v21, v21.(cf29 := v13), v21.(cf29 := v13).(cf29 := v13), v21, v21, v21, DC16(v13), v21, DC16(v13), v21, v21];
			v22[safeIndex(543, v22.Length)] := v21;
			var v23: multiset<bool> := multiset{if (v4) then v4 else false, v0 !in [i1], v4, i1 >= v0, v4};
			v2[safeIndex(804, v2.Length)] := v4;
			globalState.f3, v22[safeIndex(543, v22.Length)], v23, v2[safeIndex(804, v2.Length)] := -v0, v21, (if (v4) then v23 else v23) * v23, v4;
			globalState.f3 := if (v4) then safeModuloInt(v0, i1) else i1;
			var v24: map<int, int> := map[-i1 := v0];
			v24 := v24[v0 + i1 := v0];
			v4 := v4;
		}
		
		var v25: T0 := new C5(|v8|, v4, multiset{false, v4});
		var v26: seq<T0> := [v25, v25];
		var v27: array<T0> := new T0[24] [v25, v25, v25, v25, v26[safeIndex(i1, |v26|)], v25, v25, v25, v25, v25, v25, v25, v25, v25, v25, v25, v25, v25, v25, v25, v25, v25, v25, v25];
		var v28: multiset<array<T0>> := multiset{v27, v27, v27};
		v2[safeIndex(802, v2.Length)] := !(v28 >= v28);
		v2[safeIndex(802, v2.Length)] := v4;
	}
	var v29 := DC43(fm51(v4, v4, v4, v0, globalState));
	var v30: map<int, int> := map[v0 := v0];
	var v31 := 'm';
	var v32, v33 := m11(v29, |v30[|v3[safeIndex(v0, |v3|) := v31]| := v0]|, v4, v4, globalState);
	if (if (!!(true || v4)) then v4 else v32 <= v32) {
		var v34: array<array<int>> := new array<int>[18];
		var v35 := DC52(v0, true, v34, v4, v0);
		var v36 := DC9(v4, v0, 'l');
		var v37, v38 := m11(v29, v32 * v35.cf88, v36.cf15 in v3, v4, globalState);
		var v39: array<int> := new int[15](i2 => safeModuloInt(i2, DC4(v0, 0).cf5));
		v39[safeIndex(190, v39.Length)] := v32;
		v39[safeIndex(190, v39.Length)] := -v32;
		if (v4) {
			var v40: map<bool, seq<bool>> := map[v4 := v8];
			var v41: map<char, int> := map[v31 := |v1|];
			var v42: set<seq<int>> := {v10};
			var v43: map<string, multiset<bool>> := map["pshth" := multiset(v8)];
			var v44: multiset<bool> := multiset{v4};
			var v45: map<bool, int> := map[v4 := v39[safeIndex(190, v39.Length)]];
			var v46: map<int, bool> := map[if (v4 in v45) then v45[v4] else |v10[safeIndex(|(seq(abs(-0x271), i4  => ('r')))|, |v10|) := v37]| := v4];
			var v47 := DC5(v4, false);
			var v48 := DC34(v44, v1, v2, v47.(cf7 := v4));
			var v49: array<multiset<bool>> := new multiset<bool>[26] [multiset(if (fm23(|v8|, v4, v41, v42, globalState) in v40) then v40[fm23(|v8|, v4, v41, v42, globalState)] else v8[safeIndex(v39[safeIndex(190, v39.Length)], |v8|) := fm23(|v8|, !v4, fm24(v37, v4, v4, v39[safeIndex(190, v39.Length)], globalState), v42, globalState)]), (if ((seq(abs(-0x3b4), i3  => (v31))) in v43) then v43[seq(abs(-0x3b4), i3  => (v31))] else v44) * multiset{v4}, v44, v44, multiset([!v4, v4]), fm33(!false, globalState) + v44, v44 * v44, multiset(v8), v44, multiset{false}, v44, v44[v4 := abs(v32)], multiset{v4, !v4, if (v37 in v46) then v46[v37] else false}, v44, multiset{true}, (v44[v4 := abs(v37)])[v4 := abs(v0)], v44 + v44, v44, v44, v44, v44, multiset{v4, !v4, true, v4}, v44, (v48.(cf60 := v2, cf58 := v44)).cf58, v44 - v44, v44];
			var v50: array<C3> := new C3[9];
			var v51: map<array<C3>, int> := map[v50 := v37];
			var v52: array<char> := new char[1](i5 => v31);
			var v53 := DC51(v52);
			globalState.f3, v37, v31, v49, globalState.f9 := fm0(v32 >= |v51|, -v32, v32, v1, globalState), v37, v31, v49, v8[safeIndex(|(multiset{v52})[v53.cf83 := abs(0x23)]|, |v8|)];
			var v54: array<D14> := new D14[7];
			var v55 := DC44(v7);
			v54[safeIndex(715, v54.Length)] := v55;
			v54[safeIndex(715, v54.Length)] := v55.(cf71 := v7);
			var v56: map<int, multiset<bool>> := map[v39[safeIndex(190, v39.Length)] := v44[v4 := abs(v32)]];
			var v57 := new C5(v37, v4, (if (|(seq(abs(0x183), i6  => (v31)))| in v56) then v56[|(seq(abs(0x183), i6  => (v31)))|] else v44) + multiset(v8));
			v31 := v31;
			var v58: array<seq<bool>> := new seq<bool>[28](i7 => ([v4, v4])[safeIndex(|v3|, |[v4, v4]|) := v4] + v8);
			v58[safeIndex(201, v58.Length)] := v8;
			v58[safeIndex(201, v58.Length)] := v8;
		} else {
			globalState.f9 := v4;
			v39 := v39;
			globalState.f3 := fm0(!v4, v37, v0, v1 * v1, globalState);
			var v59: array<seq<bool>> := new seq<bool>[19];
			v59[safeIndex(307, v59.Length)] := v8;
			var v60: map<seq<bool>, int> := map[v8 + [false] := if (v4) then -|fm53(v4, v4, globalState)| else v39[safeIndex(190, v39.Length)]];
			var v61: C4 := new C4(v4, v4, (multiset{v4, v4})[v4 := abs(|v1|)]);
			var v62: map<C4, int> := map[v61 := v32];
			v59[safeIndex(307, v59.Length)], v60 := fm38(globalState), (map[v8 := |v62|])[v8 := 0x10b] + map[v8[safeIndex(v0, |v8|) := v61.f19] := v32];
			v39[safeIndex(190, v39.Length)] := -956;
		}
		
		var v64: map<int, set<bool>> := map[v10[safeIndex(v39[safeIndex(190, v39.Length)], |v10|)] := {v4}];
		var v65 := DC2(fm37(set v63 : int | v63 in v10[safeIndex(|v3|, |v10|) := |v1|] :: (safeDivisionInt(v63, 0xf0)), v64, v4, v4, globalState));
		match (v65) {
			case DC1(cf1) =>
				globalState.f3 := v37;
				v32 := v0 + fm0(!false, v37, |v5|, v1, globalState);
				var v66, v67 := m11(v29, v37, 633 < -v32, false, globalState);
				v66 := v32;
			case DC0(cf0) =>
				v4 := cf0;
				v37 := v32;
				var v68: multiset<bool> := multiset{v4, cf0};
				var v69: set<int> := {v32, -v0, if (false in v68) then v68[false] else v37};
				globalState.f9 := v69 < {0x3d5, v0};
				var v70: map<int, seq<bool>> := map[v32 := [v4, cf0]];
				globalState.f9 := ((if (v39[safeIndex(190, v39.Length)] in v70) then v70[v39[safeIndex(190, v39.Length)]] else v8) + v8) < v8;
			case DC2(cf2) =>
				var v71 := DC42(v31, !true);
				var v72, v73 := m11(DC43(v71), v37, v4, v4, globalState);
				v4 := v1 == v1;
				var v74 := DC16(v39);
				v39 := v74.cf29;
				globalState.f9 := v4;
		}
		
		v0 := v0;
	} else {
		globalState.f9 := (if (v4) then !v4 else v4) in (v8 + v8);
		var v75: map<char, int> := map[v31 := v0];
		var v77, v78 := m11(v29, v0, v4, fm23(v0, v4, v75, set v76 : seq<int> | v76 in map[[v0] := v0] :: (v76), globalState), globalState);
		var v79: C4 := new C4(v4, v4, fm33(v4, globalState));
		var v80: multiset<C4> := multiset{v79};
		var v81: multiset<bool> := multiset{v4, v4};
		var v82: map<multiset<C4>, int> := map[v80 := if (v79.f19 in v81) then v81[v79.f19] else v32];
		v82 := v82[v80 := safeModuloInt(v32, v77)];
		if (if (v4) then true else if (v79.fm2(v79.f19, v4, globalState)) then v79.f19 else v4) {
			var v83: array<int> := new int[12](i8 => i8 * |v1|);
			v83[safeIndex(845, v83.Length)] := v77;
			var v84 := DC12(v2);
			var v85 := DC15(v84);
			var v86: C6 := new C6(v4, v81);
			var v87: seq<C6> := [v86];
			var v88: map<seq<C6>, set<int>> := map[v87 := fm6(globalState)];
			v83[safeIndex(845, v83.Length)], globalState.f3, globalState.f3, v85 := safeDivisionInt(v0, 632) * |v88|, v32, v32, v85.(cf28 := v84);
			v83[safeIndex(845, v83.Length)] := v83[safeIndex(845, v83.Length)];
			var v89, v90, v91, v92 := v86.m5(globalState);
			var v93, v94, v95, v96 := v79.m0(globalState);
			var v97 := new C7();
		} else {
			var v98: array<char> := new char[6](i9 => v31);
			var v99: map<char, array<char>> := map['f' := v98];
			var v100: map<bool, int> := map[v79.f19 := v32];
			var v101: array<array<char>> := new array<char>[15] [v98, v98, v98, v98, if (fm25(v0, |v10|, v79.f19, v100, globalState) in v99) then v99[fm25(v0, |v10|, v79.f19, v100, globalState)] else v98, v98, v98, v98, v98, v98, v98, if (v79.f19) then v98 else v98, v98, v98, v98];
			v101[safeIndex(670, v101.Length)] := v98;
			var v102 := DC51(v98);
			v101[safeIndex(670, v101.Length)] := v102.cf83;
			var v103: array<int> := new int[7];
			v103[safeIndex(207, v103.Length)] := if (v79.f19) then v77 else v0;
			var v104: map<array<int>, bool> := map[v103 := if (v4) then v8[safeIndex(|v30|, |v8|)] else v4];
			var v105 := DC18(v79.f19);
			var v106: map<int, D11> := map[v0 := DC35()];
			var v107: set<string> := {v3};
			globalState.f9, globalState.f9, v4, v103[safeIndex(207, v103.Length)], v104 := v79.f19, v105.cf33, if (v106 != v106) then v79.f19 else v107 !! v107, safeModuloInt(v0, v0), v104;
			var v108: C3 := new C3(v103[safeIndex(207, v103.Length)], v3, v4, v81);
			var v109 := DC55(v108);
			var v110: multiset<C3> := multiset{v109.cf91, v108};
			var v111: seq<map<int, int>> := [(map[|v110| := v108.f20])[|v8[safeIndex(v103[safeIndex(207, v103.Length)], |v8|) := v79.f19]| := -v108.f20]];
			v111 := v111;
			v2[safeIndex(919, v2.Length)] := true;
			v2[safeIndex(919, v2.Length)] := v79.f19;
			var v112: map<bool, char> := map[v4 := v31];
			var v113: multiset<map<bool, char>> := multiset{v112[v79.f19 := v31], v112};
			globalState.f9 := v113 < v113;
		}
		
		var v114: map<array<bool>, char> := map[v2 := v31];
		var v115 := DC49(v114 + v114, -0x58);
		v115 := DC49(v114, v32);
	}
	
	v0 := v32;
	v4 := v4;
	var v116: multiset<bool> := multiset{v4, v4, v4, v4, v4};
	var v117 := new C6(v4, v116);
	if (v10 != v10) {
		v0 := fm0(v4, -v32 + 359, v32, v1, globalState);
		var v118 := DC24({v4});
		var v119: map<int, set<bool>> := map[v32 := v5];
		var v120: map<int, bool> := map[v0 := v4];
		var v121: seq<set<bool>> := [v5, {v4, v4, true}, {v4}];
		var v122: array<set<bool>> := new set<bool>[16] [v5, v5, v5, v118.cf45, v5, if (v0 in v119) then v119[v0] else v5, v5 + fm41(v32, |v120|, globalState), v121[safeIndex(v32, |v121|)] * {v4}, if (v4) then v5 else v5, fm41(|v10|, v0, globalState), {v4, v4}, v5, v5 - {v4, v4, v4, v4, v4}, v5 + v5, v5, v5];
		v122[safeIndex(78, v122.Length)] := v5 * v5;
		var v123: array<int> := new int[10] [-v0, v10[safeIndex(|v3|, |v10|)], |v8|, safeDivisionInt(v32, v0), v0, -(v0 + v0), v32, |(map[v4 := v31])[v4 := v31]|, v32 - v32, v0];
		v123[safeIndex(170, v123.Length)] := v0;
		var v124 := DC39(v32);
		v122[safeIndex(78, v122.Length)], globalState.f6, v123[safeIndex(170, v123.Length)], v32 := fm41(|v3|, v124.cf65, globalState), v5 * v5, -(v0 + |v3|), fm0(v4, v0, -0x23b, v1, globalState);
		v4 := !v4;
		v4 := true;
		v123[safeIndex(170, v123.Length)] := v0;
	} else {
		if (v4) {
			v116 := v116;
			var v125: set<int> := {78, v0};
			var v126: map<char, set<int>> := map[v31 := v125];
			var v127, v128, v129 := v117.m1((if (v31 in v126) then v126[v31] else v125) + {-0xde}, 0x1e, globalState);
			var v130: C3 := new C3(v128, v3, true, v116);
			var v131: map<int, C3> := map[safeDivisionInt(v128, v128) := v130];
			v131 := v131[v32 := v130];
			var v132: map<int, array<bool>> := map[v128 := v2];
			var v133 := new C2((map[|v132| := v4])[v32 := v4]);
			var v134: map<bool, int> := map[v129 := v0];
			v134 := v134[v129 := v0];
		} else {
			var v135, v136, v137, v138 := v117.m5(globalState);
			var v139: C7 := new C7();
			v139 := new C7();
			var v140: map<int, bool> := map[|map[v4 := v4]| := v138];
			var v141: array<array<int>> := new array<int>[28];
			var v142 := DC52(v0, v138, DC52(v136, v138, v141, true, v0).cf86, v4, v136);
			v140 := v140[v142.cf88 * v32 := v4];
			var v143: array<int> := new int[11](i10 => safeDivisionInt(i10, v136));
			v143[safeIndex(825, v143.Length)] := v10[safeIndex(v0, |v10|)];
			var v144: map<bool, int> := map[v138 := v32];
			var v145: set<int> := {v32};
			var v147: T0 := new C5(v136, v138, v116);
			var v148 := DC59(v147, v144);
			var v149: array<map<bool, int>> := new map<bool, int>[29] [v144[v138 := v135], v144, v144 + v144, fm1(v138, v32, v4, v145, globalState), v144, map[v4 := v32], v144, if (!v4) then fm1(v4, v32, !v138, set v146 : int | (0x3e2 <= v146) && (v146 < 0x307) :: (safeDivisionInt(v146, v32)), globalState) else v144, v144, v144, fm1(v4, 705, v4, v145, globalState), v148.cf98, v144, map[v138 := v135] + v144, v144, v144, v144 + map[v138 := v32], fm1(!true, v136, v138, v145, globalState) + v144, v144 + v144[v4 := v32], DC59(v147, v144).cf98, v144, v144 + v144, v144, v144, v144, v144, map[false := v0], v144 + fm1(v138, v135, v138, v145, globalState), map[v147.f14 := v0]];
			var v150: seq<map<bool, int>> := [v144];
			v149[safeIndex(514, v149.Length)] := v150[safeIndex(v135, |v150|)] + map[v147.f14 := v136];
			v31, v143[safeIndex(825, v143.Length)], v149[safeIndex(514, v149.Length)], v136, v31 := v31, v32, (v144 + v144) + (v144[v4 := v135] + v144), v0, v31;
			v0 := v136 - safeDivisionInt(v135, |v3|);
		}
		
		var v151: array<int> := new int[25];
		v151 := v151;
		v151[safeIndex(349, v151.Length)] := 0x32f;
		v151[safeIndex(349, v151.Length)] := if (true) then v32 else v32;
		v3 := v3;
		v151[safeIndex(349, v151.Length)] := v151[safeIndex(349, v151.Length)];
	}
	
	var i11 := 0;
	while (v4)
		decreases 100 - i11
	{
		if (i11 >= 100) {
			break;
		}
		
		i11 := i11 + 1;
		var v152: array<int> := new int[25](i12 => i12 * |multiset([DC10(map[v4 := v4])])|);
		v152[safeIndex(726, v152.Length)] := v32;
		v152[safeIndex(726, v152.Length)] := safeModuloInt(|v1|, v32);
		var v153: array<C3> := new C3[15];
		var v154: C3 := new C3(v0, "ihvjxhpxy", v4, v116);
		v153[safeIndex(945, v153.Length)] := v154;
		v153[safeIndex(945, v153.Length)] := new C3(|[true, v4, v4, v4]|, v154.f21 + v154.f21, v4, v116);
		var v155 := DC7(v8);
		var v156: set<D2> := {v155, v155, v155};
		globalState.f9 := DC7(v8) !in v156;
		v2[safeIndex(918, v2.Length)] := v4;
		var v157: map<bool, char> := map[v4 := v154.f21[safeIndex(-0x294, |v154.f21|)]];
		var v158: set<int> := {-|v157|};
		v2[safeIndex(918, v2.Length)] := v158 >= fm6(globalState);
	}
	var v159: array<int> := new int[13](i14 => i14 - |v8|);
	forall i13 | 0 <= i13 < v159.Length {
		v159[i13] := i13 - 0x20d;
	}
	v2[safeIndex(302, v2.Length)] := v4;
	v2[safeIndex(302, v2.Length)] := v4;
	var v160, v161, v162, v163 := v117.m0(globalState);
	var v164: map<int, bool> := map[v160 := v4];
	v164 := v164[-v0 := v2[safeIndex(302, v2.Length)]];
	var i15 := 0;
	while (|v10| <= safeDivisionInt(v0, v0))
		decreases 100 - i15
	{
		if (i15 >= 100) {
			break;
		}
		
		i15 := i15 + 1;
		var v165: multiset<seq<int>> := multiset{seq(abs(0x4c), i16  => (188)), seq(abs(0x2db), i17  => (v32)), v10};
		v165 := v165;
		var v166 := DC5(v4, v161);
		v163, v2[safeIndex(302, v2.Length)], v3 := if (v2[safeIndex(302, v2.Length)]) then v166.cf6 else v161, v163, DC13(v3, v160, v161, v2).cf22;
		var v167, v168 := m11(v29, |v3|, true, v1 > multiset([307, v32]), globalState);
		v2[safeIndex(302, v2.Length)] := false;
	}
	if (v2[safeIndex(302, v2.Length)]) {
		var v169: set<int> := {0x258};
		var v170, v171, v172 := v117.m1(v169, |v164|, globalState);
		v2[safeIndex(302, v2.Length)] := v163;
		var v173 := new C1();
		var v174: T0 := new C5(v160, v161 ==> v161, multiset{v4});
		v2[safeIndex(302, v2.Length)], v174 := v172, v174;
		v163 := v161;
	} else {
		var v175, v176 := m11(v29, fm0(v2[safeIndex(302, v2.Length)], v0, v0, fm30(v161, globalState), globalState), multiset(v8) <= v116, v163, globalState);
		var v177: map<int, string> := map[|v164| := v3];
		var v178 := DC26(v161, v177);
		var v179: set<D8> := {v178, v178, v178, v178};
		var v180: map<bool, bool> := map[v163 := !v161];
		v8, v163, v179, v1, v2[safeIndex(302, v2.Length)] := [v163] + [!v163, if (v2[safeIndex(302, v2.Length)] in v180) then v180[v2[safeIndex(302, v2.Length)]] else v161], v163, v179, v1, v2[safeIndex(302, v2.Length)];
		var v181: set<int> := {v0};
		var v182, v183, v184 := v117.m1({-v160, 0x13a} + v181, v32, globalState);
		globalState.f3 := v160;
		var v185 := new C0(!v163, !false);
	}
	
	print v0, "\n";
	print v1 == multiset{206}, "\n";
	print v2[0], "\n";
	print v2[1], "\n";
	print v2[2], "\n";
	print v3, "\n";
	print v4, "\n";
	print v5 == {true}, "\n";
	print v7[0], "\n";
	print v7[1], "\n";
	print v7[2], "\n";
	print v7[3], "\n";
	print v7[4], "\n";
	print v7[5], "\n";
	print v7[6], "\n";
	print v7[7], "\n";
	print v7[8], "\n";
	print v7[9], "\n";
	print v7[10], "\n";
	print globalState.f0, "\n";
	print globalState.f1 == multiset{}, "\n";
	print globalState.f2[0], "\n";
	print globalState.f2[1], "\n";
	print globalState.f2[2], "\n";
	print globalState.f3, "\n";
	print globalState.f4, "\n";
	print globalState.f5, "\n";
	print globalState.f6 == {true}, "\n";
	print globalState.f7, "\n";
	print globalState.f8, "\n";
	print globalState.f9, "\n";
	print globalState.f10, "\n";
	print globalState.f12[0], "\n";
	print globalState.f12[1], "\n";
	print globalState.f12[2], "\n";
	print globalState.f12[3], "\n";
	print globalState.f12[4], "\n";
	print globalState.f12[5], "\n";
	print globalState.f12[6], "\n";
	print globalState.f12[7], "\n";
	print globalState.f12[8], "\n";
	print globalState.f12[9], "\n";
	print globalState.f12[10], "\n";
	print globalState.f13 == {true, false}, "\n";
	print v8 == [true, true], "\n";
	print v9 == map[false := {true}, true := {true}], "\n";
	print v10 == [206, 206], "\n";
	print v29.cf70.cf70.cf70.cf68, "\n";
	print v29.cf70.cf70.cf70.cf69, "\n";
	print v30 == map[0 := 0], "\n";
	print v31, "\n";
	print v32, "\n";
	print |v33|, "\n";
	print v116 == multiset{true, true, true, true, true}, "\n";
	print i11, "\n";
	print v159[0], "\n";
	print v159[1], "\n";
	print v159[2], "\n";
	print v159[3], "\n";
	print v159[4], "\n";
	print v159[5], "\n";
	print v159[6], "\n";
	print v159[7], "\n";
	print v159[8], "\n";
	print v159[9], "\n";
	print v159[10], "\n";
	print v159[11], "\n";
	print v159[12], "\n";
	print v160, "\n";
	print v161, "\n";
	print v162 == map[map[-153 := true] := -804], "\n";
	print v163, "\n";
	print v164 == map[-153539075 := true, -2 := true], "\n";
	print i15, "\n";
}