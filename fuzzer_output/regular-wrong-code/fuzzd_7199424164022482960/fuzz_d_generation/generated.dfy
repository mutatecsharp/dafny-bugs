datatype D0 = DC0(cf0: int, cf1: map<string, int>, cf2: multiset<int>) | DC1(cf3: D0)
datatype D1 = DC3(cf5: set<bool>) | DC4(cf6: string, cf7: bool, cf8: seq<string>, cf9: seq<char>, cf10: char) | DC2(cf4: map<bool, int>) | DC5(cf11: D1)
datatype D2 = DC7(cf13: array<bool>) | DC6(cf12: C0) | DC8(cf14: D2)
datatype D3 = DC10(cf16: bool, cf17: array<int>, cf18: string) | DC9(cf15: map<int, string>)
datatype D4 = DC12 | DC11(cf19: map<int, bool>)
class GlobalState {
	var f0 : int
	const f1 : set<bool>
	const f2 : bool
	const f3 : bool
	var f4 : int
	const f5 : int
	const f6 : bool
	const f7 : array<map<int, bool>>
	const f8 : seq<int>
	const f9 : multiset<int>
	var f10 : int
	const f11 : int
	const f12 : char
	var f13 : bool
	const f14 : int
	var f15 : int
	const f16 : array<bool>
	const f17 : string
	constructor (f0 : int, f1 : set<bool>, f2 : bool, f3 : bool, f4 : int, f5 : int, f6 : bool, f7 : array<map<int, bool>>, f8 : seq<int>, f9 : multiset<int>, f10 : int, f11 : int, f12 : char, f13 : bool, f14 : int, f15 : int, f16 : array<bool>, f17 : string) {
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

function fm0(p0: D0, globalState: GlobalState): int {
	if (true) then 0x304 else 0x392 * 0x14d
}
function fm1(p0: bool, globalState: GlobalState): bool {
	DC0(|{map[true := |(seq(918, i0  => (137)))|], map[true := |[true]|]}|, map["tq" := -0x164], multiset{-192, |map[true := true]|}).cf0 <= (----610 - |map[true := true]|)
}
function fm2(p0: bool, globalState: GlobalState): multiset<int> {
	multiset{0x188}
}
function fm3(globalState: GlobalState): D0 {
	DC0(|"m"|, map["xnjhhmdfk" := |"tslm"|] + map["lhmcdhe" := 0x31e], multiset{-|multiset{false}|, |"xoksdycy"|} + multiset{0x12a})
}
function fm4(p0: int, p1: int, p2: int, p3: int, globalState: GlobalState): set<bool> {
	{!!false, false, false, true} - ({false, true, false} - {!true, false})
}
function fm5(p0: int, p1: int, p2: int, p3: int, globalState: GlobalState): seq<char> {
	(['l'] + (seq(780, i0  => ('s')))) + ['v', 'o']
}
function fm6(p0: bool, globalState: GlobalState): char {
	'c'
}
function fm7(p0: int, p1: seq<bool>, p2: bool, globalState: GlobalState): multiset<bool> {
	(multiset{true, false} - multiset{false, true, false}) + (multiset([false, true, true, !true]) * multiset{true})
}
function fm8(globalState: GlobalState): map<int, bool> {
	(if (false) then DC11(map v0 : int | v0 in (seq(0x289, i0  => (-|{0xcc, 26}|))) :: (v0 / 0x134) := (true)) else DC11(map v1 : int | v1 in (map v2 : int | v2 in map[|"t"| := |{|[false]|}|] :: (v2 + |"qc"|) := (|"mrjysrs"|)) :: (v1 - 0x1b0) := (true))).cf19
}
method m0(p0: D0, p1: int, globalState: GlobalState) returns (r0: int, r1: seq<int>, r2: set<int>, r3: bool) {
	var v0 := "hkx";
	var v1: seq<string> := [v0];
	var v2 := true;
	var v3: map<bool, int> := map[v2 := |(seq(0x386, i0  => ('f')))|];
	var v4: map<int, bool> := map[|[p1]| := v2];
	var v5: seq<bool> := [v2, v2];
	var v6: map<string, int> := map[v0 := p1];
	var v7 := 'i';
	var v8: map<char, bool> := map[v7 := true];
	var v9: multiset<int> := multiset{|v8|, -0x221, p1, p1};
	var v10 := DC0(p1, v6, v9);
	var v11: map<int, string> := map[fm0(v10, globalState) := v0];
	var v12: map<bool, string> := map[if (|v5| in v4) then v4[|v5|] else v2 := if (fm0(v10, globalState) in v11) then v11[fm0(v10, globalState)] else v0];
	var v13 := DC4(v0, v2, v1, v0, v7);
	var v14: array<string> := new string[28] [v1[p1] + v0, v0, fm5(|v0|, p1, p1, |v3|, globalState), v0, v0, v0 + v0, seq(713, i1  => (v0[0x275])), "hrekd", v0, v0, v0, v0, if (v2 in v12) then v12[v2] else v0[385 := v13.cf10], v0, v0, v0 + v0, v0, v0, v0, v0, v0, v0, "mxsg", v0 + v0, v0, v0, v0, v0];
	v14 := v14;
	var v15 := new C0(p1 + p1);
	var v16: set<bool> := {v2, v2};
	var v17 := DC3(v16);
	r3 := match v17 {
		case DC3(cf5) => v0 < v0
		case DC4(cf6, cf7, cf8, cf9, cf10) => cf7
		case DC2(cf4) => v2
		case DC5(cf11) => v2 || v2
	};
	var v18: array<int> := new int[19](i2 => i2 * v15.f18);
	v18 := new int[17](i3 => i3 * -0x3da);
	for i4 := v15.f18 to v15.f18 {
		v18[580] := v15.f18;
		v18[580] := if (v2) then v15.f18 else p1 * p1;
		v0 := v0 + v0;
		var v19: seq<int> := [0x5f, p1];
		var v20: array<bool> := new bool[21] [v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, v2, true, v2, v2, v2, v2, v2, v2, v2, v2, v2];
		var v21: seq<array<bool>> := [v20, v20, v20, v20];
		v5 := (v5 + (v5 + v5))[v10.cf0 := |v19| < -|v21|];
		match (v13) {
			case DC3(cf5) =>
				globalState.f13 := 0x2d1 > i4;
				globalState.f13 := v2;
				globalState.f10 := v15.f18;
				r3 := v2;
			case DC4(cf6, cf7, cf8, cf9, cf10) =>
				var v22: seq<C0> := [v15, v15, v15];
				v22 := v22 + v22;
				var v23: map<seq<int>, int> := map[v19 := v15.f18];
				v18[580], v23, globalState.f4 := v15.f18 / -fm0(v10, globalState), v23, |(((seq(-273, i5  => (|{if (v15.f18 in v9) then v9[v15.f18] else |v0|}|))) + v19) + (v19 + v19))|;
				v18[580] := i4;
				var v24: array<char> := new char[25](i6 => v0[-0x147]);
				var v25: multiset<array<char>> := multiset{v24};
				v25 := v25;
			case DC2(cf4) =>
				var v26 := new C0(i4 - fm0(v10, globalState));
				r3 := v15.f18 > p1;
				var v27: map<C0, bool> := map[v26 := v2];
				var v28: array<map<C0, bool>> := new map<C0, bool>[11] [v27, v27, map[v26 := !v2], v27, v27[v26 := v2], v27 + v27, v27, v27, v27, v27, v27];
				var v29: multiset<bool> := multiset{v2};
				v28[487] := map[v26 := v5[|v29|]] + v27;
				v28[487] := v27;
				var v30: map<int, int> := map[v15.f18 := v15.f18];
				var v31: array<int> := new int[14] [i4, v18[580], v26.f18, v15.f18, v26.f18, v18[580], -v15.f18, v18[580], -v15.f18, v15.f18, v15.f18, |v30[p1 := |v5|]|, v18[580], v15.f18];
				var v32 := DC10(if (v2) then true else v2, v31, v0 + v0);
				v32 := v32;
			case DC5(cf11) =>
				globalState.f0 := i4;
				v20[960] := v2;
				v20[960] := v2;
				var v33: array<char> := new char[18](i7 => v7);
				v33 := v33;
				v20[960] := v20[960] ==> v20[960];
		}
		
	}
	match (v13.(cf6 := seq(335, i8  => (v7)), cf8 := v1)) {
		case DC3(cf5) =>
			var v34: array<bool> := new bool[5];
			var v35 := DC7(v34);
			var v36 := DC8(v35);
			v36 := v36;
			v34[898] := v2 ==> v2;
			v34[898] := v2;
			v34 := if (v2) then v34 else v34;
			var v37 := DC1(fm3(globalState));
			v37 := p0;
		case DC4(cf6, cf7, cf8, cf9, cf10) =>
			var v38: array<bool> := new bool[27](i9 => v5[|map[|{v15.f18, v15.f18, |v3|}| := p1]|] !in v3);
			v15, v38 := v15, v38;
			var v39: multiset<bool> := multiset{cf7, cf7};
			v18[62] := |(if (v2) then v39 else fm7(p1, v5, v2, globalState))|;
			v18[62] := p1;
			var v40: map<C0, int> := map[v15 := v15.f18];
			v40 := (v40 + v40) + v40;
			v7 := v7;
		case DC2(cf4) =>
			var v41: map<bool, bool> := map[v2 := true];
			v41 := v41 + v41;
			if (v2) {
				var v42: array<bool> := new bool[15];
				var v43: seq<array<bool>> := [v42];
				var v44: map<C0, seq<array<bool>>> := map[v15 := v43];
				var v45: map<bool, array<int>> := map[v2 := v18];
				var v46: map<int, seq<array<bool>>> := map[|v45| := v43];
				var v47: array<seq<array<bool>>> := new seq<array<bool>>[9] [v43, v43 + (if (v15 in v44) then v44[v15] else v43), v43[p1 := v42] + v43, v43, if (p1 in v46) then v46[p1] else v43, v43, v43, v43, v43 + v43];
				v47[565] := [v42, v42] + v43;
				v18[368] := --|v16|;
				var v48 := DC10(v2 ==> false, v18, ((seq(-28, i10  => (v7))) + "i")[v15.f18 := v7]);
				var v49: map<bool, C0> := map[v2 := v15];
				v47[565], v0, v18[368], v48 := v43 + v43, v0, |(v49 + v49)|, v48;
				var v50: map<int, int> := map[v18[368] := v15.f18];
				v50 := v50[v18[368] := v15.f18];
				globalState.f0 := p1;
				v42 := v42;
				var v51: multiset<bool> := multiset{!v2, true};
				var v52 := new C0(|(v51 * v51)|);
			} else {
				var v53: array<set<int>> := new set<int>[20];
				v53 := v53;
				var v54 := DC2(v3);
				var v55: map<D1, int> := map[v54 := v15.f18];
				v55 := v55;
				var v56: seq<int> := [v15.f18, v15.f18, v15.f18];
				var v57: map<bool, seq<int>> := map[v2 := v56 + v56];
				r1 := if (v2 in v57) then v57[v2] else if (v2) then v56 else [-|cf4|];
				var v58: array<bool> := new bool[10](i11 => v2);
				v58[781] := true;
				v58[781] := v2;
				globalState.f15 := v15.f18;
			}
			
			var v59: multiset<bool> := multiset{false};
			r3 := (v59 - multiset{!fm1(v2, globalState), v2}) >= (multiset{v2})[!v2 := p1];
			var v60: array<seq<int>> := new seq<int>[24];
			v60 := new seq<int>[29];
		case DC5(cf11) =>
			var v61: seq<seq<bool>> := [v5, v5, v5];
			var v62: map<int, char> := map[-890 := v7];
			var v63: seq<int> := [|fm8(globalState)|, v15.f18];
			var v64: seq<int> := [|v63|];
			var v65 := DC10(v2, v18, v0);
			var v66: array<bool> := new bool[22] [true, v13.cf7, v2, fm1(v2, globalState), v61 <= [v5], !true, v2, v2, v15.f18 != v15.f18, fm1(v2, globalState), v2, v62 == map[v64[|v0|] := v7], !v2, true, v5[p1], v2, v15.f18 > p1, !v2, v65.cf16, v2, v2, fm1(!false, globalState)];
			v66 := v66;
			var v67: array<seq<C0>> := new seq<C0>[21];
			v67 := v67;
			r3 := v2;
			var v68: map<array<bool>, C0> := map[v66 := v15];
			v68 := v68;
	}
	
	r0 := -match p0 {
		case DC0(cf0, cf1, cf2) => -(cf0 * 483)
		case DC1(cf3) => v15.f18
	};
	var v69: seq<int> := [v15.f18, |"jxu"|];
	r1 := v69;
	var v70: set<int> := {p1, fm0(v10, globalState)};
	r2 := v70 + v70;
	r3 := v2;
}
class C0 {
	const f18 : int
	constructor (f18 : int) {
		this.f18 := f18;
	}
	
}

method Main() {
	var v0 := true;
	var v1: set<bool> := {v0};
	var v2: array<map<int, bool>> := new map<int, bool>[15];
	var v3 := 0x27f;
	var v4: multiset<int> := multiset{v3, -408};
	var v5: multiset<bool> := multiset{v0, v0};
	var v6: seq<int> := [|v4|, |v5|];
	var v7: array<bool> := new bool[11](i0 => v0);
	var v8 := "mgx";
	var globalState := new GlobalState(-0xce, v1 * v1, false, false, 633, 0x301, true, v2, v6[v3 := v3] + v6, v4, -0x2e2, 0x2ca, 'r', false, 0x2f8, 0x28e, v7, v8);
	var v9: map<int, int> := map[v3 := v3];
	var v10: map<bool, int> := map[v0 := |v9|];
	var v11: map<string, int> := map[v8 := if (v0 in v10) then v10[v0] else v3];
	var v12 := DC0(v3, v11, v4);
	match (v12) {
		case DC0(cf0, cf1, cf2) =>
			var v13 := DC1(DC1(DC0(|v6|, v11, cf2)));
			var v14 := DC1(v13);
			var v15, v16, v17, v18 := m0(DC1(v14), fm0(v12, globalState), globalState);
			v0 := !(fm0(v12, globalState) == (-v3 - cf0));
			if (fm1(v0, globalState)) {
				var v19 := DC1(v13);
				var v20, v21, v22, v23 := m0(if (v0) then v19 else v19, |v8| % cf0, globalState);
				globalState.f13 := fm1(v23, globalState);
				globalState.f13 := v18;
				v10 := v10[v8 < v8 := v3];
				globalState.f0 := v3;
			} else {
				v12 := v12;
				var v24 := DC1(v13);
				var v25, v26, v27, v28 := m0(v24, |v8| * v3, globalState);
				globalState.f13 := true;
				v3 := 0x1d8 % |v8|;
				var v29 := new C0(v25);
			}
			
			var v30: array<int> := new int[6];
			v30[841] := v3 - cf0;
			v30[841], v3, v0 := cf0, v3, !((cf0 * v15) > |v9|);
		case DC1(cf3) =>
			match (if (false) then v12 else v12) {
				case DC0(cf0, cf1, cf2) =>
					var v31: map<int, bool> := map[if (v0 in v5) then v5[v0] else v3 := v0];
					v0 := fm1(if (cf0 in v31) then v31[cf0] else v0, globalState);
					globalState.f15 := if (v1 !! v1) then |map[v0 := v0]| + |v8| else -v3;
					var v33: map<string, seq<bool>> := map[v8 := [!fm1(false, globalState)]];
					var v34 := DC0(cf0, map v32 : string | v32 in v33 :: (v32) := (v12.cf0), v4[if (cf0 in v4) then v4[cf0] else v3 := v3]);
					var v35 := DC1(v34);
					var v36 := DC1(v35);
					var v37 := DC1(v36);
					var v38: set<string> := {v8};
					var v39, v40, v41, v42 := m0(v37, |[---cf0, v3, -|(seq(0x33e, i1  => (v8[v3])))|, v3, -|v38|]|, globalState);
					var v43: C0 := new C0(cf0);
					v43 := v43;
				case DC1(cf3) =>
					var v44: C0 := new C0(v3 - v3);
					v44 := v44;
					var v45 := DC0(v3, v11, v4);
					var v46, v47, v48, v49 := m0(DC1(v45), DC0(|v6|, v11, fm2(v0, globalState)).cf0, globalState);
					var v50 := DC1(v45);
					var v51, v52, v53, v54 := m0(v50, |v5|, globalState);
					var v55, v56, v57, v58 := m0(v50.(cf3 := v45), |({v44.f18, v46} - v48)|, globalState);
			}
			
			var v59: multiset<multiset<bool>> := multiset{multiset{false}, multiset{v0}, v5, v5};
			var v60: seq<multiset<bool>> := [v5];
			v59 := multiset(v60 + (seq(-680, i2  => (multiset([v0, false, v0])))));
			globalState.f0 := 0xc;
			globalState.f10 := v3;
	}
	
	var v62: seq<multiset<int>> := [multiset(v6), v4, v4, v4];
	var v63: C0 := new C0(|v4|);
	var v64: map<bool, C0> := map[v0 := v63];
	var v65 := DC2(map[!fm1(false, globalState) := v3]);
	var v66: array<int> := new int[14] [fm0(fm3(globalState), globalState), v3, v3 + v3, v3, v3, -0x1b5, v6[fm0(DC0(v3, map[v8 := v3], v4), globalState)], -|(map v61 : multiset<int> | v61 in v62 :: (v61) := (v3))|, v3, |(v8 + v8)|, |v64|, |v65.cf4|, |v10| + v63.f18, v3];
	v66[580] := v63.f18 * |v8|;
	v66[831] := v63.f18 * v3;
	var v67: set<int> := {v63.f18, v63.f18};
	var v68: seq<bool> := [v0];
	globalState.f10, v66[580], v66[831] := if (v0) then v3 else |v67|, (v63.f18 - |v8|) % |v68|, v63.f18;
	var v69 := new C0(v3);
	v4 := v4;
	var v70: map<int, bool> := map[v63.f18 := !v0];
	if (if ((v63.f18 % -0x389) in v70) then v70[v63.f18 % -0x389] else !fm1(v0, globalState)) {
		var v71: array<map<bool, int>> := new map<bool, int>[4](i3 => if (v0) then v10 else v10);
		v71[584] := v10;
		v71[584] := v10;
		var v72: array<multiset<array<D1>>> := new multiset<array<D1>>[13];
		var v73: array<D1> := new D1[2] [v65, v65];
		v72[458] := multiset{v73, v73, v73, v73, v73};
		var v74: multiset<array<D1>> := multiset{v73, v73};
		v72[458] := (v74 - multiset{v73, v73}) * (v74 * v74);
		var v75: array<C0> := new C0[25];
		v75[299] := v63;
		v75[299] := v63;
		var v76: array<map<int, C0>> := new map<int, C0>[17];
		v76 := v76;
		var v77 := DC3(v1);
		var v78: array<D1> := new D1[16] [DC3(v1), v77, v77, v77, v77, v77, DC3(v1), DC3(v1), v77, v77, if (true) then v77 else v77, v77, v77, DC3(fm4(v63.f18, v69.f18, v66[580], v63.f18, globalState)), DC3(v1), v77];
		v78[53] := v77;
		v78[53] := v77;
	} else {
		v7[826] := v0;
		v7[826] := v0;
		var v79 := DC0(v63.f18, (v11[v8 := |multiset(v68)|])[v8 := v63.f18], v4);
		var v80 := DC1(v79);
		var v81 := DC1(v79);
		var v82, v83, v84, v85 := m0(v81, if (v0) then v63.f18 else -0x5e, globalState);
		var v86 := new C0(v3);
		var v87: map<bool, D0> := map[v7[826] := v81];
		v87 := v87[v85 := DC1(v79)];
		globalState.f13 := v0;
	}
	
	v9 := v9[v63.f18 := v63.f18];
	var v88 := 'j';
	var v89: map<multiset<int>, char> := map[multiset{v69.f18, v66[580], v3, v3, -242} := v88];
	var v90: map<bool, map<multiset<int>, char>> := map[v0 := v89];
	v89 := if (true in v90) then v90[true] else v89;
	globalState.f4 := -v3;
	v0 := !v0;
	var v91: array<seq<int>> := new seq<int>[18];
	v91[747] := v6;
	v91[747] := v6 + v6;
	v7 := v7;
	v0 := v0;
	if (v66[580] == |v8|) {
		var v92: seq<C0> := [v69, v69, v63, v69];
		var v93: map<char, bool> := map[v88 := v92 < [v63, v63]];
		v93 := v93[v88 := v0];
		if (v0) {
			var v94: map<int, D0> := map[54 := v12];
			var v96: seq<set<int>> := [v67, v67, set v95 : int | v95 in v94 :: (v95 + 514)];
			v96 := v96 + v96;
			var v97 := DC3(v1);
			var v98: map<set<bool>, string> := map[v1 := v8];
			v0 := (v97.(cf5 := v1)).cf5 !in v98;
			var v99 := DC0(v69.f18, v11, v4);
			var v100 := DC1(v99);
			var v101, v102, v103, v104 := m0(v100, fm0(v12, globalState), globalState);
			var v105: map<int, seq<C0>> := map[v101 := v92];
			var v106: map<bool, seq<C0>> := map[v0 := v92];
			var v107: multiset<seq<C0>> := multiset{if (v3 in v105) then v105[v3] else if (v0 in v106) then v106[v0] else v92, v92};
			v66[580] := |(v107 * v107)| - |"tprjqefl"|;
			v92 := v92;
		} else {
			globalState.f13 := |v92| < v69.f18;
			var v108: seq<string> := [v8, v8];
			var v109 := DC4("cs", v0, v108, fm5(|v4|, v91[747][fm0(v12, globalState)], v69.f18, v63.f18, globalState), v88);
			globalState.f15, v63, v70, globalState.f13, v66[580] := |(v8 + ((seq(650, i4  => (v88))) + v8))|, v69, v70, v109.cf7, v63.f18;
			var v110 := DC6(v69);
			var v111: array<C0> := new C0[29] [v69, if (!v0) then v69 else v110.cf12, v69, v63, v63, v63, v63, v63, v63, v63, v69, v92[v66[580]], v63, v63, v69, v69, v63, v63, v92[v69.f18], v63, v92[917], v69, v69, v63, v69, v63, v63, v69, v69];
			v111[892] := v69;
			v111[892] := v69;
			v12 := if (v0) then v12 else v12;
			v7[51] := v0 ==> v0;
			v7[51] := v0;
		}
		
		var v112 := DC0(v66[580], v11[v8 := v63.f18], multiset{-v3, v69.f18});
		var v113 := DC1(v112);
		var v114, v115, v116, v117 := m0(DC1(v113), v69.f18 - v69.f18, globalState);
		v0 := v0;
		v7[708] := "eqmqr" <= v8;
		v7[708] := v117;
	} else {
		var v118 := new C0(v63.f18);
		globalState.f4 := v63.f18;
		globalState.f4, v69, globalState.f13 := v63.f18, v69, !fm1(v0, globalState);
		var v119: map<char, bool> := map[v88 := fm1(v0, globalState)];
		v0, globalState.f13, v119 := !v0, v0, v119;
		var v120: map<int, string> := map[v69.f18 := seq(-0x94, i5  => (v88))];
		var v121 := DC9(v120);
		var v123: array<map<int, string>> := new map<int, string>[7] [v120, if (v0) then map[0x2b8 := v8] else v120, v121.cf15, map v122 : int | (905 <= v122) && (v122 < 0x2ca) :: (v122 / |v68|) := (v8), map[v66[580] := v8], v120, v120];
		v123[856] := v120;
		v123[856] := if (v0) then map[v69.f18 := v8] else map[-|v8| := seq(-0x10a, i6  => (v88))];
	}
	
	v88 := fm6(v67 < v67, globalState);
	v10 := v10[v0 <==> !v0 := v63.f18];
	if (v0) {
		v66[580] := (if (v0) then v69.f18 else 706) % v63.f18;
		globalState.f4, globalState.f4 := v63.f18, fm0(v12, globalState);
		var v124 := new C0(|(seq(51, i7  => (v68)))[v66[580] := v68]|);
		v11 := v11[v8 := 0x2f7];
		globalState.f4 := (v69.f18 / v6[v124.f18]) % fm0(v12, globalState);
	} else {
		var v125: map<int, string> := map[|fm7(|[if (v69.f18 in v70) then v70[v69.f18] else v0, !v0, v0]|, v68, v0, globalState)| := v8];
		var v126: array<D2> := new D2[19];
		var v127: map<map<int, string>, array<D2>> := map[v125 := v126];
		v127 := v127[v125 := v126];
		var v128 := new C0(0x364 + 0x337);
		globalState.f13 := if (true) then v0 else !v0;
		var v129: map<bool, array<int>> := map[v0 := v66];
		v66[580] := (-v3 + v63.f18) - |v129|;
		globalState.f0 := v3 + (326 + |v4|);
	}
	
}