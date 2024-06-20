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
        return not ((_dafny.MultiSet([True])).isdisjoint(_dafny.MultiSet([False, False, not(False), False]))) or ((False) or (True))

    @staticmethod
    def fm3(p0, globalState):
        if True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgwprn"))
        elif True:
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_0_i0_ in range(default__.abs(792))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsxqnnw")))

    @staticmethod
    def fm4(p0, p1, globalState):
        return (0) - ((len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fv")) if not(True) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uosyjyo"))))) * ((D4_DC15(not(False), (_dafny.MultiSet([True])).cardinality)).cf24))

    @staticmethod
    def fm5(globalState):
        return ((_dafny.Set({not(False), False, False})) | (_dafny.Set({True}))) - ((_dafny.Set({True, True})).intersection(_dafny.Set({False, not(True)})))

    @staticmethod
    def fm6(globalState):
        return D2_DC7((len(_dafny.Map({True: False}))) - (242))

    @staticmethod
    def fm7(p0, p1, globalState):
        return ((_dafny.Map({(D1_DC3(False, False, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1_i0_ in range(default__.abs(627))]), False, 45)).cf5: 7})) | (_dafny.Map({False: 177}))) | ((_dafny.Map({True: -120})) | (_dafny.Map({True: (_dafny.MultiSet([False])).cardinality})))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(59, -991):
                d_2_v0_: int = compr_0_
                if ((59) <= (d_2_v0_)) and ((d_2_v0_) < (-991)):
                    coll0_[default__.safeDivisionInt(d_2_v0_, 958)] = (True) or (True)
            return _dafny.Map(coll0_)
        return iife0_()
        

    @staticmethod
    def fm9(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(500, 122):
                d_3_v0_: int = compr_1_
                if ((500) <= (d_3_v0_)) and ((d_3_v0_) < (122)):
                    coll1_ = coll1_.union(_dafny.Set([(d_3_v0_) * (-706)]))
            return _dafny.Set(coll1_)
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(-473, 223):
                d_4_v1_: int = compr_2_
                if ((-473) <= (d_4_v1_)) and ((d_4_v1_) < (223)):
                    coll2_ = coll2_.union(_dafny.Set([default__.safeModuloInt(d_4_v1_, len(_dafny.Map({461: True})))]))
            return _dafny.Set(coll2_)
        return ((iife1_()
        ) - (iife2_()
        )) | (_dafny.Set({(0) - (-290)}))

    @staticmethod
    def fm10(p0, p1, globalState):
        if False:
            if False:
                return _dafny.SeqWithoutIsStrInference([False, not(not(False))])
            elif True:
                return _dafny.SeqWithoutIsStrInference([False])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([False, False, (D6_DC20(True)).cf34])) + (_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm11(globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(337, 918):
                d_5_v0_: int = compr_3_
                if ((337) <= (d_5_v0_)) and ((d_5_v0_) < (918)):
                    coll3_ = coll3_.union(_dafny.Set([default__.safeDivisionInt(d_5_v0_, 99)]))
            return _dafny.Set(coll3_)
        return ((_dafny.Map({False: D3_DC12(D3_DC10(False, 524, len(_dafny.Map({False: 380}))))})) | (_dafny.Map({False: D3_DC12(D3_DC12(D3_DC11()))}))) | ((_dafny.Map({True: D3_DC12(D3_DC10(False, 682, len(iife3_()
)))})) | (_dafny.Map({False: D3_DC12(D3_DC9(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.Map({True: False}))])).cardinality])))})))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        source0_ = (D1_DC2(376) if True else D1_DC2(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arg")))))
        if source0_.is_DC3:
            d_6___mcc_h0_ = source0_.cf5
            d_7___mcc_h1_ = source0_.cf6
            d_8___mcc_h2_ = source0_.cf7
            d_9___mcc_h3_ = source0_.cf8
            d_10___mcc_h4_ = source0_.cf9
            d_11_cf9_ = d_10___mcc_h4_
            d_12_cf8_ = d_9___mcc_h3_
            d_13_cf7_ = d_8___mcc_h2_
            d_14_cf6_ = d_7___mcc_h1_
            d_15_cf5_ = d_6___mcc_h0_
            return D3_DC12(D3_DC10(d_12_cf8_, d_11_cf9_, d_11_cf9_))
        elif source0_.is_DC2:
            d_16___mcc_h5_ = source0_.cf4
            d_17_cf4_ = d_16___mcc_h5_
            return D3_DC12(D3_DC11())
        elif True:
            d_18___mcc_h6_ = source0_.cf10
            d_19_cf10_ = d_18___mcc_h6_
            return D3_DC12(D3_DC12(D3_DC12(D3_DC9(_dafny.SeqWithoutIsStrInference([27])))))

    @staticmethod
    def fm13(p0, p1, globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(937, 842):
                d_20_v0_: int = compr_4_
                if ((937) <= (d_20_v0_)) and ((d_20_v0_) < (842)):
                    coll4_ = coll4_.union(_dafny.Set([default__.safeModuloInt(d_20_v0_, (0) - ((_dafny.MultiSet([True, True])).cardinality))]))
            return _dafny.Set(coll4_)
        return (_dafny.Map({254: _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nk"))), len(_dafny.Set({False, True}))})})) | (_dafny.Map({240: iife4_()
        }))

    @staticmethod
    def fm14(globalState):
        return D1_DC4(D1_DC4(D1_DC3(True, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eewew")), True, 243)))

    @staticmethod
    def fm15(p0, p1, globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False, False, False, False]))).intersection((_dafny.MultiSet([False])) | (_dafny.MultiSet([False])))

    @staticmethod
    def fm16(p0, p1, globalState):
        source1_ = D6_DC20(True)
        if source1_.is_DC19:
            d_21___mcc_h0_ = source1_.cf31
            d_22___mcc_h1_ = source1_.cf32
            d_23___mcc_h2_ = source1_.cf33
            d_24_cf33_ = d_23___mcc_h2_
            d_25_cf32_ = d_22___mcc_h1_
            d_26_cf31_ = d_21___mcc_h0_
            return D4_DC15(d_25_cf32_, len(_dafny.Map({False: (D0_DC0(d_25_cf32_)).cf0})))
        elif source1_.is_DC20:
            d_27___mcc_h3_ = source1_.cf34
            d_28_cf34_ = d_27___mcc_h3_
            return D4_DC15(d_28_cf34_, (0) - (len(_dafny.Map({(0) - (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_28_cf34_])), (_dafny.MultiSet([_dafny.CodePoint('y'), _dafny.CodePoint('m'), _dafny.CodePoint('t'), _dafny.CodePoint('q'), _dafny.CodePoint('a')])).cardinality}))): _dafny.MultiSet([673, len(_dafny.Map({d_28_cf34_: True}))])}))))
        elif source1_.is_DC21:
            def iife5_():
                def iife7_():
                    coll7_ = _dafny.Set()
                    compr_7_: int
                    for compr_7_ in _dafny.IntegerRange(572, 240):
                        d_31_v1_: int = compr_7_
                        if ((572) <= (d_31_v1_)) and ((d_31_v1_) < (240)):
                            coll7_ = coll7_.union(_dafny.Set([(d_31_v1_) * (414)]))
                    return _dafny.Set(coll7_)
                coll5_ = _dafny.Map()
                def iife6_():
                    coll6_ = _dafny.Set()
                    compr_6_: int
                    for compr_6_ in _dafny.IntegerRange(572, 240):
                        d_29_v1_: int = compr_6_
                        if ((572) <= (d_29_v1_)) and ((d_29_v1_) < (240)):
                            coll6_ = coll6_.union(_dafny.Set([(d_29_v1_) * (414)]))
                    return _dafny.Set(coll6_)
                compr_5_: int
                for compr_5_ in (iife6_()
                ).Elements:
                    d_30_v0_: int = compr_5_
                    if (d_30_v0_) in (iife7_()
                    ):
                        coll5_[default__.safeModuloInt(d_30_v0_, (0) - (len(_dafny.Set({691}))))] = True
                return _dafny.Map(coll5_)
            return D4_DC15(True, len(_dafny.Set({(0) - (len(_dafny.Set({len(_dafny.Set({True})), len(iife5_()
)})))})))
        elif True:
            d_32___mcc_h4_ = source1_.cf30
            d_33_cf30_ = d_32___mcc_h4_
            if True:
                return D4_DC15(not(True), len(_dafny.Map({False: 592})))
            elif True:
                return D4_DC15(True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngymo"))))

    @staticmethod
    def fm17(p0, p1, p2, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in (_dafny.SeqWithoutIsStrInference([445, len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([873, 28]))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfd"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cngsgffno"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "giiveerwc")))])).Elements:
                d_34_v0_: int = compr_8_
                if (d_34_v0_) in (_dafny.SeqWithoutIsStrInference([445, len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([873, 28]))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfd"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cngsgffno"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "giiveerwc")))])):
                    coll8_[(d_34_v0_) + ((D4_DC15(False, len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(_dafny.MultiSet([True])).cardinality]) for d_35_i0_ in range(default__.abs(112))])))).cf24)] = -492
            return _dafny.Map(coll8_)
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(233, 864):
                d_36_v1_: int = compr_9_
                if ((233) <= (d_36_v1_)) and ((d_36_v1_) < (864)):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeModuloInt(d_36_v1_, 830)]))
            return _dafny.Set(coll9_)
        return (_dafny.Set({_dafny.Set({len(_dafny.Map({False: (_dafny.MultiSet([True, False])).cardinality}))})})) | (_dafny.Set({_dafny.Set({len(_dafny.SeqWithoutIsStrInference([943]))}), _dafny.Set({-890, len(iife8_()
        )}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([868]))}), _dafny.Set({472}), _dafny.Set({len(_dafny.Set({(D4_DC15(False, 655)).cf24})), -685, len(_dafny.Map({(0) - (len(_dafny.Set({len(iife9_()
        ), (0) - (len(_dafny.Map({False: True})))}))): True})), (0) - ((_dafny.MultiSet([False])).cardinality)})}))

    @staticmethod
    def fm18(p0, globalState):
        return _dafny.CodePoint('h')

    @staticmethod
    def fm19(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_37_i1_ in range(default__.abs(-673))]) for d_38_i0_ in range(default__.abs(300))])

    @staticmethod
    def m0(p0, globalState):
        r0: bool = False
        r0 = p0
        d_39_v0_: D0
        d_39_v0_ = D0_DC0(p0)
        d_40_i0_: int
        d_40_i0_ = 0
        with _dafny.label("0"):
            pat_let_tv3_ = p0
            pat_let_tv4_ = p0
            def lambda0_(source3_):
                if source3_.is_DC0:
                    d_97___mcc_h6_ = source3_.cf0
                    d_98_cf0_ = d_97___mcc_h6_
                    return pat_let_tv3_
                elif True:
                    d_99___mcc_h7_ = source3_.cf1
                    d_100___mcc_h8_ = source3_.cf2
                    d_101___mcc_h9_ = source3_.cf3
                    d_102_cf3_ = d_101___mcc_h9_
                    d_103_cf2_ = d_100___mcc_h8_
                    d_104_cf1_ = d_99___mcc_h7_
                    return ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_105_i1_ in range(default__.abs(-575))])))) == ((D1_DC3(pat_let_tv4_, True, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_106_i2_ in range(default__.abs(431))]), False, (0) - (-332))).cf9)

            while lambda0_(d_39_v0_):
                with _dafny.c_label("0"):
                    if (d_40_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_40_i0_ = (d_40_i0_) + (1)
                    r0 = p0
                    d_41_v1_: int
                    d_41_v1_ = 660
                    d_42_v2_: _dafny.Seq
                    d_42_v2_ = _dafny.SeqWithoutIsStrInference([d_41_v1_, d_41_v1_])
                    d_43_v3_: _dafny.Array
                    nw0_ = _dafny.Array(None, 3)
                    nw0_[int(0)] = d_41_v1_
                    nw0_[int(1)] = d_41_v1_
                    nw0_[int(2)] = len(d_42_v2_)
                    d_43_v3_ = nw0_
                    d_44_v4_: D2
                    d_44_v4_ = D2_DC6(d_41_v1_, True, d_43_v3_)
                    pat_let_tv0_ = p0
                    pat_let_tv1_ = p0
                    pat_let_tv2_ = d_43_v3_
                    def iife10_(_pat_let0_0):
                        def iife11_(d_45_dt__update__tmp_h0_):
                            def iife12_(_pat_let1_0):
                                def iife13_(d_46_dt__update_hcf13_h0_):
                                    def iife14_(_pat_let2_0):
                                        def iife15_(d_47_dt__update_hcf14_h0_):
                                            return D2_DC6((d_45_dt__update__tmp_h0_).cf12, d_46_dt__update_hcf13_h0_, d_47_dt__update_hcf14_h0_)
                                        return iife15_(_pat_let2_0)
                                    return iife14_(pat_let_tv2_)
                                return iife13_(_pat_let1_0)
                            return iife12_((pat_let_tv0_) == (pat_let_tv1_))
                        return iife11_(_pat_let0_0)
                    source2_ = iife10_(d_44_v4_)
                    if source2_.is_DC6:
                        d_48___mcc_h0_ = source2_.cf12
                        d_49___mcc_h1_ = source2_.cf13
                        d_50___mcc_h2_ = source2_.cf14
                        d_51_cf14_ = d_50___mcc_h2_
                        d_52_cf13_ = d_49___mcc_h1_
                        d_53_cf12_ = d_48___mcc_h0_
                        d_54_v5_: _dafny.Set
                        d_54_v5_ = _dafny.Set({False})
                        d_55_v6_: C0
                        nw1_ = C0()
                        nw1_.ctor__(d_53_cf12_, (_dafny.Set({not(p0), default__.fm0(not(p0), p0, d_53_cf12_, globalState), d_52_cf13_}) if p0 else d_54_v5_))
                        d_55_v6_ = nw1_
                        d_56_v7_: _dafny.Seq
                        d_56_v7_ = _dafny.SeqWithoutIsStrInference([p0, p0, (p0) in (_dafny.Set({d_52_cf13_}))])
                        r0 = (d_56_v7_)[default__.safeIndex(d_53_cf12_, len(d_56_v7_))]
                        d_57_v8_: C0
                        nw2_ = C0()
                        nw2_.ctor__((d_55_v6_).f6, d_54_v5_)
                        d_57_v8_ = nw2_
                        d_58_v9_: _dafny.Set
                        d_58_v9_ = _dafny.Set({(d_55_v6_).f6, d_41_v1_})
                        d_59_v10_: _dafny.Seq
                        d_59_v10_ = _dafny.SeqWithoutIsStrInference([d_58_v9_])
                        d_60_v11_: _dafny.Map
                        d_60_v11_ = _dafny.Map({(d_57_v8_).f6: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upqhkypa"))})
                        d_61_v12_: _dafny.Map
                        d_61_v12_ = _dafny.Map({len(d_60_v11_): d_58_v9_})
                        d_62_v15_: _dafny.Map
                        d_62_v15_ = _dafny.Map({p0: (d_57_v8_).f6})
                        d_63_v16_: _dafny.Array
                        nw3_ = _dafny.Array(None, 15)
                        nw3_[int(0)] = (d_59_v10_) + (_dafny.SeqWithoutIsStrInference([((d_61_v12_)[(d_55_v6_).f6] if ((d_55_v6_).f6) in (d_61_v12_) else (d_59_v10_)[default__.safeIndex(d_53_cf12_, len(d_59_v10_))])]))
                        def iife16_():
                            coll10_ = _dafny.Set()
                            compr_10_: int
                            for compr_10_ in _dafny.IntegerRange(239, 328):
                                d_64_v13_: int = compr_10_
                                if ((239) <= (d_64_v13_)) and ((d_64_v13_) < (328)):
                                    coll10_ = coll10_.union(_dafny.Set([default__.safeModuloInt(d_64_v13_, d_53_cf12_)]))
                            return _dafny.Set(coll10_)
                        nw3_[int(1)] = _dafny.SeqWithoutIsStrInference([iife16_()
 for d_65_i3_ in range(default__.abs(169))])
                        nw3_[int(2)] = ((d_59_v10_).set(default__.safeIndex((d_57_v8_).f6, len(d_59_v10_)), d_58_v9_)) + (d_59_v10_)
                        nw3_[int(3)] = d_59_v10_
                        nw3_[int(4)] = (d_59_v10_) + (d_59_v10_)
                        nw3_[int(5)] = d_59_v10_
                        nw3_[int(6)] = _dafny.SeqWithoutIsStrInference([d_58_v9_ for d_66_i4_ in range(default__.abs(845))])
                        nw3_[int(7)] = (_dafny.SeqWithoutIsStrInference([d_58_v9_])) + (_dafny.SeqWithoutIsStrInference([d_58_v9_, d_58_v9_]))
                        def iife17_():
                            coll11_ = _dafny.Set()
                            compr_11_: int
                            for compr_11_ in _dafny.IntegerRange(-118, 257):
                                d_67_v14_: int = compr_11_
                                if ((-118) <= (d_67_v14_)) and ((d_67_v14_) < (257)):
                                    coll11_ = coll11_.union(_dafny.Set([default__.safeDivisionInt(d_67_v14_, len(d_56_v7_))]))
                            return _dafny.Set(coll11_)
                        nw3_[int(8)] = (d_59_v10_).set(default__.safeIndex((d_55_v6_).f6, len(d_59_v10_)), iife17_()
                        )
                        nw3_[int(9)] = (d_59_v10_) + (d_59_v10_)
                        nw3_[int(10)] = (d_59_v10_).set(default__.safeIndex(len(d_62_v15_), len(d_59_v10_)), d_58_v9_)
                        nw3_[int(11)] = d_59_v10_
                        nw3_[int(12)] = (d_59_v10_) + (d_59_v10_)
                        nw3_[int(13)] = d_59_v10_
                        nw3_[int(14)] = d_59_v10_
                        d_63_v16_ = nw3_
                        index0_ = default__.safeIndex(153, (d_63_v16_).length(0))
                        (d_63_v16_)[index0_] = d_59_v10_
                        d_68_v17_: _dafny.MultiSet
                        d_68_v17_ = _dafny.MultiSet([d_53_cf12_])
                        d_69_v18_: C0
                        nw4_ = C0()
                        nw4_.ctor__(default__.safeDivisionInt((d_57_v8_).f6, ((d_62_v15_)[p0] if (p0) in (d_62_v15_) else (d_68_v17_).cardinality)), (d_57_v8_).f7)
                        d_69_v18_ = nw4_
                        index1_ = default__.safeIndex(153, (d_63_v16_).length(0))
                        rhs0_ = d_59_v10_
                        rhs1_ = d_69_v18_
                        rhs2_ = p0
                        lhs0_ = d_63_v16_
                        lhs1_ = default__.safeIndex(153, (d_63_v16_).length(0))
                        lhs0_[lhs1_] = rhs0_
                        d_69_v18_ = rhs1_
                        r0 = rhs2_
                    elif source2_.is_DC7:
                        d_70___mcc_h3_ = source2_.cf15
                        d_71_cf15_ = d_70___mcc_h3_
                        d_72_v19_: _dafny.Set
                        d_72_v19_ = _dafny.Set({602})
                        r0 = (not(p0) if p0 else (d_72_v19_).ispropersubset(d_72_v19_))
                        r0 = default__.fm0(True, p0, default__.safeDivisionInt(d_71_cf15_, d_71_cf15_), globalState)
                        d_73_v20_: _dafny.Set
                        d_73_v20_ = _dafny.Set({d_72_v19_})
                        d_74_v21_: _dafny.Map
                        d_74_v21_ = _dafny.Map({(len(d_73_v20_) if p0 else (0) - (d_71_cf15_)): d_41_v1_})
                        d_74_v21_ = d_74_v21_
                        d_75_v22_: _dafny.Set
                        d_75_v22_ = _dafny.Set({False, p0, p0})
                        d_76_v23_: C0
                        nw5_ = C0()
                        nw5_.ctor__((0) - (d_71_cf15_), d_75_v22_)
                        d_76_v23_ = nw5_
                    elif source2_.is_DC5:
                        d_77___mcc_h4_ = source2_.cf11
                        d_78_cf11_ = d_77___mcc_h4_
                        d_79_v24_: str
                        d_79_v24_ = _dafny.CodePoint('w')
                        d_80_v25_: _dafny.Map
                        d_80_v25_ = _dafny.Map({not(p0): d_79_v24_})
                        d_80_v25_ = (d_80_v25_).set(p0, (d_79_v24_ if False else d_79_v24_))
                        d_81_v26_: C0
                        nw6_ = C0()
                        nw6_.ctor__(default__.safeModuloInt(d_41_v1_, d_41_v1_), _dafny.Set({p0, p0}))
                        d_81_v26_ = nw6_
                        d_81_v26_ = d_81_v26_
                        d_82_v27_: _dafny.Set
                        d_82_v27_ = _dafny.Set({(d_81_v26_).f6, d_41_v1_, -816, d_41_v1_})
                        d_83_v28_: C0
                        nw7_ = C0()
                        nw7_.ctor__(len(d_82_v27_), (d_81_v26_).f7)
                        d_83_v28_ = nw7_
                        (globalState).f0 = (d_41_v1_) - ((d_81_v26_).f6)
                    elif True:
                        d_84___mcc_h5_ = source2_.cf16
                        d_85_cf16_ = d_84___mcc_h5_
                        def iife18_():
                            coll12_ = _dafny.Map()
                            compr_12_: int
                            for compr_12_ in _dafny.IntegerRange(933, 887):
                                d_86_v29_: int = compr_12_
                                if ((933) <= (d_86_v29_)) and ((d_86_v29_) < (887)):
                                    coll12_[default__.safeDivisionInt(d_86_v29_, len(_dafny.SeqWithoutIsStrInference([not(p0)])))] = p0
                            return _dafny.Map(coll12_)
                        d_41_v1_ = len(iife18_()
                        )
                        d_87_v31_: _dafny.Map
                        d_87_v31_ = _dafny.Map({d_41_v1_: d_41_v1_})
                        d_88_v32_: _dafny.Map
                        d_88_v32_ = _dafny.Map({d_87_v31_: d_41_v1_})
                        index2_ = default__.safeIndex(142, (d_43_v3_).length(0))
                        def iife19_():
                            coll13_ = _dafny.Map()
                            compr_13_: int
                            for compr_13_ in (_dafny.MultiSet([(_dafny.MultiSet([p0, p0])).cardinality])).Elements:
                                d_89_v30_: int = compr_13_
                                if (d_89_v30_) in (_dafny.MultiSet([(_dafny.MultiSet([p0, p0])).cardinality])):
                                    coll13_[(d_89_v30_) * (877)] = p0
                            return _dafny.Map(coll13_)
                        (d_43_v3_)[index2_] = (d_41_v1_ if p0 else len((iife19_()
                        ).set(len(d_88_v32_), p0)))
                        index3_ = default__.safeIndex(142, (d_43_v3_).length(0))
                        (d_43_v3_)[index3_] = d_41_v1_
                        d_90_v33_: _dafny.Array
                        nw8_ = _dafny.Array(_dafny.CodePoint('D'), 7)
                        d_90_v33_ = nw8_
                        d_90_v33_ = d_90_v33_
                        index4_ = default__.safeIndex(142, (d_43_v3_).length(0))
                        (d_43_v3_)[index4_] = default__.safeModuloInt((d_43_v3_)[default__.safeIndex(142, (d_43_v3_).length(0))], (d_43_v3_)[default__.safeIndex(142, (d_43_v3_).length(0))])
                    d_91_v34_: _dafny.Array
                    nw9_ = _dafny.Array(False, 4)
                    d_91_v34_ = nw9_
                    d_92_v35_: _dafny.MultiSet
                    d_92_v35_ = _dafny.MultiSet([(d_41_v1_) >= (d_41_v1_)])
                    d_93_v36_: _dafny.Set
                    d_93_v36_ = _dafny.Set({d_41_v1_})
                    d_94_v37_: D3
                    d_94_v37_ = D3_DC11()
                    rhs3_ = d_91_v34_
                    rhs4_ = default__.safeModuloInt(default__.safeModuloInt(d_41_v1_, d_41_v1_), len((d_93_v36_ if p0 else d_93_v36_)))
                    rhs5_ = default__.fm15(d_94_v37_, p0, globalState)
                    lhs2_ = globalState
                    d_91_v34_ = rhs3_
                    lhs2_.f0 = rhs4_
                    d_92_v35_ = rhs5_
                    d_95_v38_: _dafny.Set
                    d_95_v38_ = _dafny.Set({True})
                    d_96_v39_: C0
                    nw10_ = C0()
                    nw10_.ctor__(d_41_v1_, d_95_v38_)
                    d_96_v39_ = nw10_
                    pass
            pass
        d_107_v40_: int
        d_107_v40_ = -284
        source4_ = D4_DC15((d_107_v40_) >= (305), d_107_v40_)
        if source4_.is_DC14:
            r0 = p0
            r0 = p0
            r0 = not(p0)
            d_108_v41_: _dafny.Array
            nw11_ = _dafny.Array(_dafny.Map({}), 6)
            d_108_v41_ = nw11_
            d_109_v42_: _dafny.Array
            nw12_ = _dafny.Array(None, 17)
            nw12_[int(0)] = d_108_v41_
            nw12_[int(1)] = d_108_v41_
            nw12_[int(2)] = d_108_v41_
            nw12_[int(3)] = d_108_v41_
            nw12_[int(4)] = d_108_v41_
            nw12_[int(5)] = d_108_v41_
            nw12_[int(6)] = d_108_v41_
            nw12_[int(7)] = d_108_v41_
            nw12_[int(8)] = d_108_v41_
            nw12_[int(9)] = d_108_v41_
            nw12_[int(10)] = d_108_v41_
            nw12_[int(11)] = d_108_v41_
            nw12_[int(12)] = d_108_v41_
            nw12_[int(13)] = d_108_v41_
            nw12_[int(14)] = d_108_v41_
            nw12_[int(15)] = d_108_v41_
            nw12_[int(16)] = d_108_v41_
            d_109_v42_ = nw12_
            index5_ = default__.safeIndex(30, (d_109_v42_).length(0))
            (d_109_v42_)[index5_] = d_108_v41_
            index6_ = default__.safeIndex(30, (d_109_v42_).length(0))
            (d_109_v42_)[index6_] = d_108_v41_
        elif source4_.is_DC15:
            d_110___mcc_h10_ = source4_.cf23
            d_111___mcc_h11_ = source4_.cf24
            d_112_cf24_ = d_111___mcc_h11_
            d_113_cf23_ = d_110___mcc_h10_
            d_114_v43_: _dafny.Seq
            d_114_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sficpirr"))
            d_115_v44_: str
            d_115_v44_ = _dafny.CodePoint('d')
            d_116_v45_: D1
            d_116_v45_ = D1_DC3(d_113_cf23_, default__.fm0(p0, d_113_cf23_, d_112_cf24_, globalState), ((d_114_v43_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_117_i5_ in range(default__.abs(141))]))).set(default__.safeIndex(d_107_v40_, len((d_114_v43_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_118_i5_ in range(default__.abs(141))])))), d_115_v44_), (len(d_114_v43_)) <= (len(_dafny.Map({p0: d_107_v40_}))), (0) - (d_107_v40_))
            rhs6_ = d_116_v45_
            rhs7_ = default__.fm4(d_107_v40_, p0, globalState)
            d_116_v45_ = rhs6_
            d_112_cf24_ = rhs7_
            d_114_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
            if False:
                (globalState).f4 = 206
                d_119_v46_: _dafny.Set
                d_119_v46_ = _dafny.Set({False, d_113_cf23_})
                d_120_v47_: C0
                nw13_ = C0()
                nw13_.ctor__(d_112_cf24_, d_119_v46_)
                d_120_v47_ = nw13_
                d_121_v48_: _dafny.Map
                d_121_v48_ = _dafny.Map({d_120_v47_: d_113_cf23_})
                d_122_v49_: _dafny.Map
                d_122_v49_ = _dafny.Map({838: d_120_v47_})
                d_113_cf23_ = default__.fm0(p0, (len(d_114_v43_)) not in (_dafny.Map({len((d_121_v48_).set(((d_122_v49_)[d_112_cf24_] if (d_112_cf24_) in (d_122_v49_) else d_120_v47_), p0)): _dafny.Map({d_113_cf23_: (d_114_v43_)[default__.safeIndex((d_120_v47_).f6, len(d_114_v43_))]})})), (0) - ((d_120_v47_).f6), globalState)
                d_123_v50_: C0
                nw14_ = C0()
                nw14_.ctor__((d_120_v47_).f6, d_119_v46_)
                d_123_v50_ = nw14_
                d_124_v51_: _dafny.Set
                d_124_v51_ = _dafny.Set({d_122_v49_})
                d_113_cf23_ = (d_124_v51_).issubset(d_124_v51_)
                d_125_v52_: _dafny.Map
                d_125_v52_ = _dafny.Map({(d_120_v47_).f6: p0})
                d_125_v52_ = (d_125_v52_).set((0) - ((d_123_v50_).f6), p0)
            elif True:
                r0 = not(p0)
                d_126_v53_: _dafny.Set
                d_126_v53_ = _dafny.Set({d_113_cf23_})
                d_127_v54_: C0
                nw15_ = C0()
                nw15_.ctor__(d_107_v40_, d_126_v53_)
                d_127_v54_ = nw15_
                d_128_v55_: D4
                d_128_v55_ = D4_DC13(d_127_v54_)
                d_129_v56_: _dafny.Map
                d_129_v56_ = _dafny.Map({d_128_v55_: d_113_cf23_})
                d_130_v57_: C0
                nw16_ = C0()
                nw16_.ctor__(len(d_129_v56_), d_126_v53_)
                d_130_v57_ = nw16_
                d_131_v58_: _dafny.Seq
                d_131_v58_ = _dafny.SeqWithoutIsStrInference([p0])
                d_132_v59_: _dafny.Map
                d_132_v59_ = _dafny.Map({(d_114_v43_) <= (d_114_v43_): (p0) not in (d_131_v58_)})
                d_132_v59_ = d_132_v59_
                r0 = d_113_cf23_
                d_133_v60_: C0
                nw17_ = C0()
                nw17_.ctor__((0) - ((d_130_v57_).f6), (d_126_v53_) | (_dafny.Set({d_113_cf23_})))
                d_133_v60_ = nw17_
            d_134_v61_: _dafny.Set
            d_134_v61_ = _dafny.Set({d_107_v40_})
            d_135_v62_: _dafny.Map
            d_135_v62_ = _dafny.Map({d_115_v44_: d_134_v61_})
            d_136_v63_: _dafny.Array
            nw18_ = _dafny.Array(None, 10)
            nw18_[int(0)] = (len(d_114_v43_)) - (d_112_cf24_)
            nw18_[int(1)] = d_112_cf24_
            nw18_[int(2)] = d_107_v40_
            nw18_[int(3)] = d_107_v40_
            nw18_[int(4)] = 446
            nw18_[int(5)] = default__.fm4(d_112_cf24_, p0, globalState)
            nw18_[int(6)] = default__.fm4(d_107_v40_, d_113_cf23_, globalState)
            nw18_[int(7)] = d_107_v40_
            nw18_[int(8)] = d_107_v40_
            nw18_[int(9)] = (0) - ((d_107_v40_) - (len(d_135_v62_)))
            d_136_v63_ = nw18_
            index7_ = default__.safeIndex(88, (d_136_v63_).length(0))
            (d_136_v63_)[index7_] = d_107_v40_
            d_137_v64_: _dafny.Array
            def lambda1_(d_138_v43_):
                def lambda2_(d_139_i6_):
                    return d_138_v43_

                return lambda2_

            init0_ = lambda1_(d_114_v43_)
            nw19_ = _dafny.Array(None, 10)
            for i0_0_ in range(nw19_.length(0)):
                nw19_[i0_0_] = init0_(i0_0_)
            d_137_v64_ = nw19_
            index8_ = default__.safeIndex(965, (d_137_v64_).length(0))
            (d_137_v64_)[index8_] = d_114_v43_
            d_140_v65_: _dafny.Array
            def lambda3_(d_141_v43_):
                def lambda4_(d_142_i7_):
                    return not(not((d_141_v43_) <= (d_141_v43_)))

                return lambda4_

            init1_ = lambda3_(d_114_v43_)
            nw20_ = _dafny.Array(None, 11)
            for i0_1_ in range(nw20_.length(0)):
                nw20_[i0_1_] = init1_(i0_1_)
            d_140_v65_ = nw20_
            d_143_v66_: _dafny.Seq
            d_143_v66_ = _dafny.SeqWithoutIsStrInference([p0])
            index9_ = default__.safeIndex(41, (d_140_v65_).length(0))
            (d_140_v65_)[index9_] = (_dafny.SeqWithoutIsStrInference([p0, d_113_cf23_, d_113_cf23_])) == (d_143_v66_)
            index10_ = default__.safeIndex(88, (d_136_v63_).length(0))
            index11_ = default__.safeIndex(965, (d_137_v64_).length(0))
            index12_ = default__.safeIndex(41, (d_140_v65_).length(0))
            rhs8_ = d_107_v40_
            rhs9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihfjckhnn"))
            rhs10_ = d_107_v40_
            rhs11_ = ((d_112_cf24_) * (len(d_114_v43_))) - (205)
            rhs12_ = p0
            lhs3_ = d_136_v63_
            lhs4_ = default__.safeIndex(88, (d_136_v63_).length(0))
            lhs5_ = d_137_v64_
            lhs6_ = default__.safeIndex(965, (d_137_v64_).length(0))
            lhs7_ = globalState
            lhs8_ = globalState
            lhs9_ = d_140_v65_
            lhs10_ = default__.safeIndex(41, (d_140_v65_).length(0))
            lhs3_[lhs4_] = rhs8_
            lhs5_[lhs6_] = rhs9_
            lhs7_.f0 = rhs10_
            lhs8_.f0 = rhs11_
            lhs9_[lhs10_] = rhs12_
        elif source4_.is_DC16:
            d_144___mcc_h12_ = source4_.cf25
            d_145___mcc_h13_ = source4_.cf26
            d_146___mcc_h14_ = source4_.cf27
            d_147___mcc_h15_ = source4_.cf28
            d_148_cf28_ = d_147___mcc_h15_
            d_149_cf27_ = d_146___mcc_h14_
            d_150_cf26_ = d_145___mcc_h13_
            d_151_cf25_ = d_144___mcc_h12_
            d_152_v67_: _dafny.Seq
            d_152_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adyhsx"))
            d_153_v68_: _dafny.Map
            d_153_v68_ = _dafny.Map({d_152_v67_: p0})
            d_154_v69_: _dafny.MultiSet
            d_154_v69_ = _dafny.MultiSet([d_107_v40_, len(d_153_v68_)])
            d_155_v70_: _dafny.Seq
            d_155_v70_ = _dafny.SeqWithoutIsStrInference([len(d_152_v67_), d_149_cf27_])
            d_156_v71_: D2
            d_156_v71_ = D2_DC7((d_155_v70_)[default__.safeIndex(467, len(d_155_v70_))])
            d_157_v72_: _dafny.Map
            d_157_v72_ = _dafny.Map({-151: d_156_v71_})
            d_158_v73_: _dafny.MultiSet
            d_158_v73_ = _dafny.MultiSet([(d_149_cf27_) >= (((d_154_v69_)[d_107_v40_] if (d_107_v40_) in (d_154_v69_) else len(d_157_v72_)))])
            d_159_v74_: _dafny.Array
            def lambda5_(d_160_v70_, d_161_cf26_):
                def lambda6_(d_162_i8_):
                    return (d_162_i8_) + ((d_160_v70_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_161_cf26_ for d_163_i9_ in range(default__.abs(3))])), len(d_160_v70_))])

                return lambda6_

            init2_ = lambda5_(d_155_v70_, d_150_cf26_)
            nw21_ = _dafny.Array(None, 6)
            for i0_2_ in range(nw21_.length(0)):
                nw21_[i0_2_] = init2_(i0_2_)
            d_159_v74_ = nw21_
            d_164_v75_: _dafny.Map
            d_164_v75_ = _dafny.Map({d_159_v74_: p0})
            d_165_v76_: _dafny.Map
            d_165_v76_ = _dafny.Map({365: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csu"))})
            (globalState).f0 = ((d_158_v73_)[default__.fm0(p0, ((d_164_v75_)[d_159_v74_] if (d_159_v74_) in (d_164_v75_) else p0), len(d_165_v76_), globalState)] if (default__.fm0(p0, ((d_164_v75_)[d_159_v74_] if (d_159_v74_) in (d_164_v75_) else p0), len(d_165_v76_), globalState)) in (d_158_v73_) else d_107_v40_)
            r0 = p0
            d_166_v77_: _dafny.Array
            nw22_ = _dafny.Array(False, 23)
            d_166_v77_ = nw22_
            d_166_v77_ = d_166_v77_
            index13_ = default__.safeIndex(44, (d_159_v74_).length(0))
            (d_159_v74_)[index13_] = (d_149_cf27_) + (d_107_v40_)
            index14_ = default__.safeIndex(44, (d_159_v74_).length(0))
            (d_159_v74_)[index14_] = d_149_cf27_
        elif True:
            d_167___mcc_h16_ = source4_.cf22
            d_168_cf22_ = d_167___mcc_h16_
            d_169_v78_: D3
            d_169_v78_ = D3_DC10(p0, d_107_v40_, d_107_v40_)
            source5_ = d_169_v78_
            if source5_.is_DC10:
                d_170___mcc_h17_ = source5_.cf18
                d_171___mcc_h18_ = source5_.cf19
                d_172___mcc_h19_ = source5_.cf20
                d_173_cf20_ = d_172___mcc_h19_
                d_174_cf19_ = d_171___mcc_h18_
                d_175_cf18_ = d_170___mcc_h17_
                d_176_v79_: _dafny.Seq
                d_176_v79_ = _dafny.SeqWithoutIsStrInference([d_107_v40_])
                d_175_cf18_ = (d_176_v79_) != (d_176_v79_)
                d_175_cf18_ = d_175_cf18_
                d_177_v80_: _dafny.Array
                nw23_ = _dafny.Array(None, 9)
                nw23_[int(0)] = -215
                nw23_[int(1)] = len(d_176_v79_)
                nw23_[int(2)] = (d_168_cf22_).f6
                nw23_[int(3)] = (d_168_cf22_).f6
                nw23_[int(4)] = (d_168_cf22_).f6
                nw23_[int(5)] = d_174_cf19_
                nw23_[int(6)] = d_173_cf20_
                nw23_[int(7)] = d_107_v40_
                nw23_[int(8)] = d_107_v40_
                d_177_v80_ = nw23_
                d_178_v81_: D2
                d_178_v81_ = D2_DC6(d_173_cf20_, p0, d_177_v80_)
                d_179_v82_: _dafny.MultiSet
                d_179_v82_ = _dafny.MultiSet([(d_178_v81_).cf12, (d_168_cf22_).f6])
                d_180_v83_: _dafny.Map
                d_180_v83_ = _dafny.Map({(d_168_cf22_).f6: d_179_v82_})
                (globalState).f4 = (d_168_cf22_).fm1((((d_180_v83_)[d_107_v40_] if (d_107_v40_) in (d_180_v83_) else d_179_v82_)).set((0) - ((d_168_cf22_).f6), default__.abs((d_168_cf22_).f6)), d_39_v0_, globalState)
                d_177_v80_ = d_177_v80_
            elif source5_.is_DC11:
                d_181_v84_: _dafny.Map
                d_181_v84_ = _dafny.Map({p0: True})
                d_182_v85_: _dafny.Seq
                d_182_v85_ = _dafny.SeqWithoutIsStrInference([d_107_v40_, (d_168_cf22_).f6])
                d_183_v86_: _dafny.Map
                d_183_v86_ = _dafny.Map({(d_168_cf22_).f6: d_182_v85_})
                d_184_v87_: _dafny.Map
                d_184_v87_ = _dafny.Map({len(d_181_v84_): (d_182_v85_) + (((d_183_v86_)[603] if (603) in (d_183_v86_) else d_182_v85_))})
                d_185_v88_: str
                d_185_v88_ = _dafny.CodePoint('j')
                d_186_v89_: _dafny.MultiSet
                d_186_v89_ = _dafny.MultiSet([p0, p0, not(p0)])
                d_187_v90_: _dafny.Seq
                d_187_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcyod"))
                d_188_v91_: _dafny.Seq
                d_188_v91_ = _dafny.SeqWithoutIsStrInference([(default__.fm7(d_186_v89_, d_187_v90_, globalState)).set(p0, (d_168_cf22_).f6)])
                d_183_v86_ = (d_183_v86_).set((d_168_cf22_).fm2((default__.fm16(d_107_v40_, d_185_v88_, globalState)).cf23, d_182_v85_, d_182_v85_, (d_188_v91_)[default__.safeIndex(d_107_v40_, len(d_188_v91_))], globalState), d_182_v85_)
                r0 = not(not(p0))
                d_189_v92_: _dafny.Map
                d_189_v92_ = _dafny.Map({d_107_v40_: default__.fm0(p0, p0, (d_168_cf22_).f6, globalState)})
                d_190_v93_: _dafny.Set
                d_190_v93_ = _dafny.Set({d_107_v40_, d_107_v40_, (d_168_cf22_).f6, (d_168_cf22_).f6, d_107_v40_})
                def iife20_():
                    coll14_ = _dafny.Set()
                    compr_14_: int
                    for compr_14_ in _dafny.IntegerRange(312, 235):
                        d_191_v94_: int = compr_14_
                        if ((312) <= (d_191_v94_)) and ((d_191_v94_) < (235)):
                            coll14_ = coll14_.union(_dafny.Set([default__.safeModuloInt(d_191_v94_, (d_168_cf22_).f6)]))
                    return _dafny.Set(coll14_)
                d_189_v92_ = (d_189_v92_).set((d_168_cf22_).f6, (iife20_()
                ).ispropersubset(d_190_v93_))
                (globalState).f4 = (0) - ((d_168_cf22_).f6)
            elif source5_.is_DC9:
                d_192___mcc_h20_ = source5_.cf17
                d_193_cf17_ = d_192___mcc_h20_
                d_194_v95_: _dafny.Array
                nw24_ = _dafny.Array(False, 29)
                d_194_v95_ = nw24_
                d_194_v95_ = d_194_v95_
                d_195_v96_: _dafny.Array
                nw25_ = _dafny.Array(int(0), 12)
                d_195_v96_ = nw25_
                rhs13_ = p0
                rhs14_ = d_195_v96_
                rhs15_ = d_107_v40_
                rhs16_ = (d_168_cf22_).f6
                lhs11_ = globalState
                lhs12_ = globalState
                r0 = rhs13_
                d_195_v96_ = rhs14_
                lhs11_.f0 = rhs15_
                lhs12_.f4 = rhs16_
                d_196_v97_: C0
                nw26_ = C0()
                nw26_.ctor__((d_168_cf22_).f6, (d_168_cf22_).f7)
                d_196_v97_ = nw26_
                d_197_v98_: str
                d_197_v98_ = _dafny.CodePoint('c')
                d_197_v98_ = d_197_v98_
            elif True:
                d_198___mcc_h21_ = source5_.cf21
                d_199_cf21_ = d_198___mcc_h21_
                nw27_ = C0()
                def iife21_():
                    coll15_ = _dafny.Map()
                    compr_15_: int
                    for compr_15_ in _dafny.IntegerRange(320, -184):
                        d_200_v99_: int = compr_15_
                        if ((320) <= (d_200_v99_)) and ((d_200_v99_) < (-184)):
                            coll15_[(d_200_v99_) + ((d_168_cf22_).f6)] = p0
                    return _dafny.Map(coll15_)
                nw27_.ctor__(len(iife21_()
                ), (d_168_cf22_).f7)
                d_168_cf22_ = nw27_
                d_201_v100_: _dafny.Array
                nw28_ = _dafny.Array(int(0), 12)
                d_201_v100_ = nw28_
                d_201_v100_ = d_201_v100_
                d_202_v101_: _dafny.Seq
                d_202_v101_ = _dafny.SeqWithoutIsStrInference([p0])
                d_203_v102_: _dafny.Array
                nw29_ = _dafny.Array(None, 15)
                nw29_[int(0)] = p0
                nw29_[int(1)] = p0
                nw29_[int(2)] = not (not(p0)) or (p0)
                nw29_[int(3)] = p0
                nw29_[int(4)] = p0
                nw29_[int(5)] = (d_202_v101_)[default__.safeIndex((d_168_cf22_).f6, len(d_202_v101_))]
                nw29_[int(6)] = p0
                nw29_[int(7)] = (d_169_v78_).cf18
                nw29_[int(8)] = (p0) == (p0)
                nw29_[int(9)] = False
                nw29_[int(10)] = p0
                nw29_[int(11)] = p0
                nw29_[int(12)] = p0
                nw29_[int(13)] = p0
                nw29_[int(14)] = default__.fm0(not(p0), False, (0) - (len(d_202_v101_)), globalState)
                d_203_v102_ = nw29_
                index15_ = default__.safeIndex(563, (d_203_v102_).length(0))
                (d_203_v102_)[index15_] = p0
                index16_ = default__.safeIndex(563, (d_203_v102_).length(0))
                (d_203_v102_)[index16_] = not(p0)
                d_204_v103_: _dafny.Seq
                d_204_v103_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pibona"))
                d_205_v104_: _dafny.Seq
                d_205_v104_ = _dafny.SeqWithoutIsStrInference([d_107_v40_, len(d_204_v103_)])
                d_206_v105_: _dafny.Map
                d_206_v105_ = _dafny.Map({default__.fm0(p0, (d_203_v102_)[default__.safeIndex(563, (d_203_v102_).length(0))], len(d_205_v104_), globalState): d_168_cf22_})
                d_207_v106_: _dafny.Map
                d_207_v106_ = d_206_v105_
                d_208_v107_: _dafny.Array
                nw30_ = _dafny.Array(None, 3)
                nw30_[int(0)] = _dafny.Map({default__.fm0(p0, (d_203_v102_)[default__.safeIndex(563, (d_203_v102_).length(0))], (d_168_cf22_).f6, globalState): d_168_cf22_})
                nw30_[int(1)] = _dafny.Map({False: d_168_cf22_})
                nw30_[int(2)] = ((d_207_v106_)) | (d_206_v105_)
                d_208_v107_ = nw30_
                d_208_v107_ = d_208_v107_
            if not(p0):
                d_209_v108_: str
                d_209_v108_ = _dafny.CodePoint('w')
                d_210_v109_: _dafny.Map
                d_210_v109_ = _dafny.Map({d_209_v108_: d_168_cf22_})
                d_210_v109_ = (d_210_v109_).set((d_209_v108_ if not(p0) else d_209_v108_), d_168_cf22_)
                d_211_v110_: C0
                nw31_ = C0()
                nw31_.ctor__(((d_168_cf22_).f6) * ((d_168_cf22_).f6), ((d_168_cf22_).f7).intersection(_dafny.Set({p0})))
                d_211_v110_ = nw31_
                (globalState).f4 = (d_168_cf22_).f6
                d_212_v111_: _dafny.MultiSet
                d_212_v111_ = _dafny.MultiSet([(d_168_cf22_).f6, (d_168_cf22_).f6])
                d_213_v112_: _dafny.Seq
                d_213_v112_ = _dafny.SeqWithoutIsStrInference([(d_212_v111_).cardinality, -745])
                d_214_v114_: _dafny.Seq
                def iife22_():
                    coll16_ = _dafny.Set()
                    compr_16_: int
                    for compr_16_ in (d_213_v112_).Elements:
                        d_215_v113_: int = compr_16_
                        if (d_215_v113_) in (d_213_v112_):
                            coll16_ = coll16_.union(_dafny.Set([default__.safeDivisionInt(d_215_v113_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yr"))))]))
                    return _dafny.Set(coll16_)
                d_214_v114_ = _dafny.SeqWithoutIsStrInference([len(iife22_()
                )])
                d_214_v114_ = d_214_v114_
                r0 = (((d_168_cf22_).f6) * ((d_168_cf22_).f6)) < (d_107_v40_)
            elif True:
                d_216_v115_: _dafny.Set
                d_216_v115_ = _dafny.Set({682})
                d_216_v115_ = _dafny.Set({d_107_v40_, d_107_v40_, d_107_v40_})
                d_217_v116_: _dafny.Map
                d_217_v116_ = _dafny.Map({d_168_cf22_: False})
                d_218_v117_: _dafny.Seq
                d_218_v117_ = _dafny.SeqWithoutIsStrInference([p0, False, (d_217_v116_) != (d_217_v116_)])
                d_218_v117_ = ((d_218_v117_ if False else d_218_v117_)) + (d_218_v117_)
                d_219_v118_: _dafny.Array
                nw32_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_219_v118_ = nw32_
                d_219_v118_ = d_219_v118_
                d_220_v119_: _dafny.Array
                nw33_ = _dafny.Array(False, 17)
                d_220_v119_ = nw33_
                d_221_v120_: str
                d_221_v120_ = _dafny.CodePoint('y')
                d_222_v122_: _dafny.Map
                d_222_v122_ = _dafny.Map({d_221_v120_: -9})
                index17_ = default__.safeIndex(836, (d_220_v119_).length(0))
                def iife23_():
                    coll17_ = _dafny.Map()
                    compr_17_: str
                    for compr_17_ in (d_222_v122_).keys.Elements:
                        d_223_v121_: str = compr_17_
                        if (d_223_v121_) in (d_222_v122_):
                            coll17_[d_223_v121_] = p0
                    return _dafny.Map(coll17_)
                (d_220_v119_)[index17_] = (_dafny.Map({d_221_v120_: p0})) == (iife23_()
                )
                index18_ = default__.safeIndex(836, (d_220_v119_).length(0))
                (d_220_v119_)[index18_] = p0
                d_224_v123_: _dafny.Array
                def lambda7_(d_225_cf22_):
                    def lambda8_(d_226_i10_):
                        return (d_226_i10_) + ((d_225_cf22_).f6)

                    return lambda8_

                init3_ = lambda7_(d_168_cf22_)
                nw34_ = _dafny.Array(None, 19)
                for i0_3_ in range(nw34_.length(0)):
                    nw34_[i0_3_] = init3_(i0_3_)
                d_224_v123_ = nw34_
                d_227_v124_: _dafny.Map
                d_227_v124_ = _dafny.Map({d_220_v119_: d_224_v123_})
                index19_ = default__.safeIndex(836, (d_220_v119_).length(0))
                (d_220_v119_)[index19_] = (d_220_v119_) not in (d_227_v124_)
            d_228_v125_: _dafny.Set
            d_228_v125_ = _dafny.Set({(d_168_cf22_).f6})
            d_229_v126_: _dafny.Map
            d_229_v126_ = _dafny.Map({d_228_v125_: d_107_v40_})
            d_230_v128_: _dafny.Map
            def iife24_():
                coll18_ = _dafny.Set()
                compr_18_: _dafny.Set
                for compr_18_ in (d_229_v126_).keys.Elements:
                    d_231_v127_: _dafny.Set = compr_18_
                    if (d_231_v127_) in (d_229_v126_):
                        coll18_ = coll18_.union(_dafny.Set([d_231_v127_]))
                return _dafny.Set(coll18_)
            d_230_v128_ = _dafny.Map({d_107_v40_: iife24_()
            })
            d_232_v129_: _dafny.Set
            d_232_v129_ = _dafny.Set({d_228_v125_, d_228_v125_})
            d_107_v40_ = len(((((d_230_v128_)[398] if (398) in (d_230_v128_) else d_232_v129_)) - (default__.fm17(d_107_v40_, p0, (0) - ((d_168_cf22_).f6), globalState))) | (d_232_v129_))
            d_233_v130_: _dafny.MultiSet
            d_233_v130_ = _dafny.MultiSet([(d_168_cf22_).f6])
            (globalState).f0 = ((d_168_cf22_).fm1(d_233_v130_, d_39_v0_, globalState)) * ((d_168_cf22_).f6)
        d_234_v131_: _dafny.Seq
        d_234_v131_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xukko"))
        if (((0) - (d_107_v40_)) == (d_107_v40_) if (d_234_v131_) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_235_i11_ in range(default__.abs(910))])) else ((0) - (d_107_v40_)) < (d_107_v40_)):
            r0 = p0
            if ((d_107_v40_) - (d_107_v40_)) >= (d_107_v40_):
                d_236_v132_: D4
                d_236_v132_ = D4_DC15(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rq"))))
                d_237_v133_: _dafny.Map
                d_237_v133_ = _dafny.Map({default__.fm0(p0, p0, len(d_234_v131_), globalState): (d_236_v132_).cf23})
                d_238_v134_: _dafny.Set
                d_238_v134_ = _dafny.Set({d_107_v40_, d_107_v40_})
                d_239_v135_: _dafny.Set
                d_239_v135_ = _dafny.Set({p0, p0})
                d_240_v136_: C0
                nw35_ = C0()
                nw35_.ctor__(d_107_v40_, d_239_v135_)
                d_240_v136_ = nw35_
                d_241_v137_: _dafny.Map
                d_241_v137_ = _dafny.Map({d_240_v136_: d_107_v40_})
                d_242_v138_: _dafny.MultiSet
                d_242_v138_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True])), ((d_241_v137_)[d_240_v136_] if (d_240_v136_) in (d_241_v137_) else (d_240_v136_).f6)])
                d_237_v133_ = (d_237_v133_).set((d_238_v134_).isdisjoint(_dafny.Set({d_107_v40_, len(_dafny.Set({p0})), len(d_234_v131_), ((d_242_v138_)[d_107_v40_] if (d_107_v40_) in (d_242_v138_) else d_107_v40_)})), p0)
                d_243_v139_: _dafny.Map
                d_243_v139_ = _dafny.Map({default__.fm18(d_107_v40_, globalState): d_107_v40_})
                d_244_v140_: str
                d_244_v140_ = _dafny.CodePoint('l')
                d_245_v141_: _dafny.Map
                d_245_v141_ = _dafny.Map({d_107_v40_: p0})
                (globalState).f0 = (0) - (((d_243_v139_)[d_244_v140_] if (d_244_v140_) in (d_243_v139_) else len((d_245_v141_).set(d_107_v40_, p0))))
                d_246_v142_: C0
                nw36_ = C0()
                nw36_.ctor__(((d_240_v136_).f6) * (d_107_v40_), (d_240_v136_).f7)
                d_246_v142_ = nw36_
                d_244_v140_ = d_244_v140_
                d_247_v143_: _dafny.MultiSet
                d_247_v143_ = _dafny.MultiSet([p0])
                d_248_v144_: _dafny.Map
                d_248_v144_ = _dafny.Map({d_247_v143_: d_244_v140_})
                d_249_v145_: _dafny.Seq
                d_249_v145_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_248_v144_ = (d_248_v144_).set(_dafny.MultiSet((d_249_v145_) + (d_249_v145_)), d_244_v140_)
            elif True:
                r0 = p0
                d_250_v146_: _dafny.Array
                nw37_ = _dafny.Array(None, 12)
                d_250_v146_ = nw37_
                d_251_v147_: _dafny.Set
                d_251_v147_ = _dafny.Set({p0})
                d_252_v148_: C0
                nw38_ = C0()
                nw38_.ctor__(len(_dafny.Map({p0: d_107_v40_})), d_251_v147_)
                d_252_v148_ = nw38_
                index20_ = default__.safeIndex(981, (d_250_v146_).length(0))
                (d_250_v146_)[index20_] = d_252_v148_
                index21_ = default__.safeIndex(981, (d_250_v146_).length(0))
                rhs17_ = d_252_v148_
                rhs18_ = d_107_v40_
                rhs19_ = (d_252_v148_).f6
                rhs20_ = d_252_v148_
                lhs13_ = d_250_v146_
                lhs14_ = default__.safeIndex(981, (d_250_v146_).length(0))
                lhs15_ = globalState
                lhs16_ = globalState
                lhs13_[lhs14_] = rhs17_
                lhs15_.f4 = rhs18_
                lhs16_.f4 = rhs19_
                d_252_v148_ = rhs20_
                d_253_v149_: _dafny.Array
                def lambda9_(d_254_v148_, d_255_v40_, d_256_p0_, d_257_v131_):
                    def lambda10_(d_258_i12_):
                        return (_dafny.MultiSet([(d_254_v148_).f6, (d_254_v148_).f6])).intersection(_dafny.MultiSet([(d_254_v148_).f6, d_255_v40_, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_256_p0_])), d_255_v40_, len(d_257_v131_), d_255_v40_}))]))

                    return lambda10_

                init4_ = lambda9_(d_252_v148_, d_107_v40_, p0, d_234_v131_)
                nw39_ = _dafny.Array(None, 7)
                for i0_4_ in range(nw39_.length(0)):
                    nw39_[i0_4_] = init4_(i0_4_)
                d_253_v149_ = nw39_
                d_259_v150_: _dafny.MultiSet
                d_259_v150_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_107_v40_, (d_252_v148_).f6])), (0) - (d_107_v40_)])
                index22_ = default__.safeIndex(685, (d_253_v149_).length(0))
                (d_253_v149_)[index22_] = d_259_v150_
                index23_ = default__.safeIndex(685, (d_253_v149_).length(0))
                (d_253_v149_)[index23_] = ((d_259_v150_ if p0 else d_259_v150_)) | (d_259_v150_)
                d_260_v151_: _dafny.MultiSet
                d_260_v151_ = _dafny.MultiSet([not(p0)])
                (globalState).f0 = ((d_252_v148_).f6) + ((d_260_v151_).cardinality)
                r0 = (d_234_v131_) <= (d_234_v131_)
            (globalState).f0 = ((0) - (d_107_v40_) if p0 else d_107_v40_)
            d_261_v152_: C0
            nw40_ = C0()
            nw40_.ctor__(d_107_v40_, _dafny.Set({p0}))
            d_261_v152_ = nw40_
            d_262_v153_: _dafny.Map
            d_262_v153_ = _dafny.Map({p0: (d_261_v152_).f6})
            d_263_v154_: _dafny.Seq
            d_263_v154_ = _dafny.SeqWithoutIsStrInference([d_107_v40_, (d_261_v152_).fm2(p0, (_dafny.SeqWithoutIsStrInference([d_107_v40_])).set(default__.safeIndex(d_107_v40_, len(_dafny.SeqWithoutIsStrInference([d_107_v40_]))), d_107_v40_), _dafny.SeqWithoutIsStrInference([(d_261_v152_).f6 for d_264_i14_ in range(default__.abs(629))]), d_262_v153_, globalState)])
            d_107_v40_ = default__.safeModuloInt(d_107_v40_, (d_261_v152_).fm2(True, _dafny.SeqWithoutIsStrInference([(d_261_v152_).f6 for d_265_i13_ in range(default__.abs(27))]), d_263_v154_, _dafny.Map({p0: len(default__.fm3(_dafny.SeqWithoutIsStrInference([p0]), globalState))}), globalState))
        elif True:
            (globalState).f0 = d_107_v40_
            d_266_v155_: _dafny.Array
            nw41_ = _dafny.Array(False, 14)
            d_266_v155_ = nw41_
            d_267_v156_: _dafny.Map
            d_267_v156_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_266_v155_]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))})
            d_268_v157_: _dafny.Seq
            d_268_v157_ = _dafny.SeqWithoutIsStrInference([d_266_v155_, d_266_v155_])
            d_269_v158_: str
            d_269_v158_ = _dafny.CodePoint('t')
            d_270_v159_: _dafny.Seq
            d_270_v159_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxrvpul")), (d_234_v131_).set(default__.safeIndex(len(((d_267_v156_)[(d_268_v157_).set(default__.safeIndex(d_107_v40_, len(d_268_v157_)), d_266_v155_)] if ((d_268_v157_).set(default__.safeIndex(d_107_v40_, len(d_268_v157_)), d_266_v155_)) in (d_267_v156_) else d_234_v131_)), len(d_234_v131_)), d_269_v158_)])
            d_271_v160_: _dafny.Seq
            d_271_v160_ = _dafny.SeqWithoutIsStrInference([p0])
            d_272_v161_: _dafny.Array
            nw42_ = _dafny.Array(None, 29)
            nw42_[int(0)] = (d_270_v159_) + (d_270_v159_)
            nw42_[int(1)] = (d_270_v159_) + (d_270_v159_)
            nw42_[int(2)] = _dafny.SeqWithoutIsStrInference([d_234_v131_, (d_234_v131_).set(default__.safeIndex(d_107_v40_, len(d_234_v131_)), d_269_v158_)])
            nw42_[int(3)] = d_270_v159_
            nw42_[int(4)] = d_270_v159_
            nw42_[int(5)] = _dafny.SeqWithoutIsStrInference([d_234_v131_, _dafny.SeqWithoutIsStrInference([(D6_DC19(len(d_234_v131_), p0, d_269_v158_)).cf33 for d_273_i15_ in range(default__.abs(328))])])
            nw42_[int(6)] = d_270_v159_
            nw42_[int(7)] = default__.fm19(False, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - (d_107_v40_): True})) for d_274_i16_ in range(default__.abs(-351))])), globalState)
            nw42_[int(8)] = d_270_v159_
            nw42_[int(9)] = (d_270_v159_) + (d_270_v159_)
            nw42_[int(10)] = (d_270_v159_) + (d_270_v159_)
            nw42_[int(11)] = d_270_v159_
            nw42_[int(12)] = d_270_v159_
            nw42_[int(13)] = d_270_v159_
            nw42_[int(14)] = d_270_v159_
            nw42_[int(15)] = (d_270_v159_).set(default__.safeIndex(len(d_234_v131_), len(d_270_v159_)), d_234_v131_)
            nw42_[int(16)] = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrr"))])) + (d_270_v159_)
            nw42_[int(17)] = d_270_v159_
            nw42_[int(18)] = d_270_v159_
            nw42_[int(19)] = (d_270_v159_) + ((d_270_v159_).set(default__.safeIndex(d_107_v40_, len(d_270_v159_)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igyetaiv"))))
            nw42_[int(20)] = _dafny.SeqWithoutIsStrInference([d_234_v131_ for d_275_i17_ in range(default__.abs(905))])
            nw42_[int(21)] = _dafny.SeqWithoutIsStrInference([default__.fm3(d_271_v160_, globalState)])
            nw42_[int(22)] = (d_270_v159_) + (d_270_v159_)
            nw42_[int(23)] = _dafny.SeqWithoutIsStrInference([d_234_v131_, _dafny.SeqWithoutIsStrInference([d_269_v158_ for d_276_i18_ in range(default__.abs(353))])])
            nw42_[int(24)] = (d_270_v159_) + (d_270_v159_)
            nw42_[int(25)] = d_270_v159_
            nw42_[int(26)] = d_270_v159_
            nw42_[int(27)] = (default__.fm19(p0, d_107_v40_, globalState)) + (d_270_v159_)
            nw42_[int(28)] = _dafny.SeqWithoutIsStrInference([d_234_v131_, d_234_v131_])
            d_272_v161_ = nw42_
            index24_ = default__.safeIndex(3, (d_272_v161_).length(0))
            (d_272_v161_)[index24_] = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_277_i19_ in range(default__.abs(480))])])
            index25_ = default__.safeIndex(3, (d_272_v161_).length(0))
            (d_272_v161_)[index25_] = (d_270_v159_) + (d_270_v159_)
            d_278_v162_: _dafny.Array
            nw43_ = _dafny.Array(_dafny.Array(None, 0), 25)
            d_278_v162_ = nw43_
            d_279_v163_: _dafny.Array
            nw44_ = _dafny.Array(_dafny.Map({}), 6)
            d_279_v163_ = nw44_
            index26_ = default__.safeIndex(760, (d_278_v162_).length(0))
            (d_278_v162_)[index26_] = d_279_v163_
            index27_ = default__.safeIndex(760, (d_278_v162_).length(0))
            nw45_ = _dafny.Array(_dafny.Map({}), 15)
            (d_278_v162_)[index27_] = nw45_
            d_280_v165_: _dafny.Array
            def lambda11_(d_281_p0_, d_282_v131_, d_283_v158_, d_284_v40_):
                def lambda12_(d_285_i20_):
                    def iife25_():
                        coll19_ = _dafny.Set()
                        compr_19_: str
                        for compr_19_ in (_dafny.SeqWithoutIsStrInference([d_283_v158_, d_283_v158_])).Elements:
                            d_287_v164_: str = compr_19_
                            if (d_287_v164_) in (_dafny.SeqWithoutIsStrInference([d_283_v158_, d_283_v158_])):
                                coll19_ = coll19_.union(_dafny.Set([d_287_v164_]))
                        return _dafny.Set(coll19_)
                    return _dafny.SeqWithoutIsStrInference([D1_DC3(d_281_p0_, d_281_p0_, d_282_v131_, True, len(d_282_v131_)), D1_DC3(d_281_p0_, True, _dafny.SeqWithoutIsStrInference([d_283_v158_ for d_286_i21_ in range(default__.abs(235))]), d_281_p0_, d_284_v40_), D1_DC3(d_281_p0_, d_281_p0_, d_282_v131_, not(False), len(iife25_()
))])

                return lambda12_

            init5_ = lambda11_(p0, d_234_v131_, d_269_v158_, d_107_v40_)
            nw46_ = _dafny.Array(None, 9)
            for i0_5_ in range(nw46_.length(0)):
                nw46_[i0_5_] = init5_(i0_5_)
            d_280_v165_ = nw46_
            d_288_v166_: _dafny.Map
            d_288_v166_ = _dafny.Map({p0: _dafny.MultiSet(d_271_v160_)})
            rhs21_ = (len(d_288_v166_)) * (d_107_v40_)
            rhs22_ = d_107_v40_
            rhs23_ = p0
            rhs24_ = (d_280_v165_ if p0 else d_280_v165_)
            lhs17_ = globalState
            d_107_v40_ = rhs21_
            lhs17_.f0 = rhs22_
            r0 = rhs23_
            d_280_v165_ = rhs24_
            d_289_v167_: D2
            d_289_v167_ = D2_DC7(len(d_234_v131_))
            d_290_v168_: _dafny.Set
            d_290_v168_ = _dafny.Set({p0})
            d_291_v169_: C0
            nw47_ = C0()
            nw47_.ctor__((d_107_v40_) + ((d_289_v167_).cf15), d_290_v168_)
            d_291_v169_ = nw47_
        d_292_v170_: str
        d_292_v170_ = _dafny.CodePoint('t')
        d_293_v171_: D6
        d_293_v171_ = D6_DC19(d_107_v40_, default__.fm0(not(not(p0)), p0, d_107_v40_, globalState), d_292_v170_)
        pat_let_tv5_ = d_107_v40_
        pat_let_tv6_ = d_39_v0_
        pat_let_tv7_ = d_39_v0_
        pat_let_tv8_ = p0
        pat_let_tv9_ = d_107_v40_
        pat_let_tv10_ = d_107_v40_
        pat_let_tv11_ = d_107_v40_
        pat_let_tv12_ = d_107_v40_
        pat_let_tv13_ = d_107_v40_
        def lambda13_(source6_):
            if source6_.is_DC19:
                d_294___mcc_h22_ = source6_.cf31
                d_295___mcc_h23_ = source6_.cf32
                d_296___mcc_h24_ = source6_.cf33
                d_297_cf33_ = d_296___mcc_h24_
                d_298_cf32_ = d_295___mcc_h23_
                d_299_cf31_ = d_294___mcc_h22_
                return d_298_cf32_
            elif source6_.is_DC20:
                d_300___mcc_h25_ = source6_.cf34
                d_301_cf34_ = d_300___mcc_h25_
                def iife26_():
                    coll20_ = _dafny.Map()
                    compr_20_: D0
                    for compr_20_ in (_dafny.SeqWithoutIsStrInference([pat_let_tv6_])).Elements:
                        d_302_v172_: D0 = compr_20_
                        if (d_302_v172_) in (_dafny.SeqWithoutIsStrInference([pat_let_tv7_])):
                            coll20_[d_302_v172_] = pat_let_tv8_
                    return _dafny.Map(coll20_)
                return (pat_let_tv5_) > (len(iife26_()
                ))
            elif source6_.is_DC21:
                return (_dafny.MultiSet([14, pat_let_tv9_])).issubset(_dafny.MultiSet([pat_let_tv10_, pat_let_tv11_]))
            elif True:
                d_303___mcc_h26_ = source6_.cf30
                d_304_cf30_ = d_303___mcc_h26_
                return (D3_DC10(True, pat_let_tv12_, pat_let_tv13_)).cf18

        r0 = lambda13_(d_293_v171_)
        d_305_i22_: int
        d_305_i22_ = 0
        with _dafny.label("1"):
            while not(p0):
                with _dafny.c_label("1"):
                    if (d_305_i22_) >= (100):
                        raise _dafny.Break("1")
                    d_305_i22_ = (d_305_i22_) + (1)
                    r0 = p0
                    r0 = p0
                    (globalState).f0 = d_107_v40_
                    d_306_v173_: _dafny.Set
                    d_306_v173_ = _dafny.Set({p0, p0})
                    d_307_v174_: C0
                    nw48_ = C0()
                    nw48_.ctor__((0) - (default__.fm4(d_107_v40_, p0, globalState)), d_306_v173_)
                    d_307_v174_ = nw48_
                    pass
            pass
        r0 = p0
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_308_v0_: int
        d_308_v0_ = -45
        d_309_v1_: bool
        d_309_v1_ = True
        d_310_v2_: _dafny.Seq
        d_310_v2_ = _dafny.SeqWithoutIsStrInference([d_309_v1_, d_309_v1_])
        d_311_v3_: _dafny.Map
        d_311_v3_ = _dafny.Map({d_308_v0_: d_310_v2_})
        d_312_globalState_: GlobalState
        nw49_ = GlobalState()
        nw49_.ctor__(977, 424, True, 358, -429, d_311_v3_)
        d_312_globalState_ = nw49_
        d_313_i0_: int
        d_313_i0_ = 0
        with _dafny.label("2"):
            while (671) >= ((548 if d_309_v1_ else d_308_v0_)):
                with _dafny.c_label("2"):
                    if (d_313_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_313_i0_ = (d_313_i0_) + (1)
                    d_314_v4_: str
                    d_314_v4_ = _dafny.CodePoint('d')
                    d_315_v5_: _dafny.Seq
                    d_315_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ja"))
                    d_314_v4_ = (d_315_v5_)[default__.safeIndex(len((_dafny.SeqWithoutIsStrInference([d_309_v1_])) + (_dafny.SeqWithoutIsStrInference([d_309_v1_]))), len(d_315_v5_))]
                    d_316_v6_: _dafny.Array
                    nw50_ = _dafny.Array(int(0), 13)
                    d_316_v6_ = nw50_
                    d_317_v7_: _dafny.MultiSet
                    d_317_v7_ = _dafny.MultiSet([d_316_v6_, d_316_v6_, d_316_v6_])
                    d_318_v8_: D0
                    d_318_v8_ = D0_DC0(d_309_v1_)
                    d_319_v9_: _dafny.Array
                    nw51_ = _dafny.Array(None, 29)
                    nw51_[int(0)] = not(d_309_v1_)
                    nw51_[int(1)] = d_309_v1_
                    nw51_[int(2)] = d_309_v1_
                    nw51_[int(3)] = d_309_v1_
                    nw51_[int(4)] = d_309_v1_
                    nw51_[int(5)] = d_309_v1_
                    nw51_[int(6)] = d_309_v1_
                    nw51_[int(7)] = d_309_v1_
                    nw51_[int(8)] = True
                    nw51_[int(9)] = d_309_v1_
                    nw51_[int(10)] = d_309_v1_
                    nw51_[int(11)] = True
                    nw51_[int(12)] = d_309_v1_
                    nw51_[int(13)] = d_309_v1_
                    nw51_[int(14)] = d_309_v1_
                    nw51_[int(15)] = default__.fm0(d_309_v1_, d_309_v1_, d_308_v0_, d_312_globalState_)
                    nw51_[int(16)] = d_309_v1_
                    nw51_[int(17)] = d_309_v1_
                    nw51_[int(18)] = d_309_v1_
                    nw51_[int(19)] = not(d_309_v1_)
                    nw51_[int(20)] = False
                    nw51_[int(21)] = d_309_v1_
                    nw51_[int(22)] = d_309_v1_
                    nw51_[int(23)] = d_309_v1_
                    nw51_[int(24)] = True
                    nw51_[int(25)] = d_309_v1_
                    nw51_[int(26)] = d_309_v1_
                    nw51_[int(27)] = (d_318_v8_).cf0
                    nw51_[int(28)] = d_309_v1_
                    d_319_v9_ = nw51_
                    d_320_v10_: _dafny.Array
                    def lambda14_(d_321_v1_):
                        def lambda15_(d_322_i1_):
                            return _dafny.MultiSet([(_dafny.MultiSet([d_321_v1_])).cardinality])

                        return lambda15_

                    init6_ = lambda14_(d_309_v1_)
                    nw52_ = _dafny.Array(None, 20)
                    for i0_6_ in range(nw52_.length(0)):
                        nw52_[i0_6_] = init6_(i0_6_)
                    d_320_v10_ = nw52_
                    d_323_v11_: D0
                    d_323_v11_ = D0_DC1(d_317_v7_, d_319_v9_, d_320_v10_)
                    source7_ = d_323_v11_
                    if source7_.is_DC0:
                        d_324___mcc_h0_ = source7_.cf0
                        d_325_cf0_ = d_324___mcc_h0_
                        d_326_v12_: _dafny.Set
                        d_326_v12_ = _dafny.Set({d_308_v0_, d_308_v0_})
                        d_327_v13_: _dafny.Seq
                        d_327_v13_ = _dafny.SeqWithoutIsStrInference([d_326_v12_, d_326_v12_])
                        d_327_v13_ = d_327_v13_
                        d_328_v14_: _dafny.Set
                        d_328_v14_ = _dafny.Set({True, d_309_v1_, d_309_v1_})
                        d_329_v15_: C0
                        nw53_ = C0()
                        nw53_.ctor__(default__.safeDivisionInt(d_308_v0_, d_308_v0_), (d_328_v14_) - (d_328_v14_))
                        d_329_v15_ = nw53_
                        d_330_v16_: _dafny.Map
                        d_330_v16_ = _dafny.Map({d_309_v1_: d_329_v15_})
                        d_330_v16_ = (d_330_v16_).set(d_325_cf0_, ((d_330_v16_)[d_309_v1_] if (d_309_v1_) in (d_330_v16_) else d_329_v15_))
                        d_316_v6_ = (d_316_v6_ if True else d_316_v6_)
                    elif True:
                        d_331___mcc_h1_ = source7_.cf1
                        d_332___mcc_h2_ = source7_.cf2
                        d_333___mcc_h3_ = source7_.cf3
                        d_334_cf3_ = d_333___mcc_h3_
                        d_335_cf2_ = d_332___mcc_h2_
                        d_336_cf1_ = d_331___mcc_h1_
                        d_316_v6_ = d_316_v6_
                        d_337_v17_: _dafny.MultiSet
                        d_337_v17_ = _dafny.MultiSet([d_309_v1_])
                        d_337_v17_ = (d_337_v17_) | (d_337_v17_)
                        d_338_v18_: bool
                        out0_: bool
                        out0_ = default__.m0(not((d_309_v1_ if d_309_v1_ else d_309_v1_)), d_312_globalState_)
                        d_338_v18_ = out0_
                        d_308_v0_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nroqns")))
                    (d_312_globalState_).f4 = d_308_v0_
                    d_339_v19_: _dafny.Set
                    d_339_v19_ = _dafny.Set({d_308_v0_, (d_308_v0_) - (len(d_315_v5_)), len((d_315_v5_) + (d_315_v5_)), d_308_v0_, default__.safeDivisionInt(d_308_v0_, d_308_v0_)})
                    index28_ = default__.safeIndex(693, (d_316_v6_).length(0))
                    (d_316_v6_)[index28_] = d_308_v0_
                    index29_ = default__.safeIndex(693, (d_316_v6_).length(0))
                    rhs25_ = d_309_v1_
                    rhs26_ = d_339_v19_
                    rhs27_ = d_308_v0_
                    lhs18_ = d_316_v6_
                    lhs19_ = default__.safeIndex(693, (d_316_v6_).length(0))
                    d_309_v1_ = rhs25_
                    d_339_v19_ = rhs26_
                    lhs18_[lhs19_] = rhs27_
                    pass
            pass
        d_340_v20_: _dafny.Seq
        d_340_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ss"))
        d_341_v21_: _dafny.MultiSet
        d_341_v21_ = _dafny.MultiSet([d_308_v0_, d_308_v0_])
        source8_ = D0_DC0((len(d_340_v20_)) > ((d_341_v21_).cardinality))
        if source8_.is_DC0:
            d_342___mcc_h4_ = source8_.cf0
            d_343_cf0_ = d_342___mcc_h4_
            d_344_v22_: _dafny.Array
            nw54_ = _dafny.Array(False, 7)
            d_344_v22_ = nw54_
            index30_ = default__.safeIndex(676, (d_344_v22_).length(0))
            (d_344_v22_)[index30_] = (d_343_cf0_) or (True)
            d_345_v23_: str
            d_345_v23_ = _dafny.CodePoint('s')
            d_346_v24_: _dafny.MultiSet
            d_346_v24_ = _dafny.MultiSet([default__.fm3(d_310_v2_, d_312_globalState_)])
            index31_ = default__.safeIndex(676, (d_344_v22_).length(0))
            rhs28_ = d_309_v1_
            rhs29_ = d_345_v23_
            rhs30_ = d_308_v0_
            rhs31_ = (d_346_v24_).cardinality
            lhs20_ = d_344_v22_
            lhs21_ = default__.safeIndex(676, (d_344_v22_).length(0))
            lhs22_ = d_312_globalState_
            lhs23_ = d_312_globalState_
            lhs20_[lhs21_] = rhs28_
            d_345_v23_ = rhs29_
            lhs22_.f4 = rhs30_
            lhs23_.f0 = rhs31_
            d_347_v25_: _dafny.Map
            d_347_v25_ = _dafny.Map({d_343_cf0_: default__.safeDivisionInt(d_308_v0_, d_308_v0_)})
            d_348_v26_: _dafny.MultiSet
            d_348_v26_ = _dafny.MultiSet([(d_344_v22_)[default__.safeIndex(676, (d_344_v22_).length(0))]])
            d_347_v25_ = (d_347_v25_).set((_dafny.MultiSet(d_310_v2_)).isdisjoint((d_348_v26_).set(not(not(True)), default__.abs(d_308_v0_))), (0) - ((0) - (d_308_v0_)))
            d_345_v23_ = d_345_v23_
            d_349_v27_: _dafny.Map
            d_349_v27_ = _dafny.Map({(d_344_v22_)[default__.safeIndex(676, (d_344_v22_).length(0))]: _dafny.Set({d_309_v1_, (d_344_v22_)[default__.safeIndex(676, (d_344_v22_).length(0))], d_309_v1_, (d_344_v22_)[default__.safeIndex(676, (d_344_v22_).length(0))], d_343_cf0_})})
            d_350_v28_: _dafny.Set
            d_350_v28_ = _dafny.Set({False, False, d_343_cf0_})
            d_351_v29_: _dafny.Array
            nw55_ = _dafny.Array(None, 22)
            nw55_[int(0)] = d_308_v0_
            nw55_[int(1)] = d_308_v0_
            nw55_[int(2)] = d_308_v0_
            nw55_[int(3)] = d_308_v0_
            nw55_[int(4)] = (d_308_v0_) - (default__.fm4(d_308_v0_, d_343_cf0_, d_312_globalState_))
            nw55_[int(5)] = d_308_v0_
            nw55_[int(6)] = d_308_v0_
            nw55_[int(7)] = len((d_340_v20_) + (d_340_v20_))
            nw55_[int(8)] = default__.safeDivisionInt(d_308_v0_, d_308_v0_)
            nw55_[int(9)] = d_308_v0_
            nw55_[int(10)] = d_308_v0_
            nw55_[int(11)] = 233
            nw55_[int(12)] = default__.safeDivisionInt(d_308_v0_, d_308_v0_)
            nw55_[int(13)] = d_308_v0_
            nw55_[int(14)] = (d_308_v0_) - (d_308_v0_)
            nw55_[int(15)] = (0) - ((0) - (d_308_v0_))
            nw55_[int(16)] = d_308_v0_
            nw55_[int(17)] = (len((d_349_v27_).set(d_309_v1_, d_350_v28_))) * (d_308_v0_)
            nw55_[int(18)] = default__.safeModuloInt(d_308_v0_, d_308_v0_)
            nw55_[int(19)] = d_308_v0_
            nw55_[int(20)] = d_308_v0_
            nw55_[int(21)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oywkaa"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))))
            d_351_v29_ = nw55_
            index32_ = default__.safeIndex(342, (d_351_v29_).length(0))
            (d_351_v29_)[index32_] = d_308_v0_
            index33_ = default__.safeIndex(342, (d_351_v29_).length(0))
            (d_351_v29_)[index33_] = default__.safeModuloInt(d_308_v0_, d_308_v0_)
        elif True:
            d_352___mcc_h5_ = source8_.cf1
            d_353___mcc_h6_ = source8_.cf2
            d_354___mcc_h7_ = source8_.cf3
            d_355_cf3_ = d_354___mcc_h7_
            d_356_cf2_ = d_353___mcc_h6_
            d_357_cf1_ = d_352___mcc_h5_
            (d_312_globalState_).f4 = d_308_v0_
            d_358_v30_: _dafny.Set
            d_358_v30_ = _dafny.Set({d_309_v1_})
            d_359_v31_: C0
            nw56_ = C0()
            nw56_.ctor__(d_308_v0_, d_358_v30_)
            d_359_v31_ = nw56_
            d_360_v32_: _dafny.Array
            nw57_ = _dafny.Array(int(0), 14)
            d_360_v32_ = nw57_
            d_361_v33_: _dafny.Map
            d_361_v33_ = _dafny.Map({d_359_v31_: d_360_v32_})
            d_361_v33_ = (_dafny.Map({d_359_v31_: d_360_v32_})) | (d_361_v33_)
            d_362_v34_: str
            d_362_v34_ = _dafny.CodePoint('t')
            d_363_v35_: _dafny.Array
            nw58_ = _dafny.Array(None, 16)
            nw58_[int(0)] = d_360_v32_
            nw58_[int(1)] = d_360_v32_
            nw58_[int(2)] = d_360_v32_
            nw58_[int(3)] = d_360_v32_
            nw58_[int(4)] = d_360_v32_
            nw58_[int(5)] = d_360_v32_
            nw58_[int(6)] = d_360_v32_
            nw58_[int(7)] = d_360_v32_
            nw58_[int(8)] = d_360_v32_
            nw58_[int(9)] = d_360_v32_
            nw58_[int(10)] = d_360_v32_
            nw58_[int(11)] = d_360_v32_
            nw58_[int(12)] = (d_360_v32_ if d_309_v1_ else d_360_v32_)
            nw58_[int(13)] = d_360_v32_
            nw58_[int(14)] = d_360_v32_
            nw58_[int(15)] = d_360_v32_
            d_363_v35_ = nw58_
            index34_ = default__.safeIndex(372, (d_363_v35_).length(0))
            (d_363_v35_)[index34_] = d_360_v32_
            d_364_v36_: _dafny.Map
            d_364_v36_ = _dafny.Map({d_360_v32_: d_355_cf3_})
            d_365_v37_: D0
            d_365_v37_ = D0_DC1(d_357_cf1_, d_356_cf2_, ((d_364_v36_)[d_360_v32_] if (d_360_v32_) in (d_364_v36_) else d_355_cf3_))
            d_366_v38_: _dafny.Seq
            d_366_v38_ = _dafny.SeqWithoutIsStrInference([d_356_cf2_])
            d_367_v39_: _dafny.Array
            nw59_ = _dafny.Array(None, 28)
            nw59_[int(0)] = d_356_cf2_
            nw59_[int(1)] = d_356_cf2_
            nw59_[int(2)] = d_356_cf2_
            nw59_[int(3)] = d_356_cf2_
            nw59_[int(4)] = d_356_cf2_
            nw59_[int(5)] = d_356_cf2_
            nw59_[int(6)] = d_356_cf2_
            nw59_[int(7)] = d_356_cf2_
            nw59_[int(8)] = d_356_cf2_
            nw59_[int(9)] = d_356_cf2_
            nw59_[int(10)] = (d_356_cf2_ if d_309_v1_ else d_356_cf2_)
            nw59_[int(11)] = d_356_cf2_
            nw59_[int(12)] = (d_365_v37_).cf2
            nw59_[int(13)] = d_356_cf2_
            nw59_[int(14)] = d_356_cf2_
            nw59_[int(15)] = d_356_cf2_
            nw59_[int(16)] = d_356_cf2_
            nw59_[int(17)] = d_356_cf2_
            nw59_[int(18)] = d_356_cf2_
            nw59_[int(19)] = d_356_cf2_
            nw59_[int(20)] = d_356_cf2_
            nw59_[int(21)] = d_356_cf2_
            nw59_[int(22)] = d_356_cf2_
            nw59_[int(23)] = d_356_cf2_
            nw59_[int(24)] = d_356_cf2_
            nw59_[int(25)] = d_356_cf2_
            nw59_[int(26)] = d_356_cf2_
            nw59_[int(27)] = (d_366_v38_)[default__.safeIndex(d_308_v0_, len(d_366_v38_))]
            d_367_v39_ = nw59_
            d_368_v40_: _dafny.MultiSet
            d_368_v40_ = _dafny.MultiSet([False, d_309_v1_])
            index35_ = default__.safeIndex(372, (d_363_v35_).length(0))
            rhs32_ = d_362_v34_
            rhs33_ = d_360_v32_
            rhs34_ = d_309_v1_
            rhs35_ = d_367_v39_
            rhs36_ = ((d_368_v40_ if not(d_309_v1_) else ((d_368_v40_).set(True, default__.abs(-722)) if d_309_v1_ else d_368_v40_))).cardinality
            lhs24_ = d_363_v35_
            lhs25_ = default__.safeIndex(372, (d_363_v35_).length(0))
            d_362_v34_ = rhs32_
            lhs24_[lhs25_] = rhs33_
            d_309_v1_ = rhs34_
            d_367_v39_ = rhs35_
            d_308_v0_ = rhs36_
            d_369_v42_: _dafny.Set
            d_369_v42_ = _dafny.Set({(0) - ((d_359_v31_).f6)})
            def iife27_():
                coll21_ = _dafny.Set()
                compr_21_: int
                for compr_21_ in _dafny.IntegerRange(-485, 640):
                    d_370_v41_: int = compr_21_
                    if ((-485) <= (d_370_v41_)) and ((d_370_v41_) < (640)):
                        coll21_ = coll21_.union(_dafny.Set([(d_370_v41_) - (len(_dafny.Map({d_362_v34_: d_309_v1_})))]))
                return _dafny.Set(coll21_)
            d_309_v1_ = ((iife27_()
            ).intersection(d_369_v42_)).ispropersubset(d_369_v42_)
        d_371_v43_: D1
        d_371_v43_ = D1_DC3(d_309_v1_, d_309_v1_, d_340_v20_, d_309_v1_, d_308_v0_)
        pat_let_tv14_ = d_309_v1_
        pat_let_tv15_ = d_309_v1_
        pat_let_tv16_ = d_309_v1_
        pat_let_tv17_ = d_309_v1_
        def lambda16_(source9_):
            if source9_.is_DC3:
                d_372___mcc_h8_ = source9_.cf5
                d_373___mcc_h9_ = source9_.cf6
                d_374___mcc_h10_ = source9_.cf7
                d_375___mcc_h11_ = source9_.cf8
                d_376___mcc_h12_ = source9_.cf9
                d_377_cf9_ = d_376___mcc_h12_
                d_378_cf8_ = d_375___mcc_h11_
                d_379_cf7_ = d_374___mcc_h10_
                d_380_cf6_ = d_373___mcc_h9_
                d_381_cf5_ = d_372___mcc_h8_
                return pat_let_tv14_
            elif source9_.is_DC2:
                d_382___mcc_h13_ = source9_.cf4
                d_383_cf4_ = d_382___mcc_h13_
                return pat_let_tv15_
            elif True:
                d_384___mcc_h14_ = source9_.cf10
                d_385_cf10_ = d_384___mcc_h14_
                if pat_let_tv16_:
                    return True
                elif True:
                    return pat_let_tv17_

        if lambda16_(d_371_v43_):
            (d_312_globalState_).f4 = d_308_v0_
            d_386_v44_: _dafny.MultiSet
            d_386_v44_ = _dafny.MultiSet([d_309_v1_, d_309_v1_, d_309_v1_])
            d_387_v45_: _dafny.Map
            d_387_v45_ = _dafny.Map({d_340_v20_: d_309_v1_})
            d_388_v46_: C0
            nw60_ = C0()
            nw60_.ctor__((d_386_v44_).cardinality, _dafny.Set({not(((d_387_v45_)[d_340_v20_] if (d_340_v20_) in (d_387_v45_) else d_309_v1_)), d_309_v1_}))
            d_388_v46_ = nw60_
            d_389_v47_: _dafny.Array
            nw61_ = _dafny.Array(False, 24)
            d_389_v47_ = nw61_
            d_389_v47_ = d_389_v47_
            (d_312_globalState_).f4 = d_308_v0_
            d_390_v48_: str
            d_390_v48_ = _dafny.CodePoint('b')
            d_387_v45_ = (_dafny.Map({(d_340_v20_).set(default__.safeIndex(d_308_v0_, len(d_340_v20_)), d_390_v48_): False})).set(_dafny.SeqWithoutIsStrInference([d_390_v48_ for d_391_i2_ in range(default__.abs(896))]), d_309_v1_)
        elif True:
            d_392_v49_: _dafny.MultiSet
            d_392_v49_ = _dafny.MultiSet([d_309_v1_, d_309_v1_])
            if (((d_392_v49_ if d_309_v1_ else d_392_v49_)).cardinality) <= (d_308_v0_):
                d_309_v1_ = not(d_309_v1_)
                d_309_v1_ = d_309_v1_
                d_309_v1_ = False
                d_393_v50_: _dafny.Array
                nw62_ = _dafny.Array(None, 3)
                nw62_[int(0)] = d_308_v0_
                nw62_[int(1)] = (274) - ((d_371_v43_).cf9)
                nw62_[int(2)] = (d_308_v0_) - (d_308_v0_)
                d_393_v50_ = nw62_
                index36_ = default__.safeIndex(812, (d_393_v50_).length(0))
                (d_393_v50_)[index36_] = d_308_v0_
                d_394_v51_: _dafny.Map
                d_394_v51_ = _dafny.Map({default__.fm4(d_308_v0_, d_309_v1_, d_312_globalState_): d_309_v1_})
                index37_ = default__.safeIndex(812, (d_393_v50_).length(0))
                (d_393_v50_)[index37_] = len(d_394_v51_)
                d_395_v52_: _dafny.Array
                nw63_ = _dafny.Array(False, 16)
                d_395_v52_ = nw63_
                d_396_v53_: _dafny.Map
                d_396_v53_ = _dafny.Map({d_392_v49_: d_395_v52_})
                (d_312_globalState_).f0 = len(d_396_v53_)
            elif True:
                d_309_v1_ = False
                d_397_v54_: _dafny.Array
                nw64_ = _dafny.Array(None, 7)
                nw64_[int(0)] = (d_308_v0_) >= (d_308_v0_)
                nw64_[int(1)] = d_309_v1_
                nw64_[int(2)] = d_309_v1_
                nw64_[int(3)] = d_309_v1_
                nw64_[int(4)] = d_309_v1_
                nw64_[int(5)] = d_309_v1_
                nw64_[int(6)] = d_309_v1_
                d_397_v54_ = nw64_
                d_397_v54_ = (d_397_v54_ if False else d_397_v54_)
                d_398_v55_: _dafny.Seq
                d_398_v55_ = _dafny.SeqWithoutIsStrInference([d_308_v0_])
                d_399_v56_: str
                d_399_v56_ = _dafny.CodePoint('f')
                d_400_v57_: _dafny.Set
                d_400_v57_ = _dafny.Set({d_399_v56_, d_399_v56_, d_399_v56_})
                d_401_v58_: _dafny.Seq
                d_401_v58_ = _dafny.SeqWithoutIsStrInference([len(d_398_v55_), d_308_v0_, d_308_v0_, d_308_v0_, len(d_400_v57_)])
                d_402_v59_: C0
                nw65_ = C0()
                nw65_.ctor__((d_398_v55_)[default__.safeIndex(default__.fm4(d_308_v0_, d_309_v1_, d_312_globalState_), len(d_398_v55_))], _dafny.Set({not(d_309_v1_)}))
                d_402_v59_ = nw65_
                (d_312_globalState_).f0 = d_308_v0_
                (d_312_globalState_).f4 = (d_308_v0_) * ((d_371_v43_).cf9)
            d_403_v60_: _dafny.Map
            d_403_v60_ = _dafny.Map({d_308_v0_: (d_308_v0_) + (d_308_v0_)})
            d_404_v61_: _dafny.Set
            d_404_v61_ = _dafny.Set({(0) - (d_308_v0_), d_308_v0_})
            d_403_v60_ = _dafny.Map({len(d_404_v61_): d_308_v0_})
            d_405_v62_: _dafny.Set
            d_405_v62_ = _dafny.Set({not(d_309_v1_)})
            d_406_v63_: C0
            nw66_ = C0()
            nw66_.ctor__(default__.safeModuloInt(d_308_v0_, d_308_v0_), (d_405_v62_ if d_309_v1_ else d_405_v62_))
            d_406_v63_ = nw66_
            d_407_v64_: str
            d_407_v64_ = _dafny.CodePoint('u')
            (d_312_globalState_).f0 = (177 if (d_407_v64_) not in (d_340_v20_) else ((d_406_v63_).f6) - ((d_406_v63_).f6))
            if True:
                d_408_v65_: bool
                out1_: bool
                out1_ = default__.m0(d_309_v1_, d_312_globalState_)
                d_408_v65_ = out1_
                d_309_v1_ = d_408_v65_
                d_409_v66_: _dafny.Array
                nw67_ = _dafny.Array(int(0), 28)
                d_409_v66_ = nw67_
                d_410_v67_: _dafny.Seq
                d_410_v67_ = _dafny.SeqWithoutIsStrInference([d_409_v66_, d_409_v66_, d_409_v66_])
                pat_let_tv18_ = d_409_v66_
                d_411_v68_: _dafny.Array
                nw68_ = _dafny.Array(None, 22)
                nw68_[int(0)] = d_409_v66_
                nw68_[int(1)] = d_409_v66_
                nw68_[int(2)] = (d_410_v67_)[default__.safeIndex((d_406_v63_).f6, len(d_410_v67_))]
                nw68_[int(3)] = d_409_v66_
                nw68_[int(4)] = d_409_v66_
                nw68_[int(5)] = d_409_v66_
                nw68_[int(6)] = (d_410_v67_)[default__.safeIndex(default__.fm4((d_406_v63_).f6, d_408_v65_, d_312_globalState_), len(d_410_v67_))]
                nw68_[int(7)] = d_409_v66_
                nw68_[int(8)] = d_409_v66_
                nw68_[int(9)] = d_409_v66_
                nw68_[int(10)] = (d_410_v67_)[default__.safeIndex((d_406_v63_).f6, len(d_410_v67_))]
                nw68_[int(11)] = d_409_v66_
                nw68_[int(12)] = d_409_v66_
                nw68_[int(13)] = d_409_v66_
                nw68_[int(14)] = d_409_v66_
                nw68_[int(15)] = d_409_v66_
                nw68_[int(16)] = d_409_v66_
                def iife28_(_pat_let3_0):
                    def iife29_(d_412_dt__update__tmp_h0_):
                        def iife30_(_pat_let4_0):
                            def iife31_(d_413_dt__update_hcf11_h0_):
                                return D2_DC5(d_413_dt__update_hcf11_h0_)
                            return iife31_(_pat_let4_0)
                        return iife30_(pat_let_tv18_)
                    return iife29_(_pat_let3_0)
                nw68_[int(17)] = (iife28_(D2_DC5(d_409_v66_))).cf11
                nw68_[int(18)] = d_409_v66_
                nw68_[int(19)] = d_409_v66_
                nw68_[int(20)] = d_409_v66_
                nw68_[int(21)] = d_409_v66_
                d_411_v68_ = nw68_
                index38_ = default__.safeIndex(915, (d_411_v68_).length(0))
                (d_411_v68_)[index38_] = d_409_v66_
                index39_ = default__.safeIndex(915, (d_411_v68_).length(0))
                (d_411_v68_)[index39_] = d_409_v66_
                d_414_v69_: _dafny.Array
                nw69_ = _dafny.Array(None, 28)
                nw69_[int(0)] = d_408_v65_
                nw69_[int(1)] = default__.fm0(d_309_v1_, (D1_DC3((d_371_v43_).cf6, d_408_v65_, _dafny.SeqWithoutIsStrInference([d_407_v64_ for d_415_i3_ in range(default__.abs(997))]), d_408_v65_, 639)).cf5, (d_406_v63_).f6, d_312_globalState_)
                nw69_[int(2)] = d_309_v1_
                nw69_[int(3)] = d_309_v1_
                nw69_[int(4)] = d_309_v1_
                nw69_[int(5)] = d_309_v1_
                nw69_[int(6)] = False
                nw69_[int(7)] = (d_371_v43_).cf5
                nw69_[int(8)] = d_309_v1_
                nw69_[int(9)] = d_309_v1_
                nw69_[int(10)] = d_309_v1_
                nw69_[int(11)] = not(d_408_v65_)
                nw69_[int(12)] = d_408_v65_
                nw69_[int(13)] = d_408_v65_
                nw69_[int(14)] = d_309_v1_
                nw69_[int(15)] = d_408_v65_
                nw69_[int(16)] = d_309_v1_
                nw69_[int(17)] = d_408_v65_
                nw69_[int(18)] = True
                nw69_[int(19)] = d_408_v65_
                nw69_[int(20)] = d_309_v1_
                nw69_[int(21)] = d_408_v65_
                nw69_[int(22)] = d_408_v65_
                nw69_[int(23)] = True
                nw69_[int(24)] = False
                nw69_[int(25)] = d_408_v65_
                nw69_[int(26)] = d_408_v65_
                nw69_[int(27)] = d_309_v1_
                d_414_v69_ = nw69_
                d_416_v70_: _dafny.Map
                d_416_v70_ = _dafny.Map({d_414_v69_: d_408_v65_})
                d_417_v71_: C0
                nw70_ = C0()
                nw70_.ctor__((((d_392_v49_)[d_408_v65_] if (d_408_v65_) in (d_392_v49_) else (d_406_v63_).f6)) * (len(d_416_v70_)), d_405_v62_)
                d_417_v71_ = nw70_
                d_418_v72_: _dafny.Map
                d_418_v72_ = _dafny.Map({True: (d_341_v21_).cardinality})
                (d_312_globalState_).f4 = len(d_418_v72_)
            elif True:
                d_419_v73_: C0
                nw71_ = C0()
                nw71_.ctor__(483, (d_406_v63_).f7)
                d_419_v73_ = nw71_
                d_420_v74_: _dafny.Array
                nw72_ = _dafny.Array(False, 19)
                d_420_v74_ = nw72_
                d_421_v75_: _dafny.Map
                d_421_v75_ = _dafny.Map({d_406_v63_: d_420_v74_})
                d_422_v76_: _dafny.Map
                d_422_v76_ = _dafny.Map({d_308_v0_: d_420_v74_})
                d_421_v75_ = (d_421_v75_).set(d_419_v73_, ((d_422_v76_)[(d_406_v63_).f6] if ((d_406_v63_).f6) in (d_422_v76_) else ((d_422_v76_)[d_308_v0_] if (d_308_v0_) in (d_422_v76_) else d_420_v74_)))
                d_423_v77_: _dafny.Array
                def lambda17_(d_424_v63_):
                    def lambda18_(d_425_i4_):
                        return default__.safeModuloInt(d_425_i4_, (d_424_v63_).f6)

                    return lambda18_

                init7_ = lambda17_(d_406_v63_)
                nw73_ = _dafny.Array(None, 1)
                for i0_7_ in range(nw73_.length(0)):
                    nw73_[i0_7_] = init7_(i0_7_)
                d_423_v77_ = nw73_
                index40_ = default__.safeIndex(378, (d_423_v77_).length(0))
                (d_423_v77_)[index40_] = (d_406_v63_).f6
                index41_ = default__.safeIndex(229, (d_423_v77_).length(0))
                (d_423_v77_)[index41_] = d_308_v0_
                d_426_v78_: _dafny.Map
                d_426_v78_ = _dafny.Map({(d_404_v61_) | (_dafny.Set({858})): (d_340_v20_).set(default__.safeIndex((d_419_v73_).f6, len(d_340_v20_)), d_407_v64_)})
                pat_let_tv19_ = d_309_v1_
                pat_let_tv20_ = d_309_v1_
                index42_ = default__.safeIndex(378, (d_423_v77_).length(0))
                index43_ = default__.safeIndex(229, (d_423_v77_).length(0))
                def iife32_(_pat_let5_0):
                    def iife33_(d_427_dt__update__tmp_h1_):
                        def iife34_(_pat_let6_0):
                            def iife35_(d_428_dt__update_hcf6_h0_):
                                def iife36_(_pat_let7_0):
                                    def iife37_(d_429_dt__update_hcf8_h0_):
                                        return D1_DC3((d_427_dt__update__tmp_h1_).cf5, d_428_dt__update_hcf6_h0_, (d_427_dt__update__tmp_h1_).cf7, d_429_dt__update_hcf8_h0_, (d_427_dt__update__tmp_h1_).cf9)
                                    return iife37_(_pat_let7_0)
                                return iife36_(pat_let_tv20_)
                            return iife35_(_pat_let6_0)
                        return iife34_(pat_let_tv19_)
                    return iife33_(_pat_let5_0)
                rhs37_ = ((d_426_v78_)[(d_404_v61_) | (d_404_v61_)] if ((d_404_v61_) | (d_404_v61_)) in (d_426_v78_) else (iife32_(d_371_v43_)).cf7)
                rhs38_ = (d_406_v63_).f6
                rhs39_ = default__.fm4(d_308_v0_, (d_340_v20_) == (d_340_v20_), d_312_globalState_)
                rhs40_ = ((d_406_v63_).f6) + (d_308_v0_)
                rhs41_ = (d_406_v63_).f6
                lhs26_ = d_423_v77_
                lhs27_ = default__.safeIndex(378, (d_423_v77_).length(0))
                lhs28_ = d_312_globalState_
                lhs29_ = d_423_v77_
                lhs30_ = default__.safeIndex(229, (d_423_v77_).length(0))
                lhs31_ = d_312_globalState_
                d_340_v20_ = rhs37_
                lhs26_[lhs27_] = rhs38_
                lhs28_.f4 = rhs39_
                lhs29_[lhs30_] = rhs40_
                lhs31_.f4 = rhs41_
                d_430_v79_: _dafny.Set
                d_430_v79_ = _dafny.Set({d_407_v64_})
                (d_312_globalState_).f4 = ((d_423_v77_)[default__.safeIndex(378, (d_423_v77_).length(0))]) * (len((d_430_v79_).intersection(d_430_v79_)))
                d_431_v80_: _dafny.Seq
                d_431_v80_ = _dafny.SeqWithoutIsStrInference([(d_406_v63_).f7, _dafny.Set({d_309_v1_, d_309_v1_, d_309_v1_})])
                d_432_v81_: _dafny.Map
                d_432_v81_ = _dafny.Map({d_309_v1_: d_308_v0_})
                d_433_v82_: _dafny.Seq
                d_433_v82_ = _dafny.SeqWithoutIsStrInference([-799])
                d_434_v83_: C0
                nw74_ = C0()
                nw74_.ctor__(-604, (d_431_v80_)[default__.safeIndex((0) - (((d_432_v81_)[d_309_v1_] if (d_309_v1_) in (d_432_v81_) else len(d_433_v82_))), len(d_431_v80_))])
                d_434_v83_ = nw74_
        d_435_v84_: _dafny.Map
        d_435_v84_ = _dafny.Map({d_309_v1_: len(d_340_v20_)})
        d_436_v85_: _dafny.Map
        d_436_v85_ = _dafny.Map({492: d_435_v84_})
        d_437_v86_: _dafny.Map
        d_437_v86_ = _dafny.Map({len(d_436_v85_): d_435_v84_})
        d_438_v87_: _dafny.MultiSet
        d_438_v87_ = _dafny.MultiSet([d_309_v1_, d_309_v1_, False, (len(d_436_v85_)) != ((0) - (d_308_v0_)), not (d_309_v1_) or (d_309_v1_)])
        d_438_v87_ = (d_438_v87_) - (d_438_v87_)
        d_309_v1_ = d_309_v1_
        d_439_v88_: _dafny.Array
        nw75_ = _dafny.Array(None, 7)
        d_439_v88_ = nw75_
        d_440_v89_: _dafny.Set
        d_440_v89_ = _dafny.Set({d_309_v1_})
        rhs42_ = default__.safeModuloInt(len(d_440_v89_), (d_308_v0_) * (578))
        rhs43_ = False
        rhs44_ = d_439_v88_
        rhs45_ = not(not(d_309_v1_))
        rhs46_ = d_309_v1_
        lhs32_ = d_312_globalState_
        lhs32_.f0 = rhs42_
        d_309_v1_ = rhs43_
        d_439_v88_ = rhs44_
        d_309_v1_ = rhs45_
        d_309_v1_ = rhs46_
        d_441_v90_: bool
        out2_: bool
        out2_ = default__.m0(d_309_v1_, d_312_globalState_)
        d_441_v90_ = out2_
        d_442_v91_: _dafny.Map
        d_442_v91_ = _dafny.Map({d_308_v0_: True})
        d_443_v92_: _dafny.Map
        d_443_v92_ = _dafny.Map({_dafny.CodePoint('v'): d_308_v0_})
        d_444_v93_: D1
        d_444_v93_ = D1_DC2(d_308_v0_)
        d_445_v94_: _dafny.Array
        nw76_ = _dafny.Array(None, 8)
        nw76_[int(0)] = d_438_v87_
        nw76_[int(1)] = (d_438_v87_) - (d_438_v87_)
        nw76_[int(2)] = d_438_v87_
        nw76_[int(3)] = _dafny.MultiSet([default__.fm0(not(d_309_v1_), ((d_442_v91_)[(d_371_v43_).cf9] if ((d_371_v43_).cf9) in (d_442_v91_) else d_441_v90_), len(d_443_v92_), d_312_globalState_), d_309_v1_, d_309_v1_, d_309_v1_, d_309_v1_])
        nw76_[int(4)] = (d_438_v87_) - (d_438_v87_)
        nw76_[int(5)] = (d_438_v87_).set(d_309_v1_, default__.abs((d_444_v93_).cf4))
        nw76_[int(6)] = _dafny.MultiSet([d_441_v90_, d_441_v90_])
        nw76_[int(7)] = d_438_v87_
        d_445_v94_ = nw76_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_445_v94_).length(0)):
            d_446_i5_: int = guard_loop_0_
            if (True) and (((0) <= (d_446_i5_)) and ((d_446_i5_) < ((d_445_v94_).length(0)))):
                (d_445_v94_)[(d_446_i5_)] = d_438_v87_
        d_340_v20_ = (d_340_v20_) + (d_340_v20_)
        source10_ = d_444_v93_
        if source10_.is_DC3:
            d_447___mcc_h15_ = source10_.cf5
            d_448___mcc_h16_ = source10_.cf6
            d_449___mcc_h17_ = source10_.cf7
            d_450___mcc_h18_ = source10_.cf8
            d_451___mcc_h19_ = source10_.cf9
            d_452_cf9_ = d_451___mcc_h19_
            d_453_cf8_ = d_450___mcc_h18_
            d_454_cf7_ = d_449___mcc_h17_
            d_455_cf6_ = d_448___mcc_h16_
            d_456_cf5_ = d_447___mcc_h15_
            d_457_v95_: _dafny.Map
            d_457_v95_ = _dafny.Map({d_455_cf6_: False})
            def iife38_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in _dafny.IntegerRange(954, -549):
                    d_458_v96_: int = compr_22_
                    if ((954) <= (d_458_v96_)) and ((d_458_v96_) < (-549)):
                        coll22_[default__.safeModuloInt(d_458_v96_, len(d_457_v95_))] = d_456_cf5_
                return _dafny.Map(coll22_)
            d_457_v95_ = (d_457_v95_).set((iife38_()
            ) != (d_442_v91_), d_456_cf5_)
            d_459_v97_: str
            d_459_v97_ = _dafny.CodePoint('r')
            d_459_v97_ = d_459_v97_
            (d_312_globalState_).f4 = (d_308_v0_) + (len(d_454_cf7_))
            d_455_cf6_ = ((d_457_v95_)[(d_453_cf8_) == (d_453_cf8_)] if ((d_453_cf8_) == (d_453_cf8_)) in (d_457_v95_) else True)
        elif source10_.is_DC2:
            d_460___mcc_h20_ = source10_.cf4
            d_461_cf4_ = d_460___mcc_h20_
            d_440_v89_ = default__.fm5(d_312_globalState_)
            d_462_v98_: _dafny.Seq
            d_462_v98_ = _dafny.SeqWithoutIsStrInference([len(d_310_v2_)])
            if not(((d_462_v98_)[default__.safeIndex(d_308_v0_, len(d_462_v98_))]) != (d_461_cf4_)):
                d_463_v99_: _dafny.Map
                d_463_v99_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_464_i6_ in range(default__.abs(-19))]): d_310_v2_})
                d_465_v101_: str
                d_465_v101_ = _dafny.CodePoint('d')
                d_466_v102_: _dafny.Map
                d_466_v102_ = _dafny.Map({d_308_v0_: d_465_v101_})
                d_467_v103_: _dafny.Seq
                def iife39_():
                    coll23_ = _dafny.Map()
                    compr_23_: _dafny.Seq
                    for compr_23_ in (_dafny.MultiSet([d_340_v20_, d_340_v20_])).Elements:
                        d_468_v100_: _dafny.Seq = compr_23_
                        if (d_468_v100_) in (_dafny.MultiSet([d_340_v20_, d_340_v20_])):
                            coll23_[d_468_v100_] = _dafny.SeqWithoutIsStrInference([True])
                    return _dafny.Map(coll23_)
                d_467_v103_ = _dafny.SeqWithoutIsStrInference([d_463_v99_, d_463_v99_, d_463_v99_, iife39_()
                , _dafny.Map({d_340_v20_: (d_310_v2_).set(default__.safeIndex(len(d_466_v102_), len(d_310_v2_)), d_309_v1_)})])
                d_469_v105_: _dafny.Map
                def iife40_():
                    coll24_ = _dafny.Map()
                    compr_24_: _dafny.Seq
                    for compr_24_ in (_dafny.Map({d_340_v20_: d_309_v1_})).keys.Elements:
                        d_470_v104_: _dafny.Seq = compr_24_
                        if (d_470_v104_) in (_dafny.Map({d_340_v20_: d_309_v1_})):
                            coll24_[d_470_v104_] = d_310_v2_
                    return _dafny.Map(coll24_)
                d_469_v105_ = _dafny.Map({d_308_v0_: (((d_467_v103_)[default__.safeIndex(d_308_v0_, len(d_467_v103_))]).set(d_340_v20_, d_310_v2_) if d_441_v90_ else iife40_()
                )})
                d_469_v105_ = (d_469_v105_).set(d_461_cf4_, ((d_463_v99_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kb")), d_310_v2_)) | (d_463_v99_))
                d_309_v1_ = not (d_309_v1_) or (d_309_v1_)
                d_471_v106_: _dafny.Seq
                d_471_v106_ = _dafny.SeqWithoutIsStrInference([d_435_v84_])
                d_472_v107_: _dafny.Seq
                d_472_v107_ = _dafny.SeqWithoutIsStrInference([(d_471_v106_)[default__.safeIndex(d_308_v0_, len(d_471_v106_))], d_435_v84_, d_435_v84_, d_435_v84_])
                d_473_v108_: C0
                nw77_ = C0()
                nw77_.ctor__(len(((d_472_v107_)[default__.safeIndex(len(d_340_v20_), len(d_472_v107_))]) | (d_435_v84_)), d_440_v89_)
                d_473_v108_ = nw77_
                (d_312_globalState_).f4 = default__.safeDivisionInt(len(d_340_v20_), d_308_v0_)
                d_474_v109_: D2
                d_474_v109_ = D2_DC7((d_461_cf4_) * ((d_473_v108_).f6))
                pat_let_tv21_ = d_473_v108_
                def iife41_(_pat_let8_0):
                    def iife42_(d_475_dt__update__tmp_h2_):
                        def iife43_(_pat_let9_0):
                            def iife44_(d_476_dt__update_hcf15_h0_):
                                return D2_DC7(d_476_dt__update_hcf15_h0_)
                            return iife44_(_pat_let9_0)
                        return iife43_((pat_let_tv21_).f6)
                    return iife42_(_pat_let8_0)
                d_474_v109_ = iife41_(d_474_v109_)
            elif True:
                d_477_v110_: _dafny.Set
                d_477_v110_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_478_i7_ in range(default__.abs(-884))])), d_308_v0_})
                d_479_v111_: _dafny.Map
                d_479_v111_ = _dafny.Map({d_308_v0_: (d_477_v110_) | (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_441_v90_])), d_308_v0_}))})
                d_479_v111_ = (d_479_v111_).set(d_308_v0_, d_477_v110_)
                d_480_v112_: _dafny.Array
                nw78_ = _dafny.Array(None, 2)
                nw78_[int(0)] = d_308_v0_
                nw78_[int(1)] = d_461_cf4_
                d_480_v112_ = nw78_
                index44_ = default__.safeIndex(97, (d_480_v112_).length(0))
                (d_480_v112_)[index44_] = (301) + (len(d_340_v20_))
                d_481_v113_: _dafny.Seq
                d_481_v113_ = _dafny.SeqWithoutIsStrInference([D2_DC5(d_480_v112_)])
                index45_ = default__.safeIndex(97, (d_480_v112_).length(0))
                rhs47_ = (0) - ((D2_DC6(d_308_v0_, d_309_v1_, d_480_v112_)).cf12)
                rhs48_ = _dafny.SeqWithoutIsStrInference([D2_DC5(d_480_v112_)])
                rhs49_ = default__.safeDivisionInt(len(d_440_v89_), d_308_v0_)
                lhs33_ = d_480_v112_
                lhs34_ = default__.safeIndex(97, (d_480_v112_).length(0))
                lhs33_[lhs34_] = rhs47_
                d_481_v113_ = rhs48_
                d_308_v0_ = rhs49_
                d_482_v114_: C0
                nw79_ = C0()
                nw79_.ctor__((d_308_v0_) - (d_461_cf4_), default__.fm5(d_312_globalState_))
                d_482_v114_ = nw79_
                d_483_v115_: bool
                out3_: bool
                out3_ = default__.m0(d_441_v90_, d_312_globalState_)
                d_483_v115_ = out3_
                d_484_v116_: _dafny.Map
                d_484_v116_ = _dafny.Map({d_482_v114_: d_309_v1_})
                d_483_v115_ = not(((d_484_v116_)[(d_482_v114_ if False else d_482_v114_)] if ((d_482_v114_ if False else d_482_v114_)) in (d_484_v116_) else d_483_v115_))
            d_485_v117_: _dafny.Array
            def lambda19_(d_486_v0_):
                def lambda20_(d_487_i8_):
                    return default__.safeModuloInt(d_487_i8_, d_486_v0_)

                return lambda20_

            init8_ = lambda19_(d_308_v0_)
            nw80_ = _dafny.Array(None, 8)
            for i0_8_ in range(nw80_.length(0)):
                nw80_[i0_8_] = init8_(i0_8_)
            d_485_v117_ = nw80_
            d_488_v118_: D2
            d_488_v118_ = D2_DC5(d_485_v117_)
            d_489_v119_: _dafny.Map
            d_489_v119_ = _dafny.Map({d_485_v117_: (d_488_v118_ if d_441_v90_ else d_488_v118_)})
            d_489_v119_ = (d_489_v119_).set(d_485_v117_, d_488_v118_)
            (d_312_globalState_).f0 = default__.safeModuloInt(d_461_cf4_, (0) - (d_308_v0_))
        elif True:
            d_490___mcc_h21_ = source10_.cf10
            d_491_cf10_ = d_490___mcc_h21_
            d_492_v120_: str
            d_492_v120_ = _dafny.CodePoint('f')
            d_493_v121_: _dafny.MultiSet
            d_493_v121_ = _dafny.MultiSet([d_492_v120_])
            d_494_v122_: _dafny.Array
            nw81_ = _dafny.Array(int(0), 27)
            d_494_v122_ = nw81_
            d_495_v123_: _dafny.Map
            d_495_v123_ = _dafny.Map({d_493_v121_: d_494_v122_})
            d_495_v123_ = (d_495_v123_).set(d_493_v121_, d_494_v122_)
            (d_312_globalState_).f0 = (0) - (d_308_v0_)
            d_309_v1_ = (d_308_v0_) <= (default__.safeModuloInt(d_308_v0_, d_308_v0_))
            d_496_v124_: _dafny.Set
            d_496_v124_ = _dafny.Set({d_308_v0_, d_308_v0_, d_308_v0_})
            (d_312_globalState_).f4 = (0) - (len((d_496_v124_).intersection(d_496_v124_)))
        d_497_v125_: _dafny.Seq
        d_497_v125_ = _dafny.SeqWithoutIsStrInference([d_308_v0_, d_308_v0_])
        d_498_v126_: _dafny.Array
        nw82_ = _dafny.Array(int(0), 1)
        d_498_v126_ = nw82_
        d_499_v127_: D2
        d_499_v127_ = D2_DC6(len(default__.fm3(d_310_v2_, d_312_globalState_)), (d_497_v125_) == (d_497_v125_), d_498_v126_)
        source11_ = d_499_v127_
        if source11_.is_DC6:
            d_500___mcc_h22_ = source11_.cf12
            d_501___mcc_h23_ = source11_.cf13
            d_502___mcc_h24_ = source11_.cf14
            d_503_cf14_ = d_502___mcc_h24_
            d_504_cf13_ = d_501___mcc_h23_
            d_505_cf12_ = d_500___mcc_h22_
            d_308_v0_ = ((761) * (d_505_cf12_)) - (d_505_cf12_)
            d_309_v1_ = d_441_v90_
            d_506_v128_: _dafny.Map
            d_506_v128_ = _dafny.Map({default__.fm4(d_505_cf12_, d_504_cf13_, d_312_globalState_): default__.safeModuloInt(d_308_v0_, len(d_497_v125_))})
            rhs50_ = d_441_v90_
            rhs51_ = False
            rhs52_ = (d_506_v128_) | (_dafny.Map({d_505_cf12_: d_505_cf12_}))
            d_441_v90_ = rhs50_
            d_441_v90_ = rhs51_
            d_506_v128_ = rhs52_
            d_441_v90_ = d_309_v1_
        elif source11_.is_DC7:
            d_507___mcc_h25_ = source11_.cf15
            d_508_cf15_ = d_507___mcc_h25_
            d_509_v129_: bool
            out4_: bool
            out4_ = default__.m0((((d_442_v91_)[d_508_cf15_] if (d_508_cf15_) in (d_442_v91_) else not(d_441_v90_)) if d_309_v1_ else d_309_v1_), d_312_globalState_)
            d_509_v129_ = out4_
            d_510_v130_: C0
            nw83_ = C0()
            nw83_.ctor__(d_308_v0_, d_440_v89_)
            d_510_v130_ = nw83_
            index46_ = default__.safeIndex(921, (d_439_v88_).length(0))
            (d_439_v88_)[index46_] = d_510_v130_
            index47_ = default__.safeIndex(921, (d_439_v88_).length(0))
            rhs53_ = d_510_v130_
            rhs54_ = d_308_v0_
            rhs55_ = d_497_v125_
            rhs56_ = d_510_v130_
            lhs35_ = d_439_v88_
            lhs36_ = default__.safeIndex(921, (d_439_v88_).length(0))
            lhs37_ = d_312_globalState_
            lhs35_[lhs36_] = rhs53_
            lhs37_.f0 = rhs54_
            d_497_v125_ = rhs55_
            d_510_v130_ = rhs56_
            d_511_v131_: str
            d_511_v131_ = _dafny.CodePoint('a')
            d_511_v131_ = d_511_v131_
            d_512_v132_: D2
            d_512_v132_ = D2_DC7(d_508_cf15_)
            d_513_v133_: _dafny.Set
            d_513_v133_ = _dafny.Set({d_508_cf15_})
            d_514_v134_: _dafny.Array
            nw84_ = _dafny.Array(None, 28)
            nw84_[int(0)] = d_512_v132_
            nw84_[int(1)] = d_512_v132_
            nw84_[int(2)] = d_512_v132_
            nw84_[int(3)] = (d_512_v132_ if default__.fm0(d_441_v90_, d_309_v1_, len(d_513_v133_), d_312_globalState_) else D2_DC7(len(d_340_v20_)))
            nw84_[int(4)] = d_512_v132_
            nw84_[int(5)] = d_512_v132_
            nw84_[int(6)] = D2_DC7(68)
            nw84_[int(7)] = d_512_v132_
            nw84_[int(8)] = d_512_v132_
            nw84_[int(9)] = d_512_v132_
            nw84_[int(10)] = D2_DC7(d_308_v0_)
            nw84_[int(11)] = d_512_v132_
            nw84_[int(12)] = d_512_v132_
            nw84_[int(13)] = d_512_v132_
            nw84_[int(14)] = (d_512_v132_ if False else d_512_v132_)
            nw84_[int(15)] = d_512_v132_
            nw84_[int(16)] = d_512_v132_
            nw84_[int(17)] = d_512_v132_
            nw84_[int(18)] = d_512_v132_
            nw84_[int(19)] = D2_DC7(len(d_310_v2_))
            nw84_[int(20)] = d_512_v132_
            nw84_[int(21)] = d_512_v132_
            nw84_[int(22)] = d_512_v132_
            nw84_[int(23)] = d_512_v132_
            nw84_[int(24)] = (d_512_v132_ if d_441_v90_ else d_512_v132_)
            nw84_[int(25)] = d_512_v132_
            nw84_[int(26)] = d_512_v132_
            nw84_[int(27)] = (d_512_v132_ if d_309_v1_ else d_512_v132_)
            d_514_v134_ = nw84_
            index48_ = default__.safeIndex(190, (d_514_v134_).length(0))
            (d_514_v134_)[index48_] = d_512_v132_
            index49_ = default__.safeIndex(190, (d_514_v134_).length(0))
            (d_514_v134_)[index49_] = default__.fm6(d_312_globalState_)
        elif source11_.is_DC5:
            d_515___mcc_h26_ = source11_.cf11
            d_516_cf11_ = d_515___mcc_h26_
            if (d_441_v90_ if d_309_v1_ else d_309_v1_):
                d_441_v90_ = ((0) - (d_308_v0_)) >= ((0) - (default__.fm4(-120, d_309_v1_, d_312_globalState_)))
                d_438_v87_ = (_dafny.MultiSet([d_309_v1_, d_441_v90_, d_309_v1_, d_309_v1_, d_441_v90_])).set(not(not(d_441_v90_)), default__.abs(d_308_v0_))
                (d_312_globalState_).f4 = (d_497_v125_)[default__.safeIndex(d_308_v0_, len(d_497_v125_))]
                d_517_v135_: _dafny.Array
                def lambda21_(d_518_v20_):
                    def lambda22_(d_519_i9_):
                        return d_518_v20_

                    return lambda22_

                init9_ = lambda21_(d_340_v20_)
                nw85_ = _dafny.Array(None, 26)
                for i0_9_ in range(nw85_.length(0)):
                    nw85_[i0_9_] = init9_(i0_9_)
                d_517_v135_ = nw85_
                index50_ = default__.safeIndex(63, (d_517_v135_).length(0))
                (d_517_v135_)[index50_] = (d_340_v20_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkypk")))
                index51_ = default__.safeIndex(63, (d_517_v135_).length(0))
                (d_517_v135_)[index51_] = d_340_v20_
                d_435_v84_ = (d_435_v84_).set(True, d_308_v0_)
            elif True:
                d_520_v136_: bool
                out5_: bool
                out5_ = default__.m0(((0) - (d_308_v0_)) != (len(d_497_v125_)), d_312_globalState_)
                d_520_v136_ = out5_
                d_521_v137_: bool
                out6_: bool
                out6_ = default__.m0(d_309_v1_, d_312_globalState_)
                d_521_v137_ = out6_
                d_520_v136_ = (d_340_v20_) != ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwjyx"))) + (d_340_v20_))
                (d_312_globalState_).f0 = d_308_v0_
                d_522_v138_: C0
                nw86_ = C0()
                nw86_.ctor__(d_308_v0_, d_440_v89_)
                d_522_v138_ = nw86_
            (d_312_globalState_).f4 = -260
            d_523_v139_: _dafny.Array
            nw87_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_523_v139_ = nw87_
            d_523_v139_ = d_523_v139_
            d_524_v140_: _dafny.Map
            d_524_v140_ = _dafny.Map({d_308_v0_: d_308_v0_})
            (d_312_globalState_).f0 = ((d_524_v140_)[(0) - (d_308_v0_)] if ((0) - (d_308_v0_)) in (d_524_v140_) else d_308_v0_)
        elif True:
            d_525___mcc_h27_ = source11_.cf16
            d_526_cf16_ = d_525___mcc_h27_
            d_527_v141_: _dafny.Map
            d_527_v141_ = _dafny.Map({d_441_v90_: d_309_v1_})
            d_528_v142_: C0
            nw88_ = C0()
            nw88_.ctor__(d_308_v0_, d_440_v89_)
            d_528_v142_ = nw88_
            d_529_v143_: _dafny.Map
            d_529_v143_ = _dafny.Map({d_528_v142_: (d_528_v142_).f6})
            if ((0) - ((((d_341_v21_)[len(d_310_v2_)] if (len(d_310_v2_)) in (d_341_v21_) else len(d_527_v141_))) * (d_308_v0_))) > (len((d_529_v143_) | (_dafny.Map({d_528_v142_: (d_528_v142_).f6})))):
                d_530_v144_: _dafny.MultiSet
                d_530_v144_ = _dafny.MultiSet([d_498_v126_])
                d_531_v145_: _dafny.Array
                def lambda23_(d_532_i10_):
                    return True

                init10_ = lambda23_
                nw89_ = _dafny.Array(None, 21)
                for i0_10_ in range(nw89_.length(0)):
                    nw89_[i0_10_] = init10_(i0_10_)
                d_531_v145_ = nw89_
                d_533_v146_: _dafny.Array
                nw90_ = _dafny.Array(_dafny.MultiSet({}), 16)
                d_533_v146_ = nw90_
                d_534_v147_: D0
                d_534_v147_ = D0_DC1(d_530_v144_, d_531_v145_, (d_533_v146_ if d_309_v1_ else d_533_v146_))
                d_534_v147_ = d_534_v147_
                (d_312_globalState_).f4 = (d_528_v142_).f6
                d_435_v84_ = default__.fm7(_dafny.MultiSet([((d_442_v91_)[d_308_v0_] if (d_308_v0_) in (d_442_v91_) else d_441_v90_), False]), d_340_v20_, d_312_globalState_)
                d_309_v1_ = d_309_v1_
                d_435_v84_ = (d_435_v84_).set((d_371_v43_).cf6, d_308_v0_)
            elif True:
                d_535_v148_: D0
                d_535_v148_ = D0_DC0(False)
                d_536_v149_: C0
                nw91_ = C0()
                nw91_.ctor__((d_528_v142_).fm1(_dafny.MultiSet([(d_528_v142_).f6]), d_535_v148_, d_312_globalState_), _dafny.Set({False, d_309_v1_}))
                d_536_v149_ = nw91_
                d_537_v150_: C0
                nw92_ = C0()
                nw92_.ctor__((d_528_v142_).f6, (d_440_v89_).intersection((d_536_v149_).f7))
                d_537_v150_ = nw92_
                d_538_v151_: D3
                d_538_v151_ = D3_DC9(d_497_v125_)
                d_539_v152_: _dafny.MultiSet
                d_539_v152_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_308_v0_]), ((d_538_v151_).cf17).set(default__.safeIndex(len(d_435_v84_), len((d_538_v151_).cf17)), (d_528_v142_).f6), d_497_v125_])
                d_540_v153_: C0
                nw93_ = C0()
                nw93_.ctor__((0) - (((d_539_v152_)[(d_497_v125_).set(default__.safeIndex(122, len(d_497_v125_)), len(d_340_v20_))] if ((d_497_v125_).set(default__.safeIndex(122, len(d_497_v125_)), len(d_340_v20_))) in (d_539_v152_) else 624)), _dafny.Set({d_441_v90_}))
                d_540_v153_ = nw93_
                d_441_v90_ = d_441_v90_
                d_541_v154_: bool
                out7_: bool
                out7_ = default__.m0(d_309_v1_, d_312_globalState_)
                d_541_v154_ = out7_
            d_441_v90_ = not (False) or (not((default__.fm0(False, d_309_v1_, d_308_v0_, d_312_globalState_)) in (d_438_v87_)))
            d_542_v155_: D3
            d_542_v155_ = D3_DC10(d_309_v1_, len(default__.fm9(664, (d_341_v21_).cardinality, d_312_globalState_)), d_308_v0_)
            d_543_v156_: D3
            d_543_v156_ = D3_DC12(d_542_v155_)
            d_442_v91_ = default__.fm8(d_309_v1_, d_543_v156_, d_440_v89_, d_308_v0_, d_312_globalState_)
            if (default__.safeDivisionInt((d_528_v142_).f6, (d_528_v142_).f6)) != (((d_528_v142_).f6) + (len(d_340_v20_))):
                (d_312_globalState_).f0 = (d_528_v142_).f6
                d_544_v157_: bool
                out8_: bool
                out8_ = default__.m0(not((_dafny.MultiSet([d_309_v1_, not(True)])).issubset(d_438_v87_)), d_312_globalState_)
                d_544_v157_ = out8_
                index52_ = default__.safeIndex(771, (d_439_v88_).length(0))
                (d_439_v88_)[index52_] = d_528_v142_
                index53_ = default__.safeIndex(771, (d_439_v88_).length(0))
                rhs57_ = (d_341_v21_) - (d_341_v21_)
                rhs58_ = ((d_308_v0_) * ((d_528_v142_).f6) if not (d_544_v157_) or (d_544_v157_) else len(d_527_v141_))
                rhs59_ = not(d_441_v90_)
                rhs60_ = d_528_v142_
                lhs38_ = d_312_globalState_
                lhs39_ = d_439_v88_
                lhs40_ = default__.safeIndex(771, (d_439_v88_).length(0))
                d_341_v21_ = rhs57_
                lhs38_.f0 = rhs58_
                d_309_v1_ = rhs59_
                lhs39_[lhs40_] = rhs60_
                d_545_v158_: _dafny.Map
                d_545_v158_ = _dafny.Map({(default__.fm10(len(d_442_v91_), (d_528_v142_).f6, d_312_globalState_)) < (_dafny.SeqWithoutIsStrInference([d_441_v90_])): D3_DC12(d_542_v155_)})
                d_545_v158_ = default__.fm11(d_312_globalState_)
                d_546_v159_: _dafny.Array
                nw94_ = _dafny.Array(False, 29)
                d_546_v159_ = nw94_
                index54_ = default__.safeIndex(306, (d_546_v159_).length(0))
                (d_546_v159_)[index54_] = d_441_v90_
                index55_ = default__.safeIndex(306, (d_546_v159_).length(0))
                rhs61_ = d_441_v90_
                rhs62_ = d_340_v20_
                lhs41_ = d_546_v159_
                lhs42_ = default__.safeIndex(306, (d_546_v159_).length(0))
                lhs41_[lhs42_] = rhs61_
                d_340_v20_ = rhs62_
            elif True:
                d_441_v90_ = d_441_v90_
                d_441_v90_ = True
                d_547_v160_: _dafny.Map
                d_547_v160_ = _dafny.Map({d_310_v2_: (d_528_v142_).f6})
                index56_ = default__.safeIndex(592, (d_498_v126_).length(0))
                (d_498_v126_)[index56_] = len(d_547_v160_)
                index57_ = default__.safeIndex(592, (d_498_v126_).length(0))
                (d_498_v126_)[index57_] = (d_528_v142_).f6
                d_548_v161_: _dafny.Map
                d_548_v161_ = _dafny.Map({d_528_v142_: ((d_442_v91_)[(d_498_v126_)[default__.safeIndex(592, (d_498_v126_).length(0))]] if ((d_498_v126_)[default__.safeIndex(592, (d_498_v126_).length(0))]) in (d_442_v91_) else d_309_v1_)})
                d_548_v161_ = (d_548_v161_).set(d_528_v142_, d_441_v90_)
                nw95_ = C0()
                nw95_.ctor__(default__.safeModuloInt((d_528_v142_).f6, d_308_v0_), d_440_v89_)
                d_528_v142_ = nw95_
        d_309_v1_ = (d_308_v0_) < (183)
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_498_v126_).length(0)):
            d_549_i11_: int = guard_loop_1_
            if (True) and (((0) <= (d_549_i11_)) and ((d_549_i11_) < ((d_498_v126_).length(0)))):
                (d_498_v126_)[(d_549_i11_)] = (d_549_i11_) - ((d_308_v0_ if d_441_v90_ else d_308_v0_))
        d_550_v162_: D3
        d_550_v162_ = D3_DC9(d_497_v125_)
        source12_ = (D3_DC12(D3_DC12(d_550_v162_)) if (d_441_v90_ if d_441_v90_ else not(d_441_v90_)) else default__.fm12(d_309_v1_, len(default__.fm13(d_441_v90_, d_441_v90_, d_312_globalState_)), d_441_v90_, 49, d_312_globalState_))
        if source12_.is_DC10:
            d_551___mcc_h28_ = source12_.cf18
            d_552___mcc_h29_ = source12_.cf19
            d_553___mcc_h30_ = source12_.cf20
            d_554_cf20_ = d_553___mcc_h30_
            d_555_cf19_ = d_552___mcc_h29_
            d_556_cf18_ = d_551___mcc_h28_
            d_557_v163_: _dafny.Seq
            d_557_v163_ = _dafny.SeqWithoutIsStrInference([d_435_v84_, d_435_v84_])
            d_558_v164_: bool
            out9_: bool
            out9_ = default__.m0((d_557_v163_) == (d_557_v163_), d_312_globalState_)
            d_558_v164_ = out9_
            d_340_v20_ = (d_340_v20_) + (d_340_v20_)
            d_559_v165_: _dafny.Seq
            d_559_v165_ = _dafny.SeqWithoutIsStrInference([d_310_v2_])
            d_560_v166_: C0
            nw96_ = C0()
            nw96_.ctor__(len(d_435_v84_), d_440_v89_)
            d_560_v166_ = nw96_
            d_561_v167_: _dafny.Seq
            d_561_v167_ = _dafny.SeqWithoutIsStrInference([d_560_v166_, d_560_v166_])
            d_562_v168_: _dafny.Map
            d_562_v168_ = _dafny.Map({(d_561_v167_)[default__.safeIndex(d_555_cf19_, len(d_561_v167_))]: d_441_v90_})
            d_563_v169_: _dafny.Map
            d_563_v169_ = _dafny.Map({_dafny.Map({d_560_v166_: not(d_558_v164_)}): d_309_v1_})
            d_564_v170_: _dafny.Array
            nw97_ = _dafny.Array(None, 17)
            nw97_[int(0)] = d_556_cf18_
            nw97_[int(1)] = d_309_v1_
            nw97_[int(2)] = d_309_v1_
            nw97_[int(3)] = d_556_cf18_
            nw97_[int(4)] = (d_340_v20_) == (d_340_v20_)
            nw97_[int(5)] = not((d_309_v1_) in ((d_559_v165_)[default__.safeIndex(d_555_cf19_, len(d_559_v165_))]))
            nw97_[int(6)] = d_309_v1_
            nw97_[int(7)] = (len(d_497_v125_)) != (len(d_435_v84_))
            nw97_[int(8)] = d_441_v90_
            nw97_[int(9)] = (d_554_cf20_) != (-329)
            nw97_[int(10)] = d_558_v164_
            nw97_[int(11)] = True
            nw97_[int(12)] = (_dafny.Map({d_562_v168_: d_556_cf18_})) != (d_563_v169_)
            nw97_[int(13)] = (d_556_cf18_ if d_309_v1_ else d_441_v90_)
            nw97_[int(14)] = d_558_v164_
            nw97_[int(15)] = d_556_cf18_
            nw97_[int(16)] = d_441_v90_
            d_564_v170_ = nw97_
            index58_ = default__.safeIndex(75, (d_564_v170_).length(0))
            (d_564_v170_)[index58_] = d_556_cf18_
            d_565_v172_: _dafny.Set
            d_565_v172_ = _dafny.Set({d_310_v2_})
            d_566_v173_: _dafny.Set
            d_566_v173_ = _dafny.Set({len(d_565_v172_)})
            index59_ = default__.safeIndex(75, (d_564_v170_).length(0))
            def iife45_():
                coll25_ = _dafny.Map()
                compr_25_: int
                for compr_25_ in _dafny.IntegerRange(908, -712):
                    d_567_v171_: int = compr_25_
                    if ((908) <= (d_567_v171_)) and ((d_567_v171_) < (-712)):
                        coll25_[(d_567_v171_) * (d_308_v0_)] = False
                return _dafny.Map(coll25_)
            (d_564_v170_)[index59_] = (_dafny.Set({(d_497_v125_)[default__.safeIndex((d_341_v21_).cardinality, len(d_497_v125_))], (d_560_v166_).fm2(d_556_cf18_, d_497_v125_, _dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(iife45_()
): True})) for d_568_i12_ in range(default__.abs(611))]), d_435_v84_, d_312_globalState_)})).isdisjoint(d_566_v173_)
            d_569_v174_: str
            d_569_v174_ = _dafny.CodePoint('a')
            d_555_cf19_ = len((d_340_v20_).set(default__.safeIndex(((d_560_v166_).f6) * (d_308_v0_), len(d_340_v20_)), d_569_v174_))
        elif source12_.is_DC11:
            d_570_v175_: D3
            d_570_v175_ = D3_DC9(_dafny.SeqWithoutIsStrInference([d_308_v0_]))
            pat_let_tv22_ = d_497_v125_
            d_571_v176_: _dafny.MultiSet
            def iife46_(_pat_let10_0):
                def iife47_(d_572_dt__update__tmp_h3_):
                    def iife48_(_pat_let11_0):
                        def iife49_(d_573_dt__update_hcf17_h0_):
                            return D3_DC9(d_573_dt__update_hcf17_h0_)
                        return iife49_(_pat_let11_0)
                    return iife48_(pat_let_tv22_)
                return iife47_(_pat_let10_0)
            d_571_v176_ = _dafny.MultiSet([d_570_v175_, d_570_v175_, d_570_v175_, iife46_(d_570_v175_)])
            index60_ = default__.safeIndex(608, (d_498_v126_).length(0))
            (d_498_v126_)[index60_] = (0) - ((d_571_v176_).cardinality)
            d_574_v177_: _dafny.Map
            d_574_v177_ = _dafny.Map({d_309_v1_: True})
            index61_ = default__.safeIndex(608, (d_498_v126_).length(0))
            (d_498_v126_)[index61_] = default__.fm4(d_308_v0_, ((d_574_v177_)[d_309_v1_] if (d_309_v1_) in (d_574_v177_) else d_441_v90_), d_312_globalState_)
            d_575_v178_: _dafny.Set
            d_575_v178_ = _dafny.Set({d_435_v84_, _dafny.Map({(D0_DC0(False)).cf0: len(d_310_v2_)}), d_435_v84_})
            d_576_v179_: _dafny.Set
            d_576_v179_ = _dafny.Set({default__.safeDivisionInt(-660, d_308_v0_), len(d_575_v178_), 253})
            d_576_v179_ = d_576_v179_
            d_577_v180_: C0
            nw98_ = C0()
            nw98_.ctor__(default__.fm4(len(d_576_v179_), d_441_v90_, d_312_globalState_), d_440_v89_)
            d_577_v180_ = nw98_
            index62_ = default__.safeIndex(608, (d_498_v126_).length(0))
            (d_498_v126_)[index62_] = d_308_v0_
        elif source12_.is_DC9:
            d_578___mcc_h31_ = source12_.cf17
            d_579_cf17_ = d_578___mcc_h31_
            d_435_v84_ = ((d_435_v84_) | (d_435_v84_)) | (d_435_v84_)
            if d_309_v1_:
                d_580_v181_: _dafny.Map
                d_580_v181_ = _dafny.Map({(0) - (d_308_v0_): len((d_340_v20_ if d_441_v90_ else d_340_v20_))})
                d_580_v181_ = (d_580_v181_) | (d_580_v181_)
                d_581_v182_: str
                d_581_v182_ = _dafny.CodePoint('g')
                d_582_v183_: _dafny.Map
                d_582_v183_ = _dafny.Map({False: d_581_v182_})
                d_583_v184_: _dafny.Map
                d_583_v184_ = _dafny.Map({d_340_v20_: ((d_582_v183_)[d_309_v1_] if (d_309_v1_) in (d_582_v183_) else d_581_v182_)})
                d_583_v184_ = (d_583_v184_).set(d_340_v20_, d_581_v182_)
                d_308_v0_ = (d_308_v0_) + (d_308_v0_)
                d_584_v185_: D2
                d_584_v185_ = D2_DC5(d_498_v126_)
                pat_let_tv23_ = d_498_v126_
                pat_let_tv24_ = d_498_v126_
                def iife50_(_pat_let12_0):
                    def iife51_(d_585_dt__update__tmp_h4_):
                        def iife52_(_pat_let13_0):
                            def iife53_(d_586_dt__update_hcf11_h1_):
                                return D2_DC5(d_586_dt__update_hcf11_h1_)
                            return iife53_(_pat_let13_0)
                        return iife52_((pat_let_tv23_ if True else pat_let_tv24_))
                    return iife51_(_pat_let12_0)
                d_584_v185_ = iife50_(d_584_v185_)
                d_587_v186_: C0
                nw99_ = C0()
                nw99_.ctor__(d_308_v0_, d_440_v89_)
                d_587_v186_ = nw99_
                index63_ = default__.safeIndex(792, (d_439_v88_).length(0))
                (d_439_v88_)[index63_] = d_587_v186_
                index64_ = default__.safeIndex(792, (d_439_v88_).length(0))
                (d_439_v88_)[index64_] = (d_587_v186_ if d_309_v1_ else d_587_v186_)
            elif True:
                d_588_v187_: C0
                nw100_ = C0()
                nw100_.ctor__(d_308_v0_, _dafny.Set({False, d_309_v1_}))
                d_588_v187_ = nw100_
                d_588_v187_ = d_588_v187_
                d_340_v20_ = d_340_v20_
                index65_ = default__.safeIndex(422, (d_498_v126_).length(0))
                (d_498_v126_)[index65_] = d_308_v0_
                index66_ = default__.safeIndex(422, (d_498_v126_).length(0))
                (d_498_v126_)[index66_] = d_308_v0_
                d_588_v187_ = d_588_v187_
                d_589_v188_: _dafny.Seq
                d_589_v188_ = _dafny.SeqWithoutIsStrInference([d_340_v20_])
                index67_ = default__.safeIndex(422, (d_498_v126_).length(0))
                (d_498_v126_)[index67_] = ((len(d_589_v188_)) * (-818)) * ((d_438_v87_).cardinality)
            d_590_v189_: _dafny.Array
            nw101_ = _dafny.Array(None, 26)
            nw101_[int(0)] = (d_310_v2_)[default__.safeIndex(d_308_v0_, len(d_310_v2_))]
            nw101_[int(1)] = d_309_v1_
            nw101_[int(2)] = d_309_v1_
            nw101_[int(3)] = d_441_v90_
            nw101_[int(4)] = False
            nw101_[int(5)] = True
            nw101_[int(6)] = not(False)
            nw101_[int(7)] = d_309_v1_
            nw101_[int(8)] = d_309_v1_
            nw101_[int(9)] = d_441_v90_
            nw101_[int(10)] = d_309_v1_
            nw101_[int(11)] = not(default__.fm0(d_309_v1_, d_309_v1_, default__.fm4(len(d_310_v2_), d_309_v1_, d_312_globalState_), d_312_globalState_))
            nw101_[int(12)] = d_309_v1_
            nw101_[int(13)] = d_309_v1_
            nw101_[int(14)] = d_309_v1_
            nw101_[int(15)] = False
            nw101_[int(16)] = d_441_v90_
            nw101_[int(17)] = d_309_v1_
            nw101_[int(18)] = d_441_v90_
            nw101_[int(19)] = True
            nw101_[int(20)] = d_441_v90_
            nw101_[int(21)] = default__.fm0(d_309_v1_, (d_371_v43_).cf5, d_308_v0_, d_312_globalState_)
            nw101_[int(22)] = not(d_309_v1_)
            nw101_[int(23)] = d_309_v1_
            nw101_[int(24)] = d_309_v1_
            nw101_[int(25)] = d_309_v1_
            d_590_v189_ = nw101_
            d_591_v190_: _dafny.Map
            d_591_v190_ = _dafny.Map({d_441_v90_: d_590_v189_})
            d_592_v191_: _dafny.Array
            nw102_ = _dafny.Array(None, 26)
            nw102_[int(0)] = ((d_591_v190_)[d_309_v1_] if (d_309_v1_) in (d_591_v190_) else d_590_v189_)
            nw102_[int(1)] = d_590_v189_
            nw102_[int(2)] = d_590_v189_
            nw102_[int(3)] = d_590_v189_
            nw102_[int(4)] = d_590_v189_
            nw102_[int(5)] = d_590_v189_
            nw102_[int(6)] = (d_590_v189_ if d_441_v90_ else d_590_v189_)
            nw102_[int(7)] = d_590_v189_
            nw102_[int(8)] = d_590_v189_
            nw102_[int(9)] = d_590_v189_
            nw102_[int(10)] = d_590_v189_
            nw102_[int(11)] = d_590_v189_
            nw102_[int(12)] = d_590_v189_
            nw102_[int(13)] = d_590_v189_
            nw102_[int(14)] = d_590_v189_
            nw102_[int(15)] = d_590_v189_
            nw102_[int(16)] = d_590_v189_
            nw102_[int(17)] = d_590_v189_
            nw102_[int(18)] = d_590_v189_
            nw102_[int(19)] = (d_590_v189_ if d_309_v1_ else d_590_v189_)
            nw102_[int(20)] = d_590_v189_
            nw102_[int(21)] = d_590_v189_
            nw102_[int(22)] = d_590_v189_
            nw102_[int(23)] = (d_590_v189_ if d_309_v1_ else d_590_v189_)
            nw102_[int(24)] = d_590_v189_
            nw102_[int(25)] = d_590_v189_
            d_592_v191_ = nw102_
            index68_ = default__.safeIndex(215, (d_592_v191_).length(0))
            (d_592_v191_)[index68_] = d_590_v189_
            index69_ = default__.safeIndex(215, (d_592_v191_).length(0))
            (d_592_v191_)[index69_] = d_590_v189_
            source13_ = d_444_v93_
            if source13_.is_DC3:
                d_593___mcc_h33_ = source13_.cf5
                d_594___mcc_h34_ = source13_.cf6
                d_595___mcc_h35_ = source13_.cf7
                d_596___mcc_h36_ = source13_.cf8
                d_597___mcc_h37_ = source13_.cf9
                d_598_cf9_ = d_597___mcc_h37_
                d_599_cf8_ = d_596___mcc_h36_
                d_600_cf7_ = d_595___mcc_h35_
                d_601_cf6_ = d_594___mcc_h34_
                d_602_cf5_ = d_593___mcc_h33_
                d_603_v192_: str
                d_603_v192_ = _dafny.CodePoint('v')
                d_604_v193_: _dafny.Map
                d_604_v193_ = _dafny.Map({d_599_cf8_: False})
                d_605_v194_: _dafny.Array
                nw103_ = _dafny.Array(None, 2)
                nw103_[int(0)] = _dafny.Set({d_599_cf8_, ((d_604_v193_)[d_441_v90_] if (d_441_v90_) in (d_604_v193_) else d_601_cf6_)})
                nw103_[int(1)] = _dafny.Set({d_602_cf5_, (d_310_v2_)[default__.safeIndex(len(d_497_v125_), len(d_310_v2_))], not(d_441_v90_)})
                d_605_v194_ = nw103_
                index70_ = default__.safeIndex(316, (d_605_v194_).length(0))
                (d_605_v194_)[index70_] = d_440_v89_
                d_606_v195_: _dafny.Map
                d_606_v195_ = _dafny.Map({d_308_v0_: D3_DC11()})
                index71_ = default__.safeIndex(316, (d_605_v194_).length(0))
                rhs63_ = d_598_cf9_
                rhs64_ = (_dafny.MultiSet((d_579_cf17_) + (d_497_v125_))).ispropersubset(_dafny.MultiSet([d_308_v0_, d_308_v0_, d_308_v0_]))
                rhs65_ = d_603_v192_
                rhs66_ = (d_440_v89_ if not((len(d_606_v195_)) >= (len(d_443_v92_))) else d_440_v89_)
                lhs43_ = d_312_globalState_
                lhs44_ = d_605_v194_
                lhs45_ = default__.safeIndex(316, (d_605_v194_).length(0))
                lhs43_.f0 = rhs63_
                d_599_cf8_ = rhs64_
                d_603_v192_ = rhs65_
                lhs44_[lhs45_] = rhs66_
                index72_ = default__.safeIndex(403, (d_590_v189_).length(0))
                (d_590_v189_)[index72_] = True
                index73_ = default__.safeIndex(403, (d_590_v189_).length(0))
                (d_590_v189_)[index73_] = default__.fm0(d_309_v1_, d_441_v90_, d_598_cf9_, d_312_globalState_)
                d_607_v196_: C0
                nw104_ = C0()
                nw104_.ctor__(default__.safeDivisionInt(d_598_cf9_, (0) - (len(d_591_v190_))), default__.fm5(d_312_globalState_))
                d_607_v196_ = nw104_
                d_601_cf6_ = True
            elif source13_.is_DC2:
                d_608___mcc_h38_ = source13_.cf4
                d_609_cf4_ = d_608___mcc_h38_
                d_610_v197_: C0
                nw105_ = C0()
                nw105_.ctor__(default__.safeModuloInt(default__.fm4(d_609_cf4_, d_309_v1_, d_312_globalState_), (0) - (d_308_v0_)), (_dafny.Set({d_441_v90_}) if d_441_v90_ else d_440_v89_))
                d_610_v197_ = nw105_
                d_611_v198_: _dafny.Map
                d_611_v198_ = _dafny.Map({d_441_v90_: (d_310_v2_)[default__.safeIndex(212, len(d_310_v2_))]})
                d_610_v197_ = (d_610_v197_ if ((d_611_v198_)[d_441_v90_] if (d_441_v90_) in (d_611_v198_) else d_441_v90_) else d_610_v197_)
                d_612_v199_: D4
                d_612_v199_ = D4_DC13(d_610_v197_)
                d_610_v197_ = (d_612_v199_).cf22
                (d_312_globalState_).f4 = ((d_610_v197_).f6) + (d_609_cf4_)
                (d_312_globalState_).f4 = len(((d_435_v84_) | (d_435_v84_)) | ((d_435_v84_) | (d_435_v84_)))
            elif True:
                d_613___mcc_h39_ = source13_.cf10
                d_614_cf10_ = d_613___mcc_h39_
                d_615_v200_: _dafny.Seq
                d_615_v200_ = _dafny.SeqWithoutIsStrInference([d_498_v126_])
                d_616_v201_: _dafny.Map
                d_616_v201_ = _dafny.Map({default__.fm0(d_441_v90_, d_309_v1_, d_308_v0_, d_312_globalState_): d_615_v200_})
                d_616_v201_ = (d_616_v201_).set(d_441_v90_, (d_615_v200_ if d_309_v1_ else d_615_v200_))
                (d_312_globalState_).f4 = (len((d_435_v84_) | ((d_435_v84_).set(d_309_v1_, d_308_v0_)))) * (default__.fm4((0) - (len(d_340_v20_)), default__.fm0(d_309_v1_, d_309_v1_, d_308_v0_, d_312_globalState_), d_312_globalState_))
                d_617_v202_: C0
                nw106_ = C0()
                nw106_.ctor__(382, d_440_v89_)
                d_617_v202_ = nw106_
                d_618_v203_: _dafny.Set
                d_618_v203_ = _dafny.Set({d_617_v202_, d_617_v202_})
                d_619_v204_: str
                d_619_v204_ = _dafny.CodePoint('r')
                d_620_v205_: D4
                d_620_v205_ = D4_DC16(d_618_v203_, d_619_v204_, (d_617_v202_).f6, d_619_v204_)
                d_621_v206_: _dafny.Map
                d_621_v206_ = _dafny.Map({d_308_v0_: (0) - (d_308_v0_)})
                rhs67_ = (default__.safeModuloInt(len(_dafny.Map({(d_617_v202_).f6: d_617_v202_})), (d_617_v202_).f6)) != ((0) - (d_308_v0_))
                rhs68_ = d_620_v205_
                rhs69_ = d_621_v206_
                rhs70_ = d_498_v126_
                d_441_v90_ = rhs67_
                d_620_v205_ = rhs68_
                d_621_v206_ = rhs69_
                d_498_v126_ = rhs70_
        elif True:
            d_622___mcc_h32_ = source12_.cf21
            d_623_cf21_ = d_622___mcc_h32_
            d_624_v207_: bool
            out10_: bool
            out10_ = default__.m0(not(d_441_v90_), d_312_globalState_)
            d_624_v207_ = out10_
            (d_312_globalState_).f0 = d_308_v0_
            d_625_v208_: D3
            d_625_v208_ = D3_DC10(not(d_441_v90_), d_308_v0_, -144)
            pat_let_tv25_ = d_308_v0_
            d_626_v209_: _dafny.Array
            nw107_ = _dafny.Array(None, 23)
            nw107_[int(0)] = d_441_v90_
            nw107_[int(1)] = d_441_v90_
            nw107_[int(2)] = d_441_v90_
            nw107_[int(3)] = not(d_309_v1_)
            nw107_[int(4)] = d_441_v90_
            nw107_[int(5)] = False
            nw107_[int(6)] = (d_624_v207_) and (not(d_624_v207_))
            nw107_[int(7)] = d_309_v1_
            nw107_[int(8)] = d_309_v1_
            nw107_[int(9)] = not(True)
            nw107_[int(10)] = d_441_v90_
            nw107_[int(11)] = d_309_v1_
            nw107_[int(12)] = d_624_v207_
            nw107_[int(13)] = ((_dafny.MultiSet([d_309_v1_])).cardinality) < (-508)
            nw107_[int(14)] = not (d_441_v90_) or (False)
            nw107_[int(15)] = d_441_v90_
            nw107_[int(16)] = d_309_v1_
            nw107_[int(17)] = d_624_v207_
            def iife54_(_pat_let14_0):
                def iife55_(d_627_dt__update__tmp_h5_):
                    def iife56_(_pat_let15_0):
                        def iife57_(d_628_dt__update_hcf20_h0_):
                            return D3_DC10((d_627_dt__update__tmp_h5_).cf18, (d_627_dt__update__tmp_h5_).cf19, d_628_dt__update_hcf20_h0_)
                        return iife57_(_pat_let15_0)
                    return iife56_(pat_let_tv25_)
                return iife55_(_pat_let14_0)
            nw107_[int(18)] = (iife54_(d_625_v208_)).cf18
            nw107_[int(19)] = d_441_v90_
            nw107_[int(20)] = d_441_v90_
            nw107_[int(21)] = d_624_v207_
            nw107_[int(22)] = (d_310_v2_)[default__.safeIndex(-229, len(d_310_v2_))]
            d_626_v209_ = nw107_
            index74_ = default__.safeIndex(672, (d_626_v209_).length(0))
            (d_626_v209_)[index74_] = d_441_v90_
            index75_ = default__.safeIndex(672, (d_626_v209_).length(0))
            (d_626_v209_)[index75_] = d_309_v1_
            d_629_v210_: _dafny.Map
            d_629_v210_ = _dafny.Map({len(d_442_v91_): d_308_v0_})
            d_630_v212_: C0
            nw108_ = C0()
            def iife58_():
                coll26_ = _dafny.Set()
                compr_26_: _dafny.Map
                for compr_26_ in (_dafny.Map({d_629_v210_: d_310_v2_})).keys.Elements:
                    d_631_v211_: _dafny.Map = compr_26_
                    if (d_631_v211_) in (_dafny.Map({d_629_v210_: d_310_v2_})):
                        coll26_ = coll26_.union(_dafny.Set([d_631_v211_]))
                return _dafny.Set(coll26_)
            nw108_.ctor__(len(iife58_()
            ), _dafny.Set({d_309_v1_}))
            d_630_v212_ = nw108_
        if ((d_438_v87_) | (d_438_v87_)) != (d_438_v87_):
            d_632_v213_: str
            d_632_v213_ = _dafny.CodePoint('m')
            d_632_v213_ = d_632_v213_
            d_442_v91_ = (d_442_v91_) | (d_442_v91_)
            d_340_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qyqejkm"))
            d_633_v214_: D2
            d_633_v214_ = D2_DC7(default__.fm4((0) - (d_308_v0_), d_309_v1_, d_312_globalState_))
            d_634_v215_: D2
            d_634_v215_ = D2_DC8(d_633_v214_)
            source14_ = d_634_v215_
            if source14_.is_DC6:
                d_635___mcc_h40_ = source14_.cf12
                d_636___mcc_h41_ = source14_.cf13
                d_637___mcc_h42_ = source14_.cf14
                d_638_cf14_ = d_637___mcc_h42_
                d_639_cf13_ = d_636___mcc_h41_
                d_640_cf12_ = d_635___mcc_h40_
                d_641_v216_: _dafny.Array
                nw109_ = _dafny.Array(None, 6)
                nw109_[int(0)] = (d_308_v0_) == ((d_497_v125_)[default__.safeIndex(d_640_cf12_, len(d_497_v125_))])
                nw109_[int(1)] = True
                nw109_[int(2)] = not(((d_442_v91_)[d_640_cf12_] if (d_640_cf12_) in (d_442_v91_) else d_639_cf13_))
                nw109_[int(3)] = d_441_v90_
                nw109_[int(4)] = d_309_v1_
                nw109_[int(5)] = d_309_v1_
                d_641_v216_ = nw109_
                index76_ = default__.safeIndex(855, (d_641_v216_).length(0))
                (d_641_v216_)[index76_] = d_309_v1_
                index77_ = default__.safeIndex(855, (d_641_v216_).length(0))
                (d_641_v216_)[index77_] = d_309_v1_
                d_642_v217_: _dafny.Array
                def lambda24_(d_643_v20_):
                    def lambda25_(d_644_i13_):
                        return (d_643_v20_) + (d_643_v20_)

                    return lambda25_

                init11_ = lambda24_(d_340_v20_)
                nw110_ = _dafny.Array(None, 10)
                for i0_11_ in range(nw110_.length(0)):
                    nw110_[i0_11_] = init11_(i0_11_)
                d_642_v217_ = nw110_
                index78_ = default__.safeIndex(2, (d_642_v217_).length(0))
                (d_642_v217_)[index78_] = d_340_v20_
                index79_ = default__.safeIndex(2, (d_642_v217_).length(0))
                (d_642_v217_)[index79_] = (d_340_v20_) + (d_340_v20_)
                d_645_v220_: _dafny.Array
                def lambda26_(d_646_cf12_):
                    def lambda27_(d_647_i14_):
                        def iife59_():
                            coll27_ = _dafny.Set()
                            compr_27_: int
                            for compr_27_ in _dafny.IntegerRange(873, 129):
                                d_648_v218_: int = compr_27_
                                if ((873) <= (d_648_v218_)) and ((d_648_v218_) < (129)):
                                    coll27_ = coll27_.union(_dafny.Set([(d_648_v218_) - (d_646_cf12_)]))
                            return _dafny.Set(coll27_)
                        def iife60_():
                            coll28_ = _dafny.Set()
                            compr_28_: int
                            for compr_28_ in _dafny.IntegerRange(15, 772):
                                d_649_v219_: int = compr_28_
                                if ((15) <= (d_649_v219_)) and ((d_649_v219_) < (772)):
                                    coll28_ = coll28_.union(_dafny.Set([(d_649_v219_) + (d_646_cf12_)]))
                            return _dafny.Set(coll28_)
                        return (iife59_()
                        ) - (iife60_()
                        )

                    return lambda27_

                init12_ = lambda26_(d_640_cf12_)
                nw111_ = _dafny.Array(None, 26)
                for i0_12_ in range(nw111_.length(0)):
                    nw111_[i0_12_] = init12_(i0_12_)
                d_645_v220_ = nw111_
                d_650_v221_: _dafny.Set
                d_650_v221_ = _dafny.Set({len(_dafny.Set({d_640_cf12_, len(d_340_v20_)})), d_308_v0_})
                d_651_v222_: _dafny.Seq
                d_651_v222_ = _dafny.SeqWithoutIsStrInference([d_440_v89_, d_440_v89_])
                d_652_v223_: C0
                nw112_ = C0()
                nw112_.ctor__(d_640_cf12_, (d_651_v222_)[default__.safeIndex(len((d_642_v217_)[default__.safeIndex(2, (d_642_v217_).length(0))]), len(d_651_v222_))])
                d_652_v223_ = nw112_
                d_653_v224_: _dafny.Set
                d_653_v224_ = _dafny.Set({d_652_v223_})
                d_654_v225_: D4
                d_654_v225_ = D4_DC16(d_653_v224_, d_632_v213_, d_308_v0_, d_632_v213_)
                index80_ = default__.safeIndex(332, (d_645_v220_).length(0))
                (d_645_v220_)[index80_] = (d_650_v221_ if True else _dafny.Set({len(d_650_v221_), d_640_cf12_, d_640_cf12_, (d_654_v225_).cf27, d_640_cf12_}))
                index81_ = default__.safeIndex(332, (d_645_v220_).length(0))
                (d_645_v220_)[index81_] = d_650_v221_
                d_655_v226_: D3
                d_655_v226_ = D3_DC12(d_550_v162_)
                d_656_v227_: _dafny.Map
                d_656_v227_ = _dafny.Map({d_655_v226_: d_652_v223_})
                d_656_v227_ = (d_656_v227_).set(D3_DC12(d_550_v162_), d_652_v223_)
            elif source14_.is_DC7:
                d_657___mcc_h43_ = source14_.cf15
                d_658_cf15_ = d_657___mcc_h43_
                d_659_v228_: C0
                nw113_ = C0()
                nw113_.ctor__(len(d_497_v125_), d_440_v89_)
                d_659_v228_ = nw113_
                d_660_v229_: D4
                d_660_v229_ = D4_DC13(d_659_v228_)
                d_661_v230_: _dafny.Map
                d_661_v230_ = _dafny.Map({d_632_v213_: default__.fm0(d_441_v90_, d_309_v1_, d_308_v0_, d_312_globalState_)})
                d_662_v231_: D1
                d_662_v231_ = D1_DC3(d_309_v1_, default__.fm0(default__.fm0(d_441_v90_, d_309_v1_, 844, d_312_globalState_), ((d_661_v230_)[d_632_v213_] if (d_632_v213_) in (d_661_v230_) else True), len(d_310_v2_), d_312_globalState_), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tim"))).set(default__.safeIndex((d_438_v87_).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tim")))), _dafny.CodePoint('u')), d_309_v1_, d_308_v0_)
                d_663_v232_: D1
                d_663_v232_ = D1_DC4(d_662_v231_)
                rhs71_ = D4_DC13((d_659_v228_ if d_441_v90_ else d_659_v228_))
                rhs72_ = D1_DC4(default__.fm14(d_312_globalState_))
                rhs73_ = d_658_cf15_
                lhs46_ = d_312_globalState_
                d_660_v229_ = rhs71_
                d_663_v232_ = rhs72_
                lhs46_.f0 = rhs73_
                d_664_v233_: bool
                out11_: bool
                out11_ = default__.m0(not(d_309_v1_), d_312_globalState_)
                d_664_v233_ = out11_
                d_664_v233_ = not((d_664_v233_ if default__.fm0(not(d_441_v90_), d_309_v1_, len(d_497_v125_), d_312_globalState_) else d_309_v1_))
                d_308_v0_ = 230
            elif source14_.is_DC5:
                d_665___mcc_h44_ = source14_.cf11
                d_666_cf11_ = d_665___mcc_h44_
                (d_312_globalState_).f0 = (0) - (default__.safeDivisionInt(d_308_v0_, (d_308_v0_ if (d_310_v2_)[default__.safeIndex(default__.fm4(d_308_v0_, d_309_v1_, d_312_globalState_), len(d_310_v2_))] else d_308_v0_)))
                d_667_v234_: _dafny.Array
                nw114_ = _dafny.Array(False, 20)
                d_667_v234_ = nw114_
                index82_ = default__.safeIndex(321, (d_667_v234_).length(0))
                (d_667_v234_)[index82_] = d_309_v1_
                d_668_v235_: D3
                d_668_v235_ = D3_DC12(D3_DC11())
                index83_ = default__.safeIndex(321, (d_667_v234_).length(0))
                rhs74_ = d_441_v90_
                rhs75_ = d_668_v235_
                rhs76_ = d_667_v234_
                rhs77_ = (default__.fm4(d_308_v0_, d_309_v1_, d_312_globalState_)) <= (((D2_DC6(d_308_v0_, d_441_v90_, d_498_v126_)).cf12) * (d_308_v0_))
                rhs78_ = default__.fm0(d_441_v90_, d_309_v1_, d_308_v0_, d_312_globalState_)
                lhs47_ = d_667_v234_
                lhs48_ = default__.safeIndex(321, (d_667_v234_).length(0))
                lhs47_[lhs48_] = rhs74_
                d_668_v235_ = rhs75_
                d_667_v234_ = rhs76_
                d_441_v90_ = rhs77_
                d_441_v90_ = rhs78_
                d_669_v236_: D3
                d_669_v236_ = D3_DC10(d_441_v90_, d_308_v0_, d_308_v0_)
                (d_312_globalState_).f0 = (722) + (((0) - ((d_438_v87_).cardinality)) * ((d_669_v236_).cf19))
                (d_312_globalState_).f0 = ((d_308_v0_) * (d_308_v0_)) - (d_308_v0_)
            elif True:
                d_670___mcc_h45_ = source14_.cf16
                d_671_cf16_ = d_670___mcc_h45_
                (d_312_globalState_).f0 = (d_308_v0_) * (default__.safeDivisionInt(d_308_v0_, (0) - (((_dafny.MultiSet(d_310_v2_)).set(d_309_v1_, default__.abs(d_308_v0_))).cardinality)))
                d_441_v90_ = not(True)
                d_672_v237_: C0
                nw115_ = C0()
                nw115_.ctor__(d_308_v0_, d_440_v89_)
                d_672_v237_ = nw115_
                d_673_v238_: D3
                d_673_v238_ = D3_DC10(d_441_v90_, 354, (d_672_v237_).f6)
                d_309_v1_ = (d_673_v238_).cf18
            pat_let_tv26_ = d_441_v90_
            d_674_v239_: _dafny.Seq
            def iife61_(_pat_let16_0):
                def iife62_(d_675_dt__update__tmp_h6_):
                    def iife63_(_pat_let17_0):
                        def iife64_(d_676_dt__update_hcf23_h0_):
                            return D4_DC15(d_676_dt__update_hcf23_h0_, (d_675_dt__update__tmp_h6_).cf24)
                        return iife64_(_pat_let17_0)
                    return iife63_(pat_let_tv26_)
                return iife62_(_pat_let16_0)
            d_674_v239_ = _dafny.SeqWithoutIsStrInference([iife61_(D4_DC15(d_309_v1_, d_308_v0_)), D4_DC15(not(d_309_v1_), 412)])
            (d_312_globalState_).f0 = (default__.safeModuloInt(len(d_340_v20_), d_308_v0_)) * (len((d_674_v239_) + (d_674_v239_)))
        elif True:
            d_677_v240_: C0
            nw116_ = C0()
            nw116_.ctor__(len(d_340_v20_), d_440_v89_)
            d_677_v240_ = nw116_
            (d_312_globalState_).f0 = (d_677_v240_).f6
            d_678_v241_: _dafny.Map
            d_678_v241_ = _dafny.Map({(d_340_v20_) + (d_340_v20_): (d_310_v2_)[default__.safeIndex((d_677_v240_).f6, len(d_310_v2_))]})
            d_678_v241_ = (d_678_v241_).set((d_340_v20_) + (d_340_v20_), d_441_v90_)
            d_679_v242_: C0
            nw117_ = C0()
            nw117_.ctor__(len((_dafny.SeqWithoutIsStrInference([len(d_443_v92_)])) + (d_497_v125_)), (d_677_v240_).f7)
            d_679_v242_ = nw117_
            d_680_v243_: bool
            out12_: bool
            out12_ = default__.m0(True, d_312_globalState_)
            d_680_v243_ = out12_
        d_681_v244_: _dafny.Array
        nw118_ = _dafny.Array(_dafny.Seq({}), 18)
        d_681_v244_ = nw118_
        index84_ = default__.safeIndex(834, (d_681_v244_).length(0))
        (d_681_v244_)[index84_] = _dafny.SeqWithoutIsStrInference([d_498_v126_, d_498_v126_])
        d_682_v245_: _dafny.Seq
        d_682_v245_ = _dafny.SeqWithoutIsStrInference([d_498_v126_, d_498_v126_])
        index85_ = default__.safeIndex(834, (d_681_v244_).length(0))
        (d_681_v244_)[index85_] = d_682_v245_
        _dafny.print(_dafny.string_of(d_308_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_309_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_310_v2_) == (_dafny.SeqWithoutIsStrInference([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v3_) == (_dafny.Map({-45: _dafny.SeqWithoutIsStrInference([True, True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_312_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_312_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_312_globalState_).f5) == (_dafny.Map({-45: _dafny.SeqWithoutIsStrInference([True, True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_313_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_340_v20_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_341_v21_) == (_dafny.MultiSet([6, 6]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_371_v43_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_371_v43_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_371_v43_).cf7).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_371_v43_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_371_v43_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_435_v84_) == (_dafny.Map({True: 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_436_v85_) == (_dafny.Map({492: _dafny.Map({True: 2})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_437_v86_) == (_dafny.Map({1: _dafny.Map({True: 2})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_438_v87_) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_440_v89_) == (_dafny.Set({True, False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_441_v90_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_442_v91_) == (_dafny.Map({6: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_443_v92_) == (_dafny.Map({_dafny.CodePoint('v'): 6}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_444_v93_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_445_v94_)[0]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_445_v94_)[1]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_445_v94_)[2]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_445_v94_)[3]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_445_v94_)[4]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_445_v94_)[5]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_445_v94_)[6]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_445_v94_)[7]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_497_v125_) == (_dafny.SeqWithoutIsStrInference([0, 0]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_498_v126_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_499_v127_).cf12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_499_v127_).cf13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_499_v127_).cf14)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_550_v162_).cf17) == (_dafny.SeqWithoutIsStrInference([0, 0]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_681_v244_)[6])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_682_v245_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC3(False, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0))
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

class D1_DC3(D1, NamedTuple('DC3', [('cf5', Any), ('cf6', Any), ('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {self.cf7.VerbatimString(True)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf4 == __o.cf4
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
        return lambda: D2_DC6(int(0), False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC6(D2, NamedTuple('DC6', [('cf12', Any), ('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)

class D3_DC10(D3, NamedTuple('DC10', [('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11)
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC14()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D4_DC16)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC14(D4, NamedTuple('DC14', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC16(D4, NamedTuple('DC16', [('cf25', Any), ('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC16({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC16) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)

class D5_DC17(D5, NamedTuple('DC17', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC19(int(0), False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D6_DC21)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)

class D6_DC19(D6, NamedTuple('DC19', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC20(D6, NamedTuple('DC20', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC21(D6, NamedTuple('DC21', [])):
    def __dafnystr__(self) -> str:
        return f'D6.DC21'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC21)
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f4: int = int(0)
        self._f1: int = int(0)
        self._f2: bool = False
        self._f3: int = int(0)
        self._f5: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f5 = f5

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
    def f5(self):
        return self._f5

class C0:
    def  __init__(self):
        self._f6: int = int(0)
        self._f7: _dafny.Set = _dafny.Set({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f6, f7):
        (self)._f6 = f6
        (self)._f7 = f7

    def fm1(self, p0, p1, globalState):
        return default__.safeDivisionInt((self).f6, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_683_i0_ in range(default__.abs(287))])))

    def fm2(self, p0, p1, p2, p3, globalState):
        return (D1_DC2(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))))).cf4

    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
