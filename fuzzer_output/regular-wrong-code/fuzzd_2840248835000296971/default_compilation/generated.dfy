datatype D0 = DC1(cf1: array<array<int>>) | DC0(cf0: seq<int>)
datatype D1 = DC3 | DC2(cf2: bool)
datatype D2 = DC5(cf4: int, cf5: int) | DC6 | DC7(cf6: bool, cf7: int, cf8: bool, cf9: map<multiset<bool>, int>) | DC4(cf3: int)
datatype D3 = DC9(cf11: multiset<int>) | DC10(cf12: bool, cf13: bool, cf14: int, cf15: int, cf16: int) | DC8(cf10: map<D0, bool>)
datatype D4 = DC11(cf17: seq<array<int>>)
datatype D5 = DC13(cf19: D0, cf20: int, cf21: int, cf22: int) | DC12(cf18: seq<D2>)
class GlobalState {
	const f0 : int
	var f1 : bool
	const f2 : seq<int>
	var f3 : string
	const f4 : array<multiset<char>>
	constructor (f0 : int, f1 : bool, f2 : seq<int>, f3 : string, f4 : array<multiset<char>>) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
	}
	
}

function fm0(globalState: GlobalState): int {
	|(if (false) then seq(0x3e, i0  => ('f')) else (seq(584, i1  => ('h'))) + "jkumqri")|
}
function fm1(p0: int, p1: bool, p2: bool, p3: bool, globalState: GlobalState): bool {
	!((---0x3ac + 0x38e) == |("rrv" + "ffpcdlj")|)
}
function fm2(p0: bool, p1: int, globalState: GlobalState): D0 {
	DC0(seq(-0x3ce, i0  => (0x231)))
}
function fm3(p0: bool, p1: bool, globalState: GlobalState): set<int> {
	set v0 : int | (224 <= v0) && (v0 < 0x3bd) :: (v0 % -0x3a2)
}
function fm4(p0: bool, globalState: GlobalState): char {
	if (!false || true) then 'a' else 's'
}
function fm5(p0: int, p1: int, globalState: GlobalState): seq<bool> {
	([true, true] + [false]) + ([DC7(true, -0x246, true, map[multiset{!false, true} := 0x172]).cf6, false] + [true, true])
}
function fm6(p0: int, p1: D2, globalState: GlobalState): string {
	seq(767, i0  => ('w'))
}
function fm7(globalState: GlobalState): seq<int> {
	match DC7(false, 0x2d6, true, DC7(false, |map[153 := set v1 : int | v1 in (map v0 : int | (101 <= v0) && (v0 < 0x25b) :: (v0 % |{|[false, false]|}|) := (DC2(false).cf2)) :: (v1 / |[|[|[-790]|, 0x1]|]|)]|, false, map[multiset{!false} := |"gcxukgk"|]).cf9) {
		case DC5(cf4, cf5) => [0x35b]
		case DC6() => (seq(989, i0  => (|(set v2 : multiset<bool> | v2 in [multiset{true}] :: (v2))|))) + [71]
		case DC7(cf6, cf7, cf8, cf9) => [cf7]
		case DC4(cf3) => [|map[true := !!false]|, cf3, cf3, cf3]
	}
}
function fm8(p0: bool, p1: bool, p2: seq<bool>, p3: bool, globalState: GlobalState): set<bool> {
	{true, true, !true} - ({false, true} - {!false})
}
function fm9(globalState: GlobalState): multiset<bool> {
	multiset{false} + (multiset{false, false} * multiset([true, true]))
}
method m0(globalState: GlobalState) returns (r0: array<int>, r1: int) {
	r1 := fm0(globalState);
	var v0 := 0x2d7;
	var v1: set<int> := {v0};
	var v2 := false;
	var v3: multiset<bool> := multiset{v2};
	var v4: seq<multiset<bool>> := [v3, v3];
	var v5 := 'k';
	var v6: map<multiset<bool>, int> := map[v4[|[v5]|] := |v3|];
	var v7 := DC7(v1 > v1, v0, v2, v6);
	match (v7) {
		case DC5(cf4, cf5) =>
			var v8: array<set<string>> := new set<string>[17](i0 => {"sldfemqf", "e"});
			var v9: set<string> := {"adcvwjlwm"};
			v8[144] := v9 + v9;
			var v10 := "rbff";
			var v11: seq<string> := ["oktgtfhvg", v10];
			v8[144] := (v9 + (set v12 : string | v12 in v11 :: (v12))) * {seq(0x24d, i1  => (v5)), seq(792, i2  => (v5))};
			var v13: map<char, bool> := map[if (v2) then v5 else v5 := !v2];
			globalState.f1 := if (fm4(v2, globalState) in v13) then v13[fm4(v2, globalState)] else v2;
			var v14: map<int, bool> := map[cf5 := true];
			var v15: map<int, bool> := map[-cf5 := if (|fm5(cf5, cf4, globalState)| in v14) then v14[|fm5(cf5, cf4, globalState)|] else v2];
			var v16: seq<bool> := [true, v2, v2];
			var v17: seq<seq<bool>> := [v16, v16, [v2, true, v2, v2, v2]];
			var v18: array<seq<bool>> := new seq<bool>[21] [[if (cf5 in v15) then v15[cf5] else !true, v2], v16, v16 + v16, [fm1(v0, v2, v2, true, globalState)], v16, [false, !v2], v16 + v16, v16, v16, v16, v17[v0], v16 + v16, v16 + v16, v16, v16 + v16, v16, v16, v16, v16, v16, v16];
			v18[703] := v16;
			v18[703] := v16;
			globalState.f1, v2, r1 := if (true) then v2 else v2 || v2, !v2, |"wmqcbpj"|;
		case DC6() =>
			var v19: set<bool> := {v2, v2, v2};
			var v20: C0 := new C0(|multiset{v2, v2, v2}|);
			var v21: set<C0> := {v20, v20, v20, v20, v20};
			var v22: array<bool> := new bool[11] [v2, !!false <== v2, v2, v2, v2, v19 > {v2, true, v2, false, v2}, v2, v0 != v0, v2, v2, v21 > v21];
			v22[912] := v2;
			var v23 := "h";
			var v24 := DC4(v20.f5);
			v22[912] := (v23 + v23) < (fm6(v0, v24.(cf3 := v20.f5), globalState))[|{v2, v2, v2, v2, v2}| := v5];
			var v25: array<seq<D4>> := new seq<D4>[1];
			var v26: array<int> := new int[5];
			var v27: seq<array<int>> := [v26];
			var v28 := DC11(v27);
			v25[226] := [v28];
			var v29 := DC10(v22[912], v22[912], v0, v0, v20.f5);
			var v30: seq<D4> := [DC11(v27), if (v29.cf13) then v28 else v28, v28, v28];
			v25[226] := v30;
			var v31: multiset<C0> := multiset{v20};
			var v32: seq<int> := [|v31|, fm0(globalState)];
			var v33: multiset<seq<int>> := multiset{v32};
			v33 := v33 * v33;
			if (DC10(v2, v22[912], v32[v0], v20.f5, v0).cf12) {
				var v34: map<char, int> := map[v5 := v20.f5];
				var v35: multiset<int> := multiset{-0x3ab, |v34|};
				globalState.f1 := (multiset{fm1(if (v0 in v35) then v35[v0] else v0, fm1(v20.f5, v2, v22[912], v2, globalState), v22[912], true, globalState)} > v3) <==> !v2;
				var v36: array<char> := new char[10];
				var v37: multiset<array<char>> := multiset{v36};
				var v38: map<int, bool> := map[|v37| / -0xe0 := v22[912]];
				v38 := v38[v20.f5 := v22[912]];
				var v39 := new C0(v20.f5 % 0x34c);
				v39 := new C0(v39.f5 * v0);
				v20.f5 := fm0(globalState);
			} else {
				v32 := v32;
				globalState.f1 := v2;
				var v40: array<char> := new char[27] [fm4(false, globalState), v5, v5, v5, v5, 'a', v5, v5, v5, v5, v5, v5, v5, v5, v5, v5, v5, 'n', v5, v5, v5, if (fm1(v20.f5, !v22[912], v2, v22[912], globalState)) then v5 else 'c', v5, 'n', v5, v5, v5];
				v40[672] := v5;
				var v41: map<string, seq<int>> := map[v23 := v32 + v32];
				v20.f5, v40[672], globalState.f1, v32 := v20.f5, 'v', (v20.f5 * v20.f5) == v0, if ((v23 + v23) in v41) then v41[v23 + v23] else seq(0x2cb, i3  => (v20.f5));
				var v42 := new C0(v20.f5);
				v22 := new bool[16];
			}
			
		case DC7(cf6, cf7, cf8, cf9) =>
			var v43 := DC4(v0);
			var v44 := "xautmy";
			var v45 := DC12(([v43, v43, v43, v43, v43])[|v44| := DC4(v0)]);
			var v46: seq<D2> := [v43];
			var v47: seq<int> := [-cf7];
			var v48: array<bool> := new bool[28] [if (cf8) then cf8 else v2, cf6, v2, true, cf6, !cf8, v2, v2, cf7 == |{cf7}|, cf6, multiset((v45.(cf18 := [v43])).cf18) > multiset(v46), if (!v2) then cf8 else v2, cf8, v2, !cf6, |fm7(globalState)| >= |v44|, cf8, -v0 != -v0, v2, v2, fm1(|v47|, true, false, cf8, globalState), v3 > v3, cf6, cf8, !v7.cf8, cf8, v44 != v44, v2];
			v48[203] := v2;
			var v49: map<int, int> := map[v0 := fm0(globalState)];
			var v50: array<int> := new int[3];
			r0, r0, v48[203], v49, v48 := if (true) then v50 else v50, v50, if (cf8) then v2 else cf6, map[cf7 := v47[-v0]], v48;
			var v51: C0 := new C0(v0);
			var v52: array<C0> := new C0[6] [v51, v51, v51, if (fm1(v0, false, false, cf6, globalState)) then v51 else v51, v51, v51];
			v52[831] := v51;
			var v53: map<int, string> := map[cf7 := if (v48[203]) then seq(0xd8, i4  => (v5)) else v44];
			v52[831], v53 := v51, v53 + v53;
			if (v2) {
				v49 := v49[-v51.f5 := fm0(globalState)];
				var v54: multiset<array<int>> := multiset{v50};
				v54 := v54;
				var v55: set<bool> := {!!false};
				var v56: array<array<int>> := new array<int>[14] [v50, v50, v50, v50, v50, v50, v50, v50, v50, v50, v50, v50, v50, v50];
				v56[765] := v50;
				var v57: array<seq<int>> := new seq<int>[21](i5 => v47);
				v57[719] := [|v44|, cf7];
				var v58: seq<bool> := [v48[203]];
				v55, v56[765], v57[719] := fm8(cf6, cf8, v58 + v58, v55 == v55, globalState), v50, fm7(globalState);
				var v59: multiset<multiset<bool>> := multiset{v3, v3};
				v59 := multiset((seq(0x1a6, i6  => (v3))) + [v3, multiset(v58), fm9(globalState), v3]) * v59;
				v50[165] := v51.f5;
				v50[165] := cf7;
			} else {
				v48 := v48;
				v51.f5, globalState.f1 := v0 / -cf7, cf6;
				cf8 := v2;
				v5 := v5;
				var v60: array<string> := new string[29];
				v60[452] := v44;
				v60, v60[452] := v60, (v44 + (seq(621, i7  => (v44[v51.f5])))) + v44;
			}
			
			var v61 := DC11([v50]);
			var v62: multiset<int> := multiset{-259};
			var v63: seq<array<int>> := [v50];
			v61 := if (v62 > v62) then DC11(v63) else v61.(cf17 := v63);
		case DC4(cf3) =>
			var v64: array<bool> := new bool[22];
			v64[910] := v2;
			var v65: map<int, bool> := map[v0 % v0 := true];
			var v66 := DC4(v0);
			var v67: seq<bool> := [v2];
			v64[910], globalState.f1, v65, v66, cf3 := v2, fm1(cf3, if (!fm1(cf3, v2, true, v2, globalState)) then v2 else v67[cf3], cf3 == v0, !v2, globalState), v65[v0 := v2], DC4(v0), -(cf3 % cf3);
			globalState.f1 := v64[910];
			v5 := v5;
			var v68: map<char, int> := map[v5 := 0x17e];
			var v69: C0 := new C0(v0);
			var v70: map<bool, bool> := map[v64[910] := v64[910]];
			var v71: multiset<char> := multiset{v5};
			var v72: array<int> := new int[13] [cf3, |v70|, |v71|, v69.f5, 0x2a7, -cf3, v69.f5, v0, 0x335, v69.f5, v0, v0, v69.f5];
			var v73: seq<array<int>> := [v72, v72];
			var v74 := DC11(v73);
			var v75: map<int, D4> := map[|[v69]| := v74];
			v68 := v68[v5 := |v75|];
	}
	
	var v76 := new C0(v0 * v0);
	var v77: multiset<int> := multiset{v76.f5 * v76.f5, v0};
	v77 := v77 * v77;
	v7 := v7.(cf7 := v0, cf9 := v6);
	var v78 := DC10(v2, v2, v0, v76.f5, v0);
	var v79: set<bool> := {v2};
	var v80: seq<bool> := [fm1(v76.f5, v2, v2, v2, globalState), v2, false];
	var v81 := DC9(v77);
	var v82: multiset<D3> := multiset{v81, v81};
	var v83: set<set<int>> := {v1};
	var v84: array<bool> := new bool[20] [v78.cf12, v2, v2, if (!!v2) then v2 else v2, v76.f5 >= -|v79|, v80[v76.f5], v2, v2, v2, v2, v2, v2, v2, v2, true, v2, v2, v2, v82 != multiset{v81}, v83 >= v83];
	v84[382] := v2;
	v84[382] := v78.cf12;
	r0 := new int[5];
	var v85 := "matixu";
	var v86: map<string, seq<char>> := map[v85 := [v5]];
	var v87: seq<int> := [|((if (v85 in v86) then v86[v85] else v85) + (seq(0x32e, i8  => (v5))))|, v76.f5, v0 * v76.f5];
	r1 := v87[v0];
}
class C0 {
	var f5 : int
	constructor (f5 : int) {
		this.f5 := f5;
	}
	
}

method Main() {
	var v0 := -0x230;
	var v1: seq<int> := [v0, v0];
	var v2: array<multiset<char>> := new multiset<char>[19](i0 => multiset{'r', 's'});
	var globalState := new GlobalState(0x317, false, v1, "kpno", v2);
	var v3 := true;
	var i1 := 0;
	while (v3)
		decreases 100 - i1
	{
		if (i1 >= 100) {
			break;
		}
		
		i1 := i1 + 1;
		var v4: array<bool> := new bool[19];
		v4[258] := false;
		v4[258] := v3 ==> v3;
		if (!v4[258]) {
			var v5, v6 := m0(globalState);
			globalState.f1 := v3 ==> v3;
			v1 := DC0(v1).cf0[fm0(globalState) := v0];
			globalState.f1 := v4[258];
			v0 := fm0(globalState) % (v6 / v0);
		} else {
			var v7, v8 := m0(globalState);
			v3 := v4[258];
			var v9 := DC2(v4[258]);
			v4[258] := v9.cf2;
			var v10 := new C0(v0);
			v0, v10.f5 := v8, 0x5a % v0;
		}
		
		globalState.f1 := true;
		var v11: multiset<bool> := multiset{v3, v4[258]};
		var v12 := DC0(if (v3) then v1 else [|v11|]);
		v0, v12 := if (v3) then 0x264 else v0, v12;
	}
	var v13: array<int> := new int[5];
	forall i2 | 0 <= i2 < v13.Length {
		v13[i2] := i2 % DC4(v0).cf3;
	}
	forall i3 | 0 <= i3 < v13.Length {
		v13[i3] := i3 * v0;
	}
	globalState.f3 := seq(-0x2c4, i4  => ('f'));
	var v14 := new C0(0xbe);
	var v15 := "xqlvtuvhq";
	var i5 := 0;
	while (if (v0 >= v0) then v3 else (seq(0x29f, i6  => ('g'))) != v15)
		decreases 100 - i5
	{
		if (i5 >= 100) {
			break;
		}
		
		i5 := i5 + 1;
		v14.f5 := |v15|;
		var v16: map<bool, int> := map[v3 := v14.f5];
		var v17: map<int, map<bool, int>> := map[v0 := v16];
		var v18: map<map<int, map<bool, int>>, int> := map[v17 := v14.f5];
		v18 := v18[v17 := v14.f5];
		v1 := v1;
		v0 := v14.f5;
	}
	var v19 := 'q';
	v19 := v15[0x2fc];
	var v20: array<map<int, int>> := new map<int, int>[17](i7 => map[v0 := v14.f5]);
	var v21: map<bool, int> := map[v3 := v0];
	var v22: map<int, int> := map[v14.f5 := |v21|];
	v20[590] := v22;
	v20[590] := (v22 + map[-v14.f5 := v14.f5])[v0 := v14.f5];
	for i8 := 0xe8 to v0 {
		var v23 := new C0(v0);
		var v24, v25 := m0(globalState);
		v24[212] := i8;
		var v26: array<bool> := new bool[13] [!v3, false, !v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3];
		var v27: map<array<bool>, bool> := map[v26 := fm1(v25, false, v3, v3, globalState)];
		v24[212], v27, v25 := -fm0(globalState), v27, -v14.f5;
		match (fm2(v3, 321, globalState)) {
			case DC1(cf1) =>
				v24[212] := v14.f5 - (v14.f5 * fm0(globalState));
				var v28 := new C0(v0);
				var v29 := DC5(v14.f5, i8);
				var v30: map<string, D2> := map[v15 := v29.(cf5 := v0)];
				v30 := v30;
				var v31: seq<array<bool>> := [v26];
				var v32: seq<map<bool, seq<bool>>> := [map[!v3 := [v3]]];
				var v33: array<array<bool>> := new array<bool>[25] [v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v31[|v32[v23.f5]|], v26, v26, v26, v26, v26, v26, v26, v26, v26, v26, v26];
				var v34: map<C0, array<bool>> := map[v14 := v26];
				v33[372] := if (v23 in v34) then v34[v23] else v26;
				var v35: map<int, bool> := map[|v20[590]| := v3];
				var v36: multiset<map<int, bool>> := multiset{v35};
				var v37: multiset<int> := multiset{|v36|, i8, -v24[212], v28.f5, v25};
				var v38: seq<C0> := [v28];
				globalState.f1, v33[372] := v37[|v38| := |v15|] !! v37, v26;
			case DC0(cf0) =>
				var v39 := new C0(v24[212]);
				var v40: multiset<bool> := multiset{v14.f5 >= v23.f5};
				var v41: seq<bool> := [v3, v3];
				v40 := multiset{if (v3) then v3 else false, v41[fm0(globalState)], v40 > v40};
				var v42: array<array<int>> := new array<int>[8];
				var v43 := DC1(v42);
				var v44: map<D0, bool> := map[v43 := false];
				var v45 := DC8(v44);
				var v46: seq<map<D0, bool>> := [v44 + v45.cf10, v44, v44, v44[DC1(v42) := v3]];
				v25 := |v46|;
				v39.f5 := --v25;
		}
		
	}
	if (v3) {
		var v47: array<bool> := new bool[25](i9 => false);
		v47[360] := v3;
		v47[360] := !v3;
		v14.f5 := v0;
		v13[587] := if (-v14.f5 in v20[590]) then v20[590][-v14.f5] else v0;
		var v48: seq<array<int>> := [v13, v13, v13, v13];
		var v50: set<int> := {|v1|, v0};
		var v51: map<C0, array<int>> := map[v14 := v13];
		v47[360], v13[587], v48, v14 := (set v49 : int | (803 <= v49) && (v49 < 238) :: (v49 / v14.f5)) >= (v50 + fm3(v47[360], !v3, globalState)), fm0(globalState), (v48 + DC11([if (v14 in v51) then v51[v14] else v13]).cf17) + [v13], v14;
		globalState.f3 := v15;
		var v52, v53 := m0(globalState);
	} else {
		v0 := -v14.f5;
		var v54: multiset<int> := multiset{v0};
		if (fm1(v14.f5, |(seq(0x3a2, i10  => (v14.f5)))| != v0, |v54| < v0, v3, globalState)) {
			var v55: array<bool> := new bool[12];
			v55[358] := !(v14.f5 >= -|v15|);
			v55[358] := v3;
			globalState.f1 := v3;
			var v56: seq<bool> := [v55[358], v55[358]];
			var v57: seq<seq<bool>> := [v56, v56];
			var v58: set<int> := {v0, v0 % v14.f5, |v15[|multiset{fm1(|v57[v14.f5]|, v55[358], true, v3, globalState), v55[358]}| := fm4(v55[358], globalState)]|, 321};
			v58 := (v58 + (set v59 : int | v59 in v1 :: (v59 % -0x1b5))) * {v0};
			v19 := if (true) then v19 else v19;
			var v60: array<char> := new char[16](i11 => v19);
			var v61: map<int, array<char>> := map[v0 := v60];
			v60 := if (v0 in v61) then v61[v0] else if (v55[358]) then v60 else v60;
		} else {
			var v62 := DC5(v14.f5, v14.f5);
			v62 := v62;
			var v63: multiset<bool> := multiset{v3};
			var v64: map<C0, C0> := map[v14 := v14];
			var v65: seq<bool> := [v3];
			v13 := new int[8] [v14.f5, v14.f5, -|v63| / |v15|, v14.f5 / v14.f5, |v64|, fm0(globalState), |v65|, v14.f5];
			v0 := v14.f5 - v14.f5;
			v14.f5 := v0;
			var v66, v67 := m0(globalState);
		}
		
		globalState.f1 := v3;
		var v68 := DC0(v1);
		v68 := v68.(cf0 := v1);
		var v69: set<bool> := {v3};
		var v70: set<set<bool>> := {v69, v69};
		v70 := v70 + v70;
	}
	
	var v71: array<bool> := new bool[1](i12 => v3);
	v71 := v71;
	var v72: map<bool, array<bool>> := map[v3 := v71];
	v72 := v72[false := v71];
	v14.f5 := v14.f5 % (v1[v14.f5] % fm0(globalState));
	var v73, v74 := m0(globalState);
	globalState.f1 := v3;
	for i13 := v14.f5 to v14.f5 {
		var v75: array<map<array<D4>, int>> := new map<array<D4>, int>[6];
		var v76: array<D4> := new D4[17];
		var v77: map<array<D4>, int> := map[v76 := v14.f5];
		v75[843] := if (v3) then v77 else v77;
		v75[843] := v77;
		v3 := v1[v0] < v14.f5;
		v0 := |v15| + v14.f5;
		v14.f5 := v0 * i13;
	}
}