// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm0(p0, globalState) {
      return (((false) ? (_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true, false))) : (_dafny.MultiSet.fromElements(true)))).IsProperSubsetOf(((false) ? (_dafny.MultiSet.fromElements(true, false, !(true))) : (_dafny.MultiSet.fromElements(!(!(true)), false))));
    };
    static fm1(p0, globalState) {
      return _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber(-606)), _dafny.ZERO);
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true, true), _dafny.Seq.of(false));
    };
    static fm3(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("eh"), _dafny.Seq.UnicodeFromString("rm")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qlxxw"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(992)), function (_0_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })));
    };
    static fm4(globalState) {
      return _dafny.Seq.of(_dafny.Seq.UnicodeFromString("sh"), ((true) ? (_dafny.Seq.UnicodeFromString("siqr")) : (_dafny.Seq.UnicodeFromString("ovpwyj"))), _dafny.Seq.UnicodeFromString("amlqk"), _dafny.Seq.UnicodeFromString("reultjxs"));
    };
    static fm5(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(119),true)).length), new BigNumber(921), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(836)))).Elements) {
          let _1_v0 = _compr_0;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(836)))).contains(_1_v0)) {
            _coll0.add((_1_v0).multipliedBy(new BigNumber(248)));
          }
        }
        return _coll0;
      }()).length),false)).length))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(839))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(42)), function (_2_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })).length), new BigNumber((_dafny.Set.fromElements(false, false, !(false), false)).length))));
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("iffq"),_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(609)), function (_3_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("rxl")));
    };
    static fm7(globalState) {
      let _source0 = _module.D5.create_DC10(_dafny.MultiSet.fromElements(false));
      if (_source0.is_DC11) {
        let _4___mcc_h0 = (_source0).cf17;
        let _5_cf17 = _4___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("csjiecvhl"),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),new BigNumber((function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of _dafny.IntegerRange(new BigNumber(-39), new BigNumber(780))) {
            let _6_v0 = _compr_1;
            if (((new BigNumber(-39)).isLessThanOrEqualTo(_6_v0)) && ((_6_v0).isLessThan(new BigNumber(780)))) {
              _coll1.push([_module.__default.safeModuloInt(_6_v0, new BigNumber((_dafny.Seq.UnicodeFromString("nshammh")).length)),true]);
            }
          }
          return _coll1;
        }()).length))).length))).length))).length))).length))).length),_module.D4.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(564))));
      } else if (_source0.is_DC12) {
        let _7___mcc_h1 = (_source0).cf18;
        let _8___mcc_h2 = (_source0).cf19;
        let _9___mcc_h3 = (_source0).cf20;
        let _10___mcc_h4 = (_source0).cf21;
        let _11___mcc_h5 = (_source0).cf22;
        let _12_cf22 = _11___mcc_h5;
        let _13_cf21 = _10___mcc_h4;
        let _14_cf20 = _9___mcc_h3;
        let _15_cf19 = _8___mcc_h2;
        let _16_cf18 = _7___mcc_h1;
        return _dafny.Map.Empty.slice().updateUnsafe(_12_cf22,_module.D4.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(_14_cf20,_16_cf18)));
      } else {
        let _17___mcc_h6 = (_source0).cf16;
        let _18_cf16 = _17___mcc_h6;
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-694),_module.D4.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(false)).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality()),_module.D4.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-749)))));
      }
    };
    static fm8(p0, p1, globalState) {
      return _dafny.Seq.Concat(((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(755)), function (_19_i0) {
        return (_dafny.ZERO).minus(new BigNumber(-863));
      })) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-749)), function (_20_i1) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("uhedjju")).length);
      }))), ((!(true)) ? (_dafny.Seq.of(new BigNumber(227))) : (_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false, false, false, !(false))).length), new BigNumber(215)))));
    };
    static fm9(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(221)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(6)), function (_21_i0) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length),_dafny.MultiSet.fromElements(true))).length);
      }))).length),!((new BigNumber((_dafny.Seq.UnicodeFromString("sjycrdgy")).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("dwneg")).length))));
    };
    static fm10(p0, globalState) {
      return _dafny.Seq.of(_dafny.Seq.of(true), _dafny.Seq.Concat(_dafny.Seq.of(!(true), !(true)), _dafny.Seq.of(!(false))));
    };
    static fm11(p0, p1, p2, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(!(false))));
    };
    static fm12(globalState) {
      if (false) {
        if (true) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('t'.codePointAt(0));
        }
      } else if (false) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }
    };
    static fm13(p0, p1, p2, globalState) {
      return (_module.D7.create_DC14(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("qsiewudbc"),_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(926), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-434)), function (_22_i0) {
  return new _dafny.CodePoint('t'.codePointAt(0));
})).length), new BigNumber(-827), new BigNumber(851), new BigNumber(812))).length), new BigNumber(380))))).dtor_cf24;
    };
    static m0(p0, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _23_v0;
      _23_v0 = true;
      let _24_v1;
      _24_v1 = _dafny.Seq.UnicodeFromString("wtfyrctyq");
      let _25_v2;
      _25_v2 = _dafny.Seq.UnicodeFromString("plyeo");
      let _source1 = ((_23_v0) ? (_24_v1) : (_25_v2));
      let _26___mcc_h0 = _source1;
      let _27_cf4 = _26___mcc_h0;
      let _28_v3;
      _28_v3 = new _dafny.CodePoint('m'.codePointAt(0));
      let _29_v4;
      _29_v4 = _dafny.MultiSet.fromElements(_28_v3, _28_v3, _28_v3);
      let _30_v5;
      _30_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_29_v4).cardinality()),_23_v0);
      let _31_v6;
      _31_v6 = new BigNumber(847);
      _30_v5 = (_30_v5).update(_31_v6, (new BigNumber(-326)).isLessThanOrEqualTo(_31_v6));
      if (_23_v0) {
        let _32_v7;
        _32_v7 = _dafny.Seq.of(_27_cf4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(180)), ((_33_v3) => function (_34_i0) {
          return _33_v3;
        })(_28_v3)));
        let _35_v8;
        let _nw0 = Array((new BigNumber(6)).toNumber());
        _nw0[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_32_v7, _32_v7);
        _nw0[(_dafny.ONE).toNumber()] = _32_v7;
        _nw0[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("tsajg"));
        _nw0[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_32_v7, _32_v7);
        _nw0[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_32_v7, _dafny.Seq.of(_24_v1, _27_cf4));
        _nw0[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_32_v7, _32_v7);
        _35_v8 = _nw0;
        let _index0 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_35_v8).length));
        (_35_v8)[_index0] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-648)), ((_36_cf4) => function (_37_i1) {
          return _36_cf4;
        })(_27_cf4));
        let _index1 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_35_v8).length));
        (_35_v8)[_index1] = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(822)), ((_38_v3) => function (_39_i2) {
          return _38_v3;
        })(_28_v3)));
        let _40_v9;
        let _init0 = ((_41_v6) => function (_42_i3) {
          return _dafny.Seq.Concat(_dafny.Seq.of(_41_v6, _41_v6), _dafny.Seq.of(_41_v6));
        })(_31_v6);
        let _nw1 = Array((new BigNumber(18)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
          _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _40_v9 = _nw1;
        let _index2 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_40_v9).length));
        (_40_v9)[_index2] = _module.__default.fm8(true, _29_v4, globalState);
        let _43_v10;
        _43_v10 = _dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(_31_v6)));
        let _44_v11;
        _44_v11 = _dafny.Seq.of(_dafny.Seq.Concat(_43_v10, _module.__default.fm8(_23_v0, _29_v4, globalState)));
        let _index3 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_40_v9).length));
        (_40_v9)[_index3] = (_44_v11)[_module.__default.safeIndex(_module.__default.safeModuloInt(_31_v6, _31_v6), new BigNumber((_44_v11).length))];
        let _45_v12;
        _45_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(842),_30_v5);
        let _46_v13;
        _46_v13 = _dafny.Seq.of((((_45_v12).contains((_dafny.ZERO).minus(_31_v6))) ? ((_45_v12).get((_dafny.ZERO).minus(_31_v6))) : (_dafny.Map.Empty.slice().updateUnsafe(_31_v6,_23_v0))), _30_v5);
        _46_v13 = _dafny.Seq.Concat(_dafny.Seq.of((_30_v5).update(_31_v6, _23_v0), _module.__default.fm9(_23_v0, _23_v0, _23_v0, _23_v0, globalState), _30_v5, _30_v5), _46_v13);
        let _47_v14;
        let _init1 = ((_48_v0) => function (_49_i4) {
          return _48_v0;
        })(_23_v0);
        let _nw2 = Array((new BigNumber(21)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw2.length); _i0_1++) {
          _nw2[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _47_v14 = _nw2;
        let _50_v15;
        _50_v15 = _dafny.Map.Empty.slice().updateUnsafe(_23_v0,_47_v14);
        let _51_v16;
        _51_v16 = _dafny.Map.Empty.slice().updateUnsafe((_31_v6).minus(_31_v6),_31_v6);
        let _52_v17;
        let _init2 = ((_53_v6) => function (_54_i5) {
          return (_54_i5).multipliedBy(_53_v6);
        })(_31_v6);
        let _nw3 = Array((new BigNumber(9)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw3.length); _i0_2++) {
          _nw3[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _52_v17 = _nw3;
        let _rhs0 = (_50_v15).Merge((_50_v15).Merge(_50_v15));
        let _rhs1 = (_31_v6).plus(_31_v6);
        let _rhs2 = (_dafny.Map.Empty.slice().updateUnsafe(_31_v6,_31_v6)).Merge(_51_v16);
        let _rhs3 = !(_31_v6).isEqualTo(_31_v6);
        let _rhs4 = _52_v17;
        let _lhs0 = globalState;
        _50_v15 = _rhs0;
        _31_v6 = _rhs1;
        _51_v16 = _rhs2;
        _lhs0.f1 = _rhs3;
        _52_v17 = _rhs4;
        let _55_v18;
        _55_v18 = _dafny.Seq.of(_23_v0);
        let _56_v19;
        _56_v19 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_55_v18).length)));
        let _57_v20;
        _57_v20 = _dafny.Map.Empty.slice().updateUnsafe(_23_v0,_43_v10);
        let _index4 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_47_v14).length));
        (_47_v14)[_index4] = _module.__default.fm0((((_56_v19).contains(new BigNumber((function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(_28_v3,new BigNumber((_57_v20).length))).Keys.Elements) {
            let _59_v21 = _compr_3;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_28_v3,new BigNumber((_57_v20).length))).contains(_59_v21)) {
              _coll3.add(_59_v21);
            }
          }
          return _coll3;
        }()).length))) ? ((_56_v19).get(new BigNumber((function () {
          let _coll2 = new _dafny.Set();
          for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(_28_v3,new BigNumber((_57_v20).length))).Keys.Elements) {
            let _58_v21 = _compr_2;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_28_v3,new BigNumber((_57_v20).length))).contains(_58_v21)) {
              _coll2.add(_58_v21);
            }
          }
          return _coll2;
        }()).length))) : (new BigNumber(589))), globalState);
        let _60_v22;
        _60_v22 = _dafny.Map.Empty.slice().updateUnsafe(_28_v3,_52_v17);
        let _index5 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_47_v14).length));
        (_47_v14)[_index5] = !(((_23_v0) ? (_60_v22) : (_dafny.Map.Empty.slice().updateUnsafe(_28_v3,_52_v17)))).equals((_60_v22).Merge(_60_v22));
      } else {
        let _61_v23;
        _61_v23 = _dafny.Map.Empty.slice().updateUnsafe(_23_v0,_27_cf4);
        _61_v23 = ((_dafny.Map.Empty.slice().updateUnsafe(_23_v0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(269)), ((_62_v3) => function (_63_i6) {
          return _62_v3;
        })(_28_v3)))).Merge(_61_v23)).Merge(_61_v23);
        let _64_v24;
        let _nw4 = new _module.C0();
        _nw4.__ctor();
        _64_v24 = _nw4;
        let _65_v25;
        let _nw5 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
        _65_v25 = _nw5;
        let _66_v26;
        _66_v26 = _dafny.Set.fromElements(_31_v6);
        let _67_v27;
        _67_v27 = _dafny.MultiSet.fromElements(_module.__default.fm11(_23_v0, _23_v0, new BigNumber((_66_v26).length), globalState));
        let _index6 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_65_v25).length));
        (_65_v25)[_index6] = _module.__default.fm10(new BigNumber((_67_v27).cardinality()), globalState);
        let _68_v28;
        let _nw6 = Array((new BigNumber(24)).toNumber());
        _nw6[(_dafny.ZERO).toNumber()] = _27_cf4;
        _nw6[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_24_v1, _dafny.Seq.UnicodeFromString("knesulqj"));
        _nw6[(new BigNumber(2)).toNumber()] = _27_cf4;
        _nw6[(new BigNumber(3)).toNumber()] = _27_cf4;
        _nw6[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("yxthifwc");
        _nw6[(new BigNumber(5)).toNumber()] = (((_61_v23).contains(true)) ? ((_61_v23).get(true)) : (_24_v1));
        _nw6[(new BigNumber(6)).toNumber()] = _27_cf4;
        _nw6[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_24_v1, _27_cf4);
        _nw6[(new BigNumber(8)).toNumber()] = _27_cf4;
        _nw6[(new BigNumber(9)).toNumber()] = _24_v1;
        _nw6[(new BigNumber(10)).toNumber()] = _27_cf4;
        _nw6[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_24_v1, _dafny.Seq.update(_24_v1, _module.__default.safeIndex(new BigNumber((_24_v1).length), new BigNumber((_24_v1).length)), _28_v3));
        _nw6[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("dqmctbicw");
        _nw6[(new BigNumber(13)).toNumber()] = _24_v1;
        _nw6[(new BigNumber(14)).toNumber()] = _24_v1;
        _nw6[(new BigNumber(15)).toNumber()] = _27_cf4;
        _nw6[(new BigNumber(16)).toNumber()] = _27_cf4;
        _nw6[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_24_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-182)), ((_69_v3) => function (_70_i7) {
          return _69_v3;
        })(_28_v3)));
        _nw6[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_27_cf4, _dafny.Seq.UnicodeFromString("gkykgdr"));
        _nw6[(new BigNumber(19)).toNumber()] = _24_v1;
        _nw6[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_24_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(18)), ((_71_v1, _72_v6) => function (_73_i8) {
          return (_71_v1)[_module.__default.safeIndex(_72_v6, new BigNumber((_71_v1).length))];
        })(_24_v1, _31_v6)));
        _nw6[(new BigNumber(21)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_27_cf4, _dafny.Seq.UnicodeFromString("hyohbsaxk")), _module.__default.safeIndex(_31_v6, new BigNumber((_dafny.Seq.Concat(_27_cf4, _dafny.Seq.UnicodeFromString("hyohbsaxk"))).length)), _28_v3);
        _nw6[(new BigNumber(22)).toNumber()] = _24_v1;
        _nw6[(new BigNumber(23)).toNumber()] = _dafny.Seq.UnicodeFromString("krh");
        _68_v28 = _nw6;
        let _index7 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_68_v28).length));
        (_68_v28)[_index7] = _dafny.Seq.Concat(_27_cf4, _27_cf4);
        let _74_v29;
        _74_v29 = _dafny.Set.fromElements(!(true), _23_v0, _23_v0, _23_v0);
        let _75_v30;
        _75_v30 = _dafny.Seq.of(_23_v0);
        let _76_v31;
        _76_v31 = _dafny.Seq.of(_75_v30);
        let _index8 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_65_v25).length));
        let _index9 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_68_v28).length));
        let _rhs5 = _module.__default.fm0(new BigNumber((_74_v29).length), globalState);
        let _rhs6 = _76_v31;
        let _rhs7 = _dafny.Seq.Concat(_24_v1, _dafny.Seq.update(_27_cf4, _module.__default.safeIndex(_31_v6, new BigNumber((_27_cf4).length)), (_24_v1)[_module.__default.safeIndex(_31_v6, new BigNumber((_24_v1).length))]));
        let _rhs8 = (_dafny.ZERO).minus(_31_v6);
        let _rhs9 = _64_v24;
        let _lhs1 = _65_v25;
        let _lhs2 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_65_v25).length));
        let _lhs3 = _68_v28;
        let _lhs4 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_68_v28).length));
        _23_v0 = _rhs5;
        _lhs1[_lhs2] = _rhs6;
        _lhs3[_lhs4] = _rhs7;
        r0 = _rhs8;
        _64_v24 = _rhs9;
        _25_v2 = _25_v2;
        _24_v1 = _27_cf4;
      }
      r1 = _module.__default.safeDivisionInt(_31_v6, new BigNumber((_27_cf4).length));
      _31_v6 = _31_v6;
      let _77_v32;
      _77_v32 = _dafny.Seq.of(_23_v0, true, _23_v0);
      let _78_v33;
      _78_v33 = new BigNumber(514);
      let _79_i9;
      _79_i9 = _dafny.ZERO;
      L0: {
        while ((_77_v32)[_module.__default.safeIndex(_78_v33, new BigNumber((_77_v32).length))]) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_79_i9)) {
              break L0;
            }
            _79_i9 = (_79_i9).plus(_dafny.ONE);
            let _80_v34;
            _80_v34 = new _dafny.CodePoint('b'.codePointAt(0));
            let _81_v35;
            _81_v35 = _dafny.Map.Empty.slice().updateUnsafe((_23_v0) === (_23_v0),_25_v2);
            let _82_v36;
            let _nw7 = new _module.C0();
            _nw7.__ctor();
            _82_v36 = _nw7;
            let _83_v37;
            _83_v37 = _dafny.Seq.of(_82_v36, _82_v36);
            let _84_v38;
            _84_v38 = _dafny.Map.Empty.slice().updateUnsafe((_83_v37)[_module.__default.safeIndex(_78_v33, new BigNumber((_83_v37).length))],_78_v33);
            let _85_v39;
            _85_v39 = _dafny.MultiSet.fromElements(_78_v33, _78_v33, _78_v33);
            let _rhs10 = (((_84_v38).contains(_82_v36)) ? ((_84_v38).get(_82_v36)) : ((((_85_v39).contains(new BigNumber(925))) ? ((_85_v39).get(new BigNumber(925))) : (_78_v33))));
            let _rhs11 = _80_v34;
            let _rhs12 = _81_v35;
            r1 = _rhs10;
            _80_v34 = _rhs11;
            _81_v35 = _rhs12;
            let _86_v40;
            _86_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(867)), ((_87_v34) => function (_88_i10) {
              return _87_v34;
            })(_80_v34))).length),_module.__default.fm0(_78_v33, globalState));
            let _89_v41;
            let _nw8 = Array((new BigNumber(11)).toNumber());
            _nw8[(_dafny.ZERO).toNumber()] = _23_v0;
            _nw8[(_dafny.ONE).toNumber()] = _23_v0;
            _nw8[(new BigNumber(2)).toNumber()] = _23_v0;
            _nw8[(new BigNumber(3)).toNumber()] = _23_v0;
            _nw8[(new BigNumber(4)).toNumber()] = _23_v0;
            _nw8[(new BigNumber(5)).toNumber()] = (((_86_v40).contains(_78_v33)) ? ((_86_v40).get(_78_v33)) : (_23_v0));
            _nw8[(new BigNumber(6)).toNumber()] = !(_23_v0);
            _nw8[(new BigNumber(7)).toNumber()] = _23_v0;
            _nw8[(new BigNumber(8)).toNumber()] = _23_v0;
            _nw8[(new BigNumber(9)).toNumber()] = _23_v0;
            _nw8[(new BigNumber(10)).toNumber()] = _23_v0;
            _89_v41 = _nw8;
            let _index10 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_89_v41).length));
            (_89_v41)[_index10] = ((_23_v0) ? (false) : (_23_v0));
            let _90_v42;
            _90_v42 = _module.D5.create_DC12(_78_v33, _module.__default.fm12(globalState), _23_v0, _24_v1, _78_v33);
            let _index11 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_89_v41).length));
            (_89_v41)[_index11] = !((_23_v0) || ((_90_v42).dtor_cf20));
            let _rhs13 = _78_v33;
            let _rhs14 = (_89_v41)[_module.__default.safeIndex(new BigNumber(760), new BigNumber((_89_v41).length))];
            let _rhs15 = (_89_v41)[_module.__default.safeIndex(new BigNumber(760), new BigNumber((_89_v41).length))];
            let _rhs16 = (_78_v33).plus(_78_v33);
            let _lhs5 = globalState;
            r0 = _rhs13;
            _23_v0 = _rhs14;
            _lhs5.f1 = _rhs15;
            r1 = _rhs16;
            let _91_v43;
            _91_v43 = _dafny.MultiSet.fromElements(_23_v0);
            _91_v43 = _91_v43;
          }
        }
      }
      let _92_v44;
      _92_v44 = _dafny.MultiSet.fromElements(_23_v0);
      let _93_v45;
      _93_v45 = _dafny.Set.fromElements((((_92_v44).contains(_23_v0)) ? ((_92_v44).get(_23_v0)) : (_78_v33)), new BigNumber((_77_v32).length));
      let _94_v46;
      _94_v46 = _dafny.Seq.of(_93_v45);
      let _95_v47;
      _95_v47 = _dafny.Map.Empty.slice().updateUnsafe(_24_v1,_93_v45);
      let _96_v49;
      _96_v49 = _dafny.Seq.of(_24_v1, _24_v1, _24_v1);
      let _97_v50;
      _97_v50 = new _dafny.CodePoint('l'.codePointAt(0));
      let _98_v51;
      _98_v51 = _module.D7.create_DC14(_95_v47);
      let _99_v54;
      _99_v54 = _dafny.MultiSet.fromElements(_24_v1);
      let _100_v55;
      let _nw9 = Array((new BigNumber(18)).toNumber());
      _nw9[(_dafny.ZERO).toNumber()] = _95_v47;
      _nw9[(_dafny.ONE).toNumber()] = _95_v47;
      _nw9[(new BigNumber(2)).toNumber()] = function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_96_v49).Elements) {
          let _101_v48 = _compr_4;
          if (_dafny.Seq.contains(_96_v49, _101_v48)) {
            _coll4.push([_101_v48,_93_v45]);
          }
        }
        return _coll4;
      }();
      _nw9[(new BigNumber(3)).toNumber()] = _module.__default.fm13(_23_v0, false, _24_v1, globalState);
      _nw9[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_24_v1,_93_v45);
      _nw9[(new BigNumber(5)).toNumber()] = _95_v47;
      _nw9[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_24_v1,_93_v45);
      _nw9[(new BigNumber(7)).toNumber()] = (_95_v47).Merge(_95_v47);
      _nw9[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update((_25_v2), _module.__default.safeIndex(_78_v33, new BigNumber(((_25_v2)).length)), _97_v50),_93_v45);
      _nw9[(new BigNumber(9)).toNumber()] = (_95_v47).Merge((_98_v51).dtor_cf24);
      _nw9[(new BigNumber(10)).toNumber()] = ((_23_v0) ? (_95_v47) : (function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_96_v49).Elements) {
          let _102_v52 = _compr_5;
          if (_dafny.Seq.contains(_96_v49, _102_v52)) {
            _coll5.push([_102_v52,_93_v45]);
          }
        }
        return _coll5;
      }()));
      _nw9[(new BigNumber(11)).toNumber()] = _95_v47;
      _nw9[(new BigNumber(12)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("xrfel"),_dafny.Set.fromElements(_78_v33, _78_v33))).update(_24_v1, _93_v45);
      _nw9[(new BigNumber(13)).toNumber()] = _95_v47;
      _nw9[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_24_v1,_93_v45);
      _nw9[(new BigNumber(15)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_24_v1,_93_v45)).Merge(_95_v47);
      _nw9[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_24_v1,_93_v45);
      _nw9[(new BigNumber(17)).toNumber()] = function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_99_v54).Elements) {
          let _103_v53 = _compr_6;
          if ((_99_v54).contains(_103_v53)) {
            _coll6.push([_103_v53,_93_v45]);
          }
        }
        return _coll6;
      }();
      _100_v55 = _nw9;
      let _index12 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_100_v55).length));
      (_100_v55)[_index12] = _95_v47;
      let _104_v56;
      let _nw10 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
      _104_v56 = _nw10;
      let _105_v57;
      let _init3 = ((_106_v33) => function (_107_i11) {
        return _module.D2.create_DC3(_dafny.MultiSet.fromElements(_106_v33));
      })(_78_v33);
      let _nw11 = Array((new BigNumber(23)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw11.length); _i0_3++) {
        _nw11[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _105_v57 = _nw11;
      let _108_v58;
      _108_v58 = _module.D2.create_DC3(_dafny.MultiSet.fromElements(_78_v33, _78_v33));
      let _index13 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_105_v57).length));
      (_105_v57)[_index13] = _108_v58;
      let _109_v59;
      _109_v59 = _module.D4.create_DC9(_104_v56, new BigNumber((_96_v49).length), _104_v56);
      let _110_v60;
      _110_v60 = _dafny.MultiSet.fromElements(_78_v33, _78_v33);
      let _index14 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_100_v55).length));
      let _index15 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_105_v57).length));
      let _rhs17 = _94_v46;
      let _rhs18 = _95_v47;
      let _rhs19 = ((((_23_v0) ? (_module.__default.fm0(_78_v33, globalState)) : (_23_v0))) ? ((_109_v59).dtor_cf13) : (_104_v56));
      let _rhs20 = (_78_v33).isLessThanOrEqualTo((_dafny.ZERO).minus((((_110_v60).contains(new BigNumber(353))) ? ((_110_v60).get(new BigNumber(353))) : (_78_v33))));
      let _rhs21 = _108_v58;
      let _lhs6 = _100_v55;
      let _lhs7 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_100_v55).length));
      let _lhs8 = _105_v57;
      let _lhs9 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_105_v57).length));
      _94_v46 = _rhs17;
      _lhs6[_lhs7] = _rhs18;
      _104_v56 = _rhs19;
      r2 = _rhs20;
      _lhs8[_lhs9] = _rhs21;
      let _111_v61;
      let _nw12 = new _module.C0();
      _nw12.__ctor();
      _111_v61 = _nw12;
      let _hi0 = _78_v33;
      for (let _112_i12 = _78_v33; _112_i12.isLessThan(_hi0); _112_i12 = _112_i12.plus(_dafny.ONE)) {
        let _113_v62;
        let _nw13 = Array((new BigNumber(9)).toNumber()).fill(false);
        _113_v62 = _nw13;
        let _index16 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_113_v62).length));
        (_113_v62)[_index16] = true;
        let _index17 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_113_v62).length));
        (_113_v62)[_index17] = _23_v0;
        let _114_v63;
        let _nw14 = Array((new BigNumber(20)).toNumber()).fill(_dafny.MultiSet.Empty);
        _114_v63 = _nw14;
        _114_v63 = _114_v63;
        let _115_v64;
        let _nw15 = new _module.C0();
        _nw15.__ctor();
        _115_v64 = _nw15;
        (_111_v61).m1(_112_i12, !(_23_v0), globalState);
      }
      let _116_v65;
      _116_v65 = _module.D0.create_DC0(!(_78_v33).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(89)), ((_117_v33) => function (_118_i13) {
  return _117_v33;
})(_78_v33))).length)));
      let _pat_let_tv0 = _110_v60;
      let _pat_let_tv1 = _110_v60;
      _116_v65 = function (_pat_let0_0) {
        return function (_119_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_120_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_120_dt__update_hcf0_h0);
            }(_pat_let1_0);
          }((_pat_let_tv0).IsSubsetOf(_pat_let_tv1));
        }(_pat_let0_0);
      }(_116_v65);
      r0 = _78_v33;
      r1 = new BigNumber(-928);
      r2 = _23_v0;
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _121_globalState;
      let _nw16 = new _module.GlobalState();
      _nw16.__ctor(_dafny.Seq.of(new BigNumber(879)), true);
      _121_globalState = _nw16;
      let _122_v0;
      let _init4 = function (_123_i1) {
        return !(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(645), new BigNumber(936), new BigNumber((_dafny.Seq.of(new BigNumber(901), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(945)), function (_124_i2) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        })).length), new BigNumber(989), new BigNumber(532), new BigNumber(440))).length))).length)), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(274)), function (_125_i3) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })).length))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(872)), function (_126_i4) {
          return new BigNumber(467);
        })).length)))).equals(_dafny.MultiSet.fromElements(new BigNumber(874)));
      };
      let _nw17 = Array((new BigNumber(12)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw17.length); _i0_4++) {
        _nw17[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _122_v0 = _nw17;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_122_v0).length))) {
        let _127_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_127_i0)) && ((_127_i0).isLessThan(new BigNumber((_122_v0).length))))) {
          (_122_v0)[(_127_i0)] = (_module.D0.create_DC0(true)).dtor_cf0;
        }
      }
      let _128_v1;
      _128_v1 = _module.D0.create_DC0(!(true));
      let _source2 = _128_v1;
      if (_source2.is_DC1) {
        let _129___mcc_h0 = (_source2).cf1;
        let _130___mcc_h1 = (_source2).cf2;
        let _131___mcc_h2 = (_source2).cf3;
        let _132_cf3 = _131___mcc_h2;
        let _133_cf2 = _130___mcc_h1;
        let _134_cf1 = _129___mcc_h0;
        let _135_v2;
        _135_v2 = _dafny.Seq.UnicodeFromString("hkd");
        let _136_v3;
        _136_v3 = _dafny.Map.Empty.slice().updateUnsafe(_134_cf1,_133_cf2);
        let _137_v4;
        _137_v4 = _dafny.Seq.of(new BigNumber(892), new BigNumber((_136_v3).length));
        let _138_v5;
        _138_v5 = _dafny.Map.Empty.slice().updateUnsafe(_132_cf3,_137_v4);
        if (!_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(308)), function (_139_i5) {
          return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(488)), function (_140_i6) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          })).length);
        }), (((_138_v5).contains(_132_cf3)) ? ((_138_v5).get(_132_cf3)) : (_137_v4))), new BigNumber((_135_v2).length))) {
          _134_cf1 = _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_132_cf3,_dafny.Seq.Create(_module.__default.abs(new BigNumber(669)), ((_141_cf2) => function (_142_i7) {
            return _141_cf2;
          })(_133_cf2)))).length), _133_cf2);
          let _143_v6;
          _143_v6 = _dafny.MultiSet.fromElements((new BigNumber((_135_v2).length)).isLessThanOrEqualTo(_134_cf1), _132_cf3, _132_cf3, !(_132_cf3) || (_132_cf3), !(_132_cf3) || (_module.__default.fm0(_134_cf1, _121_globalState)));
          _143_v6 = (_143_v6).Intersect(_143_v6);
          let _144_v7;
          let _145_v8;
          let _146_v9;
          let _out0;
          let _out1;
          let _out2;
          let _outcollector0 = _module.__default.m0(_dafny.MultiSet.fromElements(_137_v4, _137_v4), _121_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _144_v7 = _out0;
          _145_v8 = _out1;
          _146_v9 = _out2;
          let _147_v10;
          _147_v10 = _dafny.Map.Empty.slice().updateUnsafe(_146_v9,!(!(_146_v9) || (!((_128_v1).dtor_cf0))));
          let _148_v11;
          _148_v11 = _dafny.Seq.of(_132_cf3);
          _147_v10 = (_147_v10).update(_146_v9, (_148_v11)[_module.__default.safeIndex(_144_v7, new BigNumber((_148_v11).length))]);
          (_121_globalState).f1 = _146_v9;
        } else {
          let _149_v12;
          let _150_v13;
          let _151_v14;
          let _out3;
          let _out4;
          let _out5;
          let _outcollector1 = _module.__default.m0(_dafny.MultiSet.fromElements(_137_v4, _137_v4), _121_globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _out5 = _outcollector1[2];
          _149_v12 = _out3;
          _150_v13 = _out4;
          _151_v14 = _out5;
          (_121_globalState).f1 = (_module.__default.safeModuloInt((_dafny.ZERO).minus(_134_cf1), _module.__default.fm1(_133_cf2, _121_globalState))).isLessThanOrEqualTo(_149_v12);
          _135_v2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ne"), _135_v2), _dafny.Seq.Concat(_135_v2, _135_v2));
          (_121_globalState).f1 = false;
          let _rhs22 = _151_v14;
          let _rhs23 = ((new BigNumber(395)).minus(_134_cf1)).isEqualTo(_133_cf2);
          _151_v14 = _rhs22;
          _132_cf3 = _rhs23;
        }
        if (!(!(_132_cf3)) || ((_128_v1).dtor_cf0)) {
          let _152_v15;
          let _nw18 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
          _152_v15 = _nw18;
          let _nw19 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
          _152_v15 = _nw19;
          let _index18 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_122_v0).length));
          (_122_v0)[_index18] = _module.__default.fm0(_134_cf1, _121_globalState);
          let _index19 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_122_v0).length));
          (_122_v0)[_index19] = !(_132_cf3) || (_132_cf3);
          (_121_globalState).f0 = _137_v4;
          _134_cf1 = _module.__default.fm1((new BigNumber((_137_v4).length)).plus(_134_cf1), _121_globalState);
          _133_cf2 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(332)), function (_153_i8) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })).length), _133_cf2);
        } else {
          let _154_v16;
          let _init5 = ((_155_cf1) => function (_156_i9) {
            return (_156_i9).plus(_155_cf1);
          })(_134_cf1);
          let _nw20 = Array((new BigNumber(26)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw20.length); _i0_5++) {
            _nw20[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _154_v16 = _nw20;
          let _init6 = ((_157_cf1, _158_v3, _159_cf2) => function (_160_i10) {
            return (_160_i10).plus((((_158_v3).contains(_157_cf1)) ? ((_158_v3).get(_157_cf1)) : ((_dafny.ZERO).minus(_159_cf2))));
          })(_134_cf1, _136_v3, _133_cf2);
          let _nw21 = Array((new BigNumber(22)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw21.length); _i0_6++) {
            _nw21[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _154_v16 = _nw21;
          let _161_v17;
          _161_v17 = _dafny.Map.Empty.slice().updateUnsafe(_133_cf2,_132_cf3);
          let _162_v18;
          _162_v18 = _dafny.Seq.of(_161_v17, _161_v17);
          let _rhs24 = _132_cf3;
          let _rhs25 = _162_v18;
          let _lhs10 = _121_globalState;
          _lhs10.f1 = _rhs24;
          _162_v18 = _rhs25;
          (_121_globalState).f1 = _132_cf3;
          let _index20 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_122_v0).length));
          (_122_v0)[_index20] = _132_cf3;
          let _index21 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_122_v0).length));
          let _rhs26 = _132_cf3;
          let _rhs27 = (_134_cf1).minus(new BigNumber(653));
          let _lhs11 = _122_v0;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_122_v0).length));
          _lhs11[_lhs12] = _rhs26;
          _133_cf2 = _rhs27;
          let _163_v19;
          _163_v19 = _dafny.Map.Empty.slice().updateUnsafe((_122_v0)[_module.__default.safeIndex(new BigNumber(77), new BigNumber((_122_v0).length))],_134_cf1);
          let _164_v20;
          _164_v20 = _dafny.MultiSet.fromElements(_137_v4);
          let _165_v21;
          let _166_v22;
          let _167_v23;
          let _out6;
          let _out7;
          let _out8;
          let _outcollector2 = _module.__default.m0((_dafny.MultiSet.fromElements(_dafny.Seq.of(_133_cf2, _134_cf1), _137_v4, _dafny.Seq.of(_133_cf2, new BigNumber((_163_v19).length)), _137_v4, _137_v4)).Intersect(_164_v20), _121_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _out8 = _outcollector2[2];
          _165_v21 = _out6;
          _166_v22 = _out7;
          _167_v23 = _out8;
        }
        _133_cf2 = (_dafny.ZERO).minus(_133_cf2);
        let _168_v24;
        _168_v24 = _dafny.Map.Empty.slice().updateUnsafe(_132_cf3,true);
        _168_v24 = (_168_v24).update(((_dafny.ZERO).minus((_dafny.ZERO).minus(_134_cf1))).isLessThan(_134_cf1), _132_cf3);
      } else {
        let _169___mcc_h3 = (_source2).cf0;
        let _170_cf0 = _169___mcc_h3;
        let _171_v25;
        _171_v25 = new BigNumber(136);
        let _172_v26;
        let _init7 = ((_173_v25) => function (_174_i11) {
          return _module.__default.safeModuloInt(_174_i11, _173_v25);
        })(_171_v25);
        let _nw22 = Array((new BigNumber(2)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw22.length); _i0_7++) {
          _nw22[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _172_v26 = _nw22;
        let _index22 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_172_v26).length));
        (_172_v26)[_index22] = _module.__default.safeDivisionInt(_171_v25, _171_v25);
        let _175_v27;
        _175_v27 = _dafny.MultiSet.fromElements(_170_cf0);
        let _176_v28;
        _176_v28 = _dafny.Seq.of(_175_v27);
        let _177_v29;
        _177_v29 = _module.D0.create_DC1(_171_v25, _171_v25, _170_cf0);
        let _178_v30;
        _178_v30 = _dafny.Map.Empty.slice().updateUnsafe(!(_170_cf0),new BigNumber(-24));
        let _index23 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_172_v26).length));
        let _rhs28 = new BigNumber((((_170_cf0) ? ((_176_v28)[_module.__default.safeIndex(_171_v25, new BigNumber((_176_v28).length))]) : ((_dafny.MultiSet.fromElements((_177_v29).dtor_cf3, _170_cf0)).update(_module.__default.fm0(new BigNumber((_178_v30).length), _121_globalState), _module.__default.abs(_171_v25))))).cardinality());
        let _rhs29 = _171_v25;
        let _rhs30 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_170_cf0,_170_cf0)).length)).minus(_171_v25);
        let _rhs31 = (((_178_v30).contains(true)) ? ((_178_v30).get(true)) : (_171_v25));
        let _rhs32 = _module.__default.fm0(_171_v25, _121_globalState);
        let _lhs13 = _172_v26;
        let _lhs14 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_172_v26).length));
        _171_v25 = _rhs28;
        _171_v25 = _rhs29;
        _171_v25 = _rhs30;
        _lhs13[_lhs14] = _rhs31;
        _170_cf0 = _rhs32;
        let _179_v31;
        _179_v31 = _dafny.Set.fromElements(_178_v30, _178_v30);
        let _180_v32;
        _180_v32 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("wop"));
        let _index24 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_122_v0).length));
        (_122_v0)[_index24] = !(_180_v32).equals(_180_v32);
        let _181_v33;
        _181_v33 = _dafny.Seq.UnicodeFromString("iqinws");
        let _index25 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_122_v0).length));
        let _rhs33 = _179_v31;
        let _rhs34 = !_dafny.areEqual(_181_v33, _dafny.Seq.update(_181_v33, _module.__default.safeIndex(_171_v25, new BigNumber((_181_v33).length)), new _dafny.CodePoint('p'.codePointAt(0))));
        let _rhs35 = !(_170_cf0) || (((_dafny.ZERO).minus(_171_v25)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(786)), function (_182_i12) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        })).length)));
        let _rhs36 = (_dafny.ZERO).minus(_module.__default.fm1((_dafny.ZERO).minus(_171_v25), _121_globalState));
        let _rhs37 = _170_cf0;
        let _lhs15 = _121_globalState;
        let _lhs16 = _122_v0;
        let _lhs17 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_122_v0).length));
        _179_v31 = _rhs33;
        _lhs15.f1 = _rhs34;
        _170_cf0 = _rhs35;
        _171_v25 = _rhs36;
        _lhs16[_lhs17] = _rhs37;
        let _source3 = _177_v29;
        if (_source3.is_DC1) {
          let _183___mcc_h4 = (_source3).cf1;
          let _184___mcc_h5 = (_source3).cf2;
          let _185___mcc_h6 = (_source3).cf3;
          let _186_cf3 = _185___mcc_h6;
          let _187_cf2 = _184___mcc_h5;
          let _188_cf1 = _183___mcc_h4;
          let _189_v34;
          let _nw23 = new _module.C0();
          _nw23.__ctor();
          _189_v34 = _nw23;
          let _190_v35;
          _190_v35 = _dafny.Seq.of((_122_v0)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_122_v0).length))]);
          let _191_v36;
          _191_v36 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_122_v0)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_122_v0).length))],_171_v25)).length));
          let _192_v37;
          _192_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(496),_189_v34);
          let _rhs38 = _module.D0.create_DC0((_175_v27).equals(_175_v27));
          let _rhs39 = _dafny.Seq.of(false, (_122_v0)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_122_v0).length))], _170_cf0);
          let _rhs40 = _dafny.Seq.Concat(_dafny.Seq.Concat(_191_v36, _191_v36), _dafny.Seq.of(new BigNumber((_190_v35).length), _187_cf2, _187_cf2));
          let _rhs41 = ((_dafny.ZERO).minus(new BigNumber((_181_v33).length))).plus(new BigNumber(((_192_v37).Merge(_192_v37)).length));
          let _rhs42 = _189_v34;
          let _lhs18 = _121_globalState;
          _128_v1 = _rhs38;
          _190_v35 = _rhs39;
          _lhs18.f0 = _rhs40;
          _187_cf2 = _rhs41;
          _189_v34 = _rhs42;
          let _193_v38;
          let _nw24 = new _module.C0();
          _nw24.__ctor();
          _193_v38 = _nw24;
          (_193_v38).m1(_188_cf1, _186_cf3, _121_globalState);
        } else {
          let _194___mcc_h7 = (_source3).cf0;
          let _195_cf0 = _194___mcc_h7;
          let _196_v39;
          let _nw25 = new _module.C0();
          _nw25.__ctor();
          _196_v39 = _nw25;
          let _197_v40;
          _197_v40 = _dafny.Set.fromElements(_196_v39, _196_v39, _196_v39, _196_v39);
          _197_v40 = ((_197_v40).Intersect(_197_v40)).Difference(_dafny.Set.fromElements(_196_v39));
          let _index26 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_172_v26).length));
          (_172_v26)[_index26] = _171_v25;
          let _198_v41;
          _198_v41 = _dafny.Seq.of((_172_v26)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_172_v26).length))]);
          let _199_v42;
          _199_v42 = _dafny.MultiSet.fromElements(_198_v41);
          let _200_v43;
          let _201_v44;
          let _202_v45;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector3 = _module.__default.m0(_199_v42, _121_globalState);
          _out9 = _outcollector3[0];
          _out10 = _outcollector3[1];
          _out11 = _outcollector3[2];
          _200_v43 = _out9;
          _201_v44 = _out10;
          _202_v45 = _out11;
          _181_v33 = _181_v33;
        }
        let _203_v46;
        let _nw26 = new _module.C0();
        _nw26.__ctor();
        _203_v46 = _nw26;
        let _204_v47;
        _204_v47 = new _dafny.CodePoint('j'.codePointAt(0));
        let _205_v48;
        _205_v48 = _dafny.Map.Empty.slice().updateUnsafe(_204_v47,_203_v46);
        _203_v46 = (((_205_v48).contains(_204_v47)) ? ((_205_v48).get(_204_v47)) : (_203_v46));
      }
      let _206_v49;
      let _nw27 = new _module.C0();
      _nw27.__ctor();
      _206_v49 = _nw27;
      let _207_v50;
      _207_v50 = false;
      if (_207_v50) {
        let _208_v51;
        _208_v51 = new _dafny.CodePoint('i'.codePointAt(0));
        _208_v51 = _208_v51;
        let _209_v52;
        _209_v52 = new BigNumber(296);
        let _210_v53;
        _210_v53 = _dafny.Set.fromElements((_dafny.ZERO).minus(_209_v52), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_207_v50, _207_v50)).length)));
        _209_v52 = _module.__default.fm1(new BigNumber((((_207_v50) ? (_210_v53) : (_210_v53))).length), _121_globalState);
        _208_v51 = _208_v51;
        let _211_v54;
        let _nw28 = new _module.C0();
        _nw28.__ctor();
        _211_v54 = _nw28;
        if ((_module.__default.safeModuloInt(_209_v52, _209_v52)).isLessThanOrEqualTo(_209_v52)) {
          (_206_v49).m1((_209_v52).plus(_209_v52), _207_v50, _121_globalState);
          let _index27 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_122_v0).length));
          (_122_v0)[_index27] = _207_v50;
          let _index28 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_122_v0).length));
          (_122_v0)[_index28] = _207_v50;
          (_121_globalState).f1 = !(!((_122_v0)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_122_v0).length))]));
          let _212_v55;
          let _init8 = function (_213_i13) {
            return _dafny.Seq.UnicodeFromString("hvo");
          };
          let _nw29 = Array((new BigNumber(27)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw29.length); _i0_8++) {
            _nw29[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _212_v55 = _nw29;
          let _214_v56;
          _214_v56 = _dafny.Seq.UnicodeFromString("sdpbjuu");
          let _index29 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_212_v55).length));
          (_212_v55)[_index29] = _dafny.Seq.Concat(_214_v56, _214_v56);
          let _215_v57;
          _215_v57 = _module.D0.create_DC1(new BigNumber(692), _209_v52, (_122_v0)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_122_v0).length))]);
          let _216_v58;
          _216_v58 = _dafny.MultiSet.fromElements((_122_v0)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_122_v0).length))], (_122_v0)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_122_v0).length))], _207_v50, (_122_v0)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_122_v0).length))], !((_122_v0)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_122_v0).length))]));
          let _index30 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_212_v55).length));
          (_212_v55)[_index30] = ((!((_module.__default.fm5((_122_v0)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_122_v0).length))], (_122_v0)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_122_v0).length))], _215_v57, _121_globalState)).IsDisjointFrom((_module.D2.create_DC3(_dafny.MultiSet.fromElements(_209_v52, new BigNumber(((_216_v58).update((_122_v0)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_122_v0).length))], _module.__default.abs(_209_v52))).cardinality())))).dtor_cf5))) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(196)), ((_217_v51) => function (_218_i14) {
            return _217_v51;
          })(_208_v51))) : (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-556)), ((_219_v51) => function (_220_i15) {
            return _219_v51;
          })(_208_v51)), _214_v56)));
          let _221_v59;
          _221_v59 = _dafny.Seq.of(_209_v52, new BigNumber((_216_v58).cardinality()), (_dafny.ZERO).minus(_209_v52));
          _209_v52 = ((_221_v59)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_221_v59).length))]).multipliedBy((_209_v52).multipliedBy(_209_v52));
        } else {
          let _222_v60;
          let _init9 = function (_223_i16) {
            return (_223_i16).plus(new BigNumber((_dafny.Seq.of(false)).length));
          };
          let _nw30 = Array((new BigNumber(24)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw30.length); _i0_9++) {
            _nw30[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _222_v60 = _nw30;
          let _index31 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_222_v60).length));
          (_222_v60)[_index31] = (_dafny.ZERO).minus(_209_v52);
          let _index32 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_222_v60).length));
          (_222_v60)[_index32] = _209_v52;
          let _224_v61;
          _224_v61 = _dafny.Map.Empty.slice().updateUnsafe(_208_v51,_209_v52);
          let _225_v62;
          _225_v62 = _dafny.Map.Empty.slice().updateUnsafe((_222_v60)[_module.__default.safeIndex(new BigNumber(922), new BigNumber((_222_v60).length))],(_222_v60)[_module.__default.safeIndex(new BigNumber(922), new BigNumber((_222_v60).length))]);
          _224_v61 = (_224_v61).update(_208_v51, _module.__default.safeDivisionInt(new BigNumber((_225_v62).length), new BigNumber(196)));
          let _226_v63;
          _226_v63 = _dafny.MultiSet.fromElements(_208_v51, new _dafny.CodePoint('w'.codePointAt(0)), _208_v51, _208_v51, _208_v51);
          let _227_v64;
          _227_v64 = _dafny.Seq.of(_208_v51, _208_v51, new _dafny.CodePoint('a'.codePointAt(0)), _208_v51, _208_v51);
          let _228_v65;
          _228_v65 = _dafny.Seq.of(_226_v63, _226_v63, (_226_v63).update(_208_v51, _module.__default.abs(_209_v52)), _dafny.MultiSet.FromArray(_227_v64), _226_v63);
          _226_v63 = (_228_v65)[_module.__default.safeIndex(_209_v52, new BigNumber((_228_v65).length))];
          _209_v52 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(_209_v52, _209_v52), _209_v52);
          (_206_v49).m1((_222_v60)[_module.__default.safeIndex(new BigNumber(922), new BigNumber((_222_v60).length))], ((_207_v50) ? (!(_207_v50)) : (_207_v50)), _121_globalState);
        }
      } else {
        let _229_v66;
        _229_v66 = _dafny.MultiSet.fromElements(!(true));
        let _230_v67;
        _230_v67 = new BigNumber(399);
        let _rhs43 = _229_v66;
        let _rhs44 = _230_v67;
        _229_v66 = _rhs43;
        _230_v67 = _rhs44;
        _230_v67 = new BigNumber(644);
        _206_v49 = _206_v49;
        let _231_v68;
        let _nw31 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _231_v68 = _nw31;
        let _index33 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_231_v68).length));
        (_231_v68)[_index33] = _module.__default.fm1(_230_v67, _121_globalState);
        let _index34 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_231_v68).length));
        (_231_v68)[_index34] = _230_v67;
        let _232_v69;
        _232_v69 = _dafny.Seq.of(_207_v50);
        let _233_v70;
        _233_v70 = _dafny.MultiSet.fromElements(_232_v69);
        let _234_v71;
        _234_v71 = _dafny.Map.Empty.slice().updateUnsafe((_231_v68)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_231_v68).length))],_233_v70);
        _207_v50 = ((((((_234_v71).contains(_230_v67)) ? ((_234_v71).get(_230_v67)) : (_233_v70))).IsDisjointFrom(_233_v70)) ? (true) : (_207_v50));
      }
      let _235_v72;
      _235_v72 = new BigNumber(694);
      _235_v72 = _235_v72;
      let _236_v73;
      let _init10 = ((_237_v50) => function (_238_i18) {
        return _dafny.Seq.of(true, _237_v50);
      })(_207_v50);
      let _nw32 = Array((new BigNumber(3)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw32.length); _i0_10++) {
        _nw32[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _236_v73 = _nw32;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_236_v73).length))) {
        let _239_i17 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_239_i17)) && ((_239_i17).isLessThan(new BigNumber((_236_v73).length))))) {
          (_236_v73)[(_239_i17)] = _dafny.Seq.of(_207_v50, _207_v50, _207_v50);
        }
      }
      let _240_v74;
      let _nw33 = new _module.C0();
      _nw33.__ctor();
      _240_v74 = _nw33;
      if (_module.__default.fm0((_235_v72).minus(_235_v72), _121_globalState)) {
        let _241_v75;
        _241_v75 = new _dafny.CodePoint('o'.codePointAt(0));
        let _242_v76;
        _242_v76 = _dafny.Map.Empty.slice().updateUnsafe(_207_v50,_241_v75);
        let _243_v77;
        _243_v77 = _dafny.Map.Empty.slice().updateUnsafe(_207_v50,_242_v76);
        let _244_v78;
        _244_v78 = (((_243_v77).contains(_207_v50)) ? ((_243_v77).get(_207_v50)) : (_242_v76));
        let _245_v79;
        _245_v79 = _dafny.Seq.of(new BigNumber((((_207_v50) ? (_242_v76) : ((_244_v78)))).length));
        let _rhs45 = _235_v72;
        let _rhs46 = _245_v79;
        let _lhs19 = _121_globalState;
        _235_v72 = _rhs45;
        _lhs19.f0 = _rhs46;
        let _246_v80;
        _246_v80 = _dafny.MultiSet.fromElements(_235_v72, _235_v72);
        let _247_v81;
        _247_v81 = _dafny.Map.Empty.slice().updateUnsafe(_235_v72,_207_v50);
        let _248_v82;
        _248_v82 = _module.D2.create_DC5(_235_v72, _235_v72);
        let _249_v84;
        let _nw34 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _249_v84 = _nw34;
        let _250_v85;
        _250_v85 = _dafny.MultiSet.fromElements(_249_v84);
        let _251_v86;
        let _nw35 = Array((new BigNumber(9)).toNumber());
        _nw35[(_dafny.ZERO).toNumber()] = new BigNumber((_246_v80).cardinality());
        _nw35[(_dafny.ONE).toNumber()] = _235_v72;
        _nw35[(new BigNumber(2)).toNumber()] = new BigNumber((_module.__default.fm6(_207_v50, _247_v81, (_module.D0.create_DC0(_207_v50)).dtor_cf0, _235_v72, _121_globalState)).length);
        _nw35[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_235_v72);
        _nw35[(new BigNumber(4)).toNumber()] = (_248_v82).dtor_cf8;
        _nw35[(new BigNumber(5)).toNumber()] = _module.__default.fm1(new BigNumber((function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of _dafny.IntegerRange(new BigNumber(144), new BigNumber(803))) {
            let _252_v83 = _compr_7;
            if (((new BigNumber(144)).isLessThanOrEqualTo(_252_v83)) && ((_252_v83).isLessThan(new BigNumber(803)))) {
              _coll7.push([(_252_v83).plus(_235_v72),new BigNumber(785)]);
            }
          }
          return _coll7;
        }()).length), _121_globalState);
        _nw35[(new BigNumber(6)).toNumber()] = new BigNumber((_250_v85).cardinality());
        _nw35[(new BigNumber(7)).toNumber()] = _235_v72;
        _nw35[(new BigNumber(8)).toNumber()] = _235_v72;
        _251_v86 = _nw35;
        let _253_v87;
        _253_v87 = _dafny.Seq.of(_249_v84, _249_v84);
        let _254_v88;
        _254_v88 = _dafny.Set.fromElements(new BigNumber((_253_v87).length));
        _254_v88 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_245_v79, _module.__default.safeIndex(_235_v72, new BigNumber((_245_v79).length)), _235_v72)).length)), _235_v72);
        let _255_v89;
        _255_v89 = _dafny.Seq.UnicodeFromString("hlrf");
        let _256_v90;
        _256_v90 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_235_v72),((!(!(_207_v50))) ? (_dafny.Seq.UnicodeFromString("iff")) : (_255_v89)));
        _255_v89 = (((_256_v90).contains(_235_v72)) ? ((_256_v90).get(_235_v72)) : (_255_v89));
        let _index35 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_251_v86).length));
        (_251_v86)[_index35] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_207_v50,_235_v72)).length);
        let _index36 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_251_v86).length));
        (_251_v86)[_index36] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat(_255_v89, _255_v89)).length), (_245_v79)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_245_v79).length))]);
        _206_v49 = _240_v74;
      } else {
        let _257_v91;
        let _nw36 = new _module.C0();
        _nw36.__ctor();
        _257_v91 = _nw36;
        _206_v49 = _240_v74;
        (_240_v74).m1(_235_v72, true, _121_globalState);
        let _258_v92;
        _258_v92 = _dafny.Seq.of((new BigNumber(256)).isLessThan(_235_v72));
        _207_v50 = (_258_v92)[_module.__default.safeIndex(new BigNumber(-242), new BigNumber((_258_v92).length))];
        let _259_v93;
        _259_v93 = _dafny.Seq.of((_dafny.ZERO).minus(_235_v72), ((_207_v50) ? (_235_v72) : (_235_v72)), _235_v72);
        _235_v72 = (_259_v93)[_module.__default.safeIndex(new BigNumber(-252), new BigNumber((_259_v93).length))];
      }
      let _260_v94;
      _260_v94 = _dafny.MultiSet.fromElements(_240_v74, _240_v74, _240_v74);
      _260_v94 = _260_v94;
      let _hi1 = _235_v72;
      for (let _261_i19 = new BigNumber(762); _261_i19.isLessThan(_hi1); _261_i19 = _261_i19.plus(_dafny.ONE)) {
        (_121_globalState).f1 = _207_v50;
        (_121_globalState).f1 = _207_v50;
        _235_v72 = _235_v72;
        let _262_v95;
        let _init11 = ((_263_v72) => function (_264_i20) {
          return _module.__default.safeDivisionInt(_264_i20, _263_v72);
        })(_235_v72);
        let _nw37 = Array((new BigNumber(25)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw37.length); _i0_11++) {
          _nw37[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _262_v95 = _nw37;
        let _index37 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_262_v95).length));
        (_262_v95)[_index37] = (_dafny.ZERO).minus(_261_i19);
        let _index38 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_262_v95).length));
        (_262_v95)[_index38] = _261_i19;
      }
      let _hi2 = _235_v72;
      for (let _265_i21 = new BigNumber((_dafny.Seq.UnicodeFromString("q")).length); _265_i21.isLessThan(_hi2); _265_i21 = _265_i21.plus(_dafny.ONE)) {
        let _266_v96;
        _266_v96 = _dafny.MultiSet.fromElements(_265_i21);
        _235_v72 = ((((_266_v96).contains(new BigNumber(247))) ? ((_266_v96).get(new BigNumber(247))) : (_265_i21))).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(92)), function (_267_i22) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        })).length)));
        let _268_v97;
        _268_v97 = new _dafny.CodePoint('f'.codePointAt(0));
        _268_v97 = _268_v97;
        if (false) {
          let _269_v98;
          let _nw38 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _269_v98 = _nw38;
          let _index39 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_269_v98).length));
          (_269_v98)[_index39] = (_module.__default.fm1(new BigNumber(-790), _121_globalState)).minus(_235_v72);
          let _270_v99;
          _270_v99 = _dafny.Seq.UnicodeFromString("ensayrry");
          let _index40 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_269_v98).length));
          (_269_v98)[_index40] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm3(_121_globalState), _270_v99), _270_v99)).length);
          let _271_v100;
          _271_v100 = _module.D4.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_235_v72, _121_globalState),(_269_v98)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_269_v98).length))]));
          let _272_v101;
          _272_v101 = _dafny.Map.Empty.slice().updateUnsafe(false,(_269_v98)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_269_v98).length))]);
          let _273_v102;
          _273_v102 = _dafny.Map.Empty.slice().updateUnsafe(((_271_v100).dtor_cf12).Merge(_272_v101),_268_v97);
          _273_v102 = ((_273_v102).update(_272_v101, _268_v97)).Merge(_273_v102);
          _268_v97 = new _dafny.CodePoint('k'.codePointAt(0));
          let _274_v103;
          let _nw39 = new _module.C0();
          _nw39.__ctor();
          _274_v103 = _nw39;
          let _275_v104;
          _275_v104 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),_235_v72);
          let _index41 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_269_v98).length));
          (_269_v98)[_index41] = ((_dafny.ZERO).minus(new BigNumber(-723))).minus(_module.__default.safeDivisionInt(new BigNumber((_275_v104).length), (_269_v98)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_269_v98).length))]));
        } else {
          let _276_v105;
          let _nw40 = new _module.C0();
          _nw40.__ctor();
          _276_v105 = _nw40;
          _207_v50 = true;
          let _277_v106;
          _277_v106 = _dafny.Seq.UnicodeFromString("sxn");
          _277_v106 = _277_v106;
          let _278_v107;
          _278_v107 = _dafny.MultiSet.fromElements(_207_v50);
          let _279_v108;
          _279_v108 = _module.D5.create_DC10(_278_v107);
          (_121_globalState).f1 = ((_278_v107).Union((_279_v108).dtor_cf16)).IsProperSubsetOf((_278_v107).Difference(_278_v107));
          _277_v106 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_277_v106, _module.__default.safeIndex(_265_i21, new BigNumber((_277_v106).length)), _268_v97), _module.__default.fm3(_121_globalState)), _277_v106);
        }
        let _280_v109;
        _280_v109 = _dafny.Set.fromElements(_235_v72);
        let _281_v110;
        _281_v110 = _dafny.Seq.UnicodeFromString("ybjqinu");
        _235_v72 = (((_207_v50) ? (new BigNumber((_280_v109).length)) : (new BigNumber((_281_v110).length)))).plus((_235_v72).minus(_235_v72));
      }
      let _282_v111;
      _282_v111 = _dafny.MultiSet.fromElements(_207_v50);
      (_206_v49).m1(new BigNumber((_282_v111).cardinality()), _207_v50, _121_globalState);
      (_240_v74).m1(_235_v72, _207_v50, _121_globalState);
      let _hi3 = new BigNumber(285);
      for (let _283_i23 = new BigNumber(188); _283_i23.isLessThan(_hi3); _283_i23 = _283_i23.plus(_dafny.ONE)) {
        let _284_v112;
        _284_v112 = _dafny.Seq.UnicodeFromString("imxv");
        let _285_v113;
        _285_v113 = _dafny.Seq.of(_283_i23, (_dafny.ZERO).minus(_283_i23), (_235_v72).multipliedBy(_module.__default.fm1(new BigNumber((_284_v112).length), _121_globalState)));
        _235_v72 = (_285_v113)[_module.__default.safeIndex(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(612)), function (_286_i24) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        })).length))).plus(_module.__default.fm1(_283_i23, _121_globalState)), new BigNumber((_285_v113).length))];
        _236_v73 = _236_v73;
        if (!((_235_v72).multipliedBy(new BigNumber(-552))).isEqualTo(new BigNumber((_module.__default.fm7(_121_globalState)).length))) {
          let _287_v114;
          _287_v114 = _module.D2.create_DC5(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-94)), ((_288_v72) => function (_289_i25) {
  return _288_v72;
})(_235_v72)))).cardinality()), new BigNumber(902));
          let _290_v115;
          _290_v115 = _dafny.MultiSet.fromElements(_287_v114, _287_v114);
          let _291_v116;
          _291_v116 = _dafny.Set.fromElements(_283_i23, new BigNumber(914));
          _235_v72 = (((_290_v115).contains(_287_v114)) ? ((_290_v115).get(_287_v114)) : ((_235_v72).multipliedBy(new BigNumber((_291_v116).length))));
          let _292_v117;
          _292_v117 = new _dafny.CodePoint('b'.codePointAt(0));
          let _293_v118;
          _293_v118 = _dafny.MultiSet.fromElements(_292_v117, new _dafny.CodePoint('p'.codePointAt(0)));
          let _294_v119;
          _294_v119 = _dafny.Seq.of(_207_v50, _module.__default.fm0(_283_i23, _121_globalState));
          let _295_v120;
          let _nw41 = Array((new BigNumber(5)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = _235_v72;
          _nw41[(_dafny.ONE).toNumber()] = (_285_v113)[_module.__default.safeIndex((((_293_v118).contains(_292_v117)) ? ((_293_v118).get(_292_v117)) : (new BigNumber((_294_v119).length))), new BigNumber((_285_v113).length))];
          _nw41[(new BigNumber(2)).toNumber()] = _235_v72;
          _nw41[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(_235_v72, new BigNumber((_294_v119).length));
          _nw41[(new BigNumber(4)).toNumber()] = ((_207_v50) ? (new BigNumber(154)) : (_283_i23));
          _295_v120 = _nw41;
          let _index42 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_295_v120).length));
          (_295_v120)[_index42] = _235_v72;
          let _index43 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_295_v120).length));
          (_295_v120)[_index43] = _283_i23;
          (_121_globalState).f1 = _207_v50;
          let _296_v121;
          _296_v121 = _dafny.Map.Empty.slice().updateUnsafe(_207_v50,true);
          let _297_v122;
          let _init12 = function (_298_i26) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          };
          let _nw42 = Array((new BigNumber(17)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw42.length); _i0_12++) {
            _nw42[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _297_v122 = _nw42;
          let _index44 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_297_v122).length));
          (_297_v122)[_index44] = _292_v117;
          let _299_v123;
          let _init13 = ((_300_v111) => function (_301_i27) {
            return _300_v111;
          })(_282_v111);
          let _nw43 = Array((new BigNumber(9)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw43.length); _i0_13++) {
            _nw43[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _299_v123 = _nw43;
          let _302_v124;
          _302_v124 = _module.D5.create_DC10(_282_v111);
          let _index45 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_295_v120).length));
          let _index46 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_297_v122).length));
          let _rhs47 = (_282_v111).Union((_302_v124).dtor_cf16);
          let _rhs48 = _296_v121;
          let _rhs49 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_283_i23));
          let _rhs50 = _292_v117;
          let _rhs51 = _299_v123;
          let _lhs20 = _295_v120;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_295_v120).length));
          let _lhs22 = _297_v122;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_297_v122).length));
          _282_v111 = _rhs47;
          _296_v121 = _rhs48;
          _lhs20[_lhs21] = _rhs49;
          _lhs22[_lhs23] = _rhs50;
          _299_v123 = _rhs51;
          (_121_globalState).f1 = _module.__default.fm0(_module.__default.fm1(new BigNumber((_dafny.Seq.UnicodeFromString("ehbbvaf")).length), _121_globalState), _121_globalState);
        } else {
          _235_v72 = _235_v72;
          let _303_v125;
          _303_v125 = _module.D0.create_DC1(_283_i23, _235_v72, !(_dafny.Map.Empty.slice().updateUnsafe(_283_i23,_207_v50)).contains(_235_v72));
          _303_v125 = _303_v125;
          let _304_v127;
          let _init14 = ((_305_v72) => function (_306_i28) {
            return _module.__default.safeModuloInt(_306_i28, new BigNumber((function () {
              let _coll8 = new _dafny.Map();
              for (const _compr_8 of _dafny.IntegerRange(new BigNumber(374), new BigNumber(471))) {
                let _307_v126 = _compr_8;
                if (((new BigNumber(374)).isLessThanOrEqualTo(_307_v126)) && ((_307_v126).isLessThan(new BigNumber(471)))) {
                  _coll8.push([(_307_v126).multipliedBy(new BigNumber(-80)),_305_v72]);
                }
              }
              return _coll8;
            }()).length));
          })(_235_v72);
          let _nw44 = Array((new BigNumber(24)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw44.length); _i0_14++) {
            _nw44[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _304_v127 = _nw44;
          let _index47 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_304_v127).length));
          (_304_v127)[_index47] = _283_i23;
          let _308_v128;
          _308_v128 = new _dafny.CodePoint('o'.codePointAt(0));
          let _309_v129;
          _309_v129 = _dafny.Seq.of(_dafny.Seq.update(_284_v112, _module.__default.safeIndex(new BigNumber(525), new BigNumber((_284_v112).length)), _308_v128));
          let _310_v130;
          let _nw45 = Array((new BigNumber(18)).toNumber());
          _nw45[(_dafny.ZERO).toNumber()] = _284_v112;
          _nw45[(_dafny.ONE).toNumber()] = _284_v112;
          _nw45[(new BigNumber(2)).toNumber()] = _284_v112;
          _nw45[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_284_v112, _284_v112);
          _nw45[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("cscllebh");
          _nw45[(new BigNumber(5)).toNumber()] = _284_v112;
          _nw45[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_284_v112, _module.__default.safeIndex(_283_i23, new BigNumber((_284_v112).length)), _308_v128), _dafny.Seq.Create(_module.__default.abs(new BigNumber(860)), function (_311_i29) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          })), _module.__default.safeIndex(_283_i23, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_284_v112, _module.__default.safeIndex(_283_i23, new BigNumber((_284_v112).length)), _308_v128), _dafny.Seq.Create(_module.__default.abs(new BigNumber(860)), function (_312_i29) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }))).length)), _308_v128);
          _nw45[(new BigNumber(7)).toNumber()] = (_309_v129)[_module.__default.safeIndex((_285_v113)[_module.__default.safeIndex(_283_i23, new BigNumber((_285_v113).length))], new BigNumber((_309_v129).length))];
          _nw45[(new BigNumber(8)).toNumber()] = _284_v112;
          _nw45[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_284_v112, _284_v112);
          _nw45[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_284_v112, _284_v112);
          _nw45[(new BigNumber(11)).toNumber()] = _284_v112;
          _nw45[(new BigNumber(12)).toNumber()] = _284_v112;
          _nw45[(new BigNumber(13)).toNumber()] = _284_v112;
          _nw45[(new BigNumber(14)).toNumber()] = _284_v112;
          _nw45[(new BigNumber(15)).toNumber()] = _284_v112;
          _nw45[(new BigNumber(16)).toNumber()] = ((false) ? (_284_v112) : (_284_v112));
          _nw45[(new BigNumber(17)).toNumber()] = _284_v112;
          _310_v130 = _nw45;
          let _313_v131;
          _313_v131 = _dafny.Map.Empty.slice().updateUnsafe(_283_i23,_235_v72);
          let _314_v132;
          _314_v132 = _dafny.Map.Empty.slice().updateUnsafe(_207_v50,_240_v74);
          let _315_v133;
          _315_v133 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((((_313_v131).contains(_283_i23)) ? ((_313_v131).get(_283_i23)) : (new BigNumber((_314_v132).length)))));
          let _index48 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_304_v127).length));
          let _rhs52 = !(_207_v50);
          let _rhs53 = (_235_v72).multipliedBy(new BigNumber(((_315_v133).Union(_315_v133)).cardinality()));
          let _rhs54 = _310_v130;
          let _rhs55 = (_dafny.ZERO).minus((_283_i23).plus(_283_i23));
          let _lhs24 = _304_v127;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_304_v127).length));
          _207_v50 = _rhs52;
          _lhs24[_lhs25] = _rhs53;
          _310_v130 = _rhs54;
          _235_v72 = _rhs55;
          _304_v127 = _304_v127;
          _207_v50 = !(_207_v50) || (_module.__default.fm0(new BigNumber(340), _121_globalState));
        }
        (_121_globalState).f1 = _207_v50;
      }
      if (!(_module.__default.fm0(_module.__default.safeModuloInt(_235_v72, _235_v72), _121_globalState))) {
        let _316_v134;
        _316_v134 = _dafny.Seq.UnicodeFromString("hceiynyhv");
        _316_v134 = _316_v134;
        let _317_v135;
        let _nw46 = new _module.C0();
        _nw46.__ctor();
        _317_v135 = _nw46;
        let _318_v136;
        let _nw47 = Array((new BigNumber(19)).toNumber()).fill([]);
        _318_v136 = _nw47;
        let _319_v137;
        _319_v137 = _dafny.MultiSet.fromElements(_235_v72, (_dafny.ZERO).minus(_235_v72), new BigNumber((_316_v134).length));
        let _320_v138;
        _320_v138 = _dafny.Map.Empty.slice().updateUnsafe(_207_v50,_235_v72);
        let _321_v139;
        _321_v139 = _dafny.Map.Empty.slice().updateUnsafe(((true) ? (_318_v136) : (_318_v136)),new BigNumber(((_319_v137).Union((_319_v137).update(_235_v72, _module.__default.abs(new BigNumber((_320_v138).length))))).cardinality()));
        _321_v139 = (_321_v139).update(_318_v136, _235_v72);
        (_121_globalState).f1 = _207_v50;
        let _322_v140;
        _322_v140 = _dafny.Map.Empty.slice().updateUnsafe(_235_v72,new BigNumber(573));
        let _323_v141;
        _323_v141 = _dafny.MultiSet.fromElements(_322_v140, _322_v140);
        let _324_v142;
        _324_v142 = _dafny.Map.Empty.slice().updateUnsafe(_207_v50,_207_v50);
        if (!((((_324_v142).contains(false)) ? ((_324_v142).get(false)) : (_207_v50))) || ((_323_v141).IsSubsetOf(_323_v141))) {
          let _325_v143;
          let _nw48 = new _module.C0();
          _nw48.__ctor();
          _325_v143 = _nw48;
          (_121_globalState).f1 = ((_207_v50) ? (_207_v50) : (_207_v50));
          let _326_v144;
          _326_v144 = _dafny.Seq.of(_235_v72);
          let _index49 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_122_v0).length));
          (_122_v0)[_index49] = _dafny.areEqual(_326_v144, _326_v144);
          let _327_v145;
          _327_v145 = _module.D0.create_DC1(_235_v72, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-55),_235_v72)).length), _207_v50);
          let _328_v146;
          _328_v146 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber(-268), _121_globalState),_326_v144);
          let _index50 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_122_v0).length));
          (_122_v0)[_index50] = _dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(_235_v72, (_dafny.ZERO).minus(new BigNumber((_module.__default.fm5(_207_v50, _207_v50, _327_v145, _121_globalState)).cardinality()))), (((_328_v146).contains(_207_v50)) ? ((_328_v146).get(_207_v50)) : (_326_v144))), _module.__default.fm1(_235_v72, _121_globalState));
          _207_v50 = (_207_v50) && (true);
          _235_v72 = _module.__default.safeModuloInt((_235_v72).plus(_235_v72), _235_v72);
        } else {
          _235_v72 = ((new BigNumber((_316_v134).length)).minus(_235_v72)).multipliedBy(_235_v72);
          let _329_v147;
          let _nw49 = Array((new BigNumber(25)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _329_v147 = _nw49;
          let _330_v148;
          let _nw50 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _330_v148 = _nw50;
          let _331_v149;
          _331_v149 = _dafny.Map.Empty.slice().updateUnsafe(_122_v0,_329_v147);
          let _332_v150;
          _332_v150 = _dafny.Seq.of((((_331_v149).contains(_122_v0)) ? ((_331_v149).get(_122_v0)) : (_329_v147)), _329_v147, _329_v147, _329_v147);
          let _rhs56 = (_332_v150)[_module.__default.safeIndex(_module.__default.fm1(_235_v72, _121_globalState), new BigNumber((_332_v150).length))];
          let _rhs57 = _235_v72;
          let _rhs58 = !((_207_v50) && (_207_v50)) || (_207_v50);
          let _rhs59 = _330_v148;
          let _lhs26 = _121_globalState;
          _329_v147 = _rhs56;
          _235_v72 = _rhs57;
          _lhs26.f1 = _rhs58;
          _330_v148 = _rhs59;
          let _333_v151;
          _333_v151 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_207_v50)).length),_207_v50);
          (_121_globalState).f1 = _module.__default.fm0((_235_v72).minus(new BigNumber((_333_v151).length)), _121_globalState);
          let _334_v152;
          _334_v152 = new _dafny.CodePoint('h'.codePointAt(0));
          let _335_v153;
          _335_v153 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_module.__default.fm8(_207_v50, _dafny.MultiSet.fromElements(_334_v152), _121_globalState), _dafny.Seq.of(_235_v72, _235_v72)),_317_v135);
          let _336_v154;
          _336_v154 = _dafny.Seq.of(_235_v72, _235_v72);
          _317_v135 = (((_335_v153).contains(_336_v154)) ? ((_335_v153).get(_336_v154)) : (((_207_v50) ? (_317_v135) : (_317_v135))));
          let _337_v155;
          let _nw51 = Array((new BigNumber(8)).toNumber());
          _337_v155 = _nw51;
          _337_v155 = _337_v155;
        }
      } else {
        (_121_globalState).f1 = _207_v50;
        let _338_v156;
        _338_v156 = _dafny.Seq.of(_235_v72);
        let _339_v157;
        _339_v157 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.update(_338_v156, _module.__default.safeIndex(new BigNumber((_282_v111).cardinality()), new BigNumber((_338_v156).length)), _235_v72)).length)).minus(_235_v72),_207_v50);
        (_121_globalState).f1 = (((_339_v157).contains(_235_v72)) ? ((_339_v157).get(_235_v72)) : (_207_v50));
        let _340_v158;
        _340_v158 = _dafny.Seq.of(_207_v50, _207_v50);
        _340_v158 = _dafny.Seq.of(_207_v50, _207_v50);
        (_121_globalState).f1 = ((_282_v111).Intersect(_282_v111)).IsProperSubsetOf(_282_v111);
        let _341_v159;
        let _nw52 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _341_v159 = _nw52;
        let _342_v160;
        _342_v160 = _dafny.MultiSet.fromElements(_341_v159);
        let _343_v161;
        _343_v161 = _342_v160;
        let _344_v162;
        _344_v162 = _dafny.Map.Empty.slice().updateUnsafe(_235_v72,new BigNumber((((_343_v161)).Intersect(_342_v160)).cardinality()));
        let _345_v163;
        _345_v163 = new _dafny.CodePoint('p'.codePointAt(0));
        let _346_v164;
        _346_v164 = _dafny.Map.Empty.slice().updateUnsafe(_207_v50,_345_v163);
        let _347_v165;
        _347_v165 = _dafny.Map.Empty.slice().updateUnsafe(_346_v164,new BigNumber(133));
        _344_v162 = (_344_v162).update(_235_v72, (((_347_v165).contains(_346_v164)) ? ((_347_v165).get(_346_v164)) : (_235_v72)));
      }
      let _348_v166;
      _348_v166 = _dafny.Seq.of(true);
      _207_v50 = _dafny.Seq.IsPrefixOf(_348_v166, _348_v166);
      process.stdout.write(_dafny.toString(_dafny.areEqual(_121_globalState.f0, _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_121_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_128_v1).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_207_v50));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_235_v72));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_236_v73)[_dafny.ZERO], _dafny.Seq.of(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_236_v73)[_dafny.ONE], _dafny.Seq.of(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_236_v73)[new BigNumber(2)], _dafny.Seq.of(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_260_v94).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v111).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_348_v166, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC2(cf4) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + this.cf4.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.UnicodeFromString("");
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf6, cf7) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC5(cf8, cf9) {
      let $dt = new D2(1);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC3(cf5) {
      let $dt = new D2(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC6(cf10) {
      let $dt = new D2(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf6 === other.cf6 && this.cf7 === other.cf7;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4(false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf11) {
      let $dt = new D3(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9(cf13, cf14, cf15) {
      let $dt = new D4(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC8(cf12) {
      let $dt = new D4(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC9([], _dafny.ZERO, []);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC11(cf17) {
      let $dt = new D5(0);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC12(cf18, cf19, cf20, cf21, cf22) {
      let $dt = new D5(1);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC10(cf16) {
      let $dt = new D5(2);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + this.cf21.toVerbatimString(true) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC10" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf17 === other.cf17;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20 && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC11(null);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC13(cf23) {
      let $dt = new D6(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf25) {
      let $dt = new D7(0);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC14(cf24) {
      let $dt = new D7(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC15" + "(" + this.cf25.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC15(_dafny.Seq.UnicodeFromString(""));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.Seq.of();
      this.f1 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1) {
      let _this = this;
      (_this).f0 = f0;
      (_this).f1 = f1;
      return;
    }
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let _hi4 = p0;
      for (let _349_i0 = p0; _349_i0.isLessThan(_hi4); _349_i0 = _349_i0.plus(_dafny.ONE)) {
        let _350_v0;
        _350_v0 = new BigNumber(-15);
        let _351_v2;
        _351_v2 = new _dafny.CodePoint('e'.codePointAt(0));
        let _352_v3;
        _352_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_351_v2)).length),p0);
        _350_v0 = (_350_v0).multipliedBy(new BigNumber(((function () {
          let _coll9 = new _dafny.Map();
          for (const _compr_9 of _dafny.IntegerRange(new BigNumber(666), new BigNumber(-533))) {
            let _353_v1 = _compr_9;
            if (((new BigNumber(666)).isLessThanOrEqualTo(_353_v1)) && ((_353_v1).isLessThan(new BigNumber(-533)))) {
              _coll9.push([_module.__default.safeModuloInt(_353_v1, _350_v0),new BigNumber(((_dafny.Seq.UnicodeFromString("dchcbikbk"))).length)]);
            }
          }
          return _coll9;
        }()).Merge(_352_v3)).length));
        let _354_v4;
        _354_v4 = _dafny.Seq.of(_349_i0);
        let _355_v5;
        _355_v5 = _dafny.Seq.of(_354_v4);
        let _356_v6;
        _356_v6 = _dafny.Map.Empty.slice().updateUnsafe(_350_v0,!(p1));
        (globalState).f0 = (_355_v5)[_module.__default.safeIndex(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p0,p1)).Merge(_356_v6)).length), new BigNumber((_355_v5).length))];
        (globalState).f1 = ((_module.D0.create_DC1(_350_v0, _module.__default.fm1(_349_i0, globalState), p1)).dtor_cf1).isLessThan(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(p1), _module.__default.safeIndex(_349_i0, new BigNumber((_dafny.Seq.of(p1)).length)), p1)).length));
        (globalState).f1 = p1;
      }
      (globalState).f1 = p1;
      let _357_v7;
      _357_v7 = new BigNumber(436);
      let _358_v8;
      _358_v8 = _dafny.MultiSet.fromElements(p1, p1);
      _357_v7 = _module.__default.safeDivisionInt(_module.__default.fm1(p0, globalState), (new BigNumber(803)).plus(new BigNumber((_358_v8).cardinality())));
      let _359_v9;
      _359_v9 = _dafny.Seq.UnicodeFromString("hb");
      let _360_v10;
      _360_v10 = _dafny.Seq.of(p1);
      let _361_v11;
      _361_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_module.__default.fm0(_357_v7, globalState))).length),p1);
      let _362_v12;
      let _nw53 = Array((new BigNumber(20)).toNumber());
      _nw53[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(true, p1);
      _nw53[(_dafny.ONE).toNumber()] = _360_v10;
      _nw53[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_360_v10, _module.__default.safeIndex(_357_v7, new BigNumber((_360_v10).length)), p1), _360_v10);
      _nw53[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_360_v10, _360_v10);
      _nw53[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(p1, p1);
      _nw53[(new BigNumber(5)).toNumber()] = _360_v10;
      _nw53[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(p1, p1);
      _nw53[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_360_v10, _module.__default.safeIndex(new BigNumber(312), new BigNumber((_360_v10).length)), p1);
      _nw53[(new BigNumber(8)).toNumber()] = _360_v10;
      _nw53[(new BigNumber(9)).toNumber()] = _360_v10;
      _nw53[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_360_v10, _360_v10);
      _nw53[(new BigNumber(11)).toNumber()] = _360_v10;
      _nw53[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(p1, p1, p1);
      _nw53[(new BigNumber(13)).toNumber()] = _360_v10;
      _nw53[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(p1, p1, p1);
      _nw53[(new BigNumber(15)).toNumber()] = _module.__default.fm2(p1, p1, _361_v11, p1, globalState);
      _nw53[(new BigNumber(16)).toNumber()] = _360_v10;
      _nw53[(new BigNumber(17)).toNumber()] = ((p1) ? (_360_v10) : (_360_v10));
      _nw53[(new BigNumber(18)).toNumber()] = _360_v10;
      _nw53[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(true, p1);
      _362_v12 = _nw53;
      let _index51 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_362_v12).length));
      (_362_v12)[_index51] = _dafny.Seq.Concat(_360_v10, _360_v10);
      let _index52 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_362_v12).length));
      let _rhs60 = _dafny.Seq.Concat(_module.__default.fm3(globalState), _359_v9);
      let _rhs61 = _360_v10;
      let _rhs62 = _dafny.Seq.IsProperPrefixOf(_module.__default.fm4(globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(857)), ((_363_v9) => function (_364_i1) {
        return (_363_v9);
      })(_359_v9)));
      let _rhs63 = p1;
      let _lhs27 = _362_v12;
      let _lhs28 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_362_v12).length));
      let _lhs29 = globalState;
      let _lhs30 = globalState;
      _359_v9 = _rhs60;
      _lhs27[_lhs28] = _rhs61;
      _lhs29.f1 = _rhs62;
      _lhs30.f1 = _rhs63;
      _357_v7 = _357_v7;
      let _365_v13;
      _365_v13 = _dafny.Map.Empty.slice().updateUnsafe(_357_v7,_357_v7);
      let _366_v14;
      _366_v14 = _dafny.Map.Empty.slice().updateUnsafe((p0).isEqualTo((((_365_v13).contains(_357_v7)) ? ((_365_v13).get(_357_v7)) : (p0))),new BigNumber(-600));
      let _rhs64 = ((p0).plus((_dafny.ZERO).minus(new BigNumber((_365_v13).length)))).multipliedBy(p0);
      let _rhs65 = (((_366_v14).contains(p1)) ? ((_366_v14).get(p1)) : ((_dafny.ZERO).minus(_357_v7)));
      _357_v7 = _rhs64;
      _357_v7 = _rhs65;
      return;
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
