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
datatype D0 = DC0(cf0: int)
datatype D1 = DC2(cf2: bool) | DC1(cf1: bool)
datatype D2 = DC4(cf4: int, cf5: bool) | DC5(cf6: int, cf7: bool, cf8: bool, cf9: bool) | DC3(cf3: seq<map<int, int>>) | DC6(cf10: D2)
datatype D3 = DC8(cf12: bool) | DC9(cf13: int, cf14: char, cf15: int) | DC7(cf11: map<int, bool>) | DC10(cf16: D3)
datatype D4 = DC12(cf18: bool, cf19: array<int>, cf20: int) | DC13(cf21: array<C0>, cf22: bool, cf23: bool) | DC11(cf17: string)
datatype D5 = DC15(cf25: bool, cf26: array<bool>, cf27: string, cf28: bool) | DC14(cf24: map<int, char>) | DC16(cf29: D5)
datatype D6 = DC18(cf31: int, cf32: bool, cf33: bool, cf34: bool, cf35: int) | DC17(cf30: map<int, int>) | DC19(cf36: D6)
datatype D7 = DC21(cf38: int, cf39: map<bool, bool>) | DC22(cf40: int, cf41: D0, cf42: bool) | DC23(cf43: int, cf44: int, cf45: bool, cf46: int) | DC20(cf37: seq<bool>) | DC24(cf47: D7)
datatype D8 = DC26(cf49: array<int>, cf50: bool, cf51: char, cf52: bool, cf53: bool) | DC25(cf48: multiset<int>)
trait T0 {
	const f25 : bool
	method m1(p0: string, p1: int, p2: seq<bool>, globalState: GlobalState) returns (r0: bool, r1: int) 
}

trait T1 extends T0 {
	const f26 : char
	const f27 : array<multiset<int>>
	function fm3(p0: bool, p1: int, p2: int, p3: map<char, int>, globalState: GlobalState): bool 
	method m2(p0: set<char>, p1: bool, p2: int, globalState: GlobalState) 
}

class GlobalState {
	const f0 : int
	const f1 : int
	var f2 : bool
	const f3 : string
	const f4 : array<set<int>>
	var f5 : int
	const f6 : bool
	const f7 : int
	const f8 : bool
	const f9 : array<array<bool>>
	const f10 : int
	const f11 : bool
	const f12 : char
	const f13 : bool
	const f14 : bool
	const f15 : int
	const f16 : int
	const f17 : bool
	var f18 : int
	var f19 : seq<int>
	const f20 : int
	var f21 : int
	var f22 : int
	const f23 : array<bool>
	const f24 : char
	constructor (f0 : int, f1 : int, f2 : bool, f3 : string, f4 : array<set<int>>, f5 : int, f6 : bool, f7 : int, f8 : bool, f9 : array<array<bool>>, f10 : int, f11 : bool, f12 : char, f13 : bool, f14 : bool, f15 : int, f16 : int, f17 : bool, f18 : int, f19 : seq<int>, f20 : int, f21 : int, f22 : int, f23 : array<bool>, f24 : char) {
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
	}
	
}

class C0 {
	constructor () {
	}
	
	function fm5(globalState: GlobalState): int {
		-((if (true) then -144 else -|map[0x262 := true]|) + -131)
	}
	function fm6(p0: seq<int>, p1: bool, globalState: GlobalState): int {
		|[false]| + |("roa" + "w")|
	}
}

class C1 extends T0 {
	var f28 : string
	var f29 : int
	constructor (f28 : string, f29 : int, f25 : bool) {
		this.f28 := f28;
		this.f29 := f29;
		this.f25 := f25;
	}
	
	function fm14(p0: bool, p1: int, globalState: GlobalState): int {
		safeModuloInt(-|(multiset{f25, f25, f25, false} + multiset{f25})|, f29)
	}
	function fm15(globalState: GlobalState): seq<int> {
		seq(abs(0x4), i0  => (|([f29] + [|(set v0 : int | v0 in multiset{f29} :: (safeModuloInt(v0, |map[-186 := |(map v1 : seq<bool> | v1 in map[[false] := 'j'] :: (v1) := (873))|]|)))|])|))
	}
	method m1(p0: string, p1: int, p2: seq<bool>, globalState: GlobalState) returns (r0: bool, r1: int) {
		var v0 := DC7(map[f29 := f25]);
		var v2: map<int, bool> := map[p1 := f25];
		var v3: map<bool, bool> := map[f25 := false];
		var v4: array<bool> := new bool[14] [f25, f25, f25, f25, f25, false, f25, f25, f25, p2[safeIndex(p1, |p2|)], if (f25 in v3) then v3[f25] else f25, f25, f25, f25];
		var v5: map<array<bool>, map<int, bool>> := map[v4 := v2];
		var v7: map<int, int> := map[0x10e := p1];
		var v8: seq<map<int, bool>> := [map v6 : int | v6 in v7 :: (safeDivisionInt(v6, f29)) := (f25)];
		var v9: array<map<int, bool>> := new map<int, bool>[21] [v0.cf11, map v1 : int | (0x94 <= v1) && (v1 < 0x22d) :: (v1 - f29) := (f25), v2 + (if (v4 in v5) then v5[v4] else v2), v2, v2, if (true) then v2 else v2, v8[safeIndex(-0x212, |v8|)], v2, v2, v2, v2, v2, map[|v3| := f25], v2, v2, v2 + map[f29 := f25], v2, v2 + v2, v2, v2, v2];
		forall i0 | 0 <= i0 < v9.Length {
			v9[i0] := v2;
		}
		if (f25) {
			var v10: seq<int> := [|f28| + f29];
			var v11 := 'f';
			var v12: multiset<char> := multiset{v11, v11};
			f29 := v10[safeIndex(|v12|, |v10|)];
			f28 := ((seq(abs(0x6b), i1  => (v11))) + [v11])[safeIndex(safeDivisionInt(p1, f29), |((seq(abs(0x6b), i1  => (v11))) + [v11])|) := 'c'];
			globalState.f2 := p2[safeIndex(p1, |p2|)];
			var v13 := new C0();
			v11 := fm16(f29, globalState);
		} else {
			f28 := seq(abs(0x212), i2  => ('r'));
			var v14 := 'y';
			var v15: set<int> := {f29, |v7|};
			var v16: seq<set<int>> := [v15];
			var v17: seq<int> := [fm2(f25, |[f25]|, if (|f28| in v2) then v2[|f28|] else f25, v14, globalState), |v16[safeIndex(p1, |v16|)]|, f29];
			globalState.f19 := v17;
			v4[safeIndex(232, v4.Length)] := f25;
			v4[safeIndex(232, v4.Length)] := f25;
			var v18 := new C0();
			var v19 := DC9(p1, v14, f29);
			var v20: map<D3, bool> := map[v19 := f25];
			v20 := v20[v19 := v4[safeIndex(232, v4.Length)]];
		}
		
		var v21: array<int> := new int[7](i3 => safeModuloInt(i3, |multiset{f25, f25}|));
		v21[safeIndex(394, v21.Length)] := safeDivisionInt(0xda, p1);
		v21[safeIndex(394, v21.Length)] := p1 + -0x31e;
		v21[safeIndex(394, v21.Length)] := f29;
		var v22: array<string> := new string[22];
		v22 := v22;
		forall i4 | 0 <= i4 < v4.Length {
			v4[i4] := f25;
		}
		r0 := f25;
		r1 := f29;
	}
	method m3(globalState: GlobalState) {
		var v0: array<int> := new int[4](i0 => i0 - f29);
		v0[safeIndex(658, v0.Length)] := f29;
		var v1: seq<bool> := [f25, f25];
		v0[safeIndex(658, v0.Length)] := safeDivisionInt(f29, |v1| + f29);
		var v2 := new C0();
		var v3: multiset<string> := multiset{f28};
		var v4: map<int, bool> := map[v0[safeIndex(658, v0.Length)] := f25];
		var v5: array<bool> := new bool[19] [f25, f25, multiset{f28, f28} <= v3, f29 > 0x2b, fm0(f29, f29, f25, globalState), true, fm0(-v0[safeIndex(658, v0.Length)], |v1|, !f25, globalState), f25, f25, f25, true, f28 == f28, [f25] < v1, v1[safeIndex(v0[safeIndex(658, v0.Length)], |v1|)], f25, f25, if (0x12b in v4) then v4[0x12b] else false, f25, fm0(-0x175, |f28|, !f25, globalState)];
		v5[safeIndex(448, v5.Length)] := fm0(f29, f29, f25, globalState);
		var v6 := 'i';
		var v7: map<int, char> := map[v0[safeIndex(658, v0.Length)] := v6];
		var v8: seq<int> := [|v7|];
		var v9 := DC9(f29, v6, |v8|);
		v5[safeIndex(448, v5.Length)] := match v9 {
			case DC8(cf12) => f25
			case DC9(cf13, cf14, cf15) => v0[safeIndex(658, v0.Length)] <= cf15
			case DC7(cf11) => f25 && f25
			case DC10(cf16) => f25
		};
		globalState.f21 := |fm17(f25, v0[safeIndex(658, v0.Length)] - 0x12a, globalState)|;
		var v10: map<bool, int> := map[true := 778];
		var v11: seq<map<bool, int>> := [v10, (v10[v5[safeIndex(448, v5.Length)] := v0[safeIndex(658, v0.Length)]])[true := 48]];
		for i1 := v2.fm6(v8, f25, globalState) to |v11| {
			var v12: seq<C0> := [v2, v2, v2, v2];
			var v13: array<C0> := new C0[17] [v2, if (v5[safeIndex(448, v5.Length)]) then v2 else v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v12[safeIndex(f29, |v12|)], v2, v2, v2];
			v13[safeIndex(661, v13.Length)] := v12[safeIndex(v0[safeIndex(658, v0.Length)], |v12|)];
			var v14: array<array<bool>> := new array<bool>[9];
			v14[safeIndex(401, v14.Length)] := v5;
			var v15 := DC1(v5[safeIndex(448, v5.Length)]);
			var v16: seq<D1> := [v15];
			var v17 := DC11(fm18(|v16|, |[v5[safeIndex(448, v5.Length)], false, v5[safeIndex(448, v5.Length)], f25, v5[safeIndex(448, v5.Length)]]|, globalState));
			f28, v10, v13[safeIndex(661, v13.Length)], v14[safeIndex(401, v14.Length)] := v17.cf17, v10 + v10, v2, v5;
			var v18: array<D3> := new D3[26](i2 => DC8(f25));
			var v19 := DC8(v5[safeIndex(448, v5.Length)]);
			v18[safeIndex(131, v18.Length)] := v19;
			f29, v0, v6, v18[safeIndex(131, v18.Length)] := |f28| * f29, if (f25) then v0 else if (f25) then v0 else v0, v6, v19;
			globalState.f2 := true;
			globalState.f2, v0[safeIndex(658, v0.Length)], v9, v0[safeIndex(658, v0.Length)], globalState.f5 := true, i1 * i1, v9, i1 + v0[safeIndex(658, v0.Length)], if (f25) then f29 else v0[safeIndex(658, v0.Length)];
		}
		if (fm0(-v0[safeIndex(658, v0.Length)], v0[safeIndex(658, v0.Length)], f25, globalState)) {
			v6 := 'w';
			if (if (v5[safeIndex(448, v5.Length)]) then v5[safeIndex(448, v5.Length)] else false) {
				var v20: multiset<bool> := multiset{v5[safeIndex(448, v5.Length)], true};
				var v21: map<bool, multiset<bool>> := map[v5[safeIndex(448, v5.Length)] := v20];
				v21 := v21;
				v0[safeIndex(658, v0.Length)] := -f29;
				v0[safeIndex(658, v0.Length)] := -0x164;
				v5[safeIndex(448, v5.Length)] := f28[safeIndex(f29, |f28|) := v6] <= f28;
				globalState.f2 := v5[safeIndex(448, v5.Length)];
			} else {
				var v22: map<bool, string> := map[fm0(v0[safeIndex(658, v0.Length)], v0[safeIndex(658, v0.Length)], v5[safeIndex(448, v5.Length)], globalState) := (DC11(f28).(cf17 := f28)).cf17];
				var v23: seq<string> := [f28, f28[safeIndex(v0[safeIndex(658, v0.Length)], |f28|) := v6], f28, if (f25 in v22) then v22[f25] else "idiijp"];
				var v24: array<array<int>> := new array<int>[23];
				var v25: map<string, array<array<int>>> := map[v23[safeIndex(f29, |v23|)] := v24];
				v25 := v25[f28 := v24];
				globalState.f18 := -0x1ee * (if (!false) then v0[safeIndex(658, v0.Length)] else v0[safeIndex(658, v0.Length)]);
				globalState.f2 := v0[safeIndex(658, v0.Length)] <= f29;
				v0[safeIndex(658, v0.Length)] := v0[safeIndex(658, v0.Length)];
				var v26 := new C0();
			}
			
			globalState.f2 := v5[safeIndex(448, v5.Length)] || f25;
			f28 := f28;
			globalState.f5 := f29;
		} else {
			var v27 := DC11("qptks");
			var v28: map<int, D4> := map[f29 := v27];
			v28 := v28[v0[safeIndex(658, v0.Length)] := v27];
			var v29: set<int> := {v0[safeIndex(658, v0.Length)] * v0[safeIndex(658, v0.Length)], safeDivisionInt(f29, v0[safeIndex(658, v0.Length)])};
			v29 := v29 + v29;
			var v30 := DC4(f29, f25);
			match (v30) {
				case DC4(cf4, cf5) =>
					var v31: map<seq<int>, bool> := map[(seq(abs(-0x2c2), i3  => (v0[safeIndex(658, v0.Length)]))) + v8 := cf5];
					v31 := v31[v8 := cf5];
					var v32: map<int, multiset<string>> := map[v0[safeIndex(658, v0.Length)] := if (true) then v3 else v3];
					v32 := v32[|v8| := v3];
					var v33: set<bool> := {true, v5[safeIndex(448, v5.Length)]};
					var v34: map<int, int> := map[v0[safeIndex(658, v0.Length)] := |v33|];
					var v35: set<char> := {'t'};
					var v36: multiset<int> := multiset{|v35|, 0x131};
					var v37: map<set<int>, string> := map[if (v5[safeIndex(448, v5.Length)]) then v29 else v29 := fm18(v0[safeIndex(658, v0.Length)], if (f29 in v34) then v34[f29] else if (v0[safeIndex(658, v0.Length)] in v36) then v36[v0[safeIndex(658, v0.Length)]] else cf4, globalState) + (seq(abs(-0x6e), i4  => (v6)))];
					v37 := v37;
					globalState.f5, cf5, globalState.f18, globalState.f18, v5 := cf4, v5[safeIndex(448, v5.Length)], -f29, f29, v5;
				case DC5(cf6, cf7, cf8, cf9) =>
					v5[safeIndex(448, v5.Length)] := cf6 > 0x113;
					globalState.f22 := |fm1(v0[safeIndex(658, v0.Length)], cf6, globalState)|;
					var v38: multiset<int> := multiset{f29, f29};
					var v39 := DC0(|v38|);
					var v42: seq<map<int, char>> := [v7];
					var v43: array<map<int, char>> := new map<int, char>[26] [v7, fm19(|f28|, true, globalState) + v7, map[|v1| := 'i'], map[f29 := f28[safeIndex(f29, |f28|)]], v7[v39.cf0 := 'e'], v7 + v7, v7, v7, v7, v7[f29 := v6], fm19(f29, cf7, globalState), map[v0[safeIndex(658, v0.Length)] := f28[safeIndex(v0[safeIndex(658, v0.Length)], |f28|)]], map[f29 := v6], v7 + v7, v7 + v7, v7 + v7, v7, DC14(map[f29 := v6]).cf24 + v7, map v40 : int | v40 in v8 :: (v40 * f29) := (v6), v7, v7, (map v41 : int | (0x129 <= v41) && (v41 < 0x2a3) :: (safeDivisionInt(v41, f29)) := ('s')) + map[402 := v6], v7, v42[safeIndex(v0[safeIndex(658, v0.Length)], |v42|)], v7, v7];
					v43[safeIndex(581, v43.Length)] := v7;
					v43[safeIndex(581, v43.Length)] := (v7 + v7) + v42[safeIndex(cf6, |v42|)];
					globalState.f22 := safeModuloInt(|{cf6}|, f29) * (0x14b + f29);
				case DC3(cf3) =>
					globalState.f18 := f29;
					f28 := ("ccueaym" + f28) + (seq(abs(248), i5  => (v6)));
					var v44 := DC4(f29, f25);
					var v45 := DC6(v44);
					v45 := fm20(v0[safeIndex(658, v0.Length)], globalState);
					globalState.f18 := f29;
				case DC6(cf10) =>
					var v46: multiset<bool> := multiset{v5[safeIndex(448, v5.Length)]};
					globalState.f2 := fm0(|v46|, v0[safeIndex(658, v0.Length)], f25, globalState);
					v0 := if (!f25) then v0 else v0;
					f29 := v2.fm6(v8 + (seq(abs(792), i6  => (|multiset{0x372}|))), f25, globalState);
					v8 := fm15(globalState);
			}
			
			globalState.f2 := f25;
			var v47: map<int, int> := map[|v1| := f29];
			var v50 := DC17(fm21(globalState));
			var v52: seq<map<int, int>> := [fm21(globalState), v47, v47];
			var v53: array<map<int, int>> := new map<int, int>[24] [v47 + v47, map[f29 := 159], (map v48 : int | (155 <= v48) && (v48 < 0x154) :: (v48 + v0[safeIndex(658, v0.Length)]) := (|map[v1[safeIndex(f29, |v1|)] := true]|))[f29 := f29] + v47, (map v49 : int | (222 <= v49) && (v49 < 0x16d) :: (safeModuloInt(v49, v0[safeIndex(658, v0.Length)])) := (v0[safeIndex(658, v0.Length)])) + v47, v47, fm21(globalState), v47[v0[safeIndex(658, v0.Length)] := v2.fm5(globalState)], v47[v0[safeIndex(658, v0.Length)] := v0[safeIndex(658, v0.Length)]], v47[|fm18(v0[safeIndex(658, v0.Length)], v0[safeIndex(658, v0.Length)], globalState)| := f29], v47, v47, v47, v47, v50.cf30 + map[-f29 := -f29], v47, v47, v47, map v51 : int | (0x1cf <= v51) && (v51 < 0x1db) :: (v51 * v0[safeIndex(658, v0.Length)]) := (0x2ad), v52[safeIndex(-f29, |v52|)], v47 + v47, v47, v47, v47, v47];
			v53 := v53;
		}
		
	}
}

class C2 extends T1 {
	constructor (f26 : char, f27 : array<multiset<int>>, f25 : bool) {
		this.f26 := f26;
		this.f27 := f27;
		this.f25 := f25;
	}
	
	function fm3(p0: bool, p1: int, p2: int, p3: map<char, int>, globalState: GlobalState): bool {
		{195} !! {|map[f25 := map[f25 := 0x25a]]|}
	}
	function fm8(p0: bool, globalState: GlobalState): string {
		("w" + "goffav") + "atkuqwah"
	}
	function fm9(p0: int, p1: string, p2: bool, globalState: GlobalState): D0 {
		DC0(safeModuloInt(681, -|map[f25 := true]|))
	}
	method m2(p0: set<char>, p1: bool, p2: int, globalState: GlobalState) {
		var v0: map<int, bool> := map[p2 := false];
		globalState.f5 := |v0|;
		var v1 := "w";
		var v2: seq<int> := [|v1|];
		for i0 := |(v2 + v2)| to p2 {
			globalState.f2 := !true;
			var v3: set<bool> := {p1};
			var v4 := DC4(|v3|, p1);
			var v5 := DC6(v4);
			match (v5) {
				case DC4(cf4, cf5) =>
					var v6: array<array<D2>> := new array<D2>[26];
					var v7: map<bool, array<array<D2>>> := map[cf5 := v6];
					v7 := v7[cf5 := v6];
					var v8: map<bool, bool> := map[cf5 := f25];
					var v9: map<map<bool, bool>, bool> := map[v8 := p1];
					v9 := v9;
					cf5 := false;
					v8 := v8;
				case DC5(cf6, cf7, cf8, cf9) =>
					globalState.f2 := if (fm10(globalState) in v1) then cf9 else f25;
					var v10 := DC0(p2 * p2);
					v10 := v10;
					globalState.f19 := v2 + v2;
					var v11: array<bool> := new bool[2](i1 => f25);
					var v12: seq<bool> := [p1, cf7];
					var v13: map<bool, int> := map[p1 := |v1|];
					var v14: multiset<int> := multiset{-0x238, cf6, i0, |v13|};
					v11[safeIndex(889, v11.Length)] := !v12[safeIndex(|(set v15 : int | v15 in v14 :: (safeModuloInt(v15, |multiset{|(set v16 : int | (0x134 <= v16) && (v16 < 0x131) :: (safeModuloInt(v16, |{|[-0x3d6]|}|)))|}|)))|, |v12|)] <==> f25;
					var v17: array<set<bool>> := new set<bool>[21](i2 => v3);
					v11[safeIndex(889, v11.Length)], globalState.f2, v17 := p1, true, v17;
				case DC3(cf3) =>
					var v18: C0 := new C0();
					v18 := if (p1) then v18 else v18;
					globalState.f22 := if (p1) then i0 else i0;
					var v19 := m0(globalState);
					globalState.f2 := !f25;
				case DC6(cf10) =>
					var v20: map<seq<map<int, bool>>, bool> := map[seq(abs(0x2ad), i3  => (v0)) := if (p1) then f25 else f25];
					var v21: seq<map<int, bool>> := [v0, v0];
					var v22 := DC5(i0, p1, f25, f25);
					v20 := v20[v21 := v22.cf9];
					globalState.f22 := i0;
					globalState.f2 := p1;
					var v23: array<bool> := new bool[6](i4 => 444 > |v3|);
					v23[safeIndex(202, v23.Length)] := f25;
					v23[safeIndex(202, v23.Length)] := f25;
			}
			
			var v24 := 'y';
			v24 := fm10(globalState);
			var v25: array<int> := new int[25];
			v25 := new int[1];
		}
		var i5 := 0;
		while (f25)
			decreases 100 - i5
		{
			if (i5 >= 100) {
				break;
			}
			
			i5 := i5 + 1;
			var v26 := DC4(p2, p1);
			var v27: seq<D2> := [v26, fm11(globalState), v26];
			var v28: map<char, int> := map[f26 := p2];
			var v29: array<bool> := new bool[20] [p1, f25, p1, f25, !p1, p1, f25, f25 <== p1, p1 <== p1, f25, fm3(f25, p2, |v1|, v28, globalState), !false, fm3(p1, p2, p2, v28, globalState), f25, f25, p1, !true, !f25, f25, true];
			v29[safeIndex(207, v29.Length)] := p1;
			var v30: map<bool, int> := map[p1 := p2];
			var v31: set<bool> := {p1};
			var v32: multiset<int> := multiset{p2, p2, safeModuloInt(|v31|, 0x1f6), safeDivisionInt(p2, p2)};
			v27, globalState.f22, globalState.f18, v29[safeIndex(207, v29.Length)], v30 := v27 + (v27 + v27), 387, |v32|, if (true) then f25 else false, (v30 + map[true := p2]) + (map[f25 := p2] + v30);
			var v33: array<array<int>> := new array<int>[2];
			v33 := v33;
			v29, globalState.f2 := if (p1) then v29 else v29, !fm0(fm2(v29[safeIndex(207, v29.Length)], 0x3e2, !p1, f26, globalState), safeDivisionInt(-823, p2), v1 <= v1, globalState);
			if (p1) {
				globalState.f2 := f25;
				var v34: array<int> := new int[14](i6 => i6 * -782);
				v34[safeIndex(119, v34.Length)] := p2;
				v34[safeIndex(119, v34.Length)] := safeModuloInt(p2, p2) * p2;
				var v35: seq<bool> := [v29[safeIndex(207, v29.Length)]];
				v34[safeIndex(119, v34.Length)] := v34[safeIndex(119, v34.Length)] * --|([p1, p1, v29[safeIndex(207, v29.Length)]] + [v35[safeIndex(|v35|, |v35|)]])|;
				v29[safeIndex(207, v29.Length)] := false;
				v29[safeIndex(207, v29.Length)] := f25;
			} else {
				globalState.f2 := -0x24 <= p2;
				var v36: array<int> := new int[5];
				v36[safeIndex(284, v36.Length)] := p2;
				var v37: seq<array<bool>> := [v29, v29];
				var v38: multiset<bool> := multiset{f25};
				var v39: multiset<array<bool>> := multiset{v29, if (f25) then v29 else v29, v29, v37[safeIndex(|v38|, |v37|)]};
				globalState.f18, globalState.f5, v36[safeIndex(284, v36.Length)] := |v39|, safeModuloInt(0x32b, 0x159) - safeDivisionInt(|v38|, |v1|), safeDivisionInt(884, safeModuloInt(p2, p2));
				globalState.f22 := p2;
				globalState.f18 := -safeDivisionInt(-p2, p2);
				v29[safeIndex(207, v29.Length)] := f25;
			}
			
		}
		var v40: array<bool> := new bool[18];
		v40[safeIndex(54, v40.Length)] := !(if (p2 in v0) then v0[p2] else !false);
		v40[safeIndex(54, v40.Length)] := f25;
		var v41: array<map<bool, multiset<int>>> := new map<bool, multiset<int>>[11](i7 => map[p1 := multiset{p2, p2}]);
		var v42: multiset<bool> := multiset{f25};
		var v43: multiset<int> := multiset{|v42|};
		var v44: map<bool, multiset<int>> := map[p1 := v43];
		v41[safeIndex(879, v41.Length)] := v44;
		var v45 := DC4(p2, !f25);
		var v46: C0 := new C0();
		var v47: seq<C0> := [v46, v46];
		var v48: set<C0> := {v47[safeIndex(p2, |v47|)]};
		v2, v41[safeIndex(879, v41.Length)], v45 := [p2, p2, |v48|] + (v2 + v2), v44 + v44, DC4(p2 - -0x34d, p1);
		var v49: seq<seq<int>> := [v2];
		for i8 := |v49| to |v42| {
			v1 := (v1 + v1) + (v1 + v1);
			globalState.f5 := i8;
			var v50: map<array<bool>, int> := map[v40 := i8];
			var v51: set<int> := {|v50|};
			if (i8 !in v51) {
				var v52 := DC2(p1);
				v40, globalState.f5, globalState.f2, v52, globalState.f22 := v40, -safeDivisionInt(0x23c * 896, i8), v40[safeIndex(54, v40.Length)], v52, v46.fm5(globalState);
				globalState.f18 := |v2[safeIndex(p2, |v2|) := |v2|]|;
				var v53: map<array<bool>, bool> := map[v40 := v40[safeIndex(54, v40.Length)]];
				var v54: array<map<array<bool>, bool>> := new map<array<bool>, bool>[3] [v53, v53[v40 := v40[safeIndex(54, v40.Length)]], v53];
				v54[safeIndex(767, v54.Length)] := map[v40 := true];
				var v55: map<bool, map<array<bool>, bool>> := map[v40[safeIndex(54, v40.Length)] := v53];
				v54[safeIndex(767, v54.Length)] := (if (v40[safeIndex(54, v40.Length)] in v55) then v55[v40[safeIndex(54, v40.Length)]] else v53) + v53;
				var v56: map<multiset<int>, int> := map[v43 := p2 + |v1|];
				v56 := if (p1) then v56 else v56;
				var v57 := new C0();
			} else {
				globalState.f21 := |v1|;
				globalState.f5 := -32;
				var v58: array<multiset<string>> := new multiset<string>[20](i9 => multiset{v1, v1} + multiset{v1, seq(abs(0x1b9), i10  => (f26))});
				var v59: multiset<string> := multiset{v1, v1};
				v58[safeIndex(97, v58.Length)] := v59;
				v58[safeIndex(97, v58.Length)] := v59[v1 := abs(p2)];
				var v60: array<seq<int>> := new seq<int>[23] [v2, v2, v2, v2, v2, v2, v2 + v2, if (fm0(i8, i8, p1, globalState)) then v2 else seq(abs(-463), i11  => (p2)), v49[safeIndex(|v1|, |v49|)], fm12(p2, f25, v0[i8 := v40[safeIndex(54, v40.Length)]], p2, globalState), v2 + v2, v2, [i8, p2] + v2, v2, [|multiset(v49)|, p2, -p2], [p2] + v2, v2, [|fm13(p1, p2, globalState)|, 0x64, p2, p2], v2, v2 + (seq(abs(-853), i12  => (i8))), seq(abs(237), i13  => (p2)), v2 + v2, fm12(p2, !f25, v0, p2, globalState)];
				var v61: map<int, seq<int>> := map[p2 := [p2]];
				v60[safeIndex(171, v60.Length)] := if (f25) then if (i8 in v61) then v61[i8] else v2 else seq(abs(0x3), i14  => (i8));
				var v62: map<int, int> := map[|v42| := 669];
				var v63: seq<map<int, int>> := [v62, map[41 := p2]];
				var v64: map<char, int> := map[f26 := i8];
				var v65 := DC5(|map[DC6(DC3(v63)) := p2]|, fm3(p1, i8, p2, v64, globalState), p1, f25);
				v60[safeIndex(171, v60.Length)], globalState.f5, globalState.f22, v40[safeIndex(54, v40.Length)] := (v2 + v2) + [|v1|, i8], fm2(p1, v65.cf6 - p2, !f25, f26, globalState), i8, if (f25) then f25 else p1;
				v40[safeIndex(54, v40.Length)] := p1 <==> f25;
			}
			
			var v66 := m0(globalState);
		}
	}
	method m1(p0: string, p1: int, p2: seq<bool>, globalState: GlobalState) returns (r0: bool, r1: int) {
		if (!!false) {
			globalState.f21 := p1;
			var v0 := new C0();
			var v1 := m0(globalState);
			globalState.f2, globalState.f5 := f25, safeModuloInt(v0.fm5(globalState), safeModuloInt(p1, v1));
			if (f25) {
				var v2: map<int, int> := map[p1 := p1];
				var v3: seq<map<int, int>> := [v2, v2];
				var v4 := DC3(v3);
				var v5: map<D2, int> := map[v4 := p1];
				v5 := (v5 + v5) + v5;
				var v6: map<char, int> := map[f26 := v1];
				var v7: array<bool> := new bool[16] [true, true, false, !f25, f25, f25, fm3(f25, p1, p1, v6, globalState), f25, f25, f25, f25, f25, !f25, false, f25, f25];
				var v8: multiset<array<bool>> := multiset{v7, v7, v7, v7, v7};
				var v9: map<bool, multiset<array<bool>>> := map[f25 := v8];
				r0 := v7 !in (if (!f25 in v9) then v9[!f25] else multiset{v7, v7, v7, v7, v7});
				globalState.f22 := fm2(f25, p1, f25, f26, globalState);
				var v10: map<bool, int> := map[f25 := p1];
				var v11: set<bool> := {fm0(p1, p1, f25, globalState)};
				var v12: seq<int> := [v1];
				var v13 := DC0(p1);
				var v14: array<int> := new int[26] [v1, (if (f25 in v10) then v10[f25] else -|[f25]|) + v1, |p0|, v0.fm5(globalState), 187, v1, 0xb8, v1, v1, safeDivisionInt(v1, p1), if (f25) then v1 else v1, p1 + |p2|, |v11|, 0x3ac, p1, p1, v1, v0.fm5(globalState), v1, v1, v1, 676, v12[safeIndex(-0x272, |v12|)], 419, |fm12(v1, f25, map[v1 := f25], p1, globalState)|, v13.cf0];
				v14[safeIndex(84, v14.Length)] := v1;
				var v15: multiset<bool> := multiset{fm3(f25, p1, v1, v6, globalState)};
				v14[safeIndex(84, v14.Length)] := safeDivisionInt(if (f25 in v15) then v15[f25] else p1, v1 - |v10|);
				var v16 := new C0();
			} else {
				globalState.f2 := v1 < p1;
				var v17: array<int> := new int[26];
				v17 := v17;
				var v18: set<int> := {p1, p1};
				var v19: multiset<int> := multiset{|v18|};
				var v20: seq<bool> := [f25, f25, p1 <= |v19|];
				v20 := v20;
				var v21: array<array<int>> := new array<int>[26];
				v21[safeIndex(8, v21.Length)] := v17;
				var v22 := DC20([f25, f25, f25, !f25]);
				var v23: T0 := new C1(p0 + p0, |v22.cf37|, false);
				var v24: seq<C0> := [v0, v0];
				var v25: array<seq<C0>> := new seq<C0>[20] [v24, v24, v24[safeIndex(v1, |v24|) := v0], [v0, v24[safeIndex(p1, |v24|)]] + v24, v24, v24, v24, v24, v24, [v0], [v0, v0], v24, v24, v24, v24, [v0, v0], v24, [v0], v24 + [v0], v24];
				v25[safeIndex(997, v25.Length)] := [v0, v0];
				var v26: set<char> := {'d'};
				v21[safeIndex(8, v21.Length)], v23, v25[safeIndex(997, v25.Length)], globalState.f2 := v17, v23, v24[safeIndex(|v26|, |v24|) := v0], v23.f25;
				var v27: array<bool> := new bool[11] [p1 >= v1, v23.f25, v26 >= fm22(v1, v1, p1, f25, globalState), !v23.f25, f25, !f25, v23.f25, p2[safeIndex(-v1, |p2|)], v23.f25, f25 <==> f25, f25];
				v27[safeIndex(685, v27.Length)] := v23.f25;
				v27[safeIndex(685, v27.Length)] := f25;
			}
			
		} else {
			globalState.f5 := if (f25) then p1 else |multiset{!f25}|;
			var v28 := new C1("g", p1, f25);
			var v29: array<int> := new int[12](i0 => i0 - 586);
			v29[safeIndex(144, v29.Length)] := v28.f29;
			v29[safeIndex(144, v29.Length)] := p1;
			var v30: map<bool, bool> := map[!f25 := f25];
			var v31 := DC21(v29[safeIndex(144, v29.Length)], v30 + v30[!true := f25]);
			v31 := fm23(!!f25, globalState);
			var v32 := new C0();
		}
		
		var v33 := DC0(p1);
		match (DC22(p1, v33, f25)) {
			case DC21(cf38, cf39) =>
				var v34: map<int, char> := map[-p1 := f26];
				var v35 := DC14(v34);
				var v36 := DC16(v35);
				var v37: set<bool> := {f25, f25};
				v36 := if (!(v37 > v37)) then v36 else v36;
				globalState.f21 := p1 + p1;
				var v38: array<bool> := new bool[17](i1 => f25);
				var v39: set<int> := {p1};
				v38[safeIndex(962, v38.Length)] := v39 < v39;
				var v40: map<int, bool> := map[cf38 := f25];
				v38[safeIndex(962, v38.Length)] := ((if (cf38 in v40) then v40[cf38] else f25) || false) ==> f25;
				var v41: seq<int> := [|p2|, p1];
				var v42: map<char, bool> := map[f26 := !v38[safeIndex(962, v38.Length)]];
				var v43: array<int> := new int[12] [if (f25) then p1 else p1, cf38, cf38, safeModuloInt(|{v38[safeIndex(962, v38.Length)], p2[safeIndex(-cf38, |p2|)]}|, p1), cf38, safeModuloInt(0x61, cf38), -(if (f25) then v41[safeIndex(cf38, |v41|)] else p1), cf38, p1, --|(v42 + v42)|, DC9(p1, f26, |p0|).cf15, safeDivisionInt(p1, p1)];
				v43[safeIndex(435, v43.Length)] := v41[safeIndex(p1, |v41|)];
				var v44: multiset<bool> := multiset{v38[safeIndex(962, v38.Length)], f25};
				var v45: map<bool, int> := map[f25 := |v44|];
				v43[safeIndex(435, v43.Length)], v40 := safeModuloInt(|v45|, safeModuloInt(-0x1d2, cf38)), map[safeDivisionInt(p1, cf38) := f25];
			case DC22(cf40, cf41, cf42) =>
				var v46 := new C0();
				var v47 := m0(globalState);
				var v48: array<string> := new string[13](i2 => p0);
				var v49: map<bool, array<string>> := map[f25 := v48];
				var v50: map<bool, char> := map[cf42 := f26];
				var v51: seq<map<bool, char>> := [v50, v50];
				v49 := v49[-0x3a9 == |v51| := v48];
				var v52 := DC8(!f25);
				globalState.f2 := (p1 != v47) <== v52.cf12;
			case DC23(cf43, cf44, cf45, cf46) =>
				var v53: multiset<bool> := multiset{f25};
				cf45 := v53 > v53;
				if (!cf45) {
					globalState.f18 := if (cf45) then cf44 else -p1;
					var v54 := DC5(-0x2e3, f25, cf45, fm0(cf44, cf43, true, globalState));
					var v55: map<bool, int> := map[f25 := -cf44];
					var v56: map<bool, bool> := map[cf45 := f25];
					var v57: array<D2> := new D2[18] [v54, v54, v54, DC5(|v55|, f25, if (f25 in v56) then v56[f25] else f25, cf45), v54, v54.(cf8 := f25, cf7 := cf45), v54, v54, v54, v54, v54, v54, v54, v54, v54, v54, v54, v54];
					v57[safeIndex(899, v57.Length)] := v54;
					v57[safeIndex(899, v57.Length)] := v54;
					r0 := f25;
					var v58: array<set<C0>> := new set<C0>[7];
					var v59: C0 := new C0();
					var v60: set<C0> := {v59, v59};
					v58[safeIndex(351, v58.Length)] := v60;
					v58[safeIndex(351, v58.Length)] := {v59};
					var v61 := new C0();
				} else {
					var v62: seq<int> := [cf46, cf46];
					var v63: map<int, int> := map[cf46 := safeModuloInt(v62[safeIndex(|p0|, |v62|)], p1)];
					v63 := v63[cf46 - fm2(false, cf44, cf45, f26, globalState) := cf46];
					var v64 := new C1("awhlibqi", cf46 - p1, cf45);
					cf45 := cf45;
					globalState.f22 := safeDivisionInt(cf46, 0xa0) * v64.f29;
					var v66: map<int, bool> := map[cf44 := f25];
					var v67 := DC17(map v65 : int | v65 in v66 :: (v65 - cf44) := (cf44));
					var v68 := DC19(v67);
					v68 := v68;
				}
				
				var v69: array<char> := new char[17](i3 => f26);
				v69[safeIndex(120, v69.Length)] := f26;
				v69[safeIndex(471, v69.Length)] := f26;
				var v70: set<bool> := {!cf45};
				globalState.f22, v69[safeIndex(120, v69.Length)], v69[safeIndex(471, v69.Length)], r1, cf44 := 169 + cf43, f26, 'n', |(v70 * v70)|, p1;
				globalState.f18 := p1 + cf46;
			case DC20(cf37) =>
				globalState.f2 := f25;
				if (!true) {
					var v71: multiset<int> := multiset{p1, p1, 988, p1, p1};
					globalState.f2 := fm0(610 * |(seq(abs(0x2fb), i4  => (p1)))|, fm2(f25, -|v71[0x38d := abs(p1)]|, !f25, f26, globalState), !false, globalState);
					globalState.f21 := p1;
					globalState.f2 := f25;
					var v72: array<map<bool, bool>> := new map<bool, bool>[9](i5 => map[f25 := f25] + map[f25 := f25]);
					v72 := v72;
					var v73: array<int> := new int[18];
					v73[safeIndex(886, v73.Length)] := p1;
					v73[safeIndex(886, v73.Length)] := p1;
				} else {
					var v74: array<bool> := new bool[5](i6 => f25 ==> f25);
					v74[safeIndex(756, v74.Length)] := f25;
					v74[safeIndex(756, v74.Length)] := f25;
					v74[safeIndex(756, v74.Length)] := fm16(p1, globalState) !in "hqccefcs";
					var v75: seq<seq<bool>> := [p2, p2, cf37, [f25] + p2, p2];
					var v76 := DC20(cf37);
					var v77: map<int, bool> := map[p1 := v74[safeIndex(756, v74.Length)]];
					globalState.f2, r1, globalState.f22, v75, globalState.f22 := v74[safeIndex(756, v74.Length)], |p0| - -(p1 + |[|map[p0 := v76]|]|), if (f25) then |p0| else p1, v75, |v77|;
					cf37 := p2;
					var v78: map<bool, int> := map[f25 := p1];
					var v79: multiset<bool> := multiset{v74[safeIndex(756, v74.Length)], v74[safeIndex(756, v74.Length)], v74[safeIndex(756, v74.Length)]};
					var v80: multiset<map<bool, int>> := multiset{v78 + map[v74[safeIndex(756, v74.Length)] := |v79|]};
					v80 := v80 + v80;
				}
				
				var v81: seq<int> := [p1];
				var v82: multiset<int> := multiset{-|p0|};
				globalState.f2 := multiset(v81 + v81) >= v82;
				var v83: map<int, char> := map[0x264 := f26];
				var v84 := DC16(DC14(v83));
				var v85: multiset<seq<bool>> := multiset{p2, cf37, [f25]};
				v84, r0, v33 := if (f25) then v84 else v84, p1 <= |v85[p2 := abs(p1)]|, v33;
			case DC24(cf47) =>
				var v86 := DC5(p1, f25, !f25, f25);
				var v87 := DC6(v86);
				var v88 := DC6(v87);
				v88 := v88.(cf10 := v87);
				var v89: array<int> := new int[18](i7 => i7 * 0x2e5);
				v89[safeIndex(657, v89.Length)] := p1;
				var v90 := DC22(|p0|, v33, f25);
				v89[safeIndex(854, v89.Length)] := safeDivisionInt(v90.cf40, p1);
				var v91: multiset<int> := multiset{p1, p1, p1, p1};
				var v92: seq<int> := [p1, |fm24(f25, p1, p1, globalState)|];
				v89[safeIndex(657, v89.Length)], v89[safeIndex(854, v89.Length)], globalState.f2 := p1, p1, v91 > multiset{|v92|, p1, p1};
				globalState.f2 := f25;
				var v93: set<bool> := {!f25};
				var v94: map<set<bool>, bool> := map[v93 := f25];
				v94 := v94[v93 := p1 != p1];
		}
		
		var v95: seq<int> := [p1, p1, p1];
		for i8 := p1 to |v95| {
			var v96 := new C0();
			var v97: map<bool, bool> := map[f25 := true];
			var v98: map<bool, bool> := map[!(if (false in v97) then v97[false] else f25) := f25];
			v98 := v98[true := f25];
			var v99: seq<seq<int>> := [v95];
			globalState.f19 := v95 + v99[safeIndex(i8, |v99|)];
			var v100: array<bool> := new bool[7];
			v100[safeIndex(269, v100.Length)] := p1 == p1;
			v100[safeIndex(747, v100.Length)] := true;
			r0, v100[safeIndex(269, v100.Length)], v100[safeIndex(747, v100.Length)], globalState.f22, globalState.f22 := f25, f25, p2[safeIndex(p1, |p2|)], p1, i8;
		}
		var v101: multiset<int> := multiset{p1};
		var v102: map<bool, bool> := map[f25 := f25];
		var v103 := DC21(-p1, v102);
		var v104: array<bool> := new bool[24] [f25, p1 != -0x377, v101[p1 := abs(v103.cf38)] >= v101, !f25, -p1 < p1, f25, !p2[safeIndex(p1, |p2|)] && f25, f25, f25, f25, f25, f25, f25, p1 == p1, f26 !in p0, v101 < multiset(seq(abs(-0x176), i9  => (0xfe))), f25, p2[safeIndex(p1, |p2|)], f25, false, if (f25) then f25 else f25, f25, f25, f25];
		v104[safeIndex(468, v104.Length)] := f25;
		var v105: set<char> := {f26, f26, fm16(p1, globalState)};
		v104[safeIndex(468, v104.Length)] := safeDivisionInt(-284, p1) != |(v105 - v105)|;
		var i10 := 0;
		while (f25)
			decreases 100 - i10
		{
			if (i10 >= 100) {
				break;
			}
			
			i10 := i10 + 1;
			var v106: array<int> := new int[19](i11 => safeModuloInt(i11, p1));
			var v107: set<seq<int>> := {seq(abs(0x1f9), i12  => (|map[|map[p0 := |v105|]| := f25]|))};
			var v108 := DC23(0x3a6, p1, v104[safeIndex(468, v104.Length)], p1);
			var v109: map<seq<int>, D7> := map[v95 := v108];
			var v111: seq<array<bool>> := [v104];
			v104[safeIndex(468, v104.Length)], r0, v104, r0, v106 := v107 >= (set v110 : seq<int> | v110 in v109 :: (v110)), v104[safeIndex(468, v104.Length)], v111[safeIndex(p1, |v111|)], f25, v106;
			var v112: map<int, bool> := map[p1 := f25];
			var v113: set<int> := {fm2(f25, -0x11b, v104[safeIndex(468, v104.Length)], f26, globalState)};
			var v114: array<set<int>> := new set<int>[6] [{p1, fm2(v104[safeIndex(468, v104.Length)], |p2|, v104[safeIndex(468, v104.Length)], 'm', globalState), p1}, fm24(v104[safeIndex(468, v104.Length)], |v112|, p1, globalState), v113, v113 + {p1, -p1}, v113, v113 + {fm2(f25, p1, f25, f26, globalState), |p2|}];
			v114 := v114;
			globalState.f18 := if (true) then -p1 - p1 else p1 - -0x6f;
			v104[safeIndex(468, v104.Length)] := true;
		}
		var v115 := DC9(p1, f26, p1);
		var v116: multiset<string> := multiset{p0[safeIndex(|v95|, |p0|) := v115.cf14], p0};
		r0 := v116 > v116["v" := abs(0x2fc)];
		r0 := f25;
		r1 := p1;
	}
}

class C3 extends T1 {
	constructor (f26 : char, f27 : array<multiset<int>>, f25 : bool) {
		this.f26 := f26;
		this.f27 := f27;
		this.f25 := f25;
	}
	
	function fm3(p0: bool, p1: int, p2: int, p3: map<char, int>, globalState: GlobalState): bool {
		f25
	}
	function fm4(p0: bool, globalState: GlobalState): int {
		safeModuloInt(--0x344 * |[f25, f25]|, 314)
	}
	method m2(p0: set<char>, p1: bool, p2: int, globalState: GlobalState) {
		var v0: array<int> := new int[19];
		v0[safeIndex(994, v0.Length)] := safeModuloInt(p2, p2);
		v0[safeIndex(994, v0.Length)] := safeDivisionInt(p2, safeDivisionInt(fm2(f25, fm2(f25, p2, false, 'e', globalState), p1, f26, globalState), |(map v1 : seq<int> | v1 in (seq(abs(0x5e), i0  => ([|p0|]))) :: (v1) := (|"cela"|))|));
		var v2 := "pgjv";
		v2 := v2;
		var v3 := new C0();
		var v4 := new C0();
		globalState.f2 := f25 <==> p1;
		if (p1) {
			globalState.f18 := safeDivisionInt(p2, p2);
			var v5: set<bool> := {f25};
			globalState.f5 := |v5|;
			globalState.f22 := v0[safeIndex(994, v0.Length)] + v0[safeIndex(994, v0.Length)];
			if (p1) {
				var v6: array<C0> := new C0[21];
				v6[safeIndex(37, v6.Length)] := v3;
				v6[safeIndex(37, v6.Length)] := v3;
				var v7: array<seq<string>> := new seq<string>[1];
				var v8: seq<string> := [v2];
				v7[safeIndex(761, v7.Length)] := v8;
				v7[safeIndex(761, v7.Length)] := v8;
				globalState.f2 := !p1;
				v2 := (v2 + v2) + v2;
				v2 := v2 + "rr";
			} else {
				globalState.f21 := -(safeModuloInt(v0[safeIndex(994, v0.Length)], p2) - 0x1ef);
				globalState.f18 := p2;
				var v9: map<int, int> := map[v0[safeIndex(994, v0.Length)] := p2];
				v9 := v9[p2 := p2 - 0x13f];
				var v10: array<array<string>> := new array<string>[9];
				var v11: array<string> := new string[6](i1 => v2);
				v10[safeIndex(566, v10.Length)] := v11;
				v10[safeIndex(566, v10.Length)] := v11;
				var v12: array<T1> := new T1[4];
				var v13: seq<array<T1>> := [v12];
				v12 := v13[safeIndex(v0[safeIndex(994, v0.Length)], |v13|)];
			}
			
			var v14 := DC0(p2);
			var v15: seq<int> := [p2, v14.cf0];
			var v16: map<int, bool> := map[-0x6b := p1];
			var v17: map<bool, map<int, bool>> := map[|v15| != v0[safeIndex(994, v0.Length)] := v16];
			v17 := v17[true := v16];
		} else {
			var v18: map<char, char> := map[f26 := f26];
			globalState.f2 := fm0(-|v18|, v0[safeIndex(994, v0.Length)], f25, globalState);
			globalState.f18 := v0[safeIndex(994, v0.Length)];
			globalState.f2 := f25;
			v0[safeIndex(994, v0.Length)] := p2;
			var v19: map<int, bool> := map[p2 := f25];
			var v20: map<bool, bool> := map[p1 := p1];
			var v21: set<bool> := {p1, p1, if (!!f25 in v20) then v20[!!f25] else p1, false, p1};
			var v22 := DC2(p1);
			var v23 := DC1(f25);
			var v24: seq<bool> := [v23.cf1];
			var v25: set<seq<bool>> := {v24[safeIndex(p2, |v24|) := p1]};
			var v26: array<bool> := new bool[16] [p1, f25, f25, !f25, -v0[safeIndex(994, v0.Length)] <= v0[safeIndex(994, v0.Length)], false, p2 == p2, if (v0[safeIndex(994, v0.Length)] in v19) then v19[v0[safeIndex(994, v0.Length)]] else f25, fm3(f25, |v21|, fm2(f25, p2, fm3(p1, v0[safeIndex(994, v0.Length)], -711, map[f26 := p2], globalState), f26, globalState), map[f26 := 0x39e], globalState), f25, f25, v22.cf2, f25, p1, v25 < v25, p1];
			v26 := v26;
		}
		
	}
	method m1(p0: string, p1: int, p2: seq<bool>, globalState: GlobalState) returns (r0: bool, r1: int) {
		var v0: multiset<int> := multiset{p1};
		var v1: multiset<int> := multiset{p1, 149, |v0|, p1};
		var i0 := 0;
		while (v1 > (v0 + v0))
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v2: array<string> := new string[14] [p0, p0 + (seq(abs(0x300), i1  => (f26))), seq(abs(-96), i2  => (f26)), p0, p0[safeIndex(p1, |p0|) := f26], p0, p0 + "dhnjxv", seq(abs(0x1d5), i3  => (f26)), "nyxcrvbu", (seq(abs(992), i4  => (f26))) + "imrgbtikj", p0, "dewtel" + "fcxwj", if (f25) then p0 else seq(abs(304), i5  => (f26)), p0];
			v2[safeIndex(908, v2.Length)] := p0;
			v2[safeIndex(908, v2.Length)] := p0[safeIndex(p1, |p0|) := f26];
			var v3: map<int, int> := map[|v2[safeIndex(908, v2.Length)]| := p1];
			var v4: seq<map<int, int>> := [v3];
			var v5 := DC3(v4);
			var v6: multiset<bool> := multiset{f25};
			var v7: map<multiset<bool>, seq<map<int, int>>> := map[v6 := v4];
			globalState.f21 := |(v5.cf3 + (if (v6 in v7) then v7[v6] else [v3]))|;
			var v8: seq<string> := [p0, (v2[safeIndex(908, v2.Length)] + v2[safeIndex(908, v2.Length)])[safeIndex(0x5, |(v2[safeIndex(908, v2.Length)] + v2[safeIndex(908, v2.Length)])|) := f26], v2[safeIndex(908, v2.Length)], fm7(f26, globalState) + "h", v2[safeIndex(908, v2.Length)]];
			v2[safeIndex(908, v2.Length)] := v8[safeIndex(p1, |v8|)];
			var v9: map<bool, int> := map[f25 := fm2(f25, |(seq(abs(0x377), i6  => (f26)))|, f25, 'g', globalState)];
			var v10: array<bool> := new bool[4](i7 => true);
			v10[safeIndex(653, v10.Length)] := f25;
			var v11 := DC2(f25);
			v10[safeIndex(338, v10.Length)] := v11.cf2;
			v9, globalState.f2, v10[safeIndex(653, v10.Length)], r1, v10[safeIndex(338, v10.Length)] := v9, f25, fm0(p1 + p1, p1 - p1, f25, globalState), --p1 - p1, true;
		}
		var v12: array<int> := new int[3](i8 => safeDivisionInt(i8, p1));
		v12[safeIndex(765, v12.Length)] := 0xf9;
		v12[safeIndex(765, v12.Length)] := p1;
		if (f25) {
			r0 := f25;
			var v13 := 'w';
			var v14: set<string> := {p0, p0 + p0, p0};
			var v15: seq<int> := [v12[safeIndex(765, v12.Length)]];
			var v16 := DC0(|[f25]|);
			var v17: map<bool, D0> := map[f25 := v16];
			globalState.f2, v13, r0, globalState.f21, v14 := !fm0(-safeModuloInt(v12[safeIndex(765, v12.Length)], v12[safeIndex(765, v12.Length)]), v12[safeIndex(765, v12.Length)], f25, globalState), f26, v12[safeIndex(765, v12.Length)] == v12[safeIndex(765, v12.Length)], safeModuloInt(|v15|, safeModuloInt(v15[safeIndex(|v17|, |v15|)], |p0[safeIndex(p1, |p0|) := f26]|)), v14;
			v12[safeIndex(765, v12.Length)] := -safeDivisionInt((v16.(cf0 := |map[0x261 := f25]|)).cf0, |p0|);
			var v18: array<array<int>> := new array<int>[3];
			var v19: array<array<array<int>>> := new array<array<int>>[18] [v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, if (f25) then v18 else v18];
			v19[safeIndex(488, v19.Length)] := v18;
			v19[safeIndex(488, v19.Length)] := v18;
			globalState.f5 := p1 + p1;
		} else {
			var v20: set<int> := {0x24a, p1};
			var v21: seq<int> := [|v20|];
			if (if (!(|v21| < -|p2|)) then v12[safeIndex(765, v12.Length)] != v12[safeIndex(765, v12.Length)] else f25 ==> false) {
				v12[safeIndex(765, v12.Length)] := v12[safeIndex(765, v12.Length)];
				r0 := v12[safeIndex(765, v12.Length)] >= 0x1f0;
				globalState.f18 := |p0|;
				globalState.f2 := false;
				var v22 := 't';
				v22, globalState.f2 := v22, f25;
			} else {
				v12[safeIndex(765, v12.Length)] := p1;
				globalState.f2 := f25;
				var v23: T1 := new C2(f26, f27, fm0(v12[safeIndex(765, v12.Length)], v12[safeIndex(765, v12.Length)], f25, globalState));
				var v24: array<bool> := new bool[3](i9 => f25);
				globalState.f5 := |(map[v23 := v24] + map[v23 := v24])|;
				var v25: multiset<bool> := multiset{true, f25};
				v12[safeIndex(765, v12.Length)] := if (v23.f25) then v12[safeIndex(765, v12.Length)] + (if (f25 in v25) then v25[f25] else v12[safeIndex(765, v12.Length)]) else p1;
				var v26: array<T1> := new T1[25];
				v26[safeIndex(550, v26.Length)] := v23;
				globalState.f2, v26[safeIndex(550, v26.Length)] := f25, v23;
			}
			
			var v27: set<seq<int>> := {seq(abs(0x1b3), i10  => (|map[true := f25]|)), [v12[safeIndex(765, v12.Length)]], v21, v21};
			if (v27 !! v27) {
				globalState.f5 := p1;
				globalState.f18 := v12[safeIndex(765, v12.Length)] + v12[safeIndex(765, v12.Length)];
				var v28: map<int, bool> := map[p1 := f25];
				var v29 := DC7(v28);
				var v30: array<D3> := new D3[7] [fm25(globalState), fm25(globalState), DC7(v28), fm25(globalState), v29, v29, v29];
				v30[safeIndex(559, v30.Length)] := v29;
				v30[safeIndex(559, v30.Length)] := v29.(cf11 := v28);
				var v31: map<seq<int>, string> := map[v21 := p0];
				v31 := v31[v21 := "yjmlwg" + "x"];
				var v32 := new C2(f26, f27, f25);
			} else {
				globalState.f5 := 0x327;
				globalState.f5 := p1;
				var v33: set<bool> := {f25};
				globalState.f18 := |v33|;
				var v34: T1 := new C2(f26, f27, f25);
				globalState.f2, v34 := f25, if (f25) then v34 else v34;
				globalState.f2 := f25;
			}
			
			globalState.f2 := f25;
			var v35: array<bool> := new bool[29](i11 => false);
			v35[safeIndex(574, v35.Length)] := false;
			v35[safeIndex(574, v35.Length)] := f25 || !f25;
			var v36 := DC15(f25, v35, p0, false);
			match (v36) {
				case DC15(cf25, cf26, cf27, cf28) =>
					globalState.f18 := p1;
					globalState.f22 := fm2(v35[safeIndex(574, v35.Length)], v12[safeIndex(765, v12.Length)], f25, 'f', globalState);
					var v37 := new C2(f26, f27, cf25);
					var v38 := DC9(p1, 'r', p1);
					var v39 := DC10(v38);
					var v40: map<int, bool> := map[p1 := f25];
					v39 := v39.(cf16 := DC7(v40));
				case DC14(cf24) =>
					var v41 := DC5(204, true, v35[safeIndex(574, v35.Length)], f25);
					var v42: map<int, array<bool>> := map[v12[safeIndex(765, v12.Length)] := v35];
					var v43: map<int, map<int, array<bool>>> := map[v41.cf6 := v42];
					v12[safeIndex(765, v12.Length)] := if (v35[safeIndex(574, v35.Length)]) then v12[safeIndex(765, v12.Length)] else |(if (p1 in v43) then v43[p1] else map[0xab := v35])|;
					globalState.f18 := -safeModuloInt(335 + v12[safeIndex(765, v12.Length)], v12[safeIndex(765, v12.Length)] - -v12[safeIndex(765, v12.Length)]);
					var v44: C2 := new C2(f26, f27, f25);
					var v45: map<C2, string> := map[v44 := p0 + (seq(abs(44), i12  => (f26)))];
					v45 := v45;
					var v46 := "if";
					v46 := (p0 + "xdalhwjvr") + v46;
				case DC16(cf29) =>
					globalState.f2 := v12[safeIndex(765, v12.Length)] != v12[safeIndex(765, v12.Length)];
					globalState.f5 := v12[safeIndex(765, v12.Length)];
					var v47: map<bool, string> := map[v35[safeIndex(574, v35.Length)] := p0];
					globalState.f5 := -v21[safeIndex(-|v47|, |v21|)];
					v35[safeIndex(574, v35.Length)] := p1 !in v0[p1 := abs(v12[safeIndex(765, v12.Length)])];
			}
			
		}
		
		r1 := (v12[safeIndex(765, v12.Length)] * v12[safeIndex(765, v12.Length)]) + safeDivisionInt(p1, |p0|);
		r1 := v12[safeIndex(765, v12.Length)];
		r0 := f25;
		var v49 := DC25(v1);
		var v50 := DC17(map v48 : int | v48 in v49.cf48[p1 := abs(v12[safeIndex(765, v12.Length)])] :: (safeModuloInt(v48, |(seq(abs(0x42), i13  => (f26)))|)) := (p1));
		r0 := match v50 {
			case DC18(cf31, cf32, cf33, cf34, cf35) => cf33 || cf34
			case DC17(cf30) => true
			case DC19(cf36) => f25
		};
		r1 := v12[safeIndex(765, v12.Length)];
	}
}

function fm0(p0: int, p1: int, p2: bool, globalState: GlobalState): bool {
	if (false && false) then !true <== true else (seq(abs(0x192), i0  => ('x'))) != "pdxgcrjxl"
}
function fm1(p0: int, p1: int, globalState: GlobalState): multiset<bool> {
	multiset{!!false} * (multiset{!false, true} * multiset{true, false})
}
function fm2(p0: bool, p1: int, p2: bool, p3: char, globalState: GlobalState): int {
	(if (false) then |[0xf0, -0xe7, |"rljhwsp"|, 0xce]| else |map['d' := |map[|"xe"| := |(set v0 : int | v0 in [0xdf] :: (v0 + 741))|]|]|) + (|['d']| * -0x183)
}
function fm7(p0: char, globalState: GlobalState): string {
	if (true) then "cs" else "huytibwmr"
}
function fm10(globalState: GlobalState): char {
	match DC6(DC6(DC5(-0x8b, true, true, true))) {
		case DC4(cf4, cf5) => 't'
		case DC5(cf6, cf7, cf8, cf9) => 'v'
		case DC3(cf3) => if (false) then 'e' else 'n'
		case DC6(cf10) => 'o'
	}
}
function fm11(globalState: GlobalState): D2 {
	DC4(855 + 195, {-694, 0x129} > {0xbc})
}
function fm12(p0: int, p1: bool, p2: map<int, bool>, p3: int, globalState: GlobalState): seq<int> {
	[-(0x162 - 0x3ba)]
}
function fm13(p0: bool, p1: int, globalState: GlobalState): map<bool, int> {
	map[!!!!!false := 430] + (if (true) then map[true := |[[false, !true], [false], [true]]|] else map[true := |[false]|])
}
function fm16(p0: int, globalState: GlobalState): char {
	'a'
}
function fm17(p0: bool, p1: int, globalState: GlobalState): seq<bool> {
	[!!true, false, false, {35} >= {0x168}]
}
function fm18(p0: int, p1: int, globalState: GlobalState): string {
	"ckp" + (if (!false) then "dy" else seq(abs(0x168), i0  => ('y')))
}
function fm19(p0: int, p1: bool, globalState: GlobalState): map<int, char> {
	(map[|map[true := -|"kimkmjliq"|]| := 'x'] + (map v0 : int | v0 in [-0x2b3, 0x175] :: (v0 * |[0x136]|) := ('x'))) + map[|(map v1 : D3 | v1 in [DC9(|[false]|, 'c', 0x200)] :: (v1) := (false))| := 'd']
}
function fm20(p0: int, globalState: GlobalState): D2 {
	DC6(DC5(|multiset{'n'}|, !!true, !false, false))
}
function fm21(globalState: GlobalState): map<int, int> {
	map[|([0x280] + [|(seq(abs(-152), i0  => (|"moh"|)))|, 532])| := safeDivisionInt(0x31e, |"hnvcppjf"|)]
}
function fm22(p0: int, p1: int, p2: int, p3: bool, globalState: GlobalState): set<char> {
	if (!false) then {'p'} else {'h', 'l'}
}
function fm23(p0: bool, globalState: GlobalState): D7 {
	DC21(-405, map[!true := true] + map[true := false])
}
function fm24(p0: bool, p1: int, p2: int, globalState: GlobalState): set<int> {
	({|{!!false}|} + {--|"bdxocc"|}) + {0x1c1}
}
function fm25(globalState: GlobalState): D3 {
	if (false) then DC7(map[DC22(|map[|map[0x2b5 := |map[true := false]|]| := 0x2a]|, DC0(|{'p'}|), false).cf40 := true]) else DC7(map[0x13f := true])
}
function fm26(p0: bool, p1: int, p2: int, globalState: GlobalState): map<int, bool> {
	map[|{-|multiset{false}|, |(set v0 : int | v0 in (seq(abs(0x181), i0  => (|(seq(abs(-0x2e2), i1  => (|[true]|)))|))) :: (v0 + 0x2ae))|}| := !true] + map[-0x1d9 := true]
}
function fm27(globalState: GlobalState): map<bool, bool> {
	match DC8(false) {
		case DC8(cf12) => map[cf12 := !!cf12]
		case DC9(cf13, cf14, cf15) => map[false := true]
		case DC7(cf11) => map[!!true := !!false] + map[true := false]
		case DC10(cf16) => DC21(|map[true := false]|, map[false := true]).cf39
	}
}
function fm28(p0: bool, p1: int, p2: int, globalState: GlobalState): map<set<int>, map<int, bool>> {
	map[{|[false]|} + {-|multiset{192, 0xff}|} := map v0 : int | (0xeb <= v0) && (v0 < -647) :: (v0 - 0x36) := (true)]
}
method m0(globalState: GlobalState) returns (r0: int) {
	var v0 := 0x8e;
	var v1: seq<int> := [v0];
	var v2 := "rfkqvu";
	var v3: set<int> := {v0, v1[safeIndex(|v2|, |v1|)]};
	for i0 := 0x39f to safeDivisionInt(v0, |v3|) {
		var v4 := false;
		globalState.f2 := v4;
		var v5 := DC11(v2);
		v4 := v2 < v5.cf17;
		var v6 := new C0();
		var v7 := 'm';
		var v8: set<char> := {v7, v7, v7, v7, v7};
		var v9: multiset<set<char>> := multiset{v8};
		globalState.f18, globalState.f21, v9, globalState.f5 := i0, safeDivisionInt(v0, i0), v9 * v9, v0;
	}
	if (false) {
		var v10 := false;
		globalState.f2 := !(v2 <= (if (v10) then v2 else v2));
		var v11: set<bool> := {v10, v10, v10};
		if (!fm0(|multiset{|v11|, v0}|, safeDivisionInt(v0, 775), v10, globalState)) {
			var v12 := 'u';
			var v13: array<multiset<int>> := new multiset<int>[9];
			var v14 := new C2(v12, v13, true);
			var v15: array<int> := new int[25];
			var v16 := DC12(true, v15, v0);
			globalState.f2 := v16.cf18;
			var v17: seq<bool> := [v10];
			globalState.f2 := |v17| >= v0;
			v15 := v15;
			var v18 := new C0();
		} else {
			globalState.f2 := v10;
			var v19: array<int> := new int[23](i1 => safeModuloInt(i1, v0));
			v19[safeIndex(47, v19.Length)] := if (v10) then |v1[safeIndex(|"ec"|, |v1|) := v0]| else |v3|;
			v19[safeIndex(47, v19.Length)] := v1[safeIndex(v0, |v1|)];
			v2 := v2;
			globalState.f21 := v19[safeIndex(47, v19.Length)];
			var v20: array<bool> := new bool[7] [true, v10, !v10, v10, v10, v10, v10];
			var v21: multiset<array<bool>> := multiset{v20};
			v19[safeIndex(47, v19.Length)] := |v21|;
		}
		
		globalState.f5 := safeModuloInt(v0, 0x35c);
		var v22 := new C0();
		var v23: seq<bool> := [!v10, v10];
		var v24: array<int> := new int[12] [v0, v0, |v23|, v0, -v0 + -v0, v0, v0, v0, v0, v0, -(v0 * 250), v0];
		var v25: map<int, bool> := map[v0 := v10];
		v24[safeIndex(794, v24.Length)] := safeDivisionInt(v0, |map[v0 := v25]|);
		v24[safeIndex(794, v24.Length)] := v22.fm5(globalState);
	} else {
		var v26 := 'p';
		var v27: multiset<char> := multiset{v26};
		var v28: array<array<D2>> := new array<D2>[21];
		var v29 := false;
		var v30 := DC4(v0, v29);
		var v31: array<D2> := new D2[26] [v30, v30, v30, v30, v30, v30, v30, DC4(-v0, v29), v30, v30, v30, DC4(v0, v29), v30, v30, v30, v30, DC4(v0, v29).(cf5 := v29), v30, v30, v30, v30, DC4(|v2|, v29), v30, v30, DC4(-v0, false), v30];
		v28[safeIndex(395, v28.Length)] := v31;
		var v32: set<bool> := {v29, v29};
		var v33: seq<array<D2>> := [v31, v31, if (v29) then v31 else v31];
		v27, v28[safeIndex(395, v28.Length)], globalState.f18, v3, globalState.f5 := v27[v26 := abs(|v32|)] - v27, v33[safeIndex(v0, |v33|)], -v0, v3, fm2(v0 == |multiset(v1)|, -fm2(v29, -v0, v29, v26, globalState), v0 >= v0, 'd', globalState);
		globalState.f2 := v29;
		var v34: array<C2> := new C2[18];
		var v35: array<multiset<int>> := new multiset<int>[12](i2 => multiset{v0, v0});
		var v36: C2 := new C2(v26, v35, v29);
		v34[safeIndex(570, v34.Length)] := v36;
		v34[safeIndex(570, v34.Length)] := v36;
		var v37: array<bool> := new bool[4](i3 => 'n' !in v2);
		var v38: seq<seq<int>> := [seq(abs(-0x197), i4  => (v0))];
		v37[safeIndex(744, v37.Length)] := |v38[safeIndex(v0, |v38|)]| != v0;
		v37[safeIndex(744, v37.Length)] := if (v1 < [v0, |v1|]) then v29 else v29;
		globalState.f19 := seq(abs(-27), i5  => (v0));
	}
	
	var v39 := false;
	var v40 := 'p';
	var v41: T0 := new C1(v2, v0, v39);
	var v42: seq<T0> := [v41];
	var v43: seq<T0> := [v42[safeIndex(v0, |v42|)]];
	var v44: map<int, int> := map[fm2(v39, v0, v39, v40, globalState) := |v43|];
	globalState.f2 := (v44 + (map v45 : int | v45 in v1 :: (safeDivisionInt(v45, v0)) := (0x2))) !in [map[|v2| := v0]];
	var v46: array<multiset<int>> := new multiset<int>[7];
	forall i6 | 0 <= i6 < v46.Length {
		v46[i6] := multiset{DC5(v0, v41.f25, v41.f25, v41.f25).cf6, 0x34b};
	}
	var v47: multiset<int> := multiset{v0};
	var v49: map<int, bool> := map[v0 := fm0(-145, -v0, true, globalState)];
	var v50: set<bool> := {if (v0 in v49) then v49[v0] else v39, v39};
	var v51: map<bool, char> := map[v41.f25 := v40];
	var v52: map<int, map<bool, char>> := map[v0 := v51];
	var v53: map<bool, bool> := map[v41.f25 := v39];
	var v54 := DC21(v0, fm27(globalState));
	var v55: C3 := new C3(v40, v46, v41.f25);
	var v56: array<int> := new int[28] [|v1|, -v0, v0, v0, safeDivisionInt(fm2(fm0(v0, v0, v39, globalState), |v47|, v39, v40, globalState), v0), v0, |(seq(abs(0x268), i7  => (v40)))|, v0, --(v0 * v0), v0, v0, |(set v48 : int | (0x75 <= v48) && (v48 < 0x17c) :: (safeDivisionInt(v48, v0)))| - --v0, |v50| + 0x4, v0, v0, 0x6f, |v52| * v0, v0, fm2(v39, 0xbf, v39, v40, globalState), v0, v0, v0, |(v53[v39 := v41.f25] + v54.cf39)|, v0, safeModuloInt(v0, -v0), 0x4d, v0, |map[v55 := v0]|];
	v56[safeIndex(816, v56.Length)] := -(v0 - |v2|);
	var v57: C1 := new C1("s", |v2|, v41.f25);
	var v58 := DC1(v39);
	var v59: map<C1, D1> := map[v57 := v58];
	v56[safeIndex(816, v56.Length)] := -|(v59[v57 := v58] + v59)|;
	if (v39 <== true) {
		var v60: map<int, set<bool>> := map[v57.f29 := {v39, v41.f25, v39, v41.f25}];
		var v61: seq<bool> := [v41.f25];
		var v62: map<set<bool>, seq<bool>> := map[if (v57.f29 in v60) then v60[v57.f29] else v50 := v61];
		var v63: seq<map<set<bool>, seq<bool>>> := [v62];
		var v64: seq<multiset<int>> := [v47];
		var v66: set<multiset<int>> := {v47};
		var v67: array<C0> := new C0[23];
		var v68 := DC13(v67, v39, v41.f25);
		var v69: map<seq<int>, bool> := map[[38, |v1|] := v41.f25];
		var v70: array<bool> := new bool[29] [v39, v2 <= v2, v50 >= {v41.f25, v41.f25}, v41.f25, true, false, v41.f25, v50 in v63[safeIndex(v57.f29, |v63|)], v41.f25, v39, v41.f25, v41.f25, v41.f25, v56[safeIndex(816, v56.Length)] !in v44, true, v41.f25, v39, v41.f25, v41.f25, v41.f25, v41.f25, v41.f25, (set v65 : multiset<int> | v65 in v64 :: (v65)) !! v66, v68.cf22, v47 >= multiset(v1), !(|v69| <= 0x5b), v57.f28 <= v2, v57.f29 > v56[safeIndex(816, v56.Length)], v0 >= v56[safeIndex(816, v56.Length)]];
		var v71: map<multiset<int>, C3> := map[v47 * v47 := v55];
		v70, v55 := v70, if (v47 in v71) then v71[v47] else v55;
		var v72 := DC15(v41.f25, v70, v57.f28, v39);
		var v73 := DC16(v72);
		var v74: map<bool, D5> := map[v0 < v57.f29 := DC16(v72)];
		var v75 := DC16(v72);
		v74 := v74[v39 := v75];
		var v76: T1 := new C3(v40, v46, v41.f25);
		var v77: seq<T1> := [v76];
		v77 := v77 + (v77 + v77);
		var v78 := new C0();
		globalState.f18 := v57.f29;
	} else {
		globalState.f22 := -v57.f29;
		var v79: map<int, T0> := map[-974 := v41];
		if (v79 == map[-0x11b := v41]) {
			var v80: array<bool> := new bool[13];
			v80[safeIndex(850, v80.Length)] := v57.f28 <= v2;
			v80[safeIndex(850, v80.Length)] := v39;
			v80[safeIndex(850, v80.Length)], v55 := v57.f28 != v57.f28, v55;
			globalState.f2 := (v56[safeIndex(816, v56.Length)] + 0x2e3) == v57.f29;
			v47 := multiset{v57.f29} * (v47 - v47[v56[safeIndex(816, v56.Length)] := abs(v56[safeIndex(816, v56.Length)])]);
			globalState.f2 := v41.f25;
		} else {
			v56[safeIndex(816, v56.Length)], globalState.f18 := safeModuloInt(v56[safeIndex(816, v56.Length)], v57.f29) * v57.f29, v56[safeIndex(816, v56.Length)];
			globalState.f2 := !v41.f25;
			var v81: C2 := new C2(v40, v46, !v41.f25);
			var v82: seq<C2> := [v81];
			var v83: map<string, char> := map[fm18(|v82|, v57.f29, globalState) := v40];
			var v84: T1 := new C3(if (v57.f28 in v83) then v83[v57.f28] else v40, v46, !true);
			var v85: map<array<int>, T1> := map[v56 := v84];
			var v86: map<bool, T1> := map[false := if (v56 in v85) then v85[v56] else v84];
			v84 := if (true in v86) then v86[true] else v84;
			var v87: array<bool> := new bool[17];
			v87[safeIndex(169, v87.Length)] := v41.f25;
			v87[safeIndex(169, v87.Length)] := v84.f25;
			var v88: map<set<int>, map<int, bool>> := map[v3 := v49];
			var v90: map<char, int> := map[v84.f26 := v57.f29];
			var v91: multiset<bool> := multiset{v41.f25, v87[safeIndex(169, v87.Length)], v41.f25, v41.f25, v41.f25};
			v88 := map[set v89 : int | v89 in v3 :: (v89 + |multiset{0x139}|) := map[-325 := v87[safeIndex(169, v87.Length)]]] + (fm28(v55.fm3(v87[safeIndex(169, v87.Length)], 0x333, |map[v57.f29 := v56[safeIndex(816, v56.Length)]]|, v90, globalState), if (v41.f25 in v91) then v91[v41.f25] else v57.f29, 342, globalState))[v3 := v49];
		}
		
		globalState.f2 := v41.f25;
		globalState.f22 := v0;
		var v92: seq<bool> := [v41.f25];
		r0 := safeModuloInt(|v92| + v56[safeIndex(816, v56.Length)], v57.f29);
	}
	
	var v93: multiset<bool> := multiset{v39};
	r0 := (v56[safeIndex(816, v56.Length)] - |v93[v41.f25 := abs(0x275)]|) * |v3|;
}
method Main() {
	var v0 := "qyvk";
	var v1: array<set<int>> := new set<int>[11](i0 => {|v0|, |v0|});
	var v2: array<array<bool>> := new array<bool>[19];
	var v3 := true;
	var v4: set<bool> := {v3};
	var v5: seq<int> := [|v4|];
	var v6: array<bool> := new bool[20];
	var globalState := new GlobalState(460, 0xee, true, v0, v1, 0x2ba, false, 0x3ab, false, v2, 0x293, false, 'w', false, false, 856, 966, false, 806, v5, 962, 0x231, 0xb6, v6, 'w');
	var v7 := m0(globalState);
	var v8: multiset<string> := multiset{v0};
	var v9: multiset<bool> := multiset{v3, !(v8 != v8), !v3 ==> fm0(v7, v7, v3, globalState), v3, v3};
	var v10: seq<bool> := [v3, v3];
	v9 := fm1(v7, |(v10 + v10)|, globalState);
	globalState.f22, globalState.f2, v6 := v7 * v7, !v3, v6;
	var v11 := 'k';
	v11 := v11;
	var v12: map<int, bool> := map[fm2(v3, v7, v3, v11, globalState) := v3];
	var v13: seq<multiset<bool>> := [v9];
	var v14: array<multiset<bool>> := new multiset<bool>[21] [multiset{if (0x39e in v12) then v12[0x39e] else v3} + v9, v9, v9, v9, fm1(v7, v7, globalState), multiset{true}, v9, v13[safeIndex(v7, |v13|)], fm1(v7, v7, globalState), v9, v9 * fm1(v7, v7, globalState), v9, fm1(-362, v7, globalState), multiset(v10), v9, v9, multiset{v3} * v9, v9, multiset((([v3])[safeIndex(v7, |[v3]|) := v3])[safeIndex(317, |([v3])[safeIndex(v7, |[v3]|) := v3]|) := v3]), v9, v9 * v9];
	v14[safeIndex(115, v14.Length)] := v9;
	v6[safeIndex(939, v6.Length)] := true;
	var v15: map<bool, multiset<bool>> := map[!(fm2(v3, v7, !v3, v11, globalState) > |v0|) := multiset(v10)];
	v7, v3, v14[safeIndex(115, v14.Length)], v6[safeIndex(939, v6.Length)], v7 := v7, false, if (v3 in v15) then v15[v3] else v9, v7 >= |[fm2(v3, |v0[safeIndex(v7, |v0|) := v11]|, v3, v11, globalState)]|, v7;
	var v16 := m0(globalState);
	var v17: map<bool, bool> := map[v6[safeIndex(939, v6.Length)] := v6[safeIndex(939, v6.Length)]];
	var v18: multiset<int> := multiset{|v17|};
	var v19: set<int> := {v16, v7};
	var v20: array<int> := new int[19] [v16, v16, v16, v16, |v0|, v16, 0x2f5, v16, v16, fm2(v6[safeIndex(939, v6.Length)], v7, v3, v11, globalState), |v18|, v16, |v19|, v7, fm2(v3, v16, v3, 'w', globalState), v16, -v7, v7, fm2(v6[safeIndex(939, v6.Length)], |v4|, v3, v11, globalState)];
	for i1 := |(map[v5 := v20])[v5 := v20]| to v16 {
		var v21: array<map<int, bool>> := new map<int, bool>[17];
		v21[safeIndex(892, v21.Length)] := v12;
		v14[safeIndex(115, v14.Length)], v21[safeIndex(892, v21.Length)] := multiset{v6[safeIndex(939, v6.Length)], v6[safeIndex(939, v6.Length)]}, v12;
		v0 := (if (v6[safeIndex(939, v6.Length)]) then v0 else v0)[safeIndex(safeDivisionInt(i1, v16), |(if (v6[safeIndex(939, v6.Length)]) then v0 else v0)|) := 'n'];
		var v22 := m0(globalState);
		var v23 := new C0();
	}
	var v24 := m0(globalState);
	v24 := v16;
	globalState.f2 := v6[safeIndex(939, v6.Length)];
	v20[safeIndex(793, v20.Length)] := v7;
	v20[safeIndex(793, v20.Length)] := 0x3c7;
	globalState.f5 := safeDivisionInt(fm2(if (true in v17) then v17[true] else !v3, 0x13d, v6[safeIndex(939, v6.Length)], v11, globalState), -|((map v25 : int | (0x29e <= v25) && (v25 < 941) :: (safeDivisionInt(v25, 0x391)) := (v3)) + fm26(v6[safeIndex(939, v6.Length)], if (v3 in v9) then v9[v3] else v16, |v0|, globalState))|);
	var v26: array<multiset<int>> := new multiset<int>[14](i2 => v18);
	var v27 := new C3(v11, v26, v20[safeIndex(793, v20.Length)] == v7);
	var v28 := DC23(v24, 0x126, v3, v7);
	var v29 := DC24(v28);
	var i3 := 0;
	while (match v29 {
		case DC21(cf38, cf39) => v3
		case DC22(cf40, cf41, cf42) => v3
		case DC23(cf43, cf44, cf45, cf46) => "tdneecvnd" <= (seq(abs(-0x1cc), i4  => (v11)))
		case DC20(cf37) => v6[safeIndex(939, v6.Length)]
		case DC24(cf47) => v0 <= "qarhghwh"
	})
		decreases 100 - i3
	{
		if (i3 >= 100) {
			break;
		}
		
		i3 := i3 + 1;
		v6[safeIndex(939, v6.Length)] := v3;
		var v30 := m0(globalState);
		v0 := "mukobu" + v0;
		var v31: set<char> := {v11};
		v27.m2(if (v6[safeIndex(939, v6.Length)]) then v31 else {v11, v11, 'o', v11, v11}, v3, |(seq(abs(-0x2c8), i5  => (v0)))|, globalState);
	}
	forall i6 | 0 <= i6 < v6.Length {
		v6[i6] := v3;
	}
	var v32: seq<array<bool>> := [v6];
	var v33 := DC9(|v32|, v11, 0x146);
	var v34: map<char, int> := map[v33.cf14 := v16];
	if (v27.fm3(false, v24, fm2(v3, v16, v6[safeIndex(939, v6.Length)], v11, globalState), v34 + v34, globalState)) {
		globalState.f5 := v20[safeIndex(793, v20.Length)];
		globalState.f22 := if (v24 in v18) then v18[v24] else v24;
		v32 := v32;
		var v35: array<D3> := new D3[24];
		var v36: map<array<D3>, seq<bool>> := map[v35 := v10 + v10];
		v16, v36, v18 := v16, map[v35 := v10 + [v3]], v18 * (v18 - multiset(v5));
		globalState.f2 := v6[safeIndex(939, v6.Length)];
	} else {
		globalState.f2 := v18 != v18;
		globalState.f2 := !(v4 !! v4);
		v3 := v27.fm3(v3, 0x3cb, v20[safeIndex(793, v20.Length)], v34, globalState);
		var v37 := DC26(v20, v6[safeIndex(939, v6.Length)], v0[safeIndex(v20[safeIndex(793, v20.Length)], |v0|)], v6[safeIndex(939, v6.Length)], v3);
		var v38: seq<D8> := [v37];
		v38 := if (v3) then v38 else if (v6[safeIndex(939, v6.Length)]) then v38 else v38;
		var v39: map<int, seq<bool>> := map[v7 := v10];
		v39 := v39[v24 := [fm0(v7, 0x200, v6[safeIndex(939, v6.Length)], globalState)]];
	}
	
	print v0, "\n";
	print v1[0] == {4}, "\n";
	print v1[1] == {4}, "\n";
	print v1[2] == {4}, "\n";
	print v1[3] == {4}, "\n";
	print v1[4] == {4}, "\n";
	print v1[5] == {4}, "\n";
	print v1[6] == {4}, "\n";
	print v1[7] == {4}, "\n";
	print v1[8] == {4}, "\n";
	print v1[9] == {4}, "\n";
	print v1[10] == {4}, "\n";
	print v3, "\n";
	print v4 == {true}, "\n";
	print v5 == [1], "\n";
	print v6[0], "\n";
	print v6[1], "\n";
	print v6[2], "\n";
	print v6[3], "\n";
	print v6[4], "\n";
	print v6[5], "\n";
	print v6[6], "\n";
	print v6[7], "\n";
	print v6[8], "\n";
	print v6[9], "\n";
	print v6[10], "\n";
	print v6[11], "\n";
	print v6[12], "\n";
	print v6[13], "\n";
	print v6[14], "\n";
	print v6[15], "\n";
	print v6[16], "\n";
	print v6[17], "\n";
	print v6[18], "\n";
	print v6[19], "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3, "\n";
	print globalState.f4[0] == {4}, "\n";
	print globalState.f4[1] == {4}, "\n";
	print globalState.f4[2] == {4}, "\n";
	print globalState.f4[3] == {4}, "\n";
	print globalState.f4[4] == {4}, "\n";
	print globalState.f4[5] == {4}, "\n";
	print globalState.f4[6] == {4}, "\n";
	print globalState.f4[7] == {4}, "\n";
	print globalState.f4[8] == {4}, "\n";
	print globalState.f4[9] == {4}, "\n";
	print globalState.f4[10] == {4}, "\n";
	print globalState.f5, "\n";
	print globalState.f6, "\n";
	print globalState.f7, "\n";
	print globalState.f8, "\n";
	print globalState.f10, "\n";
	print globalState.f11, "\n";
	print globalState.f12, "\n";
	print globalState.f13, "\n";
	print globalState.f14, "\n";
	print globalState.f15, "\n";
	print globalState.f16, "\n";
	print globalState.f17, "\n";
	print globalState.f18, "\n";
	print globalState.f19 == [142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142], "\n";
	print globalState.f20, "\n";
	print globalState.f21, "\n";
	print globalState.f22, "\n";
	print globalState.f23[0], "\n";
	print globalState.f23[1], "\n";
	print globalState.f23[2], "\n";
	print globalState.f23[3], "\n";
	print globalState.f23[4], "\n";
	print globalState.f23[5], "\n";
	print globalState.f23[6], "\n";
	print globalState.f23[7], "\n";
	print globalState.f23[8], "\n";
	print globalState.f23[9], "\n";
	print globalState.f23[10], "\n";
	print globalState.f23[11], "\n";
	print globalState.f23[12], "\n";
	print globalState.f23[13], "\n";
	print globalState.f23[14], "\n";
	print globalState.f23[15], "\n";
	print globalState.f23[16], "\n";
	print globalState.f23[17], "\n";
	print globalState.f23[18], "\n";
	print globalState.f23[19], "\n";
	print globalState.f24, "\n";
	print v7, "\n";
	print v8 == multiset{"qyvk"}, "\n";
	print v9 == multiset{}, "\n";
	print v10 == [true, true], "\n";
	print v11, "\n";
	print v12 == map[-386 := true], "\n";
	print v13 == [multiset{}], "\n";
	print v14[0] == multiset{true}, "\n";
	print v14[1] == multiset{}, "\n";
	print v14[2] == multiset{}, "\n";
	print v14[3] == multiset{}, "\n";
	print v14[4] == multiset{}, "\n";
	print v14[5] == multiset{true}, "\n";
	print v14[6] == multiset{}, "\n";
	print v14[7] == multiset{}, "\n";
	print v14[8] == multiset{}, "\n";
	print v14[9] == multiset{}, "\n";
	print v14[10] == multiset{true, true}, "\n";
	print v14[11] == multiset{}, "\n";
	print v14[12] == multiset{}, "\n";
	print v14[13] == multiset{true, true}, "\n";
	print v14[14] == multiset{}, "\n";
	print v14[15] == multiset{}, "\n";
	print v14[16] == multiset{}, "\n";
	print v14[17] == multiset{}, "\n";
	print v14[18] == multiset{true}, "\n";
	print v14[19] == multiset{}, "\n";
	print v14[20] == multiset{}, "\n";
	print v15 == map[true := multiset{true, true}], "\n";
	print v16, "\n";
	print v17 == map[false := false], "\n";
	print v18 == multiset{1}, "\n";
	print v19 == {-599}, "\n";
	print v20[0], "\n";
	print v20[1], "\n";
	print v20[2], "\n";
	print v20[3], "\n";
	print v20[4], "\n";
	print v20[5], "\n";
	print v20[6], "\n";
	print v20[7], "\n";
	print v20[8], "\n";
	print v20[9], "\n";
	print v20[10], "\n";
	print v20[11], "\n";
	print v20[12], "\n";
	print v20[13], "\n";
	print v20[14], "\n";
	print v20[15], "\n";
	print v20[16], "\n";
	print v20[17], "\n";
	print v20[18], "\n";
	print v24, "\n";
	print v26[0] == multiset{1}, "\n";
	print v26[1] == multiset{1}, "\n";
	print v26[2] == multiset{1}, "\n";
	print v26[3] == multiset{1}, "\n";
	print v26[4] == multiset{1}, "\n";
	print v26[5] == multiset{1}, "\n";
	print v26[6] == multiset{1}, "\n";
	print v26[7] == multiset{1}, "\n";
	print v26[8] == multiset{1}, "\n";
	print v26[9] == multiset{1}, "\n";
	print v26[10] == multiset{1}, "\n";
	print v26[11] == multiset{1}, "\n";
	print v26[12] == multiset{1}, "\n";
	print v26[13] == multiset{1}, "\n";
	print v28.cf43, "\n";
	print v28.cf44, "\n";
	print v28.cf45, "\n";
	print v28.cf46, "\n";
	print v29.cf47.cf43, "\n";
	print v29.cf47.cf44, "\n";
	print v29.cf47.cf45, "\n";
	print v29.cf47.cf46, "\n";
	print i3, "\n";
	print |v32|, "\n";
	print v33.cf13, "\n";
	print v33.cf14, "\n";
	print v33.cf15, "\n";
	print v34 == map['k' := -599], "\n";
}