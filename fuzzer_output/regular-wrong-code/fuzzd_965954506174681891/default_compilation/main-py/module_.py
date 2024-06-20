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
        return ((len(_dafny.SeqWithoutIsStrInference([True, not(False)]))) * (840)) < ((-513) + (882))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return 736

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "euwd"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))

    @staticmethod
    def fm5(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: _dafny.Seq
            for compr_0_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_0_i0_ in range(default__.abs(577))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1_i1_ in range(default__.abs(-211))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmpbyeb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jeeg"))])).Elements:
                d_2_v0_: _dafny.Seq = compr_0_
                if (d_2_v0_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_0_i0_ in range(default__.abs(577))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1_i1_ in range(default__.abs(-211))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmpbyeb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jeeg"))])):
                    coll0_ = coll0_.union(_dafny.Set([d_2_v0_]))
            return _dafny.Set(coll0_)
        return iife0_()
        

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: _dafny.Seq
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_3_i1_ in range(default__.abs(715))]) for d_4_i0_ in range(default__.abs(240))])).Elements:
                d_5_v0_: _dafny.Seq = compr_1_
                if (d_5_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_3_i1_ in range(default__.abs(715))]) for d_4_i0_ in range(default__.abs(240))])):
                    coll1_ = coll1_.union(_dafny.Set([d_5_v0_]))
            return _dafny.Set(coll1_)
        return D1_DC4((iife1_()
).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khuk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbthcxlck"))})))

    @staticmethod
    def fm7(globalState):
        return ((_dafny.SeqWithoutIsStrInference([582, (D7_DC17(False, len(_dafny.Map({121: 901})), False, True)).cf29, -450])) + (_dafny.SeqWithoutIsStrInference([630, 714]))) + (_dafny.SeqWithoutIsStrInference([549]))

    @staticmethod
    def fm8(p0, p1, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(469, -47):
                d_6_v0_: int = compr_2_
                if ((469) <= (d_6_v0_)) and ((d_6_v0_) < (-47)):
                    coll2_ = coll2_.union(_dafny.Set([(d_6_v0_) - (437)]))
            return _dafny.Set(coll2_)
        return _dafny.SeqWithoutIsStrInference([(528) != (len(_dafny.Map({not(not(True)): (D7_DC17(True, len(_dafny.SeqWithoutIsStrInference([706, 932])), True, not(True))).cf30}))), ((_dafny.MultiSet([len(iife2_()
        ), (0) - (len(_dafny.Set({False}))), len(_dafny.Map({False: _dafny.Map({not(not(True)): len(_dafny.Map({False: True}))})})), 581, (0) - (-48)])).cardinality) >= (561), False, not((len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(True), not(True)]), _dafny.SeqWithoutIsStrInference([True, True, True]), _dafny.SeqWithoutIsStrInference([False, False, False, False, True]), _dafny.SeqWithoutIsStrInference([not(True), True, True]), _dafny.SeqWithoutIsStrInference([True, False, True])]))) == ((_dafny.MultiSet([False, False])).cardinality)), False])

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({True: True}))) | ((_dafny.Map({True: False})) | (_dafny.Map({False: True})))

    @staticmethod
    def fm10(p0, p1, globalState):
        if (True) or (not(True)):
            return _dafny.CodePoint('j')
        elif True:
            return _dafny.CodePoint('u')

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_7_i0_ in range(default__.abs(65))])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xm")))

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return ((_dafny.Set({False})) - (_dafny.Set({True, False}))) | ((_dafny.Set({True, False})) - (_dafny.Set({True})))

    @staticmethod
    def fm13(p0, p1, globalState):
        return _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({986: (_dafny.MultiSet([True])).cardinality}))])), len(_dafny.Set({_dafny.CodePoint('k')}))]))) | (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-994, len(_dafny.Map({False: -991}))]))])): default__.safeModuloInt(len(_dafny.Map({-577: _dafny.Map({660: 407})})), 584)})

    @staticmethod
    def m0(globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_8_v0_: int
        d_8_v0_ = 301
        d_9_v1_: bool
        d_9_v1_ = False
        d_10_v2_: _dafny.Map
        d_10_v2_ = _dafny.Map({d_9_v1_: _dafny.CodePoint('l')})
        hi0_ = (0) - (len((d_10_v2_) | (_dafny.Map({d_9_v1_: _dafny.CodePoint('k')}))))
        for d_11_i0_ in range(d_8_v0_, hi0_):
            r1 = d_8_v0_
            d_8_v0_ = d_8_v0_
            if d_9_v1_:
                d_12_v3_: _dafny.Seq
                d_12_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                d_13_v4_: str
                d_13_v4_ = _dafny.CodePoint('l')
                rhs0_ = (_dafny.SeqWithoutIsStrInference([d_9_v1_, (d_12_v3_) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_14_i1_ in range(default__.abs(196))])), d_9_v1_, d_9_v1_, (d_9_v1_) and (default__.fm0(d_9_v1_, d_9_v1_, d_12_v3_, globalState))])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_11_i0_ for d_15_i3_ in range(default__.abs(-835))]) for d_16_i2_ in range(default__.abs(267))])), len(_dafny.SeqWithoutIsStrInference([d_9_v1_, (d_12_v3_) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_17_i1_ in range(default__.abs(196))])), d_9_v1_, d_9_v1_, (d_9_v1_) and (default__.fm0(d_9_v1_, d_9_v1_, d_12_v3_, globalState))]))), (d_9_v1_) and (d_9_v1_))
                rhs1_ = d_11_i0_
                rhs2_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "quhelj"))) + ((d_12_v3_).set(default__.safeIndex(d_8_v0_, len(d_12_v3_)), d_13_v4_))) + (d_12_v3_)
                lhs0_ = globalState
                lhs1_ = globalState
                lhs0_.f9 = rhs0_
                lhs1_.f7 = rhs1_
                d_12_v3_ = rhs2_
                d_18_v5_: _dafny.Array
                def lambda0_(d_19_i0_):
                    def lambda1_(d_20_i4_):
                        return (_dafny.SeqWithoutIsStrInference([d_19_i0_])) + (_dafny.SeqWithoutIsStrInference([d_19_i0_]))

                    return lambda1_

                init0_ = lambda0_(d_11_i0_)
                nw0_ = _dafny.Array(None, 27)
                for i0_0_ in range(nw0_.length(0)):
                    nw0_[i0_0_] = init0_(i0_0_)
                d_18_v5_ = nw0_
                d_21_v6_: _dafny.Array
                nw1_ = _dafny.Array(False, 26)
                d_21_v6_ = nw1_
                d_22_v7_: _dafny.Array
                nw2_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_22_v7_ = nw2_
                index0_ = default__.safeIndex(651, (d_22_v7_).length(0))
                (d_22_v7_)[index0_] = d_21_v6_
                d_23_v8_: _dafny.Map
                d_23_v8_ = _dafny.Map({d_9_v1_: d_21_v6_})
                d_24_v9_: D5
                d_24_v9_ = D5_DC11(d_9_v1_, d_11_i0_, True, d_21_v6_, 319)
                d_25_v10_: _dafny.Map
                d_25_v10_ = _dafny.Map({(d_12_v3_)[default__.safeIndex(d_8_v0_, len(d_12_v3_))]: d_12_v3_})
                d_26_v11_: _dafny.Set
                d_26_v11_ = _dafny.Set({default__.fm1(False, d_8_v0_, (d_8_v0_), globalState)})
                index1_ = default__.safeIndex(651, (d_22_v7_).length(0))
                rhs3_ = d_18_v5_
                rhs4_ = ((d_23_v8_)[not(d_9_v1_)] if (not(d_9_v1_)) in (d_23_v8_) else d_21_v6_)
                rhs5_ = (d_24_v9_).cf14
                rhs6_ = (len(d_25_v10_)) > (default__.safeModuloInt(d_8_v0_, len(d_26_v11_)))
                lhs2_ = d_22_v7_
                lhs3_ = default__.safeIndex(651, (d_22_v7_).length(0))
                d_18_v5_ = rhs3_
                d_21_v6_ = rhs4_
                lhs2_[lhs3_] = rhs5_
                d_9_v1_ = rhs6_
                d_8_v0_ = d_11_i0_
                d_27_v12_: _dafny.Array
                nw3_ = _dafny.Array(int(0), 3)
                d_27_v12_ = nw3_
                d_28_v13_: _dafny.Seq
                d_28_v13_ = _dafny.SeqWithoutIsStrInference([d_8_v0_])
                d_29_v14_: _dafny.Map
                d_29_v14_ = _dafny.Map({d_9_v1_: d_9_v1_})
                d_30_v15_: C0
                nw4_ = C0()
                nw4_.ctor__(d_9_v1_, d_29_v14_)
                d_30_v15_ = nw4_
                d_31_v16_: _dafny.Seq
                d_31_v16_ = _dafny.SeqWithoutIsStrInference([len(d_28_v13_), len(_dafny.Map({d_30_v15_: d_8_v0_}))])
                d_32_v17_: _dafny.Map
                d_32_v17_ = _dafny.Map({len(d_28_v13_): d_9_v1_})
                d_33_v18_: _dafny.MultiSet
                d_33_v18_ = _dafny.MultiSet([d_11_i0_, d_8_v0_])
                d_34_v19_: _dafny.Seq
                d_34_v19_ = _dafny.SeqWithoutIsStrInference([d_9_v1_, ((d_32_v17_)[((d_33_v18_)[(0) - (d_8_v0_)] if ((0) - (d_8_v0_)) in (d_33_v18_) else d_8_v0_)] if (((d_33_v18_)[(0) - (d_8_v0_)] if ((0) - (d_8_v0_)) in (d_33_v18_) else d_8_v0_)) in (d_32_v17_) else d_30_v15_.f19), d_30_v15_.f19])
                d_35_v20_: _dafny.Map
                d_35_v20_ = _dafny.Map({d_27_v12_: d_34_v19_})
                d_35_v20_ = d_35_v20_
                d_13_v4_ = d_13_v4_
            elif True:
                (globalState).f8 = d_8_v0_
                rhs7_ = (d_8_v0_) + (d_8_v0_)
                rhs8_ = d_9_v1_
                r1 = rhs7_
                d_9_v1_ = rhs8_
                d_36_v21_: _dafny.Array
                nw5_ = _dafny.Array(int(0), 22)
                d_36_v21_ = nw5_
                d_36_v21_ = d_36_v21_
                d_37_v22_: _dafny.Seq
                d_37_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbhg"))
                d_38_v23_: C0
                nw6_ = C0()
                nw6_.ctor__(d_9_v1_, default__.fm9(d_8_v0_, len(d_37_v22_), default__.fm1(d_9_v1_, d_11_i0_, len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")): d_8_v0_})), globalState), globalState))
                d_38_v23_ = nw6_
                d_39_v24_: str
                d_39_v24_ = _dafny.CodePoint('y')
                d_40_v25_: _dafny.Map
                d_40_v25_ = _dafny.Map({d_9_v1_: 770})
                rhs9_ = default__.fm0(d_38_v23_.f19, (d_8_v0_) >= (d_8_v0_), (d_37_v22_).set(default__.safeIndex(d_11_i0_, len(d_37_v22_)), d_39_v24_), globalState)
                rhs10_ = d_38_v23_
                rhs11_ = (0) - ((d_11_i0_) + (len(d_40_v25_)))
                rhs12_ = d_37_v22_
                rhs13_ = (d_38_v23_.f19 if d_9_v1_ else d_9_v1_)
                d_9_v1_ = rhs9_
                d_38_v23_ = rhs10_
                r0 = rhs11_
                d_37_v22_ = rhs12_
                d_9_v1_ = rhs13_
                d_8_v0_ = d_11_i0_
            d_41_v26_: _dafny.Array
            nw7_ = _dafny.Array(False, 2)
            d_41_v26_ = nw7_
            d_42_v27_: _dafny.Seq
            d_42_v27_ = _dafny.SeqWithoutIsStrInference([d_41_v26_, d_41_v26_, (d_41_v26_ if d_9_v1_ else d_41_v26_)])
            d_41_v26_ = (d_42_v27_)[default__.safeIndex(984, len(d_42_v27_))]
        r0 = d_8_v0_
        d_43_v28_: str
        d_43_v28_ = _dafny.CodePoint('f')
        d_44_v29_: _dafny.Seq
        d_44_v29_ = _dafny.SeqWithoutIsStrInference([d_43_v28_, d_43_v28_, d_43_v28_])
        d_45_v30_: _dafny.Seq
        d_45_v30_ = _dafny.SeqWithoutIsStrInference([(d_44_v29_) < (_dafny.SeqWithoutIsStrInference([d_43_v28_ for d_46_i5_ in range(default__.abs(373))])), d_9_v1_])
        d_47_v31_: _dafny.Map
        d_47_v31_ = _dafny.Map({d_8_v0_: len(_dafny.SeqWithoutIsStrInference([d_8_v0_ for d_48_i6_ in range(default__.abs(709))]))})
        rhs14_ = ((12 if not(d_9_v1_) else d_8_v0_)) * (d_8_v0_)
        rhs15_ = (d_45_v30_)[default__.safeIndex((d_8_v0_) * (((d_47_v31_)[len(d_44_v29_)] if (len(d_44_v29_)) in (d_47_v31_) else default__.fm1(d_9_v1_, d_8_v0_, d_8_v0_, globalState))), len(d_45_v30_))]
        r1 = rhs14_
        d_9_v1_ = rhs15_
        d_49_v32_: _dafny.Array
        nw8_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
        d_49_v32_ = nw8_
        d_50_v33_: _dafny.Seq
        d_50_v33_ = d_44_v29_
        index2_ = default__.safeIndex(704, (d_49_v32_).length(0))
        (d_49_v32_)[index2_] = (d_50_v33_)
        index3_ = default__.safeIndex(704, (d_49_v32_).length(0))
        (d_49_v32_)[index3_] = d_44_v29_
        d_51_v34_: _dafny.Array
        def lambda2_(d_52_v0_):
            def lambda3_(d_53_i7_):
                return (d_52_v0_) < (d_52_v0_)

            return lambda3_

        init1_ = lambda2_(d_8_v0_)
        nw9_ = _dafny.Array(None, 27)
        for i0_1_ in range(nw9_.length(0)):
            nw9_[i0_1_] = init1_(i0_1_)
        d_51_v34_ = nw9_
        index4_ = default__.safeIndex(798, (d_51_v34_).length(0))
        (d_51_v34_)[index4_] = (483) > (d_8_v0_)
        d_54_v35_: _dafny.Array
        nw10_ = _dafny.Array(int(0), 9)
        d_54_v35_ = nw10_
        d_55_v36_: _dafny.Map
        d_55_v36_ = _dafny.Map({d_45_v30_: d_54_v35_})
        d_56_v37_: D0
        d_56_v37_ = D0_DC0((d_55_v36_) | (d_55_v36_))
        pat_let_tv0_ = d_55_v36_
        index5_ = default__.safeIndex(798, (d_51_v34_).length(0))
        def iife3_(_pat_let0_0):
            def iife4_(d_57_dt__update__tmp_h0_):
                def iife5_(_pat_let1_0):
                    def iife6_(d_58_dt__update_hcf0_h0_):
                        return D0_DC0(d_58_dt__update_hcf0_h0_)
                    return iife6_(_pat_let1_0)
                return iife5_(pat_let_tv0_)
            return iife4_(_pat_let0_0)
        rhs16_ = d_9_v1_
        rhs17_ = d_8_v0_
        rhs18_ = d_9_v1_
        rhs19_ = iife3_(d_56_v37_)
        lhs4_ = d_51_v34_
        lhs5_ = default__.safeIndex(798, (d_51_v34_).length(0))
        d_9_v1_ = rhs16_
        r1 = rhs17_
        lhs4_[lhs5_] = rhs18_
        d_56_v37_ = rhs19_
        if (d_8_v0_) < (d_8_v0_):
            index6_ = default__.safeIndex(798, (d_51_v34_).length(0))
            (d_51_v34_)[index6_] = not(d_9_v1_)
            r0 = d_8_v0_
            (globalState).f7 = d_8_v0_
            if d_9_v1_:
                d_59_v38_: _dafny.Map
                d_59_v38_ = _dafny.Map({d_9_v1_: not(d_9_v1_)})
                d_60_v39_: _dafny.Seq
                d_60_v39_ = _dafny.SeqWithoutIsStrInference([d_59_v38_, d_59_v38_, d_59_v38_, d_59_v38_, (d_59_v38_).set((d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))], d_9_v1_)])
                d_61_v40_: _dafny.Seq
                d_61_v40_ = _dafny.SeqWithoutIsStrInference([default__.fm1(True, len(d_47_v31_), d_8_v0_, globalState), d_8_v0_])
                d_62_v41_: _dafny.Array
                nw11_ = _dafny.Array(None, 21)
                nw11_[int(0)] = d_59_v38_
                nw11_[int(1)] = d_59_v38_
                nw11_[int(2)] = _dafny.Map({d_9_v1_: (d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]})
                nw11_[int(3)] = d_59_v38_
                nw11_[int(4)] = d_59_v38_
                nw11_[int(5)] = d_59_v38_
                nw11_[int(6)] = d_59_v38_
                nw11_[int(7)] = d_59_v38_
                nw11_[int(8)] = d_59_v38_
                nw11_[int(9)] = _dafny.Map({(d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]: (d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]})
                nw11_[int(10)] = d_59_v38_
                nw11_[int(11)] = _dafny.Map({d_9_v1_: (d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]})
                nw11_[int(12)] = d_59_v38_
                nw11_[int(13)] = (d_60_v39_)[default__.safeIndex(len(d_61_v40_), len(d_60_v39_))]
                nw11_[int(14)] = d_59_v38_
                nw11_[int(15)] = _dafny.Map({d_9_v1_: (d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]})
                nw11_[int(16)] = d_59_v38_
                nw11_[int(17)] = d_59_v38_
                nw11_[int(18)] = _dafny.Map({(d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]: d_9_v1_})
                nw11_[int(19)] = d_59_v38_
                nw11_[int(20)] = d_59_v38_
                d_62_v41_ = nw11_
                d_63_v42_: _dafny.Map
                d_63_v42_ = _dafny.Map({False: d_62_v41_})
                d_64_v43_: _dafny.Seq
                d_64_v43_ = _dafny.SeqWithoutIsStrInference([d_62_v41_, d_62_v41_, d_62_v41_])
                d_65_v44_: _dafny.Set
                d_65_v44_ = _dafny.Set({d_8_v0_, len(d_61_v40_)})
                d_66_v45_: bool
                d_66_v45_ = d_9_v1_
                d_67_v46_: _dafny.Array
                nw12_ = _dafny.Array(None, 27)
                nw12_[int(0)] = d_62_v41_
                nw12_[int(1)] = d_62_v41_
                nw12_[int(2)] = d_62_v41_
                nw12_[int(3)] = d_62_v41_
                nw12_[int(4)] = d_62_v41_
                nw12_[int(5)] = d_62_v41_
                nw12_[int(6)] = d_62_v41_
                nw12_[int(7)] = d_62_v41_
                nw12_[int(8)] = d_62_v41_
                nw12_[int(9)] = d_62_v41_
                nw12_[int(10)] = d_62_v41_
                nw12_[int(11)] = d_62_v41_
                nw12_[int(12)] = d_62_v41_
                nw12_[int(13)] = d_62_v41_
                nw12_[int(14)] = d_62_v41_
                nw12_[int(15)] = ((d_63_v42_)[(d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]] if ((d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]) in (d_63_v42_) else d_62_v41_)
                nw12_[int(16)] = d_62_v41_
                nw12_[int(17)] = d_62_v41_
                nw12_[int(18)] = (d_64_v43_)[default__.safeIndex(len(d_65_v44_), len(d_64_v43_))]
                nw12_[int(19)] = (d_62_v41_ if d_9_v1_ else d_62_v41_)
                nw12_[int(20)] = d_62_v41_
                nw12_[int(21)] = d_62_v41_
                nw12_[int(22)] = d_62_v41_
                nw12_[int(23)] = d_62_v41_
                nw12_[int(24)] = d_62_v41_
                nw12_[int(25)] = d_62_v41_
                nw12_[int(26)] = (d_62_v41_ if (d_66_v45_) else d_62_v41_)
                d_67_v46_ = nw12_
                index7_ = default__.safeIndex(74, (d_67_v46_).length(0))
                (d_67_v46_)[index7_] = d_62_v41_
                index8_ = default__.safeIndex(74, (d_67_v46_).length(0))
                def lambda4_(d_68_v38_):
                    def lambda5_(d_69_i8_):
                        return d_68_v38_

                    return lambda5_

                init2_ = lambda4_(d_59_v38_)
                nw13_ = _dafny.Array(None, 23)
                for i0_2_ in range(nw13_.length(0)):
                    nw13_[i0_2_] = init2_(i0_2_)
                (d_67_v46_)[index8_] = nw13_
                d_70_v47_: C0
                nw14_ = C0()
                nw14_.ctor__(d_9_v1_, d_59_v38_)
                d_70_v47_ = nw14_
                d_70_v47_ = d_70_v47_
                r1 = d_8_v0_
                d_9_v1_ = ((d_8_v0_) - (len(d_59_v38_))) != (((0) - (len((d_44_v29_).set(default__.safeIndex(d_8_v0_, len(d_44_v29_)), d_43_v28_)))) - ((0) - (d_8_v0_)))
                (globalState).f8 = d_8_v0_
            elif True:
                d_8_v0_ = (0) - (d_8_v0_)
                d_71_v48_: _dafny.Map
                d_71_v48_ = _dafny.Map({d_9_v1_: d_9_v1_})
                d_72_v49_: C0
                nw15_ = C0()
                nw15_.ctor__(d_9_v1_, d_71_v48_)
                d_72_v49_ = nw15_
                d_73_v50_: _dafny.Seq
                d_73_v50_ = _dafny.SeqWithoutIsStrInference([d_54_v35_])
                d_74_v51_: _dafny.MultiSet
                d_74_v51_ = _dafny.MultiSet([d_54_v35_, (d_73_v50_)[default__.safeIndex(len(d_44_v29_), len(d_73_v50_))], d_54_v35_, d_54_v35_])
                d_75_v52_: _dafny.Map
                d_75_v52_ = _dafny.Map({(d_8_v0_) - (d_8_v0_): d_72_v49_})
                d_76_v53_: _dafny.Array
                nw16_ = _dafny.Array(D1.default()(), 5)
                d_76_v53_ = nw16_
                d_77_v54_: D1
                d_77_v54_ = D1_DC4(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvldbtq")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiws"))}))
                d_78_v55_: D1
                d_78_v55_ = D1_DC5(d_77_v54_)
                index9_ = default__.safeIndex(147, (d_76_v53_).length(0))
                (d_76_v53_)[index9_] = d_78_v55_
                d_79_v56_: _dafny.MultiSet
                d_79_v56_ = _dafny.MultiSet([d_45_v30_])
                index10_ = default__.safeIndex(147, (d_76_v53_).length(0))
                rhs20_ = (d_72_v49_ if not(d_9_v1_) else d_72_v49_)
                rhs21_ = d_74_v51_
                rhs22_ = ((0) - (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_45_v30_]))).intersection(d_79_v56_)).cardinality)) + (79)
                rhs23_ = (d_75_v52_) | (d_75_v52_)
                rhs24_ = d_78_v55_
                lhs6_ = globalState
                lhs7_ = d_76_v53_
                lhs8_ = default__.safeIndex(147, (d_76_v53_).length(0))
                d_72_v49_ = rhs20_
                d_74_v51_ = rhs21_
                lhs6_.f7 = rhs22_
                d_75_v52_ = rhs23_
                lhs7_[lhs8_] = rhs24_
                d_10_v2_ = (d_10_v2_).set(d_9_v1_, d_43_v28_)
                index11_ = default__.safeIndex(798, (d_51_v34_).length(0))
                (d_51_v34_)[index11_] = (d_9_v1_) in (d_45_v30_)
                d_80_v57_: _dafny.Map
                d_80_v57_ = _dafny.Map({d_43_v28_: d_72_v49_.f19})
                d_80_v57_ = (d_80_v57_) | (d_80_v57_)
            r0 = (0) - (d_8_v0_)
        elif True:
            index12_ = default__.safeIndex(494, (d_54_v35_).length(0))
            (d_54_v35_)[index12_] = d_8_v0_
            index13_ = default__.safeIndex(494, (d_54_v35_).length(0))
            (d_54_v35_)[index13_] = d_8_v0_
            d_81_v58_: _dafny.Set
            d_81_v58_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]]), default__.fm8(-7, d_9_v1_, globalState)})
            index14_ = default__.safeIndex(798, (d_51_v34_).length(0))
            (d_51_v34_)[index14_] = (_dafny.Set({d_45_v30_})).ispropersubset(d_81_v58_)
            d_8_v0_ = d_8_v0_
            d_82_v59_: _dafny.Set
            d_82_v59_ = _dafny.Set({d_8_v0_, (d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))]})
            d_83_v60_: _dafny.MultiSet
            d_83_v60_ = _dafny.MultiSet([(d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))], (d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))]])
            d_84_v61_: _dafny.Map
            d_84_v61_ = _dafny.Map({(D7_DC14(d_83_v60_)).cf18: (d_49_v32_)[default__.safeIndex(704, (d_49_v32_).length(0))]})
            d_85_v62_: D5
            d_85_v62_ = D5_DC10(d_82_v59_, d_84_v61_)
            source0_ = d_85_v62_
            if source0_.is_DC10:
                d_86___mcc_h0_ = source0_.cf9
                d_87___mcc_h1_ = source0_.cf10
                d_88_cf10_ = d_87___mcc_h1_
                d_89_cf9_ = d_86___mcc_h0_
                index15_ = default__.safeIndex(494, (d_54_v35_).length(0))
                (d_54_v35_)[index15_] = d_8_v0_
                d_90_v63_: _dafny.Set
                d_90_v63_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_43_v28_ for d_91_i9_ in range(default__.abs(318))])})
                d_9_v1_ = (d_90_v63_).ispropersubset((d_90_v63_) | (d_90_v63_))
                d_92_v64_: _dafny.Array
                nw17_ = _dafny.Array(_dafny.MultiSet({}), 6)
                d_92_v64_ = nw17_
                d_93_v65_: D1
                d_93_v65_ = D1_DC3(d_92_v64_)
                pat_let_tv1_ = d_92_v64_
                d_94_v66_: _dafny.Array
                nw18_ = _dafny.Array(None, 28)
                nw18_[int(0)] = d_93_v65_
                nw18_[int(1)] = d_93_v65_
                nw18_[int(2)] = d_93_v65_
                nw18_[int(3)] = d_93_v65_
                nw18_[int(4)] = d_93_v65_
                nw18_[int(5)] = d_93_v65_
                nw18_[int(6)] = d_93_v65_
                nw18_[int(7)] = D1_DC3(d_92_v64_)
                nw18_[int(8)] = d_93_v65_
                nw18_[int(9)] = d_93_v65_
                nw18_[int(10)] = d_93_v65_
                nw18_[int(11)] = (d_93_v65_ if d_9_v1_ else d_93_v65_)
                nw18_[int(12)] = d_93_v65_
                nw18_[int(13)] = d_93_v65_
                nw18_[int(14)] = d_93_v65_
                nw18_[int(15)] = d_93_v65_
                nw18_[int(16)] = D1_DC3((D1_DC3(d_92_v64_)).cf2)
                nw18_[int(17)] = (d_93_v65_ if d_9_v1_ else d_93_v65_)
                nw18_[int(18)] = d_93_v65_
                nw18_[int(19)] = D1_DC3(d_92_v64_)
                nw18_[int(20)] = d_93_v65_
                nw18_[int(21)] = D1_DC3(d_92_v64_)
                nw18_[int(22)] = d_93_v65_
                def iife7_(_pat_let2_0):
                    def iife8_(d_95_dt__update__tmp_h1_):
                        def iife9_(_pat_let3_0):
                            def iife10_(d_96_dt__update_hcf2_h0_):
                                return D1_DC3(d_96_dt__update_hcf2_h0_)
                            return iife10_(_pat_let3_0)
                        return iife9_(pat_let_tv1_)
                    return iife8_(_pat_let2_0)
                nw18_[int(23)] = iife7_(d_93_v65_)
                nw18_[int(24)] = d_93_v65_
                nw18_[int(25)] = d_93_v65_
                nw18_[int(26)] = d_93_v65_
                nw18_[int(27)] = d_93_v65_
                d_94_v66_ = nw18_
                index16_ = default__.safeIndex(691, (d_94_v66_).length(0))
                (d_94_v66_)[index16_] = d_93_v65_
                index17_ = default__.safeIndex(691, (d_94_v66_).length(0))
                (d_94_v66_)[index17_] = d_93_v65_
                d_97_v67_: D7
                d_97_v67_ = D7_DC14(d_83_v60_)
                index18_ = default__.safeIndex(494, (d_54_v35_).length(0))
                rhs25_ = (d_97_v67_ if default__.fm0(not(not(d_9_v1_)), d_9_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")), globalState) else d_97_v67_)
                rhs26_ = (d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))]
                lhs9_ = d_54_v35_
                lhs10_ = default__.safeIndex(494, (d_54_v35_).length(0))
                d_97_v67_ = rhs25_
                lhs9_[lhs10_] = rhs26_
            elif source0_.is_DC11:
                d_98___mcc_h2_ = source0_.cf11
                d_99___mcc_h3_ = source0_.cf12
                d_100___mcc_h4_ = source0_.cf13
                d_101___mcc_h5_ = source0_.cf14
                d_102___mcc_h6_ = source0_.cf15
                d_103_cf15_ = d_102___mcc_h6_
                d_104_cf14_ = d_101___mcc_h5_
                d_105_cf13_ = d_100___mcc_h4_
                d_106_cf12_ = d_99___mcc_h3_
                d_107_cf11_ = d_98___mcc_h2_
                d_108_v68_: D5
                d_108_v68_ = D5_DC11((d_45_v30_)[default__.safeIndex(d_106_cf12_, len(d_45_v30_))], d_8_v0_, d_9_v1_, d_104_cf14_, (0) - (d_106_cf12_))
                d_108_v68_ = d_108_v68_
                d_8_v0_ = (d_8_v0_ if d_107_cf11_ else d_8_v0_)
                d_107_cf11_ = d_107_cf11_
                d_109_v69_: C0
                nw19_ = C0()
                nw19_.ctor__(((d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))]) < (d_8_v0_), _dafny.Map({(d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]: d_9_v1_}))
                d_109_v69_ = nw19_
                d_109_v69_ = d_109_v69_
            elif source0_.is_DC9:
                d_110___mcc_h7_ = source0_.cf8
                d_111_cf8_ = d_110___mcc_h7_
                d_112_v70_: int
                d_112_v70_ = 96
                d_113_v71_: _dafny.Map
                d_113_v71_ = _dafny.Map({not(d_9_v1_): (d_112_v70_)})
                d_114_v72_: _dafny.Array
                nw20_ = _dafny.Array(D5.default()(), 20)
                d_114_v72_ = nw20_
                d_115_v73_: _dafny.MultiSet
                d_115_v73_ = _dafny.MultiSet([d_114_v72_])
                d_113_v71_ = (d_113_v71_).set((d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))], default__.safeDivisionInt((d_115_v73_).cardinality, len((d_49_v32_)[default__.safeIndex(704, (d_49_v32_).length(0))])))
                d_116_v74_: D5
                d_116_v74_ = D5_DC11(d_111_cf8_.f19, (d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))], (d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))], d_51_v34_, (d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))])
                d_117_v75_: _dafny.MultiSet
                d_117_v75_ = _dafny.MultiSet([(d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]])
                d_118_v76_: _dafny.Seq
                d_118_v76_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_111_cf8_.f19, d_9_v1_])), d_117_v75_, d_117_v75_])
                d_119_v77_: _dafny.Array
                nw21_ = _dafny.Array(None, 11)
                nw21_[int(0)] = d_8_v0_
                nw21_[int(1)] = ((d_116_v74_).cf15) + (len((d_111_cf8_).f20))
                nw21_[int(2)] = d_8_v0_
                nw21_[int(3)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivkit")))
                nw21_[int(4)] = (d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))]
                nw21_[int(5)] = (0) - ((d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))])
                nw21_[int(6)] = (0) - ((d_8_v0_) + ((d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))]))
                nw21_[int(7)] = d_8_v0_
                nw21_[int(8)] = len(d_118_v76_)
                nw21_[int(9)] = d_8_v0_
                nw21_[int(10)] = d_8_v0_
                d_119_v77_ = nw21_
                d_119_v77_ = d_119_v77_
                d_120_v78_: _dafny.Set
                d_120_v78_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbduko"))})
                d_120_v78_ = (d_120_v78_) | (d_120_v78_)
                d_121_v79_: _dafny.Array
                def lambda6_(d_122_v30_):
                    def lambda7_(d_123_i10_):
                        return d_122_v30_

                    return lambda7_

                init3_ = lambda6_(d_45_v30_)
                nw22_ = _dafny.Array(None, 26)
                for i0_3_ in range(nw22_.length(0)):
                    nw22_[i0_3_] = init3_(i0_3_)
                d_121_v79_ = nw22_
                d_124_v80_: _dafny.Map
                d_124_v80_ = _dafny.Map({d_121_v79_: (d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))]})
                d_124_v80_ = (d_124_v80_).set(d_121_v79_, (_dafny.MultiSet([(d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))]])).cardinality)
            elif True:
                d_125___mcc_h8_ = source0_.cf16
                d_126_cf16_ = d_125___mcc_h8_
                d_127_v81_: _dafny.Map
                d_127_v81_ = _dafny.Map({d_9_v1_: not((d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))])})
                index19_ = default__.safeIndex(798, (d_51_v34_).length(0))
                rhs27_ = (d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]
                rhs28_ = len(d_127_v81_)
                rhs29_ = True
                lhs11_ = globalState
                lhs12_ = d_51_v34_
                lhs13_ = default__.safeIndex(798, (d_51_v34_).length(0))
                d_9_v1_ = rhs27_
                lhs11_.f8 = rhs28_
                lhs12_[lhs13_] = rhs29_
                d_128_v82_: _dafny.Array
                nw23_ = _dafny.Array(None, 4)
                d_128_v82_ = nw23_
                d_128_v82_ = d_128_v82_
                (globalState).f8 = default__.safeModuloInt(-109, (0) - ((d_54_v35_)[default__.safeIndex(494, (d_54_v35_).length(0))]))
                d_129_v83_: C0
                nw24_ = C0()
                nw24_.ctor__((d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))], (d_127_v81_).set(False, (d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]))
                d_129_v83_ = nw24_
            d_130_v84_: _dafny.Set
            d_130_v84_ = _dafny.Set({d_9_v1_})
            d_131_v85_: _dafny.Map
            d_131_v85_ = _dafny.Map({(d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]: d_130_v84_})
            d_132_v86_: _dafny.Seq
            d_132_v86_ = _dafny.SeqWithoutIsStrInference([((d_131_v85_)[(d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]] if ((d_51_v34_)[default__.safeIndex(798, (d_51_v34_).length(0))]) in (d_131_v85_) else d_130_v84_), d_130_v84_])
            d_132_v86_ = d_132_v86_
        r0 = d_8_v0_
        r1 = default__.safeModuloInt(d_8_v0_, 554)
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_133_v0_: str
        d_133_v0_ = _dafny.CodePoint('u')
        d_134_v1_: _dafny.Array
        def lambda8_(d_135_i0_):
            return default__.safeDivisionInt(d_135_i0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))))

        init4_ = lambda8_
        nw25_ = _dafny.Array(None, 13)
        for i0_4_ in range(nw25_.length(0)):
            nw25_[i0_4_] = init4_(i0_4_)
        d_134_v1_ = nw25_
        d_136_v2_: _dafny.Map
        d_136_v2_ = _dafny.Map({d_133_v0_: d_134_v1_})
        d_137_v3_: bool
        d_137_v3_ = False
        d_138_v4_: _dafny.Set
        d_138_v4_ = _dafny.Set({d_137_v3_, d_137_v3_})
        d_139_v5_: int
        d_139_v5_ = 172
        d_140_v6_: _dafny.Seq
        d_140_v6_ = _dafny.SeqWithoutIsStrInference([d_137_v3_, d_137_v3_, d_137_v3_])
        d_141_v7_: _dafny.Map
        d_141_v7_ = _dafny.Map({d_137_v3_: False})
        d_142_globalState_: GlobalState
        nw26_ = GlobalState()
        nw26_.ctor__(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvfmfbioj")), d_136_v2_, d_138_v4_, _dafny.Map({_dafny.SeqWithoutIsStrInference([d_133_v0_ for d_143_i1_ in range(default__.abs(492))]): d_139_v5_}), -984, False, 148, -376, (d_140_v6_) + (d_140_v6_), False, d_134_v1_, True, d_141_v7_, 363, 680, True, False, False)
        d_142_globalState_ = nw26_
        d_144_v8_: _dafny.MultiSet
        d_144_v8_ = _dafny.MultiSet([d_137_v3_])
        d_137_v3_ = (d_137_v3_) not in ((_dafny.MultiSet([d_137_v3_, d_137_v3_, d_137_v3_])).intersection(d_144_v8_))
        d_145_v9_: _dafny.Seq
        d_145_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qu"))
        d_146_i2_: int
        d_146_i2_ = 0
        with _dafny.label("0"):
            while default__.fm0(d_137_v3_, d_137_v3_, d_145_v9_, d_142_globalState_):
                with _dafny.c_label("0"):
                    if (d_146_i2_) >= (100):
                        raise _dafny.Break("0")
                    d_146_i2_ = (d_146_i2_) + (1)
                    d_147_v10_: _dafny.Array
                    nw27_ = _dafny.Array(None, 28)
                    nw27_[int(0)] = d_137_v3_
                    nw27_[int(1)] = d_137_v3_
                    nw27_[int(2)] = d_137_v3_
                    nw27_[int(3)] = d_137_v3_
                    nw27_[int(4)] = d_137_v3_
                    nw27_[int(5)] = d_137_v3_
                    nw27_[int(6)] = not(d_137_v3_)
                    nw27_[int(7)] = default__.fm0(d_137_v3_, d_137_v3_, _dafny.SeqWithoutIsStrInference([d_133_v0_ for d_148_i3_ in range(default__.abs(558))]), d_142_globalState_)
                    nw27_[int(8)] = d_137_v3_
                    nw27_[int(9)] = False
                    nw27_[int(10)] = d_137_v3_
                    nw27_[int(11)] = False
                    nw27_[int(12)] = False
                    nw27_[int(13)] = d_137_v3_
                    nw27_[int(14)] = d_137_v3_
                    nw27_[int(15)] = d_137_v3_
                    nw27_[int(16)] = d_137_v3_
                    nw27_[int(17)] = d_137_v3_
                    nw27_[int(18)] = d_137_v3_
                    nw27_[int(19)] = d_137_v3_
                    nw27_[int(20)] = d_137_v3_
                    nw27_[int(21)] = d_137_v3_
                    nw27_[int(22)] = d_137_v3_
                    nw27_[int(23)] = d_137_v3_
                    nw27_[int(24)] = d_137_v3_
                    nw27_[int(25)] = d_137_v3_
                    nw27_[int(26)] = d_137_v3_
                    nw27_[int(27)] = default__.fm0(d_137_v3_, True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")), d_142_globalState_)
                    d_147_v10_ = nw27_
                    d_149_v11_: _dafny.MultiSet
                    d_149_v11_ = _dafny.MultiSet([d_147_v10_, d_147_v10_])
                    d_149_v11_ = (((d_149_v11_).set(d_147_v10_, default__.abs(d_139_v5_))) - (_dafny.MultiSet([d_147_v10_]))) - (d_149_v11_)
                    index20_ = default__.safeIndex(42, (d_147_v10_).length(0))
                    (d_147_v10_)[index20_] = True
                    d_150_v12_: _dafny.MultiSet
                    d_150_v12_ = _dafny.MultiSet([(d_145_v9_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([len(d_141_v7_) for d_151_i4_ in range(default__.abs(46))])), len(d_145_v9_))]])
                    d_152_v13_: _dafny.Map
                    d_152_v13_ = _dafny.Map({((d_150_v12_) | (d_150_v12_)).cardinality: d_137_v3_})
                    index21_ = default__.safeIndex(42, (d_147_v10_).length(0))
                    rhs30_ = d_137_v3_
                    rhs31_ = d_152_v13_
                    lhs14_ = d_147_v10_
                    lhs15_ = default__.safeIndex(42, (d_147_v10_).length(0))
                    lhs14_[lhs15_] = rhs30_
                    d_152_v13_ = rhs31_
                    d_153_v14_: _dafny.MultiSet
                    d_153_v14_ = _dafny.MultiSet([d_140_v6_])
                    d_154_v15_: _dafny.Array
                    nw28_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
                    d_154_v15_ = nw28_
                    d_155_v16_: _dafny.Map
                    d_155_v16_ = _dafny.Map({d_153_v14_: d_154_v15_})
                    d_155_v16_ = (d_155_v16_).set(d_153_v14_, d_154_v15_)
                    d_137_v3_ = ((d_147_v10_)[default__.safeIndex(42, (d_147_v10_).length(0))]) and ((d_147_v10_)[default__.safeIndex(42, (d_147_v10_).length(0))])
                    pass
            pass
        d_156_v17_: int
        d_157_v18_: int
        out0_: int
        out1_: int
        out0_, out1_ = default__.m0(d_142_globalState_)
        d_156_v17_ = out0_
        d_157_v18_ = out1_
        hi1_ = d_156_v17_
        for d_158_i5_ in range(default__.safeDivisionInt(d_139_v5_, d_156_v17_), hi1_):
            d_137_v3_ = True
            d_159_v19_: _dafny.Map
            d_159_v19_ = _dafny.Map({d_134_v1_: len(d_140_v6_)})
            d_160_v20_: _dafny.Map
            d_160_v20_ = _dafny.Map({d_133_v0_: d_159_v19_})
            d_160_v20_ = (d_160_v20_).set(d_133_v0_, d_159_v19_)
            d_141_v7_ = d_141_v7_
            d_161_v21_: int
            d_162_v22_: int
            out2_: int
            out3_: int
            out2_, out3_ = default__.m0(d_142_globalState_)
            d_161_v21_ = out2_
            d_162_v22_ = out3_
        (d_142_globalState_).f8 = d_157_v18_
        d_137_v3_ = d_137_v3_
        d_163_v23_: _dafny.Array
        nw29_ = _dafny.Array(_dafny.Array(None, 0), 27)
        d_163_v23_ = nw29_
        d_164_v24_: _dafny.Array
        nw30_ = _dafny.Array(False, 8)
        d_164_v24_ = nw30_
        index22_ = default__.safeIndex(114, (d_163_v23_).length(0))
        (d_163_v23_)[index22_] = (d_164_v24_ if d_137_v3_ else d_164_v24_)
        index23_ = default__.safeIndex(114, (d_163_v23_).length(0))
        (d_163_v23_)[index23_] = d_164_v24_
        d_165_v25_: _dafny.Array
        def lambda9_(d_166_v0_):
            def lambda10_(d_167_i6_):
                return (_dafny.SeqWithoutIsStrInference([d_166_v0_ for d_168_i7_ in range(default__.abs(-562))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uueqdwt")))

            return lambda10_

        init5_ = lambda9_(d_133_v0_)
        nw31_ = _dafny.Array(None, 15)
        for i0_5_ in range(nw31_.length(0)):
            nw31_[i0_5_] = init5_(i0_5_)
        d_165_v25_ = nw31_
        index24_ = default__.safeIndex(195, (d_165_v25_).length(0))
        (d_165_v25_)[index24_] = d_145_v9_
        index25_ = default__.safeIndex(195, (d_165_v25_).length(0))
        (d_165_v25_)[index25_] = d_145_v9_
        (d_142_globalState_).f9 = d_140_v6_
        d_169_i8_: int
        d_169_i8_ = 0
        with _dafny.label("1"):
            while not(not (d_137_v3_) or ((d_156_v17_) == (d_157_v18_))):
                with _dafny.c_label("1"):
                    if (d_169_i8_) >= (100):
                        raise _dafny.Break("1")
                    d_169_i8_ = (d_169_i8_) + (1)
                    d_133_v0_ = _dafny.CodePoint('h')
                    d_170_v26_: _dafny.MultiSet
                    d_170_v26_ = _dafny.MultiSet([d_138_v4_])
                    index26_ = default__.safeIndex(640, (d_134_v1_).length(0))
                    (d_134_v1_)[index26_] = default__.fm1(d_137_v3_, (d_170_v26_).cardinality, d_157_v18_, d_142_globalState_)
                    d_171_v27_: _dafny.MultiSet
                    d_171_v27_ = _dafny.MultiSet([d_133_v0_, d_133_v0_])
                    index27_ = default__.safeIndex(640, (d_134_v1_).length(0))
                    (d_134_v1_)[index27_] = default__.fm1((d_171_v27_).issubset(d_171_v27_), ((0) - (d_156_v17_)) * (d_157_v18_), d_156_v17_, d_142_globalState_)
                    d_172_v28_: _dafny.Array
                    def lambda11_(d_173_v18_):
                        def lambda12_(d_174_i9_):
                            return (d_174_i9_) * (d_173_v18_)

                        return lambda12_

                    init6_ = lambda11_(d_157_v18_)
                    nw32_ = _dafny.Array(None, 24)
                    for i0_6_ in range(nw32_.length(0)):
                        nw32_[i0_6_] = init6_(i0_6_)
                    d_172_v28_ = nw32_
                    d_175_v29_: _dafny.Map
                    d_175_v29_ = _dafny.Map({d_137_v3_: d_172_v28_})
                    d_176_v30_: _dafny.Map
                    d_176_v30_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_137_v3_]): d_172_v28_})
                    d_177_v31_: D0
                    d_177_v31_ = D0_DC0(_dafny.Map({d_140_v6_: d_172_v28_}))
                    d_178_v32_: _dafny.Array
                    nw33_ = _dafny.Array(None, 22)
                    nw33_[int(0)] = (_dafny.Map({d_140_v6_: ((d_175_v29_)[d_137_v3_] if (d_137_v3_) in (d_175_v29_) else d_172_v28_)})) | (d_176_v30_)
                    nw33_[int(1)] = (d_176_v30_) | (d_176_v30_)
                    nw33_[int(2)] = (_dafny.Map({d_140_v6_: d_172_v28_}) if d_137_v3_ else d_176_v30_)
                    nw33_[int(3)] = d_176_v30_
                    nw33_[int(4)] = d_176_v30_
                    nw33_[int(5)] = d_176_v30_
                    nw33_[int(6)] = d_176_v30_
                    nw33_[int(7)] = d_176_v30_
                    nw33_[int(8)] = d_176_v30_
                    nw33_[int(9)] = d_176_v30_
                    nw33_[int(10)] = d_176_v30_
                    nw33_[int(11)] = (d_177_v31_).cf0
                    nw33_[int(12)] = d_176_v30_
                    nw33_[int(13)] = d_176_v30_
                    nw33_[int(14)] = (d_176_v30_) | (d_176_v30_)
                    nw33_[int(15)] = d_176_v30_
                    nw33_[int(16)] = d_176_v30_
                    nw33_[int(17)] = _dafny.Map({d_140_v6_: d_172_v28_})
                    nw33_[int(18)] = ((d_177_v31_).cf0) | (d_176_v30_)
                    nw33_[int(19)] = (d_177_v31_).cf0
                    nw33_[int(20)] = (d_176_v30_).set(d_140_v6_, d_172_v28_)
                    nw33_[int(21)] = d_176_v30_
                    d_178_v32_ = nw33_
                    index28_ = default__.safeIndex(845, (d_178_v32_).length(0))
                    (d_178_v32_)[index28_] = d_176_v30_
                    d_179_v33_: _dafny.Map
                    d_179_v33_ = _dafny.Map({d_137_v3_: d_176_v30_})
                    index29_ = default__.safeIndex(845, (d_178_v32_).length(0))
                    (d_178_v32_)[index29_] = (d_176_v30_ if d_137_v3_ else ((d_179_v33_)[d_137_v3_] if (d_137_v3_) in (d_179_v33_) else d_176_v30_))
                    (d_142_globalState_).f8 = d_157_v18_
                    pass
            pass
        d_180_i10_: int
        d_180_i10_ = 0
        with _dafny.label("2"):
            while ((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))]) <= ((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))]):
                with _dafny.c_label("2"):
                    if (d_180_i10_) >= (100):
                        raise _dafny.Break("2")
                    d_180_i10_ = (d_180_i10_) + (1)
                    d_181_v34_: C0
                    nw34_ = C0()
                    nw34_.ctor__(False, d_141_v7_)
                    d_181_v34_ = nw34_
                    index30_ = default__.safeIndex(51, (d_134_v1_).length(0))
                    (d_134_v1_)[index30_] = d_139_v5_
                    index31_ = default__.safeIndex(51, (d_134_v1_).length(0))
                    (d_134_v1_)[index31_] = -698
                    d_182_v35_: _dafny.Array
                    nw35_ = _dafny.Array(None, 29)
                    nw35_[int(0)] = d_144_v8_
                    nw35_[int(1)] = (_dafny.MultiSet([d_181_v34_.f19, d_137_v3_, d_181_v34_.f19, ((d_141_v7_)[d_137_v3_] if (d_137_v3_) in (d_141_v7_) else d_137_v3_)])) - (_dafny.MultiSet([d_181_v34_.f19, d_137_v3_]))
                    nw35_[int(2)] = _dafny.MultiSet([default__.fm0(d_137_v3_, d_137_v3_, default__.fm4(d_181_v34_.f19, d_137_v3_, not(d_137_v3_), len(d_138_v4_), d_142_globalState_), d_142_globalState_), d_181_v34_.f19])
                    nw35_[int(3)] = _dafny.MultiSet(d_140_v6_)
                    nw35_[int(4)] = d_144_v8_
                    nw35_[int(5)] = (d_144_v8_).set(d_137_v3_, default__.abs((d_134_v1_)[default__.safeIndex(51, (d_134_v1_).length(0))]))
                    nw35_[int(6)] = (d_144_v8_) - (_dafny.MultiSet([d_137_v3_]))
                    nw35_[int(7)] = d_144_v8_
                    nw35_[int(8)] = (d_144_v8_).set(d_181_v34_.f19, default__.abs(d_157_v18_))
                    nw35_[int(9)] = d_144_v8_
                    nw35_[int(10)] = d_144_v8_
                    nw35_[int(11)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_181_v34_.f19]))
                    nw35_[int(12)] = d_144_v8_
                    nw35_[int(13)] = (d_144_v8_).set(False, default__.abs((d_134_v1_)[default__.safeIndex(51, (d_134_v1_).length(0))]))
                    nw35_[int(14)] = d_144_v8_
                    nw35_[int(15)] = (d_144_v8_) - (_dafny.MultiSet([d_181_v34_.f19, d_181_v34_.f19]))
                    nw35_[int(16)] = d_144_v8_
                    nw35_[int(17)] = _dafny.MultiSet([d_181_v34_.f19, d_181_v34_.f19, d_181_v34_.f19])
                    nw35_[int(18)] = d_144_v8_
                    nw35_[int(19)] = (d_144_v8_) | (_dafny.MultiSet([d_137_v3_]))
                    nw35_[int(20)] = _dafny.MultiSet([d_137_v3_])
                    nw35_[int(21)] = d_144_v8_
                    nw35_[int(22)] = (d_144_v8_).set(d_137_v3_, default__.abs(d_156_v17_))
                    nw35_[int(23)] = d_144_v8_
                    nw35_[int(24)] = d_144_v8_
                    nw35_[int(25)] = _dafny.MultiSet([d_181_v34_.f19, d_137_v3_, d_137_v3_])
                    nw35_[int(26)] = d_144_v8_
                    nw35_[int(27)] = (d_144_v8_).set(default__.fm0(d_137_v3_, d_137_v3_, (d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))], d_142_globalState_), default__.abs((d_134_v1_)[default__.safeIndex(51, (d_134_v1_).length(0))]))
                    nw35_[int(28)] = d_144_v8_
                    d_182_v35_ = nw35_
                    d_182_v35_ = (D1_DC3(d_182_v35_)).cf2
                    (d_181_v34_).f19 = d_181_v34_.f19
                    pass
            pass
        source1_ = D0_DC1()
        if source1_.is_DC1:
            d_183_v36_: _dafny.Map
            d_183_v36_ = _dafny.Map({default__.fm0(not(d_137_v3_), d_137_v3_, (d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))], d_142_globalState_): default__.fm5(d_144_v8_, d_142_globalState_)})
            d_184_v37_: _dafny.Set
            d_184_v37_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odbnb"))})
            d_185_v38_: D1
            d_185_v38_ = D1_DC4(((d_183_v36_)[not(d_137_v3_)] if (not(d_137_v3_)) in (d_183_v36_) else d_184_v37_))
            d_186_v39_: _dafny.Seq
            d_186_v39_ = _dafny.SeqWithoutIsStrInference([40])
            d_187_v40_: _dafny.Seq
            d_187_v40_ = _dafny.SeqWithoutIsStrInference([d_186_v39_, default__.fm7(d_142_globalState_)])
            pat_let_tv2_ = d_184_v37_
            def iife11_(_pat_let4_0):
                def iife12_(d_188_dt__update__tmp_h0_):
                    def iife13_(_pat_let5_0):
                        def iife14_(d_189_dt__update_hcf3_h0_):
                            return D1_DC4(d_189_dt__update_hcf3_h0_)
                        return iife14_(_pat_let5_0)
                    return iife13_(pat_let_tv2_)
                return iife12_(_pat_let4_0)
            d_185_v38_ = iife11_(default__.fm6(d_187_v40_, d_137_v3_, d_137_v3_, d_142_globalState_))
            d_137_v3_ = (d_137_v3_) and (d_137_v3_)
            d_190_v41_: _dafny.Map
            d_190_v41_ = _dafny.Map({default__.fm7(d_142_globalState_): d_137_v3_})
            d_190_v41_ = (d_190_v41_).set(d_186_v39_, (d_137_v3_) == (d_137_v3_))
            index32_ = default__.safeIndex(386, (d_164_v24_).length(0))
            (d_164_v24_)[index32_] = d_137_v3_
            index33_ = default__.safeIndex(386, (d_164_v24_).length(0))
            (d_164_v24_)[index33_] = (len((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))])) >= (len(d_138_v4_))
        elif source1_.is_DC0:
            d_191___mcc_h0_ = source1_.cf0
            d_192_cf0_ = d_191___mcc_h0_
            (d_142_globalState_).f8 = 551
            d_193_v42_: C0
            nw36_ = C0()
            nw36_.ctor__(not (d_137_v3_) or (d_137_v3_), ((d_141_v7_)) | (d_141_v7_))
            d_193_v42_ = nw36_
            d_193_v42_ = d_193_v42_
            (d_142_globalState_).f8 = d_139_v5_
            arr0_ = (d_163_v23_)[default__.safeIndex(114, (d_163_v23_).length(0))]
            index34_ = default__.safeIndex(386, ((d_163_v23_)[default__.safeIndex(114, (d_163_v23_).length(0))]).length(0))
            arr0_[index34_] = d_137_v3_
            arr1_ = (d_163_v23_)[default__.safeIndex(114, (d_163_v23_).length(0))]
            index35_ = default__.safeIndex(386, ((d_163_v23_)[default__.safeIndex(114, (d_163_v23_).length(0))]).length(0))
            arr1_[index35_] = (d_157_v18_) < (-638)
        elif True:
            d_194___mcc_h1_ = source1_.cf1
            d_195_cf1_ = d_194___mcc_h1_
            d_196_v43_: _dafny.Array
            nw37_ = _dafny.Array(None, 21)
            d_196_v43_ = nw37_
            d_196_v43_ = d_196_v43_
            if d_137_v3_:
                (d_142_globalState_).f7 = d_139_v5_
                index36_ = default__.safeIndex(122, (d_164_v24_).length(0))
                (d_164_v24_)[index36_] = d_137_v3_
                d_197_v44_: _dafny.Map
                d_197_v44_ = _dafny.Map({d_137_v3_: _dafny.CodePoint('g')})
                d_198_v45_: _dafny.Set
                d_198_v45_ = _dafny.Set({d_139_v5_, d_139_v5_, default__.fm1(d_137_v3_, d_139_v5_, len(d_197_v44_), d_142_globalState_), default__.safeDivisionInt(-178, d_156_v17_), d_157_v18_})
                index37_ = default__.safeIndex(122, (d_164_v24_).length(0))
                rhs32_ = (d_139_v5_ if default__.fm0(d_137_v3_, d_137_v3_, d_145_v9_, d_142_globalState_) else d_157_v18_)
                rhs33_ = ((d_140_v6_ if (d_156_v17_) <= (d_157_v18_) else default__.fm8(d_157_v18_, d_137_v3_, d_142_globalState_))).set(default__.safeIndex(d_139_v5_, len((d_140_v6_ if (d_156_v17_) <= (d_157_v18_) else default__.fm8(d_157_v18_, d_137_v3_, d_142_globalState_)))), not (d_137_v3_) or (not(False)))
                rhs34_ = (not(((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))]) <= ((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))])) if d_137_v3_ else (-337) > ((_dafny.MultiSet([(d_163_v23_)[default__.safeIndex(114, (d_163_v23_).length(0))], d_164_v24_, d_164_v24_])).cardinality))
                rhs35_ = d_198_v45_
                lhs16_ = d_142_globalState_
                lhs17_ = d_142_globalState_
                lhs18_ = d_164_v24_
                lhs19_ = default__.safeIndex(122, (d_164_v24_).length(0))
                lhs16_.f8 = rhs32_
                lhs17_.f9 = rhs33_
                lhs18_[lhs19_] = rhs34_
                d_198_v45_ = rhs35_
                d_199_v46_: int
                d_200_v47_: int
                out4_: int
                out5_: int
                out4_, out5_ = default__.m0(d_142_globalState_)
                d_199_v46_ = out4_
                d_200_v47_ = out5_
                d_201_v48_: C0
                nw38_ = C0()
                nw38_.ctor__(not(True), d_141_v7_)
                d_201_v48_ = nw38_
                d_201_v48_ = d_201_v48_
                d_137_v3_ = ((d_201_v48_).fm2(d_133_v0_, default__.fm0(d_201_v48_.f19, (d_164_v24_)[default__.safeIndex(122, (d_164_v24_).length(0))], d_145_v9_, d_142_globalState_), d_142_globalState_)) >= (d_139_v5_)
            elif True:
                d_202_v49_: int
                d_203_v50_: int
                out6_: int
                out7_: int
                out6_, out7_ = default__.m0(d_142_globalState_)
                d_202_v49_ = out6_
                d_203_v50_ = out7_
                index38_ = default__.safeIndex(830, (d_164_v24_).length(0))
                (d_164_v24_)[index38_] = (default__.fm8(d_156_v17_, d_137_v3_, d_142_globalState_)) < (default__.fm8(d_202_v49_, d_137_v3_, d_142_globalState_))
                index39_ = default__.safeIndex(830, (d_164_v24_).length(0))
                (d_164_v24_)[index39_] = (not(d_137_v3_)) and (d_137_v3_)
                d_204_v51_: _dafny.Map
                d_204_v51_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mnvbm")): d_203_v50_})
                (d_142_globalState_).f7 = ((d_204_v51_)[d_145_v9_] if (d_145_v9_) in (d_204_v51_) else d_203_v50_)
                d_205_v52_: _dafny.Map
                d_205_v52_ = _dafny.Map({d_202_v49_: d_202_v49_})
                d_206_v53_: C0
                nw39_ = C0()
                nw39_.ctor__(d_137_v3_, default__.fm9(((d_205_v52_)[(len((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))]))] if ((len((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))]))) in (d_205_v52_) else d_202_v49_), d_139_v5_, len(d_145_v9_), d_142_globalState_))
                d_206_v53_ = nw39_
                d_207_v54_: _dafny.Map
                d_207_v54_ = _dafny.Map({d_134_v1_: (d_203_v50_) - (d_139_v5_)})
                d_207_v54_ = (d_207_v54_).set(d_134_v1_, (0) - (d_202_v49_))
            d_133_v0_ = d_133_v0_
            d_137_v3_ = not(True)
        d_208_i11_: int
        d_208_i11_ = 0
        with _dafny.label("3"):
            while d_137_v3_:
                with _dafny.c_label("3"):
                    if (d_208_i11_) >= (100):
                        raise _dafny.Break("3")
                    d_208_i11_ = (d_208_i11_) + (1)
                    rhs36_ = not (d_137_v3_) or (d_137_v3_)
                    rhs37_ = d_156_v17_
                    rhs38_ = 744
                    lhs20_ = d_142_globalState_
                    lhs21_ = d_142_globalState_
                    d_137_v3_ = rhs36_
                    lhs20_.f8 = rhs37_
                    lhs21_.f8 = rhs38_
                    if default__.fm0(d_137_v3_, d_137_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xahmwdx")), d_142_globalState_):
                        d_209_v55_: _dafny.Map
                        d_209_v55_ = _dafny.Map({(len(d_145_v9_)) * (d_157_v18_): _dafny.CodePoint('p')})
                        d_209_v55_ = (d_209_v55_).set(d_139_v5_, d_133_v0_)
                        d_210_v56_: C0
                        nw40_ = C0()
                        nw40_.ctor__(d_137_v3_, d_141_v7_)
                        d_210_v56_ = nw40_
                        d_211_v57_: int
                        d_212_v58_: int
                        out8_: int
                        out9_: int
                        out8_, out9_ = default__.m0(d_142_globalState_)
                        d_211_v57_ = out8_
                        d_212_v58_ = out9_
                        d_213_v59_: bool
                        d_213_v59_ = not(d_137_v3_)
                        (d_210_v56_).f19 = (d_213_v59_)
                        d_214_v60_: _dafny.Seq
                        d_214_v60_ = _dafny.SeqWithoutIsStrInference([((d_144_v8_)[d_210_v56_.f19] if (d_210_v56_.f19) in (d_144_v8_) else d_157_v18_), d_139_v5_])
                        d_215_v61_: _dafny.Map
                        d_215_v61_ = _dafny.Map({d_210_v56_: d_214_v60_})
                        d_215_v61_ = (d_215_v61_).set(d_210_v56_, d_214_v60_)
                    elif True:
                        d_216_v62_: C0
                        nw41_ = C0()
                        nw41_.ctor__(d_137_v3_, (d_141_v7_) | (d_141_v7_))
                        d_216_v62_ = nw41_
                        d_141_v7_ = (d_141_v7_).set(default__.fm0(d_137_v3_, d_137_v3_, (d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))], d_142_globalState_), not(d_137_v3_))
                        d_217_v63_: _dafny.Map
                        d_217_v63_ = _dafny.Map({default__.safeDivisionInt(default__.fm1(d_137_v3_, d_157_v18_, (0) - (d_139_v5_), d_142_globalState_), d_156_v17_): default__.safeDivisionInt(d_156_v17_, d_157_v18_)})
                        d_217_v63_ = (d_217_v63_).set((d_157_v18_) * (d_157_v18_), (_dafny.MultiSet([(0) - (d_156_v17_), (d_216_v62_).fm2(d_133_v0_, d_216_v62_.f19, d_142_globalState_), d_139_v5_])).cardinality)
                        d_141_v7_ = (d_141_v7_).set(False, d_137_v3_)
                        d_218_v64_: C0
                        nw42_ = C0()
                        nw42_.ctor__((d_216_v62_.f19) == ((d_140_v6_)[default__.safeIndex(d_157_v18_, len(d_140_v6_))]), d_141_v7_)
                        d_218_v64_ = nw42_
                    index40_ = default__.safeIndex(844, (d_134_v1_).length(0))
                    (d_134_v1_)[index40_] = d_157_v18_
                    index41_ = default__.safeIndex(844, (d_134_v1_).length(0))
                    rhs39_ = len((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))])
                    rhs40_ = default__.safeDivisionInt(d_156_v17_, d_156_v17_)
                    rhs41_ = d_139_v5_
                    lhs22_ = d_142_globalState_
                    lhs23_ = d_142_globalState_
                    lhs24_ = d_134_v1_
                    lhs25_ = default__.safeIndex(844, (d_134_v1_).length(0))
                    lhs22_.f8 = rhs39_
                    lhs23_.f7 = rhs40_
                    lhs24_[lhs25_] = rhs41_
                    d_219_v65_: _dafny.MultiSet
                    d_219_v65_ = _dafny.MultiSet([((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))]).set(default__.safeIndex(d_156_v17_, len((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))])), d_133_v0_), d_145_v9_])
                    d_220_v66_: _dafny.Seq
                    d_220_v66_ = _dafny.SeqWithoutIsStrInference([((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))]).set(default__.safeIndex(d_139_v5_, len((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))])), d_133_v0_)])
                    d_137_v3_ = (_dafny.MultiSet(d_220_v66_)).ispropersubset((d_219_v65_) | (d_219_v65_))
                    pass
            pass
        d_137_v3_ = d_137_v3_
        if d_137_v3_:
            if (False) and (d_137_v3_):
                d_221_v67_: C0
                nw43_ = C0()
                nw43_.ctor__(d_137_v3_, d_141_v7_)
                d_221_v67_ = nw43_
                d_221_v67_ = d_221_v67_
                def lambda13_(d_222_v5_):
                    def lambda14_(d_223_i12_):
                        return (d_223_i12_) * (d_222_v5_)

                    return lambda14_

                init7_ = lambda13_(d_139_v5_)
                nw44_ = _dafny.Array(None, 15)
                for i0_7_ in range(nw44_.length(0)):
                    nw44_[i0_7_] = init7_(i0_7_)
                d_134_v1_ = nw44_
                (d_142_globalState_).f8 = d_157_v18_
                arr2_ = (d_163_v23_)[default__.safeIndex(114, (d_163_v23_).length(0))]
                index42_ = default__.safeIndex(41, ((d_163_v23_)[default__.safeIndex(114, (d_163_v23_).length(0))]).length(0))
                arr2_[index42_] = (d_139_v5_) != (d_156_v17_)
                d_224_v68_: _dafny.Set
                d_224_v68_ = _dafny.Set({d_156_v17_, 151})
                d_225_v70_: _dafny.Map
                def iife15_():
                    coll3_ = _dafny.Set()
                    compr_3_: int
                    for compr_3_ in _dafny.IntegerRange(168, 106):
                        d_226_v69_: int = compr_3_
                        if ((168) <= (d_226_v69_)) and ((d_226_v69_) < (106)):
                            coll3_ = coll3_.union(_dafny.Set([(d_226_v69_) * (d_139_v5_)]))
                    return _dafny.Set(coll3_)
                d_225_v70_ = _dafny.Map({iife15_()
                : d_140_v6_})
                arr3_ = (d_163_v23_)[default__.safeIndex(114, (d_163_v23_).length(0))]
                index43_ = default__.safeIndex(41, ((d_163_v23_)[default__.safeIndex(114, (d_163_v23_).length(0))]).length(0))
                arr3_[index43_] = (d_224_v68_) not in (d_225_v70_)
                d_221_v67_ = d_221_v67_
            elif True:
                d_227_v71_: C0
                nw45_ = C0()
                nw45_.ctor__(((d_141_v7_)[d_137_v3_] if (d_137_v3_) in (d_141_v7_) else d_137_v3_), d_141_v7_)
                d_227_v71_ = nw45_
                d_228_v72_: D5
                d_228_v72_ = D5_DC9(d_227_v71_)
                pat_let_tv3_ = d_227_v71_
                def iife16_(_pat_let6_0):
                    def iife17_(d_229_dt__update__tmp_h1_):
                        def iife18_(_pat_let7_0):
                            def iife19_(d_230_dt__update_hcf8_h0_):
                                return D5_DC9(d_230_dt__update_hcf8_h0_)
                            return iife19_(_pat_let7_0)
                        return iife18_(pat_let_tv3_)
                    return iife17_(_pat_let6_0)
                d_227_v71_ = (iife16_(d_228_v72_)).cf8
                d_231_v73_: _dafny.Array
                nw46_ = _dafny.Array(None, 8)
                nw46_[int(0)] = d_144_v8_
                nw46_[int(1)] = d_144_v8_
                nw46_[int(2)] = d_144_v8_
                nw46_[int(3)] = d_144_v8_
                nw46_[int(4)] = d_144_v8_
                nw46_[int(5)] = d_144_v8_
                nw46_[int(6)] = d_144_v8_
                nw46_[int(7)] = _dafny.MultiSet([False])
                d_231_v73_ = nw46_
                nw47_ = _dafny.Array(_dafny.MultiSet({}), 19)
                d_231_v73_ = nw47_
                d_232_v74_: _dafny.Seq
                d_232_v74_ = _dafny.SeqWithoutIsStrInference([(d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))], d_145_v9_])
                d_145_v9_ = ((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))]) + (((d_232_v74_)[default__.safeIndex(542, len(d_232_v74_))]) + (d_145_v9_))
                d_227_v71_ = d_227_v71_
                index44_ = default__.safeIndex(217, (d_134_v1_).length(0))
                (d_134_v1_)[index44_] = -834
                d_233_v75_: _dafny.Set
                d_233_v75_ = _dafny.Set({(d_227_v71_).fm2(d_133_v0_, not(not(d_227_v71_.f19)), d_142_globalState_), d_157_v18_})
                d_234_v76_: _dafny.MultiSet
                d_234_v76_ = _dafny.MultiSet([(d_144_v8_).cardinality])
                d_235_v77_: _dafny.Map
                d_235_v77_ = _dafny.Map({(d_234_v76_).set(len(d_145_v9_), default__.abs(d_139_v5_)): d_157_v18_})
                d_236_v78_: _dafny.Seq
                d_236_v78_ = _dafny.SeqWithoutIsStrInference([d_156_v17_])
                index45_ = default__.safeIndex(217, (d_134_v1_).length(0))
                rhs42_ = (d_236_v78_)[default__.safeIndex((0) - (d_157_v18_), len(d_236_v78_))]
                rhs43_ = (d_233_v75_ if d_137_v3_ else d_233_v75_)
                rhs44_ = default__.fm10(d_236_v78_, default__.fm11(d_133_v0_, d_227_v71_.f19, d_156_v17_, d_157_v18_, d_142_globalState_), d_142_globalState_)
                rhs45_ = (d_235_v77_) | (d_235_v77_)
                rhs46_ = d_137_v3_
                lhs26_ = d_134_v1_
                lhs27_ = default__.safeIndex(217, (d_134_v1_).length(0))
                lhs26_[lhs27_] = rhs42_
                d_233_v75_ = rhs43_
                d_133_v0_ = rhs44_
                d_235_v77_ = rhs45_
                d_137_v3_ = rhs46_
            d_237_v79_: _dafny.Seq
            d_237_v79_ = _dafny.SeqWithoutIsStrInference([d_156_v17_])
            d_238_v80_: _dafny.Seq
            d_238_v80_ = _dafny.SeqWithoutIsStrInference([d_237_v79_, _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_156_v17_, 362])).cardinality for d_239_i13_ in range(default__.abs(-204))])])
            d_138_v4_ = default__.fm12(d_238_v80_, (d_139_v5_) >= (d_139_v5_), ((d_141_v7_)[not(False)] if (not(False)) in (d_141_v7_) else d_137_v3_), d_142_globalState_)
            if not (d_137_v3_) or ((d_137_v3_ if (d_140_v6_)[default__.safeIndex(d_157_v18_, len(d_140_v6_))] else d_137_v3_)):
                d_240_v81_: _dafny.Array
                nw48_ = _dafny.Array(_dafny.Set({}), 24)
                d_240_v81_ = nw48_
                d_240_v81_ = d_240_v81_
                d_163_v23_ = d_163_v23_
                d_241_v82_: _dafny.Seq
                d_241_v82_ = _dafny.SeqWithoutIsStrInference([d_164_v24_, d_164_v24_, (d_163_v23_)[default__.safeIndex(114, (d_163_v23_).length(0))]])
                d_241_v82_ = (d_241_v82_) + (d_241_v82_)
                index46_ = default__.safeIndex(277, (d_164_v24_).length(0))
                (d_164_v24_)[index46_] = not(True)
                index47_ = default__.safeIndex(277, (d_164_v24_).length(0))
                (d_164_v24_)[index47_] = (((d_141_v7_)[d_137_v3_] if (d_137_v3_) in (d_141_v7_) else d_137_v3_)) and (d_137_v3_)
                (d_142_globalState_).f8 = d_156_v17_
            elif True:
                d_242_v83_: C0
                nw49_ = C0()
                nw49_.ctor__(d_137_v3_, d_141_v7_)
                d_242_v83_ = nw49_
                d_242_v83_ = d_242_v83_
                (d_142_globalState_).f7 = default__.safeModuloInt(d_139_v5_, d_157_v18_)
                d_243_v84_: _dafny.Seq
                d_243_v84_ = _dafny.SeqWithoutIsStrInference([d_242_v83_])
                d_242_v83_ = (d_243_v84_)[default__.safeIndex((d_139_v5_) - (d_139_v5_), len(d_243_v84_))]
                (d_142_globalState_).f8 = len(_dafny.SeqWithoutIsStrInference([d_242_v83_.f19, not(d_137_v3_), d_137_v3_, (d_242_v83_.f19 if d_137_v3_ else d_137_v3_)]))
                index48_ = default__.safeIndex(721, (d_164_v24_).length(0))
                (d_164_v24_)[index48_] = d_242_v83_.f19
                d_244_v85_: _dafny.Map
                d_244_v85_ = _dafny.Map({442: d_137_v3_})
                index49_ = default__.safeIndex(721, (d_164_v24_).length(0))
                rhs47_ = d_242_v83_.f19
                rhs48_ = d_244_v85_
                lhs28_ = d_164_v24_
                lhs29_ = default__.safeIndex(721, (d_164_v24_).length(0))
                lhs28_[lhs29_] = rhs47_
                d_244_v85_ = rhs48_
            d_245_v87_: _dafny.MultiSet
            d_245_v87_ = _dafny.MultiSet([(0) - (len(d_141_v7_)), 693, d_156_v17_])
            d_246_v88_: _dafny.Set
            d_246_v88_ = _dafny.Set({_dafny.MultiSet([d_156_v17_]), d_245_v87_, d_245_v87_, d_245_v87_, d_245_v87_})
            d_247_v89_: D5
            def iife20_():
                coll4_ = _dafny.Map()
                compr_4_: _dafny.MultiSet
                for compr_4_ in (d_246_v88_).Elements:
                    d_248_v86_: _dafny.MultiSet = compr_4_
                    if (d_248_v86_) in (d_246_v88_):
                        coll4_[d_248_v86_] = d_145_v9_
                return _dafny.Map(coll4_)
            d_247_v89_ = D5_DC10(_dafny.Set({d_139_v5_, d_139_v5_, d_157_v18_}), iife20_()
)
            d_249_v90_: _dafny.Seq
            d_249_v90_ = _dafny.SeqWithoutIsStrInference([(d_247_v89_ if d_137_v3_ else d_247_v89_), d_247_v89_, d_247_v89_])
            d_250_v91_: _dafny.Set
            d_250_v91_ = _dafny.Set({d_139_v5_})
            d_251_v92_: _dafny.Map
            d_251_v92_ = _dafny.Map({d_245_v87_: (d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))]})
            rhs49_ = d_134_v1_
            rhs50_ = (_dafny.SeqWithoutIsStrInference([d_247_v89_, D5_DC10(d_250_v91_, d_251_v92_)])).set(default__.safeIndex(d_156_v17_, len(_dafny.SeqWithoutIsStrInference([d_247_v89_, D5_DC10(d_250_v91_, d_251_v92_)]))), d_247_v89_)
            rhs51_ = d_133_v0_
            d_134_v1_ = rhs49_
            d_249_v90_ = rhs50_
            d_133_v0_ = rhs51_
            d_252_v93_: _dafny.Seq
            d_252_v93_ = d_145_v9_
            d_253_v94_: C0
            nw50_ = C0()
            nw50_.ctor__(default__.fm0(d_137_v3_, d_137_v3_, (d_252_v93_), d_142_globalState_), d_141_v7_)
            d_253_v94_ = nw50_
        elif True:
            if ((d_156_v17_) * (821)) > (d_139_v5_):
                index50_ = default__.safeIndex(458, (d_134_v1_).length(0))
                (d_134_v1_)[index50_] = (d_157_v18_) - (default__.fm1(d_137_v3_, d_157_v18_, (0) - (d_156_v17_), d_142_globalState_))
                d_254_v95_: _dafny.Seq
                d_254_v95_ = _dafny.SeqWithoutIsStrInference([d_134_v1_])
                d_255_v96_: _dafny.Map
                d_255_v96_ = _dafny.Map({d_157_v18_: d_157_v18_})
                index51_ = default__.safeIndex(458, (d_134_v1_).length(0))
                rhs52_ = (len(d_145_v9_) if d_137_v3_ else ((0) - (d_156_v17_)) - (d_156_v17_))
                rhs53_ = (d_254_v95_)[default__.safeIndex(len(_dafny.Set({d_137_v3_, d_137_v3_})), len(d_254_v95_))]
                rhs54_ = default__.safeDivisionInt((0) - (d_157_v18_), default__.safeModuloInt(((d_255_v96_)[d_139_v5_] if (d_139_v5_) in (d_255_v96_) else d_156_v17_), d_139_v5_))
                lhs30_ = d_142_globalState_
                lhs31_ = d_134_v1_
                lhs32_ = default__.safeIndex(458, (d_134_v1_).length(0))
                lhs30_.f7 = rhs52_
                d_134_v1_ = rhs53_
                lhs31_[lhs32_] = rhs54_
                d_256_v97_: _dafny.MultiSet
                d_256_v97_ = _dafny.MultiSet([d_157_v18_, (d_134_v1_)[default__.safeIndex(458, (d_134_v1_).length(0))]])
                d_257_v98_: _dafny.Map
                d_257_v98_ = _dafny.Map({d_256_v97_: len((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))])})
                d_258_v99_: _dafny.Map
                d_258_v99_ = _dafny.Map({(d_134_v1_)[default__.safeIndex(458, (d_134_v1_).length(0))]: d_257_v98_})
                d_259_v100_: _dafny.Seq
                d_259_v100_ = _dafny.SeqWithoutIsStrInference([d_141_v7_, d_141_v7_, d_141_v7_])
                d_260_v102_: C0
                nw51_ = C0()
                nw51_.ctor__(d_137_v3_, d_141_v7_)
                d_260_v102_ = nw51_
                d_261_v103_: _dafny.Map
                d_261_v103_ = _dafny.Map({d_260_v102_: d_137_v3_})
                d_262_v104_: _dafny.Array
                nw52_ = _dafny.Array(None, 8)
                nw52_[int(0)] = ((d_257_v98_).set((d_256_v97_).set((d_134_v1_)[default__.safeIndex(458, (d_134_v1_).length(0))], default__.abs(d_157_v18_)), (d_134_v1_)[default__.safeIndex(458, (d_134_v1_).length(0))]) if d_137_v3_ else ((d_258_v99_)[len((d_259_v100_)[default__.safeIndex(d_157_v18_, len(d_259_v100_))])] if (len((d_259_v100_)[default__.safeIndex(d_157_v18_, len(d_259_v100_))])) in (d_258_v99_) else d_257_v98_))
                def iife21_():
                    coll5_ = _dafny.Map()
                    compr_5_: _dafny.MultiSet
                    for compr_5_ in (_dafny.SeqWithoutIsStrInference([d_256_v97_])).Elements:
                        d_263_v101_: _dafny.MultiSet = compr_5_
                        if (d_263_v101_) in (_dafny.SeqWithoutIsStrInference([d_256_v97_])):
                            coll5_[d_263_v101_] = d_139_v5_
                    return _dafny.Map(coll5_)
                nw52_[int(1)] = (iife21_()
                ) | (_dafny.Map({d_256_v97_: d_139_v5_}))
                nw52_[int(2)] = d_257_v98_
                nw52_[int(3)] = (_dafny.Map({d_256_v97_: d_139_v5_})).set(d_256_v97_, len(d_261_v103_))
                nw52_[int(4)] = _dafny.Map({_dafny.MultiSet([d_156_v17_, 274, (d_134_v1_)[default__.safeIndex(458, (d_134_v1_).length(0))], d_156_v17_]): d_156_v17_})
                nw52_[int(5)] = default__.fm13(d_137_v3_, (0) - (d_139_v5_), d_142_globalState_)
                nw52_[int(6)] = d_257_v98_
                nw52_[int(7)] = (d_257_v98_) | (d_257_v98_)
                d_262_v104_ = nw52_
                index52_ = default__.safeIndex(16, (d_262_v104_).length(0))
                (d_262_v104_)[index52_] = d_257_v98_
                index53_ = default__.safeIndex(16, (d_262_v104_).length(0))
                rhs55_ = (0) - ((d_134_v1_)[default__.safeIndex(458, (d_134_v1_).length(0))])
                rhs56_ = default__.fm1(True, d_156_v17_, (d_134_v1_)[default__.safeIndex(458, (d_134_v1_).length(0))], d_142_globalState_)
                rhs57_ = (((d_165_v25_)[default__.safeIndex(195, (d_165_v25_).length(0))]) + (_dafny.SeqWithoutIsStrInference([d_133_v0_ for d_264_i14_ in range(default__.abs(905))]))) + ((d_145_v9_) + (_dafny.SeqWithoutIsStrInference([d_133_v0_ for d_265_i15_ in range(default__.abs(934))])))
                rhs58_ = d_257_v98_
                lhs33_ = d_142_globalState_
                lhs34_ = d_262_v104_
                lhs35_ = default__.safeIndex(16, (d_262_v104_).length(0))
                lhs33_.f8 = rhs55_
                d_157_v18_ = rhs56_
                d_145_v9_ = rhs57_
                lhs34_[lhs35_] = rhs58_
                d_266_v105_: _dafny.Seq
                d_266_v105_ = _dafny.SeqWithoutIsStrInference([d_164_v24_, d_164_v24_, d_164_v24_, d_164_v24_, (d_163_v23_)[default__.safeIndex(114, (d_163_v23_).length(0))]])
                d_266_v105_ = d_266_v105_
                d_267_v106_: _dafny.Seq
                d_267_v106_ = _dafny.SeqWithoutIsStrInference([(d_134_v1_)[default__.safeIndex(458, (d_134_v1_).length(0))], (0) - (d_156_v17_)])
                d_268_v107_: _dafny.Seq
                d_268_v107_ = _dafny.SeqWithoutIsStrInference([d_267_v106_, d_267_v106_, d_267_v106_, _dafny.SeqWithoutIsStrInference([d_156_v17_, len(_dafny.SeqWithoutIsStrInference([(d_134_v1_)[default__.safeIndex(458, (d_134_v1_).length(0))]])), d_139_v5_])])
                d_269_v108_: _dafny.Set
                d_269_v108_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjjbn"))})
                index54_ = default__.safeIndex(195, (d_165_v25_).length(0))
                rhs59_ = d_137_v3_
                rhs60_ = ((d_268_v107_)[default__.safeIndex(d_139_v5_, len(d_268_v107_))]) + (_dafny.SeqWithoutIsStrInference([len(d_269_v108_)]))
                rhs61_ = d_137_v3_
                rhs62_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwwcq"))) + (d_145_v9_)
                lhs36_ = d_260_v102_
                lhs37_ = d_165_v25_
                lhs38_ = default__.safeIndex(195, (d_165_v25_).length(0))
                d_137_v3_ = rhs59_
                d_267_v106_ = rhs60_
                lhs36_.f19 = rhs61_
                lhs37_[lhs38_] = rhs62_
                (d_260_v102_).f19 = (False) == (d_260_v102_.f19)
            elif True:
                d_137_v3_ = d_137_v3_
                d_270_v109_: C0
                nw53_ = C0()
                nw53_.ctor__(d_137_v3_, d_141_v7_)
                d_270_v109_ = nw53_
                index55_ = default__.safeIndex(32, (d_134_v1_).length(0))
                (d_134_v1_)[index55_] = (d_157_v18_) * (d_156_v17_)
                d_271_v110_: _dafny.Map
                d_271_v110_ = _dafny.Map({d_144_v8_: d_156_v17_})
                index56_ = default__.safeIndex(32, (d_134_v1_).length(0))
                (d_134_v1_)[index56_] = (0) - (default__.safeDivisionInt(len(d_271_v110_), default__.safeModuloInt(d_139_v5_, d_157_v18_)))
                (d_270_v109_).f19 = d_270_v109_.f19
                (d_142_globalState_).f7 = ((0) - ((d_134_v1_)[default__.safeIndex(32, (d_134_v1_).length(0))])) - (d_156_v17_)
            d_272_v111_: C0
            nw54_ = C0()
            nw54_.ctor__((d_140_v6_)[default__.safeIndex(d_156_v17_, len(d_140_v6_))], d_141_v7_)
            d_272_v111_ = nw54_
            d_141_v7_ = (d_141_v7_).set(d_137_v3_, d_272_v111_.f19)
            d_273_v112_: _dafny.Map
            d_273_v112_ = _dafny.Map({d_133_v0_: d_272_v111_})
            d_273_v112_ = (((d_273_v112_).set(d_133_v0_, d_272_v111_)).set(_dafny.CodePoint('w'), d_272_v111_) if d_137_v3_ else d_273_v112_)
            d_137_v3_ = d_272_v111_.f19
        d_141_v7_ = (d_141_v7_).set(d_137_v3_, d_137_v3_)
        _dafny.print(_dafny.string_of(d_133_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v1_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_136_v2_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_137_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v4_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_139_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_140_v6_) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_141_v7_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_142_globalState_).f1).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_142_globalState_).f2)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f3) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_globalState_.f4) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u')]): 172}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_142_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_142_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_globalState_.f9) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f11)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_globalState_).f13) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_144_v8_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_145_v9_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_146_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_156_v17_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_157_v18_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_163_v23_)[6])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v24_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_165_v25_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_169_i8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_180_i10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_208_i11_))
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
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)

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

class D0_DC2(D0, NamedTuple('DC2', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(_dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC4(D1, NamedTuple('DC4', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC6(D2, NamedTuple('DC6', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: int(0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC7(D3, NamedTuple('DC7', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: False
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)

class D4_DC8(D4, NamedTuple('DC8', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC10(_dafny.Set({}), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D5_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D5_DC9)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC10(D5, NamedTuple('DC10', [('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC10({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC10) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC11(D5, NamedTuple('DC11', [('cf11', Any), ('cf12', Any), ('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC9(D5, NamedTuple('DC9', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC9({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC9) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)

class D6_DC13(D6, NamedTuple('DC13', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({self.cf17.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC15(None, int(0), int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D7_DC14)

class D7_DC15(D7, NamedTuple('DC15', [('cf19', Any), ('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf23', Any), ('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf28', Any), ('cf29', Any), ('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC14(D7, NamedTuple('DC14', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC14({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC14) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f4: _dafny.Map = _dafny.Map({})
        self.f7: int = int(0)
        self.f8: int = int(0)
        self.f9: _dafny.Seq = _dafny.Seq({})
        self._f0: bool = False
        self._f1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f2: _dafny.Map = _dafny.Map({})
        self._f3: _dafny.Set = _dafny.Set({})
        self._f5: int = int(0)
        self._f6: bool = False
        self._f10: bool = False
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        self._f12: bool = False
        self._f13: _dafny.Map = _dafny.Map({})
        self._f14: int = int(0)
        self._f15: int = int(0)
        self._f16: bool = False
        self._f17: bool = False
        self._f18: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self).f8 = f8
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
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
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
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
    def f18(self):
        return self._f18

class C0:
    def  __init__(self):
        self.f19: bool = False
        self._f20: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f19, f20):
        (self).f19 = f19
        (self)._f20 = f20

    def fm2(self, p0, p1, globalState):
        return (0) - (len((_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))), len((self).f20), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")))}))]))})).intersection((_dafny.Set({689})) | (_dafny.Set({468, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "doisvt")))})))))

    def fm3(self, p0, p1, globalState):
        source2_ = D0_DC1()
        if source2_.is_DC1:
            return D0_DC1()
        elif source2_.is_DC0:
            d_274___mcc_h0_ = source2_.cf0
            d_275_cf0_ = d_274___mcc_h0_
            return D0_DC1()
        elif True:
            d_276___mcc_h1_ = source2_.cf1
            d_277_cf1_ = d_276___mcc_h1_
            return D0_DC1()

    @property
    def f20(self):
        return self._f20
