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
    def fm6(p0, globalState):
        if False:
            return D0_DC0(True)
        elif True:
            return D0_DC0(True)

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        if False:
            return D3_DC8(_dafny.SeqWithoutIsStrInference([D1_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_0_i0_ in range(default__.abs(615))]))]))
        elif True:
            return D3_DC8(_dafny.SeqWithoutIsStrInference([D1_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1_i2_ in range(default__.abs(-926))])) for d_2_i1_ in range(default__.abs(635))]))

    @staticmethod
    def fm9(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qeexk"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_3_i0_ in range(default__.abs(105))]))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tcrjsmo"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_4_i1_ in range(default__.abs(995))])))

    @staticmethod
    def fm10(p0, globalState):
        return (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acxmpij")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xt"))])), (0) - (len(_dafny.Set({True, not(False), True, False}))), len(_dafny.SeqWithoutIsStrInference([500]))])})) | ((_dafny.Map({True: _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([410])).cardinality])}) if False else _dafny.Map({False: _dafny.SeqWithoutIsStrInference([115])})))

    @staticmethod
    def fm11(p0, globalState):
        source0_ = D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbtb")))
        if source0_.is_DC4:
            d_5___mcc_h0_ = source0_.cf5
            d_6___mcc_h1_ = source0_.cf6
            d_7_cf6_ = d_6___mcc_h1_
            d_8_cf5_ = d_5___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([d_8_cf5_])) + (_dafny.SeqWithoutIsStrInference([d_8_cf5_, d_7_cf6_, d_8_cf5_, d_8_cf5_]))
        elif source0_.is_DC5:
            d_9___mcc_h2_ = source0_.cf7
            d_10___mcc_h3_ = source0_.cf8
            d_11___mcc_h4_ = source0_.cf9
            d_12___mcc_h5_ = source0_.cf10
            d_13_cf10_ = d_12___mcc_h5_
            d_14_cf9_ = d_11___mcc_h4_
            d_15_cf8_ = d_10___mcc_h3_
            d_16_cf7_ = d_9___mcc_h2_
            if not(d_14_cf9_):
                return _dafny.SeqWithoutIsStrInference([(0) - (d_16_cf7_) for d_17_i0_ in range(default__.abs(145))])
            elif True:
                return _dafny.SeqWithoutIsStrInference([d_15_cf8_])
        elif True:
            d_18___mcc_h6_ = source0_.cf4
            d_19_cf4_ = d_18___mcc_h6_
            return _dafny.SeqWithoutIsStrInference([767])

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([688, 805]))), 52])).Elements:
                d_20_v0_: int = compr_0_
                if (d_20_v0_) in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([688, 805]))), 52])):
                    coll0_[(d_20_v0_) * (82)] = len(_dafny.SeqWithoutIsStrInference([True]))
            return _dafny.Map(coll0_)
        return ((_dafny.Set({492}) if False else _dafny.Set({135, len(iife0_()
        )}))).intersection((_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lql")))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_21_i0_ in range(default__.abs(210))])), 765, 79})).intersection((D10_DC28(_dafny.Set({701}))).cf44))

    @staticmethod
    def fm13(p0, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(583, 764):
                d_22_v0_: int = compr_1_
                if ((583) <= (d_22_v0_)) and ((d_22_v0_) < (764)):
                    coll1_[(d_22_v0_) * (len(_dafny.Set({37})))] = 350
            return _dafny.Map(coll1_)
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(-272, 931):
                d_23_v1_: int = compr_2_
                if ((-272) <= (d_23_v1_)) and ((d_23_v1_) < (931)):
                    coll2_[(d_23_v1_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_24_i0_ in range(default__.abs(506))])))] = 569
            return _dafny.Map(coll2_)
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(900, 332):
                d_25_v2_: int = compr_3_
                if ((900) <= (d_25_v2_)) and ((d_25_v2_) < (332)):
                    coll3_[default__.safeDivisionInt(d_25_v2_, 574)] = 492
            return _dafny.Map(coll3_)
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in (_dafny.Map({410: False})).keys.Elements:
                d_26_v3_: int = compr_4_
                if (d_26_v3_) in (_dafny.Map({410: False})):
                    coll4_[default__.safeDivisionInt(d_26_v3_, 610)] = 745
            return _dafny.Map(coll4_)
        return ((iife1_()
        ) | (iife2_()
        )) | ((iife3_()
         if True else iife4_()
        ))

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: _dafny.Set
            for compr_5_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({not(not(True))}), _dafny.Set({True, not(False)})])).Elements:
                d_27_v0_: _dafny.Set = compr_5_
                if (d_27_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({not(not(True))}), _dafny.Set({True, not(False)})])):
                    coll5_ = coll5_.union(_dafny.Set([d_27_v0_]))
            return _dafny.Set(coll5_)
        return (_dafny.Set({_dafny.Set({False})})).intersection(iife5_()
        )

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return _dafny.Map({(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_28_i0_ in range(default__.abs(804))]))) == (-960): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvbhie")))})

    @staticmethod
    def fm16(p0, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in (_dafny.Map({724: len(_dafny.SeqWithoutIsStrInference([-52]))})).keys.Elements:
                d_29_v0_: int = compr_6_
                if (d_29_v0_) in (_dafny.Map({724: len(_dafny.SeqWithoutIsStrInference([-52]))})):
                    coll6_[(d_29_v0_) * (628)] = False
            return _dafny.Map(coll6_)
        return ((iife6_()
        ) | (_dafny.Map({-925: False}))) | ((_dafny.Map({47: True})) | (_dafny.Map({623: True})))

    @staticmethod
    def fm17(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.Set({33})) | (_dafny.Set({-138}))])

    @staticmethod
    def fm18(globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(871, 44):
                d_30_v0_: int = compr_7_
                if ((871) <= (d_30_v0_)) and ((d_30_v0_) < (44)):
                    coll7_[(d_30_v0_) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False]))).cardinality)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])
            return _dafny.Map(coll7_)
        return ((_dafny.SeqWithoutIsStrInference([not(not(True))])) + (_dafny.SeqWithoutIsStrInference([False]))) + ((_dafny.SeqWithoutIsStrInference([not(False), False, (D1_DC5((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmamusrti")))), 983, False, iife7_()
)).cf9, (D0_DC0(not(True))).cf0])) + (_dafny.SeqWithoutIsStrInference([False, False, False])))

    @staticmethod
    def fm19(globalState):
        return (0) - ((len((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False]) for d_31_i0_ in range(default__.abs(-775))])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True])]))) if (_dafny.SeqWithoutIsStrInference([False])) != (_dafny.SeqWithoutIsStrInference([True, False])) else default__.safeDivisionInt((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltc")))), 236)))

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        return _dafny.CodePoint('d')

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        source1_ = D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yohqnt")))
        if source1_.is_DC4:
            d_32___mcc_h0_ = source1_.cf5
            d_33___mcc_h1_ = source1_.cf6
            d_34_cf6_ = d_33___mcc_h1_
            d_35_cf5_ = d_32___mcc_h0_
            return False
        elif source1_.is_DC5:
            d_36___mcc_h2_ = source1_.cf7
            d_37___mcc_h3_ = source1_.cf8
            d_38___mcc_h4_ = source1_.cf9
            d_39___mcc_h5_ = source1_.cf10
            d_40_cf10_ = d_39___mcc_h5_
            d_41_cf9_ = d_38___mcc_h4_
            d_42_cf8_ = d_37___mcc_h3_
            d_43_cf7_ = d_36___mcc_h2_
            return ((_dafny.MultiSet([d_42_cf8_, d_42_cf8_, d_42_cf8_, d_43_cf7_])).cardinality) == ((_dafny.MultiSet([237, len(_dafny.Map({d_43_cf7_: (0) - (d_42_cf8_)}))])).cardinality)
        elif True:
            d_44___mcc_h6_ = source1_.cf4
            d_45_cf4_ = d_44___mcc_h6_
            return (len(_dafny.SeqWithoutIsStrInference([len(d_45_cf4_), 390, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality, len(d_45_cf4_), len(_dafny.SeqWithoutIsStrInference([True]))]))) <= (-744)

    @staticmethod
    def fm22(globalState):
        return _dafny.MultiSet([not (True) or (True)])

    @staticmethod
    def fm23(p0, p1, globalState):
        source2_ = D5_DC13(_dafny.Map({True: 738}))
        if source2_.is_DC14:
            d_46___mcc_h0_ = source2_.cf25
            d_47___mcc_h1_ = source2_.cf26
            d_48___mcc_h2_ = source2_.cf27
            d_49_cf27_ = d_48___mcc_h2_
            d_50_cf26_ = d_47___mcc_h1_
            d_51_cf25_ = d_46___mcc_h0_
            return D3_DC9(d_51_cf25_)
        elif source2_.is_DC13:
            d_52___mcc_h3_ = source2_.cf24
            d_53_cf24_ = d_52___mcc_h3_
            return D3_DC9((0) - ((_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))))])).cardinality))
        elif True:
            d_54___mcc_h4_ = source2_.cf28
            d_55_cf28_ = d_54___mcc_h4_
            return D3_DC9(218)

    @staticmethod
    def m0(p0, globalState):
        d_56_v0_: _dafny.Seq
        d_56_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ir"))
        d_57_v1_: int
        d_57_v1_ = 671
        d_56_v0_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drtgu"))) + (default__.fm9((0) - (d_57_v1_), globalState))) + ((d_56_v0_) + (d_56_v0_))
        d_58_v2_: bool
        d_58_v2_ = True
        d_59_v3_: _dafny.MultiSet
        d_59_v3_ = _dafny.MultiSet([not(d_58_v2_)])
        d_60_v4_: _dafny.Map
        d_60_v4_ = _dafny.Map({d_58_v2_: d_59_v3_})
        d_61_v5_: _dafny.Seq
        d_61_v5_ = _dafny.SeqWithoutIsStrInference([(((d_60_v4_)[d_58_v2_] if (d_58_v2_) in (d_60_v4_) else d_59_v3_)).set(False, default__.abs(d_57_v1_))])
        d_62_v6_: _dafny.MultiSet
        d_62_v6_ = _dafny.MultiSet([len(d_56_v0_)])
        d_63_v7_: _dafny.Map
        d_63_v7_ = _dafny.Map({d_62_v6_: False})
        d_64_v8_: _dafny.Map
        d_64_v8_ = _dafny.Map({d_57_v1_: len(d_63_v7_)})
        d_65_v9_: _dafny.Seq
        d_65_v9_ = _dafny.SeqWithoutIsStrInference([d_58_v2_])
        d_66_v10_: D2
        d_66_v10_ = D2_DC7(d_58_v2_, False, d_57_v1_)
        d_67_v11_: _dafny.Array
        nw0_ = _dafny.Array(None, 20)
        nw0_[int(0)] = default__.fm22(globalState)
        nw0_[int(1)] = d_59_v3_
        nw0_[int(2)] = (default__.fm22(globalState)).set(d_58_v2_, default__.abs(d_57_v1_))
        nw0_[int(3)] = (d_61_v5_)[default__.safeIndex(d_57_v1_, len(d_61_v5_))]
        nw0_[int(4)] = _dafny.MultiSet([d_58_v2_, d_58_v2_, not(default__.fm21(d_57_v1_, d_58_v2_, d_58_v2_, d_64_v8_, globalState))])
        nw0_[int(5)] = d_59_v3_
        nw0_[int(6)] = d_59_v3_
        nw0_[int(7)] = d_59_v3_
        nw0_[int(8)] = d_59_v3_
        nw0_[int(9)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_58_v2_, d_58_v2_]))
        nw0_[int(10)] = d_59_v3_
        nw0_[int(11)] = d_59_v3_
        nw0_[int(12)] = (d_59_v3_) - (_dafny.MultiSet(((d_65_v9_).set(default__.safeIndex(d_57_v1_, len(d_65_v9_)), (d_66_v10_).cf12)).set(default__.safeIndex(d_57_v1_, len((d_65_v9_).set(default__.safeIndex(d_57_v1_, len(d_65_v9_)), (d_66_v10_).cf12))), d_58_v2_)))
        nw0_[int(13)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(d_58_v2_)]))
        nw0_[int(14)] = (d_59_v3_) - (default__.fm22(globalState))
        nw0_[int(15)] = (_dafny.MultiSet([d_58_v2_, d_58_v2_, d_58_v2_])) | ((d_59_v3_).set(d_58_v2_, default__.abs(len(d_56_v0_))))
        nw0_[int(16)] = (d_59_v3_).set(d_58_v2_, default__.abs(d_57_v1_))
        nw0_[int(17)] = d_59_v3_
        nw0_[int(18)] = d_59_v3_
        nw0_[int(19)] = _dafny.MultiSet(d_65_v9_)
        d_67_v11_ = nw0_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_67_v11_).length(0)):
            d_68_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_68_i0_)) and ((d_68_i0_) < ((d_67_v11_).length(0)))):
                (d_67_v11_)[(d_68_i0_)] = (d_59_v3_ if d_58_v2_ else (d_59_v3_).set(d_58_v2_, default__.abs(d_57_v1_)))
        d_69_v12_: T1
        nw1_ = C0()
        nw1_.ctor__(d_58_v2_, d_58_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtxlh")), d_57_v1_, d_58_v2_)
        d_69_v12_ = nw1_
        d_70_v13_: _dafny.Map
        d_70_v13_ = _dafny.Map({d_69_v12_: d_57_v1_})
        (globalState).f11 = ((((d_70_v13_)[d_69_v12_] if (d_69_v12_) in (d_70_v13_) else d_69_v12_.f14)) * (d_69_v12_.f14)) + (d_69_v12_.f14)
        d_58_v2_ = d_58_v2_
        source3_ = D0_DC2(d_69_v12_.f14, (_dafny.MultiSet([True, d_58_v2_, (d_69_v12_).f15, d_58_v2_])) - (_dafny.MultiSet(d_65_v9_)))
        if source3_.is_DC1:
            d_71___mcc_h0_ = source3_.cf1
            d_72_cf1_ = d_71___mcc_h0_
            d_73_v14_: _dafny.Map
            d_73_v14_ = _dafny.Map({True: (d_69_v12_).f15})
            d_74_v15_: C0
            nw2_ = C0()
            nw2_.ctor__(d_58_v2_, (((d_73_v14_)[d_58_v2_] if (d_58_v2_) in (d_73_v14_) else (d_69_v12_).f15)) == ((d_69_v12_).f15), d_69_v12_.f16, d_57_v1_, (d_69_v12_.f14) >= (d_69_v12_.f14))
            d_74_v15_ = nw2_
            d_74_v15_ = d_74_v15_
            d_57_v1_ = (0) - (((0) - (d_72_cf1_) if (_dafny.MultiSet([d_57_v1_])).issubset(d_62_v6_) else d_69_v12_.f14))
            d_75_v16_: _dafny.Map
            d_75_v16_ = _dafny.Map({(d_69_v12_).f15: d_69_v12_})
            d_76_v17_: _dafny.Array
            nw3_ = _dafny.Array(None, 1)
            nw3_[int(0)] = d_74_v15_
            d_76_v17_ = nw3_
            d_77_v18_: _dafny.MultiSet
            d_77_v18_ = _dafny.MultiSet([d_76_v17_, d_76_v17_, d_76_v17_, d_76_v17_])
            (globalState).f11 = (default__.safeModuloInt(d_69_v12_.f14, len(d_75_v16_)) if d_58_v2_ else (0) - (((d_77_v18_).set(d_76_v17_, default__.abs(d_69_v12_.f14))).cardinality))
            d_72_cf1_ = ((d_59_v3_)[(d_74_v15_).f17] if ((d_74_v15_).f17) in (d_59_v3_) else d_72_cf1_)
        elif source3_.is_DC2:
            d_78___mcc_h1_ = source3_.cf2
            d_79___mcc_h2_ = source3_.cf3
            d_80_cf3_ = d_79___mcc_h2_
            d_81_cf2_ = d_78___mcc_h1_
            d_82_v19_: _dafny.Array
            nw4_ = _dafny.Array(False, 6)
            d_82_v19_ = nw4_
            d_83_v20_: _dafny.Map
            d_83_v20_ = _dafny.Map({(d_69_v12_).f15: d_82_v19_})
            d_84_v21_: _dafny.Map
            d_84_v21_ = _dafny.Map({(d_69_v12_).fm0((d_69_v12_).f15, globalState): d_58_v2_})
            if default__.fm21(len(d_83_v20_), ((d_84_v21_)[(d_69_v12_).f15] if ((d_69_v12_).f15) in (d_84_v21_) else (d_69_v12_).f15), (d_69_v12_).f15, d_64_v8_, globalState):
                (globalState).f11 = (d_69_v12_.f14) - (d_57_v1_)
                d_85_v22_: D3
                d_85_v22_ = D3_DC9(d_81_cf2_)
                d_86_v23_: C1
                nw5_ = C1()
                nw5_.ctor__(d_69_v12_.f16, d_65_v9_, (d_59_v3_).cardinality, (d_69_v12_).f15)
                d_86_v23_ = nw5_
                d_87_v24_: _dafny.Map
                d_87_v24_ = _dafny.Map({(d_69_v12_).f15: d_86_v23_})
                d_85_v22_ = default__.fm23(_dafny.SeqWithoutIsStrInference([d_69_v12_.f14, len(d_87_v24_), d_69_v12_.f14]), not(d_58_v2_), globalState)
                d_88_v25_: _dafny.Array
                def lambda0_(d_89_v12_, d_90_cf3_, d_91_cf2_):
                    def lambda1_(d_92_i1_):
                        return _dafny.SeqWithoutIsStrInference([((d_90_cf3_)[(d_89_v12_).f15] if ((d_89_v12_).f15) in (d_90_cf3_) else d_89_v12_.f14), len(_dafny.Map({834: d_91_cf2_})), d_89_v12_.f14])

                    return lambda1_

                init0_ = lambda0_(d_69_v12_, d_80_cf3_, d_81_cf2_)
                nw6_ = _dafny.Array(None, 22)
                for i0_0_ in range(nw6_.length(0)):
                    nw6_[i0_0_] = init0_(i0_0_)
                d_88_v25_ = nw6_
                rhs0_ = d_69_v12_.f14
                rhs1_ = ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yitpb")))) - (d_57_v1_)) not in (d_62_v6_)
                rhs2_ = not((d_69_v12_.f14) < ((949) + (-463)))
                rhs3_ = d_56_v0_
                rhs4_ = (d_88_v25_ if d_58_v2_ else d_88_v25_)
                lhs0_ = globalState
                lhs1_ = globalState
                lhs0_.f11 = rhs0_
                lhs1_.f0 = rhs1_
                d_58_v2_ = rhs2_
                d_56_v0_ = rhs3_
                d_88_v25_ = rhs4_
                d_93_v26_: _dafny.Map
                d_93_v26_ = _dafny.Map({562: d_58_v2_})
                d_94_v27_: D8
                d_94_v27_ = D8_DC21(d_93_v26_)
                d_93_v26_ = ((d_94_v27_).cf36) | (d_93_v26_)
                d_95_v28_: _dafny.Array
                nw7_ = _dafny.Array(None, 28)
                d_95_v28_ = nw7_
                d_95_v28_ = d_95_v28_
            elif True:
                (d_69_v12_).f16 = (d_69_v12_.f16) + (d_69_v12_.f16)
                index0_ = default__.safeIndex(991, (d_82_v19_).length(0))
                (d_82_v19_)[index0_] = default__.fm21(d_69_v12_.f14, False, (d_69_v12_).f15, d_64_v8_, globalState)
                index1_ = default__.safeIndex(991, (d_82_v19_).length(0))
                (d_82_v19_)[index1_] = (d_80_cf3_).isdisjoint(d_59_v3_)
                d_96_v29_: _dafny.Array
                def lambda2_(d_97_v9_):
                    def lambda3_(d_98_i2_):
                        return d_97_v9_

                    return lambda3_

                init1_ = lambda2_(d_65_v9_)
                nw8_ = _dafny.Array(None, 24)
                for i0_1_ in range(nw8_.length(0)):
                    nw8_[i0_1_] = init1_(i0_1_)
                d_96_v29_ = nw8_
                d_99_v30_: D5
                d_99_v30_ = D5_DC14(d_69_v12_.f14, d_96_v29_, True)
                d_100_v31_: _dafny.Array
                nw9_ = _dafny.Array(None, 17)
                nw9_[int(0)] = (d_69_v12_).f15
                nw9_[int(1)] = False
                nw9_[int(2)] = (d_69_v12_).f15
                nw9_[int(3)] = (d_69_v12_).f15
                nw9_[int(4)] = (d_69_v12_).f15
                nw9_[int(5)] = (d_69_v12_).f15
                nw9_[int(6)] = (d_69_v12_).f15
                nw9_[int(7)] = (d_69_v12_).f15
                nw9_[int(8)] = (d_69_v12_).f15
                nw9_[int(9)] = (d_69_v12_).f15
                nw9_[int(10)] = (d_69_v12_).f15
                nw9_[int(11)] = not((d_99_v30_).cf27)
                nw9_[int(12)] = not(not((d_69_v12_).f15))
                nw9_[int(13)] = (d_69_v12_).f15
                nw9_[int(14)] = (d_69_v12_).f15
                nw9_[int(15)] = not(not(d_58_v2_))
                nw9_[int(16)] = (d_69_v12_).f15
                d_100_v31_ = nw9_
                d_101_v32_: _dafny.Map
                d_101_v32_ = _dafny.Map({d_100_v31_: 420})
                d_101_v32_ = (d_101_v32_).set(d_100_v31_, (d_62_v6_).cardinality)
                d_102_v33_: _dafny.Set
                d_102_v33_ = _dafny.Set({(d_69_v12_).f15, d_58_v2_})
                d_103_v34_: str
                d_103_v34_ = _dafny.CodePoint('y')
                d_104_v35_: _dafny.MultiSet
                d_104_v35_ = _dafny.MultiSet([d_103_v34_, d_103_v34_])
                pat_let_tv0_ = d_69_v12_
                pat_let_tv1_ = d_82_v19_
                pat_let_tv2_ = d_82_v19_
                pat_let_tv3_ = d_82_v19_
                pat_let_tv4_ = d_82_v19_
                pat_let_tv5_ = globalState
                pat_let_tv6_ = globalState
                pat_let_tv7_ = d_69_v12_
                d_105_v36_: C2
                nw10_ = C2()
                def iife8_(_pat_let0_0):
                    def iife9_(d_106_dt__update__tmp_h0_):
                        def iife10_(_pat_let1_0):
                            def iife11_(d_107_dt__update_hcf12_h0_):
                                def iife12_(_pat_let2_0):
                                    def iife13_(d_108_dt__update_hcf13_h0_):
                                        return D2_DC7(d_107_dt__update_hcf12_h0_, d_108_dt__update_hcf13_h0_, (d_106_dt__update__tmp_h0_).cf14)
                                    return iife13_(_pat_let2_0)
                                return iife12_((pat_let_tv7_).f15)
                            return iife11_(_pat_let1_0)
                        return iife10_(default__.fm21(pat_let_tv0_.f14, (pat_let_tv2_)[default__.safeIndex(991, (pat_let_tv1_).length(0))], (pat_let_tv4_)[default__.safeIndex(991, (pat_let_tv3_).length(0))], default__.fm13(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffbfp"))), pat_let_tv5_), pat_let_tv6_))
                    return iife9_(_pat_let0_0)
                nw10_.ctor__(iife8_(D2_DC7(default__.fm21(len(d_102_v33_), (d_69_v12_).f15, (d_69_v12_).f15, d_64_v8_, globalState), (d_69_v12_).f15, d_69_v12_.f14)), ((d_104_v35_)[d_103_v34_] if (d_103_v34_) in (d_104_v35_) else d_81_cf2_), (d_69_v12_).f15)
                d_105_v36_ = nw10_
                d_109_v37_: C0
                nw11_ = C0()
                nw11_.ctor__(not((d_69_v12_).f15), (d_69_v12_).f15, d_56_v0_, d_81_cf2_, (d_69_v12_).f15)
                d_109_v37_ = nw11_
                d_110_v38_: _dafny.Map
                d_110_v38_ = _dafny.Map({(d_65_v9_)[default__.safeIndex(d_81_cf2_, len(d_65_v9_))]: d_109_v37_})
                (globalState).f11 = len(((d_110_v38_) | (d_110_v38_)) | (d_110_v38_))
            d_111_v39_: C1
            nw12_ = C1()
            nw12_.ctor__(d_56_v0_, _dafny.SeqWithoutIsStrInference([d_58_v2_]), d_57_v1_, (d_69_v12_).f15)
            d_111_v39_ = nw12_
            d_112_v40_: _dafny.Map
            d_112_v40_ = _dafny.Map({d_111_v39_: (d_69_v12_).f15})
            d_113_v41_: _dafny.Map
            d_113_v41_ = _dafny.Map({(d_69_v12_).f15: len(d_112_v40_)})
            d_113_v41_ = (d_113_v41_).set((d_69_v12_.f14) <= (d_69_v12_.f14), d_69_v12_.f14)
            d_114_v42_: _dafny.Set
            d_114_v42_ = _dafny.Set({d_111_v39_, d_111_v39_, d_111_v39_})
            d_115_v43_: D1
            d_115_v43_ = D1_DC4(default__.safeDivisionInt(len(d_114_v42_), len(d_69_v12_.f16)), d_69_v12_.f14)
            d_115_v43_ = d_115_v43_
            d_116_v44_: _dafny.Array
            nw13_ = _dafny.Array(_dafny.CodePoint('D'), 24)
            d_116_v44_ = nw13_
            d_117_v45_: str
            d_117_v45_ = _dafny.CodePoint('s')
            index2_ = default__.safeIndex(783, (d_116_v44_).length(0))
            (d_116_v44_)[index2_] = d_117_v45_
            index3_ = default__.safeIndex(783, (d_116_v44_).length(0))
            (d_116_v44_)[index3_] = d_117_v45_
        elif True:
            d_118___mcc_h3_ = source3_.cf0
            d_119_cf0_ = d_118___mcc_h3_
            d_65_v9_ = (_dafny.SeqWithoutIsStrInference([False, (d_69_v12_).f15, (d_69_v12_).fm0((d_69_v12_).f15, globalState)])) + (d_65_v9_)
            d_120_v46_: C0
            nw14_ = C0()
            nw14_.ctor__(d_58_v2_, not((d_69_v12_).f15), d_56_v0_, d_69_v12_.f14, (d_69_v12_).f15)
            d_120_v46_ = nw14_
            d_121_v47_: _dafny.Map
            d_121_v47_ = _dafny.Map({d_62_v6_: d_120_v46_})
            d_122_v48_: _dafny.Map
            d_122_v48_ = _dafny.Map({d_58_v2_: d_69_v12_.f14})
            d_123_v49_: _dafny.Map
            d_123_v49_ = _dafny.Map({d_57_v1_: d_122_v48_})
            d_124_v50_: _dafny.Array
            nw15_ = _dafny.Array(None, 12)
            nw15_[int(0)] = d_120_v46_
            nw15_[int(1)] = d_120_v46_
            nw15_[int(2)] = d_120_v46_
            nw15_[int(3)] = d_120_v46_
            nw15_[int(4)] = d_120_v46_
            nw15_[int(5)] = d_120_v46_
            nw15_[int(6)] = d_120_v46_
            nw15_[int(7)] = d_120_v46_
            nw15_[int(8)] = (d_120_v46_ if (d_69_v12_).f15 else d_120_v46_)
            nw15_[int(9)] = d_120_v46_
            nw15_[int(10)] = d_120_v46_
            nw15_[int(11)] = ((d_121_v47_)[_dafny.MultiSet([len(((d_123_v49_)[512] if (512) in (d_123_v49_) else d_122_v48_))])] if (_dafny.MultiSet([len(((d_123_v49_)[512] if (512) in (d_123_v49_) else d_122_v48_))])) in (d_121_v47_) else d_120_v46_)
            d_124_v50_ = nw15_
            d_125_v51_: D9
            d_125_v51_ = D9_DC25(d_124_v50_)
            d_124_v50_ = (d_125_v51_).cf41
            d_126_v52_: _dafny.Map
            d_126_v52_ = _dafny.Map({default__.fm19(globalState): (d_69_v12_.f14) == (d_69_v12_.f14)})
            d_127_v53_: _dafny.MultiSet
            d_127_v53_ = _dafny.MultiSet([d_120_v46_, d_120_v46_])
            d_128_v54_: _dafny.Seq
            d_128_v54_ = _dafny.SeqWithoutIsStrInference([d_57_v1_])
            d_119_cf0_ = ((d_126_v52_)[(d_57_v1_ if d_119_cf0_ else ((d_127_v53_)[d_120_v46_] if (d_120_v46_) in (d_127_v53_) else len(d_128_v54_)))] if ((d_57_v1_ if d_119_cf0_ else ((d_127_v53_)[d_120_v46_] if (d_120_v46_) in (d_127_v53_) else len(d_128_v54_)))) in (d_126_v52_) else d_119_cf0_)
            d_129_v55_: _dafny.Array
            def lambda4_(d_130_v46_):
                def lambda5_(d_131_i3_):
                    return (d_130_v46_).f18

                return lambda5_

            init2_ = lambda4_(d_120_v46_)
            nw16_ = _dafny.Array(None, 1)
            for i0_2_ in range(nw16_.length(0)):
                nw16_[i0_2_] = init2_(i0_2_)
            d_129_v55_ = nw16_
            d_129_v55_ = d_129_v55_
        d_132_v56_: _dafny.Array
        nw17_ = _dafny.Array(int(0), 3)
        d_132_v56_ = nw17_
        index4_ = default__.safeIndex(889, (d_132_v56_).length(0))
        (d_132_v56_)[index4_] = (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([not((d_69_v12_).f15), d_58_v2_])) + (_dafny.SeqWithoutIsStrInference([d_58_v2_, (d_65_v9_)[default__.safeIndex(d_57_v1_, len(d_65_v9_))], (d_69_v12_).f15, d_58_v2_])))).cardinality
        index5_ = default__.safeIndex(889, (d_132_v56_).length(0))
        (d_132_v56_)[index5_] = ((0) - ((d_69_v12_.f14) + (len(d_65_v9_)))) + ((0) - (d_57_v1_))

    @staticmethod
    def Main(noArgsParameter__):
        d_133_v0_: bool
        d_133_v0_ = False
        d_134_v1_: _dafny.Seq
        d_134_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogvv"))
        d_135_v2_: _dafny.Seq
        d_135_v2_ = _dafny.SeqWithoutIsStrInference([d_133_v0_])
        d_136_v3_: int
        d_136_v3_ = 329
        d_137_v4_: _dafny.Seq
        d_137_v4_ = _dafny.SeqWithoutIsStrInference([d_133_v0_, d_133_v0_, d_133_v0_, (d_135_v2_)[default__.safeIndex((0) - (d_136_v3_), len(d_135_v2_))], not(d_133_v0_)])
        d_138_globalState_: GlobalState
        nw18_ = GlobalState()
        nw18_.ctor__(False, True, 140, True, (d_134_v1_ if d_133_v0_ else d_134_v1_), False, 188, True, 38, 937, 439, 8, False, d_135_v2_)
        d_138_globalState_ = nw18_
        if d_133_v0_:
            d_139_v5_: _dafny.Set
            d_139_v5_ = _dafny.Set({d_133_v0_, d_133_v0_})
            d_140_v6_: _dafny.Seq
            d_140_v6_ = _dafny.SeqWithoutIsStrInference([d_139_v5_])
            (d_138_globalState_).f0 = (d_140_v6_) != (_dafny.SeqWithoutIsStrInference([d_139_v5_, d_139_v5_, d_139_v5_, d_139_v5_]))
            d_141_v7_: _dafny.Map
            d_141_v7_ = _dafny.Map({d_136_v3_: not(d_133_v0_)})
            d_142_v8_: _dafny.Array
            nw19_ = _dafny.Array(False, 20)
            d_142_v8_ = nw19_
            default__.m0(_dafny.Map({d_141_v7_: d_142_v8_}), d_138_globalState_)
            d_143_v9_: _dafny.Map
            d_143_v9_ = _dafny.Map({d_136_v3_: d_134_v1_})
            d_143_v9_ = (d_143_v9_).set(d_136_v3_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkrlymcrd"))) + (d_134_v1_))
            d_144_v10_: _dafny.Array
            nw20_ = _dafny.Array(int(0), 5)
            d_144_v10_ = nw20_
            index6_ = default__.safeIndex(759, (d_144_v10_).length(0))
            (d_144_v10_)[index6_] = d_136_v3_
            d_145_v11_: _dafny.MultiSet
            d_145_v11_ = _dafny.MultiSet([d_136_v3_, 378])
            d_146_v12_: _dafny.Map
            d_146_v12_ = _dafny.Map({False: ((d_145_v11_).intersection((d_145_v11_).set(len(_dafny.SeqWithoutIsStrInference([d_136_v3_])), default__.abs(d_136_v3_)))).cardinality})
            d_147_v13_: _dafny.Map
            d_147_v13_ = _dafny.Map({not(False): d_146_v12_})
            index7_ = default__.safeIndex(759, (d_144_v10_).length(0))
            rhs5_ = True
            rhs6_ = d_139_v5_
            rhs7_ = d_136_v3_
            rhs8_ = (_dafny.Map({False: d_136_v3_})) | ((((d_147_v13_)[not(d_133_v0_)] if (not(d_133_v0_)) in (d_147_v13_) else d_146_v12_)) | (_dafny.Map({d_133_v0_: len(d_134_v1_)})))
            lhs2_ = d_144_v10_
            lhs3_ = default__.safeIndex(759, (d_144_v10_).length(0))
            d_133_v0_ = rhs5_
            d_139_v5_ = rhs6_
            lhs2_[lhs3_] = rhs7_
            d_146_v12_ = rhs8_
            d_148_v14_: _dafny.Map
            d_148_v14_ = _dafny.Map({_dafny.Map({(d_144_v10_)[default__.safeIndex(759, (d_144_v10_).length(0))]: d_133_v0_}): d_142_v8_})
            default__.m0(d_148_v14_, d_138_globalState_)
        elif True:
            d_149_v15_: _dafny.Map
            d_149_v15_ = _dafny.Map({(0) - (d_136_v3_): d_133_v0_})
            d_150_v16_: _dafny.Array
            def lambda6_(d_151_v0_):
                def lambda7_(d_152_i0_):
                    return not(d_151_v0_)

                return lambda7_

            init3_ = lambda6_(d_133_v0_)
            nw21_ = _dafny.Array(None, 7)
            for i0_3_ in range(nw21_.length(0)):
                nw21_[i0_3_] = init3_(i0_3_)
            d_150_v16_ = nw21_
            d_153_v17_: _dafny.Map
            d_153_v17_ = _dafny.Map({d_149_v15_: d_150_v16_})
            default__.m0(d_153_v17_, d_138_globalState_)
            (d_138_globalState_).f11 = d_136_v3_
            if ((d_149_v15_)[d_136_v3_] if (d_136_v3_) in (d_149_v15_) else d_133_v0_):
                d_154_v18_: _dafny.MultiSet
                d_154_v18_ = _dafny.MultiSet([d_136_v3_, d_136_v3_, (0) - (d_136_v3_)])
                d_133_v0_ = ((d_149_v15_)[(d_154_v18_).cardinality] if ((d_154_v18_).cardinality) in (d_149_v15_) else d_133_v0_)
                default__.m0((d_153_v17_) | (_dafny.Map({_dafny.Map({d_136_v3_: d_133_v0_}): d_150_v16_})), d_138_globalState_)
                d_155_v19_: C0
                nw22_ = C0()
                nw22_.ctor__(d_133_v0_, not((d_133_v0_) or (d_133_v0_)), d_134_v1_, 385, True)
                d_155_v19_ = nw22_
                d_156_v21_: _dafny.Map
                d_156_v21_ = _dafny.Map({_dafny.CodePoint('m'): d_136_v3_})
                def iife14_():
                    coll8_ = _dafny.Map()
                    compr_8_: str
                    for compr_8_ in (d_156_v21_).keys.Elements:
                        d_157_v20_: str = compr_8_
                        if (d_157_v20_) in (d_156_v21_):
                            coll8_[d_157_v20_] = (d_155_v19_).f17
                    return _dafny.Map(coll8_)
                (d_138_globalState_).f11 = len(iife14_()
                )
                d_158_v22_: D0
                d_158_v22_ = D0_DC0(True)
                (d_138_globalState_).f0 = (d_158_v22_).cf0
            elif True:
                d_159_v23_: str
                d_159_v23_ = _dafny.CodePoint('r')
                d_160_v24_: _dafny.Seq
                d_160_v24_ = _dafny.SeqWithoutIsStrInference([382])
                d_161_v25_: _dafny.Map
                d_161_v25_ = _dafny.Map({True: d_159_v23_})
                d_159_v23_ = (d_134_v1_)[default__.safeIndex((d_160_v24_)[default__.safeIndex(len(d_161_v25_), len(d_160_v24_))], len(d_134_v1_))]
                default__.m0((d_153_v17_) | (d_153_v17_), d_138_globalState_)
                d_162_v26_: _dafny.MultiSet
                d_162_v26_ = _dafny.MultiSet([d_133_v0_])
                rhs9_ = not(d_133_v0_)
                rhs10_ = (d_136_v3_) - ((len(d_135_v2_) if d_133_v0_ else d_136_v3_))
                rhs11_ = d_136_v3_
                rhs12_ = not(((_dafny.MultiSet([d_133_v0_])) - (d_162_v26_)).issubset(d_162_v26_))
                lhs4_ = d_138_globalState_
                lhs5_ = d_138_globalState_
                lhs6_ = d_138_globalState_
                lhs4_.f0 = rhs9_
                d_136_v3_ = rhs10_
                lhs5_.f11 = rhs11_
                lhs6_.f3 = rhs12_
                d_163_v27_: _dafny.Map
                d_163_v27_ = _dafny.Map({(d_136_v3_) + ((0) - (d_136_v3_)): (_dafny.MultiSet([d_133_v0_, d_133_v0_])).set(d_133_v0_, default__.abs(d_136_v3_))})
                d_163_v27_ = (D2_DC6(d_163_v27_)).cf11
                default__.m0(_dafny.Map({d_149_v15_: d_150_v16_}), d_138_globalState_)
            (d_138_globalState_).f3 = True
            (d_138_globalState_).f3 = not(d_133_v0_)
        d_164_v28_: _dafny.Array
        def lambda8_(d_165_v0_):
            def lambda9_(d_166_i2_):
                return (d_165_v0_ if d_165_v0_ else d_165_v0_)

            return lambda9_

        init4_ = lambda8_(d_133_v0_)
        nw23_ = _dafny.Array(None, 7)
        for i0_4_ in range(nw23_.length(0)):
            nw23_[i0_4_] = init4_(i0_4_)
        d_164_v28_ = nw23_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_164_v28_).length(0)):
            d_167_i1_: int = guard_loop_1_
            if (True) and (((0) <= (d_167_i1_)) and ((d_167_i1_) < ((d_164_v28_).length(0)))):
                (d_164_v28_)[(d_167_i1_)] = d_133_v0_
        d_168_v29_: _dafny.Seq
        d_168_v29_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_133_v0_]), ((d_137_v4_) + (d_137_v4_)).set(default__.safeIndex((0) - (d_136_v3_), len((d_137_v4_) + (d_137_v4_))), d_133_v0_), d_135_v2_])
        d_168_v29_ = ((_dafny.SeqWithoutIsStrInference([d_137_v4_ for d_169_i3_ in range(default__.abs(659))])) + (_dafny.SeqWithoutIsStrInference([d_137_v4_, (d_135_v2_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tknb"))), len(d_135_v2_)), d_133_v0_), d_137_v4_]))) + (d_168_v29_)
        d_136_v3_ = (default__.safeModuloInt(d_136_v3_, d_136_v3_)) + (d_136_v3_)
        if (d_136_v3_) >= (d_136_v3_):
            d_170_v30_: _dafny.Array
            def lambda10_(d_171_v4_):
                def lambda11_(d_172_i4_):
                    return default__.safeModuloInt(d_172_i4_, len(d_171_v4_))

                return lambda11_

            init5_ = lambda10_(d_137_v4_)
            nw24_ = _dafny.Array(None, 9)
            for i0_5_ in range(nw24_.length(0)):
                nw24_[i0_5_] = init5_(i0_5_)
            d_170_v30_ = nw24_
            d_173_v31_: T0
            nw25_ = C1()
            nw25_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsgjrgt")), d_137_v4_, d_136_v3_, d_133_v0_)
            d_173_v31_ = nw25_
            d_174_v32_: _dafny.Map
            d_174_v32_ = _dafny.Map({d_173_v31_: (d_173_v31_).f15})
            d_175_v33_: T1
            nw26_ = C0()
            nw26_.ctor__(True, d_133_v0_, d_134_v1_, d_136_v3_, (d_137_v4_)[default__.safeIndex(len(d_174_v32_), len(d_137_v4_))])
            d_175_v33_ = nw26_
            d_136_v3_ = len(_dafny.Map({d_170_v30_: d_175_v33_}))
            d_176_v34_: _dafny.MultiSet
            d_176_v34_ = _dafny.MultiSet([(d_173_v31_.f14) * (d_136_v3_)])
            d_176_v34_ = _dafny.MultiSet([d_175_v33_.f14, (d_173_v31_.f14) * ((d_173_v31_).fm1(d_176_v34_, not((d_173_v31_).f15), d_138_globalState_)), d_173_v31_.f14, 297])
            d_177_v35_: D1
            d_177_v35_ = D1_DC4(d_136_v3_, (0) - (d_173_v31_.f14))
            source4_ = d_177_v35_
            if source4_.is_DC4:
                d_178___mcc_h0_ = source4_.cf5
                d_179___mcc_h1_ = source4_.cf6
                d_180_cf6_ = d_179___mcc_h1_
                d_181_cf5_ = d_178___mcc_h0_
                d_182_v36_: D0
                d_182_v36_ = D0_DC0(d_133_v0_)
                d_183_v37_: D2
                d_183_v37_ = D2_DC7(not((d_175_v33_).f15), (d_175_v33_).f15, (0) - ((_dafny.MultiSet([d_182_v36_, d_182_v36_, d_182_v36_])).cardinality))
                d_184_v38_: C2
                nw27_ = C2()
                nw27_.ctor__(d_183_v37_, (d_175_v33_.f14) * (d_173_v31_.f14), not(True))
                d_184_v38_ = nw27_
                d_185_v39_: C2
                nw28_ = C2()
                nw28_.ctor__(D2_DC7(d_133_v0_, d_133_v0_, d_173_v31_.f14), default__.safeDivisionInt(d_175_v33_.f14, d_180_cf6_), (d_184_v38_).fm0((d_135_v2_)[default__.safeIndex(d_181_cf5_, len(d_135_v2_))], d_138_globalState_))
                d_185_v39_ = nw28_
                d_186_v40_: _dafny.Array
                nw29_ = _dafny.Array(D2.default()(), 24)
                d_186_v40_ = nw29_
                index8_ = default__.safeIndex(158, (d_186_v40_).length(0))
                (d_186_v40_)[index8_] = (d_184_v38_).f19
                index9_ = default__.safeIndex(158, (d_186_v40_).length(0))
                (d_186_v40_)[index9_] = (d_185_v39_).f19
                d_137_v4_ = d_137_v4_
            elif source4_.is_DC5:
                d_187___mcc_h2_ = source4_.cf7
                d_188___mcc_h3_ = source4_.cf8
                d_189___mcc_h4_ = source4_.cf9
                d_190___mcc_h5_ = source4_.cf10
                d_191_cf10_ = d_190___mcc_h5_
                d_192_cf9_ = d_189___mcc_h4_
                d_193_cf8_ = d_188___mcc_h3_
                d_194_cf7_ = d_187___mcc_h2_
                d_195_v41_: _dafny.Map
                d_195_v41_ = _dafny.Map({d_133_v0_: (d_175_v33_).f15})
                d_195_v41_ = (d_195_v41_).set((d_173_v31_).f15, (d_175_v33_).f15)
                d_196_v42_: _dafny.Array
                nw30_ = _dafny.Array(_dafny.Seq({}), 7)
                d_196_v42_ = nw30_
                d_197_v43_: D5
                d_197_v43_ = D5_DC14(d_175_v33_.f14, d_196_v42_, d_133_v0_)
                d_198_v44_: _dafny.MultiSet
                d_198_v44_ = _dafny.MultiSet([d_197_v43_, d_197_v43_, d_197_v43_, d_197_v43_, d_197_v43_])
                d_198_v44_ = ((d_198_v44_).set(d_197_v43_, default__.abs((0) - (d_193_cf8_)))) | (d_198_v44_)
                (d_138_globalState_).f3 = d_133_v0_
                d_199_v45_: D6
                d_199_v45_ = D6_DC17(True, d_175_v33_)
                (d_138_globalState_).f0 = (d_199_v45_).cf30
            elif True:
                d_200___mcc_h6_ = source4_.cf4
                d_201_cf4_ = d_200___mcc_h6_
                d_202_v46_: _dafny.Array
                nw31_ = _dafny.Array(None, 10)
                d_202_v46_ = nw31_
                d_203_v47_: C1
                nw32_ = C1()
                nw32_.ctor__(d_201_cf4_, d_137_v4_, d_175_v33_.f14, d_133_v0_)
                d_203_v47_ = nw32_
                index10_ = default__.safeIndex(578, (d_202_v46_).length(0))
                (d_202_v46_)[index10_] = d_203_v47_
                index11_ = default__.safeIndex(475, (d_164_v28_).length(0))
                (d_164_v28_)[index11_] = (d_175_v33_).f15
                index12_ = default__.safeIndex(578, (d_202_v46_).length(0))
                index13_ = default__.safeIndex(475, (d_164_v28_).length(0))
                rhs13_ = d_203_v47_
                rhs14_ = (d_173_v31_).f15
                rhs15_ = d_173_v31_.f14
                lhs7_ = d_202_v46_
                lhs8_ = default__.safeIndex(578, (d_202_v46_).length(0))
                lhs9_ = d_164_v28_
                lhs10_ = default__.safeIndex(475, (d_164_v28_).length(0))
                lhs11_ = d_138_globalState_
                lhs7_[lhs8_] = rhs13_
                lhs9_[lhs10_] = rhs14_
                lhs11_.f11 = rhs15_
                d_204_v48_: _dafny.Seq
                d_204_v48_ = _dafny.SeqWithoutIsStrInference([d_175_v33_])
                d_205_v49_: _dafny.Map
                d_205_v49_ = _dafny.Map({d_133_v0_: len(d_204_v48_)})
                (d_173_v31_).f14 = default__.safeModuloInt(len(d_205_v49_), d_175_v33_.f14)
                index14_ = default__.safeIndex(547, (d_170_v30_).length(0))
                (d_170_v30_)[index14_] = d_175_v33_.f14
                d_206_v50_: str
                d_206_v50_ = _dafny.CodePoint('c')
                d_207_v51_: _dafny.Set
                d_207_v51_ = _dafny.Set({d_206_v50_, d_206_v50_})
                d_208_v52_: _dafny.Array
                nw33_ = _dafny.Array(None, 2)
                nw33_[int(0)] = _dafny.Set({d_206_v50_})
                nw33_[int(1)] = d_207_v51_
                d_208_v52_ = nw33_
                index15_ = default__.safeIndex(683, (d_208_v52_).length(0))
                (d_208_v52_)[index15_] = d_207_v51_
                d_209_v53_: _dafny.Array
                nw34_ = _dafny.Array(None, 28)
                d_209_v53_ = nw34_
                d_210_v54_: C0
                nw35_ = C0()
                nw35_.ctor__(d_133_v0_, (d_173_v31_).f15, d_201_cf4_, d_173_v31_.f14, d_133_v0_)
                d_210_v54_ = nw35_
                index16_ = default__.safeIndex(471, (d_209_v53_).length(0))
                (d_209_v53_)[index16_] = d_210_v54_
                index17_ = default__.safeIndex(547, (d_170_v30_).length(0))
                index18_ = default__.safeIndex(683, (d_208_v52_).length(0))
                index19_ = default__.safeIndex(471, (d_209_v53_).length(0))
                rhs16_ = d_136_v3_
                rhs17_ = (_dafny.Set({d_206_v50_})) | (d_207_v51_)
                rhs18_ = d_210_v54_
                lhs12_ = d_170_v30_
                lhs13_ = default__.safeIndex(547, (d_170_v30_).length(0))
                lhs14_ = d_208_v52_
                lhs15_ = default__.safeIndex(683, (d_208_v52_).length(0))
                lhs16_ = d_209_v53_
                lhs17_ = default__.safeIndex(471, (d_209_v53_).length(0))
                lhs12_[lhs13_] = rhs16_
                lhs14_[lhs15_] = rhs17_
                lhs16_[lhs17_] = rhs18_
                index20_ = default__.safeIndex(475, (d_164_v28_).length(0))
                (d_164_v28_)[index20_] = d_133_v0_
            d_211_v55_: _dafny.Array
            def lambda12_(d_212_i5_):
                return _dafny.CodePoint('k')

            init6_ = lambda12_
            nw36_ = _dafny.Array(None, 3)
            for i0_6_ in range(nw36_.length(0)):
                nw36_[i0_6_] = init6_(i0_6_)
            d_211_v55_ = nw36_
            d_213_v56_: str
            d_213_v56_ = _dafny.CodePoint('y')
            index21_ = default__.safeIndex(457, (d_211_v55_).length(0))
            (d_211_v55_)[index21_] = (d_213_v56_ if (d_173_v31_).f15 else _dafny.CodePoint('t'))
            d_214_v57_: _dafny.Array
            def lambda13_(d_215_i6_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esgshe"))

            init7_ = lambda13_
            nw37_ = _dafny.Array(None, 15)
            for i0_7_ in range(nw37_.length(0)):
                nw37_[i0_7_] = init7_(i0_7_)
            d_214_v57_ = nw37_
            d_216_v58_: _dafny.Array
            d_216_v58_ = d_214_v57_
            d_217_v59_: D6
            d_217_v59_ = D6_DC18(d_173_v31_.f14, d_213_v56_)
            index22_ = default__.safeIndex(100, (d_170_v30_).length(0))
            (d_170_v30_)[index22_] = (d_217_v59_).cf32
            index23_ = default__.safeIndex(457, (d_211_v55_).length(0))
            index24_ = default__.safeIndex(100, (d_170_v30_).length(0))
            rhs19_ = (d_173_v31_).f15
            rhs20_ = d_213_v56_
            rhs21_ = d_214_v57_
            rhs22_ = d_175_v33_.f14
            rhs23_ = (0) - (-341)
            lhs18_ = d_211_v55_
            lhs19_ = default__.safeIndex(457, (d_211_v55_).length(0))
            lhs20_ = d_173_v31_
            lhs21_ = d_170_v30_
            lhs22_ = default__.safeIndex(100, (d_170_v30_).length(0))
            d_133_v0_ = rhs19_
            lhs18_[lhs19_] = rhs20_
            d_216_v58_ = rhs21_
            lhs20_.f14 = rhs22_
            lhs21_[lhs22_] = rhs23_
            (d_138_globalState_).f3 = d_133_v0_
        elif True:
            (d_138_globalState_).f11 = d_136_v3_
            d_218_v60_: C1
            nw38_ = C1()
            nw38_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daadh")), d_135_v2_, d_136_v3_, d_133_v0_)
            d_218_v60_ = nw38_
            d_219_v61_: _dafny.Map
            d_219_v61_ = _dafny.Map({d_218_v60_: d_133_v0_})
            d_219_v61_ = (d_219_v61_).set(d_218_v60_, d_133_v0_)
            if (d_133_v0_) or ((d_136_v3_) != (612)):
                d_220_v62_: str
                d_220_v62_ = _dafny.CodePoint('v')
                d_221_v63_: _dafny.Set
                d_221_v63_ = _dafny.Set({(0) - (d_136_v3_)})
                d_222_v64_: _dafny.Seq
                d_222_v64_ = _dafny.SeqWithoutIsStrInference([default__.fm12(d_136_v3_, (d_218_v60_).fm0(d_133_v0_, d_138_globalState_), 699, d_220_v62_, d_138_globalState_), d_221_v63_])
                d_223_v65_: _dafny.Array
                nw39_ = _dafny.Array(None, 5)
                nw39_[int(0)] = (default__.fm17(d_136_v3_, d_138_globalState_)) + (d_222_v64_)
                nw39_[int(1)] = d_222_v64_
                nw39_[int(2)] = d_222_v64_
                nw39_[int(3)] = d_222_v64_
                nw39_[int(4)] = d_222_v64_
                d_223_v65_ = nw39_
                index25_ = default__.safeIndex(49, (d_223_v65_).length(0))
                (d_223_v65_)[index25_] = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_136_v3_}) for d_224_i7_ in range(default__.abs(188))])
                d_225_v66_: _dafny.Seq
                d_225_v66_ = _dafny.SeqWithoutIsStrInference([d_222_v64_, d_222_v64_, d_222_v64_])
                d_226_v67_: _dafny.Map
                d_226_v67_ = _dafny.Map({595: d_133_v0_})
                d_227_v68_: _dafny.MultiSet
                d_227_v68_ = _dafny.MultiSet([len((d_218_v60_).fm8(d_133_v0_, d_138_globalState_)), d_136_v3_, (0) - (d_136_v3_), d_136_v3_, d_136_v3_])
                index26_ = default__.safeIndex(49, (d_223_v65_).length(0))
                (d_223_v65_)[index26_] = (d_225_v66_)[default__.safeIndex((d_136_v3_ if ((d_226_v67_)[d_136_v3_] if (d_136_v3_) in (d_226_v67_) else not(d_133_v0_)) else (d_218_v60_).fm1(d_227_v68_, d_133_v0_, d_138_globalState_)), len(d_225_v66_))]
                d_228_v69_: T1
                nw40_ = C0()
                nw40_.ctor__(not(d_133_v0_), d_133_v0_, (d_218_v60_).f20, 224, d_133_v0_)
                d_228_v69_ = nw40_
                d_229_v70_: D2
                d_229_v70_ = D2_DC7(False, False, -533)
                d_230_v71_: C2
                nw41_ = C2()
                nw41_.ctor__(d_229_v70_, len(_dafny.Set({d_136_v3_})), (d_228_v69_).f15)
                d_230_v71_ = nw41_
                d_231_v72_: _dafny.Map
                d_231_v72_ = _dafny.Map({d_228_v69_: d_230_v71_})
                (d_138_globalState_).f11 = len((d_231_v72_) | (_dafny.Map({d_228_v69_: d_230_v71_})))
                d_226_v67_ = (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhwepn"))): not((d_228_v69_).f15)})) | (_dafny.Map({(d_227_v68_).cardinality: d_133_v0_}))
                (d_138_globalState_).f11 = (0) - (d_136_v3_)
                index27_ = default__.safeIndex(625, (d_164_v28_).length(0))
                (d_164_v28_)[index27_] = (d_228_v69_).f15
                d_232_v73_: _dafny.Map
                d_232_v73_ = _dafny.Map({d_220_v62_: _dafny.SeqWithoutIsStrInference([d_220_v62_ for d_233_i8_ in range(default__.abs(50))])})
                d_234_v74_: _dafny.Map
                d_234_v74_ = _dafny.Map({(d_228_v69_).f15: d_133_v0_})
                d_235_v75_: _dafny.Map
                d_235_v75_ = _dafny.Map({(d_234_v74_ if (d_218_v60_).fm0(d_133_v0_, d_138_globalState_) else d_234_v74_): d_133_v0_})
                index28_ = default__.safeIndex(625, (d_164_v28_).length(0))
                def iife15_():
                    coll9_ = _dafny.Map()
                    compr_9_: _dafny.Map
                    for compr_9_ in (_dafny.SeqWithoutIsStrInference([d_234_v74_, d_234_v74_, d_234_v74_])).Elements:
                        d_236_v76_: _dafny.Map = compr_9_
                        if (d_236_v76_) in (_dafny.SeqWithoutIsStrInference([d_234_v74_, d_234_v74_, d_234_v74_])):
                            coll9_[d_236_v76_] = d_133_v0_
                    return _dafny.Map(coll9_)
                rhs24_ = d_133_v0_
                rhs25_ = (d_232_v73_).set(d_220_v62_, d_134_v1_)
                rhs26_ = ((d_235_v75_) | (_dafny.Map({d_234_v74_: d_133_v0_}))) | (iife15_()
                )
                rhs27_ = d_230_v71_
                rhs28_ = d_136_v3_
                lhs23_ = d_164_v28_
                lhs24_ = default__.safeIndex(625, (d_164_v28_).length(0))
                lhs25_ = d_138_globalState_
                lhs23_[lhs24_] = rhs24_
                d_232_v73_ = rhs25_
                d_235_v75_ = rhs26_
                d_230_v71_ = rhs27_
                lhs25_.f11 = rhs28_
            elif True:
                d_237_v77_: str
                d_237_v77_ = _dafny.CodePoint('y')
                d_238_v78_: _dafny.Set
                d_238_v78_ = _dafny.Set({d_237_v77_})
                d_239_v79_: T0
                nw42_ = C1()
                nw42_.ctor__((d_134_v1_).set(default__.safeIndex(len(d_238_v78_), len(d_134_v1_)), d_237_v77_), (d_218_v60_).f21, d_136_v3_, (d_218_v60_).fm0(True, d_138_globalState_))
                d_239_v79_ = nw42_
                d_239_v79_ = d_239_v79_
                (d_138_globalState_).f11 = (0) - (d_239_v79_.f14)
                d_240_v80_: T1
                nw43_ = C0()
                nw43_.ctor__(d_133_v0_, d_133_v0_, (d_218_v60_).f20, d_136_v3_, d_133_v0_)
                d_240_v80_ = nw43_
                d_241_v81_: _dafny.Map
                d_241_v81_ = _dafny.Map({d_134_v1_: d_240_v80_})
                d_242_v82_: _dafny.Array
                def lambda14_(d_243_v80_):
                    def lambda15_(d_244_i9_):
                        return default__.safeModuloInt(d_244_i9_, d_243_v80_.f14)

                    return lambda15_

                init8_ = lambda14_(d_240_v80_)
                nw44_ = _dafny.Array(None, 21)
                for i0_8_ in range(nw44_.length(0)):
                    nw44_[i0_8_] = init8_(i0_8_)
                d_242_v82_ = nw44_
                d_245_v83_: _dafny.Set
                d_245_v83_ = _dafny.Set({d_242_v82_, d_242_v82_, d_242_v82_})
                d_246_v84_: _dafny.Map
                d_246_v84_ = _dafny.Map({default__.safeModuloInt(d_136_v3_, len(d_241_v81_)): len(d_245_v83_)})
                d_246_v84_ = (d_246_v84_).set(d_240_v80_.f14, d_240_v80_.f14)
                d_247_v85_: _dafny.Map
                d_247_v85_ = _dafny.Map({d_136_v3_: (d_239_v79_).f15})
                d_248_v86_: _dafny.Map
                d_248_v86_ = _dafny.Map({(d_247_v85_).set(d_240_v80_.f14, d_133_v0_): d_164_v28_})
                default__.m0(d_248_v86_, d_138_globalState_)
                (d_138_globalState_).f11 = (d_239_v79_.f14) - (-435)
            d_249_v87_: _dafny.MultiSet
            d_249_v87_ = _dafny.MultiSet([(0) - (d_136_v3_)])
            d_250_v88_: T0
            nw45_ = C1()
            nw45_.ctor__(((d_218_v60_).f20) + ((d_218_v60_).f20), default__.fm18(d_138_globalState_), (len(d_135_v2_) if d_133_v0_ else (d_249_v87_).cardinality), not(d_133_v0_))
            d_250_v88_ = nw45_
            d_251_v89_: _dafny.Seq
            d_251_v89_ = _dafny.SeqWithoutIsStrInference([d_250_v88_])
            d_250_v88_ = (d_251_v89_)[default__.safeIndex((d_136_v3_) - (d_250_v88_.f14), len(d_251_v89_))]
            d_252_v90_: D2
            d_252_v90_ = D2_DC7((d_250_v88_).f15, d_133_v0_, d_136_v3_)
            nw46_ = C2()
            nw46_.ctor__(d_252_v90_, d_250_v88_.f14, (d_250_v88_).f15)
            d_250_v88_ = nw46_
        hi0_ = 583
        for d_253_i10_ in range(default__.safeDivisionInt(d_136_v3_, d_136_v3_), hi0_):
            d_254_v91_: _dafny.Map
            d_254_v91_ = _dafny.Map({d_253_i10_: d_253_i10_})
            d_255_v92_: D2
            d_255_v92_ = D2_DC7(d_133_v0_, False, d_253_i10_)
            d_256_v93_: _dafny.Set
            d_256_v93_ = _dafny.Set({D2_DC7(False, d_133_v0_, ((d_254_v91_)[d_136_v3_] if (d_136_v3_) in (d_254_v91_) else d_136_v3_)), d_255_v92_})
            d_257_v94_: D6
            d_257_v94_ = D6_DC16(_dafny.Set({d_256_v93_}))
            d_258_v95_: _dafny.Set
            d_258_v95_ = _dafny.Set({d_256_v93_})
            pat_let_tv8_ = d_258_v95_
            def iife16_(_pat_let3_0):
                def iife17_(d_259_dt__update__tmp_h1_):
                    def iife18_(_pat_let4_0):
                        def iife19_(d_260_dt__update_hcf29_h0_):
                            return D6_DC16(d_260_dt__update_hcf29_h0_)
                        return iife19_(_pat_let4_0)
                    return iife18_(pat_let_tv8_)
                return iife17_(_pat_let3_0)
            d_257_v94_ = iife16_(D6_DC16(d_258_v95_))
            d_261_v96_: str
            d_261_v96_ = _dafny.CodePoint('q')
            d_262_v97_: T0
            nw47_ = C1()
            nw47_.ctor__(d_134_v1_, d_135_v2_, d_253_i10_, d_133_v0_)
            d_262_v97_ = nw47_
            d_263_v98_: _dafny.Map
            d_263_v98_ = _dafny.Map({d_262_v97_: (d_262_v97_).f15})
            d_264_v99_: _dafny.Seq
            d_264_v99_ = _dafny.SeqWithoutIsStrInference([len(d_263_v98_)])
            d_265_v100_: T1
            nw48_ = C0()
            nw48_.ctor__((d_262_v97_).f15, (d_262_v97_).f15, d_134_v1_, len(_dafny.SeqWithoutIsStrInference([d_261_v96_ for d_266_i11_ in range(default__.abs(-637))])), d_133_v0_)
            d_265_v100_ = nw48_
            d_267_v101_: _dafny.Map
            d_267_v101_ = _dafny.Map({d_136_v3_: d_265_v100_})
            d_268_v102_: _dafny.Map
            d_268_v102_ = _dafny.Map({default__.fm15(d_261_v96_, d_253_i10_, len(d_264_v99_), d_138_globalState_): default__.fm20(647, len(d_267_v101_), d_133_v0_, d_138_globalState_)})
            d_269_v103_: D6
            d_269_v103_ = D6_DC17(d_133_v0_, d_265_v100_)
            d_270_v104_: _dafny.Seq
            d_270_v104_ = _dafny.SeqWithoutIsStrInference([(d_269_v103_).cf31])
            d_271_v106_: _dafny.Map
            def iife20_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(418, 104):
                    d_272_v105_: int = compr_10_
                    if ((418) <= (d_272_v105_)) and ((d_272_v105_) < (104)):
                        coll10_[(d_272_v105_) - (d_262_v97_.f14)] = d_133_v0_
                return _dafny.Map(coll10_)
            d_271_v106_ = _dafny.Map({iife20_()
            : (d_262_v97_).f15})
            d_273_v107_: _dafny.Array
            nw49_ = _dafny.Array(int(0), 4)
            d_273_v107_ = nw49_
            d_274_v108_: _dafny.Seq
            d_274_v108_ = _dafny.SeqWithoutIsStrInference([d_273_v107_, d_273_v107_])
            d_275_v109_: _dafny.Set
            d_275_v109_ = _dafny.Set({(d_265_v100_).f15})
            d_276_v110_: _dafny.Map
            d_276_v110_ = _dafny.Map({d_164_v28_: len(d_275_v109_)})
            d_277_v111_: _dafny.Map
            d_277_v111_ = _dafny.Map({d_133_v0_: d_136_v3_})
            d_278_v112_: _dafny.Map
            d_278_v112_ = _dafny.Map({len(d_277_v111_): (d_262_v97_).f15})
            d_279_v113_: _dafny.Array
            nw50_ = _dafny.Array(None, 27)
            nw50_[int(0)] = d_253_i10_
            nw50_[int(1)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ninu")))
            nw50_[int(2)] = len(d_270_v104_)
            nw50_[int(3)] = -917
            nw50_[int(4)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))
            nw50_[int(5)] = d_136_v3_
            nw50_[int(6)] = len(d_134_v1_)
            nw50_[int(7)] = d_136_v3_
            nw50_[int(8)] = len(d_271_v106_)
            nw50_[int(9)] = len(d_274_v108_)
            nw50_[int(10)] = d_253_i10_
            nw50_[int(11)] = d_253_i10_
            nw50_[int(12)] = len(d_135_v2_)
            nw50_[int(13)] = d_136_v3_
            nw50_[int(14)] = d_253_i10_
            nw50_[int(15)] = d_136_v3_
            nw50_[int(16)] = d_136_v3_
            nw50_[int(17)] = len(d_134_v1_)
            nw50_[int(18)] = (d_264_v99_)[default__.safeIndex(len(d_135_v2_), len(d_264_v99_))]
            nw50_[int(19)] = 413
            nw50_[int(20)] = d_253_i10_
            nw50_[int(21)] = len(d_276_v110_)
            nw50_[int(22)] = d_265_v100_.f14
            nw50_[int(23)] = (0) - (len(d_278_v112_))
            nw50_[int(24)] = d_265_v100_.f14
            nw50_[int(25)] = d_253_i10_
            nw50_[int(26)] = len(_dafny.Map({d_133_v0_: (d_265_v100_).f15}))
            d_279_v113_ = nw50_
            d_280_v114_: _dafny.MultiSet
            d_280_v114_ = _dafny.MultiSet([d_279_v113_, d_279_v113_, d_279_v113_])
            d_281_v115_: _dafny.Seq
            d_281_v115_ = _dafny.SeqWithoutIsStrInference([default__.fm19(d_138_globalState_), len(d_268_v102_), (d_280_v114_).cardinality, d_136_v3_])
            d_281_v115_ = d_264_v99_
            d_282_v116_: _dafny.Map
            d_282_v116_ = _dafny.Map({d_261_v96_: D1_DC3(d_134_v1_)})
            d_283_v117_: _dafny.Map
            d_283_v117_ = _dafny.Map({d_279_v113_: d_282_v116_})
            d_136_v3_ = len(_dafny.Map({(d_262_v97_.f14) + (d_262_v97_.f14): ((d_283_v117_)[d_279_v113_] if (d_279_v113_) in (d_283_v117_) else d_282_v116_)}))
            (d_262_v97_).f14 = default__.fm19(d_138_globalState_)
        d_136_v3_ = d_136_v3_
        d_284_v118_: D3
        d_284_v118_ = D3_DC10(d_134_v1_, d_164_v28_, d_133_v0_, d_136_v3_, -289)
        d_285_v119_: _dafny.Array
        nw51_ = _dafny.Array(None, 21)
        nw51_[int(0)] = d_134_v1_
        nw51_[int(1)] = ((d_284_v118_).cf17) + (d_134_v1_)
        nw51_[int(2)] = d_134_v1_
        nw51_[int(3)] = d_134_v1_
        nw51_[int(4)] = d_134_v1_
        nw51_[int(5)] = d_134_v1_
        nw51_[int(6)] = d_134_v1_
        nw51_[int(7)] = d_134_v1_
        nw51_[int(8)] = d_134_v1_
        nw51_[int(9)] = d_134_v1_
        nw51_[int(10)] = default__.fm9(d_136_v3_, d_138_globalState_)
        nw51_[int(11)] = d_134_v1_
        nw51_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwdbsyk"))
        nw51_[int(13)] = d_134_v1_
        nw51_[int(14)] = (d_134_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iiayqcdiq")))
        nw51_[int(15)] = d_134_v1_
        nw51_[int(16)] = d_134_v1_
        nw51_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbwlhhvq"))
        nw51_[int(18)] = d_134_v1_
        nw51_[int(19)] = d_134_v1_
        nw51_[int(20)] = d_134_v1_
        d_285_v119_ = nw51_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_285_v119_).length(0)):
            d_286_i12_: int = guard_loop_2_
            if (True) and (((0) <= (d_286_i12_)) and ((d_286_i12_) < ((d_285_v119_).length(0)))):
                (d_285_v119_)[(d_286_i12_)] = (d_134_v1_ if (False) == (d_133_v0_) else (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_287_i13_ in range(default__.abs(951))])).set(default__.safeIndex(len(_dafny.Set({False, d_133_v0_})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_288_i13_ in range(default__.abs(951))]))), _dafny.CodePoint('g')))
        d_289_v120_: _dafny.Map
        d_289_v120_ = _dafny.Map({d_136_v3_: 622})
        d_290_i14_: int
        d_290_i14_ = 0
        with _dafny.label("0"):
            while (d_133_v0_ if (d_133_v0_) or (d_133_v0_) else (d_133_v0_) == (default__.fm21(d_136_v3_, not(False), not(d_133_v0_), d_289_v120_, d_138_globalState_))):
                with _dafny.c_label("0"):
                    if (d_290_i14_) >= (100):
                        raise _dafny.Break("0")
                    d_290_i14_ = (d_290_i14_) + (1)
                    d_291_v121_: _dafny.Seq
                    d_291_v121_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(len(d_134_v1_), len(d_137_v4_))])
                    d_291_v121_ = ((d_291_v121_).set(default__.safeIndex((0) - ((0) - (default__.fm19(d_138_globalState_))), len(d_291_v121_)), default__.safeDivisionInt(864, d_136_v3_))).set(default__.safeIndex(d_136_v3_, len((d_291_v121_).set(default__.safeIndex((0) - ((0) - (default__.fm19(d_138_globalState_))), len(d_291_v121_)), default__.safeDivisionInt(864, d_136_v3_)))), 366)
                    d_134_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cuplutb"))
                    d_292_v122_: str
                    d_292_v122_ = _dafny.CodePoint('u')
                    if (not((len((d_134_v1_).set(default__.safeIndex(d_136_v3_, len(d_134_v1_)), d_292_v122_))) < (d_136_v3_))) or ((d_136_v3_) != (911)):
                        d_293_v123_: C0
                        nw52_ = C0()
                        nw52_.ctor__((d_136_v3_) == (79), (d_133_v0_) or (d_133_v0_), d_134_v1_, d_136_v3_, (d_133_v0_ if d_133_v0_ else False))
                        d_293_v123_ = nw52_
                        d_294_v124_: _dafny.Map
                        d_294_v124_ = _dafny.Map({d_136_v3_: (d_136_v3_) > (d_136_v3_)})
                        d_294_v124_ = d_294_v124_
                        d_295_v125_: C1
                        nw53_ = C1()
                        nw53_.ctor__(_dafny.SeqWithoutIsStrInference([d_292_v122_ for d_296_i15_ in range(default__.abs(-172))]), d_137_v4_, len(d_134_v1_), (d_293_v123_).f17)
                        d_295_v125_ = nw53_
                        d_295_v125_ = d_295_v125_
                        d_297_v126_: _dafny.Map
                        d_297_v126_ = _dafny.Map({d_294_v124_: d_164_v28_})
                        d_298_v127_: _dafny.Seq
                        d_298_v127_ = _dafny.SeqWithoutIsStrInference([d_297_v126_])
                        default__.m0(((d_298_v127_)[default__.safeIndex(d_136_v3_, len(d_298_v127_))]) | (d_297_v126_), d_138_globalState_)
                        d_299_v128_: _dafny.MultiSet
                        d_299_v128_ = _dafny.MultiSet([len(d_134_v1_), d_136_v3_])
                        (d_138_globalState_).f3 = ((d_294_v124_)[len(default__.fm13((d_299_v128_).cardinality, d_138_globalState_))] if (len(default__.fm13((d_299_v128_).cardinality, d_138_globalState_))) in (d_294_v124_) else (d_293_v123_).f17)
                    elif True:
                        d_300_v129_: _dafny.Array
                        nw54_ = _dafny.Array(int(0), 29)
                        d_300_v129_ = nw54_
                        index29_ = default__.safeIndex(293, (d_300_v129_).length(0))
                        (d_300_v129_)[index29_] = d_136_v3_
                        index30_ = default__.safeIndex(293, (d_300_v129_).length(0))
                        (d_300_v129_)[index30_] = d_136_v3_
                        d_136_v3_ = (d_300_v129_)[default__.safeIndex(293, (d_300_v129_).length(0))]
                        d_301_v130_: D3
                        d_301_v130_ = D3_DC8(_dafny.SeqWithoutIsStrInference([D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psxioq"))) for d_302_i16_ in range(default__.abs(309))]))
                        d_303_v131_: _dafny.MultiSet
                        d_303_v131_ = _dafny.MultiSet([d_133_v0_])
                        d_301_v130_ = (d_301_v130_ if (d_303_v131_).issubset(d_303_v131_) else d_301_v130_)
                        (d_138_globalState_).f0 = (not(d_133_v0_)) or (d_133_v0_)
                        d_304_v132_: _dafny.Map
                        d_304_v132_ = _dafny.Map({(d_300_v129_)[default__.safeIndex(293, (d_300_v129_).length(0))]: d_133_v0_})
                        d_305_v133_: _dafny.Map
                        d_305_v133_ = _dafny.Map({d_304_v132_: d_164_v28_})
                        default__.m0(d_305_v133_, d_138_globalState_)
                    d_306_v134_: _dafny.Map
                    d_306_v134_ = _dafny.Map({d_133_v0_: d_136_v3_})
                    d_307_v135_: _dafny.Map
                    d_307_v135_ = _dafny.Map({len(d_134_v1_): (d_306_v134_) | (_dafny.Map({d_133_v0_: len(d_289_v120_)}))})
                    d_307_v135_ = (d_307_v135_) | (_dafny.Map({935: d_306_v134_}))
                    pass
            pass
        d_308_v136_: _dafny.MultiSet
        d_308_v136_ = _dafny.MultiSet([d_136_v3_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_309_i17_ in range(default__.abs(168))])), d_136_v3_])
        d_310_v137_: _dafny.Seq
        d_310_v137_ = _dafny.SeqWithoutIsStrInference([d_136_v3_, d_136_v3_])
        d_311_v138_: C0
        nw55_ = C0()
        nw55_.ctor__((d_136_v3_) < (d_136_v3_), (d_308_v136_).ispropersubset(d_308_v136_), d_134_v1_, (0) - (len((_dafny.SeqWithoutIsStrInference([d_136_v3_])) + (d_310_v137_))), not(default__.fm21(d_136_v3_, d_133_v0_, d_133_v0_, d_289_v120_, d_138_globalState_)))
        d_311_v138_ = nw55_
        d_312_v139_: _dafny.Array
        nw56_ = _dafny.Array(int(0), 19)
        d_312_v139_ = nw56_
        d_312_v139_ = d_312_v139_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_312_v139_).length(0)):
            d_313_i18_: int = guard_loop_3_
            if (True) and (((0) <= (d_313_i18_)) and ((d_313_i18_) < ((d_312_v139_).length(0)))):
                (d_312_v139_)[(d_313_i18_)] = default__.safeModuloInt(d_313_i18_, d_136_v3_)
        d_314_v140_: _dafny.MultiSet
        d_314_v140_ = _dafny.MultiSet([True])
        (d_138_globalState_).f0 = (d_314_v140_).isdisjoint(d_314_v140_)
        if not((d_311_v138_).f18):
            index31_ = default__.safeIndex(202, (d_164_v28_).length(0))
            (d_164_v28_)[index31_] = (d_311_v138_).f17
            index32_ = default__.safeIndex(202, (d_164_v28_).length(0))
            (d_164_v28_)[index32_] = (d_311_v138_).f18
            d_315_v141_: _dafny.Array
            def lambda16_(d_316_v3_):
                def lambda17_(d_317_i19_):
                    return _dafny.Set({d_316_v3_})

                return lambda17_

            init9_ = lambda16_(d_136_v3_)
            nw57_ = _dafny.Array(None, 23)
            for i0_9_ in range(nw57_.length(0)):
                nw57_[i0_9_] = init9_(i0_9_)
            d_315_v141_ = nw57_
            d_318_v142_: _dafny.Set
            d_318_v142_ = _dafny.Set({((d_308_v136_)[((d_314_v140_)[d_133_v0_] if (d_133_v0_) in (d_314_v140_) else d_136_v3_)] if (((d_314_v140_)[d_133_v0_] if (d_133_v0_) in (d_314_v140_) else d_136_v3_)) in (d_308_v136_) else d_136_v3_)})
            index33_ = default__.safeIndex(511, (d_315_v141_).length(0))
            (d_315_v141_)[index33_] = d_318_v142_
            index34_ = default__.safeIndex(575, (d_312_v139_).length(0))
            (d_312_v139_)[index34_] = d_136_v3_
            index35_ = default__.safeIndex(511, (d_315_v141_).length(0))
            index36_ = default__.safeIndex(575, (d_312_v139_).length(0))
            rhs29_ = d_318_v142_
            rhs30_ = (0) - (d_136_v3_)
            rhs31_ = d_136_v3_
            rhs32_ = (((d_308_v136_)[d_136_v3_] if (d_136_v3_) in (d_308_v136_) else (d_311_v138_).fm2(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_319_i20_ in range(default__.abs(241))]), (d_311_v138_).f18, (d_164_v28_)[default__.safeIndex(202, (d_164_v28_).length(0))], d_136_v3_, d_138_globalState_))) * (d_136_v3_)
            lhs26_ = d_315_v141_
            lhs27_ = default__.safeIndex(511, (d_315_v141_).length(0))
            lhs28_ = d_138_globalState_
            lhs29_ = d_138_globalState_
            lhs30_ = d_312_v139_
            lhs31_ = default__.safeIndex(575, (d_312_v139_).length(0))
            lhs26_[lhs27_] = rhs29_
            lhs28_.f11 = rhs30_
            lhs29_.f11 = rhs31_
            lhs30_[lhs31_] = rhs32_
            d_320_v143_: D2
            d_320_v143_ = D2_DC7((d_311_v138_).f18, d_133_v0_, len(_dafny.Map({len(d_310_v137_): (d_164_v28_)[default__.safeIndex(202, (d_164_v28_).length(0))]})))
            d_321_v144_: C2
            nw58_ = C2()
            nw58_.ctor__(d_320_v143_, (d_314_v140_).cardinality, True)
            d_321_v144_ = nw58_
            d_322_v145_: _dafny.Map
            d_322_v145_ = _dafny.Map({default__.fm21((d_312_v139_)[default__.safeIndex(575, (d_312_v139_).length(0))], (d_311_v138_).f18, d_133_v0_, d_289_v120_, d_138_globalState_): d_321_v144_})
            d_323_v146_: _dafny.Array
            nw59_ = _dafny.Array(None, 24)
            nw59_[int(0)] = d_321_v144_
            nw59_[int(1)] = (d_321_v144_ if (d_311_v138_).f18 else d_321_v144_)
            nw59_[int(2)] = d_321_v144_
            nw59_[int(3)] = d_321_v144_
            nw59_[int(4)] = d_321_v144_
            nw59_[int(5)] = d_321_v144_
            nw59_[int(6)] = d_321_v144_
            nw59_[int(7)] = d_321_v144_
            nw59_[int(8)] = d_321_v144_
            nw59_[int(9)] = ((d_322_v145_)[d_133_v0_] if (d_133_v0_) in (d_322_v145_) else d_321_v144_)
            nw59_[int(10)] = d_321_v144_
            nw59_[int(11)] = d_321_v144_
            nw59_[int(12)] = d_321_v144_
            nw59_[int(13)] = d_321_v144_
            nw59_[int(14)] = d_321_v144_
            nw59_[int(15)] = d_321_v144_
            nw59_[int(16)] = d_321_v144_
            nw59_[int(17)] = d_321_v144_
            nw59_[int(18)] = d_321_v144_
            nw59_[int(19)] = d_321_v144_
            nw59_[int(20)] = d_321_v144_
            nw59_[int(21)] = d_321_v144_
            nw59_[int(22)] = d_321_v144_
            nw59_[int(23)] = d_321_v144_
            d_323_v146_ = nw59_
            index37_ = default__.safeIndex(146, (d_323_v146_).length(0))
            (d_323_v146_)[index37_] = d_321_v144_
            index38_ = default__.safeIndex(146, (d_323_v146_).length(0))
            (d_323_v146_)[index38_] = d_321_v144_
            d_324_v147_: C0
            nw60_ = C0()
            nw60_.ctor__((d_311_v138_).f17, d_133_v0_, d_134_v1_, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "syla")))) - (d_136_v3_), d_133_v0_)
            d_324_v147_ = nw60_
            d_325_v148_: _dafny.Map
            d_325_v148_ = _dafny.Map({d_134_v1_: (d_164_v28_)[default__.safeIndex(202, (d_164_v28_).length(0))]})
            d_326_v149_: _dafny.Map
            d_326_v149_ = _dafny.Map({d_325_v148_: (d_311_v138_).f17})
            index39_ = default__.safeIndex(202, (d_164_v28_).length(0))
            rhs33_ = (d_311_v138_).f17
            rhs34_ = (d_324_v147_).f18
            rhs35_ = (d_164_v28_)[default__.safeIndex(202, (d_164_v28_).length(0))]
            rhs36_ = ((d_326_v149_)[d_325_v148_] if (d_325_v148_) in (d_326_v149_) else (d_324_v147_).f17)
            rhs37_ = (default__.fm21(d_136_v3_, not((d_311_v138_).f17), d_133_v0_, d_289_v120_, d_138_globalState_) if False else not((d_311_v138_).f17))
            lhs32_ = d_164_v28_
            lhs33_ = default__.safeIndex(202, (d_164_v28_).length(0))
            lhs34_ = d_138_globalState_
            lhs35_ = d_138_globalState_
            lhs32_[lhs33_] = rhs33_
            d_133_v0_ = rhs34_
            lhs34_.f3 = rhs35_
            d_133_v0_ = rhs36_
            lhs35_.f0 = rhs37_
        elif True:
            (d_138_globalState_).f3 = (d_311_v138_).f18
            if (d_311_v138_).f18:
                index40_ = default__.safeIndex(570, (d_164_v28_).length(0))
                (d_164_v28_)[index40_] = (d_311_v138_).f17
                index41_ = default__.safeIndex(570, (d_164_v28_).length(0))
                (d_164_v28_)[index41_] = (d_311_v138_).f18
                d_327_v150_: _dafny.Map
                d_327_v150_ = _dafny.Map({d_136_v3_: (d_311_v138_).f17})
                d_328_v151_: _dafny.Array
                nw61_ = _dafny.Array(False, 11)
                d_328_v151_ = nw61_
                d_329_v152_: _dafny.Map
                d_329_v152_ = _dafny.Map({d_327_v150_: d_328_v151_})
                default__.m0((d_329_v152_) | (d_329_v152_), d_138_globalState_)
                d_330_v153_: D2
                d_330_v153_ = D2_DC7((d_311_v138_).f17, (d_311_v138_).f18, d_136_v3_)
                d_331_v154_: C1
                nw62_ = C1()
                nw62_.ctor__(d_134_v1_, d_137_v4_, d_136_v3_, (d_311_v138_).f18)
                d_331_v154_ = nw62_
                d_332_v155_: _dafny.Map
                d_332_v155_ = _dafny.Map({d_331_v154_: not((d_311_v138_).f17)})
                d_333_v156_: C2
                nw63_ = C2()
                nw63_.ctor__(d_330_v153_, len(d_332_v155_), (len(d_135_v2_)) >= (d_136_v3_))
                d_333_v156_ = nw63_
                (d_138_globalState_).f0 = ((d_136_v3_) not in (d_308_v136_) if (d_311_v138_).f17 else (d_311_v138_).f18)
                d_334_v157_: _dafny.Map
                d_334_v157_ = _dafny.Map({d_136_v3_: d_314_v140_})
                d_335_v158_: D2
                d_335_v158_ = D2_DC6(d_334_v157_)
                d_336_v159_: int
                d_337_v160_: int
                out0_: int
                out1_: int
                out0_, out1_ = (d_333_v156_).m2(d_335_v158_, (d_311_v138_).f17, (0) - ((0) - (d_136_v3_)), d_136_v3_, d_138_globalState_)
                d_336_v159_ = out0_
                d_337_v160_ = out1_
            elif True:
                d_338_v161_: _dafny.Array
                def lambda18_(d_339_i21_):
                    return _dafny.CodePoint('j')

                init10_ = lambda18_
                nw64_ = _dafny.Array(None, 24)
                for i0_10_ in range(nw64_.length(0)):
                    nw64_[i0_10_] = init10_(i0_10_)
                d_338_v161_ = nw64_
                d_340_v162_: str
                d_340_v162_ = _dafny.CodePoint('g')
                index42_ = default__.safeIndex(432, (d_338_v161_).length(0))
                (d_338_v161_)[index42_] = d_340_v162_
                index43_ = default__.safeIndex(432, (d_338_v161_).length(0))
                (d_338_v161_)[index43_] = d_340_v162_
                d_341_v163_: D1
                d_341_v163_ = D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tcir")))
                d_341_v163_ = D1_DC3(d_134_v1_)
                (d_138_globalState_).f11 = d_136_v3_
                (d_138_globalState_).f11 = (d_311_v138_).fm1(d_308_v136_, d_133_v0_, d_138_globalState_)
                d_164_v28_ = d_164_v28_
            d_310_v137_ = (d_310_v137_) + (d_310_v137_)
            (d_138_globalState_).f11 = default__.safeDivisionInt(default__.safeModuloInt(d_136_v3_, d_136_v3_), d_136_v3_)
            d_342_v164_: _dafny.Set
            d_342_v164_ = _dafny.Set({(d_311_v138_).f17})
            rhs38_ = d_136_v3_
            rhs39_ = (d_136_v3_) - (d_136_v3_)
            rhs40_ = d_342_v164_
            d_136_v3_ = rhs38_
            d_136_v3_ = rhs39_
            d_342_v164_ = rhs40_
        index44_ = default__.safeIndex(638, (d_164_v28_).length(0))
        (d_164_v28_)[index44_] = (d_311_v138_).fm0((d_311_v138_).f17, d_138_globalState_)
        d_343_v165_: _dafny.Map
        d_343_v165_ = _dafny.Map({255: (d_311_v138_).f18})
        d_344_v166_: D3
        d_344_v166_ = D3_DC9(len(d_343_v165_))
        pat_let_tv9_ = d_311_v138_
        pat_let_tv10_ = d_311_v138_
        pat_let_tv11_ = d_311_v138_
        pat_let_tv12_ = d_311_v138_
        index45_ = default__.safeIndex(638, (d_164_v28_).length(0))
        def lambda19_(source5_):
            if source5_.is_DC9:
                d_345___mcc_h7_ = source5_.cf16
                d_346_cf16_ = d_345___mcc_h7_
                return not((pat_let_tv9_).f18)
            elif source5_.is_DC10:
                d_347___mcc_h8_ = source5_.cf17
                d_348___mcc_h9_ = source5_.cf18
                d_349___mcc_h10_ = source5_.cf19
                d_350___mcc_h11_ = source5_.cf20
                d_351___mcc_h12_ = source5_.cf21
                d_352_cf21_ = d_351___mcc_h12_
                d_353_cf20_ = d_350___mcc_h11_
                d_354_cf19_ = d_349___mcc_h10_
                d_355_cf18_ = d_348___mcc_h9_
                d_356_cf17_ = d_347___mcc_h8_
                return (pat_let_tv10_).f18
            elif source5_.is_DC8:
                d_357___mcc_h13_ = source5_.cf15
                d_358_cf15_ = d_357___mcc_h13_
                return (pat_let_tv11_).f18
            elif True:
                d_359___mcc_h14_ = source5_.cf22
                d_360_cf22_ = d_359___mcc_h14_
                return (pat_let_tv12_).f18

        (d_164_v28_)[index45_] = lambda19_(d_344_v166_)
        d_361_v167_: _dafny.Array
        nw65_ = _dafny.Array(None, 26)
        d_361_v167_ = nw65_
        d_361_v167_ = d_361_v167_
        _dafny.print(_dafny.string_of(d_133_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_134_v1_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v2_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_136_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v4_) == (_dafny.SeqWithoutIsStrInference([False, False, False, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_138_globalState_).f4).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_globalState_.f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f13) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v28_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v28_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v28_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v28_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v28_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v28_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v28_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v29_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True, False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_284_v118_).cf17).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_284_v118_).cf18)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_284_v118_).cf18)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_284_v118_).cf18)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_284_v118_).cf18)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_284_v118_).cf18)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_284_v118_).cf18)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_284_v118_).cf18)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v118_).cf19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v118_).cf20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v118_).cf21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[18]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_285_v119_)[20]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_289_v120_) == (_dafny.Map({1: 622}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_290_i14_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_308_v136_) == (_dafny.MultiSet([1, 1, 168]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_310_v137_) == (_dafny.SeqWithoutIsStrInference([1, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v138_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v138_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v139_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_314_v140_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v165_) == (_dafny.Map({255: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_344_v166_).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
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
        return lambda: D1_DC4(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({self.cf4.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC7(D2, NamedTuple('DC7', [('cf12', Any), ('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)

class D3_DC9(D3, NamedTuple('DC9', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({self.cf17.VerbatimString(True)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC12(D4, NamedTuple('DC12', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14(int(0), _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC14(D5, NamedTuple('DC14', [('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)

class D6_DC17(D6, NamedTuple('DC17', [('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)

class D7_DC20(D7, NamedTuple('DC20', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC22(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)

class D8_DC22(D8, NamedTuple('DC22', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC26(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)

class D9_DC26(D9, NamedTuple('DC26', [('cf42', Any), ('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27)
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC29(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)

class D10_DC29(D10, NamedTuple('DC29', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value

class T1:
    pass
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value

class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self.f3: bool = False
        self.f11: int = int(0)
        self._f1: bool = False
        self._f2: int = int(0)
        self._f4: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f5: bool = False
        self._f6: int = int(0)
        self._f7: bool = False
        self._f8: int = int(0)
        self._f9: int = int(0)
        self._f10: int = int(0)
        self._f12: bool = False
        self._f13: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13

    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
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
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13

class C0(T1, T0):
    def  __init__(self):
        self._f14: int = int(0)
        self._f16: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f15: bool = False
        self._f17: bool = False
        self._f18: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f15(self):
        return self._f15
    def ctor__(self, f17, f18, f16, f14, f15):
        (self)._f17 = f17
        (self)._f18 = f18
        (self).f16 = f16
        (self).f14 = f14
        (self)._f15 = f15

    def fm0(self, p0, globalState):
        return not(not((D0_DC0((self).f18)).cf0))

    def fm1(self, p0, p1, globalState):
        return self.f14

    def fm2(self, p0, p1, p2, p3, globalState):
        def iife21_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(404, 615):
                d_362_v0_: int = compr_11_
                if ((404) <= (d_362_v0_)) and ((d_362_v0_) < (615)):
                    coll11_[(d_362_v0_) + (self.f14)] = self.f14
            return _dafny.Map(coll11_)
        return len(((D1_DC3(self.f16)).cf4).set(default__.safeIndex((D1_DC4(len(iife21_()
), self.f14)).cf5, len((D1_DC3(self.f16)).cf4)), (self.f16)[default__.safeIndex(self.f14, len(self.f16))]))

    def fm3(self, p0, p1, p2, globalState):
        return _dafny.MultiSet([self.f14, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_363_i0_ in range(default__.abs(523))]))), (0) - (len((_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18, (self).f15])) + (_dafny.SeqWithoutIsStrInference([(self).f15]))))])

    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18

class C1(T0):
    def  __init__(self):
        self._f14: int = int(0)
        self._f15: bool = False
        self._f20: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f21: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f15(self):
        return self._f15
    def ctor__(self, f20, f21, f14, f15):
        (self)._f20 = f20
        (self)._f21 = f21
        (self).f14 = f14
        (self)._f15 = f15

    def fm0(self, p0, globalState):
        return (self).f15

    def fm1(self, p0, p1, globalState):
        return (0) - (len(((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")): len(_dafny.SeqWithoutIsStrInference([244, self.f14]))})) | (_dafny.Map({(self).f20: len((self).f20)}))) | ((_dafny.Map({(self).f20: self.f14})) | (_dafny.Map({(self).f20: self.f14})))))

    def fm8(self, p0, globalState):
        source6_ = D0_DC2(self.f14, _dafny.MultiSet([(self).f15]))
        if source6_.is_DC1:
            d_364___mcc_h0_ = source6_.cf1
            d_365_cf1_ = d_364___mcc_h0_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuycpn"))
        elif source6_.is_DC2:
            d_366___mcc_h1_ = source6_.cf2
            d_367___mcc_h2_ = source6_.cf3
            d_368_cf3_ = d_367___mcc_h2_
            d_369_cf2_ = d_366___mcc_h1_
            return (self).f20
        elif True:
            d_370___mcc_h3_ = source6_.cf0
            d_371_cf0_ = d_370___mcc_h3_
            return (self).f20

    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21

class C2(T0):
    def  __init__(self):
        self._f14: int = int(0)
        self._f15: bool = False
        self._f19: D2 = D2.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f15(self):
        return self._f15
    def ctor__(self, f19, f14, f15):
        (self)._f19 = f19
        (self).f14 = f14
        (self)._f15 = f15

    def fm0(self, p0, globalState):
        return (self).f15

    def fm1(self, p0, p1, globalState):
        return 292

    def fm4(self, p0, globalState):
        return _dafny.CodePoint('m')

    def fm5(self, p0, globalState):
        return (_dafny.Set({(self).f15})) | ((_dafny.Set({(self).f15})) - (_dafny.Set({(self).f15})))

    def m1(self, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: D1 = D1.default()()
        r2: bool = False
        r3: int = int(0)
        d_372_v0_: _dafny.Array
        nw66_ = _dafny.Array(False, 22)
        d_372_v0_ = nw66_
        index46_ = default__.safeIndex(514, (d_372_v0_).length(0))
        (d_372_v0_)[index46_] = (self).f15
        d_373_v1_: _dafny.Array
        def lambda20_(d_374_i0_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eafbti"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))

        init11_ = lambda20_
        nw67_ = _dafny.Array(None, 5)
        for i0_11_ in range(nw67_.length(0)):
            nw67_[i0_11_] = init11_(i0_11_)
        d_373_v1_ = nw67_
        d_375_v2_: _dafny.Seq
        d_375_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kaueyjj"))
        index47_ = default__.safeIndex(90, (d_373_v1_).length(0))
        (d_373_v1_)[index47_] = d_375_v2_
        d_376_v3_: str
        d_376_v3_ = _dafny.CodePoint('q')
        index48_ = default__.safeIndex(514, (d_372_v0_).length(0))
        index49_ = default__.safeIndex(90, (d_373_v1_).length(0))
        rhs41_ = (self).f15
        rhs42_ = ((d_375_v2_) + (d_375_v2_)).set(default__.safeIndex((0) - (self.f14), len((d_375_v2_) + (d_375_v2_))), d_376_v3_)
        lhs36_ = d_372_v0_
        lhs37_ = default__.safeIndex(514, (d_372_v0_).length(0))
        lhs38_ = d_373_v1_
        lhs39_ = default__.safeIndex(90, (d_373_v1_).length(0))
        lhs36_[lhs37_] = rhs41_
        lhs38_[lhs39_] = rhs42_
        if (self.f14) in (_dafny.MultiSet([self.f14])):
            (globalState).f11 = self.f14
            (globalState).f3 = ((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]) == ((self).f15)
            d_377_v4_: _dafny.Map
            d_377_v4_ = _dafny.Map({(self).f15: (self.f14) - (self.f14)})
            d_378_v5_: _dafny.MultiSet
            d_378_v5_ = _dafny.MultiSet([(self).f15])
            d_377_v4_ = (d_377_v4_).set((len(d_375_v2_)) > ((d_378_v5_).cardinality), self.f14)
            d_379_v6_: _dafny.Map
            d_379_v6_ = _dafny.Map({(self).f15: not ((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]) or ((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))])})
            r2 = ((d_379_v6_)[(d_377_v4_) == (d_377_v4_)] if ((d_377_v4_) == (d_377_v4_)) in (d_379_v6_) else (self).fm0((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))], globalState))
            d_380_v7_: D0
            d_380_v7_ = D0_DC0(True)
            d_380_v7_ = default__.fm6(False, globalState)
        elif True:
            index50_ = default__.safeIndex(90, (d_373_v1_).length(0))
            (d_373_v1_)[index50_] = d_375_v2_
            d_381_v8_: _dafny.Map
            d_381_v8_ = _dafny.Map({d_376_v3_: (d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))]})
            d_381_v8_ = (d_381_v8_).set(_dafny.CodePoint('r'), (d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))])
            d_382_v9_: _dafny.Seq
            d_382_v9_ = _dafny.SeqWithoutIsStrInference([(d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))], (self).f15])
            d_383_v10_: T0
            nw68_ = C1()
            nw68_.ctor__((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_384_i1_ in range(default__.abs(823))])).set(default__.safeIndex(self.f14, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_385_i1_ in range(default__.abs(823))]))), d_376_v3_), (d_382_v9_).set(default__.safeIndex(self.f14, len(d_382_v9_)), (self).f15), self.f14, (self).f15)
            d_383_v10_ = nw68_
            d_386_v11_: _dafny.Seq
            d_386_v11_ = _dafny.SeqWithoutIsStrInference([d_383_v10_, d_383_v10_])
            d_387_v12_: _dafny.Map
            d_387_v12_ = _dafny.Map({376: (self.f14) * (len((default__.fm7((self).f15, (self).f15, len(d_386_v11_), globalState)).cf15))})
            d_387_v12_ = (d_387_v12_).set(self.f14, d_383_v10_.f14)
            (globalState).f0 = (self).f15
            (globalState).f11 = 727
        hi1_ = (self.f14) * ((0) - (-936))
        for d_388_i2_ in range(self.f14, hi1_):
            d_389_v13_: _dafny.Set
            d_389_v13_ = _dafny.Set({d_376_v3_})
            d_390_v15_: _dafny.Map
            def iife22_():
                coll12_ = _dafny.Set()
                compr_12_: str
                for compr_12_ in ((d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))]).Elements:
                    d_391_v14_: str = compr_12_
                    if (d_391_v14_) in ((d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))]):
                        coll12_ = coll12_.union(_dafny.Set([d_391_v14_]))
                return _dafny.Set(coll12_)
            d_390_v15_ = _dafny.Map({_dafny.MultiSet([-98]): iife22_()
            })
            if (d_389_v13_) == (((d_390_v15_)[_dafny.MultiSet([len(_dafny.Map({d_376_v3_: d_388_i2_}))])] if (_dafny.MultiSet([len(_dafny.Map({d_376_v3_: d_388_i2_}))])) in (d_390_v15_) else d_389_v13_)):
                d_392_v16_: C0
                nw69_ = C0()
                nw69_.ctor__(False, (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))], (d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))], len(d_375_v2_), not((self).f15))
                d_392_v16_ = nw69_
                d_393_v17_: _dafny.Array
                def lambda21_(d_394_i3_):
                    return D3_DC9(self.f14)

                init12_ = lambda21_
                nw70_ = _dafny.Array(None, 26)
                for i0_12_ in range(nw70_.length(0)):
                    nw70_[i0_12_] = init12_(i0_12_)
                d_393_v17_ = nw70_
                d_395_v18_: _dafny.Array
                def lambda22_(d_396_v16_):
                    def lambda23_(d_397_i4_):
                        return default__.safeModuloInt(d_397_i4_, (_dafny.MultiSet([(d_396_v16_).f17])).cardinality)

                    return lambda23_

                init13_ = lambda22_(d_392_v16_)
                nw71_ = _dafny.Array(None, 5)
                for i0_13_ in range(nw71_.length(0)):
                    nw71_[i0_13_] = init13_(i0_13_)
                d_395_v18_ = nw71_
                index51_ = default__.safeIndex(688, (d_395_v18_).length(0))
                (d_395_v18_)[index51_] = self.f14
                d_398_v19_: _dafny.Seq
                d_398_v19_ = _dafny.SeqWithoutIsStrInference([d_393_v17_])
                d_399_v20_: _dafny.Seq
                d_399_v20_ = _dafny.SeqWithoutIsStrInference([(d_393_v17_ if (d_392_v16_).f18 else (d_398_v19_)[default__.safeIndex(self.f14, len(d_398_v19_))])])
                index52_ = default__.safeIndex(688, (d_395_v18_).length(0))
                rhs43_ = (self).f15
                rhs44_ = d_392_v16_
                rhs45_ = (d_399_v20_)[default__.safeIndex(d_388_i2_, len(d_399_v20_))]
                rhs46_ = 474
                lhs40_ = d_395_v18_
                lhs41_ = default__.safeIndex(688, (d_395_v18_).length(0))
                r2 = rhs43_
                d_392_v16_ = rhs44_
                d_393_v17_ = rhs45_
                lhs40_[lhs41_] = rhs46_
                d_400_v21_: _dafny.Seq
                d_400_v21_ = _dafny.SeqWithoutIsStrInference([(d_392_v16_).f18])
                d_401_v22_: _dafny.MultiSet
                d_401_v22_ = _dafny.MultiSet([d_400_v21_])
                d_402_v23_: _dafny.Map
                d_402_v23_ = _dafny.Map({(d_392_v16_).f18: (d_392_v16_).f17})
                d_403_v24_: _dafny.Map
                d_403_v24_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_376_v3_ for d_404_i5_ in range(default__.abs(360))])): self.f14})
                index53_ = default__.safeIndex(514, (d_372_v0_).length(0))
                rhs47_ = (d_392_v16_).f17
                rhs48_ = ((0) - ((0) - ((0) - (self.f14))) if ((d_402_v23_)[True] if (True) in (d_402_v23_) else (d_392_v16_).f17) else (len(default__.fm9(self.f14, globalState))) + (d_388_i2_))
                rhs49_ = d_401_v22_
                rhs50_ = (d_400_v21_)[default__.safeIndex(((d_403_v24_)[769] if (769) in (d_403_v24_) else self.f14), len(d_400_v21_))]
                lhs42_ = globalState
                lhs43_ = globalState
                lhs44_ = d_372_v0_
                lhs45_ = default__.safeIndex(514, (d_372_v0_).length(0))
                lhs42_.f0 = rhs47_
                lhs43_.f11 = rhs48_
                d_401_v22_ = rhs49_
                lhs44_[lhs45_] = rhs50_
                d_405_v25_: _dafny.Array
                d_405_v25_ = d_373_v1_
                d_373_v1_ = (d_405_v25_)
                d_406_v26_: _dafny.Map
                d_406_v26_ = _dafny.Map({not(not((d_392_v16_).f18)): d_372_v0_})
                d_372_v0_ = ((d_406_v26_)[(self).f15] if ((self).f15) in (d_406_v26_) else d_372_v0_)
                d_407_v27_: C0
                nw72_ = C0()
                nw72_.ctor__((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))], (d_392_v16_).f17, (d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))], self.f14, (d_392_v16_).f18)
                d_407_v27_ = nw72_
            elif True:
                (globalState).f11 = d_388_i2_
                d_408_v28_: _dafny.Seq
                d_408_v28_ = _dafny.SeqWithoutIsStrInference([self.f14])
                d_409_v29_: _dafny.Seq
                d_409_v29_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                d_410_v30_: C1
                nw73_ = C1()
                nw73_.ctor__(d_375_v2_, _dafny.SeqWithoutIsStrInference([(self).fm0((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))], globalState), (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]]), (d_388_i2_) + ((d_408_v28_)[default__.safeIndex(d_388_i2_, len(d_408_v28_))]), (d_409_v29_)[default__.safeIndex(self.f14, len(d_409_v29_))])
                d_410_v30_ = nw73_
                (globalState).f11 = (0) - (d_388_i2_)
                (globalState).f3 = (self).f15
                d_411_v31_: _dafny.Array
                def lambda24_(d_412_v0_, d_413_v28_):
                    def lambda25_(d_414_i6_):
                        return (_dafny.Map({(d_412_v0_)[default__.safeIndex(514, (d_412_v0_).length(0))]: d_413_v28_})) | (_dafny.Map({(self).f15: d_413_v28_}))

                    return lambda25_

                init14_ = lambda24_(d_372_v0_, d_408_v28_)
                nw74_ = _dafny.Array(None, 11)
                for i0_14_ in range(nw74_.length(0)):
                    nw74_[i0_14_] = init14_(i0_14_)
                d_411_v31_ = nw74_
                index54_ = default__.safeIndex(929, (d_411_v31_).length(0))
                (d_411_v31_)[index54_] = default__.fm10((self).f15, globalState)
                d_415_v32_: _dafny.MultiSet
                d_415_v32_ = _dafny.MultiSet([d_388_i2_])
                d_416_v33_: _dafny.Map
                d_416_v33_ = _dafny.Map({(self).f15: ((d_408_v28_).set(default__.safeIndex((_dafny.MultiSet(d_408_v28_)).cardinality, len(d_408_v28_)), d_388_i2_)).set(default__.safeIndex((d_410_v30_).fm1(d_415_v32_, (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))], globalState), len((d_408_v28_).set(default__.safeIndex((_dafny.MultiSet(d_408_v28_)).cardinality, len(d_408_v28_)), d_388_i2_))), len(d_375_v2_))})
                index55_ = default__.safeIndex(929, (d_411_v31_).length(0))
                (d_411_v31_)[index55_] = (d_416_v33_).set(False, d_408_v28_)
            d_417_v34_: _dafny.Map
            d_417_v34_ = _dafny.Map({self.f14: (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]})
            d_418_v35_: _dafny.Seq
            d_418_v35_ = _dafny.SeqWithoutIsStrInference([(d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))], (self).f15])
            d_419_v36_: C0
            nw75_ = C0()
            nw75_.ctor__((self).fm0((self).fm0((self).f15, globalState), globalState), not((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_420_i7_ in range(default__.abs(101))])) + (d_375_v2_), len((d_417_v34_).set((0) - (len(d_418_v35_)), ((d_417_v34_)[660] if (660) in (d_417_v34_) else (self).fm0((self).f15, globalState)))), (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))])
            d_419_v36_ = nw75_
            d_421_v37_: _dafny.Seq
            d_421_v37_ = _dafny.SeqWithoutIsStrInference([-45, self.f14, -309, self.f14, self.f14])
            d_422_v38_: _dafny.MultiSet
            d_422_v38_ = _dafny.MultiSet([d_375_v2_])
            d_423_v39_: _dafny.MultiSet
            d_423_v39_ = _dafny.MultiSet([self.f14])
            d_424_v40_: D3
            d_424_v40_ = D3_DC10(d_375_v2_, d_372_v0_, (d_423_v39_).issubset((_dafny.MultiSet(d_421_v37_)).set(((d_422_v38_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxayt"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxayt"))) in (d_422_v38_) else len((d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))])), default__.abs(len((d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))])))), (0) - (d_388_i2_), d_388_i2_)
            d_424_v40_ = d_424_v40_
            d_417_v34_ = (d_417_v34_).set(self.f14, (d_388_i2_) <= (self.f14))
        d_425_v41_: _dafny.Seq
        d_425_v41_ = _dafny.SeqWithoutIsStrInference([((d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))]) + (d_375_v2_), ((d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))] if True else (d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urdc")), ((d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))]) + ((d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))]), (d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))]])
        d_426_v42_: _dafny.Map
        d_426_v42_ = _dafny.Map({828: _dafny.SeqWithoutIsStrInference([d_375_v2_ for d_427_i8_ in range(default__.abs(471))])})
        d_425_v41_ = ((d_426_v42_)[self.f14] if (self.f14) in (d_426_v42_) else d_425_v41_)
        d_428_v43_: _dafny.Array
        def lambda26_(d_429_v0_):
            def lambda27_(d_430_i9_):
                return (d_430_i9_) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f15, (d_429_v0_)[default__.safeIndex(514, (d_429_v0_).length(0))]])) for d_431_i10_ in range(default__.abs(375))])))

            return lambda27_

        init15_ = lambda26_(d_372_v0_)
        nw76_ = _dafny.Array(None, 19)
        for i0_15_ in range(nw76_.length(0)):
            nw76_[i0_15_] = init15_(i0_15_)
        d_428_v43_ = nw76_
        index56_ = default__.safeIndex(2, (d_428_v43_).length(0))
        (d_428_v43_)[index56_] = ((0) - (self.f14)) + (self.f14)
        d_432_v44_: _dafny.Seq
        d_432_v44_ = _dafny.SeqWithoutIsStrInference([self.f14])
        d_433_v45_: _dafny.Set
        d_433_v45_ = _dafny.Set({(d_432_v44_)[default__.safeIndex(self.f14, len(d_432_v44_))]})
        index57_ = default__.safeIndex(14, (d_428_v43_).length(0))
        (d_428_v43_)[index57_] = len(d_433_v45_)
        index58_ = default__.safeIndex(2, (d_428_v43_).length(0))
        index59_ = default__.safeIndex(14, (d_428_v43_).length(0))
        rhs51_ = 262
        rhs52_ = self.f14
        rhs53_ = self.f14
        rhs54_ = d_428_v43_
        rhs55_ = (d_432_v44_) + (_dafny.SeqWithoutIsStrInference([self.f14 for d_434_i11_ in range(default__.abs(281))]))
        lhs46_ = d_428_v43_
        lhs47_ = default__.safeIndex(2, (d_428_v43_).length(0))
        lhs48_ = self
        lhs49_ = d_428_v43_
        lhs50_ = default__.safeIndex(14, (d_428_v43_).length(0))
        lhs46_[lhs47_] = rhs51_
        lhs48_.f14 = rhs52_
        lhs49_[lhs50_] = rhs53_
        d_428_v43_ = rhs54_
        d_432_v44_ = rhs55_
        d_435_v46_: _dafny.MultiSet
        d_435_v46_ = _dafny.MultiSet([(self).f15, (self).f15, (self).f15, (self).f15])
        d_436_v47_: D0
        d_436_v47_ = D0_DC2(self.f14, d_435_v46_)
        pat_let_tv13_ = d_428_v43_
        pat_let_tv14_ = d_428_v43_
        def lambda28_(source8_):
            if source8_.is_DC1:
                d_437___mcc_h8_ = source8_.cf1
                d_438_cf1_ = d_437___mcc_h8_
                return D3_DC9(d_438_cf1_)
            elif source8_.is_DC2:
                d_439___mcc_h9_ = source8_.cf2
                d_440___mcc_h10_ = source8_.cf3
                d_441_cf3_ = d_440___mcc_h10_
                d_442_cf2_ = d_439___mcc_h9_
                return D3_DC9(d_442_cf2_)
            elif True:
                d_443___mcc_h11_ = source8_.cf0
                d_444_cf0_ = d_443___mcc_h11_
                return D3_DC9((pat_let_tv14_)[default__.safeIndex(2, (pat_let_tv13_).length(0))])

        source7_ = lambda28_(d_436_v47_)
        if source7_.is_DC9:
            d_445___mcc_h0_ = source7_.cf16
            d_446_cf16_ = d_445___mcc_h0_
            d_447_v48_: _dafny.Set
            d_447_v48_ = _dafny.Set({(self).f15})
            d_448_v49_: _dafny.MultiSet
            d_448_v49_ = _dafny.MultiSet([d_447_v48_, d_447_v48_, d_447_v48_])
            d_449_v50_: _dafny.Map
            d_449_v50_ = _dafny.Map({(self).f15: d_448_v49_})
            (self).f14 = len(default__.fm11(((d_449_v50_)[(d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]] if ((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]) in (d_449_v50_) else d_448_v49_), globalState))
            d_450_v51_: _dafny.MultiSet
            d_450_v51_ = _dafny.MultiSet([170])
            d_451_v52_: _dafny.Map
            d_451_v52_ = _dafny.Map({False: (self).f15})
            d_452_v53_: _dafny.Map
            d_452_v53_ = _dafny.Map({d_450_v51_: d_451_v52_})
            d_452_v53_ = (d_452_v53_).set((d_450_v51_).set(self.f14, default__.abs((0) - (((d_435_v46_)[(self).f15] if ((self).f15) in (d_435_v46_) else self.f14)))), (d_451_v52_) | (d_451_v52_))
            d_453_v54_: D1
            d_453_v54_ = D1_DC3(_dafny.SeqWithoutIsStrInference([d_376_v3_ for d_454_i12_ in range(default__.abs(746))]))
            index60_ = default__.safeIndex(2, (d_428_v43_).length(0))
            rhs56_ = (d_446_cf16_) + (d_446_cf16_)
            rhs57_ = ((d_447_v48_) - (d_447_v48_)) - ((self).fm5(d_453_v54_, globalState))
            lhs51_ = d_428_v43_
            lhs52_ = default__.safeIndex(2, (d_428_v43_).length(0))
            lhs51_[lhs52_] = rhs56_
            d_447_v48_ = rhs57_
            d_455_v55_: _dafny.Seq
            d_455_v55_ = _dafny.SeqWithoutIsStrInference([(self).f15, True, (self).f15, (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))], (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]])
            d_456_v56_: _dafny.Map
            d_456_v56_ = _dafny.Map({(d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))]: _dafny.MultiSet(d_455_v55_)})
            d_457_v57_: D2
            d_457_v57_ = D2_DC6(d_456_v56_)
            d_458_v58_: _dafny.Map
            d_458_v58_ = _dafny.Map({d_457_v57_: ((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]) and ((self).f15)})
            d_459_v59_: C0
            nw77_ = C0()
            nw77_.ctor__((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))], (self).f15, (d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))], 670, (self).f15)
            d_459_v59_ = nw77_
            d_460_v60_: _dafny.Map
            d_460_v60_ = _dafny.Map({(d_459_v59_ if (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))] else d_459_v59_): d_458_v58_})
            d_458_v58_ = ((d_460_v60_)[d_459_v59_] if (d_459_v59_) in (d_460_v60_) else d_458_v58_)
        elif source7_.is_DC10:
            d_461___mcc_h1_ = source7_.cf17
            d_462___mcc_h2_ = source7_.cf18
            d_463___mcc_h3_ = source7_.cf19
            d_464___mcc_h4_ = source7_.cf20
            d_465___mcc_h5_ = source7_.cf21
            d_466_cf21_ = d_465___mcc_h5_
            d_467_cf20_ = d_464___mcc_h4_
            d_468_cf19_ = d_463___mcc_h3_
            d_469_cf18_ = d_462___mcc_h2_
            d_470_cf17_ = d_461___mcc_h1_
            d_375_v2_ = (d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))]
            (globalState).f3 = (self).f15
            d_471_v61_: _dafny.Set
            d_471_v61_ = _dafny.Set({d_469_cf18_})
            d_472_v62_: _dafny.Map
            d_472_v62_ = _dafny.Map({(d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))]: (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]})
            d_473_v63_: C0
            nw78_ = C0()
            nw78_.ctor__((d_471_v61_).isdisjoint(d_471_v61_), (self).f15, (d_375_v2_) + (d_470_cf17_), d_466_cf21_, ((d_472_v62_)[d_466_cf21_] if (d_466_cf21_) in (d_472_v62_) else not(True)))
            d_473_v63_ = nw78_
            d_474_v64_: _dafny.Map
            d_474_v64_ = _dafny.Map({d_468_cf19_: self.f14})
            d_475_v65_: D5
            d_475_v65_ = D5_DC13(d_474_v64_)
            index61_ = default__.safeIndex(514, (d_372_v0_).length(0))
            rhs58_ = (d_466_cf21_) - (((d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))] if (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))] else d_467_cf20_))
            rhs59_ = True
            rhs60_ = (d_473_v63_).fm2((d_375_v2_) + (d_375_v2_), (d_473_v63_).f17, (d_474_v64_) != ((d_475_v65_).cf24), (0) - (self.f14), globalState)
            rhs61_ = d_473_v63_
            rhs62_ = d_467_cf20_
            lhs53_ = d_372_v0_
            lhs54_ = default__.safeIndex(514, (d_372_v0_).length(0))
            lhs55_ = self
            r3 = rhs58_
            lhs53_[lhs54_] = rhs59_
            r3 = rhs60_
            d_473_v63_ = rhs61_
            lhs55_.f14 = rhs62_
            d_466_cf21_ = len(d_470_cf17_)
        elif source7_.is_DC8:
            d_476___mcc_h6_ = source7_.cf15
            d_477_cf15_ = d_476___mcc_h6_
            d_478_v66_: _dafny.Map
            d_478_v66_ = _dafny.Map({self.f14: len(d_433_v45_)})
            d_478_v66_ = (d_478_v66_).set(self.f14, 855)
            d_433_v45_ = ((d_433_v45_) | (d_433_v45_) if (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))] else default__.fm12(self.f14, (self).f15, (0) - ((d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))]), d_376_v3_, globalState))
            d_479_v67_: _dafny.Map
            d_479_v67_ = _dafny.Map({935: (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]})
            d_480_v68_: C0
            nw79_ = C0()
            nw79_.ctor__(((d_479_v67_)[239] if (239) in (d_479_v67_) else not((self).f15)), ((self).f15 if (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))] else (self).f15), d_375_v2_, (0) - ((d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))]), (self).f15)
            d_480_v68_ = nw79_
            (globalState).f11 = (d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))]
        elif True:
            d_481___mcc_h7_ = source7_.cf22
            d_482_cf22_ = d_481___mcc_h7_
            d_483_v69_: _dafny.Seq
            d_483_v69_ = _dafny.SeqWithoutIsStrInference([(self).fm0((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))], globalState)])
            d_484_v70_: _dafny.Map
            d_484_v70_ = _dafny.Map({self.f14: d_483_v69_})
            (globalState).f11 = ((d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))]) + (len(((d_484_v70_)[self.f14] if (self.f14) in (d_484_v70_) else _dafny.SeqWithoutIsStrInference([(d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]]))))
            d_485_v71_: D0
            d_485_v71_ = D0_DC0((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))])
            d_485_v71_ = d_485_v71_
            d_486_v72_: _dafny.Map
            d_486_v72_ = _dafny.Map({(self).f15: (self).f15})
            d_486_v72_ = ((d_486_v72_).set((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))], not((d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]))) | (d_486_v72_)
            d_487_v73_: _dafny.Array
            nw80_ = _dafny.Array(None, 20)
            nw80_[int(0)] = (d_376_v3_ if (self).f15 else (d_375_v2_)[default__.safeIndex(self.f14, len(d_375_v2_))])
            nw80_[int(1)] = _dafny.CodePoint('h')
            nw80_[int(2)] = _dafny.CodePoint('y')
            nw80_[int(3)] = d_376_v3_
            nw80_[int(4)] = d_376_v3_
            nw80_[int(5)] = ((d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))])[default__.safeIndex(self.f14, len((d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))]))]
            nw80_[int(6)] = d_376_v3_
            nw80_[int(7)] = d_376_v3_
            nw80_[int(8)] = d_376_v3_
            nw80_[int(9)] = (d_375_v2_)[default__.safeIndex((d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))], len(d_375_v2_))]
            nw80_[int(10)] = d_376_v3_
            nw80_[int(11)] = d_376_v3_
            nw80_[int(12)] = d_376_v3_
            nw80_[int(13)] = d_376_v3_
            nw80_[int(14)] = _dafny.CodePoint('w')
            nw80_[int(15)] = d_376_v3_
            nw80_[int(16)] = d_376_v3_
            nw80_[int(17)] = d_376_v3_
            nw80_[int(18)] = d_376_v3_
            nw80_[int(19)] = d_376_v3_
            d_487_v73_ = nw80_
            index62_ = default__.safeIndex(119, (d_487_v73_).length(0))
            (d_487_v73_)[index62_] = d_376_v3_
            index63_ = default__.safeIndex(119, (d_487_v73_).length(0))
            (d_487_v73_)[index63_] = d_376_v3_
        r0 = (d_425_v41_)[default__.safeIndex(self.f14, len(d_425_v41_))]
        d_488_v74_: _dafny.Map
        d_488_v74_ = _dafny.Map({(d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))]: (d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))]})
        d_489_v75_: _dafny.Map
        d_489_v75_ = _dafny.Map({(d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))]: (self).f15})
        r1 = D1_DC5((d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))], len(_dafny.Set({(d_373_v1_)[default__.safeIndex(90, (d_373_v1_).length(0))]})), (d_372_v0_)[default__.safeIndex(514, (d_372_v0_).length(0))], (d_488_v74_).set(len(d_489_v75_), _dafny.SeqWithoutIsStrInference([d_376_v3_, d_376_v3_, _dafny.CodePoint('d')])))
        r2 = ((d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))]) != ((d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))])
        r3 = (d_428_v43_)[default__.safeIndex(2, (d_428_v43_).length(0))]
        return r0, r1, r2, r3

    def m2(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_490_v0_: _dafny.Array
        nw81_ = _dafny.Array(_dafny.Seq({}), 14)
        d_490_v0_ = nw81_
        d_491_v1_: D5
        d_491_v1_ = D5_DC14(p3, d_490_v0_, (self).f15)
        d_492_v2_: _dafny.Seq
        d_492_v2_ = _dafny.SeqWithoutIsStrInference([(self).f15])
        d_493_v3_: C0
        nw82_ = C0()
        nw82_.ctor__((self).f15, p1, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_494_i0_ in range(default__.abs(448))]), len(d_492_v2_), p1)
        d_493_v3_ = nw82_
        d_495_v4_: _dafny.MultiSet
        d_495_v4_ = _dafny.MultiSet([d_493_v3_, d_493_v3_])
        d_496_v5_: _dafny.Map
        d_496_v5_ = _dafny.Map({p1: d_495_v4_})
        d_497_v6_: _dafny.Seq
        d_497_v6_ = _dafny.SeqWithoutIsStrInference([d_496_v5_])
        d_498_v7_: _dafny.Seq
        d_498_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejencq"))
        d_499_v8_: T0
        nw83_ = C1()
        nw83_.ctor__(d_498_v7_, d_492_v2_, 335, (d_493_v3_).f18)
        d_499_v8_ = nw83_
        d_500_v9_: _dafny.Set
        d_500_v9_ = _dafny.Set({d_499_v8_, d_499_v8_})
        d_501_v10_: _dafny.Array
        nw84_ = _dafny.Array(None, 16)
        nw84_[int(0)] = (self).f15
        nw84_[int(1)] = True
        nw84_[int(2)] = (p1 if (self).f15 else not((d_491_v1_).cf27))
        nw84_[int(3)] = (self).f15
        nw84_[int(4)] = (self).f15
        nw84_[int(5)] = (self).f15
        nw84_[int(6)] = (len((d_497_v6_)[default__.safeIndex((0) - (p2), len(d_497_v6_))])) == (p3)
        nw84_[int(7)] = (p2) <= ((0) - (p3))
        nw84_[int(8)] = (d_493_v3_).f17
        nw84_[int(9)] = (d_492_v2_)[default__.safeIndex(p2, len(d_492_v2_))]
        nw84_[int(10)] = not((d_493_v3_).f18)
        nw84_[int(11)] = (self).f15
        nw84_[int(12)] = (d_500_v9_).issubset(_dafny.Set({d_499_v8_, d_499_v8_, d_499_v8_, d_499_v8_}))
        nw84_[int(13)] = (d_493_v3_).f17
        nw84_[int(14)] = p1
        nw84_[int(15)] = p1
        d_501_v10_ = nw84_
        index64_ = default__.safeIndex(920, (d_501_v10_).length(0))
        (d_501_v10_)[index64_] = (d_499_v8_).f15
        d_502_v11_: _dafny.Array
        nw85_ = _dafny.Array(_dafny.Array(None, 0), 17)
        d_502_v11_ = nw85_
        d_503_v12_: _dafny.MultiSet
        d_503_v12_ = _dafny.MultiSet([self.f14, p2])
        d_504_v13_: _dafny.Map
        d_504_v13_ = _dafny.Map({d_502_v11_: (d_503_v12_).cardinality})
        index65_ = default__.safeIndex(920, (d_501_v10_).length(0))
        rhs63_ = p3
        rhs64_ = ((d_504_v13_)[d_502_v11_] if (d_502_v11_) in (d_504_v13_) else p2)
        rhs65_ = False
        lhs56_ = d_501_v10_
        lhs57_ = default__.safeIndex(920, (d_501_v10_).length(0))
        r0 = rhs63_
        r1 = rhs64_
        lhs56_[lhs57_] = rhs65_
        hi2_ = len(d_498_v7_)
        for d_505_i1_ in range(len(default__.fm13(p3, globalState)), hi2_):
            d_506_v14_: _dafny.Array
            nw86_ = _dafny.Array(_dafny.Map({}), 27)
            d_506_v14_ = nw86_
            d_507_v15_: _dafny.Array
            nw87_ = _dafny.Array(None, 8)
            nw87_[int(0)] = d_506_v14_
            nw87_[int(1)] = d_506_v14_
            nw87_[int(2)] = d_506_v14_
            nw87_[int(3)] = d_506_v14_
            nw87_[int(4)] = d_506_v14_
            nw87_[int(5)] = d_506_v14_
            nw87_[int(6)] = d_506_v14_
            nw87_[int(7)] = d_506_v14_
            d_507_v15_ = nw87_
            index66_ = default__.safeIndex(551, (d_507_v15_).length(0))
            (d_507_v15_)[index66_] = d_506_v14_
            index67_ = default__.safeIndex(551, (d_507_v15_).length(0))
            (d_507_v15_)[index67_] = d_506_v14_
            d_508_v16_: _dafny.Set
            d_508_v16_ = _dafny.Set({(d_493_v3_).f18, (d_493_v3_).f18})
            d_509_v17_: C0
            nw88_ = C0()
            nw88_.ctor__(p1, p1, d_498_v7_, (d_503_v12_).cardinality, (_dafny.Set({(d_493_v3_).f18})).issubset(d_508_v16_))
            d_509_v17_ = nw88_
            d_498_v7_ = d_498_v7_
            (globalState).f0 = p1
        (globalState).f11 = p2
        d_510_v18_: D3
        d_510_v18_ = D3_DC11(D3_DC10(d_498_v7_, d_501_v10_, (d_493_v3_).f18, 502, d_499_v8_.f14))
        d_511_v19_: _dafny.Set
        d_511_v19_ = _dafny.Set({d_510_v18_})
        d_512_v20_: _dafny.Map
        d_512_v20_ = _dafny.Map({d_511_v19_: (0) - (p2)})
        d_513_v21_: _dafny.Map
        d_513_v21_ = _dafny.Map({(p2) * ((0) - (p2)): (d_512_v20_) != (d_512_v20_)})
        d_514_v22_: _dafny.Map
        d_514_v22_ = _dafny.Map({False: -673})
        d_515_v23_: _dafny.MultiSet
        d_515_v23_ = _dafny.MultiSet([_dafny.Map({(d_501_v10_)[default__.safeIndex(920, (d_501_v10_).length(0))]: d_499_v8_.f14}), d_514_v22_])
        d_513_v21_ = (d_513_v21_).set((d_515_v23_).cardinality, (d_493_v3_).f17)
        d_516_v24_: D0
        d_516_v24_ = D0_DC1(len(d_513_v21_))
        source9_ = d_516_v24_
        if source9_.is_DC1:
            d_517___mcc_h0_ = source9_.cf1
            d_518_cf1_ = d_517___mcc_h0_
            d_519_v25_: _dafny.Set
            d_519_v25_ = _dafny.Set({(self).f15})
            d_520_v26_: _dafny.Map
            d_520_v26_ = _dafny.Map({d_519_v25_: d_499_v8_.f14})
            rhs66_ = (d_518_cf1_) >= (self.f14)
            rhs67_ = (d_501_v10_)[default__.safeIndex(920, (d_501_v10_).length(0))]
            rhs68_ = (p3) + (p2)
            rhs69_ = (d_520_v26_) | (d_520_v26_)
            lhs58_ = globalState
            lhs59_ = globalState
            lhs60_ = globalState
            lhs58_.f0 = rhs66_
            lhs59_.f0 = rhs67_
            lhs60_.f11 = rhs68_
            d_520_v26_ = rhs69_
            index68_ = default__.safeIndex(920, (d_501_v10_).length(0))
            (d_501_v10_)[index68_] = p1
            d_514_v22_ = (d_514_v22_).set(not((d_493_v3_).f17), p3)
            d_521_v27_: C0
            nw89_ = C0()
            nw89_.ctor__(((d_513_v21_)[d_518_cf1_] if (d_518_cf1_) in (d_513_v21_) else (d_499_v8_).f15), False, d_498_v7_, len(d_498_v7_), not(not((d_493_v3_).f17)))
            d_521_v27_ = nw89_
        elif source9_.is_DC2:
            d_522___mcc_h1_ = source9_.cf2
            d_523___mcc_h2_ = source9_.cf3
            d_524_cf3_ = d_523___mcc_h2_
            d_525_cf2_ = d_522___mcc_h1_
            r1 = -780
            if (d_525_cf2_) != ((219) - (d_499_v8_.f14)):
                (globalState).f11 = ((len(d_492_v2_)) - (self.f14)) * (p3)
                d_526_v28_: _dafny.Array
                nw90_ = _dafny.Array(None, 15)
                nw90_[int(0)] = d_498_v7_
                nw90_[int(1)] = d_498_v7_
                nw90_[int(2)] = d_498_v7_
                nw90_[int(3)] = d_498_v7_
                nw90_[int(4)] = d_498_v7_
                nw90_[int(5)] = d_498_v7_
                nw90_[int(6)] = d_498_v7_
                nw90_[int(7)] = d_498_v7_
                nw90_[int(8)] = d_498_v7_
                nw90_[int(9)] = (d_498_v7_) + (default__.fm9(p2, globalState))
                nw90_[int(10)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwdgami"))) + (d_498_v7_)
                nw90_[int(11)] = d_498_v7_
                nw90_[int(12)] = d_498_v7_
                nw90_[int(13)] = d_498_v7_
                nw90_[int(14)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gamojfa")) if False else d_498_v7_)
                d_526_v28_ = nw90_
                index69_ = default__.safeIndex(437, (d_526_v28_).length(0))
                (d_526_v28_)[index69_] = (d_498_v7_ if (d_493_v3_).f17 else d_498_v7_)
                d_527_v30_: _dafny.Array
                def lambda29_(d_528_v3_, d_529_v8_, d_530_p3_):
                    def lambda30_(d_531_i2_):
                        def iife23_():
                            coll13_ = _dafny.Set()
                            compr_13_: D2
                            for compr_13_ in (_dafny.MultiSet([D2_DC7((d_528_v3_).f17, (d_529_v8_).f15, d_530_p3_), (self).f19])).Elements:
                                d_532_v29_: D2 = compr_13_
                                if (d_532_v29_) in (_dafny.MultiSet([D2_DC7((d_528_v3_).f17, (d_529_v8_).f15, d_530_p3_), (self).f19])):
                                    coll13_ = coll13_.union(_dafny.Set([d_532_v29_]))
                            return _dafny.Set(coll13_)
                        return (_dafny.Set({iife23_()
                        })) - (_dafny.Set({_dafny.Set({(self).f19})}))

                    return lambda30_

                init16_ = lambda29_(d_493_v3_, d_499_v8_, p3)
                nw91_ = _dafny.Array(None, 17)
                for i0_16_ in range(nw91_.length(0)):
                    nw91_[i0_16_] = init16_(i0_16_)
                d_527_v30_ = nw91_
                d_533_v31_: _dafny.Set
                d_533_v31_ = _dafny.Set({_dafny.Set({D2_DC7(p1, (d_493_v3_).f17, (0) - (d_499_v8_.f14)), (self).f19})})
                d_534_v32_: D6
                d_534_v32_ = D6_DC16(d_533_v31_)
                index70_ = default__.safeIndex(569, (d_527_v30_).length(0))
                (d_527_v30_)[index70_] = (d_534_v32_).cf29
                d_535_v33_: _dafny.Map
                d_535_v33_ = _dafny.Map({(0) - (d_499_v8_.f14): d_501_v10_})
                index71_ = default__.safeIndex(437, (d_526_v28_).length(0))
                index72_ = default__.safeIndex(569, (d_527_v30_).length(0))
                rhs70_ = len(d_535_v33_)
                rhs71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kykbqtx"))
                rhs72_ = (d_533_v31_) | (d_533_v31_)
                rhs73_ = d_503_v12_
                lhs61_ = globalState
                lhs62_ = d_526_v28_
                lhs63_ = default__.safeIndex(437, (d_526_v28_).length(0))
                lhs64_ = d_527_v30_
                lhs65_ = default__.safeIndex(569, (d_527_v30_).length(0))
                lhs61_.f11 = rhs70_
                lhs62_[lhs63_] = rhs71_
                lhs64_[lhs65_] = rhs72_
                d_503_v12_ = rhs73_
                d_536_v34_: _dafny.Array
                nw92_ = _dafny.Array(_dafny.Set({}), 23)
                d_536_v34_ = nw92_
                d_537_v35_: _dafny.MultiSet
                d_537_v35_ = _dafny.MultiSet([d_536_v34_, d_536_v34_])
                d_538_v36_: _dafny.Map
                d_538_v36_ = _dafny.Map({(0) - (p2): d_499_v8_.f14})
                d_539_v37_: _dafny.Map
                d_539_v37_ = _dafny.Map({p2: d_538_v36_})
                (d_499_v8_).f14 = ((d_537_v35_)[d_536_v34_] if (d_536_v34_) in (d_537_v35_) else default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([(d_493_v3_).f18, False])), len(d_539_v37_)))
                d_540_v38_: _dafny.Set
                d_540_v38_ = _dafny.Set({(0) - (d_499_v8_.f14)})
                (globalState).f0 = (d_540_v38_).ispropersubset((d_540_v38_ if (d_493_v3_).f18 else d_540_v38_))
                d_541_v39_: D1
                d_541_v39_ = D1_DC3(d_498_v7_)
                index73_ = default__.safeIndex(437, (d_526_v28_).length(0))
                (d_526_v28_)[index73_] = ((d_526_v28_)[default__.safeIndex(437, (d_526_v28_).length(0))]) + ((d_541_v39_).cf4)
            elif True:
                d_542_v40_: C1
                nw93_ = C1()
                nw93_.ctor__((d_498_v7_) + (d_498_v7_), d_492_v2_, self.f14, False)
                d_542_v40_ = nw93_
                d_542_v40_ = d_542_v40_
                d_543_v41_: _dafny.Array
                nw94_ = _dafny.Array(int(0), 13)
                d_543_v41_ = nw94_
                index74_ = default__.safeIndex(176, (d_543_v41_).length(0))
                (d_543_v41_)[index74_] = ((d_524_cf3_)[(d_499_v8_).f15] if ((d_499_v8_).f15) in (d_524_cf3_) else p3)
                index75_ = default__.safeIndex(446, (d_543_v41_).length(0))
                (d_543_v41_)[index75_] = (0) - (self.f14)
                d_544_v42_: _dafny.Map
                d_544_v42_ = _dafny.Map({d_499_v8_.f14: D5_DC13((d_514_v22_).set((d_499_v8_).f15, d_499_v8_.f14))})
                index76_ = default__.safeIndex(920, (d_501_v10_).length(0))
                index77_ = default__.safeIndex(176, (d_543_v41_).length(0))
                index78_ = default__.safeIndex(446, (d_543_v41_).length(0))
                index79_ = default__.safeIndex(920, (d_501_v10_).length(0))
                rhs74_ = (65) != (len(d_544_v42_))
                rhs75_ = default__.safeModuloInt((self.f14) + (p2), default__.safeModuloInt(p3, 109))
                rhs76_ = d_499_v8_.f14
                rhs77_ = not ((d_493_v3_).f18) or ((d_493_v3_).f18)
                rhs78_ = (d_525_cf2_) - (((d_493_v3_).fm2((d_542_v40_).f20, (d_493_v3_).f18, (d_493_v3_).f18, (0) - ((0) - (d_499_v8_.f14)), globalState) if (d_493_v3_).f18 else (0) - (len(d_513_v21_))))
                lhs66_ = d_501_v10_
                lhs67_ = default__.safeIndex(920, (d_501_v10_).length(0))
                lhs68_ = d_543_v41_
                lhs69_ = default__.safeIndex(176, (d_543_v41_).length(0))
                lhs70_ = d_543_v41_
                lhs71_ = default__.safeIndex(446, (d_543_v41_).length(0))
                lhs72_ = d_501_v10_
                lhs73_ = default__.safeIndex(920, (d_501_v10_).length(0))
                lhs74_ = globalState
                lhs66_[lhs67_] = rhs74_
                lhs68_[lhs69_] = rhs75_
                lhs70_[lhs71_] = rhs76_
                lhs72_[lhs73_] = rhs77_
                lhs74_.f11 = rhs78_
                d_545_v43_: str
                d_545_v43_ = _dafny.CodePoint('m')
                d_546_v44_: _dafny.Map
                d_546_v44_ = _dafny.Map({d_499_v8_.f14: (d_542_v40_).f20})
                d_547_v45_: _dafny.Array
                nw95_ = _dafny.Array(None, 29)
                nw95_[int(0)] = (d_542_v40_).f20
                nw95_[int(1)] = (d_542_v40_).f20
                nw95_[int(2)] = d_498_v7_
                nw95_[int(3)] = (d_542_v40_).f20
                nw95_[int(4)] = (d_542_v40_).f20
                nw95_[int(5)] = _dafny.SeqWithoutIsStrInference([d_545_v43_ for d_548_i3_ in range(default__.abs(518))])
                nw95_[int(6)] = (d_542_v40_).f20
                nw95_[int(7)] = (d_542_v40_).f20
                nw95_[int(8)] = d_498_v7_
                nw95_[int(9)] = (d_542_v40_).f20
                nw95_[int(10)] = d_498_v7_
                nw95_[int(11)] = d_498_v7_
                nw95_[int(12)] = (d_542_v40_).f20
                nw95_[int(13)] = (d_542_v40_).f20
                nw95_[int(14)] = (d_542_v40_).fm8((d_501_v10_)[default__.safeIndex(920, (d_501_v10_).length(0))], globalState)
                nw95_[int(15)] = default__.fm9(d_499_v8_.f14, globalState)
                nw95_[int(16)] = d_498_v7_
                nw95_[int(17)] = (d_542_v40_).f20
                nw95_[int(18)] = d_498_v7_
                nw95_[int(19)] = d_498_v7_
                nw95_[int(20)] = d_498_v7_
                nw95_[int(21)] = d_498_v7_
                nw95_[int(22)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gb"))
                nw95_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
                nw95_[int(24)] = (d_542_v40_).f20
                nw95_[int(25)] = d_498_v7_
                nw95_[int(26)] = (d_542_v40_).f20
                nw95_[int(27)] = (d_498_v7_).set(default__.safeIndex(len(d_546_v44_), len(d_498_v7_)), d_545_v43_)
                nw95_[int(28)] = (d_542_v40_).f20
                d_547_v45_ = nw95_
                d_549_v46_: _dafny.Map
                d_549_v46_ = _dafny.Map({d_545_v43_: d_547_v45_})
                d_549_v46_ = (d_549_v46_).set(d_545_v43_, d_547_v45_)
                (globalState).f3 = ((0) - (p3)) >= (default__.safeDivisionInt(d_499_v8_.f14, (0) - (len(d_514_v22_))))
                d_550_v47_: D3
                d_550_v47_ = D3_DC9((self.f14) + ((d_543_v41_)[default__.safeIndex(176, (d_543_v41_).length(0))]))
                d_550_v47_ = d_550_v47_
            def lambda31_(d_551_v10_):
                def lambda32_(d_552_i4_):
                    return (d_551_v10_)[default__.safeIndex(920, (d_551_v10_).length(0))]

                return lambda32_

            init17_ = lambda31_(d_501_v10_)
            nw96_ = _dafny.Array(None, 10)
            for i0_17_ in range(nw96_.length(0)):
                nw96_[i0_17_] = init17_(i0_17_)
            d_501_v10_ = nw96_
            d_553_v48_: _dafny.Array
            nw97_ = _dafny.Array(int(0), 24)
            d_553_v48_ = nw97_
            index80_ = default__.safeIndex(205, (d_502_v11_).length(0))
            (d_502_v11_)[index80_] = d_553_v48_
            d_554_v49_: str
            d_554_v49_ = _dafny.CodePoint('r')
            d_555_v50_: _dafny.Array
            d_555_v50_ = d_553_v48_
            index81_ = default__.safeIndex(205, (d_502_v11_).length(0))
            rhs79_ = len(d_498_v7_)
            rhs80_ = ((d_555_v50_) if (887) != (d_525_cf2_) else d_553_v48_)
            rhs81_ = (d_493_v3_).f18
            rhs82_ = (self).fm4((d_492_v2_) == (d_492_v2_), globalState)
            rhs83_ = (p3) > (d_499_v8_.f14)
            lhs75_ = d_502_v11_
            lhs76_ = default__.safeIndex(205, (d_502_v11_).length(0))
            lhs77_ = globalState
            lhs78_ = globalState
            r0 = rhs79_
            lhs75_[lhs76_] = rhs80_
            lhs77_.f0 = rhs81_
            d_554_v49_ = rhs82_
            lhs78_.f0 = rhs83_
        elif True:
            d_556___mcc_h3_ = source9_.cf0
            d_557_cf0_ = d_556___mcc_h3_
            index82_ = default__.safeIndex(920, (d_501_v10_).length(0))
            (d_501_v10_)[index82_] = d_557_cf0_
            d_558_v51_: _dafny.Set
            d_558_v51_ = _dafny.Set({(d_493_v3_).f18})
            index83_ = default__.safeIndex(920, (d_501_v10_).length(0))
            index84_ = default__.safeIndex(920, (d_501_v10_).length(0))
            rhs84_ = (d_558_v51_ if (self).f15 else d_558_v51_)
            rhs85_ = ((self).f15) and ((d_499_v8_).f15)
            rhs86_ = p1
            lhs79_ = d_501_v10_
            lhs80_ = default__.safeIndex(920, (d_501_v10_).length(0))
            lhs81_ = d_501_v10_
            lhs82_ = default__.safeIndex(920, (d_501_v10_).length(0))
            d_558_v51_ = rhs84_
            lhs79_[lhs80_] = rhs85_
            lhs81_[lhs82_] = rhs86_
            d_559_v52_: T1
            nw98_ = C0()
            nw98_.ctor__((d_499_v8_).f15, (self).f15, d_498_v7_, d_499_v8_.f14, (d_493_v3_).f18)
            d_559_v52_ = nw98_
            d_560_v53_: D6
            d_560_v53_ = D6_DC17((d_499_v8_).f15, d_559_v52_)
            d_560_v53_ = d_560_v53_
            d_561_v54_: str
            d_561_v54_ = _dafny.CodePoint('i')
            d_562_v55_: C1
            nw99_ = C1()
            nw99_.ctor__((d_498_v7_).set(default__.safeIndex(d_499_v8_.f14, len(d_498_v7_)), d_561_v54_), d_492_v2_, len(d_514_v22_), (d_493_v3_).f18)
            d_562_v55_ = nw99_
            d_563_v56_: _dafny.Set
            d_563_v56_ = _dafny.Set({d_559_v52_})
            rhs87_ = d_490_v0_
            rhs88_ = len(d_563_v56_)
            rhs89_ = len((d_562_v55_).f21)
            rhs90_ = ((d_562_v55_).f21 if (d_501_v10_)[default__.safeIndex(920, (d_501_v10_).length(0))] else (d_562_v55_).f21)
            rhs91_ = d_562_v55_
            lhs83_ = globalState
            lhs84_ = globalState
            d_490_v0_ = rhs87_
            lhs83_.f11 = rhs88_
            lhs84_.f11 = rhs89_
            d_492_v2_ = rhs90_
            d_562_v55_ = rhs91_
        d_564_v57_: _dafny.Seq
        d_564_v57_ = _dafny.SeqWithoutIsStrInference([d_499_v8_.f14])
        d_565_v58_: _dafny.Seq
        d_565_v58_ = _dafny.SeqWithoutIsStrInference([d_564_v57_, d_564_v57_, d_564_v57_])
        d_566_v59_: _dafny.Set
        d_566_v59_ = _dafny.Set({(0) - ((self).fm1(d_503_v12_, (self).f15, globalState))})
        d_567_v60_: _dafny.Map
        d_567_v60_ = _dafny.Map({((d_565_v58_)[default__.safeIndex(d_499_v8_.f14, len(d_565_v58_))]) <= (((d_564_v57_).set(default__.safeIndex(p2, len(d_564_v57_)), d_499_v8_.f14)).set(default__.safeIndex(self.f14, len((d_564_v57_).set(default__.safeIndex(p2, len(d_564_v57_)), d_499_v8_.f14))), p3)): (len(d_566_v59_)) > (-239)})
        if ((d_567_v60_)[not(False)] if (not(False)) in (d_567_v60_) else (d_499_v8_).f15):
            d_568_v61_: _dafny.Map
            d_568_v61_ = _dafny.Map({d_513_v21_: d_501_v10_})
            default__.m0(d_568_v61_, globalState)
            d_569_v62_: _dafny.Array
            def lambda33_(d_570_v8_):
                def lambda34_(d_571_i5_):
                    return (d_571_i5_) * (d_570_v8_.f14)

                return lambda34_

            init18_ = lambda33_(d_499_v8_)
            nw100_ = _dafny.Array(None, 16)
            for i0_18_ in range(nw100_.length(0)):
                nw100_[i0_18_] = init18_(i0_18_)
            d_569_v62_ = nw100_
            d_572_v63_: _dafny.Array
            d_572_v63_ = d_569_v62_
            d_572_v63_ = (d_572_v63_ if (d_493_v3_).f18 else d_572_v63_)
            (globalState).f0 = not((d_493_v3_).f18)
            rhs92_ = d_499_v8_.f14
            rhs93_ = p2
            lhs85_ = d_499_v8_
            lhs86_ = globalState
            lhs85_.f14 = rhs92_
            lhs86_.f11 = rhs93_
            d_573_v64_: _dafny.Set
            d_573_v64_ = _dafny.Set({(d_493_v3_).f18, p1, (d_493_v3_).f18})
            (globalState).f0 = (d_573_v64_).ispropersubset(d_573_v64_)
        elif True:
            if not((d_493_v3_).f17):
                d_574_v65_: _dafny.Array
                def lambda35_(d_575_v3_, d_576_v8_, d_577_v2_):
                    def lambda36_(d_578_i6_):
                        return (d_578_i6_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_575_v3_).f17, (d_575_v3_).f17, (d_576_v8_).f15, (d_575_v3_).f18, (d_576_v8_).f15]), d_577_v2_])))

                    return lambda36_

                init19_ = lambda35_(d_493_v3_, d_499_v8_, d_492_v2_)
                nw101_ = _dafny.Array(None, 4)
                for i0_19_ in range(nw101_.length(0)):
                    nw101_[i0_19_] = init19_(i0_19_)
                d_574_v65_ = nw101_
                index85_ = default__.safeIndex(983, (d_574_v65_).length(0))
                (d_574_v65_)[index85_] = d_499_v8_.f14
                d_579_v66_: _dafny.Seq
                d_579_v66_ = _dafny.SeqWithoutIsStrInference([d_501_v10_])
                d_580_v67_: C1
                nw102_ = C1()
                nw102_.ctor__(d_498_v7_, d_492_v2_, (0) - (p3), (self).f15)
                d_580_v67_ = nw102_
                d_581_v68_: _dafny.Map
                d_581_v68_ = _dafny.Map({d_580_v67_: (d_501_v10_)[default__.safeIndex(920, (d_501_v10_).length(0))]})
                d_582_v69_: _dafny.Seq
                d_582_v69_ = _dafny.SeqWithoutIsStrInference([d_581_v68_, d_581_v68_, d_581_v68_, d_581_v68_, d_581_v68_])
                d_583_v70_: _dafny.Set
                d_583_v70_ = _dafny.Set({(d_493_v3_).f18})
                index86_ = default__.safeIndex(983, (d_574_v65_).length(0))
                rhs94_ = (d_499_v8_.f14) - (d_499_v8_.f14)
                rhs95_ = len(((d_579_v66_) + (_dafny.SeqWithoutIsStrInference([d_501_v10_, d_501_v10_, d_501_v10_]))) + (d_579_v66_))
                rhs96_ = (0) - ((0) - ((0) - (default__.safeDivisionInt((0) - (len((d_582_v69_)[default__.safeIndex(len(d_583_v70_), len(d_582_v69_))])), default__.safeDivisionInt(self.f14, self.f14)))))
                rhs97_ = False
                lhs87_ = d_574_v65_
                lhs88_ = default__.safeIndex(983, (d_574_v65_).length(0))
                lhs89_ = d_499_v8_
                lhs90_ = globalState
                lhs87_[lhs88_] = rhs94_
                r0 = rhs95_
                lhs89_.f14 = rhs96_
                lhs90_.f0 = rhs97_
                d_584_v71_: _dafny.Map
                d_584_v71_ = _dafny.Map({(809) + (p2): d_513_v21_})
                d_585_v73_: _dafny.Map
                d_585_v73_ = _dafny.Map({(0) - (d_499_v8_.f14): (d_574_v65_)[default__.safeIndex(983, (d_574_v65_).length(0))]})
                index87_ = default__.safeIndex(920, (d_501_v10_).length(0))
                def iife24_():
                    coll14_ = _dafny.Set()
                    compr_14_: int
                    for compr_14_ in (d_566_v59_).Elements:
                        d_586_v72_: int = compr_14_
                        if (d_586_v72_) in (d_566_v59_):
                            coll14_ = coll14_.union(_dafny.Set([default__.safeDivisionInt(d_586_v72_, len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_587_i7_ in range(default__.abs(253))]))})))]))
                    return _dafny.Set(coll14_)
                def iife25_():
                    coll15_ = _dafny.Set()
                    compr_15_: int
                    for compr_15_ in (d_566_v59_).Elements:
                        d_588_v72_: int = compr_15_
                        if (d_588_v72_) in (d_566_v59_):
                            coll15_ = coll15_.union(_dafny.Set([default__.safeDivisionInt(d_588_v72_, len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_589_i7_ in range(default__.abs(253))]))})))]))
                    return _dafny.Set(coll15_)
                rhs98_ = ((d_584_v71_)[len(iife24_()
                )] if (len(iife25_()
                )) in (d_584_v71_) else d_513_v21_)
                rhs99_ = (len((d_585_v73_) | (d_585_v73_))) >= (19)
                lhs91_ = d_501_v10_
                lhs92_ = default__.safeIndex(920, (d_501_v10_).length(0))
                d_513_v21_ = rhs98_
                lhs91_[lhs92_] = rhs99_
                rhs100_ = d_499_v8_.f14
                rhs101_ = d_566_v59_
                rhs102_ = (d_493_v3_).f17
                lhs93_ = globalState
                lhs94_ = globalState
                lhs93_.f11 = rhs100_
                d_566_v59_ = rhs101_
                lhs94_.f3 = rhs102_
                r1 = default__.safeDivisionInt(p2, self.f14)
                d_590_v74_: str
                d_590_v74_ = _dafny.CodePoint('d')
                d_591_v75_: _dafny.Map
                d_591_v75_ = _dafny.Map({681: d_590_v74_})
                r1 = len(default__.fm14((default__.fm15(d_590_v74_, p3, len(d_591_v75_), globalState)).set((d_501_v10_)[default__.safeIndex(920, (d_501_v10_).length(0))], p2), _dafny.CodePoint('c'), not(((d_580_v67_).f20) != ((d_498_v7_).set(default__.safeIndex((d_580_v67_).fm1(d_503_v12_, True, globalState), len(d_498_v7_)), d_590_v74_))), globalState))
            elif True:
                d_592_v76_: _dafny.Map
                d_592_v76_ = _dafny.Map({(d_498_v7_) != (d_498_v7_): d_564_v57_})
                d_592_v76_ = d_592_v76_
                d_593_v77_: _dafny.Seq
                d_593_v77_ = _dafny.SeqWithoutIsStrInference([d_492_v2_])
                d_594_v78_: C1
                nw103_ = C1()
                nw103_.ctor__(d_498_v7_, (d_593_v77_)[default__.safeIndex((d_516_v24_).cf1, len(d_593_v77_))], 169, True)
                d_594_v78_ = nw103_
                d_595_v79_: str
                d_595_v79_ = _dafny.CodePoint('h')
                d_596_v80_: _dafny.Map
                d_596_v80_ = _dafny.Map({(d_493_v3_).f18: d_595_v79_})
                d_597_v81_: D6
                d_597_v81_ = D6_DC18(d_499_v8_.f14, ((d_596_v80_)[(d_493_v3_).f18] if ((d_493_v3_).f18) in (d_596_v80_) else d_595_v79_))
                d_598_v82_: D6
                d_598_v82_ = D6_DC19(d_597_v81_)
                d_599_v83_: D6
                d_599_v83_ = D6_DC19(d_598_v82_)
                d_600_v84_: _dafny.Map
                d_600_v84_ = _dafny.Map({d_594_v78_: d_599_v83_})
                d_600_v84_ = (d_600_v84_).set((d_594_v78_ if (d_499_v8_).f15 else d_594_v78_), d_599_v83_)
                d_601_v85_: C0
                nw104_ = C0()
                nw104_.ctor__((d_493_v3_).f17, (d_493_v3_).f18, (d_594_v78_).f20, self.f14, (d_493_v3_).f17)
                d_601_v85_ = nw104_
                d_602_v86_: _dafny.MultiSet
                d_602_v86_ = _dafny.MultiSet([(d_493_v3_).f18, (d_601_v85_).fm0((d_493_v3_).f18, globalState), (d_499_v8_).f15, (d_493_v3_).f18, (self).f15])
                d_603_v87_: _dafny.Map
                d_603_v87_ = _dafny.Map({d_499_v8_.f14: d_602_v86_})
                d_604_v88_: D2
                d_604_v88_ = D2_DC6(d_603_v87_)
                d_604_v88_ = d_604_v88_
                d_605_v89_: _dafny.Map
                d_605_v89_ = _dafny.Map({default__.fm16(len(d_566_v59_), globalState): d_501_v10_})
                default__.m0(d_605_v89_, globalState)
            (globalState).f0 = (d_492_v2_)[default__.safeIndex((0) - (((d_503_v12_)[len(d_498_v7_)] if (len(d_498_v7_)) in (d_503_v12_) else d_499_v8_.f14)), len(d_492_v2_))]
            r1 = (d_499_v8_.f14) - (self.f14)
            (globalState).f0 = (d_493_v3_).fm0((d_501_v10_)[default__.safeIndex(920, (d_501_v10_).length(0))], globalState)
            d_606_v90_: _dafny.Seq
            d_606_v90_ = _dafny.SeqWithoutIsStrInference([d_492_v2_, d_492_v2_, d_492_v2_, _dafny.SeqWithoutIsStrInference([(d_493_v3_).f18]), _dafny.SeqWithoutIsStrInference([(d_499_v8_).f15, True])])
            d_607_v91_: C1
            nw105_ = C1()
            nw105_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqvurq")), (d_606_v90_)[default__.safeIndex(d_499_v8_.f14, len(d_606_v90_))], d_499_v8_.f14, (d_493_v3_).f18)
            d_607_v91_ = nw105_
        r0 = default__.safeModuloInt((0) - ((0) - ((d_499_v8_).fm1(d_503_v12_, (self).f15, globalState))), p3)
        d_608_v92_: _dafny.Seq
        d_608_v92_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p3, p2})])
        r1 = default__.safeModuloInt(d_499_v8_.f14, (len(d_608_v92_)) - (-105))
        return r0, r1

    @property
    def f19(self):
        return self._f19
