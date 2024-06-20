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
datatype D0 = DC1(cf1: bool) | DC0(cf0: map<int, bool>) | DC2(cf2: D0)
datatype D1 = DC4(cf4: int) | DC3(cf3: array<map<int, int>>)
datatype D2 = DC6(cf6: bool) | DC7(cf7: int, cf8: map<map<bool, bool>, int>, cf9: int) | DC5(cf5: T1)
datatype D3 = DC9(cf11: int) | DC8(cf10: string) | DC10(cf12: D3)
datatype D4 = DC12(cf14: int, cf15: int, cf16: string, cf17: int, cf18: int) | DC13(cf19: int, cf20: array<bool>) | DC14 | DC11(cf13: array<bool>)
datatype D5 = DC15(cf21: array<int>)
datatype D6 = DC16(cf22: seq<int>)
datatype D7 = DC18 | DC17(cf23: seq<bool>) | DC19(cf24: D7)
datatype D8 = DC21(cf26: bool, cf27: bool, cf28: seq<bool>, cf29: bool, cf30: int) | DC20(cf25: set<T1>) | DC22(cf31: D8)
datatype D9 = DC24(cf33: bool, cf34: bool, cf35: bool) | DC23(cf32: map<array<bool>, bool>) | DC25(cf36: D9)
datatype D10 = DC27(cf38: bool, cf39: int, cf40: char, cf41: multiset<bool>) | DC26(cf37: seq<array<int>>) | DC28(cf42: D10)
datatype D11 = DC30 | DC31(cf44: int, cf45: int) | DC29(cf43: T0) | DC32(cf46: D11)
datatype D12 = DC34(cf48: D0, cf49: bool, cf50: bool) | DC35 | DC36(cf51: bool, cf52: int, cf53: bool) | DC33(cf47: set<bool>) | DC37(cf54: D12)
datatype D13 = DC39 | DC40(cf56: bool, cf57: C0, cf58: map<bool, int>) | DC38(cf55: map<bool, int>)
datatype D14 = DC41(cf59: C4)
trait T0 {
	function fm4(p0: D0, p1: D0, p2: multiset<bool>, globalState: GlobalState): bool 
	method m1(p0: int, p1: array<array<int>>, p2: int, p3: int, globalState: GlobalState) returns (r0: bool, r1: int, r2: bool, r3: bool) 
}

trait T1 extends T0 {
}

class GlobalState {
	var f0 : bool
	const f1 : int
	var f2 : bool
	var f3 : multiset<array<int>>
	const f4 : bool
	var f5 : int
	const f6 : multiset<map<int, bool>>
	const f7 : bool
	var f8 : set<bool>
	constructor (f0 : bool, f1 : int, f2 : bool, f3 : multiset<array<int>>, f4 : bool, f5 : int, f6 : multiset<map<int, bool>>, f7 : bool, f8 : set<bool>) {
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

class C0 extends T1 {
	const f10 : bool
	const f11 : array<bool>
	constructor (f10 : bool, f11 : array<bool>) {
		this.f10 := f10;
		this.f11 := f11;
	}
	
	function fm4(p0: D0, p1: D0, p2: multiset<bool>, globalState: GlobalState): bool {
		f10
	}
	method m1(p0: int, p1: array<array<int>>, p2: int, p3: int, globalState: GlobalState) returns (r0: bool, r1: int, r2: bool, r3: bool) {
		globalState.f2 := f10;
		var v0 := "asdhidii";
		var i0 := 0;
		while (v0 <= ("sqahgg" + v0))
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v1 := DC1(f10);
			var v2: map<bool, bool> := map[v1.cf1 <==> f10 := f10];
			v2 := v2[f10 := f10];
			var v3: array<int> := new int[27];
			v3[safeIndex(128, v3.Length)] := p2;
			v3[safeIndex(128, v3.Length)] := p3;
			var v4: seq<bool> := [f10];
			r0, v4, globalState.f2, globalState.f2 := f10, v4 + v4, f10, f10;
			globalState.f5 := safeModuloInt(safeDivisionInt(v3[safeIndex(128, v3.Length)], |v0|), p3);
		}
		var v5: array<map<int, int>> := new map<int, int>[17];
		var v6 := DC3(v5);
		var v7: map<bool, array<map<int, int>>> := map[f10 := v5];
		var v8: seq<array<map<int, int>>> := [v5, v5];
		var v9: array<array<map<int, int>>> := new array<map<int, int>>[21] [v5, v5, v5, v5, v5, v5, v5, v6.cf3, v5, v5, v5, v5, if (f10 in v7) then v7[f10] else v8[safeIndex(517, |v8|)], v5, if (f10) then v5 else v5, v5, v5, v5, v5, v5, v5];
		v9[safeIndex(88, v9.Length)] := v5;
		v9[safeIndex(88, v9.Length)], globalState.f5, r1 := v5, fm0(|v0|, p2, p0, f10 <==> !f10, globalState), p0;
		var v10: set<bool> := {f10};
		var v11: set<int> := {|v10|};
		v11 := fm6(globalState);
		if (!true) {
			globalState.f0 := if (f10) then f10 else !false;
			var v12: map<bool, bool> := map[f10 := !f10];
			var v13: seq<bool> := [(if (f10 in v12) then v12[f10] else f10) <==> f10];
			v13 := [f10] + (v13 + v13);
			var v14: array<int> := new int[7](i1 => safeModuloInt(i1, p3));
			v14 := v14;
			r1 := p0;
			globalState.f0 := f10;
		} else {
			var v15: multiset<bool> := multiset{f10};
			globalState.f2 := (fm7(globalState) + v15[f10 := abs(p3)]) > v15;
			r0 := true;
			var v16: array<int> := new int[13](i2 => i2 - p0);
			v16 := v16;
			var v17: seq<bool> := [false];
			f11[safeIndex(773, f11.Length)] := v17[safeIndex(p2, |v17|) := v17[safeIndex(|v17|, |v17|)]] <= v17;
			f11[safeIndex(773, f11.Length)] := p3 >= p2;
			if (f10) {
				globalState.f5 := fm0(p3, p3, p2, f11[safeIndex(773, f11.Length)], globalState);
				globalState.f0 := fm1(false, globalState);
				var v18: map<int, bool> := map[p3 := f11[safeIndex(773, f11.Length)]];
				var v19: map<int, int> := map[|v18| := safeDivisionInt(p2, 0x2aa)];
				v19 := v19[p0 := if (-|"weoxwqp"| in v19) then v19[-|"weoxwqp"|] else p2];
				globalState.f2 := f10;
				var v20: map<set<bool>, bool> := map[v10 := p3 == p3];
				v20 := v20;
			} else {
				var v21 := 'h';
				v21 := v21;
				var v22: map<int, int> := map[p0 := p3];
				var v23: map<bool, bool> := map[f11[safeIndex(773, f11.Length)] := f10];
				var v24: map<bool, int> := map[f11[safeIndex(773, f11.Length)] := (if (fm0(p2, |v23|, p2, f10, globalState) in v22) then v22[fm0(p2, |v23|, p2, f10, globalState)] else p0) * -p0];
				v24 := v24[f10 := p0 - p0];
				var v25: map<int, bool> := map[p3 := f10];
				v16[safeIndex(317, v16.Length)] := -0x3d9;
				var v26 := DC0(v25);
				var v27: set<char> := {v21};
				v25, v16[safeIndex(317, v16.Length)], v6, globalState.f0 := v26.cf0, 0x2bb, v6.(cf3 := v9[safeIndex(88, v9.Length)]), if (v27 > v27) then f10 else f11[safeIndex(773, f11.Length)];
				var v28: map<D1, int> := map[DC4(p3) := p2];
				var v29 := DC4(v16[safeIndex(317, v16.Length)]);
				var v30: multiset<int> := multiset{0x2a3};
				v28 := v28[v29 := fm0(492, p2, |v30|, false, globalState)];
				globalState.f2 := true;
			}
			
		}
		
		var v31: seq<int> := [p0];
		var v32: map<int, int> := map[v31[safeIndex(p0, |v31|)] := -p0];
		r1 := |v32| - (-p3 - p3);
		r0 := f10;
		var v33: map<int, bool> := map[p0 := true];
		var v34 := DC0(v33);
		var v35: multiset<bool> := multiset{f10, f10};
		var v36: multiset<bool> := multiset{f10, fm4(v34, v34, v35, globalState), f10};
		var v37: map<bool, int> := map[fm4(v34, v34, v35[f10 := abs(p0)], globalState) := p3];
		r1 := if (f10 in v37) then v37[f10] else p3;
		r2 := f10 && (if (p3 in v33) then v33[p3] else fm4(v34, v34, v36, globalState));
		var v38 := 'w';
		var v39: map<bool, char> := map[f10 := v38];
		r3 := fm8(f10, v39, f10, globalState) < v0;
	}
}

class C1 extends T0 {
	constructor () {
	}
	
	function fm4(p0: D0, p1: D0, p2: multiset<bool>, globalState: GlobalState): bool {
		match DC4(-0x2e2) {
			case DC4(cf4) => true
			case DC3(cf3) => true
		}
	}
	method m1(p0: int, p1: array<array<int>>, p2: int, p3: int, globalState: GlobalState) returns (r0: bool, r1: int, r2: bool, r3: bool) {
		var v0 := true;
		var v1: array<bool> := new bool[9];
		var v2 := new C0(v0, v1);
		v2.f11[safeIndex(154, v2.f11.Length)] := true;
		var v3: map<int, bool> := map[p2 := v2.f10];
		var v4 := DC0(v3);
		var v5 := DC0(v4.cf0);
		var v6: multiset<bool> := multiset{fm1(v0, globalState)};
		var v7: map<int, multiset<bool>> := map[p2 := v6];
		v2.f11[safeIndex(154, v2.f11.Length)] := v2.fm4(v4, v5, if (p3 in v7) then v7[p3] else v6, globalState);
		var v8: seq<int> := [p0];
		var v9 := DC16(v8);
		globalState.f5 := fm0(p3, p2, 0x1da, v9.cf22 <= v8, globalState);
		for i0 := p2 to p0 {
			var v10 := DC17([v2.f10]);
			var v11: seq<bool> := [v2.f11[safeIndex(154, v2.f11.Length)], v2.f10, v2.f11[safeIndex(154, v2.f11.Length)], fm4(DC0(fm22(v2.f10, globalState)), v4, v6, globalState), true];
			var v12: set<seq<bool>> := {v11};
			if (v10.cf23 in v12) {
				globalState.f0 := DC6(v2.f10).cf6;
				r2 := p3 >= p0;
				var v13: map<map<bool, bool>, int> := map[map[v0 := v0] := p0];
				var v14 := DC7(|v11|, v13, i0);
				globalState.f5 := v14.cf9;
				var v15: multiset<int> := multiset{0x1ed};
				v15 := fm23(v15[fm0(p0, p3, p2, v2.f10, globalState) := abs(p3)], p3, globalState);
				r1, globalState.f0, v1, v1, r1 := p2, v0, v2.f11, v1, |(v3 + v3)|;
			} else {
				var v16: map<int, array<bool>> := map[p3 := v1];
				v16 := (v16 + v16) + v16;
				r2 := v0;
				var v17: T1 := new C0(v2.f10, v2.f11);
				var v18: set<T1> := {v17};
				var v19 := DC20(v18);
				v1, globalState.f2, globalState.f2 := if (!v2.f10) then v2.f11 else v2.f11, if (!v2.f10) then v19.cf25 >= v18 else v2.f10, v2.f10;
				var v20 := new C0(v2.f11[safeIndex(154, v2.f11.Length)], v2.f11);
				r0 := v2.f10 && true;
			}
			
			var v21 := new C0(v2.f11[safeIndex(154, v2.f11.Length)], v2.f11);
			globalState.f5 := |multiset{v2.f10}|;
			var v22: array<int> := new int[6];
			var v23 := DC15(v22);
			var v24: map<D5, int> := map[if (v2.f10) then v23 else v23 := DC13(p3, v1).cf19];
			v24 := v24[v23 := safeDivisionInt(p0, p3)];
		}
		var v25 := DC6(false);
		var i1 := 0;
		while (v25.cf6)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			var v26 := new C0(v0, v1);
			globalState.f2 := fm0(p0, p0, p2, v2.f10, globalState) != |v3|;
			v3 := v3[p3 := v2.f11[safeIndex(154, v2.f11.Length)]];
			var v27 := 'm';
			v27 := v27;
		}
		var v28: seq<multiset<int>> := [multiset(v8)];
		r3, globalState.f5, v0, r1, r1 := !(v0 <== v0), safeDivisionInt(p3, 0x28c), (v6 + v6) !! v6[v0 := abs(-|multiset(v28)|)], match v4 {
			case DC1(cf1) => 0x20f
			case DC0(cf0) => |{v0}| * p0
			case DC2(cf2) => safeModuloInt(|map[v2.f11[safeIndex(154, v2.f11.Length)] := v2.f11[safeIndex(154, v2.f11.Length)]]|, p0)
		}, safeModuloInt(safeDivisionInt(p0, p2), -safeDivisionInt(p3, p2));
		var v29: map<bool, bool> := map[v2.f11[safeIndex(154, v2.f11.Length)] := false];
		var v30: map<map<bool, bool>, int> := map[v29 := p0];
		var v31: seq<bool> := [v2.f11[safeIndex(154, v2.f11.Length)]];
		var v32 := DC7(p3, v30, |v31|);
		r0 := p3 > v32.cf9;
		r1 := p2;
		r2 := true !in (map[v0 := |v31|] + map[fm4(DC0(v3), v5, v6, globalState) := fm0(p2, p0, p2, v2.f11[safeIndex(154, v2.f11.Length)], globalState)]);
		r3 := v0;
	}
	method m6(p0: bool, p1: int, p2: bool, globalState: GlobalState) returns (r0: int, r1: bool, r2: bool) {
		globalState.f0 := p2;
		var v0 := "igpdhpy";
		var v1: set<string> := {v0};
		v1 := v1;
		if (p2) {
			var v2 := DC6(p2);
			var v3: seq<bool> := [v2.cf6, p2, p0, p2, !p0];
			var v4: multiset<int> := multiset{safeDivisionInt(|v3|, p1)};
			var v5: map<bool, bool> := map[p2 := p0];
			var v6: array<bool> := new bool[11] [p0, (if (p2 in v5) then v5[p2] else p0) && p2, if (p2) then p0 else p2, true, p2, p2, p0, p0, p0, p2, p2];
			var v7: map<bool, seq<bool>> := map[v3[safeIndex(|v4|, |v3|)] := v3];
			v4, globalState.f5, globalState.f0, globalState.f0, v6 := v4, |(if (p0) then v0 else "inuynen")|, !(!p0 && p2), (p1 * p1) >= |v7|, v6;
			var v8: array<int> := new int[10](i0 => i0 * p1);
			v8[safeIndex(572, v8.Length)] := p1;
			var v10: array<map<bool, map<bool, int>>> := new map<bool, map<bool, int>>[6](i1 => map[p0 := map[p0 := |{{p1}, set v9 : int | v9 in [p1] :: (safeModuloInt(v9, 0x2ce))}|]] + map[p0 := map[p0 := p1]]);
			var v11: map<array<bool>, bool> := map[v6 := p0];
			var v12 := DC23(v11);
			v10[safeIndex(659, v10.Length)] := fm24(|v12.cf32|, globalState);
			var v13: set<seq<bool>> := {fm25(if (p2 in v5) then v5[p2] else p2, 'n', p0, globalState), v3, v3};
			var v14: map<bool, int> := map[false := p1];
			var v15: map<bool, map<bool, int>> := map[p0 := v14];
			v8[safeIndex(572, v8.Length)], v3, v10[safeIndex(659, v10.Length)] := |v13|, v3, v15 + v15;
			var v16: seq<int> := [|"bodi"|, p1, p1, -0x1be, v8[safeIndex(572, v8.Length)]];
			var v17: set<bool> := {!fm1(p2, globalState)};
			if ((v8[safeIndex(572, v8.Length)] - v16[safeIndex(p1, |v16|)]) != (p1 - -|v17|)) {
				var v18 := DC13(v8[safeIndex(572, v8.Length)], v6);
				v18 := DC13(v8[safeIndex(572, v8.Length)], v6);
				globalState.f2 := -p1 < (v8[safeIndex(572, v8.Length)] * |v5|);
				var v19: map<int, array<bool>> := map[v8[safeIndex(572, v8.Length)] := v6];
				v19 := v19[p1 := v6];
				var v20: map<int, bool> := map[v8[safeIndex(572, v8.Length)] := !!(p2 && p0)];
				globalState.f2, globalState.f5, globalState.f0, v0 := if (fm0(|v17|, v8[safeIndex(572, v8.Length)], -0x77, p2, globalState) in v20) then v20[fm0(|v17|, v8[safeIndex(572, v8.Length)], -0x77, p2, globalState)] else p0, 0x1b5, (v2.(cf6 := p2)).cf6, v0 + (v0 + v0);
				var v21 := new C0(p0, v6);
			} else {
				v6[safeIndex(481, v6.Length)] := p2;
				var v22 := DC15(v8);
				v6, globalState.f5, v6[safeIndex(481, v6.Length)], v8[safeIndex(572, v8.Length)], v8 := v6, v8[safeIndex(572, v8.Length)], p0, p1, v22.cf21;
				var v23: multiset<bool> := multiset{if (p0) then v2.cf6 else v3[safeIndex(v8[safeIndex(572, v8.Length)], |v3|)]};
				var v24 := DC0(fm22(p2, globalState));
				globalState.f5, globalState.f5 := -|v23|, safeDivisionInt(-(fm26(v8[safeIndex(572, v8.Length)], 0x385, false, fm4(v24, v24, v23, globalState), globalState)).cf9, p1);
				var v26: map<int, multiset<int>> := map[|v0| := v4];
				var v27: map<int, multiset<int>> := map[(if (p1 in v4) then v4[p1] else v8[safeIndex(572, v8.Length)]) * |(map v25 : int | (0x352 <= v25) && (v25 < 991) :: (safeModuloInt(v25, 0x1a0)) := (132))| := (if (|multiset{v6[safeIndex(481, v6.Length)], p0}| in v26) then v26[|multiset{v6[safeIndex(481, v6.Length)], p0}|] else v4) * v4];
				v27 := v27[p1 - p1 := v4];
				v6[safeIndex(481, v6.Length)] := safeModuloInt(p1, p1) > (v8[safeIndex(572, v8.Length)] + p1);
				var v28: array<map<bool, int>> := new map<bool, int>[1](i2 => if (false) then map[p2 := p1] else v14);
				var v29: array<bool> := new bool[8];
				v28, v8[safeIndex(572, v8.Length)], r0, globalState.f5, v29 := v28, fm0(v8[safeIndex(572, v8.Length)], p1, |v0|, !p0, globalState) + |v0|, v8[safeIndex(572, v8.Length)], p1 + |[v16, v16]|, v29;
			}
			
			var v30: map<map<bool, bool>, int> := map[map[false := p2] := v8[safeIndex(572, v8.Length)]];
			var v31 := DC7(p1, v30, v8[safeIndex(572, v8.Length)]);
			var v32: array<D2> := new D2[7] [v31, v31, v31, fm26(v8[safeIndex(572, v8.Length)], v8[safeIndex(572, v8.Length)], p2, !true, globalState), DC7(p1, v30, 587), DC7(v8[safeIndex(572, v8.Length)], v30, |v0|), fm26(p1, p1, !!p0, p2, globalState)];
			v32[safeIndex(368, v32.Length)] := DC7(v8[safeIndex(572, v8.Length)], v30, p1);
			v32[safeIndex(368, v32.Length)] := v31;
			var v33 := DC14();
			var v34: seq<D4> := [v33, DC14(), v33];
			var v35 := 's';
			var v36: map<bool, array<bool>> := map[p2 := v6];
			var v37: set<int> := {v8[safeIndex(572, v8.Length)], |v36| * -v8[safeIndex(572, v8.Length)], 0x15e, p1, -0x392};
			v34, v0, v35, v37, v0 := [v33, v33, DC14(), v33, v33], fm27(globalState), v35, v37 - (v37 - v37), "cps";
		} else {
			var v38: map<bool, bool> := map[p0 := p0];
			var v39 := DC1(p2);
			var v40: map<int, D0> := map[|v38| := v39];
			var v41 := DC2(if (p1 in v40) then v40[p1] else v39);
			v41 := v41;
			var v42: map<int, int> := map[-(0x3ad + p1) := -p1];
			var v43: set<bool> := {p0, p2, p0, p2};
			v42 := v42[|v43| := |{p2}|];
			var v45: set<int> := {p1, 0x248};
			r2 := ((set v44 : int | (0x1fc <= v44) && (v44 < 0x392) :: (v44 + -p1)) + v45) > (v45 + v45);
			r0 := p1;
			var v47: seq<set<int>> := [(set v46 : int | (0x327 <= v46) && (v46 < 0x98) :: (v46 * |[p0]|)) - v45];
			globalState.f5 := |v47[safeIndex(|v0|, |v47|)]|;
		}
		
		globalState.f0 := if (p2) then p0 else p0;
		var v48: array<int> := new int[14];
		v48[safeIndex(294, v48.Length)] := p1;
		v48[safeIndex(294, v48.Length)] := p1;
		r0 := p1;
		var v49: map<bool, int> := map[p2 := v48[safeIndex(294, v48.Length)]];
		var v50: seq<int> := [0x141];
		var v51: map<int, bool> := map[v50[safeIndex(p1, |v50|)] := false];
		r0 := |fm2(p0, true, p2, v49 + map[if (v48[safeIndex(294, v48.Length)] in v51) then v51[v48[safeIndex(294, v48.Length)]] else p0 := p1], globalState)|;
		r1 := p0;
		r2 := !(p2 || (if (p0) then p2 else p0));
	}
	method m7(p0: bool, p1: bool, p2: multiset<bool>, p3: int, globalState: GlobalState) returns (r0: map<bool, int>, r1: bool, r2: array<set<bool>>, r3: int) {
		var v0 := 'u';
		var v1 := DC1(p1);
		var v2 := DC6(v1.cf1);
		var v3: map<int, bool> := map[p3 := false];
		var v4: map<bool, bool> := map[p1 := if (p3 in v3) then v3[p3] else p1];
		var v5: map<map<bool, bool>, int> := map[v4 := p3];
		var v6: multiset<int> := multiset{-p3};
		var v7: array<bool> := new bool[17] [p0, p1, p1, p0, p0, {v0, v0} > {v0}, true, p0, p1, p0 && v2.cf6, v4 in v5, false, v6[p3 := abs(p3)] !! multiset{p3}, p1, p1, p0, p0];
		v7[safeIndex(160, v7.Length)] := p0;
		var v8 := DC9(p3);
		v7[safeIndex(160, v7.Length)] := match v8 {
			case DC9(cf11) => p1
			case DC8(cf10) => p0
			case DC10(cf12) => p1
		};
		var v9 := DC11(v7);
		match (v9) {
			case DC12(cf14, cf15, cf16, cf17, cf18) =>
				cf16 := (cf16 + cf16) + ("taufebrr" + cf16);
				r3 := cf14;
				var v10: set<char> := {v0};
				var v13: map<char, int> := map[v0 := cf14];
				var v15: multiset<char> := multiset{v0, v0};
				var v17: map<char, bool> := map[v0 := p1];
				var v19: array<set<char>> := new set<char>[22] [(set v12 : char | v12 in (set v11 : char | v11 in v10 :: (v11)) :: (v12)) + v10, v10, v10 * fm28(|v6|, cf17, !p1, globalState), {v0, fm29(globalState), v0, v0}, v10, {v0, v0} * v10, if (true) then set v14 : char | v14 in v13 :: (v14) else {v0}, v10, v10, v10 * v10, v10, (set v16 : char | v16 in v15 :: (v16)) - v10, v10, v10, v10, v10, v10, v10, v10 * (set v18 : char | v18 in v17 :: (v18)), v10, {v0}, {v0}];
				v19[safeIndex(275, v19.Length)] := fm28(cf17, |cf16|, v7[safeIndex(160, v7.Length)], globalState);
				v19[safeIndex(275, v19.Length)] := v10;
				r3 := 340;
			case DC13(cf19, cf20) =>
				var v20 := new C0(p2 >= p2, v7);
				var v21: array<int> := new int[26];
				v21 := v21;
				cf19 := cf19;
				v21[safeIndex(371, v21.Length)] := p3;
				v21[safeIndex(371, v21.Length)] := cf19;
			case DC14() =>
				var v22 := "guejok";
				var v23: array<int> := new int[23];
				var v24: map<bool, array<int>> := map[p1 := v23];
				var v25 := DC12(fm0(|"isuftd"|, p3, p3, true, globalState), p3, v22, -|v24[true := v23]|, 612);
				r3 := |v25.cf16|;
				var v26: array<string> := new seq<char>[24](i0 => (seq(abs(-0x5c), i1  => (v0))) + "elvverckl");
				var v27: seq<string> := [v22, v22, v22];
				var v28: seq<bool> := [p0, p0];
				var v29 := DC21(p0, p1, v28, true, p3);
				var v30: set<bool> := {v28[safeIndex(|v29.cf28|, |v28|)]};
				v26[safeIndex(380, v26.Length)] := (seq(abs(0x27e), i2  => ('s'))) + v27[safeIndex(|v30|, |v27|)];
				v26[safeIndex(380, v26.Length)] := v22;
				var v31 := DC0(v3);
				var v32: seq<multiset<bool>> := [multiset([true, p0])];
				var v33 := new C0(fm4(v31, DC0(map[0x2d8 := v7[safeIndex(160, v7.Length)]]), v32[safeIndex(811, |v32|)], globalState), v7);
				var v34 := new C0(v2.cf6, v33.f11);
			case DC11(cf13) =>
				globalState.f5 := p3;
				var v35 := new C0(p0, cf13);
				var v36: array<int> := new int[12](i3 => i3 - p3);
				var v37: seq<array<int>> := [v36];
				var v38 := DC26(v37);
				v37 := v38.cf37 + v37;
				v35 := new C0(p3 >= p3, v35.f11);
		}
		
		r3 := safeDivisionInt(p3, p3);
		var v39: array<int> := new int[20];
		var v40 := DC15(v39);
		match (v40) {
			case DC15(cf21) =>
				var v41: multiset<char> := multiset{v0};
				v41 := (v41 + v41) + v41[v0 := abs(p3)];
				globalState.f5 := p3;
				var v43: seq<int> := [p3];
				var v44: seq<map<int, bool>> := [v3 + (map v42 : int | v42 in v43 :: (safeDivisionInt(v42, -0x3ca)) := (p1)), map[p3 := v7[safeIndex(160, v7.Length)]], v3];
				v44 := (v44 + v44)[safeIndex(fm0(p3, p3, p3, v7[safeIndex(160, v7.Length)], globalState), |(v44 + v44)|) := fm22(p0, globalState)];
				var v45: array<string> := new string[9](i4 => "ipgn" + "uhyhwmg");
				var v46 := "knmngv";
				v45[safeIndex(544, v45.Length)] := v46;
				v45[safeIndex(544, v45.Length)] := fm27(globalState);
		}
		
		var v47: array<char> := new char[14];
		v47[safeIndex(153, v47.Length)] := v0;
		v47[safeIndex(153, v47.Length)] := v0;
		var v48: map<int, int> := map[p3 := -115];
		var v49 := DC0(v3);
		v48, r1, v7 := if (p0) then v48 else v48, if (!(p1 <==> p1) in v4) then v4[!(p1 <==> p1)] else fm4(v49, v49, p2, globalState), v7;
		var v50: map<bool, int> := map[v7[safeIndex(160, v7.Length)] := -467];
		r0 := v50[p0 := p3];
		r1 := v7[safeIndex(160, v7.Length)];
		var v51: array<set<bool>> := new set<bool>[19];
		var v52: seq<array<set<bool>>> := [v51, v51, v51];
		r2 := v52[safeIndex(p3, |v52|)];
		r3 := p3;
	}
}

class C2 extends T0 {
	const f13 : bool
	constructor (f13 : bool) {
		this.f13 := f13;
	}
	
	function fm4(p0: D0, p1: D0, p2: multiset<bool>, globalState: GlobalState): bool {
		f13
	}
	function fm18(p0: D2, p1: int, p2: int, globalState: GlobalState): int {
		0x41
	}
	method m1(p0: int, p1: array<array<int>>, p2: int, p3: int, globalState: GlobalState) returns (r0: bool, r1: int, r2: bool, r3: bool) {
		var v0: seq<bool> := [f13, f13];
		r2 := !((v0 + v0) == (v0 + v0));
		var v1 := 'j';
		var v2: map<int, char> := map[p3 := v1];
		var v3 := "bnmkw";
		var v4 := DC8(v3);
		var v5 := DC10(v4);
		var v7: map<D3, map<int, char>> := map[v5 := map v6 : int | (225 <= v6) && (v6 < 0x13c) :: (v6 * p2) := (v1)];
		var v9: set<int> := {p0, p2};
		var v10: map<bool, map<int, char>> := map[f13 := map v8 : int | v8 in v9 :: (safeDivisionInt(v8, p2)) := (v1)];
		var v11: set<bool> := {f13, f13};
		var v12: map<bool, int> := map[f13 := |v11|];
		var v14: seq<map<int, char>> := [v2];
		var v15: seq<int> := [p0, -|v3|];
		var v17: array<map<int, char>> := new map<int, char>[28] [v2[p2 := v1], v2 + map[0x3b8 := v1], v2[p0 := v1] + v2, v2 + v2, (if (v5 in v7) then v7[v5] else v2) + v2, fm19(globalState), v2, v2 + v2, v2, v2[p2 := 'd'], if (f13) then map[-p2 := v1] else v2, v2, map[p2 := v3[safeIndex(p0, |v3|)]], (map[|map[f13 := f13]| := v1])[|(seq(abs(617), i1  => (v1)))| := v1], map[p2 := v1], v2, v2[p3 := v1], v2, v2, v2, if (!f13 in v10) then v10[!f13] else v2, v2, v2, v2, map[p3 := fm20(-|v12|, f13, p3, globalState)], map v13 : int | v13 in v9 :: (safeDivisionInt(v13, p3)) := (v1), v14[safeIndex(v15[safeIndex(|v3|, |v15|)], |v14|)], (map v16 : int | (0x399 <= v16) && (v16 < -0x382) :: (v16 * p2) := ('s'))[p2 := v1]];
		forall i0 | 0 <= i0 < v17.Length {
			v17[i0] := v2 + v2[0x4a := v1];
		}
		var i2 := 0;
		while (f13)
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			globalState.f5 := -(0x41 * p0);
			r1 := safeModuloInt(|map[p3 := f13]|, p2) - (|"mow"| + -p3);
			if (f13) {
				var v18: array<int> := new int[21](i3 => i3 + p2);
				v18[safeIndex(306, v18.Length)] := p0 + p3;
				v18[safeIndex(306, v18.Length)] := p0;
				var v19: array<bool> := new bool[7] [f13, false, if (f13) then v0[safeIndex(|v3|, |v0|)] else f13, f13, false, f13, f13];
				v19 := v19;
				var v20: T1 := new C0(true, v19);
				var v21: set<D2> := {DC5(v20)};
				var v22 := DC5(v20);
				v21 := if (f13) then {DC5(v20).(cf5 := v20), v22} else {DC5(v20), v22, v22, v22, v22};
				v12 := v12[true := safeDivisionInt(v18[safeIndex(306, v18.Length)], p2)];
				var v23 := DC6(f13);
				v18[safeIndex(306, v18.Length)], globalState.f2, v18 := p3 * fm18(v23, -|fm21(globalState)|, p3, globalState), f13, v18;
			} else {
				var v24 := DC6(f13);
				var v25: array<bool> := new bool[13] [f13, f13, fm1(f13, globalState) || f13, if (f13) then f13 else f13, false, f13, f13, f13, f13, f13, v24.cf6, f13, f13];
				v25[safeIndex(295, v25.Length)] := f13;
				v25[safeIndex(295, v25.Length)] := !(p3 != p3);
				var v26: array<char> := new char[7](i4 => v1);
				v26[safeIndex(167, v26.Length)] := if (f13) then v1 else v1;
				v25[safeIndex(295, v25.Length)], v26[safeIndex(167, v26.Length)] := v25[safeIndex(295, v25.Length)], v1;
				globalState.f2 := f13;
				globalState.f5 := p3;
				var v27 := DC12(p0, p3, v3, p3, |v0|);
				var v28: map<int, D4> := map[p0 := v27];
				v28 := v28[p2 := v27];
			}
			
			var v29: T0 := new C1();
			var v30: array<T0> := new T0[7] [v29, v29, v29, v29, v29, v29, v29];
			var v31 := DC29(v29);
			v30[safeIndex(441, v30.Length)] := v31.cf43;
			v30[safeIndex(441, v30.Length)] := v29;
		}
		var v32: array<int> := new int[10];
		v32 := v32;
		var v33: T0 := new C1();
		v33 := v33;
		var v34: map<int, bool> := map[p3 := f13];
		var v36 := DC6(f13);
		var v38: array<map<int, bool>> := new map<int, bool>[17] [v34, (map v35 : int | (-111 <= v35) && (v35 < 985) :: (v35 - p0) := (false)) + v34, v34 + v34, map[fm18(v36, 0x1fe, p2, globalState) := f13], v34[-p0 := f13], v34 + (map v37 : int | (0x116 <= v37) && (v37 < 0x2ef) :: (v37 * p0) := (f13)), v34, v34, v34, v34[p0 := f13], v34, v34, v34, v34, v34, v34 + v34, v34 + v34];
		var v39: set<char> := {v3[safeIndex(0xed, |v3|)], 'j'};
		var v40: map<set<char>, array<map<int, bool>>> := map[v39 - v39 := v38];
		r3, v3, globalState.f5, v38 := f13, v3, p3, if (v39 in v40) then v40[v39] else v38;
		var v41: map<int, seq<bool>> := map[p2 := v0];
		r0 := 0x2f5 < |v41|;
		r1 := safeDivisionInt(p0, p0);
		r2 := !!false;
		var v42 := DC0(v34);
		var v43: multiset<bool> := multiset{f13};
		var v44 := DC21(f13, f13, v0, v33.fm4(v42, v42, v43, globalState), p0);
		r3 := !v44.cf29;
	}
	method m4(globalState: GlobalState) returns (r0: bool, r1: multiset<int>, r2: map<string, bool>) {
		var v0 := 945;
		for i0 := v0 to v0 {
			v0 := i0;
			var v1 := new C1();
			var v2: array<bool> := new bool[14](i1 => f13);
			var v3: T1 := new C0(f13, v2);
			var v4 := DC5(v3);
			var v5: set<D2> := {v4, v4};
			if (v5 < (v5 + v5)) {
				var v6: map<bool, int> := map[f13 := i0];
				var v7 := DC0(map[|v6| := !f13]);
				globalState.f2 := if (f13) then f13 else fm4(v7, v7, multiset{false}, globalState);
				var v8: set<int> := {i0, v0, i0, v0};
				var v9: seq<bool> := [-fm0(|multiset{i0, v0}|, v0, v0, f13, globalState) >= |v8|, f13, true, f13, f13];
				v9 := v9 + v9;
				v2[safeIndex(634, v2.Length)] := f13;
				v2[safeIndex(634, v2.Length)] := f13;
				v8 := set v10 : int | (676 <= v10) && (v10 < 277) :: (v10 * DC7(v0, map[map[!true := v2[safeIndex(634, v2.Length)]] := i0], v0).cf7);
				var v11: array<int> := new int[26];
				v11[safeIndex(525, v11.Length)] := v0;
				var v12: map<bool, string> := map[v2[safeIndex(634, v2.Length)] := seq(abs(0x2be), i2  => ('f'))];
				var v13 := "d";
				var v14: seq<string> := [v13];
				globalState.f5, v2[safeIndex(634, v2.Length)], v11[safeIndex(525, v11.Length)] := if (f13) then v0 * v0 else safeModuloInt(v0, |(if (false in v12) then v12[false] else "smmjmydnm")|), v14 != v14, i0;
			} else {
				var v15: array<string> := new string[25](i3 => "m");
				v15[safeIndex(815, v15.Length)] := "fyghxaonv";
				var v16 := "trl";
				v15[safeIndex(815, v15.Length)] := v16;
				var v17: seq<bool> := [f13, f13, f13, false, f13];
				var v18: multiset<bool> := multiset{f13, true};
				r0 := multiset(v17) > (if (f13) then v18 else v18);
				var v19 := 'w';
				var v20: map<bool, char> := map[f13 := v19];
				var v21: set<bool> := {f13};
				var v22: map<bool, set<bool>> := map[|v20| == -i0 := v21];
				v22 := v22;
				var v23: array<int> := new int[28];
				var v24 := DC26([v23, v23]);
				v24 := v24;
				v0 := v0;
			}
			
			var v25: array<multiset<bool>> := new multiset<bool>[9];
			var v26: multiset<bool> := multiset{f13};
			var v27: map<bool, int> := map[f13 := i0];
			var v28 := 's';
			v25[safeIndex(337, v25.Length)] := v26 * DC27(true, if (f13 in v27) then v27[f13] else i0, v28, v26).cf41;
			var v29: set<bool> := {!f13};
			var v30 := DC33(v29);
			v25[safeIndex(337, v25.Length)] := multiset{f13, true, true, v30.cf47 <= v29};
		}
		var v31: array<D2> := new D2[6];
		forall i4 | 0 <= i4 < v31.Length {
			v31[i4] := DC7(-v0, map[map[f13 := f13] := v0], if (f13) then |"wlpnjgxrs"| else v0);
		}
		v0, globalState.f5 := v0, safeDivisionInt(0xb1, v0);
		var v32 := new C1();
		for i5 := safeDivisionInt(v0, v0) to v0 {
			var v33: map<int, bool> := map[v0 := f13];
			var v34 := DC0(v33);
			var v35: multiset<bool> := multiset{f13, false};
			var v36: array<bool> := new bool[12] [f13, true, f13, f13, f13, f13, f13, f13, f13, f13, f13, false];
			var v37: seq<array<bool>> := [v36];
			var v38: set<map<int, bool>> := {v33, v33};
			var v39 := "vgk";
			var v40: seq<bool> := [f13, f13];
			var v41 := DC36(f13, 0x4b, f13);
			var v42: C0 := new C0(f13, v36);
			var v43: seq<C0> := [v42];
			var v44: array<bool> := new bool[29] [fm1(f13, globalState), false, !f13, f13, f13, fm4(v34, DC0(v33), v35, globalState), v37 != v37, v38 > v38, f13, !(v39 <= v39), v35 < v35, f13, f13, f13, f13, i5 > |v40[safeIndex(v0, |v40|) := !f13]|, true, !(!!f13 || v41.cf53), f13, fm21(globalState) == v39, f13, fm1(f13, globalState), !f13, !f13, f13, f13, f13, v43 < v43, if (true) then f13 else v42.f10];
			v44 := v42.f11;
			var v45: array<map<int, int>> := new map<int, int>[8];
			v45 := v45;
			v42 := v42;
			globalState.f5 := safeDivisionInt(i5, i5);
		}
		if (f13) {
			var v46: seq<int> := [v0];
			var v47 := DC16(v46);
			v47 := v47;
			var v48: array<int> := new int[17] [v0, v0, 0x9e, v0 * v0, v46[safeIndex(v0, |v46|)], v0, 0x106, v0, v0, v0, -672, v0, |v46| + |(seq(abs(-0x25e), i6  => ('x')))|, v0, v0, v0, v0];
			v48[safeIndex(646, v48.Length)] := -v0;
			v48[safeIndex(646, v48.Length)], globalState.f5 := v0, -v0;
			var v49: map<int, bool> := map[v0 := !!f13];
			var v50: seq<bool> := [f13];
			var v51 := DC0(v49);
			var v52: multiset<bool> := multiset{f13};
			var v53: array<bool> := new bool[22] [f13, [f13] < [fm1(f13, globalState)], f13, if (v0 in v49) then v49[v0] else !f13, f13, v48[safeIndex(646, v48.Length)] > v0, f13, !f13, !f13, f13, f13, f13, f13, f13, !(if (v0 in v49) then v49[v0] else f13), v50 != [f13, f13], false, f13, f13, f13, !f13, if (v32.fm4(v51.(cf0 := map[v48[safeIndex(646, v48.Length)] := f13]), v51, v52, globalState)) then true else f13];
			v53[safeIndex(244, v53.Length)] := f13;
			v53[safeIndex(244, v53.Length)] := !f13;
			var v54 := 'p';
			v54 := v54;
			var v55: array<map<C0, C1>> := new map<C0, C1>[16];
			v55 := v55;
		} else {
			var v56 := new C1();
			globalState.f2 := f13;
			var v57: map<int, int> := map[v0 := -0x137];
			var v58: set<int> := {|v57|};
			if (v58 > (fm30(globalState) * v58)) {
				var v59: array<bool> := new bool[13](i7 => false);
				v59 := v59;
				var v60 := new C1();
				globalState.f0 := f13;
				var v61: C0 := new C0(f13, v59);
				v61 := v61;
				var v62: array<int> := new int[28];
				var v63: multiset<array<int>> := multiset{v62, v62};
				var v64: map<C1, multiset<array<int>>> := map[v32 := v63];
				var v65: seq<int> := [v0];
				var v66 := DC15(v62);
				var v67: map<set<bool>, char> := map[{v61.f10} := 'v'];
				var v68: seq<multiset<array<int>>> := [v63, v63, v63, (v63[v62 := abs(|v67|)])[v62 := abs(v0)], v63];
				var v69 := DC6(f13);
				var v70: map<bool, int> := map[v61.f10 := v0];
				var v71 := 'q';
				var v72: map<int, char> := map[fm18(v69, |v58|, |v70|, globalState) := v71];
				var v73: array<multiset<array<int>>> := new multiset<array<int>>[19] [v63, multiset{v62}, (if (v56 in v64) then v64[v56] else multiset{v62})[v62 := abs(|v65|)], v63[v66.cf21 := abs(0x59)], v63 + v63, v63 - v63, multiset{v62, v62}, v63, v68[safeIndex(v0, |v68|)], v63 - v63[v62 := abs(|v72|)], if (fm1(f13, globalState)) then v63 else v63, v63, v63, v63, v63, multiset{v62, v62, v62}, v63 + v63, v63 * multiset{v62}, v68[safeIndex(v0, |v68|)]];
				v73[safeIndex(592, v73.Length)] := v63;
				v73[safeIndex(592, v73.Length)] := v63;
			} else {
				var v74: set<map<int, int>> := {v57, v57};
				var v75: array<int> := new int[5] [v0, if (f13) then v0 else v0, v0, v0, |v74|];
				v75[safeIndex(915, v75.Length)] := v0;
				var v76: seq<bool> := [f13];
				var v77: seq<seq<bool>> := [v76];
				globalState.f5, v75, v58, v75[safeIndex(915, v75.Length)] := -v0, v75, fm30(globalState), -|(v77 + (seq(abs(-310), i8  => ([f13]))))|;
				var v78: map<bool, bool> := map[f13 := f13];
				var v79 := DC7(v0, map[v78 := v0], -|"fjam"|);
				var v80: multiset<D2> := multiset{if (!f13) then v79 else v79};
				v80 := v80 + v80;
				var v81 := 'p';
				v81 := v81;
				var v82 := new C1();
				globalState.f0 := f13;
			}
			
			var v83: array<bool> := new bool[24] [f13, f13, f13, f13, fm1(f13, globalState), false, f13, !f13, f13, f13, f13, f13, f13, !f13, f13, f13, f13, f13, f13, !f13, f13, f13, f13, f13];
			var v84 := new C0(f13, v83);
			var v85 := "pswtnsaun";
			globalState.f0 := !(false ==> (v85 < (seq(abs(-0x20f), i9  => ('p')))));
		}
		
		var v86: array<bool> := new bool[11](i11 => false);
		var v87 := "gvhxsqh";
		var v88: map<array<bool>, string> := map[v86 := v87];
		var v89: seq<int> := [|v88|];
		r0 := (seq(abs(0x2b3), i10  => (v0))) < v89;
		var v90: multiset<int> := multiset{0x1fb};
		var v91: map<bool, multiset<int>> := map[f13 := v90];
		r1 := if ((|"l"| != |(map[v0 := !f13])[v0 := f13]|) in v91) then v91[|"l"| != |(map[v0 := !f13])[v0 := f13]|] else multiset(seq(abs(0x332), i12  => (v0)));
		var v92: map<string, bool> := map[v87 := f13];
		r2 := v92;
	}
	method m5(p0: seq<array<int>>, p1: int, globalState: GlobalState) returns (r0: int, r1: int, r2: bool) {
		var v0: array<bool> := new bool[10](i0 => f13);
		var v1 := new C0(f13, v0);
		var v2 := DC6(v1.f10);
		var v3: map<int, bool> := map[fm18(v2, p1, p1, globalState) := !f13];
		var v4: set<bool> := {f13, v1.f10};
		v3 := v3[|v4| - p1 := !f13];
		var v5 := 'o';
		v5 := fm20(safeDivisionInt(p1, p1), f13, p1, globalState);
		var i1 := 0;
		while (false)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			var v6: array<int> := new int[15](i2 => safeModuloInt(i2, 0x15d));
			var v7: map<bool, array<int>> := map[f13 := v6];
			var v8: map<int, array<int>> := map[p1 := v6];
			var v9 := DC15(v6);
			var v10: array<array<int>> := new array<int>[18] [v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, if (v1.f10 in v7) then v7[v1.f10] else if (p1 in v8) then v8[p1] else v6, v6, v9.cf21, v6];
			v10, v0 := v10, v0;
			var v11: seq<char> := [v5, 'y', v5];
			v5 := v11[safeIndex(940, |v11|)];
			var v12: seq<bool> := [v1.f10];
			globalState.f0 := if (f13) then v12[safeIndex(p1, |v12|)] else v1.f10 !in multiset{v1.f10, v1.f10};
			globalState.f0 := f13;
		}
		var v13: seq<int> := [-p1, p1, p1];
		var v14: map<int, int> := map[0x16c := -p1];
		var v15: map<int, int> := map[|v14| := p1];
		var i3 := 0;
		while (!(v13 <= ([0x257])[safeIndex(if (0x334 in v15) then v15[0x334] else v13[safeIndex(p1, |v13|)], |[0x257]|) := p1]))
			decreases 100 - i3
		{
			if (i3 >= 100) {
				break;
			}
			
			i3 := i3 + 1;
			var v16 := new C1();
			var v17: multiset<bool> := multiset{v1.f10};
			v17 := v17 + v17;
			globalState.f5 := -safeModuloInt(p1, 538);
			var v18: map<bool, array<int>> := map[false ==> f13 := p0[safeIndex(p1, |p0|)]];
			v18 := v18;
		}
		v15 := v15[p1 := p1];
		var v19: map<bool, int> := map[!v1.f10 := 381];
		r0 := safeDivisionInt(p1 - p1, if (f13 in v19) then v19[f13] else p1);
		r1 := p1;
		var v20 := DC36(v1.f10, p1, false);
		r2 := match v20 {
			case DC34(cf48, cf49, cf50) => v1.f10
			case DC35() => v13 <= v13
			case DC36(cf51, cf52, cf53) => cf52 != |multiset{DC1(v1.f10), DC1(cf51), DC1(cf51)}|
			case DC33(cf47) => v1.f10
			case DC37(cf54) => v1.f10
		};
	}
}

class C3 extends T0, T1 {
	const f12 : int
	constructor (f12 : int) {
		this.f12 := f12;
	}
	
	function fm4(p0: D0, p1: D0, p2: multiset<bool>, globalState: GlobalState): bool {
		f12 <= f12
	}
	function fm11(globalState: GlobalState): map<bool, int> {
		(map[false := |"vd"|] + map[false := f12]) + (map[!!!!!!false := -0xd6] + map[true := f12])
	}
	function fm12(p0: string, p1: int, p2: int, globalState: GlobalState): bool {
		[[f12], [|map[false := f12]|, f12]] < ([[f12, f12, |"cvi"|, -f12], [f12], [f12], [f12, |[false]|]] + (seq(abs(-0x36d), i0  => ([f12, f12]))))
	}
	method m1(p0: int, p1: array<array<int>>, p2: int, p3: int, globalState: GlobalState) returns (r0: bool, r1: int, r2: bool, r3: bool) {
		var v0 := true;
		if (v0 <== v0) {
			var v1: array<bool> := new bool[1](i0 => DC1(v0).cf1);
			v1 := v1;
			if (v0) {
				var v2: map<int, bool> := map[p3 := v0];
				var v3: array<int> := new int[6] [p0, f12, |v2|, p2, -|"fqcgspfpy"|, f12];
				var v4 := DC15(v3);
				var v5: seq<array<int>> := [v3, (v4.(cf21 := v3).(cf21 := v3)).cf21];
				var v6: set<bool> := {v0, v0};
				var v7: set<array<int>> := {v3, v3, v5[safeIndex(|v6|, |v5|)]};
				v7 := v7;
				var v8: multiset<bool> := multiset{fm1(v0, globalState), v0, v0};
				var v9 := "axeend";
				r0, v8 := (v9 + v9) < (v9 + v9), v8;
				var v10 := 'u';
				v10 := v10;
				p1[safeIndex(605, p1.Length)] := v3;
				p1[safeIndex(605, p1.Length)] := v3;
				var v11 := DC4(p3);
				var v12: map<D1, int> := map[v11 := p3];
				var v13: map<bool, map<D1, int>> := map[v0 := v12];
				v13 := map[v0 := map[v11 := -438]] + map[v0 := v12];
			} else {
				var v14: seq<bool> := [v0, !!v0];
				var v15 := new C0(v14[safeIndex(p2, |v14|)], v1);
				var v16: array<int> := new int[22](i1 => i1 * |map[v0 := true]|);
				v16[safeIndex(108, v16.Length)] := p0;
				v16[safeIndex(108, v16.Length)] := f12;
				var v17: set<map<bool, bool>> := {map[v0 := v15.f10]};
				v17 := fm13(globalState);
				var v18 := DC14();
				globalState.f5, v18 := if (v15.f10) then -safeModuloInt(p0, p2) else safeDivisionInt(f12, p2), v18;
				var v19 := "uqxmp";
				v19 := if (v0) then v19 + (seq(abs(-716), i2  => ('p'))) else fm14(v19, p3, globalState) + "fraq";
			}
			
			globalState.f5 := -(p0 * |[v0, v0, v0]|);
			r1 := p2;
			var v20 := DC11(v1);
			v20, globalState.f0 := v20, false;
		} else {
			var v21: map<bool, bool> := map[v0 := v0];
			var v22: set<bool> := {v0, v0};
			v21 := v21[v22 !! v22 := v0];
			var v23: multiset<int> := multiset{p3, f12};
			var v24: seq<int> := [p0];
			var v25: seq<bool> := [v0, v0, v0, v0];
			var v26: map<int, bool> := map[p3 := v0];
			var v27 := DC0(v26);
			var v28: multiset<bool> := multiset{false};
			var v29: seq<bool> := [v0, v25[safeIndex(p2, |v25|)], fm4(v27, DC0(v26), v28, globalState), true];
			globalState.f0, globalState.f5 := v23 >= multiset(v24), safeModuloInt(p3, f12) + |(v29 + v29)|;
			var v30 := DC1(v0);
			match (v30) {
				case DC1(cf1) =>
					globalState.f0 := cf1;
					var v31: map<int, int> := map[0x34 := |v29|];
					var v32: array<bool> := new bool[19] [!v30.cf1, cf1, v0, v0, cf1, v0, true, fm4(v27, v27, v28, globalState), cf1, v0, v0, cf1, cf1, v0, cf1, v0, cf1, fm4(v27, v27, v28, globalState), cf1];
					var v33 := new C0(p2 > (if (p2 in v31) then v31[p2] else -0x3d3), v32);
					var v34 := "gvwujy";
					v34 := v34;
					var v35: array<char> := new char[6];
					v35 := v35;
				case DC0(cf0) =>
					var v36: array<bool> := new bool[29] [v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, false, v0, v0, v0, true, v0, v0, v0, !v0, v0, v0, v0, v0, v0, v25[safeIndex(0x57, |v25|)], v0, v0];
					var v37: map<array<bool>, bool> := map[v36 := v0];
					v37 := v37 + v37;
					var v38 := 'x';
					var v40: map<bool, int> := map[v0 := f12];
					var v41: seq<map<bool, int>> := [v40];
					m0(v0, if (v0) then v38 else v38, |(set v39 : int | (0x16c <= v39) && (v39 < 0x2ef) :: (v39 + 0xa5))|, v40 + v41[safeIndex(f12, |v41|)], globalState);
					var v42: array<map<int, int>> := new map<int, int>[9];
					var v43 := DC3(v42);
					v43 := v43.(cf3 := v42);
					globalState.f0 := if (v0) then v0 else v0;
				case DC2(cf2) =>
					var v44: array<multiset<seq<bool>>> := new multiset<seq<bool>>[23];
					var v45: multiset<seq<bool>> := multiset{v29, v29, v25};
					v44[safeIndex(392, v44.Length)] := v45 * (multiset([v29, [v0]]))[v25 := abs(p0)];
					var v46 := "ofcufvnr";
					var v47: map<int, int> := map[|v46| := p0];
					var v48: seq<map<int, int>> := [v47];
					var v49: array<int> := new int[23];
					p1[safeIndex(666, p1.Length)] := v49;
					var v50: map<bool, int> := map[v0 := p2];
					r1, v44[safeIndex(392, v44.Length)], v48, p1[safeIndex(666, p1.Length)], globalState.f5 := |v50|, v45, [v47], v49, f12;
					var v51: array<seq<int>> := new seq<int>[23] [v24, v24 + v24, if (v0) then v24 else fm15(v0, v0, globalState), v24, v24 + (seq(abs(-0x1ce), i3  => (|(seq(abs(0x1b0), i4  => (p0)))|))), [f12], seq(abs(0x38), i5  => (p3)), v24, (seq(abs(0x2a7), i6  => (p0))) + v24, [f12, p3, fm0(0x1f3, |v47|, p3, v0, globalState)], v24, v24, v24, v24, seq(abs(985), i7  => (p3)), v24, v24, [f12, p2], v24, (seq(abs(0x24e), i8  => (if (false in v50) then v50[false] else p2))) + v24, v24, v24, v24];
					v51[safeIndex(542, v51.Length)] := v24;
					v51[safeIndex(542, v51.Length)] := seq(abs(0x3ab), i9  => (p2));
					var v52: array<bool> := new bool[25](i10 => v0);
					v52 := v52;
					var v53: map<array<int>, bool> := map[p1[safeIndex(666, p1.Length)] := v0];
					v49[safeIndex(582, v49.Length)] := fm0(p2, p3, p0, if (v49 in v53) then v53[v49] else v0, globalState);
					v49[safeIndex(582, v49.Length)] := f12;
			}
			
			var v54 := DC0(v26);
			var v55 := DC2(v54);
			var v56: multiset<D0> := multiset{v55};
			var v57 := "sfcvgsxqa";
			var v58: seq<seq<bool>> := [v25, [v0, v0], v29, v29, v29];
			var v59 := DC6(v0);
			var v60: array<bool> := new bool[18] [v56 !! v56, v0, v0, v0, v24 == [|v29|, p3, |v57|], v0, !v0, v58 < [v25, v29], false, v0, v0, v0, v0, v0, v0, v59.cf6, v0, v0];
			v60[safeIndex(498, v60.Length)] := v0 || v0;
			v60[safeIndex(498, v60.Length)], globalState.f5, globalState.f5, globalState.f5 := v0, p3, --(p2 - (p0 * f12)), if (v59.cf6) then p0 else p2;
			var v61: map<bool, int> := map[v60[safeIndex(498, v60.Length)] := p0];
			var v62: set<int> := {|v21|};
			var v63: map<int, int> := map[-|v62| := -557];
			var v64: array<int> := new int[13] [f12 * p2, f12, |v23|, p0, p0, p2, if (v0 in v61) then v61[v0] else if (|v26| in v63) then v63[|v26|] else p0, p0, 0x2af, 0x1c0, p3, 0x3b8, -p2];
			globalState.f5, globalState.f5, v64, v0 := safeDivisionInt(p0, p0), if (v60[safeIndex(498, v60.Length)]) then 0x327 else p3 + |multiset{v60[safeIndex(498, v60.Length)]}|, v64, !v0;
		}
		
		var v65 := 'u';
		v65 := fm16(globalState);
		var v66: multiset<bool> := multiset{true, false};
		r1 := f12 - |v66|;
		var i11 := 0;
		while (true)
			decreases 100 - i11
		{
			if (i11 >= 100) {
				break;
			}
			
			i11 := i11 + 1;
			if (false) {
				var v67: map<int, bool> := map[f12 := v0];
				var v68: seq<bool> := [false];
				var v69 := "kdak";
				var v70: map<map<int, string>, bool> := map[map[p2 := v69] := false];
				var v72: array<bool> := new bool[28] [true, v0, false, v0, v0, v0, if (|v68[safeIndex(p2, |v68|) := v0]| in v67) then v67[|v68[safeIndex(p2, |v68|) := v0]|] else v0, v0, fm1(v0, globalState), v0, v0, v0, v0, v0, true, v0, v0, v0, v0, v0, v0, !v0, v0, v0, v0, fm1(v0, globalState), if ((map v71 : int | v71 in v67 :: (v71 * p0) := (v69)) in v70) then v70[map v71 : int | v71 in v67 :: (v71 * p0) := (v69)] else v0, v0];
				var v73 := new C0(true, v72);
				var v75 := DC0(map v74 : int | (0x2d5 <= v74) && (v74 < -252) :: (safeDivisionInt(v74, p3)) := (v73.f10));
				v72 := if (fm4(v75.(cf0 := v67), DC0(map[p2 := v0]), multiset(v68), globalState)) then v73.f11 else v72;
				var v76: set<bool> := {v73.f10};
				globalState.f5 := |v76| * p3;
				r0 := v76 >= (v76 - v76);
				var v77 := new C0(v0, v73.f11);
			} else {
				var v78: array<bool> := new bool[4](i12 => v0);
				var v79 := new C0(v0, v78);
				var v80: array<int> := new int[4] [p2, p0, p0, p0];
				var v81: map<bool, array<int>> := map[v0 := v80];
				var v82 := "dypeuvr";
				var v83: set<char> := {v82[safeIndex(|"etmifa"|, |v82|)]};
				var v84: map<bool, int> := map[v0 := |v82[safeIndex(p3, |v82|) := v65]|];
				var v85: array<int> := new int[23] [-0x21b, |v81| * -p0, p2, p0, p0, if (v0) then f12 else p0, fm0(fm0(p0, p0, |fm17(true, p3, v0, p0, globalState)|, v0, globalState), f12, |v83|, v0, globalState), if (v0) then |v84| else -f12, f12, f12, p3, p0, p3, p2 * p2, |(if (v79.f10) then ([0x238])[safeIndex(845, |[0x238]|) := |v82|] else seq(abs(0xe1), i13  => (p2)))|, p0, |v66|, p3, safeModuloInt(|v82|, f12), safeDivisionInt(p2, p2), |multiset{f12}|, safeModuloInt(f12, f12), -487];
				v85 := v85;
				globalState.f5 := f12;
				globalState.f0 := !(!fm12(v82, f12, p0, globalState) ==> v79.f10);
				globalState.f0 := v0;
			}
			
			if (v0) {
				var v86: array<bool> := new bool[11];
				v86 := v86;
				var v87 := "jqxxg";
				var v88: array<C0> := new C0[15];
				var v89, v90 := m3(fm14(v87, p2, globalState), v0, v88, p0, globalState);
				v87 := ("u" + v87) + v87;
				var v91: seq<int> := [|v87|];
				var v92: map<seq<int>, int> := map[v91 := p2];
				v92 := (v92 + v92) + v92;
				v86[safeIndex(366, v86.Length)] := if (v0) then v0 else true;
				var v93: T0 := new C2(false);
				var v94: map<int, T0> := map[f12 := v93];
				var v95: multiset<int> := multiset{v89, -0x16, 0xb5, |v94|, -0x20e};
				var v96: map<int, bool> := map[|(seq(abs(0xc0), i14  => (p3)))| := v0];
				var v97 := DC0(v96);
				var v98 := DC2(v97);
				var v99 := DC0(v96);
				var v100 := DC34(v98, v0, v93.fm4(v99, v99, v66, globalState));
				var v101 := DC34(DC2(v97), v0, !v100.cf50);
				var v102: seq<bool> := [v0, v95 !! multiset(v91), v100.cf49, true];
				v86[safeIndex(366, v86.Length)] := v102[safeIndex(v89, |v102|)];
			} else {
				var v103: map<bool, bool> := map[v0 := v0];
				var v104: seq<map<bool, bool>> := [v103];
				var v105: map<map<bool, bool>, int> := map[v104[safeIndex(p0, |v104|)] := 762];
				var v106 := DC7(|("heyc")[safeIndex(-f12, |"heyc"|) := v65]|, v105[v103 := f12], f12);
				globalState.f5 := v106.cf9;
				var v107: seq<bool> := [!v0];
				var v108 := DC24(true, true, v107[safeIndex(-f12, |v107|)]);
				var v109: map<int, bool> := map[p3 := v0];
				var v110: map<bool, int> := map[v107[safeIndex(0x216, |v107|)] := p0];
				v108, r1 := if ((if (0x278 in v109) then v109[0x278] else if (false in v103) then v103[false] else v0) <==> v0) then DC24(false, v0, v0) else v108, safeDivisionInt(|DC38(v110).cf55|, safeDivisionInt(|v107|, p3));
				var v111: map<int, int> := map[p3 := f12];
				var v112: map<map<int, int>, bool> := map[v111 := v0];
				globalState.f0 := !!!(if (v111 in v112) then v112[v111] else v0);
				var v113: seq<int> := [f12];
				var v114: set<int> := {fm0(|v113|, p0, p0, v0, globalState), p0};
				var v116: array<bool> := new bool[9] [v0, true, v113 == v113, v0, !v0, v0, v0, v0, v114 !! (set v115 : int | (0x2ee <= v115) && (v115 < -0x12) :: (safeModuloInt(v115, p0)))];
				v116[safeIndex(933, v116.Length)] := v0;
				var v117: array<D5> := new D5[16];
				var v118: array<int> := new int[11](i15 => safeModuloInt(i15, p3));
				var v119 := DC15(v118);
				v117[safeIndex(888, v117.Length)] := v119;
				r2, v116[safeIndex(933, v116.Length)], v117[safeIndex(888, v117.Length)] := (v0 <==> !v0) ==> v0, false, v119;
				var v120 := DC9(0x199);
				var v121 := DC10(v120);
				var v122 := "fxm";
				var v123: map<int, D3> := map[p3 := v121.(cf12 := DC8(v122))];
				globalState.f0 := |v123| > p0;
			}
			
			var v124 := DC30();
			var v125 := DC32(v124);
			var v126 := DC32(v125.cf46);
			var v128: array<D4> := new D4[12](i16 => DC12(|multiset([f12, |(set v127 : int | (349 <= v127) && (v127 < 538) :: (safeDivisionInt(v127, p2)))|])|, |[v66]|, "qxnwwoae", -|{true}|, p2));
			var v129: map<D11, array<D4>> := map[if (v0) then v126 else v126 := v128];
			v129 := v129[v126 := v128];
			var v130 := DC30();
			var v131: seq<D11> := [DC30(), v130, v130, v130];
			var v132: multiset<int> := multiset{-p0};
			var v133: seq<seq<int>> := [[p0, |v132|]];
			var v134 := "sxyqh";
			var v135: array<int> := new int[19] [p0, p2, f12 * p3, -|((seq(abs(0x49), i17  => (DC30()))) + v131)|, p2, p3, p2, p3, f12 + p3, safeDivisionInt(|v133|, p3), f12, safeDivisionInt(-p3, |[v0, !v0, !v0, v0, v0]|), -586, f12, p2, safeDivisionInt(p0, f12), |(v134 + v134)|, p0, p2];
			v135[safeIndex(756, v135.Length)] := p2 + |v134|;
			v135[safeIndex(756, v135.Length)] := p2;
		}
		var v136: array<seq<bool>> := new seq<bool>[22];
		forall i18 | 0 <= i18 < v136.Length {
			v136[i18] := [v0, v0] + ([v0, v0, v0] + [false]);
		}
		var v137 := "r";
		v137 := v137 + v137;
		r0 := v0;
		r1 := p3;
		var v138 := DC31(f12, -p2);
		r2 := match v138 {
			case DC30() => v0
			case DC31(cf44, cf45) => v0
			case DC29(cf43) => v0
			case DC32(cf46) => v0
		};
		r3 := !(v0 <== v0);
	}
	method m3(p0: string, p1: bool, p2: array<C0>, p3: int, globalState: GlobalState) returns (r0: int, r1: array<int>) {
		r0 := p3;
		globalState.f2 := !!p1;
		var v0 := DC1(p1);
		var v1 := DC2(v0);
		var v2 := DC34(v1, p1, true);
		var i0 := 0;
		while (v2.cf50)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v3: C2 := new C2(p1);
			var v4: map<bool, C2> := map[p1 := v3];
			var v7 := DC0(map v5 : int | v5 in (map v6 : int | (0x2b7 <= v6) && (v6 < 0xf6) :: (safeModuloInt(v6, -p3)) := (v3.f13)) :: (safeDivisionInt(v5, p3)) := (v3.f13));
			var v8: multiset<bool> := multiset{v3.f13};
			var v9: set<bool> := {p1};
			var v10: array<bool> := new bool[17] [p1 !in v4, v3.f13, true, p1, v3.f13 && p1, p1, p3 >= f12, !v3.f13 <== p1, v3.f13, v3.f13 <==> p1, fm1(v3.f13, globalState), v3.f13, (seq(abs(-224), i1  => ('k'))) == (seq(abs(953), i2  => ('h'))), if (v3.f13) then p1 else v3.f13, fm4(v7, v7, v8, globalState), f12 == |v9|, v3.f13];
			var v11 := "akjkn";
			v10[safeIndex(273, v10.Length)] := v3.f13;
			var v12: multiset<D0> := multiset{v1};
			v10, v11, v10[safeIndex(273, v10.Length)] := v10, seq(abs(0x2f3), i3  => (DC27(p1, f12, 'n', multiset{false, v3.f13, false, p1}).cf40)), v12 !! (v12 * fm31(globalState));
			globalState.f5 := f12;
			var v13 := 's';
			v13 := 'f';
			var v14: set<int> := {f12, f12};
			var v15 := DC13(p3, v10);
			var v16 := new C0(v14 == v14, v15.cf20);
		}
		var v17: multiset<int> := multiset{p3, f12, p3};
		var v18: seq<bool> := [p1, p1, p1, !p1, fm1(!p1, globalState)];
		for i4 := |v17| + |v18| to f12 {
			v18 := v18;
			var v19 := new C2(p1);
			var v20: array<bool> := new bool[21];
			var v21 := new C0(v18[safeIndex(p3, |v18|)], v20);
			var v22: map<bool, int> := map[!p1 := -p3];
			v22 := v22[v19.f13 := f12];
		}
		for i5 := safeDivisionInt(|v17|, p3) to f12 * p3 {
			globalState.f5 := |(p0 + "wjl")|;
			var v23 := DC18();
			var v24 := DC19(v23);
			var v25 := DC19(v24);
			var v26 := DC19(v25.cf24);
			var v27 := DC19(v23);
			var v28: map<D7, bool> := map[v25 := true];
			v28 := v28[v27 := p1];
			var v29: array<bool> := new bool[19];
			var v30 := new C0(p1 <== p1, v29);
			globalState.f2 := !(|p0| >= |v18|);
		}
		var v31: array<string> := new string[23];
		v31 := v31;
		var v32: map<bool, bool> := map[false := p1];
		var v33: map<map<bool, bool>, int> := map[v32 := f12];
		r0 := DC7(f12, v33, p3).cf9;
		var v34: array<int> := new int[2];
		r1 := v34;
	}
}

class C4 extends T1 {
	const f9 : int
	constructor (f9 : int) {
		this.f9 := f9;
	}
	
	function fm4(p0: D0, p1: D0, p2: multiset<bool>, globalState: GlobalState): bool {
		safeModuloInt(f9, f9) > |(multiset{true} * multiset{false, true})|
	}
	function fm5(p0: seq<set<bool>>, p1: char, p2: seq<int>, p3: seq<int>, globalState: GlobalState): int {
		-|{|"urcwm"| == f9, "dwdsvbe" <= (seq(abs(-0x396), i0  => ('p'))), false}|
	}
	method m1(p0: int, p1: array<array<int>>, p2: int, p3: int, globalState: GlobalState) returns (r0: bool, r1: int, r2: bool, r3: bool) {
		r1 := safeDivisionInt(-(p0 * p2), p3);
		var v0: array<bool> := new bool[6];
		var v1 := new C0(true, v0);
		var v2 := new C0(v1.f10 <==> v1.f10, v0);
		v1.f11[safeIndex(204, v1.f11.Length)] := v1.f10;
		var v3 := 'm';
		var v4: array<int> := new int[8](i0 => safeDivisionInt(i0, p0));
		v4[safeIndex(136, v4.Length)] := p0;
		var v5: multiset<bool> := multiset{v1.f10};
		v1.f11[safeIndex(204, v1.f11.Length)], v3, v4[safeIndex(136, v4.Length)] := (v5 * multiset{v1.f10, v1.f10}) != (v5 - v5), v3, safeDivisionInt(p2, |v5|);
		var v6: seq<bool> := [v2.f10, v1.f10, v2.f10];
		var v7: map<bool, int> := map[v6[safeIndex(p2, |v6|)] := v4[safeIndex(136, v4.Length)]];
		v7 := v7[v1.f11[safeIndex(204, v1.f11.Length)] := v4[safeIndex(136, v4.Length)]];
		r2 := v1.f11[safeIndex(204, v1.f11.Length)];
		r0 := v1.f10;
		r1 := p2;
		r2 := fm1(v2.f10, globalState);
		var v8 := DC1(v2.f10);
		r3 := match v8 {
			case DC1(cf1) => {false} !! {true}
			case DC0(cf0) => false ==> v2.f10
			case DC2(cf2) => v1.f10 && v1.f10
		};
	}
	method m2(p0: int, p1: int, p2: seq<T0>, globalState: GlobalState) returns (r0: map<bool, int>, r1: int, r2: bool) {
		var v0 := "vg";
		var v1 := false;
		var v2 := DC0(map[|v0| := v1]);
		match (match v2 {
			case DC1(cf1) => if (true) then DC1(true) else DC1(v1)
			case DC0(cf0) => DC1(v1)
			case DC2(cf2) => DC1(!v1)
		}) {
			case DC1(cf1) =>
				r1 := p0;
				globalState.f5 := --p0;
				var v3: seq<bool> := [false];
				var v4: seq<bool> := [|v3| > p1];
				globalState.f0, v3, globalState.f5 := cf1, v4, |v4|;
				var v5 := 'c';
				v5 := v5;
			case DC0(cf0) =>
				var v6: map<bool, int> := map[v1 := p1];
				var v7: array<bool> := new bool[26] [true, v1, v1, v1, !v1, v1, v1, !v1, v1, v1, false, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, !v1, v1, v1];
				var v8: T1 := new C0(v1, v7);
				var v9: map<map<bool, int>, T1> := map[v6 := v8];
				var v10 := DC0(cf0);
				var v11 := DC2(v10);
				v9, v11 := (v9 + v9)[v6 + v6 := if (v1) then v8 else v8], v11;
				if (v1) {
					cf0 := cf0[p1 := v1];
					var v12: set<bool> := {false};
					var v13: seq<set<bool>> := [v12];
					var v14 := 'f';
					var v15: multiset<T1> := multiset{DC5(v8).cf5, v8, v8};
					globalState.f5 := fm0(f9, |v0|, p1 + fm5(v13, v14, [|v15|], [835], globalState), v1, globalState);
					var v16: multiset<bool> := multiset{v1};
					var v17: map<multiset<bool>, bool> := map[v16 := v1];
					v14 := fm9(-safeDivisionInt(|p2|, |v17|), globalState);
					globalState.f5 := p1;
					cf0 := cf0[-p1 := true];
				} else {
					var v18: array<int> := new int[21](i0 => safeDivisionInt(i0, f9));
					var v19: set<int> := {|v0|};
					v18, globalState.f0, v19 := v18, false, v19;
					var v20: multiset<bool> := multiset{v1};
					globalState.f5 := safeDivisionInt(p1, if (v8.fm4(v2, v2, v20, globalState) in v20) then v20[v8.fm4(v2, v2, v20, globalState)] else -0x9e);
					r1, globalState.f5, r1 := p1 + f9, f9, p1;
					var v21 := 'i';
					var v22: map<bool, char> := map[v1 := v21];
					var v23 := DC8(v0);
					var v24: array<string> := new string[21] [v0, v0, v0, v0 + (seq(abs(0x29b), i1  => ('r'))), v0, v0 + v0, v0, v0 + v0, v0, v0 + v0, v0, v0, v0, v0 + v0, v0, seq(abs(-0x16e), i2  => ('a')), v0, fm8(v1, map[true := v21], v1, globalState) + v0, fm8(v1, v22, v1, globalState), v0, (seq(abs(0x111), i3  => (v21))) + v23.cf10];
					v24[safeIndex(109, v24.Length)] := (("ggh" + v0)[safeIndex(p0, |("ggh" + v0)|) := v21])[safeIndex(p1, |("ggh" + v0)[safeIndex(p0, |("ggh" + v0)|) := v21]|) := v21];
					v24[safeIndex(109, v24.Length)] := (seq(abs(0x33b), i4  => (v21))) + v0;
					var v25: array<char> := new char[28] ['g', v21, v21, 'l', v21, v21, 'v', v21, v21, v21, v21, v21, v21, v21, 'l', v21, v21, v21, v21, v24[safeIndex(109, v24.Length)][safeIndex(-|v20|, |v24[safeIndex(109, v24.Length)]|)], v21, v21, v21, v21, v21, v21, v21, v21];
					v25[safeIndex(483, v25.Length)] := v21;
					v25[safeIndex(838, v25.Length)] := v21;
					v18, globalState.f5, v25[safeIndex(483, v25.Length)], v25[safeIndex(838, v25.Length)] := v18, p0 - p0, v21, v21;
				}
				
				var v26: array<array<int>> := new array<int>[23];
				var v27: array<seq<D3>> := new seq<D3>[6](i5 => [DC10(DC8("ewwnswuai")), DC10(DC9(|cf0|))] + (seq(abs(0x2d4), i6  => (DC10(DC10(DC8(v0)).cf12)))));
				var v28 := DC8(v0);
				var v29: seq<D3> := [DC10(v28)];
				v27[safeIndex(866, v27.Length)] := v29 + v29;
				var v30: seq<bool> := [true || v1];
				globalState.f2, v26, v27[safeIndex(866, v27.Length)], globalState.f5, r2 := v30[safeIndex(p1, |v30|)], v26, v29, f9, f9 < |(v30 + [v1])|;
				globalState.f5 := p0;
			case DC2(cf2) =>
				globalState.f5, v1, globalState.f5 := safeDivisionInt(p0, f9) - f9, v1, f9;
				var v31: array<int> := new int[2];
				v31[safeIndex(555, v31.Length)] := 0x387;
				v31[safeIndex(555, v31.Length)] := safeDivisionInt(p1, p0 + 0x2cd);
				var v32: seq<array<int>> := [v31];
				var v33: array<array<int>> := new array<int>[2] [v32[safeIndex(|"sbtwqs"|, |v32|)], v31];
				var v34: array<array<set<int>>> := new array<set<int>>[21];
				var v35: array<set<int>> := new set<int>[20];
				v34[safeIndex(9, v34.Length)] := v35;
				v33, v34[safeIndex(9, v34.Length)], globalState.f5, v0, globalState.f2 := v33, v35, p1, v0, v1;
				var v36 := DC9(p0);
				var v37: multiset<D3> := multiset{DC10(v36), DC10(v36).(cf12 := v36)};
				var v38 := DC10(v36);
				var v39 := DC1(v1);
				var v40: array<bool> := new bool[15] [v1, v1, v1, v1, v1, v1, true, true, true, v1, v1, v1, v1, true, v39.cf1];
				var v41: T1 := new C0(!(v37 >= v37[v38.(cf12 := v36) := abs(v31[safeIndex(555, v31.Length)])]), v40);
				v41 := v41;
		}
		
		if (false) {
			var v42: map<int, bool> := map[f9 := v1];
			var v43: seq<map<int, bool>> := [v42];
			var v44: multiset<bool> := multiset{v1};
			var v45: set<bool> := {v1, v1};
			var v46 := 'u';
			var v47: map<bool, bool> := map[v1 := v1];
			var v48: map<map<bool, bool>, int> := map[v47 := |v45|];
			var v49 := DC7(p0, v48, |v0[safeIndex(p1, |v0|) := v46]|);
			var v50: map<multiset<bool>, seq<bool>> := map[v44 := fm10(v45, v46, v49, p0, globalState)];
			v1, v1, v43 := v1, v44 in v50, ([v42[f9 := v1], map[p1 := v1]] + (seq(abs(-0x314), i7  => (v42)))) + (seq(abs(971), i8  => (v42)));
			globalState.f0 := v1;
			var v51 := DC6(v1);
			match (v51) {
				case DC6(cf6) =>
					var v52: array<bool> := new bool[15](i9 => cf6);
					var v53: C0 := new C0(cf6, v52);
					var v54: seq<C0> := [v53, v53];
					var v55: array<bool> := new bool[15] [cf6, v1, v1 <== true, cf6, v54 != [v53], v1, cf6, v53.f10, cf6 <== v1, v0 != (seq(abs(0x82), i10  => (v46))), cf6, v1, v1, cf6, v1];
					v55[safeIndex(638, v55.Length)] := !cf6;
					v55[safeIndex(638, v55.Length)] := f9 > p0;
					globalState.f2 := v53.f10;
					var v56: array<int> := new int[26];
					v56 := if (cf6) then v56 else v56;
					var v57: seq<bool> := [v55[safeIndex(638, v55.Length)], v53.f10];
					var v58: map<char, array<bool>> := map[v46 := v53.f11];
					var v59 := DC13(|v57|, if (v46 in v58) then v58[v46] else v55);
					var v60: map<array<bool>, array<bool>> := map[v59.cf20 := v52];
					var v61: T1 := new C0(v53.f10, if (v52 in v60) then v60[v52] else v52);
					var v62: map<int, T1> := map[p1 := v61];
					var v63: seq<set<bool>> := [v45];
					var v64: seq<int> := [|v57|, p1];
					var v65: set<int> := {|v57|};
					var v66: array<map<int, T1>> := new map<int, T1>[16] [v62, v62, v62, v62, v62, v62[f9 := v61], v62, v62, map[p1 := v61], v62, map[|multiset([false, false])| := v61], v62 + v62, map[v49.cf7 := v61], v62[fm5(v63, v46, v64, [-|v65|, p0], globalState) := v61], v62, v62 + v62];
					v66[safeIndex(787, v66.Length)] := v62;
					v66[safeIndex(787, v66.Length)] := v62;
				case DC7(cf7, cf8, cf9) =>
					globalState.f0 := (if (v1) then v1 else v1) && v1;
					var v67: seq<int> := [cf7, |v47|];
					v67 := v67;
					v49 := v49;
					var v68 := DC1(false);
					var v69: seq<bool> := [v1, v1, v1, v1, v68.cf1];
					var v70: array<bool> := new bool[9];
					var v71 := DC13(cf9, v70);
					globalState.f0 := v69[safeIndex(v71.cf19, |v69|)];
				case DC5(cf5) =>
					r2 := true;
					var v72: seq<bool> := [v1];
					var v73 := DC1(v72[safeIndex(p0, |v72|)]);
					var v74: map<D3, D0> := map[DC9(0x262) := v73];
					v74 := v74;
					var v75: array<bool> := new bool[10](i11 => v1);
					var v76: C0 := new C0(true, v75);
					var v77: array<int> := new int[10](i12 => i12 + p1);
					var v78: map<C0, array<int>> := map[v76 := v77];
					var v79: map<bool, map<bool, bool>> := map[v76.f10 := v47];
					var v80: map<int, int> := map[fm0(|v79|, p0, p1, v76.f10, globalState) := p1];
					globalState.f2, r1, globalState.f5 := v1, if (v1) then |(v78 + v78)| else if (v76.f10) then f9 else p0, |(map[f9 := p0] + (v80 + v80))|;
					var v81: multiset<map<int, bool>> := multiset{v42, v42 + v42, v42, v42};
					v76.f11[safeIndex(526, v76.f11.Length)] := v76.f10;
					var v82: multiset<string> := multiset{v0};
					var v83: set<int> := {p1};
					globalState.f5, globalState.f5, globalState.f5, v81, v76.f11[safeIndex(526, v76.f11.Length)] := p0, |v82| * (if (v76.f10) then f9 else fm0(|v83|, p0, p1, v51.cf6, globalState)), |((v72 + [v76.f10]) + v72)|, v81, v1;
			}
			
			if (p1 <= |v45|) {
				v1 := v1;
				var v84 := DC1(v1);
				var v85: array<bool> := new bool[3] [false, v1, v84.cf1];
				var v86 := new C0(!!(f9 >= f9), v85);
				v85[safeIndex(991, v85.Length)] := v1 ==> v86.f10;
				v85[safeIndex(991, v85.Length)] := !v1;
				var v87: map<multiset<bool>, bool> := map[v44 := v85[safeIndex(991, v85.Length)]];
				r1 := |((v87 + v87) + v87)|;
				var v88: array<seq<int>> := new seq<int>[19];
				var v89: T0 := new C2(!v86.f10);
				var v90: seq<bool> := [v85[safeIndex(991, v85.Length)], v86.f10, v85[safeIndex(991, v85.Length)]];
				var v91 := DC17(v90);
				var v92: multiset<D7> := multiset{DC17([v86.f10]), v91};
				var v93: multiset<T0> := multiset{v89};
				v88, globalState.f2, r2, v89, globalState.f8 := v88, v85[safeIndex(991, v85.Length)], v86.f10, v89, {v1, v92 !! v92, v93 !! v93};
			} else {
				globalState.f2, globalState.f5, globalState.f2, globalState.f5 := v1, p1, !!v1, f9;
				var v94: map<int, int> := map[p0 - 755 := -p0];
				var v96: seq<bool> := [v1, !v1];
				v94 := v94[fm0(p0, p1, |(map v95 : int | v95 in map[p0 := true] :: (v95 - -|{p0, p0}|) := (|v0|))|, true, globalState) := |v96|];
				globalState.f0 := true;
				var v97: array<int> := new int[19];
				v97[safeIndex(230, v97.Length)] := p0;
				var v98: array<bool> := new bool[14];
				var v99: C0 := new C0(v1, v98);
				var v100: multiset<C0> := multiset{v99, v99, v99, v99, v99};
				var v101: seq<int> := [|map[|multiset{p1}| := f9]|];
				var v102: seq<string> := [v0, v0];
				r1, globalState.f5, v97[safeIndex(230, v97.Length)] := if (v99 in v100) then v100[v99] else -(if (v99.f10) then v101[safeIndex(|v96|, |v101|)] else f9), -((p1 + p1) * p1), if (fm1(v99.f10, globalState)) then |v102| else p1 * f9;
				var v103: C3 := new C3(-p0);
				var v104: array<C3> := new C3[8] [v103, v103, v103, v103, v103, v103, v103, v103];
				var v105: set<char> := {v46};
				globalState.f2, v1, v104, v0, v97[safeIndex(230, v97.Length)] := (if (v1) then f9 else |v105|) <= -0x15c, v96 != v96, v104, fm27(globalState), v97[safeIndex(230, v97.Length)];
			}
			
			var v106: array<bool> := new bool[1](i13 => v1);
			v106[safeIndex(797, v106.Length)] := if (v1) then v1 else false;
			v106[safeIndex(797, v106.Length)] := v1;
		} else {
			var v107 := 'k';
			var v108: map<int, char> := map[p0 := v107];
			globalState.f5 := |v108| * 0x217;
			r1 := -p0;
			var v109: map<int, bool> := map[f9 := v1];
			var v110 := DC0(v109);
			var v111 := DC2(v110);
			var v112 := DC2(v111);
			match (DC2(if (v1) then v111 else v111)) {
				case DC1(cf1) =>
					var v113: set<char> := {'l', 'm', v107};
					var v114: seq<int> := [p0];
					var v115: seq<bool> := [cf1, false];
					var v116: set<bool> := {v1, cf1};
					var v117: map<set<bool>, bool> := map[v116 := cf1];
					var v118: seq<map<set<bool>, bool>> := [v117[{v1} := true], v117, map[v116 := cf1]];
					var v119: map<int, int> := map[0x379 := p1];
					var v120: seq<string> := [v0];
					var v121: seq<string> := [seq(abs(0x196), i14  => (v107)), v120[safeIndex(p0, |v120|)], v0, "nikx", v0];
					var v122: array<int> := new int[25] [p0, if (if (|v113| in v109) then v109[|v113|] else v1) then fm0(f9, f9, p0, cf1, globalState) else f9, f9, v114[safeIndex(f9, |v114|)], if (v1) then f9 else f9, f9, fm0(|v115|, -0x1e2, p1, true, globalState), f9, p1, p0, p0 * |{f9}|, p0, |v118[safeIndex(-|v114|, |v118|)]|, p0, |(v119 + v119)|, f9, |v115| * |v116|, p1, |(v121[safeIndex(fm0(|v0|, 378, f9, true, globalState), |v121|)])[safeIndex(0x23b, |v121[safeIndex(fm0(|v0|, 378, f9, true, globalState), |v121|)]|) := v107]|, p1, p0, f9, 0x15, p1, p1];
					v122 := v122;
					globalState.f8 := (v116 - v116) + (v116 + v116);
					var v123 := new C3(p0);
					var v124: array<string> := new string[12];
					v124[safeIndex(389, v124.Length)] := v0;
					var v125: array<bool> := new bool[19](i15 => v1);
					v125[safeIndex(478, v125.Length)] := v1;
					var v126: set<int> := {v123.f12};
					v124[safeIndex(389, v124.Length)], r2, v125[safeIndex(478, v125.Length)], globalState.f5, globalState.f2 := (v0 + (seq(abs(-0x11b), i16  => (v107))))[safeIndex(p0, |(v0 + (seq(abs(-0x11b), i16  => (v107))))|) := v107], v1, v1, 0x124, v126 >= (v126 - v126);
				case DC0(cf0) =>
					var v127: array<bool> := new bool[3](i17 => false <== v1);
					var v128 := DC11(v127);
					v127 := (v128.(cf13 := v127)).cf13;
					var v129 := new C1();
					var v130 := new C3(-f9);
					cf0 := cf0[v130.f12 := v1];
				case DC2(cf2) =>
					globalState.f5 := f9;
					globalState.f5 := p1;
					r1 := f9;
					var v131 := new C1();
			}
			
			var v132: array<bool> := new bool[1] [v1 <== true];
			var v133: map<string, bool> := map[v0 := v1];
			var v134: multiset<bool> := multiset{v1, v1};
			v132[safeIndex(189, v132.Length)] := if (("rngu")[safeIndex(|v134|, |"rngu"|) := v107] in v133) then v133[("rngu")[safeIndex(|v134|, |"rngu"|) := v107]] else v1;
			v132[safeIndex(189, v132.Length)] := v1;
			r1 := |v134|;
		}
		
		var v135: C1 := new C1();
		var v136: map<C1, bool> := map[v135 := v1];
		var v137: map<bool, bool> := map[v1 := v1];
		var v138 := 'o';
		var v139 := DC27(v1, p1, v138, multiset{v1});
		var v140: seq<string> := [v0, v0, v0];
		var v141: set<string> := {v140[safeIndex(p0, |v140|)], v0};
		var v142: array<bool> := new bool[27] [!v1, v1, if (!!v1) then if (v135 in v136) then v136[v135] else v1 else false, fm1(v1, globalState), v1 ==> v1, v1 <==> v1, !true, map[p1 := v1] != map[p1 := fm1(v1, globalState)], v1 ==> (if (v1 in v137) then v137[v1] else v1), DC24(v1, v1, v1).cf35, v1, v1, v1, v1, !false, !false, v1, f9 > f9, v135.fm4(v2, v2, v139.cf41, globalState), false, {v0, v0, v0, "q"} >= v141, p1 == f9, v1, v1 ==> v1, v1 && v1, v1, v1];
		v142[safeIndex(902, v142.Length)] := v1 && v1;
		var v143: array<multiset<int>> := new multiset<int>[2](i18 => multiset{|map[|(seq(abs(0x390), i19  => (0x25f)))| := {v1, v1}]|, 0xf1});
		v143[safeIndex(446, v143.Length)] := multiset{p1};
		var v144: multiset<int> := multiset{|multiset{v0, v0, v0}|, f9, f9};
		v142[safeIndex(902, v142.Length)], v143[safeIndex(446, v143.Length)] := v1, v144;
		var v145 := new C2(v1);
		var v146 := DC18();
		v146 := if (v142[safeIndex(902, v142.Length)]) then v146 else v146;
		var v147 := DC31(p0, -f9);
		v142[safeIndex(902, v142.Length)] := v147.cf45 < p1;
		var v148: seq<bool> := [v145.f13];
		var v149: seq<seq<bool>> := [v148, v148];
		var v150: map<seq<D3>, bool> := map[fm32(v149, globalState) := true];
		var v151: map<bool, int> := map[false := |v150|];
		r0 := v151;
		var v152: multiset<bool> := multiset{v142[safeIndex(902, v142.Length)]};
		var v153 := DC19(DC18());
		var v154: map<D7, bool> := map[v153 := v142[safeIndex(902, v142.Length)]];
		var v155: map<seq<string>, int> := map[v140 := f9];
		r1 := safeDivisionInt(p0, if ((if (v153 in v154) then v154[v153] else v142[safeIndex(902, v142.Length)]) in v152) then v152[if (v153 in v154) then v154[v153] else v142[safeIndex(902, v142.Length)]] else -0x294) + (if ([v0] in v155) then v155[[v0]] else p0);
		var v156: set<array<bool>> := {v142};
		r2 := safeModuloInt(|v156|, 0x146) == p1;
	}
}

function fm0(p0: int, p1: int, p2: int, p3: bool, globalState: GlobalState): int {
	if (true in [false, !true]) then |multiset{-0x2a3}| * |{|"g"|, 0x369, 0x33d, |map['b' := true]|}| else -|multiset{0x1af}| * 0x24e
}
function fm1(p0: bool, globalState: GlobalState): bool {
	!false
}
function fm2(p0: bool, p1: bool, p2: bool, p3: map<bool, int>, globalState: GlobalState): map<bool, int> {
	map[true := 0x392] + (map[false := |(seq(abs(-0x340), i0  => ('u')))|] + map[true := -|[true, false]|])
}
function fm3(p0: bool, p1: int, p2: int, p3: int, globalState: GlobalState): D0 {
	DC1({true} >= {true, false})
}
function fm6(globalState: GlobalState): set<int> {
	{|((map v0 : int | (0x2cc <= v0) && (v0 < 849) :: (safeDivisionInt(v0, |"ja"|)) := (-0xf9)) + map[0x7d := |"m"|])|}
}
function fm7(globalState: GlobalState): multiset<bool> {
	multiset{true} + (multiset([!false, true]) - multiset{true})
}
function fm8(p0: bool, p1: map<bool, char>, p2: bool, globalState: GlobalState): string {
	"khrapsoyo" + "l"
}
function fm9(p0: int, globalState: GlobalState): char {
	'r'
}
function fm10(p0: set<bool>, p1: char, p2: D2, p3: int, globalState: GlobalState): seq<bool> {
	[(set v0 : int | (0x19a <= v0) && (v0 < 0x2d5) :: (v0 - |"tsnmoy"|)) <= (set v2 : int | v2 in (map v1 : int | (0x123 <= v1) && (v1 < 256) :: (safeDivisionInt(v1, |{-0x127, -|multiset{-0x251, 360}|}|)) := (true)) :: (v2 * 0xd0)), false, 0x1de !in {-756, |(seq(abs(-0x2c5), i0  => (|[-|map[|"rdw"| := 'm']|, |(seq(abs(497), i1  => (|"smixm"|)))|, 0x152]|)))|}, true]
}
function fm13(globalState: GlobalState): set<map<bool, bool>> {
	{map[true := false], map[true := !false], map[true := false]} + ({map[true := false]} - (set v1 : map<bool, bool> | v1 in (map v0 : map<bool, bool> | v0 in {map[true := false]} :: (v0) := (true)) :: (v1)))
}
function fm14(p0: string, p1: int, globalState: GlobalState): string {
	"dybqgrbwg" + ("wxkb" + (seq(abs(756), i0  => ('p'))))
}
function fm15(p0: bool, p1: bool, globalState: GlobalState): seq<int> {
	DC16(seq(abs(0x12c), i0  => (-0x208))).cf22 + [0x1a5, |{false}|, 0x1d5]
}
function fm16(globalState: GlobalState): char {
	match DC38(map[false := |(map v0 : int | (0x36a <= v0) && (v0 < 0x26c) :: (v0 * 0x2ef) := (-0x33c))|]) {
		case DC39() => if (false) then 'h' else 't'
		case DC40(cf56, cf57, cf58) => 'p'
		case DC38(cf55) => 'u'
	}
}
function fm17(p0: bool, p1: int, p2: bool, p3: int, globalState: GlobalState): set<bool> {
	({false} - {true}) - ({false} + {true})
}
function fm19(globalState: GlobalState): map<int, char> {
	map[if (false) then 822 else -0x2b := 't']
}
function fm20(p0: int, p1: bool, p2: int, globalState: GlobalState): char {
	if ([{!false}, {false}, {true}, {!false}] != [{true, !!true}]) then 'a' else 'p'
}
function fm21(globalState: GlobalState): string {
	((seq(abs(0x317), i0  => ('f'))) + "curdiyset") + "o"
}
function fm22(p0: bool, globalState: GlobalState): map<int, bool> {
	match DC37(DC36(true, -997, false)) {
		case DC34(cf48, cf49, cf50) => map[-|{|(seq(abs(26), i0  => ('b')))|, 192}| := true]
		case DC35() => map[0x372 := false] + map[0x90 := true]
		case DC36(cf51, cf52, cf53) => map[cf52 := cf51]
		case DC33(cf47) => (map v0 : int | (0x3e5 <= v0) && (v0 < 0xa4) :: (v0 - 492) := (false)) + (map v1 : int | (0x2fe <= v1) && (v1 < 0x29e) :: (v1 * |cf47|) := (!!true))
		case DC37(cf54) => map[|"slrpor"| := false] + map[0x2fd := false]
	}
}
function fm23(p0: multiset<int>, p1: int, globalState: GlobalState): multiset<int> {
	multiset{|multiset{true}|, |"qgn"|} - (multiset{---0x29, |map[false := false]|, 0x1ca} - multiset{|(set v0 : map<int, int> | v0 in map[map[-|[!false]| := |{false, false, false, false}|] := true] :: (v0))|})
}
function fm24(p0: int, globalState: GlobalState): map<bool, map<bool, int>> {
	map[!false ==> DC24(!!true, false, true).cf35 := map[true := 855]]
}
function fm25(p0: bool, p1: char, p2: bool, globalState: GlobalState): seq<bool> {
	[true] + ([false, true] + [DC34(DC2(DC1(false)), false, true).cf50, true, !true, false])
}
function fm26(p0: int, p1: int, p2: bool, p3: bool, globalState: GlobalState): D2 {
	match DC10(DC10(DC9(528)).cf12) {
		case DC9(cf11) => DC7(-0x18a, map v0 : map<bool, bool> | v0 in [map[!true := true]] :: (v0) := (|map[cf11 := -|[false, true, true, true]|]|), 0x37a)
		case DC8(cf10) => DC7(0x152, map[map[false := false] := -77], |(seq(abs(373), i0  => ('o')))|)
		case DC10(cf12) => DC7(|(seq(abs(115), i1  => ('k')))|, map v1 : map<bool, bool> | v1 in [map[true := false], map[true := true]] :: (v1) := (0x2fa), 0x6b)
	}
}
function fm27(globalState: GlobalState): string {
	if (false ==> !false) then "j" + (seq(abs(455), i0  => ('o'))) else seq(abs(0x346), i1  => ('p'))
}
function fm28(p0: int, p1: int, p2: bool, globalState: GlobalState): set<char> {
	{'r'}
}
function fm29(globalState: GlobalState): char {
	if (!!(!false && true)) then 'i' else 'n'
}
function fm30(globalState: GlobalState): set<int> {
	match if (false) then DC30() else DC30() {
		case DC30() => {-0x5e, 0x10f}
		case DC31(cf44, cf45) => {cf44} * {-|{false}|}
		case DC29(cf43) => {0x19d} + {-0x1c0, |(map v0 : int | (0x24f <= v0) && (v0 < 0x321) :: (v0 * |"aydffd"|) := (|"ikjyjyifl"|))|}
		case DC32(cf46) => {|"ba"|} * {|{false, false}|, |map[|"ryykijq"| := 't']|}
	}
}
function fm31(globalState: GlobalState): multiset<D0> {
	(multiset{DC2(DC0(map[|multiset{multiset{-0x1db}}| := false])), DC2(DC2(DC0(map[|"jgwt"| := true]))), DC2(DC2(DC0(map v0 : int | v0 in (seq(abs(-0x3d5), i0  => (|map[-|multiset{|(seq(abs(0x118), i1  => ('c')))|}| := !false]|))) :: (v0 + 0x326) := (true))))} * multiset{DC2(DC0(map[-|[true]| := false]))}) * (multiset{DC2(DC2(DC1(true))), DC2(DC1(false)), DC2(DC2(DC1(true)))} * multiset{DC2(DC2(DC2(DC2(DC1(true)))))})
}
function fm32(p0: seq<seq<bool>>, globalState: GlobalState): seq<D3> {
	[DC9(-|[|multiset{!true}|, |map[|(seq(abs(0x2f2), i0  => ('a')))| := false]|]|), DC9(-|multiset{|multiset{false}|, 678, 0x32d, |{0x3cd, 0x3e, --0xdd, |map[true := !true]|, |[true]|}|}|)] + [DC9(-0x197)]
}
function fm33(p0: bool, p1: char, globalState: GlobalState): map<int, D8> {
	(map[|"pptit"| := DC21(false, false, [false, true], true, 249)] + map[|map[-0x276 := -844]| := DC21(true, true, [false], !false, -0x18e)]) + (map[0x7b := DC21(true, true, [false], false, |[0x30]|)] + map[|"ska"| := DC21(true, !!false, [false, true, false, false, true], true, -0x21d)])
}
function fm34(p0: int, p1: int, p2: bool, globalState: GlobalState): map<set<bool>, int> {
	map[{false, false} := |map[false := [false, true]]|] + map[{!true, !false, false, false, !true} := 0x12b]
}
method m0(p0: bool, p1: char, p2: int, p3: map<bool, int>, globalState: GlobalState) {
	var v0: map<bool, multiset<bool>> := map[p0 := multiset{p0, p0, p0, false, p0}];
	var v1 := DC1(p0);
	var v2: multiset<bool> := multiset{p0, v1.cf1};
	var v3 := DC27(p0, 0x399, p1, if (!!true in v0) then v0[!!true] else v2);
	var v4: seq<bool> := [p0];
	var v5: set<int> := {-|v4[safeIndex(p2, |v4|) := !p0]|, p2, p2, 947};
	if (safeModuloInt(-(if (true in p3) then p3[true] else |v3.cf41|), p2) in v5) {
		globalState.f0 := p0;
		var v6: C2 := new C2(p0);
		v6 := v6;
		var v7 := "osmkakap";
		v7 := (seq(abs(0x3e1), i0  => (p1))) + "baefwmqnu";
		var v8: multiset<char> := multiset{p1, p1};
		var v9: array<bool> := new bool[24](i1 => !v6.f13);
		var v10 := new C0(p2 != p2, DC13(|v8|, v9).cf20);
		var v11: array<string> := new string[8];
		v11[safeIndex(776, v11.Length)] := fm14(v7, p2, globalState);
		v11[safeIndex(776, v11.Length)] := v7 + v7;
	} else {
		var v12: map<int, bool> := map[p2 := p0 <== v3.cf38];
		v12 := v12;
		var v13: seq<int> := [p2];
		var v14 := "yxwm";
		v13, globalState.f0, globalState.f0, globalState.f5 := v13, !p0, p0, |v14|;
		var v15: map<int, int> := map[-fm0(|v14|, p2, |v14|, p0, globalState) := |(seq(abs(561), i2  => (p2)))|];
		v15 := v15[|{false, p0}| := |v12|];
		var v16: array<bool> := new bool[23](i3 => p0);
		v16[safeIndex(634, v16.Length)] := p0;
		v16[safeIndex(634, v16.Length)] := p2 > p2;
		v16[safeIndex(634, v16.Length)] := p0 !in v4;
	}
	
	globalState.f2 := p0;
	if (fm1(p0, globalState)) {
		var v17: array<bool> := new bool[14];
		v17[safeIndex(884, v17.Length)] := false;
		v17[safeIndex(884, v17.Length)] := !false;
		if (p0) {
			var v18: array<int> := new int[5](i4 => safeModuloInt(i4, p2));
			var v19: map<int, array<int>> := map[p2 := v18];
			var v20: map<bool, bool> := map[true := p0];
			var v21 := "ikfpgwiid";
			var v22: T1 := new C3(p2);
			var v23: map<T1, int> := map[v22 := p2];
			var v24: C2 := new C2(p0);
			var v25: map<int, C2> := map[p2 := v24];
			var v26: map<char, int> := map['a' := |v25|];
			var v27: map<array<bool>, int> := map[v17 := 0x326];
			var v28: set<bool> := {v17[safeIndex(884, v17.Length)], p0};
			var v29: array<int> := new int[21] [p2, p2, p2 - |v19|, p2, p2, p2, p2, p2, if (p0 in p3) then p3[p0] else -fm0(|v20|, p2, |v21|, p0, globalState), |v23|, -p2, p2, p2, safeDivisionInt(p2, -p2), if (p1 in v26) then v26[p1] else p2, p2, -p2, p2, -|(v27 + v27)|, |v28|, p2];
			var v30: C0 := new C0(v17[safeIndex(884, v17.Length)], v17);
			var v31: set<C0> := {v30};
			var v32: map<int, seq<int>> := map[|multiset{p2, |v31|}| := seq(abs(0x250), i5  => (|"pbcudgtk"|))];
			var v33: seq<int> := [p2];
			v29[safeIndex(420, v29.Length)] := |((if (p2 in v32) then v32[p2] else [552]) + v33)|;
			var v34: array<string> := new seq<char>[25](i6 => seq(abs(209), i7  => (p1)));
			v29[safeIndex(420, v29.Length)], globalState.f2, globalState.f5, v34 := p2, "agjcjgg" < (v21 + (seq(abs(944), i8  => (p1)))), p2, v34;
			globalState.f2, globalState.f5, v29[safeIndex(420, v29.Length)] := p0, fm0(v29[safeIndex(420, v29.Length)], v29[safeIndex(420, v29.Length)], p2, p0, globalState), v29[safeIndex(420, v29.Length)] * v29[safeIndex(420, v29.Length)];
			var v35: C4 := new C4(v29[safeIndex(420, v29.Length)]);
			var v36: map<int, C4> := map[p2 := v35];
			var v37: seq<C4> := [v35, if (|v28| in v36) then v36[|v28|] else v35, v35];
			v37 := v37 + v37[safeIndex(p2, |v37|) := v35];
			var v38 := 's';
			v38 := if (v24.f13) then p1 else p1;
			globalState.f0 := true !in p3;
		} else {
			var v39: array<C2> := new C2[11];
			var v40: seq<int> := [p2];
			var v41: C2 := new C2(v4[safeIndex(v40[safeIndex(p2, |v40|)], |v4|)]);
			v39[safeIndex(618, v39.Length)] := v41;
			v39[safeIndex(618, v39.Length)] := new C2(v41.f13);
			globalState.f2 := p0;
			var v42: T0 := new C2(p0);
			var v43: map<T0, int> := map[v42 := p2];
			var v44: map<int, bool> := map[|v43| := v17[safeIndex(884, v17.Length)]];
			var v45 := DC0(v44);
			v2 := multiset{!(false && v17[safeIndex(884, v17.Length)]), !v41.fm4(v45, v45, v2, globalState)};
			var v46 := "colybgq";
			var v47: map<string, seq<int>> := map[v46 := v40 + [0x28c]];
			var v48: seq<seq<int>> := [seq(abs(0x83), i9  => (p2))];
			v47 := v47["tqy" := v48[safeIndex(-0x390, |v48|)]];
			var v49 := new C3(p2);
		}
		
		var v50 := new C1();
		globalState.f0 := v2 <= v2;
		var v51 := DC18();
		var v52: seq<D7> := [v51];
		var v53: seq<seq<D7>> := [v52];
		var v54: map<bool, seq<seq<D7>>> := map[v17[safeIndex(884, v17.Length)] := v53];
		v17[safeIndex(884, v17.Length)] := |(if (p0 in v54) then v54[p0] else v53)| <= -(if (true) then -279 else p2);
	} else {
		var v55: map<int, bool> := map[p2 := p0];
		var v56: C4 := new C4(|v55|);
		var v57: array<C4> := new C4[15] [v56, v56, v56, if (!p0) then v56 else v56, v56, v56, v56, v56, v56, v56, v56, v56, v56, if (p0) then v56 else v56, v56];
		v57[safeIndex(755, v57.Length)] := v56;
		var v58 := "tlsst";
		var v59: map<bool, int> := map[p0 := |v58|];
		var v60: array<int> := new int[13];
		v60[safeIndex(504, v60.Length)] := -v56.f9;
		var v61: array<bool> := new bool[23](i10 => p0);
		var v62 := DC13(p2, v61);
		var v63 := DC41(v56);
		var v64 := DC41(v63.cf59);
		var v65: map<bool, char> := map[p0 := 'y'];
		var v66 := DC6(p0);
		var v67: C2 := new C2(p0);
		var v68: set<C2> := {v67, v67, v67};
		v57[safeIndex(755, v57.Length)], v59, v60[safeIndex(504, v60.Length)], v62, globalState.f2 := v63.cf59, v59, |fm8(p0, v65, v66.cf6, globalState)|, v62, (v68 !! v68) && v67.f13;
		var v69 := new C0(p0, v61);
		var v70: T0 := new C1();
		var v71: set<T0> := {v70, v70, v70, v70};
		var v72: seq<int> := [|v71|, p2];
		v60[safeIndex(504, v60.Length)] := -safeDivisionInt(safeModuloInt(v56.f9, p2), |v72|);
		v55 := v55;
		var v74: seq<seq<bool>> := [[v69.f10, v67.f13]];
		var v75: map<map<seq<bool>, int>, array<int>> := map[map v73 : seq<bool> | v73 in v74 :: (v73) := (|v72|) := v60];
		var v76: map<seq<bool>, int> := map[v4 := v60[safeIndex(504, v60.Length)]];
		v60 := if ((v76 + v76) in v75) then v75[v76 + v76] else v60;
	}
	
	var v77: array<int> := new int[20](i11 => safeDivisionInt(i11, |{p0}|));
	var v78: map<bool, array<int>> := map[p0 := v77];
	globalState.f5 := safeModuloInt(p2 * |v78|, fm0(p2, fm0(0x332, p2, |"dmpivcli"|, p0, globalState), p2, p0, globalState));
	var v79 := new C4(p2 - p2);
	globalState.f0 := false;
}
method Main() {
	var v0 := 292;
	var v1: seq<int> := [v0, v0];
	var v2 := false;
	var v3: array<int> := new int[3] [v0, v1[safeIndex(|map[v2 := v2]|, |v1|)], v0];
	var v4: multiset<array<int>> := multiset{v3};
	var v5: seq<bool> := [true];
	var v6: map<int, bool> := map[|v5| := v2];
	var v7: map<int, multiset<map<int, bool>>> := map[v0 := multiset{v6}];
	var v9 := DC0(map v8 : int | (388 <= v8) && (v8 < 693) :: (safeDivisionInt(v8, v0)) := (v2));
	var v10: multiset<map<int, bool>> := multiset{v6, v9.cf0};
	var v11: set<bool> := {v2, v2};
	var globalState := new GlobalState(true, -616, false, v4, true, 0x297, (if (v0 in v7) then v7[v0] else v10) + v10, false, v11);
	var v12 := "aiwqmuo";
	globalState.f2 := safeModuloInt(v0, -v0) < fm0(|v12|, v0, v0, v2, globalState);
	var i0 := 0;
	while (v2)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		globalState.f0 := fm1(v0 <= 0x137, globalState);
		v3[safeIndex(871, v3.Length)] := v0;
		v3[safeIndex(871, v3.Length)], v0 := v0, 618;
		var v13 := 'y';
		var v14: map<bool, int> := map[v2 := |v12|];
		m0(!(v2 <==> v2), v13, -(if (v2 in v14) then v14[v2] else v3[safeIndex(871, v3.Length)]), v14, globalState);
		var v15: map<int, seq<int>> := map[-v0 := v1];
		v15 := v15;
	}
	var v17: map<bool, int> := map[v2 := |v12|];
	m0(true, v12[safeIndex(fm0(|(set v16 : char | v16 in v12 :: (v16))|, -v0, v0, v2, globalState), |v12|)], fm0(v0, v0, if (v2 in v17) then v17[v2] else |v12|, v2, globalState) - v0, fm2(v2, v2, v2, v17, globalState), globalState);
	var v18 := 'n';
	m0(v2, v18, fm0(v0, v0, v0, v2, globalState), map[v2 := fm0(v0, v0, |v12|, true, globalState)], globalState);
	for i1 := v0 to -0x2f0 {
		v0 := i1;
		globalState.f0 := v2;
		m0(v2, v18, i1, v17, globalState);
		var v19: multiset<int> := multiset{i1, v0, i1};
		v19 := v19;
	}
	var v20 := DC1(v2);
	var v21: map<bool, D0> := map[v2 := v20];
	v21 := v21[v2 := DC1(v2)];
	match (v20) {
		case DC1(cf1) =>
			var v22: map<string, int> := map["qrvcjj" := v0];
			var v23: multiset<int> := multiset{|v22|};
			var v25 := DC0(v6);
			var v26 := DC2(v25);
			var v27: map<D0, string> := map[v26 := v12];
			globalState.f5, v0 := if (--|(map v24 : D0 | v24 in v27 :: (v24) := (-v0))| in v23) then v23[--|(map v24 : D0 | v24 in v27 :: (v24) := (-v0))|] else fm0(v0, v0, v0, cf1, globalState), safeDivisionInt(550, safeDivisionInt(v0, v0));
			var v28: seq<set<bool>> := [v11];
			globalState.f0 := !((cf1 || fm1(cf1, globalState)) in (v11 * v28[safeIndex(v0, |v28|)]));
			v12 := v12;
			globalState.f2 := v2;
		case DC0(cf0) =>
			m0(!v2, v18, v0, v17, globalState);
			m0(!v2 <==> v2, v18, v0, v17, globalState);
			v12 := v12 + v12;
			globalState.f5 := v0;
		case DC2(cf2) =>
			m0(v2, v18, -(v0 - v0), v17, globalState);
			var v29: multiset<D0> := multiset{v20, v20, v20.(cf1 := v2)};
			var v30: set<array<int>> := {v3};
			globalState.f2 := (if (v2) then multiset{DC1(v5[safeIndex(|v17|, |v5|)])} else v29) == (if (false) then (v29[v20 := abs(v0)])[fm3(v20.cf1, v0, |v30|, v0, globalState) := abs(v0)] else v29);
			v6 := v6[v0 := -v0 < v0];
			var v31 := new C1();
	}
	
	if (false) {
		globalState.f2 := v2;
		globalState.f0 := !v2;
		if (v2) {
			var v32: map<int, array<int>> := map[v0 := v3];
			var v33: multiset<bool> := multiset{v2, fm1(v2, globalState), v2};
			v32 := v32[|v33| := v3];
			v1 := v1;
			var v34 := DC12(|v12|, v0, v12, 0x15e, |v5|);
			v17 := v17[v12 < "fvye" := v34.cf17];
			var v35: array<bool> := new bool[24];
			v35[safeIndex(765, v35.Length)] := v2;
			v35[safeIndex(765, v35.Length)] := v2;
			globalState.f2 := v2;
		} else {
			globalState.f2 := v2;
			v0 := |(v12 + "poavktqf")|;
			var v36: array<map<int, bool>> := new map<int, bool>[3](i2 => v6);
			v36[safeIndex(47, v36.Length)] := map[v0 := v2];
			var v37 := DC21(v2, v2, (v5[safeIndex(|"wcd"|, |v5|) := v2])[safeIndex(-0x241, |v5[safeIndex(|"wcd"|, |v5|) := v2]|) := v2], v2, v0);
			v36[safeIndex(47, v36.Length)] := v6[v0 := v37.cf27];
			v5 := (v5[safeIndex(v0, |v5|) := v2] + v5) + v5;
			var v38: T1 := new C4(v0);
			var v39: seq<T1> := [v38, v38];
			m0(v2, v18, safeDivisionInt(v0, fm0(v0, v0, v0, v2, globalState)), v17[v2 := |v39|], globalState);
		}
		
		var v40: map<int, D8> := map[v0 := DC21(v2, v2, v5, v2, v0)];
		v40 := fm33(v2, v18, globalState);
		var v41: array<char> := new char[25](i3 => v18);
		v41[safeIndex(349, v41.Length)] := v18;
		v41[safeIndex(349, v41.Length)] := v18;
	} else {
		var v42: map<int, int> := map[v0 := v0];
		v42 := v42[v0 := v0];
		v0 := -|v17|;
		var v43: C1 := new C1();
		v43 := v43;
		globalState.f2 := v2;
		globalState.f2 := v2;
	}
	
	globalState.f2 := v0 < v0;
	var v44: map<string, int> := map[v12 := v0];
	v44 := v44[v12 := safeDivisionInt(v0, v0)];
	m0(v2, v18, v0, v17 + v17, globalState);
	var v45: multiset<bool> := multiset{v2};
	globalState.f2 := v45 !! v45;
	if (v2) {
		var v46: multiset<char> := multiset{v18, v18};
		if (fm0(if (v18 in v46) then v46[v18] else v0, |v1|, v0, v2, globalState) <= v0) {
			m0(v2, v18, v0, v17, globalState);
			var v47: T1 := new C3(997);
			var v48: set<T1> := {v47};
			var v49 := DC20(v48 * v48);
			v49 := if (v0 < v0) then v49 else v49;
			var v50: array<bool> := new bool[25] [v2, v2, v2, v2, v2, v2, false, v2, true, true, v2, v2, true, v2, v2, v2, v47.fm4(v9, DC0(map[0x319 := v2]), multiset([v2]), globalState), v2, v2, v2, v2, v2, v2, v2, v2];
			var v51: map<seq<int>, array<bool>> := map[v1 := v50];
			var v52: map<map<seq<int>, array<bool>>, bool> := map[v51 := DC24(v2, false, v2).cf34];
			v52 := v52[v51 := v2 in v11];
			globalState.f0 := true;
			globalState.f2 := fm1((seq(abs(671), i4  => (v18)))[safeIndex(v0, |(seq(abs(671), i4  => (v18)))|) := v18] < v12, globalState);
		} else {
			v0, globalState.f2 := fm0(v0, v0, -v0, !v2, globalState), multiset{v18} >= (if (!v2) then v46 else v46);
			var v53: map<bool, bool> := map[v2 := v2];
			var v54 := DC7(|v12|, map[v53 := v0], v0);
			var v55: multiset<D2> := multiset{v54};
			var v56: map<int, int> := map[if (v54 in v55) then v55[v54] else v0 := v0];
			var v57: T0 := new C1();
			var v58: set<T0> := {v57};
			v56 := v56[safeModuloInt(v0, v0) := |(v58 * v58)|];
			v3 := v3;
			globalState.f5 := v0;
			var v59: T1 := new C3(v0);
			var v60 := DC5(v59);
			var v61: map<multiset<bool>, T1> := map[multiset(v5) := v60.cf5];
			var v62 := DC24(v2, v2, v2);
			var v63: array<T1> := new T1[15] [v59, if (v2) then v59 else v59, v59, if (v45 in v61) then v61[v45] else v59, v59, v59, v59, if (v62.cf35) then v59 else v59, v59, v59, v59, v59, v59, v59, v59];
			v63[safeIndex(122, v63.Length)] := v59;
			var v64 := DC38(v17);
			var v65: multiset<D13> := multiset{v64, v64};
			var v66: array<bool> := new bool[3] [!v2, v65 < v65, v2];
			v66[safeIndex(556, v66.Length)] := v2;
			var v67: C3 := new C3(v0);
			v63[safeIndex(122, v63.Length)], v66[safeIndex(556, v66.Length)], v67 := v59, v2, v67;
		}
		
		var v68: array<T0> := new T0[5];
		var v69: T0 := new C2(v2);
		v68[safeIndex(929, v68.Length)] := v69;
		v12, v68[safeIndex(929, v68.Length)] := if (v2) then v12 else seq(abs(0x3aa), i5  => (DC27(v2, v0, v18, v45).cf40)), v69;
		globalState.f0 := if (v0 <= v0) then v2 else v2;
		globalState.f5 := -v0;
		globalState.f5 := |(fm34(v0, v0, !v2, globalState))[v11 := -582]|;
	} else {
		m0(v2 <== v2, v18, v0, v17, globalState);
		v2 := v0 !in v1;
		var v70: array<bool> := new bool[28] [false, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, !v2, fm1(!v2, globalState), v2, v2, v2, v2, v2, v2, v2, v2, !v2, v2, false, true, v2];
		var v71 := new C0(true, v70);
		var v72 := new C2(v71.f10);
		var v73: array<array<int>> := new array<int>[6];
		var v74, v75, v76, v77 := v72.m1(v0, v73, |(multiset(v5) + multiset(v5))|, v0, globalState);
	}
	
	var v78: array<bool> := new bool[10];
	var v79: C0 := new C0(v2, v78);
	var v80: set<C0> := {v79};
	var v81: map<bool, set<C0>> := map[v2 := v80];
	var v82: map<int, int> := map[|(if (v79.fm4(v9, DC0(v6), v45, globalState) in v81) then v81[v79.fm4(v9, DC0(v6), v45, globalState)] else v80)| := v0];
	var v83: T1 := new C0(v79.f10, v79.f11);
	var v84: map<T1, int> := map[v83 := v0];
	var v85: map<bool, map<T1, int>> := map[v2 := v84];
	v82 := v82[if (v79.f10) then v0 else fm0(v0, v0, |v85|, v79.f10, globalState) := v0];
	var v86: array<D13> := new D13[10](i6 => DC38(map[v2 := v0]));
	var v87: set<array<D13>> := {v86};
	if (v87 !! v87) {
		v11 := fm17(!v79.f10, -v0, v79.f10, v0, globalState);
		var v88: C4 := new C4(v0);
		var v89: seq<C4> := [v88];
		var v90: set<int> := {|v89|, v88.f9, v88.f9};
		var v91: map<C0, int> := map[v79 := |v90|];
		v91 := v91[v79 := -v0];
		v78 := new bool[13];
		var v92 := DC8(v12);
		globalState.f2 := v92.cf10 <= ((seq(abs(956), i7  => ('a'))) + v12);
		var v93: C1 := new C1();
		var v94: map<C1, bool> := map[v93 := v2];
		v0 := |v94| * v88.f9;
	} else {
		globalState.f2 := v0 != |v12[safeIndex(v0, |v12|) := v18]|;
		var v95 := new C4(v0);
		v2 := v2;
		var v96 := new C1();
		globalState.f5 := v0 + v0;
	}
	
	var v97: C1 := new C1();
	var v98: seq<C1> := [v97, v97];
	v97 := v98[safeIndex(v0, |v98|)];
	print v0, "\n";
	print v1 == [292, 292], "\n";
	print v2, "\n";
	print v3[0], "\n";
	print v3[1], "\n";
	print v3[2], "\n";
	print |v4|, "\n";
	print v5 == [true], "\n";
	print v6 == map[1 := false], "\n";
	print v7 == map[292 := multiset{map[1 := false]}], "\n";
	print v9.cf0 == map[1 := false, 2 := false], "\n";
	print v10 == multiset{map[1 := false], map[1 := false, 2 := false]}, "\n";
	print v11 == {false}, "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print |globalState.f3|, "\n";
	print globalState.f4, "\n";
	print globalState.f5, "\n";
	print globalState.f6 == multiset{map[1 := false], map[1 := false], map[1 := false, 2 := false]}, "\n";
	print globalState.f7, "\n";
	print globalState.f8 == {false}, "\n";
	print v12, "\n";
	print i0, "\n";
	print v17 == map[false := 7], "\n";
	print v18, "\n";
	print v20.cf1, "\n";
	print v21 == map[false := DC1(false)], "\n";
	print v44 == map["aiwqmuo" := 1], "\n";
	print v45 == multiset{false}, "\n";
	print v79.f10, "\n";
	print |v80|, "\n";
	print |v81|, "\n";
	print v82 == map[1 := -1, -1 := -1], "\n";
	print |v84|, "\n";
	print |v85|, "\n";
	print v86[0].cf55 == map[true := -1], "\n";
	print v86[1].cf55 == map[true := -1], "\n";
	print v86[2].cf55 == map[true := -1], "\n";
	print v86[3].cf55 == map[true := -1], "\n";
	print v86[4].cf55 == map[true := -1], "\n";
	print v86[5].cf55 == map[true := -1], "\n";
	print v86[6].cf55 == map[true := -1], "\n";
	print v86[7].cf55 == map[true := -1], "\n";
	print v86[8].cf55 == map[true := -1], "\n";
	print v86[9].cf55 == map[true := -1], "\n";
	print |v87|, "\n";
	print |v98|, "\n";
}