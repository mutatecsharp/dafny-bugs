import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_
import _dafny
import System_

# Module: module_

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def abs(x):
        if (x) < (0):
            return (-1) * (x)
        elif True:
            return x

    @staticmethod
    def safeIndex(x, length):
        if (x) < (0):
            return 0
        elif (x) >= (length):
            return _dafny.euclidian_modulus(x, length)
        elif True:
            return x

    @staticmethod
    def safeDivisionInt(x1, x2):
        if (x2) == (0):
            return x1
        elif True:
            return _dafny.euclidian_division(x1, x2)

    @staticmethod
    def safeModuloInt(x1, x2):
        if (x2) == (0):
            return x1
        elif True:
            return _dafny.euclidian_modulus(x1, x2)

    @staticmethod
    def fm0(globalState):
        return (_dafny.Set({True, True})) | (_dafny.Set({True}))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return True

    @staticmethod
    def fm2(p0, globalState):
        return D0_DC1()

    @staticmethod
    def fm3(globalState):
        source0_ = D0_DC0(True)
        if source0_.is_DC1:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttioiues"))
        elif True:
            d_0___mcc_h0_ = source0_.cf0
            d_1_cf0_ = d_0___mcc_h0_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return -413

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "meafobh"))) for d_2_i0_ in range(default__.abs(857))])

    @staticmethod
    def fm7(p0, globalState):
        source1_ = D0_DC1()
        if source1_.is_DC1:
            return D0_DC0(True)
        elif True:
            d_3___mcc_h0_ = source1_.cf0
            d_4_cf0_ = d_3___mcc_h0_
            return D0_DC0(d_4_cf0_)

    @staticmethod
    def m0(p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        r2: bool = False
        d_5_v0_: _dafny.Array
        nw0_ = _dafny.Array(False, 7)
        d_5_v0_ = nw0_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_5_v0_).length(0)):
            d_6_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_6_i0_)) and ((d_6_i0_) < ((d_5_v0_).length(0)))):
                (d_5_v0_)[(d_6_i0_)] = (p0 if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icxytex")))) != (410) else p0)
        d_7_v1_: int
        d_7_v1_ = 14
        d_8_v2_: _dafny.Seq
        d_8_v2_ = _dafny.SeqWithoutIsStrInference([d_7_v1_])
        d_9_v3_: C0
        nw1_ = C0()
        nw1_.ctor__(d_8_v2_)
        d_9_v3_ = nw1_
        d_10_v4_: _dafny.Map
        d_10_v4_ = _dafny.Map({d_9_v3_: p1})
        d_11_v5_: _dafny.Map
        d_11_v5_ = _dafny.Map({d_7_v1_: d_10_v4_})
        d_12_v6_: _dafny.Seq
        d_12_v6_ = _dafny.SeqWithoutIsStrInference([d_11_v5_])
        d_13_v7_: D1
        d_13_v7_ = D1_DC4((d_12_v6_)[default__.safeIndex(d_7_v1_, len(d_12_v6_))], d_7_v1_)
        source2_ = default__.fm2((d_13_v7_).cf5, globalState)
        if source2_.is_DC1:
            r2 = False
            d_14_v8_: _dafny.Seq
            d_14_v8_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_15_v9_: _dafny.Map
            d_15_v9_ = _dafny.Map({(d_14_v8_)[default__.safeIndex(d_7_v1_, len(d_14_v8_))]: d_7_v1_})
            d_16_v10_: _dafny.Array
            nw2_ = _dafny.Array(int(0), 12)
            d_16_v10_ = nw2_
            d_17_v11_: _dafny.MultiSet
            d_17_v11_ = _dafny.MultiSet([d_7_v1_, default__.fm4(d_7_v1_, d_7_v1_, p0, p0, globalState), len(d_15_v9_), len(_dafny.SeqWithoutIsStrInference([d_16_v10_, d_16_v10_, d_16_v10_]))])
            d_18_v12_: D0
            d_18_v12_ = D0_DC0((_dafny.MultiSet([d_7_v1_])) == (d_17_v11_))
            d_19_v13_: _dafny.Seq
            d_19_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "avymisae"))
            d_20_v14_: _dafny.Map
            d_20_v14_ = _dafny.Map({p1: d_19_v13_})
            d_21_v15_: _dafny.Map
            d_21_v15_ = _dafny.Map({not(p1): ((d_20_v14_)[p0] if (p0) in (d_20_v14_) else d_19_v13_)})
            d_22_v16_: _dafny.Array
            nw3_ = _dafny.Array(None, 16)
            nw3_[int(0)] = d_19_v13_
            nw3_[int(1)] = (d_19_v13_) + (d_19_v13_)
            nw3_[int(2)] = d_19_v13_
            nw3_[int(3)] = d_19_v13_
            nw3_[int(4)] = (d_19_v13_) + (d_19_v13_)
            nw3_[int(5)] = ((d_21_v15_)[p1] if (p1) in (d_21_v15_) else (d_19_v13_).set(default__.safeIndex(d_7_v1_, len(d_19_v13_)), _dafny.CodePoint('w')))
            nw3_[int(6)] = d_19_v13_
            nw3_[int(7)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_23_i1_ in range(default__.abs(490))])
            nw3_[int(8)] = d_19_v13_
            nw3_[int(9)] = (d_19_v13_ if p0 else d_19_v13_)
            nw3_[int(10)] = d_19_v13_
            nw3_[int(11)] = (d_19_v13_) + (d_19_v13_)
            nw3_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwsgvn"))
            nw3_[int(13)] = default__.fm3(globalState)
            nw3_[int(14)] = (d_19_v13_) + (d_19_v13_)
            nw3_[int(15)] = d_19_v13_
            d_22_v16_ = nw3_
            d_24_v17_: str
            d_24_v17_ = _dafny.CodePoint('v')
            d_25_v18_: _dafny.Array
            nw4_ = _dafny.Array(None, 4)
            nw4_[int(0)] = d_24_v17_
            nw4_[int(1)] = d_24_v17_
            nw4_[int(2)] = d_24_v17_
            nw4_[int(3)] = d_24_v17_
            d_25_v18_ = nw4_
            d_26_v19_: _dafny.Seq
            d_26_v19_ = _dafny.SeqWithoutIsStrInference([d_9_v3_])
            d_27_v20_: _dafny.Array
            nw5_ = _dafny.Array(None, 10)
            nw5_[int(0)] = (d_26_v19_)[default__.safeIndex(d_7_v1_, len(d_26_v19_))]
            nw5_[int(1)] = d_9_v3_
            nw5_[int(2)] = d_9_v3_
            nw5_[int(3)] = d_9_v3_
            nw5_[int(4)] = d_9_v3_
            nw5_[int(5)] = d_9_v3_
            nw5_[int(6)] = d_9_v3_
            nw5_[int(7)] = d_9_v3_
            nw5_[int(8)] = d_9_v3_
            nw5_[int(9)] = d_9_v3_
            d_27_v20_ = nw5_
            d_28_v21_: _dafny.Map
            d_28_v21_ = _dafny.Map({len(d_14_v8_): d_27_v20_})
            d_29_v22_: _dafny.Map
            d_29_v22_ = _dafny.Map({d_5_v0_: ((d_28_v21_)[d_7_v1_] if (d_7_v1_) in (d_28_v21_) else d_27_v20_)})
            d_30_v23_: _dafny.Map
            d_30_v23_ = _dafny.Map({d_25_v18_: d_29_v22_})
            d_31_v24_: _dafny.Seq
            d_31_v24_ = _dafny.SeqWithoutIsStrInference([d_29_v22_])
            d_32_v25_: _dafny.Array
            nw6_ = _dafny.Array(None, 23)
            nw6_[int(0)] = ((d_30_v23_)[d_25_v18_] if (d_25_v18_) in (d_30_v23_) else _dafny.Map({d_5_v0_: d_27_v20_}))
            nw6_[int(1)] = d_29_v22_
            nw6_[int(2)] = (d_29_v22_) | (d_29_v22_)
            nw6_[int(3)] = d_29_v22_
            nw6_[int(4)] = d_29_v22_
            nw6_[int(5)] = d_29_v22_
            nw6_[int(6)] = d_29_v22_
            nw6_[int(7)] = d_29_v22_
            nw6_[int(8)] = (d_29_v22_) | (_dafny.Map({d_5_v0_: d_27_v20_}))
            nw6_[int(9)] = (D3_DC6(_dafny.Map({d_5_v0_: d_27_v20_}))).cf7
            nw6_[int(10)] = d_29_v22_
            nw6_[int(11)] = d_29_v22_
            nw6_[int(12)] = d_29_v22_
            nw6_[int(13)] = d_29_v22_
            nw6_[int(14)] = d_29_v22_
            nw6_[int(15)] = (d_29_v22_) | (_dafny.Map({d_5_v0_: d_27_v20_}))
            nw6_[int(16)] = d_29_v22_
            nw6_[int(17)] = _dafny.Map({d_5_v0_: d_27_v20_})
            nw6_[int(18)] = (d_31_v24_)[default__.safeIndex(d_7_v1_, len(d_31_v24_))]
            nw6_[int(19)] = (d_29_v22_) | (d_29_v22_)
            nw6_[int(20)] = (d_29_v22_) | (d_29_v22_)
            nw6_[int(21)] = d_29_v22_
            nw6_[int(22)] = (d_29_v22_).set(d_5_v0_, d_27_v20_)
            d_32_v25_ = nw6_
            index0_ = default__.safeIndex(548, (d_32_v25_).length(0))
            (d_32_v25_)[index0_] = d_29_v22_
            index1_ = default__.safeIndex(548, (d_32_v25_).length(0))
            rhs0_ = p0
            rhs1_ = d_18_v12_
            rhs2_ = d_22_v16_
            rhs3_ = (d_29_v22_) | (d_29_v22_)
            lhs0_ = d_32_v25_
            lhs1_ = default__.safeIndex(548, (d_32_v25_).length(0))
            r2 = rhs0_
            d_18_v12_ = rhs1_
            d_22_v16_ = rhs2_
            lhs0_[lhs1_] = rhs3_
            d_33_v26_: _dafny.Array
            nw7_ = _dafny.Array(_dafny.Array(None, 0), 25)
            d_33_v26_ = nw7_
            index2_ = default__.safeIndex(132, (d_33_v26_).length(0))
            (d_33_v26_)[index2_] = d_16_v10_
            d_34_v27_: _dafny.Map
            d_34_v27_ = _dafny.Map({p1: d_17_v11_})
            d_35_v28_: _dafny.MultiSet
            d_35_v28_ = _dafny.MultiSet([p1, p0])
            d_36_v29_: _dafny.Seq
            d_36_v29_ = _dafny.SeqWithoutIsStrInference([d_16_v10_])
            d_37_v30_: _dafny.Map
            d_37_v30_ = _dafny.Map({(d_17_v11_).ispropersubset(((d_34_v27_)[p0] if (p0) in (d_34_v27_) else _dafny.MultiSet([d_7_v1_, (d_35_v28_).cardinality, 57, d_7_v1_]))): (d_36_v29_)[default__.safeIndex(d_7_v1_, len(d_36_v29_))]})
            d_38_v31_: _dafny.Array
            d_38_v31_ = d_16_v10_
            index3_ = default__.safeIndex(132, (d_33_v26_).length(0))
            (d_33_v26_)[index3_] = ((d_37_v30_)[p1] if (p1) in (d_37_v30_) else (d_38_v31_))
            r2 = not(not(not(p0)))
        elif True:
            d_39___mcc_h0_ = source2_.cf0
            d_40_cf0_ = d_39___mcc_h0_
            d_41_v32_: _dafny.Seq
            d_41_v32_ = _dafny.SeqWithoutIsStrInference([d_7_v1_ for d_42_i2_ in range(default__.abs(-50))])
            d_8_v2_ = ((d_9_v3_).f0) + ((d_41_v32_))
            d_43_v33_: str
            d_43_v33_ = _dafny.CodePoint('y')
            d_43_v33_ = d_43_v33_
            r1 = (d_7_v1_) - (-38)
            d_44_v34_: _dafny.Set
            d_44_v34_ = _dafny.Set({d_13_v7_})
            d_45_v35_: _dafny.Map
            d_45_v35_ = _dafny.Map({d_9_v3_: (d_44_v34_).intersection(d_44_v34_)})
            d_46_v36_: _dafny.Seq
            d_46_v36_ = _dafny.SeqWithoutIsStrInference([d_44_v34_, d_44_v34_, d_44_v34_, d_44_v34_, d_44_v34_])
            d_45_v35_ = (d_45_v35_).set(d_9_v3_, (d_46_v36_)[default__.safeIndex(d_7_v1_, len(d_46_v36_))])
        index4_ = default__.safeIndex(314, (d_5_v0_).length(0))
        (d_5_v0_)[index4_] = p0
        d_47_v37_: _dafny.Seq
        d_47_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mexshnevu"))
        d_48_v38_: _dafny.MultiSet
        d_48_v38_ = _dafny.MultiSet([default__.fm1(p1, p0, len(_dafny.SeqWithoutIsStrInference([d_8_v2_])), globalState), p0])
        d_49_v39_: _dafny.Seq
        d_49_v39_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_50_v40_: _dafny.MultiSet
        d_50_v40_ = _dafny.MultiSet([d_9_v3_])
        index5_ = default__.safeIndex(314, (d_5_v0_).length(0))
        rhs4_ = p0
        rhs5_ = ((d_48_v38_).issubset(d_48_v38_)) not in (d_49_v39_)
        rhs6_ = (d_50_v40_).ispropersubset(d_50_v40_)
        rhs7_ = d_47_v37_
        lhs2_ = d_5_v0_
        lhs3_ = default__.safeIndex(314, (d_5_v0_).length(0))
        r2 = rhs4_
        r2 = rhs5_
        lhs2_[lhs3_] = rhs6_
        d_47_v37_ = rhs7_
        d_51_v41_: C0
        nw8_ = C0()
        nw8_.ctor__((d_8_v2_) + ((d_9_v3_).f0))
        d_51_v41_ = nw8_
        d_52_v42_: _dafny.Array
        def lambda0_(d_53_v0_):
            def lambda1_(d_54_i3_):
                return (_dafny.Map({(d_53_v0_)[default__.safeIndex(314, (d_53_v0_).length(0))]: D0_DC1()}) if False else _dafny.Map({(d_53_v0_)[default__.safeIndex(314, (d_53_v0_).length(0))]: D0_DC1()}))

            return lambda1_

        init0_ = lambda0_(d_5_v0_)
        nw9_ = _dafny.Array(None, 9)
        for i0_0_ in range(nw9_.length(0)):
            nw9_[i0_0_] = init0_(i0_0_)
        d_52_v42_ = nw9_
        d_55_v43_: D0
        d_55_v43_ = D0_DC1()
        d_56_v44_: _dafny.Map
        d_56_v44_ = _dafny.Map({p0: d_55_v43_})
        index6_ = default__.safeIndex(646, (d_52_v42_).length(0))
        (d_52_v42_)[index6_] = d_56_v44_
        pat_let_tv0_ = d_56_v44_
        pat_let_tv1_ = d_56_v44_
        index7_ = default__.safeIndex(646, (d_52_v42_).length(0))
        def lambda2_(source3_):
            if source3_.is_DC1:
                return pat_let_tv0_
            elif True:
                d_57___mcc_h1_ = source3_.cf0
                d_58_cf0_ = d_57___mcc_h1_
                return pat_let_tv1_

        (d_52_v42_)[index7_] = lambda2_(default__.fm7((d_5_v0_)[default__.safeIndex(314, (d_5_v0_).length(0))], globalState))
        d_59_v45_: _dafny.Array
        nw10_ = _dafny.Array(_dafny.Array(None, 0), 7)
        d_59_v45_ = nw10_
        d_60_v46_: _dafny.Array
        def lambda3_(d_61_v1_):
            def lambda4_(d_62_i4_):
                return _dafny.Map({len(_dafny.Set({d_61_v1_, (_dafny.MultiSet([d_61_v1_])).cardinality})): d_61_v1_})

            return lambda4_

        init1_ = lambda3_(d_7_v1_)
        nw11_ = _dafny.Array(None, 26)
        for i0_1_ in range(nw11_.length(0)):
            nw11_[i0_1_] = init1_(i0_1_)
        d_60_v46_ = nw11_
        index8_ = default__.safeIndex(152, (d_59_v45_).length(0))
        (d_59_v45_)[index8_] = d_60_v46_
        index9_ = default__.safeIndex(152, (d_59_v45_).length(0))
        (d_59_v45_)[index9_] = d_60_v46_
        r0 = (_dafny.SeqWithoutIsStrInference([139, d_7_v1_, d_7_v1_])) + (d_8_v2_)
        r1 = d_7_v1_
        r2 = p0
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_63_globalState_: GlobalState
        nw12_ = GlobalState()
        nw12_.ctor__()
        d_63_globalState_ = nw12_
        d_64_v0_: _dafny.Array
        nw13_ = _dafny.Array(False, 7)
        d_64_v0_ = nw13_
        d_65_v1_: _dafny.Set
        d_65_v1_ = _dafny.Set({d_64_v0_})
        d_65_v1_ = d_65_v1_
        d_66_v2_: int
        d_66_v2_ = 98
        d_67_v3_: bool
        d_67_v3_ = True
        d_68_v4_: _dafny.Map
        d_68_v4_ = _dafny.Map({len(_dafny.Set({d_67_v3_, d_67_v3_, d_67_v3_})): d_66_v2_})
        d_69_v6_: D0
        d_69_v6_ = D0_DC0(False)
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(-725, -518):
                d_70_v5_: int = compr_0_
                if ((-725) <= (d_70_v5_)) and ((d_70_v5_) < (-518)):
                    coll0_[default__.safeModuloInt(d_70_v5_, d_66_v2_)] = (0) - ((0) - (d_66_v2_))
            return _dafny.Map(coll0_)
        rhs8_ = d_66_v2_
        rhs9_ = (iife0_()
        ) | (d_68_v4_)
        rhs10_ = (d_67_v3_) == ((d_69_v6_).cf0)
        d_66_v2_ = rhs8_
        d_68_v4_ = rhs9_
        d_67_v3_ = rhs10_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_64_v0_).length(0)):
            d_71_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_71_i0_)) and ((d_71_i0_) < ((d_64_v0_).length(0)))):
                (d_64_v0_)[(d_71_i0_)] = d_67_v3_
        d_72_v7_: str
        d_72_v7_ = _dafny.CodePoint('y')
        d_73_v8_: _dafny.Map
        d_73_v8_ = _dafny.Map({d_72_v7_: d_67_v3_})
        hi0_ = d_66_v2_
        for d_74_i1_ in range(len(d_73_v8_), hi0_):
            d_75_v9_: D0
            d_75_v9_ = D0_DC1()
            d_75_v9_ = d_75_v9_
            d_76_v10_: _dafny.Set
            d_76_v10_ = _dafny.Set({d_67_v3_, d_67_v3_})
            d_77_v11_: _dafny.Seq
            d_78_v12_: int
            d_79_v13_: bool
            out0_: _dafny.Seq
            out1_: int
            out2_: bool
            out0_, out1_, out2_ = default__.m0(not((d_76_v10_).issubset(default__.fm0(d_63_globalState_))), not (True) or (True), d_63_globalState_)
            d_77_v11_ = out0_
            d_78_v12_ = out1_
            d_79_v13_ = out2_
            rhs11_ = default__.fm1(d_79_v13_, d_79_v13_, d_78_v12_, d_63_globalState_)
            rhs12_ = d_72_v7_
            rhs13_ = d_79_v13_
            d_67_v3_ = rhs11_
            d_72_v7_ = rhs12_
            d_67_v3_ = rhs13_
            d_76_v10_ = d_76_v10_
        d_80_v14_: _dafny.Seq
        d_80_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "braqfma"))
        d_81_v15_: D0
        d_81_v15_ = D0_DC1()
        d_82_v16_: _dafny.Array
        nw14_ = _dafny.Array(None, 12)
        nw14_[int(0)] = d_81_v15_
        nw14_[int(1)] = d_81_v15_
        nw14_[int(2)] = d_81_v15_
        nw14_[int(3)] = d_81_v15_
        nw14_[int(4)] = d_81_v15_
        nw14_[int(5)] = default__.fm2(d_66_v2_, d_63_globalState_)
        nw14_[int(6)] = d_81_v15_
        nw14_[int(7)] = D0_DC1()
        nw14_[int(8)] = d_81_v15_
        nw14_[int(9)] = d_81_v15_
        nw14_[int(10)] = d_81_v15_
        nw14_[int(11)] = d_81_v15_
        d_82_v16_ = nw14_
        d_83_v17_: _dafny.Map
        d_83_v17_ = _dafny.Map({d_82_v16_: d_67_v3_})
        rhs14_ = ((d_80_v14_) + (d_80_v14_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvtinys")))
        rhs15_ = d_67_v3_
        rhs16_ = d_67_v3_
        rhs17_ = (d_83_v17_) | (d_83_v17_)
        rhs18_ = d_69_v6_
        d_80_v14_ = rhs14_
        d_67_v3_ = rhs15_
        d_67_v3_ = rhs16_
        d_83_v17_ = rhs17_
        d_69_v6_ = rhs18_
        d_84_v18_: _dafny.MultiSet
        d_84_v18_ = _dafny.MultiSet([d_80_v14_, d_80_v14_, d_80_v14_])
        rhs19_ = True
        rhs20_ = (d_66_v2_) + (((d_84_v18_).cardinality) + (d_66_v2_))
        d_67_v3_ = rhs19_
        d_66_v2_ = rhs20_
        d_85_i2_: int
        d_85_i2_ = 0
        with _dafny.label("0"):
            while (d_67_v3_) == (not(d_67_v3_)):
                with _dafny.c_label("0"):
                    if (d_85_i2_) >= (100):
                        raise _dafny.Break("0")
                    d_85_i2_ = (d_85_i2_) + (1)
                    d_66_v2_ = (175) + (d_66_v2_)
                    d_86_v19_: _dafny.Seq
                    d_87_v20_: int
                    d_88_v21_: bool
                    out3_: _dafny.Seq
                    out4_: int
                    out5_: bool
                    out3_, out4_, out5_ = default__.m0(False, d_67_v3_, d_63_globalState_)
                    d_86_v19_ = out3_
                    d_87_v20_ = out4_
                    d_88_v21_ = out5_
                    d_89_v22_: _dafny.MultiSet
                    d_89_v22_ = _dafny.MultiSet([d_88_v21_])
                    d_88_v21_ = (d_66_v2_) != (default__.safeModuloInt(d_87_v20_, (d_89_v22_).cardinality))
                    d_88_v21_ = d_88_v21_
                    pass
            pass
        d_80_v14_ = (default__.fm3(d_63_globalState_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttjivi")))
        if not(d_67_v3_):
            d_67_v3_ = d_67_v3_
            d_90_v23_: _dafny.Map
            d_90_v23_ = _dafny.Map({d_66_v2_: d_67_v3_})
            d_91_v24_: _dafny.Map
            d_91_v24_ = _dafny.Map({d_67_v3_: d_66_v2_})
            d_92_v25_: _dafny.Seq
            d_92_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_67_v3_: len(d_90_v23_)}), d_91_v24_])
            d_93_v26_: _dafny.Map
            d_93_v26_ = _dafny.Map({False: (default__.fm4(d_66_v2_, len(d_92_v25_), d_67_v3_, d_67_v3_, d_63_globalState_)) - (-707)})
            d_91_v24_ = (d_91_v24_).set(d_67_v3_, (d_66_v2_) - (d_66_v2_))
            d_94_v27_: _dafny.Seq
            d_94_v27_ = _dafny.SeqWithoutIsStrInference([d_67_v3_])
            d_67_v3_ = (d_94_v27_)[default__.safeIndex(d_66_v2_, len(d_94_v27_))]
            d_66_v2_ = 640
            d_66_v2_ = d_66_v2_
        elif True:
            if d_67_v3_:
                d_95_v28_: _dafny.Array
                nw15_ = _dafny.Array(int(0), 19)
                d_95_v28_ = nw15_
                index10_ = default__.safeIndex(611, (d_95_v28_).length(0))
                (d_95_v28_)[index10_] = d_66_v2_
                d_96_v29_: _dafny.MultiSet
                d_96_v29_ = _dafny.MultiSet([d_72_v7_, _dafny.CodePoint('k'), d_72_v7_])
                d_97_v30_: _dafny.Map
                d_97_v30_ = _dafny.Map({d_67_v3_: -34})
                d_98_v31_: _dafny.MultiSet
                d_98_v31_ = _dafny.MultiSet([d_67_v3_, d_67_v3_, d_67_v3_])
                d_99_v32_: _dafny.MultiSet
                d_99_v32_ = _dafny.MultiSet([((d_96_v29_)[d_72_v7_] if (d_72_v7_) in (d_96_v29_) else default__.fm4(d_66_v2_, d_66_v2_, False, d_67_v3_, d_63_globalState_)), (_dafny.MultiSet([d_67_v3_])).cardinality, d_66_v2_, (len((d_97_v30_).set(default__.fm1(d_67_v3_, d_67_v3_, d_66_v2_, d_63_globalState_), d_66_v2_))) + ((d_98_v31_).cardinality), d_66_v2_])
                index11_ = default__.safeIndex(611, (d_95_v28_).length(0))
                (d_95_v28_)[index11_] = ((d_99_v32_)[d_66_v2_] if (d_66_v2_) in (d_99_v32_) else default__.safeDivisionInt(850, (0) - (d_66_v2_)))
                d_100_v33_: _dafny.Seq
                d_100_v33_ = _dafny.SeqWithoutIsStrInference([False])
                d_101_v34_: C0
                nw16_ = C0()
                nw16_.ctor__(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_100_v33_)).cardinality]))
                d_101_v34_ = nw16_
                d_80_v14_ = d_80_v14_
                d_72_v7_ = (d_80_v14_)[default__.safeIndex(d_66_v2_, len(d_80_v14_))]
                d_102_v35_: _dafny.Array
                nw17_ = _dafny.Array(None, 9)
                d_102_v35_ = nw17_
                d_103_v36_: _dafny.Map
                d_103_v36_ = _dafny.Map({d_102_v35_: d_67_v3_})
                d_104_v37_: _dafny.Seq
                d_105_v38_: int
                d_106_v39_: bool
                out6_: _dafny.Seq
                out7_: int
                out8_: bool
                out6_, out7_, out8_ = default__.m0(d_67_v3_, not(((0) - (len(d_103_v36_))) < ((d_101_v34_).fm5(d_67_v3_, d_67_v3_, d_67_v3_, (d_95_v28_)[default__.safeIndex(611, (d_95_v28_).length(0))], d_63_globalState_))), d_63_globalState_)
                d_104_v37_ = out6_
                d_105_v38_ = out7_
                d_106_v39_ = out8_
            elif True:
                d_107_v40_: _dafny.Set
                d_107_v40_ = _dafny.Set({not(d_67_v3_)})
                d_108_v41_: _dafny.Map
                d_108_v41_ = _dafny.Map({d_69_v6_: d_107_v40_})
                d_67_v3_ = (d_67_v3_ if d_67_v3_ else (d_108_v41_) == (d_108_v41_))
                d_66_v2_ = (d_66_v2_) + ((0) - (d_66_v2_))
                d_66_v2_ = d_66_v2_
                d_109_v42_: _dafny.Map
                d_109_v42_ = _dafny.Map({d_67_v3_: d_67_v3_})
                d_109_v42_ = (d_109_v42_) | (d_109_v42_)
                d_110_v43_: _dafny.Seq
                d_110_v43_ = _dafny.SeqWithoutIsStrInference([True, d_67_v3_, False, d_67_v3_, d_67_v3_])
                d_111_v44_: _dafny.MultiSet
                d_111_v44_ = _dafny.MultiSet([D0_DC1()])
                d_112_v45_: _dafny.Map
                d_112_v45_ = _dafny.Map({(d_110_v43_)[default__.safeIndex(d_66_v2_, len(d_110_v43_))]: ((d_111_v44_)[D0_DC1()] if (D0_DC1()) in (d_111_v44_) else d_66_v2_)})
                d_66_v2_ = (default__.safeModuloInt(len(d_112_v45_), d_66_v2_)) + (d_66_v2_)
            d_113_v46_: _dafny.Map
            d_113_v46_ = _dafny.Map({d_67_v3_: d_67_v3_})
            d_113_v46_ = (d_113_v46_).set(d_67_v3_, (d_69_v6_).cf0)
            d_114_v47_: _dafny.MultiSet
            d_114_v47_ = _dafny.MultiSet([d_67_v3_, d_67_v3_, d_67_v3_])
            d_115_v48_: C0
            nw18_ = C0()
            nw18_.ctor__(_dafny.SeqWithoutIsStrInference([len(d_80_v14_), 706, (d_114_v47_).cardinality]))
            d_115_v48_ = nw18_
            d_116_v49_: _dafny.Map
            d_116_v49_ = _dafny.Map({False: d_115_v48_})
            d_117_v50_: D1
            d_117_v50_ = D1_DC2(((d_116_v49_)[d_67_v3_] if (d_67_v3_) in (d_116_v49_) else d_115_v48_))
            d_118_v51_: _dafny.Array
            nw19_ = _dafny.Array(None, 16)
            nw19_[int(0)] = d_115_v48_
            nw19_[int(1)] = d_115_v48_
            nw19_[int(2)] = d_115_v48_
            nw19_[int(3)] = d_115_v48_
            nw19_[int(4)] = (d_115_v48_ if True else d_115_v48_)
            nw19_[int(5)] = d_115_v48_
            nw19_[int(6)] = d_115_v48_
            nw19_[int(7)] = d_115_v48_
            nw19_[int(8)] = d_115_v48_
            nw19_[int(9)] = d_115_v48_
            nw19_[int(10)] = d_115_v48_
            nw19_[int(11)] = d_115_v48_
            nw19_[int(12)] = (d_117_v50_).cf1
            nw19_[int(13)] = d_115_v48_
            nw19_[int(14)] = d_115_v48_
            nw19_[int(15)] = d_115_v48_
            d_118_v51_ = nw19_
            index12_ = default__.safeIndex(539, (d_118_v51_).length(0))
            (d_118_v51_)[index12_] = d_115_v48_
            index13_ = default__.safeIndex(539, (d_118_v51_).length(0))
            (d_118_v51_)[index13_] = d_115_v48_
            d_67_v3_ = d_67_v3_
            index14_ = default__.safeIndex(866, (d_64_v0_).length(0))
            (d_64_v0_)[index14_] = (d_67_v3_) == (d_67_v3_)
            index15_ = default__.safeIndex(866, (d_64_v0_).length(0))
            (d_64_v0_)[index15_] = d_67_v3_
        if (d_69_v6_).cf0:
            d_119_v52_: _dafny.Seq
            d_119_v52_ = _dafny.SeqWithoutIsStrInference([d_67_v3_])
            d_120_v53_: _dafny.MultiSet
            d_120_v53_ = _dafny.MultiSet([d_67_v3_, d_67_v3_])
            d_66_v2_ = (d_66_v2_ if (d_120_v53_).ispropersubset(_dafny.MultiSet(d_119_v52_)) else ((0) - (len(d_119_v52_))) * (d_66_v2_))
            d_121_v54_: _dafny.Set
            d_121_v54_ = _dafny.Set({True})
            d_121_v54_ = ((_dafny.Set({False, d_67_v3_})) - (_dafny.Set({not(d_67_v3_)})) if d_67_v3_ else _dafny.Set({d_67_v3_, d_67_v3_, d_67_v3_}))
            d_122_v55_: _dafny.Array
            def lambda5_(d_123_v2_):
                def lambda6_(d_124_i3_):
                    return (d_124_i3_) - (d_123_v2_)

                return lambda6_

            init2_ = lambda5_(d_66_v2_)
            nw20_ = _dafny.Array(None, 25)
            for i0_2_ in range(nw20_.length(0)):
                nw20_[i0_2_] = init2_(i0_2_)
            d_122_v55_ = nw20_
            index16_ = default__.safeIndex(637, (d_122_v55_).length(0))
            (d_122_v55_)[index16_] = d_66_v2_
            index17_ = default__.safeIndex(597, (d_122_v55_).length(0))
            (d_122_v55_)[index17_] = -266
            d_125_v56_: _dafny.Seq
            d_125_v56_ = _dafny.SeqWithoutIsStrInference([d_66_v2_])
            d_126_v57_: _dafny.Seq
            d_126_v57_ = _dafny.SeqWithoutIsStrInference([len(d_125_v56_), len(d_125_v56_)])
            d_127_v58_: C0
            nw21_ = C0()
            nw21_.ctor__(d_126_v57_)
            d_127_v58_ = nw21_
            d_128_v59_: _dafny.Map
            d_128_v59_ = _dafny.Map({d_127_v58_: True})
            d_129_v60_: _dafny.Map
            d_129_v60_ = _dafny.Map({len(d_80_v14_): d_128_v59_})
            index18_ = default__.safeIndex(637, (d_122_v55_).length(0))
            index19_ = default__.safeIndex(597, (d_122_v55_).length(0))
            rhs21_ = (0) - (d_66_v2_)
            rhs22_ = (d_69_v6_).cf0
            rhs23_ = d_66_v2_
            rhs24_ = 246
            rhs25_ = ((0) - (default__.safeDivisionInt(len(d_80_v14_), (D1_DC4(d_129_v60_, d_66_v2_)).cf5))) - ((d_127_v58_).fm5(False, d_67_v3_, d_67_v3_, default__.fm4(d_66_v2_, d_66_v2_, d_67_v3_, d_67_v3_, d_63_globalState_), d_63_globalState_))
            lhs4_ = d_122_v55_
            lhs5_ = default__.safeIndex(637, (d_122_v55_).length(0))
            lhs6_ = d_122_v55_
            lhs7_ = default__.safeIndex(597, (d_122_v55_).length(0))
            d_66_v2_ = rhs21_
            d_67_v3_ = rhs22_
            lhs4_[lhs5_] = rhs23_
            d_66_v2_ = rhs24_
            lhs6_[lhs7_] = rhs25_
            d_67_v3_ = (d_120_v53_).issubset(_dafny.MultiSet(d_119_v52_))
            d_67_v3_ = True
        elif True:
            d_68_v4_ = (d_68_v4_).set(d_66_v2_, (0) - (d_66_v2_))
            d_66_v2_ = 935
            d_66_v2_ = (0) - (d_66_v2_)
            d_67_v3_ = not(d_67_v3_)
            d_130_v61_: _dafny.Map
            d_130_v61_ = _dafny.Map({_dafny.MultiSet([d_67_v3_]): d_66_v2_})
            d_131_v62_: _dafny.MultiSet
            d_131_v62_ = _dafny.MultiSet([d_67_v3_])
            d_67_v3_ = (((d_130_v61_)[d_131_v62_] if (d_131_v62_) in (d_130_v61_) else d_66_v2_)) <= (d_66_v2_)
        d_132_i4_: int
        d_132_i4_ = 0
        with _dafny.label("1"):
            while d_67_v3_:
                with _dafny.c_label("1"):
                    if (d_132_i4_) >= (100):
                        raise _dafny.Break("1")
                    d_132_i4_ = (d_132_i4_) + (1)
                    d_66_v2_ = d_66_v2_
                    d_133_v63_: _dafny.Map
                    d_133_v63_ = _dafny.Map({(0) - ((0) - (d_66_v2_)): d_67_v3_})
                    d_134_v64_: C0
                    nw22_ = C0()
                    nw22_.ctor__((default__.fm6(d_66_v2_, d_67_v3_, ((d_133_v63_)[834] if (834) in (d_133_v63_) else d_67_v3_), d_63_globalState_)))
                    d_134_v64_ = nw22_
                    d_80_v14_ = d_80_v14_
                    d_66_v2_ = (d_66_v2_ if d_67_v3_ else d_66_v2_)
                    pass
            pass
        d_67_v3_ = d_67_v3_
        d_67_v3_ = (default__.fm4(d_66_v2_, (0) - (d_66_v2_), False, True, d_63_globalState_)) <= (d_66_v2_)
        d_135_v65_: _dafny.Set
        d_135_v65_ = _dafny.Set({len(d_80_v14_)})
        d_136_v66_: _dafny.Seq
        d_136_v66_ = _dafny.SeqWithoutIsStrInference([d_66_v2_])
        d_137_v67_: C0
        nw23_ = C0()
        nw23_.ctor__(d_136_v66_)
        d_137_v67_ = nw23_
        d_138_v68_: _dafny.Map
        d_138_v68_ = _dafny.Map({d_137_v67_: d_67_v3_})
        d_139_v69_: _dafny.Map
        d_139_v69_ = _dafny.Map({232: d_138_v68_})
        d_140_v70_: D1
        d_140_v70_ = D1_DC4(d_139_v69_, (_dafny.MultiSet([not(d_67_v3_), d_67_v3_])).cardinality)
        d_141_v71_: _dafny.Map
        d_141_v71_ = _dafny.Map({d_135_v65_: d_140_v70_})
        d_142_v72_: _dafny.Seq
        d_142_v72_ = _dafny.SeqWithoutIsStrInference([len(d_141_v71_), (0) - (-911), len(d_80_v14_)])
        d_143_v73_: C0
        nw24_ = C0()
        nw24_.ctor__(d_136_v66_)
        d_143_v73_ = nw24_
        d_144_v74_: _dafny.Seq
        d_145_v75_: int
        d_146_v76_: bool
        out9_: _dafny.Seq
        out10_: int
        out11_: bool
        out9_, out10_, out11_ = default__.m0(d_67_v3_, (d_69_v6_).cf0, d_63_globalState_)
        d_144_v74_ = out9_
        d_145_v75_ = out10_
        d_146_v76_ = out11_
        d_145_v75_ = (-550 if d_67_v3_ else default__.safeModuloInt((0) - (d_66_v2_), (0) - (d_145_v75_)))
        _dafny.print(_dafny.string_of((d_64_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_64_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_64_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_64_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_64_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_64_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_64_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_65_v1_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_66_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_67_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_68_v4_) == (_dafny.Map({59: 98, 60: 98, 61: 98, 62: 98, 63: 98, 64: 98, 65: 98, 66: 98, 67: 98, 68: 98, 69: 98, 70: 98, 71: 98, 72: 98, 73: 98, 74: 98, 75: 98, 76: 98, 77: 98, 78: 98, 79: 98, 80: 98, 81: 98, 82: 98, 83: 98, 84: 98, 85: 98, 86: 98, 87: 98, 88: 98, 89: 98, 90: 98, 91: 98, 92: 98, 93: 98, 94: 98, 95: 98, 96: 98, 97: 98, 0: 98, 1: 98, 2: 98, 3: 98, 4: 98, 5: 98, 6: 98, 7: 98, 8: 98, 9: 98, 10: 98, 11: 98, 12: 98, 13: 98, 14: 98, 15: 98, 16: 98, 17: 98, 18: 98, 19: 98, 20: 98, 21: 98, 22: 98, 23: 98, 24: 98, 25: 98, 26: 98, 27: 98, 28: 98, 29: 98, 30: 98, 31: 98, 32: 98, 33: 98, 34: 98, 35: 98, 36: 98, 37: 98, 38: 98, 39: 98, 40: 98, 41: 98, 42: 98, 43: 98, 44: 98, 45: 98, 46: 98, 47: 98, 48: 98, 49: 98, 50: 98, 51: 98, 52: 98, 53: 98, 54: 98, 55: 98, 56: 98, 57: 98, 58: 98, 199: -199}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v6_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_72_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_73_v8_) == (_dafny.Map({_dafny.CodePoint('y'): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_80_v14_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_83_v17_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_v18_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "braqfmabraqfmadvtinys")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "braqfmabraqfmadvtinys")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "braqfmabraqfmadvtinys"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_85_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_132_i4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v65_) == (_dafny.Set({7}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v66_) == (_dafny.SeqWithoutIsStrInference([-935]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_137_v67_).f0) == (_dafny.SeqWithoutIsStrInference([-935]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_138_v68_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_139_v69_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_140_v70_).cf4)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_140_v70_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_141_v71_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_v72_) == (_dafny.SeqWithoutIsStrInference([1, 911, 7]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_v73_).f0) == (_dafny.SeqWithoutIsStrInference([-935]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_144_v74_) == (_dafny.SeqWithoutIsStrInference([139, 14, 14, 14]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_145_v75_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_146_v76_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC3(_dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC3(D1, NamedTuple('DC3', [('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC5(D2, NamedTuple('DC5', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC7(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)

class D3_DC7(D3, NamedTuple('DC7', [('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC6(D3, NamedTuple('DC6', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)

class D4_DC8(D4, NamedTuple('DC8', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self):
        pass
        pass


class C0:
    def  __init__(self):
        self._f0: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f0):
        (self)._f0 = f0

    def fm5(self, p0, p1, p2, p3, globalState):
        return -344

    @property
    def f0(self):
        return self._f0
