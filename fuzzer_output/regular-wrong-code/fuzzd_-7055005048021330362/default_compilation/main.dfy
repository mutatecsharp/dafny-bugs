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
datatype D0 = DC1(cf1: bool) | DC0(cf0: array<bool>) | DC2(cf2: D0)
datatype D1 = DC4(cf4: char, cf5: map<int, int>, cf6: seq<int>) | DC5(cf7: bool, cf8: map<bool, int>, cf9: int, cf10: multiset<int>) | DC3(cf3: set<set<char>>)
datatype D2 = DC6(cf11: set<int>)
datatype D3 = DC8(cf13: int, cf14: map<multiset<bool>, int>, cf15: D1, cf16: bool) | DC9(cf17: int, cf18: int, cf19: int) | DC7(cf12: multiset<bool>)
datatype D4 = DC11(cf21: set<int>) | DC10(cf20: C0) | DC12(cf22: D4)
datatype D5 = DC13(cf23: C1)
datatype D6 = DC14(cf24: C3)
datatype D7 = DC16(cf26: bool, cf27: bool, cf28: C3, cf29: bool) | DC15(cf25: multiset<array<bool>>)
datatype D8 = DC17(cf30: map<int, array<bool>>)
datatype D9 = DC19(cf32: int, cf33: int, cf34: bool, cf35: int, cf36: set<seq<bool>>) | DC20(cf37: string, cf38: bool, cf39: int) | DC21(cf40: array<bool>, cf41: bool) | DC18(cf31: set<set<bool>>)
trait T0 {
	var f21 : string
	var f22 : bool
	function fm0(globalState: GlobalState): set<bool> 
	function fm1(p0: string, p1: bool, p2: bool, globalState: GlobalState): bool 
}

trait T1 extends T0 {
	var f23 : seq<seq<int>>
	function fm2(p0: bool, globalState: GlobalState): bool 
}

trait T2 extends T1 {
	method m1(p0: bool, p1: map<map<int, bool>, array<bool>>, globalState: GlobalState) returns (r0: int) 
}

class GlobalState {
	const f0 : bool
	const f1 : map<string, int>
	const f2 : int
	var f3 : int
	const f4 : int
	const f5 : bool
	var f6 : string
	const f7 : bool
	var f8 : string
	const f9 : string
	const f10 : int
	const f11 : bool
	var f12 : char
	const f13 : int
	const f14 : int
	const f15 : int
	const f16 : array<bool>
	const f17 : array<int>
	const f18 : bool
	var f19 : multiset<int>
	var f20 : int
	constructor (f0 : bool, f1 : map<string, int>, f2 : int, f3 : int, f4 : int, f5 : bool, f6 : string, f7 : bool, f8 : string, f9 : string, f10 : int, f11 : bool, f12 : char, f13 : int, f14 : int, f15 : int, f16 : array<bool>, f17 : array<int>, f18 : bool, f19 : multiset<int>, f20 : int) {
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
	}
	
}

class C0 {
	var f26 : set<string>
	constructor (f26 : set<string>) {
		this.f26 := f26;
	}
	
	function fm7(p0: int, globalState: GlobalState): bool {
		(false && true) <== false
	}
	function fm8(globalState: GlobalState): set<set<char>> {
		DC3({{'r'}}).cf3
	}
}

class C1 {
	var f24 : array<multiset<bool>>
	const f25 : bool
	constructor (f24 : array<multiset<bool>>, f25 : bool) {
		this.f24 := f24;
		this.f25 := f25;
	}
	
	method m3(p0: bool, p1: bool, p2: bool, globalState: GlobalState) {
		var v0 := true;
		var v1 := 0x3d0;
		var v2: set<int> := {v1, v1, v1};
		v0 := p1 || (v2 != v2);
		var v3: array<bool> := new bool[21];
		var v4 := DC0(v3);
		v4 := v4.(cf0 := v3);
		var v5: seq<bool> := [p0, f25];
		var v6: multiset<bool> := multiset{!!v0, f25, p0, v5[safeIndex(fm6(globalState), |v5|)], f25};
		var v7 := "ekkjfqs";
		v3[safeIndex(442, v3.Length)] := v7 < v7;
		var v8: seq<seq<bool>> := [v5, v5, v5];
		v6, v0, v3[safeIndex(442, v3.Length)] := v6, v5 == v8[safeIndex(|v7|, |v8|)], p0;
		var i0 := 0;
		while (v3[safeIndex(442, v3.Length)])
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v9: array<map<D0, int>> := new map<D0, int>[19];
			v9 := v9;
			v3 := v3;
			var v10 := DC1(f25);
			if (v10.cf1) {
				globalState.f3 := |{-v1}| + |multiset{v1}|;
				v0 := p0;
				var v11: set<string> := {"u", "px"};
				var v12 := new C0(v11);
				globalState.f3 := |(v7 + (seq(abs(687), i1  => ('c'))))|;
				var v13: map<bool, int> := map[p1 := v1];
				var v14: seq<int> := [v1, |fm9(v0, p2, globalState)|];
				var v15: map<int, multiset<int>> := map[v1 := (multiset(v14))[v1 := abs(v1)]];
				var v16: multiset<int> := multiset{|v5|};
				var v17 := DC5(v10.cf1, v13, -|v7|, if (v1 in v15) then v15[v1] else v16);
				v17 := v17;
			} else {
				v3[safeIndex(442, v3.Length)] := p0 <== p0;
				v1 := if (p1) then 0x203 else v1;
				var v18: set<string> := {v7};
				var v19 := new C0(v18);
				globalState.f20 := safeModuloInt(0xdb, v1);
				v0 := v0;
			}
			
			match (fm10(globalState)) {
				case DC4(cf4, cf5, cf6) =>
					var v20: array<D0> := new D0[20](i2 => v10.(cf1 := p1));
					var v21: map<array<D0>, int> := map[v20 := v1];
					var v22: map<bool, map<array<D0>, int>> := map[v3[safeIndex(442, v3.Length)] := v21];
					v22 := v22[f25 := v21];
					var v23: array<int> := new int[5];
					var v24: map<bool, array<int>> := map[false := v23];
					v24 := v24[p0 := v23];
					var v25: array<string> := new seq<char>[8](i3 => seq(abs(-0x361), i4  => (cf4)));
					v25[safeIndex(96, v25.Length)] := fm11(false, globalState);
					v25[safeIndex(96, v25.Length)] := v7;
					v1 := -((v1 - v1) - cf6[safeIndex(v1, |cf6|)]);
				case DC5(cf7, cf8, cf9, cf10) =>
					var v26 := 't';
					v0 := if (p1) then v26 !in "bdncgjj" else v3[safeIndex(442, v3.Length)];
					v3[safeIndex(442, v3.Length)] := v2 > (v2 - v2);
					var v27: seq<int> := [cf9, cf9];
					var v28: multiset<multiset<bool>> := multiset{v6};
					var v29: map<int, int> := map[v1 := cf9];
					v3[safeIndex(442, v3.Length)] := fm12(|(v27 + v27)|, v6[v0 := abs(cf9)] * v6, v3[safeIndex(442, v3.Length)], multiset{multiset{fm12(cf9, v6, p2, v28, globalState), p2}, multiset(v5), multiset{v5[safeIndex(v27[safeIndex(|v29|, |v27|)], |v5|)], v0}, multiset{p0, f25}}, globalState);
					var v30: map<int, bool> := map[cf9 := p2];
					var v31: map<multiset<int>, string> := map[cf10 := fm11(if (v1 in v30) then v30[v1] else v0, globalState)];
					v31 := map[cf10 := seq(abs(786), i5  => (v26))] + v31;
				case DC3(cf3) =>
					v0 := (v1 * v1) <= v1;
					globalState.f20 := fm6(globalState);
					globalState.f3 := v1;
					var v32: set<string> := {v7};
					var v33 := new C0(v32 * v32);
			}
			
		}
		var v34: map<int, int> := map[|v5| := |v7|];
		var v35: seq<map<int, int>> := [v34];
		v3[safeIndex(442, v3.Length)] := (set v36 : int | v36 in v35[safeIndex(v1, |v35|)] :: (v36 - |{DC1(true).cf1}|)) != (v2 + v2);
		var v37: map<bool, int> := map[p2 := v1];
		var v38: multiset<int> := multiset{|v2|};
		var v39 := DC5(v3[safeIndex(442, v3.Length)], v37, v1, v38);
		globalState.f3 := match v39 {
			case DC4(cf4, cf5, cf6) => |(if (v3[safeIndex(442, v3.Length)]) then seq(abs(0x1de), i6  => (cf6)) else seq(abs(0x1b5), i7  => ([|{cf4}|])))|
			case DC5(cf7, cf8, cf9, cf10) => v1
			case DC3(cf3) => v1 - 0x249
		};
	}
}

class C2 extends T2 {
	constructor (f23 : seq<seq<int>>, f21 : string, f22 : bool) {
		this.f23 := f23;
		this.f21 := f21;
		this.f22 := f22;
	}
	
	function fm2(p0: bool, globalState: GlobalState): bool {
		(if (false) then f22 else f22) ==> !false
	}
	function fm0(globalState: GlobalState): set<bool> {
		{f22, f22, false, f22, f22}
	}
	function fm1(p0: string, p1: bool, p2: bool, globalState: GlobalState): bool {
		(99 - |[|multiset{f22}|]|) == |({f22, f22} + {true})|
	}
	function fm3(globalState: GlobalState): bool {
		!((if (f22) then -0x8d else 0x1a3) <= |[f21]|)
	}
	method m1(p0: bool, p1: map<map<int, bool>, array<bool>>, globalState: GlobalState) returns (r0: int) {
		var v0 := 191;
		var v1: seq<int> := [v0, v0];
		var v2: multiset<seq<int>> := multiset{f23[safeIndex(v0, |f23|)], v1};
		var v3: set<int> := {if (v1 in v2) then v2[v1] else v0};
		var v5: set<set<int>> := {v3, v3, v3, (set v4 : int | (132 <= v4) && (v4 < 0x3dd) :: (v4 + |map[v0 := p0]|)) - v3};
		v5 := (if (p0) then {v3} else v5) + v5;
		if (true) {
			globalState.f20 := v0;
			var v6: map<int, seq<char>> := map[|multiset([v0, |[f22]|, v0])| := f21];
			var v7 := 'u';
			var v8: seq<seq<char>> := [f21, if (v0 in v6) then v6[v0] else f21, [v7]];
			f21 := (v8[safeIndex(v0, |v8|)])[safeIndex(v0, |v8[safeIndex(v0, |v8|)]|) := v7];
			var v9: array<bool> := new bool[19](i0 => true);
			var v10 := DC0(v9);
			var v11: seq<array<bool>> := [v9];
			var v12: array<array<bool>> := new array<bool>[6] [v9, v9, v9, v10.cf0, v9, v11[safeIndex(v0, |v11|)]];
			var v13 := DC1(f22);
			var v14: multiset<bool> := multiset{f22};
			var v15: map<char, array<array<bool>>> := map[fm4(f22, v14, globalState) := v12];
			v12 := if ((v13.(cf1 := false)).cf1) then if ('x' in v15) then v15['x'] else v12 else v12;
			v12[safeIndex(623, v12.Length)] := v9;
			v12[safeIndex(623, v12.Length)] := v9;
			var v16: multiset<int> := multiset{v0};
			r0 := |v16| * v1[safeIndex(0x351, |v1|)];
		} else {
			var v17: map<bool, int> := map[p0 := v0];
			var v18: map<map<bool, int>, bool> := map[v17 := f22 <== f22];
			v18 := map[v17 := true];
			f22 := f22;
			var v19: array<bool> := new bool[16];
			var v20 := m0(v19, true && p0, globalState);
			f22 := v20;
			var v21: seq<set<int>> := [v3, v3];
			globalState.f20 := |(v21 + v21[safeIndex(0xb6, |v21|) := {v0, -v0}])|;
		}
		
		var v22 := DC1(true);
		var v23 := DC2(v22);
		var v24: seq<string> := [f21];
		var v25: seq<bool> := [f22, f22, !p0];
		var v26: array<int> := new int[15] [|v24|, v0, v0, v0, 0x1aa, safeModuloInt(v0, v0), -v0, 834, v0, |v24|, v0, -0x19c, v0, v0, |v25|];
		v26[safeIndex(925, v26.Length)] := |v1| * v0;
		globalState.f20, v0, globalState.f20, v23, v26[safeIndex(925, v26.Length)] := |f23|, v0, --v0, v23, safeDivisionInt(v0, --|(v25 + v25)|);
		r0 := v26[safeIndex(925, v26.Length)];
		var i1 := 0;
		while (false)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			var v27 := 'i';
			var v28: map<int, char> := map[v26[safeIndex(925, v26.Length)] := v27];
			var v29: map<bool, int> := map[f22 := |v28|];
			var v30: set<seq<int>> := {fm5(v26[safeIndex(925, v26.Length)], |[true]|, v0, globalState)};
			v26[safeIndex(925, v26.Length)] := (|v29| - 914) - |(if (false) then v30 else {v1, [v0, v26[safeIndex(925, v26.Length)], v0]})|;
			var v31: array<multiset<bool>> := new multiset<bool>[25](i2 => multiset{!f22, p0});
			var v32 := new C1(v31, true);
			var v33: set<char> := {v27};
			var v34 := DC3({v33});
			var v35: set<D1> := {v34, v34};
			var v36: set<set<D1>> := {v35 - v35, {fm10(globalState), v34, fm10(globalState), v34}};
			v36 := v36;
			var v37: array<multiset<string>> := new multiset<string>[20];
			var v38: multiset<string> := multiset{f21, f21[safeIndex(|v25|, |f21|) := v27], f21, f21};
			v37[safeIndex(214, v37.Length)] := v38;
			var v39 := DC1(v32.f25);
			var v40: set<bool> := {v39.cf1, false, !f22, p0};
			var v41: map<C1, set<bool>> := map[v32 := v40];
			v37[safeIndex(214, v37.Length)] := fm13(f22, v40, safeDivisionInt(|f21|, |multiset{if (v32 in v41) then v41[v32] else v40, v40, {false, p0, v32.f25, v32.f25, v32.f25}}|), p0, globalState);
		}
		var v43 := DC3({set v42 : char | v42 in multiset(f21) :: (v42)});
		var v44: map<int, bool> := map[v0 := p0];
		f21, r0, v26, globalState.f20, globalState.f20 := f21 + f21, match v43 {
			case DC4(cf4, cf5, cf6) => v26[safeIndex(925, v26.Length)] - DC5(f22, map[!f22 := |cf6|], |{f22, false}|, multiset{v0, |f21|}).cf9
			case DC5(cf7, cf8, cf9, cf10) => v0
			case DC3(cf3) => v26[safeIndex(925, v26.Length)] + 0x1bb
		}, v26, v26[safeIndex(925, v26.Length)] - v0, |(v44 + v44)|;
		var v45 := 's';
		var v46: set<bool> := {f22, f22, p0};
		var v47: map<int, int> := map[|v46| := fm6(globalState)];
		r0 := |(match DC4(v45, v47, v1) {
			case DC4(cf4, cf5, cf6) => map[p0 := f22] + map[p0 := false]
			case DC5(cf7, cf8, cf9, cf10) => map[p0 := cf7] + map[false := f22]
			case DC3(cf3) => map[p0 := true] + map[f22 := p0]
		})|;
	}
	method m2(p0: int, globalState: GlobalState) returns (r0: string, r1: bool) {
		var v0: array<int> := new int[13];
		v0[safeIndex(931, v0.Length)] := |f21|;
		v0[safeIndex(931, v0.Length)] := p0 - p0;
		var v1: seq<bool> := [f22];
		var v2: set<bool> := {true};
		v1 := [v0[safeIndex(931, v0.Length)] !in map[0x15e := |v2|], f22, f22 && f22];
		if (f22) {
			var v3: multiset<bool> := multiset{f22};
			globalState.f3 := if (!(if (!f22) then f22 else f22)) then safeModuloInt(fm6(globalState), v0[safeIndex(931, v0.Length)]) else |v3|;
			var v4 := 'u';
			var v5: map<int, int> := map[|v3| := v0[safeIndex(931, v0.Length)]];
			var v6: seq<int> := [v0[safeIndex(931, v0.Length)]];
			var v7 := DC4(v4, v5, v6);
			v7 := v7;
			var v8: map<bool, int> := map[fm12(v0[safeIndex(931, v0.Length)], multiset([f22]), f22, multiset{multiset{false, !true}, v3, v3, multiset{f22, f22, f22}, v3}, globalState) := -v0[safeIndex(931, v0.Length)]];
			var v9: map<bool, bool> := map[f22 := false];
			var v10 := DC1(if (true in v9) then v9[true] else f22);
			v8 := v8[v10.cf1 := -p0];
			var v11: array<bool> := new bool[27] [f22, f22, f22, f22, f22, f22, f22, f22, false, f22, f22, true, f22, f22, f22, f22, f22, true, f22, false, f22, f22, !f22, f22, f22, false, f22];
			var v12 := DC0(v11);
			var v13 := DC2(v12);
			var v14: map<D0, int> := map[v13 := |f21|];
			v14 := v14[v13 := p0];
			r1 := f22;
		} else {
			globalState.f20 := v0[safeIndex(931, v0.Length)] - |"etxvn"|;
			globalState.f12 := 'd';
			var v15: array<set<int>> := new set<int>[15](i0 => {v0[safeIndex(931, v0.Length)]});
			var v16 := DC6({v0[safeIndex(931, v0.Length)], p0});
			v15[safeIndex(983, v15.Length)] := v16.cf11;
			var v17: map<bool, bool> := map[fm1(f21, true, false, globalState) := f22];
			v15[safeIndex(983, v15.Length)] := fm14(v17, p0, globalState);
			match (v16) {
				case DC6(cf11) =>
					globalState.f20 := fm6(globalState) + safeDivisionInt(|f21|, v0[safeIndex(931, v0.Length)]);
					v16 := v16;
					v17 := v17;
					f22 := (if (f22) then f22 else f22) && fm2(f22, globalState);
			}
			
			var v18: array<multiset<bool>> := new multiset<bool>[4];
			var v19: map<array<multiset<bool>>, map<int, int>> := map[v18 := map[p0 := |f21|] + map[|f21| := p0]];
			var v20: map<int, int> := map[v0[safeIndex(931, v0.Length)] := p0];
			v19 := v19[v18 := v20];
		}
		
		if (true) {
			var v21: set<int> := {v0[safeIndex(931, v0.Length)]};
			var v22 := DC6(v21 * v21);
			v22 := v22;
			var v23: set<string> := {f21, f21};
			var v24: C0 := new C0(v23);
			v24 := v24;
			var v25: array<bool> := new bool[15](i1 => if (f22) then f22 else f22);
			v25[safeIndex(184, v25.Length)] := f22;
			var v26: multiset<bool> := multiset{f22, fm1(f21, f22, f22, globalState)};
			v25[safeIndex(184, v25.Length)] := !((0xd3 == v0[safeIndex(931, v0.Length)]) <==> (v26 !! multiset(v1)));
			r1 := f22;
			var v27: array<map<bool, int>> := new map<bool, int>[14];
			var v28: multiset<string> := multiset{f21, fm11(false, globalState), "bdbdchmi"};
			globalState.f20, v25[safeIndex(184, v25.Length)], f22, v27 := if (f21 in v28) then v28[f21] else v0[safeIndex(931, v0.Length)], v25[safeIndex(184, v25.Length)], v25[safeIndex(184, v25.Length)], v27;
		} else {
			var v29: array<bool> := new bool[8];
			v29 := v29;
			var v30: multiset<bool> := multiset{f22, false};
			var v32: set<multiset<bool>> := {v30};
			var v33: seq<map<multiset<bool>, int>> := [map v31 : multiset<bool> | v31 in v32 :: (v31) := (-p0)];
			var v34: map<multiset<bool>, map<multiset<bool>, int>> := map[v30 := v33[safeIndex(-v0[safeIndex(931, v0.Length)], |v33|)]];
			v34 := v34[multiset{f22} * v30 := v33[safeIndex(v0[safeIndex(931, v0.Length)], |v33|)]];
			v1 := v1 + v1;
			if (!f22) {
				v0[safeIndex(931, v0.Length)], f22, v0 := -(if (f22) then (if (f22 in v30) then v30[f22] else -0x3a8) + v0[safeIndex(931, v0.Length)] else p0), f22, v0;
				var v35: map<bool, array<bool>> := map[f22 := v29];
				var v36 := DC1(f22);
				var v37: map<array<bool>, D0> := map[if (f22 in v35) then v35[f22] else v29 := v36];
				globalState.f20 := |(v37 + v37)|;
				r1 := v1[safeIndex(p0, |v1|)];
				var v38 := 'x';
				var v39: multiset<char> := multiset{v38, v38};
				v39, f22 := v39 - v39, f22;
				var v40: map<bool, bool> := map[fm3(globalState) := f22];
				var v41: seq<int> := [p0];
				var v42: set<int> := {|v41|};
				v40 := v40[{0x1c, p0, |v30|} !! v42 := fm3(globalState)];
			} else {
				var v43: map<int, bool> := map[v0[safeIndex(931, v0.Length)] := f22];
				f22 := |f21| <= safeModuloInt(|v43|, v0[safeIndex(931, v0.Length)]);
				r1 := !(if (f22) then !f22 else if (if (643 in v43) then v43[643] else f22) then f22 else f22);
				var v44 := 'g';
				var v45: set<char> := {v44, v44};
				var v46 := DC3({v45, {v44}});
				var v47: seq<D1> := [v46, fm10(globalState)];
				globalState.f20 := |v47|;
				var v48: array<array<multiset<bool>>> := new array<multiset<bool>>[28];
				var v49: array<multiset<bool>> := new multiset<bool>[6](i2 => v30);
				v48[safeIndex(40, v48.Length)] := v49;
				v48[safeIndex(40, v48.Length)] := v49;
				v29 := v29;
			}
			
			var v50: array<multiset<bool>> := new multiset<bool>[2](i3 => v30[f22 := abs(p0)]);
			var v51: C1 := new C1(v50, f22);
			var v52: seq<C1> := [v51, v51];
			r1 := !f22 <==> ({v52, v52, v52} > {v52, v52});
		}
		
		var v53: set<string> := {f21, f21};
		var v54 := new C0(v53);
		var i4 := 0;
		while (p0 >= p0)
			decreases 100 - i4
		{
			if (i4 >= 100) {
				break;
			}
			
			i4 := i4 + 1;
			var v55: multiset<bool> := multiset{f22};
			var v56 := DC7(multiset{f22});
			v55 := v56.cf12;
			v0[safeIndex(931, v0.Length)] := p0;
			globalState.f20 := fm6(globalState);
			var v57: set<int> := {p0, v0[safeIndex(931, v0.Length)], p0};
			var v58: map<bool, bool> := map[if (f22) then f22 else f22 := f22];
			var v59 := DC1(true);
			var v60 := DC2(v59);
			v57, r1, v58, v60 := fm14(map[if (!f22 in v58) then v58[!f22] else f22 := f22], |f21|, globalState) - v57, v1[safeIndex(|fm15(globalState)|, |v1|)], if (f22) then v58 else v58, v60;
		}
		var v62 := DC6(set v61 : int | (-0x19 <= v61) && (v61 < 428) :: (safeModuloInt(v61, v0[safeIndex(931, v0.Length)])));
		r0 := match v62 {
			case DC6(cf11) => f21
		};
		var v63: multiset<bool> := multiset{f22, f22, true, f22, f22};
		var v64: map<seq<bool>, int> := map[v1 := p0];
		var v65: array<bool> := new bool[1];
		var v66: seq<array<bool>> := [v65, v65];
		var v67: map<int, array<bool>> := map[|v64| := v66[safeIndex(p0, |v66|)]];
		var v68: multiset<multiset<bool>> := multiset{multiset(v1)};
		r1 := fm12(|"hsldowxme"|, if (!f22) then v63 else v63, v67 == map[0x58 := v65], v68, globalState);
	}
}

class C3 extends T0 {
	const f27 : bool
	var f28 : bool
	constructor (f27 : bool, f28 : bool, f21 : string, f22 : bool) {
		this.f27 := f27;
		this.f28 := f28;
		this.f21 := f21;
		this.f22 := f22;
	}
	
	function fm0(globalState: GlobalState): set<bool> {
		(if (false) then {f22, true} else {f27, f28}) + {f27}
	}
	function fm1(p0: string, p1: bool, p2: bool, globalState: GlobalState): bool {
		f27
	}
	method m4(p0: int, p1: int, p2: int, p3: int, globalState: GlobalState) returns (r0: array<array<char>>) {
		f22 := false;
		for i0 := |f21| to p1 {
			var v0: array<C1> := new C1[3];
			var v1: multiset<array<C1>> := multiset{v0, v0};
			v1 := (v1 * v1) + (v1 * v1);
			globalState.f6 := f21;
			if (f27) {
				var v2: map<bool, int> := map[f22 := p1];
				var v3: set<map<bool, int>> := {v2, v2};
				f22 := ({map[f28 := p3], v2, v2, map[f28 := p2]} + v3) !! v3;
				var v4: set<string> := {"cafmfqmaf", seq(abs(-524), i1  => ('m'))};
				var v5: C0 := new C0(v4);
				var v6 := DC10(v5);
				globalState.f20, v5 := p3, v6.cf20;
				var v7: seq<int> := [p1, fm6(globalState), p2];
				v7 := v7;
				var v8: array<bool> := new bool[10];
				v8, globalState.f3 := v8, if (true) then safeDivisionInt(i0, i0) else safeDivisionInt(i0, p3);
				f22 := false;
			} else {
				var v9: multiset<int> := multiset{p3, p0};
				var v10: map<bool, int> := map[f22 := |v9|];
				var v11: seq<map<bool, int>> := [map[f27 := p0] + v10];
				v11 := v11;
				var v12: set<int> := {p1, -0x101, 555, 0x36f, i0};
				var v14: map<int, set<int>> := map[|"mkm"| := set v13 : int | v13 in (seq(abs(502), i2  => (0x2c7))) :: (v13 + -0x2fa)];
				var v15: set<set<int>> := {v12, if (p1 in v14) then v14[p1] else v12};
				globalState.f3 := |(v15 + {{p0, |v12|, p2}})|;
				var v16 := DC7(multiset{f27, f28});
				var v17: map<int, D3> := map[p3 := v16];
				v17 := v17 + v17;
				f22 := f27 || f27;
				var v18: array<string> := new string[24];
				v18[safeIndex(391, v18.Length)] := f21;
				v18[safeIndex(391, v18.Length)] := f21;
			}
			
			var v19: array<bool> := new bool[5](i3 => f21 == f21);
			v19[safeIndex(792, v19.Length)] := f27;
			v19[safeIndex(792, v19.Length)] := f22 || f27;
		}
		var v20: array<int> := new int[6];
		var v21: array<bool> := new bool[23];
		var v22: seq<bool> := [f28, false, true, f22, f22];
		v21[safeIndex(690, v21.Length)] := v22[safeIndex(p0, |v22|)];
		v21[safeIndex(791, v21.Length)] := f27;
		v20, v21[safeIndex(690, v21.Length)], v21[safeIndex(791, v21.Length)], f28 := v20, f22, v22[safeIndex(-0x213, |v22|)], if (f22) then f28 || f22 else true;
		var v23: set<int> := {|(multiset([p2, p3]))[p2 := abs(p0)]|, p1};
		var v24 := DC6(v23);
		match (v24) {
			case DC6(cf11) =>
				globalState.f3 := p3;
				var v25: seq<array<int>> := [v20, v20, if (f28) then v20 else v20, v20, v20];
				v25 := v25;
				f28 := (seq(abs(0x150), i4  => ('d'))) <= f21;
				var v26 := 't';
				var v27: set<string> := {seq(abs(468), i5  => (v26))};
				var v28: map<char, set<string>> := map[v26 := v27];
				var v29 := new C0(if (v26 in v28) then v28[v26] else {"ce"});
		}
		
		var v31: map<int, bool> := map[p0 := false];
		if ((seq(abs(0x20e), i6  => (map v30 : int | (0x2a8 <= v30) && (v30 < 0x265) :: (v30 + p2) := (f27)))) < [v31, v31, map[p1 := f27], v31]) {
			var v32: array<seq<bool>> := new seq<bool>[5];
			v32[safeIndex(617, v32.Length)] := v22 + [f22, v21[safeIndex(690, v21.Length)], v21[safeIndex(690, v21.Length)], v21[safeIndex(690, v21.Length)]];
			v32[safeIndex(617, v32.Length)] := (if (!f22) then v22 else v22) + v22;
			var v33: array<multiset<bool>> := new multiset<bool>[26](i7 => multiset{!f22, f22});
			var v34: C1 := new C1(v33, f27);
			v20[safeIndex(331, v20.Length)] := p2;
			var v35: seq<int> := [|v22|, |f21|];
			var v36: seq<int> := [0xf4, p3, |multiset(v35)|, |v22|, p2];
			var v37: multiset<int> := multiset{886};
			var v38: map<int, int> := map[|v37| := p0];
			var v39: multiset<bool> := multiset{v21[safeIndex(690, v21.Length)]};
			var v40: multiset<multiset<bool>> := multiset{v39, v39, v39, v39, v39};
			var v41: map<bool, bool> := map[f27 := false];
			v34, globalState.f12, v20[safeIndex(331, v20.Length)], globalState.f3 := v34, if (v36 < fm5(p3, |v38|, p3, globalState)) then f21[safeIndex(p0, |f21|)] else 'k', fm6(globalState) - |v32[safeIndex(617, v32.Length)]|, if (fm12(p1, v39, f28, v40, globalState)) then fm6(globalState) else v36[safeIndex(if (v21[safeIndex(690, v21.Length)] in v39) then v39[v21[safeIndex(690, v21.Length)]] else -|v41|, |v36|)] + p3;
			var v42 := 'e';
			var v43: set<char> := {'e', v42, v42, v42};
			var v44: set<set<char>> := {v43};
			var v45 := DC3(v44);
			match (v45) {
				case DC4(cf4, cf5, cf6) =>
					v21[safeIndex(690, v21.Length)] := v34.f25;
					cf4 := v42;
					v45 := v45.(cf3 := v44);
					var v46 := DC13(v34);
					v34, globalState.f3, f22 := v46.cf23, p3, if (f28) then f27 else f21 < (seq(abs(-738), i8  => (v42)));
				case DC5(cf7, cf8, cf9, cf10) =>
					f22 := multiset(v32[safeIndex(617, v32.Length)] + v32[safeIndex(617, v32.Length)]) > v39;
					v21[safeIndex(690, v21.Length)] := f22;
					f28 := cf7;
					var v47: seq<seq<int>> := [[|map[v22 := f22]|], v36, v36];
					var v48 := new C2(v47, f21, v36[safeIndex(|v32[safeIndex(617, v32.Length)]|, |v36|)] >= 0x285);
				case DC3(cf3) =>
					var v49: set<bool> := {!f27, false, !v34.f25};
					v20[safeIndex(331, v20.Length)] := (|v49| + -0x2fd) - p1;
					globalState.f20, v21[safeIndex(690, v21.Length)], v42 := -((v20[safeIndex(331, v20.Length)] * 0x36c) - safeModuloInt(|map[false := |v22|]|, fm6(globalState))), v21[safeIndex(690, v21.Length)], fm4(-v20[safeIndex(331, v20.Length)] !in [p1], multiset{v34.f25, !f28, f27, f27, v34.f25}, globalState);
					f22 := !v34.f25;
					globalState.f3 := p0;
			}
			
			var v50: map<char, int> := map[v42 := |f21|];
			v50 := v50[v42 := safeModuloInt(p0, p0)];
			f22 := f28;
		} else {
			var v51 := 'r';
			var v52: map<int, char> := map[p1 := v51];
			var v53: array<char> := new char[19] [v51, v51, v51, v51, v51, v51, v51, v51, v51, v51, v51, v51, v51, v51, v51, if (f28) then 'v' else if (fm6(globalState) in v52) then v52[fm6(globalState)] else v51, v51, v51, 'b'];
			v53[safeIndex(376, v53.Length)] := v51;
			v53[safeIndex(376, v53.Length)] := v51;
			var v54: seq<int> := [-p1];
			var v55: seq<int> := [v54[safeIndex(p0, |v54|)]];
			var v56: map<bool, int> := map[f27 := p2];
			var v57 := DC5(f22, v56, p1, multiset(v54));
			var v58: map<char, int> := map[v51 := p2];
			var v59: seq<seq<int>> := [fm5(p1, if (v53[safeIndex(376, v53.Length)] in v58) then v58[v53[safeIndex(376, v53.Length)]] else p0, p2, globalState), [p1], v55];
			var v60: array<multiset<bool>> := new multiset<bool>[17];
			var v61: C1 := new C1(v60, f28);
			var v62: map<bool, C1> := map[v21[safeIndex(690, v21.Length)] := v61];
			var v63: multiset<bool> := multiset{v61.f25};
			var v64: multiset<multiset<bool>> := multiset{v63};
			var v65: map<seq<bool>, bool> := map[v22 := if (p1 in v31) then v31[p1] else fm12(|v62|, v63, false, v64, globalState)];
			var v66: array<seq<int>> := new seq<int>[27] [v54, seq(abs(-664), i9  => (p2)), v55, [v57.cf9, p2, p0, |v54|], v54 + v59[safeIndex(p1, |v59|)], v55, v54, fm5(p3, p0, -0x171, globalState), v54 + v54, v55, fm5(p3, p3, -0x17b, globalState), [0x397], seq(abs(588), i10  => (p1)), v55, v54, v54, v55, (((seq(abs(766), i11  => (p2)))[safeIndex(p0, |(seq(abs(766), i11  => (p2)))|) := |v65|])[safeIndex(p0, |(seq(abs(766), i11  => (p2)))[safeIndex(p0, |(seq(abs(766), i11  => (p2)))|) := |v65|]|) := p1])[safeIndex(|f21|, |((seq(abs(766), i11  => (p2)))[safeIndex(p0, |(seq(abs(766), i11  => (p2)))|) := |v65|])[safeIndex(p0, |(seq(abs(766), i11  => (p2)))[safeIndex(p0, |(seq(abs(766), i11  => (p2)))|) := |v65|]|) := p1]|) := p0], v54, v54, v55, v54, [p1, |map[f27 := p1]|, p3, p2], [|v56|, p1], v54 + v55, v54, v54 + v55];
			v66[safeIndex(543, v66.Length)] := [p0];
			v66[safeIndex(543, v66.Length)] := v55 + v55;
			var v67 := DC9(p3 + p1, p0, p0);
			match (v67) {
				case DC8(cf13, cf14, cf15, cf16) =>
					globalState.f3 := p3 * fm6(globalState);
					var v68: map<int, int> := map[p1 := cf13];
					var v69: set<string> := {f21};
					var v70: C0 := new C0(v69);
					var v71: seq<C0> := [v70, v70];
					var v72: seq<seq<C0>> := [v71, v71, [v70, v70]];
					var v73: multiset<int> := multiset{|multiset(v72)|, cf13, cf13};
					globalState.f20, cf16, v57, cf16, f22 := if ((|f21| * p1) in v68) then v68[|f21| * p1] else -p3, f22, cf15.(cf9 := p2, cf10 := v73, cf8 := v56[f28 := if (false in v56) then v56[false] else p1]), v21[safeIndex(690, v21.Length)], f22;
					var v74: map<bool, bool> := map[!cf16 := |v68| != p1];
					v74, v70, f28, v74 := v74, v70, cf16, v74 + v74;
					f28, f22 := v21[safeIndex(690, v21.Length)] && (if (f28) then v61.f25 else f28), f27;
				case DC9(cf17, cf18, cf19) =>
					var v75: seq<array<multiset<bool>>> := [v60];
					var v76 := new C1(v75[safeIndex(cf18, |v75|)], v21[safeIndex(690, v21.Length)] || f27);
					var v77 := m0(v21, f22, globalState);
					var v78 := m0(v21, v61.f25, globalState);
					f22 := v77;
				case DC7(cf12) =>
					globalState.f20 := p3;
					var v79: map<bool, bool> := map[f27 := if (f22) then true else v61.f25];
					v79 := v79[f27 := v61.f25];
					var v80: array<string> := new string[11] [f21, (fm11(f22, globalState))[safeIndex(668, |fm11(f22, globalState)|) := v51], "jnsquovwx", f21 + "ittgfr", "cqonwvi" + f21, "xwfxaqtqq", "mt", seq(abs(0xf2), i12  => ('v')), (seq(abs(-696), i13  => (v51))) + f21, f21, f21];
					v80[safeIndex(733, v80.Length)] := (seq(abs(-0x7a), i14  => (v51))) + "kmvw";
					v80[safeIndex(733, v80.Length)] := f21;
					v55 := seq(abs(301), i15  => (0xde));
			}
			
			v21[safeIndex(690, v21.Length)] := true;
			var v81: seq<multiset<multiset<bool>>> := [multiset{v63}];
			v62 := v62[fm12(p2, multiset{f27, f22, v61.f25, fm12(p1, v63, f27, multiset{v63}, globalState), fm12(|f21|, multiset{f22, f28}, f22, v64, globalState)}, f22, v81[safeIndex(-0x3d3, |v81|)], globalState) := v61];
		}
		
		forall i16 | 0 <= i16 < v20.Length {
			v20[i16] := i16 * p3;
		}
		var v82: array<array<char>> := new array<char>[24];
		r0 := v82;
	}
}

class C4 extends T1 {
	var f29 : D2
	const f30 : bool
	constructor (f29 : D2, f30 : bool, f23 : seq<seq<int>>, f21 : string, f22 : bool) {
		this.f29 := f29;
		this.f30 := f30;
		this.f23 := f23;
		this.f21 := f21;
		this.f22 := f22;
	}
	
	function fm2(p0: bool, globalState: GlobalState): bool {
		if (f30) then false else f22
	}
	function fm0(globalState: GlobalState): set<bool> {
		{false} * ({f22, f30} * {f30})
	}
	function fm1(p0: string, p1: bool, p2: bool, globalState: GlobalState): bool {
		f22
	}
}

function fm4(p0: bool, p1: multiset<bool>, globalState: GlobalState): char {
	if (0x68 == |{-0x18b}|) then 'm' else 'p'
}
function fm5(p0: int, p1: int, p2: int, globalState: GlobalState): seq<int> {
	seq(abs(0x21), i0  => (|((seq(abs(0x19c), i1  => (map[true := false]))) + [map[false := false], map[false := false], map[true := !false]])|))
}
function fm6(globalState: GlobalState): int {
	safeModuloInt(|"nasbdl"| * -|['h', 'k']|, |((map v0 : int | (-0x8b <= v0) && (v0 < 404) :: (safeModuloInt(v0, 0x99)) := (map v1 : int | v1 in [741] :: (safeDivisionInt(v1, 0x125)) := (!true))) + map[|{-|"u"|}| := map v2 : int | v2 in map[0x16 := 0xcb] :: (v2 - |map[true := true]|) := (false)])|)
}
function fm9(p0: bool, p1: bool, globalState: GlobalState): map<seq<bool>, bool> {
	map v0 : seq<bool> | v0 in ((seq(abs(-0x221), i0  => ([true, false, false]))) + [[true, true, false]]) :: (v0) := (!(!false <==> true))
}
function fm10(globalState: GlobalState): D1 {
	DC3({{'x', 'e'}})
}
function fm11(p0: bool, globalState: GlobalState): string {
	seq(abs(0x22b), i0  => ('d'))
}
function fm12(p0: int, p1: multiset<bool>, p2: bool, p3: multiset<multiset<bool>>, globalState: GlobalState): bool {
	['a', 'h'] == ['k']
}
function fm13(p0: bool, p1: set<bool>, p2: int, p3: bool, globalState: GlobalState): multiset<string> {
	multiset{"jlwcc"} - (multiset{"dulq", "ouob", "wntsnxr"} - multiset{"oikjsmub"})
}
function fm14(p0: map<bool, bool>, p1: int, globalState: GlobalState): set<int> {
	set v0 : int | (-141 <= v0) && (v0 < 493) :: (v0 - |[[false, true], [true]]|)
}
function fm15(globalState: GlobalState): set<char> {
	{'a', 'a', 'r', 'a'} + ({'j'} - (set v0 : char | v0 in {'r'} :: (v0)))
}
function fm16(p0: set<bool>, p1: bool, p2: bool, globalState: GlobalState): seq<seq<int>> {
	[[|[!false]|, |multiset{{!false}}|]]
}
function fm17(p0: int, p1: bool, globalState: GlobalState): multiset<bool> {
	multiset{!DC20("lt", false, 130).cf38, false, true, false} - (multiset{!true} + DC7(multiset([!true, true])).cf12)
}
function fm18(p0: int, p1: seq<int>, p2: int, p3: int, globalState: GlobalState): map<string, int> {
	map["rkakqfc" + "uhejkofsg" := |(if (false) then "olxnmwt" else seq(abs(461), i0  => ('b')))|]
}
function fm19(globalState: GlobalState): map<multiset<int>, bool> {
	map v0 : multiset<int> | v0 in (map[multiset{-0x3e5} := DC7(multiset{!true, true})] + map[multiset{-0x275} := DC7(multiset{true})]) :: (v0) := (--516 <= |"bsjog"|)
}
method m0(p0: array<bool>, p1: bool, globalState: GlobalState) returns (r0: bool) {
	var v0: map<int, bool> := map[-0xfb := p1];
	var v1 := -0x1dd;
	globalState.f3 := |v0| - v1;
	var v2 := "wqmiaea";
	var v3: map<string, int> := map[v2 := v1];
	var v4: seq<int> := [382, v1];
	v3 := (map[v2 := v1] + v3) + (fm18(0x36e, v4, v1, v1, globalState) + v3);
	var v5: multiset<int> := multiset{v4[safeIndex(-0x1aa, |v4|)]};
	var v6 := 't';
	for i0 := -(|map[|v5| := v6]| * 582) to if (p1) then |v2| else v1 {
		if (p1) {
			p0[safeIndex(666, p0.Length)] := !(p1 <==> p1);
			p0[safeIndex(666, p0.Length)] := p1 || !(|fm19(globalState)| != i0);
			var v7: set<string> := {v2, v2, v2, v2};
			var v8: C0 := new C0(v7);
			var v9: seq<C0> := [v8];
			var v10: multiset<C0> := multiset{v8, v8};
			var v11: C3 := new C3(p0[safeIndex(666, p0.Length)], p1, v2, p1);
			var v12 := DC16(p1, false, v11, v11.f27);
			var v13: map<D7, char> := map[v12 := v6];
			var v14: multiset<bool> := multiset{v11.f27};
			var v15: array<bool> := new bool[17];
			var v16: map<int, array<bool>> := map[v4[safeIndex(v1, |v4|)] := v15];
			var v17 := DC17(v16);
			var v18: map<int, int> := map[v1 := v1];
			var v19: array<int> := new int[20] [v1 + -i0, i0, v1 + i0, i0, fm6(globalState), |(v9 + v9)|, if (v8 in v10) then v10[v8] else v1, i0, v1, v1 - |v13|, 0x274, if (v1 in v5) then v5[v1] else |v14|, |(v16 + v17.cf30)|, i0, v1, v1, safeDivisionInt(v1, i0), if (v1 in v18) then v18[v1] else v1, v1 * |v8.f26|, i0];
			v19[safeIndex(378, v19.Length)] := i0;
			var v20: array<char> := new char[9];
			v20[safeIndex(84, v20.Length)] := v6;
			v19[safeIndex(378, v19.Length)], globalState.f20, v20[safeIndex(84, v20.Length)], p0[safeIndex(666, p0.Length)], p0[safeIndex(666, p0.Length)] := i0 + (if (v11.f28) then v1 else 468), safeModuloInt(i0, |v2|), v6, !((v2 + v2)[safeIndex(v1, |(v2 + v2)|) := v6] <= v2), v11.f27;
			var v21: T0 := new C3(p1, true, v2, p0[safeIndex(666, p0.Length)]);
			var v22: map<T0, bool> := map[v21 := p1];
			var v23: seq<bool> := [v21.f22];
			var v24: map<bool, int> := map[v11.f28 := v1];
			var v25 := DC5(v21.f22, v24, v1, (multiset{0x17b})[v19[safeIndex(378, v19.Length)] := abs(|(seq(abs(0x273), i1  => (v20[safeIndex(84, v20.Length)])))|)]);
			var v26 := DC8(|v23|, map[v14 := |v23|], v25, p1);
			var v27 := DC5(v26.cf16, v24, -v1, multiset{i0, v19[safeIndex(378, v19.Length)]});
			var v28: seq<seq<int>> := [v4];
			var v29: T2 := new C2(v28, seq(abs(0x25c), i2  => (v20[safeIndex(84, v20.Length)])), v11.f28);
			var v30: seq<T2> := [v29, v29];
			var v31: map<bool, seq<T2>> := map[v22 != v22 := if (v27.cf7) then v30 else v30];
			v31 := v31[!v21.f22 := v30 + v30];
			v29 := v29;
			var v32: set<T2> := {v29};
			v32 := v32 + v32;
		} else {
			var v33: multiset<bool> := multiset{p1, false};
			var v34: multiset<multiset<bool>> := multiset{v33, v33, v33, v33, v33};
			var v35: seq<bool> := [fm12(v1, fm17(v1, fm12(v1, v33, p1, v34, globalState), globalState), p1, v34, globalState)];
			r0 := v35[safeIndex(fm6(globalState), |v35|)];
			r0 := v2 == v2;
			var v36: array<multiset<bool>> := new multiset<bool>[12](i3 => v33);
			var v37: C1 := new C1(v36, p1);
			v37 := v37;
			var v38: set<bool> := {v37.f25};
			var v39: set<set<bool>> := {v38, {true, v37.f25}};
			var v40: map<set<set<bool>>, bool> := map[DC18(v39).cf31 := p1];
			v40 := v40[v39 := p1];
			var v41: map<int, char> := map[i0 := v6];
			var v42: seq<multiset<bool>> := [v33, v33];
			r0 := (if (|v42[safeIndex(v1, |v42|)]| in v41) then v41[|v42[safeIndex(v1, |v42|)]|] else v6) in v2;
		}
		
		var v43: map<array<bool>, int> := map[p0 := v1];
		v1 := |v43|;
		v1 := -286;
		p0[safeIndex(286, p0.Length)] := p1;
		var v44: array<bool> := new bool[10];
		var v45: seq<array<bool>> := [p0];
		p0[safeIndex(286, p0.Length)], v44 := (v1 - i0) in v4, if (!true) then p0 else v45[safeIndex(0x201, |v45|)];
	}
	var v46: multiset<bool> := multiset{p1};
	var v47: map<multiset<bool>, int> := map[v46 := fm6(globalState)];
	var v48 := DC5(p1, map[p1 := -v1], v1, v5);
	v4, r0, r0 := match DC8(v1, v47, v48, !p1) {
		case DC8(cf13, cf14, cf15, cf16) => v4
		case DC9(cf17, cf18, cf19) => v4
		case DC7(cf12) => v4
	}, p1, p1;
	var v49: set<int> := {|v0|};
	var v50 := DC6(v49);
	var v51: seq<seq<int>> := [v4[safeIndex(v1, |v4|) := v1], v4, v4];
	var v52: T1 := new C4(v50, p1, v51, v2, |v49| == v1);
	var v53: set<bool> := {false};
	var v54: map<set<bool>, bool> := map[v53 := p1];
	var v56 := DC18(set v55 : set<bool> | v55 in v54 :: (v55));
	var v57: map<bool, bool> := map[true := p1];
	var v58 := DC20(v52.f21, p1, |v57|);
	var v59 := DC20(v58.cf37, true, v1);
	globalState.f3, v52, v46 := match v56 {
		case DC19(cf32, cf33, cf34, cf35, cf36) => v1
		case DC20(cf37, cf38, cf39) => |(v48.cf8 + map[cf38 := v1])|
		case DC21(cf40, cf41) => 0x3e3
		case DC18(cf31) => v1
	}, v52, match v59 {
		case DC19(cf32, cf33, cf34, cf35, cf36) => v46
		case DC20(cf37, cf38, cf39) => v46
		case DC21(cf40, cf41) => multiset([p1, p1] + [cf41, !p1, false, true, cf41])
		case DC18(cf31) => v46 + v46
	};
	v52.f22 := p1;
	var v60: seq<bool> := [v52.f22, v52.f22, v52.f22, v52.f22, true];
	r0 := !!(v52.f22 !in v60) && v52.f22;
}
method Main() {
	var v1 := "truxk";
	var v2: set<string> := {v1, v1};
	var v3 := -0x262;
	var v4 := 'g';
	var v5: array<bool> := new bool[29];
	var v6: array<int> := new int[5];
	var v7: multiset<int> := multiset{v3};
	var globalState := new GlobalState(false, map v0 : string | v0 in v2 :: (v0) := (|[0x119]|), -0x292, 2, -0x2c8, true, (("v")[safeIndex(-v3, |"v"|) := v4])[safeIndex(v3, |("v")[safeIndex(-v3, |"v"|) := v4]|) := v4], true, v1, v1, 0x380, true, 'c', -0x9e, 316, -0x28e, v5, v6, false, v7, 0x134);
	var v8 := false;
	var i0 := 0;
	while (!!v8)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		var v9: array<array<bool>> := new array<bool>[27];
		v9[safeIndex(512, v9.Length)] := v5;
		v3, v9[safeIndex(512, v9.Length)], globalState.f20 := v3 - (if (v8) then v3 else v3), v5, v3;
		globalState.f20 := v3;
		v3 := v3;
		var v10 := m0(v5, v8, globalState);
	}
	if (v8) {
		var v11: array<multiset<bool>> := new multiset<bool>[23](i1 => multiset{true});
		var v12 := new C1(v11, false ==> v8);
		var v13: multiset<bool> := multiset{true};
		var v14: multiset<multiset<bool>> := multiset{v13};
		var v15: set<bool> := {v12.f25, false};
		var v16: C2 := new C2(fm16(v15, v12.f25, v8, globalState), v1, v8);
		var v17: map<bool, C2> := map[fm12(v3, v13, v12.f25, v14[v13 := abs(|v1|)], globalState) := v16];
		var v18: map<string, map<bool, C2>> := map[v1 := v17];
		var v19: map<bool, bool> := map[v8 := false];
		var v20: seq<map<bool, bool>> := [v19, v19, map[v8 := false]];
		globalState.f3, v18, globalState.f3, v8 := v3, v18, safeDivisionInt(v3, v3), v19 !in v20;
		var v21 := m0(v5, !v12.f25, globalState);
		v8 := v12.f25 || v12.f25;
		v3 := if (v16.fm3(globalState)) then v3 else v3;
	} else {
		v6[safeIndex(347, v6.Length)] := -385;
		var v22: seq<int> := [v3];
		v6[safeIndex(347, v6.Length)] := -(v22[safeIndex(v3, |v22|)] * v3);
		var v23: multiset<bool> := multiset{v8, v8};
		var v24: map<int, bool> := map[v6[safeIndex(347, v6.Length)] := v8];
		var v25: multiset<multiset<bool>> := multiset{v23};
		var v26: array<multiset<bool>> := new multiset<bool>[27] [v23, multiset{true}, v23, v23, v23, v23, v23, v23, v23, v23, v23, v23, v23, multiset{v8, v8, v8}, v23, v23, multiset{v8}, multiset{v8, v8}, multiset{v8, v8}, multiset{v8, v8, fm12(v3, multiset{v8, true, if (v3 in v24) then v24[v3] else v8}, !fm12(v3, multiset{v8, v8}, v8, v25, globalState), v25, globalState)}, v23, fm17(v6[safeIndex(347, v6.Length)], v8, globalState), multiset{v8, v8, v8}, v23, v23, multiset{v8}, v23];
		var v27: map<bool, int> := map[fm12(|map[v8 := v3]|, v23, v8, v25, globalState) := v3];
		var v28: C1 := new C1(v26, v6[safeIndex(347, v6.Length)] >= |v27|);
		var v29: seq<bool> := [true, v8];
		var v30 := DC7(multiset(v29));
		v28, v6[safeIndex(347, v6.Length)], v30, v8 := v28, if (v3 >= v3) then |(v2 + v2)| else -0x31a, v30, safeModuloInt(|v22|, v3) >= v3;
		var v31: array<set<bool>> := new set<bool>[24](i2 => {true});
		var v32: array<array<set<bool>>> := new array<set<bool>>[4] [v31, v31, v31, v31];
		v32[safeIndex(134, v32.Length)] := v31;
		v32[safeIndex(134, v32.Length)] := v31;
		var v33: seq<seq<int>> := [v22, fm5(v3, v3, v3, globalState)];
		var v34: C2 := new C2(v33, "m", v8);
		var v35: set<C2> := {v34};
		v35 := {v34, v34, v34, v34, v34};
		var v36 := DC1(v28.f25);
		var v37 := new C1(v28.f24, v36.cf1);
	}
	
	v8 := v8;
	var v38: multiset<bool> := multiset{false};
	var v39: multiset<multiset<bool>> := multiset{v38, v38};
	var v40 := DC0(v5);
	var v41 := DC2(v40);
	match (if (fm12(v3, v38, v8, v39, globalState)) then v41 else v41) {
		case DC1(cf1) =>
			v6[safeIndex(881, v6.Length)] := 0xcb;
			var v42: seq<int> := [v3, 0x22b, v3, v3];
			var v43: seq<bool> := [v8, v8];
			var v44: array<multiset<bool>> := new multiset<bool>[19] [v38, v38, v38, v38, v38, v38, v38, multiset{cf1, cf1, cf1}, v38, fm17(|v42|, v8, globalState), v38, v38, v38, v38, multiset(v43), v38, v38, v38, fm17(v3, true, globalState)];
			var v45: seq<array<multiset<bool>>> := [v44];
			var v46: C1 := new C1(v45[safeIndex(0x268, |v45|)], v8);
			var v47: seq<C1> := [v46, v46, v46];
			v6[safeIndex(881, v6.Length)], globalState.f20 := v3 + |v47|, v3 - -0x312;
			v6[safeIndex(881, v6.Length)] := 0x2d3;
			var v48: array<T0> := new T0[19];
			var v49: T0 := new C3(v46.f25, v8, ("r")[safeIndex(-v3, |"r"|) := v4], false);
			v48[safeIndex(947, v48.Length)] := v49;
			v48[safeIndex(947, v48.Length)] := v49;
			var v50: array<string> := new string[23];
			v50 := v50;
		case DC0(cf0) =>
			var v51: seq<int> := [v3];
			var v52: seq<seq<int>> := [v51];
			var v53: T2 := new C2(v52, v1, if (v8) then v8 else v8);
			v53 := v53;
			var v54: map<string, int> := map[v1 := v3];
			var v56 := new C0(set v55 : string | v55 in v54 :: (v55));
			globalState.f8 := seq(abs(0x8c), i3  => (v4));
			var v57: seq<bool> := [!v53.f22, true, v53.f22];
			if (v57 < v57) {
				v6[safeIndex(657, v6.Length)] := 0x34a;
				v5[safeIndex(483, v5.Length)] := !v8;
				var v58: seq<T2> := [v53];
				var v59: set<map<bool, char>> := {map[true := 'y']};
				var v60: map<set<map<bool, char>>, T2> := map[v59 := v53];
				var v61: C3 := new C3(v53.f22, v8, v53.f21, false);
				var v62: multiset<C3> := multiset{v61};
				var v63 := DC14(v61);
				var v64: array<T2> := new T2[28] [v53, v53, v53, v53, v53, v58[safeIndex(-v3, |v58|)], v53, v53, v53, v53, if (v59 in v60) then v60[v59] else v53, v53, v53, if (v53.f22) then v53 else v53, v53, v53, v53, v53, v58[safeIndex(if (v63.cf24 in v62) then v62[v63.cf24] else v3, |v58|)], v58[safeIndex(|v53.f21|, |v58|)], v53, v53, v53, v53, v53, v53, v53, v53];
				v64[safeIndex(357, v64.Length)] := v53;
				var v65: seq<string> := [v53.f21, v1];
				var v66: map<bool, int> := map[v61.f28 := v3];
				var v67: map<map<bool, int>, int> := map[v66 := v3];
				var v68: array<string> := new string[14] [v65[safeIndex(|v67|, |v65|)], v1, v53.f21, v1, v53.f21, fm11(v61.f27, globalState), v53.f21, v53.f21, v53.f21, v1, fm11(v61.f28, globalState), v53.f21, seq(abs(434), i4  => (v4)), v53.f21];
				var v69: multiset<array<string>> := multiset{v68};
				var v70: set<array<int>> := {v6, v6, v6, v6, v6};
				v6[safeIndex(657, v6.Length)], v5[safeIndex(483, v5.Length)], v64[safeIndex(357, v64.Length)] := if (v68 in v69) then v69[v68] else v3, v70 !! (v70 * {v6}), v53;
				var v71: array<int> := new int[9];
				v71 := v71;
				var v72 := DC5(v8, v66, v3, v7);
				var v73: T0 := new C3(v3 == v3, v61.f28, v53.f21, v3 <= v72.cf9);
				var v74: set<int> := {v3};
				var v75 := DC6(v74);
				var v76: T1 := new C4(v75, v61.f27, v52, v1, v3 < |v57|);
				v53.f21, v73, v76 := fm11(v73.fm1(v1, v5[safeIndex(483, v5.Length)], v61.f27, globalState), globalState), v73, v76;
				var v77: map<int, bool> := map[v6[safeIndex(657, v6.Length)] := v73.f22];
				var v78: map<map<int, bool>, array<bool>> := map[v77 := cf0];
				var v79 := v53.m1(v73.f22, v78, globalState);
				var v80: array<char> := new char[19];
				v80[safeIndex(193, v80.Length)] := v4;
				var v81: set<bool> := {!v61.f27, v8};
				var v82: C4 := new C4(v75, true ==> v8, fm16(v81, v61.f28, v8, globalState), ("lpt")[safeIndex(v3, |"lpt"|) := v4], v53.f22);
				v5[safeIndex(483, v5.Length)], v80[safeIndex(193, v80.Length)], v82 := v82.fm1(v1, !v82.f30, v61.f28, globalState), v4, v82;
			} else {
				v53.f22 := v57[safeIndex(|(if (v53.f22) then v51 else fm5(-fm6(globalState), v3, 109, globalState))|, |v57|)];
				globalState.f20 := v3;
				v8 := v8;
				var v83: array<multiset<bool>> := new multiset<bool>[29];
				var v84: C1 := new C1(v83, false);
				var v85: map<bool, multiset<bool>> := map[v53.f22 := multiset([false])];
				var v86: map<C1, bool> := map[v84 := fm12(v3, if (false in v85) then v85[false] else multiset{v53.f22}, v53.f22, v39, globalState)];
				v86 := v86[v84 := v57[safeIndex(fm6(globalState), |v57|)]];
				v38 := multiset{!v53.f22};
			}
			
		case DC2(cf2) =>
			v6[safeIndex(775, v6.Length)] := v3;
			v6[safeIndex(775, v6.Length)] := v3;
			globalState.f20 := fm6(globalState);
			var v87: seq<int> := [v6[safeIndex(775, v6.Length)], v3];
			globalState.f20, globalState.f3 := v3 + v3, v87[safeIndex(v3, |v87|)];
			var v88: set<int> := {v6[safeIndex(775, v6.Length)]};
			var v89: seq<set<int>> := [v88];
			var v90 := DC6(v89[safeIndex(v3, |v89|)]);
			var v91: seq<seq<int>> := [v87, v87];
			var v92: T1 := new C4(v90.(cf11 := {v6[safeIndex(775, v6.Length)]}), v8, v91, seq(abs(-0x25), i5  => (v4)), fm6(globalState) in v87);
			v8, v92, globalState.f3 := v92.f22, if (v8 ==> v8) then v92 else v92, v3;
	}
	
	var v93: multiset<array<bool>> := multiset{v5, v5};
	var v94: map<int, bool> := map[v3 := v8];
	var v95: seq<bool> := [v8, v8];
	var v96: array<multiset<array<bool>>> := new multiset<array<bool>>[25] [v93, v93, v93, v93 * DC15(v93).cf25[v5 := abs(v3)], v93, multiset{v5, v5}, v93 - v93, v93, v93, v93, v93, v93 + v93, v93, if (if (v3 in v94) then v94[v3] else !v8) then v93 else v93, v93, multiset{v5}, v93, if (fm12(0x23c, multiset{v8, v8, v8, v8, v8}, v95[safeIndex(|v1|, |v95|)], multiset{v38}, globalState)) then v93 else v93, DC15(v93).cf25, v93, v93, v93[v5 := abs(v3)], v93, v93 + v93, multiset{v5} + v93];
	v96[safeIndex(543, v96.Length)] := multiset{v5, v5} + v93;
	v96[safeIndex(543, v96.Length)] := (v93[v5 := abs(|v94|)] + v93) - v93;
	var v97: set<int> := {v3, v3, |"ly"|};
	var v98 := DC6(v97);
	var v99: seq<int> := [v3];
	var v100: C4 := new C4(v98, v8, [[v3], v99, ([v3])[safeIndex(v3, |[v3]|) := v3]], v1, true);
	var v101: set<C4> := {v100, v100, v100};
	var v102: map<bool, set<C4>> := map[fm12(|v1|, multiset{v8}, v8, v39, globalState) := v101];
	for i6 := v3 to -|v102| {
		var v103: seq<seq<int>> := [v99];
		var v104: seq<seq<int>> := [v103[safeIndex(v3, |v103|)]];
		var v105 := new C2(v103, v1 + v1[safeIndex(i6, |v1|) := 'u'], v3 >= v3);
		var v106: seq<map<int, bool>> := [v94];
		var v107 := v105.m1(v100.f30, map[v106[safeIndex(840, |v106|)] := v5], globalState);
		v8 := v100.f30;
		v99 := v99;
	}
	v8 := v100.f30;
	var v108: array<char> := new char[8] [v4, if (v8) then v4 else v4, v4, v4, v4, v4, v4, v4];
	v108 := v108;
	v99 := v99;
	var v109 := DC7(v38);
	var v110: array<multiset<bool>> := new multiset<bool>[25] [multiset{false}, fm17(v3, v100.f30, globalState), v38, v38, v38, (v109.(cf12 := v38)).cf12, v38, v38, v38, multiset{v8, v100.f30}, v38, v38, v38, multiset(v95[safeIndex(v3, |v95|) := v100.fm1(v1, v8, v100.f30, globalState)]), multiset{!v100.f30}, v38, v38, v38, multiset{v100.f30}, v38, v38[v8 := abs(-v3)], v38, v38, fm17(0x379, v8, globalState), multiset(v95)];
	var v111 := new C1(v110, v100.f30);
	var v112: map<int, int> := map[v3 := v3];
	var v113: map<multiset<bool>, int> := map[v38 := v3];
	var v114 := DC5(false, map[v111.f25 := v3], v3, v7);
	v3 := safeModuloInt(|v112|, DC8(v3, v113, v114, v8).cf13);
	for i7 := safeDivisionInt(0x56, v3) to v3 {
		var v115: seq<string> := [fm11(v111.f25, globalState), v1, fm11(false, globalState)];
		v1 := v115[safeIndex(v3, |v115|)] + v1;
		globalState.f20 := 0x2f9;
		v108[safeIndex(578, v108.Length)] := v4;
		var v116: multiset<string> := multiset{v1};
		v108[safeIndex(578, v108.Length)] := v1[safeIndex(|(multiset{v1} * v116)|, |v1|)];
		v3 := |v38[v111.f25 := abs(-(v3 + 0x3d8))]|;
	}
	var v117 := new C3(v111.f25, v8, v1, v111.f25);
	v3 := v3;
	globalState.f12 := if (v117.f27) then fm4(v117.f28, v38, globalState) else 'q';
	v117.f28 := v111.f25;
	print v1, "\n";
	print v2 == {"truxk"}, "\n";
	print v3, "\n";
	print v4, "\n";
	print v6[0], "\n";
	print v6[2], "\n";
	print v7 == multiset{-610}, "\n";
	print globalState.f0, "\n";
	print globalState.f1 == map["truxk" := 1], "\n";
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
	print globalState.f14, "\n";
	print globalState.f15, "\n";
	print globalState.f17[0], "\n";
	print globalState.f17[2], "\n";
	print globalState.f18, "\n";
	print globalState.f19 == multiset{-610}, "\n";
	print globalState.f20, "\n";
	print v8, "\n";
	print i0, "\n";
	print v38 == multiset{false}, "\n";
	print v39 == multiset{multiset{false}, multiset{false}}, "\n";
	print |v93|, "\n";
	print v94 == map[-610 := false], "\n";
	print v95 == [false, false], "\n";
	print |v96[0]|, "\n";
	print |v96[1]|, "\n";
	print |v96[2]|, "\n";
	print |v96[3]|, "\n";
	print |v96[4]|, "\n";
	print |v96[5]|, "\n";
	print |v96[6]|, "\n";
	print |v96[7]|, "\n";
	print |v96[8]|, "\n";
	print |v96[9]|, "\n";
	print |v96[10]|, "\n";
	print |v96[11]|, "\n";
	print |v96[12]|, "\n";
	print |v96[13]|, "\n";
	print |v96[14]|, "\n";
	print |v96[15]|, "\n";
	print |v96[16]|, "\n";
	print |v96[17]|, "\n";
	print |v96[18]|, "\n";
	print |v96[19]|, "\n";
	print |v96[20]|, "\n";
	print |v96[21]|, "\n";
	print |v96[22]|, "\n";
	print |v96[23]|, "\n";
	print |v96[24]|, "\n";
	print v97 == {-610, 2}, "\n";
	print v98.cf11 == {-610, 2}, "\n";
	print v99 == [-610], "\n";
	print v100.f29.cf11 == {-610, 2}, "\n";
	print v100.f30, "\n";
	print |v101|, "\n";
	print |v102|, "\n";
	print v108[0], "\n";
	print v108[1], "\n";
	print v108[2], "\n";
	print v108[3], "\n";
	print v108[4], "\n";
	print v108[5], "\n";
	print v108[6], "\n";
	print v108[7], "\n";
	print v109.cf12 == multiset{false}, "\n";
	print v110[0] == multiset{false}, "\n";
	print v110[1] == multiset{true}, "\n";
	print v110[2] == multiset{false}, "\n";
	print v110[3] == multiset{false}, "\n";
	print v110[4] == multiset{false}, "\n";
	print v110[5] == multiset{false}, "\n";
	print v110[6] == multiset{false}, "\n";
	print v110[7] == multiset{false}, "\n";
	print v110[8] == multiset{false}, "\n";
	print v110[9] == multiset{false, false}, "\n";
	print v110[10] == multiset{false}, "\n";
	print v110[11] == multiset{false}, "\n";
	print v110[12] == multiset{false}, "\n";
	print v110[13] == multiset{true, false}, "\n";
	print v110[14] == multiset{true}, "\n";
	print v110[15] == multiset{false}, "\n";
	print v110[16] == multiset{false}, "\n";
	print v110[17] == multiset{false}, "\n";
	print v110[18] == multiset{false}, "\n";
	print v110[19] == multiset{false}, "\n";
	print v110[20] == multiset{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}, "\n";
	print v110[21] == multiset{false}, "\n";
	print v110[22] == multiset{false}, "\n";
	print v110[23] == multiset{true}, "\n";
	print v110[24] == multiset{false, false}, "\n";
	print v111.f24[0] == multiset{false}, "\n";
	print v111.f24[1] == multiset{true}, "\n";
	print v111.f24[2] == multiset{false}, "\n";
	print v111.f24[3] == multiset{false}, "\n";
	print v111.f24[4] == multiset{false}, "\n";
	print v111.f24[5] == multiset{false}, "\n";
	print v111.f24[6] == multiset{false}, "\n";
	print v111.f24[7] == multiset{false}, "\n";
	print v111.f24[8] == multiset{false}, "\n";
	print v111.f24[9] == multiset{false, false}, "\n";
	print v111.f24[10] == multiset{false}, "\n";
	print v111.f24[11] == multiset{false}, "\n";
	print v111.f24[12] == multiset{false}, "\n";
	print v111.f24[13] == multiset{true, false}, "\n";
	print v111.f24[14] == multiset{true}, "\n";
	print v111.f24[15] == multiset{false}, "\n";
	print v111.f24[16] == multiset{false}, "\n";
	print v111.f24[17] == multiset{false}, "\n";
	print v111.f24[18] == multiset{false}, "\n";
	print v111.f24[19] == multiset{false}, "\n";
	print v111.f24[20] == multiset{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}, "\n";
	print v111.f24[21] == multiset{false}, "\n";
	print v111.f24[22] == multiset{false}, "\n";
	print v111.f24[23] == multiset{true}, "\n";
	print v111.f24[24] == multiset{false, false}, "\n";
	print v111.f25, "\n";
	print v112 == map[-610 := -610], "\n";
	print v113 == map[multiset{false} := -610], "\n";
	print v114.cf7, "\n";
	print v114.cf8 == map[false := -610], "\n";
	print v114.cf9, "\n";
	print v114.cf10 == multiset{-610}, "\n";
	print v117.f27, "\n";
	print v117.f28, "\n";
}