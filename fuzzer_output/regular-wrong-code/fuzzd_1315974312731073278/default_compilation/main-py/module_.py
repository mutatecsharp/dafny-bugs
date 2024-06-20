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
    def fm0(p0, p1, p2, p3, globalState):
        return (len(_dafny.Map({-469: False}))) * (default__.safeDivisionInt(-411, len(_dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([366 for d_0_i0_ in range(default__.abs(94))]))).cardinality}))))

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(850, 668):
                d_1_v0_: int = compr_0_
                if ((850) <= (d_1_v0_)) and ((d_1_v0_) < (668)):
                    def iife1_():
                        coll1_ = _dafny.Set()
                        compr_1_: str
                        for compr_1_ in (_dafny.Map({_dafny.CodePoint('b'): False})).keys.Elements:
                            d_2_v1_: str = compr_1_
                            if (d_2_v1_) in (_dafny.Map({_dafny.CodePoint('b'): False})):
                                coll1_ = coll1_.union(_dafny.Set([d_2_v1_]))
                        return _dafny.Set(coll1_)
                    coll0_[default__.safeModuloInt(d_1_v0_, -868)] = len(iife1_()
                    )
            return _dafny.Map(coll0_)
        return ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference([iife0_()
        ]))})) | (_dafny.Set({176}))) | ((_dafny.Set({166})) | (_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "henpihgvp")))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_3_i0_ in range(default__.abs(91))])), 936])))})))

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([D1_DC4(D1_DC3(True, True, False, not(not(not(True))))) for d_4_i0_ in range(default__.abs(76))])

    @staticmethod
    def fm4(p0, globalState):
        return _dafny.MultiSet([(D4_DC11(_dafny.CodePoint('b'), True, _dafny.MultiSet([_dafny.CodePoint('t')]), True, _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([493]) for d_5_i0_ in range(default__.abs(-835))]))).cf20])

    @staticmethod
    def fm5(globalState):
        return _dafny.CodePoint('g')

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        source0_ = D3_DC7()
        if source0_.is_DC7:
            return False
        elif source0_.is_DC8:
            d_6___mcc_h0_ = source0_.cf14
            d_7_cf14_ = d_6___mcc_h0_
            return False
        elif True:
            d_8___mcc_h1_ = source0_.cf13
            d_9_cf13_ = d_8___mcc_h1_
            return not(not(False))

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('x')])

    @staticmethod
    def fm8(p0, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(939, 466):
                d_10_v0_: int = compr_2_
                if ((939) <= (d_10_v0_)) and ((d_10_v0_) < (466)):
                    coll2_[(d_10_v0_) * (len(_dafny.Map({False: False})))] = False
            return _dafny.Map(coll2_)
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, False}))]) if False else _dafny.SeqWithoutIsStrInference([len(iife2_()
        ), 844, 518]))) + (_dafny.SeqWithoutIsStrInference([244 for d_11_i0_ in range(default__.abs(681))]))

    @staticmethod
    def fm9(p0, p1, globalState):
        return D3_DC7()

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(421, 551):
                d_12_v0_: int = compr_3_
                if ((421) <= (d_12_v0_)) and ((d_12_v0_) < (551)):
                    coll3_[(d_12_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxbtxxe"))))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gi")))
            return _dafny.Map(coll3_)
        source1_ = D4_DC11(_dafny.CodePoint('h'), False, _dafny.MultiSet([_dafny.CodePoint('m'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('d'), _dafny.CodePoint('d')]), True, _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(iife3_()
)]) for d_13_i0_ in range(default__.abs(271))]))
        if source1_.is_DC10:
            d_14___mcc_h0_ = source1_.cf16
            d_15___mcc_h1_ = source1_.cf17
            d_16___mcc_h2_ = source1_.cf18
            d_17_cf18_ = d_16___mcc_h2_
            d_18_cf17_ = d_15___mcc_h1_
            d_19_cf16_ = d_14___mcc_h0_
            return _dafny.MultiSet([_dafny.Map({d_18_cf17_: True})])
        elif source1_.is_DC11:
            d_20___mcc_h3_ = source1_.cf19
            d_21___mcc_h4_ = source1_.cf20
            d_22___mcc_h5_ = source1_.cf21
            d_23___mcc_h6_ = source1_.cf22
            d_24___mcc_h7_ = source1_.cf23
            d_25_cf23_ = d_24___mcc_h7_
            d_26_cf22_ = d_23___mcc_h6_
            d_27_cf21_ = d_22___mcc_h5_
            d_28_cf20_ = d_21___mcc_h4_
            d_29_cf19_ = d_20___mcc_h3_
            return (_dafny.MultiSet([_dafny.Map({d_26_cf22_: d_26_cf22_}), _dafny.Map({d_28_cf20_: d_28_cf20_})])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({True: d_26_cf22_}), _dafny.Map({d_28_cf20_: d_26_cf22_})])))
        elif source1_.is_DC12:
            d_30___mcc_h8_ = source1_.cf24
            d_31___mcc_h9_ = source1_.cf25
            d_32___mcc_h10_ = source1_.cf26
            d_33___mcc_h11_ = source1_.cf27
            d_34_cf27_ = d_33___mcc_h11_
            d_35_cf26_ = d_32___mcc_h10_
            d_36_cf25_ = d_31___mcc_h9_
            d_37_cf24_ = d_30___mcc_h8_
            return _dafny.MultiSet([_dafny.Map({d_35_cf26_: d_35_cf26_}), _dafny.Map({d_35_cf26_: (D4_DC10(d_37_cf24_, False, d_35_cf26_)).cf18})])
        elif True:
            d_38___mcc_h12_ = source1_.cf15
            d_39_cf15_ = d_38___mcc_h12_
            return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.Map({False: True})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: False}) for d_40_i1_ in range(default__.abs(632))])))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        if p1:
            (globalState).f15 = default__.safeModuloInt(p0, default__.safeDivisionInt(p0, 639))
            d_41_v0_: _dafny.Seq
            d_41_v0_ = _dafny.SeqWithoutIsStrInference([p0, p0, (0) - (default__.fm0(p1, p3, p2, p0, globalState)), (p0) - (p0)])
            d_42_v1_: str
            d_42_v1_ = _dafny.CodePoint('y')
            d_43_v2_: C0
            nw0_ = C0()
            nw0_.ctor__(p0, d_42_v1_)
            d_43_v2_ = nw0_
            d_44_v3_: _dafny.Seq
            d_44_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "deyue"))
            d_45_v4_: D0
            d_45_v4_ = D0_DC0(d_44_v3_)
            rhs0_ = default__.fm8(not (p2) or (p2), globalState)
            rhs1_ = ((default__.fm7(False, (d_43_v2_).f19, len(d_41_v0_), globalState)) + (d_44_v3_)) == ((d_44_v3_ if not(p2) else (d_45_v4_).cf0))
            rhs2_ = (d_43_v2_).f19
            rhs3_ = d_43_v2_
            rhs4_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkykdidt"))) + (_dafny.SeqWithoutIsStrInference([d_42_v1_ for d_46_i0_ in range(default__.abs(74))]))
            lhs0_ = globalState
            lhs1_ = globalState
            d_41_v0_ = rhs0_
            lhs0_.f6 = rhs1_
            lhs1_.f13 = rhs2_
            d_43_v2_ = rhs3_
            d_44_v3_ = rhs4_
            d_47_v5_: _dafny.Array
            def lambda0_(d_48_v2_):
                def lambda1_(d_49_i1_):
                    return default__.safeModuloInt(d_49_i1_, (d_48_v2_).f19)

                return lambda1_

            init0_ = lambda0_(d_43_v2_)
            nw1_ = _dafny.Array(None, 8)
            for i0_0_ in range(nw1_.length(0)):
                nw1_[i0_0_] = init0_(i0_0_)
            d_47_v5_ = nw1_
            d_50_v6_: _dafny.Map
            d_50_v6_ = _dafny.Map({p2: d_47_v5_})
            d_51_v7_: _dafny.Map
            d_51_v7_ = _dafny.Map({p2: p2})
            d_52_v8_: _dafny.Map
            d_52_v8_ = _dafny.Map({((d_50_v6_)[True] if (True) in (d_50_v6_) else d_47_v5_): d_51_v7_})
            d_52_v8_ = (d_52_v8_).set(d_47_v5_, d_51_v7_)
            d_53_v9_: _dafny.Map
            d_53_v9_ = _dafny.Map({d_41_v0_: d_47_v5_})
            d_53_v9_ = (d_53_v9_) | (d_53_v9_)
            d_54_v10_: _dafny.Array
            nw2_ = _dafny.Array(None, 14)
            nw2_[int(0)] = d_42_v1_
            nw2_[int(1)] = d_42_v1_
            nw2_[int(2)] = (p3)[default__.safeIndex((d_43_v2_).f19, len(p3))]
            nw2_[int(3)] = (d_43_v2_).f20
            nw2_[int(4)] = (d_43_v2_).f20
            nw2_[int(5)] = (d_43_v2_).f20
            nw2_[int(6)] = d_42_v1_
            nw2_[int(7)] = _dafny.CodePoint('c')
            nw2_[int(8)] = d_42_v1_
            nw2_[int(9)] = d_42_v1_
            nw2_[int(10)] = (d_43_v2_).f20
            nw2_[int(11)] = (d_43_v2_).f20
            nw2_[int(12)] = d_42_v1_
            nw2_[int(13)] = (d_43_v2_).f20
            d_54_v10_ = nw2_
            index0_ = default__.safeIndex(614, (d_54_v10_).length(0))
            (d_54_v10_)[index0_] = d_42_v1_
            index1_ = default__.safeIndex(614, (d_54_v10_).length(0))
            (d_54_v10_)[index1_] = (d_43_v2_).f20
        elif True:
            d_55_v11_: str
            d_55_v11_ = _dafny.CodePoint('a')
            (globalState).f6 = ((d_55_v11_ if p2 else default__.fm5(globalState))) in (p3)
            d_56_v12_: D3
            d_56_v12_ = D3_DC7()
            d_56_v12_ = d_56_v12_
            d_57_v13_: _dafny.MultiSet
            d_57_v13_ = _dafny.MultiSet([(not(p1) if p1 else p1), (False) and (p1)])
            d_57_v13_ = default__.fm4(p2, globalState)
            (globalState).f13 = (0) - (p0)
            d_58_v14_: _dafny.Array
            def lambda2_(d_59_p0_):
                def lambda3_(d_60_i2_):
                    return default__.safeModuloInt(d_60_i2_, d_59_p0_)

                return lambda3_

            init1_ = lambda2_(p0)
            nw3_ = _dafny.Array(None, 18)
            for i0_1_ in range(nw3_.length(0)):
                nw3_[i0_1_] = init1_(i0_1_)
            d_58_v14_ = nw3_
            d_61_v15_: _dafny.Array
            nw4_ = _dafny.Array(None, 19)
            nw4_[int(0)] = p2
            nw4_[int(1)] = not(p1)
            nw4_[int(2)] = p2
            nw4_[int(3)] = p2
            nw4_[int(4)] = p1
            nw4_[int(5)] = p1
            nw4_[int(6)] = p1
            nw4_[int(7)] = p2
            nw4_[int(8)] = p2
            nw4_[int(9)] = True
            nw4_[int(10)] = p2
            nw4_[int(11)] = not(p1)
            nw4_[int(12)] = p2
            nw4_[int(13)] = False
            nw4_[int(14)] = p2
            nw4_[int(15)] = p1
            nw4_[int(16)] = p2
            nw4_[int(17)] = p2
            nw4_[int(18)] = p1
            d_61_v15_ = nw4_
            d_62_v16_: _dafny.Seq
            d_62_v16_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_55_v11_ for d_63_i3_ in range(default__.abs(691))]), p3])
            d_64_v17_: D0
            d_64_v17_ = D0_DC1(d_58_v14_, d_61_v15_, _dafny.MultiSet((d_62_v16_).set(default__.safeIndex(p0, len(d_62_v16_)), _dafny.SeqWithoutIsStrInference([d_55_v11_]))), p0, p0)
            index2_ = default__.safeIndex(248, (d_58_v14_).length(0))
            (d_58_v14_)[index2_] = (d_64_v17_).cf4
            d_65_v18_: D1
            d_65_v18_ = D1_DC3(p2, p1, True, default__.fm6(p0, default__.fm5(globalState), (0) - (p0), globalState))
            index3_ = default__.safeIndex(248, (d_58_v14_).length(0))
            (d_58_v14_)[index3_] = (0) - ((_dafny.MultiSet([p1, (d_65_v18_).cf8])).cardinality)
        d_66_v19_: _dafny.Array
        nw5_ = _dafny.Array(int(0), 20)
        d_66_v19_ = nw5_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_66_v19_).length(0)):
            d_67_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_67_i4_)) and ((d_67_i4_) < ((d_66_v19_).length(0)))):
                (d_66_v19_)[(d_67_i4_)] = default__.safeModuloInt(d_67_i4_, p0)
        hi0_ = p0
        for d_68_i5_ in range(p0, hi0_):
            d_69_v20_: _dafny.Array
            nw6_ = _dafny.Array(False, 13)
            d_69_v20_ = nw6_
            d_70_v21_: _dafny.Set
            d_70_v21_ = _dafny.Set({d_69_v20_})
            d_70_v21_ = d_70_v21_
            d_71_v22_: _dafny.Array
            nw7_ = _dafny.Array(_dafny.Map({}), 3)
            d_71_v22_ = nw7_
            d_71_v22_ = d_71_v22_
            d_72_v23_: _dafny.Map
            d_72_v23_ = _dafny.Map({p1: p1})
            source2_ = default__.fm9(True, len((d_72_v23_) | (d_72_v23_)), globalState)
            if source2_.is_DC7:
                (globalState).f6 = True
                d_73_v24_: _dafny.Seq
                d_73_v24_ = _dafny.SeqWithoutIsStrInference([p1])
                d_74_v25_: _dafny.Seq
                d_74_v25_ = _dafny.SeqWithoutIsStrInference([((p3).set(default__.safeIndex(p0, len(p3)), _dafny.CodePoint('i'))).set(default__.safeIndex(d_68_i5_, len((p3).set(default__.safeIndex(p0, len(p3)), _dafny.CodePoint('i')))), default__.fm5(globalState)), p3, (p3) + (p3), default__.fm7(not((d_73_v24_)[default__.safeIndex(d_68_i5_, len(d_73_v24_))]), p0, p0, globalState)])
                d_74_v25_ = (d_74_v25_) + (d_74_v25_)
                d_75_v26_: _dafny.Map
                d_75_v26_ = _dafny.Map({p3: p0})
                d_76_v27_: _dafny.Map
                d_76_v27_ = _dafny.Map({(p2) or (p1): d_75_v26_})
                d_76_v27_ = (d_76_v27_).set(p1, (d_75_v26_) | (_dafny.Map({p3: p0})))
                d_77_v28_: str
                d_77_v28_ = _dafny.CodePoint('w')
                d_78_v29_: C0
                nw8_ = C0()
                nw8_.ctor__(p0, d_77_v28_)
                d_78_v29_ = nw8_
                d_79_v30_: D4
                d_79_v30_ = D4_DC9(d_78_v29_)
                d_80_v31_: _dafny.Array
                nw9_ = _dafny.Array(None, 19)
                nw9_[int(0)] = d_78_v29_
                nw9_[int(1)] = d_78_v29_
                nw9_[int(2)] = d_78_v29_
                nw9_[int(3)] = d_78_v29_
                nw9_[int(4)] = (d_79_v30_).cf15
                nw9_[int(5)] = d_78_v29_
                nw9_[int(6)] = d_78_v29_
                nw9_[int(7)] = d_78_v29_
                nw9_[int(8)] = d_78_v29_
                nw9_[int(9)] = d_78_v29_
                nw9_[int(10)] = d_78_v29_
                nw9_[int(11)] = d_78_v29_
                nw9_[int(12)] = d_78_v29_
                nw9_[int(13)] = d_78_v29_
                nw9_[int(14)] = d_78_v29_
                nw9_[int(15)] = d_78_v29_
                nw9_[int(16)] = d_78_v29_
                nw9_[int(17)] = d_78_v29_
                nw9_[int(18)] = d_78_v29_
                d_80_v31_ = nw9_
                d_81_v32_: _dafny.Set
                d_81_v32_ = _dafny.Set({d_80_v31_, d_80_v31_, d_80_v31_, d_80_v31_, d_80_v31_})
                d_81_v32_ = _dafny.Set({(d_80_v31_ if p2 else d_80_v31_)})
            elif source2_.is_DC8:
                d_82___mcc_h0_ = source2_.cf14
                d_83_cf14_ = d_82___mcc_h0_
                d_84_v33_: str
                d_84_v33_ = _dafny.CodePoint('c')
                d_85_v34_: _dafny.Map
                d_85_v34_ = _dafny.Map({d_84_v33_: p1})
                d_86_v35_: _dafny.Map
                d_86_v35_ = _dafny.Map({d_69_v20_: d_85_v34_})
                d_87_v36_: _dafny.Seq
                d_87_v36_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({d_69_v20_: d_85_v34_}) if p1 else d_86_v35_)])
                d_88_v37_: C0
                nw10_ = C0()
                nw10_.ctor__(d_83_cf14_, d_84_v33_)
                d_88_v37_ = nw10_
                d_89_v38_: _dafny.Map
                d_89_v38_ = _dafny.Map({d_88_v37_: d_68_i5_})
                rhs5_ = (d_87_v36_)[default__.safeIndex(d_83_cf14_, len(d_87_v36_))]
                rhs6_ = (len((d_89_v38_).set(d_88_v37_, d_83_cf14_))) < (908)
                d_86_v35_ = rhs5_
                r1 = rhs6_
                d_90_v39_: _dafny.Map
                d_90_v39_ = _dafny.Map({470: len((D0_DC0(p3)).cf0)})
                r1 = (((d_88_v37_).f19) > (len(d_90_v39_)) if (p1) == (True) else p2)
                d_69_v20_ = d_69_v20_
                (globalState).f15 = d_83_cf14_
            elif True:
                d_91___mcc_h1_ = source2_.cf13
                d_92_cf13_ = d_91___mcc_h1_
                d_93_v40_: str
                d_93_v40_ = _dafny.CodePoint('t')
                d_94_v41_: C0
                nw11_ = C0()
                nw11_.ctor__(p0, d_93_v40_)
                d_94_v41_ = nw11_
                d_95_v42_: _dafny.Seq
                d_95_v42_ = _dafny.SeqWithoutIsStrInference([p0, len(p3), d_68_i5_])
                d_96_v43_: _dafny.Map
                d_96_v43_ = _dafny.Map({d_95_v42_: d_94_v41_})
                d_97_v44_: _dafny.MultiSet
                d_97_v44_ = _dafny.MultiSet([(d_94_v41_).f20, (d_94_v41_).f20])
                d_98_v45_: _dafny.MultiSet
                d_98_v45_ = _dafny.MultiSet([len(_dafny.Set({not(p2)})), 328])
                d_99_v46_: _dafny.Seq
                d_99_v46_ = _dafny.SeqWithoutIsStrInference([d_98_v45_, d_98_v45_])
                d_100_v47_: D4
                d_100_v47_ = D4_DC11((d_94_v41_).f20, p1, d_97_v44_, p1, d_99_v46_)
                rhs7_ = (d_100_v47_).cf22
                rhs8_ = _dafny.Map({d_95_v42_: d_94_v41_})
                r1 = rhs7_
                d_96_v43_ = rhs8_
                (globalState).f5 = (0) - ((((d_94_v41_).f19) * ((d_94_v41_).f19)) + (d_68_i5_))
                d_95_v42_ = ((d_95_v42_) + (d_95_v42_)) + ((d_95_v42_ if p2 else _dafny.SeqWithoutIsStrInference([(d_94_v41_).f19])))
            d_101_v48_: _dafny.Map
            d_101_v48_ = _dafny.Map({p0: p1})
            d_102_v49_: str
            d_102_v49_ = _dafny.CodePoint('c')
            index4_ = default__.safeIndex(518, (d_66_v19_).length(0))
            (d_66_v19_)[index4_] = default__.safeDivisionInt(d_68_i5_, default__.fm0(p1, p3, default__.fm6(len(d_101_v48_), d_102_v49_, len(_dafny.SeqWithoutIsStrInference([p2])), globalState), p0, globalState))
            index5_ = default__.safeIndex(518, (d_66_v19_).length(0))
            (d_66_v19_)[index5_] = (0) - (p0)
        d_103_v50_: _dafny.Array
        def lambda4_(d_104_p3_):
            def lambda5_(d_105_i6_):
                return d_104_p3_

            return lambda5_

        init2_ = lambda4_(p3)
        nw12_ = _dafny.Array(None, 25)
        for i0_2_ in range(nw12_.length(0)):
            nw12_[i0_2_] = init2_(i0_2_)
        d_103_v50_ = nw12_
        index6_ = default__.safeIndex(296, (d_103_v50_).length(0))
        (d_103_v50_)[index6_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
        index7_ = default__.safeIndex(296, (d_103_v50_).length(0))
        (d_103_v50_)[index7_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arixobu"))
        d_106_v51_: _dafny.Array
        nw13_ = _dafny.Array(False, 3)
        d_106_v51_ = nw13_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_106_v51_).length(0)):
            d_107_i7_: int = guard_loop_1_
            if (True) and (((0) <= (d_107_i7_)) and ((d_107_i7_) < ((d_106_v51_).length(0)))):
                (d_106_v51_)[(d_107_i7_)] = p1
        d_108_v52_: str
        d_108_v52_ = _dafny.CodePoint('x')
        d_109_v53_: C0
        nw14_ = C0()
        nw14_.ctor__(p0, d_108_v52_)
        d_109_v53_ = nw14_
        d_109_v53_ = d_109_v53_
        d_110_v54_: _dafny.Set
        d_110_v54_ = _dafny.Set({(d_109_v53_).f19, p0})
        d_111_v56_: _dafny.Set
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(-145, 464):
                d_112_v55_: int = compr_4_
                if ((-145) <= (d_112_v55_)) and ((d_112_v55_) < (464)):
                    coll4_ = coll4_.union(_dafny.Set([default__.safeDivisionInt(d_112_v55_, len(_dafny.SeqWithoutIsStrInference([p2, p1, True, p1, p1])))]))
            return _dafny.Set(coll4_)
        d_111_v56_ = _dafny.Set({d_110_v54_, (d_110_v54_) | (iife4_()
        ), default__.fm1(p1, (d_109_v53_).f19, p0, 503, globalState)})
        r0 = d_111_v56_
        r1 = p1
        d_113_v57_: D4
        d_113_v57_ = D4_DC10(p0, p2, p1)
        d_114_v58_: _dafny.Map
        d_114_v58_ = _dafny.Map({p0: 828})
        d_115_v60_: _dafny.MultiSet
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(335, 915):
                d_116_v59_: int = compr_5_
                if ((335) <= (d_116_v59_)) and ((d_116_v59_) < (915)):
                    coll5_[(d_116_v59_) - (len(d_110_v54_))] = (d_109_v53_).f19
            return _dafny.Map(coll5_)
        d_115_v60_ = _dafny.MultiSet([d_114_v58_, iife5_()
        ])
        r2 = default__.fm10(p0, (435) - ((d_113_v57_).cf16), d_115_v60_, p0, globalState)
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_117_v0_: int
        d_117_v0_ = 327
        d_118_v1_: _dafny.Seq
        d_118_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qljvqw"))
        d_119_v2_: bool
        d_119_v2_ = True
        d_120_v3_: _dafny.Seq
        d_120_v3_ = _dafny.SeqWithoutIsStrInference([d_119_v2_])
        d_121_v4_: _dafny.MultiSet
        d_121_v4_ = _dafny.MultiSet([d_118_v1_, d_118_v1_, d_118_v1_])
        d_122_v5_: _dafny.Array
        def lambda6_(d_123_v2_):
            def lambda7_(d_124_i0_):
                return d_123_v2_

            return lambda7_

        init3_ = lambda6_(d_119_v2_)
        nw15_ = _dafny.Array(None, 29)
        for i0_3_ in range(nw15_.length(0)):
            nw15_[i0_3_] = init3_(i0_3_)
        d_122_v5_ = nw15_
        d_125_globalState_: GlobalState
        nw16_ = GlobalState()
        nw16_.ctor__(624, 743, False, -990, False, 913, False, _dafny.SeqWithoutIsStrInference([d_117_v0_, len(d_118_v1_), d_117_v0_, (0) - (d_117_v0_)]), _dafny.Map({d_120_v3_: d_119_v2_}), 580, 647, _dafny.CodePoint('r'), d_121_v4_, 481, _dafny.SeqWithoutIsStrInference([d_122_v5_, d_122_v5_]), 31, True, -830, False)
        d_125_globalState_ = nw16_
        d_126_v6_: str
        d_126_v6_ = _dafny.CodePoint('m')
        d_127_v7_: _dafny.Seq
        d_127_v7_ = _dafny.SeqWithoutIsStrInference([-349])
        d_128_v8_: _dafny.Set
        d_128_v8_ = _dafny.Set({d_117_v0_, (d_127_v7_)[default__.safeIndex(-857, len(d_127_v7_))], d_117_v0_, d_117_v0_, d_117_v0_})
        d_129_v9_: _dafny.MultiSet
        d_129_v9_ = _dafny.MultiSet([default__.fm1(d_119_v2_, (0) - (len(_dafny.SeqWithoutIsStrInference([True, d_119_v2_, True, d_119_v2_, d_119_v2_]))), d_117_v0_, (0) - (d_117_v0_), d_125_globalState_), d_128_v8_])
        d_130_v11_: _dafny.Map
        d_130_v11_ = _dafny.Map({d_117_v0_: (0) - (d_117_v0_)})
        d_131_v12_: _dafny.Array
        nw17_ = _dafny.Array(None, 14)
        nw17_[int(0)] = len(_dafny.Map({d_117_v0_: d_119_v2_}))
        nw17_[int(1)] = default__.safeModuloInt(d_117_v0_, d_117_v0_)
        nw17_[int(2)] = d_117_v0_
        nw17_[int(3)] = -498
        nw17_[int(4)] = len((d_118_v1_).set(default__.safeIndex(default__.fm0(d_119_v2_, d_118_v1_, d_119_v2_, len(_dafny.Map({d_119_v2_: d_126_v6_})), d_125_globalState_), len(d_118_v1_)), _dafny.CodePoint('q')))
        nw17_[int(5)] = d_117_v0_
        nw17_[int(6)] = d_117_v0_
        def iife6_():
            coll6_ = _dafny.Set()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(218, 451):
                d_132_v10_: int = compr_6_
                if ((218) <= (d_132_v10_)) and ((d_132_v10_) < (451)):
                    coll6_ = coll6_.union(_dafny.Set([default__.safeDivisionInt(d_132_v10_, (0) - (len(d_128_v8_)))]))
            return _dafny.Set(coll6_)
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(218, 451):
                d_133_v10_: int = compr_7_
                if ((218) <= (d_133_v10_)) and ((d_133_v10_) < (451)):
                    coll7_ = coll7_.union(_dafny.Set([default__.safeDivisionInt(d_133_v10_, (0) - (len(d_128_v8_)))]))
            return _dafny.Set(coll7_)
        nw17_[int(7)] = ((d_129_v9_)[iife6_()
        ] if (iife7_()
        ) in (d_129_v9_) else 157)
        nw17_[int(8)] = d_117_v0_
        nw17_[int(9)] = (d_117_v0_) - (d_117_v0_)
        nw17_[int(10)] = d_117_v0_
        nw17_[int(11)] = default__.safeDivisionInt(d_117_v0_, d_117_v0_)
        nw17_[int(12)] = default__.safeModuloInt((0) - (d_117_v0_), d_117_v0_)
        nw17_[int(13)] = (0) - ((len(d_130_v11_) if d_119_v2_ else (0) - (d_117_v0_)))
        d_131_v12_ = nw17_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_131_v12_).length(0)):
            d_134_i1_: int = guard_loop_2_
            if (True) and (((0) <= (d_134_i1_)) and ((d_134_i1_) < ((d_131_v12_).length(0)))):
                (d_131_v12_)[(d_134_i1_)] = default__.safeDivisionInt(d_134_i1_, d_117_v0_)
        d_135_v13_: _dafny.MultiSet
        d_135_v13_ = _dafny.MultiSet([d_117_v0_, d_117_v0_])
        rhs9_ = d_126_v6_
        rhs10_ = d_117_v0_
        rhs11_ = (d_135_v13_).issubset(d_135_v13_)
        rhs12_ = d_131_v12_
        lhs2_ = d_125_globalState_
        lhs3_ = d_125_globalState_
        d_126_v6_ = rhs9_
        lhs2_.f5 = rhs10_
        lhs3_.f6 = rhs11_
        d_131_v12_ = rhs12_
        d_136_v14_: _dafny.Set
        d_137_v15_: bool
        d_138_v16_: _dafny.MultiSet
        out0_: _dafny.Set
        out1_: bool
        out2_: _dafny.MultiSet
        out0_, out1_, out2_ = default__.m0(d_117_v0_, False, d_119_v2_, d_118_v1_, d_125_globalState_)
        d_136_v14_ = out0_
        d_137_v15_ = out1_
        d_138_v16_ = out2_
        (d_125_globalState_).f5 = d_117_v0_
        index8_ = default__.safeIndex(514, (d_122_v5_).length(0))
        (d_122_v5_)[index8_] = not(d_137_v15_)
        index9_ = default__.safeIndex(514, (d_122_v5_).length(0))
        (d_122_v5_)[index9_] = d_137_v15_
        d_139_v17_: _dafny.Array
        def lambda8_(d_140_v1_):
            def lambda9_(d_141_i2_):
                return d_140_v1_

            return lambda9_

        init4_ = lambda8_(d_118_v1_)
        nw18_ = _dafny.Array(None, 10)
        for i0_4_ in range(nw18_.length(0)):
            nw18_[i0_4_] = init4_(i0_4_)
        d_139_v17_ = nw18_
        index10_ = default__.safeIndex(564, (d_139_v17_).length(0))
        (d_139_v17_)[index10_] = d_118_v1_
        index11_ = default__.safeIndex(564, (d_139_v17_).length(0))
        (d_139_v17_)[index11_] = (D0_DC0(d_118_v1_)).cf0
        d_142_i3_: int
        d_142_i3_ = 0
        with _dafny.label("0"):
            while not(d_137_v15_):
                with _dafny.c_label("0"):
                    if (d_142_i3_) >= (100):
                        raise _dafny.Break("0")
                    d_142_i3_ = (d_142_i3_) + (1)
                    d_143_v18_: D1
                    d_143_v18_ = D1_DC2(d_126_v6_)
                    d_144_v19_: C0
                    nw19_ = C0()
                    nw19_.ctor__(d_117_v0_, (d_143_v18_).cf6)
                    d_144_v19_ = nw19_
                    index12_ = default__.safeIndex(686, (d_131_v12_).length(0))
                    (d_131_v12_)[index12_] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofwu")))) + (len(d_118_v1_))
                    d_145_v20_: _dafny.Map
                    d_145_v20_ = _dafny.Map({d_119_v2_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oijvun")))})
                    index13_ = default__.safeIndex(686, (d_131_v12_).length(0))
                    (d_131_v12_)[index13_] = (0) - (len(default__.fm3((((d_145_v20_)[d_119_v2_] if (d_119_v2_) in (d_145_v20_) else -524) if False else (d_144_v19_).f19), ((d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))]) or ((d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))]), d_120_v3_, d_145_v20_, d_125_globalState_)))
                    d_146_v21_: _dafny.Map
                    d_146_v21_ = _dafny.Map({(d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))]: not(d_137_v15_)})
                    source3_ = D1_DC3(d_137_v15_, (True) not in (d_120_v3_), d_137_v15_, ((d_127_v7_).set(default__.safeIndex(default__.fm0(d_137_v15_, (d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))], d_137_v15_, d_117_v0_, d_125_globalState_), len(d_127_v7_)), (d_144_v19_).f19)) <= (_dafny.SeqWithoutIsStrInference([len(d_146_v21_)])))
                    if source3_.is_DC3:
                        d_147___mcc_h0_ = source3_.cf7
                        d_148___mcc_h1_ = source3_.cf8
                        d_149___mcc_h2_ = source3_.cf9
                        d_150___mcc_h3_ = source3_.cf10
                        d_151_cf10_ = d_150___mcc_h3_
                        d_152_cf9_ = d_149___mcc_h2_
                        d_153_cf8_ = d_148___mcc_h1_
                        d_154_cf7_ = d_147___mcc_h0_
                        (d_125_globalState_).f5 = (d_144_v19_).f19
                        d_155_v22_: _dafny.Set
                        d_156_v23_: bool
                        d_157_v24_: _dafny.MultiSet
                        out3_: _dafny.Set
                        out4_: bool
                        out5_: _dafny.MultiSet
                        out3_, out4_, out5_ = default__.m0(default__.safeDivisionInt((d_131_v12_)[default__.safeIndex(686, (d_131_v12_).length(0))], (d_131_v12_)[default__.safeIndex(686, (d_131_v12_).length(0))]), (d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], (d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], ((d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))]) + (d_118_v1_), d_125_globalState_)
                        d_155_v22_ = out3_
                        d_156_v23_ = out4_
                        d_157_v24_ = out5_
                        d_158_v25_: C0
                        nw20_ = C0()
                        nw20_.ctor__(d_117_v0_, d_126_v6_)
                        d_158_v25_ = nw20_
                        d_159_v26_: _dafny.MultiSet
                        d_159_v26_ = _dafny.MultiSet([d_144_v19_])
                        d_160_v27_: _dafny.Array
                        nw21_ = _dafny.Array(None, 27)
                        nw21_[int(0)] = (d_158_v25_).f19
                        nw21_[int(1)] = len(_dafny.Set({d_156_v23_}))
                        nw21_[int(2)] = (d_144_v19_).f19
                        nw21_[int(3)] = (d_131_v12_)[default__.safeIndex(686, (d_131_v12_).length(0))]
                        nw21_[int(4)] = (d_144_v19_).f19
                        nw21_[int(5)] = (d_144_v19_).f19
                        nw21_[int(6)] = (d_144_v19_).f19
                        nw21_[int(7)] = (d_131_v12_)[default__.safeIndex(686, (d_131_v12_).length(0))]
                        nw21_[int(8)] = 179
                        nw21_[int(9)] = d_117_v0_
                        nw21_[int(10)] = d_117_v0_
                        nw21_[int(11)] = (d_144_v19_).f19
                        nw21_[int(12)] = (d_144_v19_).f19
                        nw21_[int(13)] = len(d_146_v21_)
                        nw21_[int(14)] = d_117_v0_
                        nw21_[int(15)] = (d_158_v25_).f19
                        nw21_[int(16)] = len((d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))])
                        nw21_[int(17)] = (d_144_v19_).f19
                        nw21_[int(18)] = (d_144_v19_).f19
                        nw21_[int(19)] = d_117_v0_
                        nw21_[int(20)] = -124
                        nw21_[int(21)] = (d_144_v19_).f19
                        nw21_[int(22)] = (d_144_v19_).f19
                        nw21_[int(23)] = (d_131_v12_)[default__.safeIndex(686, (d_131_v12_).length(0))]
                        nw21_[int(24)] = ((d_135_v13_)[(d_159_v26_).cardinality] if ((d_159_v26_).cardinality) in (d_135_v13_) else (d_144_v19_).f19)
                        nw21_[int(25)] = (d_144_v19_).f19
                        nw21_[int(26)] = -897
                        d_160_v27_ = nw21_
                        d_161_v28_: _dafny.Map
                        d_161_v28_ = _dafny.Map({(d_144_v19_).f20: d_160_v27_})
                        d_162_v29_: _dafny.Set
                        d_163_v30_: bool
                        d_164_v31_: _dafny.MultiSet
                        out6_: _dafny.Set
                        out7_: bool
                        out8_: _dafny.MultiSet
                        out6_, out7_, out8_ = default__.m0(d_117_v0_, (d_161_v28_) != ((d_161_v28_)), d_137_v15_, (d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))], d_125_globalState_)
                        d_162_v29_ = out6_
                        d_163_v30_ = out7_
                        d_164_v31_ = out8_
                    elif source3_.is_DC2:
                        d_165___mcc_h4_ = source3_.cf6
                        d_166_cf6_ = d_165___mcc_h4_
                        d_167_v32_: _dafny.Seq
                        d_167_v32_ = _dafny.SeqWithoutIsStrInference([d_127_v7_])
                        d_168_v33_: _dafny.Map
                        d_168_v33_ = _dafny.Map({(d_167_v32_)[default__.safeIndex((d_131_v12_)[default__.safeIndex(686, (d_131_v12_).length(0))], len(d_167_v32_))]: (d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))]})
                        d_169_v34_: _dafny.Set
                        d_169_v34_ = _dafny.Set({not((d_168_v33_) == (d_168_v33_))})
                        d_170_v35_: _dafny.Map
                        d_170_v35_ = _dafny.Map({d_144_v19_: d_117_v0_})
                        d_171_v36_: _dafny.MultiSet
                        d_171_v36_ = _dafny.MultiSet([(d_144_v19_) in (d_170_v35_), False])
                        d_172_v37_: _dafny.Seq
                        d_172_v37_ = _dafny.SeqWithoutIsStrInference([d_144_v19_])
                        d_173_v38_: _dafny.Seq
                        d_173_v38_ = _dafny.SeqWithoutIsStrInference([d_169_v34_])
                        index14_ = default__.safeIndex(686, (d_131_v12_).length(0))
                        index15_ = default__.safeIndex(686, (d_131_v12_).length(0))
                        rhs13_ = default__.safeModuloInt(len((d_172_v37_).set(default__.safeIndex((d_131_v12_)[default__.safeIndex(686, (d_131_v12_).length(0))], len(d_172_v37_)), d_144_v19_)), d_117_v0_)
                        rhs14_ = (d_173_v38_)[default__.safeIndex(len(d_120_v3_), len(d_173_v38_))]
                        rhs15_ = default__.fm4(d_137_v15_, d_125_globalState_)
                        rhs16_ = (d_144_v19_).f19
                        lhs4_ = d_131_v12_
                        lhs5_ = default__.safeIndex(686, (d_131_v12_).length(0))
                        lhs6_ = d_131_v12_
                        lhs7_ = default__.safeIndex(686, (d_131_v12_).length(0))
                        lhs4_[lhs5_] = rhs13_
                        d_169_v34_ = rhs14_
                        d_171_v36_ = rhs15_
                        lhs6_[lhs7_] = rhs16_
                        d_117_v0_ = d_117_v0_
                        def iife8_():
                            coll8_ = _dafny.Set()
                            compr_8_: int
                            for compr_8_ in (d_127_v7_).Elements:
                                d_174_v39_: int = compr_8_
                                if (d_174_v39_) in (d_127_v7_):
                                    coll8_ = coll8_.union(_dafny.Set([default__.safeDivisionInt(d_174_v39_, -557)]))
                            return _dafny.Set(coll8_)
                        d_128_v8_ = iife8_()
                        
                        d_175_v40_: _dafny.Array
                        nw22_ = _dafny.Array(None, 4)
                        d_175_v40_ = nw22_
                        d_175_v40_ = d_175_v40_
                    elif True:
                        d_176___mcc_h5_ = source3_.cf11
                        d_177_cf11_ = d_176___mcc_h5_
                        (d_125_globalState_).f15 = (d_131_v12_)[default__.safeIndex(686, (d_131_v12_).length(0))]
                        d_178_v41_: _dafny.Array
                        def lambda10_(d_179_v0_):
                            def lambda11_(d_180_i4_):
                                return (d_180_i4_) * (d_179_v0_)

                            return lambda11_

                        init5_ = lambda10_(d_117_v0_)
                        nw23_ = _dafny.Array(None, 23)
                        for i0_5_ in range(nw23_.length(0)):
                            nw23_[i0_5_] = init5_(i0_5_)
                        d_178_v41_ = nw23_
                        d_181_v42_: _dafny.Map
                        d_181_v42_ = _dafny.Map({(d_144_v19_).f20: d_178_v41_})
                        d_182_v43_: _dafny.Map
                        d_182_v43_ = d_181_v42_
                        d_182_v43_ = d_181_v42_
                        d_126_v6_ = (d_144_v19_).f20
                        d_182_v43_ = d_182_v43_
                    d_183_v44_: _dafny.Set
                    d_184_v45_: bool
                    d_185_v46_: _dafny.MultiSet
                    out9_: _dafny.Set
                    out10_: bool
                    out11_: _dafny.MultiSet
                    out9_, out10_, out11_ = default__.m0((d_144_v19_).f19, d_119_v2_, d_137_v15_, ((d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))]) + (d_118_v1_), d_125_globalState_)
                    d_183_v44_ = out9_
                    d_184_v45_ = out10_
                    d_185_v46_ = out11_
                    pass
            pass
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_131_v12_).length(0)):
            d_186_i5_: int = guard_loop_3_
            if (True) and (((0) <= (d_186_i5_)) and ((d_186_i5_) < ((d_131_v12_).length(0)))):
                (d_131_v12_)[(d_186_i5_)] = (d_186_i5_) - ((d_117_v0_ if d_137_v15_ else d_117_v0_))
        d_187_v47_: D1
        d_187_v47_ = D1_DC3((d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], d_119_v2_, d_119_v2_, d_137_v15_)
        (d_125_globalState_).f6 = not((d_187_v47_).cf7)
        hi1_ = len(_dafny.Set({d_117_v0_, d_117_v0_}))
        for d_188_i6_ in range(d_117_v0_, hi1_):
            d_189_v48_: D0
            d_189_v48_ = D0_DC1(d_131_v12_, d_122_v5_, d_121_v4_, d_188_i6_, d_188_i6_)
            source4_ = d_189_v48_
            if source4_.is_DC1:
                d_190___mcc_h6_ = source4_.cf1
                d_191___mcc_h7_ = source4_.cf2
                d_192___mcc_h8_ = source4_.cf3
                d_193___mcc_h9_ = source4_.cf4
                d_194___mcc_h10_ = source4_.cf5
                d_195_cf5_ = d_194___mcc_h10_
                d_196_cf4_ = d_193___mcc_h9_
                d_197_cf3_ = d_192___mcc_h8_
                d_198_cf2_ = d_191___mcc_h7_
                d_199_cf1_ = d_190___mcc_h6_
                d_126_v6_ = (d_118_v1_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_126_v6_ for d_200_i7_ in range(default__.abs(603))])), len(d_118_v1_))]
                d_201_v49_: _dafny.Set
                d_202_v50_: bool
                d_203_v51_: _dafny.MultiSet
                out12_: _dafny.Set
                out13_: bool
                out14_: _dafny.MultiSet
                out12_, out13_, out14_ = default__.m0(d_196_cf4_, True, d_137_v15_, d_118_v1_, d_125_globalState_)
                d_201_v49_ = out12_
                d_202_v50_ = out13_
                d_203_v51_ = out14_
                (d_125_globalState_).f5 = d_195_cf5_
                d_195_cf5_ = d_196_cf4_
            elif True:
                d_204___mcc_h11_ = source4_.cf0
                d_205_cf0_ = d_204___mcc_h11_
                d_120_v3_ = d_120_v3_
                d_206_v52_: _dafny.Array
                def lambda12_(d_207_v1_):
                    def lambda13_(d_208_i8_):
                        return D0_DC0(d_207_v1_)

                    return lambda13_

                init6_ = lambda12_(d_118_v1_)
                nw24_ = _dafny.Array(None, 11)
                for i0_6_ in range(nw24_.length(0)):
                    nw24_[i0_6_] = init6_(i0_6_)
                d_206_v52_ = nw24_
                d_209_v53_: D0
                d_209_v53_ = D0_DC0(d_118_v1_)
                index16_ = default__.safeIndex(106, (d_206_v52_).length(0))
                (d_206_v52_)[index16_] = d_209_v53_
                index17_ = default__.safeIndex(106, (d_206_v52_).length(0))
                rhs17_ = (d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))]
                rhs18_ = d_117_v0_
                rhs19_ = (d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))]
                rhs20_ = not(False)
                rhs21_ = D0_DC0((d_205_cf0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fy"))))
                lhs8_ = d_125_globalState_
                lhs9_ = d_125_globalState_
                lhs10_ = d_125_globalState_
                lhs11_ = d_206_v52_
                lhs12_ = default__.safeIndex(106, (d_206_v52_).length(0))
                d_205_cf0_ = rhs17_
                lhs8_.f13 = rhs18_
                lhs9_.f6 = rhs19_
                lhs10_.f6 = rhs20_
                lhs11_[lhs12_] = rhs21_
                d_210_v54_: _dafny.Map
                d_210_v54_ = _dafny.Map({d_126_v6_: d_131_v12_})
                d_211_v55_: _dafny.Map
                d_211_v55_ = d_210_v54_
                d_212_v56_: _dafny.Map
                d_212_v56_ = _dafny.Map({d_117_v0_: d_211_v55_})
                d_119_v2_ = (d_188_i6_) >= (len((d_212_v56_).set(d_188_i6_, d_211_v55_)))
                d_213_v57_: _dafny.Set
                d_214_v58_: bool
                d_215_v59_: _dafny.MultiSet
                out15_: _dafny.Set
                out16_: bool
                out17_: _dafny.MultiSet
                out15_, out16_, out17_ = default__.m0(849, ((d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))]) == ((d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))]), (d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], d_118_v1_, d_125_globalState_)
                d_213_v57_ = out15_
                d_214_v58_ = out16_
                d_215_v59_ = out17_
            d_216_v60_: _dafny.Array
            def lambda14_(d_217_v7_, d_218_i6_):
                def lambda15_(d_219_i9_):
                    return (d_217_v7_).set(default__.safeIndex(180, len(d_217_v7_)), d_218_i6_)

                return lambda15_

            init7_ = lambda14_(d_127_v7_, d_188_i6_)
            nw25_ = _dafny.Array(None, 5)
            for i0_7_ in range(nw25_.length(0)):
                nw25_[i0_7_] = init7_(i0_7_)
            d_216_v60_ = nw25_
            d_216_v60_ = d_216_v60_
            (d_125_globalState_).f6 = not(not(not(d_137_v15_)))
            d_220_v61_: C0
            nw26_ = C0()
            nw26_.ctor__(204, d_126_v6_)
            d_220_v61_ = nw26_
            d_221_v62_: _dafny.Array
            def lambda16_(d_222_v5_):
                def lambda17_(d_223_i10_):
                    return D1_DC3((d_222_v5_)[default__.safeIndex(514, (d_222_v5_).length(0))], (d_222_v5_)[default__.safeIndex(514, (d_222_v5_).length(0))], True, (d_222_v5_)[default__.safeIndex(514, (d_222_v5_).length(0))])

                return lambda17_

            init8_ = lambda16_(d_122_v5_)
            nw27_ = _dafny.Array(None, 29)
            for i0_8_ in range(nw27_.length(0)):
                nw27_[i0_8_] = init8_(i0_8_)
            d_221_v62_ = nw27_
            d_224_v63_: _dafny.Map
            d_224_v63_ = _dafny.Map({d_220_v61_: d_221_v62_})
            index18_ = default__.safeIndex(514, (d_122_v5_).length(0))
            rhs22_ = default__.safeDivisionInt(default__.safeDivisionInt(d_117_v0_, len(d_224_v63_)), d_188_i6_)
            rhs23_ = d_119_v2_
            rhs24_ = d_117_v0_
            rhs25_ = d_139_v17_
            rhs26_ = 827
            lhs13_ = d_122_v5_
            lhs14_ = default__.safeIndex(514, (d_122_v5_).length(0))
            lhs15_ = d_125_globalState_
            d_117_v0_ = rhs22_
            lhs13_[lhs14_] = rhs23_
            d_117_v0_ = rhs24_
            d_139_v17_ = rhs25_
            lhs15_.f15 = rhs26_
        hi2_ = d_117_v0_
        for d_225_i11_ in range(default__.safeDivisionInt(d_117_v0_, d_117_v0_), hi2_):
            if d_119_v2_:
                d_226_v64_: C0
                nw28_ = C0()
                nw28_.ctor__(d_117_v0_, d_126_v6_)
                d_226_v64_ = nw28_
                d_227_v65_: _dafny.Map
                d_227_v65_ = _dafny.Map({d_119_v2_: d_226_v64_})
                d_228_v66_: _dafny.Array
                nw29_ = _dafny.Array(None, 3)
                nw29_[int(0)] = (_dafny.Map({d_137_v15_: d_226_v64_})) | (d_227_v65_)
                nw29_[int(1)] = d_227_v65_
                nw29_[int(2)] = (d_227_v65_) | (_dafny.Map({not(d_137_v15_): d_226_v64_}))
                d_228_v66_ = nw29_
                d_228_v66_ = d_228_v66_
                d_229_v67_: _dafny.Array
                def lambda18_(d_230_v3_, d_231_v15_):
                    def lambda19_(d_232_i12_):
                        return (_dafny.MultiSet(d_230_v3_)) - (_dafny.MultiSet([d_231_v15_]))

                    return lambda19_

                init9_ = lambda18_(d_120_v3_, d_137_v15_)
                nw30_ = _dafny.Array(None, 3)
                for i0_9_ in range(nw30_.length(0)):
                    nw30_[i0_9_] = init9_(i0_9_)
                d_229_v67_ = nw30_
                d_233_v68_: _dafny.MultiSet
                d_233_v68_ = _dafny.MultiSet([d_137_v15_])
                index19_ = default__.safeIndex(456, (d_229_v67_).length(0))
                (d_229_v67_)[index19_] = (d_233_v68_) - (_dafny.MultiSet([(d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], (d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], d_137_v15_, d_119_v2_]))
                index20_ = default__.safeIndex(456, (d_229_v67_).length(0))
                (d_229_v67_)[index20_] = d_233_v68_
                index21_ = default__.safeIndex(807, (d_131_v12_).length(0))
                (d_131_v12_)[index21_] = (d_226_v64_).f19
                d_234_v69_: _dafny.Map
                d_234_v69_ = _dafny.Map({(d_233_v68_).cardinality: d_137_v15_})
                index22_ = default__.safeIndex(807, (d_131_v12_).length(0))
                rhs27_ = default__.safeModuloInt(len(d_234_v69_), d_225_i11_)
                rhs28_ = d_119_v2_
                lhs16_ = d_131_v12_
                lhs17_ = default__.safeIndex(807, (d_131_v12_).length(0))
                lhs18_ = d_125_globalState_
                lhs16_[lhs17_] = rhs27_
                lhs18_.f6 = rhs28_
                d_235_v70_: _dafny.Seq
                d_235_v70_ = _dafny.SeqWithoutIsStrInference([d_122_v5_])
                d_235_v70_ = (d_235_v70_).set(default__.safeIndex((d_131_v12_)[default__.safeIndex(807, (d_131_v12_).length(0))], len(d_235_v70_)), d_122_v5_)
                (d_125_globalState_).f5 = default__.safeDivisionInt(d_225_i11_, d_225_i11_)
            elif True:
                (d_125_globalState_).f6 = (d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))]
                d_236_v71_: D0
                d_236_v71_ = D0_DC1(d_131_v12_, d_122_v5_, d_121_v4_, d_117_v0_, (d_127_v7_)[default__.safeIndex(len(d_118_v1_), len(d_127_v7_))])
                d_237_v72_: _dafny.Seq
                d_237_v72_ = _dafny.SeqWithoutIsStrInference([d_236_v71_])
                index23_ = default__.safeIndex(564, (d_139_v17_).length(0))
                rhs29_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vw"))).set(default__.safeIndex(d_225_i11_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vw")))), d_126_v6_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmpe")))
                rhs30_ = not (d_137_v15_) or (d_137_v15_)
                rhs31_ = ((d_237_v72_) + (d_237_v72_)) + (d_237_v72_)
                rhs32_ = default__.fm5(d_125_globalState_)
                lhs19_ = d_139_v17_
                lhs20_ = default__.safeIndex(564, (d_139_v17_).length(0))
                lhs21_ = d_125_globalState_
                lhs19_[lhs20_] = rhs29_
                lhs21_.f6 = rhs30_
                d_237_v72_ = rhs31_
                d_126_v6_ = rhs32_
                (d_125_globalState_).f15 = d_225_i11_
                (d_125_globalState_).f6 = d_119_v2_
                d_238_v73_: _dafny.Map
                d_238_v73_ = _dafny.Map({d_236_v71_: ((d_135_v13_).cardinality) + (d_225_i11_)})
                d_238_v73_ = (d_238_v73_).set(d_236_v71_, (0) - (d_117_v0_))
            d_239_v74_: _dafny.Seq
            d_239_v74_ = _dafny.SeqWithoutIsStrInference([d_118_v1_])
            d_239_v74_ = d_239_v74_
            d_240_v75_: _dafny.Set
            d_240_v75_ = _dafny.Set({True})
            d_241_v76_: _dafny.Set
            d_241_v76_ = _dafny.Set({d_240_v75_})
            d_242_v77_: _dafny.Map
            d_242_v77_ = _dafny.Map({d_126_v6_: len(d_241_v76_)})
            d_242_v77_ = (d_242_v77_).set(d_126_v6_, (0) - (d_117_v0_))
            (d_125_globalState_).f15 = (d_117_v0_) * (d_225_i11_)
        d_243_i13_: int
        d_243_i13_ = 0
        with _dafny.label("1"):
            while not(d_137_v15_):
                with _dafny.c_label("1"):
                    if (d_243_i13_) >= (100):
                        raise _dafny.Break("1")
                    d_243_i13_ = (d_243_i13_) + (1)
                    d_244_v78_: _dafny.Set
                    d_245_v79_: bool
                    d_246_v80_: _dafny.MultiSet
                    out18_: _dafny.Set
                    out19_: bool
                    out20_: _dafny.MultiSet
                    out18_, out19_, out20_ = default__.m0(((d_130_v11_)[len((d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))])] if (len((d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))])) in (d_130_v11_) else d_117_v0_), not(d_137_v15_), (d_117_v0_) != (d_117_v0_), d_118_v1_, d_125_globalState_)
                    d_244_v78_ = out18_
                    d_245_v79_ = out19_
                    d_246_v80_ = out20_
                    d_247_v81_: C0
                    nw31_ = C0()
                    nw31_.ctor__(d_117_v0_, d_126_v6_)
                    d_247_v81_ = nw31_
                    d_248_v82_: _dafny.Seq
                    d_248_v82_ = _dafny.SeqWithoutIsStrInference([d_128_v8_, default__.fm1((d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vityokux"))), len(d_127_v7_), d_117_v0_, d_125_globalState_)])
                    d_249_v83_: _dafny.Map
                    d_249_v83_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([_dafny.Set({719, (d_247_v81_).f19}), d_128_v8_])) <= (d_248_v82_): not((d_117_v0_) < (d_117_v0_))})
                    d_249_v83_ = (d_249_v83_).set(d_245_v79_, (d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))])
                    index24_ = default__.safeIndex(550, (d_131_v12_).length(0))
                    (d_131_v12_)[index24_] = d_117_v0_
                    index25_ = default__.safeIndex(550, (d_131_v12_).length(0))
                    (d_131_v12_)[index25_] = d_117_v0_
                    pass
            pass
        d_250_v84_: _dafny.Set
        d_250_v84_ = _dafny.Set({d_137_v15_, (d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))]})
        hi3_ = default__.safeModuloInt(d_117_v0_, d_117_v0_)
        for d_251_i14_ in range(len((d_250_v84_) | (d_250_v84_)), hi3_):
            if d_137_v15_:
                d_187_v47_ = d_187_v47_
                (d_125_globalState_).f15 = (0) - ((d_251_i14_) - (881))
                index26_ = default__.safeIndex(514, (d_122_v5_).length(0))
                (d_122_v5_)[index26_] = not(not(d_137_v15_))
                (d_125_globalState_).f13 = d_251_i14_
                d_252_v85_: _dafny.Seq
                d_252_v85_ = _dafny.SeqWithoutIsStrInference([d_122_v5_])
                d_253_v86_: D3
                d_253_v86_ = D3_DC6(d_252_v85_)
                (d_125_globalState_).f5 = len((d_253_v86_).cf13)
            elif True:
                (d_125_globalState_).f6 = d_119_v2_
                d_187_v47_ = d_187_v47_
                (d_125_globalState_).f6 = (-512) < (d_117_v0_)
                index27_ = default__.safeIndex(514, (d_122_v5_).length(0))
                (d_122_v5_)[index27_] = (d_250_v84_).isdisjoint(d_250_v84_)
                d_254_v87_: _dafny.Array
                nw32_ = _dafny.Array(None, 14)
                d_254_v87_ = nw32_
                d_255_v88_: _dafny.Map
                d_255_v88_ = _dafny.Map({d_117_v0_: d_119_v2_})
                d_256_v89_: _dafny.Map
                d_256_v89_ = _dafny.Map({default__.fm6((0) - (len(d_255_v88_)), d_126_v6_, (0) - (d_117_v0_), d_125_globalState_): d_254_v87_})
                d_254_v87_ = ((d_256_v89_)[(_dafny.Set({d_119_v2_})).ispropersubset(d_250_v84_)] if ((_dafny.Set({d_119_v2_})).ispropersubset(d_250_v84_)) in (d_256_v89_) else d_254_v87_)
            (d_125_globalState_).f15 = default__.safeDivisionInt(default__.safeDivisionInt(d_117_v0_, d_251_i14_), len(d_120_v3_))
            d_257_v90_: _dafny.Map
            d_257_v90_ = _dafny.Map({True: (D0_DC1(d_131_v12_, d_122_v5_, d_121_v4_, d_117_v0_, d_251_i14_)).cf2})
            d_258_v91_: _dafny.Map
            d_258_v91_ = _dafny.Map({d_126_v6_: (d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))]})
            d_259_v92_: _dafny.Map
            d_259_v92_ = _dafny.Map({d_251_i14_: d_137_v15_})
            d_260_v93_: D0
            d_260_v93_ = D0_DC1(d_131_v12_, ((d_257_v90_)[(d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))]] if ((d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))]) in (d_257_v90_) else d_122_v5_), _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_126_v6_ for d_261_i15_ in range(default__.abs(592))]), ((d_258_v91_)[d_126_v6_] if (d_126_v6_) in (d_258_v91_) else default__.fm7((d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], d_251_i14_, default__.fm0(((d_259_v92_)[default__.fm0((d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], d_118_v1_, (d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], d_251_i14_, d_125_globalState_)] if (default__.fm0((d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], d_118_v1_, (d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], d_251_i14_, d_125_globalState_)) in (d_259_v92_) else not(d_119_v2_)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ytn")), d_137_v15_, d_117_v0_, d_125_globalState_), d_125_globalState_)), (d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))], (d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))]]), d_251_i14_, d_117_v0_)
            d_260_v93_ = d_260_v93_
            d_262_v94_: _dafny.Set
            d_263_v95_: bool
            d_264_v96_: _dafny.MultiSet
            out21_: _dafny.Set
            out22_: bool
            out23_: _dafny.MultiSet
            out21_, out22_, out23_ = default__.m0((0) - (default__.safeDivisionInt(d_251_i14_, d_117_v0_)), d_137_v15_, (232) not in (d_128_v8_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ckhhe")), d_125_globalState_)
            d_262_v94_ = out21_
            d_263_v95_ = out22_
            d_264_v96_ = out23_
        d_265_v97_: _dafny.Array
        nw33_ = _dafny.Array(None, 4)
        d_265_v97_ = nw33_
        d_266_v98_: C0
        nw34_ = C0()
        nw34_.ctor__(d_117_v0_, d_126_v6_)
        d_266_v98_ = nw34_
        index28_ = default__.safeIndex(411, (d_265_v97_).length(0))
        (d_265_v97_)[index28_] = d_266_v98_
        index29_ = default__.safeIndex(411, (d_265_v97_).length(0))
        (d_265_v97_)[index29_] = d_266_v98_
        d_267_v99_: _dafny.Seq
        d_267_v99_ = _dafny.SeqWithoutIsStrInference([(d_265_v97_)[default__.safeIndex(411, (d_265_v97_).length(0))]])
        d_268_i16_: int
        d_268_i16_ = 0
        with _dafny.label("2"):
            while (d_267_v99_) != (d_267_v99_):
                with _dafny.c_label("2"):
                    if (d_268_i16_) >= (100):
                        raise _dafny.Break("2")
                    d_268_i16_ = (d_268_i16_) + (1)
                    index30_ = default__.safeIndex(564, (d_139_v17_).length(0))
                    (d_139_v17_)[index30_] = ((d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))]) + ((d_139_v17_)[default__.safeIndex(564, (d_139_v17_).length(0))])
                    d_269_v100_: _dafny.Set
                    d_270_v101_: bool
                    d_271_v102_: _dafny.MultiSet
                    out24_: _dafny.Set
                    out25_: bool
                    out26_: _dafny.MultiSet
                    out24_, out25_, out26_ = default__.m0(((d_135_v13_)[d_117_v0_] if (d_117_v0_) in (d_135_v13_) else (d_266_v98_).f19), not (d_137_v15_) or (not(True)), d_119_v2_, ((d_118_v1_) + ((d_118_v1_).set(default__.safeIndex((d_266_v98_).f19, len(d_118_v1_)), d_126_v6_))).set(default__.safeIndex((d_266_v98_).f19, len((d_118_v1_) + ((d_118_v1_).set(default__.safeIndex((d_266_v98_).f19, len(d_118_v1_)), d_126_v6_)))), d_126_v6_), d_125_globalState_)
                    d_269_v100_ = out24_
                    d_270_v101_ = out25_
                    d_271_v102_ = out26_
                    d_270_v101_ = d_119_v2_
                    if d_119_v2_:
                        d_272_v103_: _dafny.Array
                        nw35_ = _dafny.Array(_dafny.Map({}), 29)
                        d_272_v103_ = nw35_
                        d_273_v104_: _dafny.Seq
                        d_273_v104_ = _dafny.SeqWithoutIsStrInference([d_122_v5_])
                        d_274_v105_: D3
                        d_274_v105_ = D3_DC6(d_273_v104_)
                        d_275_v106_: _dafny.Map
                        d_275_v106_ = _dafny.Map({d_274_v105_: (d_266_v98_).f19})
                        d_276_v107_: _dafny.Map
                        d_276_v107_ = _dafny.Map({d_275_v106_: d_137_v15_})
                        index31_ = default__.safeIndex(196, (d_272_v103_).length(0))
                        (d_272_v103_)[index31_] = d_276_v107_
                        d_277_v108_: _dafny.Map
                        d_277_v108_ = _dafny.Map({(d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))]: d_276_v107_})
                        index32_ = default__.safeIndex(196, (d_272_v103_).length(0))
                        rhs33_ = (((d_277_v108_)[False] if (False) in (d_277_v108_) else d_276_v107_)) | (d_276_v107_)
                        rhs34_ = (138) == (62)
                        lhs22_ = d_272_v103_
                        lhs23_ = default__.safeIndex(196, (d_272_v103_).length(0))
                        lhs24_ = d_125_globalState_
                        lhs22_[lhs23_] = rhs33_
                        lhs24_.f6 = rhs34_
                        d_139_v17_ = d_139_v17_
                        (d_125_globalState_).f5 = default__.safeModuloInt(((d_266_v98_).f19) + ((d_266_v98_).f19), (d_266_v98_).f19)
                        d_278_v109_: D1
                        d_278_v109_ = D1_DC2((d_266_v98_).f20)
                        pat_let_tv0_ = d_126_v6_
                        def iife9_(_pat_let0_0):
                            def iife10_(d_279_dt__update__tmp_h1_):
                                def iife11_(_pat_let1_0):
                                    def iife12_(d_280_dt__update_hcf6_h0_):
                                        return D1_DC2(d_280_dt__update_hcf6_h0_)
                                    return iife12_(_pat_let1_0)
                                return iife11_(pat_let_tv0_)
                            return iife10_(_pat_let0_0)
                        d_278_v109_ = iife9_(d_278_v109_)
                        (d_125_globalState_).f15 = (0) - ((0) - (default__.safeDivisionInt(((0) - ((d_266_v98_).f19)) + (259), (d_266_v98_).fm2(d_117_v0_, (d_266_v98_).f19, (d_266_v98_).f19, d_125_globalState_))))
                    elif True:
                        d_130_v11_ = (d_130_v11_).set(len(d_120_v3_), (d_266_v98_).f19)
                        d_127_v7_ = (default__.fm8((d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))], d_125_globalState_)) + (d_127_v7_)
                        d_118_v1_ = (d_118_v1_) + (_dafny.SeqWithoutIsStrInference([(d_266_v98_).f20 for d_281_i17_ in range(default__.abs(368))]))
                        d_282_v110_: _dafny.Array
                        def lambda20_(d_283_v2_, d_284_v101_):
                            def lambda21_(d_285_i18_):
                                return _dafny.Map({d_283_v2_: d_284_v101_})

                            return lambda21_

                        init10_ = lambda20_(d_119_v2_, d_270_v101_)
                        nw36_ = _dafny.Array(None, 21)
                        for i0_10_ in range(nw36_.length(0)):
                            nw36_[i0_10_] = init10_(i0_10_)
                        d_282_v110_ = nw36_
                        index33_ = default__.safeIndex(514, (d_122_v5_).length(0))
                        rhs35_ = ((d_266_v98_).f19) < (d_117_v0_)
                        rhs36_ = d_282_v110_
                        lhs25_ = d_122_v5_
                        lhs26_ = default__.safeIndex(514, (d_122_v5_).length(0))
                        lhs25_[lhs26_] = rhs35_
                        d_282_v110_ = rhs36_
                        d_286_v111_: _dafny.Set
                        d_287_v112_: bool
                        d_288_v113_: _dafny.MultiSet
                        out27_: _dafny.Set
                        out28_: bool
                        out29_: _dafny.MultiSet
                        out27_, out28_, out29_ = default__.m0((d_266_v98_).f19, d_270_v101_, not((d_122_v5_)[default__.safeIndex(514, (d_122_v5_).length(0))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "svi")), d_125_globalState_)
                        d_286_v111_ = out27_
                        d_287_v112_ = out28_
                        d_288_v113_ = out29_
                    pass
            pass
        index34_ = default__.safeIndex(411, (d_265_v97_).length(0))
        (d_265_v97_)[index34_] = (d_265_v97_)[default__.safeIndex(411, (d_265_v97_).length(0))]
        _dafny.print(_dafny.string_of(d_117_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_118_v1_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_119_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v3_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v4_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qljvqw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qljvqw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qljvqw"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_)[28]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_125_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_125_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_globalState_).f7) == (_dafny.SeqWithoutIsStrInference([327, 6, 327, -327]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_globalState_).f8) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_globalState_).f12) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qljvqw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qljvqw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qljvqw"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_125_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_125_globalState_).f14)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_125_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_126_v6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_127_v7_) == (_dafny.SeqWithoutIsStrInference([-349]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_128_v8_) == (_dafny.Set({327, -349}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v9_) == (_dafny.MultiSet([_dafny.Set({1, 176, 166, -3}), _dafny.Set({327, -349})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_v11_) == (_dafny.Map({327: -327}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v12_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v13_) == (_dafny.MultiSet([327, 327]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v14_) == (_dafny.Set({_dafny.Set({327}), _dafny.Set({327, -29, -28, -27, -26, -25, -24, -23, -22, -21, -20, -19, -18, -17, -16, -15, -14, -13, -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92}), _dafny.Set({1, 176, 166, -3})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_137_v15_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v16_) == (_dafny.MultiSet([_dafny.Map({False: False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_139_v17_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_139_v17_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_139_v17_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_139_v17_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_139_v17_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_139_v17_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_139_v17_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_139_v17_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_139_v17_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_139_v17_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_142_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v47_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v47_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v47_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v47_).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_243_i13_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_250_v84_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v97_)[3]).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v97_)[3]).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_266_v98_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_266_v98_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_267_v99_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_268_i16_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.Array(None, 0), _dafny.Array(None, 0), _dafny.MultiSet({}), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
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

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({self.cf0.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC3(False, False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC3(D1, NamedTuple('DC3', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC5(D2, NamedTuple('DC5', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC7()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)

class D3_DC7(D3, NamedTuple('DC7', [])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7)
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC6(D3, NamedTuple('DC6', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)

class D4_DC10(D4, NamedTuple('DC10', [('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf19', Any), ('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f5: int = int(0)
        self.f6: bool = False
        self.f13: int = int(0)
        self.f15: int = int(0)
        self._f0: int = int(0)
        self._f1: int = int(0)
        self._f2: bool = False
        self._f3: int = int(0)
        self._f4: bool = False
        self._f7: _dafny.Seq = _dafny.Seq({})
        self._f8: _dafny.Map = _dafny.Map({})
        self._f9: int = int(0)
        self._f10: int = int(0)
        self._f11: str = _dafny.CodePoint('D')
        self._f12: _dafny.MultiSet = _dafny.MultiSet({})
        self._f14: _dafny.Seq = _dafny.Seq({})
        self._f16: bool = False
        self._f17: int = int(0)
        self._f18: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4
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
    def f14(self):
        return self._f14
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18

class C0:
    def  __init__(self):
        self._f19: int = int(0)
        self._f20: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f19, f20):
        (self)._f19 = f19
        (self)._f20 = f20

    def fm2(self, p0, p1, p2, globalState):
        return (self).f19

    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
