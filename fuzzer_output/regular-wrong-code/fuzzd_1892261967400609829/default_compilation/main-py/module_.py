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
    def fm0(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(294, 982):
                d_1_v0_: int = compr_0_
                if ((294) <= (d_1_v0_)) and ((d_1_v0_) < (982)):
                    coll0_[(d_1_v0_) + (len(_dafny.Map({True: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: not(True)})) for d_2_i0_ in range(default__.abs(549))]))).cardinality})))] = len(_dafny.Set({False}))
            return _dafny.Map(coll0_)
        if (_dafny.MultiSet([_dafny.Map({(_dafny.MultiSet([True, False, False])).cardinality: len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, False])).cardinality for d_0_i1_ in range(default__.abs(252))]))})])).issubset(_dafny.MultiSet([iife0_()
        ])):
            return True
        elif True:
            return False

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        return default__.safeModuloInt(len((_dafny.Map({False: False})) | (_dafny.Map({True: True}))), ((_dafny.MultiSet([-146])).cardinality) * (-601))

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))

    @staticmethod
    def fm3(p0, p1, globalState):
        return _dafny.Set({True, True, (_dafny.MultiSet([693, 689])).ispropersubset(_dafny.MultiSet([(_dafny.MultiSet([False, False])).cardinality, len(_dafny.SeqWithoutIsStrInference([True, True]))]))})

    @staticmethod
    def fm4(p0, p1, globalState):
        return _dafny.CodePoint('p')

    @staticmethod
    def fm6(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([not((not(True)) == (False)), not(not(True)), (_dafny.SeqWithoutIsStrInference([_dafny.Set({(D8_DC20(False, 962)).cf26, -800})])) <= (_dafny.SeqWithoutIsStrInference([_dafny.Set({-845})])), False, (782) == (905)])

    @staticmethod
    def fm7(p0, globalState):
        return (_dafny.Map({False: 522})) | ((_dafny.Map({not((D3_DC9(True, _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lihiauqtn"))): len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([654])), -799}))}))).cf13): len(_dafny.Map({631: False}))})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([690]))})))

    @staticmethod
    def fm8(globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.Map({(D3_DC9(False, _dafny.Map({70: 280}))).cf13: 299})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: (0) - ((_dafny.MultiSet([True])).cardinality)}) for d_3_i0_ in range(default__.abs(-53))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): len(_dafny.Set({True}))})]))

    @staticmethod
    def fm9(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: _dafny.Seq
            for compr_1_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_4_i0_ in range(default__.abs(-674))]): not(True)})).keys.Elements:
                d_5_v0_: _dafny.Seq = compr_1_
                if (d_5_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_4_i0_ in range(default__.abs(-674))]): not(True)})):
                    coll1_ = coll1_.union(_dafny.Set([d_5_v0_]))
            return _dafny.Set(coll1_)
        return iife1_()
        

    @staticmethod
    def fm10(p0, p1, globalState):
        return _dafny.Map({-122: D3_DC9(False, _dafny.Map({len(_dafny.Map({False: len(_dafny.Set({_dafny.CodePoint('w'), _dafny.CodePoint('k'), _dafny.CodePoint('e')}))})): (0) - (-32)}))})

    @staticmethod
    def fm11(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqgfpgnr"))) for d_6_i0_ in range(default__.abs(333))])

    @staticmethod
    def fm12(p0, p1, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(310, -947):
                d_7_v0_: int = compr_2_
                if ((310) <= (d_7_v0_)) and ((d_7_v0_) < (-947)):
                    coll2_[(d_7_v0_) * ((_dafny.MultiSet([(_dafny.MultiSet([True, False])).cardinality])).cardinality)] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjm"))))
            return _dafny.Map(coll2_)
        return ((_dafny.Map({68: 47})) | (_dafny.Map({len(_dafny.Map({True: True})): len(_dafny.SeqWithoutIsStrInference([79]))}))) | ((_dafny.Map({(0) - (-498): (D0_DC2(len(_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True})), len(_dafny.Map({True: False})))).cf3})) | (iife2_()
        ))

    @staticmethod
    def fm13(globalState):
        return D2_DC4((_dafny.Map({len(_dafny.Map({True: True})): (0) - ((_dafny.MultiSet([(0) - ((_dafny.MultiSet([D3_DC9(False, _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqbk")))]))).cardinality: len(_dafny.SeqWithoutIsStrInference([False, False, True, True, True]))}))])).cardinality)])).cardinality)})) | (_dafny.Map({-7: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnpmcdw")))})))

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        return _dafny.Map({(len(_dafny.Map({False: _dafny.CodePoint('y')}))) - (96): _dafny.CodePoint('o')})

    @staticmethod
    def m0(p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: str = _dafny.CodePoint('D')
        d_8_i0_: int
        d_8_i0_ = 0
        with _dafny.label("0"):
            while p0:
                with _dafny.c_label("0"):
                    if (d_8_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_8_i0_ = (d_8_i0_) + (1)
                    d_9_v0_: C0
                    nw0_ = C0()
                    nw0_.ctor__()
                    d_9_v0_ = nw0_
                    (globalState).f3 = not(p0)
                    d_10_v1_: _dafny.Seq
                    d_10_v1_ = _dafny.SeqWithoutIsStrInference([p0])
                    (globalState).f3 = (p0 if p0 else (_dafny.MultiSet(d_10_v1_)).ispropersubset(_dafny.MultiSet([p0])))
                    d_11_v2_: _dafny.Set
                    d_11_v2_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbq"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjrtckkdd")))})
                    d_12_v3_: int
                    d_12_v3_ = 346
                    d_11_v2_ = default__.fm9(p0, d_12_v3_, globalState)
                    pass
            pass
        d_13_v4_: int
        d_13_v4_ = 346
        d_14_v5_: D0
        d_14_v5_ = D0_DC2(d_13_v4_, d_13_v4_)
        d_15_v6_: _dafny.Seq
        d_15_v6_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dfufllfmk"))), d_13_v4_])
        (globalState).f22 = ((d_14_v5_).cf2) - (len(d_15_v6_))
        d_16_v7_: D2
        d_16_v7_ = D2_DC6()
        d_16_v7_ = d_16_v7_
        source0_ = D8_DC20(p0, d_13_v4_)
        if source0_.is_DC20:
            d_17___mcc_h0_ = source0_.cf25
            d_18___mcc_h1_ = source0_.cf26
            d_19_cf26_ = d_18___mcc_h1_
            d_20_cf25_ = d_17___mcc_h0_
            d_21_v8_: _dafny.Array
            def lambda0_(d_22_p0_):
                def lambda1_(d_23_i1_):
                    return d_22_p0_

                return lambda1_

            init0_ = lambda0_(p0)
            nw1_ = _dafny.Array(None, 18)
            for i0_0_ in range(nw1_.length(0)):
                nw1_[i0_0_] = init0_(i0_0_)
            d_21_v8_ = nw1_
            d_24_v9_: _dafny.Map
            d_24_v9_ = _dafny.Map({d_13_v4_: d_13_v4_})
            d_25_v10_: _dafny.Set
            d_25_v10_ = _dafny.Set({d_24_v9_, (d_24_v9_).set(len(_dafny.SeqWithoutIsStrInference([d_19_cf26_])), d_19_cf26_), d_24_v9_, d_24_v9_, d_24_v9_})
            index0_ = default__.safeIndex(554, (d_21_v8_).length(0))
            (d_21_v8_)[index0_] = (d_25_v10_) == (d_25_v10_)
            index1_ = default__.safeIndex(554, (d_21_v8_).length(0))
            (d_21_v8_)[index1_] = p0
            d_26_v11_: str
            d_26_v11_ = _dafny.CodePoint('l')
            d_27_v12_: _dafny.MultiSet
            d_27_v12_ = _dafny.MultiSet([d_13_v4_, d_13_v4_, d_19_cf26_, d_19_cf26_])
            d_28_v13_: _dafny.Map
            d_28_v13_ = _dafny.Map({d_26_v11_: (d_19_cf26_) * ((d_27_v12_).cardinality)})
            d_28_v13_ = (d_28_v13_).set(d_26_v11_, (d_13_v4_ if p0 else d_13_v4_))
            d_29_v14_: _dafny.Map
            d_29_v14_ = _dafny.Map({(d_13_v4_) != (584): default__.safeModuloInt(d_19_cf26_, d_13_v4_)})
            d_29_v14_ = (d_29_v14_) | (_dafny.Map({(d_21_v8_)[default__.safeIndex(554, (d_21_v8_).length(0))]: d_19_cf26_}))
            d_30_v15_: _dafny.Map
            d_30_v15_ = _dafny.Map({p0: (d_21_v8_)[default__.safeIndex(554, (d_21_v8_).length(0))]})
            d_30_v15_ = (d_30_v15_).set((d_21_v8_)[default__.safeIndex(554, (d_21_v8_).length(0))], p0)
        elif source0_.is_DC19:
            d_31___mcc_h2_ = source0_.cf24
            d_32_cf24_ = d_31___mcc_h2_
            (globalState).f14 = (p0 if False else (p0 if p0 else p0))
            d_33_v16_: str
            d_33_v16_ = _dafny.CodePoint('a')
            d_34_v17_: _dafny.Map
            d_34_v17_ = _dafny.Map({d_33_v16_: _dafny.CodePoint('k')})
            d_35_v18_: _dafny.Set
            d_35_v18_ = _dafny.Set({p0})
            if (d_13_v4_) <= (default__.safeModuloInt(len(d_34_v17_), len(d_35_v18_))):
                d_36_v19_: C0
                nw2_ = C0()
                nw2_.ctor__()
                d_36_v19_ = nw2_
                d_37_v20_: _dafny.Seq
                d_37_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntnbowsn"))
                (globalState).f21 = d_37_v20_
                d_38_v21_: C0
                nw3_ = C0()
                nw3_.ctor__()
                d_38_v21_ = nw3_
                d_33_v16_ = d_33_v16_
                (globalState).f22 = d_13_v4_
            elif True:
                (globalState).f14 = p0
                d_39_v22_: _dafny.MultiSet
                d_39_v22_ = _dafny.MultiSet([p0])
                d_39_v22_ = d_39_v22_
                (globalState).f25 = (d_13_v4_) * (-892)
                d_40_v23_: _dafny.Array
                nw4_ = _dafny.Array(None, 21)
                d_40_v23_ = nw4_
                d_41_v24_: C0
                nw5_ = C0()
                nw5_.ctor__()
                d_41_v24_ = nw5_
                d_42_v25_: _dafny.Map
                d_42_v25_ = _dafny.Map({d_41_v24_: d_40_v23_})
                d_43_v26_: _dafny.Seq
                d_43_v26_ = _dafny.SeqWithoutIsStrInference([d_40_v23_, ((d_42_v25_)[d_41_v24_] if (d_41_v24_) in (d_42_v25_) else d_40_v23_)])
                (globalState).f3 = (len((d_43_v26_) + (_dafny.SeqWithoutIsStrInference([d_40_v23_])))) <= (d_13_v4_)
                d_44_v27_: _dafny.Map
                d_44_v27_ = _dafny.Map({p0: p0})
                d_45_v28_: D0
                d_45_v28_ = D0_DC1(p0)
                d_46_v29_: _dafny.Seq
                d_46_v29_ = _dafny.SeqWithoutIsStrInference([d_45_v28_])
                d_47_v30_: _dafny.Map
                d_47_v30_ = _dafny.Map({d_44_v27_: d_46_v29_})
                d_48_v31_: _dafny.Map
                d_48_v31_ = _dafny.Map({(D5_DC14(d_47_v30_, d_13_v4_)).cf20: 309})
                d_49_v32_: D3
                d_49_v32_ = D3_DC9(p0, d_48_v31_)
                d_50_v33_: _dafny.Map
                d_50_v33_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_33_v16_ for d_51_i2_ in range(default__.abs(357))])): d_49_v32_})
                d_52_v34_: _dafny.Array
                nw6_ = _dafny.Array(False, 24)
                d_52_v34_ = nw6_
                d_53_v35_: D4
                d_53_v35_ = D4_DC11(d_50_v33_, d_52_v34_)
                d_53_v35_ = d_53_v35_
            d_54_v36_: _dafny.Seq
            d_54_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "llmnukb"))
            d_55_v37_: _dafny.Map
            d_55_v37_ = _dafny.Map({d_33_v16_: d_54_v36_})
            source1_ = D0_DC1(default__.fm0(len(((d_55_v37_)[d_33_v16_] if (d_33_v16_) in (d_55_v37_) else d_54_v36_)), p0, globalState))
            if source1_.is_DC1:
                d_56___mcc_h4_ = source1_.cf1
                d_57_cf1_ = d_56___mcc_h4_
                (globalState).f24 = True
                (globalState).f9 = ((d_13_v4_ if d_57_cf1_ else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwdi"))))) * (d_13_v4_)
                d_58_v38_: _dafny.Seq
                d_58_v38_ = _dafny.SeqWithoutIsStrInference([d_57_cf1_])
                d_59_v39_: _dafny.Seq
                d_59_v39_ = _dafny.SeqWithoutIsStrInference([(d_58_v38_) + (_dafny.SeqWithoutIsStrInference([p0, p0]))])
                r0 = (d_59_v39_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkf"))), len(d_59_v39_))]
                d_60_v40_: _dafny.Map
                d_60_v40_ = _dafny.Map({len(_dafny.Map({d_13_v4_: d_58_v38_})): 177})
                d_61_v41_: D3
                d_61_v41_ = D3_DC9(default__.fm0(d_13_v4_, d_57_cf1_, globalState), (d_60_v40_) | (d_60_v40_))
                d_61_v41_ = d_61_v41_
            elif source1_.is_DC2:
                d_62___mcc_h5_ = source1_.cf2
                d_63___mcc_h6_ = source1_.cf3
                d_64_cf3_ = d_63___mcc_h6_
                d_65_cf2_ = d_62___mcc_h5_
                d_66_v42_: C0
                nw7_ = C0()
                nw7_.ctor__()
                d_66_v42_ = nw7_
                d_67_v43_: _dafny.Set
                d_67_v43_ = _dafny.Set({d_33_v16_, d_33_v16_, d_33_v16_})
                d_68_v44_: _dafny.Array
                nw8_ = _dafny.Array(None, 7)
                nw8_[int(0)] = d_67_v43_
                nw8_[int(1)] = d_67_v43_
                nw8_[int(2)] = _dafny.Set({d_33_v16_, d_33_v16_, d_33_v16_, d_33_v16_})
                nw8_[int(3)] = d_67_v43_
                nw8_[int(4)] = (d_67_v43_ if not(p0) else d_67_v43_)
                nw8_[int(5)] = d_67_v43_
                nw8_[int(6)] = d_67_v43_
                d_68_v44_ = nw8_
                d_69_v45_: _dafny.Array
                def lambda2_(d_70_v36_):
                    def lambda3_(d_71_i3_):
                        return (d_71_i3_) * (len(d_70_v36_))

                    return lambda3_

                init1_ = lambda2_(d_54_v36_)
                nw9_ = _dafny.Array(None, 16)
                for i0_1_ in range(nw9_.length(0)):
                    nw9_[i0_1_] = init1_(i0_1_)
                d_69_v45_ = nw9_
                index2_ = default__.safeIndex(389, (d_69_v45_).length(0))
                (d_69_v45_)[index2_] = 551
                d_72_v46_: _dafny.Array
                def lambda4_(d_73_p0_):
                    def lambda5_(d_74_i4_):
                        return d_73_p0_

                    return lambda5_

                init2_ = lambda4_(p0)
                nw10_ = _dafny.Array(None, 5)
                for i0_2_ in range(nw10_.length(0)):
                    nw10_[i0_2_] = init2_(i0_2_)
                d_72_v46_ = nw10_
                d_75_v47_: D4
                d_75_v47_ = D4_DC11(default__.fm10(d_33_v16_, len(d_15_v6_), globalState), d_72_v46_)
                d_76_v48_: _dafny.Map
                d_76_v48_ = _dafny.Map({d_75_v47_: p0})
                d_77_v49_: _dafny.Array
                nw11_ = _dafny.Array(None, 9)
                nw11_[int(0)] = d_76_v48_
                nw11_[int(1)] = d_76_v48_
                nw11_[int(2)] = (_dafny.Map({d_75_v47_: True})) | (_dafny.Map({d_75_v47_: p0}))
                nw11_[int(3)] = _dafny.Map({d_75_v47_: p0})
                nw11_[int(4)] = (d_76_v48_) | ((_dafny.Map({d_75_v47_: p0})).set(d_75_v47_, p0))
                nw11_[int(5)] = d_76_v48_
                nw11_[int(6)] = d_76_v48_
                nw11_[int(7)] = (d_76_v48_) | (d_76_v48_)
                nw11_[int(8)] = _dafny.Map({d_75_v47_: p0})
                d_77_v49_ = nw11_
                index3_ = default__.safeIndex(150, (d_77_v49_).length(0))
                (d_77_v49_)[index3_] = d_76_v48_
                index4_ = default__.safeIndex(389, (d_69_v45_).length(0))
                index5_ = default__.safeIndex(150, (d_77_v49_).length(0))
                rhs0_ = d_68_v44_
                rhs1_ = d_64_cf3_
                rhs2_ = not(not((p0 if p0 else p0)))
                rhs3_ = (d_76_v48_) | ((d_76_v48_) | (d_76_v48_))
                rhs4_ = p0
                lhs0_ = d_69_v45_
                lhs1_ = default__.safeIndex(389, (d_69_v45_).length(0))
                lhs2_ = globalState
                lhs3_ = d_77_v49_
                lhs4_ = default__.safeIndex(150, (d_77_v49_).length(0))
                lhs5_ = globalState
                d_68_v44_ = rhs0_
                lhs0_[lhs1_] = rhs1_
                lhs2_.f24 = rhs2_
                lhs3_[lhs4_] = rhs3_
                lhs5_.f3 = rhs4_
                d_78_v50_: _dafny.Set
                d_78_v50_ = _dafny.Set({d_72_v46_, d_72_v46_})
                rhs5_ = (d_78_v50_).intersection((d_78_v50_ if p0 else _dafny.Set({d_72_v46_})))
                rhs6_ = d_33_v16_
                rhs7_ = (d_69_v45_)[default__.safeIndex(389, (d_69_v45_).length(0))]
                lhs6_ = globalState
                d_78_v50_ = rhs5_
                r3 = rhs6_
                lhs6_.f9 = rhs7_
                d_79_v51_: _dafny.Map
                d_79_v51_ = _dafny.Map({p0: d_65_cf2_})
                d_80_v52_: _dafny.Map
                d_80_v52_ = _dafny.Map({d_79_v51_: (not(p0) if p0 else not(p0))})
                d_80_v52_ = (d_80_v52_).set(d_79_v51_, p0)
            elif True:
                d_81___mcc_h7_ = source1_.cf0
                d_82_cf0_ = d_81___mcc_h7_
                d_83_v53_: C0
                nw12_ = C0()
                nw12_.ctor__()
                d_83_v53_ = nw12_
                (globalState).f21 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukla"))
                d_84_v54_: _dafny.Seq
                d_84_v54_ = _dafny.SeqWithoutIsStrInference([default__.fm11(p0, globalState), default__.fm11(p0, globalState), d_15_v6_])
                r1 = (len(default__.fm2((d_83_v53_).fm5(globalState), len(d_15_v6_), p0, d_13_v4_, globalState))) * (len(d_84_v54_))
                (globalState).f9 = d_13_v4_
            d_85_v55_: C0
            nw13_ = C0()
            nw13_.ctor__()
            d_85_v55_ = nw13_
            d_86_v56_: _dafny.Array
            nw14_ = _dafny.Array(None, 1)
            nw14_[int(0)] = d_85_v55_
            d_86_v56_ = nw14_
            index6_ = default__.safeIndex(359, (d_86_v56_).length(0))
            (d_86_v56_)[index6_] = d_85_v55_
            index7_ = default__.safeIndex(359, (d_86_v56_).length(0))
            (d_86_v56_)[index7_] = d_85_v55_
        elif True:
            d_87___mcc_h3_ = source0_.cf27
            d_88_cf27_ = d_87___mcc_h3_
            d_89_v57_: C0
            nw15_ = C0()
            nw15_.ctor__()
            d_89_v57_ = nw15_
            r1 = 788
            d_90_v58_: _dafny.Seq
            d_90_v58_ = _dafny.SeqWithoutIsStrInference([not(p0)])
            d_91_v59_: _dafny.Map
            d_91_v59_ = _dafny.Map({d_90_v58_: p0})
            d_91_v59_ = d_91_v59_
            d_92_v60_: _dafny.Array
            nw16_ = _dafny.Array(_dafny.CodePoint('D'), 8)
            d_92_v60_ = nw16_
            d_93_v61_: str
            d_93_v61_ = _dafny.CodePoint('x')
            index8_ = default__.safeIndex(970, (d_92_v60_).length(0))
            (d_92_v60_)[index8_] = (d_93_v61_ if p0 else d_93_v61_)
            d_94_v62_: _dafny.Map
            d_94_v62_ = _dafny.Map({p0: p0})
            d_95_v63_: D0
            d_95_v63_ = D0_DC1(p0)
            d_96_v64_: _dafny.Seq
            d_96_v64_ = _dafny.SeqWithoutIsStrInference([d_95_v63_, d_95_v63_])
            d_97_v65_: _dafny.Map
            d_97_v65_ = _dafny.Map({d_94_v62_: d_96_v64_})
            d_98_v66_: D5
            d_98_v66_ = D5_DC14((d_97_v65_).set(d_94_v62_, d_96_v64_), d_13_v4_)
            index9_ = default__.safeIndex(970, (d_92_v60_).length(0))
            rhs8_ = (d_98_v66_).cf20
            rhs9_ = not(not(p0))
            rhs10_ = _dafny.CodePoint('c')
            rhs11_ = ((0) - (d_13_v4_)) + (d_13_v4_)
            lhs7_ = globalState
            lhs8_ = d_92_v60_
            lhs9_ = default__.safeIndex(970, (d_92_v60_).length(0))
            lhs10_ = globalState
            r1 = rhs8_
            lhs7_.f24 = rhs9_
            lhs8_[lhs9_] = rhs10_
            lhs10_.f25 = rhs11_
        d_99_v67_: C0
        nw17_ = C0()
        nw17_.ctor__()
        d_99_v67_ = nw17_
        d_100_v68_: _dafny.Map
        d_100_v68_ = _dafny.Map({d_99_v67_: (d_99_v67_).fm5(globalState)})
        d_101_v69_: _dafny.Map
        d_101_v69_ = _dafny.Map({((d_100_v68_)[d_99_v67_] if (d_99_v67_) in (d_100_v68_) else 86): d_13_v4_})
        rhs12_ = d_13_v4_
        rhs13_ = (d_101_v69_) | ((d_101_v69_) | (d_101_v69_))
        d_13_v4_ = rhs12_
        d_101_v69_ = rhs13_
        hi0_ = d_13_v4_
        for d_102_i5_ in range(d_13_v4_, hi0_):
            d_103_v70_: _dafny.Array
            nw18_ = _dafny.Array(None, 2)
            nw18_[int(0)] = p0
            nw18_[int(1)] = default__.fm0(d_13_v4_, True, globalState)
            d_103_v70_ = nw18_
            index10_ = default__.safeIndex(919, (d_103_v70_).length(0))
            (d_103_v70_)[index10_] = False
            d_104_v71_: _dafny.MultiSet
            d_104_v71_ = _dafny.MultiSet([p0])
            index11_ = default__.safeIndex(919, (d_103_v70_).length(0))
            (d_103_v70_)[index11_] = default__.fm0(d_102_i5_, (417) == (((d_104_v71_)[p0] if (p0) in (d_104_v71_) else d_102_i5_)), globalState)
            (globalState).f14 = (d_103_v70_)[default__.safeIndex(919, (d_103_v70_).length(0))]
            d_105_v72_: _dafny.Seq
            d_105_v72_ = _dafny.SeqWithoutIsStrInference([p0, not(p0), p0, p0])
            d_106_v73_: _dafny.Seq
            d_106_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjmbsow"))
            (globalState).f9 = (((d_14_v5_).cf2) - (len(d_105_v72_))) * (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_107_i6_ in range(default__.abs(997))])) + (d_106_v73_)))
            d_108_v74_: D2
            d_108_v74_ = D2_DC4(default__.fm12(d_105_v72_, (d_103_v70_)[default__.safeIndex(919, (d_103_v70_).length(0))], globalState))
            d_108_v74_ = default__.fm13(globalState)
        d_109_v75_: _dafny.Seq
        d_109_v75_ = _dafny.SeqWithoutIsStrInference([p0, not(p0), False, False])
        d_110_v76_: _dafny.Seq
        d_110_v76_ = _dafny.SeqWithoutIsStrInference([d_109_v75_])
        d_111_v77_: _dafny.MultiSet
        d_111_v77_ = _dafny.MultiSet([d_13_v4_, d_13_v4_])
        r0 = (d_110_v76_)[default__.safeIndex(((d_111_v77_)[(0) - (d_13_v4_)] if ((0) - (d_13_v4_)) in (d_111_v77_) else len(default__.fm14(_dafny.SeqWithoutIsStrInference([d_13_v4_ for d_112_i7_ in range(default__.abs(798))]), d_13_v4_, d_13_v4_, d_111_v77_, globalState))), len(d_110_v76_))]
        r1 = d_13_v4_
        d_113_v78_: _dafny.Array
        def lambda6_(d_114_v4_):
            def lambda7_(d_115_i8_):
                return default__.safeModuloInt(d_115_i8_, d_114_v4_)

            return lambda7_

        init3_ = lambda6_(d_13_v4_)
        nw19_ = _dafny.Array(None, 8)
        for i0_3_ in range(nw19_.length(0)):
            nw19_[i0_3_] = init3_(i0_3_)
        d_113_v78_ = nw19_
        d_116_v79_: _dafny.Seq
        d_116_v79_ = _dafny.SeqWithoutIsStrInference([d_113_v78_, d_113_v78_, d_113_v78_, d_113_v78_])
        r2 = (d_116_v79_)[default__.safeIndex(d_13_v4_, len(d_116_v79_))]
        d_117_v80_: str
        d_117_v80_ = _dafny.CodePoint('f')
        r3 = d_117_v80_
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_118_v0_: int
        d_118_v0_ = 691
        d_119_v1_: _dafny.Set
        d_119_v1_ = _dafny.Set({d_118_v0_})
        d_120_v2_: _dafny.Map
        d_120_v2_ = _dafny.Map({d_119_v1_: d_118_v0_})
        d_121_v3_: _dafny.Seq
        d_121_v3_ = _dafny.SeqWithoutIsStrInference([d_118_v0_, d_118_v0_])
        d_122_v4_: _dafny.Map
        d_122_v4_ = _dafny.Map({len(_dafny.Set({d_118_v0_})): d_121_v3_})
        d_123_v5_: _dafny.Array
        nw20_ = _dafny.Array(int(0), 24)
        d_123_v5_ = nw20_
        d_124_v6_: D0
        d_124_v6_ = D0_DC0(d_123_v5_)
        d_125_v7_: _dafny.Array
        nw21_ = _dafny.Array(None, 4)
        nw21_[int(0)] = d_123_v5_
        nw21_[int(1)] = d_123_v5_
        nw21_[int(2)] = d_123_v5_
        nw21_[int(3)] = (d_124_v6_).cf0
        d_125_v7_ = nw21_
        d_126_v8_: _dafny.MultiSet
        d_126_v8_ = _dafny.MultiSet([len(d_121_v3_)])
        d_127_v9_: _dafny.Seq
        d_127_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "px"))
        d_128_v10_: _dafny.Map
        d_128_v10_ = _dafny.Map({d_125_v7_: _dafny.Set({(d_126_v8_).cardinality, d_118_v0_, len(d_127_v9_), d_118_v0_})})
        d_129_v11_: _dafny.Array
        def lambda8_(d_130_i0_):
            return (_dafny.CodePoint('j'))

        init4_ = lambda8_
        nw22_ = _dafny.Array(None, 23)
        for i0_4_ in range(nw22_.length(0)):
            nw22_[i0_4_] = init4_(i0_4_)
        d_129_v11_ = nw22_
        d_131_v12_: bool
        d_131_v12_ = True
        d_132_v13_: _dafny.Map
        d_132_v13_ = _dafny.Map({d_131_v12_: d_131_v12_})
        d_133_v14_: _dafny.Array
        nw23_ = _dafny.Array(None, 2)
        nw23_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqobbdls"))
        nw23_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgyeuu"))
        d_133_v14_ = nw23_
        d_134_v15_: _dafny.Set
        d_134_v15_ = _dafny.Set({d_121_v3_})
        d_135_globalState_: GlobalState
        nw24_ = GlobalState()
        nw24_.ctor__((d_120_v2_) | (d_120_v2_), d_119_v1_, 445, False, 988, _dafny.CodePoint('l'), d_122_v4_, 496, -233, 938, ((d_128_v10_)[d_125_v7_] if (d_125_v7_) in (d_128_v10_) else d_119_v1_), 987, False, True, False, 99, False, d_129_v11_, 161, d_132_v13_, d_133_v14_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txxtamr")), 577, d_134_v15_, False, 181)
        d_135_globalState_ = nw24_
        d_136_v16_: str
        d_136_v16_ = _dafny.CodePoint('j')
        d_137_v17_: _dafny.Map
        d_137_v17_ = _dafny.Map({_dafny.Map({d_131_v12_: d_118_v0_}): d_136_v16_})
        d_137_v17_ = d_137_v17_
        d_138_i1_: int
        d_138_i1_ = 0
        with _dafny.label("1"):
            while d_131_v12_:
                with _dafny.c_label("1"):
                    if (d_138_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_138_i1_ = (d_138_i1_) + (1)
                    d_139_v18_: _dafny.Seq
                    d_139_v18_ = _dafny.SeqWithoutIsStrInference([d_131_v12_])
                    (d_135_globalState_).f14 = (d_139_v18_)[default__.safeIndex(d_118_v0_, len(d_139_v18_))]
                    d_140_v19_: D0
                    d_140_v19_ = D0_DC1(d_131_v12_)
                    d_141_v20_: _dafny.Map
                    d_141_v20_ = _dafny.Map({d_140_v19_: (d_126_v8_).cardinality})
                    d_142_v21_: _dafny.Map
                    d_142_v21_ = _dafny.Map({d_118_v0_: d_131_v12_})
                    d_143_v22_: _dafny.Set
                    d_143_v22_ = _dafny.Set({d_131_v12_})
                    d_144_v23_: _dafny.Map
                    d_144_v23_ = _dafny.Map({((d_142_v21_)[len(d_143_v22_)] if (len(d_143_v22_)) in (d_142_v21_) else d_131_v12_): d_118_v0_})
                    d_141_v20_ = (d_141_v20_).set(d_140_v19_, len(d_144_v23_))
                    index12_ = default__.safeIndex(874, (d_123_v5_).length(0))
                    (d_123_v5_)[index12_] = len(d_127_v9_)
                    d_145_v24_: _dafny.Map
                    d_145_v24_ = _dafny.Map({(d_139_v18_)[default__.safeIndex(len(d_139_v18_), len(d_139_v18_))]: d_136_v16_})
                    d_146_v25_: _dafny.Map
                    d_146_v25_ = _dafny.Map({len(d_145_v24_): d_123_v5_})
                    index13_ = default__.safeIndex(874, (d_123_v5_).length(0))
                    (d_123_v5_)[index13_] = (d_118_v0_) + (len((d_146_v25_) | (d_146_v25_)))
                    d_147_v26_: _dafny.Map
                    d_147_v26_ = _dafny.Map({d_136_v16_: d_131_v12_})
                    d_148_v27_: str
                    d_148_v27_ = d_136_v16_
                    d_147_v26_ = (d_147_v26_).set(d_148_v27_, d_131_v12_)
                    pass
            pass
        d_149_v28_: _dafny.Seq
        d_150_v29_: int
        d_151_v30_: _dafny.Array
        d_152_v31_: str
        out0_: _dafny.Seq
        out1_: int
        out2_: _dafny.Array
        out3_: str
        out0_, out1_, out2_, out3_ = default__.m0(default__.fm0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gb"))), d_131_v12_, d_135_globalState_), d_135_globalState_)
        d_149_v28_ = out0_
        d_150_v29_ = out1_
        d_151_v30_ = out2_
        d_152_v31_ = out3_
        (d_135_globalState_).f22 = (0) - (d_150_v29_)
        if d_131_v12_:
            (d_135_globalState_).f14 = d_131_v12_
            d_153_v32_: D0
            d_153_v32_ = D0_DC1(d_131_v12_)
            source2_ = d_153_v32_
            if source2_.is_DC1:
                d_154___mcc_h0_ = source2_.cf1
                d_155_cf1_ = d_154___mcc_h0_
                d_156_v33_: _dafny.Array
                nw25_ = _dafny.Array(None, 8)
                nw25_[int(0)] = (d_150_v29_) != (d_150_v29_)
                nw25_[int(1)] = True
                nw25_[int(2)] = d_131_v12_
                nw25_[int(3)] = d_131_v12_
                nw25_[int(4)] = (d_119_v1_).isdisjoint(d_119_v1_)
                nw25_[int(5)] = d_155_cf1_
                nw25_[int(6)] = (492) == ((0) - ((0) - (d_118_v0_)))
                nw25_[int(7)] = d_155_cf1_
                d_156_v33_ = nw25_
                d_156_v33_ = d_156_v33_
                (d_135_globalState_).f14 = (not(d_131_v12_)) or (d_155_cf1_)
                (d_135_globalState_).f22 = d_150_v29_
                index14_ = default__.safeIndex(176, (d_123_v5_).length(0))
                (d_123_v5_)[index14_] = d_150_v29_
                index15_ = default__.safeIndex(176, (d_123_v5_).length(0))
                (d_123_v5_)[index15_] = 882
            elif source2_.is_DC2:
                d_157___mcc_h1_ = source2_.cf2
                d_158___mcc_h2_ = source2_.cf3
                d_159_cf3_ = d_158___mcc_h2_
                d_160_cf2_ = d_157___mcc_h1_
                d_161_v34_: _dafny.MultiSet
                d_161_v34_ = _dafny.MultiSet([d_131_v12_, not(d_131_v12_)])
                d_162_v36_: _dafny.Map
                d_162_v36_ = _dafny.Map({d_160_cf2_: d_118_v0_})
                d_163_v37_: _dafny.Map
                d_163_v37_ = _dafny.Map({d_131_v12_: default__.fm1(716, (d_149_v28_)[default__.safeIndex(d_160_cf2_, len(d_149_v28_))], d_127_v9_, d_131_v12_, d_135_globalState_)})
                d_164_v40_: _dafny.Array
                nw26_ = _dafny.Array(None, 19)
                def iife3_():
                    coll3_ = _dafny.Map()
                    compr_3_: int
                    for compr_3_ in _dafny.IntegerRange(717, 364):
                        d_165_v35_: int = compr_3_
                        if ((717) <= (d_165_v35_)) and ((d_165_v35_) < (364)):
                            coll3_[(d_165_v35_) * (len(_dafny.Map({_dafny.Map({d_131_v12_: d_160_cf2_}): True})))] = (_dafny.MultiSet([d_160_cf2_, (0) - (d_118_v0_)])).cardinality
                    return _dafny.Map(coll3_)
                nw26_[int(0)] = (_dafny.Map({(d_161_v34_).cardinality: d_159_cf3_})) | (iife3_()
                )
                nw26_[int(1)] = d_162_v36_
                nw26_[int(2)] = d_162_v36_
                nw26_[int(3)] = d_162_v36_
                nw26_[int(4)] = (_dafny.Map({d_160_cf2_: d_150_v29_})) | (d_162_v36_)
                nw26_[int(5)] = d_162_v36_
                nw26_[int(6)] = d_162_v36_
                nw26_[int(7)] = (d_162_v36_).set(len(d_163_v37_), d_159_cf3_)
                nw26_[int(8)] = d_162_v36_
                nw26_[int(9)] = (d_162_v36_ if d_131_v12_ else (D2_DC4(d_162_v36_)).cf5)
                nw26_[int(10)] = d_162_v36_
                def iife4_():
                    coll4_ = _dafny.Map()
                    compr_4_: int
                    for compr_4_ in _dafny.IntegerRange(-473, 31):
                        d_166_v38_: int = compr_4_
                        if ((-473) <= (d_166_v38_)) and ((d_166_v38_) < (31)):
                            coll4_[(d_166_v38_) + ((0) - (d_150_v29_))] = (0) - (d_160_cf2_)
                    return _dafny.Map(coll4_)
                nw26_[int(11)] = iife4_()
                
                nw26_[int(12)] = (_dafny.Map({d_150_v29_: d_160_cf2_})) | (d_162_v36_)
                nw26_[int(13)] = (d_162_v36_) | (d_162_v36_)
                nw26_[int(14)] = d_162_v36_
                nw26_[int(15)] = d_162_v36_
                nw26_[int(16)] = d_162_v36_
                nw26_[int(17)] = d_162_v36_
                def iife5_():
                    coll5_ = _dafny.Map()
                    compr_5_: int
                    for compr_5_ in ((d_162_v36_).set(len(d_149_v28_), d_159_cf3_)).keys.Elements:
                        d_167_v39_: int = compr_5_
                        if (d_167_v39_) in ((d_162_v36_).set(len(d_149_v28_), d_159_cf3_)):
                            coll5_[default__.safeModuloInt(d_167_v39_, 990)] = len(d_149_v28_)
                    return _dafny.Map(coll5_)
                nw26_[int(18)] = (iife5_()
                ).set(d_160_cf2_, (0) - (-994))
                d_164_v40_ = nw26_
                d_168_v41_: _dafny.Map
                d_168_v41_ = _dafny.Map({d_150_v29_: d_153_v32_})
                index16_ = default__.safeIndex(329, (d_151_v30_).length(0))
                (d_151_v30_)[index16_] = default__.fm1(len(d_168_v41_), d_131_v12_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjvcgp")), d_131_v12_, d_135_globalState_)
                index17_ = default__.safeIndex(329, (d_151_v30_).length(0))
                rhs14_ = d_160_cf2_
                rhs15_ = d_160_cf2_
                rhs16_ = d_164_v40_
                rhs17_ = not(d_131_v12_)
                rhs18_ = d_118_v0_
                lhs11_ = d_135_globalState_
                lhs12_ = d_135_globalState_
                lhs13_ = d_151_v30_
                lhs14_ = default__.safeIndex(329, (d_151_v30_).length(0))
                d_150_v29_ = rhs14_
                lhs11_.f22 = rhs15_
                d_164_v40_ = rhs16_
                lhs12_.f14 = rhs17_
                lhs13_[lhs14_] = rhs18_
                d_169_v42_: _dafny.Set
                d_169_v42_ = _dafny.Set({(d_162_v36_) | ((d_162_v36_).set(12, (d_151_v30_)[default__.safeIndex(329, (d_151_v30_).length(0))]))})
                d_170_v43_: _dafny.Map
                d_170_v43_ = _dafny.Map({d_162_v36_: d_131_v12_})
                pat_let_tv0_ = d_169_v42_
                pat_let_tv1_ = d_169_v42_
                def iife7_(_pat_let1_0):
                    def iife8_(d_171_dt__update__tmp_h0_):
                        def iife9_(_pat_let2_0):
                            def iife10_(d_172_dt__update_hcf12_h0_):
                                return D3_DC8(d_172_dt__update_hcf12_h0_)
                            return iife10_(_pat_let2_0)
                        return iife9_(pat_let_tv0_)
                    return iife8_(_pat_let1_0)
                def iife6_(_pat_let0_0):
                    def iife11_(d_173_dt__update__tmp_h1_):
                        def iife12_(_pat_let3_0):
                            def iife13_(d_174_dt__update_hcf12_h1_):
                                return D3_DC8(d_174_dt__update_hcf12_h1_)
                            return iife13_(_pat_let3_0)
                        return iife12_(pat_let_tv1_)
                    return iife11_(_pat_let0_0)
                def iife14_():
                    coll6_ = _dafny.Set()
                    compr_6_: _dafny.Map
                    for compr_6_ in (d_170_v43_).keys.Elements:
                        d_175_v44_: _dafny.Map = compr_6_
                        if (d_175_v44_) in (d_170_v43_):
                            coll6_ = coll6_.union(_dafny.Set([d_175_v44_]))
                    return _dafny.Set(coll6_)
                d_169_v42_ = ((iife6_(iife7_(D3_DC8(_dafny.Set({d_162_v36_}))))).cf12).intersection((d_169_v42_).intersection(iife14_()
                ))
                (d_135_globalState_).f22 = default__.fm1(default__.safeModuloInt((_dafny.MultiSet([True])).cardinality, (0) - ((0) - (d_150_v29_))), d_131_v12_, default__.fm2(d_150_v29_, d_118_v0_, (d_149_v28_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_150_v29_ for d_176_i2_ in range(default__.abs(933))])), len(d_149_v28_))], (0) - (d_159_cf3_), d_135_globalState_), d_131_v12_, d_135_globalState_)
                d_177_v45_: _dafny.Set
                d_177_v45_ = _dafny.Set({True, d_131_v12_})
                d_177_v45_ = ((d_177_v45_) - (default__.fm3((d_151_v30_)[default__.safeIndex(329, (d_151_v30_).length(0))], d_131_v12_, d_135_globalState_))) | ((_dafny.Set({d_131_v12_, d_131_v12_})) - (d_177_v45_))
            elif True:
                d_178___mcc_h3_ = source2_.cf0
                d_179_cf0_ = d_178___mcc_h3_
                (d_135_globalState_).f25 = len(d_119_v1_)
                (d_135_globalState_).f9 = d_118_v0_
                d_180_v46_: _dafny.Seq
                d_181_v47_: int
                d_182_v48_: _dafny.Array
                d_183_v49_: str
                out4_: _dafny.Seq
                out5_: int
                out6_: _dafny.Array
                out7_: str
                out4_, out5_, out6_, out7_ = default__.m0(not(d_131_v12_), d_135_globalState_)
                d_180_v46_ = out4_
                d_181_v47_ = out5_
                d_182_v48_ = out6_
                d_183_v49_ = out7_
                d_184_v50_: _dafny.MultiSet
                d_184_v50_ = _dafny.MultiSet([d_131_v12_, d_131_v12_, d_131_v12_, d_131_v12_, d_131_v12_])
                d_185_v51_: str
                d_185_v51_ = d_136_v16_
                rhs19_ = ((d_184_v50_).cardinality) - ((0) - ((0) - (default__.safeDivisionInt(d_150_v29_, len(d_121_v3_)))))
                rhs20_ = default__.fm4(d_185_v51_, _dafny.SeqWithoutIsStrInference([True]), d_135_globalState_)
                d_181_v47_ = rhs19_
                d_183_v49_ = rhs20_
            d_186_v52_: _dafny.Map
            d_186_v52_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_118_v0_]): d_118_v0_})
            d_187_v54_: _dafny.Seq
            d_187_v54_ = _dafny.SeqWithoutIsStrInference([d_121_v3_, (d_121_v3_).set(default__.safeIndex(d_118_v0_, len(d_121_v3_)), d_118_v0_)])
            def iife15_():
                coll7_ = _dafny.Map()
                compr_7_: _dafny.Seq
                for compr_7_ in (d_187_v54_).Elements:
                    d_188_v53_: _dafny.Seq = compr_7_
                    if (d_188_v53_) in (d_187_v54_):
                        coll7_[d_188_v53_] = d_118_v0_
                return _dafny.Map(coll7_)
            d_186_v52_ = (d_186_v52_) | ((iife15_()
            ) | (d_186_v52_))
            d_189_v55_: _dafny.MultiSet
            d_189_v55_ = _dafny.MultiSet([d_131_v12_, d_131_v12_, True, d_131_v12_])
            d_190_v56_: D0
            d_190_v56_ = D0_DC2(623, default__.fm1(d_150_v29_, d_131_v12_, d_127_v9_, d_131_v12_, d_135_globalState_))
            source3_ = D2_DC5((d_119_v1_).ispropersubset(d_119_v1_), (((d_189_v55_)[False] if (False) in (d_189_v55_) else (d_190_v56_).cf2)) < (d_150_v29_), True, d_150_v29_, (default__.fm1(d_150_v29_, False, default__.fm2(d_150_v29_, d_150_v29_, d_131_v12_, d_118_v0_, d_135_globalState_), d_131_v12_, d_135_globalState_)) < (d_118_v0_))
            if source3_.is_DC5:
                d_191___mcc_h4_ = source3_.cf6
                d_192___mcc_h5_ = source3_.cf7
                d_193___mcc_h6_ = source3_.cf8
                d_194___mcc_h7_ = source3_.cf9
                d_195___mcc_h8_ = source3_.cf10
                d_196_cf10_ = d_195___mcc_h8_
                d_197_cf9_ = d_194___mcc_h7_
                d_198_cf8_ = d_193___mcc_h6_
                d_199_cf7_ = d_192___mcc_h5_
                d_200_cf6_ = d_191___mcc_h4_
                d_201_v57_: _dafny.Set
                d_201_v57_ = _dafny.Set({d_119_v1_})
                d_199_cf7_ = ((d_201_v57_).intersection(d_201_v57_)).issubset(d_201_v57_)
                rhs21_ = (0) - ((0) - ((d_197_cf9_) * (d_150_v29_)))
                rhs22_ = (d_123_v5_ if d_131_v12_ else (d_123_v5_ if d_198_cf8_ else d_123_v5_))
                rhs23_ = d_127_v9_
                rhs24_ = (default__.fm2(d_197_cf9_, d_118_v0_, d_199_cf7_, len(d_149_v28_), d_135_globalState_)) + (d_127_v9_)
                lhs15_ = d_135_globalState_
                lhs16_ = d_135_globalState_
                d_150_v29_ = rhs21_
                d_123_v5_ = rhs22_
                lhs15_.f21 = rhs23_
                lhs16_.f21 = rhs24_
                index18_ = default__.safeIndex(193, (d_123_v5_).length(0))
                (d_123_v5_)[index18_] = d_118_v0_
                index19_ = default__.safeIndex(193, (d_123_v5_).length(0))
                (d_123_v5_)[index19_] = d_150_v29_
                d_202_v58_: _dafny.Seq
                d_203_v59_: int
                d_204_v60_: _dafny.Array
                d_205_v61_: str
                out8_: _dafny.Seq
                out9_: int
                out10_: _dafny.Array
                out11_: str
                out8_, out9_, out10_, out11_ = default__.m0(not(d_131_v12_), d_135_globalState_)
                d_202_v58_ = out8_
                d_203_v59_ = out9_
                d_204_v60_ = out10_
                d_205_v61_ = out11_
            elif source3_.is_DC6:
                d_206_v62_: C0
                nw27_ = C0()
                nw27_.ctor__()
                d_206_v62_ = nw27_
                d_207_v63_: _dafny.MultiSet
                d_207_v63_ = _dafny.MultiSet([d_129_v11_])
                d_208_v64_: _dafny.Map
                d_208_v64_ = _dafny.Map({d_207_v63_: d_206_v62_})
                d_209_v65_: _dafny.Seq
                d_209_v65_ = _dafny.SeqWithoutIsStrInference([d_206_v62_])
                d_210_v66_: _dafny.Map
                d_210_v66_ = _dafny.Map({d_131_v12_: d_206_v62_})
                d_211_v67_: _dafny.Array
                nw28_ = _dafny.Array(None, 26)
                nw28_[int(0)] = d_206_v62_
                nw28_[int(1)] = d_206_v62_
                nw28_[int(2)] = d_206_v62_
                nw28_[int(3)] = d_206_v62_
                nw28_[int(4)] = d_206_v62_
                nw28_[int(5)] = d_206_v62_
                nw28_[int(6)] = d_206_v62_
                nw28_[int(7)] = d_206_v62_
                nw28_[int(8)] = (d_206_v62_ if default__.fm0(d_118_v0_, d_131_v12_, d_135_globalState_) else d_206_v62_)
                nw28_[int(9)] = ((d_208_v64_)[d_207_v63_] if (d_207_v63_) in (d_208_v64_) else (d_209_v65_)[default__.safeIndex(d_118_v0_, len(d_209_v65_))])
                nw28_[int(10)] = d_206_v62_
                nw28_[int(11)] = d_206_v62_
                nw28_[int(12)] = d_206_v62_
                nw28_[int(13)] = d_206_v62_
                nw28_[int(14)] = d_206_v62_
                nw28_[int(15)] = d_206_v62_
                nw28_[int(16)] = d_206_v62_
                nw28_[int(17)] = d_206_v62_
                nw28_[int(18)] = d_206_v62_
                nw28_[int(19)] = d_206_v62_
                nw28_[int(20)] = ((d_210_v66_)[default__.fm0(d_118_v0_, d_131_v12_, d_135_globalState_)] if (default__.fm0(d_118_v0_, d_131_v12_, d_135_globalState_)) in (d_210_v66_) else d_206_v62_)
                nw28_[int(21)] = ((d_210_v66_)[d_131_v12_] if (d_131_v12_) in (d_210_v66_) else d_206_v62_)
                nw28_[int(22)] = d_206_v62_
                nw28_[int(23)] = d_206_v62_
                nw28_[int(24)] = d_206_v62_
                nw28_[int(25)] = d_206_v62_
                d_211_v67_ = nw28_
                index20_ = default__.safeIndex(521, (d_211_v67_).length(0))
                (d_211_v67_)[index20_] = d_206_v62_
                index21_ = default__.safeIndex(521, (d_211_v67_).length(0))
                (d_211_v67_)[index21_] = (d_206_v62_ if d_131_v12_ else d_206_v62_)
                d_212_v68_: C0
                nw29_ = C0()
                nw29_.ctor__()
                d_212_v68_ = nw29_
                d_213_v69_: _dafny.Seq
                d_214_v70_: int
                d_215_v71_: _dafny.Array
                d_216_v72_: str
                out12_: _dafny.Seq
                out13_: int
                out14_: _dafny.Array
                out15_: str
                out12_, out13_, out14_, out15_ = default__.m0(d_131_v12_, d_135_globalState_)
                d_213_v69_ = out12_
                d_214_v70_ = out13_
                d_215_v71_ = out14_
                d_216_v72_ = out15_
            elif source3_.is_DC4:
                d_217___mcc_h9_ = source3_.cf5
                d_218_cf5_ = d_217___mcc_h9_
                d_219_v73_: C0
                nw30_ = C0()
                nw30_.ctor__()
                d_219_v73_ = nw30_
                d_220_v74_: _dafny.Map
                d_220_v74_ = _dafny.Map({d_219_v73_: d_219_v73_})
                d_221_v75_: D4
                d_221_v75_ = D4_DC10(d_219_v73_)
                rhs25_ = ((d_220_v74_)[(d_221_v75_).cf15] if ((d_221_v75_).cf15) in (d_220_v74_) else d_219_v73_)
                rhs26_ = ((d_189_v55_)[((d_132_v13_)[d_131_v12_] if (d_131_v12_) in (d_132_v13_) else d_131_v12_)] if (((d_132_v13_)[d_131_v12_] if (d_131_v12_) in (d_132_v13_) else d_131_v12_)) in (d_189_v55_) else default__.safeDivisionInt(d_118_v0_, d_150_v29_))
                lhs17_ = d_135_globalState_
                d_219_v73_ = rhs25_
                lhs17_.f22 = rhs26_
                d_222_v76_: _dafny.Array
                nw31_ = _dafny.Array(_dafny.Set({}), 21)
                d_222_v76_ = nw31_
                rhs27_ = d_150_v29_
                rhs28_ = ((d_121_v3_)[default__.safeIndex(d_118_v0_, len(d_121_v3_))]) >= (-295)
                rhs29_ = d_152_v31_
                rhs30_ = not(d_131_v12_)
                rhs31_ = d_222_v76_
                lhs18_ = d_135_globalState_
                lhs19_ = d_135_globalState_
                lhs20_ = d_135_globalState_
                lhs18_.f25 = rhs27_
                lhs19_.f14 = rhs28_
                d_152_v31_ = rhs29_
                lhs20_.f14 = rhs30_
                d_222_v76_ = rhs31_
                (d_135_globalState_).f21 = d_127_v9_
                d_223_v77_: _dafny.Seq
                d_224_v78_: int
                d_225_v79_: _dafny.Array
                d_226_v80_: str
                out16_: _dafny.Seq
                out17_: int
                out18_: _dafny.Array
                out19_: str
                out16_, out17_, out18_, out19_ = default__.m0((845) > (d_150_v29_), d_135_globalState_)
                d_223_v77_ = out16_
                d_224_v78_ = out17_
                d_225_v79_ = out18_
                d_226_v80_ = out19_
            elif True:
                d_227___mcc_h10_ = source3_.cf11
                d_228_cf11_ = d_227___mcc_h10_
                d_229_v81_: _dafny.Array
                nw32_ = _dafny.Array(False, 11)
                d_229_v81_ = nw32_
                index22_ = default__.safeIndex(634, (d_229_v81_).length(0))
                (d_229_v81_)[index22_] = (d_153_v32_).cf1
                d_230_v82_: _dafny.Seq
                d_230_v82_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(default__.fm6(d_118_v0_, d_135_globalState_))])
                index23_ = default__.safeIndex(634, (d_229_v81_).length(0))
                (d_229_v81_)[index23_] = (((d_230_v82_)[default__.safeIndex(d_150_v29_, len(d_230_v82_))]).intersection((d_189_v55_).set(d_131_v12_, default__.abs((0) - (d_118_v0_))))).issubset((d_189_v55_ if d_131_v12_ else _dafny.MultiSet([d_131_v12_, d_131_v12_])))
                (d_135_globalState_).f24 = ((_dafny.MultiSet([d_150_v29_])).ispropersubset(d_126_v8_) if d_131_v12_ else d_131_v12_)
                d_231_v83_: _dafny.Map
                d_231_v83_ = _dafny.Map({len(d_121_v3_): d_150_v29_})
                d_232_v84_: _dafny.Set
                d_232_v84_ = _dafny.Set({d_131_v12_})
                d_233_v85_: _dafny.Set
                d_233_v85_ = _dafny.Set({d_231_v83_, _dafny.Map({len(d_127_v9_): d_150_v29_}), _dafny.Map({len(d_232_v84_): d_150_v29_}), d_231_v83_, d_231_v83_})
                d_234_v86_: D3
                d_234_v86_ = D3_DC8(d_233_v85_)
                d_234_v86_ = d_234_v86_
                (d_135_globalState_).f14 = (d_229_v81_)[default__.safeIndex(634, (d_229_v81_).length(0))]
            d_235_v87_: _dafny.Seq
            d_236_v88_: int
            d_237_v89_: _dafny.Array
            d_238_v90_: str
            out20_: _dafny.Seq
            out21_: int
            out22_: _dafny.Array
            out23_: str
            out20_, out21_, out22_, out23_ = default__.m0(d_131_v12_, d_135_globalState_)
            d_235_v87_ = out20_
            d_236_v88_ = out21_
            d_237_v89_ = out22_
            d_238_v90_ = out23_
        elif True:
            source4_ = d_152_v31_
            d_239___mcc_h11_ = source4_
            d_240_cf4_ = d_239___mcc_h11_
            d_241_v91_: _dafny.Set
            d_241_v91_ = _dafny.Set({d_131_v12_})
            d_241_v91_ = d_241_v91_
            d_242_v92_: _dafny.Array
            nw33_ = _dafny.Array(False, 27)
            d_242_v92_ = nw33_
            index24_ = default__.safeIndex(263, (d_242_v92_).length(0))
            (d_242_v92_)[index24_] = d_131_v12_
            d_243_v93_: _dafny.Array
            nw34_ = _dafny.Array(_dafny.Seq({}), 2)
            d_243_v93_ = nw34_
            d_244_v94_: _dafny.Seq
            d_244_v94_ = _dafny.SeqWithoutIsStrInference([d_123_v5_, d_151_v30_])
            index25_ = default__.safeIndex(48, (d_243_v93_).length(0))
            (d_243_v93_)[index25_] = d_244_v94_
            d_245_v95_: _dafny.Map
            d_245_v95_ = _dafny.Map({not(d_131_v12_): len(d_149_v28_)})
            d_246_v96_: _dafny.Map
            d_246_v96_ = _dafny.Map({_dafny.Map({len(d_245_v95_): d_131_v12_}): d_151_v30_})
            d_247_v97_: _dafny.Map
            d_247_v97_ = _dafny.Map({d_150_v29_: d_131_v12_})
            index26_ = default__.safeIndex(263, (d_242_v92_).length(0))
            index27_ = default__.safeIndex(48, (d_243_v93_).length(0))
            rhs32_ = ((d_246_v96_)[d_247_v97_] if (d_247_v97_) in (d_246_v96_) else (d_123_v5_ if d_131_v12_ else d_151_v30_))
            rhs33_ = d_131_v12_
            rhs34_ = not(d_131_v12_)
            rhs35_ = (d_121_v3_)[default__.safeIndex(d_150_v29_, len(d_121_v3_))]
            rhs36_ = (d_244_v94_) + (d_244_v94_)
            lhs21_ = d_135_globalState_
            lhs22_ = d_242_v92_
            lhs23_ = default__.safeIndex(263, (d_242_v92_).length(0))
            lhs24_ = d_135_globalState_
            lhs25_ = d_243_v93_
            lhs26_ = default__.safeIndex(48, (d_243_v93_).length(0))
            d_151_v30_ = rhs32_
            lhs21_.f14 = rhs33_
            lhs22_[lhs23_] = rhs34_
            lhs24_.f22 = rhs35_
            lhs25_[lhs26_] = rhs36_
            d_248_v98_: C0
            nw35_ = C0()
            nw35_.ctor__()
            d_248_v98_ = nw35_
            d_249_v99_: _dafny.Seq
            d_250_v100_: int
            d_251_v101_: _dafny.Array
            d_252_v102_: str
            out24_: _dafny.Seq
            out25_: int
            out26_: _dafny.Array
            out27_: str
            out24_, out25_, out26_, out27_ = default__.m0(default__.fm0(d_118_v0_, (d_242_v92_)[default__.safeIndex(263, (d_242_v92_).length(0))], d_135_globalState_), d_135_globalState_)
            d_249_v99_ = out24_
            d_250_v100_ = out25_
            d_251_v101_ = out26_
            d_252_v102_ = out27_
            d_127_v9_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aml"))) + (d_127_v9_)) + ((d_127_v9_) + (d_127_v9_))
            (d_135_globalState_).f14 = (d_131_v12_ if (d_150_v29_) != (d_118_v0_) else (d_131_v12_ if d_131_v12_ else d_131_v12_))
            (d_135_globalState_).f24 = d_131_v12_
            d_253_v103_: _dafny.Map
            d_253_v103_ = _dafny.Map({d_150_v29_: d_127_v9_})
            d_253_v103_ = (d_253_v103_).set(len(_dafny.Map({d_118_v0_: not(False)})), d_127_v9_)
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_123_v5_).length(0)):
            d_254_i3_: int = guard_loop_0_
            if (True) and (((0) <= (d_254_i3_)) and ((d_254_i3_) < ((d_123_v5_).length(0)))):
                (d_123_v5_)[(d_254_i3_)] = (d_254_i3_) + (110)
        d_255_v104_: _dafny.Map
        d_255_v104_ = _dafny.Map({(0) - (d_150_v29_): d_152_v31_})
        (d_135_globalState_).f24 = (((d_255_v104_)[d_118_v0_] if (d_118_v0_) in (d_255_v104_) else _dafny.CodePoint('g'))) in ((d_127_v9_).set(default__.safeIndex((0) - (d_118_v0_), len(d_127_v9_)), d_152_v31_))
        d_256_i4_: int
        d_256_i4_ = 0
        with _dafny.label("2"):
            while d_131_v12_:
                with _dafny.c_label("2"):
                    if (d_256_i4_) >= (100):
                        raise _dafny.Break("2")
                    d_256_i4_ = (d_256_i4_) + (1)
                    d_257_v105_: _dafny.Seq
                    d_257_v105_ = _dafny.SeqWithoutIsStrInference([d_123_v5_, d_123_v5_, d_123_v5_])
                    index28_ = default__.safeIndex(236, (d_151_v30_).length(0))
                    (d_151_v30_)[index28_] = (len(_dafny.SeqWithoutIsStrInference([(d_149_v28_)[default__.safeIndex(d_150_v29_, len(d_149_v28_))]]))) - (len(d_257_v105_))
                    d_258_v106_: _dafny.Seq
                    d_258_v106_ = _dafny.SeqWithoutIsStrInference([d_132_v13_])
                    d_259_v107_: _dafny.Map
                    d_259_v107_ = _dafny.Map({len(d_258_v106_): d_131_v12_})
                    d_260_v108_: _dafny.Map
                    d_260_v108_ = _dafny.Map({d_259_v107_: d_118_v0_})
                    d_261_v111_: _dafny.Seq
                    def iife16_():
                        coll8_ = _dafny.Map()
                        compr_8_: int
                        for compr_8_ in (_dafny.MultiSet([d_150_v29_, d_118_v0_])).Elements:
                            d_262_v110_: int = compr_8_
                            if (d_262_v110_) in (_dafny.MultiSet([d_150_v29_, d_118_v0_])):
                                coll8_[(d_262_v110_) + (d_118_v0_)] = d_131_v12_
                        return _dafny.Map(coll8_)
                    d_261_v111_ = _dafny.SeqWithoutIsStrInference([d_259_v107_, d_259_v107_, iife16_()
                    ])
                    index29_ = default__.safeIndex(236, (d_151_v30_).length(0))
                    def iife17_():
                        coll9_ = _dafny.Map()
                        compr_9_: _dafny.Map
                        for compr_9_ in (d_261_v111_).Elements:
                            d_263_v109_: _dafny.Map = compr_9_
                            if (d_263_v109_) in (d_261_v111_):
                                coll9_[d_263_v109_] = d_150_v29_
                        return _dafny.Map(coll9_)
                    rhs37_ = len(((d_260_v108_) | (iife17_()
                    )) | (d_260_v108_))
                    rhs38_ = d_150_v29_
                    lhs27_ = d_135_globalState_
                    lhs28_ = d_151_v30_
                    lhs29_ = default__.safeIndex(236, (d_151_v30_).length(0))
                    lhs27_.f9 = rhs37_
                    lhs28_[lhs29_] = rhs38_
                    d_264_v112_: _dafny.Set
                    d_264_v112_ = _dafny.Set({d_131_v12_, d_131_v12_})
                    d_265_v113_: _dafny.Map
                    d_265_v113_ = _dafny.Map({_dafny.Map({d_150_v29_: d_136_v16_}): True})
                    (d_135_globalState_).f9 = len(_dafny.Map({default__.safeModuloInt((0) - (default__.fm1(len(d_264_v112_), d_131_v12_, d_127_v9_, d_131_v12_, d_135_globalState_)), len(d_265_v113_)): (d_151_v30_)[default__.safeIndex(236, (d_151_v30_).length(0))]}))
                    index30_ = default__.safeIndex(236, (d_151_v30_).length(0))
                    (d_151_v30_)[index30_] = default__.safeModuloInt(d_150_v29_, len(d_259_v107_))
                    (d_135_globalState_).f22 = (0) - ((d_151_v30_)[default__.safeIndex(236, (d_151_v30_).length(0))])
                    pass
            pass
        (d_135_globalState_).f25 = d_150_v29_
        hi1_ = 63
        for d_266_i5_ in range(len(d_127_v9_), hi1_):
            d_267_v114_: _dafny.Map
            d_267_v114_ = _dafny.Map({d_266_i5_: d_150_v29_})
            d_268_v115_: _dafny.Set
            d_268_v115_ = _dafny.Set({d_267_v114_})
            d_269_v116_: D3
            d_269_v116_ = D3_DC8(d_268_v115_)
            pat_let_tv2_ = d_267_v114_
            pat_let_tv3_ = d_267_v114_
            pat_let_tv4_ = d_268_v115_
            def iife18_(_pat_let4_0):
                def iife19_(d_270_dt__update__tmp_h2_):
                    def iife20_(_pat_let5_0):
                        def iife21_(d_271_dt__update_hcf12_h2_):
                            return D3_DC8(d_271_dt__update_hcf12_h2_)
                        return iife21_(_pat_let5_0)
                    return iife20_((_dafny.Set({pat_let_tv2_, pat_let_tv3_})) | (pat_let_tv4_))
                return iife19_(_pat_let4_0)
            source5_ = iife18_(d_269_v116_)
            if source5_.is_DC9:
                d_272___mcc_h12_ = source5_.cf13
                d_273___mcc_h13_ = source5_.cf14
                d_274_cf14_ = d_273___mcc_h13_
                d_275_cf13_ = d_272___mcc_h12_
                d_276_v117_: _dafny.Map
                d_276_v117_ = _dafny.Map({(d_149_v28_) + (d_149_v28_): d_150_v29_})
                d_276_v117_ = (d_276_v117_).set(d_149_v28_, (0) - (d_266_i5_))
                d_277_v118_: _dafny.Array
                nw36_ = _dafny.Array(None, 16)
                nw36_[int(0)] = d_275_cf13_
                nw36_[int(1)] = d_131_v12_
                nw36_[int(2)] = d_131_v12_
                nw36_[int(3)] = d_275_cf13_
                nw36_[int(4)] = d_275_cf13_
                nw36_[int(5)] = d_275_cf13_
                nw36_[int(6)] = d_131_v12_
                nw36_[int(7)] = d_275_cf13_
                nw36_[int(8)] = d_275_cf13_
                nw36_[int(9)] = d_131_v12_
                nw36_[int(10)] = d_131_v12_
                nw36_[int(11)] = d_131_v12_
                nw36_[int(12)] = d_131_v12_
                nw36_[int(13)] = d_131_v12_
                nw36_[int(14)] = d_275_cf13_
                nw36_[int(15)] = d_275_cf13_
                d_277_v118_ = nw36_
                d_278_v119_: _dafny.Map
                d_278_v119_ = _dafny.Map({d_266_i5_: d_277_v118_})
                (d_135_globalState_).f3 = ((d_278_v119_ if not(d_275_cf13_) else d_278_v119_)) != (d_278_v119_)
                index31_ = default__.safeIndex(920, (d_151_v30_).length(0))
                (d_151_v30_)[index31_] = default__.safeModuloInt(default__.fm1(d_150_v29_, d_131_v12_, d_127_v9_, d_275_cf13_, d_135_globalState_), (0) - (d_150_v29_))
                index32_ = default__.safeIndex(920, (d_151_v30_).length(0))
                (d_151_v30_)[index32_] = d_118_v0_
                d_149_v28_ = default__.fm6(len(d_127_v9_), d_135_globalState_)
            elif True:
                d_279___mcc_h14_ = source5_.cf12
                d_280_cf12_ = d_279___mcc_h14_
                d_281_v120_: _dafny.Map
                d_281_v120_ = _dafny.Map({d_131_v12_: 875})
                d_281_v120_ = (d_281_v120_).set(True, (len(_dafny.SeqWithoutIsStrInference([d_118_v0_])) if d_131_v12_ else d_150_v29_))
                (d_135_globalState_).f3 = d_131_v12_
                (d_135_globalState_).f9 = d_150_v29_
                d_282_v121_: C0
                nw37_ = C0()
                nw37_.ctor__()
                d_282_v121_ = nw37_
            if False:
                d_283_v122_: _dafny.Seq
                d_284_v123_: int
                d_285_v124_: _dafny.Array
                d_286_v125_: str
                out28_: _dafny.Seq
                out29_: int
                out30_: _dafny.Array
                out31_: str
                out28_, out29_, out30_, out31_ = default__.m0(not(d_131_v12_), d_135_globalState_)
                d_283_v122_ = out28_
                d_284_v123_ = out29_
                d_285_v124_ = out30_
                d_286_v125_ = out31_
                d_287_v126_: D2
                d_287_v126_ = D2_DC4(d_267_v114_)
                d_288_v127_: D3
                d_288_v127_ = D3_DC9(d_131_v12_, (d_287_v126_).cf5)
                d_289_v128_: _dafny.Seq
                d_289_v128_ = _dafny.SeqWithoutIsStrInference([(d_267_v114_).set(d_118_v0_, len(d_127_v9_))])
                d_290_v129_: _dafny.Map
                d_290_v129_ = _dafny.Map({d_118_v0_: D3_DC9((d_288_v127_).cf13, (d_289_v128_)[default__.safeIndex(default__.fm1(d_118_v0_, d_131_v12_, _dafny.SeqWithoutIsStrInference([d_286_v125_ for d_291_i6_ in range(default__.abs(410))]), d_131_v12_, d_135_globalState_), len(d_289_v128_))])})
                d_292_v130_: _dafny.Map
                d_292_v130_ = _dafny.Map({(d_290_v129_) | (d_290_v129_): d_152_v31_})
                d_292_v130_ = (d_292_v130_).set(d_290_v129_, d_136_v16_)
                d_293_v131_: _dafny.Array
                nw38_ = _dafny.Array(_dafny.MultiSet({}), 24)
                d_293_v131_ = nw38_
                d_294_v132_: _dafny.Seq
                d_294_v132_ = _dafny.SeqWithoutIsStrInference([d_121_v3_, d_121_v3_, _dafny.SeqWithoutIsStrInference([d_150_v29_, (0) - (d_266_i5_)])])
                index33_ = default__.safeIndex(194, (d_293_v131_).length(0))
                (d_293_v131_)[index33_] = _dafny.MultiSet(d_294_v132_)
                index34_ = default__.safeIndex(194, (d_293_v131_).length(0))
                (d_293_v131_)[index34_] = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_266_i5_, 683, len(d_121_v3_)])])
                (d_135_globalState_).f22 = (default__.fm1(d_150_v29_, d_131_v12_, _dafny.SeqWithoutIsStrInference([d_286_v125_ for d_295_i7_ in range(default__.abs(807))]), default__.fm0(d_266_i5_, d_131_v12_, d_135_globalState_), d_135_globalState_)) + (228)
                d_296_v133_: C0
                nw39_ = C0()
                nw39_.ctor__()
                d_296_v133_ = nw39_
            elif True:
                d_297_v134_: _dafny.Array
                nw40_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_297_v134_ = nw40_
                d_298_v135_: _dafny.Array
                def lambda9_(d_299_i8_):
                    return True

                init5_ = lambda9_
                nw41_ = _dafny.Array(None, 19)
                for i0_5_ in range(nw41_.length(0)):
                    nw41_[i0_5_] = init5_(i0_5_)
                d_298_v135_ = nw41_
                d_300_v136_: _dafny.Array
                nw42_ = _dafny.Array(None, 12)
                nw42_[int(0)] = d_298_v135_
                nw42_[int(1)] = d_298_v135_
                nw42_[int(2)] = d_298_v135_
                nw42_[int(3)] = d_298_v135_
                nw42_[int(4)] = d_298_v135_
                nw42_[int(5)] = d_298_v135_
                nw42_[int(6)] = d_298_v135_
                nw42_[int(7)] = d_298_v135_
                nw42_[int(8)] = d_298_v135_
                nw42_[int(9)] = d_298_v135_
                nw42_[int(10)] = d_298_v135_
                nw42_[int(11)] = d_298_v135_
                d_300_v136_ = nw42_
                index35_ = default__.safeIndex(944, (d_297_v134_).length(0))
                (d_297_v134_)[index35_] = d_300_v136_
                d_301_v137_: _dafny.Map
                d_301_v137_ = _dafny.Map({d_123_v5_: d_123_v5_})
                index36_ = default__.safeIndex(944, (d_297_v134_).length(0))
                rhs39_ = ((d_131_v12_) and (default__.fm0(len(d_301_v137_), d_131_v12_, d_135_globalState_)) if not(True) else d_131_v12_)
                rhs40_ = d_300_v136_
                rhs41_ = d_118_v0_
                lhs30_ = d_297_v134_
                lhs31_ = default__.safeIndex(944, (d_297_v134_).length(0))
                lhs32_ = d_135_globalState_
                d_131_v12_ = rhs39_
                lhs30_[lhs31_] = rhs40_
                lhs32_.f25 = rhs41_
                d_302_v138_: C0
                nw43_ = C0()
                nw43_.ctor__()
                d_302_v138_ = nw43_
                index37_ = default__.safeIndex(854, (d_123_v5_).length(0))
                (d_123_v5_)[index37_] = 964
                d_303_v139_: _dafny.Map
                d_303_v139_ = _dafny.Map({d_302_v138_: d_302_v138_})
                index38_ = default__.safeIndex(854, (d_123_v5_).length(0))
                rhs42_ = default__.fm0(d_150_v29_, d_131_v12_, d_135_globalState_)
                rhs43_ = ((d_303_v139_)[d_302_v138_] if (d_302_v138_) in (d_303_v139_) else d_302_v138_)
                rhs44_ = d_150_v29_
                rhs45_ = (d_121_v3_)[default__.safeIndex(d_118_v0_, len(d_121_v3_))]
                rhs46_ = (0) - (d_118_v0_)
                lhs33_ = d_135_globalState_
                lhs34_ = d_135_globalState_
                lhs35_ = d_135_globalState_
                lhs36_ = d_123_v5_
                lhs37_ = default__.safeIndex(854, (d_123_v5_).length(0))
                lhs33_.f24 = rhs42_
                d_302_v138_ = rhs43_
                lhs34_.f22 = rhs44_
                lhs35_.f25 = rhs45_
                lhs36_[lhs37_] = rhs46_
                d_304_v140_: D4
                d_304_v140_ = D4_DC10(d_302_v138_)
                d_305_v141_: _dafny.Map
                d_305_v141_ = _dafny.Map({d_304_v140_: (d_123_v5_)[default__.safeIndex(854, (d_123_v5_).length(0))]})
                d_305_v141_ = (d_305_v141_).set(d_304_v140_, (d_123_v5_)[default__.safeIndex(854, (d_123_v5_).length(0))])
                d_306_v142_: C0
                nw44_ = C0()
                nw44_.ctor__()
                d_306_v142_ = nw44_
            d_307_v143_: _dafny.Array
            nw45_ = _dafny.Array(_dafny.Map({}), 19)
            d_307_v143_ = nw45_
            d_308_v144_: _dafny.Map
            d_308_v144_ = _dafny.Map({not(d_131_v12_): d_150_v29_})
            index39_ = default__.safeIndex(937, (d_307_v143_).length(0))
            (d_307_v143_)[index39_] = d_308_v144_
            d_309_v145_: D5
            d_309_v145_ = D5_DC12(d_308_v144_)
            pat_let_tv5_ = d_308_v144_
            pat_let_tv6_ = d_131_v12_
            pat_let_tv7_ = d_308_v144_
            index40_ = default__.safeIndex(937, (d_307_v143_).length(0))
            def iife24_(_pat_let8_0):
                def iife25_(d_310_dt__update__tmp_h3_):
                    def iife26_(_pat_let9_0):
                        def iife27_(d_311_dt__update_hcf18_h0_):
                            return D5_DC12(d_311_dt__update_hcf18_h0_)
                        return iife27_(_pat_let9_0)
                    return iife26_(pat_let_tv5_)
                return iife25_(_pat_let8_0)
            def iife23_(_pat_let7_0):
                def iife28_(d_312_dt__update__tmp_h4_):
                    def iife29_(_pat_let10_0):
                        def iife30_(d_313_dt__update_hcf18_h1_):
                            return D5_DC12(d_313_dt__update_hcf18_h1_)
                        return iife30_(_pat_let10_0)
                    return iife29_(_dafny.Map({pat_let_tv6_: d_266_i5_}))
                return iife28_(_pat_let7_0)
            def iife22_(_pat_let6_0):
                def iife31_(d_314_dt__update__tmp_h5_):
                    def iife32_(_pat_let11_0):
                        def iife33_(d_315_dt__update_hcf18_h2_):
                            return D5_DC12(d_315_dt__update_hcf18_h2_)
                        return iife33_(_pat_let11_0)
                    return iife32_(pat_let_tv7_)
                return iife31_(_pat_let6_0)
            (d_307_v143_)[index40_] = (iife22_(iife23_(iife24_(d_309_v145_)))).cf18
            d_316_v146_: _dafny.MultiSet
            d_316_v146_ = _dafny.MultiSet([d_131_v12_, d_131_v12_])
            d_317_v147_: _dafny.Seq
            d_317_v147_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True]), d_316_v146_, d_316_v146_, d_316_v146_])
            (d_135_globalState_).f14 = (d_131_v12_ if (d_126_v8_).issubset(d_126_v8_) else ((d_317_v147_)[default__.safeIndex(d_150_v29_, len(d_317_v147_))]).isdisjoint(d_316_v146_))
        d_318_v148_: C0
        nw46_ = C0()
        nw46_.ctor__()
        d_318_v148_ = nw46_
        d_319_v149_: _dafny.Seq
        d_319_v149_ = _dafny.SeqWithoutIsStrInference([d_318_v148_, d_318_v148_])
        d_320_v150_: _dafny.Map
        d_320_v150_ = _dafny.Map({d_118_v0_: d_127_v9_})
        d_318_v148_ = (d_319_v149_)[default__.safeIndex(default__.fm1(d_150_v29_, d_131_v12_, ((d_320_v150_)[d_118_v0_] if (d_118_v0_) in (d_320_v150_) else _dafny.SeqWithoutIsStrInference([d_136_v16_ for d_321_i9_ in range(default__.abs(-209))])), d_131_v12_, d_135_globalState_), len(d_319_v149_))]
        index41_ = default__.safeIndex(998, (d_123_v5_).length(0))
        (d_123_v5_)[index41_] = (0) - (d_150_v29_)
        index42_ = default__.safeIndex(998, (d_123_v5_).length(0))
        rhs47_ = d_118_v0_
        rhs48_ = d_152_v31_
        rhs49_ = (d_118_v0_) <= (d_118_v0_)
        rhs50_ = d_131_v12_
        lhs38_ = d_123_v5_
        lhs39_ = default__.safeIndex(998, (d_123_v5_).length(0))
        lhs40_ = d_135_globalState_
        lhs41_ = d_135_globalState_
        lhs38_[lhs39_] = rhs47_
        d_136_v16_ = rhs48_
        lhs40_.f3 = rhs49_
        lhs41_.f24 = rhs50_
        d_322_v151_: _dafny.Map
        d_322_v151_ = _dafny.Map({(d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))]: d_131_v12_})
        d_323_v152_: _dafny.MultiSet
        d_323_v152_ = _dafny.MultiSet([((d_322_v151_)[d_150_v29_] if (d_150_v29_) in (d_322_v151_) else not(d_131_v12_))])
        if (_dafny.MultiSet((d_149_v28_) + (d_149_v28_))).ispropersubset(d_323_v152_):
            (d_135_globalState_).f24 = True
            (d_135_globalState_).f22 = (d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))]
            d_324_v153_: _dafny.Map
            d_324_v153_ = _dafny.Map({(d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))]: (d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))]})
            d_325_v154_: _dafny.Map
            d_325_v154_ = _dafny.Map({d_118_v0_: D3_DC9(False, d_324_v153_)})
            d_326_v155_: _dafny.Array
            nw47_ = _dafny.Array(False, 15)
            d_326_v155_ = nw47_
            d_327_v156_: D4
            d_327_v156_ = D4_DC11(d_325_v154_, d_326_v155_)
            d_328_v157_: _dafny.Map
            d_328_v157_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_131_v12_]): d_327_v156_})
            d_329_v158_: _dafny.Map
            d_329_v158_ = _dafny.Map({d_118_v0_: len(d_328_v157_)})
            source6_ = D2_DC4((d_329_v158_) | (d_324_v153_))
            if source6_.is_DC5:
                d_330___mcc_h15_ = source6_.cf6
                d_331___mcc_h16_ = source6_.cf7
                d_332___mcc_h17_ = source6_.cf8
                d_333___mcc_h18_ = source6_.cf9
                d_334___mcc_h19_ = source6_.cf10
                d_335_cf10_ = d_334___mcc_h19_
                d_336_cf9_ = d_333___mcc_h18_
                d_337_cf8_ = d_332___mcc_h17_
                d_338_cf7_ = d_331___mcc_h16_
                d_339_cf6_ = d_330___mcc_h15_
                d_340_v159_: _dafny.Seq
                d_341_v160_: int
                d_342_v161_: _dafny.Array
                d_343_v162_: str
                out32_: _dafny.Seq
                out33_: int
                out34_: _dafny.Array
                out35_: str
                out32_, out33_, out34_, out35_ = default__.m0(default__.fm0((d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))], d_339_cf6_, d_135_globalState_), d_135_globalState_)
                d_340_v159_ = out32_
                d_341_v160_ = out33_
                d_342_v161_ = out34_
                d_343_v162_ = out35_
                d_344_v163_: D6
                d_344_v163_ = D6_DC16(d_119_v1_)
                d_345_v164_: _dafny.Map
                d_345_v164_ = d_132_v13_
                pat_let_tv8_ = d_119_v1_
                def iife34_(_pat_let12_0):
                    def iife35_(d_346_dt__update__tmp_h6_):
                        def iife36_(_pat_let13_0):
                            def iife37_(d_347_dt__update_hcf22_h0_):
                                return D6_DC16(d_347_dt__update_hcf22_h0_)
                            return iife37_(_pat_let13_0)
                        return iife36_(pat_let_tv8_)
                    return iife35_(_pat_let12_0)
                rhs51_ = (iife34_(d_344_v163_)).cf22
                rhs52_ = d_131_v12_
                rhs53_ = d_319_v149_
                rhs54_ = len((d_345_v164_))
                lhs42_ = d_135_globalState_
                lhs43_ = d_135_globalState_
                d_119_v1_ = rhs51_
                lhs42_.f24 = rhs52_
                d_319_v149_ = rhs53_
                lhs43_.f9 = rhs54_
                (d_135_globalState_).f14 = d_131_v12_
                d_348_v165_: C0
                nw48_ = C0()
                nw48_.ctor__()
                d_348_v165_ = nw48_
            elif source6_.is_DC6:
                d_349_v166_: _dafny.Seq
                d_350_v167_: int
                d_351_v168_: _dafny.Array
                d_352_v169_: str
                out36_: _dafny.Seq
                out37_: int
                out38_: _dafny.Array
                out39_: str
                out36_, out37_, out38_, out39_ = default__.m0(d_131_v12_, d_135_globalState_)
                d_349_v166_ = out36_
                d_350_v167_ = out37_
                d_351_v168_ = out38_
                d_352_v169_ = out39_
                index43_ = default__.safeIndex(756, (d_326_v155_).length(0))
                (d_326_v155_)[index43_] = d_131_v12_
                index44_ = default__.safeIndex(756, (d_326_v155_).length(0))
                (d_326_v155_)[index44_] = (d_150_v29_) <= (len(d_119_v1_))
                d_149_v28_ = d_349_v166_
                (d_135_globalState_).f9 = 102
            elif source6_.is_DC4:
                d_353___mcc_h20_ = source6_.cf5
                d_354_cf5_ = d_353___mcc_h20_
                (d_135_globalState_).f21 = d_127_v9_
                (d_135_globalState_).f24 = (default__.safeDivisionInt(d_150_v29_, (d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))])) > (len(d_149_v28_))
                d_355_v170_: _dafny.Array
                nw49_ = _dafny.Array(_dafny.Map({}), 1)
                d_355_v170_ = nw49_
                d_356_v171_: _dafny.Map
                d_356_v171_ = _dafny.Map({d_318_v148_: d_255_v104_})
                d_357_v172_: D8
                d_357_v172_ = D8_DC19(((d_356_v171_)[d_318_v148_] if (d_318_v148_) in (d_356_v171_) else d_255_v104_))
                index45_ = default__.safeIndex(224, (d_355_v170_).length(0))
                (d_355_v170_)[index45_] = ((d_357_v172_).cf24) | (d_255_v104_)
                d_358_v173_: _dafny.Seq
                d_358_v173_ = _dafny.SeqWithoutIsStrInference([(d_255_v104_).set((d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))], d_152_v31_), d_255_v104_, d_255_v104_])
                index46_ = default__.safeIndex(224, (d_355_v170_).length(0))
                (d_355_v170_)[index46_] = (d_255_v104_) | ((d_358_v173_)[default__.safeIndex((d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))], len(d_358_v173_))])
                d_359_v174_: _dafny.Map
                d_359_v174_ = _dafny.Map({d_149_v28_: (d_121_v3_) + (d_121_v3_)})
                d_359_v174_ = (d_359_v174_).set((d_149_v28_) + (d_149_v28_), (d_121_v3_) + (d_121_v3_))
            elif True:
                d_360___mcc_h21_ = source6_.cf11
                d_361_cf11_ = d_360___mcc_h21_
                d_362_v175_: _dafny.Seq
                d_362_v175_ = _dafny.SeqWithoutIsStrInference([default__.fm7(d_131_v12_, d_135_globalState_)])
                (d_135_globalState_).f14 = (d_362_v175_) <= ((default__.fm8(d_135_globalState_)) + (d_362_v175_))
                d_363_v176_: _dafny.Seq
                d_364_v177_: int
                d_365_v178_: _dafny.Array
                d_366_v179_: str
                out40_: _dafny.Seq
                out41_: int
                out42_: _dafny.Array
                out43_: str
                out40_, out41_, out42_, out43_ = default__.m0(d_131_v12_, d_135_globalState_)
                d_363_v176_ = out40_
                d_364_v177_ = out41_
                d_365_v178_ = out42_
                d_366_v179_ = out43_
                d_152_v31_ = default__.fm4(d_366_v179_, d_149_v28_, d_135_globalState_)
                rhs55_ = (-121) * ((d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))])
                rhs56_ = d_131_v12_
                lhs44_ = d_135_globalState_
                lhs45_ = d_135_globalState_
                lhs44_.f25 = rhs55_
                lhs45_.f24 = rhs56_
            (d_135_globalState_).f14 = not(d_131_v12_)
            if (_dafny.MultiSet(d_149_v28_)).issubset(d_323_v152_):
                d_152_v31_ = d_152_v31_
                d_152_v31_ = (d_136_v16_ if False else d_152_v31_)
                d_123_v5_ = d_151_v30_
                d_367_v180_: C0
                nw50_ = C0()
                nw50_.ctor__()
                d_367_v180_ = nw50_
                d_368_v181_: _dafny.Set
                d_368_v181_ = _dafny.Set({True, (d_150_v29_) in (d_121_v3_), d_131_v12_, (d_118_v0_) != (len(d_149_v28_)), (d_131_v12_) == (d_131_v12_)})
                index47_ = default__.safeIndex(998, (d_123_v5_).length(0))
                rhs57_ = ((d_329_v158_)[d_150_v29_] if (d_150_v29_) in (d_329_v158_) else (d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))])
                rhs58_ = d_123_v5_
                rhs59_ = ((d_368_v181_) | (d_368_v181_)) | (d_368_v181_)
                rhs60_ = d_118_v0_
                rhs61_ = d_131_v12_
                lhs46_ = d_123_v5_
                lhs47_ = default__.safeIndex(998, (d_123_v5_).length(0))
                lhs48_ = d_135_globalState_
                lhs49_ = d_135_globalState_
                lhs46_[lhs47_] = rhs57_
                d_123_v5_ = rhs58_
                d_368_v181_ = rhs59_
                lhs48_.f25 = rhs60_
                lhs49_.f24 = rhs61_
            elif True:
                d_369_v182_: C0
                nw51_ = C0()
                nw51_.ctor__()
                d_369_v182_ = nw51_
                (d_135_globalState_).f25 = (0) - (d_150_v29_)
                (d_135_globalState_).f25 = len((_dafny.SeqWithoutIsStrInference([d_131_v12_, d_131_v12_])) + (d_149_v28_))
                (d_135_globalState_).f25 = d_150_v29_
                d_370_v183_: C0
                nw52_ = C0()
                nw52_.ctor__()
                d_370_v183_ = nw52_
        elif True:
            d_371_v185_: _dafny.Map
            def iife38_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in (d_121_v3_).Elements:
                    d_372_v184_: int = compr_10_
                    if (d_372_v184_) in (d_121_v3_):
                        coll10_[(d_372_v184_) + (d_118_v0_)] = (d_126_v8_).cardinality
                return _dafny.Map(coll10_)
            d_371_v185_ = _dafny.Map({d_150_v29_: len(iife38_()
            )})
            rhs62_ = (0) - (((len(d_371_v185_)) + ((d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))])) - (default__.safeDivisionInt(d_150_v29_, d_150_v29_)))
            rhs63_ = d_131_v12_
            rhs64_ = d_119_v1_
            lhs50_ = d_135_globalState_
            lhs51_ = d_135_globalState_
            lhs50_.f25 = rhs62_
            lhs51_.f3 = rhs63_
            d_119_v1_ = rhs64_
            d_373_v186_: _dafny.Seq
            d_374_v187_: int
            d_375_v188_: _dafny.Array
            d_376_v189_: str
            out44_: _dafny.Seq
            out45_: int
            out46_: _dafny.Array
            out47_: str
            out44_, out45_, out46_, out47_ = default__.m0(d_131_v12_, d_135_globalState_)
            d_373_v186_ = out44_
            d_374_v187_ = out45_
            d_375_v188_ = out46_
            d_376_v189_ = out47_
            d_377_v190_: _dafny.Seq
            d_378_v191_: int
            d_379_v192_: _dafny.Array
            d_380_v193_: str
            out48_: _dafny.Seq
            out49_: int
            out50_: _dafny.Array
            out51_: str
            out48_, out49_, out50_, out51_ = default__.m0(not (d_131_v12_) or (d_131_v12_), d_135_globalState_)
            d_377_v190_ = out48_
            d_378_v191_ = out49_
            d_379_v192_ = out50_
            d_380_v193_ = out51_
            d_381_v194_: _dafny.Array
            nw53_ = _dafny.Array(None, 10)
            d_381_v194_ = nw53_
            index48_ = default__.safeIndex(53, (d_381_v194_).length(0))
            (d_381_v194_)[index48_] = d_318_v148_
            d_382_v195_: _dafny.Array
            nw54_ = _dafny.Array(None, 14)
            nw54_[int(0)] = d_131_v12_
            nw54_[int(1)] = d_131_v12_
            nw54_[int(2)] = (d_131_v12_ if d_131_v12_ else d_131_v12_)
            nw54_[int(3)] = (not(d_131_v12_)) and (d_131_v12_)
            nw54_[int(4)] = True
            nw54_[int(5)] = not(d_131_v12_)
            nw54_[int(6)] = d_131_v12_
            nw54_[int(7)] = d_131_v12_
            nw54_[int(8)] = d_131_v12_
            nw54_[int(9)] = d_131_v12_
            nw54_[int(10)] = d_131_v12_
            nw54_[int(11)] = (d_131_v12_ if d_131_v12_ else d_131_v12_)
            nw54_[int(12)] = (d_150_v29_) != (d_150_v29_)
            nw54_[int(13)] = d_131_v12_
            d_382_v195_ = nw54_
            index49_ = default__.safeIndex(198, (d_382_v195_).length(0))
            (d_382_v195_)[index49_] = d_131_v12_
            d_383_v196_: _dafny.Set
            d_383_v196_ = _dafny.Set({d_127_v9_, d_127_v9_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxdpjq")), default__.fm2((d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))], d_150_v29_, False, (d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))], d_135_globalState_), d_127_v9_})
            index50_ = default__.safeIndex(53, (d_381_v194_).length(0))
            index51_ = default__.safeIndex(198, (d_382_v195_).length(0))
            rhs65_ = d_318_v148_
            rhs66_ = d_129_v11_
            rhs67_ = d_121_v3_
            rhs68_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))) != (_dafny.SeqWithoutIsStrInference([d_380_v193_ for d_384_i10_ in range(default__.abs(781))]))) in ((d_377_v190_) + (d_377_v190_))
            rhs69_ = (d_383_v196_).ispropersubset(d_383_v196_)
            lhs52_ = d_381_v194_
            lhs53_ = default__.safeIndex(53, (d_381_v194_).length(0))
            lhs54_ = d_135_globalState_
            lhs55_ = d_382_v195_
            lhs56_ = default__.safeIndex(198, (d_382_v195_).length(0))
            lhs52_[lhs53_] = rhs65_
            d_129_v11_ = rhs66_
            d_121_v3_ = rhs67_
            lhs54_.f3 = rhs68_
            lhs55_[lhs56_] = rhs69_
            if False:
                (d_135_globalState_).f9 = (0) - (d_374_v187_)
                (d_135_globalState_).f25 = (d_318_v148_).fm5(d_135_globalState_)
                d_385_v197_: _dafny.Seq
                d_386_v198_: int
                d_387_v199_: _dafny.Array
                d_388_v200_: str
                out52_: _dafny.Seq
                out53_: int
                out54_: _dafny.Array
                out55_: str
                out52_, out53_, out54_, out55_ = default__.m0(d_131_v12_, d_135_globalState_)
                d_385_v197_ = out52_
                d_386_v198_ = out53_
                d_387_v199_ = out54_
                d_388_v200_ = out55_
                d_389_v201_: C0
                nw55_ = C0()
                nw55_.ctor__()
                d_389_v201_ = nw55_
                d_390_v202_: _dafny.Map
                d_390_v202_ = _dafny.Map({(d_382_v195_)[default__.safeIndex(198, (d_382_v195_).length(0))]: d_118_v0_})
                d_390_v202_ = d_390_v202_
            elif True:
                (d_135_globalState_).f25 = len(d_377_v190_)
                (d_135_globalState_).f24 = ((d_378_v191_) + ((0) - (d_150_v29_))) >= (len(d_121_v3_))
                d_391_v203_: C0
                nw56_ = C0()
                nw56_.ctor__()
                d_391_v203_ = nw56_
                (d_135_globalState_).f14 = (d_382_v195_)[default__.safeIndex(198, (d_382_v195_).length(0))]
                (d_135_globalState_).f24 = (d_382_v195_)[default__.safeIndex(198, (d_382_v195_).length(0))]
        d_392_v204_: _dafny.Set
        d_392_v204_ = _dafny.Set({(d_149_v28_)[default__.safeIndex(len(d_127_v9_), len(d_149_v28_))], d_131_v12_, d_131_v12_})
        if ((d_392_v204_) - (d_392_v204_)).isdisjoint(d_392_v204_):
            d_393_v205_: _dafny.Array
            nw57_ = _dafny.Array(False, 17)
            d_393_v205_ = nw57_
            index52_ = default__.safeIndex(644, (d_393_v205_).length(0))
            (d_393_v205_)[index52_] = d_131_v12_
            d_394_v206_: _dafny.Map
            d_394_v206_ = _dafny.Map({len(d_322_v151_): -212})
            d_395_v207_: D2
            d_395_v207_ = D2_DC4(d_394_v206_)
            index53_ = default__.safeIndex(644, (d_393_v205_).length(0))
            (d_393_v205_)[index53_] = default__.fm0(len((default__.fm2((d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))], (d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))], d_131_v12_, d_150_v29_, d_135_globalState_)) + (d_127_v9_)), ((d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))]) not in ((d_395_v207_).cf5), d_135_globalState_)
            (d_135_globalState_).f22 = (0) - ((d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))])
            d_396_v208_: _dafny.Seq
            d_397_v209_: int
            d_398_v210_: _dafny.Array
            d_399_v211_: str
            out56_: _dafny.Seq
            out57_: int
            out58_: _dafny.Array
            out59_: str
            out56_, out57_, out58_, out59_ = default__.m0((d_393_v205_)[default__.safeIndex(644, (d_393_v205_).length(0))], d_135_globalState_)
            d_396_v208_ = out56_
            d_397_v209_ = out57_
            d_398_v210_ = out58_
            d_399_v211_ = out59_
            d_151_v30_ = d_398_v210_
            (d_135_globalState_).f24 = (d_393_v205_)[default__.safeIndex(644, (d_393_v205_).length(0))]
        elif True:
            if not ((len(d_149_v28_)) != ((d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))])) or ((d_149_v28_)[default__.safeIndex(d_150_v29_, len(d_149_v28_))]):
                d_400_v212_: _dafny.Array
                nw58_ = _dafny.Array(None, 19)
                nw58_[int(0)] = d_125_v7_
                nw58_[int(1)] = d_125_v7_
                nw58_[int(2)] = d_125_v7_
                nw58_[int(3)] = d_125_v7_
                nw58_[int(4)] = (d_125_v7_ if d_131_v12_ else d_125_v7_)
                nw58_[int(5)] = d_125_v7_
                nw58_[int(6)] = d_125_v7_
                nw58_[int(7)] = d_125_v7_
                nw58_[int(8)] = d_125_v7_
                nw58_[int(9)] = d_125_v7_
                nw58_[int(10)] = d_125_v7_
                nw58_[int(11)] = d_125_v7_
                nw58_[int(12)] = d_125_v7_
                nw58_[int(13)] = (d_125_v7_ if d_131_v12_ else d_125_v7_)
                nw58_[int(14)] = d_125_v7_
                nw58_[int(15)] = d_125_v7_
                nw58_[int(16)] = d_125_v7_
                nw58_[int(17)] = d_125_v7_
                nw58_[int(18)] = d_125_v7_
                d_400_v212_ = nw58_
                index54_ = default__.safeIndex(637, (d_400_v212_).length(0))
                (d_400_v212_)[index54_] = d_125_v7_
                index55_ = default__.safeIndex(637, (d_400_v212_).length(0))
                (d_400_v212_)[index55_] = d_125_v7_
                (d_135_globalState_).f24 = (d_131_v12_) and (not (d_131_v12_) or (d_131_v12_))
                d_401_v213_: D6
                d_401_v213_ = D6_DC17()
                rhs70_ = d_127_v9_
                rhs71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smrlm"))
                rhs72_ = d_401_v213_
                lhs57_ = d_135_globalState_
                lhs58_ = d_135_globalState_
                lhs57_.f21 = rhs70_
                lhs58_.f21 = rhs71_
                d_401_v213_ = rhs72_
                d_151_v30_ = d_151_v30_
                (d_135_globalState_).f24 = default__.fm0(d_150_v29_, default__.fm0(d_118_v0_, d_131_v12_, d_135_globalState_), d_135_globalState_)
            elif True:
                d_402_v214_: _dafny.Seq
                d_403_v215_: int
                d_404_v216_: _dafny.Array
                d_405_v217_: str
                out60_: _dafny.Seq
                out61_: int
                out62_: _dafny.Array
                out63_: str
                out60_, out61_, out62_, out63_ = default__.m0(d_131_v12_, d_135_globalState_)
                d_402_v214_ = out60_
                d_403_v215_ = out61_
                d_404_v216_ = out62_
                d_405_v217_ = out63_
                (d_135_globalState_).f14 = d_131_v12_
                d_406_v218_: _dafny.Map
                d_406_v218_ = _dafny.Map({d_404_v216_: False})
                d_406_v218_ = (d_406_v218_) | (d_406_v218_)
                d_407_v219_: _dafny.Seq
                d_408_v220_: int
                d_409_v221_: _dafny.Array
                d_410_v222_: str
                out64_: _dafny.Seq
                out65_: int
                out66_: _dafny.Array
                out67_: str
                out64_, out65_, out66_, out67_ = default__.m0(d_131_v12_, d_135_globalState_)
                d_407_v219_ = out64_
                d_408_v220_ = out65_
                d_409_v221_ = out66_
                d_410_v222_ = out67_
                (d_135_globalState_).f21 = d_127_v9_
            d_411_v223_: _dafny.Seq
            d_412_v224_: int
            d_413_v225_: _dafny.Array
            d_414_v226_: str
            out68_: _dafny.Seq
            out69_: int
            out70_: _dafny.Array
            out71_: str
            out68_, out69_, out70_, out71_ = default__.m0((d_392_v204_).isdisjoint(d_392_v204_), d_135_globalState_)
            d_411_v223_ = out68_
            d_412_v224_ = out69_
            d_413_v225_ = out70_
            d_414_v226_ = out71_
            d_415_v227_: _dafny.Seq
            d_416_v228_: int
            d_417_v229_: _dafny.Array
            d_418_v230_: str
            out72_: _dafny.Seq
            out73_: int
            out74_: _dafny.Array
            out75_: str
            out72_, out73_, out74_, out75_ = default__.m0(d_131_v12_, d_135_globalState_)
            d_415_v227_ = out72_
            d_416_v228_ = out73_
            d_417_v229_ = out74_
            d_418_v230_ = out75_
            (d_135_globalState_).f24 = d_131_v12_
            (d_135_globalState_).f14 = d_131_v12_
        (d_135_globalState_).f9 = (d_123_v5_)[default__.safeIndex(998, (d_123_v5_).length(0))]
        d_118_v0_ = (d_318_v148_).fm5(d_135_globalState_)
        _dafny.print(_dafny.string_of(d_118_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v1_) == (_dafny.Set({691}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v2_) == (_dafny.Map({_dafny.Set({691}): 691}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v3_) == (_dafny.SeqWithoutIsStrInference([691, 691]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v4_) == (_dafny.Map({1: _dafny.SeqWithoutIsStrInference([691, 691])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v5_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_v6_).cf0)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[0])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[1])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[2])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_125_v7_)[3])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v8_) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_127_v9_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_128_v10_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v11_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_131_v12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v13_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_133_v14_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_133_v14_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v15_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference([691, 691])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_.f0) == (_dafny.Map({_dafny.Set({691}): 691}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1) == (_dafny.Set({691}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_135_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f6) == (_dafny.Map({1: _dafny.SeqWithoutIsStrInference([691, 691])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_135_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f10) == (_dafny.Set({1, 691, 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_135_globalState_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f17)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f19) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_135_globalState_).f20)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_135_globalState_).f20)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_135_globalState_.f21).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_135_globalState_.f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_.f23) == (_dafny.Set({_dafny.SeqWithoutIsStrInference([691, 691])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_135_globalState_.f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_135_globalState_.f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_136_v16_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v17_) == (_dafny.Map({_dafny.Map({True: 691}): _dafny.CodePoint('j')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_i1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v28_) == (_dafny.SeqWithoutIsStrInference([False, True, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_150_v29_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v30_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v30_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v30_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v30_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v30_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v30_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v30_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v30_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_152_v31_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_255_v104_) == (_dafny.Map({-119716: _dafny.CodePoint('f')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_256_i4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_319_v149_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_320_v150_) == (_dafny.Map({691: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "px"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_322_v151_) == (_dafny.Map({691: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_323_v152_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_392_v204_) == (_dafny.Set({False, True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False)
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
        return lambda: _dafny.CodePoint('D')
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC5(False, False, False, int(0), False)
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

class D2_DC5(D2, NamedTuple('DC5', [('cf6', Any), ('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(False, _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC9(D3, NamedTuple('DC9', [('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(_dafny.Map({}), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC11(D4, NamedTuple('DC11', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC13(D5, NamedTuple('DC13', [])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13)
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC17(D6, NamedTuple('DC17', [])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17)
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC18(D7, NamedTuple('DC18', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC20(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)

class D8_DC20(D8, NamedTuple('DC20', [('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: _dafny.Map = _dafny.Map({})
        self.f3: bool = False
        self.f9: int = int(0)
        self.f14: bool = False
        self.f21: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f22: int = int(0)
        self.f23: _dafny.Set = _dafny.Set({})
        self.f24: bool = False
        self.f25: int = int(0)
        self._f1: _dafny.Set = _dafny.Set({})
        self._f2: int = int(0)
        self._f4: int = int(0)
        self._f5: str = _dafny.CodePoint('D')
        self._f6: _dafny.Map = _dafny.Map({})
        self._f7: int = int(0)
        self._f8: int = int(0)
        self._f10: _dafny.Set = _dafny.Set({})
        self._f11: int = int(0)
        self._f12: bool = False
        self._f13: bool = False
        self._f15: int = int(0)
        self._f16: bool = False
        self._f17: _dafny.Array = _dafny.Array(None, 0)
        self._f18: int = int(0)
        self._f19: _dafny.Map = _dafny.Map({})
        self._f20: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self).f21 = f21
        (self).f22 = f22
        (self).f23 = f23
        (self).f24 = f24
        (self).f25 = f25

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
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm5(self, globalState):
        return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpbp")))) - ((704) * (515))

