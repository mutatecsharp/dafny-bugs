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
        return _dafny.SeqWithoutIsStrInference([616, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_0_i0_ in range(default__.abs(-434))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1_i1_ in range(default__.abs(26))])))])

    @staticmethod
    def fm1(p0, globalState):
        def iife0_():
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: str
                for compr_2_ in (_dafny.Set({_dafny.CodePoint('d')})).Elements:
                    d_3_v0_: str = compr_2_
                    if (d_3_v0_) in (_dafny.Set({_dafny.CodePoint('d')})):
                        coll2_[d_3_v0_] = False
                return _dafny.Map(coll2_)
            coll0_ = _dafny.Set()
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: str
                for compr_1_ in (_dafny.Set({_dafny.CodePoint('d')})).Elements:
                    d_3_v0_: str = compr_1_
                    if (d_3_v0_) in (_dafny.Set({_dafny.CodePoint('d')})):
                        coll1_[d_3_v0_] = False
                return _dafny.Map(coll1_)
            compr_0_: str
            for compr_0_ in (iife1_()
            ).keys.Elements:
                d_4_v1_: str = compr_0_
                if (d_4_v1_) in (iife2_()
                ):
                    coll0_ = coll0_.union(_dafny.Set([d_4_v1_]))
            return _dafny.Set(coll0_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([356 for d_2_i0_ in range(default__.abs(582))]))])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([343]), _dafny.MultiSet([(0) - (-843)]), _dafny.MultiSet([len(iife0_()
        )])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True]))]) for d_5_i1_ in range(default__.abs(863))])))

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return not (not((not(False)) or (not(False)))) or ((_dafny.Set({_dafny.CodePoint('t')})).ispropersubset(_dafny.Set({_dafny.CodePoint('k')})))

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        return ((884 if False else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjmw"))))) * ((len(_dafny.SeqWithoutIsStrInference([True, not(True)])) if True else -443))

    @staticmethod
    def fm4(p0, p1, p2, globalState):
        return _dafny.Map({D0_DC0(not(True)): False})

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return D2_DC5()

    @staticmethod
    def fm9(globalState):
        return D2_DC6(D2_DC6(D2_DC6(D2_DC4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bx")), True))))

    @staticmethod
    def fm10(p0, p1, globalState):
        return (_dafny.Set({True})) - ((_dafny.Set({False})) | (_dafny.Set({False})))

    @staticmethod
    def fm11(p0, globalState):
        return (_dafny.Map({465: 64})) | (_dafny.Map({-896: 863}))

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True, False, not(True)])))

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        if not(False):
            return D2_DC6(D2_DC5())
        elif True:
            return D2_DC6(D2_DC3(_dafny.SeqWithoutIsStrInference([D0_DC1(True), D0_DC1(False), D0_DC1(False)])))

    @staticmethod
    def fm17(p0, globalState):
        return _dafny.CodePoint('s')

    @staticmethod
    def fm18(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vn"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pyh")))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_6_i0_ in range(default__.abs(513))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nt"))))

    @staticmethod
    def fm19(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iql"))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(-443, -121):
                d_7_v0_: int = compr_3_
                if ((-443) <= (d_7_v0_)) and ((d_7_v0_) < (-121)):
                    coll3_ = coll3_.union(_dafny.Set([(d_7_v0_) - ((_dafny.MultiSet([len(_dafny.Set({False}))])).cardinality)]))
            return _dafny.Set(coll3_)
        return _dafny.Map({(len(_dafny.Map({not(False): not(False)}))) > (len(_dafny.SeqWithoutIsStrInference([iife3_()
        , _dafny.Set({917, -461}), _dafny.Set({(_dafny.MultiSet([False, not(not(True))])).cardinality}), _dafny.Set({-841, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_8_i0_ in range(default__.abs(193))])), 36}), _dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([False, True])))})]))): (_dafny.CodePoint('j')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))})

    @staticmethod
    def fm21(globalState):
        return (_dafny.Map({not(True): (_dafny.MultiSet([-222, 310])).cardinality})) | ((_dafny.Map({False: len(_dafny.Set({True, False, False}))})) | (_dafny.Map({not(not(not(not(False)))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sp")))})))

    @staticmethod
    def fm22(globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([303]))) | (_dafny.MultiSet([965]))) | (_dafny.MultiSet([610]))

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        return D5_DC13((392 if False else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntyxfwen")))), _dafny.Set({False}), _dafny.CodePoint('g'), -831, False)

    @staticmethod
    def fm24(p0, p1, p2, p3, globalState):
        if False:
            return (_dafny.MultiSet([D2_DC6(D2_DC3(_dafny.SeqWithoutIsStrInference([D0_DC1(False), D0_DC1(True)])))])).intersection(_dafny.MultiSet([D2_DC6(D2_DC5())]))
        elif True:
            return _dafny.MultiSet([D2_DC6(D2_DC6(D2_DC6(D2_DC3(_dafny.SeqWithoutIsStrInference([D0_DC1(True), D0_DC1(not(False)), D0_DC1(True), D0_DC1(True)]))))), D2_DC6(D2_DC3(_dafny.SeqWithoutIsStrInference([D0_DC1(True), D0_DC1(True)])))])

    @staticmethod
    def fm25(globalState):
        return _dafny.Map({False: False})

    @staticmethod
    def fm26(globalState):
        return D7_DC18(default__.safeDivisionInt(698, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "todmnh")))))

    @staticmethod
    def fm27(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({False, not(not(True))})) - (_dafny.Set({False, not(False)}))) | (_dafny.Set({False}))

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([not(False)])) | (_dafny.MultiSet([False, not(True), False]))).intersection((_dafny.MultiSet([not(not(True))]) if False else _dafny.MultiSet([False])))

    @staticmethod
    def fm29(p0, p1, p2, p3, globalState):
        return _dafny.Map({(_dafny.MultiSet([121, 39])).intersection(_dafny.MultiSet([len(_dafny.Map({953: False}))])): (532) == (len(_dafny.Set({len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")))}))})))})

    @staticmethod
    def fm30(p0, p1, p2, globalState):
        return (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([True])})) | ((_dafny.Map({not(False): _dafny.SeqWithoutIsStrInference([True])})) | (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([False, False, True, True, False])})))

    @staticmethod
    def fm31(globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: str
            for compr_4_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_9_i0_ in range(default__.abs(492))])).Elements:
                d_10_v0_: str = compr_4_
                if (d_10_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_9_i0_ in range(default__.abs(492))])):
                    coll4_ = coll4_.union(_dafny.Set([d_10_v0_]))
            return _dafny.Set(coll4_)
        return (_dafny.Set({_dafny.CodePoint('l')})) | ((_dafny.Set({_dafny.CodePoint('i'), _dafny.CodePoint('x')}) if True else iife4_()
        ))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: _dafny.Set
            for compr_5_ in (_dafny.MultiSet([_dafny.Set({False})])).Elements:
                d_12_v0_: _dafny.Set = compr_5_
                if (d_12_v0_) in (_dafny.MultiSet([_dafny.Set({False})])):
                    coll5_[d_12_v0_] = False
            return _dafny.Map(coll5_)
        return (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-625])), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([580 for d_11_i0_ in range(default__.abs(454))]))]) if not(False) else _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([116, 990]), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sng")))]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iajxrdi")))])), (_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pocfxd"))): False}))]))])))) - ((_dafny.MultiSet([_dafny.MultiSet([973, len(iife5_()
        ), 59, 168])])).intersection(_dafny.MultiSet([_dafny.MultiSet([-520, 279]), _dafny.MultiSet([len(_dafny.Map({150: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sq")))]))})), 789, len(_dafny.SeqWithoutIsStrInference([True])), (_dafny.MultiSet([-224, (_dafny.MultiSet([True])).cardinality, len(_dafny.Set({False, True}))])).cardinality])])))

    @staticmethod
    def fm33(p0, p1, globalState):
        return (_dafny.Map({-871: D18_DC39(True)})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbpaa"))): D18_DC39(not(True))}))

    @staticmethod
    def fm34(globalState):
        def iife6_():
            def iife8_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_13_i1_ in range(default__.abs(580))])) for d_14_i0_ in range(default__.abs(209))])).Elements:
                    d_15_v0_: int = compr_8_
                    if (d_15_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_13_i1_ in range(default__.abs(580))])) for d_14_i0_ in range(default__.abs(209))])):
                        coll8_[(d_15_v0_) - (91)] = D13_DC28(False, _dafny.CodePoint('o'))
                return _dafny.Map(coll8_)
            coll6_ = _dafny.Set()
            def iife7_():
                coll7_ = _dafny.Map()
                compr_7_: int
                for compr_7_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_13_i1_ in range(default__.abs(580))])) for d_14_i0_ in range(default__.abs(209))])).Elements:
                    d_15_v0_: int = compr_7_
                    if (d_15_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_13_i1_ in range(default__.abs(580))])) for d_14_i0_ in range(default__.abs(209))])):
                        coll7_[(d_15_v0_) - (91)] = D13_DC28(False, _dafny.CodePoint('o'))
                return _dafny.Map(coll7_)
            compr_6_: int
            for compr_6_ in ((_dafny.Map({573: D13_DC28(True, _dafny.CodePoint('j'))})) | (iife7_()
            )).keys.Elements:
                d_16_v1_: int = compr_6_
                if (d_16_v1_) in ((_dafny.Map({573: D13_DC28(True, _dafny.CodePoint('j'))})) | (iife8_()
                )):
                    coll6_ = coll6_.union(_dafny.Set([default__.safeDivisionInt(d_16_v1_, 316)]))
            return _dafny.Set(coll6_)
        return iife6_()
        

    @staticmethod
    def fm35(p0, p1, p2, globalState):
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(511, 620):
                d_17_v0_: int = compr_9_
                if ((511) <= (d_17_v0_)) and ((d_17_v0_) < (620)):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeModuloInt(d_17_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oyikk"))))]))
            return _dafny.Set(coll9_)
        return _dafny.SeqWithoutIsStrInference([_dafny.Set({(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([True])))), len(_dafny.SeqWithoutIsStrInference([True]))}), _dafny.Set({len(_dafny.Set({31, 724}))}), iife9_()
        ])

    @staticmethod
    def m0(p0, p1, globalState):
        r0: bool = False
        d_18_v0_: str
        d_18_v0_ = _dafny.CodePoint('q')
        d_19_v1_: D13
        d_19_v1_ = D13_DC28(p0, d_18_v0_)
        source0_ = d_19_v1_
        if source0_.is_DC27:
            d_20___mcc_h0_ = source0_.cf43
            d_21___mcc_h1_ = source0_.cf44
            d_22___mcc_h2_ = source0_.cf45
            d_23___mcc_h3_ = source0_.cf46
            d_24_cf46_ = d_23___mcc_h3_
            d_25_cf45_ = d_22___mcc_h2_
            d_26_cf44_ = d_21___mcc_h1_
            d_27_cf43_ = d_20___mcc_h0_
            d_28_v2_: _dafny.Array
            nw0_ = _dafny.Array(False, 3)
            d_28_v2_ = nw0_
            d_29_v3_: D11
            d_29_v3_ = D11_DC24(d_28_v2_, p0, p0)
            d_30_v4_: _dafny.Set
            d_30_v4_ = _dafny.Set({d_29_v3_})
            d_31_v5_: _dafny.Array
            d_31_v5_ = d_28_v2_
            d_32_v6_: _dafny.Map
            d_32_v6_ = _dafny.Map({p1: p0})
            d_33_v7_: _dafny.Seq
            d_33_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvt"))
            d_34_v8_: D18
            d_34_v8_ = D18_DC38(d_30_v4_)
            d_35_v9_: _dafny.Array
            nw1_ = _dafny.Array(None, 20)
            nw1_[int(0)] = (d_30_v4_) - (d_30_v4_)
            nw1_[int(1)] = (d_30_v4_ if p0 else d_30_v4_)
            nw1_[int(2)] = (_dafny.Set({d_29_v3_})).intersection(d_30_v4_)
            nw1_[int(3)] = (_dafny.Set({D11_DC24((d_31_v5_), p0, p0)})) - (d_30_v4_)
            nw1_[int(4)] = d_30_v4_
            nw1_[int(5)] = d_30_v4_
            nw1_[int(6)] = d_30_v4_
            nw1_[int(7)] = d_30_v4_
            nw1_[int(8)] = _dafny.Set({D11_DC24(d_28_v2_, p0, p0)})
            nw1_[int(9)] = d_30_v4_
            nw1_[int(10)] = d_30_v4_
            nw1_[int(11)] = d_30_v4_
            nw1_[int(12)] = d_30_v4_
            nw1_[int(13)] = d_30_v4_
            nw1_[int(14)] = d_30_v4_
            nw1_[int(15)] = d_30_v4_
            nw1_[int(16)] = d_30_v4_
            nw1_[int(17)] = _dafny.Set({D11_DC24(d_28_v2_, ((d_32_v6_)[(d_26_cf44_).f0] if ((d_26_cf44_).f0) in (d_32_v6_) else True), p0), d_29_v3_, D11_DC24(d_28_v2_, p0, (d_26_cf44_).fm6(p0, d_33_v7_, d_33_v7_, d_24_cf46_, globalState))})
            nw1_[int(18)] = _dafny.Set({d_29_v3_})
            nw1_[int(19)] = ((d_34_v8_).cf58) | (d_30_v4_)
            d_35_v9_ = nw1_
            d_35_v9_ = d_35_v9_
            d_36_v10_: _dafny.Map
            d_36_v10_ = _dafny.Map({_dafny.Map({True: p0}): p1})
            d_37_v11_: _dafny.Map
            d_37_v11_ = _dafny.Map({(d_26_cf44_).f0: (d_26_cf44_).f0})
            d_38_v12_: C5
            nw2_ = C5()
            nw2_.ctor__(d_25_cf45_, (0) - ((D5_DC12(p1, ((d_36_v10_)[_dafny.Map({p0: p0})] if (_dafny.Map({p0: p0})) in (d_36_v10_) else p1))).cf15), len(d_37_v11_))
            d_38_v12_ = nw2_
            d_39_v13_: C2
            nw3_ = C2()
            nw3_.ctor__(len(d_33_v7_))
            d_39_v13_ = nw3_
            rhs0_ = d_38_v12_
            rhs1_ = d_39_v13_
            rhs2_ = (d_38_v12_).f2
            rhs3_ = p0
            rhs4_ = ((d_38_v12_).f2) == (default__.safeModuloInt(p1, p1))
            d_38_v12_ = rhs0_
            d_39_v13_ = rhs1_
            d_24_cf46_ = rhs2_
            r0 = rhs3_
            r0 = rhs4_
            d_40_v14_: _dafny.Set
            d_40_v14_ = _dafny.Set({(d_38_v12_).f2})
            d_41_v15_: D3
            d_41_v15_ = D3_DC8(p0, (d_38_v12_).f2)
            d_42_v16_: D7
            d_42_v16_ = D7_DC17(p0, p0, p0)
            d_43_v17_: _dafny.Set
            d_43_v17_ = _dafny.Set({d_42_v16_})
            rhs5_ = default__.safeModuloInt((0) - ((d_38_v12_).f2), len(d_40_v14_))
            rhs6_ = not((d_38_v12_).fm6(p0, d_33_v7_, (D2_DC4(d_33_v7_, p0)).cf4, (d_41_v15_).cf9, globalState))
            rhs7_ = (d_38_v12_).f2
            rhs8_ = not((_dafny.MultiSet([d_43_v17_, d_43_v17_])).ispropersubset(_dafny.MultiSet([d_43_v17_])))
            d_24_cf46_ = rhs5_
            r0 = rhs6_
            d_24_cf46_ = rhs7_
            r0 = rhs8_
            d_44_v18_: C4
            nw4_ = C4()
            nw4_.ctor__(default__.fm3(_dafny.Set({p0}), (d_38_v12_).f1, -553, 258, globalState))
            d_44_v18_ = nw4_
        elif source0_.is_DC28:
            d_45___mcc_h4_ = source0_.cf47
            d_46___mcc_h5_ = source0_.cf48
            d_47_cf48_ = d_46___mcc_h5_
            d_48_cf47_ = d_45___mcc_h4_
            if (d_48_cf47_ if True else d_48_cf47_):
                d_49_v19_: T1
                nw5_ = C5()
                nw5_.ctor__(d_47_cf48_, 515, 237)
                d_49_v19_ = nw5_
                d_50_v20_: _dafny.Map
                d_50_v20_ = _dafny.Map({p1: d_49_v19_})
                d_51_v21_: _dafny.Map
                d_51_v21_ = _dafny.Map({d_18_v0_: (d_50_v20_) | (d_50_v20_)})
                d_51_v21_ = (d_51_v21_) | (_dafny.Map({d_47_cf48_: d_50_v20_}))
                d_52_v22_: int
                d_52_v22_ = 118
                d_52_v22_ = p1
                d_53_v23_: _dafny.MultiSet
                d_53_v23_ = _dafny.MultiSet([d_52_v22_, (d_49_v19_).f0, d_52_v22_, (d_49_v19_).f0])
                d_54_v24_: _dafny.MultiSet
                d_54_v24_ = _dafny.MultiSet([(d_53_v23_) - ((d_53_v23_).set(p1, default__.abs(p1))), d_53_v23_, (d_53_v23_).intersection(d_53_v23_)])
                d_55_v25_: _dafny.Map
                d_55_v25_ = _dafny.Map({p1: (0) - (p1)})
                d_56_v26_: _dafny.Map
                d_56_v26_ = _dafny.Map({d_49_v19_: p1})
                d_57_v27_: _dafny.Seq
                d_57_v27_ = _dafny.SeqWithoutIsStrInference([p1])
                d_58_v28_: _dafny.Map
                d_58_v28_ = _dafny.Map({len(d_57_v27_): _dafny.SeqWithoutIsStrInference([d_18_v0_, d_47_cf48_])})
                d_54_v24_ = default__.fm32((d_55_v25_) | (d_55_v25_), len(d_56_v26_), False, d_58_v28_, globalState)
                d_59_v29_: _dafny.Seq
                d_59_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wkht"))
                d_60_v30_: D2
                d_60_v30_ = D2_DC4(d_59_v29_, d_48_cf47_)
                d_61_v31_: _dafny.Map
                d_61_v31_ = _dafny.Map({d_60_v30_: p1})
                d_62_v32_: D18
                d_62_v32_ = D18_DC39(p0)
                pat_let_tv0_ = d_48_cf47_
                d_63_v33_: _dafny.Map
                def iife10_(_pat_let0_0):
                    def iife11_(d_64_dt__update__tmp_h0_):
                        def iife12_(_pat_let1_0):
                            def iife13_(d_65_dt__update_hcf59_h0_):
                                return D18_DC39(d_65_dt__update_hcf59_h0_)
                            return iife13_(_pat_let1_0)
                        return iife12_(pat_let_tv0_)
                    return iife11_(_pat_let0_0)
                d_63_v33_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhqkkmfj"))): iife10_(d_62_v32_)})
                d_66_v35_: _dafny.Seq
                def iife14_():
                    coll10_ = _dafny.Map()
                    compr_10_: int
                    for compr_10_ in _dafny.IntegerRange(648, 747):
                        d_67_v34_: int = compr_10_
                        if ((648) <= (d_67_v34_)) and ((d_67_v34_) < (747)):
                            coll10_[(d_67_v34_) * ((_dafny.MultiSet([d_59_v29_])).cardinality)] = d_62_v32_
                    return _dafny.Map(coll10_)
                d_66_v35_ = _dafny.SeqWithoutIsStrInference([default__.fm33((d_49_v19_).f0, _dafny.Map({D2_DC4(d_59_v29_, True): d_52_v22_}), globalState), default__.fm33(d_52_v22_, d_61_v31_, globalState), (d_63_v33_).set((0) - ((d_49_v19_).f0), d_62_v32_), default__.fm33((d_49_v19_).f0, d_61_v31_, globalState), iife14_()
                ])
                d_52_v22_ = len(d_66_v35_)
                d_68_v36_: _dafny.Array
                nw6_ = _dafny.Array(None, 12)
                d_68_v36_ = nw6_
                rhs9_ = p0
                rhs10_ = d_68_v36_
                d_48_cf47_ = rhs9_
                d_68_v36_ = rhs10_
            elif True:
                d_69_v37_: int
                d_69_v37_ = 223
                d_69_v37_ = d_69_v37_
                d_70_v38_: _dafny.Map
                d_70_v38_ = _dafny.Map({d_48_cf47_: default__.fm2(p0, p1, p0, d_48_cf47_, globalState)})
                d_71_v39_: _dafny.Map
                d_71_v39_ = d_70_v38_
                d_72_v40_: _dafny.Map
                d_72_v40_ = _dafny.Map({d_71_v39_: d_48_cf47_})
                d_72_v40_ = (d_72_v40_).set(d_71_v39_, p0)
                d_73_v41_: _dafny.Array
                nw7_ = _dafny.Array(int(0), 1)
                d_73_v41_ = nw7_
                index0_ = default__.safeIndex(367, (d_73_v41_).length(0))
                (d_73_v41_)[index0_] = d_69_v37_
                index1_ = default__.safeIndex(367, (d_73_v41_).length(0))
                (d_73_v41_)[index1_] = (default__.fm22(globalState)).cardinality
                d_48_cf47_ = p0
                d_74_v42_: _dafny.Seq
                d_74_v42_ = _dafny.SeqWithoutIsStrInference([(d_73_v41_)[default__.safeIndex(367, (d_73_v41_).length(0))]])
                index2_ = default__.safeIndex(367, (d_73_v41_).length(0))
                (d_73_v41_)[index2_] = default__.safeModuloInt(len((d_74_v42_) + (d_74_v42_)), default__.safeModuloInt(d_69_v37_, p1))
            d_75_v43_: int
            d_75_v43_ = -188
            d_76_v44_: _dafny.Set
            d_76_v44_ = _dafny.Set({not(d_48_cf47_)})
            d_77_v45_: _dafny.Array
            def lambda0_(d_78_cf47_):
                def lambda1_(d_79_i0_):
                    return d_78_cf47_

                return lambda1_

            init0_ = lambda0_(d_48_cf47_)
            nw8_ = _dafny.Array(None, 24)
            for i0_0_ in range(nw8_.length(0)):
                nw8_[i0_0_] = init0_(i0_0_)
            d_77_v45_ = nw8_
            d_80_v46_: D11
            d_80_v46_ = D11_DC24(d_77_v45_, d_48_cf47_, p0)
            d_81_v47_: _dafny.Set
            def iife15_(_pat_let2_0):
                def iife16_(d_82_dt__update__tmp_h1_):
                    def iife17_(_pat_let3_0):
                        def iife18_(d_83_dt__update_hcf39_h0_):
                            return D11_DC24((d_82_dt__update__tmp_h1_).cf38, d_83_dt__update_hcf39_h0_, (d_82_dt__update__tmp_h1_).cf40)
                        return iife18_(_pat_let3_0)
                    return iife17_(False)
                return iife16_(_pat_let2_0)
            d_81_v47_ = _dafny.Set({d_80_v46_, iife15_(d_80_v46_)})
            d_84_v48_: D18
            d_84_v48_ = D18_DC38(d_81_v47_)
            d_85_v49_: _dafny.MultiSet
            d_85_v49_ = _dafny.MultiSet([D18_DC38(_dafny.Set({d_80_v46_})), d_84_v48_])
            d_86_v50_: D19
            d_86_v50_ = D19_DC41(d_85_v49_)
            d_75_v43_ = default__.fm3(d_76_v44_, d_47_cf48_, (((d_86_v50_).cf62).cardinality if d_48_cf47_ else d_75_v43_), d_75_v43_, globalState)
            if not(not((d_75_v43_) == (182))):
                d_75_v43_ = (641) - (p1)
                d_77_v45_ = d_77_v45_
                d_87_v51_: _dafny.Seq
                d_87_v51_ = _dafny.SeqWithoutIsStrInference([p0])
                d_48_cf47_ = default__.fm2((d_76_v44_).isdisjoint(d_76_v44_), default__.safeDivisionInt(d_75_v43_, (0) - (len(d_87_v51_))), p0, True, globalState)
                pat_let_tv1_ = p0
                def iife19_(_pat_let4_0):
                    def iife20_(d_88_dt__update__tmp_h2_):
                        def iife21_(_pat_let5_0):
                            def iife22_(d_89_dt__update_hcf47_h0_):
                                return D13_DC28(d_89_dt__update_hcf47_h0_, (d_88_dt__update__tmp_h2_).cf48)
                            return iife22_(_pat_let5_0)
                        return iife21_(pat_let_tv1_)
                    return iife20_(_pat_let4_0)
                d_19_v1_ = (iife19_(d_19_v1_) if d_48_cf47_ else d_19_v1_)
                d_90_v52_: _dafny.Set
                d_90_v52_ = _dafny.Set({p1})
                d_91_v53_: C1
                nw9_ = C1()
                nw9_.ctor__(p1)
                d_91_v53_ = nw9_
                d_92_v54_: _dafny.MultiSet
                d_92_v54_ = _dafny.MultiSet([d_91_v53_])
                d_93_v55_: _dafny.Seq
                d_93_v55_ = _dafny.SeqWithoutIsStrInference([d_90_v52_, _dafny.Set({default__.fm3(d_76_v44_, _dafny.CodePoint('d'), 872, (d_92_v54_).cardinality, globalState)})])
                d_94_v56_: _dafny.Seq
                d_94_v56_ = _dafny.SeqWithoutIsStrInference([(d_93_v55_)[default__.safeIndex(p1, len(d_93_v55_))]])
                d_90_v52_ = ((d_93_v55_)[default__.safeIndex(536, len(d_93_v55_))]).intersection(d_90_v52_)
            elif True:
                d_95_v57_: C1
                nw10_ = C1()
                nw10_.ctor__((p1) * (p1))
                d_95_v57_ = nw10_
                d_96_v58_: _dafny.Seq
                d_96_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmdqlx"))
                d_97_v59_: D2
                d_97_v59_ = D2_DC4(d_96_v58_, p0)
                d_98_v60_: _dafny.Map
                d_98_v60_ = _dafny.Map({(d_97_v59_).cf4: p0})
                d_99_v61_: _dafny.Map
                d_99_v61_ = _dafny.Map({d_48_cf47_: (d_95_v57_).fm15(d_98_v60_, d_48_cf47_, d_48_cf47_, globalState)})
                rhs11_ = d_77_v45_
                rhs12_ = (d_96_v58_) < (d_96_v58_)
                rhs13_ = d_95_v57_
                rhs14_ = ((((d_99_v61_)[p0] if (p0) in (d_99_v61_) else d_75_v43_)) * (d_75_v43_) if p0 else -427)
                d_77_v45_ = rhs11_
                r0 = rhs12_
                d_95_v57_ = rhs13_
                d_75_v43_ = rhs14_
                d_100_v62_: T1
                nw11_ = C5()
                nw11_.ctor__((d_96_v58_)[default__.safeIndex(p1, len(d_96_v58_))], p1, default__.safeModuloInt(d_75_v43_, d_75_v43_))
                d_100_v62_ = nw11_
                d_100_v62_ = d_100_v62_
                d_101_v63_: C1
                nw12_ = C1()
                nw12_.ctor__((d_100_v62_).f0)
                d_101_v63_ = nw12_
                d_102_v64_: _dafny.Set
                d_102_v64_ = _dafny.Set({(0) - (d_75_v43_), p1, p1, 308})
                d_103_v65_: _dafny.MultiSet
                d_103_v65_ = _dafny.MultiSet([((d_99_v61_)[d_48_cf47_] if (d_48_cf47_) in (d_99_v61_) else (d_100_v62_).f0)])
                d_104_v66_: _dafny.Array
                nw13_ = _dafny.Array(None, 15)
                nw13_[int(0)] = p1
                nw13_[int(1)] = d_75_v43_
                nw13_[int(2)] = len(d_102_v64_)
                nw13_[int(3)] = 132
                nw13_[int(4)] = ((d_100_v62_).f0) + ((d_100_v62_).f0)
                nw13_[int(5)] = p1
                nw13_[int(6)] = 854
                nw13_[int(7)] = 842
                nw13_[int(8)] = p1
                nw13_[int(9)] = p1
                nw13_[int(10)] = 686
                nw13_[int(11)] = (d_100_v62_).f0
                nw13_[int(12)] = p1
                nw13_[int(13)] = (d_103_v65_).cardinality
                nw13_[int(14)] = p1
                d_104_v66_ = nw13_
                d_105_v67_: _dafny.Map
                d_105_v67_ = _dafny.Map({667: p0})
                d_106_v68_: _dafny.Map
                d_106_v68_ = _dafny.Map({default__.safeDivisionInt(d_75_v43_, len(d_105_v67_)): d_48_cf47_})
                rhs15_ = p1
                rhs16_ = ((d_106_v68_)[p1] if (p1) in (d_106_v68_) else d_48_cf47_)
                rhs17_ = (p1) > (d_75_v43_)
                rhs18_ = d_104_v66_
                d_75_v43_ = rhs15_
                r0 = rhs16_
                d_48_cf47_ = rhs17_
                d_104_v66_ = rhs18_
                d_107_v69_: _dafny.Seq
                d_107_v69_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
                d_108_v70_: _dafny.Map
                d_108_v70_ = _dafny.Map({d_107_v69_: d_48_cf47_})
                d_108_v70_ = _dafny.Map({(d_107_v69_) + (d_107_v69_): p0})
            index3_ = default__.safeIndex(627, (d_77_v45_).length(0))
            (d_77_v45_)[index3_] = p0
            d_109_v71_: _dafny.MultiSet
            d_109_v71_ = _dafny.MultiSet([d_48_cf47_, p0, d_48_cf47_])
            index4_ = default__.safeIndex(627, (d_77_v45_).length(0))
            (d_77_v45_)[index4_] = (((d_109_v71_)[False] if (False) in (d_109_v71_) else d_75_v43_)) < (default__.fm3(_dafny.Set({d_48_cf47_, d_48_cf47_, d_48_cf47_}), d_18_v0_, p1, p1, globalState))
        elif source0_.is_DC29:
            d_110___mcc_h6_ = source0_.cf49
            d_111___mcc_h7_ = source0_.cf50
            d_112_cf50_ = d_111___mcc_h7_
            d_113_cf49_ = d_110___mcc_h6_
            d_114_v72_: _dafny.Array
            nw14_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_114_v72_ = nw14_
            nw15_ = _dafny.Array(_dafny.Array(None, 0), 13)
            d_114_v72_ = nw15_
            d_115_v73_: _dafny.Set
            d_115_v73_ = _dafny.Set({d_112_cf50_, p0, p0, not(d_112_cf50_), d_112_cf50_})
            d_116_v74_: T1
            nw16_ = C5()
            nw16_.ctor__(_dafny.CodePoint('c'), default__.fm3(d_115_v73_, d_18_v0_, p1, p1, globalState), (0) - ((0) - (default__.safeModuloInt(p1, p1))))
            d_116_v74_ = nw16_
            d_116_v74_ = (d_116_v74_ if d_112_cf50_ else d_116_v74_)
            d_117_v75_: _dafny.Array
            nw17_ = _dafny.Array(_dafny.Map({}), 19)
            d_117_v75_ = nw17_
            d_118_v76_: _dafny.Map
            d_118_v76_ = _dafny.Map({p1: (d_116_v74_).f0})
            index5_ = default__.safeIndex(999, (d_117_v75_).length(0))
            (d_117_v75_)[index5_] = (d_118_v76_) | (d_118_v76_)
            index6_ = default__.safeIndex(999, (d_117_v75_).length(0))
            (d_117_v75_)[index6_] = (((d_118_v76_).set(p1, (d_116_v74_).f0)) | (d_118_v76_)) | ((d_118_v76_) | (d_118_v76_))
            d_119_v77_: _dafny.Array
            def lambda2_(d_120_cf50_):
                def lambda3_(d_121_i1_):
                    return d_120_cf50_

                return lambda3_

            init1_ = lambda2_(d_112_cf50_)
            nw18_ = _dafny.Array(None, 23)
            for i0_1_ in range(nw18_.length(0)):
                nw18_[i0_1_] = init1_(i0_1_)
            d_119_v77_ = nw18_
            d_119_v77_ = d_119_v77_
        elif source0_.is_DC26:
            d_122___mcc_h8_ = source0_.cf42
            d_123_cf42_ = d_122___mcc_h8_
            d_124_v78_: C1
            nw19_ = C1()
            nw19_.ctor__(p1)
            d_124_v78_ = nw19_
            d_125_v79_: _dafny.Seq
            d_125_v79_ = _dafny.SeqWithoutIsStrInference([d_124_v78_])
            d_126_v80_: _dafny.Seq
            d_126_v80_ = _dafny.SeqWithoutIsStrInference([d_125_v79_])
            d_125_v79_ = (d_126_v80_)[default__.safeIndex(p1, len(d_126_v80_))]
            d_127_v81_: int
            d_127_v81_ = 70
            d_128_v82_: _dafny.Set
            d_128_v82_ = _dafny.Set({False, p0, p0})
            d_127_v81_ = (d_127_v81_) - (default__.fm3(d_128_v82_, d_18_v0_, d_127_v81_, p1, globalState))
            r0 = p0
            d_127_v81_ = d_127_v81_
        elif True:
            d_129___mcc_h9_ = source0_.cf51
            d_130_cf51_ = d_129___mcc_h9_
            r0 = not(p0)
            d_131_v83_: _dafny.Array
            def lambda4_(d_132_p0_):
                def lambda5_(d_133_i2_):
                    return not(d_132_p0_)

                return lambda5_

            init2_ = lambda4_(p0)
            nw20_ = _dafny.Array(None, 15)
            for i0_2_ in range(nw20_.length(0)):
                nw20_[i0_2_] = init2_(i0_2_)
            d_131_v83_ = nw20_
            d_134_v84_: _dafny.MultiSet
            d_134_v84_ = _dafny.MultiSet([p1])
            index7_ = default__.safeIndex(64, (d_131_v83_).length(0))
            (d_131_v83_)[index7_] = (d_134_v84_).issubset(d_134_v84_)
            index8_ = default__.safeIndex(64, (d_131_v83_).length(0))
            (d_131_v83_)[index8_] = p0
            d_135_v85_: _dafny.Array
            nw21_ = _dafny.Array(_dafny.Map({}), 24)
            d_135_v85_ = nw21_
            d_136_v86_: _dafny.Set
            d_136_v86_ = _dafny.Set({p0})
            d_137_v87_: C1
            nw22_ = C1()
            nw22_.ctor__(p1)
            d_137_v87_ = nw22_
            d_138_v88_: _dafny.Seq
            d_138_v88_ = _dafny.SeqWithoutIsStrInference([d_137_v87_, d_137_v87_])
            d_139_v89_: _dafny.Seq
            d_139_v89_ = _dafny.SeqWithoutIsStrInference([d_131_v83_])
            d_140_v90_: _dafny.Array
            nw23_ = _dafny.Array(None, 17)
            nw23_[int(0)] = (0) - (p1)
            nw23_[int(1)] = p1
            nw23_[int(2)] = p1
            nw23_[int(3)] = p1
            nw23_[int(4)] = p1
            nw23_[int(5)] = p1
            nw23_[int(6)] = p1
            nw23_[int(7)] = 122
            nw23_[int(8)] = p1
            nw23_[int(9)] = p1
            nw23_[int(10)] = p1
            nw23_[int(11)] = len(d_136_v86_)
            nw23_[int(12)] = len(d_138_v88_)
            nw23_[int(13)] = p1
            nw23_[int(14)] = p1
            nw23_[int(15)] = p1
            nw23_[int(16)] = len(d_139_v89_)
            d_140_v90_ = nw23_
            d_141_v91_: _dafny.Map
            d_141_v91_ = _dafny.Map({(d_131_v83_)[default__.safeIndex(64, (d_131_v83_).length(0))]: d_140_v90_})
            index9_ = default__.safeIndex(963, (d_135_v85_).length(0))
            (d_135_v85_)[index9_] = d_141_v91_
            index10_ = default__.safeIndex(963, (d_135_v85_).length(0))
            (d_135_v85_)[index10_] = d_141_v91_
            d_142_v92_: int
            d_142_v92_ = 601
            d_143_v93_: _dafny.Seq
            d_143_v93_ = _dafny.SeqWithoutIsStrInference([p1, d_142_v92_])
            d_144_v94_: _dafny.Seq
            d_144_v94_ = _dafny.SeqWithoutIsStrInference([d_134_v84_])
            d_142_v92_ = ((_dafny.MultiSet(d_143_v93_) if (d_131_v83_)[default__.safeIndex(64, (d_131_v83_).length(0))] else (d_144_v94_)[default__.safeIndex(d_142_v92_, len(d_144_v94_))])).cardinality
        d_145_v95_: int
        d_145_v95_ = 898
        rhs19_ = p1
        rhs20_ = p0
        d_145_v95_ = rhs19_
        r0 = rhs20_
        d_146_v96_: _dafny.Seq
        d_146_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwej"))
        d_147_v97_: _dafny.Seq
        d_147_v97_ = _dafny.SeqWithoutIsStrInference([d_145_v95_])
        d_148_v98_: _dafny.Seq
        d_148_v98_ = _dafny.SeqWithoutIsStrInference([p0, default__.fm2(p0, len(d_147_v97_), p0, p0, globalState)])
        d_149_v99_: T1
        nw24_ = C5()
        nw24_.ctor__(_dafny.CodePoint('w'), len(d_148_v98_), d_145_v95_)
        d_149_v99_ = nw24_
        d_150_v100_: _dafny.Seq
        d_150_v100_ = _dafny.SeqWithoutIsStrInference([d_149_v99_])
        d_151_v101_: _dafny.Set
        d_151_v101_ = _dafny.Set({len(_dafny.Map({d_146_v96_: not(p0)})), p1, p1, len(d_150_v100_)})
        if (d_151_v101_).ispropersubset(default__.fm34(globalState)):
            if p0:
                d_152_v102_: _dafny.Set
                d_152_v102_ = d_151_v101_
                d_153_v103_: _dafny.Seq
                d_153_v103_ = _dafny.SeqWithoutIsStrInference([d_152_v102_, d_151_v101_, d_151_v101_])
                d_154_v104_: _dafny.Map
                d_154_v104_ = _dafny.Map({p0: p1})
                d_153_v103_ = (d_153_v103_) + (default__.fm35(((d_154_v104_)[p0] if (p0) in (d_154_v104_) else len(d_146_v96_)), (d_149_v99_).f0, p0, globalState))
                r0 = p0
                d_155_v105_: _dafny.Array
                nw25_ = _dafny.Array(int(0), 4)
                d_155_v105_ = nw25_
                d_156_v106_: C0
                nw26_ = C0()
                nw26_.ctor__(d_18_v0_)
                d_156_v106_ = nw26_
                d_157_v107_: _dafny.Map
                d_157_v107_ = _dafny.Map({d_156_v106_: _dafny.SeqWithoutIsStrInference([(d_149_v99_).f0 for d_158_i3_ in range(default__.abs(289))])})
                d_159_v108_: _dafny.Map
                d_159_v108_ = _dafny.Map({d_157_v107_: d_145_v95_})
                index11_ = default__.safeIndex(31, (d_155_v105_).length(0))
                (d_155_v105_)[index11_] = ((d_159_v108_)[d_157_v107_] if (d_157_v107_) in (d_159_v108_) else 376)
                d_160_v109_: _dafny.Array
                def lambda6_(d_161_p0_):
                    def lambda7_(d_162_i4_):
                        return d_161_p0_

                    return lambda7_

                init3_ = lambda6_(p0)
                nw27_ = _dafny.Array(None, 22)
                for i0_3_ in range(nw27_.length(0)):
                    nw27_[i0_3_] = init3_(i0_3_)
                d_160_v109_ = nw27_
                d_163_v110_: _dafny.Set
                d_163_v110_ = _dafny.Set({d_160_v109_})
                index12_ = default__.safeIndex(31, (d_155_v105_).length(0))
                (d_155_v105_)[index12_] = ((d_149_v99_).f0) - (len(d_163_v110_))
                d_164_v111_: T0
                nw28_ = C3()
                nw28_.ctor__((d_149_v99_).f0)
                d_164_v111_ = nw28_
                d_165_v112_: _dafny.Map
                d_165_v112_ = _dafny.Map({(d_155_v105_)[default__.safeIndex(31, (d_155_v105_).length(0))]: _dafny.Set({d_164_v111_})})
                d_145_v95_ = default__.safeModuloInt(len(d_165_v112_), 223)
                index13_ = default__.safeIndex(31, (d_155_v105_).length(0))
                (d_155_v105_)[index13_] = (d_149_v99_).f0
            elif True:
                d_149_v99_ = d_149_v99_
                d_166_v113_: _dafny.Set
                d_166_v113_ = _dafny.Set({d_149_v99_})
                rhs21_ = d_146_v96_
                rhs22_ = (p0 if False else p0)
                rhs23_ = (len((d_166_v113_) | (d_166_v113_))) - ((d_149_v99_).f0)
                d_146_v96_ = rhs21_
                r0 = rhs22_
                d_145_v95_ = rhs23_
                d_167_v114_: _dafny.Array
                nw29_ = _dafny.Array(None, 3)
                d_167_v114_ = nw29_
                index14_ = default__.safeIndex(345, (d_167_v114_).length(0))
                (d_167_v114_)[index14_] = d_149_v99_
                index15_ = default__.safeIndex(345, (d_167_v114_).length(0))
                (d_167_v114_)[index15_] = d_149_v99_
                d_145_v95_ = (0) - (((d_149_v99_).f0) * (((d_149_v99_).f0) - ((0) - ((d_149_v99_).f0))))
                d_145_v95_ = (-90) - ((d_149_v99_).f0)
            d_168_v115_: _dafny.Map
            d_168_v115_ = _dafny.Map({p0: p0})
            d_169_v116_: _dafny.Seq
            d_169_v116_ = _dafny.SeqWithoutIsStrInference([d_168_v115_, d_168_v115_, _dafny.Map({p0: p0}), (d_168_v115_) | (d_168_v115_)])
            d_169_v116_ = d_169_v116_
            d_170_v117_: C3
            nw30_ = C3()
            nw30_.ctor__(len(d_148_v98_))
            d_170_v117_ = nw30_
            d_171_v118_: _dafny.Set
            d_171_v118_ = _dafny.Set({p0, p0})
            def iife23_():
                coll11_ = _dafny.Set()
                compr_11_: int
                for compr_11_ in _dafny.IntegerRange(766, 816):
                    d_172_v119_: int = compr_11_
                    if ((766) <= (d_172_v119_)) and ((d_172_v119_) < (816)):
                        coll11_ = coll11_.union(_dafny.Set([default__.safeModuloInt(d_172_v119_, 918)]))
                return _dafny.Set(coll11_)
            d_145_v95_ = (720) - ((default__.fm3(d_171_v118_, d_18_v0_, len(iife23_()
            ), (d_149_v99_).f0, globalState)) - ((d_149_v99_).f0))
            d_173_v120_: C0
            nw31_ = C0()
            nw31_.ctor__(_dafny.CodePoint('m'))
            d_173_v120_ = nw31_
        elif True:
            d_174_v121_: _dafny.Array
            nw32_ = _dafny.Array(None, 16)
            d_174_v121_ = nw32_
            d_175_v122_: D3
            d_175_v122_ = D3_DC9(745, d_174_v121_, p0)
            d_176_v123_: _dafny.Map
            d_176_v123_ = _dafny.Map({p0: False})
            d_177_v124_: _dafny.Array
            nw33_ = _dafny.Array(None, 26)
            nw33_[int(0)] = p0
            nw33_[int(1)] = p0
            nw33_[int(2)] = (d_146_v96_) <= (d_146_v96_)
            nw33_[int(3)] = False
            nw33_[int(4)] = p0
            nw33_[int(5)] = p0
            nw33_[int(6)] = p0
            nw33_[int(7)] = p0
            nw33_[int(8)] = p0
            nw33_[int(9)] = p0
            nw33_[int(10)] = p0
            nw33_[int(11)] = (p0 if p0 else True)
            nw33_[int(12)] = p0
            nw33_[int(13)] = (d_18_v0_) not in (d_146_v96_)
            nw33_[int(14)] = not (not(p0)) or (p0)
            nw33_[int(15)] = p0
            nw33_[int(16)] = p0
            nw33_[int(17)] = p0
            nw33_[int(18)] = True
            nw33_[int(19)] = (d_175_v122_).cf12
            nw33_[int(20)] = (True if p0 else ((d_176_v123_)[p0] if (p0) in (d_176_v123_) else p0))
            nw33_[int(21)] = p0
            nw33_[int(22)] = (p0 if p0 else p0)
            nw33_[int(23)] = p0
            nw33_[int(24)] = p0
            nw33_[int(25)] = not(p0)
            d_177_v124_ = nw33_
            d_177_v124_ = d_177_v124_
            if p0:
                d_178_v125_: _dafny.Map
                d_178_v125_ = _dafny.Map({d_18_v0_: d_145_v95_})
                d_179_v126_: _dafny.MultiSet
                d_179_v126_ = _dafny.MultiSet([d_178_v125_])
                d_179_v126_ = d_179_v126_
                d_180_v127_: C4
                nw34_ = C4()
                nw34_.ctor__(len(d_148_v98_))
                d_180_v127_ = nw34_
                d_181_v128_: _dafny.Seq
                d_181_v128_ = _dafny.SeqWithoutIsStrInference([d_180_v127_])
                d_182_v129_: _dafny.Seq
                d_182_v129_ = _dafny.SeqWithoutIsStrInference([d_181_v128_])
                r0 = (((d_182_v129_)[default__.safeIndex(p1, len(d_182_v129_))] if p0 else _dafny.SeqWithoutIsStrInference([d_180_v127_]))) < ((d_181_v128_).set(default__.safeIndex((d_149_v99_).f0, len(d_181_v128_)), d_180_v127_))
                d_183_v130_: _dafny.Array
                nw35_ = _dafny.Array(int(0), 13)
                d_183_v130_ = nw35_
                d_184_v131_: T0
                nw36_ = C4()
                nw36_.ctor__((d_149_v99_).f0)
                d_184_v131_ = nw36_
                d_185_v132_: D10
                d_185_v132_ = D10_DC21(d_184_v131_)
                d_186_v133_: _dafny.Map
                d_186_v133_ = _dafny.Map({d_180_v127_: (d_185_v132_).cf33})
                d_187_v134_: _dafny.Array
                nw37_ = _dafny.Array(None, 18)
                nw37_[int(0)] = (d_184_v131_ if p0 else d_184_v131_)
                nw37_[int(1)] = d_184_v131_
                nw37_[int(2)] = d_184_v131_
                nw37_[int(3)] = d_184_v131_
                nw37_[int(4)] = d_184_v131_
                nw37_[int(5)] = d_184_v131_
                nw37_[int(6)] = d_184_v131_
                nw37_[int(7)] = d_184_v131_
                nw37_[int(8)] = ((d_186_v133_)[d_180_v127_] if (d_180_v127_) in (d_186_v133_) else d_184_v131_)
                nw37_[int(9)] = (d_184_v131_ if p0 else d_184_v131_)
                nw37_[int(10)] = d_184_v131_
                nw37_[int(11)] = d_184_v131_
                nw37_[int(12)] = d_184_v131_
                nw37_[int(13)] = d_184_v131_
                nw37_[int(14)] = d_184_v131_
                nw37_[int(15)] = d_184_v131_
                nw37_[int(16)] = d_184_v131_
                nw37_[int(17)] = d_184_v131_
                d_187_v134_ = nw37_
                index16_ = default__.safeIndex(587, (d_187_v134_).length(0))
                (d_187_v134_)[index16_] = d_184_v131_
                d_188_v135_: _dafny.Map
                d_188_v135_ = _dafny.Map({d_174_v121_: p0})
                index17_ = default__.safeIndex(587, (d_187_v134_).length(0))
                rhs24_ = d_183_v130_
                rhs25_ = default__.fm17(len((d_188_v135_) | (d_188_v135_)), globalState)
                rhs26_ = (d_149_v99_).f0
                rhs27_ = d_184_v131_
                lhs0_ = d_187_v134_
                lhs1_ = default__.safeIndex(587, (d_187_v134_).length(0))
                d_183_v130_ = rhs24_
                d_18_v0_ = rhs25_
                d_145_v95_ = rhs26_
                lhs0_[lhs1_] = rhs27_
                index18_ = default__.safeIndex(550, (d_177_v124_).length(0))
                (d_177_v124_)[index18_] = not(not(p0))
                d_189_v136_: _dafny.Map
                d_189_v136_ = _dafny.Map({d_149_v99_: (d_149_v99_).f0})
                index19_ = default__.safeIndex(550, (d_177_v124_).length(0))
                (d_177_v124_)[index19_] = (d_189_v136_) == (d_189_v136_)
                r0 = not(not((p0 if p0 else p0)))
            elif True:
                d_18_v0_ = d_18_v0_
                d_190_v137_: _dafny.Set
                d_190_v137_ = _dafny.Set({(d_149_v99_).fm6(p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqrcjrmk")), d_146_v96_, (d_149_v99_).f0, globalState), p0, p0, p0, p0})
                d_145_v95_ = default__.fm3(d_190_v137_, d_18_v0_, default__.safeModuloInt(p1, (d_149_v99_).f0), p1, globalState)
                r0 = p0
                d_145_v95_ = (d_149_v99_).f0
                d_191_v139_: _dafny.Array
                def lambda8_(d_192_v99_):
                    def lambda9_(d_193_i5_):
                        def iife24_():
                            coll12_ = _dafny.Set()
                            compr_12_: int
                            for compr_12_ in (_dafny.Map({(d_192_v99_).f0: True})).keys.Elements:
                                d_194_v138_: int = compr_12_
                                if (d_194_v138_) in (_dafny.Map({(d_192_v99_).f0: True})):
                                    coll12_ = coll12_.union(_dafny.Set([(d_194_v138_) - (len(_dafny.Map({True: True})))]))
                            return _dafny.Set(coll12_)
                        return iife24_()
                        

                    return lambda9_

                init4_ = lambda8_(d_149_v99_)
                nw38_ = _dafny.Array(None, 21)
                for i0_4_ in range(nw38_.length(0)):
                    nw38_[i0_4_] = init4_(i0_4_)
                d_191_v139_ = nw38_
                index20_ = default__.safeIndex(338, (d_191_v139_).length(0))
                (d_191_v139_)[index20_] = d_151_v101_
                d_195_v140_: D16
                d_195_v140_ = D16_DC35((d_149_v99_).f0)
                index21_ = default__.safeIndex(338, (d_191_v139_).length(0))
                rhs28_ = default__.fm34(globalState)
                rhs29_ = (_dafny.SeqWithoutIsStrInference([d_18_v0_ for d_196_i6_ in range(default__.abs(890))])).set(default__.safeIndex(323, len(_dafny.SeqWithoutIsStrInference([d_18_v0_ for d_197_i6_ in range(default__.abs(890))]))), (d_146_v96_)[default__.safeIndex((0) - ((d_195_v140_).cf55), len(d_146_v96_))])
                lhs2_ = d_191_v139_
                lhs3_ = default__.safeIndex(338, (d_191_v139_).length(0))
                lhs2_[lhs3_] = rhs28_
                d_146_v96_ = rhs29_
            d_198_v141_: D3
            d_198_v141_ = D3_DC8(p0, (d_149_v99_).f0)
            d_199_v142_: T0
            nw39_ = C1()
            nw39_.ctor__((d_149_v99_).f0)
            d_199_v142_ = nw39_
            d_200_v143_: _dafny.Seq
            d_200_v143_ = _dafny.SeqWithoutIsStrInference([d_199_v142_])
            d_201_v144_: _dafny.MultiSet
            d_201_v144_ = _dafny.MultiSet([d_199_v142_, d_199_v142_, (d_200_v143_)[default__.safeIndex(d_145_v95_, len(d_200_v143_))]])
            d_202_v145_: _dafny.Map
            d_202_v145_ = _dafny.Map({d_198_v141_: not((_dafny.MultiSet([d_199_v142_, d_199_v142_, d_199_v142_])).issubset(d_201_v144_))})
            d_202_v145_ = (d_202_v145_).set(d_198_v141_, p0)
            d_203_v146_: C4
            nw40_ = C4()
            nw40_.ctor__(((d_149_v99_).f0) * ((d_149_v99_).f0))
            d_203_v146_ = nw40_
            d_204_v147_: C2
            nw41_ = C2()
            nw41_.ctor__(d_145_v95_)
            d_204_v147_ = nw41_
        d_205_v148_: C2
        nw42_ = C2()
        nw42_.ctor__(len(d_146_v96_))
        d_205_v148_ = nw42_
        d_206_v149_: _dafny.Map
        d_206_v149_ = _dafny.Map({(d_145_v95_) != (p1): d_205_v148_})
        d_207_v150_: D6
        d_207_v150_ = D6_DC15((0) - (d_145_v95_), p0, len(_dafny.Map({p1: (d_149_v99_).f0})))
        d_208_v151_: _dafny.Set
        d_208_v151_ = _dafny.Set({p0})
        d_209_v152_: D5
        d_209_v152_ = D5_DC13((d_149_v99_).f0, d_208_v151_, d_18_v0_, d_145_v95_, p0)
        d_210_v153_: _dafny.Seq
        d_210_v153_ = _dafny.SeqWithoutIsStrInference([(d_207_v150_ if p0 else d_207_v150_), D6_DC15((d_209_v152_).cf20, p0, 898), d_207_v150_, d_207_v150_, d_207_v150_])
        d_211_v154_: _dafny.Array
        def lambda10_(d_212_v95_):
            def lambda11_(d_213_i7_):
                return (d_213_i7_) * (d_212_v95_)

            return lambda11_

        init5_ = lambda10_(d_145_v95_)
        nw43_ = _dafny.Array(None, 29)
        for i0_5_ in range(nw43_.length(0)):
            nw43_[i0_5_] = init5_(i0_5_)
        d_211_v154_ = nw43_
        d_214_v155_: _dafny.Seq
        d_214_v155_ = _dafny.SeqWithoutIsStrInference([d_211_v154_, d_211_v154_])
        d_215_v156_: _dafny.Map
        d_215_v156_ = _dafny.Map({p1: d_211_v154_})
        d_216_v157_: _dafny.Array
        nw44_ = _dafny.Array(None, 23)
        nw44_[int(0)] = (d_214_v155_)[default__.safeIndex((d_147_v97_)[default__.safeIndex((d_149_v99_).f0, len(d_147_v97_))], len(d_214_v155_))]
        nw44_[int(1)] = (d_211_v154_ if p0 else d_211_v154_)
        nw44_[int(2)] = d_211_v154_
        nw44_[int(3)] = d_211_v154_
        nw44_[int(4)] = d_211_v154_
        nw44_[int(5)] = d_211_v154_
        nw44_[int(6)] = d_211_v154_
        nw44_[int(7)] = d_211_v154_
        nw44_[int(8)] = d_211_v154_
        nw44_[int(9)] = (d_214_v155_)[default__.safeIndex(len(d_208_v151_), len(d_214_v155_))]
        nw44_[int(10)] = (d_211_v154_ if p0 else d_211_v154_)
        nw44_[int(11)] = d_211_v154_
        nw44_[int(12)] = d_211_v154_
        nw44_[int(13)] = ((d_215_v156_)[(0) - (len(d_147_v97_))] if ((0) - (len(d_147_v97_))) in (d_215_v156_) else d_211_v154_)
        nw44_[int(14)] = d_211_v154_
        nw44_[int(15)] = d_211_v154_
        nw44_[int(16)] = d_211_v154_
        nw44_[int(17)] = d_211_v154_
        nw44_[int(18)] = d_211_v154_
        nw44_[int(19)] = d_211_v154_
        nw44_[int(20)] = d_211_v154_
        nw44_[int(21)] = d_211_v154_
        nw44_[int(22)] = d_211_v154_
        d_216_v157_ = nw44_
        index22_ = default__.safeIndex(466, (d_216_v157_).length(0))
        (d_216_v157_)[index22_] = d_211_v154_
        d_217_v158_: _dafny.Map
        d_217_v158_ = _dafny.Map({len(d_147_v97_): d_206_v149_})
        d_218_v159_: _dafny.MultiSet
        d_218_v159_ = _dafny.MultiSet([p1])
        index23_ = default__.safeIndex(466, (d_216_v157_).length(0))
        rhs30_ = (d_206_v149_) | ((_dafny.Map({p0: d_205_v148_})) | (((d_217_v158_)[len(default__.fm19(d_218_v159_, len(d_208_v151_), p1, d_145_v95_, globalState))] if (len(default__.fm19(d_218_v159_, len(d_208_v151_), p1, d_145_v95_, globalState))) in (d_217_v158_) else d_206_v149_)))
        rhs31_ = d_210_v153_
        rhs32_ = d_211_v154_
        lhs4_ = d_216_v157_
        lhs5_ = default__.safeIndex(466, (d_216_v157_).length(0))
        d_206_v149_ = rhs30_
        d_210_v153_ = rhs31_
        lhs4_[lhs5_] = rhs32_
        d_219_i8_: int
        d_219_i8_ = 0
        with _dafny.label("0"):
            while (p1) <= (((0) - (d_145_v95_)) - (d_145_v95_)):
                with _dafny.c_label("0"):
                    if (d_219_i8_) >= (100):
                        raise _dafny.Break("0")
                    d_219_i8_ = (d_219_i8_) + (1)
                    d_220_v160_: T0
                    nw45_ = C4()
                    nw45_.ctor__(d_145_v95_)
                    d_220_v160_ = nw45_
                    d_145_v95_ = (D13_DC27(d_18_v0_, d_220_v160_, d_18_v0_, len(d_146_v96_))).cf46
                    d_221_v161_: _dafny.MultiSet
                    d_221_v161_ = _dafny.MultiSet([p0, True])
                    d_222_v162_: _dafny.Map
                    d_222_v162_ = _dafny.Map({(d_221_v161_).ispropersubset(d_221_v161_): d_208_v151_})
                    d_223_v163_: D7
                    d_223_v163_ = D7_DC18((d_220_v160_).f0)
                    rhs33_ = (d_220_v160_).f0
                    rhs34_ = default__.fm3(d_208_v151_, d_18_v0_, (d_149_v99_).f0, (d_223_v163_).cf30, globalState)
                    rhs35_ = d_222_v162_
                    rhs36_ = (d_148_v98_) <= ((d_148_v98_) + (_dafny.SeqWithoutIsStrInference([p0])))
                    d_145_v95_ = rhs33_
                    d_145_v95_ = rhs34_
                    d_222_v162_ = rhs35_
                    r0 = rhs36_
                    d_145_v95_ = len((d_146_v96_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvrdr"))) + (d_146_v96_)))
                    arr0_ = (d_216_v157_)[default__.safeIndex(466, (d_216_v157_).length(0))]
                    index24_ = default__.safeIndex(937, ((d_216_v157_)[default__.safeIndex(466, (d_216_v157_).length(0))]).length(0))
                    arr0_[index24_] = (d_149_v99_).f0
                    d_224_v164_: _dafny.Map
                    d_224_v164_ = _dafny.Map({default__.fm3(d_208_v151_, d_18_v0_, p1, (d_149_v99_).f0, globalState): p0})
                    d_225_v165_: C1
                    nw46_ = C1()
                    nw46_.ctor__(len(d_224_v164_))
                    d_225_v165_ = nw46_
                    d_226_v166_: _dafny.MultiSet
                    d_226_v166_ = _dafny.MultiSet([d_225_v165_])
                    arr1_ = (d_216_v157_)[default__.safeIndex(466, (d_216_v157_).length(0))]
                    index25_ = default__.safeIndex(937, ((d_216_v157_)[default__.safeIndex(466, (d_216_v157_).length(0))]).length(0))
                    arr1_[index25_] = default__.safeDivisionInt((((d_226_v166_)[d_225_v165_] if (d_225_v165_) in (d_226_v166_) else (d_149_v99_).f0) if p0 else (d_149_v99_).f0), (0) - ((len(d_148_v98_)) + (d_145_v95_)))
                    pass
            pass
        d_227_v167_: _dafny.Seq
        d_227_v167_ = _dafny.SeqWithoutIsStrInference([d_218_v159_])
        hi0_ = d_145_v95_
        for d_228_i9_ in range(((d_227_v167_)[default__.safeIndex(p1, len(d_227_v167_))]).cardinality, hi0_):
            d_229_v168_: _dafny.Array
            nw47_ = _dafny.Array(None, 2)
            nw47_[int(0)] = p0
            nw47_[int(1)] = p0
            d_229_v168_ = nw47_
            index26_ = default__.safeIndex(50, (d_229_v168_).length(0))
            (d_229_v168_)[index26_] = p0
            index27_ = default__.safeIndex(50, (d_229_v168_).length(0))
            (d_229_v168_)[index27_] = False
            if (d_229_v168_)[default__.safeIndex(50, (d_229_v168_).length(0))]:
                d_230_v169_: _dafny.Seq
                d_230_v169_ = _dafny.SeqWithoutIsStrInference([d_146_v96_])
                d_231_v170_: _dafny.Set
                d_231_v170_ = _dafny.Set({d_146_v96_, (d_146_v96_ if False else d_146_v96_), (d_146_v96_) + (d_146_v96_), (d_230_v169_)[default__.safeIndex(d_228_i9_, len(d_230_v169_))], d_146_v96_})
                d_231_v170_ = d_231_v170_
                d_145_v95_ = ((d_218_v159_)[(len((d_148_v98_).set(default__.safeIndex(d_228_i9_, len(d_148_v98_)), (d_229_v168_)[default__.safeIndex(50, (d_229_v168_).length(0))]))) + (d_228_i9_)] if ((len((d_148_v98_).set(default__.safeIndex(d_228_i9_, len(d_148_v98_)), (d_229_v168_)[default__.safeIndex(50, (d_229_v168_).length(0))]))) + (d_228_i9_)) in (d_218_v159_) else (len(d_151_v101_) if False else default__.fm3(d_208_v151_, d_18_v0_, len(d_146_v96_), (d_149_v99_).f0, globalState)))
                d_232_v171_: _dafny.Seq
                d_232_v171_ = _dafny.SeqWithoutIsStrInference([(d_229_v168_ if (d_229_v168_)[default__.safeIndex(50, (d_229_v168_).length(0))] else d_229_v168_), d_229_v168_])
                d_232_v171_ = _dafny.SeqWithoutIsStrInference([d_229_v168_])
                d_233_v172_: _dafny.MultiSet
                d_233_v172_ = _dafny.MultiSet([d_229_v168_])
                d_233_v172_ = ((d_233_v172_) | (_dafny.MultiSet([d_229_v168_]))) | (d_233_v172_)
                arr2_ = (d_216_v157_)[default__.safeIndex(466, (d_216_v157_).length(0))]
                index28_ = default__.safeIndex(124, ((d_216_v157_)[default__.safeIndex(466, (d_216_v157_).length(0))]).length(0))
                arr2_[index28_] = ((d_149_v99_).f0) - ((d_147_v97_)[default__.safeIndex(p1, len(d_147_v97_))])
                d_234_v173_: T0
                nw48_ = C1()
                nw48_.ctor__((d_149_v99_).f0)
                d_234_v173_ = nw48_
                d_235_v174_: D13
                d_235_v174_ = D13_DC27(_dafny.CodePoint('w'), d_234_v173_, _dafny.CodePoint('o'), (d_149_v99_).f0)
                arr3_ = (d_216_v157_)[default__.safeIndex(466, (d_216_v157_).length(0))]
                index29_ = default__.safeIndex(124, ((d_216_v157_)[default__.safeIndex(466, (d_216_v157_).length(0))]).length(0))
                arr3_[index29_] = default__.safeDivisionInt(default__.safeModuloInt(p1, (d_149_v99_).f0), (d_235_v174_).cf46)
            elif True:
                d_145_v95_ = (d_149_v99_).f0
                d_236_v175_: _dafny.Map
                d_236_v175_ = _dafny.Map({_dafny.MultiSet([d_18_v0_]): (d_229_v168_)[default__.safeIndex(50, (d_229_v168_).length(0))]})
                d_237_v176_: _dafny.MultiSet
                d_237_v176_ = _dafny.MultiSet([d_18_v0_, d_18_v0_])
                d_238_v178_: _dafny.Seq
                d_238_v178_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_18_v0_ for d_239_i10_ in range(default__.abs(788))]), d_146_v96_, d_146_v96_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hdjof")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edovto"))])
                def iife25_():
                    coll13_ = _dafny.Map()
                    compr_13_: _dafny.Seq
                    for compr_13_ in (d_238_v178_).Elements:
                        d_240_v177_: _dafny.Seq = compr_13_
                        if (d_240_v177_) in (d_238_v178_):
                            coll13_[d_240_v177_] = _dafny.SeqWithoutIsStrInference([(d_229_v168_)[default__.safeIndex(50, (d_229_v168_).length(0))]])
                    return _dafny.Map(coll13_)
                d_236_v175_ = (d_236_v175_).set((d_237_v176_) - (d_237_v176_), (_dafny.Map({d_146_v96_: d_148_v98_})) == (iife25_()
                ))
                d_241_v179_: _dafny.Map
                d_241_v179_ = _dafny.Map({71: d_146_v96_})
                d_145_v95_ = ((d_149_v99_).f0) * (len((((d_241_v179_)[(d_149_v99_).f0] if ((d_149_v99_).f0) in (d_241_v179_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lv"))) if p0 else _dafny.SeqWithoutIsStrInference([d_18_v0_ for d_242_i11_ in range(default__.abs(-667))]))))
                d_243_v180_: C5
                nw49_ = C5()
                nw49_.ctor__((d_18_v0_ if False else d_18_v0_), len(d_146_v96_), d_145_v95_)
                d_243_v180_ = nw49_
                index30_ = default__.safeIndex(50, (d_229_v168_).length(0))
                (d_229_v168_)[index30_] = p0
            d_244_v181_: _dafny.Array
            nw50_ = _dafny.Array(None, 21)
            d_244_v181_ = nw50_
            index31_ = default__.safeIndex(113, (d_244_v181_).length(0))
            (d_244_v181_)[index31_] = d_205_v148_
            index32_ = default__.safeIndex(113, (d_244_v181_).length(0))
            (d_244_v181_)[index32_] = (d_205_v148_ if ((d_149_v99_).f0) < ((0) - (d_228_i9_)) else d_205_v148_)
            d_245_v182_: _dafny.Seq
            d_245_v182_ = _dafny.SeqWithoutIsStrInference([d_229_v168_, d_229_v168_, d_229_v168_, d_229_v168_])
            d_229_v168_ = (d_245_v182_)[default__.safeIndex((0) - ((d_149_v99_).f0), len(d_245_v182_))]
        r0 = p0
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_246_globalState_: GlobalState
        nw51_ = GlobalState()
        nw51_.ctor__()
        d_246_globalState_ = nw51_
        d_247_v0_: _dafny.Array
        def lambda12_(d_248_i0_):
            return False

        init6_ = lambda12_
        nw52_ = _dafny.Array(None, 11)
        for i0_6_ in range(nw52_.length(0)):
            nw52_[i0_6_] = init6_(i0_6_)
        d_247_v0_ = nw52_
        index33_ = default__.safeIndex(231, (d_247_v0_).length(0))
        (d_247_v0_)[index33_] = (len(default__.fm0(d_246_globalState_))) <= (716)
        d_249_v1_: bool
        d_249_v1_ = False
        index34_ = default__.safeIndex(692, (d_247_v0_).length(0))
        (d_247_v0_)[index34_] = d_249_v1_
        d_250_v2_: int
        d_250_v2_ = -806
        d_251_v3_: _dafny.Array
        def lambda13_(d_252_v2_):
            def lambda14_(d_253_i1_):
                return (d_253_i1_) - (d_252_v2_)

            return lambda14_

        init7_ = lambda13_(d_250_v2_)
        nw53_ = _dafny.Array(None, 28)
        for i0_7_ in range(nw53_.length(0)):
            nw53_[i0_7_] = init7_(i0_7_)
        d_251_v3_ = nw53_
        d_254_v4_: _dafny.MultiSet
        d_254_v4_ = _dafny.MultiSet([d_251_v3_, d_251_v3_, d_251_v3_, d_251_v3_])
        d_255_v5_: _dafny.MultiSet
        d_255_v5_ = _dafny.MultiSet([d_250_v2_])
        index35_ = default__.safeIndex(231, (d_247_v0_).length(0))
        index36_ = default__.safeIndex(692, (d_247_v0_).length(0))
        rhs37_ = (True) or ((_dafny.MultiSet([d_250_v2_])).ispropersubset(d_255_v5_))
        rhs38_ = (d_247_v0_ if d_249_v1_ else d_247_v0_)
        rhs39_ = d_249_v1_
        rhs40_ = d_250_v2_
        rhs41_ = d_254_v4_
        lhs6_ = d_247_v0_
        lhs7_ = default__.safeIndex(231, (d_247_v0_).length(0))
        lhs8_ = d_247_v0_
        lhs9_ = default__.safeIndex(692, (d_247_v0_).length(0))
        lhs6_[lhs7_] = rhs37_
        d_247_v0_ = rhs38_
        lhs8_[lhs9_] = rhs39_
        d_250_v2_ = rhs40_
        d_254_v4_ = rhs41_
        if (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]:
            d_250_v2_ = d_250_v2_
            d_256_v6_: str
            d_256_v6_ = _dafny.CodePoint('p')
            d_250_v2_ = len(default__.fm1(d_256_v6_, d_246_globalState_))
            if (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]:
                d_250_v2_ = d_250_v2_
                d_257_v7_: _dafny.Map
                d_257_v7_ = _dafny.Map({d_249_v1_: len(_dafny.Set({not(d_249_v1_)}))})
                d_257_v7_ = (d_257_v7_).set(d_249_v1_, (d_255_v5_).cardinality)
                d_256_v6_ = d_256_v6_
                d_258_v8_: _dafny.Seq
                d_258_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqyjceeg"))
                d_259_v9_: bool
                out0_: bool
                out0_ = default__.m0(d_249_v1_, len(d_258_v8_), d_246_globalState_)
                d_259_v9_ = out0_
                d_260_v10_: _dafny.Seq
                d_260_v10_ = _dafny.SeqWithoutIsStrInference([False])
                d_261_v11_: D0
                d_261_v11_ = D0_DC1((d_260_v10_)[default__.safeIndex(d_250_v2_, len(d_260_v10_))])
                d_262_v12_: _dafny.Map
                d_262_v12_ = _dafny.Map({d_249_v1_: d_259_v9_})
                d_259_v9_ = (((d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))] if d_249_v1_ else d_249_v1_) if (not((d_261_v11_).cf1)) and ((d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]) else not(default__.fm2(d_259_v9_, d_250_v2_, not(True), default__.fm2((d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))], 888, d_249_v1_, ((d_262_v12_)[(d_260_v10_)[default__.safeIndex(d_250_v2_, len(d_260_v10_))]] if ((d_260_v10_)[default__.safeIndex(d_250_v2_, len(d_260_v10_))]) in (d_262_v12_) else d_259_v9_), d_246_globalState_), d_246_globalState_)))
            elif True:
                d_263_v13_: bool
                out1_: bool
                out1_ = default__.m0(((d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))] if (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))] else (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]), len(_dafny.SeqWithoutIsStrInference([d_256_v6_ for d_264_i2_ in range(default__.abs(725))])), d_246_globalState_)
                d_263_v13_ = out1_
                d_265_v14_: _dafny.Array
                nw54_ = _dafny.Array(None, 3)
                nw54_[int(0)] = d_247_v0_
                nw54_[int(1)] = d_247_v0_
                nw54_[int(2)] = d_247_v0_
                d_265_v14_ = nw54_
                index37_ = default__.safeIndex(794, (d_265_v14_).length(0))
                (d_265_v14_)[index37_] = d_247_v0_
                index38_ = default__.safeIndex(794, (d_265_v14_).length(0))
                (d_265_v14_)[index38_] = d_247_v0_
                index39_ = default__.safeIndex(231, (d_247_v0_).length(0))
                (d_247_v0_)[index39_] = (d_263_v13_) == (d_263_v13_)
                index40_ = default__.safeIndex(231, (d_247_v0_).length(0))
                (d_247_v0_)[index40_] = d_249_v1_
                d_263_v13_ = not((d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))])
            d_266_v15_: _dafny.Set
            d_266_v15_ = _dafny.Set({d_256_v6_})
            d_267_v16_: _dafny.Seq
            d_267_v16_ = _dafny.SeqWithoutIsStrInference([d_250_v2_, d_250_v2_])
            d_268_v17_: _dafny.Map
            d_268_v17_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_256_v6_ for d_269_i3_ in range(default__.abs(891))]): len(d_267_v16_)})
            d_249_v1_ = (len(d_266_v15_)) <= (len(d_268_v17_))
            d_270_v18_: _dafny.Map
            d_270_v18_ = _dafny.Map({d_250_v2_: d_250_v2_})
            d_271_v19_: _dafny.Map
            d_271_v19_ = _dafny.Map({((d_270_v18_)[d_250_v2_] if (d_250_v2_) in (d_270_v18_) else d_250_v2_): not(d_249_v1_)})
            index41_ = default__.safeIndex(231, (d_247_v0_).length(0))
            (d_247_v0_)[index41_] = (not (((d_271_v19_)[d_250_v2_] if (d_250_v2_) in (d_271_v19_) else d_249_v1_)) or (d_249_v1_) if d_249_v1_ else d_249_v1_)
        elif True:
            d_272_v20_: str
            d_272_v20_ = _dafny.CodePoint('o')
            d_273_v21_: _dafny.Seq
            d_273_v21_ = _dafny.SeqWithoutIsStrInference([(d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]])
            d_274_v22_: bool
            out2_: bool
            out2_ = default__.m0((d_250_v2_) > (default__.fm3(_dafny.Set({(d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]}), d_272_v20_, d_250_v2_, d_250_v2_, d_246_globalState_)), len(_dafny.Set({d_250_v2_, len(d_273_v21_)})), d_246_globalState_)
            d_274_v22_ = out2_
            if ((len(d_273_v21_)) <= ((0) - (len(d_273_v21_))) if d_249_v1_ else d_274_v22_):
                d_275_v23_: _dafny.Map
                d_275_v23_ = _dafny.Map({D0_DC0((d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]): d_274_v22_})
                d_276_v24_: _dafny.Set
                d_276_v24_ = _dafny.Set({d_274_v22_})
                rhs42_ = ((d_255_v5_) | (_dafny.MultiSet([default__.fm3(d_276_v24_, d_272_v20_, d_250_v2_, len(_dafny.SeqWithoutIsStrInference([D0_DC0(d_249_v1_) for d_277_i4_ in range(default__.abs(-155))])), d_246_globalState_), d_250_v2_]))) | ((d_255_v5_).set(d_250_v2_, default__.abs(d_250_v2_)))
                rhs43_ = default__.fm4(d_250_v2_, len(_dafny.SeqWithoutIsStrInference([-995 for d_278_i5_ in range(default__.abs(782))])), 185, d_246_globalState_)
                d_255_v5_ = rhs42_
                d_275_v23_ = rhs43_
                d_250_v2_ = (d_250_v2_ if d_274_v22_ else d_250_v2_)
                d_272_v20_ = d_272_v20_
                d_279_v25_: _dafny.MultiSet
                d_279_v25_ = _dafny.MultiSet([d_276_v24_])
                d_250_v2_ = default__.safeModuloInt(547, (d_279_v25_).cardinality)
                index42_ = default__.safeIndex(231, (d_247_v0_).length(0))
                (d_247_v0_)[index42_] = (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]
            elif True:
                d_250_v2_ = (((d_255_v5_) - (d_255_v5_)).cardinality) * (d_250_v2_)
                d_274_v22_ = (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]
                d_250_v2_ = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_250_v2_ for d_280_i6_ in range(default__.abs(989))])), d_250_v2_)
                d_281_v26_: _dafny.Seq
                d_281_v26_ = _dafny.SeqWithoutIsStrInference([d_250_v2_, d_250_v2_])
                d_282_v27_: _dafny.MultiSet
                d_282_v27_ = _dafny.MultiSet([d_249_v1_, (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))], d_249_v1_])
                d_283_v28_: _dafny.Set
                d_283_v28_ = _dafny.Set({(d_282_v27_).cardinality})
                d_284_v29_: _dafny.Seq
                d_284_v29_ = _dafny.SeqWithoutIsStrInference([((d_281_v26_).set(default__.safeIndex(d_250_v2_, len(d_281_v26_)), len(d_283_v28_))) + (d_281_v26_), (d_281_v26_) + (_dafny.SeqWithoutIsStrInference([d_250_v2_ for d_285_i7_ in range(default__.abs(514))]))])
                index43_ = default__.safeIndex(231, (d_247_v0_).length(0))
                rhs44_ = not(d_249_v1_)
                rhs45_ = len(d_284_v29_)
                lhs10_ = d_247_v0_
                lhs11_ = default__.safeIndex(231, (d_247_v0_).length(0))
                lhs10_[lhs11_] = rhs44_
                d_250_v2_ = rhs45_
                index44_ = default__.safeIndex(515, (d_251_v3_).length(0))
                (d_251_v3_)[index44_] = ((d_255_v5_)[d_250_v2_] if (d_250_v2_) in (d_255_v5_) else (0) - (d_250_v2_))
                index45_ = default__.safeIndex(515, (d_251_v3_).length(0))
                (d_251_v3_)[index45_] = d_250_v2_
            d_286_v30_: _dafny.Seq
            d_286_v30_ = _dafny.SeqWithoutIsStrInference([d_250_v2_])
            d_287_v31_: bool
            out3_: bool
            out3_ = default__.m0((d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))], (d_286_v30_)[default__.safeIndex(d_250_v2_, len(d_286_v30_))], d_246_globalState_)
            d_287_v31_ = out3_
            d_288_v32_: _dafny.Array
            nw55_ = _dafny.Array(None, 6)
            nw55_[int(0)] = (d_247_v0_)
            nw55_[int(1)] = d_247_v0_
            nw55_[int(2)] = d_247_v0_
            nw55_[int(3)] = d_247_v0_
            nw55_[int(4)] = d_247_v0_
            nw55_[int(5)] = d_247_v0_
            d_288_v32_ = nw55_
            index46_ = default__.safeIndex(215, (d_288_v32_).length(0))
            (d_288_v32_)[index46_] = d_247_v0_
            d_289_v33_: _dafny.Array
            nw56_ = _dafny.Array(None, 9)
            nw56_[int(0)] = d_251_v3_
            nw56_[int(1)] = d_251_v3_
            nw56_[int(2)] = d_251_v3_
            nw56_[int(3)] = d_251_v3_
            nw56_[int(4)] = d_251_v3_
            nw56_[int(5)] = d_251_v3_
            nw56_[int(6)] = d_251_v3_
            nw56_[int(7)] = d_251_v3_
            nw56_[int(8)] = d_251_v3_
            d_289_v33_ = nw56_
            d_290_v34_: _dafny.Seq
            d_290_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
            index47_ = default__.safeIndex(807, (d_251_v3_).length(0))
            (d_251_v3_)[index47_] = len(d_290_v34_)
            d_291_v35_: D0
            d_291_v35_ = D0_DC1(not(d_249_v1_))
            d_292_v36_: _dafny.Seq
            d_292_v36_ = _dafny.SeqWithoutIsStrInference([d_291_v35_])
            d_293_v37_: _dafny.Map
            d_293_v37_ = _dafny.Map({(d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]: (D2_DC3((_dafny.SeqWithoutIsStrInference([d_291_v35_ for d_294_i8_ in range(default__.abs(82))])).set(default__.safeIndex(d_250_v2_, len(_dafny.SeqWithoutIsStrInference([d_291_v35_ for d_295_i8_ in range(default__.abs(82))]))), d_291_v35_))).cf3})
            d_296_v38_: _dafny.Seq
            d_296_v38_ = _dafny.SeqWithoutIsStrInference([d_292_v36_, ((d_293_v37_)[d_274_v22_] if (d_274_v22_) in (d_293_v37_) else d_292_v36_), d_292_v36_, _dafny.SeqWithoutIsStrInference([d_291_v35_, d_291_v35_])])
            index48_ = default__.safeIndex(215, (d_288_v32_).length(0))
            index49_ = default__.safeIndex(807, (d_251_v3_).length(0))
            rhs46_ = d_247_v0_
            rhs47_ = d_289_v33_
            rhs48_ = (d_292_v36_) <= ((d_296_v38_)[default__.safeIndex(977, len(d_296_v38_))])
            rhs49_ = (d_250_v2_) - (len(_dafny.Set({d_274_v22_})))
            lhs12_ = d_288_v32_
            lhs13_ = default__.safeIndex(215, (d_288_v32_).length(0))
            lhs14_ = d_251_v3_
            lhs15_ = default__.safeIndex(807, (d_251_v3_).length(0))
            lhs12_[lhs13_] = rhs46_
            d_289_v33_ = rhs47_
            d_274_v22_ = rhs48_
            lhs14_[lhs15_] = rhs49_
            index50_ = default__.safeIndex(231, (d_247_v0_).length(0))
            (d_247_v0_)[index50_] = not((d_273_v21_)[default__.safeIndex((d_251_v3_)[default__.safeIndex(807, (d_251_v3_).length(0))], len(d_273_v21_))])
        if False:
            index51_ = default__.safeIndex(231, (d_247_v0_).length(0))
            (d_247_v0_)[index51_] = (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]
            d_297_v39_: bool
            out4_: bool
            out4_ = default__.m0(not((d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]), d_250_v2_, d_246_globalState_)
            d_297_v39_ = out4_
            d_247_v0_ = d_247_v0_
            index52_ = default__.safeIndex(231, (d_247_v0_).length(0))
            rhs50_ = (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]
            rhs51_ = d_297_v39_
            rhs52_ = d_250_v2_
            lhs16_ = d_247_v0_
            lhs17_ = default__.safeIndex(231, (d_247_v0_).length(0))
            d_297_v39_ = rhs50_
            lhs16_[lhs17_] = rhs51_
            d_250_v2_ = rhs52_
            d_298_v40_: _dafny.Seq
            d_298_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qn"))
            d_299_v41_: _dafny.Set
            d_299_v41_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fylsrkxqe"))), d_250_v2_, d_250_v2_})
            index53_ = default__.safeIndex(231, (d_247_v0_).length(0))
            (d_247_v0_)[index53_] = ((len(d_298_v40_)) == (d_250_v2_) if (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))] else (d_299_v41_).ispropersubset(d_299_v41_))
        elif True:
            d_300_v42_: _dafny.MultiSet
            d_300_v42_ = _dafny.MultiSet([d_249_v1_, not(d_249_v1_), (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]])
            d_301_v43_: _dafny.Seq
            d_301_v43_ = _dafny.SeqWithoutIsStrInference([d_249_v1_])
            d_300_v42_ = _dafny.MultiSet(((d_301_v43_ if d_249_v1_ else _dafny.SeqWithoutIsStrInference([d_249_v1_, not(d_249_v1_)]))) + (d_301_v43_))
            d_250_v2_ = (0) - ((0) - (d_250_v2_))
            d_302_v44_: _dafny.Map
            d_302_v44_ = _dafny.Map({(d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]: (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]})
            d_303_v45_: bool
            out5_: bool
            out5_ = default__.m0(not((_dafny.Map({False: d_249_v1_})) == ((d_302_v44_).set(d_249_v1_, default__.fm2((d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))], d_250_v2_, (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))], d_249_v1_, d_246_globalState_)))), d_250_v2_, d_246_globalState_)
            d_303_v45_ = out5_
            if not(d_249_v1_):
                index54_ = default__.safeIndex(226, (d_251_v3_).length(0))
                (d_251_v3_)[index54_] = len(_dafny.Set({d_303_v45_, (d_301_v43_)[default__.safeIndex(930, len(d_301_v43_))]}))
                index55_ = default__.safeIndex(226, (d_251_v3_).length(0))
                (d_251_v3_)[index55_] = default__.safeModuloInt(d_250_v2_, 501)
                d_250_v2_ = d_250_v2_
                d_249_v1_ = d_249_v1_
                d_302_v44_ = (d_302_v44_).set(d_303_v45_, d_303_v45_)
                d_303_v45_ = False
            elif True:
                d_304_v46_: bool
                out6_: bool
                out6_ = default__.m0(not(d_303_v45_), d_250_v2_, d_246_globalState_)
                d_304_v46_ = out6_
                d_305_v47_: bool
                out7_: bool
                out7_ = default__.m0(d_249_v1_, d_250_v2_, d_246_globalState_)
                d_305_v47_ = out7_
                d_250_v2_ = (0) - ((d_250_v2_) + (d_250_v2_))
                d_306_v48_: C1
                nw57_ = C1()
                nw57_.ctor__(d_250_v2_)
                d_306_v48_ = nw57_
                d_307_v49_: str
                d_307_v49_ = _dafny.CodePoint('g')
                d_308_v50_: _dafny.Seq
                d_308_v50_ = _dafny.SeqWithoutIsStrInference([d_307_v49_])
                d_309_v51_: _dafny.Map
                d_309_v51_ = _dafny.Map({d_308_v50_: d_303_v45_})
                d_309_v51_ = (d_309_v51_).set(_dafny.SeqWithoutIsStrInference([d_307_v49_]), (d_301_v43_) < (d_301_v43_))
            d_310_v52_: str
            d_310_v52_ = _dafny.CodePoint('y')
            d_311_v53_: C5
            nw58_ = C5()
            nw58_.ctor__(d_310_v52_, d_250_v2_, d_250_v2_)
            d_311_v53_ = nw58_
        d_312_v54_: C3
        nw59_ = C3()
        nw59_.ctor__(d_250_v2_)
        d_312_v54_ = nw59_
        d_313_v55_: D16
        d_313_v55_ = D16_DC33(d_312_v54_)
        d_314_v56_: _dafny.Array
        nw60_ = _dafny.Array(None, 5)
        nw60_[int(0)] = d_312_v54_
        nw60_[int(1)] = d_312_v54_
        nw60_[int(2)] = d_312_v54_
        nw60_[int(3)] = d_312_v54_
        nw60_[int(4)] = (d_313_v55_).cf54
        d_314_v56_ = nw60_
        d_315_v57_: _dafny.Seq
        d_315_v57_ = _dafny.SeqWithoutIsStrInference([d_314_v56_])
        d_316_v58_: str
        d_316_v58_ = _dafny.CodePoint('x')
        d_317_v59_: _dafny.Map
        d_317_v59_ = _dafny.Map({(d_315_v57_)[default__.safeIndex(d_250_v2_, len(d_315_v57_))]: d_316_v58_})
        d_317_v59_ = (d_317_v59_) | (d_317_v59_)
        d_250_v2_ = d_250_v2_
        d_316_v58_ = d_316_v58_
        d_318_v60_: D2
        d_318_v60_ = D2_DC4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ontdavgjc")), (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))])
        pat_let_tv2_ = d_318_v60_
        pat_let_tv3_ = d_318_v60_
        pat_let_tv4_ = d_318_v60_
        index56_ = default__.safeIndex(231, (d_247_v0_).length(0))
        def lambda15_(source1_):
            if source1_.is_DC12:
                d_319___mcc_h0_ = source1_.cf15
                d_320___mcc_h1_ = source1_.cf16
                d_321_cf16_ = d_320___mcc_h1_
                d_322_cf15_ = d_319___mcc_h0_
                return pat_let_tv2_
            elif source1_.is_DC13:
                d_323___mcc_h2_ = source1_.cf17
                d_324___mcc_h3_ = source1_.cf18
                d_325___mcc_h4_ = source1_.cf19
                d_326___mcc_h5_ = source1_.cf20
                d_327___mcc_h6_ = source1_.cf21
                d_328_cf21_ = d_327___mcc_h6_
                d_329_cf20_ = d_326___mcc_h5_
                d_330_cf19_ = d_325___mcc_h4_
                d_331_cf18_ = d_324___mcc_h3_
                d_332_cf17_ = d_323___mcc_h2_
                return pat_let_tv3_
            elif True:
                d_333___mcc_h7_ = source1_.cf14
                d_334_cf14_ = d_333___mcc_h7_
                return pat_let_tv4_

        rhs53_ = ((d_250_v2_) + (d_250_v2_)) * (-502)
        rhs54_ = (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]
        rhs55_ = lambda15_(D5_DC12(d_250_v2_, d_250_v2_))
        lhs18_ = d_247_v0_
        lhs19_ = default__.safeIndex(231, (d_247_v0_).length(0))
        d_250_v2_ = rhs53_
        lhs18_[lhs19_] = rhs54_
        d_318_v60_ = rhs55_
        d_335_v61_: _dafny.Set
        d_335_v61_ = _dafny.Set({(d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]})
        d_336_i9_: int
        d_336_i9_ = 0
        with _dafny.label("1"):
            while (d_335_v61_).ispropersubset(_dafny.Set({d_249_v1_, d_249_v1_, (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))], (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]})):
                with _dafny.c_label("1"):
                    if (d_336_i9_) >= (100):
                        raise _dafny.Break("1")
                    d_336_i9_ = (d_336_i9_) + (1)
                    d_337_v62_: D3
                    d_337_v62_ = D3_DC8((d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))], d_250_v2_)
                    d_338_v63_: bool
                    out8_: bool
                    out8_ = default__.m0(((0) - (default__.fm3(_dafny.Set({d_249_v1_, d_249_v1_}), d_316_v58_, d_250_v2_, d_250_v2_, d_246_globalState_))) < ((d_337_v62_).cf9), d_250_v2_, d_246_globalState_)
                    d_338_v63_ = out8_
                    index57_ = default__.safeIndex(777, (d_251_v3_).length(0))
                    (d_251_v3_)[index57_] = default__.safeDivisionInt(d_250_v2_, len(_dafny.SeqWithoutIsStrInference([(D13_DC29(_dafny.CodePoint('j'), (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))])).cf49 for d_339_i10_ in range(default__.abs(542))])))
                    d_340_v64_: C5
                    nw61_ = C5()
                    nw61_.ctor__(d_316_v58_, d_250_v2_, d_250_v2_)
                    d_340_v64_ = nw61_
                    d_341_v65_: _dafny.Map
                    d_341_v65_ = _dafny.Map({d_340_v64_: d_250_v2_})
                    index58_ = default__.safeIndex(777, (d_251_v3_).length(0))
                    (d_251_v3_)[index58_] = ((d_341_v65_)[d_340_v64_] if (d_340_v64_) in (d_341_v65_) else d_250_v2_)
                    d_342_v66_: _dafny.Array
                    nw62_ = _dafny.Array(_dafny.Array(None, 0), 9)
                    d_342_v66_ = nw62_
                    d_342_v66_ = d_342_v66_
                    index59_ = default__.safeIndex(777, (d_251_v3_).length(0))
                    (d_251_v3_)[index59_] = (d_340_v64_).f2
                    pass
            pass
        d_250_v2_ = default__.safeDivisionInt(d_250_v2_, d_250_v2_)
        d_343_i11_: int
        d_343_i11_ = 0
        with _dafny.label("2"):
            while not(d_249_v1_):
                with _dafny.c_label("2"):
                    if (d_343_i11_) >= (100):
                        raise _dafny.Break("2")
                    d_343_i11_ = (d_343_i11_) + (1)
                    d_249_v1_ = d_249_v1_
                    d_250_v2_ = default__.safeModuloInt(default__.safeModuloInt(d_250_v2_, 770), d_250_v2_)
                    if (d_250_v2_) == (d_250_v2_):
                        d_344_v68_: T0
                        nw63_ = C3()
                        def iife26_():
                            coll14_ = _dafny.Map()
                            compr_14_: int
                            for compr_14_ in (_dafny.Map({d_250_v2_: d_249_v1_})).keys.Elements:
                                d_345_v67_: int = compr_14_
                                if (d_345_v67_) in (_dafny.Map({d_250_v2_: d_249_v1_})):
                                    coll14_[default__.safeDivisionInt(d_345_v67_, len(d_335_v61_))] = d_250_v2_
                            return _dafny.Map(coll14_)
                        nw63_.ctor__(len(iife26_()
                        ))
                        d_344_v68_ = nw63_
                        d_346_v69_: D13
                        d_346_v69_ = D13_DC27(d_316_v58_, d_344_v68_, d_316_v58_, d_250_v2_)
                        d_250_v2_ = (d_346_v69_).cf46
                        d_347_v70_: _dafny.Set
                        d_347_v70_ = _dafny.Set({d_247_v0_})
                        d_250_v2_ = default__.safeDivisionInt(len(d_347_v70_), (0) - ((d_346_v69_).cf46))
                        d_348_v71_: _dafny.Seq
                        d_348_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmqirv"))
                        (d_312_v54_).m7((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oaqldlvsi"))) + (d_348_v71_), d_249_v1_, d_251_v3_, len(d_348_v71_), d_246_globalState_)
                        d_250_v2_ = (d_344_v68_).f0
                        d_349_v72_: _dafny.Seq
                        d_349_v72_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_316_v58_ for d_350_i12_ in range(default__.abs(580))]), _dafny.SeqWithoutIsStrInference([d_316_v58_ for d_351_i13_ in range(default__.abs(259))])])
                        d_352_v73_: _dafny.Seq
                        d_352_v73_ = _dafny.SeqWithoutIsStrInference([(d_250_v2_) <= (-839), (d_348_v71_) <= ((d_349_v72_)[default__.safeIndex((d_344_v68_).f0, len(d_349_v72_))]), d_249_v1_])
                        d_353_v74_: _dafny.Map
                        d_353_v74_ = _dafny.Map({d_316_v58_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmslbj")))})
                        d_354_v75_: _dafny.Set
                        d_354_v75_ = _dafny.Set({((d_353_v74_)[d_316_v58_] if (d_316_v58_) in (d_353_v74_) else d_250_v2_), (d_344_v68_).f0})
                        d_355_v76_: _dafny.MultiSet
                        d_355_v76_ = _dafny.MultiSet([d_316_v58_, d_316_v58_])
                        d_356_v77_: _dafny.Map
                        d_356_v77_ = _dafny.Map({(d_344_v68_).f0: _dafny.MultiSet([d_316_v58_])})
                        d_357_v78_: _dafny.Seq
                        d_357_v78_ = _dafny.SeqWithoutIsStrInference([d_355_v76_, ((d_356_v77_)[d_250_v2_] if (d_250_v2_) in (d_356_v77_) else d_355_v76_)])
                        rhs56_ = default__.fm3(d_335_v61_, (d_316_v58_ if (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))] else d_316_v58_), (d_250_v2_) * ((d_344_v68_).f0), default__.safeDivisionInt(len(d_348_v71_), len(d_354_v75_)), d_246_globalState_)
                        rhs57_ = (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]
                        rhs58_ = d_352_v73_
                        rhs59_ = default__.fm10(((d_357_v78_)[default__.safeIndex(447, len(d_357_v78_))]).intersection(d_355_v76_), (d_312_v54_).fm13(d_250_v2_, (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))], d_246_globalState_), d_246_globalState_)
                        rhs60_ = d_316_v58_
                        d_250_v2_ = rhs56_
                        d_249_v1_ = rhs57_
                        d_352_v73_ = rhs58_
                        d_335_v61_ = rhs59_
                        d_316_v58_ = rhs60_
                    elif True:
                        d_250_v2_ = d_250_v2_
                        index60_ = default__.safeIndex(231, (d_247_v0_).length(0))
                        (d_247_v0_)[index60_] = (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]
                        d_358_v79_: _dafny.Map
                        d_358_v79_ = _dafny.Map({d_250_v2_: False})
                        d_359_v80_: _dafny.MultiSet
                        d_359_v80_ = _dafny.MultiSet([d_316_v58_, d_316_v58_])
                        d_360_v81_: _dafny.Map
                        d_360_v81_ = _dafny.Map({d_359_v80_: d_251_v3_})
                        d_358_v79_ = (d_358_v79_).set(d_250_v2_, (d_359_v80_) in ((d_360_v81_).set(d_359_v80_, d_251_v3_)))
                        d_361_v82_: _dafny.Map
                        d_361_v82_ = _dafny.Map({default__.safeDivisionInt(d_250_v2_, d_250_v2_): (_dafny.Set({(d_312_v54_).fm13(-924, (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))], d_246_globalState_)})) | (d_335_v61_)})
                        d_361_v82_ = (d_361_v82_).set(d_250_v2_, d_335_v61_)
                        d_362_v83_: D6
                        d_362_v83_ = D6_DC14(d_251_v3_)
                        d_363_v84_: _dafny.Seq
                        d_363_v84_ = _dafny.SeqWithoutIsStrInference([d_250_v2_])
                        d_364_v85_: _dafny.Seq
                        d_364_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryvnthp"))
                        d_365_v86_: C4
                        nw64_ = C4()
                        nw64_.ctor__(d_250_v2_)
                        d_365_v86_ = nw64_
                        d_366_v87_: _dafny.Seq
                        d_366_v87_ = _dafny.SeqWithoutIsStrInference([d_365_v86_])
                        d_367_v88_: _dafny.Seq
                        d_367_v88_ = _dafny.SeqWithoutIsStrInference([d_365_v86_, (d_366_v87_)[default__.safeIndex(d_250_v2_, len(d_366_v87_))]])
                        (d_312_v54_).m7(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_368_i14_ in range(default__.abs(627))]), d_249_v1_, (d_251_v3_ if (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))] else (d_362_v83_).cf22), (0) - (len(((d_363_v84_).set(default__.safeIndex(180, len(d_363_v84_)), d_250_v2_)) + (_dafny.SeqWithoutIsStrInference([len(d_364_v85_), len(d_367_v88_), d_250_v2_])))), d_246_globalState_)
                    index61_ = default__.safeIndex(231, (d_247_v0_).length(0))
                    (d_247_v0_)[index61_] = d_249_v1_
                    pass
            pass
        d_369_i15_: int
        d_369_i15_ = 0
        with _dafny.label("3"):
            while (d_249_v1_) or ((d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]):
                with _dafny.c_label("3"):
                    if (d_369_i15_) >= (100):
                        raise _dafny.Break("3")
                    d_369_i15_ = (d_369_i15_) + (1)
                    d_370_v89_: T1
                    nw65_ = C5()
                    nw65_.ctor__(d_316_v58_, d_250_v2_, d_250_v2_)
                    d_370_v89_ = nw65_
                    d_371_v90_: _dafny.Array
                    nw66_ = _dafny.Array(None, 13)
                    d_371_v90_ = nw66_
                    d_372_v91_: _dafny.Map
                    d_372_v91_ = _dafny.Map({d_370_v89_: d_371_v90_})
                    d_373_v92_: D3
                    d_373_v92_ = D3_DC9(d_250_v2_, ((d_372_v91_)[d_370_v89_] if (d_370_v89_) in (d_372_v91_) else d_371_v90_), (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))])
                    d_374_v93_: _dafny.Set
                    d_374_v93_ = _dafny.Set({(d_373_v92_).cf10, (d_370_v89_).f0, 20})
                    d_375_v94_: D5
                    d_375_v94_ = D5_DC13((d_250_v2_) * (default__.fm3(_dafny.Set({d_249_v1_}), d_316_v58_, d_250_v2_, d_250_v2_, d_246_globalState_)), d_335_v61_, d_316_v58_, default__.safeModuloInt(d_250_v2_, d_250_v2_), (d_374_v93_).isdisjoint(d_374_v93_))
                    d_375_v94_ = d_375_v94_
                    d_376_v95_: _dafny.Map
                    d_376_v95_ = _dafny.Map({(d_370_v89_).f0: (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]})
                    d_376_v95_ = (d_376_v95_).set((d_370_v89_).f0, d_249_v1_)
                    d_251_v3_ = d_251_v3_
                    index62_ = default__.safeIndex(923, (d_251_v3_).length(0))
                    (d_251_v3_)[index62_] = (0) - ((d_370_v89_).f0)
                    index63_ = default__.safeIndex(923, (d_251_v3_).length(0))
                    (d_251_v3_)[index63_] = d_250_v2_
                    pass
            pass
        d_377_v96_: C3
        nw67_ = C3()
        nw67_.ctor__(default__.safeDivisionInt(747, 988))
        d_377_v96_ = nw67_
        d_378_v97_: _dafny.Seq
        d_378_v97_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")) for d_379_i16_ in range(default__.abs(811))])
        d_380_v98_: _dafny.Map
        d_380_v98_ = _dafny.Map({(d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))]: d_378_v97_})
        d_380_v98_ = (d_380_v98_).set((d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))], d_378_v97_)
        d_381_v99_: _dafny.Seq
        d_381_v99_ = _dafny.SeqWithoutIsStrInference([d_316_v58_])
        (d_312_v54_).m7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nujtypxb")), (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))], (d_251_v3_ if True else d_251_v3_), len(d_381_v99_), d_246_globalState_)
        d_382_v100_: _dafny.Array
        nw68_ = _dafny.Array(_dafny.Array(None, 0), 22)
        d_382_v100_ = nw68_
        d_383_v101_: int
        d_384_v102_: int
        d_385_v103_: int
        out9_: int
        out10_: int
        out11_: int
        out9_, out10_, out11_ = (d_312_v54_).m1(d_382_v100_, ((0) - (d_250_v2_) if d_249_v1_ else (0) - (d_250_v2_)), d_246_globalState_)
        d_383_v101_ = out9_
        d_384_v102_ = out10_
        d_385_v103_ = out11_
        d_386_v104_: _dafny.Seq
        d_386_v104_ = _dafny.SeqWithoutIsStrInference([d_250_v2_])
        hi1_ = (d_250_v2_ if (d_247_v0_)[default__.safeIndex(231, (d_247_v0_).length(0))] else len(_dafny.SeqWithoutIsStrInference([d_316_v58_ for d_387_i18_ in range(default__.abs(-266))])))
        for d_388_i17_ in range((d_386_v104_)[default__.safeIndex(d_383_v101_, len(d_386_v104_))], hi1_):
            d_389_v105_: _dafny.Map
            d_389_v105_ = _dafny.Map({(d_255_v5_).ispropersubset(_dafny.MultiSet([d_250_v2_])): d_247_v0_})
            d_389_v105_ = (d_389_v105_).set(((d_318_v60_).cf4) != (d_381_v99_), (d_247_v0_ if d_249_v1_ else d_247_v0_))
            index64_ = default__.safeIndex(231, (d_247_v0_).length(0))
            (d_247_v0_)[index64_] = d_249_v1_
            d_390_v106_: _dafny.Map
            d_390_v106_ = _dafny.Map({d_250_v2_: d_250_v2_})
            rhs61_ = ((d_390_v106_)[default__.fm3(d_335_v61_, d_316_v58_, d_388_i17_, d_388_i17_, d_246_globalState_)] if (default__.fm3(d_335_v61_, d_316_v58_, d_388_i17_, d_388_i17_, d_246_globalState_)) in (d_390_v106_) else d_384_v102_)
            rhs62_ = (d_381_v99_) + (d_381_v99_)
            d_250_v2_ = rhs61_
            d_381_v99_ = rhs62_
            d_391_v107_: _dafny.Array
            def lambda16_(d_392_v104_):
                def lambda17_(d_393_i19_):
                    return (_dafny.Set({len(d_392_v104_)}))

                return lambda17_

            init8_ = lambda16_(d_386_v104_)
            nw69_ = _dafny.Array(None, 20)
            for i0_8_ in range(nw69_.length(0)):
                nw69_[i0_8_] = init8_(i0_8_)
            d_391_v107_ = nw69_
            d_391_v107_ = d_391_v107_
        _dafny.print(_dafny.string_of((d_247_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_249_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_250_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v3_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v4_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_255_v5_) == (_dafny.MultiSet([-806]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_313_v55_).cf54).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_314_v56_)[0]).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_314_v56_)[1]).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_314_v56_)[2]).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_314_v56_)[3]).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_314_v56_)[4]).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_315_v57_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_316_v58_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_317_v59_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v60_).cf4).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_318_v60_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_v61_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_336_i9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_343_i11_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_369_i15_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_378_v97_)) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_380_v98_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgqps"))])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_381_v99_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_383_v101_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_384_v102_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_385_v103_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_v104_) == (_dafny.SeqWithoutIsStrInference([1]))))
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
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
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
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC4(D2, NamedTuple('DC4', [('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({self.cf4.VerbatimString(True)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(False, int(0))
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

class D3_DC8(D3, NamedTuple('DC8', [('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.CodePoint('D')
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC10(D4, NamedTuple('DC10', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC12(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)

class D5_DC12(D5, NamedTuple('DC12', [('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC11(D5, NamedTuple('DC11', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC15(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC15(D6, NamedTuple('DC15', [('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
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
        return lambda: D7_DC17(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)

class D7_DC17(D7, NamedTuple('DC17', [('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)

class D8_DC19(D8, NamedTuple('DC19', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)

class D9_DC20(D9, NamedTuple('DC20', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC22(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D10_DC22)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D10_DC21)

class D10_DC22(D10, NamedTuple('DC22', [('cf34', Any), ('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC21(D10, NamedTuple('DC21', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC21({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC21) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC24(_dafny.Array(None, 0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D11_DC24)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D11_DC23)

class D11_DC24(D11, NamedTuple('DC24', [('cf38', Any), ('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC23(D11, NamedTuple('DC23', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC23({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC23) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D12_DC25)

class D12_DC25(D12, NamedTuple('DC25', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC25({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC25) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC27(_dafny.CodePoint('D'), None, _dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D13_DC27)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D13_DC28)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D13_DC29)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D13_DC26)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D13_DC30)

class D13_DC27(D13, NamedTuple('DC27', [('cf43', Any), ('cf44', Any), ('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC27({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC27) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC28(D13, NamedTuple('DC28', [('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC28({_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC28) and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC29(D13, NamedTuple('DC29', [('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC29({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC29) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC26(D13, NamedTuple('DC26', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC26({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC26) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC30(D13, NamedTuple('DC30', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC30({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC30) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D14_DC31)

class D14_DC31(D14, NamedTuple('DC31', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC31({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC31) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D15_DC32)

class D15_DC32(D15, NamedTuple('DC32', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC32({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC32) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC34()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D16_DC34)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D16_DC35)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D16_DC33)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D16_DC36)

class D16_DC34(D16, NamedTuple('DC34', [])):
    def __dafnystr__(self) -> str:
        return f'D16.DC34'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC34)
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC35(D16, NamedTuple('DC35', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC35({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC35) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC33(D16, NamedTuple('DC33', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC33({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC33) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC36(D16, NamedTuple('DC36', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC36({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC36) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D17_DC37)

class D17_DC37(D17, NamedTuple('DC37', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC37({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC37) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC39(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D18_DC39)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D18_DC40)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D18_DC38)

class D18_DC39(D18, NamedTuple('DC39', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC39({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC39) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC40(D18, NamedTuple('DC40', [('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC40({_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC40) and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC38(D18, NamedTuple('DC38', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC38({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC38) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC42(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D19_DC42)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D19_DC41)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D19_DC43)

class D19_DC42(D19, NamedTuple('DC42', [('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC42({_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC42) and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC41(D19, NamedTuple('DC41', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC41({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC41) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC43(D19, NamedTuple('DC43', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC43({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC43) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, p1, globalState):
        pass


class T1:
    pass
    def m2(self, globalState):
        pass


class GlobalState:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self):
        pass
        pass


class C0:
    def  __init__(self):
        self._f3: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f3):
        (self)._f3 = f3

    @property
    def f3(self):
        return self._f3

class C1(T0):
    def  __init__(self):
        self._f0: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f0(self):
        return self._f0
    def ctor__(self, f0):
        (self)._f0 = f0

    def fm5(self, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))) != ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_394_i0_ in range(default__.abs(600))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))))

    def fm6(self, p0, p1, p2, p3, globalState):
        return not (False) or (True)

    def fm15(self, p0, p1, p2, globalState):
        return (self).f0

    def fm16(self, p0, p1, globalState):
        return not((_dafny.MultiSet([False])).issubset(_dafny.MultiSet([False, not(False)])))

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        d_395_v0_: bool
        d_395_v0_ = True
        d_396_i0_: int
        d_396_i0_ = 0
        with _dafny.label("4"):
            while (d_395_v0_ if True else (not(d_395_v0_) if d_395_v0_ else d_395_v0_)):
                with _dafny.c_label("4"):
                    if (d_396_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_396_i0_ = (d_396_i0_) + (1)
                    d_397_v1_: _dafny.Array
                    nw70_ = _dafny.Array(int(0), 17)
                    d_397_v1_ = nw70_
                    index65_ = default__.safeIndex(959, (d_397_v1_).length(0))
                    (d_397_v1_)[index65_] = (self).f0
                    index66_ = default__.safeIndex(959, (d_397_v1_).length(0))
                    (d_397_v1_)[index66_] = p1
                    d_398_v2_: bool
                    out12_: bool
                    out12_ = default__.m0(((0) - (-237)) > (p1), (self).f0, globalState)
                    d_398_v2_ = out12_
                    d_399_v3_: str
                    d_399_v3_ = _dafny.CodePoint('h')
                    d_400_v4_: _dafny.Map
                    d_400_v4_ = _dafny.Map({d_399_v3_: -500})
                    d_400_v4_ = (d_400_v4_) | (d_400_v4_)
                    d_401_v5_: C0
                    nw71_ = C0()
                    nw71_.ctor__(d_399_v3_)
                    d_401_v5_ = nw71_
                    pass
            pass
        d_402_v6_: str
        d_402_v6_ = _dafny.CodePoint('e')
        d_403_v7_: C0
        nw72_ = C0()
        nw72_.ctor__(d_402_v6_)
        d_403_v7_ = nw72_
        d_404_v8_: _dafny.Seq
        d_404_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fce"))
        d_405_v9_: _dafny.Map
        d_405_v9_ = _dafny.Map({d_404_v8_: d_395_v0_})
        d_406_v10_: _dafny.Seq
        d_406_v10_ = _dafny.SeqWithoutIsStrInference([584])
        r0 = (self).fm15(d_405_v9_, not(d_395_v0_), (d_406_v10_) == (d_406_v10_), globalState)
        d_407_v11_: _dafny.Array
        nw73_ = _dafny.Array(int(0), 13)
        d_407_v11_ = nw73_
        index67_ = default__.safeIndex(489, (d_407_v11_).length(0))
        (d_407_v11_)[index67_] = (0) - (default__.safeModuloInt((self).f0, (0) - (p1)))
        index68_ = default__.safeIndex(489, (d_407_v11_).length(0))
        (d_407_v11_)[index68_] = ((self).f0 if d_395_v0_ else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flbjqa"))))
        d_408_v12_: _dafny.MultiSet
        d_408_v12_ = _dafny.MultiSet([d_395_v0_])
        d_409_v13_: _dafny.Map
        d_409_v13_ = _dafny.Map({D3_DC7(d_406_v10_): (d_408_v12_) == (d_408_v12_)})
        d_410_v14_: D3
        d_410_v14_ = D3_DC7(d_406_v10_)
        d_409_v13_ = (d_409_v13_).set(d_410_v14_, (d_408_v12_).issubset(d_408_v12_))
        d_411_v15_: _dafny.Map
        d_411_v15_ = _dafny.Map({(0) - (p1): _dafny.SeqWithoutIsStrInference([(d_403_v7_).f3 for d_412_i2_ in range(default__.abs(718))])})
        d_413_v16_: _dafny.Array
        nw74_ = _dafny.Array(None, 25)
        nw74_[int(0)] = (d_404_v8_ if d_395_v0_ else (d_404_v8_).set(default__.safeIndex(len(d_404_v8_), len(d_404_v8_)), _dafny.CodePoint('y')))
        nw74_[int(1)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btdr"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_414_i1_ in range(default__.abs(543))]))
        nw74_[int(2)] = d_404_v8_
        nw74_[int(3)] = ((d_411_v15_)[(self).f0] if ((self).f0) in (d_411_v15_) else d_404_v8_)
        nw74_[int(4)] = d_404_v8_
        nw74_[int(5)] = (d_404_v8_).set(default__.safeIndex((d_407_v11_)[default__.safeIndex(489, (d_407_v11_).length(0))], len(d_404_v8_)), default__.fm17(p1, globalState))
        nw74_[int(6)] = d_404_v8_
        nw74_[int(7)] = d_404_v8_
        nw74_[int(8)] = d_404_v8_
        nw74_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cx"))
        nw74_[int(10)] = (d_404_v8_) + (d_404_v8_)
        nw74_[int(11)] = d_404_v8_
        nw74_[int(12)] = d_404_v8_
        nw74_[int(13)] = (d_404_v8_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")))
        nw74_[int(14)] = d_404_v8_
        nw74_[int(15)] = d_404_v8_
        nw74_[int(16)] = d_404_v8_
        nw74_[int(17)] = _dafny.SeqWithoutIsStrInference([(d_403_v7_).f3 for d_415_i3_ in range(default__.abs(730))])
        nw74_[int(18)] = d_404_v8_
        nw74_[int(19)] = d_404_v8_
        nw74_[int(20)] = (d_404_v8_) + (d_404_v8_)
        nw74_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxyrf"))
        nw74_[int(22)] = d_404_v8_
        nw74_[int(23)] = d_404_v8_
        nw74_[int(24)] = _dafny.SeqWithoutIsStrInference([d_402_v6_ for d_416_i4_ in range(default__.abs(-916))])
        d_413_v16_ = nw74_
        d_417_v17_: D0
        d_417_v17_ = D0_DC0(not((d_404_v8_) < (d_404_v8_)))
        d_418_v18_: _dafny.Map
        d_418_v18_ = _dafny.Map({d_395_v0_: d_395_v0_})
        index69_ = default__.safeIndex(489, (d_407_v11_).length(0))
        rhs63_ = (d_413_v16_ if d_395_v0_ else d_413_v16_)
        rhs64_ = p1
        rhs65_ = (d_404_v8_ if d_395_v0_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhobkl")))
        rhs66_ = (D0_DC0(((d_418_v18_)[d_395_v0_] if (d_395_v0_) in (d_418_v18_) else True)) if ((self).f0) >= ((0) - ((d_407_v11_)[default__.safeIndex(489, (d_407_v11_).length(0))])) else d_417_v17_)
        lhs20_ = d_407_v11_
        lhs21_ = default__.safeIndex(489, (d_407_v11_).length(0))
        d_413_v16_ = rhs63_
        lhs20_[lhs21_] = rhs64_
        d_404_v8_ = rhs65_
        d_417_v17_ = rhs66_
        r0 = (self).f0
        r1 = len(((d_404_v8_) + (d_404_v8_)) + (d_404_v8_))
        r2 = p1
        return r0, r1, r2


class C2(T1, T0):
    def  __init__(self):
        self._f0: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f0(self):
        return self._f0
    def ctor__(self, f0):
        (self)._f0 = f0

    def fm5(self, globalState):
        return ((True if True else True)) and (not(not (True) or (False)))

    def fm6(self, p0, p1, p2, p3, globalState):
        return not((False) or (True))

    def m2(self, globalState):
        d_419_v0_: bool
        d_419_v0_ = True
        d_420_v1_: _dafny.Array
        nw75_ = _dafny.Array(None, 2)
        nw75_[int(0)] = d_419_v0_
        nw75_[int(1)] = d_419_v0_
        d_420_v1_ = nw75_
        d_421_v2_: _dafny.MultiSet
        d_421_v2_ = _dafny.MultiSet([d_420_v1_])
        d_422_v3_: _dafny.Seq
        d_422_v3_ = _dafny.SeqWithoutIsStrInference([(d_421_v2_).cardinality])
        pat_let_tv5_ = d_422_v3_
        pat_let_tv6_ = d_422_v3_
        def lambda18_(source2_):
            if source2_.is_DC1:
                d_423___mcc_h0_ = source2_.cf1
                d_424_cf1_ = d_423___mcc_h0_
                return pat_let_tv5_
            elif True:
                d_425___mcc_h1_ = source2_.cf0
                d_426_cf0_ = d_425___mcc_h1_
                return pat_let_tv6_

        d_422_v3_ = lambda18_(D0_DC0(d_419_v0_))
        d_427_v4_: _dafny.Seq
        d_427_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdfea"))
        d_428_v5_: D2
        d_428_v5_ = D2_DC4(d_427_v4_, d_419_v0_)
        pat_let_tv7_ = d_419_v0_
        pat_let_tv8_ = d_419_v0_
        pat_let_tv9_ = d_419_v0_
        pat_let_tv10_ = d_419_v0_
        pat_let_tv11_ = d_419_v0_
        def lambda19_(source3_):
            if source3_.is_DC4:
                d_429___mcc_h2_ = source3_.cf4
                d_430___mcc_h3_ = source3_.cf5
                d_431_cf5_ = d_430___mcc_h3_
                d_432_cf4_ = d_429___mcc_h2_
                return pat_let_tv7_
            elif source3_.is_DC5:
                return pat_let_tv8_
            elif source3_.is_DC3:
                d_433___mcc_h4_ = source3_.cf3
                d_434_cf3_ = d_433___mcc_h4_
                return not((pat_let_tv9_) or (pat_let_tv10_))
            elif True:
                d_435___mcc_h5_ = source3_.cf6
                d_436_cf6_ = d_435___mcc_h5_
                return pat_let_tv11_

        d_419_v0_ = lambda19_(d_428_v5_)
        hi2_ = 221
        for d_437_i0_ in range((self).f0, hi2_):
            d_438_v6_: _dafny.Array
            nw76_ = _dafny.Array(_dafny.Map({}), 26)
            d_438_v6_ = nw76_
            d_439_v7_: _dafny.Map
            d_439_v7_ = _dafny.Map({d_420_v1_: (self).f0})
            index70_ = default__.safeIndex(184, (d_438_v6_).length(0))
            (d_438_v6_)[index70_] = d_439_v7_
            d_440_v8_: _dafny.Map
            d_440_v8_ = _dafny.Map({(self).fm5(globalState): d_439_v7_})
            index71_ = default__.safeIndex(184, (d_438_v6_).length(0))
            (d_438_v6_)[index71_] = (((d_440_v8_)[d_419_v0_] if (d_419_v0_) in (d_440_v8_) else d_439_v7_)).set(d_420_v1_, d_437_i0_)
            d_441_v9_: int
            d_441_v9_ = 715
            d_442_v10_: _dafny.MultiSet
            d_442_v10_ = _dafny.MultiSet([d_419_v0_, not(d_419_v0_), d_419_v0_, d_419_v0_, True])
            d_441_v9_ = ((d_442_v10_)[(d_427_v4_) <= (d_427_v4_)] if ((d_427_v4_) <= (d_427_v4_)) in (d_442_v10_) else d_437_i0_)
            d_443_v11_: _dafny.Seq
            d_443_v11_ = _dafny.SeqWithoutIsStrInference([d_419_v0_, d_419_v0_])
            d_444_v12_: str
            d_444_v12_ = _dafny.CodePoint('t')
            d_445_v13_: bool
            out13_: bool
            out13_ = default__.m0(((self).f0) < (d_437_i0_), default__.fm3(_dafny.Set({default__.fm2(d_419_v0_, -904, d_419_v0_, d_419_v0_, globalState), d_419_v0_, (d_443_v11_)[default__.safeIndex(610, len(d_443_v11_))], d_419_v0_, d_419_v0_}), d_444_v12_, d_441_v9_, d_441_v9_, globalState), globalState)
            d_445_v13_ = out13_
            index72_ = default__.safeIndex(89, (d_420_v1_).length(0))
            (d_420_v1_)[index72_] = False
            d_446_v14_: _dafny.Map
            d_446_v14_ = _dafny.Map({(d_437_i0_) > ((self).f0): d_441_v9_})
            d_447_v15_: _dafny.Map
            d_447_v15_ = _dafny.Map({d_446_v14_: d_443_v11_})
            d_448_v16_: _dafny.Set
            d_448_v16_ = _dafny.Set({d_419_v0_, d_419_v0_, True, d_445_v13_})
            index73_ = default__.safeIndex(89, (d_420_v1_).length(0))
            rhs67_ = ((d_446_v14_)[True] if (True) in (d_446_v14_) else (0) - ((0) - (d_441_v9_)))
            rhs68_ = (d_437_i0_) >= ((0) - (d_437_i0_))
            rhs69_ = (d_443_v11_) + ((d_443_v11_) + (((d_447_v15_)[d_446_v14_] if (d_446_v14_) in (d_447_v15_) else d_443_v11_)))
            rhs70_ = not((_dafny.Set({d_445_v13_})) == (d_448_v16_))
            lhs22_ = d_420_v1_
            lhs23_ = default__.safeIndex(89, (d_420_v1_).length(0))
            d_441_v9_ = rhs67_
            lhs22_[lhs23_] = rhs68_
            d_443_v11_ = rhs69_
            d_419_v0_ = rhs70_
        d_449_v17_: int
        d_449_v17_ = 653
        d_450_v18_: _dafny.Set
        d_450_v18_ = _dafny.Set({d_419_v0_})
        d_451_v19_: str
        d_451_v19_ = _dafny.CodePoint('h')
        d_452_v20_: _dafny.Map
        d_452_v20_ = _dafny.Map({d_449_v17_: (self).fm5(globalState)})
        d_453_v21_: _dafny.MultiSet
        d_453_v21_ = _dafny.MultiSet([d_452_v20_, d_452_v20_])
        d_449_v17_ = default__.fm3((d_450_v18_) | (d_450_v18_), d_451_v19_, (d_449_v17_) + (((d_453_v21_)[d_452_v20_] if (d_452_v20_) in (d_453_v21_) else d_449_v17_)), (self).f0, globalState)
        hi3_ = d_449_v17_
        for d_454_i1_ in range(d_449_v17_, hi3_):
            d_455_v22_: _dafny.MultiSet
            d_455_v22_ = _dafny.MultiSet([d_454_i1_, (self).f0])
            d_427_v4_ = ((d_427_v4_) + (_dafny.SeqWithoutIsStrInference([d_451_v19_ for d_456_i2_ in range(default__.abs(585))]))) + (default__.fm19(d_455_v22_, d_454_i1_, (self).f0, d_449_v17_, globalState))
            d_457_v23_: _dafny.Array
            nw77_ = _dafny.Array(int(0), 29)
            d_457_v23_ = nw77_
            d_458_v24_: _dafny.Set
            d_458_v24_ = _dafny.Set({d_457_v23_})
            rhs71_ = d_422_v3_
            rhs72_ = not (d_419_v0_) or (d_419_v0_)
            rhs73_ = len(d_458_v24_)
            rhs74_ = (0) - (d_449_v17_)
            d_422_v3_ = rhs71_
            d_419_v0_ = rhs72_
            d_449_v17_ = rhs73_
            d_449_v17_ = rhs74_
            d_459_v25_: _dafny.Map
            d_459_v25_ = _dafny.Map({not(d_419_v0_): d_420_v1_})
            d_420_v1_ = ((d_459_v25_)[d_419_v0_] if (d_419_v0_) in (d_459_v25_) else d_420_v1_)
            d_460_v26_: C0
            nw78_ = C0()
            nw78_.ctor__(_dafny.CodePoint('n'))
            d_460_v26_ = nw78_
        d_461_v27_: _dafny.Array
        nw79_ = _dafny.Array(int(0), 11)
        d_461_v27_ = nw79_
        d_462_v28_: _dafny.MultiSet
        d_462_v28_ = _dafny.MultiSet([d_461_v27_, d_461_v27_, d_461_v27_, d_461_v27_, d_461_v27_])
        d_449_v17_ = default__.safeModuloInt((self).f0, default__.safeModuloInt((d_462_v28_).cardinality, (self).f0))

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        d_463_v0_: bool
        d_463_v0_ = False
        d_463_v0_ = d_463_v0_
        d_464_v1_: _dafny.Set
        d_464_v1_ = _dafny.Set({(self).f0, (self).f0})
        d_463_v0_ = ((d_464_v1_).intersection(d_464_v1_)) != (d_464_v1_)
        d_465_v2_: _dafny.Array
        def lambda20_(d_466_i1_):
            return _dafny.CodePoint('i')

        init9_ = lambda20_
        nw80_ = _dafny.Array(None, 20)
        for i0_9_ in range(nw80_.length(0)):
            nw80_[i0_9_] = init9_(i0_9_)
        d_465_v2_ = nw80_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_465_v2_).length(0)):
            d_467_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_467_i0_)) and ((d_467_i0_) < ((d_465_v2_).length(0)))):
                (d_465_v2_)[(d_467_i0_)] = _dafny.CodePoint('h')
        d_468_v3_: _dafny.MultiSet
        d_468_v3_ = _dafny.MultiSet([d_463_v0_, (self).fm5(globalState), (self).fm5(globalState)])
        d_469_v4_: _dafny.Map
        d_469_v4_ = _dafny.Map({(d_468_v3_).cardinality: p1})
        d_470_v5_: _dafny.Array
        nw81_ = _dafny.Array(None, 20)
        nw81_[int(0)] = False
        nw81_[int(1)] = d_463_v0_
        nw81_[int(2)] = default__.fm2(d_463_v0_, 847, d_463_v0_, d_463_v0_, globalState)
        nw81_[int(3)] = d_463_v0_
        nw81_[int(4)] = d_463_v0_
        nw81_[int(5)] = d_463_v0_
        nw81_[int(6)] = d_463_v0_
        nw81_[int(7)] = d_463_v0_
        nw81_[int(8)] = d_463_v0_
        nw81_[int(9)] = d_463_v0_
        nw81_[int(10)] = d_463_v0_
        nw81_[int(11)] = d_463_v0_
        nw81_[int(12)] = True
        nw81_[int(13)] = d_463_v0_
        nw81_[int(14)] = d_463_v0_
        nw81_[int(15)] = True
        nw81_[int(16)] = d_463_v0_
        nw81_[int(17)] = d_463_v0_
        nw81_[int(18)] = d_463_v0_
        nw81_[int(19)] = d_463_v0_
        d_470_v5_ = nw81_
        d_471_v6_: _dafny.Array
        d_471_v6_ = d_470_v5_
        d_472_v7_: _dafny.Map
        d_472_v7_ = _dafny.Map({d_471_v6_: -32})
        d_473_v8_: _dafny.Array
        nw82_ = _dafny.Array(int(0), 24)
        d_473_v8_ = nw82_
        d_474_v9_: _dafny.MultiSet
        d_474_v9_ = _dafny.MultiSet([d_473_v8_])
        d_475_v10_: _dafny.Set
        d_475_v10_ = _dafny.Set({d_463_v0_, d_463_v0_})
        d_476_v11_: str
        d_476_v11_ = _dafny.CodePoint('k')
        d_477_v12_: _dafny.Seq
        d_477_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_463_v0_])])
        d_478_v13_: _dafny.Seq
        d_478_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltx"))
        d_479_v14_: _dafny.Seq
        d_479_v14_ = _dafny.SeqWithoutIsStrInference([True, d_463_v0_])
        d_480_v15_: D7
        d_480_v15_ = D7_DC16(d_479_v14_)
        d_481_v16_: _dafny.Array
        nw83_ = _dafny.Array(None, 21)
        nw83_[int(0)] = 98
        nw83_[int(1)] = len(d_469_v4_)
        nw83_[int(2)] = default__.safeModuloInt((self).f0, (self).f0)
        nw83_[int(3)] = (124) * ((self).f0)
        nw83_[int(4)] = len(d_472_v7_)
        nw83_[int(5)] = p1
        nw83_[int(6)] = (d_474_v9_).cardinality
        nw83_[int(7)] = (self).f0
        nw83_[int(8)] = p1
        nw83_[int(9)] = 71
        nw83_[int(10)] = (self).f0
        nw83_[int(11)] = p1
        nw83_[int(12)] = default__.safeDivisionInt((self).f0, (0) - (default__.fm3(d_475_v10_, d_476_v11_, len(d_477_v12_), p1, globalState)))
        nw83_[int(13)] = len((d_478_v13_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbu"))))
        nw83_[int(14)] = p1
        nw83_[int(15)] = ((d_468_v3_)[d_463_v0_] if (d_463_v0_) in (d_468_v3_) else (self).f0)
        nw83_[int(16)] = (len((d_480_v15_).cf26) if d_463_v0_ else (self).f0)
        nw83_[int(17)] = p1
        nw83_[int(18)] = (self).f0
        nw83_[int(19)] = p1
        nw83_[int(20)] = (self).f0
        d_481_v16_ = nw83_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_481_v16_).length(0)):
            d_482_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_482_i2_)) and ((d_482_i2_) < ((d_481_v16_).length(0)))):
                (d_481_v16_)[(d_482_i2_)] = (d_482_i2_) + (p1)
        d_483_v17_: D5
        d_483_v17_ = D5_DC13(p1, d_475_v10_, d_476_v11_, 701, d_463_v0_)
        d_484_i3_: int
        d_484_i3_ = 0
        with _dafny.label("5"):
            while (d_483_v17_).cf21:
                with _dafny.c_label("5"):
                    if (d_484_i3_) >= (100):
                        raise _dafny.Break("5")
                    d_484_i3_ = (d_484_i3_) + (1)
                    d_485_v18_: _dafny.Seq
                    d_485_v18_ = _dafny.SeqWithoutIsStrInference([d_473_v8_])
                    d_486_v19_: _dafny.Map
                    d_486_v19_ = _dafny.Map({(d_485_v18_) + (d_485_v18_): d_463_v0_})
                    d_487_v20_: D2
                    d_487_v20_ = D2_DC4(d_478_v13_, not(d_463_v0_))
                    d_486_v19_ = (d_486_v19_).set(d_485_v18_, (d_487_v20_).cf5)
                    d_488_v21_: _dafny.Map
                    d_488_v21_ = _dafny.Map({d_463_v0_: d_476_v11_})
                    d_489_v22_: _dafny.Seq
                    d_489_v22_ = _dafny.SeqWithoutIsStrInference([(self).f0, len(_dafny.Map({d_463_v0_: d_463_v0_})), len(default__.fm20(d_463_v0_, d_488_v21_, True, d_476_v11_, globalState)), (self).f0, p1])
                    rhs75_ = ((self).f0) != (default__.safeModuloInt((self).f0, (self).f0))
                    rhs76_ = d_476_v11_
                    rhs77_ = p1
                    rhs78_ = (d_489_v22_) == (d_489_v22_)
                    rhs79_ = default__.safeDivisionInt((self).f0, (self).f0)
                    d_463_v0_ = rhs75_
                    d_476_v11_ = rhs76_
                    r1 = rhs77_
                    d_463_v0_ = rhs78_
                    r2 = rhs79_
                    d_463_v0_ = default__.fm2((d_463_v0_) and (d_463_v0_), ((self).f0) + ((self).f0), not(d_463_v0_), d_463_v0_, globalState)
                    d_463_v0_ = (357) <= (len(d_475_v10_))
                    pass
            pass
        hi4_ = p1
        for d_490_i4_ in range(len(_dafny.SeqWithoutIsStrInference([(self).f0])), hi4_):
            d_463_v0_ = ((0) - ((self).f0)) == (p1)
            d_463_v0_ = d_463_v0_
            if d_463_v0_:
                d_491_v23_: _dafny.Map
                d_491_v23_ = _dafny.Map({len(d_478_v13_): d_463_v0_})
                d_492_v24_: _dafny.Map
                d_492_v24_ = _dafny.Map({d_463_v0_: len(default__.fm19(default__.fm22(globalState), d_490_i4_, (self).f0, len(d_491_v23_), globalState))})
                d_493_v25_: _dafny.Seq
                d_493_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_463_v0_: (self).f0}), default__.fm21(globalState), d_492_v24_])
                d_493_v25_ = d_493_v25_
                rhs80_ = d_490_i4_
                rhs81_ = d_476_v11_
                rhs82_ = ((d_490_i4_) - (p1)) < (p1)
                r1 = rhs80_
                d_476_v11_ = rhs81_
                d_463_v0_ = rhs82_
                d_463_v0_ = d_463_v0_
                r2 = default__.fm3(d_475_v10_, d_476_v11_, p1, (p1) + (-34), globalState)
                r0 = d_490_i4_
            elif True:
                r2 = d_490_i4_
                d_473_v8_ = d_473_v8_
                d_494_v26_: _dafny.MultiSet
                d_494_v26_ = _dafny.MultiSet([d_468_v3_])
                d_495_v27_: _dafny.MultiSet
                d_495_v27_ = _dafny.MultiSet([len(d_478_v13_), p1, ((d_494_v26_)[_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, d_463_v0_]))] if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, d_463_v0_]))) in (d_494_v26_) else len(d_478_v13_))])
                d_496_v28_: _dafny.Map
                d_496_v28_ = _dafny.Map({(d_470_v5_ if d_463_v0_ else d_470_v5_): d_495_v27_})
                d_496_v28_ = (d_496_v28_) | (d_496_v28_)
                d_473_v8_ = d_481_v16_
                d_497_v29_: _dafny.Array
                nw84_ = _dafny.Array(_dafny.Seq({}), 12)
                d_497_v29_ = nw84_
                index74_ = default__.safeIndex(802, (d_497_v29_).length(0))
                (d_497_v29_)[index74_] = _dafny.SeqWithoutIsStrInference([d_490_i4_])
                d_498_v30_: _dafny.Map
                d_498_v30_ = _dafny.Map({p1: d_463_v0_})
                d_499_v31_: _dafny.Seq
                d_499_v31_ = _dafny.SeqWithoutIsStrInference([(d_468_v3_).cardinality, default__.safeDivisionInt(len(d_498_v30_), (self).f0)])
                index75_ = default__.safeIndex(802, (d_497_v29_).length(0))
                (d_497_v29_)[index75_] = d_499_v31_
            if (d_490_i4_) != (d_490_i4_):
                d_500_v32_: _dafny.MultiSet
                d_500_v32_ = _dafny.MultiSet([d_490_i4_])
                d_463_v0_ = (((d_500_v32_).cardinality) * ((self).f0)) == (p1)
                r0 = ((self).f0 if False else p1)
                index76_ = default__.safeIndex(47, (d_470_v5_).length(0))
                (d_470_v5_)[index76_] = not((self).fm6(d_463_v0_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_501_i5_ in range(default__.abs(-317))]), d_478_v13_, d_490_i4_, globalState))
                index77_ = default__.safeIndex(47, (d_470_v5_).length(0))
                (d_470_v5_)[index77_] = (not((-183) >= (d_490_i4_)) if d_463_v0_ else d_463_v0_)
                r1 = len(d_478_v13_)
                r0 = default__.fm3((d_475_v10_).intersection(d_475_v10_), d_476_v11_, p1, ((self).f0) - (p1), globalState)
            elif True:
                d_476_v11_ = d_476_v11_
                d_502_v33_: _dafny.Map
                d_502_v33_ = _dafny.Map({d_463_v0_: d_473_v8_})
                d_502_v33_ = (d_502_v33_).set(d_463_v0_, d_481_v16_)
                index78_ = default__.safeIndex(157, (d_470_v5_).length(0))
                (d_470_v5_)[index78_] = d_463_v0_
                index79_ = default__.safeIndex(157, (d_470_v5_).length(0))
                (d_470_v5_)[index79_] = (d_463_v0_ if False else (d_479_v14_)[default__.safeIndex(p1, len(d_479_v14_))])
                d_503_v34_: _dafny.Map
                d_503_v34_ = _dafny.Map({D7_DC18(p1): (d_470_v5_)[default__.safeIndex(157, (d_470_v5_).length(0))]})
                d_504_v36_: D7
                d_504_v36_ = D7_DC18((self).f0)
                d_505_v37_: _dafny.Map
                def iife27_():
                    coll15_ = _dafny.Map()
                    compr_15_: D7
                    for compr_15_ in (_dafny.Map({d_504_v36_: d_463_v0_})).keys.Elements:
                        d_506_v35_: D7 = compr_15_
                        if (d_506_v35_) in (_dafny.Map({d_504_v36_: d_463_v0_})):
                            coll15_[d_506_v35_] = (d_470_v5_)[default__.safeIndex(157, (d_470_v5_).length(0))]
                    return _dafny.Map(coll15_)
                d_505_v37_ = _dafny.Map({(0) - (default__.safeModuloInt(p1, (self).f0)): (d_503_v34_ if d_463_v0_ else iife27_()
                )})
                d_505_v37_ = (d_505_v37_) | ((d_505_v37_) | (d_505_v37_))
                d_476_v11_ = d_476_v11_
        d_507_v38_: _dafny.Seq
        d_507_v38_ = _dafny.SeqWithoutIsStrInference([len(d_478_v13_)])
        r0 = len((d_507_v38_) + (d_507_v38_))
        r1 = p1
        r2 = p1
        return r0, r1, r2


class C3(T0):
    def  __init__(self):
        self._f0: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f0(self):
        return self._f0
    def ctor__(self, f0):
        (self)._f0 = f0

    def fm5(self, globalState):
        return not(not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdrpodwh"))) == ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_508_i0_ in range(default__.abs(491))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hehmjmhcu"))))))

    def fm6(self, p0, p1, p2, p3, globalState):
        if True:
            return not((False if False else True))
        elif True:
            return (True) and (True)

    def fm13(self, p0, p1, globalState):
        return True

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        d_509_v0_: _dafny.Set
        d_509_v0_ = _dafny.Set({p1, (self).f0})
        d_510_v1_: _dafny.Set
        d_510_v1_ = _dafny.Set({d_509_v0_, d_509_v0_, d_509_v0_})
        d_511_v2_: bool
        d_511_v2_ = True
        d_512_v3_: _dafny.Seq
        d_512_v3_ = _dafny.SeqWithoutIsStrInference([d_510_v1_])
        d_513_v4_: _dafny.Map
        d_513_v4_ = _dafny.Map({p1: d_511_v2_})
        d_514_v5_: _dafny.Set
        d_514_v5_ = _dafny.Set({d_511_v2_})
        d_515_v6_: str
        d_515_v6_ = _dafny.CodePoint('j')
        d_516_v7_: _dafny.Map
        d_516_v7_ = _dafny.Map({d_509_v0_: d_511_v2_})
        d_517_v9_: _dafny.Seq
        d_517_v9_ = _dafny.SeqWithoutIsStrInference([d_509_v0_, d_509_v0_])
        d_518_v11_: _dafny.Array
        nw85_ = _dafny.Array(None, 21)
        nw85_[int(0)] = d_510_v1_
        nw85_[int(1)] = d_510_v1_
        nw85_[int(2)] = d_510_v1_
        nw85_[int(3)] = (d_510_v1_).intersection(d_510_v1_)
        nw85_[int(4)] = (d_510_v1_) | (d_510_v1_)
        nw85_[int(5)] = d_510_v1_
        nw85_[int(6)] = ((d_512_v3_)[default__.safeIndex(len(d_513_v4_), len(d_512_v3_))] if d_511_v2_ else d_510_v1_)
        nw85_[int(7)] = d_510_v1_
        nw85_[int(8)] = (d_512_v3_)[default__.safeIndex(default__.fm3(d_514_v5_, d_515_v6_, (self).f0, p1, globalState), len(d_512_v3_))]
        def iife28_():
            coll16_ = _dafny.Set()
            compr_16_: _dafny.Set
            for compr_16_ in (d_516_v7_).keys.Elements:
                d_519_v8_: _dafny.Set = compr_16_
                if (d_519_v8_) in (d_516_v7_):
                    coll16_ = coll16_.union(_dafny.Set([d_519_v8_]))
            return _dafny.Set(coll16_)
        nw85_[int(9)] = iife28_()
        
        nw85_[int(10)] = (_dafny.Set({d_509_v0_})).intersection(_dafny.Set({_dafny.Set({p1}), _dafny.Set({p1}), d_509_v0_, d_509_v0_, d_509_v0_}))
        nw85_[int(11)] = d_510_v1_
        nw85_[int(12)] = d_510_v1_
        nw85_[int(13)] = _dafny.Set({d_509_v0_, d_509_v0_, d_509_v0_, d_509_v0_})
        nw85_[int(14)] = d_510_v1_
        nw85_[int(15)] = _dafny.Set({d_509_v0_})
        nw85_[int(16)] = d_510_v1_
        def iife29_():
            coll17_ = _dafny.Set()
            compr_17_: _dafny.Set
            for compr_17_ in (d_517_v9_).Elements:
                d_520_v10_: _dafny.Set = compr_17_
                if (d_520_v10_) in (d_517_v9_):
                    coll17_ = coll17_.union(_dafny.Set([d_520_v10_]))
            return _dafny.Set(coll17_)
        nw85_[int(17)] = iife29_()
        
        nw85_[int(18)] = d_510_v1_
        nw85_[int(19)] = d_510_v1_
        nw85_[int(20)] = (d_510_v1_) | (d_510_v1_)
        d_518_v11_ = nw85_
        index80_ = default__.safeIndex(125, (d_518_v11_).length(0))
        (d_518_v11_)[index80_] = _dafny.Set({d_509_v0_})
        d_521_v12_: D5
        d_521_v12_ = D5_DC11(_dafny.Set({d_509_v0_, d_509_v0_}))
        index81_ = default__.safeIndex(125, (d_518_v11_).length(0))
        (d_518_v11_)[index81_] = (d_521_v12_).cf14
        source4_ = default__.fm14((self).f0, d_511_v2_, d_515_v6_, globalState)
        if source4_.is_DC4:
            d_522___mcc_h0_ = source4_.cf4
            d_523___mcc_h1_ = source4_.cf5
            d_524_cf5_ = d_523___mcc_h1_
            d_525_cf4_ = d_522___mcc_h0_
            d_526_v13_: _dafny.Array
            nw86_ = _dafny.Array(False, 25)
            d_526_v13_ = nw86_
            d_527_v14_: _dafny.Seq
            d_527_v14_ = _dafny.SeqWithoutIsStrInference([False, True, d_524_cf5_, d_524_cf5_])
            index82_ = default__.safeIndex(315, (d_526_v13_).length(0))
            (d_526_v13_)[index82_] = not(not (d_511_v2_) or ((d_527_v14_)[default__.safeIndex(p1, len(d_527_v14_))]))
            index83_ = default__.safeIndex(315, (d_526_v13_).length(0))
            (d_526_v13_)[index83_] = d_511_v2_
            d_528_v15_: _dafny.Map
            d_528_v15_ = _dafny.Map({d_511_v2_: d_525_cf4_})
            d_528_v15_ = (d_528_v15_).set((_dafny.Set({p1, (self).f0})).ispropersubset(d_509_v0_), d_525_cf4_)
            r2 = len(d_525_cf4_)
            d_526_v13_ = d_526_v13_
        elif source4_.is_DC5:
            d_529_v16_: _dafny.Map
            d_529_v16_ = _dafny.Map({False: p1})
            d_530_v17_: _dafny.Seq
            d_530_v17_ = _dafny.SeqWithoutIsStrInference([d_511_v2_, d_511_v2_, d_511_v2_])
            d_531_v18_: _dafny.Map
            d_531_v18_ = _dafny.Map({p1: len(d_530_v17_)})
            d_532_v19_: _dafny.Seq
            d_532_v19_ = _dafny.SeqWithoutIsStrInference([default__.fm3(d_514_v5_, d_515_v6_, len(d_531_v18_), (self).f0, globalState)])
            d_533_v20_: _dafny.Map
            d_533_v20_ = _dafny.Map({p1: d_515_v6_})
            d_534_v21_: _dafny.Array
            nw87_ = _dafny.Array(None, 26)
            nw87_[int(0)] = ((d_529_v16_)[d_511_v2_] if (d_511_v2_) in (d_529_v16_) else -553)
            nw87_[int(1)] = p1
            nw87_[int(2)] = p1
            nw87_[int(3)] = ((self).f0 if (self).fm5(globalState) else (self).f0)
            nw87_[int(4)] = (d_532_v19_)[default__.safeIndex((self).f0, len(d_532_v19_))]
            nw87_[int(5)] = p1
            nw87_[int(6)] = ((self).f0 if False else 345)
            nw87_[int(7)] = (0) - (p1)
            nw87_[int(8)] = 610
            nw87_[int(9)] = (0) - (p1)
            nw87_[int(10)] = default__.fm3(d_514_v5_, ((d_533_v20_)[p1] if (p1) in (d_533_v20_) else d_515_v6_), (self).f0, (self).f0, globalState)
            nw87_[int(11)] = (self).f0
            nw87_[int(12)] = (self).f0
            nw87_[int(13)] = len(d_509_v0_)
            nw87_[int(14)] = (self).f0
            nw87_[int(15)] = -41
            nw87_[int(16)] = p1
            nw87_[int(17)] = len(d_532_v19_)
            nw87_[int(18)] = (0) - ((self).f0)
            nw87_[int(19)] = p1
            nw87_[int(20)] = 808
            nw87_[int(21)] = p1
            nw87_[int(22)] = (self).f0
            nw87_[int(23)] = p1
            nw87_[int(24)] = (self).f0
            nw87_[int(25)] = p1
            d_534_v21_ = nw87_
            index84_ = default__.safeIndex(856, (d_534_v21_).length(0))
            (d_534_v21_)[index84_] = p1
            index85_ = default__.safeIndex(856, (d_534_v21_).length(0))
            (d_534_v21_)[index85_] = (p1) + (default__.safeDivisionInt(p1, p1))
            d_535_v22_: _dafny.MultiSet
            d_535_v22_ = _dafny.MultiSet([d_511_v2_, False])
            d_536_v23_: _dafny.MultiSet
            d_536_v23_ = _dafny.MultiSet([p1, p1, (d_534_v21_)[default__.safeIndex(856, (d_534_v21_).length(0))]])
            d_537_v24_: C0
            nw88_ = C0()
            nw88_.ctor__(_dafny.CodePoint('v'))
            d_537_v24_ = nw88_
            d_538_v25_: _dafny.Map
            d_538_v25_ = _dafny.Map({d_537_v24_: d_511_v2_})
            d_539_v26_: _dafny.Map
            d_539_v26_ = _dafny.Map({((d_538_v25_)[d_537_v24_] if (d_537_v24_) in (d_538_v25_) else default__.fm2(d_511_v2_, p1, d_511_v2_, d_511_v2_, globalState)): True})
            d_540_v27_: D0
            d_540_v27_ = D0_DC0(d_511_v2_)
            d_541_v28_: _dafny.Array
            nw89_ = _dafny.Array(None, 13)
            nw89_[int(0)] = (d_540_v27_).cf0
            nw89_[int(1)] = True
            nw89_[int(2)] = False
            nw89_[int(3)] = d_511_v2_
            nw89_[int(4)] = (self).fm5(globalState)
            nw89_[int(5)] = d_511_v2_
            nw89_[int(6)] = d_511_v2_
            nw89_[int(7)] = d_511_v2_
            nw89_[int(8)] = d_511_v2_
            nw89_[int(9)] = d_511_v2_
            nw89_[int(10)] = d_511_v2_
            nw89_[int(11)] = d_511_v2_
            nw89_[int(12)] = d_511_v2_
            d_541_v28_ = nw89_
            d_542_v29_: _dafny.MultiSet
            d_542_v29_ = _dafny.MultiSet([d_541_v28_])
            d_543_v30_: _dafny.Array
            nw90_ = _dafny.Array(None, 26)
            nw90_[int(0)] = (d_535_v22_).isdisjoint(_dafny.MultiSet(d_530_v17_))
            nw90_[int(1)] = (d_511_v2_ if d_511_v2_ else d_511_v2_)
            nw90_[int(2)] = (False) == (d_511_v2_)
            nw90_[int(3)] = d_511_v2_
            nw90_[int(4)] = (d_530_v17_)[default__.safeIndex((d_534_v21_)[default__.safeIndex(856, (d_534_v21_).length(0))], len(d_530_v17_))]
            nw90_[int(5)] = not (d_511_v2_) or (d_511_v2_)
            nw90_[int(6)] = d_511_v2_
            nw90_[int(7)] = d_511_v2_
            nw90_[int(8)] = (d_530_v17_)[default__.safeIndex(p1, len(d_530_v17_))]
            nw90_[int(9)] = d_511_v2_
            nw90_[int(10)] = d_511_v2_
            nw90_[int(11)] = (d_536_v23_).isdisjoint(d_536_v23_)
            nw90_[int(12)] = (d_530_v17_)[default__.safeIndex((self).f0, len(d_530_v17_))]
            nw90_[int(13)] = ((d_530_v17_)[default__.safeIndex((d_534_v21_)[default__.safeIndex(856, (d_534_v21_).length(0))], len(d_530_v17_))]) or (((d_539_v26_)[d_511_v2_] if (d_511_v2_) in (d_539_v26_) else False))
            nw90_[int(14)] = (d_511_v2_) == (d_511_v2_)
            nw90_[int(15)] = True
            nw90_[int(16)] = (_dafny.MultiSet([d_541_v28_])).ispropersubset(d_542_v29_)
            nw90_[int(17)] = d_511_v2_
            nw90_[int(18)] = (self).fm5(globalState)
            nw90_[int(19)] = (d_509_v0_).ispropersubset(d_509_v0_)
            nw90_[int(20)] = d_511_v2_
            nw90_[int(21)] = d_511_v2_
            nw90_[int(22)] = d_511_v2_
            nw90_[int(23)] = d_511_v2_
            nw90_[int(24)] = not((self).fm5(globalState))
            nw90_[int(25)] = d_511_v2_
            d_543_v30_ = nw90_
            d_544_v31_: _dafny.Seq
            d_544_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykehlcy"))
            index86_ = default__.safeIndex(781, (d_541_v28_).length(0))
            (d_541_v28_)[index86_] = (d_515_v6_) not in (d_544_v31_)
            d_545_v32_: _dafny.Seq
            d_545_v32_ = _dafny.SeqWithoutIsStrInference([d_534_v21_, d_534_v21_, d_534_v21_])
            d_546_v33_: D2
            d_546_v33_ = D2_DC4(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_547_i0_ in range(default__.abs(382))]), d_511_v2_)
            index87_ = default__.safeIndex(781, (d_541_v28_).length(0))
            index88_ = default__.safeIndex(856, (d_534_v21_).length(0))
            rhs83_ = not(not (d_511_v2_) or ((d_534_v21_) not in (d_545_v32_)))
            rhs84_ = default__.fm3(d_514_v5_, ((d_537_v24_).f3 if d_511_v2_ else (d_537_v24_).f3), p1, p1, globalState)
            rhs85_ = (len((d_546_v33_).cf4)) <= ((self).f0)
            rhs86_ = d_511_v2_
            rhs87_ = (0) - (len(d_530_v17_))
            lhs24_ = d_541_v28_
            lhs25_ = default__.safeIndex(781, (d_541_v28_).length(0))
            lhs26_ = d_534_v21_
            lhs27_ = default__.safeIndex(856, (d_534_v21_).length(0))
            d_511_v2_ = rhs83_
            r2 = rhs84_
            lhs24_[lhs25_] = rhs85_
            d_511_v2_ = rhs86_
            lhs26_[lhs27_] = rhs87_
            nw91_ = _dafny.Array(int(0), 18)
            d_534_v21_ = nw91_
            d_548_v34_: C0
            nw92_ = C0()
            nw92_.ctor__((d_537_v24_).f3)
            d_548_v34_ = nw92_
        elif source4_.is_DC3:
            d_549___mcc_h2_ = source4_.cf3
            d_550_cf3_ = d_549___mcc_h2_
            d_551_v35_: _dafny.Array
            def lambda21_(d_552_p1_):
                def lambda22_(d_553_i1_):
                    return (d_553_i1_) - (d_552_p1_)

                return lambda22_

            init10_ = lambda21_(p1)
            nw93_ = _dafny.Array(None, 5)
            for i0_10_ in range(nw93_.length(0)):
                nw93_[i0_10_] = init10_(i0_10_)
            d_551_v35_ = nw93_
            index89_ = default__.safeIndex(993, (d_551_v35_).length(0))
            (d_551_v35_)[index89_] = default__.safeModuloInt(p1, p1)
            d_554_v36_: _dafny.MultiSet
            d_554_v36_ = _dafny.MultiSet([d_551_v35_])
            d_555_v37_: D6
            d_555_v37_ = D6_DC14(d_551_v35_)
            d_556_v38_: C0
            nw94_ = C0()
            nw94_.ctor__(_dafny.CodePoint('a'))
            d_556_v38_ = nw94_
            d_557_v39_: _dafny.Map
            d_557_v39_ = _dafny.Map({(self).f0: d_556_v38_})
            d_558_v40_: _dafny.Map
            d_558_v40_ = _dafny.Map({d_511_v2_: p1})
            d_559_v41_: _dafny.Seq
            d_559_v41_ = _dafny.SeqWithoutIsStrInference([d_556_v38_, d_556_v38_])
            d_560_v42_: _dafny.Array
            nw95_ = _dafny.Array(None, 26)
            nw95_[int(0)] = d_556_v38_
            nw95_[int(1)] = d_556_v38_
            nw95_[int(2)] = d_556_v38_
            nw95_[int(3)] = d_556_v38_
            nw95_[int(4)] = ((d_557_v39_)[((d_558_v40_)[d_511_v2_] if (d_511_v2_) in (d_558_v40_) else p1)] if (((d_558_v40_)[d_511_v2_] if (d_511_v2_) in (d_558_v40_) else p1)) in (d_557_v39_) else d_556_v38_)
            nw95_[int(5)] = (d_559_v41_)[default__.safeIndex(p1, len(d_559_v41_))]
            nw95_[int(6)] = d_556_v38_
            nw95_[int(7)] = d_556_v38_
            nw95_[int(8)] = d_556_v38_
            nw95_[int(9)] = d_556_v38_
            nw95_[int(10)] = d_556_v38_
            nw95_[int(11)] = d_556_v38_
            nw95_[int(12)] = d_556_v38_
            nw95_[int(13)] = d_556_v38_
            nw95_[int(14)] = d_556_v38_
            nw95_[int(15)] = d_556_v38_
            nw95_[int(16)] = d_556_v38_
            nw95_[int(17)] = d_556_v38_
            nw95_[int(18)] = d_556_v38_
            nw95_[int(19)] = d_556_v38_
            nw95_[int(20)] = d_556_v38_
            nw95_[int(21)] = d_556_v38_
            nw95_[int(22)] = d_556_v38_
            nw95_[int(23)] = d_556_v38_
            nw95_[int(24)] = d_556_v38_
            nw95_[int(25)] = d_556_v38_
            d_560_v42_ = nw95_
            d_561_v43_: D3
            d_561_v43_ = D3_DC9((self).f0, d_560_v42_, d_511_v2_)
            pat_let_tv12_ = d_551_v35_
            pat_let_tv13_ = d_511_v2_
            index90_ = default__.safeIndex(993, (d_551_v35_).length(0))
            def iife30_(_pat_let6_0):
                def iife31_(d_562_dt__update__tmp_h0_):
                    def iife32_(_pat_let7_0):
                        def iife33_(d_563_dt__update_hcf22_h0_):
                            return D6_DC14(d_563_dt__update_hcf22_h0_)
                        return iife33_(_pat_let7_0)
                    return iife32_(pat_let_tv12_)
                return iife31_(_pat_let6_0)
            def iife34_(_pat_let8_0):
                def iife35_(d_564_dt__update__tmp_h1_):
                    def iife36_(_pat_let9_0):
                        def iife37_(d_565_dt__update_hcf10_h0_):
                            def iife38_(_pat_let10_0):
                                def iife39_(d_566_dt__update_hcf12_h0_):
                                    return D3_DC9(d_565_dt__update_hcf10_h0_, (d_564_dt__update__tmp_h1_).cf11, d_566_dt__update_hcf12_h0_)
                                return iife39_(_pat_let10_0)
                            return iife38_(pat_let_tv13_)
                        return iife37_(_pat_let9_0)
                    return iife36_((self).f0)
                return iife35_(_pat_let8_0)
            rhs88_ = (d_554_v36_).ispropersubset((d_554_v36_).intersection(_dafny.MultiSet([(iife30_(d_555_v37_)).cf22, d_551_v35_, d_551_v35_, d_551_v35_])))
            rhs89_ = ((self).f0) * ((iife34_(d_561_v43_)).cf10)
            lhs28_ = d_551_v35_
            lhs29_ = default__.safeIndex(993, (d_551_v35_).length(0))
            d_511_v2_ = rhs88_
            lhs28_[lhs29_] = rhs89_
            d_567_v44_: _dafny.Seq
            d_567_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arknopxms"))
            d_568_v45_: D2
            d_568_v45_ = D2_DC4(d_567_v44_, d_511_v2_)
            d_569_v46_: _dafny.Set
            d_569_v46_ = _dafny.Set({d_568_v45_, d_568_v45_, d_568_v45_})
            if (_dafny.Set({d_568_v45_})).isdisjoint(d_569_v46_):
                d_511_v2_ = d_511_v2_
                d_511_v2_ = (d_511_v2_ if ((self).f0) != ((self).f0) else d_511_v2_)
                d_511_v2_ = not(default__.fm2(d_511_v2_, ((self).f0) * ((self).f0), True, False, globalState))
                d_570_v47_: _dafny.Map
                d_570_v47_ = _dafny.Map({len(default__.fm18((d_551_v35_)[default__.safeIndex(993, (d_551_v35_).length(0))], globalState)): (self).f0})
                d_571_v48_: T0
                nw96_ = C1()
                nw96_.ctor__(((d_570_v47_)[p1] if (p1) in (d_570_v47_) else default__.fm3(d_514_v5_, d_515_v6_, (self).f0, p1, globalState)))
                d_571_v48_ = nw96_
                index91_ = default__.safeIndex(993, (d_551_v35_).length(0))
                rhs90_ = d_571_v48_
                rhs91_ = len(_dafny.SeqWithoutIsStrInference([(_dafny.Set({d_511_v2_})).intersection(_dafny.Set({(self).fm6(d_511_v2_, d_567_v44_, d_567_v44_, (d_571_v48_).f0, globalState), d_511_v2_})), d_514_v5_, d_514_v5_]))
                lhs30_ = d_551_v35_
                lhs31_ = default__.safeIndex(993, (d_551_v35_).length(0))
                d_571_v48_ = rhs90_
                lhs30_[lhs31_] = rhs91_
                d_572_v49_: _dafny.Array
                nw97_ = _dafny.Array(_dafny.Seq({}), 18)
                d_572_v49_ = nw97_
                d_572_v49_ = d_572_v49_
            elif True:
                d_573_v50_: C0
                nw98_ = C0()
                nw98_.ctor__((d_556_v38_).f3)
                d_573_v50_ = nw98_
                d_574_v51_: _dafny.Array
                nw99_ = _dafny.Array(False, 11)
                d_574_v51_ = nw99_
                index92_ = default__.safeIndex(586, (d_574_v51_).length(0))
                (d_574_v51_)[index92_] = d_511_v2_
                index93_ = default__.safeIndex(586, (d_574_v51_).length(0))
                (d_574_v51_)[index93_] = d_511_v2_
                d_511_v2_ = (d_574_v51_)[default__.safeIndex(586, (d_574_v51_).length(0))]
                d_568_v45_ = D2_DC4(d_567_v44_, d_511_v2_)
                r1 = ((d_551_v35_)[default__.safeIndex(993, (d_551_v35_).length(0))]) * ((d_551_v35_)[default__.safeIndex(993, (d_551_v35_).length(0))])
            d_575_v52_: D0
            d_575_v52_ = D0_DC1(d_511_v2_)
            d_576_v53_: _dafny.Seq
            d_576_v53_ = _dafny.SeqWithoutIsStrInference([(self).f0])
            pat_let_tv14_ = d_511_v2_
            index94_ = default__.safeIndex(993, (d_551_v35_).length(0))
            def iife40_(_pat_let11_0):
                def iife41_(d_577_dt__update__tmp_h2_):
                    def iife42_(_pat_let12_0):
                        def iife43_(d_578_dt__update_hcf1_h0_):
                            return D0_DC1(d_578_dt__update_hcf1_h0_)
                        return iife43_(_pat_let12_0)
                    return iife42_(pat_let_tv14_)
                return iife41_(_pat_let11_0)
            rhs92_ = (d_567_v44_)[default__.safeIndex((d_551_v35_)[default__.safeIndex(993, (d_551_v35_).length(0))], len(d_567_v44_))]
            rhs93_ = (iife40_(d_575_v52_)).cf1
            rhs94_ = len(d_576_v53_)
            lhs32_ = d_551_v35_
            lhs33_ = default__.safeIndex(993, (d_551_v35_).length(0))
            d_515_v6_ = rhs92_
            d_511_v2_ = rhs93_
            lhs32_[lhs33_] = rhs94_
            index95_ = default__.safeIndex(993, (d_551_v35_).length(0))
            (d_551_v35_)[index95_] = 193
        elif True:
            d_579___mcc_h3_ = source4_.cf6
            d_580_cf6_ = d_579___mcc_h3_
            d_581_v54_: C0
            nw100_ = C0()
            nw100_.ctor__(d_515_v6_)
            d_581_v54_ = nw100_
            d_582_v55_: C1
            nw101_ = C1()
            nw101_.ctor__(len(_dafny.Map({p1: d_581_v54_})))
            d_582_v55_ = nw101_
            r0 = (0) - ((self).f0)
            d_583_v56_: C0
            nw102_ = C0()
            nw102_.ctor__((d_581_v54_).f3)
            d_583_v56_ = nw102_
            r1 = (self).f0
        d_584_i2_: int
        d_584_i2_ = 0
        with _dafny.label("6"):
            while d_511_v2_:
                with _dafny.c_label("6"):
                    if (d_584_i2_) >= (100):
                        raise _dafny.Break("6")
                    d_584_i2_ = (d_584_i2_) + (1)
                    d_585_v57_: _dafny.Array
                    nw103_ = _dafny.Array(None, 16)
                    d_585_v57_ = nw103_
                    d_586_v58_: D3
                    d_586_v58_ = D3_DC9(default__.safeModuloInt((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffshwyvs"))])).cardinality, 364), d_585_v57_, d_511_v2_)
                    source5_ = d_586_v58_
                    if source5_.is_DC8:
                        d_587___mcc_h4_ = source5_.cf8
                        d_588___mcc_h5_ = source5_.cf9
                        d_589_cf9_ = d_588___mcc_h5_
                        d_590_cf8_ = d_587___mcc_h4_
                        d_591_v59_: _dafny.Seq
                        d_591_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpgrnxwqu"))
                        d_592_v60_: bool
                        out14_: bool
                        out14_ = default__.m0((self).fm6(d_511_v2_, d_591_v59_, d_591_v59_, default__.fm3(d_514_v5_, d_515_v6_, len(d_514_v5_), (self).f0, globalState), globalState), p1, globalState)
                        d_592_v60_ = out14_
                        r1 = d_589_cf9_
                        d_593_v61_: C1
                        nw104_ = C1()
                        nw104_.ctor__((self).f0)
                        d_593_v61_ = nw104_
                        d_594_v62_: _dafny.Map
                        d_594_v62_ = _dafny.Map({d_592_v60_: d_592_v60_})
                        d_594_v62_ = (d_594_v62_).set(d_511_v2_, d_511_v2_)
                    elif source5_.is_DC9:
                        d_595___mcc_h6_ = source5_.cf10
                        d_596___mcc_h7_ = source5_.cf11
                        d_597___mcc_h8_ = source5_.cf12
                        d_598_cf12_ = d_597___mcc_h8_
                        d_599_cf11_ = d_596___mcc_h7_
                        d_600_cf10_ = d_595___mcc_h6_
                        d_601_v63_: T1
                        nw105_ = C2()
                        nw105_.ctor__(p1)
                        d_601_v63_ = nw105_
                        d_602_v64_: _dafny.Map
                        d_602_v64_ = _dafny.Map({((self).f0) <= (p1): d_601_v63_})
                        d_602_v64_ = (d_602_v64_).set(False, ((d_602_v64_)[(d_601_v63_).fm5(globalState)] if ((d_601_v63_).fm5(globalState)) in (d_602_v64_) else d_601_v63_))
                        d_603_v65_: _dafny.Array
                        def lambda23_(d_604_v0_):
                            def lambda24_(d_605_i3_):
                                return d_604_v0_

                            return lambda24_

                        init11_ = lambda23_(d_509_v0_)
                        nw106_ = _dafny.Array(None, 19)
                        for i0_11_ in range(nw106_.length(0)):
                            nw106_[i0_11_] = init11_(i0_11_)
                        d_603_v65_ = nw106_
                        d_603_v65_ = d_603_v65_
                        d_606_v66_: _dafny.Array
                        nw107_ = _dafny.Array(False, 26)
                        d_606_v66_ = nw107_
                        d_607_v67_: _dafny.MultiSet
                        d_607_v67_ = _dafny.MultiSet([(d_601_v63_).f0, d_600_cf10_])
                        index96_ = default__.safeIndex(423, (d_606_v66_).length(0))
                        (d_606_v66_)[index96_] = (((d_607_v67_)[d_600_cf10_] if (d_600_cf10_) in (d_607_v67_) else (self).f0)) < ((self).f0)
                        index97_ = default__.safeIndex(423, (d_606_v66_).length(0))
                        (d_606_v66_)[index97_] = True
                        r2 = (0) - ((p1) * (d_600_cf10_))
                    elif True:
                        d_608___mcc_h9_ = source5_.cf7
                        d_609_cf7_ = d_608___mcc_h9_
                        d_610_v68_: C0
                        nw108_ = C0()
                        nw108_.ctor__(d_515_v6_)
                        d_610_v68_ = nw108_
                        d_511_v2_ = d_511_v2_
                        r0 = (p1) + ((self).f0)
                        d_611_v69_: _dafny.Array
                        def lambda25_(d_612_cf7_):
                            def lambda26_(d_613_i4_):
                                return (d_613_i4_) * (len(d_612_cf7_))

                            return lambda26_

                        init12_ = lambda25_(d_609_cf7_)
                        nw109_ = _dafny.Array(None, 7)
                        for i0_12_ in range(nw109_.length(0)):
                            nw109_[i0_12_] = init12_(i0_12_)
                        d_611_v69_ = nw109_
                        d_614_v70_: _dafny.MultiSet
                        d_614_v70_ = _dafny.MultiSet([d_611_v69_, d_611_v69_, d_611_v69_, d_611_v69_, d_611_v69_])
                        d_614_v70_ = _dafny.MultiSet([d_611_v69_, d_611_v69_])
                    d_615_v71_: _dafny.Array
                    def lambda27_(d_616_p1_):
                        def lambda28_(d_617_i5_):
                            return (d_617_i5_) * (d_616_p1_)

                        return lambda28_

                    init13_ = lambda27_(p1)
                    nw110_ = _dafny.Array(None, 5)
                    for i0_13_ in range(nw110_.length(0)):
                        nw110_[i0_13_] = init13_(i0_13_)
                    d_615_v71_ = nw110_
                    d_618_v72_: _dafny.Map
                    d_618_v72_ = _dafny.Map({p1: d_615_v71_})
                    d_618_v72_ = d_618_v72_
                    d_619_v73_: _dafny.Map
                    d_619_v73_ = _dafny.Map({(self).f0: p1})
                    d_620_v74_: _dafny.Map
                    d_620_v74_ = _dafny.Map({d_515_v6_: p1})
                    d_621_v75_: _dafny.Array
                    nw111_ = _dafny.Array(None, 10)
                    nw111_[int(0)] = d_511_v2_
                    nw111_[int(1)] = d_511_v2_
                    nw111_[int(2)] = ((d_619_v73_).set(len(d_620_v74_), (self).f0)) != (d_619_v73_)
                    nw111_[int(3)] = True
                    nw111_[int(4)] = d_511_v2_
                    nw111_[int(5)] = True
                    nw111_[int(6)] = True
                    nw111_[int(7)] = False
                    nw111_[int(8)] = d_511_v2_
                    nw111_[int(9)] = d_511_v2_
                    d_621_v75_ = nw111_
                    d_622_v76_: _dafny.Seq
                    d_622_v76_ = _dafny.SeqWithoutIsStrInference([(self).fm13((self).f0, False, globalState), d_511_v2_])
                    index98_ = default__.safeIndex(972, (d_621_v75_).length(0))
                    (d_621_v75_)[index98_] = (d_622_v76_)[default__.safeIndex((0) - (len(d_619_v73_)), len(d_622_v76_))]
                    index99_ = default__.safeIndex(635, (d_621_v75_).length(0))
                    (d_621_v75_)[index99_] = (d_511_v2_) == (d_511_v2_)
                    d_623_v77_: _dafny.MultiSet
                    d_623_v77_ = _dafny.MultiSet([d_511_v2_, True])
                    d_624_v78_: _dafny.Seq
                    d_624_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvejd"))
                    d_625_v79_: _dafny.MultiSet
                    d_625_v79_ = _dafny.MultiSet([204, (self).f0])
                    d_626_v80_: _dafny.MultiSet
                    d_626_v80_ = _dafny.MultiSet([d_623_v77_, d_623_v77_])
                    index100_ = default__.safeIndex(972, (d_621_v75_).length(0))
                    index101_ = default__.safeIndex(635, (d_621_v75_).length(0))
                    rhs95_ = (((d_623_v77_) | (d_623_v77_)).cardinality) != (p1)
                    rhs96_ = p1
                    rhs97_ = d_586_v58_
                    rhs98_ = len(((d_514_v5_) | (_dafny.Set({d_511_v2_}))) | (d_514_v5_))
                    rhs99_ = (len((d_624_v78_) + (default__.fm19(d_625_v79_, (self).f0, (self).f0, (self).f0, globalState)))) == (default__.fm3(d_514_v5_, d_515_v6_, (self).f0, (d_626_v80_).cardinality, globalState))
                    lhs34_ = d_621_v75_
                    lhs35_ = default__.safeIndex(972, (d_621_v75_).length(0))
                    lhs36_ = d_621_v75_
                    lhs37_ = default__.safeIndex(635, (d_621_v75_).length(0))
                    lhs34_[lhs35_] = rhs95_
                    r0 = rhs96_
                    d_586_v58_ = rhs97_
                    r0 = rhs98_
                    lhs36_[lhs37_] = rhs99_
                    d_627_v81_: _dafny.Array
                    nw112_ = _dafny.Array(None, 5)
                    nw112_[int(0)] = d_515_v6_
                    nw112_[int(1)] = (d_515_v6_ if d_511_v2_ else d_515_v6_)
                    nw112_[int(2)] = d_515_v6_
                    nw112_[int(3)] = d_515_v6_
                    nw112_[int(4)] = d_515_v6_
                    d_627_v81_ = nw112_
                    index102_ = default__.safeIndex(132, (d_627_v81_).length(0))
                    (d_627_v81_)[index102_] = _dafny.CodePoint('b')
                    index103_ = default__.safeIndex(132, (d_627_v81_).length(0))
                    (d_627_v81_)[index103_] = d_515_v6_
                    pass
            pass
        d_628_i6_: int
        d_628_i6_ = 0
        with _dafny.label("7"):
            while (self).fm13((0) - ((default__.fm3(d_514_v5_, d_515_v6_, p1, p1, globalState)) - (p1)), True, globalState):
                with _dafny.c_label("7"):
                    if (d_628_i6_) >= (100):
                        raise _dafny.Break("7")
                    d_628_i6_ = (d_628_i6_) + (1)
                    d_511_v2_ = not(d_511_v2_)
                    d_629_v82_: _dafny.Map
                    d_629_v82_ = _dafny.Map({d_511_v2_: d_514_v5_})
                    d_511_v2_ = (default__.safeDivisionInt((self).f0, (D5_DC13(678, ((d_629_v82_)[d_511_v2_] if (d_511_v2_) in (d_629_v82_) else d_514_v5_), d_515_v6_, (self).f0, False)).cf17)) <= ((self).f0)
                    d_511_v2_ = d_511_v2_
                    d_630_v83_: _dafny.Seq
                    d_630_v83_ = _dafny.SeqWithoutIsStrInference([d_511_v2_, d_511_v2_, True, not(d_511_v2_), True])
                    d_631_v84_: _dafny.Seq
                    d_631_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oiil"))
                    d_632_v85_: _dafny.Seq
                    d_632_v85_ = _dafny.SeqWithoutIsStrInference([len(d_631_v84_), (0) - (p1)])
                    d_633_v86_: _dafny.MultiSet
                    d_633_v86_ = _dafny.MultiSet([p1, p1])
                    source6_ = default__.fm23((d_630_v83_)[default__.safeIndex((0) - ((self).f0), len(d_630_v83_))], (0) - ((len(d_514_v5_)) * (p1)), (0) - ((d_632_v85_)[default__.safeIndex((self).f0, len(d_632_v85_))]), (default__.fm19(d_633_v86_, (0) - ((self).f0), (d_633_v86_).cardinality, len(d_631_v84_), globalState)) + (d_631_v84_), globalState)
                    if source6_.is_DC12:
                        d_634___mcc_h10_ = source6_.cf15
                        d_635___mcc_h11_ = source6_.cf16
                        d_636_cf16_ = d_635___mcc_h11_
                        d_637_cf15_ = d_634___mcc_h10_
                        d_638_v88_: C2
                        nw113_ = C2()
                        def iife44_():
                            coll18_ = _dafny.Set()
                            compr_18_: int
                            for compr_18_ in _dafny.IntegerRange(908, 602):
                                d_639_v87_: int = compr_18_
                                if ((908) <= (d_639_v87_)) and ((d_639_v87_) < (602)):
                                    coll18_ = coll18_.union(_dafny.Set([(d_639_v87_) - ((0) - (p1))]))
                            return _dafny.Set(coll18_)
                        nw113_.ctor__(len((_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvaepprp"))), 913})).intersection(iife44_()
                        )))
                        d_638_v88_ = nw113_
                        r0 = (d_632_v85_)[default__.safeIndex((self).f0, len(d_632_v85_))]
                        d_511_v2_ = d_511_v2_
                        d_511_v2_ = d_511_v2_
                    elif source6_.is_DC13:
                        d_640___mcc_h12_ = source6_.cf17
                        d_641___mcc_h13_ = source6_.cf18
                        d_642___mcc_h14_ = source6_.cf19
                        d_643___mcc_h15_ = source6_.cf20
                        d_644___mcc_h16_ = source6_.cf21
                        d_645_cf21_ = d_644___mcc_h16_
                        d_646_cf20_ = d_643___mcc_h15_
                        d_647_cf19_ = d_642___mcc_h14_
                        d_648_cf18_ = d_641___mcc_h13_
                        d_649_cf17_ = d_640___mcc_h12_
                        r0 = (self).f0
                        d_646_cf20_ = -224
                        d_645_cf21_ = (len(_dafny.SeqWithoutIsStrInference([(d_633_v86_).cardinality for d_650_i7_ in range(default__.abs(165))]))) > ((self).f0)
                        d_645_cf21_ = d_511_v2_
                    elif True:
                        d_651___mcc_h17_ = source6_.cf14
                        d_652_cf14_ = d_651___mcc_h17_
                        r1 = -301
                        d_653_v89_: D3
                        d_653_v89_ = D3_DC8(d_511_v2_, p1)
                        d_654_v90_: _dafny.Map
                        d_654_v90_ = _dafny.Map({d_511_v2_: (d_653_v89_).cf9})
                        d_655_v91_: _dafny.MultiSet
                        d_655_v91_ = _dafny.MultiSet([d_654_v90_, d_654_v90_, d_654_v90_])
                        d_656_v92_: _dafny.MultiSet
                        d_656_v92_ = _dafny.MultiSet([d_514_v5_])
                        rhs100_ = (_dafny.MultiSet([d_514_v5_, d_514_v5_])) != (d_656_v92_)
                        rhs101_ = d_511_v2_
                        rhs102_ = not((252) >= (p1))
                        rhs103_ = ((d_655_v91_).intersection(_dafny.MultiSet([d_654_v90_]))) - (d_655_v91_)
                        rhs104_ = (default__.safeModuloInt((self).f0, (self).f0)) * ((self).f0)
                        d_511_v2_ = rhs100_
                        d_511_v2_ = rhs101_
                        d_511_v2_ = rhs102_
                        d_655_v91_ = rhs103_
                        r2 = rhs104_
                        d_511_v2_ = d_511_v2_
                        d_657_v93_: _dafny.Array
                        nw114_ = _dafny.Array(int(0), 27)
                        d_657_v93_ = nw114_
                        index104_ = default__.safeIndex(460, (d_657_v93_).length(0))
                        (d_657_v93_)[index104_] = p1
                        index105_ = default__.safeIndex(460, (d_657_v93_).length(0))
                        (d_657_v93_)[index105_] = (0) - ((self).f0)
                    pass
            pass
        d_658_i8_: int
        d_658_i8_ = 0
        with _dafny.label("8"):
            while (default__.fm2(True, p1, d_511_v2_, d_511_v2_, globalState)) or (d_511_v2_):
                with _dafny.c_label("8"):
                    if (d_658_i8_) >= (100):
                        raise _dafny.Break("8")
                    d_658_i8_ = (d_658_i8_) + (1)
                    d_659_v94_: _dafny.Array
                    def lambda29_(d_660_v2_):
                        def lambda30_(d_661_i9_):
                            return (D7_DC17(d_660_v2_, d_660_v2_, d_660_v2_) if d_660_v2_ else D7_DC17(d_660_v2_, d_660_v2_, d_660_v2_))

                        return lambda30_

                    init14_ = lambda29_(d_511_v2_)
                    nw115_ = _dafny.Array(None, 28)
                    for i0_14_ in range(nw115_.length(0)):
                        nw115_[i0_14_] = init14_(i0_14_)
                    d_659_v94_ = nw115_
                    nw116_ = _dafny.Array(D7.default()(), 13)
                    d_659_v94_ = nw116_
                    d_662_v95_: _dafny.Array
                    nw117_ = _dafny.Array(_dafny.Array(None, 0), 16)
                    d_662_v95_ = nw117_
                    d_663_v96_: _dafny.Array
                    nw118_ = _dafny.Array(_dafny.Map({}), 2)
                    d_663_v96_ = nw118_
                    index106_ = default__.safeIndex(208, (d_662_v95_).length(0))
                    (d_662_v95_)[index106_] = d_663_v96_
                    d_664_v97_: C1
                    nw119_ = C1()
                    nw119_.ctor__(p1)
                    d_664_v97_ = nw119_
                    d_665_v98_: _dafny.Map
                    d_665_v98_ = _dafny.Map({d_664_v97_: d_663_v96_})
                    d_666_v99_: _dafny.Array
                    d_666_v99_ = ((d_665_v98_)[d_664_v97_] if (d_664_v97_) in (d_665_v98_) else d_663_v96_)
                    index107_ = default__.safeIndex(208, (d_662_v95_).length(0))
                    (d_662_v95_)[index107_] = (d_666_v99_)
                    r0 = (0) - (-892)
                    d_667_v100_: _dafny.Map
                    d_667_v100_ = _dafny.Map({default__.fm3(d_514_v5_, d_515_v6_, (self).f0, p1, globalState): _dafny.SeqWithoutIsStrInference([d_511_v2_])})
                    d_668_v101_: _dafny.Seq
                    d_668_v101_ = _dafny.SeqWithoutIsStrInference([d_511_v2_, d_511_v2_, d_511_v2_])
                    d_667_v100_ = (d_667_v100_).set(p1, d_668_v101_)
                    pass
            pass
        d_511_v2_ = d_511_v2_
        r0 = (self).f0
        d_669_v102_: _dafny.Seq
        d_669_v102_ = _dafny.SeqWithoutIsStrInference([p1])
        r1 = (0) - ((0) - ((d_669_v102_)[default__.safeIndex(p1, len(d_669_v102_))]))
        r2 = len(d_513_v4_)
        return r0, r1, r2

    def m6(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        d_670_v0_: T0
        nw120_ = C1()
        nw120_.ctor__(p0)
        d_670_v0_ = nw120_
        d_671_v1_: _dafny.Map
        d_671_v1_ = _dafny.Map({d_670_v0_: (d_670_v0_).f0})
        d_671_v1_ = (d_671_v1_).set(d_670_v0_, (self).f0)
        d_672_v2_: bool
        d_672_v2_ = True
        d_673_v3_: _dafny.Seq
        d_673_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejn"))
        d_674_v4_: D2
        d_674_v4_ = D2_DC4(d_673_v3_, (d_670_v0_).fm5(globalState))
        d_675_v5_: D3
        d_675_v5_ = D3_DC8(d_672_v2_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxkw"))) + ((d_674_v4_).cf4)))
        source7_ = d_675_v5_
        if source7_.is_DC8:
            d_676___mcc_h0_ = source7_.cf8
            d_677___mcc_h1_ = source7_.cf9
            d_678_cf9_ = d_677___mcc_h1_
            d_679_cf8_ = d_676___mcc_h0_
            r1 = -144
            d_680_v6_: C1
            nw121_ = C1()
            nw121_.ctor__(default__.safeDivisionInt(276, len(d_673_v3_)))
            d_680_v6_ = nw121_
            d_681_v7_: str
            d_681_v7_ = _dafny.CodePoint('s')
            d_682_v8_: C0
            nw122_ = C0()
            nw122_.ctor__(d_681_v7_)
            d_682_v8_ = nw122_
            d_683_v9_: _dafny.MultiSet
            d_683_v9_ = _dafny.MultiSet([d_672_v2_])
            d_684_v10_: _dafny.Seq
            d_684_v10_ = _dafny.SeqWithoutIsStrInference([False])
            d_685_v11_: _dafny.Seq
            d_685_v11_ = _dafny.SeqWithoutIsStrInference([d_683_v9_, _dafny.MultiSet(d_684_v10_)])
            d_686_v12_: C2
            nw123_ = C2()
            nw123_.ctor__(default__.safeModuloInt(d_678_cf9_, ((d_685_v11_)[default__.safeIndex(d_678_cf9_, len(d_685_v11_))]).cardinality))
            d_686_v12_ = nw123_
        elif source7_.is_DC9:
            d_687___mcc_h2_ = source7_.cf10
            d_688___mcc_h3_ = source7_.cf11
            d_689___mcc_h4_ = source7_.cf12
            d_690_cf12_ = d_689___mcc_h4_
            d_691_cf11_ = d_688___mcc_h3_
            d_692_cf10_ = d_687___mcc_h2_
            d_693_v13_: _dafny.Seq
            d_693_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_670_v0_).f0, (d_670_v0_).f0})])
            d_694_v14_: _dafny.Set
            d_694_v14_ = _dafny.Set({(d_670_v0_).f0})
            d_695_v15_: _dafny.MultiSet
            d_695_v15_ = _dafny.MultiSet([(_dafny.Set({p0})) | ((d_693_v13_)[default__.safeIndex((d_670_v0_).f0, len(d_693_v13_))]), d_694_v14_, d_694_v14_, (d_694_v14_) | (d_694_v14_)])
            d_695_v15_ = d_695_v15_
            d_696_v16_: _dafny.Array
            nw124_ = _dafny.Array(False, 7)
            d_696_v16_ = nw124_
            d_696_v16_ = d_696_v16_
            d_697_v17_: D6
            d_697_v17_ = D6_DC15(934, d_690_cf12_, (d_670_v0_).f0)
            d_697_v17_ = d_697_v17_
            d_673_v3_ = d_673_v3_
        elif True:
            d_698___mcc_h5_ = source7_.cf7
            d_699_cf7_ = d_698___mcc_h5_
            r1 = (default__.safeDivisionInt(p0, p0) if True else (d_670_v0_).f0)
            d_672_v2_ = d_672_v2_
            d_700_v18_: str
            d_700_v18_ = _dafny.CodePoint('d')
            d_700_v18_ = default__.fm17((d_670_v0_).f0, globalState)
            d_701_v19_: _dafny.Array
            nw125_ = _dafny.Array(None, 11)
            nw125_[int(0)] = (d_672_v2_) and (d_672_v2_)
            nw125_[int(1)] = d_672_v2_
            nw125_[int(2)] = d_672_v2_
            nw125_[int(3)] = d_672_v2_
            nw125_[int(4)] = d_672_v2_
            nw125_[int(5)] = False
            nw125_[int(6)] = True
            nw125_[int(7)] = d_672_v2_
            nw125_[int(8)] = d_672_v2_
            nw125_[int(9)] = d_672_v2_
            nw125_[int(10)] = d_672_v2_
            d_701_v19_ = nw125_
            index108_ = default__.safeIndex(62, (d_701_v19_).length(0))
            (d_701_v19_)[index108_] = (self).fm5(globalState)
            index109_ = default__.safeIndex(253, (d_701_v19_).length(0))
            (d_701_v19_)[index109_] = d_672_v2_
            d_702_v20_: _dafny.Map
            d_702_v20_ = _dafny.Map({d_672_v2_: d_672_v2_})
            d_703_v21_: _dafny.Seq
            d_703_v21_ = _dafny.SeqWithoutIsStrInference([d_702_v20_, d_702_v20_, d_702_v20_])
            d_704_v22_: _dafny.Set
            d_704_v22_ = _dafny.Set({d_672_v2_, d_672_v2_, (d_670_v0_).fm5(globalState)})
            d_705_v23_: D2
            d_705_v23_ = D2_DC5()
            d_706_v24_: D2
            d_706_v24_ = D2_DC6(d_705_v23_)
            d_707_v25_: _dafny.MultiSet
            d_707_v25_ = _dafny.MultiSet([d_706_v24_])
            index110_ = default__.safeIndex(62, (d_701_v19_).length(0))
            index111_ = default__.safeIndex(253, (d_701_v19_).length(0))
            rhs105_ = (self).f0
            rhs106_ = default__.fm17(default__.safeModuloInt(p0, (0) - (p0)), globalState)
            rhs107_ = (d_672_v2_ if (d_672_v2_ if d_672_v2_ else d_672_v2_) else d_672_v2_)
            rhs108_ = ((d_707_v25_).set(d_706_v24_, default__.abs(len(d_673_v3_)))).ispropersubset(default__.fm24(d_672_v2_, d_672_v2_, len((d_703_v21_)[default__.safeIndex((d_670_v0_).f0, len(d_703_v21_))]), d_704_v22_, globalState))
            lhs38_ = d_701_v19_
            lhs39_ = default__.safeIndex(62, (d_701_v19_).length(0))
            lhs40_ = d_701_v19_
            lhs41_ = default__.safeIndex(253, (d_701_v19_).length(0))
            r1 = rhs105_
            d_700_v18_ = rhs106_
            lhs38_[lhs39_] = rhs107_
            lhs40_[lhs41_] = rhs108_
        d_708_v26_: _dafny.Array
        nw126_ = _dafny.Array(int(0), 23)
        d_708_v26_ = nw126_
        index112_ = default__.safeIndex(923, (d_708_v26_).length(0))
        (d_708_v26_)[index112_] = p0
        index113_ = default__.safeIndex(923, (d_708_v26_).length(0))
        (d_708_v26_)[index113_] = (p0) + ((self).f0)
        d_709_v27_: _dafny.Map
        d_709_v27_ = _dafny.Map({d_708_v26_: d_708_v26_})
        d_709_v27_ = (d_709_v27_).set(d_708_v26_, d_708_v26_)
        d_710_v28_: _dafny.Array
        def lambda31_(d_711_i0_):
            return True

        init15_ = lambda31_
        nw127_ = _dafny.Array(None, 28)
        for i0_15_ in range(nw127_.length(0)):
            nw127_[i0_15_] = init15_(i0_15_)
        d_710_v28_ = nw127_
        d_710_v28_ = d_710_v28_
        d_712_v29_: _dafny.Set
        d_712_v29_ = _dafny.Set({d_672_v2_, d_672_v2_})
        r0 = (d_712_v29_).issubset(d_712_v29_)
        r0 = d_672_v2_
        d_713_v30_: _dafny.Set
        d_713_v30_ = _dafny.Set({d_708_v26_, d_708_v26_})
        r1 = len(d_713_v30_)
        return r0, r1

    def m7(self, p0, p1, p2, p3, globalState):
        d_714_v0_: int
        d_714_v0_ = -360
        d_714_v0_ = default__.safeModuloInt(d_714_v0_, p3)
        d_715_v1_: _dafny.Map
        d_715_v1_ = _dafny.Map({p1: p2})
        if ((d_715_v1_ if True else _dafny.Map({p1: p2}))) != (d_715_v1_):
            d_716_v2_: _dafny.Map
            d_716_v2_ = _dafny.Map({p1: p1})
            d_714_v0_ = (len((d_716_v2_) | (d_716_v2_))) * ((self).f0)
            d_717_v3_: bool
            d_717_v3_ = False
            d_718_v4_: _dafny.Seq
            d_718_v4_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2, p2, p2])
            d_717_v3_ = (d_718_v4_) == ((d_718_v4_) + (d_718_v4_))
            d_719_v5_: _dafny.Set
            d_719_v5_ = _dafny.Set({p3})
            d_717_v3_ = not((d_717_v3_) or ((d_719_v5_).issubset(d_719_v5_)))
            d_720_v6_: str
            d_720_v6_ = _dafny.CodePoint('p')
            d_721_v7_: _dafny.Seq
            d_721_v7_ = _dafny.SeqWithoutIsStrInference([d_720_v6_, d_720_v6_])
            d_722_v8_: T1
            nw128_ = C2()
            nw128_.ctor__(len(d_719_v5_))
            d_722_v8_ = nw128_
            d_723_v9_: _dafny.Array
            nw129_ = _dafny.Array(None, 20)
            nw129_[int(0)] = d_722_v8_
            nw129_[int(1)] = d_722_v8_
            nw129_[int(2)] = (d_722_v8_ if d_717_v3_ else d_722_v8_)
            nw129_[int(3)] = d_722_v8_
            nw129_[int(4)] = d_722_v8_
            nw129_[int(5)] = d_722_v8_
            nw129_[int(6)] = d_722_v8_
            nw129_[int(7)] = d_722_v8_
            nw129_[int(8)] = d_722_v8_
            nw129_[int(9)] = d_722_v8_
            nw129_[int(10)] = d_722_v8_
            nw129_[int(11)] = d_722_v8_
            nw129_[int(12)] = d_722_v8_
            nw129_[int(13)] = d_722_v8_
            nw129_[int(14)] = d_722_v8_
            nw129_[int(15)] = d_722_v8_
            nw129_[int(16)] = d_722_v8_
            nw129_[int(17)] = d_722_v8_
            nw129_[int(18)] = d_722_v8_
            nw129_[int(19)] = d_722_v8_
            d_723_v9_ = nw129_
            index114_ = default__.safeIndex(256, (d_723_v9_).length(0))
            (d_723_v9_)[index114_] = d_722_v8_
            d_724_v10_: _dafny.Map
            d_724_v10_ = d_716_v2_
            d_725_v11_: _dafny.Seq
            d_725_v11_ = _dafny.SeqWithoutIsStrInference([(d_722_v8_).f0, len(((d_724_v10_)).set(d_717_v3_, p1))])
            index115_ = default__.safeIndex(256, (d_723_v9_).length(0))
            rhs109_ = (d_725_v11_)[default__.safeIndex(default__.safeModuloInt((d_725_v11_)[default__.safeIndex((d_722_v8_).f0, len(d_725_v11_))], (0) - (d_714_v0_)), len(d_725_v11_))]
            rhs110_ = p0
            rhs111_ = d_722_v8_
            rhs112_ = p1
            lhs42_ = d_723_v9_
            lhs43_ = default__.safeIndex(256, (d_723_v9_).length(0))
            d_714_v0_ = rhs109_
            d_721_v7_ = rhs110_
            lhs42_[lhs43_] = rhs111_
            d_717_v3_ = rhs112_
            d_726_v12_: _dafny.Array
            def lambda32_(d_727_v7_, d_728_v3_):
                def lambda33_(d_729_i0_):
                    return D2_DC6(D2_DC4(d_727_v7_, d_728_v3_))

                return lambda33_

            init16_ = lambda32_(d_721_v7_, d_717_v3_)
            nw130_ = _dafny.Array(None, 27)
            for i0_16_ in range(nw130_.length(0)):
                nw130_[i0_16_] = init16_(i0_16_)
            d_726_v12_ = nw130_
            d_730_v13_: D2
            d_730_v13_ = D2_DC3(_dafny.SeqWithoutIsStrInference([D0_DC1(d_717_v3_)]))
            index116_ = default__.safeIndex(364, (d_726_v12_).length(0))
            (d_726_v12_)[index116_] = D2_DC6(d_730_v13_)
            d_731_v14_: D2
            d_731_v14_ = D2_DC6(D2_DC6(d_730_v13_))
            index117_ = default__.safeIndex(364, (d_726_v12_).length(0))
            (d_726_v12_)[index117_] = d_731_v14_
        elif True:
            d_732_v15_: _dafny.Map
            d_732_v15_ = _dafny.Map({p3: (self).f0})
            d_733_v16_: _dafny.Set
            d_733_v16_ = _dafny.Set({p1})
            d_734_v17_: str
            d_734_v17_ = _dafny.CodePoint('b')
            d_735_v18_: _dafny.Map
            d_735_v18_ = _dafny.Map({default__.fm3(d_733_v16_, d_734_v17_, d_714_v0_, 215, globalState): p1})
            d_736_v19_: _dafny.Map
            d_736_v19_ = _dafny.Map({((d_735_v18_)[d_714_v0_] if (d_714_v0_) in (d_735_v18_) else p1): d_732_v15_})
            d_737_v20_: _dafny.MultiSet
            d_737_v20_ = _dafny.MultiSet([(d_732_v15_) | (((d_736_v19_)[p1] if (p1) in (d_736_v19_) else d_732_v15_)), d_732_v15_, d_732_v15_])
            d_738_v21_: bool
            d_738_v21_ = True
            rhs113_ = d_737_v20_
            rhs114_ = p1
            d_737_v20_ = rhs113_
            d_738_v21_ = rhs114_
            d_739_v22_: D5
            d_739_v22_ = D5_DC13(d_714_v0_, d_733_v16_, d_734_v17_, (self).f0, d_738_v21_)
            d_738_v21_ = (d_739_v22_).cf21
            d_740_v23_: _dafny.Map
            d_740_v23_ = _dafny.Map({p1: True})
            d_741_v24_: _dafny.Map
            d_741_v24_ = _dafny.Map({p1: d_740_v23_})
            d_742_v25_: _dafny.Array
            nw131_ = _dafny.Array(None, 6)
            nw131_[int(0)] = d_740_v23_
            nw131_[int(1)] = (d_740_v23_).set(p1, d_738_v21_)
            nw131_[int(2)] = d_740_v23_
            nw131_[int(3)] = d_740_v23_
            nw131_[int(4)] = ((d_741_v24_)[p1] if (p1) in (d_741_v24_) else d_740_v23_)
            nw131_[int(5)] = d_740_v23_
            d_742_v25_ = nw131_
            d_743_v26_: _dafny.Array
            d_743_v26_ = d_742_v25_
            source8_ = d_743_v26_
            d_744___mcc_h0_ = source8_
            d_745_cf31_ = d_744___mcc_h0_
            d_746_v27_: _dafny.Array
            nw132_ = _dafny.Array(False, 2)
            d_746_v27_ = nw132_
            index118_ = default__.safeIndex(132, (d_746_v27_).length(0))
            (d_746_v27_)[index118_] = (True) or (True)
            index119_ = default__.safeIndex(132, (d_746_v27_).length(0))
            rhs115_ = p1
            rhs116_ = (d_733_v16_).issubset((_dafny.Set({not(p1), True, d_738_v21_, False, d_738_v21_})).intersection(d_733_v16_))
            rhs117_ = d_738_v21_
            lhs44_ = d_746_v27_
            lhs45_ = default__.safeIndex(132, (d_746_v27_).length(0))
            d_738_v21_ = rhs115_
            lhs44_[lhs45_] = rhs116_
            d_738_v21_ = rhs117_
            d_738_v21_ = d_738_v21_
            d_747_v28_: _dafny.Map
            d_747_v28_ = _dafny.Map({p1: (0) - ((self).f0)})
            d_748_v29_: _dafny.Map
            d_748_v29_ = _dafny.Map({len(d_747_v28_): d_734_v17_})
            d_749_v30_: _dafny.Map
            d_749_v30_ = _dafny.Map({d_714_v0_: d_746_v27_})
            d_748_v29_ = (d_748_v29_).set((len(d_749_v30_) if d_738_v21_ else d_714_v0_), _dafny.CodePoint('b'))
            d_750_v31_: _dafny.Seq
            d_750_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdbsmm"))
            d_751_v32_: D0
            d_751_v32_ = D0_DC1((d_746_v27_)[default__.safeIndex(132, (d_746_v27_).length(0))])
            d_752_v33_: _dafny.Seq
            d_752_v33_ = _dafny.SeqWithoutIsStrInference([d_751_v32_])
            d_753_v34_: D2
            d_753_v34_ = D2_DC3(d_752_v33_)
            d_754_v35_: _dafny.Seq
            d_754_v35_ = _dafny.SeqWithoutIsStrInference([not(d_738_v21_)])
            d_755_v36_: _dafny.Array
            nw133_ = _dafny.Array(None, 16)
            nw133_[int(0)] = d_746_v27_
            nw133_[int(1)] = d_746_v27_
            nw133_[int(2)] = d_746_v27_
            nw133_[int(3)] = d_746_v27_
            nw133_[int(4)] = d_746_v27_
            nw133_[int(5)] = d_746_v27_
            nw133_[int(6)] = d_746_v27_
            nw133_[int(7)] = d_746_v27_
            nw133_[int(8)] = d_746_v27_
            nw133_[int(9)] = d_746_v27_
            nw133_[int(10)] = d_746_v27_
            nw133_[int(11)] = d_746_v27_
            nw133_[int(12)] = d_746_v27_
            nw133_[int(13)] = d_746_v27_
            nw133_[int(14)] = d_746_v27_
            nw133_[int(15)] = d_746_v27_
            d_755_v36_ = nw133_
            index120_ = default__.safeIndex(563, (d_755_v36_).length(0))
            (d_755_v36_)[index120_] = d_746_v27_
            d_756_v37_: T0
            nw134_ = C1()
            nw134_.ctor__(p3)
            d_756_v37_ = nw134_
            d_757_v38_: D10
            d_757_v38_ = D10_DC21(d_756_v37_)
            d_758_v39_: _dafny.Map
            d_758_v39_ = _dafny.Map({d_738_v21_: (d_757_v38_).cf33})
            d_759_v40_: _dafny.Seq
            d_759_v40_ = _dafny.SeqWithoutIsStrInference([d_758_v39_, d_758_v39_])
            d_760_v41_: _dafny.MultiSet
            d_760_v41_ = _dafny.MultiSet([len((d_759_v40_)[default__.safeIndex(p3, len(d_759_v40_))])])
            index121_ = default__.safeIndex(132, (d_746_v27_).length(0))
            index122_ = default__.safeIndex(563, (d_755_v36_).length(0))
            rhs118_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evawlaen"))).set(default__.safeIndex((self).f0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evawlaen")))), d_734_v17_)
            rhs119_ = not((d_746_v27_)[default__.safeIndex(132, (d_746_v27_).length(0))])
            rhs120_ = d_753_v34_
            rhs121_ = (_dafny.SeqWithoutIsStrInference([p1]) if ((d_760_v41_).cardinality) != ((self).f0) else d_754_v35_)
            rhs122_ = d_746_v27_
            lhs46_ = d_746_v27_
            lhs47_ = default__.safeIndex(132, (d_746_v27_).length(0))
            lhs48_ = d_755_v36_
            lhs49_ = default__.safeIndex(563, (d_755_v36_).length(0))
            d_750_v31_ = rhs118_
            lhs46_[lhs47_] = rhs119_
            d_753_v34_ = rhs120_
            d_754_v35_ = rhs121_
            lhs48_[lhs49_] = rhs122_
            d_761_v42_: _dafny.Array
            nw135_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
            d_761_v42_ = nw135_
            index123_ = default__.safeIndex(303, (d_761_v42_).length(0))
            (d_761_v42_)[index123_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcrj"))
            index124_ = default__.safeIndex(303, (d_761_v42_).length(0))
            (d_761_v42_)[index124_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrrxlikqd"))
            d_762_v43_: _dafny.Seq
            d_762_v43_ = _dafny.SeqWithoutIsStrInference([d_714_v0_])
            d_763_v44_: _dafny.Seq
            d_763_v44_ = _dafny.SeqWithoutIsStrInference([d_738_v21_])
            d_738_v21_ = ((d_762_v43_) != ((d_762_v43_).set(default__.safeIndex((self).f0, len(d_762_v43_)), len(d_763_v44_)))) and (p1)
        d_764_v45_: _dafny.Set
        d_764_v45_ = _dafny.Set({p1})
        d_714_v0_ = default__.safeDivisionInt(p3, len(_dafny.Map({p1: len(d_764_v45_)})))
        d_765_v46_: str
        d_765_v46_ = _dafny.CodePoint('g')
        d_766_v47_: _dafny.Map
        d_766_v47_ = _dafny.Map({p1: p3})
        d_767_v48_: _dafny.Map
        d_767_v48_ = _dafny.Map({d_765_v46_: d_766_v47_})
        d_768_v49_: _dafny.Array
        nw136_ = _dafny.Array(None, 3)
        nw136_[int(0)] = (d_767_v48_) | (d_767_v48_)
        nw136_[int(1)] = d_767_v48_
        nw136_[int(2)] = d_767_v48_
        d_768_v49_ = nw136_
        index125_ = default__.safeIndex(829, (d_768_v49_).length(0))
        (d_768_v49_)[index125_] = d_767_v48_
        index126_ = default__.safeIndex(829, (d_768_v49_).length(0))
        (d_768_v49_)[index126_] = d_767_v48_
        hi5_ = -100
        for d_769_i1_ in range(p3, hi5_):
            d_770_v50_: _dafny.Array
            nw137_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_770_v50_ = nw137_
            d_770_v50_ = d_770_v50_
            d_714_v0_ = d_714_v0_
            d_771_v51_: _dafny.Set
            d_771_v51_ = _dafny.Set({_dafny.Set({(self).f0, (0) - (p3)})})
            d_772_v52_: D5
            d_772_v52_ = D5_DC11(d_771_v51_)
            source9_ = d_772_v52_
            if source9_.is_DC12:
                d_773___mcc_h1_ = source9_.cf15
                d_774___mcc_h2_ = source9_.cf16
                d_775_cf16_ = d_774___mcc_h2_
                d_776_cf15_ = d_773___mcc_h1_
                d_777_v53_: D2
                d_777_v53_ = D2_DC4(p0, False)
                d_778_v54_: _dafny.Map
                d_778_v54_ = _dafny.Map({(d_777_v53_).cf5: p1})
                d_778_v54_ = d_778_v54_
                d_779_v55_: _dafny.MultiSet
                d_779_v55_ = _dafny.MultiSet([d_769_i1_, p3])
                d_780_v56_: bool
                d_781_v57_: int
                out15_: bool
                out16_: int
                out15_, out16_ = (self).m6(((d_779_v55_).set(d_775_cf16_, default__.abs(d_776_cf15_))).cardinality, globalState)
                d_780_v56_ = out15_
                d_781_v57_ = out16_
                d_782_v58_: C0
                nw138_ = C0()
                nw138_.ctor__(d_765_v46_)
                d_782_v58_ = nw138_
                d_780_v56_ = not(False)
            elif source9_.is_DC13:
                d_783___mcc_h3_ = source9_.cf17
                d_784___mcc_h4_ = source9_.cf18
                d_785___mcc_h5_ = source9_.cf19
                d_786___mcc_h6_ = source9_.cf20
                d_787___mcc_h7_ = source9_.cf21
                d_788_cf21_ = d_787___mcc_h7_
                d_789_cf20_ = d_786___mcc_h6_
                d_790_cf19_ = d_785___mcc_h5_
                d_791_cf18_ = d_784___mcc_h4_
                d_792_cf17_ = d_783___mcc_h3_
                d_793_v59_: _dafny.Seq
                d_793_v59_ = _dafny.SeqWithoutIsStrInference([p1, d_788_cf21_])
                d_794_v60_: D7
                d_794_v60_ = D7_DC16(d_793_v59_)
                d_794_v60_ = d_794_v60_
                d_795_v61_: _dafny.MultiSet
                d_795_v61_ = _dafny.MultiSet([(0) - ((0) - (p3)), p3, (0) - (d_789_cf20_)])
                d_796_v62_: _dafny.MultiSet
                d_796_v62_ = _dafny.MultiSet([(d_795_v61_).cardinality, (self).f0, d_714_v0_])
                index127_ = default__.safeIndex(273, (p2).length(0))
                (p2)[index127_] = d_769_i1_
                index128_ = default__.safeIndex(273, (p2).length(0))
                rhs123_ = (d_795_v61_ if p1 else d_796_v62_)
                rhs124_ = 268
                lhs50_ = p2
                lhs51_ = default__.safeIndex(273, (p2).length(0))
                d_796_v62_ = rhs123_
                lhs50_[lhs51_] = rhs124_
                d_714_v0_ = default__.safeDivisionInt(d_769_i1_, d_789_cf20_)
                d_797_v63_: _dafny.Array
                nw139_ = _dafny.Array(False, 25)
                d_797_v63_ = nw139_
                index129_ = default__.safeIndex(381, (d_797_v63_).length(0))
                (d_797_v63_)[index129_] = p1
                index130_ = default__.safeIndex(381, (d_797_v63_).length(0))
                (d_797_v63_)[index130_] = not(d_788_cf21_)
            elif True:
                d_798___mcc_h8_ = source9_.cf14
                d_799_cf14_ = d_798___mcc_h8_
                d_714_v0_ = d_769_i1_
                d_800_v64_: bool
                d_801_v65_: int
                out17_: bool
                out18_: int
                out17_, out18_ = (self).m6(d_714_v0_, globalState)
                d_800_v64_ = out17_
                d_801_v65_ = out18_
                d_801_v65_ = d_769_i1_
                d_802_v66_: _dafny.Map
                d_802_v66_ = _dafny.Map({d_800_v64_: d_800_v64_})
                d_803_v67_: _dafny.Map
                d_803_v67_ = d_802_v66_
                d_803_v67_ = default__.fm25(globalState)
            d_804_v68_: T0
            nw140_ = C1()
            nw140_.ctor__((self).f0)
            d_804_v68_ = nw140_
            d_805_v69_: D10
            d_805_v69_ = D10_DC21(d_804_v68_)
            d_806_v70_: D0
            d_806_v70_ = D0_DC1(p1)
            d_807_v71_: _dafny.Seq
            d_807_v71_ = _dafny.SeqWithoutIsStrInference([d_806_v70_])
            d_808_v72_: _dafny.Map
            d_808_v72_ = _dafny.Map({d_805_v69_: D2_DC3(d_807_v71_)})
            d_809_v73_: _dafny.Seq
            d_809_v73_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({d_805_v69_: D2_DC3(_dafny.SeqWithoutIsStrInference([d_806_v70_]))})) | (d_808_v72_), (d_808_v72_) | (d_808_v72_), (d_808_v72_ if p1 else d_808_v72_), (d_808_v72_) | (d_808_v72_), d_808_v72_])
            d_809_v73_ = d_809_v73_
        d_810_v74_: _dafny.MultiSet
        d_810_v74_ = _dafny.MultiSet([p1, p1, p1, True, p1])
        d_811_v75_: _dafny.Seq
        d_811_v75_ = _dafny.SeqWithoutIsStrInference([((d_810_v74_)[p1] if (p1) in (d_810_v74_) else p3), p3])
        d_812_v76_: D0
        d_812_v76_ = D0_DC1(p1)
        d_813_v77_: _dafny.Seq
        d_813_v77_ = _dafny.SeqWithoutIsStrInference([d_812_v76_])
        d_814_v78_: D2
        d_814_v78_ = D2_DC3(d_813_v77_)
        d_815_v79_: _dafny.Map
        d_815_v79_ = _dafny.Map({p1: d_814_v78_})
        d_816_v80_: D2
        d_816_v80_ = D2_DC6(((d_815_v79_)[p1] if (p1) in (d_815_v79_) else d_814_v78_))
        pat_let_tv15_ = d_811_v75_
        pat_let_tv16_ = d_714_v0_
        pat_let_tv17_ = d_811_v75_
        pat_let_tv18_ = d_811_v75_
        pat_let_tv19_ = d_811_v75_
        pat_let_tv20_ = d_811_v75_
        pat_let_tv21_ = d_714_v0_
        pat_let_tv22_ = d_814_v78_
        def lambda34_(source10_):
            if source10_.is_DC4:
                d_817___mcc_h9_ = source10_.cf4
                d_818___mcc_h10_ = source10_.cf5
                d_819_cf5_ = d_818___mcc_h10_
                d_820_cf4_ = d_817___mcc_h9_
                return pat_let_tv15_
            elif source10_.is_DC5:
                return (_dafny.SeqWithoutIsStrInference([(self).f0, (self).f0, pat_let_tv16_])) + (pat_let_tv17_)
            elif source10_.is_DC3:
                d_821___mcc_h11_ = source10_.cf3
                d_822_cf3_ = d_821___mcc_h11_
                return pat_let_tv18_
            elif True:
                d_823___mcc_h12_ = source10_.cf6
                d_824_cf6_ = d_823___mcc_h12_
                return (pat_let_tv19_).set(default__.safeIndex((self).f0, len(pat_let_tv20_)), pat_let_tv21_)

        def iife45_(_pat_let13_0):
            def iife46_(d_825_dt__update__tmp_h0_):
                def iife47_(_pat_let14_0):
                    def iife48_(d_826_dt__update_hcf6_h0_):
                        return D2_DC6(d_826_dt__update_hcf6_h0_)
                    return iife48_(_pat_let14_0)
                return iife47_(pat_let_tv22_)
            return iife46_(_pat_let13_0)
        d_811_v75_ = lambda34_(iife45_(d_816_v80_))


class C4(T0):
    def  __init__(self):
        self._f0: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f0(self):
        return self._f0
    def ctor__(self, f0):
        (self)._f0 = f0

    def fm5(self, globalState):
        return not(not(not(not((_dafny.CodePoint('i')) in ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrjq"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_827_i0_ in range(default__.abs(549))])))))))

    def fm6(self, p0, p1, p2, p3, globalState):
        return not (True) or ((D3_DC8(True, -634)).cf8)

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        d_828_v0_: bool
        d_828_v0_ = False
        d_829_v1_: _dafny.MultiSet
        d_829_v1_ = _dafny.MultiSet([d_828_v0_])
        r2 = default__.safeDivisionInt((d_829_v1_).cardinality, (p1) * (p1))
        if d_828_v0_:
            d_830_v3_: _dafny.Set
            d_830_v3_ = _dafny.Set({p1, p1})
            d_831_v4_: _dafny.Map
            d_831_v4_ = _dafny.Map({p1: p1})
            d_832_v6_: _dafny.Seq
            d_832_v6_ = _dafny.SeqWithoutIsStrInference([(self).f0])
            d_833_v7_: D3
            d_833_v7_ = D3_DC7(d_832_v6_)
            d_834_v8_: _dafny.Array
            nw141_ = _dafny.Array(int(0), 11)
            d_834_v8_ = nw141_
            d_835_v9_: str
            d_835_v9_ = _dafny.CodePoint('c')
            d_836_v10_: C0
            nw142_ = C0()
            nw142_.ctor__(d_835_v9_)
            d_836_v10_ = nw142_
            d_837_v11_: _dafny.Seq
            d_837_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgt"))
            d_838_v15_: _dafny.Array
            nw143_ = _dafny.Array(None, 29)
            def iife49_():
                coll19_ = _dafny.Map()
                compr_19_: int
                for compr_19_ in (d_830_v3_).Elements:
                    d_839_v2_: int = compr_19_
                    if (d_839_v2_) in (d_830_v3_):
                        coll19_[(d_839_v2_) * ((d_829_v1_).cardinality)] = (self).f0
                return _dafny.Map(coll19_)
            nw143_[int(0)] = iife49_()
            
            def iife50_():
                coll20_ = _dafny.Map()
                compr_20_: int
                for compr_20_ in _dafny.IntegerRange(859, 772):
                    d_840_v5_: int = compr_20_
                    if ((859) <= (d_840_v5_)) and ((d_840_v5_) < (772)):
                        coll20_[(d_840_v5_) + (len(_dafny.Map({d_828_v0_: D2_DC5()})))] = len(d_831_v4_)
                return _dafny.Map(coll20_)
            nw143_[int(1)] = (d_831_v4_) | (iife50_()
            )
            nw143_[int(2)] = d_831_v4_
            nw143_[int(3)] = _dafny.Map({len((d_833_v7_).cf7): (self).f0})
            nw143_[int(4)] = d_831_v4_
            nw143_[int(5)] = d_831_v4_
            nw143_[int(6)] = d_831_v4_
            nw143_[int(7)] = d_831_v4_
            nw143_[int(8)] = (d_831_v4_) | (d_831_v4_)
            nw143_[int(9)] = _dafny.Map({(self).f0: len(_dafny.SeqWithoutIsStrInference([d_834_v8_, d_834_v8_, d_834_v8_]))})
            nw143_[int(10)] = d_831_v4_
            nw143_[int(11)] = d_831_v4_
            nw143_[int(12)] = default__.fm11((self).f0, globalState)
            nw143_[int(13)] = d_831_v4_
            nw143_[int(14)] = _dafny.Map({len((_dafny.SeqWithoutIsStrInference([d_836_v10_])).set(default__.safeIndex((0) - (len(d_830_v3_)), len(_dafny.SeqWithoutIsStrInference([d_836_v10_]))), d_836_v10_)): len(d_837_v11_)})
            nw143_[int(15)] = d_831_v4_
            nw143_[int(16)] = d_831_v4_
            def iife51_():
                coll21_ = _dafny.Map()
                compr_21_: int
                for compr_21_ in _dafny.IntegerRange(896, 834):
                    d_841_v12_: int = compr_21_
                    if ((896) <= (d_841_v12_)) and ((d_841_v12_) < (834)):
                        coll21_[(d_841_v12_) * (-918)] = p1
                return _dafny.Map(coll21_)
            nw143_[int(17)] = (d_831_v4_) | (iife51_()
            )
            nw143_[int(18)] = (d_831_v4_) | (d_831_v4_)
            nw143_[int(19)] = d_831_v4_
            nw143_[int(20)] = d_831_v4_
            def iife52_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in _dafny.IntegerRange(156, 62):
                    d_842_v13_: int = compr_22_
                    if ((156) <= (d_842_v13_)) and ((d_842_v13_) < (62)):
                        coll22_[default__.safeDivisionInt(d_842_v13_, p1)] = (self).f0
                return _dafny.Map(coll22_)
            nw143_[int(21)] = iife52_()
            
            nw143_[int(22)] = d_831_v4_
            nw143_[int(23)] = d_831_v4_
            nw143_[int(24)] = default__.fm11(p1, globalState)
            nw143_[int(25)] = d_831_v4_
            def iife53_():
                coll23_ = _dafny.Map()
                compr_23_: int
                for compr_23_ in _dafny.IntegerRange(659, 736):
                    d_843_v14_: int = compr_23_
                    if ((659) <= (d_843_v14_)) and ((d_843_v14_) < (736)):
                        coll23_[(d_843_v14_) - (p1)] = 225
                return _dafny.Map(coll23_)
            nw143_[int(26)] = iife53_()
            
            nw143_[int(27)] = d_831_v4_
            nw143_[int(28)] = d_831_v4_
            d_838_v15_ = nw143_
            index131_ = default__.safeIndex(121, (d_838_v15_).length(0))
            (d_838_v15_)[index131_] = _dafny.Map({p1: (self).f0})
            d_844_v16_: _dafny.Array
            def lambda35_(d_845_v0_):
                def lambda36_(d_846_i0_):
                    return _dafny.SeqWithoutIsStrInference([d_845_v0_, d_845_v0_])

                return lambda36_

            init17_ = lambda35_(d_828_v0_)
            nw144_ = _dafny.Array(None, 12)
            for i0_17_ in range(nw144_.length(0)):
                nw144_[i0_17_] = init17_(i0_17_)
            d_844_v16_ = nw144_
            index132_ = default__.safeIndex(859, (d_844_v16_).length(0))
            (d_844_v16_)[index132_] = default__.fm12((self).f0, (self).f0, d_828_v0_, globalState)
            d_847_v17_: _dafny.Array
            def lambda37_(d_848_v0_):
                def lambda38_(d_849_i1_):
                    return not (d_848_v0_) or (d_848_v0_)

                return lambda38_

            init18_ = lambda37_(d_828_v0_)
            nw145_ = _dafny.Array(None, 1)
            for i0_18_ in range(nw145_.length(0)):
                nw145_[i0_18_] = init18_(i0_18_)
            d_847_v17_ = nw145_
            d_850_v18_: _dafny.Seq
            d_850_v18_ = _dafny.SeqWithoutIsStrInference([d_828_v0_, d_828_v0_])
            d_851_v19_: _dafny.Array
            d_851_v19_ = d_847_v17_
            index133_ = default__.safeIndex(121, (d_838_v15_).length(0))
            index134_ = default__.safeIndex(859, (d_844_v16_).length(0))
            rhs125_ = d_831_v4_
            rhs126_ = d_837_v11_
            rhs127_ = d_850_v18_
            rhs128_ = (d_847_v17_)
            rhs129_ = ((d_837_v11_) + ((_dafny.SeqWithoutIsStrInference([(d_836_v10_).f3 for d_852_i2_ in range(default__.abs(156))])) + (d_837_v11_))).set(default__.safeIndex(len(d_850_v18_), len((d_837_v11_) + ((_dafny.SeqWithoutIsStrInference([(d_836_v10_).f3 for d_853_i2_ in range(default__.abs(156))])) + (d_837_v11_)))), (d_836_v10_).f3)
            lhs52_ = d_838_v15_
            lhs53_ = default__.safeIndex(121, (d_838_v15_).length(0))
            lhs54_ = d_844_v16_
            lhs55_ = default__.safeIndex(859, (d_844_v16_).length(0))
            lhs52_[lhs53_] = rhs125_
            d_837_v11_ = rhs126_
            lhs54_[lhs55_] = rhs127_
            d_847_v17_ = rhs128_
            d_837_v11_ = rhs129_
            d_854_v20_: C0
            nw146_ = C0()
            nw146_.ctor__(_dafny.CodePoint('c'))
            d_854_v20_ = nw146_
            pat_let_tv23_ = d_832_v6_
            pat_let_tv24_ = d_832_v6_
            def iife54_(_pat_let15_0):
                def iife55_(d_855_dt__update__tmp_h1_):
                    def iife56_(_pat_let16_0):
                        def iife57_(d_856_dt__update_hcf7_h0_):
                            return D3_DC7(d_856_dt__update_hcf7_h0_)
                        return iife57_(_pat_let16_0)
                    return iife56_((pat_let_tv23_).set(default__.safeIndex(615, len(pat_let_tv24_)), (self).f0))
                return iife55_(_pat_let15_0)
            source11_ = iife54_(d_833_v7_)
            if source11_.is_DC8:
                d_857___mcc_h0_ = source11_.cf8
                d_858___mcc_h1_ = source11_.cf9
                d_859_cf9_ = d_858___mcc_h1_
                d_860_cf8_ = d_857___mcc_h0_
                d_861_v21_: T0
                nw147_ = C3()
                nw147_.ctor__((self).f0)
                d_861_v21_ = nw147_
                d_862_v22_: _dafny.Array
                nw148_ = _dafny.Array(None, 25)
                nw148_[int(0)] = d_861_v21_
                nw148_[int(1)] = d_861_v21_
                nw148_[int(2)] = d_861_v21_
                nw148_[int(3)] = d_861_v21_
                nw148_[int(4)] = d_861_v21_
                nw148_[int(5)] = d_861_v21_
                nw148_[int(6)] = d_861_v21_
                nw148_[int(7)] = d_861_v21_
                nw148_[int(8)] = d_861_v21_
                nw148_[int(9)] = d_861_v21_
                nw148_[int(10)] = d_861_v21_
                nw148_[int(11)] = d_861_v21_
                nw148_[int(12)] = d_861_v21_
                nw148_[int(13)] = d_861_v21_
                nw148_[int(14)] = d_861_v21_
                nw148_[int(15)] = d_861_v21_
                nw148_[int(16)] = d_861_v21_
                nw148_[int(17)] = d_861_v21_
                nw148_[int(18)] = d_861_v21_
                nw148_[int(19)] = d_861_v21_
                nw148_[int(20)] = d_861_v21_
                nw148_[int(21)] = d_861_v21_
                nw148_[int(22)] = d_861_v21_
                nw148_[int(23)] = d_861_v21_
                nw148_[int(24)] = d_861_v21_
                d_862_v22_ = nw148_
                index135_ = default__.safeIndex(854, (d_862_v22_).length(0))
                (d_862_v22_)[index135_] = d_861_v21_
                d_863_v23_: D5
                d_863_v23_ = D5_DC13(len(_dafny.Map({len(d_837_v11_): d_828_v0_})), _dafny.Set({d_860_cf8_}), (d_837_v11_)[default__.safeIndex((d_861_v21_).f0, len(d_837_v11_))], d_859_cf9_, d_828_v0_)
                index136_ = default__.safeIndex(854, (d_862_v22_).length(0))
                (d_862_v22_)[index136_] = ((d_861_v21_ if (d_863_v23_).cf21 else d_861_v21_) if d_860_cf8_ else d_861_v21_)
                r1 = d_859_cf9_
                d_864_v24_: _dafny.Set
                d_864_v24_ = _dafny.Set({d_828_v0_, d_860_cf8_, d_860_cf8_, d_860_cf8_})
                r1 = (((self).f0) * (default__.fm3(d_864_v24_, d_835_v9_, d_859_cf9_, len(d_837_v11_), globalState))) * ((d_861_v21_).f0)
                d_865_v25_: _dafny.Array
                nw149_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_865_v25_ = nw149_
                index137_ = default__.safeIndex(24, (d_865_v25_).length(0))
                (d_865_v25_)[index137_] = d_847_v17_
                index138_ = default__.safeIndex(24, (d_865_v25_).length(0))
                (d_865_v25_)[index138_] = d_847_v17_
            elif source11_.is_DC9:
                d_866___mcc_h2_ = source11_.cf10
                d_867___mcc_h3_ = source11_.cf11
                d_868___mcc_h4_ = source11_.cf12
                d_869_cf12_ = d_868___mcc_h4_
                d_870_cf11_ = d_867___mcc_h3_
                d_871_cf10_ = d_866___mcc_h2_
                d_828_v0_ = d_828_v0_
                d_872_v26_: C1
                nw150_ = C1()
                nw150_.ctor__(default__.safeModuloInt(d_871_cf10_, len((d_844_v16_)[default__.safeIndex(859, (d_844_v16_).length(0))])))
                d_872_v26_ = nw150_
                d_873_v27_: C1
                nw151_ = C1()
                nw151_.ctor__((self).f0)
                d_873_v27_ = nw151_
                d_874_v28_: _dafny.Map
                d_874_v28_ = _dafny.Map({d_828_v0_: p1})
                d_875_v29_: _dafny.Seq
                d_875_v29_ = _dafny.SeqWithoutIsStrInference([d_874_v28_, d_874_v28_, d_874_v28_])
                d_876_v30_: D3
                d_876_v30_ = D3_DC8(d_828_v0_, p1)
                d_877_v31_: _dafny.Map
                d_877_v31_ = _dafny.Map({(d_875_v29_)[default__.safeIndex(p1, len(d_875_v29_))]: (self).fm6((d_876_v30_).cf8, d_837_v11_, d_837_v11_, 432, globalState)})
                d_877_v31_ = (d_877_v31_).set(d_874_v28_, d_869_cf12_)
            elif True:
                d_878___mcc_h5_ = source11_.cf7
                d_879_cf7_ = d_878___mcc_h5_
                d_880_v32_: _dafny.Map
                d_880_v32_ = _dafny.Map({p1: d_847_v17_})
                d_881_v33_: _dafny.Set
                d_881_v33_ = _dafny.Set({d_828_v0_})
                d_882_v34_: _dafny.Map
                d_882_v34_ = _dafny.Map({d_828_v0_: _dafny.Map({len(d_881_v33_): d_847_v17_})})
                d_883_v35_: _dafny.Map
                d_883_v35_ = _dafny.Map({False: d_828_v0_})
                rhs130_ = ((d_882_v34_)[((d_883_v35_)[(d_850_v18_)[default__.safeIndex((self).f0, len(d_850_v18_))]] if ((d_850_v18_)[default__.safeIndex((self).f0, len(d_850_v18_))]) in (d_883_v35_) else default__.fm2(d_828_v0_, (self).f0, d_828_v0_, d_828_v0_, globalState))] if (((d_883_v35_)[(d_850_v18_)[default__.safeIndex((self).f0, len(d_850_v18_))]] if ((d_850_v18_)[default__.safeIndex((self).f0, len(d_850_v18_))]) in (d_883_v35_) else default__.fm2(d_828_v0_, (self).f0, d_828_v0_, d_828_v0_, globalState))) in (d_882_v34_) else (_dafny.Map({p1: d_847_v17_})).set(p1, d_847_v17_))
                rhs131_ = (0) - ((0) - (default__.safeModuloInt((self).f0, p1)))
                d_880_v32_ = rhs130_
                r2 = rhs131_
                d_884_v36_: _dafny.MultiSet
                d_884_v36_ = _dafny.MultiSet([(self).f0, default__.fm3(d_881_v33_, (d_836_v10_).f3, p1, (0) - ((self).f0), globalState)])
                d_885_v37_: T1
                nw152_ = C2()
                nw152_.ctor__(p1)
                d_885_v37_ = nw152_
                d_886_v38_: _dafny.Set
                d_886_v38_ = _dafny.Set({d_885_v37_})
                d_887_v39_: C2
                nw153_ = C2()
                nw153_.ctor__(default__.fm3(d_881_v33_, (d_836_v10_).f3, p1, ((d_884_v36_)[len(d_886_v38_)] if (len(d_886_v38_)) in (d_884_v36_) else p1), globalState))
                d_887_v39_ = nw153_
                d_887_v39_ = d_887_v39_
                d_888_v40_: _dafny.Map
                d_888_v40_ = _dafny.Map({d_834_v8_: (_dafny.Map({d_828_v0_: d_828_v0_})).set(d_828_v0_, d_828_v0_)})
                d_889_v41_: _dafny.Map
                d_889_v41_ = _dafny.Map({467: d_883_v35_})
                d_890_v42_: _dafny.Map
                d_890_v42_ = _dafny.Map({(d_836_v10_).f3: (d_885_v37_).f0})
                d_891_v43_: _dafny.Map
                d_891_v43_ = _dafny.Map({not(d_828_v0_): ((d_890_v42_)[(d_836_v10_).f3] if ((d_836_v10_).f3) in (d_890_v42_) else len(d_837_v11_))})
                d_892_v44_: _dafny.Map
                d_892_v44_ = _dafny.Map({p1: d_828_v0_})
                d_883_v35_ = (((d_888_v40_)[d_834_v8_] if (d_834_v8_) in (d_888_v40_) else ((d_889_v41_)[len(d_891_v43_)] if (len(d_891_v43_)) in (d_889_v41_) else _dafny.Map({d_828_v0_: ((d_892_v44_)[p1] if (p1) in (d_892_v44_) else d_828_v0_)})))).set(d_828_v0_, d_828_v0_)
                index139_ = default__.safeIndex(959, (d_834_v8_).length(0))
                (d_834_v8_)[index139_] = (0) - ((d_885_v37_).f0)
                d_893_v45_: _dafny.Map
                d_893_v45_ = _dafny.Map({not(d_828_v0_): _dafny.Set({d_828_v0_})})
                index140_ = default__.safeIndex(959, (d_834_v8_).length(0))
                (d_834_v8_)[index140_] = default__.fm3((((d_893_v45_)[d_828_v0_] if (d_828_v0_) in (d_893_v45_) else d_881_v33_)) - (d_881_v33_), _dafny.CodePoint('i'), (d_885_v37_).f0, (self).f0, globalState)
            d_894_v46_: _dafny.Set
            d_894_v46_ = _dafny.Set({not(d_828_v0_), d_828_v0_, d_828_v0_})
            d_895_v47_: _dafny.MultiSet
            d_895_v47_ = _dafny.MultiSet([721])
            d_896_v48_: D0
            d_896_v48_ = D0_DC1(((d_844_v16_)[default__.safeIndex(859, (d_844_v16_).length(0))])[default__.safeIndex(p1, len((d_844_v16_)[default__.safeIndex(859, (d_844_v16_).length(0))]))])
            d_897_v49_: D2
            d_897_v49_ = D2_DC3(_dafny.SeqWithoutIsStrInference([d_896_v48_]))
            d_898_v50_: _dafny.Map
            d_898_v50_ = _dafny.Map({(self).f0: d_897_v49_})
            d_899_v51_: _dafny.MultiSet
            d_899_v51_ = _dafny.MultiSet([D2_DC6(((d_898_v50_)[p1] if (p1) in (d_898_v50_) else d_897_v49_)), D2_DC6(d_897_v49_)])
            d_900_v52_: _dafny.Map
            d_900_v52_ = _dafny.Map({default__.fm18(default__.fm3(d_894_v46_, d_835_v9_, (d_895_v47_).cardinality, (self).f0, globalState), globalState): d_899_v51_})
            d_901_v53_: C3
            nw154_ = C3()
            nw154_.ctor__((d_829_v1_).cardinality)
            d_901_v53_ = nw154_
            d_902_v54_: _dafny.Map
            d_902_v54_ = _dafny.Map({p1: d_901_v53_})
            d_903_v55_: _dafny.Map
            d_903_v55_ = _dafny.Map({((d_900_v52_)[d_837_v11_] if (d_837_v11_) in (d_900_v52_) else d_899_v51_): ((d_902_v54_)[693] if (693) in (d_902_v54_) else d_901_v53_)})
            d_903_v55_ = (d_903_v55_).set(d_899_v51_, d_901_v53_)
            if (p1) < (default__.safeModuloInt(len(d_837_v11_), (self).f0)):
                index141_ = default__.safeIndex(303, (d_847_v17_).length(0))
                (d_847_v17_)[index141_] = d_828_v0_
                d_904_v56_: D2
                d_904_v56_ = D2_DC4(d_837_v11_, d_828_v0_)
                index142_ = default__.safeIndex(303, (d_847_v17_).length(0))
                rhs132_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))) + (_dafny.SeqWithoutIsStrInference([(d_854_v20_).f3 for d_905_i3_ in range(default__.abs(95))]))
                rhs133_ = not((d_837_v11_) < ((d_904_v56_).cf4))
                lhs56_ = d_847_v17_
                lhs57_ = default__.safeIndex(303, (d_847_v17_).length(0))
                d_837_v11_ = rhs132_
                lhs56_[lhs57_] = rhs133_
                (d_901_v53_).m7((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jljbktggr"))) + (d_837_v11_), d_828_v0_, d_834_v8_, len((d_832_v6_ if d_828_v0_ else d_832_v6_)), globalState)
                index143_ = default__.safeIndex(303, (d_847_v17_).length(0))
                (d_847_v17_)[index143_] = (p1) > (p1)
                d_906_v57_: C1
                nw155_ = C1()
                nw155_.ctor__((self).f0)
                d_906_v57_ = nw155_
                d_907_v58_: _dafny.Map
                d_907_v58_ = _dafny.Map({(self).f0: d_828_v0_})
                d_908_v59_: _dafny.Map
                d_908_v59_ = _dafny.Map({d_906_v57_: d_907_v58_})
                d_909_v60_: _dafny.Map
                d_909_v60_ = _dafny.Map({d_828_v0_: d_908_v59_})
                d_909_v60_ = (d_909_v60_).set((d_906_v57_).fm6(d_828_v0_, d_837_v11_, d_837_v11_, p1, globalState), (d_908_v59_ if d_828_v0_ else (d_908_v59_).set(d_906_v57_, d_907_v58_)))
                index144_ = default__.safeIndex(859, (d_844_v16_).length(0))
                (d_844_v16_)[index144_] = (d_844_v16_)[default__.safeIndex(859, (d_844_v16_).length(0))]
            elif True:
                d_835_v9_ = (d_836_v10_).f3
                d_910_v61_: C3
                nw156_ = C3()
                nw156_.ctor__(p1)
                d_910_v61_ = nw156_
                d_911_v62_: _dafny.Array
                nw157_ = _dafny.Array(_dafny.Set({}), 17)
                d_911_v62_ = nw157_
                index145_ = default__.safeIndex(898, (d_911_v62_).length(0))
                (d_911_v62_)[index145_] = d_894_v46_
                d_912_v63_: _dafny.Array
                nw158_ = _dafny.Array(_dafny.MultiSet({}), 27)
                d_912_v63_ = nw158_
                index146_ = default__.safeIndex(706, (d_912_v63_).length(0))
                (d_912_v63_)[index146_] = d_895_v47_
                index147_ = default__.safeIndex(898, (d_911_v62_).length(0))
                index148_ = default__.safeIndex(706, (d_912_v63_).length(0))
                rhs134_ = (p1 if default__.fm2(d_828_v0_, (self).f0, not(d_828_v0_), False, globalState) else default__.fm3(d_894_v46_, (d_854_v20_).f3, len(default__.fm11(len(d_837_v11_), globalState)), len(d_830_v3_), globalState))
                rhs135_ = (0) - ((self).f0)
                rhs136_ = d_834_v8_
                rhs137_ = d_894_v46_
                rhs138_ = _dafny.MultiSet([(self).f0])
                lhs58_ = d_911_v62_
                lhs59_ = default__.safeIndex(898, (d_911_v62_).length(0))
                lhs60_ = d_912_v63_
                lhs61_ = default__.safeIndex(706, (d_912_v63_).length(0))
                r2 = rhs134_
                r2 = rhs135_
                d_834_v8_ = rhs136_
                lhs58_[lhs59_] = rhs137_
                lhs60_[lhs61_] = rhs138_
                d_835_v9_ = (d_835_v9_ if d_828_v0_ else (d_854_v20_).f3)
                index149_ = default__.safeIndex(280, (d_847_v17_).length(0))
                (d_847_v17_)[index149_] = True
                index150_ = default__.safeIndex(280, (d_847_v17_).length(0))
                (d_847_v17_)[index150_] = not(default__.fm2(((d_911_v62_)[default__.safeIndex(898, (d_911_v62_).length(0))]).ispropersubset(_dafny.Set({d_828_v0_})), p1, (True) or (d_828_v0_), (d_828_v0_ if d_828_v0_ else d_828_v0_), globalState))
        elif True:
            d_913_v64_: C3
            nw159_ = C3()
            nw159_.ctor__(p1)
            d_913_v64_ = nw159_
            if d_828_v0_:
                d_914_v65_: _dafny.Array
                nw160_ = _dafny.Array(int(0), 6)
                d_914_v65_ = nw160_
                index151_ = default__.safeIndex(406, (d_914_v65_).length(0))
                (d_914_v65_)[index151_] = p1
                index152_ = default__.safeIndex(406, (d_914_v65_).length(0))
                (d_914_v65_)[index152_] = p1
                d_915_v66_: _dafny.Array
                def lambda39_(d_916_p1_):
                    def lambda40_(d_917_i4_):
                        return (d_916_p1_) >= (854)

                    return lambda40_

                init19_ = lambda39_(p1)
                nw161_ = _dafny.Array(None, 1)
                for i0_19_ in range(nw161_.length(0)):
                    nw161_[i0_19_] = init19_(i0_19_)
                d_915_v66_ = nw161_
                index153_ = default__.safeIndex(360, (d_915_v66_).length(0))
                (d_915_v66_)[index153_] = d_828_v0_
                d_918_v67_: _dafny.Seq
                d_918_v67_ = _dafny.SeqWithoutIsStrInference([not(d_828_v0_)])
                d_919_v68_: D7
                d_919_v68_ = D7_DC16(d_918_v67_)
                d_920_v69_: str
                d_920_v69_ = _dafny.CodePoint('k')
                d_921_v70_: str
                d_921_v70_ = d_920_v69_
                d_922_v71_: _dafny.Map
                d_922_v71_ = _dafny.Map({(0) - ((d_914_v65_)[default__.safeIndex(406, (d_914_v65_).length(0))]): d_921_v70_})
                d_923_v72_: _dafny.Seq
                d_923_v72_ = _dafny.SeqWithoutIsStrInference([d_914_v65_])
                d_924_v73_: _dafny.Seq
                d_924_v73_ = _dafny.SeqWithoutIsStrInference([(self).f0])
                d_925_v74_: _dafny.Set
                d_925_v74_ = _dafny.Set({d_828_v0_})
                d_926_v75_: _dafny.MultiSet
                d_926_v75_ = _dafny.MultiSet([len(d_924_v73_), default__.fm3(d_925_v74_, _dafny.CodePoint('q'), -458, -6, globalState)])
                d_927_v76_: _dafny.Seq
                d_927_v76_ = _dafny.SeqWithoutIsStrInference([p1, len((D11_DC23(d_923_v72_)).cf37), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xurqdyud"))), (self).f0, ((d_926_v75_)[(d_914_v65_)[default__.safeIndex(406, (d_914_v65_).length(0))]] if ((d_914_v65_)[default__.safeIndex(406, (d_914_v65_).length(0))]) in (d_926_v75_) else p1)])
                d_928_v77_: _dafny.Seq
                d_928_v77_ = _dafny.SeqWithoutIsStrInference([len(d_922_v71_), (len(d_927_v76_)) + ((d_926_v75_).cardinality)])
                d_929_v78_: _dafny.Seq
                d_929_v78_ = _dafny.SeqWithoutIsStrInference([d_928_v77_, d_927_v76_, _dafny.SeqWithoutIsStrInference([(self).f0]), _dafny.SeqWithoutIsStrInference([(self).f0])])
                index154_ = default__.safeIndex(360, (d_915_v66_).length(0))
                rhs139_ = d_828_v0_
                rhs140_ = d_919_v68_
                rhs141_ = (d_929_v78_)[default__.safeIndex(p1, len(d_929_v78_))]
                rhs142_ = not (d_828_v0_) or (d_828_v0_)
                lhs62_ = d_915_v66_
                lhs63_ = default__.safeIndex(360, (d_915_v66_).length(0))
                lhs62_[lhs63_] = rhs139_
                d_919_v68_ = rhs140_
                d_928_v77_ = rhs141_
                d_828_v0_ = rhs142_
                d_930_v79_: D7
                d_930_v79_ = D7_DC18((0) - ((self).f0))
                d_931_v80_: C1
                nw162_ = C1()
                nw162_.ctor__((self).f0)
                d_931_v80_ = nw162_
                pat_let_tv25_ = d_931_v80_
                pat_let_tv26_ = d_828_v0_
                d_932_v81_: _dafny.Array
                nw163_ = _dafny.Array(None, 29)
                nw163_[int(0)] = D7_DC18((0) - (((d_829_v1_)[d_828_v0_] if (d_828_v0_) in (d_829_v1_) else (d_914_v65_)[default__.safeIndex(406, (d_914_v65_).length(0))])))
                nw163_[int(1)] = d_930_v79_
                nw163_[int(2)] = d_930_v79_
                nw163_[int(3)] = d_930_v79_
                nw163_[int(4)] = d_930_v79_
                nw163_[int(5)] = d_930_v79_
                nw163_[int(6)] = D7_DC18((self).f0)
                def iife59_(_pat_let18_0):
                    def iife60_(d_933_dt__update__tmp_h2_):
                        def iife61_(_pat_let19_0):
                            def iife62_(d_934_dt__update_hcf30_h0_):
                                return D7_DC18(d_934_dt__update_hcf30_h0_)
                            return iife62_(_pat_let19_0)
                        return iife61_(len(_dafny.Map({pat_let_tv25_: pat_let_tv26_})))
                    return iife60_(_pat_let18_0)
                def iife58_(_pat_let17_0):
                    def iife63_(d_935_dt__update__tmp_h3_):
                        def iife64_(_pat_let20_0):
                            def iife65_(d_936_dt__update_hcf30_h1_):
                                return D7_DC18(d_936_dt__update_hcf30_h1_)
                            return iife65_(_pat_let20_0)
                        return iife64_((self).f0)
                    return iife63_(_pat_let17_0)
                nw163_[int(7)] = iife58_(iife59_(d_930_v79_))
                nw163_[int(8)] = d_930_v79_
                nw163_[int(9)] = D7_DC18(default__.fm3(d_925_v74_, d_920_v69_, p1, (d_829_v1_).cardinality, globalState))
                nw163_[int(10)] = default__.fm26(globalState)
                nw163_[int(11)] = d_930_v79_
                nw163_[int(12)] = D7_DC18(p1)
                nw163_[int(13)] = d_930_v79_
                nw163_[int(14)] = d_930_v79_
                nw163_[int(15)] = d_930_v79_
                nw163_[int(16)] = d_930_v79_
                nw163_[int(17)] = D7_DC18(p1)
                nw163_[int(18)] = default__.fm26(globalState)
                nw163_[int(19)] = D7_DC18((d_914_v65_)[default__.safeIndex(406, (d_914_v65_).length(0))])
                nw163_[int(20)] = d_930_v79_
                nw163_[int(21)] = d_930_v79_
                nw163_[int(22)] = (d_930_v79_ if d_828_v0_ else d_930_v79_)
                nw163_[int(23)] = d_930_v79_
                nw163_[int(24)] = d_930_v79_
                nw163_[int(25)] = d_930_v79_
                nw163_[int(26)] = d_930_v79_
                nw163_[int(27)] = d_930_v79_
                nw163_[int(28)] = d_930_v79_
                d_932_v81_ = nw163_
                index155_ = default__.safeIndex(563, (d_932_v81_).length(0))
                (d_932_v81_)[index155_] = d_930_v79_
                index156_ = default__.safeIndex(563, (d_932_v81_).length(0))
                rhs143_ = d_930_v79_
                rhs144_ = (d_925_v74_) | (d_925_v74_)
                rhs145_ = (d_914_v65_)[default__.safeIndex(406, (d_914_v65_).length(0))]
                lhs64_ = d_932_v81_
                lhs65_ = default__.safeIndex(563, (d_932_v81_).length(0))
                lhs64_[lhs65_] = rhs143_
                d_925_v74_ = rhs144_
                r1 = rhs145_
                d_937_v82_: _dafny.Seq
                d_937_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aixybf"))
                d_937_v82_ = d_937_v82_
                r1 = (default__.safeDivisionInt((d_914_v65_)[default__.safeIndex(406, (d_914_v65_).length(0))], (self).f0)) + ((0) - ((self).f0))
            elif True:
                d_938_v83_: _dafny.Array
                def lambda41_(d_939_v0_):
                    def lambda42_(d_940_i5_):
                        return d_939_v0_

                    return lambda42_

                init20_ = lambda41_(d_828_v0_)
                nw164_ = _dafny.Array(None, 4)
                for i0_20_ in range(nw164_.length(0)):
                    nw164_[i0_20_] = init20_(i0_20_)
                d_938_v83_ = nw164_
                d_941_v84_: _dafny.Seq
                d_941_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elco"))
                index157_ = default__.safeIndex(379, (d_938_v83_).length(0))
                (d_938_v83_)[index157_] = (d_913_v64_).fm6(d_828_v0_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_942_i6_ in range(default__.abs(87))]), d_941_v84_, ((d_829_v1_)[not(d_828_v0_)] if (not(d_828_v0_)) in (d_829_v1_) else p1), globalState)
                d_943_v85_: T0
                nw165_ = C3()
                nw165_.ctor__(len(d_941_v84_))
                d_943_v85_ = nw165_
                d_944_v86_: _dafny.Seq
                d_944_v86_ = _dafny.SeqWithoutIsStrInference([d_828_v0_, d_828_v0_])
                d_945_v87_: D7
                d_945_v87_ = D7_DC16(d_944_v86_)
                d_946_v88_: _dafny.MultiSet
                d_946_v88_ = _dafny.MultiSet([d_945_v87_])
                d_947_v89_: str
                d_947_v89_ = _dafny.CodePoint('d')
                d_948_v90_: _dafny.Set
                d_948_v90_ = _dafny.Set({d_828_v0_, default__.fm2(d_828_v0_, len(_dafny.Map({d_828_v0_: p1})), d_828_v0_, d_828_v0_, globalState)})
                d_949_v91_: _dafny.Map
                d_949_v91_ = _dafny.Map({(0) - ((0) - ((d_943_v85_).f0)): (self).f0})
                d_950_v92_: _dafny.MultiSet
                d_950_v92_ = _dafny.MultiSet([p1, (self).f0])
                index158_ = default__.safeIndex(379, (d_938_v83_).length(0))
                rhs146_ = (0) - (default__.fm3(default__.fm27(d_828_v0_, d_946_v88_, p1, _dafny.Set({(0) - ((self).f0), (d_943_v85_).f0, (self).f0, (d_943_v85_).f0}), globalState), d_947_v89_, (default__.fm3(d_948_v90_, d_947_v89_, len(d_949_v91_), (d_943_v85_).f0, globalState) if not(d_828_v0_) else ((d_950_v92_)[(self).f0] if ((self).f0) in (d_950_v92_) else (self).f0)), (0) - ((d_829_v1_).cardinality), globalState))
                rhs147_ = ((d_829_v1_).cardinality) < ((959) - (default__.fm3(d_948_v90_, d_947_v89_, 999, len(d_941_v84_), globalState)))
                rhs148_ = (_dafny.MultiSet([len(d_944_v86_), 191, p1])).isdisjoint(((d_950_v92_).set((d_943_v85_).f0, default__.abs((d_943_v85_).f0))).intersection(default__.fm22(globalState)))
                rhs149_ = (p1) * ((d_943_v85_).f0)
                rhs150_ = d_943_v85_
                lhs66_ = d_938_v83_
                lhs67_ = default__.safeIndex(379, (d_938_v83_).length(0))
                r2 = rhs146_
                d_828_v0_ = rhs147_
                lhs66_[lhs67_] = rhs148_
                r1 = rhs149_
                d_943_v85_ = rhs150_
                d_951_v93_: _dafny.Map
                d_951_v93_ = _dafny.Map({(default__.fm17(p1, globalState) if d_828_v0_ else d_947_v89_): (d_943_v85_).f0})
                def iife66_():
                    coll24_ = _dafny.Map()
                    compr_24_: str
                    for compr_24_ in ((_dafny.MultiSet([_dafny.CodePoint('e')])).set(d_947_v89_, default__.abs(len(d_941_v84_)))).Elements:
                        d_952_v94_: str = compr_24_
                        if (d_952_v94_) in ((_dafny.MultiSet([_dafny.CodePoint('e')])).set(d_947_v89_, default__.abs(len(d_941_v84_)))):
                            coll24_[d_952_v94_] = (d_943_v85_).f0
                    return _dafny.Map(coll24_)
                d_951_v93_ = (iife66_()
                ) | ((d_951_v93_) | (d_951_v93_))
                d_953_v95_: T1
                nw166_ = C2()
                nw166_.ctor__((d_943_v85_).f0)
                d_953_v95_ = nw166_
                d_954_v96_: _dafny.Map
                d_954_v96_ = _dafny.Map({d_953_v95_: d_941_v84_})
                d_941_v84_ = ((d_954_v96_)[d_953_v95_] if (d_953_v95_) in (d_954_v96_) else d_941_v84_)
                d_955_v97_: _dafny.Seq
                d_955_v97_ = _dafny.SeqWithoutIsStrInference([(d_953_v95_).f0])
                d_956_v98_: D3
                d_956_v98_ = D3_DC7(d_955_v97_)
                d_957_v99_: _dafny.Set
                d_957_v99_ = _dafny.Set({d_938_v83_, d_938_v83_})
                d_958_v100_: _dafny.Array
                nw167_ = _dafny.Array(None, 21)
                nw167_[int(0)] = d_913_v64_
                nw167_[int(1)] = d_913_v64_
                nw167_[int(2)] = d_913_v64_
                nw167_[int(3)] = d_913_v64_
                nw167_[int(4)] = d_913_v64_
                nw167_[int(5)] = d_913_v64_
                nw167_[int(6)] = d_913_v64_
                nw167_[int(7)] = d_913_v64_
                nw167_[int(8)] = d_913_v64_
                nw167_[int(9)] = d_913_v64_
                nw167_[int(10)] = d_913_v64_
                nw167_[int(11)] = d_913_v64_
                nw167_[int(12)] = d_913_v64_
                nw167_[int(13)] = d_913_v64_
                nw167_[int(14)] = d_913_v64_
                nw167_[int(15)] = d_913_v64_
                nw167_[int(16)] = d_913_v64_
                nw167_[int(17)] = d_913_v64_
                nw167_[int(18)] = d_913_v64_
                nw167_[int(19)] = d_913_v64_
                nw167_[int(20)] = d_913_v64_
                d_958_v100_ = nw167_
                d_959_v101_: _dafny.Map
                d_959_v101_ = _dafny.Map({d_958_v100_: not(d_828_v0_)})
                d_960_v102_: _dafny.Set
                d_960_v102_ = _dafny.Set({d_913_v64_})
                index159_ = default__.safeIndex(379, (d_938_v83_).length(0))
                index160_ = default__.safeIndex(379, (d_938_v83_).length(0))
                index161_ = default__.safeIndex(379, (d_938_v83_).length(0))
                rhs151_ = (d_938_v83_)[default__.safeIndex(379, (d_938_v83_).length(0))]
                rhs152_ = d_956_v98_
                rhs153_ = not ((d_957_v99_).ispropersubset(d_957_v99_)) or (d_828_v0_)
                rhs154_ = ((d_959_v101_)[d_958_v100_] if (d_958_v100_) in (d_959_v101_) else (d_960_v102_).issubset(d_960_v102_))
                lhs68_ = d_938_v83_
                lhs69_ = default__.safeIndex(379, (d_938_v83_).length(0))
                lhs70_ = d_938_v83_
                lhs71_ = default__.safeIndex(379, (d_938_v83_).length(0))
                lhs72_ = d_938_v83_
                lhs73_ = default__.safeIndex(379, (d_938_v83_).length(0))
                lhs68_[lhs69_] = rhs151_
                d_956_v98_ = rhs152_
                lhs70_[lhs71_] = rhs153_
                lhs72_[lhs73_] = rhs154_
                rhs155_ = d_828_v0_
                rhs156_ = ((_dafny.MultiSet([d_941_v84_])).set(((d_941_v84_) + (d_941_v84_)).set(default__.safeIndex(p1, len((d_941_v84_) + (d_941_v84_))), d_947_v89_), default__.abs(((d_953_v95_).f0) + ((self).f0)))).cardinality
                rhs157_ = _dafny.SeqWithoutIsStrInference([d_828_v0_, (d_938_v83_)[default__.safeIndex(379, (d_938_v83_).length(0))], (d_938_v83_)[default__.safeIndex(379, (d_938_v83_).length(0))]])
                d_828_v0_ = rhs155_
                r2 = rhs156_
                d_944_v86_ = rhs157_
            d_961_v103_: D2
            d_961_v103_ = D2_DC4(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_962_i7_ in range(default__.abs(-137))]), (p1) < (p1))
            d_961_v103_ = d_961_v103_
            if (d_828_v0_) and (d_828_v0_):
                d_963_v104_: _dafny.Array
                nw168_ = _dafny.Array(int(0), 5)
                d_963_v104_ = nw168_
                index162_ = default__.safeIndex(440, (d_963_v104_).length(0))
                (d_963_v104_)[index162_] = -560
                index163_ = default__.safeIndex(440, (d_963_v104_).length(0))
                (d_963_v104_)[index163_] = (0) - (p1)
                d_964_v105_: str
                d_964_v105_ = _dafny.CodePoint('n')
                d_965_v106_: _dafny.Map
                d_965_v106_ = _dafny.Map({d_828_v0_: d_964_v105_})
                d_965_v106_ = _dafny.Map({not((not(not(d_828_v0_))) == (d_828_v0_)): d_964_v105_})
                d_828_v0_ = d_828_v0_
                d_966_v107_: _dafny.Seq
                d_966_v107_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcmmxy"))
                d_966_v107_ = d_966_v107_
                d_828_v0_ = not (d_828_v0_) or (d_828_v0_)
            elif True:
                d_967_v108_: _dafny.Array
                nw169_ = _dafny.Array(False, 1)
                d_967_v108_ = nw169_
                index164_ = default__.safeIndex(491, (d_967_v108_).length(0))
                (d_967_v108_)[index164_] = d_828_v0_
                index165_ = default__.safeIndex(491, (d_967_v108_).length(0))
                (d_967_v108_)[index165_] = d_828_v0_
                d_968_v109_: _dafny.Array
                def lambda43_(d_969_p1_):
                    def lambda44_(d_970_i8_):
                        return (d_970_i8_) * (d_969_p1_)

                    return lambda44_

                init21_ = lambda43_(p1)
                nw170_ = _dafny.Array(None, 14)
                for i0_21_ in range(nw170_.length(0)):
                    nw170_[i0_21_] = init21_(i0_21_)
                d_968_v109_ = nw170_
                d_971_v110_: _dafny.Seq
                d_971_v110_ = _dafny.SeqWithoutIsStrInference([d_968_v109_])
                d_968_v109_ = (d_971_v110_)[default__.safeIndex((self).f0, len(d_971_v110_))]
                index166_ = default__.safeIndex(491, (d_967_v108_).length(0))
                (d_967_v108_)[index166_] = ((-114 if not(d_828_v0_) else (self).f0)) > (454)
                r0 = (self).f0
                d_972_v111_: _dafny.Array
                nw171_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_972_v111_ = nw171_
                d_973_v112_: _dafny.Array
                def lambda45_(d_974_v108_):
                    def lambda46_(d_975_i9_):
                        return D0_DC1((d_974_v108_)[default__.safeIndex(491, (d_974_v108_).length(0))])

                    return lambda46_

                init22_ = lambda45_(d_967_v108_)
                nw172_ = _dafny.Array(None, 7)
                for i0_22_ in range(nw172_.length(0)):
                    nw172_[i0_22_] = init22_(i0_22_)
                d_973_v112_ = nw172_
                index167_ = default__.safeIndex(772, (d_972_v111_).length(0))
                (d_972_v111_)[index167_] = d_973_v112_
                d_976_v113_: _dafny.Array
                d_976_v113_ = d_973_v112_
                d_977_v114_: _dafny.Seq
                d_977_v114_ = _dafny.SeqWithoutIsStrInference([d_973_v112_, d_973_v112_, (d_976_v113_)])
                index168_ = default__.safeIndex(772, (d_972_v111_).length(0))
                (d_972_v111_)[index168_] = (d_977_v114_)[default__.safeIndex((0) - (p1), len(d_977_v114_))]
            d_978_v115_: _dafny.Map
            d_978_v115_ = _dafny.Map({d_828_v0_: True})
            d_979_v116_: _dafny.Seq
            d_979_v116_ = _dafny.SeqWithoutIsStrInference([d_828_v0_])
            d_980_v117_: _dafny.Seq
            d_980_v117_ = _dafny.SeqWithoutIsStrInference([len(d_978_v115_), (self).f0])
            if ((d_978_v115_)[not((d_979_v116_)[default__.safeIndex(len(d_980_v117_), len(d_979_v116_))])] if (not((d_979_v116_)[default__.safeIndex(len(d_980_v117_), len(d_979_v116_))])) in (d_978_v115_) else False):
                d_828_v0_ = d_828_v0_
                d_828_v0_ = d_828_v0_
                d_828_v0_ = d_828_v0_
                d_981_v118_: _dafny.Seq
                d_981_v118_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qp"))
                d_982_v119_: D0
                d_982_v119_ = D0_DC0(False)
                d_983_v120_: str
                d_983_v120_ = _dafny.CodePoint('r')
                d_981_v118_ = (d_981_v118_).set(default__.safeIndex(default__.safeDivisionInt((D3_DC8((d_982_v119_).cf0, len(_dafny.Map({((d_829_v1_)[d_828_v0_] if (d_828_v0_) in (d_829_v1_) else p1): (self).f0})))).cf9, len(d_981_v118_)), len(d_981_v118_)), d_983_v120_)
                d_984_v121_: bool
                d_985_v122_: int
                out19_: bool
                out20_: int
                out19_, out20_ = (d_913_v64_).m6(default__.safeModuloInt(p1, (self).f0), globalState)
                d_984_v121_ = out19_
                d_985_v122_ = out20_
            elif True:
                d_986_v123_: _dafny.Array
                def lambda47_(d_987_i10_):
                    return _dafny.CodePoint('t')

                init23_ = lambda47_
                nw173_ = _dafny.Array(None, 17)
                for i0_23_ in range(nw173_.length(0)):
                    nw173_[i0_23_] = init23_(i0_23_)
                d_986_v123_ = nw173_
                d_988_v124_: str
                d_988_v124_ = _dafny.CodePoint('l')
                d_989_v125_: str
                d_989_v125_ = d_988_v124_
                index169_ = default__.safeIndex(579, (d_986_v123_).length(0))
                (d_986_v123_)[index169_] = d_989_v125_
                d_990_v126_: _dafny.Set
                d_990_v126_ = _dafny.Set({(self).f0, (self).f0})
                index170_ = default__.safeIndex(579, (d_986_v123_).length(0))
                rhs158_ = d_989_v125_
                rhs159_ = ((((d_829_v1_).set(d_828_v0_, default__.abs((d_980_v117_)[default__.safeIndex(p1, len(d_980_v117_))]))).set(d_828_v0_, default__.abs(len(d_980_v117_)))).cardinality) * (p1)
                rhs160_ = (len((d_979_v116_) + (d_979_v116_))) > ((self).f0)
                rhs161_ = d_980_v117_
                rhs162_ = (d_990_v126_).intersection((_dafny.Set({(self).f0})) - (d_990_v126_))
                lhs74_ = d_986_v123_
                lhs75_ = default__.safeIndex(579, (d_986_v123_).length(0))
                lhs74_[lhs75_] = rhs158_
                r0 = rhs159_
                d_828_v0_ = rhs160_
                d_980_v117_ = rhs161_
                d_990_v126_ = rhs162_
                d_991_v127_: _dafny.Array
                nw174_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_991_v127_ = nw174_
                d_992_v128_: _dafny.Map
                d_992_v128_ = _dafny.Map({-949: p1})
                d_993_v129_: D7
                d_993_v129_ = D7_DC17(d_828_v0_, d_828_v0_, d_828_v0_)
                pat_let_tv27_ = d_828_v0_
                d_994_v130_: _dafny.Map
                def iife67_(_pat_let21_0):
                    def iife68_(d_995_dt__update__tmp_h4_):
                        def iife69_(_pat_let22_0):
                            def iife70_(d_996_dt__update_hcf27_h0_):
                                return D7_DC17(d_996_dt__update_hcf27_h0_, (d_995_dt__update__tmp_h4_).cf28, (d_995_dt__update__tmp_h4_).cf29)
                            return iife70_(_pat_let22_0)
                        return iife69_(pat_let_tv27_)
                    return iife68_(_pat_let21_0)
                d_994_v130_ = _dafny.Map({iife67_(d_993_v129_): (p1) + (((d_992_v128_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "puclwjroo")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "puclwjroo")))) in (d_992_v128_) else 967))})
                d_997_v131_: _dafny.Seq
                d_997_v131_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "msay"))
                d_998_v132_: _dafny.Map
                d_998_v132_ = _dafny.Map({d_828_v0_: d_997_v131_})
                d_999_v133_: _dafny.Set
                d_999_v133_ = _dafny.Set({d_828_v0_, d_828_v0_})
                rhs163_ = d_991_v127_
                rhs164_ = d_992_v128_
                rhs165_ = ((p1 if d_828_v0_ else (self).f0)) + (p1)
                rhs166_ = _dafny.Map({d_993_v129_: len((((d_998_v132_)[d_828_v0_] if (d_828_v0_) in (d_998_v132_) else d_997_v131_)).set(default__.safeIndex(p1, len(((d_998_v132_)[d_828_v0_] if (d_828_v0_) in (d_998_v132_) else d_997_v131_))), d_988_v124_))})
                rhs167_ = default__.fm3(d_999_v133_, d_988_v124_, p1, p1, globalState)
                d_991_v127_ = rhs163_
                d_992_v128_ = rhs164_
                r2 = rhs165_
                d_994_v130_ = rhs166_
                r1 = rhs167_
                d_1000_v134_: C2
                nw175_ = C2()
                nw175_.ctor__((self).f0)
                d_1000_v134_ = nw175_
                d_1001_v135_: _dafny.Seq
                d_1001_v135_ = _dafny.SeqWithoutIsStrInference([d_1000_v134_, d_1000_v134_])
                d_1002_v136_: _dafny.Array
                nw176_ = _dafny.Array(int(0), 22)
                d_1002_v136_ = nw176_
                rhs168_ = ((d_961_v103_).cf4 if d_828_v0_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmljar")))
                rhs169_ = (d_979_v116_)[default__.safeIndex(p1, len(d_979_v116_))]
                rhs170_ = d_1001_v135_
                rhs171_ = d_1002_v136_
                rhs172_ = d_828_v0_
                d_997_v131_ = rhs168_
                d_828_v0_ = rhs169_
                d_1001_v135_ = rhs170_
                d_1002_v136_ = rhs171_
                d_828_v0_ = rhs172_
                d_990_v126_ = d_990_v126_
                d_828_v0_ = d_828_v0_
        d_1003_v137_: _dafny.Seq
        d_1003_v137_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1003_v137_ = d_1003_v137_
        if (d_828_v0_ if d_828_v0_ else d_828_v0_):
            d_1004_v138_: _dafny.Set
            d_1004_v138_ = _dafny.Set({False, d_828_v0_})
            d_1005_v139_: C3
            nw177_ = C3()
            nw177_.ctor__((0) - (len((d_1004_v138_).intersection(d_1004_v138_))))
            d_1005_v139_ = nw177_
            d_1006_v140_: str
            d_1006_v140_ = _dafny.CodePoint('v')
            d_1007_v141_: C0
            nw178_ = C0()
            nw178_.ctor__(d_1006_v140_)
            d_1007_v141_ = nw178_
            if d_828_v0_:
                d_1008_v142_: _dafny.MultiSet
                d_1008_v142_ = _dafny.MultiSet([(self).f0])
                d_1009_v143_: _dafny.Map
                d_1009_v143_ = _dafny.Map({default__.fm19(d_1008_v142_, p1, (self).f0, (self).f0, globalState): True})
                d_1009_v143_ = (d_1009_v143_).set(default__.fm19(d_1008_v142_, p1, len(d_1003_v137_), (self).f0, globalState), not(d_828_v0_))
                d_1010_v144_: _dafny.Seq
                d_1010_v144_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwm"))
                d_1010_v144_ = _dafny.SeqWithoutIsStrInference([d_1006_v140_ for d_1011_i11_ in range(default__.abs(9))])
                d_1012_v145_: C1
                nw179_ = C1()
                nw179_.ctor__((self).f0)
                d_1012_v145_ = nw179_
                d_1013_v146_: C1
                nw180_ = C1()
                nw180_.ctor__((0) - ((self).f0))
                d_1013_v146_ = nw180_
                d_828_v0_ = (_dafny.MultiSet([d_828_v0_, d_828_v0_, d_828_v0_, d_828_v0_])) != ((d_829_v1_) - (d_829_v1_))
            elif True:
                d_1014_v147_: _dafny.Seq
                d_1014_v147_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iltk"))
                d_1014_v147_ = ((_dafny.SeqWithoutIsStrInference([d_1006_v140_ for d_1015_i12_ in range(default__.abs(310))])) + (d_1014_v147_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gclpd")))
                d_1016_v148_: _dafny.Array
                nw181_ = _dafny.Array(False, 24)
                d_1016_v148_ = nw181_
                d_1017_v149_: _dafny.Map
                d_1017_v149_ = _dafny.Map({(d_828_v0_) and (d_828_v0_): d_1016_v148_})
                d_1017_v149_ = (d_1017_v149_).set((D3_DC8(d_828_v0_, p1)).cf8, d_1016_v148_)
                r0 = (p1) * (len((_dafny.SeqWithoutIsStrInference([d_1006_v140_, (d_1007_v141_).f3])) + (d_1014_v147_)))
                d_1018_v150_: C1
                nw182_ = C1()
                nw182_.ctor__((self).f0)
                d_1018_v150_ = nw182_
                r2 = default__.safeDivisionInt((0) - ((self).f0), (self).f0)
            d_828_v0_ = d_828_v0_
            d_828_v0_ = d_828_v0_
        elif True:
            d_828_v0_ = d_828_v0_
            d_828_v0_ = d_828_v0_
            d_1019_v151_: _dafny.Seq
            d_1019_v151_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nny"))
            d_1020_v152_: str
            d_1020_v152_ = _dafny.CodePoint('d')
            d_1021_v153_: _dafny.Set
            d_1021_v153_ = _dafny.Set({d_828_v0_, d_828_v0_})
            d_1022_v154_: _dafny.MultiSet
            d_1022_v154_ = _dafny.MultiSet([len(d_1003_v137_), p1])
            d_1023_v155_: D2
            d_1023_v155_ = D2_DC4((d_1019_v151_).set(default__.safeIndex((self).f0, len(d_1019_v151_)), d_1020_v152_), ((self).f0) <= (default__.fm3(d_1021_v153_, d_1020_v152_, (self).f0, (d_1022_v154_).cardinality, globalState)))
            d_1023_v155_ = d_1023_v155_
            d_1024_v156_: C2
            nw183_ = C2()
            nw183_.ctor__(p1)
            d_1024_v156_ = nw183_
            nw184_ = C2()
            nw184_.ctor__((p1) * (p1))
            d_1024_v156_ = nw184_
            d_1003_v137_ = d_1003_v137_
        d_1025_v157_: _dafny.Set
        d_1025_v157_ = _dafny.Set({True})
        d_1026_v158_: _dafny.Seq
        d_1026_v158_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kw"))
        hi6_ = default__.safeModuloInt(len(_dafny.Map({d_1025_v157_: len(d_1026_v158_)})), p1)
        for d_1027_i13_ in range((d_1003_v137_)[default__.safeIndex((self).f0, len(d_1003_v137_))], hi6_):
            d_828_v0_ = d_828_v0_
            r0 = default__.safeModuloInt((self).f0, len(d_1025_v157_))
            d_828_v0_ = not((len(d_1003_v137_)) > ((_dafny.MultiSet([p1, (self).f0])).cardinality))
            r2 = (self).f0
        d_1028_v159_: _dafny.Seq
        d_1028_v159_ = _dafny.SeqWithoutIsStrInference([(d_1026_v158_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))])
        d_1028_v159_ = d_1028_v159_
        d_1029_v160_: str
        d_1029_v160_ = _dafny.CodePoint('s')
        d_1030_v161_: _dafny.Set
        d_1030_v161_ = _dafny.Set({d_1029_v160_, d_1029_v160_, d_1029_v160_, _dafny.CodePoint('g'), _dafny.CodePoint('h')})
        r0 = default__.safeModuloInt((924) + (p1), default__.safeModuloInt((self).f0, len(d_1030_v161_)))
        r1 = (self).f0
        d_1031_v162_: D0
        d_1031_v162_ = D0_DC1(d_828_v0_)
        d_1032_v163_: _dafny.Seq
        d_1032_v163_ = _dafny.SeqWithoutIsStrInference([d_1031_v162_])
        d_1033_v164_: D2
        d_1033_v164_ = D2_DC3(d_1032_v163_)
        pat_let_tv28_ = d_1026_v158_
        pat_let_tv29_ = p1
        pat_let_tv30_ = p1
        def lambda48_(source12_):
            if source12_.is_DC4:
                d_1034___mcc_h6_ = source12_.cf4
                d_1035___mcc_h7_ = source12_.cf5
                d_1036_cf5_ = d_1035___mcc_h7_
                d_1037_cf4_ = d_1034___mcc_h6_
                return len(_dafny.Map({(D2_DC4(pat_let_tv28_, False)).cf5: (0) - (pat_let_tv29_)}))
            elif source12_.is_DC5:
                return pat_let_tv30_
            elif source12_.is_DC3:
                d_1038___mcc_h8_ = source12_.cf3
                d_1039_cf3_ = d_1038___mcc_h8_
                return default__.safeModuloInt((self).f0, (self).f0)
            elif True:
                d_1040___mcc_h9_ = source12_.cf6
                d_1041_cf6_ = d_1040___mcc_h9_
                return (self).f0

        r2 = lambda48_(d_1033_v164_)
        return r0, r1, r2

    def m4(self, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_1042_v0_: _dafny.Array
        def lambda49_(d_1043_i0_):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kotlm"))

        init24_ = lambda49_
        nw185_ = _dafny.Array(None, 19)
        for i0_24_ in range(nw185_.length(0)):
            nw185_[i0_24_] = init24_(i0_24_)
        d_1042_v0_ = nw185_
        d_1044_v1_: _dafny.Seq
        d_1044_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxch"))
        index171_ = default__.safeIndex(383, (d_1042_v0_).length(0))
        (d_1042_v0_)[index171_] = d_1044_v1_
        d_1045_v2_: C2
        nw186_ = C2()
        nw186_.ctor__(282)
        d_1045_v2_ = nw186_
        d_1046_v3_: _dafny.MultiSet
        d_1046_v3_ = _dafny.MultiSet([d_1045_v2_])
        d_1047_v4_: bool
        d_1047_v4_ = False
        d_1048_v5_: D2
        d_1048_v5_ = D2_DC4(d_1044_v1_, d_1047_v4_)
        index172_ = default__.safeIndex(383, (d_1042_v0_).length(0))
        (d_1042_v0_)[index172_] = ((default__.fm18((self).f0, globalState)) + (d_1044_v1_) if (d_1045_v2_) in (d_1046_v3_) else (d_1048_v5_).cf4)
        hi7_ = ((self).f0) + (-45)
        for d_1049_i1_ in range(default__.safeDivisionInt((self).f0, (0) - ((self).f0)), hi7_):
            d_1044_v1_ = (d_1042_v0_)[default__.safeIndex(383, (d_1042_v0_).length(0))]
            d_1050_v6_: _dafny.Array
            def lambda50_(d_1051_v4_):
                def lambda51_(d_1052_i2_):
                    return d_1051_v4_

                return lambda51_

            init25_ = lambda50_(d_1047_v4_)
            nw187_ = _dafny.Array(None, 25)
            for i0_25_ in range(nw187_.length(0)):
                nw187_[i0_25_] = init25_(i0_25_)
            d_1050_v6_ = nw187_
            index173_ = default__.safeIndex(232, (d_1050_v6_).length(0))
            (d_1050_v6_)[index173_] = d_1047_v4_
            d_1053_v7_: _dafny.Seq
            d_1053_v7_ = _dafny.SeqWithoutIsStrInference([d_1047_v4_])
            d_1054_v8_: _dafny.MultiSet
            d_1054_v8_ = _dafny.MultiSet([d_1047_v4_, d_1047_v4_])
            index174_ = default__.safeIndex(232, (d_1050_v6_).length(0))
            (d_1050_v6_)[index174_] = (_dafny.MultiSet(((d_1053_v7_) + (d_1053_v7_)).set(default__.safeIndex(d_1049_i1_, len((d_1053_v7_) + (d_1053_v7_))), d_1047_v4_))).isdisjoint((d_1054_v8_ if d_1047_v4_ else d_1054_v8_))
            index175_ = default__.safeIndex(232, (d_1050_v6_).length(0))
            (d_1050_v6_)[index175_] = (d_1045_v2_).fm5(globalState)
            d_1047_v4_ = not(d_1047_v4_)
        d_1055_v9_: _dafny.MultiSet
        d_1055_v9_ = _dafny.MultiSet([(self).f0])
        d_1056_v10_: _dafny.Seq
        d_1056_v10_ = _dafny.SeqWithoutIsStrInference([d_1047_v4_])
        d_1057_v11_: _dafny.Map
        d_1057_v11_ = _dafny.Map({_dafny.Map({len(d_1056_v10_): not(d_1047_v4_)}): d_1047_v4_})
        d_1058_v12_: _dafny.Map
        d_1058_v12_ = _dafny.Map({904: False})
        index176_ = default__.safeIndex(383, (d_1042_v0_).length(0))
        (d_1042_v0_)[index176_] = default__.fm19(d_1055_v9_, (self).f0, (len(d_1057_v11_)) + ((self).f0), len(d_1058_v12_), globalState)
        if (d_1045_v2_).fm6(d_1047_v4_, d_1044_v1_, d_1044_v1_, ((self).f0) + ((self).f0), globalState):
            d_1059_v13_: int
            d_1059_v13_ = 47
            d_1059_v13_ = d_1059_v13_
            if d_1047_v4_:
                d_1060_v14_: _dafny.Array
                nw188_ = _dafny.Array(_dafny.Set({}), 28)
                d_1060_v14_ = nw188_
                d_1061_v15_: _dafny.Set
                d_1061_v15_ = _dafny.Set({636})
                index177_ = default__.safeIndex(823, (d_1060_v14_).length(0))
                (d_1060_v14_)[index177_] = d_1061_v15_
                index178_ = default__.safeIndex(823, (d_1060_v14_).length(0))
                (d_1060_v14_)[index178_] = (d_1061_v15_).intersection((d_1061_v15_ if d_1047_v4_ else d_1061_v15_))
                d_1062_v16_: C3
                nw189_ = C3()
                nw189_.ctor__(d_1059_v13_)
                d_1062_v16_ = nw189_
                d_1063_v17_: _dafny.Set
                d_1063_v17_ = _dafny.Set({_dafny.Set({d_1047_v4_, d_1047_v4_})})
                d_1064_v18_: _dafny.Map
                d_1064_v18_ = _dafny.Map({d_1062_v16_: (d_1063_v17_) | (d_1063_v17_)})
                d_1064_v18_ = (d_1064_v18_).set(d_1062_v16_, d_1063_v17_)
                d_1047_v4_ = (d_1047_v4_) or (not(d_1047_v4_))
                d_1047_v4_ = d_1047_v4_
                d_1059_v13_ = 976
            elif True:
                d_1047_v4_ = d_1047_v4_
                d_1065_v19_: _dafny.Map
                d_1065_v19_ = _dafny.Map({d_1047_v4_: (0) - (d_1059_v13_)})
                d_1066_v20_: _dafny.Map
                d_1066_v20_ = _dafny.Map({d_1047_v4_: d_1065_v19_})
                d_1066_v20_ = (d_1066_v20_).set(d_1047_v4_, _dafny.Map({d_1047_v4_: len(d_1044_v1_)}))
                d_1067_v21_: C3
                nw190_ = C3()
                nw190_.ctor__(len((d_1044_v1_) + (d_1044_v1_)))
                d_1067_v21_ = nw190_
                d_1068_v22_: str
                d_1068_v22_ = _dafny.CodePoint('x')
                d_1069_v23_: _dafny.Map
                d_1069_v23_ = _dafny.Map({d_1059_v13_: d_1068_v22_})
                d_1070_v24_: _dafny.MultiSet
                d_1070_v24_ = _dafny.MultiSet([d_1069_v23_, d_1069_v23_, _dafny.Map({(self).f0: d_1068_v22_}), d_1069_v23_, _dafny.Map({(self).f0: d_1068_v22_})])
                d_1071_v25_: _dafny.Seq
                d_1071_v25_ = _dafny.SeqWithoutIsStrInference([d_1059_v13_, (0) - (d_1059_v13_), d_1059_v13_, d_1059_v13_, d_1059_v13_])
                d_1072_v26_: _dafny.Map
                d_1072_v26_ = _dafny.Map({(d_1070_v24_) | (d_1070_v24_): (d_1071_v25_)[default__.safeIndex(841, len(d_1071_v25_))]})
                d_1072_v26_ = (d_1072_v26_).set((d_1070_v24_) - (_dafny.MultiSet([d_1069_v23_])), (self).f0)
                d_1073_v27_: _dafny.Array
                nw191_ = _dafny.Array(int(0), 6)
                d_1073_v27_ = nw191_
                index179_ = default__.safeIndex(769, (d_1073_v27_).length(0))
                (d_1073_v27_)[index179_] = d_1059_v13_
                index180_ = default__.safeIndex(66, (d_1073_v27_).length(0))
                (d_1073_v27_)[index180_] = (d_1059_v13_) * ((self).f0)
                d_1074_v28_: _dafny.MultiSet
                d_1074_v28_ = _dafny.MultiSet([d_1047_v4_, d_1047_v4_])
                d_1075_v29_: _dafny.Map
                d_1075_v29_ = _dafny.Map({d_1068_v22_: ((d_1074_v28_).set((d_1067_v21_).fm5(globalState), default__.abs((self).f0))).ispropersubset(d_1074_v28_)})
                d_1076_v30_: _dafny.Seq
                d_1076_v30_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1047_v4_]), d_1056_v10_])
                d_1077_v31_: _dafny.Array
                nw192_ = _dafny.Array(False, 22)
                d_1077_v31_ = nw192_
                d_1078_v32_: _dafny.Set
                d_1078_v32_ = _dafny.Set({d_1077_v31_})
                index181_ = default__.safeIndex(769, (d_1073_v27_).length(0))
                index182_ = default__.safeIndex(66, (d_1073_v27_).length(0))
                rhs173_ = ((d_1056_v10_) + ((d_1076_v30_)[default__.safeIndex((d_1074_v28_).cardinality, len(d_1076_v30_))])) + (d_1056_v10_)
                rhs174_ = (0) - (default__.safeDivisionInt(len(d_1078_v32_), d_1059_v13_))
                rhs175_ = (d_1059_v13_) + ((self).f0)
                rhs176_ = d_1075_v29_
                rhs177_ = (0) - ((d_1059_v13_) * ((self).f0))
                lhs76_ = d_1073_v27_
                lhs77_ = default__.safeIndex(769, (d_1073_v27_).length(0))
                lhs78_ = d_1073_v27_
                lhs79_ = default__.safeIndex(66, (d_1073_v27_).length(0))
                d_1056_v10_ = rhs173_
                lhs76_[lhs77_] = rhs174_
                lhs78_[lhs79_] = rhs175_
                d_1075_v29_ = rhs176_
                d_1059_v13_ = rhs177_
            rhs178_ = d_1047_v4_
            rhs179_ = d_1047_v4_
            rhs180_ = d_1047_v4_
            rhs181_ = -963
            d_1047_v4_ = rhs178_
            d_1047_v4_ = rhs179_
            d_1047_v4_ = rhs180_
            d_1059_v13_ = rhs181_
            d_1079_v33_: str
            d_1079_v33_ = _dafny.CodePoint('s')
            d_1080_v34_: C0
            nw193_ = C0()
            nw193_.ctor__(d_1079_v33_)
            d_1080_v34_ = nw193_
            d_1080_v34_ = d_1080_v34_
            d_1059_v13_ = default__.safeDivisionInt((self).f0, (self).f0)
        elif True:
            d_1081_v35_: _dafny.Map
            d_1081_v35_ = _dafny.Map({((self).f0) < ((self).f0): (self).f0})
            d_1081_v35_ = (d_1081_v35_).set((False) or (d_1047_v4_), (self).f0)
            d_1082_v36_: _dafny.Array
            nw194_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_1082_v36_ = nw194_
            d_1083_v37_: int
            d_1084_v38_: int
            d_1085_v39_: int
            out21_: int
            out22_: int
            out23_: int
            out21_, out22_, out23_ = (d_1045_v2_).m1(d_1082_v36_, (self).f0, globalState)
            d_1083_v37_ = out21_
            d_1084_v38_ = out22_
            d_1085_v39_ = out23_
            d_1086_v40_: str
            d_1086_v40_ = _dafny.CodePoint('y')
            d_1086_v40_ = d_1086_v40_
            d_1087_v41_: _dafny.Array
            nw195_ = _dafny.Array(_dafny.Map({}), 4)
            d_1087_v41_ = nw195_
            d_1088_v42_: T0
            nw196_ = C3()
            nw196_.ctor__((self).f0)
            d_1088_v42_ = nw196_
            d_1089_v43_: D10
            d_1089_v43_ = D10_DC21(d_1088_v42_)
            d_1090_v44_: _dafny.Array
            nw197_ = _dafny.Array(None, 7)
            nw197_[int(0)] = d_1047_v4_
            nw197_[int(1)] = d_1047_v4_
            nw197_[int(2)] = d_1047_v4_
            nw197_[int(3)] = False
            nw197_[int(4)] = not(False)
            nw197_[int(5)] = d_1047_v4_
            nw197_[int(6)] = d_1047_v4_
            d_1090_v44_ = nw197_
            d_1091_v45_: _dafny.Map
            d_1091_v45_ = _dafny.Map({d_1089_v43_: d_1090_v44_})
            index183_ = default__.safeIndex(860, (d_1087_v41_).length(0))
            (d_1087_v41_)[index183_] = d_1091_v45_
            index184_ = default__.safeIndex(860, (d_1087_v41_).length(0))
            (d_1087_v41_)[index184_] = _dafny.Map({d_1089_v43_: d_1090_v44_})
            d_1092_v46_: _dafny.Array
            nw198_ = _dafny.Array(int(0), 10)
            d_1092_v46_ = nw198_
            index185_ = default__.safeIndex(378, (d_1092_v46_).length(0))
            (d_1092_v46_)[index185_] = (0) - (d_1085_v39_)
            d_1093_v47_: _dafny.Map
            d_1093_v47_ = _dafny.Map({d_1047_v4_: d_1047_v4_})
            index186_ = default__.safeIndex(378, (d_1092_v46_).length(0))
            (d_1092_v46_)[index186_] = ((0) - (len((d_1093_v47_ if not(d_1047_v4_) else _dafny.Map({d_1047_v4_: d_1047_v4_}))))) * (d_1084_v38_)
        d_1094_v48_: int
        d_1094_v48_ = -963
        d_1094_v48_ = (0) - ((self).f0)
        d_1095_v49_: str
        d_1095_v49_ = _dafny.CodePoint('t')
        d_1096_v50_: _dafny.Map
        d_1096_v50_ = _dafny.Map({d_1095_v49_: d_1094_v48_})
        hi8_ = len((d_1096_v50_) | (d_1096_v50_))
        for d_1097_i3_ in range((self).f0, hi8_):
            d_1098_v51_: C3
            nw199_ = C3()
            nw199_.ctor__(920)
            d_1098_v51_ = nw199_
            d_1099_v52_: _dafny.Seq
            d_1099_v52_ = _dafny.SeqWithoutIsStrInference([d_1097_i3_])
            d_1100_v53_: _dafny.Seq
            d_1100_v53_ = _dafny.SeqWithoutIsStrInference([d_1099_v52_, d_1099_v52_, d_1099_v52_, d_1099_v52_, d_1099_v52_])
            d_1101_v54_: _dafny.Map
            d_1101_v54_ = _dafny.Map({d_1100_v53_: _dafny.SeqWithoutIsStrInference([d_1094_v48_, 847, (self).f0])})
            d_1101_v54_ = (d_1101_v54_).set((d_1100_v53_).set(default__.safeIndex(d_1097_i3_, len(d_1100_v53_)), d_1099_v52_), d_1099_v52_)
            d_1094_v48_ = (d_1094_v48_) - ((self).f0)
            d_1102_v55_: _dafny.Array
            nw200_ = _dafny.Array(None, 2)
            nw200_[int(0)] = 11
            nw200_[int(1)] = (d_1097_i3_) * (d_1094_v48_)
            d_1102_v55_ = nw200_
            index187_ = default__.safeIndex(495, (d_1102_v55_).length(0))
            (d_1102_v55_)[index187_] = d_1097_i3_
            d_1103_v56_: _dafny.Set
            d_1103_v56_ = _dafny.Set({d_1047_v4_, d_1047_v4_})
            index188_ = default__.safeIndex(495, (d_1102_v55_).length(0))
            (d_1102_v55_)[index188_] = default__.safeModuloInt(default__.fm3(d_1103_v56_, d_1095_v49_, (self).f0, len(d_1056_v10_), globalState), d_1094_v48_)
        d_1104_v57_: _dafny.Map
        d_1104_v57_ = _dafny.Map({d_1047_v4_: d_1047_v4_})
        d_1105_v58_: _dafny.Seq
        d_1105_v58_ = _dafny.SeqWithoutIsStrInference([d_1104_v57_])
        d_1106_v59_: _dafny.Seq
        d_1106_v59_ = _dafny.SeqWithoutIsStrInference([d_1055_v9_])
        d_1107_v60_: _dafny.Seq
        d_1107_v60_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecqjshdfv"))): d_1094_v48_}))])
        d_1108_v61_: _dafny.MultiSet
        d_1108_v61_ = _dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f0])), _dafny.MultiSet([len(d_1107_v60_), d_1094_v48_])])
        r0 = ((d_1104_v57_) | ((d_1105_v58_)[default__.safeIndex((self).f0, len(d_1105_v58_))])).set(d_1047_v4_, (_dafny.MultiSet(d_1106_v59_)).isdisjoint(d_1108_v61_))
        return r0

    def m5(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: _dafny.MultiSet = _dafny.MultiSet({})
        d_1109_v0_: int
        d_1109_v0_ = 565
        d_1109_v0_ = (self).f0
        d_1110_v1_: str
        d_1110_v1_ = _dafny.CodePoint('g')
        d_1110_v1_ = d_1110_v1_
        d_1109_v0_ = d_1109_v0_
        if p0:
            d_1111_v2_: C0
            nw201_ = C0()
            nw201_.ctor__(d_1110_v1_)
            d_1111_v2_ = nw201_
            d_1112_v3_: _dafny.Set
            d_1112_v3_ = _dafny.Set({p1, p1})
            d_1113_v4_: _dafny.Seq
            d_1113_v4_ = _dafny.SeqWithoutIsStrInference([len(d_1112_v3_)])
            d_1114_v5_: D0
            d_1114_v5_ = D0_DC1(True)
            d_1115_v6_: _dafny.Seq
            d_1115_v6_ = _dafny.SeqWithoutIsStrInference([D0_DC1(default__.fm2(p1, len(d_1113_v4_), p0, p1, globalState)), d_1114_v5_, d_1114_v5_])
            d_1116_v7_: D2
            d_1116_v7_ = D2_DC3(d_1115_v6_)
            d_1117_v8_: D2
            d_1117_v8_ = D2_DC6((d_1116_v7_ if p1 else D2_DC6(d_1116_v7_)))
            d_1117_v8_ = d_1117_v8_
            d_1109_v0_ = (d_1109_v0_) * (d_1109_v0_)
            d_1109_v0_ = (default__.safeModuloInt((self).f0, (0) - (default__.fm3(d_1112_v3_, (d_1111_v2_).f3, (d_1113_v4_)[default__.safeIndex((self).f0, len(d_1113_v4_))], 759, globalState)))) + (d_1109_v0_)
            d_1118_v9_: C0
            nw202_ = C0()
            nw202_.ctor__((d_1111_v2_).f3)
            d_1118_v9_ = nw202_
        elif True:
            d_1119_v10_: _dafny.Array
            def lambda52_(d_1120_i0_):
                return (d_1120_i0_) + ((self).f0)

            init26_ = lambda52_
            nw203_ = _dafny.Array(None, 20)
            for i0_26_ in range(nw203_.length(0)):
                nw203_[i0_26_] = init26_(i0_26_)
            d_1119_v10_ = nw203_
            d_1121_v11_: _dafny.Map
            d_1121_v11_ = _dafny.Map({301: (self).f0})
            d_1122_v12_: _dafny.Seq
            d_1122_v12_ = _dafny.SeqWithoutIsStrInference([d_1110_v1_, d_1110_v1_])
            d_1123_v13_: _dafny.Seq
            d_1123_v13_ = _dafny.SeqWithoutIsStrInference([len(d_1121_v11_), (self).f0, (0) - (len(d_1122_v12_))])
            d_1124_v14_: _dafny.Seq
            d_1124_v14_ = _dafny.SeqWithoutIsStrInference([p1])
            index189_ = default__.safeIndex(990, (d_1119_v10_).length(0))
            (d_1119_v10_)[index189_] = ((d_1123_v13_)[default__.safeIndex(len(d_1124_v14_), len(d_1123_v13_))]) - (d_1109_v0_)
            d_1125_v15_: D3
            d_1125_v15_ = D3_DC7(d_1123_v13_)
            index190_ = default__.safeIndex(990, (d_1119_v10_).length(0))
            rhs182_ = len(_dafny.SeqWithoutIsStrInference([(d_1125_v15_).cf7, d_1123_v13_]))
            rhs183_ = d_1109_v0_
            rhs184_ = d_1122_v12_
            lhs80_ = d_1119_v10_
            lhs81_ = default__.safeIndex(990, (d_1119_v10_).length(0))
            lhs80_[lhs81_] = rhs182_
            d_1109_v0_ = rhs183_
            d_1122_v12_ = rhs184_
            d_1126_v16_: _dafny.Array
            nw204_ = _dafny.Array(None, 18)
            nw204_[int(0)] = p1
            nw204_[int(1)] = p0
            nw204_[int(2)] = True
            nw204_[int(3)] = p1
            nw204_[int(4)] = p1
            nw204_[int(5)] = p1
            nw204_[int(6)] = (self).fm6(p0, d_1122_v12_, d_1122_v12_, d_1109_v0_, globalState)
            nw204_[int(7)] = p0
            nw204_[int(8)] = p0
            nw204_[int(9)] = not(p0)
            nw204_[int(10)] = p0
            nw204_[int(11)] = default__.fm2(p1, d_1109_v0_, True, p1, globalState)
            nw204_[int(12)] = p0
            nw204_[int(13)] = p1
            nw204_[int(14)] = (p1) and (p1)
            nw204_[int(15)] = p0
            nw204_[int(16)] = p0
            nw204_[int(17)] = p0
            d_1126_v16_ = nw204_
            r1 = d_1126_v16_
            r0 = not (p0) or (p1)
            index191_ = default__.safeIndex(990, (d_1119_v10_).length(0))
            (d_1119_v10_)[index191_] = ((d_1119_v10_)[default__.safeIndex(990, (d_1119_v10_).length(0))]) + (d_1109_v0_)
            d_1127_v17_: D7
            d_1127_v17_ = D7_DC17(not(p1), not(p1), p1)
            source13_ = d_1127_v17_
            if source13_.is_DC17:
                d_1128___mcc_h0_ = source13_.cf27
                d_1129___mcc_h1_ = source13_.cf28
                d_1130___mcc_h2_ = source13_.cf29
                d_1131_cf29_ = d_1130___mcc_h2_
                d_1132_cf28_ = d_1129___mcc_h1_
                d_1133_cf27_ = d_1128___mcc_h0_
                d_1134_v18_: _dafny.Map
                d_1134_v18_ = _dafny.Map({d_1122_v12_: d_1126_v16_})
                d_1134_v18_ = (d_1134_v18_) | (d_1134_v18_)
                d_1135_v19_: _dafny.Array
                nw205_ = _dafny.Array(None, 8)
                d_1135_v19_ = nw205_
                d_1136_v20_: D3
                d_1136_v20_ = D3_DC9(d_1109_v0_, d_1135_v19_, p0)
                d_1131_cf29_ = (d_1136_v20_).cf12
                d_1119_v10_ = d_1119_v10_
                d_1137_v21_: _dafny.Map
                d_1137_v21_ = _dafny.Map({(_dafny.Map({not(d_1133_cf27_): len(_dafny.SeqWithoutIsStrInference([(self).f0]))})).set(d_1132_cf28_, (d_1119_v10_)[default__.safeIndex(990, (d_1119_v10_).length(0))]): (d_1119_v10_)[default__.safeIndex(990, (d_1119_v10_).length(0))]})
                d_1138_v22_: D7
                d_1138_v22_ = D7_DC18((d_1119_v10_)[default__.safeIndex(990, (d_1119_v10_).length(0))])
                d_1139_v23_: _dafny.Map
                d_1139_v23_ = _dafny.Map({d_1138_v22_: d_1109_v0_})
                d_1140_v24_: T1
                nw206_ = C2()
                nw206_.ctor__(d_1109_v0_)
                d_1140_v24_ = nw206_
                d_1141_v25_: _dafny.Map
                d_1141_v25_ = _dafny.Map({d_1140_v24_: p0})
                d_1142_v26_: _dafny.Map
                d_1142_v26_ = _dafny.Map({p0: len((d_1139_v23_).set(d_1138_v22_, len(d_1141_v25_)))})
                d_1143_v27_: _dafny.Map
                d_1143_v27_ = _dafny.Map({((d_1137_v21_)[d_1142_v26_] if (d_1142_v26_) in (d_1137_v21_) else (d_1119_v10_)[default__.safeIndex(990, (d_1119_v10_).length(0))]): d_1122_v12_})
                d_1143_v27_ = (d_1143_v27_).set((d_1119_v10_)[default__.safeIndex(990, (d_1119_v10_).length(0))], d_1122_v12_)
            elif source13_.is_DC18:
                d_1144___mcc_h3_ = source13_.cf30
                d_1145_cf30_ = d_1144___mcc_h3_
                d_1146_v28_: _dafny.Map
                d_1146_v28_ = _dafny.Map({p0: p1})
                d_1146_v28_ = (d_1146_v28_).set(not (p0) or (p0), p1)
                index192_ = default__.safeIndex(990, (d_1119_v10_).length(0))
                (d_1119_v10_)[index192_] = (668) - ((self).f0)
                index193_ = default__.safeIndex(201, (d_1126_v16_).length(0))
                (d_1126_v16_)[index193_] = p0
                d_1147_v29_: D2
                d_1147_v29_ = D2_DC4((d_1122_v12_).set(default__.safeIndex(d_1145_cf30_, len(d_1122_v12_)), d_1110_v1_), p1)
                d_1148_v30_: D0
                d_1148_v30_ = D0_DC0((d_1147_v29_).cf5)
                index194_ = default__.safeIndex(201, (d_1126_v16_).length(0))
                rhs185_ = p0
                rhs186_ = d_1148_v30_
                lhs82_ = d_1126_v16_
                lhs83_ = default__.safeIndex(201, (d_1126_v16_).length(0))
                lhs82_[lhs83_] = rhs185_
                d_1148_v30_ = rhs186_
                rhs187_ = d_1122_v12_
                rhs188_ = p1
                d_1122_v12_ = rhs187_
                r0 = rhs188_
            elif True:
                d_1149___mcc_h4_ = source13_.cf26
                d_1150_cf26_ = d_1149___mcc_h4_
                d_1151_v31_: _dafny.Map
                d_1151_v31_ = _dafny.Map({(d_1109_v0_) > ((d_1123_v13_)[default__.safeIndex(len(d_1122_v12_), len(d_1123_v13_))]): (p1 if not(p1) else p0)})
                d_1152_v32_: D2
                d_1152_v32_ = D2_DC4(d_1122_v12_, p0)
                d_1153_v33_: _dafny.Map
                d_1153_v33_ = _dafny.Map({(d_1152_v32_).cf5: d_1122_v12_})
                d_1151_v31_ = (d_1151_v31_).set(not((d_1127_v17_).cf29), (self).fm6(default__.fm2(p1, (d_1119_v10_)[default__.safeIndex(990, (d_1119_v10_).length(0))], p1, False, globalState), ((d_1153_v33_)[p0] if (p0) in (d_1153_v33_) else d_1122_v12_), d_1122_v12_, d_1109_v0_, globalState))
                d_1154_v34_: _dafny.Array
                nw207_ = _dafny.Array(_dafny.MultiSet({}), 9)
                d_1154_v34_ = nw207_
                d_1155_v35_: _dafny.MultiSet
                d_1155_v35_ = _dafny.MultiSet([d_1110_v1_])
                index195_ = default__.safeIndex(388, (d_1154_v34_).length(0))
                (d_1154_v34_)[index195_] = d_1155_v35_
                d_1156_v36_: T0
                nw208_ = C3()
                nw208_.ctor__((self).f0)
                d_1156_v36_ = nw208_
                d_1157_v37_: D10
                d_1157_v37_ = D10_DC21(d_1156_v36_)
                d_1158_v38_: _dafny.Map
                d_1158_v38_ = _dafny.Map({(d_1155_v35_).cardinality: d_1157_v37_})
                index196_ = default__.safeIndex(990, (d_1119_v10_).length(0))
                index197_ = default__.safeIndex(388, (d_1154_v34_).length(0))
                rhs189_ = not (p0) or (not((len(d_1158_v38_)) < (len(d_1122_v12_))))
                rhs190_ = p0
                rhs191_ = (d_1156_v36_).f0
                rhs192_ = d_1155_v35_
                lhs84_ = d_1119_v10_
                lhs85_ = default__.safeIndex(990, (d_1119_v10_).length(0))
                lhs86_ = d_1154_v34_
                lhs87_ = default__.safeIndex(388, (d_1154_v34_).length(0))
                r0 = rhs189_
                r0 = rhs190_
                lhs84_[lhs85_] = rhs191_
                lhs86_[lhs87_] = rhs192_
                d_1159_v39_: C1
                nw209_ = C1()
                nw209_.ctor__((d_1119_v10_)[default__.safeIndex(990, (d_1119_v10_).length(0))])
                d_1159_v39_ = nw209_
                d_1160_v40_: _dafny.Map
                d_1160_v40_ = _dafny.Map({d_1109_v0_: d_1159_v39_})
                d_1161_v41_: _dafny.Map
                d_1161_v41_ = _dafny.Map({(d_1123_v13_) < (_dafny.SeqWithoutIsStrInference([(d_1119_v10_)[default__.safeIndex(990, (d_1119_v10_).length(0))] for d_1162_i1_ in range(default__.abs(88))])): len(d_1160_v40_)})
                d_1161_v41_ = (_dafny.Map({p0: (d_1156_v36_).f0})) | (d_1161_v41_)
                d_1163_v42_: _dafny.Map
                d_1163_v42_ = _dafny.Map({d_1110_v1_: d_1123_v13_})
                d_1123_v13_ = ((d_1163_v42_)[d_1110_v1_] if (d_1110_v1_) in (d_1163_v42_) else d_1123_v13_)
        d_1164_v43_: _dafny.Set
        d_1164_v43_ = _dafny.Set({p1, p1, True})
        hi9_ = default__.fm3(d_1164_v43_, d_1110_v1_, d_1109_v0_, (self).f0, globalState)
        for d_1165_i2_ in range((self).f0, hi9_):
            d_1166_v44_: _dafny.Map
            d_1166_v44_ = _dafny.Map({p1: d_1109_v0_})
            d_1166_v44_ = (d_1166_v44_).set(p1, d_1165_i2_)
            d_1167_v45_: _dafny.Array
            nw210_ = _dafny.Array(int(0), 8)
            d_1167_v45_ = nw210_
            d_1168_v46_: _dafny.Seq
            d_1168_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
            index198_ = default__.safeIndex(188, (d_1167_v45_).length(0))
            (d_1167_v45_)[index198_] = (d_1109_v0_) + (len(d_1168_v46_))
            d_1169_v47_: _dafny.Array
            def lambda53_(d_1170_p0_):
                def lambda54_(d_1171_i3_):
                    return d_1170_p0_

                return lambda54_

            init27_ = lambda53_(p0)
            nw211_ = _dafny.Array(None, 7)
            for i0_27_ in range(nw211_.length(0)):
                nw211_[i0_27_] = init27_(i0_27_)
            d_1169_v47_ = nw211_
            d_1172_v48_: D11
            d_1172_v48_ = D11_DC24(d_1169_v47_, p1, not(p0))
            d_1173_v49_: D5
            d_1173_v49_ = D5_DC13(d_1165_i2_, _dafny.Set({True, True, p0, p1, p1}), (d_1168_v46_)[default__.safeIndex(d_1165_i2_, len(d_1168_v46_))], d_1165_i2_, p0)
            index199_ = default__.safeIndex(188, (d_1167_v45_).length(0))
            rhs193_ = ((self).f0) * (len(d_1164_v43_))
            rhs194_ = D11_DC24(d_1169_v47_, (d_1173_v49_).cf21, p1)
            rhs195_ = (p1) or (p0)
            lhs88_ = d_1167_v45_
            lhs89_ = default__.safeIndex(188, (d_1167_v45_).length(0))
            lhs88_[lhs89_] = rhs193_
            d_1172_v48_ = rhs194_
            r0 = rhs195_
            d_1174_v50_: _dafny.Map
            d_1174_v50_ = _dafny.Map({(d_1167_v45_)[default__.safeIndex(188, (d_1167_v45_).length(0))]: not(p1)})
            rhs196_ = d_1165_i2_
            rhs197_ = not(((self).f0) < ((d_1167_v45_)[default__.safeIndex(188, (d_1167_v45_).length(0))]))
            rhs198_ = (d_1174_v50_) | (d_1174_v50_)
            rhs199_ = (p1) == ((d_1164_v43_).ispropersubset(d_1164_v43_))
            d_1109_v0_ = rhs196_
            r0 = rhs197_
            d_1174_v50_ = rhs198_
            r0 = rhs199_
            d_1175_v51_: _dafny.Map
            d_1175_v51_ = _dafny.Map({(d_1167_v45_)[default__.safeIndex(188, (d_1167_v45_).length(0))]: d_1165_i2_})
            d_1176_v52_: _dafny.Map
            d_1176_v52_ = _dafny.Map({d_1175_v51_: d_1167_v45_})
            d_1176_v52_ = _dafny.Map({d_1175_v51_: d_1167_v45_})
        d_1177_v53_: D7
        d_1177_v53_ = D7_DC17(p1, p1, (p1) == (p1))
        source14_ = d_1177_v53_
        if source14_.is_DC17:
            d_1178___mcc_h5_ = source14_.cf27
            d_1179___mcc_h6_ = source14_.cf28
            d_1180___mcc_h7_ = source14_.cf29
            d_1181_cf29_ = d_1180___mcc_h7_
            d_1182_cf28_ = d_1179___mcc_h6_
            d_1183_cf27_ = d_1178___mcc_h5_
            d_1184_v54_: _dafny.Map
            d_1184_v54_ = _dafny.Map({d_1109_v0_: True})
            d_1184_v54_ = d_1184_v54_
            d_1109_v0_ = d_1109_v0_
            d_1185_v55_: _dafny.Array
            def lambda55_(d_1186_v0_):
                def lambda56_(d_1187_i4_):
                    return default__.safeDivisionInt(d_1187_i4_, d_1186_v0_)

                return lambda56_

            init28_ = lambda55_(d_1109_v0_)
            nw212_ = _dafny.Array(None, 22)
            for i0_28_ in range(nw212_.length(0)):
                nw212_[i0_28_] = init28_(i0_28_)
            d_1185_v55_ = nw212_
            d_1188_v56_: _dafny.Seq
            d_1188_v56_ = _dafny.SeqWithoutIsStrInference([not(p0)])
            index200_ = default__.safeIndex(980, (d_1185_v55_).length(0))
            (d_1185_v55_)[index200_] = len((d_1188_v56_) + (d_1188_v56_))
            index201_ = default__.safeIndex(980, (d_1185_v55_).length(0))
            (d_1185_v55_)[index201_] = d_1109_v0_
            d_1189_v57_: C1
            nw213_ = C1()
            nw213_.ctor__((self).f0)
            d_1189_v57_ = nw213_
        elif source14_.is_DC18:
            d_1190___mcc_h8_ = source14_.cf30
            d_1191_cf30_ = d_1190___mcc_h8_
            d_1192_v58_: _dafny.Array
            def lambda57_(d_1193_v0_):
                def lambda58_(d_1194_i5_):
                    return _dafny.SeqWithoutIsStrInference([D5_DC11(_dafny.Set({_dafny.Set({len(_dafny.Set({(self).f0, 810}))})})), D5_DC11(_dafny.Set({_dafny.Set({(self).f0, d_1193_v0_, (self).f0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqlvvtuaq"))), (_dafny.MultiSet([(self).f0])).cardinality})}))])

                return lambda58_

            init29_ = lambda57_(d_1109_v0_)
            nw214_ = _dafny.Array(None, 9)
            for i0_29_ in range(nw214_.length(0)):
                nw214_[i0_29_] = init29_(i0_29_)
            d_1192_v58_ = nw214_
            d_1195_v59_: _dafny.Array
            nw215_ = _dafny.Array(False, 21)
            d_1195_v59_ = nw215_
            d_1196_v60_: _dafny.MultiSet
            d_1196_v60_ = _dafny.MultiSet([-22, d_1109_v0_])
            d_1197_v61_: _dafny.Seq
            d_1197_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkw"))
            rhs200_ = d_1195_v59_
            rhs201_ = ((d_1196_v60_)[len(_dafny.Map({p1: p1}))] if (len(_dafny.Map({p1: p1}))) in (d_1196_v60_) else default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqka"))), len(_dafny.Set({p1, p1, p0}))))
            rhs202_ = (default__.safeDivisionInt(d_1109_v0_, (self).f0)) - (default__.safeDivisionInt(d_1109_v0_, len(d_1197_v61_)))
            rhs203_ = d_1192_v58_
            r1 = rhs200_
            d_1109_v0_ = rhs201_
            d_1109_v0_ = rhs202_
            d_1192_v58_ = rhs203_
            d_1198_v62_: _dafny.Array
            nw216_ = _dafny.Array(None, 18)
            nw216_[int(0)] = d_1195_v59_
            nw216_[int(1)] = (d_1195_v59_ if not((D0_DC1(p0)).cf1) else d_1195_v59_)
            nw216_[int(2)] = d_1195_v59_
            nw216_[int(3)] = d_1195_v59_
            nw216_[int(4)] = d_1195_v59_
            nw216_[int(5)] = d_1195_v59_
            nw216_[int(6)] = d_1195_v59_
            nw216_[int(7)] = d_1195_v59_
            nw216_[int(8)] = d_1195_v59_
            nw216_[int(9)] = d_1195_v59_
            nw216_[int(10)] = d_1195_v59_
            nw216_[int(11)] = d_1195_v59_
            nw216_[int(12)] = d_1195_v59_
            nw216_[int(13)] = d_1195_v59_
            nw216_[int(14)] = d_1195_v59_
            nw216_[int(15)] = (d_1195_v59_ if p0 else d_1195_v59_)
            nw216_[int(16)] = d_1195_v59_
            nw216_[int(17)] = d_1195_v59_
            d_1198_v62_ = nw216_
            index202_ = default__.safeIndex(407, (d_1198_v62_).length(0))
            (d_1198_v62_)[index202_] = d_1195_v59_
            d_1199_v63_: _dafny.Array
            nw217_ = _dafny.Array(int(0), 24)
            d_1199_v63_ = nw217_
            d_1200_v64_: _dafny.MultiSet
            d_1200_v64_ = _dafny.MultiSet([p1, (self).fm6((self).fm5(globalState), d_1197_v61_, d_1197_v61_, d_1191_cf30_, globalState), p0])
            index203_ = default__.safeIndex(407, (d_1198_v62_).length(0))
            rhs204_ = (d_1197_v61_) + (d_1197_v61_)
            rhs205_ = d_1199_v63_
            rhs206_ = p1
            rhs207_ = d_1195_v59_
            rhs208_ = ((_dafny.MultiSet([not(p0)])) | (d_1200_v64_)).ispropersubset(((d_1200_v64_).set(p1, default__.abs(d_1191_cf30_))) - (d_1200_v64_))
            lhs90_ = d_1198_v62_
            lhs91_ = default__.safeIndex(407, (d_1198_v62_).length(0))
            d_1197_v61_ = rhs204_
            r2 = rhs205_
            r0 = rhs206_
            lhs90_[lhs91_] = rhs207_
            r0 = rhs208_
            d_1201_v65_: C0
            nw218_ = C0()
            nw218_.ctor__(d_1110_v1_)
            d_1201_v65_ = nw218_
            d_1201_v65_ = d_1201_v65_
            index204_ = default__.safeIndex(400, (d_1199_v63_).length(0))
            (d_1199_v63_)[index204_] = d_1191_cf30_
            index205_ = default__.safeIndex(400, (d_1199_v63_).length(0))
            (d_1199_v63_)[index205_] = (d_1109_v0_) + ((self).f0)
        elif True:
            d_1202___mcc_h9_ = source14_.cf26
            d_1203_cf26_ = d_1202___mcc_h9_
            d_1204_v66_: _dafny.Seq
            d_1204_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqjpl"))
            d_1205_v67_: _dafny.Set
            d_1205_v67_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vk")), d_1204_v66_})
            d_1206_v68_: _dafny.Map
            d_1206_v68_ = _dafny.Map({(self).f0: p1})
            d_1207_v69_: _dafny.Map
            d_1207_v69_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jut")): d_1206_v68_})
            def iife71_():
                coll25_ = _dafny.Set()
                compr_25_: _dafny.Seq
                for compr_25_ in (d_1207_v69_).keys.Elements:
                    d_1208_v70_: _dafny.Seq = compr_25_
                    if (d_1208_v70_) in (d_1207_v69_):
                        coll25_ = coll25_.union(_dafny.Set([d_1208_v70_]))
                return _dafny.Set(coll25_)
            r0 = (iife71_()
            ).ispropersubset(d_1205_v67_)
            d_1209_v71_: _dafny.Map
            d_1209_v71_ = _dafny.Map({p0: p1})
            d_1210_v72_: _dafny.Set
            d_1210_v72_ = _dafny.Set({d_1109_v0_, len(d_1209_v71_)})
            r0 = (_dafny.Set({(self).f0, len(default__.fm0(globalState)), d_1109_v0_, d_1109_v0_})) != (d_1210_v72_)
            d_1110_v1_ = d_1110_v1_
            d_1211_v73_: C0
            nw219_ = C0()
            nw219_.ctor__(d_1110_v1_)
            d_1211_v73_ = nw219_
        r0 = ((d_1164_v43_) | (d_1164_v43_)).issubset((d_1164_v43_ if p1 else d_1164_v43_))
        d_1212_v74_: _dafny.Array
        nw220_ = _dafny.Array(False, 7)
        d_1212_v74_ = nw220_
        r1 = d_1212_v74_
        d_1213_v75_: _dafny.Seq
        d_1213_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shkq"))
        d_1214_v76_: _dafny.MultiSet
        d_1214_v76_ = _dafny.MultiSet([d_1109_v0_])
        d_1215_v77_: _dafny.Seq
        d_1215_v77_ = _dafny.SeqWithoutIsStrInference([(self).f0, d_1109_v0_])
        d_1216_v78_: _dafny.Array
        nw221_ = _dafny.Array(None, 13)
        d_1216_v78_ = nw221_
        pat_let_tv31_ = d_1109_v0_
        d_1217_v79_: _dafny.Array
        nw222_ = _dafny.Array(None, 7)
        nw222_[int(0)] = (len(d_1213_v75_)) * (d_1109_v0_)
        nw222_[int(1)] = d_1109_v0_
        nw222_[int(2)] = d_1109_v0_
        nw222_[int(3)] = ((d_1214_v76_) - (_dafny.MultiSet(d_1215_v77_))).cardinality
        nw222_[int(4)] = d_1109_v0_
        def iife72_(_pat_let23_0):
            def iife73_(d_1218_dt__update__tmp_h0_):
                def iife74_(_pat_let24_0):
                    def iife75_(d_1219_dt__update_hcf10_h0_):
                        return D3_DC9(d_1219_dt__update_hcf10_h0_, (d_1218_dt__update__tmp_h0_).cf11, (d_1218_dt__update__tmp_h0_).cf12)
                    return iife75_(_pat_let24_0)
                return iife74_(pat_let_tv31_)
            return iife73_(_pat_let23_0)
        nw222_[int(5)] = (0) - ((iife72_(D3_DC9((self).f0, d_1216_v78_, p1))).cf10)
        nw222_[int(6)] = d_1109_v0_
        d_1217_v79_ = nw222_
        r2 = d_1217_v79_
        d_1220_v80_: _dafny.MultiSet
        d_1220_v80_ = _dafny.MultiSet([False])
        r3 = ((d_1220_v80_).set(p0, default__.abs(d_1109_v0_))).intersection(d_1220_v80_)
        return r0, r1, r2, r3


class C5(T1, T0):
    def  __init__(self):
        self._f0: int = int(0)
        self._f1: str = _dafny.CodePoint('D')
        self._f2: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f0(self):
        return self._f0
    def ctor__(self, f1, f2, f0):
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f0 = f0

    def fm5(self, globalState):
        return False

    def fm6(self, p0, p1, p2, p3, globalState):
        return not (not(((self).f0) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilqiu"))) for d_1221_i0_ in range(default__.abs(573))]))))) or ((True) == (True))

    def fm7(self, p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('n')

    def m2(self, globalState):
        d_1222_v0_: C0
        nw223_ = C0()
        nw223_.ctor__((self).f1)
        d_1222_v0_ = nw223_
        d_1223_v1_: bool
        d_1223_v1_ = True
        d_1224_v2_: _dafny.Seq
        d_1224_v2_ = _dafny.SeqWithoutIsStrInference([True, d_1223_v1_])
        d_1225_v3_: _dafny.Map
        d_1225_v3_ = _dafny.Map({d_1223_v1_: d_1224_v2_})
        d_1226_v4_: _dafny.Array
        nw224_ = _dafny.Array(None, 9)
        nw224_[int(0)] = d_1223_v1_
        nw224_[int(1)] = (d_1223_v1_) == (d_1223_v1_)
        nw224_[int(2)] = d_1223_v1_
        nw224_[int(3)] = d_1223_v1_
        nw224_[int(4)] = d_1223_v1_
        nw224_[int(5)] = d_1223_v1_
        nw224_[int(6)] = (d_1224_v2_) <= (((d_1225_v3_)[d_1223_v1_] if (d_1223_v1_) in (d_1225_v3_) else d_1224_v2_))
        nw224_[int(7)] = not (d_1223_v1_) or (d_1223_v1_)
        nw224_[int(8)] = d_1223_v1_
        d_1226_v4_ = nw224_
        index206_ = default__.safeIndex(8, (d_1226_v4_).length(0))
        (d_1226_v4_)[index206_] = d_1223_v1_
        d_1227_v5_: _dafny.Map
        d_1227_v5_ = _dafny.Map({((self).fm5(globalState)) or (d_1223_v1_): _dafny.MultiSet([d_1223_v1_])})
        d_1228_v6_: _dafny.Seq
        d_1228_v6_ = _dafny.SeqWithoutIsStrInference([d_1224_v2_, _dafny.SeqWithoutIsStrInference([d_1223_v1_])])
        d_1229_v7_: _dafny.Map
        d_1229_v7_ = _dafny.Map({True: (d_1222_v0_).f3})
        d_1230_v8_: _dafny.Seq
        d_1230_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sargafyl"))
        d_1231_v9_: _dafny.Seq
        d_1231_v9_ = _dafny.SeqWithoutIsStrInference([((d_1229_v7_)[d_1223_v1_] if (d_1223_v1_) in (d_1229_v7_) else (self).f1), (d_1222_v0_).f3, (d_1230_v8_)[default__.safeIndex((self).f2, len(d_1230_v8_))], (self).f1])
        d_1232_v10_: _dafny.Seq
        d_1232_v10_ = _dafny.SeqWithoutIsStrInference([(d_1231_v9_)[default__.safeIndex((self).f0, len(d_1231_v9_))]])
        d_1233_v11_: _dafny.MultiSet
        d_1233_v11_ = _dafny.MultiSet([d_1223_v1_])
        index207_ = default__.safeIndex(8, (d_1226_v4_).length(0))
        rhs209_ = ((self).f0) < (len((d_1228_v6_).set(default__.safeIndex((self).f2, len(d_1228_v6_)), d_1224_v2_)))
        rhs210_ = ((self).f1) in ((d_1231_v9_).set(default__.safeIndex((0) - ((self).f0), len(d_1231_v9_)), (self).f1))
        rhs211_ = ((d_1227_v5_).set(d_1223_v1_, _dafny.MultiSet([d_1223_v1_, not(d_1223_v1_), d_1223_v1_]))) | (_dafny.Map({d_1223_v1_: d_1233_v11_}))
        lhs92_ = d_1226_v4_
        lhs93_ = default__.safeIndex(8, (d_1226_v4_).length(0))
        lhs92_[lhs93_] = rhs209_
        d_1223_v1_ = rhs210_
        d_1227_v5_ = rhs211_
        d_1234_v12_: int
        d_1234_v12_ = 796
        d_1234_v12_ = 234
        d_1235_v13_: _dafny.Seq
        d_1235_v13_ = _dafny.SeqWithoutIsStrInference([d_1222_v0_, d_1222_v0_, d_1222_v0_])
        d_1236_v14_: _dafny.Map
        d_1236_v14_ = _dafny.Map({d_1233_v11_: d_1223_v1_})
        d_1237_v15_: _dafny.Seq
        d_1237_v15_ = _dafny.SeqWithoutIsStrInference([182, (0) - ((self).f0)])
        if not((_dafny.MultiSet(d_1237_v15_)).issubset(_dafny.MultiSet([len((_dafny.SeqWithoutIsStrInference([len(d_1235_v13_), len(d_1224_v2_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdcu"))), len(d_1236_v14_)])).set(default__.safeIndex(d_1234_v12_, len(_dafny.SeqWithoutIsStrInference([len(d_1235_v13_), len(d_1224_v2_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdcu"))), len(d_1236_v14_)]))), (self).f2))]))):
            d_1238_v16_: _dafny.Set
            d_1238_v16_ = _dafny.Set({d_1234_v12_, (self).f2, (self).f0})
            index208_ = default__.safeIndex(8, (d_1226_v4_).length(0))
            rhs212_ = (745) != ((len(d_1237_v15_)) - (d_1234_v12_))
            rhs213_ = d_1234_v12_
            rhs214_ = (d_1223_v1_) and (False)
            rhs215_ = d_1238_v16_
            lhs94_ = d_1226_v4_
            lhs95_ = default__.safeIndex(8, (d_1226_v4_).length(0))
            lhs94_[lhs95_] = rhs212_
            d_1234_v12_ = rhs213_
            d_1223_v1_ = rhs214_
            d_1238_v16_ = rhs215_
            d_1239_v17_: _dafny.Array
            def lambda59_(d_1240_v2_):
                def lambda60_(d_1241_i0_):
                    return d_1240_v2_

                return lambda60_

            init30_ = lambda59_(d_1224_v2_)
            nw225_ = _dafny.Array(None, 12)
            for i0_30_ in range(nw225_.length(0)):
                nw225_[i0_30_] = init30_(i0_30_)
            d_1239_v17_ = nw225_
            index209_ = default__.safeIndex(793, (d_1239_v17_).length(0))
            (d_1239_v17_)[index209_] = d_1224_v2_
            d_1242_v18_: _dafny.Seq
            d_1242_v18_ = _dafny.SeqWithoutIsStrInference([d_1226_v4_])
            d_1243_v19_: D3
            d_1243_v19_ = D3_DC7(d_1237_v15_)
            index210_ = default__.safeIndex(793, (d_1239_v17_).length(0))
            rhs216_ = (d_1242_v18_)[default__.safeIndex(default__.safeDivisionInt(d_1234_v12_, len((d_1243_v19_).cf7)), len(d_1242_v18_))]
            rhs217_ = not(d_1223_v1_)
            rhs218_ = (d_1224_v2_ if d_1223_v1_ else d_1224_v2_)
            rhs219_ = (d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))]
            lhs96_ = d_1239_v17_
            lhs97_ = default__.safeIndex(793, (d_1239_v17_).length(0))
            d_1226_v4_ = rhs216_
            d_1223_v1_ = rhs217_
            lhs96_[lhs97_] = rhs218_
            d_1223_v1_ = rhs219_
            d_1244_v20_: _dafny.Map
            d_1244_v20_ = _dafny.Map({(self).f2: d_1226_v4_})
            if ((0) - (d_1234_v12_)) not in (d_1244_v20_):
                index211_ = default__.safeIndex(8, (d_1226_v4_).length(0))
                (d_1226_v4_)[index211_] = (not(not ((d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))]) or ((d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))])) if (d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))] else (d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))])
                d_1245_v21_: _dafny.MultiSet
                d_1245_v21_ = _dafny.MultiSet([default__.safeDivisionInt(d_1234_v12_, 573), ((self).f2) + (d_1234_v12_), (len(d_1230_v8_)) + ((self).f0)])
                d_1245_v21_ = d_1245_v21_
                d_1246_v22_: C0
                nw226_ = C0()
                nw226_.ctor__((d_1222_v0_).f3)
                d_1246_v22_ = nw226_
                d_1247_v23_: _dafny.Array
                def lambda61_(d_1248_v9_):
                    def lambda62_(d_1249_i1_):
                        return d_1248_v9_

                    return lambda62_

                init31_ = lambda61_(d_1231_v9_)
                nw227_ = _dafny.Array(None, 23)
                for i0_31_ in range(nw227_.length(0)):
                    nw227_[i0_31_] = init31_(i0_31_)
                d_1247_v23_ = nw227_
                index212_ = default__.safeIndex(353, (d_1247_v23_).length(0))
                (d_1247_v23_)[index212_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eblbebi"))).set(default__.safeIndex(73, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eblbebi")))), (d_1222_v0_).f3)) + (d_1231_v9_)
                index213_ = default__.safeIndex(353, (d_1247_v23_).length(0))
                (d_1247_v23_)[index213_] = (d_1231_v9_) + (d_1231_v9_)
                index214_ = default__.safeIndex(8, (d_1226_v4_).length(0))
                (d_1226_v4_)[index214_] = (d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))]
            elif True:
                index215_ = default__.safeIndex(8, (d_1226_v4_).length(0))
                (d_1226_v4_)[index215_] = d_1223_v1_
                index216_ = default__.safeIndex(8, (d_1226_v4_).length(0))
                (d_1226_v4_)[index216_] = (self).fm5(globalState)
                d_1234_v12_ = d_1234_v12_
                d_1250_v24_: C0
                nw228_ = C0()
                nw228_.ctor__((d_1222_v0_).f3)
                d_1250_v24_ = nw228_
                d_1251_v25_: _dafny.MultiSet
                d_1251_v25_ = _dafny.MultiSet([(self).f2])
                index217_ = default__.safeIndex(8, (d_1226_v4_).length(0))
                (d_1226_v4_)[index217_] = (d_1251_v25_) != ((d_1251_v25_).set((self).f0, default__.abs(46)))
            d_1252_v26_: _dafny.Array
            nw229_ = _dafny.Array(None, 12)
            d_1252_v26_ = nw229_
            index218_ = default__.safeIndex(820, (d_1252_v26_).length(0))
            (d_1252_v26_)[index218_] = d_1222_v0_
            d_1253_v27_: str
            d_1253_v27_ = (self).f1
            index219_ = default__.safeIndex(820, (d_1252_v26_).length(0))
            nw230_ = C0()
            nw230_.ctor__((d_1253_v27_))
            (d_1252_v26_)[index219_] = nw230_
            d_1237_v15_ = d_1237_v15_
        elif True:
            index220_ = default__.safeIndex(8, (d_1226_v4_).length(0))
            index221_ = default__.safeIndex(8, (d_1226_v4_).length(0))
            rhs220_ = (d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))]
            rhs221_ = not (d_1223_v1_) or ((d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))])
            lhs98_ = d_1226_v4_
            lhs99_ = default__.safeIndex(8, (d_1226_v4_).length(0))
            lhs100_ = d_1226_v4_
            lhs101_ = default__.safeIndex(8, (d_1226_v4_).length(0))
            lhs98_[lhs99_] = rhs220_
            lhs100_[lhs101_] = rhs221_
            d_1254_v28_: D2
            d_1254_v28_ = D2_DC5()
            d_1255_v29_: D3
            d_1255_v29_ = D3_DC7(_dafny.SeqWithoutIsStrInference([len(d_1237_v15_) for d_1256_i2_ in range(default__.abs(925))]))
            d_1257_v30_: _dafny.Map
            d_1257_v30_ = _dafny.Map({d_1223_v1_: d_1255_v29_})
            d_1258_v31_: _dafny.Map
            d_1258_v31_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1234_v12_, -499, (self).f0]): (d_1222_v0_).f3})
            d_1254_v28_ = default__.fm8((0) - (d_1234_v12_), d_1257_v30_, d_1258_v31_, (d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))], globalState)
            d_1259_v32_: D0
            d_1259_v32_ = D0_DC0(d_1223_v1_)
            d_1223_v1_ = (d_1259_v32_).cf0
            index222_ = default__.safeIndex(8, (d_1226_v4_).length(0))
            (d_1226_v4_)[index222_] = not(not((d_1234_v12_) >= ((d_1234_v12_) - (len(_dafny.SeqWithoutIsStrInference([d_1234_v12_]))))))
            source15_ = d_1255_v29_
            if source15_.is_DC8:
                d_1260___mcc_h0_ = source15_.cf8
                d_1261___mcc_h1_ = source15_.cf9
                d_1262_cf9_ = d_1261___mcc_h1_
                d_1263_cf8_ = d_1260___mcc_h0_
                index223_ = default__.safeIndex(8, (d_1226_v4_).length(0))
                (d_1226_v4_)[index223_] = d_1263_cf8_
                d_1264_v33_: D2
                d_1264_v33_ = D2_DC5()
                d_1265_v34_: D2
                d_1265_v34_ = D2_DC6(d_1264_v33_)
                d_1265_v34_ = default__.fm9(globalState)
                d_1266_v35_: D3
                d_1266_v35_ = D3_DC8(d_1223_v1_, (self).f0)
                d_1234_v12_ = default__.safeDivisionInt((d_1266_v35_).cf9, d_1262_cf9_)
                d_1267_v36_: _dafny.Map
                d_1267_v36_ = _dafny.Map({not(d_1263_cf8_): d_1263_cf8_})
                d_1267_v36_ = (d_1267_v36_).set(d_1263_cf8_, d_1263_cf8_)
            elif source15_.is_DC9:
                d_1268___mcc_h2_ = source15_.cf10
                d_1269___mcc_h3_ = source15_.cf11
                d_1270___mcc_h4_ = source15_.cf12
                d_1271_cf12_ = d_1270___mcc_h4_
                d_1272_cf11_ = d_1269___mcc_h3_
                d_1273_cf10_ = d_1268___mcc_h2_
                d_1274_v37_: C0
                nw231_ = C0()
                nw231_.ctor__((d_1222_v0_).f3)
                d_1274_v37_ = nw231_
                d_1275_v38_: _dafny.Seq
                d_1275_v38_ = _dafny.SeqWithoutIsStrInference([d_1231_v9_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrbh")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xh"))) + (d_1230_v8_)])
                d_1276_v39_: _dafny.Map
                d_1276_v39_ = _dafny.Map({d_1223_v1_: _dafny.SeqWithoutIsStrInference([d_1230_v8_ for d_1277_i4_ in range(default__.abs(-525))])})
                d_1278_v40_: _dafny.Seq
                d_1278_v40_ = _dafny.SeqWithoutIsStrInference([((d_1275_v38_).set(default__.safeIndex((self).f0, len(d_1275_v38_)), d_1231_v9_)) + (d_1275_v38_), _dafny.SeqWithoutIsStrInference([d_1232_v10_ for d_1279_i3_ in range(default__.abs(-71))]), ((d_1276_v39_)[not((d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))])] if (not((d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))])) in (d_1276_v39_) else d_1275_v38_)])
                rhs222_ = (d_1274_v37_ if (d_1224_v2_)[default__.safeIndex((0) - (d_1273_cf10_), len(d_1224_v2_))] else d_1274_v37_)
                rhs223_ = (d_1278_v40_)[default__.safeIndex(-693, len(d_1278_v40_))]
                d_1274_v37_ = rhs222_
                d_1275_v38_ = rhs223_
                d_1280_v41_: _dafny.MultiSet
                d_1280_v41_ = _dafny.MultiSet([_dafny.CodePoint('p'), (self).f1, (d_1274_v37_).f3, (d_1222_v0_).f3, (d_1222_v0_).f3])
                d_1281_v42_: _dafny.Set
                d_1281_v42_ = _dafny.Set({d_1271_cf12_, d_1271_cf12_, (d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))], (d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))]})
                d_1282_v43_: _dafny.MultiSet
                d_1282_v43_ = _dafny.MultiSet([(self).f2, (self).f2, (self).f0, d_1234_v12_])
                d_1271_cf12_ = ((default__.fm3(default__.fm10(d_1280_v41_, d_1271_cf12_, globalState), (d_1222_v0_).f3, (self).f2, default__.fm3(d_1281_v42_, (d_1222_v0_).f3, len(d_1231_v9_), (self).f2, globalState), globalState)) - (d_1273_cf10_)) <= (((d_1282_v43_) - (_dafny.MultiSet([(self).f2, d_1234_v12_, (self).f2]))).cardinality)
                d_1283_v44_: _dafny.Map
                d_1283_v44_ = _dafny.Map({d_1271_cf12_: ((d_1274_v37_).f3) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuhjyw")))})
                d_1283_v44_ = (d_1283_v44_).set(not(d_1223_v1_), d_1223_v1_)
                d_1234_v12_ = ((0) - (d_1234_v12_)) + (d_1234_v12_)
            elif True:
                d_1284___mcc_h5_ = source15_.cf7
                d_1285_cf7_ = d_1284___mcc_h5_
                d_1232_v10_ = ((d_1231_v9_) + (d_1232_v10_)) + ((d_1231_v9_) + (d_1232_v10_))
                d_1286_v45_: D0
                d_1286_v45_ = D0_DC1(True)
                d_1286_v45_ = D0_DC1((d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))])
                d_1287_v46_: _dafny.Map
                d_1287_v46_ = _dafny.Map({(d_1223_v1_ if (d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))] else d_1223_v1_): False})
                d_1287_v46_ = (d_1287_v46_).set(not(default__.fm2((d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))], d_1234_v12_, (d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))], (d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))], globalState)), not(not(((d_1287_v46_)[d_1223_v1_] if (d_1223_v1_) in (d_1287_v46_) else (self).fm5(globalState)))))
                d_1288_v47_: C0
                nw232_ = C0()
                nw232_.ctor__((d_1222_v0_).f3)
                d_1288_v47_ = nw232_
        d_1289_v48_: C0
        nw233_ = C0()
        nw233_.ctor__((d_1222_v0_).f3)
        d_1289_v48_ = nw233_
        source16_ = D0_DC0(d_1223_v1_)
        if source16_.is_DC1:
            d_1290___mcc_h6_ = source16_.cf1
            d_1291_cf1_ = d_1290___mcc_h6_
            d_1234_v12_ = (self).f0
            d_1234_v12_ = ((701 if not((d_1226_v4_)[default__.safeIndex(8, (d_1226_v4_).length(0))]) else (self).f2)) - (len(d_1237_v15_))
            d_1292_v50_: _dafny.Set
            d_1292_v50_ = _dafny.Set({d_1291_cf1_, d_1223_v1_})
            d_1293_v51_: _dafny.Seq
            d_1293_v51_ = _dafny.SeqWithoutIsStrInference([d_1292_v50_])
            d_1294_v52_: _dafny.Map
            d_1294_v52_ = _dafny.Map({_dafny.CodePoint('s'): (d_1293_v51_)[default__.safeIndex((self).f0, len(d_1293_v51_))]})
            d_1295_v53_: _dafny.Map
            d_1295_v53_ = _dafny.Map({default__.fm3(((d_1294_v52_)[_dafny.CodePoint('o')] if (_dafny.CodePoint('o')) in (d_1294_v52_) else d_1292_v50_), (d_1222_v0_).f3, (self).f0, d_1234_v12_, globalState): (self).f0})
            def iife76_():
                coll26_ = _dafny.Map()
                compr_26_: int
                for compr_26_ in (d_1295_v53_).keys.Elements:
                    d_1296_v49_: int = compr_26_
                    if (d_1296_v49_) in (d_1295_v53_):
                        coll26_[(d_1296_v49_) - (len(_dafny.Set({False, d_1223_v1_})))] = d_1234_v12_
                return _dafny.Map(coll26_)
            d_1234_v12_ = ((self).f2) - (len((iife76_()
            ) | (d_1295_v53_)))
            d_1223_v1_ = True
        elif True:
            d_1297___mcc_h7_ = source16_.cf0
            d_1298_cf0_ = d_1297___mcc_h7_
            d_1299_v54_: _dafny.Array
            nw234_ = _dafny.Array(None, 18)
            d_1299_v54_ = nw234_
            index224_ = default__.safeIndex(841, (d_1299_v54_).length(0))
            (d_1299_v54_)[index224_] = d_1222_v0_
            index225_ = default__.safeIndex(841, (d_1299_v54_).length(0))
            (d_1299_v54_)[index225_] = d_1222_v0_
            d_1300_v55_: D0
            d_1300_v55_ = D0_DC0(d_1298_cf0_)
            d_1301_v56_: _dafny.Map
            d_1301_v56_ = _dafny.Map({d_1300_v55_: d_1226_v4_})
            d_1301_v56_ = (d_1301_v56_).set(d_1300_v55_, d_1226_v4_)
            d_1302_v57_: _dafny.Array
            nw235_ = _dafny.Array(int(0), 9)
            d_1302_v57_ = nw235_
            d_1303_v58_: _dafny.Map
            d_1303_v58_ = _dafny.Map({(self).f2: d_1234_v12_})
            index226_ = default__.safeIndex(497, (d_1302_v57_).length(0))
            (d_1302_v57_)[index226_] = len((d_1303_v58_) | (d_1303_v58_))
            index227_ = default__.safeIndex(497, (d_1302_v57_).length(0))
            (d_1302_v57_)[index227_] = 929
            d_1223_v1_ = (self).fm5(globalState)

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        d_1304_v0_: bool
        d_1304_v0_ = False
        d_1305_v1_: _dafny.Set
        d_1305_v1_ = _dafny.Set({d_1304_v0_, d_1304_v0_, d_1304_v0_})
        hi10_ = default__.fm3(d_1305_v1_, (self).f1, (self).f2, 2, globalState)
        for d_1306_i0_ in range(p1, hi10_):
            if d_1304_v0_:
                d_1307_v2_: _dafny.Set
                d_1307_v2_ = _dafny.Set({(self).f1})
                r0 = (0) - (((self).f0 if (_dafny.Set({_dafny.CodePoint('g')})).ispropersubset(d_1307_v2_) else p1))
                d_1308_v3_: D3
                d_1308_v3_ = D3_DC8(True, (self).f2)
                d_1309_v4_: _dafny.Seq
                d_1309_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmdnh"))
                d_1310_v5_: _dafny.Map
                d_1310_v5_ = _dafny.Map({(0) - (default__.safeDivisionInt((d_1308_v3_).cf9, len(d_1309_v4_))): len((d_1305_v1_) - (d_1305_v1_))})
                d_1310_v5_ = (d_1310_v5_).set(d_1306_i0_, (self).f0)
                d_1311_v6_: C0
                nw236_ = C0()
                nw236_.ctor__((self).f1)
                d_1311_v6_ = nw236_
                d_1304_v0_ = d_1304_v0_
                r1 = default__.safeModuloInt(d_1306_i0_, (self).f0)
            elif True:
                d_1312_v7_: _dafny.Array
                nw237_ = _dafny.Array(None, 1)
                nw237_[int(0)] = (d_1306_i0_ if d_1304_v0_ else d_1306_i0_)
                d_1312_v7_ = nw237_
                d_1313_v8_: T0
                nw238_ = C3()
                nw238_.ctor__((0) - ((0) - ((self).f2)))
                d_1313_v8_ = nw238_
                d_1314_v9_: _dafny.Seq
                d_1314_v9_ = _dafny.SeqWithoutIsStrInference([d_1313_v8_, d_1313_v8_, d_1313_v8_, d_1313_v8_])
                index228_ = default__.safeIndex(563, (d_1312_v7_).length(0))
                (d_1312_v7_)[index228_] = len(d_1314_v9_)
                d_1315_v10_: _dafny.Seq
                d_1315_v10_ = _dafny.SeqWithoutIsStrInference([(self).f2, p1, (self).f2, (self).f2])
                index229_ = default__.safeIndex(563, (d_1312_v7_).length(0))
                (d_1312_v7_)[index229_] = (d_1315_v10_)[default__.safeIndex(((self).f0) - (266), len(d_1315_v10_))]
                d_1316_v11_: _dafny.Map
                d_1316_v11_ = _dafny.Map({d_1304_v0_: d_1304_v0_})
                d_1317_v12_: _dafny.Map
                d_1317_v12_ = _dafny.Map({d_1304_v0_: _dafny.CodePoint('p')})
                index230_ = default__.safeIndex(563, (d_1312_v7_).length(0))
                rhs224_ = (d_1304_v0_ if d_1304_v0_ else (d_1305_v1_).ispropersubset(d_1305_v1_))
                rhs225_ = False
                rhs226_ = (len((d_1316_v11_) | (default__.fm20(d_1304_v0_, d_1317_v12_, d_1304_v0_, (self).f1, globalState)))) + (((self).f0) * (d_1306_i0_))
                rhs227_ = -863
                rhs228_ = d_1312_v7_
                lhs102_ = d_1312_v7_
                lhs103_ = default__.safeIndex(563, (d_1312_v7_).length(0))
                d_1304_v0_ = rhs224_
                d_1304_v0_ = rhs225_
                r0 = rhs226_
                lhs102_[lhs103_] = rhs227_
                d_1312_v7_ = rhs228_
                d_1318_v13_: D0
                d_1318_v13_ = D0_DC1(d_1304_v0_)
                d_1319_v14_: D2
                d_1319_v14_ = D2_DC6(D2_DC3(_dafny.SeqWithoutIsStrInference([d_1318_v13_, d_1318_v13_])))
                d_1320_v15_: D2
                d_1320_v15_ = D2_DC6(d_1319_v14_)
                d_1321_v16_: _dafny.Map
                d_1321_v16_ = _dafny.Map({d_1320_v15_: p1})
                d_1304_v0_ = (_dafny.Map({d_1320_v15_: (d_1313_v8_).f0})) == (d_1321_v16_)
                d_1322_v17_: C2
                nw239_ = C2()
                nw239_.ctor__(409)
                d_1322_v17_ = nw239_
                d_1323_v18_: _dafny.Seq
                d_1323_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "paqg"))
                d_1324_v19_: _dafny.Array
                nw240_ = _dafny.Array(None, 22)
                nw240_[int(0)] = d_1323_v18_
                nw240_[int(1)] = d_1323_v18_
                nw240_[int(2)] = default__.fm18(p1, globalState)
                nw240_[int(3)] = d_1323_v18_
                nw240_[int(4)] = (d_1323_v18_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hibaahddh"))).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hibaahddh")))), (self).f1))
                nw240_[int(5)] = d_1323_v18_
                nw240_[int(6)] = d_1323_v18_
                nw240_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
                nw240_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvptickwg"))
                nw240_[int(9)] = d_1323_v18_
                nw240_[int(10)] = d_1323_v18_
                nw240_[int(11)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))) + (d_1323_v18_)
                nw240_[int(12)] = d_1323_v18_
                nw240_[int(13)] = d_1323_v18_
                nw240_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cidhpdr"))
                nw240_[int(15)] = _dafny.SeqWithoutIsStrInference([(d_1323_v18_)[default__.safeIndex((self).f2, len(d_1323_v18_))] for d_1325_i1_ in range(default__.abs(82))])
                nw240_[int(16)] = d_1323_v18_
                nw240_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abkkeaax"))
                nw240_[int(18)] = d_1323_v18_
                nw240_[int(19)] = d_1323_v18_
                nw240_[int(20)] = _dafny.SeqWithoutIsStrInference([(self).f1 for d_1326_i2_ in range(default__.abs(580))])
                nw240_[int(21)] = d_1323_v18_
                d_1324_v19_ = nw240_
                index231_ = default__.safeIndex(189, (d_1324_v19_).length(0))
                (d_1324_v19_)[index231_] = d_1323_v18_
                index232_ = default__.safeIndex(189, (d_1324_v19_).length(0))
                (d_1324_v19_)[index232_] = (d_1323_v18_) + (d_1323_v18_)
            d_1327_v20_: _dafny.Seq
            d_1327_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
            d_1328_v21_: _dafny.Map
            d_1328_v21_ = _dafny.Map({(self).f2: d_1327_v20_})
            d_1329_v22_: _dafny.MultiSet
            d_1329_v22_ = _dafny.MultiSet([True])
            d_1330_v23_: _dafny.MultiSet
            d_1330_v23_ = _dafny.MultiSet([749, p1])
            d_1331_v24_: _dafny.Seq
            d_1331_v24_ = _dafny.SeqWithoutIsStrInference([d_1304_v0_])
            d_1332_v25_: D10
            d_1332_v25_ = D10_DC22(default__.fm3(default__.fm10(_dafny.MultiSet([(self).f1]), d_1304_v0_, globalState), (self).f1, len(((d_1328_v21_)[(d_1329_v22_).cardinality] if ((d_1329_v22_).cardinality) in (d_1328_v21_) else default__.fm19(d_1330_v23_, len(d_1331_v24_), d_1306_i0_, p1, globalState))), d_1306_i0_, globalState), (d_1304_v0_) and (d_1304_v0_), d_1304_v0_)
            d_1332_v25_ = d_1332_v25_
            r2 = (0) - ((self).f2)
            d_1304_v0_ = (not(d_1304_v0_) if d_1304_v0_ else d_1304_v0_)
        hi11_ = ((self).f0) - (p1)
        for d_1333_i3_ in range((self).f0, hi11_):
            d_1334_v26_: _dafny.Array
            nw241_ = _dafny.Array(None, 13)
            d_1334_v26_ = nw241_
            d_1335_v27_: _dafny.MultiSet
            d_1335_v27_ = _dafny.MultiSet([d_1304_v0_])
            d_1336_v28_: _dafny.Map
            d_1336_v28_ = _dafny.Map({d_1335_v27_: not(d_1304_v0_)})
            d_1337_v29_: T0
            nw242_ = C1()
            nw242_.ctor__(len(d_1336_v28_))
            d_1337_v29_ = nw242_
            index233_ = default__.safeIndex(242, (d_1334_v26_).length(0))
            (d_1334_v26_)[index233_] = d_1337_v29_
            index234_ = default__.safeIndex(242, (d_1334_v26_).length(0))
            (d_1334_v26_)[index234_] = d_1337_v29_
            d_1338_v30_: _dafny.Seq
            d_1338_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yyrnjblm"))
            d_1338_v30_ = d_1338_v30_
            r1 = (self).f0
            d_1339_v31_: _dafny.Seq
            d_1339_v31_ = _dafny.SeqWithoutIsStrInference([d_1304_v0_, True])
            d_1340_v32_: _dafny.Map
            d_1340_v32_ = _dafny.Map({(self).f2: d_1333_i3_})
            d_1341_v33_: _dafny.Seq
            d_1341_v33_ = _dafny.SeqWithoutIsStrInference([p1, (self).f0])
            d_1342_v34_: _dafny.Seq
            d_1342_v34_ = _dafny.SeqWithoutIsStrInference([d_1341_v33_])
            d_1343_v35_: _dafny.Map
            d_1343_v35_ = _dafny.Map({d_1304_v0_: d_1304_v0_})
            d_1344_v36_: _dafny.Map
            d_1344_v36_ = _dafny.Map({-707: True})
            d_1345_v37_: _dafny.Array
            nw243_ = _dafny.Array(None, 20)
            nw243_[int(0)] = (d_1335_v27_) - (d_1335_v27_)
            nw243_[int(1)] = (d_1335_v27_) | (d_1335_v27_)
            nw243_[int(2)] = (_dafny.MultiSet([d_1304_v0_, d_1304_v0_])) | (d_1335_v27_)
            nw243_[int(3)] = (d_1335_v27_).set(d_1304_v0_, default__.abs(d_1333_i3_))
            nw243_[int(4)] = (_dafny.MultiSet(d_1339_v31_)) - (d_1335_v27_)
            nw243_[int(5)] = d_1335_v27_
            nw243_[int(6)] = d_1335_v27_
            nw243_[int(7)] = (_dafny.MultiSet(d_1339_v31_)).set(d_1304_v0_, default__.abs(((d_1340_v32_)[d_1333_i3_] if (d_1333_i3_) in (d_1340_v32_) else len(d_1342_v34_))))
            nw243_[int(8)] = _dafny.MultiSet([d_1304_v0_])
            nw243_[int(9)] = _dafny.MultiSet([d_1304_v0_, d_1304_v0_, d_1304_v0_])
            nw243_[int(10)] = _dafny.MultiSet([d_1304_v0_])
            nw243_[int(11)] = (d_1335_v27_).set(d_1304_v0_, default__.abs(895))
            nw243_[int(12)] = d_1335_v27_
            nw243_[int(13)] = (d_1335_v27_) - (d_1335_v27_)
            nw243_[int(14)] = d_1335_v27_
            nw243_[int(15)] = d_1335_v27_
            nw243_[int(16)] = _dafny.MultiSet([d_1304_v0_, d_1304_v0_, d_1304_v0_])
            nw243_[int(17)] = d_1335_v27_
            nw243_[int(18)] = d_1335_v27_
            nw243_[int(19)] = (d_1335_v27_ if ((d_1343_v35_)[d_1304_v0_] if (d_1304_v0_) in (d_1343_v35_) else d_1304_v0_) else default__.fm28(d_1338_v30_, True, len(d_1343_v35_), default__.fm2(((d_1344_v36_)[len(_dafny.SeqWithoutIsStrInference([(self).f1 for d_1346_i4_ in range(default__.abs(467))]))] if (len(_dafny.SeqWithoutIsStrInference([(self).f1 for d_1347_i4_ in range(default__.abs(467))]))) in (d_1344_v36_) else d_1304_v0_), p1, d_1304_v0_, d_1304_v0_, globalState), globalState))
            d_1345_v37_ = nw243_
            index235_ = default__.safeIndex(642, (d_1345_v37_).length(0))
            (d_1345_v37_)[index235_] = _dafny.MultiSet([(d_1337_v29_).fm5(globalState)])
            index236_ = default__.safeIndex(642, (d_1345_v37_).length(0))
            (d_1345_v37_)[index236_] = d_1335_v27_
        d_1348_v38_: _dafny.Map
        d_1348_v38_ = _dafny.Map({(self).f2: (self).f2})
        d_1349_v39_: _dafny.Set
        d_1349_v39_ = _dafny.Set({len(d_1348_v38_), (self).f2})
        d_1350_i5_: int
        d_1350_i5_ = 0
        with _dafny.label("9"):
            while not (d_1304_v0_) or (((self).f0) in (d_1349_v39_)):
                with _dafny.c_label("9"):
                    if (d_1350_i5_) >= (100):
                        raise _dafny.Break("9")
                    d_1350_i5_ = (d_1350_i5_) + (1)
                    d_1304_v0_ = not(((self).f0) < ((self).f2))
                    d_1351_v40_: _dafny.Seq
                    d_1351_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjy"))
                    d_1352_v41_: _dafny.Map
                    d_1352_v41_ = _dafny.Map({_dafny.Map({d_1304_v0_: p1}): (0) - (len(d_1351_v40_))})
                    d_1353_v42_: _dafny.Map
                    d_1353_v42_ = _dafny.Map({d_1304_v0_: (self).f0})
                    d_1352_v41_ = (d_1352_v41_).set(d_1353_v42_, (self).f2)
                    d_1304_v0_ = d_1304_v0_
                    if d_1304_v0_:
                        d_1354_v43_: _dafny.Array
                        nw244_ = _dafny.Array(_dafny.Array(None, 0), 3)
                        d_1354_v43_ = nw244_
                        d_1355_v44_: _dafny.Array
                        nw245_ = _dafny.Array(False, 11)
                        d_1355_v44_ = nw245_
                        index237_ = default__.safeIndex(797, (d_1354_v43_).length(0))
                        (d_1354_v43_)[index237_] = d_1355_v44_
                        index238_ = default__.safeIndex(797, (d_1354_v43_).length(0))
                        (d_1354_v43_)[index238_] = d_1355_v44_
                        d_1356_v45_: _dafny.Map
                        d_1356_v45_ = _dafny.Map({d_1351_v40_: (self).f2})
                        d_1357_v46_: _dafny.MultiSet
                        d_1357_v46_ = _dafny.MultiSet([(self).f0, default__.fm3(d_1305_v1_, _dafny.CodePoint('p'), ((d_1356_v45_)[_dafny.SeqWithoutIsStrInference([(self).f1 for d_1358_i6_ in range(default__.abs(919))])] if (_dafny.SeqWithoutIsStrInference([(self).f1 for d_1359_i6_ in range(default__.abs(919))])) in (d_1356_v45_) else p1), 328, globalState)])
                        d_1360_v47_: _dafny.MultiSet
                        d_1360_v47_ = _dafny.MultiSet([d_1357_v46_, d_1357_v46_])
                        r0 = ((_dafny.MultiSet([_dafny.MultiSet([(self).f0]), d_1357_v46_, d_1357_v46_, _dafny.MultiSet([(self).f2]), d_1357_v46_])) - (d_1360_v47_)).cardinality
                        d_1361_v48_: C0
                        nw246_ = C0()
                        nw246_.ctor__((self).f1)
                        d_1361_v48_ = nw246_
                        d_1351_v40_ = _dafny.SeqWithoutIsStrInference([(d_1361_v48_).f3 for d_1362_i7_ in range(default__.abs(-405))])
                        def iife77_():
                            def iife79_():
                                coll29_ = _dafny.Map()
                                compr_29_: int
                                for compr_29_ in _dafny.IntegerRange(944, -345):
                                    d_1363_v49_: int = compr_29_
                                    if ((944) <= (d_1363_v49_)) and ((d_1363_v49_) < (-345)):
                                        coll29_[(d_1363_v49_) * (p1)] = (self).f2
                                return _dafny.Map(coll29_)
                            coll27_ = _dafny.Set()
                            def iife78_():
                                coll28_ = _dafny.Map()
                                compr_28_: int
                                for compr_28_ in _dafny.IntegerRange(944, -345):
                                    d_1363_v49_: int = compr_28_
                                    if ((944) <= (d_1363_v49_)) and ((d_1363_v49_) < (-345)):
                                        coll28_[(d_1363_v49_) * (p1)] = (self).f2
                                return _dafny.Map(coll28_)
                            compr_27_: int
                            for compr_27_ in ((_dafny.Map({(self).f0: p1})) | (iife78_()
                            )).keys.Elements:
                                d_1364_v50_: int = compr_27_
                                if (d_1364_v50_) in ((_dafny.Map({(self).f0: p1})) | (iife79_()
                                )):
                                    coll27_ = coll27_.union(_dafny.Set([(d_1364_v50_) + ((0) - (len((_dafny.Map({204: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1365_i8_ in range(default__.abs(872))]))}) if not(False) else _dafny.Map({968: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pp")))}))]))})))))]))
                            return _dafny.Set(coll27_)
                        d_1349_v39_ = iife77_()
                        
                    elif True:
                        d_1366_v51_: C1
                        nw247_ = C1()
                        nw247_.ctor__(580)
                        d_1366_v51_ = nw247_
                        d_1366_v51_ = d_1366_v51_
                        d_1367_v52_: T0
                        nw248_ = C3()
                        nw248_.ctor__((self).f0)
                        d_1367_v52_ = nw248_
                        d_1368_v53_: _dafny.Map
                        d_1368_v53_ = _dafny.Map({default__.fm3(d_1305_v1_, (self).f1, (self).f2, len(d_1351_v40_), globalState): d_1367_v52_})
                        d_1369_v54_: _dafny.Seq
                        d_1369_v54_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_1370_v55_: _dafny.Seq
                        d_1370_v55_ = _dafny.SeqWithoutIsStrInference([d_1304_v0_, d_1304_v0_])
                        d_1371_v56_: _dafny.Array
                        nw249_ = _dafny.Array(None, 11)
                        nw249_[int(0)] = (self).f2
                        nw249_[int(1)] = len(d_1368_v53_)
                        nw249_[int(2)] = (self).f2
                        nw249_[int(3)] = (0) - ((d_1369_v54_)[default__.safeIndex((self).f0, len(d_1369_v54_))])
                        nw249_[int(4)] = p1
                        nw249_[int(5)] = (self).f2
                        nw249_[int(6)] = default__.safeDivisionInt(p1, len(d_1348_v38_))
                        nw249_[int(7)] = (self).f2
                        nw249_[int(8)] = (0) - ((self).f2)
                        nw249_[int(9)] = ((self).f0) + (len(_dafny.Set({d_1304_v0_, d_1304_v0_})))
                        nw249_[int(10)] = (p1) - ((d_1366_v51_).fm15(_dafny.Map({_dafny.SeqWithoutIsStrInference([(D5_DC13(209, d_1305_v1_, (self).f1, (self).f2, True)).cf19 for d_1372_i9_ in range(default__.abs(880))]): (self).fm5(globalState)}), (d_1370_v55_)[default__.safeIndex((self).f2, len(d_1370_v55_))], True, globalState))
                        d_1371_v56_ = nw249_
                        index239_ = default__.safeIndex(170, (d_1371_v56_).length(0))
                        (d_1371_v56_)[index239_] = ((self).f2) - (p1)
                        d_1373_v57_: _dafny.Map
                        d_1373_v57_ = _dafny.Map({d_1351_v40_: d_1304_v0_})
                        d_1374_v58_: _dafny.MultiSet
                        d_1374_v58_ = _dafny.MultiSet([d_1304_v0_, (d_1370_v55_)[default__.safeIndex((self).f0, len(d_1370_v55_))], d_1304_v0_])
                        index240_ = default__.safeIndex(170, (d_1371_v56_).length(0))
                        rhs229_ = (0) - (default__.safeModuloInt((d_1369_v54_)[default__.safeIndex((self).f0, len(d_1369_v54_))], len(_dafny.SeqWithoutIsStrInference([(self).f1 for d_1375_i10_ in range(default__.abs(566))]))))
                        rhs230_ = (d_1366_v51_).fm15(d_1373_v57_, False, (((d_1374_v58_)[True] if (True) in (d_1374_v58_) else default__.fm3(d_1305_v1_, (self).f1, (self).f2, (d_1367_v52_).f0, globalState))) in ((d_1348_v38_).set((d_1367_v52_).f0, (self).f0)), globalState)
                        rhs231_ = (d_1367_v52_).f0
                        lhs104_ = d_1371_v56_
                        lhs105_ = default__.safeIndex(170, (d_1371_v56_).length(0))
                        r2 = rhs229_
                        lhs104_[lhs105_] = rhs230_
                        r0 = rhs231_
                        d_1376_v59_: D6
                        d_1376_v59_ = D6_DC14(d_1371_v56_)
                        d_1376_v59_ = d_1376_v59_
                        d_1371_v56_ = d_1371_v56_
                        d_1304_v0_ = d_1304_v0_
                    pass
            pass
        d_1377_v60_: _dafny.Seq
        d_1377_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jiaaicmv"))
        d_1378_v61_: _dafny.Map
        d_1378_v61_ = _dafny.Map({d_1377_v60_: d_1304_v0_})
        d_1378_v61_ = (d_1378_v61_).set(_dafny.SeqWithoutIsStrInference([(self).f1 for d_1379_i11_ in range(default__.abs(904))]), d_1304_v0_)
        d_1380_v62_: T1
        nw250_ = C2()
        nw250_.ctor__((self).f0)
        d_1380_v62_ = nw250_
        d_1381_v63_: _dafny.Map
        d_1381_v63_ = _dafny.Map({(0) - (p1): d_1380_v62_})
        d_1382_v64_: D6
        d_1382_v64_ = D6_DC15(len((_dafny.SeqWithoutIsStrInference([p1 for d_1383_i12_ in range(default__.abs(299))])) + (default__.fm0(globalState))), (d_1381_v63_) == (d_1381_v63_), (self).f0)
        source17_ = d_1382_v64_
        if source17_.is_DC15:
            d_1384___mcc_h0_ = source17_.cf23
            d_1385___mcc_h1_ = source17_.cf24
            d_1386___mcc_h2_ = source17_.cf25
            d_1387_cf25_ = d_1386___mcc_h2_
            d_1388_cf24_ = d_1385___mcc_h1_
            d_1389_cf23_ = d_1384___mcc_h0_
            d_1390_v65_: C1
            nw251_ = C1()
            nw251_.ctor__((d_1380_v62_).f0)
            d_1390_v65_ = nw251_
            source18_ = (self).f1
            d_1391___mcc_h4_ = source18_
            d_1392_cf13_ = d_1391___mcc_h4_
            d_1393_v66_: D2
            d_1393_v66_ = D2_DC4(((d_1377_v60_).set(default__.safeIndex(default__.fm3(d_1305_v1_, (self).f1, (self).f2, d_1387_cf25_, globalState), len(d_1377_v60_)), (d_1377_v60_)[default__.safeIndex((self).f0, len(d_1377_v60_))])) + (d_1377_v60_), not((d_1304_v0_) == (d_1304_v0_)))
            d_1393_v66_ = d_1393_v66_
            pat_let_tv32_ = d_1377_v60_
            pat_let_tv33_ = d_1377_v60_
            pat_let_tv34_ = d_1392_cf13_
            def iife80_(_pat_let25_0):
                def iife81_(d_1394_dt__update__tmp_h0_):
                    def iife82_(_pat_let26_0):
                        def iife83_(d_1395_dt__update_hcf4_h0_):
                            return D2_DC4(d_1395_dt__update_hcf4_h0_, (d_1394_dt__update__tmp_h0_).cf5)
                        return iife83_(_pat_let26_0)
                    return iife82_((pat_let_tv32_).set(default__.safeIndex((self).f0, len(pat_let_tv33_)), pat_let_tv34_))
                return iife81_(_pat_let25_0)
            rhs232_ = iife80_(D2_DC4(d_1377_v60_, d_1304_v0_))
            rhs233_ = d_1304_v0_
            d_1393_v66_ = rhs232_
            d_1304_v0_ = rhs233_
            d_1304_v0_ = d_1388_cf24_
            d_1396_v67_: _dafny.Array
            nw252_ = _dafny.Array(None, 26)
            d_1396_v67_ = nw252_
            d_1397_v68_: D3
            d_1397_v68_ = D3_DC9(len(_dafny.Map({not(d_1304_v0_): (d_1380_v62_).f0})), d_1396_v67_, d_1304_v0_)
            d_1348_v38_ = (d_1348_v38_).set(d_1387_cf25_, (d_1397_v68_).cf10)
            d_1398_v69_: _dafny.MultiSet
            d_1398_v69_ = _dafny.MultiSet([not(d_1388_cf24_)])
            r2 = default__.fm3((d_1305_v1_) - (_dafny.Set({d_1304_v0_})), (self).f1, (0) - ((d_1390_v65_).fm15(d_1378_v61_, d_1304_v0_, d_1388_cf24_, globalState)), (579 if d_1304_v0_ else ((d_1398_v69_)[(d_1390_v65_).fm5(globalState)] if ((d_1390_v65_).fm5(globalState)) in (d_1398_v69_) else (self).f2)), globalState)
            d_1348_v38_ = d_1348_v38_
        elif True:
            d_1399___mcc_h3_ = source17_.cf22
            d_1400_cf22_ = d_1399___mcc_h3_
            d_1401_v70_: _dafny.Array
            nw253_ = _dafny.Array(_dafny.Set({}), 14)
            d_1401_v70_ = nw253_
            index241_ = default__.safeIndex(824, (d_1400_cf22_).length(0))
            (d_1400_cf22_)[index241_] = -340
            index242_ = default__.safeIndex(824, (d_1400_cf22_).length(0))
            rhs234_ = d_1401_v70_
            rhs235_ = ((self).f0) - (default__.safeModuloInt((d_1380_v62_).f0, p1))
            rhs236_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
            lhs106_ = d_1400_cf22_
            lhs107_ = default__.safeIndex(824, (d_1400_cf22_).length(0))
            d_1401_v70_ = rhs234_
            lhs106_[lhs107_] = rhs235_
            d_1377_v60_ = rhs236_
            d_1304_v0_ = d_1304_v0_
            d_1402_v71_: _dafny.MultiSet
            d_1402_v71_ = _dafny.MultiSet([(d_1380_v62_).f0, len(d_1377_v60_)])
            d_1305_v1_ = _dafny.Set({d_1304_v0_, d_1304_v0_, (d_1402_v71_).issubset(d_1402_v71_)})
            d_1403_v72_: _dafny.Map
            d_1403_v72_ = _dafny.Map({(p1) >= ((self).f2): d_1304_v0_})
            d_1403_v72_ = (d_1403_v72_ if d_1304_v0_ else d_1403_v72_)
        d_1404_v73_: _dafny.Array
        nw254_ = _dafny.Array(False, 8)
        d_1404_v73_ = nw254_
        d_1405_v74_: _dafny.Map
        d_1405_v74_ = _dafny.Map({d_1304_v0_: d_1404_v73_})
        d_1406_v75_: _dafny.MultiSet
        d_1406_v75_ = _dafny.MultiSet([p1])
        d_1404_v73_ = ((d_1405_v74_)[(d_1406_v75_) != (d_1406_v75_)] if ((d_1406_v75_) != (d_1406_v75_)) in (d_1405_v74_) else d_1404_v73_)
        d_1407_v76_: _dafny.Seq
        d_1407_v76_ = _dafny.SeqWithoutIsStrInference([p1])
        r0 = (((d_1348_v38_)[836] if (836) in (d_1348_v38_) else (0) - ((d_1407_v76_)[default__.safeIndex(p1, len(d_1407_v76_))]))) * (default__.safeDivisionInt((0) - (len(d_1348_v38_)), (self).f2))
        r1 = (self).f0
        r2 = (d_1380_v62_).f0
        return r0, r1, r2

    def m3(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: D2 = D2.default()()
        d_1408_v0_: bool
        d_1408_v0_ = True
        if d_1408_v0_:
            d_1409_v1_: _dafny.Set
            d_1409_v1_ = _dafny.Set({default__.fm2(not(d_1408_v0_), (p1).f0, d_1408_v0_, d_1408_v0_, globalState), False, (d_1408_v0_ if d_1408_v0_ else d_1408_v0_), not(d_1408_v0_), d_1408_v0_})
            d_1410_v2_: _dafny.Seq
            d_1410_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_1408_v0_})])
            d_1411_v3_: _dafny.Seq
            d_1411_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hxmu"))
            d_1409_v1_ = (d_1410_v2_)[default__.safeIndex(default__.fm3(_dafny.Set({False, d_1408_v0_, d_1408_v0_}), (self).f1, len(d_1411_v3_), (self).f0, globalState), len(d_1410_v2_))]
            d_1412_v4_: _dafny.Set
            d_1412_v4_ = _dafny.Set({p0, p0})
            d_1413_v5_: _dafny.MultiSet
            d_1413_v5_ = _dafny.MultiSet([(self).f0, len(d_1412_v4_)])
            d_1414_v6_: _dafny.Set
            d_1414_v6_ = _dafny.Set({d_1413_v5_, _dafny.MultiSet([-379]), d_1413_v5_})
            d_1414_v6_ = (d_1414_v6_) - (d_1414_v6_)
            d_1415_v7_: _dafny.Array
            nw255_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_1415_v7_ = nw255_
            d_1416_v10_: _dafny.Array
            nw256_ = _dafny.Array(_dafny.Seq({}), 5)
            d_1416_v10_ = nw256_
            d_1417_v11_: _dafny.Map
            def iife84_():
                def iife86_():
                    coll32_ = _dafny.Map()
                    compr_32_: int
                    for compr_32_ in _dafny.IntegerRange(758, -84):
                        d_1418_v8_: int = compr_32_
                        if ((758) <= (d_1418_v8_)) and ((d_1418_v8_) < (-84)):
                            coll32_[default__.safeModuloInt(d_1418_v8_, (_dafny.MultiSet([(self).f1])).cardinality)] = d_1408_v0_
                    return _dafny.Map(coll32_)
                coll30_ = _dafny.Set()
                def iife85_():
                    coll31_ = _dafny.Map()
                    compr_31_: int
                    for compr_31_ in _dafny.IntegerRange(758, -84):
                        d_1418_v8_: int = compr_31_
                        if ((758) <= (d_1418_v8_)) and ((d_1418_v8_) < (-84)):
                            coll31_[default__.safeModuloInt(d_1418_v8_, (_dafny.MultiSet([(self).f1])).cardinality)] = d_1408_v0_
                    return _dafny.Map(coll31_)
                compr_30_: int
                for compr_30_ in ((iife85_()
                ).set(648, d_1408_v0_)).keys.Elements:
                    d_1419_v9_: int = compr_30_
                    if (d_1419_v9_) in ((iife86_()
                    ).set(648, d_1408_v0_)):
                        coll30_ = coll30_.union(_dafny.Set([(d_1419_v9_) + (len(_dafny.Set({989, -246})))]))
                return _dafny.Set(coll30_)
            d_1417_v11_ = _dafny.Map({default__.fm3((d_1410_v2_)[default__.safeIndex((self).f2, len(d_1410_v2_))], (self).f1, (self).f2, len(iife84_()
            ), globalState): d_1416_v10_})
            index243_ = default__.safeIndex(742, (d_1415_v7_).length(0))
            (d_1415_v7_)[index243_] = ((d_1417_v11_)[(p1).f0] if ((p1).f0) in (d_1417_v11_) else d_1416_v10_)
            index244_ = default__.safeIndex(742, (d_1415_v7_).length(0))
            (d_1415_v7_)[index244_] = d_1416_v10_
            d_1420_v12_: int
            d_1420_v12_ = 399
            d_1420_v12_ = ((p1).f0) + (d_1420_v12_)
            d_1421_v13_: str
            d_1421_v13_ = _dafny.CodePoint('g')
            d_1421_v13_ = d_1421_v13_
        elif True:
            d_1422_v14_: _dafny.Seq
            d_1422_v14_ = _dafny.SeqWithoutIsStrInference([d_1408_v0_, d_1408_v0_])
            d_1423_v15_: _dafny.Seq
            d_1423_v15_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_1408_v0_, d_1408_v0_])) + (d_1422_v14_), _dafny.SeqWithoutIsStrInference([d_1408_v0_])])
            d_1422_v14_ = (d_1423_v15_)[default__.safeIndex((len(_dafny.Map({d_1408_v0_: (self).f1})) if d_1408_v0_ else (0) - ((self).f0)), len(d_1423_v15_))]
            d_1424_v16_: _dafny.Seq
            d_1424_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "biksd"))
            d_1425_v17_: _dafny.Seq
            d_1425_v17_ = _dafny.SeqWithoutIsStrInference([d_1424_v16_, d_1424_v16_])
            d_1426_v18_: _dafny.Map
            d_1426_v18_ = _dafny.Map({d_1408_v0_: False})
            d_1427_v19_: _dafny.Seq
            d_1427_v19_ = _dafny.SeqWithoutIsStrInference([d_1425_v17_, d_1425_v17_])
            rhs237_ = not (True) or (((d_1426_v18_)[(d_1422_v14_)[default__.safeIndex((p1).f0, len(d_1422_v14_))]] if ((d_1422_v14_)[default__.safeIndex((p1).f0, len(d_1422_v14_))]) in (d_1426_v18_) else True))
            rhs238_ = (d_1427_v19_)[default__.safeIndex((self).f0, len(d_1427_v19_))]
            r0 = rhs237_
            d_1425_v17_ = rhs238_
            d_1428_v20_: _dafny.Array
            def lambda63_(d_1429_v0_):
                def lambda64_(d_1430_i0_):
                    return d_1429_v0_

                return lambda64_

            init32_ = lambda63_(d_1408_v0_)
            nw257_ = _dafny.Array(None, 16)
            for i0_32_ in range(nw257_.length(0)):
                nw257_[i0_32_] = init32_(i0_32_)
            d_1428_v20_ = nw257_
            index245_ = default__.safeIndex(846, (d_1428_v20_).length(0))
            (d_1428_v20_)[index245_] = False
            index246_ = default__.safeIndex(846, (d_1428_v20_).length(0))
            (d_1428_v20_)[index246_] = True
            d_1431_v21_: _dafny.Map
            d_1431_v21_ = _dafny.Map({(self).f1: d_1408_v0_})
            d_1432_v22_: D7
            d_1432_v22_ = D7_DC17((d_1428_v20_)[default__.safeIndex(846, (d_1428_v20_).length(0))], not (((d_1431_v21_)[(self).f1] if ((self).f1) in (d_1431_v21_) else not(not((d_1428_v20_)[default__.safeIndex(846, (d_1428_v20_).length(0))])))) or ((d_1428_v20_)[default__.safeIndex(846, (d_1428_v20_).length(0))]), not ((d_1428_v20_)[default__.safeIndex(846, (d_1428_v20_).length(0))]) or (True))
            d_1433_v23_: _dafny.Map
            d_1433_v23_ = _dafny.Map({len(d_1426_v18_): p1})
            d_1432_v22_ = D7_DC17(False, d_1408_v0_, ((self).f0) in (d_1433_v23_))
            d_1434_v24_: _dafny.Set
            d_1434_v24_ = _dafny.Set({(d_1428_v20_)[default__.safeIndex(846, (d_1428_v20_).length(0))]})
            d_1435_v25_: _dafny.Set
            d_1435_v25_ = _dafny.Set({(p1).f0, len(_dafny.SeqWithoutIsStrInference([(p1).f0, len(d_1434_v24_)]))})
            d_1436_v26_: _dafny.MultiSet
            d_1436_v26_ = _dafny.MultiSet([len(d_1435_v25_), len(p2), (p1).f0])
            d_1437_v27_: D7
            d_1437_v27_ = D7_DC16(d_1422_v14_)
            d_1438_v28_: _dafny.MultiSet
            d_1438_v28_ = _dafny.MultiSet([(self).f1])
            d_1439_v29_: _dafny.Array
            nw258_ = _dafny.Array(None, 23)
            nw258_[int(0)] = d_1422_v14_
            nw258_[int(1)] = d_1422_v14_
            nw258_[int(2)] = _dafny.SeqWithoutIsStrInference([(d_1428_v20_)[default__.safeIndex(846, (d_1428_v20_).length(0))]])
            nw258_[int(3)] = d_1422_v14_
            nw258_[int(4)] = d_1422_v14_
            nw258_[int(5)] = d_1422_v14_
            nw258_[int(6)] = d_1422_v14_
            nw258_[int(7)] = _dafny.SeqWithoutIsStrInference([d_1408_v0_, d_1408_v0_])
            nw258_[int(8)] = d_1422_v14_
            nw258_[int(9)] = d_1422_v14_
            nw258_[int(10)] = d_1422_v14_
            nw258_[int(11)] = d_1422_v14_
            nw258_[int(12)] = d_1422_v14_
            nw258_[int(13)] = d_1422_v14_
            nw258_[int(14)] = d_1422_v14_
            nw258_[int(15)] = d_1422_v14_
            nw258_[int(16)] = d_1422_v14_
            nw258_[int(17)] = d_1422_v14_
            nw258_[int(18)] = d_1422_v14_
            nw258_[int(19)] = d_1422_v14_
            nw258_[int(20)] = d_1422_v14_
            nw258_[int(21)] = d_1422_v14_
            nw258_[int(22)] = d_1422_v14_
            d_1439_v29_ = nw258_
            d_1440_v30_: _dafny.Seq
            d_1440_v30_ = _dafny.SeqWithoutIsStrInference([d_1439_v29_])
            d_1441_v31_: _dafny.MultiSet
            d_1441_v31_ = _dafny.MultiSet([(d_1440_v30_)[default__.safeIndex(560, len(d_1440_v30_))], d_1439_v29_])
            d_1442_v32_: C0
            nw259_ = C0()
            nw259_.ctor__(_dafny.CodePoint('a'))
            d_1442_v32_ = nw259_
            d_1443_v33_: D13
            d_1443_v33_ = D13_DC26(d_1442_v32_)
            d_1444_v34_: _dafny.Map
            d_1444_v34_ = _dafny.Map({(p1).f0: d_1442_v32_})
            d_1445_v35_: _dafny.Array
            nw260_ = _dafny.Array(None, 27)
            nw260_[int(0)] = d_1442_v32_
            nw260_[int(1)] = d_1442_v32_
            nw260_[int(2)] = d_1442_v32_
            nw260_[int(3)] = d_1442_v32_
            nw260_[int(4)] = d_1442_v32_
            nw260_[int(5)] = d_1442_v32_
            nw260_[int(6)] = d_1442_v32_
            nw260_[int(7)] = d_1442_v32_
            nw260_[int(8)] = d_1442_v32_
            nw260_[int(9)] = d_1442_v32_
            nw260_[int(10)] = d_1442_v32_
            nw260_[int(11)] = d_1442_v32_
            nw260_[int(12)] = d_1442_v32_
            nw260_[int(13)] = d_1442_v32_
            nw260_[int(14)] = d_1442_v32_
            nw260_[int(15)] = d_1442_v32_
            nw260_[int(16)] = d_1442_v32_
            nw260_[int(17)] = (d_1443_v33_).cf42
            nw260_[int(18)] = d_1442_v32_
            nw260_[int(19)] = d_1442_v32_
            nw260_[int(20)] = d_1442_v32_
            nw260_[int(21)] = d_1442_v32_
            nw260_[int(22)] = ((d_1444_v34_)[(p1).f0] if ((p1).f0) in (d_1444_v34_) else d_1442_v32_)
            nw260_[int(23)] = d_1442_v32_
            nw260_[int(24)] = d_1442_v32_
            nw260_[int(25)] = d_1442_v32_
            nw260_[int(26)] = d_1442_v32_
            d_1445_v35_ = nw260_
            d_1446_v36_: D10
            d_1446_v36_ = D10_DC22(968, (d_1428_v20_)[default__.safeIndex(846, (d_1428_v20_).length(0))], not(d_1408_v0_))
            d_1447_v37_: _dafny.Map
            d_1447_v37_ = _dafny.Map({d_1408_v0_: (p1).f0})
            d_1448_v38_: _dafny.Set
            d_1448_v38_ = _dafny.Set({_dafny.Map({d_1408_v0_: d_1408_v0_})})
            d_1449_v39_: D5
            d_1449_v39_ = D5_DC13((self).f2, d_1434_v24_, (self).f1, (self).f2, d_1408_v0_)
            pat_let_tv35_ = p1
            pat_let_tv36_ = d_1442_v32_
            d_1450_v40_: _dafny.Array
            nw261_ = _dafny.Array(None, 28)
            nw261_[int(0)] = default__.safeDivisionInt(-289, (self).f0)
            nw261_[int(1)] = default__.safeDivisionInt(len(d_1426_v18_), (p1).f0)
            nw261_[int(2)] = (0) - ((73) + ((self).f0))
            nw261_[int(3)] = (0) - (default__.safeDivisionInt(p0, (p1).f0))
            nw261_[int(4)] = (p1).f0
            nw261_[int(5)] = (p1).f0
            nw261_[int(6)] = ((self).f0 if d_1408_v0_ else (p1).f0)
            nw261_[int(7)] = (self).f0
            nw261_[int(8)] = default__.safeModuloInt((p1).f0, ((d_1436_v26_).set((p1).f0, default__.abs(len((d_1437_v27_).cf26)))).cardinality)
            nw261_[int(9)] = -526
            nw261_[int(10)] = (p1).f0
            nw261_[int(11)] = (0) - (default__.fm3(default__.fm10(d_1438_v28_, (d_1428_v20_)[default__.safeIndex(846, (d_1428_v20_).length(0))], globalState), (self).f1, (p1).f0, (0) - ((p1).f0), globalState))
            nw261_[int(12)] = (0) - ((p1).f0)
            nw261_[int(13)] = (p1).f0
            nw261_[int(14)] = default__.safeModuloInt((self).f2, ((d_1441_v31_)[d_1439_v29_] if (d_1439_v29_) in (d_1441_v31_) else (D3_DC9((p1).f0, d_1445_v35_, (d_1428_v20_)[default__.safeIndex(846, (d_1428_v20_).length(0))])).cf10))
            nw261_[int(15)] = (d_1446_v36_).cf34
            nw261_[int(16)] = (p1).f0
            nw261_[int(17)] = (self).f0
            nw261_[int(18)] = (self).f2
            nw261_[int(19)] = len(d_1424_v16_)
            nw261_[int(20)] = len(d_1424_v16_)
            nw261_[int(21)] = (self).f0
            nw261_[int(22)] = (self).f0
            nw261_[int(23)] = ((0) - ((p1).f0)) * ((self).f2)
            nw261_[int(24)] = (p1).f0
            nw261_[int(25)] = default__.safeModuloInt((0) - (len(d_1447_v37_)), (p1).f0)
            nw261_[int(26)] = len(d_1448_v38_)
            def iife87_(_pat_let27_0):
                def iife88_(d_1451_dt__update__tmp_h0_):
                    def iife89_(_pat_let28_0):
                        def iife90_(d_1452_dt__update_hcf17_h0_):
                            def iife91_(_pat_let29_0):
                                def iife92_(d_1453_dt__update_hcf19_h0_):
                                    return D5_DC13(d_1452_dt__update_hcf17_h0_, (d_1451_dt__update__tmp_h0_).cf18, d_1453_dt__update_hcf19_h0_, (d_1451_dt__update__tmp_h0_).cf20, (d_1451_dt__update__tmp_h0_).cf21)
                                return iife92_(_pat_let29_0)
                            return iife91_((pat_let_tv36_).f3)
                        return iife90_(_pat_let28_0)
                    return iife89_((pat_let_tv35_).f0)
                return iife88_(_pat_let27_0)
            nw261_[int(27)] = len((iife87_(d_1449_v39_)).cf18)
            d_1450_v40_ = nw261_
            index247_ = default__.safeIndex(903, (d_1450_v40_).length(0))
            (d_1450_v40_)[index247_] = len(_dafny.SeqWithoutIsStrInference([p0]))
            index248_ = default__.safeIndex(903, (d_1450_v40_).length(0))
            (d_1450_v40_)[index248_] = ((p1).f0) * (814)
        d_1454_v41_: _dafny.Array
        def lambda65_(d_1455_i1_):
            return default__.safeModuloInt(d_1455_i1_, -498)

        init33_ = lambda65_
        nw262_ = _dafny.Array(None, 16)
        for i0_33_ in range(nw262_.length(0)):
            nw262_[i0_33_] = init33_(i0_33_)
        d_1454_v41_ = nw262_
        d_1456_v42_: _dafny.Seq
        d_1456_v42_ = _dafny.SeqWithoutIsStrInference([d_1454_v41_, d_1454_v41_])
        d_1457_v43_: _dafny.MultiSet
        d_1457_v43_ = _dafny.MultiSet([((p2).set(default__.safeIndex(p0, len(p2)), (0) - ((self).f0))) + (_dafny.SeqWithoutIsStrInference([(p1).f0, 531, len(d_1456_v42_), p0, p0]))])
        d_1458_v44_: _dafny.Array
        nw263_ = _dafny.Array(_dafny.CodePoint('D'), 15)
        d_1458_v44_ = nw263_
        index249_ = default__.safeIndex(776, (d_1458_v44_).length(0))
        (d_1458_v44_)[index249_] = (_dafny.CodePoint('k') if (self).fm5(globalState) else (self).f1)
        d_1459_v45_: int
        d_1459_v45_ = 113
        d_1460_v46_: _dafny.Set
        d_1460_v46_ = _dafny.Set({True})
        d_1461_v47_: _dafny.Seq
        d_1461_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "davdt"))
        index250_ = default__.safeIndex(776, (d_1458_v44_).length(0))
        rhs239_ = ((d_1457_v43_) | (d_1457_v43_)).intersection((_dafny.MultiSet([p2, p2])).set(p2, default__.abs(default__.fm3(d_1460_v46_, (self).f1, (self).f2, (p1).f0, globalState))))
        rhs240_ = not(d_1408_v0_)
        rhs241_ = (self).f1
        rhs242_ = default__.safeModuloInt((p1).f0, (d_1459_v45_) - (len(d_1461_v47_)))
        rhs243_ = d_1459_v45_
        lhs108_ = d_1458_v44_
        lhs109_ = default__.safeIndex(776, (d_1458_v44_).length(0))
        d_1457_v43_ = rhs239_
        r0 = rhs240_
        lhs108_[lhs109_] = rhs241_
        d_1459_v45_ = rhs242_
        d_1459_v45_ = rhs243_
        d_1462_v48_: C0
        nw264_ = C0()
        nw264_.ctor__(_dafny.CodePoint('h'))
        d_1462_v48_ = nw264_
        d_1463_v49_: D13
        d_1463_v49_ = D13_DC26(d_1462_v48_)
        source19_ = d_1463_v49_
        if source19_.is_DC27:
            d_1464___mcc_h0_ = source19_.cf43
            d_1465___mcc_h1_ = source19_.cf44
            d_1466___mcc_h2_ = source19_.cf45
            d_1467___mcc_h3_ = source19_.cf46
            d_1468_cf46_ = d_1467___mcc_h3_
            d_1469_cf45_ = d_1466___mcc_h2_
            d_1470_cf44_ = d_1465___mcc_h1_
            d_1471_cf43_ = d_1464___mcc_h0_
            d_1472_v50_: _dafny.Array
            nw265_ = _dafny.Array(_dafny.Seq({}), 11)
            d_1472_v50_ = nw265_
            d_1473_v51_: _dafny.Array
            nw266_ = _dafny.Array(False, 18)
            d_1473_v51_ = nw266_
            d_1474_v52_: _dafny.Seq
            d_1474_v52_ = _dafny.SeqWithoutIsStrInference([d_1473_v51_])
            index251_ = default__.safeIndex(471, (d_1472_v50_).length(0))
            (d_1472_v50_)[index251_] = d_1474_v52_
            d_1475_v53_: _dafny.MultiSet
            d_1475_v53_ = _dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference([d_1473_v51_])))])
            index252_ = default__.safeIndex(471, (d_1472_v50_).length(0))
            rhs244_ = d_1474_v52_
            rhs245_ = default__.fm22(globalState)
            lhs110_ = d_1472_v50_
            lhs111_ = default__.safeIndex(471, (d_1472_v50_).length(0))
            lhs110_[lhs111_] = rhs244_
            d_1475_v53_ = rhs245_
            d_1468_cf46_ = ((p2)[default__.safeIndex(45, len(p2))]) - (((0) - ((0) - (len(d_1461_v47_))) if d_1408_v0_ else (d_1470_cf44_).f0))
            d_1476_v54_: _dafny.Array
            def lambda66_(d_1477_v0_):
                def lambda67_(d_1478_i2_):
                    return (_dafny.SeqWithoutIsStrInference([d_1477_v0_])) + (_dafny.SeqWithoutIsStrInference([d_1477_v0_]))

                return lambda67_

            init34_ = lambda66_(d_1408_v0_)
            nw267_ = _dafny.Array(None, 9)
            for i0_34_ in range(nw267_.length(0)):
                nw267_[i0_34_] = init34_(i0_34_)
            d_1476_v54_ = nw267_
            d_1479_v55_: _dafny.Seq
            d_1479_v55_ = _dafny.SeqWithoutIsStrInference([d_1476_v54_, d_1476_v54_, d_1476_v54_])
            d_1476_v54_ = (d_1479_v55_)[default__.safeIndex((p1).f0, len(d_1479_v55_))]
            d_1480_v56_: D0
            d_1480_v56_ = D0_DC1(d_1408_v0_)
            d_1480_v56_ = d_1480_v56_
        elif source19_.is_DC28:
            d_1481___mcc_h4_ = source19_.cf47
            d_1482___mcc_h5_ = source19_.cf48
            d_1483_cf48_ = d_1482___mcc_h5_
            d_1484_cf47_ = d_1481___mcc_h4_
            d_1485_v57_: _dafny.Array
            def lambda68_(d_1486_i3_):
                return True

            init35_ = lambda68_
            nw268_ = _dafny.Array(None, 6)
            for i0_35_ in range(nw268_.length(0)):
                nw268_[i0_35_] = init35_(i0_35_)
            d_1485_v57_ = nw268_
            index253_ = default__.safeIndex(428, (d_1485_v57_).length(0))
            (d_1485_v57_)[index253_] = d_1408_v0_
            d_1487_v58_: _dafny.Seq
            d_1487_v58_ = _dafny.SeqWithoutIsStrInference([d_1408_v0_, d_1408_v0_])
            index254_ = default__.safeIndex(428, (d_1485_v57_).length(0))
            (d_1485_v57_)[index254_] = (d_1484_cf47_) and ((d_1487_v58_) <= (d_1487_v58_))
            d_1459_v45_ = (self).f0
            d_1488_v59_: _dafny.MultiSet
            d_1488_v59_ = _dafny.MultiSet([p0, -389])
            d_1489_v60_: _dafny.MultiSet
            d_1489_v60_ = d_1488_v59_
            d_1490_v61_: _dafny.Map
            d_1490_v61_ = _dafny.Map({(d_1489_v60_): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f2, (0) - (p0)]))).ispropersubset(d_1488_v59_)})
            rhs246_ = (p1).f0
            rhs247_ = default__.fm29(d_1408_v0_, d_1484_cf47_, d_1484_cf47_, d_1484_cf47_, globalState)
            d_1459_v45_ = rhs246_
            d_1490_v61_ = rhs247_
            d_1491_v62_: _dafny.Map
            d_1491_v62_ = _dafny.Map({(d_1487_v58_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gisbfojc"))), len(d_1487_v58_))]: d_1484_cf47_})
            d_1492_v63_: _dafny.Map
            d_1492_v63_ = _dafny.Map({(d_1491_v62_) | (_dafny.Map({True: d_1408_v0_})): (p0) > ((p1).f0)})
            d_1492_v63_ = _dafny.Map({d_1491_v62_: ((self).f2) <= ((p1).f0)})
        elif source19_.is_DC29:
            d_1493___mcc_h6_ = source19_.cf49
            d_1494___mcc_h7_ = source19_.cf50
            d_1495_cf50_ = d_1494___mcc_h7_
            d_1496_cf49_ = d_1493___mcc_h6_
            if d_1408_v0_:
                d_1497_v64_: _dafny.Seq
                d_1497_v64_ = _dafny.SeqWithoutIsStrInference([d_1408_v0_])
                d_1497_v64_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([(d_1462_v48_).f3 for d_1498_i4_ in range(default__.abs(751))])) <= (d_1461_v47_), d_1408_v0_, d_1495_cf50_])
                d_1499_v65_: _dafny.MultiSet
                d_1499_v65_ = _dafny.MultiSet([-707, (self).f0])
                index255_ = default__.safeIndex(36, (d_1454_v41_).length(0))
                (d_1454_v41_)[index255_] = ((_dafny.MultiSet(p2)) - (d_1499_v65_)).cardinality
                index256_ = default__.safeIndex(36, (d_1454_v41_).length(0))
                (d_1454_v41_)[index256_] = default__.safeModuloInt((d_1459_v45_) + ((p1).f0), (p1).f0)
                r0 = d_1495_cf50_
                d_1500_v66_: C2
                nw269_ = C2()
                nw269_.ctor__((d_1454_v41_)[default__.safeIndex(36, (d_1454_v41_).length(0))])
                d_1500_v66_ = nw269_
                d_1501_v67_: _dafny.MultiSet
                d_1501_v67_ = _dafny.MultiSet([d_1500_v66_])
                d_1495_cf50_ = ((d_1501_v67_).issubset(d_1501_v67_)) and (not((_dafny.CodePoint('w')) not in (d_1461_v47_)))
                d_1495_cf50_ = d_1408_v0_
            elif True:
                index257_ = default__.safeIndex(186, (d_1454_v41_).length(0))
                (d_1454_v41_)[index257_] = (self).f2
                index258_ = default__.safeIndex(186, (d_1454_v41_).length(0))
                (d_1454_v41_)[index258_] = (0) - (default__.safeModuloInt(default__.fm3(_dafny.Set({d_1408_v0_}), (self).f1, (p1).f0, (p1).f0, globalState), p0))
                d_1502_v68_: _dafny.Seq
                d_1502_v68_ = _dafny.SeqWithoutIsStrInference([False, d_1495_cf50_, d_1408_v0_])
                d_1503_v69_: D7
                d_1503_v69_ = D7_DC16(_dafny.SeqWithoutIsStrInference([d_1408_v0_, d_1495_cf50_, True, True]))
                d_1502_v68_ = (d_1503_v69_).cf26
                d_1504_v70_: _dafny.MultiSet
                d_1504_v70_ = _dafny.MultiSet([d_1503_v69_])
                d_1505_v71_: _dafny.Map
                d_1505_v71_ = _dafny.Map({(p1).f0: d_1495_cf50_})
                d_1506_v72_: _dafny.MultiSet
                d_1506_v72_ = _dafny.MultiSet([d_1459_v45_, (p1).f0])
                d_1507_v73_: _dafny.Set
                d_1507_v73_ = _dafny.Set({p0, (d_1506_v72_).cardinality})
                index259_ = default__.safeIndex(186, (d_1454_v41_).length(0))
                (d_1454_v41_)[index259_] = default__.fm3(d_1460_v46_, (d_1462_v48_).f3, default__.fm3(default__.fm27(d_1408_v0_, d_1504_v70_, len(d_1505_v71_), d_1507_v73_, globalState), (d_1462_v48_).f3, (p1).f0, (self).f0, globalState), default__.safeModuloInt((p1).f0, (p1).f0), globalState)
                d_1459_v45_ = len(_dafny.Set({(d_1454_v41_)[default__.safeIndex(186, (d_1454_v41_).length(0))]}))
                index260_ = default__.safeIndex(186, (d_1454_v41_).length(0))
                (d_1454_v41_)[index260_] = default__.fm3((d_1460_v46_).intersection(d_1460_v46_), _dafny.CodePoint('f'), len(d_1507_v73_), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "to")))) + ((self).f2), globalState)
            index261_ = default__.safeIndex(786, (d_1454_v41_).length(0))
            (d_1454_v41_)[index261_] = 86
            d_1508_v74_: _dafny.Map
            d_1508_v74_ = _dafny.Map({d_1496_cf49_: (d_1458_v44_)[default__.safeIndex(776, (d_1458_v44_).length(0))]})
            index262_ = default__.safeIndex(786, (d_1454_v41_).length(0))
            rhs248_ = ((p0) + (len(d_1461_v47_))) * ((p2)[default__.safeIndex(len((p2).set(default__.safeIndex(d_1459_v45_, len(p2)), (self).f0)), len(p2))])
            rhs249_ = len(d_1508_v74_)
            rhs250_ = not(d_1408_v0_)
            lhs112_ = d_1454_v41_
            lhs113_ = default__.safeIndex(786, (d_1454_v41_).length(0))
            d_1459_v45_ = rhs248_
            lhs112_[lhs113_] = rhs249_
            r0 = rhs250_
            d_1509_v75_: _dafny.Set
            d_1509_v75_ = _dafny.Set({p0, d_1459_v45_})
            d_1496_cf49_ = (self).fm7((len(default__.fm18((_dafny.MultiSet([d_1408_v0_])).cardinality, globalState))) + (len(_dafny.Set({not(not(d_1495_cf50_)), not(d_1495_cf50_)}))), (d_1454_v41_)[default__.safeIndex(786, (d_1454_v41_).length(0))], d_1495_cf50_, (0) - (default__.fm3(d_1460_v46_, (self).f1, len(d_1509_v75_), (0) - (p0), globalState)), globalState)
            d_1510_v76_: _dafny.Map
            d_1510_v76_ = _dafny.Map({(self).f2: d_1461_v47_})
            d_1511_v77_: D2
            d_1511_v77_ = D2_DC4((d_1461_v47_) + (((d_1510_v76_)[p0] if (p0) in (d_1510_v76_) else (d_1461_v47_).set(default__.safeIndex((d_1454_v41_)[default__.safeIndex(786, (d_1454_v41_).length(0))], len(d_1461_v47_)), (d_1462_v48_).f3))), ((d_1454_v41_)[default__.safeIndex(786, (d_1454_v41_).length(0))]) > ((self).f0))
            r1 = d_1511_v77_
        elif source19_.is_DC26:
            d_1512___mcc_h8_ = source19_.cf42
            d_1513_cf42_ = d_1512___mcc_h8_
            d_1514_v78_: _dafny.MultiSet
            d_1514_v78_ = _dafny.MultiSet([(p1).f0])
            r0 = not((d_1514_v78_).ispropersubset((d_1514_v78_ if d_1408_v0_ else d_1514_v78_)))
            d_1515_v79_: _dafny.Array
            nw270_ = _dafny.Array(None, 6)
            nw270_[int(0)] = d_1408_v0_
            nw270_[int(1)] = d_1408_v0_
            nw270_[int(2)] = d_1408_v0_
            nw270_[int(3)] = d_1408_v0_
            nw270_[int(4)] = d_1408_v0_
            nw270_[int(5)] = d_1408_v0_
            d_1515_v79_ = nw270_
            d_1516_v80_: _dafny.Seq
            d_1516_v80_ = _dafny.SeqWithoutIsStrInference([d_1515_v79_])
            d_1459_v45_ = default__.safeModuloInt(len(d_1516_v80_), (p1).f0)
            d_1517_v81_: _dafny.Array
            nw271_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_1517_v81_ = nw271_
            index263_ = default__.safeIndex(439, (d_1517_v81_).length(0))
            (d_1517_v81_)[index263_] = d_1454_v41_
            index264_ = default__.safeIndex(439, (d_1517_v81_).length(0))
            rhs251_ = d_1454_v41_
            rhs252_ = d_1515_v79_
            rhs253_ = (p1).f0
            rhs254_ = (d_1408_v0_) or ((default__.fm19(d_1514_v78_, (p1).f0, (self).f2, (p1).f0, globalState)) != (d_1461_v47_))
            rhs255_ = (p0 if (_dafny.MultiSet(p2)).issubset(d_1514_v78_) else ((self).f0) - (734))
            lhs114_ = d_1517_v81_
            lhs115_ = default__.safeIndex(439, (d_1517_v81_).length(0))
            lhs114_[lhs115_] = rhs251_
            d_1515_v79_ = rhs252_
            d_1459_v45_ = rhs253_
            r0 = rhs254_
            d_1459_v45_ = rhs255_
            d_1461_v47_ = ((d_1461_v47_) + (d_1461_v47_)) + (_dafny.SeqWithoutIsStrInference([(d_1513_cf42_).f3 for d_1518_i5_ in range(default__.abs(-630))]))
        elif True:
            d_1519___mcc_h9_ = source19_.cf51
            d_1520_cf51_ = d_1519___mcc_h9_
            d_1521_v82_: D0
            d_1521_v82_ = D0_DC0(d_1408_v0_)
            d_1522_v83_: _dafny.Seq
            d_1522_v83_ = _dafny.SeqWithoutIsStrInference([d_1521_v82_])
            d_1523_v84_: _dafny.Set
            d_1523_v84_ = _dafny.Set({(0) - (len(d_1522_v83_))})
            d_1524_v85_: _dafny.Seq
            d_1524_v85_ = _dafny.SeqWithoutIsStrInference([d_1523_v84_])
            d_1525_v86_: _dafny.MultiSet
            d_1525_v86_ = _dafny.MultiSet([len(d_1524_v85_)])
            d_1526_v87_: _dafny.Map
            d_1526_v87_ = _dafny.Map({p0: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyltafcwf")))})
            d_1527_v88_: _dafny.Seq
            d_1527_v88_ = _dafny.SeqWithoutIsStrInference([d_1525_v86_, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1459_v45_, p0, 222, (p1).f0]))])
            d_1528_v90_: _dafny.Map
            def iife93_():
                coll33_ = _dafny.Map()
                compr_33_: int
                for compr_33_ in _dafny.IntegerRange(798, -379):
                    d_1529_v89_: int = compr_33_
                    if ((798) <= (d_1529_v89_)) and ((d_1529_v89_) < (-379)):
                        coll33_[(d_1529_v89_) * ((self).f0)] = len(d_1461_v47_)
                return _dafny.Map(coll33_)
            d_1528_v90_ = _dafny.Map({iife93_()
            : _dafny.MultiSet([len(d_1461_v47_), default__.fm3(d_1460_v46_, (d_1462_v48_).f3, len(d_1523_v84_), (self).f0, globalState)])})
            d_1530_v91_: _dafny.Array
            nw272_ = _dafny.Array(None, 24)
            nw272_[int(0)] = d_1525_v86_
            nw272_[int(1)] = d_1525_v86_
            nw272_[int(2)] = d_1525_v86_
            nw272_[int(3)] = d_1525_v86_
            nw272_[int(4)] = (_dafny.MultiSet([-230])).set(len(d_1526_v87_), default__.abs((self).f2))
            nw272_[int(5)] = default__.fm22(globalState)
            nw272_[int(6)] = d_1525_v86_
            nw272_[int(7)] = _dafny.MultiSet([(0) - ((self).f2)])
            nw272_[int(8)] = d_1525_v86_
            nw272_[int(9)] = (d_1525_v86_).set((self).f0, default__.abs(p0))
            nw272_[int(10)] = (_dafny.MultiSet([(p1).f0])).intersection(d_1525_v86_)
            nw272_[int(11)] = (d_1527_v88_)[default__.safeIndex(467, len(d_1527_v88_))]
            nw272_[int(12)] = d_1525_v86_
            nw272_[int(13)] = (d_1527_v88_)[default__.safeIndex((self).f0, len(d_1527_v88_))]
            nw272_[int(14)] = (d_1525_v86_).intersection(d_1525_v86_)
            nw272_[int(15)] = d_1525_v86_
            nw272_[int(16)] = ((d_1528_v90_)[d_1526_v87_] if (d_1526_v87_) in (d_1528_v90_) else d_1525_v86_)
            nw272_[int(17)] = d_1525_v86_
            nw272_[int(18)] = d_1525_v86_
            nw272_[int(19)] = d_1525_v86_
            nw272_[int(20)] = (d_1525_v86_) - (default__.fm22(globalState))
            nw272_[int(21)] = d_1525_v86_
            nw272_[int(22)] = _dafny.MultiSet([(self).f2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cm")))])
            nw272_[int(23)] = (d_1525_v86_).intersection(d_1525_v86_)
            d_1530_v91_ = nw272_
            index265_ = default__.safeIndex(829, (d_1530_v91_).length(0))
            (d_1530_v91_)[index265_] = _dafny.MultiSet(p2)
            index266_ = default__.safeIndex(829, (d_1530_v91_).length(0))
            (d_1530_v91_)[index266_] = d_1525_v86_
            d_1531_v92_: _dafny.Array
            nw273_ = _dafny.Array(None, 3)
            d_1531_v92_ = nw273_
            d_1532_v93_: C3
            nw274_ = C3()
            nw274_.ctor__(d_1459_v45_)
            d_1532_v93_ = nw274_
            index267_ = default__.safeIndex(835, (d_1531_v92_).length(0))
            (d_1531_v92_)[index267_] = d_1532_v93_
            index268_ = default__.safeIndex(835, (d_1531_v92_).length(0))
            (d_1531_v92_)[index268_] = d_1532_v93_
            d_1533_v94_: _dafny.Map
            d_1533_v94_ = _dafny.Map({(p2) + (p2): d_1408_v0_})
            d_1533_v94_ = (d_1533_v94_) | (d_1533_v94_)
            d_1534_v95_: _dafny.Array
            nw275_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_1534_v95_ = nw275_
            d_1535_v96_: int
            d_1536_v97_: int
            d_1537_v98_: int
            out24_: int
            out25_: int
            out26_: int
            out24_, out25_, out26_ = (p1).m1(d_1534_v95_, (d_1459_v45_) + ((0) - (-94)), globalState)
            d_1535_v96_ = out24_
            d_1536_v97_ = out25_
            d_1537_v98_ = out26_
        d_1459_v45_ = (0) - (len(default__.fm0(globalState)))
        hi12_ = default__.safeDivisionInt((self).f2, (0) - ((self).f0))
        for d_1538_i6_ in range((p1).f0, hi12_):
            if not(d_1408_v0_):
                d_1539_v99_: _dafny.Array
                def lambda69_(d_1540_v0_):
                    def lambda70_(d_1541_i7_):
                        return _dafny.Map({d_1540_v0_: _dafny.SeqWithoutIsStrInference([d_1540_v0_, d_1540_v0_, d_1540_v0_])})

                    return lambda70_

                init36_ = lambda69_(d_1408_v0_)
                nw276_ = _dafny.Array(None, 1)
                for i0_36_ in range(nw276_.length(0)):
                    nw276_[i0_36_] = init36_(i0_36_)
                d_1539_v99_ = nw276_
                index269_ = default__.safeIndex(697, (d_1539_v99_).length(0))
                (d_1539_v99_)[index269_] = default__.fm30(d_1408_v0_, d_1408_v0_, d_1408_v0_, globalState)
                d_1542_v100_: D13
                d_1542_v100_ = D13_DC28(d_1408_v0_, (d_1462_v48_).f3)
                d_1543_v101_: _dafny.Map
                d_1543_v101_ = _dafny.Map({d_1408_v0_: _dafny.SeqWithoutIsStrInference([(d_1542_v100_).cf47, d_1408_v0_, d_1408_v0_, d_1408_v0_])})
                index270_ = default__.safeIndex(697, (d_1539_v99_).length(0))
                (d_1539_v99_)[index270_] = d_1543_v101_
                index271_ = default__.safeIndex(238, (d_1454_v41_).length(0))
                (d_1454_v41_)[index271_] = default__.safeDivisionInt(len(d_1461_v47_), (p1).f0)
                d_1544_v102_: _dafny.Seq
                d_1544_v102_ = _dafny.SeqWithoutIsStrInference([d_1408_v0_])
                d_1545_v103_: _dafny.MultiSet
                d_1545_v103_ = _dafny.MultiSet([(d_1544_v102_)[default__.safeIndex((p1).f0, len(d_1544_v102_))], not((d_1544_v102_)[default__.safeIndex((p1).f0, len(d_1544_v102_))]), d_1408_v0_, d_1408_v0_])
                index272_ = default__.safeIndex(238, (d_1454_v41_).length(0))
                rhs256_ = (d_1408_v0_) in (d_1545_v103_)
                rhs257_ = ((p1).f0 if d_1408_v0_ else (p1).f0)
                lhs116_ = d_1454_v41_
                lhs117_ = default__.safeIndex(238, (d_1454_v41_).length(0))
                r0 = rhs256_
                lhs116_[lhs117_] = rhs257_
                d_1546_v104_: _dafny.Set
                d_1546_v104_ = _dafny.Set({d_1461_v47_})
                d_1547_v105_: _dafny.Seq
                d_1547_v105_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1548_v106_: C3
                nw277_ = C3()
                nw277_.ctor__((p1).f0)
                d_1548_v106_ = nw277_
                d_1549_v107_: _dafny.MultiSet
                d_1549_v107_ = _dafny.MultiSet([d_1548_v106_])
                d_1550_v108_: _dafny.MultiSet
                d_1550_v108_ = _dafny.MultiSet([(self).f2, (p1).f0, d_1459_v45_, len(d_1547_v105_), len(_dafny.SeqWithoutIsStrInference([(p1).f0, p0, (d_1549_v107_).cardinality]))])
                d_1551_v109_: D7
                d_1551_v109_ = D7_DC18((d_1454_v41_)[default__.safeIndex(238, (d_1454_v41_).length(0))])
                d_1408_v0_ = (self).fm6((d_1546_v104_).issubset(_dafny.Set({d_1461_v47_, d_1461_v47_})), default__.fm19(d_1550_v108_, (self).f2, (p1).f0, (d_1551_v109_).cf30, globalState), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vr"))) + (d_1461_v47_), -812, globalState)
                index273_ = default__.safeIndex(238, (d_1454_v41_).length(0))
                (d_1454_v41_)[index273_] = (p1).f0
                d_1408_v0_ = d_1408_v0_
            elif True:
                d_1552_v110_: _dafny.Map
                d_1552_v110_ = _dafny.Map({d_1461_v47_: d_1408_v0_})
                r0 = (d_1552_v110_) != (d_1552_v110_)
                d_1553_v111_: _dafny.Map
                d_1553_v111_ = _dafny.Map({(p1).f0: default__.fm2(d_1408_v0_, (self).f2, not(d_1408_v0_), d_1408_v0_, globalState)})
                d_1553_v111_ = (d_1553_v111_).set(d_1459_v45_, False)
                d_1459_v45_ = (d_1459_v45_) - (len((d_1461_v47_) + (d_1461_v47_)))
                d_1554_v112_: _dafny.Array
                nw278_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
                d_1554_v112_ = nw278_
                index274_ = default__.safeIndex(880, (d_1554_v112_).length(0))
                (d_1554_v112_)[index274_] = _dafny.SeqWithoutIsStrInference([(d_1458_v44_)[default__.safeIndex(776, (d_1458_v44_).length(0))], (d_1462_v48_).f3])
                index275_ = default__.safeIndex(880, (d_1554_v112_).length(0))
                (d_1554_v112_)[index275_] = d_1461_v47_
                d_1555_v113_: _dafny.MultiSet
                d_1555_v113_ = _dafny.MultiSet([d_1538_i6_])
                d_1556_v114_: _dafny.Seq
                d_1556_v114_ = _dafny.SeqWithoutIsStrInference([default__.fm19(d_1555_v113_, (p1).f0, (p1).f0, 180, globalState)])
                d_1557_v115_: _dafny.Seq
                d_1557_v115_ = d_1556_v114_
                d_1556_v114_ = (d_1557_v115_)
            d_1558_v116_: _dafny.Seq
            d_1558_v116_ = _dafny.SeqWithoutIsStrInference([d_1408_v0_, (-270) >= ((self).f0)])
            d_1558_v116_ = ((d_1558_v116_).set(default__.safeIndex(939, len(d_1558_v116_)), d_1408_v0_)) + ((_dafny.SeqWithoutIsStrInference([True, d_1408_v0_, not(d_1408_v0_)])).set(default__.safeIndex((p1).f0, len(_dafny.SeqWithoutIsStrInference([True, d_1408_v0_, not(d_1408_v0_)]))), d_1408_v0_))
            index276_ = default__.safeIndex(776, (d_1458_v44_).length(0))
            (d_1458_v44_)[index276_] = (d_1462_v48_).f3
            d_1559_v117_: _dafny.Array
            nw279_ = _dafny.Array(_dafny.Set({}), 24)
            d_1559_v117_ = nw279_
            index277_ = default__.safeIndex(318, (d_1559_v117_).length(0))
            (d_1559_v117_)[index277_] = default__.fm31(globalState)
            d_1560_v118_: _dafny.Set
            d_1560_v118_ = _dafny.Set({_dafny.CodePoint('g')})
            index278_ = default__.safeIndex(318, (d_1559_v117_).length(0))
            (d_1559_v117_)[index278_] = (d_1560_v118_) | (d_1560_v118_)
        d_1561_v119_: _dafny.Seq
        d_1561_v119_ = _dafny.SeqWithoutIsStrInference([d_1461_v47_, d_1461_v47_, d_1461_v47_])
        d_1461_v47_ = (d_1461_v47_) + (((d_1561_v119_)[default__.safeIndex((self).f0, len(d_1561_v119_))]) + (d_1461_v47_))
        r0 = ((self).f1) not in ((d_1461_v47_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdiwmui"))).set(default__.safeIndex(d_1459_v45_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdiwmui")))), (self).f1)))
        r1 = D2_DC4(d_1461_v47_, not(d_1408_v0_))
        return r0, r1

    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
