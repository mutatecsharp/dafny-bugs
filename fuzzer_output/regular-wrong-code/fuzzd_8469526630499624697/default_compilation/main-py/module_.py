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
        return ((_dafny.Set({True, False}) if True else _dafny.Set({False}))) | (_dafny.Set({not(True), True, False}))

    @staticmethod
    def fm3(globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: _dafny.Map
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([False]), _dafny.MultiSet([True, False])])).cardinality: len(_dafny.SeqWithoutIsStrInference([not(False)]))}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))): len(_dafny.Map({(_dafny.MultiSet([679, -336])).cardinality: not(False)}))}), _dafny.Map({793: len(_dafny.Map({True: 8}))}), _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False]))).cardinality: 76}), _dafny.Map({len(_dafny.Set({True})): len(_dafny.SeqWithoutIsStrInference([False]))})])).Elements:
                d_0_v0_: _dafny.Map = compr_0_
                if (d_0_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([False]), _dafny.MultiSet([True, False])])).cardinality: len(_dafny.SeqWithoutIsStrInference([not(False)]))}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))): len(_dafny.Map({(_dafny.MultiSet([679, -336])).cardinality: not(False)}))}), _dafny.Map({793: len(_dafny.Map({True: 8}))}), _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False]))).cardinality: 76}), _dafny.Map({len(_dafny.Set({True})): len(_dafny.SeqWithoutIsStrInference([False]))})])):
                    coll0_ = coll0_.union(_dafny.Set([d_0_v0_]))
            return _dafny.Set(coll0_)
        return not(not(((len(iife0_()
        )) + ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True]))])).cardinality)) == ((_dafny.MultiSet([-574, len(_dafny.SeqWithoutIsStrInference([False, not(True)])), -776])).cardinality)))

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([(402) > ((_dafny.MultiSet([False])).cardinality), True, not((False) in (_dafny.MultiSet([not(not(False)), True, True]))), True, False])

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        return True

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        return default__.safeModuloInt(544, len((_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agf")))})) | (_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufyke")))}))))

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        return _dafny.Set({324, 155})

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        return D2_DC7(False)

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        if False:
            return D3_DC12(D3_DC12(D3_DC12(D3_DC12(D3_DC11(False, 707, _dafny.Set({False}))))))
        elif True:
            return D3_DC12(D3_DC11(False, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1_i0_ in range(default__.abs(-504))])), _dafny.Set({False})))

    @staticmethod
    def fm13(globalState):
        if (not(False)) or (False):
            if True:
                return _dafny.CodePoint('y')
            elif True:
                return _dafny.CodePoint('x')
        elif not(not(True)):
            return _dafny.CodePoint('f')
        elif True:
            return _dafny.CodePoint('o')

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxqveefwj"))

    @staticmethod
    def fm15(globalState):
        return (_dafny.Map({len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([not(True)])})): len(_dafny.SeqWithoutIsStrInference([True]))})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xw"))): len(_dafny.Set({False}))}))

    @staticmethod
    def fm16(globalState):
        return _dafny.SeqWithoutIsStrInference([D1_DC5(710), D1_DC5(611), D1_DC5(-796)])

    @staticmethod
    def fm17(p0, globalState):
        if (False) in (_dafny.Set({False, True, not(False)})):
            return (_dafny.Map({_dafny.CodePoint('q'): False})) | (_dafny.Map({_dafny.CodePoint('k'): True}))
        elif True:
            return _dafny.Map({_dafny.CodePoint('d'): True})

    @staticmethod
    def fm18(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([False, False, True, False])) + ((_dafny.SeqWithoutIsStrInference([True, (D12_DC34(not(True))).cf65, True])) + (_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        return D0_DC2((687) + (len(_dafny.SeqWithoutIsStrInference([True]))), (818) <= (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qyetl")) for d_2_i0_ in range(default__.abs(-321))]))), (_dafny.SeqWithoutIsStrInference([8 for d_3_i1_ in range(default__.abs(614))]) if True else _dafny.SeqWithoutIsStrInference([-302 for d_4_i2_ in range(default__.abs(818))])))

    @staticmethod
    def fm22(globalState):
        def iife1_():
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: _dafny.Map
                for compr_3_ in (_dafny.Map({_dafny.Map({False: True}): len(_dafny.Map({True: _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_5_i0_ in range(default__.abs(-951))])): 168})}))})).keys.Elements:
                    d_6_v0_: _dafny.Map = compr_3_
                    if (d_6_v0_) in (_dafny.Map({_dafny.Map({False: True}): len(_dafny.Map({True: _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_5_i0_ in range(default__.abs(-951))])): 168})}))})):
                        coll3_[d_6_v0_] = not(True)
                return _dafny.Map(coll3_)
            coll1_ = _dafny.Set()
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: _dafny.Map
                for compr_2_ in (_dafny.Map({_dafny.Map({False: True}): len(_dafny.Map({True: _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_5_i0_ in range(default__.abs(-951))])): 168})}))})).keys.Elements:
                    d_6_v0_: _dafny.Map = compr_2_
                    if (d_6_v0_) in (_dafny.Map({_dafny.Map({False: True}): len(_dafny.Map({True: _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_5_i0_ in range(default__.abs(-951))])): 168})}))})):
                        coll2_[d_6_v0_] = not(True)
                return _dafny.Map(coll2_)
            compr_1_: _dafny.Map
            for compr_1_ in (iife2_()
            ).keys.Elements:
                d_7_v1_: _dafny.Map = compr_1_
                if (d_7_v1_) in (iife3_()
                ):
                    coll1_ = coll1_.union(_dafny.Set([d_7_v1_]))
            return _dafny.Set(coll1_)
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: _dafny.Map
            for compr_4_ in (_dafny.Map({_dafny.Map({not(True): True}): (0) - (-400)})).keys.Elements:
                d_8_v2_: _dafny.Map = compr_4_
                if (d_8_v2_) in (_dafny.Map({_dafny.Map({not(True): True}): (0) - (-400)})):
                    coll4_ = coll4_.union(_dafny.Set([d_8_v2_]))
            return _dafny.Set(coll4_)
        return (iife1_()
        ) | ((_dafny.Set({_dafny.Map({True: True})}) if True else iife4_()
        ))

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({True: False}) if True else _dafny.Map({False: False}))) | ((_dafny.Map({False: True})) | (_dafny.Map({not(False): True})))

    @staticmethod
    def fm24(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({not(False): (0) - (-285)})) | (_dafny.Map({not(not(not(False))): -46}))) | ((_dafny.Map({False: (_dafny.MultiSet([289])).cardinality})) | (_dafny.Map({False: -844})))

    @staticmethod
    def fm25(globalState):
        source0_ = _dafny.SeqWithoutIsStrInference([D1_DC5(928) for d_9_i0_ in range(default__.abs(381))])
        d_10___mcc_h0_ = source0_
        d_11_cf24_ = d_10___mcc_h0_
        return D1_DC5(752)

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        return D1_DC4(_dafny.CodePoint('a'), (0) - ((len(_dafny.SeqWithoutIsStrInference([True]))) * ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfmrgm")))))), _dafny.CodePoint('q'), (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False}))])): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_12_i0_ in range(default__.abs(-71))])}))) * (961), (_dafny.MultiSet([not(True), True])).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True]))))

    @staticmethod
    def fm27(p0, p1, p2, globalState):
        return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxgruga")))) == (len(_dafny.SeqWithoutIsStrInference([D14_DC40((0) - (-75), True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnsuin")))])))

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbthlure"))

    @staticmethod
    def fm29(globalState):
        return ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([(D0_DC2(len(_dafny.Map({True: D12_DC33(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yuplewtm")))})), True, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bensvahfw")))]))).cf5]))) + (_dafny.SeqWithoutIsStrInference([True, True, False]))

    @staticmethod
    def fm30(p0, globalState):
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(363, 186):
                d_13_v0_: int = compr_5_
                if ((363) <= (d_13_v0_)) and ((d_13_v0_) < (186)):
                    coll5_ = coll5_.union(_dafny.Set([(d_13_v0_) * (-949)]))
            return _dafny.Set(coll5_)
        return ((D16_DC45(_dafny.Set({(0) - ((_dafny.MultiSet([False])).cardinality), 292}))).cf80) - (iife5_()
        )

    @staticmethod
    def fm31(globalState):
        return D6_DC17((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_14_i0_ in range(default__.abs(992))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eanbpr"))), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_15_i1_ in range(default__.abs(156))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xm"))))

    @staticmethod
    def fm32(p0, p1, globalState):
        return ((_dafny.Map({_dafny.Set({len(_dafny.Map({not(False): (_dafny.MultiSet([len(_dafny.Set({not(False)}))])).cardinality}))}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sorttx"))}) if False else _dafny.Map({_dafny.Set({len(_dafny.Set({len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_16_i0_ in range(default__.abs(599))])})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkbghw")))})), 559}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhnnprjd"))}))) | (_dafny.Map({_dafny.Set({len(_dafny.SeqWithoutIsStrInference([not(True), True])), 67}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ug"))}))

    @staticmethod
    def fm33(globalState):
        return _dafny.MultiSet([((0) - (len(_dafny.Map({True: False})))) - (77)])

    @staticmethod
    def fm34(p0, p1, p2, globalState):
        if (788) not in (_dafny.Set({919, len(_dafny.Set({not(False)})), len(_dafny.SeqWithoutIsStrInference([True])), (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sswah")))))})):
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True, False])])
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True])])

    @staticmethod
    def m0(globalState):
        d_17_v0_: bool
        d_17_v0_ = True
        if d_17_v0_:
            d_18_v1_: int
            d_18_v1_ = 146
            (globalState).f1 = (0) - (d_18_v1_)
            d_19_v2_: _dafny.MultiSet
            d_19_v2_ = _dafny.MultiSet([d_18_v1_])
            d_20_v3_: _dafny.Array
            nw0_ = _dafny.Array(None, 19)
            d_20_v3_ = nw0_
            d_21_v4_: _dafny.Map
            d_21_v4_ = _dafny.Map({(d_18_v1_) * (d_18_v1_): d_20_v3_})
            d_22_v5_: _dafny.Seq
            d_22_v5_ = _dafny.SeqWithoutIsStrInference([default__.fm13(globalState)])
            d_23_v6_: str
            d_23_v6_ = _dafny.CodePoint('f')
            d_24_v7_: T0
            nw1_ = C4()
            nw1_.ctor__(len(d_22_v5_), d_17_v0_, d_17_v0_, d_17_v0_, d_23_v6_, d_22_v5_)
            d_24_v7_ = nw1_
            d_25_v8_: _dafny.Map
            d_25_v8_ = _dafny.Map({len(d_22_v5_): d_24_v7_})
            rhs0_ = d_18_v1_
            rhs1_ = (d_19_v2_).set(82, default__.abs(397))
            rhs2_ = (d_21_v4_) | (d_21_v4_)
            rhs3_ = ((d_19_v2_)[default__.safeDivisionInt(len(d_22_v5_), d_18_v1_)] if (default__.safeDivisionInt(len(d_22_v5_), d_18_v1_)) in (d_19_v2_) else d_18_v1_)
            rhs4_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([len(d_25_v8_)])), default__.safeModuloInt(d_18_v1_, d_18_v1_))
            lhs0_ = globalState
            lhs1_ = globalState
            d_18_v1_ = rhs0_
            d_19_v2_ = rhs1_
            d_21_v4_ = rhs2_
            lhs0_.f1 = rhs3_
            lhs1_.f1 = rhs4_
            if not(d_17_v0_):
                (globalState).f1 = (d_18_v1_) * (d_18_v1_)
                d_26_v9_: D6
                d_26_v9_ = D6_DC16(d_19_v2_)
                d_27_v10_: _dafny.Seq
                d_27_v10_ = _dafny.SeqWithoutIsStrInference([d_26_v9_, d_26_v9_, d_26_v9_])
                d_28_v11_: _dafny.Seq
                d_28_v11_ = _dafny.SeqWithoutIsStrInference([d_17_v0_, True, d_17_v0_])
                d_27_v10_ = ((d_27_v10_).set(default__.safeIndex((_dafny.MultiSet(d_28_v11_)).cardinality, len(d_27_v10_)), d_26_v9_)) + (d_27_v10_)
                d_22_v5_ = ((d_24_v7_).f5 if d_17_v0_ else (d_24_v7_).f5)
                d_29_v12_: _dafny.Array
                nw2_ = _dafny.Array(_dafny.MultiSet({}), 18)
                d_29_v12_ = nw2_
                d_30_v13_: _dafny.MultiSet
                d_30_v13_ = _dafny.MultiSet([d_17_v0_])
                index0_ = default__.safeIndex(430, (d_29_v12_).length(0))
                (d_29_v12_)[index0_] = d_30_v13_
                index1_ = default__.safeIndex(430, (d_29_v12_).length(0))
                (d_29_v12_)[index1_] = d_30_v13_
                d_31_v14_: C1
                nw3_ = C1()
                nw3_.ctor__(d_18_v1_, d_17_v0_, d_17_v0_, d_17_v0_, d_23_v6_, d_22_v5_)
                d_31_v14_ = nw3_
                d_32_v15_: _dafny.Seq
                d_32_v15_ = _dafny.SeqWithoutIsStrInference([d_31_v14_, d_31_v14_, d_31_v14_, d_31_v14_])
                d_33_v16_: _dafny.Map
                d_33_v16_ = _dafny.Map({d_32_v15_: d_18_v1_})
                d_34_v17_: _dafny.Map
                d_34_v17_ = _dafny.Map({((d_33_v16_)[_dafny.SeqWithoutIsStrInference([d_31_v14_])] if (_dafny.SeqWithoutIsStrInference([d_31_v14_])) in (d_33_v16_) else d_18_v1_): d_18_v1_})
                d_35_v18_: _dafny.Seq
                d_35_v18_ = _dafny.SeqWithoutIsStrInference([d_18_v1_])
                d_34_v17_ = (d_34_v17_).set(d_18_v1_, len(d_35_v18_))
            elif True:
                d_36_v20_: _dafny.Array
                def lambda0_(d_37_i0_):
                    def iife6_():
                        coll6_ = _dafny.Set()
                        compr_6_: int
                        for compr_6_ in _dafny.IntegerRange(-674, -208):
                            d_38_v19_: int = compr_6_
                            if ((-674) <= (d_38_v19_)) and ((d_38_v19_) < (-208)):
                                coll6_ = coll6_.union(_dafny.Set([(d_38_v19_) + (382)]))
                        return _dafny.Set(coll6_)
                    return default__.safeDivisionInt(d_37_i0_, len(iife6_()
                    ))

                init0_ = lambda0_
                nw4_ = _dafny.Array(None, 18)
                for i0_0_ in range(nw4_.length(0)):
                    nw4_[i0_0_] = init0_(i0_0_)
                d_36_v20_ = nw4_
                index2_ = default__.safeIndex(243, (d_36_v20_).length(0))
                (d_36_v20_)[index2_] = (d_18_v1_ if d_17_v0_ else d_18_v1_)
                d_39_v21_: _dafny.Array
                nw5_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_39_v21_ = nw5_
                d_40_v22_: _dafny.Array
                nw6_ = _dafny.Array(False, 17)
                d_40_v22_ = nw6_
                index3_ = default__.safeIndex(657, (d_40_v22_).length(0))
                (d_40_v22_)[index3_] = (d_19_v2_) == (d_19_v2_)
                index4_ = default__.safeIndex(243, (d_36_v20_).length(0))
                index5_ = default__.safeIndex(657, (d_40_v22_).length(0))
                rhs5_ = d_24_v7_.f4
                rhs6_ = -977
                rhs7_ = (d_39_v21_ if (default__.fm7(d_24_v7_.f4, d_18_v1_, d_17_v0_, globalState) if d_17_v0_ else d_17_v0_) else d_39_v21_)
                rhs8_ = not (False) or (not(not (d_17_v0_) or (d_17_v0_)))
                lhs2_ = d_36_v20_
                lhs3_ = default__.safeIndex(243, (d_36_v20_).length(0))
                lhs4_ = d_40_v22_
                lhs5_ = default__.safeIndex(657, (d_40_v22_).length(0))
                d_23_v6_ = rhs5_
                lhs2_[lhs3_] = rhs6_
                d_39_v21_ = rhs7_
                lhs4_[lhs5_] = rhs8_
                d_23_v6_ = d_24_v7_.f4
                d_17_v0_ = d_17_v0_
                d_41_v23_: _dafny.Array
                def lambda1_(d_42_v7_):
                    def lambda2_(d_43_i1_):
                        return _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itjlfvghe")): d_42_v7_.f4})

                    return lambda2_

                init1_ = lambda1_(d_24_v7_)
                nw7_ = _dafny.Array(None, 3)
                for i0_1_ in range(nw7_.length(0)):
                    nw7_[i0_1_] = init1_(i0_1_)
                d_41_v23_ = nw7_
                d_44_v24_: _dafny.Map
                d_44_v24_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtgg")): d_24_v7_.f4})
                index6_ = default__.safeIndex(463, (d_41_v23_).length(0))
                (d_41_v23_)[index6_] = (d_44_v24_) | (d_44_v24_)
                index7_ = default__.safeIndex(463, (d_41_v23_).length(0))
                (d_41_v23_)[index7_] = d_44_v24_
                d_40_v22_ = d_40_v22_
            d_17_v0_ = default__.fm3(globalState)
            (globalState).f1 = (d_18_v1_) - (d_18_v1_)
        elif True:
            d_17_v0_ = d_17_v0_
            d_45_v25_: _dafny.Array
            def lambda3_(d_46_v0_):
                def lambda4_(d_47_i2_):
                    return d_46_v0_

                return lambda4_

            init2_ = lambda3_(d_17_v0_)
            nw8_ = _dafny.Array(None, 11)
            for i0_2_ in range(nw8_.length(0)):
                nw8_[i0_2_] = init2_(i0_2_)
            d_45_v25_ = nw8_
            d_48_v26_: _dafny.Array
            def lambda5_(d_49_v0_):
                def lambda6_(d_50_i3_):
                    return _dafny.SeqWithoutIsStrInference([d_49_v0_, d_49_v0_])

                return lambda6_

            init3_ = lambda5_(d_17_v0_)
            nw9_ = _dafny.Array(None, 6)
            for i0_3_ in range(nw9_.length(0)):
                nw9_[i0_3_] = init3_(i0_3_)
            d_48_v26_ = nw9_
            d_51_v27_: D1
            d_51_v27_ = D1_DC3(d_48_v26_)
            d_52_v28_: _dafny.Seq
            d_52_v28_ = _dafny.SeqWithoutIsStrInference([d_48_v26_])
            d_53_v29_: int
            d_53_v29_ = 80
            pat_let_tv0_ = d_48_v26_
            d_54_v30_: _dafny.Array
            nw10_ = _dafny.Array(None, 24)
            nw10_[int(0)] = d_51_v27_
            nw10_[int(1)] = d_51_v27_
            nw10_[int(2)] = d_51_v27_
            nw10_[int(3)] = d_51_v27_
            nw10_[int(4)] = d_51_v27_
            nw10_[int(5)] = D1_DC3((d_52_v28_)[default__.safeIndex(d_53_v29_, len(d_52_v28_))])
            nw10_[int(6)] = d_51_v27_
            nw10_[int(7)] = d_51_v27_
            nw10_[int(8)] = d_51_v27_
            nw10_[int(9)] = d_51_v27_
            nw10_[int(10)] = d_51_v27_
            nw10_[int(11)] = d_51_v27_
            nw10_[int(12)] = d_51_v27_
            nw10_[int(13)] = d_51_v27_
            nw10_[int(14)] = D1_DC3(d_48_v26_)
            nw10_[int(15)] = d_51_v27_
            nw10_[int(16)] = d_51_v27_
            def iife7_(_pat_let0_0):
                def iife8_(d_55_dt__update__tmp_h0_):
                    def iife9_(_pat_let1_0):
                        def iife10_(d_56_dt__update_hcf7_h0_):
                            return D1_DC3(d_56_dt__update_hcf7_h0_)
                        return iife10_(_pat_let1_0)
                    return iife9_(pat_let_tv0_)
                return iife8_(_pat_let0_0)
            nw10_[int(17)] = iife7_(d_51_v27_)
            nw10_[int(18)] = d_51_v27_
            nw10_[int(19)] = d_51_v27_
            nw10_[int(20)] = D1_DC3(d_48_v26_)
            nw10_[int(21)] = d_51_v27_
            nw10_[int(22)] = D1_DC3(d_48_v26_)
            nw10_[int(23)] = d_51_v27_
            d_54_v30_ = nw10_
            d_57_v31_: C2
            nw11_ = C2()
            nw11_.ctor__(d_45_v25_, d_54_v30_)
            d_57_v31_ = nw11_
            (globalState).f1 = d_53_v29_
            d_58_v32_: _dafny.MultiSet
            d_58_v32_ = _dafny.MultiSet([d_53_v29_, 536, d_53_v29_, default__.fm8(d_53_v29_, d_17_v0_, d_17_v0_, globalState)])
            d_59_v33_: _dafny.Array
            def lambda7_(d_60_v29_, d_61_v0_):
                def lambda8_(d_62_i4_):
                    return (d_62_i4_) * ((D14_DC40(d_60_v29_, d_61_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywpyyikps")))).cf73)

                return lambda8_

            init4_ = lambda7_(d_53_v29_, d_17_v0_)
            nw12_ = _dafny.Array(None, 19)
            for i0_4_ in range(nw12_.length(0)):
                nw12_[i0_4_] = init4_(i0_4_)
            d_59_v33_ = nw12_
            index8_ = default__.safeIndex(824, (d_59_v33_).length(0))
            (d_59_v33_)[index8_] = d_53_v29_
            index9_ = default__.safeIndex(824, (d_59_v33_).length(0))
            rhs9_ = (d_53_v29_) + (d_53_v29_)
            rhs10_ = ((d_58_v32_) | (d_58_v32_)).intersection((d_58_v32_) | (default__.fm33(globalState)))
            rhs11_ = d_53_v29_
            lhs6_ = d_59_v33_
            lhs7_ = default__.safeIndex(824, (d_59_v33_).length(0))
            d_53_v29_ = rhs9_
            d_58_v32_ = rhs10_
            lhs6_[lhs7_] = rhs11_
            d_63_v34_: _dafny.Seq
            d_63_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vejr"))
            d_63_v34_ = d_63_v34_
        d_64_v35_: _dafny.Array
        nw13_ = _dafny.Array(int(0), 28)
        d_64_v35_ = nw13_
        d_65_v36_: int
        d_65_v36_ = -214
        index10_ = default__.safeIndex(91, (d_64_v35_).length(0))
        (d_64_v35_)[index10_] = d_65_v36_
        index11_ = default__.safeIndex(91, (d_64_v35_).length(0))
        (d_64_v35_)[index11_] = d_65_v36_
        index12_ = default__.safeIndex(91, (d_64_v35_).length(0))
        (d_64_v35_)[index12_] = (d_65_v36_) + (d_65_v36_)
        d_65_v36_ = (d_64_v35_)[default__.safeIndex(91, (d_64_v35_).length(0))]
        index13_ = default__.safeIndex(91, (d_64_v35_).length(0))
        (d_64_v35_)[index13_] = d_65_v36_
        d_66_v37_: _dafny.Seq
        d_66_v37_ = _dafny.SeqWithoutIsStrInference([d_65_v36_])
        d_67_v38_: str
        d_67_v38_ = _dafny.CodePoint('x')
        d_68_v39_: T1
        nw14_ = C0()
        nw14_.ctor__(d_17_v0_, d_17_v0_, d_67_v38_, _dafny.SeqWithoutIsStrInference([d_67_v38_ for d_69_i5_ in range(default__.abs(407))]))
        d_68_v39_ = nw14_
        d_70_v40_: _dafny.MultiSet
        d_70_v40_ = _dafny.MultiSet([d_68_v39_, d_68_v39_])
        d_71_v41_: _dafny.Seq
        d_71_v41_ = _dafny.SeqWithoutIsStrInference([d_17_v0_])
        d_72_v42_: _dafny.Set
        d_72_v42_ = _dafny.Set({(d_70_v40_).cardinality, 733, len((d_71_v41_).set(default__.safeIndex(688, len(d_71_v41_)), d_17_v0_))})
        d_73_v43_: _dafny.Map
        d_73_v43_ = _dafny.Map({False: (d_68_v39_).f7})
        d_74_v44_: T2
        nw15_ = C1()
        nw15_.ctor__(len(d_72_v42_), (d_68_v39_).f6, (d_68_v39_).f6, ((d_73_v43_)[(d_68_v39_).f7] if ((d_68_v39_).f7) in (d_73_v43_) else d_17_v0_), d_67_v38_, (d_68_v39_).f5)
        d_74_v44_ = nw15_
        d_75_v45_: _dafny.Map
        d_75_v45_ = _dafny.Map({D7_DC21(d_66_v37_, d_74_v44_): d_65_v36_})
        d_76_v46_: D7
        d_76_v46_ = D7_DC21(d_66_v37_, d_74_v44_)
        pat_let_tv1_ = d_66_v37_
        pat_let_tv2_ = d_74_v44_
        pat_let_tv3_ = d_66_v37_
        pat_let_tv4_ = d_65_v36_
        def iife11_(_pat_let2_0):
            def iife12_(d_77_dt__update__tmp_h1_):
                def iife13_(_pat_let3_0):
                    def iife14_(d_78_dt__update_hcf39_h0_):
                        return D7_DC21(d_78_dt__update_hcf39_h0_, (d_77_dt__update__tmp_h1_).cf40)
                    return iife14_(_pat_let3_0)
                return iife13_((pat_let_tv1_).set(default__.safeIndex((pat_let_tv2_).f8, len(pat_let_tv3_)), pat_let_tv4_))
            return iife12_(_pat_let2_0)
        d_75_v45_ = (d_75_v45_).set(iife11_(d_76_v46_), 957)

    @staticmethod
    def Main(noArgsParameter__):
        d_79_v0_: bool
        d_79_v0_ = True
        d_80_v1_: _dafny.Array
        nw16_ = _dafny.Array(int(0), 12)
        d_80_v1_ = nw16_
        d_81_v2_: _dafny.Map
        d_81_v2_ = _dafny.Map({d_79_v0_: d_80_v1_})
        d_82_globalState_: GlobalState
        nw17_ = GlobalState()
        nw17_.ctor__(False, 130, False, ((d_81_v2_)[d_79_v0_] if (d_79_v0_) in (d_81_v2_) else d_80_v1_))
        d_82_globalState_ = nw17_
        d_83_v3_: int
        d_83_v3_ = 287
        d_79_v0_ = ((0) - (d_83_v3_)) < (398)
        d_84_v4_: _dafny.Seq
        d_84_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikmk"))
        d_85_v5_: _dafny.MultiSet
        d_85_v5_ = _dafny.MultiSet([default__.fm0(d_79_v0_, len(d_84_v4_), d_83_v3_, d_79_v0_, d_82_globalState_)])
        d_85_v5_ = d_85_v5_
        d_79_v0_ = d_79_v0_
        default__.m0(d_82_globalState_)
        d_86_v6_: _dafny.Set
        d_86_v6_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmln"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkkejxop"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnqpsmb")), d_84_v4_})
        d_87_v7_: _dafny.Seq
        d_87_v7_ = _dafny.SeqWithoutIsStrInference([d_84_v4_])
        def iife15_():
            coll7_ = _dafny.Set()
            compr_7_: _dafny.Seq
            for compr_7_ in ((d_87_v7_ if d_79_v0_ else d_87_v7_)).Elements:
                d_88_v8_: _dafny.Seq = compr_7_
                if (d_88_v8_) in ((d_87_v7_ if d_79_v0_ else d_87_v7_)):
                    coll7_ = coll7_.union(_dafny.Set([d_88_v8_]))
            return _dafny.Set(coll7_)
        d_86_v6_ = iife15_()
        
        d_89_i0_: int
        d_89_i0_ = 0
        with _dafny.label("0"):
            while d_79_v0_:
                with _dafny.c_label("0"):
                    if (d_89_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_89_i0_ = (d_89_i0_) + (1)
                    d_90_v9_: _dafny.Map
                    d_90_v9_ = _dafny.Map({349: d_84_v4_})
                    d_91_v10_: str
                    d_91_v10_ = _dafny.CodePoint('h')
                    d_92_v11_: _dafny.Map
                    d_92_v11_ = _dafny.Map({d_79_v0_: default__.fm7(d_91_v10_, d_83_v3_, True, d_82_globalState_)})
                    d_93_v12_: C3
                    nw18_ = C3()
                    nw18_.ctor__(False, d_79_v0_, len(d_90_v9_), d_79_v0_, True, (len(d_92_v11_)) <= (d_83_v3_), d_91_v10_, (d_84_v4_) + (d_84_v4_))
                    d_93_v12_ = nw18_
                    d_94_v13_: _dafny.Seq
                    d_94_v13_ = _dafny.SeqWithoutIsStrInference([d_93_v12_.f12, d_93_v12_.f12])
                    d_95_v14_: _dafny.Array
                    nw19_ = _dafny.Array(None, 29)
                    nw19_[int(0)] = d_79_v0_
                    nw19_[int(1)] = (d_93_v12_).f13
                    nw19_[int(2)] = (d_93_v12_).f13
                    nw19_[int(3)] = d_93_v12_.f12
                    nw19_[int(4)] = (d_93_v12_).f13
                    nw19_[int(5)] = d_79_v0_
                    nw19_[int(6)] = (d_93_v12_).f13
                    nw19_[int(7)] = d_79_v0_
                    nw19_[int(8)] = (d_93_v12_).f13
                    nw19_[int(9)] = (d_93_v12_).f13
                    nw19_[int(10)] = default__.fm3(d_82_globalState_)
                    nw19_[int(11)] = (d_93_v12_).f13
                    nw19_[int(12)] = (d_94_v13_)[default__.safeIndex(d_83_v3_, len(d_94_v13_))]
                    nw19_[int(13)] = not(d_79_v0_)
                    nw19_[int(14)] = d_79_v0_
                    nw19_[int(15)] = d_79_v0_
                    nw19_[int(16)] = (d_93_v12_).f13
                    nw19_[int(17)] = not((d_93_v12_).f13)
                    nw19_[int(18)] = d_93_v12_.f12
                    nw19_[int(19)] = d_79_v0_
                    nw19_[int(20)] = d_93_v12_.f12
                    nw19_[int(21)] = d_79_v0_
                    nw19_[int(22)] = (d_93_v12_).f13
                    nw19_[int(23)] = d_93_v12_.f12
                    nw19_[int(24)] = d_93_v12_.f12
                    nw19_[int(25)] = not(d_93_v12_.f12)
                    nw19_[int(26)] = d_79_v0_
                    nw19_[int(27)] = d_93_v12_.f12
                    nw19_[int(28)] = not(d_93_v12_.f12)
                    d_95_v14_ = nw19_
                    d_96_v15_: _dafny.Map
                    d_96_v15_ = _dafny.Map({d_95_v14_: ((d_92_v11_)[(d_93_v12_).f13] if ((d_93_v12_).f13) in (d_92_v11_) else (d_93_v12_).f13)})
                    d_96_v15_ = (d_96_v15_).set(d_95_v14_, (d_93_v12_).f13)
                    d_97_v16_: _dafny.Array
                    def lambda9_(d_98_v3_, d_99_v13_):
                        def lambda10_(d_100_i1_):
                            return (_dafny.MultiSet([d_98_v3_, len(d_99_v13_)])) - (_dafny.MultiSet([d_98_v3_, d_98_v3_, d_98_v3_]))

                        return lambda10_

                    init5_ = lambda9_(d_83_v3_, d_94_v13_)
                    nw20_ = _dafny.Array(None, 19)
                    for i0_5_ in range(nw20_.length(0)):
                        nw20_[i0_5_] = init5_(i0_5_)
                    d_97_v16_ = nw20_
                    d_97_v16_ = d_97_v16_
                    d_101_v17_: _dafny.Map
                    d_101_v17_ = _dafny.Map({(d_93_v12_).f13: d_94_v13_})
                    d_102_v18_: _dafny.Array
                    nw21_ = _dafny.Array(None, 29)
                    nw21_[int(0)] = d_94_v13_
                    nw21_[int(1)] = d_94_v13_
                    nw21_[int(2)] = d_94_v13_
                    nw21_[int(3)] = d_94_v13_
                    nw21_[int(4)] = ((d_101_v17_)[False] if (False) in (d_101_v17_) else d_94_v13_)
                    nw21_[int(5)] = d_94_v13_
                    nw21_[int(6)] = d_94_v13_
                    nw21_[int(7)] = d_94_v13_
                    nw21_[int(8)] = d_94_v13_
                    nw21_[int(9)] = d_94_v13_
                    nw21_[int(10)] = _dafny.SeqWithoutIsStrInference([True, d_93_v12_.f12])
                    nw21_[int(11)] = d_94_v13_
                    nw21_[int(12)] = d_94_v13_
                    nw21_[int(13)] = d_94_v13_
                    nw21_[int(14)] = d_94_v13_
                    nw21_[int(15)] = d_94_v13_
                    nw21_[int(16)] = d_94_v13_
                    nw21_[int(17)] = d_94_v13_
                    nw21_[int(18)] = d_94_v13_
                    nw21_[int(19)] = d_94_v13_
                    nw21_[int(20)] = default__.fm18(len((_dafny.SeqWithoutIsStrInference([d_91_v10_ for d_103_i2_ in range(default__.abs(640))])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqt"))), len(_dafny.SeqWithoutIsStrInference([d_91_v10_ for d_104_i2_ in range(default__.abs(640))]))), d_91_v10_)), d_93_v12_.f12, d_82_globalState_)
                    nw21_[int(21)] = d_94_v13_
                    nw21_[int(22)] = d_94_v13_
                    nw21_[int(23)] = d_94_v13_
                    nw21_[int(24)] = d_94_v13_
                    nw21_[int(25)] = d_94_v13_
                    nw21_[int(26)] = d_94_v13_
                    nw21_[int(27)] = d_94_v13_
                    nw21_[int(28)] = d_94_v13_
                    d_102_v18_ = nw21_
                    d_105_v19_: D1
                    d_105_v19_ = D1_DC3(d_102_v18_)
                    d_106_v20_: _dafny.Map
                    d_106_v20_ = _dafny.Map({d_83_v3_: d_105_v19_})
                    d_107_v21_: D2
                    d_107_v21_ = D2_DC6(d_106_v20_)
                    d_108_v22_: _dafny.Set
                    d_108_v22_ = _dafny.Set({d_107_v21_})
                    d_109_v23_: D13
                    d_109_v23_ = D13_DC35(d_108_v22_)
                    d_108_v22_ = (d_109_v23_).cf66
                    pass
            pass
        d_79_v0_ = (default__.fm3(d_82_globalState_) if True else False)
        d_79_v0_ = d_79_v0_
        d_110_v24_: _dafny.MultiSet
        d_110_v24_ = _dafny.MultiSet([d_83_v3_])
        d_110_v24_ = d_110_v24_
        d_83_v3_ = d_83_v3_
        d_111_v25_: _dafny.Seq
        d_111_v25_ = _dafny.SeqWithoutIsStrInference([372, d_83_v3_])
        d_112_v26_: _dafny.Seq
        d_112_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_111_v25_)])
        d_113_v27_: str
        d_113_v27_ = _dafny.CodePoint('b')
        d_114_v28_: C1
        nw22_ = C1()
        nw22_.ctor__(d_83_v3_, d_79_v0_, d_79_v0_, d_79_v0_, d_113_v27_, d_84_v4_)
        d_114_v28_ = nw22_
        d_115_v29_: _dafny.Array
        nw23_ = _dafny.Array(False, 2)
        d_115_v29_ = nw23_
        d_116_v30_: C5
        nw24_ = C5()
        nw24_.ctor__(d_115_v29_, d_80_v1_, not(d_79_v0_), True, (d_84_v4_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_117_i3_ in range(default__.abs(-922))])), len(d_84_v4_))], d_84_v4_, d_83_v3_, d_79_v0_)
        d_116_v30_ = nw24_
        d_118_v31_: _dafny.Map
        d_118_v31_ = _dafny.Map({d_79_v0_: d_116_v30_})
        d_119_v32_: T3
        nw25_ = C4()
        nw25_.ctor__(d_83_v3_, d_79_v0_, default__.fm27(d_83_v3_, d_79_v0_, default__.fm27(d_83_v3_, not(not(default__.fm27(d_83_v3_, d_79_v0_, d_79_v0_, d_82_globalState_))), d_79_v0_, d_82_globalState_), d_82_globalState_), d_79_v0_, _dafny.CodePoint('m'), d_84_v4_)
        d_119_v32_ = nw25_
        d_120_v33_: _dafny.Map
        d_120_v33_ = _dafny.Map({d_119_v32_: (d_119_v32_).f7})
        d_121_v34_: D0
        d_121_v34_ = D0_DC2(d_83_v3_, d_79_v0_, d_111_v25_)
        d_122_v35_: _dafny.Array
        nw26_ = _dafny.Array(None, 28)
        nw26_[int(0)] = _dafny.MultiSet(d_111_v25_)
        nw26_[int(1)] = (d_110_v24_).intersection(_dafny.MultiSet([d_83_v3_]))
        nw26_[int(2)] = d_110_v24_
        nw26_[int(3)] = ((d_112_v26_)[default__.safeIndex(d_83_v3_, len(d_112_v26_))]).intersection(d_110_v24_)
        nw26_[int(4)] = (d_110_v24_).set(d_83_v3_, default__.abs(len((_dafny.Map({d_114_v28_: ((d_110_v24_).set(d_83_v3_, default__.abs(d_83_v3_))).cardinality})).set(d_114_v28_, d_83_v3_))))
        nw26_[int(5)] = d_110_v24_
        nw26_[int(6)] = d_110_v24_
        nw26_[int(7)] = (d_110_v24_) - (d_110_v24_)
        nw26_[int(8)] = d_110_v24_
        nw26_[int(9)] = d_110_v24_
        nw26_[int(10)] = d_110_v24_
        nw26_[int(11)] = d_110_v24_
        nw26_[int(12)] = d_110_v24_
        nw26_[int(13)] = d_110_v24_
        nw26_[int(14)] = d_110_v24_
        nw26_[int(15)] = d_110_v24_
        nw26_[int(16)] = d_110_v24_
        nw26_[int(17)] = default__.fm33(d_82_globalState_)
        nw26_[int(18)] = d_110_v24_
        nw26_[int(19)] = (_dafny.MultiSet([len(d_118_v31_), d_83_v3_])) - (d_110_v24_)
        nw26_[int(20)] = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(0) - (d_83_v3_) for d_123_i4_ in range(default__.abs(608))])) + (_dafny.SeqWithoutIsStrInference([d_83_v3_, len(d_111_v25_), d_83_v3_])))
        nw26_[int(21)] = d_110_v24_
        nw26_[int(22)] = d_110_v24_
        nw26_[int(23)] = (_dafny.MultiSet([len(d_84_v4_)])) | (d_110_v24_)
        nw26_[int(24)] = d_110_v24_
        nw26_[int(25)] = d_110_v24_
        nw26_[int(26)] = (_dafny.MultiSet([d_83_v3_])).intersection(d_110_v24_)
        nw26_[int(27)] = _dafny.MultiSet([len(d_120_v33_), len(_dafny.SeqWithoutIsStrInference([(d_119_v32_).f6, (d_119_v32_).f7, (d_119_v32_).f7, (d_121_v34_).cf5, d_79_v0_])), -415, d_83_v3_, d_83_v3_])
        d_122_v35_ = nw26_
        index14_ = default__.safeIndex(124, (d_122_v35_).length(0))
        (d_122_v35_)[index14_] = d_110_v24_
        index15_ = default__.safeIndex(124, (d_122_v35_).length(0))
        (d_122_v35_)[index15_] = (d_110_v24_).intersection(d_110_v24_)
        source1_ = default__.fm26((d_119_v32_).f7, 147, d_119_v32_.f4, d_82_globalState_)
        if source1_.is_DC4:
            d_124___mcc_h0_ = source1_.cf8
            d_125___mcc_h1_ = source1_.cf9
            d_126___mcc_h2_ = source1_.cf10
            d_127___mcc_h3_ = source1_.cf11
            d_128___mcc_h4_ = source1_.cf12
            d_129_cf12_ = d_128___mcc_h4_
            d_130_cf11_ = d_127___mcc_h3_
            d_131_cf10_ = d_126___mcc_h2_
            d_132_cf9_ = d_125___mcc_h1_
            d_133_cf8_ = d_124___mcc_h0_
            d_134_v36_: _dafny.MultiSet
            d_134_v36_ = _dafny.MultiSet([d_79_v0_, (d_119_v32_).f6])
            d_135_v37_: int
            d_136_v38_: int
            d_137_v39_: _dafny.Array
            out0_: int
            out1_: int
            out2_: _dafny.Array
            out0_, out1_, out2_ = (d_116_v30_).m1((d_119_v32_).f6, (d_119_v32_).f6, (d_134_v36_).set((d_119_v32_).f6, default__.abs(71)), (d_119_v32_).f6, d_82_globalState_)
            d_135_v37_ = out0_
            d_136_v38_ = out1_
            d_137_v39_ = out2_
            d_138_v40_: _dafny.Seq
            d_138_v40_ = _dafny.SeqWithoutIsStrInference([(d_119_v32_).f9])
            d_139_v41_: _dafny.Map
            d_139_v41_ = _dafny.Map({d_79_v0_: len(d_138_v40_)})
            d_140_v42_: D10
            d_140_v42_ = D10_DC28(d_129_cf12_, len(d_139_v41_), (d_119_v32_).f9, (d_119_v32_).f9, d_114_v28_)
            d_141_v43_: T2
            nw27_ = C1()
            nw27_.ctor__((d_119_v32_).f8, (d_119_v32_).f6, (d_119_v32_).f6, (d_140_v42_).cf50, _dafny.CodePoint('t'), (d_119_v32_).f5)
            d_141_v43_ = nw27_
            d_142_v44_: _dafny.Map
            d_142_v44_ = _dafny.Map({d_141_v43_: d_129_cf12_})
            d_143_v45_: D5
            d_143_v45_ = D5_DC15((d_116_v30_).f10, (d_119_v32_).f9, d_139_v41_, len(d_142_v44_))
            d_144_v46_: T2
            nw28_ = C3()
            nw28_.ctor__(not((d_119_v32_).f7), ((d_119_v32_).f7) and (d_79_v0_), d_132_cf9_, (d_143_v45_).cf27, False, False, d_133_cf8_, ((d_119_v32_).f5) + (_dafny.SeqWithoutIsStrInference([d_119_v32_.f4 for d_145_i5_ in range(default__.abs(519))])))
            d_144_v46_ = nw28_
            d_144_v46_ = d_141_v43_
            d_129_cf12_ = ((d_141_v43_).f5) <= ((d_141_v43_).f5)
            (d_144_v46_).f4 = d_131_cf10_
        elif source1_.is_DC5:
            d_146___mcc_h5_ = source1_.cf13
            d_147_cf13_ = d_146___mcc_h5_
            (d_82_globalState_).f1 = d_83_v3_
            d_148_v47_: _dafny.Map
            d_148_v47_ = _dafny.Map({(d_83_v3_) - (d_147_cf13_): (d_119_v32_).f6})
            d_148_v47_ = d_148_v47_
            d_149_v48_: _dafny.Seq
            d_149_v48_ = _dafny.SeqWithoutIsStrInference([False])
            d_150_v49_: _dafny.Map
            d_150_v49_ = _dafny.Map({default__.safeDivisionInt(d_147_cf13_, (d_119_v32_).f8): d_149_v48_})
            d_150_v49_ = d_150_v49_
            d_151_v50_: D2
            d_151_v50_ = D2_DC7((d_119_v32_).f9)
            d_79_v0_ = (d_114_v28_).fm19((d_119_v32_).f6, d_79_v0_, _dafny.Set({d_151_v50_, d_151_v50_, d_151_v50_, d_151_v50_, d_151_v50_}), d_82_globalState_)
        elif True:
            d_152___mcc_h6_ = source1_.cf7
            d_153_cf7_ = d_152___mcc_h6_
            d_154_v51_: _dafny.Array
            nw29_ = _dafny.Array(_dafny.Set({}), 28)
            d_154_v51_ = nw29_
            d_155_v52_: D14
            d_155_v52_ = D14_DC39(d_154_v51_)
            d_156_v53_: _dafny.Map
            d_156_v53_ = _dafny.Map({(d_155_v52_).cf72: ((d_119_v32_).f5) <= (d_84_v4_)})
            if not(not(((d_156_v53_)[d_154_v51_] if (d_154_v51_) in (d_156_v53_) else (d_119_v32_).f9))):
                d_157_v54_: T2
                nw30_ = C3()
                nw30_.ctor__((d_119_v32_).f9, (d_119_v32_).f7, (d_119_v32_).f8, (d_119_v32_).f6, (d_119_v32_).f6, d_79_v0_, d_119_v32_.f4, _dafny.SeqWithoutIsStrInference([d_119_v32_.f4 for d_158_i6_ in range(default__.abs(424))]))
                d_157_v54_ = nw30_
                d_159_v55_: D7
                d_159_v55_ = D7_DC21(d_111_v25_, d_157_v54_)
                d_79_v0_ = (default__.safeModuloInt(len((d_159_v55_).cf39), d_83_v3_)) <= ((d_119_v32_).f8)
                d_160_v56_: _dafny.Set
                d_160_v56_ = _dafny.Set({353})
                d_161_v57_: _dafny.Seq
                d_161_v57_ = _dafny.SeqWithoutIsStrInference([d_160_v56_])
                d_162_v58_: _dafny.Map
                d_162_v58_ = _dafny.Map({(d_160_v56_ if (d_119_v32_).f7 else d_160_v56_): len(d_161_v57_)})
                d_162_v58_ = d_162_v58_
                d_79_v0_ = (d_119_v32_).f7
                d_163_v59_: _dafny.Map
                d_163_v59_ = _dafny.Map({(d_119_v32_).f9: (d_119_v32_).f9})
                d_164_v60_: _dafny.Seq
                d_164_v60_ = _dafny.SeqWithoutIsStrInference([(d_119_v32_).f7])
                d_163_v59_ = (d_163_v59_).set(not(not((_dafny.SeqWithoutIsStrInference([(d_157_v54_).f8, (0) - (len(d_164_v60_))])) == (d_111_v25_))), (d_157_v54_).f7)
                d_165_v61_: D5
                d_165_v61_ = D5_DC14((d_116_v30_).f11)
                d_166_v62_: _dafny.Set
                d_166_v62_ = _dafny.Set({D5_DC14(d_80_v1_), d_165_v61_, d_165_v61_, D5_DC14(d_80_v1_), d_165_v61_})
                d_167_v63_: _dafny.Map
                d_167_v63_ = _dafny.Map({d_113_v27_: ((d_119_v32_).f8) * (len(d_166_v62_))})
                d_167_v63_ = (d_167_v63_).set(d_157_v54_.f4, d_83_v3_)
            elif True:
                index16_ = default__.safeIndex(265, ((d_116_v30_).f10).length(0))
                ((d_116_v30_).f10)[index16_] = (d_119_v32_).f7
                index17_ = default__.safeIndex(265, ((d_116_v30_).f10).length(0))
                ((d_116_v30_).f10)[index17_] = not(not((d_119_v32_).f7))
                d_168_v64_: _dafny.Map
                d_168_v64_ = _dafny.Map({(d_119_v32_).f8: (d_119_v32_).f7})
                d_169_v65_: T1
                nw31_ = C0()
                nw31_.ctor__(((d_168_v64_)[(d_119_v32_).f8] if ((d_119_v32_).f8) in (d_168_v64_) else not((d_119_v32_).f6)), ((d_116_v30_).f10)[default__.safeIndex(265, ((d_116_v30_).f10).length(0))], d_119_v32_.f4, (d_119_v32_).f5)
                d_169_v65_ = nw31_
                d_169_v65_ = d_169_v65_
                d_170_v66_: _dafny.Map
                d_170_v66_ = _dafny.Map({d_116_v30_: (d_119_v32_).f8})
                d_171_v67_: _dafny.Seq
                d_171_v67_ = _dafny.SeqWithoutIsStrInference([d_116_v30_])
                (d_82_globalState_).f1 = (d_111_v25_)[default__.safeIndex((0) - ((((d_170_v66_)[(d_171_v67_)[default__.safeIndex(d_83_v3_, len(d_171_v67_))]] if ((d_171_v67_)[default__.safeIndex(d_83_v3_, len(d_171_v67_))]) in (d_170_v66_) else d_83_v3_)) - ((d_119_v32_).f8)), len(d_111_v25_))]
                d_172_v68_: D12
                d_172_v68_ = D12_DC34((d_119_v32_).f6)
                d_173_v69_: C4
                nw32_ = C4()
                nw32_.ctor__((d_119_v32_).f8, (d_119_v32_).f7, False, (d_172_v68_).cf65, d_119_v32_.f4, ((d_119_v32_).f5) + (_dafny.SeqWithoutIsStrInference([d_119_v32_.f4 for d_174_i7_ in range(default__.abs(32))])))
                d_173_v69_ = nw32_
                d_175_v70_: _dafny.Array
                def lambda11_(d_176_v32_):
                    def lambda12_(d_177_i8_):
                        return _dafny.SeqWithoutIsStrInference([d_176_v32_.f4 for d_178_i9_ in range(default__.abs(971))])

                    return lambda12_

                init6_ = lambda11_(d_119_v32_)
                nw33_ = _dafny.Array(None, 20)
                for i0_6_ in range(nw33_.length(0)):
                    nw33_[i0_6_] = init6_(i0_6_)
                d_175_v70_ = nw33_
                index18_ = default__.safeIndex(134, (d_175_v70_).length(0))
                (d_175_v70_)[index18_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlgmycrna"))
                index19_ = default__.safeIndex(134, (d_175_v70_).length(0))
                (d_175_v70_)[index19_] = ((d_119_v32_).f5) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_179_i10_ in range(default__.abs(211))]))
            d_180_v71_: _dafny.Seq
            d_180_v71_ = _dafny.SeqWithoutIsStrInference([(d_119_v32_).f7])
            d_181_v72_: _dafny.Set
            d_181_v72_ = _dafny.Set({d_180_v71_})
            d_182_v73_: C4
            nw34_ = C4()
            nw34_.ctor__(len(d_181_v72_), d_79_v0_, (d_119_v32_).f6, not(not(False)), _dafny.CodePoint('g'), ((d_119_v32_).f5) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lctrtkorn"))))
            d_182_v73_ = nw34_
            d_182_v73_ = d_182_v73_
            d_83_v3_ = (d_119_v32_).f8
            d_183_v74_: _dafny.Map
            d_183_v74_ = _dafny.Map({162: d_83_v3_})
            d_184_v75_: _dafny.MultiSet
            d_184_v75_ = _dafny.MultiSet([(d_119_v32_).f9, True, default__.fm3(d_82_globalState_)])
            d_185_v76_: int
            d_186_v77_: int
            d_187_v78_: _dafny.Array
            out3_: int
            out4_: int
            out5_: _dafny.Array
            out3_, out4_, out5_ = (d_182_v73_).m1(d_79_v0_, not((len((d_119_v32_).f5)) not in (d_183_v74_)), d_184_v75_, (d_119_v32_).f6, d_82_globalState_)
            d_185_v76_ = out3_
            d_186_v77_ = out4_
            d_187_v78_ = out5_
        d_188_v79_: _dafny.Set
        d_188_v79_ = _dafny.Set({(d_119_v32_).f9})
        d_189_v80_: _dafny.Seq
        d_189_v80_ = _dafny.SeqWithoutIsStrInference([(d_188_v79_).issubset(d_188_v79_), (d_119_v32_).f9, (d_119_v32_).f9])
        d_189_v80_ = (d_189_v80_) + (d_189_v80_)
        d_190_v81_: C4
        nw35_ = C4()
        nw35_.ctor__(d_83_v3_, d_79_v0_, d_79_v0_, True, d_113_v27_, d_84_v4_)
        d_190_v81_ = nw35_
        d_191_v82_: _dafny.Set
        d_191_v82_ = _dafny.Set({d_190_v81_, (d_190_v81_ if default__.fm27((0) - (-56), True, not((d_119_v32_).f7), d_82_globalState_) else d_190_v81_)})
        d_192_v83_: T0
        nw36_ = C4()
        nw36_.ctor__(d_83_v3_, (d_119_v32_).f7, (d_119_v32_).f9, d_79_v0_, d_119_v32_.f4, d_84_v4_)
        d_192_v83_ = nw36_
        d_193_v84_: _dafny.Map
        d_193_v84_ = _dafny.Map({d_79_v0_: (d_79_v0_) or ((d_119_v32_).f9)})
        d_194_v85_: _dafny.Array
        nw37_ = _dafny.Array(D11.default()(), 5)
        d_194_v85_ = nw37_
        rhs12_ = d_191_v82_
        rhs13_ = (d_119_v32_).f7
        rhs14_ = d_192_v83_
        rhs15_ = d_193_v84_
        rhs16_ = d_194_v85_
        d_191_v82_ = rhs12_
        d_79_v0_ = rhs13_
        d_192_v83_ = rhs14_
        d_193_v84_ = rhs15_
        d_194_v85_ = rhs16_
        d_195_v86_: _dafny.Map
        d_195_v86_ = _dafny.Map({(d_119_v32_).f8: True})
        if ((d_195_v86_)[default__.safeModuloInt((d_119_v32_).f8, len((d_119_v32_).f5))] if (default__.safeModuloInt((d_119_v32_).f8, len((d_119_v32_).f5))) in (d_195_v86_) else (d_189_v80_)[default__.safeIndex(len(d_84_v4_), len(d_189_v80_))]):
            if d_79_v0_:
                d_196_v87_: C0
                nw38_ = C0()
                nw38_.ctor__(((d_119_v32_).f7 if (d_119_v32_).f9 else (d_119_v32_).f9), (d_119_v32_).f6, ((d_192_v83_).f5)[default__.safeIndex(d_83_v3_, len((d_192_v83_).f5))], (d_119_v32_).f5)
                d_196_v87_ = nw38_
                d_196_v87_ = d_196_v87_
                (d_119_v32_).f4 = d_192_v83_.f4
                d_197_v88_: _dafny.Array
                nw39_ = _dafny.Array(D7.default()(), 22)
                d_197_v88_ = nw39_
                d_198_v89_: _dafny.Map
                d_198_v89_ = _dafny.Map({(d_119_v32_).f6: d_197_v88_})
                d_199_v90_: _dafny.Map
                d_199_v90_ = _dafny.Map({(d_198_v89_) | (d_198_v89_): ((d_114_v28_).fm1((d_119_v32_).f7, d_193_v84_, (d_119_v32_).f6, d_82_globalState_)) < ((d_119_v32_).f8)})
                d_200_v91_: _dafny.Seq
                d_200_v91_ = _dafny.SeqWithoutIsStrInference([(d_116_v30_).f10, (d_116_v30_).f10])
                d_201_v92_: _dafny.Seq
                d_201_v92_ = _dafny.SeqWithoutIsStrInference([d_200_v91_, d_200_v91_])
                d_199_v90_ = (d_199_v90_).set(d_198_v89_, ((d_122_v35_)[default__.safeIndex(124, (d_122_v35_).length(0))]).ispropersubset(_dafny.MultiSet([(d_119_v32_).f8, len((d_201_v92_)[default__.safeIndex((d_119_v32_).f8, len(d_201_v92_))]), d_83_v3_])))
                d_202_v93_: _dafny.MultiSet
                d_202_v93_ = _dafny.MultiSet([(d_119_v32_).f6])
                d_203_v94_: int
                d_204_v95_: int
                d_205_v96_: _dafny.Array
                out6_: int
                out7_: int
                out8_: _dafny.Array
                out6_, out7_, out8_ = (d_119_v32_).m1((d_119_v32_).f6, d_79_v0_, d_202_v93_, not((d_119_v32_).f9), d_82_globalState_)
                d_203_v94_ = out6_
                d_204_v95_ = out7_
                d_205_v96_ = out8_
                d_206_v97_: D0
                d_206_v97_ = D0_DC1(False, d_203_v94_, (d_119_v32_).f7, d_83_v3_)
                d_79_v0_ = not((d_206_v97_).cf2)
            elif True:
                d_207_v98_: _dafny.Seq
                d_207_v98_ = _dafny.SeqWithoutIsStrInference([(default__.fm15(d_82_globalState_)) | (_dafny.Map({(d_119_v32_).f8: d_83_v3_}))])
                d_207_v98_ = d_207_v98_
                d_208_v99_: D12
                d_208_v99_ = D12_DC34(False)
                d_79_v0_ = (d_208_v99_).cf65
                d_79_v0_ = (d_119_v32_).f9
                index20_ = default__.safeIndex(738, ((d_116_v30_).f11).length(0))
                ((d_116_v30_).f11)[index20_] = (85) * (d_83_v3_)
                index21_ = default__.safeIndex(738, ((d_116_v30_).f11).length(0))
                ((d_116_v30_).f11)[index21_] = 766
                d_209_v100_: C3
                nw40_ = C3()
                nw40_.ctor__(False, (d_119_v32_).f7, ((d_116_v30_).f11)[default__.safeIndex(738, ((d_116_v30_).f11).length(0))], True, d_79_v0_, not ((d_119_v32_).f7) or (True), d_119_v32_.f4, (d_192_v83_).f5)
                d_209_v100_ = nw40_
            d_210_v101_: D7
            d_210_v101_ = D7_DC20()
            source2_ = d_210_v101_
            if source2_.is_DC19:
                d_211___mcc_h7_ = source2_.cf34
                d_212___mcc_h8_ = source2_.cf35
                d_213___mcc_h9_ = source2_.cf36
                d_214___mcc_h10_ = source2_.cf37
                d_215___mcc_h11_ = source2_.cf38
                d_216_cf38_ = d_215___mcc_h11_
                d_217_cf37_ = d_214___mcc_h10_
                d_218_cf36_ = d_213___mcc_h9_
                d_219_cf35_ = d_212___mcc_h8_
                d_220_cf34_ = d_211___mcc_h7_
                index22_ = default__.safeIndex(147, ((d_116_v30_).f10).length(0))
                ((d_116_v30_).f10)[index22_] = d_79_v0_
                d_221_v102_: _dafny.Map
                d_221_v102_ = _dafny.Map({d_189_v80_: (d_119_v32_).f9})
                index23_ = default__.safeIndex(147, ((d_116_v30_).f10).length(0))
                ((d_116_v30_).f10)[index23_] = not (not ((d_119_v32_).f6) or (d_220_cf34_)) or (((d_119_v32_).f5) == (((d_192_v83_).f5).set(default__.safeIndex(len(d_221_v102_), len((d_192_v83_).f5)), d_192_v83_.f4)))
                d_190_v81_ = d_190_v81_
                d_222_v103_: _dafny.Map
                d_222_v103_ = _dafny.Map({(d_119_v32_).fm1((d_119_v32_).f6, d_193_v84_, (d_119_v32_).f6, d_82_globalState_): d_83_v3_})
                d_223_v104_: _dafny.MultiSet
                d_223_v104_ = _dafny.MultiSet([d_190_v81_])
                d_222_v103_ = (d_222_v103_).set((0) - ((d_119_v32_).f8), ((d_223_v104_)[d_190_v81_] if (d_190_v81_) in (d_223_v104_) else (d_119_v32_).f8))
                d_224_v105_: _dafny.Map
                d_224_v105_ = _dafny.Map({d_192_v83_: (d_119_v32_).f8})
                d_224_v105_ = (d_224_v105_).set(d_192_v83_, 732)
            elif source2_.is_DC20:
                d_225_v106_: _dafny.Array
                nw41_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
                d_225_v106_ = nw41_
                index24_ = default__.safeIndex(233, (d_225_v106_).length(0))
                (d_225_v106_)[index24_] = d_84_v4_
                index25_ = default__.safeIndex(233, (d_225_v106_).length(0))
                (d_225_v106_)[index25_] = (d_192_v83_).f5
                d_226_v107_: C5
                nw42_ = C5()
                nw42_.ctor__(d_115_v29_, (d_116_v30_).f11, (d_119_v32_).f7, (d_119_v32_).f9, d_192_v83_.f4, d_84_v4_, (d_83_v3_) - ((0) - ((d_119_v32_).f8)), (d_119_v32_).f6)
                d_226_v107_ = nw42_
                d_227_v108_: _dafny.Map
                d_227_v108_ = _dafny.Map({d_83_v3_: 956})
                d_228_v109_: _dafny.Map
                d_228_v109_ = _dafny.Map({(len((d_227_v108_).set(len(d_111_v25_), (d_119_v32_).f8))) > (904): (d_114_v28_).fm1(not((d_119_v32_).f7), default__.fm23(len(_dafny.SeqWithoutIsStrInference([827 for d_229_i11_ in range(default__.abs(991))])), (0) - ((d_119_v32_).f8), d_119_v32_.f4, (d_119_v32_).f9, d_82_globalState_), (d_119_v32_).f7, d_82_globalState_)})
                rhs17_ = (len(d_195_v86_) if not((d_119_v32_).f9) else len((_dafny.Set({False})).intersection(d_188_v79_)))
                rhs18_ = d_228_v109_
                lhs8_ = d_82_globalState_
                lhs8_.f1 = rhs17_
                d_228_v109_ = rhs18_
                (d_116_v30_).m4(len(d_195_v86_), d_195_v86_, d_82_globalState_)
            elif source2_.is_DC21:
                d_230___mcc_h12_ = source2_.cf39
                d_231___mcc_h13_ = source2_.cf40
                d_232_cf40_ = d_231___mcc_h13_
                d_233_cf39_ = d_230___mcc_h12_
                d_234_v110_: _dafny.Set
                d_234_v110_ = _dafny.Set({(d_119_v32_).f8})
                d_234_v110_ = _dafny.Set({d_83_v3_, len((d_119_v32_).f5), (d_119_v32_).f8, (d_232_cf40_).fm1((d_119_v32_).f9, d_193_v84_, False, d_82_globalState_)})
                d_235_v111_: C4
                nw43_ = C4()
                nw43_.ctor__((d_110_v24_).cardinality, (d_189_v80_) != (d_189_v80_), not((d_119_v32_).f9), default__.fm3(d_82_globalState_), d_232_cf40_.f4, ((D3_DC10((d_119_v32_).f5)).cf19).set(default__.safeIndex(d_83_v3_, len((D3_DC10((d_119_v32_).f5)).cf19)), d_113_v27_))
                d_235_v111_ = nw43_
                d_189_v80_ = (d_189_v80_) + (((d_189_v80_) + (d_189_v80_)).set(default__.safeIndex(len(d_188_v79_), len((d_189_v80_) + (d_189_v80_))), False))
                d_236_v112_: _dafny.Map
                out9_: _dafny.Map
                out9_ = (d_116_v30_).m2((d_119_v32_).f8, d_82_globalState_)
                d_236_v112_ = out9_
            elif True:
                d_237___mcc_h14_ = source2_.cf33
                d_238_cf33_ = d_237___mcc_h14_
                d_239_v113_: _dafny.Map
                d_239_v113_ = _dafny.Map({(d_119_v32_).f8: (0) - ((d_119_v32_).f8)})
                d_240_v114_: _dafny.Map
                d_240_v114_ = _dafny.Map({(d_119_v32_).f7: len(d_239_v113_)})
                d_241_v115_: D12
                d_241_v115_ = D12_DC31(d_193_v84_)
                d_242_v116_: C3
                nw44_ = C3()
                nw44_.ctor__((d_119_v32_).f6, (d_119_v32_).f7, ((d_240_v114_)[(d_119_v32_).f9] if ((d_119_v32_).f9) in (d_240_v114_) else (d_119_v32_).f8), not((d_119_v32_).f9), (d_119_v32_).f7, not((575) >= ((d_190_v81_).fm1(not((d_119_v32_).f6), (d_241_v115_).cf60, True, d_82_globalState_))), d_192_v83_.f4, (d_119_v32_).f5)
                d_242_v116_ = nw44_
                d_243_v117_: D15
                d_243_v117_ = D15_DC42(d_242_v116_)
                rhs19_ = d_83_v3_
                rhs20_ = (((d_119_v32_).f8) - ((0) - ((d_119_v32_).f8)) if (d_119_v32_).f9 else (0) - (((d_119_v32_).f8) - ((d_119_v32_).f8)))
                rhs21_ = (d_243_v117_).cf77
                lhs9_ = d_82_globalState_
                lhs10_ = d_82_globalState_
                lhs9_.f1 = rhs19_
                lhs10_.f1 = rhs20_
                d_242_v116_ = rhs21_
                d_244_v118_: _dafny.MultiSet
                d_244_v118_ = _dafny.MultiSet([(d_119_v32_).f9])
                d_245_v119_: int
                d_246_v120_: int
                d_247_v121_: _dafny.Array
                out10_: int
                out11_: int
                out12_: _dafny.Array
                out10_, out11_, out12_ = (d_242_v116_).m1((d_119_v32_).f6, not((d_119_v32_).f6), (d_244_v118_).intersection(d_244_v118_), True, d_82_globalState_)
                d_245_v119_ = out10_
                d_246_v120_ = out11_
                d_247_v121_ = out12_
                d_246_v120_ = ((0) - (d_83_v3_)) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubva"))))
                d_83_v3_ = (d_119_v32_).f8
            d_248_v122_: _dafny.Set
            d_249_v123_: _dafny.Array
            d_250_v124_: int
            d_251_v125_: _dafny.Array
            out13_: _dafny.Set
            out14_: _dafny.Array
            out15_: int
            out16_: _dafny.Array
            out13_, out14_, out15_, out16_ = (d_114_v28_).m8((d_119_v32_).f7, (d_119_v32_).f9, d_82_globalState_)
            d_248_v122_ = out13_
            d_249_v123_ = out14_
            d_250_v124_ = out15_
            d_251_v125_ = out16_
            d_252_v126_: _dafny.Map
            d_252_v126_ = _dafny.Map({(d_119_v32_).f6: -203})
            d_253_v127_: C3
            nw45_ = C3()
            nw45_.ctor__(d_79_v0_, (d_119_v32_).f9, d_250_v124_, (d_119_v32_).f7, (d_119_v32_).f6, False, _dafny.CodePoint('u'), (d_119_v32_).f5)
            d_253_v127_ = nw45_
            d_254_v128_: _dafny.Map
            d_254_v128_ = _dafny.Map({d_253_v127_: d_250_v124_})
            d_255_v129_: _dafny.Map
            d_255_v129_ = _dafny.Map({420: (0) - ((d_119_v32_).f8)})
            rhs22_ = (d_119_v32_).f6
            rhs23_ = default__.safeDivisionInt((d_83_v3_) + (len((d_252_v126_).set(d_79_v0_, len(d_254_v128_)))), default__.safeModuloInt(d_83_v3_, (0) - (((d_255_v129_)[d_250_v124_] if (d_250_v124_) in (d_255_v129_) else (d_119_v32_).f8))))
            d_79_v0_ = rhs22_
            d_83_v3_ = rhs23_
            d_256_v130_: D3
            d_256_v130_ = D3_DC11(not(d_253_v127_.f12), (d_119_v32_).f8, _dafny.Set({default__.fm7(d_113_v27_, d_250_v124_, (d_119_v32_).f6, d_82_globalState_), True}))
            d_257_v131_: _dafny.Map
            d_257_v131_ = _dafny.Map({d_256_v130_: d_83_v3_})
            d_258_v133_: _dafny.MultiSet
            d_258_v133_ = _dafny.MultiSet([D3_DC11((d_119_v32_).f6, d_250_v124_, d_188_v79_)])
            d_259_v134_: _dafny.Map
            def iife16_():
                coll8_ = _dafny.Map()
                compr_8_: D3
                for compr_8_ in (d_258_v133_).Elements:
                    d_260_v132_: D3 = compr_8_
                    if (d_260_v132_) in (d_258_v133_):
                        coll8_[d_260_v132_] = d_83_v3_
                return _dafny.Map(coll8_)
            d_259_v134_ = _dafny.Map({(d_119_v32_).f7: iife16_()
            })
            d_257_v131_ = ((_dafny.Map({d_256_v130_: (d_119_v32_).f8})) | (((d_259_v134_)[(d_119_v32_).f9] if ((d_119_v32_).f9) in (d_259_v134_) else _dafny.Map({d_256_v130_: (d_119_v32_).f8})))) | (d_257_v131_)
        elif True:
            if (d_119_v32_).f9:
                d_79_v0_ = (d_83_v3_) >= ((d_119_v32_).f8)
                d_79_v0_ = ((d_195_v86_)[d_83_v3_] if (d_83_v3_) in (d_195_v86_) else not((len((d_119_v32_).f5)) == (d_83_v3_)))
                d_261_v135_: _dafny.Map
                d_261_v135_ = _dafny.Map({(d_119_v32_).f6: d_188_v79_})
                d_262_v136_: _dafny.Map
                d_262_v136_ = _dafny.Map({len((d_192_v83_).f5): d_261_v135_})
                d_195_v86_ = (d_195_v86_).set(default__.safeDivisionInt(d_83_v3_, len(((d_262_v136_)[len(d_189_v80_)] if (len(d_189_v80_)) in (d_262_v136_) else d_261_v135_))), (d_119_v32_).f7)
                d_263_v137_: C3
                nw46_ = C3()
                nw46_.ctor__((d_119_v32_).f9, (d_119_v32_).f6, (len(d_189_v80_)) - (len(_dafny.SeqWithoutIsStrInference([d_113_v27_ for d_264_i12_ in range(default__.abs(370))]))), not ((d_189_v80_)[default__.safeIndex((d_119_v32_).f8, len(d_189_v80_))]) or ((d_119_v32_).f6), (d_119_v32_).f9, False, d_113_v27_, ((d_192_v83_).f5) + ((d_119_v32_).f5))
                d_263_v137_ = nw46_
                d_265_v138_: _dafny.Map
                d_265_v138_ = _dafny.Map({not(not((d_119_v32_).f6)): d_83_v3_})
                d_266_v139_: D5
                d_266_v139_ = D5_DC15((d_116_v30_).f10, d_79_v0_, d_265_v138_, len((d_119_v32_).f5))
                d_267_v140_: _dafny.Set
                d_267_v140_ = _dafny.Set({d_119_v32_.f4})
                d_268_v141_: _dafny.Set
                d_268_v141_ = _dafny.Set({len(d_267_v140_)})
                d_269_v142_: _dafny.Array
                nw47_ = _dafny.Array(None, 19)
                nw47_[int(0)] = d_268_v141_
                nw47_[int(1)] = d_268_v141_
                nw47_[int(2)] = d_268_v141_
                nw47_[int(3)] = d_268_v141_
                nw47_[int(4)] = d_268_v141_
                nw47_[int(5)] = d_268_v141_
                nw47_[int(6)] = d_268_v141_
                nw47_[int(7)] = d_268_v141_
                nw47_[int(8)] = d_268_v141_
                nw47_[int(9)] = d_268_v141_
                nw47_[int(10)] = d_268_v141_
                nw47_[int(11)] = d_268_v141_
                nw47_[int(12)] = _dafny.Set({(d_119_v32_).f8})
                nw47_[int(13)] = _dafny.Set({(d_119_v32_).f8})
                nw47_[int(14)] = d_268_v141_
                nw47_[int(15)] = d_268_v141_
                nw47_[int(16)] = d_268_v141_
                nw47_[int(17)] = d_268_v141_
                nw47_[int(18)] = _dafny.Set({(d_119_v32_).f8, d_83_v3_, (d_119_v32_).f8})
                d_269_v142_ = nw47_
                d_270_v143_: D14
                d_270_v143_ = D14_DC39(d_269_v142_)
                pat_let_tv5_ = d_269_v142_
                d_271_v144_: _dafny.Map
                def iife17_(_pat_let4_0):
                    def iife18_(d_272_dt__update__tmp_h0_):
                        def iife19_(_pat_let5_0):
                            def iife20_(d_273_dt__update_hcf72_h0_):
                                return D14_DC39(d_273_dt__update_hcf72_h0_)
                            return iife20_(_pat_let5_0)
                        return iife19_(pat_let_tv5_)
                    return iife18_(_pat_let4_0)
                d_271_v144_ = _dafny.Map({(d_266_v139_).cf29: iife17_(d_270_v143_)})
                d_271_v144_ = (d_271_v144_).set((d_83_v3_) - ((d_119_v32_).f8), D14_DC39(d_269_v142_))
            elif True:
                d_79_v0_ = d_79_v0_
                d_274_v145_: _dafny.Array
                nw48_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
                d_274_v145_ = nw48_
                index26_ = default__.safeIndex(39, (d_274_v145_).length(0))
                (d_274_v145_)[index26_] = (d_192_v83_).f5
                index27_ = default__.safeIndex(39, (d_274_v145_).length(0))
                (d_274_v145_)[index27_] = (_dafny.SeqWithoutIsStrInference([d_119_v32_.f4 for d_275_i13_ in range(default__.abs(183))])) + (d_84_v4_)
                d_276_v146_: C0
                nw49_ = C0()
                nw49_.ctor__((d_119_v32_).f6, d_79_v0_, d_119_v32_.f4, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igmblqb")))
                d_276_v146_ = nw49_
                d_277_v147_: _dafny.Map
                d_277_v147_ = _dafny.Map({(d_119_v32_).f9: d_276_v146_})
                (d_82_globalState_).f1 = ((len(d_277_v147_)) + ((d_119_v32_).f8)) - ((d_119_v32_).f8)
                d_193_v84_ = (d_193_v84_).set(False, (d_119_v32_).f9)
                d_278_v148_: D12
                d_278_v148_ = D12_DC33(not(False), (d_192_v83_).f5)
                d_279_v149_: _dafny.Map
                d_279_v149_ = _dafny.Map({d_278_v148_: (d_119_v32_).f7})
                d_279_v149_ = (d_279_v149_).set(d_278_v148_, (d_119_v32_).f6)
            d_280_v151_: _dafny.Set
            d_280_v151_ = _dafny.Set({271})
            d_281_v152_: _dafny.Map
            d_281_v152_ = _dafny.Map({not((d_119_v32_).f9): (d_119_v32_).f8})
            d_282_v153_: _dafny.Map
            d_282_v153_ = _dafny.Map({d_83_v3_: len(d_281_v152_)})
            d_283_v155_: C0
            nw50_ = C0()
            nw50_.ctor__((d_119_v32_).f9, (d_119_v32_).f7, d_113_v27_, (d_119_v32_).f5)
            d_283_v155_ = nw50_
            d_284_v156_: _dafny.Seq
            d_284_v156_ = _dafny.SeqWithoutIsStrInference([d_283_v155_])
            d_285_v157_: _dafny.Seq
            d_285_v157_ = _dafny.SeqWithoutIsStrInference([d_284_v156_])
            d_286_v158_: _dafny.MultiSet
            d_286_v158_ = _dafny.MultiSet([(d_119_v32_).f9, (d_119_v32_).f7, d_79_v0_])
            d_287_v160_: _dafny.Map
            def iife21_():
                coll9_ = _dafny.Set()
                compr_9_: int
                for compr_9_ in ((d_122_v35_)[default__.safeIndex(124, (d_122_v35_).length(0))]).Elements:
                    d_288_v159_: int = compr_9_
                    if (d_288_v159_) in ((d_122_v35_)[default__.safeIndex(124, (d_122_v35_).length(0))]):
                        coll9_ = coll9_.union(_dafny.Set([(d_288_v159_) * ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-752 for d_289_i14_ in range(default__.abs(852))]))).cardinality)]))
                return _dafny.Set(coll9_)
            d_287_v160_ = _dafny.Map({len((d_285_v157_)[default__.safeIndex(((d_286_v158_)[d_79_v0_] if (d_79_v0_) in (d_286_v158_) else (d_119_v32_).f8), len(d_285_v157_))]): iife21_()
            })
            d_290_v161_: _dafny.Array
            nw51_ = _dafny.Array(None, 15)
            def iife22_():
                coll10_ = _dafny.Set()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(381, -403):
                    d_291_v150_: int = compr_10_
                    if ((381) <= (d_291_v150_)) and ((d_291_v150_) < (-403)):
                        coll10_ = coll10_.union(_dafny.Set([default__.safeModuloInt(d_291_v150_, (0) - ((d_119_v32_).f8))]))
                return _dafny.Set(coll10_)
            nw51_[int(0)] = iife22_()
            
            nw51_[int(1)] = d_280_v151_
            nw51_[int(2)] = d_280_v151_
            nw51_[int(3)] = d_280_v151_
            nw51_[int(4)] = d_280_v151_
            def iife23_():
                coll11_ = _dafny.Set()
                compr_11_: int
                for compr_11_ in (d_282_v153_).keys.Elements:
                    d_292_v154_: int = compr_11_
                    if (d_292_v154_) in (d_282_v153_):
                        coll11_ = coll11_.union(_dafny.Set([default__.safeModuloInt(d_292_v154_, len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucmoqf")))})))]))
                return _dafny.Set(coll11_)
            nw51_[int(5)] = iife23_()
            
            nw51_[int(6)] = d_280_v151_
            nw51_[int(7)] = _dafny.Set({(d_119_v32_).f8})
            nw51_[int(8)] = _dafny.Set({(d_119_v32_).f8})
            nw51_[int(9)] = d_280_v151_
            nw51_[int(10)] = d_280_v151_
            nw51_[int(11)] = d_280_v151_
            nw51_[int(12)] = d_280_v151_
            nw51_[int(13)] = ((d_287_v160_)[(d_119_v32_).f8] if ((d_119_v32_).f8) in (d_287_v160_) else d_280_v151_)
            nw51_[int(14)] = d_280_v151_
            d_290_v161_ = nw51_
            d_293_v162_: D14
            d_293_v162_ = D14_DC39(d_290_v161_)
            source3_ = d_293_v162_
            if source3_.is_DC40:
                d_294___mcc_h15_ = source3_.cf73
                d_295___mcc_h16_ = source3_.cf74
                d_296___mcc_h17_ = source3_.cf75
                d_297_cf75_ = d_296___mcc_h17_
                d_298_cf74_ = d_295___mcc_h16_
                d_299_cf73_ = d_294___mcc_h15_
                default__.m0(d_82_globalState_)
                d_79_v0_ = ((d_119_v32_).f8) != (default__.safeModuloInt(d_83_v3_, d_299_cf73_))
                d_298_cf74_ = (d_119_v32_).f9
                rhs24_ = (d_119_v32_).f6
                rhs25_ = (d_119_v32_).f8
                rhs26_ = (d_298_cf74_) or (not((d_119_v32_).f6))
                rhs27_ = ((d_119_v32_).f8) - (d_299_cf73_)
                lhs11_ = d_82_globalState_
                d_79_v0_ = rhs24_
                lhs11_.f1 = rhs25_
                d_298_cf74_ = rhs26_
                d_83_v3_ = rhs27_
            elif source3_.is_DC39:
                d_300___mcc_h18_ = source3_.cf72
                d_301_cf72_ = d_300___mcc_h18_
                default__.m0(d_82_globalState_)
                d_302_v163_: _dafny.Seq
                d_302_v163_ = _dafny.SeqWithoutIsStrInference([d_114_v28_])
                d_303_v164_: _dafny.Array
                nw52_ = _dafny.Array(None, 14)
                nw52_[int(0)] = d_114_v28_
                nw52_[int(1)] = d_114_v28_
                nw52_[int(2)] = d_114_v28_
                nw52_[int(3)] = d_114_v28_
                nw52_[int(4)] = (d_302_v163_)[default__.safeIndex((d_119_v32_).f8, len(d_302_v163_))]
                nw52_[int(5)] = d_114_v28_
                nw52_[int(6)] = d_114_v28_
                nw52_[int(7)] = d_114_v28_
                nw52_[int(8)] = (d_114_v28_ if (d_119_v32_).f9 else d_114_v28_)
                nw52_[int(9)] = d_114_v28_
                nw52_[int(10)] = d_114_v28_
                nw52_[int(11)] = d_114_v28_
                nw52_[int(12)] = (D10_DC28((d_119_v32_).f9, (d_119_v32_).f8, d_79_v0_, (d_119_v32_).f7, d_114_v28_)).cf54
                nw52_[int(13)] = d_114_v28_
                d_303_v164_ = nw52_
                index28_ = default__.safeIndex(259, (d_303_v164_).length(0))
                (d_303_v164_)[index28_] = d_114_v28_
                index29_ = default__.safeIndex(259, (d_303_v164_).length(0))
                (d_303_v164_)[index29_] = d_114_v28_
                (d_82_globalState_).f1 = d_83_v3_
                d_293_v162_ = d_293_v162_
            elif True:
                d_304___mcc_h19_ = source3_.cf76
                d_305_cf76_ = d_304___mcc_h19_
                d_306_v165_: _dafny.Array
                def lambda13_(d_307_v80_):
                    def lambda14_(d_308_i15_):
                        return (d_307_v80_) + (d_307_v80_)

                    return lambda14_

                init7_ = lambda13_(d_189_v80_)
                nw53_ = _dafny.Array(None, 12)
                for i0_7_ in range(nw53_.length(0)):
                    nw53_[i0_7_] = init7_(i0_7_)
                d_306_v165_ = nw53_
                d_306_v165_ = d_306_v165_
                d_79_v0_ = (d_119_v32_).f7
                d_309_v166_: _dafny.Map
                d_309_v166_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(d_119_v32_).f8])): d_189_v80_})
                d_189_v80_ = (((((d_309_v166_)[(d_119_v32_).f8] if ((d_119_v32_).f8) in (d_309_v166_) else d_189_v80_)).set(default__.safeIndex(len((d_192_v83_).f5), len(((d_309_v166_)[(d_119_v32_).f8] if ((d_119_v32_).f8) in (d_309_v166_) else d_189_v80_))), (d_119_v32_).f9)) + (d_189_v80_)) + (d_189_v80_)
                d_79_v0_ = (d_119_v32_).f9
            d_310_v167_: D2
            d_310_v167_ = D2_DC8(d_83_v3_, (d_119_v32_).f8)
            d_311_v168_: D2
            d_311_v168_ = D2_DC9(d_310_v167_)
            d_311_v168_ = d_311_v168_
            d_312_v169_: D5
            d_312_v169_ = D5_DC14(d_80_v1_)
            d_313_v170_: _dafny.MultiSet
            d_313_v170_ = _dafny.MultiSet([d_312_v169_])
            d_313_v170_ = d_313_v170_
            d_79_v0_ = (d_119_v32_).f9
        d_314_v171_: D0
        d_314_v171_ = D0_DC0()
        d_315_v172_: _dafny.Map
        d_315_v172_ = _dafny.Map({d_314_v171_: (d_119_v32_).fm1(True, d_193_v84_, (d_119_v32_).f9, d_82_globalState_)})
        d_315_v172_ = (d_315_v172_) | (_dafny.Map({d_314_v171_: 763}))
        _dafny.print(_dafny.string_of(d_79_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_80_v1_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_80_v1_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_81_v2_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_82_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_82_globalState_).f3)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_82_globalState_).f3)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_83_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_84_v4_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_v5_) == (_dafny.MultiSet([_dafny.Set({True, False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_86_v6_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikmk"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v7_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikmk"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_89_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_110_v24_) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_111_v25_) == (_dafny.SeqWithoutIsStrInference([372, 287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v26_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([372, 287])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_113_v27_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_115_v29_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_v30_).f10)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_v30_).f11)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_v30_).f11)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_118_v31_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v32_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v32_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v32_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v32_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_119_v32_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_119_v32_).f5).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_120_v33_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v34_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v34_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_121_v34_).cf6) == (_dafny.SeqWithoutIsStrInference([372, 287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[0]) == (_dafny.MultiSet([372, 287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[1]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[2]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[3]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[4]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[5]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[6]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[7]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[8]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[9]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[10]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[11]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[12]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[13]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[14]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[15]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[16]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[17]) == (_dafny.MultiSet([-78]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[18]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[19]) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[20]) == (_dafny.MultiSet([-287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, -287, 287, 287, 2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[21]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[22]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[23]) == (_dafny.MultiSet([4, 287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[24]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[25]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[26]) == (_dafny.MultiSet([287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v35_)[27]) == (_dafny.MultiSet([1, 5, -415, 287, 287]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v79_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_189_v80_) == (_dafny.SeqWithoutIsStrInference([True, False, False, True, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_191_v82_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_192_v83_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_192_v83_).f5).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_193_v84_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v86_) == (_dafny.Map({287: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_315_v172_) == (_dafny.Map({D0_DC0(): 763}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)

class D0_DC0(D0, NamedTuple('DC0', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf0', Any), ('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(_dafny.CodePoint('D'), int(0), _dafny.CodePoint('D'), int(0), False)
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

class D1_DC4(D1, NamedTuple('DC4', [('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)

class D2_DC7(D2, NamedTuple('DC7', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC11(False, int(0), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)

class D3_DC11(D3, NamedTuple('DC11', [('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({self.cf19.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC13(D4, NamedTuple('DC13', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC15(_dafny.Array(None, 0), False, _dafny.Map({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC15(D5, NamedTuple('DC15', [('cf26', Any), ('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC17(D6, NamedTuple('DC17', [('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({self.cf31.VerbatimString(True)}, {self.cf32.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC19(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC19(D7, NamedTuple('DC19', [('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf34)}, {self.cf35.VerbatimString(True)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC23(D8, NamedTuple('DC23', [('cf42', Any), ('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC25(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)

class D9_DC25(D9, NamedTuple('DC25', [('cf45', Any), ('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC27(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)

class D10_DC27(D10, NamedTuple('DC27', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf50', Any), ('cf51', Any), ('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC26(D10, NamedTuple('DC26', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC30(D7.default()(), False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)

class D11_DC30(D11, NamedTuple('DC30', [('cf56', Any), ('cf57', Any), ('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC32(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)

class D12_DC32(D12, NamedTuple('DC32', [('cf61', Any), ('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf61 == __o.cf61 and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC33(D12, NamedTuple('DC33', [('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf63)}, {self.cf64.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC31(D12, NamedTuple('DC31', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC36(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D13_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D13_DC37)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D13_DC38)

class D13_DC36(D13, NamedTuple('DC36', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC37(D13, NamedTuple('DC37', [('cf68', Any), ('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC37({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC37) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC35(D13, NamedTuple('DC35', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC38(D13, NamedTuple('DC38', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC38({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC38) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC40(int(0), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D14_DC40)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D14_DC39)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)

class D14_DC40(D14, NamedTuple('DC40', [('cf73', Any), ('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC40({_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {self.cf75.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC40) and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC39(D14, NamedTuple('DC39', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC39({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC39) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC41(D14, NamedTuple('DC41', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC43(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D15_DC42)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D15_DC44)

class D15_DC43(D15, NamedTuple('DC43', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC42(D15, NamedTuple('DC42', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC42({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC42) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC44(D15, NamedTuple('DC44', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC44({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC44) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC46(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D16_DC46)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D16_DC47)

class D16_DC46(D16, NamedTuple('DC46', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC46({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC46) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC45(D16, NamedTuple('DC45', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC47(D16, NamedTuple('DC47', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC47({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC47) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value

class T1:
    pass

class T2:
    pass
    def m1(self, p0, p1, p2, p3, globalState):
        pass


class T3:
    pass
    def m2(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f1: int = int(0)
        self._f0: bool = False
        self._f2: bool = False
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3

    @property
    def f0(self):
        return self._f0
    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3

class C0(T1, T0):
    def  __init__(self):
        self._f4: str = _dafny.CodePoint('D')
        self._f6: bool = False
        self._f7: bool = False
        self._f5: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f6, f7, f4, f5):
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f4 = f4
        (self)._f5 = f5

    def fm9(self, p0, p1, globalState):
        def iife24_():
            coll12_ = _dafny.Map()
            compr_12_: _dafny.MultiSet
            for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(D0_DC2(821, (self).f6, _dafny.SeqWithoutIsStrInference([318 for d_316_i1_ in range(default__.abs(755))]))).cf5, (self).f6]) for d_317_i0_ in range(default__.abs(759))])).Elements:
                d_318_v0_: _dafny.MultiSet = compr_12_
                if (d_318_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(D0_DC2(821, (self).f6, _dafny.SeqWithoutIsStrInference([318 for d_316_i1_ in range(default__.abs(755))]))).cf5, (self).f6]) for d_317_i0_ in range(default__.abs(759))])):
                    coll12_[d_318_v0_] = _dafny.SeqWithoutIsStrInference([D2_DC7((self).f7), D2_DC7((self).f6)])
            return _dafny.Map(coll12_)
        return len(iife24_()
        )


class C1(T2, T1, T0):
    def  __init__(self):
        self._f4: str = _dafny.CodePoint('D')
        self._f8: int = int(0)
        self._f9: bool = False
        self._f6: bool = False
        self._f7: bool = False
        self._f5: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f8, f9, f6, f7, f4, f5):
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f4 = f4
        (self)._f5 = f5

    def fm1(self, p0, p1, p2, globalState):
        return -485

    def fm2(self, globalState):
        source4_ = D3_DC12(D3_DC12(D3_DC11((self).f9, (self).f8, _dafny.Set({(self).f9}))))
        if source4_.is_DC11:
            d_319___mcc_h0_ = source4_.cf20
            d_320___mcc_h1_ = source4_.cf21
            d_321___mcc_h2_ = source4_.cf22
            d_322_cf22_ = d_321___mcc_h2_
            d_323_cf21_ = d_320___mcc_h1_
            d_324_cf20_ = d_319___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([d_323_cf21_])
        elif source4_.is_DC10:
            d_325___mcc_h3_ = source4_.cf19
            d_326_cf19_ = d_325___mcc_h3_
            return (_dafny.SeqWithoutIsStrInference([(self).f8 for d_327_i0_ in range(default__.abs(-920))])) + (_dafny.SeqWithoutIsStrInference([537]))
        elif True:
            d_328___mcc_h4_ = source4_.cf23
            d_329_cf23_ = d_328___mcc_h4_
            def iife25_():
                coll13_ = _dafny.Set()
                compr_13_: int
                for compr_13_ in _dafny.IntegerRange(372, 170):
                    d_330_v0_: int = compr_13_
                    if ((372) <= (d_330_v0_)) and ((d_330_v0_) < (170)):
                        coll13_ = coll13_.union(_dafny.Set([(d_330_v0_) + ((self).f8)]))
                return _dafny.Set(coll13_)
            return _dafny.SeqWithoutIsStrInference([len(iife25_()
            ), (_dafny.MultiSet([self.f4])).cardinality])

    def fm19(self, p0, p1, p2, globalState):
        return (self).f7

    def fm20(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([D1_DC5((self).f8) for d_331_i0_ in range(default__.abs(639))])) + (_dafny.SeqWithoutIsStrInference([D1_DC5((self).f8), D1_DC5((self).f8), D1_DC5((self).f8)]))

    def m1(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_332_v0_: D2
        d_332_v0_ = D2_DC7(p1)
        d_333_v1_: _dafny.Seq
        d_333_v1_ = _dafny.SeqWithoutIsStrInference([(d_332_v0_).cf15])
        d_334_v2_: _dafny.Set
        d_334_v2_ = _dafny.Set({(self).f7, False})
        d_335_v3_: _dafny.Seq
        d_335_v3_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_334_v2_))])
        d_336_v4_: _dafny.Seq
        d_336_v4_ = _dafny.SeqWithoutIsStrInference([len(d_335_v3_)])
        d_337_v5_: D0
        d_337_v5_ = D0_DC2((self).f8, (d_333_v1_)[default__.safeIndex((self).f8, len(d_333_v1_))], d_335_v3_)
        d_338_v6_: _dafny.Map
        d_338_v6_ = _dafny.Map({p3: (d_337_v5_).cf5})
        d_339_v7_: _dafny.Set
        def iife26_(_pat_let6_0):
            def iife27_(d_340_dt__update__tmp_h0_):
                def iife28_(_pat_let7_0):
                    def iife29_(d_341_dt__update_hcf15_h0_):
                        return D2_DC7(d_341_dt__update_hcf15_h0_)
                    return iife29_(_pat_let7_0)
                return iife28_(True)
            return iife27_(_pat_let6_0)
        d_339_v7_ = _dafny.Set({iife26_(D2_DC7(p0)), d_332_v0_, d_332_v0_})
        d_342_v8_: D3
        d_342_v8_ = D3_DC12(D3_DC10((self).f5))
        d_343_v9_: D3
        d_343_v9_ = D3_DC12(d_342_v8_)
        d_344_v10_: D3
        d_344_v10_ = D3_DC12(d_342_v8_)
        d_345_v11_: _dafny.MultiSet
        d_345_v11_ = _dafny.MultiSet([d_344_v10_])
        d_346_v12_: _dafny.Array
        nw54_ = _dafny.Array(None, 9)
        nw54_[int(0)] = (self).fm1((self).f6, d_338_v6_, (self).f6, globalState)
        nw54_[int(1)] = (self).f8
        nw54_[int(2)] = ((self).f8 if (self).f6 else 31)
        nw54_[int(3)] = (self).f8
        nw54_[int(4)] = len((d_333_v1_).set(default__.safeIndex(531, len(d_333_v1_)), (self).f7))
        nw54_[int(5)] = (0) - ((593 if (self).fm19(False, p0, d_339_v7_, globalState) else (self).f8))
        nw54_[int(6)] = (self).f8
        nw54_[int(7)] = ((d_345_v11_).set(d_344_v10_, default__.abs((self).f8))).cardinality
        nw54_[int(8)] = (self).f8
        d_346_v12_ = nw54_
        index30_ = default__.safeIndex(696, (d_346_v12_).length(0))
        (d_346_v12_)[index30_] = ((self).f8) * ((self).f8)
        d_347_v13_: _dafny.Array
        nw55_ = _dafny.Array(False, 16)
        d_347_v13_ = nw55_
        index31_ = default__.safeIndex(378, (d_347_v13_).length(0))
        (d_347_v13_)[index31_] = ((self).f7 if p0 else p3)
        pat_let_tv6_ = p3
        pat_let_tv7_ = d_336_v4_
        index32_ = default__.safeIndex(696, (d_346_v12_).length(0))
        index33_ = default__.safeIndex(378, (d_347_v13_).length(0))
        def lambda15_(source5_):
            if source5_.is_DC0:
                return (self).f8
            elif source5_.is_DC1:
                d_348___mcc_h0_ = source5_.cf0
                d_349___mcc_h1_ = source5_.cf1
                d_350___mcc_h2_ = source5_.cf2
                d_351___mcc_h3_ = source5_.cf3
                d_352_cf3_ = d_351___mcc_h3_
                d_353_cf2_ = d_350___mcc_h2_
                d_354_cf1_ = d_349___mcc_h1_
                d_355_cf0_ = d_348___mcc_h0_
                return len((self).f5)
            elif True:
                d_356___mcc_h4_ = source5_.cf4
                d_357___mcc_h5_ = source5_.cf5
                d_358___mcc_h6_ = source5_.cf6
                d_359_cf6_ = d_358___mcc_h6_
                d_360_cf5_ = d_357___mcc_h5_
                d_361_cf4_ = d_356___mcc_h4_
                return (self).f8

        def iife30_(_pat_let8_0):
            def iife31_(d_362_dt__update__tmp_h1_):
                def iife32_(_pat_let9_0):
                    def iife33_(d_363_dt__update_hcf5_h0_):
                        def iife34_(_pat_let10_0):
                            def iife35_(d_364_dt__update_hcf6_h0_):
                                return D0_DC2((d_362_dt__update__tmp_h1_).cf4, d_363_dt__update_hcf5_h0_, d_364_dt__update_hcf6_h0_)
                            return iife35_(_pat_let10_0)
                        return iife34_(pat_let_tv7_)
                    return iife33_(_pat_let9_0)
                return iife32_(pat_let_tv6_)
            return iife31_(_pat_let8_0)
        rhs28_ = lambda15_(iife30_(default__.fm21((self).f6, True, (self).f9, p0, globalState)))
        rhs29_ = True
        rhs30_ = default__.safeDivisionInt((self).f8, (self).f8)
        lhs12_ = d_346_v12_
        lhs13_ = default__.safeIndex(696, (d_346_v12_).length(0))
        lhs14_ = d_347_v13_
        lhs15_ = default__.safeIndex(378, (d_347_v13_).length(0))
        lhs12_[lhs13_] = rhs28_
        lhs14_[lhs15_] = rhs29_
        r1 = rhs30_
        index34_ = default__.safeIndex(378, (d_347_v13_).length(0))
        (d_347_v13_)[index34_] = False
        d_365_v14_: C0
        nw56_ = C0()
        nw56_.ctor__(not(p3), p0, self.f4, ((self).f5 if (d_347_v13_)[default__.safeIndex(378, (d_347_v13_).length(0))] else (self).f5))
        d_365_v14_ = nw56_
        d_366_v15_: _dafny.Set
        d_366_v15_ = _dafny.Set({(self).f5, (self).f5})
        d_367_v16_: _dafny.Set
        d_367_v16_ = _dafny.Set({(d_346_v12_)[default__.safeIndex(696, (d_346_v12_).length(0))], (self).f8})
        d_368_v17_: _dafny.Map
        d_368_v17_ = _dafny.Map({(_dafny.Set({(self).f5, ((self).f5).set(default__.safeIndex(len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_369_i0_ in range(default__.abs(668))])).set(default__.safeIndex(len((self).f5), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_370_i0_ in range(default__.abs(668))]))), self.f4)), len((self).f5)), self.f4), _dafny.SeqWithoutIsStrInference([self.f4]), (self).f5, (self).f5})).intersection(d_366_v15_): (d_367_v16_).isdisjoint(_dafny.Set({(self).f8, (self).f8}))})
        d_368_v17_ = (d_368_v17_).set((d_366_v15_) - (_dafny.Set({(self).f5})), (self).f7)
        d_371_v18_: _dafny.Set
        d_371_v18_ = _dafny.Set({d_347_v13_, d_347_v13_})
        index35_ = default__.safeIndex(696, (d_346_v12_).length(0))
        (d_346_v12_)[index35_] = len(d_371_v18_)
        d_372_v19_: C0
        nw57_ = C0()
        nw57_.ctor__(p0, (self).f7, self.f4, (self).f5)
        d_372_v19_ = nw57_
        r0 = (self).f8
        r1 = (0) - (len(d_335_v3_))
        d_373_v20_: _dafny.Array
        nw58_ = _dafny.Array(_dafny.Array(None, 0), 14)
        d_373_v20_ = nw58_
        r2 = d_373_v20_
        return r0, r1, r2

    def m8(self, p0, p1, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        r3: _dafny.Array = _dafny.Array(None, 0)
        (globalState).f1 = (self).f8
        d_374_v0_: _dafny.Array
        nw59_ = _dafny.Array(_dafny.MultiSet({}), 7)
        d_374_v0_ = nw59_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_374_v0_).length(0)):
            d_375_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_375_i0_)) and ((d_375_i0_) < ((d_374_v0_).length(0)))):
                (d_374_v0_)[(d_375_i0_)] = (_dafny.MultiSet([(self).f8, (self).f8, (0) - ((0) - (len((self).f5))), -67])) - (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p1, (self).f9, (self).f6]))])) + (_dafny.SeqWithoutIsStrInference([173]))))
        d_376_i1_: int
        d_376_i1_ = 0
        with _dafny.label("1"):
            while False:
                with _dafny.c_label("1"):
                    if (d_376_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_376_i1_ = (d_376_i1_) + (1)
                    d_377_v1_: D1
                    d_377_v1_ = D1_DC4(self.f4, (self).f8, self.f4, (self).f8, True)
                    d_378_v2_: _dafny.Seq
                    d_378_v2_ = _dafny.SeqWithoutIsStrInference([self.f4, (d_377_v1_).cf10])
                    d_378_v2_ = ((self).f5).set(default__.safeIndex(default__.safeModuloInt((0) - ((self).f8), (self).f8), len((self).f5)), self.f4)
                    d_379_v3_: bool
                    d_379_v3_ = False
                    d_379_v3_ = (self).f6
                    (globalState).f1 = 76
                    d_380_v4_: D3
                    d_380_v4_ = D3_DC10(d_378_v2_)
                    d_381_v5_: D3
                    d_381_v5_ = D3_DC12(d_380_v4_)
                    d_381_v5_ = D3_DC12(d_380_v4_)
                    pass
            pass
        default__.m0(globalState)
        d_382_v6_: bool
        d_382_v6_ = False
        d_382_v6_ = (self).f7
        d_383_v7_: _dafny.Seq
        d_383_v7_ = _dafny.SeqWithoutIsStrInference([(self).f8])
        d_384_v8_: _dafny.Map
        d_384_v8_ = _dafny.Map({p1: (self).fm2(globalState)})
        d_385_v9_: _dafny.Set
        d_385_v9_ = _dafny.Set({p0, (self).f7, (self).f9, d_382_v6_})
        d_386_v10_: _dafny.Seq
        d_386_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_385_v9_)]), d_383_v7_, (self).fm2(globalState)])
        d_387_v11_: _dafny.Array
        nw60_ = _dafny.Array(None, 23)
        nw60_[int(0)] = _dafny.SeqWithoutIsStrInference([(self).f8])
        nw60_[int(1)] = d_383_v7_
        nw60_[int(2)] = d_383_v7_
        nw60_[int(3)] = ((d_383_v7_) + (d_383_v7_)).set(default__.safeIndex((self).f8, len((d_383_v7_) + (d_383_v7_))), (self).f8)
        nw60_[int(4)] = d_383_v7_
        nw60_[int(5)] = d_383_v7_
        nw60_[int(6)] = ((((d_384_v8_)[p1] if (p1) in (d_384_v8_) else (d_386_v10_)[default__.safeIndex(len((self).f5), len(d_386_v10_))]) if False else d_383_v7_)).set(default__.safeIndex((0) - ((self).f8), len((((d_384_v8_)[p1] if (p1) in (d_384_v8_) else (d_386_v10_)[default__.safeIndex(len((self).f5), len(d_386_v10_))]) if False else d_383_v7_))), (self).f8)
        nw60_[int(7)] = _dafny.SeqWithoutIsStrInference([(self).f8, (self).f8, (self).f8, (self).f8, (self).f8])
        nw60_[int(8)] = d_383_v7_
        nw60_[int(9)] = d_383_v7_
        nw60_[int(10)] = d_383_v7_
        nw60_[int(11)] = d_383_v7_
        nw60_[int(12)] = _dafny.SeqWithoutIsStrInference([(self).f8, (self).f8, (self).f8, (self).f8])
        nw60_[int(13)] = (d_383_v7_) + (d_383_v7_)
        nw60_[int(14)] = d_383_v7_
        nw60_[int(15)] = (d_383_v7_) + (d_383_v7_)
        nw60_[int(16)] = d_383_v7_
        nw60_[int(17)] = d_383_v7_
        nw60_[int(18)] = (self).fm2(globalState)
        nw60_[int(19)] = d_383_v7_
        nw60_[int(20)] = d_383_v7_
        nw60_[int(21)] = d_383_v7_
        nw60_[int(22)] = (d_383_v7_) + (d_383_v7_)
        d_387_v11_ = nw60_
        d_388_v12_: _dafny.Seq
        d_388_v12_ = _dafny.SeqWithoutIsStrInference([d_382_v6_])
        index36_ = default__.safeIndex(491, (d_387_v11_).length(0))
        (d_387_v11_)[index36_] = ((self).fm2(globalState)) + ((d_383_v7_).set(default__.safeIndex((d_383_v7_)[default__.safeIndex((self).f8, len(d_383_v7_))], len(d_383_v7_)), len((d_388_v12_).set(default__.safeIndex(len(d_383_v7_), len(d_388_v12_)), p1))))
        d_389_v13_: D0
        d_389_v13_ = D0_DC2(981, (self).f7, d_383_v7_)
        index37_ = default__.safeIndex(491, (d_387_v11_).length(0))
        (d_387_v11_)[index37_] = (d_389_v13_).cf6
        d_390_v14_: _dafny.Map
        d_390_v14_ = _dafny.Map({d_382_v6_: _dafny.SeqWithoutIsStrInference([self.f4 for d_391_i2_ in range(default__.abs(697))])})
        d_392_v15_: _dafny.Map
        d_392_v15_ = _dafny.Map({p0: p0})
        d_393_v16_: _dafny.Set
        d_393_v16_ = _dafny.Set({d_392_v15_})
        d_394_v17_: _dafny.Seq
        d_394_v17_ = _dafny.SeqWithoutIsStrInference([d_393_v16_])
        d_395_v18_: _dafny.Seq
        d_395_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f7: (self).f7}), default__.fm23((self).f8, (self).f8, self.f4, (self).f6, globalState)])
        def iife36_():
            coll14_ = _dafny.Set()
            compr_14_: _dafny.Map
            for compr_14_ in (d_395_v18_).Elements:
                d_396_v19_: _dafny.Map = compr_14_
                if (d_396_v19_) in (d_395_v18_):
                    coll14_ = coll14_.union(_dafny.Set([d_396_v19_]))
            return _dafny.Set(coll14_)
        r0 = (((d_394_v17_)[default__.safeIndex((self).f8, len(d_394_v17_))] if (self).f6 else default__.fm22(globalState)) if (d_382_v6_) in (d_390_v14_) else (d_393_v16_).intersection(iife36_()
        ))
        d_397_v20_: _dafny.Array
        def lambda16_(d_398_i3_):
            return (self).f7

        init8_ = lambda16_
        nw61_ = _dafny.Array(None, 2)
        for i0_8_ in range(nw61_.length(0)):
            nw61_[i0_8_] = init8_(i0_8_)
        d_397_v20_ = nw61_
        r1 = d_397_v20_
        r2 = (0) - (((0) - (len(_dafny.SeqWithoutIsStrInference([self.f4, self.f4, self.f4])))) - ((self).f8))
        r3 = d_397_v20_
        return r0, r1, r2, r3


class C2:
    def  __init__(self):
        self.f15: _dafny.Array = _dafny.Array(None, 0)
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f14, f15):
        (self)._f14 = f14
        (self).f15 = f15

    def fm6(self, globalState):
        if (_dafny.Map({not(True): False})) == (_dafny.Map({not(False): True})):
            return _dafny.SeqWithoutIsStrInference([_dafny.Set({True})])
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.Set({False}), _dafny.Set({False, True, True})])

    def m6(self, globalState):
        r0: int = int(0)
        r1: D1 = D1.default()()
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: D1 = D1.default()()
        d_399_v0_: _dafny.Seq
        d_399_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "svbxxdv"))
        d_400_i0_: int
        d_400_i0_ = 0
        with _dafny.label("2"):
            while (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfrnqnere"))) == (d_399_v0_):
                with _dafny.c_label("2"):
                    if (d_400_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_400_i0_ = (d_400_i0_) + (1)
                    d_401_v1_: bool
                    d_401_v1_ = True
                    d_402_v2_: int
                    d_402_v2_ = -884
                    d_403_v3_: _dafny.Array
                    def lambda17_(d_404_v1_):
                        def lambda18_(d_405_i1_):
                            return _dafny.SeqWithoutIsStrInference([not(d_404_v1_)])

                        return lambda18_

                    init9_ = lambda17_(d_401_v1_)
                    nw62_ = _dafny.Array(None, 2)
                    for i0_9_ in range(nw62_.length(0)):
                        nw62_[i0_9_] = init9_(i0_9_)
                    d_403_v3_ = nw62_
                    d_406_v4_: _dafny.Map
                    d_406_v4_ = _dafny.Map({d_402_v2_: D1_DC3(d_403_v3_)})
                    d_407_v5_: _dafny.Map
                    d_407_v5_ = _dafny.Map({d_401_v1_: d_406_v4_})
                    d_408_v6_: D2
                    d_408_v6_ = D2_DC6(d_406_v4_)
                    d_407_v5_ = (d_407_v5_).set((d_401_v1_) == (d_401_v1_), (d_406_v4_) | ((d_408_v6_).cf14))
                    d_409_v7_: str
                    d_409_v7_ = _dafny.CodePoint('y')
                    d_410_v8_: _dafny.Map
                    d_410_v8_ = _dafny.Map({d_401_v1_: d_409_v7_})
                    d_411_v9_: D1
                    d_411_v9_ = D1_DC3(d_403_v3_)
                    d_412_v10_: _dafny.MultiSet
                    d_412_v10_ = _dafny.MultiSet([d_411_v9_])
                    d_413_v11_: _dafny.Map
                    d_413_v11_ = _dafny.Map({default__.fm7(d_409_v7_, (d_412_v10_).cardinality, d_401_v1_, globalState): d_401_v1_})
                    d_414_v12_: _dafny.Map
                    d_414_v12_ = _dafny.Map({d_410_v8_: d_413_v11_})
                    d_414_v12_ = d_414_v12_
                    d_399_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
                    d_415_v13_: _dafny.Seq
                    d_415_v13_ = _dafny.SeqWithoutIsStrInference([not(d_401_v1_)])
                    d_416_v14_: _dafny.Set
                    d_416_v14_ = _dafny.Set({d_401_v1_, d_401_v1_})
                    d_417_v15_: D0
                    d_417_v15_ = D0_DC2(d_402_v2_, d_401_v1_, _dafny.SeqWithoutIsStrInference([len(d_416_v14_), len(d_399_v0_)]))
                    d_399_v0_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aetllib"))) + ((d_399_v0_).set(default__.safeIndex(default__.fm8(len(_dafny.SeqWithoutIsStrInference([d_409_v7_])), d_401_v1_, d_401_v1_, globalState), len(d_399_v0_)), _dafny.CodePoint('o')))) + ((d_399_v0_).set(default__.safeIndex(len((d_415_v13_).set(default__.safeIndex(d_402_v2_, len(d_415_v13_)), (d_417_v15_).cf5)), len(d_399_v0_)), d_409_v7_))
                    pass
            pass
        if True:
            d_418_v16_: bool
            d_418_v16_ = False
            d_419_v17_: str
            d_419_v17_ = _dafny.CodePoint('j')
            d_420_v18_: C0
            nw63_ = C0()
            nw63_.ctor__(d_418_v16_, d_418_v16_, d_419_v17_, d_399_v0_)
            d_420_v18_ = nw63_
            d_421_v19_: _dafny.Map
            d_421_v19_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmxt")): d_418_v16_})
            d_421_v19_ = (d_421_v19_) | ((d_421_v19_).set(d_399_v0_, d_418_v16_))
            d_422_v20_: int
            d_422_v20_ = 475
            d_423_v21_: _dafny.Set
            d_423_v21_ = _dafny.Set({d_422_v20_, d_422_v20_})
            (globalState).f1 = len((d_423_v21_ if (d_423_v21_).isdisjoint(d_423_v21_) else default__.fm10(True, d_422_v20_, default__.fm8(d_422_v20_, d_418_v16_, d_418_v16_, globalState), globalState)))
            d_424_v22_: D1
            d_424_v22_ = D1_DC5(d_422_v20_)
            d_424_v22_ = (d_424_v22_ if d_418_v16_ else d_424_v22_)
            d_425_v23_: _dafny.Seq
            d_425_v23_ = _dafny.SeqWithoutIsStrInference([d_422_v20_, (0) - (d_422_v20_)])
            d_426_v24_: _dafny.Map
            d_426_v24_ = _dafny.Map({d_399_v0_: d_422_v20_})
            d_427_v25_: _dafny.Array
            nw64_ = _dafny.Array(None, 4)
            nw64_[int(0)] = (d_425_v23_) + (d_425_v23_)
            nw64_[int(1)] = d_425_v23_
            nw64_[int(2)] = (d_425_v23_) + (_dafny.SeqWithoutIsStrInference([d_422_v20_, len(d_426_v24_), d_422_v20_, (_dafny.MultiSet([d_418_v16_, d_418_v16_])).cardinality, 782]))
            nw64_[int(3)] = (d_425_v23_).set(default__.safeIndex(d_422_v20_, len(d_425_v23_)), d_422_v20_)
            d_427_v25_ = nw64_
            d_427_v25_ = d_427_v25_
        elif True:
            d_428_v26_: int
            d_428_v26_ = 650
            (globalState).f1 = (d_428_v26_) - (d_428_v26_)
            d_429_v27_: bool
            d_429_v27_ = False
            d_429_v27_ = d_429_v27_
            d_430_v28_: _dafny.Seq
            d_430_v28_ = _dafny.SeqWithoutIsStrInference([d_428_v26_, d_428_v26_])
            d_431_v29_: D0
            d_431_v29_ = D0_DC2(d_428_v26_, False, d_430_v28_)
            if not((d_431_v29_).cf5):
                d_432_v30_: _dafny.Array
                nw65_ = _dafny.Array(int(0), 28)
                d_432_v30_ = nw65_
                d_433_v31_: _dafny.Seq
                d_433_v31_ = _dafny.SeqWithoutIsStrInference([d_429_v27_, d_429_v27_, d_429_v27_])
                d_434_v32_: _dafny.MultiSet
                d_434_v32_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_429_v27_]), d_433_v31_])
                index38_ = default__.safeIndex(333, (d_432_v30_).length(0))
                (d_432_v30_)[index38_] = ((d_434_v32_).cardinality) * (d_428_v26_)
                d_435_v33_: _dafny.Array
                nw66_ = _dafny.Array(_dafny.Seq({}), 26)
                d_435_v33_ = nw66_
                d_436_v34_: D1
                d_436_v34_ = D1_DC3(d_435_v33_)
                d_437_v35_: _dafny.Map
                d_437_v35_ = _dafny.Map({d_428_v26_: d_436_v34_})
                d_438_v36_: D2
                d_438_v36_ = D2_DC6(_dafny.Map({d_428_v26_: d_436_v34_}))
                d_439_v37_: _dafny.MultiSet
                d_439_v37_ = _dafny.MultiSet([(D2_DC6(d_437_v35_) if d_429_v27_ else d_438_v36_), d_438_v36_])
                index39_ = default__.safeIndex(333, (d_432_v30_).length(0))
                (d_432_v30_)[index39_] = (d_439_v37_).cardinality
                index40_ = default__.safeIndex(310, ((self).f14).length(0))
                ((self).f14)[index40_] = True
                d_440_v38_: _dafny.MultiSet
                d_440_v38_ = _dafny.MultiSet([d_429_v27_, (d_431_v29_).cf5, d_429_v27_])
                index41_ = default__.safeIndex(310, ((self).f14).length(0))
                ((self).f14)[index41_] = (181) > (((d_440_v38_)[d_429_v27_] if (d_429_v27_) in (d_440_v38_) else (d_432_v30_)[default__.safeIndex(333, (d_432_v30_).length(0))]))
                d_441_v39_: str
                d_441_v39_ = _dafny.CodePoint('h')
                d_442_v40_: C0
                nw67_ = C0()
                nw67_.ctor__(((self).f14)[default__.safeIndex(310, ((self).f14).length(0))], (d_429_v27_ if d_429_v27_ else ((self).f14)[default__.safeIndex(310, ((self).f14).length(0))]), d_441_v39_, d_399_v0_)
                d_442_v40_ = nw67_
                d_442_v40_ = d_442_v40_
                index42_ = default__.safeIndex(333, (d_432_v30_).length(0))
                index43_ = default__.safeIndex(310, ((self).f14).length(0))
                rhs31_ = ((self).f14)[default__.safeIndex(310, ((self).f14).length(0))]
                rhs32_ = d_428_v26_
                rhs33_ = ((self).f14)[default__.safeIndex(310, ((self).f14).length(0))]
                rhs34_ = d_441_v39_
                lhs16_ = d_432_v30_
                lhs17_ = default__.safeIndex(333, (d_432_v30_).length(0))
                lhs18_ = (self).f14
                lhs19_ = default__.safeIndex(310, ((self).f14).length(0))
                d_429_v27_ = rhs31_
                lhs16_[lhs17_] = rhs32_
                lhs18_[lhs19_] = rhs33_
                d_441_v39_ = rhs34_
                r0 = d_428_v26_
            elif True:
                d_443_v41_: _dafny.Set
                d_443_v41_ = _dafny.Set({d_429_v27_})
                d_444_v42_: D2
                d_444_v42_ = D2_DC8((d_428_v26_) + (d_428_v26_), d_428_v26_)
                d_445_v43_: _dafny.Seq
                d_445_v43_ = _dafny.SeqWithoutIsStrInference([d_429_v27_, d_429_v27_, d_429_v27_, False])
                d_446_v44_: _dafny.MultiSet
                d_446_v44_ = _dafny.MultiSet([len(((d_445_v43_ if not(d_429_v27_) else d_445_v43_)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibbeoyt"))), len((d_445_v43_ if not(d_429_v27_) else d_445_v43_))), d_429_v27_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjhcwwr"))), d_428_v26_])
                rhs35_ = ((d_446_v44_)[d_428_v26_] if (d_428_v26_) in (d_446_v44_) else (0) - (d_428_v26_))
                rhs36_ = (d_443_v41_) | (d_443_v41_)
                rhs37_ = d_444_v42_
                rhs38_ = _dafny.SeqWithoutIsStrInference([d_428_v26_ for d_447_i2_ in range(default__.abs(0))])
                rhs39_ = (d_445_v43_) + (d_445_v43_)
                lhs20_ = globalState
                lhs20_.f1 = rhs35_
                d_443_v41_ = rhs36_
                d_444_v42_ = rhs37_
                d_430_v28_ = rhs38_
                d_445_v43_ = rhs39_
                d_448_v45_: _dafny.Array
                nw68_ = _dafny.Array(_dafny.CodePoint('D'), 15)
                d_448_v45_ = nw68_
                d_448_v45_ = d_448_v45_
                index44_ = default__.safeIndex(550, ((self).f14).length(0))
                ((self).f14)[index44_] = d_429_v27_
                index45_ = default__.safeIndex(550, ((self).f14).length(0))
                ((self).f14)[index45_] = d_429_v27_
                (globalState).f1 = d_428_v26_
                d_449_v46_: _dafny.Set
                d_449_v46_ = _dafny.Set({d_428_v26_, 666})
                d_450_v47_: _dafny.Seq
                d_450_v47_ = _dafny.SeqWithoutIsStrInference([d_449_v46_, _dafny.Set({len(d_399_v0_), default__.fm8(688, d_429_v27_, not(d_429_v27_), globalState)})])
                d_429_v27_ = ((d_450_v47_)[default__.safeIndex(d_428_v26_, len(d_450_v47_))]).isdisjoint(_dafny.Set({(d_446_v44_).cardinality}))
            d_451_v48_: _dafny.Array
            nw69_ = _dafny.Array(int(0), 11)
            d_451_v48_ = nw69_
            d_452_v49_: _dafny.Seq
            d_452_v49_ = _dafny.SeqWithoutIsStrInference([d_451_v48_, d_451_v48_, d_451_v48_, d_451_v48_, d_451_v48_])
            rhs40_ = d_429_v27_
            rhs41_ = d_428_v26_
            rhs42_ = not((d_451_v48_) in ((_dafny.SeqWithoutIsStrInference([d_451_v48_, d_451_v48_])) + (d_452_v49_)))
            rhs43_ = d_429_v27_
            rhs44_ = d_429_v27_
            d_429_v27_ = rhs40_
            d_428_v26_ = rhs41_
            d_429_v27_ = rhs42_
            d_429_v27_ = rhs43_
            d_429_v27_ = rhs44_
            d_429_v27_ = not(not(d_429_v27_))
        d_453_v50_: int
        d_453_v50_ = -178
        (globalState).f1 = d_453_v50_
        d_454_v51_: _dafny.Array
        nw70_ = _dafny.Array(False, 20)
        d_454_v51_ = nw70_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_454_v51_).length(0)):
            d_455_i3_: int = guard_loop_1_
            if (True) and (((0) <= (d_455_i3_)) and ((d_455_i3_) < ((d_454_v51_).length(0)))):
                (d_454_v51_)[(d_455_i3_)] = ((_dafny.SeqWithoutIsStrInference([d_453_v50_])) + (_dafny.SeqWithoutIsStrInference([d_453_v50_]))) == ((_dafny.SeqWithoutIsStrInference([d_453_v50_, d_453_v50_, len(_dafny.Set({d_453_v50_, d_453_v50_})), 150])) + (_dafny.SeqWithoutIsStrInference([d_453_v50_ for d_456_i4_ in range(default__.abs(397))])))
        d_457_v52_: _dafny.Array
        nw71_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
        d_457_v52_ = nw71_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_457_v52_).length(0)):
            d_458_i5_: int = guard_loop_2_
            if (True) and (((0) <= (d_458_i5_)) and ((d_458_i5_) < ((d_457_v52_).length(0)))):
                (d_457_v52_)[(d_458_i5_)] = (d_399_v0_).set(default__.safeIndex(d_453_v50_, len(d_399_v0_)), _dafny.CodePoint('r'))
        d_459_v53_: bool
        d_459_v53_ = False
        if (d_459_v53_ if d_459_v53_ else (d_399_v0_) <= (d_399_v0_)):
            d_460_v54_: _dafny.Seq
            d_460_v54_ = _dafny.SeqWithoutIsStrInference([d_459_v53_, d_459_v53_])
            d_461_v55_: _dafny.Set
            d_461_v55_ = _dafny.Set({d_459_v53_, (d_460_v54_)[default__.safeIndex(len(d_460_v54_), len(d_460_v54_))]})
            index46_ = default__.safeIndex(86, (d_457_v52_).length(0))
            (d_457_v52_)[index46_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptockx"))
            index47_ = default__.safeIndex(86, (d_457_v52_).length(0))
            rhs45_ = d_461_v55_
            rhs46_ = d_399_v0_
            lhs21_ = d_457_v52_
            lhs22_ = default__.safeIndex(86, (d_457_v52_).length(0))
            d_461_v55_ = rhs45_
            lhs21_[lhs22_] = rhs46_
            d_462_v56_: _dafny.Array
            nw72_ = _dafny.Array(int(0), 13)
            d_462_v56_ = nw72_
            index48_ = default__.safeIndex(760, (d_462_v56_).length(0))
            (d_462_v56_)[index48_] = (d_453_v50_) + (len((d_457_v52_)[default__.safeIndex(86, (d_457_v52_).length(0))]))
            d_463_v57_: _dafny.Map
            d_463_v57_ = _dafny.Map({d_459_v53_: len(d_461_v55_)})
            index49_ = default__.safeIndex(760, (d_462_v56_).length(0))
            (d_462_v56_)[index49_] = (default__.safeModuloInt(len(d_463_v57_), d_453_v50_) if d_459_v53_ else (d_453_v50_) * (d_453_v50_))
            d_464_v58_: _dafny.Map
            d_464_v58_ = _dafny.Map({(d_462_v56_)[default__.safeIndex(760, (d_462_v56_).length(0))]: not(d_459_v53_)})
            d_465_v59_: _dafny.Seq
            d_465_v59_ = _dafny.SeqWithoutIsStrInference([(d_464_v58_).set(d_453_v50_, d_459_v53_)])
            d_466_v60_: _dafny.Seq
            d_466_v60_ = _dafny.SeqWithoutIsStrInference([d_465_v59_])
            d_467_v61_: _dafny.Seq
            d_467_v61_ = _dafny.SeqWithoutIsStrInference([-869, (d_462_v56_)[default__.safeIndex(760, (d_462_v56_).length(0))]])
            d_468_v62_: str
            d_468_v62_ = _dafny.CodePoint('n')
            index50_ = default__.safeIndex(760, (d_462_v56_).length(0))
            index51_ = default__.safeIndex(760, (d_462_v56_).length(0))
            rhs47_ = d_459_v53_
            rhs48_ = (d_462_v56_)[default__.safeIndex(760, (d_462_v56_).length(0))]
            rhs49_ = (d_466_v60_)[default__.safeIndex(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmbaqsn"))).set(default__.safeIndex((d_467_v61_)[default__.safeIndex(576, len(d_467_v61_))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmbaqsn")))), d_468_v62_)), len(d_466_v60_))]
            rhs50_ = d_459_v53_
            rhs51_ = -502
            lhs23_ = d_462_v56_
            lhs24_ = default__.safeIndex(760, (d_462_v56_).length(0))
            lhs25_ = d_462_v56_
            lhs26_ = default__.safeIndex(760, (d_462_v56_).length(0))
            d_459_v53_ = rhs47_
            lhs23_[lhs24_] = rhs48_
            d_465_v59_ = rhs49_
            d_459_v53_ = rhs50_
            lhs25_[lhs26_] = rhs51_
            d_469_v63_: C0
            nw73_ = C0()
            nw73_.ctor__(d_459_v53_, (d_459_v53_ if d_459_v53_ else False), _dafny.CodePoint('c'), (d_399_v0_) + (d_399_v0_))
            d_469_v63_ = nw73_
            d_468_v62_ = d_468_v62_
        elif True:
            d_470_v64_: _dafny.Array
            nw74_ = _dafny.Array(None, 2)
            nw74_[int(0)] = d_454_v51_
            nw74_[int(1)] = (self).f14
            d_470_v64_ = nw74_
            index52_ = default__.safeIndex(590, (d_470_v64_).length(0))
            (d_470_v64_)[index52_] = (self).f14
            index53_ = default__.safeIndex(590, (d_470_v64_).length(0))
            (d_470_v64_)[index53_] = d_454_v51_
            d_471_v65_: str
            d_471_v65_ = _dafny.CodePoint('p')
            d_472_v66_: C0
            nw75_ = C0()
            nw75_.ctor__(not(True), d_459_v53_, d_471_v65_, d_399_v0_)
            d_472_v66_ = nw75_
            d_473_v67_: _dafny.Map
            d_473_v67_ = _dafny.Map({(d_459_v53_) and (d_459_v53_): d_459_v53_})
            d_473_v67_ = (d_473_v67_) | (d_473_v67_)
            d_459_v53_ = (default__.fm7(d_471_v65_, d_453_v50_, d_459_v53_, globalState) if not(d_459_v53_) else d_459_v53_)
            d_471_v65_ = d_471_v65_
        r0 = d_453_v50_
        d_474_v68_: _dafny.Array
        nw76_ = _dafny.Array(_dafny.Seq({}), 8)
        d_474_v68_ = nw76_
        d_475_v69_: D1
        d_475_v69_ = D1_DC3(d_474_v68_)
        r1 = d_475_v69_
        d_476_v71_: D3
        d_476_v71_ = D3_DC10(d_399_v0_)
        d_477_v72_: _dafny.Map
        d_477_v72_ = _dafny.Map({(d_476_v71_).cf19: d_459_v53_})
        d_478_v73_: _dafny.Seq
        d_478_v73_ = _dafny.SeqWithoutIsStrInference([d_453_v50_])
        d_479_v74_: _dafny.Array
        nw77_ = _dafny.Array(None, 16)
        nw77_[int(0)] = (d_453_v50_) - (d_453_v50_)
        nw77_[int(1)] = d_453_v50_
        nw77_[int(2)] = d_453_v50_
        def iife37_():
            coll15_ = _dafny.Map()
            compr_15_: _dafny.Seq
            for compr_15_ in (d_477_v72_).keys.Elements:
                d_480_v70_: _dafny.Seq = compr_15_
                if (d_480_v70_) in (d_477_v72_):
                    coll15_[d_480_v70_] = d_453_v50_
            return _dafny.Map(coll15_)
        nw77_[int(3)] = len(iife37_()
        )
        nw77_[int(4)] = default__.safeDivisionInt(d_453_v50_, 719)
        nw77_[int(5)] = d_453_v50_
        nw77_[int(6)] = (233 if d_459_v53_ else d_453_v50_)
        nw77_[int(7)] = d_453_v50_
        nw77_[int(8)] = (d_478_v73_)[default__.safeIndex(d_453_v50_, len(d_478_v73_))]
        nw77_[int(9)] = d_453_v50_
        nw77_[int(10)] = (d_453_v50_ if d_459_v53_ else d_453_v50_)
        nw77_[int(11)] = len(d_399_v0_)
        nw77_[int(12)] = d_453_v50_
        nw77_[int(13)] = (d_478_v73_)[default__.safeIndex(854, len(d_478_v73_))]
        nw77_[int(14)] = d_453_v50_
        nw77_[int(15)] = len(d_399_v0_)
        d_479_v74_ = nw77_
        r2 = d_479_v74_
        pat_let_tv8_ = d_474_v68_
        def iife38_(_pat_let11_0):
            def iife39_(d_481_dt__update__tmp_h0_):
                def iife40_(_pat_let12_0):
                    def iife41_(d_482_dt__update_hcf7_h0_):
                        return D1_DC3(d_482_dt__update_hcf7_h0_)
                    return iife41_(_pat_let12_0)
                return iife40_(pat_let_tv8_)
            return iife39_(_pat_let11_0)
        r3 = iife38_(d_475_v69_)
        return r0, r1, r2, r3

    def m7(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: int = int(0)
        d_483_v0_: bool
        d_483_v0_ = False
        d_484_v1_: int
        d_484_v1_ = 64
        d_485_v2_: D0
        d_485_v2_ = D0_DC1(d_483_v0_, (d_484_v1_ if d_483_v0_ else d_484_v1_), d_483_v0_, d_484_v1_)
        source6_ = d_485_v2_
        if source6_.is_DC0:
            if (d_484_v1_) == (d_484_v1_):
                d_486_v3_: _dafny.MultiSet
                d_486_v3_ = _dafny.MultiSet([d_484_v1_])
                d_487_v4_: _dafny.MultiSet
                d_487_v4_ = _dafny.MultiSet([d_484_v1_, (d_486_v3_).cardinality, default__.safeModuloInt(d_484_v1_, default__.fm8(-351, d_483_v0_, d_483_v0_, globalState))])
                d_486_v3_ = d_486_v3_
                r2 = d_484_v1_
                d_488_v5_: _dafny.Seq
                d_488_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcmororjs"))
                d_489_v6_: C0
                nw78_ = C0()
                nw78_.ctor__(d_483_v0_, d_483_v0_, _dafny.CodePoint('q'), d_488_v5_)
                d_489_v6_ = nw78_
                r1 = (d_483_v0_ if True else True)
                d_489_v6_ = d_489_v6_
            elif True:
                r2 = default__.safeModuloInt(d_484_v1_, default__.fm8(d_484_v1_, d_483_v0_, d_483_v0_, globalState))
                d_490_v7_: _dafny.Array
                nw79_ = _dafny.Array(int(0), 2)
                d_490_v7_ = nw79_
                def lambda19_(d_491_i0_):
                    return default__.safeDivisionInt(d_491_i0_, 696)

                init10_ = lambda19_
                nw80_ = _dafny.Array(None, 20)
                for i0_10_ in range(nw80_.length(0)):
                    nw80_[i0_10_] = init10_(i0_10_)
                d_490_v7_ = nw80_
                d_492_v8_: _dafny.MultiSet
                d_492_v8_ = _dafny.MultiSet([not (d_483_v0_) or (d_483_v0_)])
                d_492_v8_ = d_492_v8_
                d_493_v9_: int
                d_494_v10_: D1
                d_495_v11_: _dafny.Array
                d_496_v12_: D1
                out17_: int
                out18_: D1
                out19_: _dafny.Array
                out20_: D1
                out17_, out18_, out19_, out20_ = (self).m6(globalState)
                d_493_v9_ = out17_
                d_494_v10_ = out18_
                d_495_v11_ = out19_
                d_496_v12_ = out20_
                d_483_v0_ = (d_493_v9_) == (d_484_v1_)
            d_497_v13_: _dafny.Seq
            d_497_v13_ = _dafny.SeqWithoutIsStrInference([(p0).cf8])
            d_498_v14_: _dafny.Seq
            d_498_v14_ = _dafny.SeqWithoutIsStrInference([d_483_v0_, ((d_497_v13_)[default__.safeIndex(d_484_v1_, len(d_497_v13_))]) in (d_497_v13_)])
            d_483_v0_ = (d_498_v14_)[default__.safeIndex(d_484_v1_, len(d_498_v14_))]
            (globalState).f1 = d_484_v1_
            d_499_v15_: D3
            d_499_v15_ = D3_DC10((d_497_v13_) + (d_497_v13_))
            d_499_v15_ = D3_DC10(d_497_v13_)
        elif source6_.is_DC1:
            d_500___mcc_h0_ = source6_.cf0
            d_501___mcc_h1_ = source6_.cf1
            d_502___mcc_h2_ = source6_.cf2
            d_503___mcc_h3_ = source6_.cf3
            d_504_cf3_ = d_503___mcc_h3_
            d_505_cf2_ = d_502___mcc_h2_
            d_506_cf1_ = d_501___mcc_h1_
            d_507_cf0_ = d_500___mcc_h0_
            d_508_v16_: _dafny.Set
            d_508_v16_ = _dafny.Set({d_483_v0_})
            d_509_v17_: str
            d_509_v17_ = _dafny.CodePoint('i')
            d_510_v18_: _dafny.Map
            d_510_v18_ = _dafny.Map({d_484_v1_: d_504_cf3_})
            d_511_v19_: _dafny.Set
            d_511_v19_ = _dafny.Set({d_506_cf1_, len(d_510_v18_)})
            r0 = len(((d_508_v16_ if d_505_cf2_ else d_508_v16_)).intersection((_dafny.Set({default__.fm7(d_509_v17_, len(d_511_v19_), d_507_cf0_, globalState), d_505_cf2_})) | (d_508_v16_)))
            r1 = (d_506_cf1_) == ((0) - (d_484_v1_))
            d_512_v20_: _dafny.Array
            nw81_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
            d_512_v20_ = nw81_
            index54_ = default__.safeIndex(35, (d_512_v20_).length(0))
            (d_512_v20_)[index54_] = _dafny.SeqWithoutIsStrInference([d_509_v17_ for d_513_i1_ in range(default__.abs(961))])
            d_514_v21_: _dafny.Seq
            d_514_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elp"))
            index55_ = default__.safeIndex(35, (d_512_v20_).length(0))
            (d_512_v20_)[index55_] = ((d_514_v21_) + (d_514_v21_)) + (d_514_v21_)
            if d_505_cf2_:
                d_515_v22_: C0
                nw82_ = C0()
                nw82_.ctor__(d_507_cf0_, d_483_v0_, d_509_v17_, _dafny.SeqWithoutIsStrInference([d_509_v17_ for d_516_i2_ in range(default__.abs(792))]))
                d_515_v22_ = nw82_
                d_517_v23_: D2
                d_517_v23_ = D2_DC7(d_505_cf2_)
                d_518_v24_: _dafny.Map
                d_518_v24_ = _dafny.Map({d_515_v22_: (self).f14})
                d_519_v25_: _dafny.Array
                nw83_ = _dafny.Array(_dafny.Seq({}), 2)
                d_519_v25_ = nw83_
                d_520_v26_: _dafny.Seq
                d_520_v26_ = _dafny.SeqWithoutIsStrInference([d_507_cf0_])
                index56_ = default__.safeIndex(410, (d_519_v25_).length(0))
                (d_519_v25_)[index56_] = d_520_v26_
                index57_ = default__.safeIndex(410, (d_519_v25_).length(0))
                rhs52_ = (d_515_v22_ if (966) == (d_484_v1_) else d_515_v22_)
                rhs53_ = default__.fm11(default__.fm12(d_507_cf0_, d_514_v21_, d_505_cf2_, globalState), d_505_cf2_, ((0) - (d_484_v1_)) - (810), d_483_v0_, globalState)
                rhs54_ = d_518_v24_
                rhs55_ = (d_520_v26_).set(default__.safeIndex(d_504_cf3_, len(d_520_v26_)), d_507_cf0_)
                lhs27_ = d_519_v25_
                lhs28_ = default__.safeIndex(410, (d_519_v25_).length(0))
                d_515_v22_ = rhs52_
                d_517_v23_ = rhs53_
                d_518_v24_ = rhs54_
                lhs27_[lhs28_] = rhs55_
                d_507_cf0_ = d_483_v0_
                d_521_v27_: _dafny.Seq
                d_521_v27_ = _dafny.SeqWithoutIsStrInference([d_506_cf1_])
                d_522_v28_: _dafny.Map
                d_522_v28_ = _dafny.Map({d_483_v0_: d_507_cf0_})
                d_523_v29_: _dafny.Set
                d_523_v29_ = _dafny.Set({(d_512_v20_)[default__.safeIndex(35, (d_512_v20_).length(0))]})
                d_524_v30_: _dafny.Array
                nw84_ = _dafny.Array(None, 21)
                nw84_[int(0)] = d_506_cf1_
                nw84_[int(1)] = ((d_521_v27_)[default__.safeIndex(d_504_cf3_, len(d_521_v27_))]) + (d_504_cf3_)
                nw84_[int(2)] = d_506_cf1_
                nw84_[int(3)] = d_504_cf3_
                nw84_[int(4)] = d_504_cf3_
                nw84_[int(5)] = len(d_514_v21_)
                nw84_[int(6)] = d_484_v1_
                nw84_[int(7)] = (d_484_v1_ if ((d_522_v28_)[d_483_v0_] if (d_483_v0_) in (d_522_v28_) else d_507_cf0_) else d_506_cf1_)
                nw84_[int(8)] = len(d_523_v29_)
                nw84_[int(9)] = (d_506_cf1_) * (d_506_cf1_)
                nw84_[int(10)] = default__.safeDivisionInt(d_504_cf3_, d_504_cf3_)
                nw84_[int(11)] = d_506_cf1_
                nw84_[int(12)] = (d_484_v1_ if d_505_cf2_ else len(d_508_v16_))
                nw84_[int(13)] = d_506_cf1_
                nw84_[int(14)] = d_504_cf3_
                nw84_[int(15)] = (d_506_cf1_) - (d_484_v1_)
                nw84_[int(16)] = (0) - (d_506_cf1_)
                nw84_[int(17)] = default__.fm8((0) - (d_506_cf1_), d_483_v0_, d_483_v0_, globalState)
                nw84_[int(18)] = (_dafny.MultiSet([False])).cardinality
                nw84_[int(19)] = d_506_cf1_
                nw84_[int(20)] = d_504_cf3_
                d_524_v30_ = nw84_
                index58_ = default__.safeIndex(193, (d_524_v30_).length(0))
                (d_524_v30_)[index58_] = d_504_cf3_
                d_525_v31_: _dafny.MultiSet
                d_525_v31_ = _dafny.MultiSet([d_509_v17_])
                d_526_v32_: _dafny.Map
                d_526_v32_ = _dafny.Map({d_509_v17_: d_505_cf2_})
                index59_ = default__.safeIndex(193, (d_524_v30_).length(0))
                rhs56_ = default__.fm7(d_509_v17_, ((d_525_v31_)[(D1_DC4(_dafny.CodePoint('j'), d_484_v1_, d_509_v17_, d_504_cf3_, True)).cf10] if ((D1_DC4(_dafny.CodePoint('j'), d_484_v1_, d_509_v17_, d_504_cf3_, True)).cf10) in (d_525_v31_) else 204), not((d_484_v1_) > (d_506_cf1_)), globalState)
                rhs57_ = ((len(d_526_v32_)) + (d_504_cf3_)) >= (d_504_cf3_)
                rhs58_ = (0) - ((d_521_v27_)[default__.safeIndex(d_484_v1_, len(d_521_v27_))])
                rhs59_ = d_504_cf3_
                lhs29_ = d_524_v30_
                lhs30_ = default__.safeIndex(193, (d_524_v30_).length(0))
                r1 = rhs56_
                d_483_v0_ = rhs57_
                lhs29_[lhs30_] = rhs58_
                r2 = rhs59_
                index60_ = default__.safeIndex(193, (d_524_v30_).length(0))
                (d_524_v30_)[index60_] = d_506_cf1_
                d_484_v1_ = default__.safeModuloInt(len(((d_512_v20_)[default__.safeIndex(35, (d_512_v20_).length(0))]) + (d_514_v21_)), d_484_v1_)
            elif True:
                index61_ = default__.safeIndex(35, (d_512_v20_).length(0))
                (d_512_v20_)[index61_] = (d_512_v20_)[default__.safeIndex(35, (d_512_v20_).length(0))]
                r1 = d_483_v0_
                d_514_v21_ = (d_514_v21_) + ((d_514_v21_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_504_cf3_])), len(d_514_v21_)), d_509_v17_))
                d_527_v33_: C0
                nw85_ = C0()
                nw85_.ctor__(d_507_cf0_, (d_504_cf3_) == (default__.fm8(d_504_cf3_, d_505_cf2_, not(d_505_cf2_), globalState)), d_509_v17_, (d_512_v20_)[default__.safeIndex(35, (d_512_v20_).length(0))])
                d_527_v33_ = nw85_
                nw86_ = C0()
                nw86_.ctor__(d_483_v0_, d_483_v0_, default__.fm13(globalState), default__.fm14(d_484_v1_, d_504_cf3_, d_507_cf0_, True, globalState))
                d_527_v33_ = nw86_
                d_528_v34_: D3
                d_528_v34_ = D3_DC10(d_514_v21_)
                d_529_v35_: D3
                d_529_v35_ = D3_DC12(d_528_v34_)
                d_530_v36_: _dafny.MultiSet
                d_530_v36_ = _dafny.MultiSet([d_504_cf3_])
                d_531_v37_: _dafny.Seq
                d_531_v37_ = _dafny.SeqWithoutIsStrInference([d_505_cf2_])
                d_532_v38_: _dafny.Map
                d_532_v38_ = _dafny.Map({d_530_v36_: d_531_v37_})
                d_533_v39_: _dafny.Map
                d_533_v39_ = _dafny.Map({d_529_v35_: d_532_v38_})
                d_533_v39_ = (d_533_v39_).set(D3_DC12(d_528_v34_), d_532_v38_)
        elif True:
            d_534___mcc_h4_ = source6_.cf4
            d_535___mcc_h5_ = source6_.cf5
            d_536___mcc_h6_ = source6_.cf6
            d_537_cf6_ = d_536___mcc_h6_
            d_538_cf5_ = d_535___mcc_h5_
            d_539_cf4_ = d_534___mcc_h4_
            d_485_v2_ = d_485_v2_
            d_483_v0_ = (619) > (d_539_cf4_)
            (globalState).f1 = (0) - (d_539_cf4_)
            default__.m0(globalState)
        r1 = d_483_v0_
        d_540_v40_: str
        d_540_v40_ = _dafny.CodePoint('f')
        d_541_v41_: _dafny.Map
        d_541_v41_ = _dafny.Map({d_483_v0_: d_540_v40_})
        d_542_v42_: C0
        nw87_ = C0()
        nw87_.ctor__(True, not(d_483_v0_), ((d_541_v41_)[d_483_v0_] if (d_483_v0_) in (d_541_v41_) else d_540_v40_), default__.fm14(-567, d_484_v1_, default__.fm7(d_540_v40_, d_484_v1_, d_483_v0_, globalState), d_483_v0_, globalState))
        d_542_v42_ = nw87_
        d_484_v1_ = d_484_v1_
        if (d_483_v0_) and (d_483_v0_):
            d_543_v43_: _dafny.Map
            d_543_v43_ = _dafny.Map({not(d_483_v0_): default__.fm15(globalState)})
            d_543_v43_ = (d_543_v43_) | (d_543_v43_)
            if d_483_v0_:
                d_544_v44_: _dafny.Set
                d_544_v44_ = _dafny.Set({True, True})
                r2 = (d_484_v1_) * (default__.safeModuloInt(len(d_544_v44_), d_484_v1_))
                r0 = d_484_v1_
                (globalState).f1 = ((0) - (d_484_v1_)) - (d_484_v1_)
                d_545_v45_: D1
                d_545_v45_ = D1_DC5(d_484_v1_)
                pat_let_tv9_ = d_484_v1_
                pat_let_tv10_ = d_484_v1_
                d_546_v46_: _dafny.Seq
                def iife42_(_pat_let13_0):
                    def iife43_(d_547_dt__update__tmp_h0_):
                        def iife44_(_pat_let14_0):
                            def iife45_(d_548_dt__update_hcf13_h0_):
                                return D1_DC5(d_548_dt__update_hcf13_h0_)
                            return iife45_(_pat_let14_0)
                        return iife44_(pat_let_tv9_)
                    return iife43_(_pat_let13_0)
                def iife46_(_pat_let15_0):
                    def iife47_(d_549_dt__update__tmp_h1_):
                        def iife48_(_pat_let16_0):
                            def iife49_(d_550_dt__update_hcf13_h1_):
                                return D1_DC5(d_550_dt__update_hcf13_h1_)
                            return iife49_(_pat_let16_0)
                        return iife48_(pat_let_tv10_)
                    return iife47_(_pat_let15_0)
                d_546_v46_ = _dafny.SeqWithoutIsStrInference([d_545_v45_, iife42_(d_545_v45_), iife46_(d_545_v45_)])
                d_551_v47_: _dafny.Seq
                d_551_v47_ = default__.fm16(globalState)
                d_552_v48_: _dafny.Map
                d_552_v48_ = _dafny.Map({d_484_v1_: (d_551_v47_)})
                d_553_v49_: _dafny.Seq
                d_553_v49_ = _dafny.SeqWithoutIsStrInference([default__.fm16(globalState), d_546_v46_])
                d_554_v50_: _dafny.Seq
                d_554_v50_ = _dafny.SeqWithoutIsStrInference([(d_546_v46_) + (_dafny.SeqWithoutIsStrInference([d_545_v45_])), ((d_552_v48_)[d_484_v1_] if (d_484_v1_) in (d_552_v48_) else d_546_v46_), (d_553_v49_)[default__.safeIndex(d_484_v1_, len(d_553_v49_))]])
                d_554_v50_ = d_553_v49_
                d_555_v51_: _dafny.Array
                nw88_ = _dafny.Array(int(0), 23)
                d_555_v51_ = nw88_
                index62_ = default__.safeIndex(742, (d_555_v51_).length(0))
                (d_555_v51_)[index62_] = d_484_v1_
                d_556_v52_: _dafny.Map
                d_556_v52_ = _dafny.Map({d_540_v40_: d_483_v0_})
                d_557_v53_: _dafny.Map
                d_557_v53_ = _dafny.Map({d_483_v0_: d_483_v0_})
                index63_ = default__.safeIndex(742, (d_555_v51_).length(0))
                rhs60_ = d_484_v1_
                rhs61_ = default__.fm17(d_484_v1_, globalState)
                rhs62_ = _dafny.CodePoint('j')
                rhs63_ = ((d_484_v1_) - (len(d_557_v53_))) == (d_484_v1_)
                rhs64_ = d_484_v1_
                lhs31_ = d_555_v51_
                lhs32_ = default__.safeIndex(742, (d_555_v51_).length(0))
                lhs33_ = globalState
                lhs31_[lhs32_] = rhs60_
                d_556_v52_ = rhs61_
                d_540_v40_ = rhs62_
                d_483_v0_ = rhs63_
                lhs33_.f1 = rhs64_
            elif True:
                r1 = d_483_v0_
                rhs65_ = d_484_v1_
                rhs66_ = d_484_v1_
                lhs34_ = globalState
                lhs35_ = globalState
                lhs34_.f1 = rhs65_
                lhs35_.f1 = rhs66_
                d_558_v54_: _dafny.Seq
                d_558_v54_ = _dafny.SeqWithoutIsStrInference([d_483_v0_])
                d_559_v55_: _dafny.Seq
                d_559_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kossprq"))
                d_560_v56_: _dafny.Map
                d_560_v56_ = _dafny.Map({default__.fm7(d_540_v40_, len(d_558_v54_), not(d_483_v0_), globalState): d_559_v55_})
                d_560_v56_ = (d_560_v56_).set(d_483_v0_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhcrvh"))) + (d_559_v55_))
                r2 = d_484_v1_
                d_561_v57_: _dafny.Array
                nw89_ = _dafny.Array(None, 5)
                nw89_[int(0)] = -814
                nw89_[int(1)] = (d_484_v1_) - (d_484_v1_)
                nw89_[int(2)] = d_484_v1_
                nw89_[int(3)] = d_484_v1_
                nw89_[int(4)] = 208
                d_561_v57_ = nw89_
                d_562_v58_: D5
                d_562_v58_ = D5_DC14(d_561_v57_)
                d_561_v57_ = (d_562_v58_).cf25
            r1 = not(True)
            r1 = d_483_v0_
            d_563_v59_: _dafny.Array
            nw90_ = _dafny.Array(_dafny.Seq({}), 19)
            d_563_v59_ = nw90_
            d_564_v60_: D1
            d_564_v60_ = D1_DC3(d_563_v59_)
            source7_ = d_564_v60_
            if source7_.is_DC4:
                d_565___mcc_h7_ = source7_.cf8
                d_566___mcc_h8_ = source7_.cf9
                d_567___mcc_h9_ = source7_.cf10
                d_568___mcc_h10_ = source7_.cf11
                d_569___mcc_h11_ = source7_.cf12
                d_570_cf12_ = d_569___mcc_h11_
                d_571_cf11_ = d_568___mcc_h10_
                d_572_cf10_ = d_567___mcc_h9_
                d_573_cf9_ = d_566___mcc_h8_
                d_574_cf8_ = d_565___mcc_h7_
                d_575_v61_: _dafny.Array
                nw91_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
                d_575_v61_ = nw91_
                d_576_v62_: _dafny.Seq
                d_576_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))
                d_577_v63_: _dafny.Map
                d_577_v63_ = _dafny.Map({d_576_v62_: d_570_cf12_})
                d_578_v64_: _dafny.Map
                d_578_v64_ = _dafny.Map({len(d_576_v62_): d_483_v0_})
                d_579_v65_: _dafny.Map
                d_579_v65_ = _dafny.Map({d_483_v0_: len(d_578_v64_)})
                d_580_v66_: _dafny.MultiSet
                d_580_v66_ = _dafny.MultiSet([d_542_v42_])
                d_581_v67_: _dafny.Seq
                d_581_v67_ = _dafny.SeqWithoutIsStrInference([d_571_cf11_, d_573_cf9_, d_484_v1_, (0) - (d_484_v1_)])
                d_582_v68_: _dafny.Array
                nw92_ = _dafny.Array(None, 22)
                nw92_[int(0)] = d_573_cf9_
                nw92_[int(1)] = d_573_cf9_
                nw92_[int(2)] = d_484_v1_
                nw92_[int(3)] = d_484_v1_
                nw92_[int(4)] = d_571_cf11_
                nw92_[int(5)] = d_571_cf11_
                nw92_[int(6)] = d_571_cf11_
                nw92_[int(7)] = d_571_cf11_
                nw92_[int(8)] = d_571_cf11_
                nw92_[int(9)] = d_573_cf9_
                nw92_[int(10)] = len((d_577_v63_).set(_dafny.SeqWithoutIsStrInference([d_572_cf10_ for d_583_i3_ in range(default__.abs(562))]), d_570_cf12_))
                nw92_[int(11)] = d_573_cf9_
                nw92_[int(12)] = (0) - (len(d_576_v62_))
                nw92_[int(13)] = d_484_v1_
                nw92_[int(14)] = d_571_cf11_
                nw92_[int(15)] = len(d_576_v62_)
                nw92_[int(16)] = d_484_v1_
                nw92_[int(17)] = d_484_v1_
                nw92_[int(18)] = len(d_579_v65_)
                nw92_[int(19)] = (d_580_v66_).cardinality
                nw92_[int(20)] = d_571_cf11_
                nw92_[int(21)] = len(d_581_v67_)
                d_582_v68_ = nw92_
                d_584_v69_: _dafny.Map
                d_584_v69_ = _dafny.Map({d_575_v61_: d_582_v68_})
                d_584_v69_ = (d_584_v69_).set(d_575_v61_, d_582_v68_)
                index64_ = default__.safeIndex(349, (d_582_v68_).length(0))
                (d_582_v68_)[index64_] = d_484_v1_
                index65_ = default__.safeIndex(349, (d_582_v68_).length(0))
                (d_582_v68_)[index65_] = d_484_v1_
                d_572_cf10_ = d_572_cf10_
                d_585_v70_: _dafny.Map
                d_585_v70_ = _dafny.Map({d_484_v1_: (d_484_v1_ if d_483_v0_ else 659)})
                d_585_v70_ = (d_585_v70_).set(d_571_cf11_, (d_484_v1_) * (d_573_cf9_))
            elif source7_.is_DC5:
                d_586___mcc_h12_ = source7_.cf13
                d_587_cf13_ = d_586___mcc_h12_
                d_588_v71_: _dafny.Array
                nw93_ = _dafny.Array(_dafny.MultiSet({}), 28)
                d_588_v71_ = nw93_
                d_588_v71_ = d_588_v71_
                d_589_v72_: _dafny.Map
                d_589_v72_ = _dafny.Map({default__.safeDivisionInt(d_587_cf13_, d_587_cf13_): (self).f14})
                d_590_v73_: _dafny.Seq
                d_590_v73_ = _dafny.SeqWithoutIsStrInference([not(d_483_v0_), not(d_483_v0_)])
                d_591_v74_: _dafny.Seq
                d_591_v74_ = _dafny.SeqWithoutIsStrInference([954, len(d_590_v73_)])
                d_592_v75_: _dafny.Map
                d_592_v75_ = _dafny.Map({d_540_v40_: d_483_v0_})
                d_593_v76_: D3
                d_593_v76_ = D3_DC11(d_483_v0_, len(d_591_v74_), _dafny.Set({d_483_v0_, ((d_592_v75_)[d_540_v40_] if (d_540_v40_) in (d_592_v75_) else d_483_v0_)}))
                d_589_v72_ = (d_589_v72_).set((d_593_v76_).cf21, (self).f14)
                d_594_v77_: _dafny.Array
                def lambda20_(d_595_v40_):
                    def lambda21_(d_596_i4_):
                        return default__.safeModuloInt(d_596_i4_, (0) - (len(_dafny.SeqWithoutIsStrInference([d_595_v40_ for d_597_i5_ in range(default__.abs(169))]))))

                    return lambda21_

                init11_ = lambda20_(d_540_v40_)
                nw94_ = _dafny.Array(None, 4)
                for i0_11_ in range(nw94_.length(0)):
                    nw94_[i0_11_] = init11_(i0_11_)
                d_594_v77_ = nw94_
                d_598_v78_: _dafny.Map
                d_598_v78_ = _dafny.Map({d_594_v77_: d_483_v0_})
                d_598_v78_ = (d_598_v78_).set(d_594_v77_, d_483_v0_)
                d_599_v79_: _dafny.Seq
                d_599_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
                d_599_v79_ = d_599_v79_
            elif True:
                d_600___mcc_h13_ = source7_.cf7
                d_601_cf7_ = d_600___mcc_h13_
                d_602_v80_: _dafny.Seq
                d_602_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egghuib"))
                d_602_v80_ = d_602_v80_
                d_603_v81_: _dafny.Map
                d_603_v81_ = _dafny.Map({len(d_602_v80_): d_483_v0_})
                d_604_v82_: _dafny.Seq
                d_604_v82_ = _dafny.SeqWithoutIsStrInference([d_484_v1_, len(d_603_v81_)])
                d_604_v82_ = d_604_v82_
                d_605_v83_: C0
                nw95_ = C0()
                nw95_.ctor__(d_483_v0_, False, d_540_v40_, d_602_v80_)
                d_605_v83_ = nw95_
                d_605_v83_ = d_605_v83_
                d_483_v0_ = (d_484_v1_) == (d_484_v1_)
        elif True:
            d_606_v84_: _dafny.Map
            d_606_v84_ = _dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([d_484_v1_]))})
            d_607_v85_: _dafny.Seq
            d_607_v85_ = _dafny.SeqWithoutIsStrInference([d_483_v0_, d_483_v0_])
            d_608_v86_: _dafny.Map
            d_608_v86_ = _dafny.Map({d_483_v0_: d_483_v0_})
            d_609_v87_: _dafny.Seq
            d_609_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtb"))
            d_610_v88_: _dafny.MultiSet
            d_610_v88_ = _dafny.MultiSet([d_607_v85_, default__.fm18(d_484_v1_, ((d_608_v86_)[d_483_v0_] if (d_483_v0_) in (d_608_v86_) else True), globalState), (d_607_v85_).set(default__.safeIndex(len(default__.fm18(len(d_609_v87_), d_483_v0_, globalState)), len(d_607_v85_)), False), d_607_v85_, d_607_v85_])
            d_611_v89_: D5
            d_611_v89_ = D5_DC15(((self).f14 if d_483_v0_ else (self).f14), d_483_v0_, d_606_v84_, (d_610_v88_).cardinality)
            d_611_v89_ = d_611_v89_
            d_612_v90_: T2
            nw96_ = C1()
            nw96_.ctor__(d_484_v1_, True, d_483_v0_, True, d_540_v40_, d_609_v87_)
            d_612_v90_ = nw96_
            d_613_v91_: _dafny.Map
            d_613_v91_ = _dafny.Map({d_612_v90_: (d_612_v90_).f7})
            d_614_v92_: _dafny.MultiSet
            d_614_v92_ = _dafny.MultiSet([(d_612_v90_).f7])
            d_615_v93_: _dafny.Array
            nw97_ = _dafny.Array(None, 20)
            nw97_[int(0)] = ((d_613_v91_)[d_612_v90_] if (d_612_v90_) in (d_613_v91_) else True)
            nw97_[int(1)] = d_483_v0_
            nw97_[int(2)] = (d_612_v90_).f6
            nw97_[int(3)] = not(default__.fm7(d_540_v40_, d_484_v1_, not((d_612_v90_).f9), globalState))
            nw97_[int(4)] = True
            nw97_[int(5)] = (d_612_v90_).f7
            nw97_[int(6)] = (d_612_v90_).f9
            nw97_[int(7)] = (d_612_v90_).f9
            nw97_[int(8)] = True
            nw97_[int(9)] = (d_612_v90_).f6
            nw97_[int(10)] = ((d_612_v90_).f5) != (d_609_v87_)
            nw97_[int(11)] = not((d_612_v90_).f7)
            nw97_[int(12)] = (_dafny.MultiSet([(d_612_v90_).f9])).ispropersubset(d_614_v92_)
            nw97_[int(13)] = ((d_612_v90_).f9) == ((d_612_v90_).f9)
            nw97_[int(14)] = d_483_v0_
            nw97_[int(15)] = (d_612_v90_).f6
            nw97_[int(16)] = (d_612_v90_).f7
            nw97_[int(17)] = d_483_v0_
            nw97_[int(18)] = False
            nw97_[int(19)] = not ((d_612_v90_).f9) or ((d_612_v90_).f7)
            d_615_v93_ = nw97_
            d_615_v93_ = (self).f14
            d_609_v87_ = d_609_v87_
            d_616_v94_: _dafny.Map
            d_616_v94_ = _dafny.Map({(d_612_v90_).f8: len(d_608_v86_)})
            d_617_v95_: _dafny.Set
            d_617_v95_ = _dafny.Set({(d_612_v90_).f9, (d_612_v90_).f7, (d_612_v90_).f7})
            d_618_v96_: _dafny.Seq
            d_618_v96_ = _dafny.SeqWithoutIsStrInference([((d_616_v94_)[d_484_v1_] if (d_484_v1_) in (d_616_v94_) else (d_612_v90_).f8), default__.safeModuloInt((d_612_v90_).f8, (d_612_v90_).f8), default__.safeModuloInt((d_612_v90_).f8, len(d_617_v95_)), (0) - (d_484_v1_), (d_612_v90_).f8])
            d_618_v96_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(d_484_v1_, 977)])
            if ((d_612_v90_).f8) == ((d_612_v90_).fm1(d_483_v0_, _dafny.Map({(d_612_v90_).f7: (d_612_v90_).f6}), d_483_v0_, globalState)):
                r1 = (d_612_v90_).f9
                d_619_v97_: _dafny.Seq
                d_619_v97_ = _dafny.SeqWithoutIsStrInference([(d_612_v90_).f5, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nravm")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqwbaqax"))])
                d_620_v98_: C0
                nw98_ = C0()
                nw98_.ctor__(default__.fm7(d_612_v90_.f4, (d_612_v90_).f8, False, globalState), True, d_612_v90_.f4, (d_619_v97_)[default__.safeIndex(276, len(d_619_v97_))])
                d_620_v98_ = nw98_
                d_621_v99_: D1
                d_621_v99_ = D1_DC4(_dafny.CodePoint('n'), (d_612_v90_).f8, d_540_v40_, (0) - ((d_612_v90_).f8), not ((d_612_v90_).f9) or (True))
                d_622_v100_: _dafny.Array
                nw99_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
                d_622_v100_ = nw99_
                index66_ = default__.safeIndex(343, (d_622_v100_).length(0))
                (d_622_v100_)[index66_] = (d_612_v90_).f5
                d_623_v101_: _dafny.Map
                d_623_v101_ = _dafny.Map({(d_612_v90_).f8: (d_612_v90_).f6})
                d_624_v102_: _dafny.Set
                d_624_v102_ = _dafny.Set({d_623_v101_})
                pat_let_tv11_ = d_484_v1_
                pat_let_tv12_ = d_612_v90_
                pat_let_tv13_ = d_612_v90_
                pat_let_tv14_ = d_612_v90_
                index67_ = default__.safeIndex(343, (d_622_v100_).length(0))
                def iife50_(_pat_let17_0):
                    def iife51_(d_625_dt__update__tmp_h2_):
                        def iife52_(_pat_let18_0):
                            def iife53_(d_626_dt__update_hcf9_h0_):
                                def iife54_(_pat_let19_0):
                                    def iife55_(d_627_dt__update_hcf8_h0_):
                                        def iife56_(_pat_let20_0):
                                            def iife57_(d_628_dt__update_hcf10_h0_):
                                                def iife58_(_pat_let21_0):
                                                    def iife59_(d_629_dt__update_hcf11_h0_):
                                                        return D1_DC4(d_627_dt__update_hcf8_h0_, d_626_dt__update_hcf9_h0_, d_628_dt__update_hcf10_h0_, d_629_dt__update_hcf11_h0_, (d_625_dt__update__tmp_h2_).cf12)
                                                    return iife59_(_pat_let21_0)
                                                return iife58_((pat_let_tv14_).f8)
                                            return iife57_(_pat_let20_0)
                                        return iife56_(pat_let_tv13_.f4)
                                    return iife55_(_pat_let19_0)
                                return iife54_(pat_let_tv12_.f4)
                            return iife53_(_pat_let18_0)
                        return iife52_(pat_let_tv11_)
                    return iife51_(_pat_let17_0)
                rhs67_ = not (default__.fm7(d_540_v40_, len(_dafny.Set({(d_612_v90_).f8, (d_612_v90_).f8, len(d_624_v102_)})), d_483_v0_, globalState)) or ((d_612_v90_).f7)
                rhs68_ = d_484_v1_
                rhs69_ = iife50_(p0)
                rhs70_ = (d_609_v87_) + (((d_609_v87_).set(default__.safeIndex((0) - ((0) - ((d_612_v90_).f8)), len(d_609_v87_)), d_540_v40_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agc"))))
                rhs71_ = (0) - (len((d_607_v85_) + (_dafny.SeqWithoutIsStrInference([True, (d_612_v90_).f9]))))
                lhs36_ = d_622_v100_
                lhs37_ = default__.safeIndex(343, (d_622_v100_).length(0))
                lhs38_ = globalState
                d_483_v0_ = rhs67_
                d_484_v1_ = rhs68_
                d_621_v99_ = rhs69_
                lhs36_[lhs37_] = rhs70_
                lhs38_.f1 = rhs71_
                d_630_v103_: C0
                nw100_ = C0()
                nw100_.ctor__((d_612_v90_).f9, (d_612_v90_).f9, d_612_v90_.f4, (d_612_v90_).f5)
                d_630_v103_ = nw100_
                (globalState).f1 = (d_612_v90_).f8
            elif True:
                index68_ = default__.safeIndex(634, (d_615_v93_).length(0))
                (d_615_v93_)[index68_] = (True) not in (_dafny.Map({(d_612_v90_).f7: (d_612_v90_).f8}))
                d_631_v104_: _dafny.Array
                nw101_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_631_v104_ = nw101_
                d_632_v105_: C1
                nw102_ = C1()
                nw102_.ctor__((d_612_v90_).f8, (d_612_v90_).f6, (d_612_v90_).f7, (d_612_v90_).f9, d_612_v90_.f4, (d_612_v90_).f5)
                d_632_v105_ = nw102_
                d_633_v106_: _dafny.Array
                nw103_ = _dafny.Array(None, 7)
                nw103_[int(0)] = d_632_v105_
                nw103_[int(1)] = d_632_v105_
                nw103_[int(2)] = d_632_v105_
                nw103_[int(3)] = d_632_v105_
                nw103_[int(4)] = d_632_v105_
                nw103_[int(5)] = d_632_v105_
                nw103_[int(6)] = d_632_v105_
                d_633_v106_ = nw103_
                index69_ = default__.safeIndex(842, (d_631_v104_).length(0))
                (d_631_v104_)[index69_] = d_633_v106_
                d_634_v107_: _dafny.MultiSet
                d_634_v107_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sd")), (d_612_v90_).f5])
                index70_ = default__.safeIndex(634, (d_615_v93_).length(0))
                index71_ = default__.safeIndex(842, (d_631_v104_).length(0))
                rhs72_ = not(not((d_634_v107_).isdisjoint(d_634_v107_)))
                rhs73_ = (d_612_v90_).f9
                rhs74_ = d_633_v106_
                lhs39_ = d_615_v93_
                lhs40_ = default__.safeIndex(634, (d_615_v93_).length(0))
                lhs41_ = d_631_v104_
                lhs42_ = default__.safeIndex(842, (d_631_v104_).length(0))
                lhs39_[lhs40_] = rhs72_
                d_483_v0_ = rhs73_
                lhs41_[lhs42_] = rhs74_
                r1 = not((d_612_v90_).f6)
                d_483_v0_ = not(not ((True if (d_612_v90_).f7 else (d_612_v90_).f6)) or ((d_612_v90_).f7))
                d_635_v108_: _dafny.Map
                d_635_v108_ = _dafny.Map({(d_617_v95_) | (d_617_v95_): d_612_v90_.f4})
                d_635_v108_ = d_635_v108_
                r2 = (d_612_v90_).f8
        d_636_i6_: int
        d_636_i6_ = 0
        with _dafny.label("3"):
            while not(d_483_v0_):
                with _dafny.c_label("3"):
                    if (d_636_i6_) >= (100):
                        raise _dafny.Break("3")
                    d_636_i6_ = (d_636_i6_) + (1)
                    d_483_v0_ = not(d_483_v0_)
                    d_637_v109_: _dafny.Map
                    d_637_v109_ = _dafny.Map({d_484_v1_: d_483_v0_})
                    d_638_v110_: _dafny.Seq
                    d_638_v110_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvhbgf"))
                    d_639_v111_: C0
                    nw104_ = C0()
                    nw104_.ctor__(d_483_v0_, not (((d_637_v109_)[d_484_v1_] if (d_484_v1_) in (d_637_v109_) else not(d_483_v0_))) or (d_483_v0_), default__.fm13(globalState), d_638_v110_)
                    d_639_v111_ = nw104_
                    index72_ = default__.safeIndex(975, ((self).f14).length(0))
                    ((self).f14)[index72_] = not(d_483_v0_)
                    index73_ = default__.safeIndex(975, ((self).f14).length(0))
                    ((self).f14)[index73_] = d_483_v0_
                    d_640_v112_: _dafny.Map
                    d_640_v112_ = _dafny.Map({True: ((self).f14)[default__.safeIndex(975, ((self).f14).length(0))]})
                    d_641_v113_: C1
                    nw105_ = C1()
                    nw105_.ctor__(873, (not(((d_640_v112_)[d_483_v0_] if (d_483_v0_) in (d_640_v112_) else d_483_v0_)) if ((self).f14)[default__.safeIndex(975, ((self).f14).length(0))] else d_483_v0_), d_483_v0_, ((self).f14)[default__.safeIndex(975, ((self).f14).length(0))], d_540_v40_, d_638_v110_)
                    d_641_v113_ = nw105_
                    pass
            pass
        r0 = d_484_v1_
        r1 = d_483_v0_
        r2 = d_484_v1_
        return r0, r1, r2

    @property
    def f14(self):
        return self._f14

class C3(T2, T1, T0):
    def  __init__(self):
        self._f4: str = _dafny.CodePoint('D')
        self._f8: int = int(0)
        self._f9: bool = False
        self._f6: bool = False
        self._f7: bool = False
        self._f5: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f12: bool = False
        self._f13: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f12, f13, f8, f9, f6, f7, f4, f5):
        (self).f12 = f12
        (self)._f13 = f13
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f4 = f4
        (self)._f5 = f5

    def fm1(self, p0, p1, p2, globalState):
        return (self).f8

    def fm2(self, globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f8 for d_642_i1_ in range(default__.abs(329))])) for d_643_i0_ in range(default__.abs(559))])) + ((D0_DC2((self).f8, not(True), _dafny.SeqWithoutIsStrInference([(self).f8 for d_644_i2_ in range(default__.abs(951))]))).cf6)

    def fm5(self, p0, p1, p2, p3, globalState):
        return (self).f8

    def m1(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_645_v0_: _dafny.Map
        d_645_v0_ = _dafny.Map({(self).f6: self.f12})
        d_646_v1_: C0
        nw106_ = C0()
        nw106_.ctor__((self.f12) not in (d_645_v0_), (self).f7, self.f4, (self).f5)
        d_646_v1_ = nw106_
        d_647_v2_: _dafny.Array
        nw107_ = _dafny.Array(False, 5)
        d_647_v2_ = nw107_
        d_648_v3_: _dafny.Seq
        d_648_v3_ = _dafny.SeqWithoutIsStrInference([(self).f9])
        d_649_v4_: _dafny.Seq
        d_649_v4_ = _dafny.SeqWithoutIsStrInference([d_648_v3_])
        d_650_v5_: _dafny.Array
        nw108_ = _dafny.Array(None, 16)
        nw108_[int(0)] = d_648_v3_
        nw108_[int(1)] = _dafny.SeqWithoutIsStrInference([p0])
        nw108_[int(2)] = _dafny.SeqWithoutIsStrInference([False, p3, (self).f9])
        nw108_[int(3)] = d_648_v3_
        nw108_[int(4)] = d_648_v3_
        nw108_[int(5)] = (d_649_v4_)[default__.safeIndex((self).f8, len(d_649_v4_))]
        nw108_[int(6)] = _dafny.SeqWithoutIsStrInference([(self).f7])
        nw108_[int(7)] = d_648_v3_
        nw108_[int(8)] = d_648_v3_
        nw108_[int(9)] = d_648_v3_
        nw108_[int(10)] = d_648_v3_
        nw108_[int(11)] = d_648_v3_
        nw108_[int(12)] = d_648_v3_
        nw108_[int(13)] = d_648_v3_
        nw108_[int(14)] = d_648_v3_
        nw108_[int(15)] = d_648_v3_
        d_650_v5_ = nw108_
        d_651_v6_: D1
        d_651_v6_ = D1_DC3(d_650_v5_)
        d_652_v7_: _dafny.Array
        nw109_ = _dafny.Array(None, 25)
        nw109_[int(0)] = d_651_v6_
        nw109_[int(1)] = d_651_v6_
        nw109_[int(2)] = d_651_v6_
        nw109_[int(3)] = d_651_v6_
        nw109_[int(4)] = d_651_v6_
        nw109_[int(5)] = d_651_v6_
        nw109_[int(6)] = d_651_v6_
        nw109_[int(7)] = d_651_v6_
        nw109_[int(8)] = d_651_v6_
        nw109_[int(9)] = d_651_v6_
        nw109_[int(10)] = d_651_v6_
        nw109_[int(11)] = d_651_v6_
        nw109_[int(12)] = D1_DC3(d_650_v5_)
        nw109_[int(13)] = d_651_v6_
        nw109_[int(14)] = D1_DC3(d_650_v5_)
        nw109_[int(15)] = d_651_v6_
        nw109_[int(16)] = d_651_v6_
        nw109_[int(17)] = d_651_v6_
        nw109_[int(18)] = D1_DC3(d_650_v5_)
        nw109_[int(19)] = d_651_v6_
        nw109_[int(20)] = d_651_v6_
        nw109_[int(21)] = d_651_v6_
        nw109_[int(22)] = d_651_v6_
        nw109_[int(23)] = d_651_v6_
        nw109_[int(24)] = D1_DC3(d_650_v5_)
        d_652_v7_ = nw109_
        d_653_v8_: C2
        nw110_ = C2()
        nw110_.ctor__(d_647_v2_, d_652_v7_)
        d_653_v8_ = nw110_
        (self).f4 = self.f4
        d_654_v9_: _dafny.Array
        def lambda22_(d_655_v3_):
            def lambda23_(d_656_i0_):
                return (d_656_i0_) - (len(d_655_v3_))

            return lambda23_

        init12_ = lambda22_(d_648_v3_)
        nw111_ = _dafny.Array(None, 7)
        for i0_12_ in range(nw111_.length(0)):
            nw111_[i0_12_] = init12_(i0_12_)
        d_654_v9_ = nw111_
        index74_ = default__.safeIndex(912, (d_654_v9_).length(0))
        (d_654_v9_)[index74_] = (0) - ((self).f8)
        d_657_v10_: _dafny.MultiSet
        d_657_v10_ = _dafny.MultiSet([(p2).cardinality])
        d_658_v11_: _dafny.Map
        d_658_v11_ = _dafny.Map({self.f12: d_657_v10_})
        d_659_v12_: _dafny.MultiSet
        d_659_v12_ = _dafny.MultiSet([((d_658_v11_)[(self).f6] if ((self).f6) in (d_658_v11_) else (d_657_v10_).set((self).f8, default__.abs((self).f8))), d_657_v10_])
        index75_ = default__.safeIndex(912, (d_654_v9_).length(0))
        (d_654_v9_)[index75_] = (0) - ((((self).f8) + ((self).f8) if True else (d_659_v12_).cardinality))
        (self).f12 = (self).f9
        hi0_ = (_dafny.MultiSet([p1])).cardinality
        for d_660_i1_ in range((d_654_v9_)[default__.safeIndex(912, (d_654_v9_).length(0))], hi0_):
            d_661_v13_: C0
            nw112_ = C0()
            nw112_.ctor__(p3, (self).f13, self.f4, (self).f5)
            d_661_v13_ = nw112_
            d_657_v10_ = (_dafny.MultiSet([490, (d_654_v9_)[default__.safeIndex(912, (d_654_v9_).length(0))]])).intersection(d_657_v10_)
            r0 = d_660_i1_
            (self).f12 = p1
        d_662_v14_: _dafny.Set
        d_662_v14_ = _dafny.Set({(d_654_v9_)[default__.safeIndex(912, (d_654_v9_).length(0))], (self).f8})
        r0 = ((362) * ((0) - (len(d_662_v14_)))) + (-419)
        d_663_v15_: _dafny.Map
        d_663_v15_ = _dafny.Map({p1: (self).f8})
        r1 = (len((self).f5) if not(((self).f5) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")))) else default__.safeModuloInt((d_654_v9_)[default__.safeIndex(912, (d_654_v9_).length(0))], ((d_663_v15_)[p0] if (p0) in (d_663_v15_) else len(_dafny.Map({(self).f6: 704})))))
        d_664_v16_: _dafny.Seq
        d_664_v16_ = _dafny.SeqWithoutIsStrInference([d_654_v9_, d_654_v9_])
        d_665_v17_: _dafny.Array
        nw113_ = _dafny.Array(None, 20)
        nw113_[int(0)] = d_654_v9_
        nw113_[int(1)] = d_654_v9_
        nw113_[int(2)] = (d_664_v16_)[default__.safeIndex((d_654_v9_)[default__.safeIndex(912, (d_654_v9_).length(0))], len(d_664_v16_))]
        nw113_[int(3)] = (d_664_v16_)[default__.safeIndex((d_654_v9_)[default__.safeIndex(912, (d_654_v9_).length(0))], len(d_664_v16_))]
        nw113_[int(4)] = d_654_v9_
        nw113_[int(5)] = d_654_v9_
        nw113_[int(6)] = d_654_v9_
        nw113_[int(7)] = d_654_v9_
        nw113_[int(8)] = d_654_v9_
        nw113_[int(9)] = d_654_v9_
        nw113_[int(10)] = d_654_v9_
        nw113_[int(11)] = d_654_v9_
        nw113_[int(12)] = d_654_v9_
        nw113_[int(13)] = d_654_v9_
        nw113_[int(14)] = d_654_v9_
        nw113_[int(15)] = d_654_v9_
        nw113_[int(16)] = (d_654_v9_ if (self).f9 else d_654_v9_)
        nw113_[int(17)] = d_654_v9_
        nw113_[int(18)] = d_654_v9_
        nw113_[int(19)] = d_654_v9_
        d_665_v17_ = nw113_
        r2 = d_665_v17_
        return r0, r1, r2

    def m5(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_666_v0_: _dafny.Map
        d_666_v0_ = _dafny.Map({(self).f8: ((self).f8) + (678)})
        d_666_v0_ = (d_666_v0_).set((self).f8, default__.safeModuloInt((0) - ((self).f8), (self).f8))
        index76_ = default__.safeIndex(570, (p2).length(0))
        (p2)[index76_] = p0
        d_667_v1_: _dafny.Array
        nw114_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
        d_667_v1_ = nw114_
        index77_ = default__.safeIndex(685, (d_667_v1_).length(0))
        (d_667_v1_)[index77_] = ((self).f5) + ((self).f5)
        index78_ = default__.safeIndex(570, (p2).length(0))
        index79_ = default__.safeIndex(685, (d_667_v1_).length(0))
        rhs75_ = (self).f8
        rhs76_ = p0
        rhs77_ = (self).f5
        lhs43_ = globalState
        lhs44_ = p2
        lhs45_ = default__.safeIndex(570, (p2).length(0))
        lhs46_ = d_667_v1_
        lhs47_ = default__.safeIndex(685, (d_667_v1_).length(0))
        lhs43_.f1 = rhs75_
        lhs44_[lhs45_] = rhs76_
        lhs46_[lhs47_] = rhs77_
        d_668_v2_: _dafny.Seq
        d_668_v2_ = _dafny.SeqWithoutIsStrInference([(self).f8, (self).f8])
        d_669_v3_: _dafny.Seq
        d_669_v3_ = _dafny.SeqWithoutIsStrInference([not (not(p0)) or ((self).f6), ((d_668_v2_).set(default__.safeIndex(101, len(d_668_v2_)), (self).f8)) < (d_668_v2_), (self).f13])
        index80_ = default__.safeIndex(570, (p2).length(0))
        (p2)[index80_] = (d_669_v3_)[default__.safeIndex((self).f8, len(d_669_v3_))]
        hi1_ = ((self).f8) + ((self).f8)
        for d_670_i0_ in range((self).f8, hi1_):
            d_669_v3_ = d_669_v3_
            d_671_v4_: _dafny.Array
            nw115_ = _dafny.Array(None, 24)
            d_671_v4_ = nw115_
            d_672_v5_: D0
            d_672_v5_ = D0_DC2((self).f8, default__.fm7(p3, d_670_i0_, False, globalState), _dafny.SeqWithoutIsStrInference([d_670_i0_]))
            d_673_v6_: _dafny.Map
            d_673_v6_ = _dafny.Map({d_672_v5_: (self).f8})
            d_674_v7_: C1
            nw116_ = C1()
            nw116_.ctor__(len(d_673_v6_), False, (self).f13, self.f12, p3, (self).f5)
            d_674_v7_ = nw116_
            index81_ = default__.safeIndex(534, (d_671_v4_).length(0))
            (d_671_v4_)[index81_] = d_674_v7_
            index82_ = default__.safeIndex(534, (d_671_v4_).length(0))
            (d_671_v4_)[index82_] = d_674_v7_
            (globalState).f1 = len(((self).f5).set(default__.safeIndex(((self).f8 if True else len((d_667_v1_)[default__.safeIndex(685, (d_667_v1_).length(0))])), len((self).f5)), p3))
            d_675_v8_: _dafny.Set
            d_675_v8_ = _dafny.Set({d_669_v3_})
            d_676_v9_: _dafny.Set
            d_676_v9_ = _dafny.Set({(self).f6, self.f12, (self).f9})
            d_677_v10_: D3
            d_677_v10_ = D3_DC11((self).f9, len(d_675_v8_), d_676_v9_)
            if not((d_677_v10_).cf20):
                (globalState).f1 = d_670_i0_
                (globalState).f1 = 275
                (globalState).f1 = (d_670_i0_) - (len(d_668_v2_))
                d_678_v11_: _dafny.Array
                nw117_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_678_v11_ = nw117_
                d_679_v12_: _dafny.Array
                nw118_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                d_679_v12_ = nw118_
                index83_ = default__.safeIndex(601, (d_678_v11_).length(0))
                (d_678_v11_)[index83_] = d_679_v12_
                index84_ = default__.safeIndex(601, (d_678_v11_).length(0))
                (d_678_v11_)[index84_] = d_679_v12_
                (globalState).f1 = ((self).f8 if (d_670_i0_) == (d_670_i0_) else (self).f8)
            elif True:
                d_680_v13_: _dafny.Map
                d_680_v13_ = _dafny.Map({((self).f8) - ((self).f8): (self).f5})
                d_681_v14_: _dafny.Array
                nw119_ = _dafny.Array(int(0), 23)
                d_681_v14_ = nw119_
                d_682_v15_: _dafny.MultiSet
                d_682_v15_ = _dafny.MultiSet([d_681_v14_, d_681_v14_, d_681_v14_])
                (globalState).f1 = len(((d_680_v13_)[((d_682_v15_)[d_681_v14_] if (d_681_v14_) in (d_682_v15_) else (self).f8)] if (((d_682_v15_)[d_681_v14_] if (d_681_v14_) in (d_682_v15_) else (self).f8)) in (d_680_v13_) else _dafny.SeqWithoutIsStrInference([self.f4 for d_683_i1_ in range(default__.abs(291))])))
                index85_ = default__.safeIndex(220, (d_681_v14_).length(0))
                (d_681_v14_)[index85_] = default__.safeDivisionInt((d_677_v10_).cf21, default__.fm8(len((self).f5), (self).f9, True, globalState))
                d_684_v16_: _dafny.MultiSet
                d_684_v16_ = _dafny.MultiSet([(self).f13])
                d_685_v17_: D5
                d_685_v17_ = D5_DC15(p2, p0, default__.fm24(d_670_i0_, True, p0, (d_684_v16_).cardinality, globalState), 69)
                d_686_v18_: _dafny.Map
                d_686_v18_ = _dafny.Map({(self).f9: (d_685_v17_).cf29})
                index86_ = default__.safeIndex(220, (d_681_v14_).length(0))
                index87_ = default__.safeIndex(685, (d_667_v1_).length(0))
                rhs78_ = (self).f8
                rhs79_ = (self).f8
                rhs80_ = (d_686_v18_) | ((d_686_v18_) | (d_686_v18_))
                rhs81_ = (0) - (((d_668_v2_)[default__.safeIndex(len((d_667_v1_)[default__.safeIndex(685, (d_667_v1_).length(0))]), len(d_668_v2_))]) + (d_670_i0_))
                rhs82_ = (((d_667_v1_)[default__.safeIndex(685, (d_667_v1_).length(0))]) + ((self).f5)).set(default__.safeIndex(d_670_i0_, len(((d_667_v1_)[default__.safeIndex(685, (d_667_v1_).length(0))]) + ((self).f5))), p3)
                lhs48_ = globalState
                lhs49_ = d_681_v14_
                lhs50_ = default__.safeIndex(220, (d_681_v14_).length(0))
                lhs51_ = globalState
                lhs52_ = d_667_v1_
                lhs53_ = default__.safeIndex(685, (d_667_v1_).length(0))
                lhs48_.f1 = rhs78_
                lhs49_[lhs50_] = rhs79_
                d_686_v18_ = rhs80_
                lhs51_.f1 = rhs81_
                lhs52_[lhs53_] = rhs82_
                d_687_v19_: _dafny.Map
                d_687_v19_ = _dafny.Map({d_670_i0_: p3})
                d_687_v19_ = (d_687_v19_).set(-494, (self.f4 if p0 else p3))
                index88_ = default__.safeIndex(570, (p2).length(0))
                (p2)[index88_] = False
                (globalState).f1 = ((0) - ((self).f8)) * ((d_681_v14_)[default__.safeIndex(220, (d_681_v14_).length(0))])
        d_688_v20_: _dafny.Array
        nw120_ = _dafny.Array(int(0), 26)
        d_688_v20_ = nw120_
        index89_ = default__.safeIndex(470, (d_688_v20_).length(0))
        (d_688_v20_)[index89_] = (self).f8
        d_689_v21_: _dafny.Array
        def lambda24_(d_690_i2_):
            return _dafny.MultiSet([True, self.f12, True, False, not((self).f9)])

        init13_ = lambda24_
        nw121_ = _dafny.Array(None, 24)
        for i0_13_ in range(nw121_.length(0)):
            nw121_[i0_13_] = init13_(i0_13_)
        d_689_v21_ = nw121_
        d_691_v22_: _dafny.MultiSet
        d_691_v22_ = _dafny.MultiSet([(self).f8])
        d_692_v23_: D0
        d_692_v23_ = D0_DC1((self).f9, (d_691_v22_).cardinality, (self).f7, -704)
        index90_ = default__.safeIndex(685, (d_667_v1_).length(0))
        index91_ = default__.safeIndex(570, (p2).length(0))
        index92_ = default__.safeIndex(470, (d_688_v20_).length(0))
        rhs83_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbviiq"))) if ((d_692_v23_).cf1) == ((self).f8) else (self).f5)
        rhs84_ = (self).f13
        rhs85_ = default__.safeModuloInt((self).f8, default__.safeModuloInt((self).f8, (self).f8))
        rhs86_ = d_689_v21_
        lhs54_ = d_667_v1_
        lhs55_ = default__.safeIndex(685, (d_667_v1_).length(0))
        lhs56_ = p2
        lhs57_ = default__.safeIndex(570, (p2).length(0))
        lhs58_ = d_688_v20_
        lhs59_ = default__.safeIndex(470, (d_688_v20_).length(0))
        lhs54_[lhs55_] = rhs83_
        lhs56_[lhs57_] = rhs84_
        lhs58_[lhs59_] = rhs85_
        d_689_v21_ = rhs86_
        d_693_v24_: T2
        nw122_ = C1()
        nw122_.ctor__(((d_666_v0_)[(self).f8] if ((self).f8) in (d_666_v0_) else 444), (self).f13, True, (self).f6, self.f4, (d_667_v1_)[default__.safeIndex(685, (d_667_v1_).length(0))])
        d_693_v24_ = nw122_
        d_694_v25_: _dafny.Map
        d_694_v25_ = _dafny.Map({(p2)[default__.safeIndex(570, (p2).length(0))]: d_693_v24_})
        d_695_v26_: _dafny.Array
        nw123_ = _dafny.Array(None, 14)
        nw123_[int(0)] = d_694_v25_
        nw123_[int(1)] = _dafny.Map({self.f12: d_693_v24_})
        nw123_[int(2)] = _dafny.Map({(d_693_v24_).f6: d_693_v24_})
        nw123_[int(3)] = d_694_v25_
        nw123_[int(4)] = d_694_v25_
        nw123_[int(5)] = d_694_v25_
        nw123_[int(6)] = d_694_v25_
        nw123_[int(7)] = d_694_v25_
        nw123_[int(8)] = d_694_v25_
        nw123_[int(9)] = _dafny.Map({(d_693_v24_).f7: d_693_v24_})
        nw123_[int(10)] = (d_694_v25_).set((self).f9, d_693_v24_)
        nw123_[int(11)] = d_694_v25_
        nw123_[int(12)] = d_694_v25_
        nw123_[int(13)] = (d_694_v25_).set((p2)[default__.safeIndex(570, (p2).length(0))], d_693_v24_)
        d_695_v26_ = nw123_
        d_696_v27_: _dafny.Map
        d_696_v27_ = _dafny.Map({d_695_v26_: (d_669_v3_)[default__.safeIndex((self).f8, len(d_669_v3_))]})
        (self).f12 = ((d_696_v27_)[d_695_v26_] if (d_695_v26_) in (d_696_v27_) else not(p0))
        r0 = d_688_v20_
        return r0

    @property
    def f13(self):
        return self._f13

class C4(T2, T0, T3, T1):
    def  __init__(self):
        self._f4: str = _dafny.CodePoint('D')
        self._f8: int = int(0)
        self._f9: bool = False
        self._f6: bool = False
        self._f7: bool = False
        self._f5: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f8, f9, f6, f7, f4, f5):
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f4 = f4
        (self)._f5 = f5

    def fm1(self, p0, p1, p2, globalState):
        return (self).f8

    def fm2(self, globalState):
        return (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f7])).cardinality])) + ((_dafny.SeqWithoutIsStrInference([(self).f8, (self).f8, (self).f8, (self).f8, (self).f8]) if not(not((self).f9)) else _dafny.SeqWithoutIsStrInference([len((self).f5)])))

    def m1(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_697_v0_: _dafny.Seq
        d_697_v0_ = _dafny.SeqWithoutIsStrInference([(self).f8, (self).f8])
        d_698_v1_: _dafny.MultiSet
        d_698_v1_ = _dafny.MultiSet([len((self).f5)])
        d_699_v2_: _dafny.Seq
        d_699_v2_ = _dafny.SeqWithoutIsStrInference([p3, (_dafny.MultiSet(d_697_v0_)) == (d_698_v1_)])
        d_699_v2_ = d_699_v2_
        d_700_v3_: C1
        nw124_ = C1()
        nw124_.ctor__((self).f8, False, p1, (self).f6, self.f4, (self).f5)
        d_700_v3_ = nw124_
        d_701_v4_: _dafny.Map
        d_701_v4_ = _dafny.Map({d_700_v3_: (self).f6})
        d_702_v5_: _dafny.Map
        d_702_v5_ = _dafny.Map({p3: (self).f8})
        d_703_v6_: _dafny.Map
        d_703_v6_ = _dafny.Map({(self).f8: ((d_702_v5_)[p0] if (p0) in (d_702_v5_) else (self).f8)})
        d_704_v7_: _dafny.MultiSet
        d_704_v7_ = _dafny.MultiSet([_dafny.Map({(self).f8: len(d_701_v4_)}), d_703_v6_])
        def iife60_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in _dafny.IntegerRange(908, 663):
                d_705_v8_: int = compr_16_
                if ((908) <= (d_705_v8_)) and ((d_705_v8_) < (663)):
                    coll16_[(d_705_v8_) + ((self).f8)] = (self).f8
            return _dafny.Map(coll16_)
        def iife61_():
            coll17_ = _dafny.Map()
            compr_17_: int
            for compr_17_ in _dafny.IntegerRange(908, 663):
                d_706_v8_: int = compr_17_
                if ((908) <= (d_706_v8_)) and ((d_706_v8_) < (663)):
                    coll17_[(d_706_v8_) + ((self).f8)] = (self).f8
            return _dafny.Map(coll17_)
        r0 = ((d_704_v7_)[(iife60_()
        ) | (d_703_v6_)] if ((iife61_()
        ) | (d_703_v6_)) in (d_704_v7_) else (self).f8)
        hi2_ = (self).f8
        for d_707_i0_ in range(((self).f8) + ((self).f8), hi2_):
            r1 = (0) - ((((self).f8 if p0 else len((self).f5))) - ((d_707_i0_) * (d_707_i0_)))
            d_708_v9_: _dafny.Array
            nw125_ = _dafny.Array(_dafny.Map({}), 5)
            d_708_v9_ = nw125_
            d_709_v10_: _dafny.Array
            nw126_ = _dafny.Array(False, 1)
            d_709_v10_ = nw126_
            d_710_v11_: D5
            d_710_v11_ = D5_DC15(d_709_v10_, True, d_702_v5_, -268)
            d_711_v12_: _dafny.Array
            nw127_ = _dafny.Array(None, 4)
            nw127_[int(0)] = (self).f6
            nw127_[int(1)] = p1
            nw127_[int(2)] = (d_710_v11_).cf27
            nw127_[int(3)] = True
            d_711_v12_ = nw127_
            d_712_v13_: _dafny.Map
            d_712_v13_ = _dafny.Map({(self).f8: d_709_v10_})
            index93_ = default__.safeIndex(900, (d_708_v9_).length(0))
            (d_708_v9_)[index93_] = d_712_v13_
            index94_ = default__.safeIndex(900, (d_708_v9_).length(0))
            (d_708_v9_)[index94_] = (d_712_v13_).set(-894, d_711_v12_)
            d_713_v14_: _dafny.Array
            nw128_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_713_v14_ = nw128_
            d_714_v15_: _dafny.Array
            nw129_ = _dafny.Array(None, 6)
            nw129_[int(0)] = (self).f8
            nw129_[int(1)] = d_707_i0_
            nw129_[int(2)] = (self).f8
            nw129_[int(3)] = 716
            nw129_[int(4)] = d_707_i0_
            nw129_[int(5)] = (self).f8
            d_714_v15_ = nw129_
            index95_ = default__.safeIndex(244, (d_713_v14_).length(0))
            (d_713_v14_)[index95_] = d_714_v15_
            index96_ = default__.safeIndex(244, (d_713_v14_).length(0))
            rhs87_ = d_714_v15_
            rhs88_ = d_714_v15_
            lhs60_ = d_713_v14_
            lhs61_ = default__.safeIndex(244, (d_713_v14_).length(0))
            lhs60_[lhs61_] = rhs87_
            d_714_v15_ = rhs88_
            (globalState).f1 = (d_707_i0_ if (self).f6 else d_707_i0_)
        d_698_v1_ = ((d_698_v1_).intersection((d_698_v1_).set(94, default__.abs((self).f8)))) - (d_698_v1_)
        (self).f4 = self.f4
        r0 = (self).f8
        r0 = (self).f8
        r1 = (self).f8
        d_715_v16_: _dafny.Array
        nw130_ = _dafny.Array(_dafny.Array(None, 0), 3)
        d_715_v16_ = nw130_
        r2 = d_715_v16_
        return r0, r1, r2

    def m2(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_716_v0_: _dafny.MultiSet
        d_716_v0_ = _dafny.MultiSet([(self).f9, not(True)])
        d_717_v1_: _dafny.Seq
        d_717_v1_ = _dafny.SeqWithoutIsStrInference([(self).f7])
        if ((d_716_v0_) - (_dafny.MultiSet([(self).f9]))).isdisjoint(_dafny.MultiSet(d_717_v1_)):
            default__.m0(globalState)
            d_718_v2_: bool
            d_718_v2_ = False
            d_718_v2_ = (self).f9
            d_719_v3_: _dafny.Array
            nw131_ = _dafny.Array(int(0), 25)
            d_719_v3_ = nw131_
            d_720_v4_: _dafny.Seq
            d_720_v4_ = _dafny.SeqWithoutIsStrInference([(self).f8])
            d_721_v5_: _dafny.Map
            d_721_v5_ = _dafny.Map({self.f4: default__.fm27((self).f8, d_718_v2_, d_718_v2_, globalState)})
            d_722_v6_: C1
            nw132_ = C1()
            nw132_.ctor__(len(d_720_v4_), (self).f9, ((d_721_v5_)[self.f4] if (self.f4) in (d_721_v5_) else (self).f9), not(d_718_v2_), _dafny.CodePoint('i'), (self).f5)
            d_722_v6_ = nw132_
            d_723_v7_: _dafny.Map
            d_723_v7_ = _dafny.Map({d_722_v6_: _dafny.Map({(self).f8: d_718_v2_})})
            d_724_v8_: _dafny.Map
            d_724_v8_ = _dafny.Map({p0: d_718_v2_})
            index97_ = default__.safeIndex(688, (d_719_v3_).length(0))
            (d_719_v3_)[index97_] = (0) - (len(((d_723_v7_)[d_722_v6_] if (d_722_v6_) in (d_723_v7_) else d_724_v8_)))
            index98_ = default__.safeIndex(688, (d_719_v3_).length(0))
            rhs89_ = (0) - ((self).f8)
            rhs90_ = (self).f8
            rhs91_ = not((d_718_v2_ if ((0) - (p0)) < (p0) else (self).f9))
            lhs62_ = d_719_v3_
            lhs63_ = default__.safeIndex(688, (d_719_v3_).length(0))
            lhs64_ = globalState
            lhs62_[lhs63_] = rhs89_
            lhs64_.f1 = rhs90_
            d_718_v2_ = rhs91_
            d_725_v9_: C1
            nw133_ = C1()
            nw133_.ctor__(default__.safeModuloInt((d_719_v3_)[default__.safeIndex(688, (d_719_v3_).length(0))], (self).f8), d_718_v2_, d_718_v2_, (self).f7, self.f4, (self).f5)
            d_725_v9_ = nw133_
            (self).f4 = self.f4
        elif True:
            d_726_v10_: bool
            d_726_v10_ = False
            d_726_v10_ = default__.fm27(p0, (self).f9, ((self).f7) or ((self).f6), globalState)
            d_727_v11_: D1
            d_727_v11_ = D1_DC5((self).f8)
            d_728_v12_: _dafny.Seq
            d_728_v12_ = _dafny.SeqWithoutIsStrInference([d_727_v11_, d_727_v11_])
            d_729_v13_: _dafny.Seq
            d_729_v13_ = d_728_v12_
            source8_ = d_729_v13_
            d_730___mcc_h0_ = source8_
            d_731_cf24_ = d_730___mcc_h0_
            d_732_v14_: _dafny.Seq
            d_732_v14_ = _dafny.SeqWithoutIsStrInference([272])
            d_733_v15_: _dafny.Set
            d_733_v15_ = _dafny.Set({d_732_v14_})
            d_734_v16_: C1
            nw134_ = C1()
            nw134_.ctor__((0) - (p0), not((self).f6), (653) > ((self).f8), default__.fm27(len(d_733_v15_), False, d_726_v10_, globalState), self.f4, ((self).f5) + ((self).f5))
            d_734_v16_ = nw134_
            d_735_v17_: D1
            d_735_v17_ = D1_DC4(self.f4, (self).f8, self.f4, p0, (self).f6)
            d_726_v10_ = (((self).f8) + ((d_735_v17_).cf11)) < (default__.safeModuloInt((self).f8, p0))
            d_736_v18_: _dafny.Seq
            d_736_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "euaygar"))
            d_737_v19_: _dafny.Array
            nw135_ = _dafny.Array(_dafny.MultiSet({}), 5)
            d_737_v19_ = nw135_
            index99_ = default__.safeIndex(227, (d_737_v19_).length(0))
            (d_737_v19_)[index99_] = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pas"))])
            d_738_v20_: _dafny.Array
            nw136_ = _dafny.Array(False, 4)
            d_738_v20_ = nw136_
            d_739_v21_: _dafny.Array
            nw137_ = _dafny.Array(None, 18)
            nw137_[int(0)] = d_738_v20_
            nw137_[int(1)] = d_738_v20_
            nw137_[int(2)] = d_738_v20_
            nw137_[int(3)] = d_738_v20_
            nw137_[int(4)] = d_738_v20_
            nw137_[int(5)] = d_738_v20_
            nw137_[int(6)] = d_738_v20_
            nw137_[int(7)] = d_738_v20_
            nw137_[int(8)] = d_738_v20_
            nw137_[int(9)] = d_738_v20_
            nw137_[int(10)] = d_738_v20_
            nw137_[int(11)] = d_738_v20_
            nw137_[int(12)] = d_738_v20_
            nw137_[int(13)] = d_738_v20_
            nw137_[int(14)] = d_738_v20_
            nw137_[int(15)] = d_738_v20_
            nw137_[int(16)] = (d_738_v20_ if default__.fm27(p0, False, default__.fm27(p0, (d_717_v1_)[default__.safeIndex(635, len(d_717_v1_))], (self).f6, globalState), globalState) else d_738_v20_)
            nw137_[int(17)] = d_738_v20_
            d_739_v21_ = nw137_
            index100_ = default__.safeIndex(0, (d_739_v21_).length(0))
            (d_739_v21_)[index100_] = d_738_v20_
            d_740_v22_: _dafny.Set
            d_740_v22_ = _dafny.Set({(self).f7, (self).f7})
            d_741_v23_: D3
            d_741_v23_ = D3_DC11((self).f7, p0, d_740_v22_)
            index101_ = default__.safeIndex(227, (d_737_v19_).length(0))
            index102_ = default__.safeIndex(0, (d_739_v21_).length(0))
            rhs92_ = (False) == ((self).f9)
            rhs93_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxeov"))) + (default__.fm28(len(d_732_v14_), d_726_v10_, d_717_v1_, D3_DC11(False, (self).f8, d_740_v22_), globalState))
            rhs94_ = _dafny.MultiSet([(default__.fm28((self).f8, (self).f9, d_717_v1_, d_741_v23_, globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nservh")))])
            rhs95_ = d_738_v20_
            lhs65_ = d_737_v19_
            lhs66_ = default__.safeIndex(227, (d_737_v19_).length(0))
            lhs67_ = d_739_v21_
            lhs68_ = default__.safeIndex(0, (d_739_v21_).length(0))
            d_726_v10_ = rhs92_
            d_736_v18_ = rhs93_
            lhs65_[lhs66_] = rhs94_
            lhs67_[lhs68_] = rhs95_
            d_742_v24_: _dafny.Array
            nw138_ = _dafny.Array(int(0), 9)
            d_742_v24_ = nw138_
            d_743_v25_: _dafny.Map
            d_743_v25_ = _dafny.Map({d_738_v20_: d_742_v24_})
            d_744_v26_: _dafny.MultiSet
            d_744_v26_ = _dafny.MultiSet([(self).f8])
            d_745_v27_: _dafny.Map
            d_745_v27_ = _dafny.Map({((d_743_v25_)[d_738_v20_] if (d_738_v20_) in (d_743_v25_) else d_742_v24_): d_744_v26_})
            d_746_v28_: C3
            nw139_ = C3()
            nw139_.ctor__(False, d_726_v10_, (0) - (len(d_745_v27_)), (self).f7, (self).f7, (self).f7, self.f4, _dafny.SeqWithoutIsStrInference([self.f4 for d_747_i0_ in range(default__.abs(657))]))
            d_746_v28_ = nw139_
            d_748_v29_: _dafny.Array
            def lambda25_(d_749_i1_):
                return False

            init14_ = lambda25_
            nw140_ = _dafny.Array(None, 14)
            for i0_14_ in range(nw140_.length(0)):
                nw140_[i0_14_] = init14_(i0_14_)
            d_748_v29_ = nw140_
            d_750_v30_: _dafny.Set
            d_750_v30_ = _dafny.Set({d_748_v29_})
            d_751_v31_: D7
            d_751_v31_ = D7_DC18(d_750_v30_)
            rhs96_ = (self).f7
            rhs97_ = not (d_726_v10_) or ((self).f7)
            rhs98_ = (d_751_v31_).cf33
            rhs99_ = (0) - (p0)
            lhs69_ = globalState
            d_726_v10_ = rhs96_
            d_726_v10_ = rhs97_
            d_750_v30_ = rhs98_
            lhs69_.f1 = rhs99_
            d_752_v32_: _dafny.Array
            nw141_ = _dafny.Array(int(0), 26)
            d_752_v32_ = nw141_
            index103_ = default__.safeIndex(971, (d_752_v32_).length(0))
            (d_752_v32_)[index103_] = (0) - (p0)
            index104_ = default__.safeIndex(971, (d_752_v32_).length(0))
            (d_752_v32_)[index104_] = (0) - (p0)
            index105_ = default__.safeIndex(971, (d_752_v32_).length(0))
            (d_752_v32_)[index105_] = (self).f8
        hi3_ = default__.safeDivisionInt((0) - (p0), p0)
        for d_753_i2_ in range((0) - ((self).f8), hi3_):
            default__.m0(globalState)
            d_754_v33_: _dafny.Seq
            d_754_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mm"))
            d_754_v33_ = (d_754_v33_).set(default__.safeIndex(default__.safeModuloInt(p0, p0), len(d_754_v33_)), self.f4)
            d_755_v34_: _dafny.Array
            nw142_ = _dafny.Array(None, 7)
            nw142_[int(0)] = (self).f7
            nw142_[int(1)] = True
            nw142_[int(2)] = (self).f6
            nw142_[int(3)] = default__.fm27(len((self).f5), (self).f9, (self).f6, globalState)
            nw142_[int(4)] = (self).f9
            nw142_[int(5)] = (self).f6
            nw142_[int(6)] = (self).f9
            d_755_v34_ = nw142_
            d_756_v35_: _dafny.Array
            nw143_ = _dafny.Array(None, 26)
            nw143_[int(0)] = _dafny.SeqWithoutIsStrInference([(self).f7, False])
            nw143_[int(1)] = _dafny.SeqWithoutIsStrInference([(self).f7])
            nw143_[int(2)] = d_717_v1_
            nw143_[int(3)] = d_717_v1_
            nw143_[int(4)] = d_717_v1_
            nw143_[int(5)] = d_717_v1_
            nw143_[int(6)] = _dafny.SeqWithoutIsStrInference([(self).f9])
            nw143_[int(7)] = _dafny.SeqWithoutIsStrInference([(self).f9])
            nw143_[int(8)] = d_717_v1_
            nw143_[int(9)] = d_717_v1_
            nw143_[int(10)] = d_717_v1_
            nw143_[int(11)] = d_717_v1_
            nw143_[int(12)] = _dafny.SeqWithoutIsStrInference([(self).f7])
            nw143_[int(13)] = _dafny.SeqWithoutIsStrInference([(self).f7])
            nw143_[int(14)] = d_717_v1_
            nw143_[int(15)] = d_717_v1_
            nw143_[int(16)] = d_717_v1_
            nw143_[int(17)] = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f7])
            nw143_[int(18)] = _dafny.SeqWithoutIsStrInference([(self).f9])
            nw143_[int(19)] = default__.fm29(globalState)
            nw143_[int(20)] = d_717_v1_
            nw143_[int(21)] = _dafny.SeqWithoutIsStrInference([(self).f9])
            nw143_[int(22)] = d_717_v1_
            nw143_[int(23)] = d_717_v1_
            nw143_[int(24)] = _dafny.SeqWithoutIsStrInference([(self).f6])
            nw143_[int(25)] = d_717_v1_
            d_756_v35_ = nw143_
            d_757_v36_: D1
            d_757_v36_ = D1_DC3(d_756_v35_)
            d_758_v37_: _dafny.Array
            nw144_ = _dafny.Array(None, 26)
            nw144_[int(0)] = d_757_v36_
            nw144_[int(1)] = d_757_v36_
            nw144_[int(2)] = D1_DC3(d_756_v35_)
            nw144_[int(3)] = d_757_v36_
            nw144_[int(4)] = d_757_v36_
            nw144_[int(5)] = d_757_v36_
            nw144_[int(6)] = D1_DC3(d_756_v35_)
            nw144_[int(7)] = d_757_v36_
            nw144_[int(8)] = d_757_v36_
            nw144_[int(9)] = D1_DC3(d_756_v35_)
            nw144_[int(10)] = d_757_v36_
            nw144_[int(11)] = d_757_v36_
            nw144_[int(12)] = d_757_v36_
            nw144_[int(13)] = d_757_v36_
            nw144_[int(14)] = d_757_v36_
            nw144_[int(15)] = D1_DC3(d_756_v35_)
            nw144_[int(16)] = D1_DC3(d_756_v35_)
            nw144_[int(17)] = d_757_v36_
            nw144_[int(18)] = d_757_v36_
            nw144_[int(19)] = d_757_v36_
            nw144_[int(20)] = d_757_v36_
            nw144_[int(21)] = D1_DC3(d_756_v35_)
            nw144_[int(22)] = d_757_v36_
            nw144_[int(23)] = d_757_v36_
            nw144_[int(24)] = d_757_v36_
            nw144_[int(25)] = d_757_v36_
            d_758_v37_ = nw144_
            d_759_v38_: C2
            nw145_ = C2()
            nw145_.ctor__(d_755_v34_, d_758_v37_)
            d_759_v38_ = nw145_
            d_760_v39_: _dafny.Seq
            d_760_v39_ = _dafny.SeqWithoutIsStrInference([d_759_v38_, d_759_v38_])
            d_759_v38_ = (d_760_v39_)[default__.safeIndex(default__.safeModuloInt(p0, p0), len(d_760_v39_))]
            if (self).f7:
                d_761_v40_: _dafny.Map
                d_761_v40_ = _dafny.Map({len(default__.fm30((self).f9, globalState)): (d_716_v0_).cardinality})
                d_762_v41_: _dafny.Map
                d_762_v41_ = _dafny.Map({d_761_v40_: (self).f8})
                d_763_v42_: _dafny.Map
                d_763_v42_ = _dafny.Map({(self).f7: False})
                d_764_v43_: _dafny.Seq
                d_764_v43_ = _dafny.SeqWithoutIsStrInference([p0, (self).f8])
                d_762_v41_ = (d_762_v41_).set((_dafny.Map({(d_716_v0_).cardinality: p0})).set((0) - (d_753_i2_), len((d_763_v42_).set(not((self).f7), (self).f7))), (d_753_i2_) - ((d_764_v43_)[default__.safeIndex(p0, len(d_764_v43_))]))
                d_765_v44_: bool
                d_765_v44_ = True
                d_766_v45_: _dafny.Set
                d_766_v45_ = _dafny.Set({(self).f7})
                d_767_v46_: D3
                d_767_v46_ = D3_DC11(False, p0, d_766_v45_)
                d_768_v47_: _dafny.MultiSet
                d_768_v47_ = _dafny.MultiSet([(0) - (default__.safeModuloInt(len(d_764_v43_), 646))])
                def iife62_(_pat_let22_0):
                    def iife63_(d_769_dt__update__tmp_h0_):
                        def iife64_(_pat_let23_0):
                            def iife65_(d_770_dt__update_hcf22_h0_):
                                return D3_DC11((d_769_dt__update__tmp_h0_).cf20, (d_769_dt__update__tmp_h0_).cf21, d_770_dt__update_hcf22_h0_)
                            return iife65_(_pat_let23_0)
                        return iife64_(_dafny.Set({(self).f9}))
                    return iife63_(_pat_let22_0)
                rhs100_ = (iife62_(d_767_v46_)).cf20
                rhs101_ = (d_768_v47_).cardinality
                lhs70_ = globalState
                d_765_v44_ = rhs100_
                lhs70_.f1 = rhs101_
                d_771_v48_: _dafny.Map
                d_771_v48_ = _dafny.Map({(True) or ((self).f9): self.f4})
                d_771_v48_ = d_771_v48_
                d_765_v44_ = (self).f9
                d_772_v49_: _dafny.Set
                d_772_v49_ = _dafny.Set({p0})
                d_773_v50_: _dafny.Map
                d_773_v50_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([len((self).f5)])})
                d_774_v51_: _dafny.Map
                d_774_v51_ = _dafny.Map({len(d_772_v49_): (_dafny.Map({len(d_717_v1_): ((d_773_v50_)[(self).f8] if ((self).f8) in (d_773_v50_) else d_764_v43_)})) | (_dafny.Map({264: d_764_v43_}))})
                d_774_v51_ = (d_774_v51_).set((self).f8, d_773_v50_)
            elif True:
                d_775_v52_: bool
                d_775_v52_ = False
                d_775_v52_ = not((d_775_v52_ if True else (self).f7))
                index106_ = default__.safeIndex(59, ((d_759_v38_).f14).length(0))
                ((d_759_v38_).f14)[index106_] = (p0) == (d_753_i2_)
                d_776_v53_: D1
                d_776_v53_ = D1_DC4(self.f4, d_753_i2_, self.f4, p0, (self).f6)
                d_777_v54_: _dafny.Set
                d_777_v54_ = _dafny.Set({d_776_v53_})
                index107_ = default__.safeIndex(59, ((d_759_v38_).f14).length(0))
                ((d_759_v38_).f14)[index107_] = (d_777_v54_) != (d_777_v54_)
                index108_ = default__.safeIndex(59, ((d_759_v38_).f14).length(0))
                ((d_759_v38_).f14)[index108_] = (self.f4) in (((self).f5 if not(d_775_v52_) else (self).f5))
                d_778_v55_: C0
                nw146_ = C0()
                nw146_.ctor__(((d_759_v38_).f14)[default__.safeIndex(59, ((d_759_v38_).f14).length(0))], ((d_759_v38_).f14)[default__.safeIndex(59, ((d_759_v38_).f14).length(0))], self.f4, (self).f5)
                d_778_v55_ = nw146_
                d_779_v56_: _dafny.Map
                d_779_v56_ = _dafny.Map({d_778_v55_: _dafny.CodePoint('m')})
                (globalState).f1 = default__.safeModuloInt((self).f8, len((d_779_v56_) | (d_779_v56_)))
                d_780_v57_: _dafny.Map
                d_780_v57_ = _dafny.Map({(self).f8: d_775_v52_})
                d_781_v58_: C3
                nw147_ = C3()
                nw147_.ctor__(((self).f8) in (_dafny.SeqWithoutIsStrInference([p0])), ((d_780_v57_)[d_753_i2_] if (d_753_i2_) in (d_780_v57_) else (self).f6), (0) - ((p0) - ((self).f8)), (self).f7, (self).f7, (self).f7, self.f4, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_782_i3_ in range(default__.abs(665))]))
                d_781_v58_ = nw147_
        d_783_v59_: _dafny.Array
        nw148_ = _dafny.Array(D6.default()(), 19)
        d_783_v59_ = nw148_
        d_784_v60_: D6
        d_784_v60_ = D6_DC17((self).f5, (self).f5)
        index109_ = default__.safeIndex(553, (d_783_v59_).length(0))
        (d_783_v59_)[index109_] = (d_784_v60_ if (self).f7 else D6_DC17((self).f5, (self).f5))
        d_785_v61_: _dafny.Map
        d_785_v61_ = _dafny.Map({(0) - (p0): (d_717_v1_).set(default__.safeIndex(p0, len(d_717_v1_)), (self).f9)})
        index110_ = default__.safeIndex(553, (d_783_v59_).length(0))
        rhs102_ = (self).f8
        rhs103_ = d_784_v60_
        rhs104_ = 26
        rhs105_ = _dafny.MultiSet((d_717_v1_) + (((d_785_v61_)[p0] if (p0) in (d_785_v61_) else d_717_v1_)))
        lhs71_ = globalState
        lhs72_ = d_783_v59_
        lhs73_ = default__.safeIndex(553, (d_783_v59_).length(0))
        lhs74_ = globalState
        lhs71_.f1 = rhs102_
        lhs72_[lhs73_] = rhs103_
        lhs74_.f1 = rhs104_
        d_716_v0_ = rhs105_
        hi4_ = (0) - ((self).f8)
        for d_786_i4_ in range((p0 if (self).f7 else (self).f8), hi4_):
            if (d_717_v1_)[default__.safeIndex((0) - (d_786_i4_), len(d_717_v1_))]:
                d_787_v62_: _dafny.Array
                nw149_ = _dafny.Array(int(0), 28)
                d_787_v62_ = nw149_
                index111_ = default__.safeIndex(883, (d_787_v62_).length(0))
                (d_787_v62_)[index111_] = (0) - (d_786_i4_)
                index112_ = default__.safeIndex(883, (d_787_v62_).length(0))
                (d_787_v62_)[index112_] = p0
                d_788_v63_: _dafny.Map
                d_788_v63_ = _dafny.Map({(self).f6: self.f4})
                (globalState).f1 = (d_786_i4_ if (self).f6 else ((d_787_v62_)[default__.safeIndex(883, (d_787_v62_).length(0))]) * (len(d_788_v63_)))
                d_789_v64_: _dafny.Array
                nw150_ = _dafny.Array(False, 28)
                d_789_v64_ = nw150_
                d_790_v65_: _dafny.Map
                d_790_v65_ = _dafny.Map({(self).f6: (self).f8})
                d_791_v66_: _dafny.MultiSet
                d_791_v66_ = _dafny.MultiSet([(self).f8])
                d_792_v67_: D5
                d_792_v67_ = D5_DC15(d_789_v64_, (self).f9, d_790_v65_, (d_791_v66_).cardinality)
                d_793_v68_: _dafny.Map
                d_793_v68_ = _dafny.Map({(self).f7: (d_792_v67_).cf27})
                d_794_v69_: _dafny.Seq
                d_794_v69_ = _dafny.SeqWithoutIsStrInference([d_793_v68_, d_793_v68_])
                d_795_v70_: _dafny.Map
                d_795_v70_ = _dafny.Map({(self).f7: len(d_794_v69_)})
                d_790_v65_ = (d_790_v65_).set((self).f6, (self).fm1((self).f6, d_793_v68_, (self).f9, globalState))
                (globalState).f1 = d_786_i4_
                d_796_v71_: bool
                d_796_v71_ = True
                index113_ = default__.safeIndex(883, (d_787_v62_).length(0))
                rhs106_ = (p0) > (len(_dafny.Map({self.f4: d_789_v64_})))
                rhs107_ = default__.safeDivisionInt(d_786_i4_, (self).f8)
                lhs75_ = d_787_v62_
                lhs76_ = default__.safeIndex(883, (d_787_v62_).length(0))
                d_796_v71_ = rhs106_
                lhs75_[lhs76_] = rhs107_
            elif True:
                (globalState).f1 = 382
                d_797_v72_: bool
                d_797_v72_ = False
                d_797_v72_ = (self).f9
                d_798_v73_: _dafny.Set
                d_798_v73_ = _dafny.Set({self.f4, self.f4})
                d_799_v74_: _dafny.Array
                nw151_ = _dafny.Array(None, 11)
                nw151_[int(0)] = d_797_v72_
                nw151_[int(1)] = (self).f6
                nw151_[int(2)] = (self).f6
                nw151_[int(3)] = (self).f7
                nw151_[int(4)] = (self).f9
                nw151_[int(5)] = ((self).f8) >= (p0)
                nw151_[int(6)] = (_dafny.Set({self.f4})).issubset(d_798_v73_)
                nw151_[int(7)] = (d_797_v72_) and ((self).f7)
                nw151_[int(8)] = d_797_v72_
                nw151_[int(9)] = (self).f7
                nw151_[int(10)] = True
                d_799_v74_ = nw151_
                index114_ = default__.safeIndex(174, (d_799_v74_).length(0))
                (d_799_v74_)[index114_] = d_797_v72_
                index115_ = default__.safeIndex(174, (d_799_v74_).length(0))
                (d_799_v74_)[index115_] = (self).f7
                d_800_v75_: D2
                d_800_v75_ = D2_DC8(p0, p0)
                d_801_v76_: D2
                d_801_v76_ = D2_DC9(d_800_v75_)
                d_801_v76_ = (d_801_v76_ if True else d_801_v76_)
                def lambda26_(d_802_v74_):
                    def lambda27_(d_803_i5_):
                        return (d_802_v74_)[default__.safeIndex(174, (d_802_v74_).length(0))]

                    return lambda27_

                init15_ = lambda26_(d_799_v74_)
                nw152_ = _dafny.Array(None, 20)
                for i0_15_ in range(nw152_.length(0)):
                    nw152_[i0_15_] = init15_(i0_15_)
                d_799_v74_ = nw152_
            d_804_v77_: _dafny.Array
            nw153_ = _dafny.Array(False, 27)
            d_804_v77_ = nw153_
            d_805_v78_: _dafny.Set
            d_805_v78_ = _dafny.Set({d_804_v77_})
            d_806_v79_: D7
            d_806_v79_ = D7_DC18(d_805_v78_)
            pat_let_tv15_ = d_804_v77_
            def iife66_(_pat_let24_0):
                def iife67_(d_807_dt__update__tmp_h1_):
                    def iife68_(_pat_let25_0):
                        def iife69_(d_808_dt__update_hcf33_h0_):
                            return D7_DC18(d_808_dt__update_hcf33_h0_)
                        return iife69_(_pat_let25_0)
                    return iife68_(_dafny.Set({pat_let_tv15_}))
                return iife67_(_pat_let24_0)
            d_806_v79_ = iife66_(D7_DC18(d_805_v78_))
            d_809_v80_: bool
            d_809_v80_ = False
            d_809_v80_ = False
            d_810_v81_: _dafny.Map
            d_810_v81_ = _dafny.Map({(self).f7: (self).f8})
            d_811_v82_: _dafny.Map
            d_811_v82_ = _dafny.Map({(self).f6: (self).f6})
            d_812_v83_: D3
            d_812_v83_ = D3_DC11((self).f6, p0, _dafny.Set({(self).f6}))
            d_813_v84_: _dafny.Array
            def lambda28_(d_814_v82_):
                def lambda29_(d_815_i6_):
                    return default__.safeDivisionInt(d_815_i6_, len(d_814_v82_))

                return lambda29_

            init16_ = lambda28_(d_811_v82_)
            nw154_ = _dafny.Array(None, 28)
            for i0_16_ in range(nw154_.length(0)):
                nw154_[i0_16_] = init16_(i0_16_)
            d_813_v84_ = nw154_
            d_816_v85_: _dafny.MultiSet
            d_816_v85_ = _dafny.MultiSet([d_813_v84_, d_813_v84_, d_813_v84_, d_813_v84_, d_813_v84_])
            d_817_v86_: _dafny.Array
            nw155_ = _dafny.Array(None, 24)
            nw155_[int(0)] = d_786_i4_
            nw155_[int(1)] = (((d_810_v81_)[(self).f6] if ((self).f6) in (d_810_v81_) else len((self).f5))) - (456)
            nw155_[int(2)] = 280
            nw155_[int(3)] = ((self).fm1((self).f7, _dafny.Map({(self).f7: (self).f9}), (self).f7, globalState)) * (p0)
            nw155_[int(4)] = (self).f8
            nw155_[int(5)] = (self).fm1(d_809_v80_, (d_811_v82_).set((self).f6, (self).f7), (self).f7, globalState)
            nw155_[int(6)] = (self).f8
            nw155_[int(7)] = d_786_i4_
            nw155_[int(8)] = (d_812_v83_).cf21
            nw155_[int(9)] = (0) - (d_786_i4_)
            nw155_[int(10)] = (self).f8
            nw155_[int(11)] = (len((self).f5) if (self).f6 else d_786_i4_)
            nw155_[int(12)] = (0) - (p0)
            nw155_[int(13)] = ((d_816_v85_).cardinality) * ((self).f8)
            nw155_[int(14)] = d_786_i4_
            nw155_[int(15)] = (d_786_i4_) + (p0)
            nw155_[int(16)] = (0) - (p0)
            nw155_[int(17)] = d_786_i4_
            nw155_[int(18)] = len((self).f5)
            nw155_[int(19)] = (self).f8
            nw155_[int(20)] = (0) - (default__.safeModuloInt(d_786_i4_, len((self).f5)))
            nw155_[int(21)] = d_786_i4_
            nw155_[int(22)] = (self).f8
            nw155_[int(23)] = -869
            d_817_v86_ = nw155_
            rhs108_ = (777) - ((self).f8)
            rhs109_ = d_813_v84_
            rhs110_ = default__.safeDivisionInt(578, d_786_i4_)
            rhs111_ = p0
            lhs77_ = globalState
            lhs78_ = globalState
            lhs79_ = globalState
            lhs77_.f1 = rhs108_
            d_817_v86_ = rhs109_
            lhs78_.f1 = rhs110_
            lhs79_.f1 = rhs111_
        (globalState).f1 = (self).f8
        d_818_i7_: int
        d_818_i7_ = 0
        with _dafny.label("4"):
            while (p0) < (len((self).f5)):
                with _dafny.c_label("4"):
                    if (d_818_i7_) >= (100):
                        raise _dafny.Break("4")
                    d_818_i7_ = (d_818_i7_) + (1)
                    d_819_v87_: bool
                    d_819_v87_ = False
                    d_819_v87_ = ((self).f5) <= ((self).f5)
                    d_820_v88_: _dafny.Map
                    d_820_v88_ = _dafny.Map({d_819_v87_: (self).f9})
                    d_821_v89_: _dafny.Map
                    d_821_v89_ = _dafny.Map({d_819_v87_: _dafny.SeqWithoutIsStrInference([p0 for d_822_i8_ in range(default__.abs(58))])})
                    (globalState).f1 = (self).fm1(True, (d_820_v88_) | (_dafny.Map({d_819_v87_: True})), not((d_821_v89_) != (d_821_v89_)), globalState)
                    d_823_v90_: _dafny.Array
                    nw156_ = _dafny.Array(int(0), 11)
                    d_823_v90_ = nw156_
                    index116_ = default__.safeIndex(296, (d_823_v90_).length(0))
                    (d_823_v90_)[index116_] = (self).f8
                    index117_ = default__.safeIndex(296, (d_823_v90_).length(0))
                    (d_823_v90_)[index117_] = p0
                    d_824_v91_: _dafny.Map
                    d_824_v91_ = _dafny.Map({97: False})
                    d_825_v92_: _dafny.Set
                    d_825_v92_ = _dafny.Set({True})
                    d_826_v93_: D3
                    d_826_v93_ = D3_DC11(((d_824_v91_)[(0) - (p0)] if ((0) - (p0)) in (d_824_v91_) else (self).f7), 26, d_825_v92_)
                    d_827_v94_: T1
                    nw157_ = C0()
                    nw157_.ctor__(d_819_v87_, (self).f7, self.f4, default__.fm28((d_823_v90_)[default__.safeIndex(296, (d_823_v90_).length(0))], (self).f6, d_717_v1_, d_826_v93_, globalState))
                    d_827_v94_ = nw157_
                    d_828_v95_: _dafny.Array
                    nw158_ = _dafny.Array(_dafny.Array(None, 0), 7)
                    d_828_v95_ = nw158_
                    d_829_v96_: _dafny.Array
                    nw159_ = _dafny.Array(False, 12)
                    d_829_v96_ = nw159_
                    d_830_v97_: _dafny.Array
                    nw160_ = _dafny.Array(None, 3)
                    nw160_[int(0)] = d_829_v96_
                    nw160_[int(1)] = d_829_v96_
                    nw160_[int(2)] = d_829_v96_
                    d_830_v97_ = nw160_
                    index118_ = default__.safeIndex(422, (d_828_v95_).length(0))
                    (d_828_v95_)[index118_] = d_830_v97_
                    d_831_v98_: _dafny.Array
                    def lambda30_(d_832_v94_, d_833_p0_, d_834_v87_):
                        def lambda31_(d_835_i9_):
                            return D0_DC1((d_832_v94_).f6, d_833_p0_, d_834_v87_, (self).f8)

                        return lambda31_

                    init17_ = lambda30_(d_827_v94_, p0, d_819_v87_)
                    nw161_ = _dafny.Array(None, 4)
                    for i0_17_ in range(nw161_.length(0)):
                        nw161_[i0_17_] = init17_(i0_17_)
                    d_831_v98_ = nw161_
                    d_836_v99_: D0
                    d_836_v99_ = D0_DC1((self).f9, (0) - ((self).f8), (d_827_v94_).f7, len((d_827_v94_).f5))
                    index119_ = default__.safeIndex(25, (d_831_v98_).length(0))
                    (d_831_v98_)[index119_] = d_836_v99_
                    index120_ = default__.safeIndex(422, (d_828_v95_).length(0))
                    index121_ = default__.safeIndex(25, (d_831_v98_).length(0))
                    rhs112_ = (d_827_v94_ if (self).f6 else d_827_v94_)
                    rhs113_ = d_830_v97_
                    rhs114_ = d_836_v99_
                    lhs80_ = d_828_v95_
                    lhs81_ = default__.safeIndex(422, (d_828_v95_).length(0))
                    lhs82_ = d_831_v98_
                    lhs83_ = default__.safeIndex(25, (d_831_v98_).length(0))
                    d_827_v94_ = rhs112_
                    lhs80_[lhs81_] = rhs113_
                    lhs82_[lhs83_] = rhs114_
                    pass
            pass
        d_837_v100_: _dafny.Map
        d_837_v100_ = _dafny.Map({p0: ((self).f8) + (p0)})
        r0 = d_837_v100_
        return r0


class C5(T1, T2, T3, T0):
    def  __init__(self):
        self._f4: str = _dafny.CodePoint('D')
        self._f8: int = int(0)
        self._f9: bool = False
        self._f6: bool = False
        self._f7: bool = False
        self._f5: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: _dafny.Array = _dafny.Array(None, 0)
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f4(self):
        return self._f4
    @f4.setter
    def f4(self, value):
        self._f4 = value
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f10, f11, f6, f7, f4, f5, f8, f9):
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f8 = f8
        (self)._f9 = f9

    def fm1(self, p0, p1, p2, globalState):
        return (0) - ((self).f8)

    def fm2(self, globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f6])), (self).f8])) + ((_dafny.SeqWithoutIsStrInference([(self).f8])) + (_dafny.SeqWithoutIsStrInference([(self).f8 for d_838_i0_ in range(default__.abs(342))])))

    def m1(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_839_v0_: D0
        d_839_v0_ = D0_DC0()
        def lambda32_(source10_):
            if source10_.is_DC0:
                return D0_DC1((self).f6, (self).f8, (self).f9, (self).f8)
            elif source10_.is_DC1:
                d_840___mcc_h7_ = source10_.cf0
                d_841___mcc_h8_ = source10_.cf1
                d_842___mcc_h9_ = source10_.cf2
                d_843___mcc_h10_ = source10_.cf3
                d_844_cf3_ = d_843___mcc_h10_
                d_845_cf2_ = d_842___mcc_h9_
                d_846_cf1_ = d_841___mcc_h8_
                d_847_cf0_ = d_840___mcc_h7_
                return D0_DC1((self).f9, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f8 for d_848_i0_ in range(default__.abs(528))])), len((self).f5)])), d_847_cf0_, len(_dafny.Map({False: (self).f9})))
            elif True:
                d_849___mcc_h11_ = source10_.cf4
                d_850___mcc_h12_ = source10_.cf5
                d_851___mcc_h13_ = source10_.cf6
                d_852_cf6_ = d_851___mcc_h13_
                d_853_cf5_ = d_850___mcc_h12_
                d_854_cf4_ = d_849___mcc_h11_
                return D0_DC1(True, 556, (self).f7, len(_dafny.Map({d_854_cf4_: d_853_cf5_})))

        source9_ = lambda32_(d_839_v0_)
        if source9_.is_DC0:
            d_855_v1_: bool
            d_855_v1_ = False
            d_855_v1_ = (self).f9
            d_855_v1_ = d_855_v1_
            d_856_v2_: _dafny.Seq
            d_856_v2_ = _dafny.SeqWithoutIsStrInference([(self).f6])
            d_857_v3_: _dafny.Seq
            d_857_v3_ = _dafny.SeqWithoutIsStrInference([d_856_v2_, d_856_v2_])
            d_858_v4_: _dafny.Map
            d_858_v4_ = _dafny.Map({539: (self).f8})
            d_859_v5_: _dafny.Seq
            d_859_v5_ = _dafny.SeqWithoutIsStrInference([d_858_v4_, d_858_v4_])
            d_860_v6_: _dafny.Set
            d_860_v6_ = _dafny.Set({(self).f8})
            d_861_v7_: _dafny.Map
            d_861_v7_ = _dafny.Map({727: (self).f7})
            d_862_v8_: _dafny.Array
            nw162_ = _dafny.Array(None, 22)
            nw162_[int(0)] = ((d_857_v3_)[default__.safeIndex(len(d_859_v5_), len(d_857_v3_))]) + (d_856_v2_)
            nw162_[int(1)] = d_856_v2_
            nw162_[int(2)] = d_856_v2_
            nw162_[int(3)] = d_856_v2_
            nw162_[int(4)] = d_856_v2_
            nw162_[int(5)] = d_856_v2_
            nw162_[int(6)] = ((d_857_v3_)[default__.safeIndex(len(d_860_v6_), len(d_857_v3_))]) + (d_856_v2_)
            nw162_[int(7)] = d_856_v2_
            nw162_[int(8)] = (_dafny.SeqWithoutIsStrInference([(self).f7])) + (_dafny.SeqWithoutIsStrInference([True]))
            nw162_[int(9)] = d_856_v2_
            nw162_[int(10)] = d_856_v2_
            nw162_[int(11)] = ((d_856_v2_).set(default__.safeIndex(102, len(d_856_v2_)), d_855_v1_)) + (_dafny.SeqWithoutIsStrInference([p0, p0]))
            nw162_[int(12)] = (d_856_v2_) + (d_856_v2_)
            nw162_[int(13)] = d_856_v2_
            nw162_[int(14)] = (_dafny.SeqWithoutIsStrInference([(self).f9, not(False)])) + (d_856_v2_)
            nw162_[int(15)] = d_856_v2_
            nw162_[int(16)] = _dafny.SeqWithoutIsStrInference([p1, ((d_861_v7_)[(self).f8] if ((self).f8) in (d_861_v7_) else p1), default__.fm3(globalState)])
            nw162_[int(17)] = d_856_v2_
            nw162_[int(18)] = _dafny.SeqWithoutIsStrInference([(self).f6])
            nw162_[int(19)] = d_856_v2_
            nw162_[int(20)] = (d_856_v2_) + (d_856_v2_)
            nw162_[int(21)] = (_dafny.SeqWithoutIsStrInference([(self).f6, p0])) + (d_856_v2_)
            d_862_v8_ = nw162_
            d_863_v9_: _dafny.Seq
            d_863_v9_ = _dafny.SeqWithoutIsStrInference([(self).f11])
            d_864_v10_: _dafny.Array
            nw163_ = _dafny.Array(None, 26)
            nw163_[int(0)] = (self).f11
            nw163_[int(1)] = (self).f11
            nw163_[int(2)] = (self).f11
            nw163_[int(3)] = (self).f11
            nw163_[int(4)] = (self).f11
            nw163_[int(5)] = (self).f11
            nw163_[int(6)] = (self).f11
            nw163_[int(7)] = (self).f11
            nw163_[int(8)] = (self).f11
            nw163_[int(9)] = (d_863_v9_)[default__.safeIndex((self).f8, len(d_863_v9_))]
            nw163_[int(10)] = (self).f11
            nw163_[int(11)] = (self).f11
            nw163_[int(12)] = (self).f11
            nw163_[int(13)] = (self).f11
            nw163_[int(14)] = (self).f11
            nw163_[int(15)] = (self).f11
            nw163_[int(16)] = (self).f11
            nw163_[int(17)] = (d_863_v9_)[default__.safeIndex((self).f8, len(d_863_v9_))]
            nw163_[int(18)] = (self).f11
            nw163_[int(19)] = (d_863_v9_)[default__.safeIndex((self).f8, len(d_863_v9_))]
            nw163_[int(20)] = (self).f11
            nw163_[int(21)] = (self).f11
            nw163_[int(22)] = (self).f11
            nw163_[int(23)] = (self).f11
            nw163_[int(24)] = (self).f11
            nw163_[int(25)] = (self).f11
            d_864_v10_ = nw163_
            d_865_v11_: _dafny.Map
            d_865_v11_ = _dafny.Map({(self).f8: (self).f11})
            index122_ = default__.safeIndex(395, (d_864_v10_).length(0))
            (d_864_v10_)[index122_] = ((d_865_v11_)[(default__.fm4(p0, (self).f8, (self).f7, len(d_860_v6_), globalState)).cardinality] if ((default__.fm4(p0, (self).f8, (self).f7, len(d_860_v6_), globalState)).cardinality) in (d_865_v11_) else (self).f11)
            index123_ = default__.safeIndex(395, (d_864_v10_).length(0))
            rhs115_ = (D1_DC3(d_862_v8_)).cf7
            rhs116_ = (self).f11
            lhs84_ = d_864_v10_
            lhs85_ = default__.safeIndex(395, (d_864_v10_).length(0))
            d_862_v8_ = rhs115_
            lhs84_[lhs85_] = rhs116_
            d_855_v1_ = not(not(d_855_v1_))
        elif source9_.is_DC1:
            d_866___mcc_h0_ = source9_.cf0
            d_867___mcc_h1_ = source9_.cf1
            d_868___mcc_h2_ = source9_.cf2
            d_869___mcc_h3_ = source9_.cf3
            d_870_cf3_ = d_869___mcc_h3_
            d_871_cf2_ = d_868___mcc_h2_
            d_872_cf1_ = d_867___mcc_h1_
            d_873_cf0_ = d_866___mcc_h0_
            d_871_cf2_ = not(d_871_cf2_)
            d_874_v12_: _dafny.Seq
            d_874_v12_ = _dafny.SeqWithoutIsStrInference([p1])
            index124_ = default__.safeIndex(385, ((self).f10).length(0))
            ((self).f10)[index124_] = (d_874_v12_)[default__.safeIndex((self).f8, len(d_874_v12_))]
            index125_ = default__.safeIndex(385, ((self).f10).length(0))
            ((self).f10)[index125_] = p0
            d_875_v13_: _dafny.Set
            d_875_v13_ = _dafny.Set({not(d_873_cf0_)})
            d_876_v14_: _dafny.Array
            nw164_ = _dafny.Array(None, 19)
            nw164_[int(0)] = p1
            nw164_[int(1)] = default__.fm7(self.f4, len(d_875_v13_), p0, globalState)
            nw164_[int(2)] = p3
            nw164_[int(3)] = ((self).f10)[default__.safeIndex(385, ((self).f10).length(0))]
            nw164_[int(4)] = not(default__.fm3(globalState))
            nw164_[int(5)] = (self).f9
            nw164_[int(6)] = ((self).f10)[default__.safeIndex(385, ((self).f10).length(0))]
            nw164_[int(7)] = p1
            nw164_[int(8)] = d_873_cf0_
            nw164_[int(9)] = (self).f9
            nw164_[int(10)] = (self).f9
            nw164_[int(11)] = not(d_871_cf2_)
            nw164_[int(12)] = (d_874_v12_)[default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference([self.f4 for d_877_i1_ in range(default__.abs(561))]))), len(d_874_v12_))]
            nw164_[int(13)] = d_871_cf2_
            nw164_[int(14)] = d_873_cf0_
            nw164_[int(15)] = ((self).f10)[default__.safeIndex(385, ((self).f10).length(0))]
            nw164_[int(16)] = ((self).f10)[default__.safeIndex(385, ((self).f10).length(0))]
            nw164_[int(17)] = default__.fm3(globalState)
            nw164_[int(18)] = (self).f7
            d_876_v14_ = nw164_
            d_878_v15_: _dafny.Array
            nw165_ = _dafny.Array(_dafny.Seq({}), 9)
            d_878_v15_ = nw165_
            d_879_v16_: D1
            d_879_v16_ = D1_DC3(d_878_v15_)
            d_880_v17_: _dafny.Array
            nw166_ = _dafny.Array(None, 5)
            nw166_[int(0)] = d_879_v16_
            nw166_[int(1)] = d_879_v16_
            nw166_[int(2)] = D1_DC3(d_878_v15_)
            nw166_[int(3)] = d_879_v16_
            nw166_[int(4)] = D1_DC3(d_878_v15_)
            d_880_v17_ = nw166_
            d_881_v18_: C2
            nw167_ = C2()
            nw167_.ctor__(d_876_v14_, d_880_v17_)
            d_881_v18_ = nw167_
            if ((len((self).f5)) - (-466)) > ((0) - ((self).f8)):
                d_882_v19_: C0
                nw168_ = C0()
                nw168_.ctor__(not(default__.fm7(self.f4, d_870_cf3_, not(True), globalState)), p3, self.f4, (self).f5)
                d_882_v19_ = nw168_
                d_883_v20_: _dafny.Array
                nw169_ = _dafny.Array(None, 18)
                nw169_[int(0)] = (d_882_v19_ if (self).f6 else d_882_v19_)
                nw169_[int(1)] = d_882_v19_
                nw169_[int(2)] = d_882_v19_
                nw169_[int(3)] = d_882_v19_
                nw169_[int(4)] = d_882_v19_
                nw169_[int(5)] = d_882_v19_
                nw169_[int(6)] = (d_882_v19_ if not(d_873_cf0_) else d_882_v19_)
                nw169_[int(7)] = d_882_v19_
                nw169_[int(8)] = d_882_v19_
                nw169_[int(9)] = d_882_v19_
                nw169_[int(10)] = d_882_v19_
                nw169_[int(11)] = d_882_v19_
                nw169_[int(12)] = d_882_v19_
                nw169_[int(13)] = d_882_v19_
                nw169_[int(14)] = d_882_v19_
                nw169_[int(15)] = d_882_v19_
                nw169_[int(16)] = d_882_v19_
                nw169_[int(17)] = d_882_v19_
                d_883_v20_ = nw169_
                index126_ = default__.safeIndex(792, (d_883_v20_).length(0))
                (d_883_v20_)[index126_] = d_882_v19_
                d_884_v21_: _dafny.Seq
                d_884_v21_ = _dafny.SeqWithoutIsStrInference([d_882_v19_])
                d_885_v22_: _dafny.Seq
                d_885_v22_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False]), d_874_v12_])
                index127_ = default__.safeIndex(792, (d_883_v20_).length(0))
                index128_ = default__.safeIndex(385, ((self).f10).length(0))
                rhs117_ = ((self).f8) - (d_870_cf3_)
                rhs118_ = (d_884_v21_)[default__.safeIndex(((self).f8) * ((d_882_v19_).fm9(d_872_cf1_, (d_885_v22_)[default__.safeIndex((self).f8, len(d_885_v22_))], globalState)), len(d_884_v21_))]
                rhs119_ = ((self).f7) not in ((d_874_v12_) + (d_874_v12_))
                lhs86_ = d_883_v20_
                lhs87_ = default__.safeIndex(792, (d_883_v20_).length(0))
                lhs88_ = (self).f10
                lhs89_ = default__.safeIndex(385, ((self).f10).length(0))
                d_870_cf3_ = rhs117_
                lhs86_[lhs87_] = rhs118_
                lhs88_[lhs89_] = rhs119_
                d_886_v23_: _dafny.Seq
                d_886_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "muet"))
                d_887_v24_: _dafny.Map
                d_887_v24_ = _dafny.Map({not(d_873_cf0_): False})
                d_886_v23_ = (((self).f5) + (d_886_v23_)) + (((self).f5) + (default__.fm14((self).fm1(((self).f10)[default__.safeIndex(385, ((self).f10).length(0))], d_887_v24_, not(d_873_cf0_), globalState), (self).f8, False, p1, globalState)))
                index129_ = default__.safeIndex(792, (d_883_v20_).length(0))
                (d_883_v20_)[index129_] = (d_883_v20_)[default__.safeIndex(792, (d_883_v20_).length(0))]
                index130_ = default__.safeIndex(385, ((self).f10).length(0))
                ((self).f10)[index130_] = p1
                (globalState).f1 = d_872_cf1_
            elif True:
                d_888_v25_: _dafny.Array
                nw170_ = _dafny.Array(_dafny.CodePoint('D'), 7)
                d_888_v25_ = nw170_
                d_889_v26_: _dafny.Map
                d_889_v26_ = _dafny.Map({(self).f7: (self).f9})
                index131_ = default__.safeIndex(385, ((self).f10).length(0))
                rhs120_ = d_873_cf0_
                rhs121_ = d_888_v25_
                rhs122_ = ((_dafny.Map({((self).f10)[default__.safeIndex(385, ((self).f10).length(0))]: p0})) | (d_889_v26_)) | ((d_889_v26_) | ((d_889_v26_).set((self).f7, (self).f7)))
                rhs123_ = default__.safeDivisionInt(d_872_cf1_, len(_dafny.SeqWithoutIsStrInference([d_870_cf3_])))
                rhs124_ = (d_870_cf3_) - (d_872_cf1_)
                lhs90_ = (self).f10
                lhs91_ = default__.safeIndex(385, ((self).f10).length(0))
                lhs92_ = globalState
                lhs93_ = globalState
                lhs90_[lhs91_] = rhs120_
                d_888_v25_ = rhs121_
                d_889_v26_ = rhs122_
                lhs92_.f1 = rhs123_
                lhs93_.f1 = rhs124_
                index132_ = default__.safeIndex(385, ((self).f10).length(0))
                ((self).f10)[index132_] = d_871_cf2_
                d_890_v27_: _dafny.Seq
                d_890_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsc"))
                d_890_v27_ = (d_890_v27_).set(default__.safeIndex(d_872_cf1_, len(d_890_v27_)), self.f4)
                d_870_cf3_ = d_870_cf3_
                d_891_v28_: D1
                d_891_v28_ = D1_DC5(d_872_cf1_)
                d_892_v29_: _dafny.Seq
                d_892_v29_ = _dafny.SeqWithoutIsStrInference([d_891_v28_, d_891_v28_])
                d_893_v30_: _dafny.Seq
                d_893_v30_ = (d_892_v29_) + (d_892_v29_)
                d_893_v30_ = ((d_892_v29_).set(default__.safeIndex(635, len(d_892_v29_)), d_891_v28_)).set(default__.safeIndex((self).f8, len((d_892_v29_).set(default__.safeIndex(635, len(d_892_v29_)), d_891_v28_))), default__.fm25(globalState))
        elif True:
            d_894___mcc_h4_ = source9_.cf4
            d_895___mcc_h5_ = source9_.cf5
            d_896___mcc_h6_ = source9_.cf6
            d_897_cf6_ = d_896___mcc_h6_
            d_898_cf5_ = d_895___mcc_h5_
            d_899_cf4_ = d_894___mcc_h4_
            d_900_v31_: _dafny.Seq
            d_900_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eyth"))
            d_901_v32_: _dafny.MultiSet
            d_901_v32_ = _dafny.MultiSet([555, (self).f8])
            d_900_v31_ = (d_900_v31_ if p0 else default__.fm14((self).f8, ((d_901_v32_)[(self).f8] if ((self).f8) in (d_901_v32_) else (self).f8), d_898_cf5_, True, globalState))
            d_902_v33_: _dafny.Array
            nw171_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
            d_902_v33_ = nw171_
            index133_ = default__.safeIndex(891, (d_902_v33_).length(0))
            (d_902_v33_)[index133_] = default__.fm14(d_899_cf4_, (self).f8, (self).f7, p0, globalState)
            index134_ = default__.safeIndex(891, (d_902_v33_).length(0))
            (d_902_v33_)[index134_] = ((self).f5).set(default__.safeIndex(default__.safeDivisionInt(len(d_900_v31_), (p2).cardinality), len((self).f5)), self.f4)
            def iife70_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in (d_901_v32_).Elements:
                    d_903_v34_: int = compr_18_
                    if (d_903_v34_) in (d_901_v32_):
                        coll18_[default__.safeDivisionInt(d_903_v34_, len(_dafny.SeqWithoutIsStrInference([(self).f9, True])))] = self.f4
                return _dafny.Map(coll18_)
            if ((len(iife70_()
            )) * ((self).f8)) != ((self).f8):
                d_898_cf5_ = p0
                r1 = (self).f8
                d_904_v35_: D1
                d_904_v35_ = D1_DC5(d_899_cf4_)
                d_905_v36_: _dafny.Seq
                d_905_v36_ = _dafny.SeqWithoutIsStrInference([d_904_v35_])
                d_905_v36_ = _dafny.SeqWithoutIsStrInference([D1_DC5(d_899_cf4_), d_904_v35_, d_904_v35_])
                d_898_cf5_ = ((not(p0) if (self).f6 else (self).f7)) or (not(d_898_cf5_))
                d_906_v37_: _dafny.Array
                nw172_ = _dafny.Array(_dafny.MultiSet({}), 25)
                d_906_v37_ = nw172_
                index135_ = default__.safeIndex(695, (d_906_v37_).length(0))
                (d_906_v37_)[index135_] = d_901_v32_
                index136_ = default__.safeIndex(695, (d_906_v37_).length(0))
                (d_906_v37_)[index136_] = d_901_v32_
            elif True:
                d_907_v38_: _dafny.Map
                d_907_v38_ = _dafny.Map({len(_dafny.Set({d_898_cf5_})): _dafny.CodePoint('k')})
                (self).f4 = ((d_907_v38_)[default__.safeModuloInt(d_899_cf4_, len((self).f5))] if (default__.safeModuloInt(d_899_cf4_, len((self).f5))) in (d_907_v38_) else self.f4)
                (globalState).f1 = (0) - ((0) - ((0) - ((0) - ((d_899_cf4_) - ((self).f8)))))
                d_908_v39_: C1
                nw173_ = C1()
                nw173_.ctor__(len(_dafny.SeqWithoutIsStrInference([p1])), not(p1), default__.fm3(globalState), p1, default__.fm13(globalState), (d_902_v33_)[default__.safeIndex(891, (d_902_v33_).length(0))])
                d_908_v39_ = nw173_
                index137_ = default__.safeIndex(54, ((self).f10).length(0))
                ((self).f10)[index137_] = p1
                index138_ = default__.safeIndex(54, ((self).f10).length(0))
                ((self).f10)[index138_] = p3
                d_909_v40_: _dafny.Array
                nw174_ = _dafny.Array(False, 8)
                d_909_v40_ = nw174_
                d_910_v41_: _dafny.Set
                d_910_v41_ = _dafny.Set({d_909_v40_})
                d_898_cf5_ = not((d_910_v41_).isdisjoint(d_910_v41_))
            d_911_v42_: _dafny.Map
            d_911_v42_ = _dafny.Map({(d_899_cf4_ if (self).f7 else d_899_cf4_): (self).f8})
            d_912_v43_: C3
            nw175_ = C3()
            nw175_.ctor__(p3, p3, 33, p0, p3, p0, self.f4, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkhm")))
            d_912_v43_ = nw175_
            d_913_v44_: _dafny.MultiSet
            d_913_v44_ = _dafny.MultiSet([d_912_v43_])
            d_899_cf4_ = (0) - (((d_911_v42_)[(837) + ((d_913_v44_).cardinality)] if ((837) + ((d_913_v44_).cardinality)) in (d_911_v42_) else d_899_cf4_))
        if False:
            d_914_v45_: bool
            d_914_v45_ = True
            d_914_v45_ = not((self).f7)
            d_915_v46_: _dafny.Map
            d_915_v46_ = _dafny.Map({d_914_v45_: True})
            d_916_v47_: _dafny.Map
            d_916_v47_ = _dafny.Map({p0: (self).fm1(p3, d_915_v46_, (self).f6, globalState)})
            d_914_v45_ = ((self).f8) == (((d_916_v47_)[d_914_v45_] if (d_914_v45_) in (d_916_v47_) else len((self).f5)))
            (globalState).f1 = (((self).f8 if p0 else (self).f8)) * ((self).f8)
            d_914_v45_ = not((self).f6)
            d_917_v48_: C3
            nw176_ = C3()
            nw176_.ctor__((self).f7, p0, (self).f8, p0, d_914_v45_, p0, self.f4, ((self).f5) + ((self).f5))
            d_917_v48_ = nw176_
        elif True:
            if default__.fm3(globalState):
                d_918_v49_: _dafny.Set
                d_918_v49_ = _dafny.Set({(self).f8, ((self).f8 if (self).f7 else (self).f8)})
                def iife71_():
                    coll19_ = _dafny.Set()
                    compr_19_: int
                    for compr_19_ in _dafny.IntegerRange(190, 516):
                        d_919_v50_: int = compr_19_
                        if ((190) <= (d_919_v50_)) and ((d_919_v50_) < (516)):
                            coll19_ = coll19_.union(_dafny.Set([(d_919_v50_) + ((p2).cardinality)]))
                    return _dafny.Set(coll19_)
                d_918_v49_ = iife71_()
                
                d_920_v51_: bool
                d_920_v51_ = True
                rhs125_ = self.f4
                rhs126_ = (self).f7
                lhs94_ = self
                lhs94_.f4 = rhs125_
                d_920_v51_ = rhs126_
                r1 = (self).f8
                index139_ = default__.safeIndex(107, ((self).f11).length(0))
                ((self).f11)[index139_] = (self).f8
                index140_ = default__.safeIndex(107, ((self).f11).length(0))
                ((self).f11)[index140_] = (self).f8
                d_921_v52_: D3
                d_921_v52_ = D3_DC10((self).f5)
                d_922_v53_: _dafny.MultiSet
                d_922_v53_ = _dafny.MultiSet([(self).f5, ((d_921_v52_).cf19) + ((self).f5), (self).f5])
                d_922_v53_ = d_922_v53_
            elif True:
                r1 = (default__.safeModuloInt(len((self).f5), 134)) + (default__.safeModuloInt((self).f8, (self).f8))
                index141_ = default__.safeIndex(694, ((self).f11).length(0))
                ((self).f11)[index141_] = (self).f8
                index142_ = default__.safeIndex(694, ((self).f11).length(0))
                ((self).f11)[index142_] = ((self).f8) + (((self).f8) + ((self).f8))
                d_923_v54_: bool
                d_923_v54_ = False
                d_924_v55_: _dafny.Set
                d_924_v55_ = _dafny.Set({((self).f11)[default__.safeIndex(694, ((self).f11).length(0))]})
                d_925_v56_: _dafny.Seq
                d_925_v56_ = _dafny.SeqWithoutIsStrInference([((self).f11)[default__.safeIndex(694, ((self).f11).length(0))], ((self).f11)[default__.safeIndex(694, ((self).f11).length(0))], (self).f8, len(d_924_v55_)])
                d_926_v57_: _dafny.Seq
                d_926_v57_ = _dafny.SeqWithoutIsStrInference([default__.fm3(globalState)])
                d_927_v58_: _dafny.Set
                d_927_v58_ = _dafny.Set({d_925_v56_})
                index143_ = default__.safeIndex(694, ((self).f11).length(0))
                rhs127_ = not(((self).f9 if not(default__.fm7(_dafny.CodePoint('u'), ((self).f11)[default__.safeIndex(694, ((self).f11).length(0))], p0, globalState)) else ((d_925_v56_)[default__.safeIndex((self).f8, len(d_925_v56_))]) < (-550)))
                rhs128_ = default__.fm7(self.f4, (0) - ((0) - (len((d_926_v57_) + ((d_926_v57_).set(default__.safeIndex(len(d_926_v57_), len(d_926_v57_)), not((self).f6)))))), not((self).f6), globalState)
                rhs129_ = (self).f8
                rhs130_ = ((self).f11)[default__.safeIndex(694, ((self).f11).length(0))]
                rhs131_ = (d_927_v58_) == (d_927_v58_)
                lhs95_ = globalState
                lhs96_ = (self).f11
                lhs97_ = default__.safeIndex(694, ((self).f11).length(0))
                d_923_v54_ = rhs127_
                d_923_v54_ = rhs128_
                lhs95_.f1 = rhs129_
                lhs96_[lhs97_] = rhs130_
                d_923_v54_ = rhs131_
                d_923_v54_ = (d_924_v55_).issubset(d_924_v55_)
                d_923_v54_ = p0
            d_928_v59_: _dafny.Seq
            d_928_v59_ = _dafny.SeqWithoutIsStrInference([(self).f9, (self).f9, (self).f6, (self).f6])
            if (d_928_v59_)[default__.safeIndex((self).f8, len(d_928_v59_))]:
                default__.m0(globalState)
                d_929_v60_: _dafny.Map
                d_929_v60_ = _dafny.Map({(self).f6: (self).f7})
                d_930_v61_: _dafny.Map
                d_930_v61_ = _dafny.Map({self.f4: p1})
                d_931_v62_: T2
                nw177_ = C3()
                nw177_.ctor__((self).f7, (self).f7, (self).f8, p1, False, p3, self.f4, (self).f5)
                d_931_v62_ = nw177_
                d_932_v63_: _dafny.MultiSet
                d_932_v63_ = _dafny.MultiSet([(self).fm1((self).f7, d_929_v60_, (self).f7, globalState), len(_dafny.Map({False: d_931_v62_}))])
                d_933_v64_: _dafny.Seq
                d_933_v64_ = _dafny.SeqWithoutIsStrInference([(self).f8, (d_932_v63_).cardinality])
                d_934_v65_: _dafny.Map
                d_934_v65_ = _dafny.Map({p1: len(d_933_v64_)})
                d_935_v66_: _dafny.Map
                d_935_v66_ = _dafny.Map({(self).f8: (p2).cardinality})
                d_936_v67_: _dafny.Seq
                d_936_v67_ = _dafny.SeqWithoutIsStrInference([d_929_v60_])
                d_937_v68_: _dafny.Map
                d_937_v68_ = _dafny.Map({237: (d_936_v67_)[default__.safeIndex((d_931_v62_).f8, len(d_936_v67_))]})
                d_938_v69_: _dafny.Array
                nw178_ = _dafny.Array(None, 29)
                nw178_[int(0)] = (self).f8
                nw178_[int(1)] = ((self).fm1(p0, d_929_v60_, p3, globalState)) + ((self).f8)
                nw178_[int(2)] = (self).f8
                nw178_[int(3)] = (self).f8
                nw178_[int(4)] = (self).f8
                nw178_[int(5)] = (self).f8
                nw178_[int(6)] = (self).f8
                nw178_[int(7)] = (self).f8
                nw178_[int(8)] = default__.safeDivisionInt((self).f8, len(d_930_v61_))
                nw178_[int(9)] = (self).f8
                nw178_[int(10)] = ((self).f8) + (len(d_934_v65_))
                nw178_[int(11)] = len((((self).f5).set(default__.safeIndex((self).f8, len((self).f5)), d_931_v62_.f4)) + ((self).f5))
                nw178_[int(12)] = (self).f8
                nw178_[int(13)] = (0) - ((d_931_v62_).f8)
                nw178_[int(14)] = (d_931_v62_).f8
                nw178_[int(15)] = ((d_935_v66_)[(d_931_v62_).f8] if ((d_931_v62_).f8) in (d_935_v66_) else 337)
                nw178_[int(16)] = (self).fm1((self).f6, d_929_v60_, (self).f7, globalState)
                nw178_[int(17)] = (self).f8
                nw178_[int(18)] = ((self).f8) * ((self).f8)
                nw178_[int(19)] = len((((d_937_v68_)[(self).f8] if ((self).f8) in (d_937_v68_) else d_929_v60_)).set(not(not((self).f6)), (d_931_v62_).f7))
                nw178_[int(20)] = (d_931_v62_).f8
                nw178_[int(21)] = ((p2) - (_dafny.MultiSet(d_928_v59_))).cardinality
                nw178_[int(22)] = (d_931_v62_).f8
                nw178_[int(23)] = (self).f8
                nw178_[int(24)] = ((self).f8) * ((d_931_v62_).f8)
                nw178_[int(25)] = len((d_934_v65_) | (d_934_v65_))
                nw178_[int(26)] = (0) - ((self).f8)
                nw178_[int(27)] = (self).f8
                nw178_[int(28)] = (d_931_v62_).f8
                d_938_v69_ = nw178_
                d_938_v69_ = d_938_v69_
                d_939_v70_: _dafny.Map
                d_939_v70_ = _dafny.Map({(d_931_v62_).f8: default__.fm26(p3, (d_931_v62_).f8, d_931_v62_.f4, globalState)})
                d_940_v71_: _dafny.Map
                d_940_v71_ = _dafny.Map({self.f4: d_939_v70_})
                d_941_v72_: _dafny.Array
                nw179_ = _dafny.Array(None, 7)
                nw179_[int(0)] = (d_940_v71_).set(d_931_v62_.f4, d_939_v70_)
                nw179_[int(1)] = d_940_v71_
                nw179_[int(2)] = (d_940_v71_).set(self.f4, d_939_v70_)
                nw179_[int(3)] = d_940_v71_
                nw179_[int(4)] = d_940_v71_
                nw179_[int(5)] = d_940_v71_
                nw179_[int(6)] = d_940_v71_
                d_941_v72_ = nw179_
                index144_ = default__.safeIndex(398, (d_941_v72_).length(0))
                (d_941_v72_)[index144_] = d_940_v71_
                index145_ = default__.safeIndex(398, (d_941_v72_).length(0))
                (d_941_v72_)[index145_] = (d_940_v71_) | (d_940_v71_)
                d_942_v73_: _dafny.Array
                nw180_ = _dafny.Array(_dafny.MultiSet({}), 29)
                d_942_v73_ = nw180_
                d_942_v73_ = d_942_v73_
                d_943_v74_: _dafny.Map
                d_943_v74_ = _dafny.Map({(self).f8: (self).f9})
                d_944_v75_: _dafny.Map
                d_944_v75_ = _dafny.Map({d_943_v74_: p1})
                index146_ = default__.safeIndex(223, ((self).f11).length(0))
                ((self).f11)[index146_] = (d_931_v62_).f8
                index147_ = default__.safeIndex(223, ((self).f11).length(0))
                rhs132_ = default__.fm8(default__.safeDivisionInt((self).f8, (self).f8), False, ((d_931_v62_).f8) >= ((d_931_v62_).f8), globalState)
                rhs133_ = (d_931_v62_).f8
                rhs134_ = d_944_v75_
                rhs135_ = default__.safeModuloInt((self).fm1(p3, _dafny.Map({(d_931_v62_).f9: (d_931_v62_).f6}), (self).f9, globalState), (d_931_v62_).f8)
                rhs136_ = default__.safeModuloInt(((self).f8) * ((d_931_v62_).f8), (self).f8)
                lhs98_ = globalState
                lhs99_ = (self).f11
                lhs100_ = default__.safeIndex(223, ((self).f11).length(0))
                r0 = rhs132_
                lhs98_.f1 = rhs133_
                d_944_v75_ = rhs134_
                lhs99_[lhs100_] = rhs135_
                r0 = rhs136_
            elif True:
                d_945_v76_: bool
                d_945_v76_ = True
                d_945_v76_ = True
                d_945_v76_ = not(((self).f5) < ((self).f5))
                index148_ = default__.safeIndex(463, ((self).f10).length(0))
                ((self).f10)[index148_] = p1
                index149_ = default__.safeIndex(463, ((self).f10).length(0))
                ((self).f10)[index149_] = False
                d_946_v77_: C0
                nw181_ = C0()
                nw181_.ctor__(d_945_v76_, ((self).f10)[default__.safeIndex(463, ((self).f10).length(0))], self.f4, _dafny.SeqWithoutIsStrInference([self.f4 for d_947_i2_ in range(default__.abs(38))]))
                d_946_v77_ = nw181_
                d_948_v78_: _dafny.Map
                d_948_v78_ = _dafny.Map({d_946_v77_: d_945_v76_})
                d_949_v79_: _dafny.Map
                d_949_v79_ = _dafny.Map({(self).f8: ((d_948_v78_)[d_946_v77_] if (d_946_v77_) in (d_948_v78_) else not((self).f7))})
                (self).m4((self).f8, d_949_v79_, globalState)
                index150_ = default__.safeIndex(463, ((self).f10).length(0))
                ((self).f10)[index150_] = not((d_928_v59_)[default__.safeIndex((self).f8, len(d_928_v59_))])
            (self).f4 = self.f4
            d_950_v80_: bool
            d_950_v80_ = False
            d_950_v80_ = p0
            d_951_v81_: _dafny.Seq
            d_951_v81_ = _dafny.SeqWithoutIsStrInference([(self).f8, (0) - ((self).f8)])
            d_952_v82_: _dafny.Set
            d_952_v82_ = _dafny.Set({(self).f6, (self).f7})
            d_953_v83_: D3
            d_953_v83_ = D3_DC11(False, len(d_951_v81_), d_952_v82_)
            d_954_v84_: _dafny.Seq
            d_954_v84_ = _dafny.SeqWithoutIsStrInference([(d_953_v83_).cf22])
            d_955_v85_: D0
            d_955_v85_ = D0_DC1(d_950_v80_, -224, (self).f6, (self).f8)
            (globalState).f1 = (len(d_954_v84_)) * (default__.fm8((d_955_v85_).cf3, not(p3), p1, globalState))
        d_956_v86_: _dafny.MultiSet
        d_956_v86_ = _dafny.MultiSet([(self).f8, (self).f8])
        d_957_v87_: D6
        d_957_v87_ = D6_DC16(d_956_v86_)
        r0 = (0) - (((d_957_v87_).cf30).cardinality)
        r1 = (0) - ((self).f8)
        if (p2).isdisjoint(p2):
            d_958_v88_: _dafny.Array
            nw182_ = _dafny.Array(_dafny.Set({}), 11)
            d_958_v88_ = nw182_
            d_959_v89_: _dafny.Set
            d_959_v89_ = _dafny.Set({p0, True})
            index151_ = default__.safeIndex(828, (d_958_v88_).length(0))
            (d_958_v88_)[index151_] = (default__.fm0((self).f9, 924, (p2).cardinality, p3, globalState) if (self).f6 else d_959_v89_)
            d_960_v90_: bool
            d_960_v90_ = False
            d_961_v91_: T3
            nw183_ = C4()
            nw183_.ctor__(644, p3, not((self).f6), False, self.f4, (self).f5)
            d_961_v91_ = nw183_
            d_962_v92_: _dafny.Map
            d_962_v92_ = _dafny.Map({d_961_v91_: (self).f11})
            d_963_v93_: _dafny.Array
            nw184_ = _dafny.Array(None, 27)
            nw184_[int(0)] = (self).f11
            nw184_[int(1)] = (self).f11
            nw184_[int(2)] = (self).f11
            nw184_[int(3)] = (self).f11
            nw184_[int(4)] = (self).f11
            nw184_[int(5)] = (self).f11
            nw184_[int(6)] = (self).f11
            nw184_[int(7)] = (self).f11
            nw184_[int(8)] = (self).f11
            nw184_[int(9)] = (self).f11
            nw184_[int(10)] = (self).f11
            nw184_[int(11)] = (self).f11
            nw184_[int(12)] = (self).f11
            nw184_[int(13)] = (self).f11
            nw184_[int(14)] = (self).f11
            nw184_[int(15)] = (self).f11
            nw184_[int(16)] = (self).f11
            nw184_[int(17)] = (self).f11
            nw184_[int(18)] = (self).f11
            nw184_[int(19)] = (self).f11
            nw184_[int(20)] = (self).f11
            nw184_[int(21)] = (self).f11
            nw184_[int(22)] = ((self).f11 if p0 else (self).f11)
            nw184_[int(23)] = (self).f11
            nw184_[int(24)] = (self).f11
            nw184_[int(25)] = ((d_962_v92_)[d_961_v91_] if (d_961_v91_) in (d_962_v92_) else (self).f11)
            nw184_[int(26)] = (self).f11
            d_963_v93_ = nw184_
            index152_ = default__.safeIndex(451, (d_963_v93_).length(0))
            (d_963_v93_)[index152_] = (self).f11
            d_964_v94_: C0
            nw185_ = C0()
            nw185_.ctor__(p1, p1, d_961_v91_.f4, ((d_961_v91_).f5) + ((self).f5))
            d_964_v94_ = nw185_
            index153_ = default__.safeIndex(828, (d_958_v88_).length(0))
            index154_ = default__.safeIndex(451, (d_963_v93_).length(0))
            rhs137_ = d_959_v89_
            rhs138_ = d_961_v91_.f4
            rhs139_ = not((self).f7)
            rhs140_ = ((self).f11 if ((self).f8) >= ((self).f8) else (self).f11)
            rhs141_ = d_964_v94_
            lhs101_ = d_958_v88_
            lhs102_ = default__.safeIndex(828, (d_958_v88_).length(0))
            lhs103_ = self
            lhs104_ = d_963_v93_
            lhs105_ = default__.safeIndex(451, (d_963_v93_).length(0))
            lhs101_[lhs102_] = rhs137_
            lhs103_.f4 = rhs138_
            d_960_v90_ = rhs139_
            lhs104_[lhs105_] = rhs140_
            d_964_v94_ = rhs141_
            arr0_ = (d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]
            index155_ = default__.safeIndex(341, ((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]).length(0))
            arr0_[index155_] = (self).f8
            d_965_v95_: _dafny.Map
            d_965_v95_ = _dafny.Map({d_956_v86_: (d_961_v91_).f9})
            arr1_ = (d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]
            index156_ = default__.safeIndex(341, ((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]).length(0))
            arr1_[index156_] = len((d_965_v95_) | (d_965_v95_))
            (globalState).f1 = ((self).f8) * ((d_961_v91_).f8)
            if (d_961_v91_).f9:
                (d_961_v91_).f4 = self.f4
                d_966_v96_: D1
                d_966_v96_ = D1_DC4(d_961_v91_.f4, (self).f8, d_961_v91_.f4, (self).f8, (d_961_v91_).f6)
                d_967_v97_: _dafny.Seq
                d_967_v97_ = _dafny.SeqWithoutIsStrInference([(self).f8])
                d_968_v98_: C4
                nw186_ = C4()
                nw186_.ctor__((0) - ((d_961_v91_).f8), ((d_961_v91_).f9) and (False), (d_966_v96_).cf12, (d_956_v86_).issubset(_dafny.MultiSet([(d_961_v91_).f8, (self).f8])), ((d_961_v91_).f5)[default__.safeIndex(len(d_967_v97_), len((d_961_v91_).f5))], (self).f5)
                d_968_v98_ = nw186_
                d_969_v99_: _dafny.Seq
                d_969_v99_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                index157_ = default__.safeIndex(492, ((self).f10).length(0))
                ((self).f10)[index157_] = (((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))])[default__.safeIndex(341, ((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]).length(0))]) > (((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))])[default__.safeIndex(341, ((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]).length(0))])
                d_970_v100_: _dafny.Seq
                d_970_v100_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                d_971_v101_: T2
                nw187_ = C1()
                nw187_.ctor__((self).f8, (d_970_v100_)[default__.safeIndex((self).f8, len(d_970_v100_))], (self).f6, p3, d_961_v91_.f4, (self).f5)
                d_971_v101_ = nw187_
                d_972_v102_: D0
                d_972_v102_ = D0_DC1(not(p1), (d_961_v91_).f8, p0, (d_961_v91_).f8)
                d_973_v103_: _dafny.Set
                d_973_v103_ = _dafny.Set({d_972_v102_, d_972_v102_})
                d_974_v104_: _dafny.Set
                d_974_v104_ = _dafny.Set({(self).f8})
                d_975_v105_: _dafny.Map
                d_975_v105_ = _dafny.Map({(d_971_v101_ if p0 else d_971_v101_): len((default__.fm14(len(d_973_v103_), len((d_971_v101_).f5), (d_961_v91_).f6, (d_971_v101_).f9, globalState)) + (((d_971_v101_).f5).set(default__.safeIndex(718, len((d_971_v101_).f5)), ((self).f5)[default__.safeIndex(len(d_974_v104_), len((self).f5))])))})
                index158_ = default__.safeIndex(492, ((self).f10).length(0))
                index159_ = default__.safeIndex(451, (d_963_v93_).length(0))
                rhs142_ = d_969_v99_
                rhs143_ = (d_971_v101_).f6
                rhs144_ = (d_975_v105_) | (d_975_v105_)
                rhs145_ = (d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]
                rhs146_ = (len((d_970_v100_) + (d_970_v100_))) * ((d_961_v91_).f8)
                lhs106_ = (self).f10
                lhs107_ = default__.safeIndex(492, ((self).f10).length(0))
                lhs108_ = d_963_v93_
                lhs109_ = default__.safeIndex(451, (d_963_v93_).length(0))
                lhs110_ = globalState
                d_969_v99_ = rhs142_
                lhs106_[lhs107_] = rhs143_
                d_975_v105_ = rhs144_
                lhs108_[lhs109_] = rhs145_
                lhs110_.f1 = rhs146_
                (self).f4 = d_971_v101_.f4
                d_976_v106_: _dafny.Array
                nw188_ = _dafny.Array(None, 21)
                nw188_[int(0)] = p3
                nw188_[int(1)] = default__.fm3(globalState)
                nw188_[int(2)] = p1
                nw188_[int(3)] = p1
                nw188_[int(4)] = (d_961_v91_).f7
                nw188_[int(5)] = not(not(p1))
                nw188_[int(6)] = (self).f6
                nw188_[int(7)] = p3
                nw188_[int(8)] = False
                nw188_[int(9)] = default__.fm3(globalState)
                nw188_[int(10)] = (self).f7
                nw188_[int(11)] = (self).f9
                nw188_[int(12)] = False
                nw188_[int(13)] = p3
                nw188_[int(14)] = False
                nw188_[int(15)] = (d_961_v91_).f6
                nw188_[int(16)] = (d_961_v91_).f9
                nw188_[int(17)] = default__.fm7(d_961_v91_.f4, -833, p1, globalState)
                nw188_[int(18)] = (d_961_v91_).f9
                nw188_[int(19)] = (d_961_v91_).f7
                nw188_[int(20)] = (d_961_v91_).f6
                d_976_v106_ = nw188_
                d_977_v107_: _dafny.Array
                def lambda33_(d_978_v100_):
                    def lambda34_(d_979_i3_):
                        return (D8_DC22(d_978_v100_)).cf41

                    return lambda34_

                init18_ = lambda33_(d_970_v100_)
                nw189_ = _dafny.Array(None, 6)
                for i0_18_ in range(nw189_.length(0)):
                    nw189_[i0_18_] = init18_(i0_18_)
                d_977_v107_ = nw189_
                d_980_v108_: D1
                d_980_v108_ = D1_DC3(d_977_v107_)
                pat_let_tv16_ = d_977_v107_
                d_981_v109_: _dafny.Array
                nw190_ = _dafny.Array(None, 25)
                nw190_[int(0)] = d_980_v108_
                nw190_[int(1)] = d_980_v108_
                nw190_[int(2)] = d_980_v108_
                nw190_[int(3)] = D1_DC3(d_977_v107_)
                nw190_[int(4)] = d_980_v108_
                nw190_[int(5)] = d_980_v108_
                nw190_[int(6)] = d_980_v108_
                nw190_[int(7)] = d_980_v108_
                nw190_[int(8)] = d_980_v108_
                nw190_[int(9)] = d_980_v108_
                nw190_[int(10)] = d_980_v108_
                nw190_[int(11)] = D1_DC3(d_977_v107_)
                nw190_[int(12)] = d_980_v108_
                nw190_[int(13)] = d_980_v108_
                nw190_[int(14)] = d_980_v108_
                def iife72_(_pat_let26_0):
                    def iife73_(d_982_dt__update__tmp_h2_):
                        def iife74_(_pat_let27_0):
                            def iife75_(d_983_dt__update_hcf7_h0_):
                                return D1_DC3(d_983_dt__update_hcf7_h0_)
                            return iife75_(_pat_let27_0)
                        return iife74_(pat_let_tv16_)
                    return iife73_(_pat_let26_0)
                nw190_[int(15)] = iife72_(D1_DC3(d_977_v107_))
                nw190_[int(16)] = D1_DC3(d_977_v107_)
                nw190_[int(17)] = D1_DC3(d_977_v107_)
                nw190_[int(18)] = d_980_v108_
                nw190_[int(19)] = D1_DC3(d_977_v107_)
                nw190_[int(20)] = d_980_v108_
                nw190_[int(21)] = d_980_v108_
                nw190_[int(22)] = d_980_v108_
                nw190_[int(23)] = D1_DC3(d_977_v107_)
                nw190_[int(24)] = d_980_v108_
                d_981_v109_ = nw190_
                d_984_v110_: C2
                nw191_ = C2()
                nw191_.ctor__(d_976_v106_, d_981_v109_)
                d_984_v110_ = nw191_
                d_985_v111_: _dafny.Set
                d_985_v111_ = _dafny.Set({(d_961_v91_).f5})
                index160_ = default__.safeIndex(492, ((self).f10).length(0))
                rhs147_ = (d_985_v111_).issubset((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhuqmt")), _dafny.SeqWithoutIsStrInference([d_961_v91_.f4 for d_986_i4_ in range(default__.abs(737))])}) if True else d_985_v111_))
                rhs148_ = d_984_v110_
                lhs111_ = (self).f10
                lhs112_ = default__.safeIndex(492, ((self).f10).length(0))
                lhs111_[lhs112_] = rhs147_
                d_984_v110_ = rhs148_
            elif True:
                d_987_v112_: _dafny.Seq
                d_987_v112_ = _dafny.SeqWithoutIsStrInference([910])
                d_988_v113_: _dafny.Seq
                d_988_v113_ = _dafny.SeqWithoutIsStrInference([len(d_987_v112_)])
                d_989_v114_: _dafny.Map
                d_989_v114_ = _dafny.Map({(self).f10: (d_956_v86_).isdisjoint(_dafny.MultiSet(d_988_v113_))})
                d_990_v115_: _dafny.Map
                d_990_v115_ = _dafny.Map({(d_961_v91_).f6: (self).f6})
                d_989_v114_ = (d_989_v114_).set((self).f10, ((d_990_v115_)[False] if (False) in (d_990_v115_) else (d_961_v91_).f9))
                (globalState).f1 = (d_961_v91_).f8
                d_991_v116_: _dafny.Array
                def lambda35_(d_992_v91_, d_993_v93_):
                    def lambda36_(d_994_i5_):
                        return (_dafny.Map({D1_DC4(d_992_v91_.f4, (d_992_v91_).f8, d_992_v91_.f4, (self).f8, (d_992_v91_).f7): (d_992_v91_).f8})) | (_dafny.Map({D1_DC4(d_992_v91_.f4, ((d_993_v93_)[default__.safeIndex(451, (d_993_v93_).length(0))])[default__.safeIndex(341, ((d_993_v93_)[default__.safeIndex(451, (d_993_v93_).length(0))]).length(0))], self.f4, (d_992_v91_).f8, (d_992_v91_).f6): ((d_993_v93_)[default__.safeIndex(451, (d_993_v93_).length(0))])[default__.safeIndex(341, ((d_993_v93_)[default__.safeIndex(451, (d_993_v93_).length(0))]).length(0))]}))

                    return lambda36_

                init19_ = lambda35_(d_961_v91_, d_963_v93_)
                nw192_ = _dafny.Array(None, 1)
                for i0_19_ in range(nw192_.length(0)):
                    nw192_[i0_19_] = init19_(i0_19_)
                d_991_v116_ = nw192_
                d_995_v117_: D1
                d_995_v117_ = D1_DC4(d_961_v91_.f4, (self).f8, d_961_v91_.f4, (self).f8, False)
                d_996_v118_: _dafny.Map
                d_996_v118_ = _dafny.Map({d_995_v117_: (d_961_v91_).f8})
                index161_ = default__.safeIndex(893, (d_991_v116_).length(0))
                (d_991_v116_)[index161_] = d_996_v118_
                d_997_v119_: _dafny.Seq
                d_997_v119_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ciyeo"))
                d_998_v121_: T1
                nw193_ = C0()
                nw193_.ctor__(default__.fm3(globalState), d_960_v90_, _dafny.CodePoint('f'), (d_961_v91_).f5)
                d_998_v121_ = nw193_
                d_999_v122_: _dafny.MultiSet
                d_999_v122_ = _dafny.MultiSet([d_998_v121_])
                d_1000_v123_: _dafny.Map
                d_1000_v123_ = _dafny.Map({((d_961_v91_).f8) * (((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))])[default__.safeIndex(341, ((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]).length(0))]): (self).f5})
                index162_ = default__.safeIndex(893, (d_991_v116_).length(0))
                def iife76_():
                    coll20_ = _dafny.Map()
                    compr_20_: D1
                    for compr_20_ in (_dafny.MultiSet([D1_DC4(self.f4, (d_999_v122_).cardinality, d_998_v121_.f4, len((d_961_v91_).f5), p0)])).Elements:
                        d_1001_v120_: D1 = compr_20_
                        if (d_1001_v120_) in (_dafny.MultiSet([D1_DC4(self.f4, (d_999_v122_).cardinality, d_998_v121_.f4, len((d_961_v91_).f5), p0)])):
                            coll20_[d_1001_v120_] = ((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))])[default__.safeIndex(341, ((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]).length(0))]
                    return _dafny.Map(coll20_)
                rhs149_ = (iife76_()
                 if p0 else _dafny.Map({d_995_v117_: (self).f8}))
                rhs150_ = (((self).f8) + ((d_961_v91_).f8)) >= ((len((d_961_v91_).f5)) * ((d_961_v91_).f8))
                rhs151_ = ((d_1000_v123_)[(d_961_v91_).f8] if ((d_961_v91_).f8) in (d_1000_v123_) else (d_998_v121_).f5)
                lhs113_ = d_991_v116_
                lhs114_ = default__.safeIndex(893, (d_991_v116_).length(0))
                lhs113_[lhs114_] = rhs149_
                d_960_v90_ = rhs150_
                d_997_v119_ = rhs151_
                index163_ = default__.safeIndex(653, ((self).f10).length(0))
                ((self).f10)[index163_] = not(((d_961_v91_).f7 if (self).f9 else (d_961_v91_).f9))
                index164_ = default__.safeIndex(653, ((self).f10).length(0))
                ((self).f10)[index164_] = (d_961_v91_).f6
                d_1002_v124_: _dafny.Map
                d_1002_v124_ = _dafny.Map({(d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]: (0) - ((d_961_v91_).f8)})
                index165_ = default__.safeIndex(653, ((self).f10).length(0))
                arr2_ = (d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]
                index166_ = default__.safeIndex(341, ((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]).length(0))
                rhs152_ = (d_998_v121_).f7
                rhs153_ = (len((d_961_v91_).f5)) > (len(d_1002_v124_))
                rhs154_ = d_997_v119_
                rhs155_ = 339
                lhs115_ = (self).f10
                lhs116_ = default__.safeIndex(653, ((self).f10).length(0))
                lhs117_ = (d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]
                lhs118_ = default__.safeIndex(341, ((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]).length(0))
                lhs115_[lhs116_] = rhs152_
                d_960_v90_ = rhs153_
                d_997_v119_ = rhs154_
                lhs117_[lhs118_] = rhs155_
            arr3_ = (d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]
            index167_ = default__.safeIndex(341, ((d_963_v93_)[default__.safeIndex(451, (d_963_v93_).length(0))]).length(0))
            arr3_[index167_] = (d_961_v91_).f8
        elif True:
            d_1003_v125_: _dafny.MultiSet
            d_1003_v125_ = _dafny.MultiSet([(self).f9])
            def iife77_(_pat_let28_0):
                def iife78_(d_1004_dt__update__tmp_h3_):
                    def iife79_(_pat_let29_0):
                        def iife80_(d_1005_dt__update_hcf43_h0_):
                            return D8_DC23((d_1004_dt__update__tmp_h3_).cf42, d_1005_dt__update_hcf43_h0_)
                        return iife80_(_pat_let29_0)
                    return iife79_(True)
                return iife78_(_pat_let28_0)
            d_1003_v125_ = _dafny.MultiSet([(iife77_(D8_DC23(937, p3))).cf43])
            d_1006_v126_: bool
            d_1006_v126_ = True
            d_1006_v126_ = p1
            index168_ = default__.safeIndex(757, ((self).f10).length(0))
            ((self).f10)[index168_] = ((self).f5) == ((self).f5)
            d_1007_v127_: _dafny.Array
            nw194_ = _dafny.Array(None, 9)
            nw194_[int(0)] = (self).f10
            nw194_[int(1)] = (self).f10
            nw194_[int(2)] = (self).f10
            nw194_[int(3)] = (self).f10
            nw194_[int(4)] = (self).f10
            nw194_[int(5)] = (self).f10
            nw194_[int(6)] = (self).f10
            nw194_[int(7)] = (self).f10
            nw194_[int(8)] = (self).f10
            d_1007_v127_ = nw194_
            d_1008_v128_: _dafny.Array
            nw195_ = _dafny.Array(D6.default()(), 1)
            d_1008_v128_ = nw195_
            index169_ = default__.safeIndex(432, (d_1008_v128_).length(0))
            (d_1008_v128_)[index169_] = d_957_v87_
            d_1009_v129_: _dafny.Seq
            d_1009_v129_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((0) - ((self).f8), (self).f8)])
            index170_ = default__.safeIndex(757, ((self).f10).length(0))
            index171_ = default__.safeIndex(432, (d_1008_v128_).length(0))
            rhs156_ = not((718) == ((self).f8))
            rhs157_ = d_1007_v127_
            rhs158_ = d_957_v87_
            rhs159_ = (d_1009_v129_) + (d_1009_v129_)
            lhs119_ = (self).f10
            lhs120_ = default__.safeIndex(757, ((self).f10).length(0))
            lhs121_ = d_1008_v128_
            lhs122_ = default__.safeIndex(432, (d_1008_v128_).length(0))
            lhs119_[lhs120_] = rhs156_
            d_1007_v127_ = rhs157_
            lhs121_[lhs122_] = rhs158_
            d_1009_v129_ = rhs159_
            d_1010_v130_: _dafny.Map
            d_1010_v130_ = _dafny.Map({(self).f8: (not(not(d_1006_v126_)) if True else True)})
            d_1011_v131_: _dafny.Array
            def lambda37_(d_1012_v126_):
                def lambda38_(d_1013_i6_):
                    return d_1012_v126_

                return lambda38_

            init20_ = lambda37_(d_1006_v126_)
            nw196_ = _dafny.Array(None, 27)
            for i0_20_ in range(nw196_.length(0)):
                nw196_[i0_20_] = init20_(i0_20_)
            d_1011_v131_ = nw196_
            d_1014_v132_: _dafny.Map
            d_1014_v132_ = _dafny.Map({p0: (self).f8})
            d_1015_v133_: D5
            d_1015_v133_ = D5_DC15(d_1011_v131_, p1, d_1014_v132_, (self).f8)
            index172_ = default__.safeIndex(757, ((self).f10).length(0))
            ((self).f10)[index172_] = ((d_1010_v130_)[((self).f8) - ((self).f8)] if (((self).f8) - ((self).f8)) in (d_1010_v130_) else (d_1015_v133_).cf27)
            d_1016_v134_: _dafny.Map
            d_1016_v134_ = _dafny.Map({(self).f8: -229})
            d_1017_v135_: C3
            nw197_ = C3()
            nw197_.ctor__(((self).f10)[default__.safeIndex(757, ((self).f10).length(0))], p1, ((d_1016_v134_)[(self).f8] if ((self).f8) in (d_1016_v134_) else (d_956_v86_).cardinality), (self).f6, d_1006_v126_, (self).f7, self.f4, (self).f5)
            d_1017_v135_ = nw197_
            d_1018_v136_: _dafny.Map
            d_1018_v136_ = _dafny.Map({-343: d_1017_v135_})
            d_1019_v137_: _dafny.Map
            d_1019_v137_ = _dafny.Map({(self).f6: ((d_1010_v130_)[(self).f8] if ((self).f8) in (d_1010_v130_) else (d_1017_v135_).f13)})
            d_1018_v136_ = (d_1018_v136_).set((self).fm1(p0, d_1019_v137_, p0, globalState), (d_1017_v135_ if (self).f6 else d_1017_v135_))
        r1 = default__.safeDivisionInt((self).f8, (self).f8)
        d_1020_v138_: _dafny.MultiSet
        d_1020_v138_ = _dafny.MultiSet([(self).f11, (self).f11, (self).f11, (self).f11])
        d_1021_v139_: _dafny.Seq
        d_1021_v139_ = _dafny.SeqWithoutIsStrInference([d_1020_v138_])
        r0 = (0) - (((d_1020_v138_).intersection(((d_1021_v139_)[default__.safeIndex((self).f8, len(d_1021_v139_))]) | (d_1020_v138_))).cardinality)
        r1 = (self).f8
        d_1022_v140_: _dafny.Array
        nw198_ = _dafny.Array(_dafny.Array(None, 0), 16)
        d_1022_v140_ = nw198_
        r2 = d_1022_v140_
        return r0, r1, r2

    def m2(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_1023_v0_: _dafny.Seq
        d_1023_v0_ = _dafny.SeqWithoutIsStrInference([(self).f7])
        d_1024_v1_: _dafny.Seq
        d_1024_v1_ = _dafny.SeqWithoutIsStrInference([False, (self).f6, (d_1023_v0_)[default__.safeIndex(p0, len(d_1023_v0_))], True])
        d_1025_v2_: _dafny.MultiSet
        d_1025_v2_ = _dafny.MultiSet([(self).f7, (self).f9, (self).f6, (self).f9, True])
        if default__.fm27(default__.safeDivisionInt(len(d_1024_v1_), (d_1025_v2_).cardinality), (self).f6, (self).f7, globalState):
            d_1026_v3_: _dafny.Seq
            d_1026_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgo"))
            d_1027_v4_: _dafny.Set
            d_1027_v4_ = _dafny.Set({(self).f9})
            d_1028_v5_: D3
            d_1028_v5_ = D3_DC11((self).f6, (self).f8, d_1027_v4_)
            d_1029_v6_: _dafny.Map
            d_1029_v6_ = _dafny.Map({(self).f6: _dafny.SeqWithoutIsStrInference([self.f4 for d_1030_i0_ in range(default__.abs(281))])})
            d_1026_v3_ = ((default__.fm28((self).f8, (self).f7, default__.fm29(globalState), d_1028_v5_, globalState)) + (((d_1029_v6_)[(self).f9] if ((self).f9) in (d_1029_v6_) else d_1026_v3_))) + ((self).f5)
            d_1026_v3_ = d_1026_v3_
            d_1026_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xolxenqc"))
            d_1031_v7_: _dafny.Array
            nw199_ = _dafny.Array(D1.default()(), 22)
            d_1031_v7_ = nw199_
            d_1032_v8_: C2
            nw200_ = C2()
            nw200_.ctor__((self).f10, d_1031_v7_)
            d_1032_v8_ = nw200_
            index173_ = default__.safeIndex(111, ((d_1032_v8_).f14).length(0))
            ((d_1032_v8_).f14)[index173_] = (self).f7
            d_1033_v9_: _dafny.Array
            def lambda39_(d_1034_i1_):
                return D6_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xy")), _dafny.SeqWithoutIsStrInference([self.f4 for d_1035_i2_ in range(default__.abs(-640))]))

            init21_ = lambda39_
            nw201_ = _dafny.Array(None, 6)
            for i0_21_ in range(nw201_.length(0)):
                nw201_[i0_21_] = init21_(i0_21_)
            d_1033_v9_ = nw201_
            d_1036_v10_: _dafny.Map
            d_1036_v10_ = _dafny.Map({d_1033_v9_: d_1032_v8_})
            d_1037_v11_: _dafny.Seq
            d_1037_v11_ = _dafny.SeqWithoutIsStrInference([d_1033_v9_])
            d_1038_v12_: _dafny.Map
            d_1038_v12_ = _dafny.Map({(self).f6: len((self).f5)})
            index174_ = default__.safeIndex(111, ((d_1032_v8_).f14).length(0))
            rhs160_ = ((d_1036_v10_)[(d_1037_v11_)[default__.safeIndex(len(d_1038_v12_), len(d_1037_v11_))]] if ((d_1037_v11_)[default__.safeIndex(len(d_1038_v12_), len(d_1037_v11_))]) in (d_1036_v10_) else d_1032_v8_)
            rhs161_ = (self).f9
            lhs123_ = (d_1032_v8_).f14
            lhs124_ = default__.safeIndex(111, ((d_1032_v8_).f14).length(0))
            d_1032_v8_ = rhs160_
            lhs123_[lhs124_] = rhs161_
            d_1039_v13_: _dafny.Map
            d_1039_v13_ = _dafny.Map({d_1026_v3_: (self).f6})
            index175_ = default__.safeIndex(111, ((d_1032_v8_).f14).length(0))
            ((d_1032_v8_).f14)[index175_] = ((d_1039_v13_)[(d_1026_v3_) + (d_1026_v3_)] if ((d_1026_v3_) + (d_1026_v3_)) in (d_1039_v13_) else (self).f9)
        elif True:
            d_1040_v14_: bool
            d_1040_v14_ = False
            d_1040_v14_ = d_1040_v14_
            d_1041_v15_: _dafny.MultiSet
            d_1041_v15_ = _dafny.MultiSet([(self).f8])
            d_1042_v16_: _dafny.Seq
            d_1042_v16_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1043_v17_: _dafny.Map
            d_1043_v17_ = _dafny.Map({d_1042_v16_: p0})
            d_1044_v18_: _dafny.Map
            d_1044_v18_ = _dafny.Map({d_1040_v14_: ((d_1041_v15_)[len(d_1043_v17_)] if (len(d_1043_v17_)) in (d_1041_v15_) else p0)})
            d_1044_v18_ = d_1044_v18_
            d_1045_v19_: D3
            d_1045_v19_ = D3_DC10((self).f5)
            d_1046_v20_: D3
            d_1046_v20_ = D3_DC12(d_1045_v19_)
            source11_ = d_1046_v20_
            if source11_.is_DC11:
                d_1047___mcc_h0_ = source11_.cf20
                d_1048___mcc_h1_ = source11_.cf21
                d_1049___mcc_h2_ = source11_.cf22
                d_1050_cf22_ = d_1049___mcc_h2_
                d_1051_cf21_ = d_1048___mcc_h1_
                d_1052_cf20_ = d_1047___mcc_h0_
                d_1053_v21_: C0
                nw202_ = C0()
                nw202_.ctor__((self).f6, d_1052_cf20_, (self.f4 if False else _dafny.CodePoint('v')), ((self).f5) + ((self).f5))
                d_1053_v21_ = nw202_
                d_1054_v22_: _dafny.Array
                nw203_ = _dafny.Array(D1.default()(), 16)
                d_1054_v22_ = nw203_
                d_1055_v23_: C2
                nw204_ = C2()
                nw204_.ctor__((self).f10, d_1054_v22_)
                d_1055_v23_ = nw204_
                d_1051_cf21_ = (self).f8
                (globalState).f1 = (self).f8
            elif source11_.is_DC10:
                d_1056___mcc_h3_ = source11_.cf19
                d_1057_cf19_ = d_1056___mcc_h3_
                d_1058_v24_: _dafny.Set
                d_1058_v24_ = _dafny.Set({(self).f10})
                d_1059_v25_: _dafny.Map
                d_1059_v25_ = _dafny.Map({(self).f8: False})
                d_1060_v26_: D0
                d_1060_v26_ = D0_DC1((self).f7, ((d_1041_v15_).cardinality) - ((self).f8), default__.fm7(self.f4, len(d_1058_v24_), (self).f9, globalState), len(d_1059_v25_))
                rhs162_ = d_1060_v26_
                rhs163_ = d_1041_v15_
                rhs164_ = False
                d_1060_v26_ = rhs162_
                d_1041_v15_ = rhs163_
                d_1040_v14_ = rhs164_
                index176_ = default__.safeIndex(273, ((self).f10).length(0))
                ((self).f10)[index176_] = (self).f9
                index177_ = default__.safeIndex(273, ((self).f10).length(0))
                ((self).f10)[index177_] = ((d_1041_v15_).set((self).f8, default__.abs((self).f8))).issubset(d_1041_v15_)
                d_1061_v27_: C4
                nw205_ = C4()
                nw205_.ctor__(-331, (self).f6, ((self).f10)[default__.safeIndex(273, ((self).f10).length(0))], ((self).f8) == (629), self.f4, _dafny.SeqWithoutIsStrInference([self.f4 for d_1062_i3_ in range(default__.abs(112))]))
                d_1061_v27_ = nw205_
                d_1063_v28_: D0
                d_1063_v28_ = D0_DC2(len(d_1042_v16_), (self).f6, _dafny.SeqWithoutIsStrInference([(self).f8, p0]))
                d_1064_v29_: _dafny.Map
                d_1064_v29_ = _dafny.Map({(d_1063_v28_).cf6: _dafny.Map({(d_1041_v15_).cardinality: p0})})
                index178_ = default__.safeIndex(662, ((self).f11).length(0))
                ((self).f11)[index178_] = (len(d_1064_v29_)) * ((self).f8)
                d_1065_v30_: _dafny.Set
                d_1065_v30_ = _dafny.Set({_dafny.MultiSet(d_1023_v0_), d_1025_v2_, d_1025_v2_})
                index179_ = default__.safeIndex(662, ((self).f11).length(0))
                ((self).f11)[index179_] = (len(d_1065_v30_)) - ((self).f8)
            elif True:
                d_1066___mcc_h4_ = source11_.cf23
                d_1067_cf23_ = d_1066___mcc_h4_
                d_1068_v31_: _dafny.Array
                nw206_ = _dafny.Array(int(0), 22)
                d_1068_v31_ = nw206_
                d_1068_v31_ = (self).f11
                d_1069_v32_: D8
                d_1069_v32_ = D8_DC22(_dafny.SeqWithoutIsStrInference([(self).f6, (self).f7]))
                d_1070_v33_: _dafny.Array
                nw207_ = _dafny.Array(None, 29)
                nw207_[int(0)] = d_1024_v1_
                nw207_[int(1)] = d_1023_v0_
                nw207_[int(2)] = d_1024_v1_
                nw207_[int(3)] = (d_1069_v32_).cf41
                nw207_[int(4)] = default__.fm18((self).f8, (self).f6, globalState)
                nw207_[int(5)] = d_1024_v1_
                nw207_[int(6)] = d_1023_v0_
                nw207_[int(7)] = d_1023_v0_
                nw207_[int(8)] = d_1024_v1_
                nw207_[int(9)] = d_1023_v0_
                nw207_[int(10)] = d_1023_v0_
                nw207_[int(11)] = _dafny.SeqWithoutIsStrInference([d_1040_v14_])
                nw207_[int(12)] = d_1023_v0_
                nw207_[int(13)] = d_1023_v0_
                nw207_[int(14)] = _dafny.SeqWithoutIsStrInference([True])
                nw207_[int(15)] = d_1024_v1_
                nw207_[int(16)] = d_1024_v1_
                nw207_[int(17)] = d_1023_v0_
                nw207_[int(18)] = d_1024_v1_
                nw207_[int(19)] = d_1023_v0_
                nw207_[int(20)] = _dafny.SeqWithoutIsStrInference([d_1040_v14_])
                nw207_[int(21)] = d_1023_v0_
                nw207_[int(22)] = d_1024_v1_
                nw207_[int(23)] = d_1023_v0_
                nw207_[int(24)] = d_1024_v1_
                nw207_[int(25)] = d_1023_v0_
                nw207_[int(26)] = d_1024_v1_
                nw207_[int(27)] = d_1023_v0_
                nw207_[int(28)] = d_1023_v0_
                d_1070_v33_ = nw207_
                d_1071_v34_: D1
                d_1071_v34_ = D1_DC3(d_1070_v33_)
                d_1072_v35_: _dafny.Seq
                d_1072_v35_ = _dafny.SeqWithoutIsStrInference([d_1071_v34_])
                d_1073_v36_: D9
                d_1073_v36_ = D9_DC24(d_1072_v35_)
                d_1074_v37_: _dafny.Array
                nw208_ = _dafny.Array(None, 23)
                nw208_[int(0)] = ((d_1072_v35_) + (d_1072_v35_)).set(default__.safeIndex((self).f8, len((d_1072_v35_) + (d_1072_v35_))), d_1071_v34_)
                nw208_[int(1)] = d_1072_v35_
                nw208_[int(2)] = d_1072_v35_
                nw208_[int(3)] = d_1072_v35_
                nw208_[int(4)] = (d_1072_v35_).set(default__.safeIndex((self).f8, len(d_1072_v35_)), d_1071_v34_)
                nw208_[int(5)] = (d_1072_v35_) + (d_1072_v35_)
                nw208_[int(6)] = d_1072_v35_
                nw208_[int(7)] = d_1072_v35_
                nw208_[int(8)] = d_1072_v35_
                nw208_[int(9)] = ((_dafny.SeqWithoutIsStrInference([d_1071_v34_, D1_DC3(d_1070_v33_), d_1071_v34_, d_1071_v34_])).set(default__.safeIndex((self).f8, len(_dafny.SeqWithoutIsStrInference([d_1071_v34_, D1_DC3(d_1070_v33_), d_1071_v34_, d_1071_v34_]))), D1_DC3(d_1070_v33_))) + (d_1072_v35_)
                nw208_[int(10)] = (d_1072_v35_) + (d_1072_v35_)
                nw208_[int(11)] = _dafny.SeqWithoutIsStrInference([d_1071_v34_])
                nw208_[int(12)] = _dafny.SeqWithoutIsStrInference([d_1071_v34_])
                nw208_[int(13)] = d_1072_v35_
                nw208_[int(14)] = d_1072_v35_
                nw208_[int(15)] = d_1072_v35_
                nw208_[int(16)] = d_1072_v35_
                nw208_[int(17)] = d_1072_v35_
                nw208_[int(18)] = (d_1072_v35_) + (d_1072_v35_)
                nw208_[int(19)] = d_1072_v35_
                nw208_[int(20)] = d_1072_v35_
                nw208_[int(21)] = (d_1072_v35_) + (d_1072_v35_)
                nw208_[int(22)] = ((d_1073_v36_).cf44) + (d_1072_v35_)
                d_1074_v37_ = nw208_
                d_1074_v37_ = d_1074_v37_
                d_1040_v14_ = not(False)
                index180_ = default__.safeIndex(771, ((self).f10).length(0))
                ((self).f10)[index180_] = (self).f6
                index181_ = default__.safeIndex(771, ((self).f10).length(0))
                ((self).f10)[index181_] = (self).f9
            (globalState).f1 = (0) - (p0)
            d_1075_v38_: D6
            d_1075_v38_ = D6_DC17((self).f5, (self).f5)
            d_1075_v38_ = default__.fm31(globalState)
        if (self).f7:
            d_1076_v39_: _dafny.Array
            nw209_ = _dafny.Array(_dafny.Seq({}), 3)
            d_1076_v39_ = nw209_
            d_1077_v40_: D1
            d_1077_v40_ = D1_DC3(d_1076_v39_)
            pat_let_tv17_ = d_1076_v39_
            d_1078_v41_: _dafny.Array
            nw210_ = _dafny.Array(None, 21)
            nw210_[int(0)] = d_1077_v40_
            nw210_[int(1)] = d_1077_v40_
            nw210_[int(2)] = D1_DC3(d_1076_v39_)
            nw210_[int(3)] = d_1077_v40_
            nw210_[int(4)] = D1_DC3(d_1076_v39_)
            nw210_[int(5)] = d_1077_v40_
            nw210_[int(6)] = d_1077_v40_
            nw210_[int(7)] = d_1077_v40_
            nw210_[int(8)] = d_1077_v40_
            nw210_[int(9)] = d_1077_v40_
            nw210_[int(10)] = d_1077_v40_
            nw210_[int(11)] = d_1077_v40_
            nw210_[int(12)] = D1_DC3(d_1076_v39_)
            def iife81_(_pat_let30_0):
                def iife82_(d_1079_dt__update__tmp_h0_):
                    def iife83_(_pat_let31_0):
                        def iife84_(d_1080_dt__update_hcf7_h0_):
                            return D1_DC3(d_1080_dt__update_hcf7_h0_)
                        return iife84_(_pat_let31_0)
                    return iife83_(pat_let_tv17_)
                return iife82_(_pat_let30_0)
            nw210_[int(13)] = iife81_(D1_DC3(d_1076_v39_))
            nw210_[int(14)] = d_1077_v40_
            nw210_[int(15)] = d_1077_v40_
            nw210_[int(16)] = d_1077_v40_
            nw210_[int(17)] = d_1077_v40_
            nw210_[int(18)] = d_1077_v40_
            nw210_[int(19)] = d_1077_v40_
            nw210_[int(20)] = d_1077_v40_
            d_1078_v41_ = nw210_
            d_1081_v42_: C2
            nw211_ = C2()
            nw211_.ctor__((self).f10, d_1078_v41_)
            d_1081_v42_ = nw211_
            d_1082_v43_: _dafny.Map
            d_1082_v43_ = _dafny.Map({(self).f8: 612})
            r0 = d_1082_v43_
            source12_ = d_1077_v40_
            if source12_.is_DC4:
                d_1083___mcc_h5_ = source12_.cf8
                d_1084___mcc_h6_ = source12_.cf9
                d_1085___mcc_h7_ = source12_.cf10
                d_1086___mcc_h8_ = source12_.cf11
                d_1087___mcc_h9_ = source12_.cf12
                d_1088_cf12_ = d_1087___mcc_h9_
                d_1089_cf11_ = d_1086___mcc_h8_
                d_1090_cf10_ = d_1085___mcc_h7_
                d_1091_cf9_ = d_1084___mcc_h6_
                d_1092_cf8_ = d_1083___mcc_h5_
                d_1091_cf9_ = (d_1091_cf9_) + (len((d_1023_v0_) + (d_1024_v1_)))
                d_1093_v44_: _dafny.MultiSet
                d_1093_v44_ = _dafny.MultiSet([(self).f11, (self).f11])
                index182_ = default__.safeIndex(769, ((d_1081_v42_).f14).length(0))
                ((d_1081_v42_).f14)[index182_] = (d_1093_v44_).isdisjoint(d_1093_v44_)
                index183_ = default__.safeIndex(769, ((d_1081_v42_).f14).length(0))
                ((d_1081_v42_).f14)[index183_] = (self).f7
                d_1094_v45_: C4
                nw212_ = C4()
                nw212_.ctor__((self).f8, (self).f6, (self).f9, not(((d_1081_v42_).f14)[default__.safeIndex(769, ((d_1081_v42_).f14).length(0))]), d_1092_cf8_, (self).f5)
                d_1094_v45_ = nw212_
                d_1095_v46_: _dafny.Map
                d_1095_v46_ = _dafny.Map({d_1094_v45_: (self).f6})
                d_1096_v47_: _dafny.Map
                d_1096_v47_ = _dafny.Map({d_1090_cf10_: (self).f9})
                d_1097_v48_: _dafny.Map
                d_1097_v48_ = _dafny.Map({d_1095_v46_: d_1096_v47_})
                d_1098_v49_: _dafny.Map
                d_1098_v49_ = _dafny.Map({(self).f7: (self).f7})
                rhs165_ = d_1097_v48_
                rhs166_ = (self).fm1(default__.fm3(globalState), d_1098_v49_, not((self).f7), globalState)
                lhs125_ = globalState
                d_1097_v48_ = rhs165_
                lhs125_.f1 = rhs166_
                d_1099_v50_: _dafny.Seq
                d_1099_v50_ = _dafny.SeqWithoutIsStrInference([(self).f8])
                (globalState).f1 = default__.safeDivisionInt(default__.safeDivisionInt(d_1091_cf9_, d_1091_cf9_), len(d_1099_v50_))
            elif source12_.is_DC5:
                d_1100___mcc_h10_ = source12_.cf13
                d_1101_cf13_ = d_1100___mcc_h10_
                d_1102_v51_: _dafny.Seq
                d_1102_v51_ = _dafny.SeqWithoutIsStrInference([(self).f8])
                d_1103_v52_: _dafny.Seq
                d_1103_v52_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p0: len(d_1023_v0_)}), d_1082_v43_])
                d_1104_v53_: D6
                d_1104_v53_ = D6_DC17(((self).f5).set(default__.safeIndex(555, len((self).f5)), self.f4), (self).f5)
                d_1105_v54_: _dafny.Seq
                d_1105_v54_ = _dafny.SeqWithoutIsStrInference([(d_1102_v51_)[default__.safeIndex(p0, len(d_1102_v51_))], len((d_1103_v52_)[default__.safeIndex(len((d_1104_v53_).cf31), len(d_1103_v52_))])])
                d_1106_v55_: _dafny.Array
                nw213_ = _dafny.Array(None, 5)
                nw213_[int(0)] = (len((d_1082_v43_).set(p0, d_1101_cf13_))) in (d_1102_v51_)
                nw213_[int(1)] = (self).f7
                nw213_[int(2)] = (self).f9
                nw213_[int(3)] = (p0) < ((self).f8)
                nw213_[int(4)] = ((self).f7 if (self).f9 else (self).f9)
                d_1106_v55_ = nw213_
                d_1106_v55_ = (d_1081_v42_).f14
                d_1107_v56_: bool
                d_1107_v56_ = False
                d_1107_v56_ = (self).f7
                d_1108_v57_: _dafny.Array
                nw214_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_1108_v57_ = nw214_
                d_1109_v58_: D2
                d_1109_v58_ = D2_DC8((self).f8, -507)
                d_1110_v59_: D0
                d_1110_v59_ = D0_DC0()
                d_1111_v60_: _dafny.Map
                d_1111_v60_ = _dafny.Map({d_1110_v59_: d_1107_v56_})
                d_1112_v61_: D0
                d_1112_v61_ = D0_DC1((self).f9, p0, (self).f7, len(d_1111_v60_))
                d_1113_v62_: _dafny.Array
                nw215_ = _dafny.Array(None, 26)
                nw215_[int(0)] = d_1109_v58_
                nw215_[int(1)] = d_1109_v58_
                nw215_[int(2)] = d_1109_v58_
                nw215_[int(3)] = d_1109_v58_
                nw215_[int(4)] = D2_DC8((self).f8, (d_1112_v61_).cf3)
                nw215_[int(5)] = d_1109_v58_
                nw215_[int(6)] = d_1109_v58_
                nw215_[int(7)] = d_1109_v58_
                nw215_[int(8)] = d_1109_v58_
                nw215_[int(9)] = d_1109_v58_
                nw215_[int(10)] = d_1109_v58_
                nw215_[int(11)] = d_1109_v58_
                nw215_[int(12)] = d_1109_v58_
                nw215_[int(13)] = d_1109_v58_
                nw215_[int(14)] = d_1109_v58_
                nw215_[int(15)] = d_1109_v58_
                nw215_[int(16)] = d_1109_v58_
                nw215_[int(17)] = d_1109_v58_
                nw215_[int(18)] = d_1109_v58_
                nw215_[int(19)] = d_1109_v58_
                nw215_[int(20)] = d_1109_v58_
                nw215_[int(21)] = D2_DC8(p0, (self).f8)
                nw215_[int(22)] = d_1109_v58_
                nw215_[int(23)] = d_1109_v58_
                nw215_[int(24)] = d_1109_v58_
                nw215_[int(25)] = d_1109_v58_
                d_1113_v62_ = nw215_
                d_1114_v63_: D10
                d_1114_v63_ = D10_DC26(d_1113_v62_)
                d_1115_v64_: _dafny.Array
                nw216_ = _dafny.Array(None, 8)
                nw216_[int(0)] = d_1113_v62_
                nw216_[int(1)] = (d_1114_v63_).cf48
                nw216_[int(2)] = d_1113_v62_
                nw216_[int(3)] = d_1113_v62_
                nw216_[int(4)] = d_1113_v62_
                nw216_[int(5)] = d_1113_v62_
                nw216_[int(6)] = d_1113_v62_
                nw216_[int(7)] = d_1113_v62_
                d_1115_v64_ = nw216_
                index184_ = default__.safeIndex(347, (d_1108_v57_).length(0))
                (d_1108_v57_)[index184_] = d_1115_v64_
                index185_ = default__.safeIndex(347, (d_1108_v57_).length(0))
                (d_1108_v57_)[index185_] = d_1115_v64_
                index186_ = default__.safeIndex(54, ((self).f11).length(0))
                ((self).f11)[index186_] = len((self).f5)
                d_1116_v65_: _dafny.Map
                d_1116_v65_ = _dafny.Map({(self).f6: not(d_1107_v56_)})
                index187_ = default__.safeIndex(54, ((self).f11).length(0))
                rhs167_ = self.f4
                rhs168_ = 223
                rhs169_ = default__.safeDivisionInt((self).f8, default__.safeModuloInt((self).f8, (self).f8))
                rhs170_ = self.f4
                rhs171_ = default__.fm7(_dafny.CodePoint('v'), len(d_1116_v65_), (_dafny.Set({len(default__.fm24(d_1101_cf13_, (self).f9, (self).f7, p0, globalState))})).isdisjoint(default__.fm30((self).f6, globalState)), globalState)
                lhs126_ = self
                lhs127_ = (self).f11
                lhs128_ = default__.safeIndex(54, ((self).f11).length(0))
                lhs129_ = self
                lhs126_.f4 = rhs167_
                lhs127_[lhs128_] = rhs168_
                d_1101_cf13_ = rhs169_
                lhs129_.f4 = rhs170_
                d_1107_v56_ = rhs171_
            elif True:
                d_1117___mcc_h11_ = source12_.cf7
                d_1118_cf7_ = d_1117___mcc_h11_
                (globalState).f1 = p0
                (globalState).f1 = (0) - ((self).f8)
                d_1119_v66_: _dafny.Array
                def lambda40_(d_1120_p0_):
                    def lambda41_(d_1121_i4_):
                        return D2_DC8(d_1120_p0_, 730)

                    return lambda41_

                init22_ = lambda40_(p0)
                nw217_ = _dafny.Array(None, 17)
                for i0_22_ in range(nw217_.length(0)):
                    nw217_[i0_22_] = init22_(i0_22_)
                d_1119_v66_ = nw217_
                d_1122_v67_: D10
                d_1122_v67_ = D10_DC26(d_1119_v66_)
                pat_let_tv18_ = d_1119_v66_
                def iife85_(_pat_let32_0):
                    def iife86_(d_1123_dt__update__tmp_h1_):
                        def iife87_(_pat_let33_0):
                            def iife88_(d_1124_dt__update_hcf48_h0_):
                                return D10_DC26(d_1124_dt__update_hcf48_h0_)
                            return iife88_(_pat_let33_0)
                        return iife87_(pat_let_tv18_)
                    return iife86_(_pat_let32_0)
                d_1122_v67_ = iife85_(d_1122_v67_)
                d_1125_v68_: bool
                d_1125_v68_ = True
                d_1125_v68_ = True
            (globalState).f1 = p0
            d_1126_v69_: bool
            d_1126_v69_ = True
            d_1126_v69_ = False
        elif True:
            d_1127_v70_: bool
            d_1127_v70_ = True
            d_1128_v71_: _dafny.Set
            d_1128_v71_ = _dafny.Set({len((self).f5), (self).f8})
            d_1127_v70_ = ((d_1128_v71_ if (self).f6 else d_1128_v71_)).issubset(d_1128_v71_)
            d_1127_v70_ = not(False)
            d_1129_v72_: _dafny.Array
            nw218_ = _dafny.Array(int(0), 6)
            d_1129_v72_ = nw218_
            d_1129_v72_ = d_1129_v72_
            index188_ = default__.safeIndex(136, ((self).f10).length(0))
            ((self).f10)[index188_] = (self).f7
            index189_ = default__.safeIndex(136, ((self).f10).length(0))
            ((self).f10)[index189_] = not(d_1127_v70_)
            d_1130_v73_: _dafny.Map
            d_1130_v73_ = _dafny.Map({(self).f8: (self).f8})
            r0 = ((d_1130_v73_) | (d_1130_v73_)) | (d_1130_v73_)
        (globalState).f1 = (0) - ((self).f8)
        def lambda42_(source13_):
            if source13_.is_DC7:
                d_1131___mcc_h12_ = source13_.cf15
                d_1132_cf15_ = d_1131___mcc_h12_
                return (self).f6
            elif source13_.is_DC8:
                d_1133___mcc_h13_ = source13_.cf16
                d_1134___mcc_h14_ = source13_.cf17
                d_1135_cf17_ = d_1134___mcc_h14_
                d_1136_cf16_ = d_1133___mcc_h13_
                return (self).f6
            elif source13_.is_DC6:
                d_1137___mcc_h15_ = source13_.cf14
                d_1138_cf14_ = d_1137___mcc_h15_
                return (self).f9
            elif True:
                d_1139___mcc_h16_ = source13_.cf18
                d_1140_cf18_ = d_1139___mcc_h16_
                return (_dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f7})])) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f6, (self).f9, (self).f9, (self).f7, (self).f6}), _dafny.Set({(self).f6, (self).f7, (self).f7, (self).f9})]))

        def iife89_(_pat_let34_0):
            def iife90_(d_1141_dt__update__tmp_h2_):
                def iife91_(_pat_let35_0):
                    def iife92_(d_1142_dt__update_hcf15_h0_):
                        return D2_DC7(d_1142_dt__update_hcf15_h0_)
                    return iife92_(_pat_let35_0)
                return iife91_((self).f9)
            return iife90_(_pat_let34_0)
        if lambda42_(iife89_(D2_DC7((self).f6))):
            (globalState).f1 = p0
            d_1143_v74_: _dafny.Set
            d_1143_v74_ = _dafny.Set({(self).f9, (self).f6})
            if ((self).f8) == (len(d_1143_v74_)):
                d_1144_v75_: bool
                d_1144_v75_ = False
                d_1145_v76_: _dafny.Map
                d_1145_v76_ = _dafny.Map({p0: (self).f9})
                rhs172_ = default__.safeModuloInt(-556, p0)
                rhs173_ = (d_1144_v75_) or ((self).f7)
                rhs174_ = d_1144_v75_
                rhs175_ = d_1145_v76_
                lhs130_ = globalState
                lhs130_.f1 = rhs172_
                d_1144_v75_ = rhs173_
                d_1144_v75_ = rhs174_
                d_1145_v76_ = rhs175_
                (globalState).f1 = (self).f8
                d_1146_v77_: _dafny.Seq
                d_1146_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rc"))
                d_1146_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdhcdkvjy"))
                d_1147_v78_: C3
                nw219_ = C3()
                nw219_.ctor__(False, True, 317, True, not(((self).f8) >= ((0) - (p0))), (self).f9, self.f4, default__.fm14((0) - ((self).f8), len(d_1146_v77_), (self).f7, (self).f6, globalState))
                d_1147_v78_ = nw219_
                d_1148_v79_: _dafny.Seq
                d_1148_v79_ = _dafny.SeqWithoutIsStrInference([(self).f8, p0])
                d_1149_v80_: _dafny.MultiSet
                d_1149_v80_ = _dafny.MultiSet([d_1148_v79_, d_1148_v79_])
                d_1150_v81_: D11
                d_1150_v81_ = D11_DC29(d_1149_v80_)
                d_1151_v82_: C4
                nw220_ = C4()
                nw220_.ctor__(738, (d_1148_v79_) in ((d_1150_v81_).cf55), (self).f6, (self).f9, self.f4, d_1146_v77_)
                d_1151_v82_ = nw220_
            elif True:
                d_1152_v83_: _dafny.Map
                d_1152_v83_ = _dafny.Map({p0: p0})
                d_1152_v83_ = (d_1152_v83_).set(p0, p0)
                d_1153_v84_: _dafny.Array
                nw221_ = _dafny.Array(_dafny.Set({}), 16)
                d_1153_v84_ = nw221_
                index190_ = default__.safeIndex(915, (d_1153_v84_).length(0))
                (d_1153_v84_)[index190_] = d_1143_v74_
                d_1154_v85_: bool
                d_1154_v85_ = True
                index191_ = default__.safeIndex(915, (d_1153_v84_).length(0))
                rhs176_ = d_1143_v74_
                rhs177_ = ((p0) * ((self).f8)) + ((self).f8)
                rhs178_ = not((d_1154_v85_ if (self).f6 else not((self).f7)))
                rhs179_ = p0
                lhs131_ = d_1153_v84_
                lhs132_ = default__.safeIndex(915, (d_1153_v84_).length(0))
                lhs133_ = globalState
                lhs134_ = globalState
                lhs131_[lhs132_] = rhs176_
                lhs133_.f1 = rhs177_
                d_1154_v85_ = rhs178_
                lhs134_.f1 = rhs179_
                d_1155_v86_: _dafny.Seq
                d_1155_v86_ = _dafny.SeqWithoutIsStrInference([d_1023_v0_])
                d_1156_v87_: _dafny.Array
                nw222_ = _dafny.Array(None, 11)
                nw222_[int(0)] = d_1024_v1_
                nw222_[int(1)] = d_1024_v1_
                nw222_[int(2)] = (d_1155_v86_)[default__.safeIndex(p0, len(d_1155_v86_))]
                nw222_[int(3)] = _dafny.SeqWithoutIsStrInference([(self).f9])
                nw222_[int(4)] = d_1023_v0_
                nw222_[int(5)] = d_1024_v1_
                nw222_[int(6)] = d_1023_v0_
                nw222_[int(7)] = d_1023_v0_
                nw222_[int(8)] = d_1024_v1_
                nw222_[int(9)] = d_1024_v1_
                nw222_[int(10)] = d_1024_v1_
                d_1156_v87_ = nw222_
                d_1157_v88_: D1
                d_1157_v88_ = D1_DC3(d_1156_v87_)
                d_1158_v89_: _dafny.Array
                nw223_ = _dafny.Array(None, 19)
                nw223_[int(0)] = D1_DC3(d_1156_v87_)
                nw223_[int(1)] = d_1157_v88_
                nw223_[int(2)] = d_1157_v88_
                nw223_[int(3)] = d_1157_v88_
                nw223_[int(4)] = d_1157_v88_
                nw223_[int(5)] = D1_DC3(d_1156_v87_)
                nw223_[int(6)] = d_1157_v88_
                nw223_[int(7)] = d_1157_v88_
                nw223_[int(8)] = D1_DC3(d_1156_v87_)
                nw223_[int(9)] = d_1157_v88_
                nw223_[int(10)] = d_1157_v88_
                nw223_[int(11)] = d_1157_v88_
                nw223_[int(12)] = d_1157_v88_
                nw223_[int(13)] = D1_DC3(d_1156_v87_)
                nw223_[int(14)] = d_1157_v88_
                nw223_[int(15)] = d_1157_v88_
                nw223_[int(16)] = d_1157_v88_
                nw223_[int(17)] = D1_DC3(d_1156_v87_)
                nw223_[int(18)] = d_1157_v88_
                d_1158_v89_ = nw223_
                d_1159_v90_: C2
                nw224_ = C2()
                nw224_.ctor__((self).f10, d_1158_v89_)
                d_1159_v90_ = nw224_
                d_1160_v91_: C0
                nw225_ = C0()
                nw225_.ctor__(((self).f7 if d_1154_v85_ else (self).f7), (self).f6, self.f4, ((self).f5) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mum"))))
                d_1160_v91_ = nw225_
                d_1161_v92_: _dafny.Seq
                d_1161_v92_ = _dafny.SeqWithoutIsStrInference([p0, (0) - (len(_dafny.Set({-577}))), 258])
                d_1162_v93_: _dafny.Seq
                d_1162_v93_ = _dafny.SeqWithoutIsStrInference([d_1161_v92_, _dafny.SeqWithoutIsStrInference([p0, len((self).f5)]), d_1161_v92_, d_1161_v92_])
                d_1162_v93_ = d_1162_v93_
            d_1163_v94_: bool
            d_1163_v94_ = False
            d_1163_v94_ = (p0) == (p0)
            index192_ = default__.safeIndex(782, ((self).f10).length(0))
            ((self).f10)[index192_] = True
            index193_ = default__.safeIndex(782, ((self).f10).length(0))
            ((self).f10)[index193_] = d_1163_v94_
            d_1164_v95_: _dafny.Array
            def lambda43_(d_1165_i5_):
                return (self).f6

            init23_ = lambda43_
            nw226_ = _dafny.Array(None, 16)
            for i0_23_ in range(nw226_.length(0)):
                nw226_[i0_23_] = init23_(i0_23_)
            d_1164_v95_ = nw226_
            d_1166_v96_: _dafny.Array
            nw227_ = _dafny.Array(D1.default()(), 20)
            d_1166_v96_ = nw227_
            d_1167_v97_: C2
            nw228_ = C2()
            nw228_.ctor__(d_1164_v95_, d_1166_v96_)
            d_1167_v97_ = nw228_
        elif True:
            d_1168_v98_: _dafny.Seq
            d_1168_v98_ = _dafny.SeqWithoutIsStrInference([d_1024_v1_])
            d_1169_v99_: _dafny.Array
            nw229_ = _dafny.Array(None, 23)
            nw229_[int(0)] = d_1023_v0_
            nw229_[int(1)] = d_1024_v1_
            nw229_[int(2)] = d_1024_v1_
            nw229_[int(3)] = d_1024_v1_
            nw229_[int(4)] = d_1024_v1_
            nw229_[int(5)] = d_1024_v1_
            nw229_[int(6)] = d_1023_v0_
            nw229_[int(7)] = d_1024_v1_
            nw229_[int(8)] = _dafny.SeqWithoutIsStrInference([(self).f7, (self).f9])
            nw229_[int(9)] = d_1023_v0_
            nw229_[int(10)] = d_1024_v1_
            nw229_[int(11)] = d_1023_v0_
            nw229_[int(12)] = d_1023_v0_
            nw229_[int(13)] = d_1024_v1_
            nw229_[int(14)] = d_1024_v1_
            nw229_[int(15)] = d_1023_v0_
            nw229_[int(16)] = d_1024_v1_
            nw229_[int(17)] = d_1023_v0_
            nw229_[int(18)] = _dafny.SeqWithoutIsStrInference([(self).f9])
            nw229_[int(19)] = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f7])
            nw229_[int(20)] = d_1023_v0_
            nw229_[int(21)] = d_1023_v0_
            nw229_[int(22)] = (d_1168_v98_)[default__.safeIndex((self).f8, len(d_1168_v98_))]
            d_1169_v99_ = nw229_
            d_1170_v100_: _dafny.Map
            d_1170_v100_ = _dafny.Map({(_dafny.MultiSet((d_1024_v1_ if False else d_1023_v0_))).cardinality: D1_DC3(d_1169_v99_)})
            d_1171_v101_: D1
            d_1171_v101_ = D1_DC3(d_1169_v99_)
            d_1170_v100_ = (d_1170_v100_).set((p0 if (self).f7 else p0), d_1171_v101_)
            d_1172_v102_: bool
            d_1172_v102_ = False
            d_1172_v102_ = (self).f6
            d_1173_v103_: _dafny.Set
            d_1173_v103_ = _dafny.Set({(self).f9, False, d_1172_v102_})
            d_1174_v104_: D3
            d_1174_v104_ = D3_DC11(not((self).f6), (self).f8, d_1173_v103_)
            d_1175_v105_: _dafny.Seq
            d_1175_v105_ = _dafny.SeqWithoutIsStrInference([len(default__.fm28(len(_dafny.SeqWithoutIsStrInference([self.f4 for d_1176_i6_ in range(default__.abs(66))])), (self).f6, d_1024_v1_, d_1174_v104_, globalState))])
            d_1177_v106_: _dafny.Set
            d_1177_v106_ = _dafny.Set({(self).f8, (d_1175_v105_)[default__.safeIndex(p0, len(d_1175_v105_))]})
            d_1178_v107_: _dafny.Map
            d_1178_v107_ = _dafny.Map({not((self).f6): (d_1175_v105_).set(default__.safeIndex((d_1175_v105_)[default__.safeIndex((0) - (len(d_1177_v106_)), len(d_1175_v105_))], len(d_1175_v105_)), (self).f8)})
            d_1178_v107_ = (d_1178_v107_) | (d_1178_v107_)
            d_1172_v102_ = ((self).f6) in (d_1173_v103_)
            d_1179_v108_: _dafny.Map
            d_1179_v108_ = _dafny.Map({(self).f9: (self).f6})
            d_1180_v109_: D1
            d_1180_v109_ = D1_DC5(p0)
            d_1181_v110_: _dafny.MultiSet
            d_1181_v110_ = _dafny.MultiSet([p0, p0, len(_dafny.Map({_dafny.Map({p0: (self).fm1((self).f7, d_1179_v108_, default__.fm27((self).f8, d_1172_v102_, (self).f6, globalState), globalState)}): (self).f5})), (len(_dafny.SeqWithoutIsStrInference([d_1180_v109_, D1_DC5(len((self).f5)), d_1180_v109_, D1_DC5(528), d_1180_v109_]))) * (896)])
            (globalState).f1 = ((d_1181_v110_)[p0] if (p0) in (d_1181_v110_) else p0)
        d_1182_v111_: D1
        d_1182_v111_ = D1_DC5(-740)
        d_1183_v112_: _dafny.Seq
        d_1183_v112_ = _dafny.SeqWithoutIsStrInference([d_1182_v111_, d_1182_v111_])
        d_1184_v113_: _dafny.Seq
        d_1184_v113_ = d_1183_v112_
        def lambda44_(source14_):
            d_1185___mcc_h17_ = source14_
            d_1186_cf24_ = d_1185___mcc_h17_
            return (self).f7

        if lambda44_(d_1184_v113_):
            (globalState).f1 = p0
            d_1187_v114_: _dafny.Map
            d_1187_v114_ = _dafny.Map({self.f4: (self).f6})
            d_1188_v115_: _dafny.Map
            d_1188_v115_ = _dafny.Map({(self).f7: (self).f7})
            (self).m4(p0, (_dafny.Map({(self).f8: (self).f9})).set((self).f8, ((d_1187_v114_)[self.f4] if (self.f4) in (d_1187_v114_) else ((d_1188_v115_)[(self).f7] if ((self).f7) in (d_1188_v115_) else (self).f7))), globalState)
            (globalState).f1 = len(d_1023_v0_)
            d_1189_v116_: C0
            nw230_ = C0()
            nw230_.ctor__((d_1025_v2_).isdisjoint(d_1025_v2_), (self).f9, _dafny.CodePoint('q'), ((self).f5) + ((self).f5))
            d_1189_v116_ = nw230_
            d_1190_v118_: _dafny.Seq
            d_1190_v118_ = _dafny.SeqWithoutIsStrInference([default__.fm8(p0, (self).f9, (self).f6, globalState)])
            d_1191_v119_: _dafny.Map
            d_1191_v119_ = _dafny.Map({(self).f9: p0})
            d_1192_v120_: _dafny.Map
            d_1192_v120_ = _dafny.Map({len(d_1190_v118_): d_1191_v119_})
            d_1193_v121_: _dafny.Seq
            d_1193_v121_ = _dafny.SeqWithoutIsStrInference([((d_1192_v120_)[(self).f8] if ((self).f8) in (d_1192_v120_) else d_1191_v119_)])
            d_1194_v122_: _dafny.Seq
            def iife93_():
                coll21_ = _dafny.Map()
                compr_21_: _dafny.Map
                for compr_21_ in (d_1193_v121_).Elements:
                    d_1195_v117_: _dafny.Map = compr_21_
                    if (d_1195_v117_) in (d_1193_v121_):
                        coll21_[d_1195_v117_] = (self).f6
                return _dafny.Map(coll21_)
            d_1194_v122_ = _dafny.SeqWithoutIsStrInference([(self).f8, len(iife93_()
            )])
            (globalState).f1 = ((d_1194_v122_)[default__.safeIndex((self).f8, len(d_1194_v122_))]) - ((p0) - (p0))
        elif True:
            d_1196_v123_: _dafny.MultiSet
            d_1196_v123_ = _dafny.MultiSet([(self).f5, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbxa"))])
            d_1197_v124_: _dafny.Set
            d_1197_v124_ = _dafny.Set({(self).f7, default__.fm27(p0, (self).f9, (self).f6, globalState)})
            d_1198_v125_: D3
            d_1198_v125_ = D3_DC11((self).f9, (d_1196_v123_).cardinality, d_1197_v124_)
            d_1199_v126_: _dafny.Map
            d_1199_v126_ = _dafny.Map({(self).f8: default__.fm28(p0, False, _dafny.SeqWithoutIsStrInference([(self).f7]), d_1198_v125_, globalState)})
            d_1199_v126_ = (d_1199_v126_ if (self).f9 else d_1199_v126_)
            d_1200_v127_: bool
            out21_: bool
            out21_ = (self).m3(p0, (self).f8, p0, globalState)
            d_1200_v127_ = out21_
            d_1201_v128_: _dafny.Map
            d_1201_v128_ = _dafny.Map({(self).f8: len(_dafny.SeqWithoutIsStrInference([462 for d_1202_i7_ in range(default__.abs(424))]))})
            d_1203_v129_: _dafny.Map
            d_1203_v129_ = _dafny.Map({(self).f9: (self).f9})
            (globalState).f1 = default__.safeDivisionInt(((d_1201_v128_)[p0] if (p0) in (d_1201_v128_) else len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfpkek")) for d_1204_i8_ in range(default__.abs(-830))]))), (self).fm1(default__.fm3(globalState), d_1203_v129_, (self).f7, globalState))
            d_1205_v130_: C0
            nw231_ = C0()
            nw231_.ctor__(d_1200_v127_, (self).f9, self.f4, (self).f5)
            d_1205_v130_ = nw231_
            index194_ = default__.safeIndex(21, ((self).f10).length(0))
            ((self).f10)[index194_] = (self).f9
            index195_ = default__.safeIndex(739, ((self).f10).length(0))
            ((self).f10)[index195_] = (d_1025_v2_).ispropersubset(d_1025_v2_)
            index196_ = default__.safeIndex(21, ((self).f10).length(0))
            index197_ = default__.safeIndex(739, ((self).f10).length(0))
            rhs180_ = not((not ((self).f9) or ((self).f7)) or ((self).f6))
            rhs181_ = ((self).f8) > ((self).f8)
            rhs182_ = (self).f9
            rhs183_ = (self).f7
            rhs184_ = (d_1023_v0_) + ((d_1023_v0_) + (d_1024_v1_))
            lhs135_ = (self).f10
            lhs136_ = default__.safeIndex(21, ((self).f10).length(0))
            lhs137_ = (self).f10
            lhs138_ = default__.safeIndex(739, ((self).f10).length(0))
            lhs135_[lhs136_] = rhs180_
            lhs137_[lhs138_] = rhs181_
            d_1200_v127_ = rhs182_
            d_1200_v127_ = rhs183_
            d_1024_v1_ = rhs184_
        if (self).f9:
            d_1206_v131_: _dafny.Set
            d_1206_v131_ = _dafny.Set({(d_1023_v0_) <= (_dafny.SeqWithoutIsStrInference([(self).f9, (self).f7])), not ((self).f9) or ((self).f9), True})
            d_1206_v131_ = d_1206_v131_
            d_1207_v132_: _dafny.Array
            nw232_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_1207_v132_ = nw232_
            index198_ = default__.safeIndex(348, (d_1207_v132_).length(0))
            (d_1207_v132_)[index198_] = (self).f10
            index199_ = default__.safeIndex(348, (d_1207_v132_).length(0))
            (d_1207_v132_)[index199_] = (self).f10
            d_1208_v133_: _dafny.Map
            d_1208_v133_ = _dafny.Map({(self).f9: (self).f9})
            d_1209_v134_: D3
            d_1209_v134_ = D3_DC11((self).f7, len((self).f5), d_1206_v131_)
            pat_let_tv19_ = d_1206_v131_
            d_1210_v135_: C3
            nw233_ = C3()
            def iife94_(_pat_let36_0):
                def iife95_(d_1211_dt__update__tmp_h3_):
                    def iife96_(_pat_let37_0):
                        def iife97_(d_1212_dt__update_hcf22_h0_):
                            def iife98_(_pat_let38_0):
                                def iife99_(d_1213_dt__update_hcf20_h0_):
                                    return D3_DC11(d_1213_dt__update_hcf20_h0_, (d_1211_dt__update__tmp_h3_).cf21, d_1212_dt__update_hcf22_h0_)
                                return iife99_(_pat_let38_0)
                            return iife98_((self).f7)
                        return iife97_(_pat_let37_0)
                    return iife96_(pat_let_tv19_)
                return iife95_(_pat_let36_0)
            nw233_.ctor__((428) not in (_dafny.SeqWithoutIsStrInference([p0])), ((self).f6) not in (d_1206_v131_), (self).fm1((self).f6, d_1208_v133_, (self).f6, globalState), (False) and (default__.fm7(default__.fm13(globalState), default__.fm8((self).f8, (self).f6, (d_1024_v1_)[default__.safeIndex((self).f8, len(d_1024_v1_))], globalState), True, globalState)), (self).f6, (self).f9, self.f4, default__.fm28(p0, (self).f6, d_1023_v0_, iife94_(d_1209_v134_), globalState))
            d_1210_v135_ = nw233_
            d_1210_v135_ = d_1210_v135_
            d_1214_v136_: _dafny.Seq
            d_1214_v136_ = _dafny.SeqWithoutIsStrInference([p0, (0) - (default__.fm8(p0, (d_1210_v135_).f13, not((self).f9), globalState)), (self).f8])
            d_1214_v136_ = d_1214_v136_
            d_1215_v137_: D1
            d_1215_v137_ = D1_DC4(self.f4, (0) - (p0), self.f4, p0, d_1210_v135_.f12)
            d_1216_v138_: _dafny.MultiSet
            d_1216_v138_ = _dafny.MultiSet([d_1215_v137_])
            (d_1210_v135_).f12 = (not((d_1206_v131_).isdisjoint(d_1206_v131_))) == ((d_1216_v138_).ispropersubset(d_1216_v138_))
        elif True:
            d_1217_v139_: bool
            d_1217_v139_ = True
            d_1218_v140_: _dafny.Array
            nw234_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_1218_v140_ = nw234_
            index200_ = default__.safeIndex(856, (d_1218_v140_).length(0))
            (d_1218_v140_)[index200_] = (self).f10
            index201_ = default__.safeIndex(461, ((self).f10).length(0))
            ((self).f10)[index201_] = (self.f4) in ((self).f5)
            index202_ = default__.safeIndex(856, (d_1218_v140_).length(0))
            index203_ = default__.safeIndex(461, ((self).f10).length(0))
            rhs185_ = not (default__.fm3(globalState)) or ((self).f6)
            rhs186_ = self.f4
            rhs187_ = (self).f10
            rhs188_ = (self).f8
            rhs189_ = (self).f9
            lhs139_ = self
            lhs140_ = d_1218_v140_
            lhs141_ = default__.safeIndex(856, (d_1218_v140_).length(0))
            lhs142_ = globalState
            lhs143_ = (self).f10
            lhs144_ = default__.safeIndex(461, ((self).f10).length(0))
            d_1217_v139_ = rhs185_
            lhs139_.f4 = rhs186_
            lhs140_[lhs141_] = rhs187_
            lhs142_.f1 = rhs188_
            lhs143_[lhs144_] = rhs189_
            (globalState).f1 = (self).f8
            d_1219_v141_: _dafny.Array
            def lambda45_(d_1220_i9_):
                return (d_1220_i9_) - (len(_dafny.Map({(self).f8: (self).f9})))

            init24_ = lambda45_
            nw235_ = _dafny.Array(None, 29)
            for i0_24_ in range(nw235_.length(0)):
                nw235_[i0_24_] = init24_(i0_24_)
            d_1219_v141_ = nw235_
            index204_ = default__.safeIndex(162, ((self).f11).length(0))
            ((self).f11)[index204_] = (0) - (len(_dafny.Map({(self).f11: (self).f6})))
            index205_ = default__.safeIndex(162, ((self).f11).length(0))
            rhs190_ = d_1219_v141_
            rhs191_ = d_1217_v139_
            rhs192_ = (self).f6
            rhs193_ = (self).f6
            rhs194_ = len(d_1023_v0_)
            lhs145_ = (self).f11
            lhs146_ = default__.safeIndex(162, ((self).f11).length(0))
            d_1219_v141_ = rhs190_
            d_1217_v139_ = rhs191_
            d_1217_v139_ = rhs192_
            d_1217_v139_ = rhs193_
            lhs145_[lhs146_] = rhs194_
            d_1221_v142_: _dafny.Seq
            d_1221_v142_ = _dafny.SeqWithoutIsStrInference([(self).f8])
            if (d_1221_v142_) <= ((_dafny.SeqWithoutIsStrInference([((self).f11)[default__.safeIndex(162, ((self).f11).length(0))], ((self).f11)[default__.safeIndex(162, ((self).f11).length(0))], p0]) if (self).f6 else d_1221_v142_)):
                d_1222_v143_: _dafny.Map
                d_1222_v143_ = _dafny.Map({(self).f8: (self).f9})
                (self).m4((0) - (len(default__.fm18(p0, not(((self).f10)[default__.safeIndex(461, ((self).f10).length(0))]), globalState))), d_1222_v143_, globalState)
                d_1223_v144_: _dafny.Array
                nw236_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_1223_v144_ = nw236_
                d_1224_v145_: _dafny.Array
                nw237_ = _dafny.Array(_dafny.CodePoint('D'), 17)
                d_1224_v145_ = nw237_
                index206_ = default__.safeIndex(20, (d_1223_v144_).length(0))
                (d_1223_v144_)[index206_] = d_1224_v145_
                d_1225_v146_: _dafny.Array
                nw238_ = _dafny.Array(D1.default()(), 22)
                d_1225_v146_ = nw238_
                d_1226_v147_: C2
                nw239_ = C2()
                nw239_.ctor__((d_1218_v140_)[default__.safeIndex(856, (d_1218_v140_).length(0))], d_1225_v146_)
                d_1226_v147_ = nw239_
                d_1227_v148_: _dafny.MultiSet
                d_1227_v148_ = _dafny.MultiSet([((self).f11)[default__.safeIndex(162, ((self).f11).length(0))]])
                index207_ = default__.safeIndex(20, (d_1223_v144_).length(0))
                index208_ = default__.safeIndex(162, ((self).f11).length(0))
                rhs195_ = d_1224_v145_
                rhs196_ = (0) - ((self).f8)
                rhs197_ = p0
                rhs198_ = (d_1226_v147_ if (p0) <= ((d_1227_v148_).cardinality) else d_1226_v147_)
                lhs147_ = d_1223_v144_
                lhs148_ = default__.safeIndex(20, (d_1223_v144_).length(0))
                lhs149_ = globalState
                lhs150_ = (self).f11
                lhs151_ = default__.safeIndex(162, ((self).f11).length(0))
                lhs147_[lhs148_] = rhs195_
                lhs149_.f1 = rhs196_
                lhs150_[lhs151_] = rhs197_
                d_1226_v147_ = rhs198_
                index209_ = default__.safeIndex(162, ((self).f11).length(0))
                ((self).f11)[index209_] = p0
                d_1024_v1_ = d_1024_v1_
                d_1228_v149_: D1
                d_1228_v149_ = D1_DC4(default__.fm13(globalState), p0, self.f4, (self).f8, (self).f9)
                d_1229_v150_: _dafny.Array
                nw240_ = _dafny.Array(None, 21)
                nw240_[int(0)] = d_1023_v0_
                nw240_[int(1)] = d_1023_v0_
                nw240_[int(2)] = d_1023_v0_
                nw240_[int(3)] = (d_1024_v1_).set(default__.safeIndex(p0, len(d_1024_v1_)), (self).f7)
                nw240_[int(4)] = d_1023_v0_
                nw240_[int(5)] = _dafny.SeqWithoutIsStrInference([((self).f10)[default__.safeIndex(461, ((self).f10).length(0))], (self).f6, (self).f7, d_1217_v139_])
                nw240_[int(6)] = _dafny.SeqWithoutIsStrInference([d_1217_v139_, (self).f7])
                nw240_[int(7)] = d_1024_v1_
                nw240_[int(8)] = d_1024_v1_
                nw240_[int(9)] = d_1023_v0_
                nw240_[int(10)] = _dafny.SeqWithoutIsStrInference([(self).f7])
                nw240_[int(11)] = _dafny.SeqWithoutIsStrInference([((self).f10)[default__.safeIndex(461, ((self).f10).length(0))]])
                nw240_[int(12)] = d_1024_v1_
                nw240_[int(13)] = d_1024_v1_
                nw240_[int(14)] = _dafny.SeqWithoutIsStrInference([(self).f9])
                nw240_[int(15)] = d_1023_v0_
                nw240_[int(16)] = d_1024_v1_
                nw240_[int(17)] = d_1023_v0_
                nw240_[int(18)] = d_1024_v1_
                nw240_[int(19)] = d_1023_v0_
                nw240_[int(20)] = d_1023_v0_
                d_1229_v150_ = nw240_
                d_1230_v151_: _dafny.Map
                d_1230_v151_ = _dafny.Map({d_1217_v139_: D1_DC3(d_1229_v150_)})
                d_1231_v152_: int
                d_1232_v153_: bool
                d_1233_v154_: int
                out22_: int
                out23_: bool
                out24_: int
                out22_, out23_, out24_ = (d_1226_v147_).m7(d_1228_v149_, default__.fm32((d_1023_v0_).set(default__.safeIndex(767, len(d_1023_v0_)), d_1217_v139_), True, globalState), d_1230_v151_, globalState)
                d_1231_v152_ = out22_
                d_1232_v153_ = out23_
                d_1233_v154_ = out24_
            elif True:
                (self).f4 = self.f4
                d_1217_v139_ = (self).f9
                d_1234_v155_: _dafny.Seq
                d_1234_v155_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cklypmb"))
                d_1234_v155_ = (self).f5
                d_1235_v156_: _dafny.Map
                d_1235_v156_ = _dafny.Map({(self).f8: d_1217_v139_})
                d_1236_v157_: _dafny.Map
                d_1236_v157_ = _dafny.Map({316: d_1235_v156_})
                d_1237_v158_: _dafny.Seq
                d_1237_v158_ = _dafny.SeqWithoutIsStrInference([(self).f5])
                d_1238_v159_: _dafny.Map
                d_1238_v159_ = _dafny.Map({(0) - (len(d_1236_v157_)): (d_1237_v158_)[default__.safeIndex((self).f8, len(d_1237_v158_))]})
                d_1238_v159_ = (d_1238_v159_).set(len(d_1023_v0_), d_1234_v155_)
                d_1239_v160_: D9
                d_1239_v160_ = D9_DC25(default__.fm27((self).f8, (self).f6, not(((self).f10)[default__.safeIndex(461, ((self).f10).length(0))]), globalState), (self).f7, (default__.fm33(globalState)).cardinality)
                index210_ = default__.safeIndex(461, ((self).f10).length(0))
                ((self).f10)[index210_] = ((0) - (((self).f11)[default__.safeIndex(162, ((self).f11).length(0))])) <= (((d_1239_v160_).cf47) * (p0))
            d_1240_v161_: _dafny.Map
            d_1240_v161_ = _dafny.Map({d_1217_v139_: ((self).f11)[default__.safeIndex(162, ((self).f11).length(0))]})
            d_1241_v162_: D5
            d_1241_v162_ = D5_DC15((d_1218_v140_)[default__.safeIndex(856, (d_1218_v140_).length(0))], False, d_1240_v161_, ((self).f11)[default__.safeIndex(162, ((self).f11).length(0))])
            d_1242_v163_: _dafny.Map
            d_1242_v163_ = _dafny.Map({d_1217_v139_: (d_1024_v1_)[default__.safeIndex(((self).f11)[default__.safeIndex(162, ((self).f11).length(0))], len(d_1024_v1_))]})
            d_1243_v164_: _dafny.Array
            nw241_ = _dafny.Array(None, 22)
            nw241_[int(0)] = d_1241_v162_
            nw241_[int(1)] = d_1241_v162_
            nw241_[int(2)] = d_1241_v162_
            nw241_[int(3)] = d_1241_v162_
            nw241_[int(4)] = D5_DC15((d_1218_v140_)[default__.safeIndex(856, (d_1218_v140_).length(0))], (self).f7, d_1240_v161_, ((self).f11)[default__.safeIndex(162, ((self).f11).length(0))])
            nw241_[int(5)] = D5_DC15((d_1218_v140_)[default__.safeIndex(856, (d_1218_v140_).length(0))], (self).f7, d_1240_v161_, (self).f8)
            nw241_[int(6)] = d_1241_v162_
            nw241_[int(7)] = d_1241_v162_
            nw241_[int(8)] = d_1241_v162_
            nw241_[int(9)] = D5_DC15((d_1218_v140_)[default__.safeIndex(856, (d_1218_v140_).length(0))], d_1217_v139_, _dafny.Map({(d_1024_v1_)[default__.safeIndex((self).fm1((self).f9, d_1242_v163_, ((self).f10)[default__.safeIndex(461, ((self).f10).length(0))], globalState), len(d_1024_v1_))]: p0}), ((self).f11)[default__.safeIndex(162, ((self).f11).length(0))])
            nw241_[int(10)] = D5_DC15((d_1218_v140_)[default__.safeIndex(856, (d_1218_v140_).length(0))], (self).f7, d_1240_v161_, 831)
            nw241_[int(11)] = d_1241_v162_
            nw241_[int(12)] = d_1241_v162_
            nw241_[int(13)] = d_1241_v162_
            nw241_[int(14)] = d_1241_v162_
            nw241_[int(15)] = D5_DC15((d_1218_v140_)[default__.safeIndex(856, (d_1218_v140_).length(0))], ((self).f10)[default__.safeIndex(461, ((self).f10).length(0))], d_1240_v161_, len((self).f5))
            nw241_[int(16)] = D5_DC15((d_1218_v140_)[default__.safeIndex(856, (d_1218_v140_).length(0))], default__.fm3(globalState), d_1240_v161_, 576)
            nw241_[int(17)] = d_1241_v162_
            nw241_[int(18)] = d_1241_v162_
            nw241_[int(19)] = d_1241_v162_
            nw241_[int(20)] = d_1241_v162_
            nw241_[int(21)] = d_1241_v162_
            d_1243_v164_ = nw241_
            index211_ = default__.safeIndex(457, (d_1243_v164_).length(0))
            (d_1243_v164_)[index211_] = d_1241_v162_
            d_1244_v165_: D3
            d_1244_v165_ = D3_DC11(((self).f10)[default__.safeIndex(461, ((self).f10).length(0))], ((self).f11)[default__.safeIndex(162, ((self).f11).length(0))], _dafny.Set({(self).f7}))
            pat_let_tv20_ = d_1240_v161_
            pat_let_tv21_ = d_1244_v165_
            pat_let_tv22_ = p0
            index212_ = default__.safeIndex(457, (d_1243_v164_).length(0))
            def iife101_(_pat_let40_0):
                def iife102_(d_1245_dt__update__tmp_h4_):
                    def iife103_(_pat_let41_0):
                        def iife104_(d_1246_dt__update_hcf28_h0_):
                            return D5_DC15((d_1245_dt__update__tmp_h4_).cf26, (d_1245_dt__update__tmp_h4_).cf27, d_1246_dt__update_hcf28_h0_, (d_1245_dt__update__tmp_h4_).cf29)
                        return iife104_(_pat_let41_0)
                    return iife103_(pat_let_tv20_)
                return iife102_(_pat_let40_0)
            def iife100_(_pat_let39_0):
                def iife105_(d_1247_dt__update__tmp_h5_):
                    def iife106_(_pat_let42_0):
                        def iife107_(d_1248_dt__update_hcf27_h0_):
                            def iife108_(_pat_let43_0):
                                def iife109_(d_1249_dt__update_hcf28_h1_):
                                    return D5_DC15((d_1247_dt__update__tmp_h5_).cf26, d_1248_dt__update_hcf27_h0_, d_1249_dt__update_hcf28_h1_, (d_1247_dt__update__tmp_h5_).cf29)
                                return iife109_(_pat_let43_0)
                            return iife108_((_dafny.Map({(pat_let_tv21_).cf20: ((self).f11)[default__.safeIndex(162, ((self).f11).length(0))]})).set(not(((self).f10)[default__.safeIndex(461, ((self).f10).length(0))]), pat_let_tv22_))
                        return iife107_(_pat_let42_0)
                    return iife106_(not((self).f6))
                return iife105_(_pat_let39_0)
            (d_1243_v164_)[index212_] = iife100_(iife101_(d_1241_v162_))
        d_1250_v166_: _dafny.MultiSet
        d_1250_v166_ = _dafny.MultiSet([537, p0])
        d_1251_v167_: _dafny.Map
        d_1251_v167_ = _dafny.Map({(self).f8: (d_1250_v166_).cardinality})
        r0 = (d_1251_v167_) | ((_dafny.Map({(self).f8: 66})) | (_dafny.Map({(_dafny.MultiSet([not((self).f7)])).cardinality: p0})))
        return r0

    def m3(self, p0, p1, p2, globalState):
        r0: bool = False
        d_1252_v0_: _dafny.MultiSet
        d_1252_v0_ = _dafny.MultiSet([(self).f5, (self).f5, (self).f5])
        d_1253_v1_: _dafny.Seq
        d_1253_v1_ = _dafny.SeqWithoutIsStrInference([d_1252_v0_, d_1252_v0_])
        d_1254_v2_: _dafny.Map
        d_1254_v2_ = _dafny.Map({(d_1252_v0_).ispropersubset((d_1253_v1_)[default__.safeIndex(p1, len(d_1253_v1_))]): self.f4})
        d_1254_v2_ = (d_1254_v2_).set(not ((self).f9) or ((self).f7), self.f4)
        hi5_ = p0
        for d_1255_i0_ in range(p2, hi5_):
            d_1256_v3_: _dafny.Set
            d_1256_v3_ = _dafny.Set({(self).f9})
            d_1257_v4_: _dafny.Seq
            d_1257_v4_ = _dafny.SeqWithoutIsStrInference([d_1256_v3_, d_1256_v3_])
            d_1258_v5_: _dafny.Seq
            d_1258_v5_ = _dafny.SeqWithoutIsStrInference([(self).f7])
            d_1259_v6_: _dafny.Seq
            d_1259_v6_ = _dafny.SeqWithoutIsStrInference([not((d_1256_v3_).ispropersubset((d_1257_v4_)[default__.safeIndex(len(d_1258_v5_), len(d_1257_v4_))]))])
            d_1258_v5_ = ((d_1259_v6_) + ((d_1259_v6_ if (self).f7 else _dafny.SeqWithoutIsStrInference([(self).f6])))).set(default__.safeIndex(637, len((d_1259_v6_) + ((d_1259_v6_ if (self).f7 else _dafny.SeqWithoutIsStrInference([(self).f6]))))), (self).f6)
            r0 = default__.fm7(self.f4, d_1255_i0_, (self).f9, globalState)
            d_1260_v7_: D1
            d_1260_v7_ = D1_DC4(self.f4, 501, self.f4, p2, (self).f6)
            d_1261_v8_: C0
            nw242_ = C0()
            nw242_.ctor__((d_1260_v7_).cf12, (d_1259_v6_)[default__.safeIndex(p0, len(d_1259_v6_))], self.f4, (self).f5)
            d_1261_v8_ = nw242_
            d_1262_v9_: _dafny.Seq
            d_1262_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jne"))
            index213_ = default__.safeIndex(596, ((self).f11).length(0))
            ((self).f11)[index213_] = p1
            d_1263_v10_: _dafny.MultiSet
            d_1263_v10_ = _dafny.MultiSet([(self).f9])
            index214_ = default__.safeIndex(596, ((self).f11).length(0))
            rhs199_ = ((self).f5) + ((self).f5)
            rhs200_ = True
            rhs201_ = len(_dafny.Map({(default__.fm4((self).f9, len(d_1262_v9_), (self).f9, p0, globalState)).ispropersubset(d_1263_v10_): (d_1258_v5_) + (_dafny.SeqWithoutIsStrInference([(self).f7, (self).f9]))}))
            rhs202_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxofl"))) + (_dafny.SeqWithoutIsStrInference([self.f4 for d_1264_i1_ in range(default__.abs(951))]))
            lhs152_ = (self).f11
            lhs153_ = default__.safeIndex(596, ((self).f11).length(0))
            d_1262_v9_ = rhs199_
            r0 = rhs200_
            lhs152_[lhs153_] = rhs201_
            d_1262_v9_ = rhs202_
        r0 = (_dafny.SeqWithoutIsStrInference([self.f4 for d_1265_i2_ in range(default__.abs(414))])) < ((self).f5)
        hi6_ = len(_dafny.SeqWithoutIsStrInference([(self).f8]))
        for d_1266_i3_ in range(p1, hi6_):
            d_1267_v11_: _dafny.Array
            nw243_ = _dafny.Array(int(0), 6)
            d_1267_v11_ = nw243_
            d_1267_v11_ = d_1267_v11_
            d_1268_v12_: _dafny.Set
            d_1268_v12_ = _dafny.Set({77})
            d_1269_v13_: _dafny.Map
            d_1269_v13_ = _dafny.Map({not((d_1268_v12_).ispropersubset(d_1268_v12_)): (self).f5})
            d_1269_v13_ = (d_1269_v13_).set((self).f7, (self).f5)
            d_1270_v14_: _dafny.Set
            d_1270_v14_ = _dafny.Set({self.f4})
            d_1271_v15_: _dafny.Seq
            d_1271_v15_ = _dafny.SeqWithoutIsStrInference([(d_1270_v14_) - (d_1270_v14_), d_1270_v14_, d_1270_v14_, (d_1270_v14_) | (d_1270_v14_), d_1270_v14_])
            (globalState).f1 = len((d_1271_v15_)[default__.safeIndex(d_1266_i3_, len(d_1271_v15_))])
            d_1272_v16_: _dafny.Seq
            d_1272_v16_ = _dafny.SeqWithoutIsStrInference([(self).f9, ((self).f5) < ((self).f5)])
            d_1272_v16_ = (d_1272_v16_).set(default__.safeIndex((d_1266_i3_) - ((self).f8), len(d_1272_v16_)), default__.fm27(p2, not((self).f7), (self).f9, globalState))
        d_1273_v17_: C1
        nw244_ = C1()
        nw244_.ctor__(p2, False, default__.fm27((self).f8, (self).f6, (self).f9, globalState), (self).f9, self.f4, (self).f5)
        d_1273_v17_ = nw244_
        d_1274_v18_: D10
        d_1274_v18_ = D10_DC28((self).f7, p2, (self).f6, (self).f9, d_1273_v17_)
        d_1275_v19_: _dafny.MultiSet
        d_1275_v19_ = _dafny.MultiSet([(d_1274_v18_).cf51])
        (globalState).f1 = (0) - (default__.safeDivisionInt((d_1275_v19_).cardinality, p1))
        d_1276_i4_: int
        d_1276_i4_ = 0
        with _dafny.label("5"):
            while (self).f7:
                with _dafny.c_label("5"):
                    if (d_1276_i4_) >= (100):
                        raise _dafny.Break("5")
                    d_1276_i4_ = (d_1276_i4_) + (1)
                    d_1277_v20_: _dafny.Array
                    nw245_ = _dafny.Array(_dafny.Seq({}), 9)
                    d_1277_v20_ = nw245_
                    d_1278_v21_: _dafny.Seq
                    d_1278_v21_ = _dafny.SeqWithoutIsStrInference([False])
                    d_1279_v22_: _dafny.Seq
                    d_1279_v22_ = _dafny.SeqWithoutIsStrInference([d_1278_v21_])
                    index215_ = default__.safeIndex(67, (d_1277_v20_).length(0))
                    (d_1277_v20_)[index215_] = d_1279_v22_
                    d_1280_v23_: _dafny.Set
                    d_1280_v23_ = _dafny.Set({(self).f9, (self).f6, (self).f7, (self).f9, (self).f9})
                    d_1281_v24_: D8
                    d_1281_v24_ = D8_DC22(d_1278_v21_)
                    index216_ = default__.safeIndex(67, (d_1277_v20_).length(0))
                    rhs203_ = default__.fm34(not((self).f9), d_1281_v24_, default__.safeDivisionInt((self).f8, default__.fm8(p2, (self).f7, (self).f6, globalState)), globalState)
                    rhs204_ = (len(d_1278_v21_)) + (p2)
                    rhs205_ = d_1280_v23_
                    rhs206_ = (self).f6
                    lhs154_ = d_1277_v20_
                    lhs155_ = default__.safeIndex(67, (d_1277_v20_).length(0))
                    lhs156_ = globalState
                    lhs154_[lhs155_] = rhs203_
                    lhs156_.f1 = rhs204_
                    d_1280_v23_ = rhs205_
                    r0 = rhs206_
                    d_1282_v25_: _dafny.Seq
                    d_1282_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etoajigf"))
                    d_1282_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
                    (self).f4 = self.f4
                    d_1283_v26_: _dafny.Seq
                    d_1283_v26_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f9])).set((self).f9, default__.abs(p0)), _dafny.MultiSet([(self).f6, (self).f7, (self).f6])])
                    d_1284_v27_: D3
                    d_1284_v27_ = D3_DC11((self).f9, p0, d_1280_v23_)
                    d_1285_v28_: _dafny.Array
                    def lambda46_(d_1286_i5_):
                        return self.f4

                    init25_ = lambda46_
                    nw246_ = _dafny.Array(None, 28)
                    for i0_25_ in range(nw246_.length(0)):
                        nw246_[i0_25_] = init25_(i0_25_)
                    d_1285_v28_ = nw246_
                    index217_ = default__.safeIndex(133, (d_1285_v28_).length(0))
                    (d_1285_v28_)[index217_] = ((self).f5)[default__.safeIndex(p0, len((self).f5))]
                    d_1287_v29_: _dafny.Array
                    nw247_ = _dafny.Array(D1.default()(), 6)
                    d_1287_v29_ = nw247_
                    d_1288_v30_: C2
                    nw248_ = C2()
                    nw248_.ctor__((self).f10, d_1287_v29_)
                    d_1288_v30_ = nw248_
                    d_1289_v31_: _dafny.Map
                    d_1289_v31_ = _dafny.Map({(self).f7: (self).f9})
                    d_1290_v32_: _dafny.MultiSet
                    d_1290_v32_ = _dafny.MultiSet([((d_1289_v31_)[(self).f7] if ((self).f7) in (d_1289_v31_) else True), (self).f6, (self).f6, not((self).f9), (self).f7])
                    d_1291_v33_: _dafny.Map
                    d_1291_v33_ = _dafny.Map({True: p2})
                    pat_let_tv23_ = d_1280_v23_
                    pat_let_tv24_ = globalState
                    index218_ = default__.safeIndex(133, (d_1285_v28_).length(0))
                    def iife111_(_pat_let45_0):
                        def iife112_(d_1292_dt__update__tmp_h0_):
                            def iife113_(_pat_let46_0):
                                def iife114_(d_1293_dt__update_hcf22_h0_):
                                    return D3_DC11((d_1292_dt__update__tmp_h0_).cf20, (d_1292_dt__update__tmp_h0_).cf21, d_1293_dt__update_hcf22_h0_)
                                return iife114_(_pat_let46_0)
                            return iife113_(pat_let_tv23_)
                        return iife112_(_pat_let45_0)
                    def iife110_(_pat_let44_0):
                        def iife115_(d_1294_dt__update__tmp_h1_):
                            def iife116_(_pat_let47_0):
                                def iife117_(d_1295_dt__update_hcf20_h0_):
                                    return D3_DC11(d_1295_dt__update_hcf20_h0_, (d_1294_dt__update__tmp_h1_).cf21, (d_1294_dt__update__tmp_h1_).cf22)
                                return iife117_(_pat_let47_0)
                            return iife116_(default__.fm3(pat_let_tv24_))
                        return iife115_(_pat_let44_0)
                    rhs207_ = not((self).f6)
                    rhs208_ = ((d_1283_v26_) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_1278_v21_), d_1290_v32_]))).set(default__.safeIndex(p0, len((d_1283_v26_) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_1278_v21_), d_1290_v32_])))), d_1290_v32_)
                    rhs209_ = iife110_(iife111_(D3_DC11((self).f9, len(d_1291_v33_), d_1280_v23_)))
                    rhs210_ = self.f4
                    rhs211_ = d_1288_v30_
                    lhs157_ = d_1285_v28_
                    lhs158_ = default__.safeIndex(133, (d_1285_v28_).length(0))
                    r0 = rhs207_
                    d_1283_v26_ = rhs208_
                    d_1284_v27_ = rhs209_
                    lhs157_[lhs158_] = rhs210_
                    d_1288_v30_ = rhs211_
                    pass
            pass
        r0 = not(((self).f7 if False else (self).f9))
        return r0

    def m4(self, p0, p1, globalState):
        d_1296_v0_: T2
        nw249_ = C4()
        nw249_.ctor__((self).f8, (self).f6, (self).f9, (self).f7, self.f4, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vknfksjs")))
        d_1296_v0_ = nw249_
        d_1297_v1_: _dafny.Map
        d_1297_v1_ = _dafny.Map({d_1296_v0_: (self).f8})
        d_1298_v2_: D2
        d_1298_v2_ = D2_DC8(p0, len(d_1297_v1_))
        pat_let_tv25_ = d_1296_v0_
        pat_let_tv26_ = d_1296_v0_
        pat_let_tv27_ = d_1296_v0_
        def lambda47_(source15_):
            if source15_.is_DC7:
                d_1299___mcc_h0_ = source15_.cf15
                d_1300_cf15_ = d_1299___mcc_h0_
                return (self).f8
            elif source15_.is_DC8:
                d_1301___mcc_h1_ = source15_.cf16
                d_1302___mcc_h2_ = source15_.cf17
                d_1303_cf17_ = d_1302___mcc_h2_
                d_1304_cf16_ = d_1301___mcc_h1_
                return d_1303_cf17_
            elif source15_.is_DC6:
                d_1305___mcc_h3_ = source15_.cf14
                d_1306_cf14_ = d_1305___mcc_h3_
                return (0) - (len((pat_let_tv25_).f5))
            elif True:
                d_1307___mcc_h4_ = source15_.cf18
                d_1308_cf18_ = d_1307___mcc_h4_
                return ((_dafny.MultiSet([(pat_let_tv26_).f7])) - (_dafny.MultiSet([(pat_let_tv27_).f6]))).cardinality

        (globalState).f1 = lambda47_(d_1298_v2_)
        d_1309_v3_: _dafny.MultiSet
        d_1309_v3_ = _dafny.MultiSet([False, (d_1296_v0_).f7])
        d_1310_i0_: int
        d_1310_i0_ = 0
        with _dafny.label("6"):
            while (d_1309_v3_).isdisjoint(default__.fm4((d_1296_v0_).f6, (self).f8, (d_1296_v0_).f6, p0, globalState)):
                with _dafny.c_label("6"):
                    if (d_1310_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1310_i0_ = (d_1310_i0_) + (1)
                    d_1311_v4_: _dafny.Map
                    d_1311_v4_ = _dafny.Map({(self).f9: (d_1296_v0_).f9})
                    d_1312_v5_: D12
                    d_1312_v5_ = D12_DC31((d_1311_v4_).set((self).f6, (d_1296_v0_).f7))
                    (globalState).f1 = default__.safeModuloInt((self).fm1((self).f6, (d_1312_v5_).cf60, (d_1296_v0_).f6, globalState), default__.safeModuloInt(len(_dafny.Set({True})), (self).f8))
                    d_1313_v6_: _dafny.Seq
                    d_1313_v6_ = _dafny.SeqWithoutIsStrInference([(self).f8, (d_1296_v0_).fm1((d_1296_v0_).f9, d_1311_v4_, (self).f7, globalState), default__.safeModuloInt(len(d_1311_v4_), (d_1296_v0_).f8), ((d_1296_v0_).f8) * (p0), (self).f8])
                    d_1313_v6_ = ((d_1313_v6_) + (d_1313_v6_)) + (d_1313_v6_)
                    index219_ = default__.safeIndex(374, ((self).f10).length(0))
                    ((self).f10)[index219_] = (d_1296_v0_).f7
                    index220_ = default__.safeIndex(374, ((self).f10).length(0))
                    ((self).f10)[index220_] = (d_1296_v0_).f9
                    d_1314_v7_: _dafny.Map
                    d_1314_v7_ = _dafny.Map({d_1313_v6_: True})
                    index221_ = default__.safeIndex(374, ((self).f10).length(0))
                    ((self).f10)[index221_] = (d_1314_v7_) != ((_dafny.Map({d_1313_v6_: (d_1296_v0_).f7}) if default__.fm27((d_1296_v0_).f8, (self).f9, False, globalState) else d_1314_v7_))
                    pass
            pass
        d_1315_v8_: _dafny.Array
        nw250_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
        d_1315_v8_ = nw250_
        nw251_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
        d_1315_v8_ = nw251_
        d_1316_v9_: C1
        nw252_ = C1()
        nw252_.ctor__((d_1296_v0_).f8, (self).f9, (self).f9, (d_1296_v0_).f7, d_1296_v0_.f4, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gedhynvk")))
        d_1316_v9_ = nw252_
        d_1317_v10_: _dafny.MultiSet
        d_1317_v10_ = _dafny.MultiSet([(D10_DC28((self).f9, (self).f8, (d_1296_v0_).f7, not((self).f7), d_1316_v9_)).cf51])
        hi7_ = 104
        for d_1318_i1_ in range((d_1317_v10_).cardinality, hi7_):
            d_1319_v11_: _dafny.Map
            d_1319_v11_ = _dafny.Map({(d_1296_v0_).f8: (d_1296_v0_).f8})
            d_1319_v11_ = (d_1319_v11_).set(p0, (0) - ((d_1296_v0_).f8))
            if (d_1296_v0_).f6:
                d_1320_v12_: bool
                d_1320_v12_ = True
                d_1321_v13_: D12
                d_1321_v13_ = D12_DC33(d_1320_v12_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))
                d_1322_v14_: _dafny.Map
                d_1322_v14_ = _dafny.Map({False: not((d_1296_v0_).f7)})
                pat_let_tv28_ = d_1296_v0_
                pat_let_tv29_ = globalState
                pat_let_tv30_ = d_1322_v14_
                pat_let_tv31_ = d_1296_v0_
                pat_let_tv32_ = globalState
                pat_let_tv33_ = d_1296_v0_
                pat_let_tv34_ = globalState
                def iife118_(_pat_let48_0):
                    def iife119_(d_1323_dt__update__tmp_h0_):
                        def iife120_(_pat_let49_0):
                            def iife121_(d_1324_dt__update_hcf63_h0_):
                                return D12_DC33(d_1324_dt__update_hcf63_h0_, (d_1323_dt__update__tmp_h0_).cf64)
                            return iife121_(_pat_let49_0)
                        return iife120_(default__.fm7(pat_let_tv28_.f4, (self).fm1(default__.fm3(pat_let_tv29_), pat_let_tv30_, (pat_let_tv31_).f6, pat_let_tv32_), (pat_let_tv33_).f7, pat_let_tv34_))
                    return iife119_(_pat_let48_0)
                d_1320_v12_ = (iife118_(d_1321_v13_)).cf63
                d_1325_v15_: _dafny.Set
                d_1326_v16_: _dafny.Array
                d_1327_v17_: int
                d_1328_v18_: _dafny.Array
                out25_: _dafny.Set
                out26_: _dafny.Array
                out27_: int
                out28_: _dafny.Array
                out25_, out26_, out27_, out28_ = (d_1316_v9_).m8(((d_1296_v0_).f6 if (self).f7 else (self).f6), (d_1296_v0_).f7, globalState)
                d_1325_v15_ = out25_
                d_1326_v16_ = out26_
                d_1327_v17_ = out27_
                d_1328_v18_ = out28_
                d_1329_v19_: _dafny.Array
                nw253_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_1329_v19_ = nw253_
                nw254_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_1329_v19_ = nw254_
                d_1320_v12_ = d_1320_v12_
                index222_ = default__.safeIndex(737, (d_1328_v18_).length(0))
                (d_1328_v18_)[index222_] = (self).f7
                index223_ = default__.safeIndex(737, (d_1328_v18_).length(0))
                (d_1328_v18_)[index223_] = (d_1296_v0_).f7
            elif True:
                d_1330_v20_: bool
                d_1330_v20_ = True
                rhs212_ = not (((p1)[(self).f8] if ((self).f8) in (p1) else (d_1296_v0_).f6)) or ((d_1296_v0_).f7)
                rhs213_ = not((d_1296_v0_).f7)
                d_1330_v20_ = rhs212_
                d_1330_v20_ = rhs213_
                rhs214_ = (self).f8
                rhs215_ = d_1296_v0_
                lhs159_ = globalState
                lhs159_.f1 = rhs214_
                d_1296_v0_ = rhs215_
                d_1330_v20_ = True
                d_1331_v21_: _dafny.Map
                d_1331_v21_ = _dafny.Map({d_1330_v20_: (self).f7})
                (globalState).f1 = (d_1316_v9_).fm1(True, d_1331_v21_, (self).f6, globalState)
                d_1332_v22_: _dafny.Map
                d_1332_v22_ = _dafny.Map({(d_1296_v0_).f8: True})
                d_1333_v23_: _dafny.Seq
                d_1333_v23_ = _dafny.SeqWithoutIsStrInference([(d_1296_v0_).f9])
                d_1332_v22_ = (d_1332_v22_).set(len((d_1319_v11_) | (d_1319_v11_)), ((d_1333_v23_)[default__.safeIndex((0) - (len(d_1319_v11_)), len(d_1333_v23_))]) or (not((d_1296_v0_).f7)))
            (globalState).f1 = (0) - (default__.safeModuloInt(((d_1317_v10_) - (d_1317_v10_)).cardinality, len((_dafny.SeqWithoutIsStrInference([d_1296_v0_.f4 for d_1334_i2_ in range(default__.abs(603))])).set(default__.safeIndex(default__.fm8((self).f8, (self).f9, (d_1296_v0_).f6, globalState), len(_dafny.SeqWithoutIsStrInference([d_1296_v0_.f4 for d_1335_i2_ in range(default__.abs(603))]))), d_1296_v0_.f4))))
            index224_ = default__.safeIndex(446, ((self).f10).length(0))
            ((self).f10)[index224_] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhpixis")))) != (d_1318_i1_)
            d_1336_v24_: bool
            d_1336_v24_ = True
            d_1337_v25_: _dafny.Array
            nw255_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_1337_v25_ = nw255_
            d_1338_v26_: _dafny.Seq
            d_1338_v26_ = _dafny.SeqWithoutIsStrInference([(d_1296_v0_).f7, (self).f7])
            d_1339_v27_: _dafny.Map
            d_1339_v27_ = _dafny.Map({(d_1296_v0_).f6: len(d_1338_v26_)})
            d_1340_v28_: D5
            d_1340_v28_ = D5_DC15((self).f10, (d_1296_v0_).f9, d_1339_v27_, 368)
            d_1341_v29_: _dafny.Seq
            d_1341_v29_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10, (self).f10, (self).f10])
            d_1342_v30_: _dafny.Array
            nw256_ = _dafny.Array(None, 29)
            nw256_[int(0)] = (self).f10
            nw256_[int(1)] = (self).f10
            nw256_[int(2)] = (self).f10
            nw256_[int(3)] = (self).f10
            nw256_[int(4)] = (d_1340_v28_).cf26
            nw256_[int(5)] = (self).f10
            nw256_[int(6)] = (self).f10
            nw256_[int(7)] = (self).f10
            nw256_[int(8)] = (self).f10
            nw256_[int(9)] = (d_1341_v29_)[default__.safeIndex((self).f8, len(d_1341_v29_))]
            nw256_[int(10)] = (self).f10
            nw256_[int(11)] = (self).f10
            nw256_[int(12)] = (self).f10
            nw256_[int(13)] = (self).f10
            nw256_[int(14)] = (self).f10
            nw256_[int(15)] = (self).f10
            nw256_[int(16)] = (self).f10
            nw256_[int(17)] = (d_1341_v29_)[default__.safeIndex((self).f8, len(d_1341_v29_))]
            nw256_[int(18)] = (self).f10
            nw256_[int(19)] = (self).f10
            nw256_[int(20)] = (self).f10
            nw256_[int(21)] = (self).f10
            nw256_[int(22)] = (self).f10
            nw256_[int(23)] = (self).f10
            nw256_[int(24)] = (self).f10
            nw256_[int(25)] = (self).f10
            nw256_[int(26)] = (self).f10
            nw256_[int(27)] = (self).f10
            nw256_[int(28)] = (self).f10
            d_1342_v30_ = nw256_
            index225_ = default__.safeIndex(901, (d_1337_v25_).length(0))
            (d_1337_v25_)[index225_] = d_1342_v30_
            d_1343_v31_: _dafny.Seq
            d_1343_v31_ = _dafny.SeqWithoutIsStrInference([(self).f8])
            d_1344_v32_: _dafny.Seq
            d_1344_v32_ = _dafny.SeqWithoutIsStrInference([d_1343_v31_])
            index226_ = default__.safeIndex(779, ((self).f11).length(0))
            ((self).f11)[index226_] = len((_dafny.SeqWithoutIsStrInference([p0, (self).f8])).set(default__.safeIndex((self).f8, len(_dafny.SeqWithoutIsStrInference([p0, (self).f8]))), len(d_1344_v32_)))
            d_1345_v33_: _dafny.Map
            d_1345_v33_ = _dafny.Map({(self).f6: False})
            d_1346_v34_: D8
            d_1346_v34_ = D8_DC23((d_1316_v9_).fm1((d_1296_v0_).f9, d_1345_v33_, (d_1296_v0_).f6, globalState), (self).f9)
            d_1347_v35_: _dafny.Set
            d_1347_v35_ = _dafny.Set({(self).f8})
            index227_ = default__.safeIndex(446, ((self).f10).length(0))
            index228_ = default__.safeIndex(901, (d_1337_v25_).length(0))
            index229_ = default__.safeIndex(779, ((self).f11).length(0))
            rhs216_ = ((d_1296_v0_).f9 if False else (d_1346_v34_).cf43)
            rhs217_ = default__.fm27((0) - ((d_1317_v10_).cardinality), (d_1347_v35_).issubset(d_1347_v35_), (self).f7, globalState)
            rhs218_ = d_1342_v30_
            rhs219_ = (0) - (-822)
            lhs160_ = (self).f10
            lhs161_ = default__.safeIndex(446, ((self).f10).length(0))
            lhs162_ = d_1337_v25_
            lhs163_ = default__.safeIndex(901, (d_1337_v25_).length(0))
            lhs164_ = (self).f11
            lhs165_ = default__.safeIndex(779, ((self).f11).length(0))
            lhs160_[lhs161_] = rhs216_
            d_1336_v24_ = rhs217_
            lhs162_[lhs163_] = rhs218_
            lhs164_[lhs165_] = rhs219_
        index230_ = default__.safeIndex(340, ((self).f11).length(0))
        ((self).f11)[index230_] = ((0) - (p0)) + ((self).f8)
        d_1348_v36_: _dafny.Set
        d_1348_v36_ = _dafny.Set({(d_1296_v0_).f8, (d_1296_v0_).f8, (d_1296_v0_).f8, p0, 852})
        d_1349_v37_: _dafny.Map
        d_1349_v37_ = _dafny.Map({(d_1296_v0_).f8: d_1317_v10_})
        d_1350_v38_: _dafny.Map
        d_1350_v38_ = _dafny.Map({d_1348_v36_: (D9_DC25((self).f9, (self).f6, len(d_1349_v37_))).cf47})
        index231_ = default__.safeIndex(340, ((self).f11).length(0))
        ((self).f11)[index231_] = ((d_1350_v38_)[_dafny.Set({(self).f8, (d_1296_v0_).f8, (self).f8, len((self).f5), p0})] if (_dafny.Set({(self).f8, (d_1296_v0_).f8, (self).f8, len((self).f5), p0})) in (d_1350_v38_) else default__.safeModuloInt(p0, (d_1296_v0_).f8))
        d_1351_v39_: bool
        d_1351_v39_ = False
        d_1351_v39_ = not((self).f6)

    @property
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11
