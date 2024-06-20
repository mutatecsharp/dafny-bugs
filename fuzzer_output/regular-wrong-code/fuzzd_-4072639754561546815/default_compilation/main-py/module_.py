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
    def fm0(p0, globalState):
        if not(False):
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_0_i0_ in range(default__.abs(863))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qy")))
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1_i1_ in range(default__.abs(108))])

    @staticmethod
    def fm1(p0, globalState):
        return ((716) + ((0) - (len(_dafny.Map({len(_dafny.Map({True: (D9_DC25((0) - (-341), False, True, not(True))).cf44})): _dafny.CodePoint('b')}))))) + (28)

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([True, True])) == (_dafny.SeqWithoutIsStrInference([False, False, False, False, not(True)])), not((-680) not in (_dafny.SeqWithoutIsStrInference([74]))), not(False)])

    @staticmethod
    def fm3(globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hipjlan")))) <= ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csbbvc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejenrea"))))

    @staticmethod
    def fm4(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, False, True]), _dafny.MultiSet([True, False, True, True])])

    @staticmethod
    def fm5(p0, p1, globalState):
        if (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2_i0_ in range(default__.abs(528))]))) <= (len(_dafny.Set({False, False}))):
            return _dafny.CodePoint('q')
        elif True:
            return _dafny.CodePoint('v')

    @staticmethod
    def fm7(p0, p1, globalState):
        return len((_dafny.SeqWithoutIsStrInference([986])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(False): True}))])))

    @staticmethod
    def fm8(p0, p1, globalState):
        return (-717) == (45)

    @staticmethod
    def fm10(p0, p1, globalState):
        return not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wueqr"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_3_i0_ in range(default__.abs(631))]))))

    @staticmethod
    def fm11(globalState):
        return (_dafny.SeqWithoutIsStrInference([640])) + ((_dafny.SeqWithoutIsStrInference([-931, 501])) + (_dafny.SeqWithoutIsStrInference([97])))

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False, False, not(False), True, True])) + (_dafny.SeqWithoutIsStrInference([True]))) + ((_dafny.SeqWithoutIsStrInference([True, not(False)])) + (_dafny.SeqWithoutIsStrInference([not(True), False])))

    @staticmethod
    def fm15(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in (_dafny.Map({861: True})).keys.Elements:
                d_5_v0_: int = compr_0_
                if (d_5_v0_) in (_dafny.Map({861: True})):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeModuloInt(d_5_v0_, 63)]))
            return _dafny.Set(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.Map({-462: _dafny.CodePoint('g')})).keys.Elements:
                d_6_v1_: int = compr_1_
                if (d_6_v1_) in (_dafny.Map({-462: _dafny.CodePoint('g')})):
                    def iife2_():
                        coll2_ = _dafny.Map()
                        compr_2_: str
                        for compr_2_ in (_dafny.Map({_dafny.CodePoint('q'): False})).keys.Elements:
                            d_7_v2_: str = compr_2_
                            if (d_7_v2_) in (_dafny.Map({_dafny.CodePoint('q'): False})):
                                coll2_[d_7_v2_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xgonjaa")))
                        return _dafny.Map(coll2_)
                    coll1_[(d_6_v1_) * (len(_dafny.SeqWithoutIsStrInference([103])))] = _dafny.Map({len(iife2_()
                    ): 106})
            return _dafny.Map(coll1_)
        return (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([len((D17_DC46(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_4_i0_ in range(default__.abs(-338))])): (_dafny.MultiSet([True, False])).cardinality}))).cf80)]))), len(_dafny.Map({len(iife0_()
        ): not(True)})), 190, len(iife1_()
        ), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_8_i1_ in range(default__.abs(867))]))])) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p')]))).cardinality, 62, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_9_i3_ in range(default__.abs(369))])) for d_10_i2_ in range(default__.abs(705))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbdseslwf"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))]))

    @staticmethod
    def fm18(p0, p1, globalState):
        return (D17_DC47(not(False), (0) - (-432))).cf82

    @staticmethod
    def fm19(p0, p1, globalState):
        return (D18_DC48(_dafny.Map({not(False): -652}))).cf83

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('m')

    @staticmethod
    def fm21(p0, p1, p2, globalState):
        return (_dafny.Map({707: False})) | (_dafny.Map({248: True}))

    @staticmethod
    def fm22(p0, globalState):
        if True:
            def iife3_():
                coll3_ = _dafny.Set()
                compr_3_: int
                for compr_3_ in _dafny.IntegerRange(680, 4):
                    d_12_v0_: int = compr_3_
                    if ((680) <= (d_12_v0_)) and ((d_12_v0_) < (4)):
                        coll3_ = coll3_.union(_dafny.Set([(d_12_v0_) + ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqm")))])).cardinality)]))
                return _dafny.Set(coll3_)
            return (_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.Map({-540: False}))}) for d_11_i0_ in range(default__.abs(414))])) + (_dafny.SeqWithoutIsStrInference([iife3_()
            ]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.Set({322, (0) - (-322), (_dafny.MultiSet([-173])).cardinality, (0) - (-996), -776}) for d_13_i1_ in range(default__.abs(783))])

    @staticmethod
    def fm23(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([-110, -249])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([916])) for d_14_i0_ in range(default__.abs(-74))]))) + (_dafny.SeqWithoutIsStrInference([378, len(_dafny.Map({11: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucvrgkh"))})), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_15_i1_ in range(default__.abs(178))])): _dafny.CodePoint('t')})), 933]))

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        return _dafny.MultiSet([32])

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(783, -172):
                d_16_v0_: int = compr_4_
                if ((783) <= (d_16_v0_)) and ((d_16_v0_) < (-172)):
                    coll4_[default__.safeDivisionInt(d_16_v0_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([682, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlnaqhjk"))))])) for d_17_i0_ in range(default__.abs(208))])))] = False
            return _dafny.Map(coll4_)
        return (_dafny.Set({D2_DC6((0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False)]))).cardinality))})).intersection(_dafny.Set({D2_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "msuhavx")))), D2_DC6(-218), D2_DC6(len(iife4_()
)), D2_DC6(531)}))

    @staticmethod
    def fm26(p0, p1, p2, p3, globalState):
        return (855) == (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfehftfco"))), (0) - (len(_dafny.Set({_dafny.CodePoint('u'), _dafny.CodePoint('n')})))))

    @staticmethod
    def fm27(p0, p1, p2, globalState):
        if (_dafny.MultiSet([False, True, True])).isdisjoint(_dafny.MultiSet([True, False, False, not(not(False))])):
            return D9_DC23(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_18_i0_ in range(default__.abs(984))])]))
        elif True:
            return D9_DC23(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nttb"))]))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        return _dafny.Set({len(_dafny.Map({False: -747})), -663, default__.safeDivisionInt(-818, len(_dafny.Map({D9_DC25(933, False, True, False): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-790: False})), 80]))).cardinality})))})

    @staticmethod
    def fm29(p0, p1, globalState):
        if False:
            return D0_DC1(_dafny.CodePoint('c'), False, True)
        elif True:
            return D0_DC1(_dafny.CodePoint('j'), False, False)

    @staticmethod
    def fm30(p0, p1, globalState):
        return (_dafny.Map({not(True): True})) | (_dafny.Map({False: True}))

    @staticmethod
    def fm31(p0, globalState):
        return D0_DC3(_dafny.CodePoint('u'))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(-687, 640):
                d_19_v0_: int = compr_5_
                if ((-687) <= (d_19_v0_)) and ((d_19_v0_) < (640)):
                    coll5_[(d_19_v0_) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([604])) for d_20_i0_ in range(default__.abs(379))])))] = -983
            return _dafny.Map(coll5_)
        return (iife5_()
        ) | (_dafny.Map({(_dafny.MultiSet([-585])).cardinality: -836}))

    @staticmethod
    def fm33(p0, p1, p2, globalState):
        return D8_DC21((not(False) if True else True), not(not((True) and (True))), ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "td")))])))).cardinality)

    @staticmethod
    def fm34(p0, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: str
            for compr_6_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('j')])).Elements:
                d_21_v0_: str = compr_6_
                if (d_21_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('j')])):
                    coll6_[d_21_v0_] = False
            return _dafny.Map(coll6_)
        return (D19_DC52(_dafny.MultiSet([_dafny.MultiSet([len(_dafny.Set({not(False), True, True}))]), _dafny.MultiSet([len(_dafny.Map({len(iife6_()
): 965})), (_dafny.MultiSet([719])).cardinality, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_22_i0_ in range(default__.abs(461))])))]), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vo")))]), _dafny.MultiSet([69]), _dafny.MultiSet([856, (_dafny.MultiSet([True])).cardinality])]))).cf91

    @staticmethod
    def fm35(p0, p1, p2, globalState):
        source0_ = D17_DC47(False, 897)
        if source0_.is_DC47:
            d_23___mcc_h0_ = source0_.cf81
            d_24___mcc_h1_ = source0_.cf82
            d_25_cf82_ = d_24___mcc_h1_
            d_26_cf81_ = d_23___mcc_h0_
            return (_dafny.MultiSet([_dafny.Map({d_26_cf81_: len(_dafny.SeqWithoutIsStrInference([d_26_cf81_]))}), _dafny.Map({d_26_cf81_: d_25_cf82_})])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({d_26_cf81_: d_25_cf82_})])))
        elif True:
            d_27___mcc_h2_ = source0_.cf80
            d_28_cf80_ = d_27___mcc_h2_
            def iife7_():
                def iife9_():
                    coll9_ = _dafny.Set()
                    compr_9_: int
                    for compr_9_ in _dafny.IntegerRange(869, 577):
                        d_31_v1_: int = compr_9_
                        if ((869) <= (d_31_v1_)) and ((d_31_v1_) < (577)):
                            coll9_ = coll9_.union(_dafny.Set([default__.safeModuloInt(d_31_v1_, len(_dafny.Map({True: 569})))]))
                    return _dafny.Set(coll9_)
                coll7_ = _dafny.Map()
                def iife8_():
                    coll8_ = _dafny.Set()
                    compr_8_: int
                    for compr_8_ in _dafny.IntegerRange(869, 577):
                        d_29_v1_: int = compr_8_
                        if ((869) <= (d_29_v1_)) and ((d_29_v1_) < (577)):
                            coll8_ = coll8_.union(_dafny.Set([default__.safeModuloInt(d_29_v1_, len(_dafny.Map({True: 569})))]))
                    return _dafny.Set(coll8_)
                compr_7_: _dafny.Set
                for compr_7_ in (_dafny.Map({iife8_()
                : 89})).keys.Elements:
                    d_30_v0_: _dafny.Set = compr_7_
                    if (d_30_v0_) in (_dafny.Map({iife9_()
                    : 89})):
                        coll7_[d_30_v0_] = not(False)
                return _dafny.Map(coll7_)
            return (_dafny.MultiSet([_dafny.Map({True: len(iife7_()
            )})])).intersection(_dafny.MultiSet([_dafny.Map({True: len(_dafny.Map({False: False}))})]))

    @staticmethod
    def fm36(globalState):
        return (_dafny.Set({True})).intersection((_dafny.Set({True})) | (_dafny.Set({False, False, True})))

    @staticmethod
    def m0(p0, globalState):
        d_32_v0_: D2
        d_32_v0_ = D2_DC6(p0)
        d_33_v1_: D2
        d_33_v1_ = D2_DC8(d_32_v0_)
        d_34_v2_: D2
        d_34_v2_ = D2_DC8(d_33_v1_)
        d_35_v3_: D2
        d_35_v3_ = D2_DC8(d_33_v1_)
        def lambda0_(source1_):
            if source1_.is_DC7:
                d_36___mcc_h0_ = source1_.cf9
                d_37___mcc_h1_ = source1_.cf10
                d_38_cf10_ = d_37___mcc_h1_
                d_39_cf9_ = d_36___mcc_h0_
                return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False])]))).ispropersubset(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([not(False)])]))
            elif source1_.is_DC6:
                d_40___mcc_h2_ = source1_.cf8
                d_41_cf8_ = d_40___mcc_h2_
                return False
            elif True:
                d_42___mcc_h3_ = source1_.cf11
                d_43_cf11_ = d_42___mcc_h3_
                return False

        (globalState).f16 = lambda0_(d_35_v3_)
        d_44_v4_: _dafny.Seq
        d_44_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
        d_45_v5_: bool
        d_45_v5_ = True
        d_46_v6_: D3
        d_46_v6_ = D3_DC10(d_44_v4_, p0, d_45_v5_)
        d_47_v7_: _dafny.Seq
        d_47_v7_ = _dafny.SeqWithoutIsStrInference([D3_DC10(d_44_v4_, p0, d_45_v5_), d_46_v6_, d_46_v6_, d_46_v6_, d_46_v6_])
        d_48_i0_: int
        d_48_i0_ = 0
        with _dafny.label("0"):
            while (d_47_v7_) != (_dafny.SeqWithoutIsStrInference([D3_DC10(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_112_i1_ in range(default__.abs(-571))]), p0, d_45_v5_)])):
                with _dafny.c_label("0"):
                    if (d_48_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_48_i0_ = (d_48_i0_) + (1)
                    d_49_v8_: str
                    d_49_v8_ = _dafny.CodePoint('k')
                    d_50_v9_: _dafny.MultiSet
                    d_50_v9_ = _dafny.MultiSet([d_49_v8_])
                    d_51_v10_: D14
                    d_51_v10_ = D14_DC40(d_45_v5_, d_45_v5_, (d_50_v9_).cardinality)
                    d_52_v11_: D14
                    d_52_v11_ = D14_DC42(d_51_v10_)
                    d_53_v12_: D14
                    d_53_v12_ = D14_DC42(d_52_v11_)
                    source2_ = d_53_v12_
                    if source2_.is_DC40:
                        d_54___mcc_h4_ = source2_.cf69
                        d_55___mcc_h5_ = source2_.cf70
                        d_56___mcc_h6_ = source2_.cf71
                        d_57_cf71_ = d_56___mcc_h6_
                        d_58_cf70_ = d_55___mcc_h5_
                        d_59_cf69_ = d_54___mcc_h4_
                        d_49_v8_ = d_49_v8_
                        d_60_v13_: C0
                        nw0_ = C0()
                        nw0_.ctor__(True)
                        d_60_v13_ = nw0_
                        d_61_v14_: _dafny.Array
                        nw1_ = _dafny.Array(int(0), 3)
                        d_61_v14_ = nw1_
                        index0_ = default__.safeIndex(837, (d_61_v14_).length(0))
                        (d_61_v14_)[index0_] = p0
                        d_62_v15_: _dafny.MultiSet
                        d_62_v15_ = _dafny.MultiSet([p0])
                        index1_ = default__.safeIndex(837, (d_61_v14_).length(0))
                        (d_61_v14_)[index1_] = (((0) - ((d_62_v15_).cardinality)) - (p0)) + ((0) - (p0))
                        d_63_v16_: _dafny.Map
                        d_63_v16_ = _dafny.Map({not(d_58_cf70_): d_45_v5_})
                        (globalState).f28 = (d_57_cf71_) <= ((0) - (len(d_63_v16_)))
                    elif source2_.is_DC41:
                        d_64___mcc_h7_ = source2_.cf72
                        d_65___mcc_h8_ = source2_.cf73
                        d_66_cf73_ = d_65___mcc_h8_
                        d_67_cf72_ = d_64___mcc_h7_
                        d_68_v17_: _dafny.Array
                        def lambda1_(d_69_v4_):
                            def lambda2_(d_70_i2_):
                                return d_69_v4_

                            return lambda2_

                        init0_ = lambda1_(d_44_v4_)
                        nw2_ = _dafny.Array(None, 8)
                        for i0_0_ in range(nw2_.length(0)):
                            nw2_[i0_0_] = init0_(i0_0_)
                        d_68_v17_ = nw2_
                        d_71_v18_: _dafny.MultiSet
                        d_71_v18_ = _dafny.MultiSet([d_44_v4_])
                        d_72_v19_: _dafny.Map
                        d_72_v19_ = _dafny.Map({d_68_v17_: d_71_v18_})
                        d_72_v19_ = (d_72_v19_).set(d_68_v17_, (d_71_v18_).set(d_44_v4_, default__.abs(-699)))
                        d_66_cf73_ = 312
                        d_73_v20_: _dafny.Array
                        nw3_ = _dafny.Array(int(0), 12)
                        d_73_v20_ = nw3_
                        index2_ = default__.safeIndex(636, (d_73_v20_).length(0))
                        (d_73_v20_)[index2_] = -670
                        index3_ = default__.safeIndex(636, (d_73_v20_).length(0))
                        (d_73_v20_)[index3_] = d_66_cf73_
                        d_74_v21_: _dafny.Map
                        d_74_v21_ = _dafny.Map({d_45_v5_: not (d_45_v5_) or (d_45_v5_)})
                        d_74_v21_ = (d_74_v21_).set(d_45_v5_, d_45_v5_)
                    elif source2_.is_DC39:
                        d_75___mcc_h9_ = source2_.cf68
                        d_76_cf68_ = d_75___mcc_h9_
                        d_44_v4_ = (d_44_v4_) + (d_44_v4_)
                        d_77_v22_: _dafny.Map
                        d_77_v22_ = _dafny.Map({p0: d_45_v5_})
                        d_78_v23_: _dafny.Map
                        d_78_v23_ = _dafny.Map({p0: d_77_v22_})
                        d_79_v24_: _dafny.Map
                        d_79_v24_ = _dafny.Map({p0: p0})
                        d_80_v25_: _dafny.Seq
                        d_80_v25_ = _dafny.SeqWithoutIsStrInference([d_44_v4_])
                        d_81_v26_: C3
                        nw4_ = C3()
                        nw4_.ctor__(d_76_cf68_, d_78_v23_, default__.fm26(d_45_v5_, d_45_v5_, d_45_v5_, ((d_79_v24_)[p0] if (p0) in (d_79_v24_) else len((d_80_v25_)[default__.safeIndex(p0, len(d_80_v25_))])), globalState))
                        d_81_v26_ = nw4_
                        d_82_v27_: _dafny.Set
                        d_82_v27_ = _dafny.Set({p0})
                        d_83_v28_: _dafny.Map
                        d_83_v28_ = _dafny.Map({d_45_v5_: p0})
                        d_84_v29_: _dafny.Map
                        d_84_v29_ = _dafny.Map({d_81_v26_: (len(d_82_v27_)) - (len(d_83_v28_))})
                        d_84_v29_ = (d_84_v29_).set(d_81_v26_, p0)
                        d_85_v30_: D0
                        d_85_v30_ = D0_DC1(d_49_v8_, d_45_v5_, d_45_v5_)
                        d_85_v30_ = d_85_v30_
                        d_86_v31_: _dafny.Array
                        nw5_ = _dafny.Array(None, 1)
                        nw5_[int(0)] = d_45_v5_
                        d_86_v31_ = nw5_
                        d_87_v32_: _dafny.Map
                        d_87_v32_ = _dafny.Map({d_86_v31_: p0})
                        d_88_v33_: _dafny.Map
                        d_88_v33_ = _dafny.Map({d_86_v31_: (len(d_87_v32_)) > (p0)})
                        d_88_v33_ = (d_88_v33_).set(d_86_v31_, (d_83_v28_) in (default__.fm35(p0, d_44_v4_, (0) - (p0), globalState)))
                    elif True:
                        d_89___mcc_h10_ = source2_.cf74
                        d_90_cf74_ = d_89___mcc_h10_
                        d_91_v34_: _dafny.Set
                        d_91_v34_ = _dafny.Set({d_45_v5_, d_45_v5_, d_45_v5_, False, d_45_v5_})
                        d_92_v35_: _dafny.Map
                        d_92_v35_ = _dafny.Map({len(d_91_v34_): d_45_v5_})
                        d_93_v36_: _dafny.Seq
                        d_93_v36_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0, p0])
                        d_94_v37_: D16
                        d_94_v37_ = D16_DC44(d_93_v36_)
                        d_95_v38_: _dafny.Array
                        nw6_ = _dafny.Array(None, 23)
                        nw6_[int(0)] = -58
                        nw6_[int(1)] = p0
                        nw6_[int(2)] = len(default__.fm36(globalState))
                        nw6_[int(3)] = p0
                        nw6_[int(4)] = p0
                        nw6_[int(5)] = p0
                        nw6_[int(6)] = len(d_92_v35_)
                        nw6_[int(7)] = p0
                        nw6_[int(8)] = len((d_94_v37_).cf76)
                        nw6_[int(9)] = p0
                        nw6_[int(10)] = len(d_44_v4_)
                        nw6_[int(11)] = p0
                        nw6_[int(12)] = p0
                        nw6_[int(13)] = p0
                        nw6_[int(14)] = len(d_44_v4_)
                        nw6_[int(15)] = p0
                        nw6_[int(16)] = p0
                        nw6_[int(17)] = len(d_93_v36_)
                        nw6_[int(18)] = p0
                        nw6_[int(19)] = p0
                        nw6_[int(20)] = p0
                        nw6_[int(21)] = p0
                        nw6_[int(22)] = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_96_i3_ in range(default__.abs(635))])))
                        d_95_v38_ = nw6_
                        d_97_v39_: _dafny.Array
                        d_97_v39_ = d_95_v38_
                        (globalState).f24 = (d_97_v39_)
                        d_98_v40_: _dafny.Map
                        d_98_v40_ = _dafny.Map({(d_93_v36_) + (d_93_v36_): False})
                        d_98_v40_ = (d_98_v40_).set(d_93_v36_, (29) > (p0))
                        d_99_v41_: T0
                        nw7_ = C0()
                        nw7_.ctor__(not(d_45_v5_))
                        d_99_v41_ = nw7_
                        d_100_v42_: _dafny.Map
                        d_100_v42_ = _dafny.Map({d_45_v5_: d_99_v41_})
                        d_101_v43_: _dafny.Set
                        d_101_v43_ = _dafny.Set({p0})
                        d_102_v44_: D12
                        d_102_v44_ = D12_DC32(True, (d_99_v41_).f31, d_99_v41_, d_101_v43_, p0)
                        d_103_v45_: _dafny.Map
                        d_103_v45_ = _dafny.Map({p0: d_99_v41_})
                        d_104_v46_: _dafny.Array
                        nw8_ = _dafny.Array(None, 26)
                        nw8_[int(0)] = d_99_v41_
                        nw8_[int(1)] = d_99_v41_
                        nw8_[int(2)] = d_99_v41_
                        nw8_[int(3)] = d_99_v41_
                        nw8_[int(4)] = d_99_v41_
                        nw8_[int(5)] = d_99_v41_
                        nw8_[int(6)] = d_99_v41_
                        nw8_[int(7)] = ((d_100_v42_)[(d_99_v41_).f31] if ((d_99_v41_).f31) in (d_100_v42_) else d_99_v41_)
                        nw8_[int(8)] = d_99_v41_
                        nw8_[int(9)] = d_99_v41_
                        nw8_[int(10)] = d_99_v41_
                        nw8_[int(11)] = d_99_v41_
                        nw8_[int(12)] = d_99_v41_
                        nw8_[int(13)] = d_99_v41_
                        nw8_[int(14)] = d_99_v41_
                        nw8_[int(15)] = d_99_v41_
                        nw8_[int(16)] = d_99_v41_
                        nw8_[int(17)] = d_99_v41_
                        nw8_[int(18)] = d_99_v41_
                        nw8_[int(19)] = d_99_v41_
                        nw8_[int(20)] = (d_102_v44_).cf57
                        nw8_[int(21)] = d_99_v41_
                        nw8_[int(22)] = d_99_v41_
                        nw8_[int(23)] = ((d_103_v45_)[p0] if (p0) in (d_103_v45_) else d_99_v41_)
                        nw8_[int(24)] = d_99_v41_
                        nw8_[int(25)] = d_99_v41_
                        d_104_v46_ = nw8_
                        index4_ = default__.safeIndex(10, (d_104_v46_).length(0))
                        (d_104_v46_)[index4_] = d_99_v41_
                        index5_ = default__.safeIndex(10, (d_104_v46_).length(0))
                        (d_104_v46_)[index5_] = (d_99_v41_ if not((d_99_v41_).f31) else d_99_v41_)
                        d_105_v47_: int
                        d_105_v47_ = -327
                        d_105_v47_ = 774
                    d_106_v48_: int
                    d_106_v48_ = 662
                    d_106_v48_ = d_106_v48_
                    d_107_v49_: C5
                    nw9_ = C5()
                    nw9_.ctor__(not(not(d_45_v5_)), (d_106_v48_) > (p0))
                    d_107_v49_ = nw9_
                    d_108_v50_: _dafny.Map
                    d_108_v50_ = _dafny.Map({d_45_v5_: d_106_v48_})
                    d_109_v51_: _dafny.Seq
                    d_109_v51_ = _dafny.SeqWithoutIsStrInference([d_108_v50_])
                    d_110_v52_: _dafny.MultiSet
                    d_110_v52_ = _dafny.MultiSet([d_108_v50_])
                    d_111_v53_: C5
                    nw10_ = C5()
                    nw10_.ctor__((d_107_v49_).f30, (d_110_v52_).issubset(_dafny.MultiSet(d_109_v51_)))
                    d_111_v53_ = nw10_
                    pass
            pass
        d_113_v54_: _dafny.Array
        nw11_ = _dafny.Array(int(0), 29)
        d_113_v54_ = nw11_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_113_v54_).length(0)):
            d_114_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_114_i4_)) and ((d_114_i4_) < ((d_113_v54_).length(0)))):
                (d_113_v54_)[(d_114_i4_)] = (d_114_i4_) - (57)
        d_115_v55_: _dafny.Array
        def lambda3_(d_116_i5_):
            return _dafny.CodePoint('r')

        init1_ = lambda3_
        nw12_ = _dafny.Array(None, 8)
        for i0_1_ in range(nw12_.length(0)):
            nw12_[i0_1_] = init1_(i0_1_)
        d_115_v55_ = nw12_
        d_117_v56_: str
        d_117_v56_ = _dafny.CodePoint('m')
        index6_ = default__.safeIndex(887, (d_115_v55_).length(0))
        (d_115_v55_)[index6_] = d_117_v56_
        d_118_v57_: _dafny.MultiSet
        d_118_v57_ = _dafny.MultiSet([d_45_v5_])
        index7_ = default__.safeIndex(887, (d_115_v55_).length(0))
        (d_115_v55_)[index7_] = default__.fm5(d_45_v5_, ((d_118_v57_).cardinality) * (479), globalState)
        (globalState).f28 = (d_44_v4_) != (d_44_v4_)
        (globalState).f16 = ((d_45_v5_ if d_45_v5_ else d_45_v5_)) == (d_45_v5_)

    @staticmethod
    def Main(noArgsParameter__):
        d_119_v0_: str
        d_119_v0_ = _dafny.CodePoint('v')
        d_120_v1_: _dafny.Array
        nw13_ = _dafny.Array(None, 3)
        nw13_[int(0)] = d_119_v0_
        nw13_[int(1)] = d_119_v0_
        nw13_[int(2)] = d_119_v0_
        d_120_v1_ = nw13_
        d_121_v2_: _dafny.Seq
        d_121_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mau"))
        d_122_v3_: bool
        d_122_v3_ = False
        d_123_v4_: _dafny.Map
        d_123_v4_ = _dafny.Map({d_121_v2_: d_122_v3_})
        d_124_v5_: _dafny.Array
        nw14_ = _dafny.Array(int(0), 27)
        d_124_v5_ = nw14_
        d_125_v6_: _dafny.Map
        d_125_v6_ = _dafny.Map({-850: d_124_v5_})
        d_126_v7_: int
        d_126_v7_ = 159
        d_127_v8_: _dafny.Seq
        d_127_v8_ = _dafny.SeqWithoutIsStrInference([((d_125_v6_)[d_126_v7_] if (d_126_v7_) in (d_125_v6_) else d_124_v5_)])
        d_128_v9_: _dafny.Map
        d_128_v9_ = _dafny.Map({d_122_v3_: d_126_v7_})
        d_129_v10_: _dafny.Seq
        d_129_v10_ = _dafny.SeqWithoutIsStrInference([not(d_122_v3_), d_122_v3_])
        d_130_v11_: _dafny.Map
        d_130_v11_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_126_v7_, ((d_128_v9_)[d_122_v3_] if (d_122_v3_) in (d_128_v9_) else len(d_128_v9_))]): (_dafny.MultiSet(d_129_v10_)).cardinality})
        d_131_v12_: _dafny.Array
        def lambda4_(d_132_v10_):
            def lambda5_(d_133_i0_):
                return d_132_v10_

            return lambda5_

        init2_ = lambda4_(d_129_v10_)
        nw15_ = _dafny.Array(None, 6)
        for i0_2_ in range(nw15_.length(0)):
            nw15_[i0_2_] = init2_(i0_2_)
        d_131_v12_ = nw15_
        d_134_v13_: D0
        d_134_v13_ = D0_DC0(d_131_v12_)
        d_135_v14_: _dafny.Array
        nw16_ = _dafny.Array(None, 12)
        nw16_[int(0)] = d_122_v3_
        nw16_[int(1)] = not(d_122_v3_)
        nw16_[int(2)] = False
        nw16_[int(3)] = d_122_v3_
        nw16_[int(4)] = d_122_v3_
        nw16_[int(5)] = d_122_v3_
        nw16_[int(6)] = False
        nw16_[int(7)] = False
        nw16_[int(8)] = d_122_v3_
        nw16_[int(9)] = d_122_v3_
        nw16_[int(10)] = d_122_v3_
        nw16_[int(11)] = d_122_v3_
        d_135_v14_ = nw16_
        d_136_v15_: _dafny.Array
        d_136_v15_ = d_124_v5_
        d_137_v16_: _dafny.Map
        d_137_v16_ = _dafny.Map({d_122_v3_: _dafny.Set({d_122_v3_, d_122_v3_, d_122_v3_, False, d_122_v3_})})
        d_138_globalState_: GlobalState
        nw17_ = GlobalState()
        nw17_.ctor__(d_120_v1_, False, False, d_123_v4_, False, True, d_121_v2_, 267, False, True, True, d_127_v8_, d_130_v11_, 175, True, 87, False, False, (d_134_v13_).cf0, True, d_135_v14_, (d_136_v15_), False, (d_137_v16_) | (d_137_v16_), d_124_v5_, 21, -144, d_135_v14_, True)
        d_138_globalState_ = nw17_
        d_139_v17_: D0
        d_139_v17_ = D0_DC3(d_119_v0_)
        d_140_v18_: _dafny.Set
        d_140_v18_ = _dafny.Set({d_139_v17_})
        d_141_v19_: _dafny.Seq
        d_141_v19_ = _dafny.SeqWithoutIsStrInference([d_126_v7_, d_126_v7_, (d_126_v7_) * (len(d_140_v18_)), d_126_v7_])
        d_126_v7_ = (d_141_v19_)[default__.safeIndex(d_126_v7_, len(d_141_v19_))]
        hi0_ = (d_126_v7_) * (len(d_121_v2_))
        for d_142_i1_ in range(default__.safeModuloInt(d_126_v7_, d_126_v7_), hi0_):
            d_143_v20_: _dafny.Set
            d_143_v20_ = _dafny.Set({d_122_v3_, True})
            d_144_v21_: _dafny.Map
            d_144_v21_ = _dafny.Map({d_126_v7_: len(d_143_v20_)})
            default__.m0(((d_144_v21_)[len(_dafny.Map({d_135_v14_: d_142_i1_}))] if (len(_dafny.Map({d_135_v14_: d_142_i1_}))) in (d_144_v21_) else d_126_v7_), d_138_globalState_)
            d_145_v22_: D0
            d_145_v22_ = D0_DC2(d_122_v3_)
            pat_let_tv0_ = d_129_v10_
            pat_let_tv1_ = d_129_v10_
            def iife10_(_pat_let0_0):
                def iife11_(d_146_dt__update__tmp_h0_):
                    def iife12_(_pat_let1_0):
                        def iife13_(d_147_dt__update_hcf4_h0_):
                            return D0_DC2(d_147_dt__update_hcf4_h0_)
                        return iife13_(_pat_let1_0)
                    return iife12_((pat_let_tv0_) < (pat_let_tv1_))
                return iife11_(_pat_let0_0)
            rhs0_ = iife10_(d_145_v22_)
            rhs1_ = d_122_v3_
            rhs2_ = d_135_v14_
            lhs0_ = d_138_globalState_
            lhs1_ = d_138_globalState_
            d_145_v22_ = rhs0_
            lhs0_.f16 = rhs1_
            lhs1_.f20 = rhs2_
            source3_ = (d_136_v15_ if d_122_v3_ else d_124_v5_)
            d_148___mcc_h0_ = source3_
            d_149_cf7_ = d_148___mcc_h0_
            d_122_v3_ = ((len(d_128_v9_)) * (d_142_i1_)) == (d_126_v7_)
            d_134_v13_ = d_134_v13_
            index8_ = default__.safeIndex(697, (d_149_cf7_).length(0))
            (d_149_cf7_)[index8_] = d_126_v7_
            index9_ = default__.safeIndex(697, (d_149_cf7_).length(0))
            (d_149_cf7_)[index9_] = d_126_v7_
            d_122_v3_ = not (d_122_v3_) or (d_122_v3_)
            d_150_v23_: D2
            d_150_v23_ = D2_DC6(d_126_v7_)
            (d_138_globalState_).f28 = ((d_127_v8_).set(default__.safeIndex(len(((d_121_v2_).set(default__.safeIndex(d_126_v7_, len(d_121_v2_)), _dafny.CodePoint('q'))).set(default__.safeIndex((d_150_v23_).cf8, len((d_121_v2_).set(default__.safeIndex(d_126_v7_, len(d_121_v2_)), _dafny.CodePoint('q')))), d_119_v0_)), len(d_127_v8_)), d_124_v5_)) <= ((d_127_v8_) + (d_127_v8_))
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_135_v14_).length(0)):
            d_151_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_151_i2_)) and ((d_151_i2_) < ((d_135_v14_).length(0)))):
                (d_135_v14_)[(d_151_i2_)] = False
        d_152_v24_: _dafny.Map
        d_152_v24_ = _dafny.Map({d_126_v7_: d_126_v7_})
        d_153_v25_: _dafny.Map
        d_153_v25_ = _dafny.Map({d_126_v7_: ((d_152_v24_)[d_126_v7_] if (d_126_v7_) in (d_152_v24_) else d_126_v7_)})
        (d_138_globalState_).f16 = not((d_153_v25_) == (d_153_v25_))
        if False:
            d_154_v26_: _dafny.Map
            d_154_v26_ = _dafny.Map({d_122_v3_: d_122_v3_})
            d_126_v7_ = len(d_154_v26_)
            index10_ = default__.safeIndex(86, (d_124_v5_).length(0))
            (d_124_v5_)[index10_] = d_126_v7_
            index11_ = default__.safeIndex(86, (d_124_v5_).length(0))
            rhs3_ = not (d_122_v3_) or (d_122_v3_)
            rhs4_ = (0) - (d_126_v7_)
            rhs5_ = d_129_v10_
            rhs6_ = (0) - ((0) - (d_126_v7_))
            lhs2_ = d_124_v5_
            lhs3_ = default__.safeIndex(86, (d_124_v5_).length(0))
            d_122_v3_ = rhs3_
            d_126_v7_ = rhs4_
            d_129_v10_ = rhs5_
            lhs2_[lhs3_] = rhs6_
            d_155_v27_: _dafny.Map
            d_155_v27_ = _dafny.Map({(d_126_v7_) >= ((d_124_v5_)[default__.safeIndex(86, (d_124_v5_).length(0))]): _dafny.CodePoint('u')})
            d_155_v27_ = (d_155_v27_).set(d_122_v3_, d_119_v0_)
            d_126_v7_ = (default__.safeDivisionInt(len(default__.fm0(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, not(d_122_v3_)])), d_138_globalState_)), (_dafny.MultiSet([True])).cardinality)) + ((-660 if d_122_v3_ else d_126_v7_))
            d_156_v28_: _dafny.Array
            nw18_ = _dafny.Array(_dafny.Array(None, 0), 13)
            d_156_v28_ = nw18_
            d_157_v29_: _dafny.Array
            nw19_ = _dafny.Array(None, 7)
            nw19_[int(0)] = d_141_v19_
            nw19_[int(1)] = d_141_v19_
            nw19_[int(2)] = d_141_v19_
            nw19_[int(3)] = d_141_v19_
            nw19_[int(4)] = d_141_v19_
            nw19_[int(5)] = d_141_v19_
            nw19_[int(6)] = _dafny.SeqWithoutIsStrInference([d_126_v7_, (d_124_v5_)[default__.safeIndex(86, (d_124_v5_).length(0))]])
            d_157_v29_ = nw19_
            index12_ = default__.safeIndex(363, (d_156_v28_).length(0))
            (d_156_v28_)[index12_] = d_157_v29_
            index13_ = default__.safeIndex(363, (d_156_v28_).length(0))
            (d_156_v28_)[index13_] = d_157_v29_
        elif True:
            d_122_v3_ = d_122_v3_
            d_158_v30_: _dafny.Set
            d_158_v30_ = _dafny.Set({d_141_v19_})
            d_159_v31_: _dafny.Set
            d_159_v31_ = _dafny.Set({len(d_158_v30_), d_126_v7_})
            if not((d_159_v31_).ispropersubset(_dafny.Set({d_126_v7_}))):
                d_121_v2_ = d_121_v2_
                pat_let_tv2_ = d_119_v0_
                d_160_v32_: _dafny.Array
                nw20_ = _dafny.Array(None, 26)
                nw20_[int(0)] = D0_DC3(d_119_v0_)
                nw20_[int(1)] = d_139_v17_
                def iife14_(_pat_let2_0):
                    def iife15_(d_161_dt__update__tmp_h1_):
                        def iife16_(_pat_let3_0):
                            def iife17_(d_162_dt__update_hcf5_h0_):
                                return D0_DC3(d_162_dt__update_hcf5_h0_)
                            return iife17_(_pat_let3_0)
                        return iife16_(pat_let_tv2_)
                    return iife15_(_pat_let2_0)
                nw20_[int(2)] = iife14_(d_139_v17_)
                nw20_[int(3)] = d_139_v17_
                nw20_[int(4)] = (d_139_v17_ if d_122_v3_ else D0_DC3(_dafny.CodePoint('n')))
                nw20_[int(5)] = d_139_v17_
                nw20_[int(6)] = d_139_v17_
                nw20_[int(7)] = d_139_v17_
                nw20_[int(8)] = D0_DC3(_dafny.CodePoint('p'))
                nw20_[int(9)] = d_139_v17_
                nw20_[int(10)] = d_139_v17_
                nw20_[int(11)] = d_139_v17_
                nw20_[int(12)] = d_139_v17_
                nw20_[int(13)] = d_139_v17_
                nw20_[int(14)] = d_139_v17_
                nw20_[int(15)] = d_139_v17_
                nw20_[int(16)] = d_139_v17_
                nw20_[int(17)] = d_139_v17_
                nw20_[int(18)] = d_139_v17_
                nw20_[int(19)] = d_139_v17_
                nw20_[int(20)] = D0_DC3(d_119_v0_)
                nw20_[int(21)] = d_139_v17_
                nw20_[int(22)] = d_139_v17_
                nw20_[int(23)] = d_139_v17_
                nw20_[int(24)] = d_139_v17_
                nw20_[int(25)] = d_139_v17_
                d_160_v32_ = nw20_
                d_160_v32_ = d_160_v32_
                d_163_v33_: _dafny.MultiSet
                d_163_v33_ = _dafny.MultiSet([d_126_v7_, (d_126_v7_) - (d_126_v7_)])
                rhs7_ = d_122_v3_
                rhs8_ = not (d_122_v3_) or (d_122_v3_)
                rhs9_ = _dafny.MultiSet([-665])
                lhs4_ = d_138_globalState_
                lhs5_ = d_138_globalState_
                lhs4_.f16 = rhs7_
                lhs5_.f28 = rhs8_
                d_163_v33_ = rhs9_
                d_126_v7_ = d_126_v7_
                d_164_v34_: _dafny.Map
                d_164_v34_ = _dafny.Map({d_122_v3_: _dafny.CodePoint('g')})
                d_165_v35_: _dafny.MultiSet
                d_165_v35_ = _dafny.MultiSet([D0_DC1(((d_164_v34_)[d_122_v3_] if (d_122_v3_) in (d_164_v34_) else d_119_v0_), d_122_v3_, False)])
                d_166_v36_: _dafny.Seq
                d_166_v36_ = _dafny.SeqWithoutIsStrInference([d_165_v35_])
                d_167_v37_: D0
                d_167_v37_ = D0_DC1(d_119_v0_, d_122_v3_, d_122_v3_)
                index14_ = default__.safeIndex(33, (d_135_v14_).length(0))
                (d_135_v14_)[index14_] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_167_v37_]))).issubset((d_166_v36_)[default__.safeIndex(d_126_v7_, len(d_166_v36_))])
                index15_ = default__.safeIndex(33, (d_135_v14_).length(0))
                (d_135_v14_)[index15_] = (d_126_v7_) >= (default__.safeDivisionInt(d_126_v7_, d_126_v7_))
            elif True:
                d_168_v38_: C2
                nw21_ = C2()
                nw21_.ctor__(d_129_v10_, True)
                d_168_v38_ = nw21_
                d_169_v39_: _dafny.Map
                d_169_v39_ = _dafny.Map({d_122_v3_: d_122_v3_})
                d_170_v40_: _dafny.Array
                nw22_ = _dafny.Array(None, 26)
                nw22_[int(0)] = d_169_v39_
                nw22_[int(1)] = d_169_v39_
                nw22_[int(2)] = _dafny.Map({not(d_122_v3_): d_122_v3_})
                nw22_[int(3)] = default__.fm30(d_119_v0_, d_139_v17_, d_138_globalState_)
                nw22_[int(4)] = _dafny.Map({d_122_v3_: not(d_122_v3_)})
                nw22_[int(5)] = d_169_v39_
                nw22_[int(6)] = _dafny.Map({d_122_v3_: d_122_v3_})
                nw22_[int(7)] = d_169_v39_
                nw22_[int(8)] = d_169_v39_
                nw22_[int(9)] = d_169_v39_
                nw22_[int(10)] = d_169_v39_
                nw22_[int(11)] = _dafny.Map({d_122_v3_: d_122_v3_})
                nw22_[int(12)] = (d_169_v39_).set(d_122_v3_, d_122_v3_)
                nw22_[int(13)] = d_169_v39_
                nw22_[int(14)] = d_169_v39_
                nw22_[int(15)] = d_169_v39_
                nw22_[int(16)] = d_169_v39_
                nw22_[int(17)] = d_169_v39_
                nw22_[int(18)] = d_169_v39_
                nw22_[int(19)] = d_169_v39_
                nw22_[int(20)] = d_169_v39_
                nw22_[int(21)] = d_169_v39_
                nw22_[int(22)] = d_169_v39_
                nw22_[int(23)] = d_169_v39_
                nw22_[int(24)] = d_169_v39_
                nw22_[int(25)] = d_169_v39_
                d_170_v40_ = nw22_
                d_171_v41_: _dafny.Map
                d_171_v41_ = _dafny.Map({d_126_v7_: d_122_v3_})
                d_172_v42_: T0
                nw23_ = C3()
                nw23_.ctor__(d_170_v40_, _dafny.Map({-190: d_171_v41_}), d_122_v3_)
                d_172_v42_ = nw23_
                d_172_v42_ = d_172_v42_
                index16_ = default__.safeIndex(507, (d_124_v5_).length(0))
                (d_124_v5_)[index16_] = (d_126_v7_ if d_122_v3_ else len(d_121_v2_))
                index17_ = default__.safeIndex(507, (d_124_v5_).length(0))
                (d_124_v5_)[index17_] = (d_168_v38_).fm13(d_126_v7_, d_138_globalState_)
                (d_138_globalState_).f28 = False
                d_122_v3_ = (d_172_v42_).f31
            d_173_v43_: _dafny.Seq
            d_173_v43_ = _dafny.SeqWithoutIsStrInference([d_129_v10_])
            d_174_v44_: _dafny.Seq
            d_174_v44_ = _dafny.SeqWithoutIsStrInference([d_129_v10_, (d_173_v43_)[default__.safeIndex(d_126_v7_, len(d_173_v43_))], d_129_v10_, _dafny.SeqWithoutIsStrInference([d_122_v3_])])
            d_126_v7_ = default__.safeModuloInt(len(d_174_v44_), d_126_v7_)
            if d_122_v3_:
                d_128_v9_ = (d_128_v9_).set((d_122_v3_) and (d_122_v3_), 62)
                default__.m0(d_126_v7_, d_138_globalState_)
                d_126_v7_ = default__.safeDivisionInt(d_126_v7_, d_126_v7_)
                default__.m0(d_126_v7_, d_138_globalState_)
                d_175_v45_: _dafny.Array
                nw24_ = _dafny.Array(None, 1)
                nw24_[int(0)] = (d_121_v2_).set(default__.safeIndex(d_126_v7_, len(d_121_v2_)), (d_121_v2_)[default__.safeIndex(len(d_121_v2_), len(d_121_v2_))])
                d_175_v45_ = nw24_
                d_175_v45_ = d_175_v45_
            elif True:
                d_126_v7_ = d_126_v7_
                d_139_v17_ = default__.fm31(d_126_v7_, d_138_globalState_)
                d_176_v46_: D5
                d_176_v46_ = D5_DC14(d_159_v31_, d_126_v7_, True, d_135_v14_)
                d_177_v47_: _dafny.Array
                nw25_ = _dafny.Array(None, 7)
                nw25_[int(0)] = d_135_v14_
                nw25_[int(1)] = d_135_v14_
                nw25_[int(2)] = (d_176_v46_).cf22
                nw25_[int(3)] = d_135_v14_
                nw25_[int(4)] = d_135_v14_
                nw25_[int(5)] = d_135_v14_
                nw25_[int(6)] = d_135_v14_
                d_177_v47_ = nw25_
                d_178_v48_: C1
                nw26_ = C1()
                nw26_.ctor__(d_126_v7_, d_177_v47_, d_122_v3_)
                d_178_v48_ = nw26_
                d_179_v49_: _dafny.Map
                d_179_v49_ = _dafny.Map({d_126_v7_: d_122_v3_})
                d_180_v50_: D9
                d_180_v50_ = D9_DC26(d_122_v3_, d_121_v2_, d_126_v7_, d_179_v49_, d_122_v3_)
                d_181_v51_: _dafny.Map
                d_181_v51_ = _dafny.Map({d_180_v50_: _dafny.MultiSet([-674, len(d_141_v19_)])})
                d_182_v52_: _dafny.MultiSet
                d_182_v52_ = _dafny.MultiSet([d_126_v7_])
                d_183_v53_: C5
                nw27_ = C5()
                nw27_.ctor__((default__.fm20(d_121_v2_, ((d_181_v51_)[d_180_v50_] if (d_180_v50_) in (d_181_v51_) else d_182_v52_), not(not((d_129_v10_)[default__.safeIndex(d_126_v7_, len(d_129_v10_))])), (d_182_v52_).cardinality, d_138_globalState_)) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cf"))), d_122_v3_)
                d_183_v53_ = nw27_
                d_126_v7_ = len((_dafny.Map({(d_183_v53_).f29: d_119_v0_})) | (_dafny.Map({d_122_v3_: d_119_v0_})))
            d_184_v54_: _dafny.MultiSet
            d_184_v54_ = _dafny.MultiSet([d_126_v7_])
            index18_ = default__.safeIndex(792, (d_135_v14_).length(0))
            (d_135_v14_)[index18_] = (d_126_v7_) in (d_184_v54_)
            index19_ = default__.safeIndex(858, (d_135_v14_).length(0))
            (d_135_v14_)[index19_] = d_122_v3_
            d_185_v55_: C5
            nw28_ = C5()
            nw28_.ctor__(d_122_v3_, d_122_v3_)
            d_185_v55_ = nw28_
            d_186_v56_: _dafny.Map
            d_186_v56_ = _dafny.Map({d_122_v3_: d_185_v55_})
            index20_ = default__.safeIndex(792, (d_135_v14_).length(0))
            index21_ = default__.safeIndex(858, (d_135_v14_).length(0))
            rhs10_ = ((d_126_v7_) - (len(d_121_v2_))) < (default__.safeModuloInt(d_126_v7_, d_126_v7_))
            rhs11_ = len(((d_186_v56_) | ((d_186_v56_).set((d_185_v55_).f30, d_185_v55_))) | ((_dafny.Map({default__.fm10(d_119_v0_, (d_185_v55_).f30, d_138_globalState_): d_185_v55_})) | (d_186_v56_)))
            rhs12_ = (d_185_v55_).f29
            rhs13_ = d_135_v14_
            lhs6_ = d_135_v14_
            lhs7_ = default__.safeIndex(792, (d_135_v14_).length(0))
            lhs8_ = d_135_v14_
            lhs9_ = default__.safeIndex(858, (d_135_v14_).length(0))
            lhs10_ = d_138_globalState_
            lhs6_[lhs7_] = rhs10_
            d_126_v7_ = rhs11_
            lhs8_[lhs9_] = rhs12_
            lhs10_.f20 = rhs13_
        default__.m0(default__.safeDivisionInt(d_126_v7_, d_126_v7_), d_138_globalState_)
        d_187_i3_: int
        d_187_i3_ = 0
        with _dafny.label("1"):
            while d_122_v3_:
                with _dafny.c_label("1"):
                    if (d_187_i3_) >= (100):
                        raise _dafny.Break("1")
                    d_187_i3_ = (d_187_i3_) + (1)
                    d_188_v57_: C0
                    nw29_ = C0()
                    nw29_.ctor__((d_126_v7_) < (d_126_v7_))
                    d_188_v57_ = nw29_
                    d_126_v7_ = d_126_v7_
                    d_126_v7_ = 796
                    default__.m0(d_126_v7_, d_138_globalState_)
                    pass
            pass
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_124_v5_).length(0)):
            d_189_i4_: int = guard_loop_2_
            if (True) and (((0) <= (d_189_i4_)) and ((d_189_i4_) < ((d_124_v5_).length(0)))):
                (d_124_v5_)[(d_189_i4_)] = (d_189_i4_) * (d_126_v7_)
        d_190_i5_: int
        d_190_i5_ = 0
        with _dafny.label("2"):
            while d_122_v3_:
                with _dafny.c_label("2"):
                    if (d_190_i5_) >= (100):
                        raise _dafny.Break("2")
                    d_190_i5_ = (d_190_i5_) + (1)
                    d_191_v58_: _dafny.MultiSet
                    d_191_v58_ = _dafny.MultiSet([49])
                    d_192_v59_: _dafny.Seq
                    d_192_v59_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_126_v7_]), d_191_v58_, d_191_v58_])
                    d_193_v60_: D12
                    d_193_v60_ = D12_DC31(d_129_v10_)
                    pat_let_tv3_ = d_122_v3_
                    def iife18_(_pat_let4_0):
                        def iife19_(d_194_dt__update__tmp_h2_):
                            def iife20_(_pat_let5_0):
                                def iife21_(d_195_dt__update_hcf54_h0_):
                                    return D12_DC31(d_195_dt__update_hcf54_h0_)
                                return iife21_(_pat_let5_0)
                            return iife20_(_dafny.SeqWithoutIsStrInference([not(pat_let_tv3_)]))
                        return iife19_(_pat_let4_0)
                    rhs14_ = (_dafny.MultiSet([d_126_v7_])).issubset((d_192_v59_)[default__.safeIndex(d_126_v7_, len(d_192_v59_))])
                    rhs15_ = d_124_v5_
                    rhs16_ = d_122_v3_
                    rhs17_ = (iife18_(d_193_v60_)).cf54
                    rhs18_ = d_122_v3_
                    lhs11_ = d_138_globalState_
                    lhs12_ = d_138_globalState_
                    lhs13_ = d_138_globalState_
                    lhs11_.f16 = rhs14_
                    lhs12_.f24 = rhs15_
                    d_122_v3_ = rhs16_
                    d_129_v10_ = rhs17_
                    lhs13_.f16 = rhs18_
                    d_126_v7_ = d_126_v7_
                    index22_ = default__.safeIndex(259, (d_135_v14_).length(0))
                    (d_135_v14_)[index22_] = (d_126_v7_) > (d_126_v7_)
                    index23_ = default__.safeIndex(259, (d_135_v14_).length(0))
                    rhs19_ = (d_129_v10_) + ((d_129_v10_) + (_dafny.SeqWithoutIsStrInference([d_122_v3_, True])))
                    rhs20_ = d_122_v3_
                    rhs21_ = d_122_v3_
                    lhs14_ = d_135_v14_
                    lhs15_ = default__.safeIndex(259, (d_135_v14_).length(0))
                    lhs16_ = d_138_globalState_
                    d_129_v10_ = rhs19_
                    lhs14_[lhs15_] = rhs20_
                    lhs16_.f28 = rhs21_
                    default__.m0(d_126_v7_, d_138_globalState_)
                    pass
            pass
        d_196_v61_: _dafny.MultiSet
        d_196_v61_ = _dafny.MultiSet([d_122_v3_])
        d_197_v62_: _dafny.Seq
        d_197_v62_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_122_v3_, d_122_v3_, d_122_v3_, d_122_v3_]), (d_196_v61_).set(d_122_v3_, default__.abs(d_126_v7_))])
        d_198_v63_: D5
        d_198_v63_ = D5_DC13((d_197_v62_)[default__.safeIndex(d_126_v7_, len(d_197_v62_))])
        source4_ = d_198_v63_
        if source4_.is_DC14:
            d_199___mcc_h1_ = source4_.cf19
            d_200___mcc_h2_ = source4_.cf20
            d_201___mcc_h3_ = source4_.cf21
            d_202___mcc_h4_ = source4_.cf22
            d_203_cf22_ = d_202___mcc_h4_
            d_204_cf21_ = d_201___mcc_h3_
            d_205_cf20_ = d_200___mcc_h2_
            d_206_cf19_ = d_199___mcc_h1_
            d_205_cf20_ = (len(d_141_v19_)) * (default__.fm7(d_204_cf21_, d_122_v3_, d_138_globalState_))
            d_207_v64_: C2
            nw30_ = C2()
            nw30_.ctor__((d_129_v10_).set(default__.safeIndex(703, len(d_129_v10_)), d_122_v3_), False)
            d_207_v64_ = nw30_
            d_207_v64_ = d_207_v64_
            d_208_v65_: D0
            d_208_v65_ = D0_DC1(d_119_v0_, not(d_204_cf21_), default__.fm26(d_122_v3_, d_122_v3_, d_204_cf21_, d_205_cf20_, d_138_globalState_))
            d_209_v66_: _dafny.Seq
            d_209_v66_ = _dafny.SeqWithoutIsStrInference([d_208_v65_])
            d_210_v67_: _dafny.Map
            d_210_v67_ = _dafny.Map({d_119_v0_: len(_dafny.Map({d_122_v3_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ks")))}))})
            d_211_v68_: _dafny.Map
            d_211_v68_ = _dafny.Map({d_126_v7_: d_119_v0_})
            d_212_v69_: _dafny.Array
            nw31_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_212_v69_ = nw31_
            d_213_v70_: C4
            nw32_ = C4()
            nw32_.ctor__(d_209_v66_, d_122_v3_, (0) - (((d_210_v67_)[((d_211_v68_)[len(d_121_v2_)] if (len(d_121_v2_)) in (d_211_v68_) else d_119_v0_)] if (((d_211_v68_)[len(d_121_v2_)] if (len(d_121_v2_)) in (d_211_v68_) else d_119_v0_)) in (d_210_v67_) else len(d_121_v2_))), d_212_v69_, d_122_v3_)
            d_213_v70_ = nw32_
            d_213_v70_ = d_213_v70_
            d_214_v71_: _dafny.Set
            d_214_v71_ = _dafny.Set({d_213_v70_.f35})
            (d_138_globalState_).f16 = not ((d_214_v71_).issubset(d_214_v71_)) or (d_204_cf21_)
        elif source4_.is_DC13:
            d_215___mcc_h5_ = source4_.cf18
            d_216_cf18_ = d_215___mcc_h5_
            d_124_v5_ = d_124_v5_
            d_217_v72_: _dafny.Array
            nw33_ = _dafny.Array(None, 5)
            d_217_v72_ = nw33_
            d_218_v73_: _dafny.Array
            nw34_ = _dafny.Array(_dafny.Array(None, 0), 22)
            d_218_v73_ = nw34_
            d_219_v74_: T1
            nw35_ = C1()
            nw35_.ctor__(d_126_v7_, d_218_v73_, d_122_v3_)
            d_219_v74_ = nw35_
            d_220_v75_: _dafny.Seq
            d_220_v75_ = _dafny.SeqWithoutIsStrInference([d_219_v74_, d_219_v74_])
            index24_ = default__.safeIndex(842, (d_217_v72_).length(0))
            (d_217_v72_)[index24_] = (d_220_v75_)[default__.safeIndex((d_141_v19_)[default__.safeIndex(d_126_v7_, len(d_141_v19_))], len(d_220_v75_))]
            index25_ = default__.safeIndex(842, (d_217_v72_).length(0))
            (d_217_v72_)[index25_] = d_219_v74_
            d_129_v10_ = d_129_v10_
            index26_ = default__.safeIndex(684, (d_218_v73_).length(0))
            (d_218_v73_)[index26_] = d_135_v14_
            index27_ = default__.safeIndex(684, (d_218_v73_).length(0))
            (d_218_v73_)[index27_] = d_135_v14_
        elif True:
            d_221___mcc_h6_ = source4_.cf23
            d_222_cf23_ = d_221___mcc_h6_
            d_223_v76_: D0
            d_223_v76_ = D0_DC2(d_122_v3_)
            d_224_v77_: D0
            d_224_v77_ = D0_DC1(d_119_v0_, (d_223_v76_).cf4, d_122_v3_)
            d_225_v78_: _dafny.Seq
            d_225_v78_ = _dafny.SeqWithoutIsStrInference([d_224_v77_])
            d_226_v79_: _dafny.Map
            d_226_v79_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_126_v7_]): _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(d_141_v19_)[default__.safeIndex(len(_dafny.Map({D0_DC2(d_122_v3_): d_126_v7_})), len(d_141_v19_))], len(_dafny.SeqWithoutIsStrInference([d_119_v0_ for d_227_i7_ in range(default__.abs(971))]))])) for d_228_i6_ in range(default__.abs(106))])})
            d_229_v80_: _dafny.Seq
            d_229_v80_ = _dafny.SeqWithoutIsStrInference([d_141_v19_])
            d_230_v81_: _dafny.Array
            nw36_ = _dafny.Array(_dafny.Array(None, 0), 25)
            d_230_v81_ = nw36_
            d_231_v82_: C4
            nw37_ = C4()
            nw37_.ctor__((d_225_v78_) + (d_225_v78_), d_122_v3_, len(((d_226_v79_)[(d_229_v80_)[default__.safeIndex(d_126_v7_, len(d_229_v80_))]] if ((d_229_v80_)[default__.safeIndex(d_126_v7_, len(d_229_v80_))]) in (d_226_v79_) else d_141_v19_)), d_230_v81_, not (False) or (d_122_v3_))
            d_231_v82_ = nw37_
            d_232_v83_: _dafny.Set
            d_232_v83_ = _dafny.Set({default__.fm1(d_119_v0_, d_138_globalState_)})
            d_233_v84_: _dafny.Map
            d_233_v84_ = _dafny.Map({d_122_v3_: d_231_v82_.f35})
            rhs22_ = default__.fm7(not (d_231_v82_.f35) or (d_122_v3_), ((d_233_v84_)[not(d_122_v3_)] if (not(d_122_v3_)) in (d_233_v84_) else d_231_v82_.f35), d_138_globalState_)
            rhs23_ = (d_231_v82_ if d_122_v3_ else d_231_v82_)
            rhs24_ = d_231_v82_.f35
            rhs25_ = d_232_v83_
            rhs26_ = (_dafny.SeqWithoutIsStrInference([d_119_v0_ for d_234_i8_ in range(default__.abs(674))]) if d_122_v3_ else d_121_v2_)
            lhs17_ = d_138_globalState_
            d_126_v7_ = rhs22_
            d_231_v82_ = rhs23_
            lhs17_.f28 = rhs24_
            d_232_v83_ = rhs25_
            d_121_v2_ = rhs26_
            index28_ = default__.safeIndex(209, (d_135_v14_).length(0))
            (d_135_v14_)[index28_] = (d_121_v2_) != ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))).set(default__.safeIndex(len(default__.fm32(d_126_v7_, d_231_v82_.f35, d_231_v82_.f35, d_231_v82_.f35, d_138_globalState_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))), d_119_v0_))
            index29_ = default__.safeIndex(209, (d_135_v14_).length(0))
            (d_135_v14_)[index29_] = d_231_v82_.f35
            d_235_v85_: int
            d_236_v86_: _dafny.Array
            d_237_v87_: _dafny.Seq
            d_238_v88_: int
            out0_: int
            out1_: _dafny.Array
            out2_: _dafny.Seq
            out3_: int
            out0_, out1_, out2_, out3_ = (d_231_v82_).m2((d_231_v82_.f35 if (d_135_v14_)[default__.safeIndex(209, (d_135_v14_).length(0))] else d_122_v3_), (d_196_v61_) | (d_196_v61_), d_196_v61_, d_138_globalState_)
            d_235_v85_ = out0_
            d_236_v86_ = out1_
            d_237_v87_ = out2_
            d_238_v88_ = out3_
            d_239_v89_: D3
            d_239_v89_ = D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eavwb")), d_235_v85_, False)
            d_240_v90_: D3
            d_240_v90_ = D3_DC11(d_239_v89_)
            (d_138_globalState_).f16 = (default__.fm33(d_238_v88_, d_240_v90_, d_121_v2_, d_138_globalState_)).cf32
        d_241_v91_: _dafny.Map
        d_241_v91_ = _dafny.Map({len(d_121_v2_): d_122_v3_})
        d_242_v92_: _dafny.Map
        d_242_v92_ = _dafny.Map({(d_241_v91_).set(d_126_v7_, d_122_v3_): d_126_v7_})
        d_243_v93_: D13
        d_243_v93_ = D13_DC34(d_242_v92_)
        d_244_v94_: _dafny.Map
        d_244_v94_ = _dafny.Map({d_122_v3_: d_135_v14_})
        d_245_v95_: _dafny.Seq
        d_245_v95_ = _dafny.SeqWithoutIsStrInference([((d_244_v94_)[d_122_v3_] if (d_122_v3_) in (d_244_v94_) else d_135_v14_)])
        d_246_v96_: _dafny.MultiSet
        d_246_v96_ = _dafny.MultiSet([len(d_245_v95_)])
        d_247_v97_: _dafny.Seq
        d_247_v97_ = _dafny.SeqWithoutIsStrInference([d_246_v96_])
        d_248_v98_: _dafny.MultiSet
        d_248_v98_ = _dafny.MultiSet([d_246_v96_, d_246_v96_, d_246_v96_])
        (d_138_globalState_).f16 = ((_dafny.MultiSet(d_247_v97_)) - (d_248_v98_)).ispropersubset(default__.fm34((d_243_v93_).cf61, d_138_globalState_))
        d_249_v99_: _dafny.Map
        d_249_v99_ = _dafny.Map({d_122_v3_: d_141_v19_})
        default__.m0(((d_153_v25_)[len((d_249_v99_).set(d_122_v3_, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([D13_DC36(d_122_v3_, d_122_v3_)])) for d_250_i9_ in range(default__.abs(608))])))] if (len((d_249_v99_).set(d_122_v3_, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([D13_DC36(d_122_v3_, d_122_v3_)])) for d_251_i9_ in range(default__.abs(608))])))) in (d_153_v25_) else d_126_v7_), d_138_globalState_)
        hi1_ = d_126_v7_
        for d_252_i10_ in range(d_126_v7_, hi1_):
            d_253_v100_: _dafny.Array
            nw38_ = _dafny.Array(_dafny.Map({}), 27)
            d_253_v100_ = nw38_
            d_254_v101_: D14
            d_254_v101_ = D14_DC39(d_253_v100_)
            d_255_v102_: _dafny.Map
            d_255_v102_ = _dafny.Map({(_dafny.MultiSet(d_129_v10_)).cardinality: d_241_v91_})
            d_256_v103_: C3
            nw39_ = C3()
            nw39_.ctor__((d_254_v101_).cf68, d_255_v102_, d_122_v3_)
            d_256_v103_ = nw39_
            d_257_v104_: _dafny.Set
            d_257_v104_ = _dafny.Set({d_256_v103_})
            d_257_v104_ = ((d_257_v104_).intersection(d_257_v104_)) - ((d_257_v104_) - (d_257_v104_))
            d_258_v105_: _dafny.Set
            d_258_v105_ = _dafny.Set({d_124_v5_})
            rhs27_ = (0) - (d_126_v7_)
            rhs28_ = (d_252_i10_) + (d_126_v7_)
            rhs29_ = d_126_v7_
            rhs30_ = (d_258_v105_) - (d_258_v105_)
            d_126_v7_ = rhs27_
            d_126_v7_ = rhs28_
            d_126_v7_ = rhs29_
            d_258_v105_ = rhs30_
            index30_ = default__.safeIndex(456, (d_124_v5_).length(0))
            (d_124_v5_)[index30_] = d_126_v7_
            d_259_v106_: _dafny.Array
            nw40_ = _dafny.Array(None, 6)
            nw40_[int(0)] = d_135_v14_
            nw40_[int(1)] = d_135_v14_
            nw40_[int(2)] = d_135_v14_
            nw40_[int(3)] = d_135_v14_
            nw40_[int(4)] = d_135_v14_
            nw40_[int(5)] = d_135_v14_
            d_259_v106_ = nw40_
            d_260_v107_: T1
            nw41_ = C1()
            nw41_.ctor__(d_126_v7_, d_259_v106_, d_122_v3_)
            d_260_v107_ = nw41_
            index31_ = default__.safeIndex(456, (d_124_v5_).length(0))
            (d_124_v5_)[index31_] = len(_dafny.Set({d_260_v107_}))
            index32_ = default__.safeIndex(456, (d_124_v5_).length(0))
            (d_124_v5_)[index32_] = len(d_121_v2_)
        d_261_v108_: C5
        nw42_ = C5()
        nw42_.ctor__((d_129_v10_)[default__.safeIndex(d_126_v7_, len(d_129_v10_))], d_122_v3_)
        d_261_v108_ = nw42_
        d_262_v109_: D6
        d_262_v109_ = D6_DC16(d_120_v1_)
        source5_ = d_262_v109_
        if source5_.is_DC17:
            d_263___mcc_h7_ = source5_.cf25
            d_264___mcc_h8_ = source5_.cf26
            d_265___mcc_h9_ = source5_.cf27
            d_266_cf27_ = d_265___mcc_h9_
            d_267_cf26_ = d_264___mcc_h8_
            d_268_cf25_ = d_263___mcc_h7_
            d_269_v110_: _dafny.Map
            d_269_v110_ = _dafny.Map({default__.fm10(default__.fm20(d_121_v2_, d_246_v96_, (d_261_v108_).f30, (d_196_v61_).cardinality, d_138_globalState_), (d_261_v108_).f29, d_138_globalState_): (d_261_v108_).f30})
            d_270_v111_: _dafny.Map
            d_270_v111_ = _dafny.Map({default__.fm18(((d_269_v110_)[False] if (False) in (d_269_v110_) else d_268_cf25_), d_122_v3_, d_138_globalState_): d_269_v110_})
            d_270_v111_ = (d_270_v111_).set(d_126_v7_, d_269_v110_)
            d_266_cf27_ = default__.safeDivisionInt((d_126_v7_) - (len(d_129_v10_)), d_266_cf27_)
            default__.m0(d_126_v7_, d_138_globalState_)
            d_246_v96_ = (_dafny.MultiSet([len(d_121_v2_)]) if (d_261_v108_).f30 else d_246_v96_)
        elif source5_.is_DC16:
            d_271___mcc_h10_ = source5_.cf24
            d_272_cf24_ = d_271___mcc_h10_
            d_273_v112_: D0
            d_273_v112_ = D0_DC1(d_119_v0_, (d_261_v108_).f30, (d_261_v108_).f30)
            d_274_v113_: bool
            d_275_v114_: bool
            out4_: bool
            out5_: bool
            out4_, out5_ = (d_261_v108_).m1(d_141_v19_, -33, d_273_v112_, d_138_globalState_)
            d_274_v113_ = out4_
            d_275_v114_ = out5_
            if d_275_v114_:
                d_122_v3_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ei"))) != (default__.fm0(d_196_v61_, d_138_globalState_))
                d_121_v2_ = d_121_v2_
                default__.m0(d_126_v7_, d_138_globalState_)
                d_276_v115_: bool
                d_277_v116_: bool
                out6_: bool
                out7_: bool
                out6_, out7_ = (d_261_v108_).m1(d_141_v19_, d_126_v7_, d_273_v112_, d_138_globalState_)
                d_276_v115_ = out6_
                d_277_v116_ = out7_
                d_245_v95_ = (d_245_v95_) + ((d_245_v95_) + (d_245_v95_))
            elif True:
                (d_138_globalState_).f20 = d_135_v14_
                pat_let_tv4_ = d_119_v0_
                d_278_v117_: _dafny.Array
                nw43_ = _dafny.Array(None, 21)
                def iife22_(_pat_let6_0):
                    def iife23_(d_279_dt__update__tmp_h3_):
                        def iife24_(_pat_let7_0):
                            def iife25_(d_280_dt__update_hcf5_h1_):
                                return D0_DC3(d_280_dt__update_hcf5_h1_)
                            return iife25_(_pat_let7_0)
                        return iife24_(pat_let_tv4_)
                    return iife23_(_pat_let6_0)
                nw43_[int(0)] = iife22_(d_139_v17_)
                nw43_[int(1)] = d_139_v17_
                nw43_[int(2)] = d_139_v17_
                nw43_[int(3)] = D0_DC3(d_119_v0_)
                nw43_[int(4)] = d_139_v17_
                nw43_[int(5)] = d_139_v17_
                nw43_[int(6)] = d_139_v17_
                nw43_[int(7)] = d_139_v17_
                nw43_[int(8)] = d_139_v17_
                nw43_[int(9)] = d_139_v17_
                nw43_[int(10)] = d_139_v17_
                nw43_[int(11)] = D0_DC3(d_119_v0_)
                nw43_[int(12)] = d_139_v17_
                nw43_[int(13)] = d_139_v17_
                nw43_[int(14)] = d_139_v17_
                nw43_[int(15)] = d_139_v17_
                nw43_[int(16)] = d_139_v17_
                nw43_[int(17)] = d_139_v17_
                nw43_[int(18)] = d_139_v17_
                nw43_[int(19)] = d_139_v17_
                nw43_[int(20)] = d_139_v17_
                d_278_v117_ = nw43_
                d_281_v118_: _dafny.Seq
                d_281_v118_ = _dafny.SeqWithoutIsStrInference([d_278_v117_, d_278_v117_, d_278_v117_])
                d_282_v119_: _dafny.Map
                d_282_v119_ = _dafny.Map({default__.fm7(d_274_v113_, (d_261_v108_).f30, d_138_globalState_): d_281_v118_})
                d_282_v119_ = (d_282_v119_).set(d_126_v7_, d_281_v118_)
                d_283_v120_: D3
                d_283_v120_ = D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "soglnufp")), d_126_v7_, (d_261_v108_).f29)
                d_284_v121_: _dafny.Map
                d_284_v121_ = _dafny.Map({d_283_v120_: (d_261_v108_).f30})
                d_284_v121_ = _dafny.Map({d_283_v120_: (d_122_v3_) or (d_275_v114_)})
                d_285_v122_: D9
                d_285_v122_ = D9_DC25(d_126_v7_, (d_261_v108_).f30, default__.fm3(d_138_globalState_), d_275_v114_)
                d_286_v123_: C5
                nw44_ = C5()
                nw44_.ctor__(not(not(False)), not (False) or ((d_285_v122_).cf44))
                d_286_v123_ = nw44_
                d_287_v124_: C2
                nw45_ = C2()
                nw45_.ctor__(default__.fm2(d_274_v113_, d_126_v7_, (d_261_v108_).f29, ((d_241_v91_)[len(d_121_v2_)] if (len(d_121_v2_)) in (d_241_v91_) else (d_261_v108_).f30), d_138_globalState_), (125) == (d_126_v7_))
                d_287_v124_ = nw45_
                d_287_v124_ = d_287_v124_
            d_275_v114_ = (d_261_v108_).f29
            d_288_v125_: _dafny.Array
            def lambda6_(d_289_v3_, d_290_v108_):
                def lambda7_(d_291_i11_):
                    return _dafny.Map({d_289_v3_: (d_290_v108_).f30})

                return lambda7_

            init3_ = lambda6_(d_122_v3_, d_261_v108_)
            nw46_ = _dafny.Array(None, 25)
            for i0_3_ in range(nw46_.length(0)):
                nw46_[i0_3_] = init3_(i0_3_)
            d_288_v125_ = nw46_
            d_292_v126_: _dafny.Map
            d_292_v126_ = _dafny.Map({d_126_v7_: d_241_v91_})
            d_293_v127_: C3
            nw47_ = C3()
            nw47_.ctor__(d_288_v125_, d_292_v126_, (d_261_v108_).f30)
            d_293_v127_ = nw47_
            d_294_v128_: _dafny.MultiSet
            d_294_v128_ = _dafny.MultiSet([d_293_v127_, d_293_v127_, d_293_v127_, d_293_v127_])
            d_275_v114_ = default__.fm8((d_294_v128_).issubset(_dafny.MultiSet([d_293_v127_, d_293_v127_])), (d_126_v7_) - (d_126_v7_), d_138_globalState_)
        elif True:
            d_295___mcc_h11_ = source5_.cf28
            d_296_cf28_ = d_295___mcc_h11_
            d_297_v129_: C0
            nw48_ = C0()
            nw48_.ctor__((d_261_v108_).f30)
            d_297_v129_ = nw48_
            d_298_v130_: C0
            d_298_v130_ = d_297_v129_
            d_299_v131_: _dafny.Seq
            d_299_v131_ = _dafny.SeqWithoutIsStrInference([d_297_v129_, d_297_v129_])
            d_300_v132_: _dafny.Array
            nw49_ = _dafny.Array(None, 29)
            nw49_[int(0)] = d_297_v129_
            nw49_[int(1)] = d_297_v129_
            nw49_[int(2)] = d_297_v129_
            nw49_[int(3)] = d_297_v129_
            nw49_[int(4)] = d_297_v129_
            nw49_[int(5)] = d_297_v129_
            nw49_[int(6)] = d_297_v129_
            nw49_[int(7)] = d_297_v129_
            nw49_[int(8)] = d_297_v129_
            nw49_[int(9)] = d_297_v129_
            nw49_[int(10)] = d_297_v129_
            nw49_[int(11)] = d_297_v129_
            nw49_[int(12)] = (d_298_v130_)
            nw49_[int(13)] = d_297_v129_
            nw49_[int(14)] = d_297_v129_
            nw49_[int(15)] = d_297_v129_
            nw49_[int(16)] = d_297_v129_
            nw49_[int(17)] = d_297_v129_
            nw49_[int(18)] = d_297_v129_
            nw49_[int(19)] = (d_299_v131_)[default__.safeIndex((0) - (d_126_v7_), len(d_299_v131_))]
            nw49_[int(20)] = d_297_v129_
            nw49_[int(21)] = d_297_v129_
            nw49_[int(22)] = d_297_v129_
            nw49_[int(23)] = d_297_v129_
            nw49_[int(24)] = (d_299_v131_)[default__.safeIndex(d_126_v7_, len(d_299_v131_))]
            nw49_[int(25)] = (d_298_v130_)
            nw49_[int(26)] = d_297_v129_
            nw49_[int(27)] = d_297_v129_
            nw49_[int(28)] = d_297_v129_
            d_300_v132_ = nw49_
            d_126_v7_ = len(_dafny.SeqWithoutIsStrInference([d_300_v132_, d_300_v132_, d_300_v132_, d_300_v132_, d_300_v132_]))
            rhs31_ = (d_129_v10_)[default__.safeIndex(d_126_v7_, len(d_129_v10_))]
            rhs32_ = default__.safeDivisionInt(default__.safeModuloInt(d_126_v7_, d_126_v7_), (d_126_v7_) + (len(d_129_v10_)))
            rhs33_ = default__.fm1(d_119_v0_, d_138_globalState_)
            lhs18_ = d_138_globalState_
            lhs18_.f16 = rhs31_
            d_126_v7_ = rhs32_
            d_126_v7_ = rhs33_
            if d_122_v3_:
                d_301_v133_: _dafny.Array
                nw50_ = _dafny.Array(_dafny.Seq({}), 21)
                d_301_v133_ = nw50_
                index33_ = default__.safeIndex(629, (d_301_v133_).length(0))
                (d_301_v133_)[index33_] = _dafny.SeqWithoutIsStrInference([d_135_v14_])
                index34_ = default__.safeIndex(963, (d_124_v5_).length(0))
                (d_124_v5_)[index34_] = default__.safeDivisionInt((d_246_v96_).cardinality, -566)
                d_302_v134_: _dafny.Set
                d_302_v134_ = _dafny.Set({(d_261_v108_).f30})
                d_303_v135_: _dafny.Map
                d_303_v135_ = _dafny.Map({d_126_v7_: d_302_v134_})
                index35_ = default__.safeIndex(629, (d_301_v133_).length(0))
                index36_ = default__.safeIndex(963, (d_124_v5_).length(0))
                rhs34_ = (d_245_v95_) + (d_245_v95_)
                rhs35_ = (default__.safeModuloInt(d_126_v7_, len(((d_303_v135_)[d_126_v7_] if (d_126_v7_) in (d_303_v135_) else _dafny.Set({True}))))) + ((d_126_v7_) + (((d_128_v9_)[(d_261_v108_).f30] if ((d_261_v108_).f30) in (d_128_v9_) else 885)))
                rhs36_ = default__.fm7((d_261_v108_).f29, d_122_v3_, d_138_globalState_)
                rhs37_ = 643
                rhs38_ = (d_261_v108_).f29
                lhs19_ = d_301_v133_
                lhs20_ = default__.safeIndex(629, (d_301_v133_).length(0))
                lhs21_ = d_124_v5_
                lhs22_ = default__.safeIndex(963, (d_124_v5_).length(0))
                lhs19_[lhs20_] = rhs34_
                d_126_v7_ = rhs35_
                lhs21_[lhs22_] = rhs36_
                d_126_v7_ = rhs37_
                d_122_v3_ = rhs38_
                d_126_v7_ = (0) - (default__.fm1(_dafny.CodePoint('r'), d_138_globalState_))
                d_304_v136_: _dafny.Seq
                d_304_v136_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_126_v7_: (d_261_v108_).f30})])
                (d_138_globalState_).f16 = (d_241_v91_) not in (d_304_v136_)
                (d_138_globalState_).f16 = not(not((((d_261_v108_).f29) and ((d_261_v108_).f30) if ((d_261_v108_).f30) == ((d_261_v108_).f30) else not(not((d_261_v108_).f30)))))
                rhs39_ = d_126_v7_
                rhs40_ = ((d_301_v133_)[default__.safeIndex(629, (d_301_v133_).length(0))])[default__.safeIndex(d_126_v7_, len((d_301_v133_)[default__.safeIndex(629, (d_301_v133_).length(0))]))]
                rhs41_ = d_121_v2_
                rhs42_ = (d_121_v2_).set(default__.safeIndex((0) - (len(d_128_v9_)), len(d_121_v2_)), d_119_v0_)
                lhs23_ = d_138_globalState_
                d_126_v7_ = rhs39_
                lhs23_.f27 = rhs40_
                d_121_v2_ = rhs41_
                d_121_v2_ = rhs42_
            elif True:
                d_305_v137_: D3
                d_305_v137_ = D3_DC10(d_121_v2_, d_126_v7_, d_122_v3_)
                d_126_v7_ = (d_305_v137_).cf14
                d_126_v7_ = 915
                default__.m0(d_126_v7_, d_138_globalState_)
                d_306_v138_: _dafny.Seq
                d_306_v138_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_119_v0_ for d_307_i12_ in range(default__.abs(-779))])])
                d_126_v7_ = len((((d_306_v138_ if (d_261_v108_).f29 else d_306_v138_)) + (d_306_v138_)).set(default__.safeIndex(d_126_v7_, len(((d_306_v138_ if (d_261_v108_).f29 else d_306_v138_)) + (d_306_v138_))), (d_121_v2_) + (d_121_v2_)))
                (d_138_globalState_).f16 = (d_261_v108_).f29
            d_126_v7_ = (0) - ((0) - (d_126_v7_))
        d_308_i13_: int
        d_308_i13_ = 0
        with _dafny.label("3"):
            while d_122_v3_:
                with _dafny.c_label("3"):
                    if (d_308_i13_) >= (100):
                        raise _dafny.Break("3")
                    d_308_i13_ = (d_308_i13_) + (1)
                    if (d_129_v10_)[default__.safeIndex(d_126_v7_, len(d_129_v10_))]:
                        index37_ = default__.safeIndex(148, (d_124_v5_).length(0))
                        (d_124_v5_)[index37_] = default__.safeDivisionInt(d_126_v7_, d_126_v7_)
                        d_309_v139_: _dafny.Array
                        nw51_ = _dafny.Array(_dafny.Array(None, 0), 2)
                        d_309_v139_ = nw51_
                        index38_ = default__.safeIndex(349, (d_309_v139_).length(0))
                        (d_309_v139_)[index38_] = d_135_v14_
                        d_310_v140_: _dafny.Seq
                        d_310_v140_ = _dafny.SeqWithoutIsStrInference([default__.fm14(d_126_v7_, d_126_v7_, d_126_v7_, (d_261_v108_).f29, d_138_globalState_), ((_dafny.SeqWithoutIsStrInference([(d_261_v108_).f29])).set(default__.safeIndex(-53, len(_dafny.SeqWithoutIsStrInference([(d_261_v108_).f29]))), (d_261_v108_).f30)) + (d_129_v10_)])
                        index39_ = default__.safeIndex(148, (d_124_v5_).length(0))
                        index40_ = default__.safeIndex(349, (d_309_v139_).length(0))
                        rhs43_ = len(_dafny.Map({d_141_v19_: (d_261_v108_).f30}))
                        rhs44_ = (d_121_v2_) + (d_121_v2_)
                        rhs45_ = d_135_v14_
                        rhs46_ = (d_310_v140_) + (d_310_v140_)
                        rhs47_ = 101
                        lhs24_ = d_124_v5_
                        lhs25_ = default__.safeIndex(148, (d_124_v5_).length(0))
                        lhs26_ = d_309_v139_
                        lhs27_ = default__.safeIndex(349, (d_309_v139_).length(0))
                        lhs24_[lhs25_] = rhs43_
                        d_121_v2_ = rhs44_
                        lhs26_[lhs27_] = rhs45_
                        d_310_v140_ = rhs46_
                        d_126_v7_ = rhs47_
                        index41_ = default__.safeIndex(148, (d_124_v5_).length(0))
                        (d_124_v5_)[index41_] = d_126_v7_
                        (d_138_globalState_).f16 = (d_261_v108_).f30
                        d_311_v141_: C1
                        nw52_ = C1()
                        nw52_.ctor__((d_124_v5_)[default__.safeIndex(148, (d_124_v5_).length(0))], d_309_v139_, not((d_261_v108_).f30))
                        d_311_v141_ = nw52_
                        d_311_v141_ = d_311_v141_
                        index42_ = default__.safeIndex(148, (d_124_v5_).length(0))
                        (d_124_v5_)[index42_] = (d_126_v7_) * (((d_124_v5_)[default__.safeIndex(148, (d_124_v5_).length(0))]) - ((d_141_v19_)[default__.safeIndex(d_126_v7_, len(d_141_v19_))]))
                    elif True:
                        d_312_v142_: C5
                        nw53_ = C5()
                        nw53_.ctor__((d_261_v108_).f29, (d_261_v108_).f30)
                        d_312_v142_ = nw53_
                        d_313_v143_: D0
                        d_313_v143_ = D0_DC1(d_119_v0_, (d_261_v108_).f30, True)
                        d_314_v144_: bool
                        d_315_v145_: bool
                        out8_: bool
                        out9_: bool
                        out8_, out9_ = (d_312_v142_).m1(d_141_v19_, default__.safeDivisionInt(d_126_v7_, len(_dafny.Map({d_129_v10_: d_126_v7_}))), d_313_v143_, d_138_globalState_)
                        d_314_v144_ = out8_
                        d_315_v145_ = out9_
                        d_316_v146_: bool
                        d_317_v147_: bool
                        out10_: bool
                        out11_: bool
                        out10_, out11_ = (d_312_v142_).m1((_dafny.SeqWithoutIsStrInference([d_126_v7_ for d_318_i14_ in range(default__.abs(877))])) + (d_141_v19_), d_126_v7_, d_313_v143_, d_138_globalState_)
                        d_316_v146_ = out10_
                        d_317_v147_ = out11_
                        d_316_v146_ = (d_261_v108_).f30
                        d_319_v148_: _dafny.Array
                        nw54_ = _dafny.Array(_dafny.Seq({}), 26)
                        d_319_v148_ = nw54_
                        d_320_v149_: _dafny.Array
                        nw55_ = _dafny.Array(None, 17)
                        nw55_[int(0)] = d_319_v148_
                        nw55_[int(1)] = d_319_v148_
                        nw55_[int(2)] = d_319_v148_
                        nw55_[int(3)] = d_319_v148_
                        nw55_[int(4)] = d_319_v148_
                        nw55_[int(5)] = d_319_v148_
                        nw55_[int(6)] = d_319_v148_
                        nw55_[int(7)] = d_319_v148_
                        nw55_[int(8)] = d_319_v148_
                        nw55_[int(9)] = d_319_v148_
                        nw55_[int(10)] = d_319_v148_
                        nw55_[int(11)] = d_319_v148_
                        nw55_[int(12)] = d_319_v148_
                        nw55_[int(13)] = d_319_v148_
                        nw55_[int(14)] = d_319_v148_
                        nw55_[int(15)] = d_319_v148_
                        nw55_[int(16)] = d_319_v148_
                        d_320_v149_ = nw55_
                        index43_ = default__.safeIndex(558, (d_320_v149_).length(0))
                        (d_320_v149_)[index43_] = d_319_v148_
                        index44_ = default__.safeIndex(558, (d_320_v149_).length(0))
                        (d_320_v149_)[index44_] = d_319_v148_
                    if ((d_129_v10_)[default__.safeIndex(len(d_242_v92_), len(d_129_v10_))]) and (not(default__.fm26((d_261_v108_).f29, (d_261_v108_).f29, (d_129_v10_)[default__.safeIndex(d_126_v7_, len(d_129_v10_))], ((d_153_v25_)[default__.fm1(d_119_v0_, d_138_globalState_)] if (default__.fm1(d_119_v0_, d_138_globalState_)) in (d_153_v25_) else d_126_v7_), d_138_globalState_))):
                        d_321_v150_: _dafny.Set
                        d_321_v150_ = _dafny.Set({d_119_v0_})
                        d_322_v151_: _dafny.Map
                        d_322_v151_ = _dafny.Map({d_321_v150_: d_121_v2_})
                        d_322_v151_ = (d_322_v151_).set((d_321_v150_).intersection(d_321_v150_), d_121_v2_)
                        index45_ = default__.safeIndex(810, (d_135_v14_).length(0))
                        (d_135_v14_)[index45_] = (d_261_v108_).f29
                        index46_ = default__.safeIndex(810, (d_135_v14_).length(0))
                        (d_135_v14_)[index46_] = (d_126_v7_) > ((0) - (default__.safeDivisionInt(len(d_129_v10_), d_126_v7_)))
                        d_126_v7_ = 929
                        d_323_v152_: _dafny.Map
                        d_323_v152_ = _dafny.Map({d_124_v5_: d_124_v5_})
                        d_323_v152_ = (d_323_v152_).set(d_124_v5_, d_124_v5_)
                        d_324_v153_: D0
                        d_324_v153_ = D0_DC1(d_119_v0_, (d_261_v108_).f29, (d_261_v108_).f30)
                        d_325_v154_: bool
                        d_326_v155_: bool
                        out12_: bool
                        out13_: bool
                        out12_, out13_ = (d_261_v108_).m1(d_141_v19_, d_126_v7_, d_324_v153_, d_138_globalState_)
                        d_325_v154_ = out12_
                        d_326_v155_ = out13_
                    elif True:
                        d_327_v156_: _dafny.Set
                        d_327_v156_ = _dafny.Set({324})
                        d_126_v7_ = len((_dafny.Set({-408})).intersection(d_327_v156_))
                        d_241_v91_ = (d_241_v91_).set(len(d_241_v91_), (d_261_v108_).f30)
                        d_126_v7_ = default__.safeDivisionInt(d_126_v7_, (len(d_121_v2_)) * (d_126_v7_))
                        d_126_v7_ = (0) - (d_126_v7_)
                        d_328_v157_: _dafny.Array
                        nw56_ = _dafny.Array(_dafny.Map({}), 29)
                        d_328_v157_ = nw56_
                        d_329_v158_: _dafny.Map
                        d_329_v158_ = _dafny.Map({default__.fm1(d_119_v0_, d_138_globalState_): d_241_v91_})
                        d_330_v159_: C3
                        nw57_ = C3()
                        nw57_.ctor__(d_328_v157_, d_329_v158_, (d_261_v108_).f29)
                        d_330_v159_ = nw57_
                    rhs48_ = (d_261_v108_).f29
                    rhs49_ = (d_121_v2_ if d_122_v3_ else (default__.fm0(d_196_v61_, d_138_globalState_)).set(default__.safeIndex(d_126_v7_, len(default__.fm0(d_196_v61_, d_138_globalState_))), d_119_v0_))
                    rhs50_ = default__.fm10(d_119_v0_, (d_126_v7_) == (d_126_v7_), d_138_globalState_)
                    rhs51_ = (_dafny.CodePoint('k') if (d_261_v108_).f30 else d_119_v0_)
                    lhs28_ = d_138_globalState_
                    lhs29_ = d_138_globalState_
                    lhs28_.f28 = rhs48_
                    d_121_v2_ = rhs49_
                    lhs29_.f28 = rhs50_
                    d_119_v0_ = rhs51_
                    d_331_v160_: _dafny.Array
                    def lambda8_(d_332_v108_, d_333_v7_):
                        def lambda9_(d_334_i15_):
                            return D14_DC40((d_332_v108_).f29, (d_332_v108_).f29, d_333_v7_)

                        return lambda9_

                    init4_ = lambda8_(d_261_v108_, d_126_v7_)
                    nw58_ = _dafny.Array(None, 22)
                    for i0_4_ in range(nw58_.length(0)):
                        nw58_[i0_4_] = init4_(i0_4_)
                    d_331_v160_ = nw58_
                    d_335_v161_: D14
                    d_335_v161_ = D14_DC40(not((d_261_v108_).f30), d_122_v3_, d_126_v7_)
                    index47_ = default__.safeIndex(601, (d_331_v160_).length(0))
                    (d_331_v160_)[index47_] = d_335_v161_
                    index48_ = default__.safeIndex(601, (d_331_v160_).length(0))
                    (d_331_v160_)[index48_] = (d_335_v161_ if (d_261_v108_).f29 else d_335_v161_)
                    pass
            pass
        _dafny.print(_dafny.string_of(d_119_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v1_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v1_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v1_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_121_v2_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_122_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v4_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mau")): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v5_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_125_v6_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_126_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_127_v8_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_128_v9_) == (_dafny.Map({False: 159, True: 62}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v10_) == (_dafny.SeqWithoutIsStrInference([False, False, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_v11_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([159, 159]): 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_v12_)[0]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_v12_)[1]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_v12_)[2]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_v12_)[3]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_v12_)[4]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_v12_)[5]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_134_v13_).cf0)[0]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_134_v13_).cf0)[1]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_134_v13_).cf0)[2]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_134_v13_).cf0)[3]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_134_v13_).cf0)[4]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_134_v13_).cf0)[5]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v14_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v14_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v14_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v14_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v14_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v14_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v14_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v14_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v14_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v14_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v14_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v14_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v15_))[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v16_) == (_dafny.Map({False: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f0)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f0)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f0)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f3) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mau")): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_138_globalState_).f6).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_138_globalState_).f11)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f12) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([159, 159]): 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_globalState_.f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_138_globalState_).f18)[0]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_138_globalState_).f18)[1]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_138_globalState_).f18)[2]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_138_globalState_).f18)[3]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_138_globalState_).f18)[4]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_138_globalState_).f18)[5]) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f23) == (_dafny.Map({False: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f24)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f27)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f27)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f27)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f27)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f27)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f27)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f27)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f27)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f27)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f27)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f27)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f27)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_globalState_.f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_139_v17_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_140_v18_) == (_dafny.Set({D0_DC3(_dafny.CodePoint('v'))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_141_v19_) == (_dafny.SeqWithoutIsStrInference([159, 159, 159, 159]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v24_) == (_dafny.Map({159: 159}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v25_) == (_dafny.Map({159: 159}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_187_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_190_i5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v61_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_197_v62_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, True, True, True]), _dafny.MultiSet([True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_v63_).cf18) == (_dafny.MultiSet([True, True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_241_v91_) == (_dafny.Map({3: True, 1: True, 2: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_242_v92_) == (_dafny.Map({_dafny.Map({3: True, 796: True}): 796}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_243_v93_).cf61) == (_dafny.Map({_dafny.Map({3: True, 796: True}): 796}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_244_v94_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_245_v95_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_v96_) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v97_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([1])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_248_v98_) == (_dafny.MultiSet([_dafny.MultiSet([1]), _dafny.MultiSet([1]), _dafny.MultiSet([1])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_249_v99_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([159, 159, 159, 159])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v108_).f29))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v108_).f30))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v109_).cf24)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v109_).cf24)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v109_).cf24)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_308_i13_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.CodePoint('D'), False, False)
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
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D0_DC4)

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

class D0_DC3(D0, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC4(D0, NamedTuple('DC4', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC4({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC4) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC5(D1, NamedTuple('DC5', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(_dafny.CodePoint('D'), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC7(D2, NamedTuple('DC7', [('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)

class D3_DC10(D3, NamedTuple('DC10', [('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({self.cf13.VerbatimString(True)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC12(D4, NamedTuple('DC12', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14(_dafny.Set({}), int(0), False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC14(D5, NamedTuple('DC14', [('cf19', Any), ('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(False, _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)

class D6_DC17(D6, NamedTuple('DC17', [('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)

class D7_DC19(D7, NamedTuple('DC19', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC21(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)

class D8_DC21(D8, NamedTuple('DC21', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC24(_dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)

class D9_DC24(D9, NamedTuple('DC24', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf41', Any), ('cf42', Any), ('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf45', Any), ('cf46', Any), ('cf47', Any), ('cf48', Any), ('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf45)}, {self.cf46.VerbatimString(True)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC28(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)

class D10_DC28(D10, NamedTuple('DC28', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC27(D10, NamedTuple('DC27', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC29(D10, NamedTuple('DC29', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)

class D11_DC30(D11, NamedTuple('DC30', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC32(False, False, None, _dafny.Set({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)

class D12_DC32(D12, NamedTuple('DC32', [('cf55', Any), ('cf56', Any), ('cf57', Any), ('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf55 == __o.cf55 and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC33(D12, NamedTuple('DC33', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({self.cf60.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC31(D12, NamedTuple('DC31', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC35(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D13_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D13_DC37)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D13_DC38)

class D13_DC35(D13, NamedTuple('DC35', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC36(D13, NamedTuple('DC36', [('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC37(D13, NamedTuple('DC37', [('cf65', Any), ('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC37({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC37) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC34(D13, NamedTuple('DC34', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC38(D13, NamedTuple('DC38', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC38({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC38) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC40(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D14_DC40)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D14_DC39)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D14_DC42)

class D14_DC40(D14, NamedTuple('DC40', [('cf69', Any), ('cf70', Any), ('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC40({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC40) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC41(D14, NamedTuple('DC41', [('cf72', Any), ('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf72 == __o.cf72 and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC39(D14, NamedTuple('DC39', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC39({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC39) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC42(D14, NamedTuple('DC42', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC42({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC42) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)

class D15_DC43(D15, NamedTuple('DC43', [('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43({_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43) and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC45(_dafny.CodePoint('D'), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D16_DC44)

class D16_DC45(D16, NamedTuple('DC45', [('cf77', Any), ('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC44(D16, NamedTuple('DC44', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC44({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC44) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC47(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D17_DC47)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D17_DC46)

class D17_DC47(D17, NamedTuple('DC47', [('cf81', Any), ('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC47({_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC47) and self.cf81 == __o.cf81 and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC46(D17, NamedTuple('DC46', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC46({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC46) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC49()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D18_DC49)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D18_DC50)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D18_DC51)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D18_DC48)

class D18_DC49(D18, NamedTuple('DC49', [])):
    def __dafnystr__(self) -> str:
        return f'D18.DC49'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC49)
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC50(D18, NamedTuple('DC50', [('cf84', Any), ('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC50({_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC50) and self.cf84 == __o.cf84 and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC51(D18, NamedTuple('DC51', [('cf86', Any), ('cf87', Any), ('cf88', Any), ('cf89', Any), ('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC51({_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC51) and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88 and self.cf89 == __o.cf89 and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC48(D18, NamedTuple('DC48', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC48({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC48) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC53(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D19_DC53)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D19_DC54)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D19_DC52)

class D19_DC53(D19, NamedTuple('DC53', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC53({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC53) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC54(D19, NamedTuple('DC54', [])):
    def __dafnystr__(self) -> str:
        return f'D19.DC54'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC54)
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC52(D19, NamedTuple('DC52', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC52({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC52) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    def m2(self, p0, p1, p2, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f12: _dafny.Map = _dafny.Map({})
        self.f16: bool = False
        self.f20: _dafny.Array = _dafny.Array(None, 0)
        self.f24: _dafny.Array = _dafny.Array(None, 0)
        self.f27: _dafny.Array = _dafny.Array(None, 0)
        self.f28: bool = False
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        self._f1: bool = False
        self._f2: bool = False
        self._f3: _dafny.Map = _dafny.Map({})
        self._f4: bool = False
        self._f5: bool = False
        self._f6: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f7: int = int(0)
        self._f8: bool = False
        self._f9: bool = False
        self._f10: bool = False
        self._f11: _dafny.Seq = _dafny.Seq({})
        self._f13: int = int(0)
        self._f14: bool = False
        self._f15: int = int(0)
        self._f17: bool = False
        self._f18: _dafny.Array = _dafny.Array(None, 0)
        self._f19: bool = False
        self._f21: _dafny.Array = _dafny.Array(None, 0)
        self._f22: bool = False
        self._f23: _dafny.Map = _dafny.Map({})
        self._f25: int = int(0)
        self._f26: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27, f28):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self).f20 = f20
        (self)._f21 = f21
        (self)._f22 = f22
        (self)._f23 = f23
        (self).f24 = f24
        (self)._f25 = f25
        (self)._f26 = f26
        (self).f27 = f27
        (self).f28 = f28

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
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19
    @property
    def f21(self):
        return self._f21
    @property
    def f22(self):
        return self._f22
    @property
    def f23(self):
        return self._f23
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26

class C0(T0):
    def  __init__(self):
        self._f31: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f31(self):
        return self._f31
    def ctor__(self, f31):
        (self)._f31 = f31


class C1(T1, T0):
    def  __init__(self):
        self._f31: bool = False
        self._f32: int = int(0)
        self._f33: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f31(self):
        return self._f31
    @property
    def f32(self):
        return self._f32
    @property
    def f33(self):
        return self._f33
    def ctor__(self, f32, f33, f31):
        (self)._f32 = f32
        (self)._f33 = f33
        (self)._f31 = f31

    def fm6(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(self).f31, False, (self).f31, True, not(((self).f32) == ((self).f32))])

    def fm16(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_336_i0_ in range(default__.abs(531))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_337_i1_ in range(default__.abs(432))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lriggcwfc"))))

    def fm17(self, p0, globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f32]))).issubset((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndqhhq"))), (self).f32, 679]) if (self).f31 else _dafny.MultiSet([-370, (self).f32, (_dafny.MultiSet([(self).f31, not((self).f31), True, (self).f31, (self).f31])).cardinality])))

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r3: int = int(0)
        d_338_v0_: _dafny.Seq
        d_338_v0_ = _dafny.SeqWithoutIsStrInference([not(p0), (self).f31, not(False), (self).f31])
        d_339_v1_: _dafny.Seq
        d_339_v1_ = _dafny.SeqWithoutIsStrInference([(self).f32, (self).f32, len(d_338_v0_), default__.fm18((self).f31, p0, globalState), (self).f32])
        d_340_v2_: _dafny.Map
        d_340_v2_ = _dafny.Map({269: d_339_v1_})
        d_340_v2_ = (d_340_v2_).set((self).f32, (_dafny.SeqWithoutIsStrInference([(self).f32, (self).f32])) + (d_339_v1_))
        r0 = len(((_dafny.SeqWithoutIsStrInference([(d_338_v0_)[default__.safeIndex(len(d_339_v1_), len(d_338_v0_))]])) + ((_dafny.SeqWithoutIsStrInference([(self).f31, (self).f31, (self).f31])) + (d_338_v0_))).set(default__.safeIndex(((self).f32) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_341_i0_ in range(default__.abs(-955))]))), len((_dafny.SeqWithoutIsStrInference([(d_338_v0_)[default__.safeIndex(len(d_339_v1_), len(d_338_v0_))]])) + ((_dafny.SeqWithoutIsStrInference([(self).f31, (self).f31, (self).f31])) + (d_338_v0_)))), p0))
        d_342_v3_: _dafny.Array
        nw59_ = _dafny.Array(_dafny.Array(None, 0), 8)
        d_342_v3_ = nw59_
        d_342_v3_ = d_342_v3_
        d_343_i1_: int
        d_343_i1_ = 0
        with _dafny.label("4"):
            while False:
                with _dafny.c_label("4"):
                    if (d_343_i1_) >= (100):
                        raise _dafny.Break("4")
                    d_343_i1_ = (d_343_i1_) + (1)
                    d_344_v5_: _dafny.Map
                    d_344_v5_ = _dafny.Map({p0: True})
                    d_345_v6_: _dafny.Seq
                    d_345_v6_ = _dafny.SeqWithoutIsStrInference([d_344_v5_])
                    d_346_v7_: C0
                    nw60_ = C0()
                    nw60_.ctor__(p0)
                    d_346_v7_ = nw60_
                    d_347_v8_: _dafny.MultiSet
                    d_347_v8_ = _dafny.MultiSet([d_346_v7_, d_346_v7_])
                    d_348_v9_: _dafny.Map
                    d_348_v9_ = _dafny.Map({d_344_v5_: (d_347_v8_).cardinality})
                    def iife26_():
                        coll10_ = _dafny.Map()
                        compr_10_: _dafny.Map
                        for compr_10_ in (d_345_v6_).Elements:
                            d_349_v4_: _dafny.Map = compr_10_
                            if (d_349_v4_) in (d_345_v6_):
                                coll10_[d_349_v4_] = (self).f32
                        return _dafny.Map(coll10_)
                    r3 = len((iife26_()
                    ) | (d_348_v9_))
                    r0 = (self).f32
                    d_350_v10_: str
                    d_350_v10_ = _dafny.CodePoint('m')
                    if (self).fm17(d_350_v10_, globalState):
                        d_351_v11_: _dafny.Set
                        d_351_v11_ = _dafny.Set({(self).f32})
                        (globalState).f16 = ((d_351_v11_).ispropersubset(d_351_v11_) if False else ((self).f32) > ((self).f32))
                        d_352_v12_: _dafny.MultiSet
                        d_352_v12_ = _dafny.MultiSet([((self).f32) != ((self).f32), ((self).f32) < (len(d_338_v0_))])
                        d_352_v12_ = (p1 if p0 else p2)
                        (globalState).f28 = p0
                        d_353_v13_: _dafny.Set
                        d_353_v13_ = _dafny.Set({(self).f31, p0})
                        d_354_v14_: D2
                        d_354_v14_ = D2_DC7(d_350_v10_, d_353_v13_)
                        d_355_v15_: _dafny.Seq
                        d_355_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hxuqqi"))
                        d_356_v16_: _dafny.Map
                        d_356_v16_ = _dafny.Map({(d_354_v14_).cf10: (d_355_v15_) == (_dafny.SeqWithoutIsStrInference([d_350_v10_ for d_357_i2_ in range(default__.abs(726))]))})
                        d_356_v16_ = d_356_v16_
                        d_358_v17_: C0
                        nw61_ = C0()
                        nw61_.ctor__((self).f31)
                        d_358_v17_ = nw61_
                    elif True:
                        d_359_v18_: _dafny.Seq
                        d_359_v18_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_360_v19_: _dafny.Array
                        nw62_ = _dafny.Array(None, 5)
                        nw62_[int(0)] = d_350_v10_
                        nw62_[int(1)] = d_350_v10_
                        nw62_[int(2)] = d_350_v10_
                        nw62_[int(3)] = d_350_v10_
                        nw62_[int(4)] = _dafny.CodePoint('v')
                        d_360_v19_ = nw62_
                        d_361_v20_: D6
                        d_361_v20_ = D6_DC16(d_360_v19_)
                        d_362_v21_: _dafny.Set
                        d_362_v21_ = _dafny.Set({d_360_v19_, (d_361_v20_).cf24, d_360_v19_, d_360_v19_, d_360_v19_})
                        d_363_v22_: _dafny.Map
                        d_363_v22_ = _dafny.Map({(d_359_v18_)[default__.safeIndex(len(d_362_v21_), len(d_359_v18_))]: (self).f31})
                        d_364_v24_: _dafny.Set
                        d_364_v24_ = _dafny.Set({p1, p2})
                        d_365_v25_: _dafny.Array
                        nw63_ = _dafny.Array(None, 10)
                        nw63_[int(0)] = d_363_v22_
                        nw63_[int(1)] = (d_363_v22_) | (d_363_v22_)
                        nw63_[int(2)] = (_dafny.Map({p2: p0})) | (d_363_v22_)
                        nw63_[int(3)] = _dafny.Map({p2: (self).f31})
                        nw63_[int(4)] = d_363_v22_
                        nw63_[int(5)] = d_363_v22_
                        def iife27_():
                            coll11_ = _dafny.Map()
                            compr_11_: _dafny.MultiSet
                            for compr_11_ in (d_364_v24_).Elements:
                                d_366_v23_: _dafny.MultiSet = compr_11_
                                if (d_366_v23_) in (d_364_v24_):
                                    coll11_[d_366_v23_] = p0
                            return _dafny.Map(coll11_)
                        nw63_[int(6)] = iife27_()
                        
                        nw63_[int(7)] = d_363_v22_
                        nw63_[int(8)] = (d_363_v22_) | (d_363_v22_)
                        nw63_[int(9)] = (d_363_v22_) | (d_363_v22_)
                        d_365_v25_ = nw63_
                        index49_ = default__.safeIndex(509, (d_365_v25_).length(0))
                        def iife28_():
                            coll12_ = _dafny.Map()
                            compr_12_: _dafny.MultiSet
                            for compr_12_ in (_dafny.SeqWithoutIsStrInference([p2])).Elements:
                                d_367_v26_: _dafny.MultiSet = compr_12_
                                if (d_367_v26_) in (_dafny.SeqWithoutIsStrInference([p2])):
                                    coll12_[d_367_v26_] = not(p0)
                            return _dafny.Map(coll12_)
                        (d_365_v25_)[index49_] = iife28_()
                        
                        d_368_v27_: _dafny.Map
                        d_368_v27_ = d_363_v22_
                        index50_ = default__.safeIndex(509, (d_365_v25_).length(0))
                        (d_365_v25_)[index50_] = ((d_368_v27_ if p0 else d_368_v27_))
                        d_369_v28_: _dafny.Seq
                        d_369_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyqocej"))
                        d_370_v29_: D3
                        d_370_v29_ = D3_DC10(d_369_v28_, 398, False)
                        d_371_v30_: _dafny.Map
                        d_371_v30_ = _dafny.Map({(d_370_v29_).cf13: p0})
                        d_372_v31_: _dafny.Array
                        nw64_ = _dafny.Array(None, 18)
                        nw64_[int(0)] = (self).f31
                        nw64_[int(1)] = False
                        nw64_[int(2)] = (self).f31
                        nw64_[int(3)] = p0
                        nw64_[int(4)] = (self).f31
                        nw64_[int(5)] = ((d_371_v30_)[d_369_v28_] if (d_369_v28_) in (d_371_v30_) else not(p0))
                        nw64_[int(6)] = (self).f31
                        nw64_[int(7)] = (self).f31
                        nw64_[int(8)] = (self).f31
                        nw64_[int(9)] = (self).f31
                        nw64_[int(10)] = (self).f31
                        nw64_[int(11)] = False
                        nw64_[int(12)] = (_dafny.MultiSet(d_339_v1_)).isdisjoint((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f32: p0})) for d_373_i3_ in range(default__.abs(360))]))).set((self).f32, default__.abs((self).f32)))
                        nw64_[int(13)] = p0
                        nw64_[int(14)] = ((self).f31) == (p0)
                        nw64_[int(15)] = p0
                        nw64_[int(16)] = p0
                        nw64_[int(17)] = p0
                        d_372_v31_ = nw64_
                        index51_ = default__.safeIndex(928, (d_372_v31_).length(0))
                        (d_372_v31_)[index51_] = ((self).f31) and ((self).f31)
                        index52_ = default__.safeIndex(928, (d_372_v31_).length(0))
                        (d_372_v31_)[index52_] = False
                        (globalState).f16 = (d_372_v31_)[default__.safeIndex(928, (d_372_v31_).length(0))]
                        index53_ = default__.safeIndex(928, (d_372_v31_).length(0))
                        (d_372_v31_)[index53_] = (((self).f32) <= (297) if (self).f31 else (len(d_339_v1_)) == ((p1).cardinality))
                        r0 = (self).f32
                    d_374_v32_: _dafny.Map
                    d_374_v32_ = _dafny.Map({(self).f32: (self).f32})
                    d_374_v32_ = (d_374_v32_).set(982, 652)
                    pass
            pass
        d_375_v33_: _dafny.Array
        def lambda10_(d_376_i4_):
            return (d_376_i4_) * ((self).f32)

        init5_ = lambda10_
        nw65_ = _dafny.Array(None, 14)
        for i0_5_ in range(nw65_.length(0)):
            nw65_[i0_5_] = init5_(i0_5_)
        d_375_v33_ = nw65_
        index54_ = default__.safeIndex(109, (d_375_v33_).length(0))
        (d_375_v33_)[index54_] = (self).f32
        d_377_v34_: _dafny.Map
        d_377_v34_ = _dafny.Map({(self).f32: not((self).f31)})
        index55_ = default__.safeIndex(109, (d_375_v33_).length(0))
        (d_375_v33_)[index55_] = len((d_377_v34_) | ((d_377_v34_) | (d_377_v34_)))
        d_378_v35_: _dafny.MultiSet
        d_378_v35_ = _dafny.MultiSet([(self).f32])
        d_379_v36_: _dafny.Array
        nw66_ = _dafny.Array(None, 4)
        nw66_[int(0)] = (self).f31
        nw66_[int(1)] = ((d_375_v33_)[default__.safeIndex(109, (d_375_v33_).length(0))]) <= ((self).f32)
        nw66_[int(2)] = (self).f31
        nw66_[int(3)] = p0
        d_379_v36_ = nw66_
        index56_ = default__.safeIndex(564, (d_379_v36_).length(0))
        (d_379_v36_)[index56_] = False
        d_380_v37_: D3
        d_380_v37_ = D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjaamgms")), (self).f32, (self).f31)
        d_381_v38_: _dafny.Seq
        d_381_v38_ = _dafny.SeqWithoutIsStrInference([d_378_v35_, d_378_v35_])
        index57_ = default__.safeIndex(564, (d_379_v36_).length(0))
        rhs52_ = (d_380_v37_).cf13
        rhs53_ = ((d_381_v38_)[default__.safeIndex(len(d_339_v1_), len(d_381_v38_))]) | ((d_378_v35_).set((self).f32, default__.abs((self).f32)))
        rhs54_ = True
        lhs30_ = d_379_v36_
        lhs31_ = default__.safeIndex(564, (d_379_v36_).length(0))
        r2 = rhs52_
        d_378_v35_ = rhs53_
        lhs30_[lhs31_] = rhs54_
        r0 = (self).f32
        nw67_ = _dafny.Array(_dafny.MultiSet({}), 14)
        r1 = nw67_
        d_382_v39_: _dafny.Seq
        d_382_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kip"))
        r2 = d_382_v39_
        r3 = (0) - (default__.safeDivisionInt((self).f32, (((p2)[(self).f31] if ((self).f31) in (p2) else (self).f32)) + ((d_375_v33_)[default__.safeIndex(109, (d_375_v33_).length(0))])))
        return r0, r1, r2, r3

    def m5(self, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        r3: bool = False
        r0 = (self).f32
        d_383_v0_: _dafny.MultiSet
        d_383_v0_ = _dafny.MultiSet([(self).f32])
        d_384_v1_: _dafny.Seq
        d_384_v1_ = _dafny.SeqWithoutIsStrInference([(d_383_v0_).cardinality, (self).f32])
        d_385_v2_: _dafny.Seq
        d_385_v2_ = _dafny.SeqWithoutIsStrInference([d_384_v1_, (_dafny.SeqWithoutIsStrInference([(self).f32])) + (d_384_v1_), (d_384_v1_) + (_dafny.SeqWithoutIsStrInference([(self).f32 for d_386_i0_ in range(default__.abs(595))]))])
        d_387_v3_: _dafny.Array
        nw68_ = _dafny.Array(int(0), 2)
        d_387_v3_ = nw68_
        index58_ = default__.safeIndex(373, (d_387_v3_).length(0))
        (d_387_v3_)[index58_] = 291
        d_388_v4_: str
        d_388_v4_ = _dafny.CodePoint('o')
        d_389_v5_: _dafny.Seq
        d_389_v5_ = _dafny.SeqWithoutIsStrInference([(self).f31])
        d_390_v6_: _dafny.Set
        d_390_v6_ = _dafny.Set({d_389_v5_, _dafny.SeqWithoutIsStrInference([(self).f31, (self).f31]), d_389_v5_, d_389_v5_, _dafny.SeqWithoutIsStrInference([(self).f31])})
        index59_ = default__.safeIndex(373, (d_387_v3_).length(0))
        rhs55_ = (self).f32
        rhs56_ = d_385_v2_
        rhs57_ = (self).f32
        rhs58_ = default__.fm18((self).fm17(d_388_v4_, globalState), (self).f31, globalState)
        rhs59_ = (len((d_390_v6_) | (d_390_v6_))) * ((self).f32)
        lhs32_ = d_387_v3_
        lhs33_ = default__.safeIndex(373, (d_387_v3_).length(0))
        r0 = rhs55_
        d_385_v2_ = rhs56_
        r0 = rhs57_
        r0 = rhs58_
        lhs32_[lhs33_] = rhs59_
        r1 = (self).f31
        d_391_i1_: int
        d_391_i1_ = 0
        with _dafny.label("5"):
            while (self).f31:
                with _dafny.c_label("5"):
                    if (d_391_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_391_i1_ = (d_391_i1_) + (1)
                    d_392_v7_: _dafny.Map
                    d_392_v7_ = _dafny.Map({(d_387_v3_)[default__.safeIndex(373, (d_387_v3_).length(0))]: d_389_v5_})
                    d_393_v10_: _dafny.Array
                    nw69_ = _dafny.Array(None, 14)
                    nw69_[int(0)] = d_392_v7_
                    nw69_[int(1)] = (d_392_v7_) | (d_392_v7_)
                    nw69_[int(2)] = _dafny.Map({(self).f32: _dafny.SeqWithoutIsStrInference([(self).f31])})
                    nw69_[int(3)] = ((d_392_v7_).set((d_387_v3_)[default__.safeIndex(373, (d_387_v3_).length(0))], d_389_v5_)) | ((_dafny.Map({(self).f32: d_389_v5_})).set((d_387_v3_)[default__.safeIndex(373, (d_387_v3_).length(0))], d_389_v5_))
                    nw69_[int(4)] = _dafny.Map({-777: d_389_v5_})
                    nw69_[int(5)] = d_392_v7_
                    def iife29_():
                        coll13_ = _dafny.Map()
                        compr_13_: int
                        for compr_13_ in _dafny.IntegerRange(189, 37):
                            d_394_v8_: int = compr_13_
                            if ((189) <= (d_394_v8_)) and ((d_394_v8_) < (37)):
                                coll13_[default__.safeModuloInt(d_394_v8_, 58)] = d_389_v5_
                        return _dafny.Map(coll13_)
                    nw69_[int(6)] = iife29_()
                    
                    nw69_[int(7)] = d_392_v7_
                    def iife30_():
                        coll14_ = _dafny.Map()
                        compr_14_: int
                        for compr_14_ in _dafny.IntegerRange(719, 649):
                            d_395_v9_: int = compr_14_
                            if ((719) <= (d_395_v9_)) and ((d_395_v9_) < (649)):
                                coll14_[(d_395_v9_) + ((d_387_v3_)[default__.safeIndex(373, (d_387_v3_).length(0))])] = d_389_v5_
                        return _dafny.Map(coll14_)
                    nw69_[int(8)] = iife30_()
                    
                    nw69_[int(9)] = d_392_v7_
                    nw69_[int(10)] = _dafny.Map({(d_387_v3_)[default__.safeIndex(373, (d_387_v3_).length(0))]: d_389_v5_})
                    nw69_[int(11)] = (d_392_v7_ if (self).f31 else d_392_v7_)
                    nw69_[int(12)] = d_392_v7_
                    nw69_[int(13)] = d_392_v7_
                    d_393_v10_ = nw69_
                    index60_ = default__.safeIndex(778, (d_393_v10_).length(0))
                    (d_393_v10_)[index60_] = d_392_v7_
                    index61_ = default__.safeIndex(778, (d_393_v10_).length(0))
                    (d_393_v10_)[index61_] = d_392_v7_
                    d_384_v1_ = d_384_v1_
                    d_396_v11_: _dafny.Seq
                    d_396_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smv"))
                    d_396_v11_ = d_396_v11_
                    rhs60_ = d_388_v4_
                    rhs61_ = default__.fm18(((self).f31 if (self).f31 else (self).f31), (self).f31, globalState)
                    rhs62_ = d_387_v3_
                    d_388_v4_ = rhs60_
                    r0 = rhs61_
                    d_387_v3_ = rhs62_
                    pass
            pass
        d_397_v12_: _dafny.Seq
        d_397_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxidk"))
        hi2_ = len(d_397_v12_)
        for d_398_i2_ in range((d_387_v3_)[default__.safeIndex(373, (d_387_v3_).length(0))], hi2_):
            d_399_v13_: bool
            d_400_v14_: bool
            d_401_v15_: bool
            out14_: bool
            out15_: bool
            out16_: bool
            out14_, out15_, out16_ = (self).m6((d_398_i2_) == (len(default__.fm19((self).f32, len(d_397_v12_), globalState))), globalState)
            d_399_v13_ = out14_
            d_400_v14_ = out15_
            d_401_v15_ = out16_
            def iife31_(_pat_let8_0):
                def iife32_(d_402_dt__update__tmp_h0_):
                    def iife33_(_pat_let9_0):
                        def iife34_(d_403_dt__update_hcf15_h0_):
                            def iife35_(_pat_let10_0):
                                def iife36_(d_404_dt__update_hcf14_h0_):
                                    return D3_DC10((d_402_dt__update__tmp_h0_).cf13, d_404_dt__update_hcf14_h0_, d_403_dt__update_hcf15_h0_)
                                return iife36_(_pat_let10_0)
                            return iife35_((self).f32)
                        return iife34_(_pat_let9_0)
                    return iife33_(True)
                return iife32_(_pat_let8_0)
            source6_ = iife31_(D3_DC10(d_397_v12_, 377, (self).f31))
            if source6_.is_DC10:
                d_405___mcc_h0_ = source6_.cf13
                d_406___mcc_h1_ = source6_.cf14
                d_407___mcc_h2_ = source6_.cf15
                d_408_cf15_ = d_407___mcc_h2_
                d_409_cf14_ = d_406___mcc_h1_
                d_410_cf13_ = d_405___mcc_h0_
                r0 = d_409_cf14_
                index62_ = default__.safeIndex(373, (d_387_v3_).length(0))
                (d_387_v3_)[index62_] = (self).f32
                d_411_v16_: _dafny.Array
                nw70_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_411_v16_ = nw70_
                index63_ = default__.safeIndex(136, (d_411_v16_).length(0))
                (d_411_v16_)[index63_] = d_387_v3_
                index64_ = default__.safeIndex(136, (d_411_v16_).length(0))
                (d_411_v16_)[index64_] = d_387_v3_
                d_412_v18_: _dafny.Seq
                def iife37_():
                    coll15_ = _dafny.Map()
                    compr_15_: int
                    for compr_15_ in _dafny.IntegerRange(-16, -893):
                        d_413_v17_: int = compr_15_
                        if ((-16) <= (d_413_v17_)) and ((d_413_v17_) < (-893)):
                            coll15_[(d_413_v17_) * (507)] = (self).f31
                    return _dafny.Map(coll15_)
                d_412_v18_ = _dafny.SeqWithoutIsStrInference([iife37_()
                ])
                d_414_v19_: _dafny.MultiSet
                d_414_v19_ = _dafny.MultiSet([d_399_v13_])
                d_415_v20_: _dafny.Set
                d_415_v20_ = _dafny.Set({d_398_i2_, (d_414_v19_).cardinality, (self).f32})
                d_416_v21_: C0
                nw71_ = C0()
                nw71_.ctor__((len(d_412_v18_)) in (d_415_v20_))
                d_416_v21_ = nw71_
            elif source6_.is_DC9:
                d_417___mcc_h3_ = source6_.cf12
                d_418_cf12_ = d_417___mcc_h3_
                r0 = ((d_387_v3_)[default__.safeIndex(373, (d_387_v3_).length(0))] if (d_401_v15_) and (d_401_v15_) else d_398_i2_)
                d_419_v22_: _dafny.Array
                nw72_ = _dafny.Array(False, 12)
                d_419_v22_ = nw72_
                index65_ = default__.safeIndex(865, (d_419_v22_).length(0))
                (d_419_v22_)[index65_] = d_401_v15_
                d_420_v23_: _dafny.Map
                d_420_v23_ = _dafny.Map({d_401_v15_: (d_387_v3_)[default__.safeIndex(373, (d_387_v3_).length(0))]})
                index66_ = default__.safeIndex(373, (d_387_v3_).length(0))
                index67_ = default__.safeIndex(865, (d_419_v22_).length(0))
                index68_ = default__.safeIndex(373, (d_387_v3_).length(0))
                rhs63_ = (self).f32
                rhs64_ = d_399_v13_
                rhs65_ = (d_399_v13_ if (d_400_v14_) not in (d_420_v23_) else (self).f31)
                rhs66_ = (self).f32
                rhs67_ = (self).f31
                lhs34_ = d_387_v3_
                lhs35_ = default__.safeIndex(373, (d_387_v3_).length(0))
                lhs36_ = globalState
                lhs37_ = d_419_v22_
                lhs38_ = default__.safeIndex(865, (d_419_v22_).length(0))
                lhs39_ = d_387_v3_
                lhs40_ = default__.safeIndex(373, (d_387_v3_).length(0))
                lhs34_[lhs35_] = rhs63_
                lhs36_.f16 = rhs64_
                lhs37_[lhs38_] = rhs65_
                lhs39_[lhs40_] = rhs66_
                d_399_v13_ = rhs67_
                d_421_v24_: _dafny.Array
                nw73_ = _dafny.Array(None, 6)
                nw73_[int(0)] = d_389_v5_
                nw73_[int(1)] = d_389_v5_
                nw73_[int(2)] = d_389_v5_
                nw73_[int(3)] = d_389_v5_
                nw73_[int(4)] = _dafny.SeqWithoutIsStrInference([d_399_v13_])
                nw73_[int(5)] = d_389_v5_
                d_421_v24_ = nw73_
                index69_ = default__.safeIndex(295, (d_421_v24_).length(0))
                (d_421_v24_)[index69_] = (d_389_v5_) + (d_389_v5_)
                index70_ = default__.safeIndex(295, (d_421_v24_).length(0))
                rhs68_ = (d_387_v3_)[default__.safeIndex(373, (d_387_v3_).length(0))]
                rhs69_ = d_389_v5_
                rhs70_ = d_400_v14_
                rhs71_ = d_401_v15_
                rhs72_ = default__.safeDivisionInt((self).f32, d_398_i2_)
                lhs41_ = d_421_v24_
                lhs42_ = default__.safeIndex(295, (d_421_v24_).length(0))
                lhs43_ = globalState
                r0 = rhs68_
                lhs41_[lhs42_] = rhs69_
                lhs43_.f16 = rhs70_
                d_399_v13_ = rhs71_
                r0 = rhs72_
                d_422_v25_: bool
                d_423_v26_: bool
                d_424_v27_: bool
                out17_: bool
                out18_: bool
                out19_: bool
                out17_, out18_, out19_ = (self).m6(d_400_v14_, globalState)
                d_422_v25_ = out17_
                d_423_v26_ = out18_
                d_424_v27_ = out19_
            elif True:
                d_425___mcc_h4_ = source6_.cf16
                d_426_cf16_ = d_425___mcc_h4_
                d_427_v28_: _dafny.Set
                d_427_v28_ = _dafny.Set({(d_387_v3_)[default__.safeIndex(373, (d_387_v3_).length(0))]})
                d_428_v29_: _dafny.Array
                nw74_ = _dafny.Array(False, 20)
                d_428_v29_ = nw74_
                d_429_v30_: D5
                d_429_v30_ = D5_DC14(d_427_v28_, d_398_i2_, True, d_428_v29_)
                r0 = (d_429_v30_).cf20
                index71_ = default__.safeIndex(373, (d_387_v3_).length(0))
                (d_387_v3_)[index71_] = d_398_i2_
                index72_ = default__.safeIndex(373, (d_387_v3_).length(0))
                (d_387_v3_)[index72_] = (0) - ((self).f32)
                d_430_v31_: C0
                nw75_ = C0()
                nw75_.ctor__(not(d_399_v13_))
                d_430_v31_ = nw75_
            index73_ = default__.safeIndex(373, (d_387_v3_).length(0))
            (d_387_v3_)[index73_] = (d_387_v3_)[default__.safeIndex(373, (d_387_v3_).length(0))]
            if d_399_v13_:
                d_397_v12_ = _dafny.SeqWithoutIsStrInference([d_388_v4_ for d_431_i3_ in range(default__.abs(468))])
                r0 = (self).f32
                d_432_v32_: _dafny.Set
                d_432_v32_ = _dafny.Set({d_388_v4_})
                d_432_v32_ = d_432_v32_
                d_433_v33_: D0
                d_433_v33_ = D0_DC1(d_388_v4_, (self).f31, False)
                pat_let_tv5_ = d_388_v4_
                pat_let_tv6_ = d_397_v12_
                pat_let_tv7_ = d_387_v3_
                pat_let_tv8_ = d_387_v3_
                def iife38_(_pat_let11_0):
                    def iife39_(d_434_dt__update__tmp_h1_):
                        def iife40_(_pat_let12_0):
                            def iife41_(d_435_dt__update_hcf1_h0_):
                                def iife42_(_pat_let13_0):
                                    def iife43_(d_436_dt__update_hcf3_h0_):
                                        return D0_DC1(d_435_dt__update_hcf1_h0_, (d_434_dt__update__tmp_h1_).cf2, d_436_dt__update_hcf3_h0_)
                                    return iife43_(_pat_let13_0)
                                return iife42_((len(pat_let_tv6_)) > ((pat_let_tv8_)[default__.safeIndex(373, (pat_let_tv7_).length(0))]))
                            return iife41_(_pat_let12_0)
                        return iife40_(pat_let_tv5_)
                    return iife39_(_pat_let11_0)
                d_433_v33_ = iife38_(d_433_v33_)
                d_437_v34_: C0
                nw76_ = C0()
                nw76_.ctor__(d_401_v15_)
                d_437_v34_ = nw76_
            elif True:
                d_438_v35_: _dafny.Map
                d_438_v35_ = _dafny.Map({(d_383_v0_).cardinality: (self).f31})
                d_439_v36_: _dafny.Array
                def lambda11_(d_440_i2_):
                    def lambda12_(d_441_i4_):
                        return _dafny.Set({d_440_i2_, d_440_i2_})

                    return lambda12_

                init6_ = lambda11_(d_398_i2_)
                nw77_ = _dafny.Array(None, 14)
                for i0_6_ in range(nw77_.length(0)):
                    nw77_[i0_6_] = init6_(i0_6_)
                d_439_v36_ = nw77_
                d_442_v37_: D6
                d_442_v37_ = D6_DC17(d_401_v15_, d_439_v36_, d_398_i2_)
                d_443_v38_: _dafny.Map
                d_443_v38_ = _dafny.Map({d_438_v35_: _dafny.SeqWithoutIsStrInference([(d_389_v5_)[default__.safeIndex((self).f32, len(d_389_v5_))], (d_442_v37_).cf25, d_399_v13_])})
                d_444_v39_: _dafny.Array
                def lambda13_(d_445_i5_):
                    return _dafny.CodePoint('r')

                init7_ = lambda13_
                nw78_ = _dafny.Array(None, 5)
                for i0_7_ in range(nw78_.length(0)):
                    nw78_[i0_7_] = init7_(i0_7_)
                d_444_v39_ = nw78_
                index74_ = default__.safeIndex(392, (d_444_v39_).length(0))
                (d_444_v39_)[index74_] = d_388_v4_
                d_446_v40_: D0
                d_446_v40_ = D0_DC2((self).f31)
                d_447_v41_: _dafny.Map
                d_447_v41_ = _dafny.Map({d_446_v40_: d_388_v4_})
                index75_ = default__.safeIndex(392, (d_444_v39_).length(0))
                rhs73_ = default__.fm20((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gswdyire"))) + (_dafny.SeqWithoutIsStrInference([d_388_v4_ for d_448_i6_ in range(default__.abs(564))])), d_383_v0_, (self).f31, (self).f32, globalState)
                rhs74_ = (d_400_v14_) == (not (True) or (d_399_v13_))
                rhs75_ = (d_443_v38_).set(d_438_v35_, d_389_v5_)
                rhs76_ = ((d_447_v41_)[d_446_v40_] if (d_446_v40_) in (d_447_v41_) else _dafny.CodePoint('o'))
                lhs44_ = d_444_v39_
                lhs45_ = default__.safeIndex(392, (d_444_v39_).length(0))
                d_388_v4_ = rhs73_
                d_399_v13_ = rhs74_
                d_443_v38_ = rhs75_
                lhs44_[lhs45_] = rhs76_
                rhs77_ = d_388_v4_
                rhs78_ = d_399_v13_
                lhs46_ = globalState
                d_388_v4_ = rhs77_
                lhs46_.f16 = rhs78_
                pat_let_tv9_ = d_439_v36_
                def iife44_(_pat_let14_0):
                    def iife45_(d_449_dt__update__tmp_h2_):
                        def iife46_(_pat_let15_0):
                            def iife47_(d_450_dt__update_hcf26_h0_):
                                return D6_DC17((d_449_dt__update__tmp_h2_).cf25, d_450_dt__update_hcf26_h0_, (d_449_dt__update__tmp_h2_).cf27)
                            return iife47_(_pat_let15_0)
                        return iife46_(pat_let_tv9_)
                    return iife45_(_pat_let14_0)
                d_439_v36_ = ((iife44_(d_442_v37_) if True else d_442_v37_)).cf26
                d_451_v42_: D3
                d_451_v42_ = D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpccqjlr")), len((d_397_v12_).set(default__.safeIndex(847, len(d_397_v12_)), (d_444_v39_)[default__.safeIndex(392, (d_444_v39_).length(0))])), d_399_v13_)
                d_397_v12_ = (d_451_v42_).cf13
                index76_ = default__.safeIndex(373, (d_387_v3_).length(0))
                rhs79_ = default__.fm21(d_383_v0_, d_400_v14_, ((d_397_v12_).set(default__.safeIndex((self).f32, len(d_397_v12_)), default__.fm20((self).fm16(d_384_v1_, d_401_v15_, d_399_v13_, d_400_v14_, globalState), d_383_v0_, (self).f31, (d_383_v0_).cardinality, globalState))) != (d_397_v12_), globalState)
                rhs80_ = d_398_i2_
                lhs47_ = d_387_v3_
                lhs48_ = default__.safeIndex(373, (d_387_v3_).length(0))
                d_438_v35_ = rhs79_
                lhs47_[lhs48_] = rhs80_
        (globalState).f28 = ((d_387_v3_)[default__.safeIndex(373, (d_387_v3_).length(0))]) == (((self).f32) - (141))
        r0 = (self).f32
        r1 = ((self).f32) == (default__.fm18((self).f31, not((self).f31), globalState))
        d_452_v43_: _dafny.Set
        d_452_v43_ = _dafny.Set({(self).f32, len(d_397_v12_)})
        d_453_v44_: _dafny.MultiSet
        d_453_v44_ = _dafny.MultiSet([d_452_v43_])
        d_454_v45_: _dafny.Map
        d_454_v45_ = _dafny.Map({(0) - ((self).f32): (self).f31})
        d_455_v46_: _dafny.Map
        d_455_v46_ = _dafny.Map({(self).f31: d_454_v45_})
        r2 = ((_dafny.MultiSet(default__.fm22((self).f32, globalState))) - (d_453_v44_) if (self).f31 else (d_453_v44_) | ((d_453_v44_).set(d_452_v43_, default__.abs(len(default__.fm23(d_455_v46_, _dafny.Map({False: len(d_397_v12_)}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqta")), globalState))))))
        r3 = (self).f31
        return r0, r1, r2, r3

    def m6(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        r2: bool = False
        d_456_v0_: _dafny.Array
        nw79_ = _dafny.Array(None, 6)
        nw79_[int(0)] = (self).f32
        nw79_[int(1)] = (self).f32
        nw79_[int(2)] = (self).f32
        nw79_[int(3)] = (self).f32
        nw79_[int(4)] = (self).f32
        nw79_[int(5)] = (self).f32
        d_456_v0_ = nw79_
        (globalState).f24 = d_456_v0_
        d_457_v1_: _dafny.Seq
        d_457_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jymesegv"))
        d_458_v2_: _dafny.MultiSet
        d_458_v2_ = _dafny.MultiSet([(self).f32])
        d_459_v3_: _dafny.Map
        d_459_v3_ = _dafny.Map({(0) - ((self).f32): p0})
        d_460_v4_: _dafny.Seq
        d_460_v4_ = _dafny.SeqWithoutIsStrInference([default__.fm21(d_458_v2_, (self).f31, (self).f31, globalState), d_459_v3_, d_459_v3_])
        d_461_v5_: D3
        d_461_v5_ = D3_DC10((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_462_i0_ in range(default__.abs(162))])) + (d_457_v1_), (_dafny.MultiSet(d_460_v4_)).cardinality, (self).f31)
        source7_ = d_461_v5_
        if source7_.is_DC10:
            d_463___mcc_h0_ = source7_.cf13
            d_464___mcc_h1_ = source7_.cf14
            d_465___mcc_h2_ = source7_.cf15
            d_466_cf15_ = d_465___mcc_h2_
            d_467_cf14_ = d_464___mcc_h1_
            d_468_cf13_ = d_463___mcc_h0_
            d_467_cf14_ = d_467_cf14_
            d_469_v6_: _dafny.Seq
            d_469_v6_ = _dafny.SeqWithoutIsStrInference([d_466_cf15_, (self).f31])
            d_470_v7_: _dafny.Array
            nw80_ = _dafny.Array(None, 17)
            nw80_[int(0)] = d_469_v6_
            nw80_[int(1)] = d_469_v6_
            nw80_[int(2)] = d_469_v6_
            nw80_[int(3)] = d_469_v6_
            nw80_[int(4)] = d_469_v6_
            nw80_[int(5)] = d_469_v6_
            nw80_[int(6)] = d_469_v6_
            nw80_[int(7)] = d_469_v6_
            nw80_[int(8)] = d_469_v6_
            nw80_[int(9)] = _dafny.SeqWithoutIsStrInference([d_466_cf15_, False])
            nw80_[int(10)] = d_469_v6_
            nw80_[int(11)] = d_469_v6_
            nw80_[int(12)] = d_469_v6_
            nw80_[int(13)] = d_469_v6_
            nw80_[int(14)] = d_469_v6_
            nw80_[int(15)] = d_469_v6_
            nw80_[int(16)] = _dafny.SeqWithoutIsStrInference([p0])
            d_470_v7_ = nw80_
            d_471_v8_: D0
            d_471_v8_ = D0_DC0(d_470_v7_)
            d_472_v9_: D0
            d_472_v9_ = D0_DC4(d_471_v8_)
            d_473_v10_: _dafny.Array
            nw81_ = _dafny.Array(None, 5)
            nw81_[int(0)] = d_472_v9_
            nw81_[int(1)] = d_472_v9_
            nw81_[int(2)] = d_472_v9_
            nw81_[int(3)] = d_472_v9_
            nw81_[int(4)] = d_472_v9_
            d_473_v10_ = nw81_
            d_473_v10_ = d_473_v10_
            d_474_v11_: D0
            d_474_v11_ = D0_DC2((self).f31)
            source8_ = d_474_v11_
            if source8_.is_DC1:
                d_475___mcc_h5_ = source8_.cf1
                d_476___mcc_h6_ = source8_.cf2
                d_477___mcc_h7_ = source8_.cf3
                d_478_cf3_ = d_477___mcc_h7_
                d_479_cf2_ = d_476___mcc_h6_
                d_480_cf1_ = d_475___mcc_h5_
                d_481_v12_: _dafny.Map
                d_481_v12_ = _dafny.Map({d_480_cf1_: (self).fm17(d_480_cf1_, globalState)})
                d_481_v12_ = d_481_v12_
                d_482_v13_: _dafny.Array
                nw82_ = _dafny.Array(D3.default()(), 11)
                d_482_v13_ = nw82_
                d_483_v14_: _dafny.Seq
                d_483_v14_ = _dafny.SeqWithoutIsStrInference([d_482_v13_])
                d_484_v15_: _dafny.Array
                nw83_ = _dafny.Array(None, 4)
                nw83_[int(0)] = ((_dafny.SeqWithoutIsStrInference([d_482_v13_])) + (_dafny.SeqWithoutIsStrInference([d_482_v13_, d_482_v13_]))).set(default__.safeIndex(d_467_cf14_, len((_dafny.SeqWithoutIsStrInference([d_482_v13_])) + (_dafny.SeqWithoutIsStrInference([d_482_v13_, d_482_v13_])))), d_482_v13_)
                nw83_[int(1)] = d_483_v14_
                nw83_[int(2)] = d_483_v14_
                nw83_[int(3)] = d_483_v14_
                d_484_v15_ = nw83_
                index77_ = default__.safeIndex(97, (d_484_v15_).length(0))
                (d_484_v15_)[index77_] = d_483_v14_
                index78_ = default__.safeIndex(97, (d_484_v15_).length(0))
                (d_484_v15_)[index78_] = d_483_v14_
                d_485_v16_: _dafny.Map
                d_485_v16_ = _dafny.Map({512: len(d_457_v1_)})
                d_486_v17_: _dafny.Set
                d_486_v17_ = _dafny.Set({d_485_v16_})
                d_487_v18_: _dafny.Map
                d_487_v18_ = _dafny.Map({(len(d_468_cf13_)) == ((self).f32): (d_486_v17_) | (d_486_v17_)})
                d_488_v21_: _dafny.Seq
                def iife48_():
                    def iife50_():
                        coll18_ = _dafny.Set()
                        compr_18_: _dafny.Map
                        for compr_18_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_467_cf14_: len(d_468_cf13_)}) for d_489_i1_ in range(default__.abs(682))])).Elements:
                            d_492_v19_: _dafny.Map = compr_18_
                            if (d_492_v19_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_467_cf14_: len(d_468_cf13_)}) for d_489_i1_ in range(default__.abs(682))])):
                                coll18_ = coll18_.union(_dafny.Set([d_492_v19_]))
                        return _dafny.Set(coll18_)
                    coll16_ = _dafny.Set()
                    def iife49_():
                        coll17_ = _dafny.Set()
                        compr_17_: _dafny.Map
                        for compr_17_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_467_cf14_: len(d_468_cf13_)}) for d_489_i1_ in range(default__.abs(682))])).Elements:
                            d_490_v19_: _dafny.Map = compr_17_
                            if (d_490_v19_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_467_cf14_: len(d_468_cf13_)}) for d_489_i1_ in range(default__.abs(682))])):
                                coll17_ = coll17_.union(_dafny.Set([d_490_v19_]))
                        return _dafny.Set(coll17_)
                    compr_16_: _dafny.Map
                    for compr_16_ in (iife49_()
                    ).Elements:
                        d_491_v20_: _dafny.Map = compr_16_
                        if (d_491_v20_) in (iife50_()
                        ):
                            coll16_ = coll16_.union(_dafny.Set([d_491_v20_]))
                    return _dafny.Set(coll16_)
                d_488_v21_ = _dafny.SeqWithoutIsStrInference([iife48_()
                , ((d_487_v18_)[p0] if (p0) in (d_487_v18_) else d_486_v17_)])
                d_487_v18_ = (d_487_v18_).set(p0, (d_488_v21_)[default__.safeIndex((self).f32, len(d_488_v21_))])
                d_468_cf13_ = d_468_cf13_
            elif source8_.is_DC2:
                d_493___mcc_h8_ = source8_.cf4
                d_494_cf4_ = d_493___mcc_h8_
                d_495_v22_: str
                d_495_v22_ = _dafny.CodePoint('w')
                d_496_v23_: _dafny.MultiSet
                d_496_v23_ = _dafny.MultiSet([d_495_v22_, default__.fm20(d_457_v1_, _dafny.MultiSet([(self).f32]), (self).f31, (self).f32, globalState)])
                default__.m0(((d_496_v23_)[d_495_v22_] if (d_495_v22_) in (d_496_v23_) else (self).f32), globalState)
                d_497_v24_: _dafny.Map
                d_497_v24_ = _dafny.Map({(self).f31: d_467_cf14_})
                d_498_v25_: _dafny.Map
                d_498_v25_ = _dafny.Map({((d_497_v24_)[p0] if (p0) in (d_497_v24_) else (self).f32): default__.fm18(p0, (self).f31, globalState)})
                d_498_v25_ = (d_498_v25_).set((len(d_457_v1_)) + (d_467_cf14_), d_467_cf14_)
                d_499_v26_: _dafny.MultiSet
                d_499_v26_ = _dafny.MultiSet([(D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hiecrh")), (self).f32, d_494_cf4_)).cf13])
                d_500_v27_: _dafny.Set
                d_500_v27_ = _dafny.Set({d_494_cf4_})
                d_501_v28_: _dafny.Map
                d_501_v28_ = _dafny.Map({d_499_v26_: (d_500_v27_).intersection(d_500_v27_)})
                d_501_v28_ = (d_501_v28_).set(_dafny.MultiSet([d_457_v1_, _dafny.SeqWithoutIsStrInference([d_495_v22_ for d_502_i2_ in range(default__.abs(735))])]), d_500_v27_)
                d_503_v29_: C0
                nw84_ = C0()
                nw84_.ctor__(((self).f31) == ((self).fm17(d_495_v22_, globalState)))
                d_503_v29_ = nw84_
            elif source8_.is_DC3:
                d_504___mcc_h9_ = source8_.cf5
                d_505_cf5_ = d_504___mcc_h9_
                d_506_v30_: _dafny.Map
                d_506_v30_ = _dafny.Map({(len(d_468_cf13_)) - ((self).f32): (d_458_v2_).cardinality})
                d_506_v30_ = d_506_v30_
                d_507_v31_: _dafny.Array
                nw85_ = _dafny.Array(None, 9)
                nw85_[int(0)] = not (not((self).f31)) or ((self).f31)
                nw85_[int(1)] = not (False) or ((self).f31)
                nw85_[int(2)] = not(p0)
                nw85_[int(3)] = (self).f31
                nw85_[int(4)] = (p0) == (d_466_cf15_)
                nw85_[int(5)] = ((0) - ((self).f32)) > (d_467_cf14_)
                nw85_[int(6)] = p0
                nw85_[int(7)] = (self).f31
                nw85_[int(8)] = not(d_466_cf15_)
                d_507_v31_ = nw85_
                index79_ = default__.safeIndex(865, (d_507_v31_).length(0))
                (d_507_v31_)[index79_] = d_466_cf15_
                d_508_v32_: _dafny.Map
                d_508_v32_ = _dafny.Map({d_472_v9_: d_467_cf14_})
                index80_ = default__.safeIndex(865, (d_507_v31_).length(0))
                (d_507_v31_)[index80_] = (d_469_v6_)[default__.safeIndex(((d_508_v32_)[d_472_v9_] if (d_472_v9_) in (d_508_v32_) else 759), len(d_469_v6_))]
                index81_ = default__.safeIndex(395, (d_456_v0_).length(0))
                (d_456_v0_)[index81_] = len((d_468_cf13_) + (d_468_cf13_))
                d_509_v33_: _dafny.Seq
                d_509_v33_ = _dafny.SeqWithoutIsStrInference([d_468_cf13_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufldtem")), d_468_cf13_])
                d_510_v34_: _dafny.Array
                nw86_ = _dafny.Array(None, 21)
                nw86_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eohrqs"))
                nw86_[int(1)] = d_457_v1_
                nw86_[int(2)] = d_457_v1_
                nw86_[int(3)] = d_457_v1_
                nw86_[int(4)] = d_468_cf13_
                nw86_[int(5)] = ((d_461_v5_).cf13) + (d_468_cf13_)
                nw86_[int(6)] = d_457_v1_
                nw86_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vet"))
                nw86_[int(8)] = (d_457_v1_) + (d_457_v1_)
                nw86_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "et"))
                nw86_[int(10)] = d_468_cf13_
                nw86_[int(11)] = (d_509_v33_)[default__.safeIndex(407, len(d_509_v33_))]
                nw86_[int(12)] = d_468_cf13_
                nw86_[int(13)] = d_457_v1_
                nw86_[int(14)] = d_457_v1_
                nw86_[int(15)] = (d_457_v1_) + (d_457_v1_)
                nw86_[int(16)] = (d_468_cf13_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvmafgib")))
                nw86_[int(17)] = d_468_cf13_
                nw86_[int(18)] = d_457_v1_
                nw86_[int(19)] = (d_461_v5_).cf13
                nw86_[int(20)] = _dafny.SeqWithoutIsStrInference([d_505_cf5_ for d_511_i3_ in range(default__.abs(902))])
                d_510_v34_ = nw86_
                index82_ = default__.safeIndex(44, (d_510_v34_).length(0))
                (d_510_v34_)[index82_] = d_457_v1_
                index83_ = default__.safeIndex(395, (d_456_v0_).length(0))
                index84_ = default__.safeIndex(44, (d_510_v34_).length(0))
                rhs81_ = (self).f32
                rhs82_ = (self).f32
                rhs83_ = (d_467_cf14_) * (-720)
                rhs84_ = d_468_cf13_
                lhs49_ = d_456_v0_
                lhs50_ = default__.safeIndex(395, (d_456_v0_).length(0))
                lhs51_ = d_510_v34_
                lhs52_ = default__.safeIndex(44, (d_510_v34_).length(0))
                d_467_cf14_ = rhs81_
                d_467_cf14_ = rhs82_
                lhs49_[lhs50_] = rhs83_
                lhs51_[lhs52_] = rhs84_
                d_512_v35_: D2
                d_512_v35_ = D2_DC6((self).f32)
                d_513_v36_: _dafny.Set
                d_513_v36_ = _dafny.Set({d_512_v35_})
                d_514_v37_: _dafny.Map
                d_514_v37_ = _dafny.Map({d_467_cf14_: d_505_cf5_})
                d_515_v38_: _dafny.Array
                nw87_ = _dafny.Array(None, 13)
                nw87_[int(0)] = d_505_cf5_
                nw87_[int(1)] = _dafny.CodePoint('u')
                nw87_[int(2)] = d_505_cf5_
                nw87_[int(3)] = d_505_cf5_
                nw87_[int(4)] = ((d_514_v37_)[(d_456_v0_)[default__.safeIndex(395, (d_456_v0_).length(0))]] if ((d_456_v0_)[default__.safeIndex(395, (d_456_v0_).length(0))]) in (d_514_v37_) else d_505_cf5_)
                nw87_[int(5)] = d_505_cf5_
                nw87_[int(6)] = d_505_cf5_
                nw87_[int(7)] = d_505_cf5_
                nw87_[int(8)] = d_505_cf5_
                nw87_[int(9)] = d_505_cf5_
                nw87_[int(10)] = _dafny.CodePoint('v')
                nw87_[int(11)] = d_505_cf5_
                nw87_[int(12)] = d_505_cf5_
                d_515_v38_ = nw87_
                d_516_v39_: _dafny.Map
                d_516_v39_ = _dafny.Map({d_513_v36_: d_515_v38_})
                d_516_v39_ = (d_516_v39_).set(d_513_v36_, d_515_v38_)
            elif source8_.is_DC0:
                d_517___mcc_h10_ = source8_.cf0
                d_518_cf0_ = d_517___mcc_h10_
                d_467_cf14_ = (self).f32
                d_467_cf14_ = d_467_cf14_
                d_519_v40_: _dafny.Map
                d_519_v40_ = _dafny.Map({d_466_cf15_: p0})
                d_520_v41_: _dafny.Map
                d_520_v41_ = _dafny.Map({len((d_519_v40_).set((self).f31, ((d_519_v40_)[(D0_DC2((self).f31)).cf4] if ((D0_DC2((self).f31)).cf4) in (d_519_v40_) else p0))): default__.fm24(d_466_cf15_, d_466_cf15_, 681, globalState)})
                d_521_v42_: _dafny.Map
                d_521_v42_ = _dafny.Map({(self).f31: (((d_520_v41_)[default__.fm18(d_466_cf15_, d_466_cf15_, globalState)] if (default__.fm18(d_466_cf15_, d_466_cf15_, globalState)) in (d_520_v41_) else d_458_v2_)).cardinality})
                d_521_v42_ = (d_521_v42_).set(p0, (self).f32)
                d_522_v43_: _dafny.Array
                nw88_ = _dafny.Array(None, 16)
                d_522_v43_ = nw88_
                d_523_v44_: C0
                nw89_ = C0()
                nw89_.ctor__(p0)
                d_523_v44_ = nw89_
                index85_ = default__.safeIndex(339, (d_522_v43_).length(0))
                (d_522_v43_)[index85_] = d_523_v44_
                index86_ = default__.safeIndex(339, (d_522_v43_).length(0))
                (d_522_v43_)[index86_] = d_523_v44_
            elif True:
                d_524___mcc_h11_ = source8_.cf6
                d_525_cf6_ = d_524___mcc_h11_
                d_526_v45_: _dafny.MultiSet
                d_526_v45_ = _dafny.MultiSet([p0])
                d_527_v46_: _dafny.Map
                d_527_v46_ = _dafny.Map({(d_469_v6_)[default__.safeIndex((d_526_v45_).cardinality, len(d_469_v6_))]: (self).f31})
                d_528_v47_: _dafny.Array
                nw90_ = _dafny.Array(None, 4)
                nw90_[int(0)] = d_466_cf15_
                nw90_[int(1)] = d_466_cf15_
                nw90_[int(2)] = ((d_527_v46_)[True] if (True) in (d_527_v46_) else (d_474_v11_).cf4)
                nw90_[int(3)] = (self).f31
                d_528_v47_ = nw90_
                d_529_v48_: _dafny.Map
                d_529_v48_ = _dafny.Map({d_528_v47_: (693) <= (default__.fm18((d_469_v6_)[default__.safeIndex(494, len(d_469_v6_))], d_466_cf15_, globalState))})
                d_529_v48_ = (d_529_v48_).set(d_528_v47_, d_466_cf15_)
                d_530_v49_: _dafny.Array
                nw91_ = _dafny.Array(_dafny.Map({}), 3)
                d_530_v49_ = nw91_
                d_531_v50_: _dafny.Seq
                d_531_v50_ = _dafny.SeqWithoutIsStrInference([d_530_v49_])
                d_532_v51_: _dafny.Map
                d_532_v51_ = _dafny.Map({default__.fm20(d_457_v1_, d_458_v2_, not(True), (self).f32, globalState): -279})
                d_533_v52_: str
                d_533_v52_ = _dafny.CodePoint('u')
                rhs85_ = (d_531_v50_)[default__.safeIndex(len(d_532_v51_), len(d_531_v50_))]
                rhs86_ = d_527_v46_
                rhs87_ = ((d_526_v45_) - (d_526_v45_)).ispropersubset((d_526_v45_) - (d_526_v45_))
                rhs88_ = (self).fm17(d_533_v52_, globalState)
                lhs53_ = globalState
                d_530_v49_ = rhs85_
                d_527_v46_ = rhs86_
                lhs53_.f28 = rhs87_
                r0 = rhs88_
                d_467_cf14_ = -842
                d_466_cf15_ = (d_461_v5_).cf15
            d_534_v53_: _dafny.Seq
            d_534_v53_ = _dafny.SeqWithoutIsStrInference([d_470_v7_, d_470_v7_])
            d_535_v54_: D0
            d_535_v54_ = D0_DC0((d_534_v53_)[default__.safeIndex(d_467_cf14_, len(d_534_v53_))])
            d_536_v55_: _dafny.Map
            d_536_v55_ = _dafny.Map({d_535_v54_: (self).f31})
            d_536_v55_ = d_536_v55_
        elif source7_.is_DC9:
            d_537___mcc_h3_ = source7_.cf12
            d_538_cf12_ = d_537___mcc_h3_
            default__.m0(((self).f32) + ((0) - ((self).f32)), globalState)
            d_539_v56_: _dafny.Set
            d_539_v56_ = default__.fm25(False, -932, (self).f32, (self).f31, globalState)
            d_540_v57_: _dafny.MultiSet
            d_540_v57_ = _dafny.MultiSet([(self).f31, (self).f31, p0])
            d_541_v58_: D2
            d_541_v58_ = D2_DC6((d_540_v57_).cardinality)
            d_542_v59_: _dafny.Set
            d_542_v59_ = _dafny.Set({d_541_v58_, d_541_v58_, d_541_v58_, d_541_v58_})
            d_539_v56_ = d_542_v59_
            (globalState).f16 = p0
            d_543_v60_: _dafny.Map
            d_543_v60_ = _dafny.Map({p0: (self).f32})
            d_544_v61_: _dafny.Seq
            d_544_v61_ = _dafny.SeqWithoutIsStrInference([(self).f32, len(d_543_v60_), (self).f32, (self).f32, (self).f32])
            d_545_v62_: _dafny.Set
            d_545_v62_ = _dafny.Set({(self).f32, len(d_544_v61_), (self).f32, (self).f32, (d_458_v2_).cardinality})
            d_545_v62_ = (d_545_v62_) - (_dafny.Set({(self).f32, (self).f32}))
        elif True:
            d_546___mcc_h4_ = source7_.cf16
            d_547_cf16_ = d_546___mcc_h4_
            d_548_v63_: int
            d_548_v63_ = 592
            d_548_v63_ = d_548_v63_
            d_549_v64_: _dafny.Seq
            d_549_v64_ = _dafny.SeqWithoutIsStrInference([p0])
            d_550_v65_: _dafny.Set
            d_550_v65_ = _dafny.Set({len(d_549_v64_)})
            d_551_v66_: _dafny.Map
            d_551_v66_ = _dafny.Map({d_456_v0_: len((d_550_v65_).intersection(d_550_v65_))})
            d_552_v67_: _dafny.Array
            nw92_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
            d_552_v67_ = nw92_
            d_553_v68_: str
            d_553_v68_ = _dafny.CodePoint('d')
            index87_ = default__.safeIndex(803, (d_552_v67_).length(0))
            (d_552_v67_)[index87_] = _dafny.SeqWithoutIsStrInference([d_553_v68_])
            d_554_v69_: _dafny.Seq
            d_554_v69_ = _dafny.SeqWithoutIsStrInference([(self).f32, (self).f32])
            index88_ = default__.safeIndex(803, (d_552_v67_).length(0))
            rhs89_ = ((d_551_v66_) | (d_551_v66_)) | (((d_551_v66_).set(d_456_v0_, (self).f32)) | (d_551_v66_))
            rhs90_ = (d_548_v63_) <= ((d_554_v69_)[default__.safeIndex(873, len(d_554_v69_))])
            rhs91_ = d_457_v1_
            rhs92_ = p0
            lhs54_ = d_552_v67_
            lhs55_ = default__.safeIndex(803, (d_552_v67_).length(0))
            d_551_v66_ = rhs89_
            r2 = rhs90_
            lhs54_[lhs55_] = rhs91_
            r2 = rhs92_
            d_555_v70_: _dafny.Map
            d_555_v70_ = _dafny.Map({(self).f31: (0) - ((self).f32)})
            index89_ = default__.safeIndex(531, (d_456_v0_).length(0))
            (d_456_v0_)[index89_] = (d_548_v63_ if (self).f31 else len(d_555_v70_))
            index90_ = default__.safeIndex(531, (d_456_v0_).length(0))
            (d_456_v0_)[index90_] = (self).f32
            default__.m0((d_456_v0_)[default__.safeIndex(531, (d_456_v0_).length(0))], globalState)
        d_556_v71_: _dafny.Set
        d_556_v71_ = _dafny.Set({d_456_v0_, d_456_v0_})
        d_557_v72_: _dafny.Map
        d_557_v72_ = _dafny.Map({d_556_v71_: (self).f32})
        d_557_v72_ = (d_557_v72_).set(d_556_v71_, (self).f32)
        d_558_v73_: int
        d_558_v73_ = 355
        d_558_v73_ = (d_558_v73_) - (len(_dafny.SeqWithoutIsStrInference([p0, p0])))
        d_559_v74_: str
        d_559_v74_ = _dafny.CodePoint('b')
        d_560_i4_: int
        d_560_i4_ = 0
        with _dafny.label("6"):
            while not(not (p0) or ((self).fm17(d_559_v74_, globalState))):
                with _dafny.c_label("6"):
                    if (d_560_i4_) >= (100):
                        raise _dafny.Break("6")
                    d_560_i4_ = (d_560_i4_) + (1)
                    d_457_v1_ = d_457_v1_
                    d_561_v75_: _dafny.Array
                    def lambda14_(d_562_v74_):
                        def lambda15_(d_563_i5_):
                            return _dafny.SeqWithoutIsStrInference([d_562_v74_ for d_564_i6_ in range(default__.abs(-543))])

                        return lambda15_

                    init8_ = lambda14_(d_559_v74_)
                    nw93_ = _dafny.Array(None, 1)
                    for i0_8_ in range(nw93_.length(0)):
                        nw93_[i0_8_] = init8_(i0_8_)
                    d_561_v75_ = nw93_
                    d_565_v76_: _dafny.Array
                    nw94_ = _dafny.Array(None, 6)
                    nw94_[int(0)] = d_561_v75_
                    nw94_[int(1)] = d_561_v75_
                    nw94_[int(2)] = d_561_v75_
                    nw94_[int(3)] = d_561_v75_
                    nw94_[int(4)] = d_561_v75_
                    nw94_[int(5)] = d_561_v75_
                    d_565_v76_ = nw94_
                    index91_ = default__.safeIndex(243, (d_565_v76_).length(0))
                    (d_565_v76_)[index91_] = d_561_v75_
                    index92_ = default__.safeIndex(243, (d_565_v76_).length(0))
                    (d_565_v76_)[index92_] = d_561_v75_
                    d_566_v77_: C0
                    nw95_ = C0()
                    nw95_.ctor__((d_457_v1_) == (d_457_v1_))
                    d_566_v77_ = nw95_
                    rhs93_ = (self).f32
                    rhs94_ = not((self).f31)
                    lhs56_ = globalState
                    d_558_v73_ = rhs93_
                    lhs56_.f16 = rhs94_
                    pass
            pass
        d_567_v78_: _dafny.Seq
        d_567_v78_ = _dafny.SeqWithoutIsStrInference([d_558_v73_])
        d_568_v79_: _dafny.Seq
        d_568_v79_ = _dafny.SeqWithoutIsStrInference([p0, (d_567_v78_) < (d_567_v78_)])
        d_568_v79_ = d_568_v79_
        d_569_v80_: _dafny.Map
        d_569_v80_ = _dafny.Map({p0: d_558_v73_})
        d_570_v81_: _dafny.Seq
        d_570_v81_ = _dafny.SeqWithoutIsStrInference([d_569_v80_])
        r0 = (d_570_v81_) != (d_570_v81_)
        r1 = True
        r2 = p0
        return r0, r1, r2


class C2(T0):
    def  __init__(self):
        self._f31: bool = False
        self._f38: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f31(self):
        return self._f31
    def ctor__(self, f38, f31):
        (self)._f38 = f38
        (self)._f31 = f31

    def fm12(self, p0, p1, p2, p3, globalState):
        return (D2_DC6(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "my")))])))).cf8

    def fm13(self, p0, globalState):
        return ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwst"))))) - (84)

    def m4(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_571_v0_: _dafny.Seq
        d_571_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnnicgsqf"))
        d_572_v1_: str
        d_572_v1_ = _dafny.CodePoint('r')
        d_573_v2_: D0
        d_573_v2_ = D0_DC1(d_572_v1_, (self).f31, p2)
        d_574_v3_: _dafny.Map
        d_574_v3_ = _dafny.Map({(d_573_v2_).cf2: p1})
        hi3_ = len((d_574_v3_) | (d_574_v3_))
        for d_575_i0_ in range((len(d_571_v0_)) * (p1), hi3_):
            d_576_v4_: _dafny.Array
            nw96_ = _dafny.Array(False, 26)
            d_576_v4_ = nw96_
            d_577_v5_: _dafny.Map
            d_577_v5_ = _dafny.Map({910: d_576_v4_})
            d_578_v6_: _dafny.Map
            d_578_v6_ = _dafny.Map({len(default__.fm14(d_575_i0_, d_575_i0_, p1, True, globalState)): p0})
            d_579_v7_: _dafny.Set
            d_579_v7_ = _dafny.Set({p3, False})
            d_580_v8_: _dafny.Set
            d_580_v8_ = _dafny.Set({len(d_579_v7_), p0, 151, p0, p1})
            d_581_v9_: D5
            d_581_v9_ = D5_DC14(d_580_v8_, d_575_i0_, p2, d_576_v4_)
            d_582_v10_: _dafny.Array
            nw97_ = _dafny.Array(None, 23)
            nw97_[int(0)] = ((d_577_v5_)[len((d_578_v6_).set(p1, len(d_571_v0_)))] if (len((d_578_v6_).set(p1, len(d_571_v0_)))) in (d_577_v5_) else d_576_v4_)
            nw97_[int(1)] = d_576_v4_
            nw97_[int(2)] = d_576_v4_
            nw97_[int(3)] = (d_576_v4_ if (self).f31 else d_576_v4_)
            nw97_[int(4)] = d_576_v4_
            nw97_[int(5)] = d_576_v4_
            nw97_[int(6)] = d_576_v4_
            nw97_[int(7)] = d_576_v4_
            nw97_[int(8)] = d_576_v4_
            nw97_[int(9)] = (d_581_v9_).cf22
            nw97_[int(10)] = d_576_v4_
            nw97_[int(11)] = d_576_v4_
            nw97_[int(12)] = d_576_v4_
            nw97_[int(13)] = d_576_v4_
            nw97_[int(14)] = d_576_v4_
            nw97_[int(15)] = d_576_v4_
            nw97_[int(16)] = d_576_v4_
            nw97_[int(17)] = d_576_v4_
            nw97_[int(18)] = d_576_v4_
            nw97_[int(19)] = d_576_v4_
            nw97_[int(20)] = d_576_v4_
            nw97_[int(21)] = d_576_v4_
            nw97_[int(22)] = d_576_v4_
            d_582_v10_ = nw97_
            d_582_v10_ = d_582_v10_
            if p2:
                d_583_v11_: int
                d_583_v11_ = -699
                rhs95_ = (p2 if p2 else (d_579_v7_).isdisjoint(d_579_v7_))
                rhs96_ = (D2_DC6(p0)).cf8
                rhs97_ = len(d_578_v6_)
                rhs98_ = (-292) - (d_583_v11_)
                lhs57_ = globalState
                lhs57_.f28 = rhs95_
                d_583_v11_ = rhs96_
                d_583_v11_ = rhs97_
                d_583_v11_ = rhs98_
                d_584_v12_: _dafny.Set
                d_584_v12_ = _dafny.Set({d_572_v1_, d_572_v1_, d_572_v1_, d_572_v1_, d_572_v1_})
                d_585_v13_: C0
                nw98_ = C0()
                nw98_.ctor__((d_584_v12_) == (d_584_v12_))
                d_585_v13_ = nw98_
                d_572_v1_ = d_572_v1_
                d_583_v11_ = (d_583_v11_) - (p1)
                d_586_v14_: _dafny.MultiSet
                d_586_v14_ = _dafny.MultiSet([True, p3, (d_581_v9_).cf21])
                (globalState).f28 = (d_586_v14_).ispropersubset(_dafny.MultiSet([p2, not(p2), p2]))
            elif True:
                d_587_v15_: _dafny.Map
                d_587_v15_ = _dafny.Map({p2: p2})
                d_588_v16_: _dafny.Map
                d_588_v16_ = _dafny.Map({_dafny.CodePoint('p'): p2})
                (globalState).f28 = not(((d_587_v15_)[p3] if (p3) in (d_587_v15_) else ((d_588_v16_)[d_572_v1_] if (d_572_v1_) in (d_588_v16_) else (self).f31)))
                d_589_v17_: int
                d_589_v17_ = 100
                d_589_v17_ = p1
                (globalState).f16 = p2
                d_590_v18_: C0
                nw99_ = C0()
                nw99_.ctor__(True)
                d_590_v18_ = nw99_
                d_591_v19_: _dafny.Array
                nw100_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_591_v19_ = nw100_
                d_592_v20_: _dafny.Map
                d_592_v20_ = _dafny.Map({d_591_v19_: d_589_v17_})
                d_593_v21_: _dafny.Map
                d_593_v21_ = _dafny.Map({d_575_i0_: d_591_v19_})
                d_592_v20_ = (d_592_v20_).set(((d_593_v21_)[-299] if (-299) in (d_593_v21_) else d_591_v19_), (self).fm13((0) - (d_589_v17_), globalState))
            (globalState).f28 = p3
            d_594_v22_: _dafny.Seq
            d_594_v22_ = _dafny.SeqWithoutIsStrInference([p0, p0, p1])
            d_595_v23_: _dafny.Map
            d_595_v23_ = _dafny.Map({(self).f31: d_594_v22_})
            d_595_v23_ = (d_595_v23_).set(p2, (d_594_v22_) + (_dafny.SeqWithoutIsStrInference([(self).fm12(p1, p0, d_575_i0_, d_580_v8_, globalState)])))
        hi4_ = p1
        for d_596_i1_ in range(p1, hi4_):
            d_597_v24_: _dafny.Map
            d_597_v24_ = _dafny.Map({default__.safeModuloInt(p1, len(_dafny.Map({d_596_i1_: d_596_i1_}))): (self).f31})
            d_597_v24_ = (d_597_v24_).set(p1, (self).f31)
            d_598_v25_: _dafny.Array
            nw101_ = _dafny.Array(_dafny.Seq({}), 20)
            d_598_v25_ = nw101_
            d_599_v26_: _dafny.Map
            d_599_v26_ = _dafny.Map({d_598_v25_: p1})
            d_599_v26_ = (d_599_v26_).set(d_598_v25_, default__.safeModuloInt(p0, (0) - (-690)))
            d_600_v27_: C0
            nw102_ = C0()
            nw102_.ctor__((self).f31)
            d_600_v27_ = nw102_
            d_601_v28_: _dafny.Array
            def lambda16_(d_602_p1_):
                def lambda17_(d_603_i2_):
                    return default__.safeDivisionInt(d_603_i2_, d_602_p1_)

                return lambda17_

            init9_ = lambda16_(p1)
            nw103_ = _dafny.Array(None, 8)
            for i0_9_ in range(nw103_.length(0)):
                nw103_[i0_9_] = init9_(i0_9_)
            d_601_v28_ = nw103_
            d_604_v29_: _dafny.Array
            d_604_v29_ = d_601_v28_
            source9_ = d_604_v29_
            d_605___mcc_h0_ = source9_
            d_606_cf7_ = d_605___mcc_h0_
            d_607_v30_: _dafny.Map
            d_607_v30_ = _dafny.Map({(self).f31: p3})
            d_608_v31_: _dafny.Seq
            d_608_v31_ = _dafny.SeqWithoutIsStrInference([d_607_v30_])
            d_609_v32_: _dafny.MultiSet
            d_609_v32_ = _dafny.MultiSet([d_596_i1_, (self).fm13(p1, globalState), p1])
            d_610_v33_: int
            d_610_v33_ = 465
            d_611_v34_: _dafny.Seq
            d_611_v34_ = _dafny.SeqWithoutIsStrInference([(self).f38])
            d_612_v35_: _dafny.Set
            d_612_v35_ = _dafny.Set({len(d_611_v34_), len(_dafny.SeqWithoutIsStrInference([(self).f31])), d_596_i1_})
            rhs99_ = (d_608_v31_).set(default__.safeIndex(len(default__.fm15(p2, globalState)), len(d_608_v31_)), (d_607_v30_) | (d_607_v30_))
            rhs100_ = d_609_v32_
            rhs101_ = (0) - (len(((d_571_v0_) + (d_571_v0_)) + (d_571_v0_)))
            rhs102_ = (self).fm12(len(_dafny.SeqWithoutIsStrInference([True, (self).f31, p2, (self).f31, p2])), (self).fm13(p0, globalState), p0, (_dafny.Set({d_596_i1_})) - (d_612_v35_), globalState)
            rhs103_ = p2
            lhs58_ = globalState
            d_608_v31_ = rhs99_
            d_609_v32_ = rhs100_
            d_610_v33_ = rhs101_
            d_610_v33_ = rhs102_
            lhs58_.f28 = rhs103_
            d_613_v36_: _dafny.MultiSet
            d_613_v36_ = _dafny.MultiSet([(self).f31])
            d_614_v37_: _dafny.Map
            d_614_v37_ = _dafny.Map({(self).f31: ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhpbetysj"))) + (default__.fm0(d_613_v36_, globalState))).set(default__.safeIndex((self).fm13(p1, globalState), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhpbetysj"))) + (default__.fm0(d_613_v36_, globalState)))), _dafny.CodePoint('p'))})
            d_614_v37_ = (d_614_v37_).set(True, d_571_v0_)
            rhs104_ = (self).fm13(p1, globalState)
            rhs105_ = p1
            rhs106_ = (d_604_v29_)
            d_610_v33_ = rhs104_
            d_610_v33_ = rhs105_
            d_606_cf7_ = rhs106_
            d_615_v38_: C0
            nw104_ = C0()
            nw104_.ctor__((self).f31)
            d_615_v38_ = nw104_
        d_616_v39_: _dafny.Set
        d_616_v39_ = _dafny.Set({p0})
        default__.m0(len((d_616_v39_) - (d_616_v39_)), globalState)
        r0 = ((self).f31 if p3 else p3)
        d_571_v0_ = (_dafny.SeqWithoutIsStrInference([(d_572_v1_ if p3 else d_572_v1_) for d_617_i3_ in range(default__.abs(891))])).set(default__.safeIndex(len(d_571_v0_), len(_dafny.SeqWithoutIsStrInference([(d_572_v1_ if p3 else d_572_v1_) for d_618_i3_ in range(default__.abs(891))]))), d_572_v1_)
        d_619_i4_: int
        d_619_i4_ = 0
        with _dafny.label("7"):
            while (self).f31:
                with _dafny.c_label("7"):
                    if (d_619_i4_) >= (100):
                        raise _dafny.Break("7")
                    d_619_i4_ = (d_619_i4_) + (1)
                    d_620_v40_: _dafny.Array
                    nw105_ = _dafny.Array(int(0), 20)
                    d_620_v40_ = nw105_
                    d_621_v41_: _dafny.Map
                    d_621_v41_ = _dafny.Map({d_620_v40_: (self).fm13(p0, globalState)})
                    d_622_v42_: _dafny.Seq
                    d_622_v42_ = _dafny.SeqWithoutIsStrInference([p0, len((d_571_v0_ if p3 else d_571_v0_)), default__.safeDivisionInt(len(d_621_v41_), p0), p0])
                    d_622_v42_ = d_622_v42_
                    d_623_v43_: _dafny.Array
                    def lambda18_(d_624_i5_):
                        return True

                    init10_ = lambda18_
                    nw106_ = _dafny.Array(None, 13)
                    for i0_10_ in range(nw106_.length(0)):
                        nw106_[i0_10_] = init10_(i0_10_)
                    d_623_v43_ = nw106_
                    (globalState).f20 = d_623_v43_
                    r0 = not (p2) or (p2)
                    if ((p1) not in (_dafny.MultiSet(d_622_v42_))) not in ((self).f38):
                        d_625_v44_: _dafny.Map
                        d_625_v44_ = _dafny.Map({46: d_572_v1_})
                        d_625_v44_ = (d_625_v44_).set(p0, d_572_v1_)
                        d_626_v45_: int
                        d_626_v45_ = 825
                        d_627_v46_: _dafny.Map
                        d_627_v46_ = _dafny.Map({p2: d_572_v1_})
                        d_626_v45_ = (self).fm12(p1, p1, (self).fm12(len(d_627_v46_), p0, len(d_571_v0_), d_616_v39_, globalState), d_616_v39_, globalState)
                        d_628_v47_: _dafny.MultiSet
                        d_628_v47_ = _dafny.MultiSet([not(p2)])
                        d_571_v0_ = default__.fm0(d_628_v47_, globalState)
                        d_629_v48_: _dafny.Seq
                        d_629_v48_ = _dafny.SeqWithoutIsStrInference([d_620_v40_, d_620_v40_, d_620_v40_])
                        d_630_v49_: _dafny.Array
                        nw107_ = _dafny.Array(None, 7)
                        nw107_[int(0)] = d_623_v43_
                        nw107_[int(1)] = d_623_v43_
                        nw107_[int(2)] = d_623_v43_
                        nw107_[int(3)] = d_623_v43_
                        nw107_[int(4)] = d_623_v43_
                        nw107_[int(5)] = d_623_v43_
                        nw107_[int(6)] = d_623_v43_
                        d_630_v49_ = nw107_
                        d_631_v50_: T1
                        nw108_ = C1()
                        nw108_.ctor__(p1, d_630_v49_, (self).f31)
                        d_631_v50_ = nw108_
                        d_632_v51_: _dafny.Map
                        d_632_v51_ = _dafny.Map({d_631_v50_: d_626_v45_})
                        d_626_v45_ = len((d_629_v48_).set(default__.safeIndex(((d_632_v51_)[d_631_v50_] if (d_631_v50_) in (d_632_v51_) else p1), len(d_629_v48_)), d_620_v40_))
                        d_633_v52_: _dafny.Map
                        d_633_v52_ = _dafny.Map({(0) - ((d_631_v50_).f32): default__.fm26((self).f31, p2, p3, (0) - ((d_631_v50_).f32), globalState)})
                        d_633_v52_ = (d_633_v52_).set((d_631_v50_).f32, p2)
                    elif True:
                        d_634_v53_: _dafny.Seq
                        d_634_v53_ = _dafny.SeqWithoutIsStrInference([d_571_v0_])
                        d_571_v0_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyaujnuy"))) + (d_571_v0_)) + ((d_634_v53_)[default__.safeIndex(p0, len(d_634_v53_))])
                        d_635_v54_: _dafny.Map
                        d_635_v54_ = _dafny.Map({default__.safeModuloInt(len(d_574_v3_), (0) - (p0)): d_623_v43_})
                        d_635_v54_ = (d_635_v54_) | ((D8_DC20(_dafny.Map({p0: d_623_v43_}))).cf30)
                        (globalState).f28 = p2
                        d_636_v55_: int
                        d_636_v55_ = 590
                        d_637_v56_: _dafny.Map
                        d_637_v56_ = _dafny.Map({(self).f31: d_571_v0_})
                        d_636_v55_ = len(d_637_v56_)
                        d_638_v57_: _dafny.Array
                        nw109_ = _dafny.Array(_dafny.Seq({}), 13)
                        d_638_v57_ = nw109_
                        d_638_v57_ = (d_638_v57_ if (self).f31 else d_638_v57_)
                    pass
            pass
        r0 = p3
        d_639_v59_: _dafny.Array
        def lambda19_(d_640_v1_, d_641_p2_, d_642_p3_):
            def lambda20_(d_643_i6_):
                def iife51_():
                    coll19_ = _dafny.Map()
                    compr_19_: D2
                    for compr_19_ in (_dafny.Map({D2_DC7(d_640_v1_, _dafny.Set({d_641_p2_})): False})).keys.Elements:
                        d_644_v58_: D2 = compr_19_
                        if (d_644_v58_) in (_dafny.Map({D2_DC7(d_640_v1_, _dafny.Set({d_641_p2_})): False})):
                            coll19_[d_644_v58_] = d_642_p3_
                    return _dafny.Map(coll19_)
                return iife51_()
                

            return lambda20_

        init11_ = lambda19_(d_572_v1_, p2, p3)
        nw110_ = _dafny.Array(None, 19)
        for i0_11_ in range(nw110_.length(0)):
            nw110_[i0_11_] = init11_(i0_11_)
        d_639_v59_ = nw110_
        r1 = ((d_639_v59_ if False else d_639_v59_) if p2 else d_639_v59_)
        return r0, r1

    @property
    def f38(self):
        return self._f38

class C3(T0):
    def  __init__(self):
        self._f31: bool = False
        self._f36: _dafny.Array = _dafny.Array(None, 0)
        self._f37: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f31(self):
        return self._f31
    def ctor__(self, f36, f37, f31):
        (self)._f36 = f36
        (self)._f37 = f37
        (self)._f31 = f31

    def fm9(self, p0, p1, p2, p3, globalState):
        def lambda21_(source10_):
            if source10_.is_DC7:
                d_645___mcc_h0_ = source10_.cf9
                d_646___mcc_h1_ = source10_.cf10
                d_647_cf10_ = d_646___mcc_h1_
                d_648_cf9_ = d_645___mcc_h0_
                if (self).f31:
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tj"))
                elif True:
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "puwxcvuv"))
            elif source10_.is_DC6:
                d_649___mcc_h2_ = source10_.cf8
                d_650_cf8_ = d_649___mcc_h2_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rasephr"))
            elif True:
                d_651___mcc_h3_ = source10_.cf11
                d_652_cf11_ = d_651___mcc_h3_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wlyxgncg"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mit")))

        def iife52_():
            coll20_ = _dafny.Map()
            compr_20_: int
            for compr_20_ in _dafny.IntegerRange(691, 135):
                d_653_v0_: int = compr_20_
                if ((691) <= (d_653_v0_)) and ((d_653_v0_) < (135)):
                    coll20_[(d_653_v0_) - (len(_dafny.Set({(self).f31, (self).f31, (self).f31, (self).f31, (self).f31})))] = 640
            return _dafny.Map(coll20_)
        return len(lambda21_(D2_DC6(len(iife52_()
))))

    def m3(self, p0, p1, p2, p3, globalState):
        d_654_v0_: int
        d_654_v0_ = -523
        d_654_v0_ = default__.safeModuloInt(p2, p2)
        if p3:
            d_655_v1_: _dafny.Seq
            d_655_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uda"))
            d_656_v2_: _dafny.Array
            nw111_ = _dafny.Array(None, 10)
            nw111_[int(0)] = d_655_v1_
            nw111_[int(1)] = (d_655_v1_) + (d_655_v1_)
            nw111_[int(2)] = d_655_v1_
            nw111_[int(3)] = d_655_v1_
            nw111_[int(4)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_657_i0_ in range(default__.abs(60))])) + (d_655_v1_)
            nw111_[int(5)] = d_655_v1_
            nw111_[int(6)] = d_655_v1_
            nw111_[int(7)] = d_655_v1_
            nw111_[int(8)] = (d_655_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))
            nw111_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qviqoktq"))
            d_656_v2_ = nw111_
            index93_ = default__.safeIndex(410, (d_656_v2_).length(0))
            (d_656_v2_)[index93_] = d_655_v1_
            index94_ = default__.safeIndex(410, (d_656_v2_).length(0))
            (d_656_v2_)[index94_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldgnc"))
            d_658_v4_: _dafny.Map
            d_658_v4_ = _dafny.Map({d_654_v0_: d_654_v0_})
            d_659_v5_: _dafny.MultiSet
            def iife53_():
                coll21_ = _dafny.Map()
                compr_21_: int
                for compr_21_ in (d_658_v4_).keys.Elements:
                    d_660_v3_: int = compr_21_
                    if (d_660_v3_) in (d_658_v4_):
                        coll21_[(d_660_v3_) * (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))), _dafny.CodePoint('o'))))] = (self).f31
                return _dafny.Map(coll21_)
            d_659_v5_ = _dafny.MultiSet([len(iife53_()
            )])
            d_661_v6_: _dafny.Map
            d_661_v6_ = _dafny.Map({(d_659_v5_).isdisjoint(d_659_v5_): (self).f31})
            d_661_v6_ = (d_661_v6_).set(p3, p3)
            d_662_v7_: _dafny.MultiSet
            d_662_v7_ = _dafny.MultiSet([(self).f31, p3])
            d_655_v1_ = default__.fm0(d_662_v7_, globalState)
            d_663_v8_: _dafny.Array
            nw112_ = _dafny.Array(_dafny.MultiSet({}), 25)
            d_663_v8_ = nw112_
            index95_ = default__.safeIndex(736, (d_663_v8_).length(0))
            (d_663_v8_)[index95_] = d_662_v7_
            d_664_v9_: _dafny.Array
            nw113_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_664_v9_ = nw113_
            d_665_v10_: _dafny.Set
            d_665_v10_ = _dafny.Set({(self).f31})
            d_666_v11_: str
            d_666_v11_ = _dafny.CodePoint('d')
            d_667_v12_: D0
            d_667_v12_ = D0_DC1(d_666_v11_, (self).f31, p3)
            d_668_v13_: _dafny.Array
            nw114_ = _dafny.Array(None, 25)
            nw114_[int(0)] = default__.fm10(d_666_v11_, (self).f31, globalState)
            nw114_[int(1)] = False
            nw114_[int(2)] = (self).f31
            nw114_[int(3)] = p3
            nw114_[int(4)] = (self).f31
            nw114_[int(5)] = True
            nw114_[int(6)] = p3
            nw114_[int(7)] = (self).f31
            nw114_[int(8)] = p3
            nw114_[int(9)] = (self).f31
            nw114_[int(10)] = True
            nw114_[int(11)] = p3
            nw114_[int(12)] = (self).f31
            nw114_[int(13)] = (self).f31
            nw114_[int(14)] = (d_667_v12_).cf3
            nw114_[int(15)] = p3
            nw114_[int(16)] = not((self).f31)
            nw114_[int(17)] = (self).f31
            nw114_[int(18)] = default__.fm10(d_666_v11_, (self).f31, globalState)
            nw114_[int(19)] = True
            nw114_[int(20)] = p3
            nw114_[int(21)] = True
            nw114_[int(22)] = p3
            nw114_[int(23)] = p3
            nw114_[int(24)] = not((self).f31)
            d_668_v13_ = nw114_
            d_669_v14_: _dafny.Map
            d_669_v14_ = _dafny.Map({p3: d_668_v13_})
            d_670_v15_: _dafny.Array
            nw115_ = _dafny.Array(None, 5)
            nw115_[int(0)] = len(d_665_v10_)
            nw115_[int(1)] = len(_dafny.Set({d_654_v0_}))
            nw115_[int(2)] = (0) - (((d_662_v7_)[(self).f31] if ((self).f31) in (d_662_v7_) else len(d_669_v14_)))
            nw115_[int(3)] = d_654_v0_
            nw115_[int(4)] = len(p1)
            d_670_v15_ = nw115_
            index96_ = default__.safeIndex(270, (d_664_v9_).length(0))
            (d_664_v9_)[index96_] = d_670_v15_
            d_671_v16_: D5
            d_671_v16_ = D5_DC13(d_662_v7_)
            index97_ = default__.safeIndex(736, (d_663_v8_).length(0))
            index98_ = default__.safeIndex(270, (d_664_v9_).length(0))
            rhs107_ = (d_671_v16_).cf18
            rhs108_ = d_670_v15_
            lhs59_ = d_663_v8_
            lhs60_ = default__.safeIndex(736, (d_663_v8_).length(0))
            lhs61_ = d_664_v9_
            lhs62_ = default__.safeIndex(270, (d_664_v9_).length(0))
            lhs59_[lhs60_] = rhs107_
            lhs61_[lhs62_] = rhs108_
            arr0_ = (d_664_v9_)[default__.safeIndex(270, (d_664_v9_).length(0))]
            index99_ = default__.safeIndex(843, ((d_664_v9_)[default__.safeIndex(270, (d_664_v9_).length(0))]).length(0))
            arr0_[index99_] = (0) - (p2)
            d_672_v17_: _dafny.Seq
            d_672_v17_ = _dafny.SeqWithoutIsStrInference([p2])
            d_673_v18_: _dafny.Map
            d_673_v18_ = _dafny.Map({d_670_v15_: (default__.fm11(globalState)) + (d_672_v17_)})
            d_674_v19_: _dafny.Map
            d_674_v19_ = _dafny.Map({(self).f31: p2})
            arr1_ = (d_664_v9_)[default__.safeIndex(270, (d_664_v9_).length(0))]
            index100_ = default__.safeIndex(843, ((d_664_v9_)[default__.safeIndex(270, (d_664_v9_).length(0))]).length(0))
            index101_ = default__.safeIndex(410, (d_656_v2_).length(0))
            rhs109_ = len(((d_673_v18_)[(d_670_v15_ if p3 else d_670_v15_)] if ((d_670_v15_ if p3 else d_670_v15_)) in (d_673_v18_) else (d_672_v17_ if p3 else d_672_v17_)))
            rhs110_ = p2
            rhs111_ = (d_654_v0_) * (d_654_v0_)
            rhs112_ = ((d_656_v2_)[default__.safeIndex(410, (d_656_v2_).length(0))]).set(default__.safeIndex(default__.safeDivisionInt(p2, d_654_v0_), len((d_656_v2_)[default__.safeIndex(410, (d_656_v2_).length(0))])), ((d_656_v2_)[default__.safeIndex(410, (d_656_v2_).length(0))])[default__.safeIndex(len(d_674_v19_), len((d_656_v2_)[default__.safeIndex(410, (d_656_v2_).length(0))]))])
            lhs63_ = (d_664_v9_)[default__.safeIndex(270, (d_664_v9_).length(0))]
            lhs64_ = default__.safeIndex(843, ((d_664_v9_)[default__.safeIndex(270, (d_664_v9_).length(0))]).length(0))
            lhs65_ = d_656_v2_
            lhs66_ = default__.safeIndex(410, (d_656_v2_).length(0))
            d_654_v0_ = rhs109_
            d_654_v0_ = rhs110_
            lhs63_[lhs64_] = rhs111_
            lhs65_[lhs66_] = rhs112_
        elif True:
            d_675_v20_: _dafny.Set
            d_675_v20_ = _dafny.Set({d_654_v0_, p2})
            d_676_v21_: _dafny.Map
            d_676_v21_ = _dafny.Map({(self).f31: len(d_675_v20_)})
            d_677_v22_: _dafny.MultiSet
            d_677_v22_ = _dafny.MultiSet([p2, p2, ((d_676_v21_)[p3] if (p3) in (d_676_v21_) else d_654_v0_)])
            d_678_v23_: _dafny.Seq
            d_678_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eys"))
            d_679_v24_: _dafny.Seq
            d_679_v24_ = _dafny.SeqWithoutIsStrInference([d_654_v0_])
            d_680_v25_: _dafny.Array
            nw116_ = _dafny.Array(None, 10)
            nw116_[int(0)] = (self).f31
            nw116_[int(1)] = (d_677_v22_) != (d_677_v22_)
            nw116_[int(2)] = (p2) < ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (p2) for d_681_i1_ in range(default__.abs(918))]))).cardinality)
            nw116_[int(3)] = ((d_677_v22_).set(d_654_v0_, default__.abs(d_654_v0_))).issubset(d_677_v22_)
            nw116_[int(4)] = False
            nw116_[int(5)] = True
            nw116_[int(6)] = (d_654_v0_) < (d_654_v0_)
            nw116_[int(7)] = (self).f31
            nw116_[int(8)] = (d_678_v23_) == (d_678_v23_)
            nw116_[int(9)] = (d_679_v24_) == (d_679_v24_)
            d_680_v25_ = nw116_
            index102_ = default__.safeIndex(758, (d_680_v25_).length(0))
            (d_680_v25_)[index102_] = not(p3)
            d_682_v26_: _dafny.Array
            def lambda22_(d_683_i2_):
                return _dafny.CodePoint('w')

            init12_ = lambda22_
            nw117_ = _dafny.Array(None, 16)
            for i0_12_ in range(nw117_.length(0)):
                nw117_[i0_12_] = init12_(i0_12_)
            d_682_v26_ = nw117_
            d_684_v27_: _dafny.Map
            d_684_v27_ = _dafny.Map({True: d_682_v26_})
            d_685_v28_: str
            d_685_v28_ = _dafny.CodePoint('h')
            d_686_v29_: _dafny.Map
            d_686_v29_ = _dafny.Map({(self).f31: d_685_v28_})
            d_687_v30_: _dafny.Array
            nw118_ = _dafny.Array(int(0), 19)
            d_687_v30_ = nw118_
            d_688_v31_: _dafny.Map
            d_688_v31_ = _dafny.Map({d_687_v30_: p2})
            d_689_v32_: _dafny.Map
            d_689_v32_ = _dafny.Map({p2: d_687_v30_})
            d_690_v33_: _dafny.Array
            nw119_ = _dafny.Array(None, 7)
            nw119_[int(0)] = len(d_684_v27_)
            nw119_[int(1)] = (self).fm9(977, d_679_v24_, d_654_v0_, (self).f31, globalState)
            nw119_[int(2)] = (self).fm9((self).fm9(557, d_679_v24_, d_654_v0_, p3, globalState), (_dafny.SeqWithoutIsStrInference([-745 for d_691_i3_ in range(default__.abs(-475))])).set(default__.safeIndex(len(d_686_v29_), len(_dafny.SeqWithoutIsStrInference([-745 for d_692_i3_ in range(default__.abs(-475))]))), len(d_688_v31_)), p2, (self).f31, globalState)
            nw119_[int(3)] = d_654_v0_
            nw119_[int(4)] = len((d_689_v32_) | (_dafny.Map({p2: d_687_v30_})))
            nw119_[int(5)] = (0) - (len(p1))
            nw119_[int(6)] = d_654_v0_
            d_690_v33_ = nw119_
            index103_ = default__.safeIndex(905, (d_687_v30_).length(0))
            (d_687_v30_)[index103_] = d_654_v0_
            index104_ = default__.safeIndex(758, (d_680_v25_).length(0))
            index105_ = default__.safeIndex(905, (d_687_v30_).length(0))
            rhs113_ = d_680_v25_
            rhs114_ = p2
            rhs115_ = ((self).f31 if (self).f31 else (self).f31)
            rhs116_ = (d_654_v0_) - ((self).fm9(d_654_v0_, d_679_v24_, 927, (self).f31, globalState))
            lhs67_ = globalState
            lhs68_ = d_680_v25_
            lhs69_ = default__.safeIndex(758, (d_680_v25_).length(0))
            lhs70_ = d_687_v30_
            lhs71_ = default__.safeIndex(905, (d_687_v30_).length(0))
            lhs67_.f27 = rhs113_
            d_654_v0_ = rhs114_
            lhs68_[lhs69_] = rhs115_
            lhs70_[lhs71_] = rhs116_
            d_693_v34_: C0
            nw120_ = C0()
            nw120_.ctor__((d_680_v25_)[default__.safeIndex(758, (d_680_v25_).length(0))])
            d_693_v34_ = nw120_
            d_694_v35_: _dafny.Array
            nw121_ = _dafny.Array(None, 1)
            nw121_[int(0)] = d_690_v33_
            d_694_v35_ = nw121_
            index106_ = default__.safeIndex(8, (d_694_v35_).length(0))
            (d_694_v35_)[index106_] = d_690_v33_
            d_695_v36_: _dafny.Map
            d_695_v36_ = _dafny.Map({p2: (d_687_v30_)[default__.safeIndex(905, (d_687_v30_).length(0))]})
            d_696_v37_: _dafny.Seq
            d_696_v37_ = _dafny.SeqWithoutIsStrInference([(d_677_v22_).intersection(_dafny.MultiSet([len((d_695_v36_).set(p2, 975))])), _dafny.MultiSet(d_679_v24_), d_677_v22_])
            d_697_v38_: _dafny.MultiSet
            d_697_v38_ = _dafny.MultiSet([(self).f31])
            d_698_v39_: D8
            d_698_v39_ = D8_DC21(p3, (self).f31, (d_687_v30_)[default__.safeIndex(905, (d_687_v30_).length(0))])
            pat_let_tv10_ = d_680_v25_
            pat_let_tv11_ = d_680_v25_
            index107_ = default__.safeIndex(8, (d_694_v35_).length(0))
            index108_ = default__.safeIndex(905, (d_687_v30_).length(0))
            def iife54_(_pat_let16_0):
                def iife55_(d_699_dt__update__tmp_h0_):
                    def iife56_(_pat_let17_0):
                        def iife57_(d_700_dt__update_hcf31_h0_):
                            return D8_DC21(d_700_dt__update_hcf31_h0_, (d_699_dt__update__tmp_h0_).cf32, (d_699_dt__update__tmp_h0_).cf33)
                        return iife57_(_pat_let17_0)
                    return iife56_((pat_let_tv11_)[default__.safeIndex(758, (pat_let_tv10_).length(0))])
                return iife55_(_pat_let16_0)
            rhs117_ = (0) - (len(d_696_v37_))
            rhs118_ = (d_690_v33_ if p3 else d_690_v33_)
            rhs119_ = (d_685_v28_ if (d_680_v25_)[default__.safeIndex(758, (d_680_v25_).length(0))] else d_685_v28_)
            rhs120_ = p2
            rhs121_ = default__.fm0((d_697_v38_).set((d_680_v25_)[default__.safeIndex(758, (d_680_v25_).length(0))], default__.abs((iife54_(d_698_v39_)).cf33)), globalState)
            lhs72_ = d_694_v35_
            lhs73_ = default__.safeIndex(8, (d_694_v35_).length(0))
            lhs74_ = d_687_v30_
            lhs75_ = default__.safeIndex(905, (d_687_v30_).length(0))
            d_654_v0_ = rhs117_
            lhs72_[lhs73_] = rhs118_
            d_685_v28_ = rhs119_
            lhs74_[lhs75_] = rhs120_
            d_678_v23_ = rhs121_
            d_701_v40_: _dafny.Array
            nw122_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_701_v40_ = nw122_
            d_702_v41_: D5
            d_702_v41_ = D5_DC14(d_675_v20_, d_654_v0_, p3, d_680_v25_)
            index109_ = default__.safeIndex(256, (d_701_v40_).length(0))
            (d_701_v40_)[index109_] = ((d_702_v41_).cf22 if p3 else d_680_v25_)
            index110_ = default__.safeIndex(256, (d_701_v40_).length(0))
            (d_701_v40_)[index110_] = d_680_v25_
            d_703_v42_: _dafny.Seq
            d_703_v42_ = _dafny.SeqWithoutIsStrInference([d_678_v23_])
            d_704_v43_: _dafny.Seq
            d_704_v43_ = _dafny.SeqWithoutIsStrInference([(d_679_v24_).set(default__.safeIndex((d_687_v30_)[default__.safeIndex(905, (d_687_v30_).length(0))], len(d_679_v24_)), (d_687_v30_)[default__.safeIndex(905, (d_687_v30_).length(0))]), d_679_v24_])
            d_705_v44_: T1
            nw123_ = C1()
            nw123_.ctor__(len(d_704_v43_), d_701_v40_, (self).f31)
            d_705_v44_ = nw123_
            d_706_v45_: _dafny.Map
            d_706_v45_ = _dafny.Map({d_654_v0_: d_705_v44_})
            d_703_v42_ = (default__.fm27(p3, len(d_706_v45_), (self).f31, globalState)).cf39
        d_707_v46_: _dafny.MultiSet
        d_707_v46_ = _dafny.MultiSet([p3, (self).f31, False])
        d_708_v47_: _dafny.Seq
        d_708_v47_ = _dafny.SeqWithoutIsStrInference([(d_707_v46_).issubset(d_707_v46_)])
        d_709_v48_: _dafny.Map
        d_709_v48_ = _dafny.Map({p2: p3})
        d_710_v49_: _dafny.Seq
        d_710_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugtivxrgy"))
        d_711_v50_: D9
        d_711_v50_ = D9_DC26(default__.fm26(p3, p3, ((d_709_v48_)[d_654_v0_] if (d_654_v0_) in (d_709_v48_) else p3), p2, globalState), d_710_v49_, d_654_v0_, d_709_v48_, p3)
        pat_let_tv12_ = p1
        pat_let_tv13_ = p1
        pat_let_tv14_ = d_708_v47_
        pat_let_tv15_ = d_708_v47_
        pat_let_tv16_ = p1
        def lambda23_(source11_):
            if source11_.is_DC24:
                d_712___mcc_h0_ = source11_.cf40
                d_713_cf40_ = d_712___mcc_h0_
                return pat_let_tv12_
            elif source11_.is_DC25:
                d_714___mcc_h1_ = source11_.cf41
                d_715___mcc_h2_ = source11_.cf42
                d_716___mcc_h3_ = source11_.cf43
                d_717___mcc_h4_ = source11_.cf44
                d_718_cf44_ = d_717___mcc_h4_
                d_719_cf43_ = d_716___mcc_h3_
                d_720_cf42_ = d_715___mcc_h2_
                d_721_cf41_ = d_714___mcc_h1_
                return pat_let_tv13_
            elif source11_.is_DC26:
                d_722___mcc_h5_ = source11_.cf45
                d_723___mcc_h6_ = source11_.cf46
                d_724___mcc_h7_ = source11_.cf47
                d_725___mcc_h8_ = source11_.cf48
                d_726___mcc_h9_ = source11_.cf49
                d_727_cf49_ = d_726___mcc_h9_
                d_728_cf48_ = d_725___mcc_h8_
                d_729_cf47_ = d_724___mcc_h7_
                d_730_cf46_ = d_723___mcc_h6_
                d_731_cf45_ = d_722___mcc_h5_
                return pat_let_tv14_
            elif True:
                d_732___mcc_h10_ = source11_.cf39
                d_733_cf39_ = d_732___mcc_h10_
                return (pat_let_tv15_) + (pat_let_tv16_)

        d_708_v47_ = lambda23_(d_711_v50_)
        d_654_v0_ = p2
        d_734_v51_: _dafny.Array
        nw124_ = _dafny.Array(int(0), 29)
        d_734_v51_ = nw124_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_734_v51_).length(0)):
            d_735_i4_: int = guard_loop_3_
            if (True) and (((0) <= (d_735_i4_)) and ((d_735_i4_) < ((d_734_v51_).length(0)))):
                (d_734_v51_)[(d_735_i4_)] = (d_735_i4_) * ((len(p1)) * (47))
        if (p2) >= ((d_654_v0_) * ((0) - (p2))):
            default__.m0(p2, globalState)
            d_736_v52_: _dafny.MultiSet
            d_736_v52_ = _dafny.MultiSet([p1, d_708_v47_, d_708_v47_])
            (globalState).f16 = (d_736_v52_).isdisjoint(d_736_v52_)
            d_737_v53_: _dafny.Map
            d_737_v53_ = _dafny.Map({p3: d_654_v0_})
            d_737_v53_ = d_737_v53_
            d_738_v54_: D2
            d_738_v54_ = D2_DC6(p2)
            d_739_v55_: _dafny.Set
            d_739_v55_ = _dafny.Set({d_738_v54_})
            d_740_v56_: _dafny.Set
            d_740_v56_ = d_739_v55_
            source12_ = d_740_v56_
            d_741___mcc_h11_ = source12_
            d_742_cf17_ = d_741___mcc_h11_
            d_743_v57_: _dafny.Seq
            d_743_v57_ = _dafny.SeqWithoutIsStrInference([d_711_v50_])
            d_744_v59_: _dafny.MultiSet
            def iife58_():
                coll22_ = _dafny.Set()
                compr_22_: D9
                for compr_22_ in (d_743_v57_).Elements:
                    d_745_v58_: D9 = compr_22_
                    if (d_745_v58_) in (d_743_v57_):
                        coll22_ = coll22_.union(_dafny.Set([d_745_v58_]))
                return _dafny.Set(coll22_)
            d_744_v59_ = _dafny.MultiSet([iife58_()
            ])
            d_746_v60_: _dafny.Map
            d_746_v60_ = _dafny.Map({d_744_v59_: (0) - ((p2) * ((0) - (p2)))})
            d_747_v61_: _dafny.Seq
            d_747_v61_ = _dafny.SeqWithoutIsStrInference([d_744_v59_])
            d_654_v0_ = ((d_746_v60_)[((d_747_v61_)[default__.safeIndex(d_654_v0_, len(d_747_v61_))]) | (d_744_v59_)] if (((d_747_v61_)[default__.safeIndex(d_654_v0_, len(d_747_v61_))]) | (d_744_v59_)) in (d_746_v60_) else d_654_v0_)
            d_748_v62_: _dafny.Set
            d_748_v62_ = _dafny.Set({True})
            d_749_v63_: _dafny.Array
            nw125_ = _dafny.Array(None, 1)
            nw125_[int(0)] = (_dafny.Set({p3})).intersection(d_748_v62_)
            d_749_v63_ = nw125_
            d_749_v63_ = d_749_v63_
            d_654_v0_ = p2
            d_750_v64_: _dafny.Map
            d_750_v64_ = _dafny.Map({len(d_710_v49_): d_654_v0_})
            d_751_v65_: _dafny.Map
            d_751_v65_ = _dafny.Map({len(d_750_v64_): d_710_v49_})
            d_752_v66_: _dafny.Map
            d_752_v66_ = _dafny.Map({(self).f31: d_710_v49_})
            d_753_v67_: str
            d_753_v67_ = _dafny.CodePoint('k')
            d_754_v68_: _dafny.Array
            nw126_ = _dafny.Array(None, 27)
            nw126_[int(0)] = (d_710_v49_) + (((d_751_v65_)[p2] if (p2) in (d_751_v65_) else d_710_v49_))
            nw126_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lyh"))
            nw126_[int(2)] = d_710_v49_
            nw126_[int(3)] = d_710_v49_
            nw126_[int(4)] = ((d_752_v66_)[(d_708_v47_)[default__.safeIndex(p2, len(d_708_v47_))]] if ((d_708_v47_)[default__.safeIndex(p2, len(d_708_v47_))]) in (d_752_v66_) else d_710_v49_)
            nw126_[int(5)] = d_710_v49_
            nw126_[int(6)] = (d_710_v49_) + (d_710_v49_)
            nw126_[int(7)] = d_710_v49_
            nw126_[int(8)] = (d_710_v49_ if p3 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icxd")))
            nw126_[int(9)] = d_710_v49_
            nw126_[int(10)] = d_710_v49_
            nw126_[int(11)] = (default__.fm0(d_707_v46_, globalState)) + (d_710_v49_)
            nw126_[int(12)] = d_710_v49_
            nw126_[int(13)] = (d_710_v49_) + (d_710_v49_)
            nw126_[int(14)] = d_710_v49_
            nw126_[int(15)] = d_710_v49_
            nw126_[int(16)] = (d_710_v49_) + (d_710_v49_)
            nw126_[int(17)] = d_710_v49_
            nw126_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eghuqqi"))
            nw126_[int(19)] = d_710_v49_
            nw126_[int(20)] = (d_710_v49_).set(default__.safeIndex(d_654_v0_, len(d_710_v49_)), d_753_v67_)
            nw126_[int(21)] = (d_710_v49_) + (d_710_v49_)
            nw126_[int(22)] = d_710_v49_
            nw126_[int(23)] = d_710_v49_
            nw126_[int(24)] = (d_710_v49_) + (d_710_v49_)
            nw126_[int(25)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))).set(default__.safeIndex(573, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")))), d_753_v67_)) + (d_710_v49_)
            nw126_[int(26)] = _dafny.SeqWithoutIsStrInference([d_753_v67_ for d_755_i5_ in range(default__.abs(493))])
            d_754_v68_ = nw126_
            d_754_v68_ = d_754_v68_
            d_756_v69_: _dafny.Array
            def lambda24_(d_757_i6_):
                return (self).f31

            init13_ = lambda24_
            nw127_ = _dafny.Array(None, 20)
            for i0_13_ in range(nw127_.length(0)):
                nw127_[i0_13_] = init13_(i0_13_)
            d_756_v69_ = nw127_
            d_758_v70_: _dafny.Set
            d_758_v70_ = _dafny.Set({len(d_710_v49_), len(_dafny.SeqWithoutIsStrInference([p3]))})
            index111_ = default__.safeIndex(93, (d_756_v69_).length(0))
            (d_756_v69_)[index111_] = (d_758_v70_).issubset(d_758_v70_)
            index112_ = default__.safeIndex(93, (d_756_v69_).length(0))
            (d_756_v69_)[index112_] = (self).f31
        elif True:
            d_759_v71_: _dafny.Array
            def lambda25_(d_760_v0_, d_761_p2_):
                def lambda26_(d_762_i7_):
                    return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_760_v0_, d_761_p2_]))) | (_dafny.MultiSet([d_760_v0_, 213, d_761_p2_]))

                return lambda26_

            init14_ = lambda25_(d_654_v0_, p2)
            nw128_ = _dafny.Array(None, 2)
            for i0_14_ in range(nw128_.length(0)):
                nw128_[i0_14_] = init14_(i0_14_)
            d_759_v71_ = nw128_
            d_763_v72_: _dafny.MultiSet
            d_763_v72_ = _dafny.MultiSet([p2, p2])
            index113_ = default__.safeIndex(849, (d_759_v71_).length(0))
            (d_759_v71_)[index113_] = (d_763_v72_) | (_dafny.MultiSet([len(d_708_v47_), d_654_v0_]))
            index114_ = default__.safeIndex(849, (d_759_v71_).length(0))
            (d_759_v71_)[index114_] = (d_763_v72_).set(p2, default__.abs(955))
            d_654_v0_ = (p2) - (-367)
            d_710_v49_ = ((d_710_v49_) + (d_710_v49_)) + (d_710_v49_)
            d_764_v73_: _dafny.Map
            d_764_v73_ = _dafny.Map({(d_710_v49_)[default__.safeIndex(d_654_v0_, len(d_710_v49_))]: d_654_v0_})
            d_765_v74_: str
            d_765_v74_ = _dafny.CodePoint('n')
            d_764_v73_ = (d_764_v73_).set(d_765_v74_, (self).fm9((d_711_v50_).cf47, _dafny.SeqWithoutIsStrInference([d_654_v0_, d_654_v0_]), len(d_708_v47_), (self).f31, globalState))
            default__.m0(p2, globalState)

    @property
    def f36(self):
        return self._f36
    @property
    def f37(self):
        return self._f37

class C4(T1, T0):
    def  __init__(self):
        self._f31: bool = False
        self._f32: int = int(0)
        self._f33: _dafny.Array = _dafny.Array(None, 0)
        self.f35: bool = False
        self._f34: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f31(self):
        return self._f31
    @property
    def f32(self):
        return self._f32
    @property
    def f33(self):
        return self._f33
    def ctor__(self, f34, f35, f32, f33, f31):
        (self)._f34 = f34
        (self).f35 = f35
        (self)._f32 = f32
        (self)._f33 = f33
        (self)._f31 = f31

    def fm6(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([self.f35, (self).f31])) + (_dafny.SeqWithoutIsStrInference([True]))) + (_dafny.SeqWithoutIsStrInference([not(self.f35)]))

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r3: int = int(0)
        d_766_i0_: int
        d_766_i0_ = 0
        with _dafny.label("8"):
            while not((default__.safeModuloInt((self).f32, (self).f32)) < ((self).f32)):
                with _dafny.c_label("8"):
                    if (d_766_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_766_i0_ = (d_766_i0_) + (1)
                    (globalState).f28 = (self).f31
                    r0 = 462
                    d_767_v1_: _dafny.Seq
                    def iife59_():
                        coll23_ = _dafny.Set()
                        compr_23_: int
                        for compr_23_ in _dafny.IntegerRange(180, 462):
                            d_768_v0_: int = compr_23_
                            if ((180) <= (d_768_v0_)) and ((d_768_v0_) < (462)):
                                coll23_ = coll23_.union(_dafny.Set([(d_768_v0_) - (len(_dafny.Set({_dafny.Map({(self).f31: p0}), _dafny.Map({True: self.f35})})))]))
                        return _dafny.Set(coll23_)
                    d_767_v1_ = _dafny.SeqWithoutIsStrInference([(self).f32, (0) - ((len(iife59_()
                    )) - ((self).f32)), (self).f32])
                    d_767_v1_ = d_767_v1_
                    r3 = default__.fm7((self).f31, (self).f31, globalState)
                    pass
            pass
        d_769_v2_: _dafny.Seq
        d_769_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
        d_770_v3_: _dafny.MultiSet
        d_770_v3_ = _dafny.MultiSet([d_769_v2_])
        d_771_v4_: _dafny.Seq
        d_771_v4_ = _dafny.SeqWithoutIsStrInference([p0])
        d_772_v5_: _dafny.Seq
        d_772_v5_ = _dafny.SeqWithoutIsStrInference([(self).f32, (self).f32])
        d_773_v6_: _dafny.Seq
        d_773_v6_ = _dafny.SeqWithoutIsStrInference([((d_770_v3_)[_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_774_i1_ in range(default__.abs(876))])] if (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_775_i1_ in range(default__.abs(876))])) in (d_770_v3_) else len(d_771_v4_)), len(d_772_v5_), (self).f32, (self).f32, 214])
        d_776_v7_: _dafny.Set
        d_776_v7_ = _dafny.Set({d_772_v5_, d_773_v6_})
        if ((d_776_v7_) - (d_776_v7_)).issubset((d_776_v7_) - (d_776_v7_)):
            (globalState).f16 = self.f35
            d_777_v8_: _dafny.Array
            def lambda27_(d_778_v4_):
                def lambda28_(d_779_i2_):
                    return d_778_v4_

                return lambda28_

            init15_ = lambda27_(d_771_v4_)
            nw129_ = _dafny.Array(None, 9)
            for i0_15_ in range(nw129_.length(0)):
                nw129_[i0_15_] = init15_(i0_15_)
            d_777_v8_ = nw129_
            d_780_v9_: D0
            d_780_v9_ = D0_DC0(d_777_v8_)
            d_781_v10_: _dafny.Set
            d_781_v10_ = _dafny.Set({self.f35})
            pat_let_tv17_ = d_777_v8_
            pat_let_tv18_ = d_777_v8_
            def iife61_(_pat_let19_0):
                def iife62_(d_782_dt__update__tmp_h0_):
                    def iife63_(_pat_let20_0):
                        def iife64_(d_783_dt__update_hcf0_h0_):
                            return D0_DC0(d_783_dt__update_hcf0_h0_)
                        return iife64_(_pat_let20_0)
                    return iife63_(pat_let_tv17_)
                return iife62_(_pat_let19_0)
            def iife60_(_pat_let18_0):
                def iife65_(d_784_dt__update__tmp_h1_):
                    def iife66_(_pat_let21_0):
                        def iife67_(d_785_dt__update_hcf0_h1_):
                            return D0_DC0(d_785_dt__update_hcf0_h1_)
                        return iife67_(_pat_let21_0)
                    return iife66_(pat_let_tv18_)
                return iife65_(_pat_let18_0)
            rhs122_ = len(d_781_v10_)
            rhs123_ = iife60_(iife61_(d_780_v9_))
            r3 = rhs122_
            d_780_v9_ = rhs123_
            d_771_v4_ = d_771_v4_
            d_786_v11_: _dafny.Array
            nw130_ = _dafny.Array(False, 21)
            d_786_v11_ = nw130_
            index115_ = default__.safeIndex(888, (d_786_v11_).length(0))
            (d_786_v11_)[index115_] = default__.fm8(p0, 989, globalState)
            index116_ = default__.safeIndex(888, (d_786_v11_).length(0))
            (d_786_v11_)[index116_] = (self).f31
            d_787_v12_: _dafny.MultiSet
            d_787_v12_ = _dafny.MultiSet([-823, (self).f32, (self).f32])
            d_788_v13_: _dafny.Array
            nw131_ = _dafny.Array(None, 5)
            nw131_[int(0)] = ((d_787_v12_).intersection(d_787_v12_)).cardinality
            nw131_[int(1)] = (self).f32
            nw131_[int(2)] = (0) - ((self).f32)
            nw131_[int(3)] = 592
            nw131_[int(4)] = -12
            d_788_v13_ = nw131_
            index117_ = default__.safeIndex(592, (d_788_v13_).length(0))
            (d_788_v13_)[index117_] = ((self).f32) - (439)
            index118_ = default__.safeIndex(592, (d_788_v13_).length(0))
            (d_788_v13_)[index118_] = (default__.fm7(p0, (d_786_v11_)[default__.safeIndex(888, (d_786_v11_).length(0))], globalState)) + (len((_dafny.SeqWithoutIsStrInference([self.f35, p0])) + (d_771_v4_)))
        elif True:
            (globalState).f28 = (self).f31
            d_789_v14_: D3
            d_789_v14_ = D3_DC10(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_790_i3_ in range(default__.abs(455))]), (self).f32, p0)
            r3 = ((d_789_v14_).cf14 if self.f35 else (self).f32)
            d_791_v15_: str
            d_791_v15_ = _dafny.CodePoint('y')
            (globalState).f28 = (_dafny.MultiSet([d_791_v15_])).issubset(_dafny.MultiSet([d_791_v15_, d_791_v15_]))
            d_771_v4_ = d_771_v4_
            r3 = len(d_769_v2_)
        d_792_v16_: _dafny.Array
        nw132_ = _dafny.Array(None, 22)
        nw132_[int(0)] = not(p0)
        nw132_[int(1)] = (self).f31
        nw132_[int(2)] = p0
        nw132_[int(3)] = p0
        nw132_[int(4)] = self.f35
        nw132_[int(5)] = not(p0)
        nw132_[int(6)] = p0
        nw132_[int(7)] = (self).f31
        nw132_[int(8)] = not(False)
        nw132_[int(9)] = (self).f31
        nw132_[int(10)] = (self).f31
        nw132_[int(11)] = self.f35
        nw132_[int(12)] = self.f35
        nw132_[int(13)] = True
        nw132_[int(14)] = not((self).f31)
        nw132_[int(15)] = default__.fm8(p0, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_793_i4_ in range(default__.abs(79))])), globalState)
        nw132_[int(16)] = p0
        nw132_[int(17)] = p0
        nw132_[int(18)] = p0
        nw132_[int(19)] = self.f35
        nw132_[int(20)] = p0
        nw132_[int(21)] = (self).f31
        d_792_v16_ = nw132_
        d_794_v17_: _dafny.Seq
        d_794_v17_ = _dafny.SeqWithoutIsStrInference([d_792_v16_, d_792_v16_])
        (globalState).f20 = (d_794_v17_)[default__.safeIndex((0) - (default__.safeModuloInt((self).f32, (self).f32)), len(d_794_v17_))]
        d_795_v18_: _dafny.Map
        d_795_v18_ = _dafny.Map({(self).f31: (self).f31})
        d_796_v19_: _dafny.Array
        nw133_ = _dafny.Array(None, 1)
        nw133_[int(0)] = d_795_v18_
        d_796_v19_ = nw133_
        d_797_v20_: D10
        d_797_v20_ = D10_DC27(_dafny.Map({(self).f32: _dafny.Map({(self).f32: (self).f31})}))
        d_798_v21_: T0
        nw134_ = C3()
        nw134_.ctor__(d_796_v19_, (d_797_v20_).cf50, (self).f31)
        d_798_v21_ = nw134_
        d_798_v21_ = d_798_v21_
        d_799_v22_: _dafny.Map
        d_799_v22_ = _dafny.Map({(self).f32: d_771_v4_})
        d_771_v4_ = (((d_799_v22_)[(self).f32] if ((self).f32) in (d_799_v22_) else d_771_v4_)) + (d_771_v4_)
        (globalState).f16 = ((d_795_v18_)[((self).f32) >= ((self).f32)] if (((self).f32) >= ((self).f32)) in (d_795_v18_) else self.f35)
        r0 = (self).f32
        d_800_v23_: _dafny.Array
        nw135_ = _dafny.Array(_dafny.MultiSet({}), 8)
        d_800_v23_ = nw135_
        r1 = d_800_v23_
        r2 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsi"))
        r3 = (d_772_v5_)[default__.safeIndex((self).f32, len(d_772_v5_))]
        return r0, r1, r2, r3

    @property
    def f34(self):
        return self._f34

class C5:
    def  __init__(self):
        self._f29: bool = False
        self._f30: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f29, f30):
        (self)._f29 = f29
        (self)._f30 = f30

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        if (self).f29:
            d_801_v0_: int
            d_801_v0_ = 351
            d_802_v1_: str
            d_802_v1_ = _dafny.CodePoint('r')
            d_801_v0_ = (0) - (default__.fm1(d_802_v1_, globalState))
            d_803_v2_: _dafny.Seq
            d_803_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
            d_804_v3_: _dafny.Seq
            d_804_v3_ = _dafny.SeqWithoutIsStrInference([(((d_803_v2_).set(default__.safeIndex((0) - (p1), len(d_803_v2_)), d_802_v1_)) + (d_803_v2_)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([826])), len(((d_803_v2_).set(default__.safeIndex((0) - (p1), len(d_803_v2_)), d_802_v1_)) + (d_803_v2_))), d_802_v1_)])
            d_804_v3_ = (d_804_v3_) + (d_804_v3_)
            d_805_v4_: _dafny.Array
            def lambda29_(d_806_v0_):
                def lambda30_(d_807_i0_):
                    return (d_807_i0_) + (d_806_v0_)

                return lambda30_

            init16_ = lambda29_(d_801_v0_)
            nw136_ = _dafny.Array(None, 14)
            for i0_16_ in range(nw136_.length(0)):
                nw136_[i0_16_] = init16_(i0_16_)
            d_805_v4_ = nw136_
            index119_ = default__.safeIndex(133, (d_805_v4_).length(0))
            (d_805_v4_)[index119_] = d_801_v0_
            index120_ = default__.safeIndex(133, (d_805_v4_).length(0))
            (d_805_v4_)[index120_] = p1
            d_801_v0_ = (d_805_v4_)[default__.safeIndex(133, (d_805_v4_).length(0))]
            index121_ = default__.safeIndex(133, (d_805_v4_).length(0))
            (d_805_v4_)[index121_] = p1
        elif True:
            d_808_v5_: _dafny.MultiSet
            d_808_v5_ = _dafny.MultiSet([p1])
            default__.m0(((d_808_v5_).cardinality if (self).f30 else p1), globalState)
            d_809_v6_: _dafny.Map
            d_809_v6_ = _dafny.Map({p1: (self).f30})
            if not (((d_809_v6_)[p1] if (p1) in (d_809_v6_) else (self).f29)) or ((self).f30):
                d_810_v7_: _dafny.MultiSet
                d_810_v7_ = _dafny.MultiSet([(self).f29])
                d_811_v8_: _dafny.Seq
                d_811_v8_ = _dafny.SeqWithoutIsStrInference([d_810_v7_, d_810_v7_])
                d_812_v9_: _dafny.Map
                d_812_v9_ = _dafny.Map({len(d_811_v8_): len(default__.fm2((self).f29, p1, (self).f30, (self).f29, globalState))})
                d_813_v10_: _dafny.Seq
                d_813_v10_ = _dafny.SeqWithoutIsStrInference([d_812_v9_])
                d_813_v10_ = d_813_v10_
                (globalState).f28 = (p1) == (p1)
                d_814_v11_: int
                d_814_v11_ = 497
                d_815_v12_: _dafny.Set
                d_815_v12_ = _dafny.Set({(self).f30, (self).f29})
                d_816_v13_: _dafny.Set
                d_816_v13_ = _dafny.Set({410, len(d_815_v12_), d_814_v11_})
                d_817_v14_: _dafny.Seq
                d_817_v14_ = _dafny.SeqWithoutIsStrInference([(self).f29, True])
                d_814_v11_ = (len(d_816_v13_)) + (len(d_817_v14_))
                d_818_v15_: _dafny.Map
                d_818_v15_ = _dafny.Map({(True if default__.fm3(globalState) else default__.fm3(globalState)): p2})
                d_818_v15_ = (d_818_v15_).set((self).f30, p2)
                d_819_v16_: _dafny.Array
                nw137_ = _dafny.Array(int(0), 17)
                d_819_v16_ = nw137_
                index122_ = default__.safeIndex(117, (d_819_v16_).length(0))
                (d_819_v16_)[index122_] = p1
                index123_ = default__.safeIndex(117, (d_819_v16_).length(0))
                (d_819_v16_)[index123_] = (0) - (p1)
            elif True:
                (globalState).f28 = (p1) == (p1)
                d_820_v17_: _dafny.MultiSet
                d_820_v17_ = _dafny.MultiSet([default__.fm3(globalState), (self).f30, (self).f29, (self).f29, (self).f29])
                d_821_v18_: _dafny.Seq
                d_821_v18_ = _dafny.SeqWithoutIsStrInference([d_820_v17_])
                d_822_v19_: D3
                d_822_v19_ = D3_DC9(d_821_v18_)
                d_823_v20_: D0
                d_823_v20_ = D0_DC2(True)
                d_824_v21_: _dafny.Array
                nw138_ = _dafny.Array(None, 12)
                nw138_[int(0)] = d_821_v18_
                nw138_[int(1)] = d_821_v18_
                nw138_[int(2)] = d_821_v18_
                nw138_[int(3)] = (d_821_v18_) + (default__.fm4((self).f30, 700, globalState))
                nw138_[int(4)] = (default__.fm4(not((self).f30), p1, globalState)).set(default__.safeIndex(p1, len(default__.fm4(not((self).f30), p1, globalState))), d_820_v17_)
                nw138_[int(5)] = default__.fm4((self).f30, p1, globalState)
                nw138_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_820_v17_ for d_825_i1_ in range(default__.abs(654))])) + (d_821_v18_)
                nw138_[int(7)] = (d_822_v19_).cf12
                nw138_[int(8)] = default__.fm4((d_823_v20_).cf4, -70, globalState)
                nw138_[int(9)] = (_dafny.SeqWithoutIsStrInference([d_820_v17_, d_820_v17_])) + (d_821_v18_)
                nw138_[int(10)] = d_821_v18_
                nw138_[int(11)] = (d_821_v18_) + (d_821_v18_)
                d_824_v21_ = nw138_
                index124_ = default__.safeIndex(438, (d_824_v21_).length(0))
                (d_824_v21_)[index124_] = (d_821_v18_).set(default__.safeIndex(p1, len(d_821_v18_)), d_820_v17_)
                index125_ = default__.safeIndex(438, (d_824_v21_).length(0))
                (d_824_v21_)[index125_] = _dafny.SeqWithoutIsStrInference([d_820_v17_, d_820_v17_, (d_820_v17_).intersection(d_820_v17_), d_820_v17_, d_820_v17_])
                d_826_v22_: _dafny.Seq
                d_826_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fudavpasx"))
                d_827_v23_: _dafny.Set
                d_827_v23_ = _dafny.Set({(self).f30})
                d_828_v24_: D2
                d_828_v24_ = D2_DC7((d_826_v22_)[default__.safeIndex(p1, len(d_826_v22_))], d_827_v23_)
                d_829_v25_: _dafny.Array
                nw139_ = _dafny.Array(int(0), 2)
                d_829_v25_ = nw139_
                rhs124_ = ((self).f30) or ((self).f29)
                rhs125_ = d_828_v24_
                rhs126_ = (-849) < ((p1) + ((0) - (p1)))
                rhs127_ = (default__.safeModuloInt(p1, p1)) != (p1)
                rhs128_ = d_829_v25_
                lhs76_ = globalState
                lhs77_ = globalState
                lhs78_ = globalState
                lhs76_.f28 = rhs124_
                d_828_v24_ = rhs125_
                lhs77_.f16 = rhs126_
                r0 = rhs127_
                lhs78_.f24 = rhs128_
                d_830_v26_: D2
                d_830_v26_ = D2_DC6((p1 if (self).f30 else p1))
                d_830_v26_ = d_830_v26_
                d_831_v27_: int
                d_831_v27_ = -488
                d_831_v27_ = d_831_v27_
            d_832_v28_: str
            d_832_v28_ = _dafny.CodePoint('d')
            d_833_v29_: _dafny.Array
            nw140_ = _dafny.Array(False, 13)
            d_833_v29_ = nw140_
            d_834_v30_: _dafny.Map
            d_834_v30_ = _dafny.Map({d_832_v28_: d_833_v29_})
            d_835_v31_: _dafny.Seq
            d_835_v31_ = _dafny.SeqWithoutIsStrInference([d_833_v29_])
            d_834_v30_ = (d_834_v30_).set(d_832_v28_, (d_835_v31_)[default__.safeIndex(p1, len(d_835_v31_))])
            (globalState).f16 = (self).f29
            d_836_v32_: int
            d_836_v32_ = 277
            d_836_v32_ = p1
        d_837_v33_: _dafny.Array
        nw141_ = _dafny.Array(_dafny.Seq({}), 19)
        d_837_v33_ = nw141_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_837_v33_).length(0)):
            d_838_i2_: int = guard_loop_4_
            if (True) and (((0) <= (d_838_i2_)) and ((d_838_i2_) < ((d_837_v33_).length(0)))):
                def iife68_():
                    coll24_ = _dafny.Set()
                    compr_24_: D2
                    for compr_24_ in (_dafny.Map({D2_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ho")))): p1})).keys.Elements:
                        d_839_v34_: D2 = compr_24_
                        if (d_839_v34_) in (_dafny.Map({D2_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ho")))): p1})):
                            coll24_ = coll24_.union(_dafny.Set([d_839_v34_]))
                    return _dafny.Set(coll24_)
                (d_837_v33_)[(d_838_i2_)] = _dafny.SeqWithoutIsStrInference([(self).f30, (self).f29, False, (_dafny.MultiSet([(self).f29, (self).f29])).issubset(_dafny.MultiSet([(self).f30])), ((_dafny.Set({D2_DC6(p1)}))).isdisjoint(iife68_()
                )])
        d_840_v35_: str
        d_840_v35_ = _dafny.CodePoint('u')
        d_841_v36_: _dafny.Set
        d_841_v36_ = _dafny.Set({(self).f30, True})
        d_842_v37_: D2
        d_842_v37_ = D2_DC7(d_840_v35_, d_841_v36_)
        pat_let_tv19_ = d_841_v36_
        def iife69_(_pat_let22_0):
            def iife70_(d_843_dt__update__tmp_h0_):
                def iife71_(_pat_let23_0):
                    def iife72_(d_844_dt__update_hcf10_h0_):
                        return D2_DC7((d_843_dt__update__tmp_h0_).cf9, d_844_dt__update_hcf10_h0_)
                    return iife72_(_pat_let23_0)
                return iife71_(pat_let_tv19_)
            return iife70_(_pat_let22_0)
        d_840_v35_ = (iife69_(d_842_v37_)).cf9
        d_845_v38_: _dafny.Seq
        d_845_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjubcpq"))
        d_846_v39_: D3
        d_846_v39_ = D3_DC10(d_845_v38_, p1, (self).f30)
        d_847_v40_: D3
        d_847_v40_ = D3_DC11(d_846_v39_)
        d_848_v41_: D3
        d_848_v41_ = D3_DC11(d_847_v40_)
        source13_ = d_848_v41_
        if source13_.is_DC10:
            d_849___mcc_h0_ = source13_.cf13
            d_850___mcc_h1_ = source13_.cf14
            d_851___mcc_h2_ = source13_.cf15
            d_852_cf15_ = d_851___mcc_h2_
            d_853_cf14_ = d_850___mcc_h1_
            d_854_cf13_ = d_849___mcc_h0_
            d_853_cf14_ = d_853_cf14_
            d_855_v42_: _dafny.Array
            nw142_ = _dafny.Array(D0.default()(), 1)
            d_855_v42_ = nw142_
            d_856_v43_: _dafny.Set
            d_856_v43_ = _dafny.Set({d_855_v42_})
            d_852_cf15_ = (d_856_v43_).isdisjoint(d_856_v43_)
            d_857_v44_: _dafny.Array
            def lambda31_(d_858_cf14_):
                def lambda32_(d_859_i3_):
                    return default__.safeDivisionInt(d_859_i3_, d_858_cf14_)

                return lambda32_

            init17_ = lambda31_(d_853_cf14_)
            nw143_ = _dafny.Array(None, 28)
            for i0_17_ in range(nw143_.length(0)):
                nw143_[i0_17_] = init17_(i0_17_)
            d_857_v44_ = nw143_
            d_860_v45_: _dafny.Array
            d_860_v45_ = d_857_v44_
            source14_ = d_860_v45_
            d_861___mcc_h5_ = source14_
            d_862_cf7_ = d_861___mcc_h5_
            d_863_v46_: _dafny.Array
            def lambda33_(d_864_cf15_):
                def lambda34_(d_865_i4_):
                    return not(not(d_864_cf15_))

                return lambda34_

            init18_ = lambda33_(d_852_cf15_)
            nw144_ = _dafny.Array(None, 4)
            for i0_18_ in range(nw144_.length(0)):
                nw144_[i0_18_] = init18_(i0_18_)
            d_863_v46_ = nw144_
            (globalState).f27 = d_863_v46_
            d_866_v47_: _dafny.Seq
            d_866_v47_ = _dafny.SeqWithoutIsStrInference([d_845_v38_])
            d_867_v48_: D3
            d_867_v48_ = D3_DC10((d_866_v47_)[default__.safeIndex(p1, len(d_866_v47_))], p1, d_852_cf15_)
            d_868_v49_: _dafny.Seq
            d_868_v49_ = _dafny.SeqWithoutIsStrInference([True])
            (globalState).f16 = ((d_867_v48_ if d_852_cf15_ else D3_DC10(d_845_v38_, len(d_868_v49_), d_852_cf15_))).cf15
            d_869_v50_: _dafny.Map
            d_869_v50_ = _dafny.Map({d_845_v38_: False})
            d_870_v51_: _dafny.MultiSet
            d_870_v51_ = _dafny.MultiSet([(self).f29])
            d_871_v52_: _dafny.Map
            d_871_v52_ = _dafny.Map({d_852_cf15_: default__.fm0(d_870_v51_, globalState)})
            d_869_v50_ = (d_869_v50_).set(((d_871_v52_)[(self).f29] if ((self).f29) in (d_871_v52_) else d_845_v38_), (self).f29)
            d_852_cf15_ = (len((_dafny.SeqWithoutIsStrInference([(0) - ((0) - (p1)) for d_872_i5_ in range(default__.abs(46))])) + (p0))) != ((0) - (len(d_845_v38_)))
            d_873_v53_: _dafny.MultiSet
            d_873_v53_ = _dafny.MultiSet([p0])
            d_874_v54_: _dafny.Map
            d_874_v54_ = _dafny.Map({default__.fm5((self).f29, d_853_cf14_, globalState): (self).f29})
            d_875_v55_: _dafny.MultiSet
            d_875_v55_ = _dafny.MultiSet([(self).f30])
            d_876_v57_: _dafny.Set
            d_876_v57_ = _dafny.Set({d_840_v35_})
            def iife73_():
                coll25_ = _dafny.Map()
                compr_25_: str
                for compr_25_ in (d_876_v57_).Elements:
                    d_877_v56_: str = compr_25_
                    if (d_877_v56_) in (d_876_v57_):
                        coll25_[d_877_v56_] = False
                return _dafny.Map(coll25_)
            rhs129_ = (self).f29
            rhs130_ = _dafny.MultiSet([(p0) + (p0)])
            rhs131_ = (p1) - (default__.safeModuloInt(len(_dafny.Map({d_875_v55_: (self).f29})), (p0)[default__.safeIndex((0) - ((0) - (p1)), len(p0))]))
            rhs132_ = iife73_()

            rhs133_ = (387) < (d_853_cf14_)
            lhs79_ = globalState
            lhs80_ = globalState
            lhs79_.f28 = rhs129_
            d_873_v53_ = rhs130_
            d_853_cf14_ = rhs131_
            d_874_v54_ = rhs132_
            lhs80_.f28 = rhs133_
        elif source13_.is_DC9:
            d_878___mcc_h3_ = source13_.cf12
            d_879_cf12_ = d_878___mcc_h3_
            d_880_v58_: _dafny.Map
            d_880_v58_ = _dafny.Map({p1: (self).f30})
            d_881_v59_: _dafny.MultiSet
            d_881_v59_ = _dafny.MultiSet([(self).f30, False])
            d_882_v60_: _dafny.MultiSet
            d_882_v60_ = _dafny.MultiSet([p1, len(d_880_v58_), p1, (d_881_v59_).cardinality, 364])
            d_883_v61_: _dafny.Map
            d_883_v61_ = _dafny.Map({(d_882_v60_).cardinality: (self).f29})
            d_884_v62_: D2
            d_884_v62_ = D2_DC6(p1)
            d_885_v63_: _dafny.Set
            d_885_v63_ = _dafny.Set({d_884_v62_, d_884_v62_})
            d_886_v64_: _dafny.Set
            d_886_v64_ = d_885_v63_
            d_887_v65_: _dafny.Seq
            d_887_v65_ = _dafny.SeqWithoutIsStrInference([d_886_v64_, d_886_v64_, d_886_v64_, d_886_v64_, _dafny.Set({d_884_v62_})])
            d_888_v66_: _dafny.Map
            d_888_v66_ = _dafny.Map({len((d_883_v61_) | (d_883_v61_)): d_887_v65_})
            d_889_v67_: _dafny.Map
            d_889_v67_ = _dafny.Map({(self).f30: (_dafny.Map({(0) - (p1): d_887_v65_})).set(len(p0), d_887_v65_)})
            d_888_v66_ = ((d_889_v67_)[(self).f30] if ((self).f30) in (d_889_v67_) else d_888_v66_)
            (globalState).f28 = not ((self).f29) or ((self).f30)
            r0 = ((453) + (p1)) != (p1)
            rhs134_ = (self).f30
            rhs135_ = (self).f29
            lhs81_ = globalState
            lhs82_ = globalState
            lhs81_.f16 = rhs134_
            lhs82_.f28 = rhs135_
        elif True:
            d_890___mcc_h4_ = source13_.cf16
            d_891_cf16_ = d_890___mcc_h4_
            d_892_v68_: int
            d_892_v68_ = 435
            d_893_v69_: _dafny.MultiSet
            d_893_v69_ = _dafny.MultiSet([(self).f30])
            d_894_v70_: _dafny.Map
            d_894_v70_ = _dafny.Map({d_893_v69_: (p0)[default__.safeIndex(d_892_v68_, len(p0))]})
            d_892_v68_ = (0) - ((d_892_v68_ if (self).f30 else ((d_894_v70_)[d_893_v69_] if (d_893_v69_) in (d_894_v70_) else p1)))
            d_895_v71_: _dafny.Seq
            d_895_v71_ = _dafny.SeqWithoutIsStrInference([p2, p2])
            d_896_v72_: _dafny.Map
            d_896_v72_ = _dafny.Map({p1: True})
            d_897_v73_: _dafny.Array
            nw145_ = _dafny.Array(None, 22)
            nw145_[int(0)] = (self).f29
            nw145_[int(1)] = False
            nw145_[int(2)] = False
            nw145_[int(3)] = (self).f30
            nw145_[int(4)] = (self).f30
            nw145_[int(5)] = not(True)
            nw145_[int(6)] = (self).f29
            nw145_[int(7)] = (self).f29
            nw145_[int(8)] = (self).f30
            nw145_[int(9)] = (self).f29
            nw145_[int(10)] = (self).f29
            nw145_[int(11)] = not((self).f29)
            nw145_[int(12)] = (self).f30
            nw145_[int(13)] = ((d_896_v72_)[d_892_v68_] if (d_892_v68_) in (d_896_v72_) else (self).f30)
            nw145_[int(14)] = (self).f29
            nw145_[int(15)] = (self).f29
            nw145_[int(16)] = ((d_896_v72_)[d_892_v68_] if (d_892_v68_) in (d_896_v72_) else (self).f29)
            nw145_[int(17)] = (self).f30
            nw145_[int(18)] = (self).f29
            nw145_[int(19)] = (self).f29
            nw145_[int(20)] = True
            nw145_[int(21)] = True
            d_897_v73_ = nw145_
            d_898_v74_: _dafny.Seq
            d_898_v74_ = _dafny.SeqWithoutIsStrInference([(self).f30])
            d_899_v75_: _dafny.Map
            d_899_v75_ = _dafny.Map({d_898_v74_: d_897_v73_})
            d_900_v76_: _dafny.Set
            d_900_v76_ = _dafny.Set({d_892_v68_})
            d_901_v77_: D5
            d_901_v77_ = D5_DC14(d_900_v76_, p1, (self).f30, d_897_v73_)
            d_902_v78_: _dafny.Array
            nw146_ = _dafny.Array(None, 29)
            nw146_[int(0)] = d_897_v73_
            nw146_[int(1)] = d_897_v73_
            nw146_[int(2)] = d_897_v73_
            nw146_[int(3)] = d_897_v73_
            nw146_[int(4)] = d_897_v73_
            nw146_[int(5)] = d_897_v73_
            nw146_[int(6)] = d_897_v73_
            nw146_[int(7)] = d_897_v73_
            nw146_[int(8)] = d_897_v73_
            nw146_[int(9)] = d_897_v73_
            nw146_[int(10)] = d_897_v73_
            nw146_[int(11)] = d_897_v73_
            nw146_[int(12)] = d_897_v73_
            nw146_[int(13)] = d_897_v73_
            nw146_[int(14)] = d_897_v73_
            nw146_[int(15)] = ((d_899_v75_)[d_898_v74_] if (d_898_v74_) in (d_899_v75_) else d_897_v73_)
            nw146_[int(16)] = d_897_v73_
            nw146_[int(17)] = (d_901_v77_).cf22
            nw146_[int(18)] = d_897_v73_
            nw146_[int(19)] = d_897_v73_
            nw146_[int(20)] = d_897_v73_
            nw146_[int(21)] = d_897_v73_
            nw146_[int(22)] = d_897_v73_
            nw146_[int(23)] = d_897_v73_
            nw146_[int(24)] = d_897_v73_
            nw146_[int(25)] = d_897_v73_
            nw146_[int(26)] = d_897_v73_
            nw146_[int(27)] = d_897_v73_
            nw146_[int(28)] = d_897_v73_
            d_902_v78_ = nw146_
            d_903_v79_: _dafny.Map
            d_903_v79_ = _dafny.Map({(self).f30: d_892_v68_})
            d_904_v81_: C4
            nw147_ = C4()
            def iife74_():
                coll26_ = _dafny.Set()
                compr_26_: _dafny.Map
                for compr_26_ in (_dafny.MultiSet([d_903_v79_, d_903_v79_])).Elements:
                    d_905_v80_: _dafny.Map = compr_26_
                    if (d_905_v80_) in (_dafny.MultiSet([d_903_v79_, d_903_v79_])):
                        coll26_ = coll26_.union(_dafny.Set([d_905_v80_]))
                return _dafny.Set(coll26_)
            nw147_.ctor__(d_895_v71_, (self).f30, p1, d_902_v78_, not((iife74_()
            ).isdisjoint(_dafny.Set({d_903_v79_}))))
            d_904_v81_ = nw147_
            index126_ = default__.safeIndex(848, (d_897_v73_).length(0))
            (d_897_v73_)[index126_] = not((self).f30)
            index127_ = default__.safeIndex(264, (d_902_v78_).length(0))
            (d_902_v78_)[index127_] = d_897_v73_
            d_906_v82_: _dafny.MultiSet
            d_906_v82_ = _dafny.MultiSet([280, d_892_v68_, d_892_v68_])
            d_907_v83_: _dafny.Array
            nw148_ = _dafny.Array(None, 5)
            nw148_[int(0)] = d_892_v68_
            nw148_[int(1)] = p1
            nw148_[int(2)] = len(d_845_v38_)
            nw148_[int(3)] = p1
            nw148_[int(4)] = len(d_845_v38_)
            d_907_v83_ = nw148_
            d_908_v84_: _dafny.Map
            d_908_v84_ = _dafny.Map({d_907_v83_: d_897_v73_})
            index128_ = default__.safeIndex(848, (d_897_v73_).length(0))
            index129_ = default__.safeIndex(264, (d_902_v78_).length(0))
            rhs136_ = (d_840_v35_) not in (d_845_v38_)
            rhs137_ = (d_906_v82_).isdisjoint(d_906_v82_)
            rhs138_ = ((d_908_v84_)[d_907_v83_] if (d_907_v83_) in (d_908_v84_) else d_897_v73_)
            rhs139_ = (default__.safeModuloInt(len(p0), 6)) > ((0) - (default__.safeDivisionInt(d_892_v68_, 484)))
            rhs140_ = d_845_v38_
            lhs83_ = d_897_v73_
            lhs84_ = default__.safeIndex(848, (d_897_v73_).length(0))
            lhs85_ = globalState
            lhs86_ = d_902_v78_
            lhs87_ = default__.safeIndex(264, (d_902_v78_).length(0))
            lhs88_ = d_904_v81_
            lhs83_[lhs84_] = rhs136_
            lhs85_.f16 = rhs137_
            lhs86_[lhs87_] = rhs138_
            lhs88_.f35 = rhs139_
            d_845_v38_ = rhs140_
            d_892_v68_ = p1
        d_909_v85_: _dafny.Map
        d_909_v85_ = _dafny.Map({p1: default__.fm3(globalState)})
        d_910_v86_: _dafny.Map
        d_910_v86_ = _dafny.Map({(self).f30: (self).f29})
        d_911_i6_: int
        d_911_i6_ = 0
        with _dafny.label("9"):
            while ((d_909_v85_)[p1] if (p1) in (d_909_v85_) else (len(d_845_v38_)) <= (len(d_910_v86_))):
                with _dafny.c_label("9"):
                    if (d_911_i6_) >= (100):
                        raise _dafny.Break("9")
                    d_911_i6_ = (d_911_i6_) + (1)
                    d_912_v87_: _dafny.Array
                    nw149_ = _dafny.Array(D6.default()(), 7)
                    d_912_v87_ = nw149_
                    d_913_v88_: _dafny.Array
                    def lambda35_(d_914_p1_):
                        def lambda36_(d_915_i7_):
                            return _dafny.Set({d_914_p1_, d_914_p1_})

                        return lambda36_

                    init19_ = lambda35_(p1)
                    nw150_ = _dafny.Array(None, 18)
                    for i0_19_ in range(nw150_.length(0)):
                        nw150_[i0_19_] = init19_(i0_19_)
                    d_913_v88_ = nw150_
                    d_916_v89_: _dafny.MultiSet
                    d_916_v89_ = _dafny.MultiSet([p1])
                    d_917_v90_: D6
                    d_917_v90_ = D6_DC17((self).f30, d_913_v88_, (d_916_v89_).cardinality)
                    index130_ = default__.safeIndex(956, (d_912_v87_).length(0))
                    (d_912_v87_)[index130_] = d_917_v90_
                    index131_ = default__.safeIndex(956, (d_912_v87_).length(0))
                    (d_912_v87_)[index131_] = d_917_v90_
                    d_918_v91_: D0
                    d_918_v91_ = D0_DC0(d_837_v33_)
                    d_919_v92_: D0
                    d_919_v92_ = D0_DC4(d_918_v91_)
                    d_920_v93_: _dafny.Map
                    d_920_v93_ = _dafny.Map({d_919_v92_: default__.fm1(d_840_v35_, globalState)})
                    (globalState).f28 = ((d_919_v92_ if (self).f29 else d_919_v92_)) not in (d_920_v93_)
                    d_921_v94_: _dafny.MultiSet
                    d_921_v94_ = _dafny.MultiSet([(self).f30])
                    d_922_v95_: _dafny.Seq
                    d_922_v95_ = _dafny.SeqWithoutIsStrInference([(self).f30, (self).f29, (self).f30, (self).f30])
                    d_923_v96_: _dafny.Seq
                    d_923_v96_ = _dafny.SeqWithoutIsStrInference([d_909_v85_, (d_909_v85_) | (d_909_v85_), _dafny.Map({p1: (self).f30}), _dafny.Map({(d_921_v94_).cardinality: (d_922_v95_)[default__.safeIndex(111, len(d_922_v95_))]}), d_909_v85_])
                    d_909_v85_ = (d_923_v96_)[default__.safeIndex((0) - (p1), len(d_923_v96_))]
                    d_924_v97_: _dafny.Array
                    nw151_ = _dafny.Array(False, 28)
                    d_924_v97_ = nw151_
                    index132_ = default__.safeIndex(556, (d_924_v97_).length(0))
                    (d_924_v97_)[index132_] = default__.fm10(d_840_v35_, False, globalState)
                    index133_ = default__.safeIndex(556, (d_924_v97_).length(0))
                    (d_924_v97_)[index133_] = (self).f29
                    pass
            pass
        d_925_v98_: _dafny.Seq
        d_925_v98_ = _dafny.SeqWithoutIsStrInference([d_845_v38_, d_845_v38_])
        d_926_v99_: D9
        d_926_v99_ = D9_DC23(d_925_v98_)
        source15_ = d_926_v99_
        if source15_.is_DC24:
            d_927___mcc_h6_ = source15_.cf40
            d_928_cf40_ = d_927___mcc_h6_
            d_929_v100_: _dafny.Seq
            d_929_v100_ = _dafny.SeqWithoutIsStrInference([p2])
            d_930_v101_: _dafny.Array
            nw152_ = _dafny.Array(None, 26)
            nw152_[int(0)] = (self).f29
            nw152_[int(1)] = (self).f30
            nw152_[int(2)] = (self).f29
            nw152_[int(3)] = (self).f29
            nw152_[int(4)] = (self).f30
            nw152_[int(5)] = (self).f30
            nw152_[int(6)] = (self).f30
            nw152_[int(7)] = (self).f29
            nw152_[int(8)] = True
            nw152_[int(9)] = (self).f29
            nw152_[int(10)] = (self).f29
            nw152_[int(11)] = (self).f30
            nw152_[int(12)] = (self).f30
            nw152_[int(13)] = (self).f29
            nw152_[int(14)] = (self).f29
            nw152_[int(15)] = (self).f30
            nw152_[int(16)] = (self).f30
            nw152_[int(17)] = (self).f29
            nw152_[int(18)] = (self).f29
            nw152_[int(19)] = (self).f30
            nw152_[int(20)] = (self).f29
            nw152_[int(21)] = (self).f30
            nw152_[int(22)] = not((self).f29)
            nw152_[int(23)] = True
            nw152_[int(24)] = False
            nw152_[int(25)] = (self).f29
            d_930_v101_ = nw152_
            d_931_v102_: _dafny.Seq
            d_931_v102_ = _dafny.SeqWithoutIsStrInference([d_930_v101_])
            d_932_v103_: _dafny.Array
            nw153_ = _dafny.Array(None, 28)
            nw153_[int(0)] = d_930_v101_
            nw153_[int(1)] = d_930_v101_
            nw153_[int(2)] = d_930_v101_
            nw153_[int(3)] = d_930_v101_
            nw153_[int(4)] = d_930_v101_
            nw153_[int(5)] = d_930_v101_
            nw153_[int(6)] = d_930_v101_
            nw153_[int(7)] = (d_931_v102_)[default__.safeIndex(len(p0), len(d_931_v102_))]
            nw153_[int(8)] = d_930_v101_
            nw153_[int(9)] = d_930_v101_
            nw153_[int(10)] = d_930_v101_
            nw153_[int(11)] = d_930_v101_
            nw153_[int(12)] = d_930_v101_
            nw153_[int(13)] = d_930_v101_
            nw153_[int(14)] = d_930_v101_
            nw153_[int(15)] = d_930_v101_
            nw153_[int(16)] = d_930_v101_
            nw153_[int(17)] = d_930_v101_
            nw153_[int(18)] = d_930_v101_
            nw153_[int(19)] = d_930_v101_
            nw153_[int(20)] = d_930_v101_
            nw153_[int(21)] = d_930_v101_
            nw153_[int(22)] = (d_931_v102_)[default__.safeIndex(p1, len(d_931_v102_))]
            nw153_[int(23)] = d_930_v101_
            nw153_[int(24)] = d_930_v101_
            nw153_[int(25)] = d_930_v101_
            nw153_[int(26)] = d_930_v101_
            nw153_[int(27)] = d_930_v101_
            d_932_v103_ = nw153_
            d_933_v104_: C4
            nw154_ = C4()
            nw154_.ctor__(d_929_v100_, ((d_928_cf40_).cardinality) == ((0) - (p1)), len(d_845_v38_), d_932_v103_, (self).f30)
            d_933_v104_ = nw154_
            d_934_v105_: int
            d_934_v105_ = 570
            d_934_v105_ = d_934_v105_
            d_935_v106_: _dafny.Array
            nw155_ = _dafny.Array(int(0), 14)
            d_935_v106_ = nw155_
            d_936_v107_: _dafny.Set
            d_936_v107_ = _dafny.Set({d_935_v106_})
            d_936_v107_ = _dafny.Set({d_935_v106_, d_935_v106_})
            d_937_v108_: C2
            nw156_ = C2()
            nw156_.ctor__(default__.fm14(p1, d_934_v105_, 372, (self).f29, globalState), (self).f29)
            d_937_v108_ = nw156_
        elif source15_.is_DC25:
            d_938___mcc_h7_ = source15_.cf41
            d_939___mcc_h8_ = source15_.cf42
            d_940___mcc_h9_ = source15_.cf43
            d_941___mcc_h10_ = source15_.cf44
            d_942_cf44_ = d_941___mcc_h10_
            d_943_cf43_ = d_940___mcc_h9_
            d_944_cf42_ = d_939___mcc_h8_
            d_945_cf41_ = d_938___mcc_h7_
            d_946_v109_: D10
            d_946_v109_ = D10_DC28(p1)
            source16_ = d_946_v109_
            if source16_.is_DC28:
                d_947___mcc_h17_ = source16_.cf51
                d_948_cf51_ = d_947___mcc_h17_
                d_949_v110_: _dafny.Array
                nw157_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_949_v110_ = nw157_
                d_949_v110_ = d_949_v110_
                d_950_v111_: _dafny.Array
                nw158_ = _dafny.Array(int(0), 1)
                d_950_v111_ = nw158_
                d_951_v112_: _dafny.Array
                d_951_v112_ = d_950_v111_
                d_951_v112_ = d_951_v112_
                d_952_v113_: _dafny.Map
                d_952_v113_ = _dafny.Map({(self).f30: default__.safeDivisionInt(d_948_cf51_, d_948_cf51_)})
                d_952_v113_ = (d_952_v113_).set((self).f29, d_948_cf51_)
                d_845_v38_ = d_845_v38_
            elif source16_.is_DC27:
                d_953___mcc_h18_ = source16_.cf50
                d_954_cf50_ = d_953___mcc_h18_
                d_955_v114_: _dafny.Array
                nw159_ = _dafny.Array(None, 3)
                nw159_[int(0)] = 430
                nw159_[int(1)] = -353
                nw159_[int(2)] = -342
                d_955_v114_ = nw159_
                index134_ = default__.safeIndex(357, (d_955_v114_).length(0))
                (d_955_v114_)[index134_] = (d_945_cf41_ if (self).f30 else d_945_cf41_)
                index135_ = default__.safeIndex(357, (d_955_v114_).length(0))
                (d_955_v114_)[index135_] = len((d_845_v38_) + (((d_925_v98_)[default__.safeIndex((0) - (p1), len(d_925_v98_))]) + (d_845_v38_)))
                d_956_v115_: _dafny.Seq
                d_956_v115_ = _dafny.SeqWithoutIsStrInference([True])
                d_956_v115_ = d_956_v115_
                (globalState).f16 = d_942_cf44_
                d_957_v116_: _dafny.Map
                d_957_v116_ = _dafny.Map({(self).f29: p1})
                index136_ = default__.safeIndex(357, (d_955_v114_).length(0))
                (d_955_v114_)[index136_] = ((len(d_957_v116_) if d_943_cf43_ else len(d_845_v38_))) * ((d_945_cf41_) * (default__.fm18(((d_909_v85_)[p1] if (p1) in (d_909_v85_) else (self).f30), (self).f29, globalState)))
            elif True:
                d_958___mcc_h19_ = source16_.cf52
                d_959_cf52_ = d_958___mcc_h19_
                d_942_cf44_ = default__.fm8((d_945_cf41_) <= (d_945_cf41_), -349, globalState)
                d_945_cf41_ = d_945_cf41_
                d_960_v117_: _dafny.Seq
                d_960_v117_ = _dafny.SeqWithoutIsStrInference([d_945_cf41_])
                d_960_v117_ = _dafny.SeqWithoutIsStrInference([d_945_cf41_, default__.fm7(False, (self).f30, globalState), 661])
                d_961_v118_: _dafny.Array
                nw160_ = _dafny.Array(None, 2)
                nw160_[int(0)] = not(d_944_cf42_)
                nw160_[int(1)] = (self).f30
                d_961_v118_ = nw160_
                d_962_v119_: _dafny.Map
                d_962_v119_ = _dafny.Map({p1: d_961_v118_})
                d_963_v120_: D8
                d_963_v120_ = D8_DC20(d_962_v119_)
                rhs141_ = ((0) - (default__.fm1(d_840_v35_, globalState))) > (-520)
                rhs142_ = d_963_v120_
                rhs143_ = (self).f30
                lhs89_ = globalState
                lhs89_.f28 = rhs141_
                d_963_v120_ = rhs142_
                d_943_cf43_ = rhs143_
            d_964_v122_: _dafny.Array
            def lambda37_(d_965_cf43_):
                def lambda38_(d_966_i8_):
                    return d_965_cf43_

                return lambda38_

            init20_ = lambda37_(d_943_cf43_)
            nw161_ = _dafny.Array(None, 23)
            for i0_20_ in range(nw161_.length(0)):
                nw161_[i0_20_] = init20_(i0_20_)
            d_964_v122_ = nw161_
            d_967_v123_: D5
            def iife75_():
                coll27_ = _dafny.Set()
                compr_27_: int
                for compr_27_ in _dafny.IntegerRange(-870, 990):
                    d_968_v121_: int = compr_27_
                    if ((-870) <= (d_968_v121_)) and ((d_968_v121_) < (990)):
                        coll27_ = coll27_.union(_dafny.Set([default__.safeModuloInt(d_968_v121_, p1)]))
                return _dafny.Set(coll27_)
            d_967_v123_ = D5_DC14(iife75_()
, d_945_cf41_, not(d_942_cf44_), d_964_v122_)
            d_969_v124_: _dafny.Map
            d_969_v124_ = _dafny.Map({d_967_v123_: d_840_v35_})
            d_970_v125_: _dafny.Set
            d_970_v125_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouwdevy")))})
            rhs144_ = ((d_969_v124_)[D5_DC14(d_970_v125_, p1, not((self).f29), d_964_v122_)] if (D5_DC14(d_970_v125_, p1, not((self).f29), d_964_v122_)) in (d_969_v124_) else d_840_v35_)
            rhs145_ = (d_945_cf41_) - (default__.fm1(d_840_v35_, globalState))
            rhs146_ = default__.fm3(globalState)
            rhs147_ = True
            d_840_v35_ = rhs144_
            d_945_cf41_ = rhs145_
            d_943_cf43_ = rhs146_
            d_944_cf42_ = rhs147_
            d_971_v126_: _dafny.Map
            d_971_v126_ = _dafny.Map({d_945_cf41_: _dafny.Set({d_945_cf41_, d_945_cf41_})})
            d_972_v130_: _dafny.Array
            nw162_ = _dafny.Array(None, 29)
            nw162_[int(0)] = d_970_v125_
            nw162_[int(1)] = d_970_v125_
            nw162_[int(2)] = d_970_v125_
            nw162_[int(3)] = d_970_v125_
            nw162_[int(4)] = d_970_v125_
            nw162_[int(5)] = _dafny.Set({p1, p1})
            nw162_[int(6)] = d_970_v125_
            nw162_[int(7)] = _dafny.Set({d_945_cf41_, p1, d_945_cf41_})
            nw162_[int(8)] = d_970_v125_
            nw162_[int(9)] = d_970_v125_
            nw162_[int(10)] = d_970_v125_
            nw162_[int(11)] = d_970_v125_
            nw162_[int(12)] = _dafny.Set({d_945_cf41_, p1})
            nw162_[int(13)] = d_970_v125_
            nw162_[int(14)] = (d_967_v123_).cf19
            nw162_[int(15)] = d_970_v125_
            nw162_[int(16)] = d_970_v125_
            nw162_[int(17)] = _dafny.Set({p1, (0) - (d_945_cf41_)})
            nw162_[int(18)] = _dafny.Set({(0) - (-73)})
            def iife76_():
                def iife78_():
                    coll30_ = _dafny.Set()
                    compr_30_: int
                    for compr_30_ in _dafny.IntegerRange(464, 420):
                        d_975_v127_: int = compr_30_
                        if ((464) <= (d_975_v127_)) and ((d_975_v127_) < (420)):
                            coll30_ = coll30_.union(_dafny.Set([default__.safeDivisionInt(d_975_v127_, p1)]))
                    return _dafny.Set(coll30_)
                coll28_ = _dafny.Set()
                def iife77_():
                    coll29_ = _dafny.Set()
                    compr_29_: int
                    for compr_29_ in _dafny.IntegerRange(464, 420):
                        d_973_v127_: int = compr_29_
                        if ((464) <= (d_973_v127_)) and ((d_973_v127_) < (420)):
                            coll29_ = coll29_.union(_dafny.Set([default__.safeDivisionInt(d_973_v127_, p1)]))
                    return _dafny.Set(coll29_)
                compr_28_: int
                for compr_28_ in (iife77_()
                ).Elements:
                    d_974_v128_: int = compr_28_
                    if (d_974_v128_) in (iife78_()
                    ):
                        coll28_ = coll28_.union(_dafny.Set([default__.safeModuloInt(d_974_v128_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))))]))
                return _dafny.Set(coll28_)
            nw162_[int(19)] = ((d_971_v126_)[d_945_cf41_] if (d_945_cf41_) in (d_971_v126_) else iife76_()
            )
            nw162_[int(20)] = d_970_v125_
            nw162_[int(21)] = _dafny.Set({d_945_cf41_})
            nw162_[int(22)] = _dafny.Set({p1})
            nw162_[int(23)] = d_970_v125_
            nw162_[int(24)] = d_970_v125_
            nw162_[int(25)] = default__.fm28(len(d_910_v86_), False, d_942_cf44_, globalState)
            nw162_[int(26)] = d_970_v125_
            nw162_[int(27)] = d_970_v125_
            def iife79_():
                coll31_ = _dafny.Set()
                compr_31_: int
                for compr_31_ in _dafny.IntegerRange(160, -306):
                    d_976_v129_: int = compr_31_
                    if ((160) <= (d_976_v129_)) and ((d_976_v129_) < (-306)):
                        coll31_ = coll31_.union(_dafny.Set([(d_976_v129_) * (d_945_cf41_)]))
                return _dafny.Set(coll31_)
            nw162_[int(28)] = iife79_()
            
            d_972_v130_ = nw162_
            d_977_v131_: D6
            d_977_v131_ = D6_DC17(d_942_cf44_, d_972_v130_, p1)
            d_942_cf44_ = (d_977_v131_).cf25
            d_978_v132_: _dafny.Map
            d_978_v132_ = _dafny.Map({d_840_v35_: (p0)[default__.safeIndex(p1, len(p0))]})
            d_979_v133_: _dafny.MultiSet
            d_979_v133_ = _dafny.MultiSet([p1])
            d_978_v132_ = (d_978_v132_).set(default__.fm20(d_845_v38_, d_979_v133_, False, -38, globalState), default__.safeModuloInt(d_945_cf41_, 150))
        elif source15_.is_DC26:
            d_980___mcc_h11_ = source15_.cf45
            d_981___mcc_h12_ = source15_.cf46
            d_982___mcc_h13_ = source15_.cf47
            d_983___mcc_h14_ = source15_.cf48
            d_984___mcc_h15_ = source15_.cf49
            d_985_cf49_ = d_984___mcc_h15_
            d_986_cf48_ = d_983___mcc_h14_
            d_987_cf47_ = d_982___mcc_h13_
            d_988_cf46_ = d_981___mcc_h12_
            d_989_cf45_ = d_980___mcc_h11_
            if ((d_841_v36_) | (d_841_v36_)) == (_dafny.Set({d_985_cf49_})):
                d_990_v134_: _dafny.MultiSet
                d_990_v134_ = _dafny.MultiSet([not(d_985_cf49_)])
                d_991_v135_: _dafny.Array
                nw163_ = _dafny.Array(None, 17)
                nw163_[int(0)] = True
                nw163_[int(1)] = d_985_cf49_
                nw163_[int(2)] = True
                nw163_[int(3)] = ((self).f29 if d_985_cf49_ else (self).f30)
                nw163_[int(4)] = (d_989_cf45_) == (((d_910_v86_)[(self).f30] if ((self).f30) in (d_910_v86_) else d_989_cf45_))
                nw163_[int(5)] = (self).f30
                nw163_[int(6)] = d_985_cf49_
                nw163_[int(7)] = (self).f30
                nw163_[int(8)] = d_989_cf45_
                nw163_[int(9)] = d_989_cf45_
                nw163_[int(10)] = (self).f29
                nw163_[int(11)] = (_dafny.MultiSet([d_987_cf47_, p1])).issubset(_dafny.MultiSet([d_987_cf47_]))
                nw163_[int(12)] = (not((self).f29) if (self).f29 else d_989_cf45_)
                nw163_[int(13)] = (len((d_988_cf46_).set(default__.safeIndex((d_990_v134_).cardinality, len(d_988_cf46_)), d_840_v35_))) >= (p1)
                nw163_[int(14)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vu"))) != (d_845_v38_)
                nw163_[int(15)] = d_985_cf49_
                nw163_[int(16)] = False
                d_991_v135_ = nw163_
                index137_ = default__.safeIndex(338, (d_991_v135_).length(0))
                (d_991_v135_)[index137_] = d_985_cf49_
                index138_ = default__.safeIndex(338, (d_991_v135_).length(0))
                (d_991_v135_)[index138_] = ((d_987_cf47_) >= (d_987_cf47_) if d_989_cf45_ else d_985_cf49_)
                d_992_v136_: _dafny.Array
                nw164_ = _dafny.Array(_dafny.CodePoint('D'), 5)
                d_992_v136_ = nw164_
                index139_ = default__.safeIndex(825, (d_992_v136_).length(0))
                (d_992_v136_)[index139_] = d_840_v35_
                d_993_v137_: _dafny.Array
                nw165_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_993_v137_ = nw165_
                d_994_v138_: C1
                nw166_ = C1()
                nw166_.ctor__(len(default__.fm28(p1, (self).f29, d_985_cf49_, globalState)), d_993_v137_, d_989_cf45_)
                d_994_v138_ = nw166_
                d_995_v139_: C1
                d_995_v139_ = d_994_v138_
                d_996_v140_: _dafny.Array
                nw167_ = _dafny.Array(None, 13)
                nw167_[int(0)] = d_994_v138_
                nw167_[int(1)] = d_994_v138_
                nw167_[int(2)] = d_994_v138_
                nw167_[int(3)] = (d_995_v139_)
                nw167_[int(4)] = d_994_v138_
                nw167_[int(5)] = d_994_v138_
                nw167_[int(6)] = d_994_v138_
                nw167_[int(7)] = d_994_v138_
                nw167_[int(8)] = d_994_v138_
                nw167_[int(9)] = d_994_v138_
                nw167_[int(10)] = d_994_v138_
                nw167_[int(11)] = (d_995_v139_)
                nw167_[int(12)] = d_994_v138_
                d_996_v140_ = nw167_
                index140_ = default__.safeIndex(296, (d_996_v140_).length(0))
                (d_996_v140_)[index140_] = d_994_v138_
                d_997_v141_: _dafny.Array
                nw168_ = _dafny.Array(None, 13)
                nw168_[int(0)] = (d_910_v86_).set(d_989_cf45_, d_985_cf49_)
                nw168_[int(1)] = d_910_v86_
                nw168_[int(2)] = _dafny.Map({d_985_cf49_: d_989_cf45_})
                nw168_[int(3)] = d_910_v86_
                nw168_[int(4)] = d_910_v86_
                nw168_[int(5)] = d_910_v86_
                nw168_[int(6)] = d_910_v86_
                nw168_[int(7)] = d_910_v86_
                nw168_[int(8)] = d_910_v86_
                nw168_[int(9)] = d_910_v86_
                nw168_[int(10)] = d_910_v86_
                nw168_[int(11)] = d_910_v86_
                nw168_[int(12)] = d_910_v86_
                d_997_v141_ = nw168_
                d_998_v142_: _dafny.Map
                d_998_v142_ = _dafny.Map({(0) - (d_987_cf47_): _dafny.Map({p1: d_985_cf49_})})
                d_999_v143_: C3
                nw169_ = C3()
                nw169_.ctor__(d_997_v141_, d_998_v142_, d_989_cf45_)
                d_999_v143_ = nw169_
                d_1000_v144_: _dafny.Set
                d_1000_v144_ = _dafny.Set({d_999_v143_, d_999_v143_})
                d_1001_v145_: _dafny.Map
                d_1001_v145_ = _dafny.Map({d_910_v86_: d_1000_v144_})
                d_1002_v146_: _dafny.Array
                nw170_ = _dafny.Array(None, 4)
                nw170_[int(0)] = d_1000_v144_
                nw170_[int(1)] = d_1000_v144_
                nw170_[int(2)] = ((d_1001_v145_)[d_910_v86_] if (d_910_v86_) in (d_1001_v145_) else _dafny.Set({d_999_v143_, d_999_v143_, d_999_v143_, d_999_v143_, d_999_v143_}))
                nw170_[int(3)] = _dafny.Set({d_999_v143_})
                d_1002_v146_ = nw170_
                d_1003_v147_: _dafny.Map
                d_1003_v147_ = _dafny.Map({d_1002_v146_: (self).f29})
                d_1004_v148_: _dafny.Map
                d_1004_v148_ = _dafny.Map({default__.fm29(d_987_cf47_, d_985_cf49_, globalState): (p0)[default__.safeIndex(d_987_cf47_, len(p0))]})
                index141_ = default__.safeIndex(338, (d_991_v135_).length(0))
                index142_ = default__.safeIndex(825, (d_992_v136_).length(0))
                index143_ = default__.safeIndex(296, (d_996_v140_).length(0))
                rhs148_ = ((d_1003_v147_)[d_1002_v146_] if (d_1002_v146_) in (d_1003_v147_) else (len(p0)) != (p1))
                rhs149_ = ((d_1004_v148_)[p2] if (p2) in (d_1004_v148_) else d_987_cf47_)
                rhs150_ = d_988_cf46_
                rhs151_ = d_840_v35_
                rhs152_ = (d_994_v138_ if d_985_cf49_ else (d_994_v138_ if (self).f29 else d_994_v138_))
                lhs90_ = d_991_v135_
                lhs91_ = default__.safeIndex(338, (d_991_v135_).length(0))
                lhs92_ = d_992_v136_
                lhs93_ = default__.safeIndex(825, (d_992_v136_).length(0))
                lhs94_ = d_996_v140_
                lhs95_ = default__.safeIndex(296, (d_996_v140_).length(0))
                lhs90_[lhs91_] = rhs148_
                d_987_cf47_ = rhs149_
                d_845_v38_ = rhs150_
                lhs92_[lhs93_] = rhs151_
                lhs94_[lhs95_] = rhs152_
                d_1005_v149_: D3
                d_1005_v149_ = D3_DC10(d_845_v38_, d_987_cf47_, d_985_cf49_)
                d_1006_v150_: _dafny.Array
                def lambda39_(d_1007_cf47_):
                    def lambda40_(d_1008_i9_):
                        return (d_1008_i9_) + (d_1007_cf47_)

                    return lambda40_

                init21_ = lambda39_(d_987_cf47_)
                nw171_ = _dafny.Array(None, 25)
                for i0_21_ in range(nw171_.length(0)):
                    nw171_[i0_21_] = init21_(i0_21_)
                d_1006_v150_ = nw171_
                rhs153_ = d_1005_v149_
                rhs154_ = d_1006_v150_
                rhs155_ = p1
                rhs156_ = (self).f30
                lhs96_ = globalState
                d_1005_v149_ = rhs153_
                lhs96_.f24 = rhs154_
                d_987_cf47_ = rhs155_
                d_989_cf45_ = rhs156_
                d_1009_v151_: C2
                nw172_ = C2()
                nw172_.ctor__(_dafny.SeqWithoutIsStrInference([True]), d_985_cf49_)
                d_1009_v151_ = nw172_
                d_1010_v152_: _dafny.Map
                d_1010_v152_ = _dafny.Map({d_987_cf47_: d_1009_v151_})
                rhs157_ = (self).f29
                rhs158_ = ((d_1010_v152_)[(d_999_v143_).fm9(d_987_cf47_, p0, p1, (self).f30, globalState)] if ((d_999_v143_).fm9(d_987_cf47_, p0, p1, (self).f30, globalState)) in (d_1010_v152_) else d_1009_v151_)
                d_989_cf45_ = rhs157_
                d_1009_v151_ = rhs158_
                d_1011_v153_: _dafny.Set
                d_1011_v153_ = _dafny.Set({d_987_cf47_})
                d_1012_v154_: _dafny.Map
                d_1012_v154_ = _dafny.Map({True: _dafny.Set({len(p0)})})
                d_1013_v155_: _dafny.Map
                d_1013_v155_ = _dafny.Map({len(d_1011_v153_): ((d_1012_v154_)[d_989_cf45_] if (d_989_cf45_) in (d_1012_v154_) else d_1011_v153_)})
                d_1014_v156_: _dafny.Map
                d_1014_v156_ = _dafny.Map({d_845_v38_: (d_990_v134_).cardinality})
                d_1015_v157_: _dafny.Map
                d_1015_v157_ = _dafny.Map({default__.safeModuloInt(607, len(d_1014_v156_)): d_1013_v155_})
                index144_ = default__.safeIndex(338, (d_991_v135_).length(0))
                rhs159_ = not(((d_985_cf49_) or ((d_991_v135_)[default__.safeIndex(338, (d_991_v135_).length(0))])) not in ((d_1009_v151_).f38))
                rhs160_ = ((d_1015_v157_)[len((d_988_cf46_) + (d_988_cf46_))] if (len((d_988_cf46_) + (d_988_cf46_))) in (d_1015_v157_) else d_1013_v155_)
                rhs161_ = (d_990_v134_).set(not(d_989_cf45_), default__.abs(default__.safeModuloInt((0) - ((0) - (p1)), p1)))
                rhs162_ = d_991_v135_
                lhs97_ = d_991_v135_
                lhs98_ = default__.safeIndex(338, (d_991_v135_).length(0))
                lhs99_ = globalState
                lhs97_[lhs98_] = rhs159_
                d_1013_v155_ = rhs160_
                d_990_v134_ = rhs161_
                lhs99_.f20 = rhs162_
            elif True:
                d_1016_v158_: _dafny.Array
                def lambda41_(d_1017_cf45_):
                    def lambda42_(d_1018_i10_):
                        return d_1017_cf45_

                    return lambda42_

                init22_ = lambda41_(d_989_cf45_)
                nw173_ = _dafny.Array(None, 7)
                for i0_22_ in range(nw173_.length(0)):
                    nw173_[i0_22_] = init22_(i0_22_)
                d_1016_v158_ = nw173_
                (globalState).f20 = d_1016_v158_
                default__.m0(default__.fm1(d_840_v35_, globalState), globalState)
                d_1019_v159_: _dafny.Seq
                d_1019_v159_ = _dafny.SeqWithoutIsStrInference([(self).f29, (self).f29, default__.fm26((self).f29, (self).f29, (self).f30, p1, globalState)])
                d_1020_v160_: _dafny.Map
                d_1020_v160_ = _dafny.Map({d_1019_v159_: p1})
                d_1020_v160_ = d_1020_v160_
                d_1021_v161_: _dafny.Seq
                d_1021_v161_ = _dafny.SeqWithoutIsStrInference([d_1016_v158_, d_1016_v158_, d_1016_v158_])
                index145_ = default__.safeIndex(132, (d_1016_v158_).length(0))
                (d_1016_v158_)[index145_] = (self).f30
                d_1022_v162_: _dafny.Array
                nw174_ = _dafny.Array(_dafny.Set({}), 23)
                d_1022_v162_ = nw174_
                d_1023_v163_: _dafny.Set
                d_1023_v163_ = _dafny.Set({d_840_v35_})
                index146_ = default__.safeIndex(386, (d_1022_v162_).length(0))
                (d_1022_v162_)[index146_] = (d_1023_v163_) | (d_1023_v163_)
                d_1024_v164_: _dafny.Seq
                d_1024_v164_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.CodePoint('k'), d_840_v35_, d_840_v35_})])
                index147_ = default__.safeIndex(132, (d_1016_v158_).length(0))
                index148_ = default__.safeIndex(386, (d_1022_v162_).length(0))
                rhs163_ = len(((d_1019_v159_) + ((d_1019_v159_ if (self).f29 else d_1019_v159_))).set(default__.safeIndex(p1, len((d_1019_v159_) + ((d_1019_v159_ if (self).f29 else d_1019_v159_)))), not ((self).f30) or (d_985_cf49_)))
                rhs164_ = (-756) + (p1)
                rhs165_ = d_1021_v161_
                rhs166_ = (self).f29
                rhs167_ = (d_1024_v164_)[default__.safeIndex(len(p0), len(d_1024_v164_))]
                lhs100_ = d_1016_v158_
                lhs101_ = default__.safeIndex(132, (d_1016_v158_).length(0))
                lhs102_ = d_1022_v162_
                lhs103_ = default__.safeIndex(386, (d_1022_v162_).length(0))
                d_987_cf47_ = rhs163_
                d_987_cf47_ = rhs164_
                d_1021_v161_ = rhs165_
                lhs100_[lhs101_] = rhs166_
                lhs102_[lhs103_] = rhs167_
                index149_ = default__.safeIndex(132, (d_1016_v158_).length(0))
                rhs168_ = (True if ((d_910_v86_)[(self).f29] if ((self).f29) in (d_910_v86_) else (d_1016_v158_)[default__.safeIndex(132, (d_1016_v158_).length(0))]) else (self).f29)
                rhs169_ = (d_987_cf47_) + ((0) - ((d_987_cf47_) - (d_987_cf47_)))
                lhs104_ = d_1016_v158_
                lhs105_ = default__.safeIndex(132, (d_1016_v158_).length(0))
                lhs104_[lhs105_] = rhs168_
                d_987_cf47_ = rhs169_
            d_1025_v165_: _dafny.Array
            nw175_ = _dafny.Array(False, 18)
            d_1025_v165_ = nw175_
            index150_ = default__.safeIndex(590, (d_1025_v165_).length(0))
            (d_1025_v165_)[index150_] = d_985_cf49_
            index151_ = default__.safeIndex(590, (d_1025_v165_).length(0))
            (d_1025_v165_)[index151_] = (self).f29
            d_1026_v166_: _dafny.Seq
            d_1026_v166_ = _dafny.SeqWithoutIsStrInference([(self).f30])
            index152_ = default__.safeIndex(590, (d_1025_v165_).length(0))
            index153_ = default__.safeIndex(590, (d_1025_v165_).length(0))
            rhs170_ = ((d_1026_v166_)[default__.safeIndex(p1, len(d_1026_v166_))]) not in (d_841_v36_)
            rhs171_ = (d_1025_v165_)[default__.safeIndex(590, (d_1025_v165_).length(0))]
            rhs172_ = d_1026_v166_
            rhs173_ = default__.fm3(globalState)
            lhs106_ = d_1025_v165_
            lhs107_ = default__.safeIndex(590, (d_1025_v165_).length(0))
            lhs108_ = d_1025_v165_
            lhs109_ = default__.safeIndex(590, (d_1025_v165_).length(0))
            lhs110_ = globalState
            lhs106_[lhs107_] = rhs170_
            lhs108_[lhs109_] = rhs171_
            d_1026_v166_ = rhs172_
            lhs110_.f16 = rhs173_
            d_1027_v167_: _dafny.Map
            d_1027_v167_ = _dafny.Map({d_985_cf49_: d_840_v35_})
            d_1028_v168_: _dafny.Map
            d_1028_v168_ = _dafny.Map({len(p0): d_840_v35_})
            d_1029_v169_: _dafny.Array
            nw176_ = _dafny.Array(None, 7)
            nw176_[int(0)] = ((d_1027_v167_)[d_989_cf45_] if (d_989_cf45_) in (d_1027_v167_) else d_840_v35_)
            nw176_[int(1)] = d_840_v35_
            nw176_[int(2)] = _dafny.CodePoint('k')
            nw176_[int(3)] = d_840_v35_
            nw176_[int(4)] = d_840_v35_
            nw176_[int(5)] = ((d_1028_v168_)[d_987_cf47_] if (d_987_cf47_) in (d_1028_v168_) else d_840_v35_)
            nw176_[int(6)] = d_840_v35_
            d_1029_v169_ = nw176_
            index154_ = default__.safeIndex(870, (d_1029_v169_).length(0))
            (d_1029_v169_)[index154_] = d_840_v35_
            index155_ = default__.safeIndex(870, (d_1029_v169_).length(0))
            (d_1029_v169_)[index155_] = d_840_v35_
        elif True:
            d_1030___mcc_h16_ = source15_.cf39
            d_1031_cf39_ = d_1030___mcc_h16_
            d_1032_v170_: int
            d_1032_v170_ = 18
            d_1033_v171_: _dafny.Set
            d_1033_v171_ = _dafny.Set({p1})
            d_1032_v170_ = len(d_1033_v171_)
            d_845_v38_ = d_845_v38_
            d_1034_v172_: _dafny.Array
            nw177_ = _dafny.Array(False, 6)
            d_1034_v172_ = nw177_
            index156_ = default__.safeIndex(445, (d_1034_v172_).length(0))
            (d_1034_v172_)[index156_] = ((self).f29) not in (d_841_v36_)
            index157_ = default__.safeIndex(445, (d_1034_v172_).length(0))
            (d_1034_v172_)[index157_] = (self).f30
            (globalState).f16 = (default__.safeDivisionInt(-612, d_1032_v170_)) >= (d_1032_v170_)
        r0 = (self).f29
        r1 = not((default__.safeDivisionInt(p1, p1)) != (p1))
        return r0, r1

    @property
    def f29(self):
        return self._f29
    @property
    def f30(self):
        return self._f30
