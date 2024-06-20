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
        return (-427) * (806)

    @staticmethod
    def fm1(p0, p1, globalState):
        return ((_dafny.MultiSet([False])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])))).isdisjoint(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False, not(True), True])))

    @staticmethod
    def fm2(p0, globalState):
        def iife0_():
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in (_dafny.SeqWithoutIsStrInference([(-901) for d_0_i0_ in range(default__.abs(423))])).Elements:
                    d_1_v0_: int = compr_2_
                    if (d_1_v0_) in (_dafny.SeqWithoutIsStrInference([(-901) for d_0_i0_ in range(default__.abs(423))])):
                        coll2_[default__.safeModuloInt(d_1_v0_, 949)] = True
                return _dafny.Map(coll2_)
            coll0_ = _dafny.Set()
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in (_dafny.SeqWithoutIsStrInference([(-901) for d_0_i0_ in range(default__.abs(423))])).Elements:
                    d_1_v0_: int = compr_1_
                    if (d_1_v0_) in (_dafny.SeqWithoutIsStrInference([(-901) for d_0_i0_ in range(default__.abs(423))])):
                        coll1_[default__.safeModuloInt(d_1_v0_, 949)] = True
                return _dafny.Map(coll1_)
            compr_0_: int
            for compr_0_ in (iife1_()
            ).keys.Elements:
                d_2_v1_: int = compr_0_
                if (d_2_v1_) in (iife2_()
                ):
                    coll0_ = coll0_.union(_dafny.Set([(d_2_v1_) - (len((D1_DC3(_dafny.SeqWithoutIsStrInference([57, (0) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))): True})))]), (0) - (len(_dafny.Map({not(True): True}))))).cf4))]))
            return _dafny.Set(coll0_)
        return _dafny.MultiSet([(366) < (len(_dafny.Map({False: -655}))), (252) != (907), not((_dafny.Set({424})).issubset(iife0_()
        )), (_dafny.CodePoint('y')) not in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k')])), (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcdmhndul")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cf")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdotcjy"))]))) in (_dafny.Map({403: 754}))])

    @staticmethod
    def fm5(p0, globalState):
        return D1_DC1(not(True))

    @staticmethod
    def fm6(globalState):
        return D1_DC3((_dafny.SeqWithoutIsStrInference([-626])) + (_dafny.SeqWithoutIsStrInference([-882, len(_dafny.Set({False})), len(_dafny.Map({True: _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "alqfd"))): not(not(not(True)))})})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdcr"))), (0) - ((0) - (-427))])), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vc")))) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikfvgwhey"))))))

    @staticmethod
    def fm7(globalState):
        return D1_DC2((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isjymm")))) + (687), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True}))]))).isdisjoint(_dafny.MultiSet([381, 900, len(_dafny.Set({not(False), False}))])))

    @staticmethod
    def fm8(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_3_i0_ in range(default__.abs(910))])

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        return (_dafny.Map({not(not(True)): False})) | ((D5_DC12(_dafny.Map({False: False}))).cf14)

    @staticmethod
    def fm10(globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: _dafny.Seq
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('k')])])).Elements:
                d_4_v0_: _dafny.Seq = compr_3_
                if (d_4_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('k')])])):
                    coll3_[d_4_v0_] = True
            return _dafny.Map(coll3_)
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: int
            for compr_4_ in (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "st")))]))): True})).keys.Elements:
                d_5_v1_: int = compr_4_
                if (d_5_v1_) in (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "st")))]))): True})):
                    coll4_ = coll4_.union(_dafny.Set([default__.safeModuloInt(d_5_v1_, (0) - (len(_dafny.Set({True}))))]))
            return _dafny.Set(coll4_)
        return ((_dafny.Set({679, 756, 9})) - (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apyq"))), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t')])), 272})), len(iife3_()
        )}))).intersection(iife4_()
        )

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('t')

    @staticmethod
    def fm12(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([364, len(_dafny.Set({True, not(True)})), (0) - ((_dafny.MultiSet([-803])).cardinality), -364])) + (_dafny.SeqWithoutIsStrInference([290, -495]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.CodePoint('f')}))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-600, len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([929])}))])), 616])))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        hi0_ = -648
        for d_6_i0_ in range(p1, hi0_):
            d_7_v0_: _dafny.Seq
            d_7_v0_ = _dafny.SeqWithoutIsStrInference([p1, default__.fm0(globalState), d_6_i0_])
            r2 = (d_7_v0_)[default__.safeIndex(d_6_i0_, len(d_7_v0_))]
            if (d_6_i0_) != (p1):
                d_8_v1_: str
                d_8_v1_ = _dafny.CodePoint('h')
                d_8_v1_ = d_8_v1_
                d_9_v2_: _dafny.Array
                nw0_ = _dafny.Array(_dafny.CodePoint('D'), 10)
                d_9_v2_ = nw0_
                d_10_v3_: _dafny.Array
                nw1_ = _dafny.Array(None, 13)
                nw1_[int(0)] = d_9_v2_
                nw1_[int(1)] = d_9_v2_
                nw1_[int(2)] = d_9_v2_
                nw1_[int(3)] = d_9_v2_
                nw1_[int(4)] = d_9_v2_
                nw1_[int(5)] = d_9_v2_
                nw1_[int(6)] = d_9_v2_
                nw1_[int(7)] = d_9_v2_
                nw1_[int(8)] = d_9_v2_
                nw1_[int(9)] = d_9_v2_
                nw1_[int(10)] = d_9_v2_
                nw1_[int(11)] = d_9_v2_
                nw1_[int(12)] = d_9_v2_
                d_10_v3_ = nw1_
                d_11_v4_: bool
                d_11_v4_ = False
                d_12_v5_: _dafny.Map
                d_12_v5_ = _dafny.Map({d_10_v3_: d_11_v4_})
                d_12_v5_ = (d_12_v5_).set(d_10_v3_, d_11_v4_)
                r0 = d_11_v4_
                d_13_v6_: _dafny.Array
                def lambda0_(d_14_v4_):
                    def lambda1_(d_15_i1_):
                        return (d_15_i1_) - ((0) - (len(_dafny.Map({d_14_v4_: d_14_v4_}))))

                    return lambda1_

                init0_ = lambda0_(d_11_v4_)
                nw2_ = _dafny.Array(None, 17)
                for i0_0_ in range(nw2_.length(0)):
                    nw2_[i0_0_] = init0_(i0_0_)
                d_13_v6_ = nw2_
                index0_ = default__.safeIndex(462, (d_13_v6_).length(0))
                (d_13_v6_)[index0_] = (0) - (-266)
                d_16_v7_: _dafny.Seq
                d_16_v7_ = _dafny.SeqWithoutIsStrInference([d_11_v4_])
                d_17_v8_: _dafny.MultiSet
                d_17_v8_ = _dafny.MultiSet([(0) - (p1)])
                d_18_v9_: _dafny.Map
                d_18_v9_ = _dafny.Map({d_13_v6_: d_17_v8_})
                d_19_v10_: _dafny.Seq
                d_19_v10_ = _dafny.SeqWithoutIsStrInference([d_11_v4_, (d_11_v4_) == (d_11_v4_), (_dafny.MultiSet([len(d_16_v7_)])).isdisjoint(((d_18_v9_)[d_13_v6_] if (d_13_v6_) in (d_18_v9_) else (d_17_v8_).set(d_6_i0_, default__.abs(d_6_i0_)))), d_11_v4_])
                index1_ = default__.safeIndex(462, (d_13_v6_).length(0))
                (d_13_v6_)[index1_] = len(d_19_v10_)
                d_20_v11_: _dafny.Map
                d_20_v11_ = _dafny.Map({d_11_v4_: d_17_v8_})
                d_20_v11_ = (d_20_v11_).set(d_11_v4_, d_17_v8_)
            elif True:
                d_21_v12_: _dafny.MultiSet
                d_21_v12_ = _dafny.MultiSet([p1, d_6_i0_, d_6_i0_])
                d_22_v13_: _dafny.Map
                d_22_v13_ = _dafny.Map({p1: d_21_v12_})
                d_23_v14_: str
                d_23_v14_ = _dafny.CodePoint('f')
                d_24_v15_: _dafny.MultiSet
                d_24_v15_ = _dafny.MultiSet([d_23_v14_, _dafny.CodePoint('a'), _dafny.CodePoint('x')])
                d_22_v13_ = (d_22_v13_).set(p1, _dafny.MultiSet([p1, ((d_24_v15_)[_dafny.CodePoint('b')] if (_dafny.CodePoint('b')) in (d_24_v15_) else d_6_i0_)]))
                d_25_v16_: _dafny.Array
                nw3_ = _dafny.Array(int(0), 21)
                d_25_v16_ = nw3_
                d_26_v17_: _dafny.MultiSet
                d_26_v17_ = _dafny.MultiSet([d_25_v16_])
                d_26_v17_ = (d_26_v17_) | (_dafny.MultiSet([d_25_v16_, d_25_v16_]))
                d_27_v18_: C0
                nw4_ = C0()
                nw4_.ctor__(_dafny.CodePoint('t'))
                d_27_v18_ = nw4_
                d_27_v18_ = d_27_v18_
                d_28_v19_: _dafny.Array
                nw5_ = _dafny.Array(False, 21)
                d_28_v19_ = nw5_
                d_29_v20_: bool
                d_29_v20_ = False
                d_30_v21_: _dafny.Map
                d_30_v21_ = _dafny.Map({d_27_v18_: not(d_29_v20_)})
                index2_ = default__.safeIndex(850, (d_28_v19_).length(0))
                (d_28_v19_)[index2_] = (d_27_v18_) not in (d_30_v21_)
                index3_ = default__.safeIndex(850, (d_28_v19_).length(0))
                (d_28_v19_)[index3_] = d_29_v20_
                d_31_v22_: _dafny.Array
                def lambda2_(d_32_v20_, d_33_v19_):
                    def lambda3_(d_34_i2_):
                        return (_dafny.Set({d_32_v20_})) | (_dafny.Set({(d_33_v19_)[default__.safeIndex(850, (d_33_v19_).length(0))], d_32_v20_, (d_33_v19_)[default__.safeIndex(850, (d_33_v19_).length(0))]}))

                    return lambda3_

                init1_ = lambda2_(d_29_v20_, d_28_v19_)
                nw6_ = _dafny.Array(None, 19)
                for i0_1_ in range(nw6_.length(0)):
                    nw6_[i0_1_] = init1_(i0_1_)
                d_31_v22_ = nw6_
                d_31_v22_ = d_31_v22_
            d_35_v23_: bool
            d_35_v23_ = False
            d_36_v24_: _dafny.Map
            d_36_v24_ = _dafny.Map({d_35_v23_: p1})
            d_37_v25_: _dafny.Seq
            d_37_v25_ = _dafny.SeqWithoutIsStrInference([d_35_v23_])
            d_38_v26_: _dafny.Map
            d_38_v26_ = _dafny.Map({d_6_i0_: -755})
            d_39_v27_: _dafny.Seq
            d_39_v27_ = _dafny.SeqWithoutIsStrInference([d_38_v26_, d_38_v26_])
            d_40_v28_: _dafny.Map
            d_40_v28_ = _dafny.Map({len(d_39_v27_): d_7_v0_})
            d_41_v29_: _dafny.Array
            nw7_ = _dafny.Array(None, 27)
            nw7_[int(0)] = p1
            nw7_[int(1)] = d_6_i0_
            nw7_[int(2)] = d_6_i0_
            nw7_[int(3)] = p1
            nw7_[int(4)] = d_6_i0_
            nw7_[int(5)] = ((d_36_v24_)[d_35_v23_] if (d_35_v23_) in (d_36_v24_) else d_6_i0_)
            nw7_[int(6)] = (len(_dafny.Set({p1}))) + (d_6_i0_)
            nw7_[int(7)] = (p1) - (206)
            nw7_[int(8)] = p1
            nw7_[int(9)] = (p1) - (p1)
            nw7_[int(10)] = len((d_37_v25_) + (d_37_v25_))
            nw7_[int(11)] = -159
            nw7_[int(12)] = (0) - (default__.safeDivisionInt(p1, p1))
            nw7_[int(13)] = d_6_i0_
            nw7_[int(14)] = len((d_40_v28_).set(p1, (d_7_v0_).set(default__.safeIndex(p1, len(d_7_v0_)), d_6_i0_)))
            nw7_[int(15)] = (750) - (d_6_i0_)
            nw7_[int(16)] = len(_dafny.SeqWithoutIsStrInference([p1, p1, (0) - ((0) - (p1))]))
            nw7_[int(17)] = d_6_i0_
            nw7_[int(18)] = (0) - (len(default__.fm9(d_35_v23_, d_35_v23_, d_35_v23_, globalState)))
            nw7_[int(19)] = p1
            nw7_[int(20)] = p1
            nw7_[int(21)] = (d_6_i0_) - (d_6_i0_)
            nw7_[int(22)] = (0) - (d_6_i0_)
            nw7_[int(23)] = (d_6_i0_) - (-539)
            nw7_[int(24)] = d_6_i0_
            nw7_[int(25)] = d_6_i0_
            nw7_[int(26)] = p1
            d_41_v29_ = nw7_
            index4_ = default__.safeIndex(937, (d_41_v29_).length(0))
            (d_41_v29_)[index4_] = d_6_i0_
            d_42_v30_: _dafny.Map
            d_42_v30_ = _dafny.Map({d_7_v0_: d_35_v23_})
            index5_ = default__.safeIndex(937, (d_41_v29_).length(0))
            rhs0_ = d_6_i0_
            rhs1_ = d_42_v30_
            lhs0_ = d_41_v29_
            lhs1_ = default__.safeIndex(937, (d_41_v29_).length(0))
            lhs0_[lhs1_] = rhs0_
            d_42_v30_ = rhs1_
            d_43_v31_: str
            d_43_v31_ = _dafny.CodePoint('t')
            d_44_v32_: T0
            nw8_ = C0()
            nw8_.ctor__(d_43_v31_)
            d_44_v32_ = nw8_
            d_45_v33_: _dafny.MultiSet
            d_45_v33_ = _dafny.MultiSet([d_44_v32_])
            d_46_v34_: _dafny.MultiSet
            d_46_v34_ = _dafny.MultiSet([(d_45_v33_).cardinality, len(d_7_v0_)])
            d_47_v35_: _dafny.Map
            d_47_v35_ = _dafny.Map({d_35_v23_: not((d_37_v25_)[default__.safeIndex((d_46_v34_).cardinality, len(d_37_v25_))])})
            d_47_v35_ = (d_47_v35_).set(d_35_v23_, d_35_v23_)
        d_48_v36_: str
        d_48_v36_ = _dafny.CodePoint('w')
        d_48_v36_ = d_48_v36_
        d_49_v37_: bool
        d_49_v37_ = True
        d_50_v38_: _dafny.Seq
        d_50_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejbib"))
        d_51_v39_: _dafny.MultiSet
        d_51_v39_ = _dafny.MultiSet([p1, len(d_50_v38_)])
        d_52_v40_: _dafny.Seq
        d_52_v40_ = _dafny.SeqWithoutIsStrInference([d_49_v37_])
        d_53_v41_: _dafny.Map
        d_53_v41_ = _dafny.Map({True: (d_52_v40_)[default__.safeIndex(p1, len(d_52_v40_))]})
        d_54_v42_: _dafny.Seq
        d_54_v42_ = _dafny.SeqWithoutIsStrInference([p1, len(_dafny.Set({d_49_v37_, d_49_v37_, default__.fm1(_dafny.SeqWithoutIsStrInference([d_51_v39_]), d_49_v37_, globalState)})), (p0).cardinality, p1, len(d_53_v41_)])
        d_55_v43_: _dafny.Array
        nw9_ = _dafny.Array(None, 13)
        nw9_[int(0)] = p1
        nw9_[int(1)] = p1
        nw9_[int(2)] = (p1) * (p1)
        nw9_[int(3)] = p1
        nw9_[int(4)] = len(default__.fm10(globalState))
        nw9_[int(5)] = p1
        nw9_[int(6)] = -831
        nw9_[int(7)] = (p1) + (p1)
        nw9_[int(8)] = p1
        nw9_[int(9)] = (d_54_v42_)[default__.safeIndex(p1, len(d_54_v42_))]
        nw9_[int(10)] = default__.safeModuloInt(p1, (0) - (len(d_53_v41_)))
        nw9_[int(11)] = p1
        nw9_[int(12)] = p1
        d_55_v43_ = nw9_
        index6_ = default__.safeIndex(784, (d_55_v43_).length(0))
        (d_55_v43_)[index6_] = (p1) * (p1)
        index7_ = default__.safeIndex(784, (d_55_v43_).length(0))
        (d_55_v43_)[index7_] = p1
        r2 = (d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))]
        d_56_v44_: _dafny.Array
        def lambda4_(d_57_v37_):
            def lambda5_(d_58_i3_):
                return d_57_v37_

            return lambda5_

        init2_ = lambda4_(d_49_v37_)
        nw10_ = _dafny.Array(None, 27)
        for i0_2_ in range(nw10_.length(0)):
            nw10_[i0_2_] = init2_(i0_2_)
        d_56_v44_ = nw10_
        index8_ = default__.safeIndex(561, (d_56_v44_).length(0))
        (d_56_v44_)[index8_] = d_49_v37_
        index9_ = default__.safeIndex(784, (d_55_v43_).length(0))
        index10_ = default__.safeIndex(561, (d_56_v44_).length(0))
        rhs2_ = default__.fm1(_dafny.SeqWithoutIsStrInference([d_51_v39_]), d_49_v37_, globalState)
        rhs3_ = d_52_v40_
        rhs4_ = p1
        rhs5_ = p1
        rhs6_ = d_49_v37_
        lhs2_ = globalState
        lhs3_ = d_55_v43_
        lhs4_ = default__.safeIndex(784, (d_55_v43_).length(0))
        lhs5_ = d_56_v44_
        lhs6_ = default__.safeIndex(561, (d_56_v44_).length(0))
        lhs2_.f0 = rhs2_
        d_52_v40_ = rhs3_
        r2 = rhs4_
        lhs3_[lhs4_] = rhs5_
        lhs5_[lhs6_] = rhs6_
        d_59_v45_: D1
        d_59_v45_ = D1_DC4()
        pat_let_tv0_ = d_55_v43_
        pat_let_tv1_ = d_55_v43_
        pat_let_tv2_ = d_51_v39_
        pat_let_tv3_ = d_51_v39_
        pat_let_tv4_ = d_55_v43_
        pat_let_tv5_ = d_55_v43_
        pat_let_tv6_ = d_49_v37_
        pat_let_tv7_ = d_49_v37_
        pat_let_tv8_ = d_49_v37_
        def lambda6_(source0_):
            if source0_.is_DC2:
                d_60___mcc_h0_ = source0_.cf2
                d_61___mcc_h1_ = source0_.cf3
                d_62_cf3_ = d_61___mcc_h1_
                d_63_cf2_ = d_60___mcc_h0_
                return d_62_cf3_
            elif source0_.is_DC3:
                d_64___mcc_h2_ = source0_.cf4
                d_65___mcc_h3_ = source0_.cf5
                d_66_cf5_ = d_65___mcc_h3_
                d_67_cf4_ = d_64___mcc_h2_
                return (_dafny.Set({(pat_let_tv1_)[default__.safeIndex(784, (pat_let_tv0_).length(0))]})).isdisjoint(_dafny.Set({(pat_let_tv2_).cardinality, (pat_let_tv3_).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sp"))), len(d_67_cf4_), (pat_let_tv5_)[default__.safeIndex(784, (pat_let_tv4_).length(0))]}))
            elif source0_.is_DC4:
                if True:
                    return (D1_DC1(pat_let_tv6_)).cf1
                elif True:
                    return pat_let_tv7_
            elif True:
                d_68___mcc_h4_ = source0_.cf1
                d_69_cf1_ = d_68___mcc_h4_
                return pat_let_tv8_

        if lambda6_(d_59_v45_):
            d_56_v44_ = d_56_v44_
            d_48_v36_ = default__.fm11((d_54_v42_) + (d_54_v42_), p1, 766, (d_49_v37_ if (d_56_v44_)[default__.safeIndex(561, (d_56_v44_).length(0))] else d_49_v37_), globalState)
            d_70_v46_: C0
            nw11_ = C0()
            nw11_.ctor__(d_48_v36_)
            d_70_v46_ = nw11_
            index11_ = default__.safeIndex(561, (d_56_v44_).length(0))
            (d_56_v44_)[index11_] = (d_49_v37_ if (d_70_v46_).fm3(globalState) else d_49_v37_)
            r1 = (p1) - (440)
        elif True:
            index12_ = default__.safeIndex(561, (d_56_v44_).length(0))
            (d_56_v44_)[index12_] = not((d_56_v44_)[default__.safeIndex(561, (d_56_v44_).length(0))])
            d_71_v47_: C0
            nw12_ = C0()
            nw12_.ctor__(d_48_v36_)
            d_71_v47_ = nw12_
            d_72_v48_: _dafny.Seq
            d_72_v48_ = _dafny.SeqWithoutIsStrInference([d_71_v47_])
            d_73_v49_: _dafny.Map
            d_73_v49_ = _dafny.Map({(d_72_v48_) + (_dafny.SeqWithoutIsStrInference([d_71_v47_, d_71_v47_])): (d_56_v44_)[default__.safeIndex(561, (d_56_v44_).length(0))]})
            d_73_v49_ = (d_73_v49_).set(d_72_v48_, False)
            d_49_v37_ = (p1) <= (((d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))]) * (p1))
            if d_49_v37_:
                index13_ = default__.safeIndex(561, (d_56_v44_).length(0))
                (d_56_v44_)[index13_] = (d_56_v44_)[default__.safeIndex(561, (d_56_v44_).length(0))]
                d_74_v50_: _dafny.Array
                nw13_ = _dafny.Array(None, 23)
                d_74_v50_ = nw13_
                d_75_v51_: T0
                nw14_ = C0()
                nw14_.ctor__(_dafny.CodePoint('q'))
                d_75_v51_ = nw14_
                index14_ = default__.safeIndex(611, (d_74_v50_).length(0))
                (d_74_v50_)[index14_] = d_75_v51_
                d_76_v52_: D2
                d_76_v52_ = D2_DC6(d_75_v51_)
                d_77_v53_: _dafny.Seq
                d_77_v53_ = _dafny.SeqWithoutIsStrInference([d_50_v38_])
                d_78_v54_: D4
                d_78_v54_ = D4_DC10((d_77_v53_)[default__.safeIndex((d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))], len(d_77_v53_))])
                d_79_v55_: D4
                d_79_v55_ = D4_DC10((d_78_v54_).cf11)
                d_80_v56_: _dafny.Map
                d_80_v56_ = _dafny.Map({d_49_v37_: (d_78_v54_).cf11})
                d_81_v57_: D1
                d_81_v57_ = D1_DC1((d_56_v44_)[default__.safeIndex(561, (d_56_v44_).length(0))])
                index15_ = default__.safeIndex(611, (d_74_v50_).length(0))
                index16_ = default__.safeIndex(561, (d_56_v44_).length(0))
                rhs7_ = (d_76_v52_).cf7
                rhs8_ = (((d_80_v56_)[d_49_v37_] if (d_49_v37_) in (d_80_v56_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcctiawd")))) < (default__.fm8(_dafny.SeqWithoutIsStrInference([p1, (d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))]]), globalState))
                rhs9_ = not((d_81_v57_).cf1)
                lhs7_ = d_74_v50_
                lhs8_ = default__.safeIndex(611, (d_74_v50_).length(0))
                lhs9_ = globalState
                lhs10_ = d_56_v44_
                lhs11_ = default__.safeIndex(561, (d_56_v44_).length(0))
                lhs7_[lhs8_] = rhs7_
                lhs9_.f0 = rhs8_
                lhs10_[lhs11_] = rhs9_
                d_82_v58_: _dafny.Map
                d_82_v58_ = _dafny.Map({d_59_v45_: (0) - ((d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))])})
                d_83_v59_: _dafny.Array
                nw15_ = _dafny.Array(None, 26)
                nw15_[int(0)] = (d_50_v38_) + (default__.fm8(_dafny.SeqWithoutIsStrInference([p1 for d_84_i4_ in range(default__.abs(79))]), globalState))
                nw15_[int(1)] = d_50_v38_
                nw15_[int(2)] = _dafny.SeqWithoutIsStrInference([d_48_v36_ for d_85_i5_ in range(default__.abs(684))])
                nw15_[int(3)] = (_dafny.SeqWithoutIsStrInference([(d_75_v51_).f3 for d_86_i6_ in range(default__.abs(21))])).set(default__.safeIndex(len(d_82_v58_), len(_dafny.SeqWithoutIsStrInference([(d_75_v51_).f3 for d_87_i6_ in range(default__.abs(21))]))), (d_75_v51_).f3)
                nw15_[int(4)] = (_dafny.SeqWithoutIsStrInference([(d_75_v51_).f3 for d_88_i7_ in range(default__.abs(103))])) + (d_50_v38_)
                nw15_[int(5)] = d_50_v38_
                nw15_[int(6)] = d_50_v38_
                nw15_[int(7)] = d_50_v38_
                nw15_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
                nw15_[int(9)] = (d_50_v38_).set(default__.safeIndex((d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))], len(d_50_v38_)), d_48_v36_)
                nw15_[int(10)] = d_50_v38_
                nw15_[int(11)] = d_50_v38_
                nw15_[int(12)] = (d_50_v38_).set(default__.safeIndex((d_54_v42_)[default__.safeIndex((d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))], len(d_54_v42_))], len(d_50_v38_)), (d_75_v51_).f3)
                nw15_[int(13)] = default__.fm8(d_54_v42_, globalState)
                nw15_[int(14)] = (d_50_v38_ if d_49_v37_ else default__.fm8(default__.fm12(default__.fm0(globalState), p1, globalState), globalState))
                nw15_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jumanw"))
                nw15_[int(16)] = d_50_v38_
                nw15_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
                nw15_[int(18)] = d_50_v38_
                nw15_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyeeyllc"))
                nw15_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
                nw15_[int(21)] = d_50_v38_
                nw15_[int(22)] = d_50_v38_
                nw15_[int(23)] = d_50_v38_
                nw15_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqetjr"))
                nw15_[int(25)] = (_dafny.SeqWithoutIsStrInference([(d_75_v51_).f3 for d_89_i8_ in range(default__.abs(80))])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_75_v51_).f3 for d_90_i9_ in range(default__.abs(392))])), len(_dafny.SeqWithoutIsStrInference([(d_75_v51_).f3 for d_91_i8_ in range(default__.abs(80))]))), d_48_v36_)
                d_83_v59_ = nw15_
                index17_ = default__.safeIndex(878, (d_83_v59_).length(0))
                (d_83_v59_)[index17_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqdbtkjg"))
                index18_ = default__.safeIndex(878, (d_83_v59_).length(0))
                (d_83_v59_)[index18_] = d_50_v38_
                r1 = (d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))]
                index19_ = default__.safeIndex(561, (d_56_v44_).length(0))
                (d_56_v44_)[index19_] = True
            elif True:
                r2 = len(((d_50_v38_) + (d_50_v38_)) + (d_50_v38_))
                d_92_v60_: _dafny.Array
                nw16_ = _dafny.Array(None, 16)
                nw16_[int(0)] = (d_71_v47_ if (d_56_v44_)[default__.safeIndex(561, (d_56_v44_).length(0))] else d_71_v47_)
                nw16_[int(1)] = d_71_v47_
                nw16_[int(2)] = d_71_v47_
                nw16_[int(3)] = d_71_v47_
                nw16_[int(4)] = d_71_v47_
                nw16_[int(5)] = d_71_v47_
                nw16_[int(6)] = (d_71_v47_ if not(d_49_v37_) else d_71_v47_)
                nw16_[int(7)] = d_71_v47_
                nw16_[int(8)] = d_71_v47_
                nw16_[int(9)] = d_71_v47_
                nw16_[int(10)] = d_71_v47_
                nw16_[int(11)] = d_71_v47_
                nw16_[int(12)] = d_71_v47_
                nw16_[int(13)] = d_71_v47_
                nw16_[int(14)] = d_71_v47_
                nw16_[int(15)] = d_71_v47_
                d_92_v60_ = nw16_
                d_92_v60_ = d_92_v60_
                d_50_v38_ = _dafny.SeqWithoutIsStrInference([d_48_v36_ for d_93_i10_ in range(default__.abs(128))])
                d_94_v61_: _dafny.Map
                d_94_v61_ = _dafny.Map({d_49_v37_: (d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))]})
                d_95_v62_: _dafny.Map
                d_95_v62_ = _dafny.Map({p1: d_55_v43_})
                r2 = len((d_94_v61_) | ((_dafny.Map({d_49_v37_: p1})).set(d_49_v37_, (0) - (len(_dafny.Map({d_56_v44_: len(d_95_v62_)}))))))
                r2 = (0) - (((d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))]) - ((d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))]))
        r0 = d_49_v37_
        r1 = default__.safeModuloInt(455, (d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))])
        r2 = (d_55_v43_)[default__.safeIndex(784, (d_55_v43_).length(0))]
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_96_globalState_: GlobalState
        nw17_ = GlobalState()
        nw17_.ctor__(True, True, False)
        d_96_globalState_ = nw17_
        d_97_v0_: _dafny.Seq
        d_97_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jract"))
        d_97_v0_ = (d_97_v0_) + (d_97_v0_)
        d_98_v1_: int
        d_98_v1_ = -27
        d_99_v2_: _dafny.Map
        d_99_v2_ = _dafny.Map({d_98_v1_: (0) - (d_98_v1_)})
        d_99_v2_ = (d_99_v2_).set(d_98_v1_, d_98_v1_)
        d_98_v1_ = d_98_v1_
        d_100_v3_: bool
        d_100_v3_ = True
        (d_96_globalState_).f0 = d_100_v3_
        (d_96_globalState_).f0 = True
        if d_100_v3_:
            d_101_v4_: _dafny.Array
            def lambda7_(d_102_i0_):
                return True

            init3_ = lambda7_
            nw18_ = _dafny.Array(None, 1)
            for i0_3_ in range(nw18_.length(0)):
                nw18_[i0_3_] = init3_(i0_3_)
            d_101_v4_ = nw18_
            index20_ = default__.safeIndex(784, (d_101_v4_).length(0))
            (d_101_v4_)[index20_] = d_100_v3_
            index21_ = default__.safeIndex(784, (d_101_v4_).length(0))
            (d_101_v4_)[index21_] = (d_98_v1_) >= ((default__.fm0(d_96_globalState_)) - (d_98_v1_))
            d_103_v5_: _dafny.Array
            def lambda8_(d_104_v3_):
                def lambda9_(d_105_i1_):
                    return _dafny.Map({d_104_v3_: not(d_104_v3_)})

                return lambda9_

            init4_ = lambda8_(d_100_v3_)
            nw19_ = _dafny.Array(None, 8)
            for i0_4_ in range(nw19_.length(0)):
                nw19_[i0_4_] = init4_(i0_4_)
            d_103_v5_ = nw19_
            d_106_v6_: _dafny.MultiSet
            d_106_v6_ = _dafny.MultiSet([d_98_v1_])
            d_107_v7_: _dafny.Seq
            d_107_v7_ = _dafny.SeqWithoutIsStrInference([d_106_v6_])
            d_108_v8_: _dafny.MultiSet
            d_108_v8_ = _dafny.MultiSet([(d_101_v4_)[default__.safeIndex(784, (d_101_v4_).length(0))]])
            d_109_v9_: _dafny.Map
            d_109_v9_ = _dafny.Map({(d_108_v8_).cardinality: (d_101_v4_)[default__.safeIndex(784, (d_101_v4_).length(0))]})
            d_110_v10_: _dafny.Map
            d_110_v10_ = _dafny.Map({d_103_v5_: (_dafny.Map({((d_109_v9_)[531] if (531) in (d_109_v9_) else d_100_v3_): d_108_v8_}) if default__.fm1(d_107_v7_, d_100_v3_, d_96_globalState_) else _dafny.Map({(d_101_v4_)[default__.safeIndex(784, (d_101_v4_).length(0))]: default__.fm2(d_98_v1_, d_96_globalState_)}))})
            d_111_v11_: _dafny.Map
            d_111_v11_ = _dafny.Map({True: d_108_v8_})
            d_110_v10_ = (d_110_v10_).set(d_103_v5_, d_111_v11_)
            d_112_v12_: _dafny.Seq
            d_112_v12_ = _dafny.SeqWithoutIsStrInference([d_98_v1_])
            d_113_v13_: _dafny.Map
            d_113_v13_ = _dafny.Map({d_112_v12_: d_98_v1_})
            d_114_v14_: _dafny.Set
            d_114_v14_ = _dafny.Set({d_98_v1_, (0) - (d_98_v1_), d_98_v1_})
            d_115_v15_: int
            d_115_v15_ = (0) - (len(d_114_v14_))
            d_116_v16_: _dafny.Map
            d_116_v16_ = _dafny.Map({d_98_v1_: d_109_v9_})
            d_117_v17_: _dafny.Map
            d_117_v17_ = _dafny.Map({False: (d_101_v4_)[default__.safeIndex(784, (d_101_v4_).length(0))]})
            d_118_v18_: _dafny.Array
            nw20_ = _dafny.Array(None, 17)
            nw20_[int(0)] = (len(d_97_v0_)) * (-941)
            nw20_[int(1)] = d_98_v1_
            nw20_[int(2)] = d_98_v1_
            nw20_[int(3)] = (d_98_v1_) * (d_98_v1_)
            nw20_[int(4)] = default__.safeModuloInt(len(d_113_v13_), d_98_v1_)
            nw20_[int(5)] = default__.safeDivisionInt(d_98_v1_, d_98_v1_)
            nw20_[int(6)] = (0) - (d_98_v1_)
            nw20_[int(7)] = 784
            nw20_[int(8)] = d_98_v1_
            nw20_[int(9)] = d_98_v1_
            nw20_[int(10)] = (d_115_v15_)
            nw20_[int(11)] = -619
            nw20_[int(12)] = d_98_v1_
            nw20_[int(13)] = (d_98_v1_) + (len(d_116_v16_))
            nw20_[int(14)] = (0) - (((d_106_v6_)[d_98_v1_] if (d_98_v1_) in (d_106_v6_) else d_98_v1_))
            nw20_[int(15)] = default__.safeModuloInt(d_98_v1_, len(d_117_v17_))
            nw20_[int(16)] = d_98_v1_
            d_118_v18_ = nw20_
            index22_ = default__.safeIndex(995, (d_118_v18_).length(0))
            (d_118_v18_)[index22_] = (d_98_v1_) - (888)
            index23_ = default__.safeIndex(995, (d_118_v18_).length(0))
            rhs10_ = ((d_98_v1_ if (d_101_v4_)[default__.safeIndex(784, (d_101_v4_).length(0))] else (d_106_v6_).cardinality)) * (d_98_v1_)
            rhs11_ = (d_117_v17_ if ((0) - (d_98_v1_)) <= (d_98_v1_) else (_dafny.Map({not((d_101_v4_)[default__.safeIndex(784, (d_101_v4_).length(0))]): (d_101_v4_)[default__.safeIndex(784, (d_101_v4_).length(0))]})).set((d_101_v4_)[default__.safeIndex(784, (d_101_v4_).length(0))], (d_101_v4_)[default__.safeIndex(784, (d_101_v4_).length(0))]))
            rhs12_ = (90) <= ((d_115_v15_))
            lhs12_ = d_118_v18_
            lhs13_ = default__.safeIndex(995, (d_118_v18_).length(0))
            lhs14_ = d_96_globalState_
            lhs12_[lhs13_] = rhs10_
            d_117_v17_ = rhs11_
            lhs14_.f0 = rhs12_
            index24_ = default__.safeIndex(995, (d_118_v18_).length(0))
            index25_ = default__.safeIndex(784, (d_101_v4_).length(0))
            rhs13_ = default__.fm0(d_96_globalState_)
            rhs14_ = (d_101_v4_)[default__.safeIndex(784, (d_101_v4_).length(0))]
            rhs15_ = default__.safeModuloInt(default__.safeModuloInt((d_118_v18_)[default__.safeIndex(995, (d_118_v18_).length(0))], (d_118_v18_)[default__.safeIndex(995, (d_118_v18_).length(0))]), 772)
            lhs15_ = d_118_v18_
            lhs16_ = default__.safeIndex(995, (d_118_v18_).length(0))
            lhs17_ = d_101_v4_
            lhs18_ = default__.safeIndex(784, (d_101_v4_).length(0))
            lhs15_[lhs16_] = rhs13_
            lhs17_[lhs18_] = rhs14_
            d_98_v1_ = rhs15_
            d_119_v19_: str
            d_119_v19_ = _dafny.CodePoint('c')
            d_120_v20_: C0
            nw21_ = C0()
            nw21_.ctor__(d_119_v19_)
            d_120_v20_ = nw21_
        elif True:
            d_121_v21_: _dafny.Array
            nw22_ = _dafny.Array(False, 3)
            d_121_v21_ = nw22_
            index26_ = default__.safeIndex(869, (d_121_v21_).length(0))
            (d_121_v21_)[index26_] = d_100_v3_
            index27_ = default__.safeIndex(869, (d_121_v21_).length(0))
            rhs16_ = d_100_v3_
            rhs17_ = default__.fm0(d_96_globalState_)
            lhs19_ = d_121_v21_
            lhs20_ = default__.safeIndex(869, (d_121_v21_).length(0))
            lhs19_[lhs20_] = rhs16_
            d_98_v1_ = rhs17_
            d_122_v22_: _dafny.Seq
            d_122_v22_ = _dafny.SeqWithoutIsStrInference([d_100_v3_])
            source1_ = default__.fm5(len(d_122_v22_), d_96_globalState_)
            if source1_.is_DC2:
                d_123___mcc_h0_ = source1_.cf2
                d_124___mcc_h1_ = source1_.cf3
                d_125_cf3_ = d_124___mcc_h1_
                d_126_cf2_ = d_123___mcc_h0_
                d_127_v23_: _dafny.Set
                d_127_v23_ = _dafny.Set({d_125_cf3_, not(False), d_125_cf3_})
                d_128_v24_: _dafny.Set
                d_128_v24_ = _dafny.Set({d_121_v21_, d_121_v21_})
                d_129_v25_: _dafny.MultiSet
                d_129_v25_ = _dafny.MultiSet([len(d_127_v23_), len(d_128_v24_)])
                d_125_cf3_ = ((d_121_v21_)[default__.safeIndex(869, (d_121_v21_).length(0))] if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_98_v1_, d_98_v1_]))).issubset(d_129_v25_) else d_125_cf3_)
                d_126_cf2_ = (d_98_v1_) * (d_126_cf2_)
                d_130_v26_: _dafny.Array
                nw23_ = _dafny.Array(int(0), 2)
                d_130_v26_ = nw23_
                d_131_v27_: D1
                d_131_v27_ = D1_DC1(d_125_cf3_)
                d_132_v28_: str
                d_132_v28_ = _dafny.CodePoint('o')
                d_133_v29_: C0
                nw24_ = C0()
                nw24_.ctor__(d_132_v28_)
                d_133_v29_ = nw24_
                d_134_v30_: _dafny.Map
                d_134_v30_ = _dafny.Map({d_131_v27_: d_133_v29_})
                index28_ = default__.safeIndex(967, (d_130_v26_).length(0))
                (d_130_v26_)[index28_] = len(d_134_v30_)
                index29_ = default__.safeIndex(967, (d_130_v26_).length(0))
                (d_130_v26_)[index29_] = default__.safeDivisionInt(d_98_v1_, d_126_cf2_)
                (d_96_globalState_).f0 = (d_121_v21_)[default__.safeIndex(869, (d_121_v21_).length(0))]
            elif source1_.is_DC3:
                d_135___mcc_h2_ = source1_.cf4
                d_136___mcc_h3_ = source1_.cf5
                d_137_cf5_ = d_136___mcc_h3_
                d_138_cf4_ = d_135___mcc_h2_
                (d_96_globalState_).f0 = d_100_v3_
                d_100_v3_ = not(d_100_v3_)
                d_98_v1_ = 988
                d_137_cf5_ = -855
            elif source1_.is_DC4:
                d_139_v31_: _dafny.MultiSet
                d_139_v31_ = _dafny.MultiSet([d_100_v3_])
                d_140_v32_: bool
                d_141_v33_: int
                d_142_v34_: int
                out0_: bool
                out1_: int
                out2_: int
                out0_, out1_, out2_ = default__.m0(d_139_v31_, d_98_v1_, d_96_globalState_)
                d_140_v32_ = out0_
                d_141_v33_ = out1_
                d_142_v34_ = out2_
                d_97_v0_ = (d_97_v0_) + (d_97_v0_)
                d_143_v35_: str
                d_143_v35_ = _dafny.CodePoint('v')
                d_144_v36_: C0
                nw25_ = C0()
                nw25_.ctor__(d_143_v35_)
                d_144_v36_ = nw25_
                d_145_v37_: _dafny.Set
                d_145_v37_ = _dafny.Set({d_142_v34_, 308})
                d_146_v38_: _dafny.MultiSet
                d_146_v38_ = _dafny.MultiSet([len(d_145_v37_)])
                d_147_v39_: _dafny.Seq
                d_147_v39_ = _dafny.SeqWithoutIsStrInference([d_146_v38_])
                d_148_v40_: _dafny.Set
                d_148_v40_ = _dafny.Set({(d_121_v21_)[default__.safeIndex(869, (d_121_v21_).length(0))], default__.fm1(d_147_v39_, (d_121_v21_)[default__.safeIndex(869, (d_121_v21_).length(0))], d_96_globalState_)})
                index30_ = default__.safeIndex(869, (d_121_v21_).length(0))
                rhs18_ = d_144_v36_
                rhs19_ = ((d_121_v21_)[default__.safeIndex(869, (d_121_v21_).length(0))]) == ((not(d_100_v3_)) not in (d_148_v40_))
                lhs21_ = d_121_v21_
                lhs22_ = default__.safeIndex(869, (d_121_v21_).length(0))
                d_144_v36_ = rhs18_
                lhs21_[lhs22_] = rhs19_
                d_149_v41_: _dafny.Array
                nw26_ = _dafny.Array(int(0), 26)
                d_149_v41_ = nw26_
                index31_ = default__.safeIndex(421, (d_149_v41_).length(0))
                (d_149_v41_)[index31_] = d_142_v34_
                index32_ = default__.safeIndex(421, (d_149_v41_).length(0))
                rhs20_ = d_143_v35_
                rhs21_ = d_141_v33_
                lhs23_ = d_149_v41_
                lhs24_ = default__.safeIndex(421, (d_149_v41_).length(0))
                d_143_v35_ = rhs20_
                lhs23_[lhs24_] = rhs21_
            elif True:
                d_150___mcc_h4_ = source1_.cf1
                d_151_cf1_ = d_150___mcc_h4_
                d_152_v42_: _dafny.Seq
                d_152_v42_ = _dafny.SeqWithoutIsStrInference([d_98_v1_, d_98_v1_, d_98_v1_, d_98_v1_, d_98_v1_])
                d_153_v43_: _dafny.Set
                d_153_v43_ = _dafny.Set({(d_152_v42_).set(default__.safeIndex(d_98_v1_, len(d_152_v42_)), d_98_v1_), d_152_v42_})
                (d_96_globalState_).f0 = (d_153_v43_) != (d_153_v43_)
                d_154_v44_: _dafny.Array
                nw27_ = _dafny.Array(_dafny.MultiSet({}), 29)
                d_154_v44_ = nw27_
                d_154_v44_ = (d_154_v44_ if (d_121_v21_)[default__.safeIndex(869, (d_121_v21_).length(0))] else d_154_v44_)
                d_155_v45_: _dafny.MultiSet
                d_155_v45_ = _dafny.MultiSet([d_100_v3_, d_151_cf1_])
                d_156_v46_: D2
                d_156_v46_ = D2_DC5(d_155_v45_)
                d_157_v47_: bool
                d_158_v48_: int
                d_159_v49_: int
                out3_: bool
                out4_: int
                out5_: int
                out3_, out4_, out5_ = default__.m0(((d_156_v46_).cf6) - (d_155_v45_), d_98_v1_, d_96_globalState_)
                d_157_v47_ = out3_
                d_158_v48_ = out4_
                d_159_v49_ = out5_
                d_160_v50_: D1
                d_160_v50_ = D1_DC2(d_98_v1_, d_151_cf1_)
                rhs22_ = d_121_v21_
                rhs23_ = d_160_v50_
                d_121_v21_ = rhs22_
                d_160_v50_ = rhs23_
            d_98_v1_ = (0) - (d_98_v1_)
            d_161_v51_: D1
            d_161_v51_ = D1_DC2(len(d_97_v0_), (d_121_v21_)[default__.safeIndex(869, (d_121_v21_).length(0))])
            d_162_v52_: _dafny.Map
            d_162_v52_ = _dafny.Map({d_98_v1_: _dafny.SeqWithoutIsStrInference([-327, (d_161_v51_).cf2])})
            d_163_v54_: _dafny.Map
            d_163_v54_ = _dafny.Map({_dafny.CodePoint('g'): d_98_v1_})
            d_164_v55_: _dafny.Seq
            d_164_v55_ = _dafny.SeqWithoutIsStrInference([d_122_v22_])
            d_165_v56_: _dafny.Seq
            d_165_v56_ = _dafny.SeqWithoutIsStrInference([len((d_164_v55_)[default__.safeIndex(len(d_97_v0_), len(d_164_v55_))]), d_98_v1_])
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: str
                for compr_5_ in (d_163_v54_).keys.Elements:
                    d_166_v53_: str = compr_5_
                    if (d_166_v53_) in (d_163_v54_):
                        coll5_[d_166_v53_] = d_100_v3_
                return _dafny.Map(coll5_)
            d_162_v52_ = (d_162_v52_).set(len(iife5_()
            ), d_165_v56_)
            d_167_v57_: _dafny.Map
            d_167_v57_ = _dafny.Map({d_98_v1_: (d_121_v21_)[default__.safeIndex(869, (d_121_v21_).length(0))]})
            d_168_v58_: _dafny.MultiSet
            d_168_v58_ = _dafny.MultiSet([d_100_v3_, (d_121_v21_)[default__.safeIndex(869, (d_121_v21_).length(0))], (d_121_v21_)[default__.safeIndex(869, (d_121_v21_).length(0))], ((d_167_v57_)[(default__.fm6(d_96_globalState_)).cf5] if ((default__.fm6(d_96_globalState_)).cf5) in (d_167_v57_) else d_100_v3_)])
            d_169_v59_: bool
            d_170_v60_: int
            d_171_v61_: int
            out6_: bool
            out7_: int
            out8_: int
            out6_, out7_, out8_ = default__.m0((d_168_v58_) | ((d_168_v58_).set(d_100_v3_, default__.abs(d_98_v1_))), default__.fm0(d_96_globalState_), d_96_globalState_)
            d_169_v59_ = out6_
            d_170_v60_ = out7_
            d_171_v61_ = out8_
        hi1_ = d_98_v1_
        for d_172_i2_ in range(d_98_v1_, hi1_):
            d_173_v62_: _dafny.MultiSet
            d_173_v62_ = _dafny.MultiSet([d_100_v3_])
            d_174_v63_: bool
            d_175_v64_: int
            d_176_v65_: int
            out9_: bool
            out10_: int
            out11_: int
            out9_, out10_, out11_ = default__.m0(d_173_v62_, d_98_v1_, d_96_globalState_)
            d_174_v63_ = out9_
            d_175_v64_ = out10_
            d_176_v65_ = out11_
            d_177_v66_: str
            d_177_v66_ = _dafny.CodePoint('i')
            d_178_v67_: T0
            nw28_ = C0()
            nw28_.ctor__(d_177_v66_)
            d_178_v67_ = nw28_
            d_179_v68_: _dafny.Map
            d_179_v68_ = _dafny.Map({D2_DC6(d_178_v67_): default__.fm7(d_96_globalState_)})
            d_180_v70_: _dafny.Seq
            d_180_v70_ = _dafny.SeqWithoutIsStrInference([d_174_v63_])
            d_181_v71_: _dafny.Array
            nw29_ = _dafny.Array(None, 21)
            nw29_[int(0)] = (d_173_v62_) - (_dafny.MultiSet([d_174_v63_]))
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: int
                for compr_6_ in (_dafny.Set({d_176_v65_})).Elements:
                    d_182_v69_: int = compr_6_
                    if (d_182_v69_) in (_dafny.Set({d_176_v65_})):
                        coll6_[(d_182_v69_) * (len(d_97_v0_))] = d_174_v63_
                return _dafny.Map(coll6_)
            nw29_[int(1)] = default__.fm2(len(iife6_()
            ), d_96_globalState_)
            nw29_[int(2)] = _dafny.MultiSet(d_180_v70_)
            nw29_[int(3)] = (d_173_v62_).set(d_100_v3_, default__.abs(d_175_v64_))
            nw29_[int(4)] = _dafny.MultiSet([d_174_v63_, d_100_v3_, d_100_v3_, d_100_v3_, d_174_v63_])
            nw29_[int(5)] = d_173_v62_
            nw29_[int(6)] = (default__.fm2(d_98_v1_, d_96_globalState_)).intersection(d_173_v62_)
            nw29_[int(7)] = d_173_v62_
            nw29_[int(8)] = (d_173_v62_) - (d_173_v62_)
            nw29_[int(9)] = d_173_v62_
            nw29_[int(10)] = default__.fm2((0) - (len(d_97_v0_)), d_96_globalState_)
            nw29_[int(11)] = _dafny.MultiSet(d_180_v70_)
            nw29_[int(12)] = d_173_v62_
            nw29_[int(13)] = d_173_v62_
            nw29_[int(14)] = (d_173_v62_) - (d_173_v62_)
            nw29_[int(15)] = d_173_v62_
            nw29_[int(16)] = d_173_v62_
            nw29_[int(17)] = d_173_v62_
            nw29_[int(18)] = _dafny.MultiSet([d_100_v3_])
            nw29_[int(19)] = d_173_v62_
            nw29_[int(20)] = d_173_v62_
            d_181_v71_ = nw29_
            index33_ = default__.safeIndex(673, (d_181_v71_).length(0))
            (d_181_v71_)[index33_] = d_173_v62_
            d_183_v72_: C0
            nw30_ = C0()
            nw30_.ctor__(d_177_v66_)
            d_183_v72_ = nw30_
            d_184_v73_: D2
            d_184_v73_ = D2_DC6(d_178_v67_)
            d_185_v74_: _dafny.Map
            d_185_v74_ = _dafny.Map({(0) - (d_175_v64_): d_183_v72_})
            d_186_v75_: D1
            d_186_v75_ = D1_DC2(len(d_185_v74_), True)
            d_187_v76_: _dafny.Set
            d_187_v76_ = _dafny.Set({d_175_v64_, d_98_v1_})
            index34_ = default__.safeIndex(673, (d_181_v71_).length(0))
            rhs24_ = (d_175_v64_) - (d_175_v64_)
            rhs25_ = (d_179_v68_) | (_dafny.Map({d_184_v73_: d_186_v75_}))
            rhs26_ = default__.fm2(d_175_v64_, d_96_globalState_)
            rhs27_ = (d_176_v65_) > ((0) - ((d_176_v65_) + (len(d_187_v76_))))
            rhs28_ = d_183_v72_
            lhs25_ = d_181_v71_
            lhs26_ = default__.safeIndex(673, (d_181_v71_).length(0))
            lhs27_ = d_96_globalState_
            d_98_v1_ = rhs24_
            d_179_v68_ = rhs25_
            lhs25_[lhs26_] = rhs26_
            lhs27_.f0 = rhs27_
            d_183_v72_ = rhs28_
            d_188_v77_: _dafny.Array
            nw31_ = _dafny.Array(False, 13)
            d_188_v77_ = nw31_
            index35_ = default__.safeIndex(393, (d_188_v77_).length(0))
            (d_188_v77_)[index35_] = not(d_100_v3_)
            d_189_v78_: _dafny.Seq
            d_189_v78_ = _dafny.SeqWithoutIsStrInference([d_98_v1_])
            d_190_v79_: D1
            d_190_v79_ = D1_DC3(d_189_v78_, (0) - (len(_dafny.SeqWithoutIsStrInference([d_183_v72_]))))
            d_191_v80_: _dafny.Map
            d_191_v80_ = _dafny.Map({d_100_v3_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqkshokld")))})
            d_192_v81_: _dafny.Map
            d_192_v81_ = _dafny.Map({d_178_v67_: d_100_v3_})
            d_193_v82_: _dafny.Map
            d_193_v82_ = _dafny.Map({d_178_v67_: d_98_v1_})
            index36_ = default__.safeIndex(393, (d_188_v77_).length(0))
            rhs29_ = (d_190_v79_).cf5
            rhs30_ = d_174_v63_
            rhs31_ = (_dafny.SeqWithoutIsStrInference([(d_180_v70_)[default__.safeIndex(((d_191_v80_)[d_174_v63_] if (d_174_v63_) in (d_191_v80_) else len(d_97_v0_)), len(d_180_v70_))], d_174_v63_])).set(default__.safeIndex((d_172_i2_ if d_174_v63_ else len(d_192_v81_)), len(_dafny.SeqWithoutIsStrInference([(d_180_v70_)[default__.safeIndex(((d_191_v80_)[d_174_v63_] if (d_174_v63_) in (d_191_v80_) else len(d_97_v0_)), len(d_180_v70_))], d_174_v63_]))), not(not((d_178_v67_) in (d_193_v82_))))
            rhs32_ = (d_175_v64_) != (282)
            lhs28_ = d_188_v77_
            lhs29_ = default__.safeIndex(393, (d_188_v77_).length(0))
            lhs30_ = d_96_globalState_
            d_176_v65_ = rhs29_
            lhs28_[lhs29_] = rhs30_
            d_180_v70_ = rhs31_
            lhs30_.f0 = rhs32_
            d_194_v83_: C0
            nw32_ = C0()
            nw32_.ctor__(d_177_v66_)
            d_194_v83_ = nw32_
        d_195_v84_: str
        d_195_v84_ = _dafny.CodePoint('m')
        d_196_v85_: _dafny.Array
        def lambda10_(d_197_v3_):
            def lambda11_(d_198_i3_):
                return not(d_197_v3_)

            return lambda11_

        init5_ = lambda10_(d_100_v3_)
        nw33_ = _dafny.Array(None, 6)
        for i0_5_ in range(nw33_.length(0)):
            nw33_[i0_5_] = init5_(i0_5_)
        d_196_v85_ = nw33_
        index37_ = default__.safeIndex(445, (d_196_v85_).length(0))
        (d_196_v85_)[index37_] = False
        d_199_v86_: _dafny.MultiSet
        d_199_v86_ = _dafny.MultiSet([d_100_v3_])
        d_200_v87_: _dafny.Set
        d_200_v87_ = _dafny.Set({((d_199_v86_)[d_100_v3_] if (d_100_v3_) in (d_199_v86_) else d_98_v1_)})
        d_201_v88_: _dafny.Seq
        d_201_v88_ = _dafny.SeqWithoutIsStrInference([d_200_v87_, d_200_v87_, d_200_v87_])
        d_202_v89_: _dafny.Set
        d_202_v89_ = _dafny.Set({d_100_v3_})
        index38_ = default__.safeIndex(445, (d_196_v85_).length(0))
        rhs33_ = d_97_v0_
        rhs34_ = not(d_100_v3_)
        rhs35_ = ((d_200_v87_).issubset((d_201_v88_)[default__.safeIndex(d_98_v1_, len(d_201_v88_))]) if d_100_v3_ else (d_202_v89_).ispropersubset(d_202_v89_))
        rhs36_ = d_195_v84_
        rhs37_ = not(not (d_100_v3_) or (d_100_v3_))
        lhs31_ = d_96_globalState_
        lhs32_ = d_196_v85_
        lhs33_ = default__.safeIndex(445, (d_196_v85_).length(0))
        d_97_v0_ = rhs33_
        d_100_v3_ = rhs34_
        lhs31_.f0 = rhs35_
        d_195_v84_ = rhs36_
        lhs32_[lhs33_] = rhs37_
        d_203_v90_: D1
        d_203_v90_ = D1_DC2(d_98_v1_, True)
        d_204_v91_: bool
        d_205_v92_: int
        d_206_v93_: int
        out12_: bool
        out13_: int
        out14_: int
        out12_, out13_, out14_ = default__.m0((default__.fm2(d_98_v1_, d_96_globalState_)).set((d_196_v85_)[default__.safeIndex(445, (d_196_v85_).length(0))], default__.abs(d_98_v1_)), (d_203_v90_).cf2, d_96_globalState_)
        d_204_v91_ = out12_
        d_205_v92_ = out13_
        d_206_v93_ = out14_
        if d_100_v3_:
            d_207_v94_: _dafny.Map
            d_207_v94_ = _dafny.Map({d_205_v92_: (d_196_v85_)[default__.safeIndex(445, (d_196_v85_).length(0))]})
            d_208_v95_: _dafny.Array
            def lambda12_(d_209_v93_):
                def lambda13_(d_210_i4_):
                    return (d_210_i4_) + (d_209_v93_)

                return lambda13_

            init6_ = lambda12_(d_206_v93_)
            nw34_ = _dafny.Array(None, 5)
            for i0_6_ in range(nw34_.length(0)):
                nw34_[i0_6_] = init6_(i0_6_)
            d_208_v95_ = nw34_
            d_211_v96_: _dafny.Map
            d_211_v96_ = _dafny.Map({d_208_v95_: d_196_v85_})
            d_212_v97_: _dafny.Seq
            d_212_v97_ = _dafny.SeqWithoutIsStrInference([d_204_v91_, not(d_204_v91_), ((d_207_v94_)[d_98_v1_] if (d_98_v1_) in (d_207_v94_) else (d_196_v85_)[default__.safeIndex(445, (d_196_v85_).length(0))]), (d_208_v95_) in (d_211_v96_), d_204_v91_])
            d_204_v91_ = not((d_212_v97_)[default__.safeIndex(d_206_v93_, len(d_212_v97_))])
            d_213_v98_: T0
            nw35_ = C0()
            nw35_.ctor__(d_195_v84_)
            d_213_v98_ = nw35_
            d_213_v98_ = d_213_v98_
            d_214_v99_: _dafny.Map
            d_214_v99_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(d_213_v98_).f3 for d_215_i5_ in range(default__.abs(925))])): d_195_v84_})
            d_216_v100_: _dafny.Map
            d_216_v100_ = _dafny.Map({d_100_v3_: len(d_214_v99_)})
            d_217_v101_: _dafny.Map
            d_217_v101_ = _dafny.Map({(d_212_v97_)[default__.safeIndex(513, len(d_212_v97_))]: d_100_v3_})
            d_218_v102_: _dafny.Map
            d_218_v102_ = _dafny.Map({d_206_v93_: d_216_v100_})
            d_219_v103_: _dafny.Seq
            d_219_v103_ = _dafny.SeqWithoutIsStrInference([d_206_v93_, len(d_216_v100_)])
            d_220_v104_: _dafny.Seq
            d_220_v104_ = _dafny.SeqWithoutIsStrInference([d_216_v100_, d_216_v100_])
            rhs38_ = not(((d_196_v85_)[default__.safeIndex(445, (d_196_v85_).length(0))] if (d_196_v85_)[default__.safeIndex(445, (d_196_v85_).length(0))] else not(((d_217_v101_)[(d_196_v85_)[default__.safeIndex(445, (d_196_v85_).length(0))]] if ((d_196_v85_)[default__.safeIndex(445, (d_196_v85_).length(0))]) in (d_217_v101_) else not((D1_DC2(d_98_v1_, d_204_v91_)).cf3)))))
            rhs39_ = d_195_v84_
            rhs40_ = ((((d_218_v102_)[(d_219_v103_)[default__.safeIndex(d_98_v1_, len(d_219_v103_))]] if ((d_219_v103_)[default__.safeIndex(d_98_v1_, len(d_219_v103_))]) in (d_218_v102_) else (d_220_v104_)[default__.safeIndex(d_206_v93_, len(d_220_v104_))])).set(d_204_v91_, d_205_v92_)) | (d_216_v100_)
            lhs34_ = d_96_globalState_
            lhs34_.f0 = rhs38_
            d_195_v84_ = rhs39_
            d_216_v100_ = rhs40_
            d_204_v91_ = d_100_v3_
            d_217_v101_ = (d_217_v101_).set(False, not((d_196_v85_)[default__.safeIndex(445, (d_196_v85_).length(0))]))
        elif True:
            d_221_v105_: bool
            d_222_v106_: int
            d_223_v107_: int
            out15_: bool
            out16_: int
            out17_: int
            out15_, out16_, out17_ = default__.m0(d_199_v86_, d_206_v93_, d_96_globalState_)
            d_221_v105_ = out15_
            d_222_v106_ = out16_
            d_223_v107_ = out17_
            index39_ = default__.safeIndex(445, (d_196_v85_).length(0))
            def iife7_():
                coll7_ = _dafny.Set()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(105, 36):
                    d_224_v108_: int = compr_7_
                    if ((105) <= (d_224_v108_)) and ((d_224_v108_) < (36)):
                        coll7_ = coll7_.union(_dafny.Set([default__.safeModuloInt(d_224_v108_, ((d_199_v86_)[d_221_v105_] if (d_221_v105_) in (d_199_v86_) else d_98_v1_))]))
                return _dafny.Set(coll7_)
            (d_196_v85_)[index39_] = (d_200_v87_).isdisjoint(iife7_()
            )
            d_225_v109_: _dafny.Seq
            d_225_v109_ = _dafny.SeqWithoutIsStrInference([(d_196_v85_)[default__.safeIndex(445, (d_196_v85_).length(0))]])
            d_225_v109_ = d_225_v109_
            d_226_v110_: _dafny.Array
            def lambda14_(d_227_v2_, d_228_v107_, d_229_v105_, d_230_v106_, d_231_v85_):
                def lambda15_(d_232_i6_):
                    return _dafny.MultiSet([_dafny.Map({((d_227_v2_)[354] if (354) in (d_227_v2_) else d_228_v107_): d_229_v105_}), _dafny.Map({d_230_v106_: (D1_DC1((d_231_v85_)[default__.safeIndex(445, (d_231_v85_).length(0))])).cf1})])

                return lambda15_

            init7_ = lambda14_(d_99_v2_, d_223_v107_, d_221_v105_, d_222_v106_, d_196_v85_)
            nw36_ = _dafny.Array(None, 26)
            for i0_7_ in range(nw36_.length(0)):
                nw36_[i0_7_] = init7_(i0_7_)
            d_226_v110_ = nw36_
            d_233_v111_: _dafny.Map
            d_233_v111_ = _dafny.Map({d_223_v107_: True})
            d_234_v112_: _dafny.MultiSet
            d_234_v112_ = _dafny.MultiSet([d_233_v111_, d_233_v111_])
            index40_ = default__.safeIndex(988, (d_226_v110_).length(0))
            (d_226_v110_)[index40_] = (d_234_v112_) | (d_234_v112_)
            d_235_v113_: _dafny.Array
            nw37_ = _dafny.Array(int(0), 9)
            d_235_v113_ = nw37_
            index41_ = default__.safeIndex(687, (d_235_v113_).length(0))
            (d_235_v113_)[index41_] = d_206_v93_
            d_236_v114_: _dafny.Seq
            d_236_v114_ = _dafny.SeqWithoutIsStrInference([d_206_v93_])
            index42_ = default__.safeIndex(496, (d_235_v113_).length(0))
            (d_235_v113_)[index42_] = (d_223_v107_ if d_100_v3_ else len(d_236_v114_))
            index43_ = default__.safeIndex(988, (d_226_v110_).length(0))
            index44_ = default__.safeIndex(687, (d_235_v113_).length(0))
            index45_ = default__.safeIndex(496, (d_235_v113_).length(0))
            def iife8_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in (d_233_v111_).keys.Elements:
                    d_237_v115_: int = compr_8_
                    if (d_237_v115_) in (d_233_v111_):
                        coll8_[(d_237_v115_) * (-496)] = True
                return _dafny.Map(coll8_)
            rhs41_ = (d_234_v112_).intersection(_dafny.MultiSet([iife8_()
            , d_233_v111_]))
            rhs42_ = not((d_199_v86_).isdisjoint((d_199_v86_).set(d_100_v3_, default__.abs(d_98_v1_))))
            rhs43_ = (default__.fm8(d_236_v114_, d_96_globalState_) if (d_196_v85_)[default__.safeIndex(445, (d_196_v85_).length(0))] else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))
            rhs44_ = default__.safeDivisionInt(d_98_v1_, default__.safeDivisionInt(d_223_v107_, (0) - (len(default__.fm8(d_236_v114_, d_96_globalState_)))))
            rhs45_ = (d_205_v92_) * (d_223_v107_)
            lhs35_ = d_226_v110_
            lhs36_ = default__.safeIndex(988, (d_226_v110_).length(0))
            lhs37_ = d_96_globalState_
            lhs38_ = d_235_v113_
            lhs39_ = default__.safeIndex(687, (d_235_v113_).length(0))
            lhs40_ = d_235_v113_
            lhs41_ = default__.safeIndex(496, (d_235_v113_).length(0))
            lhs35_[lhs36_] = rhs41_
            lhs37_.f0 = rhs42_
            d_97_v0_ = rhs43_
            lhs38_[lhs39_] = rhs44_
            lhs40_[lhs41_] = rhs45_
            if not(not (d_100_v3_) or (d_221_v105_)):
                d_238_v116_: _dafny.Map
                d_238_v116_ = _dafny.Map({d_233_v111_: d_199_v86_})
                d_239_v117_: D1
                d_239_v117_ = D1_DC1((d_196_v85_)[default__.safeIndex(445, (d_196_v85_).length(0))])
                d_240_v118_: bool
                d_241_v119_: int
                d_242_v120_: int
                out18_: bool
                out19_: int
                out20_: int
                out18_, out19_, out20_ = default__.m0((d_199_v86_).intersection((((d_238_v116_)[d_233_v111_] if (d_233_v111_) in (d_238_v116_) else _dafny.MultiSet([d_221_v105_]))).set((d_239_v117_).cf1, default__.abs(len(d_233_v111_)))), d_223_v107_, d_96_globalState_)
                d_240_v118_ = out18_
                d_241_v119_ = out19_
                d_242_v120_ = out20_
                d_243_v121_: D1
                d_243_v121_ = D1_DC4()
                d_243_v121_ = D1_DC4()
                index46_ = default__.safeIndex(445, (d_196_v85_).length(0))
                (d_196_v85_)[index46_] = (((d_203_v90_).cf2 if d_204_v91_ else d_206_v93_)) != ((0) - ((len(_dafny.Set({(d_199_v86_).cardinality, d_206_v93_, d_206_v93_, len(d_225_v109_), (d_235_v113_)[default__.safeIndex(687, (d_235_v113_).length(0))]}))) - (d_98_v1_)))
                index47_ = default__.safeIndex(687, (d_235_v113_).length(0))
                rhs46_ = d_242_v120_
                rhs47_ = default__.safeDivisionInt(d_222_v106_, (d_205_v92_) + (d_206_v93_))
                lhs42_ = d_235_v113_
                lhs43_ = default__.safeIndex(687, (d_235_v113_).length(0))
                lhs42_[lhs43_] = rhs46_
                d_242_v120_ = rhs47_
                index48_ = default__.safeIndex(687, (d_235_v113_).length(0))
                (d_235_v113_)[index48_] = d_241_v119_
            elif True:
                d_244_v122_: C0
                nw38_ = C0()
                nw38_.ctor__(d_195_v84_)
                d_244_v122_ = nw38_
                index49_ = default__.safeIndex(445, (d_196_v85_).length(0))
                (d_196_v85_)[index49_] = not(d_204_v91_)
                index50_ = default__.safeIndex(687, (d_235_v113_).length(0))
                (d_235_v113_)[index50_] = 514
                d_245_v123_: C0
                nw39_ = C0()
                nw39_.ctor__(d_195_v84_)
                d_245_v123_ = nw39_
                d_246_v124_: _dafny.Array
                def lambda16_(d_247_v91_, d_248_v0_, d_249_v84_):
                    def lambda17_(d_250_i7_):
                        return (d_248_v0_ if d_247_v91_ else _dafny.SeqWithoutIsStrInference([d_249_v84_ for d_251_i8_ in range(default__.abs(495))]))

                    return lambda17_

                init8_ = lambda16_(d_204_v91_, d_97_v0_, d_195_v84_)
                nw40_ = _dafny.Array(None, 26)
                for i0_8_ in range(nw40_.length(0)):
                    nw40_[i0_8_] = init8_(i0_8_)
                d_246_v124_ = nw40_
                d_252_v125_: D3
                d_252_v125_ = D3_DC8(d_246_v124_)
                index51_ = default__.safeIndex(687, (d_235_v113_).length(0))
                rhs48_ = (633) * (((d_236_v114_)[default__.safeIndex(d_206_v93_, len(d_236_v114_))]) + (-719))
                rhs49_ = (d_252_v125_).cf9
                rhs50_ = d_206_v93_
                rhs51_ = (((d_199_v86_)[d_204_v91_] if (d_204_v91_) in (d_199_v86_) else default__.fm0(d_96_globalState_))) == (d_223_v107_)
                lhs44_ = d_235_v113_
                lhs45_ = default__.safeIndex(687, (d_235_v113_).length(0))
                lhs46_ = d_96_globalState_
                lhs44_[lhs45_] = rhs48_
                d_246_v124_ = rhs49_
                d_206_v93_ = rhs50_
                lhs46_.f0 = rhs51_
        d_253_v126_: _dafny.Seq
        d_253_v126_ = _dafny.SeqWithoutIsStrInference([d_205_v92_])
        d_254_v127_: C0
        nw41_ = C0()
        nw41_.ctor__(d_195_v84_)
        d_254_v127_ = nw41_
        d_255_v128_: _dafny.Map
        d_255_v128_ = _dafny.Map({d_205_v92_: d_254_v127_})
        d_256_v129_: _dafny.Seq
        d_256_v129_ = _dafny.SeqWithoutIsStrInference([((d_255_v128_)[d_206_v93_] if (d_206_v93_) in (d_255_v128_) else d_254_v127_), d_254_v127_])
        d_98_v1_ = (d_98_v1_) + (len((d_253_v126_) + (_dafny.SeqWithoutIsStrInference([638, d_205_v92_, len(d_256_v129_)]))))
        d_98_v1_ = d_206_v93_
        d_100_v3_ = not(d_204_v91_)
        d_257_i9_: int
        d_257_i9_ = 0
        with _dafny.label("0"):
            while ((d_253_v126_).set(default__.safeIndex(d_206_v93_, len(d_253_v126_)), d_205_v92_)) == ((d_253_v126_) + (d_253_v126_)):
                with _dafny.c_label("0"):
                    if (d_257_i9_) >= (100):
                        raise _dafny.Break("0")
                    d_257_i9_ = (d_257_i9_) + (1)
                    d_258_v130_: _dafny.Seq
                    d_258_v130_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_206_v93_: d_204_v91_}), _dafny.Map({len(d_253_v126_): (D1_DC2(58, (d_196_v85_)[default__.safeIndex(445, (d_196_v85_).length(0))])).cf3})])
                    d_258_v130_ = d_258_v130_
                    d_206_v93_ = (d_253_v126_)[default__.safeIndex(d_98_v1_, len(d_253_v126_))]
                    d_205_v92_ = d_98_v1_
                    d_259_v131_: T0
                    nw42_ = C0()
                    nw42_.ctor__(_dafny.CodePoint('g'))
                    d_259_v131_ = nw42_
                    d_259_v131_ = d_259_v131_
                    pass
            pass
        (d_96_globalState_).f0 = (d_100_v3_) and (d_204_v91_)
        d_204_v91_ = d_100_v3_
        _dafny.print(_dafny.string_of(d_96_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_v0_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_98_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v2_) == (_dafny.Map({-27: -27}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_100_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_195_v84_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v85_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v85_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v85_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v85_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v85_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v85_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v86_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_200_v87_) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_201_v88_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({1}), _dafny.Set({1}), _dafny.Set({1})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_202_v89_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_v90_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_v90_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_204_v91_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_205_v92_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_206_v93_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_253_v126_) == (_dafny.SeqWithoutIsStrInference([455]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_255_v128_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_256_v129_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_257_i9_))
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
        return lambda: D1_DC2(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4)
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
        return lambda: D2_DC6(None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC6(D2, NamedTuple('DC6', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC9(D3, NamedTuple('DC9', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(D1.default()(), _dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC11(D4, NamedTuple('DC11', [('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({self.cf11.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC13(D5, NamedTuple('DC13', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self._f1: bool = False
        self._f2: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2

    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2

class C0(T0):
    def  __init__(self):
        self._f3: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f3(self):
        return self._f3
    def ctor__(self, f3):
        (self)._f3 = f3

    def fm3(self, globalState):
        source2_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqw")))
        d_260___mcc_h0_ = source2_
        d_261_cf0_ = d_260___mcc_h0_
        return not(False)

    def fm4(self, p0, p1, globalState):
        return (D1_DC1(True)).cf1

