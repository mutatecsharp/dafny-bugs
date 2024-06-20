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
        return ((_dafny.Set({len(_dafny.Map({855: _dafny.CodePoint('c')})), -525, (0) - (-448)})) | (_dafny.Set({289, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, False, not(True), False, True]))}))}))).isdisjoint(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dfanxb"))), -110, -828}))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return _dafny.CodePoint('i')

    @staticmethod
    def fm2(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: str
            for compr_0_ in (_dafny.Map({_dafny.CodePoint('m'): True})).keys.Elements:
                d_0_v0_: str = compr_0_
                if (d_0_v0_) in (_dafny.Map({_dafny.CodePoint('m'): True})):
                    coll0_[d_0_v0_] = False
            return _dafny.Map(coll0_)
        return (iife0_()
        ) | (_dafny.Map({_dafny.CodePoint('f'): True}))

    @staticmethod
    def fm3(globalState):
        return (0) - ((default__.safeModuloInt(len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([False, False]))})), len(_dafny.Set({406, 244, 694, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tudpflvdl")))})))) - (default__.safeModuloInt(len(_dafny.Map({329: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmmeb")))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdcukv"))))))

    @staticmethod
    def fm4(p0, p1, p2, globalState):
        return _dafny.Map({not((_dafny.Set({831, (_dafny.MultiSet([len(_dafny.Map({False: 251}))])).cardinality})).isdisjoint(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "famu"))), len((D3_DC8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsdvonkfs")), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))) for d_1_i0_ in range(default__.abs(467))]))).cf13)}))): True})

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([not(False)])) | (_dafny.MultiSet([not(True), False, not(True)]))) | ((_dafny.MultiSet([True, True, False, True, False])).intersection(_dafny.MultiSet([False])))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        return D3_DC8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jucvf")), (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2_i0_ in range(default__.abs(-165))])), len(_dafny.Map({_dafny.Set({604, 866}): _dafny.CodePoint('u')}))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(True)])) for d_3_i1_ in range(default__.abs(512))])))

    @staticmethod
    def fm7(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: _dafny.Seq
            for compr_1_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")): not(False)})).keys.Elements:
                d_5_v0_: _dafny.Seq = compr_1_
                if (d_5_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")): not(False)})):
                    coll1_[d_5_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcmexjuub"))
            return _dafny.Map(coll1_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_4_i0_ in range(default__.abs(132))]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxf"))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epp")): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjb"))}))) | ((iife1_()
        ) | ((D6_DC16(_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_6_i1_ in range(default__.abs(-223))]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsqfg"))}))).cf29))

    @staticmethod
    def fm8(p0, globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obosvty"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_7_i0_ in range(default__.abs(614))])): False})

    @staticmethod
    def fm9(p0, p1, globalState):
        return D3_DC7((_dafny.Map({True: False})) | (_dafny.Map({True: True})))

    @staticmethod
    def fm10(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_8_i0_ in range(default__.abs(-538))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_9_i1_ in range(default__.abs(616))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fn"))))

    @staticmethod
    def fm11(globalState):
        return _dafny.Map({(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([not(False)]))})).isdisjoint(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, True]))})): (-84 if False else -734)})

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        hi0_ = (p3) + (p0)
        for d_10_i0_ in range(p3, hi0_):
            source0_ = D2_DC6((d_10_i0_) != (p0))
            if source0_.is_DC6:
                d_11___mcc_h0_ = source0_.cf10
                d_12_cf10_ = d_11___mcc_h0_
                d_13_v0_: C0
                nw0_ = C0()
                nw0_.ctor__(p0)
                d_13_v0_ = nw0_
                d_14_v1_: _dafny.Array
                nw1_ = _dafny.Array(int(0), 20)
                d_14_v1_ = nw1_
                index0_ = default__.safeIndex(348, (d_14_v1_).length(0))
                (d_14_v1_)[index0_] = len(default__.fm10(not(p2), globalState))
                index1_ = default__.safeIndex(348, (d_14_v1_).length(0))
                (d_14_v1_)[index1_] = default__.safeModuloInt((p0) + (d_10_i0_), p3)
                d_15_v2_: _dafny.Seq
                d_15_v2_ = _dafny.SeqWithoutIsStrInference([p2])
                d_15_v2_ = d_15_v2_
                index2_ = default__.safeIndex(348, (d_14_v1_).length(0))
                (d_14_v1_)[index2_] = -844
            elif True:
                d_16___mcc_h1_ = source0_.cf9
                d_17_cf9_ = d_16___mcc_h1_
                d_18_v3_: _dafny.Array
                nw2_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_18_v3_ = nw2_
                d_19_v4_: _dafny.Array
                nw3_ = _dafny.Array(False, 18)
                d_19_v4_ = nw3_
                index3_ = default__.safeIndex(653, (d_18_v3_).length(0))
                (d_18_v3_)[index3_] = d_19_v4_
                d_20_v5_: _dafny.Array
                def lambda0_(d_21_i1_):
                    return default__.safeModuloInt(d_21_i1_, 805)

                init0_ = lambda0_
                nw4_ = _dafny.Array(None, 3)
                for i0_0_ in range(nw4_.length(0)):
                    nw4_[i0_0_] = init0_(i0_0_)
                d_20_v5_ = nw4_
                index4_ = default__.safeIndex(564, (d_20_v5_).length(0))
                (d_20_v5_)[index4_] = d_10_i0_
                d_22_v6_: _dafny.Seq
                d_22_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfaaiqx"))
                d_23_v7_: _dafny.Map
                d_23_v7_ = _dafny.Map({(p2 if (d_17_cf9_)[default__.safeIndex(d_10_i0_, len(d_17_cf9_))] else p2): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ma"))) + (d_22_v6_)})
                d_24_v8_: _dafny.Map
                d_24_v8_ = _dafny.Map({p2: p2})
                d_25_v9_: str
                d_25_v9_ = _dafny.CodePoint('p')
                d_26_v10_: _dafny.Set
                d_26_v10_ = _dafny.Set({p2})
                index5_ = default__.safeIndex(653, (d_18_v3_).length(0))
                index6_ = default__.safeIndex(564, (d_20_v5_).length(0))
                rhs0_ = ((((d_23_v7_)[((d_24_v8_)[p2] if (p2) in (d_24_v8_) else p2)] if (((d_24_v8_)[p2] if (p2) in (d_24_v8_) else p2)) in (d_23_v7_) else d_22_v6_)).set(default__.safeIndex((0) - ((0) - (p3)), len(((d_23_v7_)[((d_24_v8_)[p2] if (p2) in (d_24_v8_) else p2)] if (((d_24_v8_)[p2] if (p2) in (d_24_v8_) else p2)) in (d_23_v7_) else d_22_v6_))), d_25_v9_)).set(default__.safeIndex(default__.safeModuloInt(default__.fm3(globalState), len(d_22_v6_)), len((((d_23_v7_)[((d_24_v8_)[p2] if (p2) in (d_24_v8_) else p2)] if (((d_24_v8_)[p2] if (p2) in (d_24_v8_) else p2)) in (d_23_v7_) else d_22_v6_)).set(default__.safeIndex((0) - ((0) - (p3)), len(((d_23_v7_)[((d_24_v8_)[p2] if (p2) in (d_24_v8_) else p2)] if (((d_24_v8_)[p2] if (p2) in (d_24_v8_) else p2)) in (d_23_v7_) else d_22_v6_))), d_25_v9_))), d_25_v9_)
                rhs1_ = d_19_v4_
                rhs2_ = d_10_i0_
                rhs3_ = ((p3 if p2 else len(d_26_v10_)) if p2 else d_10_i0_)
                lhs0_ = globalState
                lhs1_ = d_18_v3_
                lhs2_ = default__.safeIndex(653, (d_18_v3_).length(0))
                lhs3_ = d_20_v5_
                lhs4_ = default__.safeIndex(564, (d_20_v5_).length(0))
                lhs5_ = globalState
                lhs0_.f4 = rhs0_
                lhs1_[lhs2_] = rhs1_
                lhs3_[lhs4_] = rhs2_
                lhs5_.f7 = rhs3_
                index7_ = default__.safeIndex(686, (d_19_v4_).length(0))
                (d_19_v4_)[index7_] = p2
                index8_ = default__.safeIndex(686, (d_19_v4_).length(0))
                (d_19_v4_)[index8_] = (p2 if False else p2)
                d_27_v11_: _dafny.Map
                d_27_v11_ = _dafny.Map({d_25_v9_: p3})
                d_28_v12_: D0
                d_28_v12_ = D0_DC2(p0)
                d_29_v13_: _dafny.Map
                d_29_v13_ = _dafny.Map({p2: p0})
                d_30_v14_: _dafny.Set
                d_30_v14_ = _dafny.Set({d_28_v12_, d_28_v12_, D0_DC2((d_20_v5_)[default__.safeIndex(564, (d_20_v5_).length(0))]), D0_DC2(len(d_29_v13_)), d_28_v12_})
                index9_ = default__.safeIndex(686, (d_19_v4_).length(0))
                rhs4_ = default__.safeModuloInt(len(d_27_v11_), (len(d_30_v14_)) + (p0))
                rhs5_ = p2
                rhs6_ = (d_22_v6_) + (d_22_v6_)
                lhs6_ = globalState
                lhs7_ = d_19_v4_
                lhs8_ = default__.safeIndex(686, (d_19_v4_).length(0))
                lhs9_ = globalState
                lhs6_.f7 = rhs4_
                lhs7_[lhs8_] = rhs5_
                lhs9_.f4 = rhs6_
                d_31_v15_: _dafny.Map
                d_31_v15_ = _dafny.Map({not((d_19_v4_)[default__.safeIndex(686, (d_19_v4_).length(0))]): (d_18_v3_)[default__.safeIndex(653, (d_18_v3_).length(0))]})
                d_32_v16_: _dafny.Map
                d_32_v16_ = _dafny.Map({(d_31_v15_) | (d_31_v15_): default__.safeModuloInt((d_20_v5_)[default__.safeIndex(564, (d_20_v5_).length(0))], p3)})
                d_32_v16_ = (d_32_v16_).set((d_31_v15_).set(False, (d_18_v3_)[default__.safeIndex(653, (d_18_v3_).length(0))]), p3)
            d_33_v17_: D0
            d_33_v17_ = D0_DC2(993)
            def iife2_(_pat_let0_0):
                def iife3_(d_34_dt__update__tmp_h0_):
                    def iife5_():
                        coll2_ = _dafny.Map()
                        compr_2_: int
                        for compr_2_ in _dafny.IntegerRange(-54, 138):
                            d_35_v18_: int = compr_2_
                            if ((-54) <= (d_35_v18_)) and ((d_35_v18_) < (138)):
                                coll2_[(d_35_v18_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqe"))))] = 776
                        return _dafny.Map(coll2_)
                    def iife4_(_pat_let1_0):
                        def iife6_(d_36_dt__update_hcf4_h0_):
                            return D0_DC2(d_36_dt__update_hcf4_h0_)
                        return iife6_(_pat_let1_0)
                    return iife4_(len(iife5_()
                    ))
                return iife3_(_pat_let0_0)
            source1_ = iife2_(d_33_v17_)
            if source1_.is_DC1:
                d_37___mcc_h2_ = source1_.cf1
                d_38___mcc_h3_ = source1_.cf2
                d_39___mcc_h4_ = source1_.cf3
                d_40_cf3_ = d_39___mcc_h4_
                d_41_cf2_ = d_38___mcc_h3_
                d_42_cf1_ = d_37___mcc_h2_
                d_43_v19_: _dafny.Array
                nw5_ = _dafny.Array(_dafny.Seq({}), 10)
                d_43_v19_ = nw5_
                d_44_v20_: C0
                nw6_ = C0()
                nw6_.ctor__(p3)
                d_44_v20_ = nw6_
                d_45_v21_: _dafny.Seq
                d_45_v21_ = _dafny.SeqWithoutIsStrInference([d_44_v20_, d_44_v20_])
                index10_ = default__.safeIndex(125, (d_43_v19_).length(0))
                (d_43_v19_)[index10_] = (d_45_v21_) + (d_45_v21_)
                index11_ = default__.safeIndex(125, (d_43_v19_).length(0))
                (d_43_v19_)[index11_] = _dafny.SeqWithoutIsStrInference([d_44_v20_])
                d_40_cf3_ = default__.safeModuloInt(p0, d_40_cf3_)
                d_40_cf3_ = default__.safeModuloInt(d_42_cf1_, d_10_i0_)
                d_46_v22_: _dafny.Seq
                d_46_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
                d_47_v23_: str
                d_47_v23_ = _dafny.CodePoint('u')
                (globalState).f4 = (((d_46_v22_) + (d_46_v22_)).set(default__.safeIndex(d_42_cf1_, len((d_46_v22_) + (d_46_v22_))), d_47_v23_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihfuqk"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unigk"))))
            elif source1_.is_DC2:
                d_48___mcc_h5_ = source1_.cf4
                d_49_cf4_ = d_48___mcc_h5_
                d_50_v24_: _dafny.Array
                nw7_ = _dafny.Array(_dafny.Seq({}), 21)
                d_50_v24_ = nw7_
                d_51_v25_: _dafny.Seq
                d_51_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vp"))
                d_52_v26_: _dafny.Seq
                d_52_v26_ = _dafny.SeqWithoutIsStrInference([d_51_v25_])
                index12_ = default__.safeIndex(444, (d_50_v24_).length(0))
                (d_50_v24_)[index12_] = d_52_v26_
                index13_ = default__.safeIndex(444, (d_50_v24_).length(0))
                (d_50_v24_)[index13_] = (_dafny.SeqWithoutIsStrInference([d_51_v25_ for d_53_i2_ in range(default__.abs(652))])) + ((d_52_v26_) + (d_52_v26_))
                d_54_v27_: _dafny.Array
                def lambda1_(d_55_cf4_):
                    def lambda2_(d_56_i3_):
                        return default__.safeModuloInt(d_56_i3_, d_55_cf4_)

                    return lambda2_

                init1_ = lambda1_(d_49_cf4_)
                nw8_ = _dafny.Array(None, 21)
                for i0_1_ in range(nw8_.length(0)):
                    nw8_[i0_1_] = init1_(i0_1_)
                d_54_v27_ = nw8_
                index14_ = default__.safeIndex(363, (d_54_v27_).length(0))
                (d_54_v27_)[index14_] = p0
                index15_ = default__.safeIndex(363, (d_54_v27_).length(0))
                (d_54_v27_)[index15_] = d_10_i0_
                d_57_v28_: bool
                d_57_v28_ = False
                d_57_v28_ = d_57_v28_
                d_57_v28_ = d_57_v28_
            elif source1_.is_DC3:
                d_58___mcc_h6_ = source1_.cf5
                d_59___mcc_h7_ = source1_.cf6
                d_60___mcc_h8_ = source1_.cf7
                d_61_cf7_ = d_60___mcc_h8_
                d_62_cf6_ = d_59___mcc_h7_
                d_63_cf5_ = d_58___mcc_h6_
                (globalState).f7 = p3
                d_62_cf6_ = d_62_cf6_
                d_64_v29_: C0
                nw9_ = C0()
                nw9_.ctor__(-122)
                d_64_v29_ = nw9_
                (d_64_v29_).f8 = d_61_cf7_
            elif True:
                d_65___mcc_h9_ = source1_.cf0
                d_66_cf0_ = d_65___mcc_h9_
                d_67_v30_: C0
                nw10_ = C0()
                nw10_.ctor__(p0)
                d_67_v30_ = nw10_
                d_68_v31_: _dafny.Array
                nw11_ = _dafny.Array(False, 1)
                d_68_v31_ = nw11_
                d_69_v32_: _dafny.Seq
                d_69_v32_ = _dafny.SeqWithoutIsStrInference([d_68_v31_, d_68_v31_, d_68_v31_])
                d_69_v32_ = d_69_v32_
                d_70_v33_: _dafny.Map
                d_70_v33_ = _dafny.Map({d_67_v30_.f8: (p3) != (d_10_i0_)})
                d_70_v33_ = (d_70_v33_).set(p3, p2)
                d_71_v34_: bool
                d_71_v34_ = True
                d_71_v34_ = d_71_v34_
            d_72_v35_: bool
            d_72_v35_ = False
            d_73_v36_: _dafny.Seq
            d_73_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
            d_72_v35_ = (len((d_73_v36_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_74_i4_ in range(default__.abs(180))])))) <= (p0)
            d_75_v37_: _dafny.Array
            def lambda3_(d_76_i5_):
                return (d_76_i5_) - (-109)

            init2_ = lambda3_
            nw12_ = _dafny.Array(None, 9)
            for i0_2_ in range(nw12_.length(0)):
                nw12_[i0_2_] = init2_(i0_2_)
            d_75_v37_ = nw12_
            index16_ = default__.safeIndex(447, (d_75_v37_).length(0))
            (d_75_v37_)[index16_] = (p3) - (p3)
            index17_ = default__.safeIndex(447, (d_75_v37_).length(0))
            (d_75_v37_)[index17_] = p0
        d_77_i6_: int
        d_77_i6_ = 0
        with _dafny.label("0"):
            while p2:
                with _dafny.c_label("0"):
                    if (d_77_i6_) >= (100):
                        raise _dafny.Break("0")
                    d_77_i6_ = (d_77_i6_) + (1)
                    d_78_v38_: str
                    d_78_v38_ = _dafny.CodePoint('k')
                    d_79_v39_: _dafny.Map
                    d_79_v39_ = _dafny.Map({d_78_v38_: p2})
                    d_80_v40_: _dafny.MultiSet
                    d_80_v40_ = _dafny.MultiSet([d_79_v39_, d_79_v39_])
                    d_81_v41_: D0
                    d_81_v41_ = D0_DC3(default__.fm11(globalState), default__.fm0(not(p2), p2, d_80_v40_, p0, globalState), p0)
                    (globalState).f7 = (d_81_v41_).cf7
                    d_82_v42_: C0
                    nw13_ = C0()
                    nw13_.ctor__(len(_dafny.SeqWithoutIsStrInference([d_78_v38_ for d_83_i7_ in range(default__.abs(317))])))
                    d_82_v42_ = nw13_
                    d_84_v43_: D4
                    d_84_v43_ = D4_DC10(p2, d_82_v42_)
                    pat_let_tv0_ = p2
                    pat_let_tv1_ = d_82_v42_
                    d_85_v44_: _dafny.Array
                    nw14_ = _dafny.Array(None, 29)
                    nw14_[int(0)] = d_84_v43_
                    def iife7_(_pat_let2_0):
                        def iife8_(d_86_dt__update__tmp_h1_):
                            def iife9_(_pat_let3_0):
                                def iife10_(d_87_dt__update_hcf15_h0_):
                                    return D4_DC10(d_87_dt__update_hcf15_h0_, (d_86_dt__update__tmp_h1_).cf16)
                                return iife10_(_pat_let3_0)
                            return iife9_(pat_let_tv0_)
                        return iife8_(_pat_let2_0)
                    nw14_[int(1)] = iife7_(d_84_v43_)
                    nw14_[int(2)] = d_84_v43_
                    nw14_[int(3)] = d_84_v43_
                    nw14_[int(4)] = D4_DC10(p2, d_82_v42_)
                    nw14_[int(5)] = d_84_v43_
                    def iife11_(_pat_let4_0):
                        def iife12_(d_88_dt__update__tmp_h2_):
                            def iife13_(_pat_let5_0):
                                def iife14_(d_89_dt__update_hcf16_h0_):
                                    return D4_DC10((d_88_dt__update__tmp_h2_).cf15, d_89_dt__update_hcf16_h0_)
                                return iife14_(_pat_let5_0)
                            return iife13_(pat_let_tv1_)
                        return iife12_(_pat_let4_0)
                    nw14_[int(6)] = iife11_(d_84_v43_)
                    nw14_[int(7)] = d_84_v43_
                    nw14_[int(8)] = d_84_v43_
                    nw14_[int(9)] = D4_DC10(p2, d_82_v42_)
                    nw14_[int(10)] = d_84_v43_
                    nw14_[int(11)] = d_84_v43_
                    nw14_[int(12)] = d_84_v43_
                    nw14_[int(13)] = d_84_v43_
                    nw14_[int(14)] = d_84_v43_
                    nw14_[int(15)] = d_84_v43_
                    nw14_[int(16)] = d_84_v43_
                    nw14_[int(17)] = d_84_v43_
                    nw14_[int(18)] = d_84_v43_
                    nw14_[int(19)] = d_84_v43_
                    nw14_[int(20)] = d_84_v43_
                    nw14_[int(21)] = D4_DC10(p2, d_82_v42_)
                    nw14_[int(22)] = D4_DC10(p2, d_82_v42_)
                    nw14_[int(23)] = D4_DC10(not(p2), d_82_v42_)
                    nw14_[int(24)] = d_84_v43_
                    nw14_[int(25)] = d_84_v43_
                    nw14_[int(26)] = D4_DC10(p2, d_82_v42_)
                    nw14_[int(27)] = d_84_v43_
                    nw14_[int(28)] = d_84_v43_
                    d_85_v44_ = nw14_
                    d_90_v45_: _dafny.Array
                    nw15_ = _dafny.Array(None, 17)
                    nw15_[int(0)] = d_85_v44_
                    nw15_[int(1)] = d_85_v44_
                    nw15_[int(2)] = d_85_v44_
                    nw15_[int(3)] = d_85_v44_
                    nw15_[int(4)] = d_85_v44_
                    nw15_[int(5)] = d_85_v44_
                    nw15_[int(6)] = d_85_v44_
                    nw15_[int(7)] = d_85_v44_
                    nw15_[int(8)] = d_85_v44_
                    nw15_[int(9)] = d_85_v44_
                    nw15_[int(10)] = d_85_v44_
                    nw15_[int(11)] = d_85_v44_
                    nw15_[int(12)] = d_85_v44_
                    nw15_[int(13)] = d_85_v44_
                    nw15_[int(14)] = d_85_v44_
                    nw15_[int(15)] = (d_85_v44_ if p2 else d_85_v44_)
                    nw15_[int(16)] = d_85_v44_
                    d_90_v45_ = nw15_
                    index18_ = default__.safeIndex(817, (d_90_v45_).length(0))
                    (d_90_v45_)[index18_] = d_85_v44_
                    pat_let_tv2_ = p2
                    index19_ = default__.safeIndex(817, (d_90_v45_).length(0))
                    nw16_ = _dafny.Array(None, 2)
                    def iife15_(_pat_let6_0):
                        def iife16_(d_91_dt__update__tmp_h3_):
                            def iife17_(_pat_let7_0):
                                def iife18_(d_92_dt__update_hcf15_h1_):
                                    return D4_DC10(d_92_dt__update_hcf15_h1_, (d_91_dt__update__tmp_h3_).cf16)
                                return iife18_(_pat_let7_0)
                            return iife17_(pat_let_tv2_)
                        return iife16_(_pat_let6_0)
                    nw16_[int(0)] = iife15_(d_84_v43_)
                    nw16_[int(1)] = d_84_v43_
                    (d_90_v45_)[index19_] = nw16_
                    d_93_v46_: C0
                    nw17_ = C0()
                    nw17_.ctor__(p0)
                    d_93_v46_ = nw17_
                    d_94_v47_: _dafny.Seq
                    d_94_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogpsjgmip"))
                    d_95_v48_: _dafny.Map
                    d_95_v48_ = _dafny.Map({d_82_v42_: d_94_v47_})
                    d_95_v48_ = (d_95_v48_).set(d_93_v46_, d_94_v47_)
                    pass
            pass
        d_96_v49_: C0
        nw18_ = C0()
        nw18_.ctor__(p3)
        d_96_v49_ = nw18_
        d_97_v50_: str
        d_97_v50_ = _dafny.CodePoint('s')
        d_98_v51_: _dafny.Map
        d_98_v51_ = _dafny.Map({d_96_v49_: d_97_v50_})
        d_99_v53_: _dafny.Array
        def lambda4_(d_100_p0_):
            def lambda5_(d_101_i8_):
                def iife19_():
                    coll3_ = _dafny.Set()
                    compr_3_: int
                    for compr_3_ in (_dafny.Set({d_100_p0_})).Elements:
                        d_102_v52_: int = compr_3_
                        if (d_102_v52_) in (_dafny.Set({d_100_p0_})):
                            coll3_ = coll3_.union(_dafny.Set([(d_102_v52_) * (len(_dafny.SeqWithoutIsStrInference([True])))]))
                    return _dafny.Set(coll3_)
                return (d_101_i8_) * ((0) - (len(iife19_()
                )))

            return lambda5_

        init3_ = lambda4_(p0)
        nw19_ = _dafny.Array(None, 13)
        for i0_3_ in range(nw19_.length(0)):
            nw19_[i0_3_] = init3_(i0_3_)
        d_99_v53_ = nw19_
        index20_ = default__.safeIndex(952, (d_99_v53_).length(0))
        (d_99_v53_)[index20_] = (d_96_v49_.f8) - ((0) - (p3))
        d_103_v54_: _dafny.Map
        d_103_v54_ = _dafny.Map({not(p2): p2})
        index21_ = default__.safeIndex(952, (d_99_v53_).length(0))
        rhs7_ = _dafny.Map({d_96_v49_: d_97_v50_})
        rhs8_ = p0
        rhs9_ = (len(d_103_v54_)) * ((p3) - (d_96_v49_.f8))
        rhs10_ = d_96_v49_
        lhs10_ = d_99_v53_
        lhs11_ = default__.safeIndex(952, (d_99_v53_).length(0))
        lhs12_ = globalState
        d_98_v51_ = rhs7_
        lhs10_[lhs11_] = rhs8_
        lhs12_.f7 = rhs9_
        d_96_v49_ = rhs10_
        d_104_v55_: _dafny.Array
        nw20_ = _dafny.Array(False, 10)
        d_104_v55_ = nw20_
        d_104_v55_ = d_104_v55_
        d_105_v56_: _dafny.Seq
        d_105_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skwswvs"))
        d_106_v57_: _dafny.Map
        d_106_v57_ = _dafny.Map({973: d_105_v56_})
        d_107_v58_: _dafny.Map
        d_107_v58_ = _dafny.Map({True: len((d_105_v56_) + (((d_106_v57_)[d_96_v49_.f8] if (d_96_v49_.f8) in (d_106_v57_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kefocyf")))))})
        d_108_v59_: _dafny.Map
        d_108_v59_ = _dafny.Map({d_97_v50_: p2})
        d_109_v60_: _dafny.MultiSet
        d_109_v60_ = _dafny.MultiSet([d_108_v59_])
        d_110_v61_: _dafny.Set
        d_110_v61_ = _dafny.Set({p2, p2, p2, p2, p2})
        d_111_v62_: _dafny.Seq
        d_111_v62_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({default__.fm0(p2, p2, d_109_v60_, p0, globalState)}), d_110_v61_])
        d_112_v63_: _dafny.Seq
        d_112_v63_ = _dafny.SeqWithoutIsStrInference([-8, d_96_v49_.f8])
        d_107_v58_ = (d_107_v58_).set(((d_111_v62_)[default__.safeIndex(len(d_112_v63_), len(d_111_v62_))]) in (_dafny.SeqWithoutIsStrInference([d_110_v61_])), (d_99_v53_)[default__.safeIndex(952, (d_99_v53_).length(0))])
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_104_v55_).length(0)):
            d_113_i9_: int = guard_loop_0_
            if (True) and (((0) <= (d_113_i9_)) and ((d_113_i9_) < ((d_104_v55_).length(0)))):
                (d_104_v55_)[(d_113_i9_)] = (d_110_v61_).issubset(d_110_v61_)

    @staticmethod
    def Main(noArgsParameter__):
        d_114_v0_: bool
        d_114_v0_ = False
        d_115_v1_: _dafny.Map
        d_115_v1_ = _dafny.Map({d_114_v0_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmoircrjr")))})
        d_116_v2_: _dafny.Array
        def lambda6_(d_117_i0_):
            return default__.safeModuloInt(d_117_i0_, 448)

        init4_ = lambda6_
        nw21_ = _dafny.Array(None, 1)
        for i0_4_ in range(nw21_.length(0)):
            nw21_[i0_4_] = init4_(i0_4_)
        d_116_v2_ = nw21_
        d_118_v3_: _dafny.Array
        def lambda7_(d_119_i1_):
            return False

        init5_ = lambda7_
        nw22_ = _dafny.Array(None, 12)
        for i0_5_ in range(nw22_.length(0)):
            nw22_[i0_5_] = init5_(i0_5_)
        d_118_v3_ = nw22_
        d_120_v4_: _dafny.Seq
        d_120_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnncmt"))
        d_121_globalState_: GlobalState
        nw23_ = GlobalState()
        nw23_.ctor__(d_115_v1_, d_116_v2_, d_118_v3_, 287, d_120_v4_, False, d_120_v4_, 632)
        d_121_globalState_ = nw23_
        d_122_v5_: int
        d_122_v5_ = 681
        d_123_v6_: _dafny.Map
        d_123_v6_ = _dafny.Map({default__.fm1(d_114_v0_, d_122_v5_, d_122_v5_, d_121_globalState_): d_114_v0_})
        d_124_v7_: _dafny.Seq
        d_124_v7_ = _dafny.SeqWithoutIsStrInference([d_123_v6_])
        d_125_v8_: _dafny.MultiSet
        d_125_v8_ = _dafny.MultiSet([d_123_v6_, d_123_v6_, d_123_v6_])
        d_114_v0_ = default__.fm0(default__.fm0(d_114_v0_, d_114_v0_, _dafny.MultiSet(d_124_v7_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcjlcyyhn"))), d_121_globalState_), default__.fm0(d_114_v0_, d_114_v0_, _dafny.MultiSet([d_123_v6_, d_123_v6_, default__.fm2(not(d_114_v0_), d_121_globalState_)]), d_122_v5_, d_121_globalState_), (d_125_v8_) - (_dafny.MultiSet(d_124_v7_)), ((0) - (len(d_120_v4_))) * (d_122_v5_), d_121_globalState_)
        index22_ = default__.safeIndex(619, (d_118_v3_).length(0))
        (d_118_v3_)[index22_] = True
        index23_ = default__.safeIndex(619, (d_118_v3_).length(0))
        (d_118_v3_)[index23_] = (d_122_v5_) >= (d_122_v5_)
        d_126_v9_: _dafny.Seq
        d_126_v9_ = _dafny.SeqWithoutIsStrInference([(d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], d_114_v0_])
        d_127_v10_: _dafny.Map
        d_127_v10_ = _dafny.Map({d_114_v0_: d_116_v2_})
        d_128_i2_: int
        d_128_i2_ = 0
        with _dafny.label("1"):
            while (d_126_v9_)[default__.safeIndex(len((d_127_v10_) | (d_127_v10_)), len(d_126_v9_))]:
                with _dafny.c_label("1"):
                    if (d_128_i2_) >= (100):
                        raise _dafny.Break("1")
                    d_128_i2_ = (d_128_i2_) + (1)
                    d_129_v11_: _dafny.Seq
                    d_129_v11_ = _dafny.SeqWithoutIsStrInference([d_122_v5_])
                    d_130_v12_: str
                    d_130_v12_ = _dafny.CodePoint('i')
                    d_131_v13_: _dafny.Map
                    d_131_v13_ = _dafny.Map({d_120_v4_: d_130_v12_})
                    d_132_v14_: _dafny.MultiSet
                    d_132_v14_ = _dafny.MultiSet([(d_129_v11_)[default__.safeIndex(len(d_129_v11_), len(d_129_v11_))], d_122_v5_, (0) - (len(d_131_v13_)), d_122_v5_])
                    default__.m0((0) - (d_122_v5_), d_132_v14_, ((d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]) == ((d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]), (0) - (d_122_v5_), d_121_globalState_)
                    d_133_v15_: _dafny.Map
                    d_133_v15_ = _dafny.Map({((d_120_v4_) + (d_120_v4_)).set(default__.safeIndex(-825, len((d_120_v4_) + (d_120_v4_))), d_130_v12_): (not(d_114_v0_) if (d_126_v9_)[default__.safeIndex(d_122_v5_, len(d_126_v9_))] else (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))])})
                    d_134_v17_: _dafny.MultiSet
                    d_134_v17_ = _dafny.MultiSet([d_120_v4_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npfn"))])
                    def iife20_():
                        coll4_ = _dafny.Map()
                        compr_4_: _dafny.Seq
                        for compr_4_ in (d_134_v17_).Elements:
                            d_135_v16_: _dafny.Seq = compr_4_
                            if (d_135_v16_) in (d_134_v17_):
                                coll4_[d_135_v16_] = d_114_v0_
                        return _dafny.Map(coll4_)
                    rhs11_ = (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]
                    rhs12_ = (0) - ((d_122_v5_) - (d_122_v5_))
                    rhs13_ = iife20_()

                    lhs13_ = d_121_globalState_
                    d_114_v0_ = rhs11_
                    lhs13_.f7 = rhs12_
                    d_133_v15_ = rhs13_
                    if not (not ((d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]) or (d_114_v0_)) or ((d_122_v5_) in (d_132_v14_)):
                        index24_ = default__.safeIndex(788, (d_116_v2_).length(0))
                        (d_116_v2_)[index24_] = d_122_v5_
                        index25_ = default__.safeIndex(788, (d_116_v2_).length(0))
                        (d_116_v2_)[index25_] = d_122_v5_
                        index26_ = default__.safeIndex(788, (d_116_v2_).length(0))
                        (d_116_v2_)[index26_] = d_122_v5_
                        (d_121_globalState_).f7 = ((d_116_v2_)[default__.safeIndex(788, (d_116_v2_).length(0))]) - (default__.safeModuloInt(default__.fm3(d_121_globalState_), (0) - (d_122_v5_)))
                        d_136_v18_: _dafny.Array
                        def lambda8_(d_137_v0_):
                            def lambda9_(d_138_i3_):
                                return _dafny.SeqWithoutIsStrInference([d_137_v0_])

                            return lambda9_

                        init6_ = lambda8_(d_114_v0_)
                        nw24_ = _dafny.Array(None, 7)
                        for i0_6_ in range(nw24_.length(0)):
                            nw24_[i0_6_] = init6_(i0_6_)
                        d_136_v18_ = nw24_
                        index27_ = default__.safeIndex(921, (d_136_v18_).length(0))
                        (d_136_v18_)[index27_] = (d_126_v9_) + (d_126_v9_)
                        index28_ = default__.safeIndex(921, (d_136_v18_).length(0))
                        (d_136_v18_)[index28_] = d_126_v9_
                        d_114_v0_ = ((d_136_v18_)[default__.safeIndex(921, (d_136_v18_).length(0))])[default__.safeIndex((_dafny.MultiSet([(d_136_v18_)[default__.safeIndex(921, (d_136_v18_).length(0))]])).cardinality, len((d_136_v18_)[default__.safeIndex(921, (d_136_v18_).length(0))]))]
                    elif True:
                        d_139_v19_: C0
                        nw25_ = C0()
                        nw25_.ctor__(d_122_v5_)
                        d_139_v19_ = nw25_
                        d_140_v20_: bool
                        d_141_v21_: bool
                        d_142_v22_: int
                        d_143_v23_: str
                        out0_: bool
                        out1_: bool
                        out2_: int
                        out3_: str
                        out0_, out1_, out2_, out3_ = (d_139_v19_).m1(d_121_globalState_)
                        d_140_v20_ = out0_
                        d_141_v21_ = out1_
                        d_142_v22_ = out2_
                        d_143_v23_ = out3_
                        d_129_v11_ = (d_129_v11_) + (d_129_v11_)
                        d_144_v24_: _dafny.Map
                        d_144_v24_ = _dafny.Map({d_139_v19_: d_139_v19_})
                        d_145_v25_: _dafny.Map
                        d_145_v25_ = _dafny.Map({587: d_144_v24_})
                        d_146_v26_: _dafny.MultiSet
                        d_146_v26_ = _dafny.MultiSet([d_139_v19_, d_139_v19_])
                        default__.m0(len((d_120_v4_) + (d_120_v4_)), d_132_v14_, (len((d_120_v4_).set(default__.safeIndex(d_139_v19_.f8, len(d_120_v4_)), _dafny.CodePoint('c')))) != (len(((d_145_v25_)[d_139_v19_.f8] if (d_139_v19_.f8) in (d_145_v25_) else d_144_v24_))), (d_146_v26_).cardinality, d_121_globalState_)
                        d_147_v27_: D2
                        d_147_v27_ = D2_DC5(d_126_v9_)
                        d_148_v28_: _dafny.Set
                        d_148_v28_ = _dafny.Set({d_141_v21_, not (d_141_v21_) or (False), (d_122_v5_) != (((d_115_v1_)[d_114_v0_] if (d_114_v0_) in (d_115_v1_) else len((d_147_v27_).cf9)))})
                        d_148_v28_ = d_148_v28_
                    default__.m0(d_122_v5_, d_132_v14_, (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], d_122_v5_, d_121_globalState_)
                    pass
            pass
        d_149_v29_: _dafny.Array
        nw26_ = _dafny.Array(_dafny.CodePoint('D'), 22)
        d_149_v29_ = nw26_
        d_150_v30_: _dafny.Set
        d_150_v30_ = _dafny.Set({d_149_v29_, d_149_v29_, d_149_v29_, d_149_v29_, d_149_v29_})
        d_151_v31_: _dafny.MultiSet
        d_151_v31_ = _dafny.MultiSet([d_122_v5_])
        d_152_v32_: D0
        d_152_v32_ = D0_DC3(d_115_v1_, d_114_v0_, d_122_v5_)
        default__.m0(len(d_150_v30_), d_151_v31_, (d_152_v32_).cf6, d_122_v5_, d_121_globalState_)
        d_153_v33_: _dafny.Map
        d_153_v33_ = _dafny.Map({d_122_v5_: d_116_v2_})
        d_153_v33_ = (d_153_v33_).set(d_122_v5_, d_116_v2_)
        d_154_v34_: D0
        d_154_v34_ = D0_DC1(d_122_v5_, (d_122_v5_) - (d_122_v5_), d_122_v5_)
        d_154_v34_ = D0_DC1(322, d_122_v5_, d_122_v5_)
        d_155_v35_: D0
        d_155_v35_ = D0_DC2(len(_dafny.SeqWithoutIsStrInference([d_122_v5_ for d_156_i4_ in range(default__.abs(-986))])))
        d_157_v36_: _dafny.MultiSet
        d_157_v36_ = _dafny.MultiSet([d_155_v35_, D0_DC2(d_122_v5_)])
        d_157_v36_ = (d_157_v36_).set(d_155_v35_, default__.abs(d_122_v5_))
        d_158_v37_: _dafny.Map
        d_158_v37_ = _dafny.Map({d_114_v0_: d_114_v0_})
        rhs14_ = (len((d_158_v37_) | (d_158_v37_)) if False else d_122_v5_)
        rhs15_ = d_122_v5_
        lhs14_ = d_121_globalState_
        d_122_v5_ = rhs14_
        lhs14_.f7 = rhs15_
        d_159_v38_: _dafny.Map
        d_159_v38_ = _dafny.Map({d_116_v2_: (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]})
        d_160_v39_: str
        d_160_v39_ = _dafny.CodePoint('c')
        if ((d_159_v38_)[d_116_v2_] if (d_116_v2_) in (d_159_v38_) else default__.fm0(not(d_114_v0_), default__.fm0(False, False, _dafny.MultiSet([(d_123_v6_).set(d_160_v39_, d_114_v0_)]), d_122_v5_, d_121_globalState_), (_dafny.MultiSet(d_124_v7_)).set(d_123_v6_, default__.abs(d_122_v5_)), d_122_v5_, d_121_globalState_)):
            d_161_v40_: _dafny.Map
            d_161_v40_ = _dafny.Map({d_114_v0_: d_118_v3_})
            d_161_v40_ = (d_161_v40_).set(d_114_v0_, d_118_v3_)
            d_162_v41_: D3
            d_162_v41_ = D3_DC7(d_158_v37_)
            d_114_v0_ = (((d_162_v41_).cf11) | (d_158_v37_)) == ((d_158_v37_) | (default__.fm4(not(d_114_v0_), d_122_v5_, (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], d_121_globalState_)))
            d_163_v42_: _dafny.MultiSet
            d_163_v42_ = _dafny.MultiSet([(d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], d_114_v0_, not (not(not(d_114_v0_))) or ((d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]), d_114_v0_, (d_120_v4_) < (d_120_v4_)])
            d_164_v43_: _dafny.Set
            d_164_v43_ = _dafny.Set({not(d_114_v0_)})
            d_165_v44_: _dafny.Map
            d_165_v44_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_122_v5_ for d_166_i5_ in range(default__.abs(604))])): not(d_114_v0_)})
            (d_121_globalState_).f7 = ((d_163_v42_)[default__.fm0(d_114_v0_, d_114_v0_, d_125_v8_, len(d_164_v43_), d_121_globalState_)] if (default__.fm0(d_114_v0_, d_114_v0_, d_125_v8_, len(d_164_v43_), d_121_globalState_)) in (d_163_v42_) else default__.safeDivisionInt((default__.fm5(17, d_114_v0_, ((d_165_v44_)[(d_151_v31_).cardinality] if ((d_151_v31_).cardinality) in (d_165_v44_) else True), d_121_globalState_)).cardinality, 136))
            d_167_v46_: _dafny.Map
            d_167_v46_ = _dafny.Map({d_122_v5_: D0_DC0(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwbs"))) for d_168_i7_ in range(default__.abs(242))])]))})
            d_169_v47_: _dafny.Seq
            d_169_v47_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_122_v5_: (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]}))])
            d_170_v48_: _dafny.Seq
            d_170_v48_ = _dafny.SeqWithoutIsStrInference([d_169_v47_])
            d_171_v49_: D0
            d_171_v49_ = D0_DC0(d_170_v48_)
            def iife21_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(954, 623):
                    d_172_v45_: int = compr_5_
                    if ((954) <= (d_172_v45_)) and ((d_172_v45_) < (623)):
                        coll5_[(d_172_v45_) + (len(_dafny.SeqWithoutIsStrInference([d_122_v5_])))] = D0_DC0(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_122_v5_ for d_173_i6_ in range(default__.abs(-263))]), _dafny.SeqWithoutIsStrInference([d_122_v5_, len(d_120_v4_)])]))
                return _dafny.Map(coll5_)
            default__.m0(len((iife21_()
            ) | ((d_167_v46_).set(len(d_120_v4_), d_171_v49_))), d_151_v31_, d_114_v0_, d_122_v5_, d_121_globalState_)
            d_122_v5_ = d_122_v5_
        elif True:
            d_174_v50_: C0
            nw27_ = C0()
            nw27_.ctor__(default__.safeModuloInt(d_122_v5_, -119))
            d_174_v50_ = nw27_
            d_175_v51_: _dafny.Set
            d_175_v51_ = _dafny.Set({d_158_v37_})
            d_176_v52_: _dafny.Set
            d_176_v52_ = d_175_v51_
            d_177_v53_: _dafny.Map
            d_177_v53_ = _dafny.Map({-303: d_176_v52_})
            d_178_v54_: _dafny.Seq
            d_178_v54_ = _dafny.SeqWithoutIsStrInference([d_122_v5_, 801, d_174_v50_.f8, d_174_v50_.f8, 333])
            d_177_v53_ = (d_177_v53_).set(len(_dafny.Map({d_178_v54_: d_120_v4_})), d_176_v52_)
            (d_121_globalState_).f7 = (d_154_v34_).cf2
            (d_121_globalState_).f7 = d_174_v50_.f8
            d_179_v55_: _dafny.Array
            nw28_ = _dafny.Array(_dafny.MultiSet({}), 22)
            d_179_v55_ = nw28_
            d_180_v56_: _dafny.Map
            d_180_v56_ = _dafny.Map({d_174_v50_: d_179_v55_})
            rhs16_ = d_155_v35_
            rhs17_ = d_152_v32_
            rhs18_ = d_174_v50_.f8
            rhs19_ = (((d_180_v56_)[d_174_v50_] if (d_174_v50_) in (d_180_v56_) else d_179_v55_) if d_114_v0_ else (d_179_v55_ if d_114_v0_ else d_179_v55_))
            lhs15_ = d_174_v50_
            d_155_v35_ = rhs16_
            d_152_v32_ = rhs17_
            lhs15_.f8 = rhs18_
            d_179_v55_ = rhs19_
        d_122_v5_ = d_122_v5_
        d_181_v57_: C0
        nw29_ = C0()
        nw29_.ctor__(530)
        d_181_v57_ = nw29_
        d_116_v2_ = d_116_v2_
        hi1_ = d_122_v5_
        for d_182_i8_ in range(d_181_v57_.f8, hi1_):
            d_183_v58_: _dafny.Map
            d_183_v58_ = _dafny.Map({d_181_v57_: False})
            d_183_v58_ = (d_183_v58_).set(d_181_v57_, not((d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]))
            d_184_v59_: _dafny.Array
            def lambda10_(d_185_v4_, d_186_v5_, d_187_v39_):
                def lambda11_(d_188_i9_):
                    return (_dafny.SeqWithoutIsStrInference([(d_185_v4_).set(default__.safeIndex(d_186_v5_, len(d_185_v4_)), d_187_v39_)])) + (_dafny.SeqWithoutIsStrInference([d_185_v4_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkc"))]))

                return lambda11_

            init7_ = lambda10_(d_120_v4_, d_122_v5_, d_160_v39_)
            nw30_ = _dafny.Array(None, 13)
            for i0_7_ in range(nw30_.length(0)):
                nw30_[i0_7_] = init7_(i0_7_)
            d_184_v59_ = nw30_
            d_189_v60_: _dafny.Seq
            d_189_v60_ = _dafny.SeqWithoutIsStrInference([d_120_v4_, (default__.fm6(d_182_i8_, d_181_v57_.f8, d_114_v0_, _dafny.CodePoint('f'), d_121_globalState_)).cf12, d_120_v4_, d_120_v4_])
            index29_ = default__.safeIndex(516, (d_184_v59_).length(0))
            (d_184_v59_)[index29_] = (d_189_v60_ if (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))] else d_189_v60_)
            index30_ = default__.safeIndex(516, (d_184_v59_).length(0))
            (d_184_v59_)[index30_] = d_189_v60_
            index31_ = default__.safeIndex(619, (d_118_v3_).length(0))
            (d_118_v3_)[index31_] = d_114_v0_
            d_190_v61_: _dafny.Map
            d_190_v61_ = _dafny.Map({_dafny.Map({d_116_v2_: d_160_v39_}): d_114_v0_})
            d_191_v62_: _dafny.Map
            d_191_v62_ = _dafny.Map({d_116_v2_: d_160_v39_})
            d_114_v0_ = not ((d_126_v9_)[default__.safeIndex(d_181_v57_.f8, len(d_126_v9_))]) or (((d_190_v61_)[d_191_v62_] if (d_191_v62_) in (d_190_v61_) else (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]))
        hi2_ = d_181_v57_.f8
        for d_192_i10_ in range(d_181_v57_.f8, hi2_):
            d_193_v63_: _dafny.Seq
            d_193_v63_ = _dafny.SeqWithoutIsStrInference([d_158_v37_])
            d_194_v64_: _dafny.Map
            d_194_v64_ = _dafny.Map({len(d_193_v63_): d_114_v0_})
            d_195_v65_: _dafny.Map
            d_195_v65_ = _dafny.Map({d_181_v57_.f8: d_122_v5_})
            d_194_v64_ = (d_194_v64_).set(default__.safeDivisionInt(len(d_195_v65_), d_192_i10_), d_114_v0_)
            index32_ = default__.safeIndex(619, (d_118_v3_).length(0))
            (d_118_v3_)[index32_] = default__.fm0((d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], d_125_v8_, d_181_v57_.f8, d_121_globalState_)
            index33_ = default__.safeIndex(619, (d_118_v3_).length(0))
            (d_118_v3_)[index33_] = (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]
            d_155_v35_ = d_155_v35_
        d_196_v66_: _dafny.Array
        nw31_ = _dafny.Array(_dafny.Seq({}), 19)
        d_196_v66_ = nw31_
        d_197_v67_: _dafny.Seq
        d_197_v67_ = _dafny.SeqWithoutIsStrInference([d_122_v5_])
        index34_ = default__.safeIndex(610, (d_196_v66_).length(0))
        (d_196_v66_)[index34_] = d_197_v67_
        index35_ = default__.safeIndex(610, (d_196_v66_).length(0))
        (d_196_v66_)[index35_] = (_dafny.SeqWithoutIsStrInference([d_181_v57_.f8 for d_198_i11_ in range(default__.abs(466))])) + (d_197_v67_)
        hi3_ = (d_122_v5_) + (-16)
        for d_199_i12_ in range((0) - (d_122_v5_), hi3_):
            if not((d_160_v39_) not in (d_120_v4_)):
                d_200_v68_: _dafny.Map
                d_200_v68_ = _dafny.Map({d_114_v0_: d_181_v57_})
                d_201_v69_: _dafny.Map
                d_201_v69_ = _dafny.Map({(d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]: d_200_v68_})
                d_201_v69_ = (d_201_v69_).set(False, (d_200_v68_).set((d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], d_181_v57_))
                d_202_v70_: _dafny.MultiSet
                d_202_v70_ = _dafny.MultiSet([d_114_v0_, (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], d_114_v0_])
                index36_ = default__.safeIndex(619, (d_118_v3_).length(0))
                rhs20_ = not((d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))])
                rhs21_ = default__.safeModuloInt(454, d_181_v57_.f8)
                rhs22_ = -651
                rhs23_ = ((d_181_v57_.f8) - (d_122_v5_)) - (d_181_v57_.f8)
                rhs24_ = (((d_158_v37_)[d_114_v0_] if (d_114_v0_) in (d_158_v37_) else default__.fm0(d_114_v0_, (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], d_125_v8_, len(_dafny.SeqWithoutIsStrInference([d_160_v39_ for d_203_i13_ in range(default__.abs(109))])), d_121_globalState_)) if True else (d_202_v70_).issubset(d_202_v70_))
                lhs16_ = d_181_v57_
                lhs17_ = d_181_v57_
                lhs18_ = d_181_v57_
                lhs19_ = d_118_v3_
                lhs20_ = default__.safeIndex(619, (d_118_v3_).length(0))
                d_114_v0_ = rhs20_
                lhs16_.f8 = rhs21_
                lhs17_.f8 = rhs22_
                lhs18_.f8 = rhs23_
                lhs19_[lhs20_] = rhs24_
                index37_ = default__.safeIndex(619, (d_118_v3_).length(0))
                (d_118_v3_)[index37_] = (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]
                d_204_v71_: D2
                d_204_v71_ = D2_DC5(_dafny.SeqWithoutIsStrInference([(d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]]))
                d_205_v72_: _dafny.Map
                d_205_v72_ = _dafny.Map({d_160_v39_: d_204_v71_})
                d_205_v72_ = (d_205_v72_).set(d_160_v39_, d_204_v71_)
                d_122_v5_ = -669
            elif True:
                d_206_v74_: D2
                d_206_v74_ = D2_DC5(d_126_v9_)
                d_207_v75_: _dafny.Map
                d_207_v75_ = _dafny.Map({d_120_v4_: (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]})
                d_208_v76_: D4
                d_208_v76_ = D4_DC9(d_207_v75_)
                d_209_v78_: _dafny.Set
                d_209_v78_ = _dafny.Set({d_120_v4_})
                d_210_v79_: _dafny.Seq
                def iife22_():
                    coll6_ = _dafny.Map()
                    compr_6_: _dafny.Seq
                    for compr_6_ in (d_209_v78_).Elements:
                        d_211_v77_: _dafny.Seq = compr_6_
                        if (d_211_v77_) in (d_209_v78_):
                            coll6_[d_211_v77_] = (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]
                    return _dafny.Map(coll6_)
                d_210_v79_ = _dafny.SeqWithoutIsStrInference([iife22_()
                , d_207_v75_])
                d_212_v80_: _dafny.Array
                nw32_ = _dafny.Array(None, 17)
                def iife23_():
                    coll7_ = _dafny.Map()
                    compr_7_: _dafny.Seq
                    for compr_7_ in (default__.fm7(d_114_v0_, d_206_v74_, (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], d_121_globalState_)).keys.Elements:
                        d_213_v73_: _dafny.Seq = compr_7_
                        if (d_213_v73_) in (default__.fm7(d_114_v0_, d_206_v74_, (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], d_121_globalState_)):
                            coll7_[d_213_v73_] = (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]
                    return _dafny.Map(coll7_)
                nw32_[int(0)] = (iife23_()
                ) | (d_207_v75_)
                nw32_[int(1)] = d_207_v75_
                nw32_[int(2)] = _dafny.Map({d_120_v4_: False})
                nw32_[int(3)] = d_207_v75_
                nw32_[int(4)] = d_207_v75_
                nw32_[int(5)] = d_207_v75_
                nw32_[int(6)] = d_207_v75_
                nw32_[int(7)] = d_207_v75_
                nw32_[int(8)] = d_207_v75_
                nw32_[int(9)] = (d_207_v75_) | (d_207_v75_)
                nw32_[int(10)] = d_207_v75_
                nw32_[int(11)] = (d_207_v75_) | (d_207_v75_)
                nw32_[int(12)] = (d_208_v76_).cf14
                nw32_[int(13)] = d_207_v75_
                nw32_[int(14)] = (d_210_v79_)[default__.safeIndex(d_122_v5_, len(d_210_v79_))]
                nw32_[int(15)] = d_207_v75_
                nw32_[int(16)] = d_207_v75_
                d_212_v80_ = nw32_
                index38_ = default__.safeIndex(35, (d_212_v80_).length(0))
                (d_212_v80_)[index38_] = default__.fm8(default__.fm0(False, False, _dafny.MultiSet([_dafny.Map({d_160_v39_: d_114_v0_}), d_123_v6_]), 10, d_121_globalState_), d_121_globalState_)
                index39_ = default__.safeIndex(35, (d_212_v80_).length(0))
                (d_212_v80_)[index39_] = (d_207_v75_) | (d_207_v75_)
                d_120_v4_ = (d_120_v4_) + (d_120_v4_)
                d_214_v81_: D3
                d_214_v81_ = D3_DC7(d_158_v37_)
                d_215_v82_: D3
                d_215_v82_ = D3_DC8(d_120_v4_, d_197_v67_)
                d_216_v83_: _dafny.MultiSet
                d_216_v83_ = _dafny.MultiSet([d_215_v82_, d_215_v82_])
                d_214_v81_ = default__.fm9(d_181_v57_.f8, (d_216_v83_).cardinality, d_121_globalState_)
                (d_121_globalState_).f7 = (0) - (d_181_v57_.f8)
                (d_181_v57_).f8 = d_181_v57_.f8
            (d_181_v57_).f8 = d_181_v57_.f8
            d_217_v84_: _dafny.Map
            d_217_v84_ = _dafny.Map({d_181_v57_: not (d_114_v0_) or (d_114_v0_)})
            d_217_v84_ = (d_217_v84_).set(d_181_v57_, d_114_v0_)
            d_218_v85_: _dafny.Array
            def lambda12_(d_219_v37_):
                def lambda13_(d_220_i14_):
                    return _dafny.Set({d_219_v37_})

                return lambda13_

            init8_ = lambda12_(d_158_v37_)
            nw33_ = _dafny.Array(None, 20)
            for i0_8_ in range(nw33_.length(0)):
                nw33_[i0_8_] = init8_(i0_8_)
            d_218_v85_ = nw33_
            d_221_v86_: D5
            d_221_v86_ = D5_DC11(d_160_v39_)
            index40_ = default__.safeIndex(436, (d_149_v29_).length(0))
            (d_149_v29_)[index40_] = (d_221_v86_).cf17
            d_222_v87_: _dafny.MultiSet
            d_222_v87_ = _dafny.MultiSet([default__.fm0(d_114_v0_, (d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))], d_125_v8_, len(d_197_v67_), d_121_globalState_)])
            index41_ = default__.safeIndex(436, (d_149_v29_).length(0))
            index42_ = default__.safeIndex(619, (d_118_v3_).length(0))
            rhs25_ = (d_122_v5_) * (((d_222_v87_)[(d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]] if ((d_118_v3_)[default__.safeIndex(619, (d_118_v3_).length(0))]) in (d_222_v87_) else d_199_i12_))
            rhs26_ = d_218_v85_
            rhs27_ = _dafny.CodePoint('r')
            rhs28_ = d_181_v57_.f8
            rhs29_ = True
            lhs21_ = d_181_v57_
            lhs22_ = d_149_v29_
            lhs23_ = default__.safeIndex(436, (d_149_v29_).length(0))
            lhs24_ = d_121_globalState_
            lhs25_ = d_118_v3_
            lhs26_ = default__.safeIndex(619, (d_118_v3_).length(0))
            lhs21_.f8 = rhs25_
            d_218_v85_ = rhs26_
            lhs22_[lhs23_] = rhs27_
            lhs24_.f7 = rhs28_
            lhs25_[lhs26_] = rhs29_
        _dafny.print(_dafny.string_of(d_114_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_115_v1_) == (_dafny.Map({False: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_v2_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v3_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v3_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v3_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v3_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v3_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v3_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v3_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v3_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v3_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v3_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v3_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v3_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_120_v4_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f0) == (_dafny.Map({False: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f1)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f2)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f2)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f2)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f2)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f2)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f2)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f2)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f2)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f2)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f2)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f2)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_globalState_).f2)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_121_globalState_.f4).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_121_globalState_).f6).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_121_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_122_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v6_) == (_dafny.Map({_dafny.CodePoint('i'): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v7_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('i'): False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v8_) == (_dafny.MultiSet([_dafny.Map({_dafny.CodePoint('i'): False}), _dafny.Map({_dafny.CodePoint('i'): False}), _dafny.Map({_dafny.CodePoint('i'): False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_) == (_dafny.SeqWithoutIsStrInference([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_127_v10_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_128_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v29_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_150_v30_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v31_) == (_dafny.MultiSet([681]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_152_v32_).cf5) == (_dafny.Map({False: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v32_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v32_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_153_v33_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v34_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v34_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v34_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v35_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v36_) == (_dafny.MultiSet([D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(986), D0_DC2(681)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v37_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_159_v38_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_160_v39_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_181_v57_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_196_v66_)[2]) == (_dafny.SeqWithoutIsStrInference([530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 530, 681]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_197_v67_) == (_dafny.SeqWithoutIsStrInference([681]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), int(0), int(0))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf5', Any), ('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
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
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC4(D1, NamedTuple('DC4', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC8(D3, NamedTuple('DC8', [('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({self.cf12.VerbatimString(True)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10(False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)

class D4_DC10(D4, NamedTuple('DC10', [('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC12(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC12(D5, NamedTuple('DC12', [('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf21', Any), ('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC11(D5, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf17 == __o.cf17
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
        return lambda: D6_DC17(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC17(D6, NamedTuple('DC17', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19)
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f4: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f7: int = int(0)
        self._f0: _dafny.Map = _dafny.Map({})
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f2: _dafny.Array = _dafny.Array(None, 0)
        self._f3: int = int(0)
        self._f5: bool = False
        self._f6: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7

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
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6

class C0:
    def  __init__(self):
        self.f8: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f8):
        (self).f8 = f8

    def m1(self, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: str = _dafny.CodePoint('D')
        d_223_v0_: _dafny.Seq
        d_223_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdjxni"))
        d_224_v1_: bool
        d_224_v1_ = True
        d_225_v2_: _dafny.Map
        d_225_v2_ = _dafny.Map({d_224_v1_: self.f8})
        d_226_v3_: _dafny.Array
        nw34_ = _dafny.Array(int(0), 17)
        d_226_v3_ = nw34_
        d_227_v4_: _dafny.Map
        d_227_v4_ = _dafny.Map({d_226_v3_: self.f8})
        d_228_v5_: _dafny.Seq
        d_228_v5_ = _dafny.SeqWithoutIsStrInference([len(d_225_v2_), len(d_227_v4_)])
        d_229_v6_: _dafny.Array
        nw35_ = _dafny.Array(None, 4)
        nw35_[int(0)] = (d_223_v0_) < (d_223_v0_)
        nw35_[int(1)] = (_dafny.SeqWithoutIsStrInference([self.f8 for d_230_i0_ in range(default__.abs(117))])) not in ((D0_DC0(_dafny.SeqWithoutIsStrInference([d_228_v5_]))).cf0)
        nw35_[int(2)] = d_224_v1_
        nw35_[int(3)] = d_224_v1_
        d_229_v6_ = nw35_
        d_231_v7_: _dafny.Map
        d_231_v7_ = _dafny.Map({not(d_224_v1_): d_224_v1_})
        d_232_v8_: _dafny.Set
        d_232_v8_ = _dafny.Set({d_231_v7_, ((d_231_v7_).set(d_224_v1_, False)).set(d_224_v1_, True), d_231_v7_})
        d_233_v9_: str
        d_233_v9_ = _dafny.CodePoint('f')
        d_234_v11_: _dafny.Set
        d_234_v11_ = _dafny.Set({_dafny.Map({True: d_224_v1_}), d_231_v7_})
        def iife24_():
            coll8_ = _dafny.Set()
            compr_8_: _dafny.Map
            for compr_8_ in (_dafny.Map({_dafny.Map({d_233_v9_: not(d_224_v1_)}): d_224_v1_})).keys.Elements:
                d_235_v10_: _dafny.Map = compr_8_
                if (d_235_v10_) in (_dafny.Map({_dafny.Map({d_233_v9_: not(d_224_v1_)}): d_224_v1_})):
                    coll8_ = coll8_.union(_dafny.Set([d_235_v10_]))
            return _dafny.Set(coll8_)
        rhs30_ = len(iife24_()
        )
        rhs31_ = d_229_v6_
        rhs32_ = default__.safeDivisionInt((self.f8) * (self.f8), (0) - (len(_dafny.SeqWithoutIsStrInference([d_224_v1_]))))
        rhs33_ = (self.f8) + (self.f8)
        rhs34_ = (d_234_v11_)
        lhs27_ = globalState
        lhs28_ = globalState
        lhs29_ = globalState
        lhs27_.f7 = rhs30_
        d_229_v6_ = rhs31_
        lhs28_.f7 = rhs32_
        lhs29_.f7 = rhs33_
        d_232_v8_ = rhs34_
        d_236_v12_: _dafny.Seq
        d_236_v12_ = _dafny.SeqWithoutIsStrInference([d_224_v1_])
        d_237_v13_: _dafny.Seq
        d_237_v13_ = _dafny.SeqWithoutIsStrInference([d_236_v12_])
        d_236_v12_ = (d_236_v12_) + ((d_237_v13_)[default__.safeIndex(self.f8, len(d_237_v13_))])
        (globalState).f7 = self.f8
        (globalState).f7 = ((0) - (self.f8)) * (616)
        (self).f8 = self.f8
        d_238_v14_: _dafny.MultiSet
        d_238_v14_ = _dafny.MultiSet([self.f8])
        default__.m0(self.f8, d_238_v14_, d_224_v1_, self.f8, globalState)
        d_239_v15_: _dafny.Map
        d_239_v15_ = _dafny.Map({d_233_v9_: d_224_v1_})
        d_240_v16_: _dafny.MultiSet
        d_240_v16_ = _dafny.MultiSet([d_239_v15_])
        def iife25_():
            coll9_ = _dafny.Map()
            compr_9_: str
            for compr_9_ in (_dafny.SeqWithoutIsStrInference([d_233_v9_ for d_241_i1_ in range(default__.abs(605))])).Elements:
                d_242_v17_: str = compr_9_
                if (d_242_v17_) in (_dafny.SeqWithoutIsStrInference([d_233_v9_ for d_241_i1_ in range(default__.abs(605))])):
                    coll9_[d_242_v17_] = not(d_224_v1_)
            return _dafny.Map(coll9_)
        r0 = default__.fm0(default__.fm0(True, d_224_v1_, d_240_v16_, len(d_232_v8_), globalState), d_224_v1_, ((d_240_v16_).set(iife25_()
        , default__.abs(self.f8))) | (d_240_v16_), self.f8, globalState)
        r1 = d_224_v1_
        d_243_v18_: _dafny.Set
        d_243_v18_ = _dafny.Set({d_224_v1_, True})
        r2 = default__.safeDivisionInt(len(d_243_v18_), ((0) - (self.f8)) - (self.f8))
        d_244_v19_: D0
        d_244_v19_ = D0_DC3((_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([not(d_224_v1_), False, d_224_v1_, not(d_224_v1_)]))})).set(d_224_v1_, self.f8), d_224_v1_, self.f8)
        pat_let_tv3_ = d_233_v9_
        pat_let_tv4_ = d_233_v9_
        pat_let_tv5_ = d_223_v0_
        pat_let_tv6_ = d_223_v0_
        def lambda14_(source2_):
            if source2_.is_DC1:
                d_245___mcc_h0_ = source2_.cf1
                d_246___mcc_h1_ = source2_.cf2
                d_247___mcc_h2_ = source2_.cf3
                d_248_cf3_ = d_247___mcc_h2_
                d_249_cf2_ = d_246___mcc_h1_
                d_250_cf1_ = d_245___mcc_h0_
                return pat_let_tv3_
            elif source2_.is_DC2:
                d_251___mcc_h3_ = source2_.cf4
                d_252_cf4_ = d_251___mcc_h3_
                return _dafny.CodePoint('d')
            elif source2_.is_DC3:
                d_253___mcc_h4_ = source2_.cf5
                d_254___mcc_h5_ = source2_.cf6
                d_255___mcc_h6_ = source2_.cf7
                d_256_cf7_ = d_255___mcc_h6_
                d_257_cf6_ = d_254___mcc_h5_
                d_258_cf5_ = d_253___mcc_h4_
                return pat_let_tv4_
            elif True:
                d_259___mcc_h7_ = source2_.cf0
                d_260_cf0_ = d_259___mcc_h7_
                return (pat_let_tv5_)[default__.safeIndex(self.f8, len(pat_let_tv6_))]

        r3 = lambda14_(d_244_v19_)
        return r0, r1, r2, r3

