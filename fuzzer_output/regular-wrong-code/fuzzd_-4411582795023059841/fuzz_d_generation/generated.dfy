datatype D0 = DC0(cf0: int)
datatype D1 = DC2(cf2: int, cf3: bool) | DC3(cf4: seq<int>, cf5: int) | DC4 | DC1(cf1: bool)
datatype D2 = DC6(cf7: T0) | DC5(cf6: multiset<bool>) | DC7(cf8: D2)
datatype D3 = DC9(cf10: char) | DC8(cf9: array<string>)
datatype D4 = DC11(cf12: D1, cf13: multiset<int>) | DC10(cf11: string)
datatype D5 = DC13(cf15: bool) | DC12(cf14: map<bool, bool>) | DC14(cf16: D5)
class GlobalState {
	var f0 : bool
	const f1 : bool
	const f2 : bool
	constructor (f0 : bool, f1 : bool, f2 : bool) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
	}
	
}

function fm0(globalState: GlobalState): int {
	-427 * 806
}
function fm1(p0: seq<multiset<int>>, p1: bool, globalState: GlobalState): bool {
	(multiset{false} * multiset([true])) !! multiset([false, false, !true, true])
}
function fm2(p0: int, globalState: GlobalState): multiset<bool> {
	multiset{0x16e < |map[false := -0x28f]|, 0xfc != 0x38b, !((set v1 : int | v1 in (map v0 : int | v0 in (seq(0x1a7, i0  => (DC0(-0x385).cf0))) :: (v0 % 0x3b5) := (true)) :: (v1 - |DC3([0x39, -|map[|"i"| := true]|], -|map[!true := true]|).cf4|)) >= {0x1a8}), 'y' !in ['k'], |["wcdmhndul", "cf", "vdotcjy"]| in map[0x193 := 0x2f2]}
}
function fm5(p0: int, globalState: GlobalState): D1 {
	DC1(!true)
}
function fm6(globalState: GlobalState): D1 {
	DC3([-626] + [-0x372, |{false}|, |map[true := map[|"alqfd"| := !!!true]]|, |"rdcr"|, ---427], |"vc"| - -|"ikfvgwhey"|)
}
function fm7(globalState: GlobalState): D1 {
	DC2(|"isjymm"| + 0x2af, multiset([|map[true := true]|]) !! multiset{0x17d, 0x384, |{!false, false}|})
}
function fm8(p0: seq<int>, globalState: GlobalState): string {
	seq(0x38e, i0  => ('b'))
}
function fm9(p0: bool, p1: bool, p2: bool, globalState: GlobalState): map<bool, bool> {
	map[!!true := false] + DC12(map[false := false]).cf14
}
function fm10(globalState: GlobalState): set<int> {
	({679, 0x2f4, 0x9} - {|"apyq"|, |{|['t']|, 0x110}|, |(map v0 : seq<char> | v0 in [['d', 'k']] :: (v0) := (true))|}) * (set v1 : int | v1 in map[-|[|"st"|]| := true] :: (v1 % -|{true}|))
}
function fm11(p0: seq<int>, p1: int, p2: int, p3: bool, globalState: GlobalState): char {
	't'
}
function fm12(p0: int, p1: int, globalState: GlobalState): seq<int> {
	([0x16c, |{true, !true}|, -|multiset{-0x323}|, -364] + [290, -0x1ef]) + ([|{'f'}|] + [|[-0x258, |map[false := [0x3a1]]|]|, 0x268])
}
method m0(p0: multiset<bool>, p1: int, globalState: GlobalState) returns (r0: bool, r1: int, r2: int) {
	for i0 := p1 to -648 {
		var v0: seq<int> := [p1, fm0(globalState), i0];
		r2 := v0[i0];
		if (i0 != p1) {
			var v1 := 'h';
			v1 := v1;
			var v2: array<char> := new char[10];
			var v3: array<array<char>> := new array<char>[13] [v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2];
			var v4 := false;
			var v5: map<array<array<char>>, bool> := map[v3 := v4];
			v5 := v5[v3 := v4];
			r0 := v4;
			var v6: array<int> := new int[17](i1 => i1 - -|map[v4 := v4]|);
			v6[462] := --266;
			var v7: seq<bool> := [v4];
			var v8: multiset<int> := multiset{-p1};
			var v9: map<array<int>, multiset<int>> := map[v6 := v8];
			var v10: seq<bool> := [v4, v4 <==> v4, multiset{|v7|} !! (if (v6 in v9) then v9[v6] else v8[i0 := i0]), v4];
			v6[462] := |v10|;
			var v11: map<bool, multiset<int>> := map[v4 := v8];
			v11 := v11[v4 := v8];
		} else {
			var v12: multiset<int> := multiset{p1, i0, i0};
			var v13: map<int, multiset<int>> := map[p1 := v12];
			var v14 := 'f';
			var v15: multiset<char> := multiset{v14, 'a', 'x'};
			v13 := v13[p1 := multiset{p1, if ('b' in v15) then v15['b'] else i0}];
			var v16: array<int> := new int[21];
			var v17: multiset<array<int>> := multiset{v16};
			v17 := v17 + multiset{v16, v16};
			var v18: C0 := new C0('t');
			v18 := v18;
			var v19: array<bool> := new bool[21];
			var v20 := false;
			var v21: map<C0, bool> := map[v18 := !v20];
			v19[850] := v18 !in v21;
			v19[850] := v20;
			var v22: array<set<bool>> := new set<bool>[19](i2 => {v20} + {v19[850], v20, v19[850]});
			v22 := v22;
		}
		
		var v23 := false;
		var v24: map<bool, int> := map[v23 := p1];
		var v25: seq<bool> := [v23];
		var v26: map<int, int> := map[i0 := -755];
		var v27: seq<map<int, int>> := [v26, v26];
		var v28: map<int, seq<int>> := map[|v27| := v0];
		var v29: array<int> := new int[27] [p1, i0, i0, p1, i0, if (v23 in v24) then v24[v23] else i0, |{p1}| + i0, p1 - 0xce, p1, p1 - p1, |(v25 + v25)|, -159, -(p1 / p1), i0, |v28[p1 := v0[p1 := i0]]|, 750 - i0, |[p1, p1, --p1]|, i0, -|fm9(v23, v23, v23, globalState)|, p1, p1, i0 - i0, -i0, i0 - -539, i0, i0, p1];
		v29[937] := i0;
		var v30: map<seq<int>, bool> := map[v0 := v23];
		v29[937], v30 := i0, v30;
		var v31 := 't';
		var v32: T0 := new C0(v31);
		var v33: multiset<T0> := multiset{v32};
		var v34: multiset<int> := multiset{|v33|, |v0|};
		var v35: map<bool, bool> := map[v23 := !v25[|v34|]];
		v35 := v35[v23 := v23];
	}
	var v36 := 'w';
	v36 := v36;
	var v37 := true;
	var v38 := "ejbib";
	var v39: multiset<int> := multiset{p1, |v38|};
	var v40: seq<bool> := [v37];
	var v41: map<bool, bool> := map[true := v40[p1]];
	var v42: seq<int> := [p1, |{v37, v37, fm1([v39], v37, globalState)}|, |p0|, p1, |v41|];
	var v43: array<int> := new int[13] [p1, p1, p1 * p1, p1, |fm10(globalState)|, p1, -0x33f, p1 + p1, p1, v42[p1], p1 % -|v41|, p1, p1];
	v43[784] := p1 * p1;
	v43[784] := p1;
	r2 := v43[784];
	var v44: array<bool> := new bool[27](i3 => v37);
	v44[561] := v37;
	globalState.f0, v40, r2, v43[784], v44[561] := fm1([v39], v37, globalState), v40, p1, p1, v37;
	var v45 := DC4();
	if (match v45 {
		case DC2(cf2, cf3) => cf3
		case DC3(cf4, cf5) => {v43[784]} !! {|v39|, |v39|, |"sp"|, |cf4|, v43[784]}
		case DC4() => if (true) then DC1(v37).cf1 else v37
		case DC1(cf1) => v37
	}) {
		v44 := v44;
		v36 := fm11(v42 + v42, p1, 0x2fe, if (v44[561]) then v37 else v37, globalState);
		var v46 := new C0(v36);
		v44[561] := if (v46.fm3(globalState)) then v37 else v37;
		r1 := p1 - 0x1b8;
	} else {
		v44[561] := !v44[561];
		var v47 := new C0(v36);
		var v48: seq<C0> := [v47];
		var v49: map<seq<C0>, bool> := map[v48 + [v47, v47] := v44[561]];
		v49 := v49[v48 := false];
		v37 := p1 <= (v43[784] * p1);
		if (v37) {
			v44[561] := v44[561];
			var v50: array<T0> := new T0[23];
			var v51: T0 := new C0('q');
			v50[611] := v51;
			var v52 := DC6(v51);
			var v53: seq<string> := [v38];
			var v54 := DC10(v53[v43[784]]);
			var v55 := DC10(v54.cf11);
			var v56: map<bool, string> := map[v37 := v54.cf11];
			var v57 := DC1(v44[561]);
			v50[611], globalState.f0, v44[561] := v52.cf7, (if (v37 in v56) then v56[v37] else "pcctiawd") < fm8([p1, v43[784]], globalState), !v57.cf1;
			var v58: map<D1, int> := map[v45 := -v43[784]];
			var v59: array<string> := new string[26] [v38 + fm8(seq(0x4f, i4  => (p1)), globalState), v38, seq(684, i5  => (v36)), (seq(21, i6  => (v51.f3)))[|v58| := v51.f3], (seq(103, i7  => (v51.f3))) + v38, v38, v38, v38, "w", v38[v43[784] := v36], v38, v38, v38[v42[v43[784]] := v51.f3], fm8(v42, globalState), if (v37) then v38 else fm8(fm12(fm0(globalState), p1, globalState), globalState), "jumanw", v38, "f", v38, "kyeeyllc", "k", v38, v38, v38, "hqetjr", (seq(0x50, i8  => (v51.f3)))[|(seq(392, i9  => (v51.f3)))| := v36]];
			v59[878] := "mqdbtkjg";
			v59[878] := v38;
			r1 := v43[784];
			v44[561] := true;
		} else {
			r2 := |((v38 + v38) + v38)|;
			var v60: array<C0> := new C0[16] [if (v44[561]) then v47 else v47, v47, v47, v47, v47, v47, if (!v37) then v47 else v47, v47, v47, v47, v47, v47, v47, v47, v47, v47];
			v60 := v60;
			v38 := seq(128, i10  => (v36));
			var v61: map<bool, int> := map[v37 := v43[784]];
			var v62: map<int, array<int>> := map[p1 := v43];
			r2 := |(v61 + (map[v37 := p1])[v37 := -|map[v44 := |v62|]|])|;
			r2 := -(v43[784] - v43[784]);
		}
		
	}
	
	r0 := v37;
	r1 := 0x1c7 % v43[784];
	r2 := v43[784];
}
trait T0 {
	const f3 : char
	function fm3(globalState: GlobalState): bool 
}

class C0 extends T0 {
	constructor (f3 : char) {
		this.f3 := f3;
	}
	
	function fm3(globalState: GlobalState): bool {
		match DC0(|"tqw"|) {
			case DC0(cf0) => !false
		}
	}
	function fm4(p0: seq<int>, p1: int, globalState: GlobalState): bool {
		DC1(true).cf1
	}
}

method Main() {
	var globalState := new GlobalState(true, true, false);
	var v0 := "jract";
	v0 := v0 + v0;
	var v1 := -0x1b;
	var v2: map<int, int> := map[v1 := -v1];
	v2 := v2[v1 := v1];
	v1 := v1;
	var v3 := true;
	globalState.f0 := v3;
	globalState.f0 := true;
	if (v3) {
		var v4: array<bool> := new bool[1](i0 => true);
		v4[784] := v3;
		v4[784] := v1 >= (fm0(globalState) - v1);
		var v5: array<map<bool, bool>> := new map<bool, bool>[8](i1 => map[v3 := !v3]);
		var v6: multiset<int> := multiset{v1};
		var v7: seq<multiset<int>> := [v6];
		var v8: multiset<bool> := multiset{v4[784]};
		var v9: map<int, bool> := map[|v8| := v4[784]];
		var v10: map<array<map<bool, bool>>, map<bool, multiset<bool>>> := map[v5 := if (fm1(v7, v3, globalState)) then map[if (0x213 in v9) then v9[0x213] else v3 := v8] else map[v4[784] := fm2(v1, globalState)]];
		var v11: map<bool, multiset<bool>> := map[true := v8];
		v10 := v10[v5 := v11];
		var v12: seq<int> := [v1];
		var v13: map<seq<int>, int> := map[v12 := v1];
		var v14: set<int> := {v1, -v1, v1};
		var v15 := DC0(-|v14|);
		var v16: map<int, map<int, bool>> := map[v1 := v9];
		var v17: map<bool, bool> := map[false := v4[784]];
		var v18: array<int> := new int[17] [|v0| * -0x3ad, v1, v1, v1 * v1, |v13| % v1, v1 / v1, -v1, 0x310, v1, v1, v15.cf0, -619, v1, v1 + |v16|, -(if (v1 in v6) then v6[v1] else v1), v1 % |v17|, v1];
		v18[995] := v1 - 0x378;
		v18[995], v17, globalState.f0 := (if (v4[784]) then v1 else |v6|) * v1, if (-v1 <= v1) then v17 else (map[!v4[784] := v4[784]])[v4[784] := v4[784]], 90 <= v15.cf0;
		v18[995], v4[784], v1 := fm0(globalState), v4[784], (v18[995] % v18[995]) % 772;
		var v19 := 'c';
		var v20 := new C0(v19);
	} else {
		var v21: array<bool> := new bool[3];
		v21[869] := v3;
		v21[869], v1 := v3, fm0(globalState);
		var v22: seq<bool> := [v3];
		match (fm5(|v22|, globalState)) {
			case DC2(cf2, cf3) =>
				var v23: set<bool> := {cf3, !false, cf3};
				var v24: set<array<bool>> := {v21, v21};
				var v25: multiset<int> := multiset{|v23|, |v24|};
				cf3 := if (v25 >= multiset([v1, v1])) then v21[869] else cf3;
				cf2 := v1 * cf2;
				var v26: array<int> := new int[2];
				var v27 := DC1(cf3);
				var v28 := 'o';
				var v29: C0 := new C0(v28);
				var v30: map<D1, C0> := map[v27 := v29];
				v26[967] := |v30|;
				v26[967] := v1 / cf2;
				globalState.f0 := v21[869];
			case DC3(cf4, cf5) =>
				globalState.f0 := v3;
				v3 := !v3;
				v1 := 988;
				cf5 := -0x357;
			case DC4() =>
				var v31: multiset<bool> := multiset{v3};
				var v32, v33, v34 := m0(v31, v1, globalState);
				v0 := v0 + v0;
				var v35 := 'v';
				var v36: C0 := new C0(v35);
				var v37: set<int> := {v34, 0x134};
				var v38: multiset<int> := multiset{|v37|};
				var v39: seq<multiset<int>> := [v38];
				var v40: set<bool> := {v21[869], fm1(v39, v21[869], globalState)};
				v36, v21[869] := v36, v21[869] <==> (!v3 !in v40);
				var v41: array<int> := new int[26];
				v41[421] := v34;
				v35, v41[421] := v35, v33;
			case DC1(cf1) =>
				var v42: seq<int> := [v1, v1, v1, v1, v1];
				var v43: set<seq<int>> := {v42[v1 := v1], v42};
				globalState.f0 := v43 != v43;
				var v44: array<multiset<bool>> := new multiset<bool>[29];
				v44 := if (v21[869]) then v44 else v44;
				var v45: multiset<bool> := multiset{v3, cf1};
				var v46 := DC5(v45);
				var v47, v48, v49 := m0(v46.cf6 - v45, v1, globalState);
				var v50 := DC2(v1, cf1);
				v21, v50 := v21, v50;
		}
		
		v1 := -v1;
		var v51 := DC2(|v0|, v21[869]);
		var v52: map<int, seq<int>> := map[v1 := [-0x147, v51.cf2]];
		var v54: map<char, int> := map['g' := v1];
		var v55: seq<seq<bool>> := [v22];
		var v56: seq<int> := [|v55[|v0|]|, v1];
		v52 := v52[|(map v53 : char | v53 in v54 :: (v53) := (v3))| := v56];
		var v57: map<int, bool> := map[v1 := v21[869]];
		var v58: multiset<bool> := multiset{v3, v21[869], v21[869], if ((fm6(globalState)).cf5 in v57) then v57[(fm6(globalState)).cf5] else v3};
		var v59, v60, v61 := m0(v58 + v58[v3 := v1], fm0(globalState), globalState);
	}
	
	for i2 := v1 to v1 {
		var v62: multiset<bool> := multiset{v3};
		var v63, v64, v65 := m0(v62, v1, globalState);
		var v66 := 'i';
		var v67: T0 := new C0(v66);
		var v68: map<D2, D1> := map[DC6(v67) := fm7(globalState)];
		var v70: seq<bool> := [v63];
		var v71: array<multiset<bool>> := new multiset<bool>[21] [v62 - multiset{v63}, fm2(|(map v69 : int | v69 in {v65} :: (v69 * |v0|) := (v63))|, globalState), multiset(v70), v62[v3 := v64], multiset{v63, v3, v3, v3, v63}, v62, fm2(v1, globalState) * v62, v62, v62 - v62, v62, fm2(-|v0|, globalState), multiset(v70), v62, v62, v62 - v62, v62, v62, v62, multiset{v3}, v62, v62];
		v71[673] := v62;
		var v72: C0 := new C0(v66);
		var v73 := DC6(v67);
		var v74: map<int, C0> := map[-v64 := v72];
		var v75 := DC2(|v74|, true);
		var v76: set<int> := {v64, v1};
		v1, v68, v71[673], globalState.f0, v72 := v64 - v64, v68 + map[v73 := v75], fm2(v64, globalState), v65 > -(v65 + |v76|), v72;
		var v77: array<bool> := new bool[13];
		v77[393] := !v3;
		var v78: seq<int> := [v1];
		var v79 := DC3(v78, -|[v72]|);
		var v80: map<bool, int> := map[v3 := |"dqkshokld"|];
		var v81: map<T0, bool> := map[v67 := v3];
		var v82: map<T0, int> := map[v67 := v1];
		v65, v77[393], v70, globalState.f0 := v79.cf5, v63, ([v70[if (v63 in v80) then v80[v63] else |v0|], v63])[if (v63) then i2 else |v81| := !!(v67 in v82)], v64 != 0x11a;
		var v83 := new C0(v66);
	}
	var v84 := 'm';
	var v85: array<bool> := new bool[6](i3 => !v3);
	v85[445] := false;
	var v86: multiset<bool> := multiset{v3};
	var v87: set<int> := {if (v3 in v86) then v86[v3] else v1};
	var v88: seq<set<int>> := [v87, v87, v87];
	var v89: set<bool> := {v3};
	v0, v3, globalState.f0, v84, v85[445] := v0, !v3, if (v3) then v88[v1] >= v87 else v89 < v89, v84, !(v3 <== v3);
	var v90 := DC2(v1, true);
	var v91, v92, v93 := m0((fm2(v1, globalState))[v85[445] := v1], v90.cf2, globalState);
	if (v3) {
		var v94: map<int, bool> := map[v92 := v85[445]];
		var v95: array<int> := new int[5](i4 => i4 + v93);
		var v96: map<array<int>, array<bool>> := map[v95 := v85];
		var v97: seq<bool> := [v91, !v91, if (v1 in v94) then v94[v1] else v85[445], v95 in v96, v91];
		v91 := !v97[v93];
		var v98: T0 := new C0(v84);
		v98 := v98;
		var v99: map<int, char> := map[|(seq(925, i5  => (v98.f3)))| := v84];
		var v100: map<bool, int> := map[v3 := |v99|];
		var v101: map<bool, bool> := map[v97[0x201] := v3];
		var v102: map<int, map<bool, int>> := map[v93 := v100];
		var v103: seq<int> := [v93, |v100|];
		var v104: seq<map<bool, int>> := [v100, v100];
		globalState.f0, v84, v100 := !(if (v85[445]) then v85[445] else !(if (v85[445] in v101) then v101[v85[445]] else !DC2(v1, v91).cf3)), v84, (if (v103[v1] in v102) then v102[v103[v1]] else v104[v93])[v91 := v92] + v100;
		v91 := v3;
		v101 := v101[false := !v85[445]];
	} else {
		var v105, v106, v107 := m0(v86, v93, globalState);
		v85[445] := v87 !! (set v108 : int | (0x69 <= v108) && (v108 < 0x24) :: (v108 % (if (v105 in v86) then v86[v105] else v1)));
		var v109: seq<bool> := [v85[445]];
		v109 := v109;
		var v110: array<multiset<map<int, bool>>> := new multiset<map<int, bool>>[26](i6 => multiset{map[if (0x162 in v2) then v2[0x162] else v107 := v105], map[v106 := DC1(v85[445]).cf1]});
		var v111: map<int, bool> := map[v107 := true];
		var v112: multiset<map<int, bool>> := multiset{v111, v111};
		v110[988] := v112 + v112;
		var v113: array<int> := new int[9];
		v113[687] := v93;
		var v114: seq<int> := [v93];
		v113[496] := if (v3) then v107 else |v114|;
		v110[988], globalState.f0, v0, v113[687], v113[496] := v112 * multiset{map v115 : int | v115 in v111 :: (v115 * -496) := (true), v111}, !(v86 !! v86[v3 := v1]), if (v85[445]) then fm8(v114, globalState) else "l", v1 / (v107 / -|fm8(v114, globalState)|), v92 * v107;
		if (!(v105 <== v3)) {
			var v116: map<map<int, bool>, multiset<bool>> := map[v111 := v86];
			var v117 := DC1(v85[445]);
			var v118, v119, v120 := m0(v86 * (if (v111 in v116) then v116[v111] else multiset{v105})[v117.cf1 := |v111|], v107, globalState);
			var v121 := DC4();
			v121 := DC4();
			v85[445] := (if (v91) then v90.cf2 else v93) != -(|{|v86|, v93, v93, |v109|, v113[687]}| - v1);
			v113[687], v120 := v120, v106 / (v92 + v93);
			v113[687] := v119;
		} else {
			var v122 := new C0(v84);
			v85[445] := !v91;
			v113[687] := 0x202;
			var v123 := new C0(v84);
			var v124: array<string> := new string[26](i7 => if (v91) then v0 else seq(0x1ef, i8  => (v84)));
			var v125 := DC8(v124);
			v113[687], v124, v93, globalState.f0 := 633 * (v114[v93] + -0x2cf), v125.cf9, v93, (if (v91 in v86) then v86[v91] else fm0(globalState)) == v107;
		}
		
	}
	
	var v126: seq<int> := [v92];
	var v127: C0 := new C0(v84);
	var v128: map<int, C0> := map[v92 := v127];
	var v129: seq<C0> := [if (v93 in v128) then v128[v93] else v127, v127];
	v1 := v1 + |(v126 + [0x27e, v92, |v129|])|;
	v1 := v93;
	v3 := !v91;
	var i9 := 0;
	while (v126[v93 := v92] == (v126 + v126))
		decreases 100 - i9
	{
		if (i9 >= 100) {
			break;
		}
		
		i9 := i9 + 1;
		var v130: seq<map<int, bool>> := [map[v93 := v91], map[|v126| := DC2(0x3a, v85[445]).cf3]];
		v130 := v130;
		v93 := v126[v1];
		v92 := v1;
		var v131: T0 := new C0('g');
		v131 := v131;
	}
	globalState.f0 := v3 && v91;
	v91 := v3;
}