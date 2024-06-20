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
    static fm0(globalState) {
      return _module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-120),false)).length), (new BigNumber(-352)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(873)), function (_0_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length)));
    };
    static fm1(p0, p1, globalState) {
      return _dafny.Seq.contains(((false) ? (_dafny.Seq.UnicodeFromString("ewqpfrw")) : (_dafny.Seq.UnicodeFromString("pfb"))), new _dafny.CodePoint('k'.codePointAt(0)));
    };
    static fm5(p0, p1, globalState) {
      return _module.D1.create_DC2(!(!(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)).isEqualTo(new BigNumber((_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(175)), function (_1_i0) {
  return new _dafny.CodePoint('c'.codePointAt(0));
}), _dafny.Seq.UnicodeFromString("ol"))).length))));
    };
    static fm6(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("k"), _dafny.Seq.UnicodeFromString("pgy"));
    };
    static fm7(p0, p1, globalState) {
      return new _dafny.CodePoint('k'.codePointAt(0));
    };
    static fm8(p0, p1, globalState) {
      return _dafny.Set.fromElements(new BigNumber(560), _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("mng")).length)), new BigNumber(987)), ((true) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)) : (new BigNumber(-276))), new BigNumber(203));
    };
    static fm9(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.fromElements((_dafny.Set.fromElements(!(false))).contains(true), !(true) || (true), true);
    };
    static fm10(globalState) {
      if (true) {
        return (_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(false, true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(true, true, true, false, false)));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.of(false));
      }
    };
    static fm11(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(((false) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(636)), function (_2_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("fqfptov")).length);
      })) : (_dafny.Seq.of(new BigNumber(-151)))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(890),false)).length), new BigNumber(505)), _dafny.Seq.of(new BigNumber(768))));
    };
    static fm12(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D5.create_DC13(new BigNumber(879), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), !(false));
      if (_source0.is_DC13) {
        let _3___mcc_h0 = (_source0).cf19;
        let _4___mcc_h1 = (_source0).cf20;
        let _5___mcc_h2 = (_source0).cf21;
        let _6_cf21 = _5___mcc_h2;
        let _7_cf20 = _4___mcc_h1;
        let _8_cf19 = _3___mcc_h0;
        return _module.D5.create_DC12(_dafny.Seq.UnicodeFromString("msdl"));
      } else {
        let _9___mcc_h3 = (_source0).cf18;
        let _10_cf18 = _9___mcc_h3;
        if (true) {
          return _module.D5.create_DC12(_10_cf18);
        } else {
          return _module.D5.create_DC12(_10_cf18);
        }
      }
    };
    static fm13(globalState) {
      let _source1 = _module.D5.create_DC12(_dafny.Seq.Create(_module.__default.abs(new BigNumber(832)), function (_11_i0) {
  return new _dafny.CodePoint('i'.codePointAt(0));
}));
      if (_source1.is_DC13) {
        let _12___mcc_h0 = (_source1).cf19;
        let _13___mcc_h1 = (_source1).cf20;
        let _14___mcc_h2 = (_source1).cf21;
        let _15_cf21 = _14___mcc_h2;
        let _16_cf20 = _13___mcc_h1;
        let _17_cf19 = _12___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_16_cf20,_15_cf21),_17_cf19)).Merge(function () {
          let _coll0 = new _dafny.Map();
          for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_15_cf21)).length))).length),_15_cf21),_15_cf21)).Keys.Elements) {
            let _18_v0 = _compr_0;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_15_cf21)).length))).length),_15_cf21),_15_cf21)).contains(_18_v0)) {
              _coll0.push([_18_v0,_16_cf20]);
            }
          }
          return _coll0;
        }());
      } else {
        let _19___mcc_h3 = (_source1).cf18;
        let _20_cf18 = _19___mcc_h3;
        return (function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll2 = new _dafny.Map();
            for (const _compr_2 of _dafny.IntegerRange(new BigNumber(857), new BigNumber(-21))) {
              let _21_v2 = _compr_2;
              if (((new BigNumber(857)).isLessThanOrEqualTo(_21_v2)) && ((_21_v2).isLessThan(new BigNumber(-21)))) {
                _coll2.push([(_21_v2).multipliedBy(new BigNumber(290)),false]);
              }
            }
            return _coll2;
          }(),true)).Keys.Elements) {
            let _22_v1 = _compr_1;
            if ((_dafny.Map.Empty.slice().updateUnsafe(function () {
              let _coll3 = new _dafny.Map();
              for (const _compr_3 of _dafny.IntegerRange(new BigNumber(857), new BigNumber(-21))) {
                let _21_v2 = _compr_3;
                if (((new BigNumber(857)).isLessThanOrEqualTo(_21_v2)) && ((_21_v2).isLessThan(new BigNumber(-21)))) {
                  _coll3.push([(_21_v2).multipliedBy(new BigNumber(290)),false]);
                }
              }
              return _coll3;
            }(),true)).contains(_22_v1)) {
              _coll1.push([_22_v1,new BigNumber(542)]);
            }
          }
          return _coll1;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),new BigNumber(327))).length),!(true)),new BigNumber(610)));
      }
    };
    static fm14(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(true, false, (new BigNumber((_dafny.Set.fromElements(new BigNumber(76), new BigNumber(112))).length)).isLessThan(new BigNumber((_dafny.Seq.of(new BigNumber(322), new BigNumber((_dafny.Seq.UnicodeFromString("t")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(558)), function (_23_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length))).length))).length)));
    };
    static fm15(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(508), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('o'.codePointAt(0)))).length))).length)))).cardinality())),_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true, true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(644)), function (_24_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false)).length));
      }),_dafny.MultiSet.FromArray(_dafny.Seq.of(false))));
    };
    static fm16(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,!(!(!(false))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false)));
    };
    static fm17(p0, p1, globalState) {
      return _module.D5.create_DC13(new BigNumber((_dafny.Seq.UnicodeFromString("qkgiwheqr")).length), (new BigNumber(303)).multipliedBy(new BigNumber(832)), !(true) || (true));
    };
    static fm18(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,!(true)), _dafny.Map.Empty.slice().updateUnsafe(true,true))).equals(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(false,!(false)))),new _dafny.CodePoint('u'.codePointAt(0)));
    };
    static m0(p0, p1, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _25_v0;
      _25_v0 = _dafny.Set.fromElements(p0);
      let _26_v1;
      let _nw0 = Array((new BigNumber(21)).toNumber()).fill(false);
      _26_v1 = _nw0;
      let _27_v2;
      let _init0 = function (_28_i0) {
        return _dafny.Seq.of(new BigNumber(468));
      };
      let _nw1 = Array((new BigNumber(16)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
        _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _27_v2 = _nw1;
      let _29_v3;
      _29_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _30_v4;
      _30_v4 = _dafny.Seq.UnicodeFromString("knce");
      let _31_v5;
      _31_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,_30_v4);
      let _32_v6;
      let _nw2 = Array((new BigNumber(22)).toNumber());
      _nw2[(_dafny.ZERO).toNumber()] = false;
      _nw2[(_dafny.ONE).toNumber()] = (_module.D9.create_DC26(p0, _25_v0, _26_v1, _27_v2)).dtor_cf51;
      _nw2[(new BigNumber(2)).toNumber()] = p0;
      _nw2[(new BigNumber(3)).toNumber()] = p0;
      _nw2[(new BigNumber(4)).toNumber()] = (((_29_v3).contains(_module.__default.fm1(true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(176)), function (_34_i1) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      }), globalState))) ? ((_29_v3).get(_module.__default.fm1(true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(176)), function (_33_i1) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      }), globalState))) : (!(p0)));
      _nw2[(new BigNumber(5)).toNumber()] = p0;
      _nw2[(new BigNumber(6)).toNumber()] = p0;
      _nw2[(new BigNumber(7)).toNumber()] = p0;
      _nw2[(new BigNumber(8)).toNumber()] = p0;
      _nw2[(new BigNumber(9)).toNumber()] = p0;
      _nw2[(new BigNumber(10)).toNumber()] = p0;
      _nw2[(new BigNumber(11)).toNumber()] = _module.__default.fm1(p0, (((_31_v5).contains(p0)) ? ((_31_v5).get(p0)) : (_30_v4)), globalState);
      _nw2[(new BigNumber(12)).toNumber()] = p0;
      _nw2[(new BigNumber(13)).toNumber()] = p0;
      _nw2[(new BigNumber(14)).toNumber()] = p0;
      _nw2[(new BigNumber(15)).toNumber()] = false;
      _nw2[(new BigNumber(16)).toNumber()] = p0;
      _nw2[(new BigNumber(17)).toNumber()] = p0;
      _nw2[(new BigNumber(18)).toNumber()] = p0;
      _nw2[(new BigNumber(19)).toNumber()] = p0;
      _nw2[(new BigNumber(20)).toNumber()] = p0;
      _nw2[(new BigNumber(21)).toNumber()] = p0;
      _32_v6 = _nw2;
      let _35_v7;
      _35_v7 = _dafny.Seq.of(_32_v6, _32_v6);
      let _36_v8;
      _36_v8 = _dafny.ZERO;
      let _37_v9;
      _37_v9 = _dafny.Map.Empty.slice().updateUnsafe(!(((p0) ? (p0) : (p0))),_dafny.Seq.update(_35_v7, _module.__default.safeIndex(_36_v8, new BigNumber((_35_v7).length)), _32_v6));
      _37_v9 = (_37_v9).update(!(p0), _35_v7);
      let _38_v10;
      _38_v10 = _dafny.Seq.of(_36_v8);
      let _39_v11;
      _39_v11 = _module.D5.create_DC13(_36_v8, (_dafny.ZERO).minus(_36_v8), p0);
      let _pat_let_tv0 = _39_v11;
      let _pat_let_tv1 = p0;
      let _pat_let_tv2 = p0;
      (globalState).f2 = new BigNumber((function (_source2) {
        if (_source2.is_DC5) {
          let _40___mcc_h0 = (_source2).cf5;
          let _41___mcc_h1 = (_source2).cf6;
          let _42___mcc_h2 = (_source2).cf7;
          let _43___mcc_h3 = (_source2).cf8;
          let _44_cf8 = _43___mcc_h3;
          let _45_cf7 = _42___mcc_h2;
          let _46_cf6 = _41___mcc_h1;
          let _47_cf5 = _40___mcc_h0;
          return _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_47_cf5,new _dafny.CodePoint('a'.codePointAt(0))));
        } else if (_source2.is_DC6) {
          let _48___mcc_h4 = (_source2).cf9;
          let _49___mcc_h5 = (_source2).cf10;
          let _50___mcc_h6 = (_source2).cf11;
          let _51___mcc_h7 = (_source2).cf12;
          let _52_cf12 = _51___mcc_h7;
          let _53_cf11 = _50___mcc_h6;
          let _54_cf10 = _49___mcc_h5;
          let _55_cf9 = _48___mcc_h4;
          return (function () {
            let _coll4 = new _dafny.Set();
            for (const _compr_4 of (function () {
              let _coll5 = new _dafny.Map();
              for (const _compr_5 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_54_cf10,new _dafny.CodePoint('e'.codePointAt(0))),_52_cf12)).Keys.Elements) {
                let _56_v12 = _compr_5;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_54_cf10,new _dafny.CodePoint('e'.codePointAt(0))),_52_cf12)).contains(_56_v12)) {
                  _coll5.push([_56_v12,new _dafny.CodePoint('t'.codePointAt(0))]);
                }
              }
              return _coll5;
            }()).Keys.Elements) {
              let _57_v13 = _compr_4;
              if ((function () {
                let _coll6 = new _dafny.Map();
                for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_54_cf10,new _dafny.CodePoint('e'.codePointAt(0))),_52_cf12)).Keys.Elements) {
                  let _56_v12 = _compr_6;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_54_cf10,new _dafny.CodePoint('e'.codePointAt(0))),_52_cf12)).contains(_56_v12)) {
                    _coll6.push([_56_v12,new _dafny.CodePoint('t'.codePointAt(0))]);
                  }
                }
                return _coll6;
              }()).contains(_57_v13)) {
                _coll4.add(_57_v13);
              }
            }
            return _coll4;
          }()).Intersect(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('k'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(_54_cf10,new _dafny.CodePoint('h'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv0).dtor_cf21,new _dafny.CodePoint('f'.codePointAt(0)))));
        } else if (_source2.is_DC4) {
          let _58___mcc_h8 = (_source2).cf4;
          let _59_cf4 = _58___mcc_h8;
          return _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv1,new _dafny.CodePoint('a'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('b'.codePointAt(0))));
        } else {
          let _60___mcc_h9 = (_source2).cf13;
          let _61_cf13 = _60___mcc_h9;
          return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv2,new _dafny.CodePoint('x'.codePointAt(0))))).Difference(function () {
            let _coll7 = new _dafny.Set();
            for (const _compr_7 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(975)), function (_62_i2) {
              return _dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('h'.codePointAt(0)));
            })).Elements) {
              let _63_v14 = _compr_7;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(975)), function (_62_i2) {
                return _dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('h'.codePointAt(0)));
              }), _63_v14)) {
                _coll7.add(_63_v14);
              }
            }
            return _coll7;
          }());
        }
      }(_module.D2.create_DC6(new BigNumber((_38_v10).length), (_39_v11).dtor_cf21, p0, _36_v8))).length);
      let _64_v15;
      _64_v15 = _36_v8;
      let _65_v16;
      let _nw3 = new _module.C3();
      _nw3.__ctor(_30_v4, _64_v15, new BigNumber(-140));
      _65_v16 = _nw3;
      let _66_v17;
      let _nw4 = Array((new BigNumber(12)).toNumber());
      _nw4[(_dafny.ZERO).toNumber()] = _65_v16;
      _nw4[(_dafny.ONE).toNumber()] = _65_v16;
      _nw4[(new BigNumber(2)).toNumber()] = _65_v16;
      _nw4[(new BigNumber(3)).toNumber()] = _65_v16;
      _nw4[(new BigNumber(4)).toNumber()] = _65_v16;
      _nw4[(new BigNumber(5)).toNumber()] = _65_v16;
      _nw4[(new BigNumber(6)).toNumber()] = _65_v16;
      _nw4[(new BigNumber(7)).toNumber()] = _65_v16;
      _nw4[(new BigNumber(8)).toNumber()] = _65_v16;
      _nw4[(new BigNumber(9)).toNumber()] = _65_v16;
      _nw4[(new BigNumber(10)).toNumber()] = _65_v16;
      _nw4[(new BigNumber(11)).toNumber()] = _65_v16;
      _66_v17 = _nw4;
      let _index0 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_66_v17).length));
      (_66_v17)[_index0] = _65_v16;
      let _index1 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_66_v17).length));
      let _rhs0 = (_36_v8).minus((new BigNumber(((_65_v16).f19).length)).plus((_dafny.ZERO).minus(_36_v8)));
      let _rhs1 = _65_v16;
      let _lhs0 = globalState;
      let _lhs1 = _66_v17;
      let _lhs2 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_66_v17).length));
      _lhs0.f2 = _rhs0;
      _lhs1[_lhs2] = _rhs1;
      let _hi0 = _36_v8;
      for (let _67_i3 = _36_v8; _67_i3.isLessThan(_hi0); _67_i3 = _67_i3.plus(_dafny.ONE)) {
        if (p0) {
          let _68_v18;
          let _nw5 = new _module.C2();
          _nw5.__ctor();
          _68_v18 = _nw5;
          let _69_v19;
          _69_v19 = _dafny.MultiSet.fromElements(p0);
          let _70_v20;
          _70_v20 = _dafny.Map.Empty.slice().updateUnsafe(_69_v19,_39_v11);
          _70_v20 = (_70_v20).update(_69_v19, _module.D5.create_DC13(_67_i3, new BigNumber((_38_v10).length), p0));
          (globalState).f9 = _67_i3;
          (globalState).f0 = false;
          let _index2 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_32_v6).length));
          (_32_v6)[_index2] = p0;
          let _index3 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_32_v6).length));
          (_32_v6)[_index3] = !_dafny.areEqual(_dafny.Seq.UnicodeFromString("ph"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(602)), function (_71_i4) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          }));
        } else {
          let _72_v21;
          let _nw6 = new _module.C0();
          _nw6.__ctor(_64_v15, new BigNumber(625));
          _72_v21 = _nw6;
          let _73_v22;
          _73_v22 = _dafny.Seq.of(_72_v21, _72_v21, _72_v21);
          let _74_v23;
          _74_v23 = _dafny.Seq.of(_dafny.Seq.update(_73_v22, _module.__default.safeIndex((_72_v21).f18, new BigNumber((_73_v22).length)), _72_v21));
          let _75_v24;
          let _nw7 = Array((new BigNumber(14)).toNumber());
          _nw7[(_dafny.ZERO).toNumber()] = _74_v23;
          _nw7[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_73_v22, _73_v22);
          _nw7[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_74_v23, _74_v23);
          _nw7[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_74_v23, _74_v23);
          _nw7[(new BigNumber(4)).toNumber()] = _74_v23;
          _nw7[(new BigNumber(5)).toNumber()] = ((p0) ? (_74_v23) : (_74_v23));
          _nw7[(new BigNumber(6)).toNumber()] = _74_v23;
          _nw7[(new BigNumber(7)).toNumber()] = _74_v23;
          _nw7[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_74_v23, _74_v23);
          _nw7[(new BigNumber(9)).toNumber()] = _74_v23;
          _nw7[(new BigNumber(10)).toNumber()] = _74_v23;
          _nw7[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_74_v23, _dafny.Seq.of(_73_v22));
          _nw7[(new BigNumber(12)).toNumber()] = _74_v23;
          _nw7[(new BigNumber(13)).toNumber()] = _74_v23;
          _75_v24 = _nw7;
          let _index4 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_75_v24).length));
          (_75_v24)[_index4] = _74_v23;
          let _index5 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_75_v24).length));
          (_75_v24)[_index5] = _dafny.Seq.update(_74_v23, _module.__default.safeIndex(_67_i3, new BigNumber((_74_v23).length)), _dafny.Seq.Concat(_73_v22, _73_v22));
          let _76_v25;
          _76_v25 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_38_v10, _module.__default.fm11(_67_i3, p0, (_72_v21).f18, globalState)),_36_v8);
          let _77_v26;
          _77_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,_36_v8);
          let _78_v27;
          _78_v27 = _dafny.Seq.of(_77_v26);
          let _79_v28;
          _79_v28 = _dafny.Set.fromElements(_36_v8);
          let _80_v29;
          _80_v29 = _dafny.Map.Empty.slice().updateUnsafe(_36_v8,(_78_v27)[_module.__default.safeIndex(new BigNumber((_79_v28).length), new BigNumber((_78_v27).length))]);
          let _81_v30;
          _81_v30 = _dafny.MultiSet.fromElements(new BigNumber(503));
          _76_v25 = (_76_v25).update(_38_v10, ((_dafny.ZERO).minus(new BigNumber(((((_80_v29).contains(new BigNumber(((_65_v16).f19).length))) ? ((_80_v29).get(new BigNumber(((_65_v16).f19).length))) : (_77_v26))).length))).minus(new BigNumber((_81_v30).cardinality())));
          let _82_v31;
          let _init1 = ((_83_p0) => function (_84_i5) {
            return (_84_i5).plus(new BigNumber((_dafny.Seq.of(_83_p0)).length));
          })(p0);
          let _nw8 = Array((new BigNumber(19)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw8.length); _i0_1++) {
            _nw8[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _82_v31 = _nw8;
          let _85_v32;
          _85_v32 = _module.D4.create_DC10(_82_v31);
          _85_v32 = _85_v32;
          let _index6 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_82_v31).length));
          (_82_v31)[_index6] = _36_v8;
          let _index7 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_82_v31).length));
          (_82_v31)[_index7] = _36_v8;
          let _86_v33;
          _86_v33 = _dafny.Seq.of(p0);
          let _87_v34;
          _87_v34 = new _dafny.CodePoint('i'.codePointAt(0));
          let _rhs2 = (_82_v31)[_module.__default.safeIndex(new BigNumber(778), new BigNumber((_82_v31).length))];
          let _rhs3 = _dafny.Seq.of(_dafny.Seq.contains((_65_v16).f19, _87_v34), p0);
          let _rhs4 = _module.__default.safeDivisionInt((_82_v31)[_module.__default.safeIndex(new BigNumber(778), new BigNumber((_82_v31).length))], (_72_v21).f18);
          let _rhs5 = p0;
          let _lhs3 = globalState;
          let _lhs4 = globalState;
          let _lhs5 = globalState;
          _lhs3.f2 = _rhs2;
          _86_v33 = _rhs3;
          _lhs4.f9 = _rhs4;
          _lhs5.f3 = _rhs5;
        }
        let _88_v35;
        let _nw9 = new _module.C2();
        _nw9.__ctor();
        _88_v35 = _nw9;
        (globalState).f2 = _36_v8;
        let _89_v36;
        _89_v36 = new _dafny.CodePoint('n'.codePointAt(0));
        let _90_v37;
        let _nw10 = new _module.C1();
        _nw10.__ctor(_89_v36);
        _90_v37 = _nw10;
        let _91_v38;
        _91_v38 = _dafny.Seq.of(_90_v37, _90_v37, _90_v37);
        let _rhs6 = _91_v38;
        let _rhs7 = (_67_i3).isLessThan(((_38_v10)[_module.__default.safeIndex(_36_v8, new BigNumber((_38_v10).length))]).multipliedBy(_67_i3));
        let _rhs8 = _dafny.Seq.IsProperPrefixOf((_65_v16).f19, _dafny.Seq.Concat((_65_v16).f19, (_65_v16).f19));
        let _rhs9 = _67_i3;
        let _lhs6 = globalState;
        let _lhs7 = globalState;
        let _lhs8 = globalState;
        _91_v38 = _rhs6;
        _lhs6.f0 = _rhs7;
        _lhs7.f10 = _rhs8;
        _lhs8.f9 = _rhs9;
      }
      let _92_v39;
      _92_v39 = _dafny.MultiSet.fromElements(p0, p0, p0);
      if ((_92_v39).IsProperSubsetOf(_dafny.MultiSet.fromElements(p0, p0))) {
        let _93_v40;
        let _init2 = ((_94_v8) => function (_95_i6) {
          return (_95_i6).plus(_94_v8);
        })(_36_v8);
        let _nw11 = Array((new BigNumber(25)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw11.length); _i0_2++) {
          _nw11[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _93_v40 = _nw11;
        _93_v40 = _93_v40;
        let _96_v41;
        _96_v41 = _dafny.Map.Empty.slice().updateUnsafe(_36_v8,(_dafny.ZERO).minus(_36_v8));
        _96_v41 = (_96_v41).update(_36_v8, _36_v8);
        (globalState).f9 = ((new BigNumber(-709)).multipliedBy(_36_v8)).minus((_36_v8).multipliedBy(_36_v8));
        (globalState).f9 = _36_v8;
        let _97_v42;
        _97_v42 = _module.D6.create_DC16(_36_v8, p0, _36_v8, _36_v8, p0);
        (globalState).f9 = ((_97_v42).dtor_cf28).plus(new BigNumber(-774));
      } else {
        let _98_v43;
        let _init3 = ((_99_v8) => function (_100_i7) {
          return (_100_i7).minus(_99_v8);
        })(_36_v8);
        let _nw12 = Array((new BigNumber(2)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw12.length); _i0_3++) {
          _nw12[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _98_v43 = _nw12;
        let _nw13 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _98_v43 = _nw13;
        let _101_v44;
        _101_v44 = _dafny.MultiSet.fromElements(_36_v8, _module.__default.fm0(globalState));
        let _102_v45;
        _102_v45 = _dafny.Set.fromElements(_36_v8);
        let _103_v46;
        let _nw14 = new _module.C0();
        _nw14.__ctor(_36_v8, new BigNumber((_102_v45).length));
        _103_v46 = _nw14;
        if ((!(!(!(p0)))) && (!(p0))) {
          let _104_v47;
          let _nw15 = new _module.C0();
          _nw15.__ctor(_64_v15, new BigNumber(178));
          _104_v47 = _nw15;
          let _index8 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_98_v43).length));
          (_98_v43)[_index8] = ((p0) ? (_36_v8) : (_36_v8));
          let _105_v48;
          _105_v48 = _dafny.Map.Empty.slice().updateUnsafe(p0,_29_v3);
          let _index9 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_98_v43).length));
          let _rhs10 = _103_v46;
          let _rhs11 = true;
          let _rhs12 = (_dafny.ZERO).minus((new BigNumber((_105_v48).length)).plus(_36_v8));
          let _lhs9 = globalState;
          let _lhs10 = _98_v43;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_98_v43).length));
          _104_v47 = _rhs10;
          _lhs9.f0 = _rhs11;
          _lhs10[_lhs11] = _rhs12;
          let _106_v49;
          let _nw16 = new _module.C0();
          _nw16.__ctor(_64_v15, (_98_v43)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_98_v43).length))]);
          _106_v49 = _nw16;
          (globalState).f3 = (((_29_v3).contains(p0)) ? ((_29_v3).get(p0)) : (false));
          (globalState).f2 = (_98_v43)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_98_v43).length))];
          let _107_v50;
          _107_v50 = _dafny.Seq.of(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("vckfvwohs"), (_65_v16).f19));
          _107_v50 = _107_v50;
        } else {
          (globalState).f2 = (_dafny.ZERO).minus(((_36_v8).multipliedBy(_36_v8)).minus(_module.__default.safeDivisionInt(_36_v8, _36_v8)));
          let _index10 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_26_v1).length));
          (_26_v1)[_index10] = p0;
          let _index11 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_26_v1).length));
          (_26_v1)[_index11] = p0;
          let _108_v51;
          _108_v51 = _module.D11.create_DC32(_36_v8);
          let _109_v52;
          _109_v52 = new _dafny.CodePoint('q'.codePointAt(0));
          let _110_v53;
          let _nw17 = new _module.C1();
          _nw17.__ctor(_109_v52);
          _110_v53 = _nw17;
          let _rhs13 = _module.__default.safeDivisionInt(_36_v8, _36_v8);
          let _rhs14 = (_26_v1)[_module.__default.safeIndex(new BigNumber(665), new BigNumber((_26_v1).length))];
          let _rhs15 = _108_v51;
          let _rhs16 = _110_v53;
          let _lhs12 = globalState;
          _36_v8 = _rhs13;
          _lhs12.f0 = _rhs14;
          _108_v51 = _rhs15;
          _110_v53 = _rhs16;
          let _111_v54;
          _111_v54 = _dafny.Map.Empty.slice().updateUnsafe(_110_v53,_103_v46);
          let _112_v55;
          _112_v55 = _dafny.Seq.of(_103_v46);
          let _113_v56;
          let _nw18 = Array((new BigNumber(10)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = ((p0) ? (_103_v46) : (_103_v46));
          _nw18[(_dafny.ONE).toNumber()] = (((_111_v54).contains(_110_v53)) ? ((_111_v54).get(_110_v53)) : (_103_v46));
          _nw18[(new BigNumber(2)).toNumber()] = _103_v46;
          _nw18[(new BigNumber(3)).toNumber()] = _103_v46;
          _nw18[(new BigNumber(4)).toNumber()] = _103_v46;
          _nw18[(new BigNumber(5)).toNumber()] = _103_v46;
          _nw18[(new BigNumber(6)).toNumber()] = (_112_v55)[_module.__default.safeIndex(new BigNumber((_38_v10).length), new BigNumber((_112_v55).length))];
          _nw18[(new BigNumber(7)).toNumber()] = _103_v46;
          _nw18[(new BigNumber(8)).toNumber()] = _103_v46;
          _nw18[(new BigNumber(9)).toNumber()] = _103_v46;
          _113_v56 = _nw18;
          let _index12 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_113_v56).length));
          (_113_v56)[_index12] = _103_v46;
          let _114_v57;
          _114_v57 = _module.D2.create_DC6(_36_v8, true, p0, _36_v8);
          let _115_v58;
          _115_v58 = _dafny.Seq.of(_114_v57);
          let _116_v59;
          _116_v59 = _dafny.MultiSet.fromElements(_39_v11, _39_v11, _module.__default.fm17(new BigNumber((_30_v4).length), _115_v58, globalState));
          let _117_v60;
          _117_v60 = _module.D1.create_DC3(_module.D1.create_DC2((_26_v1)[_module.__default.safeIndex(new BigNumber(665), new BigNumber((_26_v1).length))]));
          let _118_v61;
          _118_v61 = _module.D6.create_DC15(_36_v8, _36_v8, _module.__default.fm11((_dafny.ZERO).minus(_36_v8), false, new BigNumber(218), globalState), _36_v8, _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("c")).length))));
          let _pat_let_tv3 = p0;
          let _index13 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_113_v56).length));
          let _rhs17 = _103_v46;
          let _rhs18 = (_116_v59).Union((_116_v59).Difference(_116_v59));
          let _rhs19 = new BigNumber((_module.__default.fm18(_36_v8, globalState)).length);
          let _rhs20 = (_dafny.ZERO).minus((_118_v61).dtor_cf24);
          let _rhs21 = function (_pat_let0_0) {
            return function (_119_dt__update__tmp_h3) {
              return function (_pat_let1_0) {
                return function (_120_dt__update_hcf3_h0) {
                  return _module.D1.create_DC3(_120_dt__update_hcf3_h0);
                }(_pat_let1_0);
              }(_module.D1.create_DC2(_pat_let_tv3));
            }(_pat_let0_0);
          }(_117_v60);
          let _lhs13 = _113_v56;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_113_v56).length));
          let _lhs15 = globalState;
          let _lhs16 = globalState;
          _lhs13[_lhs14] = _rhs17;
          _116_v59 = _rhs18;
          _lhs15.f9 = _rhs19;
          _lhs16.f9 = _rhs20;
          _117_v60 = _rhs21;
          let _121_v62;
          _121_v62 = _dafny.Map.Empty.slice().updateUnsafe(_109_v52,!(!(p0)));
          _121_v62 = (_121_v62).update((_110_v53).f20, p0);
        }
        (globalState).f10 = p0;
        let _122_v63;
        let _nw19 = new _module.C2();
        _nw19.__ctor();
        _122_v63 = _nw19;
        let _123_v64;
        _123_v64 = _dafny.MultiSet.fromElements(_122_v63, _122_v63, _122_v63, _122_v63, _122_v63);
        _123_v64 = _123_v64;
      }
      let _124_v65;
      _124_v65 = new _dafny.CodePoint('j'.codePointAt(0));
      let _125_v66;
      let _nw20 = new _module.C0();
      _nw20.__ctor(_64_v15, _36_v8);
      _125_v66 = _nw20;
      let _index14 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_32_v6).length));
      (_32_v6)[_index14] = p0;
      let _126_v67;
      let _nw21 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
      _126_v67 = _nw21;
      let _127_v69;
      _127_v69 = _dafny.Seq.of(function () {
        let _coll8 = new _dafny.Set();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(784), new BigNumber(243))) {
          let _128_v68 = _compr_8;
          if (((new BigNumber(784)).isLessThanOrEqualTo(_128_v68)) && ((_128_v68).isLessThan(new BigNumber(243)))) {
            _coll8.add(_module.__default.safeDivisionInt(_128_v68, _36_v8));
          }
        }
        return _coll8;
      }());
      let _index15 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_126_v67).length));
      (_126_v67)[_index15] = (_127_v69)[_module.__default.safeIndex(new BigNumber(77), new BigNumber((_127_v69).length))];
      let _129_v70;
      _129_v70 = _dafny.Map.Empty.slice().updateUnsafe(p0,_36_v8);
      let _130_v71;
      _130_v71 = _dafny.MultiSet.fromElements(_125_v66);
      let _131_v72;
      let _nw22 = new _module.C2();
      _nw22.__ctor();
      _131_v72 = _nw22;
      let _132_v73;
      _132_v73 = _dafny.Map.Empty.slice().updateUnsafe(p0,_131_v72);
      let _133_v74;
      _133_v74 = _dafny.Set.fromElements(_36_v8, new BigNumber((_132_v73).length), _module.__default.fm0(globalState));
      let _index16 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_32_v6).length));
      let _index17 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_126_v67).length));
      let _rhs22 = _30_v4;
      let _rhs23 = ((_65_v16).f19)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_129_v70).length), new BigNumber((_130_v71).cardinality())), new BigNumber(((_65_v16).f19).length))];
      let _rhs24 = ((_module.__default.fm1(p0, _30_v4, globalState)) ? (_125_v66) : (_125_v66));
      let _rhs25 = p0;
      let _rhs26 = ((_133_v74).Intersect(_133_v74)).Intersect((_133_v74).Intersect(_133_v74));
      let _lhs17 = _32_v6;
      let _lhs18 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_32_v6).length));
      let _lhs19 = _126_v67;
      let _lhs20 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_126_v67).length));
      _30_v4 = _rhs22;
      _124_v65 = _rhs23;
      _125_v66 = _rhs24;
      _lhs17[_lhs18] = _rhs25;
      _lhs19[_lhs20] = _rhs26;
      r0 = ((_dafny.Seq.contains(_30_v4, new _dafny.CodePoint('g'.codePointAt(0)))) ? (_36_v8) : (_36_v8));
      let _134_v75;
      _134_v75 = _dafny.Map.Empty.slice().updateUnsafe((_32_v6)[_module.__default.safeIndex(new BigNumber(684), new BigNumber((_32_v6).length))],_124_v65);
      let _135_v76;
      _135_v76 = _134_v75;
      let _136_v77;
      _136_v77 = _dafny.Map.Empty.slice().updateUnsafe(_36_v8,(_134_v75));
      r1 = new BigNumber(((_136_v77).update(_module.__default.safeModuloInt(_36_v8, _36_v8), _dafny.Map.Empty.slice().updateUnsafe(p0,_124_v65))).length);
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _137_v0;
      _137_v0 = true;
      let _138_v1;
      let _init4 = ((_139_v0) => function (_140_i0) {
        return _139_v0;
      })(_137_v0);
      let _nw23 = Array((new BigNumber(21)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw23.length); _i0_4++) {
        _nw23[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _138_v1 = _nw23;
      let _141_v2;
      _141_v2 = new BigNumber(213);
      let _142_v3;
      _142_v3 = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,(_dafny.ZERO).minus(_141_v2));
      let _143_v4;
      _143_v4 = _dafny.Seq.UnicodeFromString("l");
      let _144_v5;
      _144_v5 = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_137_v0);
      let _145_globalState;
      let _nw24 = new _module.GlobalState();
      _nw24.__ctor(true, _dafny.Seq.of(_137_v0), new BigNumber(172), false, true, new _dafny.CodePoint('f'.codePointAt(0)), new BigNumber(669), _138_v1, _dafny.MultiSet.fromElements(_141_v2, _141_v2), new BigNumber(257), false, ((_137_v0) ? (_138_v1) : (_138_v1)), _142_v3, _143_v4, false, (_144_v5).Merge(_144_v5), true);
      _145_globalState = _nw24;
      if (_137_v0) {
        _142_v3 = (_142_v3).update(_137_v0, _141_v2);
        let _146_v6;
        let _init5 = function (_147_i1) {
          return _dafny.Seq.of(false);
        };
        let _nw25 = Array((new BigNumber(12)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw25.length); _i0_5++) {
          _nw25[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _146_v6 = _nw25;
        let _148_v7;
        _148_v7 = _dafny.Seq.of(_137_v0);
        let _index18 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_146_v6).length));
        (_146_v6)[_index18] = _148_v7;
        let _index19 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_146_v6).length));
        (_146_v6)[_index19] = _dafny.Seq.update(_148_v7, _module.__default.safeIndex(((_137_v0) ? ((_dafny.ZERO).minus(_141_v2)) : (_141_v2)), new BigNumber((_148_v7).length)), (_148_v7)[_module.__default.safeIndex((_dafny.ZERO).minus(_141_v2), new BigNumber((_148_v7).length))]);
        let _149_v8;
        _149_v8 = _dafny.Map.Empty.slice().updateUnsafe(_141_v2,_141_v2);
        let _150_v9;
        _150_v9 = _dafny.Map.Empty.slice().updateUnsafe(_149_v8,_137_v0);
        _141_v2 = (_141_v2).minus(new BigNumber((_150_v9).length));
        let _151_v10;
        _151_v10 = _dafny.Set.fromElements(_141_v2);
        let _152_v11;
        _152_v11 = _dafny.Seq.of(new BigNumber(313));
        let _153_v12;
        _153_v12 = _dafny.Seq.of(new BigNumber((_152_v11).length));
        let _154_v13;
        let _nw26 = Array((new BigNumber(28)).toNumber());
        _nw26[(_dafny.ZERO).toNumber()] = _141_v2;
        _nw26[(_dafny.ONE).toNumber()] = new BigNumber(77);
        _nw26[(new BigNumber(2)).toNumber()] = (new BigNumber(((_146_v6)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_146_v6).length))]).length)).multipliedBy(_141_v2);
        _nw26[(new BigNumber(3)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(4)).toNumber()] = (_module.__default.fm0(_145_globalState)).multipliedBy((_dafny.ZERO).minus(_141_v2));
        _nw26[(new BigNumber(5)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(6)).toNumber()] = _module.__default.fm0(_145_globalState);
        _nw26[(new BigNumber(7)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(8)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(355), _141_v2);
        _nw26[(new BigNumber(10)).toNumber()] = (new BigNumber((_151_v10).length)).plus(_141_v2);
        _nw26[(new BigNumber(11)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(12)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(13)).toNumber()] = new BigNumber(336);
        _nw26[(new BigNumber(14)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(15)).toNumber()] = (_module.__default.fm0(_145_globalState)).minus(_141_v2);
        _nw26[(new BigNumber(16)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(17)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(18)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(19)).toNumber()] = _module.__default.safeDivisionInt(_141_v2, _141_v2);
        _nw26[(new BigNumber(20)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(21)).toNumber()] = ((_137_v0) ? (_141_v2) : (_141_v2));
        _nw26[(new BigNumber(22)).toNumber()] = new BigNumber((_143_v4).length);
        _nw26[(new BigNumber(23)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(24)).toNumber()] = _module.__default.safeDivisionInt(_141_v2, new BigNumber((_152_v11).length));
        _nw26[(new BigNumber(25)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(26)).toNumber()] = _141_v2;
        _nw26[(new BigNumber(27)).toNumber()] = _module.__default.fm0(_145_globalState);
        _154_v13 = _nw26;
        let _155_v14;
        _155_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-657),_154_v13);
        let _rhs27 = _138_v1;
        let _rhs28 = ((_141_v2).multipliedBy(_141_v2)).minus(_module.__default.safeDivisionInt(_141_v2, _141_v2));
        let _rhs29 = (((_155_v14).contains(_module.__default.safeModuloInt(_141_v2, _141_v2))) ? ((_155_v14).get(_module.__default.safeModuloInt(_141_v2, _141_v2))) : (_154_v13));
        let _lhs21 = _145_globalState;
        _138_v1 = _rhs27;
        _lhs21.f9 = _rhs28;
        _154_v13 = _rhs29;
        let _156_v15;
        let _nw27 = Array((new BigNumber(12)).toNumber()).fill([]);
        _156_v15 = _nw27;
        _156_v15 = _156_v15;
      } else {
        let _157_v16;
        _157_v16 = new _dafny.CodePoint('r'.codePointAt(0));
        let _158_v17;
        _158_v17 = _dafny.Map.Empty.slice().updateUnsafe(_157_v16,_141_v2);
        (_145_globalState).f2 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_158_v17,_137_v0)).update(_158_v17, _137_v0)).length);
        let _159_v18;
        _159_v18 = _dafny.Seq.of(_137_v0);
        _159_v18 = _dafny.Seq.of(!(_137_v0));
        let _160_v20;
        _160_v20 = _dafny.MultiSet.fromElements(_141_v2, _141_v2, _141_v2, _141_v2, _141_v2);
        let _161_v21;
        _161_v21 = _dafny.MultiSet.fromElements(_137_v0, _137_v0);
        let _162_v22;
        let _163_v23;
        let _out0;
        let _out1;
        let _outcollector0 = _module.__default.m0(_dafny.Seq.IsPrefixOf(_143_v4, _143_v4), (function () {
          let _coll9 = new _dafny.Map();
          for (const _compr_9 of (_160_v20).Elements) {
            let _164_v19 = _compr_9;
            if ((_160_v20).contains(_164_v19)) {
              _coll9.push([_module.__default.safeModuloInt(_164_v19, _141_v2),_137_v0]);
            }
          }
          return _coll9;
        }()).update(new BigNumber((_161_v21).cardinality()), (_159_v18)[_module.__default.safeIndex(_141_v2, new BigNumber((_159_v18).length))]), _145_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _162_v22 = _out0;
        _163_v23 = _out1;
        let _165_v24;
        _165_v24 = _dafny.Seq.of(_143_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(783)), ((_166_v16) => function (_167_i2) {
          return _166_v16;
        })(_157_v16)), _143_v4, _143_v4);
        if (_module.__default.fm1(false, (_165_v24)[_module.__default.safeIndex(_162_v22, new BigNumber((_165_v24).length))], _145_globalState)) {
          _157_v16 = _157_v16;
          (_145_globalState).f3 = (_141_v2).isLessThan(_module.__default.fm0(_145_globalState));
          let _168_v25;
          _168_v25 = _163_v23;
          let _169_v26;
          _169_v26 = _dafny.Map.Empty.slice().updateUnsafe((_168_v25),!(_137_v0));
          let _170_v27;
          let _171_v28;
          let _out2;
          let _out3;
          let _outcollector1 = _module.__default.m0(_137_v0, _169_v26, _145_globalState);
          _out2 = _outcollector1[0];
          _out3 = _outcollector1[1];
          _170_v27 = _out2;
          _171_v28 = _out3;
          (_145_globalState).f9 = _171_v28;
          let _172_v29;
          _172_v29 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_159_v18, _dafny.Seq.of(false)),_137_v0);
          let _173_v30;
          _173_v30 = _dafny.Seq.of(_170_v27);
          let _174_v31;
          _174_v31 = _dafny.MultiSet.fromElements(_173_v30);
          (_145_globalState).f3 = !((((_172_v29).contains(_159_v18)) ? ((_172_v29).get(_159_v18)) : ((_dafny.MultiSet.fromElements(_173_v30, _173_v30)).IsSubsetOf(_174_v31))));
        } else {
          _157_v16 = _157_v16;
          _159_v18 = _dafny.Seq.Concat(_159_v18, _159_v18);
          let _175_v32;
          let _nw28 = new _module.C1();
          _nw28.__ctor(_157_v16);
          _175_v32 = _nw28;
          let _176_v33;
          _176_v33 = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_159_v18);
          let _177_v34;
          _177_v34 = _dafny.Set.fromElements(_159_v18, _159_v18, _159_v18, _159_v18, (((_176_v33).contains((_159_v18)[_module.__default.safeIndex(_141_v2, new BigNumber((_159_v18).length))])) ? ((_176_v33).get((_159_v18)[_module.__default.safeIndex(_141_v2, new BigNumber((_159_v18).length))])) : (_159_v18)));
          (_145_globalState).f10 = (_177_v34).IsProperSubsetOf((_177_v34).Union(_177_v34));
          _142_v3 = (_142_v3).update(((_137_v0) ? (_137_v0) : (_137_v0)), _163_v23);
        }
        let _rhs30 = !(!(true)) || (true);
        let _rhs31 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_159_v18, _dafny.Seq.of(_137_v0, _137_v0, _137_v0)))).IsDisjointFrom(_161_v21);
        let _rhs32 = (_141_v2).isLessThanOrEqualTo(_162_v22);
        let _lhs22 = _145_globalState;
        let _lhs23 = _145_globalState;
        let _lhs24 = _145_globalState;
        _lhs22.f10 = _rhs30;
        _lhs23.f0 = _rhs31;
        _lhs24.f0 = _rhs32;
      }
      let _178_v35;
      _178_v35 = _dafny.MultiSet.fromElements(_141_v2, _141_v2, _141_v2);
      let _179_v36;
      _179_v36 = _module.D6.create_DC16(_141_v2, _137_v0, new BigNumber((_dafny.Seq.UnicodeFromString("b")).length), _141_v2, _137_v0);
      let _rhs33 = (((_178_v35).contains((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_179_v36).dtor_cf31, _141_v2)))) ? ((_178_v35).get((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_179_v36).dtor_cf31, _141_v2)))) : ((_dafny.ZERO).minus(_141_v2)));
      let _rhs34 = _141_v2;
      let _lhs25 = _145_globalState;
      let _lhs26 = _145_globalState;
      _lhs25.f9 = _rhs33;
      _lhs26.f9 = _rhs34;
      let _180_v37;
      _180_v37 = _dafny.Seq.of(new BigNumber(-496), _141_v2, _141_v2);
      _180_v37 = _dafny.Seq.Concat(_dafny.Seq.Concat(_180_v37, _dafny.Seq.of(_141_v2)), _180_v37);
      let _181_v38;
      _181_v38 = _module.D1.create_DC2(false);
      let _182_v39;
      _182_v39 = _dafny.MultiSet.fromElements(_137_v0);
      let _183_v40;
      _183_v40 = _dafny.Set.fromElements(_137_v0, _137_v0);
      let _pat_let_tv4 = _137_v0;
      let _pat_let_tv5 = _137_v0;
      let _pat_let_tv6 = _137_v0;
      let _pat_let_tv7 = _145_globalState;
      let _184_v41;
      let _nw29 = Array((new BigNumber(20)).toNumber());
      _nw29[(_dafny.ZERO).toNumber()] = _181_v38;
      _nw29[(_dafny.ONE).toNumber()] = _181_v38;
      _nw29[(new BigNumber(2)).toNumber()] = function (_pat_let2_0) {
        return function (_185_dt__update__tmp_h0) {
          return function (_pat_let3_0) {
            return function (_186_dt__update_hcf2_h0) {
              return _module.D1.create_DC2(_186_dt__update_hcf2_h0);
            }(_pat_let3_0);
          }(_pat_let_tv4);
        }(_pat_let2_0);
      }(_181_v38);
      _nw29[(new BigNumber(3)).toNumber()] = _181_v38;
      _nw29[(new BigNumber(4)).toNumber()] = _181_v38;
      _nw29[(new BigNumber(5)).toNumber()] = _181_v38;
      _nw29[(new BigNumber(6)).toNumber()] = _module.D1.create_DC2(_137_v0);
      _nw29[(new BigNumber(7)).toNumber()] = _module.D1.create_DC2(_137_v0);
      _nw29[(new BigNumber(8)).toNumber()] = _181_v38;
      _nw29[(new BigNumber(9)).toNumber()] = _181_v38;
      _nw29[(new BigNumber(10)).toNumber()] = _181_v38;
      _nw29[(new BigNumber(11)).toNumber()] = _181_v38;
      _nw29[(new BigNumber(12)).toNumber()] = function (_pat_let4_0) {
        return function (_187_dt__update__tmp_h1) {
          return function (_pat_let5_0) {
            return function (_188_dt__update_hcf2_h1) {
              return _module.D1.create_DC2(_188_dt__update_hcf2_h1);
            }(_pat_let5_0);
          }(_pat_let_tv5);
        }(_pat_let4_0);
      }(_181_v38);
      _nw29[(new BigNumber(13)).toNumber()] = function (_pat_let6_0) {
        return function (_189_dt__update__tmp_h2) {
          return function (_pat_let7_0) {
            return function (_190_dt__update_hcf2_h2) {
              return _module.D1.create_DC2(_190_dt__update_hcf2_h2);
            }(_pat_let7_0);
          }(_module.__default.fm1(_pat_let_tv6, _dafny.Seq.UnicodeFromString("jol"), _pat_let_tv7));
        }(_pat_let6_0);
      }(_181_v38);
      _nw29[(new BigNumber(14)).toNumber()] = _181_v38;
      _nw29[(new BigNumber(15)).toNumber()] = _181_v38;
      _nw29[(new BigNumber(16)).toNumber()] = _181_v38;
      _nw29[(new BigNumber(17)).toNumber()] = _module.__default.fm5(_141_v2, (((_182_v39).contains(false)) ? ((_182_v39).get(false)) : (new BigNumber((_183_v40).length))), _145_globalState);
      _nw29[(new BigNumber(18)).toNumber()] = _181_v38;
      _nw29[(new BigNumber(19)).toNumber()] = _181_v38;
      _184_v41 = _nw29;
      let _191_v42;
      _191_v42 = _module.D2.create_DC4(_184_v41);
      _184_v41 = ((_dafny.Seq.IsProperPrefixOf(_143_v4, _dafny.Seq.UnicodeFromString("e"))) ? (_184_v41) : ((_191_v42).dtor_cf4));
      let _192_v43;
      _192_v43 = new _dafny.CodePoint('y'.codePointAt(0));
      let _193_v44;
      let _nw30 = new _module.C1();
      _nw30.__ctor(_192_v43);
      _193_v44 = _nw30;
      let _194_v45;
      _194_v45 = _module.D6.create_DC14(_193_v44);
      let _pat_let_tv8 = _193_v44;
      let _pat_let_tv9 = _193_v44;
      let _pat_let_tv10 = _137_v0;
      let _source3 = function (_pat_let8_0) {
        return function (_195_dt__update__tmp_h3) {
          return function (_pat_let9_0) {
            return function (_196_dt__update_hcf22_h0) {
              return _module.D6.create_DC14(_196_dt__update_hcf22_h0);
            }(_pat_let9_0);
          }(((_pat_let_tv10) ? (_pat_let_tv8) : (_pat_let_tv9)));
        }(_pat_let8_0);
      }(_194_v45);
      if (_source3.is_DC15) {
        let _197___mcc_h0 = (_source3).cf23;
        let _198___mcc_h1 = (_source3).cf24;
        let _199___mcc_h2 = (_source3).cf25;
        let _200___mcc_h3 = (_source3).cf26;
        let _201___mcc_h4 = (_source3).cf27;
        let _202_cf27 = _201___mcc_h4;
        let _203_cf26 = _200___mcc_h3;
        let _204_cf25 = _199___mcc_h2;
        let _205_cf24 = _198___mcc_h1;
        let _206_cf23 = _197___mcc_h0;
        let _207_v46;
        _207_v46 = _dafny.MultiSet.fromElements((_193_v44).f20, _192_v43, _192_v43);
        let _rhs35 = _206_cf23;
        let _rhs36 = (((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_145_globalState),_206_cf23)).length)).isLessThanOrEqualTo(new BigNumber(760))) ? (new BigNumber(446)) : (((_137_v0) ? (new BigNumber((_207_v46).cardinality())) : (_206_cf23))));
        let _lhs27 = _145_globalState;
        _lhs27.f9 = _rhs35;
        _141_v2 = _rhs36;
        let _208_v47;
        _208_v47 = new BigNumber((_143_v4).length);
        let _209_v48;
        let _nw31 = new _module.C0();
        _nw31.__ctor(_208_v47, _206_cf23);
        _209_v48 = _nw31;
        let _210_v49;
        let _nw32 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
        _210_v49 = _nw32;
        let _211_v50;
        let _nw33 = new _module.C0();
        _nw33.__ctor((_dafny.ZERO).minus(_203_cf26), (_209_v48).f18);
        _211_v50 = _nw33;
        let _212_v51;
        _212_v51 = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_211_v50);
        let _index20 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_210_v49).length));
        (_210_v49)[_index20] = (_212_v51).Merge(_212_v51);
        let _213_v52;
        _213_v52 = _dafny.MultiSet.fromElements(_138_v1);
        let _214_v53;
        _214_v53 = _dafny.Map.Empty.slice().updateUnsafe(_203_cf26,_203_cf26);
        let _215_v54;
        _215_v54 = _dafny.Map.Empty.slice().updateUnsafe(false,_214_v53);
        let _index21 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_210_v49).length));
        let _rhs37 = _209_v48;
        let _rhs38 = _module.__default.fm11(_module.__default.safeModuloInt(new BigNumber((_143_v4).length), new BigNumber((_dafny.MultiSet.FromArray(_180_v37)).cardinality())), _137_v0, (((_213_v52).contains(_138_v1)) ? ((_213_v52).get(_138_v1)) : (new BigNumber(((((_215_v54).contains(_137_v0)) ? ((_215_v54).get(_137_v0)) : (_dafny.Map.Empty.slice().updateUnsafe((_209_v48).f18,_141_v2)))).length))), _145_globalState);
        let _rhs39 = _193_v44;
        let _rhs40 = _212_v51;
        let _lhs28 = _210_v49;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_210_v49).length));
        _209_v48 = _rhs37;
        _204_cf25 = _rhs38;
        _193_v44 = _rhs39;
        _lhs28[_lhs29] = _rhs40;
        let _216_v55;
        let _nw34 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _216_v55 = _nw34;
        _216_v55 = _216_v55;
        (_145_globalState).f0 = _137_v0;
      } else if (_source3.is_DC16) {
        let _217___mcc_h5 = (_source3).cf28;
        let _218___mcc_h6 = (_source3).cf29;
        let _219___mcc_h7 = (_source3).cf30;
        let _220___mcc_h8 = (_source3).cf31;
        let _221___mcc_h9 = (_source3).cf32;
        let _222_cf32 = _221___mcc_h9;
        let _223_cf31 = _220___mcc_h8;
        let _224_cf30 = _219___mcc_h7;
        let _225_cf29 = _218___mcc_h6;
        let _226_cf28 = _217___mcc_h5;
        let _227_v56;
        let _228_v57;
        let _out4;
        let _out5;
        let _outcollector2 = (_193_v44).m3(_137_v0, _145_globalState);
        _out4 = _outcollector2[0];
        _out5 = _outcollector2[1];
        _227_v56 = _out4;
        _228_v57 = _out5;
        let _229_v58;
        _229_v58 = _dafny.Seq.of(_225_cf29);
        let _230_v59;
        _230_v59 = _dafny.Map.Empty.slice().updateUnsafe(_226_cf28,_module.__default.fm11(_224_cf30, _137_v0, _226_cf28, _145_globalState));
        let _231_v60;
        _231_v60 = _dafny.Seq.of(_dafny.Seq.update(_180_v37, _module.__default.safeIndex(new BigNumber(-15), new BigNumber((_180_v37).length)), _227_v56));
        let _232_v61;
        _232_v61 = _dafny.Map.Empty.slice().updateUnsafe(_143_v4,_222_cf32);
        let _233_v62;
        let _nw35 = Array((new BigNumber(13)).toNumber());
        _nw35[(_dafny.ZERO).toNumber()] = _180_v37;
        _nw35[(_dafny.ONE).toNumber()] = _180_v37;
        _nw35[(new BigNumber(2)).toNumber()] = _180_v37;
        _nw35[(new BigNumber(3)).toNumber()] = _180_v37;
        _nw35[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_180_v37, _180_v37);
        _nw35[(new BigNumber(5)).toNumber()] = _180_v37;
        _nw35[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_module.__default.fm0(_145_globalState), new BigNumber(295));
        _nw35[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_223_cf31, new BigNumber(173), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_229_v58, _module.__default.safeIndex(_226_cf28, new BigNumber((_229_v58).length)), _137_v0))).cardinality()), new BigNumber(3));
        _nw35[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_180_v37, _module.__default.safeIndex(_141_v2, new BigNumber((_180_v37).length)), _226_cf28);
        _nw35[(new BigNumber(9)).toNumber()] = (((_230_v59).contains(_226_cf28)) ? ((_230_v59).get(_226_cf28)) : (_180_v37));
        _nw35[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_180_v37, _dafny.Seq.Create(_module.__default.abs(new BigNumber(151)), ((_234_v40) => function (_235_i3) {
          return new BigNumber((_234_v40).length);
        })(_183_v40)));
        _nw35[(new BigNumber(11)).toNumber()] = _180_v37;
        _nw35[(new BigNumber(12)).toNumber()] = _dafny.Seq.update((_231_v60)[_module.__default.safeIndex(_223_cf31, new BigNumber((_231_v60).length))], _module.__default.safeIndex(new BigNumber((_232_v61).length), new BigNumber(((_231_v60)[_module.__default.safeIndex(_223_cf31, new BigNumber((_231_v60).length))]).length)), _224_cf30);
        _233_v62 = _nw35;
        let _index22 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_233_v62).length));
        (_233_v62)[_index22] = _180_v37;
        let _index23 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_233_v62).length));
        (_233_v62)[_index23] = _180_v37;
        let _236_v63;
        let _nw36 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _236_v63 = _nw36;
        let _237_v64;
        _237_v64 = _module.D8.create_DC23(_226_cf28, _236_v63, _223_cf31);
        let _238_v65;
        _238_v65 = _dafny.Set.fromElements(_226_cf28);
        let _239_v66;
        _239_v66 = _dafny.Map.Empty.slice().updateUnsafe(_141_v2,_238_v65);
        let _240_v67;
        _240_v67 = _dafny.Seq.of(_238_v65, (((_239_v66).contains(new BigNumber((_144_v5).length))) ? ((_239_v66).get(new BigNumber((_144_v5).length))) : (_238_v65)));
        let _source4 = _module.__default.fm12((_237_v64).dtor_cf46, _240_v67, _224_cf30, _223_cf31, _145_globalState);
        if (_source4.is_DC13) {
          let _241___mcc_h15 = (_source4).cf19;
          let _242___mcc_h16 = (_source4).cf20;
          let _243___mcc_h17 = (_source4).cf21;
          let _244_cf21 = _243___mcc_h17;
          let _245_cf20 = _242___mcc_h16;
          let _246_cf19 = _241___mcc_h15;
          (_145_globalState).f2 = _245_cf20;
          let _index24 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_236_v63).length));
          (_236_v63)[_index24] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),_142_v3)).length);
          let _247_v68;
          let _nw37 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _247_v68 = _nw37;
          let _index25 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_247_v68).length));
          (_247_v68)[_index25] = (_193_v44).f20;
          let _248_v69;
          _248_v69 = _dafny.Map.Empty.slice().updateUnsafe(_223_cf31,new BigNumber((_229_v58).length));
          let _249_v70;
          _249_v70 = _dafny.Map.Empty.slice().updateUnsafe(_248_v69,_143_v4);
          let _250_v71;
          _250_v71 = _dafny.Map.Empty.slice().updateUnsafe(_245_cf20,_248_v69);
          let _251_v72;
          _251_v72 = _module.D5.create_DC12(_143_v4);
          let _252_v73;
          let _nw38 = Array((new BigNumber(28)).toNumber());
          _nw38[(_dafny.ZERO).toNumber()] = _143_v4;
          _nw38[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sfvx"), _dafny.Seq.UnicodeFromString("tunb"));
          _nw38[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("laijoy");
          _nw38[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("l");
          _nw38[(new BigNumber(4)).toNumber()] = _143_v4;
          _nw38[(new BigNumber(5)).toNumber()] = _143_v4;
          _nw38[(new BigNumber(6)).toNumber()] = (((_249_v70).contains((((_250_v71).contains(new BigNumber(((_233_v62)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_233_v62).length))]).length))) ? ((_250_v71).get(new BigNumber(((_233_v62)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_233_v62).length))]).length))) : (_248_v69)))) ? ((_249_v70).get((((_250_v71).contains(new BigNumber(((_233_v62)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_233_v62).length))]).length))) ? ((_250_v71).get(new BigNumber(((_233_v62)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_233_v62).length))]).length))) : (_248_v69)))) : (_143_v4));
          _nw38[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_143_v4, _module.__default.safeIndex(_141_v2, new BigNumber((_143_v4).length)), (_193_v44).f20);
          _nw38[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_143_v4, _143_v4);
          _nw38[(new BigNumber(9)).toNumber()] = _dafny.Seq.update((_251_v72).dtor_cf18, _module.__default.safeIndex(_227_v56, new BigNumber(((_251_v72).dtor_cf18).length)), _192_v43);
          _nw38[(new BigNumber(10)).toNumber()] = _143_v4;
          _nw38[(new BigNumber(11)).toNumber()] = ((_225_cf29) ? (_143_v4) : (_143_v4));
          _nw38[(new BigNumber(12)).toNumber()] = ((_module.__default.fm1(_225_cf29, _143_v4, _145_globalState)) ? (_dafny.Seq.update(_module.__default.fm6(_222_cf32, _145_globalState), _module.__default.safeIndex(_227_v56, new BigNumber((_module.__default.fm6(_222_cf32, _145_globalState)).length)), (_193_v44).f20)) : (_143_v4));
          _nw38[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("vrxbpbgci");
          _nw38[(new BigNumber(14)).toNumber()] = _143_v4;
          _nw38[(new BigNumber(15)).toNumber()] = _143_v4;
          _nw38[(new BigNumber(16)).toNumber()] = _143_v4;
          _nw38[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("uci");
          _nw38[(new BigNumber(18)).toNumber()] = _143_v4;
          _nw38[(new BigNumber(19)).toNumber()] = _143_v4;
          _nw38[(new BigNumber(20)).toNumber()] = (_module.D5.create_DC12(_143_v4)).dtor_cf18;
          _nw38[(new BigNumber(21)).toNumber()] = _143_v4;
          _nw38[(new BigNumber(22)).toNumber()] = _143_v4;
          _nw38[(new BigNumber(23)).toNumber()] = _143_v4;
          _nw38[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("smf");
          _nw38[(new BigNumber(25)).toNumber()] = _143_v4;
          _nw38[(new BigNumber(26)).toNumber()] = _dafny.Seq.UnicodeFromString("hrnoni");
          _nw38[(new BigNumber(27)).toNumber()] = _dafny.Seq.UnicodeFromString("abfgyma");
          _252_v73 = _nw38;
          let _253_v74;
          _253_v74 = _dafny.Set.fromElements(_229_v58, _dafny.Seq.of(_225_cf29, _137_v0, true));
          let _254_v75;
          _254_v75 = _module.D9.create_DC24(_252_v73);
          let _index26 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_236_v63).length));
          let _index27 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_247_v68).length));
          let _rhs41 = (((_142_v3).contains((_253_v74).IsProperSubsetOf(_253_v74))) ? ((_142_v3).get((_253_v74).IsProperSubsetOf(_253_v74))) : (_224_cf30));
          let _rhs42 = _module.__default.safeModuloInt((new BigNumber(811)).minus(new BigNumber((_183_v40).length)), _226_cf28);
          let _rhs43 = _137_v0;
          let _rhs44 = (_193_v44).f20;
          let _rhs45 = (_254_v75).dtor_cf49;
          let _lhs30 = _236_v63;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_236_v63).length));
          let _lhs32 = _145_globalState;
          let _lhs33 = _145_globalState;
          let _lhs34 = _247_v68;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_247_v68).length));
          _lhs30[_lhs31] = _rhs41;
          _lhs32.f2 = _rhs42;
          _lhs33.f3 = _rhs43;
          _lhs34[_lhs35] = _rhs44;
          _252_v73 = _rhs45;
          let _255_v76;
          let _nw39 = new _module.C2();
          _nw39.__ctor();
          _255_v76 = _nw39;
          let _256_v77;
          _256_v77 = _module.D8.create_DC22(_255_v76);
          let _257_v78;
          _257_v78 = _227_v56;
          let _258_v79;
          let _nw40 = new _module.C0();
          _nw40.__ctor(_257_v78, _224_cf30);
          _258_v79 = _nw40;
          let _rhs46 = _dafny.Seq.update(_229_v58, _module.__default.safeIndex(_module.__default.safeModuloInt(_141_v2, (_236_v63)[_module.__default.safeIndex(new BigNumber(881), new BigNumber((_236_v63).length))]), new BigNumber((_229_v58).length)), true);
          let _rhs47 = _222_cf32;
          let _rhs48 = _256_v77;
          let _rhs49 = _258_v79;
          let _lhs36 = _145_globalState;
          _229_v58 = _rhs46;
          _lhs36.f10 = _rhs47;
          _256_v77 = _rhs48;
          _258_v79 = _rhs49;
          let _259_v80;
          _259_v80 = _module.D4.create_DC9(_258_v79);
          _259_v80 = _259_v80;
        } else {
          let _260___mcc_h18 = (_source4).cf18;
          let _261_cf18 = _260___mcc_h18;
          let _index28 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_233_v62).length));
          (_233_v62)[_index28] = _180_v37;
          let _262_v81;
          _262_v81 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-665),_225_cf29);
          let _263_v82;
          _263_v82 = _dafny.Map.Empty.slice().updateUnsafe(_262_v81,_141_v2);
          let _264_v84;
          _264_v84 = _dafny.MultiSet.fromElements(_262_v81);
          let _265_v86;
          _265_v86 = _module.D10.create_DC28(_dafny.Map.Empty.slice().updateUnsafe(_262_v81,(_dafny.ZERO).minus(_224_cf30)));
          let _266_v88;
          let _nw41 = Array((new BigNumber(27)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = _module.__default.fm13(_145_globalState);
          _nw41[(_dafny.ONE).toNumber()] = (_263_v82).Merge(_263_v82);
          _nw41[(new BigNumber(2)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(3)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(4)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(5)).toNumber()] = (_263_v82).Merge(_263_v82);
          _nw41[(new BigNumber(6)).toNumber()] = (function () {
            let _coll10 = new _dafny.Map();
            for (const _compr_10 of (_264_v84).Elements) {
              let _267_v83 = _compr_10;
              if ((_264_v84).contains(_267_v83)) {
                _coll10.push([_267_v83,_226_cf28]);
              }
            }
            return _coll10;
          }()).update(_262_v81, new BigNumber((_module.__default.fm8(_226_cf28, new BigNumber(-844), _145_globalState)).length));
          _nw41[(new BigNumber(7)).toNumber()] = (_263_v82).Merge(_263_v82);
          _nw41[(new BigNumber(8)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(9)).toNumber()] = (function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of (_264_v84).Elements) {
              let _268_v85 = _compr_11;
              if ((_264_v84).contains(_268_v85)) {
                _coll11.push([_268_v85,_224_cf30]);
              }
            }
            return _coll11;
          }()).update(_262_v81, _223_cf31);
          _nw41[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_262_v81,new BigNumber((_dafny.Seq.UnicodeFromString("uavoowfky")).length));
          _nw41[(new BigNumber(11)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(12)).toNumber()] = (_263_v82).Merge(_263_v82);
          _nw41[(new BigNumber(13)).toNumber()] = ((_265_v86).dtor_cf56).update(_dafny.Map.Empty.slice().updateUnsafe(_227_v56,_137_v0), _141_v2);
          _nw41[(new BigNumber(14)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(15)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(16)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(17)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(18)).toNumber()] = (_263_v82).Merge(_263_v82);
          _nw41[(new BigNumber(19)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of (_238_v65).Elements) {
              let _269_v87 = _compr_12;
              if ((_238_v65).contains(_269_v87)) {
                _coll12.push([(_269_v87).minus(_224_cf30),_222_cf32]);
              }
            }
            return _coll12;
          }(),_224_cf30);
          _nw41[(new BigNumber(20)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(21)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(22)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(23)).toNumber()] = _263_v82;
          _nw41[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_223_cf31,_222_cf32),_226_cf28);
          _nw41[(new BigNumber(25)).toNumber()] = (_263_v82).Merge(_263_v82);
          _nw41[(new BigNumber(26)).toNumber()] = (_263_v82).Merge(_263_v82);
          _266_v88 = _nw41;
          let _index29 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_266_v88).length));
          (_266_v88)[_index29] = _263_v82;
          let _index30 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_266_v88).length));
          let _rhs50 = ((_dafny.Map.Empty.slice().updateUnsafe(_262_v81,_141_v2)).Merge(function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of (_263_v82).Keys.Elements) {
              let _270_v89 = _compr_13;
              if ((_263_v82).contains(_270_v89)) {
                _coll13.push([_270_v89,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_143_v4,new BigNumber(-269))).length),new BigNumber((_dafny.Seq.of(false)).length))).length)]);
              }
            }
            return _coll13;
          }())).Merge(_263_v82);
          let _rhs51 = _226_cf28;
          let _lhs37 = _266_v88;
          let _lhs38 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_266_v88).length));
          _lhs37[_lhs38] = _rhs50;
          _223_cf31 = _rhs51;
          let _271_v90;
          _271_v90 = _module.D2.create_DC6(new BigNumber(-696), _222_cf32, _222_cf32, _227_v56);
          _271_v90 = _271_v90;
          _226_cf28 = (_141_v2).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_138_v1,_236_v63)).length));
        }
        let _272_v91;
        let _nw42 = new _module.C2();
        _nw42.__ctor();
        _272_v91 = _nw42;
        _272_v91 = _272_v91;
      } else if (_source3.is_DC17) {
        let _273___mcc_h10 = (_source3).cf33;
        let _274___mcc_h11 = (_source3).cf34;
        let _275___mcc_h12 = (_source3).cf35;
        let _276___mcc_h13 = (_source3).cf36;
        let _277_cf36 = _276___mcc_h13;
        let _278_cf35 = _275___mcc_h12;
        let _279_cf34 = _274___mcc_h11;
        let _280_cf33 = _273___mcc_h10;
        (_145_globalState).f2 = (_141_v2).multipliedBy((new BigNumber(-556)).plus(_141_v2));
        _279_cf34 = false;
        let _281_v92;
        _281_v92 = _dafny.Seq.of(_180_v37, _180_v37, _dafny.Seq.of(new BigNumber((_143_v4).length)), _180_v37);
        (_145_globalState).f10 = !(_dafny.Seq.IsProperPrefixOf((_281_v92)[_module.__default.safeIndex(new BigNumber(-944), new BigNumber((_281_v92).length))], _180_v37));
        (_145_globalState).f0 = true;
      } else {
        let _282___mcc_h14 = (_source3).cf22;
        let _283_cf22 = _282___mcc_h14;
        (_145_globalState).f0 = true;
        (_145_globalState).f9 = _141_v2;
        let _284_v93;
        _284_v93 = _module.D9.create_DC27(_141_v2);
        let _pat_let_tv11 = _141_v2;
        if (!(function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of (_dafny.Seq.of(_141_v2)).Elements) {
            let _285_v94 = _compr_14;
            if (_dafny.Seq.contains(_dafny.Seq.of(_141_v2), _285_v94)) {
              _coll14.push([(_285_v94).multipliedBy(_141_v2),_141_v2]);
            }
          }
          return _coll14;
        }()).contains((_141_v2).plus((function (_pat_let10_0) {
          return function (_286_dt__update__tmp_h4) {
            return function (_pat_let11_0) {
              return function (_287_dt__update_hcf55_h0) {
                return _module.D9.create_DC27(_287_dt__update_hcf55_h0);
              }(_pat_let11_0);
            }(_pat_let_tv11);
          }(_pat_let10_0);
        }(_284_v93)).dtor_cf55))) {
          let _288_v95;
          _288_v95 = _dafny.Seq.of(_143_v4, _143_v4, _143_v4, _143_v4, _dafny.Seq.UnicodeFromString("w"));
          let _289_v96;
          _289_v96 = _dafny.Map.Empty.slice().updateUnsafe((_288_v95)[_module.__default.safeIndex(_141_v2, new BigNumber((_288_v95).length))],!(_137_v0));
          _289_v96 = _289_v96;
          let _290_v97;
          let _291_v98;
          let _out6;
          let _out7;
          let _outcollector3 = (_283_cf22).m3(_137_v0, _145_globalState);
          _out6 = _outcollector3[0];
          _out7 = _outcollector3[1];
          _290_v97 = _out6;
          _291_v98 = _out7;
          let _292_v99;
          _292_v99 = (_dafny.ZERO).minus(_290_v97);
          let _293_v100;
          let _nw43 = new _module.C0();
          _nw43.__ctor(_292_v99, new BigNumber((_143_v4).length));
          _293_v100 = _nw43;
          _293_v100 = _293_v100;
          (_145_globalState).f0 = _137_v0;
          let _294_v101;
          _294_v101 = _dafny.Seq.of(_137_v0);
          (_145_globalState).f3 = _dafny.Seq.IsPrefixOf(_294_v101, _dafny.Seq.of(_137_v0));
        } else {
          let _295_v102;
          let _nw44 = new _module.C2();
          _nw44.__ctor();
          _295_v102 = _nw44;
          let _nw45 = new _module.C2();
          _nw45.__ctor();
          _295_v102 = _nw45;
          let _296_v103;
          let _nw46 = new _module.C1();
          _nw46.__ctor((_283_cf22).f20);
          _296_v103 = _nw46;
          let _297_v104;
          let _298_v105;
          let _out8;
          let _out9;
          let _outcollector4 = (_283_cf22).m3(!(_137_v0), _145_globalState);
          _out8 = _outcollector4[0];
          _out9 = _outcollector4[1];
          _297_v104 = _out8;
          _298_v105 = _out9;
          (_145_globalState).f10 = _137_v0;
          (_145_globalState).f9 = _141_v2;
        }
        let _299_v106;
        let _init6 = ((_300_v0) => function (_301_i4) {
          return (_301_i4).plus(new BigNumber((_dafny.Seq.of(_300_v0)).length));
        })(_137_v0);
        let _nw47 = Array((new BigNumber(5)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw47.length); _i0_6++) {
          _nw47[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _299_v106 = _nw47;
        let _index31 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_299_v106).length));
        (_299_v106)[_index31] = _module.__default.safeModuloInt(new BigNumber(170), _141_v2);
        let _index32 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_299_v106).length));
        (_299_v106)[_index32] = new BigNumber((_183_v40).length);
      }
      let _302_v107;
      _302_v107 = _module.D9.create_DC25(_141_v2);
      let _pat_let_tv12 = _141_v2;
      let _303_v108;
      let _nw48 = Array((new BigNumber(6)).toNumber());
      _nw48[(_dafny.ZERO).toNumber()] = _302_v107;
      _nw48[(_dafny.ONE).toNumber()] = function (_pat_let12_0) {
        return function (_304_dt__update__tmp_h5) {
          return function (_pat_let13_0) {
            return function (_305_dt__update_hcf50_h0) {
              return _module.D9.create_DC25(_305_dt__update_hcf50_h0);
            }(_pat_let13_0);
          }(_pat_let_tv12);
        }(_pat_let12_0);
      }(_302_v107);
      _nw48[(new BigNumber(2)).toNumber()] = _302_v107;
      _nw48[(new BigNumber(3)).toNumber()] = ((_137_v0) ? (_302_v107) : (_302_v107));
      _nw48[(new BigNumber(4)).toNumber()] = _302_v107;
      _nw48[(new BigNumber(5)).toNumber()] = _302_v107;
      _303_v108 = _nw48;
      let _index33 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_303_v108).length));
      (_303_v108)[_index33] = _302_v107;
      let _pat_let_tv13 = _143_v4;
      let _pat_let_tv14 = _143_v4;
      let _index34 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_303_v108).length));
      (_303_v108)[_index34] = function (_pat_let14_0) {
        return function (_306_dt__update__tmp_h6) {
          return function (_pat_let15_0) {
            return function (_307_dt__update_hcf50_h1) {
              return _module.D9.create_DC25(_307_dt__update_hcf50_h1);
            }(_pat_let15_0);
          }(new BigNumber((_dafny.Seq.Concat(_pat_let_tv13, _pat_let_tv14)).length));
        }(_pat_let14_0);
      }(_302_v107);
      if (!(_module.__default.safeDivisionInt(_141_v2, _141_v2)).isEqualTo(new BigNumber((((_137_v0) ? (_dafny.Seq.of(_141_v2)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(453)), ((_308_v2) => function (_309_i5) {
        return (_dafny.ZERO).minus(_308_v2);
      })(_141_v2))))).length))) {
        let _310_v109;
        _310_v109 = _dafny.Map.Empty.slice().updateUnsafe((_193_v44).f20,_137_v0);
        let _rhs52 = _310_v109;
        let _rhs53 = _137_v0;
        let _lhs39 = _145_globalState;
        _310_v109 = _rhs52;
        _lhs39.f0 = _rhs53;
        let _311_v110;
        _311_v110 = _dafny.Seq.of(_dafny.Seq.Concat(_143_v4, _143_v4), _dafny.Seq.UnicodeFromString("k"), _143_v4);
        let _rhs54 = _180_v37;
        let _rhs55 = ((((_137_v0) ? (_137_v0) : (false))) ? (_137_v0) : ((_137_v0) && (_137_v0)));
        let _rhs56 = _311_v110;
        let _lhs40 = _145_globalState;
        _180_v37 = _rhs54;
        _lhs40.f0 = _rhs55;
        _311_v110 = _rhs56;
        _143_v4 = _143_v4;
        let _312_v111;
        _312_v111 = _dafny.Map.Empty.slice().updateUnsafe(_141_v2,_192_v43);
        let _313_v112;
        _313_v112 = _dafny.Seq.of(_137_v0, true, !(_137_v0), _137_v0);
        _183_v40 = _module.__default.fm14((_141_v2).isLessThan(_141_v2), (((_312_v111).contains(new BigNumber((_313_v112).length))) ? ((_312_v111).get(new BigNumber((_313_v112).length))) : (_module.__default.fm7((_193_v44).f20, _137_v0, _145_globalState))), _137_v0, _145_globalState);
        if (!(_module.__default.fm14(_137_v0, new _dafny.CodePoint('c'.codePointAt(0)), !(_137_v0), _145_globalState)).equals(_183_v40)) {
          let _314_v113;
          _314_v113 = _module.D2.create_DC4(_184_v41);
          let _315_v114;
          _315_v114 = _module.D2.create_DC7(_314_v113);
          _315_v114 = _315_v114;
          let _316_v115;
          let _init7 = ((_317_v109) => function (_318_i6) {
            return _317_v109;
          })(_310_v109);
          let _nw49 = Array((new BigNumber(6)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw49.length); _i0_7++) {
            _nw49[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _316_v115 = _nw49;
          let _rhs57 = new BigNumber((_143_v4).length);
          let _rhs58 = _138_v1;
          let _rhs59 = _316_v115;
          let _lhs41 = _145_globalState;
          let _lhs42 = _145_globalState;
          _lhs41.f9 = _rhs57;
          _lhs42.f11 = _rhs58;
          _316_v115 = _rhs59;
          let _rhs60 = (_137_v0) === ((new BigNumber(721)).isLessThan(_141_v2));
          let _rhs61 = _module.__default.fm7((_193_v44).f20, _137_v0, _145_globalState);
          _137_v0 = _rhs60;
          _192_v43 = _rhs61;
          let _319_v116;
          _319_v116 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_module.__default.fm1(_137_v0, _143_v4, _145_globalState), _module.__default.fm6(_137_v0, _145_globalState), _145_globalState),_dafny.Seq.Concat(_143_v4, _dafny.Seq.UnicodeFromString("eph")));
          _319_v116 = (_319_v116).update(_137_v0, _dafny.Seq.UnicodeFromString("aaf"));
          (_145_globalState).f10 = !(((!(_137_v0) || (_137_v0)) ? ((_137_v0) === (_137_v0)) : (_137_v0)));
        } else {
          (_145_globalState).f10 = _137_v0;
          (_145_globalState).f2 = _module.__default.safeModuloInt(new BigNumber((_module.__default.fm8(_141_v2, new BigNumber(873), _145_globalState)).length), _141_v2);
          let _320_v117;
          let _nw50 = new _module.C2();
          _nw50.__ctor();
          _320_v117 = _nw50;
          let _321_v118;
          _321_v118 = _dafny.MultiSet.fromElements(_320_v117);
          (_145_globalState).f0 = (_dafny.MultiSet.fromElements(_320_v117)).IsSubsetOf((_321_v118).Difference(_321_v118));
          let _322_v119;
          _322_v119 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm11(_141_v2, _137_v0, _141_v2, _145_globalState),(_141_v2).isLessThanOrEqualTo((_dafny.ZERO).minus(_141_v2)));
          _322_v119 = _322_v119;
          let _323_v120;
          let _init8 = function (_324_i7) {
            return (_324_i7).minus(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('n'.codePointAt(0)))).length));
          };
          let _nw51 = Array((new BigNumber(14)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw51.length); _i0_8++) {
            _nw51[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _323_v120 = _nw51;
          let _325_v121;
          _325_v121 = _module.D6.create_DC15(new BigNumber((_143_v4).length), _141_v2, _180_v37, _141_v2, _dafny.Set.fromElements(_141_v2));
          let _326_v122;
          _326_v122 = (_325_v121).dtor_cf23;
          let _327_v123;
          let _nw52 = new _module.C0();
          _nw52.__ctor(_326_v122, new BigNumber((_143_v4).length));
          _327_v123 = _nw52;
          let _328_v124;
          _328_v124 = _dafny.Map.Empty.slice().updateUnsafe(_327_v123,_323_v120);
          let _329_v125;
          let _nw53 = Array((new BigNumber(25)).toNumber());
          _nw53[(_dafny.ZERO).toNumber()] = _323_v120;
          _nw53[(_dafny.ONE).toNumber()] = (((_328_v124).contains(_327_v123)) ? ((_328_v124).get(_327_v123)) : (_323_v120));
          _nw53[(new BigNumber(2)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(3)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(4)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(5)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(6)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(7)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(8)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(9)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(10)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(11)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(12)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(13)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(14)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(15)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(16)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(17)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(18)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(19)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(20)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(21)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(22)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(23)).toNumber()] = _323_v120;
          _nw53[(new BigNumber(24)).toNumber()] = _323_v120;
          _329_v125 = _nw53;
          let _index35 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_329_v125).length));
          (_329_v125)[_index35] = _323_v120;
          let _index36 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_329_v125).length));
          (_329_v125)[_index36] = _323_v120;
        }
      } else {
        (_145_globalState).f3 = _137_v0;
        let _330_v126;
        _330_v126 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_141_v2, new BigNumber(681)),(_182_v39).Difference(_182_v39));
        let _331_v127;
        _331_v127 = _dafny.Set.fromElements(_141_v2);
        _330_v126 = _module.__default.fm15((new BigNumber((_331_v127).length)).isEqualTo(_141_v2), (_142_v3).update(true, new BigNumber(-178)), _137_v0, false, _145_globalState);
        let _332_v128;
        _332_v128 = _dafny.Map.Empty.slice().updateUnsafe(_141_v2,(_module.D5.create_DC12(_143_v4)).dtor_cf18);
        _143_v4 = _dafny.Seq.update((((_332_v128).contains(_141_v2)) ? ((_332_v128).get(_141_v2)) : (_dafny.Seq.update(_143_v4, _module.__default.safeIndex(_141_v2, new BigNumber((_143_v4).length)), (_193_v44).f20))), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("kr")).length), new BigNumber(((((_332_v128).contains(_141_v2)) ? ((_332_v128).get(_141_v2)) : (_dafny.Seq.update(_143_v4, _module.__default.safeIndex(_141_v2, new BigNumber((_143_v4).length)), (_193_v44).f20)))).length)), _192_v43);
        (_145_globalState).f3 = _137_v0;
        let _333_v129;
        let _nw54 = Array((new BigNumber(27)).toNumber()).fill([]);
        _333_v129 = _nw54;
        let _index37 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_333_v129).length));
        (_333_v129)[_index37] = _138_v1;
        let _index38 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_333_v129).length));
        (_333_v129)[_index38] = _138_v1;
      }
      let _334_v130;
      _334_v130 = _dafny.Set.fromElements(_193_v44, _193_v44, _193_v44);
      (_145_globalState).f3 = (_334_v130).contains(_193_v44);
      let _335_v131;
      _335_v131 = _dafny.Set.fromElements(new BigNumber((_143_v4).length));
      if ((_dafny.Set.fromElements(_141_v2)).IsDisjointFrom(_335_v131)) {
        let _336_v132;
        let _337_v133;
        let _338_v134;
        let _339_v135;
        let _out10;
        let _out11;
        let _out12;
        let _out13;
        let _outcollector5 = (_193_v44).m4(_145_globalState);
        _out10 = _outcollector5[0];
        _out11 = _outcollector5[1];
        _out12 = _outcollector5[2];
        _out13 = _outcollector5[3];
        _336_v132 = _out10;
        _337_v133 = _out11;
        _338_v134 = _out12;
        _339_v135 = _out13;
        let _340_v136;
        _340_v136 = _337_v133;
        let _341_v137;
        let _nw55 = new _module.C0();
        _nw55.__ctor(_340_v136, _337_v133);
        _341_v137 = _nw55;
        let _342_v138;
        _342_v138 = _dafny.MultiSet.fromElements(_341_v137, _341_v137, _341_v137, _341_v137);
        let _343_v139;
        _343_v139 = _342_v138;
        let _source5 = _342_v138;
        let _344___mcc_h19 = _source5;
        let _345_cf14 = _344___mcc_h19;
        let _346_v140;
        let _nw56 = new _module.C2();
        _nw56.__ctor();
        _346_v140 = _nw56;
        let _347_v141;
        let _nw57 = Array((new BigNumber(17)).toNumber()).fill([]);
        _347_v141 = _nw57;
        let _348_v142;
        let _nw58 = Array((new BigNumber(6)).toNumber());
        _nw58[(_dafny.ZERO).toNumber()] = _141_v2;
        _nw58[(_dafny.ONE).toNumber()] = _module.__default.fm0(_145_globalState);
        _nw58[(new BigNumber(2)).toNumber()] = _141_v2;
        _nw58[(new BigNumber(3)).toNumber()] = _141_v2;
        _nw58[(new BigNumber(4)).toNumber()] = _141_v2;
        _nw58[(new BigNumber(5)).toNumber()] = new BigNumber((_143_v4).length);
        _348_v142 = _nw58;
        let _index39 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_347_v141).length));
        (_347_v141)[_index39] = _348_v142;
        let _index40 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_347_v141).length));
        (_347_v141)[_index40] = _348_v142;
        (_145_globalState).f2 = _337_v133;
        _192_v43 = (_193_v44).f20;
        let _rhs62 = _138_v1;
        let _rhs63 = _137_v0;
        let _lhs43 = _145_globalState;
        _138_v1 = _rhs62;
        _lhs43.f10 = _rhs63;
        (_145_globalState).f2 = _337_v133;
        let _349_v143;
        _349_v143 = _dafny.Seq.of(_338_v134, _137_v0);
        let _350_v144;
        _350_v144 = _dafny.Map.Empty.slice().updateUnsafe(_349_v143,(_181_v38).dtor_cf2);
        let _351_v145;
        _351_v145 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_350_v144).length),(((_142_v3).contains(_336_v132)) ? ((_142_v3).get(_336_v132)) : (_337_v133)));
        _337_v133 = (((_351_v145).contains(_337_v133)) ? ((_351_v145).get(_337_v133)) : (_337_v133));
      } else {
        let _352_v146;
        let _nw59 = Array((new BigNumber(22)).toNumber());
        _nw59[(_dafny.ZERO).toNumber()] = _141_v2;
        _nw59[(_dafny.ONE).toNumber()] = _141_v2;
        _nw59[(new BigNumber(2)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(3)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(4)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(5)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(6)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(7)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(8)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(9)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(10)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(11)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(12)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(13)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(14)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("ify")).length);
        _nw59[(new BigNumber(16)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(17)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(18)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(19)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(20)).toNumber()] = _141_v2;
        _nw59[(new BigNumber(21)).toNumber()] = _141_v2;
        _352_v146 = _nw59;
        let _353_v147;
        _353_v147 = _module.D4.create_DC10(_352_v146);
        let _354_v148;
        _354_v148 = _module.D4.create_DC11(_353_v147);
        let _pat_let_tv15 = _352_v146;
        let _355_v149;
        _355_v149 = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,function (_pat_let16_0) {
          return function (_356_dt__update__tmp_h8) {
            return function (_pat_let17_0) {
              return function (_357_dt__update_hcf17_h0) {
                return _module.D4.create_DC11(_357_dt__update_hcf17_h0);
              }(_pat_let17_0);
            }(_module.D4.create_DC10(_pat_let_tv15));
          }(_pat_let16_0);
        }(_354_v148));
        let _pat_let_tv16 = _353_v147;
        let _358_v150;
        _358_v150 = _dafny.Set.fromElements(_355_v149, _355_v149, _dafny.Map.Empty.slice().updateUnsafe(_137_v0,function (_pat_let18_0) {
          return function (_359_dt__update__tmp_h9) {
            return function (_pat_let19_0) {
              return function (_360_dt__update_hcf17_h1) {
                return _module.D4.create_DC11(_360_dt__update_hcf17_h1);
              }(_pat_let19_0);
            }(_pat_let_tv16);
          }(_pat_let18_0);
        }(_354_v148)), _355_v149);
        let _source6 = _module.D7.create_DC18((_358_v150).Union(_dafny.Set.fromElements(_355_v149, _355_v149, _355_v149, _355_v149)));
        if (_source6.is_DC19) {
          let _361___mcc_h20 = (_source6).cf38;
          let _362___mcc_h21 = (_source6).cf39;
          let _363___mcc_h22 = (_source6).cf40;
          let _364_cf40 = _363___mcc_h22;
          let _365_cf39 = _362___mcc_h21;
          let _366_cf38 = _361___mcc_h20;
          let _367_v151;
          _367_v151 = _module.D2.create_DC7(_module.D2.create_DC4(_184_v41));
          let _368_v152;
          _368_v152 = _module.D2.create_DC4(_184_v41);
          _367_v151 = _module.D2.create_DC7(_368_v152);
          (_145_globalState).f9 = ((false) ? ((_141_v2).plus(_141_v2)) : (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mject"), _143_v4)).length)));
          let _369_v153;
          _369_v153 = _141_v2;
          let _370_v154;
          let _nw60 = new _module.C0();
          _nw60.__ctor(_369_v153, _141_v2);
          _370_v154 = _nw60;
          let _371_v155;
          let _nw61 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _371_v155 = _nw61;
          let _372_v156;
          _372_v156 = _module.D9.create_DC24(_371_v155);
          let _373_v157;
          _373_v157 = _module.D9.create_DC24((_372_v156).dtor_cf49);
          let _374_v158;
          _374_v158 = _dafny.Seq.of(_365_cf39);
          let _rhs64 = new BigNumber(373);
          let _rhs65 = _module.__default.fm1(_365_cf39, _143_v4, _145_globalState);
          let _rhs66 = _module.__default.fm0(_145_globalState);
          let _rhs67 = ((!((_module.D2.create_DC5(_137_v0, new BigNumber((_374_v158).length), _138_v1, _178_v35)).dtor_cf5)) ? (((_365_cf39) ? (_372_v156) : (_372_v156))) : (_373_v157));
          let _lhs44 = _145_globalState;
          let _lhs45 = _145_globalState;
          _lhs44.f2 = _rhs64;
          _365_cf39 = _rhs65;
          _lhs45.f9 = _rhs66;
          _372_v156 = _rhs67;
        } else if (_source6.is_DC20) {
          let _375___mcc_h23 = (_source6).cf41;
          let _376___mcc_h24 = (_source6).cf42;
          let _377___mcc_h25 = (_source6).cf43;
          let _378_cf43 = _377___mcc_h25;
          let _379_cf42 = _376___mcc_h24;
          let _380_cf41 = _375___mcc_h23;
          (_145_globalState).f0 = _137_v0;
          let _index41 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_138_v1).length));
          (_138_v1)[_index41] = _137_v0;
          let _index42 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_138_v1).length));
          let _rhs68 = true;
          let _rhs69 = ((_378_cf43).plus(_378_cf43)).plus((_dafny.ZERO).minus(_141_v2));
          let _lhs46 = _138_v1;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_138_v1).length));
          _lhs46[_lhs47] = _rhs68;
          _378_cf43 = _rhs69;
          let _381_v159;
          _381_v159 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(68)), ((_382_cf43) => function (_383_i8) {
            return _382_cf43;
          })(_378_cf43))).length);
          let _384_v160;
          _384_v160 = _dafny.Map.Empty.slice().updateUnsafe(_378_cf43,true);
          let _385_v161;
          let _nw62 = new _module.C0();
          _nw62.__ctor(_381_v159, new BigNumber((_384_v160).length));
          _385_v161 = _nw62;
          let _386_v162;
          _386_v162 = _module.D4.create_DC9(_385_v161);
          let _387_v163;
          _387_v163 = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_385_v161);
          let _388_v164;
          let _nw63 = Array((new BigNumber(10)).toNumber());
          _nw63[(_dafny.ZERO).toNumber()] = _385_v161;
          _nw63[(_dafny.ONE).toNumber()] = _385_v161;
          _nw63[(new BigNumber(2)).toNumber()] = _385_v161;
          _nw63[(new BigNumber(3)).toNumber()] = (_386_v162).dtor_cf15;
          _nw63[(new BigNumber(4)).toNumber()] = _385_v161;
          _nw63[(new BigNumber(5)).toNumber()] = _385_v161;
          _nw63[(new BigNumber(6)).toNumber()] = _385_v161;
          _nw63[(new BigNumber(7)).toNumber()] = _385_v161;
          _nw63[(new BigNumber(8)).toNumber()] = _385_v161;
          _nw63[(new BigNumber(9)).toNumber()] = (((_387_v163).contains((_138_v1)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_138_v1).length))])) ? ((_387_v163).get((_138_v1)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_138_v1).length))])) : (_385_v161));
          _388_v164 = _nw63;
          let _index43 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_388_v164).length));
          (_388_v164)[_index43] = _385_v161;
          let _index44 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_388_v164).length));
          (_388_v164)[_index44] = _385_v161;
          let _389_v165;
          let _390_v166;
          let _out14;
          let _out15;
          let _outcollector6 = _module.__default.m0((_138_v1)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_138_v1).length))], _384_v160, _145_globalState);
          _out14 = _outcollector6[0];
          _out15 = _outcollector6[1];
          _389_v165 = _out14;
          _390_v166 = _out15;
        } else if (_source6.is_DC18) {
          let _391___mcc_h26 = (_source6).cf37;
          let _392_cf37 = _391___mcc_h26;
          let _rhs70 = _dafny.Seq.Concat(_143_v4, _dafny.Seq.Concat(_dafny.Seq.of((_193_v44).f20, (_193_v44).f20), _143_v4));
          let _rhs71 = (_141_v2).multipliedBy((_141_v2).multipliedBy(_141_v2));
          let _lhs48 = _145_globalState;
          _143_v4 = _rhs70;
          _lhs48.f9 = _rhs71;
          let _index45 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_352_v146).length));
          (_352_v146)[_index45] = _141_v2;
          let _393_v167;
          let _nw64 = Array((new BigNumber(25)).toNumber()).fill(_module.D10.Default());
          _393_v167 = _nw64;
          let _394_v168;
          _394_v168 = _module.D10.create_DC29();
          let _index46 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_393_v167).length));
          (_393_v167)[_index46] = _394_v168;
          let _index47 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_352_v146).length));
          (_352_v146)[_index47] = _module.__default.safeDivisionInt(_module.__default.fm0(_145_globalState), new BigNumber((_335_v131).length));
          let _395_v169;
          let _nw65 = new _module.C0();
          _nw65.__ctor(new BigNumber((_180_v37).length), _141_v2);
          _395_v169 = _nw65;
          let _396_v170;
          _396_v170 = _module.D2.create_DC6(_141_v2, true, _137_v0, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_144_v5).contains(_137_v0)) ? ((_144_v5).get(_137_v0)) : (true)),_395_v169)).length)));
          let _index48 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_352_v146).length));
          let _index49 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_393_v167).length));
          let _index50 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_352_v146).length));
          let _rhs72 = (_141_v2).plus(_141_v2);
          let _rhs73 = _137_v0;
          let _rhs74 = _module.D10.create_DC29();
          let _rhs75 = _module.__default.fm0(_145_globalState);
          let _rhs76 = (_396_v170).dtor_cf12;
          let _lhs49 = _352_v146;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_352_v146).length));
          let _lhs51 = _145_globalState;
          let _lhs52 = _393_v167;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_393_v167).length));
          let _lhs54 = _145_globalState;
          let _lhs55 = _352_v146;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_352_v146).length));
          _lhs49[_lhs50] = _rhs72;
          _lhs51.f3 = _rhs73;
          _lhs52[_lhs53] = _rhs74;
          _lhs54.f9 = _rhs75;
          _lhs55[_lhs56] = _rhs76;
          let _397_v171;
          _397_v171 = _dafny.Map.Empty.slice().updateUnsafe(_141_v2,(_352_v146)[_module.__default.safeIndex(new BigNumber(356), new BigNumber((_352_v146).length))]);
          _137_v0 = ((!((_397_v171).contains(new BigNumber((_dafny.Seq.UnicodeFromString("hiwhxnyyd")).length)))) ? (_137_v0) : (_137_v0));
          let _398_v172;
          _398_v172 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_141_v2, new BigNumber((_143_v4).length)),_137_v0);
          _398_v172 = (_398_v172).update(((_352_v146)[_module.__default.safeIndex(new BigNumber(356), new BigNumber((_352_v146).length))]).minus((_352_v146)[_module.__default.safeIndex(new BigNumber(356), new BigNumber((_352_v146).length))]), ((_137_v0) ? (_137_v0) : (_137_v0)));
        } else {
          let _399___mcc_h27 = (_source6).cf44;
          let _400_cf44 = _399___mcc_h27;
          let _index51 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_138_v1).length));
          (_138_v1)[_index51] = !(_141_v2).isEqualTo(_141_v2);
          let _index52 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_138_v1).length));
          (_138_v1)[_index52] = _137_v0;
          let _401_v173;
          _401_v173 = (_dafny.ZERO).minus(_141_v2);
          let _402_v174;
          _402_v174 = _dafny.MultiSet.fromElements((_193_v44).f20, _192_v43);
          let _403_v175;
          let _nw66 = new _module.C0();
          _nw66.__ctor(_401_v173, new BigNumber((_402_v174).cardinality()));
          _403_v175 = _nw66;
          let _404_v176;
          _404_v176 = _dafny.Map.Empty.slice().updateUnsafe(_141_v2,_403_v175);
          let _405_v177;
          _405_v177 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),new BigNumber((_404_v176).length));
          let _rhs77 = _352_v146;
          let _rhs78 = (_405_v177).Merge(_405_v177);
          _352_v146 = _rhs77;
          _405_v177 = _rhs78;
          _192_v43 = (_193_v44).f20;
          let _406_v178;
          let _init9 = function (_407_i9) {
            return true;
          };
          let _nw67 = Array((new BigNumber(22)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw67.length); _i0_9++) {
            _nw67[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _406_v178 = _nw67;
          let _408_v179;
          _408_v179 = _module.D2.create_DC5(_137_v0, (_403_v175).f18, _406_v178, _178_v35);
          let _409_v180;
          _409_v180 = _dafny.Set.fromElements(_408_v179);
          let _410_v181;
          _410_v181 = _dafny.Map.Empty.slice().updateUnsafe((_409_v180).Difference(_409_v180),(_dafny.ZERO).minus((_180_v37)[_module.__default.safeIndex((_403_v175).f18, new BigNumber((_180_v37).length))]));
          _410_v181 = (_410_v181).update(_409_v180, new BigNumber(((_module.__default.fm16((_403_v175).f18, (_138_v1)[_module.__default.safeIndex(new BigNumber(987), new BigNumber((_138_v1).length))], new _dafny.CodePoint('b'.codePointAt(0)), _141_v2, _145_globalState)).Merge(_144_v5)).length));
        }
        let _411_v182;
        _411_v182 = _dafny.Seq.of(_143_v4, _143_v4);
        let _412_v183;
        let _nw68 = Array((new BigNumber(28)).toNumber());
        _nw68[(_dafny.ZERO).toNumber()] = _143_v4;
        _nw68[(_dafny.ONE).toNumber()] = _module.__default.fm6(_137_v0, _145_globalState);
        _nw68[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-743)), ((_413_v44) => function (_414_i10) {
          return (_413_v44).f20;
        })(_193_v44));
        _nw68[(new BigNumber(3)).toNumber()] = _143_v4;
        _nw68[(new BigNumber(4)).toNumber()] = _143_v4;
        _nw68[(new BigNumber(5)).toNumber()] = _143_v4;
        _nw68[(new BigNumber(6)).toNumber()] = _143_v4;
        _nw68[(new BigNumber(7)).toNumber()] = _module.__default.fm6(_137_v0, _145_globalState);
        _nw68[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_143_v4, _dafny.Seq.UnicodeFromString("l"));
        _nw68[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_143_v4, _143_v4);
        _nw68[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("veh");
        _nw68[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(638)), ((_415_v43) => function (_416_i11) {
          return _415_v43;
        })(_192_v43));
        _nw68[(new BigNumber(12)).toNumber()] = _143_v4;
        _nw68[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_143_v4, _143_v4);
        _nw68[(new BigNumber(14)).toNumber()] = _143_v4;
        _nw68[(new BigNumber(15)).toNumber()] = _143_v4;
        _nw68[(new BigNumber(16)).toNumber()] = _dafny.Seq.update(_143_v4, _module.__default.safeIndex(_141_v2, new BigNumber((_143_v4).length)), (_193_v44).f20);
        _nw68[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(489)), ((_417_v44) => function (_418_i12) {
          return (_417_v44).f20;
        })(_193_v44));
        _nw68[(new BigNumber(18)).toNumber()] = _143_v4;
        _nw68[(new BigNumber(19)).toNumber()] = (_411_v182)[_module.__default.safeIndex(new BigNumber(-360), new BigNumber((_411_v182).length))];
        _nw68[(new BigNumber(20)).toNumber()] = _143_v4;
        _nw68[(new BigNumber(21)).toNumber()] = _143_v4;
        _nw68[(new BigNumber(22)).toNumber()] = _dafny.Seq.UnicodeFromString("psmshitr");
        _nw68[(new BigNumber(23)).toNumber()] = _dafny.Seq.update(_143_v4, _module.__default.safeIndex(_141_v2, new BigNumber((_143_v4).length)), new _dafny.CodePoint('f'.codePointAt(0)));
        _nw68[(new BigNumber(24)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(478)), function (_419_i13) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        });
        _nw68[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_143_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-912)), function (_420_i14) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }));
        _nw68[(new BigNumber(26)).toNumber()] = _143_v4;
        _nw68[(new BigNumber(27)).toNumber()] = _143_v4;
        _412_v183 = _nw68;
        let _index53 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_412_v183).length));
        (_412_v183)[_index53] = _143_v4;
        let _index54 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_412_v183).length));
        (_412_v183)[_index54] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(762)), ((_421_v44) => function (_422_i15) {
          return (_421_v44).f20;
        })(_193_v44));
        let _index55 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_138_v1).length));
        (_138_v1)[_index55] = _dafny.Seq.IsPrefixOf(_180_v37, _180_v37);
        let _index56 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_138_v1).length));
        (_138_v1)[_index56] = (((_303_v108)[_module.__default.safeIndex(new BigNumber(481), new BigNumber((_303_v108).length))]).dtor_cf50).isEqualTo(_141_v2);
        let _423_v184;
        _423_v184 = _module.D11.create_DC31(_182_v39);
        (_145_globalState).f2 = _module.__default.safeDivisionInt((_141_v2).multipliedBy(new BigNumber(((_423_v184).dtor_cf58).cardinality())), _141_v2);
        let _424_v185;
        let _nw69 = new _module.C1();
        _nw69.__ctor(_192_v43);
        _424_v185 = _nw69;
      }
      let _425_v186;
      let _init10 = ((_426_v131) => function (_427_i16) {
        return _426_v131;
      })(_335_v131);
      let _nw70 = Array((new BigNumber(27)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw70.length); _i0_10++) {
        _nw70[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _425_v186 = _nw70;
      let _index57 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_425_v186).length));
      (_425_v186)[_index57] = (_335_v131).Difference(_335_v131);
      let _index58 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_425_v186).length));
      (_425_v186)[_index58] = _335_v131;
      (_145_globalState).f0 = (_141_v2).isLessThanOrEqualTo((_141_v2).minus(_141_v2));
      _142_v3 = (_142_v3).Merge((_142_v3).Merge(_142_v3));
      let _428_v187;
      _428_v187 = _module.D6.create_DC15(new BigNumber(818), _141_v2, _180_v37, _141_v2, _dafny.Set.fromElements(_141_v2));
      let _hi1 = _141_v2;
      for (let _429_i17 = new BigNumber(((_428_v187).dtor_cf27).length); _429_i17.isLessThan(_hi1); _429_i17 = _429_i17.plus(_dafny.ONE)) {
        let _rhs79 = _137_v0;
        let _rhs80 = _137_v0;
        let _rhs81 = _429_i17;
        let _lhs57 = _145_globalState;
        let _lhs58 = _145_globalState;
        _lhs57.f0 = _rhs79;
        _lhs58.f0 = _rhs80;
        _141_v2 = _rhs81;
        let _430_v188;
        let _init11 = ((_431_i17) => function (_432_i18) {
          return (_432_i18).multipliedBy(_431_i17);
        })(_429_i17);
        let _nw71 = Array((new BigNumber(18)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw71.length); _i0_11++) {
          _nw71[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _430_v188 = _nw71;
        let _index59 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_430_v188).length));
        (_430_v188)[_index59] = _141_v2;
        let _433_v189;
        _433_v189 = _module.D4.create_DC10(_430_v188);
        let _pat_let_tv17 = _430_v188;
        let _434_v190;
        let _nw72 = Array((new BigNumber(13)).toNumber());
        _nw72[(_dafny.ZERO).toNumber()] = _433_v189;
        _nw72[(_dafny.ONE).toNumber()] = _433_v189;
        _nw72[(new BigNumber(2)).toNumber()] = _433_v189;
        _nw72[(new BigNumber(3)).toNumber()] = _433_v189;
        _nw72[(new BigNumber(4)).toNumber()] = _433_v189;
        _nw72[(new BigNumber(5)).toNumber()] = _433_v189;
        _nw72[(new BigNumber(6)).toNumber()] = _433_v189;
        _nw72[(new BigNumber(7)).toNumber()] = _433_v189;
        _nw72[(new BigNumber(8)).toNumber()] = function (_pat_let20_0) {
          return function (_435_dt__update__tmp_h10) {
            return function (_pat_let21_0) {
              return function (_436_dt__update_hcf16_h0) {
                return _module.D4.create_DC10(_436_dt__update_hcf16_h0);
              }(_pat_let21_0);
            }(_pat_let_tv17);
          }(_pat_let20_0);
        }(_433_v189);
        _nw72[(new BigNumber(9)).toNumber()] = _433_v189;
        _nw72[(new BigNumber(10)).toNumber()] = _433_v189;
        _nw72[(new BigNumber(11)).toNumber()] = _module.D4.create_DC10(_430_v188);
        _nw72[(new BigNumber(12)).toNumber()] = _433_v189;
        _434_v190 = _nw72;
        let _437_v191;
        _437_v191 = _dafny.Map.Empty.slice().updateUnsafe(_429_i17,(_dafny.ZERO).minus((_428_v187).dtor_cf26));
        let _438_v192;
        _438_v192 = _dafny.Set.fromElements(_434_v190, _434_v190);
        let _pat_let_tv18 = _438_v192;
        let _pat_let_tv19 = _141_v2;
        let _index60 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_430_v188).length));
        (_430_v188)[_index60] = (function (_pat_let22_0) {
          return function (_439_dt__update__tmp_h11) {
            return function (_pat_let23_0) {
              return function (_440_dt__update_hcf41_h0) {
                return function (_pat_let24_0) {
                  return function (_441_dt__update_hcf43_h0) {
                    return _module.D7.create_DC20(_440_dt__update_hcf41_h0, (_439_dt__update__tmp_h11).dtor_cf42, _441_dt__update_hcf43_h0);
                  }(_pat_let24_0);
                }(_pat_let_tv19);
              }(_pat_let23_0);
            }(_pat_let_tv18);
          }(_pat_let22_0);
        }(_module.D7.create_DC20(_dafny.Set.fromElements(_434_v190, _434_v190), _module.D5.create_DC13(_module.__default.fm0(_145_globalState), _module.__default.fm0(_145_globalState), _137_v0), new BigNumber((_437_v191).length)))).dtor_cf43;
        let _442_v193;
        _442_v193 = _429_i17;
        let _443_v194;
        let _nw73 = new _module.C0();
        _nw73.__ctor(_442_v193, (_141_v2).minus((_dafny.ZERO).minus(_141_v2)));
        _443_v194 = _nw73;
        let _444_v195;
        _444_v195 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_137_v0, _dafny.Seq.UnicodeFromString("pdygc"), _145_globalState),_443_v194);
        _443_v194 = (((_444_v195).contains(_137_v0)) ? ((_444_v195).get(_137_v0)) : (_443_v194));
        let _445_v196;
        _445_v196 = _module.D11.create_DC31(_182_v39);
        let _446_v197;
        _446_v197 = _module.D11.create_DC33(_445_v196);
        let _447_v198;
        _447_v198 = _module.D11.create_DC33(_446_v197);
        _447_v198 = _447_v198;
      }
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_138_v1).length))) {
        let _448_i19 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_448_i19)) && ((_448_i19).isLessThan(new BigNumber((_138_v1).length))))) {
          (_138_v1)[(_448_i19)] = !(_141_v2).isEqualTo((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_141_v2,false)).length)).plus(new BigNumber((_143_v4).length)));
        }
      }
      let _449_v199;
      _449_v199 = _dafny.Set.fromElements(_143_v4, _143_v4, _dafny.Seq.UnicodeFromString("r"));
      _449_v199 = _449_v199;
      let _hi2 = _141_v2;
      for (let _450_i20 = _module.__default.safeModuloInt(_141_v2, _141_v2); _450_i20.isLessThan(_hi2); _450_i20 = _450_i20.plus(_dafny.ONE)) {
        (_145_globalState).f9 = _141_v2;
        if (!(_137_v0)) {
          let _451_v200;
          let _nw74 = new _module.C3();
          _nw74.__ctor(_143_v4, new BigNumber(633), _141_v2);
          _451_v200 = _nw74;
          let _452_v201;
          _452_v201 = _450_i20;
          let _453_v202;
          let _nw75 = new _module.C0();
          _nw75.__ctor(_452_v201, (_451_v200).fm3(_137_v0, _450_i20, _145_globalState));
          _453_v202 = _nw75;
          let _454_v203;
          _454_v203 = _dafny.MultiSet.fromElements(_453_v202, _453_v202, _453_v202, _453_v202);
          let _455_v204;
          let _nw76 = Array((new BigNumber(11)).toNumber());
          _nw76[(_dafny.ZERO).toNumber()] = _141_v2;
          _nw76[(_dafny.ONE).toNumber()] = _450_i20;
          _nw76[(new BigNumber(2)).toNumber()] = _450_i20;
          _nw76[(new BigNumber(3)).toNumber()] = new BigNumber((_454_v203).cardinality());
          _nw76[(new BigNumber(4)).toNumber()] = (_451_v200).fm3(_137_v0, _141_v2, _145_globalState);
          _nw76[(new BigNumber(5)).toNumber()] = _141_v2;
          _nw76[(new BigNumber(6)).toNumber()] = _450_i20;
          _nw76[(new BigNumber(7)).toNumber()] = new BigNumber(834);
          _nw76[(new BigNumber(8)).toNumber()] = _141_v2;
          _nw76[(new BigNumber(9)).toNumber()] = _141_v2;
          _nw76[(new BigNumber(10)).toNumber()] = new BigNumber(-510);
          _455_v204 = _nw76;
          let _456_v205;
          _456_v205 = _module.D4.create_DC10(_455_v204);
          _456_v205 = _module.D4.create_DC10(_455_v204);
          let _457_v206;
          _457_v206 = _dafny.MultiSet.fromElements(_138_v1, _138_v1, _138_v1, _138_v1, _138_v1);
          _457_v206 = _457_v206;
          let _458_v207;
          let _459_v208;
          let _out16;
          let _out17;
          let _outcollector7 = (_451_v200).m1(_145_globalState);
          _out16 = _outcollector7[0];
          _out17 = _outcollector7[1];
          _458_v207 = _out16;
          _459_v208 = _out17;
          let _460_v209;
          _460_v209 = _dafny.Map.Empty.slice().updateUnsafe(_138_v1,_183_v40);
          let _rhs82 = _458_v207;
          let _rhs83 = new BigNumber((((_460_v209).Merge(_460_v209)).update(_138_v1, (_183_v40).Union(_183_v40))).length);
          let _rhs84 = _451_v200;
          _458_v207 = _rhs82;
          _141_v2 = _rhs83;
          _451_v200 = _rhs84;
        } else {
          let _461_v210;
          let _nw77 = new _module.C2();
          _nw77.__ctor();
          _461_v210 = _nw77;
          let _462_v211;
          _462_v211 = _dafny.MultiSet.fromElements(_461_v210);
          let _463_v212;
          _463_v212 = _dafny.Set.fromElements(_462_v211, _462_v211);
          (_145_globalState).f0 = ((_463_v212).Difference(_463_v212)).IsProperSubsetOf(_463_v212);
          let _464_v213;
          _464_v213 = _dafny.Map.Empty.slice().updateUnsafe(_461_v210,false);
          let _465_v215;
          _465_v215 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_141_v2),!(_137_v0));
          let _466_v216;
          _466_v216 = _dafny.Set.fromElements(_465_v215, _465_v215, _465_v215, _dafny.Map.Empty.slice().updateUnsafe(_141_v2,_137_v0), _465_v215);
          let _467_v217;
          _467_v217 = _dafny.Map.Empty.slice().updateUnsafe((((_464_v213).contains(_461_v210)) ? ((_464_v213).get(_461_v210)) : (_137_v0)),_module.D10.create_DC28(function () {
  let _coll15 = new _dafny.Map();
  for (const _compr_15 of (_466_v216).Elements) {
    let _468_v214 = _compr_15;
    if ((_466_v216).contains(_468_v214)) {
      _coll15.push([_468_v214,_141_v2]);
    }
  }
  return _coll15;
}()));
          let _469_v218;
          _469_v218 = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_467_v217);
          _469_v218 = _469_v218;
          let _470_v219;
          _470_v219 = _450_i20;
          let _471_v220;
          let _nw78 = new _module.C0();
          _nw78.__ctor(_470_v219, _450_i20);
          _471_v220 = _nw78;
          let _nw79 = new _module.C0();
          _nw79.__ctor(_470_v219, (_471_v220).f18);
          _471_v220 = _nw79;
          (_145_globalState).f2 = ((!_dafny.areEqual(_module.__default.fm6(_137_v0, _145_globalState), _module.__default.fm6((_181_v38).dtor_cf2, _145_globalState))) ? (_module.__default.safeDivisionInt(_141_v2, _141_v2)) : ((_471_v220).f18));
          _465_v215 = (_465_v215).update(((_137_v0) ? (new BigNumber(925)) : (new BigNumber((_dafny.Seq.of(_137_v0, _137_v0, _137_v0)).length))), _137_v0);
        }
        (_145_globalState).f3 = _dafny.Seq.IsProperPrefixOf(_module.__default.fm11(_450_i20, _137_v0, _450_i20, _145_globalState), _dafny.Seq.of(_450_i20, _450_i20));
        let _472_v221;
        let _nw80 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _472_v221 = _nw80;
        let _473_v222;
        _473_v222 = _dafny.Set.fromElements(_472_v221, _472_v221);
        let _474_v223;
        _474_v223 = _module.D7.create_DC19(_473_v222, _137_v0, _137_v0);
        let _475_v224;
        _475_v224 = _module.D7.create_DC21(_474_v223);
        _475_v224 = _475_v224;
      }
      process.stdout.write(_dafny.toString(_137_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_141_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_142_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(213)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_143_v4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_144_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_145_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_145_globalState).f1, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_145_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_145_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f7)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f8).equals(_dafny.MultiSet.fromElements(new BigNumber(213), new BigNumber(213)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_145_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_145_globalState.f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState.f11)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f12).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-213)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_145_globalState).f13).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f15).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v35).equals(_dafny.MultiSet.fromElements(new BigNumber(212), new BigNumber(212), new BigNumber(212)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v36).dtor_cf28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v36).dtor_cf29));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v36).dtor_cf30));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v36).dtor_cf31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v36).dtor_cf32));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_180_v37, _dafny.Seq.of(new BigNumber(-496), new BigNumber(212), new BigNumber(212), new BigNumber(212), new BigNumber(-496), new BigNumber(212), new BigNumber(212)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_181_v38).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_v39).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_183_v40).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[_dafny.ZERO]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[_dafny.ONE]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(2)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(3)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(4)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(5)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(6)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(7)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(8)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(9)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(10)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(11)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(12)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(13)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(14)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(15)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(16)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(17)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(18)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v41)[new BigNumber(19)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[_dafny.ZERO]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[_dafny.ONE]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(2)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(3)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(4)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(5)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(6)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(7)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(8)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(9)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(10)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(11)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(12)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(13)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(14)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(15)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(16)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(17)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(18)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_191_v42).dtor_cf4)[new BigNumber(19)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_192_v43));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_193_v44).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_194_v45).dtor_cf22).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v107).dtor_cf50));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_303_v108)[_dafny.ZERO]).dtor_cf50));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_303_v108)[_dafny.ONE]).dtor_cf50));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_303_v108)[new BigNumber(2)]).dtor_cf50));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_303_v108)[new BigNumber(3)]).dtor_cf50));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_303_v108)[new BigNumber(4)]).dtor_cf50));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_303_v108)[new BigNumber(5)]).dtor_cf50));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_334_v130).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v131).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[_dafny.ZERO]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[_dafny.ONE]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(2)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(3)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(4)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(5)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(6)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(7)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(8)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(9)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(10)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(11)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(12)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(13)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(14)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(15)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(16)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(17)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(18)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(19)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(20)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(21)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(22)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(23)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(24)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(25)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v186)[new BigNumber(26)]).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_428_v187).dtor_cf23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_428_v187).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_428_v187).dtor_cf25, _dafny.Seq.of(new BigNumber(-496), new BigNumber(212), new BigNumber(212), new BigNumber(212), new BigNumber(-496), new BigNumber(212), new BigNumber(212)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_428_v187).dtor_cf26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_428_v187).dtor_cf27).equals(_dafny.Set.fromElements(new BigNumber(212)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_449_v199).equals(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("y"), _dafny.Seq.UnicodeFromString("r")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.ZERO;
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
    static create_DC2(cf2) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D1(2);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf2 === other.cf2;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf1 === other.cf1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(false);
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
    static create_DC5(cf5, cf6, cf7, cf8) {
      let $dt = new D2(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC6(cf9, cf10, cf11, cf12) {
      let $dt = new D2(1);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC4(cf4) {
      let $dt = new D2(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC7(cf13) {
      let $dt = new D2(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get is_DC7() { return this.$tag === 3; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && this.cf11 === other.cf11 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf4 === other.cf4;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC5(false, _dafny.ZERO, [], _dafny.MultiSet.Empty);
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
    static create_DC8(cf14) {
      let $dt = new D3(0);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14);
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
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf16) {
      let $dt = new D4(0);
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC9(cf15) {
      let $dt = new D4(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC11(cf17) {
      let $dt = new D4(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf16 === other.cf16;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf15 === other.cf15;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC10([]);
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
    static create_DC13(cf19, cf20, cf21) {
      let $dt = new D5(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC12(cf18) {
      let $dt = new D5(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC12" + "(" + this.cf18.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20) && this.cf21 === other.cf21;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC15(cf23, cf24, cf25, cf26, cf27) {
      let $dt = new D6(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC16(cf28, cf29, cf30, cf31, cf32) {
      let $dt = new D6(1);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC17(cf33, cf34, cf35, cf36) {
      let $dt = new D6(2);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC14(cf22) {
      let $dt = new D6(3);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get is_DC14() { return this.$tag === 3; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23) && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25) && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf28, other.cf28) && this.cf29 === other.cf29 && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf33, other.cf33) && this.cf34 === other.cf34 && this.cf35 === other.cf35 && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf22 === other.cf22;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC15(_dafny.ZERO, _dafny.ZERO, _dafny.Seq.of(), _dafny.ZERO, _dafny.Set.Empty);
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
    static create_DC19(cf38, cf39, cf40) {
      let $dt = new D7(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC20(cf41, cf42, cf43) {
      let $dt = new D7(1);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC18(cf37) {
      let $dt = new D7(2);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC21(cf44) {
      let $dt = new D7(3);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get is_DC21() { return this.$tag === 3; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38) && this.cf39 === other.cf39 && this.cf40 === other.cf40;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41) && _dafny.areEqual(this.cf42, other.cf42) && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC19(_dafny.Set.Empty, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC23(cf46, cf47, cf48) {
      let $dt = new D8(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC22(cf45) {
      let $dt = new D8(1);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf46, other.cf46) && this.cf47 === other.cf47 && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf45 === other.cf45;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23(_dafny.ZERO, [], _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC25(cf50) {
      let $dt = new D9(0);
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC26(cf51, cf52, cf53, cf54) {
      let $dt = new D9(1);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC27(cf55) {
      let $dt = new D9(2);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC24(cf49) {
      let $dt = new D9(3);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get is_DC24() { return this.$tag === 3; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf51 === other.cf51 && _dafny.areEqual(this.cf52, other.cf52) && this.cf53 === other.cf53 && this.cf54 === other.cf54;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf49 === other.cf49;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC25(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC29() {
      let $dt = new D10(0);
      return $dt;
    }
    static create_DC28(cf56) {
      let $dt = new D10(1);
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC30(cf57) {
      let $dt = new D10(2);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29";
      } else if (this.$tag === 1) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC30" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf56, other.cf56);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf57, other.cf57);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC29();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32(cf59) {
      let $dt = new D11(0);
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC31(cf58) {
      let $dt = new D11(1);
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC33(cf60) {
      let $dt = new D11(2);
      $dt.cf60 = cf60;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC33() { return this.$tag === 2; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf60() { return this.cf60; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC33" + "(" + _dafny.toString(this.cf60) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf59, other.cf59);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf60, other.cf60);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC32(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC34(cf61) {
      let $dt = new D12(0);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf61) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf61, other.cf61);
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
          return D12.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = false;
      this.f2 = _dafny.ZERO;
      this.f3 = false;
      this.f8 = _dafny.MultiSet.Empty;
      this.f9 = _dafny.ZERO;
      this.f10 = false;
      this.f11 = [];
      this._f1 = _dafny.Seq.of();
      this._f4 = false;
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f6 = _dafny.ZERO;
      this._f7 = [];
      this._f12 = _dafny.Map.Empty;
      this._f13 = _dafny.Seq.UnicodeFromString("");
      this._f14 = false;
      this._f15 = _dafny.Map.Empty;
      this._f16 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this).f9 = f9;
      (_this).f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f17 = _dafny.ZERO;
      this._f18 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f17, f18) {
      let _this = this;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      return;
    }
    fm2(globalState) {
      let _this = this;
      let _source7 = (_this).f17;
      let _476___mcc_h0 = _source7;
      let _477_cf0 = _476___mcc_h0;
      return (_this).f17;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f20 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [];
    }
    __ctor(f20) {
      let _this = this;
      (_this)._f20 = f20;
      return;
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.MultiSet.Empty;
      let _478_i0;
      _478_i0 = _dafny.ZERO;
      L0: {
        while (p0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_478_i0)) {
              break L0;
            }
            _478_i0 = (_478_i0).plus(_dafny.ONE);
            let _479_v0;
            let _nw81 = Array((new BigNumber(27)).toNumber()).fill(false);
            _479_v0 = _nw81;
            let _480_v1;
            _480_v1 = new BigNumber(990);
            let _481_v2;
            _481_v2 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (_479_v0) : (_479_v0)),_480_v1);
            let _482_v3;
            _482_v3 = _dafny.Seq.of((_module.__default.fm5((_dafny.ZERO).minus(_480_v1), new BigNumber(121), globalState)).dtor_cf2);
            (globalState).f2 = (((_481_v2).contains(_479_v0)) ? ((_481_v2).get(_479_v0)) : ((_dafny.ZERO).minus(new BigNumber((_482_v3).length))));
            if (p0) {
              let _483_v4;
              _483_v4 = _dafny.Map.Empty.slice().updateUnsafe(_480_v1,(_480_v1).minus(_480_v1));
              _483_v4 = (_483_v4).update(_480_v1, (new BigNumber(304)).minus((_dafny.ZERO).minus(_480_v1)));
              let _484_v5;
              let _485_v6;
              let _486_v7;
              let _487_v8;
              let _out18;
              let _out19;
              let _out20;
              let _out21;
              let _outcollector8 = (_this).m4(globalState);
              _out18 = _outcollector8[0];
              _out19 = _outcollector8[1];
              _out20 = _outcollector8[2];
              _out21 = _outcollector8[3];
              _484_v5 = _out18;
              _485_v6 = _out19;
              _486_v7 = _out20;
              _487_v8 = _out21;
              (globalState).f10 = !(function () {
                let _coll16 = new _dafny.Set();
                for (const _compr_16 of _dafny.IntegerRange(new BigNumber(959), new BigNumber(519))) {
                  let _488_v9 = _compr_16;
                  if (((new BigNumber(959)).isLessThanOrEqualTo(_488_v9)) && ((_488_v9).isLessThan(new BigNumber(519)))) {
                    _coll16.add((_488_v9).plus(_485_v6));
                  }
                }
                return _coll16;
              }()).contains(_480_v1);
              _480_v1 = _485_v6;
              let _489_v10;
              _489_v10 = _module.D1.create_DC1(_479_v0);
              let _490_v11;
              _490_v11 = _module.D1.create_DC3(_489_v10);
              let _491_v12;
              _491_v12 = _module.D1.create_DC1(_479_v0);
              let _492_v13;
              let _nw82 = Array((new BigNumber(15)).toNumber());
              _nw82[(_dafny.ZERO).toNumber()] = _491_v12;
              _nw82[(_dafny.ONE).toNumber()] = _491_v12;
              _nw82[(new BigNumber(2)).toNumber()] = _491_v12;
              _nw82[(new BigNumber(3)).toNumber()] = _491_v12;
              _nw82[(new BigNumber(4)).toNumber()] = _491_v12;
              _nw82[(new BigNumber(5)).toNumber()] = _491_v12;
              _nw82[(new BigNumber(6)).toNumber()] = _491_v12;
              _nw82[(new BigNumber(7)).toNumber()] = _module.D1.create_DC1(_479_v0);
              _nw82[(new BigNumber(8)).toNumber()] = _491_v12;
              _nw82[(new BigNumber(9)).toNumber()] = _491_v12;
              _nw82[(new BigNumber(10)).toNumber()] = _491_v12;
              _nw82[(new BigNumber(11)).toNumber()] = _491_v12;
              _nw82[(new BigNumber(12)).toNumber()] = _491_v12;
              _nw82[(new BigNumber(13)).toNumber()] = _491_v12;
              _nw82[(new BigNumber(14)).toNumber()] = _491_v12;
              _492_v13 = _nw82;
              let _493_v14;
              _493_v14 = _dafny.Seq.of(_492_v13, _492_v13);
              let _494_v15;
              let _nw83 = Array((new BigNumber(19)).toNumber());
              _nw83[(_dafny.ZERO).toNumber()] = _492_v13;
              _nw83[(_dafny.ONE).toNumber()] = _492_v13;
              _nw83[(new BigNumber(2)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(3)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(4)).toNumber()] = ((p0) ? (_492_v13) : (_492_v13));
              _nw83[(new BigNumber(5)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(6)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(7)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(8)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(9)).toNumber()] = ((_486_v7) ? (_492_v13) : (_492_v13));
              _nw83[(new BigNumber(10)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(11)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(12)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(13)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(14)).toNumber()] = ((p0) ? (_492_v13) : (_492_v13));
              _nw83[(new BigNumber(15)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(16)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(17)).toNumber()] = _492_v13;
              _nw83[(new BigNumber(18)).toNumber()] = (_493_v14)[_module.__default.safeIndex(_480_v1, new BigNumber((_493_v14).length))];
              _494_v15 = _nw83;
              let _495_v16;
              _495_v16 = _dafny.Map.Empty.slice().updateUnsafe(_485_v6,_484_v5);
              let _496_v17;
              _496_v17 = _dafny.Seq.UnicodeFromString("k");
              let _497_v18;
              _497_v18 = _dafny.Set.fromElements((_this).f20);
              let _498_v19;
              let _nw84 = Array((new BigNumber(12)).toNumber());
              _nw84[(_dafny.ZERO).toNumber()] = _480_v1;
              _nw84[(_dafny.ONE).toNumber()] = _485_v6;
              _nw84[(new BigNumber(2)).toNumber()] = _485_v6;
              _nw84[(new BigNumber(3)).toNumber()] = _480_v1;
              _nw84[(new BigNumber(4)).toNumber()] = new BigNumber(840);
              _nw84[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_480_v1));
              _nw84[(new BigNumber(6)).toNumber()] = new BigNumber((_483_v4).length);
              _nw84[(new BigNumber(7)).toNumber()] = new BigNumber((_496_v17).length);
              _nw84[(new BigNumber(8)).toNumber()] = new BigNumber((_497_v18).length);
              _nw84[(new BigNumber(9)).toNumber()] = _485_v6;
              _nw84[(new BigNumber(10)).toNumber()] = _485_v6;
              _nw84[(new BigNumber(11)).toNumber()] = _module.__default.fm0(globalState);
              _498_v19 = _nw84;
              let _499_v20;
              _499_v20 = _dafny.Map.Empty.slice().updateUnsafe((((_495_v16).contains(_485_v6)) ? ((_495_v16).get(_485_v6)) : (_484_v5)),_498_v19);
              let _rhs85 = _490_v11;
              let _rhs86 = (_499_v20).contains(!(_485_v6).isEqualTo(_480_v1));
              let _rhs87 = _494_v15;
              let _rhs88 = (_485_v6).isEqualTo(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_480_v1), _480_v1));
              let _rhs89 = _479_v0;
              _490_v11 = _rhs85;
              _486_v7 = _rhs86;
              _494_v15 = _rhs87;
              _484_v5 = _rhs88;
              _479_v0 = _rhs89;
            } else {
              let _500_v21;
              _500_v21 = _dafny.Seq.UnicodeFromString("dcfqu");
              let _rhs90 = _500_v21;
              let _rhs91 = ((p0) ? ((_480_v1).minus(_480_v1)) : (_module.__default.fm0(globalState)));
              let _lhs59 = globalState;
              _500_v21 = _rhs90;
              _lhs59.f2 = _rhs91;
              let _501_v22;
              _501_v22 = _dafny.Map.Empty.slice().updateUnsafe(_479_v0,p0);
              (globalState).f0 = !((_501_v22).Merge(_501_v22)).equals(_501_v22);
              let _502_v23;
              _502_v23 = (_dafny.ZERO).minus(_module.__default.fm0(globalState));
              let _503_v24;
              let _nw85 = new _module.C0();
              _nw85.__ctor(_502_v23, _480_v1);
              _503_v24 = _nw85;
              let _504_v25;
              let _nw86 = new _module.C0();
              _nw86.__ctor(_502_v23, _480_v1);
              _504_v25 = _nw86;
              (globalState).f2 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_480_v1));
            }
            if ((_module.__default.safeModuloInt(new BigNumber(-85), (_dafny.ZERO).minus(new BigNumber(-525)))).isLessThan(_480_v1)) {
              (globalState).f10 = (_482_v3)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_480_v1, (_dafny.ZERO).minus(_480_v1)), new BigNumber((_482_v3).length))];
              let _index61 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_479_v0).length));
              (_479_v0)[_index61] = p0;
              let _505_v26;
              _505_v26 = _dafny.Seq.UnicodeFromString("t");
              let _506_v27;
              _506_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("nogbhr"), _module.__default.safeIndex(_480_v1, new BigNumber((_dafny.Seq.UnicodeFromString("nogbhr")).length)), (_this).f20), _dafny.Seq.update(_module.__default.fm6(_module.__default.fm1(p0, _505_v26, globalState), globalState), _module.__default.safeIndex(_480_v1, new BigNumber((_module.__default.fm6(_module.__default.fm1(p0, _505_v26, globalState), globalState)).length)), _module.__default.fm7((_this).f20, p0, globalState)))).length),(_dafny.Set.fromElements(_480_v1)).IsDisjointFrom(_module.__default.fm8(new BigNumber((_dafny.MultiSet.FromArray(_505_v26)).cardinality()), _480_v1, globalState)));
              let _index62 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_479_v0).length));
              (_479_v0)[_index62] = p0;
              let _507_v28;
              _507_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
              let _508_v29;
              _508_v29 = _dafny.Map.Empty.slice().updateUnsafe(_507_v28,(_dafny.ZERO).minus(_480_v1));
              let _509_v30;
              _509_v30 = _dafny.MultiSet.fromElements(_480_v1);
              let _510_v31;
              _510_v31 = _dafny.Set.fromElements(_480_v1);
              let _511_v32;
              _511_v32 = _dafny.Seq.of(new BigNumber((_510_v31).length), _480_v1, new BigNumber((_482_v3).length), _480_v1, _480_v1);
              let _512_v34;
              _512_v34 = _dafny.Map.Empty.slice().updateUnsafe((((_509_v30).contains(_480_v1)) ? ((_509_v30).get(_480_v1)) : (new BigNumber((_511_v32).length))),function () {
                let _coll17 = new _dafny.Map();
                for (const _compr_17 of (_506_v27).Keys.Elements) {
                  let _513_v33 = _compr_17;
                  if ((_506_v27).contains(_513_v33)) {
                    _coll17.push([(_513_v33).multipliedBy(new BigNumber(214)),p0]);
                  }
                }
                return _coll17;
              }());
              let _514_v35;
              _514_v35 = _dafny.Set.fromElements((((_507_v28).contains(p0)) ? ((_507_v28).get(p0)) : (p0)), p0, p0);
              let _515_v36;
              _515_v36 = _module.D1.create_DC2(p0);
              let _index63 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_479_v0).length));
              let _index64 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_479_v0).length));
              let _rhs92 = !(_508_v29).contains(_507_v28);
              let _rhs93 = ((p0) ? (_506_v27) : ((((_512_v34).contains(_480_v1)) ? ((_512_v34).get(_480_v1)) : (_506_v27))));
              let _rhs94 = p0;
              let _rhs95 = !((_514_v35).Union(_dafny.Set.fromElements((_515_v36).dtor_cf2, p0, p0))).contains(p0);
              let _rhs96 = _module.__default.fm0(globalState);
              let _lhs60 = _479_v0;
              let _lhs61 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_479_v0).length));
              let _lhs62 = _479_v0;
              let _lhs63 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_479_v0).length));
              let _lhs64 = globalState;
              let _lhs65 = globalState;
              _lhs60[_lhs61] = _rhs92;
              _506_v27 = _rhs93;
              _lhs62[_lhs63] = _rhs94;
              _lhs64.f0 = _rhs95;
              _lhs65.f9 = _rhs96;
              (globalState).f0 = (_515_v36).dtor_cf2;
              let _516_v37;
              _516_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,p0);
              _516_v37 = (_516_v37).update((_this).f20, p0);
              let _517_v38;
              _517_v38 = _dafny.Set.fromElements((_506_v27).Merge(_506_v27));
              _517_v38 = _517_v38;
            } else {
              let _518_v39;
              _518_v39 = _480_v1;
              let _519_v40;
              let _nw87 = new _module.C0();
              _nw87.__ctor(_518_v39, _480_v1);
              _519_v40 = _nw87;
              _519_v40 = _519_v40;
              (globalState).f0 = p0;
              (globalState).f9 = _480_v1;
              (globalState).f2 = _480_v1;
              let _520_v41;
              _520_v41 = _module.D1.create_DC1(_479_v0);
              let _521_v42;
              _521_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_520_v41);
              let _522_v43;
              _522_v43 = _dafny.Seq.UnicodeFromString("im");
              let _523_v44;
              let _nw88 = Array((new BigNumber(2)).toNumber());
              _nw88[(_dafny.ZERO).toNumber()] = _522_v43;
              _nw88[(_dafny.ONE).toNumber()] = _522_v43;
              _523_v44 = _nw88;
              let _index65 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_523_v44).length));
              (_523_v44)[_index65] = _522_v43;
              let _index66 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_523_v44).length));
              let _rhs97 = p0;
              let _rhs98 = _module.__default.safeDivisionInt(_480_v1, new BigNumber(256));
              let _rhs99 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_520_v41);
              let _rhs100 = _module.__default.fm6(p0, globalState);
              let _lhs66 = globalState;
              let _lhs67 = globalState;
              let _lhs68 = _523_v44;
              let _lhs69 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_523_v44).length));
              _lhs66.f3 = _rhs97;
              _lhs67.f9 = _rhs98;
              _521_v42 = _rhs99;
              _lhs68[_lhs69] = _rhs100;
            }
            let _524_v45;
            _524_v45 = _480_v1;
            let _525_v46;
            let _nw89 = new _module.C0();
            _nw89.__ctor(_524_v45, _480_v1);
            _525_v46 = _nw89;
          }
        }
      }
      (globalState).f3 = p0;
      let _526_v47;
      _526_v47 = new BigNumber(825);
      r0 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_526_v47), new BigNumber(127));
      let _527_i1;
      _527_i1 = _dafny.ZERO;
      L1: {
        while (!(p0)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_527_i1)) {
              break L1;
            }
            _527_i1 = (_527_i1).plus(_dafny.ONE);
            let _528_v48;
            _528_v48 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
            _528_v48 = (_528_v48).Merge(_528_v48);
            (globalState).f10 = (!(p0)) && ((new BigNumber((_dafny.Seq.of(p0, p0, p0)).length)).isLessThan((_dafny.ZERO).minus(new BigNumber(-389))));
            let _529_v49;
            _529_v49 = _dafny.Seq.UnicodeFromString("tht");
            _529_v49 = _dafny.Seq.Concat(_529_v49, _529_v49);
            let _530_v50;
            _530_v50 = _526_v47;
            let _531_v51;
            _531_v51 = _dafny.Map.Empty.slice().updateUnsafe((_530_v50),_module.__default.fm1(p0, _529_v49, globalState));
            let _532_v52;
            let _533_v53;
            let _out22;
            let _out23;
            let _outcollector9 = _module.__default.m0(p0, _531_v51, globalState);
            _out22 = _outcollector9[0];
            _out23 = _outcollector9[1];
            _532_v52 = _out22;
            _533_v53 = _out23;
          }
        }
      }
      let _534_v54;
      let _nw90 = Array((new BigNumber(26)).toNumber()).fill(false);
      _534_v54 = _nw90;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_534_v54).length))) {
        let _535_i2 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_535_i2)) && ((_535_i2).isLessThan(new BigNumber((_534_v54).length))))) {
          (_534_v54)[(_535_i2)] = !(p0);
        }
      }
      (globalState).f2 = _526_v47;
      let _536_v55;
      _536_v55 = _dafny.Seq.UnicodeFromString("iakx");
      r0 = ((p0) ? (new BigNumber((_536_v55).length)) : (_526_v47));
      let _537_v56;
      _537_v56 = _dafny.MultiSet.fromElements(p0);
      let _538_v57;
      _538_v57 = _dafny.Set.fromElements(_526_v47);
      let _539_v58;
      _539_v58 = _dafny.Seq.of(p0, p0, p0, _module.__default.fm1(p0, _536_v55, globalState), p0);
      let _540_v59;
      _540_v59 = _dafny.MultiSet.fromElements(_537_v56, _537_v56, ((p0) ? (_module.__default.fm9(new BigNumber((_538_v57).length), p0, p0, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_539_v58).length))), globalState)) : (_dafny.MultiSet.FromArray(_539_v58))), (_537_v56).update(p0, _module.__default.abs(new BigNumber((_538_v57).length))));
      r1 = _540_v59;
      return [r0, r1];
    }
    m4(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = false;
      let _541_v0;
      _541_v0 = false;
      if (_541_v0) {
        let _542_v1;
        _542_v1 = new BigNumber(785);
        let _543_v2;
        _543_v2 = _dafny.Seq.of(_541_v0);
        let _544_v3;
        _544_v3 = _dafny.MultiSet.fromElements(_541_v0, (_542_v1).isEqualTo(_542_v1), (_543_v2)[_module.__default.safeIndex(_542_v1, new BigNumber((_543_v2).length))]);
        let _545_v4;
        _545_v4 = _dafny.Seq.of(_542_v1, (_dafny.ZERO).minus(_542_v1));
        let _546_v5;
        _546_v5 = _dafny.Seq.UnicodeFromString("hfyg");
        let _rhs101 = (_dafny.MultiSet.FromArray(_dafny.Seq.update(((_541_v0) ? (_dafny.Seq.of(_541_v0, _541_v0)) : (_543_v2)), _module.__default.safeIndex(new BigNumber(947), new BigNumber((((_541_v0) ? (_dafny.Seq.of(_541_v0, _541_v0)) : (_543_v2))).length)), _541_v0))).Union(_544_v3);
        let _rhs102 = _541_v0;
        let _rhs103 = (_module.__default.safeDivisionInt(_542_v1, new BigNumber((_545_v4).length))).minus((_dafny.ZERO).minus((_542_v1).plus((_dafny.ZERO).minus(new BigNumber((_546_v5).length)))));
        let _lhs70 = globalState;
        let _lhs71 = globalState;
        _544_v3 = _rhs101;
        _lhs70.f3 = _rhs102;
        _lhs71.f2 = _rhs103;
        let _547_v6;
        let _nw91 = Array((new BigNumber(16)).toNumber()).fill(_module.D1.Default());
        _547_v6 = _nw91;
        let _548_v7;
        _548_v7 = _module.D2.create_DC4(_547_v6);
        _547_v6 = (_548_v7).dtor_cf4;
        let _549_v8;
        let _nw92 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _549_v8 = _nw92;
        let _index67 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_549_v8).length));
        (_549_v8)[_index67] = (_542_v1).minus(_542_v1);
        let _index68 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_549_v8).length));
        (_549_v8)[_index68] = _542_v1;
        (globalState).f0 = false;
        let _550_v9;
        let _nw93 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _550_v9 = _nw93;
        let _551_v10;
        _551_v10 = _dafny.Seq.of(_550_v9, _550_v9);
        let _552_v11;
        _552_v11 = _dafny.Seq.of(_550_v9, (_551_v10)[_module.__default.safeIndex(new BigNumber(-315), new BigNumber((_551_v10).length))]);
        let _553_v12;
        _553_v12 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,_550_v9);
        let _554_v13;
        let _nw94 = Array((new BigNumber(28)).toNumber());
        _nw94[(_dafny.ZERO).toNumber()] = _550_v9;
        _nw94[(_dafny.ONE).toNumber()] = _550_v9;
        _nw94[(new BigNumber(2)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(3)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(4)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(5)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(6)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(7)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(8)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(9)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(10)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(11)).toNumber()] = (_552_v11)[_module.__default.safeIndex(_542_v1, new BigNumber((_552_v11).length))];
        _nw94[(new BigNumber(12)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(13)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(14)).toNumber()] = ((true) ? (_550_v9) : (_550_v9));
        _nw94[(new BigNumber(15)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(16)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(17)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(18)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(19)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(20)).toNumber()] = ((true) ? ((_552_v11)[_module.__default.safeIndex(_542_v1, new BigNumber((_552_v11).length))]) : (_550_v9));
        _nw94[(new BigNumber(21)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(22)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(23)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(24)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(25)).toNumber()] = _550_v9;
        _nw94[(new BigNumber(26)).toNumber()] = (((_553_v12).contains(_541_v0)) ? ((_553_v12).get(_541_v0)) : (_550_v9));
        _nw94[(new BigNumber(27)).toNumber()] = _550_v9;
        _554_v13 = _nw94;
        let _index69 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_554_v13).length));
        (_554_v13)[_index69] = _550_v9;
        let _index70 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_554_v13).length));
        (_554_v13)[_index70] = _550_v9;
      } else {
        let _555_v14;
        _555_v14 = new BigNumber(-271);
        let _556_v15;
        _556_v15 = _dafny.Seq.of(_555_v14, _555_v14, _555_v14);
        (globalState).f0 = ((_556_v15)[_module.__default.safeIndex(_555_v14, new BigNumber((_556_v15).length))]).isLessThanOrEqualTo(_module.__default.safeModuloInt(_555_v14, _555_v14));
        let _557_v16;
        _557_v16 = _dafny.MultiSet.fromElements(_541_v0, _541_v0);
        let _558_v17;
        _558_v17 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,_557_v16);
        let _559_v18;
        _559_v18 = new BigNumber((_558_v17).length);
        let _source8 = _559_v18;
        let _560___mcc_h0 = _source8;
        let _561_cf0 = _560___mcc_h0;
        let _562_v19;
        let _nw95 = new _module.C0();
        _nw95.__ctor(_559_v18, new BigNumber(90));
        _562_v19 = _nw95;
        let _563_v20;
        _563_v20 = _dafny.Map.Empty.slice().updateUnsafe(!(_541_v0),new BigNumber(-251));
        let _564_v21;
        _564_v21 = _dafny.Seq.of(_563_v20);
        _564_v21 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(787)), ((_565_v20) => function (_566_i0) {
          return _565_v20;
        })(_563_v20));
        let _567_v22;
        _567_v22 = new _dafny.CodePoint('h'.codePointAt(0));
        let _568_v23;
        _568_v23 = _dafny.Seq.UnicodeFromString("nntcghypc");
        let _569_v24;
        let _nw96 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _569_v24 = _nw96;
        let _index71 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_569_v24).length));
        (_569_v24)[_index71] = _561_cf0;
        let _570_v25;
        let _nw97 = Array((new BigNumber(3)).toNumber()).fill(false);
        _570_v25 = _nw97;
        let _index72 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_570_v25).length));
        (_570_v25)[_index72] = (_557_v16).IsSubsetOf(_557_v16);
        let _571_v26;
        _571_v26 = _dafny.Seq.of(_570_v25);
        let _572_v27;
        _572_v27 = _dafny.Map.Empty.slice().updateUnsafe(_571_v26,_567_v22);
        let _index73 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_569_v24).length));
        let _index74 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_570_v25).length));
        let _rhs104 = (((_572_v27).contains(_571_v26)) ? ((_572_v27).get(_571_v26)) : (_567_v22));
        let _rhs105 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(597)), ((_573_v23, _574_cf0) => function (_575_i1) {
          return (_573_v23)[_module.__default.safeIndex(_574_cf0, new BigNumber((_573_v23).length))];
        })(_568_v23, _561_cf0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(657)), ((_576_v22) => function (_577_i2) {
          return _576_v22;
        })(_567_v22)));
        let _rhs106 = _561_cf0;
        let _rhs107 = !(_541_v0);
        let _lhs72 = _569_v24;
        let _lhs73 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_569_v24).length));
        let _lhs74 = _570_v25;
        let _lhs75 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_570_v25).length));
        _567_v22 = _rhs104;
        _568_v23 = _rhs105;
        _lhs72[_lhs73] = _rhs106;
        _lhs74[_lhs75] = _rhs107;
        _563_v20 = (_563_v20).update(_module.__default.fm1(_541_v0, _568_v23, globalState), (_569_v24)[_module.__default.safeIndex(new BigNumber(426), new BigNumber((_569_v24).length))]);
        let _578_v28;
        let _nw98 = new _module.C0();
        _nw98.__ctor(_559_v18, _555_v14);
        _578_v28 = _nw98;
        let _579_v29;
        _579_v29 = _dafny.MultiSet.fromElements(_578_v28);
        let _580_v30;
        _580_v30 = _579_v29;
        let _581_v31;
        _581_v31 = _dafny.Map.Empty.slice().updateUnsafe(_555_v14,_555_v14);
        let _582_v32;
        let _nw99 = Array((new BigNumber(25)).toNumber()).fill(false);
        _582_v32 = _nw99;
        let _583_v33;
        _583_v33 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm8(_555_v14, _555_v14, globalState)).length), _555_v14);
        let _584_v34;
        _584_v34 = _module.D2.create_DC5(_541_v0, new BigNumber((_581_v31).length), _582_v32, _583_v33);
        let _585_v35;
        _585_v35 = _dafny.Seq.of(_541_v0, !(_541_v0), _541_v0);
        let _586_v36;
        let _nw100 = Array((new BigNumber(16)).toNumber());
        _nw100[(_dafny.ZERO).toNumber()] = _dafny.Seq.of((_584_v34).dtor_cf5, _541_v0);
        _nw100[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_541_v0, _541_v0);
        _nw100[(new BigNumber(2)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(3)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(4)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(5)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(6)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(7)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(8)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(9)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(10)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(11)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(12)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(13)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(14)).toNumber()] = _585_v35;
        _nw100[(new BigNumber(15)).toNumber()] = _585_v35;
        _586_v36 = _nw100;
        let _587_v37;
        _587_v37 = _dafny.Map.Empty.slice().updateUnsafe(((_580_v30)).Union(_579_v29),_586_v36);
        _587_v37 = (_587_v37).update(_579_v29, _586_v36);
        let _588_v38;
        _588_v38 = _dafny.Seq.UnicodeFromString("eancxn");
        _588_v38 = _588_v38;
        (globalState).f0 = ((_555_v14).multipliedBy(_555_v14)).isLessThan((_dafny.ZERO).minus(_555_v14));
      }
      let _589_i3;
      _589_i3 = _dafny.ZERO;
      L2: {
        while (true) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_589_i3)) {
              break L2;
            }
            _589_i3 = (_589_i3).plus(_dafny.ONE);
            let _590_v39;
            _590_v39 = _dafny.Seq.of(_541_v0);
            let _591_v40;
            _591_v40 = _dafny.Seq.UnicodeFromString("htdkpplcq");
            let _592_v41;
            let _init12 = function (_593_i4) {
              return (_593_i4).plus(new BigNumber(221));
            };
            let _nw101 = Array((new BigNumber(22)).toNumber());
            for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw101.length); _i0_12++) {
              _nw101[_i0_12] = _init12(new BigNumber(_i0_12));
            }
            _592_v41 = _nw101;
            let _594_v42;
            _594_v42 = _dafny.Map.Empty.slice().updateUnsafe(true,_592_v41);
            if ((_590_v39)[_module.__default.safeIndex(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_591_v40).length)), new BigNumber((_594_v42).length)), new BigNumber((_590_v39).length))]) {
              let _595_v43;
              _595_v43 = _dafny.Set.fromElements(_541_v0);
              let _596_v44;
              _596_v44 = new BigNumber(192);
              let _597_v45;
              _597_v45 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,_596_v44);
              let _598_v46;
              _598_v46 = new BigNumber((_597_v45).length);
              let _599_v47;
              _599_v47 = _dafny.Map.Empty.slice().updateUnsafe(_595_v43,(_590_v39)[_module.__default.safeIndex((_598_v46), new BigNumber((_590_v39).length))]);
              _599_v47 = _599_v47;
              (globalState).f3 = (_541_v0) || (_541_v0);
              r3 = _541_v0;
              (globalState).f10 = _541_v0;
              let _600_v48;
              _600_v48 = _dafny.Seq.of(_596_v44, _596_v44);
              let _601_v49;
              _601_v49 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements((_this).f20, (_591_v40)[_module.__default.safeIndex(_596_v44, new BigNumber((_591_v40).length))])).length));
              let _602_v50;
              let _nw102 = Array((_dafny.ONE).toNumber()).fill(false);
              _602_v50 = _nw102;
              let _603_v51;
              _603_v51 = _module.D2.create_DC5(_541_v0, _596_v44, _602_v50, _dafny.MultiSet.fromElements(_596_v44));
              let _604_v52;
              let _nw103 = Array((new BigNumber(5)).toNumber());
              _nw103[(_dafny.ZERO).toNumber()] = (_601_v49).IsSubsetOf(_dafny.MultiSet.FromArray(_600_v48));
              _nw103[(_dafny.ONE).toNumber()] = (_590_v39)[_module.__default.safeIndex((_dafny.ZERO).minus(_596_v44), new BigNumber((_590_v39).length))];
              _nw103[(new BigNumber(2)).toNumber()] = _541_v0;
              _nw103[(new BigNumber(3)).toNumber()] = _541_v0;
              _nw103[(new BigNumber(4)).toNumber()] = (_603_v51).dtor_cf5;
              _604_v52 = _nw103;
              let _index75 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_604_v52).length));
              (_604_v52)[_index75] = _541_v0;
              let _index76 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_604_v52).length));
              (_604_v52)[_index76] = _541_v0;
            } else {
              _591_v40 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(764)), function (_605_i5) {
                return (_this).f20;
              }), _591_v40);
              let _606_v53;
              _606_v53 = _dafny.Map.Empty.slice().updateUnsafe(_590_v39,new BigNumber(555));
              let _607_v54;
              _607_v54 = new BigNumber((_606_v53).length);
              let _608_v55;
              let _nw104 = new _module.C0();
              _nw104.__ctor(_607_v54, new BigNumber((_591_v40).length));
              _608_v55 = _nw104;
              let _609_v56;
              _609_v56 = _dafny.Seq.of(_608_v55, _608_v55, _608_v55);
              let _610_v57;
              _610_v57 = new BigNumber(30);
              let _611_v58;
              _611_v58 = _module.D4.create_DC9(_608_v55);
              let _612_v59;
              let _nw105 = Array((new BigNumber(24)).toNumber());
              _nw105[(_dafny.ZERO).toNumber()] = _608_v55;
              _nw105[(_dafny.ONE).toNumber()] = _608_v55;
              _nw105[(new BigNumber(2)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(3)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(4)).toNumber()] = (_609_v56)[_module.__default.safeIndex(_610_v57, new BigNumber((_609_v56).length))];
              _nw105[(new BigNumber(5)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(6)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(7)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(8)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(9)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(10)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(11)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(12)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(13)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(14)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(15)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(16)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(17)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(18)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(19)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(20)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(21)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(22)).toNumber()] = _608_v55;
              _nw105[(new BigNumber(23)).toNumber()] = (_611_v58).dtor_cf15;
              _612_v59 = _nw105;
              let _index77 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_612_v59).length));
              (_612_v59)[_index77] = _608_v55;
              let _index78 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_612_v59).length));
              (_612_v59)[_index78] = _608_v55;
              (globalState).f2 = _610_v57;
              let _613_v60;
              let _nw106 = Array((new BigNumber(17)).toNumber()).fill(false);
              _613_v60 = _nw106;
              let _614_v61;
              _614_v61 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,_613_v60);
              let _615_v62;
              let _nw107 = Array((new BigNumber(17)).toNumber());
              _nw107[(_dafny.ZERO).toNumber()] = _613_v60;
              _nw107[(_dafny.ONE).toNumber()] = _613_v60;
              _nw107[(new BigNumber(2)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(3)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(4)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(5)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(6)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(7)).toNumber()] = (((_614_v61).contains(_541_v0)) ? ((_614_v61).get(_541_v0)) : (_613_v60));
              _nw107[(new BigNumber(8)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(9)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(10)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(11)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(12)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(13)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(14)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(15)).toNumber()] = _613_v60;
              _nw107[(new BigNumber(16)).toNumber()] = _613_v60;
              _615_v62 = _nw107;
              let _index79 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_615_v62).length));
              (_615_v62)[_index79] = _613_v60;
              let _616_v63;
              _616_v63 = _dafny.MultiSet.fromElements(_610_v57, _610_v57, _610_v57);
              let _index80 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_615_v62).length));
              let _rhs108 = new BigNumber((_616_v63).cardinality());
              let _rhs109 = _613_v60;
              let _lhs76 = _615_v62;
              let _lhs77 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_615_v62).length));
              _610_v57 = _rhs108;
              _lhs76[_lhs77] = _rhs109;
              (globalState).f2 = _610_v57;
            }
            _541_v0 = _541_v0;
            let _617_v64;
            let _nw108 = Array((new BigNumber(28)).toNumber());
            _nw108[(_dafny.ZERO).toNumber()] = _541_v0;
            _nw108[(_dafny.ONE).toNumber()] = _541_v0;
            _nw108[(new BigNumber(2)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(3)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(4)).toNumber()] = _module.__default.fm1(true, _591_v40, globalState);
            _nw108[(new BigNumber(5)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(6)).toNumber()] = _module.__default.fm1(_541_v0, _591_v40, globalState);
            _nw108[(new BigNumber(7)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(8)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(9)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(10)).toNumber()] = !(_541_v0);
            _nw108[(new BigNumber(11)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(12)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(13)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(14)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(15)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(16)).toNumber()] = true;
            _nw108[(new BigNumber(17)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(18)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(19)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(20)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(21)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(22)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(23)).toNumber()] = false;
            _nw108[(new BigNumber(24)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(25)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(26)).toNumber()] = _541_v0;
            _nw108[(new BigNumber(27)).toNumber()] = _module.__default.fm1(_541_v0, _dafny.Seq.update(_591_v40, _module.__default.safeIndex(new BigNumber(628), new BigNumber((_591_v40).length)), (_this).f20), globalState);
            _617_v64 = _nw108;
            let _618_v65;
            _618_v65 = _module.D1.create_DC1(_617_v64);
            let _619_v66;
            _619_v66 = _dafny.Seq.of(_617_v64);
            let _620_v67;
            let _nw109 = Array((new BigNumber(25)).toNumber());
            _nw109[(_dafny.ZERO).toNumber()] = _617_v64;
            _nw109[(_dafny.ONE).toNumber()] = _617_v64;
            _nw109[(new BigNumber(2)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(3)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(4)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(5)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(6)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(7)).toNumber()] = ((_541_v0) ? (_617_v64) : (_617_v64));
            _nw109[(new BigNumber(8)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(9)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(10)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(11)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(12)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(13)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(14)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(15)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(16)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(17)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(18)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(19)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(20)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(21)).toNumber()] = (_618_v65).dtor_cf1;
            _nw109[(new BigNumber(22)).toNumber()] = _617_v64;
            _nw109[(new BigNumber(23)).toNumber()] = ((_541_v0) ? (_617_v64) : ((_619_v66)[_module.__default.safeIndex(new BigNumber((_590_v39).length), new BigNumber((_619_v66).length))]));
            _nw109[(new BigNumber(24)).toNumber()] = _617_v64;
            _620_v67 = _nw109;
            let _nw110 = Array((new BigNumber(5)).toNumber()).fill([]);
            _620_v67 = _nw110;
            let _621_v68;
            _621_v68 = new BigNumber(14);
            r2 = (new BigNumber((_590_v39).length)).isLessThan(_module.__default.safeDivisionInt(_621_v68, new BigNumber((_dafny.Set.fromElements(_591_v40, _591_v40)).length)));
          }
        }
      }
      let _622_v69;
      let _nw111 = Array((new BigNumber(3)).toNumber()).fill(false);
      _622_v69 = _nw111;
      let _623_v70;
      _623_v70 = _module.D1.create_DC1(_622_v69);
      let _624_v71;
      _624_v71 = _dafny.Map.Empty.slice().updateUnsafe(_623_v70,new BigNumber(421));
      let _625_v72;
      _625_v72 = new BigNumber(662);
      let _626_v73;
      _626_v73 = _dafny.Map.Empty.slice().updateUnsafe((((_624_v71).contains(_623_v70)) ? ((_624_v71).get(_623_v70)) : ((_dafny.ZERO).minus(_625_v72))),_625_v72);
      let _627_v74;
      _627_v74 = _dafny.Seq.of((_this).f20);
      _626_v73 = (_626_v73).update((new BigNumber((_627_v74).length)).minus(_625_v72), new BigNumber(-351));
      let _index81 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_622_v69).length));
      (_622_v69)[_index81] = _541_v0;
      let _628_v75;
      _628_v75 = _dafny.Seq.of(_625_v72);
      let _index82 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_622_v69).length));
      (_622_v69)[_index82] = !(false) || (_dafny.Seq.IsProperPrefixOf(_628_v75, _628_v75));
      let _629_v76;
      _629_v76 = _dafny.Seq.of((_622_v69)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_622_v69).length))]);
      let _630_v77;
      _630_v77 = _module.D1.create_DC2(!((_622_v69)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_622_v69).length))]) || ((_629_v76)[_module.__default.safeIndex(_625_v72, new BigNumber((_629_v76).length))]));
      _630_v77 = function (_pat_let25_0) {
        return function (_631_dt__update__tmp_h0) {
          return function (_pat_let26_0) {
            return function (_632_dt__update_hcf2_h0) {
              return _module.D1.create_DC2(_632_dt__update_hcf2_h0);
            }(_pat_let26_0);
          }(false);
        }(_pat_let25_0);
      }(_630_v77);
      let _633_v78;
      _633_v78 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_625_v72),(_622_v69)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_622_v69).length))]);
      let _634_v79;
      let _635_v80;
      let _out24;
      let _out25;
      let _outcollector10 = _module.__default.m0((((_629_v76)[_module.__default.safeIndex((_628_v75)[_module.__default.safeIndex(_625_v72, new BigNumber((_628_v75).length))], new BigNumber((_629_v76).length))]) ? ((_622_v69)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_622_v69).length))]) : ((_622_v69)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_622_v69).length))])), (_633_v78).update(_625_v72, _541_v0), globalState);
      _out24 = _outcollector10[0];
      _out25 = _outcollector10[1];
      _634_v79 = _out24;
      _635_v80 = _out25;
      r0 = !((_622_v69)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_622_v69).length))]);
      let _636_v81;
      _636_v81 = _module.D5.create_DC12(_627_v74);
      r1 = new BigNumber((_dafny.Seq.Concat((_636_v81).dtor_cf18, _dafny.Seq.Concat(_627_v74, _627_v74))).length);
      r2 = _module.__default.fm1(_541_v0, _627_v74, globalState);
      r3 = !(((_622_v69)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_622_v69).length))]) === ((_622_v69)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_622_v69).length))])) || (!(_541_v0));
      return [r0, r1, r2, r3];
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    m2(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _637_v0;
      _637_v0 = new BigNumber(793);
      (globalState).f3 = (_637_v0).isEqualTo(_637_v0);
      let _638_v1;
      _638_v1 = new BigNumber(84);
      let _639_v2;
      _639_v2 = new _dafny.CodePoint('f'.codePointAt(0));
      let _640_v3;
      let _nw112 = new _module.C1();
      _nw112.__ctor(_639_v2);
      _640_v3 = _nw112;
      let _641_v4;
      _641_v4 = _dafny.MultiSet.fromElements(_640_v3);
      let _642_v5;
      _642_v5 = _dafny.MultiSet.fromElements(new BigNumber(906));
      let _643_v6;
      let _nw113 = new _module.C0();
      _nw113.__ctor(_638_v1, (new BigNumber((_dafny.Set.fromElements(_641_v4)).length)).plus((((_642_v5).contains(_637_v0)) ? ((_642_v5).get(_637_v0)) : (new BigNumber(302)))));
      _643_v6 = _nw113;
      let _644_v7;
      _644_v7 = true;
      let _645_v8;
      _645_v8 = _dafny.Seq.UnicodeFromString("rfvwqmbea");
      let _646_v9;
      _646_v9 = _module.D5.create_DC13(_637_v0, _637_v0, _module.__default.fm1(!(_644_v7), _645_v8, globalState));
      (globalState).f3 = function (_source9) {
        if (_source9.is_DC13) {
          let _647___mcc_h0 = (_source9).cf19;
          let _648___mcc_h1 = (_source9).cf20;
          let _649___mcc_h2 = (_source9).cf21;
          let _650_cf21 = _649___mcc_h2;
          let _651_cf20 = _648___mcc_h1;
          let _652_cf19 = _647___mcc_h0;
          return false;
        } else {
          let _653___mcc_h3 = (_source9).cf18;
          let _654_cf18 = _653___mcc_h3;
          return false;
        }
      }(_646_v9);
      let _655_v10;
      let _nw114 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
      _655_v10 = _nw114;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_655_v10).length))) {
        let _656_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_656_i0)) && ((_656_i0).isLessThan(new BigNumber((_655_v10).length))))) {
          (_655_v10)[(_656_i0)] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_644_v7), _dafny.Seq.of(_644_v7)), _dafny.Seq.Concat(_dafny.Seq.of(_644_v7, _644_v7), _dafny.Seq.of(_644_v7)));
        }
      }
      let _657_v11;
      _657_v11 = _module.D5.create_DC12(_645_v8);
      let _pat_let_tv20 = _644_v7;
      if (function (_source10) {
        if (_source10.is_DC13) {
          let _658___mcc_h4 = (_source10).cf19;
          let _659___mcc_h5 = (_source10).cf20;
          let _660___mcc_h6 = (_source10).cf21;
          let _661_cf21 = _660___mcc_h6;
          let _662_cf20 = _659___mcc_h5;
          let _663_cf19 = _658___mcc_h4;
          return false;
        } else {
          let _664___mcc_h7 = (_source10).cf18;
          let _665_cf18 = _664___mcc_h7;
          return !(_pat_let_tv20);
        }
      }(_657_v11)) {
        _644_v7 = _644_v7;
        let _666_v12;
        _666_v12 = _dafny.Set.fromElements(!(_644_v7), _644_v7, _644_v7);
        let _667_v13;
        let _init13 = function (_668_i1) {
          return false;
        };
        let _nw115 = Array((new BigNumber(24)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw115.length); _i0_13++) {
          _nw115[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _667_v13 = _nw115;
        let _669_v14;
        _669_v14 = _dafny.Map.Empty.slice().updateUnsafe(((_644_v7) ? (new BigNumber((_666_v12).length)) : (_637_v0)),_667_v13);
        _669_v14 = (_669_v14).update(new BigNumber(377), _667_v13);
        if ((_637_v0).isLessThan((_637_v0).plus(_637_v0))) {
          (globalState).f2 = _637_v0;
          (globalState).f10 = !(_dafny.Seq.contains(_645_v8, ((_644_v7) ? (_639_v2) : ((_640_v3).f20))));
          (globalState).f3 = (_module.__default.safeDivisionInt(_637_v0, (_dafny.ZERO).minus(_637_v0))).isLessThan(_637_v0);
          r0 = _637_v0;
          let _670_v15;
          _670_v15 = _module.D6.create_DC14(_640_v3);
          let _671_v16;
          let _nw116 = Array((new BigNumber(9)).toNumber());
          _nw116[(_dafny.ZERO).toNumber()] = ((_644_v7) ? (_640_v3) : ((_670_v15).dtor_cf22));
          _nw116[(_dafny.ONE).toNumber()] = _640_v3;
          _nw116[(new BigNumber(2)).toNumber()] = _640_v3;
          _nw116[(new BigNumber(3)).toNumber()] = _640_v3;
          _nw116[(new BigNumber(4)).toNumber()] = _640_v3;
          _nw116[(new BigNumber(5)).toNumber()] = _640_v3;
          _nw116[(new BigNumber(6)).toNumber()] = _640_v3;
          _nw116[(new BigNumber(7)).toNumber()] = _640_v3;
          _nw116[(new BigNumber(8)).toNumber()] = _640_v3;
          _671_v16 = _nw116;
          _671_v16 = _671_v16;
        } else {
          let _672_v17;
          let _nw117 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _672_v17 = _nw117;
          let _index83 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_672_v17).length));
          (_672_v17)[_index83] = ((_644_v7) ? (_dafny.Seq.update(_645_v8, _module.__default.safeIndex(_637_v0, new BigNumber((_645_v8).length)), _639_v2)) : (_645_v8));
          let _673_v18;
          let _nw118 = Array((new BigNumber(23)).toNumber());
          _673_v18 = _nw118;
          let _674_v19;
          _674_v19 = _dafny.MultiSet.fromElements(_673_v18);
          let _index84 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_672_v17).length));
          let _rhs110 = ((_644_v7) ? (!(_644_v7)) : (_644_v7));
          let _rhs111 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(143)), ((_675_v3) => function (_676_i2) {
            return (_675_v3).f20;
          })(_640_v3)), _module.__default.safeIndex((_dafny.ZERO).minus(_637_v0), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(143)), ((_677_v3) => function (_678_i2) {
            return (_677_v3).f20;
          })(_640_v3))).length)), new _dafny.CodePoint('h'.codePointAt(0)));
          let _rhs112 = (_674_v19).IsSubsetOf((_674_v19).Union(_674_v19));
          let _rhs113 = _637_v0;
          let _lhs78 = globalState;
          let _lhs79 = _672_v17;
          let _lhs80 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_672_v17).length));
          let _lhs81 = globalState;
          _lhs78.f0 = _rhs110;
          _lhs79[_lhs80] = _rhs111;
          _lhs81.f3 = _rhs112;
          r0 = _rhs113;
          let _679_v20;
          _679_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_637_v0, new BigNumber((_642_v5).cardinality())),_644_v7);
          let _680_v21;
          _680_v21 = _dafny.Set.fromElements(new _dafny.CodePoint('e'.codePointAt(0)));
          _679_v20 = (_679_v20).update(((_dafny.ZERO).minus(_module.__default.fm0(globalState))).plus(_637_v0), (_680_v21).IsProperSubsetOf(_680_v21));
          let _index85 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_667_v13).length));
          (_667_v13)[_index85] = true;
          let _681_v22;
          _681_v22 = _dafny.MultiSet.fromElements(_644_v7, _644_v7);
          let _index86 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_667_v13).length));
          (_667_v13)[_index86] = ((!(!(_644_v7)) || (_644_v7)) ? ((_681_v22).IsProperSubsetOf(_681_v22)) : (_644_v7));
          _644_v7 = !(_dafny.areEqual((_672_v17)[_module.__default.safeIndex(new BigNumber(553), new BigNumber((_672_v17).length))], (_672_v17)[_module.__default.safeIndex(new BigNumber(553), new BigNumber((_672_v17).length))]));
          let _682_v23;
          _682_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(606),_639_v2);
          let _683_v24;
          _683_v24 = _dafny.Seq.of(_682_v23);
          (globalState).f0 = (_637_v0).isLessThanOrEqualTo(new BigNumber(((_683_v24)[_module.__default.safeIndex(new BigNumber(744), new BigNumber((_683_v24).length))]).length));
        }
        let _684_v25;
        let _nw119 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Set.Empty);
        _684_v25 = _nw119;
        let _685_v26;
        _685_v26 = _dafny.Set.fromElements(_667_v13);
        let _index87 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_684_v25).length));
        (_684_v25)[_index87] = _685_v26;
        let _index88 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_684_v25).length));
        (_684_v25)[_index88] = (_685_v26).Difference(_685_v26);
        (globalState).f0 = _644_v7;
      } else {
        (globalState).f2 = _637_v0;
        let _686_v27;
        _686_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_645_v8).length),_640_v3);
        let _687_v28;
        _687_v28 = _dafny.Map.Empty.slice().updateUnsafe(_686_v27,!(_644_v7));
        _687_v28 = (_687_v28).update(_686_v27, false);
        if (_644_v7) {
          let _688_v29;
          let _nw120 = new _module.C1();
          _nw120.__ctor(((_644_v7) ? ((_640_v3).f20) : (new _dafny.CodePoint('m'.codePointAt(0)))));
          _688_v29 = _nw120;
          (globalState).f9 = (_module.__default.safeModuloInt(_637_v0, new BigNumber((_module.__default.fm10(globalState)).length))).minus((_637_v0).plus(new BigNumber((_645_v8).length)));
          let _689_v30;
          let _init14 = function (_690_i3) {
            return true;
          };
          let _nw121 = Array((new BigNumber(5)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw121.length); _i0_14++) {
            _nw121[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _689_v30 = _nw121;
          let _691_v31;
          _691_v31 = _module.D2.create_DC5(_644_v7, new BigNumber(324), _689_v30, _642_v5);
          let _692_v32;
          _692_v32 = _dafny.MultiSet.fromElements(_691_v31);
          let _693_v33;
          _693_v33 = _dafny.Map.Empty.slice().updateUnsafe(_644_v7,_dafny.Seq.UnicodeFromString("saqvvnypo"));
          let _694_v34;
          _694_v34 = _dafny.Seq.of(_692_v32, (_692_v32).update(_691_v31, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_693_v33).length)))));
          _694_v34 = _694_v34;
          (globalState).f0 = !(!(_644_v7)) || (_644_v7);
          (globalState).f2 = _module.__default.fm0(globalState);
        } else {
          let _695_v35;
          let _nw122 = new _module.C1();
          _nw122.__ctor((_640_v3).f20);
          _695_v35 = _nw122;
          let _696_v36;
          _696_v36 = _dafny.MultiSet.fromElements(_645_v8);
          let _697_v37;
          _697_v37 = _dafny.Seq.of(_646_v9, _module.D5.create_DC13(_637_v0, new BigNumber((_696_v36).cardinality()), _644_v7));
          let _698_v38;
          _698_v38 = _dafny.Seq.of(_644_v7, _dafny.Seq.IsPrefixOf(_697_v37, _dafny.Seq.update(_697_v37, _module.__default.safeIndex(_637_v0, new BigNumber((_697_v37).length)), _646_v9)));
          (globalState).f10 = (_698_v38)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_644_v7, false)).length), _637_v0), new BigNumber((_698_v38).length))];
          (globalState).f2 = _637_v0;
          let _699_v39;
          _699_v39 = _dafny.Map.Empty.slice().updateUnsafe(_657_v11,new _dafny.CodePoint('d'.codePointAt(0)));
          let _700_v40;
          let _nw123 = Array((new BigNumber(27)).toNumber());
          _nw123[(_dafny.ZERO).toNumber()] = _639_v2;
          _nw123[(_dafny.ONE).toNumber()] = _module.__default.fm7(_639_v2, _module.__default.fm1(_644_v7, _dafny.Seq.Create(_module.__default.abs(new BigNumber(295)), ((_701_v3) => function (_702_i4) {
            return (_701_v3).f20;
          })(_640_v3)), globalState), globalState);
          _nw123[(new BigNumber(2)).toNumber()] = (_695_v35).f20;
          _nw123[(new BigNumber(3)).toNumber()] = (_640_v3).f20;
          _nw123[(new BigNumber(4)).toNumber()] = (_640_v3).f20;
          _nw123[(new BigNumber(5)).toNumber()] = (_640_v3).f20;
          _nw123[(new BigNumber(6)).toNumber()] = (_640_v3).f20;
          _nw123[(new BigNumber(7)).toNumber()] = (_695_v35).f20;
          _nw123[(new BigNumber(8)).toNumber()] = (_640_v3).f20;
          _nw123[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
          _nw123[(new BigNumber(10)).toNumber()] = _639_v2;
          _nw123[(new BigNumber(11)).toNumber()] = (_640_v3).f20;
          _nw123[(new BigNumber(12)).toNumber()] = _639_v2;
          _nw123[(new BigNumber(13)).toNumber()] = _639_v2;
          _nw123[(new BigNumber(14)).toNumber()] = (_640_v3).f20;
          _nw123[(new BigNumber(15)).toNumber()] = _639_v2;
          _nw123[(new BigNumber(16)).toNumber()] = (_640_v3).f20;
          _nw123[(new BigNumber(17)).toNumber()] = (((_699_v39).contains(_657_v11)) ? ((_699_v39).get(_657_v11)) : (new _dafny.CodePoint('f'.codePointAt(0))));
          _nw123[(new BigNumber(18)).toNumber()] = (_640_v3).f20;
          _nw123[(new BigNumber(19)).toNumber()] = (_695_v35).f20;
          _nw123[(new BigNumber(20)).toNumber()] = (_695_v35).f20;
          _nw123[(new BigNumber(21)).toNumber()] = _639_v2;
          _nw123[(new BigNumber(22)).toNumber()] = _639_v2;
          _nw123[(new BigNumber(23)).toNumber()] = (_640_v3).f20;
          _nw123[(new BigNumber(24)).toNumber()] = new _dafny.CodePoint('j'.codePointAt(0));
          _nw123[(new BigNumber(25)).toNumber()] = (_640_v3).f20;
          _nw123[(new BigNumber(26)).toNumber()] = _639_v2;
          _700_v40 = _nw123;
          let _index89 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_700_v40).length));
          (_700_v40)[_index89] = (_695_v35).f20;
          let _703_v41;
          let _nw124 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
          _703_v41 = _nw124;
          let _704_v42;
          _704_v42 = _dafny.Map.Empty.slice().updateUnsafe(_703_v41,(_640_v3).f20);
          let _705_v43;
          _705_v43 = _dafny.Seq.of(_703_v41);
          let _index90 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_700_v40).length));
          (_700_v40)[_index90] = (((_704_v42).contains((_705_v43)[_module.__default.safeIndex((_dafny.ZERO).minus(_637_v0), new BigNumber((_705_v43).length))])) ? ((_704_v42).get((_705_v43)[_module.__default.safeIndex((_dafny.ZERO).minus(_637_v0), new BigNumber((_705_v43).length))])) : ((_640_v3).f20));
          let _706_v44;
          let _init15 = ((_707_v7) => function (_708_i5) {
            return _707_v7;
          })(_644_v7);
          let _nw125 = Array((new BigNumber(10)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw125.length); _i0_15++) {
            _nw125[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _706_v44 = _nw125;
          let _709_v45;
          _709_v45 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_637_v0));
          let _710_v46;
          _710_v46 = _dafny.Seq.of(_709_v45);
          let _index91 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_706_v44).length));
          (_706_v44)[_index91] = ((_710_v46)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_640_v3).f20)).length), new BigNumber((_710_v46).length))]).IsProperSubsetOf(_709_v45);
          let _711_v47;
          _711_v47 = _module.D6.create_DC17((_700_v40)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_700_v40).length))], _644_v7, !(_644_v7), (_640_v3).f20);
          let _index92 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_706_v44).length));
          (_706_v44)[_index92] = (_711_v47).dtor_cf34;
        }
        (globalState).f2 = new BigNumber(89);
        let _712_v48;
        let _nw126 = new _module.C1();
        _nw126.__ctor(new _dafny.CodePoint('d'.codePointAt(0)));
        _712_v48 = _nw126;
      }
      let _713_v49;
      _713_v49 = _dafny.Set.fromElements(_644_v7);
      let _714_v50;
      _714_v50 = _module.D6.create_DC16(_637_v0, _644_v7, new BigNumber((_645_v8).length), new BigNumber((_713_v49).length), _644_v7);
      if (!(_644_v7) || (((_644_v7) ? (_644_v7) : ((_714_v50).dtor_cf32)))) {
        let _715_v51;
        _715_v51 = _dafny.Seq.of((new BigNumber((_dafny.Seq.of(_644_v7, _644_v7)).length)).minus(_637_v0), _637_v0);
        _715_v51 = _715_v51;
        _644_v7 = _644_v7;
        let _716_v52;
        let _nw127 = new _module.C1();
        _nw127.__ctor(new _dafny.CodePoint('k'.codePointAt(0)));
        _716_v52 = _nw127;
        (globalState).f3 = _644_v7;
        _639_v2 = (_716_v52).f20;
      } else {
        let _717_v53;
        let _nw128 = new _module.C0();
        _nw128.__ctor(_638_v1, _637_v0);
        _717_v53 = _nw128;
        let _718_v54;
        _718_v54 = _dafny.Map.Empty.slice().updateUnsafe(_637_v0,_644_v7);
        let _719_v55;
        _719_v55 = _dafny.Seq.of(_644_v7, _644_v7);
        let _720_v56;
        let _nw129 = Array((new BigNumber(25)).toNumber());
        _nw129[(_dafny.ZERO).toNumber()] = (_718_v54).equals(_718_v54);
        _nw129[(_dafny.ONE).toNumber()] = _644_v7;
        _nw129[(new BigNumber(2)).toNumber()] = (_637_v0).isLessThanOrEqualTo(_637_v0);
        _nw129[(new BigNumber(3)).toNumber()] = !(_644_v7);
        _nw129[(new BigNumber(4)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(5)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(6)).toNumber()] = (_719_v55)[_module.__default.safeIndex(_637_v0, new BigNumber((_719_v55).length))];
        _nw129[(new BigNumber(7)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(8)).toNumber()] = ((!(false)) ? (_644_v7) : (_644_v7));
        _nw129[(new BigNumber(9)).toNumber()] = _module.__default.fm1(_644_v7, _module.__default.fm6(!(_644_v7), globalState), globalState);
        _nw129[(new BigNumber(10)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(11)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(12)).toNumber()] = (_637_v0).isLessThanOrEqualTo(new BigNumber((_713_v49).length));
        _nw129[(new BigNumber(13)).toNumber()] = !(_642_v5).equals(_642_v5);
        _nw129[(new BigNumber(14)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(15)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(16)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(17)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(18)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(19)).toNumber()] = true;
        _nw129[(new BigNumber(20)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(21)).toNumber()] = (_714_v50).dtor_cf29;
        _nw129[(new BigNumber(22)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(23)).toNumber()] = _644_v7;
        _nw129[(new BigNumber(24)).toNumber()] = _644_v7;
        _720_v56 = _nw129;
        let _rhs114 = _dafny.Seq.Concat(_645_v8, _645_v8);
        let _rhs115 = _717_v53;
        let _rhs116 = _645_v8;
        let _rhs117 = _720_v56;
        let _lhs82 = globalState;
        _645_v8 = _rhs114;
        _717_v53 = _rhs115;
        _645_v8 = _rhs116;
        _lhs82.f11 = _rhs117;
        (globalState).f9 = _637_v0;
        let _721_v57;
        _721_v57 = _dafny.Map.Empty.slice().updateUnsafe(_644_v7,_644_v7);
        let _pat_let_tv21 = _719_v55;
        let _722_v58;
        _722_v58 = _dafny.Seq.of(function (_pat_let27_0) {
          return function (_723_dt__update__tmp_h0) {
            return function (_pat_let28_0) {
              return function (_724_dt__update_hcf19_h0) {
                return _module.D5.create_DC13(_724_dt__update_hcf19_h0, (_723_dt__update__tmp_h0).dtor_cf20, (_723_dt__update__tmp_h0).dtor_cf21);
              }(_pat_let28_0);
            }(new BigNumber((_pat_let_tv21).length));
          }(_pat_let27_0);
        }(_646_v9), _646_v9);
        _721_v57 = (_721_v57).update(_module.__default.fm1(_644_v7, _dafny.Seq.Create(_module.__default.abs(new BigNumber(269)), ((_725_v3) => function (_726_i6) {
          return (_725_v3).f20;
        })(_640_v3)), globalState), _dafny.Seq.IsProperPrefixOf(_722_v58, _dafny.Seq.Create(_module.__default.abs(new BigNumber(782)), ((_727_v9) => function (_728_i7) {
          return _727_v9;
        })(_646_v9))));
        let _729_v59;
        _729_v59 = _module.D1.create_DC2(_644_v7);
        let _730_v60;
        _730_v60 = _module.D1.create_DC3(_729_v59);
        _730_v60 = _730_v60;
        if (_644_v7) {
          (globalState).f10 = ((_644_v7) ? (((_644_v7) ? (_644_v7) : (_module.__default.fm1(_644_v7, _dafny.Seq.UnicodeFromString("xtbk"), globalState)))) : (_644_v7));
          let _731_v61;
          let _nw130 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _731_v61 = _nw130;
          _731_v61 = _731_v61;
          (globalState).f3 = _644_v7;
          let _732_v62;
          _732_v62 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_721_v57).length)),_637_v0);
          _732_v62 = _732_v62;
          (globalState).f2 = (_637_v0).minus(_637_v0);
        } else {
          (globalState).f10 = _644_v7;
          let _733_v63;
          let _nw131 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _733_v63 = _nw131;
          let _734_v64;
          _734_v64 = _module.D4.create_DC10(_733_v63);
          let _735_v65;
          _735_v65 = _module.D4.create_DC11(_734_v64);
          let _736_v66;
          _736_v66 = _dafny.Seq.of(_637_v0, _637_v0);
          let _737_v67;
          _737_v67 = _dafny.Map.Empty.slice().updateUnsafe(_735_v65,_736_v66);
          let _738_v68;
          _738_v68 = _dafny.Map.Empty.slice().updateUnsafe((((_737_v67).contains(_735_v65)) ? ((_737_v67).get(_735_v65)) : (_736_v66)),_module.__default.safeModuloInt(_637_v0, _637_v0));
          _738_v68 = (_738_v68).update(((_644_v7) ? (_736_v66) : (_736_v66)), new BigNumber(329));
          _730_v60 = (((_644_v7) || (_644_v7)) ? (_730_v60) : (_730_v60));
          _637_v0 = _637_v0;
          let _739_v69;
          _739_v69 = _dafny.Map.Empty.slice().updateUnsafe(_720_v56,_644_v7);
          let _740_v71;
          _740_v71 = _dafny.Set.fromElements(_637_v0, _637_v0);
          (globalState).f0 = (((_739_v69).contains(_720_v56)) ? ((_739_v69).get(_720_v56)) : ((_740_v71).IsProperSubsetOf(function () {
            let _coll18 = new _dafny.Set();
            for (const _compr_18 of _dafny.IntegerRange(new BigNumber(-694), new BigNumber(-854))) {
              let _741_v70 = _compr_18;
              if (((new BigNumber(-694)).isLessThanOrEqualTo(_741_v70)) && ((_741_v70).isLessThan(new BigNumber(-854)))) {
                _coll18.add((_741_v70).minus(_637_v0));
              }
            }
            return _coll18;
          }())));
        }
      }
      r0 = (_dafny.ZERO).minus((_637_v0).minus((_637_v0).multipliedBy(_637_v0)));
      return r0;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f17 = _dafny.ZERO;
      this._f18 = _dafny.ZERO;
      this._f19 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f19, f17, f18) {
      let _this = this;
      (_this)._f19 = f19;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      return;
    }
    fm2(globalState) {
      let _this = this;
      return (_this).f17;
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      return (_this).f18;
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt((_this).f18, ((_this).f18).plus((_this).f18));
    };
    m1(globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let _742_v0;
      _742_v0 = _dafny.Seq.of((_this).f18);
      let _hi3 = ((_this).f18).multipliedBy(new BigNumber(((_this).f19).length));
      for (let _743_i0 = (_dafny.ZERO).minus((_742_v0)[_module.__default.safeIndex(new BigNumber(-233), new BigNumber((_742_v0).length))]); _743_i0.isLessThan(_hi3); _743_i0 = _743_i0.plus(_dafny.ONE)) {
        let _source11 = (_this).f17;
        let _744___mcc_h0 = _source11;
        let _745_cf0 = _744___mcc_h0;
        (globalState).f2 = _743_i0;
        let _746_v1;
        _746_v1 = true;
        let _747_v2;
        _747_v2 = _dafny.Map.Empty.slice().updateUnsafe(!(((_746_v1) ? (_746_v1) : (_746_v1))),_745_cf0);
        (globalState).f2 = (_dafny.ZERO).minus((((_747_v2).contains((_743_i0).isLessThanOrEqualTo(new BigNumber(((_this).f19).length)))) ? ((_747_v2).get((_743_i0).isLessThanOrEqualTo(new BigNumber(((_this).f19).length)))) : (new BigNumber(720))));
        (globalState).f9 = new BigNumber(((_this).f19).length);
        let _748_v3;
        let _nw132 = Array((new BigNumber(18)).toNumber()).fill([]);
        _748_v3 = _nw132;
        let _749_v4;
        _749_v4 = _dafny.Map.Empty.slice().updateUnsafe(_746_v1,_746_v1);
        let _750_v5;
        let _nw133 = Array((new BigNumber(6)).toNumber());
        _nw133[(_dafny.ZERO).toNumber()] = _746_v1;
        _nw133[(_dafny.ONE).toNumber()] = (((_749_v4).contains(!(_746_v1))) ? ((_749_v4).get(!(_746_v1))) : (_746_v1));
        _nw133[(new BigNumber(2)).toNumber()] = _746_v1;
        _nw133[(new BigNumber(3)).toNumber()] = _module.__default.fm1(true, _dafny.Seq.UnicodeFromString("ti"), globalState);
        _nw133[(new BigNumber(4)).toNumber()] = _746_v1;
        _nw133[(new BigNumber(5)).toNumber()] = true;
        _750_v5 = _nw133;
        let _index93 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_748_v3).length));
        (_748_v3)[_index93] = _750_v5;
        let _index94 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_750_v5).length));
        (_750_v5)[_index94] = true;
        let _751_v6;
        _751_v6 = _module.D1.create_DC1(_750_v5);
        let _752_v7;
        _752_v7 = _dafny.Seq.of(_746_v1);
        let _index95 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_748_v3).length));
        let _index96 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_750_v5).length));
        let _rhs118 = (_751_v6).dtor_cf1;
        let _rhs119 = _743_i0;
        let _rhs120 = (((_752_v7)[_module.__default.safeIndex(_743_i0, new BigNumber((_752_v7).length))]) ? (_746_v1) : (_746_v1));
        let _rhs121 = _746_v1;
        let _lhs83 = _748_v3;
        let _lhs84 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_748_v3).length));
        let _lhs85 = globalState;
        let _lhs86 = globalState;
        let _lhs87 = _750_v5;
        let _lhs88 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_750_v5).length));
        _lhs83[_lhs84] = _rhs118;
        _lhs85.f2 = _rhs119;
        _lhs86.f0 = _rhs120;
        _lhs87[_lhs88] = _rhs121;
        let _753_v8;
        let _nw134 = Array((new BigNumber(19)).toNumber()).fill(false);
        _753_v8 = _nw134;
        let _754_v9;
        _754_v9 = _module.D1.create_DC1(_753_v8);
        let _source12 = _754_v9;
        if (_source12.is_DC2) {
          let _755___mcc_h1 = (_source12).cf2;
          let _756_cf2 = _755___mcc_h1;
          let _757_v10;
          _757_v10 = new _dafny.CodePoint('x'.codePointAt(0));
          let _758_v11;
          let _nw135 = new _module.C1();
          _nw135.__ctor(_757_v10);
          _758_v11 = _nw135;
          (globalState).f9 = (_this).f18;
          let _759_v12;
          let _nw136 = new _module.C0();
          _nw136.__ctor((_this).f17, (_dafny.ZERO).minus((_this).f18));
          _759_v12 = _nw136;
          let _760_v13;
          _760_v13 = _dafny.Map.Empty.slice().updateUnsafe(_743_i0,new BigNumber(549));
          let _761_v14;
          _761_v14 = _dafny.Map.Empty.slice().updateUnsafe(_756_cf2,(_this).f18);
          let _index97 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_753_v8).length));
          (_753_v8)[_index97] = !(_dafny.Set.fromElements(_743_i0, (_this).f18)).equals(_module.__default.fm8((((_761_v14).contains(_756_cf2)) ? ((_761_v14).get(_756_cf2)) : (_743_i0)), _743_i0, globalState));
          let _index98 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_753_v8).length));
          let _rhs122 = _757_v10;
          let _rhs123 = _743_i0;
          let _rhs124 = (_760_v13).Merge(_760_v13);
          let _rhs125 = (_743_i0).isLessThan((_this).f18);
          let _lhs89 = globalState;
          let _lhs90 = _753_v8;
          let _lhs91 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_753_v8).length));
          _757_v10 = _rhs122;
          _lhs89.f9 = _rhs123;
          _760_v13 = _rhs124;
          _lhs90[_lhs91] = _rhs125;
        } else if (_source12.is_DC1) {
          let _762___mcc_h2 = (_source12).cf1;
          let _763_cf1 = _762___mcc_h2;
          let _764_v15;
          _764_v15 = false;
          let _index99 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_753_v8).length));
          (_753_v8)[_index99] = _764_v15;
          let _765_v16;
          let _nw137 = new _module.C0();
          _nw137.__ctor((_this).f18, _743_i0);
          _765_v16 = _nw137;
          let _766_v17;
          _766_v17 = _dafny.Set.fromElements(_765_v16, _765_v16);
          let _index100 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_753_v8).length));
          (_753_v8)[_index100] = (_766_v17).IsProperSubsetOf((_766_v17).Difference(_766_v17));
          let _767_v18;
          let _nw138 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _767_v18 = _nw138;
          let _index101 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_767_v18).length));
          (_767_v18)[_index101] = new BigNumber(874);
          let _index102 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_753_v8).length));
          let _index103 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_767_v18).length));
          let _rhs126 = (_753_v8)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_753_v8).length))];
          let _rhs127 = (_this).f18;
          let _lhs92 = _753_v8;
          let _lhs93 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_753_v8).length));
          let _lhs94 = _767_v18;
          let _lhs95 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_767_v18).length));
          _lhs92[_lhs93] = _rhs126;
          _lhs94[_lhs95] = _rhs127;
          _765_v16 = _765_v16;
          let _768_v19;
          _768_v19 = _dafny.Set.fromElements((_this).f18, _743_i0);
          let _769_v20;
          _769_v20 = _dafny.Map.Empty.slice().updateUnsafe(_768_v19,_763_cf1);
          _769_v20 = (_769_v20).update(_768_v19, _763_cf1);
        } else {
          let _770___mcc_h3 = (_source12).cf3;
          let _771_cf3 = _770___mcc_h3;
          let _772_v21;
          _772_v21 = true;
          let _773_v22;
          _773_v22 = new _dafny.CodePoint('t'.codePointAt(0));
          let _774_v23;
          let _nw139 = new _module.C1();
          _nw139.__ctor(_773_v22);
          _774_v23 = _nw139;
          let _775_v24;
          _775_v24 = _dafny.Seq.of(((_772_v21) ? (_774_v23) : (_774_v23)), _774_v23, _774_v23, ((_772_v21) ? (_774_v23) : (_774_v23)), _774_v23);
          _775_v24 = _dafny.Seq.Concat(_775_v24, _775_v24);
          let _776_v25;
          let _nw140 = new _module.C0();
          _nw140.__ctor(new BigNumber(((_this).f19).length), (_this).f18);
          _776_v25 = _nw140;
          let _777_v26;
          _777_v26 = _module.D4.create_DC9(_776_v25);
          let _778_v27;
          _778_v27 = _module.D4.create_DC11(_777_v26);
          let _779_v28;
          _779_v28 = _module.D4.create_DC11(_778_v27);
          let _780_v29;
          _780_v29 = _dafny.Map.Empty.slice().updateUnsafe(_772_v21,_779_v28);
          let _781_v30;
          _781_v30 = _dafny.Set.fromElements(_780_v29);
          let _782_v31;
          _782_v31 = _module.D7.create_DC18(_781_v30);
          _781_v30 = (_782_v31).dtor_cf37;
          let _783_v32;
          let _784_v33;
          let _out26;
          let _out27;
          let _outcollector11 = (_774_v23).m3(((_this).f18).isEqualTo((_this).f18), globalState);
          _out26 = _outcollector11[0];
          _out27 = _outcollector11[1];
          _783_v32 = _out26;
          _784_v33 = _out27;
          let _785_v34;
          let _nw141 = new _module.C2();
          _nw141.__ctor();
          _785_v34 = _nw141;
          let _786_v35;
          let _nw142 = Array((new BigNumber(13)).toNumber());
          _nw142[(_dafny.ZERO).toNumber()] = _785_v34;
          _nw142[(_dafny.ONE).toNumber()] = _785_v34;
          _nw142[(new BigNumber(2)).toNumber()] = _785_v34;
          _nw142[(new BigNumber(3)).toNumber()] = _785_v34;
          _nw142[(new BigNumber(4)).toNumber()] = (_module.D8.create_DC22(_785_v34)).dtor_cf45;
          _nw142[(new BigNumber(5)).toNumber()] = _785_v34;
          _nw142[(new BigNumber(6)).toNumber()] = ((_772_v21) ? (_785_v34) : (_785_v34));
          _nw142[(new BigNumber(7)).toNumber()] = _785_v34;
          _nw142[(new BigNumber(8)).toNumber()] = _785_v34;
          _nw142[(new BigNumber(9)).toNumber()] = _785_v34;
          _nw142[(new BigNumber(10)).toNumber()] = _785_v34;
          _nw142[(new BigNumber(11)).toNumber()] = _785_v34;
          _nw142[(new BigNumber(12)).toNumber()] = _785_v34;
          _786_v35 = _nw142;
          let _index104 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_786_v35).length));
          (_786_v35)[_index104] = _785_v34;
          let _index105 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_786_v35).length));
          (_786_v35)[_index105] = _785_v34;
        }
        let _787_v36;
        let _nw143 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Set.Empty);
        _787_v36 = _nw143;
        let _788_v37;
        _788_v37 = false;
        let _789_v38;
        _789_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f19).length),_788_v37);
        let _790_v40;
        _790_v40 = _dafny.Set.fromElements((_dafny.ZERO).minus((_this).f18), _743_i0);
        let _791_v41;
        _791_v41 = _dafny.Seq.of(_789_v38, _789_v38, function () {
          let _coll19 = new _dafny.Map();
          for (const _compr_19 of (_790_v40).Elements) {
            let _792_v39 = _compr_19;
            if ((_790_v40).contains(_792_v39)) {
              _coll19.push([_module.__default.safeDivisionInt(_792_v39, _743_i0),_788_v37]);
            }
          }
          return _coll19;
        }());
        let _793_v42;
        _793_v42 = _dafny.Map.Empty.slice().updateUnsafe(_788_v37,(_this).fm4(new BigNumber((_791_v41).length), (_this).f19, _742_v0, globalState));
        let _794_v43;
        let _nw144 = Array((new BigNumber(6)).toNumber());
        _794_v43 = _nw144;
        let _795_v44;
        _795_v44 = _dafny.MultiSet.fromElements(_794_v43, _794_v43, _794_v43, _794_v43, _794_v43);
        let _796_v45;
        _796_v45 = _dafny.MultiSet.fromElements(new BigNumber((_793_v42).length), (((_795_v44).contains(_794_v43)) ? ((_795_v44).get(_794_v43)) : ((_this).f18)));
        let _797_v46;
        _797_v46 = _module.D2.create_DC5(_788_v37, (_dafny.ZERO).minus((_this).f18), _753_v8, _796_v45);
        let _798_v47;
        _798_v47 = _dafny.Set.fromElements(_module.D2.create_DC5(_788_v37, _743_i0, _753_v8, _796_v45), _797_v46);
        let _index106 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_787_v36).length));
        (_787_v36)[_index106] = _798_v47;
        let _index107 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_787_v36).length));
        (_787_v36)[_index107] = _798_v47;
        let _799_v48;
        let _nw145 = new _module.C0();
        _nw145.__ctor((_this).f17, (_this).f18);
        _799_v48 = _nw145;
        let _800_v49;
        _800_v49 = _dafny.Map.Empty.slice().updateUnsafe(_799_v48,_788_v37);
        (globalState).f3 = !((_800_v49).Merge(_800_v49)).equals((_800_v49).Merge(_dafny.Map.Empty.slice().updateUnsafe(_799_v48,_788_v37)));
      }
      let _hi4 = new BigNumber(-109);
      for (let _801_i1 = (_this).f18; _801_i1.isLessThan(_hi4); _801_i1 = _801_i1.plus(_dafny.ONE)) {
        let _802_v50;
        _802_v50 = false;
        (globalState).f2 = (((_802_v50) ? ((_this).fm3(_802_v50, (_this).f18, globalState)) : ((_this).f18))).plus((_this).f18);
        let _803_v51;
        _803_v51 = new _dafny.CodePoint('x'.codePointAt(0));
        _803_v51 = _803_v51;
        (globalState).f9 = (_801_i1).plus((_dafny.ZERO).minus((_742_v0)[_module.__default.safeIndex(new BigNumber(((_this).f19).length), new BigNumber((_742_v0).length))]));
        _802_v50 = _802_v50;
      }
      let _804_v52;
      let _init16 = function (_805_i2) {
        return (_805_i2).plus(new BigNumber(-107));
      };
      let _nw146 = Array((new BigNumber(25)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw146.length); _i0_16++) {
        _nw146[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _804_v52 = _nw146;
      let _index108 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length));
      (_804_v52)[_index108] = (_this).f18;
      let _index109 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length));
      (_804_v52)[_index109] = ((_this).f18).minus((_dafny.ZERO).minus(_module.__default.fm0(globalState)));
      let _806_v53;
      let _nw147 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _806_v53 = _nw147;
      let _807_v54;
      let _init17 = function (_808_i3) {
        return !_dafny.areEqual(_dafny.Seq.UnicodeFromString("alyvhe"), (_this).f19);
      };
      let _nw148 = Array((new BigNumber(16)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw148.length); _i0_17++) {
        _nw148[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _807_v54 = _nw148;
      let _809_v55;
      _809_v55 = true;
      let _index110 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length));
      (_807_v54)[_index110] = _809_v55;
      let _810_v56;
      _810_v56 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f19).length),_809_v55)).length));
      let _index111 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length));
      let _rhs128 = _806_v53;
      let _rhs129 = (_810_v56).IsSubsetOf(_810_v56);
      let _lhs96 = _807_v54;
      let _lhs97 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length));
      _806_v53 = _rhs128;
      _lhs96[_lhs97] = _rhs129;
      let _811_v57;
      _811_v57 = _dafny.Seq.UnicodeFromString("baxh");
      let _812_v58;
      _812_v58 = _dafny.MultiSet.fromElements(_804_v52);
      let _rhs130 = (_this).f19;
      let _rhs131 = (_this).f18;
      let _rhs132 = _812_v58;
      let _rhs133 = (((_807_v54)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length))]) ? ((_807_v54)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length))]) : ((_807_v54)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length))]));
      let _lhs98 = globalState;
      let _lhs99 = globalState;
      _811_v57 = _rhs130;
      _lhs98.f9 = _rhs131;
      _812_v58 = _rhs132;
      _lhs99.f0 = _rhs133;
      let _813_i4;
      _813_i4 = _dafny.ZERO;
      L3: {
        while (_809_v55) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_813_i4)) {
              break L3;
            }
            _813_i4 = (_813_i4).plus(_dafny.ONE);
            if (_809_v55) {
              let _814_v59;
              _814_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_807_v54)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length))]);
              let _815_v60;
              let _816_v61;
              let _out28;
              let _out29;
              let _outcollector12 = _module.__default.m0(_809_v55, _814_v59, globalState);
              _out28 = _outcollector12[0];
              _out29 = _outcollector12[1];
              _815_v60 = _out28;
              _816_v61 = _out29;
              let _817_v63;
              _817_v63 = _dafny.Set.fromElements(_814_v59);
              let _818_v64;
              _818_v64 = _module.D6.create_DC15(new BigNumber((function () {
  let _coll20 = new _dafny.Map();
  for (const _compr_20 of _dafny.IntegerRange(new BigNumber(-322), new BigNumber(616))) {
    let _819_v62 = _compr_20;
    if (((new BigNumber(-322)).isLessThanOrEqualTo(_819_v62)) && ((_819_v62).isLessThan(new BigNumber(616)))) {
      _coll20.push([(_819_v62).plus((_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))]),(_dafny.ZERO).minus(_816_v61)]);
    }
  }
  return _coll20;
}()).length), _816_v61, _dafny.Seq.of((_this).f18), (_this).f18, _dafny.Set.fromElements(_815_v60, new BigNumber((_817_v63).length), (_dafny.ZERO).minus((_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))]), _815_v60, _816_v61));
              let _820_v65;
              _820_v65 = new _dafny.CodePoint('c'.codePointAt(0));
              let _821_v66;
              _821_v66 = _dafny.Map.Empty.slice().updateUnsafe((_807_v54)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length))],(_this).f18);
              let _rhs134 = _818_v64;
              let _rhs135 = !((_821_v66).Merge(_821_v66)).equals((_821_v66).Merge(_821_v66));
              let _rhs136 = _820_v65;
              let _rhs137 = (_816_v61).multipliedBy((((_807_v54)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length))]) ? (_815_v60) : (_815_v60)));
              let _rhs138 = _module.__default.fm1(_module.__default.fm1((_807_v54)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length))], _811_v57, globalState), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("g"), _module.__default.safeIndex((_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))], new BigNumber((_dafny.Seq.UnicodeFromString("g")).length)), _820_v65), globalState);
              let _lhs100 = globalState;
              let _lhs101 = globalState;
              _818_v64 = _rhs134;
              _lhs100.f10 = _rhs135;
              _820_v65 = _rhs136;
              _815_v60 = _rhs137;
              _lhs101.f3 = _rhs138;
              let _822_v67;
              let _nw149 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
              _822_v67 = _nw149;
              let _823_v68;
              let _nw150 = new _module.C1();
              _nw150.__ctor(_820_v65);
              _823_v68 = _nw150;
              let _824_v69;
              _824_v69 = _dafny.Map.Empty.slice().updateUnsafe(_823_v68,(_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))]);
              let _index112 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_822_v67).length));
              (_822_v67)[_index112] = _824_v69;
              let _825_v70;
              _825_v70 = _module.D1.create_DC1(_807_v54);
              let _826_v71;
              _826_v71 = _dafny.Seq.of(_825_v70);
              let _index113 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length));
              let _index114 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_822_v67).length));
              let _rhs139 = (_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))];
              let _rhs140 = (_module.__default.fm0(globalState)).multipliedBy(((_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))]).multipliedBy(new BigNumber((_826_v71).length)));
              let _rhs141 = (_807_v54)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length))];
              let _rhs142 = (_dafny.Map.Empty.slice().updateUnsafe(_823_v68,new BigNumber(670))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_823_v68,new BigNumber(143))).Merge(_824_v69));
              let _rhs143 = (_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))];
              let _lhs102 = globalState;
              let _lhs103 = _807_v54;
              let _lhs104 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length));
              let _lhs105 = _822_v67;
              let _lhs106 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_822_v67).length));
              let _lhs107 = globalState;
              _lhs102.f9 = _rhs139;
              _816_v61 = _rhs140;
              _lhs103[_lhs104] = _rhs141;
              _lhs105[_lhs106] = _rhs142;
              _lhs107.f2 = _rhs143;
              (globalState).f3 = (_816_v61).isLessThanOrEqualTo((_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))]);
              (globalState).f9 = new BigNumber(((_this).f19).length);
            } else {
              (globalState).f2 = _module.__default.safeModuloInt(new BigNumber(-132), (_this).f18);
              (globalState).f2 = (_this).f18;
              _811_v57 = (_this).f19;
              _742_v0 = _742_v0;
              let _827_v72;
              _827_v72 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_module.__default.safeDivisionInt(new BigNumber(((_this).f19).length), (_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))]));
              _827_v72 = (_827_v72).update((_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))], new BigNumber(55));
            }
            (globalState).f9 = ((_this).f18).minus(((_809_v55) ? ((_this).fm3(_809_v55, (_this).f18, globalState)) : (new BigNumber(850))));
            let _828_v73;
            _828_v73 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(547)), function (_829_i6) {
              return new _dafny.CodePoint('f'.codePointAt(0));
            }), _dafny.Seq.UnicodeFromString("bj"), _811_v57, _811_v57, _811_v57);
            let _830_v74;
            _830_v74 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(564)), function (_831_i5) {
              return new _dafny.CodePoint('h'.codePointAt(0));
            }), (_this).f19, (_this).f19, (_828_v73)[_module.__default.safeIndex((_this).f18, new BigNumber((_828_v73).length))]);
            let _832_v75;
            _832_v75 = _dafny.Seq.of((_807_v54)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length))]);
            let _index115 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length));
            let _index116 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length));
            let _rhs144 = (_this).f18;
            let _rhs145 = _module.__default.safeDivisionInt(new BigNumber(((_828_v73)[_module.__default.safeIndex((_this).fm3((_832_v75)[_module.__default.safeIndex((_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))], new BigNumber((_832_v75).length))], new BigNumber((_832_v75).length), globalState), new BigNumber((_828_v73).length))]).length), (_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))]);
            let _lhs108 = _804_v52;
            let _lhs109 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length));
            let _lhs110 = _804_v52;
            let _lhs111 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length));
            _lhs108[_lhs109] = _rhs144;
            _lhs110[_lhs111] = _rhs145;
            (globalState).f9 = ((_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))]).multipliedBy((_dafny.ZERO).minus((_804_v52)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_804_v52).length))]));
          }
        }
      }
      r0 = _804_v52;
      r1 = (_807_v54)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_807_v54).length))];
      return [r0, r1];
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
