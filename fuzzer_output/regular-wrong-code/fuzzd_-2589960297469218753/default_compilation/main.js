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
    static fm0(p0, p1, p2, globalState) {
      return (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber(((_module.D9.create_DC27(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("itjb"),false)).length), new BigNumber((_dafny.Set.fromElements(true, !(false))).length)))).dtor_cf40).length))).plus(new BigNumber((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Seq.of(new BigNumber(-111))).Elements) {
          let _0_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(-111)), _0_v0)) {
            _coll0.add(_module.__default.safeModuloInt(_0_v0, new BigNumber((_dafny.Seq.UnicodeFromString("exuvfvup")).length)));
          }
        }
        return _coll0;
      }()).length)));
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return false;
    };
    static fm6(p0, globalState) {
      let _source0 = _module.D15.create_DC46();
      if (_source0.is_DC44) {
        let _1___mcc_h0 = (_source0).cf71;
        let _2___mcc_h1 = (_source0).cf72;
        let _3___mcc_h2 = (_source0).cf73;
        let _4___mcc_h3 = (_source0).cf74;
        let _5___mcc_h4 = (_source0).cf75;
        let _6_cf75 = _5___mcc_h4;
        let _7_cf74 = _4___mcc_h3;
        let _8_cf73 = _3___mcc_h2;
        let _9_cf72 = _2___mcc_h1;
        let _10_cf71 = _1___mcc_h0;
        return _9_cf72;
      } else if (_source0.is_DC45) {
        let _11___mcc_h5 = (_source0).cf76;
        let _12___mcc_h6 = (_source0).cf77;
        let _13___mcc_h7 = (_source0).cf78;
        let _14___mcc_h8 = (_source0).cf79;
        let _15___mcc_h9 = (_source0).cf80;
        let _16_cf80 = _15___mcc_h9;
        let _17_cf79 = _14___mcc_h8;
        let _18_cf78 = _13___mcc_h7;
        let _19_cf77 = _12___mcc_h6;
        let _20_cf76 = _11___mcc_h5;
        return _18_cf78;
      } else if (_source0.is_DC46) {
        return !(true);
      } else {
        let _21___mcc_h10 = (_source0).cf70;
        let _22_cf70 = _21___mcc_h10;
        return (!(true)) && (!(true));
      }
    };
    static fm7(p0, globalState) {
      return _dafny.Seq.UnicodeFromString("hgxt");
    };
    static fm8(p0, p1, globalState) {
      if (!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(759)), function (_23_i0) {
        return new BigNumber((_dafny.MultiSet.fromElements(!(false), !(true), false, false, false)).cardinality());
      }), _dafny.Seq.of(new BigNumber(6)))) {
        return _module.D2.create_DC9(_dafny.Seq.UnicodeFromString("txpfrb"), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-952), new BigNumber((_dafny.Seq.of(true, false, false, true, false)).length))).cardinality())), true, new BigNumber(-911));
      } else {
        return _module.D2.create_DC9(_dafny.Seq.UnicodeFromString("asarcrdjh"), new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(-842)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(545)), function (_24_i1) {
  return new BigNumber(-674);
}), _dafny.Seq.of(new BigNumber(46)))).cardinality()), true, new BigNumber(585));
      }
    };
    static fm9(globalState) {
      return _module.__default.safeDivisionInt(new BigNumber(444), ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("wuuxa")).length)))).multipliedBy(new BigNumber(-128)));
    };
    static fm10(p0, p1, globalState) {
      return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false)), _dafny.Seq.Concat((_module.D9.create_DC28(_dafny.Seq.of(false), false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(348),new BigNumber(715))).length), false, !(true))).dtor_cf41, _dafny.Seq.of(true)));
    };
    static fm13(p0, p1, p2, p3, globalState) {
      return _module.D4.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, false))).cardinality())));
    };
    static fm14(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("wqkswi")).length));
    };
    static fm15(p0, globalState) {
      return function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(12)), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(84))).length)))).Difference(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(242),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("awlgsvyam"))).length))).length))))).Elements) {
          let _25_v0 = _compr_1;
          if (((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(12)), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(84))).length)))).Difference(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(242),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("awlgsvyam"))).length))).length))))).contains(_25_v0)) {
            _coll1.add(_25_v0);
          }
        }
        return _coll1;
      }();
    };
    static fm17(p0, globalState) {
      return (_dafny.Set.fromElements(new BigNumber(681))).Difference(_dafny.Set.fromElements(new BigNumber(340)));
    };
    static fm18(p0, p1, p2, p3, globalState) {
      return _module.D2.create_DC9(_dafny.Seq.UnicodeFromString("nc"), new BigNumber((function () {
  let _coll2 = new _dafny.Map();
  for (const _compr_2 of _dafny.IntegerRange(new BigNumber(957), new BigNumber(-55))) {
    let _26_v0 = _compr_2;
    if (((new BigNumber(957)).isLessThanOrEqualTo(_26_v0)) && ((_26_v0).isLessThan(new BigNumber(-55)))) {
      _coll2.push([(_26_v0).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)),new BigNumber(-164)]);
    }
  }
  return _coll2;
}()).length), true, (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(!(!((_module.D9.create_DC29(new BigNumber(((_module.D9.create_DC27(_dafny.Seq.Create(_module.__default.abs(new BigNumber(849)), function (_27_i0) {
  return new BigNumber(388);
}))).dtor_cf40).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qqe")).length),new BigNumber(37))).length), true)).dtor_cf48)))).cardinality())));
    };
    static fm19(p0, globalState) {
      return ((new BigNumber(-519)).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("aabgm")).length))).isLessThan((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(true, true, false)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("enqyyh")).length)))));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fnjmfaxg"), _dafny.Seq.UnicodeFromString("ms")), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(519)), function (_28_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(734)), function (_29_i1) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })));
    };
    static fm21(p0, p1, p2, globalState) {
      if ((true) && (false)) {
        return _module.D5.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(true,_module.D1.create_DC3(_dafny.Seq.of(true))));
      } else {
        return _module.D5.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D1.create_DC3(_dafny.Seq.of(false))));
      }
    };
    static fm22(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(true);
    };
    static fm23(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(!(!(true)))), _dafny.Seq.of(_dafny.Seq.of(true))), _dafny.Seq.of(_dafny.Seq.of(true), _dafny.Seq.of(true, false, false), _dafny.Seq.of(true, false)));
    };
    static fm24(globalState) {
      return _module.D1.create_DC4(false);
    };
    static fm25(p0, p1, p2, globalState) {
      let _source1 = _module.D20.create_DC62();
      if (_source1.is_DC62) {
        return function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("vkghxcks")).length)),false)).Keys.Elements) {
            let _30_v0 = _compr_3;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("vkghxcks")).length)),false)).contains(_30_v0)) {
              _coll3.add(_30_v0);
            }
          }
          return _coll3;
        }();
      } else if (_source1.is_DC63) {
        let _31___mcc_h0 = (_source1).cf112;
        let _32_cf112 = _31___mcc_h0;
        return _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-605)), function (_33_i0) {
          return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(710)), function (_34_i1) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          })).length);
        }), _dafny.Seq.of(new BigNumber(605)));
      } else if (_source1.is_DC61) {
        let _35___mcc_h1 = (_source1).cf111;
        let _36_cf111 = _35___mcc_h1;
        return (_module.D21.create_DC65(function () {
  let _coll4 = new _dafny.Set();
  for (const _compr_4 of (_dafny.Seq.of(_dafny.Seq.of(new BigNumber(747), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-938),!(true))).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-416)), function (_37_i2) {
    return new BigNumber(-799);
  }))).Elements) {
    let _38_v1 = _compr_4;
    if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(747), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-938),!(true))).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-416)), function (_37_i2) {
      return new BigNumber(-799);
    })), _38_v1)) {
      _coll4.add(_38_v1);
    }
  }
  return _coll4;
}())).dtor_cf114;
      } else {
        let _39___mcc_h2 = (_source1).cf113;
        let _40_cf113 = _39___mcc_h2;
        return (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true, true, true)).length)))).Union(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(788)), function (_41_i3) {
          return new BigNumber(530);
        })).length)), _dafny.Seq.of(new BigNumber((function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(156)), function (_42_i4) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          })).length)))).Elements) {
            let _43_v2 = _compr_5;
            if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(156)), function (_42_i4) {
              return new _dafny.CodePoint('j'.codePointAt(0));
            })).length))), _43_v2)) {
              _coll5.push([(_43_v2).minus(new BigNumber(367)),new BigNumber(-245)]);
            }
          }
          return _coll5;
        }()).length)))).length)))));
      }
    };
    static fm26(p0, p1, p2, globalState) {
      if ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).length)).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("vqoca")).length))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("xvqqeji")).length))).cardinality()))).length))).Merge(function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of (_dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)))).Elements) {
            let _44_v0 = _compr_6;
            if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0))), _44_v0)) {
              _coll6.push([_44_v0,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),new BigNumber((function () {
                let _coll7 = new _dafny.Set();
                for (const _compr_7 of (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)))).Elements) {
                  let _45_v1 = _compr_7;
                  if ((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)))).contains(_45_v1)) {
                    _coll7.add(_45_v1);
                  }
                }
                return _coll7;
              }()).length))).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("vpk")).length), new BigNumber((function () {
                let _coll8 = new _dafny.Map();
                for (const _compr_8 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(!(false), !(false))).cardinality()))).Elements) {
                  let _46_v2 = _compr_8;
                  if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(!(false), !(false))).cardinality()))).contains(_46_v2)) {
                    _coll8.push([(_46_v2).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)),false]);
                  }
                }
                return _coll8;
              }()).length))).length))).length))]);
            }
          }
          return _coll6;
        }());
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),new BigNumber(334))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),new BigNumber(642)));
      }
    };
    static fm27(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(new BigNumber(453), new BigNumber(892)), (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(472)), function (_47_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("nbv")).length)))), (new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, !(false), false))).cardinality())).plus(new BigNumber((_dafny.Seq.UnicodeFromString("uibebmq")).length)), new BigNumber(((_module.D2.create_DC9(_dafny.Seq.UnicodeFromString("vmqtwapqp"), new BigNumber((_dafny.Seq.UnicodeFromString("sdv")).length), false, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true, true, true, !(!(!(false))))).length))).cardinality())))).cardinality()))).dtor_cf14).length));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true));
    };
    static fm29(p0, p1, p2, globalState) {
      return _module.D3.create_DC12(_module.D3.create_DC11(true, false));
    };
    static fm31(p0, p1, globalState) {
      return _dafny.Seq.of(_module.D6.create_DC21(_dafny.Set.fromElements(!(false))), _module.D6.create_DC21(_dafny.Set.fromElements(!(true), true)), _module.D6.create_DC21(_dafny.Set.fromElements(false)));
    };
    static fm32(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(false,!(!((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false)).length)))).cardinality()), (_dafny.ZERO).minus(new BigNumber(-252)))).IsProperSubsetOf((_module.D17.create_DC52(false, false, new BigNumber(-932), _dafny.MultiSet.fromElements(new BigNumber(162)))).dtor_cf91))));
    };
    static fm33(p0, p1, p2, globalState) {
      return (false) && (!(false));
    };
    static fm34(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("exnq"), _dafny.Seq.UnicodeFromString("gpfxxmfqm")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(432)), function (_48_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }));
    };
    static fm35(p0, p1, globalState) {
      return (_module.D1.create_DC3(_dafny.Seq.of(true, true, false))).dtor_cf7;
    };
    static fm36(globalState) {
      return ((_dafny.ZERO).minus(new BigNumber(-821))).minus((new BigNumber(986)).minus(new BigNumber(174)));
    };
    static fm37(globalState) {
      let _source2 = _module.D21.create_DC65(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(146)), function (_49_i0) {
  return new BigNumber(452);
})));
      if (_source2.is_DC66) {
        let _50___mcc_h0 = (_source2).cf115;
        let _51_cf115 = _50___mcc_h0;
        return _module.D6.create_DC22(_51_cf115, new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('q'.codePointAt(0)))).cardinality()), true);
      } else if (_source2.is_DC67) {
        let _52___mcc_h1 = (_source2).cf116;
        let _53_cf116 = _52___mcc_h1;
        return _module.D6.create_DC22(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_53_cf116,_53_cf116)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D16.create_DC48(),_53_cf116)).length), false);
      } else {
        let _54___mcc_h2 = (_source2).cf114;
        let _55_cf114 = _54___mcc_h2;
        return _module.D6.create_DC22(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true))).cardinality()), new BigNumber(548), false);
      }
    };
    static fm38(globalState) {
      if (!_dafny.Seq.contains(_dafny.Seq.of(false, false, true), false)) {
        return _module.D9.create_DC27(_dafny.Seq.of(new BigNumber(-945)));
      } else {
        return _module.D9.create_DC27(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)))).length)));
      }
    };
    static fm39(p0, p1, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.of(!(!(false)) || (false), !(!(true) || (true))));
    };
    static fm40(p0, globalState) {
      if ((_dafny.Set.fromElements(new BigNumber(30))).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("qfejidul")).length)))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(184)), function (_56_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),true)).length)))) {
        return _module.D0.create_DC2(_dafny.MultiSet.fromElements(false, !(!(true)), true, !(true)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(19)), function (_57_i1) {
  return new _dafny.CodePoint('j'.codePointAt(0));
}), true);
      } else {
        return _module.D0.create_DC2(_dafny.MultiSet.fromElements(true, true), _dafny.Seq.Create(_module.__default.abs(new BigNumber(481)), function (_58_i2) {
  return new _dafny.CodePoint('u'.codePointAt(0));
}), false);
      }
    };
    static fm41(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(true)),_module.D4.create_DC15(new BigNumber((_dafny.Seq.UnicodeFromString("rddv")).length), new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-719)), new BigNumber(-79))).length)))).length),true)).Merge(function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(-602), new BigNumber(331))) {
          let _59_v0 = _compr_9;
          if (((new BigNumber(-602)).isLessThanOrEqualTo(_59_v0)) && ((_59_v0).isLessThan(new BigNumber(331)))) {
            _coll9.push([_module.__default.safeDivisionInt(_59_v0, new BigNumber((_dafny.Seq.UnicodeFromString("eqcdui")).length)),true]);
          }
        }
        return _coll9;
      }());
    };
    static fm42(p0, p1, globalState) {
      let _source3 = _module.D5.create_DC18();
      if (_source3.is_DC17) {
        let _60___mcc_h0 = (_source3).cf27;
        let _61_cf27 = _60___mcc_h0;
        return (_module.D10.create_DC32(new _dafny.CodePoint('a'.codePointAt(0)), !(!(true)))).dtor_cf52;
      } else if (_source3.is_DC18) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      } else if (_source3.is_DC19) {
        let _62___mcc_h1 = (_source3).cf28;
        let _63___mcc_h2 = (_source3).cf29;
        let _64___mcc_h3 = (_source3).cf30;
        let _65_cf30 = _64___mcc_h3;
        let _66_cf29 = _63___mcc_h2;
        let _67_cf28 = _62___mcc_h1;
        return _65_cf30;
      } else if (_source3.is_DC16) {
        let _68___mcc_h4 = (_source3).cf26;
        let _69_cf26 = _68___mcc_h4;
        return new _dafny.CodePoint('y'.codePointAt(0));
      } else {
        let _70___mcc_h5 = (_source3).cf31;
        let _71_cf31 = _70___mcc_h5;
        return new _dafny.CodePoint('o'.codePointAt(0));
      }
    };
    static fm43(p0, p1, globalState) {
      if (!((_dafny.MultiSet.fromElements(new BigNumber(564), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),false)).length)))),new BigNumber(-373))).length))).length))).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false, true)).length), new BigNumber(776))))) {
        return _module.D1.create_DC3(_dafny.Seq.of(true, true));
      } else {
        return _module.D1.create_DC3(_dafny.Seq.of(true, false));
      }
    };
    static fm44(p0, p1, p2, p3, globalState) {
      if ((_dafny.MultiSet.fromElements(!(false), true)).IsProperSubsetOf(_dafny.MultiSet.fromElements(true))) {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(723)), function (_72_i0) {
          return (_module.D19.create_DC59(false, new BigNumber(810))).dtor_cf109;
        }), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)));
      } else {
        return _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(!(false))).length), new BigNumber(181));
      }
    };
    static fm45(globalState) {
      return ((function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (function () {
          let _coll11 = new _dafny.Map();
          for (const _compr_11 of (_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-250)), function (_73_i0) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          })).length)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(992), (_module.D19.create_DC59(false, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()))).dtor_cf109, new BigNumber(175), new BigNumber((_dafny.Seq.of(false)).length))).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-566)), function (_74_i1) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("druvqxil")).length);
          }))).Elements) {
            let _75_v1 = _compr_11;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-250)), function (_73_i0) {
              return new _dafny.CodePoint('a'.codePointAt(0));
            })).length)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(992), (_module.D19.create_DC59(false, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()))).dtor_cf109, new BigNumber(175), new BigNumber((_dafny.Seq.of(false)).length))).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-566)), function (_74_i1) {
              return new BigNumber((_dafny.Seq.UnicodeFromString("druvqxil")).length);
            })), _75_v1)) {
              _coll11.push([_75_v1,new _dafny.CodePoint('s'.codePointAt(0))]);
            }
          }
          return _coll11;
        }()).Keys.Elements) {
          let _76_v0 = _compr_10;
          if ((function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of (_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-250)), function (_73_i0) {
              return new _dafny.CodePoint('a'.codePointAt(0));
            })).length)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(992), (_module.D19.create_DC59(false, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()))).dtor_cf109, new BigNumber(175), new BigNumber((_dafny.Seq.of(false)).length))).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-566)), function (_74_i1) {
              return new BigNumber((_dafny.Seq.UnicodeFromString("druvqxil")).length);
            }))).Elements) {
              let _75_v1 = _compr_12;
              if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-250)), function (_73_i0) {
                return new _dafny.CodePoint('a'.codePointAt(0));
              })).length)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(992), (_module.D19.create_DC59(false, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()))).dtor_cf109, new BigNumber(175), new BigNumber((_dafny.Seq.of(false)).length))).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-566)), function (_74_i1) {
                return new BigNumber((_dafny.Seq.UnicodeFromString("druvqxil")).length);
              })), _75_v1)) {
                _coll12.push([_75_v1,new _dafny.CodePoint('s'.codePointAt(0))]);
              }
            }
            return _coll12;
          }()).contains(_76_v0)) {
            _coll10.push([_76_v0,new BigNumber(590)]);
          }
        }
        return _coll10;
      }()).Merge(function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(96)), function (_77_i2) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("jqmycnx")).length),false)).length);
        }),_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),false))).Keys.Elements) {
          let _78_v2 = _compr_13;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(96)), function (_77_i2) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("jqmycnx")).length),false)).length);
          }),_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),false))).contains(_78_v2)) {
            _coll13.push([_78_v2,new BigNumber((_dafny.Seq.UnicodeFromString("fjboup")).length)]);
          }
        }
        return _coll13;
      }())).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(513), new BigNumber(779), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("yncuqfdbc"))).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ofyfibxcp")).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(162), new BigNumber((_dafny.Seq.UnicodeFromString("te")).length), new BigNumber(286)),new BigNumber((_dafny.Seq.UnicodeFromString("dv")).length))));
    };
    static fm46(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(true), true)).length),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("s")).length),true)));
    };
    static fm47(p0, p1, globalState) {
      return function () {
        let _coll14 = new _dafny.Set();
        for (const _compr_14 of (_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, false)).length),_module.D11.create_DC35(_dafny.Seq.UnicodeFromString("edl"), true)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(118),_module.D11.create_DC35(_dafny.Seq.UnicodeFromString("tah"), true)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(591))).length),_module.D11.create_DC35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(327)), function (_79_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
}), false))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-837),_module.D11.create_DC35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(831)), function (_80_i1) {
  return new _dafny.CodePoint('s'.codePointAt(0));
}), true)), function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of _dafny.IntegerRange(new BigNumber(-456), new BigNumber(4))) {
            let _81_v0 = _compr_15;
            if (((new BigNumber(-456)).isLessThanOrEqualTo(_81_v0)) && ((_81_v0).isLessThan(new BigNumber(4)))) {
              _coll15.push([(_81_v0).plus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length),true)).length))).length)),_module.D11.create_DC35(_dafny.Seq.UnicodeFromString("rh"), true)]);
            }
          }
          return _coll15;
        }()))).Elements) {
          let _82_v1 = _compr_14;
          if (_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, false)).length),_module.D11.create_DC35(_dafny.Seq.UnicodeFromString("edl"), true)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(118),_module.D11.create_DC35(_dafny.Seq.UnicodeFromString("tah"), true)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(591))).length),_module.D11.create_DC35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(327)), function (_79_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
}), false))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-837),_module.D11.create_DC35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(831)), function (_80_i1) {
  return new _dafny.CodePoint('s'.codePointAt(0));
}), true)), function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of _dafny.IntegerRange(new BigNumber(-456), new BigNumber(4))) {
              let _81_v0 = _compr_16;
              if (((new BigNumber(-456)).isLessThanOrEqualTo(_81_v0)) && ((_81_v0).isLessThan(new BigNumber(4)))) {
                _coll16.push([(_81_v0).plus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length),true)).length))).length)),_module.D11.create_DC35(_dafny.Seq.UnicodeFromString("rh"), true)]);
              }
            }
            return _coll16;
          }())), _82_v1)) {
            _coll14.add(_82_v1);
          }
        }
        return _coll14;
      }();
    };
    static fm48(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("krnlu")).length))).length),new BigNumber((_dafny.Seq.UnicodeFromString("ojsoe")).length))).Merge((function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of (_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((function () {
          let _coll18 = new _dafny.Map();
          for (const _compr_18 of (_dafny.Seq.of(new BigNumber(-922))).Elements) {
            let _83_v1 = _compr_18;
            if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(-922)), _83_v1)) {
              _coll18.push([(_83_v1).plus(new BigNumber(631)),false]);
            }
          }
          return _coll18;
        }()).length)), _dafny.Seq.of(new BigNumber(743)))).cardinality()))).length))).length), new BigNumber(482), new BigNumber((_dafny.Seq.UnicodeFromString("lhpvmy")).length), new BigNumber((_dafny.Set.fromElements(false, true)).length))).Elements) {
          let _84_v0 = _compr_17;
          if ((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of (_dafny.Seq.of(new BigNumber(-922))).Elements) {
              let _83_v1 = _compr_19;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(-922)), _83_v1)) {
                _coll19.push([(_83_v1).plus(new BigNumber(631)),false]);
              }
            }
            return _coll19;
          }()).length)), _dafny.Seq.of(new BigNumber(743)))).cardinality()))).length))).length), new BigNumber(482), new BigNumber((_dafny.Seq.UnicodeFromString("lhpvmy")).length), new BigNumber((_dafny.Set.fromElements(false, true)).length))).contains(_84_v0)) {
            _coll17.push([(_84_v0).multipliedBy(new BigNumber((_dafny.Seq.of(false)).length)),new BigNumber((_dafny.Seq.UnicodeFromString("s")).length)]);
          }
        }
        return _coll17;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(393),new BigNumber(-765))).length),new BigNumber((_dafny.MultiSet.fromElements(!(!(false)), true)).cardinality()))));
    };
    static fm49(p0, globalState) {
      return _module.D10.create_DC32(((true) ? (new _dafny.CodePoint('c'.codePointAt(0))) : (new _dafny.CodePoint('j'.codePointAt(0)))), !(!(true) || (false)));
    };
    static fm50(p0, p1, p2, p3, globalState) {
      return _module.D6.create_DC21((_dafny.Set.fromElements(true, !(true))).Intersect(_dafny.Set.fromElements(true)));
    };
    static fm51(globalState) {
      return _module.D11.create_DC35(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(251)), function (_85_i0) {
  return new _dafny.CodePoint('q'.codePointAt(0));
}), _dafny.Seq.Create(_module.__default.abs(new BigNumber(38)), function (_86_i1) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})), (_module.D1.create_DC4(false)).dtor_cf8);
    };
    static fm52(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.Create(_module.__default.abs(new BigNumber(183)), function (_87_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }));
    };
    static fm53(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_module.D10.create_DC33(!(false), new BigNumber((_dafny.Seq.of(!(true), false)).length))).dtor_cf54,_module.D3.create_DC12(_module.D3.create_DC11(!(!(false)), false)));
    };
    static fm54(p0, p1, p2, p3, globalState) {
      if (false) {
        return _module.D15.create_DC43(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(142),false));
      } else {
        return _module.D15.create_DC43(function () {
  let _coll20 = new _dafny.Map();
  for (const _compr_20 of _dafny.IntegerRange(new BigNumber(-229), new BigNumber(503))) {
    let _88_v0 = _compr_20;
    if (((new BigNumber(-229)).isLessThanOrEqualTo(_88_v0)) && ((_88_v0).isLessThan(new BigNumber(503)))) {
      _coll20.push([(_88_v0).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("gslmjy")).length))),true]);
    }
  }
  return _coll20;
}());
      }
    };
    static fm55(globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sltsbvpru"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bslnuaicl"), _dafny.Seq.UnicodeFromString("hl")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ekjt"), _dafny.Seq.UnicodeFromString("qhpp")));
    };
    static fm56(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_module.D21.create_DC65(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(229)), function (_89_i0) {
  return new BigNumber(-225);
}))))).length),true)), _dafny.Seq.of(function () {
        let _coll21 = new _dafny.Map();
        for (const _compr_21 of (_dafny.Seq.of(new BigNumber(-338))).Elements) {
          let _90_v0 = _compr_21;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(-338)), _90_v0)) {
            _coll21.push([(_90_v0).multipliedBy(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)))).length)),false]);
          }
        }
        return _coll21;
      }()));
    };
    static fm57(p0, globalState) {
      return function () {
        let _coll22 = new _dafny.Set();
        for (const _compr_22 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(891)), function (_91_i0) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        })).length),new BigNumber((_dafny.Seq.of(false, false)).length)))).Keys.Elements) {
          let _92_v0 = _compr_22;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(891)), function (_91_i0) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })).length),new BigNumber((_dafny.Seq.of(false, false)).length)))).contains(_92_v0)) {
            _coll22.add(_92_v0);
          }
        }
        return _coll22;
      }();
    };
    static m10(p0, p1, p2, p3, globalState) {
      if (((!(p2)) ? (p3) : (p2))) {
        let _93_v0;
        let _nw0 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _93_v0 = _nw0;
        let _94_v1;
        _94_v1 = _dafny.Seq.UnicodeFromString("brux");
        let _index0 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_93_v0).length));
        (_93_v0)[_index0] = new BigNumber((_94_v1).length);
        let _index1 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_93_v0).length));
        (_93_v0)[_index1] = p1;
        let _95_v2;
        _95_v2 = new _dafny.CodePoint('s'.codePointAt(0));
        let _96_v3;
        let _nw1 = Array((_dafny.ONE).toNumber());
        _nw1[(_dafny.ZERO).toNumber()] = p2;
        _96_v3 = _nw1;
        let _index2 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_96_v3).length));
        (_96_v3)[_index2] = p3;
        let _97_v4;
        _97_v4 = _dafny.Seq.of(p3);
        let _index3 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_96_v3).length));
        let _index4 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_93_v0).length));
        let _rhs0 = p0;
        let _rhs1 = p3;
        let _rhs2 = (_module.__default.safeDivisionInt(_module.__default.fm0((_93_v0)[_module.__default.safeIndex(new BigNumber(10), new BigNumber((_93_v0).length))], p3, (_97_v4)[_module.__default.safeIndex(_module.__default.fm36(globalState), new BigNumber((_97_v4).length))], globalState), (_93_v0)[_module.__default.safeIndex(new BigNumber(10), new BigNumber((_93_v0).length))])).isLessThanOrEqualTo((_93_v0)[_module.__default.safeIndex(new BigNumber(10), new BigNumber((_93_v0).length))]);
        let _rhs3 = _module.__default.safeDivisionInt((_93_v0)[_module.__default.safeIndex(new BigNumber(10), new BigNumber((_93_v0).length))], p1);
        let _lhs0 = _96_v3;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_96_v3).length));
        let _lhs2 = globalState;
        let _lhs3 = _93_v0;
        let _lhs4 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_93_v0).length));
        _95_v2 = _rhs0;
        _lhs0[_lhs1] = _rhs1;
        _lhs2.f18 = _rhs2;
        _lhs3[_lhs4] = _rhs3;
        let _nw2 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        (globalState).f17 = _nw2;
        let _98_v5;
        let _nw3 = new _module.C2();
        _nw3.__ctor((_93_v0)[_module.__default.safeIndex(new BigNumber(10), new BigNumber((_93_v0).length))], _dafny.Set.fromElements(_93_v0, _93_v0), (_96_v3)[_module.__default.safeIndex(new BigNumber(891), new BigNumber((_96_v3).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), ((_99_v2) => function (_100_i0) {
          return _99_v2;
        })(_95_v2)));
        _98_v5 = _nw3;
        let _101_v6;
        _101_v6 = _dafny.Map.Empty.slice().updateUnsafe(_98_v5,_96_v3);
        let _102_v7;
        let _nw4 = new _module.C1();
        _nw4.__ctor(p2);
        _102_v7 = _nw4;
        let _103_v8;
        let _nw5 = new _module.C0();
        _nw5.__ctor(false);
        _103_v8 = _nw5;
        let _104_v9;
        let _nw6 = Array((new BigNumber(28)).toNumber());
        _nw6[(_dafny.ZERO).toNumber()] = _103_v8;
        _nw6[(_dafny.ONE).toNumber()] = _103_v8;
        _nw6[(new BigNumber(2)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(3)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(4)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(5)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(6)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(7)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(8)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(9)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(10)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(11)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(12)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(13)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(14)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(15)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(16)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(17)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(18)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(19)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(20)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(21)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(22)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(23)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(24)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(25)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(26)).toNumber()] = _103_v8;
        _nw6[(new BigNumber(27)).toNumber()] = _103_v8;
        _104_v9 = _nw6;
        let _105_v10;
        _105_v10 = _dafny.Map.Empty.slice().updateUnsafe(_104_v9,_96_v3);
        let _106_v11;
        _106_v11 = _dafny.Map.Empty.slice().updateUnsafe(_102_v7,(((_105_v10).contains(_104_v9)) ? ((_105_v10).get(_104_v9)) : (_96_v3)));
        _101_v6 = (_101_v6).update(_98_v5, (((_106_v11).contains(_102_v7)) ? ((_106_v11).get(_102_v7)) : (_96_v3)));
        let _107_v12;
        _107_v12 = _dafny.Map.Empty.slice().updateUnsafe(_103_v8.f37,p3);
        _107_v12 = (_107_v12).update(p2, p3);
      } else {
        (globalState).f18 = p2;
        let _108_v13;
        let _nw7 = Array((new BigNumber(10)).toNumber()).fill(false);
        _108_v13 = _nw7;
        let _index5 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_108_v13).length));
        (_108_v13)[_index5] = (true) || (p3);
        let _109_v14;
        _109_v14 = _dafny.Map.Empty.slice().updateUnsafe(true,p1);
        let _110_v15;
        _110_v15 = _dafny.Seq.UnicodeFromString("ofqjjeh");
        let _111_v16;
        _111_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(864));
        let _112_v17;
        _112_v17 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_110_v15).length)), new BigNumber(-368)),(_111_v16).contains(p1));
        let _index6 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_108_v13).length));
        let _rhs4 = (!(_109_v14).contains(_module.__default.fm1(_module.__default.fm42(p3, false, globalState), p1, p1, p3, globalState))) === ((p1).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("oa")).length)));
        let _rhs5 = (((_112_v17).contains(p1)) ? ((_112_v17).get(p1)) : (p2));
        let _lhs5 = _108_v13;
        let _lhs6 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_108_v13).length));
        let _lhs7 = globalState;
        _lhs5[_lhs6] = _rhs4;
        _lhs7.f15 = _rhs5;
        let _113_v18;
        _113_v18 = _dafny.Set.fromElements(p0, p0, _module.__default.fm42(_module.__default.fm19(p0, globalState), p3, globalState));
        _113_v18 = _module.__default.fm57(!(p2), globalState);
        let _114_v19;
        let _nw8 = new _module.C4();
        _nw8.__ctor((_108_v13)[_module.__default.safeIndex(new BigNumber(560), new BigNumber((_108_v13).length))], p1);
        _114_v19 = _nw8;
        let _index7 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_108_v13).length));
        (_108_v13)[_index7] = (_114_v19).f33;
      }
      let _115_v20;
      _115_v20 = _dafny.Set.fromElements(p1, p1, new BigNumber(906));
      let _116_v21;
      _116_v21 = _dafny.Seq.of(new BigNumber((_115_v20).length));
      let _117_v22;
      _117_v22 = _dafny.Set.fromElements(!(p3));
      let _118_v23;
      _118_v23 = _dafny.MultiSet.fromElements(_117_v22, _dafny.Set.fromElements(p2));
      let _119_v24;
      let _nw9 = Array((new BigNumber(12)).toNumber());
      _nw9[(_dafny.ZERO).toNumber()] = _dafny.Seq.IsPrefixOf(_116_v21, _116_v21);
      _nw9[(_dafny.ONE).toNumber()] = false;
      _nw9[(new BigNumber(2)).toNumber()] = p2;
      _nw9[(new BigNumber(3)).toNumber()] = !(p3) || (p2);
      _nw9[(new BigNumber(4)).toNumber()] = p2;
      _nw9[(new BigNumber(5)).toNumber()] = p2;
      _nw9[(new BigNumber(6)).toNumber()] = false;
      _nw9[(new BigNumber(7)).toNumber()] = (_118_v23).IsProperSubsetOf(_118_v23);
      _nw9[(new BigNumber(8)).toNumber()] = p2;
      _nw9[(new BigNumber(9)).toNumber()] = p3;
      _nw9[(new BigNumber(10)).toNumber()] = !(_117_v22).contains(p3);
      _nw9[(new BigNumber(11)).toNumber()] = _module.__default.fm1(p0, p1, new BigNumber(371), p2, globalState);
      _119_v24 = _nw9;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_119_v24).length))) {
        let _120_i1 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_120_i1)) && ((_120_i1).isLessThan(new BigNumber((_119_v24).length))))) {
          (_119_v24)[(_120_i1)] = !(p3);
        }
      }
      (globalState).f15 = ((p3) ? (p3) : (p3));
      let _121_v25;
      let _nw10 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _121_v25 = _nw10;
      _121_v25 = _121_v25;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_121_v25).length))) {
        let _122_i2 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_122_i2)) && ((_122_i2).isLessThan(new BigNumber((_121_v25).length))))) {
          (_121_v25)[(_122_i2)] = _dafny.Seq.UnicodeFromString("mp");
        }
      }
      let _123_v26;
      _123_v26 = _dafny.Seq.of(p3, p3, p3, false, p2);
      let _index8 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_119_v24).length));
      (_119_v24)[_index8] = _module.__default.fm33(p2, new BigNumber((_117_v22).length), new BigNumber((_123_v26).length), globalState);
      let _124_v27;
      let _nw11 = new _module.C6();
      _nw11.__ctor(p2, _dafny.Seq.UnicodeFromString("bydkuf"));
      _124_v27 = _nw11;
      let _125_v28;
      _125_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1,_124_v27);
      let _126_v29;
      _126_v29 = _dafny.Map.Empty.slice().updateUnsafe((((_125_v28).contains(p1)) ? ((_125_v28).get(p1)) : (_124_v27)),p0);
      let _127_v30;
      _127_v30 = _dafny.Seq.UnicodeFromString("juplsu");
      let _index9 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_119_v24).length));
      let _rhs6 = p1;
      let _rhs7 = !_dafny.Seq.contains(_127_v30, (((_126_v29).contains(_124_v27)) ? ((_126_v29).get(_124_v27)) : (p0)));
      let _rhs8 = p1;
      let _lhs8 = globalState;
      let _lhs9 = _119_v24;
      let _lhs10 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_119_v24).length));
      let _lhs11 = globalState;
      _lhs8.f21 = _rhs6;
      _lhs9[_lhs10] = _rhs7;
      _lhs11.f14 = _rhs8;
      return;
    }
    static Main(__noArgsParameter) {
      let _128_v0;
      _128_v0 = true;
      let _129_v1;
      _129_v1 = _dafny.Set.fromElements(_128_v0, _128_v0);
      let _130_v2;
      _130_v2 = _dafny.Seq.of(_dafny.Set.fromElements(_128_v0), _dafny.Set.fromElements(_128_v0), _129_v1, _dafny.Set.fromElements(_128_v0), _129_v1);
      let _131_v3;
      _131_v3 = new BigNumber(-208);
      let _132_v4;
      _132_v4 = _dafny.Set.fromElements(_131_v3);
      let _133_v6;
      _133_v6 = _dafny.MultiSet.fromElements(new BigNumber(137));
      let _134_v7;
      _134_v7 = _dafny.Seq.of(new BigNumber((_132_v4).length), new BigNumber(((function () {
        let _coll23 = new _dafny.Map();
        for (const _compr_23 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-496)), ((_135_v3) => function (_136_i0) {
          return _dafny.MultiSet.fromElements(_135_v3);
        })(_131_v3))).Elements) {
          let _137_v5 = _compr_23;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-496)), ((_138_v3) => function (_136_i0) {
            return _dafny.MultiSet.fromElements(_138_v3);
          })(_131_v3)), _137_v5)) {
            _coll23.push([_137_v5,_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)))]);
          }
        }
        return _coll23;
      }()).update(_133_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(506)), function (_139_i1) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }))).length));
      let _140_v8;
      _140_v8 = _dafny.Seq.UnicodeFromString("rqigbg");
      let _141_v9;
      let _nw12 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _141_v9 = _nw12;
      let _142_v10;
      _142_v10 = _dafny.Map.Empty.slice().updateUnsafe(false,_141_v9);
      let _143_v11;
      _143_v11 = _dafny.Map.Empty.slice().updateUnsafe(_131_v3,(_140_v8)[_module.__default.safeIndex(_131_v3, new BigNumber((_140_v8).length))]);
      let _144_v12;
      _144_v12 = new _dafny.CodePoint('l'.codePointAt(0));
      let _145_v13;
      _145_v13 = _dafny.Map.Empty.slice().updateUnsafe((((_143_v11).contains(_131_v3)) ? ((_143_v11).get(_131_v3)) : (_144_v12)),_128_v0);
      let _146_globalState;
      let _nw13 = new _module.GlobalState();
      _nw13.__ctor(true, new BigNumber(724), new BigNumber(130), _130_v2, true, _dafny.MultiSet.fromElements(_131_v3), new BigNumber(288), _134_v7, new BigNumber(643), new BigNumber(276), new BigNumber(647), _dafny.Seq.Concat(_140_v8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-875)), function (_147_i2) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })), new BigNumber(741), false, new BigNumber(70), true, _142_v10, _141_v9, false, new BigNumber(723), _140_v8, new BigNumber(-758), _145_v13, new BigNumber(863), true, _140_v8, new BigNumber(637), false, true);
      _146_globalState = _nw13;
      if (!(_128_v0)) {
        let _148_v14;
        _148_v14 = _module.D0.create_DC0(((_128_v0) ? (_128_v0) : (false)));
        let _source4 = _148_v14;
        if (_source4.is_DC0) {
          let _149___mcc_h0 = (_source4).cf0;
          let _150_cf0 = _149___mcc_h0;
          (_146_globalState).f20 = _dafny.Seq.Concat(_140_v8, _140_v8);
          let _151_v15;
          let _nw14 = new _module.C4();
          _nw14.__ctor(!(_128_v0), _131_v3);
          _151_v15 = _nw14;
          let _152_v16;
          _152_v16 = _module.D6.create_DC21(_129_v1);
          _152_v16 = _module.D6.create_DC21(_129_v1);
          let _153_v17;
          let _nw15 = Array((new BigNumber(29)).toNumber());
          _153_v17 = _nw15;
          _153_v17 = _153_v17;
        } else if (_source4.is_DC1) {
          let _154___mcc_h1 = (_source4).cf1;
          let _155___mcc_h2 = (_source4).cf2;
          let _156___mcc_h3 = (_source4).cf3;
          let _157_cf3 = _156___mcc_h3;
          let _158_cf2 = _155___mcc_h2;
          let _159_cf1 = _154___mcc_h1;
          _157_cf3 = _module.__default.safeDivisionInt(_157_cf3, _157_cf3);
          let _160_v18;
          let _nw16 = new _module.C6();
          _nw16.__ctor(_128_v0, _140_v8);
          _160_v18 = _nw16;
          let _161_v19;
          _161_v19 = _dafny.MultiSet.fromElements(_160_v18, _160_v18, _160_v18, _160_v18, _160_v18);
          let _162_v20;
          let _nw17 = new _module.C1();
          _nw17.__ctor((_dafny.MultiSet.fromElements(_160_v18)).IsProperSubsetOf(_161_v19));
          _162_v20 = _nw17;
          let _163_v21;
          _163_v21 = _dafny.Seq.of(!(_129_v1).contains(_128_v0), _128_v0, (_160_v18).f29, false);
          let _164_v22;
          _164_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(714)), ((_165_v12) => function (_166_i3) {
            return _165_v12;
          })(_144_v12))).length),!((_160_v18).f29));
          let _167_v23;
          _167_v23 = _module.D15.create_DC43(_164_v22);
          let _pat_let_tv0 = _164_v22;
          let _168_v24;
          let _nw18 = Array((new BigNumber(13)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = _module.D15.create_DC43(_164_v22);
          _nw18[(_dafny.ONE).toNumber()] = _module.D15.create_DC43(_164_v22);
          _nw18[(new BigNumber(2)).toNumber()] = _167_v23;
          _nw18[(new BigNumber(3)).toNumber()] = _167_v23;
          _nw18[(new BigNumber(4)).toNumber()] = _167_v23;
          _nw18[(new BigNumber(5)).toNumber()] = _module.__default.fm54(!((_160_v18).f29), (_module.D2.create_DC9(_dafny.Seq.UnicodeFromString("oegckwcwk"), _131_v3, _128_v0, _158_cf2)).dtor_cf15, _132_v4, (_dafny.ZERO).minus(_158_cf2), _146_globalState);
          _nw18[(new BigNumber(6)).toNumber()] = _module.__default.fm54(_128_v0, _131_v3, _132_v4, new BigNumber(627), _146_globalState);
          _nw18[(new BigNumber(7)).toNumber()] = _167_v23;
          _nw18[(new BigNumber(8)).toNumber()] = _167_v23;
          _nw18[(new BigNumber(9)).toNumber()] = function (_pat_let0_0) {
            return function (_169_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_170_dt__update_hcf70_h0) {
                  return _module.D15.create_DC43(_170_dt__update_hcf70_h0);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_167_v23);
          _nw18[(new BigNumber(10)).toNumber()] = _167_v23;
          _nw18[(new BigNumber(11)).toNumber()] = _167_v23;
          _nw18[(new BigNumber(12)).toNumber()] = _167_v23;
          _168_v24 = _nw18;
          let _index10 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_168_v24).length));
          (_168_v24)[_index10] = _167_v23;
          let _index11 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_168_v24).length));
          let _rhs9 = _162_v20;
          let _rhs10 = ((_128_v0) ? (_163_v21) : (_163_v21));
          let _rhs11 = (((_128_v0) ? (_133_v6) : (_133_v6))).Difference(_133_v6);
          let _rhs12 = _167_v23;
          let _rhs13 = (_158_cf2).multipliedBy(new BigNumber((_129_v1).length));
          let _lhs12 = _168_v24;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_168_v24).length));
          let _lhs14 = _146_globalState;
          _162_v20 = _rhs9;
          _163_v21 = _rhs10;
          _133_v6 = _rhs11;
          _lhs12[_lhs13] = _rhs12;
          _lhs14.f14 = _rhs13;
          (_146_globalState).f14 = (_dafny.ZERO).minus((_158_cf2).minus(_131_v3));
          let _171_v25;
          let _nw19 = Array((new BigNumber(24)).toNumber()).fill([]);
          _171_v25 = _nw19;
          let _index12 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_171_v25).length));
          (_171_v25)[_index12] = _141_v9;
          let _172_v26;
          _172_v26 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,new BigNumber(-824));
          let _173_v27;
          _173_v27 = _dafny.Map.Empty.slice().updateUnsafe((_160_v18).f29,_140_v8);
          let _index13 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_171_v25).length));
          let _rhs14 = _141_v9;
          let _rhs15 = ((((_172_v26).contains(_128_v0)) ? ((_172_v26).get(_128_v0)) : (_131_v3))).minus((_dafny.ZERO).minus(_module.__default.fm36(_146_globalState)));
          let _rhs16 = new BigNumber((_dafny.Seq.Concat(_140_v8, (((_173_v27).contains(true)) ? ((_173_v27).get(true)) : (_160_v18.f30)))).length);
          let _lhs15 = _171_v25;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_171_v25).length));
          let _lhs17 = _146_globalState;
          _lhs15[_lhs16] = _rhs14;
          _131_v3 = _rhs15;
          _lhs17.f21 = _rhs16;
        } else {
          let _174___mcc_h4 = (_source4).cf4;
          let _175___mcc_h5 = (_source4).cf5;
          let _176___mcc_h6 = (_source4).cf6;
          let _177_cf6 = _176___mcc_h6;
          let _178_cf5 = _175___mcc_h5;
          let _179_cf4 = _174___mcc_h4;
          (_146_globalState).f13 = ((new BigNumber((_dafny.Seq.UnicodeFromString("rdwnsce")).length)).multipliedBy(_131_v3)).isEqualTo(new BigNumber((_module.__default.fm44(_177_cf6, _131_v3, !(false), _177_cf6, _146_globalState)).length));
          let _index14 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_141_v9).length));
          (_141_v9)[_index14] = (_dafny.ZERO).minus(_131_v3);
          let _index15 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_141_v9).length));
          (_141_v9)[_index15] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(_131_v3, new BigNumber((_134_v7).length))));
          let _180_v28;
          let _nw20 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _180_v28 = _nw20;
          let _181_v29;
          _181_v29 = _dafny.Map.Empty.slice().updateUnsafe(_180_v28,false);
          _181_v29 = (_181_v29).Merge(_181_v29);
          (_146_globalState).f1 = _131_v3;
        }
        let _182_v30;
        let _nw21 = Array((new BigNumber(15)).toNumber()).fill([]);
        _182_v30 = _nw21;
        let _183_v31;
        let _nw22 = Array((_dafny.ONE).toNumber()).fill([]);
        _183_v31 = _nw22;
        let _index16 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_182_v30).length));
        (_182_v30)[_index16] = _183_v31;
        let _184_v32;
        let _init0 = ((_185_v3) => function (_186_i4) {
          return _dafny.Map.Empty.slice().updateUnsafe(_185_v3,!(false));
        })(_131_v3);
        let _nw23 = Array((new BigNumber(17)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw23.length); _i0_0++) {
          _nw23[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _184_v32 = _nw23;
        let _187_v33;
        _187_v33 = _dafny.Map.Empty.slice().updateUnsafe(_131_v3,_128_v0);
        let _index17 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_184_v32).length));
        (_184_v32)[_index17] = _187_v33;
        let _188_v34;
        _188_v34 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,(_dafny.ZERO).minus(_131_v3));
        let _189_v35;
        _189_v35 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_128_v0,_131_v3));
        let _190_v36;
        _190_v36 = _dafny.Seq.of(_188_v34, _188_v34, (_189_v35)[_module.__default.safeIndex(new BigNumber(284), new BigNumber((_189_v35).length))]);
        let _191_v37;
        _191_v37 = _dafny.Map.Empty.slice().updateUnsafe(true,_140_v8);
        let _index18 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_182_v30).length));
        let _index19 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_184_v32).length));
        let _rhs17 = _183_v31;
        let _rhs18 = true;
        let _rhs19 = _dafny.Map.Empty.slice().updateUnsafe(_131_v3,_128_v0);
        let _rhs20 = _dafny.Seq.Concat(_190_v36, _190_v36);
        let _rhs21 = (((_191_v37).contains(_128_v0)) ? ((_191_v37).get(_128_v0)) : (_140_v8));
        let _lhs18 = _182_v30;
        let _lhs19 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_182_v30).length));
        let _lhs20 = _146_globalState;
        let _lhs21 = _184_v32;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_184_v32).length));
        let _lhs23 = _146_globalState;
        _lhs18[_lhs19] = _rhs17;
        _lhs20.f13 = _rhs18;
        _lhs21[_lhs22] = _rhs19;
        _190_v36 = _rhs20;
        _lhs23.f11 = _rhs21;
        let _192_v38;
        let _nw24 = new _module.C6();
        _nw24.__ctor(_128_v0, _dafny.Seq.Concat(_140_v8, _dafny.Seq.UnicodeFromString("onhi")));
        _192_v38 = _nw24;
        let _nw25 = new _module.C6();
        _nw25.__ctor(!(!(_128_v0)), _dafny.Seq.UnicodeFromString("aidqxkrn"));
        _192_v38 = _nw25;
        (_146_globalState).f19 = _131_v3;
        let _index20 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_141_v9).length));
        (_141_v9)[_index20] = new BigNumber((_dafny.Seq.Concat(_134_v7, _134_v7)).length);
        let _index21 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_141_v9).length));
        (_141_v9)[_index21] = _131_v3;
      } else {
        let _193_v39;
        _193_v39 = _dafny.Seq.of(_128_v0);
        let _194_v40;
        let _nw26 = new _module.C2();
        _nw26.__ctor(_131_v3, _dafny.Set.fromElements(_141_v9), !((_193_v39)[_module.__default.safeIndex(_131_v3, new BigNumber((_193_v39).length))]), _140_v8);
        _194_v40 = _nw26;
        let _195_v41;
        _195_v41 = _dafny.MultiSet.fromElements(_140_v8, _140_v8);
        _195_v41 = _module.__default.fm55(_146_globalState);
        let _196_v42;
        let _nw27 = new _module.C0();
        _nw27.__ctor(_128_v0);
        _196_v42 = _nw27;
        let _index22 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_141_v9).length));
        (_141_v9)[_index22] = new BigNumber((_dafny.MultiSet.fromElements(_128_v0, _196_v42.f37, _196_v42.f37)).cardinality());
        let _index23 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_141_v9).length));
        let _rhs22 = (_module.__default.fm27(new BigNumber((_133_v6).cardinality()), _196_v42.f37, _146_globalState)).Union(_dafny.MultiSet.fromElements(new BigNumber(-916), _131_v3));
        let _rhs23 = (new BigNumber((_module.__default.fm56(_146_globalState)).length)).minus((_194_v40).f38);
        let _rhs24 = (_194_v40).f38;
        let _rhs25 = ((!(_196_v42.f37)) ? (true) : (_128_v0));
        let _lhs24 = _146_globalState;
        let _lhs25 = _141_v9;
        let _lhs26 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_141_v9).length));
        let _lhs27 = _146_globalState;
        _133_v6 = _rhs22;
        _lhs24.f2 = _rhs23;
        _lhs25[_lhs26] = _rhs24;
        _lhs27.f15 = _rhs25;
        let _197_v43;
        let _nw28 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _197_v43 = _nw28;
        let _198_v44;
        _198_v44 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,_197_v43);
        let _199_v45;
        let _200_v46;
        let _out0;
        let _out1;
        let _outcollector0 = (_194_v40).m3((_dafny.Set.fromElements(_196_v42.f37, true)).IsDisjointFrom(_129_v1), (((_198_v44).contains(_128_v0)) ? ((_198_v44).get(_128_v0)) : (_197_v43)), _131_v3, _146_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _199_v45 = _out0;
        _200_v46 = _out1;
      }
      let _201_v47;
      _201_v47 = _module.D2.create_DC8(_144_v12, !(_module.__default.fm1(_144_v12, new BigNumber(615), new BigNumber((_134_v7).length), _128_v0, _146_globalState)));
      let _202_v48;
      _202_v48 = _dafny.Seq.of(_201_v47, _201_v47, _201_v47, _201_v47, _201_v47);
      let _203_v49;
      let _init1 = function (_204_i5) {
        return true;
      };
      let _nw29 = Array((new BigNumber(27)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw29.length); _i0_1++) {
        _nw29[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _203_v49 = _nw29;
      let _205_v50;
      _205_v50 = _module.D9.create_DC28(_module.__default.fm35(_202_v48, new BigNumber(-251), _146_globalState), _128_v0, _131_v3, (_131_v3).isEqualTo((_module.D0.create_DC1(_203_v49, new BigNumber((_140_v8).length), _131_v3)).dtor_cf3), _dafny.Seq.IsProperPrefixOf(_134_v7, _dafny.Seq.Create(_module.__default.abs(new BigNumber(60)), ((_206_v3) => function (_207_i6) {
  return _206_v3;
})(_131_v3))));
      let _source5 = _205_v50;
      if (_source5.is_DC28) {
        let _208___mcc_h7 = (_source5).cf41;
        let _209___mcc_h8 = (_source5).cf42;
        let _210___mcc_h9 = (_source5).cf43;
        let _211___mcc_h10 = (_source5).cf44;
        let _212___mcc_h11 = (_source5).cf45;
        let _213_cf45 = _212___mcc_h11;
        let _214_cf44 = _211___mcc_h10;
        let _215_cf43 = _210___mcc_h9;
        let _216_cf42 = _209___mcc_h8;
        let _217_cf41 = _208___mcc_h7;
        let _218_v51;
        let _nw30 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
        _218_v51 = _nw30;
        let _index24 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_218_v51).length));
        (_218_v51)[_index24] = _133_v6;
        let _index25 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_218_v51).length));
        (_218_v51)[_index25] = _module.__default.fm27(new BigNumber((_dafny.Seq.Concat(_140_v8, _140_v8)).length), true, _146_globalState);
        (_146_globalState).f13 = _214_cf44;
        _131_v3 = _215_cf43;
        let _219_v52;
        let _nw31 = new _module.C7();
        _nw31.__ctor(_213_cf45, _213_cf45, _213_cf45, _dafny.Seq.UnicodeFromString("fnfnj"));
        _219_v52 = _nw31;
        let _220_v53;
        _220_v53 = _dafny.Set.fromElements(_219_v52, _219_v52, _219_v52, _219_v52, _219_v52);
        let _221_v54;
        _221_v54 = _dafny.Set.fromElements(_220_v53, _220_v53, _220_v53, _220_v53, _220_v53);
        let _222_v55;
        _222_v55 = _dafny.Map.Empty.slice().updateUnsafe(false,_221_v54);
        let _223_v56;
        _223_v56 = _module.D19.create_DC56(_221_v54);
        let _rhs26 = ((_214_cf44) ? ((((_222_v55).contains(_213_cf45)) ? ((_222_v55).get(_213_cf45)) : (_dafny.Set.fromElements(_dafny.Set.fromElements(_219_v52))))) : ((_223_v56).dtor_cf99));
        let _rhs27 = (new BigNumber(((_dafny.Set.fromElements(true, _128_v0, _216_cf42)).Difference(_dafny.Set.fromElements(_216_cf42, _128_v0, _214_cf44))).length)).multipliedBy(_215_cf43);
        let _lhs28 = _146_globalState;
        _221_v54 = _rhs26;
        _lhs28.f14 = _rhs27;
      } else if (_source5.is_DC29) {
        let _224___mcc_h12 = (_source5).cf46;
        let _225___mcc_h13 = (_source5).cf47;
        let _226___mcc_h14 = (_source5).cf48;
        let _227_cf48 = _226___mcc_h14;
        let _228_cf47 = _225___mcc_h13;
        let _229_cf46 = _224___mcc_h12;
        _module.__default.m10(new _dafny.CodePoint('x'.codePointAt(0)), _229_cf46, ((true) ? (_227_cf48) : (_227_cf48)), (_131_v3).isLessThanOrEqualTo(new BigNumber(459)), _146_globalState);
        let _230_v57;
        let _nw32 = new _module.C7();
        _nw32.__ctor(_128_v0, _128_v0, _128_v0, _140_v8);
        _230_v57 = _nw32;
        let _231_v58;
        _231_v58 = _dafny.Set.fromElements(_230_v57, _230_v57);
        let _index26 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_141_v9).length));
        (_141_v9)[_index26] = _module.__default.fm0(new BigNumber((_231_v58).length), (_230_v57).f31, _227_cf48, _146_globalState);
        let _index27 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_141_v9).length));
        (_141_v9)[_index27] = _131_v3;
        let _232_v59;
        _232_v59 = _dafny.Map.Empty.slice().updateUnsafe(_134_v7,new BigNumber((_dafny.Seq.UnicodeFromString("wegckpd")).length));
        let _233_v60;
        _233_v60 = _module.D11.create_DC34(_232_v59);
        _233_v60 = _233_v60;
        (_146_globalState).f21 = (_141_v9)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_141_v9).length))];
      } else {
        let _234___mcc_h15 = (_source5).cf40;
        let _235_cf40 = _234___mcc_h15;
        if (_128_v0) {
          let _236_v61;
          let _init2 = ((_237_v8, _238_v3) => function (_239_i7) {
            return _dafny.Seq.of(_237_v8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(225)), ((_240_v8, _241_v3) => function (_242_i8) {
              return (_240_v8)[_module.__default.safeIndex(_241_v3, new BigNumber((_240_v8).length))];
            })(_237_v8, _238_v3)));
          })(_140_v8, _131_v3);
          let _nw33 = Array((new BigNumber(18)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw33.length); _i0_2++) {
            _nw33[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _236_v61 = _nw33;
          let _243_v62;
          let _nw34 = new _module.C6();
          _nw34.__ctor(_128_v0, _140_v8);
          _243_v62 = _nw34;
          let _244_v63;
          _244_v63 = _dafny.Map.Empty.slice().updateUnsafe(_243_v62,_dafny.Seq.Create(_module.__default.abs(new BigNumber(628)), ((_245_v12) => function (_246_i9) {
            return _245_v12;
          })(_144_v12)));
          let _247_v64;
          _247_v64 = _dafny.Seq.of(_243_v62, _243_v62);
          let _248_v65;
          _248_v65 = _dafny.Seq.of((((_244_v63).contains((_247_v64)[_module.__default.safeIndex(_131_v3, new BigNumber((_247_v64).length))])) ? ((_244_v63).get((_247_v64)[_module.__default.safeIndex(_131_v3, new BigNumber((_247_v64).length))])) : (_140_v8)), _140_v8, _140_v8, _dafny.Seq.UnicodeFromString("elrvfpxmc"));
          let _index28 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_236_v61).length));
          (_236_v61)[_index28] = _248_v65;
          let _249_v66;
          let _nw35 = new _module.C0();
          _nw35.__ctor(_128_v0);
          _249_v66 = _nw35;
          let _250_v67;
          _250_v67 = _dafny.Seq.of(_249_v66, _249_v66, _249_v66, _249_v66, _249_v66);
          let _index29 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_236_v61).length));
          let _rhs28 = _131_v3;
          let _rhs29 = _dafny.Seq.Concat(_dafny.Seq.Concat(_140_v8, _140_v8), _140_v8);
          let _rhs30 = _248_v65;
          let _rhs31 = (_250_v67)[_module.__default.safeIndex(_131_v3, new BigNumber((_250_v67).length))];
          let _lhs29 = _146_globalState;
          let _lhs30 = _236_v61;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_236_v61).length));
          _131_v3 = _rhs28;
          _lhs29.f20 = _rhs29;
          _lhs30[_lhs31] = _rhs30;
          _249_v66 = _rhs31;
          let _251_v68;
          _251_v68 = _module.D0.create_DC1(_203_v49, new BigNumber(-974), _131_v3);
          _251_v68 = _251_v68;
          let _252_v69;
          _252_v69 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_140_v8).length));
          let _253_v70;
          let _nw36 = new _module.C3();
          _nw36.__ctor((_module.__default.fm9(_146_globalState)).isLessThanOrEqualTo(_131_v3), (new BigNumber(579)).minus((((_252_v69).contains(_249_v66.f37)) ? ((_252_v69).get(_249_v66.f37)) : (_131_v3))), _249_v66.f37, _dafny.Seq.Concat(_140_v8, _dafny.Seq.UnicodeFromString("ogivdrrsy")));
          _253_v70 = _nw36;
          (_146_globalState).f2 = new BigNumber(750);
          let _254_v71;
          let _nw37 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _254_v71 = _nw37;
          let _index30 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_254_v71).length));
          (_254_v71)[_index30] = _140_v8;
          let _index31 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_254_v71).length));
          (_254_v71)[_index31] = _dafny.Seq.Concat(_140_v8, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("ohllvco"), _module.__default.safeIndex((_253_v70).f36, new BigNumber((_dafny.Seq.UnicodeFromString("ohllvco")).length)), _144_v12));
        } else {
          (_146_globalState).f7 = _235_cf40;
          let _255_v72;
          _255_v72 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,_144_v12);
          let _256_v73;
          _256_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_140_v8).length),_128_v0);
          let _257_v74;
          let _nw38 = new _module.C6();
          _nw38.__ctor(_128_v0, _140_v8);
          _257_v74 = _nw38;
          let _258_v75;
          _258_v75 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,_257_v74);
          let _259_v76;
          let _nw39 = Array((new BigNumber(25)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = _144_v12;
          _nw39[(_dafny.ONE).toNumber()] = _144_v12;
          _nw39[(new BigNumber(2)).toNumber()] = (((_255_v72).contains(_128_v0)) ? ((_255_v72).get(_128_v0)) : (_144_v12));
          _nw39[(new BigNumber(3)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(4)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(5)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(6)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(7)).toNumber()] = _module.__default.fm42(_128_v0, _128_v0, _146_globalState);
          _nw39[(new BigNumber(8)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(9)).toNumber()] = _module.__default.fm42((((_256_v73).contains(new BigNumber((_258_v75).length))) ? ((_256_v73).get(new BigNumber((_258_v75).length))) : (false)), false, _146_globalState);
          _nw39[(new BigNumber(10)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(11)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(12)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(13)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
          _nw39[(new BigNumber(14)).toNumber()] = ((_128_v0) ? (_144_v12) : (_144_v12));
          _nw39[(new BigNumber(15)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(16)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(17)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(18)).toNumber()] = _module.__default.fm42(_128_v0, _128_v0, _146_globalState);
          _nw39[(new BigNumber(19)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(20)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(21)).toNumber()] = _module.__default.fm42(_128_v0, _128_v0, _146_globalState);
          _nw39[(new BigNumber(22)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(23)).toNumber()] = _144_v12;
          _nw39[(new BigNumber(24)).toNumber()] = _144_v12;
          _259_v76 = _nw39;
          let _index32 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_259_v76).length));
          (_259_v76)[_index32] = _module.__default.fm42(false, _128_v0, _146_globalState);
          let _index33 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_203_v49).length));
          (_203_v49)[_index33] = _128_v0;
          let _index34 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_259_v76).length));
          let _index35 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_203_v49).length));
          let _rhs32 = !(_131_v3).isEqualTo(_131_v3);
          let _rhs33 = new _dafny.CodePoint('j'.codePointAt(0));
          let _rhs34 = (_dafny.ZERO).minus(_131_v3);
          let _rhs35 = _128_v0;
          let _rhs36 = _128_v0;
          let _lhs32 = _146_globalState;
          let _lhs33 = _259_v76;
          let _lhs34 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_259_v76).length));
          let _lhs35 = _146_globalState;
          let _lhs36 = _203_v49;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_203_v49).length));
          let _lhs38 = _146_globalState;
          _lhs32.f15 = _rhs32;
          _lhs33[_lhs34] = _rhs33;
          _lhs35.f2 = _rhs34;
          _lhs36[_lhs37] = _rhs35;
          _lhs38.f15 = _rhs36;
          let _260_v77;
          let _nw40 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _260_v77 = _nw40;
          let _261_v78;
          _261_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_235_cf40, _module.__default.safeIndex(_131_v3, new BigNumber((_235_cf40).length)), _131_v3)).length),_260_v77);
          _261_v78 = (_261_v78).update(new BigNumber(777), _260_v77);
          let _262_v79;
          let _nw41 = new _module.C5();
          _nw41.__ctor(_128_v0, _140_v8);
          _262_v79 = _nw41;
          let _263_v80;
          _263_v80 = _dafny.MultiSet.fromElements((_203_v49)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_203_v49).length))]);
          (_146_globalState).f13 = (((_145_v13).contains(_144_v12)) ? ((_145_v13).get(_144_v12)) : (((_263_v80).update(_128_v0, _module.__default.abs(_131_v3))).IsSubsetOf(_dafny.MultiSet.fromElements(_128_v0, _128_v0, _128_v0))));
        }
        _module.__default.m10(_144_v12, _131_v3, (_module.__default.fm0(_131_v3, _128_v0, _128_v0, _146_globalState)).isLessThanOrEqualTo(new BigNumber(-976)), _128_v0, _146_globalState);
        let _264_v81;
        _264_v81 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,new BigNumber(241));
        let _265_v82;
        _265_v82 = _dafny.Seq.of(_203_v49, _203_v49, _203_v49);
        let _rhs37 = (_265_v82)[_module.__default.safeIndex(_131_v3, new BigNumber((_265_v82).length))];
        let _rhs38 = _144_v12;
        let _rhs39 = _264_v81;
        _203_v49 = _rhs37;
        _144_v12 = _rhs38;
        _264_v81 = _rhs39;
        _module.__default.m10(_144_v12, (_131_v3).minus((_dafny.ZERO).minus(_131_v3)), !(_132_v4).contains(_131_v3), _128_v0, _146_globalState);
      }
      (_146_globalState).f14 = (_131_v3).multipliedBy((_131_v3).multipliedBy(new BigNumber(302)));
      if ((_131_v3).isEqualTo(_131_v3)) {
        _131_v3 = _131_v3;
        let _index36 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_141_v9).length));
        (_141_v9)[_index36] = _131_v3;
        let _266_v83;
        _266_v83 = _dafny.Map.Empty.slice().updateUnsafe(_131_v3,new BigNumber(121));
        let _267_v84;
        _267_v84 = _module.D10.create_DC32(new _dafny.CodePoint('k'.codePointAt(0)), _128_v0);
        let _pat_let_tv1 = _128_v0;
        let _268_v85;
        let _nw42 = Array((new BigNumber(12)).toNumber());
        _nw42[(_dafny.ZERO).toNumber()] = _267_v84;
        _nw42[(_dafny.ONE).toNumber()] = _267_v84;
        _nw42[(new BigNumber(2)).toNumber()] = _267_v84;
        _nw42[(new BigNumber(3)).toNumber()] = _module.D10.create_DC32(_144_v12, _128_v0);
        _nw42[(new BigNumber(4)).toNumber()] = function (_pat_let2_0) {
          return function (_269_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_270_dt__update_hcf53_h0) {
                return _module.D10.create_DC32((_269_dt__update__tmp_h1).dtor_cf52, _270_dt__update_hcf53_h0);
              }(_pat_let3_0);
            }(_pat_let_tv1);
          }(_pat_let2_0);
        }(_267_v84);
        _nw42[(new BigNumber(5)).toNumber()] = _267_v84;
        _nw42[(new BigNumber(6)).toNumber()] = _module.D10.create_DC32(_144_v12, _128_v0);
        _nw42[(new BigNumber(7)).toNumber()] = _267_v84;
        _nw42[(new BigNumber(8)).toNumber()] = _module.D10.create_DC32(_144_v12, _128_v0);
        _nw42[(new BigNumber(9)).toNumber()] = _267_v84;
        _nw42[(new BigNumber(10)).toNumber()] = _module.D10.create_DC32(_144_v12, _128_v0);
        _nw42[(new BigNumber(11)).toNumber()] = _267_v84;
        _268_v85 = _nw42;
        let _271_v86;
        let _nw43 = new _module.C4();
        _nw43.__ctor(!(_128_v0), (_module.D15.create_DC44(_131_v3, _128_v0, _268_v85, _141_v9, (_dafny.ZERO).minus(_131_v3))).dtor_cf71);
        _271_v86 = _nw43;
        let _272_v87;
        _272_v87 = _dafny.Map.Empty.slice().updateUnsafe(_271_v86,_128_v0);
        let _index37 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_141_v9).length));
        (_141_v9)[_index37] = ((_dafny.ZERO).minus((((_266_v83).contains(_131_v3)) ? ((_266_v83).get(_131_v3)) : (new BigNumber((_129_v1).length))))).minus(new BigNumber((_272_v87).length));
        let _273_v88;
        let _nw44 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _273_v88 = _nw44;
        _273_v88 = _273_v88;
        let _274_v89;
        let _275_v90;
        let _out2;
        let _out3;
        let _outcollector1 = (_271_v86).m2(_128_v0, _146_globalState);
        _out2 = _outcollector1[0];
        _out3 = _outcollector1[1];
        _274_v89 = _out2;
        _275_v90 = _out3;
        let _276_v91;
        let _nw45 = new _module.C6();
        _nw45.__ctor((_271_v86).f33, _140_v8);
        _276_v91 = _nw45;
        _276_v91 = _276_v91;
      } else {
        (_146_globalState).f15 = (_128_v0) === (_128_v0);
        let _277_v92;
        _277_v92 = _dafny.Seq.of(!(_131_v3).isEqualTo(new BigNumber(-811)));
        _131_v3 = new BigNumber((_277_v92).length);
        _144_v12 = _144_v12;
        if (_128_v0) {
          (_146_globalState).f1 = (((_133_v6).contains(_module.__default.safeModuloInt(_131_v3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(394)), ((_280_v12) => function (_281_i10) {
            return _280_v12;
          })(_144_v12))).length)))) ? ((_133_v6).get(_module.__default.safeModuloInt(_131_v3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(394)), ((_278_v12) => function (_279_i10) {
            return _278_v12;
          })(_144_v12))).length)))) : (new BigNumber(660)));
          let _282_v93;
          let _nw46 = new _module.C5();
          _nw46.__ctor(_128_v0, _140_v8);
          _282_v93 = _nw46;
          let _283_v94;
          let _nw47 = new _module.C7();
          _nw47.__ctor(_128_v0, _128_v0, _128_v0, _140_v8);
          _283_v94 = _nw47;
          (_146_globalState).f7 = _134_v7;
          let _rhs40 = (_131_v3).plus((((_133_v6).contains(_131_v3)) ? ((_133_v6).get(_131_v3)) : (_131_v3)));
          let _rhs41 = _131_v3;
          let _lhs39 = _146_globalState;
          let _lhs40 = _146_globalState;
          _lhs39.f19 = _rhs40;
          _lhs40.f14 = _rhs41;
        } else {
          let _index38 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_203_v49).length));
          (_203_v49)[_index38] = _module.__default.fm6(_131_v3, _146_globalState);
          let _index39 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_203_v49).length));
          (_203_v49)[_index39] = !(_128_v0) || (_128_v0);
          let _284_v95;
          let _nw48 = Array((new BigNumber(5)).toNumber());
          _nw48[(_dafny.ZERO).toNumber()] = _128_v0;
          _nw48[(_dafny.ONE).toNumber()] = (_203_v49)[_module.__default.safeIndex(new BigNumber(858), new BigNumber((_203_v49).length))];
          _nw48[(new BigNumber(2)).toNumber()] = (_203_v49)[_module.__default.safeIndex(new BigNumber(858), new BigNumber((_203_v49).length))];
          _nw48[(new BigNumber(3)).toNumber()] = true;
          _nw48[(new BigNumber(4)).toNumber()] = !(_128_v0);
          _284_v95 = _nw48;
          let _285_v96;
          _285_v96 = _dafny.Seq.of(_284_v95, _284_v95);
          (_146_globalState).f28 = _dafny.Seq.contains(_285_v96, _284_v95);
          let _286_v97;
          _286_v97 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(825),_128_v0);
          _module.__default.m10(((_module.__default.fm19(_144_v12, _146_globalState)) ? (_144_v12) : (_144_v12)), _131_v3, _module.__default.fm33((_203_v49)[_module.__default.safeIndex(new BigNumber(858), new BigNumber((_203_v49).length))], _131_v3, new BigNumber((_286_v97).length), _146_globalState), (_module.__default.fm36(_146_globalState)).isLessThan(_131_v3), _146_globalState);
          let _287_v98;
          _287_v98 = _dafny.Set.fromElements(_141_v9);
          let _288_v99;
          let _nw49 = new _module.C2();
          _nw49.__ctor(_131_v3, _287_v98, _128_v0, _140_v8);
          _288_v99 = _nw49;
          let _289_v100;
          _289_v100 = _dafny.Set.fromElements(_288_v99, _288_v99);
          let _290_v101;
          _290_v101 = _dafny.Map.Empty.slice().updateUnsafe((((_203_v49)[_module.__default.safeIndex(new BigNumber(858), new BigNumber((_203_v49).length))]) ? (_140_v8) : (_module.__default.fm20(_128_v0, (_203_v49)[_module.__default.safeIndex(new BigNumber(858), new BigNumber((_203_v49).length))], _277_v92, new BigNumber((_289_v100).length), _146_globalState))),_134_v7);
          _290_v101 = (_290_v101).update(_dafny.Seq.Concat(_140_v8, _140_v8), _dafny.Seq.Concat(_module.__default.fm44((_203_v49)[_module.__default.safeIndex(new BigNumber(858), new BigNumber((_203_v49).length))], (((_133_v6).contains((_288_v99).f38)) ? ((_133_v6).get((_288_v99).f38)) : (_131_v3)), _128_v0, _128_v0, _146_globalState), _134_v7));
          let _291_v102;
          let _init3 = function (_292_i12) {
            return _dafny.Seq.UnicodeFromString("ughfu");
          };
          let _nw50 = Array((new BigNumber(28)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw50.length); _i0_3++) {
            _nw50[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _291_v102 = _nw50;
          let _293_v103;
          let _294_v104;
          let _out4;
          let _out5;
          let _outcollector2 = (_288_v99).m3(_module.__default.fm6(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(502)), ((_295_v8, _296_v99) => function (_297_i11) {
            return (_295_v8)[_module.__default.safeIndex((_296_v99).f38, new BigNumber((_295_v8).length))];
          })(_140_v8, _288_v99))).length), _146_globalState), _291_v102, (_131_v3).plus((_288_v99).f38), _146_globalState);
          _out4 = _outcollector2[0];
          _out5 = _outcollector2[1];
          _293_v103 = _out4;
          _294_v104 = _out5;
        }
        (_146_globalState).f13 = _128_v0;
      }
      let _298_v105;
      let _nw51 = Array((new BigNumber(18)).toNumber());
      _298_v105 = _nw51;
      let _299_v106;
      let _nw52 = new _module.C7();
      _nw52.__ctor(_128_v0, _128_v0, _128_v0, _140_v8);
      _299_v106 = _nw52;
      let _index40 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_298_v105).length));
      (_298_v105)[_index40] = _299_v106;
      let _300_v107;
      _300_v107 = _dafny.Seq.of(_299_v106);
      let _index41 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_298_v105).length));
      (_298_v105)[_index41] = (_300_v107)[_module.__default.safeIndex(_module.__default.fm9(_146_globalState), new BigNumber((_300_v107).length))];
      _203_v49 = _203_v49;
      let _index42 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_203_v49).length));
      (_203_v49)[_index42] = _128_v0;
      let _301_v108;
      _301_v108 = _dafny.Seq.of(_128_v0, (_299_v106).f29, _128_v0);
      let _302_v109;
      _302_v109 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_144_v12, new BigNumber((_301_v108).length), _131_v3, true, _146_globalState),(_299_v106).f29);
      let _index43 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_203_v49).length));
      (_203_v49)[_index43] = !((((_302_v109).contains((_299_v106).f29)) ? ((_302_v109).get((_299_v106).f29)) : ((_299_v106).f29)));
      (_146_globalState).f28 = (_299_v106).f29;
      let _303_v110;
      let _nw53 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
      _303_v110 = _nw53;
      let _304_v111;
      _304_v111 = _dafny.Map.Empty.slice().updateUnsafe(true,_131_v3);
      let _305_v112;
      let _nw54 = new _module.C5();
      _nw54.__ctor(((_dafny.ZERO).minus(new BigNumber(((_304_v111).update(_128_v0, _131_v3)).length))).isLessThanOrEqualTo(_131_v3), _140_v8);
      _305_v112 = _nw54;
      let _306_v114;
      _306_v114 = _module.D17.create_DC52((_299_v106).f29, (_203_v49)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_203_v49).length))], _131_v3, _133_v6);
      let _pat_let_tv2 = _131_v3;
      let _pat_let_tv3 = _299_v106;
      let _pat_let_tv4 = _131_v3;
      let _pat_let_tv5 = _304_v111;
      let _pat_let_tv6 = _131_v3;
      let _pat_let_tv7 = _301_v108;
      let _pat_let_tv8 = _301_v108;
      let _rhs42 = _303_v110;
      let _rhs43 = !(((_132_v4).Intersect(function () {
        let _coll24 = new _dafny.Set();
        for (const _compr_24 of (_134_v7).Elements) {
          let _307_v113 = _compr_24;
          if (_dafny.Seq.contains(_134_v7, _307_v113)) {
            _coll24.add((_307_v113).plus(new BigNumber((_dafny.Seq.UnicodeFromString("jm")).length)));
          }
        }
        return _coll24;
      }())).IsSubsetOf(_dafny.Set.fromElements(_131_v3, new BigNumber((_dafny.Set.fromElements(true)).length))));
      let _rhs44 = (_299_v106).f29;
      let _rhs45 = function (_source6) {
        if (_source6.is_DC50) {
          let _308___mcc_h16 = (_source6).cf83;
          let _309___mcc_h17 = (_source6).cf84;
          let _310___mcc_h18 = (_source6).cf85;
          let _311_cf85 = _310___mcc_h18;
          let _312_cf84 = _309___mcc_h17;
          let _313_cf83 = _308___mcc_h16;
          return _module.__default.safeDivisionInt(_pat_let_tv2, new BigNumber((_pat_let_tv3.f30).length));
        } else if (_source6.is_DC51) {
          let _314___mcc_h19 = (_source6).cf86;
          let _315___mcc_h20 = (_source6).cf87;
          let _316_cf87 = _315___mcc_h20;
          let _317_cf86 = _314___mcc_h19;
          return _317_cf86;
        } else if (_source6.is_DC52) {
          let _318___mcc_h21 = (_source6).cf88;
          let _319___mcc_h22 = (_source6).cf89;
          let _320___mcc_h23 = (_source6).cf90;
          let _321___mcc_h24 = (_source6).cf91;
          let _322_cf91 = _321___mcc_h24;
          let _323_cf90 = _320___mcc_h23;
          let _324_cf89 = _319___mcc_h22;
          let _325_cf88 = _318___mcc_h21;
          return (_pat_let_tv4).multipliedBy(new BigNumber((_pat_let_tv5).length));
        } else if (_source6.is_DC49) {
          let _326___mcc_h25 = (_source6).cf82;
          let _327_cf82 = _326___mcc_h25;
          return (_dafny.ZERO).minus(_pat_let_tv6);
        } else {
          let _328___mcc_h26 = (_source6).cf92;
          let _329_cf92 = _328___mcc_h26;
          return new BigNumber((_dafny.MultiSet.FromArray(((true) ? (_pat_let_tv7) : (_pat_let_tv8)))).cardinality());
        }
      }(_306_v114);
      let _rhs46 = ((!(_128_v0) || ((_203_v49)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_203_v49).length))])) ? ((((_299_v106).f29) ? (_305_v112) : (_305_v112))) : (_305_v112));
      let _lhs41 = _146_globalState;
      let _lhs42 = _146_globalState;
      _303_v110 = _rhs42;
      _128_v0 = _rhs43;
      _lhs41.f18 = _rhs44;
      _lhs42.f19 = _rhs45;
      _305_v112 = _rhs46;
      let _330_v115;
      let _init4 = ((_331_v8) => function (_332_i13) {
        return _331_v8;
      })(_140_v8);
      let _nw55 = Array((new BigNumber(27)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw55.length); _i0_4++) {
        _nw55[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _330_v115 = _nw55;
      let _333_v116;
      let _nw56 = new _module.C7();
      _nw56.__ctor((_299_v106).f29, (_203_v49)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_203_v49).length))], (_299_v106).f29, _299_v106.f30);
      _333_v116 = _nw56;
      let _334_v117;
      _334_v117 = _module.D20.create_DC61(_dafny.Seq.of(_333_v116));
      let _335_v118;
      let _336_v119;
      let _out6;
      let _out7;
      let _outcollector3 = (_305_v112).m3((_299_v106).f29, _330_v115, (new BigNumber(((_334_v117).dtor_cf111).length)).minus(_module.__default.fm36(_146_globalState)), _146_globalState);
      _out6 = _outcollector3[0];
      _out7 = _outcollector3[1];
      _335_v118 = _out6;
      _336_v119 = _out7;
      let _337_i14;
      _337_i14 = _dafny.ZERO;
      L0: {
        while ((new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_333_v116).f31,_131_v3))).length)).isLessThanOrEqualTo((_305_v112).fm4(_146_globalState))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_337_i14)) {
              break L0;
            }
            _337_i14 = (_337_i14).plus(_dafny.ONE);
            (_146_globalState).f17 = _141_v9;
            _299_v106 = _299_v106;
            let _338_v120;
            let _nw57 = Array((new BigNumber(23)).toNumber());
            _338_v120 = _nw57;
            _338_v120 = _338_v120;
            let _index44 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_141_v9).length));
            (_141_v9)[_index44] = (_131_v3).multipliedBy(_131_v3);
            let _index45 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_141_v9).length));
            let _rhs47 = new _dafny.CodePoint('k'.codePointAt(0));
            let _rhs48 = _131_v3;
            let _lhs43 = _141_v9;
            let _lhs44 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_141_v9).length));
            _144_v12 = _rhs47;
            _lhs43[_lhs44] = _rhs48;
          }
        }
      }
      let _339_v121;
      let _nw58 = new _module.C4();
      _nw58.__ctor(_336_v119, _131_v3);
      _339_v121 = _nw58;
      _339_v121 = _339_v121;
      let _340_v122;
      _340_v122 = _module.D18.create_DC55(_144_v12, _128_v0, (_339_v121).f34, _141_v9, _336_v119);
      let _source7 = _340_v122;
      if (_source7.is_DC55) {
        let _341___mcc_h27 = (_source7).cf94;
        let _342___mcc_h28 = (_source7).cf95;
        let _343___mcc_h29 = (_source7).cf96;
        let _344___mcc_h30 = (_source7).cf97;
        let _345___mcc_h31 = (_source7).cf98;
        let _346_cf98 = _345___mcc_h31;
        let _347_cf97 = _344___mcc_h30;
        let _348_cf96 = _343___mcc_h29;
        let _349_cf95 = _342___mcc_h28;
        let _350_cf94 = _341___mcc_h27;
        _128_v0 = (_339_v121).f33;
        let _351_v123;
        _351_v123 = _dafny.Map.Empty.slice().updateUnsafe(_134_v7,(_128_v0) === ((_339_v121).f33));
        (_146_globalState).f15 = (((_351_v123).contains(_134_v7)) ? ((_351_v123).get(_134_v7)) : ((_333_v116).f31));
        let _352_v124;
        _352_v124 = _dafny.Map.Empty.slice().updateUnsafe(_203_v49,_301_v108);
        let _353_v125;
        _353_v125 = _module.D6.create_DC21(_129_v1);
        let _354_v126;
        _354_v126 = _dafny.MultiSet.fromElements(_353_v125);
        let _355_v127;
        _355_v127 = _dafny.Seq.of(_module.D6.create_DC21(_dafny.Set.fromElements(_336_v119)), _353_v125);
        let _356_v128;
        _356_v128 = _dafny.Seq.of(_354_v126, _dafny.MultiSet.FromArray(_355_v127), _dafny.MultiSet.fromElements(_353_v125, _353_v125), _354_v126, _354_v126);
        let _rhs49 = _352_v124;
        let _rhs50 = _module.__default.safeModuloInt((_339_v121).f34, new BigNumber((_134_v7).length));
        let _rhs51 = (_354_v126).equals((_356_v128)[_module.__default.safeIndex(new BigNumber(499), new BigNumber((_356_v128).length))]);
        let _lhs45 = _146_globalState;
        let _lhs46 = _146_globalState;
        _352_v124 = _rhs49;
        _lhs45.f2 = _rhs50;
        _lhs46.f15 = _rhs51;
        let _rhs52 = _131_v3;
        let _rhs53 = new BigNumber(734);
        let _rhs54 = (new BigNumber(-143)).isLessThanOrEqualTo((_339_v121).f34);
        let _rhs55 = _132_v4;
        let _lhs47 = _146_globalState;
        let _lhs48 = _146_globalState;
        _131_v3 = _rhs52;
        _lhs47.f2 = _rhs53;
        _lhs48.f28 = _rhs54;
        _132_v4 = _rhs55;
      } else {
        let _357___mcc_h32 = (_source7).cf93;
        let _358_cf93 = _357___mcc_h32;
        let _359_v129;
        _359_v129 = _module.D0.create_DC0((_299_v106).f29);
        let _rhs56 = _359_v129;
        let _rhs57 = _203_v49;
        _359_v129 = _rhs56;
        _203_v49 = _rhs57;
        let _360_v130;
        let _nw59 = Array((new BigNumber(4)).toNumber()).fill([]);
        _360_v130 = _nw59;
        let _index46 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_360_v130).length));
        (_360_v130)[_index46] = _203_v49;
        let _index47 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_360_v130).length));
        (_360_v130)[_index47] = _203_v49;
        let _361_v131;
        let _init5 = ((_362_v7) => function (_363_i15) {
          return _362_v7;
        })(_134_v7);
        let _nw60 = Array((new BigNumber(3)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw60.length); _i0_5++) {
          _nw60[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _361_v131 = _nw60;
        let _index48 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_361_v131).length));
        (_361_v131)[_index48] = _134_v7;
        let _index49 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_361_v131).length));
        (_361_v131)[_index49] = _134_v7;
        _304_v111 = (_304_v111).update((_333_v116).f32, (_339_v121).f34);
      }
      let _364_i16;
      _364_i16 = _dafny.ZERO;
      L1: {
        while (!((_333_v116).f31) || ((_203_v49)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_203_v49).length))])) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_364_i16)) {
              break L1;
            }
            _364_i16 = (_364_i16).plus(_dafny.ONE);
            let _365_v132;
            _365_v132 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_301_v108).length),_301_v108);
            let _366_v133;
            _366_v133 = _dafny.Map.Empty.slice().updateUnsafe(_133_v6,(_365_v132).Merge(_365_v132));
            _366_v133 = (_366_v133).update((_133_v6).update((_339_v121).f34, _module.__default.abs(_131_v3)), _365_v132);
            let _367_v134;
            _367_v134 = _module.D0.create_DC1(_203_v49, (_339_v121).f34, new BigNumber((_299_v106.f30).length));
            _203_v49 = ((((_299_v106).f29) ? (_367_v134) : (_367_v134))).dtor_cf1;
            _131_v3 = (_dafny.ZERO).minus(((_339_v121).f34).multipliedBy(new BigNumber((_129_v1).length)));
            (_146_globalState).f19 = _131_v3;
          }
        }
      }
      (_146_globalState).f2 = (_dafny.ZERO).minus((_339_v121).f34);
      let _368_v135;
      _368_v135 = _dafny.MultiSet.fromElements(_140_v8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-956)), ((_369_v12) => function (_370_i17) {
        return _369_v12;
      })(_144_v12)));
      let _371_v136;
      let _372_v137;
      let _out8;
      let _out9;
      let _outcollector4 = (_339_v121).m6(_module.__default.fm10((_339_v121).f34, new BigNumber((_368_v135).cardinality()), _146_globalState), _335_v118, _140_v8, _module.D0.create_DC0(!(!(_128_v0))), _146_globalState);
      _out8 = _outcollector4[0];
      _out9 = _outcollector4[1];
      _371_v136 = _out8;
      _372_v137 = _out9;
      process.stdout.write(_dafny.toString(_128_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_129_v1).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_130_v2, _dafny.Seq.of(_dafny.Set.fromElements(true), _dafny.Set.fromElements(true), _dafny.Set.fromElements(true), _dafny.Set.fromElements(true), _dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_131_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_132_v4).equals(_dafny.Set.fromElements(new BigNumber(-208)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v6).equals(_dafny.MultiSet.fromElements(_dafny.ZERO, new BigNumber(-1), new BigNumber(10), new BigNumber(9), new BigNumber(-916), new BigNumber(-208)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_134_v7, _dafny.Seq.of(_dafny.ONE, new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_140_v8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_141_v9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_141_v9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_141_v9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_142_v10).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-208),new _dafny.CodePoint('r'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_144_v12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_v13).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_146_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_146_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_146_globalState.f3, _dafny.Seq.of(_dafny.Set.fromElements(true), _dafny.Set.fromElements(true), _dafny.Set.fromElements(true), _dafny.Set.fromElements(true), _dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_146_globalState).f5).equals(_dafny.MultiSet.fromElements(new BigNumber(-208)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_146_globalState.f7, _dafny.Seq.of(_dafny.ONE, new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_146_globalState.f11).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_146_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_146_globalState.f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_146_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_146_globalState.f16).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState.f17)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState.f17)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState.f17)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_146_globalState.f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_146_globalState.f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_146_globalState.f20).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_146_globalState.f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_146_globalState).f22).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_146_globalState).f25).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState).f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_globalState).f27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_146_globalState.f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_201_v47).dtor_cf12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_201_v47).dtor_cf13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_202_v48, _dafny.Seq.of(_module.D2.create_DC8(new _dafny.CodePoint('l'.codePointAt(0)), true), _module.D2.create_DC8(new _dafny.CodePoint('l'.codePointAt(0)), true), _module.D2.create_DC8(new _dafny.CodePoint('l'.codePointAt(0)), true), _module.D2.create_DC8(new _dafny.CodePoint('l'.codePointAt(0)), true), _module.D2.create_DC8(new _dafny.CodePoint('l'.codePointAt(0)), true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v49)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_205_v50).dtor_cf41, _dafny.Seq.of(true, true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v50).dtor_cf42));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v50).dtor_cf43));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v50).dtor_cf44));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v50).dtor_cf45));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_298_v105)[new BigNumber(11)]).f29));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_298_v105)[new BigNumber(11)].f30).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_299_v106).f29));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_299_v106.f30).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_300_v107).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_301_v108, _dafny.Seq.of(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v109).equals(_dafny.Map.Empty.slice().updateUnsafe(false,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_304_v111).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-208)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v114).dtor_cf88));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v114).dtor_cf89));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v114).dtor_cf90));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_306_v114).dtor_cf91).equals(_dafny.MultiSet.fromElements(_dafny.ZERO, new BigNumber(-1), new BigNumber(10), new BigNumber(9), new BigNumber(-916), new BigNumber(-208)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(14)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(15)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(16)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(17)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(18)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(19)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(20)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(21)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(22)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(23)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(24)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(25)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v115)[new BigNumber(26)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_333_v116).f31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_333_v116).f32));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_334_v117).dtor_cf111).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_335_v118).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_336_v119));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_337_i14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_339_v121).f33));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_339_v121).f34));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_340_v122).dtor_cf94));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_340_v122).dtor_cf95));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_340_v122).dtor_cf96));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_340_v122).dtor_cf97)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_340_v122).dtor_cf97)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_340_v122).dtor_cf97)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_340_v122).dtor_cf98));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_364_i16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_368_v135).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("rqigbg"), _dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_371_v136));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_372_v137));
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
    static create_DC1(cf1, cf2, cf3) {
      let $dt = new D0(1);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC2(cf4, cf5, cf6) {
      let $dt = new D0(2);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ", " + this.cf5.toVerbatimString(true) + ", " + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf0 === other.cf0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(false);
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
    static create_DC4(cf8) {
      let $dt = new D1(0);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC3(cf7) {
      let $dt = new D1(1);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC5(cf9) {
      let $dt = new D1(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf8 === other.cf8;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(false);
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
    static create_DC7(cf11) {
      let $dt = new D2(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC8(cf12, cf13) {
      let $dt = new D2(1);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC9(cf14, cf15, cf16, cf17) {
      let $dt = new D2(2);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC6(cf10) {
      let $dt = new D2(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC9" + "(" + this.cf14.toVerbatimString(true) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
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
        return other.$tag === 0 && this.cf11 === other.cf11;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12) && this.cf13 === other.cf13;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(false);
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
    static create_DC11(cf19, cf20) {
      let $dt = new D3(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC10(cf18) {
      let $dt = new D3(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC12(cf21) {
      let $dt = new D3(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf19 === other.cf19 && this.cf20 === other.cf20;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC11(false, false);
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
    static create_DC14(cf23) {
      let $dt = new D4(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC15(cf24, cf25) {
      let $dt = new D4(1);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC13(cf22) {
      let $dt = new D4(2);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC14(_dafny.ZERO);
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
    static create_DC17(cf27) {
      let $dt = new D5(0);
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC18() {
      let $dt = new D5(1);
      return $dt;
    }
    static create_DC19(cf28, cf29, cf30) {
      let $dt = new D5(2);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC16(cf26) {
      let $dt = new D5(3);
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC20(cf31) {
      let $dt = new D5(4);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get is_DC16() { return this.$tag === 3; }
    get is_DC20() { return this.$tag === 4; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC18";
      } else if (this.$tag === 2) {
        return "D5.DC19" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 4) {
        return "D5.DC20" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28) && _dafny.areEqual(this.cf29, other.cf29) && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf26, other.cf26);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC17(_dafny.Set.Empty);
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
    static create_DC22(cf33, cf34, cf35) {
      let $dt = new D6(0);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC23(cf36, cf37) {
      let $dt = new D6(1);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC21(cf32) {
      let $dt = new D6(2);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC22" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC23" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC21" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34) && this.cf35 === other.cf35;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf32, other.cf32);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC22(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC24(cf38) {
      let $dt = new D7(0);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC24" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38);
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
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC26() {
      let $dt = new D8(0);
      return $dt;
    }
    static create_DC25(cf39) {
      let $dt = new D8(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC26";
      } else if (this.$tag === 1) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf39) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC26();
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
    static create_DC28(cf41, cf42, cf43, cf44, cf45) {
      let $dt = new D9(0);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC29(cf46, cf47, cf48) {
      let $dt = new D9(1);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC27(cf40) {
      let $dt = new D9(2);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC28" + "(" + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC29" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf41, other.cf41) && this.cf42 === other.cf42 && _dafny.areEqual(this.cf43, other.cf43) && this.cf44 === other.cf44 && this.cf45 === other.cf45;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47) && this.cf48 === other.cf48;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC28(_dafny.Seq.of(), false, _dafny.ZERO, false, false);
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
    static create_DC31(cf50, cf51) {
      let $dt = new D10(0);
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC32(cf52, cf53) {
      let $dt = new D10(1);
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC33(cf54, cf55) {
      let $dt = new D10(2);
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC30(cf49) {
      let $dt = new D10(3);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC33() { return this.$tag === 2; }
    get is_DC30() { return this.$tag === 3; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC31" + "(" + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC32" + "(" + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC33" + "(" + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 3) {
        return "D10.DC30" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf50 === other.cf50 && this.cf51 === other.cf51;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf52, other.cf52) && this.cf53 === other.cf53;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf54 === other.cf54 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf49 === other.cf49;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC31(false, false);
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
    static create_DC35(cf57, cf58) {
      let $dt = new D11(0);
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC36(cf59, cf60, cf61) {
      let $dt = new D11(1);
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC37(cf62, cf63) {
      let $dt = new D11(2);
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC34(cf56) {
      let $dt = new D11(3);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get is_DC34() { return this.$tag === 3; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC35" + "(" + this.cf57.toVerbatimString(true) + ", " + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC36" + "(" + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC37" + "(" + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC34" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf57, other.cf57) && this.cf58 === other.cf58;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf59, other.cf59) && _dafny.areEqual(this.cf60, other.cf60) && this.cf61 === other.cf61;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf62, other.cf62) && this.cf63 === other.cf63;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf56, other.cf56);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC35(_dafny.Seq.UnicodeFromString(""), false);
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
    static create_DC39(cf65, cf66, cf67) {
      let $dt = new D12(0);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC38(cf64) {
      let $dt = new D12(1);
      $dt.cf64 = cf64;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf64() { return this.cf64; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC39" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC38" + "(" + _dafny.toString(this.cf64) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf65, other.cf65) && this.cf66 === other.cf66 && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf64 === other.cf64;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC39(_dafny.ZERO, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC41() {
      let $dt = new D13(0);
      return $dt;
    }
    static create_DC40(cf68) {
      let $dt = new D13(1);
      $dt.cf68 = cf68;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get dtor_cf68() { return this.cf68; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC41";
      } else if (this.$tag === 1) {
        return "D13.DC40" + "(" + _dafny.toString(this.cf68) + ")";
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
        return other.$tag === 1 && this.cf68 === other.cf68;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC41();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC42(cf69) {
      let $dt = new D14(0);
      $dt.cf69 = cf69;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get dtor_cf69() { return this.cf69; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC42" + "(" + _dafny.toString(this.cf69) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf69, other.cf69);
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
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC44(cf71, cf72, cf73, cf74, cf75) {
      let $dt = new D15(0);
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC45(cf76, cf77, cf78, cf79, cf80) {
      let $dt = new D15(1);
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      return $dt;
    }
    static create_DC46() {
      let $dt = new D15(2);
      return $dt;
    }
    static create_DC43(cf70) {
      let $dt = new D15(3);
      $dt.cf70 = cf70;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get is_DC45() { return this.$tag === 1; }
    get is_DC46() { return this.$tag === 2; }
    get is_DC43() { return this.$tag === 3; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf70() { return this.cf70; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC44" + "(" + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC45" + "(" + _dafny.toString(this.cf76) + ", " + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ", " + _dafny.toString(this.cf80) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC46";
      } else if (this.$tag === 3) {
        return "D15.DC43" + "(" + _dafny.toString(this.cf70) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf71, other.cf71) && this.cf72 === other.cf72 && this.cf73 === other.cf73 && this.cf74 === other.cf74 && _dafny.areEqual(this.cf75, other.cf75);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf76, other.cf76) && _dafny.areEqual(this.cf77, other.cf77) && this.cf78 === other.cf78 && this.cf79 === other.cf79 && this.cf80 === other.cf80;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf70, other.cf70);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC44(_dafny.ZERO, false, [], [], _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC48() {
      let $dt = new D16(0);
      return $dt;
    }
    static create_DC47(cf81) {
      let $dt = new D16(1);
      $dt.cf81 = cf81;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get is_DC47() { return this.$tag === 1; }
    get dtor_cf81() { return this.cf81; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC48";
      } else if (this.$tag === 1) {
        return "D16.DC47" + "(" + _dafny.toString(this.cf81) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf81, other.cf81);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC48();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC50(cf83, cf84, cf85) {
      let $dt = new D17(0);
      $dt.cf83 = cf83;
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC51(cf86, cf87) {
      let $dt = new D17(1);
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      return $dt;
    }
    static create_DC52(cf88, cf89, cf90, cf91) {
      let $dt = new D17(2);
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      return $dt;
    }
    static create_DC49(cf82) {
      let $dt = new D17(3);
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC53(cf92) {
      let $dt = new D17(4);
      $dt.cf92 = cf92;
      return $dt;
    }
    get is_DC50() { return this.$tag === 0; }
    get is_DC51() { return this.$tag === 1; }
    get is_DC52() { return this.$tag === 2; }
    get is_DC49() { return this.$tag === 3; }
    get is_DC53() { return this.$tag === 4; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf92() { return this.cf92; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC50" + "(" + _dafny.toString(this.cf83) + ", " + _dafny.toString(this.cf84) + ", " + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC51" + "(" + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC52" + "(" + _dafny.toString(this.cf88) + ", " + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ")";
      } else if (this.$tag === 3) {
        return "D17.DC49" + "(" + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 4) {
        return "D17.DC53" + "(" + _dafny.toString(this.cf92) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf83 === other.cf83 && this.cf84 === other.cf84 && this.cf85 === other.cf85;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf86, other.cf86) && this.cf87 === other.cf87;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf88 === other.cf88 && this.cf89 === other.cf89 && _dafny.areEqual(this.cf90, other.cf90) && _dafny.areEqual(this.cf91, other.cf91);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf82, other.cf82);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf92, other.cf92);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC50(false, false, null);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC55(cf94, cf95, cf96, cf97, cf98) {
      let $dt = new D18(0);
      $dt.cf94 = cf94;
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      $dt.cf97 = cf97;
      $dt.cf98 = cf98;
      return $dt;
    }
    static create_DC54(cf93) {
      let $dt = new D18(1);
      $dt.cf93 = cf93;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get is_DC54() { return this.$tag === 1; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf93() { return this.cf93; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC55" + "(" + _dafny.toString(this.cf94) + ", " + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ", " + _dafny.toString(this.cf97) + ", " + _dafny.toString(this.cf98) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC54" + "(" + _dafny.toString(this.cf93) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf94, other.cf94) && this.cf95 === other.cf95 && _dafny.areEqual(this.cf96, other.cf96) && this.cf97 === other.cf97 && this.cf98 === other.cf98;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf93, other.cf93);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC55(new _dafny.CodePoint('D'.codePointAt(0)), false, _dafny.ZERO, [], false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC57(cf100, cf101, cf102, cf103, cf104) {
      let $dt = new D19(0);
      $dt.cf100 = cf100;
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      $dt.cf104 = cf104;
      return $dt;
    }
    static create_DC58(cf105, cf106, cf107) {
      let $dt = new D19(1);
      $dt.cf105 = cf105;
      $dt.cf106 = cf106;
      $dt.cf107 = cf107;
      return $dt;
    }
    static create_DC59(cf108, cf109) {
      let $dt = new D19(2);
      $dt.cf108 = cf108;
      $dt.cf109 = cf109;
      return $dt;
    }
    static create_DC56(cf99) {
      let $dt = new D19(3);
      $dt.cf99 = cf99;
      return $dt;
    }
    static create_DC60(cf110) {
      let $dt = new D19(4);
      $dt.cf110 = cf110;
      return $dt;
    }
    get is_DC57() { return this.$tag === 0; }
    get is_DC58() { return this.$tag === 1; }
    get is_DC59() { return this.$tag === 2; }
    get is_DC56() { return this.$tag === 3; }
    get is_DC60() { return this.$tag === 4; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf108() { return this.cf108; }
    get dtor_cf109() { return this.cf109; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf110() { return this.cf110; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC57" + "(" + _dafny.toString(this.cf100) + ", " + _dafny.toString(this.cf101) + ", " + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ", " + _dafny.toString(this.cf104) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC58" + "(" + _dafny.toString(this.cf105) + ", " + _dafny.toString(this.cf106) + ", " + _dafny.toString(this.cf107) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC59" + "(" + _dafny.toString(this.cf108) + ", " + _dafny.toString(this.cf109) + ")";
      } else if (this.$tag === 3) {
        return "D19.DC56" + "(" + _dafny.toString(this.cf99) + ")";
      } else if (this.$tag === 4) {
        return "D19.DC60" + "(" + _dafny.toString(this.cf110) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf100 === other.cf100 && _dafny.areEqual(this.cf101, other.cf101) && this.cf102 === other.cf102 && this.cf103 === other.cf103 && _dafny.areEqual(this.cf104, other.cf104);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf105 === other.cf105 && this.cf106 === other.cf106 && this.cf107 === other.cf107;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf108 === other.cf108 && _dafny.areEqual(this.cf109, other.cf109);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf99, other.cf99);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf110, other.cf110);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC57([], _dafny.ZERO, null, null, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC62() {
      let $dt = new D20(0);
      return $dt;
    }
    static create_DC63(cf112) {
      let $dt = new D20(1);
      $dt.cf112 = cf112;
      return $dt;
    }
    static create_DC61(cf111) {
      let $dt = new D20(2);
      $dt.cf111 = cf111;
      return $dt;
    }
    static create_DC64(cf113) {
      let $dt = new D20(3);
      $dt.cf113 = cf113;
      return $dt;
    }
    get is_DC62() { return this.$tag === 0; }
    get is_DC63() { return this.$tag === 1; }
    get is_DC61() { return this.$tag === 2; }
    get is_DC64() { return this.$tag === 3; }
    get dtor_cf112() { return this.cf112; }
    get dtor_cf111() { return this.cf111; }
    get dtor_cf113() { return this.cf113; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC62";
      } else if (this.$tag === 1) {
        return "D20.DC63" + "(" + _dafny.toString(this.cf112) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC61" + "(" + _dafny.toString(this.cf111) + ")";
      } else if (this.$tag === 3) {
        return "D20.DC64" + "(" + _dafny.toString(this.cf113) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf112, other.cf112);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf111, other.cf111);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf113, other.cf113);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC62();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC66(cf115) {
      let $dt = new D21(0);
      $dt.cf115 = cf115;
      return $dt;
    }
    static create_DC67(cf116) {
      let $dt = new D21(1);
      $dt.cf116 = cf116;
      return $dt;
    }
    static create_DC65(cf114) {
      let $dt = new D21(2);
      $dt.cf114 = cf114;
      return $dt;
    }
    get is_DC66() { return this.$tag === 0; }
    get is_DC67() { return this.$tag === 1; }
    get is_DC65() { return this.$tag === 2; }
    get dtor_cf115() { return this.cf115; }
    get dtor_cf116() { return this.cf116; }
    get dtor_cf114() { return this.cf114; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC66" + "(" + _dafny.toString(this.cf115) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC67" + "(" + _dafny.toString(this.cf116) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC65" + "(" + _dafny.toString(this.cf114) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf115, other.cf115);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf116, other.cf116);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf114, other.cf114);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC66(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D21.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = _dafny.ZERO;
      this.f2 = _dafny.ZERO;
      this.f3 = _dafny.Seq.of();
      this.f7 = _dafny.Seq.of();
      this.f11 = _dafny.Seq.UnicodeFromString("");
      this.f13 = false;
      this.f14 = _dafny.ZERO;
      this.f15 = false;
      this.f16 = _dafny.Map.Empty;
      this.f17 = [];
      this.f18 = false;
      this.f19 = _dafny.ZERO;
      this.f20 = _dafny.Seq.UnicodeFromString("");
      this.f21 = _dafny.ZERO;
      this.f28 = false;
      this._f0 = false;
      this._f4 = false;
      this._f5 = _dafny.MultiSet.Empty;
      this._f6 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
      this._f9 = _dafny.ZERO;
      this._f10 = _dafny.ZERO;
      this._f12 = _dafny.ZERO;
      this._f22 = _dafny.Map.Empty;
      this._f23 = _dafny.ZERO;
      this._f24 = false;
      this._f25 = _dafny.Seq.UnicodeFromString("");
      this._f26 = _dafny.ZERO;
      this._f27 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27, f28) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this).f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this).f14 = f14;
      (_this).f15 = f15;
      (_this).f16 = f16;
      (_this).f17 = f17;
      (_this).f18 = f18;
      (_this).f19 = f19;
      (_this).f20 = f20;
      (_this).f21 = f21;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      (_this)._f27 = f27;
      (_this).f28 = f28;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
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
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f37 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f37) {
      let _this = this;
      (_this).f37 = f37;
      return;
    }
    fm11(p0, globalState) {
      let _this = this;
      return _this.f37;
    };
    fm12(globalState) {
      let _this = this;
      return (new BigNumber(-345)).minus(((_this.f37) ? (new BigNumber(696)) : (new BigNumber(343))));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f40 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f40) {
      let _this = this;
      (_this)._f40 = f40;
      return;
    }
    fm30(p0, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ovq")).length), (_dafny.ZERO).minus(new BigNumber(-114)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f40,new BigNumber(-638))).length))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true, !((_this).f40))).length), new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)))).length), new BigNumber(141)));
    };
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _373_v0;
      _373_v0 = _dafny.Set.fromElements(p0);
      let _374_v1;
      _374_v1 = _module.D6.create_DC21(_373_v0);
      let _375_v2;
      _375_v2 = _dafny.Seq.of(_374_v1);
      let _376_v3;
      _376_v3 = _module.D8.create_DC25(_375_v2);
      let _377_v4;
      _377_v4 = new BigNumber(651);
      let _378_v5;
      _378_v5 = _dafny.MultiSet.fromElements((_376_v3).dtor_cf39, _module.__default.fm31(_377_v4, new BigNumber(241), globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(660)), ((_379_v1) => function (_380_i0) {
        return _379_v1;
      })(_374_v1)), _375_v2);
      let _381_v6;
      _381_v6 = _dafny.Map.Empty.slice().updateUnsafe(_377_v4,_dafny.Set.fromElements(p0, (_this).f40));
      (globalState).f18 = !((_378_v5).IsDisjointFrom(_dafny.MultiSet.fromElements(_dafny.Seq.update(_375_v2, _module.__default.safeIndex(new BigNumber(140), new BigNumber((_375_v2).length)), _module.D6.create_DC21((((_381_v6).contains(_377_v4)) ? ((_381_v6).get(_377_v4)) : (_dafny.Set.fromElements((_this).f40))))))));
      let _382_v7;
      _382_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f40,p0);
      let _383_v8;
      _383_v8 = _dafny.Seq.of(_377_v4);
      let _384_v9;
      _384_v9 = _dafny.Seq.of((_382_v7).Merge(_382_v7), _module.__default.fm32(p0, _module.__default.fm33(true, new BigNumber(-164), _377_v4, globalState), _383_v8, globalState), _382_v7);
      let _385_v10;
      _385_v10 = _dafny.Seq.of(_384_v9, _dafny.Seq.Concat(_384_v9, _384_v9));
      _384_v9 = (_385_v10)[_module.__default.safeIndex((_377_v4).plus(_377_v4), new BigNumber((_385_v10).length))];
      let _386_v11;
      _386_v11 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)));
      let _387_v12;
      _387_v12 = _dafny.Seq.UnicodeFromString("xo");
      let _388_v13;
      let _nw61 = Array((new BigNumber(17)).toNumber());
      _nw61[(_dafny.ZERO).toNumber()] = (_386_v11).IsDisjointFrom(_386_v11);
      _nw61[(_dafny.ONE).toNumber()] = !((_this).f40);
      _nw61[(new BigNumber(2)).toNumber()] = p0;
      _nw61[(new BigNumber(3)).toNumber()] = p0;
      _nw61[(new BigNumber(4)).toNumber()] = (_this).f40;
      _nw61[(new BigNumber(5)).toNumber()] = (((_this).f40) ? ((_this).f40) : ((_this).f40));
      _nw61[(new BigNumber(6)).toNumber()] = p0;
      _nw61[(new BigNumber(7)).toNumber()] = p0;
      _nw61[(new BigNumber(8)).toNumber()] = (_this).f40;
      _nw61[(new BigNumber(9)).toNumber()] = (_this).f40;
      _nw61[(new BigNumber(10)).toNumber()] = _module.__default.fm33(p0, new BigNumber(364), new BigNumber(-247), globalState);
      _nw61[(new BigNumber(11)).toNumber()] = p0;
      _nw61[(new BigNumber(12)).toNumber()] = (_this).f40;
      _nw61[(new BigNumber(13)).toNumber()] = (new BigNumber((_dafny.Seq.of(new BigNumber(286))).length)).isLessThan(new BigNumber((_387_v12).length));
      _nw61[(new BigNumber(14)).toNumber()] = (_this).f40;
      _nw61[(new BigNumber(15)).toNumber()] = false;
      _nw61[(new BigNumber(16)).toNumber()] = (p0) || (p0);
      _388_v13 = _nw61;
      let _index50 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_388_v13).length));
      (_388_v13)[_index50] = (_this).f40;
      let _index51 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_388_v13).length));
      (_388_v13)[_index51] = _dafny.Seq.IsPrefixOf(_387_v12, _387_v12);
      let _389_v14;
      let _nw62 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _389_v14 = _nw62;
      let _index52 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_389_v14).length));
      (_389_v14)[_index52] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-162)), function (_390_i1) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      });
      let _index53 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_389_v14).length));
      (_389_v14)[_index53] = _module.__default.fm34(_377_v4, _377_v4, p0, p0, globalState);
      let _391_v15;
      _391_v15 = _dafny.Seq.of((_388_v13)[_module.__default.safeIndex(new BigNumber(950), new BigNumber((_388_v13).length))], ((_this).f40) || (!((_this).f40)), ((_388_v13)[_module.__default.safeIndex(new BigNumber(950), new BigNumber((_388_v13).length))]) && ((_this).f40), true, (_this).f40);
      let _392_v16;
      _392_v16 = _module.D4.create_DC15(_377_v4, _377_v4);
      let _393_v17;
      _393_v17 = new _dafny.CodePoint('i'.codePointAt(0));
      let _394_v18;
      _394_v18 = _module.D2.create_DC8(_393_v17, (_388_v13)[_module.__default.safeIndex(new BigNumber(950), new BigNumber((_388_v13).length))]);
      let _395_v19;
      _395_v19 = _dafny.Seq.of(_module.D2.create_DC8(_393_v17, p0), _394_v18, _394_v18);
      let _rhs58 = (((_382_v7).contains((_388_v13)[_module.__default.safeIndex(new BigNumber(950), new BigNumber((_388_v13).length))])) ? ((_382_v7).get((_388_v13)[_module.__default.safeIndex(new BigNumber(950), new BigNumber((_388_v13).length))])) : (true));
      let _rhs59 = _377_v4;
      let _rhs60 = _module.__default.fm35(_dafny.Seq.Concat(_395_v19, _395_v19), _377_v4, globalState);
      let _rhs61 = _392_v16;
      let _lhs49 = globalState;
      let _lhs50 = globalState;
      _lhs49.f28 = _rhs58;
      _lhs50.f19 = _rhs59;
      _391_v15 = _rhs60;
      _392_v16 = _rhs61;
      let _396_v20;
      let _nw63 = new _module.C0();
      _nw63.__ctor(_dafny.areEqual(_dafny.Seq.UnicodeFromString("rrjunsmd"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(106)), ((_397_v17) => function (_398_i2) {
        return _397_v17;
      })(_393_v17))));
      _396_v20 = _nw63;
      let _399_v21;
      _399_v21 = _dafny.Seq.of(_391_v15);
      r0 = _399_v21;
      r1 = p0;
      return [r0, r1];
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let _400_v0;
      _400_v0 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p2));
      let _401_v1;
      _401_v1 = _dafny.Seq.of(p2, p2, p2);
      let _402_v2;
      _402_v2 = _dafny.Map.Empty.slice().updateUnsafe((_400_v0).Intersect(_dafny.MultiSet.fromElements(p2)),_dafny.Seq.Concat(_401_v1, _401_v1));
      let _403_v3;
      let _init6 = ((_404_p0) => function (_405_i0) {
        return _404_p0;
      })(p0);
      let _nw64 = Array((new BigNumber(26)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw64.length); _i0_6++) {
        _nw64[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _403_v3 = _nw64;
      let _index54 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length));
      (_403_v3)[_index54] = !((_this).f40);
      let _406_v4;
      _406_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f40,p0);
      let _407_v5;
      _407_v5 = _module.D9.create_DC27(_dafny.Seq.Create(_module.__default.abs(new BigNumber(368)), function (_408_i1) {
  return new BigNumber(102);
}));
      let _index55 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length));
      let _rhs62 = (_this).f40;
      let _rhs63 = p0;
      let _rhs64 = ((_402_v2).update(_dafny.MultiSet.fromElements(new BigNumber((_406_v4).length), p2, p2, p2, new BigNumber(478)), (_407_v5).dtor_cf40)).Merge(_402_v2);
      let _rhs65 = !(!(p0)) || (p0);
      let _lhs51 = globalState;
      let _lhs52 = globalState;
      let _lhs53 = _403_v3;
      let _lhs54 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length));
      _lhs51.f18 = _rhs62;
      _lhs52.f18 = _rhs63;
      _402_v2 = _rhs64;
      _lhs53[_lhs54] = _rhs65;
      if (!((_403_v3)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length))]) || (!((_this).f40))) {
        let _409_v6;
        _409_v6 = _module.D4.create_DC14((_dafny.ZERO).minus(p2));
        let _pat_let_tv9 = p2;
        _409_v6 = function (_pat_let4_0) {
          return function (_410_dt__update__tmp_h0) {
            return function (_pat_let5_0) {
              return function (_411_dt__update_hcf23_h0) {
                return _module.D4.create_DC14(_411_dt__update_hcf23_h0);
              }(_pat_let5_0);
            }(_pat_let_tv9);
          }(_pat_let4_0);
        }(_409_v6);
        let _412_v7;
        let _nw65 = new _module.C0();
        _nw65.__ctor((_this).f40);
        _412_v7 = _nw65;
        let _index56 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length));
        (_403_v3)[_index56] = !(_412_v7.f37);
        let _413_v8;
        _413_v8 = _module.D0.create_DC1(_403_v3, p2, p2);
        let _414_v9;
        let _nw66 = Array((new BigNumber(9)).toNumber());
        _nw66[(_dafny.ZERO).toNumber()] = (_413_v8).dtor_cf1;
        _nw66[(_dafny.ONE).toNumber()] = _403_v3;
        _nw66[(new BigNumber(2)).toNumber()] = (_413_v8).dtor_cf1;
        _nw66[(new BigNumber(3)).toNumber()] = _403_v3;
        _nw66[(new BigNumber(4)).toNumber()] = _403_v3;
        _nw66[(new BigNumber(5)).toNumber()] = ((_412_v7.f37) ? (_403_v3) : (_403_v3));
        _nw66[(new BigNumber(6)).toNumber()] = _403_v3;
        _nw66[(new BigNumber(7)).toNumber()] = _403_v3;
        _nw66[(new BigNumber(8)).toNumber()] = _403_v3;
        _414_v9 = _nw66;
        let _index57 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_414_v9).length));
        (_414_v9)[_index57] = _403_v3;
        let _415_v10;
        _415_v10 = _dafny.Seq.UnicodeFromString("eakffkqhw");
        let _416_v11;
        _416_v11 = _dafny.Seq.of(_415_v10, _dafny.Seq.UnicodeFromString("dkbchhlud"), _dafny.Seq.Concat(_415_v10, _415_v10));
        let _index58 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_414_v9).length));
        let _rhs66 = (((_this).f40) ? (_403_v3) : (_403_v3));
        let _rhs67 = _415_v10;
        let _rhs68 = new BigNumber((_416_v11).length);
        let _rhs69 = _412_v7.f37;
        let _rhs70 = _module.__default.safeModuloInt((new BigNumber(101)).multipliedBy(p2), _module.__default.safeDivisionInt((_412_v7).fm12(globalState), p2));
        let _lhs55 = _414_v9;
        let _lhs56 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_414_v9).length));
        let _lhs57 = globalState;
        let _lhs58 = globalState;
        let _lhs59 = globalState;
        let _lhs60 = globalState;
        _lhs55[_lhs56] = _rhs66;
        _lhs57.f20 = _rhs67;
        _lhs58.f19 = _rhs68;
        _lhs59.f28 = _rhs69;
        _lhs60.f21 = _rhs70;
        let _417_v12;
        _417_v12 = _dafny.Map.Empty.slice().updateUnsafe(_401_v1,p0);
        let _418_v14;
        _418_v14 = _dafny.MultiSet.fromElements(_401_v1, _dafny.Seq.of(p2));
        let _419_v15;
        let _nw67 = new _module.C0();
        _nw67.__ctor((_417_v12).equals(function () {
          let _coll25 = new _dafny.Map();
          for (const _compr_25 of (_418_v14).Elements) {
            let _420_v13 = _compr_25;
            if ((_418_v14).contains(_420_v13)) {
              _coll25.push([_420_v13,!(_412_v7.f37)]);
            }
          }
          return _coll25;
        }()));
        _419_v15 = _nw67;
      } else {
        let _421_v16;
        _421_v16 = _dafny.Set.fromElements(p2);
        let _422_v17;
        _422_v17 = new _dafny.CodePoint('q'.codePointAt(0));
        let _423_v18;
        _423_v18 = _dafny.Seq.of(_422_v17, _422_v17, new _dafny.CodePoint('j'.codePointAt(0)), _422_v17);
        let _nw68 = Array((new BigNumber(21)).toNumber());
        _nw68[(_dafny.ZERO).toNumber()] = p2;
        _nw68[(_dafny.ONE).toNumber()] = (_module.__default.fm36(globalState)).plus(_module.__default.fm36(globalState));
        _nw68[(new BigNumber(2)).toNumber()] = new BigNumber((_421_v16).length);
        _nw68[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(p2, p2);
        _nw68[(new BigNumber(4)).toNumber()] = p2;
        _nw68[(new BigNumber(5)).toNumber()] = p2;
        _nw68[(new BigNumber(6)).toNumber()] = (new BigNumber(136)).minus(p2);
        _nw68[(new BigNumber(7)).toNumber()] = p2;
        _nw68[(new BigNumber(8)).toNumber()] = p2;
        _nw68[(new BigNumber(9)).toNumber()] = p2;
        _nw68[(new BigNumber(10)).toNumber()] = new BigNumber(-137);
        _nw68[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(p2, new BigNumber(389));
        _nw68[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt(p2, p2);
        _nw68[(new BigNumber(13)).toNumber()] = p2;
        _nw68[(new BigNumber(14)).toNumber()] = p2;
        _nw68[(new BigNumber(15)).toNumber()] = p2;
        _nw68[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_423_v18, _423_v18)).length);
        _nw68[(new BigNumber(17)).toNumber()] = (_401_v1)[_module.__default.safeIndex(p2, new BigNumber((_401_v1).length))];
        _nw68[(new BigNumber(18)).toNumber()] = p2;
        _nw68[(new BigNumber(19)).toNumber()] = (((_403_v3)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length))]) ? (p2) : (new BigNumber(241)));
        _nw68[(new BigNumber(20)).toNumber()] = p2;
        (globalState).f17 = _nw68;
        let _index59 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length));
        (_403_v3)[_index59] = (p2).isLessThan(p2);
        let _424_v19;
        _424_v19 = _dafny.Set.fromElements((_this).f40, p0, (_403_v3)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length))], (_this).f40);
        let _425_v20;
        _425_v20 = _dafny.Map.Empty.slice().updateUnsafe(p2,false);
        let _426_v21;
        let _nw69 = Array((new BigNumber(9)).toNumber());
        _nw69[(_dafny.ZERO).toNumber()] = new BigNumber(((_424_v19).Difference(_424_v19)).length);
        _nw69[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(p2, (_dafny.ZERO).minus(_module.__default.fm36(globalState)));
        _nw69[(new BigNumber(2)).toNumber()] = (((_this).f40) ? (new BigNumber(619)) : (new BigNumber((_dafny.Seq.update(_423_v18, _module.__default.safeIndex(p2, new BigNumber((_423_v18).length)), _422_v17)).length)));
        _nw69[(new BigNumber(3)).toNumber()] = (_module.D4.create_DC15(p2, new BigNumber((_423_v18).length))).dtor_cf24;
        _nw69[(new BigNumber(4)).toNumber()] = p2;
        _nw69[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_425_v20).length));
        _nw69[(new BigNumber(6)).toNumber()] = new BigNumber(773);
        _nw69[(new BigNumber(7)).toNumber()] = p2;
        _nw69[(new BigNumber(8)).toNumber()] = _module.__default.fm36(globalState);
        _426_v21 = _nw69;
        let _index60 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length));
        let _rhs71 = _426_v21;
        let _rhs72 = (_module.D6.create_DC23((_this).f40, p2)).dtor_cf36;
        let _lhs61 = globalState;
        let _lhs62 = _403_v3;
        let _lhs63 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length));
        _lhs61.f17 = _rhs71;
        _lhs62[_lhs63] = _rhs72;
        if ((_this).f40) {
          let _427_v22;
          _427_v22 = _dafny.Set.fromElements(_422_v17);
          let _428_v23;
          _428_v23 = _dafny.MultiSet.fromElements(_427_v22, _dafny.Set.fromElements(_422_v17), _427_v22, _427_v22);
          let _429_v24;
          _429_v24 = _dafny.Seq.of(_427_v22);
          let _index61 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length));
          (_403_v3)[_index61] = ((_dafny.MultiSet.FromArray(_429_v24)).Intersect(_428_v23)).IsProperSubsetOf(_428_v23);
          let _430_v25;
          _430_v25 = _dafny.Seq.of(!(_424_v19).equals(_dafny.Set.fromElements((_403_v3)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length))])), ((_403_v3)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length))]) || (p0));
          (globalState).f13 = (_430_v25)[_module.__default.safeIndex(p2, new BigNumber((_430_v25).length))];
          let _431_v26;
          let _nw70 = new _module.C0();
          _nw70.__ctor((p2).isLessThan(new BigNumber(981)));
          _431_v26 = _nw70;
          let _rhs73 = new BigNumber(-164);
          let _rhs74 = (_dafny.Map.Empty.slice().updateUnsafe(p2,_425_v20)).contains(p2);
          let _rhs75 = p2;
          let _lhs64 = globalState;
          let _lhs65 = globalState;
          let _lhs66 = globalState;
          _lhs64.f19 = _rhs73;
          _lhs65.f28 = _rhs74;
          _lhs66.f1 = _rhs75;
          let _index62 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_426_v21).length));
          (_426_v21)[_index62] = new BigNumber(((((_403_v3)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length))]) ? (_430_v25) : (_430_v25))).length);
          let _index63 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_426_v21).length));
          let _rhs76 = _403_v3;
          let _rhs77 = (((p2).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_430_v25).length)))) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(74)), ((_432_v17) => function (_433_i2) {
            return _432_v17;
          })(_422_v17))) : (((_431_v26.f37) ? (_dafny.Seq.update(_423_v18, _module.__default.safeIndex(p2, new BigNumber((_423_v18).length)), _422_v17)) : (_423_v18))));
          let _rhs78 = new BigNumber(-296);
          let _rhs79 = !(p0);
          let _lhs67 = globalState;
          let _lhs68 = _426_v21;
          let _lhs69 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_426_v21).length));
          let _lhs70 = globalState;
          _403_v3 = _rhs76;
          _lhs67.f20 = _rhs77;
          _lhs68[_lhs69] = _rhs78;
          _lhs70.f28 = _rhs79;
        } else {
          let _index64 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_426_v21).length));
          (_426_v21)[_index64] = (_dafny.ZERO).minus(_module.__default.fm36(globalState));
          let _index65 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_426_v21).length));
          (_426_v21)[_index65] = p2;
          (globalState).f28 = (_this).f40;
          let _434_v27;
          _434_v27 = _dafny.Seq.of(((_dafny.ZERO).minus(new BigNumber(-877))).isLessThanOrEqualTo(_module.__default.fm36(globalState)));
          let _435_v28;
          _435_v28 = _dafny.Map.Empty.slice().updateUnsafe((_426_v21)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_426_v21).length))],_434_v27);
          _434_v27 = ((!(p0)) ? (_434_v27) : ((((_435_v28).contains(p2)) ? ((_435_v28).get(p2)) : (_434_v27))));
          _425_v20 = _dafny.Map.Empty.slice().updateUnsafe((_426_v21)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_426_v21).length))],(_dafny.MultiSet.fromElements(p0)).IsSubsetOf(_dafny.MultiSet.fromElements(p0)));
          let _436_v29;
          let _init7 = ((_437_p2) => function (_438_i3) {
            return (_438_i3).minus(_437_p2);
          })(p2);
          let _nw71 = Array((new BigNumber(23)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw71.length); _i0_7++) {
            _nw71[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _436_v29 = _nw71;
          let _rhs80 = _422_v17;
          let _rhs81 = _436_v29;
          let _rhs82 = (_426_v21)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_426_v21).length))];
          let _lhs71 = globalState;
          let _lhs72 = globalState;
          _422_v17 = _rhs80;
          _lhs71.f17 = _rhs81;
          _lhs72.f14 = _rhs82;
        }
        (globalState).f13 = p0;
      }
      let _439_v30;
      _439_v30 = _module.D3.create_DC11((_403_v3)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length))], !(p0));
      let _440_i4;
      _440_i4 = _dafny.ZERO;
      L2: {
        while ((_439_v30).dtor_cf20) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_440_i4)) {
              break L2;
            }
            _440_i4 = (_440_i4).plus(_dafny.ONE);
            let _441_v31;
            let _nw72 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
            _441_v31 = _nw72;
            let _442_v32;
            _442_v32 = _dafny.Map.Empty.slice().updateUnsafe(_441_v31,(_407_v5).dtor_cf40);
            _442_v32 = (_442_v32).update(_441_v31, _dafny.Seq.of(p2, p2));
            (globalState).f18 = (_403_v3)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length))];
            (globalState).f14 = (p2).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(706)), function (_443_i5) {
              return new _dafny.CodePoint('a'.codePointAt(0));
            })).length));
            let _444_v33;
            _444_v33 = _dafny.Seq.UnicodeFromString("evxiub");
            let _445_v34;
            _445_v34 = _dafny.MultiSet.fromElements(false);
            let _446_v35;
            _446_v35 = new _dafny.CodePoint('p'.codePointAt(0));
            let _447_v36;
            _447_v36 = _module.D5.create_DC19(new BigNumber((_444_v33).length), (_401_v1)[_module.__default.safeIndex(new BigNumber((_445_v34).cardinality()), new BigNumber((_401_v1).length))], _446_v35);
            (globalState).f13 = !(p2).isEqualTo((_447_v36).dtor_cf28);
          }
        }
      }
      let _448_v37;
      let _nw73 = new _module.C0();
      _nw73.__ctor((_403_v3)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length))]);
      _448_v37 = _nw73;
      let _449_v38;
      _449_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      let _450_v39;
      _450_v39 = _dafny.Map.Empty.slice().updateUnsafe(_449_v38,(_module.__default.fm37(globalState)).dtor_cf33);
      let _451_v40;
      _451_v40 = _dafny.Seq.UnicodeFromString("bobyhh");
      let _452_v41;
      _452_v41 = new _dafny.CodePoint('t'.codePointAt(0));
      let _rhs83 = !(_450_v39).contains((_449_v38).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm33(p0, p2, new BigNumber(945), globalState),p2)));
      let _rhs84 = _448_v37.f37;
      let _rhs85 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.update(_451_v40, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(((_module.__default.fm38(globalState)).dtor_cf40).length))).length), new BigNumber((_451_v40).length)), _452_v41), _451_v40), _451_v40);
      let _lhs73 = globalState;
      let _lhs74 = globalState;
      _lhs73.f18 = _rhs83;
      _lhs74.f15 = _rhs84;
      r1 = _rhs85;
      let _hi0 = p2;
      for (let _453_i6 = p2; _453_i6.isLessThan(_hi0); _453_i6 = _453_i6.plus(_dafny.ONE)) {
        (globalState).f1 = (p2).multipliedBy(p2);
        let _454_v42;
        let _nw74 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _454_v42 = _nw74;
        let _index66 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_454_v42).length));
        (_454_v42)[_index66] = p2;
        let _index67 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_454_v42).length));
        (_454_v42)[_index67] = _module.__default.fm36(globalState);
        let _455_v43;
        _455_v43 = _module.D6.create_DC22(p2, new BigNumber((_449_v38).length), (_403_v3)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_403_v3).length))]);
        let _456_v44;
        _456_v44 = _dafny.Map.Empty.slice().updateUnsafe(_403_v3,(_455_v43).dtor_cf33);
        _449_v38 = (_449_v38).update((p2).isLessThan(new BigNumber(((_456_v44).update(_403_v3, _453_i6)).length)), (_454_v42)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_454_v42).length))]);
        _400_v0 = (_400_v0).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_454_v42)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_454_v42).length))],_449_v38)).length), _module.__default.abs(_453_i6));
      }
      r0 = _451_v40;
      r1 = !(false);
      return [r0, r1];
    }
    get f40() {
      let _this = this;
      return _this._f40;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f30 = _dafny.Seq.UnicodeFromString("");
      this._f29 = false;
      this._f38 = _dafny.ZERO;
      this._f39 = _dafny.Set.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f30() {
      let _this = this;
      return _this._f30;
    };
    set f30(value) {
      let _this = this;
      _this._f30 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    __ctor(f38, f39, f29, f30) {
      let _this = this;
      (_this)._f38 = f38;
      (_this)._f39 = f39;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      return;
    }
    fm16(p0, globalState) {
      let _this = this;
      return (_this).f38;
    };
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      (globalState).f28 = (_this).f29;
      if (!((_this).f29) || (((_dafny.ZERO).minus((_this).f38)).isLessThan((_this).f38))) {
        (globalState).f15 = (_this).f29;
        (globalState).f2 = ((_this).f38).minus((_this).f38);
        let _457_v0;
        _457_v0 = _dafny.MultiSet.fromElements(p0);
        let _458_v1;
        _458_v1 = _module.D4.create_DC15(new BigNumber((_dafny.Seq.UnicodeFromString("ewafp")).length), new BigNumber(469));
        let _459_v2;
        let _nw75 = Array((new BigNumber(23)).toNumber()).fill(false);
        _459_v2 = _nw75;
        let _460_v3;
        _460_v3 = _dafny.MultiSet.fromElements(_459_v2, _459_v2);
        let _461_v4;
        _461_v4 = _dafny.Seq.of(p0);
        let _462_v5;
        let _nw76 = Array((new BigNumber(27)).toNumber());
        _nw76[(_dafny.ZERO).toNumber()] = (_this).f38;
        _nw76[(_dafny.ONE).toNumber()] = new BigNumber(838);
        _nw76[(new BigNumber(2)).toNumber()] = (_this).f38;
        _nw76[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt((_this).f38, (_this).f38);
        _nw76[(new BigNumber(4)).toNumber()] = new BigNumber(13);
        _nw76[(new BigNumber(5)).toNumber()] = (_this).fm16(_457_v0, globalState);
        _nw76[(new BigNumber(6)).toNumber()] = (_this).f38;
        _nw76[(new BigNumber(7)).toNumber()] = (_this).f38;
        _nw76[(new BigNumber(8)).toNumber()] = (_this).f38;
        _nw76[(new BigNumber(9)).toNumber()] = (_this).f38;
        _nw76[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus((_this).f38);
        _nw76[(new BigNumber(11)).toNumber()] = (_458_v1).dtor_cf25;
        _nw76[(new BigNumber(12)).toNumber()] = (_this).f38;
        _nw76[(new BigNumber(13)).toNumber()] = (_this).f38;
        _nw76[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt((_this).f38, new BigNumber((_module.__default.fm17((_this).f38, globalState)).length));
        _nw76[(new BigNumber(15)).toNumber()] = (_this).f38;
        _nw76[(new BigNumber(16)).toNumber()] = (_this).fm16(_457_v0, globalState);
        _nw76[(new BigNumber(17)).toNumber()] = ((_module.__default.fm18((_this).f38, new BigNumber((_460_v3).cardinality()), (_this).f38, new BigNumber((_461_v4).length), globalState)).dtor_cf15).multipliedBy((_this).f38);
        _nw76[(new BigNumber(18)).toNumber()] = (_this).f38;
        _nw76[(new BigNumber(19)).toNumber()] = new BigNumber((_461_v4).length);
        _nw76[(new BigNumber(20)).toNumber()] = (_this).f38;
        _nw76[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.of((_this).f29)).length);
        _nw76[(new BigNumber(22)).toNumber()] = _module.__default.safeDivisionInt((_this).f38, (_this).f38);
        _nw76[(new BigNumber(23)).toNumber()] = ((p0) ? ((_this).f38) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(18)), function (_463_i0) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })).length)));
        _nw76[(new BigNumber(24)).toNumber()] = (_this).f38;
        _nw76[(new BigNumber(25)).toNumber()] = (_this).f38;
        _nw76[(new BigNumber(26)).toNumber()] = (_this).f38;
        _462_v5 = _nw76;
        let _index68 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_462_v5).length));
        (_462_v5)[_index68] = (_this).f38;
        let _index69 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_462_v5).length));
        (_462_v5)[_index69] = (_this).f38;
        (globalState).f14 = (_this).f38;
        let _464_v6;
        _464_v6 = _dafny.Seq.of(_459_v2);
        (globalState).f1 = new BigNumber((_464_v6).length);
      } else {
        if (p0) {
          let _465_v7;
          let _nw77 = new _module.C0();
          _nw77.__ctor(p0);
          _465_v7 = _nw77;
          let _466_v8;
          _466_v8 = _dafny.MultiSet.fromElements(_465_v7.f37);
          let _467_v9;
          let _init8 = function (_468_i1) {
            return _module.__default.safeModuloInt(_468_i1, (_this).f38);
          };
          let _nw78 = Array((new BigNumber(16)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw78.length); _i0_8++) {
            _nw78[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _467_v9 = _nw78;
          let _469_v10;
          _469_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this.f30)[_module.__default.safeIndex((_this).fm16(_466_v8, globalState), new BigNumber((_this.f30).length))],_467_v9);
          let _470_v11;
          _470_v11 = new _dafny.CodePoint('e'.codePointAt(0));
          (globalState).f17 = (((_469_v10).contains(_470_v11)) ? ((_469_v10).get(_470_v11)) : (_467_v9));
          let _471_v12;
          let _init9 = ((_472_p0, _473_v7) => function (_474_i2) {
            return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_dafny.Set.fromElements(_472_p0, _472_p0)), _dafny.Seq.of(_dafny.Set.fromElements(_473_v7.f37), _dafny.Set.fromElements(true), _dafny.Set.fromElements(_473_v7.f37, _472_p0), _dafny.Set.fromElements(_473_v7.f37), _dafny.Set.fromElements((_this).f29)));
          })(p0, _465_v7);
          let _nw79 = Array((new BigNumber(2)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw79.length); _i0_9++) {
            _nw79[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _471_v12 = _nw79;
          _471_v12 = _471_v12;
          (globalState).f14 = (_this).f38;
          let _475_v13;
          _475_v13 = _dafny.Map.Empty.slice().updateUnsafe(_470_v11,_470_v11);
          let _476_v14;
          _476_v14 = _dafny.Seq.of(_475_v13);
          _476_v14 = _476_v14;
        } else {
          (globalState).f28 = p0;
          (globalState).f28 = ((_this).f38).isLessThan((_this).f38);
          let _477_v15;
          _477_v15 = _dafny.MultiSet.fromElements((_this).f38, (_this).f38, (_this).f38, (_this).f38, (_this).f38);
          let _478_v16;
          _478_v16 = _module.D3.create_DC10(_477_v15);
          let _479_v17;
          _479_v17 = new _dafny.CodePoint('i'.codePointAt(0));
          let _480_v18;
          let _nw80 = new _module.C0();
          _nw80.__ctor((_this).f29);
          _480_v18 = _nw80;
          let _481_v19;
          _481_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_this.f30).length),_480_v18);
          let _482_v20;
          _482_v20 = _dafny.Set.fromElements(new BigNumber((_481_v19).length), new BigNumber(-408));
          let _483_v21;
          _483_v21 = _dafny.Seq.of(p0);
          let _484_v22;
          _484_v22 = _dafny.Seq.of(new BigNumber((_482_v20).length), new BigNumber((_dafny.Seq.of(_480_v18.f37)).length), new BigNumber(-607), new BigNumber(730), new BigNumber((_483_v21).length));
          let _rhs86 = (_dafny.MultiSet.fromElements((_this).f38, (_this).f38, new BigNumber(283))).IsDisjointFrom(((_module.__default.fm19(_479_v17, globalState)) ? (_dafny.MultiSet.FromArray(_484_v22)) : (_477_v15)));
          let _rhs87 = _480_v18.f37;
          let _rhs88 = _478_v16;
          let _rhs89 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(817)), function (_485_i3) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          });
          let _lhs75 = globalState;
          let _lhs76 = globalState;
          let _lhs77 = globalState;
          _lhs75.f28 = _rhs86;
          _lhs76.f15 = _rhs87;
          _478_v16 = _rhs88;
          _lhs77.f11 = _rhs89;
          let _486_v23;
          _486_v23 = _dafny.MultiSet.fromElements((_this).f29, (_this).f29, _480_v18.f37, _480_v18.f37, _480_v18.f37);
          let _487_v24;
          _487_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_486_v23);
          let _488_v25;
          _488_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_this.f30).length));
          let _489_v26;
          _489_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_480_v18.f37);
          let _490_v27;
          let _nw81 = Array((new BigNumber(24)).toNumber());
          _nw81[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_this).f38);
          _nw81[(_dafny.ONE).toNumber()] = (_this).f38;
          _nw81[(new BigNumber(2)).toNumber()] = new BigNumber(106);
          _nw81[(new BigNumber(3)).toNumber()] = new BigNumber(-464);
          _nw81[(new BigNumber(4)).toNumber()] = new BigNumber(977);
          _nw81[(new BigNumber(5)).toNumber()] = (_this).f38;
          _nw81[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_this).f38);
          _nw81[(new BigNumber(7)).toNumber()] = (_this).fm16((((_487_v24).contains(_480_v18.f37)) ? ((_487_v24).get(_480_v18.f37)) : (_486_v23)), globalState);
          _nw81[(new BigNumber(8)).toNumber()] = (_this).f38;
          _nw81[(new BigNumber(9)).toNumber()] = ((!(!((_this).f29))) ? ((_this).f38) : ((_dafny.ZERO).minus((_this).f38)));
          _nw81[(new BigNumber(10)).toNumber()] = (_this).f38;
          _nw81[(new BigNumber(11)).toNumber()] = (_480_v18).fm12(globalState);
          _nw81[(new BigNumber(12)).toNumber()] = (_this).f38;
          _nw81[(new BigNumber(13)).toNumber()] = (_this).f38;
          _nw81[(new BigNumber(14)).toNumber()] = new BigNumber((_488_v25).length);
          _nw81[(new BigNumber(15)).toNumber()] = new BigNumber(754);
          _nw81[(new BigNumber(16)).toNumber()] = (_this).f38;
          _nw81[(new BigNumber(17)).toNumber()] = (((_this).f29) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_489_v26,_480_v18.f37)).length)) : (new BigNumber(-789)));
          _nw81[(new BigNumber(18)).toNumber()] = new BigNumber(788);
          _nw81[(new BigNumber(19)).toNumber()] = (_this).f38;
          _nw81[(new BigNumber(20)).toNumber()] = (new BigNumber((_482_v20).length)).plus((_this).f38);
          _nw81[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_484_v22).length));
          _nw81[(new BigNumber(22)).toNumber()] = (_this).fm16(_486_v23, globalState);
          _nw81[(new BigNumber(23)).toNumber()] = (_this).f38;
          _490_v27 = _nw81;
          let _index70 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_490_v27).length));
          (_490_v27)[_index70] = (_dafny.ZERO).minus((_this).f38);
          let _index71 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_490_v27).length));
          (_490_v27)[_index71] = (_this).f38;
          let _index72 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_490_v27).length));
          let _index73 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_490_v27).length));
          let _rhs90 = new BigNumber((_this.f30).length);
          let _rhs91 = (_480_v18).fm12(globalState);
          let _rhs92 = (_this).f38;
          let _rhs93 = _479_v17;
          let _lhs78 = _490_v27;
          let _lhs79 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_490_v27).length));
          let _lhs80 = globalState;
          let _lhs81 = _490_v27;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_490_v27).length));
          _lhs78[_lhs79] = _rhs90;
          _lhs80.f2 = _rhs91;
          _lhs81[_lhs82] = _rhs92;
          _479_v17 = _rhs93;
          let _491_v28;
          let _nw82 = Array((new BigNumber(24)).toNumber()).fill(false);
          _491_v28 = _nw82;
          let _index74 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_491_v28).length));
          (_491_v28)[_index74] = (_this).f29;
          let _index75 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_491_v28).length));
          (_491_v28)[_index75] = _480_v18.f37;
        }
        (globalState).f15 = p0;
        let _492_v29;
        _492_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(p0, (_this).f29, p0, (_this).f29, true)).length),!(!(true)));
        (globalState).f21 = new BigNumber(((_492_v29).Merge(_492_v29)).length);
        if (true) {
          let _493_v30;
          _493_v30 = _dafny.Seq.of(p0);
          let _494_v31;
          _494_v31 = _dafny.Seq.of(_dafny.Seq.of((_this).f29), _dafny.Seq.update(_493_v30, _module.__default.safeIndex((_this).f38, new BigNumber((_493_v30).length)), !(p0)));
          let _495_v32;
          _495_v32 = _dafny.Seq.of((_494_v31)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_this.f30).length)), new BigNumber((_494_v31).length))], _493_v30);
          let _496_v33;
          _496_v33 = new _dafny.CodePoint('t'.codePointAt(0));
          let _497_v34;
          let _nw83 = Array((new BigNumber(3)).toNumber());
          _nw83[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_this.f30, _module.__default.safeIndex(new BigNumber((_494_v31).length), new BigNumber((_this.f30).length)), _496_v33);
          _nw83[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_module.__default.fm20(true, (_this).f29, _493_v30, new BigNumber((_dafny.Seq.of(p0)).length), globalState), _module.__default.safeIndex((_this).f38, new BigNumber((_module.__default.fm20(true, (_this).f29, _493_v30, new BigNumber((_dafny.Seq.of(p0)).length), globalState)).length)), _496_v33);
          _nw83[(new BigNumber(2)).toNumber()] = _this.f30;
          _497_v34 = _nw83;
          let _index76 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_497_v34).length));
          (_497_v34)[_index76] = _dafny.Seq.UnicodeFromString("sar");
          let _498_v35;
          _498_v35 = _module.D1.create_DC3(_493_v30);
          let _pat_let_tv10 = _493_v30;
          let _499_v36;
          _499_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,function (_pat_let6_0) {
            return function (_500_dt__update__tmp_h0) {
              return function (_pat_let7_0) {
                return function (_501_dt__update_hcf7_h0) {
                  return _module.D1.create_DC3(_501_dt__update_hcf7_h0);
                }(_pat_let7_0);
              }(_pat_let_tv10);
            }(_pat_let6_0);
          }(_498_v35));
          let _502_v37;
          _502_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f30);
          let _503_v38;
          _503_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f38);
          let _504_v39;
          _504_v39 = _dafny.Seq.of((_this).f38, new BigNumber(567), new BigNumber((_dafny.Seq.update(_this.f30, _module.__default.safeIndex(new BigNumber((_503_v38).length), new BigNumber((_this.f30).length)), _496_v33)).length));
          let _index77 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_497_v34).length));
          let _rhs94 = _dafny.Seq.Concat((((_502_v37).contains((_this).f29)) ? ((_502_v37).get((_this).f29)) : (_dafny.Seq.update(_this.f30, _module.__default.safeIndex((_this).f38, new BigNumber((_this.f30).length)), _496_v33))), _module.__default.fm20(!(p0), true, _dafny.Seq.of((_this).f29, true, !((_493_v30)[_module.__default.safeIndex(new BigNumber(904), new BigNumber((_493_v30).length))]), p0), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((_this).f38, (_this).f38)).cardinality())), globalState));
          let _rhs95 = new BigNumber((_504_v39).length);
          let _rhs96 = ((_module.__default.fm21(p0, (_this).f29, p0, globalState)).dtor_cf26).Merge((_499_v36).Merge(_499_v36));
          let _lhs83 = _497_v34;
          let _lhs84 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_497_v34).length));
          let _lhs85 = globalState;
          _lhs83[_lhs84] = _rhs94;
          _lhs85.f19 = _rhs95;
          _499_v36 = _rhs96;
          (globalState).f18 = false;
          let _505_v40;
          let _nw84 = Array((new BigNumber(20)).toNumber()).fill(false);
          _505_v40 = _nw84;
          let _index78 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_505_v40).length));
          (_505_v40)[_index78] = (_this).f29;
          let _index79 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_505_v40).length));
          (_505_v40)[_index79] = ((_this).f38).isLessThan((_this).f38);
          let _506_v41;
          let _nw85 = Array((new BigNumber(12)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _506_v41 = _nw85;
          let _507_v42;
          _507_v42 = _dafny.Seq.of(_506_v41);
          _506_v41 = (_507_v42)[_module.__default.safeIndex(_module.__default.safeModuloInt((_this).f38, (_this).f38), new BigNumber((_507_v42).length))];
          (globalState).f28 = _module.__default.fm19(_496_v33, globalState);
        } else {
          let _508_v43;
          _508_v43 = new _dafny.CodePoint('s'.codePointAt(0));
          let _509_v44;
          _509_v44 = _dafny.MultiSet.fromElements(_508_v43);
          let _510_v45;
          _510_v45 = _module.D2.create_DC6(_509_v44);
          _510_v45 = _510_v45;
          let _511_v46;
          _511_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_492_v29).length),(_this).f38);
          _511_v46 = (_511_v46).update(new BigNumber(-643), (_this).f38);
          let _512_v47;
          let _nw86 = Array((new BigNumber(15)).toNumber()).fill(_dafny.MultiSet.Empty);
          _512_v47 = _nw86;
          let _513_v48;
          _513_v48 = _dafny.Seq.of(true);
          let _index80 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_512_v47).length));
          (_512_v47)[_index80] = _dafny.MultiSet.FromArray(_513_v48);
          let _514_v49;
          _514_v49 = _dafny.Seq.of(_this.f30);
          let _515_v50;
          _515_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f29);
          let _516_v51;
          _516_v51 = _dafny.MultiSet.fromElements((((_492_v29).contains((_this).f38)) ? ((_492_v29).get((_this).f38)) : (p0)), (_this).f29);
          let _517_v52;
          _517_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(301),_dafny.MultiSet.fromElements(p0, (_this).f29));
          let _index81 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_512_v47).length));
          let _rhs97 = _module.__default.safeDivisionInt(new BigNumber((_514_v49).length), (new BigNumber((_515_v50).length)).minus((_this).f38));
          let _rhs98 = (_516_v51).Difference((((_517_v52).contains((_this).f38)) ? ((_517_v52).get((_this).f38)) : (_dafny.MultiSet.FromArray(_513_v48))));
          let _lhs86 = globalState;
          let _lhs87 = _512_v47;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_512_v47).length));
          _lhs86.f1 = _rhs97;
          _lhs87[_lhs88] = _rhs98;
          let _518_v53;
          _518_v53 = _dafny.Seq.of(_513_v48);
          (globalState).f11 = _module.__default.fm20((_this).f29, p0, (_518_v53)[_module.__default.safeIndex((_this).f38, new BigNumber((_518_v53).length))], (_this).f38, globalState);
          let _519_v54;
          _519_v54 = _dafny.Set.fromElements(p0);
          let _520_v55;
          _520_v55 = _dafny.Seq.of((_module.D6.create_DC21(_519_v54)).dtor_cf32, _module.__default.fm22((_this).f38, !(p0), _508_v43, globalState), _519_v54);
          (globalState).f19 = new BigNumber(((_520_v55)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_520_v55).length))]).length);
        }
        let _521_v56;
        _521_v56 = _dafny.Map.Empty.slice().updateUnsafe((p0) === ((_this).f29),(_this).f38);
        _521_v56 = (_521_v56).update((_this).f29, _module.__default.safeDivisionInt((_this).f38, (_this).f38));
      }
      let _522_v57;
      let _init10 = function (_523_i4) {
        return _dafny.Seq.contains(_this.f30, new _dafny.CodePoint('p'.codePointAt(0)));
      };
      let _nw87 = Array((new BigNumber(15)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw87.length); _i0_10++) {
        _nw87[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _522_v57 = _nw87;
      let _524_v58;
      _524_v58 = new _dafny.CodePoint('c'.codePointAt(0));
      let _index82 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_522_v57).length));
      (_522_v57)[_index82] = !(_module.__default.fm19(_524_v58, globalState));
      let _index83 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_522_v57).length));
      (_522_v57)[_index83] = p0;
      (globalState).f19 = (_dafny.ZERO).minus((new BigNumber(-489)).multipliedBy((_this).f38));
      let _pat_let_tv11 = p0;
      let _pat_let_tv12 = p0;
      let _pat_let_tv13 = _522_v57;
      let _pat_let_tv14 = _522_v57;
      (globalState).f28 = function (_source8) {
        if (_source8.is_DC7) {
          let _525___mcc_h0 = (_source8).cf11;
          let _526_cf11 = _525___mcc_h0;
          return _dafny.Seq.IsProperPrefixOf(_this.f30, _this.f30);
        } else if (_source8.is_DC8) {
          let _527___mcc_h1 = (_source8).cf12;
          let _528___mcc_h2 = (_source8).cf13;
          let _529_cf13 = _528___mcc_h2;
          let _530_cf12 = _527___mcc_h1;
          return (_pat_let_tv11) && (_529_cf13);
        } else if (_source8.is_DC9) {
          let _531___mcc_h3 = (_source8).cf14;
          let _532___mcc_h4 = (_source8).cf15;
          let _533___mcc_h5 = (_source8).cf16;
          let _534___mcc_h6 = (_source8).cf17;
          let _535_cf17 = _534___mcc_h6;
          let _536_cf16 = _533___mcc_h5;
          let _537_cf15 = _532___mcc_h4;
          let _538_cf14 = _531___mcc_h3;
          return (false) === (_pat_let_tv12);
        } else {
          let _539___mcc_h7 = (_source8).cf10;
          let _540_cf10 = _539___mcc_h7;
          return (_pat_let_tv14)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_pat_let_tv13).length))];
        }
      }(_module.D2.create_DC9(_this.f30, (_this).f38, p0, (_this).f38));
      let _541_v59;
      _541_v59 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(p0) || (p0));
      _541_v59 = (_541_v59).update(!((_this).f29), p0);
      r0 = _module.__default.fm23(globalState);
      r1 = p0;
      return [r0, r1];
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let _542_v0;
      let _nw88 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _542_v0 = _nw88;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_542_v0).length))) {
        let _543_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_543_i0)) && ((_543_i0).isLessThan(new BigNumber((_542_v0).length))))) {
          (_542_v0)[(_543_i0)] = new _dafny.CodePoint('j'.codePointAt(0));
        }
      }
      let _544_v1;
      _544_v1 = _dafny.Set.fromElements((_this).f29);
      let _545_v3;
      _545_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(50));
      (globalState).f19 = (((_544_v1).IsSubsetOf(_544_v1)) ? ((_this).f38) : ((new BigNumber((function () {
        let _coll26 = new _dafny.Map();
        for (const _compr_26 of _dafny.IntegerRange(new BigNumber(588), new BigNumber(387))) {
          let _546_v2 = _compr_26;
          if (((new BigNumber(588)).isLessThanOrEqualTo(_546_v2)) && ((_546_v2).isLessThan(new BigNumber(387)))) {
            _coll26.push([_module.__default.safeModuloInt(_546_v2, p2),new BigNumber((_dafny.Seq.of(p2, (_this).f38, (_this).f38)).length)]);
          }
        }
        return _coll26;
      }()).length)).minus(new BigNumber((_545_v3).length))));
      let _547_v4;
      let _init11 = function (_548_i1) {
        return _module.__default.safeModuloInt(_548_i1, (_dafny.ZERO).minus(new BigNumber((_this.f30).length)));
      };
      let _nw89 = Array((new BigNumber(21)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw89.length); _i0_11++) {
        _nw89[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _547_v4 = _nw89;
      let _549_v5;
      _549_v5 = _dafny.MultiSet.fromElements(_547_v4);
      let _rhs99 = new BigNumber(727);
      let _rhs100 = _549_v5;
      let _lhs89 = globalState;
      _lhs89.f19 = _rhs99;
      _549_v5 = _rhs100;
      if (_module.__default.fm19(new _dafny.CodePoint('k'.codePointAt(0)), globalState)) {
        let _550_v6;
        let _nw90 = new _module.C0();
        _nw90.__ctor((_this).f29);
        _550_v6 = _nw90;
        (globalState).f18 = (_module.__default.fm24(globalState)).dtor_cf8;
        let _index84 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_547_v4).length));
        (_547_v4)[_index84] = new BigNumber(-588);
        let _551_v7;
        _551_v7 = _dafny.Seq.of((_this).f29);
        let _552_v8;
        _552_v8 = _dafny.Map.Empty.slice().updateUnsafe(_551_v7,(_this).f38);
        let _553_v9;
        _553_v9 = _module.D4.create_DC13(_552_v8);
        let _index85 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_547_v4).length));
        let _rhs101 = _module.__default.safeDivisionInt(new BigNumber((_this.f30).length), new BigNumber(855));
        let _rhs102 = _553_v9;
        let _rhs103 = true;
        let _lhs90 = _547_v4;
        let _lhs91 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_547_v4).length));
        let _lhs92 = globalState;
        _lhs90[_lhs91] = _rhs101;
        _553_v9 = _rhs102;
        _lhs92.f18 = _rhs103;
        let _554_v10;
        _554_v10 = _module.D2.create_DC9(_dafny.Seq.UnicodeFromString("xe"), (_547_v4)[_module.__default.safeIndex(new BigNumber(822), new BigNumber((_547_v4).length))], p0, new BigNumber((_this.f30).length));
        let _555_v11;
        _555_v11 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements((_this).f29, (_this).f29));
        let _556_v12;
        _556_v12 = _dafny.MultiSet.fromElements((_this).f29, _dafny.Seq.contains(_this.f30, new _dafny.CodePoint('t'.codePointAt(0))), ((((_555_v11).contains(_544_v1)) ? ((_555_v11).get(_544_v1)) : (p2))).isLessThanOrEqualTo((_554_v10).dtor_cf15));
        _556_v12 = _556_v12;
        let _557_v13;
        let _nw91 = Array((new BigNumber(23)).toNumber()).fill(false);
        _557_v13 = _nw91;
        let _index86 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_557_v13).length));
        (_557_v13)[_index86] = (_this).f29;
        let _558_v14;
        let _nw92 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Set.Empty);
        _558_v14 = _nw92;
        let _index87 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_558_v14).length));
        (_558_v14)[_index87] = (_544_v1).Union(_544_v1);
        let _559_v15;
        _559_v15 = new _dafny.CodePoint('h'.codePointAt(0));
        let _index88 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_547_v4).length));
        let _index89 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_557_v13).length));
        let _index90 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_558_v14).length));
        let _rhs104 = new BigNumber(211);
        let _rhs105 = ((_this).f38).isLessThanOrEqualTo(p2);
        let _rhs106 = (_547_v4)[_module.__default.safeIndex(new BigNumber(822), new BigNumber((_547_v4).length))];
        let _rhs107 = _module.__default.fm19(_559_v15, globalState);
        let _rhs108 = _544_v1;
        let _lhs93 = globalState;
        let _lhs94 = globalState;
        let _lhs95 = _547_v4;
        let _lhs96 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_547_v4).length));
        let _lhs97 = _557_v13;
        let _lhs98 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_557_v13).length));
        let _lhs99 = _558_v14;
        let _lhs100 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_558_v14).length));
        _lhs93.f14 = _rhs104;
        _lhs94.f18 = _rhs105;
        _lhs95[_lhs96] = _rhs106;
        _lhs97[_lhs98] = _rhs107;
        _lhs99[_lhs100] = _rhs108;
      } else {
        let _560_v16;
        _560_v16 = _dafny.Seq.of(new BigNumber((_this.f30).length));
        let _561_v17;
        _561_v17 = _dafny.Set.fromElements(_560_v16, (((_this).f29) ? (_560_v16) : (_dafny.Seq.of(new BigNumber(-539), p2))), _560_v16, _560_v16);
        _561_v17 = ((_561_v17).Difference(_561_v17)).Intersect(_module.__default.fm25((_this).f29, p0, _dafny.Seq.UnicodeFromString("lyyloxwhs"), globalState));
        (globalState).f1 = (_this).f38;
        let _562_v18;
        let _nw93 = Array((new BigNumber(17)).toNumber());
        _nw93[(_dafny.ZERO).toNumber()] = true;
        _nw93[(_dafny.ONE).toNumber()] = (_this).f29;
        _nw93[(new BigNumber(2)).toNumber()] = p0;
        _nw93[(new BigNumber(3)).toNumber()] = (_this).f29;
        _nw93[(new BigNumber(4)).toNumber()] = p0;
        _nw93[(new BigNumber(5)).toNumber()] = p0;
        _nw93[(new BigNumber(6)).toNumber()] = (_this).f29;
        _nw93[(new BigNumber(7)).toNumber()] = (_this).f29;
        _nw93[(new BigNumber(8)).toNumber()] = (((_this).f29) ? (true) : ((_this).f29));
        _nw93[(new BigNumber(9)).toNumber()] = (p2).isEqualTo(p2);
        _nw93[(new BigNumber(10)).toNumber()] = (_this).f29;
        _nw93[(new BigNumber(11)).toNumber()] = p0;
        _nw93[(new BigNumber(12)).toNumber()] = (_this).f29;
        _nw93[(new BigNumber(13)).toNumber()] = (_this).f29;
        _nw93[(new BigNumber(14)).toNumber()] = p0;
        _nw93[(new BigNumber(15)).toNumber()] = p0;
        _nw93[(new BigNumber(16)).toNumber()] = p0;
        _562_v18 = _nw93;
        let _index91 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_562_v18).length));
        (_562_v18)[_index91] = (_this).f29;
        let _563_v19;
        _563_v19 = new _dafny.CodePoint('x'.codePointAt(0));
        let _index92 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_562_v18).length));
        let _rhs109 = (_dafny.ZERO).minus((_560_v16)[_module.__default.safeIndex(p2, new BigNumber((_560_v16).length))]);
        let _rhs110 = (_this).f38;
        let _rhs111 = (((_dafny.ZERO).minus(p2)).multipliedBy((_this).f38)).isLessThanOrEqualTo((((_this).f29) ? (p2) : ((_this).f38)));
        let _rhs112 = p0;
        let _rhs113 = _563_v19;
        let _lhs101 = globalState;
        let _lhs102 = globalState;
        let _lhs103 = _562_v18;
        let _lhs104 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_562_v18).length));
        _lhs101.f14 = _rhs109;
        _lhs102.f14 = _rhs110;
        r1 = _rhs111;
        _lhs103[_lhs104] = _rhs112;
        _563_v19 = _rhs113;
        let _564_v20;
        _564_v20 = _module.D3.create_DC11((_this).f29, (_this).f29);
        let _565_v21;
        _565_v21 = _dafny.Map.Empty.slice().updateUnsafe(_564_v20,_563_v19);
        (globalState).f19 = (_module.__default.safeModuloInt(p2, (_dafny.ZERO).minus((_this).f38))).plus(new BigNumber((_565_v21).length));
        let _index93 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_562_v18).length));
        (_562_v18)[_index93] = (_this).f29;
      }
      let _566_v22;
      let _init12 = ((_567_p2) => function (_568_i3) {
        return !((_567_p2).isLessThan(_567_p2));
      })(p2);
      let _nw94 = Array((new BigNumber(11)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw94.length); _i0_12++) {
        _nw94[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _566_v22 = _nw94;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_566_v22).length))) {
        let _569_i2 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_569_i2)) && ((_569_i2).isLessThan(new BigNumber((_566_v22).length))))) {
          (_566_v22)[(_569_i2)] = ((new BigNumber((_dafny.MultiSet.fromElements((_module.D2.create_DC9(_this.f30, (_this).f38, p0, (_this).f38)).dtor_cf16, (_this).f29)).cardinality())).plus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(p0, p0))).cardinality())))).isLessThan(new BigNumber(474));
        }
      }
      if (p0) {
        (globalState).f20 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pwxebygro"), _this.f30);
        let _570_v23;
        _570_v23 = _dafny.Set.fromElements((_this).f38);
        let _571_v24;
        _571_v24 = _dafny.MultiSet.fromElements((_this).f38);
        let _rhs114 = _dafny.Set.fromElements((_this).f38);
        let _rhs115 = (_this).f29;
        let _rhs116 = (_570_v23).Union((_570_v23).Union(function () {
          let _coll27 = new _dafny.Set();
          for (const _compr_27 of (_571_v24).Elements) {
            let _572_v25 = _compr_27;
            if ((_571_v24).contains(_572_v25)) {
              _coll27.add((_572_v25).minus(new BigNumber((_dafny.Seq.UnicodeFromString("tsl")).length)));
            }
          }
          return _coll27;
        }()));
        let _lhs105 = globalState;
        _570_v23 = _rhs114;
        _lhs105.f28 = _rhs115;
        _570_v23 = _rhs116;
        let _573_v26;
        let _nw95 = new _module.C0();
        _nw95.__ctor(!((_this).f29));
        _573_v26 = _nw95;
        let _574_v27;
        let _nw96 = new _module.C0();
        _nw96.__ctor(_573_v26.f37);
        _574_v27 = _nw96;
        let _575_v28;
        let _576_v29;
        let _out10;
        let _out11;
        let _outcollector5 = (_this).m8(_566_v22, globalState);
        _out10 = _outcollector5[0];
        _out11 = _outcollector5[1];
        _575_v28 = _out10;
        _576_v29 = _out11;
      } else {
        (globalState).f19 = p2;
        let _577_v30;
        _577_v30 = _dafny.Seq.of((_this).f29);
        (globalState).f14 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f38, new BigNumber((_577_v30).length)));
        (globalState).f13 = (_this).f29;
        let _578_v31;
        _578_v31 = _dafny.Seq.of(_566_v22, _566_v22, _566_v22);
        let _579_v32;
        _579_v32 = _dafny.Set.fromElements(_566_v22, (_578_v31)[_module.__default.safeIndex(p2, new BigNumber((_578_v31).length))], _566_v22, _566_v22);
        (globalState).f13 = ((_dafny.Set.fromElements(_566_v22)).IsDisjointFrom(_579_v32)) === (p0);
        (globalState).f14 = p2;
      }
      r0 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("trbmjfc"), _this.f30);
      r1 = (_this).f29;
      return [r0, r1];
    }
    m8(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = _dafny.ZERO;
      let _580_v0;
      _580_v0 = _dafny.Seq.of((_this).f29);
      let _581_v1;
      _581_v1 = _dafny.Seq.of(new BigNumber((_580_v0).length), (_this).f38);
      let _582_v2;
      _582_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,((_this).fm16(_dafny.MultiSet.fromElements((_this).f29), globalState)).minus((_581_v1)[_module.__default.safeIndex((_this).f38, new BigNumber((_581_v1).length))]));
      _582_v2 = (_582_v2).update((_this).f38, (_this).f38);
      let _583_v3;
      let _init13 = function (_584_i1) {
        return (_this).f29;
      };
      let _nw97 = Array((new BigNumber(16)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw97.length); _i0_13++) {
        _nw97[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _583_v3 = _nw97;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_583_v3).length))) {
        let _585_i0 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_585_i0)) && ((_585_i0).isLessThan(new BigNumber((_583_v3).length))))) {
          (_583_v3)[(_585_i0)] = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(269)), function (_586_i2) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          }), _this.f30);
        }
      }
      (globalState).f13 = (_this).f29;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_583_v3).length))) {
        let _587_i3 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_587_i3)) && ((_587_i3).isLessThan(new BigNumber((_583_v3).length))))) {
          (_583_v3)[(_587_i3)] = (new BigNumber((_dafny.Seq.UnicodeFromString("fpqtgpu")).length)).isLessThanOrEqualTo((_this).f38);
        }
      }
      let _588_v4;
      _588_v4 = _dafny.MultiSet.fromElements(false);
      let _hi1 = (_this).f38;
      for (let _589_i4 = _module.__default.safeModuloInt((_this).fm16(_588_v4, globalState), (_this).f38); _589_i4.isLessThan(_hi1); _589_i4 = _589_i4.plus(_dafny.ONE)) {
        let _590_v5;
        _590_v5 = _dafny.Set.fromElements(_588_v4, (_588_v4).update((_this).f29, _module.__default.abs((_this).f38)));
        (globalState).f7 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_581_v1, _581_v1), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f38), new BigNumber((_dafny.Seq.Concat(_581_v1, _581_v1)).length)), new BigNumber((_590_v5).length)), _581_v1);
        let _591_v6;
        _591_v6 = _dafny.MultiSet.fromElements((_this).f38);
        let _592_v7;
        _592_v7 = _module.D3.create_DC10(_591_v6);
        let _593_v8;
        _593_v8 = _module.D3.create_DC12(_592_v7);
        let _source9 = _593_v8;
        if (_source9.is_DC11) {
          let _594___mcc_h0 = (_source9).cf19;
          let _595___mcc_h1 = (_source9).cf20;
          let _596_cf20 = _595___mcc_h1;
          let _597_cf19 = _594___mcc_h0;
          let _598_v9;
          _598_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,_dafny.Set.fromElements((_this).f38));
          let _599_v10;
          _599_v10 = _dafny.Set.fromElements(p0);
          _598_v9 = (_598_v9).update((_dafny.ZERO).minus(_589_i4), _dafny.Set.fromElements(_589_i4, new BigNumber((_this.f30).length), (_this).f38, new BigNumber((_599_v10).length)));
          let _600_v11;
          let _nw98 = new _module.C0();
          _nw98.__ctor((_this).f29);
          _600_v11 = _nw98;
          _582_v2 = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-507),(_this).f38)).Merge(_582_v2);
          let _601_v12;
          _601_v12 = _dafny.Map.Empty.slice().updateUnsafe(_600_v11.f37,_589_i4);
          let _602_v13;
          let _nw99 = Array((new BigNumber(26)).toNumber());
          _nw99[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.update(_580_v0, _module.__default.safeIndex(_589_i4, new BigNumber((_580_v0).length)), (_this).f29)).length);
          _nw99[(_dafny.ONE).toNumber()] = (_this).f38;
          _nw99[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_580_v0)).cardinality());
          _nw99[(new BigNumber(3)).toNumber()] = (_this).f38;
          _nw99[(new BigNumber(4)).toNumber()] = (_600_v11).fm12(globalState);
          _nw99[(new BigNumber(5)).toNumber()] = (_this).f38;
          _nw99[(new BigNumber(6)).toNumber()] = (_this).f38;
          _nw99[(new BigNumber(7)).toNumber()] = (_this).f38;
          _nw99[(new BigNumber(8)).toNumber()] = new BigNumber((_this.f30).length);
          _nw99[(new BigNumber(9)).toNumber()] = (_this).f38;
          _nw99[(new BigNumber(10)).toNumber()] = _589_i4;
          _nw99[(new BigNumber(11)).toNumber()] = (_this).f38;
          _nw99[(new BigNumber(12)).toNumber()] = new BigNumber(670);
          _nw99[(new BigNumber(13)).toNumber()] = (_this).f38;
          _nw99[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_this).f38);
          _nw99[(new BigNumber(15)).toNumber()] = (((_601_v12).contains(false)) ? ((_601_v12).get(false)) : ((_this).f38));
          _nw99[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus((_this).f38);
          _nw99[(new BigNumber(17)).toNumber()] = _589_i4;
          _nw99[(new BigNumber(18)).toNumber()] = (_this).fm16(_588_v4, globalState);
          _nw99[(new BigNumber(19)).toNumber()] = (_this).f38;
          _nw99[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(698)), function (_603_i5) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          })).length);
          _nw99[(new BigNumber(21)).toNumber()] = new BigNumber((_this.f30).length);
          _nw99[(new BigNumber(22)).toNumber()] = (_this).f38;
          _nw99[(new BigNumber(23)).toNumber()] = new BigNumber(329);
          _nw99[(new BigNumber(24)).toNumber()] = (_this).f38;
          _nw99[(new BigNumber(25)).toNumber()] = _589_i4;
          _602_v13 = _nw99;
          let _604_v14;
          _604_v14 = _dafny.Map.Empty.slice().updateUnsafe(_602_v13,!(_600_v11.f37));
          (globalState).f15 = (((_604_v14).contains(_602_v13)) ? ((_604_v14).get(_602_v13)) : (_596_cf20));
        } else if (_source9.is_DC10) {
          let _605___mcc_h2 = (_source9).cf18;
          let _606_cf18 = _605___mcc_h2;
          let _607_v15;
          let _nw100 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _607_v15 = _nw100;
          let _608_v16;
          _608_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm16(_588_v4, globalState),_607_v15);
          _608_v16 = (_608_v16).update(_589_i4, _607_v15);
          let _index94 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_583_v3).length));
          (_583_v3)[_index94] = (_this).f29;
          let _609_v17;
          _609_v17 = new _dafny.CodePoint('k'.codePointAt(0));
          let _610_v18;
          _610_v18 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),(_this).f38);
          let _611_v19;
          _611_v19 = _dafny.Map.Empty.slice().updateUnsafe(false,_610_v18);
          let _index95 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_583_v3).length));
          (_583_v3)[_index95] = ((_module.__default.fm19(_609_v17, globalState)) ? ((_this).f29) : (!((((_611_v19).contains((_this).f29)) ? ((_611_v19).get((_this).f29)) : (_module.__default.fm26((_this).f29, (_this).f38, (_this).f29, globalState)))).equals(_610_v18)));
          (globalState).f21 = _589_i4;
          let _612_v20;
          _612_v20 = _module.D3.create_DC10((_591_v6).update(_589_i4, _module.__default.abs((_this).f38)));
          let _613_v22;
          _613_v22 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_589_i4),(_this).f29);
          let _614_v23;
          let _nw101 = Array((new BigNumber(29)).toNumber());
          _nw101[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.FromArray(_581_v1);
          _nw101[(_dafny.ONE).toNumber()] = (_612_v20).dtor_cf18;
          _nw101[(new BigNumber(2)).toNumber()] = _606_cf18;
          _nw101[(new BigNumber(3)).toNumber()] = _606_cf18;
          _nw101[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_581_v1, _581_v1));
          _nw101[(new BigNumber(5)).toNumber()] = _module.__default.fm27((_this).f38, false, globalState);
          _nw101[(new BigNumber(6)).toNumber()] = _591_v6;
          _nw101[(new BigNumber(7)).toNumber()] = _606_cf18;
          _nw101[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.FromArray(_581_v1);
          _nw101[(new BigNumber(9)).toNumber()] = _606_cf18;
          _nw101[(new BigNumber(10)).toNumber()] = _591_v6;
          _nw101[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements(_589_i4, _589_i4);
          _nw101[(new BigNumber(12)).toNumber()] = (_606_cf18).update((_this).f38, _module.__default.abs(_589_i4));
          _nw101[(new BigNumber(13)).toNumber()] = (((_583_v3)[_module.__default.safeIndex(new BigNumber(417), new BigNumber((_583_v3).length))]) ? (_dafny.MultiSet.fromElements((_this).f38, _589_i4, _589_i4)) : (_606_cf18));
          _nw101[(new BigNumber(14)).toNumber()] = (_606_cf18).update(_589_i4, _module.__default.abs(_589_i4));
          _nw101[(new BigNumber(15)).toNumber()] = (_591_v6).Intersect(_591_v6);
          _nw101[(new BigNumber(16)).toNumber()] = _606_cf18;
          _nw101[(new BigNumber(17)).toNumber()] = (_606_cf18).Difference(_591_v6);
          _nw101[(new BigNumber(18)).toNumber()] = _606_cf18;
          _nw101[(new BigNumber(19)).toNumber()] = _module.__default.fm27((_this).f38, (_583_v3)[_module.__default.safeIndex(new BigNumber(417), new BigNumber((_583_v3).length))], globalState);
          _nw101[(new BigNumber(20)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((function () {
            let _coll28 = new _dafny.Map();
            for (const _compr_28 of (_613_v22).Keys.Elements) {
              let _615_v21 = _compr_28;
              if ((_613_v22).contains(_615_v21)) {
                _coll28.push([_module.__default.safeModuloInt(_615_v21, (_this).f38),new BigNumber((_606_cf18).cardinality())]);
              }
            }
            return _coll28;
          }()).length));
          _nw101[(new BigNumber(21)).toNumber()] = (_591_v6).Difference((_612_v20).dtor_cf18);
          _nw101[(new BigNumber(22)).toNumber()] = _606_cf18;
          _nw101[(new BigNumber(23)).toNumber()] = _606_cf18;
          _nw101[(new BigNumber(24)).toNumber()] = _606_cf18;
          _nw101[(new BigNumber(25)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber(79));
          _nw101[(new BigNumber(26)).toNumber()] = (_591_v6).Intersect(_591_v6);
          _nw101[(new BigNumber(27)).toNumber()] = _dafny.MultiSet.fromElements(_589_i4);
          _nw101[(new BigNumber(28)).toNumber()] = _dafny.MultiSet.fromElements((_this).f38, (_this).f38);
          _614_v23 = _nw101;
          let _index96 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_614_v23).length));
          (_614_v23)[_index96] = (_dafny.MultiSet.fromElements(_589_i4)).update(_589_i4, _module.__default.abs((_this).f38));
          let _index97 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_614_v23).length));
          let _rhs117 = _589_i4;
          let _rhs118 = (_591_v6).Union(_606_cf18);
          let _rhs119 = _dafny.Seq.update(_this.f30, _module.__default.safeIndex((_this).f38, new BigNumber((_this.f30).length)), _609_v17);
          let _lhs106 = globalState;
          let _lhs107 = _614_v23;
          let _lhs108 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_614_v23).length));
          let _lhs109 = globalState;
          _lhs106.f21 = _rhs117;
          _lhs107[_lhs108] = _rhs118;
          _lhs109.f11 = _rhs119;
        } else {
          let _616___mcc_h3 = (_source9).cf21;
          let _617_cf21 = _616___mcc_h3;
          let _index98 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_583_v3).length));
          (_583_v3)[_index98] = (_this).f29;
          let _index99 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_583_v3).length));
          (_583_v3)[_index99] = _dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of((_this).f29, (_this).f29), _580_v0), (_588_v4).contains(!((_this).f29)));
          (globalState).f13 = ((_this).fm16(_588_v4, globalState)).isEqualTo(_589_i4);
          (globalState).f28 = (_583_v3)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_583_v3).length))];
          let _618_v24;
          let _nw102 = Array((new BigNumber(24)).toNumber()).fill(false);
          _618_v24 = _nw102;
          let _619_v25;
          let _nw103 = Array((new BigNumber(28)).toNumber()).fill([]);
          _619_v25 = _nw103;
          let _620_v26;
          _620_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_583_v3)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_583_v3).length))]);
          let _621_v27;
          _621_v27 = new _dafny.CodePoint('v'.codePointAt(0));
          let _622_v28;
          _622_v28 = _620_v26;
          let _623_v29;
          let _nw104 = Array((new BigNumber(27)).toNumber());
          _nw104[(_dafny.ZERO).toNumber()] = _620_v26;
          _nw104[(_dafny.ONE).toNumber()] = _module.__default.fm28((_this).f29, _589_i4, (_this).f38, (_this).f29, globalState);
          _nw104[(new BigNumber(2)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(3)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(4)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(5)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f29),(_583_v3)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_583_v3).length))]);
          _nw104[(new BigNumber(7)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_module.__default.fm19(_621_v27, globalState));
          _nw104[(new BigNumber(9)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(10)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(11)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(12)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(13)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(14)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(15)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(16)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(17)).toNumber()] = (_622_v28);
          _nw104[(new BigNumber(18)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(19)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(20)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(21)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(22)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(23)).toNumber()] = _620_v26;
          _nw104[(new BigNumber(24)).toNumber()] = _module.__default.fm28((_583_v3)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_583_v3).length))], _589_i4, (_this).f38, (_this).f29, globalState);
          _nw104[(new BigNumber(25)).toNumber()] = _module.__default.fm28((_583_v3)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_583_v3).length))], (_this).f38, (_this).f38, (_this).f29, globalState);
          _nw104[(new BigNumber(26)).toNumber()] = _620_v26;
          _623_v29 = _nw104;
          let _index100 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_619_v25).length));
          (_619_v25)[_index100] = _623_v29;
          let _624_v30;
          let _nw105 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _624_v30 = _nw105;
          let _index101 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_624_v30).length));
          (_624_v30)[_index101] = (_dafny.ZERO).minus(_589_i4);
          let _index102 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_619_v25).length));
          let _index103 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_624_v30).length));
          let _rhs120 = _618_v24;
          let _rhs121 = _623_v29;
          let _rhs122 = (_589_i4).multipliedBy((_this).f38);
          let _lhs110 = _619_v25;
          let _lhs111 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_619_v25).length));
          let _lhs112 = _624_v30;
          let _lhs113 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_624_v30).length));
          _618_v24 = _rhs120;
          _lhs110[_lhs111] = _rhs121;
          _lhs112[_lhs113] = _rhs122;
        }
        (globalState).f13 = !(true) || ((_this).f29);
        (globalState).f14 = (_this).f38;
      }
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_583_v3).length))) {
        let _625_i6 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_625_i6)) && ((_625_i6).isLessThan(new BigNumber((_583_v3).length))))) {
          (_583_v3)[(_625_i6)] = ((_this).f38).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Concat(_this.f30, _this.f30)).length));
        }
      }
      let _626_v34;
      _626_v34 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f29);
      let _627_v35;
      _627_v35 = new _dafny.CodePoint('q'.codePointAt(0));
      let _628_v36;
      _628_v36 = _dafny.MultiSet.fromElements(_627_v35);
      let _629_v37;
      _629_v37 = _dafny.Set.fromElements(new BigNumber((_626_v34).length), new BigNumber((_628_v36).cardinality()), (_this).f38, new BigNumber(859));
      let _630_v38;
      _630_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(419)), ((_631_v35) => function (_632_i7) {
        return _631_v35;
      })(_627_v35))).length),_dafny.Set.fromElements((_this).f38));
      r0 = (function () {
        let _coll29 = new _dafny.Set();
        for (const _compr_29 of (_581_v1).Elements) {
          let _633_v31 = _compr_29;
          if (_dafny.Seq.contains(_581_v1, _633_v31)) {
            _coll29.add((_633_v31).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-479)), _dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((function () {
              let _coll30 = new _dafny.Set();
              for (const _compr_30 of (_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0)))).Elements) {
                let _634_v32 = _compr_30;
                if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0))), _634_v32)) {
                  _coll30.add(_634_v32);
                }
              }
              return _coll30;
            }()).length)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(317)), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((function () {
              let _coll31 = new _dafny.Map();
              for (const _compr_31 of _dafny.IntegerRange(new BigNumber(-483), new BigNumber(-255))) {
                let _635_v33 = _compr_31;
                if (((new BigNumber(-483)).isLessThanOrEqualTo(_635_v33)) && ((_635_v33).isLessThan(new BigNumber(-255)))) {
                  _coll31.push([_module.__default.safeModuloInt(_635_v33, new BigNumber((_dafny.Set.fromElements(new BigNumber(-81))).length)),true]);
                }
              }
              return _coll31;
            }()).length)))).cardinality())));
          }
        }
        return _coll29;
      }()).Intersect((_629_v37).Intersect((((_630_v38).contains((_this).f38)) ? ((_630_v38).get((_this).f38)) : (_629_v37))));
      let _636_v39;
      _636_v39 = _module.D3.create_DC12(_module.__default.fm29((_this).f29, (_this).f38, (_this).f29, globalState));
      let _637_v40;
      _637_v40 = _module.D3.create_DC11((_580_v0)[_module.__default.safeIndex(new BigNumber(587), new BigNumber((_580_v0).length))], (_this).f29);
      let _pat_let_tv15 = _637_v40;
      r1 = function (_source10) {
        if (_source10.is_DC11) {
          let _638___mcc_h4 = (_source10).cf19;
          let _639___mcc_h5 = (_source10).cf20;
          let _640_cf20 = _639___mcc_h5;
          let _641_cf19 = _638___mcc_h4;
          return new BigNumber(967);
        } else if (_source10.is_DC10) {
          let _642___mcc_h6 = (_source10).cf18;
          let _643_cf18 = _642___mcc_h6;
          return (_this).f38;
        } else {
          let _644___mcc_h7 = (_source10).cf21;
          let _645_cf21 = _644___mcc_h7;
          return (_this).f38;
        }
      }(function (_pat_let8_0) {
        return function (_646_dt__update__tmp_h0) {
          return function (_pat_let9_0) {
            return function (_647_dt__update_hcf21_h0) {
              return _module.D3.create_DC12(_647_dt__update_hcf21_h0);
            }(_pat_let9_0);
          }(_pat_let_tv15);
        }(_pat_let8_0);
      }(_636_v39));
      return [r0, r1];
    }
    m9(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let _648_v0;
      _648_v0 = _dafny.MultiSet.fromElements(p1, (_dafny.ZERO).minus((_this).f38), p1);
      let _649_v1;
      _649_v1 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f38);
      let _650_v2;
      _650_v2 = _dafny.Seq.of((_this).f38);
      let _651_v3;
      _651_v3 = _dafny.Map.Empty.slice().updateUnsafe(_649_v1,(_dafny.MultiSet.FromArray(_650_v2)).update(new BigNumber((_this.f30).length), _module.__default.abs(p1)));
      _648_v0 = (((_651_v3).contains(_649_v1)) ? ((_651_v3).get(_649_v1)) : ((_648_v0).Difference(_648_v0)));
      (globalState).f13 = (_this).f29;
      (globalState).f13 = p2;
      (globalState).f28 = !((_this).f29) || (p2);
      let _652_v4;
      _652_v4 = _module.D0.create_DC0((((_this).f29) ? (p3) : (!((_this).f29))));
      _652_v4 = (((_this).f29) ? (_652_v4) : (_652_v4));
      let _653_v5;
      let _init14 = ((_654_p2) => function (_655_i0) {
        return _654_p2;
      })(p2);
      let _nw106 = Array((new BigNumber(12)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw106.length); _i0_14++) {
        _nw106[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _653_v5 = _nw106;
      let _index104 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_653_v5).length));
      (_653_v5)[_index104] = false;
      let _656_v6;
      _656_v6 = new _dafny.CodePoint('q'.codePointAt(0));
      let _657_v7;
      _657_v7 = _dafny.Set.fromElements(_656_v6);
      let _pat_let_tv16 = p2;
      let _pat_let_tv17 = p3;
      let _pat_let_tv18 = p3;
      let _index105 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_653_v5).length));
      (_653_v5)[_index105] = function (_source11) {
        if (_source11.is_DC7) {
          let _658___mcc_h0 = (_source11).cf11;
          let _659_cf11 = _658___mcc_h0;
          return (_pat_let_tv16) && (_659_cf11);
        } else if (_source11.is_DC8) {
          let _660___mcc_h1 = (_source11).cf12;
          let _661___mcc_h2 = (_source11).cf13;
          let _662_cf13 = _661___mcc_h2;
          let _663_cf12 = _660___mcc_h1;
          return _662_cf13;
        } else if (_source11.is_DC9) {
          let _664___mcc_h3 = (_source11).cf14;
          let _665___mcc_h4 = (_source11).cf15;
          let _666___mcc_h5 = (_source11).cf16;
          let _667___mcc_h6 = (_source11).cf17;
          let _668_cf17 = _667___mcc_h6;
          let _669_cf16 = _666___mcc_h5;
          let _670_cf15 = _665___mcc_h4;
          let _671_cf14 = _664___mcc_h3;
          return _pat_let_tv17;
        } else {
          let _672___mcc_h7 = (_source11).cf10;
          let _673_cf10 = _672___mcc_h7;
          return !(_pat_let_tv18);
        }
      }(function (_pat_let10_0) {
        return function (_675_dt__update__tmp_h0) {
          return function (_pat_let11_0) {
            return function (_676_dt__update_hcf14_h0) {
              return function (_pat_let12_0) {
                return function (_677_dt__update_hcf16_h0) {
                  return function (_pat_let13_0) {
                    return function (_678_dt__update_hcf15_h0) {
                      return _module.D2.create_DC9(_676_dt__update_hcf14_h0, _678_dt__update_hcf15_h0, _677_dt__update_hcf16_h0, (_675_dt__update__tmp_h0).dtor_cf17);
                    }(_pat_let13_0);
                  }(new BigNumber(271));
                }(_pat_let12_0);
              }(true);
            }(_pat_let11_0);
          }(_this.f30);
        }(_pat_let10_0);
      }(_module.D2.create_DC9(_dafny.Seq.Create(_module.__default.abs(new BigNumber(716)), function (_674_i1) {
  return new _dafny.CodePoint('e'.codePointAt(0));
}), (_dafny.ZERO).minus(new BigNumber((_657_v7).length)), p3, p1)));
      let _679_v8;
      let _nw107 = new _module.C1();
      _nw107.__ctor(p2);
      _679_v8 = _nw107;
      let _680_v9;
      _680_v9 = _module.D10.create_DC30(_679_v8);
      let _681_v10;
      _681_v10 = _dafny.Set.fromElements(_679_v8, (_680_v9).dtor_cf49, _679_v8);
      let _682_v11;
      _682_v11 = _dafny.Set.fromElements((_681_v10).Intersect(_681_v10));
      r0 = _682_v11;
      return r0;
    }
    get f38() {
      let _this = this;
      return _this._f38;
    };
    get f39() {
      let _this = this;
      return _this._f39;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f30 = _dafny.Seq.UnicodeFromString("");
      this._f29 = false;
      this._f35 = false;
      this._f36 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f30() {
      let _this = this;
      return _this._f30;
    };
    set f30(value) {
      let _this = this;
      _this._f30 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    __ctor(f35, f36, f29, f30) {
      let _this = this;
      (_this)._f35 = f35;
      (_this)._f36 = f36;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      return;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _683_v0;
      let _nw108 = Array((new BigNumber(7)).toNumber()).fill([]);
      _683_v0 = _nw108;
      let _684_v1;
      let _init15 = function (_685_i4) {
        return (_this).f35;
      };
      let _nw109 = Array((new BigNumber(20)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw109.length); _i0_15++) {
        _nw109[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _684_v1 = _nw109;
      let _686_v2;
      _686_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f35,_684_v1);
      let _687_v3;
      _687_v3 = new _dafny.CodePoint('f'.codePointAt(0));
      let _688_v4;
      let _nw110 = Array((new BigNumber(24)).toNumber());
      _nw110[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(172)), function (_689_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      });
      _nw110[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), function (_690_i1) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      });
      _nw110[(new BigNumber(2)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(3)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(4)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(5)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(514)), function (_691_i2) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      });
      _nw110[(new BigNumber(7)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(8)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(9)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(10)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("cdiohjl");
      _nw110[(new BigNumber(12)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(13)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(14)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(107)), function (_692_i3) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      });
      _nw110[(new BigNumber(16)).toNumber()] = _dafny.Seq.update(_this.f30, _module.__default.safeIndex(new BigNumber((_686_v2).length), new BigNumber((_this.f30).length)), _687_v3);
      _nw110[(new BigNumber(17)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(18)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(19)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(20)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(21)).toNumber()] = _this.f30;
      _nw110[(new BigNumber(22)).toNumber()] = _dafny.Seq.update(_this.f30, _module.__default.safeIndex((_this).f36, new BigNumber((_this.f30).length)), _687_v3);
      _nw110[(new BigNumber(23)).toNumber()] = _this.f30;
      _688_v4 = _nw110;
      let _index106 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_683_v0).length));
      (_683_v0)[_index106] = _688_v4;
      let _index107 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_683_v0).length));
      (_683_v0)[_index107] = _688_v4;
      let _693_v5;
      _693_v5 = _dafny.Seq.of(p0, p0);
      let _source12 = _module.D1.create_DC3(_dafny.Seq.Concat(_693_v5, _dafny.Seq.of(true)));
      if (_source12.is_DC4) {
        let _694___mcc_h0 = (_source12).cf8;
        let _695_cf8 = _694___mcc_h0;
        let _696_v6;
        let _init16 = function (_697_i5) {
          return (_697_i5).minus((_this).f36);
        };
        let _nw111 = Array((new BigNumber(7)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw111.length); _i0_16++) {
          _nw111[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _696_v6 = _nw111;
        (globalState).f17 = _696_v6;
        (globalState).f15 = p0;
        let _698_v7;
        _698_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f36,_687_v3);
        let _699_v8;
        _699_v8 = _module.D2.create_DC8((((_698_v7).contains((_this).f36)) ? ((_698_v7).get((_this).f36)) : (_687_v3)), (_this).f29);
        let _source13 = _699_v8;
        if (_source13.is_DC7) {
          let _700___mcc_h3 = (_source13).cf11;
          let _701_cf11 = _700___mcc_h3;
          let _702_v9;
          _702_v9 = _dafny.Seq.of(_693_v5);
          r0 = _702_v9;
          let _703_v10;
          let _nw112 = Array((new BigNumber(23)).toNumber()).fill(_module.D3.Default());
          _703_v10 = _nw112;
          let _704_v11;
          _704_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uogups"), _this.f30),_703_v10);
          _704_v11 = (_704_v11).update(_dafny.Seq.Concat(_this.f30, _dafny.Seq.update(_this.f30, _module.__default.safeIndex((_this).f36, new BigNumber((_this.f30).length)), _687_v3)), _703_v10);
          r1 = _701_cf11;
          let _705_v12;
          _705_v12 = _dafny.Set.fromElements(_701_cf11, (_this).f35);
          (globalState).f19 = new BigNumber(((_705_v12).Difference(_705_v12)).length);
        } else if (_source13.is_DC8) {
          let _706___mcc_h4 = (_source13).cf12;
          let _707___mcc_h5 = (_source13).cf13;
          let _708_cf13 = _707___mcc_h5;
          let _709_cf12 = _706___mcc_h4;
          let _710_v13;
          _710_v13 = _module.D2.create_DC7((_this).f29);
          (globalState).f15 = (_710_v13).dtor_cf11;
          let _711_v14;
          _711_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(22),_684_v1);
          _711_v14 = (_711_v14).update(new BigNumber((_693_v5).length), (((_this).f29) ? (_684_v1) : (_684_v1)));
          let _712_v15;
          _712_v15 = _module.D0.create_DC1(_684_v1, (_this).f36, (_this).f36);
          let _rhs123 = _module.__default.safeDivisionInt((((_this).f35) ? ((_this).f36) : (_module.__default.fm9(globalState))), (new BigNumber((_this.f30).length)).multipliedBy((_this).f36));
          let _rhs124 = (_this).f36;
          let _rhs125 = _dafny.Seq.update(_dafny.Seq.of((_module.__default.fm10((_this).f36, (_this).f36, globalState)) === ((_693_v5)[_module.__default.safeIndex((_this).f36, new BigNumber((_693_v5).length))])), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f36), new BigNumber((_dafny.Seq.of((_module.__default.fm10((_this).f36, (_this).f36, globalState)) === ((_693_v5)[_module.__default.safeIndex((_this).f36, new BigNumber((_693_v5).length))]))).length)), (_this).f35);
          let _rhs126 = (_712_v15).dtor_cf2;
          let _rhs127 = ((_695_cf8) ? (_710_v13) : (_710_v13));
          let _lhs114 = globalState;
          let _lhs115 = globalState;
          let _lhs116 = globalState;
          _lhs114.f1 = _rhs123;
          _lhs115.f14 = _rhs124;
          _693_v5 = _rhs125;
          _lhs116.f21 = _rhs126;
          _710_v13 = _rhs127;
          (globalState).f1 = ((_this).f36).plus((_this).f36);
        } else if (_source13.is_DC9) {
          let _713___mcc_h6 = (_source13).cf14;
          let _714___mcc_h7 = (_source13).cf15;
          let _715___mcc_h8 = (_source13).cf16;
          let _716___mcc_h9 = (_source13).cf17;
          let _717_cf17 = _716___mcc_h9;
          let _718_cf16 = _715___mcc_h8;
          let _719_cf15 = _714___mcc_h7;
          let _720_cf14 = _713___mcc_h6;
          _687_v3 = _687_v3;
          (globalState).f2 = (_dafny.ZERO).minus((_this).f36);
          let _721_v16;
          let _nw113 = Array((new BigNumber(11)).toNumber()).fill(_module.D1.Default());
          _721_v16 = _nw113;
          let _722_v17;
          _722_v17 = _module.D1.create_DC3(_693_v5);
          let _723_v18;
          _723_v18 = _module.D1.create_DC5(_722_v17);
          let _index108 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_721_v16).length));
          (_721_v16)[_index108] = _723_v18;
          let _index109 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_721_v16).length));
          (_721_v16)[_index109] = _723_v18;
          let _724_v19;
          _724_v19 = _dafny.Map.Empty.slice().updateUnsafe(_696_v6,(_this).f36);
          (globalState).f2 = new BigNumber((_724_v19).length);
        } else {
          let _725___mcc_h10 = (_source13).cf10;
          let _726_cf10 = _725___mcc_h10;
          (globalState).f1 = (((_this).f35) ? ((_this).f36) : (((_this).f36).minus((_this).f36)));
          (globalState).f11 = _this.f30;
          let _727_v20;
          let _nw114 = new _module.C0();
          _nw114.__ctor(p0);
          _727_v20 = _nw114;
          let _728_v21;
          let _init17 = ((_729_v3) => function (_730_i6) {
            return _729_v3;
          })(_687_v3);
          let _nw115 = Array((new BigNumber(15)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw115.length); _i0_17++) {
            _nw115[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _728_v21 = _nw115;
          let _index110 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_728_v21).length));
          (_728_v21)[_index110] = (_this.f30)[_module.__default.safeIndex((_this).f36, new BigNumber((_this.f30).length))];
          let _index111 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_728_v21).length));
          (_728_v21)[_index111] = _687_v3;
        }
        let _index112 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_696_v6).length));
        (_696_v6)[_index112] = (_this).f36;
        let _index113 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_696_v6).length));
        (_696_v6)[_index113] = new BigNumber((_this.f30).length);
      } else if (_source12.is_DC3) {
        let _731___mcc_h1 = (_source12).cf7;
        let _732_cf7 = _731___mcc_h1;
        let _733_v22;
        _733_v22 = _dafny.Map.Empty.slice().updateUnsafe(_732_cf7,(_dafny.ZERO).minus((_this).f36));
        _733_v22 = ((_733_v22).Merge((_module.__default.fm13((_this).f36, new BigNumber((function () {
          let _coll32 = new _dafny.Map();
          for (const _compr_32 of _dafny.IntegerRange(new BigNumber(-784), new BigNumber(-112))) {
            let _734_v23 = _compr_32;
            if (((new BigNumber(-784)).isLessThanOrEqualTo(_734_v23)) && ((_734_v23).isLessThan(new BigNumber(-112)))) {
              _coll32.push([_module.__default.safeModuloInt(_734_v23, new BigNumber(471)),(_this).f36]);
            }
          }
          return _coll32;
        }()).length), (_this).f36, !((_this).f35), globalState)).dtor_cf22)).Merge(_733_v22);
        (globalState).f20 = _this.f30;
        (globalState).f19 = (_this).f36;
        let _735_v24;
        _735_v24 = _module.D1.create_DC4(!(p0) || (p0));
        let _source14 = _735_v24;
        if (_source14.is_DC4) {
          let _736___mcc_h11 = (_source14).cf8;
          let _737_cf8 = _736___mcc_h11;
          let _738_v25;
          let _init18 = ((_739_v3) => function (_740_i7) {
            return _739_v3;
          })(_687_v3);
          let _nw116 = Array((new BigNumber(6)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw116.length); _i0_18++) {
            _nw116[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _738_v25 = _nw116;
          let _index114 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_738_v25).length));
          (_738_v25)[_index114] = _687_v3;
          let _741_v26;
          _741_v26 = _dafny.MultiSet.fromElements((_this).f36);
          let _742_v27;
          _742_v27 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f35),_741_v26);
          let _743_v28;
          _743_v28 = _module.D3.create_DC10((((_742_v27).contains(_737_cf8)) ? ((_742_v27).get(_737_cf8)) : (_741_v26)));
          let _744_v29;
          _744_v29 = _module.D3.create_DC12(_743_v28);
          let _745_v30;
          _745_v30 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(774)).isLessThanOrEqualTo((_this).f36),_744_v29);
          let _746_v31;
          _746_v31 = _module.D2.create_DC8(_687_v3, p0);
          let _747_v32;
          let _nw117 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _747_v32 = _nw117;
          let _pat_let_tv19 = _743_v28;
          let _index115 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_738_v25).length));
          let _rhs128 = (_746_v31).dtor_cf12;
          let _rhs129 = (((_745_v30).update((_this).f35, _744_v29)).Merge((_745_v30).update(false, function (_pat_let14_0) {
            return function (_748_dt__update__tmp_h0) {
              return function (_pat_let15_0) {
                return function (_749_dt__update_hcf21_h0) {
                  return _module.D3.create_DC12(_749_dt__update_hcf21_h0);
                }(_pat_let15_0);
              }(_pat_let_tv19);
            }(_pat_let14_0);
          }(_744_v29)))).Merge(_745_v30);
          let _rhs130 = _684_v1;
          let _rhs131 = _747_v32;
          let _lhs117 = _738_v25;
          let _lhs118 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_738_v25).length));
          let _lhs119 = globalState;
          _lhs117[_lhs118] = _rhs128;
          _745_v30 = _rhs129;
          _684_v1 = _rhs130;
          _lhs119.f17 = _rhs131;
          let _750_v33;
          let _nw118 = new _module.C0();
          _nw118.__ctor((((_this).f29) ? (true) : (p0)));
          _750_v33 = _nw118;
          let _751_v34;
          _751_v34 = _dafny.Map.Empty.slice().updateUnsafe(_737_cf8,new BigNumber((_693_v5).length));
          let _752_v36;
          _752_v36 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_750_v33.f37,(_this).f36), _751_v34, _module.__default.fm14((_this).f35, _dafny.Seq.Create(_module.__default.abs(new BigNumber(900)), function (_753_i9) {
            return (_this).f36;
          }), _737_cf8, new BigNumber((_dafny.Set.fromElements(_684_v1)).length), globalState), _751_v34);
          let _754_v38;
          _754_v38 = _dafny.Set.fromElements(p0, _750_v33.f37);
          let _755_v39;
          _755_v39 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_754_v38).length)));
          let _756_v40;
          _756_v40 = _dafny.Map.Empty.slice().updateUnsafe(_751_v34,true);
          let _757_v42;
          _757_v42 = _dafny.Map.Empty.slice().updateUnsafe(_751_v34,(_this).f36);
          let _758_v44;
          let _nw119 = Array((new BigNumber(9)).toNumber());
          _nw119[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(_751_v34, _751_v34);
          _nw119[(_dafny.ONE).toNumber()] = function () {
            let _coll33 = new _dafny.Set();
            for (const _compr_33 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(677)), ((_759_v34) => function (_760_i8) {
              return _759_v34;
            })(_751_v34))).Elements) {
              let _761_v35 = _compr_33;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(677)), ((_762_v34) => function (_760_i8) {
                return _762_v34;
              })(_751_v34)), _761_v35)) {
                _coll33.add(_761_v35);
              }
            }
            return _coll33;
          }();
          _nw119[(new BigNumber(2)).toNumber()] = function () {
            let _coll34 = new _dafny.Set();
            for (const _compr_34 of (_752_v36).Elements) {
              let _763_v37 = _compr_34;
              if ((_752_v36).contains(_763_v37)) {
                _coll34.add(_763_v37);
              }
            }
            return _coll34;
          }();
          _nw119[(new BigNumber(3)).toNumber()] = (_755_v39).Intersect(function () {
            let _coll35 = new _dafny.Set();
            for (const _compr_35 of (_756_v40).Keys.Elements) {
              let _764_v41 = _compr_35;
              if ((_756_v40).contains(_764_v41)) {
                _coll35.add(_764_v41);
              }
            }
            return _coll35;
          }());
          _nw119[(new BigNumber(4)).toNumber()] = _755_v39;
          _nw119[(new BigNumber(5)).toNumber()] = function () {
            let _coll36 = new _dafny.Set();
            for (const _compr_36 of ((_757_v42).update(_751_v34, (_this).f36)).Keys.Elements) {
              let _765_v43 = _compr_36;
              if (((_757_v42).update(_751_v34, (_this).f36)).contains(_765_v43)) {
                _coll36.add(_765_v43);
              }
            }
            return _coll36;
          }();
          _nw119[(new BigNumber(6)).toNumber()] = _755_v39;
          _nw119[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements(_751_v34, _751_v34, (_751_v34).update(_737_cf8, new BigNumber((_754_v38).length)), _751_v34, _751_v34);
          _nw119[(new BigNumber(8)).toNumber()] = (_dafny.Set.fromElements(_751_v34)).Intersect(_755_v39);
          _758_v44 = _nw119;
          let _index116 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_758_v44).length));
          (_758_v44)[_index116] = _755_v39;
          let _index117 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_758_v44).length));
          (_758_v44)[_index117] = _module.__default.fm15((_this).f36, globalState);
          (globalState).f19 = (_dafny.ZERO).minus((_this).f36);
        } else if (_source14.is_DC3) {
          let _766___mcc_h12 = (_source14).cf7;
          let _767_cf7 = _766___mcc_h12;
          (globalState).f1 = new BigNumber(58);
          let _768_v45;
          let _nw120 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _768_v45 = _nw120;
          let _index118 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_768_v45).length));
          (_768_v45)[_index118] = (_this).f36;
          let _index119 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_768_v45).length));
          (_768_v45)[_index119] = new BigNumber(943);
          (globalState).f21 = (_768_v45)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_768_v45).length))];
          (globalState).f28 = p0;
        } else {
          let _769___mcc_h13 = (_source14).cf9;
          let _770_cf9 = _769___mcc_h13;
          let _771_v46;
          _771_v46 = _dafny.Set.fromElements(p0, (_this).f29);
          let _772_v47;
          _772_v47 = _dafny.Seq.of(_771_v46);
          (globalState).f3 = _dafny.Seq.Concat(_772_v47, _772_v47);
          (globalState).f1 = (_module.__default.fm9(globalState)).plus((_this).f36);
          let _773_v48;
          let _nw121 = new _module.C0();
          _nw121.__ctor((((_this).f35) ? (false) : ((_this).f35)));
          _773_v48 = _nw121;
          let _index120 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_684_v1).length));
          (_684_v1)[_index120] = (_this).f29;
          let _index121 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_684_v1).length));
          (_684_v1)[_index121] = (_this).f35;
        }
      } else {
        let _774___mcc_h2 = (_source12).cf9;
        let _775_cf9 = _774___mcc_h2;
        let _776_v49;
        _776_v49 = _dafny.Map.Empty.slice().updateUnsafe((_this).f35,(_this).f36);
        let _777_v50;
        let _nw122 = Array((new BigNumber(17)).toNumber());
        _nw122[(_dafny.ZERO).toNumber()] = (_this).f36;
        _nw122[(_dafny.ONE).toNumber()] = new BigNumber((_this.f30).length);
        _nw122[(new BigNumber(2)).toNumber()] = (_this).f36;
        _nw122[(new BigNumber(3)).toNumber()] = (_this).f36;
        _nw122[(new BigNumber(4)).toNumber()] = (_this).f36;
        _nw122[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-598)), function (_778_i10) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })).length);
        _nw122[(new BigNumber(6)).toNumber()] = (_this).f36;
        _nw122[(new BigNumber(7)).toNumber()] = (_this).f36;
        _nw122[(new BigNumber(8)).toNumber()] = new BigNumber((_776_v49).length);
        _nw122[(new BigNumber(9)).toNumber()] = (_this).f36;
        _nw122[(new BigNumber(10)).toNumber()] = new BigNumber((_693_v5).length);
        _nw122[(new BigNumber(11)).toNumber()] = (_this).f36;
        _nw122[(new BigNumber(12)).toNumber()] = (_this).f36;
        _nw122[(new BigNumber(13)).toNumber()] = (_this).f36;
        _nw122[(new BigNumber(14)).toNumber()] = (_this).f36;
        _nw122[(new BigNumber(15)).toNumber()] = (_this).f36;
        _nw122[(new BigNumber(16)).toNumber()] = (_this).f36;
        _777_v50 = _nw122;
        let _779_v51;
        _779_v51 = _dafny.Seq.of(_777_v50, _777_v50);
        let _780_v52;
        _780_v52 = _dafny.Set.fromElements(_777_v50, (_779_v51)[_module.__default.safeIndex(_module.__default.fm9(globalState), new BigNumber((_779_v51).length))]);
        let _781_v53;
        let _nw123 = new _module.C2();
        _nw123.__ctor((_this).f36, _780_v52, p0, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gt"), _this.f30));
        _781_v53 = _nw123;
        _781_v53 = _781_v53;
        let _782_v54;
        _782_v54 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("vccvow"));
        (globalState).f28 = (_782_v54).contains(_this.f30);
        let _783_v55;
        let _init19 = ((_784_v3) => function (_785_i11) {
          return _784_v3;
        })(_687_v3);
        let _nw124 = Array((new BigNumber(17)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw124.length); _i0_19++) {
          _nw124[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _783_v55 = _nw124;
        let _786_v56;
        _786_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f35,_783_v55);
        let _787_v57;
        _787_v57 = _dafny.MultiSet.fromElements(_783_v55, _783_v55, _783_v55, (((_786_v56).contains(p0)) ? ((_786_v56).get(p0)) : (_783_v55)));
        _787_v57 = _dafny.MultiSet.fromElements(_783_v55);
        (globalState).f21 = _module.__default.safeDivisionInt((_this).f36, new BigNumber(871));
      }
      let _788_v58;
      _788_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f36,(_this).f36);
      _788_v58 = (_788_v58).update(((_this).f36).minus((_this).f36), (_this).f36);
      let _789_v59;
      _789_v59 = _dafny.Map.Empty.slice().updateUnsafe(_687_v3,p0);
      let _790_v60;
      _790_v60 = _dafny.Set.fromElements(_789_v59, _789_v59);
      let _791_v61;
      _791_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f35,(_790_v60).Union(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_687_v3,false), _789_v59)));
      _791_v61 = (_791_v61).update((_this).f29, _dafny.Set.fromElements((_789_v59).update(new _dafny.CodePoint('u'.codePointAt(0)), (_this).f35), _789_v59));
      let _792_v62;
      _792_v62 = _dafny.MultiSet.fromElements((_this).f35);
      let _793_v63;
      _793_v63 = _module.D0.create_DC2(_792_v62, _dafny.Seq.UnicodeFromString("ngrrpfx"), (_this).f29);
      let _794_v64;
      let _nw125 = Array((new BigNumber(16)).toNumber());
      _nw125[(_dafny.ZERO).toNumber()] = ((false) ? ((_793_v63).dtor_cf4) : (_792_v62));
      _nw125[(_dafny.ONE).toNumber()] = ((_792_v62).update((_this).f29, _module.__default.abs((_dafny.ZERO).minus((_this).f36)))).Union(_module.__default.fm39(_687_v3, (_this).f36, globalState));
      _nw125[(new BigNumber(2)).toNumber()] = (_792_v62).Difference(_792_v62);
      _nw125[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(p0);
      _nw125[(new BigNumber(4)).toNumber()] = (_792_v62).update(!((_this).f29), _module.__default.abs((_this).f36));
      _nw125[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(p0);
      _nw125[(new BigNumber(6)).toNumber()] = _792_v62;
      _nw125[(new BigNumber(7)).toNumber()] = _792_v62;
      _nw125[(new BigNumber(8)).toNumber()] = (_793_v63).dtor_cf4;
      _nw125[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(p0, true, p0, (_this).f29, p0);
      _nw125[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.FromArray(_693_v5);
      _nw125[(new BigNumber(11)).toNumber()] = _792_v62;
      _nw125[(new BigNumber(12)).toNumber()] = (_793_v63).dtor_cf4;
      _nw125[(new BigNumber(13)).toNumber()] = (_module.__default.fm39(_687_v3, (_this).f36, globalState)).Intersect((_793_v63).dtor_cf4);
      _nw125[(new BigNumber(14)).toNumber()] = _792_v62;
      _nw125[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.fromElements(_module.__default.fm19(new _dafny.CodePoint('k'.codePointAt(0)), globalState));
      _794_v64 = _nw125;
      _794_v64 = _794_v64;
      let _795_v65;
      let _nw126 = new _module.C1();
      _nw126.__ctor((_this).f35);
      _795_v65 = _nw126;
      let _796_v66;
      _796_v66 = _dafny.Seq.of(_693_v5, _693_v5);
      r0 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_796_v66, _796_v66), _module.__default.safeIndex((_this).f36, new BigNumber((_dafny.Seq.Concat(_796_v66, _796_v66)).length)), _693_v5), _dafny.Seq.Concat(_796_v66, _796_v66));
      r1 = false;
      return [r0, r1];
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let _797_v0;
      _797_v0 = _dafny.Seq.of((_this).f29, (_this).f29);
      if (!_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.Concat(_797_v0, _module.__default.fm35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(531)), function (_798_i0) {
        return _module.D2.create_DC8(new _dafny.CodePoint('o'.codePointAt(0)), (_this).f35);
      }), (_dafny.ZERO).minus((_this).f36), globalState)), _module.__default.safeIndex((_this).f36, new BigNumber((_dafny.Seq.Concat(_797_v0, _module.__default.fm35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(531)), function (_799_i0) {
        return _module.D2.create_DC8(new _dafny.CodePoint('o'.codePointAt(0)), (_this).f35);
      }), (_dafny.ZERO).minus((_this).f36), globalState))).length)), p0), (_this).f35)) {
        let _800_v1;
        let _nw127 = new _module.C1();
        _nw127.__ctor((_this).f29);
        _800_v1 = _nw127;
        let _801_v2;
        _801_v2 = _module.D10.create_DC30(_800_v1);
        (globalState).f21 = new BigNumber((_dafny.Seq.of(_801_v2)).length);
        let _802_v3;
        _802_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,new BigNumber((_this.f30).length));
        let _803_v4;
        _803_v4 = _dafny.Seq.of((_dafny.ZERO).minus(p2), (_this).f36);
        (globalState).f1 = (_module.__default.fm9(globalState)).minus((((_802_v3).contains((_this).f35)) ? ((_802_v3).get((_this).f35)) : ((_803_v4)[_module.__default.safeIndex(p2, new BigNumber((_803_v4).length))])));
        let _804_v5;
        let _805_v6;
        let _out12;
        let _out13;
        let _outcollector6 = (_800_v1).m2(p0, globalState);
        _out12 = _outcollector6[0];
        _out13 = _outcollector6[1];
        _804_v5 = _out12;
        _805_v6 = _out13;
        let _806_v7;
        let _nw128 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _806_v7 = _nw128;
        let _807_v8;
        _807_v8 = _dafny.Set.fromElements(_806_v7);
        let _808_v9;
        let _nw129 = new _module.C2();
        _nw129.__ctor(p2, _807_v8, p0, _module.__default.fm20((_this).f29, !((_this).f35), _797_v0, p2, globalState));
        _808_v9 = _nw129;
        let _809_v10;
        let _nw130 = new _module.C1();
        _nw130.__ctor(true);
        _809_v10 = _nw130;
      } else {
        let _810_v11;
        let _nw131 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
        _810_v11 = _nw131;
        let _811_v12;
        _811_v12 = new _dafny.CodePoint('n'.codePointAt(0));
        let _812_v13;
        _812_v13 = _dafny.Seq.of(_dafny.Seq.update(_this.f30, _module.__default.safeIndex(new BigNumber((_this.f30).length), new BigNumber((_this.f30).length)), _811_v12));
        let _index122 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_810_v11).length));
        (_810_v11)[_index122] = _dafny.Seq.update(_812_v13, _module.__default.safeIndex(p2, new BigNumber((_812_v13).length)), _this.f30);
        let _813_v14;
        let _nw132 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _813_v14 = _nw132;
        let _index123 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_813_v14).length));
        (_813_v14)[_index123] = new BigNumber(62);
        let _814_v15;
        _814_v15 = _dafny.MultiSet.fromElements(p0);
        let _815_v16;
        _815_v16 = _dafny.Seq.of(_814_v15);
        let _816_v17;
        let _nw133 = Array((new BigNumber(7)).toNumber());
        _nw133[(_dafny.ZERO).toNumber()] = (_this).f29;
        _nw133[(_dafny.ONE).toNumber()] = (_this).f29;
        _nw133[(new BigNumber(2)).toNumber()] = ((_815_v16)[_module.__default.safeIndex(new BigNumber(-495), new BigNumber((_815_v16).length))]).IsProperSubsetOf((_814_v15).update(true, _module.__default.abs(p2)));
        _nw133[(new BigNumber(3)).toNumber()] = (_this).f35;
        _nw133[(new BigNumber(4)).toNumber()] = p0;
        _nw133[(new BigNumber(5)).toNumber()] = (_this).f35;
        _nw133[(new BigNumber(6)).toNumber()] = p0;
        _816_v17 = _nw133;
        let _817_v18;
        _817_v18 = _dafny.Set.fromElements(p0, (_this).f29);
        let _index124 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_810_v11).length));
        let _index125 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_813_v14).length));
        let _rhs132 = _dafny.Seq.Concat(_dafny.Seq.update(_812_v13, _module.__default.safeIndex((_this).f36, new BigNumber((_812_v13).length)), _this.f30), _dafny.Seq.Create(_module.__default.abs(new BigNumber(734)), function (_818_i1) {
          return _this.f30;
        }));
        let _rhs133 = (new BigNumber((_dafny.Seq.Concat(_this.f30, _dafny.Seq.UnicodeFromString("wjkkiurju"))).length)).minus((((_this).f35) ? ((_this).f36) : (new BigNumber(970))));
        let _rhs134 = new BigNumber((((false) ? (_817_v18) : (((p0) ? (_817_v18) : (_dafny.Set.fromElements(p0, (_this).f35)))))).length);
        let _rhs135 = _816_v17;
        let _lhs120 = _810_v11;
        let _lhs121 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_810_v11).length));
        let _lhs122 = _813_v14;
        let _lhs123 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_813_v14).length));
        let _lhs124 = globalState;
        _lhs120[_lhs121] = _rhs132;
        _lhs122[_lhs123] = _rhs133;
        _lhs124.f14 = _rhs134;
        _816_v17 = _rhs135;
        let _819_v19;
        _819_v19 = _dafny.MultiSet.fromElements((_813_v14)[_module.__default.safeIndex(new BigNumber(224), new BigNumber((_813_v14).length))]);
        (globalState).f13 = (_819_v19).IsProperSubsetOf(_819_v19);
        (globalState).f15 = !((_this).f36).isEqualTo(new BigNumber(-139));
        let _820_v20;
        let _nw134 = Array((new BigNumber(4)).toNumber());
        _nw134[(_dafny.ZERO).toNumber()] = _811_v12;
        _nw134[(_dafny.ONE).toNumber()] = _811_v12;
        _nw134[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('r'.codePointAt(0));
        _nw134[(new BigNumber(3)).toNumber()] = _811_v12;
        _820_v20 = _nw134;
        let _821_v21;
        _821_v21 = _dafny.Seq.of(_820_v20, _820_v20);
        let _822_v22;
        _822_v22 = _dafny.Seq.of(_820_v20, _820_v20, _820_v20, (_821_v21)[_module.__default.safeIndex(p2, new BigNumber((_821_v21).length))], _820_v20);
        _820_v20 = (_822_v22)[_module.__default.safeIndex((p2).multipliedBy((_dafny.ZERO).minus(new BigNumber((_797_v0).length))), new BigNumber((_822_v22).length))];
        let _rhs136 = _813_v14;
        let _rhs137 = _module.__default.fm33(((_this).f35) && ((_this).f29), (_813_v14)[_module.__default.safeIndex(new BigNumber(224), new BigNumber((_813_v14).length))], p2, globalState);
        let _lhs125 = globalState;
        let _lhs126 = globalState;
        _lhs125.f17 = _rhs136;
        _lhs126.f13 = _rhs137;
      }
      let _823_v23;
      _823_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f36,(_this).f36);
      let _824_v24;
      _824_v24 = new _dafny.CodePoint('h'.codePointAt(0));
      let _825_v25;
      _825_v25 = _dafny.Map.Empty.slice().updateUnsafe(_824_v24,(_this).f35);
      let _826_v26;
      _826_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f35,new BigNumber((_825_v25).length));
      let _827_v27;
      _827_v27 = _module.D4.create_DC15((((_823_v23).contains(new BigNumber(-597))) ? ((_823_v23).get(new BigNumber(-597))) : (new BigNumber((_826_v26).length))), (_this).f36);
      _827_v27 = function (_pat_let16_0) {
        return function (_828_dt__update__tmp_h0) {
          return function (_pat_let17_0) {
            return function (_829_dt__update_hcf25_h0) {
              return _module.D4.create_DC15((_828_dt__update__tmp_h0).dtor_cf24, _829_dt__update_hcf25_h0);
            }(_pat_let17_0);
          }((_this).f36);
        }(_pat_let16_0);
      }(_module.D4.create_DC15((_this).f36, new BigNumber((_this.f30).length)));
      if ((_this).f29) {
        let _830_v28;
        _830_v28 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length), (_this).f36);
        (globalState).f21 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_830_v28, _830_v28), _dafny.Seq.Create(_module.__default.abs(new BigNumber(875)), ((_831_p2) => function (_832_i2) {
          return _831_p2;
        })(p2)))).length);
        let _833_v29;
        let _nw135 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
        _833_v29 = _nw135;
        let _834_v30;
        _834_v30 = _dafny.Map.Empty.slice().updateUnsafe(p2,_833_v29);
        let _835_v31;
        _835_v31 = _module.D10.create_DC33(p0, (_this).f36);
        _834_v30 = (_834_v30).update((_835_v31).dtor_cf55, _833_v29);
        (globalState).f1 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_this.f30, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f36), new BigNumber((_this.f30).length)), _824_v24)).length), _module.__default.fm9(globalState)));
        r0 = _this.f30;
        let _836_v32;
        _836_v32 = _module.D6.create_DC22((_this).f36, (_this).f36, (_this).f29);
        let _837_v33;
        _837_v33 = _module.D9.create_DC29((_this).f36, (_836_v32).dtor_cf34, (_this).f29);
        let _838_v34;
        _838_v34 = _module.D9.create_DC27(_dafny.Seq.of((_837_v33).dtor_cf47, p2, (_this).f36, (_this).f36));
        let _839_v35;
        _839_v35 = _dafny.Set.fromElements(_838_v34, _838_v34, _838_v34);
        (globalState).f14 = new BigNumber((function () {
          let _coll37 = new _dafny.Set();
          for (const _compr_37 of ((_839_v35).Union(_839_v35)).Elements) {
            let _840_v36 = _compr_37;
            if (((_839_v35).Union(_839_v35)).contains(_840_v36)) {
              _coll37.add(_840_v36);
            }
          }
          return _coll37;
        }()).length);
      } else {
        (globalState).f18 = true;
        let _841_v37;
        _841_v37 = _dafny.MultiSet.fromElements((_this).f35);
        let _842_v38;
        _842_v38 = _module.D0.create_DC2((_841_v37).update((_this).f35, _module.__default.abs(p2)), _this.f30, (_this).f29);
        let _843_v39;
        let _nw136 = Array((new BigNumber(23)).toNumber());
        _nw136[(_dafny.ZERO).toNumber()] = p0;
        _nw136[(_dafny.ONE).toNumber()] = (_this).f35;
        _nw136[(new BigNumber(2)).toNumber()] = (_this).f35;
        _nw136[(new BigNumber(3)).toNumber()] = p0;
        _nw136[(new BigNumber(4)).toNumber()] = (_this).f35;
        _nw136[(new BigNumber(5)).toNumber()] = (_this).f35;
        _nw136[(new BigNumber(6)).toNumber()] = p0;
        _nw136[(new BigNumber(7)).toNumber()] = p0;
        _nw136[(new BigNumber(8)).toNumber()] = (_this).f29;
        _nw136[(new BigNumber(9)).toNumber()] = (_this).f29;
        _nw136[(new BigNumber(10)).toNumber()] = (_this).f29;
        _nw136[(new BigNumber(11)).toNumber()] = (_this).f35;
        _nw136[(new BigNumber(12)).toNumber()] = (_this).f29;
        _nw136[(new BigNumber(13)).toNumber()] = (_this).f35;
        _nw136[(new BigNumber(14)).toNumber()] = (_this).f29;
        _nw136[(new BigNumber(15)).toNumber()] = p0;
        _nw136[(new BigNumber(16)).toNumber()] = (_this).f29;
        _nw136[(new BigNumber(17)).toNumber()] = p0;
        _nw136[(new BigNumber(18)).toNumber()] = p0;
        _nw136[(new BigNumber(19)).toNumber()] = (_this).f29;
        _nw136[(new BigNumber(20)).toNumber()] = (_this).f29;
        _nw136[(new BigNumber(21)).toNumber()] = p0;
        _nw136[(new BigNumber(22)).toNumber()] = (_this).f35;
        _843_v39 = _nw136;
        let _844_v40;
        _844_v40 = _dafny.Seq.of(_843_v39);
        let _845_v41;
        _845_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f35);
        let _846_v42;
        let _nw137 = Array((new BigNumber(15)).toNumber());
        _nw137[(_dafny.ZERO).toNumber()] = _module.D0.create_DC2(_841_v37, _this.f30, (_this).f29);
        _nw137[(_dafny.ONE).toNumber()] = (((_this).f29) ? (_842_v38) : (_842_v38));
        _nw137[(new BigNumber(2)).toNumber()] = _842_v38;
        _nw137[(new BigNumber(3)).toNumber()] = _842_v38;
        _nw137[(new BigNumber(4)).toNumber()] = _module.__default.fm40(p0, globalState);
        _nw137[(new BigNumber(5)).toNumber()] = _module.D0.create_DC2(_dafny.MultiSet.FromArray(_797_v0), (_module.D2.create_DC9(_this.f30, new BigNumber((_844_v40).length), (_this).f29, p2)).dtor_cf14, (_this).f35);
        _nw137[(new BigNumber(6)).toNumber()] = _842_v38;
        _nw137[(new BigNumber(7)).toNumber()] = _module.D0.create_DC2(_module.__default.fm39(_824_v24, p2, globalState), _this.f30, (((_845_v41).contains((_this).f29)) ? ((_845_v41).get((_this).f29)) : (!(p0))));
        _nw137[(new BigNumber(8)).toNumber()] = _842_v38;
        _nw137[(new BigNumber(9)).toNumber()] = _module.D0.create_DC2(_dafny.MultiSet.fromElements((_this).f35), _this.f30, p0);
        _nw137[(new BigNumber(10)).toNumber()] = _module.D0.create_DC2(_841_v37, _this.f30, (_this).f29);
        _nw137[(new BigNumber(11)).toNumber()] = _842_v38;
        _nw137[(new BigNumber(12)).toNumber()] = _module.D0.create_DC2(_841_v37, _this.f30, true);
        _nw137[(new BigNumber(13)).toNumber()] = _842_v38;
        _nw137[(new BigNumber(14)).toNumber()] = _842_v38;
        _846_v42 = _nw137;
        let _index126 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_846_v42).length));
        (_846_v42)[_index126] = _module.D0.create_DC2(_841_v37, _dafny.Seq.Create(_module.__default.abs(new BigNumber(926)), ((_847_v24) => function (_848_i3) {
  return _847_v24;
})(_824_v24)), p0);
        let _index127 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_846_v42).length));
        (_846_v42)[_index127] = _842_v38;
        let _index128 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_843_v39).length));
        (_843_v39)[_index128] = (_this).f29;
        let _849_v43;
        _849_v43 = _dafny.MultiSet.fromElements(_843_v39);
        let _850_v44;
        _850_v44 = _dafny.Seq.of(_849_v43, _849_v43, (_849_v43).update(_843_v39, _module.__default.abs(_module.__default.fm36(globalState))));
        let _851_v45;
        _851_v45 = _dafny.Seq.of((_this).f36, p2);
        let _index129 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_843_v39).length));
        (_843_v39)[_index129] = ((_850_v44)[_module.__default.safeIndex(new BigNumber((_851_v45).length), new BigNumber((_850_v44).length))]).IsProperSubsetOf(_dafny.MultiSet.fromElements(_843_v39, _843_v39, _843_v39));
        (globalState).f18 = !((_843_v39)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_843_v39).length))]);
        (globalState).f21 = _module.__default.fm9(globalState);
      }
      let _852_v46;
      _852_v46 = _dafny.Set.fromElements(_824_v24, _824_v24, _824_v24);
      _852_v46 = _852_v46;
      let _853_v47;
      let _nw138 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
      _853_v47 = _nw138;
      let _854_v48;
      _854_v48 = _dafny.Set.fromElements(_853_v47);
      let _855_v49;
      let _nw139 = new _module.C2();
      _nw139.__ctor(new BigNumber((_dafny.Seq.Concat(_this.f30, _dafny.Seq.UnicodeFromString("ivbrw"))).length), _854_v48, (_this).f35, _this.f30);
      _855_v49 = _nw139;
      let _856_v50;
      _856_v50 = _dafny.Set.fromElements((_this).f35);
      let _857_v51;
      _857_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_this.f30).length),(_this).f29);
      let _hi2 = new BigNumber((_857_v51).length);
      for (let _858_i4 = new BigNumber(((_856_v50).Difference(_856_v50)).length); _858_i4.isLessThan(_hi2); _858_i4 = _858_i4.plus(_dafny.ONE)) {
        (globalState).f13 = !(p0) || (p0);
        (globalState).f13 = _module.__default.fm19(new _dafny.CodePoint('h'.codePointAt(0)), globalState);
        let _859_v52;
        let _nw140 = new _module.C0();
        _nw140.__ctor((_this).f35);
        _859_v52 = _nw140;
        let _860_v53;
        _860_v53 = _dafny.MultiSet.fromElements(new BigNumber(162), new BigNumber((_module.__default.fm39(_824_v24, new BigNumber(380), globalState)).cardinality()));
        let _861_v55;
        _861_v55 = _dafny.Seq.of((_855_v49).f38);
        let _862_v56;
        _862_v56 = _dafny.Seq.of(_861_v55, _861_v55);
        let _863_v57;
        let _nw141 = Array((new BigNumber(27)).toNumber());
        _nw141[(_dafny.ZERO).toNumber()] = (_this).f29;
        _nw141[(_dafny.ONE).toNumber()] = p0;
        _nw141[(new BigNumber(2)).toNumber()] = (_this).f35;
        _nw141[(new BigNumber(3)).toNumber()] = _module.__default.fm33(true, (_this).f36, (((_860_v53).contains(new BigNumber(((function () {
          let _coll39 = new _dafny.Map();
          for (const _compr_39 of _dafny.IntegerRange(new BigNumber(777), new BigNumber(-289))) {
            let _865_v54 = _compr_39;
            if (((new BigNumber(777)).isLessThanOrEqualTo(_865_v54)) && ((_865_v54).isLessThan(new BigNumber(-289)))) {
              _coll39.push([(_865_v54).multipliedBy(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality())),(_855_v49).f38]);
            }
          }
          return _coll39;
        }()).update((_this).f36, (_this).f36)).length))) ? ((_860_v53).get(new BigNumber(((function () {
          let _coll38 = new _dafny.Map();
          for (const _compr_38 of _dafny.IntegerRange(new BigNumber(777), new BigNumber(-289))) {
            let _864_v54 = _compr_38;
            if (((new BigNumber(777)).isLessThanOrEqualTo(_864_v54)) && ((_864_v54).isLessThan(new BigNumber(-289)))) {
              _coll38.push([(_864_v54).multipliedBy(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality())),(_855_v49).f38]);
            }
          }
          return _coll38;
        }()).update((_this).f36, (_this).f36)).length))) : ((_855_v49).f38)), globalState);
        _nw141[(new BigNumber(4)).toNumber()] = _859_v52.f37;
        _nw141[(new BigNumber(5)).toNumber()] = (_this).f35;
        _nw141[(new BigNumber(6)).toNumber()] = _859_v52.f37;
        _nw141[(new BigNumber(7)).toNumber()] = (_this).f29;
        _nw141[(new BigNumber(8)).toNumber()] = (_859_v52).fm11(true, globalState);
        _nw141[(new BigNumber(9)).toNumber()] = (new BigNumber((_module.__default.fm41(new BigNumber((_862_v56).length), (_this).f29, (_this).f29, _859_v52.f37, globalState)).length)).isLessThanOrEqualTo((_this).f36);
        _nw141[(new BigNumber(10)).toNumber()] = p0;
        _nw141[(new BigNumber(11)).toNumber()] = true;
        _nw141[(new BigNumber(12)).toNumber()] = _dafny.areEqual(_797_v0, _797_v0);
        _nw141[(new BigNumber(13)).toNumber()] = (_this).f29;
        _nw141[(new BigNumber(14)).toNumber()] = _859_v52.f37;
        _nw141[(new BigNumber(15)).toNumber()] = p0;
        _nw141[(new BigNumber(16)).toNumber()] = (_this).f35;
        _nw141[(new BigNumber(17)).toNumber()] = (_this).f29;
        _nw141[(new BigNumber(18)).toNumber()] = (new BigNumber(777)).isLessThan((((_823_v23).contains(new BigNumber(793))) ? ((_823_v23).get(new BigNumber(793))) : ((_this).f36)));
        _nw141[(new BigNumber(19)).toNumber()] = ((_855_v49).f38).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_859_v52,(_this).f29)).length));
        _nw141[(new BigNumber(20)).toNumber()] = p0;
        _nw141[(new BigNumber(21)).toNumber()] = (_this).f35;
        _nw141[(new BigNumber(22)).toNumber()] = (_797_v0)[_module.__default.safeIndex((_855_v49).f38, new BigNumber((_797_v0).length))];
        _nw141[(new BigNumber(23)).toNumber()] = (_this).f29;
        _nw141[(new BigNumber(24)).toNumber()] = (_this).f29;
        _nw141[(new BigNumber(25)).toNumber()] = !((_this).f35);
        _nw141[(new BigNumber(26)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_824_v24, _module.__default.fm42(_859_v52.f37, _859_v52.f37, globalState)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(688)), ((_866_v24) => function (_867_i5) {
          return _866_v24;
        })(_824_v24)));
        _863_v57 = _nw141;
        let _index130 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_863_v57).length));
        (_863_v57)[_index130] = (_this).f35;
        let _868_v58;
        let _nw142 = Array((new BigNumber(29)).toNumber()).fill(_module.D1.Default());
        _868_v58 = _nw142;
        let _869_v59;
        _869_v59 = _module.D1.create_DC3(_797_v0);
        let _870_v60;
        _870_v60 = _module.D1.create_DC3((_869_v59).dtor_cf7);
        let _871_v61;
        _871_v61 = _module.D1.create_DC5(_870_v60);
        let _index131 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_868_v58).length));
        (_868_v58)[_index131] = _871_v61;
        let _872_v62;
        _872_v62 = _module.D10.create_DC33(false, p2);
        let _index132 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_863_v57).length));
        let _index133 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_868_v58).length));
        let _rhs138 = (_859_v52).fm11((_872_v62).dtor_cf54, globalState);
        let _rhs139 = _871_v61;
        let _lhs127 = _863_v57;
        let _lhs128 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_863_v57).length));
        let _lhs129 = _868_v58;
        let _lhs130 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_868_v58).length));
        _lhs127[_lhs128] = _rhs138;
        _lhs129[_lhs130] = _rhs139;
      }
      let _873_v63;
      _873_v63 = _dafny.Seq.of(_this.f30);
      r0 = (_873_v63)[_module.__default.safeIndex(new BigNumber((_this.f30).length), new BigNumber((_873_v63).length))];
      let _874_v64;
      _874_v64 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p2), (_this).f36);
      r1 = (p2).isLessThan(_module.__default.safeDivisionInt((_855_v49).f38, new BigNumber((_874_v64).cardinality())));
      return [r0, r1];
    }
    get f35() {
      let _this = this;
      return _this._f35;
    };
    get f36() {
      let _this = this;
      return _this._f36;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f33 = false;
      this._f34 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f33, f34) {
      let _this = this;
      (_this)._f33 = f33;
      (_this)._f34 = f34;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_this).f34, (_dafny.ZERO).minus((_this).f34));
    };
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _875_i0;
      _875_i0 = _dafny.ZERO;
      L3: {
        while (_module.__default.fm6(((_this).f34).multipliedBy((_this).f34), globalState)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_875_i0)) {
              break L3;
            }
            _875_i0 = (_875_i0).plus(_dafny.ONE);
            let _876_v0;
            _876_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,(_this).f33);
            _876_v0 = (_876_v0).update((_this).f34, (_this).f33);
            if ((_this).f33) {
              (globalState).f28 = p0;
              let _877_v1;
              _877_v1 = _dafny.Set.fromElements(p0);
              (globalState).f2 = (_this).fm5((_877_v1).Union(_dafny.Set.fromElements(_module.__default.fm6((_this).f34, globalState))), globalState);
              (globalState).f18 = !((_this).f33);
              let _878_v2;
              let _nw143 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
              _878_v2 = _nw143;
              (globalState).f17 = _878_v2;
              (globalState).f21 = (((_this).f34).plus((_this).f34)).multipliedBy((_this).f34);
            } else {
              let _879_v3;
              _879_v3 = _dafny.Seq.UnicodeFromString("yk");
              (globalState).f7 = _dafny.Seq.of(((_this).f34).multipliedBy(new BigNumber((_879_v3).length)));
              (globalState).f28 = !(p0) || (p0);
              let _880_v4;
              let _nw144 = Array((new BigNumber(7)).toNumber()).fill([]);
              _880_v4 = _nw144;
              let _881_v5;
              let _init20 = ((_882_p0) => function (_883_i1) {
                return _dafny.Seq.of((_this).f33, _882_p0);
              })(p0);
              let _nw145 = Array((new BigNumber(12)).toNumber());
              for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw145.length); _i0_20++) {
                _nw145[_i0_20] = _init20(new BigNumber(_i0_20));
              }
              _881_v5 = _nw145;
              let _884_v6;
              _884_v6 = _dafny.Seq.of(_881_v5);
              let _index134 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_880_v4).length));
              (_880_v4)[_index134] = (_884_v6)[_module.__default.safeIndex((_this).f34, new BigNumber((_884_v6).length))];
              let _index135 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_880_v4).length));
              (_880_v4)[_index135] = _881_v5;
              let _885_v7;
              let _nw146 = Array((new BigNumber(11)).toNumber()).fill(false);
              _885_v7 = _nw146;
              _885_v7 = _885_v7;
              (globalState).f15 = p0;
            }
            let _886_v8;
            _886_v8 = _dafny.Seq.UnicodeFromString("ue");
            let _887_v9;
            _887_v9 = _dafny.Seq.of(_886_v8);
            (globalState).f2 = new BigNumber((_887_v9).length);
            let _888_v10;
            _888_v10 = _dafny.Seq.of((_this).f33, false, (_this).f33, false, p0);
            (globalState).f15 = ((false) ? (!_dafny.Seq.contains(_888_v10, (_this).f33)) : ((_this).f33));
          }
        }
      }
      let _hi3 = (_this).f34;
      for (let _889_i2 = new BigNumber((_dafny.Seq.UnicodeFromString("fwevwbc")).length); _889_i2.isLessThan(_hi3); _889_i2 = _889_i2.plus(_dafny.ONE)) {
        (globalState).f15 = !(_module.__default.fm6(_889_i2, globalState));
        let _890_v11;
        _890_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(748),(_this).f34);
        _890_v11 = (_890_v11).update(new BigNumber(-437), (_this).f34);
        let _891_v12;
        let _init21 = function (_892_i3) {
          return !(new BigNumber(352)).isEqualTo((_dafny.ZERO).minus((_this).f34));
        };
        let _nw147 = Array((new BigNumber(28)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw147.length); _i0_21++) {
          _nw147[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _891_v12 = _nw147;
        _891_v12 = _891_v12;
        let _893_v13;
        _893_v13 = _dafny.Seq.of(_891_v12, _891_v12);
        let _894_v14;
        let _nw148 = Array((new BigNumber(17)).toNumber()).fill([]);
        _894_v14 = _nw148;
        let _895_v15;
        let _init22 = function (_896_i4) {
          return (_896_i4).minus(new BigNumber(537));
        };
        let _nw149 = Array((new BigNumber(13)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw149.length); _i0_22++) {
          _nw149[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _895_v15 = _nw149;
        let _index136 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_894_v14).length));
        (_894_v14)[_index136] = _895_v15;
        let _897_v16;
        _897_v16 = _dafny.Seq.UnicodeFromString("hvp");
        let _898_v17;
        _898_v17 = _dafny.MultiSet.fromElements(p0);
        let _899_v18;
        _899_v18 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("jkowf"), _module.__default.fm7((_this).f34, globalState));
        let _900_v19;
        _900_v19 = new _dafny.CodePoint('j'.codePointAt(0));
        let _901_v20;
        _901_v20 = _dafny.Map.Empty.slice().updateUnsafe(_889_i2,_900_v19);
        let _pat_let_tv20 = globalState;
        let _902_v21;
        _902_v21 = _module.D2.create_DC9(_897_v16, new BigNumber(((_890_v11).update(new BigNumber((_897_v16).length), _889_i2)).length), (function (_pat_let18_0) {
  return function (_903_dt__update__tmp_h0) {
    return function (_pat_let19_0) {
      return function (_904_dt__update_hcf5_h0) {
        return function (_pat_let20_0) {
          return function (_905_dt__update_hcf6_h0) {
            return _module.D0.create_DC2((_903_dt__update__tmp_h0).dtor_cf4, _904_dt__update_hcf5_h0, _905_dt__update_hcf6_h0);
          }(_pat_let20_0);
        }(true);
      }(_pat_let19_0);
    }(_module.__default.fm7(_889_i2, _pat_let_tv20));
  }(_pat_let18_0);
}(_module.D0.create_DC2(_898_v17, (_899_v18)[_module.__default.safeIndex(new BigNumber(((_901_v20).update(_889_i2, _900_v19)).length), new BigNumber((_899_v18).length))], p0))).dtor_cf6, (_this).f34);
        let _index137 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_894_v14).length));
        let _rhs140 = _dafny.Seq.of(((p0) ? (_891_v12) : (_891_v12)), _891_v12, _891_v12, _891_v12);
        let _rhs141 = _895_v15;
        let _rhs142 = (_902_v21).dtor_cf17;
        let _lhs131 = _894_v14;
        let _lhs132 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_894_v14).length));
        let _lhs133 = globalState;
        _893_v13 = _rhs140;
        _lhs131[_lhs132] = _rhs141;
        _lhs133.f2 = _rhs142;
      }
      let _906_v22;
      let _init23 = function (_907_i5) {
        return _module.__default.safeDivisionInt(_907_i5, (_this).f34);
      };
      let _nw150 = Array((new BigNumber(9)).toNumber());
      for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw150.length); _i0_23++) {
        _nw150[_i0_23] = _init23(new BigNumber(_i0_23));
      }
      _906_v22 = _nw150;
      let _index138 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_906_v22).length));
      (_906_v22)[_index138] = (_this).f34;
      let _908_v23;
      _908_v23 = _dafny.Seq.of(true);
      let _index139 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_906_v22).length));
      let _rhs143 = (_this).f34;
      let _rhs144 = _dafny.Seq.Concat(((p0) ? (_dafny.Seq.of((_this).f33)) : (_dafny.Seq.of((_this).f33))), _908_v23);
      let _lhs134 = _906_v22;
      let _lhs135 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_906_v22).length));
      _lhs134[_lhs135] = _rhs143;
      _908_v23 = _rhs144;
      let _909_v24;
      _909_v24 = _dafny.MultiSet.fromElements((_906_v22)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_906_v22).length))], (_906_v22)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_906_v22).length))], (_this).f34);
      let _910_v25;
      _910_v25 = _module.D3.create_DC10(_909_v24);
      (globalState).f28 = (_909_v24).equals((_910_v25).dtor_cf18);
      let _index140 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_906_v22).length));
      (_906_v22)[_index140] = (_this).f34;
      let _init24 = function (_911_i6) {
        return _module.__default.safeDivisionInt(_911_i6, new BigNumber((_dafny.Seq.UnicodeFromString("bmap")).length));
      };
      let _nw151 = Array((new BigNumber(2)).toNumber());
      for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw151.length); _i0_24++) {
        _nw151[_i0_24] = _init24(new BigNumber(_i0_24));
      }
      _906_v22 = _nw151;
      let _912_v26;
      _912_v26 = _dafny.Seq.UnicodeFromString("hoalen");
      let _913_v27;
      _913_v27 = _dafny.Seq.of(_908_v23, _dafny.Seq.Concat(_dafny.Seq.of(false, (_this).f33, p0, p0, _module.__default.fm6(new BigNumber((_912_v26).length), globalState)), _908_v23));
      r0 = _913_v27;
      r1 = !((_module.__default.safeDivisionInt((_906_v22)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_906_v22).length))], (_906_v22)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_906_v22).length))])).isLessThanOrEqualTo((_906_v22)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_906_v22).length))]));
      return [r0, r1];
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let _914_v0;
      let _nw152 = Array((new BigNumber(11)).toNumber()).fill(false);
      _914_v0 = _nw152;
      let _index141 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_914_v0).length));
      (_914_v0)[_index141] = (_this).f33;
      let _index142 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_914_v0).length));
      (_914_v0)[_index142] = !((_this).f33) || ((_this).f33);
      (globalState).f28 = (_this).f33;
      let _915_v1;
      _915_v1 = _dafny.Seq.UnicodeFromString("fckyvw");
      let _916_v2;
      _916_v2 = _module.D2.create_DC9(_915_v1, p2, p0, (_this).f34);
      r0 = _dafny.Seq.update(_915_v1, _module.__default.safeIndex((_916_v2).dtor_cf15, new BigNumber((_915_v1).length)), (_915_v1)[_module.__default.safeIndex(p2, new BigNumber((_915_v1).length))]);
      let _917_v3;
      _917_v3 = _dafny.Set.fromElements(false, (_914_v0)[_module.__default.safeIndex(new BigNumber(618), new BigNumber((_914_v0).length))]);
      _916_v2 = _module.__default.fm8(!(((_this).fm5(_917_v3, globalState)).isLessThan((_this).f34)), (_914_v0)[_module.__default.safeIndex(new BigNumber(618), new BigNumber((_914_v0).length))], globalState);
      let _918_v4;
      _918_v4 = new _dafny.CodePoint('u'.codePointAt(0));
      let _919_v5;
      let _nw153 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
      _919_v5 = _nw153;
      let _index143 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((_919_v5).length));
      (_919_v5)[_index143] = p2;
      let _920_v6;
      _920_v6 = _dafny.Set.fromElements(p2);
      let _921_v7;
      _921_v7 = _dafny.MultiSet.fromElements(p2, (_this).f34, ((p0) ? ((_this).f34) : ((_this).f34)), new BigNumber((_920_v6).length), new BigNumber(223));
      let _922_v8;
      _922_v8 = _dafny.Seq.of((_this).f33, p0, (_914_v0)[_module.__default.safeIndex(new BigNumber(618), new BigNumber((_914_v0).length))]);
      let _923_v9;
      _923_v9 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
      let _924_v10;
      _924_v10 = _dafny.MultiSet.fromElements(_922_v8);
      let _925_v11;
      _925_v11 = _dafny.Seq.of((_this).f34, new BigNumber((_924_v10).cardinality()), p2);
      let _926_v12;
      _926_v12 = _dafny.Map.Empty.slice().updateUnsafe((_914_v0)[_module.__default.safeIndex(new BigNumber(618), new BigNumber((_914_v0).length))],_dafny.Seq.of((_dafny.ZERO).minus(p2)));
      let _index144 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((_919_v5).length));
      let _rhs145 = (((_921_v7).contains((_this).fm5(_dafny.Set.fromElements((_922_v8)[_module.__default.safeIndex(p2, new BigNumber((_922_v8).length))]), globalState))) ? ((_921_v7).get((_this).fm5(_dafny.Set.fromElements((_922_v8)[_module.__default.safeIndex(p2, new BigNumber((_922_v8).length))]), globalState))) : (new BigNumber((_923_v9).length)));
      let _rhs146 = !((_914_v0)[_module.__default.safeIndex(new BigNumber(618), new BigNumber((_914_v0).length))]);
      let _rhs147 = _918_v4;
      let _rhs148 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_this).f34, (_this).f34, p2, new BigNumber((_922_v8).length), (_this).f34), _925_v11), (((_926_v12).contains((_914_v0)[_module.__default.safeIndex(new BigNumber(618), new BigNumber((_914_v0).length))])) ? ((_926_v12).get((_914_v0)[_module.__default.safeIndex(new BigNumber(618), new BigNumber((_914_v0).length))])) : (_925_v11)));
      let _rhs149 = p2;
      let _lhs136 = globalState;
      let _lhs137 = globalState;
      let _lhs138 = globalState;
      let _lhs139 = _919_v5;
      let _lhs140 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((_919_v5).length));
      _lhs136.f21 = _rhs145;
      _lhs137.f28 = _rhs146;
      _918_v4 = _rhs147;
      _lhs138.f7 = _rhs148;
      _lhs139[_lhs140] = _rhs149;
      let _index145 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((p1).length));
      (p1)[_index145] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-838)), ((_927_v4) => function (_928_i0) {
        return (_module.D2.create_DC8(_927_v4, (_this).f33)).dtor_cf12;
      })(_918_v4));
      let _index146 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((p1).length));
      (p1)[_index146] = _915_v1;
      r0 = _dafny.Seq.Concat(_915_v1, _dafny.Seq.UnicodeFromString("keatndrhj"));
      r1 = _module.__default.fm6(new BigNumber(82), globalState);
      return [r0, r1];
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _929_v0;
      let _nw154 = Array((new BigNumber(29)).toNumber());
      _nw154[(_dafny.ZERO).toNumber()] = p0;
      _nw154[(_dafny.ONE).toNumber()] = p0;
      _nw154[(new BigNumber(2)).toNumber()] = (_this).f33;
      _nw154[(new BigNumber(3)).toNumber()] = false;
      _nw154[(new BigNumber(4)).toNumber()] = p0;
      _nw154[(new BigNumber(5)).toNumber()] = p0;
      _nw154[(new BigNumber(6)).toNumber()] = !(p0);
      _nw154[(new BigNumber(7)).toNumber()] = (_this).f33;
      _nw154[(new BigNumber(8)).toNumber()] = !(true);
      _nw154[(new BigNumber(9)).toNumber()] = (_this).f33;
      _nw154[(new BigNumber(10)).toNumber()] = (_this).f33;
      _nw154[(new BigNumber(11)).toNumber()] = p0;
      _nw154[(new BigNumber(12)).toNumber()] = (_this).f33;
      _nw154[(new BigNumber(13)).toNumber()] = (_this).f33;
      _nw154[(new BigNumber(14)).toNumber()] = (_this).f33;
      _nw154[(new BigNumber(15)).toNumber()] = true;
      _nw154[(new BigNumber(16)).toNumber()] = (_this).f33;
      _nw154[(new BigNumber(17)).toNumber()] = p0;
      _nw154[(new BigNumber(18)).toNumber()] = (_this).f33;
      _nw154[(new BigNumber(19)).toNumber()] = p0;
      _nw154[(new BigNumber(20)).toNumber()] = (_this).f33;
      _nw154[(new BigNumber(21)).toNumber()] = p0;
      _nw154[(new BigNumber(22)).toNumber()] = (_this).f33;
      _nw154[(new BigNumber(23)).toNumber()] = p0;
      _nw154[(new BigNumber(24)).toNumber()] = true;
      _nw154[(new BigNumber(25)).toNumber()] = true;
      _nw154[(new BigNumber(26)).toNumber()] = p0;
      _nw154[(new BigNumber(27)).toNumber()] = (_this).f33;
      _nw154[(new BigNumber(28)).toNumber()] = p0;
      _929_v0 = _nw154;
      let _930_v1;
      _930_v1 = _dafny.Seq.of(_929_v0);
      let _931_v2;
      _931_v2 = _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(new BigNumber((_930_v1).length), (_this).f34));
      let _932_v3;
      _932_v3 = _dafny.Set.fromElements((_this).f33);
      let _933_v4;
      _933_v4 = _dafny.Seq.of(p0);
      let _rhs150 = (_this).fm5(_932_v3, globalState);
      let _rhs151 = (_933_v4)[_module.__default.safeIndex(new BigNumber(978), new BigNumber((_933_v4).length))];
      let _rhs152 = new BigNumber((p2).length);
      let _rhs153 = !((_this).f33);
      let _rhs154 = _931_v2;
      let _lhs141 = globalState;
      let _lhs142 = globalState;
      let _lhs143 = globalState;
      _lhs141.f2 = _rhs150;
      r0 = _rhs151;
      _lhs142.f21 = _rhs152;
      _lhs143.f13 = _rhs153;
      _931_v2 = _rhs154;
      let _934_v5;
      _934_v5 = _dafny.Seq.of(_dafny.Seq.update(_933_v4, _module.__default.safeIndex((_this).f34, new BigNumber((_933_v4).length)), (_this).f33), _933_v4);
      let _935_v6;
      _935_v6 = _dafny.Seq.of((_934_v5)[_module.__default.safeIndex((_this).f34, new BigNumber((_934_v5).length))]);
      (globalState).f2 = ((_module.__default.fm6((_this).f34, globalState)) ? ((new BigNumber(((_934_v5)[_module.__default.safeIndex((_this).f34, new BigNumber((_934_v5).length))]).length)).minus((_this).f34)) : ((_this).fm5(_932_v3, globalState)));
      let _936_i0;
      _936_i0 = _dafny.ZERO;
      L4: {
        while (!(p0)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_936_i0)) {
              break L4;
            }
            _936_i0 = (_936_i0).plus(_dafny.ONE);
            (globalState).f28 = ((_this).f34).isLessThanOrEqualTo((_this).f34);
            let _937_v7;
            let _nw155 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
            _937_v7 = _nw155;
            let _938_v8;
            _938_v8 = new _dafny.CodePoint('k'.codePointAt(0));
            let _939_v9;
            _939_v9 = _dafny.Map.Empty.slice().updateUnsafe(_938_v8,(_933_v4)[_module.__default.safeIndex((_this).f34, new BigNumber((_933_v4).length))]);
            let _index147 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_937_v7).length));
            (_937_v7)[_index147] = _939_v9;
            let _940_v11;
            _940_v11 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),(_dafny.ZERO).minus((_this).f34));
            let _index148 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_937_v7).length));
            let _rhs155 = function () {
              let _coll40 = new _dafny.Map();
              for (const _compr_40 of (_940_v11).Keys.Elements) {
                let _941_v10 = _compr_40;
                if ((_940_v11).contains(_941_v10)) {
                  _coll40.push([_941_v10,false]);
                }
              }
              return _coll40;
            }();
            let _rhs156 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f34, (_this).f34));
            let _rhs157 = (_this).f34;
            let _lhs144 = _937_v7;
            let _lhs145 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_937_v7).length));
            let _lhs146 = globalState;
            let _lhs147 = globalState;
            _lhs144[_lhs145] = _rhs155;
            _lhs146.f2 = _rhs156;
            _lhs147.f14 = _rhs157;
            (globalState).f2 = ((_this).f34).plus((_this).f34);
            let _942_v12;
            _942_v12 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_this).f34, (_this).f34),p0);
            _942_v12 = (_942_v12).update((_this).f34, p0);
          }
        }
      }
      r0 = (_this).f33;
      let _943_v13;
      _943_v13 = _dafny.Set.fromElements((_this).f34, (_this).f34);
      (globalState).f14 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_this).f34, new BigNumber((_943_v13).length)), (_this).f34);
      let _hi4 = (_this).f34;
      for (let _944_i1 = (_this).f34; _944_i1.isLessThan(_hi4); _944_i1 = _944_i1.plus(_dafny.ONE)) {
        let _945_v14;
        _945_v14 = _dafny.MultiSet.fromElements(_module.__default.fm6((_this).f34, globalState));
        let _946_v15;
        _946_v15 = _dafny.Seq.of(_944_i1);
        let _947_v16;
        _947_v16 = new _dafny.CodePoint('p'.codePointAt(0));
        let _948_v17;
        _948_v17 = _dafny.Map.Empty.slice().updateUnsafe(_944_i1,_947_v16);
        let _949_v18;
        let _nw156 = Array((new BigNumber(11)).toNumber());
        _nw156[(_dafny.ZERO).toNumber()] = _944_i1;
        _nw156[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of((_this).f34, (_this).f34, _944_i1, _944_i1, (_dafny.ZERO).minus((_dafny.ZERO).minus(_944_i1))), _module.__default.safeIndex((((_945_v14).contains(true)) ? ((_945_v14).get(true)) : (_944_i1)), new BigNumber((_dafny.Seq.of((_this).f34, (_this).f34, _944_i1, _944_i1, (_dafny.ZERO).minus((_dafny.ZERO).minus(_944_i1)))).length)), _944_i1), _946_v15)).length);
        _nw156[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(_944_i1, _944_i1);
        _nw156[(new BigNumber(3)).toNumber()] = _944_i1;
        _nw156[(new BigNumber(4)).toNumber()] = (_this).f34;
        _nw156[(new BigNumber(5)).toNumber()] = _944_i1;
        _nw156[(new BigNumber(6)).toNumber()] = new BigNumber(305);
        _nw156[(new BigNumber(7)).toNumber()] = _944_i1;
        _nw156[(new BigNumber(8)).toNumber()] = _944_i1;
        _nw156[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_948_v17).length));
        _nw156[(new BigNumber(10)).toNumber()] = _944_i1;
        _949_v18 = _nw156;
        (globalState).f17 = _949_v18;
        (globalState).f15 = !((_this).f33);
        let _950_v19;
        let _nw157 = new _module.C1();
        _nw157.__ctor(p0);
        _950_v19 = _nw157;
        if ((_950_v19).f40) {
          (globalState).f17 = _949_v18;
          let _index149 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_929_v0).length));
          (_929_v0)[_index149] = !((_this).f33);
          let _951_v20;
          _951_v20 = _dafny.Map.Empty.slice().updateUnsafe(_944_i1,(_this).f34);
          let _index150 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_929_v0).length));
          let _rhs158 = (_module.__default.fm6(new BigNumber((_951_v20).length), globalState)) === (true);
          let _rhs159 = (_932_v3).IsDisjointFrom(((!((_this).f33)) ? (_932_v3) : (_932_v3)));
          let _rhs160 = !((_this).f34).isEqualTo((_this).f34);
          let _lhs148 = globalState;
          let _lhs149 = _929_v0;
          let _lhs150 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_929_v0).length));
          let _lhs151 = globalState;
          _lhs148.f13 = _rhs158;
          _lhs149[_lhs150] = _rhs159;
          _lhs151.f13 = _rhs160;
          let _952_v21;
          let _nw158 = Array((new BigNumber(11)).toNumber()).fill([]);
          _952_v21 = _nw158;
          let _953_v22;
          _953_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(21),_module.__default.fm19(_947_v16, globalState));
          let _954_v23;
          _954_v23 = _dafny.Map.Empty.slice().updateUnsafe(!((_950_v19).f40),p0);
          let _955_v24;
          let _nw159 = Array((new BigNumber(28)).toNumber());
          _nw159[(_dafny.ZERO).toNumber()] = (_929_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_929_v0).length))];
          _nw159[(_dafny.ONE).toNumber()] = true;
          _nw159[(new BigNumber(2)).toNumber()] = (_933_v4)[_module.__default.safeIndex((_this).f34, new BigNumber((_933_v4).length))];
          _nw159[(new BigNumber(3)).toNumber()] = (p3).dtor_cf0;
          _nw159[(new BigNumber(4)).toNumber()] = true;
          _nw159[(new BigNumber(5)).toNumber()] = (_950_v19).f40;
          _nw159[(new BigNumber(6)).toNumber()] = (_950_v19).f40;
          _nw159[(new BigNumber(7)).toNumber()] = p0;
          _nw159[(new BigNumber(8)).toNumber()] = (_950_v19).f40;
          _nw159[(new BigNumber(9)).toNumber()] = (_this).f33;
          _nw159[(new BigNumber(10)).toNumber()] = (((_953_v22).contains((_this).f34)) ? ((_953_v22).get((_this).f34)) : ((_this).f33));
          _nw159[(new BigNumber(11)).toNumber()] = (_this).f33;
          _nw159[(new BigNumber(12)).toNumber()] = (_this).f33;
          _nw159[(new BigNumber(13)).toNumber()] = (_950_v19).f40;
          _nw159[(new BigNumber(14)).toNumber()] = (((_954_v23).contains((_this).f33)) ? ((_954_v23).get((_this).f33)) : ((_950_v19).f40));
          _nw159[(new BigNumber(15)).toNumber()] = (_950_v19).f40;
          _nw159[(new BigNumber(16)).toNumber()] = (_929_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_929_v0).length))];
          _nw159[(new BigNumber(17)).toNumber()] = (_929_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_929_v0).length))];
          _nw159[(new BigNumber(18)).toNumber()] = p0;
          _nw159[(new BigNumber(19)).toNumber()] = (_950_v19).f40;
          _nw159[(new BigNumber(20)).toNumber()] = p0;
          _nw159[(new BigNumber(21)).toNumber()] = (_929_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_929_v0).length))];
          _nw159[(new BigNumber(22)).toNumber()] = (_this).f33;
          _nw159[(new BigNumber(23)).toNumber()] = (_950_v19).f40;
          _nw159[(new BigNumber(24)).toNumber()] = (_950_v19).f40;
          _nw159[(new BigNumber(25)).toNumber()] = p0;
          _nw159[(new BigNumber(26)).toNumber()] = (_929_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_929_v0).length))];
          _nw159[(new BigNumber(27)).toNumber()] = (_929_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_929_v0).length))];
          _955_v24 = _nw159;
          let _index151 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_952_v21).length));
          (_952_v21)[_index151] = _955_v24;
          let _index152 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_952_v21).length));
          (_952_v21)[_index152] = _955_v24;
          _954_v23 = (_954_v23).update((_dafny.Set.fromElements(_944_i1)).equals(_943_v13), true);
          let _956_v25;
          _956_v25 = _dafny.Map.Empty.slice().updateUnsafe(_949_v18,_944_i1);
          _956_v25 = (_956_v25).update(_949_v18, (_this).f34);
        } else {
          _948_v17 = (_948_v17).update(new BigNumber(-901), new _dafny.CodePoint('y'.codePointAt(0)));
          (globalState).f14 = new BigNumber(566);
          (globalState).f28 = (_950_v19).f40;
          (globalState).f2 = new BigNumber((_dafny.Seq.Concat(_933_v4, _933_v4)).length);
          let _957_v26;
          let _init25 = function (_958_i2) {
            return _module.D8.create_DC26();
          };
          let _nw160 = Array((new BigNumber(5)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw160.length); _i0_25++) {
            _nw160[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _957_v26 = _nw160;
          let _959_v27;
          _959_v27 = _module.D8.create_DC26();
          let _index153 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_957_v26).length));
          (_957_v26)[_index153] = (((_950_v19).f40) ? (_959_v27) : (_module.D8.create_DC26()));
          let _960_v28;
          _960_v28 = _module.D2.create_DC8(_947_v16, (_this).f33);
          let _961_v29;
          _961_v29 = _dafny.Seq.of(_960_v28);
          let _index154 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_957_v26).length));
          let _rhs161 = (((_dafny.ZERO).minus(new BigNumber((_931_v2).cardinality()))).minus(new BigNumber((_module.__default.fm35(_961_v29, _944_i1, globalState)).length))).isLessThanOrEqualTo(_944_i1);
          let _rhs162 = _959_v27;
          let _rhs163 = p0;
          let _lhs152 = globalState;
          let _lhs153 = _957_v26;
          let _lhs154 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_957_v26).length));
          let _lhs155 = globalState;
          _lhs152.f15 = _rhs161;
          _lhs153[_lhs154] = _rhs162;
          _lhs155.f13 = _rhs163;
        }
      }
      r0 = ((p0) ? (p0) : ((_this).f33));
      r1 = (_this).f33;
      return [r0, r1];
    }
    m7(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _962_v0;
      let _init26 = function (_963_i1) {
        return (_963_i1).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f33,(_this).f33)).length));
      };
      let _nw161 = Array((new BigNumber(10)).toNumber());
      for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw161.length); _i0_26++) {
        _nw161[_i0_26] = _init26(new BigNumber(_i0_26));
      }
      _962_v0 = _nw161;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_962_v0).length))) {
        let _964_i0 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_964_i0)) && ((_964_i0).isLessThan(new BigNumber((_962_v0).length))))) {
          (_962_v0)[(_964_i0)] = _module.__default.safeModuloInt(_964_i0, p0);
        }
      }
      let _965_v1;
      _965_v1 = _dafny.Seq.UnicodeFromString("aqds");
      let _966_v2;
      let _nw162 = new _module.C3();
      _nw162.__ctor((_this).f33, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f33,p0)).length), !((_this).f34).isEqualTo((_dafny.ZERO).minus((_this).f34)), _965_v1);
      _966_v2 = _nw162;
      let _967_v3;
      _967_v3 = _dafny.Seq.of((_this).f33);
      let _968_v4;
      _968_v4 = _module.D1.create_DC3(_967_v3);
      let _source15 = (((_966_v2).f35) ? (_module.D1.create_DC3(_dafny.Seq.of((_this).f33))) : (_968_v4));
      if (_source15.is_DC4) {
        let _969___mcc_h0 = (_source15).cf8;
        let _970_cf8 = _969___mcc_h0;
        let _971_v5;
        let _nw163 = Array((new BigNumber(5)).toNumber()).fill([]);
        _971_v5 = _nw163;
        let _972_v6;
        _972_v6 = new _dafny.CodePoint('a'.codePointAt(0));
        let _973_v7;
        let _nw164 = Array((new BigNumber(9)).toNumber());
        _nw164[(_dafny.ZERO).toNumber()] = !(false);
        _nw164[(_dafny.ONE).toNumber()] = (_966_v2).f35;
        _nw164[(new BigNumber(2)).toNumber()] = (_966_v2).f35;
        _nw164[(new BigNumber(3)).toNumber()] = _970_cf8;
        _nw164[(new BigNumber(4)).toNumber()] = false;
        _nw164[(new BigNumber(5)).toNumber()] = (_966_v2).f35;
        _nw164[(new BigNumber(6)).toNumber()] = _module.__default.fm19(_972_v6, globalState);
        _nw164[(new BigNumber(7)).toNumber()] = (_966_v2).f35;
        _nw164[(new BigNumber(8)).toNumber()] = _970_cf8;
        _973_v7 = _nw164;
        let _974_v8;
        _974_v8 = _dafny.Seq.of(_973_v7, _973_v7);
        let _index155 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_971_v5).length));
        (_971_v5)[_index155] = (_974_v8)[_module.__default.safeIndex(new BigNumber(-173), new BigNumber((_974_v8).length))];
        let _975_v9;
        _975_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-577),false);
        let _976_v10;
        _976_v10 = _module.D4.create_DC14(new BigNumber((_975_v9).length));
        let _977_v11;
        let _nw165 = Array((new BigNumber(8)).toNumber());
        _nw165[(_dafny.ZERO).toNumber()] = function (_pat_let21_0) {
          return function (_978_dt__update__tmp_h0) {
            return function (_pat_let22_0) {
              return function (_979_dt__update_hcf23_h0) {
                return _module.D4.create_DC14(_979_dt__update_hcf23_h0);
              }(_pat_let22_0);
            }((_dafny.ZERO).minus((_this).f34));
          }(_pat_let21_0);
        }(_976_v10);
        _nw165[(_dafny.ONE).toNumber()] = _module.D4.create_DC14((_966_v2).f36);
        _nw165[(new BigNumber(2)).toNumber()] = _976_v10;
        _nw165[(new BigNumber(3)).toNumber()] = _module.D4.create_DC14(new BigNumber((_module.__default.fm20((_this).f33, (_966_v2).f35, _967_v3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(565)), ((_980_v2) => function (_981_i2) {
  return (_980_v2).f36;
})(_966_v2))).length), globalState)).length));
        _nw165[(new BigNumber(4)).toNumber()] = _976_v10;
        _nw165[(new BigNumber(5)).toNumber()] = _976_v10;
        _nw165[(new BigNumber(6)).toNumber()] = _976_v10;
        _nw165[(new BigNumber(7)).toNumber()] = _976_v10;
        _977_v11 = _nw165;
        let _index156 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_977_v11).length));
        (_977_v11)[_index156] = _976_v10;
        let _pat_let_tv21 = _966_v2;
        let _index157 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_971_v5).length));
        let _index158 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_977_v11).length));
        let _rhs164 = _973_v7;
        let _rhs165 = function (_pat_let23_0) {
          return function (_982_dt__update__tmp_h1) {
            return function (_pat_let24_0) {
              return function (_983_dt__update_hcf23_h1) {
                return _module.D4.create_DC14(_983_dt__update_hcf23_h1);
              }(_pat_let24_0);
            }((_pat_let_tv21).f36);
          }(_pat_let23_0);
        }(_976_v10);
        let _rhs166 = _970_cf8;
        let _lhs156 = _971_v5;
        let _lhs157 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_971_v5).length));
        let _lhs158 = _977_v11;
        let _lhs159 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_977_v11).length));
        _lhs156[_lhs157] = _rhs164;
        _lhs158[_lhs159] = _rhs165;
        r0 = _rhs166;
        let _984_v12;
        _984_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f33,(_this).f33);
        let _985_v13;
        _985_v13 = _dafny.Set.fromElements(new BigNumber((_984_v12).length), (_966_v2).f36);
        let _986_v14;
        _986_v14 = _module.D9.create_DC28(_967_v3, (_966_v2).f35, (_966_v2).f36, (_966_v2).f35, (_966_v2).f35);
        let _source16 = _module.__default.fm43(_dafny.MultiSet.fromElements(new BigNumber((_985_v13).length)), (_986_v14).dtor_cf43, globalState);
        if (_source16.is_DC4) {
          let _987___mcc_h3 = (_source16).cf8;
          let _988_cf8 = _987___mcc_h3;
          let _rhs167 = (_this).f34;
          let _rhs168 = (_967_v3)[_module.__default.safeIndex((_966_v2).f36, new BigNumber((_967_v3).length))];
          let _lhs160 = globalState;
          let _lhs161 = globalState;
          _lhs160.f19 = _rhs167;
          _lhs161.f18 = _rhs168;
          let _989_v15;
          _989_v15 = _dafny.Seq.of(new BigNumber((_965_v1).length), (_966_v2).f36);
          let _index159 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_962_v0).length));
          (_962_v0)[_index159] = new BigNumber((_dafny.Seq.Concat(_989_v15, _989_v15)).length);
          let _index160 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_962_v0).length));
          (_962_v0)[_index160] = p0;
          (globalState).f11 = _965_v1;
          let _index161 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_962_v0).length));
          (_962_v0)[_index161] = ((_966_v2).f36).minus(p0);
        } else if (_source16.is_DC3) {
          let _990___mcc_h4 = (_source16).cf7;
          let _991_cf7 = _990___mcc_h4;
          _985_v13 = _985_v13;
          let _992_v17;
          _992_v17 = _dafny.Set.fromElements(_970_cf8, _970_cf8, (_this).f33);
          let _993_v18;
          _993_v18 = _dafny.Map.Empty.slice().updateUnsafe(_992_v17,(_966_v2).f35);
          let _994_v19;
          _994_v19 = _dafny.Map.Empty.slice().updateUnsafe((_966_v2).f36,new BigNumber((function () {
            let _coll41 = new _dafny.Map();
            for (const _compr_41 of (_993_v18).Keys.Elements) {
              let _995_v16 = _compr_41;
              if ((_993_v18).contains(_995_v16)) {
                _coll41.push([_995_v16,(_966_v2).f36]);
              }
            }
            return _coll41;
          }()).length));
          let _rhs169 = (_this).f33;
          let _rhs170 = (function () {
            let _coll42 = new _dafny.Map();
            for (const _compr_42 of _dafny.IntegerRange(new BigNumber(73), new BigNumber(979))) {
              let _996_v20 = _compr_42;
              if (((new BigNumber(73)).isLessThanOrEqualTo(_996_v20)) && ((_996_v20).isLessThan(new BigNumber(979)))) {
                _coll42.push([_module.__default.safeModuloInt(_996_v20, p0),new BigNumber((_dafny.Seq.of(p0)).length)]);
              }
            }
            return _coll42;
          }()).update(new BigNumber(551), new BigNumber((_dafny.Seq.update(_965_v1, _module.__default.safeIndex((_966_v2).f36, new BigNumber((_965_v1).length)), _972_v6)).length));
          let _lhs162 = globalState;
          _lhs162.f28 = _rhs169;
          _994_v19 = _rhs170;
          let _997_v21;
          _997_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_991_cf7, _991_cf7),p0);
          _997_v21 = _dafny.Map.Empty.slice().updateUnsafe(_991_cf7,new BigNumber(-239));
          (globalState).f20 = _dafny.Seq.update(_dafny.Seq.Concat(_module.__default.fm7((_966_v2).f36, globalState), _965_v1), _module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(((_966_v2).f36).multipliedBy((_966_v2).f36))), new BigNumber((_dafny.Seq.Concat(_module.__default.fm7((_966_v2).f36, globalState), _965_v1)).length)), _972_v6);
        } else {
          let _998___mcc_h5 = (_source16).cf9;
          let _999_cf9 = _998___mcc_h5;
          let _1000_v22;
          let _nw166 = new _module.C1();
          _nw166.__ctor((_966_v2).f35);
          _1000_v22 = _nw166;
          let _1001_v23;
          _1001_v23 = _dafny.Set.fromElements(_962_v0);
          let _1002_v24;
          let _nw167 = new _module.C2();
          _nw167.__ctor(new BigNumber((_dafny.Seq.Concat(_967_v3, _967_v3)).length), _1001_v23, (_966_v2).f35, _965_v1);
          _1002_v24 = _nw167;
          (globalState).f28 = _970_cf8;
          let _1003_v25;
          _1003_v25 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("jj"));
          _1003_v25 = _dafny.Seq.update(_1003_v25, _module.__default.safeIndex((_966_v2).f36, new BigNumber((_1003_v25).length)), _dafny.Seq.Concat(_965_v1, _965_v1));
        }
        let _index162 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_973_v7).length));
        (_973_v7)[_index162] = false;
        let _index163 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_973_v7).length));
        (_973_v7)[_index163] = (_966_v2).f35;
        (globalState).f14 = p0;
      } else if (_source15.is_DC3) {
        let _1004___mcc_h1 = (_source15).cf7;
        let _1005_cf7 = _1004___mcc_h1;
        (globalState).f28 = (_966_v2).f35;
        r0 = ((_966_v2).f35) === ((_966_v2).f35);
        let _1006_v26;
        let _init27 = ((_1007_v2) => function (_1008_i3) {
          return (_1007_v2).f35;
        })(_966_v2);
        let _nw168 = Array((new BigNumber(24)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw168.length); _i0_27++) {
          _nw168[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1006_v26 = _nw168;
        let _index164 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_1006_v26).length));
        (_1006_v26)[_index164] = (_966_v2).f35;
        let _1009_v27;
        _1009_v27 = _dafny.MultiSet.fromElements(_962_v0, _962_v0);
        let _index165 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_1006_v26).length));
        (_1006_v26)[_index165] = ((_1009_v27).Intersect(_1009_v27)).equals(_1009_v27);
        (globalState).f20 = _dafny.Seq.Concat(_965_v1, _dafny.Seq.Concat(_965_v1, _965_v1));
      } else {
        let _1010___mcc_h2 = (_source15).cf9;
        let _1011_cf9 = _1010___mcc_h2;
        let _1012_v28;
        let _1013_v29;
        let _out14;
        let _out15;
        let _outcollector7 = (_966_v2).m2(!((_966_v2).f35), globalState);
        _out14 = _outcollector7[0];
        _out15 = _outcollector7[1];
        _1012_v28 = _out14;
        _1013_v29 = _out15;
        let _1014_v30;
        _1014_v30 = _dafny.MultiSet.fromElements(new BigNumber(-282), (_this).f34);
        let _1015_v31;
        _1015_v31 = _module.D3.create_DC10(_1014_v30);
        let _1016_v32;
        _1016_v32 = _dafny.Seq.of(_module.D3.create_DC10(_1014_v30), _1015_v31, _1015_v31);
        let _1017_v33;
        _1017_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(26),_1016_v32);
        let _1018_v34;
        _1018_v34 = new _dafny.CodePoint('e'.codePointAt(0));
        _1017_v33 = (((_module.__default.fm19(_1018_v34, globalState)) ? (_1017_v33) : (function () {
          let _coll43 = new _dafny.Map();
          for (const _compr_43 of _dafny.IntegerRange(new BigNumber(717), new BigNumber(624))) {
            let _1019_v35 = _compr_43;
            if (((new BigNumber(717)).isLessThanOrEqualTo(_1019_v35)) && ((_1019_v35).isLessThan(new BigNumber(624)))) {
              _coll43.push([(_1019_v35).plus(p0),_dafny.Seq.Create(_module.__default.abs(new BigNumber(888)), ((_1020_v31) => function (_1021_i4) {
                return _1020_v31;
              })(_1015_v31))]);
            }
          }
          return _coll43;
        }()))).update((_dafny.ZERO).minus(((_966_v2).f36).minus(p0)), _dafny.Seq.Concat(_1016_v32, _1016_v32));
        let _1022_v36;
        let _nw169 = new _module.C1();
        _nw169.__ctor((_966_v2).f35);
        _1022_v36 = _nw169;
        let _1023_v37;
        let _nw170 = Array((new BigNumber(18)).toNumber());
        _nw170[(_dafny.ZERO).toNumber()] = _1022_v36;
        _nw170[(_dafny.ONE).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(2)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(3)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(4)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(5)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(6)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(7)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(8)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(9)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(10)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(11)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(12)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(13)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(14)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(15)).toNumber()] = ((_module.__default.fm19(new _dafny.CodePoint('e'.codePointAt(0)), globalState)) ? (_1022_v36) : (_1022_v36));
        _nw170[(new BigNumber(16)).toNumber()] = _1022_v36;
        _nw170[(new BigNumber(17)).toNumber()] = _1022_v36;
        _1023_v37 = _nw170;
        let _1024_v38;
        _1024_v38 = _dafny.Map.Empty.slice().updateUnsafe((_966_v2).f35,_1023_v37);
        _1023_v37 = (((_1024_v38).contains(!(true))) ? ((_1024_v38).get(!(true))) : (_1023_v37));
        if (!((_966_v2).f35) || (!(((_966_v2).f36).isLessThan(new BigNumber(872))))) {
          (globalState).f2 = (_966_v2).f36;
          let _1025_v39;
          let _nw171 = Array((new BigNumber(6)).toNumber()).fill(false);
          _1025_v39 = _nw171;
          let _index166 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_1025_v39).length));
          (_1025_v39)[_index166] = _1013_v29;
          let _1026_v40;
          let _nw172 = new _module.C0();
          _nw172.__ctor((_966_v2).f35);
          _1026_v40 = _nw172;
          let _1027_v41;
          _1027_v41 = _dafny.Seq.of(_1026_v40, _1026_v40, _1026_v40);
          let _1028_v42;
          _1028_v42 = _dafny.MultiSet.fromElements(_1027_v41, _1027_v41, _dafny.Seq.Concat(_1027_v41, _1027_v41));
          let _index167 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_962_v0).length));
          (_962_v0)[_index167] = (_966_v2).f36;
          let _index168 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_1025_v39).length));
          let _index169 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_962_v0).length));
          let _rhs171 = !(_1026_v40.f37);
          let _rhs172 = _1025_v39;
          let _rhs173 = (_967_v3)[_module.__default.safeIndex((_966_v2).f36, new BigNumber((_967_v3).length))];
          let _rhs174 = _1028_v42;
          let _rhs175 = (_dafny.ZERO).minus((_966_v2).f36);
          let _lhs163 = _1025_v39;
          let _lhs164 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_1025_v39).length));
          let _lhs165 = globalState;
          let _lhs166 = _962_v0;
          let _lhs167 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_962_v0).length));
          _lhs163[_lhs164] = _rhs171;
          _1025_v39 = _rhs172;
          _lhs165.f28 = _rhs173;
          _1028_v42 = _rhs174;
          _lhs166[_lhs167] = _rhs175;
          (globalState).f19 = (_966_v2).f36;
          let _1029_v43;
          _1029_v43 = _dafny.Set.fromElements((_966_v2).f35);
          let _1030_v44;
          _1030_v44 = _dafny.Seq.of(_dafny.Set.fromElements(_1013_v29));
          let _1031_v45;
          _1031_v45 = _dafny.Map.Empty.slice().updateUnsafe((_966_v2).f36,_1029_v43);
          let _1032_v46;
          let _nw173 = Array((new BigNumber(22)).toNumber());
          _nw173[(_dafny.ZERO).toNumber()] = _1029_v43;
          _nw173[(_dafny.ONE).toNumber()] = _1029_v43;
          _nw173[(new BigNumber(2)).toNumber()] = _1029_v43;
          _nw173[(new BigNumber(3)).toNumber()] = _1029_v43;
          _nw173[(new BigNumber(4)).toNumber()] = _1029_v43;
          _nw173[(new BigNumber(5)).toNumber()] = _module.__default.fm22((_this).f34, (_this).f33, _1018_v34, globalState);
          _nw173[(new BigNumber(6)).toNumber()] = (_1029_v43).Union(_1029_v43);
          _nw173[(new BigNumber(7)).toNumber()] = (_dafny.Set.fromElements(true)).Union(_1029_v43);
          _nw173[(new BigNumber(8)).toNumber()] = (_1029_v43).Intersect(_dafny.Set.fromElements(_1026_v40.f37, (_966_v2).f35, _1026_v40.f37));
          _nw173[(new BigNumber(9)).toNumber()] = _1029_v43;
          _nw173[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements((_1025_v39)[_module.__default.safeIndex(new BigNumber(62), new BigNumber((_1025_v39).length))]);
          _nw173[(new BigNumber(11)).toNumber()] = (_1030_v44)[_module.__default.safeIndex((_966_v2).f36, new BigNumber((_1030_v44).length))];
          _nw173[(new BigNumber(12)).toNumber()] = _1029_v43;
          _nw173[(new BigNumber(13)).toNumber()] = _1029_v43;
          _nw173[(new BigNumber(14)).toNumber()] = _dafny.Set.fromElements(_1026_v40.f37);
          _nw173[(new BigNumber(15)).toNumber()] = _1029_v43;
          _nw173[(new BigNumber(16)).toNumber()] = (_module.__default.fm22((_966_v2).f36, _1013_v29, _1018_v34, globalState)).Difference(_1029_v43);
          _nw173[(new BigNumber(17)).toNumber()] = _dafny.Set.fromElements((_966_v2).f35);
          _nw173[(new BigNumber(18)).toNumber()] = _dafny.Set.fromElements(_1026_v40.f37);
          _nw173[(new BigNumber(19)).toNumber()] = (_1029_v43).Difference((((_1031_v45).contains((_966_v2).f36)) ? ((_1031_v45).get((_966_v2).f36)) : (_1029_v43)));
          _nw173[(new BigNumber(20)).toNumber()] = _1029_v43;
          _nw173[(new BigNumber(21)).toNumber()] = (_1029_v43).Difference(_1029_v43);
          _1032_v46 = _nw173;
          _1032_v46 = _1032_v46;
          (globalState).f18 = (_966_v2).f35;
        } else {
          (globalState).f28 = _module.__default.fm6(new BigNumber(-894), globalState);
          _1018_v34 = _1018_v34;
          let _1033_v47;
          _1033_v47 = _dafny.Set.fromElements((_966_v2).f35, (_966_v2).f35, !((_966_v2).f35), (_966_v2).f35, (_966_v2).f35);
          _1033_v47 = _1033_v47;
          let _1034_v48;
          let _init28 = ((_1035_v2, _1036_p0) => function (_1037_i5) {
            return _dafny.Seq.Concat(_dafny.Seq.of((_1035_v2).f36), _dafny.Seq.of(_1036_p0, new BigNumber(814), _1036_p0, (_1035_v2).f36, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1035_v2).f36,(_1035_v2).f36)).length)));
          })(_966_v2, p0);
          let _nw174 = Array((new BigNumber(18)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw174.length); _i0_28++) {
            _nw174[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _1034_v48 = _nw174;
          let _1038_v49;
          _1038_v49 = _dafny.Seq.of((_this).f34, (_dafny.ZERO).minus((_966_v2).f36));
          let _index170 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1034_v48).length));
          (_1034_v48)[_index170] = _1038_v49;
          let _index171 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1034_v48).length));
          (_1034_v48)[_index171] = _1038_v49;
          (globalState).f21 = (_1038_v49)[_module.__default.safeIndex(new BigNumber(-844), new BigNumber((_1038_v49).length))];
        }
      }
      (globalState).f21 = (p0).plus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("oxjnf")).length), (_this).f34));
      (globalState).f21 = ((_966_v2).f36).multipliedBy(_module.__default.safeModuloInt((_966_v2).f36, (_966_v2).f36));
      let _1039_v50;
      _1039_v50 = new _dafny.CodePoint('g'.codePointAt(0));
      let _1040_v51;
      _1040_v51 = _dafny.Set.fromElements(_1039_v50, _module.__default.fm42(true, !((_this).f33), globalState), (((_this).f33) ? (_1039_v50) : (_1039_v50)));
      _1040_v51 = _dafny.Set.fromElements(_1039_v50, _1039_v50, _module.__default.fm42((_this).f33, (_966_v2).f35, globalState), _1039_v50, _1039_v50);
      r0 = (_966_v2).f35;
      return r0;
    }
    get f33() {
      let _this = this;
      return _this._f33;
    };
    get f34() {
      let _this = this;
      return _this._f34;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f30 = _dafny.Seq.UnicodeFromString("");
      this._f29 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f30() {
      let _this = this;
      return _this._f30;
    };
    set f30(value) {
      let _this = this;
      _this._f30 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    __ctor(f29, f30) {
      let _this = this;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      return;
    }
    fm3(p0, globalState) {
      let _this = this;
      if (((_this).f29) === ((_this).f29)) {
        return new BigNumber((_dafny.Seq.Concat(_this.f30, _this.f30)).length);
      } else {
        return (_dafny.ZERO).minus(new BigNumber((_this.f30).length));
      }
    };
    fm4(globalState) {
      let _this = this;
      return ((((_this).f29) ? (new BigNumber(((_module.D2.create_DC6(_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0))))).dtor_cf10).cardinality())) : (new BigNumber(255)))).multipliedBy(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("kxmebkric"), _dafny.Seq.UnicodeFromString("qe"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-273)), function (_1041_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("ugd"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(515)), function (_1042_i1) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }))).length));
    };
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _1043_v0;
      let _nw175 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1043_v0 = _nw175;
      let _1044_v1;
      let _nw176 = Array((new BigNumber(13)).toNumber());
      _nw176[(_dafny.ZERO).toNumber()] = (_this).f29;
      _nw176[(_dafny.ONE).toNumber()] = (_this).f29;
      _nw176[(new BigNumber(2)).toNumber()] = (_this).f29;
      _nw176[(new BigNumber(3)).toNumber()] = !(false);
      _nw176[(new BigNumber(4)).toNumber()] = (_this).f29;
      _nw176[(new BigNumber(5)).toNumber()] = p0;
      _nw176[(new BigNumber(6)).toNumber()] = (_this).f29;
      _nw176[(new BigNumber(7)).toNumber()] = p0;
      _nw176[(new BigNumber(8)).toNumber()] = p0;
      _nw176[(new BigNumber(9)).toNumber()] = (_this).f29;
      _nw176[(new BigNumber(10)).toNumber()] = (_this).f29;
      _nw176[(new BigNumber(11)).toNumber()] = (_this).f29;
      _nw176[(new BigNumber(12)).toNumber()] = (_this).f29;
      _1044_v1 = _nw176;
      let _1045_v2;
      _1045_v2 = new BigNumber(70);
      let _1046_v3;
      _1046_v3 = _dafny.Seq.of((_this).f29);
      let _1047_v4;
      _1047_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1043_v0,(_dafny.ZERO).minus((new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_1044_v1), _module.__default.safeIndex(_1045_v2, new BigNumber((_dafny.Seq.of(_1044_v1)).length)), _1044_v1)).length)).multipliedBy((_this).fm3(_1046_v3, globalState))));
      _1047_v4 = (_1047_v4).update(_1043_v0, (_this).fm4(globalState));
      let _1048_v5;
      _1048_v5 = _dafny.Seq.of(_1045_v2, _1045_v2);
      let _hi5 = (_1048_v5)[_module.__default.safeIndex(new BigNumber((_1048_v5).length), new BigNumber((_1048_v5).length))];
      for (let _1049_i0 = (_1045_v2).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(453)), ((_1053_v2) => function (_1054_i1) {
        return _1053_v2;
      })(_1045_v2))).length)); _1049_i0.isLessThan(_hi5); _1049_i0 = _1049_i0.plus(_dafny.ONE)) {
        let _1050_v6;
        _1050_v6 = _dafny.MultiSet.fromElements((_this).f29, (_this).f29);
        let _rhs176 = (_1050_v6).IsSubsetOf(_1050_v6);
        let _rhs177 = (_1049_i0).isLessThan(_1045_v2);
        let _lhs168 = globalState;
        let _lhs169 = globalState;
        _lhs168.f18 = _rhs176;
        _lhs169.f18 = _rhs177;
        (globalState).f18 = false;
        let _1051_v7;
        let _nw177 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _1051_v7 = _nw177;
        let _1052_v8;
        _1052_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_1051_v7);
        _1052_v8 = (_1052_v8).update((_1045_v2).isLessThanOrEqualTo(new BigNumber((_1050_v6).cardinality())), _1051_v7);
        let _index172 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_1051_v7).length));
        (_1051_v7)[_index172] = new BigNumber(967);
        let _index173 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_1051_v7).length));
        (_1051_v7)[_index173] = (_1045_v2).multipliedBy(_1045_v2);
      }
      (globalState).f21 = _1045_v2;
      let _hi6 = _1045_v2;
      for (let _1055_i2 = _1045_v2; _1055_i2.isLessThan(_hi6); _1055_i2 = _1055_i2.plus(_dafny.ONE)) {
        let _1056_v9;
        let _init29 = ((_1057_v3) => function (_1058_i3) {
          return _dafny.MultiSet.fromElements(_module.D1.create_DC3(_1057_v3));
        })(_1046_v3);
        let _nw178 = Array((new BigNumber(21)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw178.length); _i0_29++) {
          _nw178[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _1056_v9 = _nw178;
        _1056_v9 = _1056_v9;
        (globalState).f13 = p0;
        let _index174 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_1044_v1).length));
        (_1044_v1)[_index174] = (_this).f29;
        let _1059_v10;
        _1059_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1055_i2,p0);
        let _index175 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_1044_v1).length));
        (_1044_v1)[_index175] = ((p0) ? ((((_1059_v10).contains(new BigNumber(((function () {
          let _coll45 = new _dafny.Map();
          for (const _compr_45 of _dafny.IntegerRange(new BigNumber(228), new BigNumber(630))) {
            let _1061_v11 = _compr_45;
            if (((new BigNumber(228)).isLessThanOrEqualTo(_1061_v11)) && ((_1061_v11).isLessThan(new BigNumber(630)))) {
              _coll45.push([(_1061_v11).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("spsbhq")).length)),_1055_i2]);
            }
          }
          return _coll45;
        }()).update(new BigNumber(416), _1055_i2)).length))) ? ((_1059_v10).get(new BigNumber(((function () {
          let _coll44 = new _dafny.Map();
          for (const _compr_44 of _dafny.IntegerRange(new BigNumber(228), new BigNumber(630))) {
            let _1060_v11 = _compr_44;
            if (((new BigNumber(228)).isLessThanOrEqualTo(_1060_v11)) && ((_1060_v11).isLessThan(new BigNumber(630)))) {
              _coll44.push([(_1060_v11).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("spsbhq")).length)),_1055_i2]);
            }
          }
          return _coll44;
        }()).update(new BigNumber(416), _1055_i2)).length))) : (p0))) : ((_1055_i2).isEqualTo(_1045_v2)));
        if (p0) {
          let _1062_v12;
          let _nw179 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _1062_v12 = _nw179;
          let _1063_v13;
          _1063_v13 = _module.D1.create_DC3(_1046_v3);
          let _1064_v14;
          _1064_v14 = _module.D1.create_DC5(_1063_v13);
          let _1065_v15;
          _1065_v15 = _module.D1.create_DC5(_module.D1.create_DC5(_1063_v13));
          let _1066_v16;
          _1066_v16 = _module.D1.create_DC5(_1063_v13);
          let _1067_v17;
          _1067_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1062_v12,_1066_v16);
          _1067_v17 = ((((_this).f29) ? (_1067_v17) : ((_1067_v17).update(_1062_v12, _1066_v16)))).update(_1062_v12, _1066_v16);
          let _1068_v18;
          _1068_v18 = _dafny.Set.fromElements(p0, p0, true);
          (globalState).f21 = ((_1055_i2).plus(new BigNumber((_1068_v18).length))).plus(new BigNumber(556));
          let _1069_v19;
          _1069_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1062_v12,(_1044_v1)[_module.__default.safeIndex(new BigNumber(952), new BigNumber((_1044_v1).length))]);
          let _1070_v20;
          _1070_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1069_v19).length),_1045_v2);
          _1070_v20 = (_1070_v20).update(new BigNumber((_this.f30).length), _1045_v2);
          let _1071_v21;
          _1071_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,p0);
          _1070_v20 = (_1070_v20).update((_1045_v2).plus(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(375)), ((_1072_v20) => function (_1073_i4) {
            return new BigNumber((_1072_v20).length);
          })(_1070_v20)), _module.__default.safeIndex(new BigNumber((_1071_v21).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(375)), ((_1074_v20) => function (_1075_i4) {
            return new BigNumber((_1074_v20).length);
          })(_1070_v20))).length)), _1055_i2)).length)), (_1055_i2).minus(new BigNumber(826)));
          let _1076_v22;
          let _init30 = ((_1077_p0) => function (_1078_i5) {
            return _1077_p0;
          })(p0);
          let _nw180 = Array((new BigNumber(19)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw180.length); _i0_30++) {
            _nw180[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1076_v22 = _nw180;
          let _1079_v23;
          _1079_v23 = _module.D0.create_DC1(_1076_v22, (((_1070_v20).contains(_1055_i2)) ? ((_1070_v20).get(_1055_i2)) : (_1055_i2)), (_dafny.ZERO).minus(_1055_i2));
          _1076_v22 = (_1079_v23).dtor_cf1;
        } else {
          let _1080_v24;
          let _nw181 = Array((new BigNumber(20)).toNumber()).fill(false);
          _1080_v24 = _nw181;
          let _1081_v25;
          _1081_v25 = _dafny.MultiSet.fromElements(p0);
          _1080_v24 = (((_dafny.MultiSet.FromArray(_1046_v3)).IsDisjointFrom(_1081_v25)) ? (_1080_v24) : (_1080_v24));
          let _index176 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_1044_v1).length));
          (_1044_v1)[_index176] = !(_1055_i2).isEqualTo(_1045_v2);
          let _1082_v26;
          _1082_v26 = _dafny.Map.Empty.slice().updateUnsafe((_1044_v1)[_module.__default.safeIndex(new BigNumber(952), new BigNumber((_1044_v1).length))],_1045_v2);
          (globalState).f21 = (_dafny.ZERO).minus((_1045_v2).minus(_module.__default.safeModuloInt(_1045_v2, (((_1082_v26).contains((_this).f29)) ? ((_1082_v26).get((_this).f29)) : (_1045_v2)))));
          r1 = (_1044_v1)[_module.__default.safeIndex(new BigNumber(952), new BigNumber((_1044_v1).length))];
          let _1083_v27;
          let _nw182 = new _module.C0();
          _nw182.__ctor(p0);
          _1083_v27 = _nw182;
        }
      }
      let _1084_v28;
      _1084_v28 = _module.D9.create_DC27(_1048_v5);
      let _1085_v29;
      let _nw183 = Array((new BigNumber(6)).toNumber());
      _nw183[(_dafny.ZERO).toNumber()] = (_1084_v28).dtor_cf40;
      _nw183[(_dafny.ONE).toNumber()] = _1048_v5;
      _nw183[(new BigNumber(2)).toNumber()] = _1048_v5;
      _nw183[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1048_v5, _1048_v5);
      _nw183[(new BigNumber(4)).toNumber()] = (((_this).f29) ? (_1048_v5) : (_dafny.Seq.of(new BigNumber(566))));
      _nw183[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_1045_v2, _1045_v2);
      _1085_v29 = _nw183;
      let _index177 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_1085_v29).length));
      (_1085_v29)[_index177] = _1048_v5;
      let _index178 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_1085_v29).length));
      (_1085_v29)[_index178] = (((_this).f29) ? (((p0) ? (_1048_v5) : (_1048_v5))) : (_1048_v5));
      let _1086_v30;
      _1086_v30 = _dafny.MultiSet.fromElements((p0) === (true), (_1046_v3)[_module.__default.safeIndex(new BigNumber((_this.f30).length), new BigNumber((_1046_v3).length))], (_this).f29);
      _1086_v30 = _1086_v30;
      let _1087_v31;
      _1087_v31 = new _dafny.CodePoint('e'.codePointAt(0));
      let _1088_v32;
      _1088_v32 = _module.D2.create_DC8(_1087_v31, (_this).f29);
      let _1089_v33;
      _1089_v33 = _dafny.Seq.of(_1088_v32);
      let _1090_v34;
      _1090_v34 = _dafny.Seq.of(_module.__default.fm35(_1089_v33, new BigNumber(654), globalState), _dafny.Seq.of((_1046_v3)[_module.__default.safeIndex(_1045_v2, new BigNumber((_1046_v3).length))], (_this).f29, p0), _1046_v3);
      r0 = _dafny.Seq.Concat(_1090_v34, _1090_v34);
      r1 = (_this).f29;
      return [r0, r1];
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let _1091_v0;
      _1091_v0 = _module.D4.create_DC15(p2, p2);
      let _1092_i0;
      _1092_i0 = _dafny.ZERO;
      L5: {
        let _pat_let_tv22 = p0;
        while (function (_source17) {
          if (_source17.is_DC14) {
            let _1097___mcc_h0 = (_source17).cf23;
            let _1098_cf23 = _1097___mcc_h0;
            return (_this).f29;
          } else if (_source17.is_DC15) {
            let _1099___mcc_h1 = (_source17).cf24;
            let _1100___mcc_h2 = (_source17).cf25;
            let _1101_cf25 = _1100___mcc_h2;
            let _1102_cf24 = _1099___mcc_h1;
            return _pat_let_tv22;
          } else {
            let _1103___mcc_h3 = (_source17).cf22;
            let _1104_cf22 = _1103___mcc_h3;
            return (_this).f29;
          }
        }(_1091_v0)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1092_i0)) {
              break L5;
            }
            _1092_i0 = (_1092_i0).plus(_dafny.ONE);
            let _rhs178 = (_this).f29;
            let _rhs179 = (_this).f29;
            let _lhs170 = globalState;
            let _lhs171 = globalState;
            _lhs170.f13 = _rhs178;
            _lhs171.f15 = _rhs179;
            let _1093_v1;
            let _nw184 = Array((new BigNumber(12)).toNumber());
            _1093_v1 = _nw184;
            let _1094_v2;
            let _nw185 = new _module.C4();
            _nw185.__ctor(p0, p2);
            _1094_v2 = _nw185;
            let _index179 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_1093_v1).length));
            (_1093_v1)[_index179] = _1094_v2;
            let _index180 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_1093_v1).length));
            let _rhs180 = _this.f30;
            let _rhs181 = _1094_v2;
            let _rhs182 = _module.__default.fm6(p2, globalState);
            let _rhs183 = p2;
            let _lhs172 = _1093_v1;
            let _lhs173 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_1093_v1).length));
            let _lhs174 = globalState;
            let _lhs175 = globalState;
            r0 = _rhs180;
            _lhs172[_lhs173] = _rhs181;
            _lhs174.f18 = _rhs182;
            _lhs175.f19 = _rhs183;
            (globalState).f2 = p2;
            let _1095_v3;
            _1095_v3 = _module.D10.create_DC33(p0, p2);
            let _1096_v4;
            let _nw186 = new _module.C3();
            _nw186.__ctor((_1095_v3).dtor_cf54, p2, (p2).isEqualTo(new BigNumber(285)), _this.f30);
            _1096_v4 = _nw186;
          }
        }
      }
      let _1105_v5;
      _1105_v5 = _dafny.Seq.of(p0, (_this).f29);
      let _1106_v6;
      _1106_v6 = _dafny.Seq.of(_1105_v5, _1105_v5);
      let _1107_v7;
      _1107_v7 = _module.D1.create_DC3((_1106_v6)[_module.__default.safeIndex(p2, new BigNumber((_1106_v6).length))]);
      let _1108_v8;
      _1108_v8 = _dafny.Map.Empty.slice().updateUnsafe(false,_1107_v7);
      let _1109_v9;
      _1109_v9 = _module.D5.create_DC16(_1108_v8);
      let _pat_let_tv23 = _1108_v8;
      let _source18 = function (_pat_let25_0) {
        return function (_1110_dt__update__tmp_h0) {
          return function (_pat_let26_0) {
            return function (_1111_dt__update_hcf26_h0) {
              return _module.D5.create_DC16(_1111_dt__update_hcf26_h0);
            }(_pat_let26_0);
          }(_pat_let_tv23);
        }(_pat_let25_0);
      }(_1109_v9);
      if (_source18.is_DC17) {
        let _1112___mcc_h4 = (_source18).cf27;
        let _1113_cf27 = _1112___mcc_h4;
        (globalState).f18 = (_this).f29;
        let _1114_v10;
        _1114_v10 = new _dafny.CodePoint('r'.codePointAt(0));
        let _1115_v11;
        _1115_v11 = _dafny.MultiSet.fromElements((_this).f29);
        (globalState).f20 = _dafny.Seq.update(_dafny.Seq.update(_this.f30, _module.__default.safeIndex(p2, new BigNumber((_this.f30).length)), _1114_v10), _module.__default.safeIndex(new BigNumber((_1115_v11).cardinality()), new BigNumber((_dafny.Seq.update(_this.f30, _module.__default.safeIndex(p2, new BigNumber((_this.f30).length)), _1114_v10)).length)), (((_this).f29) ? (_1114_v10) : (_1114_v10)));
        let _1116_v12;
        _1116_v12 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(_module.__default.fm6(p2, globalState)));
        let _1117_v13;
        _1117_v13 = _dafny.Seq.of(p2);
        let _1118_v14;
        _1118_v14 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_this).f29),new BigNumber(342))).length), new BigNumber(((((_this).f29) ? (_1116_v12) : (_module.__default.fm32((_this).f29, (_this).f29, _1117_v13, globalState)))).length));
        (globalState).f7 = _1117_v13;
        (globalState).f1 = _module.__default.safeDivisionInt(p2, (_dafny.ZERO).minus(p2));
      } else if (_source18.is_DC18) {
        (globalState).f18 = (_this).f29;
        let _1119_v15;
        _1119_v15 = _dafny.Seq.of(p2);
        let _1120_v16;
        _1120_v16 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1119_v15);
        let _1121_v17;
        _1121_v17 = _module.D9.create_DC27((((_1120_v16).contains(p2)) ? ((_1120_v16).get(p2)) : (_module.__default.fm44((_this).f29, p2, p0, p0, globalState))));
        _1121_v17 = _module.D9.create_DC27(_1119_v15);
        let _1122_v18;
        _1122_v18 = _module.D0.create_DC2(_dafny.MultiSet.fromElements((_this).f29), _this.f30, true);
        let _1123_v19;
        _1123_v19 = new _dafny.CodePoint('u'.codePointAt(0));
        let _pat_let_tv24 = _1123_v19;
        let _pat_let_tv25 = p2;
        let _pat_let_tv26 = globalState;
        let _source19 = function (_pat_let27_0) {
          return function (_1124_dt__update__tmp_h1) {
            return function (_pat_let28_0) {
              return function (_1125_dt__update_hcf4_h0) {
                return _module.D0.create_DC2(_1125_dt__update_hcf4_h0, (_1124_dt__update__tmp_h1).dtor_cf5, (_1124_dt__update__tmp_h1).dtor_cf6);
              }(_pat_let28_0);
            }(_module.__default.fm39(_pat_let_tv24, _pat_let_tv25, _pat_let_tv26));
          }(_pat_let27_0);
        }(_1122_v18);
        if (_source19.is_DC0) {
          let _1126___mcc_h10 = (_source19).cf0;
          let _1127_cf0 = _1126___mcc_h10;
          (globalState).f20 = _this.f30;
          let _1128_v20;
          let _nw187 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _1128_v20 = _nw187;
          let _1129_v21;
          let _nw188 = new _module.C2();
          _nw188.__ctor(p2, (_dafny.Set.fromElements(_1128_v20, _1128_v20)).Intersect(_dafny.Set.fromElements(_1128_v20, _1128_v20, _1128_v20, _1128_v20, _1128_v20)), p0, _dafny.Seq.Concat(_this.f30, _this.f30));
          _1129_v21 = _nw188;
          let _1130_v22;
          let _nw189 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1130_v22 = _nw189;
          let _index181 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1130_v22).length));
          (_1130_v22)[_index181] = ((_this).f29) || (_1127_cf0);
          let _index182 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1130_v22).length));
          (_1130_v22)[_index182] = _1127_cf0;
          (globalState).f1 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_1129_v21).f38),p2)).length)).multipliedBy((_1129_v21).f38);
        } else if (_source19.is_DC1) {
          let _1131___mcc_h11 = (_source19).cf1;
          let _1132___mcc_h12 = (_source19).cf2;
          let _1133___mcc_h13 = (_source19).cf3;
          let _1134_cf3 = _1133___mcc_h13;
          let _1135_cf2 = _1132___mcc_h12;
          let _1136_cf1 = _1131___mcc_h11;
          let _1137_v23;
          let _nw190 = Array((new BigNumber(16)).toNumber()).fill([]);
          _1137_v23 = _nw190;
          let _1138_v24;
          let _init31 = ((_1139_v15, _1140_cf2) => function (_1141_i1) {
            return (_module.D11.create_DC34(_dafny.Map.Empty.slice().updateUnsafe(_1139_v15,_1140_cf2))).dtor_cf56;
          })(_1119_v15, _1135_cf2);
          let _nw191 = Array((new BigNumber(17)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw191.length); _i0_31++) {
            _nw191[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1138_v24 = _nw191;
          let _1142_v25;
          _1142_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1119_v15,_1134_cf3);
          let _index183 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_1138_v24).length));
          (_1138_v24)[_index183] = ((_1142_v25).update(_1119_v15, _1135_cf2)).update(_dafny.Seq.update(_dafny.Seq.of(_1135_cf2), _module.__default.safeIndex((_dafny.ZERO).minus(_1135_cf2), new BigNumber((_dafny.Seq.of(_1135_cf2)).length)), _1134_cf3), _1135_cf2);
          let _pat_let_tv27 = _1135_cf2;
          let _index184 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_1138_v24).length));
          let _rhs184 = _1137_v23;
          let _rhs185 = _module.__default.fm45(globalState);
          let _rhs186 = function (_pat_let29_0) {
            return function (_1143_dt__update__tmp_h2) {
              return function (_pat_let30_0) {
                return function (_1144_dt__update_hcf25_h0) {
                  return _module.D4.create_DC15((_1143_dt__update__tmp_h2).dtor_cf24, _1144_dt__update_hcf25_h0);
                }(_pat_let30_0);
              }(_pat_let_tv27);
            }(_pat_let29_0);
          }(_1091_v0);
          let _rhs187 = _1135_cf2;
          let _rhs188 = !(!((p2).isEqualTo((_dafny.ZERO).minus(_module.__default.fm9(globalState)))));
          let _lhs176 = _1138_v24;
          let _lhs177 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_1138_v24).length));
          let _lhs178 = globalState;
          let _lhs179 = globalState;
          _1137_v23 = _rhs184;
          _lhs176[_lhs177] = _rhs185;
          _1091_v0 = _rhs186;
          _lhs178.f21 = _rhs187;
          _lhs179.f15 = _rhs188;
          let _index185 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_1136_cf1).length));
          (_1136_cf1)[_index185] = (_this).f29;
          let _index186 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_1136_cf1).length));
          let _rhs189 = _1134_cf3;
          let _rhs190 = (_this).f29;
          let _rhs191 = p0;
          let _rhs192 = _module.__default.fm19(_1123_v19, globalState);
          let _lhs180 = globalState;
          let _lhs181 = _1136_cf1;
          let _lhs182 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_1136_cf1).length));
          let _lhs183 = globalState;
          let _lhs184 = globalState;
          _lhs180.f19 = _rhs189;
          _lhs181[_lhs182] = _rhs190;
          _lhs183.f28 = _rhs191;
          _lhs184.f15 = _rhs192;
          _1135_cf2 = (p2).minus(_1135_cf2);
          (globalState).f28 = (_1135_cf2).isLessThanOrEqualTo(p2);
        } else {
          let _1145___mcc_h14 = (_source19).cf4;
          let _1146___mcc_h15 = (_source19).cf5;
          let _1147___mcc_h16 = (_source19).cf6;
          let _1148_cf6 = _1147___mcc_h16;
          let _1149_cf5 = _1146___mcc_h15;
          let _1150_cf4 = _1145___mcc_h14;
          let _rhs193 = (_dafny.ZERO).minus(p2);
          let _rhs194 = !(_1148_cf6);
          let _rhs195 = _1150_cf4;
          let _lhs185 = globalState;
          let _lhs186 = globalState;
          _lhs185.f19 = _rhs193;
          _lhs186.f18 = _rhs194;
          _1150_cf4 = _rhs195;
          let _1151_v26;
          let _nw192 = new _module.C3();
          _nw192.__ctor(!((_this).f29), p2, ((p0) ? (p0) : ((_this).f29)), (_module.D11.create_DC35(_this.f30, _1148_cf6)).dtor_cf57);
          _1151_v26 = _nw192;
          (globalState).f13 = (_this).f29;
          let _1152_v27;
          let _nw193 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _1152_v27 = _nw193;
          let _1153_v28;
          _1153_v28 = _dafny.Seq.of(_1152_v27, _1152_v27, _1152_v27, _1152_v27, _1152_v27);
          let _1154_v29;
          let _nw194 = Array((new BigNumber(28)).toNumber());
          _nw194[(_dafny.ZERO).toNumber()] = _1152_v27;
          _nw194[(_dafny.ONE).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(2)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(3)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(4)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(5)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(6)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(7)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(8)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(9)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(10)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(11)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(12)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(13)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(14)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(15)).toNumber()] = ((!((_this).f29)) ? (_1152_v27) : (_1152_v27));
          _nw194[(new BigNumber(16)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(17)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(18)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(19)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(20)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(21)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(22)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(23)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(24)).toNumber()] = (_1153_v28)[_module.__default.safeIndex(new BigNumber(21), new BigNumber((_1153_v28).length))];
          _nw194[(new BigNumber(25)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(26)).toNumber()] = _1152_v27;
          _nw194[(new BigNumber(27)).toNumber()] = _1152_v27;
          _1154_v29 = _nw194;
          let _1155_v30;
          _1155_v30 = _module.D12.create_DC38(_1152_v27);
          let _1156_v31;
          _1156_v31 = _dafny.Map.Empty.slice().updateUnsafe((_1151_v26).f35,(_1155_v30).dtor_cf64);
          let _nw195 = Array((new BigNumber(23)).toNumber());
          _nw195[(_dafny.ZERO).toNumber()] = _1152_v27;
          _nw195[(_dafny.ONE).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(2)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(3)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(4)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(5)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(6)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(7)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(8)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(9)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(10)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(11)).toNumber()] = (((_1156_v31).contains(_1148_cf6)) ? ((_1156_v31).get(_1148_cf6)) : (_1152_v27));
          _nw195[(new BigNumber(12)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(13)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(14)).toNumber()] = (_1153_v28)[_module.__default.safeIndex(p2, new BigNumber((_1153_v28).length))];
          _nw195[(new BigNumber(15)).toNumber()] = (_1155_v30).dtor_cf64;
          _nw195[(new BigNumber(16)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(17)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(18)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(19)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(20)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(21)).toNumber()] = _1152_v27;
          _nw195[(new BigNumber(22)).toNumber()] = _1152_v27;
          _1154_v29 = _nw195;
        }
        (globalState).f1 = p2;
      } else if (_source18.is_DC19) {
        let _1157___mcc_h5 = (_source18).cf28;
        let _1158___mcc_h6 = (_source18).cf29;
        let _1159___mcc_h7 = (_source18).cf30;
        let _1160_cf30 = _1159___mcc_h7;
        let _1161_cf29 = _1158___mcc_h6;
        let _1162_cf28 = _1157___mcc_h5;
        (globalState).f15 = (p0) || ((_1161_cf29).isLessThanOrEqualTo(new BigNumber((_this.f30).length)));
        let _1163_v32;
        let _nw196 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
        _1163_v32 = _nw196;
        let _1164_v33;
        _1164_v33 = _module.D2.create_DC7((_this).f29);
        let _1165_v34;
        _1165_v34 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1164_v33).dtor_cf11);
        let _1166_v35;
        _1166_v35 = _dafny.Map.Empty.slice().updateUnsafe(true,_1165_v34);
        let _1167_v36;
        _1167_v36 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1166_v35);
        let _index187 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1163_v32).length));
        (_1163_v32)[_index187] = ((((_1167_v36).contains(false)) ? ((_1167_v36).get(false)) : (_module.__default.fm46(new _dafny.CodePoint('k'.codePointAt(0)), _this.f30, globalState)))).update(!((_this).f29), _1165_v34);
        let _1168_v37;
        _1168_v37 = _dafny.MultiSet.fromElements((_this).f29, (_this).f29);
        let _1169_v38;
        _1169_v38 = _dafny.Seq.of(new BigNumber((_1168_v37).cardinality()));
        let _1170_v39;
        _1170_v39 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1162_cf28,p0), _1165_v34, _1165_v34, _1165_v34, _1165_v34);
        let _1171_v40;
        _1171_v40 = _dafny.Set.fromElements(new BigNumber(197));
        let _1172_v41;
        let _nw197 = Array((new BigNumber(9)).toNumber());
        _nw197[(_dafny.ZERO).toNumber()] = (_1169_v38)[_module.__default.safeIndex(_1162_cf28, new BigNumber((_1169_v38).length))];
        _nw197[(_dafny.ONE).toNumber()] = _1161_cf29;
        _nw197[(new BigNumber(2)).toNumber()] = p2;
        _nw197[(new BigNumber(3)).toNumber()] = _1162_cf28;
        _nw197[(new BigNumber(4)).toNumber()] = (_this).fm4(globalState);
        _nw197[(new BigNumber(5)).toNumber()] = (((_1170_v39).contains(_module.__default.fm41(new BigNumber(18), p0, p0, !((_this).f29), globalState))) ? ((_1170_v39).get(_module.__default.fm41(new BigNumber(18), p0, p0, !((_this).f29), globalState))) : (new BigNumber((_1171_v40).length)));
        _nw197[(new BigNumber(6)).toNumber()] = _1161_cf29;
        _nw197[(new BigNumber(7)).toNumber()] = p2;
        _nw197[(new BigNumber(8)).toNumber()] = _1161_cf29;
        _1172_v41 = _nw197;
        let _1173_v42;
        _1173_v42 = _dafny.Seq.of(_1172_v41, _1172_v41);
        let _1174_v43;
        _1174_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1173_v42,new BigNumber((_this.f30).length));
        let _index188 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1163_v32).length));
        let _rhs196 = _1166_v35;
        let _rhs197 = _1161_cf29;
        let _rhs198 = _dafny.Seq.Concat(_this.f30, _this.f30);
        let _rhs199 = (((_1174_v43).contains(_1173_v42)) ? ((_1174_v43).get(_1173_v42)) : (_module.__default.safeDivisionInt(p2, _1162_cf28)));
        let _lhs187 = _1163_v32;
        let _lhs188 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1163_v32).length));
        let _lhs189 = globalState;
        let _lhs190 = globalState;
        let _lhs191 = globalState;
        _lhs187[_lhs188] = _rhs196;
        _lhs189.f14 = _rhs197;
        _lhs190.f11 = _rhs198;
        _lhs191.f21 = _rhs199;
        let _1175_v44;
        let _nw198 = new _module.C4();
        _nw198.__ctor(p0, p2);
        _1175_v44 = _nw198;
        (globalState).f21 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(832)), ((_1176_v5) => function (_1177_i2) {
          return _dafny.Seq.Concat(_1176_v5, _1176_v5);
        })(_1105_v5))).length));
      } else if (_source18.is_DC16) {
        let _1178___mcc_h8 = (_source18).cf26;
        let _1179_cf26 = _1178___mcc_h8;
        let _1180_v45;
        _1180_v45 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1181_v46;
        let _init32 = ((_1182_p2) => function (_1183_i3) {
          return _module.__default.safeModuloInt(_1183_i3, _1182_p2);
        })(p2);
        let _nw199 = Array((new BigNumber(5)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw199.length); _i0_32++) {
          _nw199[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1181_v46 = _nw199;
        let _1184_v47;
        _1184_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1180_v45,_1181_v46);
        let _1185_v48;
        let _nw200 = Array((new BigNumber(16)).toNumber());
        _nw200[(_dafny.ZERO).toNumber()] = _1184_v47;
        _nw200[(_dafny.ONE).toNumber()] = _1184_v47;
        _nw200[(new BigNumber(2)).toNumber()] = _1184_v47;
        _nw200[(new BigNumber(3)).toNumber()] = _1184_v47;
        _nw200[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1180_v45,_1181_v46)).update(_module.__default.fm42((_this).f29, p0, globalState), _1181_v46);
        _nw200[(new BigNumber(5)).toNumber()] = _1184_v47;
        _nw200[(new BigNumber(6)).toNumber()] = (_1184_v47).update(_1180_v45, _1181_v46);
        _nw200[(new BigNumber(7)).toNumber()] = _1184_v47;
        _nw200[(new BigNumber(8)).toNumber()] = _1184_v47;
        _nw200[(new BigNumber(9)).toNumber()] = (_1184_v47).update(_1180_v45, _1181_v46);
        _nw200[(new BigNumber(10)).toNumber()] = _1184_v47;
        _nw200[(new BigNumber(11)).toNumber()] = (_1184_v47).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1180_v45,_1181_v46));
        _nw200[(new BigNumber(12)).toNumber()] = _1184_v47;
        _nw200[(new BigNumber(13)).toNumber()] = _1184_v47;
        _nw200[(new BigNumber(14)).toNumber()] = _1184_v47;
        _nw200[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1180_v45,_1181_v46);
        _1185_v48 = _nw200;
        let _index189 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_1185_v48).length));
        (_1185_v48)[_index189] = _1184_v47;
        let _index190 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_1185_v48).length));
        (_1185_v48)[_index190] = (_1184_v47).update(new _dafny.CodePoint('d'.codePointAt(0)), _1181_v46);
        (globalState).f15 = (_this).f29;
        (globalState).f18 = p0;
        (globalState).f18 = (_this).f29;
      } else {
        let _1186___mcc_h9 = (_source18).cf31;
        let _1187_cf31 = _1186___mcc_h9;
        (globalState).f14 = p2;
        let _1188_v49;
        let _nw201 = Array((new BigNumber(26)).toNumber()).fill(false);
        _1188_v49 = _nw201;
        let _1189_v50;
        _1189_v50 = _dafny.Seq.of(new BigNumber(186), p2, new BigNumber((_1105_v5).length));
        let _1190_v51;
        _1190_v51 = _module.D0.create_DC1(_1188_v49, new BigNumber((_1105_v5).length), (_1189_v50)[_module.__default.safeIndex((_this).fm4(globalState), new BigNumber((_1189_v50).length))]);
        let _1191_v52;
        _1191_v52 = _dafny.Seq.of(_1188_v49);
        let _1192_v53;
        let _nw202 = Array((new BigNumber(29)).toNumber());
        _nw202[(_dafny.ZERO).toNumber()] = _1188_v49;
        _nw202[(_dafny.ONE).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(2)).toNumber()] = ((p0) ? (_1188_v49) : (_1188_v49));
        _nw202[(new BigNumber(3)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(4)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(5)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(6)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(7)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(8)).toNumber()] = (_1190_v51).dtor_cf1;
        _nw202[(new BigNumber(9)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(10)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(11)).toNumber()] = (_1191_v52)[_module.__default.safeIndex(p2, new BigNumber((_1191_v52).length))];
        _nw202[(new BigNumber(12)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(13)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(14)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(15)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(16)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(17)).toNumber()] = (((_this).f29) ? (_1188_v49) : (_1188_v49));
        _nw202[(new BigNumber(18)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(19)).toNumber()] = ((p0) ? (_1188_v49) : (_1188_v49));
        _nw202[(new BigNumber(20)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(21)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(22)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(23)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(24)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(25)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(26)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(27)).toNumber()] = _1188_v49;
        _nw202[(new BigNumber(28)).toNumber()] = _1188_v49;
        _1192_v53 = _nw202;
        _1192_v53 = (((_this).f29) ? (_1192_v53) : (_1192_v53));
        let _1193_v54;
        let _nw203 = new _module.C1();
        _nw203.__ctor((_this).f29);
        _1193_v54 = _nw203;
        let _1194_v55;
        _1194_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1193_v54,(_this).f29);
        let _1195_v56;
        _1195_v56 = _dafny.Map.Empty.slice().updateUnsafe(p2,(new BigNumber((_1194_v55).length)).isLessThan(p2));
        _1195_v56 = function () {
          let _coll46 = new _dafny.Map();
          for (const _compr_46 of (_1189_v50).Elements) {
            let _1196_v57 = _compr_46;
            if (_dafny.Seq.contains(_1189_v50, _1196_v57)) {
              _coll46.push([(_1196_v57).minus(p2),p0]);
            }
          }
          return _coll46;
        }();
        let _1197_v58;
        let _1198_v59;
        let _out16;
        let _out17;
        let _outcollector8 = (_1193_v54).m2(p0, globalState);
        _out16 = _outcollector8[0];
        _out17 = _outcollector8[1];
        _1197_v58 = _out16;
        _1198_v59 = _out17;
      }
      let _1199_v60;
      _1199_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_dafny.Set.fromElements((_this).f29));
      if (!((p2).isLessThan(new BigNumber((_1199_v60).length))) || (p0)) {
        (globalState).f21 = new BigNumber(744);
        let _1200_v61;
        _1200_v61 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        (globalState).f28 = (p2).isLessThanOrEqualTo(new BigNumber((_1200_v61).length));
        (globalState).f18 = ((p2).multipliedBy(new BigNumber((_this.f30).length))).isLessThanOrEqualTo(p2);
        (globalState).f19 = ((p0) ? (_module.__default.safeModuloInt(p2, p2)) : (new BigNumber((_dafny.Set.fromElements(p2)).length)));
        (globalState).f28 = p0;
      } else {
        let _1201_v62;
        let _nw204 = new _module.C4();
        _nw204.__ctor((_this).f29, p2);
        _1201_v62 = _nw204;
        let _1202_v63;
        _1202_v63 = _module.D10.create_DC30(_1201_v62);
        let _source20 = _1202_v63;
        if (_source20.is_DC31) {
          let _1203___mcc_h17 = (_source20).cf50;
          let _1204___mcc_h18 = (_source20).cf51;
          let _1205_cf51 = _1204___mcc_h18;
          let _1206_cf50 = _1203___mcc_h17;
          let _1207_v64;
          _1207_v64 = _module.D9.create_DC29(p2, p2, _1205_cf51);
          _1207_v64 = _1207_v64;
          let _1208_v65;
          let _nw205 = new _module.C3();
          _nw205.__ctor(_1205_cf51, _module.__default.safeDivisionInt(new BigNumber(476), p2), (_this).f29, _this.f30);
          _1208_v65 = _nw205;
          let _1209_v66;
          _1209_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1206_cf50,(_1208_v65).f35);
          let _rhs200 = (_1209_v66).Merge((_1209_v66).Merge(_1209_v66));
          let _rhs201 = p2;
          let _lhs192 = globalState;
          _1209_v66 = _rhs200;
          _lhs192.f21 = _rhs201;
          let _1210_v67;
          let _nw206 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _1210_v67 = _nw206;
          let _1211_v68;
          _1211_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1210_v67,true);
          let _1212_v69;
          _1212_v69 = new _dafny.CodePoint('n'.codePointAt(0));
          let _1213_v70;
          _1213_v70 = _module.D2.create_DC8(_1212_v69, _1206_cf50);
          let _rhs202 = _1211_v68;
          let _rhs203 = _dafny.Seq.Concat(_module.__default.fm35(_dafny.Seq.of(_1213_v70), p2, globalState), _dafny.Seq.Concat(_dafny.Seq.of((_1208_v65).f35, _1205_cf51), _1105_v5));
          let _rhs204 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_this.f30, _module.__default.safeIndex(p2, new BigNumber((_this.f30).length)), _1212_v69), _this.f30), _dafny.Seq.Concat(_this.f30, _this.f30));
          let _rhs205 = _1105_v5;
          let _rhs206 = (_1206_cf50) || (p0);
          let _lhs193 = _this;
          let _lhs194 = globalState;
          _1211_v68 = _rhs202;
          _1105_v5 = _rhs203;
          _lhs193.f30 = _rhs204;
          _1105_v5 = _rhs205;
          _lhs194.f13 = _rhs206;
        } else if (_source20.is_DC32) {
          let _1214___mcc_h19 = (_source20).cf52;
          let _1215___mcc_h20 = (_source20).cf53;
          let _1216_cf53 = _1215___mcc_h20;
          let _1217_cf52 = _1214___mcc_h19;
          let _1218_v71;
          let _nw207 = new _module.C0();
          _nw207.__ctor((_this).f29);
          _1218_v71 = _nw207;
          let _1219_v72;
          let _nw208 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
          _1219_v72 = _nw208;
          _1219_v72 = _1219_v72;
          (globalState).f19 = p2;
          let _1220_v73;
          _1220_v73 = _module.D9.create_DC28(_dafny.Seq.of(p0, !((_this).f29)), _1216_cf53, p2, _1216_cf53, _1218_v71.f37);
          let _1221_v74;
          _1221_v74 = _dafny.Map.Empty.slice().updateUnsafe(p2,_module.__default.fm42(_1218_v71.f37, (_1220_v73).dtor_cf45, globalState));
          _1221_v74 = (_1221_v74).update((p2).multipliedBy(p2), _1217_cf52);
        } else if (_source20.is_DC33) {
          let _1222___mcc_h21 = (_source20).cf54;
          let _1223___mcc_h22 = (_source20).cf55;
          let _1224_cf55 = _1223___mcc_h22;
          let _1225_cf54 = _1222___mcc_h21;
          let _1226_v75;
          let _nw209 = new _module.C1();
          _nw209.__ctor(p0);
          _1226_v75 = _nw209;
          let _1227_v76;
          let _nw210 = new _module.C1();
          _nw210.__ctor((_this).f29);
          _1227_v76 = _nw210;
          let _1228_v77;
          let _nw211 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _1228_v77 = _nw211;
          let _index191 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_1228_v77).length));
          (_1228_v77)[_index191] = p2;
          let _index192 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_1228_v77).length));
          (_1228_v77)[_index192] = p2;
          (globalState).f1 = (_1228_v77)[_module.__default.safeIndex(new BigNumber(379), new BigNumber((_1228_v77).length))];
        } else {
          let _1229___mcc_h23 = (_source20).cf49;
          let _1230_cf49 = _1229___mcc_h23;
          let _1231_v78;
          _1231_v78 = _dafny.MultiSet.fromElements(p0);
          _1231_v78 = _1231_v78;
          let _1232_v79;
          let _init33 = ((_1233_p2, _1234_v5) => function (_1235_i4) {
            return (_1235_i4).minus(new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(_1233_p2), _dafny.MultiSet.FromArray(_dafny.Seq.of(_1233_p2, _1233_p2, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_1234_v5).length), _1233_p2, _1233_p2, new BigNumber(962)))).cardinality()))))).length));
          })(p2, _1105_v5);
          let _nw212 = Array((new BigNumber(14)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw212.length); _i0_33++) {
            _nw212[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1232_v79 = _nw212;
          (globalState).f17 = _1232_v79;
          (globalState).f15 = (_this).f29;
          let _1236_v80;
          _1236_v80 = new _dafny.CodePoint('c'.codePointAt(0));
          _1236_v80 = _1236_v80;
        }
        let _1237_v81;
        let _nw213 = Array((new BigNumber(18)).toNumber()).fill(false);
        _1237_v81 = _nw213;
        let _1238_v82;
        let _nw214 = Array((new BigNumber(5)).toNumber());
        _nw214[(_dafny.ZERO).toNumber()] = _1237_v81;
        _nw214[(_dafny.ONE).toNumber()] = _1237_v81;
        _nw214[(new BigNumber(2)).toNumber()] = _1237_v81;
        _nw214[(new BigNumber(3)).toNumber()] = _1237_v81;
        _nw214[(new BigNumber(4)).toNumber()] = _1237_v81;
        _1238_v82 = _nw214;
        let _index193 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1238_v82).length));
        (_1238_v82)[_index193] = _1237_v81;
        let _1239_v83;
        _1239_v83 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v81,_module.__default.fm10(p2, p2, globalState));
        let _1240_v84;
        _1240_v84 = _dafny.Seq.of(new BigNumber((_1239_v83).length), (_dafny.ZERO).minus(new BigNumber((_this.f30).length)));
        let _index194 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1238_v82).length));
        let _rhs207 = new BigNumber((_dafny.Seq.Concat(_1240_v84, _dafny.Seq.Concat(_1240_v84, _1240_v84))).length);
        let _rhs208 = p2;
        let _rhs209 = _1237_v81;
        let _rhs210 = (_this).f29;
        let _lhs195 = globalState;
        let _lhs196 = globalState;
        let _lhs197 = _1238_v82;
        let _lhs198 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1238_v82).length));
        let _lhs199 = globalState;
        _lhs195.f19 = _rhs207;
        _lhs196.f21 = _rhs208;
        _lhs197[_lhs198] = _rhs209;
        _lhs199.f15 = _rhs210;
        let _1241_v85;
        _1241_v85 = _dafny.MultiSet.fromElements((_this).f29, p0);
        let _1242_v86;
        _1242_v86 = _dafny.Set.fromElements(p2, p2, new BigNumber((_dafny.Seq.of(new BigNumber((_1241_v85).cardinality()))).length), p2, p2);
        let _1243_v87;
        _1243_v87 = _module.D8.create_DC26();
        let _1244_v88;
        _1244_v88 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v86,_1243_v87);
        _1244_v88 = (_1244_v88).update(_1242_v86, _1243_v87);
        let _1245_v89;
        let _nw215 = Array((new BigNumber(14)).toNumber()).fill(_module.D10.Default());
        _1245_v89 = _nw215;
        let _1246_v90;
        _1246_v90 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm34(p2, p2, p0, true, globalState),_1245_v89);
        let _1247_v91;
        _1247_v91 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _rhs211 = new BigNumber((_dafny.Seq.of((((_1246_v90).contains(_dafny.Seq.UnicodeFromString("strdxawaa"))) ? ((_1246_v90).get(_dafny.Seq.UnicodeFromString("strdxawaa"))) : (_1245_v89)), _1245_v89, _1245_v89, _1245_v89, _1245_v89)).length);
        let _rhs212 = (_this).f29;
        let _rhs213 = !(false) || ((p2).isEqualTo(p2));
        let _rhs214 = (_dafny.ZERO).minus(p2);
        let _rhs215 = (p2).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_this.f30).length), p2, new BigNumber((_1247_v91).length))).length));
        let _lhs200 = globalState;
        let _lhs201 = globalState;
        let _lhs202 = globalState;
        let _lhs203 = globalState;
        let _lhs204 = globalState;
        _lhs200.f2 = _rhs211;
        _lhs201.f28 = _rhs212;
        _lhs202.f15 = _rhs213;
        _lhs203.f19 = _rhs214;
        _lhs204.f2 = _rhs215;
        let _1248_v92;
        _1248_v92 = _module.D3.create_DC11(p0, p0);
        let _1249_v93;
        _1249_v93 = _module.D3.create_DC12(_1248_v92);
        let _1250_v94;
        _1250_v94 = _dafny.Set.fromElements(_1249_v93);
        let _1251_v95;
        _1251_v95 = _dafny.Set.fromElements(_1250_v94, _1250_v94);
        (globalState).f18 = (_1251_v95).IsSubsetOf((_dafny.Set.fromElements(_1250_v94)).Union(_1251_v95));
      }
      let _1252_i5;
      _1252_i5 = _dafny.ZERO;
      L6: {
        while ((_this).f29) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1252_i5)) {
              break L6;
            }
            _1252_i5 = (_1252_i5).plus(_dafny.ONE);
            let _1253_v96;
            _1253_v96 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,p0);
            let _1254_v97;
            _1254_v97 = _module.D11.create_DC37(p2, p0);
            let _1255_v98;
            let _nw216 = Array((new BigNumber(14)).toNumber());
            _nw216[(_dafny.ZERO).toNumber()] = p0;
            _nw216[(_dafny.ONE).toNumber()] = (_this).f29;
            _nw216[(new BigNumber(2)).toNumber()] = !(!(p0));
            _nw216[(new BigNumber(3)).toNumber()] = p0;
            _nw216[(new BigNumber(4)).toNumber()] = p0;
            _nw216[(new BigNumber(5)).toNumber()] = p0;
            _nw216[(new BigNumber(6)).toNumber()] = true;
            _nw216[(new BigNumber(7)).toNumber()] = _module.__default.fm33(p0, p2, p2, globalState);
            _nw216[(new BigNumber(8)).toNumber()] = (_this).f29;
            _nw216[(new BigNumber(9)).toNumber()] = (_this).f29;
            _nw216[(new BigNumber(10)).toNumber()] = (((_1253_v96).contains(p0)) ? ((_1253_v96).get(p0)) : ((_1105_v5)[_module.__default.safeIndex(p2, new BigNumber((_1105_v5).length))]));
            _nw216[(new BigNumber(11)).toNumber()] = (_1254_v97).dtor_cf63;
            _nw216[(new BigNumber(12)).toNumber()] = (_this).f29;
            _nw216[(new BigNumber(13)).toNumber()] = p0;
            _1255_v98 = _nw216;
            let _1256_v99;
            _1256_v99 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1255_v98);
            _1256_v99 = (_1256_v99).update(p2, _1255_v98);
            let _1257_v100;
            let _nw217 = new _module.C4();
            _nw217.__ctor(true, (new BigNumber(-450)).multipliedBy(p2));
            _1257_v100 = _nw217;
            (globalState).f15 = !((_1257_v100).f33) || ((_this).f29);
            let _1258_v101;
            _1258_v101 = new _dafny.CodePoint('y'.codePointAt(0));
            let _1259_v102;
            _1259_v102 = _module.D5.create_DC19(new BigNumber(453), new BigNumber((_dafny.MultiSet.FromArray(_1105_v5)).cardinality()), _1258_v101);
            _1258_v101 = (_1259_v102).dtor_cf30;
          }
        }
      }
      let _1260_v103;
      _1260_v103 = _module.D9.create_DC29(p2, p2, (_this).f29);
      (globalState).f21 = (_1260_v103).dtor_cf46;
      if (p0) {
        let _1261_v104;
        _1261_v104 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(p2, p2));
        let _1262_v105;
        _1262_v105 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1261_v104);
        let _1263_v106;
        _1263_v106 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        let _1264_v107;
        _1264_v107 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1263_v106).length));
        _1261_v104 = (((_1262_v105).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1264_v107,!((_this).f29))).length))) ? ((_1262_v105).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1264_v107,!((_this).f29))).length))) : ((_1261_v104).update(p2, _module.__default.abs(new BigNumber((_this.f30).length)))));
        let _1265_v108;
        _1265_v108 = new _dafny.CodePoint('e'.codePointAt(0));
        _1265_v108 = _1265_v108;
        let _1266_v109;
        let _nw218 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
        _1266_v109 = _nw218;
        let _1267_v110;
        _1267_v110 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f29);
        let _index195 = _module.__default.safeIndex(new BigNumber(975), new BigNumber((_1266_v109).length));
        (_1266_v109)[_index195] = _1267_v110;
        let _index196 = _module.__default.safeIndex(new BigNumber(975), new BigNumber((_1266_v109).length));
        (_1266_v109)[_index196] = _1267_v110;
        (globalState).f20 = _dafny.Seq.Concat(_dafny.Seq.Concat(_this.f30, _dafny.Seq.update(_this.f30, _module.__default.safeIndex(p2, new BigNumber((_this.f30).length)), _1265_v108)), _dafny.Seq.Concat(_this.f30, _this.f30));
        let _1268_v111;
        _1268_v111 = _dafny.Seq.of(_this.f30);
        let _1269_v112;
        _1269_v112 = _dafny.MultiSet.fromElements(_this.f30, (_1268_v111)[_module.__default.safeIndex(new BigNumber(-182), new BigNumber((_1268_v111).length))]);
        let _1270_v113;
        _1270_v113 = _dafny.Seq.of(_1269_v112);
        (globalState).f18 = ((_1270_v113)[_module.__default.safeIndex(p2, new BigNumber((_1270_v113).length))]).IsProperSubsetOf((_1269_v112).Difference(_1269_v112));
      } else {
        let _1271_v114;
        let _init34 = ((_1272_p2) => function (_1273_i6) {
          return (_1273_i6).multipliedBy(_1272_p2);
        })(p2);
        let _nw219 = Array((new BigNumber(28)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw219.length); _i0_34++) {
          _nw219[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1271_v114 = _nw219;
        let _1274_v115;
        _1274_v115 = _dafny.Set.fromElements(_1271_v114);
        let _1275_v116;
        let _nw220 = new _module.C2();
        _nw220.__ctor(p2, _1274_v115, (_this).f29, _this.f30);
        _1275_v116 = _nw220;
        let _1276_v117;
        let _nw221 = Array((new BigNumber(13)).toNumber()).fill(false);
        _1276_v117 = _nw221;
        let _1277_v118;
        _1277_v118 = _dafny.Seq.of(_1276_v117);
        let _1278_v119;
        _1278_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1275_v116,_1277_v118);
        let _1279_v120;
        _1279_v120 = _dafny.Map.Empty.slice().updateUnsafe((((_1278_v119).contains(_1275_v116)) ? ((_1278_v119).get(_1275_v116)) : (_1277_v118)),(_1275_v116).f38);
        _1279_v120 = (_1279_v120).update(_1277_v118, new BigNumber(80));
        let _1280_v121;
        _1280_v121 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1281_v122;
        _1281_v122 = _module.D1.create_DC4((_this).f29);
        let _1282_v123;
        _1282_v123 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_this.f30).length),_module.D11.create_DC35(_dafny.Seq.update(_this.f30, _module.__default.safeIndex((_1275_v116).f38, new BigNumber((_this.f30).length)), _1280_v121), (_1281_v122).dtor_cf8));
        let _1283_v124;
        _1283_v124 = _dafny.Set.fromElements(_1282_v123);
        let _1284_v125;
        _1284_v125 = _dafny.MultiSet.fromElements(p0);
        let _1285_v126;
        _1285_v126 = _dafny.Set.fromElements((_dafny.MultiSet.fromElements((_this).f29)).equals(_1284_v125));
        let _1286_v127;
        _1286_v127 = _module.D11.create_DC37(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(948)), ((_1287_v121) => function (_1288_i7) {
  return _1287_v121;
})(_1280_v121))).length), (_this).f29);
        let _1289_v128;
        _1289_v128 = _dafny.Seq.of((_1275_v116).f38, (_1275_v116).f38);
        let _1290_v129;
        _1290_v129 = _dafny.Map.Empty.slice().updateUnsafe(_1289_v128,_this.f30);
        let _rhs216 = true;
        let _rhs217 = (_1286_v127).dtor_cf62;
        let _rhs218 = _module.__default.fm47(p0, (_this).f29, globalState);
        let _rhs219 = _1285_v126;
        let _rhs220 = _dafny.Seq.contains((((_this).f29) ? (_this.f30) : ((((_1290_v129).contains(_1289_v128)) ? ((_1290_v129).get(_1289_v128)) : (_dafny.Seq.UnicodeFromString("ysikfqdgi"))))), _1280_v121);
        let _lhs205 = globalState;
        let _lhs206 = globalState;
        let _lhs207 = globalState;
        _lhs205.f18 = _rhs216;
        _lhs206.f2 = _rhs217;
        _1283_v124 = _rhs218;
        _1285_v126 = _rhs219;
        _lhs207.f18 = _rhs220;
        let _1291_v130;
        let _nw222 = Array((new BigNumber(2)).toNumber()).fill([]);
        _1291_v130 = _nw222;
        _1291_v130 = _1291_v130;
        let _1292_v131;
        let _nw223 = new _module.C0();
        _nw223.__ctor(p0);
        _1292_v131 = _nw223;
        let _1293_v132;
        _1293_v132 = _dafny.Seq.of(_1292_v131);
        (globalState).f21 = ((_1275_v116).f38).plus((new BigNumber((_1105_v5).length)).minus(new BigNumber((_1293_v132).length)));
        _1285_v126 = _1285_v126;
      }
      r0 = _this.f30;
      let _1294_v133;
      _1294_v133 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1105_v5);
      r1 = (((_this).f29) ? (_dafny.Seq.contains((((_1294_v133).contains(false)) ? ((_1294_v133).get(false)) : (_1105_v5)), !((_this).f29))) : (p0));
      return [r0, r1];
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f30 = _dafny.Seq.UnicodeFromString("");
      this._f29 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f30() {
      let _this = this;
      return _this._f30;
    };
    set f30(value) {
      let _this = this;
      _this._f30 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    __ctor(f29, f30) {
      let _this = this;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      return;
    }
    fm2(p0, globalState) {
      let _this = this;
      return (_this).f29;
    };
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _1295_v0;
      _1295_v0 = new BigNumber(914);
      let _1296_v1;
      _1296_v1 = _dafny.MultiSet.fromElements(_1295_v0, _1295_v0, _1295_v0, _1295_v0);
      let _hi7 = _1295_v0;
      for (let _1297_i0 = new BigNumber((_1296_v1).cardinality()); _1297_i0.isLessThan(_hi7); _1297_i0 = _1297_i0.plus(_dafny.ONE)) {
        (globalState).f19 = (_1297_i0).plus((_dafny.ZERO).minus(_1297_i0));
        let _1298_v2;
        let _nw224 = Array((new BigNumber(15)).toNumber()).fill([]);
        _1298_v2 = _nw224;
        let _nw225 = Array((new BigNumber(5)).toNumber()).fill([]);
        _1298_v2 = _nw225;
        if (!(!((_this).f29) || ((_this).f29))) {
          let _1299_v3;
          let _nw226 = Array((new BigNumber(12)).toNumber()).fill(false);
          _1299_v3 = _nw226;
          let _1300_v4;
          _1300_v4 = _module.D0.create_DC1(_1299_v3, _1295_v0, _1297_i0);
          (globalState).f1 = (_1300_v4).dtor_cf3;
          let _1301_v5;
          _1301_v5 = _dafny.MultiSet.fromElements(p0);
          let _1302_v6;
          _1302_v6 = _module.D0.create_DC2(_1301_v5, _this.f30, p0);
          let _1303_v7;
          _1303_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1302_v6,_1295_v0);
          _1303_v7 = _1303_v7;
          let _1304_v8;
          let _nw227 = new _module.C1();
          _nw227.__ctor(p0);
          _1304_v8 = _nw227;
          let _1305_v9;
          _1305_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_this.f30, _this.f30)).length),(_1295_v0).isLessThanOrEqualTo(_1295_v0));
          let _1306_v10;
          _1306_v10 = _dafny.Seq.of((_1304_v8).f40);
          _1305_v9 = (_1305_v9).update((_1297_i0).multipliedBy(new BigNumber((_1306_v10).length)), (p0) === (p0));
          let _rhs221 = _module.__default.fm36(globalState);
          let _rhs222 = p0;
          let _lhs208 = globalState;
          let _lhs209 = globalState;
          _lhs208.f21 = _rhs221;
          _lhs209.f13 = _rhs222;
        } else {
          let _1307_v11;
          _1307_v11 = _dafny.Seq.of(p0);
          (globalState).f21 = (new BigNumber((_1307_v11).length)).multipliedBy(_1295_v0);
          r1 = !(true);
          let _1308_v12;
          let _nw228 = Array((new BigNumber(10)).toNumber()).fill(false);
          _1308_v12 = _nw228;
          let _index197 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_1308_v12).length));
          (_1308_v12)[_index197] = !(p0);
          let _index198 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_1308_v12).length));
          (_1308_v12)[_index198] = !((_this).f29);
          (globalState).f18 = _dafny.Seq.IsProperPrefixOf(_this.f30, _dafny.Seq.Create(_module.__default.abs(new BigNumber(207)), function (_1309_i1) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          }));
          let _1310_v13;
          let _nw229 = new _module.C5();
          _nw229.__ctor(true, _this.f30);
          _1310_v13 = _nw229;
        }
        (globalState).f14 = _1295_v0;
      }
      let _1311_v14;
      _1311_v14 = _module.D6.create_DC23((_this).f29, _1295_v0);
      let _1312_v15;
      _1312_v15 = _dafny.Set.fromElements((_this).f29);
      let _1313_v16;
      _1313_v16 = _module.D11.create_DC36(_1311_v14, _1312_v15, (_this).f29);
      let _1314_v17;
      let _nw230 = Array((_dafny.ONE).toNumber());
      _nw230[(_dafny.ZERO).toNumber()] = (_1313_v16).dtor_cf61;
      _1314_v17 = _nw230;
      let _1315_v18;
      _1315_v18 = _module.D0.create_DC1(_1314_v17, _1295_v0, _1295_v0);
      let _1316_v19;
      let _nw231 = Array((new BigNumber(11)).toNumber());
      _nw231[(_dafny.ZERO).toNumber()] = (((_this).fm2(p0, globalState)) ? (_1314_v17) : (_1314_v17));
      _nw231[(_dafny.ONE).toNumber()] = _1314_v17;
      _nw231[(new BigNumber(2)).toNumber()] = _1314_v17;
      _nw231[(new BigNumber(3)).toNumber()] = (((_this).f29) ? (_1314_v17) : (_1314_v17));
      _nw231[(new BigNumber(4)).toNumber()] = _1314_v17;
      _nw231[(new BigNumber(5)).toNumber()] = _1314_v17;
      _nw231[(new BigNumber(6)).toNumber()] = _1314_v17;
      _nw231[(new BigNumber(7)).toNumber()] = _1314_v17;
      _nw231[(new BigNumber(8)).toNumber()] = (_1315_v18).dtor_cf1;
      _nw231[(new BigNumber(9)).toNumber()] = _1314_v17;
      _nw231[(new BigNumber(10)).toNumber()] = _1314_v17;
      _1316_v19 = _nw231;
      let _index199 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_1316_v19).length));
      (_1316_v19)[_index199] = _1314_v17;
      let _index200 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_1316_v19).length));
      (_1316_v19)[_index200] = _1314_v17;
      let _1317_v20;
      _1317_v20 = _dafny.Set.fromElements(_1295_v0, _1295_v0);
      let _1318_v21;
      _1318_v21 = _dafny.Seq.of(p0, (_1317_v20).IsSubsetOf(_1317_v20), (_this).f29);
      (globalState).f13 = (_1318_v21)[_module.__default.safeIndex(_1295_v0, new BigNumber((_1318_v21).length))];
      for (const _guard_loop_8 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1314_v17).length))) {
        let _1319_i2 = _guard_loop_8;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1319_i2)) && ((_1319_i2).isLessThan(new BigNumber((_1314_v17).length))))) {
          (_1314_v17)[(_1319_i2)] = p0;
        }
      }
      let _1320_v22;
      _1320_v22 = _dafny.Map.Empty.slice().updateUnsafe(!(p0) || (false),_1295_v0);
      let _1321_v23;
      _1321_v23 = _dafny.Seq.of(_1295_v0, new BigNumber((_dafny.Seq.UnicodeFromString("kyov")).length));
      _1320_v22 = (_1320_v22).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1321_v23).length)));
      r1 = (_this).f29;
      let _1322_v24;
      _1322_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(99),_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(112)), ((_1323_v21) => function (_1324_i3) {
        return _1323_v21;
      })(_1318_v21)), _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(p0, (_this).f29)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(112)), ((_1325_v21) => function (_1326_i3) {
        return _1325_v21;
      })(_1318_v21))).length)), _dafny.Seq.of(p0)));
      let _1327_v25;
      _1327_v25 = _dafny.Seq.of(_module.__default.fm35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(948)), function (_1328_i4) {
        return _module.D2.create_DC8(new _dafny.CodePoint('d'.codePointAt(0)), (_this).f29);
      }), _1295_v0, globalState), _1318_v21);
      r0 = _dafny.Seq.Concat((((_1322_v24).contains(_1295_v0)) ? ((_1322_v24).get(_1295_v0)) : (_1327_v25)), _dafny.Seq.of(_1318_v21, _1318_v21));
      r1 = p0;
      return [r0, r1];
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let _1329_v0;
      _1329_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,!((_this).f29));
      (globalState).f15 = (((_1329_v0).contains((_this).f29)) ? ((_1329_v0).get((_this).f29)) : (((p0) ? (true) : ((_this).f29))));
      let _1330_v1;
      _1330_v1 = _dafny.MultiSet.fromElements(_1329_v0, _1329_v0, _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f29));
      let _1331_v2;
      _1331_v2 = _dafny.Seq.of((_1330_v1).IsProperSubsetOf(_1330_v1), (_this).f29);
      let _1332_v3;
      _1332_v3 = _module.D9.create_DC28(_1331_v2, p0, (_dafny.ZERO).minus(p2), !(true), (_this).f29);
      let _1333_v4;
      _1333_v4 = _dafny.Seq.of(_dafny.Seq.update(_1331_v2, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("kgmvdm")).length), new BigNumber((_1331_v2).length)), p0), _dafny.Seq.of(p0));
      _1331_v2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1331_v2, _dafny.Seq.update((_1332_v3).dtor_cf41, _module.__default.safeIndex(p2, new BigNumber(((_1332_v3).dtor_cf41).length)), p0)), (_1333_v4)[_module.__default.safeIndex(p2, new BigNumber((_1333_v4).length))]);
      let _hi8 = p2;
      for (let _1334_i0 = (p2).plus(p2); _1334_i0.isLessThan(_hi8); _1334_i0 = _1334_i0.plus(_dafny.ONE)) {
        _1331_v2 = _1331_v2;
        let _1335_v5;
        let _nw232 = Array((new BigNumber(7)).toNumber()).fill(false);
        _1335_v5 = _nw232;
        let _index201 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_1335_v5).length));
        (_1335_v5)[_index201] = true;
        let _index202 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1335_v5).length));
        (_1335_v5)[_index202] = !((_this).f29) || (p0);
        let _1336_v6;
        _1336_v6 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p2),(_this).f29);
        let _index203 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_1335_v5).length));
        let _index204 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1335_v5).length));
        let _rhs223 = p0;
        let _rhs224 = ((!(new BigNumber(40)).isEqualTo(new BigNumber((_1336_v6).length))) ? (p0) : (_module.__default.fm6(p2, globalState)));
        let _rhs225 = p0;
        let _rhs226 = _module.__default.fm36(globalState);
        let _rhs227 = p0;
        let _lhs210 = globalState;
        let _lhs211 = globalState;
        let _lhs212 = _1335_v5;
        let _lhs213 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_1335_v5).length));
        let _lhs214 = globalState;
        let _lhs215 = _1335_v5;
        let _lhs216 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1335_v5).length));
        _lhs210.f15 = _rhs223;
        _lhs211.f18 = _rhs224;
        _lhs212[_lhs213] = _rhs225;
        _lhs214.f2 = _rhs226;
        _lhs215[_lhs216] = _rhs227;
        let _1337_v7;
        _1337_v7 = _dafny.Seq.of(_this.f30, _dafny.Seq.Create(_module.__default.abs(new BigNumber(814)), function (_1338_i1) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }));
        let _1339_v8;
        _1339_v8 = _dafny.Seq.of(_1334_i0, (_dafny.ZERO).minus(p2), _1334_i0, p2, _1334_i0);
        let _1340_v9;
        let _1341_v10;
        let _out18;
        let _out19;
        let _outcollector9 = (_this).m5((_1337_v7)[_module.__default.safeIndex(_1334_i0, new BigNumber((_1337_v7).length))], ((!(p0)) ? ((_1335_v5)[_module.__default.safeIndex(new BigNumber(971), new BigNumber((_1335_v5).length))]) : (p0)), (new BigNumber((_1339_v8).length)).minus(_module.__default.fm36(globalState)), globalState);
        _out18 = _outcollector9[0];
        _out19 = _outcollector9[1];
        _1340_v9 = _out18;
        _1341_v10 = _out19;
        let _1342_v11;
        let _nw233 = new _module.C0();
        _nw233.__ctor(false);
        _1342_v11 = _nw233;
        let _1343_v12;
        _1343_v12 = _dafny.Seq.of(_1342_v11, _1342_v11, _1342_v11);
        let _1344_v13;
        _1344_v13 = _module.D9.create_DC29(new BigNumber((_1343_v12).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,_1334_i0)).length), (_this).f29);
        _1329_v0 = (_1329_v0).update((_1344_v13).dtor_cf48, (_1342_v11).fm11(_1340_v9, globalState));
      }
      if ((_this).f29) {
        let _1345_v14;
        let _nw234 = Array((new BigNumber(17)).toNumber()).fill(false);
        _1345_v14 = _nw234;
        let _rhs228 = (_dafny.ZERO).minus(p2);
        let _rhs229 = _1345_v14;
        let _rhs230 = _1345_v14;
        let _lhs217 = globalState;
        _lhs217.f2 = _rhs228;
        _1345_v14 = _rhs229;
        _1345_v14 = _rhs230;
        let _1346_v15;
        let _nw235 = new _module.C1();
        _nw235.__ctor((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-156)), function (_1347_i2) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        })).length)).isLessThanOrEqualTo(p2));
        _1346_v15 = _nw235;
        let _1348_v16;
        _1348_v16 = _dafny.MultiSet.fromElements(p2, p2);
        let _1349_v17;
        _1349_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(495),_1348_v16);
        let _1350_v18;
        _1350_v18 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        let _1351_v19;
        _1351_v19 = _dafny.Seq.of(p2, p2, new BigNumber(((((_1349_v17).contains(new BigNumber((_1350_v18).length))) ? ((_1349_v17).get(new BigNumber((_1350_v18).length))) : (_dafny.MultiSet.fromElements(_module.__default.fm36(globalState))))).cardinality()));
        if (_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1351_v19, _1351_v19), _1351_v19)) {
          let _1352_v20;
          let _nw236 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _1352_v20 = _nw236;
          let _index205 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_1352_v20).length));
          (_1352_v20)[_index205] = p2;
          let _index206 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_1352_v20).length));
          (_1352_v20)[_index206] = p2;
          let _1353_v21;
          _1353_v21 = _dafny.Set.fromElements(_1352_v20, _1352_v20);
          let _1354_v22;
          let _nw237 = new _module.C2();
          _nw237.__ctor((_1352_v20)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_1352_v20).length))], _1353_v21, (_1346_v15).f40, _this.f30);
          _1354_v22 = _nw237;
          let _1355_v23;
          _1355_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1354_v22,_1329_v0);
          let _1356_v24;
          _1356_v24 = _1329_v0;
          let _1357_v25;
          let _nw238 = Array((new BigNumber(14)).toNumber());
          _nw238[(_dafny.ZERO).toNumber()] = _1329_v0;
          _nw238[(_dafny.ONE).toNumber()] = _1329_v0;
          _nw238[(new BigNumber(2)).toNumber()] = _1329_v0;
          _nw238[(new BigNumber(3)).toNumber()] = _1329_v0;
          _nw238[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1346_v15).f40,(_1346_v15).f40);
          _nw238[(new BigNumber(5)).toNumber()] = (((_1355_v23).contains(_1354_v22)) ? ((_1355_v23).get(_1354_v22)) : (_1329_v0));
          _nw238[(new BigNumber(6)).toNumber()] = (_1329_v0).Merge(_1329_v0);
          _nw238[(new BigNumber(7)).toNumber()] = ((_1329_v0).update((_this).f29, (_this).f29)).Merge(_1329_v0);
          _nw238[(new BigNumber(8)).toNumber()] = _1329_v0;
          _nw238[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1346_v15).f40,(_this).f29);
          _nw238[(new BigNumber(10)).toNumber()] = (_1356_v24);
          _nw238[(new BigNumber(11)).toNumber()] = _1329_v0;
          _nw238[(new BigNumber(12)).toNumber()] = _1329_v0;
          _nw238[(new BigNumber(13)).toNumber()] = _1329_v0;
          _1357_v25 = _nw238;
          let _1358_v26;
          _1358_v26 = new _dafny.CodePoint('k'.codePointAt(0));
          let _1359_v27;
          _1359_v27 = _module.D2.create_DC8(_1358_v26, (_this).f29);
          let _1360_v28;
          _1360_v28 = _module.D5.create_DC19((_1352_v20)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_1352_v20).length))], new BigNumber((_1331_v2).length), (_1359_v27).dtor_cf12);
          let _rhs231 = _1357_v25;
          let _rhs232 = _dafny.Seq.update(_1351_v19, _module.__default.safeIndex((_1360_v28).dtor_cf28, new BigNumber((_1351_v19).length)), _module.__default.safeDivisionInt((_1354_v22).f38, (_1354_v22).f38));
          let _rhs233 = p2;
          let _rhs234 = _dafny.Seq.contains(_1331_v2, _dafny.areEqual(_this.f30, _this.f30));
          let _rhs235 = ((_1354_v22).f38).multipliedBy((_1354_v22).f38);
          let _lhs218 = globalState;
          let _lhs219 = globalState;
          let _lhs220 = globalState;
          let _lhs221 = globalState;
          _1357_v25 = _rhs231;
          _lhs218.f7 = _rhs232;
          _lhs219.f21 = _rhs233;
          _lhs220.f13 = _rhs234;
          _lhs221.f1 = _rhs235;
          let _1361_v29;
          let _nw239 = new _module.C0();
          _nw239.__ctor((_this).f29);
          _1361_v29 = _nw239;
          let _1362_v30;
          _1362_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1361_v29.f37,new BigNumber(994));
          (globalState).f21 = _module.__default.safeModuloInt(new BigNumber((_1362_v30).length), ((_1354_v22).f38).minus((_1361_v29).fm12(globalState)));
          let _1363_v31;
          _1363_v31 = _dafny.Set.fromElements((new BigNumber(-182)).plus(p2), (p2).minus(new BigNumber((_this.f30).length)), p2, (_1352_v20)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_1352_v20).length))]);
          _1363_v31 = _1363_v31;
        } else {
          (globalState).f1 = p2;
          let _1364_v32;
          let _nw240 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
          _1364_v32 = _nw240;
          let _1365_v33;
          let _nw241 = Array((new BigNumber(9)).toNumber());
          _nw241[(_dafny.ZERO).toNumber()] = _1364_v32;
          _nw241[(_dafny.ONE).toNumber()] = _1364_v32;
          _nw241[(new BigNumber(2)).toNumber()] = _1364_v32;
          _nw241[(new BigNumber(3)).toNumber()] = _1364_v32;
          _nw241[(new BigNumber(4)).toNumber()] = _1364_v32;
          _nw241[(new BigNumber(5)).toNumber()] = _1364_v32;
          _nw241[(new BigNumber(6)).toNumber()] = _1364_v32;
          _nw241[(new BigNumber(7)).toNumber()] = _1364_v32;
          _nw241[(new BigNumber(8)).toNumber()] = _1364_v32;
          _1365_v33 = _nw241;
          let _index207 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_1365_v33).length));
          (_1365_v33)[_index207] = _1364_v32;
          let _1366_v34;
          _1366_v34 = _dafny.Set.fromElements(p2, p2, p2, new BigNumber((_dafny.Seq.of(p0)).length), new BigNumber((_1331_v2).length));
          let _index208 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_1365_v33).length));
          let _rhs236 = _1364_v32;
          let _rhs237 = _module.__default.safeModuloInt(new BigNumber((_this.f30).length), p2);
          let _rhs238 = (new BigNumber((_1366_v34).length)).multipliedBy(((!(p0)) ? (p2) : (p2)));
          let _rhs239 = (_this).f29;
          let _lhs222 = _1365_v33;
          let _lhs223 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_1365_v33).length));
          let _lhs224 = globalState;
          let _lhs225 = globalState;
          let _lhs226 = globalState;
          _lhs222[_lhs223] = _rhs236;
          _lhs224.f21 = _rhs237;
          _lhs225.f21 = _rhs238;
          _lhs226.f28 = _rhs239;
          _1350_v18 = (_1350_v18).update(new BigNumber((_1331_v2).length), (_this).f29);
          (_this).f30 = _this.f30;
          let _1367_v35;
          _1367_v35 = new _dafny.CodePoint('e'.codePointAt(0));
          let _rhs240 = _1345_v14;
          let _rhs241 = _1367_v35;
          let _rhs242 = _1350_v18;
          _1345_v14 = _rhs240;
          _1367_v35 = _rhs241;
          _1350_v18 = _rhs242;
        }
        let _1368_v36;
        let _nw242 = Array((new BigNumber(23)).toNumber()).fill([]);
        _1368_v36 = _nw242;
        let _index209 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_1368_v36).length));
        (_1368_v36)[_index209] = _1345_v14;
        let _index210 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_1368_v36).length));
        (_1368_v36)[_index210] = _1345_v14;
        let _1369_v37;
        _1369_v37 = _module.D3.create_DC10(_1348_v16);
        if ((_1348_v16).IsSubsetOf(((_1369_v37).dtor_cf18).Union(_dafny.MultiSet.fromElements(p2)))) {
          let _1370_v38;
          _1370_v38 = _module.D12.create_DC39(new BigNumber((_1329_v0).length), (_1346_v15).f40, p2);
          let _1371_v39;
          _1371_v39 = _module.D6.create_DC23(p0, p2);
          let _pat_let_tv28 = _1346_v15;
          let _rhs243 = (function (_pat_let31_0) {
            return function (_1372_dt__update__tmp_h0) {
              return function (_pat_let32_0) {
                return function (_1373_dt__update_hcf36_h0) {
                  return _module.D6.create_DC23(_1373_dt__update_hcf36_h0, (_1372_dt__update__tmp_h0).dtor_cf37);
                }(_pat_let32_0);
              }((_pat_let_tv28).f40);
            }(_pat_let31_0);
          }(_1371_v39)).dtor_cf37;
          let _rhs244 = !(!(p0));
          let _rhs245 = _1370_v38;
          let _rhs246 = p2;
          let _rhs247 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((p2).multipliedBy(new BigNumber(-620)), new BigNumber(-108)));
          let _lhs227 = globalState;
          let _lhs228 = globalState;
          let _lhs229 = globalState;
          _lhs227.f19 = _rhs243;
          r1 = _rhs244;
          _1370_v38 = _rhs245;
          _lhs228.f19 = _rhs246;
          _lhs229.f14 = _rhs247;
          let _1374_v40;
          let _init35 = ((_1375_v15, _1376_p2) => function (_1377_i3) {
            return (((_1375_v15).f40) ? (_module.D11.create_DC37(_1376_p2, (_1375_v15).f40)) : (_module.D11.create_DC37(new BigNumber((_this.f30).length), false)));
          })(_1346_v15, p2);
          let _nw243 = Array((new BigNumber(25)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw243.length); _i0_35++) {
            _nw243[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1374_v40 = _nw243;
          let _1378_v41;
          _1378_v41 = _module.D11.create_DC37(p2, (_this).f29);
          let _index211 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_1374_v40).length));
          (_1374_v40)[_index211] = _1378_v41;
          let _arr0 = (_1368_v36)[_module.__default.safeIndex(new BigNumber(196), new BigNumber((_1368_v36).length))];
          let _index212 = _module.__default.safeIndex(new BigNumber(590), new BigNumber(((_1368_v36)[_module.__default.safeIndex(new BigNumber(196), new BigNumber((_1368_v36).length))]).length));
          _arr0[_index212] = (_1346_v15).f40;
          let _index213 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_1374_v40).length));
          let _arr1 = (_1368_v36)[_module.__default.safeIndex(new BigNumber(196), new BigNumber((_1368_v36).length))];
          let _index214 = _module.__default.safeIndex(new BigNumber(590), new BigNumber(((_1368_v36)[_module.__default.safeIndex(new BigNumber(196), new BigNumber((_1368_v36).length))]).length));
          let _rhs248 = _1378_v41;
          let _rhs249 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1351_v19, _dafny.Seq.of(new BigNumber((_1329_v0).length))), _1351_v19);
          let _rhs250 = _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of((_1346_v15).f40, p0));
          let _lhs230 = _1374_v40;
          let _lhs231 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_1374_v40).length));
          let _lhs232 = (_1368_v36)[_module.__default.safeIndex(new BigNumber(196), new BigNumber((_1368_v36).length))];
          let _lhs233 = _module.__default.safeIndex(new BigNumber(590), new BigNumber(((_1368_v36)[_module.__default.safeIndex(new BigNumber(196), new BigNumber((_1368_v36).length))]).length));
          _lhs230[_lhs231] = _rhs248;
          _lhs232[_lhs233] = _rhs249;
          _1331_v2 = _rhs250;
          (globalState).f28 = (_1346_v15).f40;
          let _1379_v42;
          _1379_v42 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1380_v44;
          _1380_v44 = _dafny.Set.fromElements(p2);
          let _1381_v45;
          _1381_v45 = _dafny.Seq.of(function () {
            let _coll47 = new _dafny.Set();
            for (const _compr_47 of _dafny.IntegerRange(new BigNumber(644), new BigNumber(-791))) {
              let _1382_v43 = _compr_47;
              if (((new BigNumber(644)).isLessThanOrEqualTo(_1382_v43)) && ((_1382_v43).isLessThan(new BigNumber(-791)))) {
                _coll47.add((_1382_v43).plus(p2));
              }
            }
            return _coll47;
          }(), _1380_v44, _1380_v44, _1380_v44);
          _1379_v42 = _module.__default.fm42((_module.__default.fm39(_1379_v42, new BigNumber(883), globalState)).equals(_module.__default.fm39(_1379_v42, (_dafny.ZERO).minus(p2), globalState)), _dafny.Seq.IsPrefixOf(_1381_v45, _1381_v45), globalState);
          let _1383_v46;
          let _init36 = function (_1384_i4) {
            return (_1384_i4).minus(new BigNumber(142));
          };
          let _nw244 = Array((new BigNumber(22)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw244.length); _i0_36++) {
            _nw244[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1383_v46 = _nw244;
          let _1385_v47;
          _1385_v47 = _module.D12.create_DC38(_1383_v46);
          (globalState).f17 = (_1385_v47).dtor_cf64;
        } else {
          (globalState).f20 = _this.f30;
          (globalState).f15 = ((p0) && ((_1346_v15).f40)) || ((_this).f29);
          let _1386_v48;
          _1386_v48 = _dafny.Map.Empty.slice().updateUnsafe(!(((_dafny.ZERO).minus((_dafny.ZERO).minus(p2))).isLessThan(p2)),(p2).plus(p2));
          let _1387_v49;
          _1387_v49 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          _1386_v48 = (_1386_v48).update(!(p0), (((_1387_v49).contains(_module.__default.fm9(globalState))) ? ((_1387_v49).get(_module.__default.fm9(globalState))) : (p2)));
          let _1388_v50;
          let _nw245 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _1388_v50 = _nw245;
          let _1389_v51;
          _1389_v51 = _dafny.Set.fromElements(_1388_v50, _1388_v50);
          let _1390_v52;
          let _nw246 = new _module.C3();
          _nw246.__ctor(!(_1350_v18).equals(_1350_v18), (p2).plus(new BigNumber((_dafny.Seq.UnicodeFromString("wgqvib")).length)), (_1389_v51).IsProperSubsetOf(_1389_v51), _this.f30);
          _1390_v52 = _nw246;
          _1390_v52 = _1390_v52;
          _1388_v50 = _1388_v50;
        }
      } else {
        let _1391_v53;
        _1391_v53 = _dafny.MultiSet.fromElements(_module.D11.create_DC37(p2, p0), _module.D11.create_DC37(p2, p0));
        let _1392_v54;
        _1392_v54 = _module.D11.create_DC37(p2, p0);
        _1391_v53 = (_1391_v53).Difference(_dafny.MultiSet.fromElements(_1392_v54));
        let _1393_v55;
        let _nw247 = new _module.C3();
        _nw247.__ctor(((p0) ? (p0) : (!(p0))), p2, _dafny.Seq.IsPrefixOf(_1331_v2, _1331_v2), _this.f30);
        _1393_v55 = _nw247;
        (globalState).f19 = (_1393_v55).f36;
        let _1394_v56;
        let _init37 = ((_1395_p2) => function (_1396_i5) {
          return (_1396_i5).plus(_1395_p2);
        })(p2);
        let _nw248 = Array((new BigNumber(12)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw248.length); _i0_37++) {
          _nw248[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1394_v56 = _nw248;
        let _index215 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_1394_v56).length));
        (_1394_v56)[_index215] = p2;
        let _1397_v57;
        let _nw249 = Array((new BigNumber(28)).toNumber()).fill([]);
        _1397_v57 = _nw249;
        let _1398_v58;
        let _nw250 = Array((new BigNumber(7)).toNumber()).fill(false);
        _1398_v58 = _nw250;
        let _index216 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1397_v57).length));
        (_1397_v57)[_index216] = _1398_v58;
        let _1399_v59;
        _1399_v59 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1393_v55).f36);
        let _1400_v60;
        _1400_v60 = _module.D0.create_DC1(_1398_v58, (_1393_v55).f36, (_1332_v3).dtor_cf43);
        let _index217 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_1394_v56).length));
        let _index218 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1397_v57).length));
        let _rhs251 = ((p2).minus((_dafny.ZERO).minus((_1393_v55).f36))).minus((((_1399_v59).contains((_1393_v55).f36)) ? ((_1399_v59).get((_1393_v55).f36)) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_1393_v55).f36)).length)))));
        let _rhs252 = (_1393_v55).f36;
        let _rhs253 = (_1400_v60).dtor_cf1;
        let _lhs234 = _1394_v56;
        let _lhs235 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_1394_v56).length));
        let _lhs236 = globalState;
        let _lhs237 = _1397_v57;
        let _lhs238 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1397_v57).length));
        _lhs234[_lhs235] = _rhs251;
        _lhs236.f19 = _rhs252;
        _lhs237[_lhs238] = _rhs253;
        let _1401_v61;
        let _nw251 = new _module.C4();
        _nw251.__ctor(p0, p2);
        _1401_v61 = _nw251;
        let _1402_v62;
        let _nw252 = Array((new BigNumber(20)).toNumber());
        _nw252[(_dafny.ZERO).toNumber()] = _1401_v61;
        _nw252[(_dafny.ONE).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(2)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(3)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(4)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(5)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(6)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(7)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(8)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(9)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(10)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(11)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(12)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(13)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(14)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(15)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(16)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(17)).toNumber()] = _1401_v61;
        _nw252[(new BigNumber(18)).toNumber()] = (((_1393_v55).f35) ? (_1401_v61) : (_1401_v61));
        _nw252[(new BigNumber(19)).toNumber()] = _1401_v61;
        _1402_v62 = _nw252;
        let _1403_v63;
        _1403_v63 = _module.D13.create_DC40(_1401_v61);
        let _1404_v64;
        _1404_v64 = _dafny.Seq.of(_1401_v61);
        let _1405_v65;
        _1405_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_1404_v64)[_module.__default.safeIndex(new BigNumber(826), new BigNumber((_1404_v64).length))]);
        let _1406_v66;
        _1406_v66 = _module.D3.create_DC11((_1393_v55).f35, (_1401_v61).f33);
        let _nw253 = Array((new BigNumber(18)).toNumber());
        _nw253[(_dafny.ZERO).toNumber()] = _1401_v61;
        _nw253[(_dafny.ONE).toNumber()] = (((_1393_v55).f35) ? ((_1403_v63).dtor_cf68) : (_1401_v61));
        _nw253[(new BigNumber(2)).toNumber()] = _1401_v61;
        _nw253[(new BigNumber(3)).toNumber()] = _1401_v61;
        _nw253[(new BigNumber(4)).toNumber()] = (((_1405_v65).contains((_1401_v61).f33)) ? ((_1405_v65).get((_1401_v61).f33)) : (_1401_v61));
        _nw253[(new BigNumber(5)).toNumber()] = _1401_v61;
        _nw253[(new BigNumber(6)).toNumber()] = _1401_v61;
        _nw253[(new BigNumber(7)).toNumber()] = _1401_v61;
        _nw253[(new BigNumber(8)).toNumber()] = _1401_v61;
        _nw253[(new BigNumber(9)).toNumber()] = _1401_v61;
        _nw253[(new BigNumber(10)).toNumber()] = ((p0) ? (_1401_v61) : (_1401_v61));
        _nw253[(new BigNumber(11)).toNumber()] = (((_this).f29) ? (_1401_v61) : (_1401_v61));
        _nw253[(new BigNumber(12)).toNumber()] = _1401_v61;
        _nw253[(new BigNumber(13)).toNumber()] = _1401_v61;
        _nw253[(new BigNumber(14)).toNumber()] = _1401_v61;
        _nw253[(new BigNumber(15)).toNumber()] = _1401_v61;
        _nw253[(new BigNumber(16)).toNumber()] = (((_1406_v66).dtor_cf20) ? (_1401_v61) : (_1401_v61));
        _nw253[(new BigNumber(17)).toNumber()] = _1401_v61;
        _1402_v62 = _nw253;
      }
      let _1407_v68;
      let _init38 = ((_1408_p0) => function (_1409_i7) {
        return (_dafny.MultiSet.fromElements(function () {
          let _coll48 = new _dafny.Map();
          for (const _compr_48 of (_this.f30).Elements) {
            let _1410_v67 = _compr_48;
            if (_dafny.Seq.contains(_this.f30, _1410_v67)) {
              _coll48.push([_1410_v67,false]);
            }
          }
          return _coll48;
        }())).IsSubsetOf(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),_1408_p0)));
      })(p0);
      let _nw254 = Array((new BigNumber(28)).toNumber());
      for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw254.length); _i0_38++) {
        _nw254[_i0_38] = _init38(new BigNumber(_i0_38));
      }
      _1407_v68 = _nw254;
      for (const _guard_loop_9 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1407_v68).length))) {
        let _1411_i6 = _guard_loop_9;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1411_i6)) && ((_1411_i6).isLessThan(new BigNumber((_1407_v68).length))))) {
          (_1407_v68)[(_1411_i6)] = (_module.D11.create_DC35(_this.f30, p0)).dtor_cf58;
        }
      }
      let _1412_v69;
      let _nw255 = Array((new BigNumber(14)).toNumber());
      _nw255[(_dafny.ZERO).toNumber()] = _this.f30;
      _nw255[(_dafny.ONE).toNumber()] = _this.f30;
      _nw255[(new BigNumber(2)).toNumber()] = _this.f30;
      _nw255[(new BigNumber(3)).toNumber()] = _this.f30;
      _nw255[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_this.f30, _this.f30);
      _nw255[(new BigNumber(5)).toNumber()] = _this.f30;
      _nw255[(new BigNumber(6)).toNumber()] = _this.f30;
      _nw255[(new BigNumber(7)).toNumber()] = _this.f30;
      _nw255[(new BigNumber(8)).toNumber()] = _this.f30;
      _nw255[(new BigNumber(9)).toNumber()] = _this.f30;
      _nw255[(new BigNumber(10)).toNumber()] = _this.f30;
      _nw255[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_this.f30, _dafny.Seq.UnicodeFromString("nnvokyy"));
      _nw255[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uqootstpb"), _this.f30);
      _nw255[(new BigNumber(13)).toNumber()] = (((_this).f29) ? (_this.f30) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-207)), function (_1413_i9) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })));
      _1412_v69 = _nw255;
      for (const _guard_loop_10 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1412_v69).length))) {
        let _1414_i8 = _guard_loop_10;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1414_i8)) && ((_1414_i8).isLessThan(new BigNumber((_1412_v69).length))))) {
          (_1412_v69)[(_1414_i8)] = _dafny.Seq.Concat(_this.f30, _this.f30);
        }
      }
      let _1415_v70;
      _1415_v70 = _dafny.Seq.of(new BigNumber(542));
      r0 = _module.__default.fm34(p2, _module.__default.fm36(globalState), p0, _dafny.Seq.IsPrefixOf(_1415_v70, _1415_v70), globalState);
      r1 = p0;
      return [r0, r1];
    }
    m5(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      if ((_this).f29) {
        let _1416_v0;
        _1416_v0 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(p1)).length));
        if (((_1416_v0).Union(_1416_v0)).equals(_1416_v0)) {
          let _1417_v1;
          let _nw256 = Array((new BigNumber(13)).toNumber()).fill(_module.D2.Default());
          _1417_v1 = _nw256;
          let _1418_v2;
          _1418_v2 = new _dafny.CodePoint('q'.codePointAt(0));
          let _1419_v3;
          _1419_v3 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0)), _1418_v2);
          let _1420_v4;
          _1420_v4 = _module.D2.create_DC6(_1419_v3);
          let _index219 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_1417_v1).length));
          (_1417_v1)[_index219] = _1420_v4;
          let _index220 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_1417_v1).length));
          (_1417_v1)[_index220] = _1420_v4;
          let _1421_v5;
          let _nw257 = new _module.C0();
          _nw257.__ctor(p1);
          _1421_v5 = _nw257;
          let _1422_v6;
          _1422_v6 = _dafny.Seq.of(!(_1421_v5.f37) || ((_this).f29));
          _1422_v6 = _dafny.Seq.update(_1422_v6, _module.__default.safeIndex(p2, new BigNumber((_1422_v6).length)), _1421_v5.f37);
          (globalState).f21 = p2;
          _1418_v2 = _module.__default.fm42(!(p1) || (p1), (_this).f29, globalState);
        } else {
          let _1423_v7;
          _1423_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm48(globalState)).length),(_dafny.ZERO).minus(p2));
          let _1424_v8;
          _1424_v8 = _1423_v7;
          let _1425_v9;
          _1425_v9 = _dafny.Seq.of((_1424_v8), _1423_v7);
          let _1426_v10;
          _1426_v10 = new _dafny.CodePoint('k'.codePointAt(0));
          _1423_v7 = ((_1425_v9)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("rksaphr"), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.UnicodeFromString("rksaphr")).length)), _1426_v10)).length), new BigNumber((_1425_v9).length))]).Merge((_dafny.Map.Empty.slice().updateUnsafe(p2,p2)).Merge((_1423_v7).update(p2, p2)));
          (globalState).f13 = (!((_this).f29)) && ((_this).f29);
          (globalState).f2 = ((p1) ? (new BigNumber(605)) : (p2));
          _1426_v10 = _1426_v10;
          (globalState).f14 = p2;
        }
        let _1427_v11;
        _1427_v11 = _dafny.MultiSet.fromElements(p2);
        _1427_v11 = (_1427_v11).Intersect(_1427_v11);
        let _1428_v12;
        _1428_v12 = _dafny.Set.fromElements(p1);
        let _1429_v13;
        _1429_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1428_v12).Union(_1428_v12)).length),(p2).isEqualTo(p2));
        _1429_v13 = _1429_v13;
        let _1430_v14;
        _1430_v14 = _module.D15.create_DC43(_1429_v13);
        let _rhs254 = (_1430_v14).dtor_cf70;
        let _rhs255 = (_this).f29;
        let _lhs239 = globalState;
        _1429_v13 = _rhs254;
        _lhs239.f28 = _rhs255;
        let _1431_v15;
        let _nw258 = new _module.C0();
        _nw258.__ctor(true);
        _1431_v15 = _nw258;
      } else {
        let _1432_v16;
        _1432_v16 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1433_v17;
        _1433_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-76),!((_this).f29));
        let _1434_v18;
        _1434_v18 = _module.D10.create_DC32(_1432_v16, p1);
        let _1435_v19;
        let _nw259 = Array((new BigNumber(13)).toNumber());
        _nw259[(_dafny.ZERO).toNumber()] = _module.__default.fm49(_1432_v16, globalState);
        _nw259[(_dafny.ONE).toNumber()] = _module.D10.create_DC32(new _dafny.CodePoint('j'.codePointAt(0)), (((_1433_v17).contains(p2)) ? ((_1433_v17).get(p2)) : ((_this).f29)));
        _nw259[(new BigNumber(2)).toNumber()] = _1434_v18;
        _nw259[(new BigNumber(3)).toNumber()] = _1434_v18;
        _nw259[(new BigNumber(4)).toNumber()] = _1434_v18;
        _nw259[(new BigNumber(5)).toNumber()] = _1434_v18;
        _nw259[(new BigNumber(6)).toNumber()] = _1434_v18;
        _nw259[(new BigNumber(7)).toNumber()] = _1434_v18;
        _nw259[(new BigNumber(8)).toNumber()] = _1434_v18;
        _nw259[(new BigNumber(9)).toNumber()] = _1434_v18;
        _nw259[(new BigNumber(10)).toNumber()] = _module.D10.create_DC32(_1432_v16, p1);
        _nw259[(new BigNumber(11)).toNumber()] = _module.D10.create_DC32(_1432_v16, p1);
        _nw259[(new BigNumber(12)).toNumber()] = _1434_v18;
        _1435_v19 = _nw259;
        let _1436_v20;
        let _nw260 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _1436_v20 = _nw260;
        let _1437_v21;
        _1437_v21 = _dafny.Seq.of(p1, (_this).f29, p1);
        let _1438_v22;
        _1438_v22 = _module.D15.create_DC44((_dafny.ZERO).minus(p2), p1, _1435_v19, _1436_v20, new BigNumber((_1437_v21).length));
        _1438_v22 = ((false) ? (_1438_v22) : (_1438_v22));
        let _1439_v23;
        _1439_v23 = _dafny.MultiSet.fromElements(p2);
        let _1440_v24;
        _1440_v24 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.of((_this).f29, p1, (_this).f29), _module.__default.fm35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-147)), function (_1441_i0) {
          return _module.D2.create_DC8(new _dafny.CodePoint('n'.codePointAt(0)), (_this).f29);
        }), new BigNumber((_1439_v23).cardinality()), globalState)));
        (globalState).f1 = (((_1440_v24).contains(_dafny.Seq.Concat(_1437_v21, _1437_v21))) ? ((_1440_v24).get(_dafny.Seq.Concat(_1437_v21, _1437_v21))) : (p2));
        let _index221 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1436_v20).length));
        (_1436_v20)[_index221] = _module.__default.fm9(globalState);
        let _index222 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1436_v20).length));
        (_1436_v20)[_index222] = (p2).minus(p2);
        let _1442_v25;
        let _nw261 = new _module.C4();
        _nw261.__ctor(p1, p2);
        _1442_v25 = _nw261;
        (globalState).f28 = (_this).f29;
      }
      let _1443_v26;
      let _nw262 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _1443_v26 = _nw262;
      let _index223 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length));
      (_1443_v26)[_index223] = _module.__default.safeModuloInt(p2, (_dafny.ZERO).minus(p2));
      let _index224 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length));
      (_1443_v26)[_index224] = _module.__default.fm36(globalState);
      let _index225 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length));
      (_1443_v26)[_index225] = _module.__default.safeModuloInt(p2, (_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))]);
      let _hi9 = (p2).plus(p2);
      for (let _1444_i1 = (_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))]; _1444_i1.isLessThan(_hi9); _1444_i1 = _1444_i1.plus(_dafny.ONE)) {
        let _1445_v27;
        let _nw263 = new _module.C1();
        _nw263.__ctor(!((_this).f29));
        _1445_v27 = _nw263;
        let _1446_v28;
        _1446_v28 = _dafny.Seq.of(_this.f30);
        let _1447_v29;
        _1447_v29 = _dafny.Seq.of((_1446_v28)[_module.__default.safeIndex(_1444_i1, new BigNumber((_1446_v28).length))]);
        let _1448_v30;
        _1448_v30 = _dafny.MultiSet.fromElements(p0, _this.f30, p0, p0);
        let _1449_v31;
        _1449_v31 = _dafny.Map.Empty.slice().updateUnsafe(false,true);
        let _1450_v32;
        _1450_v32 = new _dafny.CodePoint('k'.codePointAt(0));
        let _1451_v33;
        _1451_v33 = _dafny.MultiSet.fromElements(_1450_v32);
        let _index226 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length));
        let _index227 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length));
        let _rhs256 = _1444_i1;
        let _rhs257 = !(((_dafny.MultiSet.FromArray(_1447_v29)).update(_this.f30, _module.__default.abs(_1444_i1))).IsSubsetOf(_1448_v30));
        let _rhs258 = _module.__default.fm33(_module.__default.fm33((_1445_v27).f40, (_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))], new BigNumber((_1449_v31).length), globalState), new BigNumber(((_1451_v33).update(_1450_v32, _module.__default.abs(p2))).cardinality()), _1444_i1, globalState);
        let _rhs259 = _1444_i1;
        let _lhs240 = _1443_v26;
        let _lhs241 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length));
        let _lhs242 = globalState;
        let _lhs243 = globalState;
        let _lhs244 = _1443_v26;
        let _lhs245 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length));
        _lhs240[_lhs241] = _rhs256;
        _lhs242.f18 = _rhs257;
        _lhs243.f28 = _rhs258;
        _lhs244[_lhs245] = _rhs259;
        let _1452_v34;
        _1452_v34 = _dafny.Set.fromElements((_this).f29, (_1445_v27).f40, p1);
        let _1453_v35;
        _1453_v35 = _module.D6.create_DC21(_1452_v34);
        let _1454_v36;
        _1454_v36 = _dafny.Seq.of(p1);
        _1453_v35 = _module.__default.fm50((_this).f29, _module.__default.safeDivisionInt(_1444_i1, _1444_i1), _dafny.Seq.Concat(_this.f30, p0), (((_1445_v27).f40) ? (_1454_v36) : (_dafny.Seq.of((_this).f29))), globalState);
        let _1455_v37;
        _1455_v37 = _dafny.Seq.of(_module.__default.fm51(globalState));
        (globalState).f21 = new BigNumber((_1455_v37).length);
      }
      let _1456_v38;
      _1456_v38 = _dafny.Seq.of(true);
      let _1457_v39;
      _1457_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,new BigNumber((_1456_v38).length));
      let _hi10 = p2;
      for (let _1458_i2 = _module.__default.safeModuloInt(p2, new BigNumber((_1457_v39).length)); _1458_i2.isLessThan(_hi10); _1458_i2 = _1458_i2.plus(_dafny.ONE)) {
        let _1459_v40;
        _1459_v40 = _dafny.MultiSet.fromElements(p1);
        let _1460_v41;
        _1460_v41 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1461_v42;
        _1461_v42 = _module.D2.create_DC8(_1460_v41, p1);
        let _1462_v43;
        _1462_v43 = _dafny.Seq.of(_1461_v42, _1461_v42);
        _1459_v40 = (((p1) ? (_dafny.MultiSet.fromElements(true, (_this).f29)) : (_1459_v40))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_module.__default.fm35(_1462_v43, (_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))], globalState), _1456_v38)));
        let _1463_v44;
        _1463_v44 = _module.D9.create_DC29((_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))], _1458_i2, false);
        let _1464_v45;
        _1464_v45 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f29);
        (globalState).f28 = (_1464_v45).contains((_1463_v44).dtor_cf48);
        let _1465_v46;
        _1465_v46 = _dafny.MultiSet.fromElements(_1458_i2, (_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))], _1458_i2, (_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))], p2);
        let _1466_v47;
        _1466_v47 = _module.D3.create_DC10(_1465_v46);
        let _1467_v48;
        _1467_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1458_i2,_1466_v47);
        _1467_v48 = (_1467_v48).update(p2, _1466_v47);
        if (((_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))]).isLessThan((_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))])) {
          (globalState).f21 = _1458_i2;
          (globalState).f28 = (_this).f29;
          let _init39 = ((_1468_v26) => function (_1469_i3) {
            return (_1469_i3).plus((_1468_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1468_v26).length))]);
          })(_1443_v26);
          let _nw264 = Array((new BigNumber(21)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw264.length); _i0_39++) {
            _nw264[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          (globalState).f17 = _nw264;
          _1460_v41 = _1460_v41;
          let _rhs260 = (_1463_v44).dtor_cf46;
          let _rhs261 = _1460_v41;
          let _lhs246 = globalState;
          _lhs246.f19 = _rhs260;
          _1460_v41 = _rhs261;
        } else {
          (globalState).f15 = !(!(p1) || (false)) || ((_this).f29);
          (globalState).f20 = p0;
          (globalState).f14 = p2;
          let _1470_v49;
          let _nw265 = new _module.C0();
          _nw265.__ctor(p1);
          _1470_v49 = _nw265;
          let _rhs262 = _1470_v49;
          let _rhs263 = (_1457_v39).contains((p1) && (_module.__default.fm19(_1460_v41, globalState)));
          let _lhs247 = globalState;
          _1470_v49 = _rhs262;
          _lhs247.f28 = _rhs263;
          let _1471_v50;
          let _nw266 = Array((new BigNumber(28)).toNumber()).fill(false);
          _1471_v50 = _nw266;
          let _index228 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_1471_v50).length));
          (_1471_v50)[_index228] = (p1) && (_1470_v49.f37);
          let _1472_v51;
          _1472_v51 = _dafny.Seq.of((_1459_v40).Intersect(_dafny.MultiSet.fromElements(p1, p1, p1)));
          let _index229 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_1471_v50).length));
          let _rhs264 = _1470_v49.f37;
          let _rhs265 = _1470_v49.f37;
          let _rhs266 = _1472_v51;
          let _rhs267 = p1;
          let _lhs248 = globalState;
          let _lhs249 = _1471_v50;
          let _lhs250 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_1471_v50).length));
          let _lhs251 = globalState;
          _lhs248.f18 = _rhs264;
          _lhs249[_lhs250] = _rhs265;
          _1472_v51 = _rhs266;
          _lhs251.f28 = _rhs267;
        }
      }
      let _1473_v52;
      let _nw267 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
      _1473_v52 = _nw267;
      let _1474_v53;
      _1474_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1456_v38).length),p2);
      let _index230 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_1473_v52).length));
      (_1473_v52)[_index230] = _1474_v53;
      let _1475_v55;
      _1475_v55 = new _dafny.CodePoint('t'.codePointAt(0));
      let _1476_v56;
      _1476_v56 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19(_1475_v55, globalState),p1);
      let _index231 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_1473_v52).length));
      (_1473_v52)[_index231] = function () {
        let _coll49 = new _dafny.Map();
        for (const _compr_49 of (_dafny.MultiSet.fromElements(_module.__default.fm36(globalState), new BigNumber((_1476_v56).length), (_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))], new BigNumber(-57))).Elements) {
          let _1477_v54 = _compr_49;
          if ((_dafny.MultiSet.fromElements(_module.__default.fm36(globalState), new BigNumber((_1476_v56).length), (_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))], new BigNumber(-57))).contains(_1477_v54)) {
            _coll49.push([_module.__default.safeDivisionInt(_1477_v54, p2),_module.__default.safeDivisionInt((_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))], new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(93),p2)).length))]);
          }
        }
        return _coll49;
      }();
      r0 = p1;
      let _1478_v57;
      _1478_v57 = _dafny.MultiSet.fromElements((_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))]);
      r1 = _module.__default.fm34((p2).minus((_1443_v26)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_1443_v26).length))]), (new BigNumber(-80)).plus(p2), (_this).f29, (_1478_v57).IsSubsetOf(_dafny.MultiSet.fromElements(p2, p2)), globalState);
      return [r0, r1];
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f30 = _dafny.Seq.UnicodeFromString("");
      this._f29 = false;
      this._f31 = false;
      this._f32 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f30() {
      let _this = this;
      return _this._f30;
    };
    set f30(value) {
      let _this = this;
      _this._f30 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    __ctor(f31, f32, f29, f30) {
      let _this = this;
      (_this)._f31 = f31;
      (_this)._f32 = f32;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      return;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _1479_v0;
      _1479_v0 = new BigNumber(601);
      let _1480_v1;
      let _nw268 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
      _1480_v1 = _nw268;
      let _1481_v2;
      _1481_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1479_v0,_1480_v1);
      let _1482_v3;
      let _nw269 = Array((new BigNumber(11)).toNumber());
      _nw269[(_dafny.ZERO).toNumber()] = _1479_v0;
      _nw269[(_dafny.ONE).toNumber()] = _1479_v0;
      _nw269[(new BigNumber(2)).toNumber()] = _1479_v0;
      _nw269[(new BigNumber(3)).toNumber()] = _1479_v0;
      _nw269[(new BigNumber(4)).toNumber()] = _1479_v0;
      _nw269[(new BigNumber(5)).toNumber()] = _1479_v0;
      _nw269[(new BigNumber(6)).toNumber()] = _1479_v0;
      _nw269[(new BigNumber(7)).toNumber()] = new BigNumber((_1481_v2).length);
      _nw269[(new BigNumber(8)).toNumber()] = _1479_v0;
      _nw269[(new BigNumber(9)).toNumber()] = _1479_v0;
      _nw269[(new BigNumber(10)).toNumber()] = _1479_v0;
      _1482_v3 = _nw269;
      let _1483_v4;
      _1483_v4 = _dafny.Set.fromElements(_1482_v3, _1482_v3, _1480_v1, _1480_v1, _1480_v1);
      let _1484_v5;
      let _nw270 = new _module.C2();
      _nw270.__ctor((new BigNumber(-639)).plus(_1479_v0), (_1483_v4).Difference(_1483_v4), p0, _dafny.Seq.UnicodeFromString("nicg"));
      _1484_v5 = _nw270;
      let _index232 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1480_v1).length));
      (_1480_v1)[_index232] = _1479_v0;
      let _index233 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1480_v1).length));
      (_1480_v1)[_index233] = _1479_v0;
      let _1485_i0;
      _1485_i0 = _dafny.ZERO;
      L7: {
        while (_module.__default.fm33(false, new BigNumber(689), (_1480_v1)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((_1480_v1).length))], globalState)) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1485_i0)) {
              break L7;
            }
            _1485_i0 = (_1485_i0).plus(_dafny.ONE);
            let _1486_v6;
            _1486_v6 = _dafny.Set.fromElements(!(p0));
            _1486_v6 = ((((_this).f29) ? (_1486_v6) : (_1486_v6))).Difference((_1486_v6).Union(_1486_v6));
            let _1487_v7;
            _1487_v7 = new _dafny.CodePoint('u'.codePointAt(0));
            let _1488_v8;
            _1488_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1482_v3,(_this).f31);
            let _index234 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1480_v1).length));
            let _rhs268 = _1487_v7;
            let _rhs269 = _1488_v8;
            let _rhs270 = new BigNumber(-103);
            let _rhs271 = (_1484_v5).f38;
            let _rhs272 = !(!(_module.__default.fm10((_1484_v5).f38, (_1484_v5).f38, globalState)));
            let _lhs252 = _1480_v1;
            let _lhs253 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1480_v1).length));
            let _lhs254 = globalState;
            _1487_v7 = _rhs268;
            _1488_v8 = _rhs269;
            _lhs252[_lhs253] = _rhs270;
            _lhs254.f21 = _rhs271;
            r1 = _rhs272;
            let _1489_v9;
            _1489_v9 = _dafny.Seq.of(p0);
            let _1490_v10;
            _1490_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_this.f30, _this.f30)).length),_dafny.Seq.contains(_1489_v9, (_this).f31));
            let _1491_v11;
            _1491_v11 = _module.D12.create_DC39(new BigNumber(460), _module.__default.fm19(_1487_v7, globalState), (_1484_v5).f38);
            _1490_v10 = (_1490_v10).update((_1491_v11).dtor_cf65, (p0) || (p0));
            let _1492_v12;
            _1492_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,_1479_v0);
            let _1493_v13;
            let _nw271 = new _module.C4();
            _nw271.__ctor((_this).f32, (_1484_v5).f38);
            _1493_v13 = _nw271;
            let _1494_v14;
            _1494_v14 = _dafny.Set.fromElements(_1493_v13);
            _1492_v12 = (_1492_v12).update((_this).f29, new BigNumber(((_1494_v14).Difference(_1494_v14)).length));
          }
        }
      }
      let _1495_v15;
      _1495_v15 = _dafny.Seq.of(_1482_v3, _1480_v1);
      let _1496_v16;
      _1496_v16 = _dafny.Seq.of(_1482_v3, (_1495_v15)[_module.__default.safeIndex(_1479_v0, new BigNumber((_1495_v15).length))], _1482_v3);
      let _1497_v17;
      let _nw272 = Array((new BigNumber(5)).toNumber());
      _nw272[(_dafny.ZERO).toNumber()] = _1496_v16;
      _nw272[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_1482_v3);
      _nw272[(new BigNumber(2)).toNumber()] = _1495_v15;
      _nw272[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1495_v15, _1495_v15);
      _nw272[(new BigNumber(4)).toNumber()] = _1495_v15;
      _1497_v17 = _nw272;
      let _index235 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1497_v17).length));
      (_1497_v17)[_index235] = _dafny.Seq.Concat(_1496_v16, _1495_v15);
      let _index236 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1497_v17).length));
      (_1497_v17)[_index236] = _dafny.Seq.Concat(_1496_v16, _dafny.Seq.of(_1480_v1, _1482_v3));
      let _1498_v18;
      _1498_v18 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1479_v0));
      let _1499_v19;
      _1499_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_1498_v18);
      _1499_v19 = (_1499_v19).update((_this).f32, _dafny.MultiSet.fromElements(_1479_v0));
      let _1500_v20;
      _1500_v20 = _module.D3.create_DC10((_1498_v18).Union(_1498_v18));
      _1500_v20 = _1500_v20;
      let _1501_v21;
      _1501_v21 = _dafny.Seq.of(false);
      let _1502_v22;
      _1502_v22 = _dafny.Seq.of(_dafny.Seq.Concat(_1501_v21, _1501_v21), _dafny.Seq.of((_this).f32), _1501_v21);
      r0 = _1502_v22;
      r1 = p0;
      return [r0, r1];
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let _1503_v0;
      _1503_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(258));
      let _1504_v1;
      _1504_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,new BigNumber((_1503_v0).length));
      let _1505_v2;
      _1505_v2 = _dafny.Seq.of(!((_this).f31));
      let _1506_v3;
      _1506_v3 = _dafny.Seq.of(_1503_v0, _1503_v0, _module.__default.fm14(!(p0), _dafny.Seq.of(new BigNumber((_1505_v2).length)), (_this).f31, p2, globalState), _1504_v1, _1503_v0);
      let _1507_v4;
      let _nw273 = Array((_dafny.ONE).toNumber());
      _nw273[(_dafny.ZERO).toNumber()] = !((_1503_v0).equals((_1506_v3)[_module.__default.safeIndex(p2, new BigNumber((_1506_v3).length))]));
      _1507_v4 = _nw273;
      let _index237 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_1507_v4).length));
      (_1507_v4)[_index237] = (_this).f32;
      let _index238 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_1507_v4).length));
      (_1507_v4)[_index238] = (_this).f31;
      _1507_v4 = _1507_v4;
      r1 = (_this).f31;
      let _1508_v5;
      let _nw274 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _1508_v5 = _nw274;
      let _index239 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1508_v5).length));
      (_1508_v5)[_index239] = (_dafny.ZERO).minus(p2);
      let _1509_v6;
      _1509_v6 = _dafny.MultiSet.fromElements(p2);
      let _1510_v7;
      _1510_v7 = _dafny.MultiSet.fromElements((_this).f31, (_this).f31, (_1505_v2)[_module.__default.safeIndex(_module.__default.fm36(globalState), new BigNumber((_1505_v2).length))], true);
      let _1511_v8;
      _1511_v8 = _dafny.Map.Empty.slice().updateUnsafe((_1510_v7).equals(_1510_v7),!(!(p2).isEqualTo(p2)));
      let _1512_v9;
      _1512_v9 = _module.D3.create_DC10(_1509_v6);
      let _index240 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1508_v5).length));
      let _rhs273 = p0;
      let _rhs274 = new BigNumber((_1511_v8).length);
      let _rhs275 = (_1512_v9).dtor_cf18;
      let _rhs276 = (p2).multipliedBy(p2);
      let _rhs277 = _1507_v4;
      let _lhs255 = globalState;
      let _lhs256 = _1508_v5;
      let _lhs257 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1508_v5).length));
      let _lhs258 = globalState;
      _lhs255.f13 = _rhs273;
      _lhs256[_lhs257] = _rhs274;
      _1509_v6 = _rhs275;
      _lhs258.f19 = _rhs276;
      _1507_v4 = _rhs277;
      let _1513_v10;
      let _init40 = ((_1514_v2) => function (_1515_i0) {
        return (_dafny.Map.Empty.slice().updateUnsafe((_this).f31,_module.D1.create_DC3(_1514_v2))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D1.create_DC3(_1514_v2)));
      })(_1505_v2);
      let _nw275 = Array((new BigNumber(26)).toNumber());
      for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw275.length); _i0_40++) {
        _nw275[_i0_40] = _init40(new BigNumber(_i0_40));
      }
      _1513_v10 = _nw275;
      let _1516_v11;
      _1516_v11 = _module.D1.create_DC3(_1505_v2);
      let _1517_v12;
      _1517_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_1516_v11);
      let _index241 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_1513_v10).length));
      (_1513_v10)[_index241] = _1517_v12;
      let _index242 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_1513_v10).length));
      (_1513_v10)[_index242] = _1517_v12;
      let _1518_v13;
      let _nw276 = Array((new BigNumber(5)).toNumber());
      _nw276[(_dafny.ZERO).toNumber()] = p1;
      _nw276[(_dafny.ONE).toNumber()] = p1;
      _nw276[(new BigNumber(2)).toNumber()] = p1;
      _nw276[(new BigNumber(3)).toNumber()] = p1;
      _nw276[(new BigNumber(4)).toNumber()] = p1;
      _1518_v13 = _nw276;
      let _index243 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_1518_v13).length));
      (_1518_v13)[_index243] = p1;
      let _1519_v14;
      _1519_v14 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,new BigNumber(-828));
      let _1520_v16;
      _1520_v16 = new _dafny.CodePoint('d'.codePointAt(0));
      let _1521_v17;
      _1521_v17 = _dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(371)), function (_1522_i1) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(371)), function (_1523_i1) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })).length)), _1520_v16), _dafny.Seq.UnicodeFromString("mmynhama"));
      let _1524_v18;
      _1524_v18 = _module.D5.create_DC17(_1521_v17);
      let _1525_v19;
      _1525_v19 = _module.D5.create_DC20(_1524_v18);
      let _pat_let_tv29 = p2;
      let _pat_let_tv30 = _1508_v5;
      let _pat_let_tv31 = _1508_v5;
      let _pat_let_tv32 = _1508_v5;
      let _pat_let_tv33 = _1508_v5;
      let _index244 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_1518_v13).length));
      let _rhs278 = p1;
      let _rhs279 = function () {
        let _coll50 = new _dafny.Map();
        for (const _compr_50 of (_1521_v17).Elements) {
          let _1526_v15 = _compr_50;
          if ((_1521_v17).contains(_1526_v15)) {
            _coll50.push([_1526_v15,new BigNumber((_1509_v6).cardinality())]);
          }
        }
        return _coll50;
      }();
      let _rhs280 = function (_source21) {
        if (_source21.is_DC17) {
          let _1527___mcc_h0 = (_source21).cf27;
          let _1528_cf27 = _1527___mcc_h0;
          return _pat_let_tv29;
        } else if (_source21.is_DC18) {
          return new BigNumber((_this.f30).length);
        } else if (_source21.is_DC19) {
          let _1529___mcc_h1 = (_source21).cf28;
          let _1530___mcc_h2 = (_source21).cf29;
          let _1531___mcc_h3 = (_source21).cf30;
          let _1532_cf30 = _1531___mcc_h3;
          let _1533_cf29 = _1530___mcc_h2;
          let _1534_cf28 = _1529___mcc_h1;
          return new BigNumber(538);
        } else if (_source21.is_DC16) {
          let _1535___mcc_h4 = (_source21).cf26;
          let _1536_cf26 = _1535___mcc_h4;
          return (_pat_let_tv31)[_module.__default.safeIndex(new BigNumber(906), new BigNumber((_pat_let_tv30).length))];
        } else {
          let _1537___mcc_h5 = (_source21).cf31;
          let _1538_cf31 = _1537___mcc_h5;
          return (_pat_let_tv33)[_module.__default.safeIndex(new BigNumber(906), new BigNumber((_pat_let_tv32).length))];
        }
      }(_1525_v19);
      let _lhs259 = _1518_v13;
      let _lhs260 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_1518_v13).length));
      let _lhs261 = globalState;
      _lhs259[_lhs260] = _rhs278;
      _1519_v14 = _rhs279;
      _lhs261.f21 = _rhs280;
      let _1539_v20;
      _1539_v20 = _dafny.Seq.of(_this.f30, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-759)), ((_1540_v16) => function (_1541_i2) {
        return _1540_v16;
      })(_1520_v16)), ((true) ? (_this.f30) : (_dafny.Seq.update(_dafny.Seq.update(_this.f30, _module.__default.safeIndex(p2, new BigNumber((_this.f30).length)), new _dafny.CodePoint('i'.codePointAt(0))), _module.__default.safeIndex(new BigNumber((_1504_v1).length), new BigNumber((_dafny.Seq.update(_this.f30, _module.__default.safeIndex(p2, new BigNumber((_this.f30).length)), new _dafny.CodePoint('i'.codePointAt(0)))).length)), _1520_v16))), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("o"), _dafny.Seq.UnicodeFromString("mjfyn")), _this.f30);
      r0 = (_1539_v20)[_module.__default.safeIndex(p2, new BigNumber((_1539_v20).length))];
      r1 = ((_dafny.ZERO).minus((p2).multipliedBy(new BigNumber(89)))).isLessThanOrEqualTo((_dafny.ZERO).minus((_1508_v5)[_module.__default.safeIndex(new BigNumber(906), new BigNumber((_1508_v5).length))]));
      return [r0, r1];
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.ZERO;
      let r3 = _dafny.Set.Empty;
      let _1542_v0;
      let _init41 = function (_1543_i0) {
        return (_this).f29;
      };
      let _nw277 = Array((new BigNumber(5)).toNumber());
      for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw277.length); _i0_41++) {
        _nw277[_i0_41] = _init41(new BigNumber(_i0_41));
      }
      _1542_v0 = _nw277;
      let _1544_v1;
      _1544_v1 = _dafny.Set.fromElements(_1542_v0, _1542_v0);
      let _1545_v2;
      let _nw278 = new _module.C4();
      _nw278.__ctor(!((_module.D16.create_DC47(_1544_v1)).dtor_cf81).contains(_1542_v0), p0);
      _1545_v2 = _nw278;
      let _1546_v3;
      _1546_v3 = _module.D4.create_DC15(p0, p0);
      let _1547_v4;
      _1547_v4 = new _dafny.CodePoint('i'.codePointAt(0));
      let _1548_v5;
      _1548_v5 = _module.D10.create_DC32(_1547_v4, (_this).f32);
      let _1549_v6;
      let _nw279 = Array((new BigNumber(9)).toNumber());
      _nw279[(_dafny.ZERO).toNumber()] = _1548_v5;
      _nw279[(_dafny.ONE).toNumber()] = _1548_v5;
      _nw279[(new BigNumber(2)).toNumber()] = _1548_v5;
      _nw279[(new BigNumber(3)).toNumber()] = _module.D10.create_DC32(_1547_v4, (_this).f29);
      _nw279[(new BigNumber(4)).toNumber()] = _1548_v5;
      _nw279[(new BigNumber(5)).toNumber()] = _1548_v5;
      _nw279[(new BigNumber(6)).toNumber()] = _1548_v5;
      _nw279[(new BigNumber(7)).toNumber()] = _1548_v5;
      _nw279[(new BigNumber(8)).toNumber()] = _module.D10.create_DC32(_1547_v4, (_this).f32);
      _1549_v6 = _nw279;
      let _1550_v7;
      let _nw280 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _1550_v7 = _nw280;
      let _1551_v8;
      _1551_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_this.f30).length),!((_1545_v2).f33));
      let _1552_v9;
      _1552_v9 = _module.D15.create_DC44((_1546_v3).dtor_cf25, (_this).f32, _1549_v6, _1550_v7, new BigNumber((_1551_v8).length));
      let _1553_v10;
      _1553_v10 = _module.D3.create_DC11((_1545_v2).f33, (_this).f32);
      let _1554_v11;
      _1554_v11 = _dafny.MultiSet.fromElements((_1545_v2).f34);
      let _1555_v12;
      _1555_v12 = _module.D11.create_DC37((_1552_v9).dtor_cf75, _module.__default.fm33(!((_module.D3.create_DC11((_1553_v10).dtor_cf19, (_this).f32)).dtor_cf19), p0, new BigNumber((_1554_v11).cardinality()), globalState));
      let _source22 = _1555_v12;
      if (_source22.is_DC35) {
        let _1556___mcc_h0 = (_source22).cf57;
        let _1557___mcc_h1 = (_source22).cf58;
        let _1558_cf58 = _1557___mcc_h1;
        let _1559_cf57 = _1556___mcc_h0;
        (globalState).f15 = _1558_cf58;
        if ((p1).isLessThanOrEqualTo(p1)) {
          (globalState).f2 = (_module.__default.fm36(globalState)).minus(_module.__default.safeDivisionInt((_1545_v2).f34, (_1545_v2).f34));
          let _1560_v13;
          _1560_v13 = _dafny.Set.fromElements((_1545_v2).f34, (_1545_v2).f34);
          let _1561_v14;
          _1561_v14 = _module.D17.create_DC49(_1560_v13);
          let _1562_v15;
          _1562_v15 = _dafny.Seq.of((_1545_v2).f33);
          let _1563_v16;
          _1563_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1545_v2).f34,(_1545_v2).f34);
          let _1564_v17;
          _1564_v17 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,_1560_v13);
          let _1565_v19;
          let _nw281 = Array((new BigNumber(20)).toNumber());
          _nw281[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements((_1545_v2).f34, (_1545_v2).f34, p1, p0, (_1545_v2).f34);
          _nw281[(_dafny.ONE).toNumber()] = (_1560_v13).Intersect(_1560_v13);
          _nw281[(new BigNumber(2)).toNumber()] = _1560_v13;
          _nw281[(new BigNumber(3)).toNumber()] = (_1561_v14).dtor_cf82;
          _nw281[(new BigNumber(4)).toNumber()] = _1560_v13;
          _nw281[(new BigNumber(5)).toNumber()] = _1560_v13;
          _nw281[(new BigNumber(6)).toNumber()] = (_1560_v13).Intersect(_1560_v13);
          _nw281[(new BigNumber(7)).toNumber()] = (_1560_v13).Difference(_1560_v13);
          _nw281[(new BigNumber(8)).toNumber()] = _1560_v13;
          _nw281[(new BigNumber(9)).toNumber()] = _1560_v13;
          _nw281[(new BigNumber(10)).toNumber()] = _1560_v13;
          _nw281[(new BigNumber(11)).toNumber()] = _1560_v13;
          _nw281[(new BigNumber(12)).toNumber()] = _1560_v13;
          _nw281[(new BigNumber(13)).toNumber()] = _1560_v13;
          _nw281[(new BigNumber(14)).toNumber()] = (((_1545_v2).f33) ? (_module.__default.fm17((_1545_v2).f34, globalState)) : (_1560_v13));
          _nw281[(new BigNumber(15)).toNumber()] = (_1560_v13).Difference(_1560_v13);
          _nw281[(new BigNumber(16)).toNumber()] = _1560_v13;
          _nw281[(new BigNumber(17)).toNumber()] = (_dafny.Set.fromElements((_1545_v2).f34, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(730)), ((_1566_v4) => function (_1567_i1) {
            return _1566_v4;
          })(_1547_v4)), _module.__default.safeIndex((_1545_v2).f34, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(730)), ((_1568_v4) => function (_1569_i1) {
            return _1568_v4;
          })(_1547_v4))).length)), _1547_v4)).length), new BigNumber((_1562_v15).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1545_v2).f33,_1563_v16)).length), p1)).Difference((((_1564_v17).contains(_this.f30)) ? ((_1564_v17).get(_this.f30)) : (_dafny.Set.fromElements(new BigNumber((_this.f30).length)))));
          _nw281[(new BigNumber(18)).toNumber()] = (function () {
            let _coll51 = new _dafny.Set();
            for (const _compr_51 of _dafny.IntegerRange(new BigNumber(735), new BigNumber(598))) {
              let _1570_v18 = _compr_51;
              if (((new BigNumber(735)).isLessThanOrEqualTo(_1570_v18)) && ((_1570_v18).isLessThan(new BigNumber(598)))) {
                _coll51.add(_module.__default.safeDivisionInt(_1570_v18, p0));
              }
            }
            return _coll51;
          }()).Difference(_1560_v13);
          _nw281[(new BigNumber(19)).toNumber()] = _1560_v13;
          _1565_v19 = _nw281;
          let _index245 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_1565_v19).length));
          (_1565_v19)[_index245] = (_1560_v13).Intersect(_1560_v13);
          let _1571_v20;
          _1571_v20 = _dafny.Map.Empty.slice().updateUnsafe((_1548_v5).dtor_cf52,(_1560_v13).Difference(_1560_v13));
          let _index246 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_1565_v19).length));
          (_1565_v19)[_index246] = (((_1571_v20).contains(new _dafny.CodePoint('u'.codePointAt(0)))) ? ((_1571_v20).get(new _dafny.CodePoint('u'.codePointAt(0)))) : (_1560_v13));
          let _1572_v21;
          _1572_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1542_v0,_1547_v4);
          let _1573_v22;
          _1573_v22 = _module.D0.create_DC1(_1542_v0, p0, new BigNumber((_1559_cf57).length));
          _1572_v21 = (_1572_v21).update((_1573_v22).dtor_cf1, _1547_v4);
          let _1574_v23;
          let _nw282 = new _module.C5();
          _nw282.__ctor((((_this).f32) ? ((_this).f32) : ((_this).f31)), _1559_cf57);
          _1574_v23 = _nw282;
          let _1575_v24;
          _1575_v24 = _dafny.MultiSet.fromElements((_1545_v2).f33, (_1545_v2).f33);
          let _1576_v25;
          _1576_v25 = _dafny.Seq.of(new BigNumber((_1575_v24).cardinality()));
          let _1577_v26;
          let _nw283 = new _module.C1();
          _nw283.__ctor(!_dafny.areEqual(_module.__default.fm44((_1545_v2).f33, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(629)), ((_1578_v2) => function (_1579_i2) {
            return (_1578_v2).f34;
          })(_1545_v2))).length), (_1545_v2).f33, (_1545_v2).f33, globalState), _1576_v25));
          _1577_v26 = _nw283;
        } else {
          (globalState).f13 = _dafny.areEqual(_dafny.Seq.Concat(_this.f30, _dafny.Seq.UnicodeFromString("dnewr")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wf"), _dafny.Seq.UnicodeFromString("efgrsogq")));
          (globalState).f15 = (_1545_v2).f33;
          let _1580_v27;
          let _nw284 = Array((new BigNumber(23)).toNumber()).fill(_module.D2.Default());
          _1580_v27 = _nw284;
          let _1581_v28;
          _1581_v28 = _module.D12.create_DC39(p1, (_1545_v2).f33, new BigNumber(146));
          let _1582_v29;
          _1582_v29 = _module.D2.create_DC9(_this.f30, p1, _1558_cf58, (_1581_v28).dtor_cf67);
          let _index247 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1580_v27).length));
          (_1580_v27)[_index247] = _1582_v29;
          let _index248 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1580_v27).length));
          (_1580_v27)[_index248] = _1582_v29;
          let _1583_v30;
          _1583_v30 = _module.D17.create_DC52(_module.__default.fm6(new BigNumber((_1559_cf57).length), globalState), (_1545_v2).f33, (_1545_v2).f34, _module.__default.fm27((_1545_v2).f34, true, globalState));
          let _index249 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1550_v7).length));
          (_1550_v7)[_index249] = _module.__default.safeModuloInt((_1583_v30).dtor_cf90, (_1583_v30).dtor_cf90);
          let _index250 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1550_v7).length));
          (_1550_v7)[_index250] = new BigNumber(-602);
          let _1584_v32;
          _1584_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(622),new BigNumber((function () {
            let _coll52 = new _dafny.Map();
            for (const _compr_52 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(190)), ((_1585_v2) => function (_1586_i3) {
              return (_1585_v2).f34;
            })(_1545_v2))).Elements) {
              let _1587_v31 = _compr_52;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(190)), ((_1588_v2) => function (_1586_i3) {
                return (_1588_v2).f34;
              })(_1545_v2)), _1587_v31)) {
                _coll52.push([_module.__default.safeModuloInt(_1587_v31, (_1545_v2).f34),(_1545_v2).f34]);
              }
            }
            return _coll52;
          }()).length));
          (globalState).f19 = (new BigNumber((_1584_v32).length)).multipliedBy((_1545_v2).f34);
        }
        let _1589_v33;
        _1589_v33 = _dafny.Seq.of(p0);
        if (!(_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(301)), function (_1590_i4) {
          return new BigNumber(899);
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-501)), ((_1591_v2) => function (_1592_i5) {
          return (_1591_v2).f34;
        })(_1545_v2))), _1589_v33))) {
          _1547_v4 = _1547_v4;
          (globalState).f13 = !((_1558_cf58) || (_dafny.Seq.IsProperPrefixOf(_1559_cf57, _this.f30)));
          _1554_v11 = (_1554_v11).Intersect(_1554_v11);
          let _index251 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1550_v7).length));
          (_1550_v7)[_index251] = p1;
          let _index252 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1550_v7).length));
          let _rhs281 = (_1589_v33)[_module.__default.safeIndex((new BigNumber(847)).minus(new BigNumber((_1589_v33).length)), new BigNumber((_1589_v33).length))];
          let _rhs282 = (((_1545_v2).f33) ? ((_1545_v2).f33) : (_1558_cf58));
          let _rhs283 = (((_1545_v2).f33) ? (_dafny.Seq.Concat(_module.__default.fm34(p1, (_1545_v2).f34, (_this).f31, !(_module.__default.fm33((_this).f29, p1, new BigNumber(848), globalState)), globalState), _1559_cf57)) : (_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("c"), _this.f30)));
          let _rhs284 = (_1545_v2).f34;
          let _rhs285 = _1550_v7;
          let _lhs262 = _1550_v7;
          let _lhs263 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1550_v7).length));
          let _lhs264 = globalState;
          let _lhs265 = _this;
          let _lhs266 = globalState;
          let _lhs267 = globalState;
          _lhs262[_lhs263] = _rhs281;
          _lhs264.f18 = _rhs282;
          _lhs265.f30 = _rhs283;
          _lhs266.f19 = _rhs284;
          _lhs267.f17 = _rhs285;
          (globalState).f1 = (_1545_v2).f34;
        } else {
          let _index253 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1542_v0).length));
          (_1542_v0)[_index253] = false;
          let _index254 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1542_v0).length));
          (_1542_v0)[_index254] = (p1).isEqualTo((_1545_v2).f34);
          let _1593_v34;
          _1593_v34 = _dafny.Set.fromElements((_1545_v2).f33, (_this).f31);
          let _1594_v35;
          _1594_v35 = _dafny.Seq.of((_this).f31, _1558_cf58);
          let _1595_v36;
          _1595_v36 = _dafny.Map.Empty.slice().updateUnsafe((_1545_v2).f33,_module.__default.fm20((_1542_v0)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1542_v0).length))], (_1542_v0)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1542_v0).length))], _1594_v35, (_1545_v2).f34, globalState));
          let _rhs286 = (_1554_v11).IsProperSubsetOf(_1554_v11);
          let _rhs287 = _dafny.Seq.Concat(_this.f30, _module.__default.fm34(new BigNumber((_1593_v34).length), new BigNumber(((((_1595_v36).contains(false)) ? ((_1595_v36).get(false)) : (_1559_cf57))).length), (_this).f31, (_1542_v0)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1542_v0).length))], globalState));
          let _rhs288 = (_1589_v33)[_module.__default.safeIndex(((_1545_v2).f34).plus((_1545_v2).f34), new BigNumber((_1589_v33).length))];
          let _rhs289 = p1;
          let _lhs268 = globalState;
          let _lhs269 = globalState;
          let _lhs270 = globalState;
          let _lhs271 = globalState;
          _lhs268.f18 = _rhs286;
          _lhs269.f11 = _rhs287;
          _lhs270.f1 = _rhs288;
          _lhs271.f14 = _rhs289;
          let _1596_v37;
          _1596_v37 = _dafny.Set.fromElements(_1550_v7);
          let _1597_v38;
          _1597_v38 = _dafny.Seq.of(_1596_v37);
          let _1598_v39;
          let _nw285 = new _module.C2();
          _nw285.__ctor((_1545_v2).f34, (_1597_v38)[_module.__default.safeIndex((_1545_v2).f34, new BigNumber((_1597_v38).length))], (_this).f29, _this.f30);
          _1598_v39 = _nw285;
          let _1599_v40;
          _1599_v40 = _dafny.MultiSet.fromElements(_1550_v7, _1550_v7);
          let _1600_v41;
          let _nw286 = Array((new BigNumber(20)).toNumber()).fill(false);
          _1600_v41 = _nw286;
          let _1601_v42;
          _1601_v42 = _dafny.MultiSet.fromElements(_1600_v41, _1600_v41);
          let _index255 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1542_v0).length));
          (_1542_v0)[_index255] = ((_1599_v40).Union((_1599_v40).update(_1550_v7, _module.__default.abs(new BigNumber((_1601_v42).cardinality()))))).IsSubsetOf((_dafny.MultiSet.fromElements(_1550_v7)).Union(_1599_v40));
          (globalState).f2 = _module.__default.fm9(globalState);
        }
        let _1602_v43;
        let _nw287 = new _module.C3();
        _nw287.__ctor(true, ((_1545_v2).f34).multipliedBy(p1), (_this).f32, _1559_cf57);
        _1602_v43 = _nw287;
      } else if (_source22.is_DC36) {
        let _1603___mcc_h2 = (_source22).cf59;
        let _1604___mcc_h3 = (_source22).cf60;
        let _1605___mcc_h4 = (_source22).cf61;
        let _1606_cf61 = _1605___mcc_h4;
        let _1607_cf60 = _1604___mcc_h3;
        let _1608_cf59 = _1603___mcc_h2;
        let _1609_v44;
        _1609_v44 = _dafny.Seq.of((_1545_v2).f33, true, (((_this).f32) ? ((_1545_v2).f33) : ((_1545_v2).f33)), (_this).f29);
        (globalState).f13 = (_1609_v44)[_module.__default.safeIndex(_module.__default.fm9(globalState), new BigNumber((_1609_v44).length))];
        let _index256 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_1550_v7).length));
        (_1550_v7)[_index256] = (_1545_v2).f34;
        let _1610_v45;
        _1610_v45 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(250)), ((_1611_v4) => function (_1612_i6) {
          return _1611_v4;
        })(_1547_v4))).length));
        let _index257 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_1550_v7).length));
        (_1550_v7)[_index257] = (p0).plus((new BigNumber((_1610_v45).length)).minus(p1));
        let _1613_v46;
        _1613_v46 = _module.D3.create_DC10(_1554_v11);
        let _1614_v47;
        _1614_v47 = _module.D3.create_DC12(_1613_v46);
        let _source23 = (((_this).f29) ? (_1614_v47) : (_1614_v47));
        if (_source23.is_DC11) {
          let _1615___mcc_h8 = (_source23).cf19;
          let _1616___mcc_h9 = (_source23).cf20;
          let _1617_cf20 = _1616___mcc_h9;
          let _1618_cf19 = _1615___mcc_h8;
          let _1619_v48;
          let _nw288 = new _module.C1();
          _nw288.__ctor((_this).f29);
          _1619_v48 = _nw288;
          let _1620_v49;
          _1620_v49 = _dafny.Seq.of(_1619_v48, _1619_v48);
          let _1621_v50;
          _1621_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1617_cf20,_dafny.MultiSet.fromElements(_1619_v48, _1619_v48, (_1620_v49)[_module.__default.safeIndex((_1545_v2).f34, new BigNumber((_1620_v49).length))], _1619_v48));
          let _1622_v51;
          _1622_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1621_v50,(((_1551_v8).contains((_1550_v7)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_1550_v7).length))])) ? ((_1551_v8).get((_1550_v7)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_1550_v7).length))])) : (!(false))));
          let _1623_v52;
          _1623_v52 = _dafny.MultiSet.fromElements(_1619_v48);
          let _1624_v53;
          _1624_v53 = _dafny.Seq.of(p1, (_dafny.ZERO).minus((_1550_v7)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_1550_v7).length))]), (new BigNumber((_1609_v44).length)).plus(new BigNumber((_dafny.Set.fromElements((_1550_v7)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_1550_v7).length))], (_1545_v2).f34, p0)).length)), p1);
          let _index258 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_1550_v7).length));
          let _rhs290 = (((_1622_v51).contains((_1621_v50).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1606_cf61,_1623_v52)))) ? ((_1622_v51).get((_1621_v50).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1606_cf61,_1623_v52)))) : (((_this).f31) === (_1618_cf19)));
          let _rhs291 = (_module.__default.safeModuloInt((_1545_v2).f34, (_1545_v2).f34)).minus((_1545_v2).f34);
          let _rhs292 = _1624_v53;
          let _lhs272 = globalState;
          let _lhs273 = _1550_v7;
          let _lhs274 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_1550_v7).length));
          let _lhs275 = globalState;
          _lhs272.f28 = _rhs290;
          _lhs273[_lhs274] = _rhs291;
          _lhs275.f7 = _rhs292;
          let _1625_v54;
          _1625_v54 = _dafny.Seq.of(_1614_v47, _module.__default.fm29((_1545_v2).f33, (_1550_v7)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_1550_v7).length))], _1606_cf61, globalState));
          let _rhs293 = !(_1617_cf20) || (_1617_cf20);
          let _rhs294 = !(new BigNumber(847)).isEqualTo(new BigNumber((_dafny.Seq.Concat(_1625_v54, _1625_v54)).length));
          let _rhs295 = (_1607_cf60).Union(((_1617_cf20) ? (_1607_cf60) : (_1607_cf60)));
          let _rhs296 = (p0).multipliedBy((_1545_v2).f34);
          let _lhs276 = globalState;
          let _lhs277 = globalState;
          let _lhs278 = globalState;
          _lhs276.f15 = _rhs293;
          _lhs277.f18 = _rhs294;
          _1607_cf60 = _rhs295;
          _lhs278.f14 = _rhs296;
          let _1626_v55;
          let _init42 = ((_1627_v4) => function (_1628_i7) {
            return _1627_v4;
          })(_1547_v4);
          let _nw289 = Array((new BigNumber(21)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw289.length); _i0_42++) {
            _nw289[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1626_v55 = _nw289;
          let _rhs297 = _1626_v55;
          let _rhs298 = !(_1606_cf61);
          let _rhs299 = new BigNumber(((function () {
            let _coll53 = new _dafny.Map();
            for (const _compr_53 of _dafny.IntegerRange(new BigNumber(228), new BigNumber(-272))) {
              let _1629_v56 = _compr_53;
              if (((new BigNumber(228)).isLessThanOrEqualTo(_1629_v56)) && ((_1629_v56).isLessThan(new BigNumber(-272)))) {
                _coll53.push([(_1629_v56).plus(p1),(_1619_v48).f40]);
              }
            }
            return _coll53;
          }()).Merge((_module.__default.fm41(p0, (_this).f29, _1606_cf61, (_1545_v2).f33, globalState)).Merge(_1551_v8))).length);
          let _lhs279 = globalState;
          let _lhs280 = globalState;
          r1 = _rhs297;
          _lhs279.f13 = _rhs298;
          _lhs280.f2 = _rhs299;
          (globalState).f15 = _1618_cf19;
        } else if (_source23.is_DC10) {
          let _1630___mcc_h10 = (_source23).cf18;
          let _1631_cf18 = _1630___mcc_h10;
          let _1632_v57;
          _1632_v57 = _dafny.MultiSet.fromElements(_1547_v4);
          (globalState).f21 = new BigNumber((((_1632_v57).Intersect(_1632_v57)).Difference(_dafny.MultiSet.fromElements(_1547_v4))).cardinality());
          let _rhs300 = ((_dafny.ZERO).minus((_1545_v2).f34)).multipliedBy((_1545_v2).f34);
          let _rhs301 = _dafny.Seq.Concat(_this.f30, _this.f30);
          let _rhs302 = (_1545_v2).fm5(_dafny.Set.fromElements((_this).f31, _1606_cf61), globalState);
          let _rhs303 = _module.__default.fm33(((_1545_v2).f34).isLessThanOrEqualTo((_1545_v2).f34), p1, (_1545_v2).f34, globalState);
          let _rhs304 = new BigNumber(-554);
          let _lhs281 = globalState;
          let _lhs282 = globalState;
          let _lhs283 = globalState;
          let _lhs284 = globalState;
          let _lhs285 = globalState;
          _lhs281.f2 = _rhs300;
          _lhs282.f11 = _rhs301;
          _lhs283.f21 = _rhs302;
          _lhs284.f28 = _rhs303;
          _lhs285.f21 = _rhs304;
          let _rhs305 = _1542_v0;
          let _rhs306 = _dafny.Seq.Concat(_this.f30, _dafny.Seq.UnicodeFromString("uef"));
          let _lhs286 = globalState;
          _1542_v0 = _rhs305;
          _lhs286.f11 = _rhs306;
          let _1633_v58;
          let _nw290 = Array((new BigNumber(28)).toNumber());
          _1633_v58 = _nw290;
          let _1634_v59;
          let _nw291 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _1634_v59 = _nw291;
          let _1635_v60;
          _1635_v60 = _dafny.Set.fromElements(_1634_v59);
          let _1636_v61;
          let _nw292 = new _module.C2();
          _nw292.__ctor(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1609_v44,(_this).f31)).length), _1635_v60, (_this).f31, _this.f30);
          _1636_v61 = _nw292;
          let _index259 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_1633_v58).length));
          (_1633_v58)[_index259] = _1636_v61;
          let _index260 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_1633_v58).length));
          (_1633_v58)[_index260] = _1636_v61;
        } else {
          let _1637___mcc_h11 = (_source23).cf21;
          let _1638_cf21 = _1637___mcc_h11;
          let _1639_v62;
          let _init43 = ((_1640_v2) => function (_1641_i8) {
            return _dafny.Set.fromElements(new BigNumber(146), (_1640_v2).f34);
          })(_1545_v2);
          let _nw293 = Array((new BigNumber(28)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw293.length); _i0_43++) {
            _nw293[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1639_v62 = _nw293;
          let _index261 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_1639_v62).length));
          (_1639_v62)[_index261] = _1610_v45;
          let _index262 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_1639_v62).length));
          (_1639_v62)[_index262] = (((_1545_v2).f33) ? (_1610_v45) : ((_dafny.Set.fromElements(new BigNumber((_this.f30).length), (_1545_v2).f34)).Intersect(_1610_v45)));
          let _1642_v63;
          let _nw294 = new _module.C1();
          _nw294.__ctor(_1606_cf61);
          _1642_v63 = _nw294;
          let _1643_v64;
          _1643_v64 = _dafny.Seq.of(_1642_v63);
          let _1644_v65;
          let _nw295 = Array((new BigNumber(11)).toNumber());
          _nw295[(_dafny.ZERO).toNumber()] = _1642_v63;
          _nw295[(_dafny.ONE).toNumber()] = _1642_v63;
          _nw295[(new BigNumber(2)).toNumber()] = (_1643_v64)[_module.__default.safeIndex(p1, new BigNumber((_1643_v64).length))];
          _nw295[(new BigNumber(3)).toNumber()] = _1642_v63;
          _nw295[(new BigNumber(4)).toNumber()] = _1642_v63;
          _nw295[(new BigNumber(5)).toNumber()] = _1642_v63;
          _nw295[(new BigNumber(6)).toNumber()] = _1642_v63;
          _nw295[(new BigNumber(7)).toNumber()] = (((_this).f29) ? (_1642_v63) : (_1642_v63));
          _nw295[(new BigNumber(8)).toNumber()] = _1642_v63;
          _nw295[(new BigNumber(9)).toNumber()] = _1642_v63;
          _nw295[(new BigNumber(10)).toNumber()] = _1642_v63;
          _1644_v65 = _nw295;
          let _index263 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_1644_v65).length));
          (_1644_v65)[_index263] = _1642_v63;
          let _index264 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1542_v0).length));
          (_1542_v0)[_index264] = _module.__default.fm19(_1547_v4, globalState);
          let _1645_v66;
          let _init44 = function (_1646_i9) {
            return _dafny.Seq.UnicodeFromString("agpxc");
          };
          let _nw296 = Array((new BigNumber(29)).toNumber());
          for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw296.length); _i0_44++) {
            _nw296[_i0_44] = _init44(new BigNumber(_i0_44));
          }
          _1645_v66 = _nw296;
          let _index265 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1645_v66).length));
          (_1645_v66)[_index265] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(870)), ((_1647_v4) => function (_1648_i10) {
            return _1647_v4;
          })(_1547_v4));
          let _index266 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_1644_v65).length));
          let _index267 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1542_v0).length));
          let _index268 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1645_v66).length));
          let _rhs307 = (_1545_v2).f33;
          let _rhs308 = _1642_v63;
          let _rhs309 = (_1545_v2).f33;
          let _rhs310 = new BigNumber(-943);
          let _rhs311 = _dafny.Seq.UnicodeFromString("wrfecmi");
          let _lhs287 = globalState;
          let _lhs288 = _1644_v65;
          let _lhs289 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_1644_v65).length));
          let _lhs290 = _1542_v0;
          let _lhs291 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1542_v0).length));
          let _lhs292 = globalState;
          let _lhs293 = _1645_v66;
          let _lhs294 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1645_v66).length));
          _lhs287.f18 = _rhs307;
          _lhs288[_lhs289] = _rhs308;
          _lhs290[_lhs291] = _rhs309;
          _lhs292.f21 = _rhs310;
          _lhs293[_lhs294] = _rhs311;
          let _1649_v67;
          _1649_v67 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,(_1542_v0)[_module.__default.safeIndex(new BigNumber(667), new BigNumber((_1542_v0).length))]);
          let _1650_v68;
          _1650_v68 = _dafny.Seq.of(p0, new BigNumber((_module.__default.fm7((_dafny.ZERO).minus((_1545_v2).f34), globalState)).length), (_1545_v2).f34, (_1545_v2).f34, (_1550_v7)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_1550_v7).length))]);
          let _rhs312 = (((_this).f32) ? (_1650_v68) : (_dafny.Seq.of((_1545_v2).f34)));
          let _rhs313 = _dafny.Seq.Concat(_1650_v68, _1650_v68);
          let _rhs314 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,(_1542_v0)[_module.__default.safeIndex(new BigNumber(667), new BigNumber((_1542_v0).length))]);
          let _lhs295 = globalState;
          let _lhs296 = globalState;
          _lhs295.f7 = _rhs312;
          _lhs296.f7 = _rhs313;
          _1649_v67 = _rhs314;
          let _1651_v69;
          let _nw297 = new _module.C0();
          _nw297.__ctor(((_1639_v62)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_1639_v62).length))]).IsDisjointFrom((_1639_v62)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_1639_v62).length))]));
          _1651_v69 = _nw297;
          let _nw298 = new _module.C0();
          _nw298.__ctor((_1554_v11).equals(_1554_v11));
          _1651_v69 = _nw298;
        }
        (globalState).f28 = (_1551_v8).contains(new BigNumber(-874));
      } else if (_source22.is_DC37) {
        let _1652___mcc_h5 = (_source22).cf62;
        let _1653___mcc_h6 = (_source22).cf63;
        let _1654_cf63 = _1653___mcc_h6;
        let _1655_cf62 = _1652___mcc_h5;
        if ((((new BigNumber(-359)).isLessThanOrEqualTo(p0)) ? ((((_1545_v2).f33) ? (_1654_cf63) : (false))) : ((p1).isLessThanOrEqualTo((_1545_v2).f34)))) {
          let _1656_v70;
          let _init45 = ((_1657_v4) => function (_1658_i11) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(550)), ((_1659_v4) => function (_1660_i12) {
              return _1659_v4;
            })(_1657_v4));
          })(_1547_v4);
          let _nw299 = Array((new BigNumber(6)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw299.length); _i0_45++) {
            _nw299[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1656_v70 = _nw299;
          let _index269 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_1656_v70).length));
          (_1656_v70)[_index269] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-415)), ((_1661_v4) => function (_1662_i13) {
            return _1661_v4;
          })(_1547_v4));
          let _index270 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_1656_v70).length));
          (_1656_v70)[_index270] = _dafny.Seq.Concat(_this.f30, _dafny.Seq.Create(_module.__default.abs(new BigNumber(205)), function (_1663_i14) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          }));
          let _1664_v71;
          _1664_v71 = _dafny.Map.Empty.slice().updateUnsafe((_1545_v2).f33,_1547_v4);
          let _1665_v72;
          _1665_v72 = _dafny.Seq.of(_1664_v71);
          let _1666_v73;
          _1666_v73 = _module.D2.create_DC8(_1547_v4, (_this).f31);
          (globalState).f18 = (((_1551_v8).contains(_1655_cf62)) ? ((_1551_v8).get(_1655_cf62)) : ((new BigNumber(-222)).isLessThan(new BigNumber((((_1665_v72)[_module.__default.safeIndex(p0, new BigNumber((_1665_v72).length))]).update(_1654_cf63, (_1666_v73).dtor_cf12)).length))));
          let _rhs315 = _dafny.Seq.Concat(_dafny.Seq.Concat(_this.f30, (_1656_v70)[_module.__default.safeIndex(new BigNumber(274), new BigNumber((_1656_v70).length))]), _dafny.Seq.Create(_module.__default.abs(new BigNumber(795)), ((_1667_v4) => function (_1668_i15) {
            return _1667_v4;
          })(_1547_v4)));
          let _rhs316 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1545_v2).f33,(_dafny.ZERO).minus((_1545_v2).f34))).length), (_1545_v2).f34)).length), (_1545_v2).f34);
          let _lhs297 = globalState;
          let _lhs298 = globalState;
          _lhs297.f20 = _rhs315;
          _lhs298.f1 = _rhs316;
          let _1669_v74;
          _1669_v74 = _dafny.Seq.of((_1656_v70)[_module.__default.safeIndex(new BigNumber(274), new BigNumber((_1656_v70).length))], _this.f30, _dafny.Seq.UnicodeFromString("pbew"), (_1656_v70)[_module.__default.safeIndex(new BigNumber(274), new BigNumber((_1656_v70).length))], _this.f30);
          _1551_v8 = (_1551_v8).update(new BigNumber((_dafny.Seq.Concat(_1669_v74, _1669_v74)).length), (_this).f32);
          let _1670_v75;
          _1670_v75 = _dafny.Seq.of(_1654_cf63, (_this).f31);
          (globalState).f28 = _dafny.areEqual(_1670_v75, _dafny.Seq.Concat(_1670_v75, _dafny.Seq.of((_this).f32)));
        } else {
          _1655_cf62 = (((_1554_v11).contains(_1655_cf62)) ? ((_1554_v11).get(_1655_cf62)) : (p0));
          let _1671_v76;
          let _nw300 = new _module.C5();
          _nw300.__ctor(_1654_cf63, _this.f30);
          _1671_v76 = _nw300;
          (globalState).f2 = p1;
          r2 = p1;
          let _1672_v77;
          _1672_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1654_cf63,_this.f30);
          _1655_cf62 = new BigNumber((((_module.__default.fm52(globalState)).update((_this).f29, _this.f30)).Merge(_1672_v77)).length);
        }
        _1547_v4 = new _dafny.CodePoint('c'.codePointAt(0));
        let _index271 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1542_v0).length));
        (_1542_v0)[_index271] = _dafny.Seq.IsPrefixOf(_this.f30, _dafny.Seq.UnicodeFromString("ejv"));
        let _index272 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1542_v0).length));
        (_1542_v0)[_index272] = _module.__default.fm19(_1547_v4, globalState);
        let _1673_v78;
        _1673_v78 = _module.D11.create_DC35(_dafny.Seq.UnicodeFromString("qqnwtssk"), (_1545_v2).f33);
        let _1674_v79;
        _1674_v79 = _dafny.MultiSet.fromElements(!((_1673_v78).dtor_cf58));
        let _1675_v80;
        _1675_v80 = _dafny.MultiSet.fromElements(_1547_v4, _1547_v4);
        let _1676_v81;
        _1676_v81 = _dafny.Seq.of(_1654_cf63);
        let _1677_v82;
        _1677_v82 = _dafny.Seq.of(_1676_v81, _1676_v81, _1676_v81, _1676_v81);
        _1674_v79 = (_dafny.MultiSet.fromElements((_this).f31, (_1545_v2).f33)).Union(((_1674_v79).update(_1654_cf63, _module.__default.abs(new BigNumber((_1675_v80).cardinality())))).Union(_dafny.MultiSet.FromArray((_1677_v82)[_module.__default.safeIndex((_1545_v2).f34, new BigNumber((_1677_v82).length))])));
      } else {
        let _1678___mcc_h7 = (_source22).cf56;
        let _1679_cf56 = _1678___mcc_h7;
        let _1680_v83;
        _1680_v83 = _dafny.Seq.of((_1545_v2).f34);
        let _1681_v84;
        _1681_v84 = _module.D9.create_DC27(_1680_v83);
        let _1682_v85;
        _1682_v85 = _module.D2.create_DC9(_this.f30, new BigNumber(((_1681_v84).dtor_cf40).length), (_1545_v2).f33, (_1545_v2).f34);
        let _source24 = _1682_v85;
        if (_source24.is_DC7) {
          let _1683___mcc_h12 = (_source24).cf11;
          let _1684_cf11 = _1683___mcc_h12;
          let _1685_v86;
          _1685_v86 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_1545_v2).f34, (_1545_v2).f34),_this.f30);
          _1685_v86 = (_1685_v86).update(p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(663)), function (_1686_i16) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          }));
          (globalState).f13 = _1684_cf11;
          (globalState).f1 = (_dafny.ZERO).minus(p1);
          let _1687_v87;
          _1687_v87 = _dafny.Seq.of((_this).f32, (_this).f32, ((_1545_v2).f34).isLessThanOrEqualTo(p1));
          (globalState).f13 = (_1687_v87)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_this.f30, _this.f30)).length), new BigNumber((_1687_v87).length))];
        } else if (_source24.is_DC8) {
          let _1688___mcc_h13 = (_source24).cf12;
          let _1689___mcc_h14 = (_source24).cf13;
          let _1690_cf13 = _1689___mcc_h14;
          let _1691_cf12 = _1688___mcc_h13;
          let _1692_v88;
          _1692_v88 = _dafny.MultiSet.fromElements(_1690_cf13, (_this).f31);
          let _1693_v89;
          _1693_v89 = _dafny.Seq.of((_this).f32, (_1545_v2).f33, false);
          _1692_v88 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1693_v89, _1693_v89));
          (globalState).f2 = (_1545_v2).f34;
          let _rhs317 = ((_1690_cf13) ? ((p1).isLessThanOrEqualTo(p0)) : ((_this).f32));
          let _rhs318 = _1690_cf13;
          let _rhs319 = (_1545_v2).f33;
          let _rhs320 = (_1545_v2).f34;
          let _lhs299 = globalState;
          let _lhs300 = globalState;
          let _lhs301 = globalState;
          let _lhs302 = globalState;
          _lhs299.f18 = _rhs317;
          _lhs300.f18 = _rhs318;
          _lhs301.f13 = _rhs319;
          _lhs302.f2 = _rhs320;
          let _1694_v90;
          _1694_v90 = _dafny.Set.fromElements((_1545_v2).f33);
          let _1695_v91;
          _1695_v91 = _module.D6.create_DC21(_1694_v90);
          let _1696_v92;
          _1696_v92 = _dafny.Seq.of(_1695_v91);
          let _1697_v93;
          _1697_v93 = _module.D8.create_DC25(_1696_v92);
          let _1698_v94;
          _1698_v94 = _module.D8.create_DC25((_1697_v93).dtor_cf39);
          let _1699_v95;
          let _nw301 = Array((new BigNumber(20)).toNumber());
          _nw301[(_dafny.ZERO).toNumber()] = _1697_v93;
          _nw301[(_dafny.ONE).toNumber()] = _1698_v94;
          _nw301[(new BigNumber(2)).toNumber()] = _1698_v94;
          _nw301[(new BigNumber(3)).toNumber()] = _1697_v93;
          _nw301[(new BigNumber(4)).toNumber()] = _1698_v94;
          _nw301[(new BigNumber(5)).toNumber()] = _1698_v94;
          _nw301[(new BigNumber(6)).toNumber()] = _module.D8.create_DC25(_1696_v92);
          _nw301[(new BigNumber(7)).toNumber()] = _1698_v94;
          _nw301[(new BigNumber(8)).toNumber()] = _module.D8.create_DC25(_dafny.Seq.of(_1695_v91, _module.D6.create_DC21(_1694_v90)));
          _nw301[(new BigNumber(9)).toNumber()] = _1697_v93;
          _nw301[(new BigNumber(10)).toNumber()] = _1698_v94;
          _nw301[(new BigNumber(11)).toNumber()] = _1697_v93;
          _nw301[(new BigNumber(12)).toNumber()] = _1697_v93;
          _nw301[(new BigNumber(13)).toNumber()] = _1697_v93;
          _nw301[(new BigNumber(14)).toNumber()] = _1697_v93;
          _nw301[(new BigNumber(15)).toNumber()] = _1698_v94;
          _nw301[(new BigNumber(16)).toNumber()] = _1698_v94;
          _nw301[(new BigNumber(17)).toNumber()] = _1698_v94;
          _nw301[(new BigNumber(18)).toNumber()] = _1698_v94;
          _nw301[(new BigNumber(19)).toNumber()] = _1697_v93;
          _1699_v95 = _nw301;
          let _1700_v96;
          _1700_v96 = _dafny.Map.Empty.slice().updateUnsafe(_1699_v95,(_this).f29);
          let _rhs321 = new BigNumber((_1700_v96).length);
          let _rhs322 = p0;
          let _lhs303 = globalState;
          let _lhs304 = globalState;
          _lhs303.f19 = _rhs321;
          _lhs304.f21 = _rhs322;
        } else if (_source24.is_DC9) {
          let _1701___mcc_h15 = (_source24).cf14;
          let _1702___mcc_h16 = (_source24).cf15;
          let _1703___mcc_h17 = (_source24).cf16;
          let _1704___mcc_h18 = (_source24).cf17;
          let _1705_cf17 = _1704___mcc_h18;
          let _1706_cf16 = _1703___mcc_h17;
          let _1707_cf15 = _1702___mcc_h16;
          let _1708_cf14 = _1701___mcc_h15;
          let _1709_v97;
          let _nw302 = new _module.C5();
          _nw302.__ctor((_this).f31, _this.f30);
          _1709_v97 = _nw302;
          (globalState).f13 = !(((_1545_v2).f33) === ((_this).f29));
          let _rhs323 = _1550_v7;
          let _rhs324 = _1705_cf17;
          let _rhs325 = (_1545_v2).f33;
          let _lhs305 = globalState;
          let _lhs306 = globalState;
          _lhs305.f17 = _rhs323;
          r2 = _rhs324;
          _lhs306.f18 = _rhs325;
          let _1710_v98;
          let _nw303 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.of());
          _1710_v98 = _nw303;
          _1710_v98 = _1710_v98;
        } else {
          let _1711___mcc_h19 = (_source24).cf10;
          let _1712_cf10 = _1711___mcc_h19;
          let _1713_v99;
          _1713_v99 = _dafny.Map.Empty.slice().updateUnsafe(((_1545_v2).f34).isLessThanOrEqualTo((_1545_v2).f34),p0);
          let _1714_v100;
          _1714_v100 = _dafny.Seq.of(_1542_v0, _1542_v0, _1542_v0);
          let _1715_v101;
          _1715_v101 = _module.D11.create_DC34(_module.__default.fm45(globalState));
          let _1716_v102;
          _1716_v102 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f32,p1), _1713_v99, _1713_v99, _1713_v99);
          let _1717_v103;
          _1717_v103 = _dafny.Map.Empty.slice().updateUnsafe((_1545_v2).f34,p0);
          let _pat_let_tv34 = _1679_cf56;
          let _rhs326 = (_1716_v102)[_module.__default.safeIndex((new BigNumber((_1717_v103).length)).plus((_dafny.ZERO).minus((_1545_v2).f34)), new BigNumber((_1716_v102).length))];
          let _rhs327 = _1714_v100;
          let _rhs328 = function (_pat_let33_0) {
            return function (_1718_dt__update__tmp_h0) {
              return function (_pat_let34_0) {
                return function (_1719_dt__update_hcf56_h0) {
                  return _module.D11.create_DC34(_1719_dt__update_hcf56_h0);
                }(_pat_let34_0);
              }(_pat_let_tv34);
            }(_pat_let33_0);
          }(_1715_v101);
          _1713_v99 = _rhs326;
          _1714_v100 = _rhs327;
          _1715_v101 = _rhs328;
          let _1720_v104;
          let _nw304 = new _module.C1();
          _nw304.__ctor((_1545_v2).f33);
          _1720_v104 = _nw304;
          (globalState).f2 = (_1545_v2).f34;
          let _1721_v105;
          _1721_v105 = _module.D3.create_DC11((_this).f31, (_this).f31);
          let _1722_v106;
          _1722_v106 = _module.D3.create_DC12(_1721_v105);
          let _1723_v107;
          _1723_v107 = _module.D3.create_DC12((_1722_v106).dtor_cf21);
          let _1724_v108;
          _1724_v108 = _dafny.Map.Empty.slice().updateUnsafe((_1720_v104).f40,_1723_v107);
          let _index273 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_1542_v0).length));
          (_1542_v0)[_index273] = true;
          let _1725_v109;
          let _nw305 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.of());
          _1725_v109 = _nw305;
          let _1726_v110;
          _1726_v110 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_1545_v2).f34));
          let _index274 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1725_v109).length));
          (_1725_v109)[_index274] = _1726_v110;
          let _1727_v111;
          _1727_v111 = _dafny.Set.fromElements((_1545_v2).f33);
          let _1728_v112;
          let _nw306 = Array((new BigNumber(7)).toNumber());
          _nw306[(_dafny.ZERO).toNumber()] = _1727_v111;
          _nw306[(_dafny.ONE).toNumber()] = (_1727_v111).Difference(_1727_v111);
          _nw306[(new BigNumber(2)).toNumber()] = (((_1720_v104).f40) ? (_1727_v111) : (_module.__default.fm22(_module.__default.fm36(globalState), _module.__default.fm33((_this).f29, (_1545_v2).f34, p1, globalState), _1547_v4, globalState)));
          _nw306[(new BigNumber(3)).toNumber()] = (_1727_v111).Union(_1727_v111);
          _nw306[(new BigNumber(4)).toNumber()] = _1727_v111;
          _nw306[(new BigNumber(5)).toNumber()] = _1727_v111;
          _nw306[(new BigNumber(6)).toNumber()] = _module.__default.fm22((_1545_v2).f34, (_this).f31, _1547_v4, globalState);
          _1728_v112 = _nw306;
          let _index275 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_1728_v112).length));
          (_1728_v112)[_index275] = _1727_v111;
          let _index276 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_1542_v0).length));
          let _index277 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1725_v109).length));
          let _index278 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_1728_v112).length));
          let _rhs329 = _module.__default.fm53(!(_module.__default.fm6((_1545_v2).f34, globalState)), (_1545_v2).f34, _this.f30, p0, globalState);
          let _rhs330 = (_1545_v2).f33;
          let _rhs331 = _1726_v110;
          let _rhs332 = new BigNumber(384);
          let _rhs333 = _module.__default.fm22((((_this).f31) ? (new BigNumber((_1727_v111).length)) : ((_1545_v2).f34)), (_1545_v2).f33, _1547_v4, globalState);
          let _lhs307 = _1542_v0;
          let _lhs308 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_1542_v0).length));
          let _lhs309 = _1725_v109;
          let _lhs310 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1725_v109).length));
          let _lhs311 = globalState;
          let _lhs312 = _1728_v112;
          let _lhs313 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_1728_v112).length));
          _1724_v108 = _rhs329;
          _lhs307[_lhs308] = _rhs330;
          _lhs309[_lhs310] = _rhs331;
          _lhs311.f2 = _rhs332;
          _lhs312[_lhs313] = _rhs333;
        }
        let _index279 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1542_v0).length));
        (_1542_v0)[_index279] = (_1545_v2).f33;
        let _1729_v113;
        _1729_v113 = _dafny.Seq.of((_this).f29);
        let _index280 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1542_v0).length));
        (_1542_v0)[_index280] = ((_1545_v2).f33) && (_dafny.Seq.contains(_1729_v113, !((_this).f29)));
        (globalState).f15 = (_module.D17.create_DC52((_1542_v0)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_1542_v0).length))], (_1545_v2).f33, (_1545_v2).f34, _1554_v11)).dtor_cf89;
        r2 = (_1545_v2).fm5(_dafny.Set.fromElements((_this).f29), globalState);
      }
      let _1730_v114;
      let _nw307 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
      _1730_v114 = _nw307;
      let _1731_v115;
      _1731_v115 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_1542_v0);
      let _index281 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_1730_v114).length));
      (_1730_v114)[_index281] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f31,_1542_v0)).Merge(_1731_v115);
      let _1732_v116;
      _1732_v116 = _dafny.Set.fromElements(_this.f30);
      let _1733_v117;
      _1733_v117 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("rjcvfdu"), _dafny.Seq.UnicodeFromString("lnab"));
      let _index282 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_1730_v114).length));
      let _rhs334 = (_1545_v2).fm5(_dafny.Set.fromElements(false), globalState);
      let _rhs335 = (_1545_v2).f33;
      let _rhs336 = _1731_v115;
      let _rhs337 = _dafny.Seq.Concat(_dafny.Seq.Concat(_this.f30, (_1733_v117)[_module.__default.safeIndex(new BigNumber(-811), new BigNumber((_1733_v117).length))]), _dafny.Seq.Concat(_this.f30, _this.f30));
      let _rhs338 = _1732_v116;
      let _lhs314 = globalState;
      let _lhs315 = globalState;
      let _lhs316 = _1730_v114;
      let _lhs317 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_1730_v114).length));
      let _lhs318 = _this;
      _lhs314.f19 = _rhs334;
      _lhs315.f18 = _rhs335;
      _lhs316[_lhs317] = _rhs336;
      _lhs318.f30 = _rhs337;
      _1732_v116 = _rhs338;
      if ((_this).f29) {
        let _1734_v118;
        _1734_v118 = _dafny.MultiSet.fromElements(_1542_v0);
        if (!((_1734_v118).Union(_1734_v118)).equals((_1734_v118).update(_1542_v0, _module.__default.abs(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1547_v4,(_this).f32)).length))))) {
          (globalState).f20 = _this.f30;
          let _1735_v119;
          let _nw308 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Map.Empty);
          _1735_v119 = _nw308;
          let _1736_v120;
          let _init46 = function (_1737_i17) {
            return _this.f30;
          };
          let _nw309 = Array((new BigNumber(16)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw309.length); _i0_46++) {
            _nw309[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1736_v120 = _nw309;
          let _1738_v121;
          _1738_v121 = _dafny.Seq.of(p1, p0);
          let _1739_v122;
          _1739_v122 = _dafny.Map.Empty.slice().updateUnsafe(_1736_v120,_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(360)), ((_1740_v4) => function (_1741_i18) {
            return _1740_v4;
          })(_1547_v4)),new BigNumber((_1738_v121).length)));
          let _1742_v123;
          _1742_v123 = _module.D15.create_DC45((_1545_v2).f34, p1, (_1545_v2).f33, (_this).f31, _1736_v120);
          let _1743_v125;
          _1743_v125 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,new BigNumber((_this.f30).length));
          let _index283 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1735_v119).length));
          (_1735_v119)[_index283] = (((_1739_v122).contains((_1742_v123).dtor_cf80)) ? ((_1739_v122).get((_1742_v123).dtor_cf80)) : (function () {
            let _coll54 = new _dafny.Map();
            for (const _compr_54 of (_1743_v125).Keys.Elements) {
              let _1744_v124 = _compr_54;
              if ((_1743_v125).contains(_1744_v124)) {
                _coll54.push([_1744_v124,p0]);
              }
            }
            return _coll54;
          }()));
          let _index284 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1735_v119).length));
          (_1735_v119)[_index284] = _1743_v125;
          (globalState).f28 = (_1545_v2).f33;
          let _1745_v126;
          _1745_v126 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1545_v2).f33,(_1545_v2).f33)).length),_1542_v0);
          let _1746_v127;
          let _nw310 = new _module.C4();
          _nw310.__ctor(_module.__default.fm33((_this).f32, (_1545_v2).f34, p0, globalState), new BigNumber(((_1745_v126).Merge(_1745_v126)).length));
          _1746_v127 = _nw310;
          _1542_v0 = _1542_v0;
        } else {
          let _1747_v128;
          _1747_v128 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f29);
          let _1748_v129;
          _1748_v129 = _dafny.Seq.of((_this).f31);
          let _1749_v130;
          let _nw311 = new _module.C5();
          _nw311.__ctor(false, _module.__default.fm7(p0, globalState));
          _1749_v130 = _nw311;
          let _1750_v131;
          _1750_v131 = _module.D17.create_DC50((((_1747_v128).contains((_1748_v129)[_module.__default.safeIndex((_1545_v2).f34, new BigNumber((_1748_v129).length))])) ? ((_1747_v128).get((_1748_v129)[_module.__default.safeIndex((_1545_v2).f34, new BigNumber((_1748_v129).length))])) : ((_this).f31)), true, _1749_v130);
          _1750_v131 = _module.D17.create_DC50((_1545_v2).f33, (true) && ((_1545_v2).f33), _1749_v130);
          _1748_v129 = _1748_v129;
          (globalState).f13 = true;
          let _1751_v133;
          _1751_v133 = _dafny.Map.Empty.slice().updateUnsafe((_1545_v2).f34,p1);
          let _1752_v134;
          _1752_v134 = _dafny.Seq.of((_1545_v2).f34, (_dafny.ZERO).minus((_1545_v2).f34), _module.__default.fm36(globalState), (((_1751_v133).contains(p0)) ? ((_1751_v133).get(p0)) : ((_1545_v2).f34)));
          (globalState).f15 = !(_1551_v8).equals((function () {
            let _coll55 = new _dafny.Map();
            for (const _compr_55 of (_1752_v134).Elements) {
              let _1753_v132 = _compr_55;
              if (_dafny.Seq.contains(_1752_v134, _1753_v132)) {
                _coll55.push([_module.__default.safeDivisionInt(_1753_v132, p1),(_this).f29]);
              }
            }
            return _coll55;
          }()).Merge(_1551_v8));
          let _1754_v135;
          let _nw312 = Array((new BigNumber(27)).toNumber());
          _nw312[(_dafny.ZERO).toNumber()] = _1550_v7;
          _nw312[(_dafny.ONE).toNumber()] = (_1552_v9).dtor_cf74;
          _nw312[(new BigNumber(2)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(3)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(4)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(5)).toNumber()] = (((_1545_v2).f33) ? (_1550_v7) : (_1550_v7));
          _nw312[(new BigNumber(6)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(7)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(8)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(9)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(10)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(11)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(12)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(13)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(14)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(15)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(16)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(17)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(18)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(19)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(20)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(21)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(22)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(23)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(24)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(25)).toNumber()] = _1550_v7;
          _nw312[(new BigNumber(26)).toNumber()] = _1550_v7;
          _1754_v135 = _nw312;
          let _index285 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_1754_v135).length));
          (_1754_v135)[_index285] = _1550_v7;
          let _index286 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_1754_v135).length));
          (_1754_v135)[_index286] = (((_1545_v2).f33) ? (_1550_v7) : ((((_1545_v2).f33) ? (_1550_v7) : (_1550_v7))));
        }
        let _1755_v136;
        _1755_v136 = _dafny.Map.Empty.slice().updateUnsafe(_1550_v7,(_1545_v2).f33);
        let _1756_v137;
        _1756_v137 = _dafny.Seq.of(_1550_v7);
        let _1757_v138;
        _1757_v138 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_1545_v2).f34);
        _1755_v136 = (_1755_v136).update((_1756_v137)[_module.__default.safeIndex((((_1757_v138).contains((_this).f32)) ? ((_1757_v138).get((_this).f32)) : (p0)), new BigNumber((_1756_v137).length))], (true) || (!((_this).f32)));
        (globalState).f11 = _this.f30;
        let _1758_v139;
        _1758_v139 = _dafny.Set.fromElements((_this).f29, (_this).f29, (_this).f31, (_this).f31, (_1545_v2).f33);
        let _index287 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_1550_v7).length));
        (_1550_v7)[_index287] = (_1545_v2).fm5(_1758_v139, globalState);
        let _1759_v140;
        _1759_v140 = _dafny.Seq.of(_module.__default.fm9(globalState));
        let _index288 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_1550_v7).length));
        let _rhs339 = (_this).f32;
        let _rhs340 = (((_this).f32) ? (new BigNumber((_1759_v140).length)) : ((_1545_v2).f34));
        let _lhs319 = globalState;
        let _lhs320 = _1550_v7;
        let _lhs321 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_1550_v7).length));
        _lhs319.f28 = _rhs339;
        _lhs320[_lhs321] = _rhs340;
        if ((((_this).f31) ? ((_1548_v5).dtor_cf53) : ((_this).f29))) {
          _1551_v8 = (_1551_v8).update(p0, (_1545_v2).f33);
          (globalState).f18 = (_1545_v2).f33;
          let _1760_v141;
          let _nw313 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _1760_v141 = _nw313;
          let _1761_v142;
          _1761_v142 = _module.D17.create_DC51(new BigNumber((_this.f30).length), _1760_v141);
          let _1762_v143;
          let _nw314 = Array((new BigNumber(13)).toNumber());
          _nw314[(_dafny.ZERO).toNumber()] = _1760_v141;
          _nw314[(_dafny.ONE).toNumber()] = (_1761_v142).dtor_cf87;
          _nw314[(new BigNumber(2)).toNumber()] = _1760_v141;
          _nw314[(new BigNumber(3)).toNumber()] = _1760_v141;
          _nw314[(new BigNumber(4)).toNumber()] = _1760_v141;
          _nw314[(new BigNumber(5)).toNumber()] = (_1761_v142).dtor_cf87;
          _nw314[(new BigNumber(6)).toNumber()] = _1760_v141;
          _nw314[(new BigNumber(7)).toNumber()] = _1760_v141;
          _nw314[(new BigNumber(8)).toNumber()] = _1760_v141;
          _nw314[(new BigNumber(9)).toNumber()] = _1760_v141;
          _nw314[(new BigNumber(10)).toNumber()] = _1760_v141;
          _nw314[(new BigNumber(11)).toNumber()] = _1760_v141;
          _nw314[(new BigNumber(12)).toNumber()] = _1760_v141;
          _1762_v143 = _nw314;
          let _1763_v144;
          _1763_v144 = _dafny.Seq.of(_1762_v143);
          _1762_v143 = (_1763_v144)[_module.__default.safeIndex((_1545_v2).f34, new BigNumber((_1763_v144).length))];
          let _index289 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_1550_v7).length));
          let _rhs341 = (_1550_v7)[_module.__default.safeIndex(new BigNumber(200), new BigNumber((_1550_v7).length))];
          let _rhs342 = !((_1545_v2).f33);
          let _rhs343 = _module.__default.safeDivisionInt((_1545_v2).f34, (p0).minus((_1545_v2).f34));
          let _lhs322 = globalState;
          let _lhs323 = globalState;
          let _lhs324 = _1550_v7;
          let _lhs325 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_1550_v7).length));
          _lhs322.f19 = _rhs341;
          _lhs323.f18 = _rhs342;
          _lhs324[_lhs325] = _rhs343;
          let _1764_v145;
          _1764_v145 = _module.D0.create_DC1(_1542_v0, (_1545_v2).f34, (_1545_v2).f34);
          let _1765_v146;
          _1765_v146 = _module.D0.create_DC1((_1764_v145).dtor_cf1, p0, new BigNumber((_this.f30).length));
          _1764_v145 = _1765_v146;
        } else {
          let _1766_v147;
          _1766_v147 = _dafny.Seq.of((_1545_v2).f33, ((_1545_v2).f33) === ((_1545_v2).f33));
          _1766_v147 = _1766_v147;
          let _1767_v148;
          let _init47 = ((_1768_v140) => function (_1769_i19) {
            return _dafny.Seq.Concat(_1768_v140, _1768_v140);
          })(_1759_v140);
          let _nw315 = Array((new BigNumber(10)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw315.length); _i0_47++) {
            _nw315[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1767_v148 = _nw315;
          let _index290 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_1767_v148).length));
          (_1767_v148)[_index290] = _1759_v140;
          let _index291 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_1550_v7).length));
          let _index292 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_1767_v148).length));
          let _rhs344 = (_this).f32;
          let _rhs345 = p1;
          let _rhs346 = (((_1545_v2).f33) ? (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(652)), ((_1770_p1) => function (_1771_i20) {
            return _1770_p1;
          })(p1)), _1759_v140)) : (_dafny.Seq.Concat(_1759_v140, _1759_v140)));
          let _rhs347 = (_1545_v2).f33;
          let _lhs326 = globalState;
          let _lhs327 = _1550_v7;
          let _lhs328 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_1550_v7).length));
          let _lhs329 = _1767_v148;
          let _lhs330 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_1767_v148).length));
          let _lhs331 = globalState;
          _lhs326.f15 = _rhs344;
          _lhs327[_lhs328] = _rhs345;
          _lhs329[_lhs330] = _rhs346;
          _lhs331.f18 = _rhs347;
          (globalState).f2 = (_dafny.ZERO).minus((_1550_v7)[_module.__default.safeIndex(new BigNumber(200), new BigNumber((_1550_v7).length))]);
          let _1772_v149;
          let _nw316 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1772_v149 = _nw316;
          let _index293 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1772_v149).length));
          (_1772_v149)[_index293] = (_1548_v5).dtor_cf52;
          let _index294 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1772_v149).length));
          (_1772_v149)[_index294] = _1547_v4;
          (globalState).f20 = _dafny.Seq.Concat(_this.f30, _this.f30);
        }
      } else {
        _1547_v4 = _1547_v4;
        (globalState).f15 = (((((_1551_v8).contains((_1545_v2).f34)) ? ((_1551_v8).get((_1545_v2).f34)) : ((_1545_v2).f33))) ? ((_this).f32) : ((_1545_v2).f33));
        if ((_1545_v2).f33) {
          (globalState).f2 = (_1545_v2).f34;
          (globalState).f1 = (p0).multipliedBy((_1545_v2).fm5(_module.__default.fm22(new BigNumber(641), (_this).f31, _1547_v4, globalState), globalState));
          (globalState).f21 = (_1545_v2).f34;
          (globalState).f21 = ((_dafny.ZERO).minus((_1545_v2).f34)).multipliedBy((_1545_v2).f34);
          let _1773_v150;
          let _nw317 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
          _1773_v150 = _nw317;
          let _1774_v151;
          _1774_v151 = _dafny.Seq.of(new BigNumber(970), (_1545_v2).f34);
          let _index295 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_1773_v150).length));
          (_1773_v150)[_index295] = _1774_v151;
          let _index296 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_1773_v150).length));
          let _rhs348 = _1774_v151;
          let _rhs349 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1774_v151, _dafny.Seq.of(new BigNumber((_this.f30).length), new BigNumber(965), (_1774_v151)[_module.__default.safeIndex((_1545_v2).f34, new BigNumber((_1774_v151).length))], (_1545_v2).f34, _module.__default.fm36(globalState))), _1774_v151);
          let _lhs332 = globalState;
          let _lhs333 = _1773_v150;
          let _lhs334 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_1773_v150).length));
          _lhs332.f7 = _rhs348;
          _lhs333[_lhs334] = _rhs349;
        } else {
          let _1775_v152;
          _1775_v152 = _dafny.Seq.of((_this).f31);
          r0 = ((_1775_v152)[_module.__default.safeIndex(new BigNumber(302), new BigNumber((_1775_v152).length))]) === ((_this).f32);
          let _1776_v153;
          _1776_v153 = _dafny.Seq.of(new BigNumber((_this.f30).length));
          let _1777_v154;
          _1777_v154 = _dafny.Map.Empty.slice().updateUnsafe(((_1545_v2).f33) && (_module.__default.fm33((_1545_v2).f33, new BigNumber((_1776_v153).length), (_1545_v2).f34, globalState)),(_this).f32);
          _1777_v154 = (_1777_v154).update((_1545_v2).f33, !(((_1545_v2).f34).isLessThan((_1545_v2).f34)));
          let _1778_v155;
          let _init48 = ((_1779_v8, _1780_v2) => function (_1781_i21) {
            return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1779_v8).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(194)), ((_1782_v2) => function (_1783_i22) {
              return (_1782_v2).f34;
            })(_1780_v2))).length));
          })(_1551_v8, _1545_v2);
          let _nw318 = Array((new BigNumber(2)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw318.length); _i0_48++) {
            _nw318[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _1778_v155 = _nw318;
          let _1784_v156;
          _1784_v156 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("tlk")).length),new BigNumber((_1554_v11).cardinality()));
          let _1785_v158;
          _1785_v158 = (_1784_v156).Merge(function () {
            let _coll56 = new _dafny.Map();
            for (const _compr_56 of _dafny.IntegerRange(new BigNumber(18), new BigNumber(970))) {
              let _1786_v157 = _compr_56;
              if (((new BigNumber(18)).isLessThanOrEqualTo(_1786_v157)) && ((_1786_v157).isLessThan(new BigNumber(970)))) {
                _coll56.push([(_1786_v157).plus(p1),new BigNumber(716)]);
              }
            }
            return _coll56;
          }());
          let _rhs350 = _1778_v155;
          let _rhs351 = _this.f30;
          let _rhs352 = !((_this).f29);
          let _rhs353 = (((_this).f32) ? (_1785_v158) : (_1785_v158));
          let _rhs354 = (((_this).f31) ? ((_1545_v2).f34) : ((_1545_v2).f34));
          let _lhs335 = globalState;
          let _lhs336 = globalState;
          let _lhs337 = globalState;
          _1778_v155 = _rhs350;
          _lhs335.f11 = _rhs351;
          _lhs336.f13 = _rhs352;
          _1785_v158 = _rhs353;
          _lhs337.f14 = _rhs354;
          (globalState).f15 = (((_this).f31) ? (false) : (_module.__default.fm6((_1545_v2).f34, globalState)));
          let _1787_v159;
          let _init49 = function (_1788_i23) {
            return _dafny.Seq.UnicodeFromString("suupvv");
          };
          let _nw319 = Array((new BigNumber(17)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw319.length); _i0_49++) {
            _nw319[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _1787_v159 = _nw319;
          let _index297 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_1787_v159).length));
          (_1787_v159)[_index297] = (_1733_v117)[_module.__default.safeIndex(new BigNumber((_1776_v153).length), new BigNumber((_1733_v117).length))];
          let _index298 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_1787_v159).length));
          let _rhs355 = _dafny.Seq.UnicodeFromString("y");
          let _rhs356 = _module.__default.fm6(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(569)), ((_1789_p0) => function (_1790_i24) {
            return _1789_p0;
          })(p0))).length), globalState);
          let _rhs357 = _this.f30;
          let _rhs358 = (_1775_v152)[_module.__default.safeIndex(p0, new BigNumber((_1775_v152).length))];
          let _lhs338 = globalState;
          let _lhs339 = globalState;
          let _lhs340 = _1787_v159;
          let _lhs341 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_1787_v159).length));
          let _lhs342 = globalState;
          _lhs338.f20 = _rhs355;
          _lhs339.f28 = _rhs356;
          _lhs340[_lhs341] = _rhs357;
          _lhs342.f15 = _rhs358;
        }
        let _1791_v160;
        _1791_v160 = _dafny.Seq.of(p1, (_1545_v2).f34);
        let _index299 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_1550_v7).length));
        (_1550_v7)[_index299] = (((_this).f29) ? (p1) : ((_1791_v160)[_module.__default.safeIndex(p0, new BigNumber((_1791_v160).length))]));
        let _1792_v161;
        _1792_v161 = _dafny.Set.fromElements((_1545_v2).f34, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(609)), ((_1793_v4) => function (_1794_i25) {
          return _1793_v4;
        })(_1547_v4))).length)), (_1545_v2).f34);
        let _index300 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_1550_v7).length));
        (_1550_v7)[_index300] = _module.__default.safeDivisionInt((_1545_v2).f34, new BigNumber((_1792_v161).length));
        let _1795_v162;
        let _nw320 = new _module.C4();
        _nw320.__ctor((_this).f29, (_1545_v2).f34);
        _1795_v162 = _nw320;
      }
      let _1796_v163;
      _1796_v163 = _dafny.MultiSet.fromElements((_this).f29, false, (_this).f32);
      if ((((false) ? (_1796_v163) : (_1796_v163))).IsProperSubsetOf((_1796_v163).Intersect(_1796_v163))) {
        r2 = (_1545_v2).f34;
        let _1797_v164;
        let _nw321 = new _module.C3();
        _nw321.__ctor(!(((_1545_v2).f33) && ((_this).f32)), p0, false, _this.f30);
        _1797_v164 = _nw321;
        let _1798_v165;
        _1798_v165 = _module.D2.create_DC8(_1547_v4, (_1545_v2).f33);
        let _1799_v166;
        let _nw322 = Array((new BigNumber(20)).toNumber());
        _nw322[(_dafny.ZERO).toNumber()] = _1547_v4;
        _nw322[(_dafny.ONE).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(2)).toNumber()] = _module.__default.fm42((_1545_v2).f33, !((_this).f32), globalState);
        _nw322[(new BigNumber(3)).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
        _nw322[(new BigNumber(5)).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(6)).toNumber()] = _module.__default.fm42((_this).f31, false, globalState);
        _nw322[(new BigNumber(7)).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(8)).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(9)).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(10)).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(11)).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(12)).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(13)).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(14)).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(15)).toNumber()] = (_1798_v165).dtor_cf12;
        _nw322[(new BigNumber(16)).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(17)).toNumber()] = new _dafny.CodePoint('m'.codePointAt(0));
        _nw322[(new BigNumber(18)).toNumber()] = _1547_v4;
        _nw322[(new BigNumber(19)).toNumber()] = _module.__default.fm42((_this).f31, (_this).f32, globalState);
        _1799_v166 = _nw322;
        let _index301 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_1799_v166).length));
        (_1799_v166)[_index301] = _1547_v4;
        let _index302 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_1799_v166).length));
        (_1799_v166)[_index302] = _1547_v4;
        let _1800_v167;
        _1800_v167 = _dafny.Seq.of((_this).f29);
        let _1801_v168;
        _1801_v168 = _dafny.Map.Empty.slice().updateUnsafe((_1799_v166)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_1799_v166).length))],new BigNumber((_1800_v167).length));
        let _1802_v169;
        _1802_v169 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_this.f30, _this.f30),(((_1801_v168).contains(_1547_v4)) ? ((_1801_v168).get(_1547_v4)) : (p1)));
        let _1803_v170;
        let _nw323 = new _module.C6();
        _nw323.__ctor(true, _this.f30);
        _1803_v170 = _nw323;
        let _1804_v171;
        _1804_v171 = _dafny.Seq.of(_1803_v170);
        let _1805_v172;
        _1805_v172 = _dafny.Seq.of((_1797_v164).f36, new BigNumber((_dafny.Seq.update(_1804_v171, _module.__default.safeIndex(new BigNumber((_this.f30).length), new BigNumber((_1804_v171).length)), _1803_v170)).length));
        _1802_v169 = ((_1802_v169).Merge(_1802_v169)).Merge((_1802_v169).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f30,new BigNumber((_1805_v172).length))));
        let _1806_v173;
        _1806_v173 = _dafny.Map.Empty.slice().updateUnsafe((p0).multipliedBy(p1),(_1545_v2).f34);
        let _pat_let_tv35 = _1547_v4;
        let _1807_v174;
        _1807_v174 = _dafny.Seq.of(_module.D2.create_DC8((_1799_v166)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_1799_v166).length))], (_1797_v164).f35), function (_pat_let35_0) {
          return function (_1808_dt__update__tmp_h1) {
            return function (_pat_let36_0) {
              return function (_1809_dt__update_hcf12_h0) {
                return _module.D2.create_DC8(_1809_dt__update_hcf12_h0, (_1808_dt__update__tmp_h1).dtor_cf13);
              }(_pat_let36_0);
            }(_pat_let_tv35);
          }(_pat_let35_0);
        }(_module.D2.create_DC8((_1799_v166)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_1799_v166).length))], true)));
        _1806_v173 = (_1806_v173).update(_module.__default.safeModuloInt((_1545_v2).f34, new BigNumber((_module.__default.fm35(_1807_v174, (_1797_v164).f36, globalState)).length)), new BigNumber(980));
      } else {
        let _1810_v175;
        _1810_v175 = _dafny.Seq.of(p0);
        let _1811_v176;
        _1811_v176 = _dafny.Seq.of(new BigNumber(265), p0, (_dafny.ZERO).minus(new BigNumber((_1810_v175).length)));
        (globalState).f15 = !((_this).f29) || (_dafny.Seq.contains(_1811_v176, (_1545_v2).f34));
        let _1812_v177;
        let _nw324 = Array((new BigNumber(10)).toNumber()).fill([]);
        _1812_v177 = _nw324;
        let _index303 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length));
        (_1812_v177)[_index303] = _1550_v7;
        let _1813_v178;
        _1813_v178 = _module.D10.create_DC33((_this).f31, (_1545_v2).f34);
        let _index304 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length));
        let _rhs359 = (((_1545_v2).f33) ? (_1550_v7) : (_1550_v7));
        let _rhs360 = _this.f30;
        let _rhs361 = (_1813_v178).dtor_cf54;
        let _lhs343 = _1812_v177;
        let _lhs344 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length));
        let _lhs345 = globalState;
        let _lhs346 = globalState;
        _lhs343[_lhs344] = _rhs359;
        _lhs345.f11 = _rhs360;
        _lhs346.f15 = _rhs361;
        let _arr2 = (_1812_v177)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length))];
        let _index305 = _module.__default.safeIndex(new BigNumber(471), new BigNumber(((_1812_v177)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length))]).length));
        _arr2[_index305] = (_dafny.ZERO).minus(new BigNumber((_this.f30).length));
        let _1814_v180;
        _1814_v180 = _module.D11.create_DC36(_module.D6.create_DC23(true, (_1545_v2).f34), _dafny.Set.fromElements(false, (_this).f31), (_1545_v2).f33);
        let _1815_v181;
        _1815_v181 = _dafny.Seq.of(_1550_v7);
        let _1816_v182;
        _1816_v182 = _module.D6.create_DC23((_1545_v2).f33, new BigNumber((_1815_v181).length));
        let _pat_let_tv36 = _1816_v182;
        let _1817_v183;
        _1817_v183 = _dafny.Map.Empty.slice().updateUnsafe((_1545_v2).f33,function (_pat_let37_0) {
          return function (_1818_dt__update__tmp_h2) {
            return function (_pat_let38_0) {
              return function (_1819_dt__update_hcf59_h0) {
                return _module.D11.create_DC36(_1819_dt__update_hcf59_h0, (_1818_dt__update__tmp_h2).dtor_cf60, (_1818_dt__update__tmp_h2).dtor_cf61);
              }(_pat_let38_0);
            }(_pat_let_tv36);
          }(_pat_let37_0);
        }(_1814_v180));
        let _1820_v184;
        _1820_v184 = _dafny.Set.fromElements(_1817_v183, _1817_v183);
        let _arr3 = (_1812_v177)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length))];
        let _index306 = _module.__default.safeIndex(new BigNumber(471), new BigNumber(((_1812_v177)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length))]).length));
        let _rhs362 = (_1545_v2).f34;
        let _rhs363 = (function () {
          let _coll57 = new _dafny.Set();
          for (const _compr_57 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), ((_1821_v2) => function (_1822_i26) {
            return _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_module.D11.create_DC36(_module.D6.create_DC23((_1821_v2).f33, new BigNumber((_dafny.Seq.of((_1821_v2).f33)).length)), _dafny.Set.fromElements(false), (_this).f29));
          })(_1545_v2))).Elements) {
            let _1823_v179 = _compr_57;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), ((_1824_v2) => function (_1822_i26) {
              return _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_module.D11.create_DC36(_module.D6.create_DC23((_1824_v2).f33, new BigNumber((_dafny.Seq.of((_1824_v2).f33)).length)), _dafny.Set.fromElements(false), (_this).f29));
            })(_1545_v2)), _1823_v179)) {
              _coll57.add(_1823_v179);
            }
          }
          return _coll57;
        }()).equals(_1820_v184);
        let _lhs347 = (_1812_v177)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length))];
        let _lhs348 = _module.__default.safeIndex(new BigNumber(471), new BigNumber(((_1812_v177)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length))]).length));
        let _lhs349 = globalState;
        _lhs347[_lhs348] = _rhs362;
        _lhs349.f13 = _rhs363;
        let _1825_v185;
        let _nw325 = new _module.C5();
        _nw325.__ctor((_1545_v2).f33, _this.f30);
        _1825_v185 = _nw325;
        let _1826_v186;
        let _nw326 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
        _1826_v186 = _nw326;
        let _1827_v187;
        _1827_v187 = _dafny.Map.Empty.slice().updateUnsafe(_1796_v163,p0);
        let _index307 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_1826_v186).length));
        (_1826_v186)[_index307] = _1827_v187;
        let _1828_v188;
        let _nw327 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Set.Empty);
        _1828_v188 = _nw327;
        let _1829_v189;
        _1829_v189 = _dafny.Map.Empty.slice().updateUnsafe((_1545_v2).f34,((_1812_v177)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length))])[_module.__default.safeIndex(new BigNumber(471), new BigNumber(((_1812_v177)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length))]).length))]);
        let _1830_v190;
        _1830_v190 = _dafny.Set.fromElements(_1829_v189);
        let _index308 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_1828_v188).length));
        (_1828_v188)[_index308] = _1830_v190;
        let _1831_v191;
        _1831_v191 = _dafny.Seq.of(_1825_v185, _1825_v185);
        let _arr4 = (_1812_v177)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length))];
        let _index309 = _module.__default.safeIndex(new BigNumber(471), new BigNumber(((_1812_v177)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length))]).length));
        let _index310 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_1826_v186).length));
        let _index311 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_1828_v188).length));
        let _rhs364 = new BigNumber(493);
        let _rhs365 = (_1831_v191)[_module.__default.safeIndex(p0, new BigNumber((_1831_v191).length))];
        let _rhs366 = _1827_v187;
        let _rhs367 = _1830_v190;
        let _lhs350 = (_1812_v177)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length))];
        let _lhs351 = _module.__default.safeIndex(new BigNumber(471), new BigNumber(((_1812_v177)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1812_v177).length))]).length));
        let _lhs352 = _1826_v186;
        let _lhs353 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_1826_v186).length));
        let _lhs354 = _1828_v188;
        let _lhs355 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_1828_v188).length));
        _lhs350[_lhs351] = _rhs364;
        _1825_v185 = _rhs365;
        _lhs352[_lhs353] = _rhs366;
        _lhs354[_lhs355] = _rhs367;
        let _1832_v192;
        let _nw328 = new _module.C1();
        _nw328.__ctor((_this).f31);
        _1832_v192 = _nw328;
      }
      let _1833_i27;
      _1833_i27 = _dafny.ZERO;
      L8: {
        while ((_this).f31) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1833_i27)) {
              break L8;
            }
            _1833_i27 = (_1833_i27).plus(_dafny.ONE);
            (globalState).f18 = _dafny.Seq.contains(_this.f30, _1547_v4);
            (globalState).f18 = (_this).f31;
            let _1834_v193;
            _1834_v193 = _dafny.Seq.of((_1545_v2).f33);
            (globalState).f18 = !(_dafny.Seq.contains(_1834_v193, (_1545_v2).f33));
            if ((_1545_v2).f33) {
              let _1835_v194;
              let _nw329 = new _module.C3();
              _nw329.__ctor((_this).f31, (_1545_v2).f34, (_this).f31, _dafny.Seq.Create(_module.__default.abs(new BigNumber(784)), ((_1836_v4) => function (_1837_i28) {
                return _1836_v4;
              })(_1547_v4)));
              _1835_v194 = _nw329;
              _1835_v194 = _1835_v194;
              _1542_v0 = _1542_v0;
              let _index312 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_1542_v0).length));
              (_1542_v0)[_index312] = (p1).isLessThanOrEqualTo(new BigNumber((_1544_v1).length));
              let _index313 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_1542_v0).length));
              (_1542_v0)[_index313] = (((!((_this).f32)) ? ((_dafny.ZERO).minus(p1)) : ((_1545_v2).f34))).isLessThanOrEqualTo((_dafny.ZERO).minus((((_1545_v2).f33) ? ((_1835_v194).f36) : (new BigNumber((_this.f30).length)))));
              (globalState).f1 = p1;
              let _1838_v195;
              _1838_v195 = _module.D11.create_DC35(_this.f30, (_this).f32);
              let _1839_v196;
              _1839_v196 = _module.D2.create_DC9((_1838_v195).dtor_cf57, (_1835_v194).f36, (_1545_v2).f33, new BigNumber(((_1554_v11).update((_1545_v2).f34, _module.__default.abs((_1835_v194).f36))).cardinality()));
              let _1840_v197;
              _1840_v197 = _dafny.MultiSet.fromElements(_1796_v163);
              let _1841_v198;
              _1841_v198 = _dafny.Seq.of(p1);
              _1839_v196 = _module.__default.fm18((((_1796_v163).contains((_this).f31)) ? ((_1796_v163).get((_this).f31)) : ((_1545_v2).f34)), (_dafny.ZERO).minus(((_1545_v2).f34).plus(new BigNumber((_1840_v197).cardinality()))), (_1545_v2).f34, new BigNumber((_dafny.Seq.Concat(_1841_v198, _1841_v198)).length), globalState);
            } else {
              let _1842_v200;
              _1842_v200 = _dafny.Set.fromElements(p1);
              let _1843_v201;
              _1843_v201 = _dafny.Map.Empty.slice().updateUnsafe((_1545_v2).f33,(function () {
                let _coll58 = new _dafny.Set();
                for (const _compr_58 of _dafny.IntegerRange(new BigNumber(555), new BigNumber(896))) {
                  let _1844_v199 = _compr_58;
                  if (((new BigNumber(555)).isLessThanOrEqualTo(_1844_v199)) && ((_1844_v199).isLessThan(new BigNumber(896)))) {
                    _coll58.add(_module.__default.safeDivisionInt(_1844_v199, p0));
                  }
                }
                return _coll58;
              }()).IsDisjointFrom(_1842_v200));
              let _1845_v202;
              _1845_v202 = _dafny.Set.fromElements((_1545_v2).f33);
              _1843_v201 = (_1843_v201).update(!(_1845_v202).contains((_this).f31), (_this).f31);
              let _1846_v203;
              _1846_v203 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,(_1834_v193)[_module.__default.safeIndex((_1545_v2).f34, new BigNumber((_1834_v193).length))]);
              let _1847_v204;
              _1847_v204 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_1545_v2).f33,(_1545_v2).f34));
              let _rhs368 = (((_1847_v204).IsProperSubsetOf(_1847_v204)) ? (_1550_v7) : (((false) ? (_1550_v7) : (_1550_v7))));
              let _rhs369 = function () {
                let _coll59 = new _dafny.Map();
                for (const _compr_59 of (_1733_v117).Elements) {
                  let _1848_v205 = _compr_59;
                  if (_dafny.Seq.contains(_1733_v117, _1848_v205)) {
                    _coll59.push([_1848_v205,!((_1545_v2).f33)]);
                  }
                }
                return _coll59;
              }();
              let _lhs356 = globalState;
              _lhs356.f17 = _rhs368;
              _1846_v203 = _rhs369;
              let _index314 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1550_v7).length));
              (_1550_v7)[_index314] = (_1545_v2).f34;
              let _index315 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1550_v7).length));
              let _rhs370 = _1547_v4;
              let _rhs371 = (p1).plus((_1545_v2).f34);
              let _rhs372 = !_dafny.areEqual(_this.f30, _dafny.Seq.UnicodeFromString("jtnaivbv"));
              let _lhs357 = _1550_v7;
              let _lhs358 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1550_v7).length));
              let _lhs359 = globalState;
              _1547_v4 = _rhs370;
              _lhs357[_lhs358] = _rhs371;
              _lhs359.f15 = _rhs372;
              (globalState).f21 = new BigNumber((_this.f30).length);
              (globalState).f13 = (_this).f31;
            }
          }
        }
      }
      r0 = !((_this).f32);
      let _1849_v206;
      let _init50 = ((_1850_v4) => function (_1851_i29) {
        return _1850_v4;
      })(_1547_v4);
      let _nw330 = Array((new BigNumber(17)).toNumber());
      for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw330.length); _i0_50++) {
        _nw330[_i0_50] = _init50(new BigNumber(_i0_50));
      }
      _1849_v206 = _nw330;
      r1 = _1849_v206;
      r2 = p0;
      let _1852_v207;
      _1852_v207 = _dafny.Seq.of((_1545_v2).f34);
      let _1853_v208;
      _1853_v208 = _dafny.Map.Empty.slice().updateUnsafe(_1852_v207,_1547_v4);
      let _1854_v209;
      _1854_v209 = _dafny.Seq.of(new BigNumber((_1853_v208).length), new BigNumber((_1554_v11).cardinality()));
      let _1855_v210;
      _1855_v210 = _dafny.Map.Empty.slice().updateUnsafe(_1852_v207,(_1545_v2).f33);
      r3 = function () {
        let _coll60 = new _dafny.Set();
        for (const _compr_60 of (_1855_v210).Keys.Elements) {
          let _1856_v211 = _compr_60;
          if ((_1855_v210).contains(_1856_v211)) {
            _coll60.add(_1856_v211);
          }
        }
        return _coll60;
      }();
      return [r0, r1, r2, r3];
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    m0(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let _1857_v0;
      _1857_v0 = true;
      (globalState).f28 = _1857_v0;
      let _1858_v1;
      let _init51 = function (_1859_i1) {
        return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("r"), _dafny.Seq.UnicodeFromString("g"));
      };
      let _nw331 = Array((new BigNumber(24)).toNumber());
      for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw331.length); _i0_51++) {
        _nw331[_i0_51] = _init51(new BigNumber(_i0_51));
      }
      _1858_v1 = _nw331;
      for (const _guard_loop_11 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1858_v1).length))) {
        let _1860_i0 = _guard_loop_11;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1860_i0)) && ((_1860_i0).isLessThan(new BigNumber((_1858_v1).length))))) {
          (_1858_v1)[(_1860_i0)] = (p0).isLessThanOrEqualTo(new BigNumber(((_dafny.Set.fromElements(false, _1857_v0)).Union(_dafny.Set.fromElements(_1857_v0))).length));
        }
      }
      let _1861_v2;
      let _nw332 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
      _1861_v2 = _nw332;
      let _1862_v3;
      _1862_v3 = _dafny.MultiSet.fromElements(_1857_v0);
      let _1863_v4;
      _1863_v4 = _module.D0.create_DC2(_1862_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(617)), function (_1864_i2) {
  return new _dafny.CodePoint('c'.codePointAt(0));
}), _1857_v0);
      let _1865_v5;
      _1865_v5 = _dafny.Seq.of(_1857_v0);
      let _1866_v6;
      _1866_v6 = _dafny.Seq.UnicodeFromString("xnoahtmjj");
      let _1867_v7;
      _1867_v7 = _module.D0.create_DC0(_1857_v0);
      let _pat_let_tv37 = _1866_v6;
      let _pat_let_tv38 = _1866_v6;
      let _pat_let_tv39 = _1866_v6;
      let _pat_let_tv40 = _1865_v5;
      let _pat_let_tv41 = p0;
      let _pat_let_tv42 = _1865_v5;
      let _pat_let_tv43 = p3;
      let _pat_let_tv44 = _1857_v0;
      let _rhs373 = !(!(function (_source25) {
        if (_source25.is_DC0) {
          let _1868___mcc_h0 = (_source25).cf0;
          let _1869_cf0 = _1868___mcc_h0;
          return _1869_cf0;
        } else if (_source25.is_DC1) {
          let _1870___mcc_h1 = (_source25).cf1;
          let _1871___mcc_h2 = (_source25).cf2;
          let _1872___mcc_h3 = (_source25).cf3;
          let _1873_cf3 = _1872___mcc_h3;
          let _1874_cf2 = _1871___mcc_h2;
          let _1875_cf1 = _1870___mcc_h1;
          return _dafny.Seq.IsProperPrefixOf(_pat_let_tv37, _pat_let_tv38);
        } else {
          let _1876___mcc_h4 = (_source25).cf4;
          let _1877___mcc_h5 = (_source25).cf5;
          let _1878___mcc_h6 = (_source25).cf6;
          let _1879_cf6 = _1878___mcc_h6;
          let _1880_cf5 = _1877___mcc_h5;
          let _1881_cf4 = _1876___mcc_h4;
          return _1879_cf6;
        }
      }(function (_pat_let39_0) {
        return function (_1882_dt__update__tmp_h0) {
          return function (_pat_let40_0) {
            return function (_1883_dt__update_hcf5_h0) {
              return function (_pat_let41_0) {
                return function (_1884_dt__update_hcf6_h0) {
                  return _module.D0.create_DC2((_1882_dt__update__tmp_h0).dtor_cf4, _1883_dt__update_hcf5_h0, _1884_dt__update_hcf6_h0);
                }(_pat_let41_0);
              }((_pat_let_tv40)[_module.__default.safeIndex(_pat_let_tv41, new BigNumber((_pat_let_tv42).length))]);
            }(_pat_let40_0);
          }(_pat_let_tv39);
        }(_pat_let39_0);
      }(_1863_v4))));
      let _rhs374 = _1861_v2;
      let _rhs375 = function (_source26) {
        if (_source26.is_DC0) {
          let _1885___mcc_h7 = (_source26).cf0;
          let _1886_cf0 = _1885___mcc_h7;
          return _1886_cf0;
        } else if (_source26.is_DC1) {
          let _1887___mcc_h8 = (_source26).cf1;
          let _1888___mcc_h9 = (_source26).cf2;
          let _1889___mcc_h10 = (_source26).cf3;
          let _1890_cf3 = _1889___mcc_h10;
          let _1891_cf2 = _1888___mcc_h9;
          let _1892_cf1 = _1887___mcc_h8;
          return (_pat_let_tv43).isLessThanOrEqualTo(new BigNumber(27));
        } else {
          let _1893___mcc_h11 = (_source26).cf4;
          let _1894___mcc_h12 = (_source26).cf5;
          let _1895___mcc_h13 = (_source26).cf6;
          let _1896_cf6 = _1895___mcc_h13;
          let _1897_cf5 = _1894___mcc_h12;
          let _1898_cf4 = _1893___mcc_h11;
          return _pat_let_tv44;
        }
      }(_1867_v7);
      let _rhs376 = _dafny.Seq.Concat(_1866_v6, _1866_v6);
      let _rhs377 = _1857_v0;
      let _lhs360 = globalState;
      let _lhs361 = globalState;
      let _lhs362 = globalState;
      _1857_v0 = _rhs373;
      _1861_v2 = _rhs374;
      _lhs360.f13 = _rhs375;
      _lhs361.f11 = _rhs376;
      _lhs362.f13 = _rhs377;
      let _1899_i3;
      _1899_i3 = _dafny.ZERO;
      L9: {
        while (!(_1857_v0)) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1899_i3)) {
              break L9;
            }
            _1899_i3 = (_1899_i3).plus(_dafny.ONE);
            let _1900_v8;
            let _nw333 = Array((new BigNumber(20)).toNumber()).fill(_module.D0.Default());
            _1900_v8 = _nw333;
            let _pat_let_tv45 = p0;
            let _index316 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1900_v8).length));
            (_1900_v8)[_index316] = function (_pat_let42_0) {
              return function (_1901_dt__update__tmp_h1) {
                return function (_pat_let43_0) {
                  return function (_1902_dt__update_hcf3_h0) {
                    return _module.D0.create_DC1((_1901_dt__update__tmp_h1).dtor_cf1, (_1901_dt__update__tmp_h1).dtor_cf2, _1902_dt__update_hcf3_h0);
                  }(_pat_let43_0);
                }(_pat_let_tv45);
              }(_pat_let42_0);
            }(_module.D0.create_DC1(_1858_v1, p3, _module.__default.fm0(p3, _1857_v0, _1857_v0, globalState)));
            let _1903_v9;
            _1903_v9 = _module.D1.create_DC3(_1865_v5);
            let _1904_v10;
            _1904_v10 = _module.D0.create_DC1(_1858_v1, _module.__default.fm0(new BigNumber(((_1903_v9).dtor_cf7).length), _1857_v0, _1857_v0, globalState), p0);
            let _index317 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1900_v8).length));
            (_1900_v8)[_index317] = ((false) ? (_1904_v10) : (_module.D0.create_DC1(_1858_v1, p0, p3)));
            let _1905_v11;
            _1905_v11 = _dafny.Seq.of(p3, p3);
            let _1906_v12;
            _1906_v12 = _dafny.Map.Empty.slice().updateUnsafe((_1905_v11)[_module.__default.safeIndex(p0, new BigNumber((_1905_v11).length))],_dafny.Seq.UnicodeFromString("jr"));
            let _1907_v13;
            _1907_v13 = _dafny.Seq.of(p1);
            let _1908_v14;
            let _nw334 = Array((new BigNumber(21)).toNumber());
            _nw334[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_1907_v13, _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1907_v13).length)), p1);
            _nw334[(_dafny.ONE).toNumber()] = ((_1857_v0) ? (_1907_v13) : (_dafny.Seq.of(p1, p1)));
            _nw334[(new BigNumber(2)).toNumber()] = _1907_v13;
            _nw334[(new BigNumber(3)).toNumber()] = _1907_v13;
            _nw334[(new BigNumber(4)).toNumber()] = _1907_v13;
            _nw334[(new BigNumber(5)).toNumber()] = _1907_v13;
            _nw334[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_1907_v13, _module.__default.safeIndex(p0, new BigNumber((_1907_v13).length)), p1);
            _nw334[(new BigNumber(7)).toNumber()] = _1907_v13;
            _nw334[(new BigNumber(8)).toNumber()] = _1907_v13;
            _nw334[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(p1, p1, p1);
            _nw334[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1907_v13, _1907_v13);
            _nw334[(new BigNumber(11)).toNumber()] = _1907_v13;
            _nw334[(new BigNumber(12)).toNumber()] = _1907_v13;
            _nw334[(new BigNumber(13)).toNumber()] = _1907_v13;
            _nw334[(new BigNumber(14)).toNumber()] = _1907_v13;
            _nw334[(new BigNumber(15)).toNumber()] = _1907_v13;
            _nw334[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(p1, p1);
            _nw334[(new BigNumber(17)).toNumber()] = _1907_v13;
            _nw334[(new BigNumber(18)).toNumber()] = ((_1857_v0) ? (_1907_v13) : (_1907_v13));
            _nw334[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_1907_v13, _dafny.Seq.of(p1));
            _nw334[(new BigNumber(20)).toNumber()] = _1907_v13;
            _1908_v14 = _nw334;
            let _index318 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1908_v14).length));
            (_1908_v14)[_index318] = _dafny.Seq.update(_1907_v13, _module.__default.safeIndex(p3, new BigNumber((_1907_v13).length)), p1);
            let _1909_v15;
            _1909_v15 = new _dafny.CodePoint('e'.codePointAt(0));
            let _1910_v16;
            _1910_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of(p3));
            let _index319 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1908_v14).length));
            let _rhs378 = ((_1857_v0) ? (_module.__default.safeDivisionInt(p0, (_dafny.ZERO).minus(p3))) : (p0));
            let _rhs379 = _1906_v12;
            let _rhs380 = _dafny.Seq.Concat(_1907_v13, _1907_v13);
            let _rhs381 = p3;
            let _rhs382 = ((_module.__default.fm1(_1909_v15, new BigNumber(((((_1910_v16).contains(p3)) ? ((_1910_v16).get(p3)) : (_1905_v11))).length), (_dafny.ZERO).minus(p3), _1857_v0, globalState)) ? (!(_1857_v0) || (_1857_v0)) : ((p3).isLessThan(_module.__default.fm0(p0, _1857_v0, _1857_v0, globalState))));
            let _lhs363 = globalState;
            let _lhs364 = _1908_v14;
            let _lhs365 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1908_v14).length));
            let _lhs366 = globalState;
            let _lhs367 = globalState;
            _lhs363.f1 = _rhs378;
            _1906_v12 = _rhs379;
            _lhs364[_lhs365] = _rhs380;
            _lhs366.f21 = _rhs381;
            _lhs367.f13 = _rhs382;
            let _1911_v17;
            let _nw335 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _1911_v17 = _nw335;
            let _index320 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1911_v17).length));
            (_1911_v17)[_index320] = _1866_v6;
            let _index321 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1911_v17).length));
            (_1911_v17)[_index321] = (((_1906_v12).contains((_module.__default.fm0((_dafny.ZERO).minus(p3), _1857_v0, _1857_v0, globalState)).plus(_module.__default.fm0(new BigNumber(74), _1857_v0, _1857_v0, globalState)))) ? ((_1906_v12).get((_module.__default.fm0((_dafny.ZERO).minus(p3), _1857_v0, _1857_v0, globalState)).plus(_module.__default.fm0(new BigNumber(74), _1857_v0, _1857_v0, globalState)))) : (_1866_v6));
            let _1912_v18;
            _1912_v18 = _dafny.Set.fromElements(_1857_v0);
            let _rhs383 = p1;
            let _rhs384 = _1905_v11;
            let _rhs385 = new BigNumber(((_1911_v17)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1911_v17).length))]).length);
            let _rhs386 = _dafny.Set.fromElements(_dafny.areEqual(_1905_v11, _dafny.Seq.of(new BigNumber(-408))));
            let _lhs368 = globalState;
            let _lhs369 = globalState;
            let _lhs370 = globalState;
            _lhs368.f17 = _rhs383;
            _lhs369.f7 = _rhs384;
            _lhs370.f21 = _rhs385;
            _1912_v18 = _rhs386;
          }
        }
      }
      let _source27 = _1863_v4;
      if (_source27.is_DC0) {
        let _1913___mcc_h14 = (_source27).cf0;
        let _1914_cf0 = _1913___mcc_h14;
        let _1915_v19;
        let _nw336 = new _module.C0();
        _nw336.__ctor((_1862_v3).IsDisjointFrom(_dafny.MultiSet.FromArray(_1865_v5)));
        _1915_v19 = _nw336;
        let _1916_v20;
        let _init52 = ((_1917_v19, _1918_p3) => function (_1919_i4) {
          return _module.D6.create_DC23(_1917_v19.f37, (_dafny.ZERO).minus(_1918_p3));
        })(_1915_v19, p3);
        let _nw337 = Array((new BigNumber(19)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw337.length); _i0_52++) {
          _nw337[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _1916_v20 = _nw337;
        _1916_v20 = _1916_v20;
        (globalState).f20 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(276)), function (_1920_i5) {
          return ((false) ? (new _dafny.CodePoint('s'.codePointAt(0))) : (new _dafny.CodePoint('g'.codePointAt(0))));
        });
        let _1921_v21;
        _1921_v21 = new _dafny.CodePoint('r'.codePointAt(0));
        _1921_v21 = _1921_v21;
      } else if (_source27.is_DC1) {
        let _1922___mcc_h15 = (_source27).cf1;
        let _1923___mcc_h16 = (_source27).cf2;
        let _1924___mcc_h17 = (_source27).cf3;
        let _1925_cf3 = _1924___mcc_h17;
        let _1926_cf2 = _1923___mcc_h16;
        let _1927_cf1 = _1922___mcc_h15;
        let _1928_v22;
        _1928_v22 = _module.D17.create_DC51(p0, p1);
        if ((_module.__default.safeModuloInt((_dafny.ZERO).minus(p3), (_1928_v22).dtor_cf86)).isLessThanOrEqualTo(p3)) {
          let _1929_v23;
          _1929_v23 = new _dafny.CodePoint('f'.codePointAt(0));
          _1929_v23 = _1929_v23;
          let _index322 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_1927_cf1).length));
          (_1927_cf1)[_index322] = _1857_v0;
          let _index323 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_1927_cf1).length));
          (_1927_cf1)[_index323] = _dafny.Seq.contains(_1866_v6, _1929_v23);
          let _1930_v24;
          let _nw338 = new _module.C6();
          _nw338.__ctor((_1927_cf1)[_module.__default.safeIndex(new BigNumber(39), new BigNumber((_1927_cf1).length))], _1866_v6);
          _1930_v24 = _nw338;
          let _1931_v25;
          _1931_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1857_v0,new BigNumber(181));
          _1931_v25 = (_1931_v25).update(_1857_v0, _1925_cf3);
          (globalState).f13 = _1857_v0;
        } else {
          (globalState).f13 = _1857_v0;
          let _1932_v26;
          let _nw339 = new _module.C3();
          _nw339.__ctor(_1857_v0, _1925_cf3, _dafny.areEqual(_1866_v6, _1866_v6), _1866_v6);
          _1932_v26 = _nw339;
          (globalState).f21 = new BigNumber((_dafny.Seq.Concat(_1866_v6, _1866_v6)).length);
          (globalState).f18 = _module.__default.fm1(_module.__default.fm42(_1857_v0, (_1932_v26).f35, globalState), (_dafny.ZERO).minus(p3), new BigNumber(753), (_1857_v0) || (_1857_v0), globalState);
          let _index324 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_1858_v1).length));
          (_1858_v1)[_index324] = !(_dafny.MultiSet.FromArray(_1865_v5)).contains(false);
          let _index325 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_1858_v1).length));
          (_1858_v1)[_index325] = _1857_v0;
        }
        let _1933_v27;
        _1933_v27 = _dafny.Seq.of(p3, p0, new BigNumber(743), p3);
        let _1934_v28;
        _1934_v28 = _dafny.Set.fromElements(_1933_v27);
        let _1935_v30;
        _1935_v30 = _dafny.Seq.of(_1933_v27, _1933_v27);
        let _1936_v32;
        _1936_v32 = new _dafny.CodePoint('n'.codePointAt(0));
        let _1937_v33;
        _1937_v33 = _dafny.Seq.of((_1935_v30)[_module.__default.safeIndex(new BigNumber((function () {
          let _coll61 = new _dafny.Map();
          for (const _compr_61 of (_dafny.Map.Empty.slice().updateUnsafe(_1867_v7,_1936_v32)).Keys.Elements) {
            let _1938_v31 = _compr_61;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_1867_v7,_1936_v32)).contains(_1938_v31)) {
              _coll61.push([_1938_v31,true]);
            }
          }
          return _coll61;
        }()).length), new BigNumber((_1935_v30).length))]);
        let _1939_v35;
        _1939_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1857_v0,_1857_v0);
        let _1940_v36;
        _1940_v36 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_1865_v5).length), _1926_cf2),new BigNumber(((_1939_v35).update(!(_1857_v0), _1857_v0)).length));
        let _1941_v42;
        let _nw340 = Array((new BigNumber(24)).toNumber());
        _nw340[(_dafny.ZERO).toNumber()] = (_1934_v28).Difference(_1934_v28);
        _nw340[(_dafny.ONE).toNumber()] = _1934_v28;
        _nw340[(new BigNumber(2)).toNumber()] = _1934_v28;
        _nw340[(new BigNumber(3)).toNumber()] = _1934_v28;
        _nw340[(new BigNumber(4)).toNumber()] = (_module.__default.fm25(false, _1857_v0, _1866_v6, globalState)).Union(function () {
          let _coll62 = new _dafny.Set();
          for (const _compr_62 of (_dafny.MultiSet.fromElements(_1933_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(833)), ((_1942_p3) => function (_1943_i6) {
            return _1942_p3;
          })(p3)))).Elements) {
            let _1944_v29 = _compr_62;
            if ((_dafny.MultiSet.fromElements(_1933_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(833)), ((_1945_p3) => function (_1943_i6) {
              return _1945_p3;
            })(p3)))).contains(_1944_v29)) {
              _coll62.add(_1944_v29);
            }
          }
          return _coll62;
        }());
        _nw340[(new BigNumber(5)).toNumber()] = _module.__default.fm25(_1857_v0, _1857_v0, _dafny.Seq.UnicodeFromString("t"), globalState);
        _nw340[(new BigNumber(6)).toNumber()] = function () {
          let _coll63 = new _dafny.Set();
          for (const _compr_63 of (_1935_v30).Elements) {
            let _1946_v34 = _compr_63;
            if (_dafny.Seq.contains(_1935_v30, _1946_v34)) {
              _coll63.add(_1946_v34);
            }
          }
          return _coll63;
        }();
        _nw340[(new BigNumber(7)).toNumber()] = _1934_v28;
        _nw340[(new BigNumber(8)).toNumber()] = _1934_v28;
        _nw340[(new BigNumber(9)).toNumber()] = (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_1939_v35).length)), _1933_v27)).Intersect(function () {
          let _coll64 = new _dafny.Set();
          for (const _compr_64 of (_1940_v36).Keys.Elements) {
            let _1947_v37 = _compr_64;
            if ((_1940_v36).contains(_1947_v37)) {
              _coll64.add(_1947_v37);
            }
          }
          return _coll64;
        }());
        _nw340[(new BigNumber(10)).toNumber()] = function () {
          let _coll65 = new _dafny.Set();
          for (const _compr_65 of (_1940_v36).Keys.Elements) {
            let _1948_v38 = _compr_65;
            if ((_1940_v36).contains(_1948_v38)) {
              _coll65.add(_1948_v38);
            }
          }
          return _coll65;
        }();
        _nw340[(new BigNumber(11)).toNumber()] = _1934_v28;
        _nw340[(new BigNumber(12)).toNumber()] = (_1934_v28).Intersect(_1934_v28);
        _nw340[(new BigNumber(13)).toNumber()] = (_1934_v28).Difference(_dafny.Set.fromElements(_1933_v27));
        _nw340[(new BigNumber(14)).toNumber()] = (function () {
          let _coll66 = new _dafny.Set();
          for (const _compr_66 of (_1937_v33).Elements) {
            let _1949_v39 = _compr_66;
            if (_dafny.Seq.contains(_1937_v33, _1949_v39)) {
              _coll66.add(_1949_v39);
            }
          }
          return _coll66;
        }()).Difference(_1934_v28);
        _nw340[(new BigNumber(15)).toNumber()] = _1934_v28;
        _nw340[(new BigNumber(16)).toNumber()] = _1934_v28;
        _nw340[(new BigNumber(17)).toNumber()] = (_1934_v28).Intersect(_1934_v28);
        _nw340[(new BigNumber(18)).toNumber()] = _1934_v28;
        _nw340[(new BigNumber(19)).toNumber()] = (_1934_v28).Intersect(_1934_v28);
        _nw340[(new BigNumber(20)).toNumber()] = function () {
          let _coll67 = new _dafny.Set();
          for (const _compr_67 of (function () {
            let _coll68 = new _dafny.Set();
            for (const _compr_68 of (_1935_v30).Elements) {
              let _1950_v40 = _compr_68;
              if (_dafny.Seq.contains(_1935_v30, _1950_v40)) {
                _coll68.add(_1950_v40);
              }
            }
            return _coll68;
          }()).Elements) {
            let _1951_v41 = _compr_67;
            if ((function () {
              let _coll69 = new _dafny.Set();
              for (const _compr_69 of (_1935_v30).Elements) {
                let _1952_v40 = _compr_69;
                if (_dafny.Seq.contains(_1935_v30, _1952_v40)) {
                  _coll69.add(_1952_v40);
                }
              }
              return _coll69;
            }()).contains(_1951_v41)) {
              _coll67.add(_1951_v41);
            }
          }
          return _coll67;
        }();
        _nw340[(new BigNumber(21)).toNumber()] = _dafny.Set.fromElements(_1933_v27, _1933_v27);
        _nw340[(new BigNumber(22)).toNumber()] = _1934_v28;
        _nw340[(new BigNumber(23)).toNumber()] = _1934_v28;
        _1941_v42 = _nw340;
        let _index326 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_1941_v42).length));
        (_1941_v42)[_index326] = _1934_v28;
        let _index327 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_1941_v42).length));
        (_1941_v42)[_index327] = _1934_v28;
        let _1953_v43;
        _1953_v43 = _module.D9.create_DC28(_1865_v5, _1857_v0, (_1933_v27)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_1926_cf2)).length), new BigNumber((_1933_v27).length))], true, _1857_v0);
        (globalState).f19 = (_1953_v43).dtor_cf43;
        (globalState).f1 = p0;
      } else {
        let _1954___mcc_h18 = (_source27).cf4;
        let _1955___mcc_h19 = (_source27).cf5;
        let _1956___mcc_h20 = (_source27).cf6;
        let _1957_cf6 = _1956___mcc_h20;
        let _1958_cf5 = _1955___mcc_h19;
        let _1959_cf4 = _1954___mcc_h18;
        let _1960_v44;
        _1960_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_1857_v0, true)).cardinality()),_1957_cf6);
        let _1961_v45;
        _1961_v45 = _dafny.Seq.of(p1);
        let _1962_v46;
        _1962_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1960_v44,(_1961_v45)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_1961_v45).length))]);
        let _1963_v47;
        _1963_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1957_cf6,new BigNumber((_dafny.Seq.UnicodeFromString("l")).length));
        let _1964_v48;
        _1964_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1963_v47).length),p1);
        let _1965_v49;
        let _nw341 = Array((new BigNumber(19)).toNumber());
        _nw341[(_dafny.ZERO).toNumber()] = ((_1857_v0) ? (p1) : (p1));
        _nw341[(_dafny.ONE).toNumber()] = p1;
        _nw341[(new BigNumber(2)).toNumber()] = p1;
        _nw341[(new BigNumber(3)).toNumber()] = p1;
        _nw341[(new BigNumber(4)).toNumber()] = p1;
        _nw341[(new BigNumber(5)).toNumber()] = (((_1962_v46).contains(_1960_v44)) ? ((_1962_v46).get(_1960_v44)) : (p1));
        _nw341[(new BigNumber(6)).toNumber()] = p1;
        _nw341[(new BigNumber(7)).toNumber()] = p1;
        _nw341[(new BigNumber(8)).toNumber()] = p1;
        _nw341[(new BigNumber(9)).toNumber()] = (((_1964_v48).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-917)), function (_1967_i7) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        })).length))) ? ((_1964_v48).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-917)), function (_1966_i7) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        })).length))) : (p1));
        _nw341[(new BigNumber(10)).toNumber()] = p1;
        _nw341[(new BigNumber(11)).toNumber()] = p1;
        _nw341[(new BigNumber(12)).toNumber()] = (((_1964_v48).contains(_module.__default.fm0(p3, (_1865_v5)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1960_v44).length)), new BigNumber((_1865_v5).length))], _1957_cf6, globalState))) ? ((_1964_v48).get(_module.__default.fm0(p3, (_1865_v5)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1960_v44).length)), new BigNumber((_1865_v5).length))], _1957_cf6, globalState))) : (p1));
        _nw341[(new BigNumber(13)).toNumber()] = p1;
        _nw341[(new BigNumber(14)).toNumber()] = p1;
        _nw341[(new BigNumber(15)).toNumber()] = p1;
        _nw341[(new BigNumber(16)).toNumber()] = p1;
        _nw341[(new BigNumber(17)).toNumber()] = ((false) ? (p1) : (p1));
        _nw341[(new BigNumber(18)).toNumber()] = p1;
        _1965_v49 = _nw341;
        let _index328 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1965_v49).length));
        (_1965_v49)[_index328] = ((_1857_v0) ? (p1) : (p1));
        let _index329 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1965_v49).length));
        (_1965_v49)[_index329] = p1;
        let _1968_v50;
        _1968_v50 = new _dafny.CodePoint('u'.codePointAt(0));
        let _1969_v51;
        _1969_v51 = _module.D2.create_DC8(_1968_v50, _1957_cf6);
        let _1970_v52;
        _1970_v52 = _dafny.Seq.of(_1969_v51, _1969_v51);
        _1865_v5 = _dafny.Seq.Concat(_module.__default.fm35(_1970_v52, p0, globalState), _1865_v5);
        let _1971_v53;
        let _nw342 = new _module.C6();
        _nw342.__ctor(_1857_v0, _1866_v6);
        _1971_v53 = _nw342;
        (globalState).f11 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1958_cf5, _1958_cf5), _1866_v6);
      }
      if (_1857_v0) {
        let _pat_let_tv46 = _1866_v6;
        let _1972_v54;
        _1972_v54 = _dafny.Map.Empty.slice().updateUnsafe(p1,function (_pat_let44_0) {
          return function (_1973_dt__update__tmp_h2) {
            return function (_pat_let45_0) {
              return function (_1974_dt__update_hcf5_h1) {
                return _module.D0.create_DC2((_1973_dt__update__tmp_h2).dtor_cf4, _1974_dt__update_hcf5_h1, (_1973_dt__update__tmp_h2).dtor_cf6);
              }(_pat_let45_0);
            }(_pat_let_tv46);
          }(_pat_let44_0);
        }(_1863_v4));
        _1972_v54 = (_1972_v54).update(p1, _1863_v4);
        (globalState).f15 = !((!(((_1857_v0) ? (_module.__default.fm10(p0, p0, globalState)) : (!(false))))) === (_1857_v0));
        (globalState).f13 = _1857_v0;
        (globalState).f1 = ((new BigNumber(483)).plus(p3)).multipliedBy(p0);
        let _1975_v55;
        _1975_v55 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1976_v56;
        _1976_v56 = _module.D2.create_DC8(_1975_v55, _1857_v0);
        let _1977_v57;
        _1977_v57 = _dafny.Seq.of(_1976_v56, _module.D2.create_DC8(_1975_v55, _1857_v0));
        let _1978_v58;
        _1978_v58 = _module.D18.create_DC54(_1977_v57);
        let _1979_v59;
        _1979_v59 = _module.D1.create_DC3(_module.__default.fm35((_1978_v58).dtor_cf93, p0, globalState));
        let _pat_let_tv47 = _1865_v5;
        let _source28 = function (_pat_let46_0) {
          return function (_1980_dt__update__tmp_h3) {
            return function (_pat_let47_0) {
              return function (_1981_dt__update_hcf7_h0) {
                return _module.D1.create_DC3(_1981_dt__update_hcf7_h0);
              }(_pat_let47_0);
            }(_pat_let_tv47);
          }(_pat_let46_0);
        }(_1979_v59);
        if (_source28.is_DC4) {
          let _1982___mcc_h21 = (_source28).cf8;
          let _1983_cf8 = _1982___mcc_h21;
          let _index330 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((p1).length));
          (p1)[_index330] = (p3).plus(p0);
          let _1984_v60;
          _1984_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1858_v1,p3);
          let _index331 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((p1).length));
          (p1)[_index331] = (((_1984_v60).contains(_1858_v1)) ? ((_1984_v60).get(_1858_v1)) : ((p3).minus(p0)));
          let _1985_v61;
          _1985_v61 = _dafny.MultiSet.fromElements(p0, p3);
          let _1986_v62;
          let _nw343 = new _module.C7();
          _nw343.__ctor(_1983_cf8, !(_1983_cf8) || (_module.__default.fm10(new BigNumber((_1866_v6).length), (p1)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((p1).length))], globalState)), (_1985_v61).IsSubsetOf(_1985_v61), _1866_v6);
          _1986_v62 = _nw343;
          let _1987_v63;
          let _nw344 = new _module.C1();
          _nw344.__ctor(true);
          _1987_v63 = _nw344;
          let _1988_v64;
          _1988_v64 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(349)), ((_1989_v55) => function (_1990_i8) {
            return _1989_v55;
          })(_1975_v55)));
          let _1991_v65;
          let _init53 = ((_1992_v6) => function (_1993_i9) {
            return _module.__default.safeDivisionInt(_1993_i9, new BigNumber((_1992_v6).length));
          })(_1866_v6);
          let _nw345 = Array((new BigNumber(15)).toNumber());
          for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw345.length); _i0_53++) {
            _nw345[_i0_53] = _init53(new BigNumber(_i0_53));
          }
          _1991_v65 = _nw345;
          let _1994_v66;
          _1994_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1991_v65,_module.__default.fm6((p1)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((p1).length))], globalState));
          _1983_cf8 = ((new BigNumber((_1988_v64).length)).minus(new BigNumber((_1994_v66).length))).isEqualTo(p0);
        } else if (_source28.is_DC3) {
          let _1995___mcc_h22 = (_source28).cf7;
          let _1996_cf7 = _1995___mcc_h22;
          let _index332 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_1858_v1).length));
          (_1858_v1)[_index332] = !(_1857_v0) || (_1857_v0);
          let _1997_v67;
          _1997_v67 = _dafny.MultiSet.fromElements(_1996_cf7);
          let _index333 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_1858_v1).length));
          (_1858_v1)[_index333] = (_1997_v67).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1865_v5));
          let _1998_v68;
          _1998_v68 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_1975_v55, (_dafny.ZERO).minus(p0), p0, false, globalState),_1979_v59);
          let _1999_v69;
          _1999_v69 = _module.D5.create_DC20(_module.D5.create_DC16(_1998_v68));
          let _2000_v70;
          _2000_v70 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1999_v69);
          _2000_v70 = _2000_v70;
          (globalState).f11 = _1866_v6;
          let _2001_v71;
          _2001_v71 = _dafny.Set.fromElements(_1866_v6);
          let _2002_v72;
          _2002_v72 = _module.D5.create_DC17(_2001_v71);
          _2002_v72 = _2002_v72;
        } else {
          let _2003___mcc_h23 = (_source28).cf9;
          let _2004_cf9 = _2003___mcc_h23;
          (globalState).f28 = !((_1863_v4).dtor_cf6);
          let _index334 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((p1).length));
          (p1)[_index334] = p0;
          let _index335 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((p1).length));
          (p1)[_index335] = new BigNumber(360);
          let _2005_v73;
          _2005_v73 = _dafny.MultiSet.fromElements((p1)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((p1).length))]);
          let _2006_v74;
          let _nw346 = Array((new BigNumber(20)).toNumber());
          _nw346[(_dafny.ZERO).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((p1).length))];
          _nw346[(_dafny.ONE).toNumber()] = p0;
          _nw346[(new BigNumber(2)).toNumber()] = p3;
          _nw346[(new BigNumber(3)).toNumber()] = p3;
          _nw346[(new BigNumber(4)).toNumber()] = (p3).minus(new BigNumber((_2005_v73).cardinality()));
          _nw346[(new BigNumber(5)).toNumber()] = p0;
          _nw346[(new BigNumber(6)).toNumber()] = p0;
          _nw346[(new BigNumber(7)).toNumber()] = p0;
          _nw346[(new BigNumber(8)).toNumber()] = new BigNumber(-609);
          _nw346[(new BigNumber(9)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((p1).length))];
          _nw346[(new BigNumber(10)).toNumber()] = _module.__default.fm0(p0, _1857_v0, _1857_v0, globalState);
          _nw346[(new BigNumber(11)).toNumber()] = (p3).plus(p3);
          _nw346[(new BigNumber(12)).toNumber()] = p3;
          _nw346[(new BigNumber(13)).toNumber()] = new BigNumber((_1866_v6).length);
          _nw346[(new BigNumber(14)).toNumber()] = new BigNumber(-837);
          _nw346[(new BigNumber(15)).toNumber()] = p3;
          _nw346[(new BigNumber(16)).toNumber()] = p3;
          _nw346[(new BigNumber(17)).toNumber()] = _module.__default.fm9(globalState);
          _nw346[(new BigNumber(18)).toNumber()] = new BigNumber((_2005_v73).cardinality());
          _nw346[(new BigNumber(19)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((p1).length))];
          _2006_v74 = _nw346;
          let _2007_v75;
          _2007_v75 = _dafny.Seq.of(_1866_v6);
          let _index336 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((p1).length));
          let _rhs387 = _1975_v55;
          let _rhs388 = _2006_v74;
          let _rhs389 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((p1)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((p1).length))]), new BigNumber((_dafny.Seq.Concat(_1866_v6, _1866_v6)).length));
          let _rhs390 = _dafny.Seq.Concat((_2007_v75)[_module.__default.safeIndex(new BigNumber(367), new BigNumber((_2007_v75).length))], ((_1857_v0) ? (_1866_v6) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(950)), ((_2008_v55) => function (_2009_i10) {
            return _2008_v55;
          })(_1975_v55)))));
          let _lhs371 = globalState;
          let _lhs372 = p1;
          let _lhs373 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((p1).length));
          let _lhs374 = globalState;
          _1975_v55 = _rhs387;
          _lhs371.f17 = _rhs388;
          _lhs372[_lhs373] = _rhs389;
          _lhs374.f11 = _rhs390;
          let _2010_v76;
          let _nw347 = new _module.C7();
          _nw347.__ctor(_1857_v0, _1857_v0, _1857_v0, _dafny.Seq.update(_dafny.Seq.Concat(_1866_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(998)), ((_2011_v55) => function (_2012_i11) {
            return _2011_v55;
          })(_1975_v55))), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_1866_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(998)), ((_2013_v55) => function (_2014_i11) {
            return _2013_v55;
          })(_1975_v55)))).length)), _1975_v55));
          _2010_v76 = _nw347;
        }
      } else {
        let _2015_v77;
        let _nw348 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2015_v77 = _nw348;
        let _2016_v78;
        _2016_v78 = _dafny.Map.Empty.slice().updateUnsafe(_2015_v77,_1857_v0);
        let _2017_v79;
        _2017_v79 = _dafny.Seq.of(_2015_v77);
        if (!((((_2016_v78).contains((_2017_v79)[_module.__default.safeIndex(p0, new BigNumber((_2017_v79).length))])) ? ((_2016_v78).get((_2017_v79)[_module.__default.safeIndex(p0, new BigNumber((_2017_v79).length))])) : (_1857_v0)))) {
          (globalState).f13 = !(_1857_v0);
          let _2018_v80;
          _2018_v80 = _dafny.Seq.of(p0, p0);
          (globalState).f19 = new BigNumber((_dafny.Seq.update(_2018_v80, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_1857_v0, _1857_v0)).cardinality())), new BigNumber((_2018_v80).length)), p0)).length);
          let _2019_v81;
          _2019_v81 = _dafny.Map.Empty.slice().updateUnsafe(p1,p3);
          _2019_v81 = ((_2019_v81).Merge(_2019_v81)).Merge(_2019_v81);
          let _index337 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_2015_v77).length));
          (_2015_v77)[_index337] = _dafny.Seq.UnicodeFromString("rpkac");
          let _index338 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_2015_v77).length));
          (_2015_v77)[_index338] = _1866_v6;
          let _2020_v82;
          _2020_v82 = _dafny.MultiSet.fromElements(_2018_v80);
          let _rhs391 = new BigNumber(((_2020_v82).update(_2018_v80, _module.__default.abs(p0))).cardinality());
          let _rhs392 = !_dafny.areEqual(_dafny.Seq.Concat(_1866_v6, _1866_v6), _dafny.Seq.Concat(_1866_v6, (_2015_v77)[_module.__default.safeIndex(new BigNumber(588), new BigNumber((_2015_v77).length))]));
          let _rhs393 = !(_1857_v0);
          let _lhs375 = globalState;
          let _lhs376 = globalState;
          let _lhs377 = globalState;
          _lhs375.f1 = _rhs391;
          _lhs376.f15 = _rhs392;
          _lhs377.f13 = _rhs393;
        } else {
          let _index339 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((p1).length));
          (p1)[_index339] = p0;
          let _2021_v83;
          _2021_v83 = _dafny.MultiSet.fromElements(p3, p0);
          let _2022_v84;
          _2022_v84 = _dafny.MultiSet.fromElements((_2021_v83).Difference(_2021_v83));
          let _index340 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((p1).length));
          (p1)[_index340] = new BigNumber((_2022_v84).cardinality());
          let _2023_v85;
          let _nw349 = Array((new BigNumber(25)).toNumber()).fill([]);
          _2023_v85 = _nw349;
          let _index341 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_2023_v85).length));
          (_2023_v85)[_index341] = p2;
          let _index342 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_2023_v85).length));
          (_2023_v85)[_index342] = p2;
          let _2024_v86;
          let _nw350 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
          _2024_v86 = _nw350;
          _2024_v86 = _2024_v86;
          _1858_v1 = _1858_v1;
          (globalState).f13 = !(_module.__default.fm19((_1866_v6)[_module.__default.safeIndex((p1)[_module.__default.safeIndex(new BigNumber(65), new BigNumber((p1).length))], new BigNumber((_1866_v6).length))], globalState));
        }
        let _2025_v87;
        _2025_v87 = new _dafny.CodePoint('d'.codePointAt(0));
        _1857_v0 = _module.__default.fm19(_2025_v87, globalState);
        (globalState).f13 = _1857_v0;
        let _2026_v88;
        _2026_v88 = _dafny.Set.fromElements(p1, p1, p1);
        let _2027_v89;
        let _nw351 = new _module.C2();
        _nw351.__ctor(new BigNumber((_dafny.Seq.Concat(_1865_v5, _1865_v5)).length), _2026_v88, false, _dafny.Seq.UnicodeFromString("wxlbbb"));
        _2027_v89 = _nw351;
        let _2028_v90;
        let _nw352 = Array((new BigNumber(6)).toNumber()).fill([]);
        _2028_v90 = _nw352;
        let _2029_v91;
        let _nw353 = Array((_dafny.ONE).toNumber());
        _2029_v91 = _nw353;
        let _index343 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_2028_v90).length));
        (_2028_v90)[_index343] = _2029_v91;
        let _index344 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_2028_v90).length));
        (_2028_v90)[_index344] = _2029_v91;
      }
      r0 = !(_1857_v0);
      return r0;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2030_v0;
      _2030_v0 = new BigNumber(553);
      let _2031_v1;
      _2031_v1 = _dafny.MultiSet.fromElements(_2030_v0);
      _2031_v1 = _2031_v1;
      let _2032_v2;
      let _init54 = function (_2033_i1) {
        return (_module.D0.create_DC0(false)).dtor_cf0;
      };
      let _nw354 = Array((new BigNumber(24)).toNumber());
      for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw354.length); _i0_54++) {
        _nw354[_i0_54] = _init54(new BigNumber(_i0_54));
      }
      _2032_v2 = _nw354;
      for (const _guard_loop_12 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2032_v2).length))) {
        let _2034_i0 = _guard_loop_12;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2034_i0)) && ((_2034_i0).isLessThan(new BigNumber((_2032_v2).length))))) {
          (_2032_v2)[(_2034_i0)] = ((_dafny.Set.fromElements(new BigNumber(852), _2030_v0)).Intersect(_dafny.Set.fromElements(_2030_v0))).IsDisjointFrom(function () {
            let _coll70 = new _dafny.Set();
            for (const _compr_70 of _dafny.IntegerRange(new BigNumber(833), new BigNumber(56))) {
              let _2035_v3 = _compr_70;
              if (((new BigNumber(833)).isLessThanOrEqualTo(_2035_v3)) && ((_2035_v3).isLessThan(new BigNumber(56)))) {
                _coll70.add((_2035_v3).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality())));
              }
            }
            return _coll70;
          }());
        }
      }
      let _hi11 = new BigNumber(518);
      for (let _2036_i2 = _module.__default.safeModuloInt(new BigNumber(969), _2030_v0); _2036_i2.isLessThan(_hi11); _2036_i2 = _2036_i2.plus(_dafny.ONE)) {
        (globalState).f1 = _2030_v0;
        (globalState).f14 = _2030_v0;
        let _2037_v4;
        _2037_v4 = _dafny.Seq.of(p0);
        _2037_v4 = _2037_v4;
        (globalState).f20 = _dafny.Seq.UnicodeFromString("iepcaxj");
      }
      let _2038_v5;
      _2038_v5 = false;
      let _2039_v6;
      _2039_v6 = _dafny.Seq.UnicodeFromString("d");
      let _2040_v7;
      let _nw355 = new _module.C4();
      _nw355.__ctor(!(_2038_v5), new BigNumber((_2039_v6).length));
      _2040_v7 = _nw355;
      let _2041_v8;
      _2041_v8 = _dafny.MultiSet.fromElements(_2040_v7, _2040_v7);
      _2041_v8 = (_2041_v8).Difference(_dafny.MultiSet.fromElements(_2040_v7));
      if (false) {
        let _2042_v9;
        _2042_v9 = _dafny.Seq.of(_2031_v1);
        let _2043_v10;
        _2043_v10 = _dafny.Seq.of(_dafny.Seq.update(((true) ? (_dafny.Seq.of(_2031_v1, _2031_v1)) : (_2042_v9)), _module.__default.safeIndex(_2030_v0, new BigNumber((((true) ? (_dafny.Seq.of(_2031_v1, _2031_v1)) : (_2042_v9))).length)), _2031_v1));
        _2042_v9 = (_2043_v10)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_2043_v10).length))];
        let _2044_v11;
        _2044_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2038_v5,(_dafny.ONE).isLessThanOrEqualTo((_2040_v7).f34));
        (globalState).f28 = (((_2044_v11).contains((_2040_v7).f33)) ? ((_2044_v11).get((_2040_v7).f33)) : (_2038_v5));
        let _2045_v12;
        let _nw356 = Array((_dafny.ONE).toNumber());
        _nw356[(_dafny.ZERO).toNumber()] = _2031_v1;
        _2045_v12 = _nw356;
        _2045_v12 = _2045_v12;
        let _nw357 = Array((new BigNumber(10)).toNumber()).fill(false);
        _2032_v2 = _nw357;
        if ((((_2040_v7).f33) ? ((_2040_v7).f33) : ((_2040_v7).f33))) {
          r0 = (new BigNumber((_2039_v6).length)).plus((_2040_v7).f34);
          let _2046_v13;
          _2046_v13 = _module.D2.create_DC7((_2040_v7).f33);
          let _2047_v14;
          _2047_v14 = _dafny.Map.Empty.slice().updateUnsafe(_2046_v13,(_2040_v7).f33);
          _2047_v14 = (_2047_v14).update(_2046_v13, true);
          let _2048_v15;
          let _nw358 = Array((new BigNumber(13)).toNumber()).fill([]);
          _2048_v15 = _nw358;
          let _2049_v16;
          _2049_v16 = new _dafny.CodePoint('h'.codePointAt(0));
          let _2050_v17;
          let _nw359 = Array((new BigNumber(20)).toNumber());
          _nw359[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_2049_v16);
          _nw359[(_dafny.ONE).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_2049_v16);
          _nw359[(new BigNumber(3)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-964)), ((_2051_v16) => function (_2052_i3) {
            return _2051_v16;
          })(_2049_v16)), _module.__default.safeIndex((_2040_v7).f34, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-964)), ((_2053_v16) => function (_2054_i3) {
            return _2053_v16;
          })(_2049_v16))).length)), _2049_v16);
          _nw359[(new BigNumber(5)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(6)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(7)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(8)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(9)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(10)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(11)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(12)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(new _dafny.CodePoint('i'.codePointAt(0)));
          _nw359[(new BigNumber(14)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(15)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(16)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(17)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(18)).toNumber()] = _2039_v6;
          _nw359[(new BigNumber(19)).toNumber()] = _2039_v6;
          _2050_v17 = _nw359;
          let _index345 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_2048_v15).length));
          (_2048_v15)[_index345] = _2050_v17;
          let _2055_v18;
          let _init55 = ((_2056_v0) => function (_2057_i4) {
            return (_2057_i4).multipliedBy(_2056_v0);
          })(_2030_v0);
          let _nw360 = Array((new BigNumber(9)).toNumber());
          for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw360.length); _i0_55++) {
            _nw360[_i0_55] = _init55(new BigNumber(_i0_55));
          }
          _2055_v18 = _nw360;
          let _index346 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_2055_v18).length));
          (_2055_v18)[_index346] = (_dafny.ZERO).minus((p0)[_module.__default.safeIndex(_2030_v0, new BigNumber((p0).length))]);
          let _index347 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_2048_v15).length));
          let _index348 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_2055_v18).length));
          let _rhs394 = _2050_v17;
          let _rhs395 = (_2030_v0).isLessThan(_2030_v0);
          let _rhs396 = new BigNumber(-241);
          let _lhs378 = _2048_v15;
          let _lhs379 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_2048_v15).length));
          let _lhs380 = globalState;
          let _lhs381 = _2055_v18;
          let _lhs382 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_2055_v18).length));
          _lhs378[_lhs379] = _rhs394;
          _lhs380.f15 = _rhs395;
          _lhs381[_lhs382] = _rhs396;
          let _2058_v19;
          _2058_v19 = _dafny.Seq.of((_2040_v7).f33);
          (globalState).f18 = (_2058_v19)[_module.__default.safeIndex((_2040_v7).f34, new BigNumber((_2058_v19).length))];
          (globalState).f15 = (_2040_v7).f33;
        } else {
          (globalState).f14 = _module.__default.fm0(_2030_v0, (_2040_v7).f33, (((_2044_v11).contains((_2040_v7).f33)) ? ((_2044_v11).get((_2040_v7).f33)) : ((_2040_v7).f33)), globalState);
          let _2059_v20;
          _2059_v20 = new _dafny.CodePoint('i'.codePointAt(0));
          _2059_v20 = _2059_v20;
          let _2060_v21;
          _2060_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2038_v5,new BigNumber(-52));
          let _2061_v22;
          _2061_v22 = _module.D10.create_DC32(_2059_v20, _module.__default.fm6((((_2060_v21).contains(_2038_v5)) ? ((_2060_v21).get(_2038_v5)) : ((_2040_v7).f34)), globalState));
          (globalState).f28 = (_2061_v22).dtor_cf53;
          let _2062_v23;
          let _nw361 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Set.Empty);
          _2062_v23 = _nw361;
          let _2063_v24;
          _2063_v24 = _dafny.Set.fromElements((_2040_v7).f33);
          let _index349 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_2062_v23).length));
          (_2062_v23)[_index349] = _2063_v24;
          let _2064_v25;
          _2064_v25 = _dafny.Seq.of((_2040_v7).f33, (_2040_v7).f33, _2038_v5);
          let _2065_v26;
          _2065_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2064_v25).length),(_2040_v7).f33);
          let _2066_v29;
          let _nw362 = Array((new BigNumber(21)).toNumber());
          _nw362[(_dafny.ZERO).toNumber()] = _2065_v26;
          _nw362[(_dafny.ONE).toNumber()] = _2065_v26;
          _nw362[(new BigNumber(2)).toNumber()] = _2065_v26;
          _nw362[(new BigNumber(3)).toNumber()] = _module.__default.fm41((_dafny.ZERO).minus((_2040_v7).f34), (_2040_v7).f33, !((_2040_v7).f33), _2038_v5, globalState);
          _nw362[(new BigNumber(4)).toNumber()] = function () {
            let _coll71 = new _dafny.Map();
            for (const _compr_71 of (p0).Elements) {
              let _2067_v27 = _compr_71;
              if (_dafny.Seq.contains(p0, _2067_v27)) {
                _coll71.push([(_2067_v27).plus((_2040_v7).f34),_2038_v5]);
              }
            }
            return _coll71;
          }();
          _nw362[(new BigNumber(5)).toNumber()] = _2065_v26;
          _nw362[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(354),(_2040_v7).f33)).update((_2040_v7).f34, (_2040_v7).f33);
          _nw362[(new BigNumber(7)).toNumber()] = _2065_v26;
          _nw362[(new BigNumber(8)).toNumber()] = function () {
            let _coll72 = new _dafny.Map();
            for (const _compr_72 of _dafny.IntegerRange(new BigNumber(-623), new BigNumber(476))) {
              let _2068_v28 = _compr_72;
              if (((new BigNumber(-623)).isLessThanOrEqualTo(_2068_v28)) && ((_2068_v28).isLessThan(new BigNumber(476)))) {
                _coll72.push([(_2068_v28).multipliedBy((_2040_v7).f34),(_2040_v7).f33]);
              }
            }
            return _coll72;
          }();
          _nw362[(new BigNumber(9)).toNumber()] = _2065_v26;
          _nw362[(new BigNumber(10)).toNumber()] = _2065_v26;
          _nw362[(new BigNumber(11)).toNumber()] = (_2065_v26).Merge(_2065_v26);
          _nw362[(new BigNumber(12)).toNumber()] = _2065_v26;
          _nw362[(new BigNumber(13)).toNumber()] = (_2065_v26).Merge(_2065_v26);
          _nw362[(new BigNumber(14)).toNumber()] = _2065_v26;
          _nw362[(new BigNumber(15)).toNumber()] = _2065_v26;
          _nw362[(new BigNumber(16)).toNumber()] = _2065_v26;
          _nw362[(new BigNumber(17)).toNumber()] = _2065_v26;
          _nw362[(new BigNumber(18)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_2040_v7).f34,(_2040_v7).f33)).Merge(_2065_v26);
          _nw362[(new BigNumber(19)).toNumber()] = _2065_v26;
          _nw362[(new BigNumber(20)).toNumber()] = (_2065_v26).Merge(_2065_v26);
          _2066_v29 = _nw362;
          let _index350 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_2066_v29).length));
          (_2066_v29)[_index350] = _2065_v26;
          let _2069_v30;
          let _nw363 = new _module.C3();
          _nw363.__ctor((_2040_v7).f33, (_2040_v7).f34, false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-792)), ((_2070_v20) => function (_2071_i5) {
            return _2070_v20;
          })(_2059_v20)));
          _2069_v30 = _nw363;
          let _2072_v31;
          _2072_v31 = _dafny.Seq.of(_2069_v30, _2069_v30);
          let _index351 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_2062_v23).length));
          let _index352 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_2066_v29).length));
          let _rhs397 = _2063_v24;
          let _rhs398 = (_2031_v1).update(_module.__default.fm36(globalState), _module.__default.abs((_dafny.ZERO).minus(new BigNumber(((_2060_v21).update(!(false), new BigNumber((_2072_v31).length))).length))));
          let _rhs399 = (_2064_v25)[_module.__default.safeIndex(_2030_v0, new BigNumber((_2064_v25).length))];
          let _rhs400 = _dafny.Seq.Concat(p0, p0);
          let _rhs401 = (_2065_v26).Merge(_2065_v26);
          let _lhs383 = _2062_v23;
          let _lhs384 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_2062_v23).length));
          let _lhs385 = globalState;
          let _lhs386 = globalState;
          let _lhs387 = _2066_v29;
          let _lhs388 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_2066_v29).length));
          _lhs383[_lhs384] = _rhs397;
          _2031_v1 = _rhs398;
          _lhs385.f13 = _rhs399;
          _lhs386.f7 = _rhs400;
          _lhs387[_lhs388] = _rhs401;
          let _2073_v32;
          let _nw364 = Array((new BigNumber(4)).toNumber());
          _nw364[(_dafny.ZERO).toNumber()] = (_2040_v7).f34;
          _nw364[(_dafny.ONE).toNumber()] = new BigNumber((_2039_v6).length);
          _nw364[(new BigNumber(2)).toNumber()] = (_2040_v7).f34;
          _nw364[(new BigNumber(3)).toNumber()] = (_2040_v7).f34;
          _2073_v32 = _nw364;
          let _2074_v33;
          _2074_v33 = _module.D12.create_DC38(_2073_v32);
          let _index353 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_2032_v2).length));
          (_2032_v2)[_index353] = (_2064_v25)[_module.__default.safeIndex((_dafny.ZERO).minus((_2040_v7).f34), new BigNumber((_2064_v25).length))];
          let _2075_v34;
          let _nw365 = new _module.C4();
          _nw365.__ctor((_2040_v7).f33, new BigNumber((_2039_v6).length));
          _2075_v34 = _nw365;
          let _index354 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_2032_v2).length));
          let _rhs402 = (((_2040_v7).f33) ? (_2074_v33) : (_2074_v33));
          let _rhs403 = !(_module.__default.fm1(new _dafny.CodePoint('o'.codePointAt(0)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2038_v5,new BigNumber(673))).length), (_2040_v7).f34, !(_2038_v5), globalState));
          let _rhs404 = new BigNumber(406);
          let _rhs405 = _2075_v34;
          let _rhs406 = _2060_v21;
          let _lhs389 = _2032_v2;
          let _lhs390 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_2032_v2).length));
          let _lhs391 = globalState;
          _2074_v33 = _rhs402;
          _lhs389[_lhs390] = _rhs403;
          _lhs391.f19 = _rhs404;
          _2075_v34 = _rhs405;
          _2060_v21 = _rhs406;
        }
      } else {
        let _2076_v35;
        let _nw366 = new _module.C5();
        _nw366.__ctor((_2040_v7).f33, _dafny.Seq.Concat(_2039_v6, _2039_v6));
        _2076_v35 = _nw366;
        let _nw367 = new _module.C5();
        _nw367.__ctor((_2040_v7).f33, _2039_v6);
        _2076_v35 = _nw367;
        r0 = new BigNumber((_2039_v6).length);
        let _2077_v36;
        let _nw368 = new _module.C4();
        _nw368.__ctor(_2038_v5, _2030_v0);
        _2077_v36 = _nw368;
        let _2078_v37;
        _2078_v37 = _dafny.Map.Empty.slice().updateUnsafe((_2040_v7).f34,_2077_v36);
        let _2079_v38;
        let _nw369 = Array((new BigNumber(24)).toNumber());
        _nw369[(_dafny.ZERO).toNumber()] = _2077_v36;
        _nw369[(_dafny.ONE).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(2)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(3)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(4)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(5)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(6)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(7)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(8)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(9)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(10)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(11)).toNumber()] = (((_2078_v37).contains((_2040_v7).f34)) ? ((_2078_v37).get((_2040_v7).f34)) : (_2077_v36));
        _nw369[(new BigNumber(12)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(13)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(14)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(15)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(16)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(17)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(18)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(19)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(20)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(21)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(22)).toNumber()] = _2077_v36;
        _nw369[(new BigNumber(23)).toNumber()] = _2077_v36;
        _2079_v38 = _nw369;
        let _2080_v39;
        _2080_v39 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_2039_v6, _2039_v6),_2079_v38);
        _2080_v39 = (_2080_v39).update(_2039_v6, ((_2038_v5) ? (_2079_v38) : (_2079_v38)));
        if ((!(((true) ? ((_2040_v7).f33) : ((_2040_v7).f33)))) && ((_2040_v7).f33)) {
          (globalState).f14 = _module.__default.safeModuloInt(_2030_v0, (_2040_v7).f34);
          let _2081_v40;
          _2081_v40 = _dafny.MultiSet.fromElements(_2038_v5, (_2040_v7).f33, (_2040_v7).f33, (_2040_v7).f33, (_2040_v7).f33);
          (globalState).f28 = ((_2081_v40).Intersect(_2081_v40)).IsDisjointFrom((_2081_v40).update((_2040_v7).f33, _module.__default.abs(_2030_v0)));
          let _2082_v41;
          _2082_v41 = _dafny.Set.fromElements((_2040_v7).f33);
          let _2083_v42;
          _2083_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2082_v41).length),new BigNumber((_2039_v6).length));
          let _2084_v43;
          let _nw370 = Array((new BigNumber(19)).toNumber());
          _nw370[(_dafny.ZERO).toNumber()] = new BigNumber(((_2031_v1).Difference(_2031_v1)).cardinality());
          _nw370[(_dafny.ONE).toNumber()] = (_2040_v7).f34;
          _nw370[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt((_2040_v7).f34, (_2040_v7).f34);
          _nw370[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("kljunxavc")).length);
          _nw370[(new BigNumber(4)).toNumber()] = (_2040_v7).f34;
          _nw370[(new BigNumber(5)).toNumber()] = (_2040_v7).f34;
          _nw370[(new BigNumber(6)).toNumber()] = new BigNumber(-367);
          _nw370[(new BigNumber(7)).toNumber()] = _2030_v0;
          _nw370[(new BigNumber(8)).toNumber()] = (((_2083_v42).contains(new BigNumber(179))) ? ((_2083_v42).get(new BigNumber(179))) : ((_2040_v7).f34));
          _nw370[(new BigNumber(9)).toNumber()] = _2030_v0;
          _nw370[(new BigNumber(10)).toNumber()] = (_2040_v7).f34;
          _nw370[(new BigNumber(11)).toNumber()] = (_2040_v7).f34;
          _nw370[(new BigNumber(12)).toNumber()] = (_2040_v7).f34;
          _nw370[(new BigNumber(13)).toNumber()] = _module.__default.fm36(globalState);
          _nw370[(new BigNumber(14)).toNumber()] = _2030_v0;
          _nw370[(new BigNumber(15)).toNumber()] = (_2040_v7).f34;
          _nw370[(new BigNumber(16)).toNumber()] = _module.__default.safeDivisionInt((_2040_v7).f34, (_2040_v7).f34);
          _nw370[(new BigNumber(17)).toNumber()] = (((_2040_v7).f33) ? (new BigNumber((_dafny.Seq.UnicodeFromString("xgc")).length)) : ((_2040_v7).f34));
          _nw370[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_2040_v7).f34));
          _2084_v43 = _nw370;
          (globalState).f17 = _2084_v43;
          (globalState).f14 = _module.__default.safeDivisionInt(_2030_v0, (_2040_v7).f34);
          let _index355 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_2084_v43).length));
          (_2084_v43)[_index355] = (_dafny.ZERO).minus(_2030_v0);
          let _index356 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_2084_v43).length));
          (_2084_v43)[_index356] = _2030_v0;
          let _index357 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_2084_v43).length));
          let _index358 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_2084_v43).length));
          let _rhs407 = (_2040_v7).f34;
          let _rhs408 = new BigNumber(309);
          let _lhs392 = _2084_v43;
          let _lhs393 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_2084_v43).length));
          let _lhs394 = _2084_v43;
          let _lhs395 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_2084_v43).length));
          _lhs392[_lhs393] = _rhs407;
          _lhs394[_lhs395] = _rhs408;
        } else {
          let _2085_v44;
          let _init56 = ((_2086_v6) => function (_2087_i6) {
            return _2086_v6;
          })(_2039_v6);
          let _nw371 = Array((new BigNumber(17)).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw371.length); _i0_56++) {
            _nw371[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _2085_v44 = _nw371;
          let _index359 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_2085_v44).length));
          (_2085_v44)[_index359] = _2039_v6;
          let _index360 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_2085_v44).length));
          (_2085_v44)[_index360] = _dafny.Seq.Concat(_2039_v6, _2039_v6);
          (globalState).f15 = true;
          (globalState).f15 = (_2040_v7).f33;
          (globalState).f20 = _dafny.Seq.Concat((_2085_v44)[_module.__default.safeIndex(new BigNumber(419), new BigNumber((_2085_v44).length))], (_2085_v44)[_module.__default.safeIndex(new BigNumber(419), new BigNumber((_2085_v44).length))]);
          (globalState).f1 = new BigNumber((function () {
            let _coll73 = new _dafny.Map();
            for (const _compr_73 of (_2031_v1).Elements) {
              let _2088_v45 = _compr_73;
              if ((_2031_v1).contains(_2088_v45)) {
                _coll73.push([(_2088_v45).multipliedBy((_2040_v7).f34),_dafny.Map.Empty.slice().updateUnsafe((_2040_v7).f34,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
                  let _coll74 = new _dafny.Set();
                  for (const _compr_74 of (_2031_v1).Elements) {
                    let _2089_v46 = _compr_74;
                    if ((_2031_v1).contains(_2089_v46)) {
                      _coll74.add((_2089_v46).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).length)));
                    }
                  }
                  return _coll74;
                }()).length))).length))]);
              }
            }
            return _coll73;
          }()).length);
        }
        let _2090_v47;
        _2090_v47 = _dafny.MultiSet.fromElements(false);
        let _2091_v48;
        _2091_v48 = _dafny.Map.Empty.slice().updateUnsafe((_2040_v7).f34,_2090_v47);
        let _2092_v49;
        _2092_v49 = _dafny.Seq.of(_2039_v6);
        let _2093_v50;
        _2093_v50 = _module.D0.create_DC2((((_2091_v48).contains(_2030_v0)) ? ((_2091_v48).get(_2030_v0)) : (_2090_v47)), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("umiruj"), (_2092_v49)[_module.__default.safeIndex(_2030_v0, new BigNumber((_2092_v49).length))]), (_2040_v7).f33);
        _2093_v50 = _module.D0.create_DC2(_dafny.MultiSet.fromElements((_2040_v7).f33, !((_2040_v7).f33), (_2040_v7).f33), _2039_v6, false);
      }
      let _hi12 = _2030_v0;
      for (let _2094_i7 = (_2030_v0).minus((_dafny.ZERO).minus((_2040_v7).f34)); _2094_i7.isLessThan(_hi12); _2094_i7 = _2094_i7.plus(_dafny.ONE)) {
        let _2095_v51;
        let _nw372 = Array((new BigNumber(12)).toNumber());
        _2095_v51 = _nw372;
        let _2096_v52;
        let _init57 = ((_2097_v0) => function (_2098_i8) {
          return (_2098_i8).plus(_2097_v0);
        })(_2030_v0);
        let _nw373 = Array((new BigNumber(12)).toNumber());
        for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw373.length); _i0_57++) {
          _nw373[_i0_57] = _init57(new BigNumber(_i0_57));
        }
        _2096_v52 = _nw373;
        let _2099_v53;
        _2099_v53 = _dafny.Set.fromElements(_2096_v52, _2096_v52);
        let _2100_v54;
        _2100_v54 = new _dafny.CodePoint('r'.codePointAt(0));
        let _2101_v55;
        let _nw374 = new _module.C2();
        _nw374.__ctor(new BigNumber((_dafny.MultiSet.fromElements(_2040_v7)).cardinality()), _2099_v53, _2038_v5, _dafny.Seq.update(_2039_v6, _module.__default.safeIndex(new BigNumber((_2039_v6).length), new BigNumber((_2039_v6).length)), _2100_v54));
        _2101_v55 = _nw374;
        let _index361 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_2095_v51).length));
        (_2095_v51)[_index361] = _2101_v55;
        let _index362 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_2095_v51).length));
        (_2095_v51)[_index362] = _2101_v55;
        let _2102_v56;
        _2102_v56 = _dafny.Map.Empty.slice().updateUnsafe(_2038_v5,new BigNumber(-916));
        let _2103_v57;
        let _nw375 = new _module.C3();
        _nw375.__ctor((_2040_v7).f33, new BigNumber(941), !(_2102_v56).equals(_2102_v56), _2039_v6);
        _2103_v57 = _nw375;
        _2103_v57 = _2103_v57;
        let _2104_v58;
        _2104_v58 = _dafny.MultiSet.fromElements((_2103_v57).f29);
        let _2105_v59;
        _2105_v59 = _module.D11.create_DC37(new BigNumber((_2104_v58).cardinality()), _2038_v5);
        (globalState).f2 = (function (_pat_let48_0) {
          return function (_2106_dt__update__tmp_h0) {
            return function (_pat_let49_0) {
              return function (_2107_dt__update_hcf63_h0) {
                return _module.D11.create_DC37((_2106_dt__update__tmp_h0).dtor_cf62, _2107_dt__update_hcf63_h0);
              }(_pat_let49_0);
            }(false);
          }(_pat_let48_0);
        }(_2105_v59)).dtor_cf62;
        let _2108_v60;
        let _nw376 = new _module.C1();
        _nw376.__ctor(false);
        _2108_v60 = _nw376;
        let _2109_v61;
        _2109_v61 = _dafny.Map.Empty.slice().updateUnsafe(_2108_v60,!((_2101_v55).f38).isEqualTo((_2101_v55).f38));
        _2109_v61 = (_2109_v61).update(_2108_v60, _2038_v5);
      }
      r0 = _2030_v0;
      return r0;
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
