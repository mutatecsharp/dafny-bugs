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
        return -249

    @staticmethod
    def fm1(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([(False) or (not(True)), (252) != (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_0_i0_ in range(default__.abs(-707))]))), (len(_dafny.Map({False: _dafny.CodePoint('m')}))) >= (176)])

    @staticmethod
    def fm2(p0, globalState):
        return (D0_DC1((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([594])) for d_1_i0_ in range(default__.abs(961))])), len(_dafny.Set({len(_dafny.Map({_dafny.CodePoint('j'): False}))}))])).cardinality, -495, False)).cf3

    @staticmethod
    def fm5(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('g')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('u'), _dafny.CodePoint('r'), _dafny.CodePoint('f'), _dafny.CodePoint('u')]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2_i0_ in range(default__.abs(582))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i')])))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        return _dafny.Map({(_dafny.Map({True: False})) | (_dafny.Map({not(False): not(not(True))})): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "raqnsyr"))})

    @staticmethod
    def fm7(p0, p1, globalState):
        return _dafny.CodePoint('m')

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        source0_ = _dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ixyg"))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auue")))}), _dafny.Map({len(_dafny.Set({_dafny.CodePoint('c')})): (_dafny.MultiSet([True, False, False, False, (D0_DC2(False, not(True), 362, True)).cf5])).cardinality})})
        d_3___mcc_h0_ = source0_
        d_4_cf35_ = d_3___mcc_h0_
        return _dafny.Map({not(True): _dafny.SeqWithoutIsStrInference([124, len(_dafny.Map({not(True): len(_dafny.SeqWithoutIsStrInference([True]))}))])})

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: _dafny.Seq
            for compr_0_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_5_i0_ in range(default__.abs(331))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwyvjb"))})).Elements:
                d_6_v0_: _dafny.Seq = compr_0_
                if (d_6_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_5_i0_ in range(default__.abs(331))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwyvjb"))})):
                    coll0_ = coll0_.union(_dafny.Set([d_6_v0_]))
            return _dafny.Set(coll0_)
        return (iife0_()
        )

    @staticmethod
    def fm10(globalState):
        return _dafny.Map({default__.safeModuloInt((0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-230 for d_7_i0_ in range(default__.abs(564))]))).cardinality), -762): ((0) - (len(_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kl"))), 977, (0) - ((_dafny.MultiSet([-882])).cardinality), 983]))})))) + (993)})

    @staticmethod
    def fm11(p0, globalState):
        source1_ = D2_DC8()
        if source1_.is_DC8:
            return (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False])), 116, -258, 647, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hklxydp"))))})) - (_dafny.Set({len(_dafny.Map({len((D9_DC27(_dafny.Set({673}))).cf37): True}))}))
        elif source1_.is_DC9:
            d_8___mcc_h0_ = source1_.cf13
            d_9___mcc_h1_ = source1_.cf14
            d_10___mcc_h2_ = source1_.cf15
            d_11_cf15_ = d_10___mcc_h2_
            d_12_cf14_ = d_9___mcc_h1_
            d_13_cf13_ = d_8___mcc_h0_
            return (_dafny.Set({d_11_cf15_.f19, d_13_cf13_, d_11_cf15_.f19, (0) - (d_12_cf14_.f19)})) | (_dafny.Set({d_11_cf15_.f19, d_11_cf15_.f19}))
        elif source1_.is_DC7:
            d_14___mcc_h3_ = source1_.cf12
            d_15_cf12_ = d_14___mcc_h3_
            def iife1_():
                coll1_ = _dafny.Set()
                compr_1_: int
                for compr_1_ in (_dafny.SeqWithoutIsStrInference([631 for d_16_i0_ in range(default__.abs(158))])).Elements:
                    d_17_v0_: int = compr_1_
                    if (d_17_v0_) in (_dafny.SeqWithoutIsStrInference([631 for d_16_i0_ in range(default__.abs(158))])):
                        coll1_ = coll1_.union(_dafny.Set([default__.safeDivisionInt(d_17_v0_, len(_dafny.Set({True, True, False})))]))
                return _dafny.Set(coll1_)
            return (iife1_()
            ).intersection(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxglsvqlj")))}))
        elif True:
            d_18___mcc_h4_ = source1_.cf16
            d_19_cf16_ = d_18___mcc_h4_
            return _dafny.Set({134, 321})

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return D0_DC1((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erpx")))), 36, (False) and (True))

    @staticmethod
    def fm13(p0, globalState):
        return D5_DC19((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpwtnk"))) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_20_i0_ in range(default__.abs(116))])))

    @staticmethod
    def fm14(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([870])

    @staticmethod
    def fm15(globalState):
        return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, not(False), not(True), True, not(False)])]))])

    @staticmethod
    def fm16(globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in (_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "slbcnr"))): not(False)})): 953})).keys.Elements:
                d_21_v0_: int = compr_2_
                if (d_21_v0_) in (_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "slbcnr"))): not(False)})): 953})):
                    coll2_[(d_21_v0_) + (len(_dafny.SeqWithoutIsStrInference([True, False])))] = len(_dafny.Set({-705, len(_dafny.SeqWithoutIsStrInference([True]))}))
            return _dafny.Map(coll2_)
        return _dafny.Set({(iife2_()
        ) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rihcohcjh"))): len(_dafny.SeqWithoutIsStrInference([True, not(True)]))}))})

    @staticmethod
    def fm17(globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: _dafny.Seq
            for compr_3_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([not(False)]), _dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([True, True]), _dafny.SeqWithoutIsStrInference([True])})).Elements:
                d_23_v0_: _dafny.Seq = compr_3_
                if (d_23_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([not(False)]), _dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([True, True]), _dafny.SeqWithoutIsStrInference([True])})):
                    coll3_[d_23_v0_] = D0_DC0(566)
            return _dafny.Map(coll3_)
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: _dafny.Seq
            for compr_4_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])])).Elements:
                d_24_v1_: _dafny.Seq = compr_4_
                if (d_24_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])])):
                    coll4_[d_24_v1_] = D0_DC0(421)
            return _dafny.Map(coll4_)
        return _dafny.Set({True, (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_22_i0_ in range(default__.abs(-453))]))) != (70), (iife3_()
        ) == (iife4_()
        ), (_dafny.MultiSet([True, False, not(False)])).isdisjoint(_dafny.MultiSet([False]))})

    @staticmethod
    def fm18(p0, p1, globalState):
        source2_ = (_dafny.Set({_dafny.Map({722: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))})), -185]))})}) if False else _dafny.Set({_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_25_i0_ in range(default__.abs(734))])): False})): 265})}))
        d_26___mcc_h0_ = source2_
        d_27_cf35_ = d_26___mcc_h0_
        return D6_DC23((D0_DC2(not(True), not(True), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmttkoul"))), False)).cf4, 136, 541)

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: int = int(0)
        r3: _dafny.MultiSet = _dafny.MultiSet({})
        d_28_v0_: _dafny.MultiSet
        d_28_v0_ = _dafny.MultiSet([p0, p0, False])
        d_29_v1_: _dafny.Seq
        d_29_v1_ = _dafny.SeqWithoutIsStrInference([d_28_v0_])
        d_30_v2_: _dafny.Seq
        d_30_v2_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_28_v0_, d_28_v0_])) + (d_29_v1_), (d_29_v1_) + (d_29_v1_)])
        d_31_v3_: _dafny.Seq
        d_31_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rg"))
        pat_let_tv0_ = p2
        pat_let_tv1_ = p2
        def lambda0_(source3_):
            if source3_.is_DC15:
                return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))
            elif source3_.is_DC14:
                d_32___mcc_h0_ = source3_.cf21
                d_33_cf21_ = d_32___mcc_h0_
                return pat_let_tv0_
            elif True:
                d_34___mcc_h1_ = source3_.cf22
                d_35_cf22_ = d_34___mcc_h1_
                return pat_let_tv1_

        rhs0_ = len((d_30_v2_)[default__.safeIndex(p2, len(d_30_v2_))])
        rhs1_ = d_31_v3_
        rhs2_ = lambda0_(D4_DC15())
        rhs3_ = p1
        lhs0_ = globalState
        lhs1_ = globalState
        lhs2_ = globalState
        lhs0_.f12 = rhs0_
        r1 = rhs1_
        lhs1_.f2 = rhs2_
        lhs2_.f3 = rhs3_
        d_36_v4_: str
        d_36_v4_ = _dafny.CodePoint('b')
        d_37_v5_: D5
        d_37_v5_ = D5_DC17(d_36_v4_)
        pat_let_tv2_ = p2
        pat_let_tv3_ = p2
        pat_let_tv4_ = p2
        pat_let_tv5_ = p2
        def lambda1_(source4_):
            if source4_.is_DC18:
                d_38___mcc_h2_ = source4_.cf24
                d_39___mcc_h3_ = source4_.cf25
                d_40___mcc_h4_ = source4_.cf26
                d_41___mcc_h5_ = source4_.cf27
                d_42_cf27_ = d_41___mcc_h5_
                d_43_cf26_ = d_40___mcc_h4_
                d_44_cf25_ = d_39___mcc_h3_
                d_45_cf24_ = d_38___mcc_h2_
                return pat_let_tv2_
            elif source4_.is_DC19:
                d_46___mcc_h6_ = source4_.cf28
                d_47_cf28_ = d_46___mcc_h6_
                return pat_let_tv3_
            elif source4_.is_DC17:
                d_48___mcc_h7_ = source4_.cf23
                d_49_cf23_ = d_48___mcc_h7_
                return pat_let_tv4_
            elif True:
                d_50___mcc_h8_ = source4_.cf29
                d_51_cf29_ = d_50___mcc_h8_
                return (0) - (pat_let_tv5_)

        (globalState).f2 = lambda1_(d_37_v5_)
        if (p1) and (p1):
            d_52_v6_: _dafny.Array
            def lambda2_(d_53_p2_):
                def lambda3_(d_54_i0_):
                    return default__.safeModuloInt(d_54_i0_, d_53_p2_)

                return lambda3_

            init0_ = lambda2_(p2)
            nw0_ = _dafny.Array(None, 6)
            for i0_0_ in range(nw0_.length(0)):
                nw0_[i0_0_] = init0_(i0_0_)
            d_52_v6_ = nw0_
            d_55_v7_: _dafny.Seq
            d_55_v7_ = _dafny.SeqWithoutIsStrInference([p0])
            d_56_v8_: _dafny.Map
            d_56_v8_ = _dafny.Map({p2: d_55_v7_})
            index0_ = default__.safeIndex(181, (d_52_v6_).length(0))
            (d_52_v6_)[index0_] = len(d_56_v8_)
            d_57_v9_: _dafny.Seq
            d_57_v9_ = _dafny.SeqWithoutIsStrInference([p2])
            index1_ = default__.safeIndex(277, (d_52_v6_).length(0))
            (d_52_v6_)[index1_] = default__.safeDivisionInt(len(d_57_v9_), p2)
            d_58_v10_: _dafny.Seq
            d_58_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p2 for d_59_i2_ in range(default__.abs(172))])])
            d_60_v11_: _dafny.Seq
            d_60_v11_ = _dafny.SeqWithoutIsStrInference([d_58_v10_, d_58_v10_])
            d_61_v12_: _dafny.Map
            d_61_v12_ = _dafny.Map({p3: -913})
            index2_ = default__.safeIndex(181, (d_52_v6_).length(0))
            index3_ = default__.safeIndex(277, (d_52_v6_).length(0))
            rhs4_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_62_i1_ in range(default__.abs(812))])) + (d_31_v3_)
            rhs5_ = (p2) not in (d_57_v9_)
            rhs6_ = (len((d_60_v11_)[default__.safeIndex((0) - (p2), len(d_60_v11_))])) - ((0) - (default__.safeDivisionInt(len(d_61_v12_), (d_57_v9_)[default__.safeIndex(p2, len(d_57_v9_))])))
            rhs7_ = p2
            lhs3_ = globalState
            lhs4_ = d_52_v6_
            lhs5_ = default__.safeIndex(181, (d_52_v6_).length(0))
            lhs6_ = d_52_v6_
            lhs7_ = default__.safeIndex(277, (d_52_v6_).length(0))
            r1 = rhs4_
            lhs3_.f3 = rhs5_
            lhs4_[lhs5_] = rhs6_
            lhs6_[lhs7_] = rhs7_
            d_36_v4_ = d_36_v4_
            (globalState).f3 = p0
            index4_ = default__.safeIndex(181, (d_52_v6_).length(0))
            (d_52_v6_)[index4_] = default__.safeModuloInt(default__.safeModuloInt(p2, len(d_31_v3_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgihoex"))))
            d_63_v13_: _dafny.Array
            nw1_ = _dafny.Array(None, 1)
            nw1_[int(0)] = p1
            d_63_v13_ = nw1_
            index5_ = default__.safeIndex(689, (d_63_v13_).length(0))
            (d_63_v13_)[index5_] = default__.fm2(p0, globalState)
            index6_ = default__.safeIndex(689, (d_63_v13_).length(0))
            (d_63_v13_)[index6_] = default__.fm2((d_55_v7_)[default__.safeIndex((d_52_v6_)[default__.safeIndex(181, (d_52_v6_).length(0))], len(d_55_v7_))], globalState)
        elif True:
            d_64_v14_: _dafny.Map
            d_64_v14_ = _dafny.Map({False: p3})
            d_64_v14_ = (d_64_v14_).set(p3, p3)
            d_65_v15_: C0
            nw2_ = C0()
            nw2_.ctor__(536)
            d_65_v15_ = nw2_
            d_66_v16_: _dafny.Seq
            d_66_v16_ = _dafny.SeqWithoutIsStrInference([d_65_v15_])
            d_67_v17_: D4
            d_67_v17_ = D4_DC14(d_66_v16_)
            d_68_v18_: _dafny.Map
            d_68_v18_ = _dafny.Map({D6_DC22(): d_67_v17_})
            d_69_v19_: _dafny.Array
            nw3_ = _dafny.Array(None, 3)
            nw3_[int(0)] = d_68_v18_
            nw3_[int(1)] = (d_68_v18_) | (d_68_v18_)
            nw3_[int(2)] = d_68_v18_
            d_69_v19_ = nw3_
            d_70_v20_: D6
            d_70_v20_ = D6_DC22()
            index7_ = default__.safeIndex(387, (d_69_v19_).length(0))
            (d_69_v19_)[index7_] = _dafny.Map({d_70_v20_: d_67_v17_})
            index8_ = default__.safeIndex(387, (d_69_v19_).length(0))
            (d_69_v19_)[index8_] = d_68_v18_
            d_71_v21_: _dafny.Array
            nw4_ = _dafny.Array(False, 25)
            d_71_v21_ = nw4_
            index9_ = default__.safeIndex(813, (d_71_v21_).length(0))
            (d_71_v21_)[index9_] = p1
            index10_ = default__.safeIndex(813, (d_71_v21_).length(0))
            (d_71_v21_)[index10_] = p1
            d_72_v22_: _dafny.Set
            d_72_v22_ = _dafny.Set({p0})
            index11_ = default__.safeIndex(813, (d_71_v21_).length(0))
            (d_71_v21_)[index11_] = (len((default__.fm17(globalState)).intersection(d_72_v22_))) >= ((d_65_v15_.f19 if p0 else (0) - (d_65_v15_.f19)))
            d_73_v23_: _dafny.Seq
            d_73_v23_ = _dafny.SeqWithoutIsStrInference([d_65_v15_.f19])
            d_74_v24_: C0
            nw5_ = C0()
            nw5_.ctor__((d_73_v23_)[default__.safeIndex(d_65_v15_.f19, len(d_73_v23_))])
            d_74_v24_ = nw5_
        d_75_v25_: _dafny.Map
        d_75_v25_ = _dafny.Map({p2: p2})
        hi0_ = p2
        for d_76_i3_ in range(len(d_75_v25_), hi0_):
            d_77_v26_: _dafny.Array
            nw6_ = _dafny.Array(int(0), 24)
            d_77_v26_ = nw6_
            d_78_v27_: _dafny.MultiSet
            d_78_v27_ = _dafny.MultiSet([d_77_v26_, d_77_v26_, d_77_v26_])
            (globalState).f3 = ((d_78_v27_).intersection(d_78_v27_)).issubset(d_78_v27_)
            d_77_v26_ = d_77_v26_
            (globalState).f2 = p2
            d_79_v28_: _dafny.Array
            nw7_ = _dafny.Array(None, 5)
            nw7_[int(0)] = (p2) == (d_76_i3_)
            nw7_[int(1)] = p0
            nw7_[int(2)] = (default__.fm2(not(p1), globalState)) and (default__.fm2(default__.fm2(True, globalState), globalState))
            nw7_[int(3)] = p1
            nw7_[int(4)] = (p1) and (p1)
            d_79_v28_ = nw7_
            d_80_v29_: C0
            nw8_ = C0()
            nw8_.ctor__((343) * (p2))
            d_80_v29_ = nw8_
            index12_ = default__.safeIndex(761, (d_79_v28_).length(0))
            (d_79_v28_)[index12_] = (d_76_i3_) != ((0) - (d_76_i3_))
            d_81_v30_: _dafny.Seq
            d_81_v30_ = _dafny.SeqWithoutIsStrInference([p0])
            d_82_v31_: _dafny.Map
            d_82_v31_ = _dafny.Map({not(p0): (0) - (d_76_i3_)})
            d_83_v32_: _dafny.Seq
            d_83_v32_ = _dafny.SeqWithoutIsStrInference([d_76_i3_])
            index13_ = default__.safeIndex(761, (d_79_v28_).length(0))
            rhs8_ = d_79_v28_
            rhs9_ = d_80_v29_
            rhs10_ = ((_dafny.Set({(d_81_v30_)[default__.safeIndex(d_76_i3_, len(d_81_v30_))], (d_81_v30_)[default__.safeIndex(d_80_v29_.f19, len(d_81_v30_))]})) | (_dafny.Set({p1}))).ispropersubset(default__.fm17(globalState))
            rhs11_ = (d_80_v29_.f19) >= (((d_82_v31_)[p0] if (p0) in (d_82_v31_) else d_76_i3_))
            rhs12_ = ((len(d_83_v32_)) < (len(d_31_v3_))) and (p1)
            lhs8_ = d_79_v28_
            lhs9_ = default__.safeIndex(761, (d_79_v28_).length(0))
            lhs10_ = globalState
            lhs11_ = globalState
            d_79_v28_ = rhs8_
            d_80_v29_ = rhs9_
            lhs8_[lhs9_] = rhs10_
            lhs10_.f3 = rhs11_
            lhs11_.f3 = rhs12_
        pat_let_tv6_ = p3
        pat_let_tv7_ = p3
        pat_let_tv8_ = p1
        pat_let_tv9_ = p0
        pat_let_tv10_ = p1
        pat_let_tv11_ = p0
        pat_let_tv12_ = p1
        pat_let_tv13_ = p3
        pat_let_tv14_ = p0
        def lambda4_(source5_):
            if source5_.is_DC1:
                d_84___mcc_h9_ = source5_.cf1
                d_85___mcc_h10_ = source5_.cf2
                d_86___mcc_h11_ = source5_.cf3
                d_87_cf3_ = d_86___mcc_h11_
                d_88_cf2_ = d_85___mcc_h10_
                d_89_cf1_ = d_84___mcc_h9_
                return (_dafny.SeqWithoutIsStrInference([pat_let_tv6_])) == (_dafny.SeqWithoutIsStrInference([pat_let_tv7_, pat_let_tv8_, pat_let_tv9_, pat_let_tv10_]))
            elif source5_.is_DC2:
                d_90___mcc_h12_ = source5_.cf4
                d_91___mcc_h13_ = source5_.cf5
                d_92___mcc_h14_ = source5_.cf6
                d_93___mcc_h15_ = source5_.cf7
                d_94_cf7_ = d_93___mcc_h15_
                d_95_cf6_ = d_92___mcc_h14_
                d_96_cf5_ = d_91___mcc_h13_
                d_97_cf4_ = d_90___mcc_h12_
                return pat_let_tv11_
            elif True:
                d_98___mcc_h16_ = source5_.cf0
                d_99_cf0_ = d_98___mcc_h16_
                if pat_let_tv12_:
                    return pat_let_tv13_
                elif True:
                    return pat_let_tv14_

        if lambda4_(D0_DC0(-878)):
            d_100_v33_: C0
            nw9_ = C0()
            nw9_.ctor__(p2)
            d_100_v33_ = nw9_
            d_101_v34_: _dafny.Seq
            d_101_v34_ = _dafny.SeqWithoutIsStrInference([d_100_v33_])
            d_102_v35_: D4
            d_102_v35_ = D4_DC14(d_101_v34_)
            d_103_v36_: D4
            d_103_v36_ = D4_DC16(d_102_v35_)
            pat_let_tv15_ = d_102_v35_
            d_104_v37_: _dafny.Array
            nw10_ = _dafny.Array(None, 19)
            nw10_[int(0)] = d_103_v36_
            nw10_[int(1)] = d_103_v36_
            nw10_[int(2)] = d_103_v36_
            nw10_[int(3)] = d_103_v36_
            nw10_[int(4)] = d_103_v36_
            nw10_[int(5)] = d_103_v36_
            nw10_[int(6)] = d_103_v36_
            nw10_[int(7)] = d_103_v36_
            nw10_[int(8)] = d_103_v36_
            nw10_[int(9)] = d_103_v36_
            nw10_[int(10)] = d_103_v36_
            nw10_[int(11)] = d_103_v36_
            nw10_[int(12)] = d_103_v36_
            def iife5_(_pat_let0_0):
                def iife6_(d_105_dt__update__tmp_h0_):
                    def iife7_(_pat_let1_0):
                        def iife8_(d_106_dt__update_hcf22_h0_):
                            return D4_DC16(d_106_dt__update_hcf22_h0_)
                        return iife8_(_pat_let1_0)
                    return iife7_(pat_let_tv15_)
                return iife6_(_pat_let0_0)
            nw10_[int(13)] = iife5_(d_103_v36_)
            nw10_[int(14)] = d_103_v36_
            nw10_[int(15)] = d_103_v36_
            nw10_[int(16)] = D4_DC16(d_102_v35_)
            nw10_[int(17)] = D4_DC16(d_102_v35_)
            nw10_[int(18)] = d_103_v36_
            d_104_v37_ = nw10_
            index14_ = default__.safeIndex(594, (d_104_v37_).length(0))
            (d_104_v37_)[index14_] = d_103_v36_
            index15_ = default__.safeIndex(594, (d_104_v37_).length(0))
            (d_104_v37_)[index15_] = d_103_v36_
            r0 = d_100_v33_.f19
            d_107_v38_: C0
            nw11_ = C0()
            nw11_.ctor__(d_100_v33_.f19)
            d_107_v38_ = nw11_
            d_108_v39_: _dafny.Seq
            d_108_v39_ = _dafny.SeqWithoutIsStrInference([p3, default__.fm2(p3, globalState)])
            d_109_v40_: _dafny.Map
            d_109_v40_ = _dafny.Map({p3: len(default__.fm14(p1, globalState))})
            if (d_108_v39_)[default__.safeIndex(((d_109_v40_)[default__.fm2(p0, globalState)] if (default__.fm2(p0, globalState)) in (d_109_v40_) else -950), len(d_108_v39_))]:
                r1 = default__.fm5(len(d_31_v3_), globalState)
                rhs13_ = not((d_28_v0_).ispropersubset((d_28_v0_).set(p3, default__.abs(p2))))
                rhs14_ = (d_100_v33_ if not((384) < (d_100_v33_.f19)) else d_107_v38_)
                lhs12_ = globalState
                lhs12_.f3 = rhs13_
                d_107_v38_ = rhs14_
                d_110_v41_: _dafny.Array
                def lambda5_(d_111_p0_):
                    def lambda6_(d_112_i4_):
                        return d_111_p0_

                    return lambda6_

                init1_ = lambda5_(p0)
                nw12_ = _dafny.Array(None, 20)
                for i0_1_ in range(nw12_.length(0)):
                    nw12_[i0_1_] = init1_(i0_1_)
                d_110_v41_ = nw12_
                index16_ = default__.safeIndex(471, (d_110_v41_).length(0))
                (d_110_v41_)[index16_] = (False if p1 else p3)
                d_113_v42_: _dafny.Map
                d_113_v42_ = _dafny.Map({d_36_v4_: (d_107_v38_.f19) <= (len(d_31_v3_))})
                index17_ = default__.safeIndex(471, (d_110_v41_).length(0))
                (d_110_v41_)[index17_] = ((d_113_v42_)[d_36_v4_] if (d_36_v4_) in (d_113_v42_) else p0)
                index18_ = default__.safeIndex(471, (d_110_v41_).length(0))
                (d_110_v41_)[index18_] = not ((d_75_v25_) == (d_75_v25_)) or ((d_28_v0_).issubset((_dafny.MultiSet([(d_110_v41_)[default__.safeIndex(471, (d_110_v41_).length(0))], not(True)])).set((d_110_v41_)[default__.safeIndex(471, (d_110_v41_).length(0))], default__.abs(p2))))
                index19_ = default__.safeIndex(471, (d_110_v41_).length(0))
                rhs15_ = d_100_v33_.f19
                rhs16_ = default__.fm2((d_110_v41_)[default__.safeIndex(471, (d_110_v41_).length(0))], globalState)
                lhs13_ = globalState
                lhs14_ = d_110_v41_
                lhs15_ = default__.safeIndex(471, (d_110_v41_).length(0))
                lhs13_.f2 = rhs15_
                lhs14_[lhs15_] = rhs16_
            elif True:
                d_114_v43_: D6
                d_114_v43_ = D6_DC23(p3, d_100_v33_.f19, d_100_v33_.f19)
                d_115_v44_: _dafny.Map
                d_115_v44_ = _dafny.Map({(d_114_v43_).cf31: d_107_v38_})
                d_116_v45_: _dafny.Array
                nw13_ = _dafny.Array(False, 22)
                d_116_v45_ = nw13_
                d_117_v46_: D5
                d_117_v46_ = D5_DC18(d_107_v38_.f19, p1, d_116_v45_, p1)
                d_118_v47_: _dafny.Array
                nw14_ = _dafny.Array(None, 27)
                nw14_[int(0)] = d_107_v38_
                nw14_[int(1)] = d_107_v38_
                nw14_[int(2)] = d_107_v38_
                nw14_[int(3)] = d_100_v33_
                nw14_[int(4)] = d_107_v38_
                nw14_[int(5)] = d_107_v38_
                nw14_[int(6)] = d_100_v33_
                nw14_[int(7)] = d_107_v38_
                nw14_[int(8)] = d_107_v38_
                nw14_[int(9)] = d_107_v38_
                nw14_[int(10)] = d_107_v38_
                nw14_[int(11)] = ((d_115_v44_)[(d_117_v46_).cf25] if ((d_117_v46_).cf25) in (d_115_v44_) else d_100_v33_)
                nw14_[int(12)] = d_107_v38_
                nw14_[int(13)] = d_107_v38_
                nw14_[int(14)] = d_107_v38_
                nw14_[int(15)] = d_100_v33_
                nw14_[int(16)] = d_107_v38_
                nw14_[int(17)] = d_107_v38_
                nw14_[int(18)] = d_100_v33_
                nw14_[int(19)] = d_107_v38_
                nw14_[int(20)] = d_107_v38_
                nw14_[int(21)] = d_107_v38_
                nw14_[int(22)] = d_100_v33_
                nw14_[int(23)] = d_107_v38_
                nw14_[int(24)] = d_100_v33_
                nw14_[int(25)] = d_107_v38_
                nw14_[int(26)] = d_100_v33_
                d_118_v47_ = nw14_
                d_119_v48_: _dafny.Map
                d_119_v48_ = _dafny.Map({d_107_v38_.f19: d_118_v47_})
                d_119_v48_ = (d_119_v48_).set(d_107_v38_.f19, d_118_v47_)
                (globalState).f3 = p1
                d_120_v49_: C0
                nw15_ = C0()
                nw15_.ctor__((0) - ((0) - ((d_100_v33_).fm3(globalState))))
                d_120_v49_ = nw15_
                d_121_v50_: C0
                nw16_ = C0()
                nw16_.ctor__(d_107_v38_.f19)
                d_121_v50_ = nw16_
                (globalState).f3 = not (p1) or (p0)
            d_122_v51_: _dafny.MultiSet
            d_122_v51_ = _dafny.MultiSet([103, d_107_v38_.f19])
            d_122_v51_ = d_122_v51_
        elif True:
            d_123_v52_: _dafny.Seq
            d_123_v52_ = _dafny.SeqWithoutIsStrInference([p3])
            d_124_v53_: C0
            nw17_ = C0()
            nw17_.ctor__(default__.fm0(len(d_123_v52_), d_36_v4_, _dafny.SeqWithoutIsStrInference([d_31_v3_, d_31_v3_, d_31_v3_, d_31_v3_, _dafny.SeqWithoutIsStrInference([d_36_v4_ for d_125_i5_ in range(default__.abs(750))])]), p2, globalState))
            d_124_v53_ = nw17_
            d_126_v54_: C0
            nw18_ = C0()
            nw18_.ctor__(d_124_v53_.f19)
            d_126_v54_ = nw18_
            d_127_v55_: _dafny.Set
            d_127_v55_ = _dafny.Set({d_126_v54_.f19})
            if (d_127_v55_).isdisjoint((default__.fm11(d_126_v54_.f19, globalState)) | (d_127_v55_)):
                d_128_v56_: _dafny.Set
                d_128_v56_ = _dafny.Set({p0})
                d_129_v57_: _dafny.Map
                d_129_v57_ = _dafny.Map({d_36_v4_: len(d_128_v56_)})
                rhs17_ = d_126_v54_
                rhs18_ = ((0) - (((d_129_v57_)[d_36_v4_] if (d_36_v4_) in (d_129_v57_) else d_124_v53_.f19))) != (d_126_v54_.f19)
                rhs19_ = d_126_v54_.f19
                lhs16_ = globalState
                lhs17_ = globalState
                d_126_v54_ = rhs17_
                lhs16_.f3 = rhs18_
                lhs17_.f8 = rhs19_
                d_130_v58_: C0
                nw19_ = C0()
                nw19_.ctor__((d_126_v54_.f19) + (d_124_v53_.f19))
                d_130_v58_ = nw19_
                d_131_v59_: _dafny.Map
                d_131_v59_ = _dafny.Map({_dafny.CodePoint('c'): d_126_v54_})
                d_131_v59_ = (d_131_v59_) | (d_131_v59_)
                (globalState).f3 = ((len(d_31_v3_)) >= (-942)) and (p0)
                (globalState).f2 = d_124_v53_.f19
            elif True:
                (globalState).f8 = d_124_v53_.f19
                d_132_v60_: D6
                d_132_v60_ = D6_DC23(p1, d_124_v53_.f19, d_124_v53_.f19)
                d_133_v61_: _dafny.MultiSet
                d_133_v61_ = _dafny.MultiSet([p2, len(d_31_v3_), d_124_v53_.f19])
                d_134_v62_: _dafny.Map
                d_134_v62_ = _dafny.Map({d_126_v54_.f19: p3})
                d_135_v63_: _dafny.Map
                d_135_v63_ = _dafny.Map({((d_134_v62_)[d_124_v53_.f19] if (d_124_v53_.f19) in (d_134_v62_) else p0): p0})
                pat_let_tv16_ = d_126_v54_
                pat_let_tv17_ = d_133_v61_
                d_136_v64_: _dafny.Array
                nw20_ = _dafny.Array(None, 29)
                nw20_[int(0)] = d_132_v60_
                nw20_[int(1)] = d_132_v60_
                nw20_[int(2)] = d_132_v60_
                nw20_[int(3)] = d_132_v60_
                nw20_[int(4)] = d_132_v60_
                nw20_[int(5)] = d_132_v60_
                nw20_[int(6)] = d_132_v60_
                nw20_[int(7)] = d_132_v60_
                nw20_[int(8)] = d_132_v60_
                nw20_[int(9)] = d_132_v60_
                def iife9_(_pat_let2_0):
                    def iife10_(d_137_dt__update__tmp_h1_):
                        def iife11_(_pat_let3_0):
                            def iife12_(d_138_dt__update_hcf33_h0_):
                                return D6_DC23((d_137_dt__update__tmp_h1_).cf31, (d_137_dt__update__tmp_h1_).cf32, d_138_dt__update_hcf33_h0_)
                            return iife12_(_pat_let3_0)
                        return iife11_(pat_let_tv16_.f19)
                    return iife10_(_pat_let2_0)
                nw20_[int(10)] = iife9_(D6_DC23(p3, len(d_31_v3_), d_124_v53_.f19))
                nw20_[int(11)] = d_132_v60_
                def iife13_(_pat_let4_0):
                    def iife14_(d_139_dt__update__tmp_h2_):
                        def iife15_(_pat_let5_0):
                            def iife16_(d_140_dt__update_hcf32_h0_):
                                return D6_DC23((d_139_dt__update__tmp_h2_).cf31, d_140_dt__update_hcf32_h0_, (d_139_dt__update__tmp_h2_).cf33)
                            return iife16_(_pat_let5_0)
                        return iife15_((pat_let_tv17_).cardinality)
                    return iife14_(_pat_let4_0)
                nw20_[int(12)] = iife13_(d_132_v60_)
                nw20_[int(13)] = d_132_v60_
                nw20_[int(14)] = D6_DC23(p1, d_124_v53_.f19, d_124_v53_.f19)
                nw20_[int(15)] = D6_DC23(((d_135_v63_)[p0] if (p0) in (d_135_v63_) else p0), -328, (0) - (d_126_v54_.f19))
                nw20_[int(16)] = (d_132_v60_ if p0 else d_132_v60_)
                nw20_[int(17)] = (d_132_v60_ if False else d_132_v60_)
                nw20_[int(18)] = D6_DC23(p0, p2, d_124_v53_.f19)
                nw20_[int(19)] = d_132_v60_
                nw20_[int(20)] = d_132_v60_
                nw20_[int(21)] = d_132_v60_
                nw20_[int(22)] = d_132_v60_
                nw20_[int(23)] = d_132_v60_
                nw20_[int(24)] = d_132_v60_
                nw20_[int(25)] = (d_132_v60_ if p3 else d_132_v60_)
                nw20_[int(26)] = default__.fm18(d_126_v54_.f19, p1, globalState)
                nw20_[int(27)] = d_132_v60_
                nw20_[int(28)] = d_132_v60_
                d_136_v64_ = nw20_
                index20_ = default__.safeIndex(617, (d_136_v64_).length(0))
                (d_136_v64_)[index20_] = (d_132_v60_ if p3 else D6_DC23(p0, d_126_v54_.f19, ((d_28_v0_)[((d_134_v62_)[d_124_v53_.f19] if (d_124_v53_.f19) in (d_134_v62_) else not(p0))] if (((d_134_v62_)[d_124_v53_.f19] if (d_124_v53_.f19) in (d_134_v62_) else not(p0))) in (d_28_v0_) else d_126_v54_.f19)))
                index21_ = default__.safeIndex(617, (d_136_v64_).length(0))
                (d_136_v64_)[index21_] = default__.fm18(d_124_v53_.f19, False, globalState)
                (globalState).f3 = (d_31_v3_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hnidaeeo")))
                d_141_v65_: _dafny.Map
                d_141_v65_ = _dafny.Map({p3: (0) - (d_124_v53_.f19)})
                d_141_v65_ = (d_141_v65_).set(p1, d_124_v53_.f19)
                d_142_v66_: _dafny.Array
                nw21_ = _dafny.Array(int(0), 23)
                d_142_v66_ = nw21_
                index22_ = default__.safeIndex(348, (d_142_v66_).length(0))
                (d_142_v66_)[index22_] = d_126_v54_.f19
                index23_ = default__.safeIndex(348, (d_142_v66_).length(0))
                (d_142_v66_)[index23_] = d_126_v54_.f19
            d_143_v67_: _dafny.Array
            nw22_ = _dafny.Array(None, 27)
            d_143_v67_ = nw22_
            index24_ = default__.safeIndex(277, (d_143_v67_).length(0))
            (d_143_v67_)[index24_] = d_124_v53_
            index25_ = default__.safeIndex(277, (d_143_v67_).length(0))
            nw23_ = C0()
            nw23_.ctor__(default__.safeDivisionInt(115, (0) - (p2)))
            (d_143_v67_)[index25_] = nw23_
            d_144_v68_: _dafny.Map
            d_144_v68_ = _dafny.Map({p1: 198})
            d_145_v69_: _dafny.Seq
            d_145_v69_ = _dafny.SeqWithoutIsStrInference([d_31_v3_])
            d_144_v68_ = (d_144_v68_).set(((d_124_v53_).fm4(p2, d_126_v54_.f19, p0, p0, globalState)) != (d_126_v54_.f19), len((d_145_v69_)[default__.safeIndex(d_124_v53_.f19, len(d_145_v69_))]))
        d_146_v70_: _dafny.Seq
        d_146_v70_ = _dafny.SeqWithoutIsStrInference([False, p0])
        d_147_v71_: _dafny.Map
        d_147_v71_ = _dafny.Map({p1: p1})
        d_148_v72_: _dafny.MultiSet
        d_148_v72_ = _dafny.MultiSet([d_147_v71_, _dafny.Map({p3: p3})])
        d_149_v73_: _dafny.Map
        d_149_v73_ = _dafny.Map({p3: (d_148_v72_).cardinality})
        d_150_v74_: _dafny.Map
        d_150_v74_ = _dafny.Map({p1: (d_149_v73_) | (_dafny.Map({p0: p2}))})
        rhs20_ = d_146_v70_
        rhs21_ = (d_31_v3_) + ((d_31_v3_) + (d_31_v3_))
        rhs22_ = d_150_v74_
        rhs23_ = d_37_v5_
        d_146_v70_ = rhs20_
        r1 = rhs21_
        d_150_v74_ = rhs22_
        d_37_v5_ = rhs23_
        d_151_v75_: C0
        nw24_ = C0()
        nw24_.ctor__(p2)
        d_151_v75_ = nw24_
        r0 = (len(_dafny.SeqWithoutIsStrInference([d_151_v75_]))) * (p2)
        r1 = ((d_31_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ey")))) + (default__.fm5(d_151_v75_.f19, globalState))
        r2 = d_151_v75_.f19
        d_152_v76_: _dafny.Set
        d_152_v76_ = _dafny.Set({p0, True, p1})
        r3 = (d_28_v0_).set((p2) < (len(d_152_v76_)), default__.abs(default__.safeModuloInt(p2, p2)))
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_153_v0_: bool
        d_153_v0_ = False
        d_154_v1_: _dafny.Seq
        d_154_v1_ = _dafny.SeqWithoutIsStrInference([d_153_v0_])
        d_155_globalState_: GlobalState
        nw25_ = GlobalState()
        nw25_.ctor__(-782, False, 351, False, -139, 973, False, -608, 117, -409, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adqetglt")), False, -956, True, 26, False, 957, 182, d_154_v1_)
        d_155_globalState_ = nw25_
        d_156_v2_: int
        d_156_v2_ = 270
        d_157_v3_: _dafny.MultiSet
        d_157_v3_ = _dafny.MultiSet([d_153_v0_])
        hi1_ = (0) - ((d_156_v2_) - ((0) - ((0) - ((d_157_v3_).cardinality))))
        for d_158_i0_ in range(115, hi1_):
            d_159_v4_: str
            d_159_v4_ = _dafny.CodePoint('f')
            d_160_v5_: _dafny.Seq
            d_160_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oiwhio"))
            d_161_v6_: _dafny.Array
            nw26_ = _dafny.Array(None, 3)
            nw26_[int(0)] = d_160_v5_
            nw26_[int(1)] = d_160_v5_
            nw26_[int(2)] = d_160_v5_
            d_161_v6_ = nw26_
            d_162_v7_: _dafny.Set
            d_162_v7_ = _dafny.Set({d_153_v0_})
            d_163_v8_: _dafny.Seq
            d_163_v8_ = _dafny.SeqWithoutIsStrInference([d_160_v5_])
            d_164_v9_: _dafny.Map
            d_164_v9_ = _dafny.Map({default__.fm0(len(d_162_v7_), d_159_v4_, d_163_v8_, 591, d_155_globalState_): d_161_v6_})
            d_165_v10_: _dafny.Map
            d_165_v10_ = _dafny.Map({d_156_v2_: d_153_v0_})
            d_166_v11_: _dafny.Array
            nw27_ = _dafny.Array(None, 19)
            nw27_[int(0)] = d_161_v6_
            nw27_[int(1)] = d_161_v6_
            nw27_[int(2)] = d_161_v6_
            nw27_[int(3)] = d_161_v6_
            nw27_[int(4)] = d_161_v6_
            nw27_[int(5)] = d_161_v6_
            nw27_[int(6)] = d_161_v6_
            nw27_[int(7)] = d_161_v6_
            nw27_[int(8)] = d_161_v6_
            nw27_[int(9)] = d_161_v6_
            nw27_[int(10)] = ((d_164_v9_)[len(d_165_v10_)] if (len(d_165_v10_)) in (d_164_v9_) else d_161_v6_)
            nw27_[int(11)] = (d_161_v6_ if d_153_v0_ else d_161_v6_)
            nw27_[int(12)] = d_161_v6_
            nw27_[int(13)] = d_161_v6_
            nw27_[int(14)] = d_161_v6_
            nw27_[int(15)] = d_161_v6_
            nw27_[int(16)] = d_161_v6_
            nw27_[int(17)] = d_161_v6_
            nw27_[int(18)] = d_161_v6_
            d_166_v11_ = nw27_
            index26_ = default__.safeIndex(689, (d_166_v11_).length(0))
            (d_166_v11_)[index26_] = d_161_v6_
            d_167_v12_: _dafny.Array
            nw28_ = _dafny.Array(_dafny.Set({}), 26)
            d_167_v12_ = nw28_
            index27_ = default__.safeIndex(153, (d_167_v12_).length(0))
            (d_167_v12_)[index27_] = d_162_v7_
            d_168_v13_: _dafny.Map
            d_168_v13_ = _dafny.Map({d_156_v2_: d_162_v7_})
            index28_ = default__.safeIndex(689, (d_166_v11_).length(0))
            index29_ = default__.safeIndex(153, (d_167_v12_).length(0))
            rhs24_ = _dafny.CodePoint('r')
            rhs25_ = d_161_v6_
            rhs26_ = d_153_v0_
            rhs27_ = (d_162_v7_).intersection(((d_168_v13_)[d_158_i0_] if (d_158_i0_) in (d_168_v13_) else d_162_v7_))
            lhs18_ = d_166_v11_
            lhs19_ = default__.safeIndex(689, (d_166_v11_).length(0))
            lhs20_ = d_155_globalState_
            lhs21_ = d_167_v12_
            lhs22_ = default__.safeIndex(153, (d_167_v12_).length(0))
            d_159_v4_ = rhs24_
            lhs18_[lhs19_] = rhs25_
            lhs20_.f3 = rhs26_
            lhs21_[lhs22_] = rhs27_
            d_169_v14_: _dafny.Seq
            d_169_v14_ = _dafny.SeqWithoutIsStrInference([172])
            d_170_v15_: _dafny.Map
            d_170_v15_ = _dafny.Map({(d_169_v14_)[default__.safeIndex(-542, len(d_169_v14_))]: d_158_i0_})
            rhs28_ = d_156_v2_
            rhs29_ = (0) - ((((0) - (((d_170_v15_)[d_156_v2_] if (d_156_v2_) in (d_170_v15_) else d_156_v2_))) + (d_156_v2_)) * (d_158_i0_))
            rhs30_ = ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gydcso"))))) - (-566)
            rhs31_ = d_156_v2_
            lhs23_ = d_155_globalState_
            lhs24_ = d_155_globalState_
            lhs25_ = d_155_globalState_
            lhs26_ = d_155_globalState_
            lhs23_.f12 = rhs28_
            lhs24_.f8 = rhs29_
            lhs25_.f2 = rhs30_
            lhs26_.f12 = rhs31_
            d_171_v16_: _dafny.Array
            def lambda7_(d_172_v5_):
                def lambda8_(d_173_i1_):
                    return (_dafny.Set({d_172_v5_})) - (_dafny.Set({d_172_v5_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvf"))}))

                return lambda8_

            init2_ = lambda7_(d_160_v5_)
            nw29_ = _dafny.Array(None, 6)
            for i0_2_ in range(nw29_.length(0)):
                nw29_[i0_2_] = init2_(i0_2_)
            d_171_v16_ = nw29_
            d_174_v17_: _dafny.Set
            d_174_v17_ = _dafny.Set({d_160_v5_, d_160_v5_})
            index30_ = default__.safeIndex(397, (d_171_v16_).length(0))
            (d_171_v16_)[index30_] = d_174_v17_
            index31_ = default__.safeIndex(397, (d_171_v16_).length(0))
            (d_171_v16_)[index31_] = ((d_174_v17_).intersection(d_174_v17_)).intersection(d_174_v17_)
            (d_155_globalState_).f8 = ((0) - (d_156_v2_)) - (d_156_v2_)
        d_175_v18_: _dafny.Seq
        d_175_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcqnkipj"))
        d_175_v18_ = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('g') if True else _dafny.CodePoint('x')) for d_176_i2_ in range(default__.abs(650))])
        if d_153_v0_:
            d_177_v19_: int
            d_178_v20_: _dafny.Seq
            d_179_v21_: int
            d_180_v22_: _dafny.MultiSet
            out0_: int
            out1_: _dafny.Seq
            out2_: int
            out3_: _dafny.MultiSet
            out0_, out1_, out2_, out3_ = default__.m0(d_153_v0_, d_153_v0_, d_156_v2_, d_153_v0_, d_155_globalState_)
            d_177_v19_ = out0_
            d_178_v20_ = out1_
            d_179_v21_ = out2_
            d_180_v22_ = out3_
            d_153_v0_ = not(d_153_v0_)
            d_181_v23_: str
            d_181_v23_ = _dafny.CodePoint('q')
            d_182_v24_: _dafny.Map
            d_182_v24_ = _dafny.Map({d_156_v2_: d_181_v23_})
            d_183_v25_: _dafny.Set
            d_183_v25_ = _dafny.Set({d_182_v24_})
            d_184_v26_: D0
            d_184_v26_ = D0_DC0(d_156_v2_)
            pat_let_tv18_ = d_179_v21_
            d_185_v27_: _dafny.Array
            nw30_ = _dafny.Array(None, 5)
            nw30_[int(0)] = len(d_183_v25_)
            nw30_[int(1)] = (0) - (len(d_175_v18_))
            nw30_[int(2)] = d_177_v19_
            nw30_[int(3)] = 694
            def iife17_(_pat_let6_0):
                def iife18_(d_186_dt__update__tmp_h0_):
                    def iife19_(_pat_let7_0):
                        def iife20_(d_187_dt__update_hcf0_h0_):
                            return D0_DC0(d_187_dt__update_hcf0_h0_)
                        return iife20_(_pat_let7_0)
                    return iife19_(pat_let_tv18_)
                return iife18_(_pat_let6_0)
            nw30_[int(4)] = (iife17_(d_184_v26_)).cf0
            d_185_v27_ = nw30_
            d_188_v28_: _dafny.MultiSet
            d_188_v28_ = _dafny.MultiSet([d_185_v27_])
            d_188_v28_ = (d_188_v28_).intersection(d_188_v28_)
            d_189_v29_: _dafny.Map
            d_189_v29_ = _dafny.Map({d_156_v2_: d_180_v22_})
            d_190_v30_: D0
            d_190_v30_ = D0_DC2(d_153_v0_, (d_177_v19_) >= (d_156_v2_), d_179_v21_, True)
            rhs32_ = d_177_v19_
            rhs33_ = (d_189_v29_).set(d_177_v19_, _dafny.MultiSet([d_153_v0_]))
            rhs34_ = d_190_v30_
            lhs27_ = d_155_globalState_
            lhs27_.f2 = rhs32_
            d_189_v29_ = rhs33_
            d_190_v30_ = rhs34_
            d_191_v31_: _dafny.Array
            nw31_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_191_v31_ = nw31_
            d_192_v32_: _dafny.Seq
            d_192_v32_ = _dafny.SeqWithoutIsStrInference([d_175_v18_, d_175_v18_])
            d_193_v33_: _dafny.Array
            nw32_ = _dafny.Array(None, 23)
            nw32_[int(0)] = (_dafny.SeqWithoutIsStrInference([d_181_v23_ for d_194_i3_ in range(default__.abs(277))])).set(default__.safeIndex((0) - (default__.fm0(d_177_v19_, _dafny.CodePoint('r'), d_192_v32_, d_177_v19_, d_155_globalState_)), len(_dafny.SeqWithoutIsStrInference([d_181_v23_ for d_195_i3_ in range(default__.abs(277))]))), _dafny.CodePoint('h'))
            nw32_[int(1)] = d_175_v18_
            nw32_[int(2)] = d_175_v18_
            nw32_[int(3)] = d_178_v20_
            nw32_[int(4)] = d_178_v20_
            nw32_[int(5)] = d_175_v18_
            nw32_[int(6)] = _dafny.SeqWithoutIsStrInference([d_181_v23_ for d_196_i4_ in range(default__.abs(238))])
            nw32_[int(7)] = d_175_v18_
            nw32_[int(8)] = d_175_v18_
            nw32_[int(9)] = d_178_v20_
            nw32_[int(10)] = d_175_v18_
            nw32_[int(11)] = d_178_v20_
            nw32_[int(12)] = d_178_v20_
            nw32_[int(13)] = d_175_v18_
            nw32_[int(14)] = d_178_v20_
            nw32_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xeumplyd"))
            nw32_[int(16)] = d_175_v18_
            nw32_[int(17)] = _dafny.SeqWithoutIsStrInference([d_181_v23_ for d_197_i5_ in range(default__.abs(550))])
            nw32_[int(18)] = d_175_v18_
            nw32_[int(19)] = d_175_v18_
            nw32_[int(20)] = _dafny.SeqWithoutIsStrInference([d_181_v23_ for d_198_i6_ in range(default__.abs(885))])
            nw32_[int(21)] = d_175_v18_
            nw32_[int(22)] = d_178_v20_
            d_193_v33_ = nw32_
            index32_ = default__.safeIndex(717, (d_191_v31_).length(0))
            (d_191_v31_)[index32_] = d_193_v33_
            d_199_v34_: _dafny.MultiSet
            d_199_v34_ = _dafny.MultiSet([d_184_v26_])
            d_200_v35_: _dafny.MultiSet
            d_200_v35_ = _dafny.MultiSet([d_177_v19_, d_179_v21_])
            index33_ = default__.safeIndex(379, (d_185_v27_).length(0))
            (d_185_v27_)[index33_] = default__.safeModuloInt(((d_199_v34_).set(d_184_v26_, default__.abs((d_180_v22_).cardinality))).cardinality, default__.fm0(d_156_v2_, d_181_v23_, d_192_v32_, (d_200_v35_).cardinality, d_155_globalState_))
            index34_ = default__.safeIndex(717, (d_191_v31_).length(0))
            index35_ = default__.safeIndex(379, (d_185_v27_).length(0))
            rhs35_ = d_185_v27_
            rhs36_ = d_156_v2_
            rhs37_ = d_193_v33_
            rhs38_ = (0) - (d_156_v2_)
            lhs28_ = d_155_globalState_
            lhs29_ = d_191_v31_
            lhs30_ = default__.safeIndex(717, (d_191_v31_).length(0))
            lhs31_ = d_185_v27_
            lhs32_ = default__.safeIndex(379, (d_185_v27_).length(0))
            d_185_v27_ = rhs35_
            lhs28_.f0 = rhs36_
            lhs29_[lhs30_] = rhs37_
            lhs31_[lhs32_] = rhs38_
        elif True:
            d_201_v36_: str
            d_201_v36_ = _dafny.CodePoint('s')
            d_154_v1_ = default__.fm1(d_201_v36_, d_153_v0_, d_155_globalState_)
            if default__.fm2((d_153_v0_) == (d_153_v0_), d_155_globalState_):
                d_202_v37_: C0
                nw33_ = C0()
                nw33_.ctor__((d_156_v2_) + ((0) - (d_156_v2_)))
                d_202_v37_ = nw33_
                d_203_v38_: _dafny.Seq
                d_203_v38_ = _dafny.SeqWithoutIsStrInference([d_175_v18_])
                d_204_v39_: _dafny.Set
                d_204_v39_ = _dafny.Set({d_153_v0_})
                (d_155_globalState_).f0 = default__.fm0((d_202_v37_.f19) * (d_202_v37_.f19), d_201_v36_, (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_201_v36_]), d_175_v18_])) + (d_203_v38_), len((d_204_v39_).intersection(d_204_v39_)), d_155_globalState_)
                d_205_v40_: int
                d_206_v41_: _dafny.Seq
                d_207_v42_: int
                d_208_v43_: _dafny.MultiSet
                out4_: int
                out5_: _dafny.Seq
                out6_: int
                out7_: _dafny.MultiSet
                out4_, out5_, out6_, out7_ = default__.m0(d_153_v0_, (d_175_v18_) <= (d_175_v18_), d_202_v37_.f19, d_153_v0_, d_155_globalState_)
                d_205_v40_ = out4_
                d_206_v41_ = out5_
                d_207_v42_ = out6_
                d_208_v43_ = out7_
                d_156_v2_ = (0) - (d_205_v40_)
                d_209_v44_: _dafny.Array
                nw34_ = _dafny.Array(False, 14)
                d_209_v44_ = nw34_
                index36_ = default__.safeIndex(609, (d_209_v44_).length(0))
                (d_209_v44_)[index36_] = True
                index37_ = default__.safeIndex(609, (d_209_v44_).length(0))
                (d_209_v44_)[index37_] = d_153_v0_
            elif True:
                d_210_v45_: _dafny.Map
                d_210_v45_ = _dafny.Map({d_156_v2_: (d_153_v0_) or (d_153_v0_)})
                d_210_v45_ = (d_210_v45_).set(d_156_v2_, default__.fm2(False, d_155_globalState_))
                d_211_v46_: _dafny.Array
                nw35_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_211_v46_ = nw35_
                d_211_v46_ = d_211_v46_
                d_212_v47_: _dafny.Set
                d_212_v47_ = _dafny.Set({(d_153_v0_ if d_153_v0_ else d_153_v0_)})
                d_213_v48_: C0
                nw36_ = C0()
                nw36_.ctor__(d_156_v2_)
                d_213_v48_ = nw36_
                d_214_v49_: _dafny.Map
                d_214_v49_ = _dafny.Map({d_201_v36_: default__.fm2(d_153_v0_, d_155_globalState_)})
                d_215_v50_: _dafny.Set
                d_215_v50_ = _dafny.Set({d_213_v48_, d_213_v48_})
                d_216_v51_: _dafny.Map
                d_216_v51_ = _dafny.Map({d_156_v2_: d_201_v36_})
                d_217_v52_: _dafny.Seq
                d_217_v52_ = _dafny.SeqWithoutIsStrInference([d_216_v51_, d_216_v51_])
                d_218_v53_: _dafny.Map
                d_218_v53_ = _dafny.Map({d_153_v0_: len(d_154_v1_)})
                d_219_v54_: _dafny.Set
                d_219_v54_ = _dafny.Set({(d_217_v52_)[default__.safeIndex(d_156_v2_, len(d_217_v52_))], (d_216_v51_).set(len(d_218_v53_), d_201_v36_), (d_216_v51_).set(-407, d_201_v36_)})
                d_220_v55_: _dafny.Seq
                d_220_v55_ = _dafny.SeqWithoutIsStrInference([(0) - (d_156_v2_)])
                d_221_v56_: D0
                d_221_v56_ = D0_DC2(not(not(d_153_v0_)), d_153_v0_, d_156_v2_, d_153_v0_)
                d_222_v57_: _dafny.Array
                nw37_ = _dafny.Array(None, 16)
                nw37_[int(0)] = d_153_v0_
                nw37_[int(1)] = default__.fm2(((d_214_v49_)[d_201_v36_] if (d_201_v36_) in (d_214_v49_) else d_153_v0_), d_155_globalState_)
                nw37_[int(2)] = (d_215_v50_).ispropersubset(d_215_v50_)
                nw37_[int(3)] = (d_216_v51_) in (d_219_v54_)
                nw37_[int(4)] = (d_154_v1_)[default__.safeIndex(len(d_220_v55_), len(d_154_v1_))]
                nw37_[int(5)] = d_153_v0_
                nw37_[int(6)] = d_153_v0_
                nw37_[int(7)] = (d_221_v56_).cf4
                nw37_[int(8)] = not((d_153_v0_) == (d_153_v0_))
                nw37_[int(9)] = (d_220_v55_) != (_dafny.SeqWithoutIsStrInference([d_213_v48_.f19 for d_223_i7_ in range(default__.abs(116))]))
                nw37_[int(10)] = d_153_v0_
                nw37_[int(11)] = (d_153_v0_) or (d_153_v0_)
                nw37_[int(12)] = d_153_v0_
                nw37_[int(13)] = True
                nw37_[int(14)] = d_153_v0_
                nw37_[int(15)] = d_153_v0_
                d_222_v57_ = nw37_
                index38_ = default__.safeIndex(53, (d_222_v57_).length(0))
                (d_222_v57_)[index38_] = d_153_v0_
                d_224_v58_: _dafny.Seq
                d_224_v58_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_201_v36_]), d_175_v18_, d_175_v18_])
                index39_ = default__.safeIndex(53, (d_222_v57_).length(0))
                rhs39_ = ((d_212_v47_).intersection(d_212_v47_)) | (d_212_v47_)
                rhs40_ = (d_156_v2_) - (default__.fm0(d_213_v48_.f19, d_201_v36_, d_224_v58_, len(d_175_v18_), d_155_globalState_))
                rhs41_ = d_213_v48_
                rhs42_ = not ((d_213_v48_.f19) == (len(d_218_v53_))) or (d_153_v0_)
                lhs33_ = d_155_globalState_
                lhs34_ = d_222_v57_
                lhs35_ = default__.safeIndex(53, (d_222_v57_).length(0))
                d_212_v47_ = rhs39_
                lhs33_.f2 = rhs40_
                d_213_v48_ = rhs41_
                lhs34_[lhs35_] = rhs42_
                d_225_v59_: int
                d_226_v60_: _dafny.Seq
                d_227_v61_: int
                d_228_v62_: _dafny.MultiSet
                out8_: int
                out9_: _dafny.Seq
                out10_: int
                out11_: _dafny.MultiSet
                out8_, out9_, out10_, out11_ = default__.m0((d_213_v48_.f19) < (d_156_v2_), (default__.fm2(d_153_v0_, d_155_globalState_)) not in (d_154_v1_), (0) - (d_213_v48_.f19), d_153_v0_, d_155_globalState_)
                d_225_v59_ = out8_
                d_226_v60_ = out9_
                d_227_v61_ = out10_
                d_228_v62_ = out11_
                (d_155_globalState_).f3 = (d_222_v57_)[default__.safeIndex(53, (d_222_v57_).length(0))]
            (d_155_globalState_).f2 = d_156_v2_
            (d_155_globalState_).f2 = d_156_v2_
            (d_155_globalState_).f3 = d_153_v0_
        d_229_v63_: _dafny.Map
        d_229_v63_ = _dafny.Map({_dafny.Set({d_153_v0_, d_153_v0_, d_153_v0_, d_153_v0_}): (len(d_175_v18_)) == (d_156_v2_)})
        d_230_v64_: _dafny.Seq
        d_230_v64_ = _dafny.SeqWithoutIsStrInference([d_156_v2_])
        d_231_v65_: C0
        nw38_ = C0()
        nw38_.ctor__(d_156_v2_)
        d_231_v65_ = nw38_
        d_232_v66_: _dafny.Map
        d_232_v66_ = _dafny.Map({d_156_v2_: d_231_v65_})
        d_233_v67_: _dafny.Map
        d_233_v67_ = _dafny.Map({len(d_232_v66_): d_156_v2_})
        d_234_v68_: C0
        nw39_ = C0()
        nw39_.ctor__((d_230_v64_)[default__.safeIndex(len(d_233_v67_), len(d_230_v64_))])
        d_234_v68_ = nw39_
        d_235_v69_: _dafny.Map
        d_235_v69_ = _dafny.Map({d_231_v65_: d_153_v0_})
        d_229_v63_ = (d_229_v63_).set(_dafny.Set({not(((d_235_v69_)[d_231_v65_] if (d_231_v65_) in (d_235_v69_) else d_153_v0_))}), d_153_v0_)
        d_153_v0_ = d_153_v0_
        d_236_i8_: int
        d_236_i8_ = 0
        with _dafny.label("0"):
            while d_153_v0_:
                with _dafny.c_label("0"):
                    if (d_236_i8_) >= (100):
                        raise _dafny.Break("0")
                    d_236_i8_ = (d_236_i8_) + (1)
                    if d_153_v0_:
                        d_237_v70_: D0
                        d_237_v70_ = D0_DC2(True, d_153_v0_, d_234_v68_.f19, False)
                        d_238_v71_: _dafny.Array
                        nw40_ = _dafny.Array(None, 27)
                        nw40_[int(0)] = d_153_v0_
                        nw40_[int(1)] = d_153_v0_
                        nw40_[int(2)] = d_153_v0_
                        nw40_[int(3)] = d_153_v0_
                        nw40_[int(4)] = d_153_v0_
                        nw40_[int(5)] = (d_154_v1_)[default__.safeIndex(d_231_v65_.f19, len(d_154_v1_))]
                        nw40_[int(6)] = d_153_v0_
                        nw40_[int(7)] = d_153_v0_
                        nw40_[int(8)] = default__.fm2(not(d_153_v0_), d_155_globalState_)
                        nw40_[int(9)] = d_153_v0_
                        nw40_[int(10)] = d_153_v0_
                        nw40_[int(11)] = d_153_v0_
                        nw40_[int(12)] = d_153_v0_
                        nw40_[int(13)] = (d_237_v70_).cf4
                        nw40_[int(14)] = d_153_v0_
                        nw40_[int(15)] = not(d_153_v0_)
                        nw40_[int(16)] = d_153_v0_
                        nw40_[int(17)] = d_153_v0_
                        nw40_[int(18)] = d_153_v0_
                        nw40_[int(19)] = d_153_v0_
                        nw40_[int(20)] = d_153_v0_
                        nw40_[int(21)] = d_153_v0_
                        nw40_[int(22)] = d_153_v0_
                        nw40_[int(23)] = not(d_153_v0_)
                        nw40_[int(24)] = d_153_v0_
                        nw40_[int(25)] = d_153_v0_
                        nw40_[int(26)] = True
                        d_238_v71_ = nw40_
                        d_239_v72_: _dafny.Map
                        d_239_v72_ = _dafny.Map({d_238_v71_: d_153_v0_})
                        d_239_v72_ = (d_239_v72_).set(d_238_v71_, d_153_v0_)
                        d_240_v73_: _dafny.Array
                        nw41_ = _dafny.Array(None, 27)
                        d_240_v73_ = nw41_
                        index40_ = default__.safeIndex(854, (d_240_v73_).length(0))
                        (d_240_v73_)[index40_] = d_231_v65_
                        index41_ = default__.safeIndex(854, (d_240_v73_).length(0))
                        (d_240_v73_)[index41_] = d_234_v68_
                        d_241_v74_: _dafny.Array
                        nw42_ = _dafny.Array(int(0), 20)
                        d_241_v74_ = nw42_
                        index42_ = default__.safeIndex(987, (d_241_v74_).length(0))
                        (d_241_v74_)[index42_] = d_156_v2_
                        d_242_v75_: _dafny.Array
                        nw43_ = _dafny.Array(_dafny.Map({}), 2)
                        d_242_v75_ = nw43_
                        d_243_v76_: D0
                        d_243_v76_ = D0_DC0(d_231_v65_.f19)
                        d_244_v77_: _dafny.Seq
                        d_244_v77_ = _dafny.SeqWithoutIsStrInference([d_154_v1_])
                        d_245_v78_: _dafny.MultiSet
                        d_245_v78_ = _dafny.MultiSet([d_154_v1_, (d_244_v77_)[default__.safeIndex(d_234_v68_.f19, len(d_244_v77_))], d_154_v1_])
                        d_246_v79_: str
                        d_246_v79_ = _dafny.CodePoint('s')
                        d_247_v80_: _dafny.Seq
                        d_247_v80_ = _dafny.SeqWithoutIsStrInference([default__.fm5(d_156_v2_, d_155_globalState_), _dafny.SeqWithoutIsStrInference([d_246_v79_ for d_248_i9_ in range(default__.abs(382))]), d_175_v18_])
                        d_249_v81_: _dafny.Map
                        d_249_v81_ = _dafny.Map({d_240_v73_: d_242_v75_})
                        index43_ = default__.safeIndex(987, (d_241_v74_).length(0))
                        rhs43_ = not ((d_153_v0_ if d_153_v0_ else d_153_v0_)) or (d_153_v0_)
                        rhs44_ = d_234_v68_.f19
                        rhs45_ = default__.fm0(((_dafny.MultiSet(d_244_v77_)).intersection(d_245_v78_)).cardinality, d_246_v79_, d_247_v80_, d_231_v65_.f19, d_155_globalState_)
                        rhs46_ = ((d_249_v81_)[d_240_v73_] if (d_240_v73_) in (d_249_v81_) else d_242_v75_)
                        rhs47_ = d_243_v76_
                        lhs36_ = d_155_globalState_
                        lhs37_ = d_234_v68_
                        lhs38_ = d_241_v74_
                        lhs39_ = default__.safeIndex(987, (d_241_v74_).length(0))
                        lhs36_.f3 = rhs43_
                        lhs37_.f19 = rhs44_
                        lhs38_[lhs39_] = rhs45_
                        d_242_v75_ = rhs46_
                        d_243_v76_ = rhs47_
                        (d_231_v65_).f19 = (0) - ((0) - (d_234_v68_.f19))
                        index44_ = default__.safeIndex(7, (d_238_v71_).length(0))
                        (d_238_v71_)[index44_] = d_153_v0_
                        d_250_v82_: _dafny.Array
                        def lambda9_(d_251_v79_):
                            def lambda10_(d_252_i10_):
                                return d_251_v79_

                            return lambda10_

                        init3_ = lambda9_(d_246_v79_)
                        nw44_ = _dafny.Array(None, 9)
                        for i0_3_ in range(nw44_.length(0)):
                            nw44_[i0_3_] = init3_(i0_3_)
                        d_250_v82_ = nw44_
                        d_253_v83_: _dafny.MultiSet
                        d_253_v83_ = _dafny.MultiSet([d_250_v82_, d_250_v82_])
                        index45_ = default__.safeIndex(7, (d_238_v71_).length(0))
                        rhs48_ = d_156_v2_
                        rhs49_ = (((d_253_v83_).cardinality) <= (d_231_v65_.f19) if d_153_v0_ else d_153_v0_)
                        rhs50_ = (d_234_v68_.f19) >= (d_234_v68_.f19)
                        lhs40_ = d_155_globalState_
                        lhs41_ = d_238_v71_
                        lhs42_ = default__.safeIndex(7, (d_238_v71_).length(0))
                        lhs40_.f8 = rhs48_
                        lhs41_[lhs42_] = rhs49_
                        d_153_v0_ = rhs50_
                    elif True:
                        d_153_v0_ = False
                        d_254_v84_: _dafny.Array
                        def lambda11_(d_255_v0_):
                            def lambda12_(d_256_i11_):
                                return d_255_v0_

                            return lambda12_

                        init4_ = lambda11_(d_153_v0_)
                        nw45_ = _dafny.Array(None, 8)
                        for i0_4_ in range(nw45_.length(0)):
                            nw45_[i0_4_] = init4_(i0_4_)
                        d_254_v84_ = nw45_
                        d_257_v85_: _dafny.Seq
                        d_257_v85_ = _dafny.SeqWithoutIsStrInference([d_254_v84_, d_254_v84_])
                        d_258_v86_: _dafny.MultiSet
                        d_258_v86_ = _dafny.MultiSet([d_254_v84_, d_254_v84_, (d_257_v85_)[default__.safeIndex(len(d_175_v18_), len(d_257_v85_))]])
                        d_258_v86_ = (d_258_v86_) | ((d_258_v86_).set(d_254_v84_, default__.abs(d_231_v65_.f19)))
                        d_259_v87_: _dafny.Map
                        d_259_v87_ = _dafny.Map({d_153_v0_: d_153_v0_})
                        d_260_v88_: _dafny.Map
                        d_260_v88_ = _dafny.Map({d_259_v87_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))})
                        d_261_v89_: D0
                        d_261_v89_ = D0_DC2(d_153_v0_, d_153_v0_, d_234_v68_.f19, d_153_v0_)
                        d_262_v90_: _dafny.Map
                        d_262_v90_ = _dafny.Map({_dafny.MultiSet([d_261_v89_, d_261_v89_]): d_153_v0_})
                        d_263_v91_: _dafny.Seq
                        d_263_v91_ = _dafny.SeqWithoutIsStrInference([d_260_v88_])
                        d_264_v92_: D1
                        d_264_v92_ = D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))
                        d_265_v94_: _dafny.Seq
                        d_265_v94_ = _dafny.SeqWithoutIsStrInference([d_259_v87_, _dafny.Map({d_153_v0_: ((d_259_v87_)[True] if (True) in (d_259_v87_) else d_153_v0_)})])
                        d_266_v96_: _dafny.Seq
                        d_266_v96_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbbrxwsjm")), d_175_v18_])
                        d_267_v97_: _dafny.Array
                        nw46_ = _dafny.Array(None, 26)
                        nw46_[int(0)] = (d_260_v88_) | (d_260_v88_)
                        nw46_[int(1)] = d_260_v88_
                        nw46_[int(2)] = (d_260_v88_) | (d_260_v88_)
                        nw46_[int(3)] = default__.fm6(d_231_v65_.f19, d_156_v2_, d_153_v0_, d_262_v90_, d_155_globalState_)
                        nw46_[int(4)] = d_260_v88_
                        nw46_[int(5)] = (d_260_v88_).set(_dafny.Map({d_153_v0_: default__.fm2(d_153_v0_, d_155_globalState_)}), d_175_v18_)
                        nw46_[int(6)] = d_260_v88_
                        nw46_[int(7)] = _dafny.Map({d_259_v87_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yv"))})
                        nw46_[int(8)] = d_260_v88_
                        nw46_[int(9)] = (d_260_v88_) | (d_260_v88_)
                        nw46_[int(10)] = d_260_v88_
                        nw46_[int(11)] = (d_263_v91_)[default__.safeIndex(d_156_v2_, len(d_263_v91_))]
                        def iife21_():
                            coll5_ = _dafny.Map()
                            compr_5_: _dafny.Map
                            for compr_5_ in (_dafny.MultiSet(d_265_v94_)).Elements:
                                d_268_v93_: _dafny.Map = compr_5_
                                if (d_268_v93_) in (_dafny.MultiSet(d_265_v94_)):
                                    coll5_[d_268_v93_] = d_175_v18_
                            return _dafny.Map(coll5_)
                        nw46_[int(12)] = (_dafny.Map({d_259_v87_: (d_264_v92_).cf8})) | (iife21_()
                        )
                        nw46_[int(13)] = (d_260_v88_) | (d_260_v88_)
                        def iife22_():
                            coll6_ = _dafny.Map()
                            compr_6_: _dafny.Map
                            for compr_6_ in (d_265_v94_).Elements:
                                d_269_v95_: _dafny.Map = compr_6_
                                if (d_269_v95_) in (d_265_v94_):
                                    coll6_[d_269_v95_] = d_175_v18_
                            return _dafny.Map(coll6_)
                        nw46_[int(14)] = iife22_()
                        
                        nw46_[int(15)] = d_260_v88_
                        nw46_[int(16)] = ((d_260_v88_).set(d_259_v87_, (d_266_v96_)[default__.safeIndex(len(d_230_v64_), len(d_266_v96_))])) | (d_260_v88_)
                        nw46_[int(17)] = _dafny.Map({d_259_v87_: d_175_v18_})
                        nw46_[int(18)] = d_260_v88_
                        nw46_[int(19)] = (d_263_v91_)[default__.safeIndex(d_231_v65_.f19, len(d_263_v91_))]
                        nw46_[int(20)] = d_260_v88_
                        nw46_[int(21)] = _dafny.Map({d_259_v87_: (d_175_v18_).set(default__.safeIndex(-222, len(d_175_v18_)), (d_175_v18_)[default__.safeIndex(d_234_v68_.f19, len(d_175_v18_))])})
                        nw46_[int(22)] = d_260_v88_
                        nw46_[int(23)] = d_260_v88_
                        nw46_[int(24)] = d_260_v88_
                        nw46_[int(25)] = d_260_v88_
                        d_267_v97_ = nw46_
                        index46_ = default__.safeIndex(661, (d_267_v97_).length(0))
                        (d_267_v97_)[index46_] = (d_260_v88_) | (_dafny.Map({d_259_v87_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_270_i12_ in range(default__.abs(37))])}))
                        index47_ = default__.safeIndex(661, (d_267_v97_).length(0))
                        (d_267_v97_)[index47_] = d_260_v88_
                        d_271_v98_: _dafny.Array
                        def lambda13_(d_272_v0_):
                            def lambda14_(d_273_i14_):
                                return _dafny.Set({d_272_v0_, d_272_v0_, False})

                            return lambda14_

                        init5_ = lambda13_(d_153_v0_)
                        nw47_ = _dafny.Array(None, 13)
                        for i0_5_ in range(nw47_.length(0)):
                            nw47_[i0_5_] = init5_(i0_5_)
                        d_271_v98_ = nw47_
                        d_274_v99_: _dafny.Map
                        d_274_v99_ = _dafny.Map({d_271_v98_: d_175_v18_})
                        d_275_v100_: _dafny.Array
                        nw48_ = _dafny.Array(None, 27)
                        nw48_[int(0)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_276_i13_ in range(default__.abs(-94))])
                        nw48_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjdvjrwn"))
                        nw48_[int(2)] = (d_175_v18_ if d_153_v0_ else d_175_v18_)
                        nw48_[int(3)] = ((d_274_v99_)[d_271_v98_] if (d_271_v98_) in (d_274_v99_) else d_175_v18_)
                        nw48_[int(4)] = default__.fm5(d_234_v68_.f19, d_155_globalState_)
                        nw48_[int(5)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_277_i15_ in range(default__.abs(-417))])
                        nw48_[int(6)] = d_175_v18_
                        nw48_[int(7)] = d_175_v18_
                        nw48_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxgk"))
                        nw48_[int(9)] = d_175_v18_
                        nw48_[int(10)] = d_175_v18_
                        nw48_[int(11)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vb"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_278_i16_ in range(default__.abs(-384))]))
                        nw48_[int(12)] = (d_175_v18_) + (d_175_v18_)
                        nw48_[int(13)] = d_175_v18_
                        nw48_[int(14)] = d_175_v18_
                        nw48_[int(15)] = d_175_v18_
                        nw48_[int(16)] = d_175_v18_
                        nw48_[int(17)] = d_175_v18_
                        nw48_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
                        nw48_[int(19)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_279_i17_ in range(default__.abs(340))])
                        nw48_[int(20)] = (d_175_v18_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkieswsc")))
                        nw48_[int(21)] = (d_175_v18_) + (d_175_v18_)
                        nw48_[int(22)] = d_175_v18_
                        nw48_[int(23)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_280_i18_ in range(default__.abs(551))])
                        nw48_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uscna"))
                        nw48_[int(25)] = ((d_266_v96_)[default__.safeIndex(628, len(d_266_v96_))]) + (d_175_v18_)
                        nw48_[int(26)] = d_175_v18_
                        d_275_v100_ = nw48_
                        index48_ = default__.safeIndex(571, (d_275_v100_).length(0))
                        (d_275_v100_)[index48_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btynhwo")) if d_153_v0_ else d_175_v18_)
                        index49_ = default__.safeIndex(571, (d_275_v100_).length(0))
                        (d_275_v100_)[index49_] = d_175_v18_
                        rhs51_ = d_234_v68_.f19
                        rhs52_ = ((d_233_v67_)[len(d_230_v64_)] if (len(d_230_v64_)) in (d_233_v67_) else d_234_v68_.f19)
                        rhs53_ = not(d_153_v0_)
                        rhs54_ = len((d_275_v100_)[default__.safeIndex(571, (d_275_v100_).length(0))])
                        rhs55_ = ((d_234_v68_.f19) * (d_234_v68_.f19)) + (((d_157_v3_)[d_153_v0_] if (d_153_v0_) in (d_157_v3_) else d_231_v65_.f19))
                        lhs43_ = d_231_v65_
                        lhs44_ = d_155_globalState_
                        lhs45_ = d_155_globalState_
                        lhs46_ = d_155_globalState_
                        lhs47_ = d_155_globalState_
                        lhs43_.f19 = rhs51_
                        lhs44_.f0 = rhs52_
                        lhs45_.f3 = rhs53_
                        lhs46_.f8 = rhs54_
                        lhs47_.f12 = rhs55_
                    (d_155_globalState_).f3 = d_153_v0_
                    d_281_v101_: D1
                    d_281_v101_ = D1_DC5((not(d_153_v0_)) and (d_153_v0_))
                    source6_ = d_281_v101_
                    if source6_.is_DC4:
                        d_282___mcc_h0_ = source6_.cf9
                        d_283_cf9_ = d_282___mcc_h0_
                        d_154_v1_ = d_154_v1_
                        d_157_v3_ = (d_157_v3_) - (d_157_v3_)
                        (d_231_v65_).f19 = (0) - (d_231_v65_.f19)
                        d_284_v102_: str
                        d_284_v102_ = _dafny.CodePoint('s')
                        d_285_v103_: _dafny.Map
                        d_285_v103_ = _dafny.Map({d_153_v0_: d_230_v64_})
                        rhs56_ = (_dafny.CodePoint('a') if d_153_v0_ else default__.fm7(d_153_v0_, d_231_v65_.f19, d_155_globalState_))
                        rhs57_ = d_153_v0_
                        rhs58_ = (d_234_v68_).fm4((0) - (d_283_cf9_.f19), d_156_v2_, d_153_v0_, (d_285_v103_) != (default__.fm8(len(d_175_v18_), d_283_cf9_.f19, d_153_v0_, d_155_globalState_)), d_155_globalState_)
                        lhs48_ = d_155_globalState_
                        lhs49_ = d_155_globalState_
                        d_284_v102_ = rhs56_
                        lhs48_.f3 = rhs57_
                        lhs49_.f8 = rhs58_
                    elif source6_.is_DC5:
                        d_286___mcc_h1_ = source6_.cf10
                        d_287_cf10_ = d_286___mcc_h1_
                        (d_155_globalState_).f3 = d_153_v0_
                        d_233_v67_ = (d_233_v67_).set(d_231_v65_.f19, default__.safeDivisionInt(d_234_v68_.f19, d_234_v68_.f19))
                        d_288_v104_: _dafny.Array
                        def lambda15_(d_289_i19_):
                            return _dafny.CodePoint('i')

                        init6_ = lambda15_
                        nw49_ = _dafny.Array(None, 26)
                        for i0_6_ in range(nw49_.length(0)):
                            nw49_[i0_6_] = init6_(i0_6_)
                        d_288_v104_ = nw49_
                        index50_ = default__.safeIndex(978, (d_288_v104_).length(0))
                        (d_288_v104_)[index50_] = _dafny.CodePoint('i')
                        d_290_v105_: str
                        d_290_v105_ = _dafny.CodePoint('l')
                        d_291_v106_: _dafny.Set
                        d_291_v106_ = _dafny.Set({d_290_v105_, d_290_v105_, d_290_v105_})
                        d_292_v107_: _dafny.MultiSet
                        d_292_v107_ = _dafny.MultiSet([d_290_v105_, d_290_v105_])
                        d_293_v108_: _dafny.Seq
                        d_293_v108_ = _dafny.SeqWithoutIsStrInference([d_175_v18_, default__.fm5(((d_292_v107_)[d_290_v105_] if (d_290_v105_) in (d_292_v107_) else d_234_v68_.f19), d_155_globalState_)])
                        index51_ = default__.safeIndex(978, (d_288_v104_).length(0))
                        rhs59_ = not (d_287_cf10_) or ((d_291_v106_) != (_dafny.Set({d_290_v105_, d_290_v105_, d_290_v105_})))
                        rhs60_ = d_290_v105_
                        rhs61_ = ((d_154_v1_) + (d_154_v1_)) < (_dafny.SeqWithoutIsStrInference([d_153_v0_]))
                        rhs62_ = (d_293_v108_) != ((d_293_v108_) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owfqc"))])))
                        lhs50_ = d_155_globalState_
                        lhs51_ = d_288_v104_
                        lhs52_ = default__.safeIndex(978, (d_288_v104_).length(0))
                        lhs53_ = d_155_globalState_
                        lhs50_.f3 = rhs59_
                        lhs51_[lhs52_] = rhs60_
                        lhs53_.f3 = rhs61_
                        d_287_cf10_ = rhs62_
                        d_294_v110_: _dafny.Array
                        nw50_ = _dafny.Array(None, 3)
                        def iife23_():
                            coll7_ = _dafny.Map()
                            compr_7_: int
                            for compr_7_ in (d_230_v64_).Elements:
                                d_295_v109_: int = compr_7_
                                if (d_295_v109_) in (d_230_v64_):
                                    coll7_[default__.safeDivisionInt(d_295_v109_, (0) - (d_156_v2_))] = (D0_DC1(d_234_v68_.f19, 24, True)).cf1
                            return _dafny.Map(coll7_)
                        nw50_[int(0)] = len(iife23_()
                        )
                        nw50_[int(1)] = d_234_v68_.f19
                        nw50_[int(2)] = d_231_v65_.f19
                        d_294_v110_ = nw50_
                        index52_ = default__.safeIndex(817, (d_294_v110_).length(0))
                        (d_294_v110_)[index52_] = d_156_v2_
                        d_296_v111_: _dafny.MultiSet
                        d_296_v111_ = _dafny.MultiSet([d_233_v67_])
                        index53_ = default__.safeIndex(817, (d_294_v110_).length(0))
                        (d_294_v110_)[index53_] = (d_230_v64_)[default__.safeIndex((d_296_v111_).cardinality, len(d_230_v64_))]
                    elif source6_.is_DC3:
                        d_297___mcc_h2_ = source6_.cf8
                        d_298_cf8_ = d_297___mcc_h2_
                        d_299_v112_: D0
                        d_299_v112_ = D0_DC2(not(d_153_v0_), d_153_v0_, d_234_v68_.f19, d_153_v0_)
                        d_300_v113_: _dafny.Array
                        nw51_ = _dafny.Array(None, 1)
                        nw51_[int(0)] = d_299_v112_
                        d_300_v113_ = nw51_
                        d_301_v114_: D1
                        d_301_v114_ = D1_DC3(d_175_v18_)
                        d_302_v115_: _dafny.Map
                        d_302_v115_ = _dafny.Map({d_300_v113_: d_301_v114_})
                        d_302_v115_ = ((d_302_v115_) | (d_302_v115_)) | (d_302_v115_)
                        (d_155_globalState_).f3 = not((d_234_v68_.f19) == ((d_231_v65_).fm3(d_155_globalState_)))
                        (d_155_globalState_).f3 = default__.fm2(d_153_v0_, d_155_globalState_)
                        (d_155_globalState_).f8 = d_231_v65_.f19
                    elif True:
                        d_303___mcc_h3_ = source6_.cf11
                        d_304_cf11_ = d_303___mcc_h3_
                        d_305_v116_: C0
                        nw52_ = C0()
                        nw52_.ctor__(d_231_v65_.f19)
                        d_305_v116_ = nw52_
                        d_306_v117_: _dafny.Map
                        d_306_v117_ = _dafny.Map({default__.fm2(True, d_155_globalState_): d_234_v68_.f19})
                        d_306_v117_ = (d_306_v117_).set(True, d_234_v68_.f19)
                        d_230_v64_ = d_230_v64_
                        d_307_v118_: _dafny.Array
                        def lambda16_(d_308_i20_):
                            return default__.safeDivisionInt(d_308_i20_, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nffg")))))

                        init7_ = lambda16_
                        nw53_ = _dafny.Array(None, 10)
                        for i0_7_ in range(nw53_.length(0)):
                            nw53_[i0_7_] = init7_(i0_7_)
                        d_307_v118_ = nw53_
                        d_309_v119_: _dafny.MultiSet
                        d_309_v119_ = _dafny.MultiSet([d_307_v118_])
                        d_310_v120_: _dafny.Map
                        d_310_v120_ = _dafny.Map({len(d_306_v117_): d_153_v0_})
                        d_311_v121_: D0
                        d_311_v121_ = D0_DC1(len(d_310_v120_), d_234_v68_.f19, d_153_v0_)
                        d_312_v122_: int
                        d_313_v123_: _dafny.Seq
                        d_314_v124_: int
                        d_315_v125_: _dafny.MultiSet
                        out12_: int
                        out13_: _dafny.Seq
                        out14_: int
                        out15_: _dafny.MultiSet
                        out12_, out13_, out14_, out15_ = default__.m0((d_231_v65_.f19) >= ((d_309_v119_).cardinality), d_153_v0_, (d_157_v3_).cardinality, ((0) - ((d_234_v68_).fm4(len(d_175_v18_), d_231_v65_.f19, (d_311_v121_).cf3, d_153_v0_, d_155_globalState_))) == (d_305_v116_.f19), d_155_globalState_)
                        d_312_v122_ = out12_
                        d_313_v123_ = out13_
                        d_314_v124_ = out14_
                        d_315_v125_ = out15_
                    (d_155_globalState_).f0 = d_231_v65_.f19
                    pass
            pass
        d_316_v126_: _dafny.Set
        d_316_v126_ = _dafny.Set({d_153_v0_, d_153_v0_})
        hi2_ = d_231_v65_.f19
        for d_317_i21_ in range(((0) - (d_156_v2_) if not(d_153_v0_) else len(_dafny.Map({(0) - (d_234_v68_.f19): len(d_316_v126_)}))), hi2_):
            d_175_v18_ = d_175_v18_
            d_318_v127_: D1
            d_318_v127_ = D1_DC5(d_153_v0_)
            d_319_v128_: _dafny.Array
            nw54_ = _dafny.Array(None, 18)
            nw54_[int(0)] = (d_318_v127_).cf10
            nw54_[int(1)] = d_153_v0_
            nw54_[int(2)] = False
            nw54_[int(3)] = d_153_v0_
            nw54_[int(4)] = d_153_v0_
            nw54_[int(5)] = True
            nw54_[int(6)] = d_153_v0_
            nw54_[int(7)] = d_153_v0_
            nw54_[int(8)] = d_153_v0_
            nw54_[int(9)] = d_153_v0_
            nw54_[int(10)] = not(d_153_v0_)
            nw54_[int(11)] = default__.fm2(d_153_v0_, d_155_globalState_)
            nw54_[int(12)] = True
            nw54_[int(13)] = d_153_v0_
            nw54_[int(14)] = d_153_v0_
            nw54_[int(15)] = not(d_153_v0_)
            nw54_[int(16)] = d_153_v0_
            nw54_[int(17)] = default__.fm2(d_153_v0_, d_155_globalState_)
            d_319_v128_ = nw54_
            d_320_v129_: _dafny.Array
            nw55_ = _dafny.Array(None, 2)
            nw55_[int(0)] = (d_319_v128_ if d_153_v0_ else d_319_v128_)
            nw55_[int(1)] = d_319_v128_
            d_320_v129_ = nw55_
            index54_ = default__.safeIndex(187, (d_320_v129_).length(0))
            (d_320_v129_)[index54_] = d_319_v128_
            index55_ = default__.safeIndex(187, (d_320_v129_).length(0))
            (d_320_v129_)[index55_] = (D2_DC7(d_319_v128_)).cf12
            d_321_v130_: _dafny.Seq
            d_321_v130_ = _dafny.SeqWithoutIsStrInference([d_157_v3_, d_157_v3_])
            d_322_v131_: C0
            nw56_ = C0()
            nw56_.ctor__(len(d_321_v130_))
            d_322_v131_ = nw56_
            d_323_v132_: D2
            d_323_v132_ = D2_DC8()
            source7_ = d_323_v132_
            if source7_.is_DC8:
                d_324_v133_: _dafny.Array
                def lambda17_(d_325_i22_):
                    return (d_325_i22_) + (469)

                init8_ = lambda17_
                nw57_ = _dafny.Array(None, 17)
                for i0_8_ in range(nw57_.length(0)):
                    nw57_[i0_8_] = init8_(i0_8_)
                d_324_v133_ = nw57_
                index56_ = default__.safeIndex(567, (d_324_v133_).length(0))
                (d_324_v133_)[index56_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pyjrottw")))
                d_326_v134_: _dafny.Set
                d_326_v134_ = _dafny.Set({len(default__.fm5(len(d_233_v67_), d_155_globalState_))})
                d_327_v135_: _dafny.Map
                d_327_v135_ = _dafny.Map({d_234_v68_.f19: d_153_v0_})
                d_328_v136_: D3
                d_328_v136_ = D3_DC11(d_327_v135_)
                index57_ = default__.safeIndex(567, (d_324_v133_).length(0))
                rhs63_ = len((d_326_v134_) - (d_326_v134_))
                rhs64_ = ((d_230_v64_)[default__.safeIndex(d_322_v131_.f19, len(d_230_v64_))]) + (len((d_328_v136_).cf17))
                rhs65_ = d_231_v65_.f19
                lhs54_ = d_324_v133_
                lhs55_ = default__.safeIndex(567, (d_324_v133_).length(0))
                lhs56_ = d_155_globalState_
                d_156_v2_ = rhs63_
                lhs54_[lhs55_] = rhs64_
                lhs56_.f0 = rhs65_
                d_153_v0_ = d_153_v0_
                d_329_v137_: _dafny.Map
                d_329_v137_ = _dafny.Map({d_153_v0_: d_153_v0_})
                d_330_v138_: int
                d_331_v139_: _dafny.Seq
                d_332_v140_: int
                d_333_v141_: _dafny.MultiSet
                out16_: int
                out17_: _dafny.Seq
                out18_: int
                out19_: _dafny.MultiSet
                out16_, out17_, out18_, out19_ = default__.m0(d_153_v0_, (d_153_v0_) not in (d_329_v137_), d_317_i21_, ((d_329_v137_)[d_153_v0_] if (d_153_v0_) in (d_329_v137_) else d_153_v0_), d_155_globalState_)
                d_330_v138_ = out16_
                d_331_v139_ = out17_
                d_332_v140_ = out18_
                d_333_v141_ = out19_
                d_331_v139_ = d_175_v18_
            elif source7_.is_DC9:
                d_334___mcc_h4_ = source7_.cf13
                d_335___mcc_h5_ = source7_.cf14
                d_336___mcc_h6_ = source7_.cf15
                d_337_cf15_ = d_336___mcc_h6_
                d_338_cf14_ = d_335___mcc_h5_
                d_339_cf13_ = d_334___mcc_h4_
                (d_155_globalState_).f0 = (default__.safeModuloInt(763, d_231_v65_.f19)) + ((0) - (d_322_v131_.f19))
                d_153_v0_ = False
                d_340_v142_: str
                d_340_v142_ = _dafny.CodePoint('x')
                d_341_v143_: _dafny.Map
                d_341_v143_ = _dafny.Map({False: d_340_v142_})
                d_341_v143_ = (d_341_v143_).set(d_153_v0_, d_340_v142_)
                d_342_v144_: D1
                d_342_v144_ = D1_DC4(d_338_cf14_)
                pat_let_tv19_ = d_234_v68_
                pat_let_tv20_ = d_338_cf14_
                d_343_v145_: _dafny.Array
                nw58_ = _dafny.Array(None, 25)
                nw58_[int(0)] = D1_DC4(d_322_v131_)
                nw58_[int(1)] = (d_342_v144_ if d_153_v0_ else D1_DC4(d_234_v68_))
                nw58_[int(2)] = d_342_v144_
                nw58_[int(3)] = d_342_v144_
                nw58_[int(4)] = d_342_v144_
                nw58_[int(5)] = D1_DC4(d_231_v65_)
                nw58_[int(6)] = d_342_v144_
                nw58_[int(7)] = d_342_v144_
                nw58_[int(8)] = d_342_v144_
                nw58_[int(9)] = d_342_v144_
                nw58_[int(10)] = d_342_v144_
                def iife24_(_pat_let8_0):
                    def iife25_(d_344_dt__update__tmp_h1_):
                        def iife26_(_pat_let9_0):
                            def iife27_(d_345_dt__update_hcf9_h0_):
                                return D1_DC4(d_345_dt__update_hcf9_h0_)
                            return iife27_(_pat_let9_0)
                        return iife26_(pat_let_tv19_)
                    return iife25_(_pat_let8_0)
                nw58_[int(11)] = (d_342_v144_ if d_153_v0_ else iife24_(d_342_v144_))
                nw58_[int(12)] = D1_DC4(d_234_v68_)
                nw58_[int(13)] = d_342_v144_
                nw58_[int(14)] = d_342_v144_
                nw58_[int(15)] = d_342_v144_
                nw58_[int(16)] = d_342_v144_
                nw58_[int(17)] = d_342_v144_
                nw58_[int(18)] = d_342_v144_
                nw58_[int(19)] = d_342_v144_
                nw58_[int(20)] = d_342_v144_
                def iife28_(_pat_let10_0):
                    def iife29_(d_346_dt__update__tmp_h2_):
                        def iife30_(_pat_let11_0):
                            def iife31_(d_347_dt__update_hcf9_h1_):
                                return D1_DC4(d_347_dt__update_hcf9_h1_)
                            return iife31_(_pat_let11_0)
                        return iife30_(pat_let_tv20_)
                    return iife29_(_pat_let10_0)
                nw58_[int(21)] = iife28_(d_342_v144_)
                nw58_[int(22)] = D1_DC4(d_322_v131_)
                nw58_[int(23)] = D1_DC4(d_234_v68_)
                nw58_[int(24)] = d_342_v144_
                d_343_v145_ = nw58_
                index58_ = default__.safeIndex(295, (d_343_v145_).length(0))
                (d_343_v145_)[index58_] = d_342_v144_
                index59_ = default__.safeIndex(295, (d_343_v145_).length(0))
                rhs66_ = d_342_v144_
                rhs67_ = (0) - (d_317_i21_)
                rhs68_ = (D0_DC2(d_153_v0_, d_153_v0_, d_234_v68_.f19, d_153_v0_)).cf6
                rhs69_ = (d_231_v65_.f19) * (d_234_v68_.f19)
                lhs57_ = d_343_v145_
                lhs58_ = default__.safeIndex(295, (d_343_v145_).length(0))
                lhs59_ = d_231_v65_
                lhs60_ = d_234_v68_
                lhs61_ = d_155_globalState_
                lhs57_[lhs58_] = rhs66_
                lhs59_.f19 = rhs67_
                lhs60_.f19 = rhs68_
                lhs61_.f2 = rhs69_
            elif source7_.is_DC7:
                d_348___mcc_h7_ = source7_.cf12
                d_349_cf12_ = d_348___mcc_h7_
                d_350_v146_: _dafny.Array
                nw59_ = _dafny.Array(_dafny.Map({}), 17)
                d_350_v146_ = nw59_
                d_351_v147_: D2
                d_351_v147_ = D2_DC9(d_234_v68_.f19, d_322_v131_, d_231_v65_)
                d_352_v148_: _dafny.Seq
                d_352_v148_ = _dafny.SeqWithoutIsStrInference([d_351_v147_])
                d_353_v149_: _dafny.Map
                d_353_v149_ = _dafny.Map({d_234_v68_.f19: D2_DC10((d_352_v148_)[default__.safeIndex(d_322_v131_.f19, len(d_352_v148_))])})
                index60_ = default__.safeIndex(316, (d_350_v146_).length(0))
                (d_350_v146_)[index60_] = d_353_v149_
                index61_ = default__.safeIndex(316, (d_350_v146_).length(0))
                (d_350_v146_)[index61_] = (d_353_v149_) | (d_353_v149_)
                d_354_v150_: _dafny.MultiSet
                d_354_v150_ = _dafny.MultiSet([d_231_v65_, d_231_v65_])
                d_355_v151_: _dafny.Map
                d_355_v151_ = _dafny.Map({d_354_v150_: d_317_i21_})
                d_356_v152_: _dafny.Map
                d_356_v152_ = _dafny.Map({False: d_153_v0_})
                d_357_v153_: _dafny.Seq
                d_357_v153_ = _dafny.SeqWithoutIsStrInference([d_234_v68_])
                d_358_v154_: D4
                d_358_v154_ = D4_DC14(d_357_v153_)
                d_359_v155_: _dafny.MultiSet
                d_359_v155_ = _dafny.MultiSet([d_156_v2_, d_231_v65_.f19])
                d_360_v156_: _dafny.Array
                nw60_ = _dafny.Array(None, 8)
                nw60_[int(0)] = d_355_v151_
                nw60_[int(1)] = (d_355_v151_) | (((d_355_v151_).set(d_354_v150_, len(d_356_v152_))).set(d_354_v150_, (d_230_v64_)[default__.safeIndex(d_234_v68_.f19, len(d_230_v64_))]))
                nw60_[int(2)] = d_355_v151_
                nw60_[int(3)] = _dafny.Map({_dafny.MultiSet((d_358_v154_).cf21): d_317_i21_})
                nw60_[int(4)] = (_dafny.Map({d_354_v150_: d_231_v65_.f19})) | (_dafny.Map({d_354_v150_: d_317_i21_}))
                nw60_[int(5)] = (d_355_v151_).set(d_354_v150_, (0) - ((d_359_v155_).cardinality))
                nw60_[int(6)] = d_355_v151_
                nw60_[int(7)] = d_355_v151_
                d_360_v156_ = nw60_
                d_361_v157_: _dafny.Map
                d_361_v157_ = _dafny.Map({d_153_v0_: d_322_v131_.f19})
                index62_ = default__.safeIndex(915, (d_360_v156_).length(0))
                (d_360_v156_)[index62_] = (d_355_v151_).set(d_354_v150_, len(d_361_v157_))
                d_362_v158_: _dafny.Array
                def lambda18_(d_363_v68_):
                    def lambda19_(d_364_i23_):
                        return (d_364_i23_) - (d_363_v68_.f19)

                    return lambda19_

                init9_ = lambda18_(d_234_v68_)
                nw61_ = _dafny.Array(None, 2)
                for i0_9_ in range(nw61_.length(0)):
                    nw61_[i0_9_] = init9_(i0_9_)
                d_362_v158_ = nw61_
                d_365_v159_: _dafny.Map
                d_365_v159_ = _dafny.Map({(d_233_v67_) | (d_233_v67_): d_362_v158_})
                index63_ = default__.safeIndex(915, (d_360_v156_).length(0))
                def iife32_():
                    coll8_ = _dafny.Map()
                    compr_8_: int
                    for compr_8_ in _dafny.IntegerRange(333, 713):
                        d_366_v160_: int = compr_8_
                        if ((333) <= (d_366_v160_)) and ((d_366_v160_) < (713)):
                            coll8_[(d_366_v160_) + ((0) - (len(d_175_v18_)))] = 719
                    return _dafny.Map(coll8_)
                def iife33_():
                    coll9_ = _dafny.Map()
                    compr_9_: int
                    for compr_9_ in _dafny.IntegerRange(333, 713):
                        d_367_v160_: int = compr_9_
                        if ((333) <= (d_367_v160_)) and ((d_367_v160_) < (713)):
                            coll9_[(d_367_v160_) + ((0) - (len(d_175_v18_)))] = 719
                    return _dafny.Map(coll9_)
                rhs70_ = len((d_230_v64_) + (d_230_v64_))
                rhs71_ = d_355_v151_
                rhs72_ = ((d_365_v159_)[iife32_()
                ] if (iife33_()
                ) in (d_365_v159_) else d_362_v158_)
                lhs62_ = d_155_globalState_
                lhs63_ = d_360_v156_
                lhs64_ = default__.safeIndex(915, (d_360_v156_).length(0))
                lhs62_.f12 = rhs70_
                lhs63_[lhs64_] = rhs71_
                d_362_v158_ = rhs72_
                (d_155_globalState_).f8 = default__.safeDivisionInt(701, d_231_v65_.f19)
                index64_ = default__.safeIndex(31, (d_319_v128_).length(0))
                (d_319_v128_)[index64_] = d_153_v0_
                index65_ = default__.safeIndex(31, (d_319_v128_).length(0))
                (d_319_v128_)[index65_] = d_153_v0_
            elif True:
                d_368___mcc_h8_ = source7_.cf16
                d_369_cf16_ = d_368___mcc_h8_
                (d_155_globalState_).f3 = (d_153_v0_ if default__.fm2(d_153_v0_, d_155_globalState_) else (d_234_v68_.f19) == (d_322_v131_.f19))
                d_370_v161_: str
                d_370_v161_ = _dafny.CodePoint('e')
                d_371_v162_: D5
                d_371_v162_ = D5_DC17((d_175_v18_)[default__.safeIndex(len(d_175_v18_), len(d_175_v18_))])
                d_370_v161_ = (d_371_v162_).cf23
                d_372_v163_: _dafny.Map
                d_372_v163_ = _dafny.Map({d_153_v0_: True})
                d_373_v164_: _dafny.MultiSet
                d_373_v164_ = _dafny.MultiSet([len(d_372_v163_)])
                rhs73_ = d_234_v68_
                rhs74_ = ((_dafny.SeqWithoutIsStrInference([d_370_v161_, d_370_v161_, (d_175_v18_)[default__.safeIndex(((d_373_v164_)[d_231_v65_.f19] if (d_231_v65_.f19) in (d_373_v164_) else (d_157_v3_).cardinality), len(d_175_v18_))], (_dafny.CodePoint('j') if d_153_v0_ else _dafny.CodePoint('t')), d_370_v161_])).set(default__.safeIndex((d_231_v65_.f19) * (d_317_i21_), len(_dafny.SeqWithoutIsStrInference([d_370_v161_, d_370_v161_, (d_175_v18_)[default__.safeIndex(((d_373_v164_)[d_231_v65_.f19] if (d_231_v65_.f19) in (d_373_v164_) else (d_157_v3_).cardinality), len(d_175_v18_))], (_dafny.CodePoint('j') if d_153_v0_ else _dafny.CodePoint('t')), d_370_v161_]))), d_370_v161_)).set(default__.safeIndex(d_231_v65_.f19, len((_dafny.SeqWithoutIsStrInference([d_370_v161_, d_370_v161_, (d_175_v18_)[default__.safeIndex(((d_373_v164_)[d_231_v65_.f19] if (d_231_v65_.f19) in (d_373_v164_) else (d_157_v3_).cardinality), len(d_175_v18_))], (_dafny.CodePoint('j') if d_153_v0_ else _dafny.CodePoint('t')), d_370_v161_])).set(default__.safeIndex((d_231_v65_.f19) * (d_317_i21_), len(_dafny.SeqWithoutIsStrInference([d_370_v161_, d_370_v161_, (d_175_v18_)[default__.safeIndex(((d_373_v164_)[d_231_v65_.f19] if (d_231_v65_.f19) in (d_373_v164_) else (d_157_v3_).cardinality), len(d_175_v18_))], (_dafny.CodePoint('j') if d_153_v0_ else _dafny.CodePoint('t')), d_370_v161_]))), d_370_v161_))), d_370_v161_)
                d_322_v131_ = rhs73_
                d_175_v18_ = rhs74_
                d_374_v165_: _dafny.Array
                nw62_ = _dafny.Array(int(0), 21)
                d_374_v165_ = nw62_
                d_375_v166_: _dafny.MultiSet
                d_375_v166_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_153_v0_, d_153_v0_]), d_154_v1_])
                index66_ = default__.safeIndex(713, (d_374_v165_).length(0))
                (d_374_v165_)[index66_] = (d_375_v166_).cardinality
                index67_ = default__.safeIndex(713, (d_374_v165_).length(0))
                (d_374_v165_)[index67_] = ((d_156_v2_) * (d_322_v131_.f19)) - (d_317_i21_)
        d_376_v167_: _dafny.Array
        nw63_ = _dafny.Array(False, 28)
        d_376_v167_ = nw63_
        d_377_v168_: D5
        d_377_v168_ = D5_DC18(d_234_v68_.f19, d_153_v0_, d_376_v167_, d_153_v0_)
        d_378_v169_: D5
        d_378_v169_ = D5_DC20(d_377_v168_)
        d_379_v170_: D5
        d_379_v170_ = D5_DC20(d_378_v169_)
        source8_ = d_379_v170_
        if source8_.is_DC18:
            d_380___mcc_h9_ = source8_.cf24
            d_381___mcc_h10_ = source8_.cf25
            d_382___mcc_h11_ = source8_.cf26
            d_383___mcc_h12_ = source8_.cf27
            d_384_cf27_ = d_383___mcc_h12_
            d_385_cf26_ = d_382___mcc_h11_
            d_386_cf25_ = d_381___mcc_h10_
            d_387_cf24_ = d_380___mcc_h9_
            index68_ = default__.safeIndex(317, (d_385_cf26_).length(0))
            (d_385_cf26_)[index68_] = d_386_cf25_
            index69_ = default__.safeIndex(317, (d_385_cf26_).length(0))
            (d_385_cf26_)[index69_] = default__.fm2((d_154_v1_) != (_dafny.SeqWithoutIsStrInference([d_386_cf25_])), d_155_globalState_)
            index70_ = default__.safeIndex(317, (d_385_cf26_).length(0))
            rhs75_ = d_231_v65_.f19
            rhs76_ = not(not (d_153_v0_) or (d_384_cf27_))
            rhs77_ = (len(d_316_v126_)) * (d_234_v68_.f19)
            rhs78_ = d_386_cf25_
            lhs65_ = d_155_globalState_
            lhs66_ = d_385_cf26_
            lhs67_ = default__.safeIndex(317, (d_385_cf26_).length(0))
            lhs68_ = d_155_globalState_
            lhs65_.f0 = rhs75_
            lhs66_[lhs67_] = rhs76_
            lhs68_.f12 = rhs77_
            d_153_v0_ = rhs78_
            d_388_v171_: _dafny.MultiSet
            d_388_v171_ = _dafny.MultiSet([(d_232_v66_) | (d_232_v66_)])
            d_388_v171_ = d_388_v171_
            d_389_v172_: D2
            d_389_v172_ = D2_DC8()
            source9_ = d_389_v172_
            if source9_.is_DC8:
                def iife34_():
                    coll10_ = _dafny.Map()
                    compr_10_: int
                    for compr_10_ in _dafny.IntegerRange(143, -7):
                        d_390_v173_: int = compr_10_
                        if ((143) <= (d_390_v173_)) and ((d_390_v173_) < (-7)):
                            coll10_[(d_390_v173_) * (d_234_v68_.f19)] = d_387_cf24_
                    return _dafny.Map(coll10_)
                (d_155_globalState_).f3 = (d_233_v67_) == ((iife34_()
                 if d_153_v0_ else d_233_v67_))
                (d_155_globalState_).f0 = 174
                d_175_v18_ = (d_175_v18_ if (d_385_cf26_)[default__.safeIndex(317, (d_385_cf26_).length(0))] else d_175_v18_)
                d_391_v174_: str
                d_391_v174_ = _dafny.CodePoint('l')
                d_391_v174_ = d_391_v174_
            elif source9_.is_DC9:
                d_392___mcc_h16_ = source9_.cf13
                d_393___mcc_h17_ = source9_.cf14
                d_394___mcc_h18_ = source9_.cf15
                d_395_cf15_ = d_394___mcc_h18_
                d_396_cf14_ = d_393___mcc_h17_
                d_397_cf13_ = d_392___mcc_h16_
                (d_231_v65_).f19 = d_397_cf13_
                (d_231_v65_).f19 = default__.safeModuloInt(d_231_v65_.f19, (d_234_v68_).fm3(d_155_globalState_))
                d_398_v175_: _dafny.Seq
                d_398_v175_ = _dafny.SeqWithoutIsStrInference([(d_154_v1_) + (d_154_v1_), d_154_v1_, d_154_v1_, _dafny.SeqWithoutIsStrInference([d_153_v0_, d_386_cf25_, d_386_cf25_]), d_154_v1_])
                d_398_v175_ = (d_398_v175_) + (d_398_v175_)
                d_399_v176_: _dafny.Array
                nw64_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                d_399_v176_ = nw64_
                d_400_v177_: str
                d_400_v177_ = _dafny.CodePoint('n')
                index71_ = default__.safeIndex(150, (d_399_v176_).length(0))
                (d_399_v176_)[index71_] = d_400_v177_
                d_401_v178_: _dafny.Map
                d_401_v178_ = _dafny.Map({_dafny.Map({d_400_v177_: d_153_v0_}): d_384_cf27_})
                d_402_v179_: D1
                d_402_v179_ = D1_DC5(d_386_cf25_)
                d_403_v180_: _dafny.Map
                d_403_v180_ = _dafny.Map({d_400_v177_: (d_402_v179_).cf10})
                index72_ = default__.safeIndex(150, (d_399_v176_).length(0))
                (d_399_v176_)[index72_] = default__.fm7(((d_401_v178_)[d_403_v180_] if (d_403_v180_) in (d_401_v178_) else not(d_384_cf27_)), 318, d_155_globalState_)
            elif source9_.is_DC7:
                d_404___mcc_h19_ = source9_.cf12
                d_405_cf12_ = d_404___mcc_h19_
                d_153_v0_ = (d_153_v0_) == (d_153_v0_)
                (d_234_v68_).f19 = ((d_230_v64_)[default__.safeIndex(d_231_v65_.f19, len(d_230_v64_))]) * (len(_dafny.SeqWithoutIsStrInference([d_234_v68_.f19])))
                d_406_v181_: int
                d_407_v182_: _dafny.Seq
                d_408_v183_: int
                d_409_v184_: _dafny.MultiSet
                out20_: int
                out21_: _dafny.Seq
                out22_: int
                out23_: _dafny.MultiSet
                out20_, out21_, out22_, out23_ = default__.m0(d_153_v0_, d_386_cf25_, d_231_v65_.f19, ((d_154_v1_)[default__.safeIndex(d_234_v68_.f19, len(d_154_v1_))] if False else d_384_cf27_), d_155_globalState_)
                d_406_v181_ = out20_
                d_407_v182_ = out21_
                d_408_v183_ = out22_
                d_409_v184_ = out23_
                d_410_v185_: str
                d_410_v185_ = _dafny.CodePoint('t')
                d_411_v186_: _dafny.Map
                d_411_v186_ = _dafny.Map({d_231_v65_: d_407_v182_})
                d_412_v187_: _dafny.Seq
                d_412_v187_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_410_v185_ for d_413_i24_ in range(default__.abs(340))]), d_175_v18_])
                d_414_v188_: _dafny.Array
                nw65_ = _dafny.Array(None, 21)
                nw65_[int(0)] = ((d_175_v18_) + (d_407_v182_)).set(default__.safeIndex(d_406_v181_, len((d_175_v18_) + (d_407_v182_))), d_410_v185_)
                nw65_[int(1)] = d_407_v182_
                nw65_[int(2)] = (d_175_v18_) + (d_407_v182_)
                nw65_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rcptp"))
                nw65_[int(4)] = ((d_411_v186_)[d_234_v68_] if (d_234_v68_) in (d_411_v186_) else d_175_v18_)
                nw65_[int(5)] = (d_412_v187_)[default__.safeIndex(d_234_v68_.f19, len(d_412_v187_))]
                nw65_[int(6)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bbyhlyjx"))) + (d_407_v182_)
                nw65_[int(7)] = (d_407_v182_) + (d_407_v182_)
                nw65_[int(8)] = d_407_v182_
                nw65_[int(9)] = (default__.fm5(d_234_v68_.f19, d_155_globalState_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xa")))
                nw65_[int(10)] = d_175_v18_
                nw65_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgjqjb"))
                nw65_[int(12)] = d_175_v18_
                nw65_[int(13)] = (d_175_v18_) + (d_175_v18_)
                nw65_[int(14)] = d_407_v182_
                nw65_[int(15)] = d_175_v18_
                nw65_[int(16)] = _dafny.SeqWithoutIsStrInference([d_410_v185_ for d_415_i25_ in range(default__.abs(576))])
                nw65_[int(17)] = d_407_v182_
                nw65_[int(18)] = default__.fm5(d_231_v65_.f19, d_155_globalState_)
                nw65_[int(19)] = (d_175_v18_).set(default__.safeIndex(d_231_v65_.f19, len(d_175_v18_)), d_410_v185_)
                nw65_[int(20)] = d_175_v18_
                d_414_v188_ = nw65_
                index73_ = default__.safeIndex(502, (d_414_v188_).length(0))
                (d_414_v188_)[index73_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tb"))
                d_416_v189_: _dafny.Array
                nw66_ = _dafny.Array(_dafny.Seq({}), 17)
                d_416_v189_ = nw66_
                d_417_v190_: _dafny.Map
                d_417_v190_ = _dafny.Map({d_386_cf25_: d_416_v189_})
                index74_ = default__.safeIndex(502, (d_414_v188_).length(0))
                rhs79_ = d_407_v182_
                rhs80_ = d_405_cf12_
                rhs81_ = (d_384_cf27_ if False else (d_385_cf26_)[default__.safeIndex(317, (d_385_cf26_).length(0))])
                rhs82_ = ((d_417_v190_)[d_386_cf25_] if (d_386_cf25_) in (d_417_v190_) else d_416_v189_)
                lhs69_ = d_414_v188_
                lhs70_ = default__.safeIndex(502, (d_414_v188_).length(0))
                lhs69_[lhs70_] = rhs79_
                d_405_cf12_ = rhs80_
                d_153_v0_ = rhs81_
                d_416_v189_ = rhs82_
            elif True:
                d_418___mcc_h20_ = source9_.cf16
                d_419_cf16_ = d_418___mcc_h20_
                (d_155_globalState_).f3 = d_384_cf27_
                (d_155_globalState_).f0 = d_231_v65_.f19
                d_386_cf25_ = ((d_154_v1_)[default__.safeIndex(656, len(d_154_v1_))]) and (True)
                d_420_v191_: _dafny.MultiSet
                d_420_v191_ = _dafny.MultiSet([d_154_v1_, d_154_v1_, d_154_v1_])
                d_420_v191_ = (d_420_v191_).intersection((d_420_v191_).intersection(d_420_v191_))
        elif source8_.is_DC19:
            d_421___mcc_h13_ = source8_.cf28
            d_422_cf28_ = d_421___mcc_h13_
            d_423_v192_: str
            d_423_v192_ = _dafny.CodePoint('w')
            d_423_v192_ = (d_175_v18_)[default__.safeIndex(d_156_v2_, len(d_175_v18_))]
            d_424_v193_: _dafny.MultiSet
            d_424_v193_ = _dafny.MultiSet([d_156_v2_])
            if (d_424_v193_).ispropersubset(d_424_v193_):
                d_316_v126_ = d_316_v126_
                d_425_v194_: C0
                nw67_ = C0()
                nw67_.ctor__(d_234_v68_.f19)
                d_425_v194_ = nw67_
                d_426_v195_: _dafny.Set
                d_426_v195_ = _dafny.Set({d_175_v18_})
                d_426_v195_ = default__.fm9((0) - (d_234_v68_.f19), _dafny.CodePoint('u'), default__.fm2(d_153_v0_, d_155_globalState_), d_155_globalState_)
                d_231_v65_ = d_231_v65_
                d_422_cf28_ = d_153_v0_
            elif True:
                (d_155_globalState_).f3 = ((d_422_cf28_) or (d_153_v0_)) or (False)
                (d_155_globalState_).f3 = d_422_cf28_
                d_427_v196_: int
                d_428_v197_: _dafny.Seq
                d_429_v198_: int
                d_430_v199_: _dafny.MultiSet
                out24_: int
                out25_: _dafny.Seq
                out26_: int
                out27_: _dafny.MultiSet
                out24_, out25_, out26_, out27_ = default__.m0(d_422_cf28_, (d_231_v65_.f19) not in (d_230_v64_), d_231_v65_.f19, (d_234_v68_.f19) < (len(_dafny.Map({(0) - (d_234_v68_.f19): d_423_v192_}))), d_155_globalState_)
                d_427_v196_ = out24_
                d_428_v197_ = out25_
                d_429_v198_ = out26_
                d_430_v199_ = out27_
                d_431_v200_: _dafny.Array
                nw68_ = _dafny.Array(_dafny.Map({}), 23)
                d_431_v200_ = nw68_
                index75_ = default__.safeIndex(519, (d_431_v200_).length(0))
                (d_431_v200_)[index75_] = d_233_v67_
                d_432_v201_: _dafny.Map
                d_432_v201_ = _dafny.Map({d_156_v2_: (d_233_v67_) | (d_233_v67_)})
                index76_ = default__.safeIndex(519, (d_431_v200_).length(0))
                (d_431_v200_)[index76_] = ((d_432_v201_)[d_427_v196_] if (d_427_v196_) in (d_432_v201_) else d_233_v67_)
                d_376_v167_ = d_376_v167_
            d_433_v202_: _dafny.Map
            d_433_v202_ = _dafny.Map({d_153_v0_: _dafny.Set({d_422_cf28_})})
            d_434_v203_: D0
            d_434_v203_ = D0_DC1(len(((d_433_v202_)[True] if (True) in (d_433_v202_) else d_316_v126_)), d_234_v68_.f19, d_153_v0_)
            d_435_v204_: C0
            nw69_ = C0()
            nw69_.ctor__((d_434_v203_).cf1)
            d_435_v204_ = nw69_
            d_436_v205_: _dafny.Array
            nw70_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 12)
            d_436_v205_ = nw70_
            d_437_v206_: _dafny.MultiSet
            d_437_v206_ = _dafny.MultiSet([d_233_v67_])
            d_438_v207_: _dafny.Map
            d_438_v207_ = _dafny.Map({d_436_v205_: d_437_v206_})
            d_438_v207_ = (d_438_v207_).set(d_436_v205_, (d_437_v206_).set(default__.fm10(d_155_globalState_), default__.abs(d_156_v2_)))
        elif source8_.is_DC17:
            d_439___mcc_h14_ = source8_.cf23
            d_440_cf23_ = d_439___mcc_h14_
            d_156_v2_ = (d_234_v68_.f19) - (d_234_v68_.f19)
            d_441_v208_: _dafny.Map
            d_441_v208_ = _dafny.Map({d_156_v2_: True})
            d_153_v0_ = ((d_441_v208_)[(d_157_v3_).cardinality] if ((d_157_v3_).cardinality) in (d_441_v208_) else d_153_v0_)
            d_442_v209_: _dafny.MultiSet
            d_442_v209_ = _dafny.MultiSet([(0) - (d_234_v68_.f19)])
            if (D0_DC1(d_156_v2_, ((d_442_v209_)[len(d_175_v18_)] if (len(d_175_v18_)) in (d_442_v209_) else d_231_v65_.f19), d_153_v0_)).cf3:
                d_230_v64_ = d_230_v64_
                (d_155_globalState_).f3 = d_153_v0_
                index77_ = default__.safeIndex(647, (d_376_v167_).length(0))
                (d_376_v167_)[index77_] = not(d_153_v0_)
                d_443_v210_: _dafny.Array
                nw71_ = _dafny.Array(int(0), 26)
                d_443_v210_ = nw71_
                index78_ = default__.safeIndex(819, (d_443_v210_).length(0))
                (d_443_v210_)[index78_] = default__.safeModuloInt(d_231_v65_.f19, len(_dafny.SeqWithoutIsStrInference([d_440_cf23_ for d_444_i26_ in range(default__.abs(280))])))
                d_445_v211_: _dafny.Set
                d_445_v211_ = _dafny.Set({d_234_v68_})
                index79_ = default__.safeIndex(647, (d_376_v167_).length(0))
                index80_ = default__.safeIndex(819, (d_443_v210_).length(0))
                rhs83_ = True
                rhs84_ = not(default__.fm2(d_153_v0_, d_155_globalState_))
                rhs85_ = default__.safeModuloInt(d_231_v65_.f19, -237)
                rhs86_ = ((d_175_v18_).set(default__.safeIndex((0) - (default__.safeDivisionInt(len(d_175_v18_), -470)), len(d_175_v18_)), d_440_cf23_)).set(default__.safeIndex((d_442_v209_).cardinality, len((d_175_v18_).set(default__.safeIndex((0) - (default__.safeDivisionInt(len(d_175_v18_), -470)), len(d_175_v18_)), d_440_cf23_))), (d_175_v18_)[default__.safeIndex((d_442_v209_).cardinality, len(d_175_v18_))])
                rhs87_ = (d_445_v211_).isdisjoint(d_445_v211_)
                lhs71_ = d_376_v167_
                lhs72_ = default__.safeIndex(647, (d_376_v167_).length(0))
                lhs73_ = d_155_globalState_
                lhs74_ = d_443_v210_
                lhs75_ = default__.safeIndex(819, (d_443_v210_).length(0))
                lhs71_[lhs72_] = rhs83_
                lhs73_.f3 = rhs84_
                lhs74_[lhs75_] = rhs85_
                d_175_v18_ = rhs86_
                d_153_v0_ = rhs87_
                (d_155_globalState_).f3 = not((d_153_v0_ if d_153_v0_ else (d_376_v167_)[default__.safeIndex(647, (d_376_v167_).length(0))]))
                d_446_v212_: D2
                d_446_v212_ = D2_DC9(d_231_v65_.f19, d_234_v68_, d_234_v68_)
                d_447_v213_: D2
                d_447_v213_ = D2_DC10(d_446_v212_)
                pat_let_tv21_ = d_446_v212_
                def iife35_(_pat_let12_0):
                    def iife36_(d_448_dt__update__tmp_h3_):
                        def iife37_(_pat_let13_0):
                            def iife38_(d_449_dt__update_hcf16_h0_):
                                return D2_DC10(d_449_dt__update_hcf16_h0_)
                            return iife38_(_pat_let13_0)
                        return iife37_(pat_let_tv21_)
                    return iife36_(_pat_let12_0)
                d_447_v213_ = iife35_(d_447_v213_)
            elif True:
                d_450_v214_: _dafny.Array
                def lambda20_(d_451_v2_):
                    def lambda21_(d_452_i27_):
                        return default__.safeDivisionInt(d_452_i27_, (0) - (d_451_v2_))

                    return lambda21_

                init10_ = lambda20_(d_156_v2_)
                nw72_ = _dafny.Array(None, 2)
                for i0_10_ in range(nw72_.length(0)):
                    nw72_[i0_10_] = init10_(i0_10_)
                d_450_v214_ = nw72_
                d_450_v214_ = d_450_v214_
                d_453_v215_: int
                d_454_v216_: _dafny.Seq
                d_455_v217_: int
                d_456_v218_: _dafny.MultiSet
                out28_: int
                out29_: _dafny.Seq
                out30_: int
                out31_: _dafny.MultiSet
                out28_, out29_, out30_, out31_ = default__.m0(False, d_153_v0_, len((d_441_v208_) | (d_441_v208_)), False, d_155_globalState_)
                d_453_v215_ = out28_
                d_454_v216_ = out29_
                d_455_v217_ = out30_
                d_456_v218_ = out31_
                d_457_v219_: _dafny.Set
                d_457_v219_ = _dafny.Set({d_455_v217_, len(d_454_v216_)})
                d_458_v220_: _dafny.Seq
                d_458_v220_ = _dafny.SeqWithoutIsStrInference([default__.fm5(d_231_v65_.f19, d_155_globalState_), d_175_v18_])
                (d_155_globalState_).f3 = (len(d_457_v219_)) <= (default__.fm0(d_234_v68_.f19, d_440_cf23_, d_458_v220_, d_231_v65_.f19, d_155_globalState_))
                d_233_v67_ = (d_233_v67_).set(len(d_175_v18_), d_231_v65_.f19)
                d_376_v167_ = d_376_v167_
            d_459_v221_: _dafny.Array
            def lambda22_(d_460_v65_):
                def lambda23_(d_461_i28_):
                    return (d_461_i28_) + (d_460_v65_.f19)

                return lambda23_

            init11_ = lambda22_(d_231_v65_)
            nw73_ = _dafny.Array(None, 24)
            for i0_11_ in range(nw73_.length(0)):
                nw73_[i0_11_] = init11_(i0_11_)
            d_459_v221_ = nw73_
            d_459_v221_ = d_459_v221_
        elif True:
            d_462___mcc_h15_ = source8_.cf29
            d_463_cf29_ = d_462___mcc_h15_
            d_464_v222_: _dafny.MultiSet
            d_464_v222_ = _dafny.MultiSet([d_231_v65_])
            d_464_v222_ = d_464_v222_
            (d_155_globalState_).f3 = (d_175_v18_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eevhh")))
            d_465_v223_: D2
            d_465_v223_ = D2_DC9(d_231_v65_.f19, d_234_v68_, d_234_v68_)
            nw74_ = C0()
            nw74_.ctor__((d_465_v223_).cf13)
            d_234_v68_ = nw74_
            d_466_v224_: str
            d_466_v224_ = _dafny.CodePoint('i')
            d_467_v225_: _dafny.Seq
            d_467_v225_ = _dafny.SeqWithoutIsStrInference([d_175_v18_, _dafny.SeqWithoutIsStrInference([d_466_v224_, d_466_v224_])])
            d_468_v226_: _dafny.Map
            d_468_v226_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_466_v224_ for d_469_i29_ in range(default__.abs(777))]): d_234_v68_.f19})
            d_466_v224_ = default__.fm7(d_153_v0_, default__.fm0(d_231_v65_.f19, d_466_v224_, d_467_v225_, ((d_468_v226_)[_dafny.SeqWithoutIsStrInference([d_466_v224_ for d_470_i30_ in range(default__.abs(889))])] if (_dafny.SeqWithoutIsStrInference([d_466_v224_ for d_471_i30_ in range(default__.abs(889))])) in (d_468_v226_) else len(_dafny.SeqWithoutIsStrInference([d_153_v0_, d_153_v0_]))), d_155_globalState_), d_155_globalState_)
        d_472_v227_: str
        d_472_v227_ = _dafny.CodePoint('s')
        d_473_v228_: _dafny.Array
        nw75_ = _dafny.Array(None, 4)
        nw75_[int(0)] = d_472_v227_
        nw75_[int(1)] = d_472_v227_
        nw75_[int(2)] = (d_175_v18_)[default__.safeIndex(d_234_v68_.f19, len(d_175_v18_))]
        nw75_[int(3)] = (d_472_v227_ if False else d_472_v227_)
        d_473_v228_ = nw75_
        index81_ = default__.safeIndex(241, (d_473_v228_).length(0))
        (d_473_v228_)[index81_] = default__.fm7(True, 157, d_155_globalState_)
        index82_ = default__.safeIndex(241, (d_473_v228_).length(0))
        (d_473_v228_)[index82_] = d_472_v227_
        d_156_v2_ = d_156_v2_
        hi3_ = d_156_v2_
        for d_474_i31_ in range((len(d_230_v64_)) + (d_156_v2_), hi3_):
            d_175_v18_ = _dafny.SeqWithoutIsStrInference([(d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))] for d_475_i32_ in range(default__.abs(-856))])
            if d_153_v0_:
                d_476_v229_: _dafny.Seq
                d_476_v229_ = _dafny.SeqWithoutIsStrInference([d_175_v18_])
                d_477_v230_: D0
                d_477_v230_ = D0_DC1(default__.fm0(d_234_v68_.f19, _dafny.CodePoint('m'), d_476_v229_, d_156_v2_, d_155_globalState_), (d_157_v3_).cardinality, d_153_v0_)
                d_478_v231_: _dafny.Map
                d_478_v231_ = _dafny.Map({D5_DC17((d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))]): d_477_v230_})
                d_479_v233_: D5
                d_479_v233_ = D5_DC17((d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))])
                pat_let_tv22_ = d_472_v227_
                d_480_v234_: _dafny.Set
                def iife39_(_pat_let14_0):
                    def iife40_(d_481_dt__update__tmp_h4_):
                        def iife41_(_pat_let15_0):
                            def iife42_(d_482_dt__update_hcf23_h0_):
                                return D5_DC17(d_482_dt__update_hcf23_h0_)
                            return iife42_(_pat_let15_0)
                        return iife41_(pat_let_tv22_)
                    return iife40_(_pat_let14_0)
                d_480_v234_ = _dafny.Set({d_479_v233_, d_479_v233_, iife39_(d_479_v233_)})
                d_483_v235_: _dafny.Seq
                def iife43_():
                    coll11_ = _dafny.Map()
                    compr_11_: D5
                    for compr_11_ in (d_480_v234_).Elements:
                        d_484_v232_: D5 = compr_11_
                        if (d_484_v232_) in (d_480_v234_):
                            coll11_[d_484_v232_] = d_477_v230_
                    return _dafny.Map(coll11_)
                d_483_v235_ = _dafny.SeqWithoutIsStrInference([d_478_v231_, iife43_()
                , (d_478_v231_) | (d_478_v231_)])
                d_483_v235_ = ((d_483_v235_) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_479_v233_: d_477_v230_}) for d_485_i33_ in range(default__.abs(204))]))) + (((d_483_v235_) + (d_483_v235_)).set(default__.safeIndex(583, len((d_483_v235_) + (d_483_v235_))), _dafny.Map({d_479_v233_: d_477_v230_})))
                (d_231_v65_).f19 = d_156_v2_
                d_486_v236_: D5
                d_486_v236_ = D5_DC18(d_156_v2_, d_153_v0_, d_376_v167_, d_153_v0_)
                d_487_v237_: _dafny.Map
                d_487_v237_ = _dafny.Map({(d_486_v236_).cf26: (d_153_v0_ if d_153_v0_ else True)})
                d_488_v238_: _dafny.Map
                d_488_v238_ = _dafny.Map({d_153_v0_: _dafny.SeqWithoutIsStrInference([d_234_v68_.f19 for d_489_i34_ in range(default__.abs(218))])})
                d_490_v239_: _dafny.MultiSet
                d_490_v239_ = _dafny.MultiSet([(d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))], d_472_v227_])
                d_491_v240_: _dafny.Array
                def lambda24_(d_492_i35_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvlqdof"))

                init12_ = lambda24_
                nw76_ = _dafny.Array(None, 18)
                for i0_12_ in range(nw76_.length(0)):
                    nw76_[i0_12_] = init12_(i0_12_)
                d_491_v240_ = nw76_
                d_493_v241_: _dafny.Map
                d_493_v241_ = _dafny.Map({d_491_v240_: d_153_v0_})
                index83_ = default__.safeIndex(241, (d_473_v228_).length(0))
                rhs88_ = _dafny.Map({d_376_v167_: (d_154_v1_)[default__.safeIndex(d_234_v68_.f19, len(d_154_v1_))]})
                rhs89_ = (_dafny.CodePoint('h') if d_153_v0_ else (d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))])
                rhs90_ = (((d_488_v238_)[d_153_v0_] if (d_153_v0_) in (d_488_v238_) else _dafny.SeqWithoutIsStrInference([(0) - (d_234_v68_.f19)]))) == (((d_230_v64_) + (d_230_v64_)).set(default__.safeIndex((d_230_v64_)[default__.safeIndex((0) - (((d_490_v239_)[(d_175_v18_)[default__.safeIndex(d_234_v68_.f19, len(d_175_v18_))]] if ((d_175_v18_)[default__.safeIndex(d_234_v68_.f19, len(d_175_v18_))]) in (d_490_v239_) else (0) - (d_474_i31_))), len(d_230_v64_))], len((d_230_v64_) + (d_230_v64_))), (0) - ((d_157_v3_).cardinality)))
                rhs91_ = ((d_493_v241_)[d_491_v240_] if (d_491_v240_) in (d_493_v241_) else d_153_v0_)
                lhs76_ = d_473_v228_
                lhs77_ = default__.safeIndex(241, (d_473_v228_).length(0))
                lhs78_ = d_155_globalState_
                lhs79_ = d_155_globalState_
                d_487_v237_ = rhs88_
                lhs76_[lhs77_] = rhs89_
                lhs78_.f3 = rhs90_
                lhs79_.f3 = rhs91_
                (d_155_globalState_).f8 = d_231_v65_.f19
                d_494_v242_: _dafny.MultiSet
                d_494_v242_ = _dafny.MultiSet([d_231_v65_.f19])
                d_495_v243_: D0
                d_495_v243_ = D0_DC0((d_234_v68_).fm4(len(d_230_v64_), d_234_v68_.f19, d_153_v0_, d_153_v0_, d_155_globalState_))
                d_496_v244_: _dafny.Array
                nw77_ = _dafny.Array(None, 8)
                nw77_[int(0)] = (d_230_v64_)[default__.safeIndex(d_234_v68_.f19, len(d_230_v64_))]
                nw77_[int(1)] = (d_234_v68_).fm3(d_155_globalState_)
                nw77_[int(2)] = d_156_v2_
                nw77_[int(3)] = (d_157_v3_).cardinality
                nw77_[int(4)] = len(d_175_v18_)
                nw77_[int(5)] = d_231_v65_.f19
                nw77_[int(6)] = ((_dafny.MultiSet(d_230_v64_)) - (d_494_v242_)).cardinality
                nw77_[int(7)] = (d_474_i31_ if d_153_v0_ else (d_495_v243_).cf0)
                d_496_v244_ = nw77_
                index84_ = default__.safeIndex(203, (d_496_v244_).length(0))
                (d_496_v244_)[index84_] = d_156_v2_
                d_497_v245_: _dafny.MultiSet
                d_497_v245_ = _dafny.MultiSet([d_234_v68_])
                d_498_v246_: _dafny.Map
                d_498_v246_ = _dafny.Map({d_497_v245_: d_474_i31_})
                index85_ = default__.safeIndex(203, (d_496_v244_).length(0))
                (d_496_v244_)[index85_] = (0) - (((d_498_v246_)[d_497_v245_] if (d_497_v245_) in (d_498_v246_) else default__.safeModuloInt(len(d_230_v64_), d_156_v2_)))
            elif True:
                d_499_v247_: _dafny.Array
                nw78_ = _dafny.Array(None, 10)
                nw78_[int(0)] = d_231_v65_.f19
                nw78_[int(1)] = d_474_i31_
                nw78_[int(2)] = d_231_v65_.f19
                nw78_[int(3)] = d_234_v68_.f19
                nw78_[int(4)] = (586) + (d_234_v68_.f19)
                nw78_[int(5)] = d_231_v65_.f19
                nw78_[int(6)] = len((d_154_v1_) + (d_154_v1_))
                nw78_[int(7)] = len(_dafny.SeqWithoutIsStrInference([d_474_i31_, d_231_v65_.f19]))
                nw78_[int(8)] = d_231_v65_.f19
                nw78_[int(9)] = default__.safeModuloInt(d_231_v65_.f19, d_234_v68_.f19)
                d_499_v247_ = nw78_
                index86_ = default__.safeIndex(208, (d_499_v247_).length(0))
                (d_499_v247_)[index86_] = (0) - (d_231_v65_.f19)
                index87_ = default__.safeIndex(208, (d_499_v247_).length(0))
                (d_499_v247_)[index87_] = default__.safeModuloInt(d_474_i31_, (0) - ((d_474_i31_) + (-138)))
                index88_ = default__.safeIndex(262, (d_376_v167_).length(0))
                (d_376_v167_)[index88_] = d_153_v0_
                d_500_v248_: _dafny.Seq
                d_500_v248_ = _dafny.SeqWithoutIsStrInference([d_231_v65_])
                index89_ = default__.safeIndex(262, (d_376_v167_).length(0))
                (d_376_v167_)[index89_] = (d_154_v1_)[default__.safeIndex(len(d_500_v248_), len(d_154_v1_))]
                (d_155_globalState_).f8 = (len(_dafny.SeqWithoutIsStrInference([d_153_v0_]))) + ((d_230_v64_)[default__.safeIndex(d_234_v68_.f19, len(d_230_v64_))])
                (d_155_globalState_).f3 = not(((d_376_v167_)[default__.safeIndex(262, (d_376_v167_).length(0))] if (d_376_v167_)[default__.safeIndex(262, (d_376_v167_).length(0))] else (d_153_v0_) and (not((d_376_v167_)[default__.safeIndex(262, (d_376_v167_).length(0))]))))
                d_501_v249_: C0
                nw79_ = C0()
                nw79_.ctor__(d_234_v68_.f19)
                d_501_v249_ = nw79_
            d_502_v250_: _dafny.Map
            d_502_v250_ = _dafny.Map({d_234_v68_.f19: (d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))]})
            d_503_v251_: D5
            d_503_v251_ = D5_DC18(len(d_502_v250_), default__.fm2(d_153_v0_, d_155_globalState_), d_376_v167_, not(d_153_v0_))
            source10_ = d_503_v251_
            if source10_.is_DC18:
                d_504___mcc_h21_ = source10_.cf24
                d_505___mcc_h22_ = source10_.cf25
                d_506___mcc_h23_ = source10_.cf26
                d_507___mcc_h24_ = source10_.cf27
                d_508_cf27_ = d_507___mcc_h24_
                d_509_cf26_ = d_506___mcc_h23_
                d_510_cf25_ = d_505___mcc_h22_
                d_511_cf24_ = d_504___mcc_h21_
                d_512_v252_: _dafny.Map
                d_512_v252_ = _dafny.Map({d_156_v2_: d_508_cf27_})
                d_512_v252_ = (d_512_v252_).set((d_231_v65_).fm3(d_155_globalState_), d_153_v0_)
                index90_ = default__.safeIndex(864, (d_509_cf26_).length(0))
                (d_509_cf26_)[index90_] = d_510_cf25_
                index91_ = default__.safeIndex(228, (d_376_v167_).length(0))
                (d_376_v167_)[index91_] = default__.fm2(True, d_155_globalState_)
                d_513_v253_: _dafny.Array
                nw80_ = _dafny.Array(None, 7)
                nw80_[int(0)] = d_175_v18_
                nw80_[int(1)] = d_175_v18_
                nw80_[int(2)] = d_175_v18_
                nw80_[int(3)] = d_175_v18_
                nw80_[int(4)] = d_175_v18_
                nw80_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqsphj"))
                nw80_[int(6)] = _dafny.SeqWithoutIsStrInference([d_472_v227_ for d_514_i36_ in range(default__.abs(579))])
                d_513_v253_ = nw80_
                d_515_v254_: D3
                d_515_v254_ = D3_DC12(d_513_v253_, d_510_cf25_)
                pat_let_tv23_ = d_508_cf27_
                index92_ = default__.safeIndex(864, (d_509_cf26_).length(0))
                index93_ = default__.safeIndex(228, (d_376_v167_).length(0))
                def iife44_(_pat_let16_0):
                    def iife45_(d_516_dt__update__tmp_h5_):
                        def iife46_(_pat_let17_0):
                            def iife47_(d_517_dt__update_hcf19_h0_):
                                return D3_DC12((d_516_dt__update__tmp_h5_).cf18, d_517_dt__update_hcf19_h0_)
                            return iife47_(_pat_let17_0)
                        return iife46_(pat_let_tv23_)
                    return iife45_(_pat_let16_0)
                rhs92_ = ((d_512_v252_)[(d_230_v64_)[default__.safeIndex(d_234_v68_.f19, len(d_230_v64_))]] if ((d_230_v64_)[default__.safeIndex(d_234_v68_.f19, len(d_230_v64_))]) in (d_512_v252_) else d_510_cf25_)
                rhs93_ = (default__.fm2(not(d_153_v0_), d_155_globalState_)) == ((d_472_v227_) not in (d_175_v18_))
                rhs94_ = (d_231_v65_).fm4(d_234_v68_.f19, (d_231_v65_).fm4(d_231_v65_.f19, d_234_v68_.f19, (iife44_(d_515_v254_)).cf19, d_153_v0_, d_155_globalState_), d_510_cf25_, d_153_v0_, d_155_globalState_)
                rhs95_ = d_153_v0_
                rhs96_ = d_510_cf25_
                lhs80_ = d_509_cf26_
                lhs81_ = default__.safeIndex(864, (d_509_cf26_).length(0))
                lhs82_ = d_234_v68_
                lhs83_ = d_376_v167_
                lhs84_ = default__.safeIndex(228, (d_376_v167_).length(0))
                lhs80_[lhs81_] = rhs92_
                d_153_v0_ = rhs93_
                lhs82_.f19 = rhs94_
                lhs83_[lhs84_] = rhs95_
                d_510_cf25_ = rhs96_
                (d_155_globalState_).f2 = d_231_v65_.f19
                d_232_v66_ = d_232_v66_
            elif source10_.is_DC19:
                d_518___mcc_h25_ = source10_.cf28
                d_519_cf28_ = d_518___mcc_h25_
                d_520_v255_: _dafny.Map
                d_520_v255_ = _dafny.Map({default__.fm2(not(d_519_cf28_), d_155_globalState_): d_234_v68_.f19})
                d_521_v256_: _dafny.Array
                nw81_ = _dafny.Array(None, 13)
                nw81_[int(0)] = default__.safeModuloInt(d_474_i31_, d_234_v68_.f19)
                nw81_[int(1)] = (0) - (d_474_i31_)
                nw81_[int(2)] = 777
                nw81_[int(3)] = d_231_v65_.f19
                nw81_[int(4)] = (d_474_i31_) - (d_474_i31_)
                nw81_[int(5)] = d_231_v65_.f19
                nw81_[int(6)] = ((d_520_v255_)[d_153_v0_] if (d_153_v0_) in (d_520_v255_) else d_234_v68_.f19)
                nw81_[int(7)] = (d_231_v65_).fm4(d_234_v68_.f19, (0) - (d_234_v68_.f19), d_153_v0_, d_153_v0_, d_155_globalState_)
                nw81_[int(8)] = d_156_v2_
                nw81_[int(9)] = 324
                nw81_[int(10)] = d_231_v65_.f19
                nw81_[int(11)] = (55) + (d_231_v65_.f19)
                nw81_[int(12)] = (d_231_v65_.f19 if d_153_v0_ else d_231_v65_.f19)
                d_521_v256_ = nw81_
                index94_ = default__.safeIndex(543, (d_521_v256_).length(0))
                (d_521_v256_)[index94_] = len((d_520_v255_) | (d_520_v255_))
                index95_ = default__.safeIndex(543, (d_521_v256_).length(0))
                (d_521_v256_)[index95_] = d_234_v68_.f19
                d_522_v257_: _dafny.Map
                d_522_v257_ = _dafny.Map({d_231_v65_.f19: d_519_cf28_})
                d_522_v257_ = (d_522_v257_).set(223, (d_519_cf28_ if d_519_cf28_ else d_153_v0_))
                index96_ = default__.safeIndex(829, (d_376_v167_).length(0))
                (d_376_v167_)[index96_] = (d_234_v68_.f19) < (d_234_v68_.f19)
                index97_ = default__.safeIndex(829, (d_376_v167_).length(0))
                (d_376_v167_)[index97_] = (True) == (d_519_cf28_)
                d_523_v258_: _dafny.Array
                nw82_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_523_v258_ = nw82_
                nw83_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_523_v258_ = nw83_
            elif source10_.is_DC17:
                d_524___mcc_h26_ = source10_.cf23
                d_525_cf23_ = d_524___mcc_h26_
                d_526_v259_: _dafny.Map
                d_526_v259_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))).set(default__.safeIndex((0) - (d_234_v68_.f19), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))), _dafny.CodePoint('b')): not(d_153_v0_)})
                def iife48_():
                    coll12_ = _dafny.Map()
                    compr_12_: _dafny.Seq
                    for compr_12_ in (d_526_v259_).keys.Elements:
                        d_527_v260_: _dafny.Seq = compr_12_
                        if (d_527_v260_) in (d_526_v259_):
                            coll12_[d_527_v260_] = d_153_v0_
                    return _dafny.Map(coll12_)
                d_526_v259_ = (iife48_()
                ) | ((d_526_v259_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mnlnnbyv")), not(not(d_153_v0_))))
                d_528_v261_: int
                d_529_v262_: _dafny.Seq
                d_530_v263_: int
                d_531_v264_: _dafny.MultiSet
                out32_: int
                out33_: _dafny.Seq
                out34_: int
                out35_: _dafny.MultiSet
                out32_, out33_, out34_, out35_ = default__.m0((d_231_v65_.f19) in (d_230_v64_), d_153_v0_, d_474_i31_, not(d_153_v0_), d_155_globalState_)
                d_528_v261_ = out32_
                d_529_v262_ = out33_
                d_530_v263_ = out34_
                d_531_v264_ = out35_
                d_532_v265_: _dafny.Map
                d_532_v265_ = _dafny.Map({d_528_v261_: d_153_v0_})
                d_533_v266_: _dafny.Map
                d_533_v266_ = _dafny.Map({(d_532_v265_).set(d_156_v2_, True): (not(True)) and (d_153_v0_)})
                d_533_v266_ = (d_533_v266_) | (d_533_v266_)
                d_534_v267_: _dafny.Set
                d_534_v267_ = _dafny.Set({d_376_v167_})
                d_534_v267_ = ((d_534_v267_) - (d_534_v267_)) | (d_534_v267_)
            elif True:
                d_535___mcc_h27_ = source10_.cf29
                d_536_cf29_ = d_535___mcc_h27_
                d_537_v268_: int
                d_538_v269_: _dafny.Seq
                d_539_v270_: int
                d_540_v271_: _dafny.MultiSet
                out36_: int
                out37_: _dafny.Seq
                out38_: int
                out39_: _dafny.MultiSet
                out36_, out37_, out38_, out39_ = default__.m0((default__.fm5((0) - (d_474_i31_), d_155_globalState_)) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmxvhnl"))), default__.fm2(d_153_v0_, d_155_globalState_), d_231_v65_.f19, d_153_v0_, d_155_globalState_)
                d_537_v268_ = out36_
                d_538_v269_ = out37_
                d_539_v270_ = out38_
                d_540_v271_ = out39_
                d_541_v272_: D2
                d_541_v272_ = D2_DC9(d_231_v65_.f19, d_234_v68_, d_234_v68_)
                d_542_v273_: D2
                d_542_v273_ = D2_DC10(d_541_v272_)
                d_543_v274_: _dafny.Map
                d_543_v274_ = _dafny.Map({d_231_v65_.f19: d_153_v0_})
                d_544_v275_: D3
                d_544_v275_ = D3_DC11(d_543_v274_)
                d_545_v276_: _dafny.MultiSet
                d_545_v276_ = _dafny.MultiSet([d_544_v275_, D3_DC11(d_543_v274_)])
                rhs97_ = d_542_v273_
                rhs98_ = (676 if (d_153_v0_) or (d_153_v0_) else d_539_v270_)
                rhs99_ = (d_545_v276_).set(d_544_v275_, default__.abs((0) - (d_231_v65_.f19)))
                lhs85_ = d_231_v65_
                d_542_v273_ = rhs97_
                lhs85_.f19 = rhs98_
                d_545_v276_ = rhs99_
                d_231_v65_ = ((d_231_v65_ if d_153_v0_ else d_231_v65_) if d_153_v0_ else d_231_v65_)
                d_546_v277_: _dafny.Array
                def lambda25_(d_547_i37_):
                    return (d_547_i37_) * (936)

                init13_ = lambda25_
                nw84_ = _dafny.Array(None, 25)
                for i0_13_ in range(nw84_.length(0)):
                    nw84_[i0_13_] = init13_(i0_13_)
                d_546_v277_ = nw84_
                index98_ = default__.safeIndex(385, (d_546_v277_).length(0))
                (d_546_v277_)[index98_] = d_156_v2_
                index99_ = default__.safeIndex(894, (d_546_v277_).length(0))
                (d_546_v277_)[index99_] = d_234_v68_.f19
                d_548_v278_: _dafny.Map
                d_548_v278_ = _dafny.Map({d_538_v269_: d_153_v0_})
                d_549_v279_: _dafny.Set
                d_549_v279_ = _dafny.Set({(d_230_v64_)[default__.safeIndex(len(d_548_v278_), len(d_230_v64_))]})
                d_550_v280_: _dafny.Map
                d_550_v280_ = _dafny.Map({d_153_v0_: (d_156_v2_) - (d_537_v268_)})
                index100_ = default__.safeIndex(385, (d_546_v277_).length(0))
                index101_ = default__.safeIndex(894, (d_546_v277_).length(0))
                rhs100_ = d_156_v2_
                rhs101_ = ((d_550_v280_)[True] if (True) in (d_550_v280_) else d_234_v68_.f19)
                rhs102_ = d_153_v0_
                rhs103_ = d_474_i31_
                rhs104_ = d_549_v279_
                lhs86_ = d_546_v277_
                lhs87_ = default__.safeIndex(385, (d_546_v277_).length(0))
                lhs88_ = d_155_globalState_
                lhs89_ = d_155_globalState_
                lhs90_ = d_546_v277_
                lhs91_ = default__.safeIndex(894, (d_546_v277_).length(0))
                lhs86_[lhs87_] = rhs100_
                lhs88_.f8 = rhs101_
                lhs89_.f3 = rhs102_
                lhs90_[lhs91_] = rhs103_
                d_549_v279_ = rhs104_
            (d_155_globalState_).f8 = (d_234_v68_.f19) + (len(_dafny.SeqWithoutIsStrInference([d_472_v227_ for d_551_i38_ in range(default__.abs(74))])))
        d_552_v281_: D0
        d_552_v281_ = D0_DC1((_dafny.MultiSet([d_153_v0_])).cardinality, d_156_v2_, not(d_153_v0_))
        source11_ = ((d_552_v281_ if d_153_v0_ else d_552_v281_) if not(d_153_v0_) else d_552_v281_)
        if source11_.is_DC1:
            d_553___mcc_h28_ = source11_.cf1
            d_554___mcc_h29_ = source11_.cf2
            d_555___mcc_h30_ = source11_.cf3
            d_556_cf3_ = d_555___mcc_h30_
            d_557_cf2_ = d_554___mcc_h29_
            d_558_cf1_ = d_553___mcc_h28_
            index102_ = default__.safeIndex(704, (d_376_v167_).length(0))
            (d_376_v167_)[index102_] = d_556_cf3_
            index103_ = default__.safeIndex(704, (d_376_v167_).length(0))
            (d_376_v167_)[index103_] = d_556_cf3_
            if d_556_cf3_:
                d_559_v282_: _dafny.Array
                nw85_ = _dafny.Array(_dafny.Map({}), 19)
                d_559_v282_ = nw85_
                index104_ = default__.safeIndex(648, (d_559_v282_).length(0))
                (d_559_v282_)[index104_] = _dafny.Map({(d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))]: d_153_v0_})
                d_560_v283_: _dafny.Seq
                d_560_v283_ = _dafny.SeqWithoutIsStrInference([d_234_v68_, d_234_v68_])
                d_561_v284_: _dafny.Map
                d_561_v284_ = _dafny.Map({(d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))]: ((0) - (d_558_cf1_)) < (len(d_560_v283_))})
                index105_ = default__.safeIndex(648, (d_559_v282_).length(0))
                (d_559_v282_)[index105_] = d_561_v284_
                (d_234_v68_).f19 = default__.safeDivisionInt(d_231_v65_.f19, d_231_v65_.f19)
                (d_155_globalState_).f3 = d_153_v0_
                d_562_v285_: _dafny.Array
                def lambda26_(d_563_v68_, d_564_v0_, d_565_v167_):
                    def lambda27_(d_566_i39_):
                        return (d_563_v68_.f19) > ((0) - (len(_dafny.Map({d_564_v0_: (d_565_v167_)[default__.safeIndex(704, (d_565_v167_).length(0))]}))))

                    return lambda27_

                init14_ = lambda26_(d_234_v68_, d_153_v0_, d_376_v167_)
                nw86_ = _dafny.Array(None, 13)
                for i0_14_ in range(nw86_.length(0)):
                    nw86_[i0_14_] = init14_(i0_14_)
                d_562_v285_ = nw86_
                nw87_ = _dafny.Array(None, 4)
                nw87_[int(0)] = (-48) == (d_234_v68_.f19)
                nw87_[int(1)] = d_153_v0_
                nw87_[int(2)] = (d_376_v167_)[default__.safeIndex(704, (d_376_v167_).length(0))]
                nw87_[int(3)] = (d_376_v167_)[default__.safeIndex(704, (d_376_v167_).length(0))]
                d_562_v285_ = nw87_
                d_567_v286_: _dafny.Array
                nw88_ = _dafny.Array(_dafny.Map({}), 4)
                d_567_v286_ = nw88_
                d_568_v287_: D5
                d_568_v287_ = D5_DC17(_dafny.CodePoint('v'))
                d_569_v288_: _dafny.Map
                d_569_v288_ = _dafny.Map({d_231_v65_.f19: (d_568_v287_).cf23})
                index106_ = default__.safeIndex(549, (d_567_v286_).length(0))
                (d_567_v286_)[index106_] = d_569_v288_
                index107_ = default__.safeIndex(549, (d_567_v286_).length(0))
                rhs105_ = (d_569_v288_) | (_dafny.Map({893: d_472_v227_}))
                rhs106_ = d_234_v68_
                lhs92_ = d_567_v286_
                lhs93_ = default__.safeIndex(549, (d_567_v286_).length(0))
                lhs92_[lhs93_] = rhs105_
                d_234_v68_ = rhs106_
            elif True:
                d_156_v2_ = d_156_v2_
                (d_155_globalState_).f3 = d_153_v0_
                d_570_v289_: _dafny.MultiSet
                d_570_v289_ = _dafny.MultiSet([d_234_v68_])
                d_571_v290_: D1
                d_571_v290_ = D1_DC4(d_234_v68_)
                d_572_v291_: _dafny.Seq
                d_572_v291_ = _dafny.SeqWithoutIsStrInference([(d_571_v290_).cf9])
                d_573_v292_: _dafny.Array
                nw89_ = _dafny.Array(None, 14)
                nw89_[int(0)] = d_570_v289_
                nw89_[int(1)] = (d_570_v289_) - (_dafny.MultiSet([d_234_v68_, d_231_v65_, d_231_v65_, d_234_v68_]))
                nw89_[int(2)] = (d_570_v289_) | (_dafny.MultiSet(d_572_v291_))
                nw89_[int(3)] = d_570_v289_
                nw89_[int(4)] = d_570_v289_
                nw89_[int(5)] = d_570_v289_
                nw89_[int(6)] = (_dafny.MultiSet([d_231_v65_])).intersection(d_570_v289_)
                nw89_[int(7)] = (d_570_v289_) - (d_570_v289_)
                nw89_[int(8)] = (d_570_v289_).intersection((_dafny.MultiSet([d_234_v68_])).set(d_234_v68_, default__.abs(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihibjhg"))))))
                nw89_[int(9)] = d_570_v289_
                nw89_[int(10)] = (d_570_v289_) - (d_570_v289_)
                nw89_[int(11)] = (d_570_v289_) | (d_570_v289_)
                nw89_[int(12)] = d_570_v289_
                nw89_[int(13)] = d_570_v289_
                d_573_v292_ = nw89_
                index108_ = default__.safeIndex(645, (d_573_v292_).length(0))
                (d_573_v292_)[index108_] = d_570_v289_
                index109_ = default__.safeIndex(645, (d_573_v292_).length(0))
                (d_573_v292_)[index109_] = (d_570_v289_).set(d_231_v65_, default__.abs(d_557_cf2_))
                d_574_v293_: _dafny.Array
                def lambda28_(d_575_v281_, d_576_v167_):
                    def lambda29_(d_577_i40_):
                        return _dafny.Map({d_575_v281_: (d_576_v167_)[default__.safeIndex(704, (d_576_v167_).length(0))]})

                    return lambda29_

                init15_ = lambda28_(d_552_v281_, d_376_v167_)
                nw90_ = _dafny.Array(None, 22)
                for i0_15_ in range(nw90_.length(0)):
                    nw90_[i0_15_] = init15_(i0_15_)
                d_574_v293_ = nw90_
                d_578_v294_: _dafny.Map
                d_578_v294_ = _dafny.Map({d_552_v281_: d_153_v0_})
                index110_ = default__.safeIndex(133, (d_574_v293_).length(0))
                (d_574_v293_)[index110_] = d_578_v294_
                index111_ = default__.safeIndex(133, (d_574_v293_).length(0))
                (d_574_v293_)[index111_] = (_dafny.Map({d_552_v281_: d_153_v0_})) | (d_578_v294_)
                d_579_v295_: _dafny.MultiSet
                d_579_v295_ = _dafny.MultiSet([len(d_154_v1_), d_234_v68_.f19])
                d_580_v296_: _dafny.Map
                d_580_v296_ = _dafny.Map({D5_DC20(d_377_v168_): (d_579_v295_) - (_dafny.MultiSet([d_558_cf1_]))})
                d_581_v297_: _dafny.Seq
                d_581_v297_ = _dafny.SeqWithoutIsStrInference([d_377_v168_, d_378_v169_])
                d_580_v296_ = (d_580_v296_).set(D5_DC20((d_581_v297_)[default__.safeIndex(-900, len(d_581_v297_))]), d_579_v295_)
            d_582_v298_: _dafny.MultiSet
            d_582_v298_ = _dafny.MultiSet([d_231_v65_.f19])
            (d_155_globalState_).f8 = (0) - ((d_234_v68_).fm4(default__.safeModuloInt((d_582_v298_).cardinality, d_231_v65_.f19), -138, d_153_v0_, d_556_cf3_, d_155_globalState_))
            (d_155_globalState_).f3 = (d_376_v167_)[default__.safeIndex(704, (d_376_v167_).length(0))]
        elif source11_.is_DC2:
            d_583___mcc_h31_ = source11_.cf4
            d_584___mcc_h32_ = source11_.cf5
            d_585___mcc_h33_ = source11_.cf6
            d_586___mcc_h34_ = source11_.cf7
            d_587_cf7_ = d_586___mcc_h34_
            d_588_cf6_ = d_585___mcc_h33_
            d_589_cf5_ = d_584___mcc_h32_
            d_590_cf4_ = d_583___mcc_h31_
            d_591_v299_: _dafny.MultiSet
            d_591_v299_ = _dafny.MultiSet([d_233_v67_])
            d_592_v300_: D0
            d_592_v300_ = D0_DC0(d_231_v65_.f19)
            d_593_v301_: _dafny.Array
            nw91_ = _dafny.Array(None, 17)
            nw91_[int(0)] = d_234_v68_
            nw91_[int(1)] = d_234_v68_
            nw91_[int(2)] = d_234_v68_
            nw91_[int(3)] = d_231_v65_
            nw91_[int(4)] = d_234_v68_
            nw91_[int(5)] = d_234_v68_
            nw91_[int(6)] = d_231_v65_
            nw91_[int(7)] = d_231_v65_
            nw91_[int(8)] = d_234_v68_
            nw91_[int(9)] = d_234_v68_
            nw91_[int(10)] = d_231_v65_
            nw91_[int(11)] = d_234_v68_
            nw91_[int(12)] = d_231_v65_
            nw91_[int(13)] = d_231_v65_
            nw91_[int(14)] = d_234_v68_
            nw91_[int(15)] = d_231_v65_
            nw91_[int(16)] = d_234_v68_
            d_593_v301_ = nw91_
            d_594_v302_: _dafny.Map
            d_594_v302_ = _dafny.Map({d_593_v301_: d_231_v65_.f19})
            d_595_v303_: _dafny.Map
            d_595_v303_ = _dafny.Map({d_590_cf4_: len(d_233_v67_)})
            rhs107_ = True
            rhs108_ = (d_234_v68_.f19) != (((d_591_v299_).set(d_233_v67_, default__.abs(d_234_v68_.f19))).cardinality)
            rhs109_ = (d_231_v65_.f19) < ((d_592_v300_).cf0)
            rhs110_ = default__.safeDivisionInt((d_231_v65_.f19) - (len(d_594_v302_)), len(d_595_v303_))
            d_589_cf5_ = rhs107_
            d_589_cf5_ = rhs108_
            d_153_v0_ = rhs109_
            d_588_cf6_ = rhs110_
            if False:
                d_596_v304_: _dafny.Seq
                d_596_v304_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))]])])
                (d_234_v68_).f19 = default__.safeDivisionInt(default__.fm0(d_156_v2_, (d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))], d_596_v304_, d_231_v65_.f19, d_155_globalState_), d_156_v2_)
                d_597_v305_: _dafny.Array
                nw92_ = _dafny.Array(_dafny.Array(None, 0), 20)
                d_597_v305_ = nw92_
                index112_ = default__.safeIndex(424, (d_597_v305_).length(0))
                (d_597_v305_)[index112_] = d_376_v167_
                index113_ = default__.safeIndex(424, (d_597_v305_).length(0))
                (d_597_v305_)[index113_] = d_376_v167_
                d_598_v306_: _dafny.Array
                nw93_ = _dafny.Array(_dafny.MultiSet({}), 9)
                d_598_v306_ = nw93_
                d_599_v307_: _dafny.MultiSet
                d_599_v307_ = _dafny.MultiSet([d_234_v68_, d_231_v65_, d_231_v65_, d_231_v65_, d_234_v68_])
                index114_ = default__.safeIndex(399, (d_598_v306_).length(0))
                (d_598_v306_)[index114_] = (d_599_v307_) - (_dafny.MultiSet([d_234_v68_, d_234_v68_]))
                index115_ = default__.safeIndex(399, (d_598_v306_).length(0))
                (d_598_v306_)[index115_] = (d_599_v307_) | (d_599_v307_)
                d_600_v308_: int
                d_601_v309_: _dafny.Seq
                d_602_v310_: int
                d_603_v311_: _dafny.MultiSet
                out40_: int
                out41_: _dafny.Seq
                out42_: int
                out43_: _dafny.MultiSet
                out40_, out41_, out42_, out43_ = default__.m0(d_589_cf5_, not((d_587_cf7_ if d_590_cf4_ else d_587_cf7_)), 602, (d_157_v3_).issubset(d_157_v3_), d_155_globalState_)
                d_600_v308_ = out40_
                d_601_v309_ = out41_
                d_602_v310_ = out42_
                d_603_v311_ = out43_
                d_175_v18_ = (((d_601_v309_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhymh")))).set(default__.safeIndex(d_234_v68_.f19, len((d_601_v309_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhymh"))))), d_472_v227_)) + (default__.fm5(d_156_v2_, d_155_globalState_))
            elif True:
                d_587_cf7_ = not(d_587_cf7_)
                d_604_v312_: _dafny.Map
                d_604_v312_ = _dafny.Map({d_175_v18_: d_589_cf5_})
                d_604_v312_ = (d_604_v312_).set(d_175_v18_, d_590_cf4_)
                d_605_v313_: _dafny.Map
                d_605_v313_ = _dafny.Map({d_154_v1_: d_589_cf5_})
                (d_155_globalState_).f0 = default__.safeDivisionInt(default__.safeDivisionInt(len((_dafny.Map({d_153_v0_: d_231_v65_.f19})).set(((d_605_v313_)[d_154_v1_] if (d_154_v1_) in (d_605_v313_) else d_590_cf4_), d_231_v65_.f19)), d_588_cf6_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hryki"))))
                index116_ = default__.safeIndex(263, (d_593_v301_).length(0))
                (d_593_v301_)[index116_] = d_231_v65_
                index117_ = default__.safeIndex(263, (d_593_v301_).length(0))
                (d_593_v301_)[index117_] = d_234_v68_
                d_606_v314_: int
                d_607_v315_: _dafny.Seq
                d_608_v316_: int
                d_609_v317_: _dafny.MultiSet
                out44_: int
                out45_: _dafny.Seq
                out46_: int
                out47_: _dafny.MultiSet
                out44_, out45_, out46_, out47_ = default__.m0((d_234_v68_.f19) <= (d_231_v65_.f19), (d_154_v1_)[default__.safeIndex(d_588_cf6_, len(d_154_v1_))], d_231_v65_.f19, False, d_155_globalState_)
                d_606_v314_ = out44_
                d_607_v315_ = out45_
                d_608_v316_ = out46_
                d_609_v317_ = out47_
            if d_589_cf5_:
                d_610_v318_: _dafny.Array
                nw94_ = _dafny.Array(_dafny.Map({}), 2)
                d_610_v318_ = nw94_
                d_611_v319_: _dafny.Map
                d_611_v319_ = _dafny.Map({d_590_cf4_: d_587_cf7_})
                index118_ = default__.safeIndex(805, (d_610_v318_).length(0))
                (d_610_v318_)[index118_] = d_611_v319_
                d_612_v320_: _dafny.Seq
                d_612_v320_ = _dafny.SeqWithoutIsStrInference([d_376_v167_, d_376_v167_])
                index119_ = default__.safeIndex(805, (d_610_v318_).length(0))
                rhs111_ = (d_611_v319_) | ((d_611_v319_) | (d_611_v319_))
                rhs112_ = (default__.safeModuloInt(d_234_v68_.f19, len((d_154_v1_).set(default__.safeIndex(d_231_v65_.f19, len(d_154_v1_)), d_153_v0_)))) + ((d_231_v65_.f19 if d_587_cf7_ else d_234_v68_.f19))
                rhs113_ = (len(d_175_v18_) if default__.fm2((d_552_v281_).cf3, d_155_globalState_) else (0) - (d_234_v68_.f19))
                rhs114_ = (_dafny.SeqWithoutIsStrInference([d_376_v167_])) == ((d_612_v320_ if d_587_cf7_ else d_612_v320_))
                lhs94_ = d_610_v318_
                lhs95_ = default__.safeIndex(805, (d_610_v318_).length(0))
                lhs96_ = d_155_globalState_
                lhs97_ = d_155_globalState_
                lhs94_[lhs95_] = rhs111_
                lhs96_.f2 = rhs112_
                lhs97_.f0 = rhs113_
                d_153_v0_ = rhs114_
                d_613_v321_: _dafny.Set
                d_613_v321_ = _dafny.Set({d_593_v301_})
                d_613_v321_ = (d_613_v321_) | (d_613_v321_)
                d_614_v322_: int
                d_615_v323_: _dafny.Seq
                d_616_v324_: int
                d_617_v325_: _dafny.MultiSet
                out48_: int
                out49_: _dafny.Seq
                out50_: int
                out51_: _dafny.MultiSet
                out48_, out49_, out50_, out51_ = default__.m0(d_590_cf4_, d_589_cf5_, d_231_v65_.f19, not(d_587_cf7_), d_155_globalState_)
                d_614_v322_ = out48_
                d_615_v323_ = out49_
                d_616_v324_ = out50_
                d_617_v325_ = out51_
                d_618_v326_: D0
                d_618_v326_ = D0_DC2(d_590_cf4_, d_587_cf7_, d_231_v65_.f19, d_589_cf5_)
                d_619_v327_: int
                d_620_v328_: _dafny.Seq
                d_621_v329_: int
                d_622_v330_: _dafny.MultiSet
                out52_: int
                out53_: _dafny.Seq
                out54_: int
                out55_: _dafny.MultiSet
                out52_, out53_, out54_, out55_ = default__.m0((d_618_v326_).cf5, default__.fm2(True, d_155_globalState_), d_588_cf6_, d_589_cf5_, d_155_globalState_)
                d_619_v327_ = out52_
                d_620_v328_ = out53_
                d_621_v329_ = out54_
                d_622_v330_ = out55_
                d_621_v329_ = d_616_v324_
            elif True:
                d_623_v331_: _dafny.Array
                nw95_ = _dafny.Array(_dafny.Map({}), 16)
                d_623_v331_ = nw95_
                d_624_v332_: _dafny.Map
                d_624_v332_ = _dafny.Map({d_175_v18_: d_590_cf4_})
                index120_ = default__.safeIndex(103, (d_623_v331_).length(0))
                (d_623_v331_)[index120_] = d_624_v332_
                index121_ = default__.safeIndex(103, (d_623_v331_).length(0))
                (d_623_v331_)[index121_] = d_624_v332_
                d_472_v227_ = d_472_v227_
                d_625_v333_: D2
                d_625_v333_ = D2_DC9(d_234_v68_.f19, d_234_v68_, d_231_v65_)
                d_626_v334_: D1
                d_626_v334_ = D1_DC3(default__.fm5(d_231_v65_.f19, d_155_globalState_))
                d_627_v335_: _dafny.Set
                d_627_v335_ = _dafny.Set({d_626_v334_})
                d_628_v336_: _dafny.Map
                def iife49_(_pat_let18_0):
                    def iife50_(d_629_dt__update__tmp_h6_):
                        def iife51_(_pat_let19_0):
                            def iife52_(d_630_dt__update_hcf13_h0_):
                                return D2_DC9(d_630_dt__update_hcf13_h0_, (d_629_dt__update__tmp_h6_).cf14, (d_629_dt__update__tmp_h6_).cf15)
                            return iife52_(_pat_let19_0)
                        return iife51_(961)
                    return iife50_(_pat_let18_0)
                d_628_v336_ = _dafny.Map({iife49_(d_625_v333_): (d_627_v335_) - (_dafny.Set({d_626_v334_, D1_DC3(d_175_v18_), d_626_v334_}))})
                d_628_v336_ = (d_628_v336_).set(D2_DC9(d_588_cf6_, d_231_v65_, d_231_v65_), (d_627_v335_) - (d_627_v335_))
                d_595_v303_ = (d_595_v303_).set(d_590_cf4_, d_234_v68_.f19)
                d_631_v337_: _dafny.Array
                nw96_ = _dafny.Array(int(0), 23)
                d_631_v337_ = nw96_
                d_632_v338_: _dafny.Seq
                d_632_v338_ = _dafny.SeqWithoutIsStrInference([d_175_v18_, d_175_v18_])
                index122_ = default__.safeIndex(300, (d_631_v337_).length(0))
                (d_631_v337_)[index122_] = len(d_632_v338_)
                index123_ = default__.safeIndex(300, (d_631_v337_).length(0))
                (d_631_v337_)[index123_] = d_231_v65_.f19
            d_633_v339_: _dafny.Seq
            d_633_v339_ = _dafny.SeqWithoutIsStrInference([d_234_v68_, d_231_v65_, d_234_v68_, d_234_v68_, d_231_v65_])
            d_634_v340_: D4
            d_634_v340_ = D4_DC14(d_633_v339_)
            d_635_v341_: _dafny.Set
            d_635_v341_ = _dafny.Set({d_230_v64_, d_230_v64_})
            def iife53_():
                coll13_ = _dafny.Set()
                compr_13_: _dafny.Seq
                for compr_13_ in (_dafny.SeqWithoutIsStrInference([d_230_v64_])).Elements:
                    d_636_v342_: _dafny.Seq = compr_13_
                    if (d_636_v342_) in (_dafny.SeqWithoutIsStrInference([d_230_v64_])):
                        coll13_ = coll13_.union(_dafny.Set([d_636_v342_]))
                return _dafny.Set(coll13_)
            rhs115_ = (d_635_v341_) != (iife53_()
            )
            rhs116_ = d_634_v340_
            rhs117_ = _dafny.SeqWithoutIsStrInference([d_472_v227_ for d_637_i41_ in range(default__.abs(633))])
            rhs118_ = (d_234_v68_.f19) >= (d_588_cf6_)
            lhs98_ = d_155_globalState_
            lhs99_ = d_155_globalState_
            lhs98_.f3 = rhs115_
            d_634_v340_ = rhs116_
            d_175_v18_ = rhs117_
            lhs99_.f3 = rhs118_
        elif True:
            d_638___mcc_h35_ = source11_.cf0
            d_639_cf0_ = d_638___mcc_h35_
            d_640_v343_: C0
            nw97_ = C0()
            nw97_.ctor__(d_639_cf0_)
            d_640_v343_ = nw97_
            d_641_v344_: _dafny.Array
            nw98_ = _dafny.Array(None, 4)
            d_641_v344_ = nw98_
            index124_ = default__.safeIndex(883, (d_641_v344_).length(0))
            (d_641_v344_)[index124_] = d_234_v68_
            index125_ = default__.safeIndex(883, (d_641_v344_).length(0))
            (d_641_v344_)[index125_] = d_231_v65_
            d_642_v345_: _dafny.MultiSet
            d_642_v345_ = _dafny.MultiSet([-949])
            if ((d_640_v343_.f19) == (d_234_v68_.f19)) == ((d_642_v345_).ispropersubset(d_642_v345_)):
                d_643_v346_: _dafny.Set
                d_643_v346_ = _dafny.Set({d_640_v343_.f19, d_640_v343_.f19, d_640_v343_.f19, d_231_v65_.f19})
                d_644_v349_: _dafny.MultiSet
                def iife54_():
                    coll14_ = _dafny.Set()
                    compr_14_: int
                    for compr_14_ in _dafny.IntegerRange(964, -496):
                        d_645_v347_: int = compr_14_
                        if ((964) <= (d_645_v347_)) and ((d_645_v347_) < (-496)):
                            coll14_ = coll14_.union(_dafny.Set([(d_645_v347_) * (d_640_v343_.f19)]))
                    return _dafny.Set(coll14_)
                def iife55_():
                    coll15_ = _dafny.Set()
                    compr_15_: int
                    for compr_15_ in _dafny.IntegerRange(-993, -987):
                        d_646_v348_: int = compr_15_
                        if ((-993) <= (d_646_v348_)) and ((d_646_v348_) < (-987)):
                            coll15_ = coll15_.union(_dafny.Set([(d_646_v348_) * (len(d_175_v18_))]))
                    return _dafny.Set(coll15_)
                d_644_v349_ = _dafny.MultiSet([d_643_v346_, (d_643_v346_) | (d_643_v346_), iife54_()
                , default__.fm11(d_234_v68_.f19, d_155_globalState_), iife55_()
                ])
                d_647_v350_: _dafny.Seq
                d_647_v350_ = _dafny.SeqWithoutIsStrInference([d_644_v349_])
                d_644_v349_ = (d_644_v349_) - (((d_647_v350_)[default__.safeIndex(d_234_v68_.f19, len(d_647_v350_))]) - (d_644_v349_))
                d_648_v351_: _dafny.Array
                nw99_ = _dafny.Array(None, 11)
                nw99_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnbj"))
                nw99_[int(1)] = d_175_v18_
                nw99_[int(2)] = d_175_v18_
                nw99_[int(3)] = d_175_v18_
                nw99_[int(4)] = d_175_v18_
                nw99_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omisqr"))
                nw99_[int(6)] = d_175_v18_
                nw99_[int(7)] = d_175_v18_
                nw99_[int(8)] = d_175_v18_
                nw99_[int(9)] = d_175_v18_
                nw99_[int(10)] = d_175_v18_
                d_648_v351_ = nw99_
                d_649_v352_: D3
                d_649_v352_ = D3_DC12(d_648_v351_, d_153_v0_)
                d_650_v353_: D1
                d_650_v353_ = D1_DC5(False)
                d_651_v354_: int
                d_652_v355_: _dafny.Seq
                d_653_v356_: int
                d_654_v357_: _dafny.MultiSet
                out56_: int
                out57_: _dafny.Seq
                out58_: int
                out59_: _dafny.MultiSet
                out56_, out57_, out58_, out59_ = default__.m0((d_154_v1_)[default__.safeIndex(187, len(d_154_v1_))], d_153_v0_, (d_231_v65_).fm4(d_640_v343_.f19, d_231_v65_.f19, (d_649_v352_).cf19, (d_650_v353_).cf10, d_155_globalState_), (d_640_v343_.f19) < ((d_234_v68_).fm4(d_156_v2_, (0) - ((0) - (d_234_v68_.f19)), d_153_v0_, d_153_v0_, d_155_globalState_)), d_155_globalState_)
                d_651_v354_ = out56_
                d_652_v355_ = out57_
                d_653_v356_ = out58_
                d_654_v357_ = out59_
                d_655_v358_: C0
                nw100_ = C0()
                nw100_.ctor__(d_234_v68_.f19)
                d_655_v358_ = nw100_
                d_656_v359_: C0
                nw101_ = C0()
                nw101_.ctor__(d_231_v65_.f19)
                d_656_v359_ = nw101_
                d_657_v360_: _dafny.Array
                nw102_ = _dafny.Array(int(0), 2)
                d_657_v360_ = nw102_
                d_657_v360_ = d_657_v360_
            elif True:
                (d_155_globalState_).f3 = d_153_v0_
                d_640_v343_ = d_231_v65_
                index126_ = default__.safeIndex(169, (d_376_v167_).length(0))
                (d_376_v167_)[index126_] = d_153_v0_
                index127_ = default__.safeIndex(169, (d_376_v167_).length(0))
                (d_376_v167_)[index127_] = (d_154_v1_)[default__.safeIndex((d_234_v68_.f19 if d_153_v0_ else d_234_v68_.f19), len(d_154_v1_))]
                d_642_v345_ = d_642_v345_
                d_658_v361_: _dafny.Seq
                d_658_v361_ = _dafny.SeqWithoutIsStrInference([d_175_v18_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")), d_175_v18_])
                d_659_v362_: D1
                d_659_v362_ = D1_DC3((d_658_v361_)[default__.safeIndex((0) - (d_231_v65_.f19), len(d_658_v361_))])
                rhs119_ = d_659_v362_
                rhs120_ = default__.fm5(d_156_v2_, d_155_globalState_)
                d_659_v362_ = rhs119_
                d_175_v18_ = rhs120_
            (d_155_globalState_).f3 = not(not (not(d_153_v0_)) or (d_153_v0_))
        d_660_v363_: _dafny.Array
        nw103_ = _dafny.Array(None, 9)
        nw103_[int(0)] = d_552_v281_
        nw103_[int(1)] = default__.fm12(d_231_v65_.f19, d_234_v68_.f19, d_234_v68_.f19, d_153_v0_, d_155_globalState_)
        nw103_[int(2)] = d_552_v281_
        nw103_[int(3)] = D0_DC1(d_234_v68_.f19, d_156_v2_, d_153_v0_)
        nw103_[int(4)] = D0_DC1((d_231_v65_).fm3(d_155_globalState_), d_231_v65_.f19, d_153_v0_)
        nw103_[int(5)] = d_552_v281_
        nw103_[int(6)] = d_552_v281_
        nw103_[int(7)] = d_552_v281_
        nw103_[int(8)] = d_552_v281_
        d_660_v363_ = nw103_
        index128_ = default__.safeIndex(864, (d_660_v363_).length(0))
        (d_660_v363_)[index128_] = d_552_v281_
        index129_ = default__.safeIndex(864, (d_660_v363_).length(0))
        (d_660_v363_)[index129_] = d_552_v281_
        d_661_v365_: _dafny.Map
        def iife56_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in _dafny.IntegerRange(-50, 413):
                d_662_v364_: int = compr_16_
                if ((-50) <= (d_662_v364_)) and ((d_662_v364_) < (413)):
                    coll16_[default__.safeDivisionInt(d_662_v364_, len(d_154_v1_))] = d_234_v68_.f19
            return _dafny.Map(coll16_)
        d_661_v365_ = _dafny.Map({d_376_v167_: iife56_()
        })
        d_153_v0_ = (d_376_v167_) in (d_661_v365_)
        d_663_v366_: _dafny.Map
        d_663_v366_ = _dafny.Map({len(d_175_v18_): d_157_v3_})
        d_664_i42_: int
        d_664_i42_ = 0
        with _dafny.label("1"):
            while (d_234_v68_.f19) in (d_663_v366_):
                with _dafny.c_label("1"):
                    if (d_664_i42_) >= (100):
                        raise _dafny.Break("1")
                    d_664_i42_ = (d_664_i42_) + (1)
                    d_665_v367_: _dafny.Map
                    d_665_v367_ = _dafny.Map({d_234_v68_.f19: d_153_v0_})
                    d_666_v368_: D3
                    d_666_v368_ = D3_DC11(d_665_v367_)
                    d_667_v369_: D3
                    d_667_v369_ = D3_DC13(d_666_v368_)
                    d_668_v370_: D3
                    d_668_v370_ = D3_DC13((d_667_v369_).cf20)
                    source12_ = d_668_v370_
                    if source12_.is_DC12:
                        d_669___mcc_h36_ = source12_.cf18
                        d_670___mcc_h37_ = source12_.cf19
                        d_671_cf19_ = d_670___mcc_h37_
                        d_672_cf18_ = d_669___mcc_h36_
                        index130_ = default__.safeIndex(709, (d_376_v167_).length(0))
                        (d_376_v167_)[index130_] = (d_671_cf19_) or (d_671_cf19_)
                        d_673_v371_: _dafny.Set
                        d_673_v371_ = _dafny.Set({(0) - (d_234_v68_.f19)})
                        index131_ = default__.safeIndex(709, (d_376_v167_).length(0))
                        (d_376_v167_)[index131_] = (True if (d_234_v68_.f19) >= (len(d_673_v371_)) else d_153_v0_)
                        (d_155_globalState_).f8 = len(d_175_v18_)
                        d_674_v372_: _dafny.Map
                        d_674_v372_ = _dafny.Map({True: d_379_v170_})
                        (d_155_globalState_).f2 = len((d_674_v372_).set((d_234_v68_.f19) != (d_156_v2_), d_379_v170_))
                        index132_ = default__.safeIndex(709, (d_376_v167_).length(0))
                        (d_376_v167_)[index132_] = (d_234_v68_.f19) < (default__.safeModuloInt(len(d_175_v18_), d_231_v65_.f19))
                    elif source12_.is_DC11:
                        d_675___mcc_h38_ = source12_.cf17
                        d_676_cf17_ = d_675___mcc_h38_
                        d_156_v2_ = 554
                        d_156_v2_ = d_231_v65_.f19
                        d_677_v373_: D5
                        d_677_v373_ = D5_DC19(d_153_v0_)
                        pat_let_tv24_ = d_153_v0_
                        d_678_v374_: _dafny.Seq
                        def iife57_(_pat_let20_0):
                            def iife58_(d_679_dt__update__tmp_h7_):
                                def iife59_(_pat_let21_0):
                                    def iife60_(d_680_dt__update_hcf28_h0_):
                                        return D5_DC19(d_680_dt__update_hcf28_h0_)
                                    return iife60_(_pat_let21_0)
                                return iife59_(pat_let_tv24_)
                            return iife58_(_pat_let20_0)
                        d_678_v374_ = _dafny.SeqWithoutIsStrInference([d_677_v373_, iife57_(d_677_v373_)])
                        d_681_v375_: _dafny.Array
                        nw104_ = _dafny.Array(int(0), 13)
                        d_681_v375_ = nw104_
                        index133_ = default__.safeIndex(140, (d_681_v375_).length(0))
                        (d_681_v375_)[index133_] = d_231_v65_.f19
                        index134_ = default__.safeIndex(140, (d_681_v375_).length(0))
                        rhs121_ = _dafny.SeqWithoutIsStrInference([d_677_v373_, default__.fm13((0) - (d_234_v68_.f19), d_155_globalState_), d_677_v373_])
                        rhs122_ = (d_234_v68_.f19) * (d_156_v2_)
                        rhs123_ = ((d_230_v64_) + (d_230_v64_)) + ((d_230_v64_) + (d_230_v64_))
                        rhs124_ = d_153_v0_
                        lhs100_ = d_681_v375_
                        lhs101_ = default__.safeIndex(140, (d_681_v375_).length(0))
                        d_678_v374_ = rhs121_
                        lhs100_[lhs101_] = rhs122_
                        d_230_v64_ = rhs123_
                        d_153_v0_ = rhs124_
                        d_682_v376_: _dafny.MultiSet
                        d_682_v376_ = _dafny.MultiSet([d_234_v68_.f19, d_231_v65_.f19, d_234_v68_.f19, d_234_v68_.f19])
                        (d_155_globalState_).f3 = ((d_234_v68_.f19) * ((d_231_v65_).fm4(d_231_v65_.f19, d_231_v65_.f19, d_153_v0_, d_153_v0_, d_155_globalState_))) in (d_682_v376_)
                    elif True:
                        d_683___mcc_h39_ = source12_.cf20
                        d_684_cf20_ = d_683___mcc_h39_
                        d_685_v377_: int
                        d_686_v378_: _dafny.Seq
                        d_687_v379_: int
                        d_688_v380_: _dafny.MultiSet
                        out60_: int
                        out61_: _dafny.Seq
                        out62_: int
                        out63_: _dafny.MultiSet
                        out60_, out61_, out62_, out63_ = default__.m0((d_153_v0_) == (d_153_v0_), d_153_v0_, d_234_v68_.f19, d_153_v0_, d_155_globalState_)
                        d_685_v377_ = out60_
                        d_686_v378_ = out61_
                        d_687_v379_ = out62_
                        d_688_v380_ = out63_
                        d_153_v0_ = (d_153_v0_) == (d_153_v0_)
                        d_689_v381_: _dafny.Set
                        d_689_v381_ = _dafny.Set({d_687_v379_, -70})
                        d_690_v382_: _dafny.Array
                        nw105_ = _dafny.Array(None, 26)
                        nw105_[int(0)] = d_230_v64_
                        nw105_[int(1)] = (d_230_v64_ if d_153_v0_ else d_230_v64_)
                        nw105_[int(2)] = d_230_v64_
                        nw105_[int(3)] = d_230_v64_
                        nw105_[int(4)] = d_230_v64_
                        nw105_[int(5)] = (d_230_v64_) + (d_230_v64_)
                        nw105_[int(6)] = (d_230_v64_) + (d_230_v64_)
                        nw105_[int(7)] = d_230_v64_
                        nw105_[int(8)] = d_230_v64_
                        nw105_[int(9)] = d_230_v64_
                        nw105_[int(10)] = d_230_v64_
                        nw105_[int(11)] = d_230_v64_
                        nw105_[int(12)] = d_230_v64_
                        nw105_[int(13)] = (_dafny.SeqWithoutIsStrInference([d_234_v68_.f19 for d_691_i43_ in range(default__.abs(177))])).set(default__.safeIndex((0) - ((d_234_v68_).fm4(d_234_v68_.f19, 248, d_153_v0_, True, d_155_globalState_)), len(_dafny.SeqWithoutIsStrInference([d_234_v68_.f19 for d_692_i43_ in range(default__.abs(177))]))), d_156_v2_)
                        nw105_[int(14)] = d_230_v64_
                        nw105_[int(15)] = (d_230_v64_) + (d_230_v64_)
                        nw105_[int(16)] = d_230_v64_
                        nw105_[int(17)] = (_dafny.SeqWithoutIsStrInference([(0) - (d_234_v68_.f19) for d_693_i44_ in range(default__.abs(809))])).set(default__.safeIndex(d_156_v2_, len(_dafny.SeqWithoutIsStrInference([(0) - (d_234_v68_.f19) for d_694_i44_ in range(default__.abs(809))]))), d_231_v65_.f19)
                        nw105_[int(18)] = d_230_v64_
                        nw105_[int(19)] = _dafny.SeqWithoutIsStrInference([d_234_v68_.f19 for d_695_i45_ in range(default__.abs(302))])
                        nw105_[int(20)] = (_dafny.SeqWithoutIsStrInference([d_234_v68_.f19, len(d_689_v381_)])).set(default__.safeIndex(136, len(_dafny.SeqWithoutIsStrInference([d_234_v68_.f19, len(d_689_v381_)]))), d_687_v379_)
                        nw105_[int(21)] = (d_230_v64_) + (_dafny.SeqWithoutIsStrInference([d_685_v377_]))
                        nw105_[int(22)] = d_230_v64_
                        nw105_[int(23)] = d_230_v64_
                        nw105_[int(24)] = default__.fm14(d_153_v0_, d_155_globalState_)
                        nw105_[int(25)] = (_dafny.SeqWithoutIsStrInference([750, d_234_v68_.f19])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_153_v0_])), d_231_v65_.f19]))
                        d_690_v382_ = nw105_
                        index135_ = default__.safeIndex(300, (d_690_v382_).length(0))
                        (d_690_v382_)[index135_] = (d_230_v64_) + (d_230_v64_)
                        d_696_v383_: _dafny.Array
                        def lambda30_(d_697_v2_):
                            def lambda31_(d_698_i46_):
                                return (d_698_i46_) * (d_697_v2_)

                            return lambda31_

                        init16_ = lambda30_(d_156_v2_)
                        nw106_ = _dafny.Array(None, 26)
                        for i0_16_ in range(nw106_.length(0)):
                            nw106_[i0_16_] = init16_(i0_16_)
                        d_696_v383_ = nw106_
                        index136_ = default__.safeIndex(279, (d_696_v383_).length(0))
                        (d_696_v383_)[index136_] = d_231_v65_.f19
                        d_699_v384_: _dafny.Map
                        d_699_v384_ = _dafny.Map({d_153_v0_: _dafny.SeqWithoutIsStrInference([d_231_v65_.f19])})
                        index137_ = default__.safeIndex(300, (d_690_v382_).length(0))
                        index138_ = default__.safeIndex(279, (d_696_v383_).length(0))
                        rhs125_ = (d_153_v0_) in (d_154_v1_)
                        rhs126_ = ((d_230_v64_) + (_dafny.SeqWithoutIsStrInference([d_156_v2_]))) + (d_230_v64_)
                        rhs127_ = d_316_v126_
                        rhs128_ = (d_234_v68_).fm4(len(d_699_v384_), d_231_v65_.f19, ((d_231_v65_).fm3(d_155_globalState_)) == (d_234_v68_.f19), d_153_v0_, d_155_globalState_)
                        lhs102_ = d_155_globalState_
                        lhs103_ = d_690_v382_
                        lhs104_ = default__.safeIndex(300, (d_690_v382_).length(0))
                        lhs105_ = d_696_v383_
                        lhs106_ = default__.safeIndex(279, (d_696_v383_).length(0))
                        lhs102_.f3 = rhs125_
                        lhs103_[lhs104_] = rhs126_
                        d_316_v126_ = rhs127_
                        lhs105_[lhs106_] = rhs128_
                        (d_155_globalState_).f3 = default__.fm2(not(False), d_155_globalState_)
                    if d_153_v0_:
                        rhs129_ = (_dafny.SeqWithoutIsStrInference([d_472_v227_ for d_700_i47_ in range(default__.abs(-882))])).set(default__.safeIndex(d_231_v65_.f19, len(_dafny.SeqWithoutIsStrInference([d_472_v227_ for d_701_i47_ in range(default__.abs(-882))]))), ((d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))] if (d_154_v1_)[default__.safeIndex(d_234_v68_.f19, len(d_154_v1_))] else d_472_v227_))
                        rhs130_ = d_472_v227_
                        rhs131_ = (((d_233_v67_)[d_231_v65_.f19] if (d_231_v65_.f19) in (d_233_v67_) else d_231_v65_.f19)) <= (d_234_v68_.f19)
                        rhs132_ = d_234_v68_.f19
                        rhs133_ = d_231_v65_
                        lhs107_ = d_155_globalState_
                        d_175_v18_ = rhs129_
                        d_472_v227_ = rhs130_
                        d_153_v0_ = rhs131_
                        lhs107_.f2 = rhs132_
                        d_234_v68_ = rhs133_
                        nw107_ = C0()
                        nw107_.ctor__(d_231_v65_.f19)
                        d_231_v65_ = nw107_
                        d_702_v385_: _dafny.Map
                        d_702_v385_ = _dafny.Map({d_153_v0_: d_472_v227_})
                        d_702_v385_ = (d_702_v385_).set(d_153_v0_, _dafny.CodePoint('m'))
                        d_703_v386_: _dafny.Array
                        def lambda32_(d_704_v65_):
                            def lambda33_(d_705_i48_):
                                return (d_705_i48_) - (d_704_v65_.f19)

                            return lambda33_

                        init17_ = lambda32_(d_231_v65_)
                        nw108_ = _dafny.Array(None, 10)
                        for i0_17_ in range(nw108_.length(0)):
                            nw108_[i0_17_] = init17_(i0_17_)
                        d_703_v386_ = nw108_
                        index139_ = default__.safeIndex(448, (d_703_v386_).length(0))
                        (d_703_v386_)[index139_] = d_234_v68_.f19
                        index140_ = default__.safeIndex(448, (d_703_v386_).length(0))
                        rhs134_ = d_230_v64_
                        rhs135_ = d_231_v65_
                        rhs136_ = d_234_v68_.f19
                        rhs137_ = 742
                        rhs138_ = ((d_702_v385_)[d_153_v0_] if (d_153_v0_) in (d_702_v385_) else _dafny.CodePoint('e'))
                        lhs108_ = d_155_globalState_
                        lhs109_ = d_703_v386_
                        lhs110_ = default__.safeIndex(448, (d_703_v386_).length(0))
                        d_230_v64_ = rhs134_
                        d_234_v68_ = rhs135_
                        lhs108_.f8 = rhs136_
                        lhs109_[lhs110_] = rhs137_
                        d_472_v227_ = rhs138_
                        (d_231_v65_).f19 = d_231_v65_.f19
                    elif True:
                        index141_ = default__.safeIndex(383, (d_376_v167_).length(0))
                        (d_376_v167_)[index141_] = (d_472_v227_) in (d_175_v18_)
                        d_706_v387_: _dafny.Set
                        d_706_v387_ = _dafny.Set({(d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))], _dafny.CodePoint('p')})
                        index142_ = default__.safeIndex(383, (d_376_v167_).length(0))
                        (d_376_v167_)[index142_] = (d_153_v0_) == ((d_706_v387_).issubset(d_706_v387_))
                        d_707_v388_: C0
                        nw109_ = C0()
                        nw109_.ctor__(default__.safeDivisionInt(len(d_175_v18_), d_231_v65_.f19))
                        d_707_v388_ = nw109_
                        d_708_v389_: _dafny.Set
                        d_708_v389_ = _dafny.Set({512, -517, d_234_v68_.f19})
                        def iife61_():
                            coll17_ = _dafny.Set()
                            compr_17_: int
                            for compr_17_ in _dafny.IntegerRange(779, 398):
                                d_709_v390_: int = compr_17_
                                if ((779) <= (d_709_v390_)) and ((d_709_v390_) < (398)):
                                    coll17_ = coll17_.union(_dafny.Set([default__.safeModuloInt(d_709_v390_, d_231_v65_.f19)]))
                            return _dafny.Set(coll17_)
                        rhs139_ = (default__.fm11(d_231_v65_.f19, d_155_globalState_) if not(d_153_v0_) else iife61_()
                        )
                        rhs140_ = (0) - ((d_231_v65_.f19) - (168))
                        rhs141_ = ((0) - (d_234_v68_.f19)) - ((d_231_v65_).fm3(d_155_globalState_))
                        rhs142_ = d_153_v0_
                        lhs111_ = d_234_v68_
                        lhs112_ = d_155_globalState_
                        lhs113_ = d_155_globalState_
                        d_708_v389_ = rhs139_
                        lhs111_.f19 = rhs140_
                        lhs112_.f12 = rhs141_
                        lhs113_.f3 = rhs142_
                        d_710_v391_: C0
                        nw110_ = C0()
                        nw110_.ctor__((264 if (d_376_v167_)[default__.safeIndex(383, (d_376_v167_).length(0))] else d_234_v68_.f19))
                        d_710_v391_ = nw110_
                        (d_707_v388_).f19 = d_156_v2_
                    d_711_v392_: D5
                    d_711_v392_ = D5_DC17(d_472_v227_)
                    d_712_v393_: _dafny.Seq
                    d_712_v393_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_472_v227_ for d_713_i49_ in range(default__.abs(458))]), d_175_v18_])
                    d_714_v394_: D6
                    d_714_v394_ = D6_DC21(d_712_v393_)
                    pat_let_tv25_ = d_712_v393_
                    def iife62_(_pat_let22_0):
                        def iife63_(d_715_dt__update__tmp_h8_):
                            def iife64_(_pat_let23_0):
                                def iife65_(d_716_dt__update_hcf30_h0_):
                                    return D6_DC21(d_716_dt__update_hcf30_h0_)
                                return iife65_(_pat_let23_0)
                            return iife64_(pat_let_tv25_)
                        return iife63_(_pat_let22_0)
                    d_665_v367_ = (d_665_v367_).set((d_231_v65_.f19) + (default__.fm0(d_234_v68_.f19, (d_711_v392_).cf23, (iife62_(d_714_v394_)).cf30, (0) - ((d_231_v65_).fm4(d_231_v65_.f19, len(d_175_v18_), d_153_v0_, ((d_665_v367_)[d_231_v65_.f19] if (d_231_v65_.f19) in (d_665_v367_) else d_153_v0_), d_155_globalState_)), d_155_globalState_)), (d_156_v2_) >= (d_231_v65_.f19))
                    if d_153_v0_:
                        d_717_v395_: int
                        d_718_v396_: _dafny.Seq
                        d_719_v397_: int
                        d_720_v398_: _dafny.MultiSet
                        out64_: int
                        out65_: _dafny.Seq
                        out66_: int
                        out67_: _dafny.MultiSet
                        out64_, out65_, out66_, out67_ = default__.m0(False, True, d_231_v65_.f19, d_153_v0_, d_155_globalState_)
                        d_717_v395_ = out64_
                        d_718_v396_ = out65_
                        d_719_v397_ = out66_
                        d_720_v398_ = out67_
                        d_721_v399_: _dafny.MultiSet
                        d_721_v399_ = _dafny.MultiSet([(0) - ((0) - (d_717_v395_)), d_234_v68_.f19])
                        d_722_v400_: _dafny.Map
                        d_722_v400_ = _dafny.Map({True: d_721_v399_})
                        d_723_v401_: _dafny.Array
                        nw111_ = _dafny.Array(None, 13)
                        nw111_[int(0)] = (d_721_v399_).intersection(d_721_v399_)
                        nw111_[int(1)] = ((d_722_v400_)[d_153_v0_] if (d_153_v0_) in (d_722_v400_) else _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([-279])).set(default__.safeIndex(d_231_v65_.f19, len(_dafny.SeqWithoutIsStrInference([-279]))), d_231_v65_.f19)))
                        nw111_[int(2)] = d_721_v399_
                        nw111_[int(3)] = _dafny.MultiSet([(0) - (d_717_v395_)])
                        nw111_[int(4)] = ((d_721_v399_).set(d_231_v65_.f19, default__.abs(d_231_v65_.f19))) | (default__.fm15(d_155_globalState_))
                        nw111_[int(5)] = d_721_v399_
                        nw111_[int(6)] = d_721_v399_
                        nw111_[int(7)] = d_721_v399_
                        nw111_[int(8)] = (((d_722_v400_)[d_153_v0_] if (d_153_v0_) in (d_722_v400_) else d_721_v399_)) - (d_721_v399_)
                        nw111_[int(9)] = (d_721_v399_) - (_dafny.MultiSet([d_234_v68_.f19, d_231_v65_.f19, d_231_v65_.f19, d_156_v2_, (0) - (d_156_v2_)]))
                        nw111_[int(10)] = default__.fm15(d_155_globalState_)
                        nw111_[int(11)] = d_721_v399_
                        nw111_[int(12)] = d_721_v399_
                        d_723_v401_ = nw111_
                        d_723_v401_ = d_723_v401_
                        d_724_v402_: _dafny.Map
                        d_724_v402_ = _dafny.Map({d_153_v0_: d_234_v68_.f19})
                        d_725_v403_: _dafny.Set
                        d_725_v403_ = _dafny.Set({(d_724_v402_) | (d_724_v402_)})
                        d_726_v404_: _dafny.Array
                        nw112_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
                        d_726_v404_ = nw112_
                        index143_ = default__.safeIndex(316, (d_726_v404_).length(0))
                        (d_726_v404_)[index143_] = (_dafny.SeqWithoutIsStrInference([(d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))] for d_727_i50_ in range(default__.abs(32))]) if not(d_153_v0_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mamgipjj")))
                        d_728_v405_: _dafny.Array
                        nw113_ = _dafny.Array(int(0), 9)
                        d_728_v405_ = nw113_
                        index144_ = default__.safeIndex(973, (d_728_v405_).length(0))
                        (d_728_v405_)[index144_] = -902
                        d_729_v406_: _dafny.Seq
                        d_729_v406_ = _dafny.SeqWithoutIsStrInference([d_231_v65_, d_234_v68_, d_231_v65_])
                        index145_ = default__.safeIndex(316, (d_726_v404_).length(0))
                        index146_ = default__.safeIndex(973, (d_728_v405_).length(0))
                        rhs143_ = (d_718_v396_) + (default__.fm5(len(d_729_v406_), d_155_globalState_))
                        rhs144_ = (d_725_v403_).intersection(d_725_v403_)
                        rhs145_ = ((d_175_v18_) + (d_718_v396_)).set(default__.safeIndex((d_231_v65_.f19) + (d_719_v397_), len((d_175_v18_) + (d_718_v396_))), (d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))])
                        rhs146_ = d_156_v2_
                        lhs114_ = d_726_v404_
                        lhs115_ = default__.safeIndex(316, (d_726_v404_).length(0))
                        lhs116_ = d_728_v405_
                        lhs117_ = default__.safeIndex(973, (d_728_v405_).length(0))
                        d_718_v396_ = rhs143_
                        d_725_v403_ = rhs144_
                        lhs114_[lhs115_] = rhs145_
                        lhs116_[lhs117_] = rhs146_
                        d_730_v407_: _dafny.Map
                        d_730_v407_ = _dafny.Map({(d_157_v3_) - (d_720_v398_): d_234_v68_.f19})
                        d_731_v408_: _dafny.Array
                        nw114_ = _dafny.Array(_dafny.Set({}), 28)
                        d_731_v408_ = nw114_
                        rhs147_ = (-575) + (len(_dafny.SeqWithoutIsStrInference([d_472_v227_ for d_732_i51_ in range(default__.abs(420))])))
                        rhs148_ = d_730_v407_
                        rhs149_ = d_153_v0_
                        rhs150_ = d_231_v65_
                        rhs151_ = d_731_v408_
                        lhs118_ = d_155_globalState_
                        lhs119_ = d_155_globalState_
                        lhs118_.f12 = rhs147_
                        d_730_v407_ = rhs148_
                        lhs119_.f3 = rhs149_
                        d_234_v68_ = rhs150_
                        d_731_v408_ = rhs151_
                        (d_155_globalState_).f3 = (d_234_v68_.f19) < (d_231_v65_.f19)
                    elif True:
                        (d_231_v65_).f19 = default__.safeDivisionInt(len((d_175_v18_) + (d_175_v18_)), d_234_v68_.f19)
                        d_233_v67_ = default__.fm10(d_155_globalState_)
                        d_153_v0_ = d_153_v0_
                        d_733_v410_: _dafny.Array
                        def lambda34_(d_734_v67_):
                            def lambda35_(d_735_i52_):
                                def iife66_():
                                    coll18_ = _dafny.Set()
                                    compr_18_: _dafny.Map
                                    for compr_18_ in (_dafny.Set({d_734_v67_})).Elements:
                                        d_736_v409_: _dafny.Map = compr_18_
                                        if (d_736_v409_) in (_dafny.Set({d_734_v67_})):
                                            coll18_ = coll18_.union(_dafny.Set([d_736_v409_]))
                                    return _dafny.Set(coll18_)
                                return (_dafny.Set({d_734_v67_})) - (iife66_()
                                )

                            return lambda35_

                        init18_ = lambda34_(d_233_v67_)
                        nw115_ = _dafny.Array(None, 9)
                        for i0_18_ in range(nw115_.length(0)):
                            nw115_[i0_18_] = init18_(i0_18_)
                        d_733_v410_ = nw115_
                        d_737_v411_: D1
                        d_737_v411_ = D1_DC5(default__.fm2(d_153_v0_, d_155_globalState_))
                        d_738_v412_: _dafny.Map
                        d_738_v412_ = _dafny.Map({d_233_v67_: d_737_v411_})
                        index147_ = default__.safeIndex(990, (d_733_v410_).length(0))
                        def iife67_():
                            coll19_ = _dafny.Set()
                            compr_19_: _dafny.Map
                            for compr_19_ in (d_738_v412_).keys.Elements:
                                d_739_v413_: _dafny.Map = compr_19_
                                if (d_739_v413_) in (d_738_v412_):
                                    coll19_ = coll19_.union(_dafny.Set([d_739_v413_]))
                            return _dafny.Set(coll19_)
                        (d_733_v410_)[index147_] = ((iife67_()
                        )) - (default__.fm16(d_155_globalState_))
                        d_740_v414_: _dafny.Set
                        d_740_v414_ = _dafny.Set({d_233_v67_})
                        index148_ = default__.safeIndex(990, (d_733_v410_).length(0))
                        (d_733_v410_)[index148_] = d_740_v414_
                        d_741_v415_: _dafny.Array
                        nw116_ = _dafny.Array(None, 19)
                        d_741_v415_ = nw116_
                        d_742_v418_: _dafny.Array
                        nw117_ = _dafny.Array(None, 21)
                        nw117_[int(0)] = d_233_v67_
                        nw117_[int(1)] = d_233_v67_
                        def iife68_():
                            coll20_ = _dafny.Map()
                            compr_20_: int
                            for compr_20_ in _dafny.IntegerRange(203, -284):
                                d_743_v416_: int = compr_20_
                                if ((203) <= (d_743_v416_)) and ((d_743_v416_) < (-284)):
                                    coll20_[(d_743_v416_) - (d_156_v2_)] = len(_dafny.Map({(d_473_v228_)[default__.safeIndex(241, (d_473_v228_).length(0))]: not(d_153_v0_)}))
                            return _dafny.Map(coll20_)
                        nw117_[int(2)] = iife68_()
                        
                        nw117_[int(3)] = d_233_v67_
                        nw117_[int(4)] = d_233_v67_
                        nw117_[int(5)] = d_233_v67_
                        nw117_[int(6)] = d_233_v67_
                        nw117_[int(7)] = d_233_v67_
                        nw117_[int(8)] = d_233_v67_
                        nw117_[int(9)] = d_233_v67_
                        nw117_[int(10)] = d_233_v67_
                        nw117_[int(11)] = d_233_v67_
                        nw117_[int(12)] = _dafny.Map({d_234_v68_.f19: 908})
                        nw117_[int(13)] = d_233_v67_
                        nw117_[int(14)] = d_233_v67_
                        nw117_[int(15)] = d_233_v67_
                        nw117_[int(16)] = d_233_v67_
                        nw117_[int(17)] = d_233_v67_
                        nw117_[int(18)] = d_233_v67_
                        nw117_[int(19)] = (d_233_v67_).set(d_234_v68_.f19, (0) - (d_231_v65_.f19))
                        def iife69_():
                            coll21_ = _dafny.Map()
                            compr_21_: int
                            for compr_21_ in (d_233_v67_).keys.Elements:
                                d_744_v417_: int = compr_21_
                                if (d_744_v417_) in (d_233_v67_):
                                    coll21_[(d_744_v417_) - (d_234_v68_.f19)] = (d_157_v3_).cardinality
                            return _dafny.Map(coll21_)
                        nw117_[int(20)] = iife69_()
                        
                        d_742_v418_ = nw117_
                        d_745_v419_: _dafny.Map
                        d_745_v419_ = _dafny.Map({d_742_v418_: d_741_v415_})
                        d_741_v415_ = ((d_745_v419_)[d_742_v418_] if (d_742_v418_) in (d_745_v419_) else d_741_v415_)
                    pass
            pass
        index149_ = default__.safeIndex(823, (d_376_v167_).length(0))
        (d_376_v167_)[index149_] = d_153_v0_
        d_746_v420_: _dafny.Array
        nw118_ = _dafny.Array(D3.default()(), 29)
        d_746_v420_ = nw118_
        index150_ = default__.safeIndex(823, (d_376_v167_).length(0))
        rhs152_ = (519) != (default__.safeModuloInt(d_231_v65_.f19, len(_dafny.SeqWithoutIsStrInference([not(d_153_v0_), d_153_v0_]))))
        rhs153_ = (d_746_v420_ if d_153_v0_ else d_746_v420_)
        lhs120_ = d_376_v167_
        lhs121_ = default__.safeIndex(823, (d_376_v167_).length(0))
        lhs120_[lhs121_] = rhs152_
        d_746_v420_ = rhs153_
        _dafny.print(_dafny.string_of(d_153_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v1_) == (_dafny.SeqWithoutIsStrInference([False, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_155_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_155_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_155_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_155_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_155_globalState_).f10).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_155_globalState_.f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f18) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_156_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v3_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v18_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v63_) == (_dafny.Map({_dafny.Set({False}): False, _dafny.Set({True}): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_230_v64_) == (_dafny.SeqWithoutIsStrInference([270]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_231_v65_.f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_232_v66_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v67_) == (_dafny.Map({1: 270}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_234_v68_.f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_235_v69_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_236_i8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_316_v126_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v167_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v167_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_377_v168_).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_377_v168_).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_377_v168_).cf26)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_377_v168_).cf26)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_377_v168_).cf27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_378_v169_).cf29).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_378_v169_).cf29).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_378_v169_).cf29).cf26)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_378_v169_).cf29).cf26)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_378_v169_).cf29).cf27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_379_v170_).cf29).cf29).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_379_v170_).cf29).cf29).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_379_v170_).cf29).cf29).cf26)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_379_v170_).cf29).cf29).cf26)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_379_v170_).cf29).cf29).cf27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_472_v227_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_473_v228_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_473_v228_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_473_v228_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_473_v228_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_552_v281_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_552_v281_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_552_v281_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[0]).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[0]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[0]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[1]).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[1]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[1]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[2]).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[2]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[2]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[3]).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[3]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[3]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[4]).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[4]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[4]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[5]).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[5]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[5]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[6]).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[6]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[6]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[7]).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[7]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[7]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[8]).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[8]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_660_v363_)[8]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_661_v365_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_663_v366_) == (_dafny.Map({650: _dafny.MultiSet([False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_664_i42_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), int(0), False)
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any), ('cf5', Any), ('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
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
        return lambda: D1_DC4(None)
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
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)

class D1_DC4(D1, NamedTuple('DC4', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({self.cf8.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D2_DC10)

class D2_DC8(D2, NamedTuple('DC8', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC10(D2, NamedTuple('DC10', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC10({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC10) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC12(_dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D3_DC13)

class D3_DC12(D3, NamedTuple('DC12', [('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC13(D3, NamedTuple('DC13', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC13({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC13) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC15()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D4_DC16)

class D4_DC15(D4, NamedTuple('DC15', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC16(D4, NamedTuple('DC16', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC16({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC16) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC18(int(0), False, _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D5_DC19)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D5_DC20)

class D5_DC18(D5, NamedTuple('DC18', [('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC19(D5, NamedTuple('DC19', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC19({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC19) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC17(D5, NamedTuple('DC17', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC20(D5, NamedTuple('DC20', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC20({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC20) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC22()
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

class D6_DC22(D6, NamedTuple('DC22', [])):
    def __dafnystr__(self) -> str:
        return f'D6.DC22'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC22)
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC23(D6, NamedTuple('DC23', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC23({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC23) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC21(D6, NamedTuple('DC21', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC21({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC21) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC24(D6, NamedTuple('DC24', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC24({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC24) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D7_DC25)

class D7_DC25(D7, NamedTuple('DC25', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC25({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC25) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D8_DC26)

class D8_DC26(D8, NamedTuple('DC26', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC26({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC26) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC28(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D9_DC28)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D9_DC29)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D9_DC30)

class D9_DC28(D9, NamedTuple('DC28', [('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC28({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC28) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC29(D9, NamedTuple('DC29', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC29({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC29) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC30(D9, NamedTuple('DC30', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC30({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC30) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f2: int = int(0)
        self.f3: bool = False
        self.f8: int = int(0)
        self.f12: int = int(0)
        self._f1: bool = False
        self._f4: int = int(0)
        self._f5: int = int(0)
        self._f6: bool = False
        self._f7: int = int(0)
        self._f9: int = int(0)
        self._f10: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f11: bool = False
        self._f13: bool = False
        self._f14: int = int(0)
        self._f15: bool = False
        self._f16: int = int(0)
        self._f17: int = int(0)
        self._f18: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18):
        (self).f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18

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
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11
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
    def f18(self):
        return self._f18

class C0:
    def  __init__(self):
        self.f19: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f19):
        (self).f19 = f19

    def fm3(self, globalState):
        return self.f19

    def fm4(self, p0, p1, p2, p3, globalState):
        return ((0) - (self.f19)) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False}))])))

