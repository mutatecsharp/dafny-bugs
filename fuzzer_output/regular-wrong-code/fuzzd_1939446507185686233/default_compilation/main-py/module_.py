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
        return (864) != ((len(_dafny.Set({572}))) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_0_i0_ in range(default__.abs(-483))]))))

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in (_dafny.Map({587: False})).keys.Elements:
                d_1_v0_: int = compr_0_
                if (d_1_v0_) in (_dafny.Map({587: False})):
                    coll0_ = coll0_.union(_dafny.Set([(d_1_v0_) - ((_dafny.MultiSet([not(True)])).cardinality)]))
            return _dafny.Set(coll0_)
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(182, 877):
                d_2_v1_: int = compr_1_
                if ((182) <= (d_2_v1_)) and ((d_2_v1_) < (877)):
                    coll1_ = coll1_.union(_dafny.Set([(d_2_v1_) + (len(_dafny.Set({False})))]))
            return _dafny.Set(coll1_)
        return ((iife0_()
        ) | (iife1_()
        )).intersection(_dafny.Set({-216}))

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return default__.safeDivisionInt(len((_dafny.Set({False})).intersection(_dafny.Set({False}))), 81)

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nargnu"))

    @staticmethod
    def fm5(p0, globalState):
        source0_ = D2_DC2(_dafny.CodePoint('a'))
        if source0_.is_DC3:
            d_3___mcc_h0_ = source0_.cf3
            d_4___mcc_h1_ = source0_.cf4
            d_5_cf4_ = d_4___mcc_h1_
            d_6_cf3_ = d_3___mcc_h0_
            return _dafny.CodePoint('x')
        elif source0_.is_DC4:
            d_7___mcc_h2_ = source0_.cf5
            d_8___mcc_h3_ = source0_.cf6
            d_9_cf6_ = d_8___mcc_h3_
            d_10_cf5_ = d_7___mcc_h2_
            return (D2_DC2(_dafny.CodePoint('l'))).cf2
        elif True:
            d_11___mcc_h4_ = source0_.cf2
            d_12_cf2_ = d_11___mcc_h4_
            return d_12_cf2_

    @staticmethod
    def fm6(p0, p1, globalState):
        return ((_dafny.MultiSet([not(False), True, not(True)])) - (_dafny.MultiSet([False, False]))) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True, True]))) - (_dafny.MultiSet([False, False, not(not(not(True)))])))

    @staticmethod
    def fm7(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({-644: 320})) | (_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False]))).cardinality: (0) - (len(_dafny.SeqWithoutIsStrInference([-380])))}))) | ((_dafny.Map({-569: -889})) | (_dafny.Map({432: 210})))

    @staticmethod
    def fm8(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([531 for d_13_i0_ in range(default__.abs(601))])) + (_dafny.SeqWithoutIsStrInference([238]))) + (_dafny.SeqWithoutIsStrInference([209 for d_14_i1_ in range(default__.abs(779))]))

    @staticmethod
    def fm9(p0, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: D2
            for compr_2_ in (_dafny.Map({D2_DC3(not(not(False)), True): _dafny.Set({False})})).keys.Elements:
                d_15_v0_: D2 = compr_2_
                if (d_15_v0_) in (_dafny.Map({D2_DC3(not(not(False)), True): _dafny.Set({False})})):
                    coll2_ = coll2_.union(_dafny.Set([d_15_v0_]))
            return _dafny.Set(coll2_)
        return ((_dafny.Map({(_dafny.MultiSet([len(_dafny.Set({827})), len(_dafny.Map({True: True}))])).cardinality: _dafny.Set({D2_DC3(False, True)})})) | (_dafny.Map({-47: iife2_()
        }))) | ((_dafny.Map({-487: _dafny.Set({D2_DC3(False, True)})})) | (_dafny.Map({len(_dafny.Map({-654: not(True)})): _dafny.Set({D2_DC3(True, True)})})))

    @staticmethod
    def fm10(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_16_i1_ in range(default__.abs(493))]) for d_17_i0_ in range(default__.abs(220))])

    @staticmethod
    def fm11(p0, globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(808, -796):
                d_18_v0_: int = compr_3_
                if ((808) <= (d_18_v0_)) and ((d_18_v0_) < (-796)):
                    coll3_ = coll3_.union(_dafny.Set([(d_18_v0_) + (10)]))
            return _dafny.Set(coll3_)
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwsaa"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmyeoj"))), (_dafny.Map({len(_dafny.Map({False: False})): False})) != (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(0) - (len(iife3_()
        )), 754, 462, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkpurq")))])): not(False)})), (_dafny.SeqWithoutIsStrInference([898, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agfup")))}))])) != (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: -752})), len(_dafny.SeqWithoutIsStrInference([False])), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.Set({not(True)})), (D4_DC7(102, True)).cf9])]))])).cardinality, 86, len(_dafny.Map({False: 741}))]))])

    @staticmethod
    def fm12(p0, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(857, 355):
                d_19_v0_: int = compr_4_
                if ((857) <= (d_19_v0_)) and ((d_19_v0_) < (355)):
                    coll4_[default__.safeDivisionInt(d_19_v0_, 886)] = _dafny.SeqWithoutIsStrInference([True])
            return _dafny.Map(coll4_)
        if not(((D8_DC17(_dafny.MultiSet([_dafny.Map({False: 977}), _dafny.Map({False: 101})]))).cf28).isdisjoint(_dafny.MultiSet([_dafny.Map({True: len(iife4_()
        )})]))):
            return D2_DC3(not(True), False)
        elif True:
            return D2_DC3(not(True), False)

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        if True:
            return (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_20_i0_ in range(default__.abs(-551))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))}))
        elif True:
            return _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ll"))})

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: _dafny.MultiSet
            for compr_5_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([not(True)])])).Elements:
                d_21_v0_: _dafny.MultiSet = compr_5_
                if (d_21_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([not(True)])])):
                    coll5_[d_21_v0_] = False
            return _dafny.Map(coll5_)
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: _dafny.MultiSet
            for compr_6_ in (_dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])), _dafny.MultiSet([True]), _dafny.MultiSet([True, False]), _dafny.MultiSet([False])])).Elements:
                d_22_v1_: _dafny.MultiSet = compr_6_
                if (d_22_v1_) in (_dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])), _dafny.MultiSet([True]), _dafny.MultiSet([True, False]), _dafny.MultiSet([False])])):
                    coll6_[d_22_v1_] = False
            return _dafny.Map(coll6_)
        return ((_dafny.Map({_dafny.MultiSet([True]): not(not(False))})) | (iife5_()
        )) | ((iife6_()
        ) | (_dafny.Map({_dafny.MultiSet([(D8_DC18((0) - (len(_dafny.Set({-998}))), True)).cf30]): not(True)})))

    @staticmethod
    def fm15(globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({True: False}))) | ((_dafny.Map({True: False})) | (_dafny.Map({False: True})))

    @staticmethod
    def fm16(globalState):
        source1_ = D6_DC13()
        if source1_.is_DC11:
            d_23___mcc_h0_ = source1_.cf17
            d_24___mcc_h1_ = source1_.cf18
            d_25___mcc_h2_ = source1_.cf19
            d_26___mcc_h3_ = source1_.cf20
            d_27___mcc_h4_ = source1_.cf21
            d_28_cf21_ = d_27___mcc_h4_
            d_29_cf20_ = d_26___mcc_h3_
            d_30_cf19_ = d_25___mcc_h2_
            d_31_cf18_ = d_24___mcc_h1_
            d_32_cf17_ = d_23___mcc_h0_
            return _dafny.Map({_dafny.Set({True, True}): d_28_cf21_})
        elif source1_.is_DC12:
            d_33___mcc_h5_ = source1_.cf22
            d_34_cf22_ = d_33___mcc_h5_
            return _dafny.Map({_dafny.Set({d_34_cf22_}): -495})
        elif source1_.is_DC13:
            return _dafny.Map({_dafny.Set({not(False), True, True, True, False}): (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hw"))))})
        elif source1_.is_DC10:
            d_35___mcc_h6_ = source1_.cf16
            d_36_cf16_ = d_35___mcc_h6_
            return _dafny.Map({_dafny.Set({True}): 414})
        elif True:
            d_37___mcc_h7_ = source1_.cf23
            d_38_cf23_ = d_37___mcc_h7_
            return _dafny.Map({_dafny.Set({True}): -560})

    @staticmethod
    def fm17(p0, p1, p2, globalState):
        return _dafny.Set({True})

    @staticmethod
    def m1(p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_39_v0_: bool
        d_39_v0_ = False
        index0_ = default__.safeIndex(666, (p1).length(0))
        (p1)[index0_] = d_39_v0_
        d_40_v1_: _dafny.Seq
        d_40_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdfamt"))
        d_41_v2_: _dafny.MultiSet
        d_41_v2_ = _dafny.MultiSet([p0, (0) - (p0)])
        d_42_v3_: _dafny.Seq
        d_42_v3_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, len(d_40_v1_), ((d_41_v2_)[p0] if (p0) in (d_41_v2_) else p0)])
        index1_ = default__.safeIndex(666, (p1).length(0))
        (p1)[index1_] = not ((_dafny.SeqWithoutIsStrInference([p0, p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "benlgnvi"))), p0])) < (d_42_v3_)) or (d_39_v0_)
        d_43_v4_: D2
        d_43_v4_ = D2_DC3((p1)[default__.safeIndex(666, (p1).length(0))], (p1)[default__.safeIndex(666, (p1).length(0))])
        d_44_v5_: _dafny.Array
        nw0_ = _dafny.Array(None, 3)
        nw0_[int(0)] = D2_DC3((p1)[default__.safeIndex(666, (p1).length(0))], (p1)[default__.safeIndex(666, (p1).length(0))])
        nw0_[int(1)] = d_43_v4_
        nw0_[int(2)] = d_43_v4_
        d_44_v5_ = nw0_
        d_45_v6_: _dafny.Seq
        d_45_v6_ = _dafny.SeqWithoutIsStrInference([d_44_v5_])
        d_45_v6_ = d_45_v6_
        d_46_v7_: _dafny.Seq
        d_46_v7_ = _dafny.SeqWithoutIsStrInference([(p1)[default__.safeIndex(666, (p1).length(0))]])
        if (((p1)[default__.safeIndex(666, (p1).length(0))]) not in (d_46_v7_)) and (not(not (default__.fm0(d_39_v0_, True, 444, p0, globalState)) or (not((p1)[default__.safeIndex(666, (p1).length(0))])))):
            d_47_v8_: bool
            d_47_v8_ = (p1)[default__.safeIndex(666, (p1).length(0))]
            d_48_v9_: C1
            nw1_ = C1()
            nw1_.ctor__(d_47_v8_)
            d_48_v9_ = nw1_
            d_49_v10_: str
            d_49_v10_ = _dafny.CodePoint('i')
            d_50_v11_: D5
            d_50_v11_ = D5_DC9(d_40_v1_, d_39_v0_, d_48_v9_, d_49_v10_)
            pat_let_tv0_ = d_48_v9_
            def iife7_(_pat_let0_0):
                def iife8_(d_51_dt__update__tmp_h0_):
                    def iife9_(_pat_let1_0):
                        def iife10_(d_52_dt__update_hcf14_h0_):
                            return D5_DC9((d_51_dt__update__tmp_h0_).cf12, (d_51_dt__update__tmp_h0_).cf13, d_52_dt__update_hcf14_h0_, (d_51_dt__update__tmp_h0_).cf15)
                        return iife10_(_pat_let1_0)
                    return iife9_(pat_let_tv0_)
                return iife8_(_pat_let0_0)
            source2_ = iife7_((D5_DC9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")), d_39_v0_, d_48_v9_, d_49_v10_) if not(d_39_v0_) else d_50_v11_))
            if source2_.is_DC9:
                d_53___mcc_h0_ = source2_.cf12
                d_54___mcc_h1_ = source2_.cf13
                d_55___mcc_h2_ = source2_.cf14
                d_56___mcc_h3_ = source2_.cf15
                d_57_cf15_ = d_56___mcc_h3_
                d_58_cf14_ = d_55___mcc_h2_
                d_59_cf13_ = d_54___mcc_h1_
                d_60_cf12_ = d_53___mcc_h0_
                d_61_v12_: _dafny.Set
                d_61_v12_ = _dafny.Set({d_59_cf13_})
                d_62_v13_: _dafny.Map
                d_62_v13_ = _dafny.Map({d_61_v12_: (_dafny.MultiSet([(p1)[default__.safeIndex(666, (p1).length(0))]])).cardinality})
                d_63_v14_: C0
                nw2_ = C0()
                nw2_.ctor__((d_60_cf12_)[default__.safeIndex(p0, len(d_60_cf12_))])
                d_63_v14_ = nw2_
                d_64_v15_: _dafny.Map
                d_64_v15_ = _dafny.Map({(default__.fm16(globalState)) != (d_62_v13_): d_63_v14_})
                def iife11_():
                    coll7_ = _dafny.Set()
                    compr_7_: int
                    for compr_7_ in _dafny.IntegerRange(294, 485):
                        d_65_v16_: int = compr_7_
                        if ((294) <= (d_65_v16_)) and ((d_65_v16_) < (485)):
                            coll7_ = coll7_.union(_dafny.Set([(d_65_v16_) - (len(_dafny.SeqWithoutIsStrInference([d_40_v1_, d_60_cf12_])))]))
                    return _dafny.Set(coll7_)
                rhs0_ = (len((d_46_v7_ if (p1)[default__.safeIndex(666, (p1).length(0))] else d_46_v7_))) not in (iife11_()
                )
                rhs1_ = (d_39_v0_) == (default__.fm0(d_39_v0_, d_39_v0_, p0, p0, globalState))
                rhs2_ = d_64_v15_
                lhs0_ = globalState
                d_39_v0_ = rhs0_
                lhs0_.f9 = rhs1_
                d_64_v15_ = rhs2_
                d_66_v17_: C0
                nw3_ = C0()
                nw3_.ctor__(d_49_v10_)
                d_66_v17_ = nw3_
                d_67_v18_: _dafny.Map
                d_67_v18_ = _dafny.Map({(d_41_v2_).cardinality: p1})
                d_68_v19_: D5
                d_68_v19_ = D5_DC8((d_67_v18_) | (d_67_v18_))
                d_68_v19_ = (d_68_v19_ if d_39_v0_ else d_68_v19_)
                d_69_v20_: _dafny.Map
                d_69_v20_ = _dafny.Map({(0) - (default__.fm2(d_39_v0_, p0, p0, d_60_cf12_, globalState)): d_49_v10_})
                d_70_v21_: _dafny.Map
                d_70_v21_ = _dafny.Map({d_39_v0_: p0})
                d_71_v22_: _dafny.Map
                d_71_v22_ = _dafny.Map({p0: d_70_v21_})
                d_72_v23_: _dafny.Array
                nw4_ = _dafny.Array(None, 28)
                nw4_[int(0)] = _dafny.CodePoint('k')
                nw4_[int(1)] = d_57_cf15_
                nw4_[int(2)] = d_57_cf15_
                nw4_[int(3)] = d_57_cf15_
                nw4_[int(4)] = ((d_69_v20_)[285] if (285) in (d_69_v20_) else d_49_v10_)
                nw4_[int(5)] = d_57_cf15_
                nw4_[int(6)] = _dafny.CodePoint('e')
                nw4_[int(7)] = (d_63_v14_).f12
                nw4_[int(8)] = (d_66_v17_).f12
                nw4_[int(9)] = d_49_v10_
                nw4_[int(10)] = (d_66_v17_).f12
                nw4_[int(11)] = d_49_v10_
                nw4_[int(12)] = (d_66_v17_).f12
                nw4_[int(13)] = d_57_cf15_
                nw4_[int(14)] = (d_40_v1_)[default__.safeIndex(-569, len(d_40_v1_))]
                nw4_[int(15)] = d_49_v10_
                nw4_[int(16)] = (d_63_v14_).f12
                nw4_[int(17)] = (d_63_v14_).f12
                nw4_[int(18)] = default__.fm5(len(d_71_v22_), globalState)
                nw4_[int(19)] = d_49_v10_
                nw4_[int(20)] = d_49_v10_
                nw4_[int(21)] = (d_66_v17_).f12
                nw4_[int(22)] = d_49_v10_
                nw4_[int(23)] = d_57_cf15_
                nw4_[int(24)] = d_49_v10_
                nw4_[int(25)] = (d_63_v14_).f12
                nw4_[int(26)] = _dafny.CodePoint('f')
                nw4_[int(27)] = (d_63_v14_).f12
                d_72_v23_ = nw4_
                d_73_v24_: _dafny.Map
                d_73_v24_ = _dafny.Map({d_59_cf13_: d_72_v23_})
                d_74_v25_: _dafny.Seq
                d_74_v25_ = _dafny.SeqWithoutIsStrInference([((d_73_v24_)[(p1)[default__.safeIndex(666, (p1).length(0))]] if ((p1)[default__.safeIndex(666, (p1).length(0))]) in (d_73_v24_) else d_72_v23_)])
                d_39_v0_ = not (True) or ((d_72_v23_) not in ((d_74_v25_).set(default__.safeIndex(p0, len(d_74_v25_)), d_72_v23_)))
            elif True:
                d_75___mcc_h4_ = source2_.cf11
                d_76_cf11_ = d_75___mcc_h4_
                index2_ = default__.safeIndex(666, (p1).length(0))
                (p1)[index2_] = (p1)[default__.safeIndex(666, (p1).length(0))]
                (globalState).f4 = (p0) * (p0)
                d_77_v26_: _dafny.Array
                nw5_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                d_77_v26_ = nw5_
                index3_ = default__.safeIndex(104, (d_77_v26_).length(0))
                (d_77_v26_)[index3_] = d_49_v10_
                index4_ = default__.safeIndex(104, (d_77_v26_).length(0))
                rhs3_ = -951
                rhs4_ = d_40_v1_
                rhs5_ = d_49_v10_
                lhs1_ = globalState
                lhs2_ = d_77_v26_
                lhs3_ = default__.safeIndex(104, (d_77_v26_).length(0))
                lhs1_.f4 = rhs3_
                d_40_v1_ = rhs4_
                lhs2_[lhs3_] = rhs5_
                d_78_v27_: _dafny.Map
                d_78_v27_ = _dafny.Map({p0: (d_77_v26_)[default__.safeIndex(104, (d_77_v26_).length(0))]})
                d_79_v28_: _dafny.Map
                d_79_v28_ = _dafny.Map({p0: d_78_v27_})
                (globalState).f4 = default__.safeModuloInt(default__.safeDivisionInt(p0, p0), (p0 if (p1)[default__.safeIndex(666, (p1).length(0))] else (0) - ((0) - (len(((d_79_v28_)[p0] if (p0) in (d_79_v28_) else _dafny.Map({p0: d_49_v10_})))))))
            d_80_v29_: _dafny.MultiSet
            d_80_v29_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqp")), _dafny.SeqWithoutIsStrInference([d_49_v10_ for d_81_i0_ in range(default__.abs(12))])])
            d_82_v30_: _dafny.MultiSet
            d_82_v30_ = _dafny.MultiSet([d_39_v0_])
            d_83_v31_: _dafny.Map
            d_83_v31_ = _dafny.Map({p0: d_82_v30_})
            d_84_v32_: _dafny.Map
            d_84_v32_ = _dafny.Map({(p1)[default__.safeIndex(666, (p1).length(0))]: d_39_v0_})
            d_85_v33_: _dafny.Map
            d_85_v33_ = _dafny.Map({p0: (p1)[default__.safeIndex(666, (p1).length(0))]})
            d_86_v34_: _dafny.Map
            d_86_v34_ = _dafny.Map({len(d_85_v33_): p0})
            d_87_v35_: _dafny.Array
            nw6_ = _dafny.Array(None, 21)
            nw6_[int(0)] = p0
            nw6_[int(1)] = (p0) * (p0)
            nw6_[int(2)] = p0
            nw6_[int(3)] = ((d_80_v29_)[d_40_v1_] if (d_40_v1_) in (d_80_v29_) else len(d_40_v1_))
            nw6_[int(4)] = len(_dafny.SeqWithoutIsStrInference([d_40_v1_ for d_88_i1_ in range(default__.abs(255))]))
            nw6_[int(5)] = p0
            nw6_[int(6)] = p0
            nw6_[int(7)] = (0) - (p0)
            nw6_[int(8)] = (p0) - (p0)
            nw6_[int(9)] = p0
            nw6_[int(10)] = len(d_83_v31_)
            nw6_[int(11)] = 936
            nw6_[int(12)] = p0
            nw6_[int(13)] = (0) - (len((d_84_v32_).set(d_39_v0_, True)))
            nw6_[int(14)] = p0
            nw6_[int(15)] = p0
            nw6_[int(16)] = p0
            nw6_[int(17)] = p0
            nw6_[int(18)] = (p0) * (((d_82_v30_)[d_39_v0_] if (d_39_v0_) in (d_82_v30_) else len(d_86_v34_)))
            nw6_[int(19)] = 405
            nw6_[int(20)] = (p0 if d_39_v0_ else p0)
            d_87_v35_ = nw6_
            index5_ = default__.safeIndex(878, (d_87_v35_).length(0))
            (d_87_v35_)[index5_] = p0
            index6_ = default__.safeIndex(878, (d_87_v35_).length(0))
            (d_87_v35_)[index6_] = ((d_86_v34_)[(D4_DC7(((d_82_v30_)[d_39_v0_] if (d_39_v0_) in (d_82_v30_) else p0), (p1)[default__.safeIndex(666, (p1).length(0))])).cf9] if ((D4_DC7(((d_82_v30_)[d_39_v0_] if (d_39_v0_) in (d_82_v30_) else p0), (p1)[default__.safeIndex(666, (p1).length(0))])).cf9) in (d_86_v34_) else p0)
            d_89_v36_: C0
            nw7_ = C0()
            nw7_.ctor__((default__.fm5(len((d_40_v1_).set(default__.safeIndex(768, len(d_40_v1_)), d_49_v10_)), globalState) if default__.fm0((p1)[default__.safeIndex(666, (p1).length(0))], True, (0) - (p0), (d_87_v35_)[default__.safeIndex(878, (d_87_v35_).length(0))], globalState) else d_49_v10_))
            d_89_v36_ = nw7_
            (globalState).f10 = not((p1)[default__.safeIndex(666, (p1).length(0))])
            d_90_v37_: _dafny.Set
            d_90_v37_ = _dafny.Set({d_40_v1_, d_40_v1_})
            nw8_ = _dafny.Array(None, 12)
            nw8_[int(0)] = (d_87_v35_)[default__.safeIndex(878, (d_87_v35_).length(0))]
            nw8_[int(1)] = (d_87_v35_)[default__.safeIndex(878, (d_87_v35_).length(0))]
            nw8_[int(2)] = p0
            nw8_[int(3)] = p0
            nw8_[int(4)] = p0
            nw8_[int(5)] = (d_87_v35_)[default__.safeIndex(878, (d_87_v35_).length(0))]
            nw8_[int(6)] = (d_87_v35_)[default__.safeIndex(878, (d_87_v35_).length(0))]
            nw8_[int(7)] = len(d_90_v37_)
            nw8_[int(8)] = (d_87_v35_)[default__.safeIndex(878, (d_87_v35_).length(0))]
            nw8_[int(9)] = default__.safeModuloInt(p0, p0)
            nw8_[int(10)] = len((d_46_v7_) + (d_46_v7_))
            nw8_[int(11)] = (d_87_v35_)[default__.safeIndex(878, (d_87_v35_).length(0))]
            d_87_v35_ = nw8_
        elif True:
            d_91_v38_: bool
            d_91_v38_ = not(d_39_v0_)
            d_92_v39_: C1
            nw9_ = C1()
            nw9_.ctor__(d_91_v38_)
            d_92_v39_ = nw9_
            d_93_v40_: _dafny.Map
            d_93_v40_ = _dafny.Map({d_92_v39_: d_39_v0_})
            d_94_v41_: D6
            d_94_v41_ = D6_DC10(d_93_v40_)
            d_95_v42_: _dafny.Map
            d_95_v42_ = _dafny.Map({d_41_v2_: (d_94_v41_).cf16})
            d_96_v43_: _dafny.Seq
            d_96_v43_ = _dafny.SeqWithoutIsStrInference([d_93_v40_])
            d_97_v44_: _dafny.Map
            d_97_v44_ = _dafny.Map({-989: d_93_v40_})
            d_98_v45_: _dafny.Map
            d_98_v45_ = _dafny.Map({(p1)[default__.safeIndex(666, (p1).length(0))]: d_92_v39_})
            d_99_v46_: _dafny.Seq
            d_99_v46_ = _dafny.SeqWithoutIsStrInference([d_92_v39_])
            d_100_v47_: _dafny.Array
            nw10_ = _dafny.Array(None, 25)
            nw10_[int(0)] = ((d_95_v42_)[d_41_v2_] if (d_41_v2_) in (d_95_v42_) else d_93_v40_)
            nw10_[int(1)] = d_93_v40_
            nw10_[int(2)] = d_93_v40_
            nw10_[int(3)] = d_93_v40_
            nw10_[int(4)] = d_93_v40_
            nw10_[int(5)] = (d_93_v40_) | (d_93_v40_)
            nw10_[int(6)] = _dafny.Map({d_92_v39_: (p1)[default__.safeIndex(666, (p1).length(0))]})
            nw10_[int(7)] = d_93_v40_
            nw10_[int(8)] = (d_96_v43_)[default__.safeIndex(p0, len(d_96_v43_))]
            nw10_[int(9)] = d_93_v40_
            nw10_[int(10)] = (d_93_v40_) | (d_93_v40_)
            nw10_[int(11)] = d_93_v40_
            nw10_[int(12)] = d_93_v40_
            nw10_[int(13)] = ((d_93_v40_).set(d_92_v39_, True)) | (_dafny.Map({d_92_v39_: d_39_v0_}))
            nw10_[int(14)] = d_93_v40_
            nw10_[int(15)] = _dafny.Map({d_92_v39_: default__.fm0(d_39_v0_, default__.fm0(True, d_39_v0_, p0, p0, globalState), p0, p0, globalState)})
            nw10_[int(16)] = d_93_v40_
            nw10_[int(17)] = (((d_97_v44_)[len(d_42_v3_)] if (len(d_42_v3_)) in (d_97_v44_) else d_93_v40_)) | (d_93_v40_)
            nw10_[int(18)] = d_93_v40_
            nw10_[int(19)] = d_93_v40_
            nw10_[int(20)] = d_93_v40_
            nw10_[int(21)] = _dafny.Map({((d_98_v45_)[d_39_v0_] if (d_39_v0_) in (d_98_v45_) else d_92_v39_): d_39_v0_})
            nw10_[int(22)] = d_93_v40_
            nw10_[int(23)] = _dafny.Map({(d_99_v46_)[default__.safeIndex(p0, len(d_99_v46_))]: (p1)[default__.safeIndex(666, (p1).length(0))]})
            nw10_[int(24)] = (d_93_v40_) | (_dafny.Map({d_92_v39_: d_39_v0_}))
            d_100_v47_ = nw10_
            d_101_v48_: _dafny.Map
            d_101_v48_ = _dafny.Map({p0: (0) - (p0)})
            d_102_v49_: _dafny.Array
            nw11_ = _dafny.Array(None, 3)
            nw11_[int(0)] = p0
            nw11_[int(1)] = ((d_101_v48_)[p0] if (p0) in (d_101_v48_) else p0)
            nw11_[int(2)] = p0
            d_102_v49_ = nw11_
            index7_ = default__.safeIndex(242, (d_102_v49_).length(0))
            (d_102_v49_)[index7_] = default__.fm2(d_39_v0_, p0, p0, d_40_v1_, globalState)
            d_103_v50_: _dafny.Array
            nw12_ = _dafny.Array(_dafny.Set({}), 7)
            d_103_v50_ = nw12_
            d_104_v51_: _dafny.Set
            d_104_v51_ = _dafny.Set({d_39_v0_})
            index8_ = default__.safeIndex(267, (d_103_v50_).length(0))
            (d_103_v50_)[index8_] = d_104_v51_
            d_105_v52_: _dafny.Map
            d_105_v52_ = _dafny.Map({d_41_v2_: d_100_v47_})
            d_106_v53_: str
            d_106_v53_ = _dafny.CodePoint('h')
            d_107_v54_: C0
            nw13_ = C0()
            nw13_.ctor__(d_106_v53_)
            d_107_v54_ = nw13_
            d_108_v55_: _dafny.Map
            d_108_v55_ = _dafny.Map({p0: _dafny.Map({p0: (p1)[default__.safeIndex(666, (p1).length(0))]})})
            index9_ = default__.safeIndex(242, (d_102_v49_).length(0))
            index10_ = default__.safeIndex(267, (d_103_v50_).length(0))
            rhs6_ = ((d_105_v52_)[_dafny.MultiSet([len((D6_DC11(d_40_v1_, d_107_v54_, len(_dafny.SeqWithoutIsStrInference([d_107_v54_])), _dafny.SeqWithoutIsStrInference([p0, p0]), p0)).cf20), p0, p0, p0, -582])] if (_dafny.MultiSet([len((D6_DC11(d_40_v1_, d_107_v54_, len(_dafny.SeqWithoutIsStrInference([d_107_v54_])), _dafny.SeqWithoutIsStrInference([p0, p0]), p0)).cf20), p0, p0, p0, -582])) in (d_105_v52_) else d_100_v47_)
            rhs7_ = p0
            rhs8_ = (0) - (p0)
            rhs9_ = d_39_v0_
            rhs10_ = default__.fm17(d_108_v55_, d_40_v1_, p0, globalState)
            lhs4_ = d_102_v49_
            lhs5_ = default__.safeIndex(242, (d_102_v49_).length(0))
            lhs6_ = globalState
            lhs7_ = globalState
            lhs8_ = d_103_v50_
            lhs9_ = default__.safeIndex(267, (d_103_v50_).length(0))
            d_100_v47_ = rhs6_
            lhs4_[lhs5_] = rhs7_
            lhs6_.f4 = rhs8_
            lhs7_.f9 = rhs9_
            lhs8_[lhs9_] = rhs10_
            (globalState).f9 = d_39_v0_
            (globalState).f4 = (p0) + (p0)
            d_109_v56_: _dafny.Array
            def lambda0_(d_110_v7_):
                def lambda1_(d_111_i2_):
                    return d_110_v7_

                return lambda1_

            init0_ = lambda0_(d_46_v7_)
            nw14_ = _dafny.Array(None, 16)
            for i0_0_ in range(nw14_.length(0)):
                nw14_[i0_0_] = init0_(i0_0_)
            d_109_v56_ = nw14_
            index11_ = default__.safeIndex(633, (d_109_v56_).length(0))
            (d_109_v56_)[index11_] = (d_46_v7_) + (d_46_v7_)
            d_112_v57_: D7
            d_112_v57_ = D7_DC15(d_46_v7_)
            index12_ = default__.safeIndex(633, (d_109_v56_).length(0))
            (d_109_v56_)[index12_] = (d_112_v57_).cf24
            d_113_v58_: _dafny.Array
            nw15_ = _dafny.Array(None, 17)
            nw15_[int(0)] = p1
            nw15_[int(1)] = (p1 if d_39_v0_ else p1)
            nw15_[int(2)] = p1
            nw15_[int(3)] = p1
            nw15_[int(4)] = p1
            nw15_[int(5)] = p1
            nw15_[int(6)] = p1
            nw15_[int(7)] = p1
            nw15_[int(8)] = p1
            nw15_[int(9)] = p1
            nw15_[int(10)] = p1
            nw15_[int(11)] = p1
            nw15_[int(12)] = p1
            nw15_[int(13)] = p1
            nw15_[int(14)] = p1
            nw15_[int(15)] = p1
            nw15_[int(16)] = p1
            d_113_v58_ = nw15_
            index13_ = default__.safeIndex(85, (d_113_v58_).length(0))
            (d_113_v58_)[index13_] = p1
            index14_ = default__.safeIndex(85, (d_113_v58_).length(0))
            (d_113_v58_)[index14_] = p1
        d_114_v59_: _dafny.Array
        def lambda2_(d_115_v3_):
            def lambda3_(d_116_i3_):
                return d_115_v3_

            return lambda3_

        init1_ = lambda2_(d_42_v3_)
        nw16_ = _dafny.Array(None, 10)
        for i0_1_ in range(nw16_.length(0)):
            nw16_[i0_1_] = init1_(i0_1_)
        d_114_v59_ = nw16_
        index15_ = default__.safeIndex(21, (d_114_v59_).length(0))
        (d_114_v59_)[index15_] = d_42_v3_
        d_117_v60_: _dafny.Array
        nw17_ = _dafny.Array(None, 18)
        d_117_v60_ = nw17_
        d_118_v61_: bool
        d_118_v61_ = False
        d_119_v62_: C1
        nw18_ = C1()
        nw18_.ctor__(d_118_v61_)
        d_119_v62_ = nw18_
        index16_ = default__.safeIndex(430, (d_117_v60_).length(0))
        (d_117_v60_)[index16_] = d_119_v62_
        d_120_v63_: _dafny.Set
        d_120_v63_ = _dafny.Set({p0, (d_41_v2_).cardinality})
        d_121_v64_: _dafny.MultiSet
        d_121_v64_ = _dafny.MultiSet([d_120_v63_])
        index17_ = default__.safeIndex(21, (d_114_v59_).length(0))
        index18_ = default__.safeIndex(430, (d_117_v60_).length(0))
        rhs11_ = (d_42_v3_) + ((d_42_v3_) + (d_42_v3_))
        rhs12_ = d_119_v62_
        rhs13_ = d_119_v62_
        rhs14_ = d_121_v64_
        lhs10_ = d_114_v59_
        lhs11_ = default__.safeIndex(21, (d_114_v59_).length(0))
        lhs12_ = d_117_v60_
        lhs13_ = default__.safeIndex(430, (d_117_v60_).length(0))
        lhs10_[lhs11_] = rhs11_
        lhs12_[lhs13_] = rhs12_
        d_119_v62_ = rhs13_
        d_121_v64_ = rhs14_
        d_122_v65_: _dafny.Set
        d_122_v65_ = _dafny.Set({(p1)[default__.safeIndex(666, (p1).length(0))]})
        if (d_122_v65_).issubset(d_122_v65_):
            d_123_v66_: _dafny.Map
            d_123_v66_ = _dafny.Map({d_118_v61_: not((p1)[default__.safeIndex(666, (p1).length(0))])})
            d_123_v66_ = (d_123_v66_).set((d_119_v62_).f11, d_39_v0_)
            d_40_v1_ = ((d_40_v1_ if True else d_40_v1_) if d_39_v0_ else (d_40_v1_) + (d_40_v1_))
            d_124_v67_: _dafny.Map
            d_124_v67_ = _dafny.Map({(d_46_v7_) + (d_46_v7_): d_39_v0_})
            d_124_v67_ = (d_124_v67_).set((_dafny.SeqWithoutIsStrInference([not((p1)[default__.safeIndex(666, (p1).length(0))])])) + (d_46_v7_), (p1)[default__.safeIndex(666, (p1).length(0))])
            (globalState).f3 = p1
            (globalState).f4 = 31
        elif True:
            d_125_v68_: C1
            nw19_ = C1()
            nw19_.ctor__((d_119_v62_).f11)
            d_125_v68_ = nw19_
            d_126_v69_: _dafny.Map
            d_126_v69_ = _dafny.Map({p0: d_39_v0_})
            d_127_v70_: _dafny.Set
            d_127_v70_ = _dafny.Set({d_126_v69_, d_126_v69_})
            d_127_v70_ = d_127_v70_
            d_128_v71_: D6
            d_128_v71_ = D6_DC13()
            d_129_v72_: _dafny.Seq
            d_129_v72_ = _dafny.SeqWithoutIsStrInference([d_128_v71_, d_128_v71_, d_128_v71_])
            d_130_v73_: D6
            d_130_v73_ = D6_DC14((d_129_v72_)[default__.safeIndex(len(d_40_v1_), len(d_129_v72_))])
            d_130_v73_ = d_130_v73_
            d_40_v1_ = (d_40_v1_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_131_i4_ in range(default__.abs(771))]))
            (globalState).f9 = (p1)[default__.safeIndex(666, (p1).length(0))]
        d_132_v74_: _dafny.Seq
        out0_: _dafny.Seq
        out0_ = (d_119_v62_).m0(len(d_40_v1_), d_40_v1_, d_39_v0_, p0, globalState)
        d_132_v74_ = out0_
        r0 = p1
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_133_v0_: _dafny.Array
        nw20_ = _dafny.Array(False, 14)
        d_133_v0_ = nw20_
        d_134_globalState_: GlobalState
        nw21_ = GlobalState()
        nw21_.ctor__(d_133_v0_, 319, False, d_133_v0_, 250, 888, False, 746, False, False, True)
        d_134_globalState_ = nw21_
        d_135_v1_: bool
        d_135_v1_ = True
        d_136_v2_: _dafny.MultiSet
        d_136_v2_ = _dafny.MultiSet([d_135_v1_])
        d_137_v3_: _dafny.Map
        d_137_v3_ = _dafny.Map({(d_136_v2_) | (d_136_v2_): d_135_v1_})
        d_137_v3_ = (d_137_v3_) | (d_137_v3_)
        d_138_v4_: int
        d_138_v4_ = 384
        hi0_ = ((0) - (d_138_v4_)) - (d_138_v4_)
        for d_139_i0_ in range(d_138_v4_, hi0_):
            d_140_v5_: _dafny.Seq
            d_140_v5_ = _dafny.SeqWithoutIsStrInference([d_135_v1_])
            d_135_v1_ = (d_135_v1_ if default__.fm0(d_135_v1_, d_135_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfjvu"))), len(d_140_v5_), d_134_globalState_) else default__.fm0(d_135_v1_, not(True), d_138_v4_, d_139_i0_, d_134_globalState_))
            d_141_v6_: _dafny.Seq
            d_141_v6_ = _dafny.SeqWithoutIsStrInference([d_133_v0_])
            d_141_v6_ = _dafny.SeqWithoutIsStrInference([d_133_v0_, d_133_v0_])
            d_142_v7_: _dafny.Array
            def lambda4_(d_143_i0_):
                def lambda5_(d_144_i1_):
                    return _dafny.Map({d_143_i0_: _dafny.CodePoint('g')})

                return lambda5_

            init2_ = lambda4_(d_139_i0_)
            nw22_ = _dafny.Array(None, 23)
            for i0_2_ in range(nw22_.length(0)):
                nw22_[i0_2_] = init2_(i0_2_)
            d_142_v7_ = nw22_
            d_145_v8_: str
            d_145_v8_ = _dafny.CodePoint('o')
            d_146_v9_: _dafny.Map
            d_146_v9_ = _dafny.Map({d_138_v4_: d_145_v8_})
            index19_ = default__.safeIndex(840, (d_142_v7_).length(0))
            (d_142_v7_)[index19_] = d_146_v9_
            d_147_v10_: bool
            d_147_v10_ = False
            index20_ = default__.safeIndex(840, (d_142_v7_).length(0))
            (d_142_v7_)[index20_] = _dafny.Map({(d_139_i0_) - (len(default__.fm1(_dafny.Map({d_139_i0_: d_138_v4_}), (d_147_v10_), d_139_i0_, d_138_v4_, d_134_globalState_))): d_145_v8_})
            d_148_v11_: _dafny.Seq
            d_148_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewk"))
            d_148_v11_ = d_148_v11_
        d_149_v12_: _dafny.Map
        d_149_v12_ = _dafny.Map({(d_138_v4_) - (261): not(default__.fm0(d_135_v1_, d_135_v1_, d_138_v4_, 789, d_134_globalState_))})
        d_149_v12_ = d_149_v12_
        d_150_v13_: _dafny.Map
        d_150_v13_ = _dafny.Map({d_135_v1_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "du"))})
        d_151_v14_: _dafny.Seq
        d_151_v14_ = _dafny.SeqWithoutIsStrInference([d_138_v4_, default__.safeModuloInt((0) - (len(d_150_v13_)), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_152_i2_ in range(default__.abs(671))])))])
        d_153_v15_: bool
        d_153_v15_ = d_135_v1_
        d_154_v16_: _dafny.Map
        d_154_v16_ = _dafny.Map({d_135_v1_: d_138_v4_})
        d_155_v17_: _dafny.Seq
        d_155_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajftdewh"))
        (d_134_globalState_).f4 = (d_151_v14_)[default__.safeIndex(default__.fm2(d_135_v1_, ((d_154_v16_)[d_153_v15_] if (d_153_v15_) in (d_154_v16_) else d_138_v4_), d_138_v4_, d_155_v17_, d_134_globalState_), len(d_151_v14_))]
        d_156_v18_: _dafny.Map
        d_156_v18_ = _dafny.Map({d_135_v1_: (False if d_135_v1_ else d_135_v1_)})
        d_156_v18_ = (d_156_v18_).set(d_135_v1_, d_135_v1_)
        source3_ = d_153_v15_
        d_157___mcc_h0_ = source3_
        d_158_cf0_ = d_157___mcc_h0_
        d_159_v19_: str
        d_159_v19_ = _dafny.CodePoint('t')
        d_160_v20_: C0
        nw23_ = C0()
        nw23_.ctor__(d_159_v19_)
        d_160_v20_ = nw23_
        d_161_v21_: _dafny.Seq
        d_161_v21_ = _dafny.SeqWithoutIsStrInference([d_155_v17_])
        d_162_v22_: _dafny.Array
        nw24_ = _dafny.Array(None, 18)
        nw24_[int(0)] = (d_155_v17_) + (d_155_v17_)
        nw24_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwmdrfli"))
        nw24_[int(2)] = d_155_v17_
        nw24_[int(3)] = (d_161_v21_)[default__.safeIndex(d_138_v4_, len(d_161_v21_))]
        nw24_[int(4)] = _dafny.SeqWithoutIsStrInference([d_159_v19_ for d_163_i3_ in range(default__.abs(101))])
        nw24_[int(5)] = d_155_v17_
        nw24_[int(6)] = ((d_150_v13_)[default__.fm0(d_135_v1_, d_135_v1_, d_138_v4_, d_138_v4_, d_134_globalState_)] if (default__.fm0(d_135_v1_, d_135_v1_, d_138_v4_, d_138_v4_, d_134_globalState_)) in (d_150_v13_) else d_155_v17_)
        nw24_[int(7)] = (_dafny.SeqWithoutIsStrInference([d_159_v19_ for d_164_i4_ in range(default__.abs(195))])) + (d_155_v17_)
        nw24_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krisnrwt"))
        nw24_[int(9)] = (d_155_v17_) + (_dafny.SeqWithoutIsStrInference([(d_160_v20_).f12 for d_165_i5_ in range(default__.abs(496))]))
        nw24_[int(10)] = d_155_v17_
        nw24_[int(11)] = _dafny.SeqWithoutIsStrInference([(d_160_v20_).f12 for d_166_i6_ in range(default__.abs(649))])
        nw24_[int(12)] = (d_155_v17_) + (d_155_v17_)
        nw24_[int(13)] = d_155_v17_
        nw24_[int(14)] = d_155_v17_
        nw24_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkm"))
        nw24_[int(16)] = default__.fm4((0) - (d_138_v4_), d_158_cf0_, d_158_cf0_, (0) - (d_138_v4_), d_134_globalState_)
        nw24_[int(17)] = ((d_155_v17_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfkthhxah")))).set(default__.safeIndex(d_138_v4_, len((d_155_v17_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfkthhxah"))))), d_159_v19_)
        d_162_v22_ = nw24_
        index21_ = default__.safeIndex(280, (d_162_v22_).length(0))
        (d_162_v22_)[index21_] = d_155_v17_
        index22_ = default__.safeIndex(280, (d_162_v22_).length(0))
        (d_162_v22_)[index22_] = d_155_v17_
        d_167_v23_: _dafny.Array
        nw25_ = _dafny.Array(_dafny.CodePoint('D'), 10)
        d_167_v23_ = nw25_
        index23_ = default__.safeIndex(16, (d_167_v23_).length(0))
        (d_167_v23_)[index23_] = (d_160_v20_).f12
        index24_ = default__.safeIndex(16, (d_167_v23_).length(0))
        (d_167_v23_)[index24_] = d_159_v19_
        (d_134_globalState_).f10 = (not (not(d_135_v1_)) or (d_135_v1_)) and (d_135_v1_)
        (d_134_globalState_).f10 = d_135_v1_
        d_168_v24_: _dafny.Array
        def lambda6_(d_169_v17_):
            def lambda7_(d_170_i8_):
                return (d_170_i8_) - (len(d_169_v17_))

            return lambda7_

        init3_ = lambda6_(d_155_v17_)
        nw26_ = _dafny.Array(None, 29)
        for i0_3_ in range(nw26_.length(0)):
            nw26_[i0_3_] = init3_(i0_3_)
        d_168_v24_ = nw26_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_168_v24_).length(0)):
            d_171_i7_: int = guard_loop_0_
            if (True) and (((0) <= (d_171_i7_)) and ((d_171_i7_) < ((d_168_v24_).length(0)))):
                (d_168_v24_)[(d_171_i7_)] = (d_171_i7_) * (d_138_v4_)
        if d_135_v1_:
            index25_ = default__.safeIndex(891, (d_168_v24_).length(0))
            (d_168_v24_)[index25_] = (0) - (d_138_v4_)
            d_172_v25_: _dafny.Seq
            d_172_v25_ = _dafny.SeqWithoutIsStrInference([d_135_v1_, d_135_v1_])
            index26_ = default__.safeIndex(891, (d_168_v24_).length(0))
            rhs15_ = (d_138_v4_) - (d_138_v4_)
            rhs16_ = (d_135_v1_) not in ((d_172_v25_).set(default__.safeIndex(d_138_v4_, len(d_172_v25_)), d_135_v1_))
            rhs17_ = default__.safeDivisionInt(d_138_v4_, len(default__.fm11((0) - (d_138_v4_), d_134_globalState_)))
            rhs18_ = d_135_v1_
            lhs14_ = d_168_v24_
            lhs15_ = default__.safeIndex(891, (d_168_v24_).length(0))
            lhs16_ = d_134_globalState_
            lhs17_ = d_134_globalState_
            lhs14_[lhs15_] = rhs15_
            d_135_v1_ = rhs16_
            lhs16_.f4 = rhs17_
            lhs17_.f9 = rhs18_
            source4_ = default__.fm12(d_138_v4_, d_134_globalState_)
            if source4_.is_DC3:
                d_173___mcc_h1_ = source4_.cf3
                d_174___mcc_h2_ = source4_.cf4
                d_175_cf4_ = d_174___mcc_h2_
                d_176_cf3_ = d_173___mcc_h1_
                d_177_v26_: _dafny.Set
                d_177_v26_ = _dafny.Set({((d_149_v12_)[d_138_v4_] if (d_138_v4_) in (d_149_v12_) else d_176_cf3_)})
                (d_134_globalState_).f9 = (d_177_v26_) != (d_177_v26_)
                d_151_v14_ = d_151_v14_
                index27_ = default__.safeIndex(891, (d_168_v24_).length(0))
                (d_168_v24_)[index27_] = len(default__.fm11((d_168_v24_)[default__.safeIndex(891, (d_168_v24_).length(0))], d_134_globalState_))
                d_178_v27_: D2
                d_178_v27_ = D2_DC3((d_172_v25_)[default__.safeIndex((d_168_v24_)[default__.safeIndex(891, (d_168_v24_).length(0))], len(d_172_v25_))], False)
                d_179_v28_: C0
                nw27_ = C0()
                nw27_.ctor__((_dafny.CodePoint('l') if (d_178_v27_).cf4 else _dafny.CodePoint('k')))
                d_179_v28_ = nw27_
            elif source4_.is_DC4:
                d_180___mcc_h3_ = source4_.cf5
                d_181___mcc_h4_ = source4_.cf6
                d_182_cf6_ = d_181___mcc_h4_
                d_183_cf5_ = d_180___mcc_h3_
                d_184_v29_: C1
                nw28_ = C1()
                nw28_.ctor__(d_153_v15_)
                d_184_v29_ = nw28_
                d_184_v29_ = d_184_v29_
                d_185_v30_: _dafny.Array
                nw29_ = _dafny.Array(int(0), 18)
                d_185_v30_ = nw29_
                rhs19_ = d_185_v30_
                rhs20_ = d_135_v1_
                rhs21_ = d_172_v25_
                rhs22_ = d_135_v1_
                lhs18_ = d_134_globalState_
                lhs19_ = d_134_globalState_
                d_185_v30_ = rhs19_
                lhs18_.f9 = rhs20_
                d_172_v25_ = rhs21_
                lhs19_.f9 = rhs22_
                d_135_v1_ = d_135_v1_
                d_186_v31_: _dafny.Set
                d_186_v31_ = _dafny.Set({d_136_v2_, d_136_v2_})
                d_187_v32_: _dafny.MultiSet
                d_187_v32_ = _dafny.MultiSet([d_136_v2_, d_136_v2_])
                def iife12_():
                    coll8_ = _dafny.Set()
                    compr_8_: _dafny.MultiSet
                    for compr_8_ in (d_187_v32_).Elements:
                        d_188_v33_: _dafny.MultiSet = compr_8_
                        if (d_188_v33_) in (d_187_v32_):
                            coll8_ = coll8_.union(_dafny.Set([d_188_v33_]))
                    return _dafny.Set(coll8_)
                d_186_v31_ = iife12_()
                
            elif True:
                d_189___mcc_h5_ = source4_.cf2
                d_190_cf2_ = d_189___mcc_h5_
                d_155_v17_ = (d_155_v17_) + (d_155_v17_)
                d_191_v34_: C1
                nw30_ = C1()
                nw30_.ctor__(d_153_v15_)
                d_191_v34_ = nw30_
                d_191_v34_ = d_191_v34_
                d_192_v35_: C0
                nw31_ = C0()
                nw31_.ctor__(d_190_cf2_)
                d_192_v35_ = nw31_
                d_192_v35_ = d_192_v35_
                d_193_v36_: _dafny.Map
                d_193_v36_ = _dafny.Map({d_136_v2_: d_138_v4_})
                d_194_v37_: _dafny.Seq
                out1_: _dafny.Seq
                out1_ = (d_191_v34_).m0(len(d_193_v36_), (d_155_v17_) + (_dafny.SeqWithoutIsStrInference([d_190_cf2_ for d_195_i9_ in range(default__.abs(918))])), not (d_135_v1_) or (d_135_v1_), (d_168_v24_)[default__.safeIndex(891, (d_168_v24_).length(0))], d_134_globalState_)
                d_194_v37_ = out1_
            index28_ = default__.safeIndex(891, (d_168_v24_).length(0))
            (d_168_v24_)[index28_] = (d_168_v24_)[default__.safeIndex(891, (d_168_v24_).length(0))]
            d_196_v38_: str
            d_196_v38_ = _dafny.CodePoint('f')
            d_197_v39_: D2
            d_197_v39_ = D2_DC3((d_155_v17_) <= ((d_155_v17_).set(default__.safeIndex(d_138_v4_, len(d_155_v17_)), d_196_v38_)), (d_155_v17_) < ((d_155_v17_).set(default__.safeIndex(d_138_v4_, len(d_155_v17_)), d_196_v38_)))
            d_197_v39_ = d_197_v39_
            d_198_v40_: _dafny.Array
            nw32_ = _dafny.Array(int(0), 25)
            d_198_v40_ = nw32_
            d_198_v40_ = d_198_v40_
        elif True:
            d_199_v41_: _dafny.Set
            d_199_v41_ = _dafny.Set({d_155_v17_, d_155_v17_, d_155_v17_})
            d_199_v41_ = default__.fm13(839, d_138_v4_, d_155_v17_, d_135_v1_, d_134_globalState_)
            d_200_v42_: str
            d_200_v42_ = _dafny.CodePoint('w')
            d_201_v43_: C0
            nw33_ = C0()
            nw33_.ctor__(d_200_v42_)
            d_201_v43_ = nw33_
            d_202_v44_: _dafny.Seq
            d_202_v44_ = _dafny.SeqWithoutIsStrInference([d_201_v43_])
            d_203_v45_: _dafny.Array
            nw34_ = _dafny.Array(None, 21)
            nw34_[int(0)] = d_201_v43_
            nw34_[int(1)] = d_201_v43_
            nw34_[int(2)] = d_201_v43_
            nw34_[int(3)] = d_201_v43_
            nw34_[int(4)] = d_201_v43_
            nw34_[int(5)] = d_201_v43_
            nw34_[int(6)] = (d_202_v44_)[default__.safeIndex(d_138_v4_, len(d_202_v44_))]
            nw34_[int(7)] = d_201_v43_
            nw34_[int(8)] = d_201_v43_
            nw34_[int(9)] = d_201_v43_
            nw34_[int(10)] = d_201_v43_
            nw34_[int(11)] = d_201_v43_
            nw34_[int(12)] = d_201_v43_
            nw34_[int(13)] = d_201_v43_
            nw34_[int(14)] = d_201_v43_
            nw34_[int(15)] = d_201_v43_
            nw34_[int(16)] = d_201_v43_
            nw34_[int(17)] = d_201_v43_
            nw34_[int(18)] = d_201_v43_
            nw34_[int(19)] = d_201_v43_
            nw34_[int(20)] = d_201_v43_
            d_203_v45_ = nw34_
            d_203_v45_ = d_203_v45_
            (d_134_globalState_).f4 = (default__.fm2(d_135_v1_, 942, d_138_v4_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")), d_134_globalState_)) + (d_138_v4_)
            d_204_v46_: _dafny.Array
            out2_: _dafny.Array
            out2_ = default__.m1(d_138_v4_, d_133_v0_, d_134_globalState_)
            d_204_v46_ = out2_
            d_205_v47_: D4
            d_205_v47_ = D4_DC6(d_204_v46_)
            d_206_v48_: _dafny.Array
            out3_: _dafny.Array
            out3_ = default__.m1(default__.safeDivisionInt((0) - (d_138_v4_), d_138_v4_), (d_205_v47_).cf8, d_134_globalState_)
            d_206_v48_ = out3_
        d_207_v49_: C1
        nw35_ = C1()
        nw35_.ctor__(d_153_v15_)
        d_207_v49_ = nw35_
        d_207_v49_ = d_207_v49_
        d_208_v50_: _dafny.Seq
        d_208_v50_ = _dafny.SeqWithoutIsStrInference([d_135_v1_, d_135_v1_])
        d_135_v1_ = (d_208_v50_) != (d_208_v50_)
        if d_135_v1_:
            (d_134_globalState_).f10 = d_135_v1_
            d_209_v52_: _dafny.Set
            def iife13_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(-52, -65):
                    d_210_v51_: int = compr_9_
                    if ((-52) <= (d_210_v51_)) and ((d_210_v51_) < (-65)):
                        coll9_[(d_210_v51_) + (d_138_v4_)] = len(_dafny.Map({False: len(d_155_v17_)}))
                return _dafny.Map(coll9_)
            d_209_v52_ = _dafny.Set({(len(iife13_()
            )) * (d_138_v4_)})
            d_209_v52_ = (d_209_v52_).intersection(d_209_v52_)
            d_211_v53_: str
            d_211_v53_ = _dafny.CodePoint('y')
            d_212_v54_: D2
            d_212_v54_ = D2_DC4(d_138_v4_, d_211_v53_)
            d_213_v55_: _dafny.Map
            d_213_v55_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([14]): (d_138_v4_ if d_135_v1_ else (d_212_v54_).cf5)})
            d_213_v55_ = (d_213_v55_).set((((_dafny.SeqWithoutIsStrInference([d_138_v4_])).set(default__.safeIndex(d_138_v4_, len(_dafny.SeqWithoutIsStrInference([d_138_v4_]))), d_138_v4_)).set(default__.safeIndex(d_138_v4_, len((_dafny.SeqWithoutIsStrInference([d_138_v4_])).set(default__.safeIndex(d_138_v4_, len(_dafny.SeqWithoutIsStrInference([d_138_v4_]))), d_138_v4_))), d_138_v4_)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_135_v1_])), len(((_dafny.SeqWithoutIsStrInference([d_138_v4_])).set(default__.safeIndex(d_138_v4_, len(_dafny.SeqWithoutIsStrInference([d_138_v4_]))), d_138_v4_)).set(default__.safeIndex(d_138_v4_, len((_dafny.SeqWithoutIsStrInference([d_138_v4_])).set(default__.safeIndex(d_138_v4_, len(_dafny.SeqWithoutIsStrInference([d_138_v4_]))), d_138_v4_))), d_138_v4_))), len(_dafny.Map({d_135_v1_: d_135_v1_}))), -114)
            d_214_v56_: C0
            nw36_ = C0()
            nw36_.ctor__(d_211_v53_)
            d_214_v56_ = nw36_
            d_215_v57_: _dafny.Array
            def lambda8_(d_216_v14_, d_217_v4_):
                def lambda9_(d_218_i10_):
                    return (d_216_v14_).set(default__.safeIndex(d_217_v4_, len(d_216_v14_)), d_217_v4_)

                return lambda9_

            init4_ = lambda8_(d_151_v14_, d_138_v4_)
            nw37_ = _dafny.Array(None, 11)
            for i0_4_ in range(nw37_.length(0)):
                nw37_[i0_4_] = init4_(i0_4_)
            d_215_v57_ = nw37_
            d_219_v58_: _dafny.Seq
            d_219_v58_ = _dafny.SeqWithoutIsStrInference([d_215_v57_])
            d_220_v59_: _dafny.Array
            nw38_ = _dafny.Array(None, 18)
            nw38_[int(0)] = d_215_v57_
            nw38_[int(1)] = d_215_v57_
            nw38_[int(2)] = d_215_v57_
            nw38_[int(3)] = (d_219_v58_)[default__.safeIndex(316, len(d_219_v58_))]
            nw38_[int(4)] = d_215_v57_
            nw38_[int(5)] = d_215_v57_
            nw38_[int(6)] = d_215_v57_
            nw38_[int(7)] = d_215_v57_
            nw38_[int(8)] = d_215_v57_
            nw38_[int(9)] = d_215_v57_
            nw38_[int(10)] = d_215_v57_
            nw38_[int(11)] = d_215_v57_
            nw38_[int(12)] = d_215_v57_
            nw38_[int(13)] = (d_215_v57_ if not(d_135_v1_) else d_215_v57_)
            nw38_[int(14)] = d_215_v57_
            nw38_[int(15)] = d_215_v57_
            nw38_[int(16)] = d_215_v57_
            nw38_[int(17)] = d_215_v57_
            d_220_v59_ = nw38_
            index29_ = default__.safeIndex(757, (d_220_v59_).length(0))
            (d_220_v59_)[index29_] = d_215_v57_
            d_221_v60_: _dafny.Array
            nw39_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_221_v60_ = nw39_
            d_222_v61_: _dafny.Map
            d_222_v61_ = _dafny.Map({(0) - (d_138_v4_): d_214_v56_})
            index30_ = default__.safeIndex(757, (d_220_v59_).length(0))
            rhs23_ = ((d_222_v61_)[len(default__.fm8(d_135_v1_, d_155_v17_, d_134_globalState_))] if (len(default__.fm8(d_135_v1_, d_155_v17_, d_134_globalState_))) in (d_222_v61_) else (d_214_v56_ if d_135_v1_ else d_214_v56_))
            rhs24_ = ((d_156_v18_)[(d_135_v1_ if d_135_v1_ else d_135_v1_)] if ((d_135_v1_ if d_135_v1_ else d_135_v1_)) in (d_156_v18_) else d_135_v1_)
            rhs25_ = d_215_v57_
            rhs26_ = d_221_v60_
            rhs27_ = d_138_v4_
            lhs20_ = d_134_globalState_
            lhs21_ = d_220_v59_
            lhs22_ = default__.safeIndex(757, (d_220_v59_).length(0))
            lhs23_ = d_134_globalState_
            d_214_v56_ = rhs23_
            lhs20_.f10 = rhs24_
            lhs21_[lhs22_] = rhs25_
            d_221_v60_ = rhs26_
            lhs23_.f4 = rhs27_
            d_223_v62_: C1
            nw40_ = C1()
            nw40_.ctor__((d_207_v49_).f11)
            d_223_v62_ = nw40_
        elif True:
            d_224_v63_: _dafny.Array
            nw41_ = _dafny.Array(None, 13)
            nw41_[int(0)] = d_133_v0_
            nw41_[int(1)] = d_133_v0_
            nw41_[int(2)] = d_133_v0_
            nw41_[int(3)] = d_133_v0_
            nw41_[int(4)] = d_133_v0_
            nw41_[int(5)] = d_133_v0_
            nw41_[int(6)] = d_133_v0_
            nw41_[int(7)] = d_133_v0_
            nw41_[int(8)] = d_133_v0_
            nw41_[int(9)] = d_133_v0_
            nw41_[int(10)] = d_133_v0_
            nw41_[int(11)] = d_133_v0_
            nw41_[int(12)] = d_133_v0_
            d_224_v63_ = nw41_
            d_224_v63_ = d_224_v63_
            d_225_v64_: _dafny.Seq
            d_225_v64_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfb"))])
            d_226_v65_: _dafny.Map
            d_226_v65_ = d_149_v12_
            if default__.fm0(d_135_v1_, (d_155_v17_) < ((d_225_v64_)[default__.safeIndex(d_138_v4_, len(d_225_v64_))]), -756, default__.fm2(default__.fm0(d_135_v1_, d_135_v1_, d_138_v4_, d_138_v4_, d_134_globalState_), len(_dafny.SeqWithoutIsStrInference([d_226_v65_])), len(d_208_v50_), d_155_v17_, d_134_globalState_), d_134_globalState_):
                d_227_v66_: str
                d_227_v66_ = _dafny.CodePoint('w')
                d_228_v67_: C0
                nw42_ = C0()
                nw42_.ctor__(d_227_v66_)
                d_228_v67_ = nw42_
                (d_134_globalState_).f4 = d_138_v4_
                d_229_v69_: _dafny.Array
                def lambda10_(d_230_v4_):
                    def lambda11_(d_231_i11_):
                        def iife14_():
                            coll10_ = _dafny.Map()
                            compr_10_: int
                            for compr_10_ in _dafny.IntegerRange(-874, 754):
                                d_232_v68_: int = compr_10_
                                if ((-874) <= (d_232_v68_)) and ((d_232_v68_) < (754)):
                                    coll10_[default__.safeModuloInt(d_232_v68_, d_230_v4_)] = d_230_v4_
                            return _dafny.Map(coll10_)
                        return (_dafny.Map({d_230_v4_: d_230_v4_})) | (iife14_()
                        )

                    return lambda11_

                init5_ = lambda10_(d_138_v4_)
                nw43_ = _dafny.Array(None, 20)
                for i0_5_ in range(nw43_.length(0)):
                    nw43_[i0_5_] = init5_(i0_5_)
                d_229_v69_ = nw43_
                d_233_v70_: _dafny.Map
                d_233_v70_ = _dafny.Map({d_138_v4_: d_138_v4_})
                index31_ = default__.safeIndex(565, (d_229_v69_).length(0))
                (d_229_v69_)[index31_] = d_233_v70_
                d_234_v71_: _dafny.Seq
                d_234_v71_ = _dafny.SeqWithoutIsStrInference([d_207_v49_])
                d_235_v72_: _dafny.Map
                d_235_v72_ = _dafny.Map({d_138_v4_: d_133_v0_})
                d_236_v73_: D5
                d_236_v73_ = D5_DC8(d_235_v72_)
                d_237_v74_: _dafny.Map
                d_237_v74_ = _dafny.Map({_dafny.MultiSet(d_234_v71_): len(((d_236_v73_).cf11) | (d_235_v72_))})
                d_238_v77_: _dafny.Map
                d_238_v77_ = _dafny.Map({(d_228_v67_).f12: d_135_v1_})
                d_239_v78_: _dafny.Map
                def iife15_():
                    coll11_ = _dafny.Map()
                    compr_11_: str
                    for compr_11_ in (d_238_v77_).keys.Elements:
                        d_240_v76_: str = compr_11_
                        if (d_240_v76_) in (d_238_v77_):
                            coll11_[d_240_v76_] = d_138_v4_
                    return _dafny.Map(coll11_)
                d_239_v78_ = _dafny.Map({d_136_v2_: iife15_()
                })
                d_241_v79_: _dafny.Set
                d_241_v79_ = _dafny.Set({len(d_155_v17_), d_138_v4_, d_138_v4_})
                d_242_v80_: _dafny.Array
                nw44_ = _dafny.Array(None, 5)
                d_242_v80_ = nw44_
                d_243_v81_: _dafny.Map
                d_243_v81_ = _dafny.Map({d_135_v1_: d_242_v80_})
                d_244_v82_: _dafny.Map
                d_244_v82_ = _dafny.Map({d_243_v81_: d_237_v74_})
                d_245_v83_: _dafny.MultiSet
                d_245_v83_ = _dafny.MultiSet([d_207_v49_])
                index32_ = default__.safeIndex(565, (d_229_v69_).length(0))
                def iife16_():
                    coll12_ = _dafny.Map()
                    compr_12_: _dafny.MultiSet
                    for compr_12_ in (d_239_v78_).keys.Elements:
                        d_246_v75_: _dafny.MultiSet = compr_12_
                        if (d_246_v75_) in (d_239_v78_):
                            coll12_[d_246_v75_] = d_135_v1_
                    return _dafny.Map(coll12_)
                rhs28_ = d_233_v70_
                rhs29_ = (_dafny.Map({d_136_v2_: d_135_v1_})) | ((iife16_()
                ) | (default__.fm14(True, d_138_v4_, d_241_v79_, d_134_globalState_)))
                rhs30_ = ((d_244_v82_)[d_243_v81_] if (d_243_v81_) in (d_244_v82_) else (d_237_v74_).set(d_245_v83_, d_138_v4_))
                lhs24_ = d_229_v69_
                lhs25_ = default__.safeIndex(565, (d_229_v69_).length(0))
                lhs24_[lhs25_] = rhs28_
                d_137_v3_ = rhs29_
                d_237_v74_ = rhs30_
                d_247_v84_: _dafny.Array
                out4_: _dafny.Array
                out4_ = default__.m1(d_138_v4_, d_133_v0_, d_134_globalState_)
                d_247_v84_ = out4_
                d_248_v85_: D2
                d_248_v85_ = D2_DC2((d_228_v67_).f12)
                d_249_v86_: _dafny.MultiSet
                d_249_v86_ = _dafny.MultiSet([d_248_v85_])
                (d_134_globalState_).f10 = not((d_249_v86_).ispropersubset((d_249_v86_).intersection((d_249_v86_).set(d_248_v85_, default__.abs(d_138_v4_)))))
            elif True:
                d_250_v87_: _dafny.Map
                d_250_v87_ = _dafny.Map({(d_135_v1_ if d_135_v1_ else d_135_v1_): d_138_v4_})
                d_251_v88_: str
                d_251_v88_ = _dafny.CodePoint('n')
                rhs31_ = not(d_135_v1_)
                rhs32_ = d_138_v4_
                rhs33_ = (d_155_v17_) == ((d_155_v17_).set(default__.safeIndex(d_138_v4_, len(d_155_v17_)), d_251_v88_))
                rhs34_ = d_168_v24_
                rhs35_ = (d_250_v87_) | (_dafny.Map({d_135_v1_: d_138_v4_}))
                lhs26_ = d_134_globalState_
                lhs27_ = d_134_globalState_
                lhs28_ = d_134_globalState_
                lhs26_.f10 = rhs31_
                lhs27_.f4 = rhs32_
                lhs28_.f10 = rhs33_
                d_168_v24_ = rhs34_
                d_250_v87_ = rhs35_
                d_252_v89_: _dafny.Array
                nw45_ = _dafny.Array(None, 4)
                d_252_v89_ = nw45_
                index33_ = default__.safeIndex(264, (d_252_v89_).length(0))
                (d_252_v89_)[index33_] = (d_207_v49_ if d_135_v1_ else d_207_v49_)
                index34_ = default__.safeIndex(183, (d_168_v24_).length(0))
                (d_168_v24_)[index34_] = len(default__.fm15(d_134_globalState_))
                d_253_v90_: _dafny.Set
                d_253_v90_ = _dafny.Set({d_133_v0_})
                index35_ = default__.safeIndex(264, (d_252_v89_).length(0))
                index36_ = default__.safeIndex(183, (d_168_v24_).length(0))
                rhs36_ = d_207_v49_
                rhs37_ = default__.fm0(d_135_v1_, d_135_v1_, d_138_v4_, d_138_v4_, d_134_globalState_)
                rhs38_ = default__.safeModuloInt(default__.fm2(d_135_v1_, d_138_v4_, d_138_v4_, d_155_v17_, d_134_globalState_), (len(_dafny.SeqWithoutIsStrInference([d_138_v4_ for d_254_i12_ in range(default__.abs(-324))]))) * (d_138_v4_))
                rhs39_ = d_253_v90_
                rhs40_ = True
                lhs29_ = d_252_v89_
                lhs30_ = default__.safeIndex(264, (d_252_v89_).length(0))
                lhs31_ = d_168_v24_
                lhs32_ = default__.safeIndex(183, (d_168_v24_).length(0))
                lhs33_ = d_134_globalState_
                lhs29_[lhs30_] = rhs36_
                d_135_v1_ = rhs37_
                lhs31_[lhs32_] = rhs38_
                d_253_v90_ = rhs39_
                lhs33_.f10 = rhs40_
                d_255_v91_: _dafny.Array
                out5_: _dafny.Array
                out5_ = default__.m1((d_168_v24_)[default__.safeIndex(183, (d_168_v24_).length(0))], d_133_v0_, d_134_globalState_)
                d_255_v91_ = out5_
                d_256_v92_: _dafny.Seq
                d_256_v92_ = _dafny.SeqWithoutIsStrInference([d_133_v0_, d_255_v91_, d_255_v91_, d_133_v0_])
                d_257_v93_: _dafny.Array
                out6_: _dafny.Array
                out6_ = default__.m1(len(d_155_v17_), (d_256_v92_)[default__.safeIndex(d_138_v4_, len(d_256_v92_))], d_134_globalState_)
                d_257_v93_ = out6_
                index37_ = default__.safeIndex(183, (d_168_v24_).length(0))
                (d_168_v24_)[index37_] = 412
            d_138_v4_ = (d_151_v14_)[default__.safeIndex(d_138_v4_, len(d_151_v14_))]
            d_258_v94_: _dafny.MultiSet
            d_258_v94_ = _dafny.MultiSet([d_138_v4_, d_138_v4_])
            d_259_v95_: _dafny.Seq
            d_259_v95_ = _dafny.SeqWithoutIsStrInference([d_258_v94_])
            d_260_v96_: _dafny.Map
            d_260_v96_ = _dafny.Map({d_259_v95_: not(not((d_135_v1_ if d_135_v1_ else d_135_v1_)))})
            d_260_v96_ = (d_260_v96_).set(d_259_v95_, d_135_v1_)
            d_138_v4_ = d_138_v4_
        d_138_v4_ = d_138_v4_
        (d_134_globalState_).f4 = d_138_v4_
        if d_135_v1_:
            d_261_v97_: _dafny.Array
            out7_: _dafny.Array
            out7_ = default__.m1(d_138_v4_, d_133_v0_, d_134_globalState_)
            d_261_v97_ = out7_
            index38_ = default__.safeIndex(144, (d_168_v24_).length(0))
            (d_168_v24_)[index38_] = d_138_v4_
            index39_ = default__.safeIndex(144, (d_168_v24_).length(0))
            (d_168_v24_)[index39_] = d_138_v4_
            (d_134_globalState_).f10 = d_135_v1_
            d_262_v98_: C1
            nw46_ = C1()
            nw46_.ctor__((d_207_v49_).f11)
            d_262_v98_ = nw46_
            (d_134_globalState_).f4 = (d_138_v4_ if d_135_v1_ else d_138_v4_)
        elif True:
            d_263_v99_: str
            d_263_v99_ = _dafny.CodePoint('r')
            d_264_v100_: _dafny.Map
            d_264_v100_ = _dafny.Map({default__.fm5(d_138_v4_, d_134_globalState_): 692})
            d_265_v101_: _dafny.Map
            d_265_v101_ = _dafny.Map({d_133_v0_: d_156_v18_})
            d_266_v102_: _dafny.Map
            d_266_v102_ = _dafny.Map({(_dafny.Map({d_263_v99_: -509}) if d_135_v1_ else d_264_v100_): len(d_265_v101_)})
            d_266_v102_ = (d_266_v102_).set(d_264_v100_, d_138_v4_)
            (d_134_globalState_).f4 = d_138_v4_
            (d_134_globalState_).f3 = d_133_v0_
            d_267_v103_: _dafny.Seq
            out8_: _dafny.Seq
            out8_ = (d_207_v49_).m0(d_138_v4_, (d_155_v17_) + (d_155_v17_), not(d_135_v1_), d_138_v4_, d_134_globalState_)
            d_267_v103_ = out8_
            (d_134_globalState_).f10 = ((d_207_v49_).f11)
        (d_134_globalState_).f4 = default__.fm2(d_135_v1_, d_138_v4_, d_138_v4_, d_155_v17_, d_134_globalState_)
        _dafny.print(_dafny.string_of((d_133_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_globalState_).f0)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f3)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_134_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_134_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_134_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_135_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v2_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v3_) == (_dafny.Map({_dafny.MultiSet([True]): False, _dafny.MultiSet([False]): False, _dafny.MultiSet([True, False]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v12_) == (_dafny.Map({123: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_150_v13_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "du"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v14_) == (_dafny.SeqWithoutIsStrInference([384, 670]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v15_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v16_) == (_dafny.Map({True: 384}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_155_v17_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_v18_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v24_)[28]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_207_v49_).f11)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v50_) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: False
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
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

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
        return lambda: D2_DC3(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D2_DC2)

class D2_DC3(D2, NamedTuple('DC3', [('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC2(D2, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D3_DC5)

class D3_DC5(D3, NamedTuple('DC5', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC5({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC5) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC7(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D4_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D4_DC6)

class D4_DC7(D4, NamedTuple('DC7', [('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC7({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC7) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC6(D4, NamedTuple('DC6', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC6({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC6) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, None, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D5_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D5_DC8)

class D5_DC9(D5, NamedTuple('DC9', [('cf12', Any), ('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC9({self.cf12.VerbatimString(True)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC9) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC8(D5, NamedTuple('DC8', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC8({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC8) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC11(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), None, int(0), _dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D6_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D6_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D6_DC10)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC11(D6, NamedTuple('DC11', [('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC11({self.cf17.VerbatimString(True)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC11) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC12(D6, NamedTuple('DC12', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC12({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC12) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC13(D6, NamedTuple('DC13', [])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13)
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC10(D6, NamedTuple('DC10', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC10({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC10) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC14(D6, NamedTuple('DC14', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC16(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.MultiSet({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)

class D7_DC16(D7, NamedTuple('DC16', [('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({self.cf25.VerbatimString(True)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC15(D7, NamedTuple('DC15', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC18(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)

class D8_DC18(D8, NamedTuple('DC18', [('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC17(D8, NamedTuple('DC17', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC17({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC17) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f3: _dafny.Array = _dafny.Array(None, 0)
        self.f4: int = int(0)
        self.f9: bool = False
        self.f10: bool = False
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        self._f1: int = int(0)
        self._f2: bool = False
        self._f5: int = int(0)
        self._f6: bool = False
        self._f7: int = int(0)
        self._f8: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self).f9 = f9
        (self).f10 = f10

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

class C0:
    def  __init__(self):
        self._f12: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f12):
        (self)._f12 = f12

    def fm3(self, p0, globalState):
        if (False if (not(True)) else True):
            return False
        elif True:
            return False

    @property
    def f12(self):
        return self._f12

class C1:
    def  __init__(self):
        self._f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f11):
        (self)._f11 = f11

    def m0(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        d_268_v0_: _dafny.Seq
        d_268_v0_ = _dafny.SeqWithoutIsStrInference([len(p1), p3])
        d_269_i0_: int
        d_269_i0_ = 0
        with _dafny.label("0"):
            while (d_268_v0_) < (_dafny.SeqWithoutIsStrInference([(0) - (-33)])):
                with _dafny.c_label("0"):
                    if (d_269_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_269_i0_ = (d_269_i0_) + (1)
                    (globalState).f4 = (0) - (p3)
                    d_270_v1_: _dafny.Map
                    d_270_v1_ = _dafny.Map({-768: p1})
                    d_270_v1_ = (d_270_v1_).set(p0, p1)
                    (globalState).f10 = default__.fm0(p2, (p2 if p2 else True), p3, default__.safeDivisionInt(p3, p3), globalState)
                    d_271_v2_: _dafny.Array
                    def lambda12_(d_272_p1_):
                        def lambda13_(d_273_i1_):
                            return d_272_p1_

                        return lambda13_

                    init6_ = lambda12_(p1)
                    nw47_ = _dafny.Array(None, 26)
                    for i0_6_ in range(nw47_.length(0)):
                        nw47_[i0_6_] = init6_(i0_6_)
                    d_271_v2_ = nw47_
                    index40_ = default__.safeIndex(105, (d_271_v2_).length(0))
                    (d_271_v2_)[index40_] = p1
                    index41_ = default__.safeIndex(105, (d_271_v2_).length(0))
                    (d_271_v2_)[index41_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqgbnrs"))
                    pass
            pass
        (globalState).f10 = p2
        hi1_ = p3
        for d_274_i2_ in range(p3, hi1_):
            d_275_v3_: str
            d_275_v3_ = _dafny.CodePoint('p')
            d_276_v4_: _dafny.Map
            d_276_v4_ = _dafny.Map({len(_dafny.Map({d_275_v3_: p2})): p2})
            d_277_v5_: _dafny.Seq
            d_277_v5_ = _dafny.SeqWithoutIsStrInference([p1])
            d_278_v6_: _dafny.Array
            nw48_ = _dafny.Array(None, 6)
            nw48_[int(0)] = p3
            nw48_[int(1)] = d_274_i2_
            nw48_[int(2)] = d_274_i2_
            nw48_[int(3)] = default__.fm2(p2, (_dafny.MultiSet([p2, p2, p2])).cardinality, len(d_276_v4_), (d_277_v5_)[default__.safeIndex(len(p1), len(d_277_v5_))], globalState)
            nw48_[int(4)] = p3
            nw48_[int(5)] = (p0) - (p3)
            d_278_v6_ = nw48_
            index42_ = default__.safeIndex(323, (d_278_v6_).length(0))
            (d_278_v6_)[index42_] = d_274_i2_
            d_279_v7_: _dafny.Seq
            d_279_v7_ = _dafny.SeqWithoutIsStrInference([p2])
            index43_ = default__.safeIndex(323, (d_278_v6_).length(0))
            (d_278_v6_)[index43_] = (0) - ((len(d_279_v7_)) - (d_274_i2_))
            if p2:
                d_280_v8_: _dafny.MultiSet
                d_280_v8_ = _dafny.MultiSet([866, d_274_i2_])
                (globalState).f4 = ((d_278_v6_)[default__.safeIndex(323, (d_278_v6_).length(0))]) * ((0) - (((0) - (((d_280_v8_)[len((p1).set(default__.safeIndex(p0, len(p1)), d_275_v3_))] if (len((p1).set(default__.safeIndex(p0, len(p1)), d_275_v3_))) in (d_280_v8_) else d_274_i2_))) * ((0) - (p3))))
                index44_ = default__.safeIndex(323, (d_278_v6_).length(0))
                (d_278_v6_)[index44_] = p3
                (globalState).f9 = not (p2) or (default__.fm0(p2, False, (d_268_v0_)[default__.safeIndex((d_278_v6_)[default__.safeIndex(323, (d_278_v6_).length(0))], len(d_268_v0_))], (0) - (len(d_279_v7_)), globalState))
                index45_ = default__.safeIndex(323, (d_278_v6_).length(0))
                (d_278_v6_)[index45_] = (((d_280_v8_)[len(p1)] if (len(p1)) in (d_280_v8_) else (0) - (p3))) * ((d_268_v0_)[default__.safeIndex(p3, len(d_268_v0_))])
                d_281_v9_: _dafny.Map
                d_281_v9_ = _dafny.Map({p2: d_268_v0_})
                (globalState).f10 = (((d_281_v9_)[p2] if (p2) in (d_281_v9_) else d_268_v0_)) == (d_268_v0_)
            elif True:
                d_282_v10_: C0
                nw49_ = C0()
                nw49_.ctor__(d_275_v3_)
                d_282_v10_ = nw49_
                d_283_v11_: _dafny.Array
                nw50_ = _dafny.Array(_dafny.Map({}), 6)
                d_283_v11_ = nw50_
                index46_ = default__.safeIndex(209, (d_283_v11_).length(0))
                (d_283_v11_)[index46_] = _dafny.Map({d_274_i2_: p2})
                d_284_v12_: _dafny.Map
                d_284_v12_ = _dafny.Map({d_274_i2_: p2})
                d_285_v13_: _dafny.Seq
                d_285_v13_ = _dafny.SeqWithoutIsStrInference([d_276_v4_, ((d_284_v12_)) | (_dafny.Map({p3: p2})), (d_276_v4_) | (d_276_v4_), d_276_v4_, d_276_v4_])
                index47_ = default__.safeIndex(209, (d_283_v11_).length(0))
                (d_283_v11_)[index47_] = (d_285_v13_)[default__.safeIndex(((d_278_v6_)[default__.safeIndex(323, (d_278_v6_).length(0))]) * (p3), len(d_285_v13_))]
                d_286_v14_: C0
                nw51_ = C0()
                nw51_.ctor__((d_282_v10_).f12)
                d_286_v14_ = nw51_
                d_287_v15_: _dafny.Array
                def lambda14_(d_288_p1_):
                    def lambda15_(d_289_i3_):
                        return d_288_p1_

                    return lambda15_

                init7_ = lambda14_(p1)
                nw52_ = _dafny.Array(None, 13)
                for i0_7_ in range(nw52_.length(0)):
                    nw52_[i0_7_] = init7_(i0_7_)
                d_287_v15_ = nw52_
                index48_ = default__.safeIndex(933, (d_287_v15_).length(0))
                (d_287_v15_)[index48_] = _dafny.SeqWithoutIsStrInference([(D2_DC4(p3, d_275_v3_)).cf6 for d_290_i4_ in range(default__.abs(-706))])
                d_291_v16_: _dafny.Set
                d_291_v16_ = _dafny.Set({p3})
                d_292_v17_: _dafny.Seq
                d_292_v17_ = _dafny.SeqWithoutIsStrInference([d_278_v6_])
                d_293_v18_: _dafny.Map
                d_293_v18_ = _dafny.Map({(d_292_v17_)[default__.safeIndex(p0, len(d_292_v17_))]: len(d_268_v0_)})
                d_294_v19_: _dafny.MultiSet
                d_294_v19_ = _dafny.MultiSet([p2])
                d_295_v20_: _dafny.Map
                d_295_v20_ = _dafny.Map({615: d_274_i2_})
                index49_ = default__.safeIndex(933, (d_287_v15_).length(0))
                rhs41_ = p1
                rhs42_ = (d_291_v16_).issubset(d_291_v16_)
                rhs43_ = (((d_294_v19_)[(d_279_v7_)[default__.safeIndex(len(d_268_v0_), len(d_279_v7_))]] if ((d_279_v7_)[default__.safeIndex(len(d_268_v0_), len(d_279_v7_))]) in (d_294_v19_) else d_274_i2_) if (d_293_v18_) != (d_293_v18_) else ((d_295_v20_)[p0] if (p0) in (d_295_v20_) else p3))
                rhs44_ = not(p2)
                lhs34_ = d_287_v15_
                lhs35_ = default__.safeIndex(933, (d_287_v15_).length(0))
                lhs36_ = globalState
                lhs37_ = globalState
                lhs38_ = globalState
                lhs34_[lhs35_] = rhs41_
                lhs36_.f9 = rhs42_
                lhs37_.f4 = rhs43_
                lhs38_.f9 = rhs44_
                d_296_v21_: _dafny.MultiSet
                d_296_v21_ = _dafny.MultiSet([d_282_v10_, d_282_v10_])
                index50_ = default__.safeIndex(933, (d_287_v15_).length(0))
                (d_287_v15_)[index50_] = default__.fm4((0) - ((d_296_v21_).cardinality), p2, not(p2), ((d_278_v6_)[default__.safeIndex(323, (d_278_v6_).length(0))]) * ((d_278_v6_)[default__.safeIndex(323, (d_278_v6_).length(0))]), globalState)
            d_297_v22_: _dafny.Map
            d_297_v22_ = d_276_v4_
            d_298_v23_: _dafny.Seq
            d_298_v23_ = _dafny.SeqWithoutIsStrInference([d_297_v22_, d_297_v22_])
            (globalState).f10 = ((d_298_v23_ if p2 else d_298_v23_)) <= (d_298_v23_)
            (globalState).f4 = default__.safeModuloInt(default__.safeDivisionInt(default__.fm2(p2, (d_278_v6_)[default__.safeIndex(323, (d_278_v6_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_275_v3_ for d_299_i5_ in range(default__.abs(199))])), p1, globalState), d_274_i2_), (d_278_v6_)[default__.safeIndex(323, (d_278_v6_).length(0))])
        if p2:
            d_300_v24_: C0
            nw53_ = C0()
            nw53_.ctor__(default__.fm5((0) - (len(default__.fm4(p3, p2, p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aroliv"))), globalState))), globalState))
            d_300_v24_ = nw53_
            d_301_v25_: _dafny.Seq
            d_301_v25_ = _dafny.SeqWithoutIsStrInference([d_268_v0_])
            d_301_v25_ = d_301_v25_
            (globalState).f10 = not((False) == ((636) > (p0)))
            d_302_v26_: _dafny.Array
            nw54_ = _dafny.Array(int(0), 26)
            d_302_v26_ = nw54_
            index51_ = default__.safeIndex(933, (d_302_v26_).length(0))
            (d_302_v26_)[index51_] = p0
            d_303_v27_: _dafny.Set
            d_303_v27_ = _dafny.Set({p2, p2, p2, not(p2)})
            d_304_v28_: _dafny.Array
            nw55_ = _dafny.Array(False, 15)
            d_304_v28_ = nw55_
            d_305_v29_: _dafny.Seq
            d_305_v29_ = _dafny.SeqWithoutIsStrInference([p2])
            d_306_v30_: D2
            d_306_v30_ = D2_DC3(p2, p2)
            index52_ = default__.safeIndex(933, (d_302_v26_).length(0))
            rhs45_ = p3
            rhs46_ = (d_303_v27_).issubset(d_303_v27_)
            rhs47_ = (0) - (default__.safeDivisionInt(p0, p0))
            rhs48_ = default__.safeDivisionInt(p0, p0)
            rhs49_ = (len(default__.fm4((_dafny.MultiSet([d_304_v28_])).cardinality, p2, (d_305_v29_)[default__.safeIndex(p0, len(d_305_v29_))], p0, globalState))) > ((p3 if (d_306_v30_).cf4 else p0))
            lhs39_ = d_302_v26_
            lhs40_ = default__.safeIndex(933, (d_302_v26_).length(0))
            lhs41_ = globalState
            lhs42_ = globalState
            lhs43_ = globalState
            lhs44_ = globalState
            lhs39_[lhs40_] = rhs45_
            lhs41_.f9 = rhs46_
            lhs42_.f4 = rhs47_
            lhs43_.f4 = rhs48_
            lhs44_.f9 = rhs49_
            d_307_v31_: _dafny.MultiSet
            d_307_v31_ = _dafny.MultiSet([p2, p2, p2, p2, not(p2)])
            d_308_v32_: _dafny.Map
            d_308_v32_ = _dafny.Map({d_307_v31_: d_300_v24_})
            d_309_v33_: _dafny.Map
            d_309_v33_ = d_308_v32_
            d_310_v34_: _dafny.Map
            d_310_v34_ = _dafny.Map({p2: (d_309_v33_)})
            d_311_v35_: _dafny.Array
            nw56_ = _dafny.Array(None, 8)
            nw56_[int(0)] = ((d_310_v34_)[p2] if (p2) in (d_310_v34_) else d_308_v32_)
            nw56_[int(1)] = _dafny.Map({default__.fm6(True, (d_302_v26_)[default__.safeIndex(933, (d_302_v26_).length(0))], globalState): d_300_v24_})
            nw56_[int(2)] = _dafny.Map({d_307_v31_: d_300_v24_})
            nw56_[int(3)] = d_308_v32_
            nw56_[int(4)] = _dafny.Map({d_307_v31_: d_300_v24_})
            nw56_[int(5)] = d_308_v32_
            nw56_[int(6)] = d_308_v32_
            nw56_[int(7)] = d_308_v32_
            d_311_v35_ = nw56_
            index53_ = default__.safeIndex(352, (d_311_v35_).length(0))
            (d_311_v35_)[index53_] = d_308_v32_
            index54_ = default__.safeIndex(352, (d_311_v35_).length(0))
            (d_311_v35_)[index54_] = (d_309_v33_)
        elif True:
            (globalState).f10 = p2
            d_312_v36_: D2
            d_312_v36_ = D2_DC3(p2, p2)
            source5_ = d_312_v36_
            if source5_.is_DC3:
                d_313___mcc_h0_ = source5_.cf3
                d_314___mcc_h1_ = source5_.cf4
                d_315_cf4_ = d_314___mcc_h1_
                d_316_cf3_ = d_313___mcc_h0_
                d_317_v37_: str
                d_317_v37_ = _dafny.CodePoint('b')
                d_318_v38_: _dafny.Array
                def lambda16_(d_319_p2_):
                    def lambda17_(d_320_i6_):
                        return d_319_p2_

                    return lambda17_

                init8_ = lambda16_(p2)
                nw57_ = _dafny.Array(None, 22)
                for i0_8_ in range(nw57_.length(0)):
                    nw57_[i0_8_] = init8_(i0_8_)
                d_318_v38_ = nw57_
                d_321_v39_: _dafny.Map
                d_321_v39_ = _dafny.Map({d_317_v37_: d_318_v38_})
                d_321_v39_ = _dafny.Map({(default__.fm5(len(d_268_v0_), globalState) if p2 else d_317_v37_): d_318_v38_})
                d_322_v40_: _dafny.Seq
                d_322_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovmhyj"))
                index55_ = default__.safeIndex(415, (d_318_v38_).length(0))
                (d_318_v38_)[index55_] = d_316_cf3_
                d_323_v41_: _dafny.MultiSet
                d_323_v41_ = _dafny.MultiSet([p2])
                index56_ = default__.safeIndex(415, (d_318_v38_).length(0))
                rhs50_ = d_315_cf4_
                rhs51_ = (p1) + ((d_322_v40_) + (p1))
                rhs52_ = ((False) in (d_323_v41_)) == (p2)
                lhs45_ = globalState
                lhs46_ = d_318_v38_
                lhs47_ = default__.safeIndex(415, (d_318_v38_).length(0))
                lhs45_.f10 = rhs50_
                d_322_v40_ = rhs51_
                lhs46_[lhs47_] = rhs52_
                d_324_v42_: _dafny.Array
                nw58_ = _dafny.Array(None, 14)
                nw58_[int(0)] = default__.fm2(not(p2), (0) - (len(_dafny.Set({p0, p0}))), p0, _dafny.SeqWithoutIsStrInference([d_317_v37_ for d_325_i7_ in range(default__.abs(-331))]), globalState)
                nw58_[int(1)] = -31
                nw58_[int(2)] = p3
                nw58_[int(3)] = p3
                nw58_[int(4)] = p3
                nw58_[int(5)] = p3
                nw58_[int(6)] = p3
                nw58_[int(7)] = p0
                nw58_[int(8)] = p3
                nw58_[int(9)] = p0
                nw58_[int(10)] = p0
                nw58_[int(11)] = (0) - (p0)
                nw58_[int(12)] = p3
                nw58_[int(13)] = p3
                d_324_v42_ = nw58_
                d_326_v43_: _dafny.Map
                d_326_v43_ = _dafny.Map({p2: d_324_v42_})
                d_327_v44_: _dafny.Seq
                d_327_v44_ = _dafny.SeqWithoutIsStrInference([d_316_cf3_])
                d_328_v45_: _dafny.Map
                d_328_v45_ = _dafny.Map({d_317_v37_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))})
                d_329_v46_: _dafny.MultiSet
                d_329_v46_ = _dafny.MultiSet([43])
                d_330_v47_: _dafny.Seq
                d_330_v47_ = _dafny.SeqWithoutIsStrInference([False, default__.fm0(default__.fm0(d_316_cf3_, (d_327_v44_)[default__.safeIndex(p0, len(d_327_v44_))], 808, p0, globalState), p2, ((d_328_v45_)[_dafny.CodePoint('r')] if (_dafny.CodePoint('r')) in (d_328_v45_) else 550), (d_329_v46_).cardinality, globalState), d_315_cf4_, d_316_cf3_, p2])
                d_326_v43_ = (d_326_v43_).set((d_327_v44_)[default__.safeIndex(-217, len(d_327_v44_))], d_324_v42_)
                (globalState).f9 = d_316_cf3_
            elif source5_.is_DC4:
                d_331___mcc_h2_ = source5_.cf5
                d_332___mcc_h3_ = source5_.cf6
                d_333_cf6_ = d_332___mcc_h3_
                d_334_cf5_ = d_331___mcc_h2_
                d_335_v48_: D2
                d_335_v48_ = D2_DC4(d_334_cf5_, d_333_cf6_)
                d_336_v49_: _dafny.Map
                def iife17_(_pat_let2_0):
                    def iife18_(d_337_dt__update__tmp_h0_):
                        def iife19_(_pat_let3_0):
                            def iife20_(d_338_dt__update_hcf6_h0_):
                                return D2_DC4((d_337_dt__update__tmp_h0_).cf5, d_338_dt__update_hcf6_h0_)
                            return iife20_(_pat_let3_0)
                        return iife19_(_dafny.CodePoint('d'))
                    return iife18_(_pat_let2_0)
                d_336_v49_ = _dafny.Map({iife17_(d_335_v48_): True})
                d_336_v49_ = (d_336_v49_).set(d_335_v48_, (p1) < (p1))
                d_339_v50_: C0
                nw59_ = C0()
                nw59_.ctor__(d_333_cf6_)
                d_339_v50_ = nw59_
                d_340_v51_: _dafny.MultiSet
                d_340_v51_ = _dafny.MultiSet([d_339_v50_, d_339_v50_, d_339_v50_, d_339_v50_])
                d_340_v51_ = (d_340_v51_).set(d_339_v50_, default__.abs(d_334_cf5_))
                d_333_cf6_ = (d_339_v50_).f12
                d_341_v52_: _dafny.Set
                d_341_v52_ = _dafny.Set({not(p2), p2, p2, default__.fm0(True, p2, 585, d_334_cf5_, globalState), True})
                d_342_v53_: _dafny.MultiSet
                d_342_v53_ = _dafny.MultiSet([p3])
                d_343_v54_: _dafny.Map
                d_343_v54_ = _dafny.Map({d_339_v50_: p3})
                d_344_v55_: _dafny.Map
                d_344_v55_ = _dafny.Map({_dafny.MultiSet([p2]): d_339_v50_})
                d_345_v56_: _dafny.Map
                d_345_v56_ = d_344_v55_
                d_346_v57_: _dafny.Map
                d_346_v57_ = _dafny.Map({(d_312_v36_).cf3: d_345_v56_})
                d_347_v58_: _dafny.Seq
                d_347_v58_ = _dafny.SeqWithoutIsStrInference([p2])
                d_348_v59_: _dafny.Array
                nw60_ = _dafny.Array(None, 17)
                nw60_[int(0)] = len(d_268_v0_)
                nw60_[int(1)] = p0
                nw60_[int(2)] = p3
                nw60_[int(3)] = p0
                nw60_[int(4)] = len(d_341_v52_)
                nw60_[int(5)] = 608
                nw60_[int(6)] = (d_342_v53_).cardinality
                nw60_[int(7)] = len(p1)
                nw60_[int(8)] = ((d_342_v53_)[d_334_cf5_] if (d_334_cf5_) in (d_342_v53_) else (0) - (default__.fm2(True, p0, p0, p1, globalState)))
                nw60_[int(9)] = d_334_cf5_
                nw60_[int(10)] = len(d_343_v54_)
                nw60_[int(11)] = len((d_346_v57_).set(p2, d_345_v56_))
                nw60_[int(12)] = len(d_347_v58_)
                nw60_[int(13)] = p0
                nw60_[int(14)] = p3
                nw60_[int(15)] = d_334_cf5_
                nw60_[int(16)] = default__.fm2(p2, p0, d_334_cf5_, p1, globalState)
                d_348_v59_ = nw60_
                d_349_v60_: _dafny.Array
                nw61_ = _dafny.Array(None, 7)
                nw61_[int(0)] = len(p1)
                nw61_[int(1)] = d_334_cf5_
                nw61_[int(2)] = len(_dafny.Map({_dafny.Set({d_339_v50_}): d_348_v59_}))
                nw61_[int(3)] = p3
                nw61_[int(4)] = 671
                nw61_[int(5)] = d_334_cf5_
                nw61_[int(6)] = p3
                d_349_v60_ = nw61_
                index57_ = default__.safeIndex(655, (d_349_v60_).length(0))
                (d_349_v60_)[index57_] = (-38) + (p3)
                index58_ = default__.safeIndex(655, (d_349_v60_).length(0))
                (d_349_v60_)[index58_] = p0
            elif True:
                d_350___mcc_h4_ = source5_.cf2
                d_351_cf2_ = d_350___mcc_h4_
                d_352_v61_: _dafny.Set
                d_352_v61_ = _dafny.Set({p3})
                d_353_v62_: _dafny.Map
                d_353_v62_ = _dafny.Map({(self).f11: p3})
                d_354_v63_: _dafny.Map
                d_354_v63_ = _dafny.Map({p2: p0})
                d_355_v64_: C0
                nw62_ = C0()
                nw62_.ctor__(d_351_cf2_)
                d_355_v64_ = nw62_
                d_356_v65_: _dafny.Map
                d_356_v65_ = _dafny.Map({d_355_v64_: p2})
                d_357_v66_: _dafny.MultiSet
                d_357_v66_ = _dafny.MultiSet([p0, p3])
                d_358_v67_: _dafny.MultiSet
                d_358_v67_ = _dafny.MultiSet([p2, p2])
                d_359_v68_: _dafny.Array
                nw63_ = _dafny.Array(None, 21)
                nw63_[int(0)] = default__.fm2(p2, p3, len(p1), p1, globalState)
                nw63_[int(1)] = (0) - ((0) - (default__.fm2(p2, p3, p0, p1, globalState)))
                nw63_[int(2)] = ((0) - (p0) if p2 else p3)
                nw63_[int(3)] = p0
                nw63_[int(4)] = (len(d_352_v61_) if not(p2) else p0)
                nw63_[int(5)] = default__.safeDivisionInt(807, p0)
                nw63_[int(6)] = default__.safeModuloInt(len(p1), p3)
                nw63_[int(7)] = p3
                nw63_[int(8)] = p3
                nw63_[int(9)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0 for d_360_i8_ in range(default__.abs(750))]))).cardinality
                nw63_[int(10)] = p3
                nw63_[int(11)] = p3
                nw63_[int(12)] = (p0 if p2 else ((d_353_v62_)[(self).f11] if ((self).f11) in (d_353_v62_) else p3))
                nw63_[int(13)] = p0
                nw63_[int(14)] = len(p1)
                nw63_[int(15)] = -629
                nw63_[int(16)] = (0) - (default__.safeModuloInt(-111, len(d_354_v63_)))
                nw63_[int(17)] = len(d_356_v65_)
                nw63_[int(18)] = (0) - ((d_357_v66_).cardinality)
                nw63_[int(19)] = p3
                nw63_[int(20)] = (p3 if p2 else len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([((d_358_v67_)[p2] if (p2) in (d_358_v67_) else p0)])).cardinality])))
                d_359_v68_ = nw63_
                d_361_v69_: _dafny.Array
                def lambda18_(d_362_v64_):
                    def lambda19_(d_363_i9_):
                        return (d_362_v64_).f12

                    return lambda19_

                init9_ = lambda18_(d_355_v64_)
                nw64_ = _dafny.Array(None, 21)
                for i0_9_ in range(nw64_.length(0)):
                    nw64_[i0_9_] = init9_(i0_9_)
                d_361_v69_ = nw64_
                index59_ = default__.safeIndex(650, (d_361_v69_).length(0))
                (d_361_v69_)[index59_] = d_351_cf2_
                index60_ = default__.safeIndex(650, (d_361_v69_).length(0))
                rhs53_ = d_359_v68_
                rhs54_ = d_351_cf2_
                rhs55_ = p2
                rhs56_ = p2
                lhs48_ = d_361_v69_
                lhs49_ = default__.safeIndex(650, (d_361_v69_).length(0))
                lhs50_ = globalState
                lhs51_ = globalState
                d_359_v68_ = rhs53_
                lhs48_[lhs49_] = rhs54_
                lhs50_.f10 = rhs55_
                lhs51_.f9 = rhs56_
                nw65_ = C0()
                nw65_.ctor__((d_355_v64_).f12)
                d_355_v64_ = nw65_
                d_364_v70_: _dafny.Map
                d_364_v70_ = _dafny.Map({p3: default__.fm0(p2, p2, p3, p3, globalState)})
                def iife21_():
                    coll13_ = _dafny.Map()
                    compr_13_: int
                    for compr_13_ in _dafny.IntegerRange(531, 273):
                        d_365_v71_: int = compr_13_
                        if ((531) <= (d_365_v71_)) and ((d_365_v71_) < (273)):
                            coll13_[(d_365_v71_) - (len(d_268_v0_))] = not((_dafny.SeqWithoutIsStrInference([p2])) == (_dafny.SeqWithoutIsStrInference([p2, p2])))
                    return _dafny.Map(coll13_)
                d_364_v70_ = iife21_()
                
                d_366_v72_: C0
                nw66_ = C0()
                nw66_.ctor__((d_355_v64_).f12)
                d_366_v72_ = nw66_
            d_367_v73_: _dafny.Map
            d_367_v73_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqmai"))})
            d_368_v74_: _dafny.Map
            d_368_v74_ = _dafny.Map({p2: p2})
            d_369_v75_: _dafny.Set
            d_369_v75_ = _dafny.Set({len(d_368_v74_)})
            d_370_v76_: str
            d_370_v76_ = _dafny.CodePoint('o')
            d_371_v77_: _dafny.MultiSet
            d_371_v77_ = _dafny.MultiSet([p2])
            d_372_v78_: _dafny.Seq
            d_372_v78_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twnjp")), p1])
            d_373_v79_: _dafny.Array
            nw67_ = _dafny.Array(None, 28)
            nw67_[int(0)] = len(_dafny.Set({p3}))
            nw67_[int(1)] = p3
            nw67_[int(2)] = (0) - (default__.safeModuloInt(p3, p3))
            nw67_[int(3)] = p0
            nw67_[int(4)] = (p0 if p2 else (0) - (p3))
            nw67_[int(5)] = (len(default__.fm7(p2, p0, p2, len(d_367_v73_), globalState))) - ((0) - (p3))
            nw67_[int(6)] = p0
            nw67_[int(7)] = p3
            nw67_[int(8)] = p3
            nw67_[int(9)] = p3
            nw67_[int(10)] = len(d_369_v75_)
            nw67_[int(11)] = p0
            nw67_[int(12)] = default__.safeDivisionInt(p0, len(d_368_v74_))
            nw67_[int(13)] = (0) - (p0)
            nw67_[int(14)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")))), d_370_v76_))
            nw67_[int(15)] = (d_268_v0_)[default__.safeIndex(250, len(d_268_v0_))]
            nw67_[int(16)] = p3
            nw67_[int(17)] = len(default__.fm8(p2, p1, globalState))
            nw67_[int(18)] = (0) - (p3)
            nw67_[int(19)] = p3
            nw67_[int(20)] = (-652) - (p0)
            nw67_[int(21)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlpfc")))
            nw67_[int(22)] = 739
            nw67_[int(23)] = (((_dafny.MultiSet([p2])).set(p2, default__.abs(p3))) | (d_371_v77_)).cardinality
            nw67_[int(24)] = p3
            nw67_[int(25)] = p3
            nw67_[int(26)] = 655
            nw67_[int(27)] = len((p1) + ((d_372_v78_)[default__.safeIndex(p0, len(d_372_v78_))]))
            d_373_v79_ = nw67_
            d_374_v80_: _dafny.MultiSet
            d_374_v80_ = _dafny.MultiSet([len(p1), p3])
            d_375_v81_: C0
            nw68_ = C0()
            nw68_.ctor__(_dafny.CodePoint('j'))
            d_375_v81_ = nw68_
            d_376_v82_: _dafny.Map
            d_376_v82_ = _dafny.Map({d_371_v77_: d_375_v81_})
            d_377_v83_: _dafny.Map
            d_377_v83_ = d_376_v82_
            d_378_v84_: _dafny.Set
            d_378_v84_ = _dafny.Set({d_377_v83_})
            index61_ = default__.safeIndex(566, (d_373_v79_).length(0))
            (d_373_v79_)[index61_] = ((d_374_v80_)[p0] if (p0) in (d_374_v80_) else (0) - (len(d_378_v84_)))
            d_379_v85_: _dafny.Map
            d_379_v85_ = _dafny.Map({p0: p0})
            d_380_v86_: _dafny.Map
            d_380_v86_ = _dafny.Map({p3: ((d_379_v85_)[p0] if (p0) in (d_379_v85_) else p3)})
            d_381_v87_: _dafny.Map
            d_381_v87_ = _dafny.Map({p3: _dafny.Set({D2_DC3(p2, p2), d_312_v36_, d_312_v36_, d_312_v36_})})
            index62_ = default__.safeIndex(566, (d_373_v79_).length(0))
            rhs57_ = len(default__.fm1(d_379_v85_, not(p2), len(p1), (0) - (p0), globalState))
            rhs58_ = (p1)[default__.safeIndex(414, len(p1))]
            rhs59_ = (default__.fm9(p2, globalState)) != (d_381_v87_)
            lhs52_ = d_373_v79_
            lhs53_ = default__.safeIndex(566, (d_373_v79_).length(0))
            lhs54_ = globalState
            lhs52_[lhs53_] = rhs57_
            d_370_v76_ = rhs58_
            lhs54_.f10 = rhs59_
            (globalState).f9 = p2
            d_382_v88_: _dafny.Map
            d_382_v88_ = _dafny.Map({p2: p0})
            d_382_v88_ = (d_382_v88_).set(p2, (-902) + (len(d_382_v88_)))
        d_383_v89_: _dafny.MultiSet
        d_383_v89_ = _dafny.MultiSet([p2])
        d_384_v90_: str
        d_384_v90_ = _dafny.CodePoint('y')
        d_385_v91_: C0
        nw69_ = C0()
        nw69_.ctor__(d_384_v90_)
        d_385_v91_ = nw69_
        d_386_v92_: _dafny.Map
        d_386_v92_ = _dafny.Map({d_383_v89_: d_385_v91_})
        d_387_v93_: _dafny.Map
        d_387_v93_ = d_386_v92_
        d_388_v94_: _dafny.Seq
        d_388_v94_ = _dafny.SeqWithoutIsStrInference([d_387_v93_])
        hi2_ = default__.fm2(p2, (0) - (-285), p0, p1, globalState)
        for d_389_i10_ in range(default__.safeDivisionInt((_dafny.MultiSet(d_388_v94_)).cardinality, p0), hi2_):
            d_390_v95_: _dafny.Map
            d_390_v95_ = _dafny.Map({p3: p2})
            d_391_v96_: _dafny.Map
            d_391_v96_ = _dafny.Map({p0: len(d_390_v95_)})
            d_392_v97_: _dafny.Map
            d_392_v97_ = _dafny.Map({((d_391_v96_)[p0] if (p0) in (d_391_v96_) else p3): p2})
            d_393_v98_: _dafny.Seq
            d_393_v98_ = _dafny.SeqWithoutIsStrInference([(p2) == (not(((d_392_v97_)[p3] if (p3) in (d_392_v97_) else p2)))])
            (globalState).f10 = (d_393_v98_)[default__.safeIndex(d_389_i10_, len(d_393_v98_))]
            d_394_v99_: _dafny.Array
            nw70_ = _dafny.Array(None, 21)
            nw70_[int(0)] = d_387_v93_
            nw70_[int(1)] = (d_387_v93_ if p2 else d_387_v93_)
            nw70_[int(2)] = d_387_v93_
            nw70_[int(3)] = d_387_v93_
            nw70_[int(4)] = _dafny.Map({d_383_v89_: d_385_v91_})
            nw70_[int(5)] = d_387_v93_
            nw70_[int(6)] = d_387_v93_
            nw70_[int(7)] = d_387_v93_
            nw70_[int(8)] = d_387_v93_
            nw70_[int(9)] = d_387_v93_
            nw70_[int(10)] = d_387_v93_
            nw70_[int(11)] = d_387_v93_
            nw70_[int(12)] = d_387_v93_
            nw70_[int(13)] = d_386_v92_
            nw70_[int(14)] = d_386_v92_
            nw70_[int(15)] = d_387_v93_
            nw70_[int(16)] = d_387_v93_
            nw70_[int(17)] = d_386_v92_
            nw70_[int(18)] = d_386_v92_
            nw70_[int(19)] = d_386_v92_
            nw70_[int(20)] = (d_386_v92_)
            d_394_v99_ = nw70_
            index63_ = default__.safeIndex(950, (d_394_v99_).length(0))
            (d_394_v99_)[index63_] = d_387_v93_
            index64_ = default__.safeIndex(950, (d_394_v99_).length(0))
            (d_394_v99_)[index64_] = d_386_v92_
            d_395_v100_: _dafny.Map
            d_395_v100_ = _dafny.Map({p2: d_389_i10_})
            d_396_v101_: _dafny.Array
            nw71_ = _dafny.Array(False, 11)
            d_396_v101_ = nw71_
            d_397_v102_: _dafny.Map
            d_397_v102_ = _dafny.Map({d_395_v100_: d_396_v101_})
            d_397_v102_ = (d_397_v102_).set(d_395_v100_, d_396_v101_)
            d_398_v103_: C0
            nw72_ = C0()
            nw72_.ctor__(d_384_v90_)
            d_398_v103_ = nw72_
        d_399_v104_: _dafny.Seq
        d_399_v104_ = _dafny.SeqWithoutIsStrInference([p2])
        d_400_v105_: _dafny.Array
        nw73_ = _dafny.Array(None, 26)
        nw73_[int(0)] = p2
        nw73_[int(1)] = p2
        nw73_[int(2)] = (self).f11
        nw73_[int(3)] = (self).f11
        nw73_[int(4)] = (self).f11
        nw73_[int(5)] = (self).f11
        nw73_[int(6)] = (self).f11
        nw73_[int(7)] = (self).f11
        nw73_[int(8)] = (self).f11
        nw73_[int(9)] = True
        nw73_[int(10)] = (self).f11
        nw73_[int(11)] = (self).f11
        nw73_[int(12)] = (self).f11
        nw73_[int(13)] = p2
        nw73_[int(14)] = (self).f11
        nw73_[int(15)] = (self).f11
        nw73_[int(16)] = (self).f11
        nw73_[int(17)] = (self).f11
        nw73_[int(18)] = (self).f11
        nw73_[int(19)] = (self).f11
        nw73_[int(20)] = (d_399_v104_)[default__.safeIndex(p3, len(d_399_v104_))]
        nw73_[int(21)] = (self).f11
        nw73_[int(22)] = (self).f11
        nw73_[int(23)] = (self).f11
        nw73_[int(24)] = p2
        nw73_[int(25)] = (self).f11
        d_400_v105_ = nw73_
        index65_ = default__.safeIndex(463, (d_400_v105_).length(0))
        (d_400_v105_)[index65_] = (self).f11
        index66_ = default__.safeIndex(463, (d_400_v105_).length(0))
        (d_400_v105_)[index66_] = (self).f11
        r0 = default__.fm10(globalState)
        return r0

    @property
    def f11(self):
        return self._f11
