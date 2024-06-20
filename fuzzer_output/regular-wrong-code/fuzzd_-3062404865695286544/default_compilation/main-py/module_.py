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
    def fm1(globalState):
        return (((_dafny.MultiSet([not(True), True])).cardinality) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))))) < ((0) - ((0) - ((0) - (-257))))

    @staticmethod
    def fm2(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(98, 163):
                d_0_v0_: int = compr_0_
                if ((98) <= (d_0_v0_)) and ((d_0_v0_) < (163)):
                    coll0_[default__.safeDivisionInt(d_0_v0_, 131)] = False
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.Map({122: 265})).keys.Elements:
                d_2_v1_: int = compr_1_
                if (d_2_v1_) in (_dafny.Map({122: 265})):
                    coll1_[(d_2_v1_) - (798)] = True
            return _dafny.Map(coll1_)
        return ((len(iife0_()
        )) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({-722: True}), _dafny.Map({393: False}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1_i0_ in range(default__.abs(-391))])): False}), iife1_()
        , _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_3_i1_ in range(default__.abs(-223))])): True})])))) - (117)

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): True}))])).cardinality])).Elements:
                d_4_v0_: int = compr_2_
                if (d_4_v0_) in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): True}))])).cardinality])):
                    coll2_[(d_4_v0_) * (802)] = -596
            return _dafny.Map(coll2_)
        return (iife2_()
        ) | (_dafny.Map({567: len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True]), _dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([True])]))}))

    @staticmethod
    def fm6(globalState):
        if (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_5_i1_ in range(default__.abs(837))])), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(False): 47})), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okfva"))))]))).cardinality, -567, 792, -57])).issubset(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([985 for d_6_i0_ in range(default__.abs(-769))])), -713])):
            return _dafny.CodePoint('l')
        elif True:
            return _dafny.CodePoint('k')

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        return D1_DC4(not(False))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: D2
            for compr_3_ in (_dafny.Map({D2_DC9(_dafny.CodePoint('k')): False})).keys.Elements:
                d_7_v0_: D2 = compr_3_
                if (d_7_v0_) in (_dafny.Map({D2_DC9(_dafny.CodePoint('k')): False})):
                    coll3_[d_7_v0_] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ev"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbujxi"))))
            return _dafny.Map(coll3_)
        return iife3_()
        

    @staticmethod
    def fm9(p0, p1, globalState):
        if True:
            return (_dafny.MultiSet([_dafny.Map({True: True})])) | (_dafny.MultiSet([_dafny.Map({True: True}), _dafny.Map({False: (D2_DC10(False, not(False))).cf20})]))
        elif True:
            return (_dafny.MultiSet([_dafny.Map({True: False})])) | (_dafny.MultiSet([_dafny.Map({True: True}), _dafny.Map({False: False})]))

    @staticmethod
    def fm10(p0, p1, globalState):
        return _dafny.MultiSet([not((-677) > (772)), (_dafny.Set({_dafny.MultiSet([False, not(False)])})).ispropersubset((D6_DC21(_dafny.Set({_dafny.MultiSet([False, not(True)])}))).cf42)])

    @staticmethod
    def Main(noArgsParameter__):
        d_8_globalState_: GlobalState
        nw0_ = GlobalState()
        nw0_.ctor__(470, -748, True, 220, 551, False)
        d_8_globalState_ = nw0_
        d_9_v0_: str
        d_9_v0_ = _dafny.CodePoint('e')
        d_10_v1_: C1
        nw1_ = C1()
        nw1_.ctor__(d_9_v0_, False)
        d_10_v1_ = nw1_
        (d_8_globalState_).f1 = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_10_v1_).f7]) for d_11_i0_ in range(default__.abs(-275))]))).cardinality
        d_12_v2_: D1
        d_12_v2_ = D1_DC6(_dafny.CodePoint('o'), default__.fm6(d_8_globalState_))
        d_13_v3_: _dafny.Seq
        d_13_v3_ = _dafny.SeqWithoutIsStrInference([D1_DC6(d_9_v0_, (d_10_v1_).f6), (d_12_v2_ if False else d_12_v2_)])
        d_13_v3_ = _dafny.SeqWithoutIsStrInference([d_12_v2_, d_12_v2_, d_12_v2_, d_12_v2_, d_12_v2_])
        d_14_v4_: int
        d_14_v4_ = 22
        d_15_v5_: _dafny.Map
        d_15_v5_ = _dafny.Map({d_14_v4_: d_14_v4_})
        d_16_v6_: _dafny.Map
        d_16_v6_ = _dafny.Map({not((d_10_v1_).f7): True})
        d_17_v7_: _dafny.Array
        def lambda0_(d_18_v1_):
            def lambda1_(d_19_i1_):
                return (d_18_v1_).f7

            return lambda1_

        init0_ = lambda0_(d_10_v1_)
        nw2_ = _dafny.Array(None, 16)
        for i0_0_ in range(nw2_.length(0)):
            nw2_[i0_0_] = init0_(i0_0_)
        d_17_v7_ = nw2_
        d_20_v8_: D0
        d_20_v8_ = D0_DC2(d_15_v5_, d_16_v6_, d_17_v7_)
        pat_let_tv0_ = d_17_v7_
        def iife4_(_pat_let0_0):
            def iife5_(d_21_dt__update__tmp_h0_):
                def iife6_(_pat_let1_0):
                    def iife7_(d_22_dt__update_hcf8_h0_):
                        return D0_DC2((d_21_dt__update__tmp_h0_).cf6, (d_21_dt__update__tmp_h0_).cf7, d_22_dt__update_hcf8_h0_)
                    return iife7_(_pat_let1_0)
                return iife6_(pat_let_tv0_)
            return iife5_(_pat_let0_0)
        source0_ = iife4_(d_20_v8_)
        if source0_.is_DC1:
            d_23___mcc_h0_ = source0_.cf1
            d_24___mcc_h1_ = source0_.cf2
            d_25___mcc_h2_ = source0_.cf3
            d_26___mcc_h3_ = source0_.cf4
            d_27___mcc_h4_ = source0_.cf5
            d_28_cf5_ = d_27___mcc_h4_
            d_29_cf4_ = d_26___mcc_h3_
            d_30_cf3_ = d_25___mcc_h2_
            d_31_cf2_ = d_24___mcc_h1_
            d_32_cf1_ = d_23___mcc_h0_
            nw3_ = _dafny.Array(None, 2)
            nw3_[int(0)] = (d_10_v1_).f7
            nw3_[int(1)] = (default__.fm7(758, d_14_v4_, False, d_8_globalState_)).cf10
            d_17_v7_ = nw3_
            d_33_v9_: _dafny.Map
            d_33_v9_ = _dafny.Map({False: (d_10_v1_ if (d_10_v1_).f7 else d_10_v1_)})
            d_33_v9_ = (d_33_v9_).set((d_10_v1_).f7, d_10_v1_)
            d_15_v5_ = (d_15_v5_).set(d_14_v4_, d_14_v4_)
            d_34_v10_: _dafny.MultiSet
            d_34_v10_ = _dafny.MultiSet([d_32_cf1_])
            (d_8_globalState_).f5 = ((d_34_v10_).issubset(d_34_v10_)) == (default__.fm1(d_8_globalState_))
        elif source0_.is_DC2:
            d_35___mcc_h5_ = source0_.cf6
            d_36___mcc_h6_ = source0_.cf7
            d_37___mcc_h7_ = source0_.cf8
            d_38_cf8_ = d_37___mcc_h7_
            d_39_cf7_ = d_36___mcc_h6_
            d_40_cf6_ = d_35___mcc_h5_
            d_41_v11_: _dafny.Seq
            d_41_v11_ = _dafny.SeqWithoutIsStrInference([d_14_v4_, d_14_v4_, d_14_v4_])
            d_42_v12_: _dafny.Map
            d_42_v12_ = _dafny.Map({d_41_v11_: d_14_v4_})
            d_43_v13_: int
            d_44_v14_: int
            d_45_v15_: _dafny.Array
            out0_: int
            out1_: int
            out2_: _dafny.Array
            out0_, out1_, out2_ = (d_10_v1_).m0(((d_42_v12_)[d_41_v11_] if (d_41_v11_) in (d_42_v12_) else 105), d_14_v4_, d_8_globalState_)
            d_43_v13_ = out0_
            d_44_v14_ = out1_
            d_45_v15_ = out2_
            d_46_v16_: _dafny.Seq
            d_46_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olipx"))
            d_46_v16_ = (_dafny.SeqWithoutIsStrInference([(d_10_v1_).f6 for d_47_i2_ in range(default__.abs(301))])) + (_dafny.SeqWithoutIsStrInference([(d_10_v1_).f6 for d_48_i3_ in range(default__.abs(434))]))
            (d_8_globalState_).f5 = (d_10_v1_).f7
            d_49_v17_: D0
            d_49_v17_ = D0_DC3(d_44_v14_)
            d_50_v18_: D1
            d_50_v18_ = D1_DC5((d_49_v17_).cf9, d_40_cf6_, len(d_46_v16_))
            index0_ = default__.safeIndex(445, (d_45_v15_).length(0))
            (d_45_v15_)[index0_] = (d_50_v18_).cf11
            d_51_v19_: _dafny.Set
            d_51_v19_ = _dafny.Set({(d_10_v1_).f7})
            d_52_v20_: _dafny.Array
            def lambda2_(d_53_v1_):
                def lambda3_(d_54_i4_):
                    return _dafny.SeqWithoutIsStrInference([(d_53_v1_).f7])

                return lambda3_

            init1_ = lambda2_(d_10_v1_)
            nw4_ = _dafny.Array(None, 23)
            for i0_1_ in range(nw4_.length(0)):
                nw4_[i0_1_] = init1_(i0_1_)
            d_52_v20_ = nw4_
            d_55_v21_: D0
            d_55_v21_ = D0_DC1(d_14_v4_, d_51_v19_, d_44_v14_, d_52_v20_, d_9_v0_)
            d_56_v22_: _dafny.Map
            d_56_v22_ = _dafny.Map({(d_55_v21_).cf5: not((d_10_v1_).f7)})
            d_57_v23_: _dafny.Seq
            d_57_v23_ = _dafny.SeqWithoutIsStrInference([(d_56_v22_).set((d_10_v1_).f6, (d_10_v1_).f7)])
            index1_ = default__.safeIndex(445, (d_45_v15_).length(0))
            (d_45_v15_)[index1_] = len(d_57_v23_)
        elif source0_.is_DC3:
            d_58___mcc_h8_ = source0_.cf9
            d_59_cf9_ = d_58___mcc_h8_
            d_60_v24_: _dafny.Array
            nw5_ = _dafny.Array(int(0), 14)
            d_60_v24_ = nw5_
            d_61_v25_: _dafny.Seq
            d_61_v25_ = _dafny.SeqWithoutIsStrInference([d_60_v24_])
            d_61_v25_ = (d_61_v25_) + ((d_61_v25_).set(default__.safeIndex(d_14_v4_, len(d_61_v25_)), d_60_v24_))
            (d_8_globalState_).f5 = not((d_10_v1_).f7)
            d_62_v26_: _dafny.Set
            d_62_v26_ = _dafny.Set({(D2_DC8(d_10_v1_)).cf17})
            if (d_62_v26_).isdisjoint((d_62_v26_) | (d_62_v26_)):
                d_63_v27_: _dafny.Array
                def lambda4_(d_64_v1_, d_65_v4_):
                    def lambda5_(d_66_i5_):
                        return _dafny.Map({D2_DC9((d_64_v1_).f6): d_65_v4_})

                    return lambda5_

                init2_ = lambda4_(d_10_v1_, d_14_v4_)
                nw6_ = _dafny.Array(None, 16)
                for i0_2_ in range(nw6_.length(0)):
                    nw6_[i0_2_] = init2_(i0_2_)
                d_63_v27_ = nw6_
                d_67_v28_: _dafny.Map
                d_67_v28_ = _dafny.Map({D2_DC9((d_10_v1_).f6): d_14_v4_})
                index2_ = default__.safeIndex(150, (d_63_v27_).length(0))
                (d_63_v27_)[index2_] = (d_67_v28_ if not((d_10_v1_).f7) else d_67_v28_)
                index3_ = default__.safeIndex(150, (d_63_v27_).length(0))
                (d_63_v27_)[index3_] = (d_67_v28_) | (default__.fm8(d_14_v4_, (d_10_v1_).f7, (d_10_v1_).f7, _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_10_v1_).f7]): d_59_cf9_}), d_8_globalState_))
                rhs0_ = d_10_v1_
                rhs1_ = (d_10_v1_).f7
                lhs0_ = d_8_globalState_
                d_10_v1_ = rhs0_
                lhs0_.f5 = rhs1_
                index4_ = default__.safeIndex(698, (d_17_v7_).length(0))
                (d_17_v7_)[index4_] = (d_10_v1_).f7
                d_68_v29_: _dafny.Set
                d_68_v29_ = _dafny.Set({not((d_10_v1_).f7), ((d_10_v1_).f7) or ((d_10_v1_).f7), (d_10_v1_).f7})
                d_69_v30_: _dafny.Map
                d_69_v30_ = _dafny.Map({d_15_v5_: not((d_10_v1_).f7)})
                index5_ = default__.safeIndex(698, (d_17_v7_).length(0))
                rhs2_ = (len(d_69_v30_)) >= (d_14_v4_)
                rhs3_ = d_68_v29_
                lhs1_ = d_17_v7_
                lhs2_ = default__.safeIndex(698, (d_17_v7_).length(0))
                lhs1_[lhs2_] = rhs2_
                d_68_v29_ = rhs3_
                d_70_v31_: C0
                nw7_ = C0()
                nw7_.ctor__(d_15_v5_, (d_17_v7_)[default__.safeIndex(698, (d_17_v7_).length(0))])
                d_70_v31_ = nw7_
                d_70_v31_ = d_70_v31_
                d_71_v32_: _dafny.Array
                nw8_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
                d_71_v32_ = nw8_
                d_72_v33_: _dafny.Seq
                d_72_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhphawf"))
                index6_ = default__.safeIndex(990, (d_71_v32_).length(0))
                (d_71_v32_)[index6_] = d_72_v33_
                index7_ = default__.safeIndex(990, (d_71_v32_).length(0))
                (d_71_v32_)[index7_] = (d_72_v33_) + (d_72_v33_)
            elif True:
                index8_ = default__.safeIndex(981, (d_17_v7_).length(0))
                (d_17_v7_)[index8_] = (d_10_v1_).f7
                index9_ = default__.safeIndex(981, (d_17_v7_).length(0))
                (d_17_v7_)[index9_] = ((d_59_cf9_) == (-190) if (d_10_v1_).f7 else default__.fm1(d_8_globalState_))
                d_10_v1_ = d_10_v1_
                d_73_v34_: C1
                nw9_ = C1()
                nw9_.ctor__((d_10_v1_).f6, (d_10_v1_).f7)
                d_73_v34_ = nw9_
                d_74_v35_: _dafny.Seq
                d_74_v35_ = _dafny.SeqWithoutIsStrInference([d_14_v4_, d_59_cf9_])
                (d_8_globalState_).f5 = not(not ((d_10_v1_).f7) or ((d_74_v35_) <= (d_74_v35_)))
                d_75_v36_: D2
                d_75_v36_ = D2_DC8(d_10_v1_)
                d_76_v37_: _dafny.Array
                nw10_ = _dafny.Array(None, 14)
                nw10_[int(0)] = (d_75_v36_).cf17
                nw10_[int(1)] = d_10_v1_
                nw10_[int(2)] = d_10_v1_
                nw10_[int(3)] = d_73_v34_
                nw10_[int(4)] = d_73_v34_
                nw10_[int(5)] = d_73_v34_
                nw10_[int(6)] = d_10_v1_
                nw10_[int(7)] = d_73_v34_
                nw10_[int(8)] = d_73_v34_
                nw10_[int(9)] = d_10_v1_
                nw10_[int(10)] = (d_10_v1_ if (d_10_v1_).f7 else d_10_v1_)
                nw10_[int(11)] = d_10_v1_
                nw10_[int(12)] = d_73_v34_
                nw10_[int(13)] = d_73_v34_
                d_76_v37_ = nw10_
                index10_ = default__.safeIndex(895, (d_76_v37_).length(0))
                (d_76_v37_)[index10_] = d_10_v1_
                index11_ = default__.safeIndex(895, (d_76_v37_).length(0))
                (d_76_v37_)[index11_] = d_10_v1_
            d_77_v38_: _dafny.Seq
            d_77_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsq"))
            d_14_v4_ = (d_14_v4_) - (len(d_77_v38_))
        elif True:
            d_78___mcc_h9_ = source0_.cf0
            d_79_cf0_ = d_78___mcc_h9_
            d_80_v39_: _dafny.Array
            def lambda6_(d_81_v4_):
                def lambda7_(d_82_i6_):
                    return (d_82_i6_) + (d_81_v4_)

                return lambda7_

            init3_ = lambda6_(d_14_v4_)
            nw11_ = _dafny.Array(None, 24)
            for i0_3_ in range(nw11_.length(0)):
                nw11_[i0_3_] = init3_(i0_3_)
            d_80_v39_ = nw11_
            index12_ = default__.safeIndex(624, (d_80_v39_).length(0))
            (d_80_v39_)[index12_] = d_14_v4_
            index13_ = default__.safeIndex(624, (d_80_v39_).length(0))
            rhs4_ = d_14_v4_
            rhs5_ = 280
            lhs3_ = d_80_v39_
            lhs4_ = default__.safeIndex(624, (d_80_v39_).length(0))
            lhs5_ = d_8_globalState_
            lhs3_[lhs4_] = rhs4_
            lhs5_.f1 = rhs5_
            d_83_v40_: D0
            d_83_v40_ = D0_DC0(d_17_v7_)
            d_84_v41_: _dafny.Map
            d_84_v41_ = _dafny.Map({d_83_v40_: (d_80_v39_)[default__.safeIndex(624, (d_80_v39_).length(0))]})
            d_85_v42_: _dafny.Seq
            d_85_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htil"))
            d_86_v43_: D1
            d_86_v43_ = D1_DC5(len(d_85_v42_), d_15_v5_, (d_80_v39_)[default__.safeIndex(624, (d_80_v39_).length(0))])
            index14_ = default__.safeIndex(624, (d_80_v39_).length(0))
            (d_80_v39_)[index14_] = (default__.safeModuloInt(((d_84_v41_)[d_83_v40_] if (d_83_v40_) in (d_84_v41_) else (d_80_v39_)[default__.safeIndex(624, (d_80_v39_).length(0))]), 623) if (d_10_v1_).f7 else (0) - ((d_86_v43_).cf13))
            d_17_v7_ = d_17_v7_
            d_16_v6_ = (d_16_v6_).set((d_10_v1_).f7, (d_10_v1_).f7)
        hi0_ = d_14_v4_
        for d_87_i7_ in range(d_14_v4_, hi0_):
            (d_8_globalState_).f1 = (d_87_i7_) - (default__.fm2(False, d_8_globalState_))
            d_88_v44_: _dafny.Array
            nw12_ = _dafny.Array(int(0), 22)
            d_88_v44_ = nw12_
            d_89_v45_: _dafny.Seq
            d_89_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            index15_ = default__.safeIndex(413, (d_88_v44_).length(0))
            (d_88_v44_)[index15_] = (d_14_v4_) + ((0) - (len(d_89_v45_)))
            index16_ = default__.safeIndex(413, (d_88_v44_).length(0))
            (d_88_v44_)[index16_] = d_14_v4_
            d_90_v46_: _dafny.Map
            d_90_v46_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekadg")): d_14_v4_})
            d_90_v46_ = (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpsq")): (d_88_v44_)[default__.safeIndex(413, (d_88_v44_).length(0))]})) | (d_90_v46_)
            index17_ = default__.safeIndex(132, (d_17_v7_).length(0))
            (d_17_v7_)[index17_] = not(False)
            index18_ = default__.safeIndex(132, (d_17_v7_).length(0))
            rhs6_ = (d_10_v1_).f7
            rhs7_ = (d_10_v1_).f7
            rhs8_ = (d_10_v1_).f6
            lhs6_ = d_17_v7_
            lhs7_ = default__.safeIndex(132, (d_17_v7_).length(0))
            lhs8_ = d_8_globalState_
            lhs6_[lhs7_] = rhs6_
            lhs8_.f5 = rhs7_
            d_9_v0_ = rhs8_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_17_v7_).length(0)):
            d_91_i8_: int = guard_loop_0_
            if (True) and (((0) <= (d_91_i8_)) and ((d_91_i8_) < ((d_17_v7_).length(0)))):
                (d_17_v7_)[(d_91_i8_)] = (d_10_v1_).f7
        d_92_v47_: _dafny.Seq
        d_92_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyyyndfmr"))
        d_92_v47_ = (d_92_v47_ if (d_10_v1_).f7 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vintro")))
        hi1_ = 320
        for d_93_i9_ in range(d_14_v4_, hi1_):
            d_94_v48_: int
            d_95_v49_: int
            d_96_v50_: _dafny.Array
            out3_: int
            out4_: int
            out5_: _dafny.Array
            out3_, out4_, out5_ = (d_10_v1_).m0(d_14_v4_, d_14_v4_, d_8_globalState_)
            d_94_v48_ = out3_
            d_95_v49_ = out4_
            d_96_v50_ = out5_
            d_97_v51_: D2
            d_97_v51_ = D2_DC8(d_10_v1_)
            d_98_v52_: D2
            d_98_v52_ = D2_DC12(d_97_v51_)
            source1_ = d_98_v52_
            if source1_.is_DC9:
                d_99___mcc_h10_ = source1_.cf18
                d_100_cf18_ = d_99___mcc_h10_
                index19_ = default__.safeIndex(687, (d_96_v50_).length(0))
                (d_96_v50_)[index19_] = (0) - (len(_dafny.Map({(d_10_v1_).f7: d_93_i9_})))
                index20_ = default__.safeIndex(687, (d_96_v50_).length(0))
                (d_96_v50_)[index20_] = d_93_i9_
                d_101_v53_: _dafny.Array
                nw13_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_101_v53_ = nw13_
                index21_ = default__.safeIndex(613, (d_101_v53_).length(0))
                (d_101_v53_)[index21_] = d_17_v7_
                index22_ = default__.safeIndex(613, (d_101_v53_).length(0))
                (d_101_v53_)[index22_] = d_17_v7_
                d_102_v54_: D1
                d_102_v54_ = D1_DC7((d_10_v1_).f6)
                d_103_v55_: _dafny.Map
                d_103_v55_ = _dafny.Map({d_102_v54_: (0) - (default__.safeDivisionInt((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pb")))), (0) - (d_95_v49_)))})
                d_103_v55_ = d_103_v55_
                d_15_v5_ = (d_15_v5_).set((d_14_v4_) + (d_93_i9_), 321)
            elif source1_.is_DC10:
                d_104___mcc_h11_ = source1_.cf19
                d_105___mcc_h12_ = source1_.cf20
                d_106_cf20_ = d_105___mcc_h12_
                d_107_cf19_ = d_104___mcc_h11_
                d_108_v56_: _dafny.Map
                d_108_v56_ = _dafny.Map({True: d_94_v48_})
                index23_ = default__.safeIndex(536, (d_96_v50_).length(0))
                (d_96_v50_)[index23_] = len((d_108_v56_).set((d_10_v1_).f7, d_94_v48_))
                index24_ = default__.safeIndex(536, (d_96_v50_).length(0))
                (d_96_v50_)[index24_] = default__.safeDivisionInt(len(d_92_v47_), d_95_v49_)
                index25_ = default__.safeIndex(6, (d_17_v7_).length(0))
                (d_17_v7_)[index25_] = d_107_cf19_
                index26_ = default__.safeIndex(6, (d_17_v7_).length(0))
                (d_17_v7_)[index26_] = not (True) or ((d_9_v0_) not in (d_92_v47_))
                d_109_v57_: _dafny.Map
                d_109_v57_ = _dafny.Map({(default__.fm2((d_10_v1_).f7, d_8_globalState_)) != ((d_96_v50_)[default__.safeIndex(536, (d_96_v50_).length(0))]): _dafny.Set({(d_17_v7_)[default__.safeIndex(6, (d_17_v7_).length(0))]})})
                d_110_v58_: D1
                d_110_v58_ = D1_DC4(d_107_cf19_)
                d_111_v59_: _dafny.Map
                d_111_v59_ = _dafny.Map({d_110_v58_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))})
                d_112_v60_: _dafny.Seq
                d_112_v60_ = _dafny.SeqWithoutIsStrInference([d_95_v49_])
                d_113_v61_: _dafny.Set
                d_113_v61_ = _dafny.Set({d_106_cf20_, True, default__.fm1(d_8_globalState_), d_107_cf19_, d_107_cf19_})
                rhs9_ = (((d_111_v59_)[D1_DC4(d_107_cf19_)] if (D1_DC4(d_107_cf19_)) in (d_111_v59_) else d_92_v47_)) != (d_92_v47_)
                rhs10_ = (0) - (d_95_v49_)
                rhs11_ = (d_109_v57_).set(((d_96_v50_)[default__.safeIndex(536, (d_96_v50_).length(0))]) in (d_112_v60_), (_dafny.Set({d_107_cf19_, d_106_cf20_, d_107_cf19_})).intersection(d_113_v61_))
                rhs12_ = d_10_v1_
                lhs9_ = d_8_globalState_
                d_107_cf19_ = rhs9_
                lhs9_.f1 = rhs10_
                d_109_v57_ = rhs11_
                d_10_v1_ = rhs12_
                d_114_v62_: _dafny.Map
                d_114_v62_ = _dafny.Map({d_95_v49_: _dafny.MultiSet([d_10_v1_, d_10_v1_])})
                d_115_v63_: D2
                d_115_v63_ = D2_DC8(d_10_v1_)
                d_116_v64_: _dafny.MultiSet
                d_116_v64_ = _dafny.MultiSet([d_10_v1_, d_10_v1_, d_10_v1_, (d_115_v63_).cf17, d_10_v1_])
                d_117_v65_: int
                d_118_v66_: int
                d_119_v67_: _dafny.Array
                out6_: int
                out7_: int
                out8_: _dafny.Array
                out6_, out7_, out8_ = (d_10_v1_).m0(len(d_112_v60_), (((d_114_v62_)[d_95_v49_] if (d_95_v49_) in (d_114_v62_) else d_116_v64_)).cardinality, d_8_globalState_)
                d_117_v65_ = out6_
                d_118_v66_ = out7_
                d_119_v67_ = out8_
            elif source1_.is_DC11:
                d_120___mcc_h13_ = source1_.cf21
                d_121___mcc_h14_ = source1_.cf22
                d_122_cf22_ = d_121___mcc_h14_
                d_123_cf21_ = d_120___mcc_h13_
                rhs13_ = d_14_v4_
                rhs14_ = d_96_v50_
                d_123_cf21_ = rhs13_
                d_96_v50_ = rhs14_
                (d_8_globalState_).f5 = True
                (d_8_globalState_).f5 = (d_10_v1_).f7
                d_124_v68_: _dafny.Map
                d_124_v68_ = _dafny.Map({d_95_v49_: True})
                d_124_v68_ = (d_124_v68_).set(-12, False)
            elif source1_.is_DC8:
                d_125___mcc_h15_ = source1_.cf17
                d_126_cf17_ = d_125___mcc_h15_
                d_17_v7_ = d_17_v7_
                (d_8_globalState_).f1 = d_94_v48_
                d_127_v69_: _dafny.Map
                d_127_v69_ = _dafny.Map({len(d_92_v47_): not((d_10_v1_).f7)})
                d_128_v70_: int
                d_129_v71_: int
                d_130_v72_: _dafny.Array
                out9_: int
                out10_: int
                out11_: _dafny.Array
                out9_, out10_, out11_ = (d_10_v1_).m0(d_95_v49_, default__.safeModuloInt(len(d_127_v69_), 440), d_8_globalState_)
                d_128_v70_ = out9_
                d_129_v71_ = out10_
                d_130_v72_ = out11_
                (d_8_globalState_).f5 = (d_10_v1_).f7
            elif True:
                d_131___mcc_h16_ = source1_.cf23
                d_132_cf23_ = d_131___mcc_h16_
                (d_8_globalState_).f5 = not(default__.fm1(d_8_globalState_))
                d_133_v73_: C0
                nw14_ = C0()
                nw14_.ctor__(d_15_v5_, (d_10_v1_).f7)
                d_133_v73_ = nw14_
                d_134_v74_: _dafny.MultiSet
                d_134_v74_ = _dafny.MultiSet([d_133_v73_, d_133_v73_, d_133_v73_, d_133_v73_])
                d_135_v75_: _dafny.Seq
                d_135_v75_ = _dafny.SeqWithoutIsStrInference([(d_134_v74_).cardinality])
                d_136_v76_: _dafny.Seq
                d_136_v76_ = _dafny.SeqWithoutIsStrInference([d_135_v75_])
                d_137_v77_: _dafny.Set
                d_137_v77_ = _dafny.Set({d_94_v48_, len(d_136_v76_)})
                rhs15_ = (d_10_v1_).f7
                rhs16_ = d_137_v77_
                lhs10_ = d_8_globalState_
                lhs10_.f5 = rhs15_
                d_137_v77_ = rhs16_
                d_9_v0_ = d_9_v0_
                d_138_v78_: _dafny.MultiSet
                d_138_v78_ = _dafny.MultiSet([d_17_v7_, d_17_v7_])
                d_139_v79_: _dafny.Map
                d_139_v79_ = _dafny.Map({d_95_v49_: (_dafny.MultiSet([d_17_v7_])).set(d_17_v7_, default__.abs(d_95_v49_))})
                d_140_v80_: _dafny.Array
                nw15_ = _dafny.Array(None, 14)
                nw15_[int(0)] = d_138_v78_
                nw15_[int(1)] = d_138_v78_
                nw15_[int(2)] = d_138_v78_
                nw15_[int(3)] = d_138_v78_
                nw15_[int(4)] = (d_138_v78_).set(d_17_v7_, default__.abs(d_94_v48_))
                nw15_[int(5)] = d_138_v78_
                nw15_[int(6)] = (d_138_v78_) - (d_138_v78_)
                nw15_[int(7)] = (d_138_v78_) | (d_138_v78_)
                nw15_[int(8)] = d_138_v78_
                nw15_[int(9)] = d_138_v78_
                nw15_[int(10)] = d_138_v78_
                nw15_[int(11)] = ((d_139_v79_)[d_94_v48_] if (d_94_v48_) in (d_139_v79_) else d_138_v78_)
                nw15_[int(12)] = (d_138_v78_) - (d_138_v78_)
                nw15_[int(13)] = _dafny.MultiSet([d_17_v7_])
                d_140_v80_ = nw15_
                index27_ = default__.safeIndex(52, (d_140_v80_).length(0))
                (d_140_v80_)[index27_] = ((_dafny.MultiSet([d_17_v7_])).set(d_17_v7_, default__.abs(d_93_i9_))).set(d_17_v7_, default__.abs(len(d_92_v47_)))
                d_141_v81_: _dafny.Seq
                d_141_v81_ = _dafny.SeqWithoutIsStrInference([d_17_v7_])
                d_142_v82_: _dafny.Seq
                d_142_v82_ = _dafny.SeqWithoutIsStrInference([d_138_v78_, d_138_v78_, d_138_v78_, d_138_v78_, _dafny.MultiSet([d_17_v7_, d_17_v7_, (d_141_v81_)[default__.safeIndex(d_14_v4_, len(d_141_v81_))], d_17_v7_, d_17_v7_])])
                d_143_v83_: _dafny.Map
                d_143_v83_ = _dafny.Map({(d_10_v1_).f7: d_14_v4_})
                d_144_v84_: _dafny.Map
                d_144_v84_ = _dafny.Map({d_137_v77_: d_143_v83_})
                index28_ = default__.safeIndex(52, (d_140_v80_).length(0))
                (d_140_v80_)[index28_] = ((d_142_v82_)[default__.safeIndex(len(((d_144_v84_)[d_137_v77_] if (d_137_v77_) in (d_144_v84_) else d_143_v83_)), len(d_142_v82_))]) - (d_138_v78_)
            index29_ = default__.safeIndex(859, (d_17_v7_).length(0))
            (d_17_v7_)[index29_] = (d_10_v1_).f7
            index30_ = default__.safeIndex(859, (d_17_v7_).length(0))
            (d_17_v7_)[index30_] = True
            d_145_v85_: C1
            nw16_ = C1()
            nw16_.ctor__(_dafny.CodePoint('b'), False)
            d_145_v85_ = nw16_
        d_146_i10_: int
        d_146_i10_ = 0
        with _dafny.label("0"):
            while (d_10_v1_).f7:
                with _dafny.c_label("0"):
                    if (d_146_i10_) >= (100):
                        raise _dafny.Break("0")
                    d_146_i10_ = (d_146_i10_) + (1)
                    if (d_10_v1_).f7:
                        (d_8_globalState_).f1 = (0) - (default__.safeDivisionInt(d_14_v4_, d_14_v4_))
                        d_147_v86_: D2
                        d_147_v86_ = D2_DC9((d_10_v1_).f6)
                        d_147_v86_ = d_147_v86_
                        d_148_v87_: int
                        d_149_v88_: int
                        d_150_v89_: _dafny.Array
                        out12_: int
                        out13_: int
                        out14_: _dafny.Array
                        out12_, out13_, out14_ = (d_10_v1_).m0(d_14_v4_, d_14_v4_, d_8_globalState_)
                        d_148_v87_ = out12_
                        d_149_v88_ = out13_
                        d_150_v89_ = out14_
                        d_151_v90_: _dafny.Map
                        d_151_v90_ = _dafny.Map({d_17_v7_: d_14_v4_})
                        d_151_v90_ = (d_151_v90_).set(d_17_v7_, d_14_v4_)
                        d_92_v47_ = (d_10_v1_).fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arwyvlusb")), (d_10_v1_).f6, d_8_globalState_)
                    elif True:
                        d_9_v0_ = (d_92_v47_)[default__.safeIndex(d_14_v4_, len(d_92_v47_))]
                        index31_ = default__.safeIndex(41, (d_17_v7_).length(0))
                        (d_17_v7_)[index31_] = (d_10_v1_).f7
                        index32_ = default__.safeIndex(41, (d_17_v7_).length(0))
                        rhs17_ = ((d_92_v47_) + (d_92_v47_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgpygk")))
                        rhs18_ = (d_10_v1_).f7
                        rhs19_ = (d_10_v1_).f7
                        lhs11_ = d_8_globalState_
                        lhs12_ = d_17_v7_
                        lhs13_ = default__.safeIndex(41, (d_17_v7_).length(0))
                        d_92_v47_ = rhs17_
                        lhs11_.f5 = rhs18_
                        lhs12_[lhs13_] = rhs19_
                        (d_8_globalState_).f5 = not((d_14_v4_) < (d_14_v4_))
                        d_152_v91_: _dafny.Array
                        nw17_ = _dafny.Array(None, 2)
                        nw17_[int(0)] = d_9_v0_
                        nw17_[int(1)] = d_9_v0_
                        d_152_v91_ = nw17_
                        d_153_v92_: _dafny.Seq
                        d_153_v92_ = _dafny.SeqWithoutIsStrInference([d_152_v91_])
                        d_154_v93_: D3
                        d_154_v93_ = D3_DC13(d_152_v91_)
                        d_155_v94_: _dafny.Array
                        nw18_ = _dafny.Array(None, 26)
                        nw18_[int(0)] = d_152_v91_
                        nw18_[int(1)] = d_152_v91_
                        nw18_[int(2)] = d_152_v91_
                        nw18_[int(3)] = d_152_v91_
                        nw18_[int(4)] = d_152_v91_
                        nw18_[int(5)] = d_152_v91_
                        nw18_[int(6)] = d_152_v91_
                        nw18_[int(7)] = d_152_v91_
                        nw18_[int(8)] = d_152_v91_
                        nw18_[int(9)] = (d_152_v91_ if (d_10_v1_).f7 else d_152_v91_)
                        nw18_[int(10)] = d_152_v91_
                        nw18_[int(11)] = (d_153_v92_)[default__.safeIndex(d_14_v4_, len(d_153_v92_))]
                        nw18_[int(12)] = d_152_v91_
                        nw18_[int(13)] = d_152_v91_
                        nw18_[int(14)] = d_152_v91_
                        nw18_[int(15)] = d_152_v91_
                        nw18_[int(16)] = d_152_v91_
                        nw18_[int(17)] = d_152_v91_
                        nw18_[int(18)] = d_152_v91_
                        nw18_[int(19)] = (d_154_v93_).cf24
                        nw18_[int(20)] = d_152_v91_
                        nw18_[int(21)] = d_152_v91_
                        nw18_[int(22)] = d_152_v91_
                        nw18_[int(23)] = d_152_v91_
                        nw18_[int(24)] = d_152_v91_
                        nw18_[int(25)] = d_152_v91_
                        d_155_v94_ = nw18_
                        d_155_v94_ = d_155_v94_
                        d_156_v95_: _dafny.MultiSet
                        d_156_v95_ = _dafny.MultiSet([d_14_v4_])
                        d_157_v96_: _dafny.Array
                        nw19_ = _dafny.Array(_dafny.Seq({}), 22)
                        d_157_v96_ = nw19_
                        d_158_v97_: _dafny.Array
                        def lambda8_(d_159_v1_):
                            def lambda9_(d_160_i11_):
                                return D2_DC9((d_159_v1_).f6)

                            return lambda9_

                        init4_ = lambda8_(d_10_v1_)
                        nw20_ = _dafny.Array(None, 3)
                        for i0_4_ in range(nw20_.length(0)):
                            nw20_[i0_4_] = init4_(i0_4_)
                        d_158_v97_ = nw20_
                        d_161_v98_: _dafny.Seq
                        d_161_v98_ = _dafny.SeqWithoutIsStrInference([d_158_v97_, d_158_v97_])
                        index33_ = default__.safeIndex(652, (d_157_v96_).length(0))
                        (d_157_v96_)[index33_] = d_161_v98_
                        d_162_v99_: _dafny.Set
                        d_162_v99_ = _dafny.Set({(d_17_v7_)[default__.safeIndex(41, (d_17_v7_).length(0))]})
                        index34_ = default__.safeIndex(41, (d_17_v7_).length(0))
                        index35_ = default__.safeIndex(652, (d_157_v96_).length(0))
                        rhs20_ = (d_17_v7_)[default__.safeIndex(41, (d_17_v7_).length(0))]
                        rhs21_ = d_156_v95_
                        rhs22_ = default__.safeModuloInt(d_14_v4_, d_14_v4_)
                        rhs23_ = _dafny.SeqWithoutIsStrInference([d_158_v97_])
                        rhs24_ = (_dafny.Set({(d_10_v1_).f7})).ispropersubset((d_162_v99_ if (d_10_v1_).f7 else d_162_v99_))
                        lhs14_ = d_17_v7_
                        lhs15_ = default__.safeIndex(41, (d_17_v7_).length(0))
                        lhs16_ = d_8_globalState_
                        lhs17_ = d_157_v96_
                        lhs18_ = default__.safeIndex(652, (d_157_v96_).length(0))
                        lhs19_ = d_8_globalState_
                        lhs14_[lhs15_] = rhs20_
                        d_156_v95_ = rhs21_
                        lhs16_.f1 = rhs22_
                        lhs17_[lhs18_] = rhs23_
                        lhs19_.f5 = rhs24_
                    if (d_92_v47_) == (d_92_v47_):
                        d_163_v100_: _dafny.Map
                        d_163_v100_ = _dafny.Map({d_14_v4_: (d_10_v1_).f7})
                        d_163_v100_ = (d_163_v100_).set(default__.safeDivisionInt(d_14_v4_, (_dafny.MultiSet([d_92_v47_])).cardinality), (d_10_v1_).f7)
                        index36_ = default__.safeIndex(823, (d_17_v7_).length(0))
                        (d_17_v7_)[index36_] = (d_10_v1_).f7
                        d_164_v101_: _dafny.Array
                        nw21_ = _dafny.Array(None, 7)
                        d_164_v101_ = nw21_
                        d_165_v102_: _dafny.Map
                        d_165_v102_ = _dafny.Map({d_164_v101_: (d_10_v1_).f7})
                        d_166_v103_: _dafny.MultiSet
                        d_166_v103_ = _dafny.MultiSet([d_14_v4_, d_14_v4_, d_14_v4_])
                        d_167_v104_: _dafny.Set
                        d_167_v104_ = _dafny.Set({not((d_10_v1_).f7), (d_10_v1_).f7})
                        index37_ = default__.safeIndex(823, (d_17_v7_).length(0))
                        rhs25_ = (d_10_v1_).f7
                        rhs26_ = (0) - (-735)
                        rhs27_ = len((d_165_v102_).set(d_164_v101_, (_dafny.MultiSet([d_14_v4_, d_14_v4_, len(d_167_v104_), d_14_v4_, d_14_v4_])).ispropersubset(d_166_v103_)))
                        lhs20_ = d_17_v7_
                        lhs21_ = default__.safeIndex(823, (d_17_v7_).length(0))
                        lhs22_ = d_8_globalState_
                        lhs23_ = d_8_globalState_
                        lhs20_[lhs21_] = rhs25_
                        lhs22_.f1 = rhs26_
                        lhs23_.f0 = rhs27_
                        d_168_v105_: _dafny.Array
                        nw22_ = _dafny.Array(_dafny.Array(None, 0), 7)
                        d_168_v105_ = nw22_
                        d_169_v106_: _dafny.Map
                        d_169_v106_ = _dafny.Map({d_168_v105_: (d_10_v1_).f7})
                        d_169_v106_ = (d_169_v106_).set(d_168_v105_, (d_10_v1_).f7)
                        d_170_v107_: _dafny.Map
                        d_170_v107_ = _dafny.Map({d_10_v1_: 806})
                        d_171_v108_: _dafny.Seq
                        d_171_v108_ = _dafny.SeqWithoutIsStrInference([(d_10_v1_).f7, default__.fm1(d_8_globalState_), (d_10_v1_).f7, not((d_10_v1_).f7), (d_17_v7_)[default__.safeIndex(823, (d_17_v7_).length(0))]])
                        d_172_v109_: C0
                        nw23_ = C0()
                        nw23_.ctor__((d_15_v5_).set(((d_170_v107_)[d_10_v1_] if (d_10_v1_) in (d_170_v107_) else len(d_171_v108_)), d_14_v4_), (d_17_v7_)[default__.safeIndex(823, (d_17_v7_).length(0))])
                        d_172_v109_ = nw23_
                        d_14_v4_ = d_14_v4_
                    elif True:
                        index38_ = default__.safeIndex(315, (d_17_v7_).length(0))
                        (d_17_v7_)[index38_] = default__.fm1(d_8_globalState_)
                        index39_ = default__.safeIndex(315, (d_17_v7_).length(0))
                        (d_17_v7_)[index39_] = (d_10_v1_).f7
                        d_173_v110_: _dafny.Map
                        d_173_v110_ = _dafny.Map({d_14_v4_: d_92_v47_})
                        d_174_v111_: _dafny.Map
                        d_174_v111_ = _dafny.Map({-45: d_173_v110_})
                        d_174_v111_ = (d_174_v111_).set(d_14_v4_, d_173_v110_)
                        d_175_v112_: D1
                        d_175_v112_ = D1_DC4((d_10_v1_).f7)
                        pat_let_tv1_ = d_17_v7_
                        pat_let_tv2_ = d_17_v7_
                        def iife8_(_pat_let2_0):
                            def iife9_(d_176_dt__update__tmp_h1_):
                                def iife10_(_pat_let3_0):
                                    def iife11_(d_177_dt__update_hcf10_h0_):
                                        return D1_DC4(d_177_dt__update_hcf10_h0_)
                                    return iife11_(_pat_let3_0)
                                return iife10_((pat_let_tv2_)[default__.safeIndex(315, (pat_let_tv1_).length(0))])
                            return iife9_(_pat_let2_0)
                        d_175_v112_ = iife8_(d_175_v112_)
                        (d_8_globalState_).f0 = (0) - (d_14_v4_)
                        d_178_v113_: C1
                        nw24_ = C1()
                        nw24_.ctor__(_dafny.CodePoint('j'), (d_17_v7_)[default__.safeIndex(315, (d_17_v7_).length(0))])
                        d_178_v113_ = nw24_
                    d_179_v114_: _dafny.Seq
                    d_179_v114_ = _dafny.SeqWithoutIsStrInference([default__.fm2((d_10_v1_).f7, d_8_globalState_), 555])
                    d_180_v116_: _dafny.Set
                    d_180_v116_ = _dafny.Set({d_14_v4_})
                    def iife12_():
                        coll4_ = _dafny.Set()
                        compr_4_: int
                        for compr_4_ in (d_179_v114_).Elements:
                            d_181_v115_: int = compr_4_
                            if (d_181_v115_) in (d_179_v114_):
                                coll4_ = coll4_.union(_dafny.Set([(d_181_v115_) - (-668)]))
                        return _dafny.Set(coll4_)
                    (d_8_globalState_).f5 = (d_180_v116_).ispropersubset(iife12_()
                    )
                    d_182_v117_: _dafny.Array
                    nw25_ = _dafny.Array(None, 2)
                    nw25_[int(0)] = d_92_v47_
                    nw25_[int(1)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tdaboijva")) if not((d_10_v1_).f7) else d_92_v47_)
                    d_182_v117_ = nw25_
                    index40_ = default__.safeIndex(996, (d_182_v117_).length(0))
                    (d_182_v117_)[index40_] = d_92_v47_
                    index41_ = default__.safeIndex(996, (d_182_v117_).length(0))
                    (d_182_v117_)[index41_] = (d_10_v1_).fm0((d_10_v1_).fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqains")), d_9_v0_, d_8_globalState_), (d_10_v1_).f6, d_8_globalState_)
                    pass
            pass
        d_183_v118_: _dafny.MultiSet
        d_183_v118_ = _dafny.MultiSet([d_16_v6_])
        d_184_v119_: _dafny.Array
        nw26_ = _dafny.Array(None, 11)
        nw26_[int(0)] = len(_dafny.SeqWithoutIsStrInference([(d_10_v1_).f6 for d_185_i12_ in range(default__.abs(-948))]))
        nw26_[int(1)] = d_14_v4_
        nw26_[int(2)] = d_14_v4_
        nw26_[int(3)] = default__.safeModuloInt(621, 665)
        nw26_[int(4)] = len(d_16_v6_)
        nw26_[int(5)] = default__.fm2(((d_16_v6_)[False] if (False) in (d_16_v6_) else (d_10_v1_).f7), d_8_globalState_)
        nw26_[int(6)] = d_14_v4_
        nw26_[int(7)] = d_14_v4_
        nw26_[int(8)] = d_14_v4_
        nw26_[int(9)] = d_14_v4_
        nw26_[int(10)] = ((default__.fm9((d_10_v1_).f7, (d_10_v1_).f7, d_8_globalState_)) - (d_183_v118_)).cardinality
        d_184_v119_ = nw26_
        index42_ = default__.safeIndex(229, (d_184_v119_).length(0))
        (d_184_v119_)[index42_] = d_14_v4_
        d_186_v120_: _dafny.MultiSet
        d_186_v120_ = _dafny.MultiSet([d_14_v4_, d_14_v4_])
        d_187_v121_: _dafny.Seq
        d_187_v121_ = _dafny.SeqWithoutIsStrInference([len(d_92_v47_)])
        index43_ = default__.safeIndex(229, (d_184_v119_).length(0))
        rhs28_ = d_14_v4_
        rhs29_ = ((d_186_v120_)[-694] if (-694) in (d_186_v120_) else (d_187_v121_)[default__.safeIndex(default__.fm2((d_10_v1_).f7, d_8_globalState_), len(d_187_v121_))])
        rhs30_ = d_14_v4_
        rhs31_ = not(not ((d_10_v1_).f7) or (not((d_10_v1_).f7)))
        lhs24_ = d_184_v119_
        lhs25_ = default__.safeIndex(229, (d_184_v119_).length(0))
        lhs26_ = d_8_globalState_
        lhs24_[lhs25_] = rhs28_
        d_14_v4_ = rhs29_
        d_14_v4_ = rhs30_
        lhs26_.f5 = rhs31_
        index44_ = default__.safeIndex(229, (d_184_v119_).length(0))
        (d_184_v119_)[index44_] = len(_dafny.SeqWithoutIsStrInference([d_9_v0_ for d_188_i13_ in range(default__.abs(872))]))
        d_189_v122_: _dafny.Seq
        d_189_v122_ = _dafny.SeqWithoutIsStrInference([not ((d_10_v1_).f7) or ((d_10_v1_).f7)])
        (d_8_globalState_).f5 = (d_189_v122_)[default__.safeIndex(d_14_v4_, len(d_189_v122_))]
        d_190_v123_: _dafny.Map
        d_190_v123_ = _dafny.Map({(d_10_v1_).f7: (d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))]})
        d_191_v124_: D0
        d_191_v124_ = D0_DC3(d_14_v4_)
        d_192_v125_: _dafny.Map
        d_192_v125_ = _dafny.Map({(d_189_v122_)[default__.safeIndex(((d_186_v120_).set(((d_190_v123_)[not(not((d_10_v1_).f7))] if (not(not((d_10_v1_).f7))) in (d_190_v123_) else d_14_v4_), default__.abs((d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))]))).cardinality, len(d_189_v122_))]: d_191_v124_})
        d_192_v125_ = (d_192_v125_).set((d_10_v1_).f7, d_191_v124_)
        if (len(d_92_v47_)) <= (d_14_v4_):
            d_193_v126_: _dafny.Set
            d_193_v126_ = _dafny.Set({132, d_14_v4_})
            def iife13_():
                coll5_ = _dafny.Set()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(928, 795):
                    d_194_v127_: int = compr_5_
                    if ((928) <= (d_194_v127_)) and ((d_194_v127_) < (795)):
                        coll5_ = coll5_.union(_dafny.Set([(d_194_v127_) * (len(_dafny.SeqWithoutIsStrInference([(d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))] for d_195_i14_ in range(default__.abs(347))])))]))
                return _dafny.Set(coll5_)
            d_193_v126_ = (d_193_v126_).intersection((iife13_()
            ) - (d_193_v126_))
            (d_8_globalState_).f5 = not((d_10_v1_).f7)
            d_196_v128_: _dafny.Seq
            d_196_v128_ = _dafny.SeqWithoutIsStrInference([d_15_v5_])
            d_197_v130_: _dafny.Map
            d_197_v130_ = _dafny.Map({(d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))]: (d_10_v1_).f7})
            d_198_v131_: _dafny.Seq
            d_198_v131_ = _dafny.SeqWithoutIsStrInference([d_197_v130_])
            d_199_v132_: C0
            nw27_ = C0()
            def iife14_():
                coll6_ = _dafny.Map()
                compr_6_: _dafny.Map
                for compr_6_ in (d_198_v131_).Elements:
                    d_200_v129_: _dafny.Map = compr_6_
                    if (d_200_v129_) in (d_198_v131_):
                        coll6_[d_200_v129_] = ((d_16_v6_)[True] if (True) in (d_16_v6_) else (d_10_v1_).f7)
                return _dafny.Map(coll6_)
            nw27_.ctor__(((d_196_v128_)[default__.safeIndex(len(d_187_v121_), len(d_196_v128_))]).set(d_14_v4_, len(iife14_()
            )), (d_10_v1_).f7)
            d_199_v132_ = nw27_
            d_201_v133_: D1
            d_201_v133_ = D1_DC5(default__.safeDivisionInt((d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))], d_14_v4_), (d_199_v132_.f8) | (d_15_v5_), ((d_199_v132_.f8)[d_14_v4_] if (d_14_v4_) in (d_199_v132_.f8) else d_14_v4_))
            source2_ = d_201_v133_
            if source2_.is_DC5:
                d_202___mcc_h17_ = source2_.cf11
                d_203___mcc_h18_ = source2_.cf12
                d_204___mcc_h19_ = source2_.cf13
                d_205_cf13_ = d_204___mcc_h19_
                d_206_cf12_ = d_203___mcc_h18_
                d_207_cf11_ = d_202___mcc_h17_
                d_208_v134_: int
                d_209_v135_: int
                d_210_v136_: _dafny.Array
                out15_: int
                out16_: int
                out17_: _dafny.Array
                out15_, out16_, out17_ = (d_10_v1_).m0((d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))], d_207_cf11_, d_8_globalState_)
                d_208_v134_ = out15_
                d_209_v135_ = out16_
                d_210_v136_ = out17_
                d_211_v137_: C0
                nw28_ = C0()
                nw28_.ctor__((_dafny.Map({len(d_187_v121_): len(d_92_v47_)})) | (d_206_cf12_), True)
                d_211_v137_ = nw28_
                d_212_v138_: _dafny.MultiSet
                d_212_v138_ = _dafny.MultiSet([(d_10_v1_).f6])
                d_212_v138_ = d_212_v138_
                d_9_v0_ = (d_92_v47_)[default__.safeIndex(d_207_cf11_, len(d_92_v47_))]
            elif source2_.is_DC6:
                d_213___mcc_h20_ = source2_.cf14
                d_214___mcc_h21_ = source2_.cf15
                d_215_cf15_ = d_214___mcc_h21_
                d_216_cf14_ = d_213___mcc_h20_
                rhs32_ = (d_199_v132_).f9
                rhs33_ = default__.safeDivisionInt(596, d_14_v4_)
                lhs27_ = d_8_globalState_
                lhs28_ = d_8_globalState_
                lhs27_.f5 = rhs32_
                lhs28_.f1 = rhs33_
                d_217_v139_: _dafny.Array
                nw29_ = _dafny.Array(None, 12)
                nw29_[int(0)] = d_199_v132_
                nw29_[int(1)] = d_199_v132_
                nw29_[int(2)] = d_199_v132_
                nw29_[int(3)] = d_199_v132_
                nw29_[int(4)] = d_199_v132_
                nw29_[int(5)] = d_199_v132_
                nw29_[int(6)] = (d_199_v132_ if True else d_199_v132_)
                nw29_[int(7)] = d_199_v132_
                nw29_[int(8)] = d_199_v132_
                nw29_[int(9)] = d_199_v132_
                nw29_[int(10)] = d_199_v132_
                nw29_[int(11)] = d_199_v132_
                d_217_v139_ = nw29_
                index45_ = default__.safeIndex(857, (d_217_v139_).length(0))
                (d_217_v139_)[index45_] = d_199_v132_
                d_218_v140_: _dafny.Array
                nw30_ = _dafny.Array(_dafny.CodePoint('D'), 5)
                d_218_v140_ = nw30_
                d_219_v141_: _dafny.Array
                nw31_ = _dafny.Array(None, 17)
                nw31_[int(0)] = d_218_v140_
                nw31_[int(1)] = d_218_v140_
                nw31_[int(2)] = d_218_v140_
                nw31_[int(3)] = d_218_v140_
                nw31_[int(4)] = d_218_v140_
                nw31_[int(5)] = d_218_v140_
                nw31_[int(6)] = d_218_v140_
                nw31_[int(7)] = d_218_v140_
                nw31_[int(8)] = d_218_v140_
                nw31_[int(9)] = (d_218_v140_ if (d_10_v1_).f7 else d_218_v140_)
                nw31_[int(10)] = d_218_v140_
                nw31_[int(11)] = d_218_v140_
                nw31_[int(12)] = d_218_v140_
                nw31_[int(13)] = d_218_v140_
                nw31_[int(14)] = d_218_v140_
                nw31_[int(15)] = d_218_v140_
                nw31_[int(16)] = d_218_v140_
                d_219_v141_ = nw31_
                index46_ = default__.safeIndex(857, (d_217_v139_).length(0))
                rhs34_ = d_199_v132_
                rhs35_ = (d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))]
                rhs36_ = d_219_v141_
                rhs37_ = (d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))]
                rhs38_ = not((d_10_v1_).f7)
                lhs29_ = d_217_v139_
                lhs30_ = default__.safeIndex(857, (d_217_v139_).length(0))
                lhs31_ = d_8_globalState_
                lhs32_ = d_8_globalState_
                lhs33_ = d_8_globalState_
                lhs29_[lhs30_] = rhs34_
                lhs31_.f1 = rhs35_
                d_219_v141_ = rhs36_
                lhs32_.f1 = rhs37_
                lhs33_.f5 = rhs38_
                d_215_cf15_ = (d_10_v1_).f6
                (d_8_globalState_).f5 = ((d_199_v132_).f9) or ((d_10_v1_).f7)
            elif source2_.is_DC7:
                d_220___mcc_h22_ = source2_.cf16
                d_221_cf16_ = d_220___mcc_h22_
                d_222_v142_: int
                d_223_v143_: int
                d_224_v144_: _dafny.Array
                out18_: int
                out19_: int
                out20_: _dafny.Array
                out18_, out19_, out20_ = (d_10_v1_).m0((d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))], d_14_v4_, d_8_globalState_)
                d_222_v142_ = out18_
                d_223_v143_ = out19_
                d_224_v144_ = out20_
                d_225_v145_: _dafny.Map
                d_225_v145_ = _dafny.Map({_dafny.MultiSet([(d_10_v1_).f7, (d_10_v1_).f7, (d_10_v1_).f7, (d_10_v1_).f7]): d_14_v4_})
                d_225_v145_ = (d_225_v145_).set(default__.fm10(105, _dafny.Set({(d_10_v1_).f7, False, (d_10_v1_).f7, (d_199_v132_).f9, not(not(not((d_10_v1_).f7)))}), d_8_globalState_), (d_14_v4_ if not((d_10_v1_).f7) else (d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))]))
                d_226_v146_: D2
                d_226_v146_ = D2_DC10((d_199_v132_).f9, (d_10_v1_).f7)
                d_226_v146_ = D2_DC10((d_199_v132_).f9, (d_10_v1_).f7)
                d_227_v147_: _dafny.MultiSet
                d_227_v147_ = _dafny.MultiSet([(d_10_v1_).f7, (d_199_v132_).f9])
                index47_ = default__.safeIndex(807, (d_17_v7_).length(0))
                (d_17_v7_)[index47_] = ((0) - ((d_227_v147_).cardinality)) >= (d_14_v4_)
                index48_ = default__.safeIndex(807, (d_17_v7_).length(0))
                rhs39_ = ((d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))]) != ((d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))])
                rhs40_ = (d_199_v132_).f9
                rhs41_ = (0) - (default__.safeModuloInt(d_222_v142_, (d_222_v142_) - (len(d_190_v123_))))
                lhs34_ = d_8_globalState_
                lhs35_ = d_17_v7_
                lhs36_ = default__.safeIndex(807, (d_17_v7_).length(0))
                lhs37_ = d_8_globalState_
                lhs34_.f5 = rhs39_
                lhs35_[lhs36_] = rhs40_
                lhs37_.f0 = rhs41_
            elif True:
                d_228___mcc_h23_ = source2_.cf10
                d_229_cf10_ = d_228___mcc_h23_
                index49_ = default__.safeIndex(137, (d_17_v7_).length(0))
                (d_17_v7_)[index49_] = d_229_cf10_
                index50_ = default__.safeIndex(137, (d_17_v7_).length(0))
                (d_17_v7_)[index50_] = d_229_cf10_
                d_230_v148_: D2
                d_230_v148_ = D2_DC10(True, (d_199_v132_).fm4(d_8_globalState_))
                index51_ = default__.safeIndex(137, (d_17_v7_).length(0))
                (d_17_v7_)[index51_] = (d_230_v148_).cf20
                d_9_v0_ = (d_10_v1_).f6
                (d_8_globalState_).f5 = (d_229_cf10_) == ((d_199_v132_).f9)
            d_231_v149_: _dafny.Array
            nw32_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
            d_231_v149_ = nw32_
            index52_ = default__.safeIndex(774, (d_231_v149_).length(0))
            (d_231_v149_)[index52_] = (d_92_v47_) + (d_92_v47_)
            d_232_v150_: D2
            d_232_v150_ = D2_DC10((d_10_v1_).f7, (d_10_v1_).f7)
            d_233_v151_: D3
            d_233_v151_ = D3_DC14(d_10_v1_, (d_10_v1_).f6, d_10_v1_, d_199_v132_, (d_199_v132_).f9)
            d_234_v152_: _dafny.Set
            d_234_v152_ = _dafny.Set({D3_DC14(d_10_v1_, (d_10_v1_).f6, d_10_v1_, d_199_v132_, (d_199_v132_).f9), d_233_v151_})
            index53_ = default__.safeIndex(229, (d_184_v119_).length(0))
            index54_ = default__.safeIndex(774, (d_231_v149_).length(0))
            rhs42_ = not((not(not((d_232_v150_).cf19))) and ((d_234_v152_).isdisjoint(d_234_v152_)))
            rhs43_ = (d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))]
            rhs44_ = d_92_v47_
            lhs38_ = d_8_globalState_
            lhs39_ = d_184_v119_
            lhs40_ = default__.safeIndex(229, (d_184_v119_).length(0))
            lhs41_ = d_231_v149_
            lhs42_ = default__.safeIndex(774, (d_231_v149_).length(0))
            lhs38_.f5 = rhs42_
            lhs39_[lhs40_] = rhs43_
            lhs41_[lhs42_] = rhs44_
        elif True:
            d_235_v153_: _dafny.Seq
            d_235_v153_ = _dafny.SeqWithoutIsStrInference([d_92_v47_, (d_92_v47_) + (d_92_v47_), d_92_v47_])
            d_236_v154_: D4
            d_236_v154_ = D4_DC16(d_235_v153_)
            d_235_v153_ = ((d_235_v153_) + (((d_236_v154_).cf31).set(default__.safeIndex((0) - ((d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))]), len((d_236_v154_).cf31)), d_92_v47_))) + (d_235_v153_)
            d_237_v155_: D5
            d_237_v155_ = D5_DC18(d_184_v119_)
            d_184_v119_ = (d_237_v155_).cf37
            d_238_v156_: int
            d_239_v157_: int
            d_240_v158_: _dafny.Array
            out21_: int
            out22_: int
            out23_: _dafny.Array
            out21_, out22_, out23_ = (d_10_v1_).m0((d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))], ((d_190_v123_)[(d_10_v1_).f7] if ((d_10_v1_).f7) in (d_190_v123_) else d_14_v4_), d_8_globalState_)
            d_238_v156_ = out21_
            d_239_v157_ = out22_
            d_240_v158_ = out23_
            (d_8_globalState_).f5 = (d_10_v1_).f7
            index55_ = default__.safeIndex(704, (d_17_v7_).length(0))
            (d_17_v7_)[index55_] = (d_10_v1_).f7
            d_241_v159_: C0
            nw33_ = C0()
            nw33_.ctor__(_dafny.Map({(d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))]: d_239_v157_}), (d_10_v1_).f7)
            d_241_v159_ = nw33_
            d_242_v160_: _dafny.Seq
            d_242_v160_ = _dafny.SeqWithoutIsStrInference([d_241_v159_, d_241_v159_, d_241_v159_])
            index56_ = default__.safeIndex(704, (d_17_v7_).length(0))
            index57_ = default__.safeIndex(229, (d_184_v119_).length(0))
            rhs45_ = d_239_v157_
            rhs46_ = default__.fm6(d_8_globalState_)
            rhs47_ = ((0) - (len(d_242_v160_))) < (d_14_v4_)
            rhs48_ = (d_239_v157_) - (d_239_v157_)
            rhs49_ = 943
            lhs43_ = d_8_globalState_
            lhs44_ = d_17_v7_
            lhs45_ = default__.safeIndex(704, (d_17_v7_).length(0))
            lhs46_ = d_184_v119_
            lhs47_ = default__.safeIndex(229, (d_184_v119_).length(0))
            lhs43_.f1 = rhs45_
            d_9_v0_ = rhs46_
            lhs44_[lhs45_] = rhs47_
            d_238_v156_ = rhs48_
            lhs46_[lhs47_] = rhs49_
        d_243_v161_: _dafny.Set
        d_243_v161_ = _dafny.Set({110, (d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))]})
        d_244_v162_: _dafny.MultiSet
        d_244_v162_ = _dafny.MultiSet([d_92_v47_])
        d_245_v163_: int
        d_246_v164_: int
        d_247_v165_: _dafny.Array
        out24_: int
        out25_: int
        out26_: _dafny.Array
        out24_, out25_, out26_ = (d_10_v1_).m0((len(d_243_v161_)) - ((d_244_v162_).cardinality), ((d_184_v119_)[default__.safeIndex(229, (d_184_v119_).length(0))] if (d_10_v1_).f7 else d_14_v4_), d_8_globalState_)
        d_245_v163_ = out24_
        d_246_v164_ = out25_
        d_247_v165_ = out26_
        d_246_v164_ = 749
        _dafny.print(_dafny.string_of(d_8_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_8_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_8_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_8_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_8_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_8_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_9_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_10_v1_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_10_v1_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_12_v2_).cf14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_12_v2_).cf15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_13_v3_) == (_dafny.SeqWithoutIsStrInference([D1_DC6(_dafny.CodePoint('o'), _dafny.CodePoint('k')), D1_DC6(_dafny.CodePoint('o'), _dafny.CodePoint('k')), D1_DC6(_dafny.CodePoint('o'), _dafny.CodePoint('k')), D1_DC6(_dafny.CodePoint('o'), _dafny.CodePoint('k')), D1_DC6(_dafny.CodePoint('o'), _dafny.CodePoint('k'))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_14_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_15_v5_) == (_dafny.Map({22: 22}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_16_v6_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v7_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf6) == (_dafny.Map({22: 22}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf7) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_20_v8_).cf8)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_92_v47_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_146_i10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_183_v118_) == (_dafny.MultiSet([_dafny.Map({True: True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v119_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v119_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v119_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v119_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v119_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v119_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v119_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v119_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v119_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v119_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v119_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v120_) == (_dafny.MultiSet([22, 22]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v121_) == (_dafny.SeqWithoutIsStrInference([6]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_189_v122_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_v123_) == (_dafny.Map({False: 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_191_v124_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_192_v125_) == (_dafny.Map({True: D0_DC3(22), False: D0_DC3(22)}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_243_v161_) == (_dafny.Set({110, 872}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_244_v162_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vintro"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_245_v163_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_246_v164_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), _dafny.Set({}), int(0), _dafny.Array(None, 0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf9 == __o.cf9
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
        return lambda: D1_DC5(int(0), _dafny.Map({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D1_DC7)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC5(D1, NamedTuple('DC5', [('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC7(D1, NamedTuple('DC7', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC7({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC7) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC9(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D2_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D2_DC11)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D2_DC12)

class D2_DC9(D2, NamedTuple('DC9', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC10(D2, NamedTuple('DC10', [('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC10({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC10) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC11(D2, NamedTuple('DC11', [('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC11({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC11) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC12(D2, NamedTuple('DC12', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC12({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC12) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC14(None, _dafny.CodePoint('D'), None, None, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D3_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D3_DC13)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D3_DC15)

class D3_DC14(D3, NamedTuple('DC14', [('cf25', Any), ('cf26', Any), ('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC14({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC14) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC13(D3, NamedTuple('DC13', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC13({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC13) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC15(D3, NamedTuple('DC15', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC15({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC15) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC17(int(0), False, int(0), int(0), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D4_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D4_DC16)

class D4_DC17(D4, NamedTuple('DC17', [('cf32', Any), ('cf33', Any), ('cf34', Any), ('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC17({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC17) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC16(D4, NamedTuple('DC16', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC16({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC16) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC19(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D5_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D5_DC20)

class D5_DC19(D5, NamedTuple('DC19', [('cf38', Any), ('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC19({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC19) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC18(D5, NamedTuple('DC18', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC20(D5, NamedTuple('DC20', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC20({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC20) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC22(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D6_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D6_DC23)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D6_DC21)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D6_DC24)

class D6_DC22(D6, NamedTuple('DC22', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC22({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC22) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC23(D6, NamedTuple('DC23', [('cf44', Any), ('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC23({_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC23) and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC21(D6, NamedTuple('DC21', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC21({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC21) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC24(D6, NamedTuple('DC24', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC24({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC24) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f1: int = int(0)
        self.f5: bool = False
        self._f2: bool = False
        self._f3: int = int(0)
        self._f4: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5):
        (self).f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5

    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4

class C0:
    def  __init__(self):
        self.f8: _dafny.Map = _dafny.Map({})
        self._f9: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f8, f9):
        (self).f8 = f8
        (self)._f9 = f9

    def fm3(self, p0, p1, p2, globalState):
        return ((0) - (-152)) - ((851) * (len(_dafny.Map({(self).f9: 940}))))

    def fm4(self, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([(self).f9]), _dafny.SeqWithoutIsStrInference([(self).f9, (self).f9])])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, (self).f9]), _dafny.SeqWithoutIsStrInference([(self).f9])]))).isdisjoint((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f9]), _dafny.SeqWithoutIsStrInference([(self).f9])])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f9]), _dafny.SeqWithoutIsStrInference([(self).f9]), _dafny.SeqWithoutIsStrInference([(self).f9, False, (self).f9])])))

    @property
    def f9(self):
        return self._f9

class C1:
    def  __init__(self):
        self._f6: str = _dafny.CodePoint('D')
        self._f7: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f6, f7):
        (self)._f6 = f6
        (self)._f7 = f7

    def fm0(self, p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([(self).f6 for d_248_i0_ in range(default__.abs(-188))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkbbvoi"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcitg"))))

    def m0(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        if (self).f7:
            d_249_v0_: _dafny.Array
            def lambda10_(d_250_p1_):
                def lambda11_(d_251_i0_):
                    return (_dafny.Map({d_250_p1_: _dafny.SeqWithoutIsStrInference([(self).f7])}) if (self).f7 else _dafny.Map({d_250_p1_: _dafny.SeqWithoutIsStrInference([(self).f7, (self).f7])}))

                return lambda11_

            init5_ = lambda10_(p1)
            nw34_ = _dafny.Array(None, 21)
            for i0_5_ in range(nw34_.length(0)):
                nw34_[i0_5_] = init5_(i0_5_)
            d_249_v0_ = nw34_
            d_252_v1_: _dafny.Seq
            d_252_v1_ = _dafny.SeqWithoutIsStrInference([(self).f7])
            d_253_v2_: _dafny.Map
            d_253_v2_ = _dafny.Map({p0: d_252_v1_})
            index58_ = default__.safeIndex(572, (d_249_v0_).length(0))
            (d_249_v0_)[index58_] = d_253_v2_
            d_254_v3_: _dafny.Array
            nw35_ = _dafny.Array(int(0), 5)
            d_254_v3_ = nw35_
            index59_ = default__.safeIndex(365, (d_254_v3_).length(0))
            (d_254_v3_)[index59_] = p1
            index60_ = default__.safeIndex(572, (d_249_v0_).length(0))
            index61_ = default__.safeIndex(365, (d_254_v3_).length(0))
            rhs50_ = default__.safeModuloInt(p1, len(d_252_v1_))
            rhs51_ = (_dafny.Map({p1: d_252_v1_}) if default__.fm1(globalState) else (d_253_v2_) | (d_253_v2_))
            rhs52_ = 48
            rhs53_ = 682
            lhs48_ = globalState
            lhs49_ = d_249_v0_
            lhs50_ = default__.safeIndex(572, (d_249_v0_).length(0))
            lhs51_ = d_254_v3_
            lhs52_ = default__.safeIndex(365, (d_254_v3_).length(0))
            lhs48_.f1 = rhs50_
            lhs49_[lhs50_] = rhs51_
            r1 = rhs52_
            lhs51_[lhs52_] = rhs53_
            (globalState).f5 = ((self).f7 if False else (self).f7)
            d_255_v4_: _dafny.Seq
            d_255_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfkjko"))
            d_256_v5_: _dafny.Map
            d_256_v5_ = _dafny.Map({d_255_v4_: (d_254_v3_)[default__.safeIndex(365, (d_254_v3_).length(0))]})
            d_256_v5_ = (d_256_v5_).set(d_255_v4_, (d_254_v3_)[default__.safeIndex(365, (d_254_v3_).length(0))])
            d_257_v6_: _dafny.Map
            d_257_v6_ = _dafny.Map({(self).f7: p1})
            d_258_v7_: _dafny.Map
            d_258_v7_ = _dafny.Map({d_255_v4_: _dafny.Map({(self).f7: (self).f7})})
            d_259_v8_: _dafny.Map
            d_259_v8_ = _dafny.Map({not((self).f7): (self).f7})
            (globalState).f1 = (default__.fm2(default__.fm1(globalState), globalState) if ((self).f7) not in (d_257_v6_) else (len(((d_258_v7_)[d_255_v4_] if (d_255_v4_) in (d_258_v7_) else d_259_v8_)) if (self).f7 else (d_254_v3_)[default__.safeIndex(365, (d_254_v3_).length(0))]))
            d_260_v9_: _dafny.Map
            d_260_v9_ = _dafny.Map({(self).f6: ((self).f7) or ((self).f7)})
            d_260_v9_ = (d_260_v9_).set((self).f6, False)
        elif True:
            d_261_v10_: _dafny.Map
            d_261_v10_ = _dafny.Map({default__.fm2((self).f7, globalState): (self).f7})
            d_262_v11_: _dafny.MultiSet
            d_262_v11_ = _dafny.MultiSet([(self).f7])
            d_263_v12_: C0
            nw36_ = C0()
            nw36_.ctor__(default__.fm5(len(d_261_v10_), (self).f7, (self).f7, globalState), (_dafny.MultiSet([(self).f7, (self).f7])) == (d_262_v11_))
            d_263_v12_ = nw36_
            d_264_v13_: _dafny.Array
            nw37_ = _dafny.Array(False, 28)
            d_264_v13_ = nw37_
            d_265_v14_: D0
            d_265_v14_ = D0_DC0(d_264_v13_)
            d_264_v13_ = (d_265_v14_).cf0
            d_266_v15_: _dafny.Map
            d_266_v15_ = _dafny.Map({(d_263_v12_).f9: p0})
            d_266_v15_ = (d_266_v15_).set((d_263_v12_).f9, default__.fm2((self).f7, globalState))
            d_267_v16_: str
            d_267_v16_ = _dafny.CodePoint('q')
            d_267_v16_ = ((self).f6 if not((d_263_v12_).f9) else (self).f6)
            d_268_v17_: _dafny.Array
            nw38_ = _dafny.Array(_dafny.Set({}), 18)
            d_268_v17_ = nw38_
            d_269_v18_: _dafny.Set
            d_269_v18_ = _dafny.Set({d_264_v13_, d_264_v13_})
            d_270_v19_: _dafny.Map
            d_270_v19_ = _dafny.Map({d_269_v18_: False})
            d_268_v17_ = (d_268_v17_ if (((d_270_v19_)[d_269_v18_] if (d_269_v18_) in (d_270_v19_) else (d_263_v12_).fm4(globalState)) if (d_263_v12_).f9 else (self).f7) else d_268_v17_)
        d_271_v20_: _dafny.Seq
        d_271_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swfjwfa"))
        d_272_v21_: D1
        d_272_v21_ = D1_DC4(False)
        d_273_v22_: _dafny.Set
        d_273_v22_ = _dafny.Set({p0, len(d_271_v20_), p1})
        d_274_v23_: _dafny.MultiSet
        d_274_v23_ = _dafny.MultiSet([len(d_273_v22_)])
        d_275_v24_: _dafny.Array
        nw39_ = _dafny.Array(None, 15)
        nw39_[int(0)] = (self).f7
        nw39_[int(1)] = not((self).f7)
        nw39_[int(2)] = (p1) > (p0)
        nw39_[int(3)] = not((d_271_v20_) == (d_271_v20_))
        nw39_[int(4)] = not((d_272_v21_).cf10)
        nw39_[int(5)] = (self).f7
        nw39_[int(6)] = not(default__.fm1(globalState))
        nw39_[int(7)] = (self).f7
        nw39_[int(8)] = not((self).f7)
        nw39_[int(9)] = (d_274_v23_) == (d_274_v23_)
        nw39_[int(10)] = default__.fm1(globalState)
        nw39_[int(11)] = (self).f7
        nw39_[int(12)] = (self).f7
        nw39_[int(13)] = (self).f7
        nw39_[int(14)] = (self).f7
        d_275_v24_ = nw39_
        index62_ = default__.safeIndex(448, (d_275_v24_).length(0))
        (d_275_v24_)[index62_] = (p0) < (default__.fm2((self).f7, globalState))
        index63_ = default__.safeIndex(448, (d_275_v24_).length(0))
        (d_275_v24_)[index63_] = (self).f7
        (globalState).f5 = (self).f7
        d_276_v25_: _dafny.Map
        d_276_v25_ = _dafny.Map({-673: p1})
        d_277_v26_: _dafny.Set
        d_277_v26_ = _dafny.Set({((d_271_v20_) + (d_271_v20_)).set(default__.safeIndex(len(d_276_v25_), len((d_271_v20_) + (d_271_v20_))), _dafny.CodePoint('h'))})
        r1 = len(d_277_v26_)
        d_278_v27_: _dafny.Set
        d_278_v27_ = _dafny.Set({(self).f7, (self).f7, (self).f7})
        d_279_v28_: _dafny.Array
        nw40_ = _dafny.Array(_dafny.Seq({}), 14)
        d_279_v28_ = nw40_
        d_280_v29_: D0
        d_280_v29_ = D0_DC1(p1, d_278_v27_, p0, d_279_v28_, (self).f6)
        d_281_v30_: _dafny.Array
        nw41_ = _dafny.Array(None, 19)
        nw41_[int(0)] = d_280_v29_
        nw41_[int(1)] = d_280_v29_
        nw41_[int(2)] = d_280_v29_
        nw41_[int(3)] = d_280_v29_
        nw41_[int(4)] = d_280_v29_
        nw41_[int(5)] = d_280_v29_
        nw41_[int(6)] = d_280_v29_
        nw41_[int(7)] = d_280_v29_
        nw41_[int(8)] = d_280_v29_
        nw41_[int(9)] = d_280_v29_
        nw41_[int(10)] = d_280_v29_
        nw41_[int(11)] = d_280_v29_
        nw41_[int(12)] = d_280_v29_
        nw41_[int(13)] = d_280_v29_
        nw41_[int(14)] = d_280_v29_
        nw41_[int(15)] = d_280_v29_
        nw41_[int(16)] = d_280_v29_
        nw41_[int(17)] = d_280_v29_
        nw41_[int(18)] = d_280_v29_
        d_281_v30_ = nw41_
        d_282_v31_: _dafny.Map
        d_282_v31_ = _dafny.Map({d_281_v30_: (self).f7})
        d_282_v31_ = (d_282_v31_).set(d_281_v30_, (d_275_v24_)[default__.safeIndex(448, (d_275_v24_).length(0))])
        (globalState).f5 = (d_275_v24_)[default__.safeIndex(448, (d_275_v24_).length(0))]
        d_283_v32_: _dafny.Map
        d_283_v32_ = _dafny.Map({False: 192})
        r0 = default__.safeModuloInt((len(_dafny.Set({True})) if (d_275_v24_)[default__.safeIndex(448, (d_275_v24_).length(0))] else p1), len(d_283_v32_))
        d_284_v33_: C0
        nw42_ = C0()
        nw42_.ctor__(d_276_v25_, not((self).f7))
        d_284_v33_ = nw42_
        d_285_v34_: _dafny.Array
        nw43_ = _dafny.Array(None, 14)
        nw43_[int(0)] = d_284_v33_
        nw43_[int(1)] = d_284_v33_
        nw43_[int(2)] = d_284_v33_
        nw43_[int(3)] = d_284_v33_
        nw43_[int(4)] = d_284_v33_
        nw43_[int(5)] = d_284_v33_
        nw43_[int(6)] = d_284_v33_
        nw43_[int(7)] = d_284_v33_
        nw43_[int(8)] = d_284_v33_
        nw43_[int(9)] = d_284_v33_
        nw43_[int(10)] = d_284_v33_
        nw43_[int(11)] = d_284_v33_
        nw43_[int(12)] = d_284_v33_
        nw43_[int(13)] = d_284_v33_
        d_285_v34_ = nw43_
        d_286_v35_: _dafny.Map
        d_286_v35_ = _dafny.Map({True: False})
        r1 = len(_dafny.Map({(d_285_v34_ if (self).f7 else d_285_v34_): ((self).f7 if ((d_286_v35_)[(self).f7] if ((self).f7) in (d_286_v35_) else (self).f7) else True)}))
        d_287_v36_: _dafny.Array
        nw44_ = _dafny.Array(int(0), 4)
        d_287_v36_ = nw44_
        r2 = d_287_v36_
        return r0, r1, r2

    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
