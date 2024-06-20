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
datatype D0 = DC1(cf1: array<int>, cf2: array<bool>, cf3: multiset<seq<char>>, cf4: int, cf5: int) | DC0(cf0: string)
datatype D1 = DC3(cf7: bool, cf8: bool, cf9: bool, cf10: bool) | DC2(cf6: char) | DC4(cf11: D1)
datatype D2 = DC5(cf12: map<char, array<int>>)
datatype D3 = DC7 | DC8(cf14: int) | DC6(cf13: seq<array<bool>>)
datatype D4 = DC10(cf16: int, cf17: bool, cf18: bool) | DC11(cf19: char, cf20: bool, cf21: multiset<char>, cf22: bool, cf23: seq<multiset<int>>) | DC12(cf24: int, cf25: array<map<int, bool>>, cf26: bool, cf27: int) | DC9(cf15: C0)
class GlobalState {
	const f0 : int
	const f1 : int
	const f2 : bool
	const f3 : int
	const f4 : bool
	var f5 : int
	var f6 : bool
	const f7 : seq<int>
	const f8 : map<seq<bool>, bool>
	const f9 : int
	const f10 : int
	const f11 : char
	const f12 : multiset<string>
	var f13 : int
	const f14 : seq<array<bool>>
	var f15 : int
	const f16 : bool
	const f17 : int
	const f18 : bool
	constructor (f0 : int, f1 : int, f2 : bool, f3 : int, f4 : bool, f5 : int, f6 : bool, f7 : seq<int>, f8 : map<seq<bool>, bool>, f9 : int, f10 : int, f11 : char, f12 : multiset<string>, f13 : int, f14 : seq<array<bool>>, f15 : int, f16 : bool, f17 : int, f18 : bool) {
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
	}
	
}

class C0 {
	const f19 : int
	const f20 : char
	constructor (f19 : int, f20 : char) {
		this.f19 := f19;
		this.f20 := f20;
	}
	
	function fm2(p0: int, p1: int, p2: int, globalState: GlobalState): int {
		f19
	}
}

function fm0(p0: bool, p1: string, p2: bool, p3: int, globalState: GlobalState): int {
	|map[-469 := false]| * safeDivisionInt(-0x19b, |{|multiset(seq(abs(94), i0  => (0x16e)))|}|)
}
function fm1(p0: bool, p1: int, p2: int, p3: int, globalState: GlobalState): set<int> {
	({|[map v0 : int | (0x352 <= v0) && (v0 < 668) :: (safeModuloInt(v0, -0x364)) := (|(set v1 : char | v1 in map['b' := false] :: (v1))|)]|} + {0xb0}) + ({0xa6} + {-|[-|"henpihgvp"|, |(seq(abs(0x5b), i0  => ('u')))|, 0x3a8]|})
}
function fm3(p0: int, p1: bool, p2: seq<bool>, p3: map<bool, int>, globalState: GlobalState): seq<D1> {
	seq(abs(0x4c), i0  => (DC4(DC3(true, true, false, !!!true))))
}
function fm4(p0: bool, globalState: GlobalState): multiset<bool> {
	multiset{DC11('b', true, multiset{'t'}, true, seq(abs(-0x343), i0  => (multiset{493}))).cf20}
}
function fm5(globalState: GlobalState): char {
	'g'
}
function fm6(p0: int, p1: char, p2: int, globalState: GlobalState): bool {
	match DC7() {
		case DC7() => false
		case DC8(cf14) => false
		case DC6(cf13) => !!false
	}
}
function fm7(p0: bool, p1: int, p2: int, globalState: GlobalState): seq<char> {
	['j', 'x']
}
function fm8(p0: bool, globalState: GlobalState): seq<int> {
	(if (false) then [|{false, false}|] else [|(map v0 : int | (0x3ab <= v0) && (v0 < 0x1d2) :: (v0 * |map[false := false]|) := (false))|, 0x34c, 0x206]) + (seq(abs(681), i0  => (0xf4)))
}
function fm9(p0: bool, p1: int, globalState: GlobalState): D3 {
	DC7()
}
function fm10(p0: int, p1: int, p2: multiset<map<int, int>>, p3: int, globalState: GlobalState): multiset<map<bool, bool>> {
	match DC11('h', false, multiset{'m', 'f', 'p', 'd', 'd'}, true, seq(abs(0x10f), i0  => (multiset{|(map v0 : int | (0x1a5 <= v0) && (v0 < 0x227) :: (v0 - |"dxbtxxe"|) := (|"gi"|))|}))) {
		case DC10(cf16, cf17, cf18) => multiset{map[cf17 := true]}
		case DC11(cf19, cf20, cf21, cf22, cf23) => multiset{map[cf22 := cf22], map[cf20 := cf20]} - multiset([map[true := cf22], map[cf20 := cf22]])
		case DC12(cf24, cf25, cf26, cf27) => multiset{map[cf26 := cf26], map[cf26 := DC10(cf24, false, cf26).cf18]}
		case DC9(cf15) => multiset([map[false := true]] + (seq(abs(0x278), i1  => (map[true := false]))))
	}
}
method m0(p0: int, p1: bool, p2: bool, p3: string, globalState: GlobalState) returns (r0: set<set<int>>, r1: bool, r2: multiset<map<bool, bool>>) {
	if (p1) {
		globalState.f15 := safeModuloInt(p0, safeDivisionInt(p0, 0x27f));
		var v0: seq<int> := [p0, p0, -fm0(p1, p3, p2, p0, globalState), p0 - p0];
		var v1 := 'y';
		var v2: C0 := new C0(p0, v1);
		var v3 := "deyue";
		var v4 := DC0(v3);
		v0, globalState.f6, globalState.f13, v2, v3 := fm8(p2 ==> p2, globalState), (fm7(false, v2.f19, |v0|, globalState) + v3) == (if (!p2) then v3 else v4.cf0), v2.f19, v2, "hkykdidt" + (seq(abs(0x4a), i0  => (v1)));
		var v5: array<int> := new int[8](i1 => safeModuloInt(i1, v2.f19));
		var v6: map<bool, array<int>> := map[p2 := v5];
		var v7: map<bool, bool> := map[p2 := p2];
		var v8: map<array<int>, map<bool, bool>> := map[if (true in v6) then v6[true] else v5 := v7];
		v8 := v8[v5 := v7];
		var v9: map<seq<int>, array<int>> := map[v0 := v5];
		v9 := v9 + v9;
		var v10: array<char> := new char[14] [v1, v1, p3[safeIndex(v2.f19, |p3|)], v2.f20, v2.f20, v2.f20, v1, 'c', v1, v1, v2.f20, v2.f20, v1, v2.f20];
		v10[safeIndex(614, v10.Length)] := v1;
		v10[safeIndex(614, v10.Length)] := v2.f20;
	} else {
		var v11 := 'a';
		globalState.f6 := (if (p2) then v11 else fm5(globalState)) in p3;
		var v12 := DC7();
		v12 := v12;
		var v13: multiset<bool> := multiset{if (p1) then !p1 else p1, false && p1};
		v13 := fm4(p2, globalState);
		globalState.f13 := -p0;
		var v14: array<int> := new int[18](i2 => safeModuloInt(i2, p0));
		var v15: array<bool> := new bool[19] [p2, !p1, p2, p2, p1, p1, p1, p2, p2, true, p2, !p1, p2, false, p2, p1, p2, p2, p1];
		var v16: seq<seq<char>> := [seq(abs(0x2b3), i3  => (v11)), p3];
		var v17 := DC1(v14, v15, multiset(v16[safeIndex(p0, |v16|) := [v11]]), p0, p0);
		v14[safeIndex(248, v14.Length)] := v17.cf4;
		var v18 := DC3(p2, p1, true, fm6(p0, fm5(globalState), -p0, globalState));
		v14[safeIndex(248, v14.Length)] := -|multiset{p1, v18.cf8}|;
	}
	
	var v19: array<int> := new int[20];
	forall i4 | 0 <= i4 < v19.Length {
		v19[i4] := safeModuloInt(i4, p0);
	}
	for i5 := p0 to p0 {
		var v20: array<bool> := new bool[13];
		var v21: set<array<bool>> := {v20};
		v21 := v21;
		var v22: array<map<C0, int>> := new map<C0, int>[3];
		v22 := v22;
		var v23: map<bool, bool> := map[p1 := p1];
		match (fm9(true, |(v23 + v23)|, globalState)) {
			case DC7() =>
				globalState.f6 := true;
				var v24: seq<bool> := [p1];
				var v25: seq<string> := [(p3[safeIndex(p0, |p3|) := 'i'])[safeIndex(i5, |p3[safeIndex(p0, |p3|) := 'i']|) := fm5(globalState)], p3, p3 + p3, fm7(!v24[safeIndex(i5, |v24|)], p0, p0, globalState)];
				v25 := v25 + v25;
				var v26: map<string, int> := map[p3 := p0];
				var v27: map<bool, map<string, int>> := map[p2 || p1 := v26];
				v27 := v27[p1 := v26 + map[p3 := p0]];
				var v28 := 'w';
				var v29: C0 := new C0(p0, v28);
				var v30 := DC9(v29);
				var v31: array<C0> := new C0[19] [v29, v29, v29, v29, v30.cf15, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29, v29];
				var v32: set<array<C0>> := {v31, v31, v31, v31, v31};
				v32 := {if (p2) then v31 else v31};
			case DC8(cf14) =>
				var v33 := 'c';
				var v34: map<char, bool> := map[v33 := p1];
				var v35: map<array<bool>, map<char, bool>> := map[v20 := v34];
				var v36: seq<map<array<bool>, map<char, bool>>> := [if (p1) then map[v20 := v34] else v35];
				var v37: C0 := new C0(cf14, v33);
				var v38: map<C0, int> := map[v37 := i5];
				v35, r1 := v36[safeIndex(cf14, |v36|)], |v38[v37 := cf14]| < 0x38c;
				var v39: map<int, int> := map[0x1d6 := |DC0(p3).cf0|];
				r1 := if (p1 <==> true) then v37.f19 > |v39| else p2;
				v20 := v20;
				globalState.f15 := cf14;
			case DC6(cf13) =>
				var v40 := 't';
				var v41 := new C0(p0, v40);
				var v42: seq<int> := [p0, |p3|, i5];
				var v43: map<seq<int>, C0> := map[v42 := v41];
				var v44: multiset<char> := multiset{v41.f20, v41.f20};
				var v45: multiset<int> := multiset{|{!p2}|, 0x148};
				var v46: seq<multiset<int>> := [v45, v45];
				var v47 := DC11(v41.f20, p1, v44, p1, v46);
				r1, v43 := v47.cf22, map[v42 := v41];
				globalState.f5 := -((v41.f19 * v41.f19) + i5);
				v42 := (v42 + v42) + (if (p2) then v42 else [v41.f19]);
		}
		
		var v48: map<int, bool> := map[p0 := p1];
		var v49 := 'c';
		v19[safeIndex(518, v19.Length)] := safeDivisionInt(i5, fm0(p1, p3, fm6(|v48|, v49, |[p2]|, globalState), p0, globalState));
		v19[safeIndex(518, v19.Length)] := -p0;
	}
	var v50: array<string> := new string[25](i6 => p3);
	v50[safeIndex(296, v50.Length)] := "d";
	v50[safeIndex(296, v50.Length)] := "arixobu";
	var v51: array<bool> := new bool[3];
	forall i7 | 0 <= i7 < v51.Length {
		v51[i7] := p1;
	}
	var v52 := 'x';
	var v53: C0 := new C0(p0, v52);
	v53 := v53;
	var v54: set<int> := {v53.f19, p0};
	var v56: set<set<int>> := {v54, v54 + (set v55 : int | (-0x91 <= v55) && (v55 < 464) :: (safeDivisionInt(v55, |[p2, p1, true, p1, p1]|))), fm1(p1, v53.f19, p0, 503, globalState)};
	r0 := v56;
	r1 := p1;
	var v57 := DC10(p0, p2, p1);
	var v58: map<int, int> := map[p0 := 0x33c];
	var v60: multiset<map<int, int>> := multiset{v58, map v59 : int | (335 <= v59) && (v59 < 0x393) :: (v59 - |v54|) := (v53.f19)};
	r2 := fm10(p0, 0x1b3 - v57.cf16, v60, p0, globalState);
}
method Main() {
	var v0 := 0x147;
	var v1 := "qljvqw";
	var v2 := true;
	var v3: seq<bool> := [v2];
	var v4: multiset<string> := multiset{v1, v1, v1};
	var v5: array<bool> := new bool[29](i0 => v2);
	var globalState := new GlobalState(624, 0x2e7, false, -990, false, 0x391, false, [v0, |v1|, v0, -v0], map[v3 := v2], 580, 0x287, 'r', v4, 0x1e1, [v5, v5], 31, true, -0x33e, false);
	var v6 := 'm';
	var v7: seq<int> := [-0x15d];
	var v8: set<int> := {v0, v7[safeIndex(-857, |v7|)], v0, v0, v0};
	var v9: multiset<set<int>> := multiset{fm1(v2, -|[true, v2, true, v2, v2]|, v0, -v0, globalState), v8};
	var v11: map<int, int> := map[v0 := -v0];
	var v12: array<int> := new int[14] [|map[v0 := v2]|, safeModuloInt(v0, v0), v0, -0x1f2, |v1[safeIndex(fm0(v2, v1, v2, |map[v2 := v6]|, globalState), |v1|) := 'q']|, v0, v0, if ((set v10 : int | (0xda <= v10) && (v10 < 0x1c3) :: (safeDivisionInt(v10, -|v8|))) in v9) then v9[set v10 : int | (0xda <= v10) && (v10 < 0x1c3) :: (safeDivisionInt(v10, -|v8|))] else 157, v0, v0 - v0, v0, safeDivisionInt(v0, v0), safeModuloInt(-v0, v0), -(if (v2) then |v11| else -v0)];
	forall i1 | 0 <= i1 < v12.Length {
		v12[i1] := safeDivisionInt(i1, v0);
	}
	var v13: multiset<int> := multiset{v0, v0};
	v6, globalState.f5, globalState.f6, v12 := v6, v0, v13 <= v13, v12;
	var v14, v15, v16 := m0(v0, false, v2, v1, globalState);
	globalState.f5 := v0;
	v5[safeIndex(514, v5.Length)] := !v15;
	v5[safeIndex(514, v5.Length)] := v15;
	var v17: array<string> := new string[10](i2 => v1);
	v17[safeIndex(564, v17.Length)] := v1;
	v17[safeIndex(564, v17.Length)] := DC0(v1).cf0;
	var i3 := 0;
	while (!v15)
		decreases 100 - i3
	{
		if (i3 >= 100) {
			break;
		}
		
		i3 := i3 + 1;
		var v18 := DC2(v6);
		var v19 := new C0(v0, v18.cf6);
		v12[safeIndex(686, v12.Length)] := |"ofwu"| + |v1|;
		var v20: map<bool, int> := map[v2 := |"oijvun"|];
		v12[safeIndex(686, v12.Length)] := -|fm3(if (false) then if (v2 in v20) then v20[v2] else -524 else v19.f19, v5[safeIndex(514, v5.Length)] || v5[safeIndex(514, v5.Length)], v3, v20, globalState)|;
		var v21: map<bool, bool> := map[v5[safeIndex(514, v5.Length)] := !v15];
		match (DC3(v15, true !in v3, v15, v7[safeIndex(fm0(v15, v17[safeIndex(564, v17.Length)], v15, v0, globalState), |v7|) := v19.f19] <= [|v21|])) {
			case DC3(cf7, cf8, cf9, cf10) =>
				globalState.f5 := v19.f19;
				var v22, v23, v24 := m0(safeDivisionInt(v12[safeIndex(686, v12.Length)], v12[safeIndex(686, v12.Length)]), v5[safeIndex(514, v5.Length)], v5[safeIndex(514, v5.Length)], v17[safeIndex(564, v17.Length)] + v1, globalState);
				var v25 := new C0(v0, v6);
				var v26: multiset<C0> := multiset{v19};
				var v27: array<int> := new int[27] [v25.f19, |{v23}|, v19.f19, v12[safeIndex(686, v12.Length)], v19.f19, v19.f19, v19.f19, v12[safeIndex(686, v12.Length)], 0xb3, v0, v0, v19.f19, v19.f19, |v21|, v0, v25.f19, |v17[safeIndex(564, v17.Length)]|, v19.f19, v19.f19, v0, -124, v19.f19, v19.f19, v12[safeIndex(686, v12.Length)], if (|v26| in v13) then v13[|v26|] else v19.f19, v19.f19, -0x381];
				var v28: map<char, array<int>> := map[v19.f20 := v27];
				var v29, v30, v31 := m0(v0, v28 != DC5(v28).cf12, v15, v17[safeIndex(564, v17.Length)], globalState);
			case DC2(cf6) =>
				var v32: seq<seq<int>> := [v7];
				var v33: map<seq<int>, string> := map[v32[safeIndex(v12[safeIndex(686, v12.Length)], |v32|)] := v17[safeIndex(564, v17.Length)]];
				var v34: set<bool> := {!(v33 == v33)};
				var v35: map<C0, int> := map[v19 := v0];
				var v36: multiset<bool> := multiset{v19 in v35, false};
				var v37: seq<C0> := [v19];
				var v38: seq<set<bool>> := [v34];
				v12[safeIndex(686, v12.Length)], v34, v36, v12[safeIndex(686, v12.Length)] := safeModuloInt(|v37[safeIndex(v12[safeIndex(686, v12.Length)], |v37|) := v19]|, v0), v38[safeIndex(|v3|, |v38|)], fm4(v15, globalState), v19.f19;
				v0 := v0;
				v8 := set v39 : int | v39 in v7 :: (safeDivisionInt(v39, -0x22d));
				var v40: array<C0> := new C0[4];
				v40 := v40;
			case DC4(cf11) =>
				globalState.f15 := v12[safeIndex(686, v12.Length)];
				var v41: array<int> := new int[23](i4 => i4 * v0);
				var v42: map<char, array<int>> := map[v19.f20 := v41];
				var v43 := DC5(v42);
				v43 := v43.(cf12 := v42);
				v6 := v19.f20;
				v43 := v43;
		}
		
		var v44, v45, v46 := m0(v19.f19, v2, v15, v17[safeIndex(564, v17.Length)] + v1, globalState);
	}
	forall i5 | 0 <= i5 < v12.Length {
		v12[i5] := i5 - (if (v15) then v0 else v0);
	}
	var v47 := DC3(v5[safeIndex(514, v5.Length)], v2, v2, v15);
	globalState.f6 := !v47.cf7;
	for i6 := v0 to |{v0, v0}| {
		var v48 := DC1(v12, v5, v4, i6, i6);
		match (v48) {
			case DC1(cf1, cf2, cf3, cf4, cf5) =>
				v6 := v1[safeIndex(|(seq(abs(603), i7  => (v6)))|, |v1|)];
				var v49, v50, v51 := m0(cf4, true, v15, v1, globalState);
				globalState.f5 := cf5;
				cf5 := cf4;
			case DC0(cf0) =>
				v3 := v3;
				var v52: array<D0> := new D0[11](i8 => DC0(v1));
				var v53 := DC0(v1);
				v52[safeIndex(106, v52.Length)] := v53;
				cf0, globalState.f13, globalState.f6, globalState.f6, v52[safeIndex(106, v52.Length)] := v17[safeIndex(564, v17.Length)], v0, v5[safeIndex(514, v5.Length)], !false, DC0(cf0 + "fy");
				var v54: map<char, array<int>> := map[v6 := v12];
				var v55 := DC5(v54);
				var v56: map<int, D2> := map[v0 := v55];
				v2 := i6 >= |v56[i6 := v55]|;
				var v57, v58, v59 := m0(0x351, v5[safeIndex(514, v5.Length)] <==> v5[safeIndex(514, v5.Length)], v5[safeIndex(514, v5.Length)], v1, globalState);
		}
		
		var v60: array<seq<int>> := new seq<int>[5](i9 => v7[safeIndex(180, |v7|) := i6]);
		v60 := v60;
		globalState.f6 := !!!v15;
		var v61: C0 := new C0(204, v6);
		var v62: array<D1> := new D1[29](i10 => DC3(v5[safeIndex(514, v5.Length)], v5[safeIndex(514, v5.Length)], true, v5[safeIndex(514, v5.Length)]));
		var v63: map<C0, array<D1>> := map[v61 := v62];
		v0, v5[safeIndex(514, v5.Length)], v0, v17, globalState.f15 := safeDivisionInt(safeDivisionInt(v0, |v63|), i6), v2, v0, v17, 827;
	}
	for i11 := safeDivisionInt(v0, v0) to v0 {
		if (v2) {
			var v64: C0 := new C0(v0, v6);
			var v65: map<bool, C0> := map[v2 := v64];
			var v66: array<map<bool, C0>> := new map<bool, C0>[3] [map[v15 := v64] + v65, v65, v65 + map[!v15 := v64]];
			v66 := v66;
			var v67: array<multiset<bool>> := new multiset<bool>[3](i12 => multiset(v3) - multiset{v15});
			var v68: multiset<bool> := multiset{v15};
			v67[safeIndex(456, v67.Length)] := v68 - multiset{v5[safeIndex(514, v5.Length)], v5[safeIndex(514, v5.Length)], v15, v2};
			v67[safeIndex(456, v67.Length)] := v68;
			v12[safeIndex(807, v12.Length)] := v64.f19;
			var v69: map<int, bool> := map[|v68| := v15];
			v12[safeIndex(807, v12.Length)], globalState.f6 := safeModuloInt(|v69|, i11), v2;
			var v70: seq<array<bool>> := [v5];
			v70 := v70[safeIndex(v12[safeIndex(807, v12.Length)], |v70|) := v5];
			globalState.f5 := safeDivisionInt(i11, i11);
		} else {
			globalState.f6 := v5[safeIndex(514, v5.Length)];
			var v71 := DC1(v12, v5, v4, v0, v7[safeIndex(|v1|, |v7|)]);
			var v72: seq<D0> := [v71];
			v17[safeIndex(564, v17.Length)], globalState.f6, v72, v6 := ("vw")[safeIndex(i11, |"vw"|) := v6] + "pmpe", v15 ==> v15, (v72 + v72) + v72, fm5(globalState);
			globalState.f15 := i11;
			globalState.f6 := v2;
			var v73: map<D0, int> := map[v71 := |v13| + i11];
			v73 := v73[v71 := -v0];
		}
		
		var v74: seq<string> := [v1];
		v74 := v74;
		var v75: set<bool> := {true};
		var v76: set<set<bool>> := {v75};
		var v77: map<char, int> := map[v6 := |v76|];
		v77 := v77[v6 := -v0];
		globalState.f15 := v0 * i11;
	}
	var i13 := 0;
	while (!v15)
		decreases 100 - i13
	{
		if (i13 >= 100) {
			break;
		}
		
		i13 := i13 + 1;
		var v78, v79, v80 := m0(if (|v17[safeIndex(564, v17.Length)]| in v11) then v11[|v17[safeIndex(564, v17.Length)]|] else v0, !v15, v0 != v0, v1, globalState);
		var v81 := new C0(v0, v6);
		var v82: seq<set<int>> := [v8, fm1(v5[safeIndex(514, v5.Length)], |"vityokux"|, |v7|, v0, globalState)];
		var v83: map<bool, bool> := map[[{0x2cf, v81.f19}, v8] <= v82 := !(v0 < v0)];
		v83 := v83[v79 := v5[safeIndex(514, v5.Length)]];
		v12[safeIndex(550, v12.Length)] := v0;
		v12[safeIndex(550, v12.Length)] := v0;
	}
	var v84: set<bool> := {v15, v5[safeIndex(514, v5.Length)]};
	for i14 := |(v84 + v84)| to safeModuloInt(v0, v0) {
		if (v15) {
			v47 := v47;
			globalState.f15 := -(i14 - 881);
			v5[safeIndex(514, v5.Length)] := !!v15;
			globalState.f13 := i14;
			var v85: seq<array<bool>> := [v5];
			var v86 := DC6(v85);
			globalState.f5 := |v86.cf13|;
		} else {
			globalState.f6 := v2;
			v47 := v47;
			globalState.f6 := -0x200 < v0;
			v5[safeIndex(514, v5.Length)] := v84 !! v84;
			var v87: array<C0> := new C0[14];
			var v88: map<int, bool> := map[v0 := v2];
			var v89: map<bool, array<C0>> := map[fm6(-|v88|, v6, -v0, globalState) := v87];
			v87 := if ((v84 > {v2}) in v89) then v89[v84 > {v2}] else v87;
		}
		
		globalState.f15 := safeDivisionInt(safeDivisionInt(v0, i14), |v3|);
		var v90: map<bool, array<bool>> := map[true := DC1(v12, v5, v4, v0, i14).cf2];
		var v91: map<char, seq<char>> := map[v6 := v17[safeIndex(564, v17.Length)]];
		var v92: map<int, bool> := map[i14 := v15];
		var v93 := DC1(v12, if (v5[safeIndex(514, v5.Length)] in v90) then v90[v5[safeIndex(514, v5.Length)]] else v5, multiset{seq(abs(592), i15  => (v6)), if (v6 in v91) then v91[v6] else fm7(v5[safeIndex(514, v5.Length)], i14, fm0(if (fm0(v5[safeIndex(514, v5.Length)], v1, v5[safeIndex(514, v5.Length)], i14, globalState) in v92) then v92[fm0(v5[safeIndex(514, v5.Length)], v1, v5[safeIndex(514, v5.Length)], i14, globalState)] else !v2, "ytn", v15, v0, globalState), globalState), v17[safeIndex(564, v17.Length)], v17[safeIndex(564, v17.Length)]}, i14, v0);
		v93 := v93;
		var v94, v95, v96 := m0(-safeDivisionInt(i14, v0), v15, 0xe8 !in v8, "ckhhe", globalState);
	}
	var v97: array<C0> := new C0[4];
	var v98: C0 := new C0(v0, v6);
	v97[safeIndex(411, v97.Length)] := v98;
	v97[safeIndex(411, v97.Length)] := v98;
	var v99: seq<C0> := [v97[safeIndex(411, v97.Length)]];
	var i16 := 0;
	while (v99 != v99)
		decreases 100 - i16
	{
		if (i16 >= 100) {
			break;
		}
		
		i16 := i16 + 1;
		v17[safeIndex(564, v17.Length)] := v17[safeIndex(564, v17.Length)] + v17[safeIndex(564, v17.Length)];
		var v100, v101, v102 := m0(if (v0 in v13) then v13[v0] else v98.f19, !true <== v15, v2, (v1 + v1[safeIndex(v98.f19, |v1|) := v6])[safeIndex(v98.f19, |(v1 + v1[safeIndex(v98.f19, |v1|) := v6])|) := v6], globalState);
		v101 := v2;
		if (v2) {
			var v103: array<map<map<D3, int>, bool>> := new map<map<D3, int>, bool>[29];
			var v104: seq<array<bool>> := [v5];
			var v105 := DC6(v104);
			var v106: map<D3, int> := map[v105 := v98.f19];
			var v107: map<map<D3, int>, bool> := map[v106 := v15];
			v103[safeIndex(196, v103.Length)] := v107;
			var v108: map<bool, map<map<D3, int>, bool>> := map[v5[safeIndex(514, v5.Length)] := v107];
			v103[safeIndex(196, v103.Length)], globalState.f6 := (if (false in v108) then v108[false] else v107) + v107, 0x8a == 0x3e;
			v17 := v17;
			globalState.f5 := safeModuloInt(v98.f19 + v98.f19, v98.f19);
			var v109 := DC2(v98.f20);
			v109 := v109.(cf6 := v6);
			globalState.f15 := --safeDivisionInt(-v98.f19 + 0x103, v98.fm2(v0, v98.f19, v98.f19, globalState));
		} else {
			v11 := v11[|v3| := v98.f19];
			v7 := fm8(v5[safeIndex(514, v5.Length)], globalState) + v7;
			v1 := v1 + (seq(abs(0x170), i17  => (v98.f20)));
			var v110: array<map<bool, bool>> := new map<bool, bool>[21](i18 => map[v2 := v101]);
			v5[safeIndex(514, v5.Length)], v110 := v98.f19 < v0, v110;
			var v111, v112, v113 := m0(v98.f19, v101, !v5[safeIndex(514, v5.Length)], "svi", globalState);
		}
		
	}
	v97[safeIndex(411, v97.Length)] := v97[safeIndex(411, v97.Length)];
	print v0, "\n";
	print v1, "\n";
	print v2, "\n";
	print v3 == [true], "\n";
	print v4 == multiset{"qljvqw", "qljvqw", "qljvqw"}, "\n";
	print v5[0], "\n";
	print v5[1], "\n";
	print v5[2], "\n";
	print v5[3], "\n";
	print v5[4], "\n";
	print v5[5], "\n";
	print v5[6], "\n";
	print v5[7], "\n";
	print v5[8], "\n";
	print v5[9], "\n";
	print v5[10], "\n";
	print v5[11], "\n";
	print v5[12], "\n";
	print v5[13], "\n";
	print v5[14], "\n";
	print v5[15], "\n";
	print v5[16], "\n";
	print v5[17], "\n";
	print v5[18], "\n";
	print v5[19], "\n";
	print v5[20], "\n";
	print v5[21], "\n";
	print v5[22], "\n";
	print v5[23], "\n";
	print v5[24], "\n";
	print v5[25], "\n";
	print v5[26], "\n";
	print v5[27], "\n";
	print v5[28], "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3, "\n";
	print globalState.f4, "\n";
	print globalState.f5, "\n";
	print globalState.f6, "\n";
	print globalState.f7 == [327, 6, 327, -327], "\n";
	print globalState.f8 == map[[true] := true], "\n";
	print globalState.f9, "\n";
	print globalState.f10, "\n";
	print globalState.f11, "\n";
	print globalState.f12 == multiset{"qljvqw", "qljvqw", "qljvqw"}, "\n";
	print globalState.f13, "\n";
	print |globalState.f14|, "\n";
	print globalState.f15, "\n";
	print globalState.f16, "\n";
	print globalState.f17, "\n";
	print globalState.f18, "\n";
	print v6, "\n";
	print v7 == [-349], "\n";
	print v8 == {327, -349}, "\n";
	print v9 == multiset{{1, 176, 166, -3}, {327, -349}}, "\n";
	print v11 == map[327 := -327], "\n";
	print v12[0], "\n";
	print v12[1], "\n";
	print v12[2], "\n";
	print v12[3], "\n";
	print v12[4], "\n";
	print v12[5], "\n";
	print v12[6], "\n";
	print v12[7], "\n";
	print v12[8], "\n";
	print v12[9], "\n";
	print v12[10], "\n";
	print v12[11], "\n";
	print v12[12], "\n";
	print v12[13], "\n";
	print v13 == multiset{327, 327}, "\n";
	print v14 == {{327}, {327, -29, -28, -27, -26, -25, -24, -23, -22, -21, -20, -19, -18, -17, -16, -15, -14, -13, -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92}, {1, 176, 166, -3}}, "\n";
	print v15, "\n";
	print v16 == multiset{map[false := false]}, "\n";
	print v17[0], "\n";
	print v17[1], "\n";
	print v17[2], "\n";
	print v17[3], "\n";
	print v17[4], "\n";
	print v17[5], "\n";
	print v17[6], "\n";
	print v17[7], "\n";
	print v17[8], "\n";
	print v17[9], "\n";
	print i3, "\n";
	print v47.cf7, "\n";
	print v47.cf8, "\n";
	print v47.cf9, "\n";
	print v47.cf10, "\n";
	print i13, "\n";
	print v84 == {false}, "\n";
	print v97[3].f19, "\n";
	print v97[3].f20, "\n";
	print v98.f19, "\n";
	print v98.f20, "\n";
	print |v99|, "\n";
	print i16, "\n";
}