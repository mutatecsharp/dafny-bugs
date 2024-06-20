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
datatype D0 = DC1(cf1: int, cf2: int, cf3: int) | DC0(cf0: bool) | DC2(cf4: D0)
datatype D1 = DC4(cf6: multiset<string>, cf7: int) | DC5(cf8: bool, cf9: bool) | DC3(cf5: char) | DC6(cf10: D1)
datatype D2 = DC8(cf12: int) | DC7(cf11: array<D1>)
datatype D3 = DC10(cf14: bool, cf15: int, cf16: bool, cf17: array<bool>, cf18: map<int, bool>) | DC9(cf13: array<bool>) | DC11(cf19: D3)
datatype D4 = DC13(cf21: char) | DC12(cf20: set<bool>)
datatype D5 = DC14(cf22: map<int, array<int>>)
datatype D6 = DC15(cf23: multiset<D4>)
datatype D7 = DC17(cf25: bool) | DC16(cf24: C0) | DC18(cf26: D7)
datatype D8 = DC20(cf28: bool, cf29: bool, cf30: bool, cf31: bool) | DC19(cf27: C2)
datatype D9 = DC21(cf32: string)
datatype D10 = DC22(cf33: seq<C3>)
datatype D11 = DC24(cf35: int) | DC25 | DC26(cf36: int, cf37: array<int>, cf38: bool) | DC23(cf34: seq<int>) | DC27(cf39: D11)
datatype D12 = DC28(cf40: C1)
datatype D13 = DC29(cf41: seq<bool>)
datatype D14 = DC30(cf42: seq<C0>)
datatype D15 = DC32(cf44: bool) | DC31(cf43: map<D1, char>)
class GlobalState {
	const f0 : int
	const f1 : bool
	var f2 : int
	const f3 : bool
	const f4 : set<set<bool>>
	const f5 : map<int, bool>
	const f6 : seq<int>
	var f7 : seq<int>
	const f8 : bool
	const f9 : bool
	var f10 : array<int>
	constructor (f0 : int, f1 : bool, f2 : int, f3 : bool, f4 : set<set<bool>>, f5 : map<int, bool>, f6 : seq<int>, f7 : seq<int>, f8 : bool, f9 : bool, f10 : array<int>) {
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
	}
	
}

class C0 {
	constructor () {
	}
	
	function fm7(p0: bool, p1: D1, p2: string, p3: bool, globalState: GlobalState): bool {
		(0x6b - |"wt"|) < -0xe1
	}
	function fm8(p0: int, p1: set<bool>, p2: int, globalState: GlobalState): int {
		(341 - -0x100) + |({789, |"lan"|, |{|[0x96]|, |[false]|}|, |map[|{|"ghqr"|}| := false]|} * {0x14a})|
	}
}

class C1 {
	var f14 : bool
	constructor (f14 : bool) {
		this.f14 := f14;
	}
	
	method m4(p0: map<bool, int>, globalState: GlobalState) returns (r0: string, r1: D1, r2: int) {
		var v0 := -769;
		r2 := v0;
		v0 := v0 + 0x166;
		var v1: seq<int> := [341, v0];
		var v2 := 'q';
		var v3: multiset<int> := multiset{v0};
		for i0 := v1[safeIndex(|("n")[safeIndex(v0, |"n"|) := v2]|, |v1|)] to |v3| * v0 {
			var v4: array<array<int>> := new array<int>[9];
			var v5: array<bool> := new bool[9];
			var v6: multiset<array<bool>> := multiset{v5, v5};
			var v7: map<bool, bool> := map[true := f14];
			var v8: seq<map<bool, bool>> := [v7, v7, v7, v7];
			var v9: map<int, int> := map[|v8| := v0];
			var v10: array<int> := new int[19] [|v1|, 0x395, i0, v0, -|{f14, f14}|, |v6|, v0, i0, 0x2b3, |v7|, i0, v0, i0, i0, 0x2db, i0, if (i0 in v9) then v9[i0] else fm4(globalState), v0, i0];
			v4[safeIndex(770, v4.Length)] := v10;
			var v11: map<int, bool> := map[i0 := f14];
			v4[safeIndex(770, v4.Length)], f14, f14 := v10, f14, !f14 <==> (if (i0 in v11) then v11[i0] else f14);
			var v12: set<int> := {|map[|p0| := f14]|, |p0|};
			var v13: map<set<int>, int> := map[v12 := v0];
			var v14: multiset<bool> := multiset{i0 >= |v13|};
			v14 := v14;
			globalState.f2 := i0;
			f14 := if (f14) then f14 else f14;
		}
		var v15: map<int, bool> := map[v0 := f14];
		v15 := v15[|v1| + v0 := f14];
		var v16: array<map<bool, bool>> := new map<bool, bool>[9];
		var v17: map<bool, bool> := map[f14 := !f14];
		v16[safeIndex(345, v16.Length)] := v17;
		v16[safeIndex(345, v16.Length)] := map[f14 := f14];
		var v18: seq<bool> := [!f14, f14, f14];
		var v19: array<bool> := new bool[11](i1 => f14);
		var v20 := DC10(true, v0, f14, v19, v15);
		match (fm5(v0, v0, v0, v18, globalState).(cf5 := fm6(|v15|, v20.cf14, !f14, 'a', globalState))) {
			case DC4(cf6, cf7) =>
				var v21 := new C0();
				var v22: seq<array<bool>> := [v19];
				var v23: array<array<bool>> := new array<bool>[23] [v19, if (f14) then v19 else v19, v19, v19, v19, v19, v19, v20.cf17, v19, v19, v19, v19, v22[safeIndex(cf7, |v22|)], v19, v19, v19, v19, v19, v19, v19, v19, v19, v19];
				v23[safeIndex(347, v23.Length)] := v19;
				v23[safeIndex(347, v23.Length)] := v19;
				var v24: map<bool, array<bool>> := map[fm0(f14, f14, f14, f14, globalState) := v19];
				v19 := if (f14 in v24) then v24[f14] else v19;
				v3 := v3;
			case DC5(cf8, cf9) =>
				var v25: map<int, string> := map[v0 := "vsixj"];
				var v26 := "oukgvej";
				var v27: map<int, int> := map[v0 := if (cf8) then 0x373 else |(if (v0 in v25) then v25[v0] else v26)|];
				v27 := v27[v0 := |"lysyusnd"|];
				var v28: map<bool, int> := map[if (cf8) then false else f14 := v0];
				v28 := v28[f14 || cf8 := 0xb1];
				var v29: array<char> := new char[15];
				v29[safeIndex(725, v29.Length)] := v2;
				v29[safeIndex(725, v29.Length)] := v2;
				var v30 := new C0();
			case DC3(cf5) =>
				var v31 := new C0();
				f14 := v18[safeIndex(if (|v16[safeIndex(345, v16.Length)]| in v3) then v3[|v16[safeIndex(345, v16.Length)]|] else v0, |v18|)];
				v19[safeIndex(870, v19.Length)] := false;
				var v32: map<array<bool>, int> := map[v19 := v0];
				var v33: seq<map<array<bool>, int>> := [map[v19 := v0], map[v19 := -v0]];
				v19[safeIndex(870, v19.Length)], v32, f14, f14, r0 := f14, (if (false) then v32 else v32) + v33[safeIndex(-v0, |v33|)], f14 <==> true, f14, fm9(f14, globalState);
				var v34 := "ihdetle";
				var v35: set<int> := {-v0, v0};
				var v36: multiset<string> := multiset{v34};
				var v37 := DC4(v36, v0);
				var v38: map<bool, seq<bool>> := map[v19[safeIndex(870, v19.Length)] := v18];
				var v39: array<int> := new int[17] [v0, v0, v0, v0, 0x89, safeDivisionInt(v0, v0), |v34| + |v35|, v0, v37.cf7, v0 + |v38|, if (v0 in v3) then v3[v0] else 0x2e9, v0, v0 + v0, fm4(globalState), safeModuloInt(0x182, v0), v0, v1[safeIndex(|v35|, |v1|)]];
				v39[safeIndex(802, v39.Length)] := v0;
				v39[safeIndex(802, v39.Length)] := v0;
			case DC6(cf10) =>
				var v40: array<D1> := new D1[8](i2 => DC6(DC3(v2)));
				var v41 := DC7(v40);
				v41 := v41;
				f14 := if (f14) then f14 else f14;
				var v42 := DC8(675);
				var v43: multiset<D2> := multiset{v42};
				globalState.f2 := |(if (f14) then multiset{v42} - multiset{v42, v42} else (multiset{v42})[v42 := abs(v0)] + v43)|;
				var v44 := new C0();
		}
		
		var v45 := "nwhkcx";
		var v46: seq<string> := [v45, v45[safeIndex(-v0, |v45|) := v2], v45, v45];
		r0 := v46[safeIndex(|v1| * v0, |v46|)];
		var v47 := DC5(f14, v45 < v45);
		r1 := v47;
		var v48: seq<array<bool>> := [v19, v19, v19, v19];
		r2 := safeModuloInt(|v48|, v0 * |v45|);
	}
}

class C2 {
	var f13 : int
	constructor (f13 : int) {
		this.f13 := f13;
	}
	
	method m3(p0: array<bool>, p1: string, p2: seq<int>, globalState: GlobalState) {
		for i0 := f13 to f13 {
			var v0 := true;
			var v1 := 'e';
			var v2: map<char, bool> := map[v1 := v0];
			var v3: map<int, bool> := map[i0 := v0];
			var v4: map<int, int> := map[|v3| := f13];
			f13 := (if (v0) then f13 else |v2|) + (i0 - |v4|);
			var v5: array<seq<bool>> := new seq<bool>[21](i1 => [v0]);
			v5 := v5;
			var v6: seq<bool> := [false];
			var v7: map<bool, bool> := map[fm0(true, !v0, v0, v6[safeIndex(f13, |v6|)], globalState) := v0];
			var v8: seq<map<int, int>> := [map[|p1| := i0]];
			var v9: array<int> := new int[29] [|p1|, p2[safeIndex(i0, |p2|)], 0x38c, i0, |v7|, f13 - fm2(v0, globalState), i0, f13, -i0, 0x3af, f13, f13, |p1|, 917, |(v8 + (seq(abs(681), i2  => (map[i0 := |{f13}|]))))|, -545 * i0, i0, i0, f13, f13, i0, f13, f13 - f13, f13, f13, |("ngpktyvor" + p1)|, i0, f13, |[v0]|];
			v9[safeIndex(171, v9.Length)] := |(seq(abs(-0x9b), i3  => (v1)))|;
			var v10 := DC0(v0);
			var v11: set<D0> := {v10};
			var v12: map<bool, set<D0>> := map[v0 := v11];
			var v13: map<D0, int> := map[v10 := 0x2a8];
			var v15: map<bool, set<D0>> := map[false := if (false in v12) then v12[false] else set v14 : D0 | v14 in v13 :: (v14)];
			var v16: set<bool> := {v0, v0, !false, v0, v0};
			v9[safeIndex(171, v9.Length)], v0 := |(if (v0 in v12) then v12[v0] else if (v0) then v11 else v11)|, v16 > {v0, v0};
			v0 := v0;
		}
		var v17: seq<array<bool>> := [p0];
		f13 := |v17|;
		var v18 := true;
		v18, v18 := f13 != f13, safeDivisionInt(f13, f13) > f13;
		v18 := false;
		for i4 := f13 to safeModuloInt(f13, f13) {
			var v19: set<bool> := {v18, false, true, fm0(v18, v18, v18, v18, globalState)};
			if (!!(v19 <= v19)) {
				var v20: array<string> := new string[25];
				var v21 := 'l';
				var v22: map<char, string> := map[v21 := p1];
				var v23 := DC3(v21);
				v20[safeIndex(281, v20.Length)] := if (v23.cf5 in v22) then v22[v23.cf5] else p1;
				v20[safeIndex(281, v20.Length)] := fm3(v18, p1, f13, globalState);
				var v24: array<int> := new int[20];
				v24[safeIndex(677, v24.Length)] := -i4;
				v24[safeIndex(677, v24.Length)] := i4;
				var v25: seq<seq<int>> := [p2];
				v25 := v25;
				var v26: map<bool, bool> := map[!v18 := v18];
				var v27: array<bool> := new bool[21] [if (v18) then fm0(v18, v18, v18, !v18, globalState) else v18, v19 < v19, !v18, v18, v18, v18, fm0(v18, v18, v18, v18, globalState), if (v18) then v18 else v18, fm0(v18, false, v18, false, globalState), v18, v18, if (v18 in v26) then v26[v18] else v18, if (v18 in v26) then v26[v18] else v18, v18, v18, v18, v18, v18, !v18, v18, false];
				v27 := new bool[17];
				var v28 := DC5(v18, v18);
				var v29: seq<D1> := [v28];
				var v30: array<D1> := new D1[5];
				var v31 := DC7(v30);
				var v32: multiset<string> := multiset{v20[safeIndex(281, v20.Length)], p1};
				var v33 := DC4(v32, i4);
				var v34: map<D1, array<D1>> := map[v33 := v30];
				var v35: array<array<D1>> := new array<D1>[29] [v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, if (!v18) then v30 else v30, v31.cf11, v30, if (v33 in v34) then v34[v33] else v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, if (v18) then v30 else v30, v30];
				v29, v18, v35 := ([DC5(v18, v18)])[safeIndex(safeDivisionInt(0x26c, f13), |[DC5(v18, v18)]|) := v28], !v18, v35;
			} else {
				var v36: map<int, bool> := map[0xdf := v18];
				var v37 := DC10(v18, f13, v18, p0, v36);
				var v38 := 'i';
				var v39: map<array<bool>, char> := map[v37.cf17 := v38];
				v39 := v39[p0 := if (v18) then v38 else v38];
				var v40: array<int> := new int[1];
				v40[safeIndex(425, v40.Length)] := i4;
				f13, globalState.f10, v40[safeIndex(425, v40.Length)] := f13, v40, f13;
				var v41: array<bool> := new bool[15](i5 => v18);
				v41 := v41;
				v18, f13, globalState.f2, globalState.f2, v40[safeIndex(425, v40.Length)] := v18, i4 * fm2(v18, globalState), f13 + (v40[safeIndex(425, v40.Length)] * f13), i4, f13;
				var v42: map<int, int> := map[safeDivisionInt(0x1ca, f13) := |p2| + i4];
				v42 := v42;
			}
			
			var v43: array<map<D0, int>> := new map<D0, int>[13];
			v43 := v43;
			var v44 := 'a';
			var v45: set<char> := {v44};
			var v46: map<int, bool> := map[|p1| := v18];
			var v47 := DC10(v18, |v45|, v18, v17[safeIndex(f13, |v17|)], v46);
			if (v47.cf16) {
				var v48: map<bool, bool> := map[v18 := v18 || v18];
				v18 := if (!(v18 ==> v18) in v48) then v48[!(v18 ==> v18)] else false;
				f13 := f13;
				var v49: array<bool> := new bool[17];
				v49 := new bool[3];
				var v50: multiset<int> := multiset{i4};
				f13, globalState.f2, v18, v44, v18 := f13, f13, (if (fm2(true, globalState) in v50) then v50[fm2(true, globalState)] else i4) in (map[--i4 := v18] + v46[f13 := v18]), 'c', v18;
				var v51: array<map<bool, bool>> := new map<bool, bool>[13];
				var v52: multiset<bool> := multiset{v18};
				m0(p0, f13, v51, |v52|, globalState);
			} else {
				p0[safeIndex(573, p0.Length)] := v18;
				var v53: array<map<string, bool>> := new map<string, bool>[29];
				var v54: map<string, bool> := map[p1 := true];
				v53[safeIndex(729, v53.Length)] := v54;
				p0[safeIndex(573, p0.Length)], v53[safeIndex(729, v53.Length)] := v18, v54;
				var v55: array<array<bool>> := new array<bool>[25];
				var v56: array<bool> := new bool[18];
				v55[safeIndex(897, v55.Length)] := v56;
				v55[safeIndex(897, v55.Length)] := v56;
				var v57: multiset<bool> := multiset{p0[safeIndex(573, p0.Length)], true, !p0[safeIndex(573, p0.Length)]};
				v57 := v57;
				var v58 := new C1(v18);
				var v59: seq<string> := [seq(abs(-0x3a2), i6  => (v44)), p1];
				var v60: map<bool, bool> := map[p0[safeIndex(573, p0.Length)] := true];
				var v61: map<bool, int> := map[v58.f14 := |v60|];
				var v62: set<int> := {|p2|, if (fm0(v18, v58.f14, v58.f14, p0[safeIndex(573, p0.Length)], globalState) in v61) then v61[fm0(v18, v58.f14, v58.f14, p0[safeIndex(573, p0.Length)], globalState)] else i4, f13};
				var v63: array<int> := new int[5](i7 => i7 - i4);
				var v64: seq<array<int>> := [v63, v63, v63, v63];
				v18, v44, f13, v59, p0[safeIndex(573, p0.Length)] := v19 !! {v18}, p1[safeIndex(|v62|, |p1|)], |(v64 + ([v63, v63] + v64))|, (v59[safeIndex(i4, |v59|) := p1])[safeIndex(-f13, |v59[safeIndex(i4, |v59|) := p1]|) := p1 + p1], false;
			}
			
			var v65: map<bool, set<bool>> := map[v18 := v19];
			var v66 := DC12(v19);
			var v67: map<int, set<bool>> := map[i4 := v66.cf20];
			var v68: multiset<int> := multiset{f13};
			var v69: seq<bool> := [v18, false];
			var v70: map<bool, bool> := map[v18 := v69[safeIndex(i4, |v69|)]];
			var v71: map<bool, char> := map[!v18 := v44];
			var v72: array<int> := new int[29] [647, |(if (v18) then p1 else p1)[safeIndex(f13, |(if (v18) then p1 else p1)|) := v44]|, safeDivisionInt(|(seq(abs(466), i8  => ('b')))|, p2[safeIndex(i4, |p2|)]), f13 - |v19|, |(if (v18 in v65) then v65[v18] else if (f13 in v67) then v67[f13] else v19)|, i4, safeModuloInt(f13, 0xb6), safeModuloInt(f13, f13), i4, f13, |p1|, if (|v69| in v68) then v68[|v69|] else i4, i4, f13, i4, i4, 0x305, |v70|, fm4(globalState), |(p1 + p1)|, f13, i4, i4, --|multiset{v44, v44, v44, if (v18 in v71) then v71[v18] else v44, v44}|, f13, f13, i4, i4, f13];
			v72[safeIndex(878, v72.Length)] := f13;
			v72[safeIndex(878, v72.Length)] := f13;
		}
		v18 := ([f13] + p2[safeIndex(f13, |p2|) := f13]) <= (seq(abs(0x322), i9  => (f13)));
	}
}

class C3 {
	const f11 : bool
	const f12 : int
	constructor (f11 : bool, f12 : int) {
		this.f11 := f11;
		this.f12 := f12;
	}
	
	function fm1(p0: string, p1: int, globalState: GlobalState): int {
		safeDivisionInt(safeDivisionInt(|"ngc"|, |multiset{f11}|), safeDivisionInt(|"ttwonxung"|, |multiset([seq(abs(0xad), i0  => (f12))])|))
	}
	method m1(p0: bool, globalState: GlobalState) {
		var v0 := DC0(p0);
		var v1 := DC2(v0);
		var v2: array<char> := new char[14];
		var v3: map<D0, array<char>> := map[v1 := v2];
		if (DC2(v0) in v3) {
			var v4 := new C2(-f12);
			var v5: multiset<int> := multiset{v4.f13 - f12, f12};
			v5 := v5;
			var v6: array<int> := new int[5];
			var v7: seq<bool> := [p0];
			v6[safeIndex(559, v6.Length)] := |v7|;
			v6[safeIndex(559, v6.Length)] := safeModuloInt(-0x36e, |fm10(globalState)|);
			var v8 := false;
			var v9: set<int> := {v6[safeIndex(559, v6.Length)]};
			v8 := !(v9 != v9);
			var v10 := 'e';
			var v11 := "wdfkv";
			if (fm0(v10 in v11, if (f11) then v8 else f11, f11, f11, globalState)) {
				globalState.f2 := v6[safeIndex(559, v6.Length)] * v4.f13;
				v8 := |v7| < f12;
				var v12: array<bool> := new bool[23];
				v12[safeIndex(560, v12.Length)] := p0;
				var v13: multiset<seq<bool>> := multiset{v7};
				v12[safeIndex(560, v12.Length)] := (v13 + v13) >= v13;
				v4.f13 := v6[safeIndex(559, v6.Length)];
				v6[safeIndex(559, v6.Length)] := f12 * v4.f13;
			} else {
				v1 := v1;
				var v14 := new C1(v5 > v5);
				var v15: map<int, char> := map[v6[safeIndex(559, v6.Length)] := v10];
				var v16: map<char, bool> := map[if (f12 in v15) then v15[f12] else v10 := v10 in v11];
				v8 := if ('o' in v16) then v16['o'] else f11;
				var v17: array<bool> := new bool[9] [v8 <==> v7[safeIndex(v4.f13, |v7|)], f11, !v14.f14, f11, v7[safeIndex(-640, |v7|)], v14.f14, p0, v14.f14, f11];
				v17 := v17;
				var v18: array<seq<bool>> := new seq<bool>[17](i0 => v7);
				v18 := v18;
			}
			
		} else {
			var v19: seq<int> := [f12, f12];
			var v20: seq<int> := [f12, |v19|, 0x203];
			var v21: multiset<bool> := multiset{v19 < v20};
			v21 := if (p0) then v21 else v21;
			var v22: C2 := new C2(fm4(globalState));
			v22 := v22;
			var v23 := true;
			v23 := v23;
			var v24: map<int, int> := map[v22.f13 := f12 - f12];
			v23, v23, v23, v23, v24 := v23, p0, f12 > 80, p0, v24 + v24;
			var v25 := new C1(f11);
		}
		
		var v26 := false;
		v26 := !p0;
		var i1 := 0;
		while (p0)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			var v27: multiset<bool> := multiset{f11, f11};
			var v28 := "dcu";
			var v29 := 'h';
			var v30: multiset<string> := multiset{v28[safeIndex(f12, |v28|) := v29], "rbbmc"};
			var v31, v32, v33 := m2(f12 >= |v27|, v30 < v30, globalState);
			var v34: map<bool, set<bool>> := map[!v26 := {f11, v31, v31}];
			var v35: set<bool> := {f11};
			var v36 := DC12(v35);
			v34 := v34[p0 := v36.cf20];
			v28 := v28 + "ycmdjta";
			var v37: map<int, bool> := map[0x228 := v26];
			globalState.f2 := f12 + safeModuloInt(f12, |v37|);
		}
		for i2 := safeModuloInt(f12, f12) to -f12 {
			var v38: seq<bool> := [true];
			v38 := v38[safeIndex(i2, |v38|) := p0] + (v38[safeIndex(f12, |v38|) := f11])[safeIndex(-i2, |v38[safeIndex(f12, |v38|) := f11]|) := !f11];
			var v39: array<bool> := new bool[15](i3 => v26);
			var v40: map<array<bool>, array<bool>> := map[v39 := v39];
			var v41: seq<int> := [f12];
			var v42: multiset<int> := multiset{f12, |v41|};
			var v43: map<bool, multiset<int>> := map[v26 := v42];
			var v44: map<bool, bool> := map[f11 := f11];
			var v45: map<int, bool> := map[f12 := p0];
			var v46: map<int, int> := map[|v44| := |v45|];
			var v47: array<int> := new int[9] [f12, i2, |v46|, i2, f12, i2, f12, i2, i2];
			var v48 := 'e';
			var v49: seq<char> := [v48];
			var v50: seq<multiset<int>> := [v42];
			var v51 := DC13(v48);
			var v52: multiset<D4> := multiset{v51};
			var v53 := DC15(v52);
			var v54: set<multiset<int>> := {(if (v26 in v43) then v43[v26] else multiset{|DC14(map[i2 := v47]).cf22|})[|v49| := abs(-13)], v42 + multiset{f12, f12}, v42, v50[safeIndex(|multiset{|map[v53.cf23 := f11]|, i2, f12}|, |v50|)], v42 - v42};
			var v55: set<bool> := {f11};
			v26, v40, v54 := (v55 * v55) !! v55, v40, v54;
			v45 := v45[f12 - f12 := v26];
			v26 := true;
		}
		var v56 := 'g';
		if (fm0(v56 !in (seq(abs(-0x2b1), i4  => (v56))), v26, v26 && f11, f11, globalState)) {
			var v57 := new C2(|"my"|);
			if (f11) {
				v26 := v26 ==> !!p0;
				var v58 := "xoku";
				var v59: set<bool> := {p0, !!f11, v26, v26, v26};
				v58, v26 := "sedfiwcj", v59 != v59;
				var v60: set<int> := {|[!v26]|};
				v26 := !(v60 >= (v60 + {0x32, v57.f13}));
				v26 := v26;
				var v61 := DC0(true);
				var v62: seq<bool> := [p0 ==> v61.cf0, !f11];
				v62 := v62;
			} else {
				var v63 := new C2(v57.f13);
				var v64, v65, v66 := m2(v26, v26, globalState);
				var v67: map<bool, bool> := map[v66 := p0];
				var v68 := "yfahqywud";
				var v70: seq<string> := [seq(abs(80), i5  => ('l'))];
				var v71: multiset<int> := multiset{v57.f13, |v67|, |v68|, |[|(map v69 : string | v69 in v70 :: (v69) := (f11))|]|};
				var v72: C1 := new C1(v71 <= fm11(fm0(v26, f11, p0, f11, globalState), v56, globalState));
				v72 := v72;
				var v73 := new C1(p0);
				v72.f14 := v64;
			}
			
			globalState.f2 := safeModuloInt(v57.f13, f12);
			v26 := v26;
			var v74 := new C1(v57.f13 == 0xb8);
		} else {
			var v75: multiset<char> := multiset{v56, v56};
			v26 := v75 !! (v75 + v75);
			var v76: C2 := new C2(f12);
			var v77: seq<bool> := [f11, f11];
			var v78: map<int, bool> := map[|v77| := f11];
			var v79: multiset<bool> := multiset{f11, v26, f11};
			var v80 := "onxfcxpf";
			var v81: array<bool> := new bool[6];
			var v82: array<int> := new int[15] [f12, f12, |v79|, 584 - v76.f13, -(v76.f13 + f12), safeModuloInt(|v80|, v76.f13), f12, f12 * v76.f13, f12, if (f11) then f12 else v76.f13, fm1("lnlmx", fm2(f11, globalState), globalState), f12, |map[v81 := true]|, f12, f12];
			v82[safeIndex(700, v82.Length)] := v76.f13;
			var v83: C0 := new C0();
			var v84 := DC16(v83);
			var v85: array<C0> := new C0[27] [v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v83, v84.cf24, v83, v83];
			var v86: seq<array<C0>> := [v85, v85, v85, v85];
			v76, v26, v78, v82[safeIndex(700, v82.Length)], v86 := DC19(v76).cf27, f11, v78 + (DC10(v26, f12, p0, v81, v78).cf18 + v78), f12, v86 + v86;
			var v87 := DC13('f');
			var v88: map<bool, D4> := map[p0 := v87];
			var v89: seq<int> := [v76.f13, |v88|];
			var v90: seq<int> := [|v89|, v76.f13];
			var v91: seq<int> := [|v90|];
			v76.m3(v81, v80, v89 + (seq(abs(0x38f), i6  => (-0x12a))), globalState);
			if (v26) {
				v76.f13 := v76.f13;
				v76.f13, v80 := v76.f13 * (f12 + |v91|), "ebyufrei" + v80;
				globalState.f2 := v82[safeIndex(700, v82.Length)];
				v80 := v80[safeIndex(f12, |v80|) := v56];
				v81[safeIndex(891, v81.Length)] := p0;
				v81[safeIndex(891, v81.Length)] := true;
			} else {
				var v92: multiset<string> := multiset{"wtrbr"};
				var v93 := DC4(v92 - v92, -v82[safeIndex(700, v82.Length)]);
				v93 := v93;
				globalState.f2 := v82[safeIndex(700, v82.Length)] + f12;
				v81[safeIndex(395, v81.Length)] := p0;
				v81[safeIndex(395, v81.Length)] := "brsji" != v80;
				globalState.f2 := --safeModuloInt(v76.f13, |(map v94 : int | v94 in v78 :: (safeDivisionInt(v94, f12)) := (v56))|);
				var v95: map<int, string> := map[if (v26) then v76.f13 else v82[safeIndex(700, v82.Length)] := v80];
				v95 := v95;
			}
			
			var v96: set<bool> := {!p0, f11};
			v26 := !({v89[safeIndex(v76.f13, |v89|)], |multiset{true}|, f12} !! {v82[safeIndex(700, v82.Length)]}) in v96;
		}
		
		var v97 := "ilyuyef";
		var v98: multiset<bool> := multiset{v26};
		var v99: map<string, int> := map[v97 := |v98|];
		var v100: array<int> := new int[27];
		var v101: map<map<string, int>, array<int>> := map[v99 := v100];
		var v102: multiset<int> := multiset{192};
		var v103: multiset<multiset<int>> := multiset{v102, v102, multiset{f12}, (multiset{f12})[f12 := abs(f12)], v102};
		v101 := v101[(v99["ufaahfrp" := f12])["yaws" := if (v102 in v103) then v103[v102] else |v98|] := v100];
	}
	method m2(p0: bool, p1: bool, globalState: GlobalState) returns (r0: bool, r1: map<seq<bool>, bool>, r2: bool) {
		var v0 := "iteghrvta";
		var v1: map<string, int> := map[v0 := f12];
		v1 := v1[v0 := f12];
		var v2: array<bool> := new bool[23];
		var v3: array<map<bool, bool>> := new map<bool, bool>[19];
		var v4 := 'w';
		var v5: multiset<char> := multiset{v4};
		var v6: seq<int> := [fm2(p0, globalState)];
		m0(v2, f12, v3, if (v4 in v5) then v5[v4] else |v6|, globalState);
		var v7 := new C0();
		var v8: array<int> := new int[29](i0 => safeDivisionInt(i0, f12));
		globalState.f10 := v8;
		if (p0) {
			v2[safeIndex(312, v2.Length)] := p0 <==> f11;
			var v9: map<bool, bool> := map[p1 := p1];
			var v10: set<bool> := {if (p1 in v9) then v9[p1] else p1, f11};
			var v11: map<set<bool>, bool> := map[v10 := p0];
			v2[safeIndex(312, v2.Length)] := if (v10 in v11) then v11[v10] else f11;
			if (f11) {
				var v12: multiset<bool> := multiset{p1};
				r2 := (v12 > v12) ==> p0;
				var v13: multiset<set<bool>> := multiset{v10};
				var v14: seq<bool> := [!!(v13 !! v13)];
				v14 := v14;
				v4, v2[safeIndex(312, v2.Length)] := fm6(f12, p0, fm0(f11, true, p1, f11, globalState), v4, globalState), !v2[safeIndex(312, v2.Length)];
				var v15 := DC5(p0, p0);
				r2 := v7.fm7(if (p1) then v15.cf9 else f11, DC3(v4), v0, v2[safeIndex(312, v2.Length)], globalState);
				v8[safeIndex(941, v8.Length)] := f12;
				v8[safeIndex(941, v8.Length)] := safeDivisionInt(safeModuloInt(f12, f12), f12);
			} else {
				var v16: array<D8> := new D8[10];
				var v17: C2 := new C2(f12);
				var v18 := DC19(v17);
				v16[safeIndex(588, v16.Length)] := if (p1) then v18 else DC19(v17);
				v16[safeIndex(588, v16.Length)] := v18;
				var v19: map<int, int> := map[v17.f13 := |v0|];
				var v20: map<bool, map<int, int>> := map[v2[safeIndex(312, v2.Length)] := v19];
				v20 := v20[false := v19];
				var v21 := DC17(v2[safeIndex(312, v2.Length)]);
				var v22 := DC18(v21);
				var v23 := DC18(v21);
				var v24: map<D7, bool> := map[v23 := v2[safeIndex(312, v2.Length)]];
				v8[safeIndex(897, v8.Length)] := v6[safeIndex(|v24|, |v6|)];
				v8[safeIndex(897, v8.Length)] := v17.f13;
				var v25: array<int> := new int[21](i1 => i1 * v8[safeIndex(897, v8.Length)]);
				globalState.f10, globalState.f2, v2[safeIndex(312, v2.Length)], globalState.f2, globalState.f7 := v25, v17.f13, p0, v8[safeIndex(897, v8.Length)], fm12(v17.f13, globalState) + (v6 + v6)[safeIndex(f12, |(v6 + v6)|) := 0x196];
				var v26: array<array<bool>> := new array<bool>[26];
				var v27: array<bool> := new bool[12];
				v26[safeIndex(441, v26.Length)] := v27;
				v26[safeIndex(441, v26.Length)] := v27;
			}
			
			globalState.f10 := v8;
			v9 := v9[f11 := v2[safeIndex(312, v2.Length)]];
			var v28 := new C2(f12 * f12);
		} else {
			var v29: map<int, array<int>> := map[f12 := v8];
			v29 := v29;
			globalState.f7 := v6;
			var v30: C2 := new C2(|v0|);
			var v31: set<C2> := {v30};
			globalState.f2 := |({v30} - v31)|;
			var v32 := DC21(v0);
			var v33: array<string> := new string[13] [v32.cf32 + v0, v0, v0 + v0, v0, v0, "kctfg" + (seq(abs(908), i2  => (v4))), v0, v0, v0, v0 + v0, v0, v0 + fm9(p1, globalState), v0];
			v33[safeIndex(651, v33.Length)] := v0;
			v33[safeIndex(651, v33.Length)] := v0;
			var v34: set<int> := {f12};
			var v35: multiset<bool> := multiset{f11};
			var v36: map<int, int> := map[f12 := |v35|];
			v34 := (set v37 : int | v37 in v36 :: (safeModuloInt(v37, -|multiset([365])|))) - v34;
		}
		
		globalState.f2 := f12;
		r0 := f11;
		var v38: seq<bool> := [!f11, p0, f11];
		var v39: map<seq<bool>, bool> := map[v38 := p0];
		r1 := v39;
		r2 := v38 <= v38;
	}
}

function fm0(p0: bool, p1: bool, p2: bool, p3: bool, globalState: GlobalState): bool {
	false
}
function fm2(p0: bool, globalState: GlobalState): int {
	-safeDivisionInt(-490, safeModuloInt(|["guwwkcoj"]|, -0x264))
}
function fm3(p0: bool, p1: string, p2: int, globalState: GlobalState): string {
	((seq(abs(309), i0  => ('n'))) + "dwtkddnh") + ("kq" + (seq(abs(218), i1  => ('x'))))
}
function fm4(globalState: GlobalState): int {
	-(if (!(true && false)) then 512 + |map[true := 's']| else safeModuloInt(|[|map[268 := true]|, 0x24d, |map[false := false]|, |"ddwiwxcq"|]|, -0x22b))
}
function fm5(p0: int, p1: int, p2: int, p3: seq<bool>, globalState: GlobalState): D1 {
	DC3('d')
}
function fm6(p0: int, p1: bool, p2: bool, p3: char, globalState: GlobalState): char {
	's'
}
function fm9(p0: bool, globalState: GlobalState): string {
	(seq(abs(0x31b), i0  => ('p'))) + "l"
}
function fm10(globalState: GlobalState): set<int> {
	(set v2 : int | v2 in map[|multiset{144, |(set v1 : int | v1 in (set v0 : int | (66 <= v0) && (v0 < 817) :: (safeModuloInt(v0, 0x191))) :: (safeModuloInt(v1, |{true, true}|)))|, -|[true]|}| := true] :: (safeDivisionInt(v2, -|"vup"|))) * (if (!true) then set v3 : int | (0x2c0 <= v3) && (v3 < -0x1af) :: (v3 - -676) else {|"qdem"|, |[true, true, true, !DC0(true).cf0, false]|})
}
function fm11(p0: bool, p1: char, globalState: GlobalState): multiset<int> {
	(multiset{|[true, false, false, !!!false, true]|} + multiset{-0x301, |[0x31e]|}) - (multiset([-0x26b]) + multiset{|map[true := true]|, |(seq(abs(752), i0  => ('v')))|, 803, |map[DC5(false, false).cf8 := true]|, 0x288})
}
function fm12(p0: int, globalState: GlobalState): seq<int> {
	if (if (true) then true else true) then [-0x192] + [|"mbafox"|, |{map[!false := !!false], map[false := false], map[true := true], map[true := true]}|] else [-643]
}
function fm13(p0: char, p1: multiset<bool>, globalState: GlobalState): D1 {
	DC5(true, {[false], [!true, true, false, true, true]} >= {[false], [true]})
}
function fm14(p0: string, globalState: GlobalState): map<bool, map<int, int>> {
	map[false := map v0 : int | (920 <= v0) && (v0 < 0x65) :: (v0 + |[false]|) := (|multiset(seq(abs(0x26d), i0  => (|"yfuahe"|)))|)] + (map[true := map[|(set v1 : int | v1 in [|[DC25(), DC25(), DC25()]|, -0x19b] :: (safeModuloInt(v1, 0x118)))| := 0x72]] + map[false := map[0x1a8 := 0x9a]])
}
function fm15(p0: int, p1: int, p2: bool, globalState: GlobalState): set<bool> {
	({true} - {false}) - ({false} * {false})
}
function fm16(p0: int, p1: bool, globalState: GlobalState): seq<bool> {
	[!(['g'] == (seq(abs(0x102), i0  => ('w')))), true, false || false]
}
function fm17(p0: bool, p1: seq<int>, p2: bool, p3: bool, globalState: GlobalState): D1 {
	DC4(multiset(seq(abs(216), i0  => ("nrkvxhch"))) * multiset(["ugtro"]), 0x95)
}
function fm18(p0: int, globalState: GlobalState): set<string> {
	({"odvbie"} * {"efm", "ctebj"}) - ({"wmdxukk", "m", "dlv"} - {"affp", "v"})
}
function fm19(globalState: GlobalState): map<int, int> {
	map[-|"cjrbft"| := 0x24b] + (map[-296 := 0x1d0] + map[0x219 := |[|(map v0 : int | (932 <= v0) && (v0 < 0x12e) :: (safeDivisionInt(v0, --0x6a)) := (false))|]|])
}
function fm20(p0: bool, p1: int, p2: bool, globalState: GlobalState): map<int, bool> {
	map[|{!false, true}| := false] + (map[|[[|multiset{true}|], [0x31b], [0x209]]| := !false] + map[|multiset{|map[0x1f9 := true]|, |map['q' := |[!true]|]|, 0x109}| := !false])
}
function fm21(p0: bool, p1: bool, p2: int, p3: bool, globalState: GlobalState): map<D1, char> {
	match DC1(-|(seq(abs(0x38c), i0  => ('y')))|, -506, 0x232) {
		case DC1(cf1, cf2, cf3) => map v0 : D1 | v0 in [DC6(DC3('f')), DC6(DC5(false, true))] :: (v0) := ('y')
		case DC0(cf0) => map[DC6(DC4(multiset([seq(abs(0x378), i1  => ('r'))]), 0x2d0)) := 'k']
		case DC2(cf4) => DC31(map[DC6(DC6(DC3('g'))) := 'r']).cf43
	}
}
method m0(p0: array<bool>, p1: int, p2: array<map<bool, bool>>, p3: int, globalState: GlobalState) {
	var v0 := false;
	var v1: seq<bool> := [v0, v0, v0];
	var v2: map<bool, seq<bool>> := map[v0 := v1];
	v2 := v2[true := v1];
	if (v1[safeIndex(p3, |v1|)]) {
		var v3 := DC10(v0, p3, v0, p0, map[p1 := !v0]);
		var v4 := DC11(v3);
		v4 := v4;
		var v5 := DC5(v0, v0);
		v0 := v5.cf8;
		var v6 := DC5(v0, v0);
		var v7 := DC6(v6);
		var v8: C0 := new C0();
		var v9: multiset<bool> := multiset{v0, true};
		var v10: seq<int> := [p1];
		var v11: seq<int> := [-|v9|, p3, |v10|];
		var v12: seq<int> := [p1, |v11[safeIndex(p3, |v11|) := p1]|];
		var v13: C2 := new C2(p1);
		var v14: map<int, C2> := map[p1 := v13];
		var v15: map<bool, int> := map[v0 := |v14|];
		var v16: C1 := new C1(false);
		var v17: map<bool, C1> := map[v0 := v16];
		var v18: array<int> := new int[24] [p1, p3, p3, v12[safeIndex(p1, |v12|)], p1, p3, p1, p3, p1, p3, |v10|, p3, p3, -p1, 0x1, p1, if (v0 in v15) then v15[v0] else p3, 0x255, p3, p1, -0x1d5, v13.f13, 0x380, |v17|];
		var v19: seq<array<int>> := [v18, v18];
		var v20: array<array<int>> := new array<int>[22] [v18, v18, v18, v18, v18, v18, v18, v18, v19[safeIndex(p3, |v19|)], v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18];
		v7, globalState.f10, v0, v8, v20 := v7, v18, v0, v8, v20;
		var v21 := DC29(v1);
		var v22: array<seq<bool>> := new seq<bool>[13] [if (v16.f14) then v1 else v1, v1, v1, [v16.f14, v16.f14, v16.f14, v16.f14] + v1, [v16.f14], v1, v1, v1 + v1, [v0, v16.f14], v1 + v1, v1 + v1, v21.cf41 + v1, v1];
		var v23 := 'l';
		v22, v23, v16.f14 := v22, v23, p3 >= -v13.f13;
		var v24: set<seq<bool>> := {v1};
		v24 := v24 + v24;
	} else {
		p0[safeIndex(491, p0.Length)] := v0;
		var v25: array<seq<int>> := new seq<int>[10];
		var v26: seq<int> := [p1, p3];
		v25[safeIndex(855, v25.Length)] := v26;
		var v27 := "wwcctfdwq";
		var v28: map<int, bool> := map[p3 := !v0];
		var v29 := DC17(v0);
		var v30 := DC23(v26);
		p0[safeIndex(491, p0.Length)], globalState.f2, v0, v25[safeIndex(855, v25.Length)] := (233 in v26) && (v27 <= "g"), |(v28 + fm20(v0, -p3, fm0(true, v0, v29.cf25, v0, globalState), globalState))|, if (if (!fm0(v0, v0, v0, v0, globalState)) then v0 else false) then v0 else if (v0) then v0 else v0, v30.cf34[safeIndex(p1 * p3, |v30.cf34|) := -p1];
		p0[safeIndex(491, p0.Length)] := v0;
		var v31: C1 := new C1(v0);
		var v32: map<bool, C1> := map[p1 <= p1 := v31];
		var v33: map<int, C1> := map[v26[safeIndex(p1, |v26|)] := v31];
		v31 := if (p0[safeIndex(491, p0.Length)] in v32) then v32[p0[safeIndex(491, p0.Length)]] else if (p3 in v33) then v33[p3] else v31;
		globalState.f2 := v25[safeIndex(855, v25.Length)][safeIndex(|v27|, |v25[safeIndex(855, v25.Length)]|)];
		var v34: multiset<bool> := multiset{v0, v0};
		var v35: C2 := new C2(|v34|);
		var v36 := 'x';
		var v37: map<C2, char> := map[v35 := v36];
		globalState.f2 := p1 - |v37|;
	}
	
	if (v0) {
		v0 := v0;
		var v38: C0 := new C0();
		var v39: seq<C0> := [v38, v38];
		var v40 := DC30(v39);
		var v41: map<seq<C0>, int> := map[v40.cf42 := -p1];
		var v42: map<int, int> := map[0x2ce := fm4(globalState)];
		var v43: map<bool, bool> := map[v0 := v0];
		var v44: seq<map<bool, bool>> := [v43];
		v41 := v41[v39[safeIndex(p1, |v39|) := v38] := (if (|v44| in v42) then v42[|v44|] else p3) - p1];
		var v45: set<int> := {p3};
		var v46: map<bool, int> := map[!({p1} > v45) := p3];
		v46 := v46[fm0(v0, v0, !v0, v0, globalState) := p3];
		var v47 := DC24(p3);
		v47 := v47;
		if (!v0) {
			var v48 := new C0();
			v0 := v0;
			var v49: array<int> := new int[26];
			var v50: C2 := new C2(p3);
			var v51: map<C2, array<int>> := map[v50 := v49];
			var v52: array<array<int>> := new array<int>[25] [v49, v49, v49, v49, v49, v49, v49, v49, v49, v49, v49, v49, v49, v49, if (v50 in v51) then v51[v50] else v49, v49, v49, v49, v49, v49, v49, v49, v49, v49, v49];
			var v53: map<char, array<int>> := map['k' := v49];
			var v54 := 'g';
			v52[safeIndex(704, v52.Length)] := if (v54 in v53) then v53[v54] else v49;
			v52[safeIndex(704, v52.Length)] := v49;
			var v55: map<C0, int> := map[v38 := |v46|];
			v55 := v55[v38 := p3];
			globalState.f2 := v50.f13;
		} else {
			var v56: array<int> := new int[12](i0 => safeModuloInt(i0, p3));
			var v57: map<array<int>, int> := map[v56 := p1];
			var v58: multiset<int> := multiset{p3, p3};
			var v59: set<bool> := {v0, v0, v0};
			v42 := v42[if (v56 in v57) then v57[v56] else v38.fm8(|v58|, v59, p1, globalState) := p1];
			var v60: array<D1> := new D1[2](i1 => DC6(DC6(DC3('h'))));
			var v61 := DC7(v60);
			var v62 := 'm';
			globalState.f2, v0, v61 := p1 - safeDivisionInt(p3, |fm11(v0, v62, globalState)|), -p3 >= -p1, v61;
			globalState.f2 := --p1;
			var v63 := new C2(p1);
			v56[safeIndex(895, v56.Length)] := fm2(v0, globalState);
			v56[safeIndex(895, v56.Length)] := v63.f13;
		}
		
	} else {
		var v64 := DC18(DC17(v0));
		match (v64) {
			case DC17(cf25) =>
				cf25 := (p1 != p3) && !v0;
				var v65: map<int, bool> := map[p1 := cf25];
				var v66 := "dpwnr";
				var v67: map<bool, string> := map[!DC10(v0, p3, v0, p0, v65).cf16 := v66];
				var v68: seq<int> := [|v67|];
				globalState.f2, globalState.f2 := v68[safeIndex(0x328, |v68|)], p1;
				var v69: C2 := new C2(if (v0) then p1 else p1);
				v69 := v69;
				var v70 := new C3(cf25, 7);
			case DC16(cf24) =>
				var v71: array<seq<char>> := new seq<char>[23](i2 => ['q', 'q', 'a', 'y', 'f']);
				var v72: array<int> := new int[14](i3 => i3 - -p3);
				var v73 := "joaugkus";
				v72[safeIndex(235, v72.Length)] := |v73|;
				var v74: map<bool, string> := map[p1 != p3 := v73 + v73];
				var v75: array<char> := new char[16];
				v71, v72[safeIndex(235, v72.Length)], v71, v74, v75 := if (v0) then v71 else v71, p1, v71, v74 + v74, v75;
				var v76: set<bool> := {v0};
				p0[safeIndex(979, p0.Length)] := v76 !! v76;
				var v77: array<map<int, int>> := new map<int, int>[2](i4 => map[p3 := p1] + map[v72[safeIndex(235, v72.Length)] := p3]);
				var v78: map<int, int> := map[|"nufq"| := p1];
				v77[safeIndex(296, v77.Length)] := v78;
				var v79: multiset<int> := multiset{p1};
				globalState.f2, p0[safeIndex(979, p0.Length)], v77[safeIndex(296, v77.Length)], v72[safeIndex(235, v72.Length)] := p1 - p3, v79 >= v79, v78, p1;
				var v80: C1 := new C1(DC17(false).cf25);
				var v81: map<bool, C1> := map[p0[safeIndex(979, p0.Length)] := v80];
				p0[safeIndex(979, p0.Length)] := !!(|(v73 + v73)| > |(v81 + v81)|);
				globalState.f2 := -(if (p0[safeIndex(979, p0.Length)]) then -(v72[safeIndex(235, v72.Length)] + |v73|) else -|(v76 + v76)|);
			case DC18(cf26) =>
				var v82: multiset<string> := multiset{"gwvsosvy"};
				var v83 := DC4(v82, p1);
				var v84 := "qcm";
				v83 := v83.(cf6 := v82 * multiset{v84, v84});
				globalState.f2 := safeDivisionInt(p1, p1);
				var v85 := new C3(fm9(v0, globalState) <= v84, p1);
				v0 := v0;
		}
		
		var v86: multiset<bool> := multiset{false};
		var v87: multiset<int> := multiset{|v86|, p3};
		var v88 := "fpbqp";
		var v89: map<string, bool> := map[v88 := fm0(v0, v0, v0, v0, globalState)];
		var v90: set<bool> := {0x17f > (if (p3 in v87) then v87[p3] else p3), v0, if (v88 in v89) then v89[v88] else v0};
		v90 := {v0, v0, false} * (v90 * v90);
		globalState.f2 := if (0x212 in v87) then v87[0x212] else |multiset{v0, v0}|;
		var v91: C1 := new C1(v0);
		v91 := v91;
		var v92: multiset<string> := multiset{v88, v88};
		var v93: set<int> := {if (v88 in v92) then v92[v88] else p1, 0x137, p1, p1};
		v93 := v93;
	}
	
	var v94 := DC5(v0, fm0(v0, v0, false, v0, globalState));
	v0 := !v94.cf9;
	var v95 := "jfta";
	var v96: map<bool, string> := map[v0 := v95];
	var v97: multiset<int> := multiset{p3};
	if (multiset{|v96|} < v97) {
		var v98: array<bool> := new bool[3];
		v98 := v98;
		var v99 := new C1(v0);
		var v100: C3 := new C3(v0, p1);
		var v101: seq<C3> := [v100];
		v1, v0, v1, globalState.f2, globalState.f2 := v1 + v1, v99.f14, v1, p3, p3 + |v101|;
		var v102 := 'v';
		var v103: array<char> := new char[21] [v102, v102, v102, v102, if (fm0(v99.f14, v100.f11, v99.f14, v99.f14, globalState)) then v102 else v102, v102, v102, if (v0) then v102 else v102, v102, v102, v102, v102, if (v0) then v102 else 'x', v102, v102, v102, v102, v102, v102, v102, 'y'];
		v103 := new char[8];
		v99.f14 := p1 == |v1|;
	} else {
		var v104 := new C0();
		globalState.f2 := p3;
		var v105: array<int> := new int[18];
		var v106: seq<array<int>> := [v105, v105];
		var v107: set<bool> := {v0};
		var v108: array<array<int>> := new array<int>[19] [v105, v105, v105, v105, v105, v105, v106[safeIndex(|v107|, |v106|)], v105, v105, v105, v105, v105, v105, v105, v105, v105, v105, v105, v105];
		v108[safeIndex(744, v108.Length)] := v105;
		var v109: multiset<string> := multiset{v95};
		var v110 := DC6(DC4(v109, p1));
		var v111 := DC3('q');
		var v112 := 'a';
		var v113: map<D1, char> := map[v110.(cf10 := v111) := v112];
		globalState.f10, v108[safeIndex(744, v108.Length)], v113 := v105, v105, (fm21(v0, v0, -0x273, v0, globalState) + v113) + map[v110 := 'w'];
		v0 := v0 <== v0;
		v0 := if (v0) then v0 else !v0;
	}
	
	if (((seq(abs(0x315), i5  => ('k'))) + v95) != v95) {
		var v114: map<int, int> := map[p1 := p1];
		var v115: map<int, bool> := map[p3 := v0];
		v0 := fm0(v0, p3 >= -p3, p1 != |v114|, !(if (61 in v115) then v115[61] else v0), globalState);
		v0 := v0;
		var v116: C2 := new C2(p3);
		var v117 := DC19(v116);
		v117 := if (p3 >= v116.f13) then v117 else v117;
		var v118: array<string> := new string[23](i6 => v95[safeIndex(v116.f13, |v95|) := 'o'] + (seq(abs(0x1ad), i7  => ('v'))));
		v118[safeIndex(301, v118.Length)] := v95 + fm9(v0, globalState);
		v118[safeIndex(301, v118.Length)] := seq(abs(-589), i8  => ('v'));
		var v119: set<int> := {v116.f13};
		var v120: seq<set<int>> := [v119];
		v0 := v120[safeIndex(-v116.f13, |v120|)] > v119;
	} else {
		var v121: array<map<bool, int>> := new map<bool, int>[24];
		var v122: map<bool, int> := map[v0 := p1];
		v121[safeIndex(978, v121.Length)] := map[v0 := |(seq(abs(0x6b), i9  => ('n')))|] + v122;
		var v123: seq<int> := [p1];
		var v124: set<bool> := {v0};
		var v125: seq<set<bool>> := [v124, v124];
		var v127: array<int> := new int[17] [p3, p3, p3, p3, |v123|, |v124|, p3, |v97|, |(map v126 : int | (0xaf <= v126) && (v126 < 0x27d) :: (v126 - p1) := (v0))|, p3, |v97|, p1, -p1, p3, p3, p1, p3];
		var v128: seq<array<int>> := [v127, v127];
		var v129 := DC26(p3, v128[safeIndex(p1, |v128|)], v0);
		var v130: array<set<bool>> := new set<bool>[15] [fm15(p3, |v123|, fm0(fm0(v0, v0, v0, !v0, globalState), v0, v0, !v0, globalState), globalState), v124, v125[safeIndex(v123[safeIndex(p3, |v123|)], |v125|)], v124, v124 * v124, v124, if (v0) then v124 else fm15(p1, p1, v0, globalState), v124, v124 + {!v0, v129.cf38}, v124, fm15(p1, p1, v0, globalState), v124, {false}, v124, v124];
		v130[safeIndex(387, v130.Length)] := v124;
		var v131: set<int> := {p1, p3};
		v121[safeIndex(978, v121.Length)], v130[safeIndex(387, v130.Length)] := v122, {v131 !! v131};
		var v132 := DC0(v0);
		v0 := fm0(v0 <== v132.cf0, fm0(v0, v0, v0, v0, globalState), -p3 < p3, fm0(v0, fm0(v0, v0, v0, v0, globalState), v0, v0, globalState), globalState);
		v0 := !true;
		var v133: array<seq<int>> := new seq<int>[24];
		v133[safeIndex(765, v133.Length)] := v123;
		v133[safeIndex(765, v133.Length)] := v123;
		v127[safeIndex(45, v127.Length)] := p3;
		v127[safeIndex(45, v127.Length)] := -p1;
	}
	
}
method Main() {
	var v0 := true;
	var v1: set<bool> := {!!v0};
	var v2: set<set<bool>> := {{v0, !!v0}, v1};
	var v3 := 0x350;
	var v4 := "bk";
	var v5: map<bool, string> := map[v0 := v4];
	var v6: seq<int> := [|v5|, |map[v3 := v3]|, v3];
	var v7: seq<set<bool>> := [v1, v1, v1];
	var v8: array<int> := new int[14];
	var v9: map<int, array<int>> := map[|v7[safeIndex(v3, |v7|)]| := v8];
	var v10: multiset<bool> := multiset{!v0};
	var globalState := new GlobalState(0xf5, false, 0x35e, false, {v1, v1} - v2, if (v0) then map[v3 := v0] else map[v3 := !v0], v6, v6, false, true, if ((if (v0 in v10) then v10[v0] else v3) in v9) then v9[if (v0 in v10) then v10[v0] else v3] else v8);
	globalState.f2 := safeModuloInt(-0x295, -v3) * v3;
	var v11: map<bool, bool> := map[v0 := 758 >= v3];
	var v12: map<bool, int> := map[v0 := v3];
	v11 := v11[v3 <= v3 := v3 == |v12|];
	var v13: array<bool> := new bool[15](i0 => !v0);
	var v14: array<map<bool, bool>> := new map<bool, bool>[14](i1 => v11);
	m0(v13, v3, v14, v3, globalState);
	if (!v0) {
		var v15 := DC0(true);
		var v16: seq<bool> := [fm0(v0, v0, v0, v15.cf0, globalState)];
		m0(v13, |v16| * v3, v14, |v1|, globalState);
		v0 := v0;
		var v17 := new C2(v3);
		v0 := v0;
		v0 := v0;
	} else {
		globalState.f2 := v3;
		v0 := v0;
		var v18 := 't';
		var v19: seq<bool> := [v0, v0];
		match (fm13(v18, multiset(v19), globalState)) {
			case DC4(cf6, cf7) =>
				v0, globalState.f2, v13, v3 := v0, safeDivisionInt(v3, cf7), v13, -cf7;
				globalState.f2 := cf7;
				var v20: C1 := new C1(true);
				v20 := v20;
				var v21, v22, v23 := v20.m4(v12, globalState);
			case DC5(cf8, cf9) =>
				var v24 := DC13(v18);
				var v25: map<bool, D4> := map[cf9 := v24];
				var v26: set<int> := {v3, v3};
				var v27 := DC10(v0, v3, cf9, v13, map[|v26| := cf9]);
				v25 := v25[v27.cf16 := v24];
				v11 := v11[v0 := v0];
				cf9 := cf9;
				var v28 := new C3(cf9, v6[safeIndex(v3, |v6|)] * |v19|);
			case DC3(cf5) =>
				globalState.f2 := v3;
				v0 := v0;
				globalState.f2 := safeDivisionInt(v3, v3);
				v3 := fm2(!v0, globalState);
			case DC6(cf10) =>
				var v29 := DC3(v18);
				var v30: seq<D1> := [fm5(0x1c6, 0x2f, 0x399, v19, globalState), v29];
				var v31: map<string, int> := map["mawdmqm" + v4 := |v30|];
				v31 := (if (true) then v31 else v31) + v31;
				m0(v13, -|multiset(v19)| * v3, v14, v3, globalState);
				var v32: set<int> := {|v19|, v3};
				v8[safeIndex(188, v8.Length)] := |v32|;
				v8[safeIndex(188, v8.Length)] := v3;
				globalState.f2 := v3;
		}
		
		globalState.f2 := v3;
		v10 := v10;
	}
	
	for i2 := v3 * v3 to v3 {
		var v33: C2 := new C2(v3);
		var v34: C1 := new C1(v0);
		var v35: map<int, C1> := map[v3 := v34];
		var v36: C0 := new C0();
		var v37: map<map<int, C1>, C0> := map[v35 := v36];
		var v38: map<C0, C2> := map[if (v35 in v37) then v37[v35] else v36 := v33];
		var v39: seq<bool> := [!v34.f14];
		var v40: array<C2> := new C2[25] [v33, v33, v33, v33, v33, v33, v33, v33, v33, v33, v33, v33, v33, if (v0) then v33 else v33, v33, v33, v33, v33, v33, v33, if (v36 in v38) then v38[v36] else v33, if (v39[safeIndex(v33.f13, |v39|)]) then v33 else v33, v33, v33, v33];
		v40[safeIndex(930, v40.Length)] := v33;
		v0, v40[safeIndex(930, v40.Length)] := !((v3 * -0x327) <= v33.f13), v33;
		v4 := v4 + v4;
		var v42: multiset<int> := multiset{i2, v33.f13};
		var v43: set<int> := {|v42|};
		v34.f14 := v3 !in (map v41 : int | v41 in v43 :: (safeModuloInt(v41, v3)) := (v33.f13))[-v33.f13 := |v43|];
		var v44: multiset<array<int>> := multiset{v8, v8, v8, v8};
		if (v8 in v44) {
			globalState.f2 := -(v6[safeIndex(v33.f13, |v6|)] * i2);
			v4 := (v4 + v4) + v4;
			var v45: multiset<string> := multiset{v4 + "ux", v4, v4 + v4, v4, v4};
			v45 := v45;
			var v46: array<string> := new string[15](i3 => "k");
			var v47 := 'd';
			v46[safeIndex(63, v46.Length)] := ("dfldjy" + "e")[safeIndex(v33.f13, |("dfldjy" + "e")|) := v47];
			var v48 := DC8(|v39|);
			v33.f13, v46[safeIndex(63, v46.Length)], globalState.f2, globalState.f2 := v33.f13 * v48.cf12, "cqdmq", v33.f13, v33.f13 + v33.f13;
			var v49: map<map<bool, int>, map<int, C1>> := map[v12 := map[v3 := v34]];
			v49 := v49;
		} else {
			v11 := v11;
			var v50 := 'h';
			var v51 := DC3(v50);
			var v52: C3 := new C3(!v36.fm7(false, v51, v4, v34.f14, globalState), v3);
			v52 := v52;
			var v53: array<D4> := new D4[1];
			var v54 := DC12(v1);
			v53[safeIndex(230, v53.Length)] := v54;
			v53[safeIndex(230, v53.Length)] := v54;
			var v55 := DC5(v0, v34.f14);
			v0 := (v55.(cf9 := v34.f14)).cf9;
			v34.f14 := v34.f14 <==> v34.f14;
		}
		
	}
	var v56: map<string, int> := map["sharml" := v3];
	var v57 := 'c';
	var v58: C3 := new C3(v0, v3);
	var v59: map<char, C3> := map[v57 := v58];
	var v60: map<int, C3> := map[fm4(globalState) := if (v57 in v59) then v59[v57] else v58];
	var v61: map<map<int, C3>, array<int>> := map[v60 := v8];
	var v62: seq<bool> := [|v56| > |v61|];
	v62 := v62;
	var v63, v64, v65 := v58.m2(v0, v58.f11, globalState);
	var v66: map<D0, bool> := map[DC1(v58.f12, v58.f12, v58.f12) := v58.f11];
	var v67: map<int, int> := map[v58.f12 := v3];
	var v68: map<bool, map<int, int>> := map[v0 := v67];
	var v69 := DC14(v9);
	var v70: multiset<map<D5, int>> := multiset{map[v69 := v58.f12]};
	v8[safeIndex(595, v8.Length)] := |v70|;
	v3, v66, v68, v8[safeIndex(595, v8.Length)] := -safeModuloInt(v58.f12, |(v11 + v11)|), v66, fm14(v4 + v4, globalState), |v4| + v58.f12;
	var i4 := 0;
	while (v58.f12 >= |v10|)
		decreases 100 - i4
	{
		if (i4 >= 100) {
			break;
		}
		
		i4 := i4 + 1;
		var v71: map<int, seq<char>> := map[safeModuloInt(v8[safeIndex(595, v8.Length)], v58.f12) := v4];
		var v72: seq<C3> := [v58];
		var v73 := DC22(v72);
		v71 := v71[|v73.cf33| := [fm6(v58.f12, v0, v65, 'n', globalState)]];
		v0 := (v10 * multiset{true}) == v10;
		var v74 := DC13(if (v58.f11) then v57 else v57);
		v74 := v74;
		var v75 := DC5(v58.f11, v0);
		var v76: seq<D1> := [v75];
		v76 := v76;
	}
	var v77 := new C3(!v58.f11, 0x28b);
	var v78: map<C3, int> := map[v77 := v77.f12];
	v78 := v78[v77 := v77.f12];
	v3 := v58.f12;
	v77.m1(!v63, globalState);
	var v79: array<char> := new char[20];
	var v80: multiset<array<char>> := multiset{v79, v79, v79};
	v80 := v80;
	match (fm13(v57, multiset{v58.f11, v65, !v77.f11}, globalState)) {
		case DC4(cf6, cf7) =>
			var v81 := DC3(v57);
			match (v81) {
				case DC4(cf6, cf7) =>
					var v82: set<int> := {0x340, v77.f12, v77.f12};
					var v83 := new C3(!v0, |v82|);
					var v84: map<int, bool> := map[|"nuwybql"| := v83.f11 || fm0(false, v58.f11, v0, v77.f11, globalState)];
					var v85: multiset<map<bool, bool>> := multiset{v11};
					v84 := v84[-|(v85 - multiset{v11, v11, v11})| := v77.f11];
					v8[safeIndex(595, v8.Length)] := v58.f12;
					v63 := v65;
				case DC5(cf8, cf9) =>
					var v86: multiset<int> := multiset{v58.fm1(v4, v58.f12, globalState), v77.f12, v77.f12, |v4|, |v5|};
					m0(v13, (if (-0x198 in v86) then v86[-0x198] else v77.f12) + v77.f12, v14, v3, globalState);
					v4 := v4 + (seq(abs(0x379), i5  => ('j')));
					var v87, v88, v89 := v58.m2(v3 != cf7, v63, globalState);
					v8[safeIndex(595, v8.Length)] := -v77.f12;
				case DC3(cf5) =>
					v63 := !!((if (false) then v58.f11 else v58.f11) ==> !v58.f11);
					cf7 := cf7 - (10 - |(set v90 : int | (0x189 <= v90) && (v90 < 611) :: (safeDivisionInt(v90, v58.f12)))|);
					v11 := v11[v58.f11 := v63];
					var v91 := new C0();
				case DC6(cf10) =>
					var v92: C1 := new C1(v77.f11);
					var v93: set<C1> := {v92, v92, v92};
					v58.m1(v93 < v93, globalState);
					var v94: array<string> := new string[1];
					v94[safeIndex(269, v94.Length)] := v4;
					v94[safeIndex(269, v94.Length)] := "adpwq";
					var v95: multiset<C3> := multiset{v58};
					v8[safeIndex(595, v8.Length)], v95, v63 := v3, v95, safeModuloInt(v77.f12, v8[safeIndex(595, v8.Length)]) == (v77.f12 * v3);
					v3 := v58.f12;
			}
			
			v13 := v13;
			if (v0) {
				globalState.f2 := if (v58.f11) then -v77.f12 else -0xc9;
				var v96 := DC17(v0);
				var v97: set<int> := {|[!v63, v96.cf25]|};
				v0, globalState.f2 := v97 >= {cf7}, |(if (v77.f11) then seq(abs(-994), i6  => (v57)) else seq(abs(-963), i7  => (v57)))|;
				var v98 := DC1(v58.f12, |v6|, v77.f12);
				var v99 := new C1(v98.cf1 == cf7);
				var v100: array<string> := new string[8](i8 => "n");
				v100[safeIndex(469, v100.Length)] := v4 + v4;
				v100[safeIndex(469, v100.Length)] := v4;
				var v101 := new C1(v77.f11);
			} else {
				v62 := v62 + (v62 + v62);
				m0(v13, v58.f12, v14, v77.f12, globalState);
				var v102: map<map<C3, int>, int> := map[v78 := v77.fm1(v4, v3, globalState)];
				cf7 := if (!v58.f11 <==> fm0(true, v58.f11, v58.f11, v58.f11, globalState)) then v58.f12 else if (v78 in v102) then v102[v78] else v8[safeIndex(595, v8.Length)];
				globalState.f2 := cf7 - v77.f12;
				v13 := if (v77.f11) then v13 else v13;
			}
			
			var v103: multiset<array<int>> := multiset{v8, if (v77.f11) then v8 else v8};
			var v104: seq<C3> := [v77, v58];
			var v105 := DC22(v104);
			var v106: array<D10> := new D10[9] [v105.(cf33 := v104), v105, v105, v105.(cf33 := v104), v105.(cf33 := v104), v105, v105, DC22(v104[safeIndex(v77.f12, |v104|) := v58]), v105];
			var v107: map<array<D10>, multiset<array<int>>> := map[v106 := multiset{v8}];
			v103 := if (v106 in v107) then v107[v106] else v103;
		case DC5(cf8, cf9) =>
			var v108: multiset<string> := multiset{fm9(v77.f11, globalState)};
			var v109 := DC4(v108, v77.f12);
			v109 := v109;
			v13[safeIndex(333, v13.Length)] := v58.f11;
			v13[safeIndex(333, v13.Length)] := false;
			var v110: map<bool, seq<bool>> := map[cf8 := v62];
			var v111: map<seq<bool>, int> := map[if (cf9 in v110) then v110[cf9] else v62[safeIndex(v3, |v62|) := v63] := |v62[safeIndex(-0x207, |v62|) := !!cf8]|];
			var v112: multiset<int> := multiset{|v111|};
			v3 := |{v58.f12 - |v112|}|;
			var v113 := DC5(cf9, cf8);
			if (v113.cf9) {
				v4 := v4 + v4;
				var v114 := new C3(v77.f11, safeDivisionInt(|fm15(v77.f12, v77.f12, v0, globalState)|, v3));
				var v115: seq<array<int>> := [v8, v8];
				var v116: map<char, array<int>> := map[v57 := v8];
				var v117: array<array<int>> := new array<int>[20] [if (83 in v9) then v9[83] else v115[safeIndex(-0x231, |v115|)], v8, if (v77.f11) then v8 else v8, if (v57 in v116) then v116[v57] else v8, v8, v8, v8, v8, v8, v8, v8, v8, v8, v8, v8, v8, v8, v8, v8, v8];
				v117[safeIndex(20, v117.Length)] := v8;
				v117[safeIndex(20, v117.Length)] := v8;
				v114 := v58;
				v8[safeIndex(595, v8.Length)] := -v114.f12 - 500;
			} else {
				var v118: array<bool> := new bool[5];
				m0(v118, safeDivisionInt(v3, v58.f12), v14, v58.f12, globalState);
				var v119, v120, v121 := v58.m2(cf8, v58.f12 <= v77.f12, globalState);
				globalState.f2, v119, v63 := v6[safeIndex(if (v0) then v77.f12 else v3, |v6|)], v62[safeIndex(0x1b1, |v62|)], v57 in v4;
				globalState.f2, v62, v13[safeIndex(333, v13.Length)], globalState.f10, cf9 := if (cf8 in v10) then v10[cf8] else v77.f12, v62 + fm16(0x159, v113.cf9, globalState), true, if (!v65) then v8 else v8, !v63;
				var v122 := new C2(v3);
			}
			
		case DC3(cf5) =>
			var v123: C0 := new C0();
			var v124 := DC16(v123);
			v124 := v124;
			var v125 := new C0();
			m0(v13, v77.f12, v14, v77.f12, globalState);
			v13[safeIndex(775, v13.Length)] := v63;
			v13[safeIndex(775, v13.Length)] := v63;
		case DC6(cf10) =>
			v13[safeIndex(923, v13.Length)] := true;
			v13[safeIndex(923, v13.Length)] := !v63;
			v77.m1(v77.f11, globalState);
			var v126: array<bool> := new bool[9];
			var v127: map<int, bool> := map[v77.f12 := v0];
			var v128 := DC10(v13[safeIndex(923, v13.Length)], v8[safeIndex(595, v8.Length)], true, v126, map[v58.f12 := v77.f11]);
			match ((if (v77.f11) then DC10(v65, v8[safeIndex(595, v8.Length)], v58.f11, v126, v127) else v128).(cf16 := v0, cf14 := v58.f11, cf17 := v126)) {
				case DC10(cf14, cf15, cf16, cf17, cf18) =>
					v8[safeIndex(595, v8.Length)] := v58.fm1(v4, |v67|, globalState);
					v57 := 'w';
					v3 := -v58.f12 * v58.f12;
					v64 := v64[v62 := fm0(v0, v58.f11, v63, v58.f11, globalState)];
				case DC9(cf13) =>
					v79[safeIndex(46, v79.Length)] := if (v58.f11) then v57 else v57;
					v79[safeIndex(46, v79.Length)] := if (v63 !in v62) then v57 else v57;
					var v129: multiset<int> := multiset{v58.f12};
					v129 := v129;
					var v130 := DC20(v77.f11, v13[safeIndex(923, v13.Length)], if (|v11| in v127) then v127[|v11|] else v77.f11, v58.f11);
					var v131: set<string> := {v4};
					var v132: C1 := new C1(v77.f11);
					var v133: map<C1, bool> := map[v132 := v65];
					v65 := fm0(v130.cf28, v63, v131 !! v131, if (v132 in v133) then v133[v132] else false, globalState);
					var v134: array<seq<int>> := new seq<int>[17](i9 => v6);
					var v135: map<array<seq<int>>, bool> := map[v134 := true];
					v13[safeIndex(923, v13.Length)] := if (v134 in v135) then v135[v134] else v77.f11;
				case DC11(cf19) =>
					globalState.f2, v13[safeIndex(923, v13.Length)] := safeModuloInt(fm2(v65, globalState), v77.f12), true <== v0;
					var v136: multiset<int> := multiset{v3, v77.f12, v77.f12};
					globalState.f2 := if (v58.f12 < |v136|) then if (v77.f11 in v10) then v10[v77.f11] else v77.f12 else -v58.f12;
					v65 := v128.cf16;
					v57 := if ({v13[safeIndex(923, v13.Length)], !v63, v63} >= {!v77.f11, v65, v77.f11, v58.f11}) then 'c' else v57;
			}
			
			v67 := v67[v3 := -|v12|];
	}
	
	match (DC0(v0)) {
		case DC1(cf1, cf2, cf3) =>
			match (fm17(v77.f11, v6 + v6, v62[safeIndex(cf2, |v62|)], v0, globalState)) {
				case DC4(cf6, cf7) =>
					var v137 := DC23(v6);
					v6 := v137.cf34 + (if (v58.f11) then v6 else v6)[safeIndex(|v4|, |(if (v58.f11) then v6 else v6)|) := |v5|];
					var v138, v139, v140 := v77.m2(v58.f11, -cf1 <= v77.f12, globalState);
					var v141, v142, v143 := v58.m2(v77.f11 || v63, v58.f12 != v77.f12, globalState);
					var v144: seq<seq<bool>> := [if (v65) then [v58.f11, v77.f11] else v62, v62, fm16(v77.f12, v77.f11, globalState)];
					v144 := v144;
				case DC5(cf8, cf9) =>
					var v145: map<int, bool> := map[cf3 := v58.f12 <= v8[safeIndex(595, v8.Length)]];
					v145 := v145;
					globalState.f2 := v77.f12;
					var v146 := DC1(0x11f, v8[safeIndex(595, v8.Length)], v58.f12);
					v3 := v6[safeIndex(safeDivisionInt(v146.cf2, -(if (v58.f11 in v10) then v10[v58.f11] else v58.f12)), |v6|)];
					v4 := v4;
				case DC3(cf5) =>
					v62 := v62;
					var v147: map<int, array<bool>> := map[|v4| := v13];
					v13 := if (v8[safeIndex(595, v8.Length)] in v147) then v147[v8[safeIndex(595, v8.Length)]] else v13;
					v4 := v4 + v4;
					v63 := safeDivisionInt(v58.f12, -v58.f12) <= v77.f12;
				case DC6(cf10) =>
					var v148: array<array<bool>> := new array<bool>[15] [v13, v13, v13, v13, v13, v13, v13, v13, v13, v13, v13, v13, v13, v13, v13];
					var v149: map<int, array<array<bool>>> := map[0x2a1 := v148];
					var v150: map<bool, array<array<bool>>> := map[v58.f11 := if (v58.f12 in v149) then v149[v58.f12] else v148];
					v150 := v150[true := v148];
					v57 := v57;
					globalState.f2 := -cf3;
					var v151: array<D11> := new D11[29](i10 => if (v58.f11) then DC25() else DC25());
					var v152 := DC25();
					v151[safeIndex(150, v151.Length)] := v152;
					var v153: map<int, seq<int>> := map[safeDivisionInt(cf3, v58.fm1(v4, v8[safeIndex(595, v8.Length)], globalState)) := [v77.f12]];
					v8[safeIndex(595, v8.Length)], v65, v8[safeIndex(595, v8.Length)], v151[safeIndex(150, v151.Length)], v8[safeIndex(595, v8.Length)] := |(if (cf1 in v153) then v153[cf1] else v6)|, false, if (v57 in v4) then cf1 else v77.f12, v152, safeModuloInt(v3, |v4|);
			}
			
			var v154 := new C3(v63, safeModuloInt(|v67|, cf3));
			match (fm17(v58.f11, v6, !(|[v58.f11]| <= v77.f12), v4 <= v4, globalState)) {
				case DC4(cf6, cf7) =>
					v65 := v4 < (seq(abs(0x17b), i11  => (v57)));
					v65 := fm0(v12 != v12, v77.f11, v0, !v154.f11, globalState);
					v63 := if (v58.f11) then v77.f11 else v154.f11;
					var v155 := new C1(!v77.f11 <==> v58.f11);
				case DC5(cf8, cf9) =>
					v4 := v4;
					v77.m1(v58.f12 >= v154.f12, globalState);
					var v156 := new C1(v62 < v62);
					var v157 := new C3(v77.f11, safeModuloInt(cf3, |fm18(cf2, globalState)|));
				case DC3(cf5) =>
					var v158 := DC3(cf5);
					cf5 := v158.cf5;
					var v159: seq<multiset<bool>> := [v10[v58.f11 := abs(v77.f12)]];
					v57 := fm6(v154.f12, v159 != (seq(abs(0x2c8), i12  => (multiset{v77.f11}))), v77.f11, cf5, globalState);
					globalState.f2 := |v6|;
					v13[safeIndex(918, v13.Length)] := true;
					v13[safeIndex(918, v13.Length)] := v63;
				case DC6(cf10) =>
					var v160: map<array<int>, bool> := map[v8 := v58.f11];
					v160 := v160[v8 := fm6(v3, v77.f11, v0, v57, globalState) !in v4];
					cf3 := safeDivisionInt(-cf2, v154.f12) + v154.fm1("sn", cf1, globalState);
					var v161: C0 := new C0();
					var v162 := DC16(v161);
					var v163: map<C0, int> := map[v162.cf24 := |v6|];
					var v164: map<map<C0, int>, int> := map[v163 := -|(seq(abs(376), i13  => (DC25())))|];
					v164 := v164;
					globalState.f2 := v3;
			}
			
			var v165 := DC9(if (true) then v13 else v13);
			match (v165) {
				case DC10(cf14, cf15, cf16, cf17, cf18) =>
					cf16, v57, v63 := v58.f11, v4[safeIndex(cf15, |v4|)], !!false;
					var v166: seq<seq<bool>> := [v62 + v62];
					v166 := v166;
					cf15 := |(v67 + (v67 + v67))|;
					v11 := v11;
				case DC9(cf13) =>
					v3 := |((v4 + "qgqwbmtnb") + ((seq(abs(58), i14  => ('r')))[safeIndex(0x9c, |(seq(abs(58), i14  => ('r')))|) := fm6(v3, v77.f11, v65, 'b', globalState)] + v4))|;
					v8[safeIndex(595, v8.Length)] := v58.f12;
					var v167 := DC1(v58.f12, -v58.fm1(seq(abs(637), i15  => (v57)), v58.f12, globalState), v154.f12);
					v3 := v167.cf3;
					cf13[safeIndex(714, cf13.Length)] := v65 <== true;
					var v168 := DC26(v77.f12, v8, v154.f11);
					var v169 := DC0(v168.cf38);
					cf13[safeIndex(714, cf13.Length)] := v169.cf0;
				case DC11(cf19) =>
					cf2 := |v1| - cf2;
					var v170: array<C1> := new C1[9];
					var v171: C1 := new C1(v77.f11);
					v170[safeIndex(23, v170.Length)] := v171;
					var v172 := DC28(v171);
					v170[safeIndex(23, v170.Length)] := v172.cf40;
					v8[safeIndex(595, v8.Length)] := v8[safeIndex(595, v8.Length)] * v58.f12;
					v13[safeIndex(32, v13.Length)] := v65;
					v13[safeIndex(32, v13.Length)] := v154.f11;
			}
			
		case DC0(cf0) =>
			var v173: map<bool, char> := map[v65 := v57];
			var v174, v175, v176 := v77.m2(|v173| == v8[safeIndex(595, v8.Length)], v58.f11, globalState);
			v58.m1(v65, globalState);
			v174 := v6 != (v6 + v6);
			var v177: map<int, bool> := map[|v11| := v174];
			var v178 := DC10(v174, v58.f12, v63, v13, v177);
			v11 := v11[v178.cf15 >= v8[safeIndex(595, v8.Length)] := v63];
		case DC2(cf4) =>
			var v179: seq<string> := [v4, v4, v4, seq(abs(0x23a), i16  => (v57))];
			if (v179 < (seq(abs(0x18), i17  => (v4)))) {
				var v180: map<int, bool> := map[-v58.f12 := v58.f11];
				v180 := v180[v3 := true];
				v13[safeIndex(578, v13.Length)] := v77.f11;
				v13[safeIndex(578, v13.Length)] := DC10(v77.f11, v77.f12, v65, v13, v180).cf16;
				v77.m1(v58.f12 >= v8[safeIndex(595, v8.Length)], globalState);
				v4 := v4;
				v77.m1(v77.f11, globalState);
			} else {
				var v181: array<array<bool>> := new array<bool>[17];
				v181[safeIndex(12, v181.Length)] := v13;
				v181[safeIndex(12, v181.Length)] := v13;
				var v182 := new C3(v77.f11, v77.f12);
				v8[safeIndex(595, v8.Length)] := fm4(globalState);
				var v183 := DC21(v4);
				globalState.f2 := |v183.cf32|;
				var v184 := DC24(v182.f12);
				v8[safeIndex(595, v8.Length)] := v184.cf35;
			}
			
			var v185 := DC1(v3, v77.f12, v58.f12);
			v3, v65, v67, v3 := v77.f12 * v6[safeIndex(|v62|, |v6|)], 0x2b6 > v185.cf2, map[-v77.f12 := v58.f12], safeModuloInt(safeModuloInt(v58.f12, 0x2b9), v77.f12);
			var v186: map<string, bool> := map[v4 := true];
			var v187 := DC0(if ("ddnbjifo" in v186) then v186["ddnbjifo"] else v65);
			match (v187) {
				case DC1(cf1, cf2, cf3) =>
					var v188: map<map<int, int>, int> := map[fm19(globalState) := v3];
					cf2 := if (v58.f11) then safeModuloInt(v77.f12, cf2) else |v188|;
					v65 := v77.f11;
					v0 := v77.f12 <= cf3;
					v0 := v77.f11;
				case DC0(cf0) =>
					var v189 := new C1(false);
					cf0 := cf0;
					v186 := v186[v4 := true];
					var v190, v191, v192 := v77.m2(if (v65) then v77.f11 else v58.f11, !v63, globalState);
				case DC2(cf4) =>
					var v193, v194, v195 := v77.m2(!v58.f11, v65, globalState);
					var v196 := DC17(v77.f11);
					v65 := fm0(v195 <==> v77.f11, v196.cf25, v193, v63, globalState);
					v195 := v0;
					var v197: array<string> := new string[19](i18 => DC21(v4).cf32);
					v197[safeIndex(130, v197.Length)] := v4;
					v197[safeIndex(130, v197.Length)], v193 := "c" + v4, v195;
			}
			
			var v198: set<int> := {v77.f12, |(seq(abs(-0x9), i19  => (-|v1|)))|, v8[safeIndex(595, v8.Length)]};
			var v199: map<set<int>, bool> := map[v198 := v0];
			var v200: seq<map<set<int>, bool>> := [v199 + v199, v199, v199, map[v198 := false]];
			v3, v199 := if (v58.f11) then -(v58.f12 - v3) else safeDivisionInt(-v58.f12, |v4|), v200[safeIndex(v3, |v200|)];
	}
	
	print v0, "\n";
	print v1 == {true}, "\n";
	print v2 == {{true}}, "\n";
	print v3, "\n";
	print v4, "\n";
	print v5 == map[true := "bk"], "\n";
	print v6 == [1, 1, 848], "\n";
	print v7 == [{true}, {true}, {true}], "\n";
	print v8[7], "\n";
	print |v9|, "\n";
	print v10 == multiset{false}, "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3, "\n";
	print globalState.f4 == {}, "\n";
	print globalState.f5 == map[848 := true], "\n";
	print globalState.f6 == [1, 1, 848], "\n";
	print globalState.f7 == [490], "\n";
	print globalState.f8, "\n";
	print globalState.f9, "\n";
	print globalState.f10[0], "\n";
	print globalState.f10[1], "\n";
	print globalState.f10[2], "\n";
	print globalState.f10[3], "\n";
	print globalState.f10[4], "\n";
	print globalState.f10[5], "\n";
	print globalState.f10[6], "\n";
	print globalState.f10[7], "\n";
	print globalState.f10[8], "\n";
	print globalState.f10[9], "\n";
	print globalState.f10[10], "\n";
	print globalState.f10[11], "\n";
	print globalState.f10[12], "\n";
	print globalState.f10[13], "\n";
	print globalState.f10[14], "\n";
	print globalState.f10[15], "\n";
	print globalState.f10[16], "\n";
	print globalState.f10[17], "\n";
	print globalState.f10[18], "\n";
	print globalState.f10[19], "\n";
	print globalState.f10[20], "\n";
	print globalState.f10[21], "\n";
	print globalState.f10[22], "\n";
	print globalState.f10[23], "\n";
	print globalState.f10[24], "\n";
	print globalState.f10[25], "\n";
	print globalState.f10[26], "\n";
	print globalState.f10[27], "\n";
	print globalState.f10[28], "\n";
	print v11 == map[true := true], "\n";
	print v12 == map[true := 848], "\n";
	print v13[0], "\n";
	print v13[1], "\n";
	print v13[2], "\n";
	print v13[3], "\n";
	print v13[4], "\n";
	print v13[5], "\n";
	print v13[6], "\n";
	print v13[7], "\n";
	print v13[8], "\n";
	print v13[9], "\n";
	print v13[10], "\n";
	print v13[11], "\n";
	print v13[12], "\n";
	print v13[13], "\n";
	print v13[14], "\n";
	print v14[0] == map[true := false], "\n";
	print v14[1] == map[true := false], "\n";
	print v14[2] == map[true := false], "\n";
	print v14[3] == map[true := false], "\n";
	print v14[4] == map[true := false], "\n";
	print v14[5] == map[true := false], "\n";
	print v14[6] == map[true := false], "\n";
	print v14[7] == map[true := false], "\n";
	print v14[8] == map[true := false], "\n";
	print v14[9] == map[true := false], "\n";
	print v14[10] == map[true := false], "\n";
	print v14[11] == map[true := false], "\n";
	print v14[12] == map[true := false], "\n";
	print v14[13] == map[true := false], "\n";
	print v56 == map["sharml" := 848], "\n";
	print v57, "\n";
	print v58.f11, "\n";
	print v58.f12, "\n";
	print |v59|, "\n";
	print |v60|, "\n";
	print |v61|, "\n";
	print v62 == [false], "\n";
	print v63, "\n";
	print v64 == map[[false, true, true] := true], "\n";
	print v65, "\n";
	print v66 == map[DC1(848, 848, 848) := true], "\n";
	print v67 == map[848 := 848], "\n";
	print v68 == map[false := map[424 := 154], true := map[2 := 114]], "\n";
	print |v69.cf22|, "\n";
	print |v70|, "\n";
	print i4, "\n";
	print v77.f11, "\n";
	print v77.f12, "\n";
	print |v78|, "\n";
	print |v80|, "\n";
}