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
        return default__.safeDivisionInt(len(_dafny.Map({-120: False})), (-352) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_0_i0_ in range(default__.abs(873))]))))

    @staticmethod
    def fm1(p0, p1, globalState):
        return (_dafny.CodePoint('k')) in ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewqpfrw")) if False else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfb"))))

    @staticmethod
    def fm5(p0, p1, globalState):
        return D1_DC2(not((len(_dafny.Map({False: False}))) != (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1_i0_ in range(default__.abs(175))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ol"))])))))

    @staticmethod
    def fm6(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pgy")))

    @staticmethod
    def fm7(p0, p1, globalState):
        return _dafny.CodePoint('k')

    @staticmethod
    def fm8(p0, p1, globalState):
        return _dafny.Set({560, default__.safeDivisionInt((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mng")))), 987), (len(_dafny.Map({False: False})) if True else -276), 203})

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([(True) in (_dafny.Set({not(False)})), not (True) or (True), True])

    @staticmethod
    def fm10(globalState):
        if True:
            return (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([False, True])})) | (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([True, True, True, False, False])}))
        elif True:
            return _dafny.Map({not(True): _dafny.SeqWithoutIsStrInference([False])})

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqfptov"))) for d_2_i0_ in range(default__.abs(636))]) if False else _dafny.SeqWithoutIsStrInference([-151]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({890: False})), 505])) + (_dafny.SeqWithoutIsStrInference([768])))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        source0_ = D5_DC13(879, len(_dafny.Map({True: True})), not(False))
        if source0_.is_DC13:
            d_3___mcc_h0_ = source0_.cf19
            d_4___mcc_h1_ = source0_.cf20
            d_5___mcc_h2_ = source0_.cf21
            d_6_cf21_ = d_5___mcc_h2_
            d_7_cf20_ = d_4___mcc_h1_
            d_8_cf19_ = d_3___mcc_h0_
            return D5_DC12(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "msdl")))
        elif True:
            d_9___mcc_h3_ = source0_.cf18
            d_10_cf18_ = d_9___mcc_h3_
            if True:
                return D5_DC12(d_10_cf18_)
            elif True:
                return D5_DC12(d_10_cf18_)

    @staticmethod
    def fm13(globalState):
        source1_ = D5_DC12(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_11_i0_ in range(default__.abs(832))]))
        if source1_.is_DC13:
            d_12___mcc_h0_ = source1_.cf19
            d_13___mcc_h1_ = source1_.cf20
            d_14___mcc_h2_ = source1_.cf21
            d_15_cf21_ = d_14___mcc_h2_
            d_16_cf20_ = d_13___mcc_h1_
            d_17_cf19_ = d_12___mcc_h0_
            def iife0_():
                coll0_ = _dafny.Map()
                compr_0_: _dafny.Map
                for compr_0_ in (_dafny.Map({_dafny.Map({len(_dafny.Set({len(_dafny.Set({d_15_cf21_}))})): d_15_cf21_}): d_15_cf21_})).keys.Elements:
                    d_18_v0_: _dafny.Map = compr_0_
                    if (d_18_v0_) in (_dafny.Map({_dafny.Map({len(_dafny.Set({len(_dafny.Set({d_15_cf21_}))})): d_15_cf21_}): d_15_cf21_})):
                        coll0_[d_18_v0_] = d_16_cf20_
                return _dafny.Map(coll0_)
            return (_dafny.Map({_dafny.Map({d_16_cf20_: d_15_cf21_}): d_17_cf19_})) | (iife0_()
            )
        elif True:
            d_19___mcc_h3_ = source1_.cf18
            d_20_cf18_ = d_19___mcc_h3_
            def iife1_():
                def iife3_():
                    coll3_ = _dafny.Map()
                    compr_3_: int
                    for compr_3_ in _dafny.IntegerRange(857, -21):
                        d_21_v2_: int = compr_3_
                        if ((857) <= (d_21_v2_)) and ((d_21_v2_) < (-21)):
                            coll3_[(d_21_v2_) * (290)] = False
                    return _dafny.Map(coll3_)
                coll1_ = _dafny.Map()
                def iife2_():
                    coll2_ = _dafny.Map()
                    compr_2_: int
                    for compr_2_ in _dafny.IntegerRange(857, -21):
                        d_21_v2_: int = compr_2_
                        if ((857) <= (d_21_v2_)) and ((d_21_v2_) < (-21)):
                            coll2_[(d_21_v2_) * (290)] = False
                    return _dafny.Map(coll2_)
                compr_1_: _dafny.Map
                for compr_1_ in (_dafny.Map({iife2_()
                : True})).keys.Elements:
                    d_22_v1_: _dafny.Map = compr_1_
                    if (d_22_v1_) in (_dafny.Map({iife3_()
                    : True})):
                        coll1_[d_22_v1_] = 542
                return _dafny.Map(coll1_)
            return (iife1_()
            ) | (_dafny.Map({_dafny.Map({len(_dafny.Map({_dafny.CodePoint('x'): 327})): not(True)}): 610}))

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        return _dafny.Set({True, False, (len(_dafny.SeqWithoutIsStrInference([322, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))), len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_23_i0_ in range(default__.abs(558))]))}))]))) > (len(_dafny.Set({76, 112})))})

    @staticmethod
    def fm15(p0, p1, p2, p3, globalState):
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True})), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([508, len(_dafny.Map({False: _dafny.CodePoint('o')}))]))]))).cardinality]): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True, True]))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([False]))) for d_24_i0_ in range(default__.abs(644))]): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))}))

    @staticmethod
    def fm16(p0, p1, p2, p3, globalState):
        return (_dafny.Map({True: not(not(not(False)))})) | ((_dafny.Map({True: False})) | (_dafny.Map({True: False})))

    @staticmethod
    def fm17(p0, p1, globalState):
        return D5_DC13(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkgiwheqr"))), (303) * (832), not (True) or (True))

    @staticmethod
    def fm18(p0, globalState):
        return _dafny.Map({(_dafny.Set({_dafny.Map({True: not(True)}), _dafny.Map({True: True})})) != (_dafny.Set({_dafny.Map({True: False}), _dafny.Map({False: not(False)})})): _dafny.CodePoint('u')})

    @staticmethod
    def m0(p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_25_v0_: _dafny.Set
        d_25_v0_ = _dafny.Set({p0})
        d_26_v1_: _dafny.Array
        nw0_ = _dafny.Array(False, 21)
        d_26_v1_ = nw0_
        d_27_v2_: _dafny.Array
        def lambda0_(d_28_i0_):
            return _dafny.SeqWithoutIsStrInference([468])

        init0_ = lambda0_
        nw1_ = _dafny.Array(None, 16)
        for i0_0_ in range(nw1_.length(0)):
            nw1_[i0_0_] = init0_(i0_0_)
        d_27_v2_ = nw1_
        d_29_v3_: _dafny.Map
        d_29_v3_ = _dafny.Map({p0: p0})
        d_30_v4_: _dafny.Seq
        d_30_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knce"))
        d_31_v5_: _dafny.Map
        d_31_v5_ = _dafny.Map({p0: d_30_v4_})
        d_32_v6_: _dafny.Array
        nw2_ = _dafny.Array(None, 22)
        nw2_[int(0)] = False
        nw2_[int(1)] = (D9_DC26(p0, d_25_v0_, d_26_v1_, d_27_v2_)).cf51
        nw2_[int(2)] = p0
        nw2_[int(3)] = p0
        nw2_[int(4)] = ((d_29_v3_)[default__.fm1(True, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_33_i1_ in range(default__.abs(176))]), globalState)] if (default__.fm1(True, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_34_i1_ in range(default__.abs(176))]), globalState)) in (d_29_v3_) else not(p0))
        nw2_[int(5)] = p0
        nw2_[int(6)] = p0
        nw2_[int(7)] = p0
        nw2_[int(8)] = p0
        nw2_[int(9)] = p0
        nw2_[int(10)] = p0
        nw2_[int(11)] = default__.fm1(p0, ((d_31_v5_)[p0] if (p0) in (d_31_v5_) else d_30_v4_), globalState)
        nw2_[int(12)] = p0
        nw2_[int(13)] = p0
        nw2_[int(14)] = p0
        nw2_[int(15)] = False
        nw2_[int(16)] = p0
        nw2_[int(17)] = p0
        nw2_[int(18)] = p0
        nw2_[int(19)] = p0
        nw2_[int(20)] = p0
        nw2_[int(21)] = p0
        d_32_v6_ = nw2_
        d_35_v7_: _dafny.Seq
        d_35_v7_ = _dafny.SeqWithoutIsStrInference([d_32_v6_, d_32_v6_])
        d_36_v8_: int
        d_36_v8_ = 0
        d_37_v9_: _dafny.Map
        d_37_v9_ = _dafny.Map({not((p0 if p0 else p0)): (d_35_v7_).set(default__.safeIndex(d_36_v8_, len(d_35_v7_)), d_32_v6_)})
        d_37_v9_ = (d_37_v9_).set(not(p0), d_35_v7_)
        d_38_v10_: _dafny.Seq
        d_38_v10_ = _dafny.SeqWithoutIsStrInference([d_36_v8_])
        d_39_v11_: D5
        d_39_v11_ = D5_DC13(d_36_v8_, (0) - (d_36_v8_), p0)
        pat_let_tv0_ = d_39_v11_
        pat_let_tv1_ = p0
        pat_let_tv2_ = p0
        def lambda1_(source2_):
            if source2_.is_DC5:
                d_40___mcc_h0_ = source2_.cf5
                d_41___mcc_h1_ = source2_.cf6
                d_42___mcc_h2_ = source2_.cf7
                d_43___mcc_h3_ = source2_.cf8
                d_44_cf8_ = d_43___mcc_h3_
                d_45_cf7_ = d_42___mcc_h2_
                d_46_cf6_ = d_41___mcc_h1_
                d_47_cf5_ = d_40___mcc_h0_
                return _dafny.Set({_dafny.Map({d_47_cf5_: _dafny.CodePoint('a')})})
            elif source2_.is_DC6:
                d_48___mcc_h4_ = source2_.cf9
                d_49___mcc_h5_ = source2_.cf10
                d_50___mcc_h6_ = source2_.cf11
                d_51___mcc_h7_ = source2_.cf12
                d_52_cf12_ = d_51___mcc_h7_
                d_53_cf11_ = d_50___mcc_h6_
                d_54_cf10_ = d_49___mcc_h5_
                d_55_cf9_ = d_48___mcc_h4_
                def iife4_():
                    def iife6_():
                        coll6_ = _dafny.Map()
                        compr_6_: _dafny.Map
                        for compr_6_ in (_dafny.Map({_dafny.Map({d_54_cf10_: _dafny.CodePoint('e')}): d_52_cf12_})).keys.Elements:
                            d_56_v12_: _dafny.Map = compr_6_
                            if (d_56_v12_) in (_dafny.Map({_dafny.Map({d_54_cf10_: _dafny.CodePoint('e')}): d_52_cf12_})):
                                coll6_[d_56_v12_] = _dafny.CodePoint('t')
                        return _dafny.Map(coll6_)
                    coll4_ = _dafny.Set()
                    def iife5_():
                        coll5_ = _dafny.Map()
                        compr_5_: _dafny.Map
                        for compr_5_ in (_dafny.Map({_dafny.Map({d_54_cf10_: _dafny.CodePoint('e')}): d_52_cf12_})).keys.Elements:
                            d_56_v12_: _dafny.Map = compr_5_
                            if (d_56_v12_) in (_dafny.Map({_dafny.Map({d_54_cf10_: _dafny.CodePoint('e')}): d_52_cf12_})):
                                coll5_[d_56_v12_] = _dafny.CodePoint('t')
                        return _dafny.Map(coll5_)
                    compr_4_: _dafny.Map
                    for compr_4_ in (iife5_()
                    ).keys.Elements:
                        d_57_v13_: _dafny.Map = compr_4_
                        if (d_57_v13_) in (iife6_()
                        ):
                            coll4_ = coll4_.union(_dafny.Set([d_57_v13_]))
                    return _dafny.Set(coll4_)
                return (iife4_()
                ).intersection(_dafny.Set({_dafny.Map({True: _dafny.CodePoint('k')}), _dafny.Map({d_54_cf10_: _dafny.CodePoint('h')}), _dafny.Map({(pat_let_tv0_).cf21: _dafny.CodePoint('f')})}))
            elif source2_.is_DC4:
                d_58___mcc_h8_ = source2_.cf4
                d_59_cf4_ = d_58___mcc_h8_
                return _dafny.Set({_dafny.Map({pat_let_tv1_: _dafny.CodePoint('a')}), _dafny.Map({False: _dafny.CodePoint('b')})})
            elif True:
                d_60___mcc_h9_ = source2_.cf13
                d_61_cf13_ = d_60___mcc_h9_
                def iife7_():
                    coll7_ = _dafny.Set()
                    compr_7_: _dafny.Map
                    for compr_7_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: _dafny.CodePoint('h')}) for d_62_i2_ in range(default__.abs(975))])).Elements:
                        d_63_v14_: _dafny.Map = compr_7_
                        if (d_63_v14_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: _dafny.CodePoint('h')}) for d_62_i2_ in range(default__.abs(975))])):
                            coll7_ = coll7_.union(_dafny.Set([d_63_v14_]))
                    return _dafny.Set(coll7_)
                return (_dafny.Set({_dafny.Map({pat_let_tv2_: _dafny.CodePoint('x')})})) - (iife7_()
                )

        (globalState).f2 = len(lambda1_(D2_DC6(len(d_38_v10_), (d_39_v11_).cf21, p0, d_36_v8_)))
        d_64_v15_: int
        d_64_v15_ = d_36_v8_
        d_65_v16_: C3
        nw3_ = C3()
        nw3_.ctor__(d_30_v4_, d_64_v15_, -140)
        d_65_v16_ = nw3_
        d_66_v17_: _dafny.Array
        nw4_ = _dafny.Array(None, 12)
        nw4_[int(0)] = d_65_v16_
        nw4_[int(1)] = d_65_v16_
        nw4_[int(2)] = d_65_v16_
        nw4_[int(3)] = d_65_v16_
        nw4_[int(4)] = d_65_v16_
        nw4_[int(5)] = d_65_v16_
        nw4_[int(6)] = d_65_v16_
        nw4_[int(7)] = d_65_v16_
        nw4_[int(8)] = d_65_v16_
        nw4_[int(9)] = d_65_v16_
        nw4_[int(10)] = d_65_v16_
        nw4_[int(11)] = d_65_v16_
        d_66_v17_ = nw4_
        index0_ = default__.safeIndex(999, (d_66_v17_).length(0))
        (d_66_v17_)[index0_] = d_65_v16_
        index1_ = default__.safeIndex(999, (d_66_v17_).length(0))
        rhs0_ = (d_36_v8_) - ((len((d_65_v16_).f19)) + ((0) - (d_36_v8_)))
        rhs1_ = d_65_v16_
        lhs0_ = globalState
        lhs1_ = d_66_v17_
        lhs2_ = default__.safeIndex(999, (d_66_v17_).length(0))
        lhs0_.f2 = rhs0_
        lhs1_[lhs2_] = rhs1_
        hi0_ = d_36_v8_
        for d_67_i3_ in range(d_36_v8_, hi0_):
            if p0:
                d_68_v18_: C2
                nw5_ = C2()
                nw5_.ctor__()
                d_68_v18_ = nw5_
                d_69_v19_: _dafny.MultiSet
                d_69_v19_ = _dafny.MultiSet([p0])
                d_70_v20_: _dafny.Map
                d_70_v20_ = _dafny.Map({d_69_v19_: d_39_v11_})
                d_70_v20_ = (d_70_v20_).set(d_69_v19_, D5_DC13(d_67_i3_, len(d_38_v10_), p0))
                (globalState).f9 = d_67_i3_
                (globalState).f0 = False
                index2_ = default__.safeIndex(701, (d_32_v6_).length(0))
                (d_32_v6_)[index2_] = p0
                index3_ = default__.safeIndex(701, (d_32_v6_).length(0))
                (d_32_v6_)[index3_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ph"))) != (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_71_i4_ in range(default__.abs(602))]))
            elif True:
                d_72_v21_: T0
                nw6_ = C0()
                nw6_.ctor__(d_64_v15_, 625)
                d_72_v21_ = nw6_
                d_73_v22_: _dafny.Seq
                d_73_v22_ = _dafny.SeqWithoutIsStrInference([d_72_v21_, d_72_v21_, d_72_v21_])
                d_74_v23_: _dafny.Seq
                d_74_v23_ = _dafny.SeqWithoutIsStrInference([(d_73_v22_).set(default__.safeIndex((d_72_v21_).f18, len(d_73_v22_)), d_72_v21_)])
                d_75_v24_: _dafny.Array
                nw7_ = _dafny.Array(None, 14)
                nw7_[int(0)] = d_74_v23_
                nw7_[int(1)] = _dafny.SeqWithoutIsStrInference([d_73_v22_, d_73_v22_])
                nw7_[int(2)] = (d_74_v23_) + (d_74_v23_)
                nw7_[int(3)] = (d_74_v23_) + (d_74_v23_)
                nw7_[int(4)] = d_74_v23_
                nw7_[int(5)] = (d_74_v23_ if p0 else d_74_v23_)
                nw7_[int(6)] = d_74_v23_
                nw7_[int(7)] = d_74_v23_
                nw7_[int(8)] = (d_74_v23_) + (d_74_v23_)
                nw7_[int(9)] = d_74_v23_
                nw7_[int(10)] = d_74_v23_
                nw7_[int(11)] = (d_74_v23_) + (_dafny.SeqWithoutIsStrInference([d_73_v22_]))
                nw7_[int(12)] = d_74_v23_
                nw7_[int(13)] = d_74_v23_
                d_75_v24_ = nw7_
                index4_ = default__.safeIndex(327, (d_75_v24_).length(0))
                (d_75_v24_)[index4_] = d_74_v23_
                index5_ = default__.safeIndex(327, (d_75_v24_).length(0))
                (d_75_v24_)[index5_] = (d_74_v23_).set(default__.safeIndex(d_67_i3_, len(d_74_v23_)), (d_73_v22_) + (d_73_v22_))
                d_76_v25_: _dafny.Map
                d_76_v25_ = _dafny.Map({(d_38_v10_) + (default__.fm11(d_67_i3_, p0, (d_72_v21_).f18, globalState)): d_36_v8_})
                d_77_v26_: _dafny.Map
                d_77_v26_ = _dafny.Map({p0: d_36_v8_})
                d_78_v27_: _dafny.Seq
                d_78_v27_ = _dafny.SeqWithoutIsStrInference([d_77_v26_])
                d_79_v28_: _dafny.Set
                d_79_v28_ = _dafny.Set({d_36_v8_})
                d_80_v29_: _dafny.Map
                d_80_v29_ = _dafny.Map({d_36_v8_: (d_78_v27_)[default__.safeIndex(len(d_79_v28_), len(d_78_v27_))]})
                d_81_v30_: _dafny.MultiSet
                d_81_v30_ = _dafny.MultiSet([503])
                d_76_v25_ = (d_76_v25_).set(d_38_v10_, ((0) - (len(((d_80_v29_)[len((d_65_v16_).f19)] if (len((d_65_v16_).f19)) in (d_80_v29_) else d_77_v26_)))) - ((d_81_v30_).cardinality))
                d_82_v31_: _dafny.Array
                def lambda2_(d_83_p0_):
                    def lambda3_(d_84_i5_):
                        return (d_84_i5_) + (len(_dafny.SeqWithoutIsStrInference([d_83_p0_])))

                    return lambda3_

                init1_ = lambda2_(p0)
                nw8_ = _dafny.Array(None, 19)
                for i0_1_ in range(nw8_.length(0)):
                    nw8_[i0_1_] = init1_(i0_1_)
                d_82_v31_ = nw8_
                d_85_v32_: D4
                d_85_v32_ = D4_DC10(d_82_v31_)
                d_85_v32_ = d_85_v32_
                index6_ = default__.safeIndex(778, (d_82_v31_).length(0))
                (d_82_v31_)[index6_] = d_36_v8_
                index7_ = default__.safeIndex(778, (d_82_v31_).length(0))
                (d_82_v31_)[index7_] = d_36_v8_
                d_86_v33_: _dafny.Seq
                d_86_v33_ = _dafny.SeqWithoutIsStrInference([p0])
                d_87_v34_: str
                d_87_v34_ = _dafny.CodePoint('i')
                rhs2_ = (d_82_v31_)[default__.safeIndex(778, (d_82_v31_).length(0))]
                rhs3_ = _dafny.SeqWithoutIsStrInference([(d_87_v34_) in ((d_65_v16_).f19), p0])
                rhs4_ = default__.safeDivisionInt((d_82_v31_)[default__.safeIndex(778, (d_82_v31_).length(0))], (d_72_v21_).f18)
                rhs5_ = p0
                lhs3_ = globalState
                lhs4_ = globalState
                lhs5_ = globalState
                lhs3_.f2 = rhs2_
                d_86_v33_ = rhs3_
                lhs4_.f9 = rhs4_
                lhs5_.f3 = rhs5_
            d_88_v35_: C2
            nw9_ = C2()
            nw9_.ctor__()
            d_88_v35_ = nw9_
            (globalState).f2 = d_36_v8_
            d_89_v36_: str
            d_89_v36_ = _dafny.CodePoint('n')
            d_90_v37_: C1
            nw10_ = C1()
            nw10_.ctor__(d_89_v36_)
            d_90_v37_ = nw10_
            d_91_v38_: _dafny.Seq
            d_91_v38_ = _dafny.SeqWithoutIsStrInference([d_90_v37_, d_90_v37_, d_90_v37_])
            rhs6_ = d_91_v38_
            rhs7_ = (d_67_i3_) < (((d_38_v10_)[default__.safeIndex(d_36_v8_, len(d_38_v10_))]) * (d_67_i3_))
            rhs8_ = ((d_65_v16_).f19) < (((d_65_v16_).f19) + ((d_65_v16_).f19))
            rhs9_ = d_67_i3_
            lhs6_ = globalState
            lhs7_ = globalState
            lhs8_ = globalState
            d_91_v38_ = rhs6_
            lhs6_.f0 = rhs7_
            lhs7_.f10 = rhs8_
            lhs8_.f9 = rhs9_
        d_92_v39_: _dafny.MultiSet
        d_92_v39_ = _dafny.MultiSet([p0, p0, p0])
        if (d_92_v39_).ispropersubset(_dafny.MultiSet([p0, p0])):
            d_93_v40_: _dafny.Array
            def lambda4_(d_94_v8_):
                def lambda5_(d_95_i6_):
                    return (d_95_i6_) + (d_94_v8_)

                return lambda5_

            init2_ = lambda4_(d_36_v8_)
            nw11_ = _dafny.Array(None, 25)
            for i0_2_ in range(nw11_.length(0)):
                nw11_[i0_2_] = init2_(i0_2_)
            d_93_v40_ = nw11_
            d_93_v40_ = d_93_v40_
            d_96_v41_: _dafny.Map
            d_96_v41_ = _dafny.Map({d_36_v8_: (0) - (d_36_v8_)})
            d_96_v41_ = (d_96_v41_).set(d_36_v8_, d_36_v8_)
            (globalState).f9 = ((-709) * (d_36_v8_)) - ((d_36_v8_) * (d_36_v8_))
            (globalState).f9 = d_36_v8_
            d_97_v42_: D6
            d_97_v42_ = D6_DC16(d_36_v8_, p0, d_36_v8_, d_36_v8_, p0)
            (globalState).f9 = ((d_97_v42_).cf28) + (-774)
        elif True:
            d_98_v43_: _dafny.Array
            def lambda6_(d_99_v8_):
                def lambda7_(d_100_i7_):
                    return (d_100_i7_) - (d_99_v8_)

                return lambda7_

            init3_ = lambda6_(d_36_v8_)
            nw12_ = _dafny.Array(None, 2)
            for i0_3_ in range(nw12_.length(0)):
                nw12_[i0_3_] = init3_(i0_3_)
            d_98_v43_ = nw12_
            nw13_ = _dafny.Array(int(0), 11)
            d_98_v43_ = nw13_
            d_101_v44_: _dafny.MultiSet
            d_101_v44_ = _dafny.MultiSet([d_36_v8_, default__.fm0(globalState)])
            d_102_v45_: _dafny.Set
            d_102_v45_ = _dafny.Set({d_36_v8_})
            d_103_v46_: C0
            nw14_ = C0()
            nw14_.ctor__(d_36_v8_, len(d_102_v45_))
            d_103_v46_ = nw14_
            if (not(not(not(p0)))) and (not(p0)):
                d_104_v47_: C0
                nw15_ = C0()
                nw15_.ctor__(d_64_v15_, 178)
                d_104_v47_ = nw15_
                index8_ = default__.safeIndex(642, (d_98_v43_).length(0))
                (d_98_v43_)[index8_] = (d_36_v8_ if p0 else d_36_v8_)
                d_105_v48_: _dafny.Map
                d_105_v48_ = _dafny.Map({p0: d_29_v3_})
                index9_ = default__.safeIndex(642, (d_98_v43_).length(0))
                rhs10_ = d_103_v46_
                rhs11_ = True
                rhs12_ = (0) - ((len(d_105_v48_)) + (d_36_v8_))
                lhs9_ = globalState
                lhs10_ = d_98_v43_
                lhs11_ = default__.safeIndex(642, (d_98_v43_).length(0))
                d_104_v47_ = rhs10_
                lhs9_.f0 = rhs11_
                lhs10_[lhs11_] = rhs12_
                d_106_v49_: C0
                nw16_ = C0()
                nw16_.ctor__(d_64_v15_, (d_98_v43_)[default__.safeIndex(642, (d_98_v43_).length(0))])
                d_106_v49_ = nw16_
                (globalState).f3 = ((d_29_v3_)[p0] if (p0) in (d_29_v3_) else False)
                (globalState).f2 = (d_98_v43_)[default__.safeIndex(642, (d_98_v43_).length(0))]
                d_107_v50_: _dafny.Seq
                d_107_v50_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vckfvwohs"))) < ((d_65_v16_).f19)])
                d_107_v50_ = d_107_v50_
            elif True:
                (globalState).f2 = (0) - (((d_36_v8_) * (d_36_v8_)) - (default__.safeDivisionInt(d_36_v8_, d_36_v8_)))
                index10_ = default__.safeIndex(665, (d_26_v1_).length(0))
                (d_26_v1_)[index10_] = p0
                index11_ = default__.safeIndex(665, (d_26_v1_).length(0))
                (d_26_v1_)[index11_] = p0
                d_108_v51_: D11
                d_108_v51_ = D11_DC32(d_36_v8_)
                d_109_v52_: str
                d_109_v52_ = _dafny.CodePoint('q')
                d_110_v53_: C1
                nw17_ = C1()
                nw17_.ctor__(d_109_v52_)
                d_110_v53_ = nw17_
                rhs13_ = default__.safeDivisionInt(d_36_v8_, d_36_v8_)
                rhs14_ = (d_26_v1_)[default__.safeIndex(665, (d_26_v1_).length(0))]
                rhs15_ = d_108_v51_
                rhs16_ = d_110_v53_
                lhs12_ = globalState
                d_36_v8_ = rhs13_
                lhs12_.f0 = rhs14_
                d_108_v51_ = rhs15_
                d_110_v53_ = rhs16_
                d_111_v54_: _dafny.Map
                d_111_v54_ = _dafny.Map({d_110_v53_: d_103_v46_})
                d_112_v55_: _dafny.Seq
                d_112_v55_ = _dafny.SeqWithoutIsStrInference([d_103_v46_])
                d_113_v56_: _dafny.Array
                nw18_ = _dafny.Array(None, 10)
                nw18_[int(0)] = (d_103_v46_ if p0 else d_103_v46_)
                nw18_[int(1)] = ((d_111_v54_)[d_110_v53_] if (d_110_v53_) in (d_111_v54_) else d_103_v46_)
                nw18_[int(2)] = d_103_v46_
                nw18_[int(3)] = d_103_v46_
                nw18_[int(4)] = d_103_v46_
                nw18_[int(5)] = d_103_v46_
                nw18_[int(6)] = (d_112_v55_)[default__.safeIndex(len(d_38_v10_), len(d_112_v55_))]
                nw18_[int(7)] = d_103_v46_
                nw18_[int(8)] = d_103_v46_
                nw18_[int(9)] = d_103_v46_
                d_113_v56_ = nw18_
                index12_ = default__.safeIndex(611, (d_113_v56_).length(0))
                (d_113_v56_)[index12_] = d_103_v46_
                d_114_v57_: D2
                d_114_v57_ = D2_DC6(d_36_v8_, True, p0, d_36_v8_)
                d_115_v58_: _dafny.Seq
                d_115_v58_ = _dafny.SeqWithoutIsStrInference([d_114_v57_])
                d_116_v59_: _dafny.MultiSet
                d_116_v59_ = _dafny.MultiSet([d_39_v11_, d_39_v11_, default__.fm17(len(d_30_v4_), d_115_v58_, globalState)])
                d_117_v60_: D1
                d_117_v60_ = D1_DC3(D1_DC2((d_26_v1_)[default__.safeIndex(665, (d_26_v1_).length(0))]))
                d_118_v61_: D6
                d_118_v61_ = D6_DC15(d_36_v8_, d_36_v8_, default__.fm11((0) - (d_36_v8_), False, 218, globalState), d_36_v8_, _dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))))}))
                pat_let_tv3_ = p0
                index13_ = default__.safeIndex(611, (d_113_v56_).length(0))
                def iife8_(_pat_let0_0):
                    def iife9_(d_119_dt__update__tmp_h3_):
                        def iife10_(_pat_let1_0):
                            def iife11_(d_120_dt__update_hcf3_h0_):
                                return D1_DC3(d_120_dt__update_hcf3_h0_)
                            return iife11_(_pat_let1_0)
                        return iife10_(D1_DC2(pat_let_tv3_))
                    return iife9_(_pat_let0_0)
                rhs17_ = d_103_v46_
                rhs18_ = (d_116_v59_) | ((d_116_v59_) - (d_116_v59_))
                rhs19_ = len(default__.fm18(d_36_v8_, globalState))
                rhs20_ = (0) - ((d_118_v61_).cf24)
                rhs21_ = iife8_(d_117_v60_)
                lhs13_ = d_113_v56_
                lhs14_ = default__.safeIndex(611, (d_113_v56_).length(0))
                lhs15_ = globalState
                lhs16_ = globalState
                lhs13_[lhs14_] = rhs17_
                d_116_v59_ = rhs18_
                lhs15_.f9 = rhs19_
                lhs16_.f9 = rhs20_
                d_117_v60_ = rhs21_
                d_121_v62_: _dafny.Map
                d_121_v62_ = _dafny.Map({d_109_v52_: not(not(p0))})
                d_121_v62_ = (d_121_v62_).set((d_110_v53_).f20, p0)
            (globalState).f10 = p0
            d_122_v63_: C2
            nw19_ = C2()
            nw19_.ctor__()
            d_122_v63_ = nw19_
            d_123_v64_: _dafny.MultiSet
            d_123_v64_ = _dafny.MultiSet([d_122_v63_, d_122_v63_, d_122_v63_, d_122_v63_, d_122_v63_])
            d_123_v64_ = d_123_v64_
        d_124_v65_: str
        d_124_v65_ = _dafny.CodePoint('j')
        d_125_v66_: C0
        nw20_ = C0()
        nw20_.ctor__(d_64_v15_, d_36_v8_)
        d_125_v66_ = nw20_
        index14_ = default__.safeIndex(684, (d_32_v6_).length(0))
        (d_32_v6_)[index14_] = p0
        d_126_v67_: _dafny.Array
        nw21_ = _dafny.Array(_dafny.Set({}), 14)
        d_126_v67_ = nw21_
        d_127_v69_: _dafny.Seq
        def iife12_():
            coll8_ = _dafny.Set()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(784, 243):
                d_128_v68_: int = compr_8_
                if ((784) <= (d_128_v68_)) and ((d_128_v68_) < (243)):
                    coll8_ = coll8_.union(_dafny.Set([default__.safeDivisionInt(d_128_v68_, d_36_v8_)]))
            return _dafny.Set(coll8_)
        d_127_v69_ = _dafny.SeqWithoutIsStrInference([iife12_()
        ])
        index15_ = default__.safeIndex(925, (d_126_v67_).length(0))
        (d_126_v67_)[index15_] = (d_127_v69_)[default__.safeIndex(77, len(d_127_v69_))]
        d_129_v70_: _dafny.Map
        d_129_v70_ = _dafny.Map({p0: d_36_v8_})
        d_130_v71_: _dafny.MultiSet
        d_130_v71_ = _dafny.MultiSet([d_125_v66_])
        d_131_v72_: C2
        nw22_ = C2()
        nw22_.ctor__()
        d_131_v72_ = nw22_
        d_132_v73_: _dafny.Map
        d_132_v73_ = _dafny.Map({p0: d_131_v72_})
        d_133_v74_: _dafny.Set
        d_133_v74_ = _dafny.Set({d_36_v8_, len(d_132_v73_), default__.fm0(globalState)})
        index16_ = default__.safeIndex(684, (d_32_v6_).length(0))
        index17_ = default__.safeIndex(925, (d_126_v67_).length(0))
        rhs22_ = d_30_v4_
        rhs23_ = ((d_65_v16_).f19)[default__.safeIndex(default__.safeModuloInt(len(d_129_v70_), (d_130_v71_).cardinality), len((d_65_v16_).f19))]
        rhs24_ = (d_125_v66_ if default__.fm1(p0, d_30_v4_, globalState) else d_125_v66_)
        rhs25_ = p0
        rhs26_ = ((d_133_v74_).intersection(d_133_v74_)).intersection((d_133_v74_).intersection(d_133_v74_))
        lhs17_ = d_32_v6_
        lhs18_ = default__.safeIndex(684, (d_32_v6_).length(0))
        lhs19_ = d_126_v67_
        lhs20_ = default__.safeIndex(925, (d_126_v67_).length(0))
        d_30_v4_ = rhs22_
        d_124_v65_ = rhs23_
        d_125_v66_ = rhs24_
        lhs17_[lhs18_] = rhs25_
        lhs19_[lhs20_] = rhs26_
        r0 = (d_36_v8_ if (_dafny.CodePoint('g')) in (d_30_v4_) else d_36_v8_)
        d_134_v75_: _dafny.Map
        d_134_v75_ = _dafny.Map({(d_32_v6_)[default__.safeIndex(684, (d_32_v6_).length(0))]: d_124_v65_})
        d_135_v76_: _dafny.Map
        d_135_v76_ = d_134_v75_
        d_136_v77_: _dafny.Map
        d_136_v77_ = _dafny.Map({d_36_v8_: (d_134_v75_)})
        r1 = len((d_136_v77_).set(default__.safeModuloInt(d_36_v8_, d_36_v8_), _dafny.Map({p0: d_124_v65_})))
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_137_v0_: bool
        d_137_v0_ = True
        d_138_v1_: _dafny.Array
        def lambda8_(d_139_v0_):
            def lambda9_(d_140_i0_):
                return d_139_v0_

            return lambda9_

        init4_ = lambda8_(d_137_v0_)
        nw23_ = _dafny.Array(None, 21)
        for i0_4_ in range(nw23_.length(0)):
            nw23_[i0_4_] = init4_(i0_4_)
        d_138_v1_ = nw23_
        d_141_v2_: int
        d_141_v2_ = 213
        d_142_v3_: _dafny.Map
        d_142_v3_ = _dafny.Map({d_137_v0_: (0) - (d_141_v2_)})
        d_143_v4_: _dafny.Seq
        d_143_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
        d_144_v5_: _dafny.Map
        d_144_v5_ = _dafny.Map({d_137_v0_: d_137_v0_})
        d_145_globalState_: GlobalState
        nw24_ = GlobalState()
        nw24_.ctor__(True, _dafny.SeqWithoutIsStrInference([d_137_v0_]), 172, False, True, _dafny.CodePoint('f'), 669, d_138_v1_, _dafny.MultiSet([d_141_v2_, d_141_v2_]), 257, False, (d_138_v1_ if d_137_v0_ else d_138_v1_), d_142_v3_, d_143_v4_, False, (d_144_v5_) | (d_144_v5_), True)
        d_145_globalState_ = nw24_
        if d_137_v0_:
            d_142_v3_ = (d_142_v3_).set(d_137_v0_, d_141_v2_)
            d_146_v6_: _dafny.Array
            def lambda10_(d_147_i1_):
                return _dafny.SeqWithoutIsStrInference([False])

            init5_ = lambda10_
            nw25_ = _dafny.Array(None, 12)
            for i0_5_ in range(nw25_.length(0)):
                nw25_[i0_5_] = init5_(i0_5_)
            d_146_v6_ = nw25_
            d_148_v7_: _dafny.Seq
            d_148_v7_ = _dafny.SeqWithoutIsStrInference([d_137_v0_])
            index18_ = default__.safeIndex(676, (d_146_v6_).length(0))
            (d_146_v6_)[index18_] = d_148_v7_
            index19_ = default__.safeIndex(676, (d_146_v6_).length(0))
            (d_146_v6_)[index19_] = (d_148_v7_).set(default__.safeIndex(((0) - (d_141_v2_) if d_137_v0_ else d_141_v2_), len(d_148_v7_)), (d_148_v7_)[default__.safeIndex((0) - (d_141_v2_), len(d_148_v7_))])
            d_149_v8_: _dafny.Map
            d_149_v8_ = _dafny.Map({d_141_v2_: d_141_v2_})
            d_150_v9_: _dafny.Map
            d_150_v9_ = _dafny.Map({d_149_v8_: d_137_v0_})
            d_141_v2_ = (d_141_v2_) - (len(d_150_v9_))
            d_151_v10_: _dafny.Set
            d_151_v10_ = _dafny.Set({d_141_v2_})
            d_152_v11_: _dafny.Seq
            d_152_v11_ = _dafny.SeqWithoutIsStrInference([313])
            d_153_v12_: _dafny.Seq
            d_153_v12_ = _dafny.SeqWithoutIsStrInference([len(d_152_v11_)])
            d_154_v13_: _dafny.Array
            nw26_ = _dafny.Array(None, 28)
            nw26_[int(0)] = d_141_v2_
            nw26_[int(1)] = 77
            nw26_[int(2)] = (len((d_146_v6_)[default__.safeIndex(676, (d_146_v6_).length(0))])) * (d_141_v2_)
            nw26_[int(3)] = d_141_v2_
            nw26_[int(4)] = (default__.fm0(d_145_globalState_)) * ((0) - (d_141_v2_))
            nw26_[int(5)] = d_141_v2_
            nw26_[int(6)] = default__.fm0(d_145_globalState_)
            nw26_[int(7)] = d_141_v2_
            nw26_[int(8)] = d_141_v2_
            nw26_[int(9)] = default__.safeDivisionInt(355, d_141_v2_)
            nw26_[int(10)] = (len(d_151_v10_)) + (d_141_v2_)
            nw26_[int(11)] = d_141_v2_
            nw26_[int(12)] = d_141_v2_
            nw26_[int(13)] = 336
            nw26_[int(14)] = d_141_v2_
            nw26_[int(15)] = (default__.fm0(d_145_globalState_)) - (d_141_v2_)
            nw26_[int(16)] = d_141_v2_
            nw26_[int(17)] = d_141_v2_
            nw26_[int(18)] = d_141_v2_
            nw26_[int(19)] = default__.safeDivisionInt(d_141_v2_, d_141_v2_)
            nw26_[int(20)] = d_141_v2_
            nw26_[int(21)] = (d_141_v2_ if d_137_v0_ else d_141_v2_)
            nw26_[int(22)] = len(d_143_v4_)
            nw26_[int(23)] = d_141_v2_
            nw26_[int(24)] = default__.safeDivisionInt(d_141_v2_, len(d_152_v11_))
            nw26_[int(25)] = d_141_v2_
            nw26_[int(26)] = d_141_v2_
            nw26_[int(27)] = default__.fm0(d_145_globalState_)
            d_154_v13_ = nw26_
            d_155_v14_: _dafny.Map
            d_155_v14_ = _dafny.Map({-657: d_154_v13_})
            rhs27_ = d_138_v1_
            rhs28_ = ((d_141_v2_) * (d_141_v2_)) - (default__.safeDivisionInt(d_141_v2_, d_141_v2_))
            rhs29_ = ((d_155_v14_)[default__.safeModuloInt(d_141_v2_, d_141_v2_)] if (default__.safeModuloInt(d_141_v2_, d_141_v2_)) in (d_155_v14_) else d_154_v13_)
            lhs21_ = d_145_globalState_
            d_138_v1_ = rhs27_
            lhs21_.f9 = rhs28_
            d_154_v13_ = rhs29_
            d_156_v15_: _dafny.Array
            nw27_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_156_v15_ = nw27_
            d_156_v15_ = d_156_v15_
        elif True:
            d_157_v16_: str
            d_157_v16_ = _dafny.CodePoint('r')
            d_158_v17_: _dafny.Map
            d_158_v17_ = _dafny.Map({d_157_v16_: d_141_v2_})
            (d_145_globalState_).f2 = len((_dafny.Map({d_158_v17_: d_137_v0_})).set(d_158_v17_, d_137_v0_))
            d_159_v18_: _dafny.Seq
            d_159_v18_ = _dafny.SeqWithoutIsStrInference([d_137_v0_])
            d_159_v18_ = _dafny.SeqWithoutIsStrInference([not(d_137_v0_)])
            d_160_v20_: _dafny.MultiSet
            d_160_v20_ = _dafny.MultiSet([d_141_v2_, d_141_v2_, d_141_v2_, d_141_v2_, d_141_v2_])
            d_161_v21_: _dafny.MultiSet
            d_161_v21_ = _dafny.MultiSet([d_137_v0_, d_137_v0_])
            d_162_v22_: int
            d_163_v23_: int
            out0_: int
            out1_: int
            def iife13_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in (d_160_v20_).Elements:
                    d_164_v19_: int = compr_9_
                    if (d_164_v19_) in (d_160_v20_):
                        coll9_[default__.safeModuloInt(d_164_v19_, d_141_v2_)] = d_137_v0_
                return _dafny.Map(coll9_)
            out0_, out1_ = default__.m0((d_143_v4_) <= (d_143_v4_), (iife13_()
            ).set((d_161_v21_).cardinality, (d_159_v18_)[default__.safeIndex(d_141_v2_, len(d_159_v18_))]), d_145_globalState_)
            d_162_v22_ = out0_
            d_163_v23_ = out1_
            d_165_v24_: _dafny.Seq
            d_165_v24_ = _dafny.SeqWithoutIsStrInference([d_143_v4_, _dafny.SeqWithoutIsStrInference([d_157_v16_ for d_166_i2_ in range(default__.abs(783))]), d_143_v4_, d_143_v4_])
            if default__.fm1(False, (d_165_v24_)[default__.safeIndex(d_162_v22_, len(d_165_v24_))], d_145_globalState_):
                d_157_v16_ = d_157_v16_
                (d_145_globalState_).f3 = (default__.fm0(d_145_globalState_)) > (d_141_v2_)
                d_167_v25_: int
                d_167_v25_ = d_163_v23_
                d_168_v26_: _dafny.Map
                d_168_v26_ = _dafny.Map({(d_167_v25_): not(d_137_v0_)})
                d_169_v27_: int
                d_170_v28_: int
                out2_: int
                out3_: int
                out2_, out3_ = default__.m0(d_137_v0_, d_168_v26_, d_145_globalState_)
                d_169_v27_ = out2_
                d_170_v28_ = out3_
                (d_145_globalState_).f9 = d_170_v28_
                d_171_v29_: _dafny.Map
                d_171_v29_ = _dafny.Map({(d_159_v18_) + (_dafny.SeqWithoutIsStrInference([False])): d_137_v0_})
                d_172_v30_: _dafny.Seq
                d_172_v30_ = _dafny.SeqWithoutIsStrInference([d_169_v27_])
                d_173_v31_: _dafny.MultiSet
                d_173_v31_ = _dafny.MultiSet([d_172_v30_])
                (d_145_globalState_).f3 = not(((d_171_v29_)[d_159_v18_] if (d_159_v18_) in (d_171_v29_) else (_dafny.MultiSet([d_172_v30_, d_172_v30_])).issubset(d_173_v31_)))
            elif True:
                d_157_v16_ = d_157_v16_
                d_159_v18_ = (d_159_v18_) + (d_159_v18_)
                d_174_v32_: C1
                nw28_ = C1()
                nw28_.ctor__(d_157_v16_)
                d_174_v32_ = nw28_
                d_175_v33_: _dafny.Map
                d_175_v33_ = _dafny.Map({d_137_v0_: d_159_v18_})
                d_176_v34_: _dafny.Set
                d_176_v34_ = _dafny.Set({d_159_v18_, d_159_v18_, d_159_v18_, d_159_v18_, ((d_175_v33_)[(d_159_v18_)[default__.safeIndex(d_141_v2_, len(d_159_v18_))]] if ((d_159_v18_)[default__.safeIndex(d_141_v2_, len(d_159_v18_))]) in (d_175_v33_) else d_159_v18_)})
                (d_145_globalState_).f10 = (d_176_v34_).ispropersubset((d_176_v34_) | (d_176_v34_))
                d_142_v3_ = (d_142_v3_).set((d_137_v0_ if d_137_v0_ else d_137_v0_), d_163_v23_)
            rhs30_ = not (not(True)) or (True)
            rhs31_ = (_dafny.MultiSet((d_159_v18_) + (_dafny.SeqWithoutIsStrInference([d_137_v0_, d_137_v0_, d_137_v0_])))).isdisjoint(d_161_v21_)
            rhs32_ = (d_141_v2_) <= (d_162_v22_)
            lhs22_ = d_145_globalState_
            lhs23_ = d_145_globalState_
            lhs24_ = d_145_globalState_
            lhs22_.f10 = rhs30_
            lhs23_.f0 = rhs31_
            lhs24_.f0 = rhs32_
        d_177_v35_: _dafny.MultiSet
        d_177_v35_ = _dafny.MultiSet([d_141_v2_, d_141_v2_, d_141_v2_])
        d_178_v36_: D6
        d_178_v36_ = D6_DC16(d_141_v2_, d_137_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))), d_141_v2_, d_137_v0_)
        rhs33_ = ((d_177_v35_)[(0) - (default__.safeDivisionInt((d_178_v36_).cf31, d_141_v2_))] if ((0) - (default__.safeDivisionInt((d_178_v36_).cf31, d_141_v2_))) in (d_177_v35_) else (0) - (d_141_v2_))
        rhs34_ = d_141_v2_
        lhs25_ = d_145_globalState_
        lhs26_ = d_145_globalState_
        lhs25_.f9 = rhs33_
        lhs26_.f9 = rhs34_
        d_179_v37_: _dafny.Seq
        d_179_v37_ = _dafny.SeqWithoutIsStrInference([-496, d_141_v2_, d_141_v2_])
        d_179_v37_ = ((d_179_v37_) + (_dafny.SeqWithoutIsStrInference([d_141_v2_]))) + (d_179_v37_)
        d_180_v38_: D1
        d_180_v38_ = D1_DC2(False)
        d_181_v39_: _dafny.MultiSet
        d_181_v39_ = _dafny.MultiSet([d_137_v0_])
        d_182_v40_: _dafny.Set
        d_182_v40_ = _dafny.Set({d_137_v0_, d_137_v0_})
        pat_let_tv4_ = d_137_v0_
        pat_let_tv5_ = d_137_v0_
        pat_let_tv6_ = d_137_v0_
        pat_let_tv7_ = d_145_globalState_
        d_183_v41_: _dafny.Array
        nw29_ = _dafny.Array(None, 20)
        nw29_[int(0)] = d_180_v38_
        nw29_[int(1)] = d_180_v38_
        def iife14_(_pat_let2_0):
            def iife15_(d_184_dt__update__tmp_h0_):
                def iife16_(_pat_let3_0):
                    def iife17_(d_185_dt__update_hcf2_h0_):
                        return D1_DC2(d_185_dt__update_hcf2_h0_)
                    return iife17_(_pat_let3_0)
                return iife16_(pat_let_tv4_)
            return iife15_(_pat_let2_0)
        nw29_[int(2)] = iife14_(d_180_v38_)
        nw29_[int(3)] = d_180_v38_
        nw29_[int(4)] = d_180_v38_
        nw29_[int(5)] = d_180_v38_
        nw29_[int(6)] = D1_DC2(d_137_v0_)
        nw29_[int(7)] = D1_DC2(d_137_v0_)
        nw29_[int(8)] = d_180_v38_
        nw29_[int(9)] = d_180_v38_
        nw29_[int(10)] = d_180_v38_
        nw29_[int(11)] = d_180_v38_
        def iife18_(_pat_let4_0):
            def iife19_(d_186_dt__update__tmp_h1_):
                def iife20_(_pat_let5_0):
                    def iife21_(d_187_dt__update_hcf2_h1_):
                        return D1_DC2(d_187_dt__update_hcf2_h1_)
                    return iife21_(_pat_let5_0)
                return iife20_(pat_let_tv5_)
            return iife19_(_pat_let4_0)
        nw29_[int(12)] = iife18_(d_180_v38_)
        def iife22_(_pat_let6_0):
            def iife23_(d_188_dt__update__tmp_h2_):
                def iife24_(_pat_let7_0):
                    def iife25_(d_189_dt__update_hcf2_h2_):
                        return D1_DC2(d_189_dt__update_hcf2_h2_)
                    return iife25_(_pat_let7_0)
                return iife24_(default__.fm1(pat_let_tv6_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jol")), pat_let_tv7_))
            return iife23_(_pat_let6_0)
        nw29_[int(13)] = iife22_(d_180_v38_)
        nw29_[int(14)] = d_180_v38_
        nw29_[int(15)] = d_180_v38_
        nw29_[int(16)] = d_180_v38_
        nw29_[int(17)] = default__.fm5(d_141_v2_, ((d_181_v39_)[False] if (False) in (d_181_v39_) else len(d_182_v40_)), d_145_globalState_)
        nw29_[int(18)] = d_180_v38_
        nw29_[int(19)] = d_180_v38_
        d_183_v41_ = nw29_
        d_190_v42_: D2
        d_190_v42_ = D2_DC4(d_183_v41_)
        d_183_v41_ = (d_183_v41_ if (d_143_v4_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))) else (d_190_v42_).cf4)
        d_191_v43_: str
        d_191_v43_ = _dafny.CodePoint('y')
        d_192_v44_: C1
        nw30_ = C1()
        nw30_.ctor__(d_191_v43_)
        d_192_v44_ = nw30_
        d_193_v45_: D6
        d_193_v45_ = D6_DC14(d_192_v44_)
        pat_let_tv8_ = d_192_v44_
        pat_let_tv9_ = d_137_v0_
        pat_let_tv10_ = d_192_v44_
        def iife26_(_pat_let8_0):
            def iife27_(d_194_dt__update__tmp_h3_):
                def iife28_(_pat_let9_0):
                    def iife29_(d_195_dt__update_hcf22_h0_):
                        return D6_DC14(d_195_dt__update_hcf22_h0_)
                    return iife29_(_pat_let9_0)
                return iife28_((pat_let_tv8_ if pat_let_tv9_ else pat_let_tv10_))
            return iife27_(_pat_let8_0)
        source3_ = iife26_(d_193_v45_)
        if source3_.is_DC15:
            d_196___mcc_h0_ = source3_.cf23
            d_197___mcc_h1_ = source3_.cf24
            d_198___mcc_h2_ = source3_.cf25
            d_199___mcc_h3_ = source3_.cf26
            d_200___mcc_h4_ = source3_.cf27
            d_201_cf27_ = d_200___mcc_h4_
            d_202_cf26_ = d_199___mcc_h3_
            d_203_cf25_ = d_198___mcc_h2_
            d_204_cf24_ = d_197___mcc_h1_
            d_205_cf23_ = d_196___mcc_h0_
            d_206_v46_: _dafny.MultiSet
            d_206_v46_ = _dafny.MultiSet([(d_192_v44_).f20, d_191_v43_, d_191_v43_])
            rhs35_ = d_205_cf23_
            rhs36_ = (446 if (760) >= (len(_dafny.Map({default__.fm0(d_145_globalState_): d_205_cf23_}))) else ((d_206_v46_).cardinality if d_137_v0_ else d_205_cf23_))
            lhs27_ = d_145_globalState_
            lhs27_.f9 = rhs35_
            d_141_v2_ = rhs36_
            d_207_v47_: int
            d_207_v47_ = len(d_143_v4_)
            d_208_v48_: T0
            nw31_ = C0()
            nw31_.ctor__(d_207_v47_, d_205_cf23_)
            d_208_v48_ = nw31_
            d_209_v49_: _dafny.Array
            nw32_ = _dafny.Array(_dafny.Map({}), 16)
            d_209_v49_ = nw32_
            d_210_v50_: C0
            nw33_ = C0()
            nw33_.ctor__((0) - (d_202_cf26_), (d_208_v48_).f18)
            d_210_v50_ = nw33_
            d_211_v51_: _dafny.Map
            d_211_v51_ = _dafny.Map({d_137_v0_: d_210_v50_})
            index20_ = default__.safeIndex(787, (d_209_v49_).length(0))
            (d_209_v49_)[index20_] = (d_211_v51_) | (d_211_v51_)
            d_212_v52_: _dafny.MultiSet
            d_212_v52_ = _dafny.MultiSet([d_138_v1_])
            d_213_v53_: _dafny.Map
            d_213_v53_ = _dafny.Map({d_202_cf26_: d_202_cf26_})
            d_214_v54_: _dafny.Map
            d_214_v54_ = _dafny.Map({False: d_213_v53_})
            index21_ = default__.safeIndex(787, (d_209_v49_).length(0))
            rhs37_ = d_208_v48_
            rhs38_ = default__.fm11(default__.safeModuloInt(len(d_143_v4_), (_dafny.MultiSet(d_179_v37_)).cardinality), d_137_v0_, ((d_212_v52_)[d_138_v1_] if (d_138_v1_) in (d_212_v52_) else len(((d_214_v54_)[d_137_v0_] if (d_137_v0_) in (d_214_v54_) else _dafny.Map({(d_208_v48_).f18: d_141_v2_})))), d_145_globalState_)
            rhs39_ = d_192_v44_
            rhs40_ = d_211_v51_
            lhs28_ = d_209_v49_
            lhs29_ = default__.safeIndex(787, (d_209_v49_).length(0))
            d_208_v48_ = rhs37_
            d_203_cf25_ = rhs38_
            d_192_v44_ = rhs39_
            lhs28_[lhs29_] = rhs40_
            d_215_v55_: _dafny.Array
            nw34_ = _dafny.Array(int(0), 11)
            d_215_v55_ = nw34_
            d_215_v55_ = d_215_v55_
            (d_145_globalState_).f0 = d_137_v0_
        elif source3_.is_DC16:
            d_216___mcc_h5_ = source3_.cf28
            d_217___mcc_h6_ = source3_.cf29
            d_218___mcc_h7_ = source3_.cf30
            d_219___mcc_h8_ = source3_.cf31
            d_220___mcc_h9_ = source3_.cf32
            d_221_cf32_ = d_220___mcc_h9_
            d_222_cf31_ = d_219___mcc_h8_
            d_223_cf30_ = d_218___mcc_h7_
            d_224_cf29_ = d_217___mcc_h6_
            d_225_cf28_ = d_216___mcc_h5_
            d_226_v56_: int
            d_227_v57_: _dafny.MultiSet
            out4_: int
            out5_: _dafny.MultiSet
            out4_, out5_ = (d_192_v44_).m3(d_137_v0_, d_145_globalState_)
            d_226_v56_ = out4_
            d_227_v57_ = out5_
            d_228_v58_: _dafny.Seq
            d_228_v58_ = _dafny.SeqWithoutIsStrInference([d_224_cf29_])
            d_229_v59_: _dafny.Map
            d_229_v59_ = _dafny.Map({d_225_cf28_: default__.fm11(d_223_cf30_, d_137_v0_, d_225_cf28_, d_145_globalState_)})
            d_230_v60_: _dafny.Seq
            d_230_v60_ = _dafny.SeqWithoutIsStrInference([(d_179_v37_).set(default__.safeIndex(-15, len(d_179_v37_)), d_226_v56_)])
            d_231_v61_: _dafny.Map
            d_231_v61_ = _dafny.Map({d_143_v4_: d_221_cf32_})
            d_232_v62_: _dafny.Array
            nw35_ = _dafny.Array(None, 13)
            nw35_[int(0)] = d_179_v37_
            nw35_[int(1)] = d_179_v37_
            nw35_[int(2)] = d_179_v37_
            nw35_[int(3)] = d_179_v37_
            nw35_[int(4)] = (d_179_v37_) + (d_179_v37_)
            nw35_[int(5)] = d_179_v37_
            nw35_[int(6)] = _dafny.SeqWithoutIsStrInference([default__.fm0(d_145_globalState_), 295])
            nw35_[int(7)] = _dafny.SeqWithoutIsStrInference([d_222_cf31_, 173, (_dafny.MultiSet((d_228_v58_).set(default__.safeIndex(d_225_cf28_, len(d_228_v58_)), d_137_v0_))).cardinality, 3])
            nw35_[int(8)] = (d_179_v37_).set(default__.safeIndex(d_141_v2_, len(d_179_v37_)), d_225_cf28_)
            nw35_[int(9)] = ((d_229_v59_)[d_225_cf28_] if (d_225_cf28_) in (d_229_v59_) else d_179_v37_)
            nw35_[int(10)] = (d_179_v37_) + (_dafny.SeqWithoutIsStrInference([len(d_182_v40_) for d_233_i3_ in range(default__.abs(151))]))
            nw35_[int(11)] = d_179_v37_
            nw35_[int(12)] = ((d_230_v60_)[default__.safeIndex(d_222_cf31_, len(d_230_v60_))]).set(default__.safeIndex(len(d_231_v61_), len((d_230_v60_)[default__.safeIndex(d_222_cf31_, len(d_230_v60_))])), d_223_cf30_)
            d_232_v62_ = nw35_
            index22_ = default__.safeIndex(373, (d_232_v62_).length(0))
            (d_232_v62_)[index22_] = d_179_v37_
            index23_ = default__.safeIndex(373, (d_232_v62_).length(0))
            (d_232_v62_)[index23_] = d_179_v37_
            d_234_v63_: _dafny.Array
            nw36_ = _dafny.Array(int(0), 3)
            d_234_v63_ = nw36_
            d_235_v64_: D8
            d_235_v64_ = D8_DC23(d_225_cf28_, d_234_v63_, d_222_cf31_)
            d_236_v65_: _dafny.Set
            d_236_v65_ = _dafny.Set({d_225_cf28_})
            d_237_v66_: _dafny.Map
            d_237_v66_ = _dafny.Map({d_141_v2_: d_236_v65_})
            d_238_v67_: _dafny.Seq
            d_238_v67_ = _dafny.SeqWithoutIsStrInference([d_236_v65_, ((d_237_v66_)[len(d_144_v5_)] if (len(d_144_v5_)) in (d_237_v66_) else d_236_v65_)])
            source4_ = default__.fm12((d_235_v64_).cf46, d_238_v67_, d_223_cf30_, d_222_cf31_, d_145_globalState_)
            if source4_.is_DC13:
                d_239___mcc_h15_ = source4_.cf19
                d_240___mcc_h16_ = source4_.cf20
                d_241___mcc_h17_ = source4_.cf21
                d_242_cf21_ = d_241___mcc_h17_
                d_243_cf20_ = d_240___mcc_h16_
                d_244_cf19_ = d_239___mcc_h15_
                (d_145_globalState_).f2 = d_243_cf20_
                index24_ = default__.safeIndex(881, (d_234_v63_).length(0))
                (d_234_v63_)[index24_] = len(_dafny.Map({_dafny.CodePoint('p'): d_142_v3_}))
                d_245_v68_: _dafny.Array
                nw37_ = _dafny.Array(_dafny.CodePoint('D'), 5)
                d_245_v68_ = nw37_
                index25_ = default__.safeIndex(898, (d_245_v68_).length(0))
                (d_245_v68_)[index25_] = (d_192_v44_).f20
                d_246_v69_: _dafny.Map
                d_246_v69_ = _dafny.Map({d_222_cf31_: len(d_228_v58_)})
                d_247_v70_: _dafny.Map
                d_247_v70_ = _dafny.Map({d_246_v69_: d_143_v4_})
                d_248_v71_: _dafny.Map
                d_248_v71_ = _dafny.Map({d_243_cf20_: d_246_v69_})
                d_249_v72_: D5
                d_249_v72_ = D5_DC12(d_143_v4_)
                d_250_v73_: _dafny.Array
                nw38_ = _dafny.Array(None, 28)
                nw38_[int(0)] = d_143_v4_
                nw38_[int(1)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfvx"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tunb")))
                nw38_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "laijoy"))
                nw38_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
                nw38_[int(4)] = d_143_v4_
                nw38_[int(5)] = d_143_v4_
                nw38_[int(6)] = ((d_247_v70_)[((d_248_v71_)[len((d_232_v62_)[default__.safeIndex(373, (d_232_v62_).length(0))])] if (len((d_232_v62_)[default__.safeIndex(373, (d_232_v62_).length(0))])) in (d_248_v71_) else d_246_v69_)] if (((d_248_v71_)[len((d_232_v62_)[default__.safeIndex(373, (d_232_v62_).length(0))])] if (len((d_232_v62_)[default__.safeIndex(373, (d_232_v62_).length(0))])) in (d_248_v71_) else d_246_v69_)) in (d_247_v70_) else d_143_v4_)
                nw38_[int(7)] = (d_143_v4_).set(default__.safeIndex(d_141_v2_, len(d_143_v4_)), (d_192_v44_).f20)
                nw38_[int(8)] = (d_143_v4_) + (d_143_v4_)
                nw38_[int(9)] = ((d_249_v72_).cf18).set(default__.safeIndex(d_226_v56_, len((d_249_v72_).cf18)), d_191_v43_)
                nw38_[int(10)] = d_143_v4_
                nw38_[int(11)] = (d_143_v4_ if d_224_cf29_ else d_143_v4_)
                nw38_[int(12)] = ((default__.fm6(d_221_cf32_, d_145_globalState_)).set(default__.safeIndex(d_226_v56_, len(default__.fm6(d_221_cf32_, d_145_globalState_))), (d_192_v44_).f20) if default__.fm1(d_224_cf29_, d_143_v4_, d_145_globalState_) else d_143_v4_)
                nw38_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrxbpbgci"))
                nw38_[int(14)] = d_143_v4_
                nw38_[int(15)] = d_143_v4_
                nw38_[int(16)] = d_143_v4_
                nw38_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uci"))
                nw38_[int(18)] = d_143_v4_
                nw38_[int(19)] = d_143_v4_
                nw38_[int(20)] = (D5_DC12(d_143_v4_)).cf18
                nw38_[int(21)] = d_143_v4_
                nw38_[int(22)] = d_143_v4_
                nw38_[int(23)] = d_143_v4_
                nw38_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smf"))
                nw38_[int(25)] = d_143_v4_
                nw38_[int(26)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrnoni"))
                nw38_[int(27)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abfgyma"))
                d_250_v73_ = nw38_
                d_251_v74_: _dafny.Set
                d_251_v74_ = _dafny.Set({d_228_v58_, _dafny.SeqWithoutIsStrInference([d_224_cf29_, d_137_v0_, True])})
                d_252_v75_: D9
                d_252_v75_ = D9_DC24(d_250_v73_)
                index26_ = default__.safeIndex(881, (d_234_v63_).length(0))
                index27_ = default__.safeIndex(898, (d_245_v68_).length(0))
                rhs41_ = ((d_142_v3_)[(d_251_v74_).ispropersubset(d_251_v74_)] if ((d_251_v74_).ispropersubset(d_251_v74_)) in (d_142_v3_) else d_223_cf30_)
                rhs42_ = default__.safeModuloInt((811) - (len(d_182_v40_)), d_225_cf28_)
                rhs43_ = d_137_v0_
                rhs44_ = (d_192_v44_).f20
                rhs45_ = (d_252_v75_).cf49
                lhs30_ = d_234_v63_
                lhs31_ = default__.safeIndex(881, (d_234_v63_).length(0))
                lhs32_ = d_145_globalState_
                lhs33_ = d_145_globalState_
                lhs34_ = d_245_v68_
                lhs35_ = default__.safeIndex(898, (d_245_v68_).length(0))
                lhs30_[lhs31_] = rhs41_
                lhs32_.f2 = rhs42_
                lhs33_.f3 = rhs43_
                lhs34_[lhs35_] = rhs44_
                d_250_v73_ = rhs45_
                d_253_v76_: C2
                nw39_ = C2()
                nw39_.ctor__()
                d_253_v76_ = nw39_
                d_254_v77_: D8
                d_254_v77_ = D8_DC22(d_253_v76_)
                d_255_v78_: int
                d_255_v78_ = d_226_v56_
                d_256_v79_: C0
                nw40_ = C0()
                nw40_.ctor__(d_255_v78_, d_223_cf30_)
                d_256_v79_ = nw40_
                rhs46_ = (d_228_v58_).set(default__.safeIndex(default__.safeModuloInt(d_141_v2_, (d_234_v63_)[default__.safeIndex(881, (d_234_v63_).length(0))]), len(d_228_v58_)), True)
                rhs47_ = d_221_cf32_
                rhs48_ = d_254_v77_
                rhs49_ = d_256_v79_
                lhs36_ = d_145_globalState_
                d_228_v58_ = rhs46_
                lhs36_.f10 = rhs47_
                d_254_v77_ = rhs48_
                d_256_v79_ = rhs49_
                d_257_v80_: D4
                d_257_v80_ = D4_DC9(d_256_v79_)
                d_257_v80_ = d_257_v80_
            elif True:
                d_258___mcc_h18_ = source4_.cf18
                d_259_cf18_ = d_258___mcc_h18_
                index28_ = default__.safeIndex(373, (d_232_v62_).length(0))
                (d_232_v62_)[index28_] = d_179_v37_
                d_260_v81_: _dafny.Map
                d_260_v81_ = _dafny.Map({-665: d_224_cf29_})
                d_261_v82_: _dafny.Map
                d_261_v82_ = _dafny.Map({d_260_v81_: d_141_v2_})
                d_262_v84_: _dafny.MultiSet
                d_262_v84_ = _dafny.MultiSet([d_260_v81_])
                d_263_v86_: D10
                d_263_v86_ = D10_DC28(_dafny.Map({d_260_v81_: (0) - (d_223_cf30_)}))
                d_264_v88_: _dafny.Array
                nw41_ = _dafny.Array(None, 27)
                nw41_[int(0)] = default__.fm13(d_145_globalState_)
                nw41_[int(1)] = (d_261_v82_) | (d_261_v82_)
                nw41_[int(2)] = d_261_v82_
                nw41_[int(3)] = d_261_v82_
                nw41_[int(4)] = d_261_v82_
                nw41_[int(5)] = (d_261_v82_) | (d_261_v82_)
                def iife30_():
                    coll10_ = _dafny.Map()
                    compr_10_: _dafny.Map
                    for compr_10_ in (d_262_v84_).Elements:
                        d_265_v83_: _dafny.Map = compr_10_
                        if (d_265_v83_) in (d_262_v84_):
                            coll10_[d_265_v83_] = d_225_cf28_
                    return _dafny.Map(coll10_)
                nw41_[int(6)] = (iife30_()
                ).set(d_260_v81_, len(default__.fm8(d_225_cf28_, -844, d_145_globalState_)))
                nw41_[int(7)] = (d_261_v82_) | (d_261_v82_)
                nw41_[int(8)] = d_261_v82_
                def iife31_():
                    coll11_ = _dafny.Map()
                    compr_11_: _dafny.Map
                    for compr_11_ in (d_262_v84_).Elements:
                        d_266_v85_: _dafny.Map = compr_11_
                        if (d_266_v85_) in (d_262_v84_):
                            coll11_[d_266_v85_] = d_223_cf30_
                    return _dafny.Map(coll11_)
                nw41_[int(9)] = (iife31_()
                ).set(d_260_v81_, d_222_cf31_)
                nw41_[int(10)] = _dafny.Map({d_260_v81_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uavoowfky")))})
                nw41_[int(11)] = d_261_v82_
                nw41_[int(12)] = (d_261_v82_) | (d_261_v82_)
                nw41_[int(13)] = ((d_263_v86_).cf56).set(_dafny.Map({d_226_v56_: d_137_v0_}), d_141_v2_)
                nw41_[int(14)] = d_261_v82_
                nw41_[int(15)] = d_261_v82_
                nw41_[int(16)] = d_261_v82_
                nw41_[int(17)] = d_261_v82_
                nw41_[int(18)] = (d_261_v82_) | (d_261_v82_)
                def iife32_():
                    coll12_ = _dafny.Map()
                    compr_12_: int
                    for compr_12_ in (d_236_v65_).Elements:
                        d_267_v87_: int = compr_12_
                        if (d_267_v87_) in (d_236_v65_):
                            coll12_[(d_267_v87_) - (d_223_cf30_)] = d_221_cf32_
                    return _dafny.Map(coll12_)
                nw41_[int(19)] = _dafny.Map({iife32_()
                : d_223_cf30_})
                nw41_[int(20)] = d_261_v82_
                nw41_[int(21)] = d_261_v82_
                nw41_[int(22)] = d_261_v82_
                nw41_[int(23)] = d_261_v82_
                nw41_[int(24)] = _dafny.Map({_dafny.Map({d_222_cf31_: d_221_cf32_}): d_225_cf28_})
                nw41_[int(25)] = (d_261_v82_) | (d_261_v82_)
                nw41_[int(26)] = (d_261_v82_) | (d_261_v82_)
                d_264_v88_ = nw41_
                index29_ = default__.safeIndex(509, (d_264_v88_).length(0))
                (d_264_v88_)[index29_] = d_261_v82_
                index30_ = default__.safeIndex(509, (d_264_v88_).length(0))
                def iife33_():
                    coll13_ = _dafny.Map()
                    compr_13_: _dafny.Map
                    for compr_13_ in (d_261_v82_).keys.Elements:
                        d_268_v89_: _dafny.Map = compr_13_
                        if (d_268_v89_) in (d_261_v82_):
                            coll13_[d_268_v89_] = len(_dafny.Map({len(_dafny.Map({d_143_v4_: -269})): len(_dafny.SeqWithoutIsStrInference([False]))}))
                    return _dafny.Map(coll13_)
                rhs50_ = ((_dafny.Map({d_260_v81_: d_141_v2_})) | (iife33_()
                )) | (d_261_v82_)
                rhs51_ = d_225_cf28_
                lhs37_ = d_264_v88_
                lhs38_ = default__.safeIndex(509, (d_264_v88_).length(0))
                lhs37_[lhs38_] = rhs50_
                d_222_cf31_ = rhs51_
                d_269_v90_: D2
                d_269_v90_ = D2_DC6(-696, d_221_cf32_, d_221_cf32_, d_226_v56_)
                d_269_v90_ = d_269_v90_
                d_225_cf28_ = (d_141_v2_) * (len(_dafny.Map({d_138_v1_: d_234_v63_})))
            d_270_v91_: C2
            nw42_ = C2()
            nw42_.ctor__()
            d_270_v91_ = nw42_
            d_270_v91_ = d_270_v91_
        elif source3_.is_DC17:
            d_271___mcc_h10_ = source3_.cf33
            d_272___mcc_h11_ = source3_.cf34
            d_273___mcc_h12_ = source3_.cf35
            d_274___mcc_h13_ = source3_.cf36
            d_275_cf36_ = d_274___mcc_h13_
            d_276_cf35_ = d_273___mcc_h12_
            d_277_cf34_ = d_272___mcc_h11_
            d_278_cf33_ = d_271___mcc_h10_
            (d_145_globalState_).f2 = (d_141_v2_) * ((-556) + (d_141_v2_))
            d_277_cf34_ = False
            d_279_v92_: _dafny.Seq
            d_279_v92_ = _dafny.SeqWithoutIsStrInference([d_179_v37_, d_179_v37_, _dafny.SeqWithoutIsStrInference([len(d_143_v4_)]), d_179_v37_])
            (d_145_globalState_).f10 = not(((d_279_v92_)[default__.safeIndex(-944, len(d_279_v92_))]) < (d_179_v37_))
            (d_145_globalState_).f0 = True
        elif True:
            d_280___mcc_h14_ = source3_.cf22
            d_281_cf22_ = d_280___mcc_h14_
            (d_145_globalState_).f0 = True
            (d_145_globalState_).f9 = d_141_v2_
            d_282_v93_: D9
            d_282_v93_ = D9_DC27(d_141_v2_)
            pat_let_tv11_ = d_141_v2_
            def iife34_(_pat_let10_0):
                def iife35_(d_283_dt__update__tmp_h4_):
                    def iife36_(_pat_let11_0):
                        def iife37_(d_284_dt__update_hcf55_h0_):
                            return D9_DC27(d_284_dt__update_hcf55_h0_)
                        return iife37_(_pat_let11_0)
                    return iife36_(pat_let_tv11_)
                return iife35_(_pat_let10_0)
            def iife38_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in (_dafny.SeqWithoutIsStrInference([d_141_v2_])).Elements:
                    d_285_v94_: int = compr_14_
                    if (d_285_v94_) in (_dafny.SeqWithoutIsStrInference([d_141_v2_])):
                        coll14_[(d_285_v94_) * (d_141_v2_)] = d_141_v2_
                return _dafny.Map(coll14_)
            if ((d_141_v2_) + ((iife34_(d_282_v93_)).cf55)) not in (iife38_()
            ):
                d_286_v95_: _dafny.Seq
                d_286_v95_ = _dafny.SeqWithoutIsStrInference([d_143_v4_, d_143_v4_, d_143_v4_, d_143_v4_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))])
                d_287_v96_: _dafny.Map
                d_287_v96_ = _dafny.Map({(d_286_v95_)[default__.safeIndex(d_141_v2_, len(d_286_v95_))]: not(d_137_v0_)})
                d_287_v96_ = d_287_v96_
                d_288_v97_: int
                d_289_v98_: _dafny.MultiSet
                out6_: int
                out7_: _dafny.MultiSet
                out6_, out7_ = (d_281_cf22_).m3(d_137_v0_, d_145_globalState_)
                d_288_v97_ = out6_
                d_289_v98_ = out7_
                d_290_v99_: int
                d_290_v99_ = (0) - (d_288_v97_)
                d_291_v100_: T0
                nw43_ = C0()
                nw43_.ctor__(d_290_v99_, len(d_143_v4_))
                d_291_v100_ = nw43_
                d_291_v100_ = d_291_v100_
                (d_145_globalState_).f0 = d_137_v0_
                d_292_v101_: _dafny.Seq
                d_292_v101_ = _dafny.SeqWithoutIsStrInference([d_137_v0_])
                (d_145_globalState_).f3 = (d_292_v101_) <= (_dafny.SeqWithoutIsStrInference([d_137_v0_]))
            elif True:
                d_293_v102_: C2
                nw44_ = C2()
                nw44_.ctor__()
                d_293_v102_ = nw44_
                nw45_ = C2()
                nw45_.ctor__()
                d_293_v102_ = nw45_
                d_294_v103_: C1
                nw46_ = C1()
                nw46_.ctor__((d_281_cf22_).f20)
                d_294_v103_ = nw46_
                d_295_v104_: int
                d_296_v105_: _dafny.MultiSet
                out8_: int
                out9_: _dafny.MultiSet
                out8_, out9_ = (d_281_cf22_).m3(not(d_137_v0_), d_145_globalState_)
                d_295_v104_ = out8_
                d_296_v105_ = out9_
                (d_145_globalState_).f10 = d_137_v0_
                (d_145_globalState_).f9 = d_141_v2_
            d_297_v106_: _dafny.Array
            def lambda11_(d_298_v0_):
                def lambda12_(d_299_i4_):
                    return (d_299_i4_) + (len(_dafny.SeqWithoutIsStrInference([d_298_v0_])))

                return lambda12_

            init6_ = lambda11_(d_137_v0_)
            nw47_ = _dafny.Array(None, 5)
            for i0_6_ in range(nw47_.length(0)):
                nw47_[i0_6_] = init6_(i0_6_)
            d_297_v106_ = nw47_
            index31_ = default__.safeIndex(565, (d_297_v106_).length(0))
            (d_297_v106_)[index31_] = default__.safeModuloInt(170, d_141_v2_)
            index32_ = default__.safeIndex(565, (d_297_v106_).length(0))
            (d_297_v106_)[index32_] = len(d_182_v40_)
        d_300_v107_: D9
        d_300_v107_ = D9_DC25(d_141_v2_)
        pat_let_tv12_ = d_141_v2_
        d_301_v108_: _dafny.Array
        nw48_ = _dafny.Array(None, 6)
        nw48_[int(0)] = d_300_v107_
        def iife39_(_pat_let12_0):
            def iife40_(d_302_dt__update__tmp_h5_):
                def iife41_(_pat_let13_0):
                    def iife42_(d_303_dt__update_hcf50_h0_):
                        return D9_DC25(d_303_dt__update_hcf50_h0_)
                    return iife42_(_pat_let13_0)
                return iife41_(pat_let_tv12_)
            return iife40_(_pat_let12_0)
        nw48_[int(1)] = iife39_(d_300_v107_)
        nw48_[int(2)] = d_300_v107_
        nw48_[int(3)] = (d_300_v107_ if d_137_v0_ else d_300_v107_)
        nw48_[int(4)] = d_300_v107_
        nw48_[int(5)] = d_300_v107_
        d_301_v108_ = nw48_
        index33_ = default__.safeIndex(481, (d_301_v108_).length(0))
        (d_301_v108_)[index33_] = d_300_v107_
        pat_let_tv13_ = d_143_v4_
        pat_let_tv14_ = d_143_v4_
        index34_ = default__.safeIndex(481, (d_301_v108_).length(0))
        def iife43_(_pat_let14_0):
            def iife44_(d_304_dt__update__tmp_h6_):
                def iife45_(_pat_let15_0):
                    def iife46_(d_305_dt__update_hcf50_h1_):
                        return D9_DC25(d_305_dt__update_hcf50_h1_)
                    return iife46_(_pat_let15_0)
                return iife45_(len((pat_let_tv13_) + (pat_let_tv14_)))
            return iife44_(_pat_let14_0)
        (d_301_v108_)[index34_] = iife43_(d_300_v107_)
        if (default__.safeDivisionInt(d_141_v2_, d_141_v2_)) != (len((_dafny.SeqWithoutIsStrInference([d_141_v2_]) if d_137_v0_ else _dafny.SeqWithoutIsStrInference([(0) - (d_141_v2_) for d_306_i5_ in range(default__.abs(453))])))):
            d_307_v109_: _dafny.Map
            d_307_v109_ = _dafny.Map({(d_192_v44_).f20: d_137_v0_})
            rhs52_ = d_307_v109_
            rhs53_ = d_137_v0_
            lhs39_ = d_145_globalState_
            d_307_v109_ = rhs52_
            lhs39_.f0 = rhs53_
            d_308_v110_: _dafny.Seq
            d_308_v110_ = _dafny.SeqWithoutIsStrInference([(d_143_v4_) + (d_143_v4_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), d_143_v4_])
            rhs54_ = d_179_v37_
            rhs55_ = (d_137_v0_ if (d_137_v0_ if d_137_v0_ else False) else (d_137_v0_) and (d_137_v0_))
            rhs56_ = d_308_v110_
            lhs40_ = d_145_globalState_
            d_179_v37_ = rhs54_
            lhs40_.f0 = rhs55_
            d_308_v110_ = rhs56_
            d_143_v4_ = d_143_v4_
            d_309_v111_: _dafny.Map
            d_309_v111_ = _dafny.Map({d_141_v2_: d_191_v43_})
            d_310_v112_: _dafny.Seq
            d_310_v112_ = _dafny.SeqWithoutIsStrInference([d_137_v0_, True, not(d_137_v0_), d_137_v0_])
            d_182_v40_ = default__.fm14((d_141_v2_) > (d_141_v2_), ((d_309_v111_)[len(d_310_v112_)] if (len(d_310_v112_)) in (d_309_v111_) else default__.fm7((d_192_v44_).f20, d_137_v0_, d_145_globalState_)), d_137_v0_, d_145_globalState_)
            if (default__.fm14(d_137_v0_, _dafny.CodePoint('c'), not(d_137_v0_), d_145_globalState_)) != (d_182_v40_):
                d_311_v113_: D2
                d_311_v113_ = D2_DC4(d_183_v41_)
                d_312_v114_: D2
                d_312_v114_ = D2_DC7(d_311_v113_)
                d_312_v114_ = d_312_v114_
                d_313_v115_: _dafny.Array
                def lambda13_(d_314_v109_):
                    def lambda14_(d_315_i6_):
                        return d_314_v109_

                    return lambda14_

                init7_ = lambda13_(d_307_v109_)
                nw49_ = _dafny.Array(None, 6)
                for i0_7_ in range(nw49_.length(0)):
                    nw49_[i0_7_] = init7_(i0_7_)
                d_313_v115_ = nw49_
                rhs57_ = len(d_143_v4_)
                rhs58_ = d_138_v1_
                rhs59_ = d_313_v115_
                lhs41_ = d_145_globalState_
                lhs42_ = d_145_globalState_
                lhs41_.f9 = rhs57_
                lhs42_.f11 = rhs58_
                d_313_v115_ = rhs59_
                rhs60_ = (d_137_v0_) == ((d_141_v2_) > (721))
                rhs61_ = default__.fm7((d_192_v44_).f20, d_137_v0_, d_145_globalState_)
                d_137_v0_ = rhs60_
                d_191_v43_ = rhs61_
                d_316_v116_: _dafny.Map
                d_316_v116_ = _dafny.Map({default__.fm1(default__.fm1(d_137_v0_, d_143_v4_, d_145_globalState_), default__.fm6(d_137_v0_, d_145_globalState_), d_145_globalState_): (d_143_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eph")))})
                d_316_v116_ = (d_316_v116_).set(d_137_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aaf")))
                (d_145_globalState_).f10 = not(((d_137_v0_) == (d_137_v0_) if not (d_137_v0_) or (d_137_v0_) else d_137_v0_))
            elif True:
                (d_145_globalState_).f10 = d_137_v0_
                (d_145_globalState_).f2 = default__.safeModuloInt(len(default__.fm8(d_141_v2_, 873, d_145_globalState_)), d_141_v2_)
                d_317_v117_: C2
                nw50_ = C2()
                nw50_.ctor__()
                d_317_v117_ = nw50_
                d_318_v118_: _dafny.MultiSet
                d_318_v118_ = _dafny.MultiSet([d_317_v117_])
                (d_145_globalState_).f0 = (_dafny.MultiSet([d_317_v117_])).issubset((d_318_v118_) - (d_318_v118_))
                d_319_v119_: _dafny.Map
                d_319_v119_ = _dafny.Map({default__.fm11(d_141_v2_, d_137_v0_, d_141_v2_, d_145_globalState_): ((0) - (d_141_v2_)) >= (d_141_v2_)})
                d_319_v119_ = d_319_v119_
                d_320_v120_: _dafny.Array
                def lambda15_(d_321_i7_):
                    return (d_321_i7_) - (len(_dafny.Set({_dafny.CodePoint('n')})))

                init8_ = lambda15_
                nw51_ = _dafny.Array(None, 14)
                for i0_8_ in range(nw51_.length(0)):
                    nw51_[i0_8_] = init8_(i0_8_)
                d_320_v120_ = nw51_
                d_322_v121_: D6
                d_322_v121_ = D6_DC15(len(d_143_v4_), d_141_v2_, d_179_v37_, d_141_v2_, _dafny.Set({d_141_v2_}))
                d_323_v122_: int
                d_323_v122_ = (d_322_v121_).cf23
                d_324_v123_: C0
                nw52_ = C0()
                nw52_.ctor__(d_323_v122_, len(d_143_v4_))
                d_324_v123_ = nw52_
                d_325_v124_: _dafny.Map
                d_325_v124_ = _dafny.Map({d_324_v123_: d_320_v120_})
                d_326_v125_: _dafny.Array
                nw53_ = _dafny.Array(None, 25)
                nw53_[int(0)] = d_320_v120_
                nw53_[int(1)] = ((d_325_v124_)[d_324_v123_] if (d_324_v123_) in (d_325_v124_) else d_320_v120_)
                nw53_[int(2)] = d_320_v120_
                nw53_[int(3)] = d_320_v120_
                nw53_[int(4)] = d_320_v120_
                nw53_[int(5)] = d_320_v120_
                nw53_[int(6)] = d_320_v120_
                nw53_[int(7)] = d_320_v120_
                nw53_[int(8)] = d_320_v120_
                nw53_[int(9)] = d_320_v120_
                nw53_[int(10)] = d_320_v120_
                nw53_[int(11)] = d_320_v120_
                nw53_[int(12)] = d_320_v120_
                nw53_[int(13)] = d_320_v120_
                nw53_[int(14)] = d_320_v120_
                nw53_[int(15)] = d_320_v120_
                nw53_[int(16)] = d_320_v120_
                nw53_[int(17)] = d_320_v120_
                nw53_[int(18)] = d_320_v120_
                nw53_[int(19)] = d_320_v120_
                nw53_[int(20)] = d_320_v120_
                nw53_[int(21)] = d_320_v120_
                nw53_[int(22)] = d_320_v120_
                nw53_[int(23)] = d_320_v120_
                nw53_[int(24)] = d_320_v120_
                d_326_v125_ = nw53_
                index35_ = default__.safeIndex(587, (d_326_v125_).length(0))
                (d_326_v125_)[index35_] = d_320_v120_
                index36_ = default__.safeIndex(587, (d_326_v125_).length(0))
                (d_326_v125_)[index36_] = d_320_v120_
        elif True:
            (d_145_globalState_).f3 = d_137_v0_
            d_327_v126_: _dafny.Map
            d_327_v126_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_141_v2_, 681]): (d_181_v39_) - (d_181_v39_)})
            d_328_v127_: _dafny.Set
            d_328_v127_ = _dafny.Set({d_141_v2_})
            d_327_v126_ = default__.fm15((len(d_328_v127_)) == (d_141_v2_), (d_142_v3_).set(True, -178), d_137_v0_, False, d_145_globalState_)
            d_329_v128_: _dafny.Map
            d_329_v128_ = _dafny.Map({d_141_v2_: (D5_DC12(d_143_v4_)).cf18})
            d_143_v4_ = (((d_329_v128_)[d_141_v2_] if (d_141_v2_) in (d_329_v128_) else (d_143_v4_).set(default__.safeIndex(d_141_v2_, len(d_143_v4_)), (d_192_v44_).f20))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kr"))), len(((d_329_v128_)[d_141_v2_] if (d_141_v2_) in (d_329_v128_) else (d_143_v4_).set(default__.safeIndex(d_141_v2_, len(d_143_v4_)), (d_192_v44_).f20)))), d_191_v43_)
            (d_145_globalState_).f3 = d_137_v0_
            d_330_v129_: _dafny.Array
            nw54_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_330_v129_ = nw54_
            index37_ = default__.safeIndex(543, (d_330_v129_).length(0))
            (d_330_v129_)[index37_] = d_138_v1_
            index38_ = default__.safeIndex(543, (d_330_v129_).length(0))
            (d_330_v129_)[index38_] = d_138_v1_
        d_331_v130_: _dafny.Set
        d_331_v130_ = _dafny.Set({d_192_v44_, d_192_v44_, d_192_v44_})
        (d_145_globalState_).f3 = (d_192_v44_) in (d_331_v130_)
        d_332_v131_: _dafny.Set
        d_332_v131_ = _dafny.Set({len(d_143_v4_)})
        if (_dafny.Set({d_141_v2_})).isdisjoint(d_332_v131_):
            d_333_v132_: bool
            d_334_v133_: int
            d_335_v134_: bool
            d_336_v135_: bool
            out10_: bool
            out11_: int
            out12_: bool
            out13_: bool
            out10_, out11_, out12_, out13_ = (d_192_v44_).m4(d_145_globalState_)
            d_333_v132_ = out10_
            d_334_v133_ = out11_
            d_335_v134_ = out12_
            d_336_v135_ = out13_
            d_337_v136_: int
            d_337_v136_ = d_334_v133_
            d_338_v137_: C0
            nw55_ = C0()
            nw55_.ctor__(d_337_v136_, d_334_v133_)
            d_338_v137_ = nw55_
            d_339_v138_: _dafny.MultiSet
            d_339_v138_ = _dafny.MultiSet([d_338_v137_, d_338_v137_, d_338_v137_, d_338_v137_])
            d_340_v139_: _dafny.MultiSet
            d_340_v139_ = d_339_v138_
            source5_ = d_339_v138_
            d_341___mcc_h19_ = source5_
            d_342_cf14_ = d_341___mcc_h19_
            d_343_v140_: C2
            nw56_ = C2()
            nw56_.ctor__()
            d_343_v140_ = nw56_
            d_344_v141_: _dafny.Array
            nw57_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_344_v141_ = nw57_
            d_345_v142_: _dafny.Array
            nw58_ = _dafny.Array(None, 6)
            nw58_[int(0)] = d_141_v2_
            nw58_[int(1)] = default__.fm0(d_145_globalState_)
            nw58_[int(2)] = d_141_v2_
            nw58_[int(3)] = d_141_v2_
            nw58_[int(4)] = d_141_v2_
            nw58_[int(5)] = len(d_143_v4_)
            d_345_v142_ = nw58_
            index39_ = default__.safeIndex(718, (d_344_v141_).length(0))
            (d_344_v141_)[index39_] = d_345_v142_
            index40_ = default__.safeIndex(718, (d_344_v141_).length(0))
            (d_344_v141_)[index40_] = d_345_v142_
            (d_145_globalState_).f2 = d_334_v133_
            d_191_v43_ = (d_192_v44_).f20
            rhs62_ = d_138_v1_
            rhs63_ = d_137_v0_
            lhs43_ = d_145_globalState_
            d_138_v1_ = rhs62_
            lhs43_.f10 = rhs63_
            (d_145_globalState_).f2 = d_334_v133_
            d_346_v143_: _dafny.Seq
            d_346_v143_ = _dafny.SeqWithoutIsStrInference([d_335_v134_, d_137_v0_])
            d_347_v144_: _dafny.Map
            d_347_v144_ = _dafny.Map({d_346_v143_: (d_180_v38_).cf2})
            d_348_v145_: _dafny.Map
            d_348_v145_ = _dafny.Map({len(d_347_v144_): ((d_142_v3_)[d_333_v132_] if (d_333_v132_) in (d_142_v3_) else d_334_v133_)})
            d_334_v133_ = ((d_348_v145_)[d_334_v133_] if (d_334_v133_) in (d_348_v145_) else d_334_v133_)
        elif True:
            d_349_v146_: _dafny.Array
            nw59_ = _dafny.Array(None, 22)
            nw59_[int(0)] = d_141_v2_
            nw59_[int(1)] = d_141_v2_
            nw59_[int(2)] = d_141_v2_
            nw59_[int(3)] = d_141_v2_
            nw59_[int(4)] = d_141_v2_
            nw59_[int(5)] = d_141_v2_
            nw59_[int(6)] = d_141_v2_
            nw59_[int(7)] = d_141_v2_
            nw59_[int(8)] = d_141_v2_
            nw59_[int(9)] = d_141_v2_
            nw59_[int(10)] = d_141_v2_
            nw59_[int(11)] = d_141_v2_
            nw59_[int(12)] = d_141_v2_
            nw59_[int(13)] = d_141_v2_
            nw59_[int(14)] = d_141_v2_
            nw59_[int(15)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ify")))
            nw59_[int(16)] = d_141_v2_
            nw59_[int(17)] = d_141_v2_
            nw59_[int(18)] = d_141_v2_
            nw59_[int(19)] = d_141_v2_
            nw59_[int(20)] = d_141_v2_
            nw59_[int(21)] = d_141_v2_
            d_349_v146_ = nw59_
            d_350_v147_: D4
            d_350_v147_ = D4_DC10(d_349_v146_)
            d_351_v148_: D4
            d_351_v148_ = D4_DC11(d_350_v147_)
            pat_let_tv15_ = d_349_v146_
            d_352_v149_: _dafny.Map
            def iife47_(_pat_let16_0):
                def iife48_(d_353_dt__update__tmp_h8_):
                    def iife49_(_pat_let17_0):
                        def iife50_(d_354_dt__update_hcf17_h0_):
                            return D4_DC11(d_354_dt__update_hcf17_h0_)
                        return iife50_(_pat_let17_0)
                    return iife49_(D4_DC10(pat_let_tv15_))
                return iife48_(_pat_let16_0)
            d_352_v149_ = _dafny.Map({d_137_v0_: iife47_(d_351_v148_)})
            pat_let_tv16_ = d_350_v147_
            d_355_v150_: _dafny.Set
            def iife51_(_pat_let18_0):
                def iife52_(d_356_dt__update__tmp_h9_):
                    def iife53_(_pat_let19_0):
                        def iife54_(d_357_dt__update_hcf17_h1_):
                            return D4_DC11(d_357_dt__update_hcf17_h1_)
                        return iife54_(_pat_let19_0)
                    return iife53_(pat_let_tv16_)
                return iife52_(_pat_let18_0)
            d_355_v150_ = _dafny.Set({d_352_v149_, d_352_v149_, _dafny.Map({d_137_v0_: iife51_(d_351_v148_)}), d_352_v149_})
            source6_ = D7_DC18((d_355_v150_) | (_dafny.Set({d_352_v149_, d_352_v149_, d_352_v149_, d_352_v149_})))
            if source6_.is_DC19:
                d_358___mcc_h20_ = source6_.cf38
                d_359___mcc_h21_ = source6_.cf39
                d_360___mcc_h22_ = source6_.cf40
                d_361_cf40_ = d_360___mcc_h22_
                d_362_cf39_ = d_359___mcc_h21_
                d_363_cf38_ = d_358___mcc_h20_
                d_364_v151_: D2
                d_364_v151_ = D2_DC7(D2_DC4(d_183_v41_))
                d_365_v152_: D2
                d_365_v152_ = D2_DC4(d_183_v41_)
                d_364_v151_ = D2_DC7(d_365_v152_)
                (d_145_globalState_).f9 = ((d_141_v2_) + (d_141_v2_) if False else len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mject"))) + (d_143_v4_)))
                d_366_v153_: int
                d_366_v153_ = d_141_v2_
                d_367_v154_: C0
                nw60_ = C0()
                nw60_.ctor__(d_366_v153_, d_141_v2_)
                d_367_v154_ = nw60_
                d_368_v155_: _dafny.Array
                nw61_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
                d_368_v155_ = nw61_
                d_369_v156_: D9
                d_369_v156_ = D9_DC24(d_368_v155_)
                d_370_v157_: D9
                d_370_v157_ = D9_DC24((d_369_v156_).cf49)
                d_371_v158_: _dafny.Seq
                d_371_v158_ = _dafny.SeqWithoutIsStrInference([d_362_cf39_])
                rhs64_ = 373
                rhs65_ = default__.fm1(d_362_cf39_, d_143_v4_, d_145_globalState_)
                rhs66_ = default__.fm0(d_145_globalState_)
                rhs67_ = ((d_369_v156_ if d_362_cf39_ else d_369_v156_) if not((D2_DC5(d_137_v0_, len(d_371_v158_), d_138_v1_, d_177_v35_)).cf5) else d_370_v157_)
                lhs44_ = d_145_globalState_
                lhs45_ = d_145_globalState_
                lhs44_.f2 = rhs64_
                d_362_cf39_ = rhs65_
                lhs45_.f9 = rhs66_
                d_369_v156_ = rhs67_
            elif source6_.is_DC20:
                d_372___mcc_h23_ = source6_.cf41
                d_373___mcc_h24_ = source6_.cf42
                d_374___mcc_h25_ = source6_.cf43
                d_375_cf43_ = d_374___mcc_h25_
                d_376_cf42_ = d_373___mcc_h24_
                d_377_cf41_ = d_372___mcc_h23_
                (d_145_globalState_).f0 = d_137_v0_
                index41_ = default__.safeIndex(870, (d_138_v1_).length(0))
                (d_138_v1_)[index41_] = d_137_v0_
                index42_ = default__.safeIndex(870, (d_138_v1_).length(0))
                rhs68_ = True
                rhs69_ = ((d_375_cf43_) + (d_375_cf43_)) + ((0) - (d_141_v2_))
                lhs46_ = d_138_v1_
                lhs47_ = default__.safeIndex(870, (d_138_v1_).length(0))
                lhs46_[lhs47_] = rhs68_
                d_375_cf43_ = rhs69_
                d_378_v159_: int
                d_378_v159_ = len(_dafny.SeqWithoutIsStrInference([d_375_cf43_ for d_379_i8_ in range(default__.abs(68))]))
                d_380_v160_: _dafny.Map
                d_380_v160_ = _dafny.Map({d_375_cf43_: True})
                d_381_v161_: C0
                nw62_ = C0()
                nw62_.ctor__(d_378_v159_, len(d_380_v160_))
                d_381_v161_ = nw62_
                d_382_v162_: D4
                d_382_v162_ = D4_DC9(d_381_v161_)
                d_383_v163_: _dafny.Map
                d_383_v163_ = _dafny.Map({d_137_v0_: d_381_v161_})
                d_384_v164_: _dafny.Array
                nw63_ = _dafny.Array(None, 10)
                nw63_[int(0)] = d_381_v161_
                nw63_[int(1)] = d_381_v161_
                nw63_[int(2)] = d_381_v161_
                nw63_[int(3)] = (d_382_v162_).cf15
                nw63_[int(4)] = d_381_v161_
                nw63_[int(5)] = d_381_v161_
                nw63_[int(6)] = d_381_v161_
                nw63_[int(7)] = d_381_v161_
                nw63_[int(8)] = d_381_v161_
                nw63_[int(9)] = ((d_383_v163_)[(d_138_v1_)[default__.safeIndex(870, (d_138_v1_).length(0))]] if ((d_138_v1_)[default__.safeIndex(870, (d_138_v1_).length(0))]) in (d_383_v163_) else d_381_v161_)
                d_384_v164_ = nw63_
                index43_ = default__.safeIndex(182, (d_384_v164_).length(0))
                (d_384_v164_)[index43_] = d_381_v161_
                index44_ = default__.safeIndex(182, (d_384_v164_).length(0))
                (d_384_v164_)[index44_] = d_381_v161_
                d_385_v165_: int
                d_386_v166_: int
                out14_: int
                out15_: int
                out14_, out15_ = default__.m0((d_138_v1_)[default__.safeIndex(870, (d_138_v1_).length(0))], d_380_v160_, d_145_globalState_)
                d_385_v165_ = out14_
                d_386_v166_ = out15_
            elif source6_.is_DC18:
                d_387___mcc_h26_ = source6_.cf37
                d_388_cf37_ = d_387___mcc_h26_
                rhs70_ = (d_143_v4_) + ((_dafny.SeqWithoutIsStrInference([(d_192_v44_).f20, (d_192_v44_).f20])) + (d_143_v4_))
                rhs71_ = (d_141_v2_) * ((d_141_v2_) * (d_141_v2_))
                lhs48_ = d_145_globalState_
                d_143_v4_ = rhs70_
                lhs48_.f9 = rhs71_
                index45_ = default__.safeIndex(356, (d_349_v146_).length(0))
                (d_349_v146_)[index45_] = d_141_v2_
                d_389_v167_: _dafny.Array
                nw64_ = _dafny.Array(D10.default()(), 25)
                d_389_v167_ = nw64_
                d_390_v168_: D10
                d_390_v168_ = D10_DC29()
                index46_ = default__.safeIndex(109, (d_389_v167_).length(0))
                (d_389_v167_)[index46_] = d_390_v168_
                index47_ = default__.safeIndex(805, (d_349_v146_).length(0))
                (d_349_v146_)[index47_] = default__.safeDivisionInt(default__.fm0(d_145_globalState_), len(d_332_v131_))
                d_391_v169_: C0
                nw65_ = C0()
                nw65_.ctor__(len(d_179_v37_), d_141_v2_)
                d_391_v169_ = nw65_
                d_392_v170_: D2
                d_392_v170_ = D2_DC6(d_141_v2_, True, d_137_v0_, (0) - (len(_dafny.Map({((d_144_v5_)[d_137_v0_] if (d_137_v0_) in (d_144_v5_) else True): d_391_v169_}))))
                index48_ = default__.safeIndex(356, (d_349_v146_).length(0))
                index49_ = default__.safeIndex(109, (d_389_v167_).length(0))
                index50_ = default__.safeIndex(805, (d_349_v146_).length(0))
                rhs72_ = (d_141_v2_) + (d_141_v2_)
                rhs73_ = d_137_v0_
                rhs74_ = D10_DC29()
                rhs75_ = default__.fm0(d_145_globalState_)
                rhs76_ = (d_392_v170_).cf12
                lhs49_ = d_349_v146_
                lhs50_ = default__.safeIndex(356, (d_349_v146_).length(0))
                lhs51_ = d_145_globalState_
                lhs52_ = d_389_v167_
                lhs53_ = default__.safeIndex(109, (d_389_v167_).length(0))
                lhs54_ = d_145_globalState_
                lhs55_ = d_349_v146_
                lhs56_ = default__.safeIndex(805, (d_349_v146_).length(0))
                lhs49_[lhs50_] = rhs72_
                lhs51_.f3 = rhs73_
                lhs52_[lhs53_] = rhs74_
                lhs54_.f9 = rhs75_
                lhs55_[lhs56_] = rhs76_
                d_393_v171_: _dafny.Map
                d_393_v171_ = _dafny.Map({d_141_v2_: (d_349_v146_)[default__.safeIndex(356, (d_349_v146_).length(0))]})
                d_137_v0_ = (d_137_v0_ if not((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hiwhxnyyd")))) in (d_393_v171_)) else d_137_v0_)
                d_394_v172_: _dafny.Map
                d_394_v172_ = _dafny.Map({default__.safeDivisionInt(d_141_v2_, len(d_143_v4_)): d_137_v0_})
                d_394_v172_ = (d_394_v172_).set(((d_349_v146_)[default__.safeIndex(356, (d_349_v146_).length(0))]) - ((d_349_v146_)[default__.safeIndex(356, (d_349_v146_).length(0))]), (d_137_v0_ if d_137_v0_ else d_137_v0_))
            elif True:
                d_395___mcc_h27_ = source6_.cf44
                d_396_cf44_ = d_395___mcc_h27_
                index51_ = default__.safeIndex(987, (d_138_v1_).length(0))
                (d_138_v1_)[index51_] = (d_141_v2_) != (d_141_v2_)
                index52_ = default__.safeIndex(987, (d_138_v1_).length(0))
                (d_138_v1_)[index52_] = d_137_v0_
                d_397_v173_: int
                d_397_v173_ = (0) - (d_141_v2_)
                d_398_v174_: _dafny.MultiSet
                d_398_v174_ = _dafny.MultiSet([(d_192_v44_).f20, d_191_v43_])
                d_399_v175_: T0
                nw66_ = C0()
                nw66_.ctor__(d_397_v173_, (d_398_v174_).cardinality)
                d_399_v175_ = nw66_
                d_400_v176_: _dafny.Map
                d_400_v176_ = _dafny.Map({d_141_v2_: d_399_v175_})
                d_401_v177_: _dafny.Map
                d_401_v177_ = _dafny.Map({len(_dafny.Set({True})): len(d_400_v176_)})
                rhs77_ = d_349_v146_
                rhs78_ = (d_401_v177_) | (d_401_v177_)
                d_349_v146_ = rhs77_
                d_401_v177_ = rhs78_
                d_191_v43_ = (d_192_v44_).f20
                d_402_v178_: _dafny.Array
                def lambda16_(d_403_i9_):
                    return True

                init9_ = lambda16_
                nw67_ = _dafny.Array(None, 22)
                for i0_9_ in range(nw67_.length(0)):
                    nw67_[i0_9_] = init9_(i0_9_)
                d_402_v178_ = nw67_
                d_404_v179_: D2
                d_404_v179_ = D2_DC5(d_137_v0_, (d_399_v175_).f18, d_402_v178_, d_177_v35_)
                d_405_v180_: _dafny.Set
                d_405_v180_ = _dafny.Set({d_404_v179_})
                d_406_v181_: _dafny.Map
                d_406_v181_ = _dafny.Map({(d_405_v180_) - (d_405_v180_): (0) - ((d_179_v37_)[default__.safeIndex((d_399_v175_).f18, len(d_179_v37_))])})
                d_406_v181_ = (d_406_v181_).set(d_405_v180_, len((default__.fm16((d_399_v175_).f18, (d_138_v1_)[default__.safeIndex(987, (d_138_v1_).length(0))], _dafny.CodePoint('b'), d_141_v2_, d_145_globalState_)) | (d_144_v5_)))
            d_407_v182_: _dafny.Seq
            d_407_v182_ = _dafny.SeqWithoutIsStrInference([d_143_v4_, d_143_v4_])
            d_408_v183_: _dafny.Array
            nw68_ = _dafny.Array(None, 28)
            nw68_[int(0)] = d_143_v4_
            nw68_[int(1)] = default__.fm6(d_137_v0_, d_145_globalState_)
            nw68_[int(2)] = _dafny.SeqWithoutIsStrInference([(d_192_v44_).f20 for d_409_i10_ in range(default__.abs(-743))])
            nw68_[int(3)] = d_143_v4_
            nw68_[int(4)] = d_143_v4_
            nw68_[int(5)] = d_143_v4_
            nw68_[int(6)] = d_143_v4_
            nw68_[int(7)] = default__.fm6(d_137_v0_, d_145_globalState_)
            nw68_[int(8)] = (d_143_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))
            nw68_[int(9)] = (d_143_v4_) + (d_143_v4_)
            nw68_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "veh"))
            nw68_[int(11)] = _dafny.SeqWithoutIsStrInference([d_191_v43_ for d_410_i11_ in range(default__.abs(638))])
            nw68_[int(12)] = d_143_v4_
            nw68_[int(13)] = (d_143_v4_) + (d_143_v4_)
            nw68_[int(14)] = d_143_v4_
            nw68_[int(15)] = d_143_v4_
            nw68_[int(16)] = (d_143_v4_).set(default__.safeIndex(d_141_v2_, len(d_143_v4_)), (d_192_v44_).f20)
            nw68_[int(17)] = _dafny.SeqWithoutIsStrInference([(d_192_v44_).f20 for d_411_i12_ in range(default__.abs(489))])
            nw68_[int(18)] = d_143_v4_
            nw68_[int(19)] = (d_407_v182_)[default__.safeIndex(-360, len(d_407_v182_))]
            nw68_[int(20)] = d_143_v4_
            nw68_[int(21)] = d_143_v4_
            nw68_[int(22)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psmshitr"))
            nw68_[int(23)] = (d_143_v4_).set(default__.safeIndex(d_141_v2_, len(d_143_v4_)), _dafny.CodePoint('f'))
            nw68_[int(24)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_412_i13_ in range(default__.abs(478))])
            nw68_[int(25)] = (d_143_v4_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_413_i14_ in range(default__.abs(-912))]))
            nw68_[int(26)] = d_143_v4_
            nw68_[int(27)] = d_143_v4_
            d_408_v183_ = nw68_
            index53_ = default__.safeIndex(306, (d_408_v183_).length(0))
            (d_408_v183_)[index53_] = d_143_v4_
            index54_ = default__.safeIndex(306, (d_408_v183_).length(0))
            (d_408_v183_)[index54_] = _dafny.SeqWithoutIsStrInference([(d_192_v44_).f20 for d_414_i15_ in range(default__.abs(762))])
            index55_ = default__.safeIndex(379, (d_138_v1_).length(0))
            (d_138_v1_)[index55_] = (d_179_v37_) <= (d_179_v37_)
            index56_ = default__.safeIndex(379, (d_138_v1_).length(0))
            (d_138_v1_)[index56_] = (((d_301_v108_)[default__.safeIndex(481, (d_301_v108_).length(0))]).cf50) == (d_141_v2_)
            d_415_v184_: D11
            d_415_v184_ = D11_DC31(d_181_v39_)
            (d_145_globalState_).f2 = default__.safeDivisionInt((d_141_v2_) * (((d_415_v184_).cf58).cardinality), d_141_v2_)
            d_416_v185_: C1
            nw69_ = C1()
            nw69_.ctor__(d_191_v43_)
            d_416_v185_ = nw69_
        d_417_v186_: _dafny.Array
        def lambda17_(d_418_v131_):
            def lambda18_(d_419_i16_):
                return d_418_v131_

            return lambda18_

        init10_ = lambda17_(d_332_v131_)
        nw70_ = _dafny.Array(None, 27)
        for i0_10_ in range(nw70_.length(0)):
            nw70_[i0_10_] = init10_(i0_10_)
        d_417_v186_ = nw70_
        index57_ = default__.safeIndex(599, (d_417_v186_).length(0))
        (d_417_v186_)[index57_] = (d_332_v131_) - (d_332_v131_)
        index58_ = default__.safeIndex(599, (d_417_v186_).length(0))
        (d_417_v186_)[index58_] = d_332_v131_
        (d_145_globalState_).f0 = ((d_141_v2_) - (d_141_v2_)) >= (d_141_v2_)
        d_142_v3_ = (d_142_v3_) | ((d_142_v3_) | (d_142_v3_))
        d_420_v187_: D6
        d_420_v187_ = D6_DC15(818, d_141_v2_, d_179_v37_, d_141_v2_, _dafny.Set({d_141_v2_}))
        hi1_ = d_141_v2_
        for d_421_i17_ in range(len((d_420_v187_).cf27), hi1_):
            rhs79_ = d_137_v0_
            rhs80_ = d_137_v0_
            rhs81_ = d_421_i17_
            lhs57_ = d_145_globalState_
            lhs58_ = d_145_globalState_
            lhs57_.f0 = rhs79_
            lhs58_.f0 = rhs80_
            d_141_v2_ = rhs81_
            d_422_v188_: _dafny.Array
            def lambda19_(d_423_i17_):
                def lambda20_(d_424_i18_):
                    return (d_424_i18_) * (d_423_i17_)

                return lambda20_

            init11_ = lambda19_(d_421_i17_)
            nw71_ = _dafny.Array(None, 18)
            for i0_11_ in range(nw71_.length(0)):
                nw71_[i0_11_] = init11_(i0_11_)
            d_422_v188_ = nw71_
            index59_ = default__.safeIndex(44, (d_422_v188_).length(0))
            (d_422_v188_)[index59_] = d_141_v2_
            d_425_v189_: D4
            d_425_v189_ = D4_DC10(d_422_v188_)
            pat_let_tv17_ = d_422_v188_
            d_426_v190_: _dafny.Array
            nw72_ = _dafny.Array(None, 13)
            nw72_[int(0)] = d_425_v189_
            nw72_[int(1)] = d_425_v189_
            nw72_[int(2)] = d_425_v189_
            nw72_[int(3)] = d_425_v189_
            nw72_[int(4)] = d_425_v189_
            nw72_[int(5)] = d_425_v189_
            nw72_[int(6)] = d_425_v189_
            nw72_[int(7)] = d_425_v189_
            def iife55_(_pat_let20_0):
                def iife56_(d_427_dt__update__tmp_h10_):
                    def iife57_(_pat_let21_0):
                        def iife58_(d_428_dt__update_hcf16_h0_):
                            return D4_DC10(d_428_dt__update_hcf16_h0_)
                        return iife58_(_pat_let21_0)
                    return iife57_(pat_let_tv17_)
                return iife56_(_pat_let20_0)
            nw72_[int(8)] = iife55_(d_425_v189_)
            nw72_[int(9)] = d_425_v189_
            nw72_[int(10)] = d_425_v189_
            nw72_[int(11)] = D4_DC10(d_422_v188_)
            nw72_[int(12)] = d_425_v189_
            d_426_v190_ = nw72_
            d_429_v191_: _dafny.Map
            d_429_v191_ = _dafny.Map({d_421_i17_: (0) - ((d_420_v187_).cf26)})
            d_430_v192_: _dafny.Set
            d_430_v192_ = _dafny.Set({d_426_v190_, d_426_v190_})
            pat_let_tv18_ = d_430_v192_
            pat_let_tv19_ = d_141_v2_
            index60_ = default__.safeIndex(44, (d_422_v188_).length(0))
            def iife59_(_pat_let22_0):
                def iife60_(d_431_dt__update__tmp_h11_):
                    def iife61_(_pat_let23_0):
                        def iife62_(d_432_dt__update_hcf41_h0_):
                            def iife63_(_pat_let24_0):
                                def iife64_(d_433_dt__update_hcf43_h0_):
                                    return D7_DC20(d_432_dt__update_hcf41_h0_, (d_431_dt__update__tmp_h11_).cf42, d_433_dt__update_hcf43_h0_)
                                return iife64_(_pat_let24_0)
                            return iife63_(pat_let_tv19_)
                        return iife62_(_pat_let23_0)
                    return iife61_(pat_let_tv18_)
                return iife60_(_pat_let22_0)
            (d_422_v188_)[index60_] = (iife59_(D7_DC20(_dafny.Set({d_426_v190_, d_426_v190_}), D5_DC13(default__.fm0(d_145_globalState_), default__.fm0(d_145_globalState_), d_137_v0_), len(d_429_v191_)))).cf43
            d_434_v193_: int
            d_434_v193_ = d_421_i17_
            d_435_v194_: T0
            nw73_ = C0()
            nw73_.ctor__(d_434_v193_, (d_141_v2_) - ((0) - (d_141_v2_)))
            d_435_v194_ = nw73_
            d_436_v195_: _dafny.Map
            d_436_v195_ = _dafny.Map({default__.fm1(d_137_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdygc")), d_145_globalState_): d_435_v194_})
            d_435_v194_ = ((d_436_v195_)[d_137_v0_] if (d_137_v0_) in (d_436_v195_) else d_435_v194_)
            d_437_v196_: D11
            d_437_v196_ = D11_DC31(d_181_v39_)
            d_438_v197_: D11
            d_438_v197_ = D11_DC33(d_437_v196_)
            d_439_v198_: D11
            d_439_v198_ = D11_DC33(d_438_v197_)
            d_439_v198_ = d_439_v198_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_138_v1_).length(0)):
            d_440_i19_: int = guard_loop_0_
            if (True) and (((0) <= (d_440_i19_)) and ((d_440_i19_) < ((d_138_v1_).length(0)))):
                (d_138_v1_)[(d_440_i19_)] = (d_141_v2_) != ((len(_dafny.Map({d_141_v2_: False}))) + (len(d_143_v4_)))
        d_441_v199_: _dafny.Set
        d_441_v199_ = _dafny.Set({d_143_v4_, d_143_v4_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))})
        d_441_v199_ = d_441_v199_
        hi2_ = d_141_v2_
        for d_442_i20_ in range(default__.safeModuloInt(d_141_v2_, d_141_v2_), hi2_):
            (d_145_globalState_).f9 = d_141_v2_
            if not(d_137_v0_):
                d_443_v200_: C3
                nw74_ = C3()
                nw74_.ctor__(d_143_v4_, 633, d_141_v2_)
                d_443_v200_ = nw74_
                d_444_v201_: int
                d_444_v201_ = d_442_i20_
                d_445_v202_: C0
                nw75_ = C0()
                nw75_.ctor__(d_444_v201_, (d_443_v200_).fm3(d_137_v0_, d_442_i20_, d_145_globalState_))
                d_445_v202_ = nw75_
                d_446_v203_: _dafny.MultiSet
                d_446_v203_ = _dafny.MultiSet([d_445_v202_, d_445_v202_, d_445_v202_, d_445_v202_])
                d_447_v204_: _dafny.Array
                nw76_ = _dafny.Array(None, 11)
                nw76_[int(0)] = d_141_v2_
                nw76_[int(1)] = d_442_i20_
                nw76_[int(2)] = d_442_i20_
                nw76_[int(3)] = (d_446_v203_).cardinality
                nw76_[int(4)] = (d_443_v200_).fm3(d_137_v0_, d_141_v2_, d_145_globalState_)
                nw76_[int(5)] = d_141_v2_
                nw76_[int(6)] = d_442_i20_
                nw76_[int(7)] = 834
                nw76_[int(8)] = d_141_v2_
                nw76_[int(9)] = d_141_v2_
                nw76_[int(10)] = -510
                d_447_v204_ = nw76_
                d_448_v205_: D4
                d_448_v205_ = D4_DC10(d_447_v204_)
                d_448_v205_ = D4_DC10(d_447_v204_)
                d_449_v206_: _dafny.MultiSet
                d_449_v206_ = _dafny.MultiSet([d_138_v1_, d_138_v1_, d_138_v1_, d_138_v1_, d_138_v1_])
                d_449_v206_ = d_449_v206_
                d_450_v207_: _dafny.Array
                d_451_v208_: bool
                out16_: _dafny.Array
                out17_: bool
                out16_, out17_ = (d_443_v200_).m1(d_145_globalState_)
                d_450_v207_ = out16_
                d_451_v208_ = out17_
                d_452_v209_: _dafny.Map
                d_452_v209_ = _dafny.Map({d_138_v1_: d_182_v40_})
                rhs82_ = d_450_v207_
                rhs83_ = len(((d_452_v209_) | (d_452_v209_)).set(d_138_v1_, (d_182_v40_) | (d_182_v40_)))
                rhs84_ = d_443_v200_
                d_450_v207_ = rhs82_
                d_141_v2_ = rhs83_
                d_443_v200_ = rhs84_
            elif True:
                d_453_v210_: C2
                nw77_ = C2()
                nw77_.ctor__()
                d_453_v210_ = nw77_
                d_454_v211_: _dafny.MultiSet
                d_454_v211_ = _dafny.MultiSet([d_453_v210_])
                d_455_v212_: _dafny.Set
                d_455_v212_ = _dafny.Set({d_454_v211_, d_454_v211_})
                (d_145_globalState_).f0 = ((d_455_v212_) - (d_455_v212_)).ispropersubset(d_455_v212_)
                d_456_v213_: _dafny.Map
                d_456_v213_ = _dafny.Map({d_453_v210_: False})
                d_457_v215_: _dafny.Map
                d_457_v215_ = _dafny.Map({(0) - (d_141_v2_): not(d_137_v0_)})
                d_458_v216_: _dafny.Set
                d_458_v216_ = _dafny.Set({d_457_v215_, d_457_v215_, d_457_v215_, _dafny.Map({d_141_v2_: d_137_v0_}), d_457_v215_})
                d_459_v217_: _dafny.Map
                def iife65_():
                    coll15_ = _dafny.Map()
                    compr_15_: _dafny.Map
                    for compr_15_ in (d_458_v216_).Elements:
                        d_460_v214_: _dafny.Map = compr_15_
                        if (d_460_v214_) in (d_458_v216_):
                            coll15_[d_460_v214_] = d_141_v2_
                    return _dafny.Map(coll15_)
                d_459_v217_ = _dafny.Map({((d_456_v213_)[d_453_v210_] if (d_453_v210_) in (d_456_v213_) else d_137_v0_): D10_DC28(iife65_()
)})
                d_461_v218_: _dafny.Map
                d_461_v218_ = _dafny.Map({d_137_v0_: d_459_v217_})
                d_461_v218_ = d_461_v218_
                d_462_v219_: int
                d_462_v219_ = d_442_i20_
                d_463_v220_: T0
                nw78_ = C0()
                nw78_.ctor__(d_462_v219_, d_442_i20_)
                d_463_v220_ = nw78_
                nw79_ = C0()
                nw79_.ctor__(d_462_v219_, (d_463_v220_).f18)
                d_463_v220_ = nw79_
                (d_145_globalState_).f2 = (default__.safeDivisionInt(d_141_v2_, d_141_v2_) if (default__.fm6(d_137_v0_, d_145_globalState_)) != (default__.fm6((d_180_v38_).cf2, d_145_globalState_)) else (d_463_v220_).f18)
                d_457_v215_ = (d_457_v215_).set((925 if d_137_v0_ else len(_dafny.SeqWithoutIsStrInference([d_137_v0_, d_137_v0_, d_137_v0_]))), d_137_v0_)
            (d_145_globalState_).f3 = (default__.fm11(d_442_i20_, d_137_v0_, d_442_i20_, d_145_globalState_)) < (_dafny.SeqWithoutIsStrInference([d_442_i20_, d_442_i20_]))
            d_464_v221_: _dafny.Array
            nw80_ = _dafny.Array(_dafny.CodePoint('D'), 14)
            d_464_v221_ = nw80_
            d_465_v222_: _dafny.Set
            d_465_v222_ = _dafny.Set({d_464_v221_, d_464_v221_})
            d_466_v223_: D7
            d_466_v223_ = D7_DC19(d_465_v222_, d_137_v0_, d_137_v0_)
            d_467_v224_: D7
            d_467_v224_ = D7_DC21(d_466_v223_)
            d_467_v224_ = d_467_v224_
        _dafny.print(_dafny.string_of(d_137_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v1_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_141_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_v3_) == (_dafny.Map({True: 213}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_143_v4_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_144_v5_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_145_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f1) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_145_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_145_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f7)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f8) == (_dafny.MultiSet([213, 213]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_145_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_145_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f11)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f12) == (_dafny.Map({True: -213}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_145_globalState_).f13).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f15) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_v35_) == (_dafny.MultiSet([212, 212, 212]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_v36_).cf28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_v36_).cf29))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_v36_).cf30))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_v36_).cf31))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_v36_).cf32))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_179_v37_) == (_dafny.SeqWithoutIsStrInference([-496, 212, 212, 212, -496, 212, 212]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v38_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_v39_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_182_v40_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[0]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[1]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[2]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[3]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[4]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[5]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[6]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[7]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[8]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[9]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[10]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[11]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[12]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[13]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[14]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[15]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[16]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[17]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[18]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v41_)[19]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[0]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[1]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[2]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[3]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[4]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[5]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[6]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[7]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[8]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[9]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[10]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[11]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[12]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[13]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[14]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[15]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[16]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[17]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[18]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_190_v42_).cf4)[19]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_191_v43_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_192_v44_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_193_v45_).cf22).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_300_v107_).cf50))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v108_)[0]).cf50))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v108_)[1]).cf50))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v108_)[2]).cf50))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v108_)[3]).cf50))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v108_)[4]).cf50))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v108_)[5]).cf50))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_331_v130_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_332_v131_) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[0]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[1]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[2]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[3]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[4]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[5]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[6]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[7]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[8]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[9]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[10]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[11]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[12]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[13]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[14]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[15]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[16]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[17]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[18]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[19]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[20]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[21]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[22]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[23]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[24]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[25]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v186_)[26]) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_420_v187_).cf23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_420_v187_).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_420_v187_).cf25) == (_dafny.SeqWithoutIsStrInference([-496, 212, 212, 212, -496, 212, 212]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_420_v187_).cf26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_420_v187_).cf27) == (_dafny.Set({212}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_441_v199_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))}))))
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
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

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

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC5(False, int(0), _dafny.Array(None, 0), _dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC5(D2, NamedTuple('DC5', [('cf5', Any), ('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC8(D3, NamedTuple('DC8', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC10(D4, NamedTuple('DC10', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC13(D5, NamedTuple('DC13', [('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({self.cf18.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC15(int(0), int(0), _dafny.Seq({}), int(0), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC15(D6, NamedTuple('DC15', [('cf23', Any), ('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf28', Any), ('cf29', Any), ('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf33', Any), ('cf34', Any), ('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC14(D6, NamedTuple('DC14', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC19(_dafny.Set({}), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC19(D7, NamedTuple('DC19', [('cf38', Any), ('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf41', Any), ('cf42', Any), ('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42 and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23(int(0), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC23(D8, NamedTuple('DC23', [('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC25(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)

class D9_DC25(D9, NamedTuple('DC25', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf51', Any), ('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC29()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D10_DC30)

class D10_DC29(D10, NamedTuple('DC29', [])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29)
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC30(D10, NamedTuple('DC30', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC30({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC30) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC32(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)

class D11_DC32(D11, NamedTuple('DC32', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC31(D11, NamedTuple('DC31', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC33(D11, NamedTuple('DC33', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)

class D12_DC34(D12, NamedTuple('DC34', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self.f2: int = int(0)
        self.f3: bool = False
        self.f8: _dafny.MultiSet = _dafny.MultiSet({})
        self.f9: int = int(0)
        self.f10: bool = False
        self.f11: _dafny.Array = _dafny.Array(None, 0)
        self._f1: _dafny.Seq = _dafny.Seq({})
        self._f4: bool = False
        self._f5: str = _dafny.CodePoint('D')
        self._f6: int = int(0)
        self._f7: _dafny.Array = _dafny.Array(None, 0)
        self._f12: _dafny.Map = _dafny.Map({})
        self._f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f14: bool = False
        self._f15: _dafny.Map = _dafny.Map({})
        self._f16: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16):
        (self).f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self).f9 = f9
        (self).f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16

    @property
    def f1(self):
        return self._f1
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

class C0(T0):
    def  __init__(self):
        self._f17: int = int(0)
        self._f18: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f17, f18):
        (self)._f17 = f17
        (self)._f18 = f18

    def fm2(self, globalState):
        source7_ = (self).f17
        d_468___mcc_h0_ = source7_
        d_469_cf0_ = d_468___mcc_h0_
        return (self).f17


class C1:
    def  __init__(self):
        self._f20: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f20):
        (self)._f20 = f20

    def m3(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        d_470_i0_: int
        d_470_i0_ = 0
        with _dafny.label("0"):
            while p0:
                with _dafny.c_label("0"):
                    if (d_470_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_470_i0_ = (d_470_i0_) + (1)
                    d_471_v0_: _dafny.Array
                    nw81_ = _dafny.Array(False, 27)
                    d_471_v0_ = nw81_
                    d_472_v1_: int
                    d_472_v1_ = 990
                    d_473_v2_: _dafny.Map
                    d_473_v2_ = _dafny.Map({(d_471_v0_ if p0 else d_471_v0_): d_472_v1_})
                    d_474_v3_: _dafny.Seq
                    d_474_v3_ = _dafny.SeqWithoutIsStrInference([(default__.fm5((0) - (d_472_v1_), 121, globalState)).cf2])
                    (globalState).f2 = ((d_473_v2_)[d_471_v0_] if (d_471_v0_) in (d_473_v2_) else (0) - (len(d_474_v3_)))
                    if p0:
                        d_475_v4_: _dafny.Map
                        d_475_v4_ = _dafny.Map({d_472_v1_: (d_472_v1_) - (d_472_v1_)})
                        d_475_v4_ = (d_475_v4_).set(d_472_v1_, (304) - ((0) - (d_472_v1_)))
                        d_476_v5_: bool
                        d_477_v6_: int
                        d_478_v7_: bool
                        d_479_v8_: bool
                        out18_: bool
                        out19_: int
                        out20_: bool
                        out21_: bool
                        out18_, out19_, out20_, out21_ = (self).m4(globalState)
                        d_476_v5_ = out18_
                        d_477_v6_ = out19_
                        d_478_v7_ = out20_
                        d_479_v8_ = out21_
                        def iife66_():
                            coll16_ = _dafny.Set()
                            compr_16_: int
                            for compr_16_ in _dafny.IntegerRange(959, 519):
                                d_480_v9_: int = compr_16_
                                if ((959) <= (d_480_v9_)) and ((d_480_v9_) < (519)):
                                    coll16_ = coll16_.union(_dafny.Set([(d_480_v9_) + (d_477_v6_)]))
                            return _dafny.Set(coll16_)
                        (globalState).f10 = (d_472_v1_) not in (iife66_()
                        )
                        d_472_v1_ = d_477_v6_
                        d_481_v10_: D1
                        d_481_v10_ = D1_DC1(d_471_v0_)
                        d_482_v11_: D1
                        d_482_v11_ = D1_DC3(d_481_v10_)
                        d_483_v12_: D1
                        d_483_v12_ = D1_DC1(d_471_v0_)
                        d_484_v13_: _dafny.Array
                        nw82_ = _dafny.Array(None, 15)
                        nw82_[int(0)] = d_483_v12_
                        nw82_[int(1)] = d_483_v12_
                        nw82_[int(2)] = d_483_v12_
                        nw82_[int(3)] = d_483_v12_
                        nw82_[int(4)] = d_483_v12_
                        nw82_[int(5)] = d_483_v12_
                        nw82_[int(6)] = d_483_v12_
                        nw82_[int(7)] = D1_DC1(d_471_v0_)
                        nw82_[int(8)] = d_483_v12_
                        nw82_[int(9)] = d_483_v12_
                        nw82_[int(10)] = d_483_v12_
                        nw82_[int(11)] = d_483_v12_
                        nw82_[int(12)] = d_483_v12_
                        nw82_[int(13)] = d_483_v12_
                        nw82_[int(14)] = d_483_v12_
                        d_484_v13_ = nw82_
                        d_485_v14_: _dafny.Seq
                        d_485_v14_ = _dafny.SeqWithoutIsStrInference([d_484_v13_, d_484_v13_])
                        d_486_v15_: _dafny.Array
                        nw83_ = _dafny.Array(None, 19)
                        nw83_[int(0)] = d_484_v13_
                        nw83_[int(1)] = d_484_v13_
                        nw83_[int(2)] = d_484_v13_
                        nw83_[int(3)] = d_484_v13_
                        nw83_[int(4)] = (d_484_v13_ if p0 else d_484_v13_)
                        nw83_[int(5)] = d_484_v13_
                        nw83_[int(6)] = d_484_v13_
                        nw83_[int(7)] = d_484_v13_
                        nw83_[int(8)] = d_484_v13_
                        nw83_[int(9)] = (d_484_v13_ if d_478_v7_ else d_484_v13_)
                        nw83_[int(10)] = d_484_v13_
                        nw83_[int(11)] = d_484_v13_
                        nw83_[int(12)] = d_484_v13_
                        nw83_[int(13)] = d_484_v13_
                        nw83_[int(14)] = (d_484_v13_ if p0 else d_484_v13_)
                        nw83_[int(15)] = d_484_v13_
                        nw83_[int(16)] = d_484_v13_
                        nw83_[int(17)] = d_484_v13_
                        nw83_[int(18)] = (d_485_v14_)[default__.safeIndex(d_472_v1_, len(d_485_v14_))]
                        d_486_v15_ = nw83_
                        d_487_v16_: _dafny.Map
                        d_487_v16_ = _dafny.Map({d_477_v6_: d_476_v5_})
                        d_488_v17_: _dafny.Seq
                        d_488_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
                        d_489_v18_: _dafny.Set
                        d_489_v18_ = _dafny.Set({(self).f20})
                        d_490_v19_: _dafny.Array
                        nw84_ = _dafny.Array(None, 12)
                        nw84_[int(0)] = d_472_v1_
                        nw84_[int(1)] = d_477_v6_
                        nw84_[int(2)] = d_477_v6_
                        nw84_[int(3)] = d_472_v1_
                        nw84_[int(4)] = 840
                        nw84_[int(5)] = (0) - ((0) - (d_472_v1_))
                        nw84_[int(6)] = len(d_475_v4_)
                        nw84_[int(7)] = len(d_488_v17_)
                        nw84_[int(8)] = len(d_489_v18_)
                        nw84_[int(9)] = d_477_v6_
                        nw84_[int(10)] = d_477_v6_
                        nw84_[int(11)] = default__.fm0(globalState)
                        d_490_v19_ = nw84_
                        d_491_v20_: _dafny.Map
                        d_491_v20_ = _dafny.Map({((d_487_v16_)[d_477_v6_] if (d_477_v6_) in (d_487_v16_) else d_476_v5_): d_490_v19_})
                        rhs85_ = d_482_v11_
                        rhs86_ = ((d_477_v6_) != (d_472_v1_)) in (d_491_v20_)
                        rhs87_ = d_486_v15_
                        rhs88_ = (d_477_v6_) == (default__.safeDivisionInt((0) - (d_472_v1_), d_472_v1_))
                        rhs89_ = d_471_v0_
                        d_482_v11_ = rhs85_
                        d_478_v7_ = rhs86_
                        d_486_v15_ = rhs87_
                        d_476_v5_ = rhs88_
                        d_471_v0_ = rhs89_
                    elif True:
                        d_492_v21_: _dafny.Seq
                        d_492_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcfqu"))
                        rhs90_ = d_492_v21_
                        rhs91_ = ((d_472_v1_) - (d_472_v1_) if p0 else default__.fm0(globalState))
                        lhs59_ = globalState
                        d_492_v21_ = rhs90_
                        lhs59_.f2 = rhs91_
                        d_493_v22_: _dafny.Map
                        d_493_v22_ = _dafny.Map({d_471_v0_: p0})
                        (globalState).f0 = ((d_493_v22_) | (d_493_v22_)) != (d_493_v22_)
                        d_494_v23_: int
                        d_494_v23_ = (0) - (default__.fm0(globalState))
                        d_495_v24_: C0
                        nw85_ = C0()
                        nw85_.ctor__(d_494_v23_, d_472_v1_)
                        d_495_v24_ = nw85_
                        d_496_v25_: C0
                        nw86_ = C0()
                        nw86_.ctor__(d_494_v23_, d_472_v1_)
                        d_496_v25_ = nw86_
                        (globalState).f2 = (0) - ((0) - (d_472_v1_))
                    if (default__.safeModuloInt(-85, (0) - (-525))) < (d_472_v1_):
                        (globalState).f10 = (d_474_v3_)[default__.safeIndex(default__.safeDivisionInt(d_472_v1_, (0) - (d_472_v1_)), len(d_474_v3_))]
                        index61_ = default__.safeIndex(733, (d_471_v0_).length(0))
                        (d_471_v0_)[index61_] = p0
                        d_497_v26_: _dafny.Seq
                        d_497_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
                        d_498_v27_: _dafny.Map
                        d_498_v27_ = _dafny.Map({len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nogbhr"))).set(default__.safeIndex(d_472_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nogbhr")))), (self).f20)) + ((default__.fm6(default__.fm1(p0, d_497_v26_, globalState), globalState)).set(default__.safeIndex(d_472_v1_, len(default__.fm6(default__.fm1(p0, d_497_v26_, globalState), globalState))), default__.fm7((self).f20, p0, globalState)))): (_dafny.Set({d_472_v1_})).isdisjoint(default__.fm8((_dafny.MultiSet(d_497_v26_)).cardinality, d_472_v1_, globalState))})
                        index62_ = default__.safeIndex(336, (d_471_v0_).length(0))
                        (d_471_v0_)[index62_] = p0
                        d_499_v28_: _dafny.Map
                        d_499_v28_ = _dafny.Map({p0: p0})
                        d_500_v29_: _dafny.Map
                        d_500_v29_ = _dafny.Map({d_499_v28_: (0) - (d_472_v1_)})
                        d_501_v30_: _dafny.MultiSet
                        d_501_v30_ = _dafny.MultiSet([d_472_v1_])
                        d_502_v31_: _dafny.Set
                        d_502_v31_ = _dafny.Set({d_472_v1_})
                        d_503_v32_: _dafny.Seq
                        d_503_v32_ = _dafny.SeqWithoutIsStrInference([len(d_502_v31_), d_472_v1_, len(d_474_v3_), d_472_v1_, d_472_v1_])
                        d_504_v34_: _dafny.Map
                        def iife67_():
                            coll17_ = _dafny.Map()
                            compr_17_: int
                            for compr_17_ in (d_498_v27_).keys.Elements:
                                d_505_v33_: int = compr_17_
                                if (d_505_v33_) in (d_498_v27_):
                                    coll17_[(d_505_v33_) * (214)] = p0
                            return _dafny.Map(coll17_)
                        d_504_v34_ = _dafny.Map({((d_501_v30_)[d_472_v1_] if (d_472_v1_) in (d_501_v30_) else len(d_503_v32_)): iife67_()
                        })
                        d_506_v35_: _dafny.Set
                        d_506_v35_ = _dafny.Set({((d_499_v28_)[p0] if (p0) in (d_499_v28_) else p0), p0, p0})
                        d_507_v36_: D1
                        d_507_v36_ = D1_DC2(p0)
                        index63_ = default__.safeIndex(733, (d_471_v0_).length(0))
                        index64_ = default__.safeIndex(336, (d_471_v0_).length(0))
                        rhs92_ = (d_499_v28_) not in (d_500_v29_)
                        rhs93_ = (d_498_v27_ if p0 else ((d_504_v34_)[d_472_v1_] if (d_472_v1_) in (d_504_v34_) else d_498_v27_))
                        rhs94_ = p0
                        rhs95_ = (p0) not in ((d_506_v35_) | (_dafny.Set({(d_507_v36_).cf2, p0, p0})))
                        rhs96_ = default__.fm0(globalState)
                        lhs60_ = d_471_v0_
                        lhs61_ = default__.safeIndex(733, (d_471_v0_).length(0))
                        lhs62_ = d_471_v0_
                        lhs63_ = default__.safeIndex(336, (d_471_v0_).length(0))
                        lhs64_ = globalState
                        lhs65_ = globalState
                        lhs60_[lhs61_] = rhs92_
                        d_498_v27_ = rhs93_
                        lhs62_[lhs63_] = rhs94_
                        lhs64_.f0 = rhs95_
                        lhs65_.f9 = rhs96_
                        (globalState).f0 = (d_507_v36_).cf2
                        d_508_v37_: _dafny.Map
                        d_508_v37_ = _dafny.Map({(self).f20: p0})
                        d_508_v37_ = (d_508_v37_).set((self).f20, p0)
                        d_509_v38_: _dafny.Set
                        d_509_v38_ = _dafny.Set({(d_498_v27_) | (d_498_v27_)})
                        d_509_v38_ = d_509_v38_
                    elif True:
                        d_510_v39_: int
                        d_510_v39_ = d_472_v1_
                        d_511_v40_: C0
                        nw87_ = C0()
                        nw87_.ctor__(d_510_v39_, d_472_v1_)
                        d_511_v40_ = nw87_
                        d_511_v40_ = d_511_v40_
                        (globalState).f0 = p0
                        (globalState).f9 = d_472_v1_
                        (globalState).f2 = d_472_v1_
                        d_512_v41_: D1
                        d_512_v41_ = D1_DC1(d_471_v0_)
                        d_513_v42_: _dafny.Map
                        d_513_v42_ = _dafny.Map({(self).f20: d_512_v41_})
                        d_514_v43_: _dafny.Seq
                        d_514_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "im"))
                        d_515_v44_: _dafny.Array
                        nw88_ = _dafny.Array(None, 2)
                        nw88_[int(0)] = d_514_v43_
                        nw88_[int(1)] = d_514_v43_
                        d_515_v44_ = nw88_
                        index65_ = default__.safeIndex(75, (d_515_v44_).length(0))
                        (d_515_v44_)[index65_] = d_514_v43_
                        index66_ = default__.safeIndex(75, (d_515_v44_).length(0))
                        rhs97_ = p0
                        rhs98_ = default__.safeDivisionInt(d_472_v1_, 256)
                        rhs99_ = _dafny.Map({(self).f20: d_512_v41_})
                        rhs100_ = default__.fm6(p0, globalState)
                        lhs66_ = globalState
                        lhs67_ = globalState
                        lhs68_ = d_515_v44_
                        lhs69_ = default__.safeIndex(75, (d_515_v44_).length(0))
                        lhs66_.f3 = rhs97_
                        lhs67_.f9 = rhs98_
                        d_513_v42_ = rhs99_
                        lhs68_[lhs69_] = rhs100_
                    d_516_v45_: int
                    d_516_v45_ = d_472_v1_
                    d_517_v46_: C0
                    nw89_ = C0()
                    nw89_.ctor__(d_516_v45_, d_472_v1_)
                    d_517_v46_ = nw89_
                    pass
            pass
        (globalState).f3 = p0
        d_518_v47_: int
        d_518_v47_ = 825
        r0 = default__.safeDivisionInt((0) - (d_518_v47_), 127)
        d_519_i1_: int
        d_519_i1_ = 0
        with _dafny.label("1"):
            while not(p0):
                with _dafny.c_label("1"):
                    if (d_519_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_519_i1_ = (d_519_i1_) + (1)
                    d_520_v48_: _dafny.Map
                    d_520_v48_ = _dafny.Map({p0: p0})
                    d_520_v48_ = (d_520_v48_) | (d_520_v48_)
                    (globalState).f10 = (not(p0)) and (((0) - (-389)) > (len(_dafny.SeqWithoutIsStrInference([p0, p0, p0]))))
                    d_521_v49_: _dafny.Seq
                    d_521_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tht"))
                    d_521_v49_ = (d_521_v49_) + (d_521_v49_)
                    d_522_v50_: int
                    d_522_v50_ = d_518_v47_
                    d_523_v51_: _dafny.Map
                    d_523_v51_ = _dafny.Map({(d_522_v50_): default__.fm1(p0, d_521_v49_, globalState)})
                    d_524_v52_: int
                    d_525_v53_: int
                    out22_: int
                    out23_: int
                    out22_, out23_ = default__.m0(p0, d_523_v51_, globalState)
                    d_524_v52_ = out22_
                    d_525_v53_ = out23_
                    pass
            pass
        d_526_v54_: _dafny.Array
        nw90_ = _dafny.Array(False, 26)
        d_526_v54_ = nw90_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_526_v54_).length(0)):
            d_527_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_527_i2_)) and ((d_527_i2_) < ((d_526_v54_).length(0)))):
                (d_526_v54_)[(d_527_i2_)] = not(p0)
        (globalState).f2 = d_518_v47_
        d_528_v55_: _dafny.Seq
        d_528_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iakx"))
        r0 = (len(d_528_v55_) if p0 else d_518_v47_)
        d_529_v56_: _dafny.MultiSet
        d_529_v56_ = _dafny.MultiSet([p0])
        d_530_v57_: _dafny.Set
        d_530_v57_ = _dafny.Set({d_518_v47_})
        d_531_v58_: _dafny.Seq
        d_531_v58_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, default__.fm1(p0, d_528_v55_, globalState), p0])
        d_532_v59_: _dafny.MultiSet
        d_532_v59_ = _dafny.MultiSet([d_529_v56_, d_529_v56_, (default__.fm9(len(d_530_v57_), p0, p0, (0) - ((0) - (len(d_531_v58_))), globalState) if p0 else _dafny.MultiSet(d_531_v58_)), (d_529_v56_).set(p0, default__.abs(len(d_530_v57_)))])
        r1 = d_532_v59_
        return r0, r1

    def m4(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        d_533_v0_: bool
        d_533_v0_ = False
        if d_533_v0_:
            d_534_v1_: int
            d_534_v1_ = 785
            d_535_v2_: _dafny.Seq
            d_535_v2_ = _dafny.SeqWithoutIsStrInference([d_533_v0_])
            d_536_v3_: _dafny.MultiSet
            d_536_v3_ = _dafny.MultiSet([d_533_v0_, (d_534_v1_) == (d_534_v1_), (d_535_v2_)[default__.safeIndex(d_534_v1_, len(d_535_v2_))]])
            d_537_v4_: _dafny.Seq
            d_537_v4_ = _dafny.SeqWithoutIsStrInference([d_534_v1_, (0) - (d_534_v1_)])
            d_538_v5_: _dafny.Seq
            d_538_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfyg"))
            rhs101_ = (_dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([d_533_v0_, d_533_v0_]) if d_533_v0_ else d_535_v2_)).set(default__.safeIndex(947, len((_dafny.SeqWithoutIsStrInference([d_533_v0_, d_533_v0_]) if d_533_v0_ else d_535_v2_))), d_533_v0_))) | (d_536_v3_)
            rhs102_ = d_533_v0_
            rhs103_ = (default__.safeDivisionInt(d_534_v1_, len(d_537_v4_))) - ((0) - ((d_534_v1_) + ((0) - (len(d_538_v5_)))))
            lhs70_ = globalState
            lhs71_ = globalState
            d_536_v3_ = rhs101_
            lhs70_.f3 = rhs102_
            lhs71_.f2 = rhs103_
            d_539_v6_: _dafny.Array
            nw91_ = _dafny.Array(D1.default()(), 16)
            d_539_v6_ = nw91_
            d_540_v7_: D2
            d_540_v7_ = D2_DC4(d_539_v6_)
            d_539_v6_ = (d_540_v7_).cf4
            d_541_v8_: _dafny.Array
            nw92_ = _dafny.Array(int(0), 27)
            d_541_v8_ = nw92_
            index67_ = default__.safeIndex(644, (d_541_v8_).length(0))
            (d_541_v8_)[index67_] = (d_534_v1_) - (d_534_v1_)
            index68_ = default__.safeIndex(644, (d_541_v8_).length(0))
            (d_541_v8_)[index68_] = d_534_v1_
            (globalState).f0 = False
            d_542_v9_: _dafny.Array
            nw93_ = _dafny.Array(_dafny.Seq({}), 23)
            d_542_v9_ = nw93_
            d_543_v10_: _dafny.Seq
            d_543_v10_ = _dafny.SeqWithoutIsStrInference([d_542_v9_, d_542_v9_])
            d_544_v11_: _dafny.Seq
            d_544_v11_ = _dafny.SeqWithoutIsStrInference([d_542_v9_, (d_543_v10_)[default__.safeIndex(-315, len(d_543_v10_))]])
            d_545_v12_: _dafny.Map
            d_545_v12_ = _dafny.Map({d_533_v0_: d_542_v9_})
            d_546_v13_: _dafny.Array
            nw94_ = _dafny.Array(None, 28)
            nw94_[int(0)] = d_542_v9_
            nw94_[int(1)] = d_542_v9_
            nw94_[int(2)] = d_542_v9_
            nw94_[int(3)] = d_542_v9_
            nw94_[int(4)] = d_542_v9_
            nw94_[int(5)] = d_542_v9_
            nw94_[int(6)] = d_542_v9_
            nw94_[int(7)] = d_542_v9_
            nw94_[int(8)] = d_542_v9_
            nw94_[int(9)] = d_542_v9_
            nw94_[int(10)] = d_542_v9_
            nw94_[int(11)] = (d_544_v11_)[default__.safeIndex(d_534_v1_, len(d_544_v11_))]
            nw94_[int(12)] = d_542_v9_
            nw94_[int(13)] = d_542_v9_
            nw94_[int(14)] = (d_542_v9_ if True else d_542_v9_)
            nw94_[int(15)] = d_542_v9_
            nw94_[int(16)] = d_542_v9_
            nw94_[int(17)] = d_542_v9_
            nw94_[int(18)] = d_542_v9_
            nw94_[int(19)] = d_542_v9_
            nw94_[int(20)] = ((d_544_v11_)[default__.safeIndex(d_534_v1_, len(d_544_v11_))] if True else d_542_v9_)
            nw94_[int(21)] = d_542_v9_
            nw94_[int(22)] = d_542_v9_
            nw94_[int(23)] = d_542_v9_
            nw94_[int(24)] = d_542_v9_
            nw94_[int(25)] = d_542_v9_
            nw94_[int(26)] = ((d_545_v12_)[d_533_v0_] if (d_533_v0_) in (d_545_v12_) else d_542_v9_)
            nw94_[int(27)] = d_542_v9_
            d_546_v13_ = nw94_
            index69_ = default__.safeIndex(554, (d_546_v13_).length(0))
            (d_546_v13_)[index69_] = d_542_v9_
            index70_ = default__.safeIndex(554, (d_546_v13_).length(0))
            (d_546_v13_)[index70_] = d_542_v9_
        elif True:
            d_547_v14_: int
            d_547_v14_ = -271
            d_548_v15_: _dafny.Seq
            d_548_v15_ = _dafny.SeqWithoutIsStrInference([d_547_v14_, d_547_v14_, d_547_v14_])
            (globalState).f0 = (default__.safeModuloInt(d_547_v14_, d_547_v14_)) >= ((d_548_v15_)[default__.safeIndex(d_547_v14_, len(d_548_v15_))])
            d_549_v16_: _dafny.MultiSet
            d_549_v16_ = _dafny.MultiSet([d_533_v0_, d_533_v0_])
            d_550_v17_: _dafny.Map
            d_550_v17_ = _dafny.Map({d_533_v0_: d_549_v16_})
            d_551_v18_: int
            d_551_v18_ = len(d_550_v17_)
            source8_ = d_551_v18_
            d_552___mcc_h0_ = source8_
            d_553_cf0_ = d_552___mcc_h0_
            d_554_v19_: C0
            nw95_ = C0()
            nw95_.ctor__(d_551_v18_, 90)
            d_554_v19_ = nw95_
            d_555_v20_: _dafny.Map
            d_555_v20_ = _dafny.Map({not(d_533_v0_): -251})
            d_556_v21_: _dafny.Seq
            d_556_v21_ = _dafny.SeqWithoutIsStrInference([d_555_v20_])
            d_556_v21_ = _dafny.SeqWithoutIsStrInference([d_555_v20_ for d_557_i0_ in range(default__.abs(787))])
            d_558_v22_: str
            d_558_v22_ = _dafny.CodePoint('h')
            d_559_v23_: _dafny.Seq
            d_559_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nntcghypc"))
            d_560_v24_: _dafny.Array
            nw96_ = _dafny.Array(int(0), 8)
            d_560_v24_ = nw96_
            index71_ = default__.safeIndex(426, (d_560_v24_).length(0))
            (d_560_v24_)[index71_] = d_553_cf0_
            d_561_v25_: _dafny.Array
            nw97_ = _dafny.Array(False, 3)
            d_561_v25_ = nw97_
            index72_ = default__.safeIndex(148, (d_561_v25_).length(0))
            (d_561_v25_)[index72_] = (d_549_v16_).issubset(d_549_v16_)
            d_562_v26_: _dafny.Seq
            d_562_v26_ = _dafny.SeqWithoutIsStrInference([d_561_v25_])
            d_563_v27_: _dafny.Map
            d_563_v27_ = _dafny.Map({d_562_v26_: d_558_v22_})
            index73_ = default__.safeIndex(426, (d_560_v24_).length(0))
            index74_ = default__.safeIndex(148, (d_561_v25_).length(0))
            rhs104_ = ((d_563_v27_)[d_562_v26_] if (d_562_v26_) in (d_563_v27_) else d_558_v22_)
            rhs105_ = (_dafny.SeqWithoutIsStrInference([(d_559_v23_)[default__.safeIndex(d_553_cf0_, len(d_559_v23_))] for d_564_i1_ in range(default__.abs(597))])) + (_dafny.SeqWithoutIsStrInference([d_558_v22_ for d_565_i2_ in range(default__.abs(657))]))
            rhs106_ = d_553_cf0_
            rhs107_ = not(d_533_v0_)
            lhs72_ = d_560_v24_
            lhs73_ = default__.safeIndex(426, (d_560_v24_).length(0))
            lhs74_ = d_561_v25_
            lhs75_ = default__.safeIndex(148, (d_561_v25_).length(0))
            d_558_v22_ = rhs104_
            d_559_v23_ = rhs105_
            lhs72_[lhs73_] = rhs106_
            lhs74_[lhs75_] = rhs107_
            d_555_v20_ = (d_555_v20_).set(default__.fm1(d_533_v0_, d_559_v23_, globalState), (d_560_v24_)[default__.safeIndex(426, (d_560_v24_).length(0))])
            d_566_v28_: C0
            nw98_ = C0()
            nw98_.ctor__(d_551_v18_, d_547_v14_)
            d_566_v28_ = nw98_
            d_567_v29_: _dafny.MultiSet
            d_567_v29_ = _dafny.MultiSet([d_566_v28_])
            d_568_v30_: _dafny.MultiSet
            d_568_v30_ = d_567_v29_
            d_569_v31_: _dafny.Map
            d_569_v31_ = _dafny.Map({d_547_v14_: d_547_v14_})
            d_570_v32_: _dafny.Array
            nw99_ = _dafny.Array(False, 25)
            d_570_v32_ = nw99_
            d_571_v33_: _dafny.MultiSet
            d_571_v33_ = _dafny.MultiSet([len(default__.fm8(d_547_v14_, d_547_v14_, globalState)), d_547_v14_])
            d_572_v34_: D2
            d_572_v34_ = D2_DC5(d_533_v0_, len(d_569_v31_), d_570_v32_, d_571_v33_)
            d_573_v35_: _dafny.Seq
            d_573_v35_ = _dafny.SeqWithoutIsStrInference([d_533_v0_, not(d_533_v0_), d_533_v0_])
            d_574_v36_: _dafny.Array
            nw100_ = _dafny.Array(None, 16)
            nw100_[int(0)] = _dafny.SeqWithoutIsStrInference([(d_572_v34_).cf5, d_533_v0_])
            nw100_[int(1)] = _dafny.SeqWithoutIsStrInference([d_533_v0_, d_533_v0_])
            nw100_[int(2)] = d_573_v35_
            nw100_[int(3)] = d_573_v35_
            nw100_[int(4)] = d_573_v35_
            nw100_[int(5)] = d_573_v35_
            nw100_[int(6)] = d_573_v35_
            nw100_[int(7)] = d_573_v35_
            nw100_[int(8)] = d_573_v35_
            nw100_[int(9)] = d_573_v35_
            nw100_[int(10)] = d_573_v35_
            nw100_[int(11)] = d_573_v35_
            nw100_[int(12)] = d_573_v35_
            nw100_[int(13)] = d_573_v35_
            nw100_[int(14)] = d_573_v35_
            nw100_[int(15)] = d_573_v35_
            d_574_v36_ = nw100_
            d_575_v37_: _dafny.Map
            d_575_v37_ = _dafny.Map({((d_568_v30_)) | (d_567_v29_): d_574_v36_})
            d_575_v37_ = (d_575_v37_).set(d_567_v29_, d_574_v36_)
            d_576_v38_: _dafny.Seq
            d_576_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eancxn"))
            d_576_v38_ = d_576_v38_
            (globalState).f0 = ((d_547_v14_) * (d_547_v14_)) < ((0) - (d_547_v14_))
        d_577_i3_: int
        d_577_i3_ = 0
        with _dafny.label("2"):
            while True:
                with _dafny.c_label("2"):
                    if (d_577_i3_) >= (100):
                        raise _dafny.Break("2")
                    d_577_i3_ = (d_577_i3_) + (1)
                    d_578_v39_: _dafny.Seq
                    d_578_v39_ = _dafny.SeqWithoutIsStrInference([d_533_v0_])
                    d_579_v40_: _dafny.Seq
                    d_579_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htdkpplcq"))
                    d_580_v41_: _dafny.Array
                    def lambda21_(d_581_i4_):
                        return (d_581_i4_) + (221)

                    init12_ = lambda21_
                    nw101_ = _dafny.Array(None, 22)
                    for i0_12_ in range(nw101_.length(0)):
                        nw101_[i0_12_] = init12_(i0_12_)
                    d_580_v41_ = nw101_
                    d_582_v42_: _dafny.Map
                    d_582_v42_ = _dafny.Map({True: d_580_v41_})
                    if (d_578_v39_)[default__.safeIndex(default__.safeModuloInt((0) - (len(d_579_v40_)), len(d_582_v42_)), len(d_578_v39_))]:
                        d_583_v43_: _dafny.Set
                        d_583_v43_ = _dafny.Set({d_533_v0_})
                        d_584_v44_: int
                        d_584_v44_ = 192
                        d_585_v45_: _dafny.Map
                        d_585_v45_ = _dafny.Map({d_533_v0_: d_584_v44_})
                        d_586_v46_: int
                        d_586_v46_ = len(d_585_v45_)
                        d_587_v47_: _dafny.Map
                        d_587_v47_ = _dafny.Map({d_583_v43_: (d_578_v39_)[default__.safeIndex((d_586_v46_), len(d_578_v39_))]})
                        d_587_v47_ = d_587_v47_
                        (globalState).f3 = (d_533_v0_) or (d_533_v0_)
                        r3 = d_533_v0_
                        (globalState).f10 = d_533_v0_
                        d_588_v48_: _dafny.Seq
                        d_588_v48_ = _dafny.SeqWithoutIsStrInference([d_584_v44_, d_584_v44_])
                        d_589_v49_: _dafny.MultiSet
                        d_589_v49_ = _dafny.MultiSet([len(_dafny.Set({(self).f20, (d_579_v40_)[default__.safeIndex(d_584_v44_, len(d_579_v40_))]}))])
                        d_590_v50_: _dafny.Array
                        nw102_ = _dafny.Array(False, 1)
                        d_590_v50_ = nw102_
                        d_591_v51_: D2
                        d_591_v51_ = D2_DC5(d_533_v0_, d_584_v44_, d_590_v50_, _dafny.MultiSet([d_584_v44_]))
                        d_592_v52_: _dafny.Array
                        nw103_ = _dafny.Array(None, 5)
                        nw103_[int(0)] = (d_589_v49_).issubset(_dafny.MultiSet(d_588_v48_))
                        nw103_[int(1)] = (d_578_v39_)[default__.safeIndex((0) - (d_584_v44_), len(d_578_v39_))]
                        nw103_[int(2)] = d_533_v0_
                        nw103_[int(3)] = d_533_v0_
                        nw103_[int(4)] = (d_591_v51_).cf5
                        d_592_v52_ = nw103_
                        index75_ = default__.safeIndex(784, (d_592_v52_).length(0))
                        (d_592_v52_)[index75_] = d_533_v0_
                        index76_ = default__.safeIndex(784, (d_592_v52_).length(0))
                        (d_592_v52_)[index76_] = d_533_v0_
                    elif True:
                        d_579_v40_ = (_dafny.SeqWithoutIsStrInference([(self).f20 for d_593_i5_ in range(default__.abs(764))])) + (d_579_v40_)
                        d_594_v53_: _dafny.Map
                        d_594_v53_ = _dafny.Map({d_578_v39_: 555})
                        d_595_v54_: int
                        d_595_v54_ = len(d_594_v53_)
                        d_596_v55_: C0
                        nw104_ = C0()
                        nw104_.ctor__(d_595_v54_, len(d_579_v40_))
                        d_596_v55_ = nw104_
                        d_597_v56_: _dafny.Seq
                        d_597_v56_ = _dafny.SeqWithoutIsStrInference([d_596_v55_, d_596_v55_, d_596_v55_])
                        d_598_v57_: int
                        d_598_v57_ = 30
                        d_599_v58_: D4
                        d_599_v58_ = D4_DC9(d_596_v55_)
                        d_600_v59_: _dafny.Array
                        nw105_ = _dafny.Array(None, 24)
                        nw105_[int(0)] = d_596_v55_
                        nw105_[int(1)] = d_596_v55_
                        nw105_[int(2)] = d_596_v55_
                        nw105_[int(3)] = d_596_v55_
                        nw105_[int(4)] = (d_597_v56_)[default__.safeIndex(d_598_v57_, len(d_597_v56_))]
                        nw105_[int(5)] = d_596_v55_
                        nw105_[int(6)] = d_596_v55_
                        nw105_[int(7)] = d_596_v55_
                        nw105_[int(8)] = d_596_v55_
                        nw105_[int(9)] = d_596_v55_
                        nw105_[int(10)] = d_596_v55_
                        nw105_[int(11)] = d_596_v55_
                        nw105_[int(12)] = d_596_v55_
                        nw105_[int(13)] = d_596_v55_
                        nw105_[int(14)] = d_596_v55_
                        nw105_[int(15)] = d_596_v55_
                        nw105_[int(16)] = d_596_v55_
                        nw105_[int(17)] = d_596_v55_
                        nw105_[int(18)] = d_596_v55_
                        nw105_[int(19)] = d_596_v55_
                        nw105_[int(20)] = d_596_v55_
                        nw105_[int(21)] = d_596_v55_
                        nw105_[int(22)] = d_596_v55_
                        nw105_[int(23)] = (d_599_v58_).cf15
                        d_600_v59_ = nw105_
                        index77_ = default__.safeIndex(899, (d_600_v59_).length(0))
                        (d_600_v59_)[index77_] = d_596_v55_
                        index78_ = default__.safeIndex(899, (d_600_v59_).length(0))
                        (d_600_v59_)[index78_] = d_596_v55_
                        (globalState).f2 = d_598_v57_
                        d_601_v60_: _dafny.Array
                        nw106_ = _dafny.Array(False, 17)
                        d_601_v60_ = nw106_
                        d_602_v61_: _dafny.Map
                        d_602_v61_ = _dafny.Map({d_533_v0_: d_601_v60_})
                        d_603_v62_: _dafny.Array
                        nw107_ = _dafny.Array(None, 17)
                        nw107_[int(0)] = d_601_v60_
                        nw107_[int(1)] = d_601_v60_
                        nw107_[int(2)] = d_601_v60_
                        nw107_[int(3)] = d_601_v60_
                        nw107_[int(4)] = d_601_v60_
                        nw107_[int(5)] = d_601_v60_
                        nw107_[int(6)] = d_601_v60_
                        nw107_[int(7)] = ((d_602_v61_)[d_533_v0_] if (d_533_v0_) in (d_602_v61_) else d_601_v60_)
                        nw107_[int(8)] = d_601_v60_
                        nw107_[int(9)] = d_601_v60_
                        nw107_[int(10)] = d_601_v60_
                        nw107_[int(11)] = d_601_v60_
                        nw107_[int(12)] = d_601_v60_
                        nw107_[int(13)] = d_601_v60_
                        nw107_[int(14)] = d_601_v60_
                        nw107_[int(15)] = d_601_v60_
                        nw107_[int(16)] = d_601_v60_
                        d_603_v62_ = nw107_
                        index79_ = default__.safeIndex(191, (d_603_v62_).length(0))
                        (d_603_v62_)[index79_] = d_601_v60_
                        d_604_v63_: _dafny.MultiSet
                        d_604_v63_ = _dafny.MultiSet([d_598_v57_, d_598_v57_, d_598_v57_])
                        index80_ = default__.safeIndex(191, (d_603_v62_).length(0))
                        rhs108_ = (d_604_v63_).cardinality
                        rhs109_ = d_601_v60_
                        lhs76_ = d_603_v62_
                        lhs77_ = default__.safeIndex(191, (d_603_v62_).length(0))
                        d_598_v57_ = rhs108_
                        lhs76_[lhs77_] = rhs109_
                        (globalState).f2 = d_598_v57_
                    d_533_v0_ = d_533_v0_
                    d_605_v64_: _dafny.Array
                    nw108_ = _dafny.Array(None, 28)
                    nw108_[int(0)] = d_533_v0_
                    nw108_[int(1)] = d_533_v0_
                    nw108_[int(2)] = d_533_v0_
                    nw108_[int(3)] = d_533_v0_
                    nw108_[int(4)] = default__.fm1(True, d_579_v40_, globalState)
                    nw108_[int(5)] = d_533_v0_
                    nw108_[int(6)] = default__.fm1(d_533_v0_, d_579_v40_, globalState)
                    nw108_[int(7)] = d_533_v0_
                    nw108_[int(8)] = d_533_v0_
                    nw108_[int(9)] = d_533_v0_
                    nw108_[int(10)] = not(d_533_v0_)
                    nw108_[int(11)] = d_533_v0_
                    nw108_[int(12)] = d_533_v0_
                    nw108_[int(13)] = d_533_v0_
                    nw108_[int(14)] = d_533_v0_
                    nw108_[int(15)] = d_533_v0_
                    nw108_[int(16)] = True
                    nw108_[int(17)] = d_533_v0_
                    nw108_[int(18)] = d_533_v0_
                    nw108_[int(19)] = d_533_v0_
                    nw108_[int(20)] = d_533_v0_
                    nw108_[int(21)] = d_533_v0_
                    nw108_[int(22)] = d_533_v0_
                    nw108_[int(23)] = False
                    nw108_[int(24)] = d_533_v0_
                    nw108_[int(25)] = d_533_v0_
                    nw108_[int(26)] = d_533_v0_
                    nw108_[int(27)] = default__.fm1(d_533_v0_, (d_579_v40_).set(default__.safeIndex(628, len(d_579_v40_)), (self).f20), globalState)
                    d_605_v64_ = nw108_
                    d_606_v65_: D1
                    d_606_v65_ = D1_DC1(d_605_v64_)
                    d_607_v66_: _dafny.Seq
                    d_607_v66_ = _dafny.SeqWithoutIsStrInference([d_605_v64_])
                    d_608_v67_: _dafny.Array
                    nw109_ = _dafny.Array(None, 25)
                    nw109_[int(0)] = d_605_v64_
                    nw109_[int(1)] = d_605_v64_
                    nw109_[int(2)] = d_605_v64_
                    nw109_[int(3)] = d_605_v64_
                    nw109_[int(4)] = d_605_v64_
                    nw109_[int(5)] = d_605_v64_
                    nw109_[int(6)] = d_605_v64_
                    nw109_[int(7)] = (d_605_v64_ if d_533_v0_ else d_605_v64_)
                    nw109_[int(8)] = d_605_v64_
                    nw109_[int(9)] = d_605_v64_
                    nw109_[int(10)] = d_605_v64_
                    nw109_[int(11)] = d_605_v64_
                    nw109_[int(12)] = d_605_v64_
                    nw109_[int(13)] = d_605_v64_
                    nw109_[int(14)] = d_605_v64_
                    nw109_[int(15)] = d_605_v64_
                    nw109_[int(16)] = d_605_v64_
                    nw109_[int(17)] = d_605_v64_
                    nw109_[int(18)] = d_605_v64_
                    nw109_[int(19)] = d_605_v64_
                    nw109_[int(20)] = d_605_v64_
                    nw109_[int(21)] = (d_606_v65_).cf1
                    nw109_[int(22)] = d_605_v64_
                    nw109_[int(23)] = (d_605_v64_ if d_533_v0_ else (d_607_v66_)[default__.safeIndex(len(d_578_v39_), len(d_607_v66_))])
                    nw109_[int(24)] = d_605_v64_
                    d_608_v67_ = nw109_
                    nw110_ = _dafny.Array(_dafny.Array(None, 0), 5)
                    d_608_v67_ = nw110_
                    d_609_v68_: int
                    d_609_v68_ = 14
                    r2 = (default__.safeDivisionInt(d_609_v68_, len(_dafny.Set({d_579_v40_, d_579_v40_})))) > (len(d_578_v39_))
                    pass
            pass
        d_610_v69_: _dafny.Array
        nw111_ = _dafny.Array(False, 3)
        d_610_v69_ = nw111_
        d_611_v70_: D1
        d_611_v70_ = D1_DC1(d_610_v69_)
        d_612_v71_: _dafny.Map
        d_612_v71_ = _dafny.Map({d_611_v70_: 421})
        d_613_v72_: int
        d_613_v72_ = 662
        d_614_v73_: _dafny.Map
        d_614_v73_ = _dafny.Map({((d_612_v71_)[d_611_v70_] if (d_611_v70_) in (d_612_v71_) else (0) - (d_613_v72_)): d_613_v72_})
        d_615_v74_: _dafny.Seq
        d_615_v74_ = _dafny.SeqWithoutIsStrInference([(self).f20])
        d_614_v73_ = (d_614_v73_).set((len(d_615_v74_)) - (d_613_v72_), -351)
        index81_ = default__.safeIndex(401, (d_610_v69_).length(0))
        (d_610_v69_)[index81_] = d_533_v0_
        d_616_v75_: _dafny.Seq
        d_616_v75_ = _dafny.SeqWithoutIsStrInference([d_613_v72_])
        index82_ = default__.safeIndex(401, (d_610_v69_).length(0))
        (d_610_v69_)[index82_] = not (False) or ((d_616_v75_) < (d_616_v75_))
        d_617_v76_: _dafny.Seq
        d_617_v76_ = _dafny.SeqWithoutIsStrInference([(d_610_v69_)[default__.safeIndex(401, (d_610_v69_).length(0))]])
        d_618_v77_: D1
        d_618_v77_ = D1_DC2(not ((d_610_v69_)[default__.safeIndex(401, (d_610_v69_).length(0))]) or ((d_617_v76_)[default__.safeIndex(d_613_v72_, len(d_617_v76_))]))
        def iife68_(_pat_let25_0):
            def iife69_(d_619_dt__update__tmp_h0_):
                def iife70_(_pat_let26_0):
                    def iife71_(d_620_dt__update_hcf2_h0_):
                        return D1_DC2(d_620_dt__update_hcf2_h0_)
                    return iife71_(_pat_let26_0)
                return iife70_(False)
            return iife69_(_pat_let25_0)
        d_618_v77_ = iife68_(d_618_v77_)
        d_621_v78_: _dafny.Map
        d_621_v78_ = _dafny.Map({(0) - (d_613_v72_): (d_610_v69_)[default__.safeIndex(401, (d_610_v69_).length(0))]})
        d_622_v79_: int
        d_623_v80_: int
        out24_: int
        out25_: int
        out24_, out25_ = default__.m0(((d_610_v69_)[default__.safeIndex(401, (d_610_v69_).length(0))] if (d_617_v76_)[default__.safeIndex((d_616_v75_)[default__.safeIndex(d_613_v72_, len(d_616_v75_))], len(d_617_v76_))] else (d_610_v69_)[default__.safeIndex(401, (d_610_v69_).length(0))]), (d_621_v78_).set(d_613_v72_, d_533_v0_), globalState)
        d_622_v79_ = out24_
        d_623_v80_ = out25_
        r0 = not((d_610_v69_)[default__.safeIndex(401, (d_610_v69_).length(0))])
        d_624_v81_: D5
        d_624_v81_ = D5_DC12(d_615_v74_)
        r1 = len(((d_624_v81_).cf18) + ((d_615_v74_) + (d_615_v74_)))
        r2 = default__.fm1(d_533_v0_, d_615_v74_, globalState)
        r3 = not (((d_610_v69_)[default__.safeIndex(401, (d_610_v69_).length(0))]) == ((d_610_v69_)[default__.safeIndex(401, (d_610_v69_).length(0))])) or (not(d_533_v0_))
        return r0, r1, r2, r3

    @property
    def f20(self):
        return self._f20

class C2:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self):
        pass
        pass

    def m2(self, globalState):
        r0: int = int(0)
        d_625_v0_: int
        d_625_v0_ = 793
        (globalState).f3 = (d_625_v0_) == (d_625_v0_)
        d_626_v1_: int
        d_626_v1_ = 84
        d_627_v2_: str
        d_627_v2_ = _dafny.CodePoint('f')
        d_628_v3_: C1
        nw112_ = C1()
        nw112_.ctor__(d_627_v2_)
        d_628_v3_ = nw112_
        d_629_v4_: _dafny.MultiSet
        d_629_v4_ = _dafny.MultiSet([d_628_v3_])
        d_630_v5_: _dafny.MultiSet
        d_630_v5_ = _dafny.MultiSet([906])
        d_631_v6_: C0
        nw113_ = C0()
        nw113_.ctor__(d_626_v1_, (len(_dafny.Set({d_629_v4_}))) + (((d_630_v5_)[d_625_v0_] if (d_625_v0_) in (d_630_v5_) else 302)))
        d_631_v6_ = nw113_
        d_632_v7_: bool
        d_632_v7_ = True
        d_633_v8_: _dafny.Seq
        d_633_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfvwqmbea"))
        d_634_v9_: D5
        d_634_v9_ = D5_DC13(d_625_v0_, d_625_v0_, default__.fm1(not(d_632_v7_), d_633_v8_, globalState))
        def lambda22_(source9_):
            if source9_.is_DC13:
                d_635___mcc_h0_ = source9_.cf19
                d_636___mcc_h1_ = source9_.cf20
                d_637___mcc_h2_ = source9_.cf21
                d_638_cf21_ = d_637___mcc_h2_
                d_639_cf20_ = d_636___mcc_h1_
                d_640_cf19_ = d_635___mcc_h0_
                return False
            elif True:
                d_641___mcc_h3_ = source9_.cf18
                d_642_cf18_ = d_641___mcc_h3_
                return False

        (globalState).f3 = lambda22_(d_634_v9_)
        d_643_v10_: _dafny.Array
        nw114_ = _dafny.Array(_dafny.Seq({}), 10)
        d_643_v10_ = nw114_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_643_v10_).length(0)):
            d_644_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_644_i0_)) and ((d_644_i0_) < ((d_643_v10_).length(0)))):
                (d_643_v10_)[(d_644_i0_)] = ((_dafny.SeqWithoutIsStrInference([d_632_v7_])) + (_dafny.SeqWithoutIsStrInference([d_632_v7_]))) + ((_dafny.SeqWithoutIsStrInference([d_632_v7_, d_632_v7_])) + (_dafny.SeqWithoutIsStrInference([d_632_v7_])))
        d_645_v11_: D5
        d_645_v11_ = D5_DC12(d_633_v8_)
        pat_let_tv20_ = d_632_v7_
        def lambda23_(source10_):
            if source10_.is_DC13:
                d_646___mcc_h4_ = source10_.cf19
                d_647___mcc_h5_ = source10_.cf20
                d_648___mcc_h6_ = source10_.cf21
                d_649_cf21_ = d_648___mcc_h6_
                d_650_cf20_ = d_647___mcc_h5_
                d_651_cf19_ = d_646___mcc_h4_
                return False
            elif True:
                d_652___mcc_h7_ = source10_.cf18
                d_653_cf18_ = d_652___mcc_h7_
                return not(pat_let_tv20_)

        if lambda23_(d_645_v11_):
            d_632_v7_ = d_632_v7_
            d_654_v12_: _dafny.Set
            d_654_v12_ = _dafny.Set({not(d_632_v7_), d_632_v7_, d_632_v7_})
            d_655_v13_: _dafny.Array
            def lambda24_(d_656_i1_):
                return False

            init13_ = lambda24_
            nw115_ = _dafny.Array(None, 24)
            for i0_13_ in range(nw115_.length(0)):
                nw115_[i0_13_] = init13_(i0_13_)
            d_655_v13_ = nw115_
            d_657_v14_: _dafny.Map
            d_657_v14_ = _dafny.Map({(len(d_654_v12_) if d_632_v7_ else d_625_v0_): d_655_v13_})
            d_657_v14_ = (d_657_v14_).set(377, d_655_v13_)
            if ((d_625_v0_) + (d_625_v0_)) > (d_625_v0_):
                (globalState).f2 = d_625_v0_
                (globalState).f10 = not(((d_627_v2_ if d_632_v7_ else (d_628_v3_).f20)) in (d_633_v8_))
                (globalState).f3 = (default__.safeDivisionInt(d_625_v0_, (0) - (d_625_v0_))) < (d_625_v0_)
                r0 = d_625_v0_
                d_658_v15_: D6
                d_658_v15_ = D6_DC14(d_628_v3_)
                d_659_v16_: _dafny.Array
                nw116_ = _dafny.Array(None, 9)
                nw116_[int(0)] = (d_628_v3_ if d_632_v7_ else (d_658_v15_).cf22)
                nw116_[int(1)] = d_628_v3_
                nw116_[int(2)] = d_628_v3_
                nw116_[int(3)] = d_628_v3_
                nw116_[int(4)] = d_628_v3_
                nw116_[int(5)] = d_628_v3_
                nw116_[int(6)] = d_628_v3_
                nw116_[int(7)] = d_628_v3_
                nw116_[int(8)] = d_628_v3_
                d_659_v16_ = nw116_
                d_659_v16_ = d_659_v16_
            elif True:
                d_660_v17_: _dafny.Array
                nw117_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                d_660_v17_ = nw117_
                index83_ = default__.safeIndex(553, (d_660_v17_).length(0))
                (d_660_v17_)[index83_] = ((d_633_v8_).set(default__.safeIndex(d_625_v0_, len(d_633_v8_)), d_627_v2_) if d_632_v7_ else d_633_v8_)
                d_661_v18_: _dafny.Array
                nw118_ = _dafny.Array(None, 23)
                d_661_v18_ = nw118_
                d_662_v19_: _dafny.MultiSet
                d_662_v19_ = _dafny.MultiSet([d_661_v18_])
                index84_ = default__.safeIndex(553, (d_660_v17_).length(0))
                rhs110_ = (not(d_632_v7_) if d_632_v7_ else d_632_v7_)
                rhs111_ = (_dafny.SeqWithoutIsStrInference([(d_628_v3_).f20 for d_663_i2_ in range(default__.abs(143))])).set(default__.safeIndex((0) - (d_625_v0_), len(_dafny.SeqWithoutIsStrInference([(d_628_v3_).f20 for d_664_i2_ in range(default__.abs(143))]))), _dafny.CodePoint('h'))
                rhs112_ = (d_662_v19_).issubset((d_662_v19_) | (d_662_v19_))
                rhs113_ = d_625_v0_
                lhs78_ = globalState
                lhs79_ = d_660_v17_
                lhs80_ = default__.safeIndex(553, (d_660_v17_).length(0))
                lhs81_ = globalState
                lhs78_.f0 = rhs110_
                lhs79_[lhs80_] = rhs111_
                lhs81_.f3 = rhs112_
                r0 = rhs113_
                d_665_v20_: _dafny.Map
                d_665_v20_ = _dafny.Map({default__.safeDivisionInt(d_625_v0_, (d_630_v5_).cardinality): d_632_v7_})
                d_666_v21_: _dafny.Set
                d_666_v21_ = _dafny.Set({_dafny.CodePoint('e')})
                d_665_v20_ = (d_665_v20_).set(((0) - (default__.fm0(globalState))) + (d_625_v0_), (d_666_v21_).ispropersubset(d_666_v21_))
                index85_ = default__.safeIndex(626, (d_655_v13_).length(0))
                (d_655_v13_)[index85_] = True
                d_667_v22_: _dafny.MultiSet
                d_667_v22_ = _dafny.MultiSet([d_632_v7_, d_632_v7_])
                index86_ = default__.safeIndex(626, (d_655_v13_).length(0))
                (d_655_v13_)[index86_] = ((d_667_v22_).ispropersubset(d_667_v22_) if not (not(d_632_v7_)) or (d_632_v7_) else d_632_v7_)
                d_632_v7_ = not(((d_660_v17_)[default__.safeIndex(553, (d_660_v17_).length(0))]) == ((d_660_v17_)[default__.safeIndex(553, (d_660_v17_).length(0))]))
                d_668_v23_: _dafny.Map
                d_668_v23_ = _dafny.Map({606: d_627_v2_})
                d_669_v24_: _dafny.Seq
                d_669_v24_ = _dafny.SeqWithoutIsStrInference([d_668_v23_])
                (globalState).f0 = (d_625_v0_) <= (len((d_669_v24_)[default__.safeIndex(744, len(d_669_v24_))]))
            d_670_v25_: _dafny.Array
            nw119_ = _dafny.Array(_dafny.Set({}), 6)
            d_670_v25_ = nw119_
            d_671_v26_: _dafny.Set
            d_671_v26_ = _dafny.Set({d_655_v13_})
            index87_ = default__.safeIndex(453, (d_670_v25_).length(0))
            (d_670_v25_)[index87_] = d_671_v26_
            index88_ = default__.safeIndex(453, (d_670_v25_).length(0))
            (d_670_v25_)[index88_] = (d_671_v26_) - (d_671_v26_)
            (globalState).f0 = d_632_v7_
        elif True:
            (globalState).f2 = d_625_v0_
            d_672_v27_: _dafny.Map
            d_672_v27_ = _dafny.Map({len(d_633_v8_): d_628_v3_})
            d_673_v28_: _dafny.Map
            d_673_v28_ = _dafny.Map({d_672_v27_: not(d_632_v7_)})
            d_673_v28_ = (d_673_v28_).set(d_672_v27_, False)
            if d_632_v7_:
                d_674_v29_: C1
                nw120_ = C1()
                nw120_.ctor__(((d_628_v3_).f20 if d_632_v7_ else _dafny.CodePoint('m')))
                d_674_v29_ = nw120_
                (globalState).f9 = (default__.safeModuloInt(d_625_v0_, len(default__.fm10(globalState)))) - ((d_625_v0_) + (len(d_633_v8_)))
                d_675_v30_: _dafny.Array
                def lambda25_(d_676_i3_):
                    return True

                init14_ = lambda25_
                nw121_ = _dafny.Array(None, 5)
                for i0_14_ in range(nw121_.length(0)):
                    nw121_[i0_14_] = init14_(i0_14_)
                d_675_v30_ = nw121_
                d_677_v31_: D2
                d_677_v31_ = D2_DC5(d_632_v7_, 324, d_675_v30_, d_630_v5_)
                d_678_v32_: _dafny.MultiSet
                d_678_v32_ = _dafny.MultiSet([d_677_v31_])
                d_679_v33_: _dafny.Map
                d_679_v33_ = _dafny.Map({d_632_v7_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "saqvvnypo"))})
                d_680_v34_: _dafny.Seq
                d_680_v34_ = _dafny.SeqWithoutIsStrInference([d_678_v32_, (d_678_v32_).set(d_677_v31_, default__.abs((0) - (len(d_679_v33_))))])
                d_680_v34_ = d_680_v34_
                (globalState).f0 = not (not(d_632_v7_)) or (d_632_v7_)
                (globalState).f2 = default__.fm0(globalState)
            elif True:
                d_681_v35_: C1
                nw122_ = C1()
                nw122_.ctor__((d_628_v3_).f20)
                d_681_v35_ = nw122_
                d_682_v36_: _dafny.MultiSet
                d_682_v36_ = _dafny.MultiSet([d_633_v8_])
                d_683_v37_: _dafny.Seq
                d_683_v37_ = _dafny.SeqWithoutIsStrInference([d_634_v9_, D5_DC13(d_625_v0_, (d_682_v36_).cardinality, d_632_v7_)])
                d_684_v38_: _dafny.Seq
                d_684_v38_ = _dafny.SeqWithoutIsStrInference([d_632_v7_, (d_683_v37_) <= ((d_683_v37_).set(default__.safeIndex(d_625_v0_, len(d_683_v37_)), d_634_v9_))])
                (globalState).f10 = (d_684_v38_)[default__.safeIndex(default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_632_v7_, False])), d_625_v0_), len(d_684_v38_))]
                (globalState).f2 = d_625_v0_
                d_685_v39_: _dafny.Map
                d_685_v39_ = _dafny.Map({d_645_v11_: _dafny.CodePoint('d')})
                d_686_v40_: _dafny.Array
                nw123_ = _dafny.Array(None, 27)
                nw123_[int(0)] = d_627_v2_
                nw123_[int(1)] = default__.fm7(d_627_v2_, default__.fm1(d_632_v7_, _dafny.SeqWithoutIsStrInference([(d_628_v3_).f20 for d_687_i4_ in range(default__.abs(295))]), globalState), globalState)
                nw123_[int(2)] = (d_681_v35_).f20
                nw123_[int(3)] = (d_628_v3_).f20
                nw123_[int(4)] = (d_628_v3_).f20
                nw123_[int(5)] = (d_628_v3_).f20
                nw123_[int(6)] = (d_628_v3_).f20
                nw123_[int(7)] = (d_681_v35_).f20
                nw123_[int(8)] = (d_628_v3_).f20
                nw123_[int(9)] = _dafny.CodePoint('b')
                nw123_[int(10)] = d_627_v2_
                nw123_[int(11)] = (d_628_v3_).f20
                nw123_[int(12)] = d_627_v2_
                nw123_[int(13)] = d_627_v2_
                nw123_[int(14)] = (d_628_v3_).f20
                nw123_[int(15)] = d_627_v2_
                nw123_[int(16)] = (d_628_v3_).f20
                nw123_[int(17)] = ((d_685_v39_)[d_645_v11_] if (d_645_v11_) in (d_685_v39_) else _dafny.CodePoint('f'))
                nw123_[int(18)] = (d_628_v3_).f20
                nw123_[int(19)] = (d_681_v35_).f20
                nw123_[int(20)] = (d_681_v35_).f20
                nw123_[int(21)] = d_627_v2_
                nw123_[int(22)] = d_627_v2_
                nw123_[int(23)] = (d_628_v3_).f20
                nw123_[int(24)] = _dafny.CodePoint('j')
                nw123_[int(25)] = (d_628_v3_).f20
                nw123_[int(26)] = d_627_v2_
                d_686_v40_ = nw123_
                index89_ = default__.safeIndex(199, (d_686_v40_).length(0))
                (d_686_v40_)[index89_] = (d_681_v35_).f20
                d_688_v41_: _dafny.Array
                nw124_ = _dafny.Array(_dafny.Set({}), 5)
                d_688_v41_ = nw124_
                d_689_v42_: _dafny.Map
                d_689_v42_ = _dafny.Map({d_688_v41_: (d_628_v3_).f20})
                d_690_v43_: _dafny.Seq
                d_690_v43_ = _dafny.SeqWithoutIsStrInference([d_688_v41_])
                index90_ = default__.safeIndex(199, (d_686_v40_).length(0))
                (d_686_v40_)[index90_] = ((d_689_v42_)[(d_690_v43_)[default__.safeIndex((0) - (d_625_v0_), len(d_690_v43_))]] if ((d_690_v43_)[default__.safeIndex((0) - (d_625_v0_), len(d_690_v43_))]) in (d_689_v42_) else (d_628_v3_).f20)
                d_691_v44_: _dafny.Array
                def lambda26_(d_692_v7_):
                    def lambda27_(d_693_i5_):
                        return d_692_v7_

                    return lambda27_

                init15_ = lambda26_(d_632_v7_)
                nw125_ = _dafny.Array(None, 10)
                for i0_15_ in range(nw125_.length(0)):
                    nw125_[i0_15_] = init15_(i0_15_)
                d_691_v44_ = nw125_
                d_694_v45_: _dafny.MultiSet
                d_694_v45_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_625_v0_])])
                d_695_v46_: _dafny.Seq
                d_695_v46_ = _dafny.SeqWithoutIsStrInference([d_694_v45_])
                index91_ = default__.safeIndex(330, (d_691_v44_).length(0))
                (d_691_v44_)[index91_] = ((d_695_v46_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_628_v3_).f20])), len(d_695_v46_))]).ispropersubset(d_694_v45_)
                d_696_v47_: D6
                d_696_v47_ = D6_DC17((d_686_v40_)[default__.safeIndex(199, (d_686_v40_).length(0))], d_632_v7_, not(d_632_v7_), (d_628_v3_).f20)
                index92_ = default__.safeIndex(330, (d_691_v44_).length(0))
                (d_691_v44_)[index92_] = (d_696_v47_).cf34
            (globalState).f2 = 89
            d_697_v48_: C1
            nw126_ = C1()
            nw126_.ctor__(_dafny.CodePoint('d'))
            d_697_v48_ = nw126_
        d_698_v49_: _dafny.Set
        d_698_v49_ = _dafny.Set({d_632_v7_})
        d_699_v50_: D6
        d_699_v50_ = D6_DC16(d_625_v0_, d_632_v7_, len(d_633_v8_), len(d_698_v49_), d_632_v7_)
        if not (d_632_v7_) or ((d_632_v7_ if d_632_v7_ else (d_699_v50_).cf32)):
            d_700_v51_: _dafny.Seq
            d_700_v51_ = _dafny.SeqWithoutIsStrInference([(len(_dafny.SeqWithoutIsStrInference([d_632_v7_, d_632_v7_]))) - (d_625_v0_), d_625_v0_])
            d_700_v51_ = d_700_v51_
            d_632_v7_ = d_632_v7_
            d_701_v52_: C1
            nw127_ = C1()
            nw127_.ctor__(_dafny.CodePoint('k'))
            d_701_v52_ = nw127_
            (globalState).f3 = d_632_v7_
            d_627_v2_ = (d_701_v52_).f20
        elif True:
            d_702_v53_: C0
            nw128_ = C0()
            nw128_.ctor__(d_626_v1_, d_625_v0_)
            d_702_v53_ = nw128_
            d_703_v54_: _dafny.Map
            d_703_v54_ = _dafny.Map({d_625_v0_: d_632_v7_})
            d_704_v55_: _dafny.Seq
            d_704_v55_ = _dafny.SeqWithoutIsStrInference([d_632_v7_, d_632_v7_])
            d_705_v56_: _dafny.Array
            nw129_ = _dafny.Array(None, 25)
            nw129_[int(0)] = (d_703_v54_) == (d_703_v54_)
            nw129_[int(1)] = d_632_v7_
            nw129_[int(2)] = (d_625_v0_) <= (d_625_v0_)
            nw129_[int(3)] = not(d_632_v7_)
            nw129_[int(4)] = d_632_v7_
            nw129_[int(5)] = d_632_v7_
            nw129_[int(6)] = (d_704_v55_)[default__.safeIndex(d_625_v0_, len(d_704_v55_))]
            nw129_[int(7)] = d_632_v7_
            nw129_[int(8)] = (d_632_v7_ if not(False) else d_632_v7_)
            nw129_[int(9)] = default__.fm1(d_632_v7_, default__.fm6(not(d_632_v7_), globalState), globalState)
            nw129_[int(10)] = d_632_v7_
            nw129_[int(11)] = d_632_v7_
            nw129_[int(12)] = (len(d_698_v49_)) >= (d_625_v0_)
            nw129_[int(13)] = (d_630_v5_) != (d_630_v5_)
            nw129_[int(14)] = d_632_v7_
            nw129_[int(15)] = d_632_v7_
            nw129_[int(16)] = d_632_v7_
            nw129_[int(17)] = d_632_v7_
            nw129_[int(18)] = d_632_v7_
            nw129_[int(19)] = True
            nw129_[int(20)] = d_632_v7_
            nw129_[int(21)] = (d_699_v50_).cf29
            nw129_[int(22)] = d_632_v7_
            nw129_[int(23)] = d_632_v7_
            nw129_[int(24)] = d_632_v7_
            d_705_v56_ = nw129_
            rhs114_ = (d_633_v8_) + (d_633_v8_)
            rhs115_ = d_702_v53_
            rhs116_ = d_633_v8_
            rhs117_ = d_705_v56_
            lhs82_ = globalState
            d_633_v8_ = rhs114_
            d_702_v53_ = rhs115_
            d_633_v8_ = rhs116_
            lhs82_.f11 = rhs117_
            (globalState).f9 = d_625_v0_
            d_706_v57_: _dafny.Map
            d_706_v57_ = _dafny.Map({d_632_v7_: d_632_v7_})
            pat_let_tv21_ = d_704_v55_
            d_707_v58_: _dafny.Seq
            def iife72_(_pat_let27_0):
                def iife73_(d_708_dt__update__tmp_h0_):
                    def iife74_(_pat_let28_0):
                        def iife75_(d_709_dt__update_hcf19_h0_):
                            return D5_DC13(d_709_dt__update_hcf19_h0_, (d_708_dt__update__tmp_h0_).cf20, (d_708_dt__update__tmp_h0_).cf21)
                        return iife75_(_pat_let28_0)
                    return iife74_(len(pat_let_tv21_))
                return iife73_(_pat_let27_0)
            d_707_v58_ = _dafny.SeqWithoutIsStrInference([iife72_(d_634_v9_), d_634_v9_])
            d_706_v57_ = (d_706_v57_).set(default__.fm1(d_632_v7_, _dafny.SeqWithoutIsStrInference([(d_628_v3_).f20 for d_710_i6_ in range(default__.abs(269))]), globalState), (d_707_v58_) < (_dafny.SeqWithoutIsStrInference([d_634_v9_ for d_711_i7_ in range(default__.abs(782))])))
            d_712_v59_: D1
            d_712_v59_ = D1_DC2(d_632_v7_)
            d_713_v60_: D1
            d_713_v60_ = D1_DC3(d_712_v59_)
            d_713_v60_ = d_713_v60_
            if d_632_v7_:
                (globalState).f10 = ((d_632_v7_ if d_632_v7_ else default__.fm1(d_632_v7_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xtbk")), globalState)) if d_632_v7_ else d_632_v7_)
                d_714_v61_: _dafny.Array
                nw130_ = _dafny.Array(int(0), 21)
                d_714_v61_ = nw130_
                d_714_v61_ = d_714_v61_
                (globalState).f3 = d_632_v7_
                d_715_v62_: _dafny.Map
                d_715_v62_ = _dafny.Map({_dafny.MultiSet([len(d_706_v57_)]): d_625_v0_})
                d_715_v62_ = d_715_v62_
                (globalState).f2 = (d_625_v0_) - (d_625_v0_)
            elif True:
                (globalState).f10 = d_632_v7_
                d_716_v63_: _dafny.Array
                nw131_ = _dafny.Array(int(0), 21)
                d_716_v63_ = nw131_
                d_717_v64_: D4
                d_717_v64_ = D4_DC10(d_716_v63_)
                d_718_v65_: D4
                d_718_v65_ = D4_DC11(d_717_v64_)
                d_719_v66_: _dafny.Seq
                d_719_v66_ = _dafny.SeqWithoutIsStrInference([d_625_v0_, d_625_v0_])
                d_720_v67_: _dafny.Map
                d_720_v67_ = _dafny.Map({d_718_v65_: d_719_v66_})
                d_721_v68_: _dafny.Map
                d_721_v68_ = _dafny.Map({((d_720_v67_)[d_718_v65_] if (d_718_v65_) in (d_720_v67_) else d_719_v66_): default__.safeModuloInt(d_625_v0_, d_625_v0_)})
                d_721_v68_ = (d_721_v68_).set((d_719_v66_ if d_632_v7_ else d_719_v66_), 329)
                d_713_v60_ = (d_713_v60_ if (d_632_v7_) or (d_632_v7_) else d_713_v60_)
                d_625_v0_ = d_625_v0_
                d_722_v69_: _dafny.Map
                d_722_v69_ = _dafny.Map({d_705_v56_: d_632_v7_})
                d_723_v71_: _dafny.Set
                d_723_v71_ = _dafny.Set({d_625_v0_, d_625_v0_})
                def iife76_():
                    coll18_ = _dafny.Set()
                    compr_18_: int
                    for compr_18_ in _dafny.IntegerRange(-694, -854):
                        d_724_v70_: int = compr_18_
                        if ((-694) <= (d_724_v70_)) and ((d_724_v70_) < (-854)):
                            coll18_ = coll18_.union(_dafny.Set([(d_724_v70_) - (d_625_v0_)]))
                    return _dafny.Set(coll18_)
                (globalState).f0 = ((d_722_v69_)[d_705_v56_] if (d_705_v56_) in (d_722_v69_) else (d_723_v71_).ispropersubset(iife76_()
                ))
        r0 = (0) - ((d_625_v0_) - ((d_625_v0_) * (d_625_v0_)))
        return r0


class C3(T0):
    def  __init__(self):
        self._f17: int = int(0)
        self._f18: int = int(0)
        self._f19: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f19, f17, f18):
        (self)._f19 = f19
        (self)._f17 = f17
        (self)._f18 = f18

    def fm2(self, globalState):
        return (self).f17

    def fm3(self, p0, p1, globalState):
        return (self).f18

    def fm4(self, p0, p1, p2, globalState):
        return default__.safeModuloInt((self).f18, ((self).f18) + ((self).f18))

    def m1(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        d_725_v0_: _dafny.Seq
        d_725_v0_ = _dafny.SeqWithoutIsStrInference([(self).f18])
        hi3_ = ((self).f18) * (len((self).f19))
        for d_726_i0_ in range((0) - ((d_725_v0_)[default__.safeIndex(-233, len(d_725_v0_))]), hi3_):
            source11_ = (self).f17
            d_727___mcc_h0_ = source11_
            d_728_cf0_ = d_727___mcc_h0_
            (globalState).f2 = d_726_i0_
            d_729_v1_: bool
            d_729_v1_ = True
            d_730_v2_: _dafny.Map
            d_730_v2_ = _dafny.Map({not((d_729_v1_ if d_729_v1_ else d_729_v1_)): d_728_cf0_})
            (globalState).f2 = (0) - (((d_730_v2_)[(d_726_i0_) <= (len((self).f19))] if ((d_726_i0_) <= (len((self).f19))) in (d_730_v2_) else 720))
            (globalState).f9 = len((self).f19)
            d_731_v3_: _dafny.Array
            nw132_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_731_v3_ = nw132_
            d_732_v4_: _dafny.Map
            d_732_v4_ = _dafny.Map({d_729_v1_: d_729_v1_})
            d_733_v5_: _dafny.Array
            nw133_ = _dafny.Array(None, 6)
            nw133_[int(0)] = d_729_v1_
            nw133_[int(1)] = ((d_732_v4_)[not(d_729_v1_)] if (not(d_729_v1_)) in (d_732_v4_) else d_729_v1_)
            nw133_[int(2)] = d_729_v1_
            nw133_[int(3)] = default__.fm1(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ti")), globalState)
            nw133_[int(4)] = d_729_v1_
            nw133_[int(5)] = True
            d_733_v5_ = nw133_
            index93_ = default__.safeIndex(652, (d_731_v3_).length(0))
            (d_731_v3_)[index93_] = d_733_v5_
            index94_ = default__.safeIndex(99, (d_733_v5_).length(0))
            (d_733_v5_)[index94_] = True
            d_734_v6_: D1
            d_734_v6_ = D1_DC1(d_733_v5_)
            d_735_v7_: _dafny.Seq
            d_735_v7_ = _dafny.SeqWithoutIsStrInference([d_729_v1_])
            index95_ = default__.safeIndex(652, (d_731_v3_).length(0))
            index96_ = default__.safeIndex(99, (d_733_v5_).length(0))
            rhs118_ = (d_734_v6_).cf1
            rhs119_ = d_726_i0_
            rhs120_ = (d_729_v1_ if (d_735_v7_)[default__.safeIndex(d_726_i0_, len(d_735_v7_))] else d_729_v1_)
            rhs121_ = d_729_v1_
            lhs83_ = d_731_v3_
            lhs84_ = default__.safeIndex(652, (d_731_v3_).length(0))
            lhs85_ = globalState
            lhs86_ = globalState
            lhs87_ = d_733_v5_
            lhs88_ = default__.safeIndex(99, (d_733_v5_).length(0))
            lhs83_[lhs84_] = rhs118_
            lhs85_.f2 = rhs119_
            lhs86_.f0 = rhs120_
            lhs87_[lhs88_] = rhs121_
            d_736_v8_: _dafny.Array
            nw134_ = _dafny.Array(False, 19)
            d_736_v8_ = nw134_
            d_737_v9_: D1
            d_737_v9_ = D1_DC1(d_736_v8_)
            source12_ = d_737_v9_
            if source12_.is_DC2:
                d_738___mcc_h1_ = source12_.cf2
                d_739_cf2_ = d_738___mcc_h1_
                d_740_v10_: str
                d_740_v10_ = _dafny.CodePoint('x')
                d_741_v11_: C1
                nw135_ = C1()
                nw135_.ctor__(d_740_v10_)
                d_741_v11_ = nw135_
                (globalState).f9 = (self).f18
                d_742_v12_: C0
                nw136_ = C0()
                nw136_.ctor__((self).f17, (0) - ((self).f18))
                d_742_v12_ = nw136_
                d_743_v13_: _dafny.Map
                d_743_v13_ = _dafny.Map({d_726_i0_: 549})
                d_744_v14_: _dafny.Map
                d_744_v14_ = _dafny.Map({d_739_cf2_: (self).f18})
                index97_ = default__.safeIndex(734, (d_736_v8_).length(0))
                (d_736_v8_)[index97_] = (_dafny.Set({d_726_i0_, (self).f18})) != (default__.fm8(((d_744_v14_)[d_739_cf2_] if (d_739_cf2_) in (d_744_v14_) else d_726_i0_), d_726_i0_, globalState))
                index98_ = default__.safeIndex(734, (d_736_v8_).length(0))
                rhs122_ = d_740_v10_
                rhs123_ = d_726_i0_
                rhs124_ = (d_743_v13_) | (d_743_v13_)
                rhs125_ = (d_726_i0_) < ((self).f18)
                lhs89_ = globalState
                lhs90_ = d_736_v8_
                lhs91_ = default__.safeIndex(734, (d_736_v8_).length(0))
                d_740_v10_ = rhs122_
                lhs89_.f9 = rhs123_
                d_743_v13_ = rhs124_
                lhs90_[lhs91_] = rhs125_
            elif source12_.is_DC1:
                d_745___mcc_h2_ = source12_.cf1
                d_746_cf1_ = d_745___mcc_h2_
                d_747_v15_: bool
                d_747_v15_ = False
                index99_ = default__.safeIndex(660, (d_736_v8_).length(0))
                (d_736_v8_)[index99_] = d_747_v15_
                d_748_v16_: T0
                nw137_ = C0()
                nw137_.ctor__((self).f18, d_726_i0_)
                d_748_v16_ = nw137_
                d_749_v17_: _dafny.Set
                d_749_v17_ = _dafny.Set({d_748_v16_, d_748_v16_})
                index100_ = default__.safeIndex(660, (d_736_v8_).length(0))
                (d_736_v8_)[index100_] = (d_749_v17_).ispropersubset((d_749_v17_) - (d_749_v17_))
                d_750_v18_: _dafny.Array
                nw138_ = _dafny.Array(int(0), 27)
                d_750_v18_ = nw138_
                index101_ = default__.safeIndex(162, (d_750_v18_).length(0))
                (d_750_v18_)[index101_] = 874
                index102_ = default__.safeIndex(660, (d_736_v8_).length(0))
                index103_ = default__.safeIndex(162, (d_750_v18_).length(0))
                rhs126_ = (d_736_v8_)[default__.safeIndex(660, (d_736_v8_).length(0))]
                rhs127_ = (self).f18
                lhs92_ = d_736_v8_
                lhs93_ = default__.safeIndex(660, (d_736_v8_).length(0))
                lhs94_ = d_750_v18_
                lhs95_ = default__.safeIndex(162, (d_750_v18_).length(0))
                lhs92_[lhs93_] = rhs126_
                lhs94_[lhs95_] = rhs127_
                d_748_v16_ = d_748_v16_
                d_751_v19_: _dafny.Set
                d_751_v19_ = _dafny.Set({(self).f18, d_726_i0_})
                d_752_v20_: _dafny.Map
                d_752_v20_ = _dafny.Map({d_751_v19_: d_746_cf1_})
                d_752_v20_ = (d_752_v20_).set(d_751_v19_, d_746_cf1_)
            elif True:
                d_753___mcc_h3_ = source12_.cf3
                d_754_cf3_ = d_753___mcc_h3_
                d_755_v21_: bool
                d_755_v21_ = True
                d_756_v22_: str
                d_756_v22_ = _dafny.CodePoint('t')
                d_757_v23_: C1
                nw139_ = C1()
                nw139_.ctor__(d_756_v22_)
                d_757_v23_ = nw139_
                d_758_v24_: _dafny.Seq
                d_758_v24_ = _dafny.SeqWithoutIsStrInference([(d_757_v23_ if d_755_v21_ else d_757_v23_), d_757_v23_, d_757_v23_, (d_757_v23_ if d_755_v21_ else d_757_v23_), d_757_v23_])
                d_758_v24_ = (d_758_v24_) + (d_758_v24_)
                d_759_v25_: C0
                nw140_ = C0()
                nw140_.ctor__(len((self).f19), (self).f18)
                d_759_v25_ = nw140_
                d_760_v26_: D4
                d_760_v26_ = D4_DC9(d_759_v25_)
                d_761_v27_: D4
                d_761_v27_ = D4_DC11(d_760_v26_)
                d_762_v28_: D4
                d_762_v28_ = D4_DC11(d_761_v27_)
                d_763_v29_: _dafny.Map
                d_763_v29_ = _dafny.Map({d_755_v21_: d_762_v28_})
                d_764_v30_: _dafny.Set
                d_764_v30_ = _dafny.Set({d_763_v29_})
                d_765_v31_: D7
                d_765_v31_ = D7_DC18(d_764_v30_)
                d_764_v30_ = (d_765_v31_).cf37
                d_766_v32_: int
                d_767_v33_: _dafny.MultiSet
                out26_: int
                out27_: _dafny.MultiSet
                out26_, out27_ = (d_757_v23_).m3(((self).f18) == ((self).f18), globalState)
                d_766_v32_ = out26_
                d_767_v33_ = out27_
                d_768_v34_: C2
                nw141_ = C2()
                nw141_.ctor__()
                d_768_v34_ = nw141_
                d_769_v35_: _dafny.Array
                nw142_ = _dafny.Array(None, 13)
                nw142_[int(0)] = d_768_v34_
                nw142_[int(1)] = d_768_v34_
                nw142_[int(2)] = d_768_v34_
                nw142_[int(3)] = d_768_v34_
                nw142_[int(4)] = (D8_DC22(d_768_v34_)).cf45
                nw142_[int(5)] = d_768_v34_
                nw142_[int(6)] = (d_768_v34_ if d_755_v21_ else d_768_v34_)
                nw142_[int(7)] = d_768_v34_
                nw142_[int(8)] = d_768_v34_
                nw142_[int(9)] = d_768_v34_
                nw142_[int(10)] = d_768_v34_
                nw142_[int(11)] = d_768_v34_
                nw142_[int(12)] = d_768_v34_
                d_769_v35_ = nw142_
                index104_ = default__.safeIndex(400, (d_769_v35_).length(0))
                (d_769_v35_)[index104_] = d_768_v34_
                index105_ = default__.safeIndex(400, (d_769_v35_).length(0))
                (d_769_v35_)[index105_] = d_768_v34_
            d_770_v36_: _dafny.Array
            nw143_ = _dafny.Array(_dafny.Set({}), 12)
            d_770_v36_ = nw143_
            d_771_v37_: bool
            d_771_v37_ = False
            d_772_v38_: _dafny.Map
            d_772_v38_ = _dafny.Map({len((self).f19): d_771_v37_})
            d_773_v40_: _dafny.Set
            d_773_v40_ = _dafny.Set({(0) - ((self).f18), d_726_i0_})
            d_774_v41_: _dafny.Seq
            def iife77_():
                coll19_ = _dafny.Map()
                compr_19_: int
                for compr_19_ in (d_773_v40_).Elements:
                    d_775_v39_: int = compr_19_
                    if (d_775_v39_) in (d_773_v40_):
                        coll19_[default__.safeDivisionInt(d_775_v39_, d_726_i0_)] = d_771_v37_
                return _dafny.Map(coll19_)
            d_774_v41_ = _dafny.SeqWithoutIsStrInference([d_772_v38_, d_772_v38_, iife77_()
            ])
            d_776_v42_: _dafny.Map
            d_776_v42_ = _dafny.Map({d_771_v37_: (self).fm4(len(d_774_v41_), (self).f19, d_725_v0_, globalState)})
            d_777_v43_: _dafny.Array
            nw144_ = _dafny.Array(None, 6)
            d_777_v43_ = nw144_
            d_778_v44_: _dafny.MultiSet
            d_778_v44_ = _dafny.MultiSet([d_777_v43_, d_777_v43_, d_777_v43_, d_777_v43_, d_777_v43_])
            d_779_v45_: _dafny.MultiSet
            d_779_v45_ = _dafny.MultiSet([len(d_776_v42_), ((d_778_v44_)[d_777_v43_] if (d_777_v43_) in (d_778_v44_) else (self).f18)])
            d_780_v46_: D2
            d_780_v46_ = D2_DC5(d_771_v37_, (0) - ((self).f18), d_736_v8_, d_779_v45_)
            d_781_v47_: _dafny.Set
            d_781_v47_ = _dafny.Set({D2_DC5(d_771_v37_, d_726_i0_, d_736_v8_, d_779_v45_), d_780_v46_})
            index106_ = default__.safeIndex(807, (d_770_v36_).length(0))
            (d_770_v36_)[index106_] = d_781_v47_
            index107_ = default__.safeIndex(807, (d_770_v36_).length(0))
            (d_770_v36_)[index107_] = d_781_v47_
            d_782_v48_: T0
            nw145_ = C0()
            nw145_.ctor__((self).f17, (self).f18)
            d_782_v48_ = nw145_
            d_783_v49_: _dafny.Map
            d_783_v49_ = _dafny.Map({d_782_v48_: d_771_v37_})
            (globalState).f3 = ((d_783_v49_) | (d_783_v49_)) != ((d_783_v49_) | (_dafny.Map({d_782_v48_: d_771_v37_})))
        hi4_ = -109
        for d_784_i1_ in range((self).f18, hi4_):
            d_785_v50_: bool
            d_785_v50_ = False
            (globalState).f2 = (((self).fm3(d_785_v50_, (self).f18, globalState) if d_785_v50_ else (self).f18)) + ((self).f18)
            d_786_v51_: str
            d_786_v51_ = _dafny.CodePoint('x')
            d_786_v51_ = d_786_v51_
            (globalState).f9 = (d_784_i1_) + ((0) - ((d_725_v0_)[default__.safeIndex(len((self).f19), len(d_725_v0_))]))
            d_785_v50_ = d_785_v50_
        d_787_v52_: _dafny.Array
        def lambda28_(d_788_i2_):
            return (d_788_i2_) + (-107)

        init16_ = lambda28_
        nw146_ = _dafny.Array(None, 25)
        for i0_16_ in range(nw146_.length(0)):
            nw146_[i0_16_] = init16_(i0_16_)
        d_787_v52_ = nw146_
        index108_ = default__.safeIndex(497, (d_787_v52_).length(0))
        (d_787_v52_)[index108_] = (self).f18
        index109_ = default__.safeIndex(497, (d_787_v52_).length(0))
        (d_787_v52_)[index109_] = ((self).f18) - ((0) - (default__.fm0(globalState)))
        d_789_v53_: _dafny.Array
        nw147_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
        d_789_v53_ = nw147_
        d_790_v54_: _dafny.Array
        def lambda29_(d_791_i3_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "alyvhe"))) != ((self).f19)

        init17_ = lambda29_
        nw148_ = _dafny.Array(None, 16)
        for i0_17_ in range(nw148_.length(0)):
            nw148_[i0_17_] = init17_(i0_17_)
        d_790_v54_ = nw148_
        d_792_v55_: bool
        d_792_v55_ = True
        index110_ = default__.safeIndex(961, (d_790_v54_).length(0))
        (d_790_v54_)[index110_] = d_792_v55_
        d_793_v56_: _dafny.MultiSet
        d_793_v56_ = _dafny.MultiSet([len(_dafny.Map({len((self).f19): d_792_v55_}))])
        index111_ = default__.safeIndex(961, (d_790_v54_).length(0))
        rhs128_ = d_789_v53_
        rhs129_ = (d_793_v56_).issubset(d_793_v56_)
        lhs96_ = d_790_v54_
        lhs97_ = default__.safeIndex(961, (d_790_v54_).length(0))
        d_789_v53_ = rhs128_
        lhs96_[lhs97_] = rhs129_
        d_794_v57_: _dafny.Seq
        d_794_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "baxh"))
        d_795_v58_: _dafny.MultiSet
        d_795_v58_ = _dafny.MultiSet([d_787_v52_])
        rhs130_ = (self).f19
        rhs131_ = (self).f18
        rhs132_ = d_795_v58_
        rhs133_ = ((d_790_v54_)[default__.safeIndex(961, (d_790_v54_).length(0))] if (d_790_v54_)[default__.safeIndex(961, (d_790_v54_).length(0))] else (d_790_v54_)[default__.safeIndex(961, (d_790_v54_).length(0))])
        lhs98_ = globalState
        lhs99_ = globalState
        d_794_v57_ = rhs130_
        lhs98_.f9 = rhs131_
        d_795_v58_ = rhs132_
        lhs99_.f0 = rhs133_
        d_796_i4_: int
        d_796_i4_ = 0
        with _dafny.label("3"):
            while d_792_v55_:
                with _dafny.c_label("3"):
                    if (d_796_i4_) >= (100):
                        raise _dafny.Break("3")
                    d_796_i4_ = (d_796_i4_) + (1)
                    if d_792_v55_:
                        d_797_v59_: _dafny.Map
                        d_797_v59_ = _dafny.Map({(self).f18: (d_790_v54_)[default__.safeIndex(961, (d_790_v54_).length(0))]})
                        d_798_v60_: int
                        d_799_v61_: int
                        out28_: int
                        out29_: int
                        out28_, out29_ = default__.m0(d_792_v55_, d_797_v59_, globalState)
                        d_798_v60_ = out28_
                        d_799_v61_ = out29_
                        d_800_v63_: _dafny.Set
                        d_800_v63_ = _dafny.Set({d_797_v59_})
                        d_801_v64_: D6
                        def iife78_():
                            coll20_ = _dafny.Map()
                            compr_20_: int
                            for compr_20_ in _dafny.IntegerRange(-322, 616):
                                d_802_v62_: int = compr_20_
                                if ((-322) <= (d_802_v62_)) and ((d_802_v62_) < (616)):
                                    coll20_[(d_802_v62_) + ((d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))])] = (0) - (d_799_v61_)
                            return _dafny.Map(coll20_)
                        d_801_v64_ = D6_DC15(len(iife78_()
), d_799_v61_, _dafny.SeqWithoutIsStrInference([(self).f18]), (self).f18, _dafny.Set({d_798_v60_, len(d_800_v63_), (0) - ((d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))]), d_798_v60_, d_799_v61_}))
                        d_803_v65_: str
                        d_803_v65_ = _dafny.CodePoint('c')
                        d_804_v66_: _dafny.Map
                        d_804_v66_ = _dafny.Map({(d_790_v54_)[default__.safeIndex(961, (d_790_v54_).length(0))]: (self).f18})
                        rhs134_ = d_801_v64_
                        rhs135_ = ((d_804_v66_) | (d_804_v66_)) != ((d_804_v66_) | (d_804_v66_))
                        rhs136_ = d_803_v65_
                        rhs137_ = (d_799_v61_) * ((d_798_v60_ if (d_790_v54_)[default__.safeIndex(961, (d_790_v54_).length(0))] else d_798_v60_))
                        rhs138_ = default__.fm1(default__.fm1((d_790_v54_)[default__.safeIndex(961, (d_790_v54_).length(0))], d_794_v57_, globalState), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))).set(default__.safeIndex((d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))), d_803_v65_), globalState)
                        lhs100_ = globalState
                        lhs101_ = globalState
                        d_801_v64_ = rhs134_
                        lhs100_.f10 = rhs135_
                        d_803_v65_ = rhs136_
                        d_798_v60_ = rhs137_
                        lhs101_.f3 = rhs138_
                        d_805_v67_: _dafny.Array
                        nw149_ = _dafny.Array(_dafny.Map({}), 1)
                        d_805_v67_ = nw149_
                        d_806_v68_: C1
                        nw150_ = C1()
                        nw150_.ctor__(d_803_v65_)
                        d_806_v68_ = nw150_
                        d_807_v69_: _dafny.Map
                        d_807_v69_ = _dafny.Map({d_806_v68_: (d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))]})
                        index112_ = default__.safeIndex(264, (d_805_v67_).length(0))
                        (d_805_v67_)[index112_] = d_807_v69_
                        d_808_v70_: D1
                        d_808_v70_ = D1_DC1(d_790_v54_)
                        d_809_v71_: _dafny.Seq
                        d_809_v71_ = _dafny.SeqWithoutIsStrInference([d_808_v70_])
                        index113_ = default__.safeIndex(961, (d_790_v54_).length(0))
                        index114_ = default__.safeIndex(264, (d_805_v67_).length(0))
                        rhs139_ = (d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))]
                        rhs140_ = (default__.fm0(globalState)) * (((d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))]) * (len(d_809_v71_)))
                        rhs141_ = (d_790_v54_)[default__.safeIndex(961, (d_790_v54_).length(0))]
                        rhs142_ = (_dafny.Map({d_806_v68_: 670})) | ((_dafny.Map({d_806_v68_: 143})) | (d_807_v69_))
                        rhs143_ = (d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))]
                        lhs102_ = globalState
                        lhs103_ = d_790_v54_
                        lhs104_ = default__.safeIndex(961, (d_790_v54_).length(0))
                        lhs105_ = d_805_v67_
                        lhs106_ = default__.safeIndex(264, (d_805_v67_).length(0))
                        lhs107_ = globalState
                        lhs102_.f9 = rhs139_
                        d_799_v61_ = rhs140_
                        lhs103_[lhs104_] = rhs141_
                        lhs105_[lhs106_] = rhs142_
                        lhs107_.f2 = rhs143_
                        (globalState).f3 = (d_799_v61_) <= ((d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))])
                        (globalState).f9 = len((self).f19)
                    elif True:
                        (globalState).f2 = default__.safeModuloInt(-132, (self).f18)
                        (globalState).f2 = (self).f18
                        d_794_v57_ = (self).f19
                        d_725_v0_ = d_725_v0_
                        d_810_v72_: _dafny.Map
                        d_810_v72_ = _dafny.Map({(self).f18: default__.safeDivisionInt(len((self).f19), (d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))])})
                        d_810_v72_ = (d_810_v72_).set((d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))], 55)
                    (globalState).f9 = ((self).f18) - (((self).fm3(d_792_v55_, (self).f18, globalState) if d_792_v55_ else 850))
                    d_811_v73_: _dafny.Seq
                    d_811_v73_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_812_i6_ in range(default__.abs(547))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bj")), d_794_v57_, d_794_v57_, d_794_v57_])
                    d_813_v74_: _dafny.Seq
                    d_813_v74_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_814_i5_ in range(default__.abs(564))]), (self).f19, (self).f19, (d_811_v73_)[default__.safeIndex((self).f18, len(d_811_v73_))]])
                    d_815_v75_: _dafny.Seq
                    d_815_v75_ = _dafny.SeqWithoutIsStrInference([(d_790_v54_)[default__.safeIndex(961, (d_790_v54_).length(0))]])
                    index115_ = default__.safeIndex(497, (d_787_v52_).length(0))
                    index116_ = default__.safeIndex(497, (d_787_v52_).length(0))
                    rhs144_ = (self).f18
                    rhs145_ = default__.safeDivisionInt(len((d_811_v73_)[default__.safeIndex((self).fm3((d_815_v75_)[default__.safeIndex((d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))], len(d_815_v75_))], len(d_815_v75_), globalState), len(d_811_v73_))]), (d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))])
                    lhs108_ = d_787_v52_
                    lhs109_ = default__.safeIndex(497, (d_787_v52_).length(0))
                    lhs110_ = d_787_v52_
                    lhs111_ = default__.safeIndex(497, (d_787_v52_).length(0))
                    lhs108_[lhs109_] = rhs144_
                    lhs110_[lhs111_] = rhs145_
                    (globalState).f9 = ((d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))]) * ((0) - ((d_787_v52_)[default__.safeIndex(497, (d_787_v52_).length(0))]))
                    pass
            pass
        r0 = d_787_v52_
        r1 = (d_790_v54_)[default__.safeIndex(961, (d_790_v54_).length(0))]
        return r0, r1

    @property
    def f19(self):
        return self._f19
