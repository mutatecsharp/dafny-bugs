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
    def fm0(p0, p1, p2, globalState):
        if (False) and (False):
            return not (True) or (not(True))
        elif True:
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_0_i0_ in range(default__.abs(402))])) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdxgcrjxl")))

    @staticmethod
    def fm1(p0, p1, globalState):
        return (_dafny.MultiSet([not(not(False))])).intersection((_dafny.MultiSet([not(False), True])).intersection(_dafny.MultiSet([True, False])))

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([223])).Elements:
                d_1_v0_: int = compr_0_
                if (d_1_v0_) in (_dafny.SeqWithoutIsStrInference([223])):
                    coll0_ = coll0_.union(_dafny.Set([(d_1_v0_) + (741)]))
            return _dafny.Set(coll0_)
        return ((len(_dafny.SeqWithoutIsStrInference([240, -231, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rljhwsp"))), 206])) if False else len(_dafny.Map({_dafny.CodePoint('d'): len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xe"))): len(iife0_()
        )}))})))) + ((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d')]))) * (-387))

    @staticmethod
    def fm7(p0, globalState):
        if True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cs"))
        elif True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "huytibwmr"))

    @staticmethod
    def fm10(globalState):
        source0_ = D2_DC6(D2_DC6(D2_DC5(-139, True, True, True)))
        if source0_.is_DC4:
            d_2___mcc_h0_ = source0_.cf4
            d_3___mcc_h1_ = source0_.cf5
            d_4_cf5_ = d_3___mcc_h1_
            d_5_cf4_ = d_2___mcc_h0_
            return _dafny.CodePoint('t')
        elif source0_.is_DC5:
            d_6___mcc_h2_ = source0_.cf6
            d_7___mcc_h3_ = source0_.cf7
            d_8___mcc_h4_ = source0_.cf8
            d_9___mcc_h5_ = source0_.cf9
            d_10_cf9_ = d_9___mcc_h5_
            d_11_cf8_ = d_8___mcc_h4_
            d_12_cf7_ = d_7___mcc_h3_
            d_13_cf6_ = d_6___mcc_h2_
            return _dafny.CodePoint('v')
        elif source0_.is_DC3:
            d_14___mcc_h6_ = source0_.cf3
            d_15_cf3_ = d_14___mcc_h6_
            if False:
                return _dafny.CodePoint('e')
            elif True:
                return _dafny.CodePoint('n')
        elif True:
            d_16___mcc_h7_ = source0_.cf10
            d_17_cf10_ = d_16___mcc_h7_
            return _dafny.CodePoint('o')

    @staticmethod
    def fm11(globalState):
        return D2_DC4((855) + (195), (_dafny.Set({188})).ispropersubset(_dafny.Set({-694, 297})))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(0) - ((354) - (954))])

    @staticmethod
    def fm13(p0, p1, globalState):
        return (_dafny.Map({not(not(not(not(not(False))))): 430})) | ((_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, not(True)]), _dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([True])]))}) if True else _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([False]))})))

    @staticmethod
    def fm16(p0, globalState):
        return _dafny.CodePoint('a')

    @staticmethod
    def fm17(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([not(not(True)), False, False, (_dafny.Set({360})).issubset(_dafny.Set({35}))])

    @staticmethod
    def fm18(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ckp"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dy")) if not(False) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_18_i0_ in range(default__.abs(360))])))

    @staticmethod
    def fm19(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([-691, 373])).Elements:
                d_19_v0_: int = compr_1_
                if (d_19_v0_) in (_dafny.SeqWithoutIsStrInference([-691, 373])):
                    coll1_[(d_19_v0_) * (len(_dafny.SeqWithoutIsStrInference([310])))] = _dafny.CodePoint('x')
            return _dafny.Map(coll1_)
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: D3
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([D3_DC9(len(_dafny.SeqWithoutIsStrInference([False])), _dafny.CodePoint('c'), 512)])).Elements:
                d_20_v1_: D3 = compr_2_
                if (d_20_v1_) in (_dafny.SeqWithoutIsStrInference([D3_DC9(len(_dafny.SeqWithoutIsStrInference([False])), _dafny.CodePoint('c'), 512)])):
                    coll2_[d_20_v1_] = False
            return _dafny.Map(coll2_)
        return ((_dafny.Map({len(_dafny.Map({True: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kimkmjliq"))))})): _dafny.CodePoint('x')})) | (iife1_()
        )) | (_dafny.Map({len(iife2_()
        ): _dafny.CodePoint('d')}))

    @staticmethod
    def fm20(p0, globalState):
        return D2_DC6(D2_DC5((_dafny.MultiSet([_dafny.CodePoint('n')])).cardinality, not(not(True)), not(False), False))

    @staticmethod
    def fm21(globalState):
        return _dafny.Map({len((_dafny.SeqWithoutIsStrInference([640])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "moh"))) for d_21_i0_ in range(default__.abs(-152))])), 532]))): default__.safeDivisionInt(798, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hnvcppjf"))))})

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        if not(False):
            return _dafny.Set({_dafny.CodePoint('p')})
        elif True:
            return _dafny.Set({_dafny.CodePoint('h'), _dafny.CodePoint('l')})

    @staticmethod
    def fm23(p0, globalState):
        return D7_DC21(-405, (_dafny.Map({not(True): True})) | (_dafny.Map({True: False})))

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        return ((_dafny.Set({len(_dafny.Set({not(not(False))}))})) | (_dafny.Set({(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdxocc")))))}))) | (_dafny.Set({449}))

    @staticmethod
    def fm25(globalState):
        if False:
            return D3_DC7(_dafny.Map({(D7_DC22(len(_dafny.Map({len(_dafny.Map({693: len(_dafny.Map({True: False}))})): 42})), len(_dafny.Set({_dafny.CodePoint('p')})), False)).cf40: True}))
        elif True:
            return D3_DC7(_dafny.Map({319: True}))

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_22_i1_ in range(default__.abs(-738))])) for d_23_i0_ in range(default__.abs(385))])).Elements:
                d_24_v0_: int = compr_3_
                if (d_24_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_22_i1_ in range(default__.abs(-738))])) for d_23_i0_ in range(default__.abs(385))])):
                    coll3_ = coll3_.union(_dafny.Set([(d_24_v0_) + (686)]))
            return _dafny.Set(coll3_)
        return (_dafny.Map({len(_dafny.Set({(0) - ((_dafny.MultiSet([False])).cardinality), len(iife3_()
        )})): not(True)})) | (_dafny.Map({-473: True}))

    @staticmethod
    def fm27(globalState):
        source1_ = D3_DC8(False)
        if source1_.is_DC8:
            d_25___mcc_h0_ = source1_.cf12
            d_26_cf12_ = d_25___mcc_h0_
            return _dafny.Map({d_26_cf12_: not(not(d_26_cf12_))})
        elif source1_.is_DC9:
            d_27___mcc_h1_ = source1_.cf13
            d_28___mcc_h2_ = source1_.cf14
            d_29___mcc_h3_ = source1_.cf15
            d_30_cf15_ = d_29___mcc_h3_
            d_31_cf14_ = d_28___mcc_h2_
            d_32_cf13_ = d_27___mcc_h1_
            return _dafny.Map({False: True})
        elif source1_.is_DC7:
            d_33___mcc_h4_ = source1_.cf11
            d_34_cf11_ = d_33___mcc_h4_
            return (_dafny.Map({not(not(True)): not(not(False))})) | (_dafny.Map({True: False}))
        elif True:
            d_35___mcc_h5_ = source1_.cf16
            d_36_cf16_ = d_35___mcc_h5_
            return (D7_DC21(len(_dafny.Map({True: False})), _dafny.Map({False: True}))).cf39

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(235, -647):
                d_37_v0_: int = compr_4_
                if ((235) <= (d_37_v0_)) and ((d_37_v0_) < (-647)):
                    coll4_[(d_37_v0_) - (54)] = True
            return _dafny.Map(coll4_)
        return _dafny.Map({(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))})) | (_dafny.Set({(0) - ((_dafny.MultiSet([192, 255])).cardinality)})): iife4_()
        })

    @staticmethod
    def m0(globalState):
        r0: int = int(0)
        d_38_v0_: int
        d_38_v0_ = 142
        d_39_v1_: _dafny.Seq
        d_39_v1_ = _dafny.SeqWithoutIsStrInference([d_38_v0_])
        d_40_v2_: _dafny.Seq
        d_40_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfkqvu"))
        d_41_v3_: _dafny.Set
        d_41_v3_ = _dafny.Set({d_38_v0_, (d_39_v1_)[default__.safeIndex(len(d_40_v2_), len(d_39_v1_))]})
        hi0_ = default__.safeDivisionInt(d_38_v0_, len(d_41_v3_))
        for d_42_i0_ in range(927, hi0_):
            d_43_v4_: bool
            d_43_v4_ = False
            (globalState).f2 = d_43_v4_
            d_44_v5_: D4
            d_44_v5_ = D4_DC11(d_40_v2_)
            d_43_v4_ = (d_40_v2_) < ((d_44_v5_).cf17)
            d_45_v6_: C0
            nw0_ = C0()
            nw0_.ctor__()
            d_45_v6_ = nw0_
            d_46_v7_: str
            d_46_v7_ = _dafny.CodePoint('m')
            d_47_v8_: _dafny.Set
            d_47_v8_ = _dafny.Set({d_46_v7_, d_46_v7_, d_46_v7_, d_46_v7_, d_46_v7_})
            d_48_v9_: _dafny.MultiSet
            d_48_v9_ = _dafny.MultiSet([d_47_v8_])
            rhs0_ = d_42_i0_
            rhs1_ = default__.safeDivisionInt(d_38_v0_, d_42_i0_)
            rhs2_ = (d_48_v9_).intersection(d_48_v9_)
            rhs3_ = d_38_v0_
            lhs0_ = globalState
            lhs1_ = globalState
            lhs2_ = globalState
            lhs0_.f18 = rhs0_
            lhs1_.f21 = rhs1_
            d_48_v9_ = rhs2_
            lhs2_.f5 = rhs3_
        if False:
            d_49_v10_: bool
            d_49_v10_ = False
            (globalState).f2 = not((d_40_v2_) <= ((d_40_v2_ if d_49_v10_ else d_40_v2_)))
            d_50_v11_: _dafny.Set
            d_50_v11_ = _dafny.Set({d_49_v10_, d_49_v10_, d_49_v10_})
            if not(default__.fm0((_dafny.MultiSet([len(d_50_v11_), d_38_v0_])).cardinality, default__.safeDivisionInt(d_38_v0_, 775), d_49_v10_, globalState)):
                d_51_v12_: str
                d_51_v12_ = _dafny.CodePoint('u')
                d_52_v13_: _dafny.Array
                nw1_ = _dafny.Array(_dafny.MultiSet({}), 9)
                d_52_v13_ = nw1_
                d_53_v14_: C2
                nw2_ = C2()
                nw2_.ctor__(d_51_v12_, d_52_v13_, True)
                d_53_v14_ = nw2_
                d_54_v15_: _dafny.Array
                nw3_ = _dafny.Array(int(0), 25)
                d_54_v15_ = nw3_
                d_55_v16_: D4
                d_55_v16_ = D4_DC12(True, d_54_v15_, d_38_v0_)
                (globalState).f2 = (d_55_v16_).cf18
                d_56_v17_: _dafny.Seq
                d_56_v17_ = _dafny.SeqWithoutIsStrInference([d_49_v10_])
                (globalState).f2 = (len(d_56_v17_)) >= (d_38_v0_)
                d_54_v15_ = d_54_v15_
                d_57_v18_: C0
                nw4_ = C0()
                nw4_.ctor__()
                d_57_v18_ = nw4_
            elif True:
                (globalState).f2 = d_49_v10_
                d_58_v19_: _dafny.Array
                def lambda0_(d_59_v0_):
                    def lambda1_(d_60_i1_):
                        return default__.safeModuloInt(d_60_i1_, d_59_v0_)

                    return lambda1_

                init0_ = lambda0_(d_38_v0_)
                nw5_ = _dafny.Array(None, 23)
                for i0_0_ in range(nw5_.length(0)):
                    nw5_[i0_0_] = init0_(i0_0_)
                d_58_v19_ = nw5_
                index0_ = default__.safeIndex(47, (d_58_v19_).length(0))
                (d_58_v19_)[index0_] = (len((d_39_v1_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ec"))), len(d_39_v1_)), d_38_v0_)) if d_49_v10_ else len(d_41_v3_))
                index1_ = default__.safeIndex(47, (d_58_v19_).length(0))
                (d_58_v19_)[index1_] = (d_39_v1_)[default__.safeIndex(d_38_v0_, len(d_39_v1_))]
                d_40_v2_ = d_40_v2_
                (globalState).f21 = (d_58_v19_)[default__.safeIndex(47, (d_58_v19_).length(0))]
                d_61_v20_: _dafny.Array
                nw6_ = _dafny.Array(None, 7)
                nw6_[int(0)] = True
                nw6_[int(1)] = d_49_v10_
                nw6_[int(2)] = not(d_49_v10_)
                nw6_[int(3)] = d_49_v10_
                nw6_[int(4)] = d_49_v10_
                nw6_[int(5)] = d_49_v10_
                nw6_[int(6)] = d_49_v10_
                d_61_v20_ = nw6_
                d_62_v21_: _dafny.MultiSet
                d_62_v21_ = _dafny.MultiSet([d_61_v20_])
                index2_ = default__.safeIndex(47, (d_58_v19_).length(0))
                (d_58_v19_)[index2_] = (d_62_v21_).cardinality
            (globalState).f5 = default__.safeModuloInt(d_38_v0_, 860)
            d_63_v22_: C0
            nw7_ = C0()
            nw7_.ctor__()
            d_63_v22_ = nw7_
            d_64_v23_: _dafny.Seq
            d_64_v23_ = _dafny.SeqWithoutIsStrInference([not(d_49_v10_), d_49_v10_])
            d_65_v24_: _dafny.Array
            nw8_ = _dafny.Array(None, 12)
            nw8_[int(0)] = d_38_v0_
            nw8_[int(1)] = d_38_v0_
            nw8_[int(2)] = len(d_64_v23_)
            nw8_[int(3)] = d_38_v0_
            nw8_[int(4)] = ((0) - (d_38_v0_)) + ((0) - (d_38_v0_))
            nw8_[int(5)] = d_38_v0_
            nw8_[int(6)] = d_38_v0_
            nw8_[int(7)] = d_38_v0_
            nw8_[int(8)] = d_38_v0_
            nw8_[int(9)] = d_38_v0_
            nw8_[int(10)] = (0) - ((d_38_v0_) * (250))
            nw8_[int(11)] = d_38_v0_
            d_65_v24_ = nw8_
            d_66_v25_: _dafny.Map
            d_66_v25_ = _dafny.Map({d_38_v0_: d_49_v10_})
            index3_ = default__.safeIndex(794, (d_65_v24_).length(0))
            (d_65_v24_)[index3_] = default__.safeDivisionInt(d_38_v0_, len(_dafny.Map({d_38_v0_: d_66_v25_})))
            index4_ = default__.safeIndex(794, (d_65_v24_).length(0))
            (d_65_v24_)[index4_] = (d_63_v22_).fm5(globalState)
        elif True:
            d_67_v26_: str
            d_67_v26_ = _dafny.CodePoint('p')
            d_68_v27_: _dafny.MultiSet
            d_68_v27_ = _dafny.MultiSet([d_67_v26_])
            d_69_v28_: _dafny.Array
            nw9_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_69_v28_ = nw9_
            d_70_v29_: bool
            d_70_v29_ = False
            d_71_v30_: D2
            d_71_v30_ = D2_DC4(d_38_v0_, d_70_v29_)
            pat_let_tv0_ = d_70_v29_
            d_72_v31_: _dafny.Array
            nw10_ = _dafny.Array(None, 26)
            nw10_[int(0)] = d_71_v30_
            nw10_[int(1)] = d_71_v30_
            nw10_[int(2)] = d_71_v30_
            nw10_[int(3)] = d_71_v30_
            nw10_[int(4)] = d_71_v30_
            nw10_[int(5)] = d_71_v30_
            nw10_[int(6)] = d_71_v30_
            nw10_[int(7)] = D2_DC4((0) - (d_38_v0_), d_70_v29_)
            nw10_[int(8)] = d_71_v30_
            nw10_[int(9)] = d_71_v30_
            nw10_[int(10)] = d_71_v30_
            nw10_[int(11)] = D2_DC4(d_38_v0_, d_70_v29_)
            nw10_[int(12)] = d_71_v30_
            nw10_[int(13)] = d_71_v30_
            nw10_[int(14)] = d_71_v30_
            nw10_[int(15)] = d_71_v30_
            def iife5_(_pat_let0_0):
                def iife6_(d_73_dt__update__tmp_h0_):
                    def iife7_(_pat_let1_0):
                        def iife8_(d_74_dt__update_hcf5_h0_):
                            return D2_DC4((d_73_dt__update__tmp_h0_).cf4, d_74_dt__update_hcf5_h0_)
                        return iife8_(_pat_let1_0)
                    return iife7_(pat_let_tv0_)
                return iife6_(_pat_let0_0)
            nw10_[int(16)] = iife5_(D2_DC4(d_38_v0_, d_70_v29_))
            nw10_[int(17)] = d_71_v30_
            nw10_[int(18)] = d_71_v30_
            nw10_[int(19)] = d_71_v30_
            nw10_[int(20)] = d_71_v30_
            nw10_[int(21)] = D2_DC4(len(d_40_v2_), d_70_v29_)
            nw10_[int(22)] = d_71_v30_
            nw10_[int(23)] = d_71_v30_
            nw10_[int(24)] = D2_DC4((0) - (d_38_v0_), False)
            nw10_[int(25)] = d_71_v30_
            d_72_v31_ = nw10_
            index5_ = default__.safeIndex(395, (d_69_v28_).length(0))
            (d_69_v28_)[index5_] = d_72_v31_
            d_75_v32_: _dafny.Set
            d_75_v32_ = _dafny.Set({d_70_v29_, d_70_v29_})
            d_76_v33_: _dafny.Seq
            d_76_v33_ = _dafny.SeqWithoutIsStrInference([d_72_v31_, d_72_v31_, (d_72_v31_ if d_70_v29_ else d_72_v31_)])
            index6_ = default__.safeIndex(395, (d_69_v28_).length(0))
            rhs4_ = ((d_68_v27_).set(d_67_v26_, default__.abs(len(d_75_v32_)))) - (d_68_v27_)
            rhs5_ = (d_76_v33_)[default__.safeIndex(d_38_v0_, len(d_76_v33_))]
            rhs6_ = (0) - (d_38_v0_)
            rhs7_ = d_41_v3_
            rhs8_ = default__.fm2((d_38_v0_) == ((_dafny.MultiSet(d_39_v1_)).cardinality), (0) - (default__.fm2(d_70_v29_, (0) - (d_38_v0_), d_70_v29_, d_67_v26_, globalState)), (d_38_v0_) >= (d_38_v0_), _dafny.CodePoint('d'), globalState)
            lhs3_ = d_69_v28_
            lhs4_ = default__.safeIndex(395, (d_69_v28_).length(0))
            lhs5_ = globalState
            lhs6_ = globalState
            d_68_v27_ = rhs4_
            lhs3_[lhs4_] = rhs5_
            lhs5_.f18 = rhs6_
            d_41_v3_ = rhs7_
            lhs6_.f5 = rhs8_
            (globalState).f2 = d_70_v29_
            d_77_v34_: _dafny.Array
            nw11_ = _dafny.Array(None, 18)
            d_77_v34_ = nw11_
            d_78_v35_: _dafny.Array
            def lambda2_(d_79_v0_):
                def lambda3_(d_80_i2_):
                    return _dafny.MultiSet([d_79_v0_, d_79_v0_])

                return lambda3_

            init1_ = lambda2_(d_38_v0_)
            nw12_ = _dafny.Array(None, 12)
            for i0_1_ in range(nw12_.length(0)):
                nw12_[i0_1_] = init1_(i0_1_)
            d_78_v35_ = nw12_
            d_81_v36_: C2
            nw13_ = C2()
            nw13_.ctor__(d_67_v26_, d_78_v35_, d_70_v29_)
            d_81_v36_ = nw13_
            index7_ = default__.safeIndex(570, (d_77_v34_).length(0))
            (d_77_v34_)[index7_] = d_81_v36_
            index8_ = default__.safeIndex(570, (d_77_v34_).length(0))
            (d_77_v34_)[index8_] = d_81_v36_
            d_82_v37_: _dafny.Array
            def lambda4_(d_83_v2_):
                def lambda5_(d_84_i3_):
                    return (_dafny.CodePoint('n')) not in (d_83_v2_)

                return lambda5_

            init2_ = lambda4_(d_40_v2_)
            nw14_ = _dafny.Array(None, 4)
            for i0_2_ in range(nw14_.length(0)):
                nw14_[i0_2_] = init2_(i0_2_)
            d_82_v37_ = nw14_
            d_85_v38_: _dafny.Seq
            d_85_v38_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_38_v0_ for d_86_i4_ in range(default__.abs(-407))])])
            index9_ = default__.safeIndex(744, (d_82_v37_).length(0))
            (d_82_v37_)[index9_] = (len((d_85_v38_)[default__.safeIndex(d_38_v0_, len(d_85_v38_))])) != (d_38_v0_)
            index10_ = default__.safeIndex(744, (d_82_v37_).length(0))
            (d_82_v37_)[index10_] = (d_70_v29_ if (d_39_v1_) < (_dafny.SeqWithoutIsStrInference([d_38_v0_, len(d_39_v1_)])) else d_70_v29_)
            (globalState).f19 = _dafny.SeqWithoutIsStrInference([d_38_v0_ for d_87_i5_ in range(default__.abs(-27))])
        d_88_v39_: bool
        d_88_v39_ = False
        d_89_v40_: str
        d_89_v40_ = _dafny.CodePoint('p')
        d_90_v41_: T0
        nw15_ = C1()
        nw15_.ctor__(d_40_v2_, d_38_v0_, d_88_v39_)
        d_90_v41_ = nw15_
        d_91_v42_: _dafny.Seq
        d_91_v42_ = _dafny.SeqWithoutIsStrInference([d_90_v41_])
        d_92_v43_: _dafny.Seq
        d_92_v43_ = _dafny.SeqWithoutIsStrInference([(d_91_v42_)[default__.safeIndex(d_38_v0_, len(d_91_v42_))]])
        d_93_v44_: _dafny.Map
        d_93_v44_ = _dafny.Map({default__.fm2(d_88_v39_, d_38_v0_, d_88_v39_, d_89_v40_, globalState): len(d_92_v43_)})
        def iife9_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in (d_39_v1_).Elements:
                d_94_v45_: int = compr_5_
                if (d_94_v45_) in (d_39_v1_):
                    coll5_[default__.safeDivisionInt(d_94_v45_, d_38_v0_)] = 2
            return _dafny.Map(coll5_)
        (globalState).f2 = ((d_93_v44_) | (iife9_()
        )) not in (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(d_40_v2_): d_38_v0_})]))
        d_95_v46_: _dafny.Array
        nw16_ = _dafny.Array(_dafny.MultiSet({}), 7)
        d_95_v46_ = nw16_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_95_v46_).length(0)):
            d_96_i6_: int = guard_loop_0_
            if (True) and (((0) <= (d_96_i6_)) and ((d_96_i6_) < ((d_95_v46_).length(0)))):
                (d_95_v46_)[(d_96_i6_)] = _dafny.MultiSet([(D2_DC5(d_38_v0_, (d_90_v41_).f25, (d_90_v41_).f25, (d_90_v41_).f25)).cf6, 843])
        d_97_v47_: _dafny.MultiSet
        d_97_v47_ = _dafny.MultiSet([d_38_v0_])
        d_98_v49_: _dafny.Map
        d_98_v49_ = _dafny.Map({d_38_v0_: default__.fm0(-145, (0) - (d_38_v0_), True, globalState)})
        d_99_v50_: _dafny.Set
        d_99_v50_ = _dafny.Set({((d_98_v49_)[d_38_v0_] if (d_38_v0_) in (d_98_v49_) else d_88_v39_), d_88_v39_})
        d_100_v51_: _dafny.Map
        d_100_v51_ = _dafny.Map({(d_90_v41_).f25: d_89_v40_})
        d_101_v52_: _dafny.Map
        d_101_v52_ = _dafny.Map({d_38_v0_: d_100_v51_})
        d_102_v53_: _dafny.Map
        d_102_v53_ = _dafny.Map({(d_90_v41_).f25: d_88_v39_})
        d_103_v54_: D7
        d_103_v54_ = D7_DC21(d_38_v0_, default__.fm27(globalState))
        d_104_v55_: C3
        nw17_ = C3()
        nw17_.ctor__(d_89_v40_, d_95_v46_, (d_90_v41_).f25)
        d_104_v55_ = nw17_
        d_105_v56_: _dafny.Array
        nw18_ = _dafny.Array(None, 28)
        nw18_[int(0)] = len(d_39_v1_)
        nw18_[int(1)] = (0) - (d_38_v0_)
        nw18_[int(2)] = d_38_v0_
        nw18_[int(3)] = d_38_v0_
        nw18_[int(4)] = default__.safeDivisionInt(default__.fm2(default__.fm0(d_38_v0_, d_38_v0_, d_88_v39_, globalState), (d_97_v47_).cardinality, d_88_v39_, d_89_v40_, globalState), d_38_v0_)
        nw18_[int(5)] = d_38_v0_
        nw18_[int(6)] = len(_dafny.SeqWithoutIsStrInference([d_89_v40_ for d_106_i7_ in range(default__.abs(616))]))
        nw18_[int(7)] = d_38_v0_
        nw18_[int(8)] = (0) - ((0) - ((d_38_v0_) * (d_38_v0_)))
        nw18_[int(9)] = d_38_v0_
        nw18_[int(10)] = d_38_v0_
        def iife10_():
            coll6_ = _dafny.Set()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(117, 380):
                d_107_v48_: int = compr_6_
                if ((117) <= (d_107_v48_)) and ((d_107_v48_) < (380)):
                    coll6_ = coll6_.union(_dafny.Set([default__.safeDivisionInt(d_107_v48_, d_38_v0_)]))
            return _dafny.Set(coll6_)
        nw18_[int(11)] = (len(iife10_()
        )) - ((0) - ((0) - (d_38_v0_)))
        nw18_[int(12)] = (len(d_99_v50_)) + (4)
        nw18_[int(13)] = d_38_v0_
        nw18_[int(14)] = d_38_v0_
        nw18_[int(15)] = 111
        nw18_[int(16)] = (len(d_101_v52_)) * (d_38_v0_)
        nw18_[int(17)] = d_38_v0_
        nw18_[int(18)] = default__.fm2(d_88_v39_, 191, d_88_v39_, d_89_v40_, globalState)
        nw18_[int(19)] = d_38_v0_
        nw18_[int(20)] = d_38_v0_
        nw18_[int(21)] = d_38_v0_
        nw18_[int(22)] = len(((d_102_v53_).set(d_88_v39_, (d_90_v41_).f25)) | ((d_103_v54_).cf39))
        nw18_[int(23)] = d_38_v0_
        nw18_[int(24)] = default__.safeModuloInt(d_38_v0_, (0) - (d_38_v0_))
        nw18_[int(25)] = 77
        nw18_[int(26)] = d_38_v0_
        nw18_[int(27)] = len(_dafny.Map({d_104_v55_: d_38_v0_}))
        d_105_v56_ = nw18_
        index11_ = default__.safeIndex(816, (d_105_v56_).length(0))
        (d_105_v56_)[index11_] = (0) - ((d_38_v0_) - (len(d_40_v2_)))
        d_108_v57_: C1
        nw19_ = C1()
        nw19_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")), len(d_40_v2_), (d_90_v41_).f25)
        d_108_v57_ = nw19_
        d_109_v58_: D1
        d_109_v58_ = D1_DC1(d_88_v39_)
        d_110_v59_: _dafny.Map
        d_110_v59_ = _dafny.Map({d_108_v57_: d_109_v58_})
        index12_ = default__.safeIndex(816, (d_105_v56_).length(0))
        (d_105_v56_)[index12_] = (0) - (len(((d_110_v59_).set(d_108_v57_, d_109_v58_)) | (d_110_v59_)))
        if not (True) or (d_88_v39_):
            d_111_v60_: _dafny.Map
            d_111_v60_ = _dafny.Map({d_108_v57_.f29: _dafny.Set({d_88_v39_, (d_90_v41_).f25, d_88_v39_, (d_90_v41_).f25})})
            d_112_v61_: _dafny.Seq
            d_112_v61_ = _dafny.SeqWithoutIsStrInference([(d_90_v41_).f25])
            d_113_v62_: _dafny.Map
            d_113_v62_ = _dafny.Map({((d_111_v60_)[d_108_v57_.f29] if (d_108_v57_.f29) in (d_111_v60_) else d_99_v50_): d_112_v61_})
            d_114_v63_: _dafny.Seq
            d_114_v63_ = _dafny.SeqWithoutIsStrInference([d_113_v62_])
            d_115_v64_: _dafny.Seq
            d_115_v64_ = _dafny.SeqWithoutIsStrInference([d_97_v47_])
            d_116_v66_: _dafny.Set
            d_116_v66_ = _dafny.Set({d_97_v47_})
            d_117_v67_: _dafny.Array
            nw20_ = _dafny.Array(None, 23)
            d_117_v67_ = nw20_
            d_118_v68_: D4
            d_118_v68_ = D4_DC13(d_117_v67_, d_88_v39_, (d_90_v41_).f25)
            d_119_v69_: _dafny.Map
            d_119_v69_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([38, len(d_39_v1_)]): (d_90_v41_).f25})
            d_120_v70_: _dafny.Array
            nw21_ = _dafny.Array(None, 29)
            nw21_[int(0)] = d_88_v39_
            nw21_[int(1)] = (d_40_v2_) <= (d_40_v2_)
            nw21_[int(2)] = (_dafny.Set({(d_90_v41_).f25, (d_90_v41_).f25})).issubset(d_99_v50_)
            nw21_[int(3)] = (d_90_v41_).f25
            nw21_[int(4)] = True
            nw21_[int(5)] = False
            nw21_[int(6)] = (d_90_v41_).f25
            nw21_[int(7)] = (d_99_v50_) in ((d_114_v63_)[default__.safeIndex(d_108_v57_.f29, len(d_114_v63_))])
            nw21_[int(8)] = (d_90_v41_).f25
            nw21_[int(9)] = d_88_v39_
            nw21_[int(10)] = (d_90_v41_).f25
            nw21_[int(11)] = (d_90_v41_).f25
            nw21_[int(12)] = (d_90_v41_).f25
            nw21_[int(13)] = ((d_105_v56_)[default__.safeIndex(816, (d_105_v56_).length(0))]) not in (d_93_v44_)
            nw21_[int(14)] = True
            nw21_[int(15)] = (d_90_v41_).f25
            nw21_[int(16)] = d_88_v39_
            nw21_[int(17)] = (d_90_v41_).f25
            nw21_[int(18)] = (d_90_v41_).f25
            nw21_[int(19)] = (d_90_v41_).f25
            nw21_[int(20)] = (d_90_v41_).f25
            nw21_[int(21)] = (d_90_v41_).f25
            def iife11_():
                coll7_ = _dafny.Set()
                compr_7_: _dafny.MultiSet
                for compr_7_ in (d_115_v64_).Elements:
                    d_121_v65_: _dafny.MultiSet = compr_7_
                    if (d_121_v65_) in (d_115_v64_):
                        coll7_ = coll7_.union(_dafny.Set([d_121_v65_]))
                return _dafny.Set(coll7_)
            nw21_[int(22)] = (iife11_()
            ).isdisjoint(d_116_v66_)
            nw21_[int(23)] = (d_118_v68_).cf22
            nw21_[int(24)] = (_dafny.MultiSet(d_39_v1_)).issubset(d_97_v47_)
            nw21_[int(25)] = not((len(d_119_v69_)) <= (91))
            nw21_[int(26)] = (d_108_v57_.f28) <= (d_40_v2_)
            nw21_[int(27)] = (d_108_v57_.f29) > ((d_105_v56_)[default__.safeIndex(816, (d_105_v56_).length(0))])
            nw21_[int(28)] = (d_38_v0_) >= ((d_105_v56_)[default__.safeIndex(816, (d_105_v56_).length(0))])
            d_120_v70_ = nw21_
            d_122_v71_: _dafny.Map
            d_122_v71_ = _dafny.Map({(d_97_v47_).intersection(d_97_v47_): d_104_v55_})
            rhs9_ = d_120_v70_
            rhs10_ = ((d_122_v71_)[d_97_v47_] if (d_97_v47_) in (d_122_v71_) else d_104_v55_)
            d_120_v70_ = rhs9_
            d_104_v55_ = rhs10_
            d_123_v72_: D5
            d_123_v72_ = D5_DC15((d_90_v41_).f25, d_120_v70_, d_108_v57_.f28, d_88_v39_)
            d_124_v73_: D5
            d_124_v73_ = D5_DC16(d_123_v72_)
            d_125_v74_: _dafny.Map
            d_125_v74_ = _dafny.Map({(d_38_v0_) < (d_108_v57_.f29): D5_DC16(d_123_v72_)})
            d_126_v75_: D5
            d_126_v75_ = D5_DC16(d_123_v72_)
            d_125_v74_ = (d_125_v74_).set(d_88_v39_, d_126_v75_)
            d_127_v76_: T1
            nw22_ = C3()
            nw22_.ctor__(d_89_v40_, d_95_v46_, (d_90_v41_).f25)
            d_127_v76_ = nw22_
            d_128_v77_: _dafny.Seq
            d_128_v77_ = _dafny.SeqWithoutIsStrInference([d_127_v76_])
            d_128_v77_ = (d_128_v77_) + ((d_128_v77_) + (d_128_v77_))
            d_129_v78_: C0
            nw23_ = C0()
            nw23_.ctor__()
            d_129_v78_ = nw23_
            (globalState).f18 = d_108_v57_.f29
        elif True:
            (globalState).f22 = (0) - (d_108_v57_.f29)
            d_130_v79_: _dafny.Map
            d_130_v79_ = _dafny.Map({-974: d_90_v41_})
            if (d_130_v79_) == (_dafny.Map({-283: d_90_v41_})):
                d_131_v80_: _dafny.Array
                nw24_ = _dafny.Array(False, 13)
                d_131_v80_ = nw24_
                index13_ = default__.safeIndex(850, (d_131_v80_).length(0))
                (d_131_v80_)[index13_] = (d_108_v57_.f28) <= (d_40_v2_)
                index14_ = default__.safeIndex(850, (d_131_v80_).length(0))
                (d_131_v80_)[index14_] = d_88_v39_
                index15_ = default__.safeIndex(850, (d_131_v80_).length(0))
                rhs11_ = (d_108_v57_.f28) != (d_108_v57_.f28)
                rhs12_ = d_104_v55_
                lhs7_ = d_131_v80_
                lhs8_ = default__.safeIndex(850, (d_131_v80_).length(0))
                lhs7_[lhs8_] = rhs11_
                d_104_v55_ = rhs12_
                (globalState).f2 = (((d_105_v56_)[default__.safeIndex(816, (d_105_v56_).length(0))]) + (739)) == (d_108_v57_.f29)
                d_97_v47_ = (_dafny.MultiSet([d_108_v57_.f29])).intersection((d_97_v47_) - ((d_97_v47_).set((d_105_v56_)[default__.safeIndex(816, (d_105_v56_).length(0))], default__.abs((d_105_v56_)[default__.safeIndex(816, (d_105_v56_).length(0))]))))
                (globalState).f2 = (d_90_v41_).f25
            elif True:
                index16_ = default__.safeIndex(816, (d_105_v56_).length(0))
                rhs13_ = (default__.safeModuloInt((d_105_v56_)[default__.safeIndex(816, (d_105_v56_).length(0))], d_108_v57_.f29)) * (d_108_v57_.f29)
                rhs14_ = (d_105_v56_)[default__.safeIndex(816, (d_105_v56_).length(0))]
                lhs9_ = d_105_v56_
                lhs10_ = default__.safeIndex(816, (d_105_v56_).length(0))
                lhs11_ = globalState
                lhs9_[lhs10_] = rhs13_
                lhs11_.f18 = rhs14_
                (globalState).f2 = not((d_90_v41_).f25)
                d_132_v81_: C2
                nw25_ = C2()
                nw25_.ctor__(d_89_v40_, d_95_v46_, not((d_90_v41_).f25))
                d_132_v81_ = nw25_
                d_133_v82_: _dafny.Seq
                d_133_v82_ = _dafny.SeqWithoutIsStrInference([d_132_v81_])
                d_134_v83_: _dafny.Map
                d_134_v83_ = _dafny.Map({default__.fm18(len(d_133_v82_), d_108_v57_.f29, globalState): d_89_v40_})
                d_135_v84_: T1
                nw26_ = C3()
                nw26_.ctor__(((d_134_v83_)[d_108_v57_.f28] if (d_108_v57_.f28) in (d_134_v83_) else d_89_v40_), d_95_v46_, not(True))
                d_135_v84_ = nw26_
                d_136_v85_: _dafny.Map
                d_136_v85_ = _dafny.Map({d_105_v56_: d_135_v84_})
                d_137_v86_: _dafny.Map
                d_137_v86_ = _dafny.Map({False: ((d_136_v85_)[d_105_v56_] if (d_105_v56_) in (d_136_v85_) else d_135_v84_)})
                d_135_v84_ = ((d_137_v86_)[True] if (True) in (d_137_v86_) else d_135_v84_)
                d_138_v87_: _dafny.Array
                nw27_ = _dafny.Array(False, 17)
                d_138_v87_ = nw27_
                index17_ = default__.safeIndex(169, (d_138_v87_).length(0))
                (d_138_v87_)[index17_] = (d_90_v41_).f25
                index18_ = default__.safeIndex(169, (d_138_v87_).length(0))
                (d_138_v87_)[index18_] = (d_135_v84_).f25
                d_139_v88_: _dafny.Map
                d_139_v88_ = _dafny.Map({d_41_v3_: d_98_v49_})
                d_140_v90_: _dafny.Map
                d_140_v90_ = _dafny.Map({(d_135_v84_).f26: d_108_v57_.f29})
                d_141_v91_: _dafny.MultiSet
                d_141_v91_ = _dafny.MultiSet([(d_90_v41_).f25, (d_138_v87_)[default__.safeIndex(169, (d_138_v87_).length(0))], (d_90_v41_).f25, (d_90_v41_).f25, (d_90_v41_).f25])
                def iife12_():
                    coll8_ = _dafny.Set()
                    compr_8_: int
                    for compr_8_ in (d_41_v3_).Elements:
                        d_142_v89_: int = compr_8_
                        if (d_142_v89_) in (d_41_v3_):
                            coll8_ = coll8_.union(_dafny.Set([(d_142_v89_) + ((_dafny.MultiSet([313])).cardinality)]))
                    return _dafny.Set(coll8_)
                d_139_v88_ = (_dafny.Map({iife12_()
                : _dafny.Map({-325: (d_138_v87_)[default__.safeIndex(169, (d_138_v87_).length(0))]})})) | ((default__.fm28((d_104_v55_).fm3((d_138_v87_)[default__.safeIndex(169, (d_138_v87_).length(0))], 819, len(_dafny.Map({d_108_v57_.f29: (d_105_v56_)[default__.safeIndex(816, (d_105_v56_).length(0))]})), d_140_v90_, globalState), ((d_141_v91_)[(d_90_v41_).f25] if ((d_90_v41_).f25) in (d_141_v91_) else d_108_v57_.f29), 342, globalState)).set(d_41_v3_, d_98_v49_))
            (globalState).f2 = (d_90_v41_).f25
            (globalState).f22 = d_38_v0_
            d_143_v92_: _dafny.Seq
            d_143_v92_ = _dafny.SeqWithoutIsStrInference([(d_90_v41_).f25])
            r0 = default__.safeModuloInt((len(d_143_v92_)) + ((d_105_v56_)[default__.safeIndex(816, (d_105_v56_).length(0))]), d_108_v57_.f29)
        d_144_v93_: _dafny.MultiSet
        d_144_v93_ = _dafny.MultiSet([d_88_v39_])
        r0 = (((d_105_v56_)[default__.safeIndex(816, (d_105_v56_).length(0))]) - (((d_144_v93_).set((d_90_v41_).f25, default__.abs(629))).cardinality)) * (len(d_41_v3_))
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_145_v0_: _dafny.Seq
        d_145_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qyvk"))
        d_146_v1_: _dafny.Array
        def lambda6_(d_147_v0_):
            def lambda7_(d_148_i0_):
                return _dafny.Set({len(d_147_v0_), len(d_147_v0_)})

            return lambda7_

        init3_ = lambda6_(d_145_v0_)
        nw28_ = _dafny.Array(None, 11)
        for i0_3_ in range(nw28_.length(0)):
            nw28_[i0_3_] = init3_(i0_3_)
        d_146_v1_ = nw28_
        d_149_v2_: _dafny.Array
        nw29_ = _dafny.Array(_dafny.Array(None, 0), 19)
        d_149_v2_ = nw29_
        d_150_v3_: bool
        d_150_v3_ = True
        d_151_v4_: _dafny.Set
        d_151_v4_ = _dafny.Set({d_150_v3_})
        d_152_v5_: _dafny.Seq
        d_152_v5_ = _dafny.SeqWithoutIsStrInference([len(d_151_v4_)])
        d_153_v6_: _dafny.Array
        nw30_ = _dafny.Array(False, 20)
        d_153_v6_ = nw30_
        d_154_globalState_: GlobalState
        nw31_ = GlobalState()
        nw31_.ctor__(460, 238, True, d_145_v0_, d_146_v1_, 698, False, 939, False, d_149_v2_, 659, False, _dafny.CodePoint('w'), False, False, 856, 966, False, 806, d_152_v5_, 962, 561, 182, d_153_v6_, _dafny.CodePoint('w'))
        d_154_globalState_ = nw31_
        d_155_v7_: int
        out0_: int
        out0_ = default__.m0(d_154_globalState_)
        d_155_v7_ = out0_
        d_156_v8_: _dafny.MultiSet
        d_156_v8_ = _dafny.MultiSet([d_145_v0_])
        d_157_v9_: _dafny.MultiSet
        d_157_v9_ = _dafny.MultiSet([d_150_v3_, not((d_156_v8_) != (d_156_v8_)), not (not(d_150_v3_)) or (default__.fm0(d_155_v7_, d_155_v7_, d_150_v3_, d_154_globalState_)), d_150_v3_, d_150_v3_])
        d_158_v10_: _dafny.Seq
        d_158_v10_ = _dafny.SeqWithoutIsStrInference([d_150_v3_, d_150_v3_])
        d_157_v9_ = default__.fm1(d_155_v7_, len((d_158_v10_) + (d_158_v10_)), d_154_globalState_)
        rhs15_ = (d_155_v7_) * (d_155_v7_)
        rhs16_ = not(d_150_v3_)
        rhs17_ = d_153_v6_
        lhs12_ = d_154_globalState_
        lhs13_ = d_154_globalState_
        lhs12_.f22 = rhs15_
        lhs13_.f2 = rhs16_
        d_153_v6_ = rhs17_
        d_159_v11_: str
        d_159_v11_ = _dafny.CodePoint('k')
        d_159_v11_ = d_159_v11_
        d_160_v12_: _dafny.Map
        d_160_v12_ = _dafny.Map({default__.fm2(d_150_v3_, d_155_v7_, d_150_v3_, d_159_v11_, d_154_globalState_): d_150_v3_})
        d_161_v13_: _dafny.Seq
        d_161_v13_ = _dafny.SeqWithoutIsStrInference([d_157_v9_])
        d_162_v14_: _dafny.Array
        nw32_ = _dafny.Array(None, 21)
        nw32_[int(0)] = (_dafny.MultiSet([((d_160_v12_)[926] if (926) in (d_160_v12_) else d_150_v3_)])) | (d_157_v9_)
        nw32_[int(1)] = d_157_v9_
        nw32_[int(2)] = d_157_v9_
        nw32_[int(3)] = d_157_v9_
        nw32_[int(4)] = default__.fm1(d_155_v7_, d_155_v7_, d_154_globalState_)
        nw32_[int(5)] = _dafny.MultiSet([True])
        nw32_[int(6)] = d_157_v9_
        nw32_[int(7)] = (d_161_v13_)[default__.safeIndex(d_155_v7_, len(d_161_v13_))]
        nw32_[int(8)] = default__.fm1(d_155_v7_, d_155_v7_, d_154_globalState_)
        nw32_[int(9)] = d_157_v9_
        nw32_[int(10)] = (d_157_v9_).intersection(default__.fm1(d_155_v7_, d_155_v7_, d_154_globalState_))
        nw32_[int(11)] = d_157_v9_
        nw32_[int(12)] = default__.fm1(-362, d_155_v7_, d_154_globalState_)
        nw32_[int(13)] = _dafny.MultiSet(d_158_v10_)
        nw32_[int(14)] = d_157_v9_
        nw32_[int(15)] = d_157_v9_
        nw32_[int(16)] = (_dafny.MultiSet([d_150_v3_])).intersection(d_157_v9_)
        nw32_[int(17)] = d_157_v9_
        nw32_[int(18)] = _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([d_150_v3_])).set(default__.safeIndex(d_155_v7_, len(_dafny.SeqWithoutIsStrInference([d_150_v3_]))), d_150_v3_)).set(default__.safeIndex(317, len((_dafny.SeqWithoutIsStrInference([d_150_v3_])).set(default__.safeIndex(d_155_v7_, len(_dafny.SeqWithoutIsStrInference([d_150_v3_]))), d_150_v3_))), d_150_v3_))
        nw32_[int(19)] = d_157_v9_
        nw32_[int(20)] = (d_157_v9_).intersection(d_157_v9_)
        d_162_v14_ = nw32_
        index19_ = default__.safeIndex(115, (d_162_v14_).length(0))
        (d_162_v14_)[index19_] = d_157_v9_
        index20_ = default__.safeIndex(939, (d_153_v6_).length(0))
        (d_153_v6_)[index20_] = True
        d_163_v15_: _dafny.Map
        d_163_v15_ = _dafny.Map({not((default__.fm2(d_150_v3_, d_155_v7_, not(d_150_v3_), d_159_v11_, d_154_globalState_)) > (len(d_145_v0_))): _dafny.MultiSet(d_158_v10_)})
        index21_ = default__.safeIndex(115, (d_162_v14_).length(0))
        index22_ = default__.safeIndex(939, (d_153_v6_).length(0))
        rhs18_ = d_155_v7_
        rhs19_ = False
        rhs20_ = ((d_163_v15_)[d_150_v3_] if (d_150_v3_) in (d_163_v15_) else d_157_v9_)
        rhs21_ = (d_155_v7_) >= (len(_dafny.SeqWithoutIsStrInference([default__.fm2(d_150_v3_, len((d_145_v0_).set(default__.safeIndex(d_155_v7_, len(d_145_v0_)), d_159_v11_)), d_150_v3_, d_159_v11_, d_154_globalState_)])))
        rhs22_ = d_155_v7_
        lhs14_ = d_162_v14_
        lhs15_ = default__.safeIndex(115, (d_162_v14_).length(0))
        lhs16_ = d_153_v6_
        lhs17_ = default__.safeIndex(939, (d_153_v6_).length(0))
        d_155_v7_ = rhs18_
        d_150_v3_ = rhs19_
        lhs14_[lhs15_] = rhs20_
        lhs16_[lhs17_] = rhs21_
        d_155_v7_ = rhs22_
        d_164_v16_: int
        out1_: int
        out1_ = default__.m0(d_154_globalState_)
        d_164_v16_ = out1_
        d_165_v17_: _dafny.Map
        d_165_v17_ = _dafny.Map({(d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))]: (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))]})
        d_166_v18_: _dafny.MultiSet
        d_166_v18_ = _dafny.MultiSet([len(d_165_v17_)])
        d_167_v19_: _dafny.Set
        d_167_v19_ = _dafny.Set({d_164_v16_, d_155_v7_})
        d_168_v20_: _dafny.Array
        nw33_ = _dafny.Array(None, 19)
        nw33_[int(0)] = d_164_v16_
        nw33_[int(1)] = d_164_v16_
        nw33_[int(2)] = d_164_v16_
        nw33_[int(3)] = d_164_v16_
        nw33_[int(4)] = len(d_145_v0_)
        nw33_[int(5)] = d_164_v16_
        nw33_[int(6)] = 757
        nw33_[int(7)] = d_164_v16_
        nw33_[int(8)] = d_164_v16_
        nw33_[int(9)] = default__.fm2((d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))], d_155_v7_, d_150_v3_, d_159_v11_, d_154_globalState_)
        nw33_[int(10)] = (d_166_v18_).cardinality
        nw33_[int(11)] = d_164_v16_
        nw33_[int(12)] = len(d_167_v19_)
        nw33_[int(13)] = d_155_v7_
        nw33_[int(14)] = default__.fm2(d_150_v3_, d_164_v16_, d_150_v3_, _dafny.CodePoint('w'), d_154_globalState_)
        nw33_[int(15)] = d_164_v16_
        nw33_[int(16)] = (0) - (d_155_v7_)
        nw33_[int(17)] = d_155_v7_
        nw33_[int(18)] = default__.fm2((d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))], len(d_151_v4_), d_150_v3_, d_159_v11_, d_154_globalState_)
        d_168_v20_ = nw33_
        hi1_ = d_164_v16_
        for d_169_i1_ in range(len((_dafny.Map({d_152_v5_: d_168_v20_})).set(d_152_v5_, d_168_v20_)), hi1_):
            d_170_v21_: _dafny.Array
            nw34_ = _dafny.Array(_dafny.Map({}), 17)
            d_170_v21_ = nw34_
            index23_ = default__.safeIndex(892, (d_170_v21_).length(0))
            (d_170_v21_)[index23_] = d_160_v12_
            index24_ = default__.safeIndex(115, (d_162_v14_).length(0))
            index25_ = default__.safeIndex(892, (d_170_v21_).length(0))
            rhs23_ = _dafny.MultiSet([(d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))], (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))]])
            rhs24_ = d_160_v12_
            lhs18_ = d_162_v14_
            lhs19_ = default__.safeIndex(115, (d_162_v14_).length(0))
            lhs20_ = d_170_v21_
            lhs21_ = default__.safeIndex(892, (d_170_v21_).length(0))
            lhs18_[lhs19_] = rhs23_
            lhs20_[lhs21_] = rhs24_
            d_145_v0_ = ((d_145_v0_ if (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))] else d_145_v0_)).set(default__.safeIndex(default__.safeDivisionInt(d_169_i1_, d_164_v16_), len((d_145_v0_ if (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))] else d_145_v0_))), _dafny.CodePoint('n'))
            d_171_v22_: int
            out2_: int
            out2_ = default__.m0(d_154_globalState_)
            d_171_v22_ = out2_
            d_172_v23_: C0
            nw35_ = C0()
            nw35_.ctor__()
            d_172_v23_ = nw35_
        d_173_v24_: int
        out3_: int
        out3_ = default__.m0(d_154_globalState_)
        d_173_v24_ = out3_
        d_173_v24_ = d_164_v16_
        (d_154_globalState_).f2 = (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))]
        index26_ = default__.safeIndex(793, (d_168_v20_).length(0))
        (d_168_v20_)[index26_] = d_155_v7_
        index27_ = default__.safeIndex(793, (d_168_v20_).length(0))
        (d_168_v20_)[index27_] = 967
        def iife13_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(670, 941):
                d_174_v25_: int = compr_9_
                if ((670) <= (d_174_v25_)) and ((d_174_v25_) < (941)):
                    coll9_[default__.safeDivisionInt(d_174_v25_, 913)] = d_150_v3_
            return _dafny.Map(coll9_)
        (d_154_globalState_).f5 = default__.safeDivisionInt(default__.fm2(((d_165_v17_)[True] if (True) in (d_165_v17_) else not(d_150_v3_)), 317, (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))], d_159_v11_, d_154_globalState_), (0) - (len((iife13_()
        ) | (default__.fm26((d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))], ((d_157_v9_)[d_150_v3_] if (d_150_v3_) in (d_157_v9_) else d_164_v16_), len(d_145_v0_), d_154_globalState_)))))
        d_175_v26_: _dafny.Array
        def lambda8_(d_176_v18_):
            def lambda9_(d_177_i2_):
                return d_176_v18_

            return lambda9_

        init4_ = lambda8_(d_166_v18_)
        nw36_ = _dafny.Array(None, 14)
        for i0_4_ in range(nw36_.length(0)):
            nw36_[i0_4_] = init4_(i0_4_)
        d_175_v26_ = nw36_
        d_178_v27_: C3
        nw37_ = C3()
        nw37_.ctor__(d_159_v11_, d_175_v26_, ((d_168_v20_)[default__.safeIndex(793, (d_168_v20_).length(0))]) == (d_155_v7_))
        d_178_v27_ = nw37_
        d_179_v28_: D7
        d_179_v28_ = D7_DC23(d_173_v24_, 294, d_150_v3_, d_155_v7_)
        d_180_v29_: D7
        d_180_v29_ = D7_DC24(d_179_v28_)
        d_181_i3_: int
        d_181_i3_ = 0
        with _dafny.label("0"):
            pat_let_tv1_ = d_150_v3_
            pat_let_tv2_ = d_150_v3_
            pat_let_tv3_ = d_159_v11_
            pat_let_tv4_ = d_153_v6_
            pat_let_tv5_ = d_153_v6_
            pat_let_tv6_ = d_145_v0_
            def lambda10_(source2_):
                if source2_.is_DC21:
                    d_185___mcc_h0_ = source2_.cf38
                    d_186___mcc_h1_ = source2_.cf39
                    d_187_cf39_ = d_186___mcc_h1_
                    d_188_cf38_ = d_185___mcc_h0_
                    return pat_let_tv1_
                elif source2_.is_DC22:
                    d_189___mcc_h2_ = source2_.cf40
                    d_190___mcc_h3_ = source2_.cf41
                    d_191___mcc_h4_ = source2_.cf42
                    d_192_cf42_ = d_191___mcc_h4_
                    d_193_cf41_ = d_190___mcc_h3_
                    d_194_cf40_ = d_189___mcc_h2_
                    return pat_let_tv2_
                elif source2_.is_DC23:
                    d_195___mcc_h5_ = source2_.cf43
                    d_196___mcc_h6_ = source2_.cf44
                    d_197___mcc_h7_ = source2_.cf45
                    d_198___mcc_h8_ = source2_.cf46
                    d_199_cf46_ = d_198___mcc_h8_
                    d_200_cf45_ = d_197___mcc_h7_
                    d_201_cf44_ = d_196___mcc_h6_
                    d_202_cf43_ = d_195___mcc_h5_
                    return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tdneecvnd"))) <= (_dafny.SeqWithoutIsStrInference([pat_let_tv3_ for d_203_i4_ in range(default__.abs(-460))]))
                elif source2_.is_DC20:
                    d_204___mcc_h9_ = source2_.cf37
                    d_205_cf37_ = d_204___mcc_h9_
                    return (pat_let_tv5_)[default__.safeIndex(939, (pat_let_tv4_).length(0))]
                elif True:
                    d_206___mcc_h10_ = source2_.cf47
                    d_207_cf47_ = d_206___mcc_h10_
                    return (pat_let_tv6_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qarhghwh")))

            while lambda10_(d_180_v29_):
                with _dafny.c_label("0"):
                    if (d_181_i3_) >= (100):
                        raise _dafny.Break("0")
                    d_181_i3_ = (d_181_i3_) + (1)
                    index28_ = default__.safeIndex(939, (d_153_v6_).length(0))
                    (d_153_v6_)[index28_] = d_150_v3_
                    d_182_v30_: int
                    out4_: int
                    out4_ = default__.m0(d_154_globalState_)
                    d_182_v30_ = out4_
                    d_145_v0_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mukobu"))) + (d_145_v0_)
                    d_183_v31_: _dafny.Set
                    d_183_v31_ = _dafny.Set({d_159_v11_})
                    (d_178_v27_).m2((d_183_v31_ if (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))] else _dafny.Set({d_159_v11_, d_159_v11_, _dafny.CodePoint('o'), d_159_v11_, d_159_v11_})), d_150_v3_, len(_dafny.SeqWithoutIsStrInference([d_145_v0_ for d_184_i5_ in range(default__.abs(-712))])), d_154_globalState_)
                    pass
            pass
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_153_v6_).length(0)):
            d_208_i6_: int = guard_loop_1_
            if (True) and (((0) <= (d_208_i6_)) and ((d_208_i6_) < ((d_153_v6_).length(0)))):
                (d_153_v6_)[(d_208_i6_)] = d_150_v3_
        d_209_v32_: _dafny.Seq
        d_209_v32_ = _dafny.SeqWithoutIsStrInference([d_153_v6_])
        d_210_v33_: D3
        d_210_v33_ = D3_DC9(len(d_209_v32_), d_159_v11_, 326)
        d_211_v34_: _dafny.Map
        d_211_v34_ = _dafny.Map({(d_210_v33_).cf14: d_164_v16_})
        if (d_178_v27_).fm3(False, d_173_v24_, default__.fm2(d_150_v3_, d_164_v16_, (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))], d_159_v11_, d_154_globalState_), (d_211_v34_) | (d_211_v34_), d_154_globalState_):
            (d_154_globalState_).f5 = (d_168_v20_)[default__.safeIndex(793, (d_168_v20_).length(0))]
            (d_154_globalState_).f22 = ((d_166_v18_)[d_173_v24_] if (d_173_v24_) in (d_166_v18_) else d_173_v24_)
            d_209_v32_ = d_209_v32_
            d_212_v35_: _dafny.Array
            nw38_ = _dafny.Array(D3.default()(), 24)
            d_212_v35_ = nw38_
            d_213_v36_: _dafny.Map
            d_213_v36_ = _dafny.Map({d_212_v35_: (d_158_v10_) + (d_158_v10_)})
            rhs25_ = d_164_v16_
            rhs26_ = _dafny.Map({d_212_v35_: (d_158_v10_) + (_dafny.SeqWithoutIsStrInference([d_150_v3_]))})
            rhs27_ = (d_166_v18_).intersection((d_166_v18_) - (_dafny.MultiSet(d_152_v5_)))
            d_164_v16_ = rhs25_
            d_213_v36_ = rhs26_
            d_166_v18_ = rhs27_
            (d_154_globalState_).f2 = (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))]
        elif True:
            (d_154_globalState_).f2 = (d_166_v18_) != (d_166_v18_)
            (d_154_globalState_).f2 = not((d_151_v4_).isdisjoint(d_151_v4_))
            d_150_v3_ = (d_178_v27_).fm3(d_150_v3_, 971, (d_168_v20_)[default__.safeIndex(793, (d_168_v20_).length(0))], d_211_v34_, d_154_globalState_)
            d_214_v37_: D8
            d_214_v37_ = D8_DC26(d_168_v20_, (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))], (d_145_v0_)[default__.safeIndex((d_168_v20_)[default__.safeIndex(793, (d_168_v20_).length(0))], len(d_145_v0_))], (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))], d_150_v3_)
            d_215_v38_: _dafny.Seq
            d_215_v38_ = _dafny.SeqWithoutIsStrInference([d_214_v37_])
            d_215_v38_ = (d_215_v38_ if d_150_v3_ else (d_215_v38_ if (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))] else d_215_v38_))
            d_216_v39_: _dafny.Map
            d_216_v39_ = _dafny.Map({d_155_v7_: d_158_v10_})
            d_216_v39_ = (d_216_v39_).set(d_173_v24_, _dafny.SeqWithoutIsStrInference([default__.fm0(d_155_v7_, 512, (d_153_v6_)[default__.safeIndex(939, (d_153_v6_).length(0))], d_154_globalState_)]))
        _dafny.print((d_145_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v1_)[0]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v1_)[1]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v1_)[2]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v1_)[3]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v1_)[4]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v1_)[5]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v1_)[6]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v1_)[7]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v1_)[8]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v1_)[9]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v1_)[10]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_150_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v4_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v5_) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v6_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_154_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_154_globalState_).f3).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_154_globalState_).f4)[0]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_154_globalState_).f4)[1]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_154_globalState_).f4)[2]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_154_globalState_).f4)[3]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_154_globalState_).f4)[4]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_154_globalState_).f4)[5]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_154_globalState_).f4)[6]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_154_globalState_).f4)[7]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_154_globalState_).f4)[8]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_154_globalState_).f4)[9]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_154_globalState_).f4)[10]) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_154_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_154_globalState_.f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_.f19) == (_dafny.SeqWithoutIsStrInference([142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_154_globalState_.f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_154_globalState_.f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_154_globalState_).f23)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_globalState_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_155_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_v8_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qyvk"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v9_) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v10_) == (_dafny.SeqWithoutIsStrInference([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_159_v11_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v12_) == (_dafny.Map({-386: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_v13_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[0]) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[1]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[2]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[3]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[4]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[5]) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[6]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[7]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[8]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[9]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[10]) == (_dafny.MultiSet([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[11]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[12]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[13]) == (_dafny.MultiSet([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[14]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[15]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[16]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[17]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[18]) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[19]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v14_)[20]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_163_v15_) == (_dafny.Map({True: _dafny.MultiSet([True, True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_164_v16_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v17_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v18_) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_167_v19_) == (_dafny.Set({-599}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v20_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_173_v24_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[0]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[1]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[2]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[3]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[4]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[5]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[6]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[7]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[8]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[9]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[10]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[11]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[12]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v26_)[13]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_179_v28_).cf43))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_179_v28_).cf44))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_179_v28_).cf45))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_179_v28_).cf46))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v29_).cf47).cf43))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v29_).cf47).cf44))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v29_).cf47).cf45))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v29_).cf47).cf46))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_181_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_209_v32_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_210_v33_).cf13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_210_v33_).cf14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_210_v33_).cf15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v34_) == (_dafny.Map({_dafny.CodePoint('k'): -599}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: int(0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

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
        return lambda: D1_DC2(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC4(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC4(D2, NamedTuple('DC4', [('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf6', Any), ('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC8(D3, NamedTuple('DC8', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC12(False, _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC12(D4, NamedTuple('DC12', [('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({self.cf17.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC15(False, _dafny.Array(None, 0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)

class D5_DC15(D5, NamedTuple('DC15', [('cf25', Any), ('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {self.cf27.VerbatimString(True)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC18(int(0), False, False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)

class D6_DC18(D6, NamedTuple('DC18', [('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC21(int(0), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D7_DC23)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D7_DC24)

class D7_DC21(D7, NamedTuple('DC21', [('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC22(D7, NamedTuple('DC22', [('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC23(D7, NamedTuple('DC23', [('cf43', Any), ('cf44', Any), ('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC23({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC23) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC24(D7, NamedTuple('DC24', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC24({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC24) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC26(_dafny.Array(None, 0), False, _dafny.CodePoint('D'), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D8_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)

class D8_DC26(D8, NamedTuple('DC26', [('cf49', Any), ('cf50', Any), ('cf51', Any), ('cf52', Any), ('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC26({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC26) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, p1, p2, globalState):
        pass


class T1:
    pass
    def m2(self, p0, p1, p2, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: bool = False
        self.f5: int = int(0)
        self.f18: int = int(0)
        self.f19: _dafny.Seq = _dafny.Seq({})
        self.f21: int = int(0)
        self.f22: int = int(0)
        self._f0: int = int(0)
        self._f1: int = int(0)
        self._f3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f6: bool = False
        self._f7: int = int(0)
        self._f8: bool = False
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        self._f10: int = int(0)
        self._f11: bool = False
        self._f12: str = _dafny.CodePoint('D')
        self._f13: bool = False
        self._f14: bool = False
        self._f15: int = int(0)
        self._f16: int = int(0)
        self._f17: bool = False
        self._f20: int = int(0)
        self._f23: _dafny.Array = _dafny.Array(None, 0)
        self._f24: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self).f18 = f18
        (self).f19 = f19
        (self)._f20 = f20
        (self).f21 = f21
        (self).f22 = f22
        (self)._f23 = f23
        (self)._f24 = f24

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17
    @property
    def f20(self):
        return self._f20
    @property
    def f23(self):
        return self._f23
    @property
    def f24(self):
        return self._f24

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm5(self, globalState):
        return (0) - (((-144 if True else (0) - (len(_dafny.Map({610: True}))))) + (-131))

    def fm6(self, p0, p1, globalState):
        return (len(_dafny.SeqWithoutIsStrInference([False]))) + (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "roa"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))))


class C1(T0):
    def  __init__(self):
        self._f25: bool = False
        self.f28: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f29: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f28, f29, f25):
        (self).f28 = f28
        (self).f29 = f29
        (self)._f25 = f25

    def fm14(self, p0, p1, globalState):
        return default__.safeModuloInt((0) - (((_dafny.MultiSet([(self).f25, (self).f25, (self).f25, False])) | (_dafny.MultiSet([(self).f25]))).cardinality), self.f29)

    def fm15(self, globalState):
        def iife14_():
            coll10_ = _dafny.Set()
            compr_10_: int
            for compr_10_ in (_dafny.MultiSet([self.f29])).Elements:
                d_217_v0_: int = compr_10_
                if (d_217_v0_) in (_dafny.MultiSet([self.f29])):
                    def iife15_():
                        coll11_ = _dafny.Map()
                        compr_11_: _dafny.Seq
                        for compr_11_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): _dafny.CodePoint('j')})).keys.Elements:
                            d_218_v1_: _dafny.Seq = compr_11_
                            if (d_218_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): _dafny.CodePoint('j')})):
                                coll11_[d_218_v1_] = 873
                        return _dafny.Map(coll11_)
                    coll10_ = coll10_.union(_dafny.Set([default__.safeModuloInt(d_217_v0_, len(_dafny.Map({-186: len(iife15_()
)})))]))
            return _dafny.Set(coll10_)
        return _dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference([self.f29])) + (_dafny.SeqWithoutIsStrInference([len(iife14_()
)]))) for d_219_i0_ in range(default__.abs(4))])

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: int = int(0)
        d_220_v0_: D3
        d_220_v0_ = D3_DC7(_dafny.Map({self.f29: (self).f25}))
        d_221_v2_: _dafny.Map
        d_221_v2_ = _dafny.Map({p1: (self).f25})
        d_222_v3_: _dafny.Map
        d_222_v3_ = _dafny.Map({(self).f25: False})
        d_223_v4_: _dafny.Array
        nw39_ = _dafny.Array(None, 14)
        nw39_[int(0)] = (self).f25
        nw39_[int(1)] = (self).f25
        nw39_[int(2)] = (self).f25
        nw39_[int(3)] = (self).f25
        nw39_[int(4)] = (self).f25
        nw39_[int(5)] = False
        nw39_[int(6)] = (self).f25
        nw39_[int(7)] = (self).f25
        nw39_[int(8)] = (self).f25
        nw39_[int(9)] = (p2)[default__.safeIndex(p1, len(p2))]
        nw39_[int(10)] = ((d_222_v3_)[(self).f25] if ((self).f25) in (d_222_v3_) else (self).f25)
        nw39_[int(11)] = (self).f25
        nw39_[int(12)] = (self).f25
        nw39_[int(13)] = (self).f25
        d_223_v4_ = nw39_
        d_224_v5_: _dafny.Map
        d_224_v5_ = _dafny.Map({d_223_v4_: d_221_v2_})
        d_225_v7_: _dafny.Map
        d_225_v7_ = _dafny.Map({270: p1})
        d_226_v8_: _dafny.Seq
        def iife16_():
            coll12_ = _dafny.Map()
            compr_12_: int
            for compr_12_ in (d_225_v7_).keys.Elements:
                d_227_v6_: int = compr_12_
                if (d_227_v6_) in (d_225_v7_):
                    coll12_[default__.safeDivisionInt(d_227_v6_, self.f29)] = (self).f25
            return _dafny.Map(coll12_)
        d_226_v8_ = _dafny.SeqWithoutIsStrInference([iife16_()
        ])
        d_228_v9_: _dafny.Array
        nw40_ = _dafny.Array(None, 21)
        nw40_[int(0)] = (d_220_v0_).cf11
        def iife17_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in _dafny.IntegerRange(148, 557):
                d_229_v1_: int = compr_13_
                if ((148) <= (d_229_v1_)) and ((d_229_v1_) < (557)):
                    coll13_[(d_229_v1_) - (self.f29)] = (self).f25
            return _dafny.Map(coll13_)
        nw40_[int(1)] = iife17_()
        
        nw40_[int(2)] = (d_221_v2_) | (((d_224_v5_)[d_223_v4_] if (d_223_v4_) in (d_224_v5_) else d_221_v2_))
        nw40_[int(3)] = d_221_v2_
        nw40_[int(4)] = d_221_v2_
        nw40_[int(5)] = (d_221_v2_ if True else d_221_v2_)
        nw40_[int(6)] = (d_226_v8_)[default__.safeIndex(-530, len(d_226_v8_))]
        nw40_[int(7)] = d_221_v2_
        nw40_[int(8)] = d_221_v2_
        nw40_[int(9)] = d_221_v2_
        nw40_[int(10)] = d_221_v2_
        nw40_[int(11)] = d_221_v2_
        nw40_[int(12)] = _dafny.Map({len(d_222_v3_): (self).f25})
        nw40_[int(13)] = d_221_v2_
        nw40_[int(14)] = d_221_v2_
        nw40_[int(15)] = (d_221_v2_) | (_dafny.Map({self.f29: (self).f25}))
        nw40_[int(16)] = d_221_v2_
        nw40_[int(17)] = (d_221_v2_) | (d_221_v2_)
        nw40_[int(18)] = d_221_v2_
        nw40_[int(19)] = d_221_v2_
        nw40_[int(20)] = d_221_v2_
        d_228_v9_ = nw40_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_228_v9_).length(0)):
            d_230_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_230_i0_)) and ((d_230_i0_) < ((d_228_v9_).length(0)))):
                (d_228_v9_)[(d_230_i0_)] = d_221_v2_
        if (self).f25:
            d_231_v10_: _dafny.Seq
            d_231_v10_ = _dafny.SeqWithoutIsStrInference([(len(self.f28)) + (self.f29)])
            d_232_v11_: str
            d_232_v11_ = _dafny.CodePoint('f')
            d_233_v12_: _dafny.MultiSet
            d_233_v12_ = _dafny.MultiSet([d_232_v11_, d_232_v11_])
            (self).f29 = (d_231_v10_)[default__.safeIndex((d_233_v12_).cardinality, len(d_231_v10_))]
            (self).f28 = ((_dafny.SeqWithoutIsStrInference([d_232_v11_ for d_234_i1_ in range(default__.abs(107))])) + (_dafny.SeqWithoutIsStrInference([d_232_v11_]))).set(default__.safeIndex(default__.safeDivisionInt(p1, self.f29), len((_dafny.SeqWithoutIsStrInference([d_232_v11_ for d_235_i1_ in range(default__.abs(107))])) + (_dafny.SeqWithoutIsStrInference([d_232_v11_])))), _dafny.CodePoint('c'))
            (globalState).f2 = (p2)[default__.safeIndex(p1, len(p2))]
            d_236_v13_: C0
            nw41_ = C0()
            nw41_.ctor__()
            d_236_v13_ = nw41_
            d_232_v11_ = default__.fm16(self.f29, globalState)
        elif True:
            (self).f28 = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_237_i2_ in range(default__.abs(530))])
            d_238_v14_: str
            d_238_v14_ = _dafny.CodePoint('y')
            d_239_v15_: _dafny.Set
            d_239_v15_ = _dafny.Set({self.f29, len(d_225_v7_)})
            d_240_v16_: _dafny.Seq
            d_240_v16_ = _dafny.SeqWithoutIsStrInference([d_239_v15_])
            d_241_v17_: _dafny.Seq
            d_241_v17_ = _dafny.SeqWithoutIsStrInference([default__.fm2((self).f25, len(_dafny.SeqWithoutIsStrInference([(self).f25])), ((d_221_v2_)[len(self.f28)] if (len(self.f28)) in (d_221_v2_) else (self).f25), d_238_v14_, globalState), len((d_240_v16_)[default__.safeIndex(p1, len(d_240_v16_))]), self.f29])
            (globalState).f19 = d_241_v17_
            index29_ = default__.safeIndex(232, (d_223_v4_).length(0))
            (d_223_v4_)[index29_] = (self).f25
            index30_ = default__.safeIndex(232, (d_223_v4_).length(0))
            (d_223_v4_)[index30_] = (self).f25
            d_242_v18_: C0
            nw42_ = C0()
            nw42_.ctor__()
            d_242_v18_ = nw42_
            d_243_v19_: D3
            d_243_v19_ = D3_DC9(p1, d_238_v14_, self.f29)
            d_244_v20_: _dafny.Map
            d_244_v20_ = _dafny.Map({d_243_v19_: (self).f25})
            d_244_v20_ = (d_244_v20_).set(d_243_v19_, (d_223_v4_)[default__.safeIndex(232, (d_223_v4_).length(0))])
        d_245_v21_: _dafny.Array
        def lambda11_(d_246_i3_):
            return default__.safeModuloInt(d_246_i3_, (_dafny.MultiSet([(self).f25, (self).f25])).cardinality)

        init5_ = lambda11_
        nw43_ = _dafny.Array(None, 7)
        for i0_5_ in range(nw43_.length(0)):
            nw43_[i0_5_] = init5_(i0_5_)
        d_245_v21_ = nw43_
        index31_ = default__.safeIndex(394, (d_245_v21_).length(0))
        (d_245_v21_)[index31_] = default__.safeDivisionInt(218, p1)
        index32_ = default__.safeIndex(394, (d_245_v21_).length(0))
        (d_245_v21_)[index32_] = (p1) + (-798)
        index33_ = default__.safeIndex(394, (d_245_v21_).length(0))
        (d_245_v21_)[index33_] = self.f29
        d_247_v22_: _dafny.Array
        nw44_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
        d_247_v22_ = nw44_
        d_247_v22_ = d_247_v22_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_223_v4_).length(0)):
            d_248_i4_: int = guard_loop_3_
            if (True) and (((0) <= (d_248_i4_)) and ((d_248_i4_) < ((d_223_v4_).length(0)))):
                (d_223_v4_)[(d_248_i4_)] = (self).f25
        r0 = (self).f25
        r1 = self.f29
        return r0, r1

    def m3(self, globalState):
        d_249_v0_: _dafny.Array
        def lambda12_(d_250_i0_):
            return (d_250_i0_) - (self.f29)

        init6_ = lambda12_
        nw45_ = _dafny.Array(None, 4)
        for i0_6_ in range(nw45_.length(0)):
            nw45_[i0_6_] = init6_(i0_6_)
        d_249_v0_ = nw45_
        index34_ = default__.safeIndex(658, (d_249_v0_).length(0))
        (d_249_v0_)[index34_] = self.f29
        d_251_v1_: _dafny.Seq
        d_251_v1_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f25])
        index35_ = default__.safeIndex(658, (d_249_v0_).length(0))
        (d_249_v0_)[index35_] = default__.safeDivisionInt(self.f29, (len(d_251_v1_)) + (self.f29))
        d_252_v2_: C0
        nw46_ = C0()
        nw46_.ctor__()
        d_252_v2_ = nw46_
        d_253_v3_: _dafny.MultiSet
        d_253_v3_ = _dafny.MultiSet([self.f28])
        d_254_v4_: _dafny.Map
        d_254_v4_ = _dafny.Map({(d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]: (self).f25})
        d_255_v5_: _dafny.Array
        nw47_ = _dafny.Array(None, 19)
        nw47_[int(0)] = (self).f25
        nw47_[int(1)] = (self).f25
        nw47_[int(2)] = (_dafny.MultiSet([self.f28, self.f28])).issubset(d_253_v3_)
        nw47_[int(3)] = (self.f29) > (43)
        nw47_[int(4)] = default__.fm0(self.f29, self.f29, (self).f25, globalState)
        nw47_[int(5)] = True
        nw47_[int(6)] = default__.fm0((0) - ((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]), len(d_251_v1_), not((self).f25), globalState)
        nw47_[int(7)] = (self).f25
        nw47_[int(8)] = (self).f25
        nw47_[int(9)] = (self).f25
        nw47_[int(10)] = True
        nw47_[int(11)] = (self.f28) == (self.f28)
        nw47_[int(12)] = (_dafny.SeqWithoutIsStrInference([(self).f25])) < (d_251_v1_)
        nw47_[int(13)] = (d_251_v1_)[default__.safeIndex((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], len(d_251_v1_))]
        nw47_[int(14)] = (self).f25
        nw47_[int(15)] = (self).f25
        nw47_[int(16)] = ((d_254_v4_)[299] if (299) in (d_254_v4_) else False)
        nw47_[int(17)] = (self).f25
        nw47_[int(18)] = default__.fm0(-373, len(self.f28), not((self).f25), globalState)
        d_255_v5_ = nw47_
        index36_ = default__.safeIndex(448, (d_255_v5_).length(0))
        (d_255_v5_)[index36_] = default__.fm0(self.f29, self.f29, (self).f25, globalState)
        d_256_v6_: str
        d_256_v6_ = _dafny.CodePoint('i')
        d_257_v7_: _dafny.Map
        d_257_v7_ = _dafny.Map({(d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]: d_256_v6_})
        d_258_v8_: _dafny.Seq
        d_258_v8_ = _dafny.SeqWithoutIsStrInference([len(d_257_v7_)])
        d_259_v9_: D3
        d_259_v9_ = D3_DC9(self.f29, d_256_v6_, len(d_258_v8_))
        pat_let_tv7_ = d_249_v0_
        pat_let_tv8_ = d_249_v0_
        index37_ = default__.safeIndex(448, (d_255_v5_).length(0))
        def lambda13_(source3_):
            if source3_.is_DC8:
                d_260___mcc_h0_ = source3_.cf12
                d_261_cf12_ = d_260___mcc_h0_
                return (self).f25
            elif source3_.is_DC9:
                d_262___mcc_h1_ = source3_.cf13
                d_263___mcc_h2_ = source3_.cf14
                d_264___mcc_h3_ = source3_.cf15
                d_265_cf15_ = d_264___mcc_h3_
                d_266_cf14_ = d_263___mcc_h2_
                d_267_cf13_ = d_262___mcc_h1_
                return ((pat_let_tv8_)[default__.safeIndex(658, (pat_let_tv7_).length(0))]) <= (d_265_cf15_)
            elif source3_.is_DC7:
                d_268___mcc_h4_ = source3_.cf11
                d_269_cf11_ = d_268___mcc_h4_
                return ((self).f25) and ((self).f25)
            elif True:
                d_270___mcc_h5_ = source3_.cf16
                d_271_cf16_ = d_270___mcc_h5_
                return (self).f25

        (d_255_v5_)[index37_] = lambda13_(d_259_v9_)
        (globalState).f21 = len(default__.fm17((self).f25, ((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]) - (298), globalState))
        d_272_v10_: _dafny.Map
        d_272_v10_ = _dafny.Map({True: 778})
        d_273_v11_: _dafny.Seq
        d_273_v11_ = _dafny.SeqWithoutIsStrInference([d_272_v10_, ((d_272_v10_).set((d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))], (d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))])).set(True, 48)])
        hi2_ = len(d_273_v11_)
        for d_274_i1_ in range((d_252_v2_).fm6(d_258_v8_, (self).f25, globalState), hi2_):
            d_275_v12_: _dafny.Seq
            d_275_v12_ = _dafny.SeqWithoutIsStrInference([d_252_v2_, d_252_v2_, d_252_v2_, d_252_v2_])
            d_276_v13_: _dafny.Array
            nw48_ = _dafny.Array(None, 17)
            nw48_[int(0)] = d_252_v2_
            nw48_[int(1)] = (d_252_v2_ if (d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))] else d_252_v2_)
            nw48_[int(2)] = d_252_v2_
            nw48_[int(3)] = d_252_v2_
            nw48_[int(4)] = d_252_v2_
            nw48_[int(5)] = d_252_v2_
            nw48_[int(6)] = d_252_v2_
            nw48_[int(7)] = d_252_v2_
            nw48_[int(8)] = d_252_v2_
            nw48_[int(9)] = d_252_v2_
            nw48_[int(10)] = d_252_v2_
            nw48_[int(11)] = d_252_v2_
            nw48_[int(12)] = d_252_v2_
            nw48_[int(13)] = (d_275_v12_)[default__.safeIndex(self.f29, len(d_275_v12_))]
            nw48_[int(14)] = d_252_v2_
            nw48_[int(15)] = d_252_v2_
            nw48_[int(16)] = d_252_v2_
            d_276_v13_ = nw48_
            index38_ = default__.safeIndex(661, (d_276_v13_).length(0))
            (d_276_v13_)[index38_] = (d_275_v12_)[default__.safeIndex((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], len(d_275_v12_))]
            d_277_v14_: _dafny.Array
            nw49_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_277_v14_ = nw49_
            index39_ = default__.safeIndex(401, (d_277_v14_).length(0))
            (d_277_v14_)[index39_] = d_255_v5_
            d_278_v15_: D1
            d_278_v15_ = D1_DC1((d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))])
            d_279_v16_: _dafny.Seq
            d_279_v16_ = _dafny.SeqWithoutIsStrInference([d_278_v15_])
            d_280_v17_: D4
            d_280_v17_ = D4_DC11(default__.fm18(len(d_279_v16_), len(_dafny.SeqWithoutIsStrInference([(d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))], False, (d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))], (self).f25, (d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))]])), globalState))
            index40_ = default__.safeIndex(661, (d_276_v13_).length(0))
            index41_ = default__.safeIndex(401, (d_277_v14_).length(0))
            rhs28_ = (d_280_v17_).cf17
            rhs29_ = (d_272_v10_) | (d_272_v10_)
            rhs30_ = d_252_v2_
            rhs31_ = d_255_v5_
            lhs22_ = self
            lhs23_ = d_276_v13_
            lhs24_ = default__.safeIndex(661, (d_276_v13_).length(0))
            lhs25_ = d_277_v14_
            lhs26_ = default__.safeIndex(401, (d_277_v14_).length(0))
            lhs22_.f28 = rhs28_
            d_272_v10_ = rhs29_
            lhs23_[lhs24_] = rhs30_
            lhs25_[lhs26_] = rhs31_
            d_281_v18_: _dafny.Array
            def lambda14_(d_282_i2_):
                return D3_DC8((self).f25)

            init7_ = lambda14_
            nw50_ = _dafny.Array(None, 26)
            for i0_7_ in range(nw50_.length(0)):
                nw50_[i0_7_] = init7_(i0_7_)
            d_281_v18_ = nw50_
            d_283_v19_: D3
            d_283_v19_ = D3_DC8((d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))])
            index42_ = default__.safeIndex(131, (d_281_v18_).length(0))
            (d_281_v18_)[index42_] = d_283_v19_
            index43_ = default__.safeIndex(131, (d_281_v18_).length(0))
            rhs32_ = (len(self.f28)) * (self.f29)
            rhs33_ = (d_249_v0_ if (self).f25 else (d_249_v0_ if (self).f25 else d_249_v0_))
            rhs34_ = d_256_v6_
            rhs35_ = d_283_v19_
            lhs27_ = self
            lhs28_ = d_281_v18_
            lhs29_ = default__.safeIndex(131, (d_281_v18_).length(0))
            lhs27_.f29 = rhs32_
            d_249_v0_ = rhs33_
            d_256_v6_ = rhs34_
            lhs28_[lhs29_] = rhs35_
            (globalState).f2 = True
            index44_ = default__.safeIndex(658, (d_249_v0_).length(0))
            index45_ = default__.safeIndex(658, (d_249_v0_).length(0))
            rhs36_ = True
            rhs37_ = (d_274_i1_) * (d_274_i1_)
            rhs38_ = d_259_v9_
            rhs39_ = (d_274_i1_) + ((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))])
            rhs40_ = (self.f29 if (self).f25 else (d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))])
            lhs30_ = globalState
            lhs31_ = d_249_v0_
            lhs32_ = default__.safeIndex(658, (d_249_v0_).length(0))
            lhs33_ = d_249_v0_
            lhs34_ = default__.safeIndex(658, (d_249_v0_).length(0))
            lhs35_ = globalState
            lhs30_.f2 = rhs36_
            lhs31_[lhs32_] = rhs37_
            d_259_v9_ = rhs38_
            lhs33_[lhs34_] = rhs39_
            lhs35_.f5 = rhs40_
        if default__.fm0((0) - ((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]), (d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], (self).f25, globalState):
            d_256_v6_ = _dafny.CodePoint('w')
            if ((d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))] if (d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))] else False):
                d_284_v20_: _dafny.MultiSet
                d_284_v20_ = _dafny.MultiSet([(d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))], True])
                d_285_v21_: _dafny.Map
                d_285_v21_ = _dafny.Map({(d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))]: d_284_v20_})
                d_285_v21_ = d_285_v21_
                index46_ = default__.safeIndex(658, (d_249_v0_).length(0))
                (d_249_v0_)[index46_] = (0) - (self.f29)
                index47_ = default__.safeIndex(658, (d_249_v0_).length(0))
                (d_249_v0_)[index47_] = -356
                index48_ = default__.safeIndex(448, (d_255_v5_).length(0))
                (d_255_v5_)[index48_] = ((self.f28).set(default__.safeIndex(self.f29, len(self.f28)), d_256_v6_)) <= (self.f28)
                (globalState).f2 = (d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))]
            elif True:
                d_286_v22_: _dafny.Map
                def iife18_(_pat_let2_0):
                    def iife19_(d_287_dt__update__tmp_h0_):
                        def iife20_(_pat_let3_0):
                            def iife21_(d_288_dt__update_hcf17_h0_):
                                return D4_DC11(d_288_dt__update_hcf17_h0_)
                            return iife21_(_pat_let3_0)
                        return iife20_(self.f28)
                    return iife19_(_pat_let2_0)
                d_286_v22_ = _dafny.Map({default__.fm0((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], (d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], (d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))], globalState): (iife18_(D4_DC11(self.f28))).cf17})
                d_289_v23_: _dafny.Seq
                d_289_v23_ = _dafny.SeqWithoutIsStrInference([self.f28, (self.f28).set(default__.safeIndex((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], len(self.f28)), d_256_v6_), self.f28, ((d_286_v22_)[(self).f25] if ((self).f25) in (d_286_v22_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idiijp")))])
                d_290_v24_: _dafny.Array
                nw51_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_290_v24_ = nw51_
                d_291_v25_: _dafny.Map
                d_291_v25_ = _dafny.Map({(d_289_v23_)[default__.safeIndex(self.f29, len(d_289_v23_))]: d_290_v24_})
                d_291_v25_ = (d_291_v25_).set(self.f28, d_290_v24_)
                (globalState).f18 = (-494) * (((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))] if not(False) else (d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]))
                (globalState).f2 = ((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]) <= (self.f29)
                index49_ = default__.safeIndex(658, (d_249_v0_).length(0))
                (d_249_v0_)[index49_] = (d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]
                d_292_v26_: C0
                nw52_ = C0()
                nw52_.ctor__()
                d_292_v26_ = nw52_
            (globalState).f2 = ((d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))]) or ((self).f25)
            (self).f28 = self.f28
            (globalState).f5 = self.f29
        elif True:
            d_293_v27_: D4
            d_293_v27_ = D4_DC11(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qptks")))
            d_294_v28_: _dafny.Map
            d_294_v28_ = _dafny.Map({self.f29: d_293_v27_})
            d_294_v28_ = (d_294_v28_).set((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], d_293_v27_)
            d_295_v29_: _dafny.Set
            d_295_v29_ = _dafny.Set({((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]) * ((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]), default__.safeDivisionInt(self.f29, (d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))])})
            d_295_v29_ = (d_295_v29_) | (d_295_v29_)
            d_296_v30_: D2
            d_296_v30_ = D2_DC4(self.f29, (self).f25)
            source4_ = d_296_v30_
            if source4_.is_DC4:
                d_297___mcc_h6_ = source4_.cf4
                d_298___mcc_h7_ = source4_.cf5
                d_299_cf5_ = d_298___mcc_h7_
                d_300_cf4_ = d_297___mcc_h6_
                d_301_v31_: _dafny.Map
                d_301_v31_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))] for d_302_i3_ in range(default__.abs(-706))])) + (d_258_v8_): d_299_cf5_})
                d_301_v31_ = (d_301_v31_).set(d_258_v8_, d_299_cf5_)
                d_303_v32_: _dafny.Map
                d_303_v32_ = _dafny.Map({(d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]: (d_253_v3_ if True else d_253_v3_)})
                d_303_v32_ = (d_303_v32_).set(len(d_258_v8_), d_253_v3_)
                d_304_v33_: _dafny.Set
                d_304_v33_ = _dafny.Set({True, (d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))]})
                d_305_v34_: _dafny.Map
                d_305_v34_ = _dafny.Map({(d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]: len(d_304_v33_)})
                d_306_v35_: _dafny.Set
                d_306_v35_ = _dafny.Set({_dafny.CodePoint('t')})
                d_307_v36_: _dafny.MultiSet
                d_307_v36_ = _dafny.MultiSet([len(d_306_v35_), 305])
                d_308_v37_: _dafny.Map
                d_308_v37_ = _dafny.Map({(d_295_v29_ if (d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))] else d_295_v29_): (default__.fm18((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], ((d_305_v34_)[self.f29] if (self.f29) in (d_305_v34_) else ((d_307_v36_)[(d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]] if ((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]) in (d_307_v36_) else d_300_cf4_)), globalState)) + (_dafny.SeqWithoutIsStrInference([d_256_v6_ for d_309_i4_ in range(default__.abs(-110))]))})
                d_308_v37_ = d_308_v37_
                rhs41_ = d_300_cf4_
                rhs42_ = (d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))]
                rhs43_ = (0) - (self.f29)
                rhs44_ = self.f29
                rhs45_ = d_255_v5_
                lhs36_ = globalState
                lhs37_ = globalState
                lhs38_ = globalState
                lhs36_.f5 = rhs41_
                d_299_cf5_ = rhs42_
                lhs37_.f18 = rhs43_
                lhs38_.f18 = rhs44_
                d_255_v5_ = rhs45_
            elif source4_.is_DC5:
                d_310___mcc_h8_ = source4_.cf6
                d_311___mcc_h9_ = source4_.cf7
                d_312___mcc_h10_ = source4_.cf8
                d_313___mcc_h11_ = source4_.cf9
                d_314_cf9_ = d_313___mcc_h11_
                d_315_cf8_ = d_312___mcc_h10_
                d_316_cf7_ = d_311___mcc_h9_
                d_317_cf6_ = d_310___mcc_h8_
                index50_ = default__.safeIndex(448, (d_255_v5_).length(0))
                (d_255_v5_)[index50_] = (d_317_cf6_) > (275)
                (globalState).f22 = (default__.fm1((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], d_317_cf6_, globalState)).cardinality
                d_318_v38_: _dafny.MultiSet
                d_318_v38_ = _dafny.MultiSet([self.f29, self.f29])
                d_319_v39_: int
                d_319_v39_ = (d_318_v38_).cardinality
                d_320_v42_: _dafny.Seq
                d_320_v42_ = _dafny.SeqWithoutIsStrInference([d_257_v7_])
                d_321_v43_: _dafny.Array
                nw53_ = _dafny.Array(None, 26)
                nw53_[int(0)] = d_257_v7_
                nw53_[int(1)] = (default__.fm19(len(self.f28), True, globalState)) | (d_257_v7_)
                nw53_[int(2)] = _dafny.Map({len(d_251_v1_): _dafny.CodePoint('i')})
                nw53_[int(3)] = _dafny.Map({self.f29: (self.f28)[default__.safeIndex(self.f29, len(self.f28))]})
                nw53_[int(4)] = (d_257_v7_).set((d_319_v39_), _dafny.CodePoint('e'))
                nw53_[int(5)] = (d_257_v7_) | (d_257_v7_)
                nw53_[int(6)] = d_257_v7_
                nw53_[int(7)] = d_257_v7_
                nw53_[int(8)] = d_257_v7_
                nw53_[int(9)] = (d_257_v7_).set(self.f29, d_256_v6_)
                nw53_[int(10)] = default__.fm19(self.f29, d_316_cf7_, globalState)
                nw53_[int(11)] = _dafny.Map({(d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]: (self.f28)[default__.safeIndex((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], len(self.f28))]})
                nw53_[int(12)] = _dafny.Map({self.f29: d_256_v6_})
                nw53_[int(13)] = (d_257_v7_) | (d_257_v7_)
                nw53_[int(14)] = (d_257_v7_) | (d_257_v7_)
                nw53_[int(15)] = (d_257_v7_) | (d_257_v7_)
                nw53_[int(16)] = d_257_v7_
                nw53_[int(17)] = ((D5_DC14(_dafny.Map({self.f29: d_256_v6_}))).cf24) | (d_257_v7_)
                def iife22_():
                    coll14_ = _dafny.Map()
                    compr_14_: int
                    for compr_14_ in (d_258_v8_).Elements:
                        d_322_v40_: int = compr_14_
                        if (d_322_v40_) in (d_258_v8_):
                            coll14_[(d_322_v40_) * (self.f29)] = d_256_v6_
                    return _dafny.Map(coll14_)
                nw53_[int(18)] = iife22_()
                
                nw53_[int(19)] = d_257_v7_
                nw53_[int(20)] = d_257_v7_
                def iife23_():
                    coll15_ = _dafny.Map()
                    compr_15_: int
                    for compr_15_ in _dafny.IntegerRange(297, 675):
                        d_323_v41_: int = compr_15_
                        if ((297) <= (d_323_v41_)) and ((d_323_v41_) < (675)):
                            coll15_[default__.safeDivisionInt(d_323_v41_, self.f29)] = _dafny.CodePoint('s')
                    return _dafny.Map(coll15_)
                nw53_[int(21)] = (iife23_()
                ) | (_dafny.Map({402: d_256_v6_}))
                nw53_[int(22)] = d_257_v7_
                nw53_[int(23)] = (d_320_v42_)[default__.safeIndex((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], len(d_320_v42_))]
                nw53_[int(24)] = d_257_v7_
                nw53_[int(25)] = d_257_v7_
                d_321_v43_ = nw53_
                index51_ = default__.safeIndex(581, (d_321_v43_).length(0))
                (d_321_v43_)[index51_] = d_257_v7_
                index52_ = default__.safeIndex(581, (d_321_v43_).length(0))
                (d_321_v43_)[index52_] = ((d_257_v7_) | (d_257_v7_)) | ((d_320_v42_)[default__.safeIndex(d_317_cf6_, len(d_320_v42_))])
                (globalState).f22 = (default__.safeModuloInt(len(_dafny.Set({d_317_cf6_})), self.f29)) * ((331) + (self.f29))
            elif source4_.is_DC3:
                d_324___mcc_h12_ = source4_.cf3
                d_325_cf3_ = d_324___mcc_h12_
                (globalState).f18 = self.f29
                (self).f28 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccueaym"))) + (self.f28)) + (_dafny.SeqWithoutIsStrInference([d_256_v6_ for d_326_i5_ in range(default__.abs(248))]))
                d_327_v44_: D2
                d_327_v44_ = D2_DC4(self.f29, (self).f25)
                d_328_v45_: D2
                d_328_v45_ = D2_DC6(d_327_v44_)
                d_328_v45_ = default__.fm20((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], globalState)
                (globalState).f18 = self.f29
            elif True:
                d_329___mcc_h13_ = source4_.cf10
                d_330_cf10_ = d_329___mcc_h13_
                d_331_v46_: _dafny.MultiSet
                d_331_v46_ = _dafny.MultiSet([(d_255_v5_)[default__.safeIndex(448, (d_255_v5_).length(0))]])
                (globalState).f2 = default__.fm0((d_331_v46_).cardinality, (d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], (self).f25, globalState)
                d_249_v0_ = (d_249_v0_ if not((self).f25) else d_249_v0_)
                (self).f29 = (d_252_v2_).fm6((d_258_v8_) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([882])).cardinality for d_332_i6_ in range(default__.abs(792))])), (self).f25, globalState)
                d_258_v8_ = (self).fm15(globalState)
            (globalState).f2 = (self).f25
            d_333_v47_: _dafny.Map
            d_333_v47_ = _dafny.Map({len(d_251_v1_): self.f29})
            d_334_v50_: D6
            d_334_v50_ = D6_DC17(default__.fm21(globalState))
            d_335_v52_: _dafny.Seq
            d_335_v52_ = _dafny.SeqWithoutIsStrInference([default__.fm21(globalState), d_333_v47_, d_333_v47_])
            d_336_v53_: _dafny.Array
            nw54_ = _dafny.Array(None, 24)
            nw54_[int(0)] = (d_333_v47_) | (d_333_v47_)
            nw54_[int(1)] = _dafny.Map({self.f29: 159})
            def iife24_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(155, 340):
                    d_337_v48_: int = compr_16_
                    if ((155) <= (d_337_v48_)) and ((d_337_v48_) < (340)):
                        coll16_[(d_337_v48_) + ((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))])] = len(_dafny.Map({(d_251_v1_)[default__.safeIndex(self.f29, len(d_251_v1_))]: True}))
                return _dafny.Map(coll16_)
            nw54_[int(2)] = ((iife24_()
            ).set(self.f29, self.f29)) | (d_333_v47_)
            def iife25_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in _dafny.IntegerRange(222, 365):
                    d_338_v49_: int = compr_17_
                    if ((222) <= (d_338_v49_)) and ((d_338_v49_) < (365)):
                        coll17_[default__.safeModuloInt(d_338_v49_, (d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))])] = (d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))]
                return _dafny.Map(coll17_)
            nw54_[int(3)] = (iife25_()
            ) | (d_333_v47_)
            nw54_[int(4)] = d_333_v47_
            nw54_[int(5)] = default__.fm21(globalState)
            nw54_[int(6)] = (d_333_v47_).set((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], (d_252_v2_).fm5(globalState))
            nw54_[int(7)] = (d_333_v47_).set((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], (d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))])
            nw54_[int(8)] = (d_333_v47_).set(len(default__.fm18((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], (d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))], globalState)), self.f29)
            nw54_[int(9)] = d_333_v47_
            nw54_[int(10)] = d_333_v47_
            nw54_[int(11)] = d_333_v47_
            nw54_[int(12)] = d_333_v47_
            nw54_[int(13)] = ((d_334_v50_).cf30) | (_dafny.Map({(0) - (self.f29): (0) - (self.f29)}))
            nw54_[int(14)] = d_333_v47_
            nw54_[int(15)] = d_333_v47_
            nw54_[int(16)] = d_333_v47_
            def iife26_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in _dafny.IntegerRange(463, 475):
                    d_339_v51_: int = compr_18_
                    if ((463) <= (d_339_v51_)) and ((d_339_v51_) < (475)):
                        coll18_[(d_339_v51_) * ((d_249_v0_)[default__.safeIndex(658, (d_249_v0_).length(0))])] = 685
                return _dafny.Map(coll18_)
            nw54_[int(17)] = iife26_()
            
            nw54_[int(18)] = (d_335_v52_)[default__.safeIndex((0) - (self.f29), len(d_335_v52_))]
            nw54_[int(19)] = (d_333_v47_) | (d_333_v47_)
            nw54_[int(20)] = d_333_v47_
            nw54_[int(21)] = d_333_v47_
            nw54_[int(22)] = d_333_v47_
            nw54_[int(23)] = d_333_v47_
            d_336_v53_ = nw54_
            d_336_v53_ = d_336_v53_


class C2(T1, T0):
    def  __init__(self):
        self._f26: str = _dafny.CodePoint('D')
        self._f27: _dafny.Array = _dafny.Array(None, 0)
        self._f25: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f26, f27, f25):
        (self)._f26 = f26
        (self)._f27 = f27
        (self)._f25 = f25

    def fm3(self, p0, p1, p2, p3, globalState):
        return (_dafny.Set({195})).isdisjoint(_dafny.Set({len(_dafny.Map({(self).f25: _dafny.Map({(self).f25: 602})}))}))

    def fm8(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "goffav")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "atkuqwah")))

    def fm9(self, p0, p1, p2, globalState):
        return default__.safeModuloInt(681, (0) - (len(_dafny.Map({(self).f25: True}))))

    def m2(self, p0, p1, p2, globalState):
        d_340_v0_: _dafny.Map
        d_340_v0_ = _dafny.Map({p2: False})
        (globalState).f5 = len(d_340_v0_)
        d_341_v1_: _dafny.Seq
        d_341_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
        d_342_v2_: _dafny.Seq
        d_342_v2_ = _dafny.SeqWithoutIsStrInference([len(d_341_v1_)])
        hi3_ = p2
        for d_343_i0_ in range(len((d_342_v2_) + (d_342_v2_)), hi3_):
            (globalState).f2 = not(True)
            d_344_v3_: _dafny.Set
            d_344_v3_ = _dafny.Set({p1})
            d_345_v4_: D2
            d_345_v4_ = D2_DC4(len(d_344_v3_), p1)
            d_346_v5_: D2
            d_346_v5_ = D2_DC6(d_345_v4_)
            source5_ = d_346_v5_
            if source5_.is_DC4:
                d_347___mcc_h0_ = source5_.cf4
                d_348___mcc_h1_ = source5_.cf5
                d_349_cf5_ = d_348___mcc_h1_
                d_350_cf4_ = d_347___mcc_h0_
                d_351_v6_: _dafny.Array
                nw55_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_351_v6_ = nw55_
                d_352_v7_: _dafny.Map
                d_352_v7_ = _dafny.Map({d_349_cf5_: d_351_v6_})
                d_352_v7_ = (d_352_v7_).set(d_349_cf5_, d_351_v6_)
                d_353_v8_: _dafny.Map
                d_353_v8_ = _dafny.Map({d_349_cf5_: (self).f25})
                d_354_v9_: _dafny.Map
                d_354_v9_ = _dafny.Map({d_353_v8_: p1})
                d_354_v9_ = d_354_v9_
                d_349_cf5_ = False
                d_353_v8_ = d_353_v8_
            elif source5_.is_DC5:
                d_355___mcc_h2_ = source5_.cf6
                d_356___mcc_h3_ = source5_.cf7
                d_357___mcc_h4_ = source5_.cf8
                d_358___mcc_h5_ = source5_.cf9
                d_359_cf9_ = d_358___mcc_h5_
                d_360_cf8_ = d_357___mcc_h4_
                d_361_cf7_ = d_356___mcc_h3_
                d_362_cf6_ = d_355___mcc_h2_
                (globalState).f2 = (d_359_cf9_ if (default__.fm10(globalState)) in (d_341_v1_) else (self).f25)
                d_363_v10_: int
                d_363_v10_ = (p2) * (p2)
                d_363_v10_ = d_363_v10_
                (globalState).f19 = (d_342_v2_) + (d_342_v2_)
                d_364_v11_: _dafny.Array
                def lambda15_(d_365_i1_):
                    return (self).f25

                init8_ = lambda15_
                nw56_ = _dafny.Array(None, 2)
                for i0_8_ in range(nw56_.length(0)):
                    nw56_[i0_8_] = init8_(i0_8_)
                d_364_v11_ = nw56_
                d_366_v12_: _dafny.Seq
                d_366_v12_ = _dafny.SeqWithoutIsStrInference([p1, d_361_cf7_])
                d_367_v13_: _dafny.Map
                d_367_v13_ = _dafny.Map({p1: len(d_341_v1_)})
                d_368_v14_: _dafny.MultiSet
                d_368_v14_ = _dafny.MultiSet([-568, d_362_cf6_, d_343_i0_, len(d_367_v13_)])
                index53_ = default__.safeIndex(889, (d_364_v11_).length(0))
                def iife27_():
                    coll19_ = _dafny.Set()
                    compr_19_: int
                    for compr_19_ in (d_368_v14_).Elements:
                        d_369_v15_: int = compr_19_
                        if (d_369_v15_) in (d_368_v14_):
                            def iife28_():
                                coll20_ = _dafny.Set()
                                compr_20_: int
                                for compr_20_ in _dafny.IntegerRange(308, 305):
                                    d_370_v16_: int = compr_20_
                                    if ((308) <= (d_370_v16_)) and ((d_370_v16_) < (305)):
                                        coll20_ = coll20_.union(_dafny.Set([default__.safeModuloInt(d_370_v16_, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([-982]))})))]))
                                return _dafny.Set(coll20_)
                            coll19_ = coll19_.union(_dafny.Set([default__.safeModuloInt(d_369_v15_, (_dafny.MultiSet([len(iife28_()
)])).cardinality)]))
                    return _dafny.Set(coll19_)
                (d_364_v11_)[index53_] = (not((d_366_v12_)[default__.safeIndex(len(iife27_()
                ), len(d_366_v12_))])) == ((self).f25)
                d_371_v17_: _dafny.Array
                def lambda16_(d_372_v3_):
                    def lambda17_(d_373_i2_):
                        return d_372_v3_

                    return lambda17_

                init9_ = lambda16_(d_344_v3_)
                nw57_ = _dafny.Array(None, 21)
                for i0_9_ in range(nw57_.length(0)):
                    nw57_[i0_9_] = init9_(i0_9_)
                d_371_v17_ = nw57_
                index54_ = default__.safeIndex(889, (d_364_v11_).length(0))
                rhs46_ = p1
                rhs47_ = True
                rhs48_ = d_371_v17_
                lhs39_ = d_364_v11_
                lhs40_ = default__.safeIndex(889, (d_364_v11_).length(0))
                lhs41_ = globalState
                lhs39_[lhs40_] = rhs46_
                lhs41_.f2 = rhs47_
                d_371_v17_ = rhs48_
            elif source5_.is_DC3:
                d_374___mcc_h6_ = source5_.cf3
                d_375_cf3_ = d_374___mcc_h6_
                d_376_v18_: C0
                nw58_ = C0()
                nw58_.ctor__()
                d_376_v18_ = nw58_
                d_376_v18_ = (d_376_v18_ if p1 else d_376_v18_)
                (globalState).f22 = (d_343_i0_ if p1 else d_343_i0_)
                d_377_v19_: int
                out5_: int
                out5_ = default__.m0(globalState)
                d_377_v19_ = out5_
                (globalState).f2 = not((self).f25)
            elif True:
                d_378___mcc_h7_ = source5_.cf10
                d_379_cf10_ = d_378___mcc_h7_
                d_380_v20_: _dafny.Map
                d_380_v20_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_340_v0_ for d_381_i3_ in range(default__.abs(685))]): ((self).f25 if p1 else (self).f25)})
                d_382_v21_: _dafny.Seq
                d_382_v21_ = _dafny.SeqWithoutIsStrInference([d_340_v0_, d_340_v0_])
                d_383_v22_: D2
                d_383_v22_ = D2_DC5(d_343_i0_, p1, (self).f25, (self).f25)
                d_380_v20_ = (d_380_v20_).set(d_382_v21_, (d_383_v22_).cf9)
                (globalState).f22 = d_343_i0_
                (globalState).f2 = p1
                d_384_v23_: _dafny.Array
                def lambda18_(d_385_v3_):
                    def lambda19_(d_386_i4_):
                        return (444) > (len(d_385_v3_))

                    return lambda19_

                init10_ = lambda18_(d_344_v3_)
                nw59_ = _dafny.Array(None, 6)
                for i0_10_ in range(nw59_.length(0)):
                    nw59_[i0_10_] = init10_(i0_10_)
                d_384_v23_ = nw59_
                index55_ = default__.safeIndex(202, (d_384_v23_).length(0))
                (d_384_v23_)[index55_] = (self).f25
                index56_ = default__.safeIndex(202, (d_384_v23_).length(0))
                (d_384_v23_)[index56_] = (self).f25
            d_387_v24_: str
            d_387_v24_ = _dafny.CodePoint('y')
            d_387_v24_ = default__.fm10(globalState)
            d_388_v25_: _dafny.Array
            nw60_ = _dafny.Array(int(0), 25)
            d_388_v25_ = nw60_
            nw61_ = _dafny.Array(int(0), 1)
            d_388_v25_ = nw61_
        d_389_i5_: int
        d_389_i5_ = 0
        with _dafny.label("1"):
            while (self).f25:
                with _dafny.c_label("1"):
                    if (d_389_i5_) >= (100):
                        raise _dafny.Break("1")
                    d_389_i5_ = (d_389_i5_) + (1)
                    d_390_v26_: D2
                    d_390_v26_ = D2_DC4(p2, p1)
                    d_391_v27_: _dafny.Seq
                    d_391_v27_ = _dafny.SeqWithoutIsStrInference([d_390_v26_, default__.fm11(globalState), d_390_v26_])
                    d_392_v28_: _dafny.Map
                    d_392_v28_ = _dafny.Map({(self).f26: p2})
                    d_393_v29_: _dafny.Array
                    nw62_ = _dafny.Array(None, 20)
                    nw62_[int(0)] = p1
                    nw62_[int(1)] = (self).f25
                    nw62_[int(2)] = p1
                    nw62_[int(3)] = (self).f25
                    nw62_[int(4)] = not(p1)
                    nw62_[int(5)] = p1
                    nw62_[int(6)] = (self).f25
                    nw62_[int(7)] = not (p1) or ((self).f25)
                    nw62_[int(8)] = not (p1) or (p1)
                    nw62_[int(9)] = (self).f25
                    nw62_[int(10)] = (self).fm3((self).f25, p2, len(d_341_v1_), d_392_v28_, globalState)
                    nw62_[int(11)] = not(False)
                    nw62_[int(12)] = (self).fm3(p1, p2, p2, d_392_v28_, globalState)
                    nw62_[int(13)] = (self).f25
                    nw62_[int(14)] = (self).f25
                    nw62_[int(15)] = p1
                    nw62_[int(16)] = not(True)
                    nw62_[int(17)] = not((self).f25)
                    nw62_[int(18)] = (self).f25
                    nw62_[int(19)] = True
                    d_393_v29_ = nw62_
                    index57_ = default__.safeIndex(207, (d_393_v29_).length(0))
                    (d_393_v29_)[index57_] = p1
                    d_394_v30_: _dafny.Map
                    d_394_v30_ = _dafny.Map({p1: p2})
                    d_395_v31_: _dafny.Set
                    d_395_v31_ = _dafny.Set({p1})
                    d_396_v32_: _dafny.MultiSet
                    d_396_v32_ = _dafny.MultiSet([p2, p2, default__.safeModuloInt(len(d_395_v31_), 502), default__.safeDivisionInt(p2, p2)])
                    index58_ = default__.safeIndex(207, (d_393_v29_).length(0))
                    rhs49_ = (d_391_v27_) + ((d_391_v27_) + (d_391_v27_))
                    rhs50_ = 387
                    rhs51_ = (d_396_v32_).cardinality
                    rhs52_ = ((self).f25 if True else False)
                    rhs53_ = ((d_394_v30_) | (_dafny.Map({True: p2}))) | ((_dafny.Map({(self).f25: p2})) | (d_394_v30_))
                    lhs42_ = globalState
                    lhs43_ = globalState
                    lhs44_ = d_393_v29_
                    lhs45_ = default__.safeIndex(207, (d_393_v29_).length(0))
                    d_391_v27_ = rhs49_
                    lhs42_.f22 = rhs50_
                    lhs43_.f18 = rhs51_
                    lhs44_[lhs45_] = rhs52_
                    d_394_v30_ = rhs53_
                    d_397_v33_: _dafny.Array
                    nw63_ = _dafny.Array(_dafny.Array(None, 0), 2)
                    d_397_v33_ = nw63_
                    d_397_v33_ = d_397_v33_
                    rhs54_ = (d_393_v29_ if p1 else d_393_v29_)
                    rhs55_ = not(default__.fm0(default__.fm2((d_393_v29_)[default__.safeIndex(207, (d_393_v29_).length(0))], 994, not(p1), (self).f26, globalState), default__.safeDivisionInt(-823, p2), (d_341_v1_) <= (d_341_v1_), globalState))
                    lhs46_ = globalState
                    d_393_v29_ = rhs54_
                    lhs46_.f2 = rhs55_
                    if p1:
                        (globalState).f2 = (self).f25
                        d_398_v34_: _dafny.Array
                        def lambda20_(d_399_i6_):
                            return (d_399_i6_) * (-782)

                        init11_ = lambda20_
                        nw64_ = _dafny.Array(None, 14)
                        for i0_11_ in range(nw64_.length(0)):
                            nw64_[i0_11_] = init11_(i0_11_)
                        d_398_v34_ = nw64_
                        index59_ = default__.safeIndex(119, (d_398_v34_).length(0))
                        (d_398_v34_)[index59_] = p2
                        index60_ = default__.safeIndex(119, (d_398_v34_).length(0))
                        (d_398_v34_)[index60_] = (default__.safeModuloInt(p2, p2)) * (p2)
                        d_400_v35_: _dafny.Seq
                        d_400_v35_ = _dafny.SeqWithoutIsStrInference([(d_393_v29_)[default__.safeIndex(207, (d_393_v29_).length(0))]])
                        index61_ = default__.safeIndex(119, (d_398_v34_).length(0))
                        (d_398_v34_)[index61_] = ((d_398_v34_)[default__.safeIndex(119, (d_398_v34_).length(0))]) * ((0) - ((0) - (len((_dafny.SeqWithoutIsStrInference([p1, p1, (d_393_v29_)[default__.safeIndex(207, (d_393_v29_).length(0))]])) + (_dafny.SeqWithoutIsStrInference([(d_400_v35_)[default__.safeIndex(len(d_400_v35_), len(d_400_v35_))]]))))))
                        index62_ = default__.safeIndex(207, (d_393_v29_).length(0))
                        (d_393_v29_)[index62_] = False
                        index63_ = default__.safeIndex(207, (d_393_v29_).length(0))
                        (d_393_v29_)[index63_] = (self).f25
                    elif True:
                        (globalState).f2 = (-36) <= (p2)
                        d_401_v36_: _dafny.Array
                        nw65_ = _dafny.Array(int(0), 5)
                        d_401_v36_ = nw65_
                        index64_ = default__.safeIndex(284, (d_401_v36_).length(0))
                        (d_401_v36_)[index64_] = p2
                        d_402_v37_: _dafny.Seq
                        d_402_v37_ = _dafny.SeqWithoutIsStrInference([d_393_v29_, d_393_v29_])
                        d_403_v38_: _dafny.MultiSet
                        d_403_v38_ = _dafny.MultiSet([(self).f25])
                        d_404_v39_: _dafny.MultiSet
                        d_404_v39_ = _dafny.MultiSet([d_393_v29_, (d_393_v29_ if (self).f25 else d_393_v29_), d_393_v29_, (d_402_v37_)[default__.safeIndex((d_403_v38_).cardinality, len(d_402_v37_))]])
                        index65_ = default__.safeIndex(284, (d_401_v36_).length(0))
                        rhs56_ = (d_404_v39_).cardinality
                        rhs57_ = (default__.safeModuloInt(811, 345)) - (default__.safeDivisionInt((d_403_v38_).cardinality, len(d_341_v1_)))
                        rhs58_ = default__.safeDivisionInt(884, default__.safeModuloInt(p2, p2))
                        lhs47_ = globalState
                        lhs48_ = globalState
                        lhs49_ = d_401_v36_
                        lhs50_ = default__.safeIndex(284, (d_401_v36_).length(0))
                        lhs47_.f18 = rhs56_
                        lhs48_.f5 = rhs57_
                        lhs49_[lhs50_] = rhs58_
                        (globalState).f22 = p2
                        (globalState).f18 = (0) - (default__.safeDivisionInt((0) - (p2), p2))
                        index66_ = default__.safeIndex(207, (d_393_v29_).length(0))
                        (d_393_v29_)[index66_] = (self).f25
                    pass
            pass
        d_405_v40_: _dafny.Array
        nw66_ = _dafny.Array(False, 18)
        d_405_v40_ = nw66_
        index67_ = default__.safeIndex(54, (d_405_v40_).length(0))
        (d_405_v40_)[index67_] = not(((d_340_v0_)[p2] if (p2) in (d_340_v0_) else not(False)))
        index68_ = default__.safeIndex(54, (d_405_v40_).length(0))
        (d_405_v40_)[index68_] = (self).f25
        d_406_v41_: _dafny.Array
        def lambda21_(d_407_p1_, d_408_p2_):
            def lambda22_(d_409_i7_):
                return _dafny.Map({d_407_p1_: _dafny.MultiSet([d_408_p2_, d_408_p2_])})

            return lambda22_

        init12_ = lambda21_(p1, p2)
        nw67_ = _dafny.Array(None, 11)
        for i0_12_ in range(nw67_.length(0)):
            nw67_[i0_12_] = init12_(i0_12_)
        d_406_v41_ = nw67_
        d_410_v42_: _dafny.MultiSet
        d_410_v42_ = _dafny.MultiSet([(self).f25])
        d_411_v43_: _dafny.MultiSet
        d_411_v43_ = _dafny.MultiSet([(d_410_v42_).cardinality])
        d_412_v44_: _dafny.Map
        d_412_v44_ = _dafny.Map({p1: d_411_v43_})
        index69_ = default__.safeIndex(879, (d_406_v41_).length(0))
        (d_406_v41_)[index69_] = d_412_v44_
        d_413_v45_: D2
        d_413_v45_ = D2_DC4(p2, not((self).f25))
        d_414_v46_: C0
        nw68_ = C0()
        nw68_.ctor__()
        d_414_v46_ = nw68_
        d_415_v47_: _dafny.Seq
        d_415_v47_ = _dafny.SeqWithoutIsStrInference([d_414_v46_, d_414_v46_])
        d_416_v48_: _dafny.Set
        d_416_v48_ = _dafny.Set({(d_415_v47_)[default__.safeIndex(p2, len(d_415_v47_))]})
        index70_ = default__.safeIndex(879, (d_406_v41_).length(0))
        rhs59_ = (_dafny.SeqWithoutIsStrInference([p2, p2, len(d_416_v48_)])) + ((d_342_v2_) + (d_342_v2_))
        rhs60_ = (d_412_v44_) | (d_412_v44_)
        rhs61_ = D2_DC4((p2) - (-845), p1)
        lhs51_ = d_406_v41_
        lhs52_ = default__.safeIndex(879, (d_406_v41_).length(0))
        d_342_v2_ = rhs59_
        lhs51_[lhs52_] = rhs60_
        d_413_v45_ = rhs61_
        d_417_v49_: _dafny.Seq
        d_417_v49_ = _dafny.SeqWithoutIsStrInference([d_342_v2_])
        hi4_ = (d_410_v42_).cardinality
        for d_418_i8_ in range(len(d_417_v49_), hi4_):
            d_341_v1_ = ((d_341_v1_) + (d_341_v1_)) + ((d_341_v1_) + (d_341_v1_))
            (globalState).f5 = d_418_i8_
            d_419_v50_: _dafny.Map
            d_419_v50_ = _dafny.Map({d_405_v40_: d_418_i8_})
            d_420_v51_: _dafny.Set
            d_420_v51_ = _dafny.Set({len(d_419_v50_)})
            if (d_418_i8_) not in (d_420_v51_):
                d_421_v52_: D1
                d_421_v52_ = D1_DC2(p1)
                rhs62_ = d_405_v40_
                rhs63_ = (0) - (default__.safeDivisionInt((572) * (896), d_418_i8_))
                rhs64_ = (d_405_v40_)[default__.safeIndex(54, (d_405_v40_).length(0))]
                rhs65_ = d_421_v52_
                rhs66_ = (d_414_v46_).fm5(globalState)
                lhs53_ = globalState
                lhs54_ = globalState
                lhs55_ = globalState
                d_405_v40_ = rhs62_
                lhs53_.f5 = rhs63_
                lhs54_.f2 = rhs64_
                d_421_v52_ = rhs65_
                lhs55_.f22 = rhs66_
                (globalState).f18 = len((d_342_v2_).set(default__.safeIndex(p2, len(d_342_v2_)), len(d_342_v2_)))
                d_422_v53_: _dafny.Map
                d_422_v53_ = _dafny.Map({d_405_v40_: (d_405_v40_)[default__.safeIndex(54, (d_405_v40_).length(0))]})
                d_423_v54_: _dafny.Array
                nw69_ = _dafny.Array(None, 3)
                nw69_[int(0)] = d_422_v53_
                nw69_[int(1)] = (d_422_v53_).set(d_405_v40_, (d_405_v40_)[default__.safeIndex(54, (d_405_v40_).length(0))])
                nw69_[int(2)] = d_422_v53_
                d_423_v54_ = nw69_
                index71_ = default__.safeIndex(767, (d_423_v54_).length(0))
                (d_423_v54_)[index71_] = _dafny.Map({d_405_v40_: True})
                d_424_v55_: _dafny.Map
                d_424_v55_ = _dafny.Map({(d_405_v40_)[default__.safeIndex(54, (d_405_v40_).length(0))]: d_422_v53_})
                index72_ = default__.safeIndex(767, (d_423_v54_).length(0))
                (d_423_v54_)[index72_] = (((d_424_v55_)[(d_405_v40_)[default__.safeIndex(54, (d_405_v40_).length(0))]] if ((d_405_v40_)[default__.safeIndex(54, (d_405_v40_).length(0))]) in (d_424_v55_) else d_422_v53_)) | (d_422_v53_)
                d_425_v56_: _dafny.Map
                d_425_v56_ = _dafny.Map({d_411_v43_: (p2) + (len(d_341_v1_))})
                d_425_v56_ = (d_425_v56_ if p1 else d_425_v56_)
                d_426_v57_: C0
                nw70_ = C0()
                nw70_.ctor__()
                d_426_v57_ = nw70_
            elif True:
                (globalState).f21 = len(d_341_v1_)
                (globalState).f5 = -32
                d_427_v58_: _dafny.Array
                def lambda23_(d_428_v1_):
                    def lambda24_(d_429_i9_):
                        return (_dafny.MultiSet([d_428_v1_, d_428_v1_])) | (_dafny.MultiSet([d_428_v1_, _dafny.SeqWithoutIsStrInference([(self).f26 for d_430_i10_ in range(default__.abs(441))])]))

                    return lambda24_

                init13_ = lambda23_(d_341_v1_)
                nw71_ = _dafny.Array(None, 20)
                for i0_13_ in range(nw71_.length(0)):
                    nw71_[i0_13_] = init13_(i0_13_)
                d_427_v58_ = nw71_
                d_431_v59_: _dafny.MultiSet
                d_431_v59_ = _dafny.MultiSet([d_341_v1_, d_341_v1_])
                index73_ = default__.safeIndex(97, (d_427_v58_).length(0))
                (d_427_v58_)[index73_] = d_431_v59_
                index74_ = default__.safeIndex(97, (d_427_v58_).length(0))
                (d_427_v58_)[index74_] = (d_431_v59_).set(d_341_v1_, default__.abs(p2))
                d_432_v60_: _dafny.Array
                nw72_ = _dafny.Array(None, 23)
                nw72_[int(0)] = d_342_v2_
                nw72_[int(1)] = d_342_v2_
                nw72_[int(2)] = d_342_v2_
                nw72_[int(3)] = d_342_v2_
                nw72_[int(4)] = d_342_v2_
                nw72_[int(5)] = d_342_v2_
                nw72_[int(6)] = (d_342_v2_) + (d_342_v2_)
                nw72_[int(7)] = (d_342_v2_ if default__.fm0(d_418_i8_, d_418_i8_, p1, globalState) else _dafny.SeqWithoutIsStrInference([p2 for d_433_i11_ in range(default__.abs(-463))]))
                nw72_[int(8)] = (d_417_v49_)[default__.safeIndex(len(d_341_v1_), len(d_417_v49_))]
                nw72_[int(9)] = default__.fm12(p2, (self).f25, (d_340_v0_).set(d_418_i8_, (d_405_v40_)[default__.safeIndex(54, (d_405_v40_).length(0))]), p2, globalState)
                nw72_[int(10)] = (d_342_v2_) + (d_342_v2_)
                nw72_[int(11)] = d_342_v2_
                nw72_[int(12)] = (_dafny.SeqWithoutIsStrInference([d_418_i8_, p2])) + (d_342_v2_)
                nw72_[int(13)] = d_342_v2_
                nw72_[int(14)] = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_417_v49_)).cardinality, p2, (0) - (p2)])
                nw72_[int(15)] = (_dafny.SeqWithoutIsStrInference([p2])) + (d_342_v2_)
                nw72_[int(16)] = d_342_v2_
                nw72_[int(17)] = _dafny.SeqWithoutIsStrInference([len(default__.fm13(p1, p2, globalState)), 100, p2, p2])
                nw72_[int(18)] = d_342_v2_
                nw72_[int(19)] = (d_342_v2_) + (_dafny.SeqWithoutIsStrInference([d_418_i8_ for d_434_i12_ in range(default__.abs(-853))]))
                nw72_[int(20)] = _dafny.SeqWithoutIsStrInference([p2 for d_435_i13_ in range(default__.abs(237))])
                nw72_[int(21)] = (d_342_v2_) + (d_342_v2_)
                nw72_[int(22)] = default__.fm12(p2, not((self).f25), d_340_v0_, p2, globalState)
                d_432_v60_ = nw72_
                d_436_v61_: _dafny.Map
                d_436_v61_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([p2])})
                index75_ = default__.safeIndex(171, (d_432_v60_).length(0))
                (d_432_v60_)[index75_] = (((d_436_v61_)[d_418_i8_] if (d_418_i8_) in (d_436_v61_) else d_342_v2_) if (self).f25 else _dafny.SeqWithoutIsStrInference([d_418_i8_ for d_437_i14_ in range(default__.abs(3))]))
                d_438_v62_: _dafny.Map
                d_438_v62_ = _dafny.Map({(d_410_v42_).cardinality: 669})
                d_439_v63_: _dafny.Seq
                d_439_v63_ = _dafny.SeqWithoutIsStrInference([d_438_v62_, _dafny.Map({41: p2})])
                d_440_v64_: _dafny.Map
                d_440_v64_ = _dafny.Map({(self).f26: d_418_i8_})
                d_441_v65_: D2
                d_441_v65_ = D2_DC5(len(_dafny.Map({D2_DC6(D2_DC3(d_439_v63_)): p2})), (self).fm3(p1, d_418_i8_, p2, d_440_v64_, globalState), p1, (self).f25)
                index76_ = default__.safeIndex(171, (d_432_v60_).length(0))
                index77_ = default__.safeIndex(54, (d_405_v40_).length(0))
                rhs67_ = ((d_342_v2_) + (d_342_v2_)) + (_dafny.SeqWithoutIsStrInference([len(d_341_v1_), d_418_i8_]))
                rhs68_ = default__.fm2(p1, ((d_441_v65_).cf6) - (p2), not((self).f25), (self).f26, globalState)
                rhs69_ = d_418_i8_
                rhs70_ = ((self).f25 if (self).f25 else p1)
                lhs56_ = d_432_v60_
                lhs57_ = default__.safeIndex(171, (d_432_v60_).length(0))
                lhs58_ = globalState
                lhs59_ = globalState
                lhs60_ = d_405_v40_
                lhs61_ = default__.safeIndex(54, (d_405_v40_).length(0))
                lhs56_[lhs57_] = rhs67_
                lhs58_.f5 = rhs68_
                lhs59_.f22 = rhs69_
                lhs60_[lhs61_] = rhs70_
                index78_ = default__.safeIndex(54, (d_405_v40_).length(0))
                (d_405_v40_)[index78_] = (p1) == ((self).f25)
            d_442_v66_: int
            out6_: int
            out6_ = default__.m0(globalState)
            d_442_v66_ = out6_

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: int = int(0)
        if not(not(False)):
            (globalState).f21 = p1
            d_443_v0_: C0
            nw73_ = C0()
            nw73_.ctor__()
            d_443_v0_ = nw73_
            d_444_v1_: int
            out7_: int
            out7_ = default__.m0(globalState)
            d_444_v1_ = out7_
            rhs71_ = (self).f25
            rhs72_ = default__.safeModuloInt((d_443_v0_).fm5(globalState), default__.safeModuloInt(p1, d_444_v1_))
            lhs62_ = globalState
            lhs63_ = globalState
            lhs62_.f2 = rhs71_
            lhs63_.f5 = rhs72_
            if (self).f25:
                d_445_v2_: _dafny.Map
                d_445_v2_ = _dafny.Map({p1: p1})
                d_446_v3_: _dafny.Seq
                d_446_v3_ = _dafny.SeqWithoutIsStrInference([d_445_v2_, d_445_v2_])
                d_447_v4_: D2
                d_447_v4_ = D2_DC3(d_446_v3_)
                d_448_v5_: _dafny.Map
                d_448_v5_ = _dafny.Map({d_447_v4_: p1})
                d_448_v5_ = ((d_448_v5_) | (d_448_v5_)) | (d_448_v5_)
                d_449_v6_: _dafny.Map
                d_449_v6_ = _dafny.Map({(self).f26: d_444_v1_})
                d_450_v7_: _dafny.Array
                nw74_ = _dafny.Array(None, 16)
                nw74_[int(0)] = True
                nw74_[int(1)] = True
                nw74_[int(2)] = False
                nw74_[int(3)] = not((self).f25)
                nw74_[int(4)] = (self).f25
                nw74_[int(5)] = (self).f25
                nw74_[int(6)] = (self).fm3((self).f25, p1, p1, d_449_v6_, globalState)
                nw74_[int(7)] = (self).f25
                nw74_[int(8)] = (self).f25
                nw74_[int(9)] = (self).f25
                nw74_[int(10)] = (self).f25
                nw74_[int(11)] = (self).f25
                nw74_[int(12)] = not((self).f25)
                nw74_[int(13)] = False
                nw74_[int(14)] = (self).f25
                nw74_[int(15)] = (self).f25
                d_450_v7_ = nw74_
                d_451_v8_: _dafny.MultiSet
                d_451_v8_ = _dafny.MultiSet([d_450_v7_, d_450_v7_, d_450_v7_, d_450_v7_, d_450_v7_])
                d_452_v9_: _dafny.Map
                d_452_v9_ = _dafny.Map({(self).f25: d_451_v8_})
                r0 = (d_450_v7_) not in (((d_452_v9_)[not((self).f25)] if (not((self).f25)) in (d_452_v9_) else _dafny.MultiSet([d_450_v7_, d_450_v7_, d_450_v7_, d_450_v7_, d_450_v7_])))
                (globalState).f22 = default__.fm2((self).f25, p1, (self).f25, (self).f26, globalState)
                d_453_v10_: _dafny.Map
                d_453_v10_ = _dafny.Map({(self).f25: p1})
                d_454_v11_: _dafny.Set
                d_454_v11_ = _dafny.Set({default__.fm0(p1, p1, (self).f25, globalState)})
                d_455_v12_: _dafny.Seq
                d_455_v12_ = _dafny.SeqWithoutIsStrInference([d_444_v1_])
                d_456_v13_: int
                d_456_v13_ = p1
                d_457_v14_: _dafny.Array
                nw75_ = _dafny.Array(None, 26)
                nw75_[int(0)] = d_444_v1_
                nw75_[int(1)] = (((d_453_v10_)[(self).f25] if ((self).f25) in (d_453_v10_) else (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f25]))))) + (d_444_v1_)
                nw75_[int(2)] = len(p0)
                nw75_[int(3)] = (d_443_v0_).fm5(globalState)
                nw75_[int(4)] = 187
                nw75_[int(5)] = d_444_v1_
                nw75_[int(6)] = 184
                nw75_[int(7)] = d_444_v1_
                nw75_[int(8)] = d_444_v1_
                nw75_[int(9)] = default__.safeDivisionInt(d_444_v1_, p1)
                nw75_[int(10)] = (d_444_v1_ if (self).f25 else d_444_v1_)
                nw75_[int(11)] = (p1) + (len(p2))
                nw75_[int(12)] = len(d_454_v11_)
                nw75_[int(13)] = 940
                nw75_[int(14)] = p1
                nw75_[int(15)] = p1
                nw75_[int(16)] = d_444_v1_
                nw75_[int(17)] = (d_443_v0_).fm5(globalState)
                nw75_[int(18)] = d_444_v1_
                nw75_[int(19)] = d_444_v1_
                nw75_[int(20)] = d_444_v1_
                nw75_[int(21)] = 676
                nw75_[int(22)] = (d_455_v12_)[default__.safeIndex(-626, len(d_455_v12_))]
                nw75_[int(23)] = 419
                nw75_[int(24)] = len(default__.fm12(d_444_v1_, (self).f25, _dafny.Map({d_444_v1_: (self).f25}), p1, globalState))
                nw75_[int(25)] = (d_456_v13_)
                d_457_v14_ = nw75_
                index79_ = default__.safeIndex(84, (d_457_v14_).length(0))
                (d_457_v14_)[index79_] = d_444_v1_
                d_458_v15_: _dafny.MultiSet
                d_458_v15_ = _dafny.MultiSet([(self).fm3((self).f25, p1, d_444_v1_, d_449_v6_, globalState)])
                index80_ = default__.safeIndex(84, (d_457_v14_).length(0))
                (d_457_v14_)[index80_] = default__.safeDivisionInt(((d_458_v15_)[(self).f25] if ((self).f25) in (d_458_v15_) else p1), (d_444_v1_) - (len(d_453_v10_)))
                d_459_v16_: C0
                nw76_ = C0()
                nw76_.ctor__()
                d_459_v16_ = nw76_
            elif True:
                (globalState).f2 = (d_444_v1_) < (p1)
                d_460_v17_: _dafny.Array
                nw77_ = _dafny.Array(int(0), 26)
                d_460_v17_ = nw77_
                d_460_v17_ = d_460_v17_
                d_461_v18_: _dafny.Set
                d_461_v18_ = _dafny.Set({p1, p1})
                d_462_v19_: _dafny.MultiSet
                d_462_v19_ = _dafny.MultiSet([len(d_461_v18_)])
                d_463_v20_: _dafny.Seq
                d_463_v20_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f25, (p1) <= ((d_462_v19_).cardinality)])
                d_463_v20_ = d_463_v20_
                d_464_v21_: _dafny.Array
                nw78_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_464_v21_ = nw78_
                index81_ = default__.safeIndex(8, (d_464_v21_).length(0))
                (d_464_v21_)[index81_] = d_460_v17_
                d_465_v22_: D7
                d_465_v22_ = D7_DC20(_dafny.SeqWithoutIsStrInference([(self).f25, (self).f25, (self).f25, not((self).f25)]))
                d_466_v23_: T0
                nw79_ = C1()
                nw79_.ctor__((p0) + (p0), len((d_465_v22_).cf37), False)
                d_466_v23_ = nw79_
                d_467_v24_: _dafny.Seq
                d_467_v24_ = _dafny.SeqWithoutIsStrInference([d_443_v0_, d_443_v0_])
                d_468_v25_: _dafny.Array
                nw80_ = _dafny.Array(None, 20)
                nw80_[int(0)] = d_467_v24_
                nw80_[int(1)] = d_467_v24_
                nw80_[int(2)] = (d_467_v24_).set(default__.safeIndex(d_444_v1_, len(d_467_v24_)), d_443_v0_)
                nw80_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_443_v0_, (d_467_v24_)[default__.safeIndex(p1, len(d_467_v24_))]])) + (d_467_v24_)
                nw80_[int(4)] = d_467_v24_
                nw80_[int(5)] = d_467_v24_
                nw80_[int(6)] = d_467_v24_
                nw80_[int(7)] = d_467_v24_
                nw80_[int(8)] = d_467_v24_
                nw80_[int(9)] = _dafny.SeqWithoutIsStrInference([d_443_v0_])
                nw80_[int(10)] = _dafny.SeqWithoutIsStrInference([d_443_v0_, d_443_v0_])
                nw80_[int(11)] = d_467_v24_
                nw80_[int(12)] = d_467_v24_
                nw80_[int(13)] = d_467_v24_
                nw80_[int(14)] = d_467_v24_
                nw80_[int(15)] = _dafny.SeqWithoutIsStrInference([d_443_v0_, d_443_v0_])
                nw80_[int(16)] = d_467_v24_
                nw80_[int(17)] = _dafny.SeqWithoutIsStrInference([d_443_v0_])
                nw80_[int(18)] = (d_467_v24_) + (_dafny.SeqWithoutIsStrInference([d_443_v0_]))
                nw80_[int(19)] = d_467_v24_
                d_468_v25_ = nw80_
                index82_ = default__.safeIndex(997, (d_468_v25_).length(0))
                (d_468_v25_)[index82_] = _dafny.SeqWithoutIsStrInference([d_443_v0_, d_443_v0_])
                d_469_v26_: _dafny.Set
                d_469_v26_ = _dafny.Set({_dafny.CodePoint('d')})
                index83_ = default__.safeIndex(8, (d_464_v21_).length(0))
                index84_ = default__.safeIndex(997, (d_468_v25_).length(0))
                rhs73_ = d_460_v17_
                rhs74_ = d_466_v23_
                rhs75_ = (d_467_v24_).set(default__.safeIndex(len(d_469_v26_), len(d_467_v24_)), d_443_v0_)
                rhs76_ = (d_466_v23_).f25
                lhs64_ = d_464_v21_
                lhs65_ = default__.safeIndex(8, (d_464_v21_).length(0))
                lhs66_ = d_468_v25_
                lhs67_ = default__.safeIndex(997, (d_468_v25_).length(0))
                lhs68_ = globalState
                lhs64_[lhs65_] = rhs73_
                d_466_v23_ = rhs74_
                lhs66_[lhs67_] = rhs75_
                lhs68_.f2 = rhs76_
                d_470_v27_: _dafny.Array
                nw81_ = _dafny.Array(None, 11)
                nw81_[int(0)] = (p1) >= (d_444_v1_)
                nw81_[int(1)] = (d_466_v23_).f25
                nw81_[int(2)] = (default__.fm22(d_444_v1_, d_444_v1_, p1, (self).f25, globalState)).issubset(d_469_v26_)
                nw81_[int(3)] = not((d_466_v23_).f25)
                nw81_[int(4)] = (self).f25
                nw81_[int(5)] = not((self).f25)
                nw81_[int(6)] = (d_466_v23_).f25
                nw81_[int(7)] = (p2)[default__.safeIndex((0) - (d_444_v1_), len(p2))]
                nw81_[int(8)] = (d_466_v23_).f25
                nw81_[int(9)] = ((self).f25) == ((self).f25)
                nw81_[int(10)] = (self).f25
                d_470_v27_ = nw81_
                index85_ = default__.safeIndex(685, (d_470_v27_).length(0))
                (d_470_v27_)[index85_] = (d_466_v23_).f25
                index86_ = default__.safeIndex(685, (d_470_v27_).length(0))
                (d_470_v27_)[index86_] = (self).f25
        elif True:
            (globalState).f5 = (p1 if (self).f25 else (_dafny.MultiSet([not((self).f25)])).cardinality)
            d_471_v28_: C1
            nw82_ = C1()
            nw82_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")), p1, (self).f25)
            d_471_v28_ = nw82_
            d_472_v29_: _dafny.Array
            def lambda25_(d_473_i0_):
                return (d_473_i0_) - (586)

            init14_ = lambda25_
            nw83_ = _dafny.Array(None, 12)
            for i0_14_ in range(nw83_.length(0)):
                nw83_[i0_14_] = init14_(i0_14_)
            d_472_v29_ = nw83_
            index87_ = default__.safeIndex(144, (d_472_v29_).length(0))
            (d_472_v29_)[index87_] = d_471_v28_.f29
            index88_ = default__.safeIndex(144, (d_472_v29_).length(0))
            (d_472_v29_)[index88_] = p1
            d_474_v30_: _dafny.Map
            d_474_v30_ = _dafny.Map({not((self).f25): (self).f25})
            d_475_v31_: D7
            d_475_v31_ = D7_DC21((d_472_v29_)[default__.safeIndex(144, (d_472_v29_).length(0))], (d_474_v30_) | ((d_474_v30_).set(not(True), (self).f25)))
            d_475_v31_ = default__.fm23(not(not((self).f25)), globalState)
            d_476_v32_: C0
            nw84_ = C0()
            nw84_.ctor__()
            d_476_v32_ = nw84_
        d_477_v33_: int
        d_477_v33_ = p1
        source6_ = D7_DC22(p1, d_477_v33_, (self).f25)
        if source6_.is_DC21:
            d_478___mcc_h0_ = source6_.cf38
            d_479___mcc_h1_ = source6_.cf39
            d_480_cf39_ = d_479___mcc_h1_
            d_481_cf38_ = d_478___mcc_h0_
            d_482_v34_: _dafny.Map
            d_482_v34_ = _dafny.Map({(0) - (p1): (self).f26})
            d_483_v35_: D5
            d_483_v35_ = D5_DC14(d_482_v34_)
            d_484_v36_: D5
            d_484_v36_ = D5_DC16(d_483_v35_)
            d_485_v37_: _dafny.Set
            d_485_v37_ = _dafny.Set({(self).f25, (self).f25})
            d_484_v36_ = (d_484_v36_ if not((d_485_v37_).ispropersubset(d_485_v37_)) else d_484_v36_)
            (globalState).f21 = (p1) + (p1)
            d_486_v38_: _dafny.Array
            def lambda26_(d_487_i1_):
                return (self).f25

            init15_ = lambda26_
            nw85_ = _dafny.Array(None, 17)
            for i0_15_ in range(nw85_.length(0)):
                nw85_[i0_15_] = init15_(i0_15_)
            d_486_v38_ = nw85_
            d_488_v39_: _dafny.Set
            d_488_v39_ = _dafny.Set({p1})
            index89_ = default__.safeIndex(962, (d_486_v38_).length(0))
            (d_486_v38_)[index89_] = (d_488_v39_).ispropersubset(d_488_v39_)
            d_489_v40_: _dafny.Map
            d_489_v40_ = _dafny.Map({d_481_cf38_: (self).f25})
            index90_ = default__.safeIndex(962, (d_486_v38_).length(0))
            (d_486_v38_)[index90_] = not ((((d_489_v40_)[d_481_cf38_] if (d_481_cf38_) in (d_489_v40_) else (self).f25)) or (False)) or ((self).f25)
            d_490_v41_: _dafny.Seq
            d_490_v41_ = _dafny.SeqWithoutIsStrInference([len(p2), p1])
            d_491_v42_: _dafny.Map
            d_491_v42_ = _dafny.Map({(self).f26: not((d_486_v38_)[default__.safeIndex(962, (d_486_v38_).length(0))])})
            d_492_v43_: _dafny.Array
            nw86_ = _dafny.Array(None, 12)
            nw86_[int(0)] = (p1 if (self).f25 else p1)
            nw86_[int(1)] = d_481_cf38_
            nw86_[int(2)] = d_481_cf38_
            nw86_[int(3)] = default__.safeModuloInt(len(_dafny.Set({(d_486_v38_)[default__.safeIndex(962, (d_486_v38_).length(0))], (p2)[default__.safeIndex((0) - (d_481_cf38_), len(p2))]})), p1)
            nw86_[int(4)] = d_481_cf38_
            nw86_[int(5)] = default__.safeModuloInt(97, d_481_cf38_)
            nw86_[int(6)] = (0) - (((d_490_v41_)[default__.safeIndex(d_481_cf38_, len(d_490_v41_))] if (self).f25 else p1))
            nw86_[int(7)] = d_481_cf38_
            nw86_[int(8)] = p1
            nw86_[int(9)] = (0) - ((0) - (len((d_491_v42_) | (d_491_v42_))))
            nw86_[int(10)] = (D3_DC9(p1, (self).f26, len(p0))).cf15
            nw86_[int(11)] = default__.safeDivisionInt(p1, p1)
            d_492_v43_ = nw86_
            index91_ = default__.safeIndex(435, (d_492_v43_).length(0))
            (d_492_v43_)[index91_] = (d_490_v41_)[default__.safeIndex(p1, len(d_490_v41_))]
            d_493_v44_: _dafny.MultiSet
            d_493_v44_ = _dafny.MultiSet([(d_486_v38_)[default__.safeIndex(962, (d_486_v38_).length(0))], (self).f25])
            d_494_v45_: _dafny.Map
            d_494_v45_ = _dafny.Map({(self).f25: (d_493_v44_).cardinality})
            index92_ = default__.safeIndex(435, (d_492_v43_).length(0))
            rhs77_ = default__.safeModuloInt(len(d_494_v45_), default__.safeModuloInt(-466, d_481_cf38_))
            rhs78_ = _dafny.Map({default__.safeDivisionInt(p1, d_481_cf38_): (self).f25})
            lhs69_ = d_492_v43_
            lhs70_ = default__.safeIndex(435, (d_492_v43_).length(0))
            lhs69_[lhs70_] = rhs77_
            d_489_v40_ = rhs78_
        elif source6_.is_DC22:
            d_495___mcc_h2_ = source6_.cf40
            d_496___mcc_h3_ = source6_.cf41
            d_497___mcc_h4_ = source6_.cf42
            d_498_cf42_ = d_497___mcc_h4_
            d_499_cf41_ = d_496___mcc_h3_
            d_500_cf40_ = d_495___mcc_h2_
            d_501_v46_: C0
            nw87_ = C0()
            nw87_.ctor__()
            d_501_v46_ = nw87_
            d_502_v47_: int
            out8_: int
            out8_ = default__.m0(globalState)
            d_502_v47_ = out8_
            d_503_v48_: _dafny.Array
            def lambda27_(d_504_p0_):
                def lambda28_(d_505_i2_):
                    return d_504_p0_

                return lambda28_

            init16_ = lambda27_(p0)
            nw88_ = _dafny.Array(None, 13)
            for i0_16_ in range(nw88_.length(0)):
                nw88_[i0_16_] = init16_(i0_16_)
            d_503_v48_ = nw88_
            d_506_v49_: _dafny.Map
            d_506_v49_ = _dafny.Map({(self).f25: d_503_v48_})
            d_507_v50_: _dafny.Map
            d_507_v50_ = _dafny.Map({d_498_cf42_: (self).f26})
            d_508_v51_: _dafny.Seq
            d_508_v51_ = _dafny.SeqWithoutIsStrInference([d_507_v50_, d_507_v50_])
            d_506_v49_ = (d_506_v49_).set((-937) == (len(d_508_v51_)), d_503_v48_)
            d_509_v52_: D3
            d_509_v52_ = D3_DC8(not((self).f25))
            (globalState).f2 = not ((d_509_v52_).cf12) or ((p1) != (d_502_v47_))
        elif source6_.is_DC23:
            d_510___mcc_h5_ = source6_.cf43
            d_511___mcc_h6_ = source6_.cf44
            d_512___mcc_h7_ = source6_.cf45
            d_513___mcc_h8_ = source6_.cf46
            d_514_cf46_ = d_513___mcc_h8_
            d_515_cf45_ = d_512___mcc_h7_
            d_516_cf44_ = d_511___mcc_h6_
            d_517_cf43_ = d_510___mcc_h5_
            d_518_v53_: _dafny.MultiSet
            d_518_v53_ = _dafny.MultiSet([(self).f25])
            d_515_cf45_ = (d_518_v53_).ispropersubset(d_518_v53_)
            if not(d_515_cf45_):
                (globalState).f18 = (d_516_cf44_ if d_515_cf45_ else (0) - (p1))
                d_519_v54_: D2
                d_519_v54_ = D2_DC5(-739, (self).f25, d_515_cf45_, default__.fm0(d_516_cf44_, d_517_cf43_, True, globalState))
                d_520_v55_: _dafny.Map
                d_520_v55_ = _dafny.Map({(self).f25: (0) - (d_516_cf44_)})
                d_521_v56_: _dafny.Map
                d_521_v56_ = _dafny.Map({d_515_cf45_: (self).f25})
                pat_let_tv9_ = d_515_cf45_
                d_522_v57_: _dafny.Array
                nw89_ = _dafny.Array(None, 18)
                nw89_[int(0)] = d_519_v54_
                nw89_[int(1)] = d_519_v54_
                nw89_[int(2)] = d_519_v54_
                nw89_[int(3)] = D2_DC5(len(d_520_v55_), (self).f25, ((d_521_v56_)[(self).f25] if ((self).f25) in (d_521_v56_) else (self).f25), d_515_cf45_)
                nw89_[int(4)] = d_519_v54_
                def iife29_(_pat_let4_0):
                    def iife30_(d_523_dt__update__tmp_h0_):
                        def iife31_(_pat_let5_0):
                            def iife32_(d_524_dt__update_hcf7_h0_):
                                def iife33_(_pat_let6_0):
                                    def iife34_(d_525_dt__update_hcf8_h0_):
                                        return D2_DC5((d_523_dt__update__tmp_h0_).cf6, d_524_dt__update_hcf7_h0_, d_525_dt__update_hcf8_h0_, (d_523_dt__update__tmp_h0_).cf9)
                                    return iife34_(_pat_let6_0)
                                return iife33_((self).f25)
                            return iife32_(_pat_let5_0)
                        return iife31_(pat_let_tv9_)
                    return iife30_(_pat_let4_0)
                nw89_[int(5)] = iife29_(d_519_v54_)
                nw89_[int(6)] = d_519_v54_
                nw89_[int(7)] = d_519_v54_
                nw89_[int(8)] = d_519_v54_
                nw89_[int(9)] = d_519_v54_
                nw89_[int(10)] = d_519_v54_
                nw89_[int(11)] = d_519_v54_
                nw89_[int(12)] = d_519_v54_
                nw89_[int(13)] = d_519_v54_
                nw89_[int(14)] = d_519_v54_
                nw89_[int(15)] = d_519_v54_
                nw89_[int(16)] = d_519_v54_
                nw89_[int(17)] = d_519_v54_
                d_522_v57_ = nw89_
                index93_ = default__.safeIndex(899, (d_522_v57_).length(0))
                (d_522_v57_)[index93_] = d_519_v54_
                index94_ = default__.safeIndex(899, (d_522_v57_).length(0))
                (d_522_v57_)[index94_] = d_519_v54_
                r0 = (self).f25
                d_526_v58_: _dafny.Array
                nw90_ = _dafny.Array(_dafny.Set({}), 7)
                d_526_v58_ = nw90_
                d_527_v59_: C0
                nw91_ = C0()
                nw91_.ctor__()
                d_527_v59_ = nw91_
                d_528_v60_: _dafny.Set
                d_528_v60_ = _dafny.Set({d_527_v59_, d_527_v59_})
                index95_ = default__.safeIndex(351, (d_526_v58_).length(0))
                (d_526_v58_)[index95_] = d_528_v60_
                index96_ = default__.safeIndex(351, (d_526_v58_).length(0))
                (d_526_v58_)[index96_] = _dafny.Set({d_527_v59_})
                d_529_v61_: C0
                nw92_ = C0()
                nw92_.ctor__()
                d_529_v61_ = nw92_
            elif True:
                d_530_v62_: _dafny.Seq
                d_530_v62_ = _dafny.SeqWithoutIsStrInference([d_514_cf46_, d_514_cf46_])
                d_531_v63_: _dafny.Map
                d_531_v63_ = _dafny.Map({d_514_cf46_: default__.safeModuloInt((d_530_v62_)[default__.safeIndex(len(p0), len(d_530_v62_))], p1)})
                d_531_v63_ = (d_531_v63_).set((d_514_cf46_) - (default__.fm2(False, d_516_cf44_, d_515_cf45_, (self).f26, globalState)), d_514_cf46_)
                d_532_v64_: C1
                nw93_ = C1()
                nw93_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awhlibqi")), (d_514_cf46_) - (p1), d_515_cf45_)
                d_532_v64_ = nw93_
                d_515_cf45_ = d_515_cf45_
                (globalState).f22 = (default__.safeDivisionInt(d_514_cf46_, 160)) * (d_532_v64_.f29)
                d_533_v66_: _dafny.Map
                d_533_v66_ = _dafny.Map({d_516_cf44_: (self).f25})
                d_534_v67_: D6
                def iife35_():
                    coll21_ = _dafny.Map()
                    compr_21_: int
                    for compr_21_ in (d_533_v66_).keys.Elements:
                        d_535_v65_: int = compr_21_
                        if (d_535_v65_) in (d_533_v66_):
                            coll21_[(d_535_v65_) - (d_516_cf44_)] = d_516_cf44_
                    return _dafny.Map(coll21_)
                d_534_v67_ = D6_DC17(iife35_()
)
                d_536_v68_: D6
                d_536_v68_ = D6_DC19(d_534_v67_)
                d_536_v68_ = d_536_v68_
            d_537_v69_: _dafny.Array
            def lambda29_(d_538_i3_):
                return (self).f26

            init17_ = lambda29_
            nw94_ = _dafny.Array(None, 17)
            for i0_17_ in range(nw94_.length(0)):
                nw94_[i0_17_] = init17_(i0_17_)
            d_537_v69_ = nw94_
            index97_ = default__.safeIndex(120, (d_537_v69_).length(0))
            (d_537_v69_)[index97_] = (self).f26
            index98_ = default__.safeIndex(471, (d_537_v69_).length(0))
            (d_537_v69_)[index98_] = (self).f26
            d_539_v70_: _dafny.Set
            d_539_v70_ = _dafny.Set({not(d_515_cf45_)})
            index99_ = default__.safeIndex(120, (d_537_v69_).length(0))
            index100_ = default__.safeIndex(471, (d_537_v69_).length(0))
            rhs79_ = (169) + (d_517_cf43_)
            rhs80_ = (self).f26
            rhs81_ = _dafny.CodePoint('n')
            rhs82_ = len((d_539_v70_).intersection(d_539_v70_))
            rhs83_ = p1
            lhs71_ = globalState
            lhs72_ = d_537_v69_
            lhs73_ = default__.safeIndex(120, (d_537_v69_).length(0))
            lhs74_ = d_537_v69_
            lhs75_ = default__.safeIndex(471, (d_537_v69_).length(0))
            lhs71_.f22 = rhs79_
            lhs72_[lhs73_] = rhs80_
            lhs74_[lhs75_] = rhs81_
            r1 = rhs82_
            d_516_cf44_ = rhs83_
            (globalState).f18 = (p1) + (d_514_cf46_)
        elif source6_.is_DC20:
            d_540___mcc_h9_ = source6_.cf37
            d_541_cf37_ = d_540___mcc_h9_
            (globalState).f2 = (self).f25
            if not(True):
                d_542_v71_: _dafny.MultiSet
                d_542_v71_ = _dafny.MultiSet([p1, p1, 988, p1, p1])
                (globalState).f2 = default__.fm0((610) * (len(_dafny.SeqWithoutIsStrInference([p1 for d_543_i4_ in range(default__.abs(763))]))), default__.fm2((self).f25, (0) - (((d_542_v71_).set(909, default__.abs(p1))).cardinality), not((self).f25), (self).f26, globalState), not(False), globalState)
                (globalState).f21 = p1
                (globalState).f2 = (self).f25
                d_544_v72_: _dafny.Array
                def lambda30_(d_545_i5_):
                    return (_dafny.Map({(self).f25: (self).f25})) | (_dafny.Map({(self).f25: (self).f25}))

                init18_ = lambda30_
                nw95_ = _dafny.Array(None, 9)
                for i0_18_ in range(nw95_.length(0)):
                    nw95_[i0_18_] = init18_(i0_18_)
                d_544_v72_ = nw95_
                d_544_v72_ = d_544_v72_
                d_546_v73_: _dafny.Array
                nw96_ = _dafny.Array(int(0), 18)
                d_546_v73_ = nw96_
                index101_ = default__.safeIndex(886, (d_546_v73_).length(0))
                (d_546_v73_)[index101_] = p1
                index102_ = default__.safeIndex(886, (d_546_v73_).length(0))
                (d_546_v73_)[index102_] = p1
            elif True:
                d_547_v74_: _dafny.Array
                def lambda31_(d_548_i6_):
                    return not ((self).f25) or ((self).f25)

                init19_ = lambda31_
                nw97_ = _dafny.Array(None, 5)
                for i0_19_ in range(nw97_.length(0)):
                    nw97_[i0_19_] = init19_(i0_19_)
                d_547_v74_ = nw97_
                index103_ = default__.safeIndex(756, (d_547_v74_).length(0))
                (d_547_v74_)[index103_] = (self).f25
                index104_ = default__.safeIndex(756, (d_547_v74_).length(0))
                (d_547_v74_)[index104_] = (self).f25
                index105_ = default__.safeIndex(756, (d_547_v74_).length(0))
                (d_547_v74_)[index105_] = (default__.fm16(p1, globalState)) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqccefcs")))
                d_549_v75_: _dafny.Seq
                d_549_v75_ = _dafny.SeqWithoutIsStrInference([p2, p2, d_541_cf37_, (_dafny.SeqWithoutIsStrInference([(self).f25])) + (p2), p2])
                d_550_v76_: D7
                d_550_v76_ = D7_DC20(d_541_cf37_)
                d_551_v77_: _dafny.Map
                d_551_v77_ = _dafny.Map({p1: (d_547_v74_)[default__.safeIndex(756, (d_547_v74_).length(0))]})
                rhs84_ = (d_547_v74_)[default__.safeIndex(756, (d_547_v74_).length(0))]
                rhs85_ = (len(p0)) - ((0) - ((p1) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({p0: d_550_v76_}))])))))
                rhs86_ = (len(p0) if (self).f25 else p1)
                rhs87_ = d_549_v75_
                rhs88_ = len(d_551_v77_)
                lhs76_ = globalState
                lhs77_ = globalState
                lhs78_ = globalState
                lhs76_.f2 = rhs84_
                r1 = rhs85_
                lhs77_.f22 = rhs86_
                d_549_v75_ = rhs87_
                lhs78_.f22 = rhs88_
                d_541_cf37_ = p2
                d_552_v78_: _dafny.Map
                d_552_v78_ = _dafny.Map({(self).f25: p1})
                d_553_v79_: _dafny.MultiSet
                d_553_v79_ = _dafny.MultiSet([(d_547_v74_)[default__.safeIndex(756, (d_547_v74_).length(0))], (d_547_v74_)[default__.safeIndex(756, (d_547_v74_).length(0))], (d_547_v74_)[default__.safeIndex(756, (d_547_v74_).length(0))]])
                d_554_v80_: _dafny.MultiSet
                d_554_v80_ = _dafny.MultiSet([(d_552_v78_) | (_dafny.Map({(d_547_v74_)[default__.safeIndex(756, (d_547_v74_).length(0))]: (d_553_v79_).cardinality}))])
                d_554_v80_ = (d_554_v80_) | (d_554_v80_)
            d_555_v81_: _dafny.Seq
            d_555_v81_ = _dafny.SeqWithoutIsStrInference([p1])
            d_556_v82_: _dafny.MultiSet
            d_556_v82_ = _dafny.MultiSet([(0) - (len(p0))])
            (globalState).f2 = (d_556_v82_).issubset(_dafny.MultiSet((d_555_v81_) + (d_555_v81_)))
            d_557_v83_: _dafny.Map
            d_557_v83_ = _dafny.Map({612: (self).f26})
            d_558_v84_: D5
            d_558_v84_ = D5_DC16(D5_DC14(d_557_v83_))
            d_559_v85_: _dafny.MultiSet
            d_559_v85_ = _dafny.MultiSet([p2, d_541_cf37_, _dafny.SeqWithoutIsStrInference([(self).f25])])
            rhs89_ = (d_558_v84_ if (self).f25 else d_558_v84_)
            rhs90_ = (p1) <= (((d_559_v85_).set(p2, default__.abs(p1))).cardinality)
            rhs91_ = d_477_v33_
            d_558_v84_ = rhs89_
            r0 = rhs90_
            d_477_v33_ = rhs91_
        elif True:
            d_560___mcc_h10_ = source6_.cf47
            d_561_cf47_ = d_560___mcc_h10_
            d_562_v86_: D2
            d_562_v86_ = D2_DC5(p1, (self).f25, not((self).f25), (self).f25)
            d_563_v87_: D2
            d_563_v87_ = D2_DC6(d_562_v86_)
            d_564_v88_: D2
            d_564_v88_ = D2_DC6(d_563_v87_)
            pat_let_tv10_ = d_563_v87_
            def iife36_(_pat_let7_0):
                def iife37_(d_565_dt__update__tmp_h1_):
                    def iife38_(_pat_let8_0):
                        def iife39_(d_566_dt__update_hcf10_h0_):
                            return D2_DC6(d_566_dt__update_hcf10_h0_)
                        return iife39_(_pat_let8_0)
                    return iife38_(pat_let_tv10_)
                return iife37_(_pat_let7_0)
            d_564_v88_ = iife36_(d_564_v88_)
            d_567_v89_: _dafny.Array
            def lambda32_(d_568_i7_):
                return (d_568_i7_) * (741)

            init20_ = lambda32_
            nw98_ = _dafny.Array(None, 18)
            for i0_20_ in range(nw98_.length(0)):
                nw98_[i0_20_] = init20_(i0_20_)
            d_567_v89_ = nw98_
            index106_ = default__.safeIndex(657, (d_567_v89_).length(0))
            (d_567_v89_)[index106_] = p1
            d_569_v90_: D7
            d_569_v90_ = D7_DC22(len(p0), d_477_v33_, (self).f25)
            index107_ = default__.safeIndex(854, (d_567_v89_).length(0))
            (d_567_v89_)[index107_] = default__.safeDivisionInt((d_569_v90_).cf40, p1)
            d_570_v91_: _dafny.MultiSet
            d_570_v91_ = _dafny.MultiSet([p1, p1, p1, p1])
            d_571_v92_: _dafny.Seq
            d_571_v92_ = _dafny.SeqWithoutIsStrInference([p1, len(default__.fm24((self).f25, p1, p1, globalState))])
            index108_ = default__.safeIndex(657, (d_567_v89_).length(0))
            index109_ = default__.safeIndex(854, (d_567_v89_).length(0))
            rhs92_ = p1
            rhs93_ = p1
            rhs94_ = (_dafny.MultiSet([len(d_571_v92_), p1, p1])).ispropersubset(d_570_v91_)
            lhs79_ = d_567_v89_
            lhs80_ = default__.safeIndex(657, (d_567_v89_).length(0))
            lhs81_ = d_567_v89_
            lhs82_ = default__.safeIndex(854, (d_567_v89_).length(0))
            lhs83_ = globalState
            lhs79_[lhs80_] = rhs92_
            lhs81_[lhs82_] = rhs93_
            lhs83_.f2 = rhs94_
            (globalState).f2 = (self).f25
            d_572_v93_: _dafny.Set
            d_572_v93_ = _dafny.Set({not((self).f25)})
            d_573_v94_: _dafny.Map
            d_573_v94_ = _dafny.Map({d_572_v93_: (self).f25})
            d_573_v94_ = (d_573_v94_).set(d_572_v93_, (p1) != (p1))
        d_574_v95_: _dafny.Seq
        d_574_v95_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
        hi5_ = len(d_574_v95_)
        for d_575_i8_ in range(p1, hi5_):
            d_576_v96_: C0
            nw99_ = C0()
            nw99_.ctor__()
            d_576_v96_ = nw99_
            d_577_v97_: _dafny.Map
            d_577_v97_ = _dafny.Map({(self).f25: True})
            d_578_v98_: _dafny.Map
            d_578_v98_ = _dafny.Map({not(((d_577_v97_)[False] if (False) in (d_577_v97_) else (self).f25)): (self).f25})
            d_578_v98_ = (d_578_v98_).set(True, (self).f25)
            d_579_v99_: _dafny.Seq
            d_579_v99_ = _dafny.SeqWithoutIsStrInference([d_574_v95_])
            (globalState).f19 = (d_574_v95_) + ((d_579_v99_)[default__.safeIndex(d_575_i8_, len(d_579_v99_))])
            d_580_v100_: _dafny.Array
            nw100_ = _dafny.Array(False, 7)
            d_580_v100_ = nw100_
            index110_ = default__.safeIndex(269, (d_580_v100_).length(0))
            (d_580_v100_)[index110_] = (p1) == (p1)
            index111_ = default__.safeIndex(747, (d_580_v100_).length(0))
            (d_580_v100_)[index111_] = True
            index112_ = default__.safeIndex(269, (d_580_v100_).length(0))
            index113_ = default__.safeIndex(747, (d_580_v100_).length(0))
            rhs95_ = (self).f25
            rhs96_ = (self).f25
            rhs97_ = (p2)[default__.safeIndex(p1, len(p2))]
            rhs98_ = p1
            rhs99_ = d_575_i8_
            lhs84_ = d_580_v100_
            lhs85_ = default__.safeIndex(269, (d_580_v100_).length(0))
            lhs86_ = d_580_v100_
            lhs87_ = default__.safeIndex(747, (d_580_v100_).length(0))
            lhs88_ = globalState
            lhs89_ = globalState
            r0 = rhs95_
            lhs84_[lhs85_] = rhs96_
            lhs86_[lhs87_] = rhs97_
            lhs88_.f22 = rhs98_
            lhs89_.f22 = rhs99_
        d_581_v101_: _dafny.MultiSet
        d_581_v101_ = _dafny.MultiSet([p1])
        d_582_v102_: _dafny.Map
        d_582_v102_ = _dafny.Map({(self).f25: (self).f25})
        d_583_v103_: D7
        d_583_v103_ = D7_DC21((0) - (p1), d_582_v102_)
        d_584_v104_: _dafny.Array
        nw101_ = _dafny.Array(None, 24)
        nw101_[int(0)] = (self).f25
        nw101_[int(1)] = (p1) != (-887)
        nw101_[int(2)] = (d_581_v101_).issubset((d_581_v101_).set(p1, default__.abs((d_583_v103_).cf38)))
        nw101_[int(3)] = not((self).f25)
        nw101_[int(4)] = ((0) - (p1)) < (p1)
        nw101_[int(5)] = (self).f25
        nw101_[int(6)] = (not((p2)[default__.safeIndex(p1, len(p2))])) and ((self).f25)
        nw101_[int(7)] = (self).f25
        nw101_[int(8)] = (self).f25
        nw101_[int(9)] = (self).f25
        nw101_[int(10)] = (self).f25
        nw101_[int(11)] = (self).f25
        nw101_[int(12)] = (self).f25
        nw101_[int(13)] = (p1) == (p1)
        nw101_[int(14)] = ((self).f26) not in (p0)
        nw101_[int(15)] = (d_581_v101_).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([254 for d_585_i9_ in range(default__.abs(-374))])))
        nw101_[int(16)] = (self).f25
        nw101_[int(17)] = (p2)[default__.safeIndex(p1, len(p2))]
        nw101_[int(18)] = (self).f25
        nw101_[int(19)] = False
        nw101_[int(20)] = ((self).f25 if (self).f25 else (self).f25)
        nw101_[int(21)] = (self).f25
        nw101_[int(22)] = (self).f25
        nw101_[int(23)] = (self).f25
        d_584_v104_ = nw101_
        index114_ = default__.safeIndex(468, (d_584_v104_).length(0))
        (d_584_v104_)[index114_] = (self).f25
        d_586_v105_: _dafny.Set
        d_586_v105_ = _dafny.Set({(self).f26, (self).f26, default__.fm16(p1, globalState)})
        index115_ = default__.safeIndex(468, (d_584_v104_).length(0))
        (d_584_v104_)[index115_] = (default__.safeDivisionInt(-284, p1)) != (len((d_586_v105_) - (d_586_v105_)))
        d_587_i10_: int
        d_587_i10_ = 0
        with _dafny.label("2"):
            while (self).f25:
                with _dafny.c_label("2"):
                    if (d_587_i10_) >= (100):
                        raise _dafny.Break("2")
                    d_587_i10_ = (d_587_i10_) + (1)
                    d_588_v106_: _dafny.Array
                    def lambda33_(d_589_p1_):
                        def lambda34_(d_590_i11_):
                            return default__.safeModuloInt(d_590_i11_, d_589_p1_)

                        return lambda34_

                    init21_ = lambda33_(p1)
                    nw102_ = _dafny.Array(None, 19)
                    for i0_21_ in range(nw102_.length(0)):
                        nw102_[i0_21_] = init21_(i0_21_)
                    d_588_v106_ = nw102_
                    d_591_v107_: _dafny.Set
                    d_591_v107_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({p0: len(d_586_v105_)})): (self).f25})) for d_592_i12_ in range(default__.abs(505))])})
                    d_593_v108_: D7
                    d_593_v108_ = D7_DC23(934, p1, (d_584_v104_)[default__.safeIndex(468, (d_584_v104_).length(0))], p1)
                    d_594_v109_: _dafny.Map
                    d_594_v109_ = _dafny.Map({d_574_v95_: d_593_v108_})
                    d_595_v111_: _dafny.Seq
                    d_595_v111_ = _dafny.SeqWithoutIsStrInference([d_584_v104_])
                    index116_ = default__.safeIndex(468, (d_584_v104_).length(0))
                    def iife40_():
                        coll22_ = _dafny.Set()
                        compr_22_: _dafny.Seq
                        for compr_22_ in (d_594_v109_).keys.Elements:
                            d_596_v110_: _dafny.Seq = compr_22_
                            if (d_596_v110_) in (d_594_v109_):
                                coll22_ = coll22_.union(_dafny.Set([d_596_v110_]))
                        return _dafny.Set(coll22_)
                    rhs100_ = (iife40_()
                    ).issubset(d_591_v107_)
                    rhs101_ = (d_584_v104_)[default__.safeIndex(468, (d_584_v104_).length(0))]
                    rhs102_ = (d_595_v111_)[default__.safeIndex(p1, len(d_595_v111_))]
                    rhs103_ = (self).f25
                    rhs104_ = d_588_v106_
                    lhs90_ = d_584_v104_
                    lhs91_ = default__.safeIndex(468, (d_584_v104_).length(0))
                    lhs90_[lhs91_] = rhs100_
                    r0 = rhs101_
                    d_584_v104_ = rhs102_
                    r0 = rhs103_
                    d_588_v106_ = rhs104_
                    d_597_v112_: _dafny.Map
                    d_597_v112_ = _dafny.Map({p1: (self).f25})
                    d_598_v113_: _dafny.Set
                    d_598_v113_ = _dafny.Set({default__.fm2((self).f25, -283, (d_584_v104_)[default__.safeIndex(468, (d_584_v104_).length(0))], (self).f26, globalState)})
                    d_599_v114_: _dafny.Array
                    nw103_ = _dafny.Array(None, 6)
                    nw103_[int(0)] = _dafny.Set({p1, default__.fm2((d_584_v104_)[default__.safeIndex(468, (d_584_v104_).length(0))], len(p2), (d_584_v104_)[default__.safeIndex(468, (d_584_v104_).length(0))], _dafny.CodePoint('m'), globalState), p1})
                    nw103_[int(1)] = default__.fm24((d_584_v104_)[default__.safeIndex(468, (d_584_v104_).length(0))], len(d_597_v112_), p1, globalState)
                    nw103_[int(2)] = d_598_v113_
                    nw103_[int(3)] = (d_598_v113_) | (_dafny.Set({p1, (0) - (p1)}))
                    nw103_[int(4)] = d_598_v113_
                    nw103_[int(5)] = (d_598_v113_) | (_dafny.Set({default__.fm2((self).f25, p1, (self).f25, (self).f26, globalState), len(p2)}))
                    d_599_v114_ = nw103_
                    d_599_v114_ = d_599_v114_
                    (globalState).f18 = (((0) - (p1)) - (p1) if True else (p1) - (-111))
                    index117_ = default__.safeIndex(468, (d_584_v104_).length(0))
                    (d_584_v104_)[index117_] = True
                    pass
            pass
        d_600_v115_: D3
        d_600_v115_ = D3_DC9(p1, (self).f26, p1)
        d_601_v116_: _dafny.MultiSet
        d_601_v116_ = _dafny.MultiSet([(p0).set(default__.safeIndex(len(d_574_v95_), len(p0)), (d_600_v115_).cf14), p0])
        r0 = ((d_601_v116_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")), default__.abs(764))).ispropersubset(d_601_v116_)
        r0 = (self).f25
        r1 = p1
        return r0, r1


class C3(T1, T0):
    def  __init__(self):
        self._f26: str = _dafny.CodePoint('D')
        self._f27: _dafny.Array = _dafny.Array(None, 0)
        self._f25: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f26, f27, f25):
        (self)._f26 = f26
        (self)._f27 = f27
        (self)._f25 = f25

    def fm3(self, p0, p1, p2, p3, globalState):
        return (self).f25

    def fm4(self, p0, globalState):
        return default__.safeModuloInt(((0) - (-836)) * (len(_dafny.SeqWithoutIsStrInference([(self).f25, (self).f25]))), 314)

    def m2(self, p0, p1, p2, globalState):
        d_602_v0_: _dafny.Array
        nw104_ = _dafny.Array(int(0), 19)
        d_602_v0_ = nw104_
        index118_ = default__.safeIndex(994, (d_602_v0_).length(0))
        (d_602_v0_)[index118_] = default__.safeModuloInt(p2, p2)
        index119_ = default__.safeIndex(994, (d_602_v0_).length(0))
        def iife41_():
            coll23_ = _dafny.Map()
            compr_23_: _dafny.Seq
            for compr_23_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(p0)]) for d_603_i0_ in range(default__.abs(94))])).Elements:
                d_604_v1_: _dafny.Seq = compr_23_
                if (d_604_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(p0)]) for d_603_i0_ in range(default__.abs(94))])):
                    coll23_[d_604_v1_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cela")))
            return _dafny.Map(coll23_)
        (d_602_v0_)[index119_] = default__.safeDivisionInt(p2, default__.safeDivisionInt(default__.fm2((self).f25, default__.fm2((self).f25, p2, False, _dafny.CodePoint('e'), globalState), p1, (self).f26, globalState), len(iife41_()
        )))
        d_605_v2_: _dafny.Seq
        d_605_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pgjv"))
        d_605_v2_ = d_605_v2_
        d_606_v3_: C0
        nw105_ = C0()
        nw105_.ctor__()
        d_606_v3_ = nw105_
        d_607_v4_: C0
        nw106_ = C0()
        nw106_.ctor__()
        d_607_v4_ = nw106_
        (globalState).f2 = ((self).f25) == (p1)
        if p1:
            (globalState).f18 = default__.safeDivisionInt(p2, p2)
            d_608_v5_: _dafny.Set
            d_608_v5_ = _dafny.Set({(self).f25})
            (globalState).f5 = len(d_608_v5_)
            (globalState).f22 = ((d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))]) + ((d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))])
            if p1:
                d_609_v6_: _dafny.Array
                nw107_ = _dafny.Array(None, 21)
                d_609_v6_ = nw107_
                index120_ = default__.safeIndex(37, (d_609_v6_).length(0))
                (d_609_v6_)[index120_] = d_606_v3_
                index121_ = default__.safeIndex(37, (d_609_v6_).length(0))
                (d_609_v6_)[index121_] = d_606_v3_
                d_610_v7_: _dafny.Array
                nw108_ = _dafny.Array(_dafny.Seq({}), 1)
                d_610_v7_ = nw108_
                d_611_v8_: _dafny.Seq
                d_611_v8_ = _dafny.SeqWithoutIsStrInference([d_605_v2_])
                index122_ = default__.safeIndex(761, (d_610_v7_).length(0))
                (d_610_v7_)[index122_] = d_611_v8_
                index123_ = default__.safeIndex(761, (d_610_v7_).length(0))
                (d_610_v7_)[index123_] = d_611_v8_
                (globalState).f2 = not(p1)
                d_605_v2_ = ((d_605_v2_) + (d_605_v2_)) + (d_605_v2_)
                d_605_v2_ = (d_605_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rr")))
            elif True:
                (globalState).f21 = (0) - ((default__.safeModuloInt((d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))], p2)) - (495))
                (globalState).f18 = p2
                d_612_v9_: _dafny.Map
                d_612_v9_ = _dafny.Map({(d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))]: p2})
                d_612_v9_ = (d_612_v9_).set(p2, (p2) - (319))
                d_613_v10_: _dafny.Array
                nw109_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_613_v10_ = nw109_
                d_614_v11_: _dafny.Array
                def lambda35_(d_615_v2_):
                    def lambda36_(d_616_i1_):
                        return d_615_v2_

                    return lambda36_

                init22_ = lambda35_(d_605_v2_)
                nw110_ = _dafny.Array(None, 6)
                for i0_22_ in range(nw110_.length(0)):
                    nw110_[i0_22_] = init22_(i0_22_)
                d_614_v11_ = nw110_
                index124_ = default__.safeIndex(566, (d_613_v10_).length(0))
                (d_613_v10_)[index124_] = d_614_v11_
                index125_ = default__.safeIndex(566, (d_613_v10_).length(0))
                (d_613_v10_)[index125_] = d_614_v11_
                d_617_v12_: _dafny.Array
                nw111_ = _dafny.Array(None, 4)
                d_617_v12_ = nw111_
                d_618_v13_: _dafny.Seq
                d_618_v13_ = _dafny.SeqWithoutIsStrInference([d_617_v12_])
                d_617_v12_ = (d_618_v13_)[default__.safeIndex((d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))], len(d_618_v13_))]
            d_619_v14_: int
            d_619_v14_ = p2
            d_620_v15_: _dafny.Seq
            d_620_v15_ = _dafny.SeqWithoutIsStrInference([p2, (d_619_v14_)])
            d_621_v16_: _dafny.Map
            d_621_v16_ = _dafny.Map({-107: p1})
            d_622_v17_: _dafny.Map
            d_622_v17_ = _dafny.Map({(len(d_620_v15_)) != ((d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))]): d_621_v16_})
            d_622_v17_ = (d_622_v17_).set(True, d_621_v16_)
        elif True:
            d_623_v18_: _dafny.Map
            d_623_v18_ = _dafny.Map({(self).f26: (self).f26})
            (globalState).f2 = default__.fm0((0) - (len(d_623_v18_)), (d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))], (self).f25, globalState)
            (globalState).f18 = (d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))]
            (globalState).f2 = (self).f25
            index126_ = default__.safeIndex(994, (d_602_v0_).length(0))
            (d_602_v0_)[index126_] = p2
            d_624_v19_: _dafny.Map
            d_624_v19_ = _dafny.Map({p2: (self).f25})
            d_625_v20_: _dafny.Map
            d_625_v20_ = _dafny.Map({p1: p1})
            d_626_v21_: _dafny.Set
            d_626_v21_ = _dafny.Set({p1, p1, ((d_625_v20_)[not(not((self).f25))] if (not(not((self).f25))) in (d_625_v20_) else p1), False, p1})
            d_627_v22_: D1
            d_627_v22_ = D1_DC2(p1)
            d_628_v23_: D1
            d_628_v23_ = D1_DC1((self).f25)
            d_629_v24_: _dafny.Seq
            d_629_v24_ = _dafny.SeqWithoutIsStrInference([(d_628_v23_).cf1])
            d_630_v25_: _dafny.Set
            d_630_v25_ = _dafny.Set({(d_629_v24_).set(default__.safeIndex(p2, len(d_629_v24_)), p1)})
            d_631_v26_: _dafny.Array
            nw112_ = _dafny.Array(None, 16)
            nw112_[int(0)] = p1
            nw112_[int(1)] = (self).f25
            nw112_[int(2)] = (self).f25
            nw112_[int(3)] = not((self).f25)
            nw112_[int(4)] = ((0) - ((d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))])) <= ((d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))])
            nw112_[int(5)] = False
            nw112_[int(6)] = (p2) == (p2)
            nw112_[int(7)] = ((d_624_v19_)[(d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))]] if ((d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))]) in (d_624_v19_) else (self).f25)
            nw112_[int(8)] = (self).fm3((self).f25, len(d_626_v21_), default__.fm2((self).f25, p2, (self).fm3(p1, (d_602_v0_)[default__.safeIndex(994, (d_602_v0_).length(0))], -711, _dafny.Map({(self).f26: p2}), globalState), (self).f26, globalState), _dafny.Map({(self).f26: 926}), globalState)
            nw112_[int(9)] = (self).f25
            nw112_[int(10)] = (self).f25
            nw112_[int(11)] = (d_627_v22_).cf2
            nw112_[int(12)] = (self).f25
            nw112_[int(13)] = p1
            nw112_[int(14)] = (d_630_v25_).ispropersubset(d_630_v25_)
            nw112_[int(15)] = p1
            d_631_v26_ = nw112_
            d_631_v26_ = d_631_v26_

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: int = int(0)
        d_632_v0_: _dafny.MultiSet
        d_632_v0_ = _dafny.MultiSet([p1])
        d_633_v1_: _dafny.MultiSet
        d_633_v1_ = _dafny.MultiSet([p1, 149, (d_632_v0_).cardinality, p1])
        d_634_i0_: int
        d_634_i0_ = 0
        with _dafny.label("3"):
            while ((d_632_v0_) | (d_632_v0_)).ispropersubset(d_633_v1_):
                with _dafny.c_label("3"):
                    if (d_634_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_634_i0_ = (d_634_i0_) + (1)
                    d_635_v2_: _dafny.Array
                    nw113_ = _dafny.Array(None, 14)
                    nw113_[int(0)] = p0
                    nw113_[int(1)] = (p0) + (_dafny.SeqWithoutIsStrInference([(self).f26 for d_636_i1_ in range(default__.abs(768))]))
                    nw113_[int(2)] = _dafny.SeqWithoutIsStrInference([(self).f26 for d_637_i2_ in range(default__.abs(-96))])
                    nw113_[int(3)] = p0
                    nw113_[int(4)] = (p0).set(default__.safeIndex(p1, len(p0)), (self).f26)
                    nw113_[int(5)] = p0
                    nw113_[int(6)] = (p0) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhnjxv")))
                    nw113_[int(7)] = _dafny.SeqWithoutIsStrInference([(self).f26 for d_638_i3_ in range(default__.abs(469))])
                    nw113_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nyxcrvbu"))
                    nw113_[int(9)] = (_dafny.SeqWithoutIsStrInference([(self).f26 for d_639_i4_ in range(default__.abs(992))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imrgbtikj")))
                    nw113_[int(10)] = p0
                    nw113_[int(11)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dewtel"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fcxwj")))
                    nw113_[int(12)] = (p0 if (self).f25 else _dafny.SeqWithoutIsStrInference([(self).f26 for d_640_i5_ in range(default__.abs(304))]))
                    nw113_[int(13)] = p0
                    d_635_v2_ = nw113_
                    index127_ = default__.safeIndex(908, (d_635_v2_).length(0))
                    (d_635_v2_)[index127_] = p0
                    index128_ = default__.safeIndex(908, (d_635_v2_).length(0))
                    (d_635_v2_)[index128_] = (p0).set(default__.safeIndex(p1, len(p0)), (self).f26)
                    d_641_v3_: _dafny.Map
                    d_641_v3_ = _dafny.Map({len((d_635_v2_)[default__.safeIndex(908, (d_635_v2_).length(0))]): p1})
                    d_642_v4_: _dafny.Seq
                    d_642_v4_ = _dafny.SeqWithoutIsStrInference([d_641_v3_])
                    d_643_v5_: D2
                    d_643_v5_ = D2_DC3(d_642_v4_)
                    d_644_v6_: _dafny.MultiSet
                    d_644_v6_ = _dafny.MultiSet([(self).f25])
                    d_645_v7_: _dafny.Map
                    d_645_v7_ = _dafny.Map({d_644_v6_: d_642_v4_})
                    (globalState).f21 = len(((d_643_v5_).cf3) + (((d_645_v7_)[d_644_v6_] if (d_644_v6_) in (d_645_v7_) else _dafny.SeqWithoutIsStrInference([d_641_v3_]))))
                    d_646_v8_: _dafny.Seq
                    d_646_v8_ = _dafny.SeqWithoutIsStrInference([p0, (((d_635_v2_)[default__.safeIndex(908, (d_635_v2_).length(0))]) + ((d_635_v2_)[default__.safeIndex(908, (d_635_v2_).length(0))])).set(default__.safeIndex(5, len(((d_635_v2_)[default__.safeIndex(908, (d_635_v2_).length(0))]) + ((d_635_v2_)[default__.safeIndex(908, (d_635_v2_).length(0))]))), (self).f26), (d_635_v2_)[default__.safeIndex(908, (d_635_v2_).length(0))], (default__.fm7((self).f26, globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))), (d_635_v2_)[default__.safeIndex(908, (d_635_v2_).length(0))]])
                    index129_ = default__.safeIndex(908, (d_635_v2_).length(0))
                    (d_635_v2_)[index129_] = (d_646_v8_)[default__.safeIndex(p1, len(d_646_v8_))]
                    d_647_v9_: _dafny.Map
                    d_647_v9_ = _dafny.Map({(self).f25: default__.fm2((self).f25, len(_dafny.SeqWithoutIsStrInference([(self).f26 for d_648_i6_ in range(default__.abs(887))])), (self).f25, _dafny.CodePoint('g'), globalState)})
                    d_649_v10_: _dafny.Array
                    def lambda37_(d_650_i7_):
                        return True

                    init23_ = lambda37_
                    nw114_ = _dafny.Array(None, 4)
                    for i0_23_ in range(nw114_.length(0)):
                        nw114_[i0_23_] = init23_(i0_23_)
                    d_649_v10_ = nw114_
                    index130_ = default__.safeIndex(653, (d_649_v10_).length(0))
                    (d_649_v10_)[index130_] = (self).f25
                    d_651_v11_: D1
                    d_651_v11_ = D1_DC2((self).f25)
                    index131_ = default__.safeIndex(338, (d_649_v10_).length(0))
                    (d_649_v10_)[index131_] = (d_651_v11_).cf2
                    index132_ = default__.safeIndex(653, (d_649_v10_).length(0))
                    index133_ = default__.safeIndex(338, (d_649_v10_).length(0))
                    rhs105_ = d_647_v9_
                    rhs106_ = (self).f25
                    rhs107_ = default__.fm0((p1) + (p1), (p1) - (p1), (self).f25, globalState)
                    rhs108_ = ((0) - ((0) - (p1))) - (p1)
                    rhs109_ = True
                    lhs92_ = globalState
                    lhs93_ = d_649_v10_
                    lhs94_ = default__.safeIndex(653, (d_649_v10_).length(0))
                    lhs95_ = d_649_v10_
                    lhs96_ = default__.safeIndex(338, (d_649_v10_).length(0))
                    d_647_v9_ = rhs105_
                    lhs92_.f2 = rhs106_
                    lhs93_[lhs94_] = rhs107_
                    r1 = rhs108_
                    lhs95_[lhs96_] = rhs109_
                    pass
            pass
        d_652_v12_: _dafny.Array
        def lambda38_(d_653_p1_):
            def lambda39_(d_654_i8_):
                return default__.safeDivisionInt(d_654_i8_, d_653_p1_)

            return lambda39_

        init24_ = lambda38_(p1)
        nw115_ = _dafny.Array(None, 3)
        for i0_24_ in range(nw115_.length(0)):
            nw115_[i0_24_] = init24_(i0_24_)
        d_652_v12_ = nw115_
        index134_ = default__.safeIndex(765, (d_652_v12_).length(0))
        (d_652_v12_)[index134_] = 249
        index135_ = default__.safeIndex(765, (d_652_v12_).length(0))
        (d_652_v12_)[index135_] = p1
        if (self).f25:
            r0 = (self).f25
            d_655_v13_: str
            d_655_v13_ = _dafny.CodePoint('w')
            d_656_v14_: _dafny.Set
            d_656_v14_ = _dafny.Set({p0, (p0) + (p0), p0})
            d_657_v15_: _dafny.Seq
            d_657_v15_ = _dafny.SeqWithoutIsStrInference([(d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]])
            d_658_v16_: int
            d_658_v16_ = len(_dafny.SeqWithoutIsStrInference([(self).f25]))
            d_659_v17_: _dafny.Map
            d_659_v17_ = _dafny.Map({(self).f25: d_658_v16_})
            rhs110_ = not(default__.fm0((0) - (default__.safeModuloInt((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))], (d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))])), (d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))], (self).f25, globalState))
            rhs111_ = (self).f26
            rhs112_ = ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]) == ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))])
            rhs113_ = default__.safeModuloInt(len(d_657_v15_), default__.safeModuloInt((d_657_v15_)[default__.safeIndex(len(d_659_v17_), len(d_657_v15_))], len((p0).set(default__.safeIndex(p1, len(p0)), (self).f26))))
            rhs114_ = d_656_v14_
            lhs97_ = globalState
            lhs98_ = globalState
            lhs97_.f2 = rhs110_
            d_655_v13_ = rhs111_
            r0 = rhs112_
            lhs98_.f21 = rhs113_
            d_656_v14_ = rhs114_
            index136_ = default__.safeIndex(765, (d_652_v12_).length(0))
            (d_652_v12_)[index136_] = (0) - (default__.safeDivisionInt((len(_dafny.Map({609: (self).f25}))), len(p0)))
            d_660_v18_: _dafny.Array
            nw116_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_660_v18_ = nw116_
            d_661_v19_: _dafny.Array
            nw117_ = _dafny.Array(None, 18)
            nw117_[int(0)] = d_660_v18_
            nw117_[int(1)] = d_660_v18_
            nw117_[int(2)] = d_660_v18_
            nw117_[int(3)] = d_660_v18_
            nw117_[int(4)] = d_660_v18_
            nw117_[int(5)] = d_660_v18_
            nw117_[int(6)] = d_660_v18_
            nw117_[int(7)] = d_660_v18_
            nw117_[int(8)] = d_660_v18_
            nw117_[int(9)] = d_660_v18_
            nw117_[int(10)] = d_660_v18_
            nw117_[int(11)] = d_660_v18_
            nw117_[int(12)] = d_660_v18_
            nw117_[int(13)] = d_660_v18_
            nw117_[int(14)] = d_660_v18_
            nw117_[int(15)] = d_660_v18_
            nw117_[int(16)] = d_660_v18_
            nw117_[int(17)] = (d_660_v18_ if (self).f25 else d_660_v18_)
            d_661_v19_ = nw117_
            index137_ = default__.safeIndex(488, (d_661_v19_).length(0))
            (d_661_v19_)[index137_] = d_660_v18_
            index138_ = default__.safeIndex(488, (d_661_v19_).length(0))
            (d_661_v19_)[index138_] = d_660_v18_
            (globalState).f5 = (p1) + (p1)
        elif True:
            d_662_v20_: _dafny.Set
            d_662_v20_ = _dafny.Set({586, p1})
            d_663_v21_: _dafny.Seq
            d_663_v21_ = _dafny.SeqWithoutIsStrInference([len(d_662_v20_)])
            if (((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]) != ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]) if not((len(d_663_v21_)) < ((0) - (len(p2)))) else not ((self).f25) or (False)):
                index139_ = default__.safeIndex(765, (d_652_v12_).length(0))
                (d_652_v12_)[index139_] = (d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]
                r0 = ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]) >= (496)
                (globalState).f18 = len(p0)
                (globalState).f2 = False
                d_664_v22_: str
                d_664_v22_ = _dafny.CodePoint('t')
                rhs115_ = d_664_v22_
                rhs116_ = (self).f25
                lhs99_ = globalState
                d_664_v22_ = rhs115_
                lhs99_.f2 = rhs116_
            elif True:
                index140_ = default__.safeIndex(765, (d_652_v12_).length(0))
                (d_652_v12_)[index140_] = p1
                (globalState).f2 = (self).f25
                d_665_v23_: T1
                nw118_ = C2()
                nw118_.ctor__((self).f26, (self).f27, default__.fm0((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))], (d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))], (self).f25, globalState))
                d_665_v23_ = nw118_
                d_666_v24_: _dafny.Array
                def lambda40_(d_667_i9_):
                    return (self).f25

                init25_ = lambda40_
                nw119_ = _dafny.Array(None, 3)
                for i0_25_ in range(nw119_.length(0)):
                    nw119_[i0_25_] = init25_(i0_25_)
                d_666_v24_ = nw119_
                (globalState).f5 = len((_dafny.Map({d_665_v23_: d_666_v24_})) | (_dafny.Map({d_665_v23_: d_666_v24_})))
                d_668_v25_: _dafny.MultiSet
                d_668_v25_ = _dafny.MultiSet([True, (self).f25])
                index141_ = default__.safeIndex(765, (d_652_v12_).length(0))
                (d_652_v12_)[index141_] = (((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]) + (((d_668_v25_)[(self).f25] if ((self).f25) in (d_668_v25_) else (d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))])) if (d_665_v23_).f25 else p1)
                d_669_v26_: _dafny.Array
                nw120_ = _dafny.Array(None, 25)
                d_669_v26_ = nw120_
                index142_ = default__.safeIndex(550, (d_669_v26_).length(0))
                (d_669_v26_)[index142_] = d_665_v23_
                index143_ = default__.safeIndex(550, (d_669_v26_).length(0))
                rhs117_ = (self).f25
                rhs118_ = d_665_v23_
                lhs100_ = globalState
                lhs101_ = d_669_v26_
                lhs102_ = default__.safeIndex(550, (d_669_v26_).length(0))
                lhs100_.f2 = rhs117_
                lhs101_[lhs102_] = rhs118_
            d_670_v27_: _dafny.Set
            d_670_v27_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: (self).f25})) for d_671_i10_ in range(default__.abs(435))]), _dafny.SeqWithoutIsStrInference([(d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]]), d_663_v21_, d_663_v21_})
            if (d_670_v27_).isdisjoint(d_670_v27_):
                (globalState).f5 = p1
                (globalState).f18 = ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]) + ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))])
                d_672_v28_: _dafny.Map
                d_672_v28_ = _dafny.Map({p1: (self).f25})
                d_673_v29_: D3
                d_673_v29_ = D3_DC7(d_672_v28_)
                d_674_v30_: _dafny.Array
                nw121_ = _dafny.Array(None, 7)
                nw121_[int(0)] = default__.fm25(globalState)
                nw121_[int(1)] = default__.fm25(globalState)
                nw121_[int(2)] = D3_DC7(d_672_v28_)
                nw121_[int(3)] = default__.fm25(globalState)
                nw121_[int(4)] = d_673_v29_
                nw121_[int(5)] = d_673_v29_
                nw121_[int(6)] = d_673_v29_
                d_674_v30_ = nw121_
                index144_ = default__.safeIndex(559, (d_674_v30_).length(0))
                (d_674_v30_)[index144_] = d_673_v29_
                pat_let_tv11_ = d_672_v28_
                index145_ = default__.safeIndex(559, (d_674_v30_).length(0))
                def iife42_(_pat_let9_0):
                    def iife43_(d_675_dt__update__tmp_h1_):
                        def iife44_(_pat_let10_0):
                            def iife45_(d_676_dt__update_hcf11_h0_):
                                return D3_DC7(d_676_dt__update_hcf11_h0_)
                            return iife45_(_pat_let10_0)
                        return iife44_(pat_let_tv11_)
                    return iife43_(_pat_let9_0)
                (d_674_v30_)[index145_] = iife42_(d_673_v29_)
                d_677_v31_: _dafny.Map
                d_677_v31_ = _dafny.Map({d_663_v21_: p0})
                d_677_v31_ = (d_677_v31_).set(d_663_v21_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjmlwg"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))))
                d_678_v32_: C2
                nw122_ = C2()
                nw122_.ctor__((self).f26, (self).f27, (self).f25)
                d_678_v32_ = nw122_
            elif True:
                (globalState).f5 = 807
                (globalState).f5 = p1
                d_679_v33_: _dafny.Set
                d_679_v33_ = _dafny.Set({(self).f25})
                (globalState).f18 = len(d_679_v33_)
                d_680_v34_: T1
                nw123_ = C2()
                nw123_.ctor__((self).f26, (self).f27, (self).f25)
                d_680_v34_ = nw123_
                rhs119_ = (self).f25
                rhs120_ = (d_680_v34_ if (self).f25 else d_680_v34_)
                lhs103_ = globalState
                lhs103_.f2 = rhs119_
                d_680_v34_ = rhs120_
                (globalState).f2 = (self).f25
            (globalState).f2 = (self).f25
            d_681_v35_: _dafny.Array
            def lambda41_(d_682_i11_):
                return False

            init26_ = lambda41_
            nw124_ = _dafny.Array(None, 29)
            for i0_26_ in range(nw124_.length(0)):
                nw124_[i0_26_] = init26_(i0_26_)
            d_681_v35_ = nw124_
            index146_ = default__.safeIndex(574, (d_681_v35_).length(0))
            (d_681_v35_)[index146_] = False
            index147_ = default__.safeIndex(574, (d_681_v35_).length(0))
            (d_681_v35_)[index147_] = ((self).f25) or (not((self).f25))
            d_683_v36_: D5
            d_683_v36_ = D5_DC15((self).f25, d_681_v35_, p0, False)
            source7_ = d_683_v36_
            if source7_.is_DC15:
                d_684___mcc_h0_ = source7_.cf25
                d_685___mcc_h1_ = source7_.cf26
                d_686___mcc_h2_ = source7_.cf27
                d_687___mcc_h3_ = source7_.cf28
                d_688_cf28_ = d_687___mcc_h3_
                d_689_cf27_ = d_686___mcc_h2_
                d_690_cf26_ = d_685___mcc_h1_
                d_691_cf25_ = d_684___mcc_h0_
                (globalState).f18 = p1
                (globalState).f22 = default__.fm2((d_681_v35_)[default__.safeIndex(574, (d_681_v35_).length(0))], (d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))], (self).f25, _dafny.CodePoint('f'), globalState)
                d_692_v37_: C2
                nw125_ = C2()
                nw125_.ctor__((self).f26, (self).f27, d_691_cf25_)
                d_692_v37_ = nw125_
                d_693_v38_: D3
                d_693_v38_ = D3_DC9(p1, _dafny.CodePoint('r'), p1)
                d_694_v39_: D3
                d_694_v39_ = D3_DC10(d_693_v38_)
                d_695_v40_: _dafny.Map
                d_695_v40_ = _dafny.Map({p1: (self).f25})
                pat_let_tv12_ = d_695_v40_
                def iife46_(_pat_let11_0):
                    def iife47_(d_696_dt__update__tmp_h2_):
                        def iife48_(_pat_let12_0):
                            def iife49_(d_697_dt__update_hcf16_h0_):
                                return D3_DC10(d_697_dt__update_hcf16_h0_)
                            return iife49_(_pat_let12_0)
                        return iife48_(D3_DC7(pat_let_tv12_))
                    return iife47_(_pat_let11_0)
                d_694_v39_ = iife46_(d_694_v39_)
            elif source7_.is_DC14:
                d_698___mcc_h4_ = source7_.cf24
                d_699_cf24_ = d_698___mcc_h4_
                d_700_v41_: D2
                d_700_v41_ = D2_DC5(204, True, (d_681_v35_)[default__.safeIndex(574, (d_681_v35_).length(0))], (self).f25)
                d_701_v42_: _dafny.Map
                d_701_v42_ = _dafny.Map({(d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]: d_681_v35_})
                d_702_v43_: _dafny.Map
                d_702_v43_ = _dafny.Map({(d_700_v41_).cf6: d_701_v42_})
                index148_ = default__.safeIndex(765, (d_652_v12_).length(0))
                (d_652_v12_)[index148_] = ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))] if (d_681_v35_)[default__.safeIndex(574, (d_681_v35_).length(0))] else len(((d_702_v43_)[p1] if (p1) in (d_702_v43_) else _dafny.Map({171: d_681_v35_}))))
                (globalState).f18 = (0) - (default__.safeModuloInt((335) + ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]), ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]) - ((0) - ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]))))
                d_703_v44_: C2
                nw126_ = C2()
                nw126_.ctor__((self).f26, (self).f27, (self).f25)
                d_703_v44_ = nw126_
                d_704_v45_: _dafny.Map
                d_704_v45_ = _dafny.Map({d_703_v44_: (p0) + (_dafny.SeqWithoutIsStrInference([(self).f26 for d_705_i12_ in range(default__.abs(44))]))})
                d_704_v45_ = d_704_v45_
                d_706_v46_: _dafny.Seq
                d_706_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "if"))
                d_706_v46_ = ((p0) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdalhwjvr")))) + (d_706_v46_)
            elif True:
                d_707___mcc_h5_ = source7_.cf29
                d_708_cf29_ = d_707___mcc_h5_
                (globalState).f2 = ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]) != ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))])
                (globalState).f5 = (d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]
                d_709_v47_: _dafny.Map
                d_709_v47_ = _dafny.Map({(d_681_v35_)[default__.safeIndex(574, (d_681_v35_).length(0))]: p0})
                (globalState).f5 = (0) - ((d_663_v21_)[default__.safeIndex((0) - (len(d_709_v47_)), len(d_663_v21_))])
                index149_ = default__.safeIndex(574, (d_681_v35_).length(0))
                (d_681_v35_)[index149_] = (p1) not in ((d_632_v0_).set(p1, default__.abs((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))])))
        r1 = (((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]) * ((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))])) + (default__.safeDivisionInt(p1, len(p0)))
        r1 = (d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]
        r0 = (self).f25
        d_710_v49_: D8
        d_710_v49_ = D8_DC25(d_633_v1_)
        d_711_v50_: D6
        def iife50_():
            coll24_ = _dafny.Map()
            compr_24_: int
            for compr_24_ in (((d_710_v49_).cf48).set(p1, default__.abs((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]))).Elements:
                d_712_v48_: int = compr_24_
                if (d_712_v48_) in (((d_710_v49_).cf48).set(p1, default__.abs((d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]))):
                    coll24_[default__.safeModuloInt(d_712_v48_, len(_dafny.SeqWithoutIsStrInference([(self).f26 for d_713_i13_ in range(default__.abs(66))])))] = p1
            return _dafny.Map(coll24_)
        d_711_v50_ = D6_DC17(iife50_()
)
        def lambda42_(source8_):
            if source8_.is_DC18:
                d_714___mcc_h6_ = source8_.cf31
                d_715___mcc_h7_ = source8_.cf32
                d_716___mcc_h8_ = source8_.cf33
                d_717___mcc_h9_ = source8_.cf34
                d_718___mcc_h10_ = source8_.cf35
                d_719_cf35_ = d_718___mcc_h10_
                d_720_cf34_ = d_717___mcc_h9_
                d_721_cf33_ = d_716___mcc_h8_
                d_722_cf32_ = d_715___mcc_h7_
                d_723_cf31_ = d_714___mcc_h6_
                return (d_721_cf33_) or (d_720_cf34_)
            elif source8_.is_DC17:
                d_724___mcc_h11_ = source8_.cf30
                d_725_cf30_ = d_724___mcc_h11_
                return True
            elif True:
                d_726___mcc_h12_ = source8_.cf36
                d_727_cf36_ = d_726___mcc_h12_
                return (self).f25

        r0 = lambda42_(d_711_v50_)
        r1 = (d_652_v12_)[default__.safeIndex(765, (d_652_v12_).length(0))]
        return r0, r1

