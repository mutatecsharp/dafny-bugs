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
        return ((len(_dafny.Set({D0_DC1(), D0_DC1()}))) == (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uakfefor"))))) or (False)

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        return ((D10_DC28(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "syud"))), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfngghe"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_0_i0_ in range(default__.abs(-859))])), 313])): 420}))]))).cf38).cardinality

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return D0_DC1()

    @staticmethod
    def fm3(globalState):
        source0_ = D1_DC3(False, not(False), not(False))
        if source0_.is_DC3:
            d_1___mcc_h0_ = source0_.cf2
            d_2___mcc_h1_ = source0_.cf3
            d_3___mcc_h2_ = source0_.cf4
            d_4_cf4_ = d_3___mcc_h2_
            d_5_cf3_ = d_2___mcc_h1_
            d_6_cf2_ = d_1___mcc_h0_
            return _dafny.CodePoint('y')
        elif True:
            d_7___mcc_h3_ = source0_.cf1
            d_8_cf1_ = d_7___mcc_h3_
            return d_8_cf1_

    @staticmethod
    def fm4(p0, p1, globalState):
        return _dafny.Map({not((False if not(True) else False)): _dafny.CodePoint('t')})

    @staticmethod
    def fm5(p0, p1, p2, p3, globalState):
        return (_dafny.Set({not(True), False})).intersection(_dafny.Set({True, False}))

    @staticmethod
    def fm12(globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sua"))

    @staticmethod
    def fm13(globalState):
        return D2_DC5(not((True) and (not(not(True)))), 989, 149)

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        source1_ = D7_DC21(_dafny.MultiSet([True]), (0) - (len(_dafny.Set({False}))), True)
        if source1_.is_DC21:
            d_9___mcc_h0_ = source1_.cf30
            d_10___mcc_h1_ = source1_.cf31
            d_11___mcc_h2_ = source1_.cf32
            d_12_cf32_ = d_11___mcc_h2_
            d_13_cf31_ = d_10___mcc_h1_
            d_14_cf30_ = d_9___mcc_h0_
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_15_i0_ in range(default__.abs(935))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgxrugvjl"))]))) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpq"))]))
        elif source1_.is_DC22:
            return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyaxwe")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwwoqsahs")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_16_i1_ in range(default__.abs(-594))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))])
        elif source1_.is_DC23:
            d_17___mcc_h3_ = source1_.cf33
            d_18_cf33_ = d_17___mcc_h3_
            return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybamd"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qndkpncb"))])))
        elif True:
            d_19___mcc_h4_ = source1_.cf29
            d_20_cf29_ = d_19___mcc_h4_
            return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sca"))])

    @staticmethod
    def fm15(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + ((_dafny.SeqWithoutIsStrInference([True, False, False])) + (_dafny.SeqWithoutIsStrInference([True, True])))

    @staticmethod
    def fm16(p0, globalState):
        return (_dafny.MultiSet([(0) - (len(_dafny.Map({False: False})))])).intersection(_dafny.MultiSet([698]))

    @staticmethod
    def fm19(p0, p1, globalState):
        return (_dafny.Map({False: not(False)})) | ((_dafny.Map({not(False): True})) | (_dafny.Map({True: True})))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        return (((D11_DC31(_dafny.Map({len(_dafny.Set({794})): True}))).cf43 if True else _dafny.Map({922: False}))) | (_dafny.Map({-337: True}))

    @staticmethod
    def fm21(p0, p1, globalState):
        return ((_dafny.Map({_dafny.MultiSet([260]): _dafny.Set({not(not(not(True)))})})) | (_dafny.Map({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False])), -886]): _dafny.Set({False, True, True})}))) | ((_dafny.Map({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([272])), 332, -582, 968, 104]): _dafny.Set({True})})) | (_dafny.Map({_dafny.MultiSet([74]): _dafny.Set({False})})))

    @staticmethod
    def fm22(p0, p1, globalState):
        if (_dafny.Set({True})).isdisjoint(_dafny.Set({not(not(True)), True})):
            def iife0_():
                coll0_ = _dafny.Map()
                compr_0_: int
                for compr_0_ in (_dafny.MultiSet([-156, (0) - (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len((D12_DC34(_dafny.Map({199: _dafny.CodePoint('f')}))).cf46) for d_21_i0_ in range(default__.abs(287))]))}))), -761])).Elements:
                    d_22_v0_: int = compr_0_
                    if (d_22_v0_) in (_dafny.MultiSet([-156, (0) - (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len((D12_DC34(_dafny.Map({199: _dafny.CodePoint('f')}))).cf46) for d_21_i0_ in range(default__.abs(287))]))}))), -761])):
                        coll0_[(d_22_v0_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = 950
                return _dafny.Map(coll0_)
            return iife0_()
            
        elif True:
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in (_dafny.SeqWithoutIsStrInference([940 for d_24_i2_ in range(default__.abs(34))])).Elements:
                    d_25_v1_: int = compr_1_
                    if (d_25_v1_) in (_dafny.SeqWithoutIsStrInference([940 for d_24_i2_ in range(default__.abs(34))])):
                        coll1_[(d_25_v1_) - (678)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_26_i3_ in range(default__.abs(-429))]))]))
                return _dafny.Map(coll1_)
            return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdhfk"))])): len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_23_i1_ in range(default__.abs(-889))])}))})) | (_dafny.Map({734: len(iife1_()
            )}))

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        return (_dafny.Map({not(False): _dafny.Set({True})})) | (_dafny.Map({False: _dafny.Set({False})}))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: bool = False
        d_27_v0_: _dafny.Map
        d_27_v0_ = _dafny.Map({p2: p3})
        d_27_v0_ = (d_27_v0_).set(p2, 776)
        r2 = p3
        d_28_v1_: _dafny.MultiSet
        d_28_v1_ = _dafny.MultiSet([p2, p2, p2, not(True), p2])
        d_29_v2_: _dafny.MultiSet
        d_29_v2_ = _dafny.MultiSet([(d_28_v1_).cardinality, p3, p3])
        d_30_v3_: C2
        nw0_ = C2()
        nw0_.ctor__(-881)
        d_30_v3_ = nw0_
        d_31_v4_: _dafny.Map
        d_31_v4_ = _dafny.Map({p2: d_30_v3_})
        d_32_v5_: _dafny.Seq
        d_32_v5_ = _dafny.SeqWithoutIsStrInference([p3])
        d_33_v6_: _dafny.Seq
        d_33_v6_ = _dafny.SeqWithoutIsStrInference([d_30_v3_])
        d_34_v7_: _dafny.Seq
        d_34_v7_ = _dafny.SeqWithoutIsStrInference([((d_29_v2_)[p3] if (p3) in (d_29_v2_) else (0) - (p3)), default__.safeModuloInt(len(d_31_v4_), p3), len(d_32_v5_), default__.safeModuloInt((d_30_v3_).f3, len(d_33_v6_))])
        d_32_v5_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aymdpyag"))) for d_35_i0_ in range(default__.abs(446))])
        d_36_v8_: _dafny.Seq
        d_36_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wpx"))
        d_36_v8_ = (d_36_v8_) + (d_36_v8_)
        d_37_v9_: D7
        d_37_v9_ = D7_DC21(d_28_v1_, p3, p2)
        pat_let_tv0_ = p2
        pat_let_tv1_ = p2
        pat_let_tv2_ = p2
        def lambda0_(source2_):
            if source2_.is_DC21:
                d_38___mcc_h0_ = source2_.cf30
                d_39___mcc_h1_ = source2_.cf31
                d_40___mcc_h2_ = source2_.cf32
                d_41_cf32_ = d_40___mcc_h2_
                d_42_cf31_ = d_39___mcc_h1_
                d_43_cf30_ = d_38___mcc_h0_
                return d_41_cf32_
            elif source2_.is_DC22:
                return pat_let_tv0_
            elif source2_.is_DC23:
                d_44___mcc_h3_ = source2_.cf33
                d_45_cf33_ = d_44___mcc_h3_
                return pat_let_tv1_
            elif True:
                d_46___mcc_h4_ = source2_.cf29
                d_47_cf29_ = d_46___mcc_h4_
                return pat_let_tv2_

        if lambda0_(d_37_v9_):
            d_48_v10_: _dafny.Array
            def lambda1_(d_49_p2_):
                def lambda2_(d_50_i1_):
                    return d_49_p2_

                return lambda2_

            init0_ = lambda1_(p2)
            nw1_ = _dafny.Array(None, 15)
            for i0_0_ in range(nw1_.length(0)):
                nw1_[i0_0_] = init0_(i0_0_)
            d_48_v10_ = nw1_
            d_48_v10_ = d_48_v10_
            index0_ = default__.safeIndex(842, (d_48_v10_).length(0))
            (d_48_v10_)[index0_] = True
            index1_ = default__.safeIndex(842, (d_48_v10_).length(0))
            rhs0_ = (0) - (default__.safeDivisionInt((d_30_v3_).f3, (0) - (p3)))
            rhs1_ = p2
            lhs0_ = d_48_v10_
            lhs1_ = default__.safeIndex(842, (d_48_v10_).length(0))
            r2 = rhs0_
            lhs0_[lhs1_] = rhs1_
            d_51_v11_: str
            d_51_v11_ = _dafny.CodePoint('n')
            d_52_v12_: _dafny.Set
            d_52_v12_ = _dafny.Set({p2})
            d_53_v13_: _dafny.Map
            d_53_v13_ = _dafny.Map({p2: d_52_v12_})
            r2 = len(((default__.fm23((d_30_v3_).f3, d_51_v11_, (d_48_v10_)[default__.safeIndex(842, (d_48_v10_).length(0))], (d_30_v3_).f3, globalState)) | (d_53_v13_)).set(p2, d_52_v12_))
            r2 = ((d_30_v3_).f3) + ((len(d_27_v0_) if (d_48_v10_)[default__.safeIndex(842, (d_48_v10_).length(0))] else (d_30_v3_).f3))
            r2 = (d_30_v3_).f3
        elif True:
            r2 = -257
            d_54_v14_: _dafny.Seq
            d_54_v14_ = _dafny.SeqWithoutIsStrInference([p2])
            r1 = (d_54_v14_)[default__.safeIndex(((d_30_v3_).f3) + ((d_30_v3_).f3), len(d_54_v14_))]
            d_55_v15_: _dafny.Array
            def lambda3_(d_56_p2_, d_57_v8_):
                def lambda4_(d_58_i2_):
                    return (D2_DC6(D2_DC4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mw")))) if d_56_p2_ else D2_DC6(D2_DC6(D2_DC6(D2_DC4(d_57_v8_)))))

                return lambda4_

            init1_ = lambda3_(p2, d_36_v8_)
            nw2_ = _dafny.Array(None, 5)
            for i0_1_ in range(nw2_.length(0)):
                nw2_[i0_1_] = init1_(i0_1_)
            d_55_v15_ = nw2_
            d_59_v16_: D2
            d_59_v16_ = D2_DC4(d_36_v8_)
            d_60_v17_: D2
            d_60_v17_ = D2_DC6(d_59_v16_)
            index2_ = default__.safeIndex(67, (d_55_v15_).length(0))
            (d_55_v15_)[index2_] = d_60_v17_
            index3_ = default__.safeIndex(67, (d_55_v15_).length(0))
            (d_55_v15_)[index3_] = (d_60_v17_ if p2 else D2_DC6(d_59_v16_))
            d_61_v18_: _dafny.Set
            d_61_v18_ = _dafny.Set({p2, p2, p2})
            d_62_v19_: _dafny.Map
            d_62_v19_ = _dafny.Map({(d_61_v18_) | (_dafny.Set({p2, p2, p2, True})): default__.fm12(globalState)})
            d_62_v19_ = _dafny.Map({d_61_v18_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_63_i3_ in range(default__.abs(741))])})
            d_64_v20_: _dafny.Array
            nw3_ = _dafny.Array(None, 2)
            nw3_[int(0)] = _dafny.SeqWithoutIsStrInference([p3])
            nw3_[int(1)] = _dafny.SeqWithoutIsStrInference([(0) - (len(d_36_v8_))])
            d_64_v20_ = nw3_
            d_65_v21_: C0
            nw4_ = C0()
            nw4_.ctor__(d_64_v20_)
            d_65_v21_ = nw4_
            d_66_v22_: _dafny.Map
            d_66_v22_ = _dafny.Map({d_30_v3_: (d_65_v21_ if True else d_65_v21_)})
            d_66_v22_ = (d_66_v22_).set(d_30_v3_, d_65_v21_)
        d_67_v23_: _dafny.Array
        nw5_ = _dafny.Array(_dafny.Seq({}), 14)
        d_67_v23_ = nw5_
        d_68_v24_: _dafny.Map
        d_68_v24_ = _dafny.Map({p2: True})
        index4_ = default__.safeIndex(453, (d_67_v23_).length(0))
        (d_67_v23_)[index4_] = _dafny.SeqWithoutIsStrInference([d_68_v24_, d_68_v24_, d_68_v24_, d_68_v24_, default__.fm19(p2, (d_30_v3_).fm18(D2_DC5(p2, (d_30_v3_).f3, (d_30_v3_).f3), globalState), globalState)])
        d_69_v25_: _dafny.Array
        def lambda5_(d_70_v3_):
            def lambda6_(d_71_i4_):
                return default__.safeDivisionInt(d_71_i4_, (d_70_v3_).f3)

            return lambda6_

        init2_ = lambda5_(d_30_v3_)
        nw6_ = _dafny.Array(None, 27)
        for i0_2_ in range(nw6_.length(0)):
            nw6_[i0_2_] = init2_(i0_2_)
        d_69_v25_ = nw6_
        index5_ = default__.safeIndex(753, (d_69_v25_).length(0))
        (d_69_v25_)[index5_] = p3
        d_72_v26_: _dafny.Seq
        d_72_v26_ = _dafny.SeqWithoutIsStrInference([(d_68_v24_) | (d_68_v24_), d_68_v24_, d_68_v24_])
        d_73_v27_: C1
        nw7_ = C1()
        nw7_.ctor__()
        d_73_v27_ = nw7_
        d_74_v28_: _dafny.MultiSet
        d_74_v28_ = _dafny.MultiSet([d_73_v27_])
        index6_ = default__.safeIndex(453, (d_67_v23_).length(0))
        index7_ = default__.safeIndex(753, (d_69_v25_).length(0))
        rhs2_ = d_72_v26_
        rhs3_ = len((d_68_v24_ if p2 else d_68_v24_))
        rhs4_ = ((d_74_v28_).ispropersubset(d_74_v28_)) or (p2)
        lhs2_ = d_67_v23_
        lhs3_ = default__.safeIndex(453, (d_67_v23_).length(0))
        lhs4_ = d_69_v25_
        lhs5_ = default__.safeIndex(753, (d_69_v25_).length(0))
        lhs2_[lhs3_] = rhs2_
        lhs4_[lhs5_] = rhs3_
        r1 = rhs4_
        d_75_v29_: _dafny.Seq
        d_75_v29_ = _dafny.SeqWithoutIsStrInference([p2])
        r0 = ((d_75_v29_)[default__.safeIndex(p3, len(d_75_v29_))]) and (p2)
        r1 = False
        r2 = (0) - ((d_30_v3_).f3)
        r3 = p2
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_76_v0_: bool
        d_76_v0_ = False
        d_77_v1_: int
        d_77_v1_ = -894
        d_78_globalState_: GlobalState
        nw8_ = GlobalState()
        nw8_.ctor__(_dafny.Map({d_76_v0_: d_77_v1_}), False)
        d_78_globalState_ = nw8_
        d_79_v2_: _dafny.Set
        d_79_v2_ = _dafny.Set({d_76_v0_})
        d_80_v3_: D0
        d_80_v3_ = D0_DC0(len(d_79_v2_))
        hi0_ = ((d_80_v3_).cf0) - (d_77_v1_)
        for d_81_i0_ in range(d_77_v1_, hi0_):
            d_82_v4_: _dafny.Array
            nw9_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_82_v4_ = nw9_
            d_83_v5_: _dafny.Array
            nw10_ = _dafny.Array(None, 1)
            nw10_[int(0)] = d_81_i0_
            d_83_v5_ = nw10_
            index8_ = default__.safeIndex(863, (d_82_v4_).length(0))
            (d_82_v4_)[index8_] = d_83_v5_
            index9_ = default__.safeIndex(863, (d_82_v4_).length(0))
            (d_82_v4_)[index9_] = d_83_v5_
            d_84_v6_: _dafny.Seq
            d_84_v6_ = _dafny.SeqWithoutIsStrInference([662])
            d_85_v7_: _dafny.Seq
            d_85_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            d_86_v8_: _dafny.MultiSet
            d_86_v8_ = _dafny.MultiSet([d_76_v0_])
            d_87_v9_: _dafny.Map
            d_87_v9_ = _dafny.Map({(d_81_i0_) * (((d_86_v8_)[d_76_v0_] if (d_76_v0_) in (d_86_v8_) else 545)): d_76_v0_})
            index10_ = default__.safeIndex(863, (d_82_v4_).length(0))
            rhs5_ = default__.fm0(default__.safeModuloInt(len(d_84_v6_), d_81_i0_), default__.fm1(len(d_79_v2_), d_81_i0_, len(d_85_v7_), False, d_78_globalState_), _dafny.CodePoint('j'), d_78_globalState_)
            rhs6_ = d_76_v0_
            rhs7_ = (0) - (len(d_87_v9_))
            rhs8_ = (d_82_v4_)[default__.safeIndex(863, (d_82_v4_).length(0))]
            lhs6_ = d_82_v4_
            lhs7_ = default__.safeIndex(863, (d_82_v4_).length(0))
            d_76_v0_ = rhs5_
            d_76_v0_ = rhs6_
            d_77_v1_ = rhs7_
            lhs6_[lhs7_] = rhs8_
            arr0_ = (d_82_v4_)[default__.safeIndex(863, (d_82_v4_).length(0))]
            index11_ = default__.safeIndex(117, ((d_82_v4_)[default__.safeIndex(863, (d_82_v4_).length(0))]).length(0))
            arr0_[index11_] = (d_81_i0_) + ((0) - (d_81_i0_))
            d_88_v10_: _dafny.Seq
            d_88_v10_ = _dafny.SeqWithoutIsStrInference([d_76_v0_, True])
            arr1_ = (d_82_v4_)[default__.safeIndex(863, (d_82_v4_).length(0))]
            index12_ = default__.safeIndex(117, ((d_82_v4_)[default__.safeIndex(863, (d_82_v4_).length(0))]).length(0))
            arr1_[index12_] = default__.fm1(len(d_88_v10_), d_81_i0_, (d_81_i0_) * (d_81_i0_), (d_81_i0_) != (-356), d_78_globalState_)
            d_89_v11_: _dafny.Array
            def lambda7_(d_90_i1_):
                return True

            init3_ = lambda7_
            nw11_ = _dafny.Array(None, 2)
            for i0_3_ in range(nw11_.length(0)):
                nw11_[i0_3_] = init3_(i0_3_)
            d_89_v11_ = nw11_
            index13_ = default__.safeIndex(297, (d_89_v11_).length(0))
            (d_89_v11_)[index13_] = d_76_v0_
            index14_ = default__.safeIndex(297, (d_89_v11_).length(0))
            (d_89_v11_)[index14_] = ((d_86_v8_).set(d_76_v0_, default__.abs((0) - (len(_dafny.Set({d_76_v0_, d_76_v0_, d_76_v0_, not(d_76_v0_)})))))).issubset((d_86_v8_).intersection(d_86_v8_))
        d_91_v12_: _dafny.Array
        def lambda8_(d_92_i3_):
            return default__.safeModuloInt(d_92_i3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))))

        init4_ = lambda8_
        nw12_ = _dafny.Array(None, 2)
        for i0_4_ in range(nw12_.length(0)):
            nw12_[i0_4_] = init4_(i0_4_)
        d_91_v12_ = nw12_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_91_v12_).length(0)):
            d_93_i2_: int = guard_loop_0_
            if (True) and (((0) <= (d_93_i2_)) and ((d_93_i2_) < ((d_91_v12_).length(0)))):
                (d_91_v12_)[(d_93_i2_)] = default__.safeDivisionInt(d_93_i2_, d_77_v1_)
        d_94_v13_: _dafny.MultiSet
        d_94_v13_ = _dafny.MultiSet([d_77_v1_, d_77_v1_])
        d_95_v14_: D0
        d_95_v14_ = D0_DC1()
        source3_ = (default__.fm2(d_79_v2_, d_76_v0_, d_76_v0_, d_76_v0_, d_78_globalState_) if ((d_94_v13_).cardinality) >= (d_77_v1_) else d_95_v14_)
        if source3_.is_DC1:
            d_96_v15_: str
            d_96_v15_ = _dafny.CodePoint('l')
            d_97_v16_: _dafny.Map
            d_97_v16_ = _dafny.Map({d_76_v0_: d_96_v15_})
            d_98_v18_: _dafny.Map
            d_98_v18_ = _dafny.Map({d_94_v13_: d_96_v15_})
            d_99_v19_: bool
            d_100_v20_: bool
            d_101_v21_: int
            d_102_v22_: bool
            out0_: bool
            out1_: bool
            out2_: int
            out3_: bool
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: _dafny.MultiSet
                for compr_2_ in (d_98_v18_).keys.Elements:
                    d_103_v17_: _dafny.MultiSet = compr_2_
                    if (d_103_v17_) in (d_98_v18_):
                        coll2_[d_103_v17_] = _dafny.Set({d_76_v0_, d_76_v0_, d_76_v0_})
                return _dafny.Map(coll2_)
            out0_, out1_, out2_, out3_ = default__.m0(d_97_v16_, iife2_()
            , False, d_77_v1_, d_78_globalState_)
            d_99_v19_ = out0_
            d_100_v20_ = out1_
            d_101_v21_ = out2_
            d_102_v22_ = out3_
            d_104_v23_: _dafny.Map
            d_104_v23_ = _dafny.Map({d_96_v15_: d_76_v0_})
            d_105_v24_: _dafny.Map
            d_105_v24_ = _dafny.Map({d_102_v22_: d_101_v21_})
            d_106_v25_: _dafny.Map
            d_106_v25_ = _dafny.Map({d_77_v1_: d_77_v1_})
            d_107_v26_: _dafny.Seq
            d_107_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojlppbr"))
            d_108_v27_: _dafny.Map
            d_108_v27_ = _dafny.Map({d_107_v26_: d_100_v20_})
            d_109_v28_: _dafny.MultiSet
            d_109_v28_ = _dafny.MultiSet([True])
            rhs9_ = (d_101_v21_) - (151)
            rhs10_ = d_101_v21_
            rhs11_ = (_dafny.Map({default__.fm3(d_78_globalState_): d_99_v19_})) | ((d_104_v23_).set(d_96_v15_, d_76_v0_))
            rhs12_ = ((d_105_v24_).set(default__.fm0(((d_106_v25_)[d_101_v21_] if (d_101_v21_) in (d_106_v25_) else (_dafny.MultiSet([d_101_v21_])).cardinality), d_77_v1_, (D1_DC2(d_96_v15_)).cf1, d_78_globalState_), d_77_v1_)).set(((d_108_v27_)[d_107_v26_] if (d_107_v26_) in (d_108_v27_) else d_102_v22_), default__.fm1((d_109_v28_).cardinality, d_77_v1_, 2, d_100_v20_, d_78_globalState_))
            rhs13_ = d_100_v20_
            lhs8_ = d_78_globalState_
            d_101_v21_ = rhs9_
            d_101_v21_ = rhs10_
            d_104_v23_ = rhs11_
            lhs8_.f0 = rhs12_
            d_99_v19_ = rhs13_
            d_110_v29_: _dafny.Seq
            d_110_v29_ = _dafny.SeqWithoutIsStrInference([-209])
            d_111_v30_: _dafny.Seq
            d_111_v30_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - (-91), d_101_v21_]), (d_110_v29_) + (_dafny.SeqWithoutIsStrInference([333]))])
            d_101_v21_ = len(d_111_v30_)
            if not(d_102_v22_):
                d_77_v1_ = d_77_v1_
                d_112_v31_: _dafny.MultiSet
                d_112_v31_ = _dafny.MultiSet([d_107_v26_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))])
                d_113_v32_: _dafny.Seq
                d_113_v32_ = _dafny.SeqWithoutIsStrInference([d_76_v0_, not(d_102_v22_)])
                d_114_v33_: _dafny.Map
                d_114_v33_ = _dafny.Map({d_107_v26_: (d_106_v25_) | ((d_106_v25_).set(((d_112_v31_).set(d_107_v26_, default__.abs(len(d_113_v32_)))).cardinality, (d_110_v29_)[default__.safeIndex(d_77_v1_, len(d_110_v29_))]))})
                d_114_v33_ = (d_114_v33_).set(d_107_v26_, _dafny.Map({d_77_v1_: d_77_v1_}))
                d_96_v15_ = d_96_v15_
                d_115_v34_: _dafny.MultiSet
                d_115_v34_ = _dafny.MultiSet([d_110_v29_, _dafny.SeqWithoutIsStrInference([136 for d_116_i4_ in range(default__.abs(46))])])
                d_117_v35_: _dafny.Map
                d_117_v35_ = _dafny.Map({d_94_v13_: default__.fm5((0) - (-664), d_102_v22_, d_99_v19_, -128, d_78_globalState_)})
                d_118_v36_: bool
                d_119_v37_: bool
                d_120_v38_: int
                d_121_v39_: bool
                out4_: bool
                out5_: bool
                out6_: int
                out7_: bool
                out4_, out5_, out6_, out7_ = default__.m0(default__.fm4(default__.fm1(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "th"))), (d_115_v34_).cardinality, d_101_v21_, default__.fm0((d_94_v13_).cardinality, d_101_v21_, d_96_v15_, d_78_globalState_), d_78_globalState_), d_99_v19_, d_78_globalState_), d_117_v35_, (d_100_v20_) == (d_99_v19_), d_77_v1_, d_78_globalState_)
                d_118_v36_ = out4_
                d_119_v37_ = out5_
                d_120_v38_ = out6_
                d_121_v39_ = out7_
                d_77_v1_ = default__.safeModuloInt(d_101_v21_, -822)
            elif True:
                d_76_v0_ = not(default__.fm0(default__.safeDivisionInt(d_77_v1_, d_77_v1_), d_101_v21_, d_96_v15_, d_78_globalState_))
                d_122_v40_: _dafny.Array
                def lambda9_(d_123_v0_):
                    def lambda10_(d_124_i5_):
                        return (d_123_v0_) == (False)

                    return lambda10_

                init5_ = lambda9_(d_76_v0_)
                nw13_ = _dafny.Array(None, 6)
                for i0_5_ in range(nw13_.length(0)):
                    nw13_[i0_5_] = init5_(i0_5_)
                d_122_v40_ = nw13_
                index15_ = default__.safeIndex(376, (d_91_v12_).length(0))
                (d_91_v12_)[index15_] = d_77_v1_
                index16_ = default__.safeIndex(376, (d_91_v12_).length(0))
                rhs14_ = d_76_v0_
                rhs15_ = d_122_v40_
                rhs16_ = not (d_76_v0_) or (d_102_v22_)
                rhs17_ = 0
                rhs18_ = (_dafny.SeqWithoutIsStrInference([d_96_v15_ for d_125_i6_ in range(default__.abs(819))])).set(default__.safeIndex(d_77_v1_, len(_dafny.SeqWithoutIsStrInference([d_96_v15_ for d_126_i6_ in range(default__.abs(819))]))), d_96_v15_)
                lhs9_ = d_91_v12_
                lhs10_ = default__.safeIndex(376, (d_91_v12_).length(0))
                d_100_v20_ = rhs14_
                d_122_v40_ = rhs15_
                d_102_v22_ = rhs16_
                lhs9_[lhs10_] = rhs17_
                d_107_v26_ = rhs18_
                rhs19_ = (d_91_v12_)[default__.safeIndex(376, (d_91_v12_).length(0))]
                rhs20_ = d_109_v28_
                rhs21_ = True
                d_101_v21_ = rhs19_
                d_109_v28_ = rhs20_
                d_99_v19_ = rhs21_
                d_127_v41_: _dafny.Array
                nw14_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_127_v41_ = nw14_
                d_127_v41_ = d_127_v41_
                d_101_v21_ = (d_91_v12_)[default__.safeIndex(376, (d_91_v12_).length(0))]
        elif True:
            d_128___mcc_h0_ = source3_.cf0
            d_129_cf0_ = d_128___mcc_h0_
            if (d_79_v2_).issubset(d_79_v2_):
                d_130_v42_: str
                d_130_v42_ = _dafny.CodePoint('x')
                d_76_v0_ = (d_130_v42_) in (_dafny.SeqWithoutIsStrInference([d_130_v42_ for d_131_i7_ in range(default__.abs(-201))]))
                d_132_v43_: C1
                nw15_ = C1()
                nw15_.ctor__()
                d_132_v43_ = nw15_
                d_133_v44_: C1
                nw16_ = C1()
                nw16_.ctor__()
                d_133_v44_ = nw16_
                d_129_cf0_ = d_129_cf0_
                d_76_v0_ = d_76_v0_
            elif True:
                d_134_v45_: C1
                nw17_ = C1()
                nw17_.ctor__()
                d_134_v45_ = nw17_
                d_135_v46_: _dafny.Seq
                d_135_v46_ = _dafny.SeqWithoutIsStrInference([d_134_v45_])
                d_136_v47_: _dafny.Seq
                d_136_v47_ = _dafny.SeqWithoutIsStrInference([d_129_cf0_])
                d_137_v48_: _dafny.Array
                nw18_ = _dafny.Array(None, 12)
                nw18_[int(0)] = d_134_v45_
                nw18_[int(1)] = d_134_v45_
                nw18_[int(2)] = d_134_v45_
                nw18_[int(3)] = d_134_v45_
                nw18_[int(4)] = d_134_v45_
                nw18_[int(5)] = d_134_v45_
                nw18_[int(6)] = d_134_v45_
                nw18_[int(7)] = d_134_v45_
                nw18_[int(8)] = d_134_v45_
                nw18_[int(9)] = (d_135_v46_)[default__.safeIndex((d_136_v47_)[default__.safeIndex(d_77_v1_, len(d_136_v47_))], len(d_135_v46_))]
                nw18_[int(10)] = d_134_v45_
                nw18_[int(11)] = d_134_v45_
                d_137_v48_ = nw18_
                d_138_v49_: _dafny.Map
                d_138_v49_ = _dafny.Map({d_137_v48_: d_91_v12_})
                d_91_v12_ = ((d_138_v49_)[d_137_v48_] if (d_137_v48_) in (d_138_v49_) else d_91_v12_)
                d_139_v50_: _dafny.Array
                def lambda11_(d_140_v47_):
                    def lambda12_(d_141_i8_):
                        return d_140_v47_

                    return lambda12_

                init6_ = lambda11_(d_136_v47_)
                nw19_ = _dafny.Array(None, 27)
                for i0_6_ in range(nw19_.length(0)):
                    nw19_[i0_6_] = init6_(i0_6_)
                d_139_v50_ = nw19_
                d_142_v51_: C0
                nw20_ = C0()
                nw20_.ctor__(d_139_v50_)
                d_142_v51_ = nw20_
                d_143_v52_: D1
                d_143_v52_ = D1_DC3(d_76_v0_, d_76_v0_, d_76_v0_)
                d_143_v52_ = d_143_v52_
                d_144_v53_: _dafny.Seq
                d_144_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppn"))
                d_145_v54_: str
                d_145_v54_ = _dafny.CodePoint('m')
                d_144_v53_ = (d_144_v53_).set(default__.safeIndex(default__.safeModuloInt(d_129_cf0_, 10), len(d_144_v53_)), d_145_v54_)
                d_146_v55_: T1
                nw21_ = C0()
                nw21_.ctor__(d_139_v50_)
                d_146_v55_ = nw21_
                d_146_v55_ = d_146_v55_
            index17_ = default__.safeIndex(281, (d_91_v12_).length(0))
            (d_91_v12_)[index17_] = d_129_cf0_
            index18_ = default__.safeIndex(281, (d_91_v12_).length(0))
            (d_91_v12_)[index18_] = d_129_cf0_
            d_147_v56_: C1
            nw22_ = C1()
            nw22_.ctor__()
            d_147_v56_ = nw22_
            index19_ = default__.safeIndex(281, (d_91_v12_).length(0))
            (d_91_v12_)[index19_] = d_77_v1_
        d_77_v1_ = 882
        d_148_v57_: _dafny.Set
        d_148_v57_ = _dafny.Set({d_79_v2_})
        d_149_v58_: D2
        d_149_v58_ = D2_DC5(d_76_v0_, d_77_v1_, len(d_148_v57_))
        d_150_v59_: D2
        d_150_v59_ = D2_DC6(D2_DC6(d_149_v58_))
        source4_ = d_150_v59_
        if source4_.is_DC5:
            d_151___mcc_h1_ = source4_.cf6
            d_152___mcc_h2_ = source4_.cf7
            d_153___mcc_h3_ = source4_.cf8
            d_154_cf8_ = d_153___mcc_h3_
            d_155_cf7_ = d_152___mcc_h2_
            d_156_cf6_ = d_151___mcc_h1_
            d_157_v60_: _dafny.Map
            d_157_v60_ = _dafny.Map({d_154_cf8_: d_154_cf8_})
            d_158_v61_: _dafny.MultiSet
            d_158_v61_ = _dafny.MultiSet([d_156_cf6_])
            d_159_v62_: D2
            d_159_v62_ = D2_DC5((not(d_76_v0_)) or (d_156_cf6_), ((d_157_v60_)[default__.fm1(d_154_cf8_, d_154_cf8_, d_154_cf8_, d_76_v0_, d_78_globalState_)] if (default__.fm1(d_154_cf8_, d_154_cf8_, d_154_cf8_, d_76_v0_, d_78_globalState_)) in (d_157_v60_) else d_155_cf7_), (d_158_v61_).cardinality)
            source5_ = d_159_v62_
            if source5_.is_DC5:
                d_160___mcc_h6_ = source5_.cf6
                d_161___mcc_h7_ = source5_.cf7
                d_162___mcc_h8_ = source5_.cf8
                d_163_cf8_ = d_162___mcc_h8_
                d_164_cf7_ = d_161___mcc_h7_
                d_165_cf6_ = d_160___mcc_h6_
                index20_ = default__.safeIndex(380, (d_91_v12_).length(0))
                (d_91_v12_)[index20_] = default__.safeDivisionInt(d_164_cf7_, d_163_cf8_)
                index21_ = default__.safeIndex(380, (d_91_v12_).length(0))
                (d_91_v12_)[index21_] = d_163_cf8_
                d_166_v63_: _dafny.Set
                d_166_v63_ = _dafny.Set({42})
                d_167_v64_: _dafny.Seq
                d_167_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btxl"))
                d_168_v65_: _dafny.Array
                nw23_ = _dafny.Array(None, 13)
                nw23_[int(0)] = d_76_v0_
                nw23_[int(1)] = (d_94_v13_).isdisjoint(d_94_v13_)
                nw23_[int(2)] = d_165_cf6_
                nw23_[int(3)] = d_76_v0_
                nw23_[int(4)] = d_76_v0_
                nw23_[int(5)] = (d_166_v63_).issubset(d_166_v63_)
                nw23_[int(6)] = False
                nw23_[int(7)] = True
                nw23_[int(8)] = d_165_cf6_
                nw23_[int(9)] = (d_167_v64_) == (d_167_v64_)
                nw23_[int(10)] = d_76_v0_
                nw23_[int(11)] = d_76_v0_
                nw23_[int(12)] = (d_159_v62_).cf6
                d_168_v65_ = nw23_
                d_168_v65_ = d_168_v65_
                d_169_v66_: _dafny.Array
                nw24_ = _dafny.Array(int(0), 21)
                d_169_v66_ = nw24_
                d_170_v67_: _dafny.Seq
                d_170_v67_ = _dafny.SeqWithoutIsStrInference([d_169_v66_, d_169_v66_, d_169_v66_])
                d_171_v68_: _dafny.Set
                d_171_v68_ = _dafny.Set({(d_170_v67_)[default__.safeIndex(len(d_167_v64_), len(d_170_v67_))], d_169_v66_, d_169_v66_, d_169_v66_, d_169_v66_})
                d_171_v68_ = d_171_v68_
                d_76_v0_ = d_76_v0_
            elif source5_.is_DC4:
                d_172___mcc_h9_ = source5_.cf5
                d_173_cf5_ = d_172___mcc_h9_
                d_174_v69_: _dafny.Array
                nw25_ = _dafny.Array(False, 11)
                d_174_v69_ = nw25_
                index22_ = default__.safeIndex(97, (d_174_v69_).length(0))
                (d_174_v69_)[index22_] = not(not(d_156_cf6_))
                d_175_v70_: C1
                nw26_ = C1()
                nw26_.ctor__()
                d_175_v70_ = nw26_
                d_176_v71_: _dafny.Map
                d_176_v71_ = _dafny.Map({d_175_v70_: d_174_v69_})
                index23_ = default__.safeIndex(97, (d_174_v69_).length(0))
                (d_174_v69_)[index23_] = (d_175_v70_) not in (d_176_v71_)
                d_177_v72_: _dafny.Map
                d_177_v72_ = _dafny.Map({d_91_v12_: 666})
                d_177_v72_ = (d_177_v72_).set(d_91_v12_, d_155_cf7_)
                d_178_v73_: _dafny.Map
                d_178_v73_ = _dafny.Map({d_156_cf6_: d_156_cf6_})
                index24_ = default__.safeIndex(97, (d_174_v69_).length(0))
                (d_174_v69_)[index24_] = (d_156_cf6_ if (len(d_178_v73_)) > (d_155_cf7_) else d_156_cf6_)
                d_179_v74_: _dafny.Array
                def lambda13_(d_180_cf5_):
                    def lambda14_(d_181_i9_):
                        return (_dafny.MultiSet([d_180_cf5_])) - (_dafny.MultiSet([d_180_cf5_]))

                    return lambda14_

                init7_ = lambda13_(d_173_cf5_)
                nw27_ = _dafny.Array(None, 9)
                for i0_7_ in range(nw27_.length(0)):
                    nw27_[i0_7_] = init7_(i0_7_)
                d_179_v74_ = nw27_
                index25_ = default__.safeIndex(986, (d_179_v74_).length(0))
                (d_179_v74_)[index25_] = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_182_i10_ in range(default__.abs(-309))]), d_173_cf5_, d_173_cf5_])
                d_183_v75_: _dafny.Seq
                d_183_v75_ = _dafny.SeqWithoutIsStrInference([d_156_cf6_])
                d_184_v76_: _dafny.Seq
                d_184_v76_ = _dafny.SeqWithoutIsStrInference([d_183_v75_, default__.fm15(default__.fm1(d_77_v1_, d_155_cf7_, d_155_cf7_, (d_174_v69_)[default__.safeIndex(97, (d_174_v69_).length(0))], d_78_globalState_), d_78_globalState_), d_183_v75_, d_183_v75_])
                index26_ = default__.safeIndex(986, (d_179_v74_).length(0))
                (d_179_v74_)[index26_] = default__.fm14((d_173_cf5_) < (d_173_cf5_), d_184_v76_, False, d_78_globalState_)
            elif True:
                d_185___mcc_h10_ = source5_.cf9
                d_186_cf9_ = d_185___mcc_h10_
                d_187_v77_: _dafny.Array
                def lambda15_(d_188_cf6_):
                    def lambda16_(d_189_i11_):
                        return d_188_cf6_

                    return lambda16_

                init8_ = lambda15_(d_156_cf6_)
                nw28_ = _dafny.Array(None, 16)
                for i0_8_ in range(nw28_.length(0)):
                    nw28_[i0_8_] = init8_(i0_8_)
                d_187_v77_ = nw28_
                index27_ = default__.safeIndex(249, (d_187_v77_).length(0))
                (d_187_v77_)[index27_] = d_76_v0_
                index28_ = default__.safeIndex(249, (d_187_v77_).length(0))
                (d_187_v77_)[index28_] = not((d_155_cf7_) <= (d_155_cf7_))
                d_76_v0_ = (d_156_cf6_) == (d_76_v0_)
                d_190_v78_: C1
                nw29_ = C1()
                nw29_.ctor__()
                d_190_v78_ = nw29_
                d_191_v79_: _dafny.Seq
                d_191_v79_ = _dafny.SeqWithoutIsStrInference([d_77_v1_])
                d_192_v80_: _dafny.Map
                d_192_v80_ = _dafny.Map({_dafny.Map({d_91_v12_: 909}): default__.fm1(d_154_cf8_, default__.fm1((_dafny.MultiSet([d_190_v78_])).cardinality, d_77_v1_, d_155_cf7_, d_156_cf6_, d_78_globalState_), len(d_191_v79_), d_156_cf6_, d_78_globalState_)})
                d_193_v81_: str
                d_193_v81_ = _dafny.CodePoint('t')
                d_194_v82_: _dafny.Seq
                d_194_v82_ = _dafny.SeqWithoutIsStrInference([d_76_v0_, d_156_cf6_, d_76_v0_])
                d_195_v83_: _dafny.Map
                d_195_v83_ = _dafny.Map({d_193_v81_: len(d_194_v82_)})
                d_196_v84_: _dafny.Seq
                d_196_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gx"))
                d_155_cf7_ = ((d_192_v80_)[(_dafny.Map({d_91_v12_: len(d_195_v83_)})).set(d_91_v12_, len(d_196_v84_))] if ((_dafny.Map({d_91_v12_: len(d_195_v83_)})).set(d_91_v12_, len(d_196_v84_))) in (d_192_v80_) else d_154_cf8_)
                d_197_v85_: _dafny.Seq
                d_197_v85_ = _dafny.SeqWithoutIsStrInference([d_196_v84_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enxbdbrux"))])
                d_198_v86_: _dafny.Array
                nw30_ = _dafny.Array(None, 9)
                nw30_[int(0)] = (d_197_v85_) + (d_197_v85_)
                nw30_[int(1)] = d_197_v85_
                nw30_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_196_v84_ for d_199_i12_ in range(default__.abs(725))])) + (d_197_v85_)
                nw30_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_196_v84_ for d_200_i13_ in range(default__.abs(632))]) if not(False) else d_197_v85_)
                nw30_[int(4)] = (d_197_v85_) + (d_197_v85_)
                nw30_[int(5)] = (d_197_v85_) + ((d_197_v85_).set(default__.safeIndex(665, len(d_197_v85_)), d_196_v84_))
                nw30_[int(6)] = d_197_v85_
                nw30_[int(7)] = _dafny.SeqWithoutIsStrInference([d_196_v84_, d_196_v84_])
                nw30_[int(8)] = d_197_v85_
                d_198_v86_ = nw30_
                index29_ = default__.safeIndex(433, (d_198_v86_).length(0))
                (d_198_v86_)[index29_] = d_197_v85_
                index30_ = default__.safeIndex(914, (d_91_v12_).length(0))
                (d_91_v12_)[index30_] = (d_155_cf7_) + (d_154_cf8_)
                d_201_v87_: _dafny.Array
                nw31_ = _dafny.Array(None, 8)
                d_201_v87_ = nw31_
                d_202_v88_: _dafny.Array
                def lambda17_(d_203_cf8_):
                    def lambda18_(d_204_i14_):
                        return _dafny.SeqWithoutIsStrInference([d_203_cf8_ for d_205_i15_ in range(default__.abs(624))])

                    return lambda18_

                init9_ = lambda17_(d_154_cf8_)
                nw32_ = _dafny.Array(None, 14)
                for i0_9_ in range(nw32_.length(0)):
                    nw32_[i0_9_] = init9_(i0_9_)
                d_202_v88_ = nw32_
                d_206_v89_: C0
                nw33_ = C0()
                nw33_.ctor__(d_202_v88_)
                d_206_v89_ = nw33_
                index31_ = default__.safeIndex(580, (d_201_v87_).length(0))
                (d_201_v87_)[index31_] = d_206_v89_
                index32_ = default__.safeIndex(189, (d_91_v12_).length(0))
                (d_91_v12_)[index32_] = d_77_v1_
                index33_ = default__.safeIndex(433, (d_198_v86_).length(0))
                index34_ = default__.safeIndex(914, (d_91_v12_).length(0))
                index35_ = default__.safeIndex(580, (d_201_v87_).length(0))
                index36_ = default__.safeIndex(189, (d_91_v12_).length(0))
                rhs22_ = d_197_v85_
                rhs23_ = default__.safeDivisionInt(d_154_cf8_, default__.fm1(d_154_cf8_, d_77_v1_, len(_dafny.SeqWithoutIsStrInference([d_193_v81_ for d_207_i16_ in range(default__.abs(886))])), d_76_v0_, d_78_globalState_))
                rhs24_ = d_206_v89_
                rhs25_ = (d_187_v77_)[default__.safeIndex(249, (d_187_v77_).length(0))]
                rhs26_ = (0) - (d_155_cf7_)
                lhs11_ = d_198_v86_
                lhs12_ = default__.safeIndex(433, (d_198_v86_).length(0))
                lhs13_ = d_91_v12_
                lhs14_ = default__.safeIndex(914, (d_91_v12_).length(0))
                lhs15_ = d_201_v87_
                lhs16_ = default__.safeIndex(580, (d_201_v87_).length(0))
                lhs17_ = d_91_v12_
                lhs18_ = default__.safeIndex(189, (d_91_v12_).length(0))
                lhs11_[lhs12_] = rhs22_
                lhs13_[lhs14_] = rhs23_
                lhs15_[lhs16_] = rhs24_
                d_156_cf6_ = rhs25_
                lhs17_[lhs18_] = rhs26_
            d_208_v90_: _dafny.Array
            def lambda19_(d_209_v1_):
                def lambda20_(d_210_i17_):
                    return _dafny.SeqWithoutIsStrInference([d_209_v1_])

                return lambda20_

            init10_ = lambda19_(d_77_v1_)
            nw34_ = _dafny.Array(None, 12)
            for i0_10_ in range(nw34_.length(0)):
                nw34_[i0_10_] = init10_(i0_10_)
            d_208_v90_ = nw34_
            d_211_v91_: C0
            nw35_ = C0()
            nw35_.ctor__(d_208_v90_)
            d_211_v91_ = nw35_
            d_212_v92_: _dafny.Seq
            d_212_v92_ = _dafny.SeqWithoutIsStrInference([d_211_v91_])
            if (d_94_v13_) == (default__.fm16(len(_dafny.Set({d_155_cf7_, d_155_cf7_, len(d_212_v92_), d_154_cf8_})), d_78_globalState_)):
                d_213_v93_: _dafny.Map
                d_213_v93_ = _dafny.Map({d_154_cf8_: (d_79_v2_).issubset(d_79_v2_)})
                d_214_v94_: C1
                nw36_ = C1()
                nw36_.ctor__()
                d_214_v94_ = nw36_
                d_215_v95_: _dafny.Map
                d_215_v95_ = _dafny.Map({d_214_v94_: d_154_cf8_})
                d_216_v96_: _dafny.Map
                d_216_v96_ = _dafny.Map({d_215_v95_: d_79_v2_})
                d_217_v97_: _dafny.Seq
                d_217_v97_ = _dafny.SeqWithoutIsStrInference([d_77_v1_])
                d_218_v98_: _dafny.Map
                d_218_v98_ = _dafny.Map({d_154_cf8_: d_217_v97_})
                d_219_v99_: T1
                nw37_ = C0()
                nw37_.ctor__(d_211_v91_.f2)
                d_219_v99_ = nw37_
                d_220_v100_: D3
                d_220_v100_ = D3_DC8(d_156_cf6_, D2_DC6(d_149_v58_), d_76_v0_, d_79_v2_, d_214_v94_)
                d_221_v101_: _dafny.Array
                nw38_ = _dafny.Array(None, 24)
                nw38_[int(0)] = (d_79_v2_).intersection(d_79_v2_)
                nw38_[int(1)] = (d_79_v2_).intersection(d_79_v2_)
                nw38_[int(2)] = ((d_216_v96_)[(d_215_v95_).set(d_214_v94_, d_154_cf8_)] if ((d_215_v95_).set(d_214_v94_, d_154_cf8_)) in (d_216_v96_) else _dafny.Set({d_76_v0_}))
                nw38_[int(3)] = d_79_v2_
                nw38_[int(4)] = d_79_v2_
                nw38_[int(5)] = _dafny.Set({d_156_cf6_})
                nw38_[int(6)] = _dafny.Set({d_156_cf6_, d_76_v0_, d_76_v0_})
                nw38_[int(7)] = _dafny.Set({d_76_v0_, False, d_156_cf6_, d_156_cf6_})
                nw38_[int(8)] = (d_79_v2_) | (default__.fm5(len(((d_218_v98_)[len(_dafny.Map({_dafny.MultiSet([736]): d_219_v99_}))] if (len(_dafny.Map({_dafny.MultiSet([736]): d_219_v99_}))) in (d_218_v98_) else d_217_v97_)), d_76_v0_, d_156_cf6_, d_155_cf7_, d_78_globalState_))
                nw38_[int(9)] = (d_79_v2_ if d_156_cf6_ else d_79_v2_)
                nw38_[int(10)] = _dafny.Set({d_76_v0_, False})
                nw38_[int(11)] = d_79_v2_
                nw38_[int(12)] = d_79_v2_
                nw38_[int(13)] = (d_79_v2_) - (d_79_v2_)
                nw38_[int(14)] = (default__.fm5(d_77_v1_, d_156_cf6_, d_76_v0_, default__.fm1(d_154_cf8_, d_155_cf7_, 361, d_156_cf6_, d_78_globalState_), d_78_globalState_) if d_76_v0_ else d_79_v2_)
                nw38_[int(15)] = d_79_v2_
                nw38_[int(16)] = ((d_220_v100_).cf14) | (d_79_v2_)
                nw38_[int(17)] = (d_79_v2_) | (d_79_v2_)
                nw38_[int(18)] = d_79_v2_
                nw38_[int(19)] = d_79_v2_
                nw38_[int(20)] = d_79_v2_
                nw38_[int(21)] = _dafny.Set({True, d_76_v0_, d_156_cf6_, d_76_v0_})
                nw38_[int(22)] = (d_79_v2_).intersection(d_79_v2_)
                nw38_[int(23)] = d_79_v2_
                d_221_v101_ = nw38_
                d_222_v102_: _dafny.Seq
                d_222_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cr"))
                d_223_v103_: _dafny.Map
                d_223_v103_ = _dafny.Map({d_76_v0_: d_76_v0_})
                d_224_v104_: D1
                d_224_v104_ = D1_DC3(d_76_v0_, d_76_v0_, d_156_cf6_)
                d_225_v105_: _dafny.Map
                d_225_v105_ = _dafny.Map({d_224_v104_: d_76_v0_})
                d_226_v106_: _dafny.Seq
                d_226_v106_ = _dafny.SeqWithoutIsStrInference([d_225_v105_])
                rhs27_ = d_213_v93_
                rhs28_ = d_155_cf7_
                rhs29_ = d_221_v101_
                rhs30_ = (d_222_v102_ if d_76_v0_ else d_222_v102_)
                rhs31_ = ((d_223_v103_)[d_76_v0_] if (d_76_v0_) in (d_223_v103_) else (d_211_v91_).fm7(d_226_v106_, d_76_v0_, d_154_cf8_, d_78_globalState_))
                d_213_v93_ = rhs27_
                d_77_v1_ = rhs28_
                d_221_v101_ = rhs29_
                d_222_v102_ = rhs30_
                d_156_cf6_ = rhs31_
                d_156_cf6_ = d_76_v0_
                d_227_v107_: _dafny.Array
                def lambda21_(d_228_cf6_):
                    def lambda22_(d_229_i18_):
                        return d_228_cf6_

                    return lambda22_

                init11_ = lambda21_(d_156_cf6_)
                nw39_ = _dafny.Array(None, 3)
                for i0_11_ in range(nw39_.length(0)):
                    nw39_[i0_11_] = init11_(i0_11_)
                d_227_v107_ = nw39_
                index37_ = default__.safeIndex(917, (d_227_v107_).length(0))
                (d_227_v107_)[index37_] = (d_94_v13_).issubset(d_94_v13_)
                d_230_v108_: _dafny.Seq
                d_230_v108_ = _dafny.SeqWithoutIsStrInference([d_156_cf6_])
                d_231_v109_: str
                d_231_v109_ = _dafny.CodePoint('y')
                index38_ = default__.safeIndex(917, (d_227_v107_).length(0))
                rhs32_ = not(d_76_v0_)
                rhs33_ = default__.fm0(len(d_230_v108_), d_154_cf8_, d_231_v109_, d_78_globalState_)
                rhs34_ = D0_DC0((0) - (d_155_cf7_))
                lhs19_ = d_227_v107_
                lhs20_ = default__.safeIndex(917, (d_227_v107_).length(0))
                lhs19_[lhs20_] = rhs32_
                d_156_cf6_ = rhs33_
                d_80_v3_ = rhs34_
                d_214_v94_ = ((d_214_v94_ if d_156_cf6_ else d_214_v94_) if (d_227_v107_)[default__.safeIndex(917, (d_227_v107_).length(0))] else d_214_v94_)
                rhs35_ = (0) - (default__.safeModuloInt(698, default__.fm1(d_154_cf8_, (D3_DC9(d_155_cf7_)).cf16, d_154_cf8_, d_76_v0_, d_78_globalState_)))
                rhs36_ = d_77_v1_
                d_154_cf8_ = rhs35_
                d_77_v1_ = rhs36_
            elif True:
                d_232_v110_: _dafny.Map
                d_232_v110_ = _dafny.Map({(0) - (d_77_v1_): d_79_v2_})
                d_233_v111_: _dafny.Seq
                d_233_v111_ = _dafny.SeqWithoutIsStrInference([d_79_v2_, ((d_232_v110_)[len(d_79_v2_)] if (len(d_79_v2_)) in (d_232_v110_) else d_79_v2_), default__.fm5(d_77_v1_, d_76_v0_, d_156_cf6_, d_77_v1_, d_78_globalState_)])
                d_233_v111_ = ((_dafny.SeqWithoutIsStrInference([d_79_v2_ for d_234_i19_ in range(default__.abs(137))])) + (d_233_v111_)) + (d_233_v111_)
                d_235_v112_: _dafny.Array
                nw40_ = _dafny.Array(False, 1)
                d_235_v112_ = nw40_
                index39_ = default__.safeIndex(539, (d_235_v112_).length(0))
                (d_235_v112_)[index39_] = d_76_v0_
                index40_ = default__.safeIndex(539, (d_235_v112_).length(0))
                (d_235_v112_)[index40_] = (d_155_cf7_) > (d_154_cf8_)
                d_236_v113_: _dafny.Seq
                d_236_v113_ = _dafny.SeqWithoutIsStrInference([(d_155_cf7_) < (d_155_cf7_), (d_235_v112_)[default__.safeIndex(539, (d_235_v112_).length(0))]])
                d_236_v113_ = (_dafny.SeqWithoutIsStrInference([d_156_cf6_, d_76_v0_, True])) + (d_236_v113_)
                d_237_v114_: T0
                nw41_ = C2()
                nw41_.ctor__(d_77_v1_)
                d_237_v114_ = nw41_
                d_238_v115_: _dafny.Map
                d_238_v115_ = _dafny.Map({d_76_v0_: d_237_v114_})
                d_239_v116_: _dafny.Seq
                d_239_v116_ = _dafny.SeqWithoutIsStrInference([(d_238_v115_) | (d_238_v115_)])
                d_240_v117_: _dafny.Seq
                d_240_v117_ = _dafny.SeqWithoutIsStrInference([d_239_v116_, d_239_v116_])
                d_239_v116_ = (((d_240_v117_)[default__.safeIndex(-164, len(d_240_v117_))]) + (d_239_v116_)) + (d_239_v116_)
                d_235_v112_ = d_235_v112_
            d_241_v118_: _dafny.Map
            d_241_v118_ = _dafny.Map({d_156_cf6_: (d_76_v0_ if d_76_v0_ else d_156_cf6_)})
            d_241_v118_ = (d_241_v118_).set(d_156_cf6_, d_76_v0_)
            d_76_v0_ = d_76_v0_
        elif source4_.is_DC4:
            d_242___mcc_h4_ = source4_.cf5
            d_243_cf5_ = d_242___mcc_h4_
            d_244_v119_: _dafny.Array
            def lambda23_(d_245_v2_):
                def lambda24_(d_246_i20_):
                    return D3_DC7(d_245_v2_)

                return lambda24_

            init12_ = lambda23_(d_79_v2_)
            nw42_ = _dafny.Array(None, 28)
            for i0_12_ in range(nw42_.length(0)):
                nw42_[i0_12_] = init12_(i0_12_)
            d_244_v119_ = nw42_
            d_247_v120_: D3
            d_247_v120_ = D3_DC7(d_79_v2_)
            index41_ = default__.safeIndex(863, (d_244_v119_).length(0))
            (d_244_v119_)[index41_] = d_247_v120_
            d_248_v121_: _dafny.Array
            nw43_ = _dafny.Array(False, 24)
            d_248_v121_ = nw43_
            d_249_v122_: _dafny.Array
            nw44_ = _dafny.Array(None, 7)
            nw44_[int(0)] = d_248_v121_
            nw44_[int(1)] = d_248_v121_
            nw44_[int(2)] = d_248_v121_
            nw44_[int(3)] = d_248_v121_
            nw44_[int(4)] = d_248_v121_
            nw44_[int(5)] = d_248_v121_
            nw44_[int(6)] = d_248_v121_
            d_249_v122_ = nw44_
            index42_ = default__.safeIndex(247, (d_249_v122_).length(0))
            (d_249_v122_)[index42_] = d_248_v121_
            d_250_v123_: _dafny.Map
            d_250_v123_ = _dafny.Map({_dafny.Set({d_248_v121_, d_248_v121_, d_248_v121_}): d_91_v12_})
            d_251_v124_: _dafny.Set
            d_251_v124_ = _dafny.Set({d_248_v121_})
            index43_ = default__.safeIndex(863, (d_244_v119_).length(0))
            index44_ = default__.safeIndex(247, (d_249_v122_).length(0))
            rhs37_ = ((d_250_v123_)[(d_251_v124_) - (d_251_v124_)] if ((d_251_v124_) - (d_251_v124_)) in (d_250_v123_) else d_91_v12_)
            rhs38_ = d_247_v120_
            rhs39_ = d_248_v121_
            rhs40_ = ((d_77_v1_) * (d_77_v1_)) * (d_77_v1_)
            rhs41_ = d_76_v0_
            lhs21_ = d_244_v119_
            lhs22_ = default__.safeIndex(863, (d_244_v119_).length(0))
            lhs23_ = d_249_v122_
            lhs24_ = default__.safeIndex(247, (d_249_v122_).length(0))
            d_91_v12_ = rhs37_
            lhs21_[lhs22_] = rhs38_
            lhs23_[lhs24_] = rhs39_
            d_77_v1_ = rhs40_
            d_76_v0_ = rhs41_
            d_252_v125_: _dafny.Seq
            d_252_v125_ = _dafny.SeqWithoutIsStrInference([(d_77_v1_) == (d_77_v1_), d_76_v0_])
            d_252_v125_ = _dafny.SeqWithoutIsStrInference([d_76_v0_, d_76_v0_, d_76_v0_])
            d_76_v0_ = d_76_v0_
            d_77_v1_ = (0) - ((((0) - (d_77_v1_)) - (default__.fm1(d_77_v1_, d_77_v1_, d_77_v1_, d_76_v0_, d_78_globalState_))) + (d_77_v1_))
        elif True:
            d_253___mcc_h5_ = source4_.cf9
            d_254_cf9_ = d_253___mcc_h5_
            d_255_v126_: _dafny.Array
            def lambda25_(d_256_v0_):
                def lambda26_(d_257_i21_):
                    return (_dafny.SeqWithoutIsStrInference([not(d_256_v0_)])) + (_dafny.SeqWithoutIsStrInference([d_256_v0_]))

                return lambda26_

            init13_ = lambda25_(d_76_v0_)
            nw45_ = _dafny.Array(None, 27)
            for i0_13_ in range(nw45_.length(0)):
                nw45_[i0_13_] = init13_(i0_13_)
            d_255_v126_ = nw45_
            d_258_v127_: _dafny.Seq
            d_258_v127_ = _dafny.SeqWithoutIsStrInference([d_76_v0_])
            index45_ = default__.safeIndex(851, (d_255_v126_).length(0))
            (d_255_v126_)[index45_] = (d_258_v127_) + (d_258_v127_)
            index46_ = default__.safeIndex(851, (d_255_v126_).length(0))
            (d_255_v126_)[index46_] = d_258_v127_
            d_76_v0_ = (-732) < (d_77_v1_)
            d_259_v128_: str
            d_259_v128_ = _dafny.CodePoint('v')
            d_260_v129_: _dafny.Map
            d_260_v129_ = _dafny.Map({False: d_259_v128_})
            d_261_v130_: _dafny.Map
            d_261_v130_ = _dafny.Map({d_94_v13_: d_79_v2_})
            d_262_v131_: bool
            d_263_v132_: bool
            d_264_v133_: int
            d_265_v134_: bool
            out8_: bool
            out9_: bool
            out10_: int
            out11_: bool
            out8_, out9_, out10_, out11_ = default__.m0(d_260_v129_, d_261_v130_, d_76_v0_, default__.safeModuloInt(d_77_v1_, d_77_v1_), d_78_globalState_)
            d_262_v131_ = out8_
            d_263_v132_ = out9_
            d_264_v133_ = out10_
            d_265_v134_ = out11_
            d_266_v135_: _dafny.Seq
            d_266_v135_ = _dafny.SeqWithoutIsStrInference([d_264_v133_, d_77_v1_])
            d_267_v136_: _dafny.Array
            nw46_ = _dafny.Array(None, 7)
            nw46_[int(0)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(True)]))])
            nw46_[int(1)] = d_266_v135_
            nw46_[int(2)] = d_266_v135_
            nw46_[int(3)] = d_266_v135_
            nw46_[int(4)] = d_266_v135_
            nw46_[int(5)] = d_266_v135_
            nw46_[int(6)] = (d_266_v135_).set(default__.safeIndex(d_77_v1_, len(d_266_v135_)), d_264_v133_)
            d_267_v136_ = nw46_
            d_268_v137_: T1
            nw47_ = C0()
            nw47_.ctor__(d_267_v136_)
            d_268_v137_ = nw47_
            d_269_v138_: D4
            d_269_v138_ = D4_DC10(d_268_v137_)
            d_270_v139_: _dafny.MultiSet
            d_270_v139_ = _dafny.MultiSet([(d_269_v138_).cf17, d_268_v137_, d_268_v137_, d_268_v137_, d_268_v137_])
            d_270_v139_ = (((d_270_v139_).set(d_268_v137_, default__.abs(d_264_v133_))) | (d_270_v139_)).set(d_268_v137_, default__.abs(default__.safeDivisionInt(d_264_v133_, -895)))
        source6_ = default__.fm2(d_79_v2_, not(d_76_v0_), d_76_v0_, False, d_78_globalState_)
        if source6_.is_DC1:
            d_76_v0_ = False
            d_271_v140_: _dafny.Array
            nw48_ = _dafny.Array(False, 27)
            d_271_v140_ = nw48_
            d_272_v141_: _dafny.Seq
            d_272_v141_ = _dafny.SeqWithoutIsStrInference([d_271_v140_])
            d_273_v142_: _dafny.Array
            nw49_ = _dafny.Array(None, 14)
            nw49_[int(0)] = d_271_v140_
            nw49_[int(1)] = d_271_v140_
            nw49_[int(2)] = d_271_v140_
            nw49_[int(3)] = d_271_v140_
            nw49_[int(4)] = d_271_v140_
            nw49_[int(5)] = (d_272_v141_)[default__.safeIndex((0) - (d_77_v1_), len(d_272_v141_))]
            nw49_[int(6)] = d_271_v140_
            nw49_[int(7)] = d_271_v140_
            nw49_[int(8)] = d_271_v140_
            nw49_[int(9)] = d_271_v140_
            nw49_[int(10)] = d_271_v140_
            nw49_[int(11)] = d_271_v140_
            nw49_[int(12)] = d_271_v140_
            nw49_[int(13)] = d_271_v140_
            d_273_v142_ = nw49_
            index47_ = default__.safeIndex(760, (d_91_v12_).length(0))
            (d_91_v12_)[index47_] = d_77_v1_
            d_274_v143_: str
            d_274_v143_ = _dafny.CodePoint('s')
            index48_ = default__.safeIndex(760, (d_91_v12_).length(0))
            rhs42_ = (d_273_v142_ if default__.fm0(d_77_v1_, d_77_v1_, d_274_v143_, d_78_globalState_) else d_273_v142_)
            rhs43_ = d_77_v1_
            rhs44_ = default__.safeDivisionInt(d_77_v1_, d_77_v1_)
            lhs25_ = d_91_v12_
            lhs26_ = default__.safeIndex(760, (d_91_v12_).length(0))
            d_273_v142_ = rhs42_
            d_77_v1_ = rhs43_
            lhs25_[lhs26_] = rhs44_
            d_275_v144_: _dafny.Map
            d_275_v144_ = _dafny.Map({d_76_v0_: d_274_v143_})
            d_276_v145_: _dafny.Seq
            d_276_v145_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmrorbj"))
            d_277_v146_: _dafny.Map
            d_277_v146_ = _dafny.Map({len(d_276_v145_): d_79_v2_})
            d_278_v147_: _dafny.Map
            d_278_v147_ = _dafny.Map({(_dafny.MultiSet([(d_91_v12_)[default__.safeIndex(760, (d_91_v12_).length(0))]])).set(len(d_276_v145_), default__.abs((d_91_v12_)[default__.safeIndex(760, (d_91_v12_).length(0))])): ((d_277_v146_)[len(default__.fm12(d_78_globalState_))] if (len(default__.fm12(d_78_globalState_))) in (d_277_v146_) else d_79_v2_)})
            d_279_v148_: bool
            d_280_v149_: bool
            d_281_v150_: int
            d_282_v151_: bool
            out12_: bool
            out13_: bool
            out14_: int
            out15_: bool
            out12_, out13_, out14_, out15_ = default__.m0(d_275_v144_, d_278_v147_, (d_76_v0_ if d_76_v0_ else d_76_v0_), d_77_v1_, d_78_globalState_)
            d_279_v148_ = out12_
            d_280_v149_ = out13_
            d_281_v150_ = out14_
            d_282_v151_ = out15_
            index49_ = default__.safeIndex(209, (d_271_v140_).length(0))
            (d_271_v140_)[index49_] = not (d_282_v151_) or (d_76_v0_)
            index50_ = default__.safeIndex(209, (d_271_v140_).length(0))
            (d_271_v140_)[index50_] = ((d_76_v0_) == (d_280_v149_)) == (d_282_v151_)
        elif True:
            d_283___mcc_h11_ = source6_.cf0
            d_284_cf0_ = d_283___mcc_h11_
            if not ((d_76_v0_) == (d_76_v0_)) or (d_76_v0_):
                d_285_v152_: _dafny.Map
                d_285_v152_ = _dafny.Map({not(d_76_v0_): not(d_76_v0_)})
                d_286_v153_: _dafny.Seq
                d_286_v153_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysdj"))
                d_77_v1_ = len(((d_286_v153_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))) if ((d_285_v152_)[d_76_v0_] if (d_76_v0_) in (d_285_v152_) else True) else d_286_v153_))
                d_287_v154_: _dafny.Map
                d_287_v154_ = _dafny.Map({(d_79_v2_) | (d_79_v2_): d_284_cf0_})
                d_287_v154_ = ((d_287_v154_) | (d_287_v154_)) | (d_287_v154_)
                d_288_v155_: _dafny.Seq
                d_288_v155_ = _dafny.SeqWithoutIsStrInference([d_77_v1_, d_77_v1_])
                d_288_v155_ = ((d_288_v155_) + (d_288_v155_)) + (d_288_v155_)
                d_289_v156_: _dafny.Array
                nw50_ = _dafny.Array(False, 28)
                d_289_v156_ = nw50_
                index51_ = default__.safeIndex(52, (d_289_v156_).length(0))
                (d_289_v156_)[index51_] = d_76_v0_
                index52_ = default__.safeIndex(52, (d_289_v156_).length(0))
                (d_289_v156_)[index52_] = (d_79_v2_).issubset(_dafny.Set({d_76_v0_}))
                d_290_v157_: C2
                nw51_ = C2()
                nw51_.ctor__(d_284_cf0_)
                d_290_v157_ = nw51_
                d_290_v157_ = d_290_v157_
            elif True:
                d_291_v158_: _dafny.Seq
                d_291_v158_ = _dafny.SeqWithoutIsStrInference([d_284_cf0_, d_284_cf0_])
                d_292_v159_: _dafny.MultiSet
                d_292_v159_ = _dafny.MultiSet([d_76_v0_])
                d_293_v160_: str
                d_293_v160_ = _dafny.CodePoint('f')
                d_294_v161_: _dafny.Map
                d_294_v161_ = _dafny.Map({((d_292_v159_)[d_76_v0_] if (d_76_v0_) in (d_292_v159_) else d_77_v1_): d_293_v160_})
                index53_ = default__.safeIndex(393, (d_91_v12_).length(0))
                (d_91_v12_)[index53_] = len(d_294_v161_)
                d_295_v162_: _dafny.Seq
                d_295_v162_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icr"))
                d_296_v163_: _dafny.Set
                d_296_v163_ = _dafny.Set({(d_295_v162_).set(default__.safeIndex(d_77_v1_, len(d_295_v162_)), d_293_v160_)})
                d_297_v164_: _dafny.Seq
                d_297_v164_ = _dafny.SeqWithoutIsStrInference([d_295_v162_])
                index54_ = default__.safeIndex(393, (d_91_v12_).length(0))
                rhs45_ = (d_296_v163_) == ((d_296_v163_).intersection(d_296_v163_))
                rhs46_ = d_291_v158_
                rhs47_ = d_77_v1_
                rhs48_ = (d_297_v164_)[default__.safeIndex(d_77_v1_, len(d_297_v164_))]
                lhs27_ = d_91_v12_
                lhs28_ = default__.safeIndex(393, (d_91_v12_).length(0))
                d_76_v0_ = rhs45_
                d_291_v158_ = rhs46_
                lhs27_[lhs28_] = rhs47_
                d_295_v162_ = rhs48_
                d_76_v0_ = d_76_v0_
                d_77_v1_ = ((0) - (default__.safeDivisionInt(d_284_cf0_, len(d_295_v162_)))) - ((d_91_v12_)[default__.safeIndex(393, (d_91_v12_).length(0))])
                d_284_cf0_ = (d_91_v12_)[default__.safeIndex(393, (d_91_v12_).length(0))]
                index55_ = default__.safeIndex(393, (d_91_v12_).length(0))
                (d_91_v12_)[index55_] = (d_284_cf0_) + (18)
            d_298_v165_: str
            d_298_v165_ = _dafny.CodePoint('l')
            d_299_v166_: _dafny.Map
            d_299_v166_ = _dafny.Map({d_76_v0_: d_298_v165_})
            d_300_v167_: _dafny.Seq
            d_300_v167_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "me"))
            d_301_v168_: _dafny.Map
            d_301_v168_ = _dafny.Map({(_dafny.MultiSet([d_284_cf0_, d_77_v1_])).set(len(d_300_v167_), default__.abs((d_94_v13_).cardinality)): _dafny.Set({d_76_v0_})})
            d_302_v169_: _dafny.Set
            d_302_v169_ = _dafny.Set({d_77_v1_})
            d_303_v170_: _dafny.Map
            d_303_v170_ = _dafny.Map({866: d_76_v0_})
            d_304_v172_: bool
            d_305_v173_: bool
            d_306_v174_: int
            d_307_v175_: bool
            out16_: bool
            out17_: bool
            out18_: int
            out19_: bool
            def iife3_():
                coll3_ = _dafny.Set()
                compr_3_: int
                for compr_3_ in (d_303_v170_).keys.Elements:
                    d_308_v171_: int = compr_3_
                    if (d_308_v171_) in (d_303_v170_):
                        coll3_ = coll3_.union(_dafny.Set([(d_308_v171_) - (304)]))
                return _dafny.Set(coll3_)
            out16_, out17_, out18_, out19_ = default__.m0(d_299_v166_, d_301_v168_, (d_302_v169_).isdisjoint(iife3_()
            ), default__.safeDivisionInt(d_77_v1_, d_77_v1_), d_78_globalState_)
            d_304_v172_ = out16_
            d_305_v173_ = out17_
            d_306_v174_ = out18_
            d_307_v175_ = out19_
            d_95_v14_ = d_95_v14_
            d_77_v1_ = (d_77_v1_) + (d_284_cf0_)
        d_309_v176_: str
        d_309_v176_ = _dafny.CodePoint('q')
        d_310_v177_: _dafny.Map
        d_310_v177_ = _dafny.Map({d_309_v176_: d_76_v0_})
        d_311_v178_: D2
        d_311_v178_ = D2_DC5(not(((d_310_v177_)[d_309_v176_] if (d_309_v176_) in (d_310_v177_) else False)), d_77_v1_, d_77_v1_)
        source7_ = d_311_v178_
        if source7_.is_DC5:
            d_312___mcc_h12_ = source7_.cf6
            d_313___mcc_h13_ = source7_.cf7
            d_314___mcc_h14_ = source7_.cf8
            d_315_cf8_ = d_314___mcc_h14_
            d_316_cf7_ = d_313___mcc_h13_
            d_317_cf6_ = d_312___mcc_h12_
            pat_let_tv3_ = d_316_cf7_
            def iife4_(_pat_let0_0):
                def iife5_(d_318_dt__update__tmp_h0_):
                    def iife6_(_pat_let1_0):
                        def iife7_(d_319_dt__update_hcf0_h0_):
                            return D0_DC0(d_319_dt__update_hcf0_h0_)
                        return iife7_(_pat_let1_0)
                    return iife6_(pat_let_tv3_)
                return iife5_(_pat_let0_0)
            source8_ = iife4_(d_80_v3_)
            if source8_.is_DC1:
                d_320_v179_: C1
                nw52_ = C1()
                nw52_.ctor__()
                d_320_v179_ = nw52_
                d_309_v176_ = default__.fm3(d_78_globalState_)
                d_321_v180_: D0
                d_322_v181_: int
                out20_: D0
                out21_: int
                out20_, out21_ = (d_320_v179_).m1(d_78_globalState_)
                d_321_v180_ = out20_
                d_322_v181_ = out21_
                index56_ = default__.safeIndex(585, (d_91_v12_).length(0))
                (d_91_v12_)[index56_] = d_316_cf7_
                index57_ = default__.safeIndex(585, (d_91_v12_).length(0))
                (d_91_v12_)[index57_] = (0) - (d_77_v1_)
            elif True:
                d_323___mcc_h17_ = source8_.cf0
                d_324_cf0_ = d_323___mcc_h17_
                d_325_v182_: _dafny.Array
                nw53_ = _dafny.Array(_dafny.Seq({}), 2)
                d_325_v182_ = nw53_
                d_326_v183_: C0
                nw54_ = C0()
                nw54_.ctor__(d_325_v182_)
                d_326_v183_ = nw54_
                d_327_v184_: _dafny.Array
                def lambda27_(d_328_v1_, d_329_cf8_):
                    def lambda28_(d_330_i22_):
                        return (d_328_v1_) in (_dafny.Set({d_329_cf8_}))

                    return lambda28_

                init14_ = lambda27_(d_77_v1_, d_315_cf8_)
                nw55_ = _dafny.Array(None, 23)
                for i0_14_ in range(nw55_.length(0)):
                    nw55_[i0_14_] = init14_(i0_14_)
                d_327_v184_ = nw55_
                d_331_v185_: _dafny.Seq
                d_331_v185_ = _dafny.SeqWithoutIsStrInference([d_327_v184_, d_327_v184_])
                d_332_v186_: _dafny.Seq
                d_332_v186_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnhtq"))
                d_327_v184_ = (d_331_v185_)[default__.safeIndex(len(d_332_v186_), len(d_331_v185_))]
                d_77_v1_ = (d_324_cf0_ if d_76_v0_ else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qixb"))))
                d_333_v187_: _dafny.Map
                d_333_v187_ = _dafny.Map({default__.fm12(d_78_globalState_): d_309_v176_})
                rhs49_ = d_333_v187_
                rhs50_ = (0) - (d_77_v1_)
                d_333_v187_ = rhs49_
                d_77_v1_ = rhs50_
            d_334_v188_: C1
            nw56_ = C1()
            nw56_.ctor__()
            d_334_v188_ = nw56_
            d_335_v189_: D3
            d_335_v189_ = D3_DC8(d_317_cf6_, d_150_v59_, d_76_v0_, d_79_v2_, d_334_v188_)
            d_336_v190_: _dafny.Map
            d_336_v190_ = _dafny.Map({d_309_v176_: d_335_v189_})
            d_337_v191_: D5
            d_337_v191_ = D5_DC13(d_336_v190_)
            d_338_v192_: D3
            d_338_v192_ = D3_DC9(len((d_337_v191_).cf19))
            d_339_v193_: _dafny.Seq
            d_339_v193_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ka"))
            d_340_v194_: _dafny.Seq
            d_340_v194_ = _dafny.SeqWithoutIsStrInference([len(d_339_v193_)])
            d_341_v195_: _dafny.MultiSet
            d_341_v195_ = _dafny.MultiSet([d_340_v194_, d_340_v194_, d_340_v194_, _dafny.SeqWithoutIsStrInference([d_77_v1_ for d_342_i23_ in range(default__.abs(829))])])
            d_343_v196_: _dafny.Map
            d_343_v196_ = _dafny.Map({d_77_v1_: d_341_v195_})
            rhs51_ = (d_338_v192_).cf16
            rhs52_ = d_77_v1_
            rhs53_ = ((((d_343_v196_)[d_316_cf7_] if (d_316_cf7_) in (d_343_v196_) else d_341_v195_)).set(d_340_v194_, default__.abs(d_77_v1_))).ispropersubset(d_341_v195_)
            d_316_cf7_ = rhs51_
            d_77_v1_ = rhs52_
            d_317_cf6_ = rhs53_
            d_344_v197_: _dafny.Array
            nw57_ = _dafny.Array(False, 21)
            d_344_v197_ = nw57_
            d_344_v197_ = d_344_v197_
            d_345_v198_: _dafny.Array
            def lambda29_(d_346_v194_):
                def lambda30_(d_347_i24_):
                    return d_346_v194_

                return lambda30_

            init15_ = lambda29_(d_340_v194_)
            nw58_ = _dafny.Array(None, 17)
            for i0_15_ in range(nw58_.length(0)):
                nw58_[i0_15_] = init15_(i0_15_)
            d_345_v198_ = nw58_
            d_348_v199_: T1
            nw59_ = C0()
            nw59_.ctor__(d_345_v198_)
            d_348_v199_ = nw59_
            d_349_v200_: _dafny.Map
            d_349_v200_ = _dafny.Map({d_316_cf7_: d_348_v199_})
            d_350_v201_: _dafny.Set
            d_350_v201_ = _dafny.Set({d_349_v200_})
            index58_ = default__.safeIndex(609, (d_91_v12_).length(0))
            (d_91_v12_)[index58_] = d_315_cf8_
            index59_ = default__.safeIndex(609, (d_91_v12_).length(0))
            rhs54_ = (d_350_v201_) | ((d_350_v201_ if d_76_v0_ else d_350_v201_))
            rhs55_ = d_76_v0_
            rhs56_ = d_348_v199_
            rhs57_ = d_316_cf7_
            lhs29_ = d_91_v12_
            lhs30_ = default__.safeIndex(609, (d_91_v12_).length(0))
            d_350_v201_ = rhs54_
            d_317_cf6_ = rhs55_
            d_348_v199_ = rhs56_
            lhs29_[lhs30_] = rhs57_
        elif source7_.is_DC4:
            d_351___mcc_h15_ = source7_.cf5
            d_352_cf5_ = d_351___mcc_h15_
            if d_76_v0_:
                d_91_v12_ = d_91_v12_
                d_353_v202_: _dafny.Array
                def lambda31_(d_354_i25_):
                    return True

                init16_ = lambda31_
                nw60_ = _dafny.Array(None, 14)
                for i0_16_ in range(nw60_.length(0)):
                    nw60_[i0_16_] = init16_(i0_16_)
                d_353_v202_ = nw60_
                d_355_v203_: D1
                d_355_v203_ = D1_DC3(d_76_v0_, not(d_76_v0_), d_76_v0_)
                d_356_v204_: _dafny.Map
                d_356_v204_ = _dafny.Map({(d_355_v203_).cf4: -932})
                d_357_v205_: _dafny.Array
                nw61_ = _dafny.Array(_dafny.Seq({}), 28)
                d_357_v205_ = nw61_
                d_358_v206_: C0
                nw62_ = C0()
                nw62_.ctor__(d_357_v205_)
                d_358_v206_ = nw62_
                d_359_v207_: _dafny.MultiSet
                d_359_v207_ = _dafny.MultiSet([d_358_v206_])
                d_360_v208_: _dafny.Map
                d_360_v208_ = _dafny.Map({((d_356_v204_)[d_76_v0_] if (d_76_v0_) in (d_356_v204_) else (d_359_v207_).cardinality): not(d_76_v0_)})
                index60_ = default__.safeIndex(935, (d_353_v202_).length(0))
                (d_353_v202_)[index60_] = ((d_360_v208_)[204] if (204) in (d_360_v208_) else d_76_v0_)
                index61_ = default__.safeIndex(935, (d_353_v202_).length(0))
                (d_353_v202_)[index61_] = d_76_v0_
                d_150_v59_ = D2_DC6(d_149_v58_)
                d_361_v209_: _dafny.Map
                d_361_v209_ = _dafny.Map({(d_94_v13_).intersection(d_94_v13_): d_77_v1_})
                d_362_v210_: _dafny.Seq
                d_362_v210_ = _dafny.SeqWithoutIsStrInference([d_352_cf5_])
                d_363_v211_: _dafny.Array
                nw63_ = _dafny.Array(None, 1)
                nw63_[int(0)] = d_362_v210_
                d_363_v211_ = nw63_
                d_364_v212_: D6
                d_364_v212_ = D6_DC17(d_362_v210_)
                d_365_v213_: _dafny.Seq
                d_365_v213_ = _dafny.SeqWithoutIsStrInference([(d_364_v212_).cf23])
                index62_ = default__.safeIndex(779, (d_363_v211_).length(0))
                (d_363_v211_)[index62_] = (d_365_v213_)[default__.safeIndex(d_77_v1_, len(d_365_v213_))]
                index63_ = default__.safeIndex(935, (d_353_v202_).length(0))
                index64_ = default__.safeIndex(779, (d_363_v211_).length(0))
                rhs58_ = (d_353_v202_)[default__.safeIndex(935, (d_353_v202_).length(0))]
                rhs59_ = (d_77_v1_) <= (default__.safeDivisionInt(len((default__.fm15(len(d_356_v204_), d_78_globalState_)).set(default__.safeIndex(d_77_v1_, len(default__.fm15(len(d_356_v204_), d_78_globalState_))), d_76_v0_)), d_77_v1_))
                rhs60_ = d_361_v209_
                rhs61_ = d_362_v210_
                lhs31_ = d_353_v202_
                lhs32_ = default__.safeIndex(935, (d_353_v202_).length(0))
                lhs33_ = d_363_v211_
                lhs34_ = default__.safeIndex(779, (d_363_v211_).length(0))
                d_76_v0_ = rhs58_
                lhs31_[lhs32_] = rhs59_
                d_361_v209_ = rhs60_
                lhs33_[lhs34_] = rhs61_
                d_366_v214_: _dafny.Map
                d_366_v214_ = _dafny.Map({(d_353_v202_)[default__.safeIndex(935, (d_353_v202_).length(0))]: d_309_v176_})
                d_367_v215_: bool
                d_368_v216_: bool
                d_369_v217_: int
                d_370_v218_: bool
                out22_: bool
                out23_: bool
                out24_: int
                out25_: bool
                out22_, out23_, out24_, out25_ = default__.m0(d_366_v214_, _dafny.Map({d_94_v13_: default__.fm5(d_77_v1_, False, (d_353_v202_)[default__.safeIndex(935, (d_353_v202_).length(0))], d_77_v1_, d_78_globalState_)}), (d_353_v202_)[default__.safeIndex(935, (d_353_v202_).length(0))], d_77_v1_, d_78_globalState_)
                d_367_v215_ = out22_
                d_368_v216_ = out23_
                d_369_v217_ = out24_
                d_370_v218_ = out25_
            elif True:
                d_80_v3_ = d_80_v3_
                d_371_v219_: _dafny.Array
                nw64_ = _dafny.Array(False, 19)
                d_371_v219_ = nw64_
                d_372_v220_: T0
                nw65_ = C2()
                nw65_.ctor__(53)
                d_372_v220_ = nw65_
                d_373_v221_: D6
                d_373_v221_ = D6_DC18(d_77_v1_, d_372_v220_, D5_DC15(d_76_v0_), d_77_v1_)
                d_374_v222_: D6
                d_374_v222_ = D6_DC19(d_373_v221_)
                d_375_v223_: _dafny.Set
                d_375_v223_ = _dafny.Set({d_374_v222_})
                d_376_v224_: _dafny.Map
                d_376_v224_ = _dafny.Map({d_375_v223_: d_76_v0_})
                index65_ = default__.safeIndex(397, (d_371_v219_).length(0))
                (d_371_v219_)[index65_] = (d_77_v1_) < (len(d_376_v224_))
                index66_ = default__.safeIndex(397, (d_371_v219_).length(0))
                (d_371_v219_)[index66_] = d_76_v0_
                d_77_v1_ = 742
                d_377_v225_: _dafny.Seq
                d_377_v225_ = _dafny.SeqWithoutIsStrInference([d_352_cf5_])
                d_377_v225_ = d_377_v225_
                d_378_v226_: _dafny.Map
                d_378_v226_ = _dafny.Map({d_76_v0_: True})
                d_379_v227_: _dafny.Map
                d_379_v227_ = _dafny.Map({d_77_v1_: d_378_v226_})
                d_380_v228_: _dafny.Seq
                d_380_v228_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_76_v0_: (d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))]}), d_378_v226_])
                d_381_v229_: C1
                nw66_ = C1()
                nw66_.ctor__()
                d_381_v229_ = nw66_
                d_382_v230_: _dafny.Seq
                d_382_v230_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))]: d_381_v229_})])
                d_383_v231_: _dafny.Array
                nw67_ = _dafny.Array(None, 29)
                nw67_[int(0)] = ((d_379_v227_)[d_77_v1_] if (d_77_v1_) in (d_379_v227_) else d_378_v226_)
                nw67_[int(1)] = _dafny.Map({not((d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))]): (d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))]})
                nw67_[int(2)] = (_dafny.Map({d_76_v0_: (d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))]})) | (default__.fm19((d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))], d_77_v1_, d_78_globalState_))
                nw67_[int(3)] = d_378_v226_
                nw67_[int(4)] = ((d_378_v226_).set((d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))], (d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))])) | (_dafny.Map({not((d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))]): (d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))]}))
                nw67_[int(5)] = (d_380_v228_)[default__.safeIndex(d_77_v1_, len(d_380_v228_))]
                nw67_[int(6)] = (_dafny.Map({(d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))]: d_76_v0_})) | (d_378_v226_)
                nw67_[int(7)] = d_378_v226_
                nw67_[int(8)] = default__.fm19(not(False), 747, d_78_globalState_)
                nw67_[int(9)] = d_378_v226_
                nw67_[int(10)] = (d_378_v226_) | (d_378_v226_)
                nw67_[int(11)] = d_378_v226_
                nw67_[int(12)] = d_378_v226_
                nw67_[int(13)] = _dafny.Map({(d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))]: (d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))]})
                nw67_[int(14)] = d_378_v226_
                nw67_[int(15)] = (d_380_v228_)[default__.safeIndex(len(d_382_v230_), len(d_380_v228_))]
                nw67_[int(16)] = d_378_v226_
                nw67_[int(17)] = (d_378_v226_) | (d_378_v226_)
                nw67_[int(18)] = (d_380_v228_)[default__.safeIndex(d_77_v1_, len(d_380_v228_))]
                nw67_[int(19)] = d_378_v226_
                nw67_[int(20)] = d_378_v226_
                nw67_[int(21)] = _dafny.Map({(d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))]: False})
                nw67_[int(22)] = default__.fm19(d_76_v0_, len(d_352_cf5_), d_78_globalState_)
                nw67_[int(23)] = _dafny.Map({d_76_v0_: d_76_v0_})
                nw67_[int(24)] = (d_378_v226_) | (_dafny.Map({False: ((d_310_v177_)[d_309_v176_] if (d_309_v176_) in (d_310_v177_) else (d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))])}))
                nw67_[int(25)] = d_378_v226_
                nw67_[int(26)] = (d_378_v226_) | (d_378_v226_)
                nw67_[int(27)] = _dafny.Map({False: (d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))]})
                nw67_[int(28)] = (d_378_v226_) | ((d_378_v226_).set((d_371_v219_)[default__.safeIndex(397, (d_371_v219_).length(0))], d_76_v0_))
                d_383_v231_ = nw67_
                d_383_v231_ = d_383_v231_
            d_384_v232_: _dafny.Map
            d_384_v232_ = _dafny.Map({d_76_v0_: d_76_v0_})
            d_385_v233_: T0
            nw68_ = C2()
            nw68_.ctor__((0) - ((0) - (len(d_352_cf5_))))
            d_385_v233_ = nw68_
            d_386_v234_: _dafny.Array
            def lambda32_(d_387_v0_):
                def lambda33_(d_388_i26_):
                    return not (d_387_v0_) or (d_387_v0_)

                return lambda33_

            init17_ = lambda32_(d_76_v0_)
            nw69_ = _dafny.Array(None, 2)
            for i0_17_ in range(nw69_.length(0)):
                nw69_[i0_17_] = init17_(i0_17_)
            d_386_v234_ = nw69_
            index67_ = default__.safeIndex(756, (d_386_v234_).length(0))
            (d_386_v234_)[index67_] = True
            d_389_v235_: _dafny.Array
            nw70_ = _dafny.Array(_dafny.Seq({}), 9)
            d_389_v235_ = nw70_
            d_390_v236_: C0
            nw71_ = C0()
            nw71_.ctor__(d_389_v235_)
            d_390_v236_ = nw71_
            d_391_v237_: _dafny.Set
            d_391_v237_ = _dafny.Set({d_390_v236_})
            d_392_v238_: _dafny.Seq
            d_392_v238_ = _dafny.SeqWithoutIsStrInference([d_94_v13_])
            d_393_v239_: _dafny.Seq
            d_393_v239_ = _dafny.SeqWithoutIsStrInference([d_77_v1_])
            index68_ = default__.safeIndex(756, (d_386_v234_).length(0))
            rhs62_ = (len((d_391_v237_) - (d_391_v237_))) + (len(d_352_cf5_))
            rhs63_ = default__.fm19(not(((d_392_v238_)[default__.safeIndex(d_77_v1_, len(d_392_v238_))]) == (d_94_v13_)), default__.safeDivisionInt(len(d_393_v239_), len(default__.fm20(d_77_v1_, d_76_v0_, d_76_v0_, d_77_v1_, d_78_globalState_))), d_78_globalState_)
            rhs64_ = d_385_v233_
            rhs65_ = d_76_v0_
            rhs66_ = d_76_v0_
            lhs35_ = d_386_v234_
            lhs36_ = default__.safeIndex(756, (d_386_v234_).length(0))
            d_77_v1_ = rhs62_
            d_384_v232_ = rhs63_
            d_385_v233_ = rhs64_
            lhs35_[lhs36_] = rhs65_
            d_76_v0_ = rhs66_
            d_394_v240_: _dafny.Array
            nw72_ = _dafny.Array(None, 8)
            d_394_v240_ = nw72_
            d_395_v241_: D7
            d_395_v241_ = D7_DC20(d_390_v236_)
            index69_ = default__.safeIndex(638, (d_394_v240_).length(0))
            (d_394_v240_)[index69_] = (d_395_v241_).cf29
            index70_ = default__.safeIndex(638, (d_394_v240_).length(0))
            nw73_ = C0()
            nw73_.ctor__(d_389_v235_)
            (d_394_v240_)[index70_] = nw73_
            d_396_v242_: _dafny.Array
            nw74_ = _dafny.Array(_dafny.Map({}), 24)
            d_396_v242_ = nw74_
            d_397_v243_: _dafny.Map
            d_397_v243_ = _dafny.Map({((d_384_v232_)[(d_386_v234_)[default__.safeIndex(756, (d_386_v234_).length(0))]] if ((d_386_v234_)[default__.safeIndex(756, (d_386_v234_).length(0))]) in (d_384_v232_) else False): d_386_v234_})
            index71_ = default__.safeIndex(357, (d_396_v242_).length(0))
            (d_396_v242_)[index71_] = (d_397_v243_) | (d_397_v243_)
            d_398_v244_: _dafny.Map
            d_398_v244_ = _dafny.Map({(d_386_v234_)[default__.safeIndex(756, (d_386_v234_).length(0))]: d_352_cf5_})
            d_399_v245_: _dafny.Array
            nw75_ = _dafny.Array(D5.default()(), 10)
            d_399_v245_ = nw75_
            d_400_v246_: D8
            d_400_v246_ = D8_DC24(d_399_v245_)
            d_401_v247_: _dafny.Map
            d_401_v247_ = _dafny.Map({d_76_v0_: d_399_v245_})
            pat_let_tv4_ = d_401_v247_
            pat_let_tv5_ = d_76_v0_
            pat_let_tv6_ = d_76_v0_
            pat_let_tv7_ = d_401_v247_
            pat_let_tv8_ = d_399_v245_
            index72_ = default__.safeIndex(357, (d_396_v242_).length(0))
            def iife8_(_pat_let2_0):
                def iife9_(d_402_dt__update__tmp_h1_):
                    def iife10_(_pat_let3_0):
                        def iife11_(d_403_dt__update_hcf34_h0_):
                            return D8_DC24(d_403_dt__update_hcf34_h0_)
                        return iife11_(_pat_let3_0)
                    return iife10_(((pat_let_tv4_)[pat_let_tv5_] if (pat_let_tv6_) in (pat_let_tv7_) else pat_let_tv8_))
                return iife9_(_pat_let2_0)
            rhs67_ = d_397_v243_
            rhs68_ = _dafny.Map({True: d_352_cf5_})
            rhs69_ = (iife8_(d_400_v246_)).cf34
            rhs70_ = (d_77_v1_) + (d_77_v1_)
            rhs71_ = (d_393_v239_) + (d_393_v239_)
            lhs37_ = d_396_v242_
            lhs38_ = default__.safeIndex(357, (d_396_v242_).length(0))
            lhs37_[lhs38_] = rhs67_
            d_398_v244_ = rhs68_
            d_399_v245_ = rhs69_
            d_77_v1_ = rhs70_
            d_393_v239_ = rhs71_
        elif True:
            d_404___mcc_h16_ = source7_.cf9
            d_405_cf9_ = d_404___mcc_h16_
            d_406_v248_: C2
            nw76_ = C2()
            nw76_.ctor__(d_77_v1_)
            d_406_v248_ = nw76_
            d_407_v249_: bool
            d_408_v250_: bool
            d_409_v251_: int
            d_410_v252_: bool
            out26_: bool
            out27_: bool
            out28_: int
            out29_: bool
            out26_, out27_, out28_, out29_ = default__.m0(_dafny.Map({d_76_v0_: d_309_v176_}), default__.fm21(d_77_v1_, default__.fm1(d_77_v1_, (0) - (d_77_v1_), len(_dafny.Map({d_77_v1_: d_406_v248_})), d_76_v0_, d_78_globalState_), d_78_globalState_), d_76_v0_, d_77_v1_, d_78_globalState_)
            d_407_v249_ = out26_
            d_408_v250_ = out27_
            d_409_v251_ = out28_
            d_410_v252_ = out29_
            d_411_v253_: C1
            nw77_ = C1()
            nw77_.ctor__()
            d_411_v253_ = nw77_
            d_412_v254_: D3
            d_412_v254_ = D3_DC8(d_410_v252_, d_150_v59_, True, _dafny.Set({d_76_v0_}), d_411_v253_)
            d_413_v255_: _dafny.Seq
            d_413_v255_ = _dafny.SeqWithoutIsStrInference([d_412_v254_])
            index73_ = default__.safeIndex(345, (d_91_v12_).length(0))
            (d_91_v12_)[index73_] = (0) - (len((_dafny.SeqWithoutIsStrInference([d_412_v254_, d_412_v254_, d_412_v254_, d_412_v254_])) + (d_413_v255_)))
            index74_ = default__.safeIndex(345, (d_91_v12_).length(0))
            (d_91_v12_)[index74_] = (0) - ((0) - ((624 if d_76_v0_ else (d_406_v248_).fm18(d_311_v178_, d_78_globalState_))))
            d_414_v256_: C2
            nw78_ = C2()
            nw78_.ctor__(((_dafny.MultiSet([d_408_v250_])).cardinality) + (d_77_v1_))
            d_414_v256_ = nw78_
            d_415_v257_: _dafny.Map
            d_415_v257_ = _dafny.Map({d_76_v0_: d_408_v250_})
            d_416_v258_: _dafny.Seq
            d_416_v258_ = _dafny.SeqWithoutIsStrInference([d_76_v0_, (d_76_v0_) and (((d_415_v257_)[d_410_v252_] if (d_410_v252_) in (d_415_v257_) else d_76_v0_)), d_408_v250_])
            d_416_v258_ = d_416_v258_
        if d_76_v0_:
            d_417_v259_: _dafny.Map
            d_417_v259_ = _dafny.Map({d_76_v0_: d_309_v176_})
            d_418_v260_: _dafny.Map
            d_418_v260_ = _dafny.Map({d_94_v13_: _dafny.Set({d_76_v0_})})
            d_419_v261_: _dafny.Seq
            d_419_v261_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klwl"))
            d_420_v262_: bool
            d_421_v263_: bool
            d_422_v264_: int
            d_423_v265_: bool
            out30_: bool
            out31_: bool
            out32_: int
            out33_: bool
            out30_, out31_, out32_, out33_ = default__.m0((d_417_v259_) | (d_417_v259_), d_418_v260_, d_76_v0_, len((d_419_v261_) + (d_419_v261_)), d_78_globalState_)
            d_420_v262_ = out30_
            d_421_v263_ = out31_
            d_422_v264_ = out32_
            d_423_v265_ = out33_
            d_424_v266_: _dafny.Map
            d_424_v266_ = _dafny.Map({d_77_v1_: 717})
            rhs72_ = d_421_v263_
            rhs73_ = len(d_424_v266_)
            rhs74_ = d_77_v1_
            rhs75_ = d_77_v1_
            d_423_v265_ = rhs72_
            d_422_v264_ = rhs73_
            d_422_v264_ = rhs74_
            d_77_v1_ = rhs75_
            d_425_v267_: T0
            nw79_ = C2()
            nw79_.ctor__(d_77_v1_)
            d_425_v267_ = nw79_
            d_426_v268_: _dafny.Seq
            d_426_v268_ = _dafny.SeqWithoutIsStrInference([d_425_v267_])
            d_427_v269_: C2
            nw80_ = C2()
            nw80_.ctor__((len(d_426_v268_)) * (d_77_v1_))
            d_427_v269_ = nw80_
            d_427_v269_ = d_427_v269_
            d_419_v261_ = d_419_v261_
            if (d_311_v178_).cf6:
                d_428_v270_: C2
                nw81_ = C2()
                nw81_.ctor__(d_77_v1_)
                d_428_v270_ = nw81_
                d_429_v271_: _dafny.Seq
                d_429_v271_ = _dafny.SeqWithoutIsStrInference([d_420_v262_, d_423_v265_, d_420_v262_])
                d_430_v272_: _dafny.MultiSet
                d_430_v272_ = _dafny.MultiSet([d_420_v262_])
                d_431_v273_: _dafny.Map
                d_431_v273_ = _dafny.Map({d_430_v272_: d_422_v264_})
                d_429_v271_ = _dafny.SeqWithoutIsStrInference([d_421_v263_, not((d_77_v1_) == ((0) - (len(d_431_v273_)))), d_421_v263_])
                d_420_v262_ = False
                d_432_v274_: C2
                nw82_ = C2()
                nw82_.ctor__(d_422_v264_)
                d_432_v274_ = nw82_
                d_433_v276_: _dafny.Map
                d_433_v276_ = _dafny.Map({d_94_v13_: True})
                d_434_v277_: bool
                d_435_v278_: bool
                d_436_v279_: int
                d_437_v280_: bool
                out34_: bool
                out35_: bool
                out36_: int
                out37_: bool
                def iife12_():
                    coll4_ = _dafny.Map()
                    compr_4_: _dafny.MultiSet
                    for compr_4_ in (d_433_v276_).keys.Elements:
                        d_438_v275_: _dafny.MultiSet = compr_4_
                        if (d_438_v275_) in (d_433_v276_):
                            coll4_[d_438_v275_] = _dafny.Set({d_423_v265_})
                    return _dafny.Map(coll4_)
                out34_, out35_, out36_, out37_ = default__.m0(d_417_v259_, iife12_()
                , d_423_v265_, (d_432_v274_).fm18(d_311_v178_, d_78_globalState_), d_78_globalState_)
                d_434_v277_ = out34_
                d_435_v278_ = out35_
                d_436_v279_ = out36_
                d_437_v280_ = out37_
            elif True:
                d_439_v281_: _dafny.Seq
                d_439_v281_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt((d_427_v269_).f3, (d_427_v269_).f3)])
                d_439_v281_ = d_439_v281_
                d_440_v282_: _dafny.Array
                nw83_ = _dafny.Array(None, 28)
                d_440_v282_ = nw83_
                d_440_v282_ = d_440_v282_
                d_441_v283_: bool
                d_442_v284_: bool
                d_443_v285_: int
                d_444_v286_: bool
                out38_: bool
                out39_: bool
                out40_: int
                out41_: bool
                out38_, out39_, out40_, out41_ = default__.m0((default__.fm4(d_422_v264_, d_421_v263_, d_78_globalState_)) | (d_417_v259_), d_418_v260_, d_421_v263_, (d_427_v269_).f3, d_78_globalState_)
                d_441_v283_ = out38_
                d_442_v284_ = out39_
                d_443_v285_ = out40_
                d_444_v286_ = out41_
                d_445_v287_: _dafny.Array
                nw84_ = _dafny.Array(False, 16)
                d_445_v287_ = nw84_
                d_446_v288_: _dafny.Map
                d_446_v288_ = _dafny.Map({(d_427_v269_).f3: d_445_v287_})
                d_447_v289_: _dafny.Array
                nw85_ = _dafny.Array(None, 17)
                nw85_[int(0)] = d_445_v287_
                nw85_[int(1)] = d_445_v287_
                nw85_[int(2)] = d_445_v287_
                nw85_[int(3)] = d_445_v287_
                nw85_[int(4)] = d_445_v287_
                nw85_[int(5)] = d_445_v287_
                nw85_[int(6)] = d_445_v287_
                nw85_[int(7)] = d_445_v287_
                nw85_[int(8)] = d_445_v287_
                nw85_[int(9)] = d_445_v287_
                nw85_[int(10)] = d_445_v287_
                nw85_[int(11)] = d_445_v287_
                nw85_[int(12)] = d_445_v287_
                nw85_[int(13)] = ((d_446_v288_)[(d_427_v269_).f3] if ((d_427_v269_).f3) in (d_446_v288_) else d_445_v287_)
                nw85_[int(14)] = d_445_v287_
                nw85_[int(15)] = d_445_v287_
                nw85_[int(16)] = d_445_v287_
                d_447_v289_ = nw85_
                d_448_v290_: _dafny.Map
                d_448_v290_ = _dafny.Map({d_447_v289_: d_422_v264_})
                d_449_v291_: D9
                d_449_v291_ = D9_DC26(d_439_v281_)
                d_448_v290_ = (d_448_v290_).set(d_447_v289_, (0) - ((0) - ((d_443_v285_) * ((_dafny.MultiSet((d_449_v291_).cf37)).cardinality))))
                d_450_v292_: _dafny.Map
                d_450_v292_ = _dafny.Map({d_421_v263_: d_419_v261_})
                d_451_v293_: _dafny.Map
                d_451_v293_ = _dafny.Map({d_425_v267_: (883) <= (812)})
                rhs76_ = (d_427_v269_).fm8(d_439_v281_, False, (d_427_v269_).f3, d_78_globalState_)
                rhs77_ = (d_450_v292_) | ((d_450_v292_) | (d_450_v292_))
                rhs78_ = (d_451_v293_) | ((d_451_v293_) | (d_451_v293_))
                rhs79_ = (0) - ((default__.fm13(d_78_globalState_)).cf7)
                d_442_v284_ = rhs76_
                d_450_v292_ = rhs77_
                d_451_v293_ = rhs78_
                d_443_v285_ = rhs79_
        elif True:
            d_452_v294_: _dafny.Seq
            d_452_v294_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqshqoby"))
            d_453_v295_: _dafny.Array
            nw86_ = _dafny.Array(None, 23)
            nw86_[int(0)] = (d_452_v294_) + (d_452_v294_)
            nw86_[int(1)] = d_452_v294_
            nw86_[int(2)] = (d_452_v294_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjnhobbkk")))
            nw86_[int(3)] = (d_452_v294_).set(default__.safeIndex(len(d_79_v2_), len(d_452_v294_)), d_309_v176_)
            nw86_[int(4)] = d_452_v294_
            nw86_[int(5)] = d_452_v294_
            nw86_[int(6)] = d_452_v294_
            nw86_[int(7)] = _dafny.SeqWithoutIsStrInference([(d_452_v294_)[default__.safeIndex(d_77_v1_, len(d_452_v294_))] for d_454_i27_ in range(default__.abs(455))])
            nw86_[int(8)] = d_452_v294_
            nw86_[int(9)] = (d_452_v294_) + (d_452_v294_)
            nw86_[int(10)] = d_452_v294_
            nw86_[int(11)] = d_452_v294_
            nw86_[int(12)] = (default__.fm12(d_78_globalState_)) + (d_452_v294_)
            nw86_[int(13)] = d_452_v294_
            nw86_[int(14)] = d_452_v294_
            nw86_[int(15)] = _dafny.SeqWithoutIsStrInference([d_309_v176_ for d_455_i28_ in range(default__.abs(184))])
            nw86_[int(16)] = (d_452_v294_ if d_76_v0_ else d_452_v294_)
            nw86_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifdyyjm"))
            nw86_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibqnypvji"))
            nw86_[int(19)] = d_452_v294_
            nw86_[int(20)] = d_452_v294_
            nw86_[int(21)] = d_452_v294_
            nw86_[int(22)] = d_452_v294_
            d_453_v295_ = nw86_
            index75_ = default__.safeIndex(920, (d_453_v295_).length(0))
            (d_453_v295_)[index75_] = d_452_v294_
            index76_ = default__.safeIndex(920, (d_453_v295_).length(0))
            (d_453_v295_)[index76_] = default__.fm12(d_78_globalState_)
            d_77_v1_ = (d_77_v1_ if d_76_v0_ else -181)
            d_456_v296_: _dafny.Seq
            d_456_v296_ = _dafny.SeqWithoutIsStrInference([119, 411])
            d_457_v297_: _dafny.Map
            d_457_v297_ = _dafny.Map({d_452_v294_: (d_456_v296_).set(default__.safeIndex(d_77_v1_, len(d_456_v296_)), 567)})
            d_457_v297_ = (d_457_v297_).set((d_453_v295_)[default__.safeIndex(920, (d_453_v295_).length(0))], d_456_v296_)
            d_458_v298_: _dafny.Array
            nw87_ = _dafny.Array(None, 9)
            nw87_[int(0)] = d_76_v0_
            nw87_[int(1)] = d_76_v0_
            nw87_[int(2)] = not(not(d_76_v0_))
            nw87_[int(3)] = d_76_v0_
            nw87_[int(4)] = d_76_v0_
            nw87_[int(5)] = d_76_v0_
            nw87_[int(6)] = not(d_76_v0_)
            nw87_[int(7)] = d_76_v0_
            nw87_[int(8)] = False
            d_458_v298_ = nw87_
            d_459_v299_: _dafny.Map
            d_459_v299_ = _dafny.Map({d_77_v1_: d_76_v0_})
            d_460_v300_: _dafny.Seq
            d_460_v300_ = _dafny.SeqWithoutIsStrInference([d_459_v299_, _dafny.Map({d_77_v1_: False})])
            index77_ = default__.safeIndex(667, (d_458_v298_).length(0))
            (d_458_v298_)[index77_] = (_dafny.SeqWithoutIsStrInference([d_459_v299_])) <= (d_460_v300_)
            d_461_v301_: D4
            d_461_v301_ = D4_DC11(d_453_v295_)
            d_462_v302_: _dafny.Map
            d_462_v302_ = _dafny.Map({d_76_v0_: d_76_v0_})
            d_463_v303_: _dafny.Seq
            d_463_v303_ = _dafny.SeqWithoutIsStrInference([((d_462_v302_)[d_76_v0_] if (d_76_v0_) in (d_462_v302_) else d_76_v0_), d_76_v0_])
            index78_ = default__.safeIndex(667, (d_458_v298_).length(0))
            rhs80_ = d_79_v2_
            rhs81_ = (d_463_v303_)[default__.safeIndex(d_77_v1_, len(d_463_v303_))]
            rhs82_ = d_76_v0_
            rhs83_ = D4_DC11(d_453_v295_)
            rhs84_ = d_76_v0_
            lhs39_ = d_458_v298_
            lhs40_ = default__.safeIndex(667, (d_458_v298_).length(0))
            d_79_v2_ = rhs80_
            d_76_v0_ = rhs81_
            lhs39_[lhs40_] = rhs82_
            d_461_v301_ = rhs83_
            d_76_v0_ = rhs84_
            d_91_v12_ = d_91_v12_
        d_464_v304_: _dafny.Seq
        d_464_v304_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmlrc"))
        d_465_v305_: _dafny.Seq
        d_465_v305_ = _dafny.SeqWithoutIsStrInference([False, d_76_v0_, d_76_v0_])
        d_466_v306_: C2
        nw88_ = C2()
        nw88_.ctor__(d_77_v1_)
        d_466_v306_ = nw88_
        d_467_v307_: _dafny.Set
        d_467_v307_ = _dafny.Set({d_466_v306_, d_466_v306_})
        d_468_v308_: _dafny.Map
        d_468_v308_ = _dafny.Map({d_76_v0_: d_76_v0_})
        d_469_v309_: _dafny.Seq
        d_469_v309_ = _dafny.SeqWithoutIsStrInference([899, d_77_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlokosuc")))])
        d_470_v310_: _dafny.Set
        d_470_v310_ = _dafny.Set({(d_469_v309_)[default__.safeIndex(d_77_v1_, len(d_469_v309_))], d_77_v1_})
        d_471_v311_: _dafny.Array
        nw89_ = _dafny.Array(None, 28)
        nw89_[int(0)] = d_76_v0_
        nw89_[int(1)] = d_76_v0_
        nw89_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obq"))) <= (d_464_v304_)
        nw89_[int(3)] = d_76_v0_
        nw89_[int(4)] = d_76_v0_
        nw89_[int(5)] = (d_76_v0_ if d_76_v0_ else d_76_v0_)
        nw89_[int(6)] = d_76_v0_
        nw89_[int(7)] = (d_76_v0_ if d_76_v0_ else d_76_v0_)
        nw89_[int(8)] = d_76_v0_
        nw89_[int(9)] = d_76_v0_
        nw89_[int(10)] = ((_dafny.MultiSet(d_465_v305_)).set(d_76_v0_, default__.abs(d_77_v1_))).isdisjoint(_dafny.MultiSet([default__.fm0(d_77_v1_, len(d_467_v307_), d_309_v176_, d_78_globalState_), ((d_468_v308_)[d_76_v0_] if (d_76_v0_) in (d_468_v308_) else not(not(d_76_v0_))), d_76_v0_]))
        nw89_[int(11)] = (d_79_v2_).isdisjoint(d_79_v2_)
        nw89_[int(12)] = d_76_v0_
        nw89_[int(13)] = d_76_v0_
        nw89_[int(14)] = d_76_v0_
        nw89_[int(15)] = (d_470_v310_).ispropersubset(d_470_v310_)
        nw89_[int(16)] = d_76_v0_
        nw89_[int(17)] = (d_77_v1_) < ((d_466_v306_).f3)
        nw89_[int(18)] = d_76_v0_
        nw89_[int(19)] = (not(d_76_v0_) if True else d_76_v0_)
        nw89_[int(20)] = d_76_v0_
        nw89_[int(21)] = True
        nw89_[int(22)] = d_76_v0_
        nw89_[int(23)] = d_76_v0_
        nw89_[int(24)] = d_76_v0_
        nw89_[int(25)] = d_76_v0_
        nw89_[int(26)] = d_76_v0_
        nw89_[int(27)] = (d_466_v306_).fm7((_dafny.SeqWithoutIsStrInference([_dafny.Map({D1_DC3(d_76_v0_, not(d_76_v0_), d_76_v0_): d_76_v0_}) for d_472_i29_ in range(default__.abs(171))])).set(default__.safeIndex(d_77_v1_, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({D1_DC3(d_76_v0_, not(d_76_v0_), d_76_v0_): d_76_v0_}) for d_473_i29_ in range(default__.abs(171))]))), _dafny.Map({D1_DC3(d_76_v0_, d_76_v0_, d_76_v0_): d_76_v0_})), d_76_v0_, len(d_465_v305_), d_78_globalState_)
        d_471_v311_ = nw89_
        d_471_v311_ = d_471_v311_
        d_474_v312_: _dafny.Seq
        d_474_v312_ = _dafny.SeqWithoutIsStrInference([d_464_v304_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")), _dafny.SeqWithoutIsStrInference([d_309_v176_ for d_475_i30_ in range(default__.abs(173))])])
        d_476_v313_: D6
        d_476_v313_ = D6_DC17((d_474_v312_) + (d_474_v312_))
        d_476_v313_ = d_476_v313_
        d_477_v314_: _dafny.Array
        nw90_ = _dafny.Array(D4.default()(), 7)
        d_477_v314_ = nw90_
        d_478_v315_: _dafny.Array
        def lambda34_(d_479_v309_):
            def lambda35_(d_480_i31_):
                return d_479_v309_

            return lambda35_

        init18_ = lambda34_(d_469_v309_)
        nw91_ = _dafny.Array(None, 22)
        for i0_18_ in range(nw91_.length(0)):
            nw91_[i0_18_] = init18_(i0_18_)
        d_478_v315_ = nw91_
        d_481_v316_: T1
        nw92_ = C0()
        nw92_.ctor__(d_478_v315_)
        d_481_v316_ = nw92_
        d_482_v317_: _dafny.Map
        d_482_v317_ = _dafny.Map({d_77_v1_: d_481_v316_})
        index79_ = default__.safeIndex(727, (d_477_v314_).length(0))
        (d_477_v314_)[index79_] = D4_DC10(((d_482_v317_)[d_77_v1_] if (d_77_v1_) in (d_482_v317_) else d_481_v316_))
        d_483_v318_: _dafny.Array
        nw93_ = _dafny.Array(None, 27)
        d_483_v318_ = nw93_
        d_484_v319_: _dafny.Map
        d_484_v319_ = _dafny.Map({d_483_v318_: d_91_v12_})
        d_485_v320_: D4
        d_485_v320_ = D4_DC10(d_481_v316_)
        d_486_v321_: _dafny.Map
        d_486_v321_ = _dafny.Map({d_469_v309_: (d_466_v306_).f3})
        index80_ = default__.safeIndex(727, (d_477_v314_).length(0))
        def iife13_():
            coll5_ = _dafny.Set()
            compr_5_: _dafny.Seq
            for compr_5_ in (d_486_v321_).keys.Elements:
                d_487_v322_: _dafny.Seq = compr_5_
                if (d_487_v322_) in (d_486_v321_):
                    coll5_ = coll5_.union(_dafny.Set([d_487_v322_]))
            return _dafny.Set(coll5_)
        rhs85_ = len(d_464_v304_)
        rhs86_ = d_309_v176_
        rhs87_ = d_485_v320_
        rhs88_ = (d_469_v309_) not in (iife13_()
        )
        rhs89_ = d_484_v319_
        lhs41_ = d_477_v314_
        lhs42_ = default__.safeIndex(727, (d_477_v314_).length(0))
        d_77_v1_ = rhs85_
        d_309_v176_ = rhs86_
        lhs41_[lhs42_] = rhs87_
        d_76_v0_ = rhs88_
        d_484_v319_ = rhs89_
        d_488_v324_: _dafny.Seq
        def iife14_():
            coll6_ = _dafny.Set()
            compr_6_: str
            for compr_6_ in (d_310_v177_).keys.Elements:
                d_489_v323_: str = compr_6_
                if (d_489_v323_) in (d_310_v177_):
                    coll6_ = coll6_.union(_dafny.Set([d_489_v323_]))
            return _dafny.Set(coll6_)
        d_488_v324_ = _dafny.SeqWithoutIsStrInference([iife14_()
        ])
        rhs90_ = d_471_v311_
        rhs91_ = d_488_v324_
        d_471_v311_ = rhs90_
        d_488_v324_ = rhs91_
        d_77_v1_ = (0) - (default__.safeDivisionInt(default__.safeModuloInt((d_466_v306_).f3, d_77_v1_), (542) - ((d_466_v306_).f3)))
        index81_ = default__.safeIndex(845, (d_471_v311_).length(0))
        (d_471_v311_)[index81_] = (d_77_v1_) != (d_77_v1_)
        index82_ = default__.safeIndex(845, (d_471_v311_).length(0))
        (d_471_v311_)[index82_] = d_76_v0_
        d_490_v325_: _dafny.MultiSet
        d_490_v325_ = _dafny.MultiSet([(d_471_v311_)[default__.safeIndex(845, (d_471_v311_).length(0))]])
        d_491_v326_: _dafny.Map
        d_491_v326_ = _dafny.Map({660: -463})
        d_492_v328_: _dafny.Seq
        def iife15_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in (d_470_v310_).Elements:
                d_493_v327_: int = compr_7_
                if (d_493_v327_) in (d_470_v310_):
                    coll7_[(d_493_v327_) - ((d_466_v306_).f3)] = (d_466_v306_).f3
            return _dafny.Map(coll7_)
        d_492_v328_ = _dafny.SeqWithoutIsStrInference([d_491_v326_, _dafny.Map({(d_466_v306_).f3: (d_466_v306_).f3}), d_491_v326_, _dafny.Map({-381: d_77_v1_}), iife15_()
        ])
        d_494_v329_: D2
        d_494_v329_ = D2_DC4(d_464_v304_)
        d_495_v330_: _dafny.Array
        nw94_ = _dafny.Array(None, 29)
        nw94_[int(0)] = (d_494_v329_).cf5
        nw94_[int(1)] = d_464_v304_
        nw94_[int(2)] = d_464_v304_
        nw94_[int(3)] = d_464_v304_
        nw94_[int(4)] = (d_464_v304_) + (d_464_v304_)
        nw94_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pajtyxdav"))
        nw94_[int(6)] = d_464_v304_
        nw94_[int(7)] = (d_464_v304_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btgoc")))
        nw94_[int(8)] = (d_464_v304_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvmwjuth")))
        nw94_[int(9)] = d_464_v304_
        nw94_[int(10)] = d_464_v304_
        nw94_[int(11)] = d_464_v304_
        nw94_[int(12)] = (d_464_v304_) + (d_464_v304_)
        nw94_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnb"))
        nw94_[int(14)] = d_464_v304_
        nw94_[int(15)] = d_464_v304_
        nw94_[int(16)] = (d_464_v304_ if (d_471_v311_)[default__.safeIndex(845, (d_471_v311_).length(0))] else _dafny.SeqWithoutIsStrInference([d_309_v176_ for d_496_i32_ in range(default__.abs(673))]))
        nw94_[int(17)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_497_i33_ in range(default__.abs(-614))])
        nw94_[int(18)] = _dafny.SeqWithoutIsStrInference([d_309_v176_ for d_498_i34_ in range(default__.abs(283))])
        nw94_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpv"))
        nw94_[int(20)] = _dafny.SeqWithoutIsStrInference([d_309_v176_ for d_499_i35_ in range(default__.abs(91))])
        nw94_[int(21)] = d_464_v304_
        nw94_[int(22)] = d_464_v304_
        nw94_[int(23)] = d_464_v304_
        nw94_[int(24)] = d_464_v304_
        nw94_[int(25)] = (d_464_v304_) + (d_464_v304_)
        nw94_[int(26)] = (d_464_v304_).set(default__.safeIndex(((d_94_v13_)[d_77_v1_] if (d_77_v1_) in (d_94_v13_) else (d_466_v306_).f3), len(d_464_v304_)), d_309_v176_)
        nw94_[int(27)] = (d_464_v304_) + (d_464_v304_)
        nw94_[int(28)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdd"))).set(default__.safeIndex(891, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdd")))), d_309_v176_)
        d_495_v330_ = nw94_
        d_500_v331_: _dafny.Array
        nw95_ = _dafny.Array(_dafny.Array(None, 0), 23)
        d_500_v331_ = nw95_
        index83_ = default__.safeIndex(981, (d_500_v331_).length(0))
        (d_500_v331_)[index83_] = d_471_v311_
        d_501_v333_: C0
        nw96_ = C0()
        nw96_.ctor__(d_478_v315_)
        d_501_v333_ = nw96_
        d_502_v334_: _dafny.Map
        d_502_v334_ = _dafny.Map({d_501_v333_: _dafny.Map({d_77_v1_: default__.fm1(d_77_v1_, d_77_v1_, 336, d_76_v0_, d_78_globalState_)})})
        d_503_v335_: _dafny.Map
        d_503_v335_ = _dafny.Map({(d_471_v311_)[default__.safeIndex(845, (d_471_v311_).length(0))]: ((d_502_v334_)[d_501_v333_] if (d_501_v333_) in (d_502_v334_) else d_491_v326_)})
        d_504_v336_: _dafny.Seq
        d_504_v336_ = _dafny.SeqWithoutIsStrInference([d_91_v12_])
        index84_ = default__.safeIndex(981, (d_500_v331_).length(0))
        def iife16_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(645, 744):
                d_505_v332_: int = compr_8_
                if ((645) <= (d_505_v332_)) and ((d_505_v332_) < (744)):
                    coll8_[(d_505_v332_) - ((d_466_v306_).f3)] = d_77_v1_
            return _dafny.Map(coll8_)
        rhs92_ = d_490_v325_
        rhs93_ = (_dafny.SeqWithoutIsStrInference([d_491_v326_, d_491_v326_, default__.fm22(d_77_v1_, d_76_v0_, d_78_globalState_), d_491_v326_, iife16_()
        ])) + (_dafny.SeqWithoutIsStrInference([d_491_v326_, ((d_503_v335_)[d_76_v0_] if (d_76_v0_) in (d_503_v335_) else d_491_v326_)]))
        rhs94_ = d_495_v330_
        rhs95_ = (d_504_v336_)[default__.safeIndex(d_77_v1_, len(d_504_v336_))]
        rhs96_ = d_471_v311_
        lhs43_ = d_500_v331_
        lhs44_ = default__.safeIndex(981, (d_500_v331_).length(0))
        d_490_v325_ = rhs92_
        d_492_v328_ = rhs93_
        d_495_v330_ = rhs94_
        d_91_v12_ = rhs95_
        lhs43_[lhs44_] = rhs96_
        d_77_v1_ = default__.safeModuloInt(d_77_v1_, d_77_v1_)
        _dafny.print(_dafny.string_of(d_76_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_77_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_globalState_.f0) == (_dafny.Map({False: 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_79_v2_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_80_v3_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_v12_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_v12_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_94_v13_) == (_dafny.MultiSet([-1, -1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_v57_) == (_dafny.Set({_dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v58_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v58_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v58_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_150_v59_).cf9).cf9).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_150_v59_).cf9).cf9).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_150_v59_).cf9).cf9).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_309_v176_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_310_v177_) == (_dafny.Map({_dafny.CodePoint('q'): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v178_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v178_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v178_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_464_v304_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_465_v305_) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_466_v306_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_467_v307_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_468_v308_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_469_v309_) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_470_v310_) == (_dafny.Set({899, -181}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_471_v311_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_474_v312_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmlrc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q')])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_476_v313_).cf23) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmlrc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q')]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmlrc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q')])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[0]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[1]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[2]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[3]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[4]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[5]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[6]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[7]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[8]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[9]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[10]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[11]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[12]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[13]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[14]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[15]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[16]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[17]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[18]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[19]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[20]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_478_v315_)[21]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_482_v317_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_484_v319_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_486_v321_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([899, -181, 8]): -181}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_488_v324_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.CodePoint('q')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_490_v325_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_491_v326_) == (_dafny.Map({660: -463}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_492_v328_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({660: -463}), _dafny.Map({660: -463}), _dafny.Map({2: 1, 734: 1}), _dafny.Map({660: -463}), _dafny.Map({826: 0, 827: 0, 828: 0, 829: 0, 830: 0, 831: 0, 832: 0, 833: 0, 834: 0, 835: 0, 836: 0, 837: 0, 838: 0, 839: 0, 840: 0, 841: 0, 842: 0, 843: 0, 844: 0, 845: 0, 846: 0, 847: 0, 848: 0, 849: 0, 850: 0, 851: 0, 852: 0, 853: 0, 854: 0, 855: 0, 856: 0, 857: 0, 858: 0, 859: 0, 860: 0, 861: 0, 862: 0, 863: 0, 864: 0, 865: 0, 866: 0, 867: 0, 868: 0, 869: 0, 870: 0, 871: 0, 872: 0, 873: 0, 874: 0, 875: 0, 876: 0, 877: 0, 878: 0, 879: 0, 880: 0, 881: 0, 882: 0, 883: 0, 884: 0, 885: 0, 886: 0, 887: 0, 888: 0, 889: 0, 890: 0, 891: 0, 892: 0, 893: 0, 894: 0, 895: 0, 896: 0, 897: 0, 898: 0, 899: 0, 900: 0, 901: 0, 902: 0, 903: 0, 904: 0, 905: 0, 906: 0, 907: 0, 908: 0, 909: 0, 910: 0, 911: 0, 912: 0, 913: 0, 914: 0, 915: 0, 916: 0, 917: 0, 918: 0, 919: 0, 920: 0, 921: 0, 922: 0, 923: 0, 924: 0}), _dafny.Map({660: -463}), _dafny.Map({0: 2})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_494_v329_).cf5).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_495_v330_)[16]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_495_v330_)[17]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_495_v330_)[18]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_495_v330_)[20]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[21]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[22]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[23]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[24]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[25]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[26]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[27]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_495_v330_)[28]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_500_v331_)[15])[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[0]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[1]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[2]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[3]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[4]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[5]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[6]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[7]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[8]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[9]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[10]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[11]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[12]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[13]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[14]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[15]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[16]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[17]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[18]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[19]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[20]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v333_.f2)[21]) == (_dafny.SeqWithoutIsStrInference([899, -181, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_502_v334_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_503_v335_) == (_dafny.Map({False: _dafny.Map({0: 2})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_504_v336_)))
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
        return lambda: D1_DC3(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC3(D1, NamedTuple('DC3', [('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
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
        return lambda: D2_DC5(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC5(D2, NamedTuple('DC5', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({self.cf5.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(False, D2.default()(), False, _dafny.Set({}), None)
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

class D3_DC8(D3, NamedTuple('DC8', [('cf11', Any), ('cf12', Any), ('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC11(D4, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)

class D5_DC14(D5, NamedTuple('DC14', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC18(int(0), None, D5.default()(), int(0))
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

class D6_DC18(D6, NamedTuple('DC18', [('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC21(_dafny.MultiSet({}), int(0), False)
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

class D7_DC21(D7, NamedTuple('DC21', [('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC22(D7, NamedTuple('DC22', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC23(D7, NamedTuple('DC23', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC23({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC23) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC25(_dafny.Map({}), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)

class D8_DC25(D8, NamedTuple('DC25', [('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC27()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)

class D9_DC27(D9, NamedTuple('DC27', [])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27)
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC29(int(0), int(0), False)
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

class D10_DC29(D10, NamedTuple('DC29', [('cf39', Any), ('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC30(D10, NamedTuple('DC30', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC30({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC30) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC32(False)
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

class D11_DC32(D11, NamedTuple('DC32', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC31(D11, NamedTuple('DC31', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC33(D11, NamedTuple('DC33', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC35(int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)

class D12_DC35(D12, NamedTuple('DC35', [('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35({_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35) and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC36(D12, NamedTuple('DC36', [])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36)
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass

class GlobalState:
    def  __init__(self):
        self.f0: _dafny.Map = _dafny.Map({})
        self._f1: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1):
        (self).f0 = f0
        (self)._f1 = f1

    @property
    def f1(self):
        return self._f1

class C0(T1, T0):
    def  __init__(self):
        self.f2: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f2):
        (self).f2 = f2

    def fm9(self, p0, p1, p2, globalState):
        return _dafny.Set({(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngubj")))) - (len(_dafny.Map({-318: False})))})

    def fm10(self, p0, p1, p2, p3, globalState):
        return True

    def fm7(self, p0, p1, p2, globalState):
        return not(((-806) + (781)) != (522))

    def fm8(self, p0, p1, p2, globalState):
        return ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmlbrpfx")))) * (len((D2_DC4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))).cf5))) in ((_dafny.SeqWithoutIsStrInference([685])) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality])))

    def fm11(self, p0, p1, p2, globalState):
        if (True) and (False):
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True]))) | (_dafny.MultiSet([not(not(True))]))
        elif True:
            return _dafny.MultiSet([True])


class C1:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def fm6(self, p0, p1, p2, p3, globalState):
        return (not((False if not(False) else True))) or (False)

    def m1(self, globalState):
        r0: D0 = D0.default()()
        r1: int = int(0)
        d_506_v0_: int
        d_506_v0_ = 786
        d_507_v1_: _dafny.MultiSet
        d_507_v1_ = _dafny.MultiSet([d_506_v0_])
        d_508_i0_: int
        d_508_i0_ = 0
        with _dafny.label("0"):
            while ((_dafny.MultiSet([d_506_v0_, d_506_v0_])) | (d_507_v1_)).issubset(d_507_v1_):
                with _dafny.c_label("0"):
                    if (d_508_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_508_i0_ = (d_508_i0_) + (1)
                    d_509_v2_: _dafny.Array
                    def lambda36_(d_510_v0_):
                        def lambda37_(d_511_i1_):
                            return _dafny.SeqWithoutIsStrInference([d_510_v0_, d_510_v0_])

                        return lambda37_

                    init19_ = lambda36_(d_506_v0_)
                    nw97_ = _dafny.Array(None, 28)
                    for i0_19_ in range(nw97_.length(0)):
                        nw97_[i0_19_] = init19_(i0_19_)
                    d_509_v2_ = nw97_
                    d_512_v3_: C0
                    nw98_ = C0()
                    nw98_.ctor__(d_509_v2_)
                    d_512_v3_ = nw98_
                    d_513_v4_: C0
                    nw99_ = C0()
                    nw99_.ctor__(d_512_v3_.f2)
                    d_513_v4_ = nw99_
                    r1 = d_506_v0_
                    d_514_v5_: bool
                    d_514_v5_ = True
                    d_515_v6_: _dafny.Array
                    nw100_ = _dafny.Array(None, 6)
                    nw100_[int(0)] = d_506_v0_
                    nw100_[int(1)] = d_506_v0_
                    nw100_[int(2)] = (0) - (d_506_v0_)
                    nw100_[int(3)] = len(default__.fm12(globalState))
                    nw100_[int(4)] = (d_506_v0_) - (len(_dafny.Set({d_514_v5_, d_514_v5_, d_514_v5_})))
                    nw100_[int(5)] = (0) - (default__.safeModuloInt((0) - (d_506_v0_), d_506_v0_))
                    d_515_v6_ = nw100_
                    d_515_v6_ = d_515_v6_
                    pass
            pass
        d_516_v7_: bool
        d_516_v7_ = True
        d_517_v8_: _dafny.Seq
        d_517_v8_ = _dafny.SeqWithoutIsStrInference([d_516_v7_])
        d_518_v9_: _dafny.Seq
        d_518_v9_ = _dafny.SeqWithoutIsStrInference([d_517_v8_, d_517_v8_, d_517_v8_, d_517_v8_, d_517_v8_])
        d_519_v10_: _dafny.Map
        d_519_v10_ = _dafny.Map({d_516_v7_: d_516_v7_})
        if (d_516_v7_) and (((_dafny.SeqWithoutIsStrInference([(0) - (d_506_v0_), -736, d_506_v0_, len(d_518_v9_), len(d_519_v10_)])).set(default__.safeIndex(d_506_v0_, len(_dafny.SeqWithoutIsStrInference([(0) - (d_506_v0_), -736, d_506_v0_, len(d_518_v9_), len(d_519_v10_)]))), d_506_v0_)) < (_dafny.SeqWithoutIsStrInference([d_506_v0_]))):
            d_520_v11_: _dafny.Array
            nw101_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_520_v11_ = nw101_
            d_521_v12_: _dafny.Seq
            d_521_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eelvyw"))
            d_522_v13_: _dafny.Seq
            d_522_v13_ = _dafny.SeqWithoutIsStrInference([d_521_v12_, d_521_v12_])
            d_523_v14_: _dafny.Map
            d_523_v14_ = _dafny.Map({(d_522_v13_)[default__.safeIndex(d_506_v0_, len(d_522_v13_))]: d_520_v11_})
            d_520_v11_ = (((d_523_v14_)[d_521_v12_] if (d_521_v12_) in (d_523_v14_) else d_520_v11_) if False else d_520_v11_)
            r1 = (d_506_v0_) + ((d_506_v0_) * (d_506_v0_))
            d_524_v15_: _dafny.Array
            nw102_ = _dafny.Array(_dafny.Seq({}), 1)
            d_524_v15_ = nw102_
            d_525_v16_: C0
            nw103_ = C0()
            nw103_.ctor__(d_524_v15_)
            d_525_v16_ = nw103_
            d_526_v17_: _dafny.Array
            def lambda38_(d_527_v0_, d_528_v7_):
                def lambda39_(d_529_i2_):
                    return _dafny.Map({d_527_v0_: d_528_v7_})

                return lambda39_

            init20_ = lambda38_(d_506_v0_, d_516_v7_)
            nw104_ = _dafny.Array(None, 10)
            for i0_20_ in range(nw104_.length(0)):
                nw104_[i0_20_] = init20_(i0_20_)
            d_526_v17_ = nw104_
            d_526_v17_ = d_526_v17_
            d_530_v18_: D2
            d_530_v18_ = D2_DC6(default__.fm13(globalState))
            source9_ = d_530_v18_
            if source9_.is_DC5:
                d_531___mcc_h0_ = source9_.cf6
                d_532___mcc_h1_ = source9_.cf7
                d_533___mcc_h2_ = source9_.cf8
                d_534_cf8_ = d_533___mcc_h2_
                d_535_cf7_ = d_532___mcc_h1_
                d_536_cf6_ = d_531___mcc_h0_
                d_520_v11_ = d_520_v11_
                d_537_v19_: _dafny.Array
                nw105_ = _dafny.Array(False, 2)
                d_537_v19_ = nw105_
                d_537_v19_ = d_537_v19_
                d_538_v20_: str
                d_538_v20_ = _dafny.CodePoint('i')
                d_539_v21_: _dafny.Map
                d_539_v21_ = _dafny.Map({d_536_cf6_: d_538_v20_})
                d_540_v22_: _dafny.Set
                d_540_v22_ = _dafny.Set({d_516_v7_})
                d_541_v23_: _dafny.Map
                d_541_v23_ = _dafny.Map({d_507_v1_: d_540_v22_})
                d_542_v24_: _dafny.Map
                d_542_v24_ = _dafny.Map({len(d_517_v8_): d_506_v0_})
                d_543_v25_: bool
                d_544_v26_: bool
                d_545_v27_: int
                d_546_v28_: bool
                out42_: bool
                out43_: bool
                out44_: int
                out45_: bool
                out42_, out43_, out44_, out45_ = default__.m0((d_539_v21_ if d_536_cf6_ else d_539_v21_), d_541_v23_, d_516_v7_, len(d_542_v24_), globalState)
                d_543_v25_ = out42_
                d_544_v26_ = out43_
                d_545_v27_ = out44_
                d_546_v28_ = out45_
                d_516_v7_ = default__.fm0(d_535_cf7_, (d_534_cf8_) + (d_535_cf7_), d_538_v20_, globalState)
            elif source9_.is_DC4:
                d_547___mcc_h3_ = source9_.cf5
                d_548_cf5_ = d_547___mcc_h3_
                d_549_v29_: _dafny.Set
                d_549_v29_ = _dafny.Set({(0) - (d_506_v0_)})
                d_550_v30_: _dafny.Seq
                d_550_v30_ = _dafny.SeqWithoutIsStrInference([d_549_v29_, d_549_v29_])
                d_551_v31_: _dafny.Seq
                d_551_v31_ = _dafny.SeqWithoutIsStrInference([d_506_v0_, d_506_v0_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_506_v0_, d_506_v0_])) for d_552_i3_ in range(default__.abs(-389))])), d_506_v0_, len(d_549_v29_)])
                d_553_v32_: str
                d_553_v32_ = _dafny.CodePoint('p')
                d_554_v33_: _dafny.Map
                d_554_v33_ = _dafny.Map({(self).fm6(d_553_v32_, d_516_v7_, len(_dafny.SeqWithoutIsStrInference([(d_521_v12_)[default__.safeIndex(len(d_521_v12_), len(d_521_v12_))] for d_555_i5_ in range(default__.abs(504))])), d_516_v7_, globalState): d_506_v0_})
                d_556_v34_: _dafny.Array
                nw106_ = _dafny.Array(None, 17)
                nw106_[int(0)] = (d_516_v7_ if d_516_v7_ else d_516_v7_)
                nw106_[int(1)] = ((d_550_v30_)[default__.safeIndex(len(d_551_v31_), len(d_550_v30_))]).issubset(d_549_v29_)
                nw106_[int(2)] = (len(_dafny.SeqWithoutIsStrInference([d_506_v0_ for d_557_i4_ in range(default__.abs(510))]))) <= (d_506_v0_)
                nw106_[int(3)] = not(d_516_v7_)
                nw106_[int(4)] = not((d_553_v32_) not in (_dafny.MultiSet([d_553_v32_])))
                nw106_[int(5)] = d_516_v7_
                nw106_[int(6)] = d_516_v7_
                nw106_[int(7)] = True
                nw106_[int(8)] = (d_516_v7_) not in (d_519_v10_)
                nw106_[int(9)] = d_516_v7_
                nw106_[int(10)] = d_516_v7_
                nw106_[int(11)] = d_516_v7_
                nw106_[int(12)] = d_516_v7_
                nw106_[int(13)] = d_516_v7_
                nw106_[int(14)] = (d_507_v1_).ispropersubset(d_507_v1_)
                nw106_[int(15)] = False
                nw106_[int(16)] = not(((d_525_v16_).fm9(d_516_v7_, d_554_v33_, False, globalState)) == (d_549_v29_))
                d_556_v34_ = nw106_
                index85_ = default__.safeIndex(727, (d_556_v34_).length(0))
                (d_556_v34_)[index85_] = d_516_v7_
                index86_ = default__.safeIndex(727, (d_556_v34_).length(0))
                (d_556_v34_)[index86_] = d_516_v7_
                d_558_v35_: _dafny.Array
                nw107_ = _dafny.Array(int(0), 29)
                d_558_v35_ = nw107_
                index87_ = default__.safeIndex(72, (d_558_v35_).length(0))
                (d_558_v35_)[index87_] = default__.safeModuloInt(d_506_v0_, d_506_v0_)
                d_559_v36_: _dafny.Map
                d_559_v36_ = _dafny.Map({d_506_v0_: (d_506_v0_) * ((0) - (len(d_521_v12_)))})
                index88_ = default__.safeIndex(403, (d_558_v35_).length(0))
                (d_558_v35_)[index88_] = 702
                d_560_v37_: _dafny.Map
                d_560_v37_ = _dafny.Map({default__.fm0(d_506_v0_, d_506_v0_, d_553_v32_, globalState): d_548_cf5_})
                index89_ = default__.safeIndex(72, (d_558_v35_).length(0))
                index90_ = default__.safeIndex(403, (d_558_v35_).length(0))
                index91_ = default__.safeIndex(727, (d_556_v34_).length(0))
                rhs97_ = (len(d_548_cf5_)) - (d_506_v0_)
                rhs98_ = ((d_559_v36_) | (d_559_v36_)) | (d_559_v36_)
                rhs99_ = d_506_v0_
                rhs100_ = ((d_516_v7_) or (d_516_v7_) if d_516_v7_ else not((d_549_v29_).ispropersubset(d_549_v29_)))
                rhs101_ = (((d_556_v34_)[default__.safeIndex(727, (d_556_v34_).length(0))]) and (d_516_v7_) if d_516_v7_ else (d_525_v16_).fm10(d_516_v7_, d_506_v0_, (d_525_v16_).fm8(d_551_v31_, d_516_v7_, d_506_v0_, globalState), len(d_560_v37_), globalState))
                lhs45_ = d_558_v35_
                lhs46_ = default__.safeIndex(72, (d_558_v35_).length(0))
                lhs47_ = d_558_v35_
                lhs48_ = default__.safeIndex(403, (d_558_v35_).length(0))
                lhs49_ = d_556_v34_
                lhs50_ = default__.safeIndex(727, (d_556_v34_).length(0))
                lhs45_[lhs46_] = rhs97_
                d_559_v36_ = rhs98_
                lhs47_[lhs48_] = rhs99_
                d_516_v7_ = rhs100_
                lhs49_[lhs50_] = rhs101_
                d_561_v38_: C0
                nw108_ = C0()
                nw108_.ctor__(d_524_v15_)
                d_561_v38_ = nw108_
                d_562_v39_: _dafny.Array
                def lambda40_(d_563_v0_, d_564_v8_, d_565_cf5_, d_566_v7_, d_567_v35_, d_568_v1_):
                    def lambda41_(d_569_i6_):
                        return (_dafny.MultiSet([d_563_v0_, (_dafny.MultiSet(d_564_v8_)).cardinality, len(d_565_cf5_), (D2_DC5(d_566_v7_, 468, (d_567_v35_)[default__.safeIndex(72, (d_567_v35_).length(0))])).cf8, (d_567_v35_)[default__.safeIndex(72, (d_567_v35_).length(0))]])).intersection(d_568_v1_)

                    return lambda41_

                init21_ = lambda40_(d_506_v0_, d_517_v8_, d_548_cf5_, d_516_v7_, d_558_v35_, d_507_v1_)
                nw109_ = _dafny.Array(None, 7)
                for i0_21_ in range(nw109_.length(0)):
                    nw109_[i0_21_] = init21_(i0_21_)
                d_562_v39_ = nw109_
                index92_ = default__.safeIndex(203, (d_562_v39_).length(0))
                (d_562_v39_)[index92_] = _dafny.MultiSet([(d_558_v35_)[default__.safeIndex(72, (d_558_v35_).length(0))]])
                index93_ = default__.safeIndex(203, (d_562_v39_).length(0))
                (d_562_v39_)[index93_] = (d_507_v1_ if (d_556_v34_)[default__.safeIndex(727, (d_556_v34_).length(0))] else (d_507_v1_) | (d_507_v1_))
            elif True:
                d_570___mcc_h4_ = source9_.cf9
                d_571_cf9_ = d_570___mcc_h4_
                d_572_v40_: _dafny.Map
                d_572_v40_ = _dafny.Map({default__.safeModuloInt(d_506_v0_, (0) - ((0) - (d_506_v0_))): d_506_v0_})
                d_572_v40_ = (d_572_v40_).set(d_506_v0_, d_506_v0_)
                rhs102_ = ((d_521_v12_) + (d_521_v12_)) + ((d_521_v12_) + (d_521_v12_))
                rhs103_ = d_506_v0_
                d_521_v12_ = rhs102_
                d_506_v0_ = rhs103_
                d_573_v41_: _dafny.Seq
                d_573_v41_ = _dafny.SeqWithoutIsStrInference([len(d_519_v10_), 288])
                d_574_v42_: _dafny.Array
                nw110_ = _dafny.Array(None, 7)
                nw110_[int(0)] = 858
                nw110_[int(1)] = len((d_573_v41_).set(default__.safeIndex(835, len(d_573_v41_)), d_506_v0_))
                nw110_[int(2)] = default__.fm1(d_506_v0_, d_506_v0_, 574, d_516_v7_, globalState)
                nw110_[int(3)] = d_506_v0_
                nw110_[int(4)] = len(d_521_v12_)
                nw110_[int(5)] = d_506_v0_
                nw110_[int(6)] = 351
                d_574_v42_ = nw110_
                d_575_v43_: str
                d_575_v43_ = _dafny.CodePoint('v')
                d_576_v44_: _dafny.Map
                d_576_v44_ = _dafny.Map({d_574_v42_: d_575_v43_})
                d_576_v44_ = (d_576_v44_).set(d_574_v42_, _dafny.CodePoint('u'))
                d_506_v0_ = d_506_v0_
        elif True:
            d_577_v45_: _dafny.Array
            nw111_ = _dafny.Array(_dafny.Seq({}), 25)
            d_577_v45_ = nw111_
            d_578_v46_: C0
            nw112_ = C0()
            nw112_.ctor__(d_577_v45_)
            d_578_v46_ = nw112_
            d_516_v7_ = d_516_v7_
            d_579_v47_: D0
            d_579_v47_ = D0_DC0(d_506_v0_)
            source10_ = d_579_v47_
            if source10_.is_DC1:
                d_580_v48_: _dafny.Array
                def lambda42_(d_581_v0_):
                    def lambda43_(d_582_i7_):
                        return (d_582_i7_) * (d_581_v0_)

                    return lambda43_

                init22_ = lambda42_(d_506_v0_)
                nw113_ = _dafny.Array(None, 28)
                for i0_22_ in range(nw113_.length(0)):
                    nw113_[i0_22_] = init22_(i0_22_)
                d_580_v48_ = nw113_
                d_583_v49_: _dafny.Map
                d_583_v49_ = _dafny.Map({d_516_v7_: d_580_v48_})
                d_584_v50_: str
                d_584_v50_ = _dafny.CodePoint('d')
                r1 = default__.safeDivisionInt(len((d_583_v49_) | ((d_583_v49_).set(default__.fm0(d_506_v0_, d_506_v0_, d_584_v50_, globalState), d_580_v48_))), d_506_v0_)
                d_516_v7_ = ((0) - (d_506_v0_)) > (d_506_v0_)
                d_585_v51_: _dafny.Array
                nw114_ = _dafny.Array(False, 21)
                d_585_v51_ = nw114_
                index94_ = default__.safeIndex(0, (d_585_v51_).length(0))
                (d_585_v51_)[index94_] = (d_517_v8_)[default__.safeIndex(d_506_v0_, len(d_517_v8_))]
                d_586_v52_: _dafny.Seq
                d_586_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "llor"))
                d_587_v53_: D2
                d_587_v53_ = D2_DC4(d_586_v52_)
                d_588_v54_: _dafny.Seq
                d_588_v54_ = _dafny.SeqWithoutIsStrInference([d_586_v52_])
                pat_let_tv9_ = d_588_v54_
                pat_let_tv10_ = d_506_v0_
                pat_let_tv11_ = d_588_v54_
                d_589_v55_: _dafny.Seq
                def iife17_(_pat_let4_0):
                    def iife18_(d_590_dt__update__tmp_h0_):
                        def iife19_(_pat_let5_0):
                            def iife20_(d_591_dt__update_hcf5_h0_):
                                return D2_DC4(d_591_dt__update_hcf5_h0_)
                            return iife20_(_pat_let5_0)
                        return iife19_((pat_let_tv9_)[default__.safeIndex(pat_let_tv10_, len(pat_let_tv11_))])
                    return iife18_(_pat_let4_0)
                d_589_v55_ = _dafny.SeqWithoutIsStrInference([len((iife17_(d_587_v53_)).cf5)])
                d_592_v56_: _dafny.Set
                d_592_v56_ = _dafny.Set({d_516_v7_})
                d_593_v57_: _dafny.MultiSet
                d_593_v57_ = _dafny.MultiSet([d_516_v7_])
                index95_ = default__.safeIndex(0, (d_585_v51_).length(0))
                rhs104_ = (len(d_589_v55_)) != ((0) - ((len(d_592_v56_)) * (d_506_v0_)))
                rhs105_ = (not(d_516_v7_)) or (((d_578_v46_).fm11(d_506_v0_, d_506_v0_, d_517_v8_, globalState)).ispropersubset(d_593_v57_))
                rhs106_ = default__.safeDivisionInt(15, 389)
                lhs51_ = d_585_v51_
                lhs52_ = default__.safeIndex(0, (d_585_v51_).length(0))
                lhs51_[lhs52_] = rhs104_
                d_516_v7_ = rhs105_
                r1 = rhs106_
                d_506_v0_ = d_506_v0_
            elif True:
                d_594___mcc_h5_ = source10_.cf0
                d_595_cf0_ = d_594___mcc_h5_
                d_596_v58_: _dafny.Seq
                d_596_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbodes"))
                d_597_v59_: D2
                d_597_v59_ = D2_DC6(D2_DC4(d_596_v58_))
                d_598_v60_: _dafny.Map
                d_598_v60_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_599_i8_ in range(default__.abs(892))])): d_597_v59_})
                d_600_v61_: D2
                d_600_v61_ = D2_DC6(((d_598_v60_)[d_506_v0_] if (d_506_v0_) in (d_598_v60_) else d_597_v59_))
                d_601_v62_: _dafny.Array
                nw115_ = _dafny.Array(_dafny.Set({}), 25)
                d_601_v62_ = nw115_
                d_602_v63_: _dafny.Map
                d_602_v63_ = _dafny.Map({d_506_v0_: d_595_cf0_})
                d_603_v64_: _dafny.Seq
                d_603_v64_ = _dafny.SeqWithoutIsStrInference([(0) - (((d_602_v63_)[295] if (295) in (d_602_v63_) else d_595_cf0_)), d_595_cf0_, default__.fm1(d_506_v0_, d_506_v0_, -203, d_516_v7_, globalState), d_506_v0_])
                d_604_v65_: _dafny.Map
                d_604_v65_ = _dafny.Map({(_dafny.MultiSet(d_603_v64_)).ispropersubset(d_507_v1_): d_506_v0_})
                rhs107_ = len(d_604_v65_)
                rhs108_ = D2_DC6(d_597_v59_)
                rhs109_ = d_601_v62_
                rhs110_ = ((d_516_v7_) and (d_516_v7_)) == (not((_dafny.Set({d_516_v7_})).issubset(_dafny.Set({False, d_516_v7_}))))
                r1 = rhs107_
                d_600_v61_ = rhs108_
                d_601_v62_ = rhs109_
                d_516_v7_ = rhs110_
                d_605_v66_: _dafny.Array
                nw116_ = _dafny.Array(int(0), 6)
                d_605_v66_ = nw116_
                d_606_v67_: _dafny.Set
                d_606_v67_ = _dafny.Set({d_516_v7_, (d_517_v8_)[default__.safeIndex((0) - (d_595_cf0_), len(d_517_v8_))], d_516_v7_, d_516_v7_})
                index96_ = default__.safeIndex(817, (d_605_v66_).length(0))
                (d_605_v66_)[index96_] = len(d_606_v67_)
                index97_ = default__.safeIndex(817, (d_605_v66_).length(0))
                (d_605_v66_)[index97_] = (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_607_i9_ in range(default__.abs(59))]))) - ((d_595_cf0_) + (d_595_cf0_))
                d_608_v68_: C0
                nw117_ = C0()
                nw117_.ctor__(d_578_v46_.f2)
                d_608_v68_ = nw117_
                d_609_v69_: D0
                d_609_v69_ = D0_DC1()
                d_610_v70_: _dafny.Array
                nw118_ = _dafny.Array(None, 11)
                nw118_[int(0)] = d_516_v7_
                nw118_[int(1)] = d_516_v7_
                nw118_[int(2)] = d_516_v7_
                nw118_[int(3)] = d_516_v7_
                nw118_[int(4)] = d_516_v7_
                nw118_[int(5)] = not(d_516_v7_)
                nw118_[int(6)] = d_516_v7_
                nw118_[int(7)] = d_516_v7_
                nw118_[int(8)] = d_516_v7_
                nw118_[int(9)] = d_516_v7_
                nw118_[int(10)] = d_516_v7_
                d_610_v70_ = nw118_
                d_611_v71_: _dafny.Map
                d_611_v71_ = _dafny.Map({d_609_v69_: d_610_v70_})
                d_611_v71_ = _dafny.Map({d_609_v69_: d_610_v70_})
            d_612_v72_: str
            d_612_v72_ = _dafny.CodePoint('v')
            d_612_v72_ = d_612_v72_
            d_613_v73_: C0
            nw119_ = C0()
            nw119_.ctor__(d_577_v45_)
            d_613_v73_ = nw119_
        d_614_v74_: _dafny.Array
        def lambda44_(d_615_v0_):
            def lambda45_(d_616_i10_):
                return (d_616_i10_) - (d_615_v0_)

            return lambda45_

        init23_ = lambda44_(d_506_v0_)
        nw120_ = _dafny.Array(None, 16)
        for i0_23_ in range(nw120_.length(0)):
            nw120_[i0_23_] = init23_(i0_23_)
        d_614_v74_ = nw120_
        d_617_v75_: _dafny.Map
        d_617_v75_ = _dafny.Map({d_516_v7_: d_506_v0_})
        index98_ = default__.safeIndex(964, (d_614_v74_).length(0))
        (d_614_v74_)[index98_] = default__.fm1((0) - (d_506_v0_), len(d_617_v75_), 846, d_516_v7_, globalState)
        d_618_v76_: _dafny.Seq
        d_618_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irlyr"))
        d_619_v77_: str
        d_619_v77_ = _dafny.CodePoint('r')
        d_620_v78_: D1
        d_620_v78_ = D1_DC2(d_619_v77_)
        pat_let_tv12_ = d_506_v0_
        pat_let_tv13_ = d_506_v0_
        index99_ = default__.safeIndex(964, (d_614_v74_).length(0))
        def lambda46_(source11_):
            if source11_.is_DC3:
                d_621___mcc_h6_ = source11_.cf2
                d_622___mcc_h7_ = source11_.cf3
                d_623___mcc_h8_ = source11_.cf4
                d_624_cf4_ = d_623___mcc_h8_
                d_625_cf3_ = d_622___mcc_h7_
                d_626_cf2_ = d_621___mcc_h6_
                return pat_let_tv12_
            elif True:
                d_627___mcc_h9_ = source11_.cf1
                d_628_cf1_ = d_627___mcc_h9_
                return pat_let_tv13_

        rhs111_ = d_506_v0_
        rhs112_ = (d_617_v75_ if d_516_v7_ else (d_617_v75_) | (d_617_v75_))
        rhs113_ = (0) - (len(d_618_v76_))
        rhs114_ = lambda46_(d_620_v78_)
        lhs53_ = d_614_v74_
        lhs54_ = default__.safeIndex(964, (d_614_v74_).length(0))
        lhs55_ = globalState
        lhs53_[lhs54_] = rhs111_
        lhs55_.f0 = rhs112_
        r1 = rhs113_
        r1 = rhs114_
        d_506_v0_ = 599
        d_629_v79_: _dafny.Array
        nw121_ = _dafny.Array(_dafny.CodePoint('D'), 18)
        d_629_v79_ = nw121_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_629_v79_).length(0)):
            d_630_i11_: int = guard_loop_1_
            if (True) and (((0) <= (d_630_i11_)) and ((d_630_i11_) < ((d_629_v79_).length(0)))):
                (d_629_v79_)[(d_630_i11_)] = d_619_v77_
        d_631_i12_: int
        d_631_i12_ = 0
        with _dafny.label("1"):
            while d_516_v7_:
                with _dafny.c_label("1"):
                    if (d_631_i12_) >= (100):
                        raise _dafny.Break("1")
                    d_631_i12_ = (d_631_i12_) + (1)
                    d_632_v80_: _dafny.Array
                    nw122_ = _dafny.Array(_dafny.Array(None, 0), 1)
                    d_632_v80_ = nw122_
                    d_633_v81_: _dafny.Array
                    nw123_ = _dafny.Array(_dafny.Set({}), 19)
                    d_633_v81_ = nw123_
                    index100_ = default__.safeIndex(53, (d_632_v80_).length(0))
                    (d_632_v80_)[index100_] = d_633_v81_
                    index101_ = default__.safeIndex(53, (d_632_v80_).length(0))
                    rhs115_ = d_633_v81_
                    rhs116_ = d_506_v0_
                    rhs117_ = d_516_v7_
                    lhs56_ = d_632_v80_
                    lhs57_ = default__.safeIndex(53, (d_632_v80_).length(0))
                    lhs56_[lhs57_] = rhs115_
                    d_506_v0_ = rhs116_
                    d_516_v7_ = rhs117_
                    d_634_v82_: _dafny.Map
                    d_634_v82_ = _dafny.Map({d_516_v7_: (d_618_v76_)[default__.safeIndex(d_506_v0_, len(d_618_v76_))]})
                    d_635_v83_: _dafny.Set
                    d_635_v83_ = _dafny.Set({d_516_v7_, d_516_v7_, d_516_v7_, d_516_v7_})
                    d_636_v84_: _dafny.Map
                    d_636_v84_ = _dafny.Map({_dafny.MultiSet([d_506_v0_]): d_635_v83_})
                    d_637_v85_: _dafny.Map
                    d_637_v85_ = _dafny.Map({d_506_v0_: d_506_v0_})
                    d_638_v86_: _dafny.Seq
                    d_638_v86_ = _dafny.SeqWithoutIsStrInference([-204])
                    d_639_v87_: _dafny.Seq
                    d_639_v87_ = _dafny.SeqWithoutIsStrInference([len(d_637_v85_), (d_638_v86_)[default__.safeIndex(d_506_v0_, len(d_638_v86_))]])
                    d_640_v88_: _dafny.MultiSet
                    d_640_v88_ = _dafny.MultiSet([d_516_v7_])
                    d_641_v89_: _dafny.Array
                    nw124_ = _dafny.Array(None, 13)
                    nw124_[int(0)] = d_639_v87_
                    nw124_[int(1)] = _dafny.SeqWithoutIsStrInference([len(d_638_v86_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbnakhyf")))])
                    nw124_[int(2)] = d_638_v86_
                    nw124_[int(3)] = d_638_v86_
                    nw124_[int(4)] = d_639_v87_
                    nw124_[int(5)] = d_638_v86_
                    nw124_[int(6)] = d_639_v87_
                    nw124_[int(7)] = (d_638_v86_).set(default__.safeIndex((d_614_v74_)[default__.safeIndex(964, (d_614_v74_).length(0))], len(d_638_v86_)), (d_640_v88_).cardinality)
                    nw124_[int(8)] = d_638_v86_
                    nw124_[int(9)] = d_638_v86_
                    nw124_[int(10)] = d_639_v87_
                    nw124_[int(11)] = d_638_v86_
                    nw124_[int(12)] = d_638_v86_
                    d_641_v89_ = nw124_
                    d_642_v90_: C0
                    nw125_ = C0()
                    nw125_.ctor__(d_641_v89_)
                    d_642_v90_ = nw125_
                    d_643_v91_: _dafny.Map
                    d_643_v91_ = _dafny.Map({d_642_v90_: d_516_v7_})
                    d_644_v92_: bool
                    d_645_v93_: bool
                    d_646_v94_: int
                    d_647_v95_: bool
                    out46_: bool
                    out47_: bool
                    out48_: int
                    out49_: bool
                    out46_, out47_, out48_, out49_ = default__.m0(d_634_v82_, d_636_v84_, ((d_643_v91_)[d_642_v90_] if (d_642_v90_) in (d_643_v91_) else not(d_516_v7_)), (d_614_v74_)[default__.safeIndex(964, (d_614_v74_).length(0))], globalState)
                    d_644_v92_ = out46_
                    d_645_v93_ = out47_
                    d_646_v94_ = out48_
                    d_647_v95_ = out49_
                    d_645_v93_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_648_i13_ in range(default__.abs(48))])) < (d_618_v76_)
                    d_649_v96_: _dafny.Map
                    d_649_v96_ = _dafny.Map({(d_614_v74_)[default__.safeIndex(964, (d_614_v74_).length(0))]: d_516_v7_})
                    d_650_v97_: _dafny.Map
                    d_650_v97_ = _dafny.Map({(d_638_v86_) + (_dafny.SeqWithoutIsStrInference([len(d_638_v86_), len(d_649_v96_), 878, d_646_v94_, 951])): (d_614_v74_)[default__.safeIndex(964, (d_614_v74_).length(0))]})
                    d_650_v97_ = (d_650_v97_).set((d_638_v86_) + (d_639_v87_), default__.fm1((d_614_v74_)[default__.safeIndex(964, (d_614_v74_).length(0))], (d_614_v74_)[default__.safeIndex(964, (d_614_v74_).length(0))], (d_614_v74_)[default__.safeIndex(964, (d_614_v74_).length(0))], d_516_v7_, globalState))
                    pass
            pass
        d_651_v98_: D0
        d_651_v98_ = D0_DC1()
        r0 = d_651_v98_
        d_652_v99_: _dafny.MultiSet
        d_652_v99_ = _dafny.MultiSet([d_516_v7_])
        d_653_v100_: _dafny.Seq
        d_653_v100_ = _dafny.SeqWithoutIsStrInference([default__.fm1(d_506_v0_, d_506_v0_, (d_614_v74_)[default__.safeIndex(964, (d_614_v74_).length(0))], d_516_v7_, globalState), d_506_v0_])
        d_654_v101_: _dafny.Map
        d_654_v101_ = _dafny.Map({(d_652_v99_).cardinality: d_653_v100_})
        r1 = len(((d_654_v101_)[(d_653_v100_)[default__.safeIndex(d_506_v0_, len(d_653_v100_))]] if ((d_653_v100_)[default__.safeIndex(d_506_v0_, len(d_653_v100_))]) in (d_654_v101_) else d_653_v100_))
        return r0, r1


class C2(T0):
    def  __init__(self):
        self._f3: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f3):
        (self)._f3 = f3

    def fm7(self, p0, p1, p2, globalState):
        return ((self).f3) == (327)

    def fm8(self, p0, p1, p2, globalState):
        return True

    def fm17(self, p0, p1, p2, p3, globalState):
        return ((self).f3) + ((0) - ((self).f3))

    def fm18(self, p0, globalState):
        return (self).f3

    @property
    def f3(self):
        return self._f3
