datatype D0 = DC1(cf1: int, cf2: int) | DC2 | DC0(cf0: int)
datatype D1 = DC4(cf4: bool) | DC3(cf3: bool) | DC5(cf5: D1)
datatype D2 = DC7 | DC6(cf6: seq<bool>)
datatype D3 = DC9(cf8: int, cf9: bool, cf10: int, cf11: bool) | DC10(cf12: char) | DC8(cf7: string) | DC11(cf13: D3)
datatype D4 = DC13(cf15: bool) | DC14(cf16: map<bool, T0>, cf17: bool, cf18: seq<array<bool>>) | DC12(cf14: set<bool>)
datatype D5 = DC16(cf20: int, cf21: multiset<bool>, cf22: seq<bool>, cf23: D3) | DC15(cf19: array<bool>)
datatype D6 = DC18(cf25: string, cf26: bool, cf27: int) | DC17(cf24: C3) | DC19(cf28: D6)
datatype D7 = DC21(cf30: bool, cf31: char, cf32: multiset<int>) | DC22(cf33: int) | DC20(cf29: seq<int>)
datatype D8 = DC24(cf35: int, cf36: array<multiset<bool>>, cf37: array<bool>, cf38: int, cf39: bool) | DC23(cf34: C0) | DC25(cf40: D8)
datatype D9 = DC26(cf41: map<bool, int>)
class GlobalState {
	const f0 : bool
	const f1 : bool
	var f2 : bool
	constructor (f0 : bool, f1 : bool, f2 : bool) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
	}
	
}

function fm3(p0: bool, globalState: GlobalState): int {
	(|map[true := false]| + 0x26a) * (|(map v0 : int | v0 in (seq(0x3b2, i0  => (0x377))) :: (v0 / |"lll"|) := (true))| + |(seq(-226, i1  => ('w')))|)
}
function fm4(p0: bool, p1: int, p2: bool, p3: int, globalState: GlobalState): char {
	match if (!!false) then DC11(DC10('y')) else DC11(DC10('u')) {
		case DC9(cf8, cf9, cf10, cf11) => 'o'
		case DC10(cf12) => cf12
		case DC8(cf7) => 'b'
		case DC11(cf13) => 'h'
	}
}
function fm5(p0: map<int, string>, p1: int, p2: bool, p3: bool, globalState: GlobalState): seq<bool> {
	([false] + [!true]) + [true, true]
}
function fm6(p0: bool, p1: int, p2: bool, globalState: GlobalState): string {
	"nxm" + "uaxq"
}
function fm7(p0: bool, p1: map<bool, int>, p2: char, p3: bool, globalState: GlobalState): D1 {
	if (multiset{0x29c} == multiset{|[-644, |multiset{false, true, !false, false}|]|, |map['u' := true]|, |multiset{map[|[-0x141, 0x66, 0x3ce, 0xfa]| := 'h']}|}) then if (true) then DC4(false) else DC4(false) else if (DC21(false, 'o', multiset{|(seq(-439, i0  => (DC1(-|[0x169]|, |"w"|))))|}).cf30) then DC4(true) else DC4(false)
}
function fm8(p0: char, p1: bool, p2: int, p3: bool, globalState: GlobalState): map<D1, bool> {
	(map[DC4(false) := !!true] + map[DC4(false) := true]) + (map[DC4(true) := true] + map[DC4(true) := true])
}
function fm9(p0: int, p1: bool, p2: int, globalState: GlobalState): set<bool> {
	{false} * {true, true, false, true, true}
}
function fm10(p0: int, p1: bool, p2: int, p3: int, globalState: GlobalState): map<int, bool> {
	map[0x367 := false] + ((map v0 : int | (896 <= v0) && (v0 < 0x15f) :: (v0 * |multiset(seq(-0x18, i0  => (|map[true := false]|)))|) := (true)) + map[0x335 := false])
}
function fm11(p0: int, p1: bool, p2: int, globalState: GlobalState): multiset<int> {
	multiset{|(map[!false := |(map v0 : int | (-0x2f6 <= v0) && (v0 < 774) :: (v0 / |[-|multiset{0x213}|, |[-0x380]|]|) := (|"fixlptlef"|))|] + map[false := 0x31])|, |map[0x33d := |map[0x2ec := false]|]| - 795}
}
function fm12(p0: bool, p1: bool, p2: bool, globalState: GlobalState): bool {
	if (true) then multiset{false} >= multiset{true, true, !!true, true, true} else -|{0x2ec, 26}| >= 0x2f9
}
function fm13(p0: bool, p1: int, p2: int, globalState: GlobalState): multiset<bool> {
	multiset{!!false, [true] < [true, !!false], false}
}
function fm14(p0: int, globalState: GlobalState): map<bool, bool> {
	map[true := false] + (map[true := false] + map[false := false])
}
function fm15(globalState: GlobalState): map<int, map<int, bool>> {
	if (!DC21(true, 's', multiset{0x2f4}).cf30) then map[|"nfbfbvbsb"| := map[0x1c5 := false]] else map[514 := map[-|"jdxlfgub"| := false]] + (map v0 : int | (-0x6f <= v0) && (v0 < 887) :: (v0 % |"g"|) := (map[|[!false]| := !true]))
}
function fm16(p0: bool, p1: int, p2: bool, globalState: GlobalState): seq<map<bool, int>> {
	([map[true := |map[|"uxeiish"| := 755]|]] + (seq(-0x35f, i0  => (map[!false := |[false, false]|])))) + ([map[true := |map[false := |"fgghpfsb"|]|]] + [map[false := |(set v0 : int | (0x158 <= v0) && (v0 < 0xc2) :: (v0 * -608))|], map[false := |{true, !false}|]])
}
function fm17(globalState: GlobalState): multiset<char> {
	multiset{'a', 'u', 'c'}
}
function fm18(p0: int, globalState: GlobalState): D1 {
	match DC8("ycfygtdb") {
		case DC9(cf8, cf9, cf10, cf11) => DC5(DC5(DC4(cf9)))
		case DC10(cf12) => DC5(DC5(DC3(!true)))
		case DC8(cf7) => DC5(DC5(DC5(DC4(false))))
		case DC11(cf13) => DC5(DC5(DC5(DC3(!true))))
	}
}
function fm19(p0: bool, p1: bool, globalState: GlobalState): D1 {
	if (false ==> DC21(false, 'f', multiset{|{map[false := 0x392], map[true := |multiset{|map[true := true]|, 0x163, |"lloh"|, 0x180}|]}|, -0x240}).cf30) then DC3(!!false) else DC3(false)
}
function fm20(p0: int, p1: string, p2: seq<bool>, globalState: GlobalState): seq<string> {
	["cytq", "r"] + (["kcqqyceot"] + ["omkpbmjjq"])
}
function fm21(p0: int, p1: bool, p2: int, p3: int, globalState: GlobalState): map<char, char> {
	map['j' := 'g'] + (map['n' := 'd'] + (map v0 : char | v0 in multiset{'x', 'm'} :: (v0) := ('p')))
}
method m5(p0: seq<int>, p1: set<int>, globalState: GlobalState) returns (r0: bool, r1: map<string, int>) {
	var v0 := "ojcommo";
	var v1: set<string> := {v0};
	var i0 := 0;
	while (0x1c9 > |v1|)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		var v2: array<bool> := new bool[25];
		var v3 := true;
		v2[675] := fm12(v3, true, v3, globalState);
		var v4: array<int> := new int[16](i1 => i1 + 0x16c);
		var v5 := 0x209;
		v4[965] := v5;
		var v6 := DC8(v0);
		v2[675], v4[965], v6 := DC3(v3).cf3, v5, v6;
		if (true) {
			var v7: array<array<int>> := new array<int>[19];
			v7[708] := v4;
			var v8: set<bool> := {v2[675], v3};
			var v10: map<int, bool> := map[|(map v9 : int | v9 in p1 :: (v9 - v4[965]) := (v3))| := v3];
			v7[708] := if (v8 <= v8) then v4 else if (if (p0[v5] in v10) then v10[p0[v5]] else v3) then v4 else v4;
			v4[965] := v4[965];
			v4[965] := v4[965];
			var v11: map<array<bool>, int> := map[v2 := v5];
			v11 := (v11 + v11) + v11;
			v5 := v4[965] / fm3(!false, globalState);
		} else {
			v4[965] := -v4[965];
			var v12: T0 := new C2(v2[675], v5, v4[965]);
			var v13: map<int, T0> := map[v5 := v12];
			var v14: array<T0> := new T0[28] [v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, if (-282 in v13) then v13[-282] else v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12, v12];
			var v15: map<array<T0>, bool> := map[v14 := v2[675]];
			var v16: map<map<array<T0>, bool>, int> := map[v15 := p0[-v4[965]]];
			v16 := v16[v15 + v15 := v12.f4];
			var v17: seq<string> := [v0];
			var v18: array<string> := new string[12] [v0, v17[v4[965]], v0, v0 + v0, v0, seq(-331, i2  => ('u')), v0, v0, v0, v0, v0, v0];
			v18[86] := "owumuar";
			globalState.f2, v2[675], v12.f5, v18[86] := v3, v3, v4[965], v0;
			var v19 := new C1(-v12.f5 % v12.f5, v12.f5);
			var v20 := new C2(|([v3])[|"cbblxtt"| := v3]| != v12.f4, v5, v12.f5);
		}
		
		var v21: map<map<int, int>, bool> := map[map[|multiset{v3, false}| := -512] := v3];
		var v22: map<int, int> := map[fm3(v2[675], globalState) := v5];
		if (if (v22 in v21) then v21[v22] else v3) {
			var v23: seq<bool> := [v2[675], v3, v2[675], v2[675], v2[675]];
			var v24: map<seq<string>, bool> := map[fm20(v4[965], v0, v23, globalState) := v2[675]];
			var v25: seq<string> := [v0, v0, v0];
			v24 := v24[v25 := v23[v4[965]]];
			v2[675] := false;
			v4[965], v4[965] := v4[965], p0[v4[965]];
			var v26 := new C1(v5 * v4[965], v4[965]);
			var v27 := new C3(v2[675]);
		} else {
			v0 := v0;
			var v28: multiset<bool> := multiset{v2[675]};
			var v29: seq<bool> := [v2[675]];
			v4[965], globalState.f2, v4[965] := v4[965] / -|v28|, v29[v4[965] % v4[965]], 0x1e7;
			var v30: map<bool, int> := map[v3 := -v4[965]];
			var v31 := DC26(v30);
			v30 := v30[v3 := v5] + v31.cf41;
			var v32 := 'f';
			var v33: map<char, char> := map[fm4(true, v4[965], fm12(v2[675], v3, v3, globalState), |v0|, globalState) := v32];
			var v34: map<bool, bool> := map[!false := v3];
			var v35: multiset<map<char, char>> := multiset{v33, fm21(v5, if (v3 in v34) then v34[v3] else v3, |v34|, v5, globalState), v33, map[v32 := v32]};
			var v36: C2 := new C2(v2[675], v5, v5);
			var v37: seq<C2> := [v36];
			v35, v5, v4[965] := v35 + v35[v33 := |v37|], -v5, -0x213 + (--v5 + v4[965]);
			var v38: map<bool, multiset<char>> := map[DC13(v2[675]).cf15 := multiset{'i'}];
			v38 := v38[v3 := multiset(v0)];
		}
		
		v5 := v4[965];
	}
	globalState.f2 := false;
	var v39 := false;
	r0 := v39;
	var v40: array<array<bool>> := new array<bool>[4];
	v40 := v40;
	var v41 := -0x1f0;
	var v42: set<bool> := {v39, v39};
	for i3 := v41 to v41 % |v42| {
		var v43: multiset<bool> := multiset{v39};
		var v44: map<int, multiset<bool>> := map[-i3 := v43];
		var v45 := DC18(v0, v39, -v41);
		var v46: set<multiset<bool>> := {if (v45.cf27 in v44) then v44[v45.cf27] else v43};
		v46 := (v46 + {v43}) - v46;
		var v47: array<C1> := new C1[26];
		v47, globalState.f2 := v47, if (!v39) then v39 else v39;
		var v48: map<int, string> := map[v41 := v0];
		var v49 := 'l';
		v48 := v48[i3 := v0[v41 := v49]];
		var v50: array<int> := new int[15];
		var v51: seq<bool> := [v39, v39];
		v50[548] := |v51| - i3;
		var v52: multiset<char> := multiset{'p'};
		var v53 := DC9(|p0|, fm12(v39, v39, v39, globalState), |v52|, v39);
		v50[548] := (v53.(cf11 := false)).cf8;
	}
	globalState.f2 := v39;
	var v54: map<int, bool> := map[v41 := !v39];
	var v55 := DC0(|v54|);
	r0 := match v55.(cf0 := v41) {
		case DC1(cf1, cf2) => v39
		case DC2() => v39
		case DC0(cf0) => v39
	};
	var v56 := 'e';
	var v57: map<string, int> := map[((v0 + v0)[v41 := v56])[v41 := v56] := fm3(v39, globalState)];
	r1 := v57;
}
trait T0 {
	const f4 : int
	var f5 : int
	method m1(p0: array<int>, p1: char, globalState: GlobalState) returns (r0: bool, r1: bool, r2: bool) 
}

class C0 {
	const f7 : multiset<char>
	constructor (f7 : multiset<char>) {
		this.f7 := f7;
	}
	
}

class C1 extends T0 {
	constructor (f4 : int, f5 : int) {
		this.f4 := f4;
		this.f5 := f5;
	}
	
	function fm1(p0: bool, p1: int, p2: set<bool>, p3: bool, globalState: GlobalState): bool {
		true
	}
	function fm2(p0: bool, p1: int, p2: bool, globalState: GlobalState): bool {
		(multiset{f4} - multiset{|map[false := false]|}) > (multiset{|{|[false]|}|} + multiset{|map[true := f4]|})
	}
	method m1(p0: array<int>, p1: char, globalState: GlobalState) returns (r0: bool, r1: bool, r2: bool) {
		var v0 := "ffyidwc";
		var v1 := true;
		var v2: seq<bool> := [v1];
		f5 := |(v0 + "ykiwthwa")| % -|{v2[f5]}|;
		var v3: set<bool> := {!v1};
		r2 := (v3 - v3) !! v3;
		f5 := f5;
		r1 := v1;
		var v4: map<int, bool> := map[f4 := !true];
		if (!(|(v4 + v4)| >= (|v3| - f5))) {
			r0, r1 := |v0| != f4, f5 >= f5;
			var v5: array<bool> := new bool[8];
			var v6: seq<array<bool>> := [v5, v5, v5, v5];
			p0[656] := |v6| / f5;
			var v7: seq<int> := [f4];
			var v8: map<seq<int>, int> := map[v7[f4 := f5] + [|v0|] := f5];
			var v9 := DC0(|map[true := |"khdwjltpl"|]|);
			p0[656], f5, v8 := f4, -v9.cf0 / f4, v8;
			f5 := f4;
			var v10 := DC3(v1);
			var v11: map<bool, int> := map[v1 := f4];
			var v12: multiset<int> := multiset{|v11|};
			var v13, v14, v15, v16 := m4(v10.cf3, if (0x171 in v12) then v12[0x171] else fm3(v1, globalState), seq(0x2aa, i0  => (p1)), globalState);
			var v17: map<int, int> := map[v13 % v13 := f5 + 0x3a5];
			var v18: multiset<bool> := multiset{false};
			v17 := v17[f4 % f5 := |(multiset{true, v1, v1, true, v1} * v18[v1 := f5])|];
		} else {
			var v19: multiset<char> := multiset{'i', p1};
			var v20 := new C0(v19);
			var v21: array<char> := new char[10];
			v21[917] := p1;
			var v22: seq<int> := [f4];
			v21[917] := fm4(fm2(!v1, f4, v1, globalState), f5, v1, f5 / v22[f4], globalState);
			r2 := !true;
			var v23: seq<seq<int>> := [v22];
			r0 := v23 != (seq(0x152, i1  => (v22)));
			p0[339] := f5;
			p0[339] := --fm3(v1, globalState);
		}
		
		var v24 := DC3(v1);
		var v25: map<bool, bool> := map[true := v1];
		var v26: seq<int> := [160];
		var v27: multiset<int> := multiset{f5, f5, |v26|, fm3(v1, globalState)};
		var v28: array<bool> := new bool[25] [if (v1) then v1 else fm2(v1, f5, v1, globalState), v1 ==> v1, v1, v1, v1, !!v24.cf3, fm1(v1, f5, {false, v1}, if (v1 in v25) then v25[v1] else v1, globalState), v1, v1, v1, v1, v1, v27 <= v27, v1, !(fm3(v1, globalState) != 0xfc), v1, v1, true || v1, v1 <== v1, v1, v1, v1, v1, v24.cf3, v2[0x261]];
		v28[948] := !v1;
		v28[948] := v1;
		r0 := v28[948];
		var v29: map<int, int> := map[f5 := f5];
		r1 := |v29| <= -0x3c7;
		r2 := if (v28[948]) then v1 else v1;
	}
	method m3(p0: seq<bool>, globalState: GlobalState) returns (r0: int, r1: int) {
		var v0 := true;
		globalState.f2 := v0;
		if (v0) {
			var v1: multiset<char> := multiset{'x'};
			var v2 := new C0(v1);
			var v3 := new C0(v2.f7);
			var v4: array<bool> := new bool[25];
			v4[205] := !v0;
			v4[205] := v0;
			v3 := v3;
			var v5 := DC0(f5);
			match (v5) {
				case DC1(cf1, cf2) =>
					var v6: seq<int> := [cf2];
					var v7: seq<bool> := [|v6| >= f5];
					var v8 := "omrnrrxhq";
					v7 := ["hyq" == v8];
					var v9: set<bool> := {v0, v4[205]};
					var v10: map<bool, bool> := map[true := v9 > v9];
					v4[205] := if (!v4[205] in v10) then v10[!v4[205]] else v4[205];
					var v11: map<bool, seq<int>> := map[fm2(v4[205], -0x48, v4[205], globalState) := [0x86, 0x11f]];
					v6 := if (v0 in v11) then v11[v0] else v6;
					globalState.f2 := v0;
				case DC2() =>
					var v12: array<D0> := new D0[21];
					v12[198] := DC2();
					var v13: map<bool, bool> := map[v4[205] ==> v0 := v0];
					var v14: set<bool> := {v0};
					var v16 := "cxo";
					var v17: map<int, string> := map[f4 := v16];
					var v18: array<seq<bool>> := new seq<bool>[22] [p0, p0, [v0, v4[205]] + p0[f4 := v0], p0, [v0], [v4[205], false], p0[f5 := fm1(v0, -f4, v14, v0, globalState)], p0 + p0, p0, [v4[205], v4[205]], [p0[f5]], [v0], p0, fm5(map v15 : int | (-0x168 <= v15) && (v15 < 0xfa) :: (v15 + f5) := ("dfrfshin"), f5, v0, v4[205], globalState), p0, fm5(v17, f5, v4[205], v0, globalState) + p0, p0, p0, p0, p0, p0, p0];
					v18[450] := p0 + p0;
					var v19: map<int, int> := map[-f4 := f5];
					var v20 := DC2();
					var v21: map<int, bool> := map[f4 := v4[205]];
					globalState.f2, v12[198], v13, v18[450], r1 := f5 <= |v19|, v20, v13, p0, -((v5.cf0 % f5) + (|v21| % f4));
					v19 := v19[f5 % |(set v22 : int | (0x32e <= v22) && (v22 < 0x101) :: (v22 + f5))| := f5 / f4];
					var v23 := DC3(v4[205] <== v0);
					v23 := v23;
					f5 := fm3(v0, globalState) - f4;
				case DC0(cf0) =>
					v5 := v5;
					var v24: set<bool> := {v4[205]};
					f5 := fm3(v4[205] !in v24, globalState);
					var v25: map<bool, int> := map[v4[205] := f4];
					v25 := v25[!v4[205] := f4];
					var v26: multiset<int> := multiset{f5};
					var v27 := "x";
					r0 := -(if (v0) then f5 else |v26|) + |v27|;
			}
			
		} else {
			var v28 := 'e';
			var v29: array<char> := new char[3] [v28, v28, v28];
			v0, globalState.f2, r0, v29, r1 := v0, f4 == f4, fm3(p0 < p0, globalState), v29, f4;
			var v30: multiset<char> := multiset{v28};
			var v31 := new C0(v30);
			var v32: map<bool, int> := map[false := f5];
			v32 := v32[v0 := f5];
			var v33: array<array<int>> := new array<int>[24];
			v33 := v33;
			globalState.f2 := v0;
		}
		
		var v34 := 'b';
		var v35: seq<char> := [v34];
		var v36 := new C0(multiset(v35));
		var v37: array<char> := new char[21](i0 => v34);
		var v38: map<char, array<char>> := map[v34 := v37];
		var v39: seq<array<char>> := [if (v34 in v38) then v38[v34] else v37, v37, v37];
		if ((if (v0) then v39 else v39) == [v37, v37]) {
			var v40: array<int> := new int[19];
			var v41: map<bool, array<int>> := map[v0 := v40];
			v41 := v41 + map[v0 := v40];
			var v42: seq<int> := [|p0|, |"dgn"|];
			globalState.f2 := (f5 / |v42|) < (f5 - f5);
			var v43 := new C0(multiset(v35));
			var v44 := DC1(f5 + f5, f5);
			match (v44) {
				case DC1(cf1, cf2) =>
					v40[379] := f5;
					v40[379] := cf2;
					var v45, v46, v47, v48 := m4(v0, f4, fm6(v0, f5, v0, globalState), globalState);
					v41 := v41;
					var v49: array<int> := new int[27](i1 => i1 * cf1);
					v49 := v49;
				case DC2() =>
					var v50: map<bool, bool> := map[v0 <==> v0 := v0];
					v50 := v50[!v0 := v0];
					v44 := v44;
					v0 := v0;
					v50 := v50[v0 := false];
				case DC0(cf0) =>
					globalState.f2 := v0;
					var v51: map<bool, int> := map[false := f4];
					var v52 := DC4(v0);
					var v53: set<D1> := {fm7(v0, v51, v34, v0, globalState), v52, fm7(v0, v51, v34, v0, globalState), v52, v52.(cf4 := v0)};
					v53 := set v54 : D1 | v54 in fm8(v34, v0, cf0, v0, globalState) :: (v54);
					var v55: array<bool> := new bool[10](i2 => v0);
					var v56: map<int, char> := map[cf0 := v34];
					v34, v34, v0, v55 := if (f4 in v56) then v56[f4] else v34, v34, !v0 <== (v0 <== v0), v55;
					globalState.f2 := 'u' in v35;
			}
			
			r0 := v42[-(|multiset{v0, v0}| * f4)];
		} else {
			var v57: array<bool> := new bool[25](i3 => v0 ==> v0);
			var v58: array<map<bool, char>> := new map<bool, char>[4];
			var v59: set<bool> := {v0};
			v57, globalState.f2, v58, v0 := v57, (v59 - v59) != v59, v58, fm2(v0, f5, v0, globalState);
			if (true) {
				var v60: seq<multiset<char>> := [v36.f7];
				var v61 := new C0(v60[f5]);
				var v62: array<array<int>> := new array<int>[26];
				v62 := new array<int>[22];
				var v63 := new C0(v36.f7);
				r0 := f4 + f4;
				var v64 := DC0(f5);
				r0 := v64.cf0;
			} else {
				var v65: array<int> := new int[5];
				v65[826] := f5;
				v65[826] := -f4;
				globalState.f2 := !!v0;
				v57 := v57;
				v65 := v65;
				globalState.f2 := v35 == v35;
			}
			
			var v66: multiset<bool> := multiset{v0};
			var v67: seq<string> := [v35];
			var v68: map<int, bool> := map[|fm6(v0, f4, v0, globalState)| := false];
			var v69: set<map<int, bool>> := {v68, v68, v68};
			var v70: map<bool, bool> := map[v66 !! v66 := !(fm6(!v0, f5, v0, globalState) < v67[|v69|])];
			v70 := v70[v0 := v0];
			var v71: multiset<seq<bool>> := multiset{p0};
			globalState.f2, r0 := v0 && v0, -(if (p0 in v71) then v71[p0] else f5) - 215;
			f5 := f4 % |(if (v0) then v35 else v35[f4 := v34])|;
		}
		
		var v72: multiset<bool> := multiset{v0};
		v72 := v72;
		var v73: map<string, bool> := map[if (v0) then v35 else v35 := v0];
		v73 := v73[(seq(584, i4  => ('x'))) + "pu" := v0];
		var v74: map<bool, char> := map[v0 := v34];
		r0 := (|v74| - f4) / f4;
		r1 := f4;
	}
	method m4(p0: bool, p1: int, p2: string, globalState: GlobalState) returns (r0: int, r1: seq<int>, r2: multiset<int>, r3: seq<bool>) {
		var v0: seq<int> := [p1];
		var v1: multiset<int> := multiset{|{v0, seq(0x292, i0  => (p1))}|};
		var v2: seq<bool> := [v1 >= v1];
		globalState.f2 := v2[f5];
		var v3 := DC3(p0);
		var v4: set<D1> := {DC5(v3)};
		var v5 := DC5(v3);
		v4 := {v5, v5};
		var v6: multiset<bool> := multiset{p0};
		globalState.f2 := fm1(p0, f5, fm9(p1, p0, |v6|, globalState), p1 >= f5, globalState);
		if (p0) {
			var v7: map<bool, seq<int>> := map[p0 := v0];
			var v8: map<bool, int> := map[p0 := p1];
			var v9: set<bool> := {p0};
			var v10 := DC4(p0);
			var v11: seq<set<bool>> := [{v10.cf4, p0}, v9, v9];
			var v12: multiset<seq<D1>> := multiset{seq(0x1af, i1  => (DC4(p0)))};
			var v13: seq<D1> := [v10, v10];
			var v14: array<int> := new int[21] [p1, -(p1 + -v0[0x138]), |v7| + f5, |v8|, -|(v9 * v11[p1])|, -p1, f4, p1, if (v13 in v12) then v12[v13] else f4, f5, p1, f5, f5, -836, fm3(fm1(!p0, fm3(p0, globalState), v9, p0, globalState), globalState), 0xfe, f4, -(f5 * |[p0]|), p1, v0[f5], |(v0 + v0)|];
			v14[852] := p1 - f5;
			var v15: array<char> := new char[15](i2 => if (p0) then 'v' else 'f');
			var v16 := 's';
			v15[206] := v16;
			var v17: map<int, bool> := map[f4 := p0];
			v14[852], v15[206], globalState.f2, globalState.f2, v17 := -f4, fm4(p0, -652, p0 ==> false, fm3(!false, globalState), globalState), p0 ==> p0, !p0, if (p0) then v17 else v17 + fm10(f5, p0, f4, p1, globalState);
			var v18 := DC1(f5, v14[852]);
			v18 := v18;
			r0 := (-0x2a4 % p1) % f4;
			v14[852] := f4;
			globalState.f2 := p0;
		} else {
			globalState.f2 := (v1 - v1) >= fm11(-0x3d5, p0, --f5, globalState);
			f5 := (f4 - -p1) - f5;
			globalState.f2 := p1 < 0xc2;
			var v19: map<bool, bool> := map[p0 := p0];
			if (fm2(p0, f5, if (p0 in v19) then v19[p0] else p0, globalState)) {
				var v20: array<T0> := new T0[10];
				v20 := v20;
				var v21: array<seq<D0>> := new seq<D0>[11];
				r0, v21 := p1, v21;
				r0 := p1;
				v0 := v0;
				var v22: array<int> := new int[18](i3 => i3 * p1);
				f5, v22, globalState.f2, f5 := f4, v22, p0, -(if (v1 !! v1) then -0xc9 else -|(v6 * v6)|);
			} else {
				r0 := |(v2 + v2)|;
				var v23: array<seq<int>> := new seq<int>[8];
				v23[645] := v0;
				v23[645] := v0[fm3(p0, globalState) := p1];
				var v24 := DC0(f5);
				f5 := v24.cf0;
				var v25: array<bool> := new bool[3](i4 => p0);
				var v26: map<seq<array<bool>>, int> := map[[v25, v25] := p1];
				var v27: map<bool, array<bool>> := map[p0 := v25];
				var v28: seq<array<bool>> := [if (false in v27) then v27[false] else v25, v25, v25, v25, v25];
				var v29: seq<seq<array<bool>>> := [v28];
				globalState.f2 := p1 < (if (v29[f4] in v26) then v26[v29[f4]] else |p2|);
				var v30: set<bool> := {p0};
				v25[805] := fm1(!fm2(false, p1, p0, globalState), -p1, v30, p0, globalState);
				v25[805] := p0;
			}
			
			r0 := f4 / f4;
		}
		
		globalState.f2 := p0;
		var v31: multiset<string> := multiset{"k"};
		v31 := multiset{"nxhuqgf", p2};
		r0 := f4;
		r1 := seq(0x2cb, i5  => (f4));
		r2 := v1;
		r3 := v2;
	}
}

class C2 extends T0 {
	const f6 : bool
	constructor (f6 : bool, f4 : int, f5 : int) {
		this.f6 := f6;
		this.f4 := f4;
		this.f5 := f5;
	}
	
	function fm0(p0: bool, p1: multiset<int>, p2: int, p3: string, globalState: GlobalState): seq<map<bool, bool>> {
		[map[f6 := f6] + map[f6 := f6]]
	}
	method m1(p0: array<int>, p1: char, globalState: GlobalState) returns (r0: bool, r1: bool, r2: bool) {
		var v0: array<bool> := new bool[15];
		forall i0 | 0 <= i0 < v0.Length {
			v0[i0] := f6;
		}
		var v1 := 'o';
		v1 := v1;
		f5 := f5;
		f5, globalState.f2 := f5, f6;
		var v2 := new C0(multiset{v1});
		r0 := f6;
		var v3 := DC4(f6);
		r0 := !v3.cf4;
		var v4: seq<bool> := [fm12(f6, f6, f6, globalState)];
		r1 := v4[f4];
		r2 := f6;
	}
	method m2(p0: int, globalState: GlobalState) returns (r0: bool, r1: int) {
		if (false) {
			r0 := f6;
			if (f6 && f6) {
				r0 := !f6;
				var v0: seq<bool> := [f6];
				var v1: multiset<seq<bool>> := multiset{v0};
				var v2 := DC6(v0);
				var v3: map<int, bool> := map[p0 := f6];
				var v4 := "quepdvot";
				var v5 := DC8(v4);
				var v6: set<bool> := {f6};
				var v7 := 'v';
				var v8: array<int> := new int[18] [f5, if ([!f6] in v1) then v1[[!f6]] else fm3(f6, globalState), f5, |v2.cf6|, f4, |v3|, f4, p0, f5, f5, -265, |v5.cf7[|v6| := v7]|, f4, p0, f4, p0, f5, f4];
				var v9: array<bool> := new bool[14];
				var v10: map<array<int>, array<bool>> := map[v8 := v9];
				var v11 := DC3(f6);
				var v12: map<map<array<int>, array<bool>>, bool> := map[v10 := !(f6 <==> fm12(v11.cf3, f6, f6, globalState))];
				var v13: map<bool, bool> := map[f6 := f6];
				v12 := v12[v10[v8 := v9] := if (f6 in v13) then v13[f6] else f6];
				v3 := v3[f5 := f6];
				f5 := p0;
				var v14: multiset<char> := multiset{v7, v7};
				var v15 := new C0(v14 - multiset{v7});
			} else {
				var v16 := 'v';
				var v17: multiset<char> := multiset{v16};
				var v18 := new C0(v17);
				globalState.f2 := f6;
				v16 := v16;
				var v19 := "ngmxo";
				var v20 := DC8(v19);
				v19 := v20.cf7;
				v19 := fm6(f6, f5, !f6, globalState);
			}
			
			f5 := -fm3(f6, globalState);
			var v21: multiset<bool> := multiset{f6, f6, f6, f6, f6};
			var v22: seq<int> := [f4];
			var v23 := 't';
			var v24: multiset<char> := multiset{v23, v23, v23};
			var v25: C0 := new C0(v24);
			var v26 := DC1(f4, 469);
			var v27 := DC0(p0);
			var v28 := DC4(f6);
			var v29 := DC9(p0, f6, f5, v28.cf4);
			var v30: array<int> := new int[19] [0x164, f5, |v21|, fm3(f6, globalState), f5, f4, -v22[|[v25]|], f5, v26.cf1 + f5, 0xc3, f4, p0, v26.cf2, v27.cf0, v22[-p0], f5, v29.cf10, f5, f4];
			v30[619] := f4;
			v30[619] := f4;
			var v31: seq<bool> := [f6, f6];
			var v32: map<array<int>, seq<bool>> := map[v30 := v31];
			v21 := multiset(if (v30 in v32) then v32[v30] else v31);
		} else {
			globalState.f2 := f6;
			var v33 := 'e';
			var v34: multiset<char> := multiset{v33};
			var v35: C0 := new C0(v34);
			var v36 := DC4(f6);
			var v37: array<D1> := new D1[10] [v36, v36, v36, v36, v36, v36, v36, v36.(cf4 := f6), v36, DC4(f6)];
			var v38: array<int> := new int[24];
			var v39: seq<array<int>> := [v38];
			v35, v37, v38, globalState.f2, f5 := v35, v37, v39[f4 - fm3(f6, globalState)], f6 && f6, -(p0 * f4) - f4;
			r0 := !!f6;
			var v40: set<bool> := {f6};
			globalState.f2 := v40 != v40;
			var v41: T0 := new C1(p0, f4);
			var v42: array<T0> := new T0[7] [if (f6) then v41 else v41, v41, v41, v41, v41, v41, v41];
			v42[386] := v41;
			v42[386] := v41;
		}
		
		var v43: array<map<array<int>, char>> := new map<array<int>, char>[2];
		var v44: array<int> := new int[15](i0 => i0 % p0);
		var v45 := 'q';
		var v46: map<array<int>, char> := map[v44 := v45];
		v43[663] := v46;
		var v47: map<seq<int>, map<array<int>, char>> := map[seq(-0x2be, i1  => (f4)) := v46];
		var v48 := "spwqw";
		var v49: map<bool, string> := map[f6 := v48];
		var v50: seq<int> := [-|"qsprbn"|, p0, -|(if (f6 in v49) then v49[f6] else v48)|, 0x2bf, f4];
		v43[663] := if (v50 in v47) then v47[v50] else v46;
		if (f6) {
			var v51: seq<bool> := [f6];
			var v52: array<bool> := new bool[10] [fm12(f6, fm12(f6, f6, f6, globalState), f6, globalState), f6, (DC3(f6).(cf3 := f6)).cf3, true, f6, v51[p0], f6, f6, f6, !f6 <== f6];
			var v53 := DC4(f6);
			v52[29] := v53.cf4;
			r1, v52, v52[29] := f5, v52, f6;
			if (v52[29]) {
				var v54: multiset<char> := multiset{v45};
				var v55: seq<multiset<char>> := [v54];
				var v56: C0 := new C0(v55[0x3c0]);
				var v57: seq<C0> := [v56, v56];
				var v58: map<bool, C0> := map[v52[29] := v57[f4]];
				v58 := v58[true := v56];
				v52[29] := v52[29];
				r1 := fm3(!f6, globalState) + f4;
				var v59: multiset<C0> := multiset{v56};
				var v60: map<int, int> := map[if (v52[29]) then p0 else |(seq(0x10f, i2  => (v45)))| := f4];
				r1, v59, v52[29], v60, globalState.f2 := --f4, multiset{v56, v56, v56, v56, v56}, !v52[29], v60 + v60, fm12(f5 >= p0, f6, f6, globalState);
				var v61: multiset<int> := multiset{-0x166};
				v52[29] := v52[29] || (v61 > v61);
			} else {
				globalState.f2 := v48[p0 := 'n'] == fm6(f6, p0, f6, globalState);
				var v62: map<bool, bool> := map[f6 := v52[29]];
				v62 := if (f6) then v62 else v62;
				var v63: array<multiset<set<int>>> := new multiset<set<int>>[26];
				var v64: set<int> := {f4, f4, f4};
				v63[305] := multiset{v64};
				var v65: multiset<set<int>> := multiset{{|v48|, 0x224, DC9(f4, v52[29], p0, f6).cf8}};
				v63[305] := v65;
				v64 := v64 * v64;
				r1 := (f5 - f5) + fm3(f6, globalState);
			}
			
			if (f6) {
				f5 := f4;
				f5 := -f4;
				v52[29] := f6;
				v45 := v45;
				globalState.f2 := v52[29];
			} else {
				v52[29] := f6;
				v52[29] := f6;
				f5 := -f5;
				var v66: map<D2, bool> := map[DC6([f6, v52[29]]) := f6];
				v44[26] := -|v66| * f5;
				v44[26] := p0 / f4;
				var v67: T0 := new C1(f4, f4);
				var v68: seq<T0> := [v67, v67];
				var v69: map<int, int> := map[f4 := |v68|];
				var v70: set<int> := {|v69|};
				var v72: seq<set<int>> := [v70, set v71 : int | (0x186 <= v71) && (v71 < 0x228) :: (v71 - f5)];
				globalState.f2 := !(v72[|"jucpayd"|] !! v70);
			}
			
			f5 := -p0;
			r1 := f4;
		} else {
			r1 := p0 + (f4 - 0xd1);
			var v73: multiset<char> := multiset{v45};
			r1 := |(v73 - v73)| - (615 / |multiset{f6, f6}|);
			var v74 := DC9(f4, f6, f4, f6);
			var v75: seq<bool> := [true, v74.cf9];
			r0 := (v75 + v75) <= v75;
			var v76: array<char> := new char[10] [if (f6) then v45 else v45, if (f6) then v45 else v45, if (f6) then v48[f4] else v45, v45, v45, v45, v45, v45, v45, v45];
			v76[799] := v45;
			v76[799] := v45;
			f5 := fm3(f6, globalState);
		}
		
		r0 := f6;
		r0 := f6;
		v44[243] := f5 * p0;
		var v77: array<set<int>> := new set<int>[2](i3 => {|(seq(0x14, i4  => (v45)))|, f4, -0x295, f4});
		v44[243], v77 := v50[f4], v77;
		r0 := f6;
		r1 := f5 + v44[243];
	}
}

class C3 {
	const f3 : bool
	constructor (f3 : bool) {
		this.f3 := f3;
	}
	
	method m0(p0: array<string>, p1: int, p2: bool, p3: map<array<int>, bool>, globalState: GlobalState) returns (r0: seq<int>, r1: map<map<int, char>, bool>) {
		var v0 := new C2(fm12(p2, p2, f3, globalState), p1, p1);
		var v1 := "iqwgvcpna";
		if (v1 < v1) {
			var v2: seq<bool> := [!v0.f6, fm12(v0.f6, p2, v0.f6, globalState), true];
			globalState.f2 := v2[p1];
			if (if (p2) then v0.f6 else p2) {
				var v3: map<char, bool> := map['e' := f3];
				var v4: multiset<map<char, bool>> := multiset{v3};
				var v5: seq<int> := [-|v4|];
				var v6: multiset<bool> := multiset{v0.f6, v0.f6, v5 < v5, p1 <= |v1|, !v0.f6};
				var v7: seq<seq<bool>> := [v2];
				v6 := v6 + v6[p2 := |v7|];
				var v8 := 'r';
				v8 := if (p2) then v8 else 'd';
				var v9: map<char, string> := map[v8 := v1];
				v9 := v9[v8 := v1];
				var v10: set<bool> := {p2};
				globalState.f2 := (v10 - v10) !! v10;
				var v11 := -0x6b;
				v11 := p1;
			} else {
				var v12: array<char> := new char[2];
				var v13 := 'p';
				v12[880] := v13;
				v12[880] := v13;
				var v14: map<int, int> := map[p1 := |v1|];
				var v15: map<int, bool> := map[p1 := false];
				var v16: set<bool> := {v0.f6};
				var v17 := DC12(v16);
				v14 := map[|fm13(v0.f6, |v15|, p1, globalState)| := |v17.cf14|] + v14;
				var v18 := 0x336;
				var v19: map<bool, int> := map[v0.f6 := v18];
				var v20: seq<int> := [v18];
				var v21: T0 := new C2(v0.f6, |v20|, v18);
				var v22: map<int, T0> := map[v18 := v21];
				v18 := if (v0.f6 in v19) then v19[v0.f6] else |(fm6(false, |v22|, f3, globalState))[0x75 := v12[880]]| * p1;
				var v23: array<bool> := new bool[26] [p2, v0.f6, v0.f6, v0.f6, v0.f6, v0.f6, v0.f6, v0.f6, v0.f6, p2, p2, v0.f6, p2, v0.f6, p2, v0.f6, !p2, v0.f6, p2, f3, f3, v0.f6, v0.f6, v0.f6, v0.f6, v0.f6];
				var v24: seq<array<bool>> := [v23, v23];
				var v25: seq<seq<array<bool>>> := [v24];
				var v26: seq<seq<array<bool>>> := [v24, (v25[p1])[v21.f4 := v23]];
				var v27: map<char, int> := map[v13 := |v26[v21.f5]|];
				var v28: array<bool> := new bool[3] [p2, v27 in (seq(0x266, i0  => (map[v13 := |v20|]))), !(v21.f4 != p1)];
				v12[880], v23, v18 := if (if (v0.f6) then v0.f6 else p2) then v1[p1] else v13, DC15(v28).cf19, -fm3(p2, globalState);
				v22 := v22[p1 + -p1 := v21];
			}
			
			var v29 := 0x1be;
			var v30: multiset<bool> := multiset{!f3};
			v29 := |v30|;
			var v31: seq<int> := [if (v0.f6) then p1 else p1];
			v29 := -v31[p1];
			globalState.f2 := !(p1 < (if (v0.f6) then p1 else -|v1|));
		} else {
			var v32: T0 := new C1(p1, p1);
			var v33: map<bool, T0> := map[f3 := v32];
			var v34: array<bool> := new bool[7];
			var v35: seq<array<bool>> := [v34];
			var v36 := DC14(v33 + v33, p2, v35);
			v36 := v36.(cf17 := f3, cf16 := v33);
			v32.f5 := v32.f5;
			var v37: array<multiset<bool>> := new multiset<bool>[6](i1 => multiset{v0.f6, p2});
			var v38: multiset<bool> := multiset{v0.f6};
			v37[147] := v38 * v38;
			v37[147] := multiset{!v0.f6, f3};
			var v39: seq<bool> := [p2];
			var v40 := DC6(v39);
			match (v40) {
				case DC7() =>
					var v41: array<array<T0>> := new array<T0>[26];
					var v42 := 'e';
					var v43: map<char, string> := map[v42 := fm6(p2, v32.f5, !f3, globalState)];
					var v44: C1 := new C1(|(if (v42 in v43) then v43[v42] else v1)|, |(seq(0x288, i2  => (|[502, v32.f4, |v1|, p1]|)))|);
					var v45: map<C1, int> := map[v44 := v32.f5];
					v34, v41, globalState.f2 := v34, v41, (if (v44 in v45) then v45[v44] else 0x174) != p1;
					var v46: array<D3> := new D3[17];
					var v47: map<array<D3>, int> := map[v46 := p1];
					var v48: map<bool, bool> := map[v0.f6 := p2];
					var v49: map<map<array<D3>, int>, map<bool, bool>> := map[v47 + v47 := map[v0.f6 := p2] + v48];
					v49 := v49[v47 := v48];
					globalState.f2 := v0.f6;
					var v50: array<map<bool, bool>> := new map<bool, bool>[5];
					v50[771] := v48[true := false];
					v50[771] := fm14(v32.f5, globalState) + v48;
				case DC6(cf6) =>
					var v51 := new C2(p2, -v32.f5, 0x285);
					var v52 := DC8(v1);
					var v53: map<int, D5> := map[|v1| + -v32.f5 := DC16(v32.f5, v37[147], [fm12(f3, v0.f6, v51.f6, globalState), fm12(v0.f6, p2, !f3, globalState)], v52)];
					var v54 := DC16(|v1|, multiset(v39), cf6, v52);
					v53 := v53[|fm15(globalState)| := v54];
					var v55 := 'i';
					var v56 := new C0(multiset{v55});
					var v57: array<int> := new int[21];
					var v58: seq<array<int>> := [v57];
					v58 := [v57, v57, v57, v57, v57];
			}
			
			globalState.f2 := f3;
		}
		
		for i3 := p1 to |v1| {
			var v59: array<bool> := new bool[9](i4 => !f3);
			v59[575] := if (p2) then f3 else f3;
			v59[575] := p2 <==> v0.f6;
			var v60 := 'w';
			var v61: C0 := new C0(multiset{'y', v60});
			var v62: multiset<C0> := multiset{v61, v61};
			var v63: C1 := new C1(|v62|, 0x33f);
			v63 := new C1(fm3(DC4(!false).cf4, globalState), i3);
			var v64 := 0x95;
			var v65: set<bool> := {v0.f6};
			var v66: set<int> := {|v65|};
			var v68: set<set<int>> := {v66, set v67 : int | (0x253 <= v67) && (v67 < 1) :: (v67 / -0x3e0)};
			v64 := -|v68| - i3;
			var v69 := new C0(v61.f7);
		}
		var v70: seq<bool> := [true];
		var v71 := DC4(v70[p1]);
		var v72 := DC5(v71);
		v72 := v72;
		var v73 := 'y';
		var v74: map<bool, char> := map[p2 := v73];
		v74 := v74[p2 := v73];
		for i5 := p1 to p1 {
			var v75: map<int, bool> := map[i5 := v0.f6];
			var v76: multiset<int> := multiset{|fm16(v0.f6, fm3(v0.f6, globalState), v0.f6, globalState)|, 0x1f8, i5};
			var v77 := new C1(-609, |v75| - |v76|);
			var v78: map<bool, int> := map[f3 := i5];
			v78 := v78[true := if (v0.f6) then p1 else p1];
			var v79 := 0x23;
			v79 := -v79;
			var v80 := DC2();
			v80 := v80;
		}
		var v81: seq<int> := [|v1|, -p1];
		r0 := v81 + v81;
		var v82: map<int, char> := map[p1 := v73];
		var v83: map<map<int, char>, bool> := map[v82 := v0.f6];
		r1 := v83;
	}
}

method Main() {
	var globalState := new GlobalState(true, false, true);
	var v0 := 'h';
	var v1: multiset<char> := multiset{v0, v0, v0, v0, v0};
	var v2 := new C0(v1);
	var v3 := 295;
	v3 := -0x312;
	var v4: array<bool> := new bool[29];
	var v5 := true;
	v4[831] := v5;
	v4[831] := v3 >= v3;
	var v6: array<D4> := new D4[21];
	var v7: set<bool> := {v4[831], v4[831]};
	v6[885] := DC12(v7);
	var v8 := DC12(v7);
	v6[885] := v8;
	var i0 := 0;
	while (false)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		var v9: array<char> := new char[18](i1 => v0);
		v9[45] := 'h';
		var v10: map<int, int> := map[v3 := v3];
		var v11: seq<int> := [|v10|];
		var v12: seq<bool> := [v5, v5];
		v9[45], v11, v5 := if (v12[|[v3, v3]|]) then 't' else v0, v11 + v11, v4[831];
		var v13 := "mpqme";
		v13 := v13;
		var v14 := DC7();
		v14 := v14;
		var v15: map<int, bool> := map[v11[v3] := v5];
		globalState.f2 := !((if (v5) then v15 else v15) != map[v3 := v5]);
	}
	v7 := v7 + v7;
	for i2 := v3 to |v7| {
		var v16: multiset<multiset<char>> := multiset{fm17(globalState), v1};
		if (!!((if (v2.f7 in v16) then v16[v2.f7] else v3) < i2)) {
			var v17: seq<int> := [i2, i2];
			var v18 := "nqlglyph";
			globalState.f2 := v17[|v18|] <= i2;
			var v19: C3 := new C3(v4[831]);
			var v20 := DC17(v19);
			v19 := v20.cf24;
			globalState.f2 := v5;
			var v21: array<int> := new int[14];
			v21[535] := |v18|;
			v21[535] := fm3(v3 <= |(seq(946, i3  => (v18)))|, globalState);
			var v22: map<char, C3> := map[v0 := v19];
			v22 := v22;
		} else {
			var v23: seq<int> := [0x363];
			var v24: map<int, int> := map[i2 := v3];
			var v25: map<int, array<bool>> := map[v23[|v24|] := v4];
			var v26: seq<array<bool>> := [v4];
			v4 := if (i2 in v25) then v25[i2] else v26[fm3(v4[831], globalState)];
			var v27: map<bool, char> := map[v5 := v0];
			v27 := v27[!v4[831] := v0];
			var v28: array<int> := new int[10];
			v28[436] := v3 + i2;
			var v29 := "oasbjj";
			var v30: map<string, int> := map[v29 := 0x2b4];
			v28[436] := if (v29 in v30) then v30[v29] else i2;
			v0, v3, v3 := v29[v3], i2, -i2 - i2;
			globalState.f2 := v4[831];
		}
		
		var v31 := "tvpces";
		var v32 := DC18(v31, v5, i2);
		v32 := v32;
		var v33 := DC8(v31);
		v31 := v33.cf7 + (fm6(v4[831], |map[v5 := i2]|, v4[831], globalState) + v31);
		var v34: array<D2> := new D2[24](i4 => DC6([v4[831]]));
		var v35: map<array<D2>, bool> := map[v34 := !v5];
		var v36: array<int> := new int[5];
		var v37: map<array<int>, map<array<D2>, bool>> := map[v36 := v35];
		v35 := if (v36 in v37) then v37[v36] else v35;
	}
	v3 := v3;
	v3 := v3;
	var v38: seq<int> := [v3];
	var v39: array<int> := new int[5] [v3, 0x86, v3, -v3, |v38|];
	var v40: seq<array<int>> := [v39];
	v40 := v40;
	var v41 := "oayxq";
	var v42 := DC2();
	v41 := match v42 {
		case DC1(cf1, cf2) => "n" + v41
		case DC2() => v41 + v41
		case DC0(cf0) => v41
	};
	v4[831] := !v4[831];
	if (v5) {
		var v43: map<int, array<int>> := map[v3 := v39];
		v39 := if ((v3 / v3) in v43) then v43[v3 / v3] else v39;
		var v44 := new C3(v4[831]);
		globalState.f2 := v44.f3;
		var v45: seq<bool> := [v4[831], true];
		var v46: array<seq<int>> := new seq<int>[8](i5 => v38);
		v45, v46 := v45, v46;
		var v47 := DC4(true);
		var v48 := DC5(v47);
		var v49 := DC5(v47);
		var v50 := DC5(v48);
		var v51: C1 := new C1(v3, |map[v3 := v5]|);
		var v52: seq<C1> := [v51, v51, v51, v51];
		v50 := fm18(|v52|, globalState);
	} else {
		var v53: array<map<T0, bool>> := new map<T0, bool>[12];
		var v54: T0 := new C2(true, v3, v3);
		var v55: map<T0, bool> := map[v54 := true];
		v53[151] := v55;
		v53[151] := v55;
		v41 := v41;
		var v56: map<bool, bool> := map[v41 == (seq(-0x15a, i6  => (v0))) := v5];
		v4[831] := if (v4[831] in v56) then v56[v4[831]] else v5;
		var v57: seq<bool> := [v4[831], v5, v4[831], v5, v4[831]];
		var v58: map<set<bool>, bool> := map[v7 := v4[831]];
		if (if (v57[v54.f4]) then true else if (fm9(v54.f4, v5, |v41|, globalState) in v58) then v58[fm9(v54.f4, v5, |v41|, globalState)] else v5) {
			var v59: multiset<array<int>> := multiset{v39, v39, v39, v39};
			var v60: seq<multiset<array<int>>> := [v59, v59, v59];
			v0, v59, v5, v4[831], v54.f5 := v0, v60[v54.f4 + v54.f4], v5 ==> v4[831], v5, v54.f5;
			var v61: C3 := new C3(v57[|v57|]);
			v61 := v61;
			v3 := -v3;
			var v62: map<bool, int> := map[v61.f3 := 0x174];
			v54.f5 := if (v5 in v62) then v62[v5] else v54.f4;
			var v63: multiset<seq<int>> := multiset{v38, v38 + v38, seq(641, i7  => (DC1(|map[v54.f4 := v3]|, v54.f5).cf2))};
			v54.f5 := if ((if (v4[831]) then v38 else [v54.f4, v3]) in v63) then v63[if (v4[831]) then v38 else [v54.f4, v3]] else v54.f5;
		} else {
			v4[831] := fm12(false, v4[831], v5, globalState);
			v54.f5 := 0x25d;
			v39[491] := v54.f4;
			v39[491] := v54.f5;
			v39[491] := v3 + v3;
			var v64: array<int> := new int[5] [v3, |v38|, 0x2e9, v54.f4, 728 % v54.f4];
			var v65: array<C1> := new C1[10];
			var v66: set<array<C1>> := {v65};
			v64, v39[491], globalState.f2 := v64, v3, !(v66 >= {v65});
		}
		
		v54.f5 := v54.f5;
	}
	
	match (v42) {
		case DC1(cf1, cf2) =>
			var v67: map<bool, bool> := map[v4[831] := fm12(v4[831], !v5, v5, globalState) in [v5]];
			var v68: C2 := new C2(v5, cf2, cf2);
			var v69: map<int, C2> := map[v3 := v68];
			v67 := v67[!!v4[831] := !(|v38| >= |v69|)];
			cf2 := cf2;
			var v70: C1 := new C1(fm3(v4[831], globalState), v3);
			v70 := v70;
			var v71: multiset<bool> := multiset{v5};
			var v72 := DC18(v41, v5, |v71[v5 := v3]|);
			var v73 := DC19(v72);
			v73 := v73;
		case DC2() =>
			var v74: C1 := new C1(-v3, 933);
			v74 := v74;
			v3 := v3 + v3;
			var v75: array<C0> := new C0[4];
			v75[348] := v2;
			v75[348] := v2;
			var v76: array<C1> := new C1[14];
			v76 := v76;
		case DC0(cf0) =>
			var v77: map<int, int> := map[v3 := v3];
			var v78: map<int, int> := map[|v77| := v3];
			var v79 := DC0(|v78|);
			var v80: seq<D0> := [v79];
			v80 := (v80 + v80)[cf0 + cf0 := v79];
			var v81: T0 := new C1(-fm3(v4[831], globalState), |v7|);
			var v82: seq<T0> := [v81];
			var v83: seq<bool> := [v4[831], v5];
			var v84: set<seq<bool>> := {v83};
			var v85: map<int, T0> := map[0x275 / -v3 := v82[|v84|]];
			v85 := (v85 + v85) + v85;
			globalState.f2 := false;
			var v86, v87, v88 := v81.m1(if (v4[831]) then v39 else v39, 'u', globalState);
	}
	
	match (DC8(v41)) {
		case DC9(cf8, cf9, cf10, cf11) =>
			cf9 := v4[831];
			var v89: multiset<int> := multiset{cf8, -0x3b8};
			var v90: set<int> := {if (cf8 in v89) then v89[cf8] else v3, cf10, -0x3e2, v3};
			var v91, v92 := m5([cf10] + v38, v90, globalState);
			var v94, v95 := m5(v38, set v93 : int | (-0x154 <= v93) && (v93 < 994) :: (v93 - cf8), globalState);
			if (!v91) {
				var v96, v97 := m5(v38, v90, globalState);
				v40, cf8 := [if (cf9) then v39 else v39, v39, v39, v39], v3;
				globalState.f2 := v5 <== cf11;
				var v98 := DC3(cf9);
				var v99: multiset<D1> := multiset{fm19(cf9, !v4[831], globalState), v98, v98, v98, v98.(cf3 := v4[831])};
				var v100: array<multiset<D1>> := new multiset<D1>[6] [v99, v99, v99 * (multiset{v98})[DC3(v94) := cf10], v99, v99, v99];
				v100 := v100;
				var v101 := DC20(v38);
				globalState.f2, v41, v4[831] := v101.cf29 <= v38, fm6(v5, cf10, v4[831], globalState) + "ygn", !(cf10 > v3);
			} else {
				var v102: seq<string> := [v41];
				var v103: multiset<bool> := multiset{v4[831]};
				var v104 := DC18(v41, true, |v103|);
				var v105 := DC18("q", v91, v104.cf27);
				var v106: seq<string> := [v41, v102[v3], v104.cf25 + v41];
				v102 := v106;
				var v107 := DC0(v3);
				cf8 := cf8 + v107.cf0;
				var v108: C3 := new C3(cf11);
				var v109: T0 := new C2(true, |v89|, |v41| % cf10);
				var v110: C2 := new C2(v108.f3, v109.f4 * cf10, 685);
				var v111: map<bool, bool> := map[cf11 := v4[831]];
				globalState.f2, v108, v109, v110 := v111 == (map[v5 := v94] + v111), if (v91) then v108 else v108, v109, v110;
				v111 := fm14(0x21f, globalState);
				cf9 := v108.f3;
			}
			
		case DC10(cf12) =>
			v4[831] := !v4[831];
			var v112 := new C2(v4[831], v3, v3);
			var v113: C3 := new C3(v112.f6);
			var v114: map<bool, C3> := map[v112.f6 := v113];
			var v115: map<int, bool> := map[-0x2e8 := v112.f6];
			var v116: map<C3, map<int, bool>> := map[if (v5 in v114) then v114[v5] else v113 := v115];
			var v118: seq<seq<int>> := [v38];
			var v119: seq<bool> := [v4[831], false, !v5, v112.f6];
			v116 := v116[v113 := map v117 : int | v117 in v118[|v119|] :: (v117 - |v41|) := (v112.f6)];
			v3 := v3 % (|v38| / v3);
		case DC8(cf7) =>
			v39[339] := |(seq(-0x386, i8  => (v3)))|;
			v39[339] := v3 % v3;
			var v120: seq<bool> := [v4[831]];
			globalState.f2 := v120[v3 := v4[831]] <= v120;
			var v121 := DC22(|v120|);
			var v122 := DC0(v121.cf33);
			v3 := v122.cf0;
			v38 := (v38 + [|{v4[831], v5}|]) + v38;
		case DC11(cf13) =>
			var v123: map<bool, bool> := map[v5 := v4[831]];
			var v124: map<map<bool, bool>, int> := map[v123[v5 := v5] := v3];
			var v125: map<map<map<bool, bool>, int>, array<int>> := map[v124 := v39];
			v125 := v125[v124 := v39];
			var v126 := DC23(v2);
			v2 := v126.cf34;
			var v127 := new C1(v3 - v3, v3 * -v3);
			v5 := v5;
	}
	
	for i9 := v38[0x14a] to v3 {
		var v128: set<int> := {v3};
		var v129, v130 := m5(v38, v128, globalState);
		var v131: T0 := new C2(false, v3, v3);
		v131 := new C2(v129, i9, v3);
		var v132: C3 := new C3(fm12(v5, v5, v4[831], globalState));
		v5, v4[831], globalState.f2, v132 := v5, |({v132, v132} * {v132})| < v131.f5, v5, v132;
		v3 := (-0x12e / --v131.f5) % (-i9 % v131.f4);
	}
}