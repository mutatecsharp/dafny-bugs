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
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(160, 497):
                d_0_v0_: int = compr_0_
                if ((160) <= (d_0_v0_)) and ((d_0_v0_) < (497)):
                    coll0_[(d_0_v0_) + ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False, False, True, True]))).cardinality)] = False
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1_i0_ in range(default__.abs(926))])): not(True)})).keys.Elements:
                d_2_v1_: int = compr_1_
                if (d_2_v1_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1_i0_ in range(default__.abs(926))])): not(True)})):
                    coll1_[default__.safeModuloInt(d_2_v1_, len(_dafny.SeqWithoutIsStrInference([334, len(_dafny.SeqWithoutIsStrInference([False])), -137, 864, 242])))] = True
            return _dafny.Map(coll1_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({874: True}), _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([False, False]))): False}), _dafny.Map({773: not(True)})])) + ((_dafny.SeqWithoutIsStrInference([_dafny.Map({431: False}), iife0_()
        ])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({416: True}), iife1_()
        ])))

    @staticmethod
    def fm1(p0, p1, globalState):
        return (len((_dafny.Set({True})) | (_dafny.Set({not(True)})))) + ((len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([583])): False}))) * (-101))

    @staticmethod
    def fm8(p0, globalState):
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wlhgm"))): 865})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, False])): len(_dafny.Set({not(False)}))}))) | (_dafny.Map({-747: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfpaxux")))}))

    @staticmethod
    def fm9(p0, globalState):
        return (_dafny.Map({not(False): _dafny.Map({len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xtwhavygd")): 606})): not(False)})})) != (_dafny.Map({not(True): _dafny.Map({794: not(True)})}))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({False: False})) | (_dafny.Map({False: not(True)}))) | ((_dafny.Map({False: True}) if not(False) else _dafny.Map({True: True})))

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        return D1_DC4(not(False), ((0) - (-641)) >= (912))

    @staticmethod
    def fm12(p0, globalState):
        return _dafny.CodePoint('a')

    @staticmethod
    def fm18(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yathjppk"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))), (_dafny.Set({not(False)})).issubset(_dafny.Set({False, False}))])

    @staticmethod
    def fm20(p0, globalState):
        return ((D3_DC8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdigkgpmo")))).cf14) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybuwlsavw")))

    @staticmethod
    def fm21(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([True])) | (_dafny.MultiSet([True, False, not(False)]))) | (_dafny.MultiSet([True, True]))

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        if (False) == (not(not(True))):
            return _dafny.MultiSet([False])
        elif True:
            return (_dafny.MultiSet([True])) - (_dafny.MultiSet([True, False, True, not(False)]))

    @staticmethod
    def fm23(p0, globalState):
        return ((_dafny.Set({D0_DC2(D0_DC0(True)), D0_DC2(D0_DC1(False, len(_dafny.SeqWithoutIsStrInference([True, True, False, True, not(not(not(False)))])))), D0_DC2(D0_DC1(not(True), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqkjudxne")))})))), D0_DC2(D0_DC1(False, 478))})) | (_dafny.Set({D0_DC2(D0_DC1(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcy"))))), D0_DC2(D0_DC1(not(True), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True]))).cardinality))}))) | (_dafny.Set({D0_DC2(D0_DC2(D0_DC2(D0_DC2(D0_DC0(False))))), D0_DC2(D0_DC2(D0_DC1(not(True), len(_dafny.SeqWithoutIsStrInference([650])))))}))

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        return D0_DC2(D0_DC2(D0_DC1(True, (0) - ((_dafny.MultiSet([False])).cardinality))))

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjukjem"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wuyw")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epjwidlg"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ieovtcopc"))))

    @staticmethod
    def fm26(p0, p1, globalState):
        source0_ = D9_DC27(D9_DC26())
        if source0_.is_DC26:
            return _dafny.CodePoint('l')
        elif source0_.is_DC25:
            d_3___mcc_h0_ = source0_.cf50
            d_4_cf50_ = d_3___mcc_h0_
            return _dafny.CodePoint('u')
        elif True:
            d_5___mcc_h1_ = source0_.cf51
            d_6_cf51_ = d_5___mcc_h1_
            if False:
                return _dafny.CodePoint('j')
            elif True:
                return _dafny.CodePoint('i')

    @staticmethod
    def fm27(p0, p1, p2, p3, globalState):
        return (_dafny.Map({True: _dafny.Map({False: False})})) | ((_dafny.Map({False: _dafny.Map({True: True})})) | (_dafny.Map({False: _dafny.Map({not(True): False})})))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(354, 741):
                d_7_v0_: int = compr_2_
                if ((354) <= (d_7_v0_)) and ((d_7_v0_) < (741)):
                    coll2_ = coll2_.union(_dafny.Set([default__.safeModuloInt(d_7_v0_, 168)]))
            return _dafny.Set(coll2_)
        return (_dafny.Map({True: D1_DC5(not(False), _dafny.Map({False: not(True)}), len(_dafny.SeqWithoutIsStrInference([57])), len(_dafny.Map({True: (_dafny.MultiSet([True])).cardinality})), _dafny.Map({False: 859}))})) | ((_dafny.Map({not(True): D1_DC5(True, _dafny.Map({False: True}), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "np"))), 692, _dafny.Map({False: (_dafny.MultiSet([not(False)])).cardinality}))})) | (_dafny.Map({True: D1_DC5(True, _dafny.Map({False: True}), len(iife2_()
), -782, _dafny.Map({True: 652}))})))

    @staticmethod
    def fm29(p0, p1, p2, globalState):
        return not ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eaasr")))) > ((D15_DC37(False, _dafny.Map({True: True}), len(_dafny.SeqWithoutIsStrInference([218 for d_8_i0_ in range(default__.abs(-626))])))).cf67)) or (False)

    @staticmethod
    def fm30(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + ((_dafny.SeqWithoutIsStrInference([True, True])) + (_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm31(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([437, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_9_i0_ in range(default__.abs(-733))]))])) + ((D10_DC28(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False})), len(_dafny.SeqWithoutIsStrInference([True]))]))).cf52)

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        return _dafny.Set({True, (_dafny.Set({True, False})) != (_dafny.Set({True, False, True}))})

    @staticmethod
    def fm33(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([D1_DC4(False, False) for d_10_i0_ in range(default__.abs(261))])

    @staticmethod
    def fm34(p0, globalState):
        return (_dafny.Map({938: True})) | ((_dafny.Map({(0) - (len(_dafny.Set({len(_dafny.Map({True: True}))}))): True})) | (_dafny.Map({495: False})))

    @staticmethod
    def fm35(globalState):
        return _dafny.Map({(_dafny.Set({False, True})).intersection(_dafny.Set({(D3_DC10(True, _dafny.Map({False: False}))).cf16, True, True})): (536) >= (-239)})

    @staticmethod
    def fm36(globalState):
        return _dafny.Map({(30) > (650): default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.Set({False, not(True)})))})

    @staticmethod
    def fm37(p0, p1, globalState):
        return (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({False: not(False)})))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({676: True})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crt")))])))) | ((_dafny.MultiSet([-349, -137, 200])) | (_dafny.MultiSet([633])))

    @staticmethod
    def fm38(p0, p1, globalState):
        return (_dafny.Set({_dafny.Set({False, True})})).intersection((_dafny.Set({_dafny.Set({not(True)})})).intersection(_dafny.Set({_dafny.Set({False, True, True})})))

    @staticmethod
    def fm39(p0, p1, p2, p3, globalState):
        return D3_DC8(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_11_i0_ in range(default__.abs(-876))]))

    @staticmethod
    def fm40(p0, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: _dafny.Seq
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D6_DC16(_dafny.SeqWithoutIsStrInference([False, True]))])])).Elements:
                d_12_v0_: _dafny.Seq = compr_3_
                if (d_12_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D6_DC16(_dafny.SeqWithoutIsStrInference([False, True]))])])):
                    coll3_[d_12_v0_] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-294, (_dafny.MultiSet([238, 653, 98])).cardinality, len(_dafny.SeqWithoutIsStrInference([False]))]))]))).isdisjoint(_dafny.MultiSet([-670]))
            return _dafny.Map(coll3_)
        return iife3_()
        

    @staticmethod
    def fm41(globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(-842, 321):
                d_15_v0_: int = compr_4_
                if ((-842) <= (d_15_v0_)) and ((d_15_v0_) < (321)):
                    coll4_ = coll4_.union(_dafny.Set([default__.safeModuloInt(d_15_v0_, 646)]))
            return _dafny.Set(coll4_)
        return (_dafny.Set({((D7_DC19(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False])))).cf34).cardinality, -751, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_13_i0_ in range(default__.abs(146))])), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len((D10_DC28(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({254: False}))]))).cf52) for d_14_i1_ in range(default__.abs(467))]))).cardinality, 192})) | (((D8_DC23(319, _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hm")))}), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tcggd")))])).cardinality, 581, -312)).cf41) | (iife4_()
        ))

    @staticmethod
    def Main(noArgsParameter__):
        d_16_v0_: bool
        d_16_v0_ = True
        d_17_v1_: _dafny.Map
        d_17_v1_ = _dafny.Map({d_16_v0_: (0) - (-313)})
        d_18_v2_: _dafny.Seq
        d_18_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_16_v0_, d_16_v0_})])
        d_19_v3_: _dafny.Array
        nw0_ = _dafny.Array(False, 26)
        d_19_v3_ = nw0_
        d_20_v4_: int
        d_20_v4_ = 333
        d_21_v5_: _dafny.Map
        d_21_v5_ = _dafny.Map({d_20_v4_: d_20_v4_})
        d_22_v6_: _dafny.Map
        d_22_v6_ = _dafny.Map({d_19_v3_: d_21_v5_})
        d_23_v7_: _dafny.Seq
        d_23_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hllu"))
        d_24_v8_: _dafny.Array
        def lambda0_(d_25_v4_, d_26_v0_):
            def lambda1_(d_27_i0_):
                return _dafny.MultiSet([d_25_v4_, (0) - (d_25_v4_), (_dafny.MultiSet([d_26_v0_, d_26_v0_])).cardinality, 837])

            return lambda1_

        init0_ = lambda0_(d_20_v4_, d_16_v0_)
        nw1_ = _dafny.Array(None, 12)
        for i0_0_ in range(nw1_.length(0)):
            nw1_[i0_0_] = init0_(i0_0_)
        d_24_v8_ = nw1_
        d_28_globalState_: GlobalState
        nw2_ = GlobalState()
        nw2_.ctor__(d_17_v1_, True, -406, False, d_18_v2_, 145, False, 873, 506, True, d_22_v6_, d_23_v7_, 895, 755, 670, True, 550, 75, d_24_v8_, True, False, True, -889)
        d_28_globalState_ = nw2_
        d_29_v9_: _dafny.Seq
        d_29_v9_ = _dafny.SeqWithoutIsStrInference([len(d_23_v7_)])
        hi0_ = ((d_29_v9_)[default__.safeIndex(d_20_v4_, len(d_29_v9_))] if d_16_v0_ else d_20_v4_)
        for d_30_i1_ in range((d_20_v4_) - (len(d_23_v7_)), hi0_):
            d_31_v10_: str
            d_31_v10_ = _dafny.CodePoint('v')
            d_32_v11_: _dafny.Map
            d_32_v11_ = _dafny.Map({d_31_v10_: _dafny.Map({-829: d_16_v0_})})
            d_33_v13_: _dafny.Map
            d_33_v13_ = _dafny.Map({d_20_v4_: d_16_v0_})
            d_34_v14_: _dafny.Seq
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in (d_33_v13_).keys.Elements:
                    d_35_v12_: int = compr_5_
                    if (d_35_v12_) in (d_33_v13_):
                        coll5_[default__.safeDivisionInt(d_35_v12_, d_20_v4_)] = d_16_v0_
                return _dafny.Map(coll5_)
            d_34_v14_ = _dafny.SeqWithoutIsStrInference([iife5_()
            ])
            d_36_v16_: _dafny.Seq
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(901, 437):
                    d_37_v15_: int = compr_6_
                    if ((901) <= (d_37_v15_)) and ((d_37_v15_) < (437)):
                        coll6_[(d_37_v15_) - ((0) - (len(_dafny.Set({d_16_v0_, d_16_v0_}))))] = d_16_v0_
                return _dafny.Map(coll6_)
            d_36_v16_ = _dafny.SeqWithoutIsStrInference([((d_32_v11_)[d_31_v10_] if (d_31_v10_) in (d_32_v11_) else (d_34_v14_)[default__.safeIndex(len(d_29_v9_), len(d_34_v14_))]), iife6_()
            ])
            d_38_v17_: _dafny.Set
            d_38_v17_ = _dafny.Set({not(d_16_v0_)})
            d_36_v16_ = default__.fm0(((0) - (-98) if d_16_v0_ else default__.fm1(default__.fm1(d_20_v4_, d_16_v0_, d_28_globalState_), not(d_16_v0_), d_28_globalState_)), d_20_v4_, d_23_v7_, len(d_38_v17_), d_28_globalState_)
            (d_28_globalState_).f5 = d_20_v4_
            d_39_v18_: _dafny.Set
            d_39_v18_ = _dafny.Set({d_33_v13_})
            d_39_v18_ = (d_39_v18_) - ((d_39_v18_).intersection(d_39_v18_))
            rhs0_ = d_16_v0_
            rhs1_ = d_31_v10_
            lhs0_ = d_28_globalState_
            lhs0_.f9 = rhs0_
            d_31_v10_ = rhs1_
        (d_28_globalState_).f9 = d_16_v0_
        if d_16_v0_:
            d_40_v19_: _dafny.Set
            d_40_v19_ = _dafny.Set({d_16_v0_, d_16_v0_})
            d_41_v20_: _dafny.Map
            d_41_v20_ = _dafny.Map({d_16_v0_: (d_40_v19_).ispropersubset(d_40_v19_)})
            d_42_v21_: _dafny.Seq
            d_42_v21_ = _dafny.SeqWithoutIsStrInference([d_16_v0_])
            (d_28_globalState_).f6 = ((d_41_v20_)[(d_42_v21_)[default__.safeIndex(d_20_v4_, len(d_42_v21_))]] if ((d_42_v21_)[default__.safeIndex(d_20_v4_, len(d_42_v21_))]) in (d_41_v20_) else not(d_16_v0_))
            (d_28_globalState_).f14 = d_20_v4_
            d_43_v22_: _dafny.Set
            d_43_v22_ = _dafny.Set({d_23_v7_})
            if not((d_43_v22_).issubset((d_43_v22_) - (d_43_v22_))):
                d_44_v23_: _dafny.Array
                def lambda2_(d_45_v0_):
                    def lambda3_(d_46_i2_):
                        return _dafny.MultiSet([_dafny.MultiSet([False, d_45_v0_])])

                    return lambda3_

                init1_ = lambda2_(d_16_v0_)
                nw3_ = _dafny.Array(None, 8)
                for i0_1_ in range(nw3_.length(0)):
                    nw3_[i0_1_] = init1_(i0_1_)
                d_44_v23_ = nw3_
                d_47_v24_: _dafny.MultiSet
                d_47_v24_ = _dafny.MultiSet([d_16_v0_])
                d_48_v25_: _dafny.MultiSet
                d_48_v25_ = _dafny.MultiSet([d_47_v24_])
                index0_ = default__.safeIndex(536, (d_44_v23_).length(0))
                (d_44_v23_)[index0_] = d_48_v25_
                d_49_v26_: _dafny.Map
                d_49_v26_ = _dafny.Map({d_16_v0_: d_48_v25_})
                d_50_v27_: _dafny.Array
                nw4_ = _dafny.Array(int(0), 24)
                d_50_v27_ = nw4_
                d_51_v28_: _dafny.Map
                d_51_v28_ = _dafny.Map({d_50_v27_: d_16_v0_})
                d_52_v30_: str
                d_52_v30_ = _dafny.CodePoint('t')
                d_53_v31_: _dafny.MultiSet
                d_53_v31_ = _dafny.MultiSet([d_52_v30_])
                d_54_v32_: _dafny.Set
                def iife7_():
                    coll7_ = _dafny.Map()
                    compr_7_: str
                    for compr_7_ in (d_53_v31_).Elements:
                        d_55_v29_: str = compr_7_
                        if (d_55_v29_) in (d_53_v31_):
                            coll7_[d_55_v29_] = d_16_v0_
                    return _dafny.Map(coll7_)
                d_54_v32_ = _dafny.Set({iife7_()
                })
                index1_ = default__.safeIndex(536, (d_44_v23_).length(0))
                rhs2_ = (d_48_v25_) - (((d_49_v26_)[((d_51_v28_)[d_50_v27_] if (d_50_v27_) in (d_51_v28_) else d_16_v0_)] if (((d_51_v28_)[d_50_v27_] if (d_50_v27_) in (d_51_v28_) else d_16_v0_)) in (d_49_v26_) else d_48_v25_))
                rhs3_ = d_17_v1_
                rhs4_ = len(((d_54_v32_).intersection(d_54_v32_)) | (d_54_v32_))
                lhs1_ = d_44_v23_
                lhs2_ = default__.safeIndex(536, (d_44_v23_).length(0))
                lhs3_ = d_28_globalState_
                lhs1_[lhs2_] = rhs2_
                d_17_v1_ = rhs3_
                lhs3_.f5 = rhs4_
                d_56_v33_: _dafny.Array
                nw5_ = _dafny.Array(_dafny.Array(None, 0), 5)
                d_56_v33_ = nw5_
                rhs5_ = (d_56_v33_ if (d_42_v21_)[default__.safeIndex(125, len(d_42_v21_))] else d_56_v33_)
                rhs6_ = (d_20_v4_) - (d_20_v4_)
                lhs4_ = d_28_globalState_
                d_56_v33_ = rhs5_
                lhs4_.f17 = rhs6_
                d_57_v34_: C2
                nw6_ = C2()
                nw6_.ctor__(d_20_v4_)
                d_57_v34_ = nw6_
                index2_ = default__.safeIndex(727, (d_56_v33_).length(0))
                (d_56_v33_)[index2_] = (d_50_v27_ if d_16_v0_ else d_50_v27_)
                index3_ = default__.safeIndex(727, (d_56_v33_).length(0))
                rhs7_ = not(d_16_v0_)
                rhs8_ = d_16_v0_
                rhs9_ = d_17_v1_
                rhs10_ = d_50_v27_
                lhs5_ = d_28_globalState_
                lhs6_ = d_28_globalState_
                lhs7_ = d_28_globalState_
                lhs8_ = d_56_v33_
                lhs9_ = default__.safeIndex(727, (d_56_v33_).length(0))
                lhs5_.f1 = rhs7_
                lhs6_.f1 = rhs8_
                lhs7_.f0 = rhs9_
                lhs8_[lhs9_] = rhs10_
                d_58_v35_: _dafny.Array
                nw7_ = _dafny.Array(None, 27)
                d_58_v35_ = nw7_
                d_59_v36_: _dafny.Set
                d_59_v36_ = _dafny.Set({(_dafny.MultiSet([d_58_v35_])).cardinality, len(d_21_v5_)})
                d_60_v37_: bool
                out0_: bool
                out0_ = (d_57_v34_).m7(d_19_v3_, d_16_v0_, d_16_v0_, default__.fm29(len(d_59_v36_), d_20_v4_, d_16_v0_, d_28_globalState_), d_28_globalState_)
                d_60_v37_ = out0_
            elif True:
                d_61_v38_: str
                d_61_v38_ = _dafny.CodePoint('j')
                rhs11_ = d_61_v38_
                rhs12_ = d_20_v4_
                lhs10_ = d_28_globalState_
                d_61_v38_ = rhs11_
                lhs10_.f14 = rhs12_
                d_62_v39_: _dafny.Set
                d_62_v39_ = _dafny.Set({d_19_v3_})
                d_63_v40_: C0
                nw8_ = C0()
                nw8_.ctor__((d_62_v39_) | (d_62_v39_), d_20_v4_, d_20_v4_)
                d_63_v40_ = nw8_
                d_64_v41_: C8
                nw9_ = C8()
                nw9_.ctor__(d_20_v4_, d_16_v0_, d_16_v0_)
                d_64_v41_ = nw9_
                d_64_v41_ = d_64_v41_
                d_65_v42_: _dafny.Array
                nw10_ = _dafny.Array(_dafny.Map({}), 1)
                d_65_v42_ = nw10_
                d_66_v43_: T1
                nw11_ = C4()
                nw11_.ctor__(d_16_v0_, d_16_v0_, d_16_v0_)
                d_66_v43_ = nw11_
                d_67_v44_: _dafny.Map
                d_67_v44_ = _dafny.Map({d_66_v43_: len(d_42_v21_)})
                d_68_v45_: _dafny.Map
                d_68_v45_ = _dafny.Map({d_63_v40_: d_67_v44_})
                index4_ = default__.safeIndex(919, (d_65_v42_).length(0))
                (d_65_v42_)[index4_] = (d_68_v45_ if d_16_v0_ else d_68_v45_)
                index5_ = default__.safeIndex(919, (d_65_v42_).length(0))
                (d_65_v42_)[index5_] = (d_68_v45_) | ((d_68_v45_) | ((d_68_v45_).set(d_63_v40_, d_67_v44_)))
                d_69_v46_: _dafny.Array
                nw12_ = _dafny.Array(None, 19)
                d_69_v46_ = nw12_
                d_70_v47_: _dafny.Set
                d_70_v47_ = _dafny.Set({d_69_v46_})
                d_71_v48_: _dafny.Array
                nw13_ = _dafny.Array(None, 17)
                d_71_v48_ = nw13_
                d_72_v49_: C4
                nw14_ = C4()
                nw14_.ctor__(d_66_v43_.f24, d_66_v43_.f24, d_16_v0_)
                d_72_v49_ = nw14_
                index6_ = default__.safeIndex(574, (d_71_v48_).length(0))
                (d_71_v48_)[index6_] = d_72_v49_
                index7_ = default__.safeIndex(574, (d_71_v48_).length(0))
                rhs13_ = ((len(_dafny.Set({d_63_v40_.f31})) if (d_66_v43_).f25 else d_63_v40_.f31)) < (d_20_v4_)
                rhs14_ = d_70_v47_
                rhs15_ = d_63_v40_.f31
                rhs16_ = d_72_v49_
                lhs11_ = d_28_globalState_
                lhs12_ = d_63_v40_
                lhs13_ = d_71_v48_
                lhs14_ = default__.safeIndex(574, (d_71_v48_).length(0))
                lhs11_.f15 = rhs13_
                d_70_v47_ = rhs14_
                lhs12_.f31 = rhs15_
                lhs13_[lhs14_] = rhs16_
            d_73_v50_: _dafny.Set
            d_73_v50_ = _dafny.Set({d_19_v3_, d_19_v3_, d_19_v3_})
            d_74_v51_: C0
            nw15_ = C0()
            nw15_.ctor__(d_73_v50_, d_20_v4_, default__.safeModuloInt(d_20_v4_, d_20_v4_))
            d_74_v51_ = nw15_
            d_75_v52_: _dafny.Map
            d_75_v52_ = _dafny.Map({(d_23_v7_)[default__.safeIndex(d_74_v51_.f31, len(d_23_v7_))]: _dafny.SeqWithoutIsStrInference([False, d_16_v0_])})
            (d_28_globalState_).f8 = (0) - ((d_74_v51_.f31) * (((0) - (len(d_75_v52_))) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ja"))))))
        elif True:
            d_76_v53_: D3
            d_76_v53_ = D3_DC10(d_16_v0_, _dafny.Map({d_16_v0_: d_16_v0_}))
            d_77_v54_: _dafny.Seq
            d_77_v54_ = _dafny.SeqWithoutIsStrInference([not((d_20_v4_) <= (d_20_v4_)), (d_76_v53_).cf16, d_16_v0_])
            d_77_v54_ = ((d_77_v54_).set(default__.safeIndex(d_20_v4_, len(d_77_v54_)), d_16_v0_)) + (d_77_v54_)
            d_78_v55_: _dafny.Set
            d_78_v55_ = _dafny.Set({(len(d_23_v7_)) == (d_20_v4_), not(d_16_v0_)})
            d_79_v56_: _dafny.MultiSet
            d_79_v56_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_80_i4_ in range(default__.abs(626))])])
            rhs17_ = d_19_v3_
            rhs18_ = d_78_v55_
            rhs19_ = ((((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_81_i3_ in range(default__.abs(689))])])).set(d_23_v7_, default__.abs(d_20_v4_))) | (d_79_v56_)) - (d_79_v56_)).cardinality
            d_19_v3_ = rhs17_
            d_78_v55_ = rhs18_
            d_20_v4_ = rhs19_
            d_82_v57_: C3
            nw16_ = C3()
            nw16_.ctor__(d_16_v0_, (-487 if True else d_20_v4_), not(not(d_16_v0_)), not((d_20_v4_) > (-148)), d_20_v4_)
            d_82_v57_ = nw16_
            d_83_v58_: _dafny.Set
            d_83_v58_ = _dafny.Set({d_19_v3_})
            d_84_v59_: _dafny.MultiSet
            d_84_v59_ = _dafny.MultiSet([not(d_16_v0_), d_16_v0_])
            d_85_v60_: C0
            nw17_ = C0()
            nw17_.ctor__((_dafny.Set({d_19_v3_})) | (d_83_v58_), d_20_v4_, (d_20_v4_) + ((d_84_v59_).cardinality))
            d_85_v60_ = nw17_
            d_86_v61_: C0
            nw18_ = C0()
            nw18_.ctor__(_dafny.Set({d_19_v3_, d_19_v3_}), len(_dafny.Map({d_19_v3_: (d_82_v57_).f28})), (d_82_v57_).f29)
            d_86_v61_ = nw18_
        d_87_v62_: T1
        nw19_ = C4()
        nw19_.ctor__(d_16_v0_, d_16_v0_, d_16_v0_)
        d_87_v62_ = nw19_
        d_88_v63_: _dafny.Array
        nw20_ = _dafny.Array(D6.default()(), 24)
        d_88_v63_ = nw20_
        d_89_v64_: _dafny.Map
        d_89_v64_ = _dafny.Map({d_87_v62_: d_88_v63_})
        d_90_v65_: _dafny.Map
        d_90_v65_ = _dafny.Map({((d_89_v64_)[d_87_v62_] if (d_87_v62_) in (d_89_v64_) else d_88_v63_): (22) == (d_20_v4_)})
        d_90_v65_ = (d_90_v65_).set(d_88_v63_, d_16_v0_)
        index8_ = default__.safeIndex(621, (d_19_v3_).length(0))
        (d_19_v3_)[index8_] = d_16_v0_
        index9_ = default__.safeIndex(621, (d_19_v3_).length(0))
        (d_19_v3_)[index9_] = d_87_v62_.f24
        (d_28_globalState_).f1 = ((d_20_v4_) - (660)) not in (default__.fm34(d_20_v4_, d_28_globalState_))
        d_91_v66_: _dafny.Array
        nw21_ = _dafny.Array(int(0), 28)
        d_91_v66_ = nw21_
        index10_ = default__.safeIndex(328, (d_91_v66_).length(0))
        (d_91_v66_)[index10_] = d_20_v4_
        index11_ = default__.safeIndex(328, (d_91_v66_).length(0))
        rhs20_ = d_20_v4_
        rhs21_ = not((d_87_v62_).f25)
        rhs22_ = d_16_v0_
        lhs15_ = d_91_v66_
        lhs16_ = default__.safeIndex(328, (d_91_v66_).length(0))
        lhs17_ = d_28_globalState_
        lhs18_ = d_28_globalState_
        lhs15_[lhs16_] = rhs20_
        lhs17_.f9 = rhs21_
        lhs18_.f1 = rhs22_
        d_92_v67_: _dafny.Array
        nw22_ = _dafny.Array(_dafny.Array(None, 0), 14)
        d_92_v67_ = nw22_
        d_93_v68_: _dafny.Array
        nw23_ = _dafny.Array(_dafny.Set({}), 28)
        d_93_v68_ = nw23_
        index12_ = default__.safeIndex(699, (d_92_v67_).length(0))
        (d_92_v67_)[index12_] = d_93_v68_
        index13_ = default__.safeIndex(699, (d_92_v67_).length(0))
        (d_92_v67_)[index13_] = d_93_v68_
        d_94_v69_: _dafny.Set
        d_94_v69_ = _dafny.Set({d_19_v3_, d_19_v3_, d_19_v3_})
        d_95_v70_: T0
        nw24_ = C0()
        nw24_.ctor__(d_94_v69_, (d_20_v4_) * ((d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))]), d_20_v4_)
        d_95_v70_ = nw24_
        d_96_v71_: _dafny.MultiSet
        d_96_v71_ = _dafny.MultiSet([d_87_v62_.f24])
        d_97_v72_: _dafny.Map
        d_97_v72_ = _dafny.Map({d_96_v71_: d_19_v3_})
        rhs23_ = d_95_v70_
        rhs24_ = (0) - ((d_95_v70_).f23)
        rhs25_ = d_97_v72_
        lhs19_ = d_28_globalState_
        d_95_v70_ = rhs23_
        lhs19_.f8 = rhs24_
        d_97_v72_ = rhs25_
        d_98_v73_: _dafny.Map
        d_98_v73_ = _dafny.Map({(d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))]: d_16_v0_})
        d_98_v73_ = (d_98_v73_).set((d_95_v70_).f23, (not(d_87_v62_.f24) if (d_87_v62_).f25 else (d_87_v62_).f25))
        d_99_v74_: _dafny.Seq
        d_99_v74_ = _dafny.SeqWithoutIsStrInference([d_87_v62_.f24])
        d_100_v75_: _dafny.Map
        d_100_v75_ = _dafny.Map({len(d_99_v74_): d_23_v7_})
        d_23_v7_ = (((d_100_v75_)[456] if (456) in (d_100_v75_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gstyka")))) + ((d_23_v7_) + (d_23_v7_))
        d_101_v76_: _dafny.Map
        d_101_v76_ = _dafny.Map({(d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))]: (d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))]})
        d_102_v77_: _dafny.Map
        d_102_v77_ = _dafny.Map({d_101_v76_: len(d_21_v5_)})
        d_103_v78_: _dafny.Map
        d_103_v78_ = _dafny.Map({d_102_v77_: (d_87_v62_).f25})
        d_104_v79_: C1
        nw25_ = C1()
        nw25_.ctor__(True, ((d_103_v78_)[d_102_v77_] if (d_102_v77_) in (d_103_v78_) else (d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))]))
        d_104_v79_ = nw25_
        d_105_i5_: int
        d_105_i5_ = 0
        with _dafny.label("0"):
            while (d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))]:
                with _dafny.c_label("0"):
                    if (d_105_i5_) >= (100):
                        raise _dafny.Break("0")
                    d_105_i5_ = (d_105_i5_) + (1)
                    d_106_v80_: C0
                    nw26_ = C0()
                    nw26_.ctor__(d_94_v69_, (d_95_v70_).f23, (d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))])
                    d_106_v80_ = nw26_
                    d_107_v81_: _dafny.Map
                    d_107_v81_ = _dafny.Map({default__.safeModuloInt((d_95_v70_).f23, (_dafny.MultiSet(d_99_v74_)).cardinality): d_106_v80_})
                    d_107_v81_ = (d_107_v81_).set(default__.safeDivisionInt((d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))], d_20_v4_), d_106_v80_)
                    index14_ = default__.safeIndex(621, (d_19_v3_).length(0))
                    (d_19_v3_)[index14_] = False
                    d_108_v82_: _dafny.MultiSet
                    d_108_v82_ = _dafny.MultiSet([(d_95_v70_).f23, d_106_v80_.f31, len(d_23_v7_)])
                    d_109_v83_: C7
                    nw27_ = C7()
                    nw27_.ctor__()
                    d_109_v83_ = nw27_
                    d_110_v84_: _dafny.Map
                    d_110_v84_ = _dafny.Map({d_109_v83_: (d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))]})
                    (d_28_globalState_).f9 = (d_108_v82_) != (_dafny.MultiSet([(d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))], len((d_29_v9_).set(default__.safeIndex(((d_21_v5_)[d_106_v80_.f31] if (d_106_v80_.f31) in (d_21_v5_) else ((d_21_v5_)[(d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))]] if ((d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))]) in (d_21_v5_) else d_106_v80_.f31)), len(d_29_v9_)), len(d_110_v84_)))]))
                    if ((d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))]) >= ((d_95_v70_).f23):
                        index15_ = default__.safeIndex(328, (d_91_v66_).length(0))
                        (d_91_v66_)[index15_] = ((d_20_v4_ if (d_87_v62_).f25 else (d_95_v70_).f23)) - ((d_95_v70_).f23)
                        d_111_v85_: str
                        d_111_v85_ = _dafny.CodePoint('a')
                        d_111_v85_ = d_111_v85_
                        (d_28_globalState_).f17 = 297
                        d_112_v86_: _dafny.Map
                        d_112_v86_ = _dafny.Map({d_101_v76_: d_87_v62_})
                        d_113_v87_: _dafny.Map
                        d_113_v87_ = _dafny.Map({(d_87_v62_).f25: d_87_v62_})
                        d_114_v88_: _dafny.Array
                        nw28_ = _dafny.Array(None, 23)
                        nw28_[int(0)] = (d_87_v62_ if not((d_87_v62_).f25) else d_87_v62_)
                        nw28_[int(1)] = d_87_v62_
                        nw28_[int(2)] = d_87_v62_
                        nw28_[int(3)] = d_87_v62_
                        nw28_[int(4)] = d_87_v62_
                        nw28_[int(5)] = d_87_v62_
                        nw28_[int(6)] = d_87_v62_
                        nw28_[int(7)] = d_87_v62_
                        nw28_[int(8)] = d_87_v62_
                        nw28_[int(9)] = d_87_v62_
                        nw28_[int(10)] = d_87_v62_
                        nw28_[int(11)] = d_87_v62_
                        nw28_[int(12)] = d_87_v62_
                        nw28_[int(13)] = d_87_v62_
                        nw28_[int(14)] = (d_87_v62_ if (d_87_v62_).f25 else ((d_112_v86_)[d_101_v76_] if (d_101_v76_) in (d_112_v86_) else ((d_113_v87_)[d_16_v0_] if (d_16_v0_) in (d_113_v87_) else d_87_v62_)))
                        nw28_[int(15)] = (d_87_v62_ if d_87_v62_.f24 else d_87_v62_)
                        nw28_[int(16)] = d_87_v62_
                        nw28_[int(17)] = d_87_v62_
                        nw28_[int(18)] = d_87_v62_
                        nw28_[int(19)] = d_87_v62_
                        nw28_[int(20)] = d_87_v62_
                        nw28_[int(21)] = d_87_v62_
                        nw28_[int(22)] = d_87_v62_
                        d_114_v88_ = nw28_
                        d_114_v88_ = d_114_v88_
                        d_115_v91_: _dafny.MultiSet
                        def iife8_():
                            coll8_ = _dafny.Map()
                            compr_8_: int
                            for compr_8_ in _dafny.IntegerRange(391, 463):
                                d_116_v89_: int = compr_8_
                                if ((391) <= (d_116_v89_)) and ((d_116_v89_) < (463)):
                                    coll8_[default__.safeDivisionInt(d_116_v89_, (d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))])] = (d_87_v62_).f25
                            return _dafny.Map(coll8_)
                        def iife9_():
                            coll9_ = _dafny.Map()
                            compr_9_: int
                            for compr_9_ in (d_98_v73_).keys.Elements:
                                d_117_v90_: int = compr_9_
                                if (d_117_v90_) in (d_98_v73_):
                                    coll9_[default__.safeDivisionInt(d_117_v90_, d_20_v4_)] = True
                            return _dafny.Map(coll9_)
                        d_115_v91_ = _dafny.MultiSet([d_98_v73_, (iife8_()
                        ) | (iife9_()
                        ), d_98_v73_, d_98_v73_, d_98_v73_])
                        rhs26_ = d_115_v91_
                        rhs27_ = d_16_v0_
                        lhs20_ = d_28_globalState_
                        d_115_v91_ = rhs26_
                        lhs20_.f1 = rhs27_
                    elif True:
                        d_118_v92_: C3
                        nw29_ = C3()
                        nw29_.ctor__((d_87_v62_).f25, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('s'): (d_87_v62_).f25}) for d_119_i6_ in range(default__.abs(288))])), (d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))], d_87_v62_.f24, d_106_v80_.f31)
                        d_118_v92_ = nw29_
                        d_120_v93_: _dafny.Seq
                        d_120_v93_ = _dafny.SeqWithoutIsStrInference([d_118_v92_, d_118_v92_])
                        d_120_v93_ = (d_120_v93_) + ((d_120_v93_) + (d_120_v93_))
                        d_121_v94_: _dafny.Map
                        d_121_v94_ = _dafny.Map({d_106_v80_: d_96_v71_})
                        rhs28_ = d_87_v62_.f24
                        rhs29_ = ((d_121_v94_) | (d_121_v94_)) | (d_121_v94_)
                        lhs21_ = d_87_v62_
                        lhs21_.f24 = rhs28_
                        d_121_v94_ = rhs29_
                        d_122_v95_: C3
                        d_122_v95_ = d_118_v92_
                        index16_ = default__.safeIndex(621, (d_19_v3_).length(0))
                        rhs30_ = (d_87_v62_).f25
                        rhs31_ = False
                        rhs32_ = (d_122_v95_)
                        rhs33_ = d_19_v3_
                        lhs22_ = d_19_v3_
                        lhs23_ = default__.safeIndex(621, (d_19_v3_).length(0))
                        lhs24_ = d_28_globalState_
                        lhs22_[lhs23_] = rhs30_
                        lhs24_.f6 = rhs31_
                        d_118_v92_ = rhs32_
                        d_19_v3_ = rhs33_
                        d_123_v96_: bool
                        d_124_v97_: T0
                        d_125_v98_: bool
                        d_126_v99_: int
                        out1_: bool
                        out2_: T0
                        out3_: bool
                        out4_: int
                        out1_, out2_, out3_, out4_ = (d_104_v79_).m1(d_28_globalState_)
                        d_123_v96_ = out1_
                        d_124_v97_ = out2_
                        d_125_v98_ = out3_
                        d_126_v99_ = out4_
                        d_19_v3_ = d_19_v3_
                    pass
            pass
        d_127_v100_: _dafny.Array
        nw30_ = _dafny.Array(None, 9)
        d_127_v100_ = nw30_
        d_128_v101_: _dafny.Seq
        d_128_v101_ = _dafny.SeqWithoutIsStrInference([d_127_v100_])
        d_129_v102_: _dafny.Seq
        d_129_v102_ = _dafny.SeqWithoutIsStrInference([(d_128_v101_)[default__.safeIndex((d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))], len(d_128_v101_))]])
        d_130_v103_: D3
        d_130_v103_ = D3_DC10((d_87_v62_).f25, d_101_v76_)
        pat_let_tv0_ = d_101_v76_
        d_131_v104_: _dafny.Seq
        def iife10_(_pat_let0_0):
            def iife11_(d_132_dt__update__tmp_h0_):
                def iife12_(_pat_let1_0):
                    def iife13_(d_133_dt__update_hcf17_h0_):
                        return D3_DC10((d_132_dt__update__tmp_h0_).cf16, d_133_dt__update_hcf17_h0_)
                    return iife13_(_pat_let1_0)
                return iife12_(pat_let_tv0_)
            return iife11_(_pat_let0_0)
        d_131_v104_ = _dafny.SeqWithoutIsStrInference([iife10_(D3_DC10((d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))], d_101_v76_)), d_130_v103_])
        d_134_v105_: _dafny.Seq
        d_134_v105_ = _dafny.SeqWithoutIsStrInference([d_87_v62_, d_87_v62_, d_87_v62_])
        d_135_v106_: _dafny.Set
        d_135_v106_ = _dafny.Set({d_134_v105_})
        index17_ = default__.safeIndex(621, (d_19_v3_).length(0))
        index18_ = default__.safeIndex(328, (d_91_v66_).length(0))
        index19_ = default__.safeIndex(621, (d_19_v3_).length(0))
        rhs34_ = ((d_95_v70_).f23) > (len(d_131_v104_))
        rhs35_ = len((d_135_v106_ if (d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))] else d_135_v106_))
        rhs36_ = d_16_v0_
        rhs37_ = d_20_v4_
        rhs38_ = (d_129_v102_) + (d_128_v101_)
        lhs25_ = d_19_v3_
        lhs26_ = default__.safeIndex(621, (d_19_v3_).length(0))
        lhs27_ = d_91_v66_
        lhs28_ = default__.safeIndex(328, (d_91_v66_).length(0))
        lhs29_ = d_19_v3_
        lhs30_ = default__.safeIndex(621, (d_19_v3_).length(0))
        lhs31_ = d_28_globalState_
        lhs25_[lhs26_] = rhs34_
        lhs27_[lhs28_] = rhs35_
        lhs29_[lhs30_] = rhs36_
        lhs31_.f8 = rhs37_
        d_128_v101_ = rhs38_
        d_136_v107_: str
        d_136_v107_ = _dafny.CodePoint('d')
        d_137_v108_: _dafny.Map
        d_137_v108_ = _dafny.Map({(d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))]: d_136_v107_})
        d_138_v109_: _dafny.Map
        d_138_v109_ = _dafny.Map({d_137_v108_: d_16_v0_})
        (d_104_v79_).m8(not(((d_87_v62_).f25 if (d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))] else not(d_16_v0_))), (d_95_v70_).fm2(d_28_globalState_), d_138_v109_, False, d_28_globalState_)
        if d_16_v0_:
            if (d_136_v107_) not in (d_23_v7_):
                d_139_v110_: bool
                d_140_v111_: T0
                d_141_v112_: bool
                d_142_v113_: int
                out5_: bool
                out6_: T0
                out7_: bool
                out8_: int
                out5_, out6_, out7_, out8_ = (d_87_v62_).m1(d_28_globalState_)
                d_139_v110_ = out5_
                d_140_v111_ = out6_
                d_141_v112_ = out7_
                d_142_v113_ = out8_
                d_143_v114_: C5
                nw31_ = C5()
                nw31_.ctor__((d_23_v7_) == ((default__.fm25((d_87_v62_).f25, (d_87_v62_).f25, d_136_v107_, d_28_globalState_)).set(default__.safeIndex((d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))], len(default__.fm25((d_87_v62_).f25, (d_87_v62_).f25, d_136_v107_, d_28_globalState_))), _dafny.CodePoint('t'))), d_139_v110_)
                d_143_v114_ = nw31_
                d_144_v115_: C0
                nw32_ = C0()
                nw32_.ctor__(d_94_v69_, (d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))], (d_95_v70_).f23)
                d_144_v115_ = nw32_
                d_144_v115_ = d_144_v115_
                d_145_v116_: _dafny.Seq
                d_145_v116_ = _dafny.SeqWithoutIsStrInference([d_91_v66_])
                d_91_v66_ = (d_145_v116_)[default__.safeIndex((d_20_v4_) * ((d_29_v9_)[default__.safeIndex((d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))], len(d_29_v9_))]), len(d_145_v116_))]
                d_146_v117_: _dafny.Seq
                d_146_v117_ = _dafny.SeqWithoutIsStrInference([d_134_v105_])
                d_147_v118_: _dafny.MultiSet
                d_147_v118_ = _dafny.MultiSet([(d_140_v111_).f23, d_144_v115_.f31])
                d_148_v119_: _dafny.Map
                d_148_v119_ = _dafny.Map({default__.safeModuloInt(d_142_v113_, d_20_v4_): (d_146_v117_)[default__.safeIndex((d_147_v118_).cardinality, len(d_146_v117_))]})
                d_148_v119_ = (d_148_v119_).set(d_20_v4_, (d_134_v105_).set(default__.safeIndex((d_140_v111_).f23, len(d_134_v105_)), d_87_v62_))
            elif True:
                d_23_v7_ = ((d_23_v7_) + (d_23_v7_)) + (d_23_v7_)
                index20_ = default__.safeIndex(621, (d_19_v3_).length(0))
                (d_19_v3_)[index20_] = (d_87_v62_).f25
                d_149_v120_: C0
                nw33_ = C0()
                nw33_.ctor__(d_94_v69_, (d_95_v70_).f23, (d_95_v70_).f23)
                d_149_v120_ = nw33_
                rhs39_ = default__.safeDivisionInt(default__.safeDivisionInt((d_95_v70_).f23, (d_95_v70_).f23), (d_95_v70_).f23)
                rhs40_ = d_149_v120_
                lhs32_ = d_28_globalState_
                lhs32_.f8 = rhs39_
                d_149_v120_ = rhs40_
                d_150_v121_: _dafny.Set
                d_150_v121_ = _dafny.Set({not(d_87_v62_.f24), ((d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))] if d_87_v62_.f24 else not(not(not((d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))]))))})
                d_150_v121_ = d_150_v121_
                d_151_v122_: _dafny.MultiSet
                d_151_v122_ = _dafny.MultiSet([d_23_v7_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "guaqmlf"))])
                d_16_v0_ = not((((d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))]) + (len(d_23_v7_))) == (((d_151_v122_)[d_23_v7_] if (d_23_v7_) in (d_151_v122_) else (d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))])))
            d_104_v79_ = d_104_v79_
            d_152_v123_: _dafny.Map
            d_152_v123_ = _dafny.Map({D3_DC10((d_87_v62_).f25, d_101_v76_): d_17_v1_})
            d_152_v123_ = ((d_152_v123_) | (d_152_v123_) if d_87_v62_.f24 else d_152_v123_)
            d_153_v124_: _dafny.Set
            d_153_v124_ = _dafny.Set({(d_87_v62_).f25, (d_87_v62_).f25})
            index21_ = default__.safeIndex(328, (d_91_v66_).length(0))
            (d_91_v66_)[index21_] = (0) - (len(d_153_v124_))
            d_16_v0_ = (d_153_v124_).ispropersubset(d_153_v124_)
        elif True:
            (d_28_globalState_).f5 = default__.safeModuloInt((0) - ((d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))]), len(d_23_v7_))
            d_154_v125_: _dafny.Array
            nw34_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_154_v125_ = nw34_
            index22_ = default__.safeIndex(511, (d_154_v125_).length(0))
            (d_154_v125_)[index22_] = d_19_v3_
            d_155_v126_: _dafny.Seq
            d_155_v126_ = _dafny.SeqWithoutIsStrInference([d_19_v3_])
            index23_ = default__.safeIndex(511, (d_154_v125_).length(0))
            (d_154_v125_)[index23_] = ((d_19_v3_ if d_16_v0_ else d_19_v3_) if True else (d_155_v126_)[default__.safeIndex((d_95_v70_).f23, len(d_155_v126_))])
            d_156_v127_: C7
            nw35_ = C7()
            nw35_.ctor__()
            d_156_v127_ = nw35_
            d_157_v128_: _dafny.Seq
            d_157_v128_ = _dafny.SeqWithoutIsStrInference([d_156_v127_])
            d_156_v127_ = (d_157_v128_)[default__.safeIndex((d_91_v66_)[default__.safeIndex(328, (d_91_v66_).length(0))], len(d_157_v128_))]
            d_158_v129_: _dafny.Map
            d_158_v129_ = _dafny.Map({(d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))]: d_94_v69_})
            d_159_v130_: C0
            nw36_ = C0()
            nw36_.ctor__(((d_158_v129_)[(d_87_v62_).f25] if ((d_87_v62_).f25) in (d_158_v129_) else d_94_v69_), (d_95_v70_).fm2(d_28_globalState_), (d_95_v70_).f23)
            d_159_v130_ = nw36_
            d_159_v130_ = d_159_v130_
            d_160_v131_: _dafny.Seq
            d_160_v131_ = _dafny.SeqWithoutIsStrInference([d_91_v66_, d_91_v66_])
            index24_ = default__.safeIndex(621, (d_19_v3_).length(0))
            rhs41_ = (default__.safeModuloInt((d_95_v70_).f23, d_159_v130_.f31)) + ((d_95_v70_).f23)
            rhs42_ = (d_19_v3_)[default__.safeIndex(621, (d_19_v3_).length(0))]
            rhs43_ = ((d_160_v131_) + (_dafny.SeqWithoutIsStrInference([d_91_v66_]))) <= (d_160_v131_)
            lhs33_ = d_28_globalState_
            lhs34_ = d_19_v3_
            lhs35_ = default__.safeIndex(621, (d_19_v3_).length(0))
            lhs36_ = d_28_globalState_
            lhs33_.f8 = rhs41_
            lhs34_[lhs35_] = rhs42_
            lhs36_.f9 = rhs43_
        _dafny.print(_dafny.string_of(d_16_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v1_) == (_dafny.Map({True: 313}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_v2_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v3_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v3_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_20_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_21_v5_) == (_dafny.Map({333: 333}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_22_v6_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_23_v7_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_24_v8_)[0]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_24_v8_)[1]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_24_v8_)[2]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_24_v8_)[3]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_24_v8_)[4]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_24_v8_)[5]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_24_v8_)[6]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_24_v8_)[7]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_24_v8_)[8]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_24_v8_)[9]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_24_v8_)[10]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_24_v8_)[11]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_28_globalState_.f0) == (_dafny.Map({True: 313}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_28_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_28_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_28_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_28_globalState_.f4) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_28_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_28_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_28_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_28_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_28_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_28_globalState_).f10)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_28_globalState_).f11).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_28_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_28_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_28_globalState_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_28_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_28_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_28_globalState_.f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_28_globalState_).f18)[0]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_28_globalState_).f18)[1]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_28_globalState_).f18)[2]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_28_globalState_).f18)[3]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_28_globalState_).f18)[4]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_28_globalState_).f18)[5]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_28_globalState_).f18)[6]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_28_globalState_).f18)[7]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_28_globalState_).f18)[8]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_28_globalState_).f18)[9]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_28_globalState_).f18)[10]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_28_globalState_).f18)[11]) == (_dafny.MultiSet([333, -333, 2, 837]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_28_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_28_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_28_globalState_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_28_globalState_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_29_v9_) == (_dafny.SeqWithoutIsStrInference([4]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_87_v62_.f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v62_).f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_89_v64_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_90_v65_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_v66_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_94_v69_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_95_v70_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v71_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_97_v72_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v73_) == (_dafny.Map({333: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v74_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v75_) == (_dafny.Map({1: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hllu"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v76_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v77_) == (_dafny.Map({_dafny.Map({True: True}): 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v78_) == (_dafny.Map({_dafny.Map({_dafny.Map({True: True}): 1}): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_105_i5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_128_v101_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_129_v102_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_v103_).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_130_v103_).cf17) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v104_) == (_dafny.SeqWithoutIsStrInference([D3_DC10(False, _dafny.Map({True: True})), D3_DC10(True, _dafny.Map({True: True}))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_134_v105_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_135_v106_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_136_v107_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v108_) == (_dafny.Map({True: _dafny.CodePoint('d')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v109_) == (_dafny.Map({_dafny.Map({True: _dafny.CodePoint('d')}): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, int(0))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(False, False)
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

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

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
        return lambda: D3_DC9(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC9(D3, NamedTuple('DC9', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({self.cf14.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC12(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC12(D4, NamedTuple('DC12', [('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14(int(0), False, _dafny.Set({}))
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

class D5_DC14(D5, NamedTuple('DC14', [('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf25', Any), ('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf21 == __o.cf21
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
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC17(D6, NamedTuple('DC17', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC20(False, _dafny.Map({}), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC20(D7, NamedTuple('DC20', [('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23(int(0), _dafny.Set({}), int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC23(D8, NamedTuple('DC23', [('cf40', Any), ('cf41', Any), ('cf42', Any), ('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf45', Any), ('cf46', Any), ('cf47', Any), ('cf48', Any), ('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC26()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)

class D9_DC26(D9, NamedTuple('DC26', [])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26)
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC29(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)

class D10_DC29(D10, NamedTuple('DC29', [('cf53', Any), ('cf54', Any), ('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC31(_dafny.Array(None, 0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)

class D11_DC31(D11, NamedTuple('DC31', [('cf57', Any), ('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC33(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)

class D12_DC33(D12, NamedTuple('DC33', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)

class D13_DC34(D13, NamedTuple('DC34', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D14_DC35)

class D14_DC35(D14, NamedTuple('DC35', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC35({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC35) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC37(False, _dafny.Map({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D15_DC37)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D15_DC38)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D15_DC36)

class D15_DC37(D15, NamedTuple('DC37', [('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC37({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC37) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC38(D15, NamedTuple('DC38', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC38({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC38) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC36(D15, NamedTuple('DC36', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC36({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC36) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC40(int(0), _dafny.Array(None, 0), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D16_DC40)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D16_DC39)

class D16_DC40(D16, NamedTuple('DC40', [('cf70', Any), ('cf71', Any), ('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC40({_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC40) and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC39(D16, NamedTuple('DC39', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC39({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC39) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC42()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D17_DC42)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D17_DC41)

class D17_DC42(D17, NamedTuple('DC42', [])):
    def __dafnystr__(self) -> str:
        return f'D17.DC42'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC42)
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC41(D17, NamedTuple('DC41', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC41({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC41) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D18_DC43)

class D18_DC43(D18, NamedTuple('DC43', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC43({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC43) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m0(self, p0, p1, globalState):
        pass


class T1:
    pass
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    def m1(self, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f0: _dafny.Map = _dafny.Map({})
        self.f1: bool = False
        self.f4: _dafny.Seq = _dafny.Seq({})
        self.f5: int = int(0)
        self.f6: bool = False
        self.f8: int = int(0)
        self.f9: bool = False
        self.f14: int = int(0)
        self.f15: bool = False
        self.f17: int = int(0)
        self._f2: int = int(0)
        self._f3: bool = False
        self._f7: int = int(0)
        self._f10: _dafny.Map = _dafny.Map({})
        self._f11: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f12: int = int(0)
        self._f13: int = int(0)
        self._f16: int = int(0)
        self._f18: _dafny.Array = _dafny.Array(None, 0)
        self._f19: bool = False
        self._f20: bool = False
        self._f21: bool = False
        self._f22: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22):
        (self).f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self).f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f14 = f14
        (self).f15 = f15
        (self)._f16 = f16
        (self).f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21
        (self)._f22 = f22

    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3
    @property
    def f7(self):
        return self._f7
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
    def f16(self):
        return self._f16
    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    @property
    def f22(self):
        return self._f22

class C0(T0):
    def  __init__(self):
        self._f23: int = int(0)
        self.f31: int = int(0)
        self._f30: _dafny.Set = _dafny.Set({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f23(self):
        return self._f23
    def ctor__(self, f30, f31, f23):
        (self)._f30 = f30
        (self).f31 = f31
        (self)._f23 = f23

    def fm2(self, globalState):
        return (867) - (self.f31)

    def fm3(self, p0, p1, p2, p3, globalState):
        source1_ = D0_DC2(D0_DC2(D0_DC1(True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwdbhn"))))))
        if source1_.is_DC1:
            d_161___mcc_h0_ = source1_.cf1
            d_162___mcc_h1_ = source1_.cf2
            d_163_cf2_ = d_162___mcc_h1_
            d_164_cf1_ = d_161___mcc_h0_
            return not(d_164_cf1_)
        elif source1_.is_DC0:
            d_165___mcc_h2_ = source1_.cf0
            d_166_cf0_ = d_165___mcc_h2_
            return d_166_cf0_
        elif True:
            d_167___mcc_h3_ = source1_.cf3
            d_168_cf3_ = d_167___mcc_h3_
            return False

    def m0(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: bool = False
        r2 = p0
        d_169_v0_: str
        d_169_v0_ = _dafny.CodePoint('h')
        d_170_v1_: _dafny.Seq
        d_170_v1_ = _dafny.SeqWithoutIsStrInference([d_169_v0_, d_169_v0_])
        d_170_v1_ = d_170_v1_
        d_171_v2_: _dafny.Array
        nw37_ = _dafny.Array(_dafny.Array(None, 0), 6)
        d_171_v2_ = nw37_
        d_172_v3_: _dafny.Array
        nw38_ = _dafny.Array(False, 10)
        d_172_v3_ = nw38_
        d_173_v4_: _dafny.Array
        d_173_v4_ = d_172_v3_
        index25_ = default__.safeIndex(498, (d_171_v2_).length(0))
        (d_171_v2_)[index25_] = d_172_v3_
        index26_ = default__.safeIndex(498, (d_171_v2_).length(0))
        (d_171_v2_)[index26_] = d_173_v4_
        d_174_v5_: D0
        d_174_v5_ = D0_DC2(D0_DC0(p1))
        d_175_v6_: _dafny.Set
        d_175_v6_ = _dafny.Set({default__.fm24((0) - ((self).f23), p0, p0, globalState), d_174_v5_})
        (globalState).f1 = ((default__.fm23(self.f31, globalState)) != (d_175_v6_) if not(not(p1)) else True)
        index27_ = default__.safeIndex(21, (d_172_v3_).length(0))
        (d_172_v3_)[index27_] = p1
        index28_ = default__.safeIndex(21, (d_172_v3_).length(0))
        (d_172_v3_)[index28_] = p1
        d_176_v7_: _dafny.MultiSet
        d_176_v7_ = _dafny.MultiSet([p1, p1])
        d_177_v8_: _dafny.Seq
        d_177_v8_ = _dafny.SeqWithoutIsStrInference([44, self.f31, (self).f23, self.f31])
        d_178_v9_: _dafny.Map
        d_178_v9_ = _dafny.Map({(d_172_v3_)[default__.safeIndex(21, (d_172_v3_).length(0))]: d_170_v1_})
        d_179_v10_: _dafny.Map
        d_179_v10_ = _dafny.Map({(self).fm3(d_170_v1_, p0, self.f31, d_177_v8_, globalState): (self).f23})
        d_180_v11_: _dafny.Seq
        d_180_v11_ = _dafny.SeqWithoutIsStrInference([p0])
        d_181_v12_: D4
        d_181_v12_ = D4_DC12((self).f23, self.f31)
        d_182_v13_: _dafny.Array
        nw39_ = _dafny.Array(None, 26)
        nw39_[int(0)] = (self).f23
        nw39_[int(1)] = (0) - ((0) - (self.f31))
        nw39_[int(2)] = self.f31
        nw39_[int(3)] = self.f31
        nw39_[int(4)] = (self).f23
        nw39_[int(5)] = (self).f23
        nw39_[int(6)] = (0) - (len(d_177_v8_))
        nw39_[int(7)] = len(d_178_v9_)
        nw39_[int(8)] = (self).f23
        nw39_[int(9)] = (d_177_v8_)[default__.safeIndex((self).f23, len(d_177_v8_))]
        nw39_[int(10)] = (self).f23
        nw39_[int(11)] = (self).f23
        nw39_[int(12)] = 656
        nw39_[int(13)] = (self).f23
        nw39_[int(14)] = ((d_179_v10_)[p0] if (p0) in (d_179_v10_) else self.f31)
        nw39_[int(15)] = default__.fm1(len(d_170_v1_), p0, globalState)
        nw39_[int(16)] = 678
        nw39_[int(17)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ft")))
        nw39_[int(18)] = self.f31
        nw39_[int(19)] = (self).f23
        nw39_[int(20)] = (0) - (self.f31)
        nw39_[int(21)] = self.f31
        nw39_[int(22)] = self.f31
        nw39_[int(23)] = len(d_180_v11_)
        nw39_[int(24)] = 970
        nw39_[int(25)] = len(_dafny.Map({(d_181_v12_).cf19: (d_176_v7_).cardinality}))
        d_182_v13_ = nw39_
        d_183_v14_: _dafny.Seq
        d_183_v14_ = _dafny.SeqWithoutIsStrInference([d_182_v13_, d_182_v13_])
        def iife14_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(61, 671):
                d_184_v15_: int = compr_10_
                if ((61) <= (d_184_v15_)) and ((d_184_v15_) < (671)):
                    coll10_[(d_184_v15_) - (758)] = _dafny.MultiSet([d_169_v0_, d_169_v0_])
            return _dafny.Map(coll10_)
        rhs44_ = (d_183_v14_)[default__.safeIndex(self.f31, len(d_183_v14_))]
        rhs45_ = (d_176_v7_).intersection(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([p1])) + (_dafny.SeqWithoutIsStrInference([(d_172_v3_)[default__.safeIndex(21, (d_172_v3_).length(0))]]))))
        rhs46_ = (len(iife14_()
        )) - ((self.f31) + ((self).f23))
        rhs47_ = default__.safeModuloInt(len((d_179_v10_ if p1 else d_179_v10_)), (self.f31) * (self.f31))
        rhs48_ = (self).f23
        lhs37_ = globalState
        lhs38_ = globalState
        lhs39_ = globalState
        r1 = rhs44_
        d_176_v7_ = rhs45_
        lhs37_.f17 = rhs46_
        lhs38_.f14 = rhs47_
        lhs39_.f14 = rhs48_
        r0 = ((0) - (self.f31)) + ((self).f23)
        r1 = d_182_v13_
        d_185_v16_: _dafny.Map
        d_185_v16_ = _dafny.Map({self.f31: 151})
        d_186_v17_: _dafny.Map
        d_186_v17_ = _dafny.Map({d_172_v3_: (self).fm3(d_170_v1_, True, (self).f23, _dafny.SeqWithoutIsStrInference([len(d_185_v16_), self.f31]), globalState)})
        r2 = (d_172_v3_) not in ((d_186_v17_) | (d_186_v17_))
        return r0, r1, r2

    def m10(self, p0, p1, p2, p3, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        r3: bool = False
        r2 = p2
        d_187_v0_: _dafny.Seq
        d_187_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mevk"))
        d_187_v0_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_188_i0_ in range(default__.abs(717))])) + ((d_187_v0_) + (d_187_v0_))
        d_189_v1_: str
        d_189_v1_ = _dafny.CodePoint('w')
        d_190_v2_: _dafny.Seq
        d_190_v2_ = _dafny.SeqWithoutIsStrInference([p0])
        d_191_v3_: _dafny.Map
        d_191_v3_ = _dafny.Map({p3: p0})
        d_192_v4_: _dafny.Map
        d_192_v4_ = _dafny.Map({p0: p0})
        d_193_v5_: D3
        d_193_v5_ = D3_DC10(p0, d_192_v4_)
        d_194_v6_: _dafny.Array
        nw40_ = _dafny.Array(None, 26)
        nw40_[int(0)] = (d_187_v0_) + (d_187_v0_)
        nw40_[int(1)] = d_187_v0_
        nw40_[int(2)] = (d_187_v0_).set(default__.safeIndex(self.f31, len(d_187_v0_)), d_189_v1_)
        nw40_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkd"))
        nw40_[int(4)] = d_187_v0_
        nw40_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aai"))
        nw40_[int(6)] = d_187_v0_
        nw40_[int(7)] = d_187_v0_
        nw40_[int(8)] = d_187_v0_
        nw40_[int(9)] = default__.fm25(p0, p0, d_189_v1_, globalState)
        nw40_[int(10)] = d_187_v0_
        nw40_[int(11)] = _dafny.SeqWithoutIsStrInference([d_189_v1_ for d_195_i1_ in range(default__.abs(58))])
        nw40_[int(12)] = (d_187_v0_) + (d_187_v0_)
        nw40_[int(13)] = d_187_v0_
        nw40_[int(14)] = d_187_v0_
        nw40_[int(15)] = (_dafny.SeqWithoutIsStrInference([d_189_v1_ for d_196_i2_ in range(default__.abs(285))])) + (d_187_v0_)
        nw40_[int(16)] = d_187_v0_
        nw40_[int(17)] = default__.fm25((d_190_v2_)[default__.safeIndex((0) - (p2), len(d_190_v2_))], p0, default__.fm26(len(d_191_v3_), p0, globalState), globalState)
        nw40_[int(18)] = d_187_v0_
        nw40_[int(19)] = d_187_v0_
        nw40_[int(20)] = default__.fm25(((d_192_v4_)[p0] if (p0) in (d_192_v4_) else not(p0)), not((d_193_v5_).cf16), d_189_v1_, globalState)
        nw40_[int(21)] = d_187_v0_
        nw40_[int(22)] = d_187_v0_
        nw40_[int(23)] = d_187_v0_
        nw40_[int(24)] = d_187_v0_
        nw40_[int(25)] = d_187_v0_
        d_194_v6_ = nw40_
        rhs49_ = not(((self).f23) != (p2))
        rhs50_ = d_194_v6_
        rhs51_ = p0
        lhs40_ = globalState
        lhs40_.f1 = rhs49_
        d_194_v6_ = rhs50_
        r3 = rhs51_
        d_197_v7_: _dafny.Map
        d_197_v7_ = _dafny.Map({(d_187_v0_) != (d_187_v0_): d_192_v4_})
        d_198_v8_: _dafny.MultiSet
        d_198_v8_ = _dafny.MultiSet([p0, p0, p0])
        d_199_v9_: _dafny.MultiSet
        d_199_v9_ = _dafny.MultiSet([default__.fm1(p2, p0, globalState)])
        d_197_v7_ = (default__.fm27(self.f31, d_198_v8_, (d_199_v9_).cardinality, p2, globalState)) | (d_197_v7_)
        d_200_v10_: _dafny.Map
        d_200_v10_ = _dafny.Map({p0: (self).f23})
        d_201_v11_: _dafny.Seq
        d_201_v11_ = _dafny.SeqWithoutIsStrInference([self.f31, p3])
        hi1_ = len(d_201_v11_)
        for d_202_i3_ in range((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hk")))) + (((d_200_v10_)[((d_192_v4_)[False] if (False) in (d_192_v4_) else not(p0))] if (((d_192_v4_)[False] if (False) in (d_192_v4_) else not(p0))) in (d_200_v10_) else (self).f23)), hi1_):
            d_203_v12_: _dafny.Array
            nw41_ = _dafny.Array(False, 5)
            d_203_v12_ = nw41_
            index29_ = default__.safeIndex(869, (d_203_v12_).length(0))
            (d_203_v12_)[index29_] = p0
            d_204_v13_: _dafny.Array
            nw42_ = _dafny.Array(None, 6)
            nw42_[int(0)] = d_189_v1_
            nw42_[int(1)] = d_189_v1_
            nw42_[int(2)] = (d_189_v1_ if p0 else d_189_v1_)
            nw42_[int(3)] = d_189_v1_
            nw42_[int(4)] = d_189_v1_
            nw42_[int(5)] = d_189_v1_
            d_204_v13_ = nw42_
            index30_ = default__.safeIndex(869, (d_203_v12_).length(0))
            rhs52_ = d_190_v2_
            rhs53_ = p0
            rhs54_ = (d_187_v0_) + ((d_187_v0_ if p0 else d_187_v0_))
            rhs55_ = d_204_v13_
            lhs41_ = d_203_v12_
            lhs42_ = default__.safeIndex(869, (d_203_v12_).length(0))
            d_190_v2_ = rhs52_
            lhs41_[lhs42_] = rhs53_
            d_187_v0_ = rhs54_
            d_204_v13_ = rhs55_
            index31_ = default__.safeIndex(869, (d_203_v12_).length(0))
            index32_ = default__.safeIndex(869, (d_203_v12_).length(0))
            rhs56_ = not(((d_190_v2_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahkl"))), len(d_190_v2_))]) not in (d_190_v2_))
            rhs57_ = p0
            rhs58_ = (d_203_v12_)[default__.safeIndex(869, (d_203_v12_).length(0))]
            rhs59_ = (not ((d_203_v12_)[default__.safeIndex(869, (d_203_v12_).length(0))]) or (False)) and ((d_203_v12_)[default__.safeIndex(869, (d_203_v12_).length(0))])
            lhs43_ = globalState
            lhs44_ = d_203_v12_
            lhs45_ = default__.safeIndex(869, (d_203_v12_).length(0))
            lhs46_ = globalState
            lhs47_ = d_203_v12_
            lhs48_ = default__.safeIndex(869, (d_203_v12_).length(0))
            lhs43_.f15 = rhs56_
            lhs44_[lhs45_] = rhs57_
            lhs46_.f1 = rhs58_
            lhs47_[lhs48_] = rhs59_
            (globalState).f8 = (0) - (default__.safeDivisionInt(self.f31, d_202_i3_))
            (globalState).f17 = (self).fm2(globalState)
        (self).f31 = default__.safeDivisionInt((self.f31) - ((0) - ((self).f23)), (d_198_v8_).cardinality)
        r0 = d_199_v9_
        d_205_v14_: _dafny.Array
        nw43_ = _dafny.Array(None, 3)
        nw43_[int(0)] = (p0) in (_dafny.Map({p0: p0}))
        nw43_[int(1)] = p0
        nw43_[int(2)] = p0
        d_205_v14_ = nw43_
        r1 = d_205_v14_
        d_206_v15_: _dafny.Map
        d_206_v15_ = _dafny.Map({len(d_187_v0_): 497})
        r2 = ((d_199_v9_)[(0) - (p3)] if ((0) - (p3)) in (d_199_v9_) else ((d_206_v15_)[(self).fm2(globalState)] if ((self).fm2(globalState)) in (d_206_v15_) else (0) - (self.f31)))
        r3 = p0
        return r0, r1, r2, r3

    @property
    def f30(self):
        return self._f30

class C1(T1):
    def  __init__(self):
        self._f24: bool = False
        self._f25: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f24, f25):
        (self).f24 = f24
        (self)._f25 = f25

    def fm4(self, p0, p1, p2, globalState):
        return (_dafny.Map({self.f24: _dafny.Map({self.f24: self.f24})})) | ((_dafny.Map({self.f24: _dafny.Map({(self).f25: True})})) | (_dafny.Map({(self).f25: _dafny.Map({self.f24: (self).f25})})))

    def fm5(self, p0, p1, globalState):
        return ((len(_dafny.SeqWithoutIsStrInference([True, (self).f25]))) * (953)) + ((954 if self.f24 else (_dafny.MultiSet([36, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lidbxgdo"))))])).cardinality))

    def m1(self, globalState):
        r0: bool = False
        r1: T0 = None
        r2: bool = False
        r3: int = int(0)
        d_207_v0_: _dafny.Seq
        d_207_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqbcphfi"))
        d_208_v1_: D3
        d_208_v1_ = D3_DC8(d_207_v0_)
        source2_ = d_208_v1_
        if source2_.is_DC9:
            d_209___mcc_h0_ = source2_.cf15
            d_210_cf15_ = d_209___mcc_h0_
            d_211_v2_: _dafny.Array
            nw44_ = _dafny.Array(False, 26)
            d_211_v2_ = nw44_
            index33_ = default__.safeIndex(856, (d_211_v2_).length(0))
            (d_211_v2_)[index33_] = (self).f25
            d_212_v3_: _dafny.Map
            d_212_v3_ = _dafny.Map({(self).f25: self.f24})
            d_213_v4_: _dafny.Seq
            d_213_v4_ = _dafny.SeqWithoutIsStrInference([d_210_cf15_, len(d_212_v3_), d_210_cf15_])
            d_214_v5_: _dafny.Set
            d_214_v5_ = _dafny.Set({self.f24})
            d_215_v6_: str
            d_215_v6_ = _dafny.CodePoint('l')
            index34_ = default__.safeIndex(856, (d_211_v2_).length(0))
            rhs60_ = ((_dafny.SeqWithoutIsStrInference([d_210_cf15_ for d_216_i0_ in range(default__.abs(139))])) + (d_213_v4_)) <= ((_dafny.SeqWithoutIsStrInference([len(d_213_v4_), d_210_cf15_, d_210_cf15_])) + (d_213_v4_))
            rhs61_ = ((d_214_v5_) | (d_214_v5_)).ispropersubset(d_214_v5_)
            rhs62_ = (d_207_v0_).set(default__.safeIndex(d_210_cf15_, len(d_207_v0_)), d_215_v6_)
            rhs63_ = not ((self).f25) or (self.f24)
            rhs64_ = d_207_v0_
            lhs49_ = globalState
            lhs50_ = d_211_v2_
            lhs51_ = default__.safeIndex(856, (d_211_v2_).length(0))
            lhs52_ = globalState
            lhs49_.f6 = rhs60_
            lhs50_[lhs51_] = rhs61_
            d_207_v0_ = rhs62_
            lhs52_.f1 = rhs63_
            d_207_v0_ = rhs64_
            d_217_v7_: _dafny.Array
            nw45_ = _dafny.Array(_dafny.MultiSet({}), 11)
            d_217_v7_ = nw45_
            index35_ = default__.safeIndex(559, (d_217_v7_).length(0))
            (d_217_v7_)[index35_] = _dafny.MultiSet([d_210_cf15_])
            d_218_v8_: _dafny.MultiSet
            d_218_v8_ = _dafny.MultiSet([d_210_cf15_])
            d_219_v9_: _dafny.Map
            d_219_v9_ = _dafny.Map({self.f24: d_210_cf15_})
            index36_ = default__.safeIndex(559, (d_217_v7_).length(0))
            (d_217_v7_)[index36_] = ((d_218_v8_) | (d_218_v8_)) - ((_dafny.MultiSet([len(_dafny.Map({len(d_219_v9_): self.f24}))])) | (d_218_v8_))
            if False:
                (globalState).f8 = d_210_cf15_
                d_220_v10_: _dafny.Array
                nw46_ = _dafny.Array(int(0), 14)
                d_220_v10_ = nw46_
                index37_ = default__.safeIndex(521, (d_220_v10_).length(0))
                (d_220_v10_)[index37_] = d_210_cf15_
                index38_ = default__.safeIndex(521, (d_220_v10_).length(0))
                (d_220_v10_)[index38_] = (0) - ((0) - (d_210_cf15_))
                (globalState).f17 = default__.fm1(((d_220_v10_)[default__.safeIndex(521, (d_220_v10_).length(0))]) - (582), ((self).f25) == ((self).f25), globalState)
                d_221_v11_: _dafny.MultiSet
                d_221_v11_ = _dafny.MultiSet([(d_213_v4_) + (d_213_v4_)])
                d_221_v11_ = d_221_v11_
                d_222_v12_: _dafny.MultiSet
                d_222_v12_ = _dafny.MultiSet([(d_211_v2_)[default__.safeIndex(856, (d_211_v2_).length(0))], self.f24])
                r2 = ((d_222_v12_).intersection(default__.fm22(d_210_cf15_, (d_211_v2_)[default__.safeIndex(856, (d_211_v2_).length(0))], self.f24, (d_220_v10_)[default__.safeIndex(521, (d_220_v10_).length(0))], globalState))).isdisjoint(_dafny.MultiSet([(d_211_v2_)[default__.safeIndex(856, (d_211_v2_).length(0))], self.f24, (self).f25, True]))
            elif True:
                (globalState).f1 = self.f24
                (self).f24 = (default__.safeModuloInt(615, d_210_cf15_)) > ((d_210_cf15_) + (d_210_cf15_))
                r2 = (d_211_v2_)[default__.safeIndex(856, (d_211_v2_).length(0))]
                d_207_v0_ = (_dafny.SeqWithoutIsStrInference([d_215_v6_ for d_223_i1_ in range(default__.abs(517))])) + (d_207_v0_)
                index39_ = default__.safeIndex(856, (d_211_v2_).length(0))
                (d_211_v2_)[index39_] = (False) or ((self).f25)
            (globalState).f5 = (0) - (d_210_cf15_)
        elif source2_.is_DC10:
            d_224___mcc_h1_ = source2_.cf16
            d_225___mcc_h2_ = source2_.cf17
            d_226_cf17_ = d_225___mcc_h2_
            d_227_cf16_ = d_224___mcc_h1_
            d_228_v13_: _dafny.Array
            nw47_ = _dafny.Array(False, 12)
            d_228_v13_ = nw47_
            d_229_v14_: _dafny.Array
            d_229_v14_ = d_228_v13_
            d_229_v14_ = d_229_v14_
            d_230_v16_: _dafny.Array
            def lambda4_(d_231_cf16_):
                def lambda5_(d_232_i2_):
                    def iife15_():
                        coll11_ = _dafny.Set()
                        compr_11_: int
                        for compr_11_ in _dafny.IntegerRange(380, 719):
                            d_233_v15_: int = compr_11_
                            if ((380) <= (d_233_v15_)) and ((d_233_v15_) < (719)):
                                coll11_ = coll11_.union(_dafny.Set([default__.safeDivisionInt(d_233_v15_, len(_dafny.Map({d_231_cf16_: _dafny.MultiSet([self.f24])})))]))
                        return _dafny.Set(coll11_)
                    return iife15_()
                    

                return lambda5_

            init2_ = lambda4_(d_227_cf16_)
            nw48_ = _dafny.Array(None, 9)
            for i0_2_ in range(nw48_.length(0)):
                nw48_[i0_2_] = init2_(i0_2_)
            d_230_v16_ = nw48_
            d_234_v17_: int
            d_234_v17_ = 721
            d_235_v18_: _dafny.Map
            d_235_v18_ = _dafny.Map({d_227_cf16_: d_234_v17_})
            d_236_v19_: _dafny.Set
            d_236_v19_ = _dafny.Set({len(d_235_v18_), (0) - (d_234_v17_), len(_dafny.Set({d_227_cf16_})), 63})
            index40_ = default__.safeIndex(595, (d_230_v16_).length(0))
            (d_230_v16_)[index40_] = d_236_v19_
            index41_ = default__.safeIndex(595, (d_230_v16_).length(0))
            (d_230_v16_)[index41_] = d_236_v19_
            d_237_v20_: _dafny.Array
            nw49_ = _dafny.Array(_dafny.MultiSet({}), 22)
            d_237_v20_ = nw49_
            d_238_v21_: _dafny.Seq
            d_238_v21_ = _dafny.SeqWithoutIsStrInference([d_237_v20_])
            d_239_v23_: _dafny.Seq
            def iife16_():
                coll12_ = _dafny.Map()
                compr_12_: int
                for compr_12_ in _dafny.IntegerRange(628, 146):
                    d_240_v22_: int = compr_12_
                    if ((628) <= (d_240_v22_)) and ((d_240_v22_) < (146)):
                        coll12_[default__.safeDivisionInt(d_240_v22_, len(_dafny.SeqWithoutIsStrInference([d_227_cf16_])))] = d_227_cf16_
                return _dafny.Map(coll12_)
            d_239_v23_ = _dafny.SeqWithoutIsStrInference([d_237_v20_, d_237_v20_, (d_238_v21_)[default__.safeIndex(len(iife16_()
            ), len(d_238_v21_))]])
            d_237_v20_ = (d_239_v23_)[default__.safeIndex((d_234_v17_) + (155), len(d_239_v23_))]
            d_241_v24_: _dafny.MultiSet
            d_241_v24_ = _dafny.MultiSet([(self).f25])
            rhs65_ = d_234_v17_
            rhs66_ = False
            rhs67_ = (_dafny.MultiSet([d_227_cf16_, d_227_cf16_])).issubset((D7_DC19(d_241_v24_)).cf34)
            lhs53_ = globalState
            lhs54_ = globalState
            lhs53_.f5 = rhs65_
            lhs54_.f15 = rhs66_
            r2 = rhs67_
        elif True:
            d_242___mcc_h3_ = source2_.cf14
            d_243_cf14_ = d_242___mcc_h3_
            d_244_v25_: str
            d_244_v25_ = _dafny.CodePoint('l')
            d_244_v25_ = d_244_v25_
            d_245_v26_: _dafny.Array
            def lambda6_(d_246_i3_):
                return (_dafny.SeqWithoutIsStrInference([552])) + (_dafny.SeqWithoutIsStrInference([-777, len(_dafny.SeqWithoutIsStrInference([self.f24, self.f24]))]))

            init3_ = lambda6_
            nw50_ = _dafny.Array(None, 10)
            for i0_3_ in range(nw50_.length(0)):
                nw50_[i0_3_] = init3_(i0_3_)
            d_245_v26_ = nw50_
            index42_ = default__.safeIndex(530, (d_245_v26_).length(0))
            (d_245_v26_)[index42_] = _dafny.SeqWithoutIsStrInference([-996 for d_247_i4_ in range(default__.abs(396))])
            d_248_v27_: int
            d_248_v27_ = 40
            d_249_v28_: _dafny.Seq
            d_249_v28_ = _dafny.SeqWithoutIsStrInference([d_248_v27_])
            index43_ = default__.safeIndex(530, (d_245_v26_).length(0))
            (d_245_v26_)[index43_] = d_249_v28_
            d_250_v29_: _dafny.Seq
            d_250_v29_ = _dafny.SeqWithoutIsStrInference([self.f24, self.f24, not((self).f25)])
            d_251_v30_: _dafny.MultiSet
            d_251_v30_ = _dafny.MultiSet([self.f24, (self).f25])
            d_252_v31_: _dafny.Set
            d_252_v31_ = _dafny.Set({d_248_v27_})
            (self).m9(d_244_v25_, not ((self).f25) or (not(self.f24)), (d_250_v29_)[default__.safeIndex(d_248_v27_, len(d_250_v29_))], default__.safeModuloInt((d_251_v30_).cardinality, len(d_252_v31_)), globalState)
            d_253_v32_: _dafny.Array
            nw51_ = _dafny.Array(None, 25)
            nw51_[int(0)] = self.f24
            nw51_[int(1)] = True
            nw51_[int(2)] = self.f24
            nw51_[int(3)] = self.f24
            nw51_[int(4)] = (self).f25
            nw51_[int(5)] = self.f24
            nw51_[int(6)] = (self).f25
            nw51_[int(7)] = (self).f25
            nw51_[int(8)] = self.f24
            nw51_[int(9)] = not(self.f24)
            nw51_[int(10)] = not((self).f25)
            nw51_[int(11)] = self.f24
            nw51_[int(12)] = (self).f25
            nw51_[int(13)] = True
            nw51_[int(14)] = (d_250_v29_)[default__.safeIndex(d_248_v27_, len(d_250_v29_))]
            nw51_[int(15)] = self.f24
            nw51_[int(16)] = (self).f25
            nw51_[int(17)] = (self).f25
            nw51_[int(18)] = (self).f25
            nw51_[int(19)] = (self).f25
            nw51_[int(20)] = self.f24
            nw51_[int(21)] = (self).f25
            nw51_[int(22)] = (self).f25
            nw51_[int(23)] = self.f24
            nw51_[int(24)] = (self).f25
            d_253_v32_ = nw51_
            d_254_v33_: _dafny.Set
            d_254_v33_ = _dafny.Set({d_253_v32_})
            d_255_v34_: T0
            nw52_ = C0()
            nw52_.ctor__((d_254_v33_).intersection(d_254_v33_), default__.fm1(d_248_v27_, (self).f25, globalState), d_248_v27_)
            d_255_v34_ = nw52_
            r1 = d_255_v34_
        d_256_v35_: _dafny.Seq
        d_256_v35_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        d_257_v36_: int
        d_257_v36_ = -850
        d_258_v37_: _dafny.Map
        d_258_v37_ = _dafny.Map({(self).f25: (self).f25})
        d_259_v38_: _dafny.Map
        d_259_v38_ = _dafny.Map({self.f24: d_257_v36_})
        d_260_v39_: D1
        d_260_v39_ = D1_DC5((d_256_v35_)[default__.safeIndex(d_257_v36_, len(d_256_v35_))], d_258_v37_, d_257_v36_, d_257_v36_, d_259_v38_)
        d_261_v40_: _dafny.Map
        d_261_v40_ = _dafny.Map({self.f24: d_260_v39_})
        d_262_v41_: _dafny.MultiSet
        d_262_v41_ = _dafny.MultiSet([default__.fm28(d_257_v36_, d_207_v0_, d_257_v36_, globalState), _dafny.Map({self.f24: d_260_v39_}), d_261_v40_])
        d_263_i5_: int
        d_263_i5_ = 0
        with _dafny.label("1"):
            while ((_dafny.MultiSet([d_261_v40_, d_261_v40_])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f25: d_260_v39_})])))).ispropersubset((_dafny.MultiSet([d_261_v40_])) - (d_262_v41_)):
                with _dafny.c_label("1"):
                    if (d_263_i5_) >= (100):
                        raise _dafny.Break("1")
                    d_263_i5_ = (d_263_i5_) + (1)
                    d_264_v42_: _dafny.Array
                    def lambda7_(d_265_i6_):
                        return self.f24

                    init4_ = lambda7_
                    nw53_ = _dafny.Array(None, 29)
                    for i0_4_ in range(nw53_.length(0)):
                        nw53_[i0_4_] = init4_(i0_4_)
                    d_264_v42_ = nw53_
                    d_266_v43_: _dafny.Set
                    d_266_v43_ = _dafny.Set({d_264_v42_, d_264_v42_})
                    d_267_v44_: _dafny.Map
                    d_267_v44_ = _dafny.Map({d_257_v36_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvrqm"))})
                    d_268_v45_: C0
                    nw54_ = C0()
                    nw54_.ctor__(d_266_v43_, d_257_v36_, default__.safeDivisionInt(d_257_v36_, len(d_267_v44_)))
                    d_268_v45_ = nw54_
                    d_269_v46_: _dafny.Seq
                    d_269_v46_ = _dafny.SeqWithoutIsStrInference([d_257_v36_])
                    d_270_v47_: _dafny.Set
                    d_270_v47_ = _dafny.Set({len(d_269_v46_)})
                    d_271_v48_: _dafny.MultiSet
                    d_271_v48_ = _dafny.MultiSet([d_270_v47_])
                    d_272_v49_: _dafny.Map
                    d_272_v49_ = _dafny.Map({(_dafny.MultiSet([_dafny.Set({d_257_v36_})])) - (d_271_v48_): (d_268_v45_.f31) - ((self).fm5(len(_dafny.Map({292: d_264_v42_})), d_268_v45_.f31, globalState))})
                    d_273_v50_: _dafny.Seq
                    d_273_v50_ = _dafny.SeqWithoutIsStrInference([d_270_v47_, d_270_v47_, d_270_v47_])
                    d_274_v51_: _dafny.Map
                    d_274_v51_ = _dafny.Map({not(self.f24): d_264_v42_})
                    d_275_v52_: _dafny.MultiSet
                    d_275_v52_ = _dafny.MultiSet([d_264_v42_, ((d_274_v51_)[self.f24] if (self.f24) in (d_274_v51_) else d_264_v42_), d_264_v42_, d_264_v42_])
                    d_272_v49_ = (d_272_v49_).set(_dafny.MultiSet(d_273_v50_), ((d_275_v52_)[d_264_v42_] if (d_264_v42_) in (d_275_v52_) else (0) - ((self).fm5(d_268_v45_.f31, d_268_v45_.f31, globalState))))
                    r3 = d_268_v45_.f31
                    d_276_v53_: D0
                    d_276_v53_ = D0_DC1(self.f24, d_257_v36_)
                    d_277_v54_: _dafny.Map
                    d_277_v54_ = _dafny.Map({(d_269_v46_).set(default__.safeIndex((d_276_v53_).cf2, len(d_269_v46_)), d_268_v45_.f31): (self).f25})
                    d_278_v55_: _dafny.MultiSet
                    d_278_v55_ = _dafny.MultiSet([default__.fm22(d_268_v45_.f31, (self).f25, not(True), d_268_v45_.f31, globalState), default__.fm22(d_268_v45_.f31, False, (self).f25, len(d_277_v54_), globalState)])
                    d_278_v55_ = (d_278_v55_).intersection((d_278_v55_) | (d_278_v55_))
                    pass
            pass
        d_279_v56_: _dafny.Array
        nw55_ = _dafny.Array(_dafny.Map({}), 24)
        d_279_v56_ = nw55_
        d_280_v57_: _dafny.Map
        d_280_v57_ = _dafny.Map({d_257_v36_: (self).f25})
        index44_ = default__.safeIndex(621, (d_279_v56_).length(0))
        (d_279_v56_)[index44_] = d_280_v57_
        index45_ = default__.safeIndex(621, (d_279_v56_).length(0))
        (d_279_v56_)[index45_] = d_280_v57_
        if self.f24:
            d_281_v58_: _dafny.Array
            def lambda8_(d_282_i7_):
                return (_dafny.CodePoint('g') if (self).f25 else _dafny.CodePoint('c'))

            init5_ = lambda8_
            nw56_ = _dafny.Array(None, 21)
            for i0_5_ in range(nw56_.length(0)):
                nw56_[i0_5_] = init5_(i0_5_)
            d_281_v58_ = nw56_
            d_283_v59_: str
            d_283_v59_ = _dafny.CodePoint('a')
            index46_ = default__.safeIndex(72, (d_281_v58_).length(0))
            (d_281_v58_)[index46_] = (d_283_v59_ if (self).f25 else d_283_v59_)
            d_284_v60_: _dafny.Seq
            d_284_v60_ = _dafny.SeqWithoutIsStrInference([973])
            d_285_v61_: _dafny.MultiSet
            d_285_v61_ = _dafny.MultiSet([(d_284_v60_) + (d_284_v60_)])
            index47_ = default__.safeIndex(72, (d_281_v58_).length(0))
            index48_ = default__.safeIndex(621, (d_279_v56_).length(0))
            rhs68_ = ((d_285_v61_)[(d_284_v60_) + (d_284_v60_)] if ((d_284_v60_) + (d_284_v60_)) in (d_285_v61_) else d_257_v36_)
            rhs69_ = d_283_v59_
            rhs70_ = (d_279_v56_)[default__.safeIndex(621, (d_279_v56_).length(0))]
            lhs55_ = globalState
            lhs56_ = d_281_v58_
            lhs57_ = default__.safeIndex(72, (d_281_v58_).length(0))
            lhs58_ = d_279_v56_
            lhs59_ = default__.safeIndex(621, (d_279_v56_).length(0))
            lhs55_.f8 = rhs68_
            lhs56_[lhs57_] = rhs69_
            lhs58_[lhs59_] = rhs70_
            d_286_v62_: _dafny.Array
            nw57_ = _dafny.Array(int(0), 21)
            d_286_v62_ = nw57_
            rhs71_ = (len(d_256_v35_)) * ((d_257_v36_) + (d_257_v36_))
            rhs72_ = d_286_v62_
            rhs73_ = (d_256_v35_)[default__.safeIndex(default__.fm1(d_257_v36_, (self).f25, globalState), len(d_256_v35_))]
            lhs60_ = globalState
            lhs61_ = globalState
            lhs60_.f14 = rhs71_
            d_286_v62_ = rhs72_
            lhs61_.f9 = rhs73_
            (globalState).f8 = d_257_v36_
            d_287_v63_: D1
            d_287_v63_ = D1_DC4((self).f25, self.f24)
            d_288_v64_: _dafny.Array
            nw58_ = _dafny.Array(None, 4)
            nw58_[int(0)] = (d_287_v63_).cf5
            nw58_[int(1)] = self.f24
            nw58_[int(2)] = self.f24
            nw58_[int(3)] = not ((self).f25) or ((self).f25)
            d_288_v64_ = nw58_
            index49_ = default__.safeIndex(930, (d_288_v64_).length(0))
            (d_288_v64_)[index49_] = self.f24
            d_289_v65_: _dafny.Set
            d_289_v65_ = _dafny.Set({self.f24, (self).f25})
            index50_ = default__.safeIndex(930, (d_288_v64_).length(0))
            (d_288_v64_)[index50_] = (d_289_v65_).issubset(d_289_v65_)
            (globalState).f15 = (self).f25
        elif True:
            d_290_v66_: str
            d_290_v66_ = _dafny.CodePoint('p')
            d_290_v66_ = d_290_v66_
            d_291_v67_: _dafny.Array
            nw59_ = _dafny.Array(None, 25)
            nw59_[int(0)] = self.f24
            nw59_[int(1)] = not(False)
            nw59_[int(2)] = True
            nw59_[int(3)] = (self).f25
            nw59_[int(4)] = (self).f25
            nw59_[int(5)] = True
            nw59_[int(6)] = True
            nw59_[int(7)] = self.f24
            nw59_[int(8)] = True
            nw59_[int(9)] = (self).f25
            nw59_[int(10)] = False
            nw59_[int(11)] = (self).f25
            nw59_[int(12)] = default__.fm29(d_257_v36_, d_257_v36_, (self).f25, globalState)
            nw59_[int(13)] = self.f24
            nw59_[int(14)] = (self).f25
            nw59_[int(15)] = (self).f25
            nw59_[int(16)] = self.f24
            nw59_[int(17)] = self.f24
            nw59_[int(18)] = (self).f25
            nw59_[int(19)] = self.f24
            nw59_[int(20)] = (self).f25
            nw59_[int(21)] = (self).f25
            nw59_[int(22)] = (self).f25
            nw59_[int(23)] = (self).f25
            nw59_[int(24)] = not(self.f24)
            d_291_v67_ = nw59_
            d_292_v68_: _dafny.Set
            d_292_v68_ = _dafny.Set({d_291_v67_, d_291_v67_})
            d_293_v69_: C0
            nw60_ = C0()
            nw60_.ctor__(d_292_v68_, len(d_256_v35_), (d_257_v36_) * (d_257_v36_))
            d_293_v69_ = nw60_
            d_294_v70_: _dafny.Array
            def lambda9_(d_295_i8_):
                return D0_DC2(D0_DC0((self).f25))

            init6_ = lambda9_
            nw61_ = _dafny.Array(None, 29)
            for i0_6_ in range(nw61_.length(0)):
                nw61_[i0_6_] = init6_(i0_6_)
            d_294_v70_ = nw61_
            d_296_v71_: _dafny.Map
            d_296_v71_ = _dafny.Map({d_294_v70_: d_290_v66_})
            d_297_v72_: T0
            nw62_ = C0()
            nw62_.ctor__((d_293_v69_).f30, d_257_v36_, d_293_v69_.f31)
            d_297_v72_ = nw62_
            d_298_v73_: D3
            d_298_v73_ = D3_DC9(len(_dafny.Map({d_297_v72_: d_293_v69_.f31})))
            d_299_v74_: _dafny.Map
            d_299_v74_ = _dafny.Map({d_257_v36_: d_207_v0_})
            d_300_v75_: _dafny.MultiSet
            d_300_v75_ = _dafny.MultiSet([d_257_v36_, -199])
            d_301_v76_: _dafny.MultiSet
            d_301_v76_ = _dafny.MultiSet([self.f24])
            d_302_v77_: _dafny.Seq
            d_302_v77_ = _dafny.SeqWithoutIsStrInference([d_257_v36_, d_257_v36_])
            d_303_v79_: _dafny.Set
            d_303_v79_ = _dafny.Set({978})
            d_304_v80_: _dafny.Array
            nw63_ = _dafny.Array(None, 25)
            nw63_[int(0)] = (d_297_v72_).f23
            nw63_[int(1)] = d_293_v69_.f31
            nw63_[int(2)] = 966
            nw63_[int(3)] = d_293_v69_.f31
            nw63_[int(4)] = (d_297_v72_).f23
            nw63_[int(5)] = len(d_256_v35_)
            nw63_[int(6)] = d_257_v36_
            nw63_[int(7)] = d_257_v36_
            nw63_[int(8)] = d_293_v69_.f31
            nw63_[int(9)] = len(d_299_v74_)
            nw63_[int(10)] = d_257_v36_
            nw63_[int(11)] = (d_300_v75_).cardinality
            nw63_[int(12)] = (d_297_v72_).f23
            nw63_[int(13)] = d_257_v36_
            nw63_[int(14)] = ((d_301_v76_)[(self).f25] if ((self).f25) in (d_301_v76_) else len(d_302_v77_))
            nw63_[int(15)] = (d_297_v72_).f23
            nw63_[int(16)] = len(d_302_v77_)
            def iife17_():
                coll13_ = _dafny.Set()
                compr_13_: int
                for compr_13_ in (d_302_v77_).Elements:
                    d_305_v78_: int = compr_13_
                    if (d_305_v78_) in (d_302_v77_):
                        coll13_ = coll13_.union(_dafny.Set([(d_305_v78_) - (421)]))
                return _dafny.Set(coll13_)
            nw63_[int(17)] = len(_dafny.Map({(_dafny.MultiSet([iife17_()
            , d_303_v79_])).cardinality: self.f24}))
            nw63_[int(18)] = (d_302_v77_)[default__.safeIndex(d_257_v36_, len(d_302_v77_))]
            nw63_[int(19)] = d_257_v36_
            nw63_[int(20)] = d_293_v69_.f31
            nw63_[int(21)] = len(d_207_v0_)
            nw63_[int(22)] = d_257_v36_
            nw63_[int(23)] = d_257_v36_
            nw63_[int(24)] = d_257_v36_
            d_304_v80_ = nw63_
            d_306_v81_: D8
            d_306_v81_ = D8_DC22(d_304_v80_)
            d_307_v82_: _dafny.Map
            d_307_v82_ = _dafny.Map({d_298_v73_: (d_306_v81_).cf39})
            d_308_v83_: D5
            d_308_v83_ = D5_DC15((d_297_v72_).f23, d_290_v66_, _dafny.Map({d_294_v70_: d_290_v66_}), d_307_v82_)
            d_309_v84_: D5
            d_309_v84_ = D5_DC15(589, d_290_v66_, ((d_296_v71_).set(d_294_v70_, d_290_v66_)) | (d_296_v71_), (_dafny.Map({d_298_v73_: d_304_v80_})) | ((d_308_v83_).cf28))
            d_309_v84_ = d_309_v84_
            if (default__.fm22((d_297_v72_).f23, (self).f25, not((self).f25), (d_297_v72_).f23, globalState)).ispropersubset(d_301_v76_):
                d_310_v85_: _dafny.Array
                def lambda10_(d_311_v76_):
                    def lambda11_(d_312_i9_):
                        return d_311_v76_

                    return lambda11_

                init7_ = lambda10_(d_301_v76_)
                nw64_ = _dafny.Array(None, 1)
                for i0_7_ in range(nw64_.length(0)):
                    nw64_[i0_7_] = init7_(i0_7_)
                d_310_v85_ = nw64_
                d_313_v86_: _dafny.Seq
                d_313_v86_ = _dafny.SeqWithoutIsStrInference([d_310_v85_])
                d_310_v85_ = (d_313_v86_)[default__.safeIndex(-367, len(d_313_v86_))]
                (globalState).f5 = (d_297_v72_).f23
                d_314_v87_: _dafny.Map
                d_314_v87_ = _dafny.Map({d_293_v69_.f31: len(d_302_v77_)})
                (globalState).f1 = (d_303_v79_).isdisjoint(_dafny.Set({((d_314_v87_)[len(_dafny.Map({d_314_v87_: (d_297_v72_).f23}))] if (len(_dafny.Map({d_314_v87_: (d_297_v72_).f23}))) in (d_314_v87_) else (d_297_v72_).f23), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdjrt"))): (self).f25})), d_293_v69_.f31}))
                d_315_v88_: _dafny.Array
                nw65_ = _dafny.Array(D5.default()(), 9)
                d_315_v88_ = nw65_
                index51_ = default__.safeIndex(872, (d_315_v88_).length(0))
                (d_315_v88_)[index51_] = (D5_DC15(d_257_v36_, d_290_v66_, d_296_v71_, ((d_309_v84_).cf28).set(D3_DC9(d_257_v36_), d_304_v80_)) if self.f24 else d_309_v84_)
                index52_ = default__.safeIndex(872, (d_315_v88_).length(0))
                (d_315_v88_)[index52_] = d_308_v83_
                d_314_v87_ = (d_314_v87_).set(446, default__.safeModuloInt(len(d_256_v35_), (d_297_v72_).f23))
            elif True:
                index53_ = default__.safeIndex(908, (d_304_v80_).length(0))
                (d_304_v80_)[index53_] = d_293_v69_.f31
                index54_ = default__.safeIndex(908, (d_304_v80_).length(0))
                (d_304_v80_)[index54_] = (0) - (len((d_293_v69_).f30))
                d_316_v89_: _dafny.Array
                nw66_ = _dafny.Array(int(0), 1)
                d_316_v89_ = nw66_
                d_317_v90_: _dafny.Seq
                d_317_v90_ = _dafny.SeqWithoutIsStrInference([d_316_v89_])
                d_318_v91_: _dafny.Array
                nw67_ = _dafny.Array(None, 9)
                nw67_[int(0)] = d_316_v89_
                nw67_[int(1)] = d_316_v89_
                nw67_[int(2)] = d_316_v89_
                nw67_[int(3)] = d_316_v89_
                nw67_[int(4)] = d_316_v89_
                nw67_[int(5)] = d_316_v89_
                nw67_[int(6)] = (d_317_v90_)[default__.safeIndex((d_302_v77_)[default__.safeIndex((d_297_v72_).f23, len(d_302_v77_))], len(d_317_v90_))]
                nw67_[int(7)] = d_316_v89_
                nw67_[int(8)] = d_316_v89_
                d_318_v91_ = nw67_
                index55_ = default__.safeIndex(82, (d_318_v91_).length(0))
                (d_318_v91_)[index55_] = d_316_v89_
                index56_ = default__.safeIndex(82, (d_318_v91_).length(0))
                (d_318_v91_)[index56_] = d_316_v89_
                r3 = (0) - (d_293_v69_.f31)
                d_319_v92_: C0
                nw68_ = C0()
                nw68_.ctor__(d_292_v68_, (d_297_v72_).f23, (d_297_v72_).f23)
                d_319_v92_ = nw68_
                (globalState).f6 = self.f24
            d_320_v93_: _dafny.Array
            def lambda12_(d_321_i10_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqbkxbwhd"))

            init8_ = lambda12_
            nw69_ = _dafny.Array(None, 24)
            for i0_8_ in range(nw69_.length(0)):
                nw69_[i0_8_] = init8_(i0_8_)
            d_320_v93_ = nw69_
            index57_ = default__.safeIndex(193, (d_320_v93_).length(0))
            (d_320_v93_)[index57_] = d_207_v0_
            index58_ = default__.safeIndex(193, (d_320_v93_).length(0))
            (d_320_v93_)[index58_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "easmd"))
        r2 = (self).f25
        d_322_v94_: str
        d_322_v94_ = _dafny.CodePoint('x')
        d_322_v94_ = d_322_v94_
        r0 = self.f24
        d_323_v95_: _dafny.Array
        nw70_ = _dafny.Array(False, 18)
        d_323_v95_ = nw70_
        d_324_v96_: _dafny.Set
        d_324_v96_ = _dafny.Set({d_323_v95_, d_323_v95_, d_323_v95_})
        d_325_v97_: T0
        nw71_ = C0()
        nw71_.ctor__((_dafny.Set({d_323_v95_, d_323_v95_})) | (d_324_v96_), d_257_v36_, len(d_258_v37_))
        d_325_v97_ = nw71_
        r1 = d_325_v97_
        r2 = not(self.f24)
        d_326_v98_: _dafny.Map
        d_326_v98_ = _dafny.Map({-174: default__.safeDivisionInt(d_257_v36_, -599)})
        d_327_v99_: _dafny.Map
        d_327_v99_ = _dafny.Map({d_256_v35_: _dafny.Set({(self).f25, (self).f25})})
        r3 = (0) - ((0) - (((d_326_v98_)[d_257_v36_] if (d_257_v36_) in (d_326_v98_) else default__.safeModuloInt(len(d_327_v99_), d_257_v36_))))
        return r0, r1, r2, r3

    def m8(self, p0, p1, p2, p3, globalState):
        (globalState).f8 = p1
        (globalState).f6 = default__.fm29(856, p1, (self).f25, globalState)
        d_328_v0_: _dafny.Array
        nw72_ = _dafny.Array(None, 6)
        nw72_[int(0)] = p3
        nw72_[int(1)] = self.f24
        nw72_[int(2)] = p3
        nw72_[int(3)] = self.f24
        nw72_[int(4)] = (self).f25
        nw72_[int(5)] = (self).f25
        d_328_v0_ = nw72_
        d_329_v1_: _dafny.Set
        d_329_v1_ = _dafny.Set({d_328_v0_, d_328_v0_})
        d_330_v2_: C0
        nw73_ = C0()
        nw73_.ctor__(d_329_v1_, (0) - ((0) - (p1)), 372)
        d_330_v2_ = nw73_
        d_331_v3_: _dafny.Set
        d_331_v3_ = _dafny.Set({p0})
        d_332_i0_: int
        d_332_i0_ = 0
        with _dafny.label("2"):
            while (d_331_v3_).ispropersubset((_dafny.Set({(self).f25})) | (d_331_v3_)):
                with _dafny.c_label("2"):
                    if (d_332_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_332_i0_ = (d_332_i0_) + (1)
                    d_333_v4_: _dafny.Seq
                    d_333_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "helwwm"))
                    d_334_v5_: _dafny.Seq
                    d_334_v5_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_335_i1_ in range(default__.abs(-671))])])
                    d_336_v6_: _dafny.Map
                    d_336_v6_ = _dafny.Map({self.f24: p1})
                    d_337_v7_: _dafny.Seq
                    d_337_v7_ = _dafny.SeqWithoutIsStrInference([len(d_336_v6_)])
                    index59_ = default__.safeIndex(659, (d_328_v0_).length(0))
                    (d_328_v0_)[index59_] = (d_333_v4_) == ((d_334_v5_)[default__.safeIndex(len(d_337_v7_), len(d_334_v5_))])
                    index60_ = default__.safeIndex(659, (d_328_v0_).length(0))
                    (d_328_v0_)[index60_] = True
                    (globalState).f9 = not(not(self.f24))
                    (globalState).f15 = (d_328_v0_)[default__.safeIndex(659, (d_328_v0_).length(0))]
                    d_338_v8_: _dafny.Array
                    def lambda13_(d_339_p0_):
                        def lambda14_(d_340_i2_):
                            return not(d_339_p0_)

                        return lambda14_

                    init9_ = lambda13_(p0)
                    nw74_ = _dafny.Array(None, 16)
                    for i0_9_ in range(nw74_.length(0)):
                        nw74_[i0_9_] = init9_(i0_9_)
                    d_338_v8_ = nw74_
                    d_341_v9_: _dafny.Map
                    d_341_v9_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfq")): d_338_v8_})
                    d_341_v9_ = ((d_341_v9_) | (d_341_v9_)) | (d_341_v9_)
                    pass
            pass
        (globalState).f9 = p3
        d_342_v10_: _dafny.Map
        d_342_v10_ = _dafny.Map({p3: d_330_v2_.f31})
        source3_ = D0_DC1(not((d_342_v10_) == (d_342_v10_)), p1)
        if source3_.is_DC1:
            d_343___mcc_h0_ = source3_.cf1
            d_344___mcc_h1_ = source3_.cf2
            d_345_cf2_ = d_344___mcc_h1_
            d_346_cf1_ = d_343___mcc_h0_
            (globalState).f1 = p3
            d_347_v11_: _dafny.Map
            d_347_v11_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_330_v2_.f31, d_330_v2_.f31, 483, d_345_cf2_]): False})
            d_347_v11_ = d_347_v11_
            d_348_v12_: _dafny.Seq
            d_348_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovboiepi"))
            (globalState).f5 = (0) - (((self).fm5(d_345_cf2_, p1, globalState)) * (default__.fm1(len(d_348_v12_), (self).f25, globalState)))
            (d_330_v2_).f31 = (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqn"))), d_330_v2_.f31)) + (d_330_v2_.f31)
        elif source3_.is_DC0:
            d_349___mcc_h2_ = source3_.cf0
            d_350_cf0_ = d_349___mcc_h2_
            d_351_v13_: D4
            d_351_v13_ = D4_DC12(len(_dafny.SeqWithoutIsStrInference([(self).f25])), p1)
            d_352_v14_: _dafny.Array
            nw75_ = _dafny.Array(None, 3)
            nw75_[int(0)] = 985
            nw75_[int(1)] = (self).fm5((d_351_v13_).cf19, ((d_342_v10_)[self.f24] if (self.f24) in (d_342_v10_) else d_330_v2_.f31), globalState)
            nw75_[int(2)] = len(_dafny.SeqWithoutIsStrInference([p3, not((self).f25)]))
            d_352_v14_ = nw75_
            d_353_v15_: _dafny.Map
            d_353_v15_ = _dafny.Map({304: d_352_v14_})
            d_354_v16_: _dafny.Seq
            d_354_v16_ = _dafny.SeqWithoutIsStrInference([d_352_v14_, d_352_v14_, d_352_v14_, d_352_v14_, ((d_353_v15_)[p1] if (p1) in (d_353_v15_) else d_352_v14_)])
            index61_ = default__.safeIndex(451, (d_328_v0_).length(0))
            (d_328_v0_)[index61_] = default__.fm29(d_330_v2_.f31, len(d_342_v10_), not(d_350_cf0_), globalState)
            d_355_v17_: _dafny.Set
            d_355_v17_ = _dafny.Set({d_330_v2_.f31, d_330_v2_.f31, (0) - (p1)})
            d_356_v18_: _dafny.Map
            d_356_v18_ = _dafny.Map({d_330_v2_.f31: len(d_355_v17_)})
            d_357_v19_: _dafny.Seq
            d_357_v19_ = _dafny.SeqWithoutIsStrInference([(self).f25])
            index62_ = default__.safeIndex(451, (d_328_v0_).length(0))
            rhs74_ = (d_354_v16_) + (d_354_v16_)
            rhs75_ = self.f24
            rhs76_ = default__.safeDivisionInt(p1, d_330_v2_.f31)
            rhs77_ = ((p0 if self.f24 else False)) == ((335) >= (((d_356_v18_)[len(d_357_v19_)] if (len(d_357_v19_)) in (d_356_v18_) else -467)))
            rhs78_ = d_330_v2_.f31
            lhs62_ = globalState
            lhs63_ = globalState
            lhs64_ = d_328_v0_
            lhs65_ = default__.safeIndex(451, (d_328_v0_).length(0))
            lhs66_ = globalState
            d_354_v16_ = rhs74_
            lhs62_.f6 = rhs75_
            lhs63_.f14 = rhs76_
            lhs64_[lhs65_] = rhs77_
            lhs66_.f8 = rhs78_
            d_358_v20_: _dafny.Array
            nw76_ = _dafny.Array(False, 11)
            d_358_v20_ = nw76_
            d_359_v21_: str
            d_359_v21_ = _dafny.CodePoint('d')
            d_360_v22_: C0
            nw77_ = C0()
            nw77_.ctor__(_dafny.Set({d_358_v20_}), len(default__.fm25((d_357_v19_)[default__.safeIndex(len(d_356_v18_), len(d_357_v19_))], self.f24, d_359_v21_, globalState)), p1)
            d_360_v22_ = nw77_
            (globalState).f6 = (d_357_v19_)[default__.safeIndex((p1) - (d_360_v22_.f31), len(d_357_v19_))]
            d_361_v24_: _dafny.Array
            def lambda15_(d_362_v21_, d_363_v22_):
                def lambda16_(d_364_i3_):
                    def iife18_():
                        coll14_ = _dafny.Map()
                        compr_14_: str
                        for compr_14_ in (_dafny.Map({d_362_v21_: d_362_v21_})).keys.Elements:
                            d_365_v23_: str = compr_14_
                            if (d_365_v23_) in (_dafny.Map({d_362_v21_: d_362_v21_})):
                                coll14_[d_365_v23_] = d_363_v22_.f31
                        return _dafny.Map(coll14_)
                    return (D9_DC25(_dafny.SeqWithoutIsStrInference([iife18_()
]))).cf50

                return lambda16_

            init10_ = lambda15_(d_359_v21_, d_360_v22_)
            nw78_ = _dafny.Array(None, 5)
            for i0_10_ in range(nw78_.length(0)):
                nw78_[i0_10_] = init10_(i0_10_)
            d_361_v24_ = nw78_
            d_366_v25_: _dafny.Map
            d_366_v25_ = _dafny.Map({d_359_v21_: (0) - (d_360_v22_.f31)})
            d_367_v27_: _dafny.Seq
            d_367_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('f')])
            d_368_v28_: _dafny.Seq
            def iife19_():
                coll15_ = _dafny.Map()
                compr_15_: str
                for compr_15_ in (d_367_v27_).Elements:
                    d_369_v26_: str = compr_15_
                    if (d_369_v26_) in (d_367_v27_):
                        coll15_[d_369_v26_] = d_330_v2_.f31
                return _dafny.Map(coll15_)
            d_368_v28_ = _dafny.SeqWithoutIsStrInference([d_366_v25_, d_366_v25_, d_366_v25_, iife19_()
            ])
            index63_ = default__.safeIndex(468, (d_361_v24_).length(0))
            (d_361_v24_)[index63_] = d_368_v28_
            index64_ = default__.safeIndex(468, (d_361_v24_).length(0))
            (d_361_v24_)[index64_] = d_368_v28_
        elif True:
            d_370___mcc_h3_ = source3_.cf3
            d_371_cf3_ = d_370___mcc_h3_
            d_372_v29_: _dafny.Seq
            d_372_v29_ = _dafny.SeqWithoutIsStrInference([d_330_v2_.f31])
            d_373_v30_: C0
            nw79_ = C0()
            nw79_.ctor__((d_330_v2_).f30, d_330_v2_.f31, len(d_372_v29_))
            d_373_v30_ = nw79_
            d_374_v31_: _dafny.MultiSet
            d_374_v31_ = _dafny.MultiSet([p0, p0, p3])
            d_375_v32_: _dafny.Map
            d_375_v32_ = _dafny.Map({(self).f25: d_374_v31_})
            d_376_v33_: _dafny.Seq
            d_376_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vb"))
            d_375_v32_ = (d_375_v32_).set((d_330_v2_).fm3(d_376_v33_, (d_373_v30_).fm3(d_376_v33_, (self).f25, p1, _dafny.SeqWithoutIsStrInference([len(d_376_v33_)]), globalState), d_373_v30_.f31, d_372_v29_, globalState), d_374_v31_)
            d_377_v34_: _dafny.MultiSet
            d_377_v34_ = _dafny.MultiSet([d_331_v3_])
            d_378_v35_: _dafny.MultiSet
            d_378_v35_ = _dafny.MultiSet([d_373_v30_.f31])
            d_379_v36_: _dafny.MultiSet
            d_379_v36_ = _dafny.MultiSet([(d_377_v34_).cardinality, ((d_378_v35_)[len(d_372_v29_)] if (len(d_372_v29_)) in (d_378_v35_) else d_330_v2_.f31)])
            d_380_v37_: _dafny.Map
            d_380_v37_ = _dafny.Map({len(d_376_v33_): p0})
            d_381_v38_: _dafny.Array
            nw80_ = _dafny.Array(None, 10)
            nw80_[int(0)] = d_330_v2_.f31
            nw80_[int(1)] = (0) - (d_330_v2_.f31)
            nw80_[int(2)] = default__.safeDivisionInt(d_330_v2_.f31, d_330_v2_.f31)
            nw80_[int(3)] = ((d_378_v35_)[d_373_v30_.f31] if (d_373_v30_.f31) in (d_378_v35_) else (0) - (d_373_v30_.f31))
            nw80_[int(4)] = (d_373_v30_.f31) + (d_330_v2_.f31)
            nw80_[int(5)] = (d_330_v2_.f31) + ((0) - (d_330_v2_.f31))
            nw80_[int(6)] = (p1) * (p1)
            nw80_[int(7)] = (d_330_v2_).fm2(globalState)
            nw80_[int(8)] = len(d_331_v3_)
            nw80_[int(9)] = len((d_380_v37_) | (d_380_v37_))
            d_381_v38_ = nw80_
            d_381_v38_ = d_381_v38_
            d_376_v33_ = d_376_v33_

    def m9(self, p0, p1, p2, p3, globalState):
        d_382_v0_: _dafny.Seq
        d_382_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nyt"))
        d_383_v1_: _dafny.Map
        d_383_v1_ = _dafny.Map({self.f24: p3})
        d_384_v2_: D3
        d_384_v2_ = D3_DC9(len(d_383_v1_))
        hi2_ = (d_384_v2_).cf15
        for d_385_i0_ in range(len((d_382_v0_) + (d_382_v0_)), hi2_):
            d_386_v3_: _dafny.Seq
            d_386_v3_ = _dafny.SeqWithoutIsStrInference([(self).f25])
            (globalState).f9 = default__.fm29(444, d_385_i0_, (d_386_v3_)[default__.safeIndex(d_385_i0_, len(d_386_v3_))], globalState)
            d_387_v4_: _dafny.Map
            d_387_v4_ = _dafny.Map({p2: p0})
            d_388_v5_: _dafny.Map
            d_388_v5_ = _dafny.Map({d_387_v4_: (self).f25})
            (self).m8(self.f24, p3, (d_388_v5_).set(d_387_v4_, self.f24), default__.fm29(d_385_i0_, (0) - (d_385_i0_), p2, globalState), globalState)
            d_389_v6_: _dafny.Map
            d_389_v6_ = _dafny.Map({self.f24: self.f24})
            d_390_v7_: _dafny.Map
            d_390_v7_ = _dafny.Map({d_389_v6_: p3})
            (self).f24 = default__.fm29((0) - (d_385_i0_), ((d_390_v7_)[d_389_v6_] if (d_389_v6_) in (d_390_v7_) else d_385_i0_), (self).f25, globalState)
            d_391_v8_: _dafny.Array
            nw81_ = _dafny.Array(False, 10)
            d_391_v8_ = nw81_
            d_392_v9_: _dafny.Set
            d_392_v9_ = _dafny.Set({d_391_v8_})
            d_393_v10_: C0
            nw82_ = C0()
            nw82_.ctor__((d_392_v9_ if p1 else d_392_v9_), d_385_i0_, (len(d_382_v0_)) - (len(d_382_v0_)))
            d_393_v10_ = nw82_
        d_394_v11_: _dafny.Map
        d_394_v11_ = _dafny.Map({p0: p2})
        d_395_v12_: _dafny.Array
        nw83_ = _dafny.Array(None, 10)
        nw83_[int(0)] = p3
        nw83_[int(1)] = p3
        nw83_[int(2)] = len(d_382_v0_)
        nw83_[int(3)] = -13
        nw83_[int(4)] = p3
        nw83_[int(5)] = p3
        nw83_[int(6)] = p3
        nw83_[int(7)] = p3
        nw83_[int(8)] = len(d_382_v0_)
        nw83_[int(9)] = (len(d_394_v11_)) * (p3)
        d_395_v12_ = nw83_
        index65_ = default__.safeIndex(378, (d_395_v12_).length(0))
        (d_395_v12_)[index65_] = (0) - (len(d_382_v0_))
        d_396_v13_: _dafny.Seq
        d_396_v13_ = _dafny.SeqWithoutIsStrInference([self.f24, p2])
        index66_ = default__.safeIndex(378, (d_395_v12_).length(0))
        (d_395_v12_)[index66_] = (p3) - (len(d_396_v13_))
        (globalState).f14 = (0) - (p3)
        index67_ = default__.safeIndex(378, (d_395_v12_).length(0))
        (d_395_v12_)[index67_] = (d_395_v12_)[default__.safeIndex(378, (d_395_v12_).length(0))]
        d_397_v14_: _dafny.Seq
        d_397_v14_ = _dafny.SeqWithoutIsStrInference([p3])
        d_398_v15_: _dafny.MultiSet
        d_398_v15_ = _dafny.MultiSet([p3, (d_395_v12_)[default__.safeIndex(378, (d_395_v12_).length(0))], (d_395_v12_)[default__.safeIndex(378, (d_395_v12_).length(0))]])
        d_399_v16_: _dafny.Array
        nw84_ = _dafny.Array(None, 23)
        nw84_[int(0)] = p2
        nw84_[int(1)] = False
        nw84_[int(2)] = p1
        nw84_[int(3)] = self.f24
        nw84_[int(4)] = (self).f25
        nw84_[int(5)] = not(p1)
        nw84_[int(6)] = (self).f25
        nw84_[int(7)] = p2
        nw84_[int(8)] = self.f24
        nw84_[int(9)] = (self).f25
        nw84_[int(10)] = (self).f25
        nw84_[int(11)] = p1
        nw84_[int(12)] = p1
        nw84_[int(13)] = p2
        nw84_[int(14)] = (self).f25
        nw84_[int(15)] = (self).f25
        nw84_[int(16)] = self.f24
        nw84_[int(17)] = (self).f25
        nw84_[int(18)] = p2
        nw84_[int(19)] = p2
        nw84_[int(20)] = self.f24
        nw84_[int(21)] = p1
        nw84_[int(22)] = (self).f25
        d_399_v16_ = nw84_
        d_400_v17_: _dafny.Seq
        d_400_v17_ = _dafny.SeqWithoutIsStrInference([d_399_v16_, d_399_v16_])
        d_401_v18_: _dafny.MultiSet
        d_401_v18_ = _dafny.MultiSet([p2])
        d_402_v19_: _dafny.Seq
        d_402_v19_ = _dafny.SeqWithoutIsStrInference([len(d_397_v14_), ((d_398_v15_)[len(d_400_v17_)] if (len(d_400_v17_)) in (d_398_v15_) else (d_401_v18_).cardinality)])
        d_403_v20_: _dafny.Seq
        d_403_v20_ = _dafny.SeqWithoutIsStrInference([(d_397_v14_)[default__.safeIndex(p3, len(d_397_v14_))], (d_395_v12_)[default__.safeIndex(378, (d_395_v12_).length(0))]])
        rhs79_ = (0) - (p3)
        rhs80_ = default__.safeModuloInt((0) - ((d_395_v12_)[default__.safeIndex(378, (d_395_v12_).length(0))]), p3)
        rhs81_ = p1
        rhs82_ = _dafny.SeqWithoutIsStrInference([p0 for d_404_i1_ in range(default__.abs(812))])
        rhs83_ = ((d_403_v20_ if not(default__.fm29(724, p3, (self).f25, globalState)) else ((D10_DC28(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))) for d_405_i2_ in range(default__.abs(820))]))).cf52).set(default__.safeIndex(p3, len((D10_DC28(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))) for d_406_i2_ in range(default__.abs(820))]))).cf52)), p3))) <= (d_397_v14_)
        lhs67_ = globalState
        lhs68_ = globalState
        lhs69_ = self
        lhs70_ = globalState
        lhs67_.f14 = rhs79_
        lhs68_.f8 = rhs80_
        lhs69_.f24 = rhs81_
        d_382_v0_ = rhs82_
        lhs70_.f15 = rhs83_
        (globalState).f15 = default__.fm29(p3, (p3) * ((d_395_v12_)[default__.safeIndex(378, (d_395_v12_).length(0))]), (self).f25, globalState)


class C2(T0):
    def  __init__(self):
        self._f23: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f23(self):
        return self._f23
    def ctor__(self, f23):
        (self)._f23 = f23

    def fm2(self, globalState):
        return (self).f23

    def fm3(self, p0, p1, p2, p3, globalState):
        return not ((_dafny.SeqWithoutIsStrInference([not(False), False])) <= (_dafny.SeqWithoutIsStrInference([not(False)]))) or ((_dafny.MultiSet([(self).f23])).ispropersubset(_dafny.MultiSet([(self).f23, (self).f23])))

    def fm19(self, p0, p1, p2, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, True, True]), _dafny.SeqWithoutIsStrInference([False, True]), _dafny.SeqWithoutIsStrInference([True])])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, False, True]), _dafny.SeqWithoutIsStrInference([False])]))

    def m0(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: bool = False
        r0 = (self).f23
        if (p0) or (p0):
            d_407_v0_: _dafny.Array
            def lambda17_(d_408_p1_):
                def lambda18_(d_409_i0_):
                    return (d_408_p1_ if d_408_p1_ else not(d_408_p1_))

                return lambda18_

            init11_ = lambda17_(p1)
            nw85_ = _dafny.Array(None, 5)
            for i0_11_ in range(nw85_.length(0)):
                nw85_[i0_11_] = init11_(i0_11_)
            d_407_v0_ = nw85_
            index68_ = default__.safeIndex(90, (d_407_v0_).length(0))
            (d_407_v0_)[index68_] = p0
            d_410_v1_: _dafny.Seq
            d_410_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbhsiypk"))
            index69_ = default__.safeIndex(90, (d_407_v0_).length(0))
            rhs84_ = default__.safeDivisionInt((self).f23, 982)
            rhs85_ = p0
            rhs86_ = False
            rhs87_ = d_410_v1_
            lhs71_ = globalState
            lhs72_ = d_407_v0_
            lhs73_ = default__.safeIndex(90, (d_407_v0_).length(0))
            lhs74_ = globalState
            lhs71_.f8 = rhs84_
            lhs72_[lhs73_] = rhs85_
            lhs74_.f1 = rhs86_
            d_410_v1_ = rhs87_
            d_411_v2_: D3
            d_411_v2_ = D3_DC8(_dafny.SeqWithoutIsStrInference([(D1_DC3(_dafny.CodePoint('m'))).cf4 for d_412_i1_ in range(default__.abs(-517))]))
            d_413_v3_: _dafny.Map
            d_413_v3_ = _dafny.Map({p1: (d_407_v0_)[default__.safeIndex(90, (d_407_v0_).length(0))]})
            d_414_v4_: _dafny.Set
            d_414_v4_ = _dafny.Set({p0, p0})
            rhs88_ = (len((d_413_v3_) | (d_413_v3_))) - ((self).f23)
            rhs89_ = d_411_v2_
            rhs90_ = ((self).f23) - ((self).f23)
            rhs91_ = (p1) and (p0)
            rhs92_ = default__.safeDivisionInt(len((d_414_v4_) - (_dafny.Set({False}))), (self).f23)
            lhs75_ = globalState
            lhs76_ = globalState
            lhs77_ = globalState
            lhs78_ = globalState
            lhs75_.f17 = rhs88_
            d_411_v2_ = rhs89_
            lhs76_.f14 = rhs90_
            lhs77_.f6 = rhs91_
            lhs78_.f17 = rhs92_
            (globalState).f14 = (self).f23
            d_415_v5_: _dafny.Map
            d_415_v5_ = _dafny.Map({d_410_v1_: (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqfaosxj")))) < ((self).f23)})
            d_416_v6_: _dafny.Set
            d_416_v6_ = _dafny.Set({(self).f23})
            d_417_v7_: _dafny.Seq
            d_417_v7_ = _dafny.SeqWithoutIsStrInference([(d_407_v0_)[default__.safeIndex(90, (d_407_v0_).length(0))], (d_407_v0_)[default__.safeIndex(90, (d_407_v0_).length(0))]])
            rhs93_ = d_415_v5_
            rhs94_ = default__.safeDivisionInt((self).f23, ((self).f23) + ((self).f23))
            rhs95_ = ((self).f23 if (d_416_v6_).ispropersubset(d_416_v6_) else len(d_417_v7_))
            rhs96_ = (self).f23
            lhs79_ = globalState
            lhs80_ = globalState
            lhs81_ = globalState
            d_415_v5_ = rhs93_
            lhs79_.f17 = rhs94_
            lhs80_.f14 = rhs95_
            lhs81_.f17 = rhs96_
            d_418_v8_: _dafny.Array
            nw86_ = _dafny.Array(_dafny.Map({}), 17)
            d_418_v8_ = nw86_
            d_419_v9_: _dafny.Map
            d_419_v9_ = _dafny.Map({d_407_v0_: p1})
            index70_ = default__.safeIndex(4, (d_418_v8_).length(0))
            (d_418_v8_)[index70_] = d_419_v9_
            index71_ = default__.safeIndex(4, (d_418_v8_).length(0))
            (d_418_v8_)[index71_] = (D4_DC11(d_419_v9_)).cf18
        elif True:
            d_420_v10_: _dafny.Array
            nw87_ = _dafny.Array(None, 23)
            nw87_[int(0)] = not(p1)
            nw87_[int(1)] = p1
            nw87_[int(2)] = p0
            nw87_[int(3)] = p1
            nw87_[int(4)] = p1
            nw87_[int(5)] = False
            nw87_[int(6)] = p0
            nw87_[int(7)] = not(p1)
            nw87_[int(8)] = p0
            nw87_[int(9)] = p0
            nw87_[int(10)] = not(False)
            nw87_[int(11)] = p0
            nw87_[int(12)] = p0
            nw87_[int(13)] = p0
            nw87_[int(14)] = p1
            nw87_[int(15)] = p1
            nw87_[int(16)] = p1
            nw87_[int(17)] = p1
            nw87_[int(18)] = p1
            nw87_[int(19)] = p0
            nw87_[int(20)] = p1
            nw87_[int(21)] = p0
            nw87_[int(22)] = p0
            d_420_v10_ = nw87_
            d_421_v11_: _dafny.Seq
            d_421_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wt"))
            d_422_v12_: _dafny.Map
            d_422_v12_ = _dafny.Map({p0: d_421_v11_})
            d_423_v13_: _dafny.MultiSet
            d_423_v13_ = _dafny.MultiSet([d_421_v11_, ((d_422_v12_)[p1] if (p1) in (d_422_v12_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_424_i2_ in range(default__.abs(-505))])), d_421_v11_])
            d_425_v14_: _dafny.Seq
            d_425_v14_ = _dafny.SeqWithoutIsStrInference([d_423_v13_, d_423_v13_, d_423_v13_])
            d_426_v15_: _dafny.Map
            d_426_v15_ = _dafny.Map({341: ((d_425_v14_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_427_i3_ in range(default__.abs(861))])), len(d_425_v14_))]).cardinality})
            d_428_v16_: _dafny.Map
            d_428_v16_ = _dafny.Map({d_420_v10_: d_426_v15_})
            d_428_v16_ = (d_428_v16_) | (d_428_v16_)
            (globalState).f1 = False
            d_429_v17_: _dafny.Array
            nw88_ = _dafny.Array(int(0), 21)
            d_429_v17_ = nw88_
            index72_ = default__.safeIndex(406, (d_429_v17_).length(0))
            (d_429_v17_)[index72_] = ((self).f23) - ((self).f23)
            index73_ = default__.safeIndex(406, (d_429_v17_).length(0))
            (d_429_v17_)[index73_] = default__.safeModuloInt((self).f23, ((self).f23) * ((self).f23))
            (globalState).f15 = p1
            d_430_v18_: _dafny.Map
            d_430_v18_ = _dafny.Map({True: p1})
            d_431_v19_: _dafny.Map
            d_431_v19_ = _dafny.Map({p1: 431})
            d_432_v20_: D1
            d_432_v20_ = D1_DC5(p0, d_430_v18_, (self).f23, ((d_426_v15_)[(self).f23] if ((self).f23) in (d_426_v15_) else (d_429_v17_)[default__.safeIndex(406, (d_429_v17_).length(0))]), d_431_v19_)
            d_433_v21_: _dafny.Map
            d_433_v21_ = _dafny.Map({(d_429_v17_)[default__.safeIndex(406, (d_429_v17_).length(0))]: d_432_v20_})
            d_433_v21_ = (d_433_v21_).set((d_429_v17_)[default__.safeIndex(406, (d_429_v17_).length(0))], d_432_v20_)
        d_434_v22_: _dafny.Map
        d_434_v22_ = _dafny.Map({not(not(p0)): p0})
        d_435_v23_: D3
        d_435_v23_ = D3_DC10(not(p1), (d_434_v22_) | (d_434_v22_))
        d_436_v24_: str
        d_436_v24_ = _dafny.CodePoint('g')
        d_437_v26_: _dafny.Map
        def iife20_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in _dafny.IntegerRange(240, 994):
                d_438_v25_: int = compr_16_
                if ((240) <= (d_438_v25_)) and ((d_438_v25_) < (994)):
                    coll16_[(d_438_v25_) - ((self).f23)] = (self).f23
            return _dafny.Map(coll16_)
        d_437_v26_ = _dafny.Map({iife20_()
        : (self).f23})
        d_439_v27_: D5
        d_439_v27_ = D5_DC13(d_437_v26_)
        d_440_v28_: _dafny.Seq
        d_440_v28_ = _dafny.SeqWithoutIsStrInference([False])
        d_441_v29_: _dafny.Seq
        d_441_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngg"))
        rhs97_ = len((d_439_v27_).cf21)
        rhs98_ = ((d_440_v28_)[default__.safeIndex((self).f23, len(d_440_v28_))]) and (not((d_441_v29_) < (d_441_v29_)))
        rhs99_ = d_435_v23_
        rhs100_ = d_436_v24_
        lhs82_ = globalState
        lhs83_ = globalState
        lhs82_.f8 = rhs97_
        lhs83_.f15 = rhs98_
        d_435_v23_ = rhs99_
        d_436_v24_ = rhs100_
        d_441_v29_ = (d_441_v29_) + (default__.fm20(p0, globalState))
        d_442_v30_: _dafny.Map
        d_442_v30_ = _dafny.Map({p1: (self).f23})
        d_443_v31_: D1
        d_443_v31_ = D1_DC5(True, d_434_v22_, (self).f23, (self).f23, d_442_v30_)
        d_444_v32_: _dafny.Map
        d_444_v32_ = _dafny.Map({p0: len(((d_443_v31_).cf8) | (d_434_v22_))})
        d_445_v33_: _dafny.Seq
        d_445_v33_ = _dafny.SeqWithoutIsStrInference([(self).f23, (self).f23, (self).f23])
        d_444_v32_ = (d_444_v32_).set((self).fm3(d_441_v29_, p1, (self).f23, d_445_v33_, globalState), ((self).f23) + (default__.fm1((self).f23, p1, globalState)))
        d_446_v34_: _dafny.MultiSet
        d_446_v34_ = _dafny.MultiSet([p0])
        d_447_v35_: _dafny.Seq
        d_447_v35_ = _dafny.SeqWithoutIsStrInference([d_445_v33_])
        d_448_v36_: _dafny.MultiSet
        d_448_v36_ = _dafny.MultiSet([-143])
        d_449_v37_: _dafny.Set
        d_449_v37_ = _dafny.Set({(d_448_v36_).cardinality})
        rhs101_ = (self).fm3((d_441_v29_ if p1 else d_441_v29_), True, (self).f23, (d_447_v35_)[default__.safeIndex((self).f23, len(d_447_v35_))], globalState)
        rhs102_ = (d_449_v37_).issubset(d_449_v37_)
        rhs103_ = (_dafny.MultiSet([p0, p1])) | ((_dafny.MultiSet([p1, p0])).set(not(p0), default__.abs(len(_dafny.SeqWithoutIsStrInference([p0])))))
        lhs84_ = globalState
        r2 = rhs101_
        lhs84_.f6 = rhs102_
        d_446_v34_ = rhs103_
        r0 = default__.safeDivisionInt(30, (self).f23)
        d_450_v38_: _dafny.Array
        nw89_ = _dafny.Array(int(0), 20)
        d_450_v38_ = nw89_
        r1 = d_450_v38_
        r2 = (384) < ((self).f23)
        return r0, r1, r2

    def m6(self, p0, p1, p2, p3, globalState):
        r0: D0 = D0.default()()
        r1: int = int(0)
        r2: int = int(0)
        (globalState).f9 = p2
        d_451_v0_: _dafny.Seq
        d_451_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqyjmxlhd"))
        hi3_ = (self).fm2(globalState)
        for d_452_i0_ in range(((0) - (len(d_451_v0_))) + ((_dafny.MultiSet([D3_DC9(p3)])).cardinality), hi3_):
            d_453_v1_: _dafny.Map
            d_453_v1_ = _dafny.Map({p2: len(_dafny.Map({d_452_i0_: p3}))})
            d_454_v2_: D1
            d_454_v2_ = D1_DC5(False, _dafny.Map({p0: p2}), d_452_i0_, p3, d_453_v1_)
            d_455_v3_: _dafny.Map
            d_455_v3_ = _dafny.Map({(d_454_v2_).cf7: p1})
            d_456_v4_: _dafny.Map
            d_456_v4_ = _dafny.Map({d_454_v2_: d_452_i0_})
            d_457_v5_: _dafny.Seq
            d_457_v5_ = _dafny.SeqWithoutIsStrInference([d_456_v4_, d_456_v4_, d_456_v4_])
            d_458_v6_: _dafny.Seq
            d_458_v6_ = _dafny.SeqWithoutIsStrInference([d_457_v5_, d_457_v5_, d_457_v5_])
            d_459_v7_: _dafny.Map
            d_459_v7_ = _dafny.Map({len(d_455_v3_): ((d_458_v6_)[default__.safeIndex(d_452_i0_, len(d_458_v6_))]) + ((d_458_v6_)[default__.safeIndex(d_452_i0_, len(d_458_v6_))])})
            d_459_v7_ = (d_459_v7_).set((self).f23, d_457_v5_)
            d_453_v1_ = (d_453_v1_).set(not(p2), d_452_i0_)
            d_460_v8_: D3
            d_460_v8_ = D3_DC8(d_451_v0_)
            d_461_v9_: str
            d_461_v9_ = _dafny.CodePoint('i')
            d_462_v10_: _dafny.Seq
            d_462_v10_ = _dafny.SeqWithoutIsStrInference([p2, p0])
            d_463_v11_: _dafny.Set
            d_463_v11_ = _dafny.Set({p0})
            d_464_v12_: _dafny.MultiSet
            d_464_v12_ = _dafny.MultiSet([d_452_i0_])
            d_465_v13_: _dafny.Array
            nw90_ = _dafny.Array(None, 29)
            nw90_[int(0)] = p2
            nw90_[int(1)] = (self).fm3((d_460_v8_).cf14, p0, d_452_i0_, p1, globalState)
            nw90_[int(2)] = p2
            nw90_[int(3)] = p2
            nw90_[int(4)] = (d_451_v0_) <= ((d_451_v0_).set(default__.safeIndex(d_452_i0_, len(d_451_v0_)), d_461_v9_))
            nw90_[int(5)] = p0
            nw90_[int(6)] = p0
            nw90_[int(7)] = p0
            nw90_[int(8)] = p0
            nw90_[int(9)] = not(p0)
            nw90_[int(10)] = (d_462_v10_) < (_dafny.SeqWithoutIsStrInference([(d_462_v10_)[default__.safeIndex(default__.fm1((self).f23, not(p0), globalState), len(d_462_v10_))], p2]))
            nw90_[int(11)] = p0
            nw90_[int(12)] = not(p2)
            nw90_[int(13)] = not((p3) < ((self).f23))
            nw90_[int(14)] = (d_463_v11_).issubset(d_463_v11_)
            nw90_[int(15)] = (not((d_462_v10_)[default__.safeIndex(d_452_i0_, len(d_462_v10_))]) if p2 else not(not(p0)))
            nw90_[int(16)] = p2
            nw90_[int(17)] = p0
            nw90_[int(18)] = p2
            nw90_[int(19)] = p2
            nw90_[int(20)] = (self).fm3(d_451_v0_, p0, d_452_i0_, p1, globalState)
            nw90_[int(21)] = p0
            nw90_[int(22)] = True
            nw90_[int(23)] = p2
            nw90_[int(24)] = p0
            nw90_[int(25)] = p0
            nw90_[int(26)] = p0
            nw90_[int(27)] = (d_464_v12_).isdisjoint(d_464_v12_)
            nw90_[int(28)] = p0
            d_465_v13_ = nw90_
            index74_ = default__.safeIndex(748, (d_465_v13_).length(0))
            (d_465_v13_)[index74_] = True
            index75_ = default__.safeIndex(748, (d_465_v13_).length(0))
            (d_465_v13_)[index75_] = p2
            d_466_v14_: _dafny.Set
            d_466_v14_ = _dafny.Set({-582, d_452_i0_, len(_dafny.SeqWithoutIsStrInference([d_452_i0_ for d_467_i1_ in range(default__.abs(366))])), p3})
            d_468_v15_: _dafny.Map
            d_468_v15_ = _dafny.Map({p0: (_dafny.Set({(self).f23, p3})) != (d_466_v14_)})
            if ((d_468_v15_)[(len(d_451_v0_)) <= (len(d_451_v0_))] if ((len(d_451_v0_)) <= (len(d_451_v0_))) in (d_468_v15_) else p0):
                d_469_v16_: _dafny.Map
                d_469_v16_ = _dafny.Map({d_461_v9_: (self).f23})
                d_470_v17_: _dafny.Seq
                d_470_v17_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_461_v9_: default__.fm1((self).f23, True, globalState)}), d_469_v16_])
                d_470_v17_ = (d_470_v17_) + (d_470_v17_)
                d_471_v18_: _dafny.Array
                nw91_ = _dafny.Array(D0.default()(), 14)
                d_471_v18_ = nw91_
                d_472_v19_: _dafny.Map
                d_472_v19_ = _dafny.Map({d_471_v18_: d_461_v9_})
                d_473_v20_: D3
                d_473_v20_ = D3_DC9(d_452_i0_)
                d_474_v21_: _dafny.Array
                nw92_ = _dafny.Array(None, 29)
                nw92_[int(0)] = -441
                nw92_[int(1)] = (self).fm2(globalState)
                nw92_[int(2)] = p3
                nw92_[int(3)] = (self).f23
                nw92_[int(4)] = p3
                nw92_[int(5)] = d_452_i0_
                nw92_[int(6)] = (self).f23
                nw92_[int(7)] = (self).f23
                nw92_[int(8)] = len(p1)
                nw92_[int(9)] = (self).f23
                nw92_[int(10)] = (self).f23
                nw92_[int(11)] = -676
                nw92_[int(12)] = (self).f23
                nw92_[int(13)] = d_452_i0_
                nw92_[int(14)] = d_452_i0_
                nw92_[int(15)] = (self).f23
                nw92_[int(16)] = p3
                nw92_[int(17)] = (0) - (d_452_i0_)
                nw92_[int(18)] = (self).f23
                nw92_[int(19)] = (self).f23
                nw92_[int(20)] = len(d_463_v11_)
                nw92_[int(21)] = (self).f23
                nw92_[int(22)] = (self).f23
                nw92_[int(23)] = (self).f23
                nw92_[int(24)] = (self).f23
                nw92_[int(25)] = p3
                nw92_[int(26)] = p3
                nw92_[int(27)] = p3
                nw92_[int(28)] = (self).f23
                d_474_v21_ = nw92_
                d_475_v22_: _dafny.Map
                d_475_v22_ = _dafny.Map({d_473_v20_: d_474_v21_})
                d_476_v23_: D5
                d_476_v23_ = D5_DC15(292, ((d_451_v0_)[default__.safeIndex((self).f23, len(d_451_v0_))] if (d_465_v13_)[default__.safeIndex(748, (d_465_v13_).length(0))] else d_461_v9_), d_472_v19_, (d_475_v22_) | (d_475_v22_))
                d_477_v24_: _dafny.Array
                def lambda19_(d_478_p0_):
                    def lambda20_(d_479_i2_):
                        return _dafny.SeqWithoutIsStrInference([d_478_p0_])

                    return lambda20_

                init12_ = lambda19_(p0)
                nw93_ = _dafny.Array(None, 21)
                for i0_12_ in range(nw93_.length(0)):
                    nw93_[i0_12_] = init12_(i0_12_)
                d_477_v24_ = nw93_
                d_480_v25_: D6
                d_480_v25_ = D6_DC16(d_462_v10_)
                index76_ = default__.safeIndex(186, (d_477_v24_).length(0))
                (d_477_v24_)[index76_] = (d_480_v25_).cf29
                d_481_v26_: _dafny.Map
                d_481_v26_ = _dafny.Map({(-978 if p2 else (p1)[default__.safeIndex(-663, len(p1))]): d_461_v9_})
                pat_let_tv1_ = d_472_v19_
                pat_let_tv2_ = d_472_v19_
                pat_let_tv3_ = p3
                index77_ = default__.safeIndex(186, (d_477_v24_).length(0))
                def iife21_(_pat_let2_0):
                    def iife22_(d_482_dt__update__tmp_h0_):
                        def iife23_(_pat_let3_0):
                            def iife24_(d_483_dt__update_hcf27_h0_):
                                def iife25_(_pat_let4_0):
                                    def iife26_(d_484_dt__update_hcf25_h0_):
                                        return D5_DC15(d_484_dt__update_hcf25_h0_, (d_482_dt__update__tmp_h0_).cf26, d_483_dt__update_hcf27_h0_, (d_482_dt__update__tmp_h0_).cf28)
                                    return iife26_(_pat_let4_0)
                                return iife25_((pat_let_tv3_) - ((self).f23))
                            return iife24_(_pat_let3_0)
                        return iife23_((pat_let_tv1_) | (pat_let_tv2_))
                    return iife22_(_pat_let2_0)
                rhs104_ = iife21_(d_476_v23_)
                rhs105_ = d_465_v13_
                rhs106_ = d_465_v13_
                rhs107_ = d_462_v10_
                rhs108_ = (d_481_v26_) | (d_481_v26_)
                lhs85_ = d_477_v24_
                lhs86_ = default__.safeIndex(186, (d_477_v24_).length(0))
                d_476_v23_ = rhs104_
                d_465_v13_ = rhs105_
                d_465_v13_ = rhs106_
                lhs85_[lhs86_] = rhs107_
                d_481_v26_ = rhs108_
                (globalState).f5 = default__.fm1(p3, True, globalState)
                (globalState).f1 = (d_463_v11_).issubset(d_463_v11_)
                (globalState).f9 = not((d_465_v13_)[default__.safeIndex(748, (d_465_v13_).length(0))])
            elif True:
                d_485_v27_: D6
                d_485_v27_ = D6_DC16(d_462_v10_)
                d_485_v27_ = d_485_v27_
                index78_ = default__.safeIndex(748, (d_465_v13_).length(0))
                (d_465_v13_)[index78_] = p2
                d_486_v28_: _dafny.Map
                d_486_v28_ = _dafny.Map({p2: _dafny.MultiSet([(d_462_v10_)[default__.safeIndex((0) - (p3), len(d_462_v10_))]])})
                d_486_v28_ = (d_486_v28_).set((d_465_v13_)[default__.safeIndex(748, (d_465_v13_).length(0))], default__.fm21(p2, True, (self).f23, globalState))
                d_487_v29_: _dafny.Map
                d_487_v29_ = _dafny.Map({(d_452_i0_) - ((self).f23): (p3 if (d_465_v13_)[default__.safeIndex(748, (d_465_v13_).length(0))] else d_452_i0_)})
                d_487_v29_ = (d_487_v29_).set(p3, (self).f23)
                (globalState).f6 = (not(p0)) == (False)
        d_488_v30_: _dafny.Map
        d_488_v30_ = _dafny.Map({p2: (p3) - (p3)})
        (globalState).f0 = d_488_v30_
        d_489_v31_: _dafny.Array
        nw94_ = _dafny.Array(_dafny.Set({}), 9)
        d_489_v31_ = nw94_
        d_490_v32_: _dafny.Set
        d_490_v32_ = _dafny.Set({not(p0)})
        index79_ = default__.safeIndex(560, (d_489_v31_).length(0))
        (d_489_v31_)[index79_] = d_490_v32_
        index80_ = default__.safeIndex(560, (d_489_v31_).length(0))
        (d_489_v31_)[index80_] = d_490_v32_
        d_488_v30_ = (d_488_v30_) | (d_488_v30_)
        (globalState).f17 = (self).f23
        d_491_v33_: D0
        d_491_v33_ = D0_DC1(p2, p3)
        r0 = d_491_v33_
        r1 = p3
        d_492_v34_: D0
        d_492_v34_ = D0_DC0(p0)
        pat_let_tv4_ = d_451_v0_
        def lambda21_(source4_):
            if source4_.is_DC1:
                d_493___mcc_h0_ = source4_.cf1
                d_494___mcc_h1_ = source4_.cf2
                d_495_cf2_ = d_494___mcc_h1_
                d_496_cf1_ = d_493___mcc_h0_
                return len(pat_let_tv4_)
            elif source4_.is_DC0:
                d_497___mcc_h2_ = source4_.cf0
                d_498_cf0_ = d_497___mcc_h2_
                return len(_dafny.SeqWithoutIsStrInference([-577]))
            elif True:
                d_499___mcc_h3_ = source4_.cf3
                d_500_cf3_ = d_499___mcc_h3_
                return (self).f23

        r2 = lambda21_(D0_DC2(d_492_v34_))
        return r0, r1, r2

    def m7(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        (globalState).f14 = (self).f23
        index81_ = default__.safeIndex(901, (p0).length(0))
        (p0)[index81_] = p2
        d_501_v0_: _dafny.Seq
        d_501_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qiu"))
        index82_ = default__.safeIndex(901, (p0).length(0))
        (p0)[index82_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mseqsel"))) != (d_501_v0_) if True else p1)
        (globalState).f17 = (self).f23
        d_502_v1_: _dafny.MultiSet
        d_502_v1_ = _dafny.MultiSet([p1, p1, p1, p2])
        d_503_v2_: _dafny.MultiSet
        d_503_v2_ = _dafny.MultiSet([(self).f23])
        d_504_v3_: _dafny.Seq
        d_504_v3_ = _dafny.SeqWithoutIsStrInference([p2, p2, p1])
        d_505_v4_: _dafny.Map
        d_505_v4_ = _dafny.Map({(d_503_v2_).cardinality: p2})
        d_506_v5_: _dafny.Map
        d_506_v5_ = _dafny.Map({len(d_505_v4_): d_501_v0_})
        d_507_v6_: _dafny.Map
        d_507_v6_ = _dafny.Map({((d_506_v5_)[(self).f23] if ((self).f23) in (d_506_v5_) else d_501_v0_): d_504_v3_})
        d_508_v7_: str
        d_508_v7_ = _dafny.CodePoint('j')
        d_509_v8_: _dafny.Set
        d_509_v8_ = _dafny.Set({d_508_v7_})
        d_510_v9_: _dafny.Array
        def lambda22_(d_511_p1_):
            def lambda23_(d_512_i1_):
                return D0_DC2(D0_DC1(d_511_p1_, (self).f23))

            return lambda23_

        init13_ = lambda22_(p1)
        nw95_ = _dafny.Array(None, 20)
        for i0_13_ in range(nw95_.length(0)):
            nw95_[i0_13_] = init13_(i0_13_)
        d_510_v9_ = nw95_
        d_513_v10_: _dafny.Map
        d_513_v10_ = _dafny.Map({d_510_v9_: d_508_v7_})
        d_514_v11_: D3
        d_514_v11_ = D3_DC9((self).f23)
        d_515_v12_: _dafny.Array
        def lambda24_(d_516_i2_):
            return default__.safeDivisionInt(d_516_i2_, (self).f23)

        init14_ = lambda24_
        nw96_ = _dafny.Array(None, 9)
        for i0_14_ in range(nw96_.length(0)):
            nw96_[i0_14_] = init14_(i0_14_)
        d_515_v12_ = nw96_
        d_517_v13_: D5
        d_517_v13_ = D5_DC15(len((d_501_v0_).set(default__.safeIndex(len(d_509_v8_), len(d_501_v0_)), d_508_v7_)), d_508_v7_, d_513_v10_, _dafny.Map({d_514_v11_: d_515_v12_}))
        d_518_v14_: _dafny.Array
        nw97_ = _dafny.Array(None, 15)
        nw97_[int(0)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_519_i0_ in range(default__.abs(223))]))
        nw97_[int(1)] = (self).f23
        nw97_[int(2)] = (self).f23
        nw97_[int(3)] = ((d_502_v1_)[p2] if (p2) in (d_502_v1_) else (self).f23)
        nw97_[int(4)] = (d_503_v2_).cardinality
        nw97_[int(5)] = (default__.fm21(True, p2, (self).f23, globalState)).cardinality
        nw97_[int(6)] = (self).f23
        nw97_[int(7)] = (0) - (len(default__.fm20((p0)[default__.safeIndex(901, (p0).length(0))], globalState)))
        nw97_[int(8)] = (len(d_501_v0_)) * ((self).f23)
        nw97_[int(9)] = (self).f23
        nw97_[int(10)] = (self).f23
        nw97_[int(11)] = (self).fm2(globalState)
        nw97_[int(12)] = len((d_504_v3_) + (((d_507_v6_)[d_501_v0_] if (d_501_v0_) in (d_507_v6_) else d_504_v3_)))
        nw97_[int(13)] = (d_517_v13_).cf25
        nw97_[int(14)] = (self).f23
        d_518_v14_ = nw97_
        index83_ = default__.safeIndex(726, (d_515_v12_).length(0))
        (d_515_v12_)[index83_] = default__.safeDivisionInt((self).f23, (self).f23)
        d_520_v15_: _dafny.Map
        d_520_v15_ = _dafny.Map({(self).fm2(globalState): p0})
        d_521_v16_: _dafny.Set
        d_521_v16_ = _dafny.Set({((d_520_v15_)[(self).f23] if ((self).f23) in (d_520_v15_) else p0)})
        index84_ = default__.safeIndex(726, (d_515_v12_).length(0))
        (d_515_v12_)[index84_] = (0) - ((len((_dafny.Set({p0}) if p3 else d_521_v16_))) - ((self).f23))
        d_522_i3_: int
        d_522_i3_ = 0
        with _dafny.label("3"):
            while (d_508_v7_) not in (d_501_v0_):
                with _dafny.c_label("3"):
                    if (d_522_i3_) >= (100):
                        raise _dafny.Break("3")
                    d_522_i3_ = (d_522_i3_) + (1)
                    if not (p2) or (p2):
                        d_523_v17_: C1
                        nw98_ = C1()
                        nw98_.ctor__(p3, False)
                        d_523_v17_ = nw98_
                        d_524_v18_: C0
                        nw99_ = C0()
                        nw99_.ctor__((d_521_v16_) - (d_521_v16_), (d_515_v12_)[default__.safeIndex(726, (d_515_v12_).length(0))], 262)
                        d_524_v18_ = nw99_
                        d_525_v19_: _dafny.Map
                        d_525_v19_ = _dafny.Map({p1: p3})
                        (globalState).f1 = (len((_dafny.Map({default__.fm29(d_524_v18_.f31, (self).f23, (p0)[default__.safeIndex(901, (p0).length(0))], globalState): p3})) | (d_525_v19_))) <= ((d_515_v12_)[default__.safeIndex(726, (d_515_v12_).length(0))])
                        d_526_v20_: _dafny.Array
                        nw100_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
                        d_526_v20_ = nw100_
                        index85_ = default__.safeIndex(595, (d_526_v20_).length(0))
                        (d_526_v20_)[index85_] = (d_501_v0_) + (d_501_v0_)
                        index86_ = default__.safeIndex(595, (d_526_v20_).length(0))
                        (d_526_v20_)[index86_] = (d_501_v0_) + (default__.fm20((p0)[default__.safeIndex(901, (p0).length(0))], globalState))
                        d_527_v21_: _dafny.Seq
                        d_527_v21_ = _dafny.SeqWithoutIsStrInference([d_524_v18_.f31, 715])
                        d_528_v22_: _dafny.Set
                        d_528_v22_ = _dafny.Set({(p0)[default__.safeIndex(901, (p0).length(0))]})
                        d_529_v23_: _dafny.Map
                        d_529_v23_ = _dafny.Map({(d_527_v21_).set(default__.safeIndex((self).f23, len(d_527_v21_)), (d_515_v12_)[default__.safeIndex(726, (d_515_v12_).length(0))]): (d_528_v22_).isdisjoint(d_528_v22_)})
                        d_530_v25_: _dafny.Seq
                        d_530_v25_ = _dafny.SeqWithoutIsStrInference([d_527_v21_, d_527_v21_])
                        def iife27_():
                            coll17_ = _dafny.Map()
                            compr_17_: _dafny.Seq
                            for compr_17_ in (d_530_v25_).Elements:
                                d_531_v24_: _dafny.Seq = compr_17_
                                if (d_531_v24_) in (d_530_v25_):
                                    coll17_[d_531_v24_] = (p0)[default__.safeIndex(901, (p0).length(0))]
                            return _dafny.Map(coll17_)
                        def iife28_():
                            coll18_ = _dafny.Map()
                            compr_18_: _dafny.Seq
                            for compr_18_ in (d_529_v23_).keys.Elements:
                                d_532_v26_: _dafny.Seq = compr_18_
                                if (d_532_v26_) in (d_529_v23_):
                                    coll18_[d_532_v26_] = p1
                            return _dafny.Map(coll18_)
                        d_529_v23_ = ((iife27_()
                        ) | (iife28_()
                        )) | (_dafny.Map({d_527_v21_: p1}))
                    elif True:
                        d_533_v27_: _dafny.Map
                        d_533_v27_ = _dafny.Map({(0) - (default__.safeModuloInt((d_515_v12_)[default__.safeIndex(726, (d_515_v12_).length(0))], -120)): (_dafny.MultiSet([(p0)[default__.safeIndex(901, (p0).length(0))], p3, p1, True, p3])).cardinality})
                        d_533_v27_ = (_dafny.Map({len(_dafny.Map({((d_503_v2_)[(self).f23] if ((self).f23) in (d_503_v2_) else -265): _dafny.CodePoint('p')})): default__.fm1(default__.fm1((_dafny.MultiSet(d_504_v3_)).cardinality, p2, globalState), p2, globalState)})) | (_dafny.Map({(self).f23: (self).f23}))
                        r0 = not((d_508_v7_) not in ((d_501_v0_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_534_i4_ in range(default__.abs(511))]))))
                        (globalState).f17 = default__.safeDivisionInt(((self).f23) + ((d_515_v12_)[default__.safeIndex(726, (d_515_v12_).length(0))]), (self).f23)
                        (globalState).f17 = 368
                        d_535_v30_: _dafny.Array
                        def lambda25_(d_536_v12_, d_537_v7_, d_538_v4_):
                            def lambda26_(d_539_i5_):
                                def iife29_():
                                    def iife31_():
                                        coll21_ = _dafny.Set()
                                        compr_21_: int
                                        for compr_21_ in (d_538_v4_).keys.Elements:
                                            d_542_v29_: int = compr_21_
                                            if (d_542_v29_) in (d_538_v4_):
                                                coll21_ = coll21_.union(_dafny.Set([default__.safeModuloInt(d_542_v29_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnqjr"))))]))
                                        return _dafny.Set(coll21_)
                                    coll19_ = _dafny.Map()
                                    def iife30_():
                                        coll20_ = _dafny.Set()
                                        compr_20_: int
                                        for compr_20_ in (d_538_v4_).keys.Elements:
                                            d_540_v29_: int = compr_20_
                                            if (d_540_v29_) in (d_538_v4_):
                                                coll20_ = coll20_.union(_dafny.Set([default__.safeModuloInt(d_540_v29_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnqjr"))))]))
                                        return _dafny.Set(coll20_)
                                    compr_19_: int
                                    for compr_19_ in (iife30_()
                                    ).Elements:
                                        d_541_v28_: int = compr_19_
                                        if (d_541_v28_) in (iife31_()
                                        ):
                                            coll19_[default__.safeDivisionInt(d_541_v28_, (d_536_v12_)[default__.safeIndex(726, (d_536_v12_).length(0))])] = len(_dafny.Map({d_537_v7_: d_537_v7_}))
                                    return _dafny.Map(coll19_)
                                return iife29_()
                                

                            return lambda26_

                        init15_ = lambda25_(d_515_v12_, d_508_v7_, d_505_v4_)
                        nw101_ = _dafny.Array(None, 18)
                        for i0_15_ in range(nw101_.length(0)):
                            nw101_[i0_15_] = init15_(i0_15_)
                        d_535_v30_ = nw101_
                        def lambda27_(d_543_v27_):
                            def lambda28_(d_544_i6_):
                                return d_543_v27_

                            return lambda28_

                        init16_ = lambda27_(d_533_v27_)
                        nw102_ = _dafny.Array(None, 24)
                        for i0_16_ in range(nw102_.length(0)):
                            nw102_[i0_16_] = init16_(i0_16_)
                        d_535_v30_ = nw102_
                    d_545_v31_: _dafny.Map
                    d_545_v31_ = _dafny.Map({p3: (p0)[default__.safeIndex(901, (p0).length(0))]})
                    d_546_v32_: _dafny.Set
                    d_546_v32_ = _dafny.Set({(self).f23})
                    d_547_v33_: _dafny.Map
                    d_547_v33_ = _dafny.Map({False: len(d_546_v32_)})
                    d_548_v34_: D1
                    d_548_v34_ = D1_DC5((p0)[default__.safeIndex(901, (p0).length(0))], d_545_v31_, -897, len((default__.fm30(((d_502_v1_)[p2] if (p2) in (d_502_v1_) else -629), d_501_v0_, p3, ((d_502_v1_)[True] if (True) in (d_502_v1_) else ((d_547_v33_)[(p0)[default__.safeIndex(901, (p0).length(0))]] if ((p0)[default__.safeIndex(901, (p0).length(0))]) in (d_547_v33_) else len(d_501_v0_))), globalState)).set(default__.safeIndex((d_515_v12_)[default__.safeIndex(726, (d_515_v12_).length(0))], len(default__.fm30(((d_502_v1_)[p2] if (p2) in (d_502_v1_) else -629), d_501_v0_, p3, ((d_502_v1_)[True] if (True) in (d_502_v1_) else ((d_547_v33_)[(p0)[default__.safeIndex(901, (p0).length(0))]] if ((p0)[default__.safeIndex(901, (p0).length(0))]) in (d_547_v33_) else len(d_501_v0_))), globalState))), (p0)[default__.safeIndex(901, (p0).length(0))])), (d_547_v33_) | (d_547_v33_))
                    d_548_v34_ = d_548_v34_
                    (globalState).f8 = (d_515_v12_)[default__.safeIndex(726, (d_515_v12_).length(0))]
                    (globalState).f8 = (self).f23
                    pass
            pass
        d_549_v35_: _dafny.Seq
        d_549_v35_ = _dafny.SeqWithoutIsStrInference([d_503_v2_])
        d_550_v36_: _dafny.Seq
        d_550_v36_ = _dafny.SeqWithoutIsStrInference([d_549_v35_, d_549_v35_, _dafny.SeqWithoutIsStrInference([d_503_v2_ for d_551_i7_ in range(default__.abs(-187))])])
        (globalState).f9 = (d_549_v35_) < ((d_550_v36_)[default__.safeIndex((d_515_v12_)[default__.safeIndex(726, (d_515_v12_).length(0))], len(d_550_v36_))])
        d_552_v37_: C0
        nw103_ = C0()
        nw103_.ctor__(d_521_v16_, (self).f23, (0) - ((d_515_v12_)[default__.safeIndex(726, (d_515_v12_).length(0))]))
        d_552_v37_ = nw103_
        d_553_v38_: _dafny.Seq
        d_553_v38_ = _dafny.SeqWithoutIsStrInference([d_552_v37_, d_552_v37_])
        r0 = (p3) in ((_dafny.Map({p3: d_553_v38_})).set((p0)[default__.safeIndex(901, (p0).length(0))], d_553_v38_))
        return r0


class C3(T1, T0):
    def  __init__(self):
        self._f24: bool = False
        self._f23: int = int(0)
        self._f25: bool = False
        self._f28: bool = False
        self._f29: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    @property
    def f23(self):
        return self._f23
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f28, f29, f24, f25, f23):
        (self)._f28 = f28
        (self)._f29 = f29
        (self).f24 = f24
        (self)._f25 = f25
        (self)._f23 = f23

    def fm4(self, p0, p1, p2, globalState):
        return _dafny.Map({(self).f28: (_dafny.Map({(self).f25: (self).f28})) | (_dafny.Map({(self).f28: (self).f28}))})

    def fm5(self, p0, p1, globalState):
        return ((0) - ((self).f29)) * ((self).f29)

    def fm2(self, globalState):
        return (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), _dafny.CodePoint('u'), _dafny.CodePoint('m')]))) + ((self).f29)

    def fm3(self, p0, p1, p2, p3, globalState):
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njpng"))}) if (self).f28 else _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arqurnxwi"))}))).ispropersubset((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfvvyx")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htve"))})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exnwakoad")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfhvlkukr"))})))

    def fm16(self, p0, p1, p2, globalState):
        return not(not (not(((self).f29) != ((self).f23))) or ((len(_dafny.SeqWithoutIsStrInference([(self).f25]))) <= ((self).f23)))

    def fm17(self, p0, globalState):
        return (self).f25

    def m1(self, globalState):
        r0: bool = False
        r1: T0 = None
        r2: bool = False
        r3: int = int(0)
        d_554_v0_: _dafny.Array
        nw104_ = _dafny.Array(_dafny.Set({}), 8)
        d_554_v0_ = nw104_
        d_555_v1_: str
        d_555_v1_ = _dafny.CodePoint('d')
        d_556_v2_: D1
        d_556_v2_ = D1_DC3(d_555_v1_)
        d_557_v3_: D1
        d_557_v3_ = D1_DC6(d_556_v2_)
        d_558_v4_: _dafny.Array
        nw105_ = _dafny.Array(None, 6)
        nw105_[int(0)] = (self).f25
        nw105_[int(1)] = self.f24
        nw105_[int(2)] = (self).f28
        nw105_[int(3)] = (self).fm17(d_557_v3_, globalState)
        nw105_[int(4)] = (self).f25
        nw105_[int(5)] = self.f24
        d_558_v4_ = nw105_
        d_559_v5_: _dafny.Set
        d_559_v5_ = _dafny.Set({d_558_v4_})
        index87_ = default__.safeIndex(581, (d_554_v0_).length(0))
        (d_554_v0_)[index87_] = d_559_v5_
        d_560_v6_: _dafny.Seq
        d_560_v6_ = _dafny.SeqWithoutIsStrInference([d_559_v5_])
        d_561_v7_: _dafny.Map
        d_561_v7_ = _dafny.Map({(self).f23: d_559_v5_})
        index88_ = default__.safeIndex(581, (d_554_v0_).length(0))
        (d_554_v0_)[index88_] = ((d_560_v6_)[default__.safeIndex((self).f23, len(d_560_v6_))]) - (((d_561_v7_)[20] if (20) in (d_561_v7_) else d_559_v5_))
        d_562_v8_: _dafny.MultiSet
        d_562_v8_ = _dafny.MultiSet([(self).f25, ((self).f28) or (True)])
        d_562_v8_ = d_562_v8_
        (globalState).f15 = (self).f28
        d_563_v9_: _dafny.Seq
        d_563_v9_ = _dafny.SeqWithoutIsStrInference([(self).f28])
        d_564_v10_: _dafny.Seq
        d_564_v10_ = _dafny.SeqWithoutIsStrInference([(self).f23])
        d_565_v11_: _dafny.Seq
        d_565_v11_ = _dafny.SeqWithoutIsStrInference([d_564_v10_])
        d_566_v12_: _dafny.MultiSet
        d_566_v12_ = _dafny.MultiSet([len(d_565_v11_), (self).f29])
        d_567_v13_: _dafny.Map
        d_567_v13_ = _dafny.Map({(self).f28: ((d_566_v12_)[197] if (197) in (d_566_v12_) else 410)})
        d_568_v14_: D1
        d_568_v14_ = D1_DC5(self.f24, _dafny.Map({self.f24: self.f24}), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhhqehb"))), len(d_563_v9_), d_567_v13_)
        d_569_v15_: _dafny.Seq
        d_569_v15_ = _dafny.SeqWithoutIsStrInference([(d_568_v14_).cf10, (self).f23])
        (globalState).f8 = (d_564_v10_)[default__.safeIndex((0) - (((self).f29) * ((self).f29)), len(d_564_v10_))]
        d_570_v16_: _dafny.Seq
        d_570_v16_ = _dafny.SeqWithoutIsStrInference([d_568_v14_])
        r3 = ((len(d_570_v16_) if True else (self).f23) if (self).f25 else (self).f23)
        d_571_v17_: _dafny.Array
        nw106_ = _dafny.Array(_dafny.Array(None, 0), 15)
        d_571_v17_ = nw106_
        index89_ = default__.safeIndex(346, (d_571_v17_).length(0))
        (d_571_v17_)[index89_] = d_558_v4_
        index90_ = default__.safeIndex(90, (d_558_v4_).length(0))
        (d_558_v4_)[index90_] = ((self).f25) or ((d_563_v9_)[default__.safeIndex((self).f23, len(d_563_v9_))])
        d_572_v18_: D3
        d_572_v18_ = D3_DC9((self).f23)
        d_573_v19_: _dafny.Array
        nw107_ = _dafny.Array(None, 9)
        nw107_[int(0)] = len(default__.fm18((self).f23, (self).f29, d_563_v9_, d_572_v18_, globalState))
        nw107_[int(1)] = ((self).f23) - ((self).f23)
        nw107_[int(2)] = (self).f29
        nw107_[int(3)] = (self).f23
        nw107_[int(4)] = (default__.fm1((self).f29, True, globalState)) - ((self).f29)
        nw107_[int(5)] = (self).f29
        nw107_[int(6)] = len(_dafny.SeqWithoutIsStrInference([(self).f28, (self).f25]))
        nw107_[int(7)] = (self).f29
        nw107_[int(8)] = (self).f29
        d_573_v19_ = nw107_
        index91_ = default__.safeIndex(52, (d_573_v19_).length(0))
        (d_573_v19_)[index91_] = (self).f23
        d_574_v20_: _dafny.Set
        d_574_v20_ = _dafny.Set({181, (self).f23})
        d_575_v21_: _dafny.Set
        d_575_v21_ = _dafny.Set({d_574_v20_})
        index92_ = default__.safeIndex(346, (d_571_v17_).length(0))
        index93_ = default__.safeIndex(90, (d_558_v4_).length(0))
        index94_ = default__.safeIndex(52, (d_573_v19_).length(0))
        rhs109_ = d_558_v4_
        rhs110_ = (self).fm16((self).f25, (self).f28, self.f24, globalState)
        rhs111_ = default__.safeModuloInt(default__.fm1((0) - ((self).f23), self.f24, globalState), ((self).f29) - (len(d_575_v21_)))
        rhs112_ = (0) - ((self).f29)
        lhs87_ = d_571_v17_
        lhs88_ = default__.safeIndex(346, (d_571_v17_).length(0))
        lhs89_ = d_558_v4_
        lhs90_ = default__.safeIndex(90, (d_558_v4_).length(0))
        lhs91_ = d_573_v19_
        lhs92_ = default__.safeIndex(52, (d_573_v19_).length(0))
        lhs93_ = globalState
        lhs87_[lhs88_] = rhs109_
        lhs89_[lhs90_] = rhs110_
        lhs91_[lhs92_] = rhs111_
        lhs93_.f8 = rhs112_
        d_576_v22_: _dafny.Map
        d_576_v22_ = _dafny.Map({d_555_v1_: _dafny.CodePoint('h')})
        pat_let_tv5_ = d_563_v9_
        pat_let_tv6_ = d_573_v19_
        pat_let_tv7_ = d_573_v19_
        pat_let_tv8_ = d_563_v9_
        pat_let_tv9_ = d_558_v4_
        pat_let_tv10_ = d_558_v4_
        def lambda29_(source5_):
            if source5_.is_DC4:
                d_577___mcc_h0_ = source5_.cf5
                d_578___mcc_h1_ = source5_.cf6
                d_579_cf6_ = d_578___mcc_h1_
                d_580_cf5_ = d_577___mcc_h0_
                return d_579_cf6_
            elif source5_.is_DC5:
                d_581___mcc_h2_ = source5_.cf7
                d_582___mcc_h3_ = source5_.cf8
                d_583___mcc_h4_ = source5_.cf9
                d_584___mcc_h5_ = source5_.cf10
                d_585___mcc_h6_ = source5_.cf11
                d_586_cf11_ = d_585___mcc_h6_
                d_587_cf10_ = d_584___mcc_h5_
                d_588_cf9_ = d_583___mcc_h4_
                d_589_cf8_ = d_582___mcc_h3_
                d_590_cf7_ = d_581___mcc_h2_
                return d_590_cf7_
            elif source5_.is_DC3:
                d_591___mcc_h7_ = source5_.cf4
                d_592_cf4_ = d_591___mcc_h7_
                return (_dafny.Map({(pat_let_tv5_)[default__.safeIndex((pat_let_tv7_)[default__.safeIndex(52, (pat_let_tv6_).length(0))], len(pat_let_tv8_))]: (self).f28})) == (_dafny.Map({(self).f28: (pat_let_tv10_)[default__.safeIndex(90, (pat_let_tv9_).length(0))]}))
            elif True:
                d_593___mcc_h8_ = source5_.cf12
                d_594_cf12_ = d_593___mcc_h8_
                return (self).f28

        r0 = lambda29_(D1_DC3(((d_576_v22_)[d_555_v1_] if (d_555_v1_) in (d_576_v22_) else _dafny.CodePoint('k'))))
        d_595_v23_: T0
        nw108_ = C2()
        nw108_.ctor__((0) - ((self).f29))
        d_595_v23_ = nw108_
        r1 = d_595_v23_
        r2 = not ((self).f25) or ((self).f28)
        r3 = (self).f23
        return r0, r1, r2, r3

    def m0(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: bool = False
        d_596_v0_: _dafny.Set
        d_596_v0_ = _dafny.Set({(self).f29})
        d_597_v1_: D8
        d_597_v1_ = D8_DC23((self).f23, d_596_v0_, (self).f23, (0) - (default__.fm1((self).f29, p1, globalState)), 850)
        d_598_v2_: _dafny.Map
        d_598_v2_ = _dafny.Map({d_597_v1_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))})
        d_599_v3_: _dafny.Seq
        d_599_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnilurti"))
        d_598_v2_ = (d_598_v2_).set((d_597_v1_ if p0 else d_597_v1_), d_599_v3_)
        d_600_v4_: _dafny.Seq
        d_600_v4_ = _dafny.SeqWithoutIsStrInference([default__.fm31(p0, (self).f28, globalState)])
        d_601_v5_: _dafny.Seq
        d_601_v5_ = _dafny.SeqWithoutIsStrInference([692])
        d_602_i0_: int
        d_602_i0_ = 0
        with _dafny.label("4"):
            while (((d_600_v4_)[default__.safeIndex((self).f29, len(d_600_v4_))]).set(default__.safeIndex((0) - ((self).f29), len((d_600_v4_)[default__.safeIndex((self).f29, len(d_600_v4_))])), (0) - ((self).f29))) == (d_601_v5_):
                with _dafny.c_label("4"):
                    if (d_602_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_602_i0_ = (d_602_i0_) + (1)
                    d_603_v6_: _dafny.Array
                    nw109_ = _dafny.Array(False, 24)
                    d_603_v6_ = nw109_
                    d_604_v7_: _dafny.Set
                    d_604_v7_ = _dafny.Set({d_603_v6_, d_603_v6_, d_603_v6_, d_603_v6_})
                    d_605_v8_: C0
                    nw110_ = C0()
                    nw110_.ctor__(d_604_v7_, (self).f23, default__.safeDivisionInt((self).f23, (0) - ((self).f23)))
                    d_605_v8_ = nw110_
                    d_606_v9_: _dafny.Array
                    nw111_ = _dafny.Array(int(0), 1)
                    d_606_v9_ = nw111_
                    d_607_v10_: _dafny.Set
                    d_607_v10_ = _dafny.Set({d_606_v9_})
                    d_608_v11_: _dafny.Set
                    d_608_v11_ = _dafny.Set({(self).f28})
                    d_609_v12_: _dafny.Map
                    d_609_v12_ = _dafny.Map({len(_dafny.Set({(self).f23, (self).f29, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_610_i1_ in range(default__.abs(845))])), len(d_607_v10_), len(d_608_v11_)})): (self).f23})
                    d_611_v13_: C1
                    nw112_ = C1()
                    nw112_.ctor__(False, (d_609_v12_) == (d_609_v12_))
                    d_611_v13_ = nw112_
                    d_612_v14_: _dafny.Seq
                    d_612_v14_ = _dafny.SeqWithoutIsStrInference([self.f24, p1])
                    index95_ = default__.safeIndex(550, (d_603_v6_).length(0))
                    (d_603_v6_)[index95_] = (d_612_v14_) <= (d_612_v14_)
                    index96_ = default__.safeIndex(550, (d_603_v6_).length(0))
                    (d_603_v6_)[index96_] = True
                    d_613_v15_: _dafny.Map
                    d_613_v15_ = _dafny.Map({d_599_v3_: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_614_i2_ in range(default__.abs(848))]))})
                    (globalState).f6 = not((d_599_v3_) not in (d_613_v15_))
                    pass
            pass
        d_615_v16_: _dafny.MultiSet
        d_615_v16_ = _dafny.MultiSet([(self).f23])
        d_616_v17_: _dafny.MultiSet
        d_616_v17_ = _dafny.MultiSet([not(p1)])
        d_617_v18_: _dafny.Map
        d_617_v18_ = _dafny.Map({(d_616_v17_).cardinality: d_615_v16_})
        d_618_v19_: D4
        d_618_v19_ = D4_DC12(((d_615_v16_) - (((d_617_v18_)[(self).f29] if ((self).f29) in (d_617_v18_) else d_615_v16_))).cardinality, ((self).f23 if (self).f28 else (self).f23))
        d_618_v19_ = D4_DC12((self).f29, (0) - ((self).f23))
        hi4_ = -402
        for d_619_i3_ in range((self).f29, hi4_):
            d_620_v20_: _dafny.Set
            d_620_v20_ = _dafny.Set({p0})
            d_621_v21_: _dafny.Seq
            d_621_v21_ = _dafny.SeqWithoutIsStrInference([p1])
            d_622_v22_: _dafny.Seq
            d_622_v22_ = _dafny.SeqWithoutIsStrInference([d_621_v21_, d_621_v21_, d_621_v21_, d_621_v21_])
            d_623_v23_: D1
            d_623_v23_ = D1_DC4(p1, True)
            d_624_v24_: _dafny.Seq
            d_624_v24_ = _dafny.SeqWithoutIsStrInference([d_623_v23_, d_623_v23_, d_623_v23_, d_623_v23_, d_623_v23_])
            d_625_v26_: _dafny.Array
            def lambda30_(d_626_p1_):
                def lambda31_(d_627_i4_):
                    return d_626_p1_

                return lambda31_

            init17_ = lambda30_(p1)
            nw113_ = _dafny.Array(None, 15)
            for i0_17_ in range(nw113_.length(0)):
                nw113_[i0_17_] = init17_(i0_17_)
            d_625_v26_ = nw113_
            d_628_v27_: _dafny.Map
            d_628_v27_ = _dafny.Map({d_625_v26_: (self).f28})
            d_629_v28_: _dafny.Map
            d_629_v28_ = _dafny.Map({self.f24: d_625_v26_})
            d_630_v29_: _dafny.Set
            d_630_v29_ = _dafny.Set({d_623_v23_})
            d_631_v31_: _dafny.Map
            d_631_v31_ = _dafny.Map({-446: d_620_v20_})
            d_632_v32_: _dafny.Seq
            d_632_v32_ = _dafny.SeqWithoutIsStrInference([d_620_v20_, d_620_v20_, d_620_v20_, d_620_v20_, d_620_v20_])
            d_633_v33_: _dafny.Map
            d_633_v33_ = _dafny.Map({(self).f28: p1})
            d_634_v34_: D1
            d_634_v34_ = D1_DC6(D1_DC5((self).f28, d_633_v33_, d_619_i3_, (self).f23, _dafny.Map({p0: (0) - (d_619_i3_)})))
            d_635_v35_: _dafny.Array
            nw114_ = _dafny.Array(None, 22)
            nw114_[int(0)] = d_620_v20_
            nw114_[int(1)] = d_620_v20_
            nw114_[int(2)] = _dafny.Set({True})
            nw114_[int(3)] = d_620_v20_
            nw114_[int(4)] = d_620_v20_
            def iife32_():
                coll22_ = _dafny.Set()
                compr_22_: D1
                for compr_22_ in (d_624_v24_).Elements:
                    d_636_v25_: D1 = compr_22_
                    if (d_636_v25_) in (d_624_v24_):
                        coll22_ = coll22_.union(_dafny.Set([d_636_v25_]))
                return _dafny.Set(coll22_)
            nw114_[int(5)] = default__.fm32(self.f24, d_601_v5_, (d_622_v22_)[default__.safeIndex((self).f29, len(d_622_v22_))], iife32_()
            , globalState)
            nw114_[int(6)] = d_620_v20_
            nw114_[int(7)] = d_620_v20_
            nw114_[int(8)] = default__.fm32(((d_628_v27_)[((d_629_v28_)[not(p0)] if (not(p0)) in (d_629_v28_) else d_625_v26_)] if (((d_629_v28_)[not(p0)] if (not(p0)) in (d_629_v28_) else d_625_v26_)) in (d_628_v27_) else p1), d_601_v5_, d_621_v21_, d_630_v29_, globalState)
            def iife33_():
                coll23_ = _dafny.Set()
                compr_23_: D1
                for compr_23_ in (default__.fm33((self).f28, d_599_v3_, 846, (self).f28, globalState)).Elements:
                    d_637_v30_: D1 = compr_23_
                    if (d_637_v30_) in (default__.fm33((self).f28, d_599_v3_, 846, (self).f28, globalState)):
                        coll23_ = coll23_.union(_dafny.Set([d_637_v30_]))
                return _dafny.Set(coll23_)
            nw114_[int(9)] = default__.fm32((self).f28, d_601_v5_, d_621_v21_, iife33_()
            , globalState)
            nw114_[int(10)] = d_620_v20_
            nw114_[int(11)] = ((d_631_v31_)[len(_dafny.Map({(d_621_v21_)[default__.safeIndex(len(d_599_v3_), len(d_621_v21_))]: (self).f28}))] if (len(_dafny.Map({(d_621_v21_)[default__.safeIndex(len(d_599_v3_), len(d_621_v21_))]: (self).f28}))) in (d_631_v31_) else d_620_v20_)
            nw114_[int(12)] = ((d_632_v32_)[default__.safeIndex(len(d_633_v33_), len(d_632_v32_))]).intersection(d_620_v20_)
            nw114_[int(13)] = _dafny.Set({self.f24, self.f24, (d_621_v21_)[default__.safeIndex((0) - ((self).f29), len(d_621_v21_))]})
            nw114_[int(14)] = (d_620_v20_ if True else _dafny.Set({self.f24}))
            nw114_[int(15)] = _dafny.Set({not(p1)})
            nw114_[int(16)] = d_620_v20_
            nw114_[int(17)] = d_620_v20_
            nw114_[int(18)] = d_620_v20_
            nw114_[int(19)] = _dafny.Set({not((self).fm17(d_634_v34_, globalState)), (self).f25, not(p0)})
            nw114_[int(20)] = d_620_v20_
            nw114_[int(21)] = default__.fm32(self.f24, d_601_v5_, d_621_v21_, d_630_v29_, globalState)
            d_635_v35_ = nw114_
            index97_ = default__.safeIndex(696, (d_635_v35_).length(0))
            (d_635_v35_)[index97_] = d_620_v20_
            index98_ = default__.safeIndex(820, (d_625_v26_).length(0))
            (d_625_v26_)[index98_] = self.f24
            index99_ = default__.safeIndex(696, (d_635_v35_).length(0))
            index100_ = default__.safeIndex(820, (d_625_v26_).length(0))
            rhs113_ = (d_620_v20_) | (d_620_v20_)
            rhs114_ = (default__.safeDivisionInt(len(d_599_v3_), (self).f23)) <= (386)
            rhs115_ = (0) - ((self).f23)
            rhs116_ = (0) - ((0) - (default__.safeModuloInt((self).f23, (self).f23)))
            lhs94_ = d_635_v35_
            lhs95_ = default__.safeIndex(696, (d_635_v35_).length(0))
            lhs96_ = d_625_v26_
            lhs97_ = default__.safeIndex(820, (d_625_v26_).length(0))
            lhs98_ = globalState
            lhs99_ = globalState
            lhs94_[lhs95_] = rhs113_
            lhs96_[lhs97_] = rhs114_
            lhs98_.f17 = rhs115_
            lhs99_.f14 = rhs116_
            (globalState).f5 = ((d_616_v17_)[((d_625_v26_)[default__.safeIndex(820, (d_625_v26_).length(0))]) == (self.f24)] if (((d_625_v26_)[default__.safeIndex(820, (d_625_v26_).length(0))]) == (self.f24)) in (d_616_v17_) else ((self).f29) - (len(d_599_v3_)))
            d_638_v36_: _dafny.MultiSet
            d_638_v36_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_639_i5_ in range(default__.abs(553))])])
            d_640_v37_: C1
            nw115_ = C1()
            nw115_.ctor__((self).f28, (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "riafjk"))])).isdisjoint(d_638_v36_))
            d_640_v37_ = nw115_
            (globalState).f14 = (self).f29
        if p0:
            if (self).fm16(True, (self).f28, self.f24, globalState):
                d_641_v38_: C2
                nw116_ = C2()
                nw116_.ctor__((D10_DC29(self.f24, self.f24, (self).f23)).cf55)
                d_641_v38_ = nw116_
                (globalState).f15 = ((self).f25) and (self.f24)
                d_642_v39_: str
                d_642_v39_ = _dafny.CodePoint('g')
                d_599_v3_ = (d_599_v3_) + (((D3_DC8((d_599_v3_).set(default__.safeIndex((self).f29, len(d_599_v3_)), d_642_v39_))).cf14).set(default__.safeIndex((self).f23, len((D3_DC8((d_599_v3_).set(default__.safeIndex((self).f29, len(d_599_v3_)), d_642_v39_))).cf14)), d_642_v39_))
                d_643_v40_: C1
                nw117_ = C1()
                nw117_.ctor__((self).f25, p1)
                d_643_v40_ = nw117_
                d_644_v41_: _dafny.Array
                nw118_ = _dafny.Array(None, 5)
                nw118_[int(0)] = d_643_v40_
                nw118_[int(1)] = d_643_v40_
                nw118_[int(2)] = d_643_v40_
                nw118_[int(3)] = d_643_v40_
                nw118_[int(4)] = d_643_v40_
                d_644_v41_ = nw118_
                d_645_v42_: _dafny.MultiSet
                d_645_v42_ = _dafny.MultiSet([d_644_v41_])
                d_646_v43_: _dafny.Array
                nw119_ = _dafny.Array(False, 28)
                d_646_v43_ = nw119_
                d_647_v44_: _dafny.Map
                d_647_v44_ = _dafny.Map({d_646_v43_: ((d_616_v17_)[(self).f28] if ((self).f28) in (d_616_v17_) else (self).f29)})
                rhs117_ = ((d_645_v42_).set(d_644_v41_, default__.abs((self).f29))).isdisjoint(d_645_v42_)
                rhs118_ = default__.safeDivisionInt(-779, len((d_647_v44_) | (d_647_v44_)))
                lhs100_ = globalState
                lhs101_ = globalState
                lhs100_.f1 = rhs117_
                lhs101_.f17 = rhs118_
                d_648_v45_: _dafny.Map
                d_648_v45_ = _dafny.Map({(self).f25: True})
                d_649_v46_: _dafny.Map
                d_649_v46_ = _dafny.Map({d_616_v17_: p0})
                d_650_v47_: _dafny.Map
                d_650_v47_ = _dafny.Map({p0: 65})
                d_651_v48_: D1
                d_651_v48_ = D1_DC5(p1, d_648_v45_, -776, len(d_649_v46_), d_650_v47_)
                d_652_v49_: _dafny.Seq
                d_652_v49_ = _dafny.SeqWithoutIsStrInference([(self).f28, self.f24])
                d_653_v50_: _dafny.Array
                nw120_ = _dafny.Array(None, 8)
                nw120_[int(0)] = (self).f23
                nw120_[int(1)] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({858: True})) for d_654_i6_ in range(default__.abs(156))]))
                nw120_[int(2)] = (self).f23
                nw120_[int(3)] = (d_651_v48_).cf10
                nw120_[int(4)] = (self).f23
                nw120_[int(5)] = (self).f29
                nw120_[int(6)] = len(d_652_v49_)
                nw120_[int(7)] = (self).f23
                d_653_v50_ = nw120_
                d_655_v51_: _dafny.Array
                nw121_ = _dafny.Array(None, 14)
                nw121_[int(0)] = d_653_v50_
                nw121_[int(1)] = d_653_v50_
                nw121_[int(2)] = d_653_v50_
                nw121_[int(3)] = d_653_v50_
                nw121_[int(4)] = d_653_v50_
                nw121_[int(5)] = d_653_v50_
                nw121_[int(6)] = d_653_v50_
                nw121_[int(7)] = d_653_v50_
                nw121_[int(8)] = (d_653_v50_ if p1 else d_653_v50_)
                nw121_[int(9)] = d_653_v50_
                nw121_[int(10)] = d_653_v50_
                nw121_[int(11)] = d_653_v50_
                nw121_[int(12)] = d_653_v50_
                nw121_[int(13)] = d_653_v50_
                d_655_v51_ = nw121_
                rhs119_ = (self).f29
                rhs120_ = (self).f23
                rhs121_ = p0
                rhs122_ = _dafny.CodePoint('t')
                rhs123_ = d_655_v51_
                lhs102_ = globalState
                lhs103_ = globalState
                lhs104_ = globalState
                lhs102_.f8 = rhs119_
                lhs103_.f8 = rhs120_
                lhs104_.f9 = rhs121_
                d_642_v39_ = rhs122_
                d_655_v51_ = rhs123_
            elif True:
                d_656_v52_: _dafny.Map
                d_656_v52_ = _dafny.Map({self.f24: True})
                d_657_v53_: _dafny.Map
                d_657_v53_ = _dafny.Map({(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_658_i7_ in range(default__.abs(-899))]))) + (len(d_656_v52_)): (self).f25})
                d_657_v53_ = (d_657_v53_).set((self).f29, p1)
                (globalState).f5 = default__.safeDivisionInt((self).f29, (self).f29)
                d_659_v54_: str
                d_659_v54_ = _dafny.CodePoint('m')
                d_659_v54_ = d_659_v54_
                d_660_v55_: _dafny.Map
                d_660_v55_ = _dafny.Map({(self).f28: (self).f23})
                (globalState).f8 = (((self).f29) * (len(d_660_v55_))) * (((self).f29) - ((self).f29))
                (globalState).f5 = (self).fm2(globalState)
            d_661_v56_: _dafny.Map
            d_661_v56_ = _dafny.Map({len(d_596_v0_): p0})
            d_661_v56_ = d_661_v56_
            d_662_v57_: _dafny.Array
            nw122_ = _dafny.Array(int(0), 13)
            d_662_v57_ = nw122_
            r1 = d_662_v57_
            d_663_v58_: str
            d_663_v58_ = _dafny.CodePoint('n')
            d_663_v58_ = d_663_v58_
            (globalState).f6 = False
        elif True:
            (globalState).f1 = (d_616_v17_).isdisjoint(d_616_v17_)
            d_664_v59_: _dafny.Array
            def lambda32_(d_665_i8_):
                return (self).f25

            init18_ = lambda32_
            nw123_ = _dafny.Array(None, 5)
            for i0_18_ in range(nw123_.length(0)):
                nw123_[i0_18_] = init18_(i0_18_)
            d_664_v59_ = nw123_
            d_664_v59_ = d_664_v59_
            rhs124_ = (d_599_v3_) + (d_599_v3_)
            rhs125_ = (self).f29
            lhs105_ = globalState
            d_599_v3_ = rhs124_
            lhs105_.f5 = rhs125_
            index101_ = default__.safeIndex(763, (d_664_v59_).length(0))
            (d_664_v59_)[index101_] = (self).f28
            index102_ = default__.safeIndex(763, (d_664_v59_).length(0))
            (d_664_v59_)[index102_] = self.f24
            d_666_v60_: T1
            nw124_ = C1()
            nw124_.ctor__(p0, (self).f28)
            d_666_v60_ = nw124_
            d_667_v61_: _dafny.Set
            d_667_v61_ = _dafny.Set({d_666_v60_})
            (self).f24 = (p0) == (not((d_667_v61_) == (d_667_v61_)))
        d_668_v62_: _dafny.Array
        def lambda33_(d_669_i9_):
            return (d_669_i9_) + ((self).f29)

        init19_ = lambda33_
        nw125_ = _dafny.Array(None, 27)
        for i0_19_ in range(nw125_.length(0)):
            nw125_[i0_19_] = init19_(i0_19_)
        d_668_v62_ = nw125_
        index103_ = default__.safeIndex(342, (d_668_v62_).length(0))
        (d_668_v62_)[index103_] = (self).f23
        index104_ = default__.safeIndex(342, (d_668_v62_).length(0))
        (d_668_v62_)[index104_] = (self).f23
        r0 = (self).f29
        r1 = d_668_v62_
        r2 = not(self.f24)
        return r0, r1, r2

    def m5(self, p0, p1, globalState):
        d_670_v0_: _dafny.Array
        def lambda34_(d_671_i1_):
            return self.f24

        init20_ = lambda34_
        nw126_ = _dafny.Array(None, 15)
        for i0_20_ in range(nw126_.length(0)):
            nw126_[i0_20_] = init20_(i0_20_)
        d_670_v0_ = nw126_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_670_v0_).length(0)):
            d_672_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_672_i0_)) and ((d_672_i0_) < ((d_670_v0_).length(0)))):
                (d_670_v0_)[(d_672_i0_)] = ((self).f23) not in (_dafny.SeqWithoutIsStrInference([(0) - ((self).f23)]))
        d_673_v1_: _dafny.Array
        def lambda35_(d_674_i2_):
            return (d_674_i2_) * ((self).f29)

        init21_ = lambda35_
        nw127_ = _dafny.Array(None, 24)
        for i0_21_ in range(nw127_.length(0)):
            nw127_[i0_21_] = init21_(i0_21_)
        d_673_v1_ = nw127_
        d_675_v2_: _dafny.Map
        d_675_v2_ = _dafny.Map({d_673_v1_: d_670_v0_})
        d_675_v2_ = (d_675_v2_).set(d_673_v1_, d_670_v0_)
        d_676_i3_: int
        d_676_i3_ = 0
        with _dafny.label("5"):
            while ((self).f28) and ((self).f25):
                with _dafny.c_label("5"):
                    if (d_676_i3_) >= (100):
                        raise _dafny.Break("5")
                    d_676_i3_ = (d_676_i3_) + (1)
                    d_677_v3_: _dafny.MultiSet
                    d_677_v3_ = _dafny.MultiSet([not(p0)])
                    d_678_v4_: D10
                    d_678_v4_ = D10_DC29(self.f24, (d_677_v3_).isdisjoint(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f25, (self).f25, (self).f25]))), p1)
                    d_679_v5_: _dafny.Seq
                    d_679_v5_ = _dafny.SeqWithoutIsStrInference([self.f24, p0])
                    index105_ = default__.safeIndex(726, (d_670_v0_).length(0))
                    (d_670_v0_)[index105_] = (d_679_v5_) != (_dafny.SeqWithoutIsStrInference([self.f24, (self).f28, self.f24, (self).f25, self.f24]))
                    index106_ = default__.safeIndex(726, (d_670_v0_).length(0))
                    rhs126_ = d_678_v4_
                    rhs127_ = self.f24
                    rhs128_ = (self).f29
                    lhs106_ = d_670_v0_
                    lhs107_ = default__.safeIndex(726, (d_670_v0_).length(0))
                    lhs108_ = globalState
                    d_678_v4_ = rhs126_
                    lhs106_[lhs107_] = rhs127_
                    lhs108_.f14 = rhs128_
                    d_680_v6_: _dafny.Seq
                    d_680_v6_ = _dafny.SeqWithoutIsStrInference([(self).f29])
                    d_681_v7_: D10
                    d_681_v7_ = D10_DC28(d_680_v6_)
                    d_682_v8_: _dafny.Map
                    d_682_v8_ = _dafny.Map({(self).f25: (self).f23})
                    d_683_v9_: _dafny.Map
                    d_683_v9_ = _dafny.Map({(d_681_v7_).cf52: len(d_682_v8_)})
                    d_684_v10_: _dafny.Map
                    d_684_v10_ = _dafny.Map({d_683_v9_: (self).f29})
                    d_685_v11_: _dafny.Seq
                    d_685_v11_ = _dafny.SeqWithoutIsStrInference([(0) - (default__.safeModuloInt((self).f29, (0) - ((self).f29))), 68, len(d_680_v6_), 800, ((d_684_v10_)[d_683_v9_] if (d_683_v9_) in (d_684_v10_) else (self).f29)])
                    d_685_v11_ = _dafny.SeqWithoutIsStrInference([(self).f23 for d_686_i4_ in range(default__.abs(581))])
                    d_687_v12_: C1
                    nw128_ = C1()
                    nw128_.ctor__((self.f24) or (self.f24), p0)
                    d_687_v12_ = nw128_
                    d_688_v13_: _dafny.Seq
                    d_688_v13_ = _dafny.SeqWithoutIsStrInference([d_673_v1_, d_673_v1_])
                    d_689_v14_: _dafny.Map
                    d_689_v14_ = _dafny.Map({True: d_673_v1_})
                    d_690_v15_: _dafny.Seq
                    d_690_v15_ = _dafny.SeqWithoutIsStrInference([((d_688_v13_)[default__.safeIndex((self).f29, len(d_688_v13_))] if (d_670_v0_)[default__.safeIndex(726, (d_670_v0_).length(0))] else d_673_v1_), (d_673_v1_ if (self).f25 else ((d_689_v14_)[self.f24] if (self.f24) in (d_689_v14_) else d_673_v1_)), d_673_v1_, d_673_v1_])
                    d_691_v16_: D0
                    d_691_v16_ = D0_DC0(p0)
                    d_692_v17_: str
                    d_692_v17_ = _dafny.CodePoint('e')
                    d_693_v18_: _dafny.Set
                    d_693_v18_ = _dafny.Set({self.f24, (d_670_v0_)[default__.safeIndex(726, (d_670_v0_).length(0))], (self).f28})
                    rhs129_ = (0) - (len((default__.fm25((d_691_v16_).cf0, (self).f28, d_692_v17_, globalState)) + (_dafny.SeqWithoutIsStrInference([d_692_v17_ for d_694_i5_ in range(default__.abs(24))]))))
                    rhs130_ = (d_693_v18_).ispropersubset(d_693_v18_)
                    rhs131_ = (self).f29
                    rhs132_ = d_688_v13_
                    lhs109_ = globalState
                    lhs110_ = globalState
                    lhs111_ = globalState
                    lhs109_.f17 = rhs129_
                    lhs110_.f6 = rhs130_
                    lhs111_.f8 = rhs131_
                    d_690_v15_ = rhs132_
                    pass
            pass
        d_695_v19_: str
        d_695_v19_ = _dafny.CodePoint('v')
        d_696_v20_: _dafny.Map
        d_696_v20_ = _dafny.Map({(self).f28: d_695_v19_})
        d_697_v21_: _dafny.Seq
        d_697_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvejejg"))
        (globalState).f6 = (len((d_696_v20_) | (d_696_v20_))) != (len(d_697_v21_))
        d_698_v22_: _dafny.Map
        d_698_v22_ = _dafny.Map({p1: (643) - ((self).f23)})
        d_698_v22_ = (d_698_v22_).set(p1, (self).f29)
        d_699_v23_: _dafny.Map
        d_699_v23_ = _dafny.Map({p1: True})
        d_700_v25_: _dafny.Seq
        d_700_v25_ = _dafny.SeqWithoutIsStrInference([(self).f23])
        d_701_v26_: _dafny.Array
        nw129_ = _dafny.Array(None, 8)
        nw129_[int(0)] = (_dafny.Map({(self).f23: self.f24})) | (d_699_v23_)
        nw129_[int(1)] = (d_699_v23_) | (default__.fm34((0) - (len(d_697_v21_)), globalState))
        nw129_[int(2)] = d_699_v23_
        nw129_[int(3)] = d_699_v23_
        nw129_[int(4)] = d_699_v23_
        def iife34_():
            coll24_ = _dafny.Map()
            compr_24_: int
            for compr_24_ in (d_700_v25_).Elements:
                d_702_v24_: int = compr_24_
                if (d_702_v24_) in (d_700_v25_):
                    coll24_[(d_702_v24_) + ((self).f29)] = (self).f25
            return _dafny.Map(coll24_)
        nw129_[int(5)] = iife34_()
        
        nw129_[int(6)] = d_699_v23_
        nw129_[int(7)] = d_699_v23_
        d_701_v26_ = nw129_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_701_v26_).length(0)):
            d_703_i6_: int = guard_loop_1_
            if (True) and (((0) <= (d_703_i6_)) and ((d_703_i6_) < ((d_701_v26_).length(0)))):
                (d_701_v26_)[(d_703_i6_)] = (_dafny.Map({(0) - (-59): self.f24})) | (d_699_v23_)

    @property
    def f28(self):
        return self._f28
    @property
    def f29(self):
        return self._f29

class C4(T1):
    def  __init__(self):
        self._f24: bool = False
        self._f25: bool = False
        self._f27: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f27, f24, f25):
        (self)._f27 = f27
        (self).f24 = f24
        (self)._f25 = f25

    def fm4(self, p0, p1, p2, globalState):
        return ((_dafny.Map({self.f24: _dafny.Map({(self).f25: self.f24})})) | (_dafny.Map({False: _dafny.Map({self.f24: (D0_DC1(not((self).f25), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aeynwfee"))))).cf1})}))) | ((_dafny.Map({(self).f27: _dafny.Map({(self).f25: (self).f27})})) | (_dafny.Map({not(self.f24): _dafny.Map({(self).f25: self.f24})})))

    def fm5(self, p0, p1, globalState):
        def iife35_():
            coll25_ = _dafny.Map()
            compr_25_: int
            for compr_25_ in _dafny.IntegerRange(923, 571):
                d_704_v0_: int = compr_25_
                if ((923) <= (d_704_v0_)) and ((d_704_v0_) < (571)):
                    def iife36_():
                        coll26_ = _dafny.Set()
                        compr_26_: int
                        for compr_26_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({362}))])).Elements:
                            d_705_v1_: int = compr_26_
                            if (d_705_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({362}))])):
                                coll26_ = coll26_.union(_dafny.Set([default__.safeDivisionInt(d_705_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))))]))
                        return _dafny.Set(coll26_)
                    coll25_[(d_704_v0_) - (248)] = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpuhcuieu"))), len(iife36_()
                    ), 220])
            return _dafny.Map(coll25_)
        return len((_dafny.Map({(0) - (-771): _dafny.MultiSet([233])})) | (iife35_()
        ))

    def fm14(self, globalState):
        return (self).f27

    def fm15(self, p0, p1, p2, globalState):
        source6_ = D1_DC4((self).f27, self.f24)
        if source6_.is_DC4:
            d_706___mcc_h0_ = source6_.cf5
            d_707___mcc_h1_ = source6_.cf6
            d_708_cf6_ = d_707___mcc_h1_
            d_709_cf5_ = d_706___mcc_h0_
            return 981
        elif source6_.is_DC5:
            d_710___mcc_h2_ = source6_.cf7
            d_711___mcc_h3_ = source6_.cf8
            d_712___mcc_h4_ = source6_.cf9
            d_713___mcc_h5_ = source6_.cf10
            d_714___mcc_h6_ = source6_.cf11
            d_715_cf11_ = d_714___mcc_h6_
            d_716_cf10_ = d_713___mcc_h5_
            d_717_cf9_ = d_712___mcc_h4_
            d_718_cf8_ = d_711___mcc_h3_
            d_719_cf7_ = d_710___mcc_h2_
            return d_716_cf10_
        elif source6_.is_DC3:
            d_720___mcc_h7_ = source6_.cf4
            d_721_cf4_ = d_720___mcc_h7_
            return 810
        elif True:
            d_722___mcc_h8_ = source6_.cf12
            d_723_cf12_ = d_722___mcc_h8_
            return len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "caeqeee"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_724_i0_ in range(default__.abs(945))])))

    def m1(self, globalState):
        r0: bool = False
        r1: T0 = None
        r2: bool = False
        r3: int = int(0)
        d_725_v0_: int
        d_725_v0_ = 428
        hi5_ = d_725_v0_
        for d_726_i0_ in range((d_725_v0_) * (d_725_v0_), hi5_):
            (globalState).f15 = not ((self).f27) or (self.f24)
            d_727_v1_: _dafny.Array
            nw130_ = _dafny.Array(False, 27)
            d_727_v1_ = nw130_
            index107_ = default__.safeIndex(955, (d_727_v1_).length(0))
            (d_727_v1_)[index107_] = (self).f27
            d_728_v2_: _dafny.Map
            d_728_v2_ = _dafny.Map({not(self.f24): (self).f27})
            d_729_v3_: _dafny.Map
            d_729_v3_ = _dafny.Map({(self).f27: len(d_728_v2_)})
            d_730_v4_: D1
            d_730_v4_ = D1_DC5((self).f27, d_728_v2_, d_726_i0_, 996, d_729_v3_)
            d_731_v5_: D1
            d_731_v5_ = D1_DC6(d_730_v4_)
            pat_let_tv11_ = d_730_v4_
            d_732_v6_: _dafny.MultiSet
            def iife37_(_pat_let5_0):
                def iife38_(d_733_dt__update__tmp_h0_):
                    def iife39_(_pat_let6_0):
                        def iife40_(d_734_dt__update_hcf12_h0_):
                            return D1_DC6(d_734_dt__update_hcf12_h0_)
                        return iife40_(_pat_let6_0)
                    return iife39_(pat_let_tv11_)
                return iife38_(_pat_let5_0)
            d_732_v6_ = _dafny.MultiSet([iife37_(d_731_v5_), d_731_v5_, d_731_v5_])
            index108_ = default__.safeIndex(955, (d_727_v1_).length(0))
            (d_727_v1_)[index108_] = ((d_732_v6_).ispropersubset(d_732_v6_)) or (False)
            d_735_v7_: D1
            d_735_v7_ = D1_DC5(False, (d_728_v2_) | (d_728_v2_), d_725_v0_, d_726_i0_, d_729_v3_)
            d_735_v7_ = d_735_v7_
            d_736_v8_: _dafny.Array
            nw131_ = _dafny.Array(None, 1)
            nw131_[int(0)] = d_727_v1_
            d_736_v8_ = nw131_
            index109_ = default__.safeIndex(275, (d_736_v8_).length(0))
            (d_736_v8_)[index109_] = d_727_v1_
            d_737_v9_: _dafny.Array
            nw132_ = _dafny.Array(_dafny.Map({}), 28)
            d_737_v9_ = nw132_
            d_738_v10_: _dafny.Array
            def lambda36_(d_739_i0_):
                def lambda37_(d_740_i1_):
                    return (d_740_i1_) - (d_739_i0_)

                return lambda37_

            init22_ = lambda36_(d_726_i0_)
            nw133_ = _dafny.Array(None, 29)
            for i0_22_ in range(nw133_.length(0)):
                nw133_[i0_22_] = init22_(i0_22_)
            d_738_v10_ = nw133_
            d_741_v11_: _dafny.Map
            d_741_v11_ = _dafny.Map({d_725_v0_: d_738_v10_})
            index110_ = default__.safeIndex(144, (d_737_v9_).length(0))
            (d_737_v9_)[index110_] = d_741_v11_
            d_742_v12_: _dafny.Array
            d_742_v12_ = d_727_v1_
            index111_ = default__.safeIndex(275, (d_736_v8_).length(0))
            index112_ = default__.safeIndex(144, (d_737_v9_).length(0))
            rhs133_ = (d_727_v1_)
            rhs134_ = d_726_i0_
            rhs135_ = d_741_v11_
            rhs136_ = d_725_v0_
            lhs112_ = d_736_v8_
            lhs113_ = default__.safeIndex(275, (d_736_v8_).length(0))
            lhs114_ = globalState
            lhs115_ = d_737_v9_
            lhs116_ = default__.safeIndex(144, (d_737_v9_).length(0))
            lhs117_ = globalState
            lhs112_[lhs113_] = rhs133_
            lhs114_.f14 = rhs134_
            lhs115_[lhs116_] = rhs135_
            lhs117_.f14 = rhs136_
        d_743_v13_: _dafny.Array
        def lambda38_(d_744_i2_):
            return self.f24

        init23_ = lambda38_
        nw134_ = _dafny.Array(None, 26)
        for i0_23_ in range(nw134_.length(0)):
            nw134_[i0_23_] = init23_(i0_23_)
        d_743_v13_ = nw134_
        index113_ = default__.safeIndex(226, (d_743_v13_).length(0))
        (d_743_v13_)[index113_] = (self).f25
        d_745_v14_: _dafny.Map
        d_745_v14_ = _dafny.Map({d_725_v0_: d_725_v0_})
        d_746_v16_: _dafny.MultiSet
        def iife41_():
            coll27_ = _dafny.Map()
            compr_27_: int
            for compr_27_ in _dafny.IntegerRange(-595, 441):
                d_747_v15_: int = compr_27_
                if ((-595) <= (d_747_v15_)) and ((d_747_v15_) < (441)):
                    coll27_[default__.safeModuloInt(d_747_v15_, d_725_v0_)] = d_725_v0_
            return _dafny.Map(coll27_)
        d_746_v16_ = _dafny.MultiSet([d_745_v14_, iife41_()
        , d_745_v14_])
        d_748_v17_: _dafny.Map
        d_748_v17_ = _dafny.Map({d_746_v16_: d_725_v0_})
        index114_ = default__.safeIndex(226, (d_743_v13_).length(0))
        rhs137_ = not ((d_725_v0_) < (d_725_v0_)) or (self.f24)
        rhs138_ = self.f24
        rhs139_ = self.f24
        rhs140_ = ((d_748_v17_)[d_746_v16_] if (d_746_v16_) in (d_748_v17_) else (d_725_v0_) - (d_725_v0_))
        lhs118_ = globalState
        lhs119_ = self
        lhs120_ = d_743_v13_
        lhs121_ = default__.safeIndex(226, (d_743_v13_).length(0))
        lhs122_ = globalState
        lhs118_.f9 = rhs137_
        lhs119_.f24 = rhs138_
        lhs120_[lhs121_] = rhs139_
        lhs122_.f5 = rhs140_
        d_749_v18_: _dafny.Seq
        d_749_v18_ = _dafny.SeqWithoutIsStrInference([(d_743_v13_)[default__.safeIndex(226, (d_743_v13_).length(0))], False])
        d_750_v19_: _dafny.Map
        d_750_v19_ = _dafny.Map({(self).f27: self.f24})
        d_751_v20_: _dafny.MultiSet
        d_751_v20_ = _dafny.MultiSet([((d_750_v19_)[self.f24] if (self.f24) in (d_750_v19_) else (self).f25)])
        d_752_v21_: _dafny.Seq
        d_752_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwbojnlo"))
        d_753_v22_: _dafny.MultiSet
        d_753_v22_ = _dafny.MultiSet([d_725_v0_, d_725_v0_])
        d_754_v23_: _dafny.Set
        d_754_v23_ = _dafny.Set({d_725_v0_})
        d_755_v24_: _dafny.Array
        nw135_ = _dafny.Array(None, 19)
        nw135_[int(0)] = d_725_v0_
        nw135_[int(1)] = d_725_v0_
        nw135_[int(2)] = d_725_v0_
        nw135_[int(3)] = d_725_v0_
        nw135_[int(4)] = d_725_v0_
        nw135_[int(5)] = d_725_v0_
        nw135_[int(6)] = len((d_749_v18_).set(default__.safeIndex(d_725_v0_, len(d_749_v18_)), self.f24))
        nw135_[int(7)] = d_725_v0_
        nw135_[int(8)] = d_725_v0_
        nw135_[int(9)] = (d_751_v20_).cardinality
        nw135_[int(10)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))
        nw135_[int(11)] = d_725_v0_
        nw135_[int(12)] = d_725_v0_
        nw135_[int(13)] = len(d_752_v21_)
        nw135_[int(14)] = ((d_753_v22_)[d_725_v0_] if (d_725_v0_) in (d_753_v22_) else (0) - ((0) - (d_725_v0_)))
        nw135_[int(15)] = d_725_v0_
        nw135_[int(16)] = len(d_754_v23_)
        nw135_[int(17)] = d_725_v0_
        nw135_[int(18)] = d_725_v0_
        d_755_v24_ = nw135_
        d_756_v25_: _dafny.Seq
        d_756_v25_ = _dafny.SeqWithoutIsStrInference([d_755_v24_])
        d_757_v26_: _dafny.Seq
        d_757_v26_ = _dafny.SeqWithoutIsStrInference([662, d_725_v0_, len(d_752_v21_)])
        d_758_v27_: _dafny.Seq
        d_758_v27_ = _dafny.SeqWithoutIsStrInference([(d_757_v26_)[default__.safeIndex(d_725_v0_, len(d_757_v26_))]])
        hi6_ = (len((d_756_v25_).set(default__.safeIndex(len(d_757_v26_), len(d_756_v25_)), (d_756_v25_)[default__.safeIndex(d_725_v0_, len(d_756_v25_))]))) - (845)
        for d_759_i3_ in range(d_725_v0_, hi6_):
            if (self).f25:
                d_752_v21_ = d_752_v21_
                d_749_v18_ = d_749_v18_
                d_760_v28_: _dafny.Array
                def lambda39_(d_761_i3_, d_762_v21_):
                    def lambda40_(d_763_i4_):
                        return (_dafny.Map({_dafny.CodePoint('i'): d_761_i3_})) | (_dafny.Map({_dafny.CodePoint('o'): len((D3_DC8(d_762_v21_)).cf14)}))

                    return lambda40_

                init24_ = lambda39_(d_759_i3_, d_752_v21_)
                nw136_ = _dafny.Array(None, 26)
                for i0_24_ in range(nw136_.length(0)):
                    nw136_[i0_24_] = init24_(i0_24_)
                d_760_v28_ = nw136_
                d_764_v29_: str
                d_764_v29_ = _dafny.CodePoint('n')
                d_765_v30_: _dafny.Map
                d_765_v30_ = _dafny.Map({d_764_v29_: d_759_i3_})
                index115_ = default__.safeIndex(716, (d_760_v28_).length(0))
                (d_760_v28_)[index115_] = d_765_v30_
                index116_ = default__.safeIndex(716, (d_760_v28_).length(0))
                (d_760_v28_)[index116_] = (_dafny.Map({d_764_v29_: d_725_v0_})).set(d_764_v29_, d_759_i3_)
                d_764_v29_ = ((d_752_v21_)[default__.safeIndex(d_725_v0_, len(d_752_v21_))] if (d_759_i3_) < (d_725_v0_) else d_764_v29_)
                d_766_v31_: D0
                d_766_v31_ = D0_DC1((d_743_v13_)[default__.safeIndex(226, (d_743_v13_).length(0))], d_725_v0_)
                d_767_v32_: D0
                d_767_v32_ = D0_DC2(d_766_v31_)
                d_768_v33_: _dafny.Map
                d_768_v33_ = _dafny.Map({(self).f25: d_767_v32_})
                d_768_v33_ = (d_768_v33_).set(self.f24, d_767_v32_)
            elif True:
                d_769_v34_: _dafny.Map
                d_769_v34_ = _dafny.Map({(self).fm14(globalState): d_743_v13_})
                d_769_v34_ = d_769_v34_
                d_750_v19_ = (d_750_v19_) | ((d_750_v19_) | (d_750_v19_))
                (globalState).f6 = (D0_DC0(self.f24)).cf0
                (globalState).f9 = (d_725_v0_) <= ((d_759_i3_) - (468))
                d_770_v35_: _dafny.MultiSet
                d_770_v35_ = _dafny.MultiSet([d_753_v22_])
                d_770_v35_ = d_770_v35_
            d_771_v36_: _dafny.Array
            def lambda41_(d_772_v18_):
                def lambda42_(d_773_i5_):
                    return d_772_v18_

                return lambda42_

            init25_ = lambda41_(d_749_v18_)
            nw137_ = _dafny.Array(None, 15)
            for i0_25_ in range(nw137_.length(0)):
                nw137_[i0_25_] = init25_(i0_25_)
            d_771_v36_ = nw137_
            d_774_v37_: D3
            d_774_v37_ = D3_DC9(len(d_752_v21_))
            d_775_v38_: _dafny.Map
            d_775_v38_ = _dafny.Map({d_774_v37_: 666})
            d_776_v39_: _dafny.Map
            d_776_v39_ = _dafny.Map({d_775_v38_: d_771_v36_})
            d_777_v40_: _dafny.Array
            nw138_ = _dafny.Array(None, 19)
            nw138_[int(0)] = d_771_v36_
            nw138_[int(1)] = d_771_v36_
            nw138_[int(2)] = d_771_v36_
            nw138_[int(3)] = d_771_v36_
            nw138_[int(4)] = ((d_776_v39_)[d_775_v38_] if (d_775_v38_) in (d_776_v39_) else d_771_v36_)
            nw138_[int(5)] = d_771_v36_
            nw138_[int(6)] = d_771_v36_
            nw138_[int(7)] = d_771_v36_
            nw138_[int(8)] = d_771_v36_
            nw138_[int(9)] = d_771_v36_
            nw138_[int(10)] = d_771_v36_
            nw138_[int(11)] = d_771_v36_
            nw138_[int(12)] = d_771_v36_
            nw138_[int(13)] = d_771_v36_
            nw138_[int(14)] = d_771_v36_
            nw138_[int(15)] = d_771_v36_
            nw138_[int(16)] = d_771_v36_
            nw138_[int(17)] = d_771_v36_
            nw138_[int(18)] = d_771_v36_
            d_777_v40_ = nw138_
            index117_ = default__.safeIndex(901, (d_777_v40_).length(0))
            (d_777_v40_)[index117_] = d_771_v36_
            d_778_v41_: _dafny.Set
            d_778_v41_ = _dafny.Set({d_757_v26_})
            index118_ = default__.safeIndex(901, (d_777_v40_).length(0))
            rhs141_ = (self).fm15((len(d_778_v41_)) + (d_725_v0_), 375, (self).fm14(globalState), globalState)
            rhs142_ = d_771_v36_
            lhs123_ = globalState
            lhs124_ = d_777_v40_
            lhs125_ = default__.safeIndex(901, (d_777_v40_).length(0))
            lhs123_.f14 = rhs141_
            lhs124_[lhs125_] = rhs142_
            if (self).f27:
                r0 = (self).f25
                d_779_v42_: str
                d_779_v42_ = _dafny.CodePoint('o')
                d_780_v43_: _dafny.Map
                d_780_v43_ = _dafny.Map({(self).f25: d_779_v42_})
                index119_ = default__.safeIndex(643, (d_755_v24_).length(0))
                (d_755_v24_)[index119_] = len(d_780_v43_)
                index120_ = default__.safeIndex(643, (d_755_v24_).length(0))
                (d_755_v24_)[index120_] = d_759_i3_
                d_781_v44_: _dafny.Map
                d_781_v44_ = _dafny.Map({len(d_752_v21_): d_743_v13_})
                d_781_v44_ = (d_781_v44_).set(94, d_743_v13_)
                d_782_v45_: _dafny.Map
                d_782_v45_ = _dafny.Map({d_743_v13_: d_753_v22_})
                d_783_v46_: _dafny.Seq
                d_783_v46_ = _dafny.SeqWithoutIsStrInference([d_753_v22_, _dafny.MultiSet(d_758_v27_)])
                d_782_v45_ = (d_782_v45_).set(d_743_v13_, (d_783_v46_)[default__.safeIndex(d_759_i3_, len(d_783_v46_))])
                (globalState).f5 = d_725_v0_
            elif True:
                (globalState).f8 = (0) - (default__.safeModuloInt(d_759_i3_, (0) - ((d_759_i3_) - (d_759_i3_))))
                index121_ = default__.safeIndex(226, (d_743_v13_).length(0))
                (d_743_v13_)[index121_] = (self).fm14(globalState)
                r0 = (self).f27
                d_784_v47_: C1
                nw139_ = C1()
                nw139_.ctor__(self.f24, True)
                d_784_v47_ = nw139_
                d_785_v48_: str
                d_785_v48_ = _dafny.CodePoint('c')
                d_786_v49_: _dafny.Array
                nw140_ = _dafny.Array(_dafny.Seq({}), 10)
                d_786_v49_ = nw140_
                d_787_v50_: _dafny.Array
                nw141_ = _dafny.Array(_dafny.Map({}), 22)
                d_787_v50_ = nw141_
                rhs143_ = (d_752_v21_)[default__.safeIndex(d_725_v0_, len(d_752_v21_))]
                rhs144_ = (0) - (default__.safeDivisionInt(d_759_i3_, len(d_749_v18_)))
                rhs145_ = d_786_v49_
                rhs146_ = d_752_v21_
                rhs147_ = (d_787_v50_ if (d_743_v13_)[default__.safeIndex(226, (d_743_v13_).length(0))] else d_787_v50_)
                lhs126_ = globalState
                d_785_v48_ = rhs143_
                lhs126_.f8 = rhs144_
                d_786_v49_ = rhs145_
                d_752_v21_ = rhs146_
                d_787_v50_ = rhs147_
            d_788_v51_: C3
            nw142_ = C3()
            nw142_.ctor__(not((d_753_v22_).isdisjoint(_dafny.MultiSet([d_725_v0_]))), (0) - (d_759_i3_), (self).f27, self.f24, d_725_v0_)
            d_788_v51_ = nw142_
        d_789_v52_: _dafny.Map
        d_789_v52_ = _dafny.Map({(self).f25: d_755_v24_})
        d_789_v52_ = d_789_v52_
        (globalState).f5 = d_725_v0_
        d_790_i6_: int
        d_790_i6_ = 0
        with _dafny.label("6"):
            while self.f24:
                with _dafny.c_label("6"):
                    if (d_790_i6_) >= (100):
                        raise _dafny.Break("6")
                    d_790_i6_ = (d_790_i6_) + (1)
                    d_791_v53_: str
                    d_791_v53_ = _dafny.CodePoint('d')
                    d_792_v54_: _dafny.Seq
                    d_792_v54_ = _dafny.SeqWithoutIsStrInference([d_752_v21_, (d_752_v21_).set(default__.safeIndex(d_725_v0_, len(d_752_v21_)), d_791_v53_)])
                    d_792_v54_ = d_792_v54_
                    nw143_ = _dafny.Array(int(0), 15)
                    d_755_v24_ = nw143_
                    d_793_v55_: _dafny.Set
                    d_793_v55_ = _dafny.Set({not(self.f24), True})
                    d_794_v56_: _dafny.Map
                    d_794_v56_ = _dafny.Map({d_793_v55_: (d_743_v13_)[default__.safeIndex(226, (d_743_v13_).length(0))]})
                    d_795_v58_: _dafny.Seq
                    d_795_v58_ = _dafny.SeqWithoutIsStrInference([d_794_v56_])
                    d_796_v59_: _dafny.Map
                    d_796_v59_ = _dafny.Map({d_725_v0_: d_794_v56_})
                    d_797_v61_: _dafny.Array
                    nw144_ = _dafny.Array(None, 14)
                    nw144_[int(0)] = d_794_v56_
                    nw144_[int(1)] = d_794_v56_
                    def iife42_():
                        coll28_ = _dafny.Map()
                        compr_28_: _dafny.Set
                        for compr_28_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_793_v55_ for d_798_i7_ in range(default__.abs(659))]))).Elements:
                            d_799_v57_: _dafny.Set = compr_28_
                            if (d_799_v57_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_793_v55_ for d_798_i7_ in range(default__.abs(659))]))):
                                coll28_[d_799_v57_] = (self).f27
                        return _dafny.Map(coll28_)
                    nw144_[int(2)] = (iife42_()
                     if (self).f27 else d_794_v56_)
                    nw144_[int(3)] = (d_794_v56_) | ((d_795_v58_)[default__.safeIndex(d_725_v0_, len(d_795_v58_))])
                    nw144_[int(4)] = ((d_796_v59_)[d_725_v0_] if (d_725_v0_) in (d_796_v59_) else default__.fm35(globalState))
                    nw144_[int(5)] = d_794_v56_
                    nw144_[int(6)] = (d_794_v56_) | (_dafny.Map({_dafny.Set({(d_743_v13_)[default__.safeIndex(226, (d_743_v13_).length(0))], (self).f25}): self.f24}))
                    nw144_[int(7)] = d_794_v56_
                    nw144_[int(8)] = (d_794_v56_) | (d_794_v56_)
                    nw144_[int(9)] = d_794_v56_
                    nw144_[int(10)] = d_794_v56_
                    def iife43_():
                        coll29_ = _dafny.Map()
                        compr_29_: _dafny.Set
                        for compr_29_ in (_dafny.SeqWithoutIsStrInference([d_793_v55_ for d_800_i8_ in range(default__.abs(592))])).Elements:
                            d_801_v60_: _dafny.Set = compr_29_
                            if (d_801_v60_) in (_dafny.SeqWithoutIsStrInference([d_793_v55_ for d_800_i8_ in range(default__.abs(592))])):
                                coll29_[d_801_v60_] = (self).f27
                        return _dafny.Map(coll29_)
                    nw144_[int(11)] = (iife43_()
                     if (d_743_v13_)[default__.safeIndex(226, (d_743_v13_).length(0))] else d_794_v56_)
                    nw144_[int(12)] = d_794_v56_
                    nw144_[int(13)] = d_794_v56_
                    d_797_v61_ = nw144_
                    index122_ = default__.safeIndex(802, (d_797_v61_).length(0))
                    (d_797_v61_)[index122_] = d_794_v56_
                    index123_ = default__.safeIndex(802, (d_797_v61_).length(0))
                    (d_797_v61_)[index123_] = d_794_v56_
                    (globalState).f1 = self.f24
                    pass
            pass
        r0 = (not((self).f27) if ((self).f25) == ((self).f25) else (235) == (d_725_v0_))
        d_802_v62_: T0
        nw145_ = C2()
        nw145_.ctor__(len(_dafny.SeqWithoutIsStrInference([self.f24, (self).f27])))
        d_802_v62_ = nw145_
        d_803_v63_: D11
        d_803_v63_ = D11_DC30(d_802_v62_)
        r1 = (d_803_v63_).cf56
        r2 = (self.f24) == (self.f24)
        d_804_v65_: _dafny.MultiSet
        d_804_v65_ = _dafny.MultiSet([d_749_v18_])
        def iife44_():
            coll30_ = _dafny.Map()
            compr_30_: _dafny.Seq
            for compr_30_ in (d_804_v65_).Elements:
                d_805_v64_: _dafny.Seq = compr_30_
                if (d_805_v64_) in (d_804_v65_):
                    coll30_[d_805_v64_] = self.f24
            return _dafny.Map(coll30_)
        r3 = (len((_dafny.Map({d_749_v18_: False})) | (iife44_()
        )) if (self).f25 else d_725_v0_)
        return r0, r1, r2, r3

    @property
    def f27(self):
        return self._f27

class C5(T1):
    def  __init__(self):
        self._f24: bool = False
        self._f25: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f24, f25):
        (self).f24 = f24
        (self)._f25 = f25

    def fm4(self, p0, p1, p2, globalState):
        source7_ = D0_DC0((self).f25)
        if source7_.is_DC1:
            d_806___mcc_h0_ = source7_.cf1
            d_807___mcc_h1_ = source7_.cf2
            d_808_cf2_ = d_807___mcc_h1_
            d_809_cf1_ = d_806___mcc_h0_
            return (_dafny.Map({self.f24: _dafny.Map({d_809_cf1_: not(False)})})) | (_dafny.Map({d_809_cf1_: _dafny.Map({True: d_809_cf1_})}))
        elif source7_.is_DC0:
            d_810___mcc_h2_ = source7_.cf0
            d_811_cf0_ = d_810___mcc_h2_
            return _dafny.Map({self.f24: _dafny.Map({(self).f25: (self).f25})})
        elif True:
            d_812___mcc_h3_ = source7_.cf3
            d_813_cf3_ = d_812___mcc_h3_
            return (_dafny.Map({not(self.f24): _dafny.Map({True: self.f24})})) | (_dafny.Map({self.f24: _dafny.Map({(self).f25: (self).f25})}))

    def fm5(self, p0, p1, globalState):
        return len((_dafny.SeqWithoutIsStrInference([_dafny.Set({86, 469, len(_dafny.SeqWithoutIsStrInference([self.f24]))}) for d_814_i0_ in range(default__.abs(5))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({224})])))

    def fm13(self, p0, p1, globalState):
        return default__.safeDivisionInt((len(_dafny.SeqWithoutIsStrInference([939]))) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({634})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))]))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_815_i0_ in range(default__.abs(680))])))

    def m1(self, globalState):
        r0: bool = False
        r1: T0 = None
        r2: bool = False
        r3: int = int(0)
        d_816_v0_: int
        d_816_v0_ = 82
        d_817_v1_: str
        d_817_v1_ = _dafny.CodePoint('v')
        d_818_v2_: _dafny.Seq
        d_818_v2_ = _dafny.SeqWithoutIsStrInference([d_817_v1_, d_817_v1_])
        hi7_ = len((d_818_v2_) + (_dafny.SeqWithoutIsStrInference([d_817_v1_, d_817_v1_])))
        for d_819_i0_ in range(d_816_v0_, hi7_):
            d_820_v3_: C2
            nw146_ = C2()
            nw146_.ctor__(d_816_v0_)
            d_820_v3_ = nw146_
            if not(self.f24):
                d_821_v4_: _dafny.Map
                d_821_v4_ = _dafny.Map({d_819_i0_: 134})
                d_821_v4_ = (d_821_v4_).set(d_819_i0_, d_819_i0_)
                (globalState).f1 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "radkjif"))) < (d_818_v2_)
                d_822_v5_: _dafny.Array
                def lambda43_(d_823_i1_):
                    return (self).f25

                init26_ = lambda43_
                nw147_ = _dafny.Array(None, 10)
                for i0_26_ in range(nw147_.length(0)):
                    nw147_[i0_26_] = init26_(i0_26_)
                d_822_v5_ = nw147_
                d_824_v6_: C0
                nw148_ = C0()
                nw148_.ctor__(_dafny.Set({d_822_v5_, d_822_v5_, d_822_v5_, d_822_v5_}), (self).fm13(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arwqolr"))), (self).f25, globalState), -476)
                d_824_v6_ = nw148_
                d_825_v7_: _dafny.Map
                d_825_v7_ = _dafny.Map({d_824_v6_: self.f24})
                d_825_v7_ = (d_825_v7_).set(d_824_v6_, (self).f25)
                d_826_v8_: _dafny.Map
                d_826_v8_ = _dafny.Map({d_822_v5_: d_824_v6_.f31})
                d_827_v9_: _dafny.Set
                d_827_v9_ = _dafny.Set({self.f24, False, False})
                d_828_v10_: _dafny.Array
                nw149_ = _dafny.Array(None, 2)
                nw149_[int(0)] = len(d_827_v9_)
                nw149_[int(1)] = d_816_v0_
                d_828_v10_ = nw149_
                d_829_v11_: D11
                d_829_v11_ = D11_DC31(d_828_v10_, d_819_i0_, self.f24)
                d_830_v12_: _dafny.Array
                nw150_ = _dafny.Array(None, 8)
                nw150_[int(0)] = d_817_v1_
                nw150_[int(1)] = default__.fm26((d_829_v11_).cf58, (self).f25, globalState)
                nw150_[int(2)] = d_817_v1_
                nw150_[int(3)] = d_817_v1_
                nw150_[int(4)] = _dafny.CodePoint('y')
                nw150_[int(5)] = d_817_v1_
                nw150_[int(6)] = d_817_v1_
                nw150_[int(7)] = d_817_v1_
                d_830_v12_ = nw150_
                index124_ = default__.safeIndex(677, (d_830_v12_).length(0))
                (d_830_v12_)[index124_] = d_817_v1_
                d_831_v13_: _dafny.Array
                nw151_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_831_v13_ = nw151_
                index125_ = default__.safeIndex(181, (d_831_v13_).length(0))
                (d_831_v13_)[index125_] = d_830_v12_
                index126_ = default__.safeIndex(740, (d_822_v5_).length(0))
                (d_822_v5_)[index126_] = (self).f25
                d_832_v14_: _dafny.Seq
                d_832_v14_ = _dafny.SeqWithoutIsStrInference([(self).f25, self.f24])
                d_833_v15_: _dafny.MultiSet
                d_833_v15_ = _dafny.MultiSet([(self).f25, self.f24, ((self).f25) and (not(self.f24)), (d_832_v14_)[default__.safeIndex((d_820_v3_).fm2(globalState), len(d_832_v14_))], (self).f25])
                index127_ = default__.safeIndex(677, (d_830_v12_).length(0))
                index128_ = default__.safeIndex(181, (d_831_v13_).length(0))
                index129_ = default__.safeIndex(740, (d_822_v5_).length(0))
                rhs148_ = d_826_v8_
                rhs149_ = d_817_v1_
                rhs150_ = ((d_833_v15_)[(_dafny.CodePoint('n')) in (default__.fm25(True, self.f24, d_817_v1_, globalState))] if ((_dafny.CodePoint('n')) in (default__.fm25(True, self.f24, d_817_v1_, globalState))) in (d_833_v15_) else 191)
                rhs151_ = d_830_v12_
                rhs152_ = default__.fm29(len(d_818_v2_), d_824_v6_.f31, default__.fm29(d_824_v6_.f31, d_824_v6_.f31, self.f24, globalState), globalState)
                lhs127_ = d_830_v12_
                lhs128_ = default__.safeIndex(677, (d_830_v12_).length(0))
                lhs129_ = globalState
                lhs130_ = d_831_v13_
                lhs131_ = default__.safeIndex(181, (d_831_v13_).length(0))
                lhs132_ = d_822_v5_
                lhs133_ = default__.safeIndex(740, (d_822_v5_).length(0))
                d_826_v8_ = rhs148_
                lhs127_[lhs128_] = rhs149_
                lhs129_.f17 = rhs150_
                lhs130_[lhs131_] = rhs151_
                lhs132_[lhs133_] = rhs152_
                r0 = (d_822_v5_)[default__.safeIndex(740, (d_822_v5_).length(0))]
            elif True:
                d_834_v16_: _dafny.MultiSet
                d_834_v16_ = _dafny.MultiSet([not(self.f24)])
                d_835_v17_: _dafny.Seq
                d_835_v17_ = _dafny.SeqWithoutIsStrInference([(d_834_v16_).isdisjoint(d_834_v16_), (self).f25])
                d_836_v18_: _dafny.Map
                d_836_v18_ = _dafny.Map({d_817_v1_: d_816_v0_})
                (globalState).f15 = (d_835_v17_)[default__.safeIndex(default__.safeModuloInt(len(d_836_v18_), d_816_v0_), len(d_835_v17_))]
                (globalState).f14 = d_816_v0_
                d_837_v19_: _dafny.Array
                nw152_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                d_837_v19_ = nw152_
                index130_ = default__.safeIndex(195, (d_837_v19_).length(0))
                (d_837_v19_)[index130_] = (d_818_v2_ if (d_835_v17_)[default__.safeIndex(281, len(d_835_v17_))] else d_818_v2_)
                index131_ = default__.safeIndex(195, (d_837_v19_).length(0))
                (d_837_v19_)[index131_] = ((d_818_v2_) + (d_818_v2_)) + (d_818_v2_)
                d_838_v20_: _dafny.Map
                d_838_v20_ = _dafny.Map({len(d_835_v17_): len((d_837_v19_)[default__.safeIndex(195, (d_837_v19_).length(0))])})
                d_839_v21_: T0
                nw153_ = C2()
                nw153_.ctor__(len(d_838_v20_))
                d_839_v21_ = nw153_
                d_840_v22_: D11
                d_840_v22_ = D11_DC30(d_839_v21_)
                d_840_v22_ = d_840_v22_
                d_841_v23_: _dafny.Array
                nw154_ = _dafny.Array(None, 1)
                nw154_[int(0)] = (self).f25
                d_841_v23_ = nw154_
                d_842_v24_: _dafny.Array
                nw155_ = _dafny.Array(int(0), 18)
                d_842_v24_ = nw155_
                d_843_v25_: D11
                d_843_v25_ = D11_DC31(d_842_v24_, d_816_v0_, (d_835_v17_)[default__.safeIndex((d_839_v21_).f23, len(d_835_v17_))])
                index132_ = default__.safeIndex(340, (d_841_v23_).length(0))
                (d_841_v23_)[index132_] = (d_843_v25_).cf59
                index133_ = default__.safeIndex(340, (d_841_v23_).length(0))
                (d_841_v23_)[index133_] = (self).f25
            if (self).f25:
                d_844_v26_: _dafny.Array
                nw156_ = _dafny.Array(int(0), 22)
                d_844_v26_ = nw156_
                index134_ = default__.safeIndex(787, (d_844_v26_).length(0))
                (d_844_v26_)[index134_] = d_816_v0_
                d_845_v27_: _dafny.Seq
                d_845_v27_ = _dafny.SeqWithoutIsStrInference([d_844_v26_])
                index135_ = default__.safeIndex(787, (d_844_v26_).length(0))
                rhs153_ = d_816_v0_
                rhs154_ = (d_819_i0_) >= ((0) - (default__.safeDivisionInt(-158, 526)))
                rhs155_ = len(d_845_v27_)
                rhs156_ = (self).f25
                rhs157_ = d_819_i0_
                lhs134_ = globalState
                lhs135_ = globalState
                lhs136_ = d_844_v26_
                lhs137_ = default__.safeIndex(787, (d_844_v26_).length(0))
                lhs138_ = globalState
                lhs139_ = globalState
                lhs134_.f17 = rhs153_
                lhs135_.f9 = rhs154_
                lhs136_[lhs137_] = rhs155_
                lhs138_.f1 = rhs156_
                lhs139_.f14 = rhs157_
                d_846_v28_: _dafny.Map
                d_846_v28_ = _dafny.Map({(self).f25: d_817_v1_})
                d_847_v29_: _dafny.MultiSet
                d_847_v29_ = _dafny.MultiSet([d_846_v28_, d_846_v28_])
                (globalState).f15 = (d_847_v29_).ispropersubset(d_847_v29_)
                index136_ = default__.safeIndex(787, (d_844_v26_).length(0))
                rhs158_ = d_845_v27_
                rhs159_ = d_819_i0_
                rhs160_ = d_816_v0_
                lhs140_ = d_844_v26_
                lhs141_ = default__.safeIndex(787, (d_844_v26_).length(0))
                d_845_v27_ = rhs158_
                lhs140_[lhs141_] = rhs159_
                d_816_v0_ = rhs160_
                (globalState).f8 = (d_816_v0_) + (d_819_i0_)
                rhs161_ = True
                rhs162_ = not((self).f25)
                rhs163_ = (d_844_v26_)[default__.safeIndex(787, (d_844_v26_).length(0))]
                rhs164_ = d_819_i0_
                lhs142_ = globalState
                lhs143_ = globalState
                lhs144_ = globalState
                lhs145_ = globalState
                lhs142_.f1 = rhs161_
                lhs143_.f1 = rhs162_
                lhs144_.f14 = rhs163_
                lhs145_.f17 = rhs164_
            elif True:
                (globalState).f1 = self.f24
                (globalState).f17 = d_816_v0_
                d_848_v30_: _dafny.Array
                nw157_ = _dafny.Array(False, 26)
                d_848_v30_ = nw157_
                d_849_v31_: _dafny.Seq
                d_849_v31_ = _dafny.SeqWithoutIsStrInference([d_848_v30_, d_848_v30_, (d_848_v30_ if (self).f25 else d_848_v30_)])
                d_849_v31_ = (((d_849_v31_) + (d_849_v31_)).set(default__.safeIndex(d_819_i0_, len((d_849_v31_) + (d_849_v31_))), d_848_v30_)) + (d_849_v31_)
                (globalState).f8 = (default__.safeDivisionInt(d_816_v0_, (0) - (d_819_i0_))) - (d_819_i0_)
                d_850_v32_: _dafny.Set
                d_850_v32_ = _dafny.Set({not(self.f24)})
                d_851_v33_: _dafny.Array
                nw158_ = _dafny.Array(None, 3)
                nw158_[int(0)] = d_850_v32_
                nw158_[int(1)] = d_850_v32_
                nw158_[int(2)] = d_850_v32_
                d_851_v33_ = nw158_
                index137_ = default__.safeIndex(775, (d_851_v33_).length(0))
                (d_851_v33_)[index137_] = d_850_v32_
                d_852_v34_: _dafny.Map
                d_852_v34_ = _dafny.Map({self.f24: d_816_v0_})
                d_853_v35_: _dafny.Array
                nw159_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                d_853_v35_ = nw159_
                d_854_v36_: D7
                d_854_v36_ = D7_DC20((self).f25, d_852_v34_, d_853_v35_)
                d_855_v37_: D7
                d_855_v37_ = D7_DC21(d_854_v36_)
                d_856_v38_: D7
                d_856_v38_ = D7_DC21(d_855_v37_)
                pat_let_tv12_ = d_852_v34_
                pat_let_tv13_ = d_853_v35_
                d_857_v39_: _dafny.MultiSet
                def iife45_(_pat_let7_0):
                    def iife46_(d_858_dt__update__tmp_h0_):
                        def iife47_(_pat_let8_0):
                            def iife48_(d_859_dt__update_hcf38_h0_):
                                return D7_DC21(d_859_dt__update_hcf38_h0_)
                            return iife48_(_pat_let8_0)
                        return iife47_(D7_DC20((self).f25, pat_let_tv12_, pat_let_tv13_))
                    return iife46_(_pat_let7_0)
                d_857_v39_ = _dafny.MultiSet([iife45_(d_856_v38_), d_856_v38_])
                d_860_v40_: _dafny.Seq
                d_860_v40_ = _dafny.SeqWithoutIsStrInference([(d_857_v39_).cardinality])
                d_861_v41_: _dafny.Seq
                d_861_v41_ = _dafny.SeqWithoutIsStrInference([(self).f25, self.f24])
                d_862_v42_: D1
                d_862_v42_ = D1_DC4((self).f25, self.f24)
                d_863_v43_: _dafny.Set
                d_863_v43_ = _dafny.Set({d_862_v42_})
                index138_ = default__.safeIndex(775, (d_851_v33_).length(0))
                (d_851_v33_)[index138_] = (d_850_v32_).intersection((default__.fm32((self).f25, d_860_v40_, d_861_v41_, d_863_v43_, globalState)) | (d_850_v32_))
            d_864_v44_: _dafny.Array
            nw160_ = _dafny.Array(False, 25)
            d_864_v44_ = nw160_
            d_865_v45_: _dafny.Map
            d_865_v45_ = _dafny.Map({False: False})
            d_866_v46_: _dafny.Map
            d_866_v46_ = _dafny.Map({self.f24: d_819_i0_})
            d_867_v47_: D1
            d_867_v47_ = D1_DC5((self).f25, d_865_v45_, d_819_i0_, d_819_i0_, d_866_v46_)
            index139_ = default__.safeIndex(136, (d_864_v44_).length(0))
            (d_864_v44_)[index139_] = not((d_867_v47_).cf7)
            index140_ = default__.safeIndex(136, (d_864_v44_).length(0))
            (d_864_v44_)[index140_] = not(((d_819_i0_) - (d_816_v0_)) == (d_819_i0_))
        d_868_v48_: _dafny.Seq
        d_868_v48_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        d_869_v49_: _dafny.MultiSet
        d_869_v49_ = _dafny.MultiSet([False])
        d_870_v50_: _dafny.Seq
        d_870_v50_ = _dafny.SeqWithoutIsStrInference([d_816_v0_, d_816_v0_, ((d_869_v49_)[self.f24] if (self.f24) in (d_869_v49_) else d_816_v0_)])
        r0 = (d_868_v48_)[default__.safeIndex((d_870_v50_)[default__.safeIndex(d_816_v0_, len(d_870_v50_))], len(d_868_v48_))]
        d_871_v51_: _dafny.Map
        d_871_v51_ = _dafny.Map({self.f24: d_816_v0_})
        d_872_v52_: _dafny.Map
        d_872_v52_ = _dafny.Map({d_816_v0_: _dafny.Map({d_816_v0_: d_816_v0_})})
        d_873_v53_: _dafny.Array
        nw161_ = _dafny.Array(None, 18)
        nw161_[int(0)] = (d_816_v0_) - ((0) - (d_816_v0_))
        nw161_[int(1)] = len(d_871_v51_)
        nw161_[int(2)] = d_816_v0_
        nw161_[int(3)] = d_816_v0_
        nw161_[int(4)] = d_816_v0_
        nw161_[int(5)] = len(d_872_v52_)
        nw161_[int(6)] = d_816_v0_
        nw161_[int(7)] = d_816_v0_
        nw161_[int(8)] = d_816_v0_
        nw161_[int(9)] = len(d_818_v2_)
        nw161_[int(10)] = d_816_v0_
        nw161_[int(11)] = (0) - (d_816_v0_)
        nw161_[int(12)] = d_816_v0_
        nw161_[int(13)] = default__.safeModuloInt(d_816_v0_, d_816_v0_)
        nw161_[int(14)] = (0) - (d_816_v0_)
        nw161_[int(15)] = d_816_v0_
        nw161_[int(16)] = len(d_870_v50_)
        nw161_[int(17)] = (d_816_v0_ if self.f24 else d_816_v0_)
        d_873_v53_ = nw161_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_873_v53_).length(0)):
            d_874_i2_: int = guard_loop_2_
            if (True) and (((0) <= (d_874_i2_)) and ((d_874_i2_) < ((d_873_v53_).length(0)))):
                (d_873_v53_)[(d_874_i2_)] = (d_874_i2_) + ((d_816_v0_ if (self).f25 else d_816_v0_))
        (globalState).f5 = d_816_v0_
        d_875_v54_: _dafny.Map
        d_875_v54_ = _dafny.Map({d_871_v51_: (self).f25})
        d_876_v57_: _dafny.Map
        d_876_v57_ = _dafny.Map({d_816_v0_: d_816_v0_})
        d_877_v58_: _dafny.Array
        nw162_ = _dafny.Array(None, 28)
        nw162_[int(0)] = (self).f25
        nw162_[int(1)] = self.f24
        nw162_[int(2)] = ((d_868_v48_)[default__.safeIndex((0) - ((0) - ((d_870_v50_)[default__.safeIndex(d_816_v0_, len(d_870_v50_))])), len(d_868_v48_))]) == (not(self.f24))
        nw162_[int(3)] = self.f24
        nw162_[int(4)] = (d_875_v54_) == (d_875_v54_)
        nw162_[int(5)] = ((self).f25) or (True)
        nw162_[int(6)] = self.f24
        nw162_[int(7)] = (self).f25
        nw162_[int(8)] = self.f24
        nw162_[int(9)] = (self).f25
        nw162_[int(10)] = True
        nw162_[int(11)] = (self).f25
        nw162_[int(12)] = (self).f25
        nw162_[int(13)] = (self).f25
        def iife49_():
            coll31_ = _dafny.Map()
            compr_31_: int
            for compr_31_ in _dafny.IntegerRange(322, 256):
                d_878_v55_: int = compr_31_
                if ((322) <= (d_878_v55_)) and ((d_878_v55_) < (256)):
                    coll31_[(d_878_v55_) + (d_816_v0_)] = self.f24
            return _dafny.Map(coll31_)
        def iife50_():
            coll32_ = _dafny.Map()
            compr_32_: int
            for compr_32_ in ((d_870_v50_).set(default__.safeIndex(674, len(d_870_v50_)), (d_869_v49_).cardinality)).Elements:
                d_879_v56_: int = compr_32_
                if (d_879_v56_) in ((d_870_v50_).set(default__.safeIndex(674, len(d_870_v50_)), (d_869_v49_).cardinality)):
                    coll32_[(d_879_v56_) - (d_816_v0_)] = (self).f25
            return _dafny.Map(coll32_)
        nw162_[int(14)] = default__.fm29(len(iife49_()
        ), len(iife50_()
        ), (self).f25, globalState)
        nw162_[int(15)] = self.f24
        nw162_[int(16)] = self.f24
        nw162_[int(17)] = (((d_876_v57_)[d_816_v0_] if (d_816_v0_) in (d_876_v57_) else (0) - (d_816_v0_))) != ((0) - (d_816_v0_))
        nw162_[int(18)] = (d_869_v49_).ispropersubset(d_869_v49_)
        nw162_[int(19)] = self.f24
        nw162_[int(20)] = not((self).f25)
        nw162_[int(21)] = (self).f25
        nw162_[int(22)] = default__.fm29(d_816_v0_, d_816_v0_, self.f24, globalState)
        nw162_[int(23)] = not(self.f24)
        nw162_[int(24)] = not((d_816_v0_) <= (d_816_v0_))
        nw162_[int(25)] = (len(d_870_v50_)) < (d_816_v0_)
        nw162_[int(26)] = False
        nw162_[int(27)] = (self).f25
        d_877_v58_ = nw162_
        index141_ = default__.safeIndex(344, (d_877_v58_).length(0))
        (d_877_v58_)[index141_] = True
        index142_ = default__.safeIndex(344, (d_877_v58_).length(0))
        (d_877_v58_)[index142_] = self.f24
        d_880_v59_: C1
        nw163_ = C1()
        nw163_.ctor__(self.f24, True)
        d_880_v59_ = nw163_
        d_881_v60_: _dafny.Map
        d_881_v60_ = _dafny.Map({d_816_v0_: d_880_v59_})
        rhs165_ = _dafny.MultiSet([(d_877_v58_)[default__.safeIndex(344, (d_877_v58_).length(0))], (self).f25, (d_868_v48_)[default__.safeIndex(d_816_v0_, len(d_868_v48_))]])
        rhs166_ = default__.fm29(d_816_v0_, d_816_v0_, (d_816_v0_) == (len(d_881_v60_)), globalState)
        lhs146_ = globalState
        d_869_v49_ = rhs165_
        lhs146_.f9 = rhs166_
        r0 = (d_816_v0_) >= ((0) - (d_816_v0_))
        nw164_ = C3()
        nw164_.ctor__(self.f24, d_816_v0_, (d_877_v58_)[default__.safeIndex(344, (d_877_v58_).length(0))], (d_877_v58_)[default__.safeIndex(344, (d_877_v58_).length(0))], (0) - ((0) - (d_816_v0_)))
        r1 = nw164_
        r2 = self.f24
        r3 = d_816_v0_
        return r0, r1, r2, r3


class C6(T1, T0):
    def  __init__(self):
        self._f24: bool = False
        self._f23: int = int(0)
        self._f25: bool = False
        self.f26: D0 = D0.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    @property
    def f23(self):
        return self._f23
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f26, f24, f25, f23):
        (self).f26 = f26
        (self).f24 = f24
        (self)._f25 = f25
        (self)._f23 = f23

    def fm4(self, p0, p1, p2, globalState):
        return ((_dafny.Map({(self).f25: _dafny.Map({(self).f25: (self).f25})})) | (_dafny.Map({(self).f25: _dafny.Map({self.f24: self.f24})}))) | ((_dafny.Map({self.f24: _dafny.Map({False: True})})) | (_dafny.Map({(self).f25: _dafny.Map({self.f24: self.f24})})))

    def fm5(self, p0, p1, globalState):
        return (self).f23

    def fm2(self, globalState):
        return (self).f23

    def fm3(self, p0, p1, p2, p3, globalState):
        return (self).f25

    def m1(self, globalState):
        r0: bool = False
        r1: T0 = None
        r2: bool = False
        r3: int = int(0)
        d_882_v0_: _dafny.Seq
        d_882_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "koa"))
        d_883_v1_: _dafny.Map
        d_883_v1_ = _dafny.Map({d_882_v0_: self.f24})
        d_883_v1_ = (d_883_v1_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xx")), not(self.f24))
        d_884_v2_: _dafny.Array
        nw165_ = _dafny.Array(_dafny.Array(None, 0), 13)
        d_884_v2_ = nw165_
        d_885_v3_: str
        d_885_v3_ = _dafny.CodePoint('r')
        d_886_v4_: _dafny.Array
        out9_: _dafny.Array
        out9_ = (self).m4(d_884_v2_, (self).f23, not(not(self.f24)), d_885_v3_, globalState)
        d_886_v4_ = out9_
        d_887_v5_: _dafny.Map
        d_887_v5_ = _dafny.Map({(self).f25: (self).f23})
        d_888_i0_: int
        d_888_i0_ = 0
        with _dafny.label("7"):
            while not((((self).f23) == (-781) if (D1_DC5(False, _dafny.Map({(self).f25: (self).f25}), (self).f23, (self).f23, d_887_v5_)).cf7 else (self).f25)):
                with _dafny.c_label("7"):
                    if (d_888_i0_) >= (100):
                        raise _dafny.Break("7")
                    d_888_i0_ = (d_888_i0_) + (1)
                    r2 = (self).f25
                    d_889_v6_: _dafny.Seq
                    d_889_v6_ = _dafny.SeqWithoutIsStrInference([(self).f25])
                    d_890_v7_: _dafny.Array
                    def lambda44_(d_891_i1_):
                        return (d_891_i1_) - ((self).f23)

                    init27_ = lambda44_
                    nw166_ = _dafny.Array(None, 11)
                    for i0_27_ in range(nw166_.length(0)):
                        nw166_[i0_27_] = init27_(i0_27_)
                    d_890_v7_ = nw166_
                    d_892_v8_: _dafny.Set
                    d_892_v8_ = _dafny.Set({d_890_v7_})
                    d_893_v9_: _dafny.Seq
                    d_893_v9_ = _dafny.SeqWithoutIsStrInference([(self).f23])
                    d_894_v10_: _dafny.Map
                    d_894_v10_ = _dafny.Map({self.f24: not((self).f25)})
                    d_895_v11_: _dafny.Array
                    nw167_ = _dafny.Array(None, 26)
                    nw167_[int(0)] = (self).f23
                    nw167_[int(1)] = (self).f23
                    nw167_[int(2)] = ((self).f23 if (self).f25 else len(d_889_v6_))
                    nw167_[int(3)] = len(d_889_v6_)
                    nw167_[int(4)] = len(d_882_v0_)
                    nw167_[int(5)] = (self).f23
                    nw167_[int(6)] = 142
                    nw167_[int(7)] = (self).f23
                    nw167_[int(8)] = (self).f23
                    nw167_[int(9)] = (0) - ((self).f23)
                    nw167_[int(10)] = (self).f23
                    nw167_[int(11)] = ((self).f23) - ((self).f23)
                    nw167_[int(12)] = len(_dafny.SeqWithoutIsStrInference([default__.fm11((self).f23, self.f24, self.f24, _dafny.CodePoint('a'), globalState)]))
                    nw167_[int(13)] = 89
                    nw167_[int(14)] = ((self).f23 if self.f24 else (self).f23)
                    nw167_[int(15)] = -763
                    nw167_[int(16)] = (self).f23
                    nw167_[int(17)] = (self).f23
                    nw167_[int(18)] = (self).f23
                    nw167_[int(19)] = (0) - (((self).f23) + (len(d_892_v8_)))
                    nw167_[int(20)] = (self).f23
                    nw167_[int(21)] = (0) - ((self).f23)
                    nw167_[int(22)] = (self).f23
                    nw167_[int(23)] = (len(d_893_v9_)) + ((self).f23)
                    nw167_[int(24)] = default__.fm1((D1_DC5(not(False), d_894_v10_, (self).f23, (self).f23, d_887_v5_)).cf10, (self).f25, globalState)
                    nw167_[int(25)] = (self).fm2(globalState)
                    d_895_v11_ = nw167_
                    d_896_v12_: _dafny.MultiSet
                    d_896_v12_ = _dafny.MultiSet([d_895_v11_, d_895_v11_, d_890_v7_, d_895_v11_])
                    index143_ = default__.safeIndex(785, (d_890_v7_).length(0))
                    (d_890_v7_)[index143_] = ((d_896_v12_)[d_895_v11_] if (d_895_v11_) in (d_896_v12_) else (self).f23)
                    d_897_v13_: D1
                    d_897_v13_ = D1_DC5(self.f24, d_894_v10_, len(d_893_v9_), len(_dafny.Map({(self).f23: (self).f25})), d_887_v5_)
                    index144_ = default__.safeIndex(785, (d_890_v7_).length(0))
                    (d_890_v7_)[index144_] = (d_897_v13_).cf9
                    d_885_v3_ = default__.fm12((self).f25, globalState)
                    d_898_v14_: T1
                    nw168_ = C5()
                    nw168_.ctor__(self.f24, (self).f25)
                    d_898_v14_ = nw168_
                    d_899_v15_: _dafny.Array
                    out10_: _dafny.Array
                    out10_ = (self).m4(d_884_v2_, (self).f23, (self).fm3(d_882_v0_, self.f24, len(_dafny.Map({d_898_v14_: d_894_v10_})), _dafny.SeqWithoutIsStrInference([(self).f23 for d_900_i2_ in range(default__.abs(930))]), globalState), _dafny.CodePoint('a'), globalState)
                    d_899_v15_ = out10_
                    pass
            pass
        index145_ = default__.safeIndex(911, (d_886_v4_).length(0))
        (d_886_v4_)[index145_] = self.f24
        index146_ = default__.safeIndex(911, (d_886_v4_).length(0))
        (d_886_v4_)[index146_] = False
        d_901_v16_: _dafny.Seq
        d_901_v16_ = _dafny.SeqWithoutIsStrInference([(self).f23])
        d_902_v18_: _dafny.Set
        d_902_v18_ = _dafny.Set({(self).f23, (0) - ((self).f23), (self).f23, (self).f23, (self).f23})
        d_903_v19_: D0
        d_903_v19_ = D0_DC0((self).f25)
        d_904_v20_: D0
        d_904_v20_ = D0_DC2(D0_DC2(d_903_v19_))
        pat_let_tv14_ = d_904_v20_
        d_905_v21_: _dafny.Array
        nw169_ = _dafny.Array(None, 21)
        nw169_[int(0)] = self.f26
        nw169_[int(1)] = self.f26
        nw169_[int(2)] = self.f26
        nw169_[int(3)] = self.f26
        nw169_[int(4)] = self.f26
        nw169_[int(5)] = self.f26
        nw169_[int(6)] = self.f26
        nw169_[int(7)] = self.f26
        nw169_[int(8)] = D0_DC2(d_903_v19_)
        nw169_[int(9)] = self.f26
        nw169_[int(10)] = self.f26
        nw169_[int(11)] = self.f26
        def iife51_(_pat_let9_0):
            def iife52_(d_906_dt__update__tmp_h0_):
                def iife53_(_pat_let10_0):
                    def iife54_(d_907_dt__update_hcf3_h0_):
                        return D0_DC2(d_907_dt__update_hcf3_h0_)
                    return iife54_(_pat_let10_0)
                return iife53_(pat_let_tv14_)
            return iife52_(_pat_let9_0)
        nw169_[int(12)] = iife51_(self.f26)
        nw169_[int(13)] = self.f26
        nw169_[int(14)] = self.f26
        nw169_[int(15)] = self.f26
        nw169_[int(16)] = self.f26
        nw169_[int(17)] = self.f26
        nw169_[int(18)] = self.f26
        nw169_[int(19)] = self.f26
        nw169_[int(20)] = self.f26
        d_905_v21_ = nw169_
        d_908_v22_: _dafny.Array
        def lambda45_(d_909_i3_):
            return default__.safeModuloInt(d_909_i3_, (self).f23)

        init28_ = lambda45_
        nw170_ = _dafny.Array(None, 8)
        for i0_28_ in range(nw170_.length(0)):
            nw170_[i0_28_] = init28_(i0_28_)
        d_908_v22_ = nw170_
        d_910_v23_: _dafny.Map
        d_910_v23_ = _dafny.Map({D3_DC9((self).f23): d_908_v22_})
        d_911_v24_: _dafny.Array
        nw171_ = _dafny.Array(None, 13)
        nw171_[int(0)] = len(d_901_v16_)
        nw171_[int(1)] = ((self).f23) + (968)
        def iife55_():
            coll33_ = _dafny.Map()
            compr_33_: int
            for compr_33_ in (d_902_v18_).Elements:
                d_912_v17_: int = compr_33_
                if (d_912_v17_) in (d_902_v18_):
                    coll33_[default__.safeModuloInt(d_912_v17_, 529)] = self.f24
            return _dafny.Map(coll33_)
        nw171_[int(2)] = default__.fm1((0) - (len(iife55_()
        )), (self).f25, globalState)
        nw171_[int(3)] = (self).f23
        nw171_[int(4)] = (0) - (((self).f23) + ((self).f23))
        nw171_[int(5)] = default__.safeModuloInt((self).f23, (D5_DC15((self).f23, d_885_v3_, _dafny.Map({d_905_v21_: d_885_v3_}), d_910_v23_)).cf25)
        nw171_[int(6)] = (self).f23
        nw171_[int(7)] = (self).f23
        nw171_[int(8)] = (self).f23
        nw171_[int(9)] = (self).f23
        nw171_[int(10)] = ((self).f23) - ((self).f23)
        nw171_[int(11)] = (self).f23
        nw171_[int(12)] = (self).f23
        d_911_v24_ = nw171_
        d_913_v25_: _dafny.Map
        d_913_v25_ = _dafny.Map({d_911_v24_: self.f24})
        index147_ = default__.safeIndex(911, (d_886_v4_).length(0))
        rhs167_ = (d_886_v4_)[default__.safeIndex(911, (d_886_v4_).length(0))]
        rhs168_ = (d_911_v24_ if (140) != (926) else d_908_v22_)
        rhs169_ = d_913_v25_
        lhs147_ = d_886_v4_
        lhs148_ = default__.safeIndex(911, (d_886_v4_).length(0))
        lhs147_[lhs148_] = rhs167_
        d_908_v22_ = rhs168_
        d_913_v25_ = rhs169_
        hi8_ = (self).f23
        for d_914_i4_ in range(825, hi8_):
            d_887_v5_ = ((default__.fm36(globalState)).set((self).f25, len(_dafny.Map({self.f24: self.f24})))).set((self).f25, len((_dafny.Map({not((self).f25): 285})) | (d_887_v5_)))
            d_915_v26_: C4
            nw172_ = C4()
            nw172_.ctor__(not((self).f25), self.f24, not((self).f25))
            d_915_v26_ = nw172_
            r3 = (self).f23
            d_916_v27_: C4
            nw173_ = C4()
            nw173_.ctor__(((self).f25) and (not((self).f25)), (self).f25, not ((d_886_v4_)[default__.safeIndex(911, (d_886_v4_).length(0))]) or (self.f24))
            d_916_v27_ = nw173_
            d_917_v28_: _dafny.Seq
            d_917_v28_ = _dafny.SeqWithoutIsStrInference([(d_916_v27_).f27, (d_915_v26_).f27, (d_886_v4_)[default__.safeIndex(911, (d_886_v4_).length(0))]])
            rhs170_ = self.f24
            rhs171_ = (d_917_v28_)[default__.safeIndex(default__.safeDivisionInt(d_914_i4_, (self).f23), len(d_917_v28_))]
            rhs172_ = d_915_v26_
            rhs173_ = (d_916_v27_).f27
            rhs174_ = d_886_v4_
            lhs149_ = globalState
            lhs150_ = self
            r0 = rhs170_
            lhs149_.f6 = rhs171_
            d_916_v27_ = rhs172_
            lhs150_.f24 = rhs173_
            d_886_v4_ = rhs174_
        d_918_v29_: _dafny.Seq
        d_918_v29_ = _dafny.SeqWithoutIsStrInference([(d_886_v4_)[default__.safeIndex(911, (d_886_v4_).length(0))], (self).f25, self.f24])
        r0 = ((d_886_v4_)[default__.safeIndex(911, (d_886_v4_).length(0))]) in (d_918_v29_)
        d_919_v30_: T0
        nw174_ = C2()
        nw174_.ctor__((self).f23)
        d_919_v30_ = nw174_
        r1 = d_919_v30_
        r2 = (self).f25
        d_920_v31_: _dafny.Map
        d_920_v31_ = _dafny.Map({False: (d_886_v4_)[default__.safeIndex(911, (d_886_v4_).length(0))]})
        d_921_v32_: D1
        d_921_v32_ = D1_DC5(self.f24, d_920_v31_, (self).f23, (self).f23, d_887_v5_)
        d_922_v33_: _dafny.Map
        d_922_v33_ = _dafny.Map({(d_919_v30_).f23: (d_921_v32_).cf9})
        d_923_v34_: _dafny.Map
        d_923_v34_ = _dafny.Map({(d_919_v30_).f23: len(d_922_v33_)})
        r3 = default__.safeDivisionInt((d_919_v30_).f23, len((d_923_v34_) | (_dafny.Map({(d_919_v30_).f23: (self).f23}))))
        return r0, r1, r2, r3

    def m0(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: bool = False
        d_924_v0_: _dafny.Seq
        d_924_v0_ = _dafny.SeqWithoutIsStrInference([(self).f23])
        d_925_v1_: D10
        d_925_v1_ = D10_DC28((_dafny.SeqWithoutIsStrInference([(self).f23, (self).f23])).set(default__.safeIndex((self).f23, len(_dafny.SeqWithoutIsStrInference([(self).f23, (self).f23]))), 952))
        d_926_v2_: _dafny.Seq
        d_926_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itfta"))
        d_927_v3_: _dafny.Array
        nw175_ = _dafny.Array(None, 18)
        nw175_[int(0)] = d_924_v0_
        nw175_[int(1)] = d_924_v0_
        nw175_[int(2)] = d_924_v0_
        nw175_[int(3)] = _dafny.SeqWithoutIsStrInference([-698])
        nw175_[int(4)] = d_924_v0_
        nw175_[int(5)] = d_924_v0_
        nw175_[int(6)] = (d_925_v1_).cf52
        nw175_[int(7)] = d_924_v0_
        nw175_[int(8)] = (d_924_v0_).set(default__.safeIndex(len(d_926_v2_), len(d_924_v0_)), (self).f23)
        nw175_[int(9)] = default__.fm31(p1, True, globalState)
        nw175_[int(10)] = _dafny.SeqWithoutIsStrInference([(self).f23 for d_928_i1_ in range(default__.abs(180))])
        nw175_[int(11)] = d_924_v0_
        nw175_[int(12)] = d_924_v0_
        nw175_[int(13)] = d_924_v0_
        nw175_[int(14)] = d_924_v0_
        nw175_[int(15)] = (d_924_v0_) + (d_924_v0_)
        nw175_[int(16)] = d_924_v0_
        nw175_[int(17)] = _dafny.SeqWithoutIsStrInference([(self).f23])
        d_927_v3_ = nw175_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_927_v3_).length(0)):
            d_929_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_929_i0_)) and ((d_929_i0_) < ((d_927_v3_).length(0)))):
                (d_927_v3_)[(d_929_i0_)] = ((d_924_v0_) + ((d_924_v0_).set(default__.safeIndex((self).f23, len(d_924_v0_)), (self).f23))).set(default__.safeIndex((self).f23, len((d_924_v0_) + ((d_924_v0_).set(default__.safeIndex((self).f23, len(d_924_v0_)), (self).f23)))), (self).f23)
        d_930_v4_: _dafny.Array
        def lambda46_(d_931_i2_):
            return True

        init29_ = lambda46_
        nw176_ = _dafny.Array(None, 8)
        for i0_29_ in range(nw176_.length(0)):
            nw176_[i0_29_] = init29_(i0_29_)
        d_930_v4_ = nw176_
        d_930_v4_ = d_930_v4_
        (globalState).f15 = (((self).f23) + ((self).f23)) >= ((self).f23)
        d_924_v0_ = d_924_v0_
        d_932_v5_: _dafny.Map
        d_932_v5_ = _dafny.Map({(self).f23: p1})
        d_933_v6_: _dafny.Set
        d_933_v6_ = _dafny.Set({(self).f25})
        d_934_v7_: _dafny.Seq
        d_934_v7_ = _dafny.SeqWithoutIsStrInference([d_933_v6_])
        d_935_v8_: _dafny.Seq
        d_935_v8_ = _dafny.SeqWithoutIsStrInference([d_926_v2_])
        d_936_v9_: _dafny.Seq
        d_936_v9_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        d_937_v10_: _dafny.Array
        def lambda47_(d_938_i4_):
            return default__.safeModuloInt(d_938_i4_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sl"))))

        init30_ = lambda47_
        nw177_ = _dafny.Array(None, 22)
        for i0_30_ in range(nw177_.length(0)):
            nw177_[i0_30_] = init30_(i0_30_)
        d_937_v10_ = nw177_
        d_939_v11_: _dafny.MultiSet
        d_939_v11_ = _dafny.MultiSet([d_937_v10_])
        d_940_v12_: _dafny.MultiSet
        d_940_v12_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, self.f24, (self).f25, self.f24, (self).f25])])
        d_941_v13_: _dafny.MultiSet
        d_941_v13_ = _dafny.MultiSet([self.f24])
        d_942_v14_: _dafny.Set
        d_942_v14_ = _dafny.Set({d_941_v13_})
        d_943_v15_: _dafny.Array
        nw178_ = _dafny.Array(None, 21)
        nw178_[int(0)] = (0) - ((d_924_v0_)[default__.safeIndex(len(d_926_v2_), len(d_924_v0_))])
        nw178_[int(1)] = (self).f23
        nw178_[int(2)] = (self).f23
        nw178_[int(3)] = (self).f23
        nw178_[int(4)] = (self).f23
        nw178_[int(5)] = (self).f23
        nw178_[int(6)] = (0) - ((d_924_v0_)[default__.safeIndex(len(_dafny.Map({p1: d_932_v5_})), len(d_924_v0_))])
        nw178_[int(7)] = (self).f23
        nw178_[int(8)] = (0) - (len(d_934_v7_))
        nw178_[int(9)] = (self).f23
        nw178_[int(10)] = default__.safeModuloInt((self).f23, (0) - (len((d_935_v8_)[default__.safeIndex((self).f23, len(d_935_v8_))])))
        nw178_[int(11)] = default__.safeDivisionInt((self).f23, (self).f23)
        nw178_[int(12)] = default__.fm1((self).f23, (self).f25, globalState)
        nw178_[int(13)] = ((d_939_v11_ if (d_936_v9_)[default__.safeIndex((self).f23, len(d_936_v9_))] else d_939_v11_)).cardinality
        nw178_[int(14)] = ((self).f23) - ((self).f23)
        nw178_[int(15)] = (self).f23
        nw178_[int(16)] = default__.safeModuloInt(len(d_926_v2_), 996)
        nw178_[int(17)] = ((d_940_v12_)[d_936_v9_] if (d_936_v9_) in (d_940_v12_) else (self).f23)
        nw178_[int(18)] = default__.safeDivisionInt((self).f23, (self).f23)
        nw178_[int(19)] = len(d_942_v14_)
        nw178_[int(20)] = (self).f23
        d_943_v15_ = nw178_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_937_v10_).length(0)):
            d_944_i3_: int = guard_loop_4_
            if (True) and (((0) <= (d_944_i3_)) and ((d_944_i3_) < ((d_937_v10_).length(0)))):
                (d_937_v10_)[(d_944_i3_)] = (d_944_i3_) * (((self).f23) + (len(d_924_v0_)))
        hi9_ = (d_924_v0_)[default__.safeIndex((self).f23, len(d_924_v0_))]
        for d_945_i5_ in range((len(d_933_v6_)) + ((self).f23), hi9_):
            (globalState).f14 = (default__.safeDivisionInt(d_945_i5_, d_945_i5_)) + ((self).f23)
            d_946_v16_: _dafny.Map
            d_946_v16_ = _dafny.Map({_dafny.CodePoint('i'): p0})
            d_947_v17_: _dafny.MultiSet
            d_947_v17_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_948_i6_ in range(default__.abs(757))])), (default__.fm37((self).f25, len(d_946_v16_), globalState)).cardinality, d_945_i5_, d_945_i5_, (self).f23])
            d_949_v18_: _dafny.Set
            d_949_v18_ = _dafny.Set({(d_947_v17_).set((self).f23, default__.abs((self).f23))})
            d_950_v19_: _dafny.Map
            d_950_v19_ = _dafny.Map({p0: _dafny.Set({False, not(p1)})})
            d_951_v20_: _dafny.Array
            nw179_ = _dafny.Array(None, 20)
            nw179_[int(0)] = (d_933_v6_) - (d_933_v6_)
            nw179_[int(1)] = d_933_v6_
            nw179_[int(2)] = d_933_v6_
            nw179_[int(3)] = _dafny.Set({p1, p1, (self).f25, p1})
            nw179_[int(4)] = (_dafny.Set({self.f24})) - (d_933_v6_)
            nw179_[int(5)] = (d_933_v6_ if self.f24 else d_933_v6_)
            nw179_[int(6)] = (d_933_v6_) | (d_933_v6_)
            nw179_[int(7)] = (d_933_v6_) | (_dafny.Set({self.f24}))
            nw179_[int(8)] = d_933_v6_
            nw179_[int(9)] = d_933_v6_
            nw179_[int(10)] = d_933_v6_
            nw179_[int(11)] = d_933_v6_
            nw179_[int(12)] = d_933_v6_
            nw179_[int(13)] = (d_933_v6_).intersection(d_933_v6_)
            nw179_[int(14)] = d_933_v6_
            nw179_[int(15)] = (d_933_v6_) | (d_933_v6_)
            nw179_[int(16)] = (_dafny.Set({p1, (self).f25, self.f24, self.f24})).intersection(((d_950_v19_)[(self).f25] if ((self).f25) in (d_950_v19_) else d_933_v6_))
            nw179_[int(17)] = _dafny.Set({p0, p0, p0, self.f24, (self).fm3(d_926_v2_, (self).f25, d_945_i5_, d_924_v0_, globalState)})
            nw179_[int(18)] = d_933_v6_
            nw179_[int(19)] = d_933_v6_
            d_951_v20_ = nw179_
            index148_ = default__.safeIndex(234, (d_951_v20_).length(0))
            (d_951_v20_)[index148_] = d_933_v6_
            index149_ = default__.safeIndex(234, (d_951_v20_).length(0))
            rhs175_ = d_949_v18_
            rhs176_ = default__.safeDivisionInt(d_945_i5_, (0) - (default__.fm1((0) - (d_945_i5_), p0, globalState)))
            rhs177_ = _dafny.Set({p1, (d_945_i5_) < (d_945_i5_), not(False)})
            lhs151_ = globalState
            lhs152_ = d_951_v20_
            lhs153_ = default__.safeIndex(234, (d_951_v20_).length(0))
            d_949_v18_ = rhs175_
            lhs151_.f8 = rhs176_
            lhs152_[lhs153_] = rhs177_
            if p0:
                d_952_v21_: str
                d_952_v21_ = _dafny.CodePoint('y')
                d_953_v22_: _dafny.MultiSet
                d_953_v22_ = _dafny.MultiSet([d_947_v17_])
                d_954_v23_: _dafny.Map
                d_954_v23_ = _dafny.Map({(len(_dafny.SeqWithoutIsStrInference([d_952_v21_]))) * ((self).f23): d_953_v22_})
                d_954_v23_ = (d_954_v23_).set((((d_947_v17_)[len(d_936_v9_)] if (len(d_936_v9_)) in (d_947_v17_) else (self).f23)) + (895), d_953_v22_)
                d_955_v24_: _dafny.Set
                d_955_v24_ = _dafny.Set({d_930_v4_})
                d_956_v25_: C0
                nw180_ = C0()
                nw180_.ctor__((d_955_v24_) - (d_955_v24_), (self).f23, d_945_i5_)
                d_956_v25_ = nw180_
                (globalState).f17 = len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, self.f24]), (d_936_v9_) + (d_936_v9_), d_936_v9_]))
                d_957_v26_: _dafny.Array
                nw181_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_957_v26_ = nw181_
                index150_ = default__.safeIndex(214, (d_957_v26_).length(0))
                (d_957_v26_)[index150_] = d_937_v10_
                index151_ = default__.safeIndex(214, (d_957_v26_).length(0))
                def lambda48_(d_958_v25_):
                    def lambda49_(d_959_i7_):
                        return (d_959_i7_) * ((0) - (d_958_v25_.f31))

                    return lambda49_

                init31_ = lambda48_(d_956_v25_)
                nw182_ = _dafny.Array(None, 28)
                for i0_31_ in range(nw182_.length(0)):
                    nw182_[i0_31_] = init31_(i0_31_)
                (d_957_v26_)[index151_] = nw182_
                d_930_v4_ = d_930_v4_
            elif True:
                d_926_v2_ = (d_926_v2_) + (d_926_v2_)
                d_960_v27_: _dafny.Set
                d_960_v27_ = _dafny.Set({d_933_v6_, d_933_v6_})
                (globalState).f9 = ((_dafny.Set({(d_951_v20_)[default__.safeIndex(234, (d_951_v20_).length(0))], (d_951_v20_)[default__.safeIndex(234, (d_951_v20_).length(0))]}) if p1 else default__.fm38((self).f23, (0) - (d_945_i5_), globalState))).isdisjoint(d_960_v27_)
                (globalState).f6 = (d_936_v9_)[default__.safeIndex(981, len(d_936_v9_))]
                (globalState).f15 = (self).fm3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvwh")), p0, default__.safeDivisionInt(len(d_924_v0_), d_945_i5_), d_924_v0_, globalState)
                d_961_v28_: _dafny.Map
                d_961_v28_ = _dafny.Map({d_936_v9_: d_945_i5_})
                d_962_v29_: C3
                nw183_ = C3()
                nw183_.ctor__(p0, len(d_961_v28_), p0, (d_945_i5_) > (d_945_i5_), d_945_i5_)
                d_962_v29_ = nw183_
            (globalState).f14 = len((d_936_v9_) + ((d_936_v9_ if self.f24 else _dafny.SeqWithoutIsStrInference([self.f24, self.f24]))))
        r0 = (self).f23
        r1 = d_943_v15_
        r2 = (self).f25
        return r0, r1, r2

    def m4(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_963_v1_: _dafny.Seq
        d_963_v1_ = _dafny.SeqWithoutIsStrInference([p3, _dafny.CodePoint('u')])
        def iife56_():
            coll34_ = _dafny.Map()
            compr_34_: int
            for compr_34_ in _dafny.IntegerRange(-489, -153):
                d_964_v0_: int = compr_34_
                if ((-489) <= (d_964_v0_)) and ((d_964_v0_) < (-153)):
                    coll34_[(d_964_v0_) + (len(_dafny.SeqWithoutIsStrInference([(self).f23, p1, p1])))] = (self).f25
            return _dafny.Map(coll34_)
        rhs178_ = (self).fm2(globalState)
        rhs179_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")))) <= ((len(iife56_()
        )) - (len((d_963_v1_).set(default__.safeIndex((self).f23, len(d_963_v1_)), _dafny.CodePoint('q')))))
        lhs154_ = globalState
        lhs155_ = globalState
        lhs154_.f8 = rhs178_
        lhs155_.f9 = rhs179_
        d_965_v2_: _dafny.Map
        d_965_v2_ = _dafny.Map({_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhkyjg"))), (self).f23}): (self).f23})
        d_966_v3_: _dafny.MultiSet
        d_966_v3_ = _dafny.MultiSet([(self).f23, p1])
        d_967_v4_: _dafny.Map
        d_967_v4_ = _dafny.Map({p2: (self).f23})
        d_968_v5_: _dafny.Array
        nw184_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
        d_968_v5_ = nw184_
        d_969_v6_: D7
        d_969_v6_ = D7_DC21(D7_DC20(self.f24, d_967_v4_, d_968_v5_))
        d_970_v7_: D7
        d_970_v7_ = D7_DC21(d_969_v6_)
        d_971_v8_: _dafny.Seq
        d_971_v8_ = _dafny.SeqWithoutIsStrInference([p1])
        d_972_v9_: _dafny.Map
        d_972_v9_ = _dafny.Map({d_970_v7_: (d_971_v8_)[default__.safeIndex(p1, len(d_971_v8_))]})
        d_973_v10_: _dafny.Seq
        d_973_v10_ = _dafny.SeqWithoutIsStrInference([((d_972_v9_)[d_970_v7_] if (d_970_v7_) in (d_972_v9_) else (self).f23)])
        (globalState).f1 = (self).fm3((d_963_v1_) + ((d_963_v1_).set(default__.safeIndex((self).fm5(p1, len(d_965_v2_), globalState), len(d_963_v1_)), p3)), (self).f25, (d_966_v3_).cardinality, d_973_v10_, globalState)
        if (754) == ((self).f23):
            d_974_v11_: C5
            nw185_ = C5()
            nw185_.ctor__((self).f25, (self).f25)
            d_974_v11_ = nw185_
            d_974_v11_ = d_974_v11_
            d_975_v12_: _dafny.Set
            d_975_v12_ = _dafny.Set({(self).f25})
            source8_ = default__.fm39(p2, (p1) - ((self).f23), True, (_dafny.Set({p2, self.f24, (self).f25, p2, True})) - (d_975_v12_), globalState)
            if source8_.is_DC9:
                d_976___mcc_h0_ = source8_.cf15
                d_977_cf15_ = d_976___mcc_h0_
                (globalState).f8 = (((_dafny.MultiSet([(self).f23, d_977_cf15_])).intersection(d_966_v3_)).cardinality) * (d_977_cf15_)
                d_978_v13_: _dafny.Map
                d_978_v13_ = _dafny.Map({_dafny.CodePoint('v'): (0) - ((self).f23)})
                d_979_v14_: _dafny.Array
                nw186_ = _dafny.Array(None, 28)
                nw186_[int(0)] = p3
                nw186_[int(1)] = p3
                nw186_[int(2)] = p3
                nw186_[int(3)] = p3
                nw186_[int(4)] = p3
                nw186_[int(5)] = p3
                nw186_[int(6)] = p3
                nw186_[int(7)] = p3
                nw186_[int(8)] = p3
                nw186_[int(9)] = p3
                nw186_[int(10)] = p3
                nw186_[int(11)] = _dafny.CodePoint('t')
                nw186_[int(12)] = p3
                nw186_[int(13)] = p3
                nw186_[int(14)] = p3
                nw186_[int(15)] = p3
                nw186_[int(16)] = p3
                nw186_[int(17)] = _dafny.CodePoint('n')
                nw186_[int(18)] = p3
                nw186_[int(19)] = p3
                nw186_[int(20)] = p3
                nw186_[int(21)] = p3
                nw186_[int(22)] = p3
                nw186_[int(23)] = p3
                nw186_[int(24)] = p3
                nw186_[int(25)] = _dafny.CodePoint('g')
                nw186_[int(26)] = p3
                nw186_[int(27)] = p3
                d_979_v14_ = nw186_
                d_980_v15_: _dafny.Set
                d_980_v15_ = _dafny.Set({d_979_v14_})
                d_981_v16_: D5
                d_981_v16_ = D5_DC14((self).f23, p2, d_980_v15_)
                d_982_v17_: _dafny.Map
                d_982_v17_ = _dafny.Map({(self).f25: p2})
                d_983_v18_: _dafny.Array
                nw187_ = _dafny.Array(None, 16)
                nw187_[int(0)] = self.f24
                nw187_[int(1)] = (d_981_v16_).cf23
                nw187_[int(2)] = (self).f25
                nw187_[int(3)] = (self).fm3(d_963_v1_, self.f24, len(d_967_v4_), d_971_v8_, globalState)
                nw187_[int(4)] = False
                nw187_[int(5)] = ((d_982_v17_)[(self).f25] if ((self).f25) in (d_982_v17_) else self.f24)
                nw187_[int(6)] = True
                nw187_[int(7)] = p2
                nw187_[int(8)] = (self).f25
                nw187_[int(9)] = p2
                nw187_[int(10)] = self.f24
                nw187_[int(11)] = self.f24
                nw187_[int(12)] = (self).fm3(d_963_v1_, self.f24, (self).f23, d_973_v10_, globalState)
                nw187_[int(13)] = self.f24
                nw187_[int(14)] = (self).f25
                nw187_[int(15)] = (self).f25
                d_983_v18_ = nw187_
                d_984_v19_: _dafny.Array
                nw188_ = _dafny.Array(None, 26)
                nw188_[int(0)] = (self).f23
                nw188_[int(1)] = (_dafny.MultiSet([p1])).cardinality
                nw188_[int(2)] = d_977_cf15_
                nw188_[int(3)] = d_977_cf15_
                nw188_[int(4)] = p1
                nw188_[int(5)] = 283
                nw188_[int(6)] = (self).f23
                nw188_[int(7)] = p1
                nw188_[int(8)] = (self).f23
                nw188_[int(9)] = d_977_cf15_
                nw188_[int(10)] = (self).f23
                nw188_[int(11)] = 406
                nw188_[int(12)] = (self).f23
                nw188_[int(13)] = p1
                nw188_[int(14)] = len(d_978_v13_)
                nw188_[int(15)] = d_977_cf15_
                nw188_[int(16)] = p1
                nw188_[int(17)] = len(_dafny.Map({self.f24: p2}))
                nw188_[int(18)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "huhwbkh")))
                nw188_[int(19)] = len(d_963_v1_)
                nw188_[int(20)] = d_977_cf15_
                nw188_[int(21)] = len(d_973_v10_)
                nw188_[int(22)] = (self).f23
                nw188_[int(23)] = len(_dafny.SeqWithoutIsStrInference([d_983_v18_]))
                nw188_[int(24)] = (self).f23
                nw188_[int(25)] = p1
                d_984_v19_ = nw188_
                d_985_v20_: _dafny.MultiSet
                d_985_v20_ = _dafny.MultiSet([d_984_v19_])
                d_986_v21_: _dafny.Map
                d_986_v21_ = _dafny.Map({d_985_v20_: (d_974_v11_).fm5((self).f23, d_977_cf15_, globalState)})
                d_986_v21_ = (d_986_v21_).set((D12_DC32(d_985_v20_)).cf60, ((self).f23) + (d_977_cf15_))
                (globalState).f9 = True
                (globalState).f1 = (d_977_cf15_) != ((self).f23)
            elif source8_.is_DC10:
                d_987___mcc_h1_ = source8_.cf16
                d_988___mcc_h2_ = source8_.cf17
                d_989_cf17_ = d_988___mcc_h2_
                d_990_cf16_ = d_987___mcc_h1_
                (globalState).f8 = default__.safeDivisionInt((self).f23, (self).f23)
                (globalState).f14 = p1
                rhs180_ = p1
                rhs181_ = (self).fm2(globalState)
                rhs182_ = (False) or (self.f24)
                rhs183_ = p1
                rhs184_ = p2
                lhs156_ = globalState
                lhs157_ = globalState
                lhs158_ = self
                lhs159_ = globalState
                lhs160_ = globalState
                lhs156_.f5 = rhs180_
                lhs157_.f5 = rhs181_
                lhs158_.f24 = rhs182_
                lhs159_.f17 = rhs183_
                lhs160_.f15 = rhs184_
                d_991_v22_: C1
                nw189_ = C1()
                nw189_.ctor__(p2, False)
                d_991_v22_ = nw189_
            elif True:
                d_992___mcc_h3_ = source8_.cf14
                d_993_cf14_ = d_992___mcc_h3_
                d_994_v23_: _dafny.Array
                def lambda50_(d_995_i0_):
                    return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f24]))).intersection(_dafny.MultiSet([(self).f25]))

                init32_ = lambda50_
                nw190_ = _dafny.Array(None, 24)
                for i0_32_ in range(nw190_.length(0)):
                    nw190_[i0_32_] = init32_(i0_32_)
                d_994_v23_ = nw190_
                d_996_v24_: _dafny.MultiSet
                d_996_v24_ = _dafny.MultiSet([p2])
                index152_ = default__.safeIndex(331, (d_994_v23_).length(0))
                (d_994_v23_)[index152_] = d_996_v24_
                d_997_v25_: D7
                d_997_v25_ = D7_DC19(d_996_v24_)
                index153_ = default__.safeIndex(331, (d_994_v23_).length(0))
                (d_994_v23_)[index153_] = (d_997_v25_).cf34
                (globalState).f8 = (_dafny.MultiSet([p2, p2])).cardinality
                d_998_v26_: _dafny.Seq
                d_998_v26_ = _dafny.SeqWithoutIsStrInference([d_968_v5_, d_968_v5_, d_968_v5_, d_968_v5_])
                d_968_v5_ = (d_998_v26_)[default__.safeIndex((self).f23, len(d_998_v26_))]
                (globalState).f17 = p1
            d_999_v27_: _dafny.Seq
            d_999_v27_ = _dafny.SeqWithoutIsStrInference([d_966_v3_])
            rhs185_ = ((d_999_v27_)[default__.safeIndex((self).f23, len(d_999_v27_))]) - (d_966_v3_)
            rhs186_ = p2
            rhs187_ = p2
            lhs161_ = globalState
            lhs162_ = globalState
            d_966_v3_ = rhs185_
            lhs161_.f1 = rhs186_
            lhs162_.f1 = rhs187_
            (globalState).f17 = (self).f23
            d_1000_v28_: str
            d_1000_v28_ = _dafny.CodePoint('m')
            d_1000_v28_ = d_1000_v28_
        elif True:
            (globalState).f15 = p2
            (globalState).f17 = (0) - (p1)
            (globalState).f8 = (self).f23
            (globalState).f8 = 129
            (globalState).f8 = ((_dafny.MultiSet([(self).f25, (self).f25, not(True)])).set((self).f25, default__.abs(default__.safeDivisionInt((self).f23, (self).f23)))).cardinality
        (globalState).f15 = self.f24
        (globalState).f9 = (self.f24 if False else ((self).fm3(d_963_v1_, p2, (self).f23, d_971_v8_, globalState) if (self).f25 else p2))
        d_1001_v29_: _dafny.Array
        nw191_ = _dafny.Array(_dafny.CodePoint('D'), 16)
        d_1001_v29_ = nw191_
        d_1002_v30_: _dafny.Set
        d_1002_v30_ = _dafny.Set({d_1001_v29_, d_1001_v29_})
        d_1003_v31_: D5
        d_1003_v31_ = D5_DC14((self).f23, self.f24, d_1002_v30_)
        d_1004_v32_: _dafny.MultiSet
        d_1004_v32_ = _dafny.MultiSet([d_1003_v31_, d_1003_v31_, d_1003_v31_, d_1003_v31_, d_1003_v31_])
        (globalState).f14 = ((d_1004_v32_).cardinality if (self).f25 else p1)
        d_1005_v33_: _dafny.Array
        nw192_ = _dafny.Array(None, 2)
        nw192_[int(0)] = self.f24
        nw192_[int(1)] = self.f24
        d_1005_v33_ = nw192_
        r0 = d_1005_v33_
        return r0


class C7:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self):
        pass
        pass

    def fm7(self, globalState):
        return (_dafny.Map({True: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, not(True), False, not(False), not(True)]))})) | (_dafny.Map({True: _dafny.MultiSet([False, False, True, True])}))

    def m2(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        r2: D0 = D0.default()()
        r0 = p1
        d_1006_v0_: _dafny.Seq
        d_1006_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrd"))
        d_1007_v1_: D0
        d_1007_v1_ = D0_DC1(p1, p0)
        pat_let_tv15_ = d_1007_v1_
        pat_let_tv16_ = p0
        pat_let_tv17_ = p0
        def lambda51_(source9_):
            if source9_.is_DC1:
                d_1008___mcc_h0_ = source9_.cf1
                d_1009___mcc_h1_ = source9_.cf2
                d_1010_cf2_ = d_1009___mcc_h1_
                d_1011_cf1_ = d_1008___mcc_h0_
                return (d_1010_cf2_) - (d_1010_cf2_)
            elif source9_.is_DC0:
                d_1012___mcc_h2_ = source9_.cf0
                d_1013_cf0_ = d_1012___mcc_h2_
                return (pat_let_tv15_).cf2
            elif True:
                d_1014___mcc_h3_ = source9_.cf3
                d_1015_cf3_ = d_1014___mcc_h3_
                return default__.safeDivisionInt(pat_let_tv16_, pat_let_tv17_)

        rhs188_ = (len(d_1006_v0_)) - (p0)
        rhs189_ = (0) - (lambda51_(d_1007_v1_))
        lhs163_ = globalState
        lhs164_ = globalState
        lhs163_.f8 = rhs188_
        lhs164_.f8 = rhs189_
        d_1016_v2_: _dafny.Array
        nw193_ = _dafny.Array(_dafny.Seq({}), 14)
        d_1016_v2_ = nw193_
        d_1017_v3_: _dafny.Array
        nw194_ = _dafny.Array(False, 5)
        d_1017_v3_ = nw194_
        d_1018_v4_: _dafny.Seq
        d_1018_v4_ = _dafny.SeqWithoutIsStrInference([d_1017_v3_, d_1017_v3_, d_1017_v3_])
        index154_ = default__.safeIndex(399, (d_1016_v2_).length(0))
        (d_1016_v2_)[index154_] = (d_1018_v4_) + (d_1018_v4_)
        index155_ = default__.safeIndex(399, (d_1016_v2_).length(0))
        (d_1016_v2_)[index155_] = (d_1018_v4_) + (d_1018_v4_)
        d_1019_v5_: _dafny.Seq
        d_1019_v5_ = _dafny.SeqWithoutIsStrInference([p0])
        (globalState).f8 = default__.fm1(len(default__.fm8(len(d_1019_v5_), globalState)), (len(d_1006_v0_)) == (639), globalState)
        d_1020_v7_: _dafny.Array
        def lambda52_(d_1021_v5_, d_1022_p1_):
            def lambda53_(d_1023_i0_):
                def iife57_():
                    coll35_ = _dafny.Map()
                    compr_35_: _dafny.Seq
                    for compr_35_ in (_dafny.MultiSet([d_1021_v5_, d_1021_v5_])).Elements:
                        d_1024_v6_: _dafny.Seq = compr_35_
                        if (d_1024_v6_) in (_dafny.MultiSet([d_1021_v5_, d_1021_v5_])):
                            coll35_[d_1024_v6_] = _dafny.SeqWithoutIsStrInference([d_1022_p1_, d_1022_p1_])
                    return _dafny.Map(coll35_)
                return (_dafny.Map({d_1021_v5_: _dafny.SeqWithoutIsStrInference([d_1022_p1_, d_1022_p1_])})) | (iife57_()
                )

            return lambda53_

        init33_ = lambda52_(d_1019_v5_, p1)
        nw195_ = _dafny.Array(None, 10)
        for i0_33_ in range(nw195_.length(0)):
            nw195_[i0_33_] = init33_(i0_33_)
        d_1020_v7_ = nw195_
        d_1025_v8_: _dafny.Seq
        d_1025_v8_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1026_v9_: _dafny.Map
        d_1026_v9_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0, p0, p0]): d_1025_v8_})
        index156_ = default__.safeIndex(220, (d_1020_v7_).length(0))
        (d_1020_v7_)[index156_] = (d_1026_v9_) | (d_1026_v9_)
        index157_ = default__.safeIndex(220, (d_1020_v7_).length(0))
        (d_1020_v7_)[index157_] = d_1026_v9_
        index158_ = default__.safeIndex(117, (d_1017_v3_).length(0))
        (d_1017_v3_)[index158_] = p1
        index159_ = default__.safeIndex(117, (d_1017_v3_).length(0))
        (d_1017_v3_)[index159_] = (p1 if False else not (default__.fm9(p1, globalState)) or (default__.fm9(True, globalState)))
        r0 = p1
        r1 = default__.fm9(p1, globalState)
        r2 = d_1007_v1_
        return r0, r1, r2

    def m3(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_1027_v0_: bool
        d_1027_v0_ = False
        (globalState).f15 = d_1027_v0_
        d_1028_v1_: str
        d_1028_v1_ = _dafny.CodePoint('l')
        d_1029_v2_: _dafny.Map
        d_1029_v2_ = _dafny.Map({d_1027_v0_: _dafny.CodePoint('i')})
        d_1030_v3_: D1
        d_1030_v3_ = D1_DC3(d_1028_v1_)
        d_1031_v4_: _dafny.Array
        nw196_ = _dafny.Array(None, 14)
        nw196_[int(0)] = d_1028_v1_
        nw196_[int(1)] = d_1028_v1_
        nw196_[int(2)] = ((d_1029_v2_)[d_1027_v0_] if (d_1027_v0_) in (d_1029_v2_) else d_1028_v1_)
        nw196_[int(3)] = d_1028_v1_
        nw196_[int(4)] = d_1028_v1_
        nw196_[int(5)] = d_1028_v1_
        nw196_[int(6)] = d_1028_v1_
        nw196_[int(7)] = _dafny.CodePoint('o')
        nw196_[int(8)] = (d_1030_v3_).cf4
        nw196_[int(9)] = d_1028_v1_
        nw196_[int(10)] = (d_1028_v1_ if d_1027_v0_ else d_1028_v1_)
        nw196_[int(11)] = ((d_1029_v2_)[d_1027_v0_] if (d_1027_v0_) in (d_1029_v2_) else d_1028_v1_)
        nw196_[int(12)] = d_1028_v1_
        nw196_[int(13)] = _dafny.CodePoint('g')
        d_1031_v4_ = nw196_
        index160_ = default__.safeIndex(219, (d_1031_v4_).length(0))
        (d_1031_v4_)[index160_] = d_1028_v1_
        d_1032_v5_: _dafny.Map
        d_1032_v5_ = _dafny.Map({d_1027_v0_: d_1027_v0_})
        pat_let_tv18_ = d_1027_v0_
        pat_let_tv19_ = d_1027_v0_
        pat_let_tv20_ = d_1027_v0_
        pat_let_tv21_ = d_1027_v0_
        index161_ = default__.safeIndex(219, (d_1031_v4_).length(0))
        def lambda54_(source10_):
            if source10_.is_DC4:
                d_1033___mcc_h0_ = source10_.cf5
                d_1034___mcc_h1_ = source10_.cf6
                d_1035_cf6_ = d_1034___mcc_h1_
                d_1036_cf5_ = d_1033___mcc_h0_
                return pat_let_tv18_
            elif source10_.is_DC5:
                d_1037___mcc_h2_ = source10_.cf7
                d_1038___mcc_h3_ = source10_.cf8
                d_1039___mcc_h4_ = source10_.cf9
                d_1040___mcc_h5_ = source10_.cf10
                d_1041___mcc_h6_ = source10_.cf11
                d_1042_cf11_ = d_1041___mcc_h6_
                d_1043_cf10_ = d_1040___mcc_h5_
                d_1044_cf9_ = d_1039___mcc_h4_
                d_1045_cf8_ = d_1038___mcc_h3_
                d_1046_cf7_ = d_1037___mcc_h2_
                return pat_let_tv19_
            elif source10_.is_DC3:
                d_1047___mcc_h7_ = source10_.cf4
                d_1048_cf4_ = d_1047___mcc_h7_
                return pat_let_tv20_
            elif True:
                d_1049___mcc_h8_ = source10_.cf12
                d_1050_cf12_ = d_1049___mcc_h8_
                return pat_let_tv21_

        rhs190_ = d_1028_v1_
        rhs191_ = lambda54_(D1_DC3(d_1028_v1_))
        rhs192_ = len((_dafny.Map({d_1027_v0_: d_1027_v0_})) | (d_1032_v5_))
        rhs193_ = False
        lhs165_ = d_1031_v4_
        lhs166_ = default__.safeIndex(219, (d_1031_v4_).length(0))
        lhs167_ = globalState
        lhs168_ = globalState
        lhs169_ = globalState
        lhs165_[lhs166_] = rhs190_
        lhs167_.f15 = rhs191_
        lhs168_.f8 = rhs192_
        lhs169_.f1 = rhs193_
        d_1051_i0_: int
        d_1051_i0_ = 0
        with _dafny.label("8"):
            while d_1027_v0_:
                with _dafny.c_label("8"):
                    if (d_1051_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_1051_i0_ = (d_1051_i0_) + (1)
                    d_1052_v6_: _dafny.Seq
                    d_1052_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asg"))
                    d_1052_v6_ = d_1052_v6_
                    (globalState).f8 = default__.safeDivisionInt(p2, p0)
                    d_1053_v7_: _dafny.Map
                    d_1053_v7_ = _dafny.Map({d_1027_v0_: p0})
                    d_1054_v8_: D1
                    d_1054_v8_ = D1_DC5(d_1027_v0_, (_dafny.Map({d_1027_v0_: d_1027_v0_})) | (d_1032_v5_), (p0) + (p1), 357, d_1053_v7_)
                    source11_ = d_1054_v8_
                    if source11_.is_DC4:
                        d_1055___mcc_h9_ = source11_.cf5
                        d_1056___mcc_h10_ = source11_.cf6
                        d_1057_cf6_ = d_1056___mcc_h10_
                        d_1058_cf5_ = d_1055___mcc_h9_
                        d_1059_v9_: _dafny.Seq
                        d_1059_v9_ = _dafny.SeqWithoutIsStrInference([d_1058_cf5_])
                        d_1060_v10_: _dafny.Seq
                        d_1060_v10_ = _dafny.SeqWithoutIsStrInference([default__.fm1(p0, d_1027_v0_, globalState), p0])
                        d_1061_v11_: _dafny.Seq
                        d_1061_v11_ = _dafny.SeqWithoutIsStrInference([len(d_1060_v10_), p1, len(d_1059_v9_), p1])
                        d_1062_v12_: _dafny.Seq
                        d_1062_v12_ = _dafny.SeqWithoutIsStrInference([(d_1059_v9_)[default__.safeIndex(len(d_1061_v11_), len(d_1059_v9_))], d_1058_cf5_])
                        d_1062_v12_ = d_1062_v12_
                        d_1028_v1_ = (d_1028_v1_ if not((d_1058_cf5_) and (d_1027_v0_)) else _dafny.CodePoint('b'))
                        d_1063_v13_: _dafny.Map
                        d_1063_v13_ = _dafny.Map({p0: p2})
                        d_1063_v13_ = (d_1063_v13_).set(len(d_1052_v6_), p1)
                        (globalState).f14 = (len(default__.fm10(d_1027_v0_, len(d_1062_v12_), d_1028_v1_, (d_1031_v4_)[default__.safeIndex(219, (d_1031_v4_).length(0))], globalState)) if d_1057_cf6_ else p1)
                    elif source11_.is_DC5:
                        d_1064___mcc_h11_ = source11_.cf7
                        d_1065___mcc_h12_ = source11_.cf8
                        d_1066___mcc_h13_ = source11_.cf9
                        d_1067___mcc_h14_ = source11_.cf10
                        d_1068___mcc_h15_ = source11_.cf11
                        d_1069_cf11_ = d_1068___mcc_h15_
                        d_1070_cf10_ = d_1067___mcc_h14_
                        d_1071_cf9_ = d_1066___mcc_h13_
                        d_1072_cf8_ = d_1065___mcc_h12_
                        d_1073_cf7_ = d_1064___mcc_h11_
                        d_1074_v14_: _dafny.Array
                        def lambda55_(d_1075_v0_):
                            def lambda56_(d_1076_i1_):
                                return d_1075_v0_

                            return lambda56_

                        init34_ = lambda55_(d_1027_v0_)
                        nw197_ = _dafny.Array(None, 9)
                        for i0_34_ in range(nw197_.length(0)):
                            nw197_[i0_34_] = init34_(i0_34_)
                        d_1074_v14_ = nw197_
                        d_1077_v15_: _dafny.Set
                        d_1077_v15_ = _dafny.Set({d_1074_v14_})
                        d_1078_v16_: _dafny.Set
                        d_1078_v16_ = d_1077_v15_
                        d_1079_v17_: C0
                        nw198_ = C0()
                        nw198_.ctor__((d_1077_v15_).intersection((d_1078_v16_)), 675, (0) - (d_1071_cf9_))
                        d_1079_v17_ = nw198_
                        d_1080_v18_: _dafny.Set
                        d_1080_v18_ = _dafny.Set({d_1071_cf9_})
                        d_1081_v19_: _dafny.Seq
                        d_1081_v19_ = _dafny.SeqWithoutIsStrInference([d_1080_v18_])
                        rhs194_ = len(d_1081_v19_)
                        rhs195_ = d_1027_v0_
                        rhs196_ = not(d_1073_cf7_)
                        lhs170_ = globalState
                        lhs171_ = globalState
                        lhs172_ = globalState
                        lhs170_.f8 = rhs194_
                        lhs171_.f1 = rhs195_
                        lhs172_.f6 = rhs196_
                        rhs197_ = default__.safeDivisionInt((-321) + (p2), default__.safeModuloInt(d_1071_cf9_, 319))
                        rhs198_ = d_1028_v1_
                        lhs173_ = globalState
                        lhs173_.f14 = rhs197_
                        d_1028_v1_ = rhs198_
                        d_1082_v20_: _dafny.Seq
                        d_1082_v20_ = _dafny.SeqWithoutIsStrInference([d_1071_cf9_])
                        d_1083_v21_: C5
                        nw199_ = C5()
                        nw199_.ctor__((d_1079_v17_).fm3(d_1052_v6_, d_1027_v0_, p1, d_1082_v20_, globalState), not(d_1027_v0_))
                        d_1083_v21_ = nw199_
                        d_1084_v22_: _dafny.MultiSet
                        d_1084_v22_ = _dafny.MultiSet([d_1083_v21_])
                        d_1085_v23_: _dafny.Map
                        d_1085_v23_ = _dafny.Map({(0) - ((p1) + (d_1070_cf10_)): ((d_1084_v22_) - (d_1084_v22_)).cardinality})
                        d_1086_v25_: _dafny.MultiSet
                        d_1086_v25_ = _dafny.MultiSet([p2, d_1079_v17_.f31])
                        def iife58_():
                            coll36_ = _dafny.Map()
                            compr_36_: int
                            for compr_36_ in ((d_1086_v25_).set((d_1083_v21_).fm13(d_1070_cf10_, d_1027_v0_, globalState), default__.abs(len(d_1052_v6_)))).Elements:
                                d_1087_v24_: int = compr_36_
                                if (d_1087_v24_) in ((d_1086_v25_).set((d_1083_v21_).fm13(d_1070_cf10_, d_1027_v0_, globalState), default__.abs(len(d_1052_v6_)))):
                                    coll36_[(d_1087_v24_) + (d_1071_cf9_)] = default__.safeModuloInt(p0, p0)
                            return _dafny.Map(coll36_)
                        d_1085_v23_ = iife58_()
                        
                    elif source11_.is_DC3:
                        d_1088___mcc_h16_ = source11_.cf4
                        d_1089_cf4_ = d_1088___mcc_h16_
                        (globalState).f17 = default__.fm1(p1, True, globalState)
                        d_1090_v26_: _dafny.Set
                        d_1090_v26_ = _dafny.Set({p0, p0})
                        d_1091_v27_: _dafny.Map
                        d_1091_v27_ = _dafny.Map({len(d_1090_v26_): default__.fm9(d_1027_v0_, globalState)})
                        d_1092_v28_: _dafny.MultiSet
                        d_1092_v28_ = _dafny.MultiSet([default__.fm1(len(d_1090_v26_), d_1027_v0_, globalState)])
                        d_1093_v29_: _dafny.Map
                        d_1093_v29_ = _dafny.Map({d_1091_v27_: (d_1092_v28_).issubset(d_1092_v28_)})
                        d_1093_v29_ = ((d_1093_v29_) | (d_1093_v29_)) | (d_1093_v29_)
                        (globalState).f8 = (0) - (p0)
                        d_1032_v5_ = (d_1032_v5_) | (default__.fm10(d_1027_v0_, p2, d_1089_cf4_, d_1028_v1_, globalState))
                    elif True:
                        d_1094___mcc_h17_ = source11_.cf12
                        d_1095_cf12_ = d_1094___mcc_h17_
                        d_1096_v30_: _dafny.Array
                        def lambda57_(d_1097_p1_):
                            def lambda58_(d_1098_i2_):
                                return (d_1098_i2_) - (d_1097_p1_)

                            return lambda58_

                        init35_ = lambda57_(p1)
                        nw200_ = _dafny.Array(None, 27)
                        for i0_35_ in range(nw200_.length(0)):
                            nw200_[i0_35_] = init35_(i0_35_)
                        d_1096_v30_ = nw200_
                        d_1096_v30_ = p3
                        (globalState).f17 = p0
                        d_1099_v31_: _dafny.Map
                        d_1099_v31_ = _dafny.Map({d_1027_v0_: d_1052_v6_})
                        d_1100_v32_: _dafny.Seq
                        d_1100_v32_ = _dafny.SeqWithoutIsStrInference([((d_1099_v31_)[d_1027_v0_] if (d_1027_v0_) in (d_1099_v31_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "buqwaquu"))), d_1052_v6_])
                        d_1101_v34_: _dafny.Seq
                        d_1101_v34_ = _dafny.SeqWithoutIsStrInference([d_1027_v0_, d_1027_v0_, d_1027_v0_])
                        d_1102_v35_: _dafny.Seq
                        d_1102_v35_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_1103_v36_: _dafny.Set
                        d_1103_v36_ = _dafny.Set({d_1102_v35_})
                        d_1104_v37_: _dafny.Array
                        nw201_ = _dafny.Array(None, 25)
                        nw201_[int(0)] = d_1027_v0_
                        def iife59_():
                            coll37_ = _dafny.Set()
                            compr_37_: _dafny.Seq
                            for compr_37_ in (_dafny.MultiSet(d_1100_v32_)).Elements:
                                d_1105_v33_: _dafny.Seq = compr_37_
                                if (d_1105_v33_) in (_dafny.MultiSet(d_1100_v32_)):
                                    coll37_ = coll37_.union(_dafny.Set([d_1105_v33_]))
                            return _dafny.Set(coll37_)
                        nw201_[int(1)] = (iife59_()
                        ).isdisjoint(_dafny.Set({d_1052_v6_}))
                        nw201_[int(2)] = not(True)
                        nw201_[int(3)] = d_1027_v0_
                        nw201_[int(4)] = True
                        nw201_[int(5)] = d_1027_v0_
                        nw201_[int(6)] = (d_1027_v0_) or (True)
                        nw201_[int(7)] = False
                        nw201_[int(8)] = d_1027_v0_
                        nw201_[int(9)] = (d_1101_v34_)[default__.safeIndex(p1, len(d_1101_v34_))]
                        nw201_[int(10)] = d_1027_v0_
                        nw201_[int(11)] = d_1027_v0_
                        nw201_[int(12)] = d_1027_v0_
                        nw201_[int(13)] = d_1027_v0_
                        nw201_[int(14)] = True
                        nw201_[int(15)] = d_1027_v0_
                        nw201_[int(16)] = d_1027_v0_
                        nw201_[int(17)] = d_1027_v0_
                        nw201_[int(18)] = (_dafny.Map({len((d_1052_v6_).set(default__.safeIndex(p2, len(d_1052_v6_)), (d_1031_v4_)[default__.safeIndex(219, (d_1031_v4_).length(0))])): d_1027_v0_})) != (_dafny.Map({p2: d_1027_v0_}))
                        nw201_[int(19)] = not(d_1027_v0_)
                        nw201_[int(20)] = d_1027_v0_
                        nw201_[int(21)] = d_1027_v0_
                        nw201_[int(22)] = d_1027_v0_
                        nw201_[int(23)] = not(d_1027_v0_)
                        nw201_[int(24)] = (d_1102_v35_) not in (d_1103_v36_)
                        d_1104_v37_ = nw201_
                        index162_ = default__.safeIndex(741, (d_1104_v37_).length(0))
                        (d_1104_v37_)[index162_] = d_1027_v0_
                        index163_ = default__.safeIndex(741, (d_1104_v37_).length(0))
                        (d_1104_v37_)[index163_] = (default__.safeModuloInt(len(d_1101_v34_), p2)) != ((p1) - (p0))
                        d_1106_v38_: D0
                        d_1106_v38_ = D0_DC2(D0_DC1(d_1027_v0_, p2))
                        d_1107_v39_: D0
                        d_1107_v39_ = D0_DC2(d_1106_v38_)
                        d_1108_v40_: D0
                        d_1108_v40_ = D0_DC2(d_1106_v38_)
                        d_1109_v41_: D0
                        d_1109_v41_ = D0_DC2(d_1106_v38_)
                        d_1110_v42_: D0
                        d_1110_v42_ = D0_DC2((d_1109_v41_).cf3)
                        d_1111_v43_: D0
                        d_1111_v43_ = D0_DC2(d_1106_v38_)
                        d_1112_v44_: D0
                        d_1112_v44_ = D0_DC2(d_1108_v40_)
                        d_1113_v45_: _dafny.Seq
                        d_1113_v45_ = _dafny.SeqWithoutIsStrInference([d_1107_v39_])
                        d_1114_v46_: D0
                        d_1114_v46_ = D0_DC2((d_1113_v45_)[default__.safeIndex(p1, len(d_1113_v45_))])
                        d_1115_v47_: D0
                        d_1115_v47_ = D0_DC2(d_1112_v44_)
                        d_1116_v48_: C6
                        nw202_ = C6()
                        nw202_.ctor__(d_1115_v47_, (d_1104_v37_)[default__.safeIndex(741, (d_1104_v37_).length(0))], (d_1104_v37_)[default__.safeIndex(741, (d_1104_v37_).length(0))], default__.safeModuloInt((0) - (p0), p1))
                        d_1116_v48_ = nw202_
                    d_1117_v49_: _dafny.Array
                    nw203_ = _dafny.Array(False, 28)
                    d_1117_v49_ = nw203_
                    d_1118_v50_: _dafny.Map
                    d_1118_v50_ = _dafny.Map({d_1027_v0_: d_1117_v49_})
                    d_1119_v51_: _dafny.Seq
                    d_1119_v51_ = _dafny.SeqWithoutIsStrInference([d_1118_v50_, d_1118_v50_])
                    d_1118_v50_ = (((d_1119_v51_)[default__.safeIndex(p0, len(d_1119_v51_))]) | (d_1118_v50_)) | (_dafny.Map({d_1027_v0_: d_1117_v49_}))
                    pass
            pass
        d_1120_v52_: _dafny.Array
        def lambda59_(d_1121_p1_):
            def lambda60_(d_1122_i4_):
                return (d_1122_i4_) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([d_1121_p1_])))])}))])))

            return lambda60_

        init36_ = lambda59_(p1)
        nw204_ = _dafny.Array(None, 19)
        for i0_36_ in range(nw204_.length(0)):
            nw204_[i0_36_] = init36_(i0_36_)
        d_1120_v52_ = nw204_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1120_v52_).length(0)):
            d_1123_i3_: int = guard_loop_5_
            if (True) and (((0) <= (d_1123_i3_)) and ((d_1123_i3_) < ((d_1120_v52_).length(0)))):
                (d_1120_v52_)[(d_1123_i3_)] = default__.safeModuloInt(d_1123_i3_, (p2) - (p2))
        d_1124_v53_: _dafny.Map
        d_1124_v53_ = _dafny.Map({d_1027_v0_: p3})
        d_1125_v54_: _dafny.Seq
        d_1125_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etgsrnnx"))
        d_1124_v53_ = (d_1124_v53_).set((d_1125_v54_) < (d_1125_v54_), d_1120_v52_)
        d_1126_v56_: _dafny.Seq
        d_1126_v56_ = _dafny.SeqWithoutIsStrInference([p2])
        d_1127_v57_: D10
        d_1127_v57_ = D10_DC28(d_1126_v56_)
        def iife60_():
            coll38_ = _dafny.Map()
            compr_38_: int
            for compr_38_ in ((d_1127_v57_).cf52).Elements:
                d_1128_v55_: int = compr_38_
                if (d_1128_v55_) in ((d_1127_v57_).cf52):
                    coll38_[default__.safeModuloInt(d_1128_v55_, p2)] = len(_dafny.Map({p2: d_1027_v0_}))
            return _dafny.Map(coll38_)
        if (_dafny.Map({p0: 778})) != ((iife60_()
        )):
            d_1129_v58_: C2
            nw205_ = C2()
            nw205_.ctor__(p1)
            d_1129_v58_ = nw205_
            (globalState).f17 = ((0) - (p2)) - ((p1) - (518))
            d_1130_v59_: _dafny.MultiSet
            d_1130_v59_ = _dafny.MultiSet([p1])
            if (d_1130_v59_).ispropersubset(d_1130_v59_):
                (globalState).f6 = d_1027_v0_
                (globalState).f14 = -370
                (globalState).f9 = d_1027_v0_
                d_1131_v60_: _dafny.Seq
                d_1131_v60_ = _dafny.SeqWithoutIsStrInference([(d_1130_v59_) | (d_1130_v59_), (d_1130_v59_).intersection(d_1130_v59_)])
                d_1131_v60_ = _dafny.SeqWithoutIsStrInference([d_1130_v59_])
                d_1132_v61_: int
                d_1133_v62_: _dafny.Array
                d_1134_v63_: bool
                out11_: int
                out12_: _dafny.Array
                out13_: bool
                out11_, out12_, out13_ = (d_1129_v58_).m0(d_1027_v0_, d_1027_v0_, globalState)
                d_1132_v61_ = out11_
                d_1133_v62_ = out12_
                d_1134_v63_ = out13_
            elif True:
                d_1135_v64_: _dafny.MultiSet
                d_1135_v64_ = _dafny.MultiSet([d_1027_v0_, (d_1027_v0_) == (d_1027_v0_), not(False), (d_1032_v5_) != (d_1032_v5_)])
                d_1135_v64_ = d_1135_v64_
                (globalState).f15 = d_1027_v0_
                d_1136_v65_: _dafny.Seq
                d_1136_v65_ = _dafny.SeqWithoutIsStrInference([d_1027_v0_])
                d_1137_v66_: D6
                d_1137_v66_ = D6_DC16(d_1136_v65_)
                d_1138_v67_: _dafny.Seq
                d_1138_v67_ = _dafny.SeqWithoutIsStrInference([d_1137_v66_, D6_DC16(d_1136_v65_)])
                d_1139_v68_: _dafny.Map
                d_1139_v68_ = _dafny.Map({d_1138_v67_: d_1027_v0_})
                d_1139_v68_ = default__.fm40(d_1027_v0_, globalState)
                d_1140_v69_: _dafny.Array
                nw206_ = _dafny.Array(_dafny.Seq({}), 3)
                d_1140_v69_ = nw206_
                d_1141_v70_: _dafny.Map
                d_1141_v70_ = _dafny.Map({d_1140_v69_: (d_1027_v0_) or (d_1027_v0_)})
                d_1141_v70_ = (d_1141_v70_).set(d_1140_v69_, d_1027_v0_)
                d_1142_v71_: _dafny.Map
                d_1142_v71_ = _dafny.Map({(p0) + (p0): d_1140_v69_})
                d_1142_v71_ = (d_1142_v71_).set((p2) - (p2), d_1140_v69_)
            (globalState).f8 = p1
            rhs199_ = not(d_1027_v0_)
            rhs200_ = (d_1027_v0_) == ((d_1027_v0_) or (d_1027_v0_))
            rhs201_ = d_1027_v0_
            rhs202_ = _dafny.SeqWithoutIsStrInference([d_1028_v1_ for d_1143_i5_ in range(default__.abs(879))])
            rhs203_ = 264
            lhs174_ = globalState
            lhs175_ = globalState
            lhs176_ = globalState
            lhs177_ = globalState
            lhs174_.f1 = rhs199_
            lhs175_.f6 = rhs200_
            lhs176_.f6 = rhs201_
            d_1125_v54_ = rhs202_
            lhs177_.f5 = rhs203_
        elif True:
            if True:
                d_1125_v54_ = _dafny.SeqWithoutIsStrInference([d_1028_v1_])
                d_1144_v72_: C4
                nw207_ = C4()
                nw207_.ctor__(not(d_1027_v0_), not(True), (p2) > (456))
                d_1144_v72_ = nw207_
                d_1144_v72_ = d_1144_v72_
                d_1125_v54_ = (d_1125_v54_) + (d_1125_v54_)
                d_1145_v73_: _dafny.Map
                d_1145_v73_ = _dafny.Map({(d_1144_v72_).f27: p1})
                d_1145_v73_ = (d_1145_v73_).set(not(not((d_1144_v72_).f27)), p0)
                d_1146_v74_: _dafny.Seq
                d_1146_v74_ = _dafny.SeqWithoutIsStrInference([d_1027_v0_])
                d_1147_v75_: _dafny.Array
                nw208_ = _dafny.Array(None, 8)
                nw208_[int(0)] = d_1126_v56_
                nw208_[int(1)] = (_dafny.SeqWithoutIsStrInference([p2 for d_1148_i6_ in range(default__.abs(428))])) + (_dafny.SeqWithoutIsStrInference([p1 for d_1149_i7_ in range(default__.abs(787))]))
                nw208_[int(2)] = d_1126_v56_
                nw208_[int(3)] = (_dafny.SeqWithoutIsStrInference([len(d_1146_v74_), p1])) + (d_1126_v56_)
                nw208_[int(4)] = (d_1127_v57_).cf52
                nw208_[int(5)] = d_1126_v56_
                nw208_[int(6)] = d_1126_v56_
                nw208_[int(7)] = (d_1126_v56_) + (d_1126_v56_)
                d_1147_v75_ = nw208_
                index164_ = default__.safeIndex(475, (d_1147_v75_).length(0))
                (d_1147_v75_)[index164_] = d_1126_v56_
                d_1150_v76_: _dafny.Set
                d_1150_v76_ = _dafny.Set({len(d_1126_v56_), p2, p1, len(d_1146_v74_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "luqkew")))})
                d_1151_v77_: _dafny.MultiSet
                d_1151_v77_ = _dafny.MultiSet([p1])
                d_1152_v80_: D8
                d_1152_v80_ = D8_DC23(p0, d_1150_v76_, p2, p0, p2)
                d_1153_v81_: D0
                d_1153_v81_ = D0_DC1(d_1027_v0_, 244)
                d_1154_v82_: _dafny.MultiSet
                d_1154_v82_ = _dafny.MultiSet([d_1153_v81_, d_1153_v81_])
                d_1155_v83_: _dafny.Array
                nw209_ = _dafny.Array(False, 23)
                d_1155_v83_ = nw209_
                d_1156_v84_: _dafny.Set
                d_1156_v84_ = _dafny.Set({d_1155_v83_})
                d_1157_v85_: C0
                nw210_ = C0()
                nw210_.ctor__(d_1156_v84_, len(d_1125_v54_), p0)
                d_1157_v85_ = nw210_
                d_1158_v86_: D8
                d_1158_v86_ = D8_DC24(p2, (d_1031_v4_)[default__.safeIndex(219, (d_1031_v4_).length(0))], (d_1154_v82_).set(d_1153_v81_, default__.abs(len(default__.fm32(d_1027_v0_, _dafny.SeqWithoutIsStrInference([len(d_1029_v2_)]), d_1146_v74_, _dafny.Set({D1_DC4((d_1144_v72_).f27, (d_1144_v72_).f27)}), globalState)))), d_1027_v0_, d_1157_v85_)
                d_1159_v87_: _dafny.Map
                d_1159_v87_ = _dafny.Map({d_1157_v85_.f31: d_1157_v85_})
                d_1160_v90_: _dafny.Array
                nw211_ = _dafny.Array(None, 21)
                nw211_[int(0)] = d_1150_v76_
                nw211_[int(1)] = d_1150_v76_
                nw211_[int(2)] = d_1150_v76_
                nw211_[int(3)] = d_1150_v76_
                def iife61_():
                    coll39_ = _dafny.Set()
                    compr_39_: int
                    for compr_39_ in (d_1151_v77_).Elements:
                        d_1161_v78_: int = compr_39_
                        if (d_1161_v78_) in (d_1151_v77_):
                            coll39_ = coll39_.union(_dafny.Set([(d_1161_v78_) + (len(_dafny.SeqWithoutIsStrInference([292])))]))
                    return _dafny.Set(coll39_)
                nw211_[int(4)] = iife61_()
                
                nw211_[int(5)] = (_dafny.Set({-874, len(d_1032_v5_)})) - (d_1150_v76_)
                nw211_[int(6)] = d_1150_v76_
                nw211_[int(7)] = d_1150_v76_
                nw211_[int(8)] = d_1150_v76_
                def iife62_():
                    coll40_ = _dafny.Set()
                    compr_40_: int
                    for compr_40_ in _dafny.IntegerRange(628, 436):
                        d_1162_v79_: int = compr_40_
                        if ((628) <= (d_1162_v79_)) and ((d_1162_v79_) < (436)):
                            coll40_ = coll40_.union(_dafny.Set([(d_1162_v79_) + (p1)]))
                    return _dafny.Set(coll40_)
                nw211_[int(9)] = (d_1150_v76_) | (iife62_()
                )
                nw211_[int(10)] = _dafny.Set({p2, p1, len(d_1032_v5_)})
                nw211_[int(11)] = (d_1152_v80_).cf41
                nw211_[int(12)] = (d_1150_v76_) - (d_1150_v76_)
                nw211_[int(13)] = d_1150_v76_
                nw211_[int(14)] = _dafny.Set({p2, p2, (d_1158_v86_).cf45, len(d_1159_v87_)})
                nw211_[int(15)] = d_1150_v76_
                nw211_[int(16)] = d_1150_v76_
                def iife63_():
                    coll41_ = _dafny.Set()
                    compr_41_: int
                    for compr_41_ in _dafny.IntegerRange(51, 172):
                        d_1163_v88_: int = compr_41_
                        if ((51) <= (d_1163_v88_)) and ((d_1163_v88_) < (172)):
                            coll41_ = coll41_.union(_dafny.Set([default__.safeDivisionInt(d_1163_v88_, len(d_1125_v54_))]))
                    return _dafny.Set(coll41_)
                nw211_[int(17)] = (iife63_()
                ) | (d_1150_v76_)
                nw211_[int(18)] = default__.fm41(globalState)
                nw211_[int(19)] = d_1150_v76_
                def iife64_():
                    coll42_ = _dafny.Set()
                    compr_42_: int
                    for compr_42_ in _dafny.IntegerRange(806, -107):
                        d_1164_v89_: int = compr_42_
                        if ((806) <= (d_1164_v89_)) and ((d_1164_v89_) < (-107)):
                            coll42_ = coll42_.union(_dafny.Set([(d_1164_v89_) - (d_1157_v85_.f31)]))
                    return _dafny.Set(coll42_)
                nw211_[int(20)] = iife64_()
                
                d_1160_v90_ = nw211_
                index165_ = default__.safeIndex(475, (d_1147_v75_).length(0))
                rhs204_ = d_1126_v56_
                rhs205_ = d_1160_v90_
                rhs206_ = (d_1144_v72_).f27
                lhs178_ = d_1147_v75_
                lhs179_ = default__.safeIndex(475, (d_1147_v75_).length(0))
                lhs180_ = globalState
                lhs178_[lhs179_] = rhs204_
                d_1160_v90_ = rhs205_
                lhs180_.f1 = rhs206_
            elif True:
                d_1165_v91_: _dafny.MultiSet
                d_1165_v91_ = _dafny.MultiSet([d_1027_v0_, d_1027_v0_])
                d_1027_v0_ = (False) in ((d_1165_v91_) - (d_1165_v91_))
                d_1166_v92_: C5
                nw212_ = C5()
                nw212_.ctor__(default__.fm29((d_1165_v91_).cardinality, p2, d_1027_v0_, globalState), d_1027_v0_)
                d_1166_v92_ = nw212_
                nw213_ = C5()
                nw213_.ctor__(d_1027_v0_, d_1027_v0_)
                d_1166_v92_ = nw213_
                index166_ = default__.safeIndex(219, (d_1031_v4_).length(0))
                (d_1031_v4_)[index166_] = d_1028_v1_
                rhs207_ = (default__.safeModuloInt(698, p2) if False else len(d_1125_v54_))
                rhs208_ = d_1125_v54_
                lhs181_ = globalState
                lhs181_.f17 = rhs207_
                d_1125_v54_ = rhs208_
                index167_ = default__.safeIndex(219, (d_1031_v4_).length(0))
                (d_1031_v4_)[index167_] = d_1028_v1_
            (globalState).f14 = default__.fm1(p0, True, globalState)
            d_1167_v93_: _dafny.Map
            d_1167_v93_ = _dafny.Map({p1: d_1027_v0_})
            d_1168_v94_: D0
            d_1168_v94_ = D0_DC0(((d_1167_v93_)[p0] if (p0) in (d_1167_v93_) else d_1027_v0_))
            d_1169_v95_: D0
            d_1169_v95_ = D0_DC2(d_1168_v94_)
            source12_ = d_1169_v95_
            if source12_.is_DC1:
                d_1170___mcc_h18_ = source12_.cf1
                d_1171___mcc_h19_ = source12_.cf2
                d_1172_cf2_ = d_1171___mcc_h19_
                d_1173_cf1_ = d_1170___mcc_h18_
                d_1174_v96_: C6
                nw214_ = C6()
                nw214_.ctor__(d_1169_v95_, d_1173_cf1_, d_1173_cf1_, default__.fm1(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohhj"))), d_1173_cf1_, globalState))
                d_1174_v96_ = nw214_
                d_1175_v97_: _dafny.Map
                d_1175_v97_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urgll")): d_1174_v96_})
                d_1175_v97_ = (d_1175_v97_).set(d_1125_v54_, d_1174_v96_)
                d_1176_v98_: C2
                nw215_ = C2()
                nw215_.ctor__((d_1174_v96_).fm5(p0, len(default__.fm31(d_1173_cf1_, d_1027_v0_, globalState)), globalState))
                d_1176_v98_ = nw215_
                index168_ = default__.safeIndex(507, (d_1120_v52_).length(0))
                (d_1120_v52_)[index168_] = d_1172_cf2_
                d_1177_v99_: _dafny.Seq
                d_1177_v99_ = _dafny.SeqWithoutIsStrInference([d_1173_cf1_, d_1173_cf1_])
                d_1178_v100_: _dafny.Map
                d_1178_v100_ = _dafny.Map({d_1177_v99_: p2})
                d_1179_v101_: _dafny.Map
                d_1179_v101_ = _dafny.Map({(d_1174_v96_).fm2(globalState): d_1176_v98_})
                d_1180_v102_: _dafny.Seq
                d_1180_v102_ = _dafny.SeqWithoutIsStrInference([(d_1179_v101_).set(p1, d_1176_v98_), d_1179_v101_, d_1179_v101_])
                d_1181_v103_: _dafny.Map
                d_1181_v103_ = _dafny.Map({len(d_1180_v102_): p1})
                d_1182_v104_: _dafny.MultiSet
                d_1182_v104_ = _dafny.MultiSet([d_1028_v1_, (d_1031_v4_)[default__.safeIndex(219, (d_1031_v4_).length(0))], d_1028_v1_])
                d_1183_v105_: D0
                d_1183_v105_ = D0_DC0((d_1182_v104_).ispropersubset(d_1182_v104_))
                d_1184_v106_: D15
                d_1184_v106_ = D15_DC36(d_1178_v100_)
                index169_ = default__.safeIndex(507, (d_1120_v52_).length(0))
                def iife65_():
                    coll43_ = _dafny.Map()
                    compr_43_: int
                    for compr_43_ in _dafny.IntegerRange(109, 642):
                        d_1185_v107_: int = compr_43_
                        if ((109) <= (d_1185_v107_)) and ((d_1185_v107_) < (642)):
                            coll43_[(d_1185_v107_) * (p1)] = d_1173_cf1_
                    return _dafny.Map(coll43_)
                rhs209_ = d_1172_cf2_
                rhs210_ = (d_1178_v100_) | ((d_1184_v106_).cf64)
                rhs211_ = d_1125_v54_
                rhs212_ = (default__.fm8(len(_dafny.Map({d_1172_cf2_: len((d_1167_v93_).set(len(iife65_()
                ), d_1173_cf1_))})), globalState)) | ((d_1181_v103_) | (d_1181_v103_))
                rhs213_ = d_1183_v105_
                lhs182_ = d_1120_v52_
                lhs183_ = default__.safeIndex(507, (d_1120_v52_).length(0))
                lhs182_[lhs183_] = rhs209_
                d_1178_v100_ = rhs210_
                d_1125_v54_ = rhs211_
                d_1181_v103_ = rhs212_
                d_1183_v105_ = rhs213_
                d_1172_cf2_ = (d_1120_v52_)[default__.safeIndex(507, (d_1120_v52_).length(0))]
            elif source12_.is_DC0:
                d_1186___mcc_h20_ = source12_.cf0
                d_1187_cf0_ = d_1186___mcc_h20_
                index170_ = default__.safeIndex(889, (p3).length(0))
                (p3)[index170_] = p0
                index171_ = default__.safeIndex(889, (p3).length(0))
                (p3)[index171_] = default__.fm1(p0, d_1187_cf0_, globalState)
                (globalState).f9 = False
                d_1188_v108_: C3
                nw216_ = C3()
                nw216_.ctor__(d_1027_v0_, p1, d_1027_v0_, d_1187_cf0_, p1)
                d_1188_v108_ = nw216_
                d_1189_v109_: _dafny.Array
                nw217_ = _dafny.Array(None, 18)
                nw217_[int(0)] = d_1188_v108_
                nw217_[int(1)] = d_1188_v108_
                nw217_[int(2)] = d_1188_v108_
                nw217_[int(3)] = d_1188_v108_
                nw217_[int(4)] = d_1188_v108_
                nw217_[int(5)] = d_1188_v108_
                nw217_[int(6)] = d_1188_v108_
                nw217_[int(7)] = d_1188_v108_
                nw217_[int(8)] = d_1188_v108_
                nw217_[int(9)] = d_1188_v108_
                nw217_[int(10)] = d_1188_v108_
                nw217_[int(11)] = d_1188_v108_
                nw217_[int(12)] = d_1188_v108_
                nw217_[int(13)] = d_1188_v108_
                nw217_[int(14)] = d_1188_v108_
                nw217_[int(15)] = d_1188_v108_
                nw217_[int(16)] = d_1188_v108_
                nw217_[int(17)] = d_1188_v108_
                d_1189_v109_ = nw217_
                d_1190_v110_: D16
                d_1190_v110_ = D16_DC39(d_1189_v109_)
                d_1189_v109_ = (d_1190_v110_).cf69
                d_1191_v111_: T0
                nw218_ = C3()
                nw218_.ctor__(d_1187_cf0_, len(_dafny.Map({p0: d_1187_cf0_})), d_1187_cf0_, (d_1188_v108_).f28, 150)
                d_1191_v111_ = nw218_
                d_1191_v111_ = d_1191_v111_
            elif True:
                d_1192___mcc_h21_ = source12_.cf3
                d_1193_cf3_ = d_1192___mcc_h21_
                d_1194_v112_: _dafny.Array
                nw219_ = _dafny.Array(_dafny.MultiSet({}), 6)
                d_1194_v112_ = nw219_
                d_1195_v113_: _dafny.MultiSet
                d_1195_v113_ = _dafny.MultiSet([d_1027_v0_])
                index172_ = default__.safeIndex(922, (d_1194_v112_).length(0))
                (d_1194_v112_)[index172_] = d_1195_v113_
                index173_ = default__.safeIndex(922, (d_1194_v112_).length(0))
                (d_1194_v112_)[index173_] = d_1195_v113_
                d_1196_v114_: _dafny.Seq
                d_1196_v114_ = _dafny.SeqWithoutIsStrInference([d_1027_v0_])
                d_1197_v115_: _dafny.Set
                d_1197_v115_ = _dafny.Set({(d_1196_v114_)[default__.safeIndex(len((d_1125_v54_).set(default__.safeIndex(p2, len(d_1125_v54_)), (d_1031_v4_)[default__.safeIndex(219, (d_1031_v4_).length(0))])), len(d_1196_v114_))]})
                d_1198_v116_: _dafny.MultiSet
                d_1198_v116_ = _dafny.MultiSet([d_1197_v115_, d_1197_v115_, d_1197_v115_, _dafny.Set({d_1027_v0_, d_1027_v0_, d_1027_v0_, not(d_1027_v0_)})])
                d_1198_v116_ = _dafny.MultiSet([d_1197_v115_])
                d_1199_v117_: _dafny.Map
                d_1199_v117_ = _dafny.Map({d_1196_v114_: d_1196_v114_})
                d_1199_v117_ = (d_1199_v117_).set(d_1196_v114_, d_1196_v114_)
                d_1125_v54_ = ((((d_1125_v54_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')]))).set(default__.safeIndex(p0, len((d_1125_v54_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])))), d_1028_v1_)).set(default__.safeIndex(p1, len(((d_1125_v54_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')]))).set(default__.safeIndex(p0, len((d_1125_v54_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])))), d_1028_v1_))), d_1028_v1_)) + ((d_1125_v54_) + (d_1125_v54_))
            d_1200_v118_: T1
            nw220_ = C5()
            nw220_.ctor__(not(d_1027_v0_), d_1027_v0_)
            d_1200_v118_ = nw220_
            d_1201_v119_: _dafny.Map
            d_1201_v119_ = _dafny.Map({d_1200_v118_: ((d_1167_v93_)[p1] if (p1) in (d_1167_v93_) else (d_1200_v118_).f25)})
            if default__.fm29(len(d_1201_v119_), (0) - ((0) - (p2)), not ((d_1200_v118_).f25) or (d_1200_v118_.f24), globalState):
                d_1202_v120_: _dafny.Map
                d_1202_v120_ = _dafny.Map({p1: p2})
                (globalState).f5 = ((d_1202_v120_)[p2] if (p2) in (d_1202_v120_) else (0) - (p1))
                d_1203_v121_: D3
                d_1203_v121_ = D3_DC8(d_1125_v54_)
                d_1204_v122_: D8
                d_1204_v122_ = D8_DC22(p3)
                d_1205_v123_: _dafny.MultiSet
                d_1205_v123_ = _dafny.MultiSet([d_1120_v52_, (d_1204_v122_).cf39])
                d_1206_v124_: _dafny.Map
                d_1206_v124_ = _dafny.Map({(default__.fm25(d_1027_v0_, (d_1200_v118_).f25, _dafny.CodePoint('j'), globalState)) + ((d_1203_v121_).cf14): d_1205_v123_})
                d_1206_v124_ = (d_1206_v124_).set(d_1125_v54_, (d_1205_v123_).set(d_1120_v52_, default__.abs(p0)))
                (globalState).f5 = default__.safeDivisionInt(len(d_1125_v54_), len(d_1202_v120_))
                d_1028_v1_ = (d_1031_v4_)[default__.safeIndex(219, (d_1031_v4_).length(0))]
                d_1125_v54_ = ((d_1125_v54_) + (_dafny.SeqWithoutIsStrInference([d_1028_v1_ for d_1207_i8_ in range(default__.abs(151))]))).set(default__.safeIndex(p0, len((d_1125_v54_) + (_dafny.SeqWithoutIsStrInference([d_1028_v1_ for d_1208_i8_ in range(default__.abs(151))])))), (d_1031_v4_)[default__.safeIndex(219, (d_1031_v4_).length(0))])
            elif True:
                d_1209_v125_: _dafny.Map
                d_1209_v125_ = _dafny.Map({(-822) + (p0): d_1120_v52_})
                d_1210_v126_: D15
                d_1210_v126_ = D15_DC37((d_1200_v118_).f25, _dafny.Map({d_1027_v0_: (d_1200_v118_).f25}), p2)
                d_1211_v127_: _dafny.Map
                d_1211_v127_ = _dafny.Map({(d_1210_v126_).cf65: p0})
                d_1212_v128_: _dafny.MultiSet
                d_1212_v128_ = _dafny.MultiSet([d_1211_v127_])
                d_1209_v125_ = (d_1209_v125_).set(default__.safeDivisionInt((d_1212_v128_).cardinality, p2), p3)
                index174_ = default__.safeIndex(884, (d_1120_v52_).length(0))
                (d_1120_v52_)[index174_] = len(d_1125_v54_)
                d_1213_v129_: C6
                nw221_ = C6()
                nw221_.ctor__(d_1169_v95_, (d_1200_v118_).f25, (d_1200_v118_).f25, p0)
                d_1213_v129_ = nw221_
                d_1214_v130_: _dafny.Map
                d_1214_v130_ = _dafny.Map({d_1213_v129_: (p2) < (151)})
                d_1215_v131_: _dafny.Array
                def lambda61_(d_1216_v118_):
                    def lambda62_(d_1217_i9_):
                        return d_1216_v118_.f24

                    return lambda62_

                init37_ = lambda61_(d_1200_v118_)
                nw222_ = _dafny.Array(None, 24)
                for i0_37_ in range(nw222_.length(0)):
                    nw222_[i0_37_] = init37_(i0_37_)
                d_1215_v131_ = nw222_
                index175_ = default__.safeIndex(884, (d_1120_v52_).length(0))
                rhs214_ = ((d_1214_v130_)[d_1213_v129_] if (d_1213_v129_) in (d_1214_v130_) else (d_1200_v118_).f25)
                rhs215_ = d_1027_v0_
                rhs216_ = (p0) * ((d_1126_v56_)[default__.safeIndex(p0, len(d_1126_v56_))])
                rhs217_ = d_1215_v131_
                lhs184_ = globalState
                lhs185_ = d_1120_v52_
                lhs186_ = default__.safeIndex(884, (d_1120_v52_).length(0))
                d_1027_v0_ = rhs214_
                lhs184_.f6 = rhs215_
                lhs185_[lhs186_] = rhs216_
                r1 = rhs217_
                (globalState).f15 = d_1027_v0_
                d_1218_v132_: _dafny.Array
                d_1218_v132_ = d_1215_v131_
                d_1219_v133_: _dafny.Map
                d_1219_v133_ = _dafny.Map({d_1027_v0_: (d_1218_v132_)})
                d_1220_v134_: _dafny.Array
                nw223_ = _dafny.Array(None, 28)
                nw223_[int(0)] = d_1215_v131_
                nw223_[int(1)] = d_1215_v131_
                nw223_[int(2)] = d_1215_v131_
                nw223_[int(3)] = d_1215_v131_
                nw223_[int(4)] = d_1215_v131_
                nw223_[int(5)] = (d_1218_v132_)
                nw223_[int(6)] = d_1215_v131_
                nw223_[int(7)] = d_1215_v131_
                nw223_[int(8)] = d_1215_v131_
                nw223_[int(9)] = d_1215_v131_
                nw223_[int(10)] = d_1215_v131_
                nw223_[int(11)] = d_1215_v131_
                nw223_[int(12)] = d_1215_v131_
                nw223_[int(13)] = d_1215_v131_
                nw223_[int(14)] = d_1215_v131_
                nw223_[int(15)] = ((d_1219_v133_)[d_1200_v118_.f24] if (d_1200_v118_.f24) in (d_1219_v133_) else d_1215_v131_)
                nw223_[int(16)] = d_1215_v131_
                nw223_[int(17)] = d_1215_v131_
                nw223_[int(18)] = d_1215_v131_
                nw223_[int(19)] = d_1215_v131_
                nw223_[int(20)] = d_1215_v131_
                nw223_[int(21)] = d_1215_v131_
                nw223_[int(22)] = d_1215_v131_
                nw223_[int(23)] = d_1215_v131_
                nw223_[int(24)] = d_1215_v131_
                nw223_[int(25)] = d_1215_v131_
                nw223_[int(26)] = d_1215_v131_
                nw223_[int(27)] = d_1215_v131_
                d_1220_v134_ = nw223_
                index176_ = default__.safeIndex(95, (d_1220_v134_).length(0))
                (d_1220_v134_)[index176_] = ((d_1219_v133_)[(d_1200_v118_).f25] if ((d_1200_v118_).f25) in (d_1219_v133_) else d_1215_v131_)
                index177_ = default__.safeIndex(95, (d_1220_v134_).length(0))
                rhs218_ = d_1215_v131_
                rhs219_ = (d_1120_v52_)[default__.safeIndex(884, (d_1120_v52_).length(0))]
                rhs220_ = not(not(not(default__.fm29((d_1120_v52_)[default__.safeIndex(884, (d_1120_v52_).length(0))], p2, (d_1200_v118_).f25, globalState))))
                lhs187_ = d_1220_v134_
                lhs188_ = default__.safeIndex(95, (d_1220_v134_).length(0))
                lhs189_ = globalState
                lhs190_ = globalState
                lhs187_[lhs188_] = rhs218_
                lhs189_.f17 = rhs219_
                lhs190_.f15 = rhs220_
                d_1221_v135_: _dafny.Array
                def lambda63_(d_1222_v93_):
                    def lambda64_(d_1223_i10_):
                        return (d_1222_v93_) | (d_1222_v93_)

                    return lambda64_

                init38_ = lambda63_(d_1167_v93_)
                nw224_ = _dafny.Array(None, 8)
                for i0_38_ in range(nw224_.length(0)):
                    nw224_[i0_38_] = init38_(i0_38_)
                d_1221_v135_ = nw224_
                d_1224_v136_: _dafny.MultiSet
                d_1224_v136_ = _dafny.MultiSet([p2])
                index178_ = default__.safeIndex(884, (d_1120_v52_).length(0))
                rhs221_ = 481
                rhs222_ = d_1221_v135_
                rhs223_ = (d_1200_v118_).f25
                rhs224_ = ((d_1120_v52_)[default__.safeIndex(884, (d_1120_v52_).length(0))]) * (((d_1224_v136_).intersection(d_1224_v136_)).cardinality)
                lhs191_ = globalState
                lhs192_ = globalState
                lhs193_ = d_1120_v52_
                lhs194_ = default__.safeIndex(884, (d_1120_v52_).length(0))
                lhs191_.f5 = rhs221_
                d_1221_v135_ = rhs222_
                lhs192_.f1 = rhs223_
                lhs193_[lhs194_] = rhs224_
            d_1225_v137_: _dafny.Map
            d_1225_v137_ = _dafny.Map({d_1125_v54_: p0})
            d_1225_v137_ = (d_1225_v137_).set(d_1125_v54_, p2)
        d_1226_v138_: T1
        nw225_ = C3()
        nw225_.ctor__(d_1027_v0_, len(d_1126_v56_), d_1027_v0_, d_1027_v0_, p2)
        d_1226_v138_ = nw225_
        d_1227_v139_: _dafny.Map
        d_1227_v139_ = _dafny.Map({len(d_1125_v54_): d_1226_v138_})
        d_1228_v140_: _dafny.Map
        d_1228_v140_ = _dafny.Map({p0: d_1227_v139_})
        d_1229_v141_: _dafny.Seq
        d_1229_v141_ = _dafny.SeqWithoutIsStrInference([default__.fm9(d_1226_v138_.f24, globalState), default__.fm29(p1, p2, d_1027_v0_, globalState)])
        r0 = ((d_1228_v140_)[len(d_1229_v141_)] if (len(d_1229_v141_)) in (d_1228_v140_) else (d_1227_v139_) | (d_1227_v139_))
        d_1230_v142_: _dafny.Array
        nw226_ = _dafny.Array(False, 26)
        d_1230_v142_ = nw226_
        r1 = d_1230_v142_
        return r0, r1


class C8(T0, T1):
    def  __init__(self):
        self._f24: bool = False
        self._f23: int = int(0)
        self._f25: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f24(self):
        return self._f24
    @f24.setter
    def f24(self, value):
        self._f24 = value
    @property
    def f23(self):
        return self._f23
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f23, f24, f25):
        (self)._f23 = f23
        (self).f24 = f24
        (self)._f25 = f25

    def fm2(self, globalState):
        return (self).f23

    def fm3(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))) <= ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "motvhqyi"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcsjvtutw"))))

    def fm4(self, p0, p1, p2, globalState):
        return ((_dafny.Map({self.f24: _dafny.Map({(self).f25: self.f24})})) | (_dafny.Map({self.f24: _dafny.Map({True: (self).f25})}))) | (_dafny.Map({(self).f25: _dafny.Map({self.f24: (self).f25})}))

    def fm5(self, p0, p1, globalState):
        return (self).f23

    def fm6(self, p0, p1, p2, p3, globalState):
        return not((self).f25)

    def m0(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: bool = False
        (self).f24 = False
        d_1231_i0_: int
        d_1231_i0_ = 0
        with _dafny.label("9"):
            while not(not (not ((self).f25) or (p1)) or (self.f24)):
                with _dafny.c_label("9"):
                    if (d_1231_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_1231_i0_ = (d_1231_i0_) + (1)
                    d_1232_v0_: _dafny.Seq
                    d_1232_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "luxyg"))
                    d_1233_v1_: _dafny.Seq
                    d_1233_v1_ = _dafny.SeqWithoutIsStrInference([len(d_1232_v0_)])
                    d_1234_v2_: D0
                    d_1234_v2_ = D0_DC0(p0)
                    d_1235_v3_: _dafny.Array
                    nw227_ = _dafny.Array(None, 22)
                    nw227_[int(0)] = (self).fm6(len(_dafny.SeqWithoutIsStrInference([(self).f23])), _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p1, True])) for d_1236_i2_ in range(default__.abs(44))]), d_1233_v1_]), (self).f25, (self).f23, globalState)
                    nw227_[int(1)] = (self).f25
                    nw227_[int(2)] = (self).f25
                    nw227_[int(3)] = True
                    nw227_[int(4)] = p1
                    nw227_[int(5)] = (self).f25
                    nw227_[int(6)] = True
                    nw227_[int(7)] = False
                    nw227_[int(8)] = (self).f25
                    nw227_[int(9)] = p1
                    nw227_[int(10)] = (d_1234_v2_).cf0
                    nw227_[int(11)] = False
                    nw227_[int(12)] = self.f24
                    nw227_[int(13)] = (self).f25
                    nw227_[int(14)] = (d_1234_v2_).cf0
                    nw227_[int(15)] = True
                    nw227_[int(16)] = p0
                    nw227_[int(17)] = p0
                    nw227_[int(18)] = p1
                    nw227_[int(19)] = p0
                    nw227_[int(20)] = (self).f25
                    nw227_[int(21)] = True
                    d_1235_v3_ = nw227_
                    d_1237_v4_: _dafny.Map
                    d_1237_v4_ = _dafny.Map({d_1235_v3_: (self).f23})
                    (globalState).f8 = ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1238_i1_ in range(default__.abs(600))])))) - (len(d_1237_v4_))
                    d_1239_v5_: C6
                    nw228_ = C6()
                    nw228_.ctor__(default__.fm24((self).f23, (self).f25, (self).f25, globalState), self.f24, (d_1232_v0_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))), default__.fm1((self).f23, True, globalState))
                    d_1239_v5_ = nw228_
                    d_1240_v6_: _dafny.Map
                    d_1240_v6_ = _dafny.Map({(self).f23: self.f24})
                    d_1240_v6_ = (d_1240_v6_).set((self).f23, p0)
                    d_1241_v7_: C7
                    nw229_ = C7()
                    nw229_.ctor__()
                    d_1241_v7_ = nw229_
                    pass
            pass
        (globalState).f14 = 846
        d_1242_v8_: C5
        nw230_ = C5()
        nw230_.ctor__(False, (self).f25)
        d_1242_v8_ = nw230_
        d_1243_v9_: _dafny.Seq
        d_1243_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwvnhkl"))
        d_1244_v10_: _dafny.Set
        d_1244_v10_ = _dafny.Set({len(d_1243_v9_), (self).f23})
        d_1245_v11_: _dafny.Set
        d_1245_v11_ = _dafny.Set({d_1244_v10_, d_1244_v10_, _dafny.Set({(self).f23})})
        (globalState).f5 = len(d_1245_v11_)
        d_1246_v12_: D0
        d_1246_v12_ = D0_DC0(self.f24)
        d_1247_v13_: D0
        d_1247_v13_ = D0_DC2(d_1246_v12_)
        d_1247_v13_ = d_1247_v13_
        r0 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iasboupi")))
        d_1248_v14_: _dafny.Array
        nw231_ = _dafny.Array(int(0), 14)
        d_1248_v14_ = nw231_
        r1 = (d_1248_v14_ if (self).f25 else d_1248_v14_)
        r2 = p1
        return r0, r1, r2

    def m1(self, globalState):
        r0: bool = False
        r1: T0 = None
        r2: bool = False
        r3: int = int(0)
        d_1249_v0_: _dafny.Array
        nw232_ = _dafny.Array(False, 15)
        d_1249_v0_ = nw232_
        d_1250_v1_: _dafny.Set
        d_1250_v1_ = _dafny.Set({d_1249_v0_})
        d_1251_v2_: _dafny.Map
        d_1251_v2_ = _dafny.Map({122: -910})
        d_1252_v3_: C0
        nw233_ = C0()
        nw233_.ctor__(d_1250_v1_, ((self).f23) + (len(d_1251_v2_)), (self).f23)
        d_1252_v3_ = nw233_
        d_1253_i0_: int
        d_1253_i0_ = 0
        with _dafny.label("10"):
            while not(self.f24):
                with _dafny.c_label("10"):
                    if (d_1253_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_1253_i0_ = (d_1253_i0_) + (1)
                    d_1254_v4_: _dafny.Array
                    nw234_ = _dafny.Array(int(0), 11)
                    d_1254_v4_ = nw234_
                    d_1254_v4_ = d_1254_v4_
                    d_1255_v5_: C7
                    nw235_ = C7()
                    nw235_.ctor__()
                    d_1255_v5_ = nw235_
                    d_1256_v6_: _dafny.Seq
                    d_1256_v6_ = _dafny.SeqWithoutIsStrInference([d_1252_v3_.f31])
                    d_1257_v7_: _dafny.Map
                    d_1257_v7_ = _dafny.Map({d_1255_v5_: d_1256_v6_})
                    d_1258_v8_: _dafny.MultiSet
                    d_1258_v8_ = _dafny.MultiSet([(self).f25])
                    (globalState).f14 = (len((((d_1257_v7_)[d_1255_v5_] if (d_1255_v5_) in (d_1257_v7_) else d_1256_v6_)) + (d_1256_v6_)) if (self).f25 else (d_1258_v8_).cardinality)
                    d_1259_v9_: _dafny.Seq
                    d_1259_v9_ = _dafny.SeqWithoutIsStrInference([(self).f25])
                    d_1260_v10_: C5
                    nw236_ = C5()
                    nw236_.ctor__(not((self).f25), (d_1259_v9_)[default__.safeIndex(d_1252_v3_.f31, len(d_1259_v9_))])
                    d_1260_v10_ = nw236_
                    d_1260_v10_ = d_1260_v10_
                    (globalState).f8 = d_1252_v3_.f31
                    pass
            pass
        d_1261_v11_: _dafny.MultiSet
        d_1261_v11_ = _dafny.MultiSet([(self).f23])
        d_1262_v12_: _dafny.Seq
        d_1262_v12_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f25, self.f24])
        d_1263_v13_: _dafny.Seq
        d_1263_v13_ = _dafny.SeqWithoutIsStrInference([True, (d_1261_v11_) == (d_1261_v11_), (d_1262_v12_)[default__.safeIndex(d_1252_v3_.f31, len(d_1262_v12_))], not(default__.fm29(d_1252_v3_.f31, ((d_1261_v11_)[(self).f23] if ((self).f23) in (d_1261_v11_) else (self).f23), (self).f25, globalState))])
        rhs225_ = (self).f23
        rhs226_ = (d_1263_v13_)[default__.safeIndex(default__.safeModuloInt((self).f23, d_1252_v3_.f31), len(d_1263_v13_))]
        lhs195_ = globalState
        lhs196_ = self
        lhs195_.f5 = rhs225_
        lhs196_.f24 = rhs226_
        d_1264_v14_: _dafny.Array
        nw237_ = _dafny.Array(int(0), 15)
        d_1264_v14_ = nw237_
        d_1265_v15_: D8
        d_1265_v15_ = D8_DC22(d_1264_v14_)
        source13_ = d_1265_v15_
        if source13_.is_DC23:
            d_1266___mcc_h0_ = source13_.cf40
            d_1267___mcc_h1_ = source13_.cf41
            d_1268___mcc_h2_ = source13_.cf42
            d_1269___mcc_h3_ = source13_.cf43
            d_1270___mcc_h4_ = source13_.cf44
            d_1271_cf44_ = d_1270___mcc_h4_
            d_1272_cf43_ = d_1269___mcc_h3_
            d_1273_cf42_ = d_1268___mcc_h2_
            d_1274_cf41_ = d_1267___mcc_h1_
            d_1275_cf40_ = d_1266___mcc_h0_
            d_1276_v16_: _dafny.Set
            d_1276_v16_ = (d_1252_v3_).f30
            d_1277_v17_: _dafny.Seq
            d_1277_v17_ = _dafny.SeqWithoutIsStrInference([d_1276_v16_])
            (globalState).f5 = (904) - (len(d_1277_v17_))
            (globalState).f1 = self.f24
            if (self.f24) or (self.f24):
                d_1262_v12_ = d_1262_v12_
                d_1278_v19_: _dafny.Seq
                d_1278_v19_ = _dafny.SeqWithoutIsStrInference([d_1273_cf42_])
                def iife66_():
                    coll44_ = _dafny.Map()
                    compr_44_: int
                    for compr_44_ in (d_1278_v19_).Elements:
                        d_1279_v18_: int = compr_44_
                        if (d_1279_v18_) in (d_1278_v19_):
                            coll44_[default__.safeDivisionInt(d_1279_v18_, d_1271_cf44_)] = (self).f25
                    return _dafny.Map(coll44_)
                (globalState).f1 = (D10_DC29((self).f25, (self).f25, len(iife66_()
))).cf54
                d_1280_v20_: _dafny.Seq
                d_1280_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrebo"))
                d_1281_v21_: C4
                nw238_ = C4()
                nw238_.ctor__(not(self.f24), not(self.f24), (self).f25)
                d_1281_v21_ = nw238_
                d_1282_v22_: D17
                d_1282_v22_ = D17_DC41(d_1281_v21_)
                d_1283_v23_: _dafny.MultiSet
                d_1283_v23_ = _dafny.MultiSet([self.f24, self.f24])
                rhs227_ = d_1274_cf41_
                rhs228_ = d_1280_v20_
                rhs229_ = (d_1281_v21_ if (self).f25 else (d_1282_v22_).cf73)
                rhs230_ = (d_1283_v23_).issubset(d_1283_v23_)
                d_1274_cf41_ = rhs227_
                d_1280_v20_ = rhs228_
                d_1281_v21_ = rhs229_
                r0 = rhs230_
                d_1284_v24_: _dafny.Map
                d_1284_v24_ = _dafny.Map({d_1261_v11_: not((self).f25)})
                d_1285_v25_: _dafny.Seq
                d_1285_v25_ = _dafny.SeqWithoutIsStrInference([d_1278_v19_])
                (globalState).f15 = ((d_1281_v21_).f27 if (self.f24 if (d_1281_v21_).f27 else (self).f25) else ((d_1284_v24_)[((_dafny.MultiSet((d_1285_v25_)[default__.safeIndex(d_1271_cf44_, len(d_1285_v25_))])).set(d_1273_cf42_, default__.abs((self).f23))).set(d_1275_cf40_, default__.abs(d_1252_v3_.f31))] if (((_dafny.MultiSet((d_1285_v25_)[default__.safeIndex(d_1271_cf44_, len(d_1285_v25_))])).set(d_1273_cf42_, default__.abs((self).f23))).set(d_1275_cf40_, default__.abs(d_1252_v3_.f31))) in (d_1284_v24_) else self.f24))
                d_1286_v26_: _dafny.Array
                nw239_ = _dafny.Array(None, 25)
                d_1286_v26_ = nw239_
                d_1287_v27_: T1
                nw240_ = C4()
                nw240_.ctor__((d_1281_v21_).f27, (self).f25, (d_1281_v21_).f27)
                d_1287_v27_ = nw240_
                d_1288_v28_: _dafny.Seq
                d_1288_v28_ = _dafny.SeqWithoutIsStrInference([d_1287_v27_])
                nw241_ = _dafny.Array(None, 21)
                nw241_[int(0)] = d_1287_v27_
                nw241_[int(1)] = d_1287_v27_
                nw241_[int(2)] = d_1287_v27_
                nw241_[int(3)] = d_1287_v27_
                nw241_[int(4)] = (d_1287_v27_ if (self).fm3(d_1280_v20_, (d_1287_v27_).f25, d_1252_v3_.f31, d_1278_v19_, globalState) else d_1287_v27_)
                nw241_[int(5)] = d_1287_v27_
                nw241_[int(6)] = d_1287_v27_
                nw241_[int(7)] = d_1287_v27_
                nw241_[int(8)] = d_1287_v27_
                nw241_[int(9)] = d_1287_v27_
                nw241_[int(10)] = d_1287_v27_
                nw241_[int(11)] = d_1287_v27_
                nw241_[int(12)] = d_1287_v27_
                nw241_[int(13)] = d_1287_v27_
                nw241_[int(14)] = d_1287_v27_
                nw241_[int(15)] = d_1287_v27_
                nw241_[int(16)] = d_1287_v27_
                nw241_[int(17)] = d_1287_v27_
                nw241_[int(18)] = d_1287_v27_
                nw241_[int(19)] = (d_1288_v28_)[default__.safeIndex(len(d_1274_cf41_), len(d_1288_v28_))]
                nw241_[int(20)] = d_1287_v27_
                d_1286_v26_ = nw241_
            elif True:
                d_1289_v29_: _dafny.Seq
                d_1289_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gigjdhnxs"))
                d_1290_v30_: str
                d_1290_v30_ = _dafny.CodePoint('p')
                d_1289_v29_ = (d_1289_v29_) + (((d_1289_v29_) + (d_1289_v29_)).set(default__.safeIndex(len(d_1262_v12_), len((d_1289_v29_) + (d_1289_v29_))), d_1290_v30_))
                def iife67_():
                    coll45_ = _dafny.Set()
                    compr_45_: int
                    for compr_45_ in _dafny.IntegerRange(-720, 674):
                        d_1291_v31_: int = compr_45_
                        if ((-720) <= (d_1291_v31_)) and ((d_1291_v31_) < (674)):
                            coll45_ = coll45_.union(_dafny.Set([(d_1291_v31_) + (d_1252_v3_.f31)]))
                    return _dafny.Set(coll45_)
                d_1274_cf41_ = (iife67_()
                ) - ((d_1274_cf41_).intersection(d_1274_cf41_))
                index179_ = default__.safeIndex(675, (d_1249_v0_).length(0))
                (d_1249_v0_)[index179_] = (self).f25
                index180_ = default__.safeIndex(413, (d_1249_v0_).length(0))
                (d_1249_v0_)[index180_] = not((self).f25)
                d_1292_v32_: _dafny.Map
                d_1292_v32_ = _dafny.Map({d_1252_v3_.f31: (self).f25})
                d_1293_v33_: _dafny.Seq
                d_1293_v33_ = _dafny.SeqWithoutIsStrInference([d_1252_v3_.f31, 702, d_1275_cf40_, d_1275_cf40_])
                index181_ = default__.safeIndex(675, (d_1249_v0_).length(0))
                index182_ = default__.safeIndex(413, (d_1249_v0_).length(0))
                rhs231_ = self.f24
                rhs232_ = not(True)
                rhs233_ = (d_1274_cf41_).issubset(d_1274_cf41_)
                rhs234_ = (self).fm3(d_1289_v29_, not(self.f24), len(d_1292_v32_), d_1293_v33_, globalState)
                lhs197_ = d_1249_v0_
                lhs198_ = default__.safeIndex(675, (d_1249_v0_).length(0))
                lhs199_ = d_1249_v0_
                lhs200_ = default__.safeIndex(413, (d_1249_v0_).length(0))
                lhs201_ = globalState
                lhs202_ = globalState
                lhs197_[lhs198_] = rhs231_
                lhs199_[lhs200_] = rhs232_
                lhs201_.f15 = rhs233_
                lhs202_.f15 = rhs234_
                d_1294_v34_: _dafny.Array
                def lambda65_(d_1295_v29_):
                    def lambda66_(d_1296_i1_):
                        return d_1295_v29_

                    return lambda66_

                init39_ = lambda65_(d_1289_v29_)
                nw242_ = _dafny.Array(None, 15)
                for i0_39_ in range(nw242_.length(0)):
                    nw242_[i0_39_] = init39_(i0_39_)
                d_1294_v34_ = nw242_
                index183_ = default__.safeIndex(899, (d_1294_v34_).length(0))
                (d_1294_v34_)[index183_] = (d_1289_v29_) + (d_1289_v29_)
                index184_ = default__.safeIndex(899, (d_1294_v34_).length(0))
                (d_1294_v34_)[index184_] = d_1289_v29_
                d_1297_v35_: _dafny.Map
                d_1297_v35_ = _dafny.Map({d_1264_v14_: (d_1294_v34_)[default__.safeIndex(899, (d_1294_v34_).length(0))]})
                (globalState).f9 = (d_1264_v14_) in (d_1297_v35_)
            d_1298_v36_: _dafny.Array
            def lambda67_(d_1299_i2_):
                return _dafny.CodePoint('w')

            init40_ = lambda67_
            nw243_ = _dafny.Array(None, 11)
            for i0_40_ in range(nw243_.length(0)):
                nw243_[i0_40_] = init40_(i0_40_)
            d_1298_v36_ = nw243_
            index185_ = default__.safeIndex(196, (d_1298_v36_).length(0))
            (d_1298_v36_)[index185_] = _dafny.CodePoint('l')
            d_1300_v37_: str
            d_1300_v37_ = _dafny.CodePoint('m')
            index186_ = default__.safeIndex(196, (d_1298_v36_).length(0))
            (d_1298_v36_)[index186_] = (d_1300_v37_ if False else d_1300_v37_)
        elif source13_.is_DC24:
            d_1301___mcc_h5_ = source13_.cf45
            d_1302___mcc_h6_ = source13_.cf46
            d_1303___mcc_h7_ = source13_.cf47
            d_1304___mcc_h8_ = source13_.cf48
            d_1305___mcc_h9_ = source13_.cf49
            d_1306_cf49_ = d_1305___mcc_h9_
            d_1307_cf48_ = d_1304___mcc_h8_
            d_1308_cf47_ = d_1303___mcc_h7_
            d_1309_cf46_ = d_1302___mcc_h6_
            d_1310_cf45_ = d_1301___mcc_h5_
            d_1311_v38_: _dafny.Seq
            d_1311_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            d_1311_v38_ = d_1311_v38_
            d_1312_v39_: _dafny.Array
            def lambda68_(d_1313_i3_):
                return D17_DC42()

            init41_ = lambda68_
            nw244_ = _dafny.Array(None, 26)
            for i0_41_ in range(nw244_.length(0)):
                nw244_[i0_41_] = init41_(i0_41_)
            d_1312_v39_ = nw244_
            d_1312_v39_ = d_1312_v39_
            d_1314_v40_: _dafny.Seq
            d_1314_v40_ = _dafny.SeqWithoutIsStrInference([(self).f23])
            d_1315_v41_: D11
            d_1315_v41_ = D11_DC31(d_1264_v14_, d_1306_cf49_.f31, (self).f25)
            rhs235_ = d_1314_v40_
            rhs236_ = d_1262_v12_
            rhs237_ = default__.fm25(not(d_1307_cf48_), (d_1306_cf49_.f31) > ((0) - (d_1310_cf45_)), d_1309_cf46_, globalState)
            rhs238_ = (default__.fm31((self).f25, (d_1315_v41_).cf59, globalState)) <= (((d_1314_v40_).set(default__.safeIndex(d_1310_cf45_, len(d_1314_v40_)), d_1252_v3_.f31) if d_1307_cf48_ else d_1314_v40_))
            lhs203_ = globalState
            d_1314_v40_ = rhs235_
            d_1262_v12_ = rhs236_
            d_1311_v38_ = rhs237_
            lhs203_.f1 = rhs238_
            (globalState).f8 = d_1306_cf49_.f31
        elif True:
            d_1316___mcc_h10_ = source13_.cf39
            d_1317_cf39_ = d_1316___mcc_h10_
            if self.f24:
                d_1318_v42_: _dafny.Map
                d_1318_v42_ = _dafny.Map({self.f24: d_1261_v11_})
                d_1319_v43_: _dafny.Set
                d_1319_v43_ = _dafny.Set({(self).f25})
                d_1320_v44_: _dafny.Seq
                d_1320_v44_ = _dafny.SeqWithoutIsStrInference([d_1252_v3_.f31])
                d_1321_v45_: D1
                d_1321_v45_ = D1_DC4(self.f24, self.f24)
                d_1322_v46_: _dafny.Set
                d_1322_v46_ = _dafny.Set({d_1321_v45_, d_1321_v45_, d_1321_v45_})
                d_1323_v47_: _dafny.Map
                d_1323_v47_ = _dafny.Map({(d_1319_v43_) | (default__.fm32(self.f24, d_1320_v44_, d_1262_v12_, d_1322_v46_, globalState)): d_1261_v11_})
                d_1324_v48_: _dafny.Seq
                d_1324_v48_ = _dafny.SeqWithoutIsStrInference([d_1319_v43_])
                d_1325_v49_: _dafny.Map
                d_1325_v49_ = _dafny.Map({self.f24: -195})
                rhs239_ = ((d_1323_v47_)[(d_1324_v48_)[default__.safeIndex(d_1252_v3_.f31, len(d_1324_v48_))]] if ((d_1324_v48_)[default__.safeIndex(d_1252_v3_.f31, len(d_1324_v48_))]) in (d_1323_v47_) else _dafny.MultiSet([(0) - (len(d_1325_v49_))]))
                rhs240_ = (d_1320_v44_)[default__.safeIndex(d_1252_v3_.f31, len(d_1320_v44_))]
                rhs241_ = (d_1252_v3_.f31) - (d_1252_v3_.f31)
                rhs242_ = d_1318_v42_
                lhs204_ = globalState
                d_1261_v11_ = rhs239_
                lhs204_.f5 = rhs240_
                r3 = rhs241_
                d_1318_v42_ = rhs242_
                d_1326_v50_: str
                d_1326_v50_ = _dafny.CodePoint('v')
                d_1327_v51_: D0
                d_1327_v51_ = D0_DC1(self.f24, d_1252_v3_.f31)
                d_1328_v52_: _dafny.Map
                d_1328_v52_ = _dafny.Map({d_1326_v50_: not((d_1327_v51_).cf1)})
                d_1329_v53_: D16
                d_1329_v53_ = D16_DC40((self).f23, d_1317_cf39_, d_1328_v52_)
                d_1330_v54_: D0
                d_1330_v54_ = D0_DC2(D0_DC0((self).f25))
                d_1331_v55_: D0
                d_1331_v55_ = D0_DC2(d_1330_v54_)
                d_1332_v56_: C6
                nw245_ = C6()
                nw245_.ctor__(d_1331_v55_, True, self.f24, d_1252_v3_.f31)
                d_1332_v56_ = nw245_
                d_1333_v57_: _dafny.MultiSet
                d_1333_v57_ = _dafny.MultiSet([d_1332_v56_, d_1332_v56_, d_1332_v56_, d_1332_v56_])
                d_1334_v58_: _dafny.Map
                d_1334_v58_ = _dafny.Map({(d_1329_v53_).cf71: len(_dafny.SeqWithoutIsStrInference([d_1333_v57_, d_1333_v57_, d_1333_v57_, d_1333_v57_]))})
                d_1335_v59_: _dafny.Seq
                d_1335_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqj"))
                (globalState).f5 = default__.safeModuloInt(len((d_1251_v2_).set(d_1252_v3_.f31, (self).f23)), ((d_1334_v58_)[d_1317_cf39_] if (d_1317_cf39_) in (d_1334_v58_) else len(d_1335_v59_)))
                d_1336_v60_: C1
                nw246_ = C1()
                nw246_.ctor__(self.f24, self.f24)
                d_1336_v60_ = nw246_
                d_1337_v61_: _dafny.Seq
                d_1337_v61_ = _dafny.SeqWithoutIsStrInference([d_1249_v0_])
                d_1338_v62_: _dafny.Map
                d_1338_v62_ = _dafny.Map({len(d_1337_v61_): (self).f25})
                d_1339_v63_: _dafny.Seq
                d_1339_v63_ = _dafny.SeqWithoutIsStrInference([d_1338_v62_, d_1338_v62_])
                d_1340_v64_: _dafny.Map
                d_1340_v64_ = _dafny.Map({d_1336_v60_: len(d_1339_v63_)})
                d_1341_v65_: _dafny.Set
                d_1341_v65_ = _dafny.Set({len(d_1340_v64_)})
                d_1342_v66_: _dafny.Map
                d_1342_v66_ = _dafny.Map({d_1341_v65_: (0) - ((self).f23)})
                index187_ = default__.safeIndex(910, (d_1249_v0_).length(0))
                (d_1249_v0_)[index187_] = (d_1252_v3_.f31) > (((d_1342_v66_)[d_1341_v65_] if (d_1341_v65_) in (d_1342_v66_) else (self).f23))
                index188_ = default__.safeIndex(910, (d_1249_v0_).length(0))
                (d_1249_v0_)[index188_] = self.f24
                d_1326_v50_ = d_1326_v50_
                (globalState).f15 = not(not((self).f25))
            elif True:
                index189_ = default__.safeIndex(793, (d_1317_cf39_).length(0))
                (d_1317_cf39_)[index189_] = (d_1252_v3_.f31) * (272)
                d_1343_v67_: str
                d_1343_v67_ = _dafny.CodePoint('g')
                d_1344_v68_: _dafny.Map
                d_1344_v68_ = _dafny.Map({(self).f25: (self).f25})
                index190_ = default__.safeIndex(793, (d_1317_cf39_).length(0))
                (d_1317_cf39_)[index190_] = len((_dafny.Map({self.f24: False})) | ((default__.fm10(self.f24, d_1252_v3_.f31, default__.fm12(False, globalState), d_1343_v67_, globalState)) | (d_1344_v68_)))
                d_1345_v69_: C5
                nw247_ = C5()
                nw247_.ctor__((6) == (d_1252_v3_.f31), (self).f25)
                d_1345_v69_ = nw247_
                d_1346_v70_: _dafny.Set
                d_1346_v70_ = _dafny.Set({(d_1262_v12_)[default__.safeIndex(default__.fm1(d_1252_v3_.f31, (self).f25, globalState), len(d_1262_v12_))], not(self.f24), self.f24, (self).f25})
                d_1346_v70_ = (d_1346_v70_).intersection(_dafny.Set({(self).f25, self.f24, not(self.f24)}))
                d_1347_v71_: _dafny.Set
                d_1347_v71_ = _dafny.Set({d_1252_v3_.f31, (self).f23})
                (globalState).f15 = not(((d_1347_v71_).intersection(d_1347_v71_)) != (d_1347_v71_))
                d_1348_v72_: _dafny.Array
                nw248_ = _dafny.Array(_dafny.Set({}), 11)
                d_1348_v72_ = nw248_
                nw249_ = _dafny.Array(None, 4)
                nw249_[int(0)] = d_1347_v71_
                def iife68_():
                    coll46_ = _dafny.Set()
                    compr_46_: int
                    for compr_46_ in _dafny.IntegerRange(988, 324):
                        d_1349_v73_: int = compr_46_
                        if ((988) <= (d_1349_v73_)) and ((d_1349_v73_) < (324)):
                            coll46_ = coll46_.union(_dafny.Set([default__.safeDivisionInt(d_1349_v73_, len(d_1346_v70_))]))
                    return _dafny.Set(coll46_)
                nw249_[int(1)] = iife68_()
                
                nw249_[int(2)] = (d_1347_v71_) - (d_1347_v71_)
                nw249_[int(3)] = (d_1347_v71_) | (d_1347_v71_)
                d_1348_v72_ = nw249_
            d_1350_v74_: T1
            nw250_ = C3()
            nw250_.ctor__((self).f25, d_1252_v3_.f31, (self).f25, self.f24, d_1252_v3_.f31)
            d_1350_v74_ = nw250_
            d_1351_v75_: _dafny.Seq
            d_1351_v75_ = _dafny.SeqWithoutIsStrInference([d_1350_v74_, d_1350_v74_])
            d_1352_v76_: _dafny.Map
            d_1352_v76_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([default__.fm29(len(d_1351_v75_), d_1252_v3_.f31, d_1350_v74_.f24, globalState)]): d_1263_v13_})
            index191_ = default__.safeIndex(66, (d_1249_v0_).length(0))
            (d_1249_v0_)[index191_] = (((d_1352_v76_)[d_1262_v12_] if (d_1262_v12_) in (d_1352_v76_) else d_1262_v12_)) <= (d_1262_v12_)
            index192_ = default__.safeIndex(66, (d_1249_v0_).length(0))
            (d_1249_v0_)[index192_] = (self).f25
            (globalState).f6 = self.f24
            if (d_1350_v74_).f25:
                d_1353_v77_: _dafny.Array
                def lambda69_(d_1354_v74_, d_1355_v3_):
                    def lambda70_(d_1356_i4_):
                        return (D10_DC29((self).f25, not((d_1354_v74_).f25), d_1355_v3_.f31) if (d_1354_v74_).f25 else D10_DC29(d_1354_v74_.f24, d_1354_v74_.f24, (self).f23))

                    return lambda70_

                init42_ = lambda69_(d_1350_v74_, d_1252_v3_)
                nw251_ = _dafny.Array(None, 28)
                for i0_42_ in range(nw251_.length(0)):
                    nw251_[i0_42_] = init42_(i0_42_)
                d_1353_v77_ = nw251_
                d_1357_v78_: _dafny.Seq
                d_1357_v78_ = _dafny.SeqWithoutIsStrInference([d_1252_v3_.f31])
                d_1358_v79_: _dafny.Seq
                d_1358_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aawnfu"))
                d_1359_v80_: _dafny.MultiSet
                d_1359_v80_ = _dafny.MultiSet([d_1357_v78_, _dafny.SeqWithoutIsStrInference([len(d_1358_v79_)])])
                d_1360_v81_: _dafny.Set
                d_1360_v81_ = _dafny.Set({True, True, not((self).f25)})
                d_1361_v82_: D10
                d_1361_v82_ = D10_DC29((self).fm6(d_1252_v3_.f31, d_1359_v80_, self.f24, len(d_1360_v81_), globalState), (d_1350_v74_).f25, (self).f23)
                index193_ = default__.safeIndex(939, (d_1353_v77_).length(0))
                (d_1353_v77_)[index193_] = d_1361_v82_
                index194_ = default__.safeIndex(66, (d_1249_v0_).length(0))
                index195_ = default__.safeIndex(939, (d_1353_v77_).length(0))
                rhs243_ = (d_1361_v82_).cf54
                rhs244_ = d_1361_v82_
                lhs205_ = d_1249_v0_
                lhs206_ = default__.safeIndex(66, (d_1249_v0_).length(0))
                lhs207_ = d_1353_v77_
                lhs208_ = default__.safeIndex(939, (d_1353_v77_).length(0))
                lhs205_[lhs206_] = rhs243_
                lhs207_[lhs208_] = rhs244_
                d_1362_v83_: _dafny.MultiSet
                d_1362_v83_ = _dafny.MultiSet([(d_1350_v74_).f25])
                index196_ = default__.safeIndex(169, (d_1264_v14_).length(0))
                (d_1264_v14_)[index196_] = (d_1362_v83_).cardinality
                index197_ = default__.safeIndex(169, (d_1264_v14_).length(0))
                rhs245_ = d_1358_v79_
                rhs246_ = ((self).f23) == ((self).f23)
                rhs247_ = ((d_1362_v83_)[not ((d_1350_v74_).f25) or ((self).f25)] if (not ((d_1350_v74_).f25) or ((self).f25)) in (d_1362_v83_) else d_1252_v3_.f31)
                rhs248_ = d_1252_v3_.f31
                lhs209_ = globalState
                lhs210_ = globalState
                lhs211_ = d_1264_v14_
                lhs212_ = default__.safeIndex(169, (d_1264_v14_).length(0))
                d_1358_v79_ = rhs245_
                lhs209_.f1 = rhs246_
                lhs210_.f5 = rhs247_
                lhs211_[lhs212_] = rhs248_
                d_1363_v84_: C4
                nw252_ = C4()
                nw252_.ctor__((d_1249_v0_)[default__.safeIndex(66, (d_1249_v0_).length(0))], d_1350_v74_.f24, self.f24)
                d_1363_v84_ = nw252_
                rhs249_ = ((d_1261_v11_).issubset(d_1261_v11_) if not((d_1362_v83_).issubset(d_1362_v83_)) else (d_1249_v0_)[default__.safeIndex(66, (d_1249_v0_).length(0))])
                rhs250_ = (d_1363_v84_).f27
                rhs251_ = (d_1262_v12_) + (d_1262_v12_)
                rhs252_ = d_1252_v3_.f31
                lhs213_ = globalState
                lhs214_ = globalState
                lhs215_ = globalState
                lhs213_.f15 = rhs249_
                lhs214_.f1 = rhs250_
                d_1262_v12_ = rhs251_
                lhs215_.f8 = rhs252_
                rhs253_ = d_1317_cf39_
                rhs254_ = ((self).f25) or ((self).f25)
                lhs216_ = globalState
                d_1317_cf39_ = rhs253_
                lhs216_.f1 = rhs254_
            elif True:
                d_1364_v85_: _dafny.Array
                def lambda71_(d_1365_v11_, d_1366_v74_):
                    def lambda72_(d_1367_i5_):
                        return (_dafny.MultiSet([d_1365_v11_, d_1365_v11_])) - (_dafny.MultiSet([_dafny.MultiSet([len(_dafny.Map({(self).f25: d_1366_v74_.f24})), 981])]))

                    return lambda72_

                init43_ = lambda71_(d_1261_v11_, d_1350_v74_)
                nw253_ = _dafny.Array(None, 23)
                for i0_43_ in range(nw253_.length(0)):
                    nw253_[i0_43_] = init43_(i0_43_)
                d_1364_v85_ = nw253_
                d_1368_v86_: _dafny.MultiSet
                d_1368_v86_ = _dafny.MultiSet([d_1261_v11_, d_1261_v11_])
                index198_ = default__.safeIndex(886, (d_1364_v85_).length(0))
                (d_1364_v85_)[index198_] = (d_1368_v86_ if d_1350_v74_.f24 else d_1368_v86_)
                index199_ = default__.safeIndex(886, (d_1364_v85_).length(0))
                (d_1364_v85_)[index199_] = d_1368_v86_
                d_1251_v2_ = d_1251_v2_
                d_1369_v87_: str
                d_1369_v87_ = _dafny.CodePoint('f')
                d_1370_v88_: _dafny.Map
                d_1370_v88_ = _dafny.Map({d_1369_v87_: (self).f23})
                d_1371_v89_: _dafny.Seq
                d_1371_v89_ = _dafny.SeqWithoutIsStrInference([d_1370_v88_])
                d_1372_v90_: D9
                d_1372_v90_ = D9_DC27(D9_DC25(d_1371_v89_))
                d_1373_v91_: D9
                d_1373_v91_ = D9_DC25(d_1371_v89_)
                d_1372_v90_ = D9_DC27(d_1373_v91_)
                d_1374_v92_: _dafny.Array
                nw254_ = _dafny.Array(None, 24)
                d_1374_v92_ = nw254_
                d_1374_v92_ = d_1374_v92_
                d_1375_v93_: _dafny.Array
                nw255_ = _dafny.Array(D6.default()(), 2)
                d_1375_v93_ = nw255_
                def lambda73_(d_1376_v74_):
                    def lambda74_(d_1377_i6_):
                        return (D6_DC17((d_1376_v74_).f25) if not(False) else D6_DC17(d_1376_v74_.f24))

                    return lambda74_

                init44_ = lambda73_(d_1350_v74_)
                nw256_ = _dafny.Array(None, 21)
                for i0_44_ in range(nw256_.length(0)):
                    nw256_[i0_44_] = init44_(i0_44_)
                d_1375_v93_ = nw256_
        d_1378_v94_: _dafny.Seq
        d_1378_v94_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erevcsyu"))
        d_1378_v94_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
        if self.f24:
            d_1379_v95_: _dafny.Map
            d_1379_v95_ = _dafny.Map({83: d_1249_v0_})
            d_1380_v96_: C3
            nw257_ = C3()
            nw257_.ctor__(self.f24, len(d_1378_v94_), self.f24, (self).f25, len(d_1379_v95_))
            d_1380_v96_ = nw257_
            d_1381_v97_: _dafny.Map
            d_1381_v97_ = _dafny.Map({(d_1380_v96_).f28: d_1380_v96_})
            rhs255_ = self.f24
            rhs256_ = len((_dafny.Map({self.f24: d_1380_v96_})) | (d_1381_v97_))
            rhs257_ = not(self.f24)
            lhs217_ = globalState
            lhs218_ = globalState
            lhs219_ = globalState
            lhs217_.f9 = rhs255_
            lhs218_.f8 = rhs256_
            lhs219_.f1 = rhs257_
            d_1382_v98_: str
            d_1382_v98_ = _dafny.CodePoint('k')
            d_1383_v99_: D0
            d_1383_v99_ = D0_DC1((d_1380_v96_).f28, len(default__.fm20((d_1380_v96_).f28, globalState)))
            d_1384_v100_: D0
            d_1384_v100_ = D0_DC2(d_1383_v99_)
            d_1385_v101_: C6
            nw258_ = C6()
            nw258_.ctor__(d_1384_v100_, (d_1380_v96_).f28, True, 219)
            d_1385_v101_ = nw258_
            d_1386_v102_: _dafny.MultiSet
            d_1386_v102_ = _dafny.MultiSet([d_1385_v101_])
            pat_let_tv22_ = d_1383_v99_
            d_1387_v103_: _dafny.Array
            nw259_ = _dafny.Array(None, 6)
            nw259_[int(0)] = D0_DC2(d_1383_v99_)
            nw259_[int(1)] = d_1384_v100_
            nw259_[int(2)] = d_1384_v100_
            nw259_[int(3)] = d_1384_v100_
            nw259_[int(4)] = d_1384_v100_
            def iife69_(_pat_let11_0):
                def iife70_(d_1388_dt__update__tmp_h0_):
                    def iife71_(_pat_let12_0):
                        def iife72_(d_1389_dt__update_hcf3_h0_):
                            return D0_DC2(d_1389_dt__update_hcf3_h0_)
                        return iife72_(_pat_let12_0)
                    return iife71_(pat_let_tv22_)
                return iife70_(_pat_let11_0)
            nw259_[int(5)] = iife69_(d_1385_v101_.f26)
            d_1387_v103_ = nw259_
            d_1390_v104_: _dafny.Map
            d_1390_v104_ = _dafny.Map({d_1387_v103_: d_1382_v98_})
            d_1391_v105_: D3
            d_1391_v105_ = D3_DC9(101)
            d_1392_v106_: _dafny.Map
            d_1392_v106_ = _dafny.Map({d_1391_v105_: d_1264_v14_})
            d_1393_v107_: D5
            d_1393_v107_ = D5_DC15(((d_1251_v2_)[(d_1386_v102_).cardinality] if ((d_1386_v102_).cardinality) in (d_1251_v2_) else (self).f23), d_1382_v98_, d_1390_v104_, d_1392_v106_)
            d_1394_v108_: _dafny.Array
            nw260_ = _dafny.Array(None, 18)
            nw260_[int(0)] = d_1382_v98_
            nw260_[int(1)] = d_1382_v98_
            nw260_[int(2)] = d_1382_v98_
            nw260_[int(3)] = d_1382_v98_
            nw260_[int(4)] = d_1382_v98_
            nw260_[int(5)] = d_1382_v98_
            nw260_[int(6)] = d_1382_v98_
            nw260_[int(7)] = d_1382_v98_
            nw260_[int(8)] = _dafny.CodePoint('g')
            nw260_[int(9)] = d_1382_v98_
            nw260_[int(10)] = d_1382_v98_
            nw260_[int(11)] = d_1382_v98_
            nw260_[int(12)] = _dafny.CodePoint('t')
            nw260_[int(13)] = d_1382_v98_
            nw260_[int(14)] = (d_1393_v107_).cf26
            nw260_[int(15)] = d_1382_v98_
            nw260_[int(16)] = _dafny.CodePoint('y')
            nw260_[int(17)] = d_1382_v98_
            d_1394_v108_ = nw260_
            d_1395_v109_: _dafny.Set
            d_1395_v109_ = _dafny.Set({d_1394_v108_, d_1394_v108_})
            d_1396_v110_: D5
            d_1396_v110_ = D5_DC14(d_1252_v3_.f31, ((self).f23) <= (len(d_1378_v94_)), d_1395_v109_)
            d_1397_v111_: _dafny.Seq
            d_1397_v111_ = _dafny.SeqWithoutIsStrInference([(self).f23, (self).f23])
            pat_let_tv23_ = d_1397_v111_
            pat_let_tv24_ = d_1252_v3_
            pat_let_tv25_ = d_1397_v111_
            def iife73_(_pat_let13_0):
                def iife74_(d_1398_dt__update__tmp_h1_):
                    def iife75_(_pat_let14_0):
                        def iife76_(d_1399_dt__update_hcf23_h0_):
                            def iife77_(_pat_let15_0):
                                def iife78_(d_1400_dt__update_hcf22_h0_):
                                    return D5_DC14(d_1400_dt__update_hcf22_h0_, d_1399_dt__update_hcf23_h0_, (d_1398_dt__update__tmp_h1_).cf24)
                                return iife78_(_pat_let15_0)
                            return iife77_((pat_let_tv23_)[default__.safeIndex(pat_let_tv24_.f31, len(pat_let_tv25_))])
                        return iife76_(_pat_let14_0)
                    return iife75_(self.f24)
                return iife74_(_pat_let13_0)
            d_1396_v110_ = iife73_(d_1396_v110_)
            d_1401_v113_: _dafny.Map
            d_1401_v113_ = _dafny.Map({999: (d_1380_v96_).f28})
            d_1402_v114_: _dafny.Seq
            def iife79_():
                coll47_ = _dafny.Map()
                compr_47_: int
                for compr_47_ in (d_1401_v113_).keys.Elements:
                    d_1403_v112_: int = compr_47_
                    if (d_1403_v112_) in (d_1401_v113_):
                        coll47_[default__.safeModuloInt(d_1403_v112_, (d_1380_v96_).f29)] = d_1252_v3_.f31
                return _dafny.Map(coll47_)
            d_1402_v114_ = _dafny.SeqWithoutIsStrInference([d_1251_v2_, iife79_()
            , d_1251_v2_, _dafny.Map({d_1252_v3_.f31: (d_1380_v96_).f29}), _dafny.Map({(self).f23: d_1252_v3_.f31})])
            d_1402_v114_ = ((_dafny.SeqWithoutIsStrInference([d_1251_v2_])).set(default__.safeIndex((d_1380_v96_).f29, len(_dafny.SeqWithoutIsStrInference([d_1251_v2_]))), d_1251_v2_)) + ((d_1402_v114_) + (d_1402_v114_))
            d_1404_v115_: _dafny.Array
            nw261_ = _dafny.Array(_dafny.Seq({}), 8)
            d_1404_v115_ = nw261_
            index200_ = default__.safeIndex(470, (d_1404_v115_).length(0))
            (d_1404_v115_)[index200_] = d_1262_v12_
            index201_ = default__.safeIndex(470, (d_1404_v115_).length(0))
            (d_1404_v115_)[index201_] = _dafny.SeqWithoutIsStrInference([(self).f25, not((d_1262_v12_)[default__.safeIndex((d_1380_v96_).f29, len(d_1262_v12_))]), ((d_1380_v96_).f28 if (d_1380_v96_).f28 else (d_1380_v96_).f28), (d_1380_v96_).f28])
            d_1405_v116_: _dafny.Array
            nw262_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_1405_v116_ = nw262_
            d_1406_v117_: C1
            nw263_ = C1()
            nw263_.ctor__((d_1380_v96_).f28, True)
            d_1406_v117_ = nw263_
            d_1407_v118_: _dafny.Map
            d_1407_v118_ = _dafny.Map({d_1406_v117_: d_1382_v98_})
            index202_ = default__.safeIndex(955, (d_1264_v14_).length(0))
            (d_1264_v14_)[index202_] = ((d_1380_v96_).f29) * (d_1252_v3_.f31)
            index203_ = default__.safeIndex(955, (d_1264_v14_).length(0))
            rhs258_ = (d_1380_v96_).fm5(d_1252_v3_.f31, (self).f23, globalState)
            rhs259_ = d_1405_v116_
            rhs260_ = d_1252_v3_.f31
            rhs261_ = d_1407_v118_
            rhs262_ = d_1252_v3_.f31
            lhs220_ = globalState
            lhs221_ = globalState
            lhs222_ = d_1264_v14_
            lhs223_ = default__.safeIndex(955, (d_1264_v14_).length(0))
            lhs220_.f14 = rhs258_
            d_1405_v116_ = rhs259_
            lhs221_.f8 = rhs260_
            d_1407_v118_ = rhs261_
            lhs222_[lhs223_] = rhs262_
        elif True:
            d_1408_v119_: str
            d_1408_v119_ = _dafny.CodePoint('d')
            d_1409_v120_: _dafny.Map
            d_1409_v120_ = _dafny.Map({d_1252_v3_.f31: default__.fm11(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdnmeq"))), self.f24, self.f24, d_1408_v119_, globalState)})
            d_1410_v121_: D1
            d_1410_v121_ = D1_DC4((self).f25, (self).f25)
            (self).f24 = ((d_1409_v120_).set(d_1252_v3_.f31, d_1410_v121_)) != (d_1409_v120_)
            d_1411_v122_: _dafny.Seq
            d_1411_v122_ = _dafny.SeqWithoutIsStrInference([d_1378_v94_])
            d_1412_v123_: _dafny.Seq
            d_1412_v123_ = _dafny.SeqWithoutIsStrInference([d_1411_v122_])
            rhs263_ = (d_1411_v122_) + ((d_1412_v123_)[default__.safeIndex(d_1252_v3_.f31, len(d_1412_v123_))])
            rhs264_ = d_1408_v119_
            d_1411_v122_ = rhs263_
            d_1408_v119_ = rhs264_
            d_1413_v124_: C7
            nw264_ = C7()
            nw264_.ctor__()
            d_1413_v124_ = nw264_
            d_1414_v125_: _dafny.MultiSet
            d_1414_v125_ = _dafny.MultiSet([d_1413_v124_, d_1413_v124_])
            d_1415_v126_: _dafny.Seq
            d_1415_v126_ = _dafny.SeqWithoutIsStrInference([(self).f23, (d_1252_v3_.f31) * ((d_1414_v125_).cardinality), (d_1252_v3_.f31 if not(self.f24) else d_1252_v3_.f31), d_1252_v3_.f31, (d_1252_v3_.f31) + ((self).f23)])
            d_1415_v126_ = _dafny.SeqWithoutIsStrInference([(self).f23])
            if (d_1252_v3_.f31) < ((d_1261_v11_).cardinality):
                d_1416_v127_: bool
                d_1417_v128_: bool
                d_1418_v129_: D0
                out14_: bool
                out15_: bool
                out16_: D0
                out14_, out15_, out16_ = (d_1413_v124_).m2((self).f23, ((d_1261_v11_).set(745, default__.abs(len(d_1251_v2_)))).ispropersubset(d_1261_v11_), globalState)
                d_1416_v127_ = out14_
                d_1417_v128_ = out15_
                d_1418_v129_ = out16_
                d_1419_v130_: C7
                nw265_ = C7()
                nw265_.ctor__()
                d_1419_v130_ = nw265_
                (globalState).f1 = self.f24
                d_1416_v127_ = self.f24
                d_1417_v128_ = self.f24
            elif True:
                index204_ = default__.safeIndex(461, (d_1264_v14_).length(0))
                def iife80_():
                    coll48_ = _dafny.Map()
                    compr_48_: int
                    for compr_48_ in _dafny.IntegerRange(421, 182):
                        d_1420_v131_: int = compr_48_
                        if ((421) <= (d_1420_v131_)) and ((d_1420_v131_) < (182)):
                            coll48_[default__.safeDivisionInt(d_1420_v131_, len(d_1378_v94_))] = self.f24
                    return _dafny.Map(coll48_)
                (d_1264_v14_)[index204_] = len((iife80_()
                ).set(d_1252_v3_.f31, (self).f25))
                index205_ = default__.safeIndex(461, (d_1264_v14_).length(0))
                (d_1264_v14_)[index205_] = default__.safeModuloInt((self).f23, (self).f23)
                index206_ = default__.safeIndex(531, (d_1249_v0_).length(0))
                (d_1249_v0_)[index206_] = (self).f25
                index207_ = default__.safeIndex(531, (d_1249_v0_).length(0))
                (d_1249_v0_)[index207_] = self.f24
                d_1421_v132_: C3
                nw266_ = C3()
                nw266_.ctor__((d_1249_v0_)[default__.safeIndex(531, (d_1249_v0_).length(0))], d_1252_v3_.f31, (d_1249_v0_)[default__.safeIndex(531, (d_1249_v0_).length(0))], (d_1249_v0_)[default__.safeIndex(531, (d_1249_v0_).length(0))], d_1252_v3_.f31)
                d_1421_v132_ = nw266_
                d_1422_v133_: _dafny.Seq
                d_1422_v133_ = _dafny.SeqWithoutIsStrInference([d_1421_v132_])
                d_1423_v134_: _dafny.Seq
                d_1423_v134_ = _dafny.SeqWithoutIsStrInference([d_1422_v133_, d_1422_v133_])
                d_1424_v135_: _dafny.Array
                nw267_ = _dafny.Array(None, 12)
                nw267_[int(0)] = (d_1422_v133_) + (d_1422_v133_)
                nw267_[int(1)] = d_1422_v133_
                nw267_[int(2)] = (d_1423_v134_)[default__.safeIndex((self).f23, len(d_1423_v134_))]
                nw267_[int(3)] = (d_1422_v133_) + (_dafny.SeqWithoutIsStrInference([d_1421_v132_]))
                nw267_[int(4)] = d_1422_v133_
                nw267_[int(5)] = d_1422_v133_
                nw267_[int(6)] = d_1422_v133_
                nw267_[int(7)] = (d_1422_v133_ if (d_1421_v132_).f28 else _dafny.SeqWithoutIsStrInference([d_1421_v132_]))
                nw267_[int(8)] = d_1422_v133_
                nw267_[int(9)] = d_1422_v133_
                nw267_[int(10)] = d_1422_v133_
                nw267_[int(11)] = (d_1422_v133_).set(default__.safeIndex(d_1252_v3_.f31, len(d_1422_v133_)), d_1421_v132_)
                d_1424_v135_ = nw267_
                d_1425_v136_: C1
                nw268_ = C1()
                nw268_.ctor__((d_1421_v132_).f28, not((d_1421_v132_).f28))
                d_1425_v136_ = nw268_
                rhs265_ = d_1252_v3_.f31
                rhs266_ = d_1424_v135_
                rhs267_ = ((d_1252_v3_.f31) - (d_1252_v3_.f31)) * (d_1252_v3_.f31)
                rhs268_ = d_1425_v136_
                lhs224_ = globalState
                lhs225_ = globalState
                lhs224_.f5 = rhs265_
                d_1424_v135_ = rhs266_
                lhs225_.f5 = rhs267_
                d_1425_v136_ = rhs268_
                d_1426_v137_: T1
                nw269_ = C5()
                nw269_.ctor__((d_1249_v0_)[default__.safeIndex(531, (d_1249_v0_).length(0))], (d_1249_v0_)[default__.safeIndex(531, (d_1249_v0_).length(0))])
                d_1426_v137_ = nw269_
                d_1427_v138_: _dafny.Map
                d_1427_v138_ = _dafny.Map({d_1426_v137_: d_1426_v137_.f24})
                d_1428_v139_: _dafny.Map
                d_1428_v139_ = _dafny.Map({(d_1249_v0_)[default__.safeIndex(531, (d_1249_v0_).length(0))]: d_1427_v138_})
                d_1429_v140_: _dafny.Array
                nw270_ = _dafny.Array(D0.default()(), 3)
                d_1429_v140_ = nw270_
                d_1430_v141_: _dafny.Map
                d_1430_v141_ = _dafny.Map({d_1429_v140_: d_1408_v119_})
                d_1431_v142_: D3
                d_1431_v142_ = D3_DC9((d_1264_v14_)[default__.safeIndex(461, (d_1264_v14_).length(0))])
                d_1432_v143_: _dafny.Array
                nw271_ = _dafny.Array(int(0), 24)
                d_1432_v143_ = nw271_
                d_1433_v144_: _dafny.Map
                d_1433_v144_ = _dafny.Map({d_1431_v142_: d_1432_v143_})
                d_1434_v145_: D5
                d_1434_v145_ = D5_DC15(len((d_1428_v139_).set(False, d_1427_v138_)), d_1408_v119_, d_1430_v141_, (d_1433_v144_) | (d_1433_v144_))
                d_1435_v146_: _dafny.Set
                d_1435_v146_ = _dafny.Set({d_1408_v119_, d_1408_v119_})
                d_1434_v145_ = (d_1434_v145_ if (d_1262_v12_)[default__.safeIndex(len(d_1435_v146_), len(d_1262_v12_))] else d_1434_v145_)
                d_1436_v147_: D0
                d_1436_v147_ = D0_DC0((d_1421_v132_).f28)
                d_1437_v148_: D0
                d_1437_v148_ = D0_DC2(d_1436_v147_)
                d_1438_v149_: _dafny.Map
                d_1438_v149_ = _dafny.Map({(d_1421_v132_).f29: self.f24})
                d_1439_v150_: C6
                nw272_ = C6()
                nw272_.ctor__((d_1437_v148_ if (d_1249_v0_)[default__.safeIndex(531, (d_1249_v0_).length(0))] else D0_DC2(D0_DC0(((d_1438_v149_)[(d_1421_v132_).f29] if ((d_1421_v132_).f29) in (d_1438_v149_) else self.f24)))), (d_1421_v132_).f28, (d_1426_v137_).f25, d_1252_v3_.f31)
                d_1439_v150_ = nw272_
            (globalState).f15 = (self).f25
        d_1440_v151_: _dafny.Seq
        d_1440_v151_ = _dafny.SeqWithoutIsStrInference([171, (0) - ((0) - (d_1252_v3_.f31)), d_1252_v3_.f31, d_1252_v3_.f31, 443])
        r0 = (d_1252_v3_).fm3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1441_i7_ in range(default__.abs(-291))]), self.f24, len(d_1440_v151_), (d_1440_v151_ if not((self).f25) else (d_1440_v151_).set(default__.safeIndex((0) - (-625), len(d_1440_v151_)), (self).f23)), globalState)
        d_1442_v152_: T0
        nw273_ = C2()
        nw273_.ctor__((0) - (d_1252_v3_.f31))
        d_1442_v152_ = nw273_
        r1 = d_1442_v152_
        r2 = (self).f25
        r3 = default__.fm1((self).fm2(globalState), False, globalState)
        return r0, r1, r2, r3

