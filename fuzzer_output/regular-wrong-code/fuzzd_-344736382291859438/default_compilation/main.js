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
      return _module.__default.safeModuloInt(new BigNumber(280), (new BigNumber(-786)).multipliedBy((_dafny.ZERO).minus(new BigNumber(-585))));
    };
    static fm7(p0, p1, globalState) {
      return _dafny.Set.fromElements(true);
    };
    static fm10(p0, p1, p2, globalState) {
      if (((_dafny.ZERO).minus(new BigNumber(-465))).isLessThanOrEqualTo(new BigNumber(701))) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(377)), function (_0_i0) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        });
      } else {
        return _dafny.Seq.UnicodeFromString("xyrvtep");
      }
    };
    static fm11(p0, globalState) {
      return new _dafny.CodePoint('e'.codePointAt(0));
    };
    static fm12(p0, p1, globalState) {
      return !((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('l'.codePointAt(0)))).length),!(true)))).IsDisjointFrom(((false) ? (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),true)))) : (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("vkne")).length),false), function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(828), new BigNumber(917))) {
          let _1_v0 = _compr_0;
          if (((new BigNumber(828)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(917)))) {
            _coll0.push([(_1_v0).minus(new BigNumber(890)),!(false)]);
          }
        }
        return _coll0;
      }(), function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length),true)).Keys.Elements) {
          let _2_v1 = _compr_1;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length),true)).contains(_2_v1)) {
            _coll1.push([(_2_v1).multipliedBy(new BigNumber((_dafny.Seq.of(true)).length)),false]);
          }
        }
        return _coll1;
      }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("iwq")).length))).length),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(562),false)))))));
    };
    static fm14(p0, p1, p2, globalState) {
      return _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(606)), function (_3_i0) {
        return _dafny.Seq.UnicodeFromString("gcrwhf");
      })).length), (new BigNumber(83)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(518),new _dafny.CodePoint('k'.codePointAt(0)))).length)), ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(641),false)).length))).minus((_dafny.ZERO).minus(new BigNumber((function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(518), new BigNumber(111))) {
          let _4_v0 = _compr_2;
          if (((new BigNumber(518)).isLessThanOrEqualTo(_4_v0)) && ((_4_v0).isLessThan(new BigNumber(111)))) {
            _coll2.add((_4_v0).plus(new BigNumber(828)));
          }
        }
        return _coll2;
      }()).length))));
    };
    static fm15(p0, p1, p2, globalState) {
      return _module.D3.create_DC10(_dafny.MultiSet.fromElements(false, true));
    };
    static fm16(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(!(false)));
    };
    static fm17(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(490)), function (_5_i0) {
        return new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(757), new BigNumber(86))).cardinality());
      }))).Difference(_dafny.MultiSet.fromElements(new BigNumber(488)))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(880)), function (_6_i1) {
        return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-421)), function (_7_i2) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        })).length);
      })));
    };
    static fm18(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(914),!(true))).Merge(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).Elements) {
          let _8_v0 = _compr_3;
          if ((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).contains(_8_v0)) {
            _coll3.push([(_8_v0).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("wnrwxf")).length))).length))).cardinality())),true]);
          }
        }
        return _coll3;
      }());
    };
    static fm19(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(true,true);
    };
    static fm20(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(true), false, true, false)).length)))).length), new BigNumber(671)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(996)), function (_9_i0) {
        return new BigNumber(-936);
      }))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(125))).length)),_dafny.Seq.of(new BigNumber(443), new BigNumber(719))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(931)),_dafny.Seq.of(new BigNumber(856))));
    };
    static fm22(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true, false, !(true)), _dafny.Seq.of(!(true)));
    };
    static fm23(p0, p1, p2, globalState) {
      let _source0 = _module.D12.create_DC30(_dafny.Seq.UnicodeFromString("oausrgoo"), new BigNumber(442), new _dafny.CodePoint('l'.codePointAt(0)));
      if (_source0.is_DC29) {
        let _10___mcc_h0 = (_source0).cf45;
        let _11___mcc_h1 = (_source0).cf46;
        let _12___mcc_h2 = (_source0).cf47;
        let _13_cf47 = _12___mcc_h2;
        let _14_cf46 = _11___mcc_h1;
        let _15_cf45 = _10___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(494)), function (_16_i0) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(539)), function (_17_i1) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }));
      } else if (_source0.is_DC30) {
        let _18___mcc_h3 = (_source0).cf48;
        let _19___mcc_h4 = (_source0).cf49;
        let _20___mcc_h5 = (_source0).cf50;
        let _21_cf50 = _20___mcc_h5;
        let _22_cf49 = _19___mcc_h4;
        let _23_cf48 = _18___mcc_h3;
        return _dafny.Seq.Concat(_23_cf48, _dafny.Seq.update(_23_cf48, _module.__default.safeIndex(new BigNumber(-209), new BigNumber((_23_cf48).length)), new _dafny.CodePoint('o'.codePointAt(0))));
      } else if (_source0.is_DC31) {
        let _24___mcc_h6 = (_source0).cf51;
        let _25___mcc_h7 = (_source0).cf52;
        let _26___mcc_h8 = (_source0).cf53;
        let _27___mcc_h9 = (_source0).cf54;
        let _28_cf54 = _27___mcc_h9;
        let _29_cf53 = _26___mcc_h8;
        let _30_cf52 = _25___mcc_h7;
        let _31_cf51 = _24___mcc_h6;
        return _dafny.Seq.UnicodeFromString("xdvd");
      } else {
        let _32___mcc_h10 = (_source0).cf44;
        let _33_cf44 = _32___mcc_h10;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(995)), function (_34_i2) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        });
      }
    };
    static fm24(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(false, !(!(false)), false, true)).Difference((_dafny.MultiSet.fromElements(false)).Intersect(_dafny.MultiSet.fromElements(true)));
    };
    static fm25(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC5(new BigNumber(((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ptgi"))).Union(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(474)), function (_35_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
})))).length), false, !(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_dafny.MultiSet.fromElements(false, true)), _dafny.Seq.of(_dafny.MultiSet.fromElements(false, !(true), true, true, !(false)), _dafny.MultiSet.fromElements(!(true), true)))), _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(77)),_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true, true)).length))));
    };
    static fm26(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(479),false)).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true)).length))).cardinality()), new BigNumber(538), new BigNumber(754), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, true)).length))).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.Create(_module.__default.abs(new BigNumber(949)), function (_36_i0) {
        return new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)))).length))).length);
      })))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber(230), new BigNumber(-506))));
    };
    static fm27(p0, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(105)), function (_37_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }), _dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)))), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(486)), function (_38_i1) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(660)), function (_39_i2) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })), _dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0))));
    };
    static fm28(p0, globalState) {
      return (_dafny.Set.fromElements(true)).Difference(_dafny.Set.fromElements(!(false), true));
    };
    static fm29(globalState) {
      return _module.D1.create_DC6(new BigNumber(538), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(!(false), !(true), false, false, false), _dafny.Seq.of(true)), _dafny.Seq.of(_dafny.Seq.of(false))), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-211),true)).length), false);
    };
    static fm30(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true));
    };
    static fm32(p0, p1, p2, globalState) {
      return _dafny.MultiSet.FromArray(((true) ? (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(974)), function (_40_i0) {
        return _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(507))).length)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("xwenmsl")).length))).cardinality()))).length))).length));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-309)), function (_41_i1) {
        return _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length));
      }))) : (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(904)), function (_42_i2) {
        return new BigNumber(-134);
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-210)), function (_43_i3) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("qqmyop")).length);
      }), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_module.D15.create_DC37(function () {
  let _coll4 = new _dafny.Map();
  for (const _compr_4 of (function () {
    let _coll5 = new _dafny.Map();
    for (const _compr_5 of (_dafny.Set.fromElements(new _dafny.CodePoint('g'.codePointAt(0)))).Elements) {
      let _44_v1 = _compr_5;
      if ((_dafny.Set.fromElements(new _dafny.CodePoint('g'.codePointAt(0)))).contains(_44_v1)) {
        _coll5.push([_44_v1,new BigNumber((_dafny.Seq.of(!(true))).length)]);
      }
    }
    return _coll5;
  }()).Keys.Elements) {
    let _45_v0 = _compr_4;
    if ((function () {
      let _coll6 = new _dafny.Map();
      for (const _compr_6 of (_dafny.Set.fromElements(new _dafny.CodePoint('g'.codePointAt(0)))).Elements) {
        let _44_v1 = _compr_6;
        if ((_dafny.Set.fromElements(new _dafny.CodePoint('g'.codePointAt(0)))).contains(_44_v1)) {
          _coll6.push([_44_v1,new BigNumber((_dafny.Seq.of(!(true))).length)]);
        }
      }
      return _coll6;
    }()).contains(_45_v0)) {
      _coll4.push([_45_v0,!(!(true))]);
    }
  }
  return _coll4;
}())).dtor_cf64).length),true)).length))))));
    };
    static fm33(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(805),new BigNumber(-464))).Merge(function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(301))).Keys.Elements) {
          let _46_v0 = _compr_7;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(301))).contains(_46_v0)) {
            _coll7.push([(_46_v0).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D12.create_DC28(_dafny.MultiSet.fromElements(new BigNumber(-568), new BigNumber(314)))).dtor_cf44,new BigNumber(341))).length)),new BigNumber(503)]);
          }
        }
        return _coll7;
      }());
    };
    static fm34(p0, p1, p2, globalState) {
      return _dafny.Seq.of(false, _dafny.areEqual(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(507))).length), new BigNumber((_dafny.Seq.UnicodeFromString("ghvpdnr")).length), new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),true)),new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).length)), new BigNumber(179))).length)), _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(504)), function (_47_i0) {
        return new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).length), new BigNumber((_dafny.Seq.of(new BigNumber(94), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(449),new BigNumber((_dafny.Seq.of(new BigNumber(985), new BigNumber(165))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).length), new BigNumber(425))).length))).length))).length);
      })), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(188),new BigNumber(-128))).length), new BigNumber((_dafny.Seq.UnicodeFromString("ahqc")).length))), _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(942)))));
    };
    static fm35(p0, p1, p2, globalState) {
      return _dafny.Seq.UnicodeFromString("cxtyubjr");
    };
    static fm36(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(758)), function (_48_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })).length))), _dafny.Seq.of(function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)))).Elements) {
          let _49_v0 = _compr_8;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0))), _49_v0)) {
            _coll8.push([_49_v0,new BigNumber(386)]);
          }
        }
        return _coll8;
      }())), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber(538)), _dafny.Map.Empty.slice().updateUnsafe((_module.D4.create_DC13(new _dafny.CodePoint('e'.codePointAt(0)))).dtor_cf24,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(762), new BigNumber((_dafny.Seq.UnicodeFromString("quneatwcc")).length))).cardinality()))));
    };
    static fm37(p0, p1, p2, globalState) {
      let _source1 = _module.D15.create_DC38();
      if (_source1.is_DC38) {
        return _module.D7.create_DC18(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(-356))));
      } else {
        let _50___mcc_h0 = (_source1).cf64;
        let _51_cf64 = _50___mcc_h0;
        return _module.D7.create_DC18(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),new BigNumber(217))));
      }
    };
    static fm38(p0, p1, globalState) {
      return ((_dafny.Set.fromElements(true)).Intersect(_dafny.Set.fromElements(!(false)))).Difference(_dafny.Set.fromElements(false, false, false, false));
    };
    static fm39(globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(426), new BigNumber(976), new BigNumber((_dafny.Seq.UnicodeFromString("rkh")).length)))).Difference(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-561)), function (_52_i0) {
        return new BigNumber((function () {
          let _coll9 = new _dafny.Set();
          for (const _compr_9 of _dafny.IntegerRange(new BigNumber(562), new BigNumber(427))) {
            let _53_v0 = _compr_9;
            if (((new BigNumber(562)).isLessThanOrEqualTo(_53_v0)) && ((_53_v0).isLessThan(new BigNumber(427)))) {
              _coll9.add(_module.__default.safeDivisionInt(_53_v0, new BigNumber((_dafny.Seq.of(new BigNumber(357), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()))).length)));
            }
          }
          return _coll9;
        }()).length);
      })))).Intersect((_dafny.Set.fromElements(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ddtk")).length))).length))))).Intersect(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(80)), _dafny.Seq.of(new BigNumber(798)), _dafny.Seq.of(new BigNumber(-939)))));
    };
    static fm40(p0, p1, p2, globalState) {
      return _module.D3.create_DC11(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(560),true)).length));
    };
    static fm41(p0, p1, p2, p3, globalState) {
      let _source2 = _module.D12.create_DC31(new BigNumber((_dafny.Seq.of(false, (_module.D12.create_DC29(_dafny.Seq.of(new BigNumber(887), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(355)), function (_54_i0) {
  return new _dafny.CodePoint('c'.codePointAt(0));
})).length)), new BigNumber((function () {
  let _coll10 = new _dafny.Map();
  for (const _compr_10 of _dafny.IntegerRange(new BigNumber(273), new BigNumber(623))) {
    let _55_v0 = _compr_10;
    if (((new BigNumber(273)).isLessThanOrEqualTo(_55_v0)) && ((_55_v0).isLessThan(new BigNumber(623)))) {
      _coll10.push([(_55_v0).plus(new BigNumber(518)),new BigNumber((_dafny.Seq.UnicodeFromString("dbeopaxvs")).length)]);
    }
  }
  return _coll10;
}()).length), false)).dtor_cf47)).length), false, false, new BigNumber((function () {
  let _coll11 = new _dafny.Map();
  for (const _compr_11 of _dafny.IntegerRange(new BigNumber(478), new BigNumber(125))) {
    let _56_v1 = _compr_11;
    if (((new BigNumber(478)).isLessThanOrEqualTo(_56_v1)) && ((_56_v1).isLessThan(new BigNumber(125)))) {
      _coll11.push([_module.__default.safeDivisionInt(_56_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-611))).length)),new _dafny.CodePoint('g'.codePointAt(0))]);
    }
  }
  return _coll11;
}()).length));
      if (_source2.is_DC29) {
        let _57___mcc_h0 = (_source2).cf45;
        let _58___mcc_h1 = (_source2).cf46;
        let _59___mcc_h2 = (_source2).cf47;
        let _60_cf47 = _59___mcc_h2;
        let _61_cf46 = _58___mcc_h1;
        let _62_cf45 = _57___mcc_h0;
        return _module.D4.create_DC13(new _dafny.CodePoint('i'.codePointAt(0)));
      } else if (_source2.is_DC30) {
        let _63___mcc_h3 = (_source2).cf48;
        let _64___mcc_h4 = (_source2).cf49;
        let _65___mcc_h5 = (_source2).cf50;
        let _66_cf50 = _65___mcc_h5;
        let _67_cf49 = _64___mcc_h4;
        let _68_cf48 = _63___mcc_h3;
        return _module.D4.create_DC13(_66_cf50);
      } else if (_source2.is_DC31) {
        let _69___mcc_h6 = (_source2).cf51;
        let _70___mcc_h7 = (_source2).cf52;
        let _71___mcc_h8 = (_source2).cf53;
        let _72___mcc_h9 = (_source2).cf54;
        let _73_cf54 = _72___mcc_h9;
        let _74_cf53 = _71___mcc_h8;
        let _75_cf52 = _70___mcc_h7;
        let _76_cf51 = _69___mcc_h6;
        return _module.D4.create_DC13(new _dafny.CodePoint('t'.codePointAt(0)));
      } else {
        let _77___mcc_h10 = (_source2).cf44;
        let _78_cf44 = _77___mcc_h10;
        return _module.D4.create_DC13(new _dafny.CodePoint('c'.codePointAt(0)));
      }
    };
    static fm42(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(new BigNumber(130), (new BigNumber(483)).multipliedBy(new BigNumber(460)), (new BigNumber(954)).minus(new BigNumber((function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("iatc")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
          let _coll13 = new _dafny.Map();
          for (const _compr_13 of _dafny.IntegerRange(new BigNumber(322), new BigNumber(477))) {
            let _79_v1 = _compr_13;
            if (((new BigNumber(322)).isLessThanOrEqualTo(_79_v1)) && ((_79_v1).isLessThan(new BigNumber(477)))) {
              _coll13.push([(_79_v1).minus(new BigNumber(26)),false]);
            }
          }
          return _coll13;
        }()).length),true)).length), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).Elements) {
          let _80_v0 = _compr_12;
          if ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("iatc")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll14 = new _dafny.Map();
            for (const _compr_14 of _dafny.IntegerRange(new BigNumber(322), new BigNumber(477))) {
              let _79_v1 = _compr_14;
              if (((new BigNumber(322)).isLessThanOrEqualTo(_79_v1)) && ((_79_v1).isLessThan(new BigNumber(477)))) {
                _coll14.push([(_79_v1).minus(new BigNumber(26)),false]);
              }
            }
            return _coll14;
          }()).length),true)).length), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).contains(_80_v0)) {
            _coll12.push([(_80_v0).plus(new BigNumber(905)),false]);
          }
        }
        return _coll12;
      }()).length)), new BigNumber(997), ((_dafny.ZERO).minus(new BigNumber(-292))).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(253))).length), new BigNumber(-790)),new BigNumber(334))).length))));
    };
    static fm43(p0, p1, globalState) {
      if (!(new BigNumber(908)).isEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("qpimlci")).length))) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(668),_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-988))));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('x'.codePointAt(0)))).cardinality()),_dafny.Seq.Create(_module.__default.abs(new BigNumber(956)), function (_81_i0) {
          return new BigNumber(251);
        }));
      }
    };
    static fm44(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("nlcjilh"),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("cec"),false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("q"),true));
    };
    static fm45(p0, globalState) {
      return _dafny.MultiSet.fromElements((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(480))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(554))));
    };
    static fm46(p0, p1, p2, globalState) {
      if ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)).isLessThanOrEqualTo(new BigNumber(-344))) {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-546)), function (_82_i0) {
          return _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(102));
        }), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(true)).length))));
      } else {
        return _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(-804)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(625)));
      }
    };
    static m8(globalState) {
      let _83_v0;
      _83_v0 = false;
      let _84_v1;
      _84_v1 = _dafny.Seq.of(_83_v0, true, _83_v0, _83_v0, _83_v0);
      let _85_v2;
      _85_v2 = new BigNumber(667);
      let _86_v3;
      _86_v3 = _dafny.Seq.UnicodeFromString("rcmkx");
      (globalState).f13 = ((_dafny.ZERO).minus(new BigNumber((_84_v1).length))).isLessThan((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_85_v2, new BigNumber((_86_v3).length))));
      let _87_v4;
      let _nw0 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _87_v4 = _nw0;
      let _88_v5;
      _88_v5 = _87_v4;
      let _source3 = _88_v5;
      let _89___mcc_h0 = _source3;
      let _90_cf42 = _89___mcc_h0;
      let _91_v6;
      let _nw1 = Array((new BigNumber(23)).toNumber()).fill(false);
      _91_v6 = _nw1;
      let _92_v7;
      _92_v7 = _module.D4.create_DC14(_83_v0);
      let _93_v8;
      _93_v8 = _module.D4.create_DC15(_92_v7);
      let _94_v9;
      _94_v9 = _dafny.MultiSet.fromElements(_85_v2);
      let _95_v10;
      let _nw2 = Array((new BigNumber(10)).toNumber());
      _nw2[(_dafny.ZERO).toNumber()] = new BigNumber(-911);
      _nw2[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_85_v2);
      _nw2[(new BigNumber(2)).toNumber()] = new BigNumber((_94_v9).cardinality());
      _nw2[(new BigNumber(3)).toNumber()] = new BigNumber(994);
      _nw2[(new BigNumber(4)).toNumber()] = _85_v2;
      _nw2[(new BigNumber(5)).toNumber()] = _85_v2;
      _nw2[(new BigNumber(6)).toNumber()] = _85_v2;
      _nw2[(new BigNumber(7)).toNumber()] = _85_v2;
      _nw2[(new BigNumber(8)).toNumber()] = _85_v2;
      _nw2[(new BigNumber(9)).toNumber()] = _85_v2;
      _95_v10 = _nw2;
      let _96_v11;
      _96_v11 = _95_v10;
      let _97_v12;
      let _nw3 = Array((new BigNumber(14)).toNumber());
      _nw3[(_dafny.ZERO).toNumber()] = _95_v10;
      _nw3[(_dafny.ONE).toNumber()] = _95_v10;
      _nw3[(new BigNumber(2)).toNumber()] = _95_v10;
      _nw3[(new BigNumber(3)).toNumber()] = (_96_v11);
      _nw3[(new BigNumber(4)).toNumber()] = _95_v10;
      _nw3[(new BigNumber(5)).toNumber()] = _95_v10;
      _nw3[(new BigNumber(6)).toNumber()] = _95_v10;
      _nw3[(new BigNumber(7)).toNumber()] = _95_v10;
      _nw3[(new BigNumber(8)).toNumber()] = _95_v10;
      _nw3[(new BigNumber(9)).toNumber()] = _95_v10;
      _nw3[(new BigNumber(10)).toNumber()] = _95_v10;
      _nw3[(new BigNumber(11)).toNumber()] = _95_v10;
      _nw3[(new BigNumber(12)).toNumber()] = _95_v10;
      _nw3[(new BigNumber(13)).toNumber()] = _95_v10;
      _97_v12 = _nw3;
      let _98_v13;
      let _nw4 = new _module.C5();
      _nw4.__ctor(_97_v12, _85_v2);
      _98_v13 = _nw4;
      let _99_v14;
      _99_v14 = _dafny.Seq.of(_90_cf42);
      let _100_v15;
      _100_v15 = _dafny.Map.Empty.slice().updateUnsafe(_85_v2,_85_v2);
      let _101_v16;
      _101_v16 = new _dafny.CodePoint('i'.codePointAt(0));
      let _102_v17;
      let _nw5 = Array((new BigNumber(28)).toNumber());
      _nw5[(_dafny.ZERO).toNumber()] = new BigNumber(670);
      _nw5[(_dafny.ONE).toNumber()] = _module.__default.fm0(new BigNumber((_86_v3).length), globalState);
      _nw5[(new BigNumber(2)).toNumber()] = ((_83_v0) ? (new BigNumber(203)) : (_85_v2));
      _nw5[(new BigNumber(3)).toNumber()] = _85_v2;
      _nw5[(new BigNumber(4)).toNumber()] = new BigNumber(-330);
      _nw5[(new BigNumber(5)).toNumber()] = _85_v2;
      _nw5[(new BigNumber(6)).toNumber()] = _85_v2;
      _nw5[(new BigNumber(7)).toNumber()] = (_module.D7.create_DC20(_dafny.Seq.Create(_module.__default.abs(new BigNumber(133)), ((_103_v2) => function (_104_i0) {
  return _103_v2;
})(_85_v2)), _91_v6, _93_v8, _83_v0, _85_v2)).dtor_cf34;
      _nw5[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_85_v2);
      _nw5[(new BigNumber(9)).toNumber()] = new BigNumber((_84_v1).length);
      _nw5[(new BigNumber(10)).toNumber()] = _85_v2;
      _nw5[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_98_v13,_91_v6)).length);
      _nw5[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(_85_v2);
      _nw5[(new BigNumber(13)).toNumber()] = _85_v2;
      _nw5[(new BigNumber(14)).toNumber()] = _85_v2;
      _nw5[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_83_v0, _module.__default.fm12(_83_v0, _83_v0, globalState), _83_v0, !(false)), _module.__default.safeIndex(_85_v2, new BigNumber((_dafny.Seq.of(_83_v0, _module.__default.fm12(_83_v0, _83_v0, globalState), _83_v0, !(false))).length)), _83_v0)).length);
      _nw5[(new BigNumber(16)).toNumber()] = _85_v2;
      _nw5[(new BigNumber(17)).toNumber()] = _85_v2;
      _nw5[(new BigNumber(18)).toNumber()] = _85_v2;
      _nw5[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_99_v14, _99_v14)).length);
      _nw5[(new BigNumber(20)).toNumber()] = _85_v2;
      _nw5[(new BigNumber(21)).toNumber()] = new BigNumber(115);
      _nw5[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm44(_85_v2, new BigNumber(361), _dafny.Seq.update(_dafny.Seq.update(_86_v3, _module.__default.safeIndex(_module.__default.fm0(new BigNumber((_100_v15).length), globalState), new BigNumber((_86_v3).length)), _101_v16), _module.__default.safeIndex((_98_v13).fm2(globalState), new BigNumber((_dafny.Seq.update(_86_v3, _module.__default.safeIndex(_module.__default.fm0(new BigNumber((_100_v15).length), globalState), new BigNumber((_86_v3).length)), _101_v16)).length)), _101_v16), globalState)).length));
      _nw5[(new BigNumber(23)).toNumber()] = new BigNumber(-60);
      _nw5[(new BigNumber(24)).toNumber()] = _85_v2;
      _nw5[(new BigNumber(25)).toNumber()] = _module.__default.safeModuloInt(_85_v2, new BigNumber((_86_v3).length));
      _nw5[(new BigNumber(26)).toNumber()] = _85_v2;
      _nw5[(new BigNumber(27)).toNumber()] = _85_v2;
      _102_v17 = _nw5;
      let _105_v18;
      _105_v18 = _dafny.Set.fromElements(_85_v2);
      let _index0 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_95_v10).length));
      (_95_v10)[_index0] = _module.__default.safeModuloInt(new BigNumber((_105_v18).length), _85_v2);
      let _index1 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_102_v17).length));
      (_102_v17)[_index1] = _85_v2;
      let _106_v19;
      let _nw6 = Array((new BigNumber(27)).toNumber()).fill([]);
      _106_v19 = _nw6;
      let _107_v20;
      let _nw7 = Array((new BigNumber(4)).toNumber()).fill([]);
      _107_v20 = _nw7;
      let _index2 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_106_v19).length));
      (_106_v19)[_index2] = _107_v20;
      let _108_v21;
      let _nw8 = new _module.C5();
      _nw8.__ctor(_97_v12, (_dafny.ZERO).minus(_85_v2));
      _108_v21 = _nw8;
      let _109_v22;
      _109_v22 = _dafny.Set.fromElements(_108_v21, _108_v21);
      let _110_v23;
      _110_v23 = _dafny.Map.Empty.slice().updateUnsafe(_83_v0,_83_v0);
      let _111_v24;
      _111_v24 = _dafny.MultiSet.fromElements(_83_v0, _83_v0);
      let _112_v25;
      _112_v25 = _dafny.Map.Empty.slice().updateUnsafe(_101_v16,(_111_v24).IsDisjointFrom(_module.__default.fm24(_85_v2, _83_v0, globalState)));
      let _113_v26;
      _113_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(143),_107_v20);
      let _index3 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_95_v10).length));
      let _index4 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_102_v17).length));
      let _index5 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_106_v19).length));
      let _rhs0 = !((((_109_v22).IsSubsetOf(_109_v22)) ? (_83_v0) : ((((_110_v23).contains(_83_v0)) ? ((_110_v23).get(_83_v0)) : (_83_v0)))));
      let _rhs1 = _85_v2;
      let _rhs2 = (((_112_v25).contains(new _dafny.CodePoint('r'.codePointAt(0)))) ? ((_112_v25).get(new _dafny.CodePoint('r'.codePointAt(0)))) : (!((_83_v0) || (_83_v0))));
      let _rhs3 = (((_83_v0) ? (_108_v21.f30) : (_85_v2))).minus(new BigNumber((_dafny.Seq.Concat(_86_v3, _86_v3)).length));
      let _rhs4 = ((_83_v0) ? ((((_113_v26).contains(_108_v21.f30)) ? ((_113_v26).get(_108_v21.f30)) : (_107_v20))) : (_107_v20));
      let _lhs0 = _95_v10;
      let _lhs1 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_95_v10).length));
      let _lhs2 = globalState;
      let _lhs3 = _102_v17;
      let _lhs4 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_102_v17).length));
      let _lhs5 = _106_v19;
      let _lhs6 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_106_v19).length));
      _83_v0 = _rhs0;
      _lhs0[_lhs1] = _rhs1;
      _lhs2.f13 = _rhs2;
      _lhs3[_lhs4] = _rhs3;
      _lhs5[_lhs6] = _rhs4;
      (globalState).f24 = _105_v18;
      if (((_83_v0) ? ((_84_v1)[_module.__default.safeIndex((_102_v17)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_102_v17).length))], new BigNumber((_84_v1).length))]) : (!(_83_v0)))) {
        let _114_v27;
        let _nw9 = new _module.C0();
        _nw9.__ctor(!(false));
        _114_v27 = _nw9;
        let _rhs5 = _114_v27;
        let _rhs6 = _91_v6;
        let _rhs7 = _85_v2;
        let _lhs7 = _108_v21;
        _114_v27 = _rhs5;
        _91_v6 = _rhs6;
        _lhs7.f30 = _rhs7;
        let _115_v28;
        let _nw10 = Array((new BigNumber(27)).toNumber());
        _115_v28 = _nw10;
        let _index6 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_115_v28).length));
        (_115_v28)[_index6] = _98_v13;
        let _index7 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_115_v28).length));
        (_115_v28)[_index7] = _98_v13;
        let _116_v29;
        _116_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(211),_91_v6);
        let _117_v30;
        _117_v30 = _dafny.Set.fromElements(_83_v0);
        let _118_v31;
        _118_v31 = _dafny.Seq.of(_105_v18, _105_v18, _module.__default.fm42((_95_v10)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_95_v10).length))], _117_v30, (_95_v10)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_95_v10).length))], globalState));
        let _119_v32;
        _119_v32 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_108_v21.f30),(_114_v27).f32);
        let _index8 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_95_v10).length));
        let _index9 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_102_v17).length));
        let _rhs8 = _116_v29;
        let _rhs9 = ((_105_v18).Difference(_105_v18)).IsSubsetOf((_105_v18).Union((_118_v31)[_module.__default.safeIndex((_95_v10)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_95_v10).length))], new BigNumber((_118_v31).length))]));
        let _rhs10 = _module.__default.safeModuloInt(new BigNumber((_110_v23).length), _85_v2);
        let _rhs11 = ((!((_84_v1)[_module.__default.safeIndex(new BigNumber((_94_v9).cardinality()), new BigNumber((_84_v1).length))])) ? ((_dafny.ZERO).minus(_108_v21.f30)) : (((_83_v0) ? (new BigNumber((_dafny.Seq.of((((_119_v32).contains((_102_v17)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_102_v17).length))])) ? ((_119_v32).get((_102_v17)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_102_v17).length))])) : (_83_v0)), false)).length)) : (_108_v21.f30))));
        let _rhs12 = (_dafny.ZERO).minus((_95_v10)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_95_v10).length))]);
        let _lhs8 = globalState;
        let _lhs9 = globalState;
        let _lhs10 = _95_v10;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_95_v10).length));
        let _lhs12 = _102_v17;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_102_v17).length));
        _116_v29 = _rhs8;
        _lhs8.f22 = _rhs9;
        _lhs9.f10 = _rhs10;
        _lhs10[_lhs11] = _rhs11;
        _lhs12[_lhs13] = _rhs12;
        let _120_v33;
        _120_v33 = _dafny.Map.Empty.slice().updateUnsafe((_114_v27).f32,_108_v21.f30);
        let _121_v34;
        _121_v34 = _dafny.MultiSet.fromElements(_120_v33);
        let _122_v35;
        _122_v35 = _dafny.Seq.of(_117_v30, _117_v30);
        let _123_v36;
        let _nw11 = Array((new BigNumber(25)).toNumber());
        _nw11[(_dafny.ZERO).toNumber()] = (_121_v34).Intersect(_dafny.MultiSet.fromElements(_120_v33));
        _nw11[(_dafny.ONE).toNumber()] = _module.__default.fm45((_dafny.ZERO).minus((_95_v10)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_95_v10).length))]), globalState);
        _nw11[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(_120_v33, _120_v33, _120_v33, _120_v33, _120_v33);
        _nw11[(new BigNumber(3)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(4)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(5)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(6)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(7)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(8)).toNumber()] = (_121_v34).update(_120_v33, _module.__default.abs((_98_v13).fm2(globalState)));
        _nw11[(new BigNumber(9)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.FromArray(_module.__default.fm46((_102_v17)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_102_v17).length))], _85_v2, _108_v21.f30, globalState));
        _nw11[(new BigNumber(11)).toNumber()] = (_121_v34).Intersect(_121_v34);
        _nw11[(new BigNumber(12)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(13)).toNumber()] = (_121_v34).Intersect(_121_v34);
        _nw11[(new BigNumber(14)).toNumber()] = _module.__default.fm45((_108_v21).fm3((_102_v17)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_102_v17).length))], (_114_v27).f32, new BigNumber(-344), _module.D0.create_DC2((_102_v17)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_102_v17).length))], true), globalState), globalState);
        _nw11[(new BigNumber(15)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(16)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(17)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(18)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(19)).toNumber()] = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_114_v27).f32,new BigNumber((_122_v35).length)));
        _nw11[(new BigNumber(20)).toNumber()] = (_121_v34).Intersect(_121_v34);
        _nw11[(new BigNumber(21)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(22)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(23)).toNumber()] = _121_v34;
        _nw11[(new BigNumber(24)).toNumber()] = ((_121_v34).update(_120_v33, _module.__default.abs(_108_v21.f30))).Union(_121_v34);
        _123_v36 = _nw11;
        let _index10 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_123_v36).length));
        (_123_v36)[_index10] = _121_v34;
        let _index11 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_123_v36).length));
        (_123_v36)[_index11] = _121_v34;
        (globalState).f19 = _101_v16;
      } else {
        (globalState).f7 = _83_v0;
        (globalState).f13 = _83_v0;
        let _124_v37;
        let _nw12 = new _module.C1();
        _nw12.__ctor();
        _124_v37 = _nw12;
        _124_v37 = _124_v37;
        _108_v21 = _108_v21;
        (_108_v21).f30 = new BigNumber(667);
      }
      _84_v1 = _84_v1;
      let _125_v38;
      _125_v38 = _module.D14.create_DC34(true);
      let _126_v39;
      let _nw13 = Array((new BigNumber(9)).toNumber());
      _nw13[(_dafny.ZERO).toNumber()] = _83_v0;
      _nw13[(_dafny.ONE).toNumber()] = _83_v0;
      _nw13[(new BigNumber(2)).toNumber()] = _dafny.areEqual(_84_v1, _84_v1);
      _nw13[(new BigNumber(3)).toNumber()] = _83_v0;
      _nw13[(new BigNumber(4)).toNumber()] = (_83_v0) === (_83_v0);
      _nw13[(new BigNumber(5)).toNumber()] = _83_v0;
      _nw13[(new BigNumber(6)).toNumber()] = _83_v0;
      _nw13[(new BigNumber(7)).toNumber()] = _83_v0;
      _nw13[(new BigNumber(8)).toNumber()] = false;
      _126_v39 = _nw13;
      let _index12 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length));
      (_126_v39)[_index12] = _83_v0;
      let _127_v40;
      _127_v40 = new _dafny.CodePoint('v'.codePointAt(0));
      let _index13 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length));
      let _rhs13 = _module.D14.create_DC34(_83_v0);
      let _rhs14 = _85_v2;
      let _rhs15 = _83_v0;
      let _rhs16 = _85_v2;
      let _rhs17 = _module.__default.fm10(_dafny.Map.Empty.slice().updateUnsafe(_85_v2,_85_v2), _127_v40, _83_v0, globalState);
      let _lhs14 = globalState;
      let _lhs15 = _126_v39;
      let _lhs16 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length));
      _125_v38 = _rhs13;
      _lhs14.f10 = _rhs14;
      _lhs15[_lhs16] = _rhs15;
      _85_v2 = _rhs16;
      _86_v3 = _rhs17;
      (globalState).f10 = new BigNumber((_84_v1).length);
      (globalState).f7 = (_126_v39)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length))];
      let _128_v41;
      _128_v41 = _dafny.Map.Empty.slice().updateUnsafe((_126_v39)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length))],(_126_v39)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length))]);
      let _129_v42;
      _129_v42 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_85_v2),_85_v2);
      let _130_v43;
      _130_v43 = _dafny.Seq.of(new BigNumber((_128_v41).length), new BigNumber(81), new BigNumber((_129_v42).length));
      let _hi0 = new BigNumber((_130_v43).length);
      for (let _131_i1 = _85_v2; _131_i1.isLessThan(_hi0); _131_i1 = _131_i1.plus(_dafny.ONE)) {
        let _132_v44;
        _132_v44 = _module.D4.create_DC14(_83_v0);
        let _133_v45;
        _133_v45 = _dafny.Seq.of(_132_v44, _132_v44, _132_v44);
        _133_v45 = _133_v45;
        let _134_v46;
        _134_v46 = _dafny.MultiSet.fromElements(_83_v0, _83_v0);
        (globalState).f10 = ((((_134_v46).contains(_83_v0)) ? ((_134_v46).get(_83_v0)) : (_85_v2))).plus(_85_v2);
        let _135_v47;
        _135_v47 = _dafny.Set.fromElements(_131_i1, _131_i1, _85_v2);
        let _rhs18 = _87_v4;
        let _rhs19 = _126_v39;
        let _rhs20 = (_128_v41).equals(_128_v41);
        let _rhs21 = ((!(_135_v47).contains((_dafny.ZERO).minus(_131_i1))) ? ((_126_v39)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length))]) : ((_126_v39)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length))]));
        let _lhs17 = globalState;
        let _lhs18 = globalState;
        _87_v4 = _rhs18;
        _126_v39 = _rhs19;
        _lhs17.f13 = _rhs20;
        _lhs18.f7 = _rhs21;
        let _136_v48;
        _136_v48 = _module.D7.create_DC19();
        let _source4 = _136_v48;
        if (_source4.is_DC19) {
          _126_v39 = _126_v39;
          (globalState).f13 = true;
          let _137_v49;
          let _nw14 = Array((new BigNumber(7)).toNumber());
          _nw14[(_dafny.ZERO).toNumber()] = _86_v3;
          _nw14[(_dafny.ONE).toNumber()] = _86_v3;
          _nw14[(new BigNumber(2)).toNumber()] = _86_v3;
          _nw14[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_127_v40), _86_v3);
          _nw14[(new BigNumber(4)).toNumber()] = _86_v3;
          _nw14[(new BigNumber(5)).toNumber()] = _86_v3;
          _nw14[(new BigNumber(6)).toNumber()] = _86_v3;
          _137_v49 = _nw14;
          let _index14 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_137_v49).length));
          (_137_v49)[_index14] = _86_v3;
          let _138_v50;
          let _nw15 = new _module.C0();
          _nw15.__ctor(false);
          _138_v50 = _nw15;
          let _index15 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_137_v49).length));
          let _rhs22 = _86_v3;
          let _rhs23 = _138_v50;
          let _lhs19 = _137_v49;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_137_v49).length));
          _lhs19[_lhs20] = _rhs22;
          _138_v50 = _rhs23;
          let _139_v51;
          _139_v51 = _dafny.Map.Empty.slice().updateUnsafe(_127_v40,_85_v2);
          let _140_v52;
          _140_v52 = _dafny.Seq.of(_139_v51, _139_v51, _139_v51, _139_v51, _dafny.Map.Empty.slice().updateUnsafe(_127_v40,_131_i1));
          let _141_v53;
          _141_v53 = _module.D7.create_DC18(_140_v52);
          _141_v53 = _141_v53;
        } else if (_source4.is_DC20) {
          let _142___mcc_h1 = (_source4).cf30;
          let _143___mcc_h2 = (_source4).cf31;
          let _144___mcc_h3 = (_source4).cf32;
          let _145___mcc_h4 = (_source4).cf33;
          let _146___mcc_h5 = (_source4).cf34;
          let _147_cf34 = _146___mcc_h5;
          let _148_cf33 = _145___mcc_h4;
          let _149_cf32 = _144___mcc_h3;
          let _150_cf31 = _143___mcc_h2;
          let _151_cf30 = _142___mcc_h1;
          (globalState).f22 = _83_v0;
          let _152_v54;
          _152_v54 = _dafny.Map.Empty.slice().updateUnsafe(_85_v2,!(true) || (_148_cf33));
          _152_v54 = (_152_v54).update(_module.__default.safeModuloInt(_147_cf34, _147_cf34), _83_v0);
          _130_v43 = _dafny.Seq.of(_module.__default.fm0(_85_v2, globalState), _147_cf34, _131_i1);
          let _153_v55;
          let _init0 = ((_154_v3) => function (_155_i2) {
            return _154_v3;
          })(_86_v3);
          let _nw16 = Array((new BigNumber(19)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw16.length); _i0_0++) {
            _nw16[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _153_v55 = _nw16;
          let _index16 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_153_v55).length));
          (_153_v55)[_index16] = _86_v3;
          let _index17 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_153_v55).length));
          (_153_v55)[_index17] = _dafny.Seq.Concat(_dafny.Seq.update(((_148_cf33) ? (_86_v3) : (_86_v3)), _module.__default.safeIndex(_131_i1, new BigNumber((((_148_cf33) ? (_86_v3) : (_86_v3))).length)), _127_v40), _dafny.Seq.UnicodeFromString("uyvafgoga"));
        } else {
          let _156___mcc_h6 = (_source4).cf29;
          let _157_cf29 = _156___mcc_h6;
          let _index18 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length));
          let _rhs24 = (_126_v39)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length))];
          let _rhs25 = (_126_v39)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length))];
          let _rhs26 = (_dafny.ZERO).minus(_85_v2);
          let _lhs21 = globalState;
          let _lhs22 = _126_v39;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length));
          let _lhs24 = globalState;
          _lhs21.f22 = _rhs24;
          _lhs22[_lhs23] = _rhs25;
          _lhs24.f10 = _rhs26;
          let _158_v56;
          _158_v56 = _dafny.Seq.of(_84_v1);
          let _159_v57;
          let _nw17 = Array((new BigNumber(2)).toNumber());
          _nw17[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((_158_v56)[_module.__default.safeIndex(_131_i1, new BigNumber((_158_v56).length))], _84_v1);
          _nw17[(_dafny.ONE).toNumber()] = _dafny.Seq.of((_126_v39)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length))], _83_v0, (_126_v39)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length))], (_126_v39)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length))], _83_v0);
          _159_v57 = _nw17;
          let _index19 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_159_v57).length));
          (_159_v57)[_index19] = _84_v1;
          let _index20 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_159_v57).length));
          (_159_v57)[_index20] = ((_83_v0) ? (_dafny.Seq.of((_126_v39)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length))], (_126_v39)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_126_v39).length))])) : (_84_v1));
          (globalState).f10 = (_dafny.ZERO).minus(_85_v2);
          _85_v2 = new BigNumber(889);
        }
      }
      return;
    }
    static Main(__noArgsParameter) {
      let _160_v0;
      let _nw18 = Array((new BigNumber(13)).toNumber()).fill(false);
      _160_v0 = _nw18;
      let _161_v1;
      _161_v1 = _dafny.MultiSet.fromElements(_160_v0, _160_v0);
      let _162_v2;
      _162_v2 = false;
      let _163_v3;
      _163_v3 = _dafny.Seq.of(_162_v2);
      let _164_v4;
      _164_v4 = _dafny.Seq.of((((_161_v1).contains(_160_v0)) ? ((_161_v1).get(_160_v0)) : (new BigNumber((_163_v3).length))));
      let _165_v5;
      _165_v5 = _dafny.Map.Empty.slice().updateUnsafe(_164_v4,_160_v0);
      let _166_v6;
      _166_v6 = new BigNumber(-551);
      let _167_v7;
      let _nw19 = Array((new BigNumber(2)).toNumber());
      _nw19[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_163_v3, _module.__default.safeIndex(_166_v6, new BigNumber((_163_v3).length)), _162_v2);
      _nw19[(_dafny.ONE).toNumber()] = _163_v3;
      _167_v7 = _nw19;
      let _168_v8;
      _168_v8 = _dafny.Seq.UnicodeFromString("wrohtasav");
      let _169_v9;
      _169_v9 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(177)), ((_170_v6) => function (_171_i0) {
        return _170_v6;
      })(_166_v6))).length));
      let _172_v10;
      _172_v10 = _dafny.Map.Empty.slice().updateUnsafe((_168_v8)[_module.__default.safeIndex(_166_v6, new BigNumber((_168_v8).length))],_169_v9);
      let _173_v11;
      let _nw20 = Array((new BigNumber(4)).toNumber());
      _nw20[(_dafny.ZERO).toNumber()] = (((_172_v10).contains(new _dafny.CodePoint('n'.codePointAt(0)))) ? ((_172_v10).get(new _dafny.CodePoint('n'.codePointAt(0)))) : (_169_v9));
      _nw20[(_dafny.ONE).toNumber()] = _169_v9;
      _nw20[(new BigNumber(2)).toNumber()] = _169_v9;
      _nw20[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(_166_v6, new BigNumber(985));
      _173_v11 = _nw20;
      let _174_v12;
      let _nw21 = Array((new BigNumber(12)).toNumber()).fill([]);
      _174_v12 = _nw21;
      let _175_v13;
      _175_v13 = new _dafny.CodePoint('s'.codePointAt(0));
      let _176_v14;
      _176_v14 = _dafny.Map.Empty.slice().updateUnsafe(_175_v13,_162_v2);
      let _177_v15;
      _177_v15 = _dafny.Map.Empty.slice().updateUnsafe(_176_v14,_162_v2);
      let _178_v16;
      _178_v16 = _dafny.Map.Empty.slice().updateUnsafe(_160_v0,_162_v2);
      let _179_v17;
      _179_v17 = _dafny.Seq.of(_168_v8);
      let _180_v18;
      let _nw22 = Array((new BigNumber(18)).toNumber());
      _nw22[(_dafny.ZERO).toNumber()] = _168_v8;
      _nw22[(_dafny.ONE).toNumber()] = _168_v8;
      _nw22[(new BigNumber(2)).toNumber()] = _168_v8;
      _nw22[(new BigNumber(3)).toNumber()] = _168_v8;
      _nw22[(new BigNumber(4)).toNumber()] = _168_v8;
      _nw22[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_168_v8, _module.__default.safeIndex(_166_v6, new BigNumber((_168_v8).length)), _175_v13);
      _nw22[(new BigNumber(6)).toNumber()] = _168_v8;
      _nw22[(new BigNumber(7)).toNumber()] = (_179_v17)[_module.__default.safeIndex(_166_v6, new BigNumber((_179_v17).length))];
      _nw22[(new BigNumber(8)).toNumber()] = _168_v8;
      _nw22[(new BigNumber(9)).toNumber()] = _168_v8;
      _nw22[(new BigNumber(10)).toNumber()] = (_module.D0.create_DC0(_168_v8)).dtor_cf0;
      _nw22[(new BigNumber(11)).toNumber()] = _168_v8;
      _nw22[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("vfacyd"), _module.__default.safeIndex(_166_v6, new BigNumber((_dafny.Seq.UnicodeFromString("vfacyd")).length)), (_168_v8)[_module.__default.safeIndex(new BigNumber((_163_v3).length), new BigNumber((_168_v8).length))]);
      _nw22[(new BigNumber(13)).toNumber()] = _168_v8;
      _nw22[(new BigNumber(14)).toNumber()] = _168_v8;
      _nw22[(new BigNumber(15)).toNumber()] = _168_v8;
      _nw22[(new BigNumber(16)).toNumber()] = _168_v8;
      _nw22[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_168_v8, _module.__default.safeIndex(_166_v6, new BigNumber((_168_v8).length)), _175_v13);
      _180_v18 = _nw22;
      let _181_v19;
      _181_v19 = _dafny.Set.fromElements(new BigNumber((_163_v3).length));
      let _182_globalState;
      let _nw23 = new _module.GlobalState();
      _nw23.__ctor((_165_v5).Merge(_165_v5), _167_v7, true, new _dafny.CodePoint('h'.codePointAt(0)), false, _173_v11, _174_v12, false, _169_v9, new BigNumber(974), new BigNumber(-767), _177_v15, _178_v16, true, true, true, false, new BigNumber(811), true, new _dafny.CodePoint('d'.codePointAt(0)), true, _180_v18, true, true, _181_v19, _dafny.Seq.Concat(_163_v3, _163_v3), true, new BigNumber(80), new _dafny.CodePoint('x'.codePointAt(0)));
      _182_globalState = _nw23;
      let _183_v20;
      _183_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_164_v4).length),false);
      _183_v20 = _183_v20;
      let _184_v21;
      _184_v21 = _module.D0.create_DC1(_module.__default.fm0(new BigNumber((_dafny.MultiSet.FromArray(_164_v4)).cardinality()), _182_globalState), _166_v6, _166_v6, _160_v0);
      let _pat_let_tv0 = _184_v21;
      let _source5 = function (_pat_let0_0) {
        return function (_185_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_186_dt__update_hcf7_h0) {
              return _module.D0.create_DC3(_186_dt__update_hcf7_h0);
            }(_pat_let1_0);
          }(_pat_let_tv0);
        }(_pat_let0_0);
      }(_module.D0.create_DC3(_184_v21));
      if (_source5.is_DC1) {
        let _187___mcc_h0 = (_source5).cf1;
        let _188___mcc_h1 = (_source5).cf2;
        let _189___mcc_h2 = (_source5).cf3;
        let _190___mcc_h3 = (_source5).cf4;
        let _191_cf4 = _190___mcc_h3;
        let _192_cf3 = _189___mcc_h2;
        let _193_cf2 = _188___mcc_h1;
        let _194_cf1 = _187___mcc_h0;
        (_182_globalState).f7 = (_192_cf3).isLessThanOrEqualTo(_192_cf3);
        let _195_v22;
        _195_v22 = _module.D0.create_DC2(_193_cf2, _162_v2);
        _195_v22 = _195_v22;
        let _196_v23;
        let _nw24 = new _module.C1();
        _nw24.__ctor();
        _196_v23 = _nw24;
        let _197_v24;
        let _198_v25;
        let _out0;
        let _out1;
        let _outcollector0 = (_196_v23).m1(_dafny.Seq.update(_164_v4, _module.__default.safeIndex(_194_cf1, new BigNumber((_164_v4).length)), _194_cf1), ((_162_v2) ? (_194_cf1) : (_194_cf1)), _162_v2, _182_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _197_v24 = _out0;
        _198_v25 = _out1;
      } else if (_source5.is_DC2) {
        let _199___mcc_h4 = (_source5).cf5;
        let _200___mcc_h5 = (_source5).cf6;
        let _201_cf6 = _200___mcc_h5;
        let _202_cf5 = _199___mcc_h4;
        let _203_v26;
        _203_v26 = _dafny.Map.Empty.slice().updateUnsafe(_162_v2,_201_cf6);
        _183_v20 = (_183_v20).update(_module.__default.safeModuloInt(new BigNumber((_183_v20).length), _202_cf5), !((new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(!((((_203_v26).contains(_162_v2)) ? ((_203_v26).get(_162_v2)) : (_162_v2))), _201_cf6), _module.__default.safeIndex(_202_cf5, new BigNumber((_dafny.Seq.of(!((((_203_v26).contains(_162_v2)) ? ((_203_v26).get(_162_v2)) : (_162_v2))), _201_cf6)).length)), _201_cf6), _module.__default.safeIndex(_202_cf5, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(!((((_203_v26).contains(_162_v2)) ? ((_203_v26).get(_162_v2)) : (_162_v2))), _201_cf6), _module.__default.safeIndex(_202_cf5, new BigNumber((_dafny.Seq.of(!((((_203_v26).contains(_162_v2)) ? ((_203_v26).get(_162_v2)) : (_162_v2))), _201_cf6)).length)), _201_cf6)).length)), _162_v2)).length)).isLessThanOrEqualTo(_202_cf5)));
        _module.__default.m8(_182_globalState);
        let _204_v27;
        let _init1 = function (_205_i1) {
          return (_205_i1).multipliedBy(new BigNumber(-936));
        };
        let _nw25 = Array((new BigNumber(12)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw25.length); _i0_1++) {
          _nw25[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _204_v27 = _nw25;
        let _index21 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_204_v27).length));
        (_204_v27)[_index21] = _166_v6;
        let _index22 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_204_v27).length));
        (_204_v27)[_index22] = _202_cf5;
        let _index23 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_160_v0).length));
        (_160_v0)[_index23] = (_163_v3)[_module.__default.safeIndex((_204_v27)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_204_v27).length))], new BigNumber((_163_v3).length))];
        let _index24 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_160_v0).length));
        (_160_v0)[_index24] = _201_cf6;
      } else if (_source5.is_DC0) {
        let _206___mcc_h6 = (_source5).cf0;
        let _207_cf0 = _206___mcc_h6;
        _module.__default.m8(_182_globalState);
        let _208_v28;
        let _nw26 = Array((new BigNumber(21)).toNumber());
        _nw26[(_dafny.ZERO).toNumber()] = _175_v13;
        _nw26[(_dafny.ONE).toNumber()] = _175_v13;
        _nw26[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('j'.codePointAt(0));
        _nw26[(new BigNumber(3)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(4)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(5)).toNumber()] = (_207_cf0)[_module.__default.safeIndex((_dafny.ZERO).minus(_166_v6), new BigNumber((_207_cf0).length))];
        _nw26[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('d'.codePointAt(0));
        _nw26[(new BigNumber(7)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(8)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(9)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(10)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
        _nw26[(new BigNumber(11)).toNumber()] = (_168_v8)[_module.__default.safeIndex(new BigNumber((_169_v9).cardinality()), new BigNumber((_168_v8).length))];
        _nw26[(new BigNumber(12)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(13)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(14)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(15)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(16)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(17)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(18)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(19)).toNumber()] = _175_v13;
        _nw26[(new BigNumber(20)).toNumber()] = _175_v13;
        _208_v28 = _nw26;
        let _index25 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_208_v28).length));
        (_208_v28)[_index25] = new _dafny.CodePoint('l'.codePointAt(0));
        let _209_v29;
        _209_v29 = _dafny.Map.Empty.slice().updateUnsafe(_160_v0,_175_v13);
        let _index26 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_208_v28).length));
        (_208_v28)[_index26] = ((_162_v2) ? (_175_v13) : ((((_209_v29).contains(_160_v0)) ? ((_209_v29).get(_160_v0)) : (_175_v13))));
        (_182_globalState).f10 = (_166_v6).multipliedBy((_dafny.ZERO).minus((_166_v6).plus(_166_v6)));
        let _210_v30;
        _210_v30 = _module.D0.create_DC3(((_162_v2) ? (_184_v21) : (_184_v21)));
        let _source6 = _210_v30;
        if (_source6.is_DC1) {
          let _211___mcc_h8 = (_source6).cf1;
          let _212___mcc_h9 = (_source6).cf2;
          let _213___mcc_h10 = (_source6).cf3;
          let _214___mcc_h11 = (_source6).cf4;
          let _215_cf4 = _214___mcc_h11;
          let _216_cf3 = _213___mcc_h10;
          let _217_cf2 = _212___mcc_h9;
          let _218_cf1 = _211___mcc_h8;
          let _index27 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_215_cf4).length));
          (_215_cf4)[_index27] = _module.__default.fm12(_162_v2, _162_v2, _182_globalState);
          let _index28 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_215_cf4).length));
          (_215_cf4)[_index28] = _162_v2;
          _216_cf3 = _216_cf3;
          _module.__default.m8(_182_globalState);
          _207_cf0 = _207_cf0;
        } else if (_source6.is_DC2) {
          let _219___mcc_h12 = (_source6).cf5;
          let _220___mcc_h13 = (_source6).cf6;
          let _221_cf6 = _220___mcc_h13;
          let _222_cf5 = _219___mcc_h12;
          let _223_v31;
          let _nw27 = Array((new BigNumber(16)).toNumber());
          _223_v31 = _nw27;
          let _224_v32;
          let _nw28 = Array((new BigNumber(25)).toNumber()).fill([]);
          _224_v32 = _nw28;
          let _225_v33;
          let _nw29 = new _module.C4();
          _nw29.__ctor(_224_v32);
          _225_v33 = _nw29;
          let _226_v34;
          _226_v34 = _dafny.Map.Empty.slice().updateUnsafe(_223_v31,_225_v33);
          _226_v34 = (_226_v34).update(_223_v31, _225_v33);
          let _227_v35;
          let _nw30 = Array((new BigNumber(4)).toNumber());
          _nw30[(_dafny.ZERO).toNumber()] = _222_cf5;
          _nw30[(_dafny.ONE).toNumber()] = _166_v6;
          _nw30[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_222_cf5);
          _nw30[(new BigNumber(3)).toNumber()] = _166_v6;
          _227_v35 = _nw30;
          _227_v35 = ((_221_cf6) ? (_227_v35) : (_227_v35));
          let _228_v36;
          _228_v36 = _208_v28;
          let _rhs27 = _227_v35;
          let _rhs28 = !(_166_v6).isEqualTo(_222_cf5);
          let _rhs29 = _208_v28;
          let _lhs25 = _182_globalState;
          _227_v35 = _rhs27;
          _lhs25.f7 = _rhs28;
          _228_v36 = _rhs29;
          let _index29 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_227_v35).length));
          (_227_v35)[_index29] = _222_cf5;
          let _index30 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_227_v35).length));
          (_227_v35)[_index30] = _222_cf5;
        } else if (_source6.is_DC0) {
          let _229___mcc_h14 = (_source6).cf0;
          let _230_cf0 = _229___mcc_h14;
          let _231_v37;
          _231_v37 = _dafny.Seq.of(_183_v20);
          _231_v37 = ((_162_v2) ? (_dafny.Seq.Concat(_231_v37, _dafny.Seq.of(_183_v20))) : (_dafny.Seq.of(_183_v20)));
          let _232_v38;
          let _nw31 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _232_v38 = _nw31;
          let _233_v39;
          _233_v39 = _dafny.Seq.of(_232_v38, _232_v38, _232_v38, _232_v38, _232_v38);
          let _index31 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_232_v38).length));
          (_232_v38)[_index31] = new BigNumber((_233_v39).length);
          let _index32 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_232_v38).length));
          (_232_v38)[_index32] = (_164_v4)[_module.__default.safeIndex((new BigNumber(640)).minus(new BigNumber((_181_v19).length)), new BigNumber((_164_v4).length))];
          let _index33 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_232_v38).length));
          (_232_v38)[_index33] = _166_v6;
          let _index34 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_160_v0).length));
          (_160_v0)[_index34] = _module.__default.fm12(false, false, _182_globalState);
          let _index35 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_160_v0).length));
          (_160_v0)[_index35] = !(!_dafny.Seq.contains(_179_v17, _dafny.Seq.UnicodeFromString("gr")));
        } else {
          let _234___mcc_h15 = (_source6).cf7;
          let _235_cf7 = _234___mcc_h15;
          _166_v6 = _166_v6;
          (_182_globalState).f10 = (_164_v4)[_module.__default.safeIndex((_dafny.ZERO).minus(_166_v6), new BigNumber((_164_v4).length))];
          _module.__default.m8(_182_globalState);
          let _236_v40;
          _236_v40 = _dafny.Map.Empty.slice().updateUnsafe(_162_v2,_208_v28);
          _236_v40 = (_236_v40).update(_162_v2, _208_v28);
        }
      } else {
        let _237___mcc_h7 = (_source5).cf7;
        let _238_cf7 = _237___mcc_h7;
        _167_v7 = _167_v7;
        if ((_169_v9).IsSubsetOf(_dafny.MultiSet.fromElements(_166_v6))) {
          let _239_v41;
          let _nw32 = new _module.C0();
          _nw32.__ctor(!(!(_module.__default.fm12(_162_v2, _162_v2, _182_globalState)) || (_162_v2)));
          _239_v41 = _nw32;
          (_182_globalState).f10 = _module.__default.fm0(_module.__default.safeModuloInt((_dafny.ZERO).minus(_166_v6), _166_v6), _182_globalState);
          (_182_globalState).f7 = (_239_v41).f32;
          let _240_v42;
          let _init2 = ((_241_v6) => function (_242_i2) {
            return (_242_i2).multipliedBy(_241_v6);
          })(_166_v6);
          let _nw33 = Array((new BigNumber(16)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw33.length); _i0_2++) {
            _nw33[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _240_v42 = _nw33;
          let _index36 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_240_v42).length));
          (_240_v42)[_index36] = _166_v6;
          let _243_v44;
          _243_v44 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_166_v6, _182_globalState),true), function () {
            let _coll15 = new _dafny.Map();
            for (const _compr_15 of _dafny.IntegerRange(new BigNumber(599), new BigNumber(-29))) {
              let _244_v43 = _compr_15;
              if (((new BigNumber(599)).isLessThanOrEqualTo(_244_v43)) && ((_244_v43).isLessThan(new BigNumber(-29)))) {
                _coll15.push([_module.__default.safeDivisionInt(_244_v43, new BigNumber((_168_v8).length)),_162_v2]);
              }
            }
            return _coll15;
          }());
          let _index37 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_240_v42).length));
          (_240_v42)[_index37] = (_module.__default.safeDivisionInt(_166_v6, _166_v6)).plus((_166_v6).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_243_v44).length))).length))));
          (_182_globalState).f22 = (_239_v41).f32;
        } else {
          let _245_v45;
          let _nw34 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _245_v45 = _nw34;
          _245_v45 = _245_v45;
          let _246_v46;
          let _nw35 = new _module.C1();
          _nw35.__ctor();
          _246_v46 = _nw35;
          let _247_v47;
          _247_v47 = _dafny.MultiSet.fromElements(_162_v2);
          let _248_v48;
          _248_v48 = _dafny.Map.Empty.slice().updateUnsafe(_247_v47,_246_v46);
          let _249_v49;
          _249_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_168_v8).length),_246_v46);
          let _250_v50;
          let _nw36 = Array((new BigNumber(26)).toNumber());
          _nw36[(_dafny.ZERO).toNumber()] = _246_v46;
          _nw36[(_dafny.ONE).toNumber()] = _246_v46;
          _nw36[(new BigNumber(2)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(3)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(4)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(5)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(6)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(7)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(8)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(9)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(10)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(11)).toNumber()] = (((_248_v48).contains(_247_v47)) ? ((_248_v48).get(_247_v47)) : (_246_v46));
          _nw36[(new BigNumber(12)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(13)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(14)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(15)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(16)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(17)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(18)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(19)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(20)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(21)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(22)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(23)).toNumber()] = _246_v46;
          _nw36[(new BigNumber(24)).toNumber()] = (((_249_v49).contains(new BigNumber((_dafny.Seq.of(_162_v2, _162_v2, _162_v2)).length))) ? ((_249_v49).get(new BigNumber((_dafny.Seq.of(_162_v2, _162_v2, _162_v2)).length))) : (_246_v46));
          _nw36[(new BigNumber(25)).toNumber()] = _246_v46;
          _250_v50 = _nw36;
          let _index38 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_250_v50).length));
          (_250_v50)[_index38] = _246_v46;
          let _index39 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_250_v50).length));
          (_250_v50)[_index39] = _246_v46;
          let _index40 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_250_v50).length));
          (_250_v50)[_index40] = ((_162_v2) ? (_246_v46) : (_246_v46));
          _module.__default.m8(_182_globalState);
          let _index41 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_160_v0).length));
          (_160_v0)[_index41] = _162_v2;
          let _index42 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_160_v0).length));
          (_160_v0)[_index42] = !(false);
        }
        let _251_v51;
        let _nw37 = Array((new BigNumber(6)).toNumber()).fill([]);
        _251_v51 = _nw37;
        let _252_v52;
        _252_v52 = _dafny.Map.Empty.slice().updateUnsafe(_166_v6,_251_v51);
        let _253_v53;
        let _nw38 = new _module.C4();
        _nw38.__ctor((((_252_v52).contains(_166_v6)) ? ((_252_v52).get(_166_v6)) : (_251_v51)));
        _253_v53 = _nw38;
        let _rhs30 = !(!(!_dafny.areEqual(_179_v17, _179_v17))) || (_162_v2);
        let _rhs31 = _253_v53;
        let _lhs26 = _182_globalState;
        _lhs26.f7 = _rhs30;
        _253_v53 = _rhs31;
        let _rhs32 = (_162_v2) === (_162_v2);
        let _rhs33 = _175_v13;
        let _lhs27 = _182_globalState;
        let _lhs28 = _182_globalState;
        _lhs27.f13 = _rhs32;
        _lhs28.f19 = _rhs33;
      }
      let _254_v54;
      _254_v54 = _dafny.Set.fromElements(_162_v2);
      let _source7 = _module.D8.create_DC22(new BigNumber((_254_v54).length), _162_v2, _175_v13);
      if (_source7.is_DC22) {
        let _255___mcc_h16 = (_source7).cf36;
        let _256___mcc_h17 = (_source7).cf37;
        let _257___mcc_h18 = (_source7).cf38;
        let _258_cf38 = _257___mcc_h18;
        let _259_cf37 = _256___mcc_h17;
        let _260_cf36 = _255___mcc_h16;
        let _261_v55;
        let _nw39 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _261_v55 = _nw39;
        let _262_v56;
        _262_v56 = _dafny.Seq.of(_261_v55);
        (_182_globalState).f22 = (new BigNumber((_262_v56).length)).isLessThan(_166_v6);
        _259_cf37 = _dafny.Seq.IsPrefixOf(((!(_259_cf37)) ? (_164_v4) : (_164_v4)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(251)), ((_263_v6) => function (_264_i3) {
          return _263_v6;
        })(_166_v6)), _164_v4));
        let _265_v57;
        let _nw40 = new _module.C1();
        _nw40.__ctor();
        _265_v57 = _nw40;
        (_182_globalState).f10 = _module.__default.fm0(_166_v6, _182_globalState);
      } else if (_source7.is_DC23) {
        let _266___mcc_h19 = (_source7).cf39;
        let _267_cf39 = _266___mcc_h19;
        let _268_v58;
        _268_v58 = _dafny.Seq.of(_160_v0);
        let _269_v59;
        let _nw41 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _269_v59 = _nw41;
        let _index43 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_269_v59).length));
        (_269_v59)[_index43] = _166_v6;
        let _270_v60;
        let _nw42 = Array((new BigNumber(25)).toNumber()).fill([]);
        _270_v60 = _nw42;
        let _271_v61;
        let _nw43 = new _module.C4();
        _nw43.__ctor(_270_v60);
        _271_v61 = _nw43;
        let _272_v62;
        _272_v62 = _dafny.Seq.of(_271_v61);
        let _273_v63;
        _273_v63 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_166_v6, (_dafny.ZERO).minus(_166_v6)),(_272_v62)[_module.__default.safeIndex(_166_v6, new BigNumber((_272_v62).length))]);
        let _274_v64;
        _274_v64 = _dafny.Map.Empty.slice().updateUnsafe(_162_v2,_162_v2);
        let _index44 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_269_v59).length));
        let _rhs34 = _268_v58;
        let _rhs35 = (_166_v6).multipliedBy((new BigNumber((_274_v64).length)).multipliedBy(_166_v6));
        let _rhs36 = (_273_v63).Merge((_273_v63).Merge(_273_v63));
        let _lhs29 = _269_v59;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_269_v59).length));
        _268_v58 = _rhs34;
        _lhs29[_lhs30] = _rhs35;
        _273_v63 = _rhs36;
        let _275_v65;
        _275_v65 = _dafny.Map.Empty.slice().updateUnsafe(_269_v59,_162_v2);
        let _276_v66;
        _276_v66 = _dafny.Map.Empty.slice().updateUnsafe(_162_v2,_175_v13);
        let _277_v67;
        let _nw44 = Array((new BigNumber(16)).toNumber());
        _nw44[(_dafny.ZERO).toNumber()] = _175_v13;
        _nw44[(_dafny.ONE).toNumber()] = _175_v13;
        _nw44[(new BigNumber(2)).toNumber()] = _175_v13;
        _nw44[(new BigNumber(3)).toNumber()] = _175_v13;
        _nw44[(new BigNumber(4)).toNumber()] = _175_v13;
        _nw44[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
        _nw44[(new BigNumber(6)).toNumber()] = _175_v13;
        _nw44[(new BigNumber(7)).toNumber()] = ((_267_cf39) ? (_175_v13) : (_175_v13));
        _nw44[(new BigNumber(8)).toNumber()] = _175_v13;
        _nw44[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
        _nw44[(new BigNumber(10)).toNumber()] = _175_v13;
        _nw44[(new BigNumber(11)).toNumber()] = _module.__default.fm11(new BigNumber(((_275_v65).update(_269_v59, _267_cf39)).length), _182_globalState);
        _nw44[(new BigNumber(12)).toNumber()] = (((_276_v66).contains(_267_cf39)) ? ((_276_v66).get(_267_cf39)) : (_175_v13));
        _nw44[(new BigNumber(13)).toNumber()] = _175_v13;
        _nw44[(new BigNumber(14)).toNumber()] = _175_v13;
        _nw44[(new BigNumber(15)).toNumber()] = _175_v13;
        _277_v67 = _nw44;
        let _index45 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_277_v67).length));
        (_277_v67)[_index45] = _175_v13;
        let _278_v68;
        _278_v68 = _module.D8.create_DC22(new BigNumber(245), _162_v2, _175_v13);
        let _index46 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_277_v67).length));
        (_277_v67)[_index46] = ((_267_cf39) ? (_175_v13) : ((_278_v68).dtor_cf38));
        let _279_v69;
        _279_v69 = _dafny.Map.Empty.slice().updateUnsafe((_269_v59)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_269_v59).length))],_164_v4);
        _279_v69 = (_279_v69).update(_166_v6, _164_v4);
        let _280_v70;
        let _nw45 = new _module.C3();
        _nw45.__ctor();
        _280_v70 = _nw45;
        let _281_v71;
        _281_v71 = _dafny.Seq.of(_280_v70);
        _166_v6 = (_dafny.ZERO).minus((_166_v6).multipliedBy((new BigNumber((_281_v71).length)).minus(_166_v6)));
      } else if (_source7.is_DC21) {
        let _282___mcc_h20 = (_source7).cf35;
        let _283_cf35 = _282___mcc_h20;
        let _index47 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_160_v0).length));
        (_160_v0)[_index47] = _162_v2;
        let _index48 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_160_v0).length));
        (_160_v0)[_index48] = !(_162_v2);
        _254_v54 = _254_v54;
        (_182_globalState).f10 = _module.__default.safeModuloInt(_166_v6, (_166_v6).minus(_166_v6));
        let _284_v72;
        _284_v72 = _dafny.Map.Empty.slice().updateUnsafe(_166_v6,_166_v6);
        let _285_v73;
        _285_v73 = _dafny.MultiSet.fromElements(_175_v13);
        (_182_globalState).f10 = (((_284_v72).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_285_v73).cardinality()))))) ? ((_284_v72).get((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_285_v73).cardinality()))))) : (_166_v6));
      } else {
        let _286___mcc_h21 = (_source7).cf40;
        let _287_cf40 = _286___mcc_h21;
        let _288_v74;
        let _nw46 = Array((new BigNumber(23)).toNumber());
        _nw46[(_dafny.ZERO).toNumber()] = _175_v13;
        _nw46[(_dafny.ONE).toNumber()] = _175_v13;
        _nw46[(new BigNumber(2)).toNumber()] = _module.__default.fm11(_166_v6, _182_globalState);
        _nw46[(new BigNumber(3)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(4)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(5)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(6)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(7)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(8)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(9)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(10)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(11)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(12)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(13)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(14)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(15)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(16)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(17)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(18)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(19)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(20)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(21)).toNumber()] = _175_v13;
        _nw46[(new BigNumber(22)).toNumber()] = _175_v13;
        _288_v74 = _nw46;
        let _index49 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_288_v74).length));
        (_288_v74)[_index49] = _175_v13;
        let _index50 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_288_v74).length));
        (_288_v74)[_index50] = new _dafny.CodePoint('n'.codePointAt(0));
        let _289_v75;
        _289_v75 = _module.D4.create_DC14(_162_v2);
        let _290_v76;
        _290_v76 = _dafny.Seq.of(_164_v4, _164_v4);
        let _291_v77;
        _291_v77 = _module.D0.create_DC2(new BigNumber((_290_v76).length), _162_v2);
        let _source8 = (((_289_v75).dtor_cf25) ? (_module.D0.create_DC2(new BigNumber(334), !(_162_v2))) : (_291_v77));
        if (_source8.is_DC1) {
          let _292___mcc_h22 = (_source8).cf1;
          let _293___mcc_h23 = (_source8).cf2;
          let _294___mcc_h24 = (_source8).cf3;
          let _295___mcc_h25 = (_source8).cf4;
          let _296_cf4 = _295___mcc_h25;
          let _297_cf3 = _294___mcc_h24;
          let _298_cf2 = _293___mcc_h23;
          let _299_cf1 = _292___mcc_h22;
          (_182_globalState).f13 = !((_module.__default.fm12(false, _162_v2, _182_globalState)) === (!(_162_v2)));
          let _index51 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_288_v74).length));
          (_288_v74)[_index51] = (_288_v74)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_288_v74).length))];
          _module.__default.m8(_182_globalState);
          _module.__default.m8(_182_globalState);
        } else if (_source8.is_DC2) {
          let _300___mcc_h26 = (_source8).cf5;
          let _301___mcc_h27 = (_source8).cf6;
          let _302_cf6 = _301___mcc_h27;
          let _303_cf5 = _300___mcc_h26;
          _module.__default.m8(_182_globalState);
          (_182_globalState).f7 = _162_v2;
          let _304_v78;
          let _nw47 = new _module.C1();
          _nw47.__ctor();
          _304_v78 = _nw47;
          _304_v78 = _304_v78;
          let _305_v79;
          let _nw48 = Array((new BigNumber(15)).toNumber()).fill(_module.D3.Default());
          _305_v79 = _nw48;
          let _306_v80;
          _306_v80 = _dafny.MultiSet.fromElements(_302_cf6);
          let _index52 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_305_v79).length));
          (_305_v79)[_index52] = _module.D3.create_DC10(_306_v80);
          let _307_v81;
          _307_v81 = _module.D3.create_DC10(_306_v80);
          let _index53 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_305_v79).length));
          (_305_v79)[_index53] = _307_v81;
        } else if (_source8.is_DC0) {
          let _308___mcc_h28 = (_source8).cf0;
          let _309_cf0 = _308___mcc_h28;
          let _310_v82;
          let _nw49 = new _module.C0();
          _nw49.__ctor(false);
          _310_v82 = _nw49;
          let _311_v83;
          _311_v83 = _dafny.Map.Empty.slice().updateUnsafe(_310_v82,_166_v6);
          _311_v83 = (_311_v83).update(_310_v82, _166_v6);
          let _312_v84;
          _312_v84 = _dafny.Map.Empty.slice().updateUnsafe(_162_v2,_162_v2);
          let _index54 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_160_v0).length));
          (_160_v0)[_index54] = (((_312_v84).contains(_module.__default.fm12((_310_v82).f32, (((_183_v20).contains(_166_v6)) ? ((_183_v20).get(_166_v6)) : (_162_v2)), _182_globalState))) ? ((_312_v84).get(_module.__default.fm12((_310_v82).f32, (((_183_v20).contains(_166_v6)) ? ((_183_v20).get(_166_v6)) : (_162_v2)), _182_globalState))) : (!(_162_v2)));
          let _index55 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_160_v0).length));
          (_160_v0)[_index55] = _162_v2;
          _169_v9 = _169_v9;
          (_182_globalState).f10 = _module.__default.safeDivisionInt(_166_v6, (_166_v6).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_166_v6)).length))));
        } else {
          let _313___mcc_h29 = (_source8).cf7;
          let _314_cf7 = _313___mcc_h29;
          let _315_v85;
          let _nw50 = Array((new BigNumber(3)).toNumber()).fill([]);
          _315_v85 = _nw50;
          let _index56 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_315_v85).length));
          (_315_v85)[_index56] = _160_v0;
          let _index57 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_288_v74).length));
          let _index58 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_315_v85).length));
          let _rhs37 = _module.__default.fm11(_166_v6, _182_globalState);
          let _rhs38 = _160_v0;
          let _rhs39 = !(_162_v2);
          let _lhs31 = _288_v74;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_288_v74).length));
          let _lhs33 = _315_v85;
          let _lhs34 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_315_v85).length));
          _lhs31[_lhs32] = _rhs37;
          _lhs33[_lhs34] = _rhs38;
          _162_v2 = _rhs39;
          _166_v6 = _166_v6;
          (_182_globalState).f10 = _166_v6;
          let _index59 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_173_v11).length));
          (_173_v11)[_index59] = _169_v9;
          let _index60 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_173_v11).length));
          let _rhs40 = _166_v6;
          let _rhs41 = (_dafny.ZERO).minus(_166_v6);
          let _rhs42 = _dafny.MultiSet.fromElements(_166_v6, _166_v6);
          let _lhs35 = _182_globalState;
          let _lhs36 = _173_v11;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_173_v11).length));
          _lhs35.f10 = _rhs40;
          _166_v6 = _rhs41;
          _lhs36[_lhs37] = _rhs42;
        }
        let _316_v86;
        _316_v86 = _dafny.Map.Empty.slice().updateUnsafe(_166_v6,_module.__default.safeDivisionInt(new BigNumber((_254_v54).length), new BigNumber((_168_v8).length)));
        let _317_v87;
        let _nw51 = new _module.C1();
        _nw51.__ctor();
        _317_v87 = _nw51;
        let _318_v88;
        _318_v88 = _dafny.Seq.of(_317_v87);
        _316_v86 = (_316_v86).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_318_v88, _module.__default.safeIndex((((_169_v9).contains(_166_v6)) ? ((_169_v9).get(_166_v6)) : (_166_v6)), new BigNumber((_318_v88).length)), _317_v87),_162_v2)).length), _166_v6);
        let _319_v89;
        let _nw52 = Array((new BigNumber(11)).toNumber()).fill([]);
        _319_v89 = _nw52;
        let _320_v90;
        let _nw53 = new _module.C5();
        _nw53.__ctor(_319_v89, new BigNumber(-624));
        _320_v90 = _nw53;
        _320_v90 = ((_162_v2) ? (_320_v90) : (_320_v90));
      }
      (_182_globalState).f13 = _162_v2;
      let _321_i4;
      _321_i4 = _dafny.ZERO;
      L0: {
        while (_162_v2) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_321_i4)) {
              break L0;
            }
            _321_i4 = (_321_i4).plus(_dafny.ONE);
            let _322_v91;
            _322_v91 = _dafny.Map.Empty.slice().updateUnsafe(_163_v3,_166_v6);
            let _323_v92;
            _323_v92 = _dafny.Map.Empty.slice().updateUnsafe(_166_v6,(_dafny.ZERO).minus(new BigNumber((_322_v91).length)));
            let _324_v93;
            let _nw54 = new _module.C0();
            _nw54.__ctor(_dafny.Seq.IsPrefixOf(_module.__default.fm10(_323_v92, _175_v13, true, _182_globalState), _168_v8));
            _324_v93 = _nw54;
            let _325_v94;
            let _init3 = ((_326_v6) => function (_327_i5) {
              return (_327_i5).minus(_326_v6);
            })(_166_v6);
            let _nw55 = Array((_dafny.ONE).toNumber());
            for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw55.length); _i0_3++) {
              _nw55[_i0_3] = _init3(new BigNumber(_i0_3));
            }
            _325_v94 = _nw55;
            let _328_v95;
            _328_v95 = _dafny.Map.Empty.slice().updateUnsafe(_325_v94,_166_v6);
            (_182_globalState).f22 = _module.__default.fm12(_162_v2, !(_328_v95).equals(_328_v95), _182_globalState);
            let _329_v96;
            let _nw56 = new _module.C3();
            _nw56.__ctor();
            _329_v96 = _nw56;
            _329_v96 = _329_v96;
            let _330_v97;
            let _nw57 = Array((new BigNumber(22)).toNumber()).fill([]);
            _330_v97 = _nw57;
            let _331_v98;
            let _nw58 = new _module.C4();
            _nw58.__ctor(_330_v97);
            _331_v98 = _nw58;
            let _index61 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_325_v94).length));
            (_325_v94)[_index61] = _166_v6;
            let _index62 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_325_v94).length));
            let _rhs43 = !_dafny.Seq.contains(_168_v8, _175_v13);
            let _rhs44 = _331_v98;
            let _rhs45 = _168_v8;
            let _rhs46 = _166_v6;
            let _lhs38 = _182_globalState;
            let _lhs39 = _325_v94;
            let _lhs40 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_325_v94).length));
            _lhs38.f7 = _rhs43;
            _331_v98 = _rhs44;
            _168_v8 = _rhs45;
            _lhs39[_lhs40] = _rhs46;
          }
        }
      }
      let _332_v100;
      _332_v100 = _dafny.Map.Empty.slice().updateUnsafe(_166_v6,function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_168_v8).Elements) {
          let _333_v99 = _compr_16;
          if (_dafny.Seq.contains(_168_v8, _333_v99)) {
            _coll16.push([_333_v99,_166_v6]);
          }
        }
        return _coll16;
      }());
      (_182_globalState).f10 = ((_module.__default.fm12(!(_162_v2), _162_v2, _182_globalState)) ? ((new BigNumber((_332_v100).length)).multipliedBy(_166_v6)) : (_module.__default.safeDivisionInt(new BigNumber((_168_v8).length), _166_v6)));
      let _334_v101;
      _334_v101 = _module.D12.create_DC28(_dafny.MultiSet.FromArray(_164_v4));
      (_182_globalState).f13 = ((_169_v9).Intersect((_334_v101).dtor_cf44)).IsProperSubsetOf(_169_v9);
      let _335_v102;
      let _nw59 = Array((new BigNumber(3)).toNumber());
      _nw59[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_162_v2), _163_v3)).length);
      _nw59[(_dafny.ONE).toNumber()] = _166_v6;
      _nw59[(new BigNumber(2)).toNumber()] = new BigNumber((_168_v8).length);
      _335_v102 = _nw59;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_335_v102).length))) {
        let _336_i6 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_336_i6)) && ((_336_i6).isLessThan(new BigNumber((_335_v102).length))))) {
          (_335_v102)[(_336_i6)] = (_336_i6).minus((_166_v6).minus(_166_v6));
        }
      }
      if (!(_dafny.areEqual(_dafny.Seq.Concat(_163_v3, (_module.D8.create_DC21(_163_v3)).dtor_cf35), _163_v3))) {
        let _337_v103;
        _337_v103 = _dafny.Map.Empty.slice().updateUnsafe(_175_v13,(_166_v6).plus(_166_v6));
        _337_v103 = (_337_v103).update(_175_v13, new BigNumber(526));
        let _338_v104;
        _338_v104 = _module.D3.create_DC10(_dafny.MultiSet.FromArray(_163_v3));
        let _pat_let_tv1 = _162_v2;
        let _pat_let_tv2 = _162_v2;
        let _pat_let_tv3 = _182_globalState;
        let _pat_let_tv4 = _163_v3;
        let _source9 = function (_pat_let2_0) {
          return function (_339_dt__update__tmp_h2) {
            return function (_pat_let3_0) {
              return function (_340_dt__update_hcf21_h0) {
                return _module.D3.create_DC10(_340_dt__update_hcf21_h0);
              }(_pat_let3_0);
            }((_dafny.MultiSet.fromElements(_module.__default.fm12(_pat_let_tv1, _pat_let_tv2, _pat_let_tv3))).Difference(_dafny.MultiSet.FromArray(_pat_let_tv4)));
          }(_pat_let2_0);
        }(_338_v104);
        if (_source9.is_DC11) {
          let _341___mcc_h30 = (_source9).cf22;
          let _342_cf22 = _341___mcc_h30;
          let _343_v105;
          let _nw60 = new _module.C3();
          _nw60.__ctor();
          _343_v105 = _nw60;
          let _344_v106;
          _344_v106 = _dafny.MultiSet.fromElements(_343_v105);
          (_182_globalState).f22 = !(((new BigNumber((_169_v9).cardinality())).multipliedBy(_342_cf22)).isLessThan((((_344_v106).contains(_343_v105)) ? ((_344_v106).get(_343_v105)) : ((_343_v105).fm8(_342_cf22, _166_v6, _342_cf22, _182_globalState)))));
          let _345_v107;
          _345_v107 = _module.D4.create_DC13(_175_v13);
          let _346_v108;
          _346_v108 = _module.D4.create_DC15(_345_v107);
          let _347_v109;
          _347_v109 = _module.D7.create_DC20(_dafny.Seq.of(_342_cf22), _160_v0, _346_v108, _162_v2, ((_162_v2) ? (_166_v6) : (_342_cf22)));
          _347_v109 = _347_v109;
          let _348_v110;
          let _nw61 = Array((new BigNumber(16)).toNumber()).fill([]);
          _348_v110 = _nw61;
          let _349_v111;
          let _nw62 = new _module.C4();
          _nw62.__ctor(_348_v110);
          _349_v111 = _nw62;
          let _350_v112;
          let _nw63 = new _module.C4();
          _nw63.__ctor((_349_v111).f31);
          _350_v112 = _nw63;
          let _351_v113;
          _351_v113 = _335_v102;
          let _352_v114;
          let _nw64 = Array((new BigNumber(12)).toNumber());
          _nw64[(_dafny.ZERO).toNumber()] = _335_v102;
          _nw64[(_dafny.ONE).toNumber()] = (_351_v113);
          _nw64[(new BigNumber(2)).toNumber()] = _335_v102;
          _nw64[(new BigNumber(3)).toNumber()] = _335_v102;
          _nw64[(new BigNumber(4)).toNumber()] = _335_v102;
          _nw64[(new BigNumber(5)).toNumber()] = _335_v102;
          _nw64[(new BigNumber(6)).toNumber()] = _335_v102;
          _nw64[(new BigNumber(7)).toNumber()] = _335_v102;
          _nw64[(new BigNumber(8)).toNumber()] = _335_v102;
          _nw64[(new BigNumber(9)).toNumber()] = _335_v102;
          _nw64[(new BigNumber(10)).toNumber()] = _335_v102;
          _nw64[(new BigNumber(11)).toNumber()] = _335_v102;
          _352_v114 = _nw64;
          let _nw65 = new _module.C5();
          _nw65.__ctor(_352_v114, ((_dafny.ZERO).minus(_166_v6)).plus(_166_v6));
          _350_v112 = _nw65;
        } else if (_source9.is_DC12) {
          let _353___mcc_h31 = (_source9).cf23;
          let _354_cf23 = _353___mcc_h31;
          (_182_globalState).f10 = _166_v6;
          let _355_v115;
          _355_v115 = _dafny.Map.Empty.slice().updateUnsafe(_175_v13,_254_v54);
          let _356_v116;
          let _nw66 = Array((new BigNumber(13)).toNumber());
          _nw66[(_dafny.ZERO).toNumber()] = _254_v54;
          _nw66[(_dafny.ONE).toNumber()] = (_dafny.Set.fromElements(_354_cf23)).Union(_254_v54);
          _nw66[(new BigNumber(2)).toNumber()] = _254_v54;
          _nw66[(new BigNumber(3)).toNumber()] = _254_v54;
          _nw66[(new BigNumber(4)).toNumber()] = _254_v54;
          _nw66[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(_162_v2, _354_cf23);
          _nw66[(new BigNumber(6)).toNumber()] = _254_v54;
          _nw66[(new BigNumber(7)).toNumber()] = ((_354_cf23) ? ((((_355_v115).contains(_175_v13)) ? ((_355_v115).get(_175_v13)) : (_254_v54))) : (_dafny.Set.fromElements(_162_v2, _354_cf23)));
          _nw66[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(_354_cf23, (_163_v3)[_module.__default.safeIndex(new BigNumber(-580), new BigNumber((_163_v3).length))], _354_cf23, _module.__default.fm12(!(_354_cf23), false, _182_globalState));
          _nw66[(new BigNumber(9)).toNumber()] = _254_v54;
          _nw66[(new BigNumber(10)).toNumber()] = _254_v54;
          _nw66[(new BigNumber(11)).toNumber()] = _254_v54;
          _nw66[(new BigNumber(12)).toNumber()] = _254_v54;
          _356_v116 = _nw66;
          let _357_v117;
          _357_v117 = _dafny.Map.Empty.slice().updateUnsafe(_162_v2,!(true));
          let _358_v118;
          _358_v118 = (_dafny.Map.Empty.slice().updateUnsafe(_162_v2,_354_cf23)).update(false, _354_cf23);
          let _359_v119;
          _359_v119 = _module.D12.create_DC29(_dafny.Seq.Create(_module.__default.abs(new BigNumber(993)), function (_360_i7) {
  return new BigNumber(351);
}), _166_v6, _162_v2);
          let _361_v120;
          let _nw67 = Array((new BigNumber(26)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = ((_162_v2) ? (_357_v117) : (_358_v118));
          _nw67[(_dafny.ONE).toNumber()] = _358_v118;
          _nw67[(new BigNumber(2)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(3)).toNumber()] = _357_v117;
          _nw67[(new BigNumber(4)).toNumber()] = _357_v117;
          _nw67[(new BigNumber(5)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(6)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(7)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(8)).toNumber()] = _357_v117;
          _nw67[(new BigNumber(9)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(10)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(11)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(12)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_162_v2,_354_cf23);
          _nw67[(new BigNumber(14)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(15)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(16)).toNumber()] = (_357_v117).update(_162_v2, !(_354_cf23));
          _nw67[(new BigNumber(17)).toNumber()] = _357_v117;
          _nw67[(new BigNumber(18)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(19)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(20)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(21)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(22)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(23)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(24)).toNumber()] = _358_v118;
          _nw67[(new BigNumber(25)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_162_v2,(_359_v119).dtor_cf47);
          _361_v120 = _nw67;
          let _index63 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_361_v120).length));
          (_361_v120)[_index63] = (_357_v117).update(_354_cf23, _162_v2);
          let _index64 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_335_v102).length));
          (_335_v102)[_index64] = _module.__default.fm0(new BigNumber(504), _182_globalState);
          let _362_v121;
          _362_v121 = _dafny.Set.fromElements(_dafny.Seq.update(_168_v8, _module.__default.safeIndex((_dafny.ZERO).minus(_166_v6), new BigNumber((_168_v8).length)), (_168_v8)[_module.__default.safeIndex(_166_v6, new BigNumber((_168_v8).length))]));
          let _index65 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_361_v120).length));
          let _index66 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_335_v102).length));
          let _rhs47 = _356_v116;
          let _rhs48 = (_166_v6).plus(new BigNumber((_168_v8).length));
          let _rhs49 = _358_v118;
          let _rhs50 = _module.__default.safeDivisionInt(_166_v6, new BigNumber((_362_v121).length));
          let _rhs51 = new BigNumber(951);
          let _lhs41 = _361_v120;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_361_v120).length));
          let _lhs43 = _335_v102;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_335_v102).length));
          _356_v116 = _rhs47;
          _166_v6 = _rhs48;
          _lhs41[_lhs42] = _rhs49;
          _lhs43[_lhs44] = _rhs50;
          _166_v6 = _rhs51;
          (_182_globalState).f10 = (_335_v102)[_module.__default.safeIndex(new BigNumber(860), new BigNumber((_335_v102).length))];
          let _363_v122;
          _363_v122 = _dafny.Map.Empty.slice().updateUnsafe(_168_v8,(_183_v20).equals(_183_v20));
          let _364_v123;
          let _nw68 = Array((new BigNumber(18)).toNumber()).fill([]);
          _364_v123 = _nw68;
          let _index67 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_364_v123).length));
          (_364_v123)[_index67] = _180_v18;
          let _365_v124;
          let _nw69 = Array((new BigNumber(16)).toNumber()).fill([]);
          _365_v124 = _nw69;
          let _366_v125;
          let _nw70 = new _module.C4();
          _nw70.__ctor(_365_v124);
          _366_v125 = _nw70;
          let _367_v126;
          _367_v126 = _dafny.MultiSet.fromElements(_366_v125);
          let _368_v127;
          _368_v127 = _dafny.Seq.of(_367_v126);
          let _index68 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_160_v0).length));
          (_160_v0)[_index68] = _dafny.Seq.contains(_368_v127, _dafny.MultiSet.fromElements(_366_v125));
          let _369_v128;
          _369_v128 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(387),_168_v8);
          let _370_v129;
          _370_v129 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(((((_369_v128).contains(new BigNumber((_168_v8).length))) ? ((_369_v128).get(new BigNumber((_168_v8).length))) : (_168_v8))).length));
          let _371_v130;
          let _nw71 = Array((new BigNumber(27)).toNumber()).fill([]);
          _371_v130 = _nw71;
          let _372_v131;
          let _nw72 = new _module.C5();
          _nw72.__ctor(_371_v130, new BigNumber(-280));
          _372_v131 = _nw72;
          let _373_v132;
          _373_v132 = _dafny.Map.Empty.slice().updateUnsafe(_169_v9,_372_v131);
          let _index69 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_335_v102).length));
          let _index70 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_364_v123).length));
          let _index71 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_160_v0).length));
          let _rhs52 = (_166_v6).multipliedBy((((_370_v129).contains(_162_v2)) ? ((_370_v129).get(_162_v2)) : (_166_v6)));
          let _rhs53 = (_dafny.Map.Empty.slice().updateUnsafe(_168_v8,_354_cf23)).Merge(_363_v122);
          let _rhs54 = _180_v18;
          let _rhs55 = (new BigNumber(-786)).isLessThanOrEqualTo(_module.__default.fm0(new BigNumber((_373_v132).length), _182_globalState));
          let _rhs56 = _162_v2;
          let _lhs45 = _335_v102;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_335_v102).length));
          let _lhs47 = _364_v123;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_364_v123).length));
          let _lhs49 = _182_globalState;
          let _lhs50 = _160_v0;
          let _lhs51 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_160_v0).length));
          _lhs45[_lhs46] = _rhs52;
          _363_v122 = _rhs53;
          _lhs47[_lhs48] = _rhs54;
          _lhs49.f7 = _rhs55;
          _lhs50[_lhs51] = _rhs56;
        } else {
          let _374___mcc_h32 = (_source9).cf21;
          let _375_cf21 = _374___mcc_h32;
          let _rhs57 = _module.__default.fm12(_162_v2, _162_v2, _182_globalState);
          let _rhs58 = true;
          let _rhs59 = _dafny.Seq.of(_166_v6, _166_v6, _166_v6, _166_v6, _166_v6);
          let _lhs52 = _182_globalState;
          let _lhs53 = _182_globalState;
          _lhs52.f7 = _rhs57;
          _lhs53.f7 = _rhs58;
          _164_v4 = _rhs59;
          _168_v8 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qmpwft"), _168_v8), _168_v8), _module.__default.safeIndex(_166_v6, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qmpwft"), _168_v8), _168_v8)).length)), _175_v13);
          let _376_v133;
          let _nw73 = Array((new BigNumber(24)).toNumber());
          _376_v133 = _nw73;
          let _377_v134;
          let _nw74 = new _module.C3();
          _nw74.__ctor();
          _377_v134 = _nw74;
          let _index72 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_376_v133).length));
          (_376_v133)[_index72] = _377_v134;
          let _index73 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_376_v133).length));
          (_376_v133)[_index73] = _377_v134;
          let _index74 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_160_v0).length));
          (_160_v0)[_index74] = _module.__default.fm12(!(_162_v2), _162_v2, _182_globalState);
          let _index75 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_160_v0).length));
          let _rhs60 = _166_v6;
          let _rhs61 = _module.__default.fm12(_162_v2, _162_v2, _182_globalState);
          let _rhs62 = ((true) ? ((_377_v134).fm8(_166_v6, new BigNumber(184), _166_v6, _182_globalState)) : ((_377_v134).fm8(new BigNumber((_181_v19).length), _166_v6, new BigNumber((_163_v3).length), _182_globalState)));
          let _lhs54 = _160_v0;
          let _lhs55 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_160_v0).length));
          _166_v6 = _rhs60;
          _lhs54[_lhs55] = _rhs61;
          _166_v6 = _rhs62;
        }
        let _378_v135;
        let _nw75 = Array((new BigNumber(5)).toNumber());
        _nw75[(_dafny.ZERO).toNumber()] = _335_v102;
        _nw75[(_dafny.ONE).toNumber()] = _335_v102;
        _nw75[(new BigNumber(2)).toNumber()] = _335_v102;
        _nw75[(new BigNumber(3)).toNumber()] = _335_v102;
        _nw75[(new BigNumber(4)).toNumber()] = _335_v102;
        _378_v135 = _nw75;
        let _379_v136;
        _379_v136 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_164_v4),_166_v6);
        let _380_v137;
        _380_v137 = _dafny.Map.Empty.slice().updateUnsafe(_379_v136,_dafny.Set.fromElements(_162_v2));
        let _381_v139;
        _381_v139 = _dafny.Map.Empty.slice().updateUnsafe((((_380_v137).contains(_379_v136)) ? ((_380_v137).get(_379_v136)) : (_254_v54)),(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of (_181_v19).Elements) {
            let _382_v138 = _compr_17;
            if ((_181_v19).contains(_382_v138)) {
              _coll17.push([(_382_v138).plus(new BigNumber(-583)),_166_v6]);
            }
          }
          return _coll17;
        }())).cardinality())));
        let _383_v140;
        let _nw76 = new _module.C5();
        _nw76.__ctor(_378_v135, (_dafny.ZERO).minus((((_381_v139).contains(_254_v54)) ? ((_381_v139).get(_254_v54)) : (_166_v6))));
        _383_v140 = _nw76;
        let _384_v141;
        let _nw77 = Array((new BigNumber(2)).toNumber());
        _nw77[(_dafny.ZERO).toNumber()] = _383_v140;
        _nw77[(_dafny.ONE).toNumber()] = _383_v140;
        _384_v141 = _nw77;
        let _index76 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_384_v141).length));
        (_384_v141)[_index76] = _383_v140;
        let _index77 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_384_v141).length));
        let _rhs63 = _module.__default.safeModuloInt(_166_v6, (_dafny.ZERO).minus((_dafny.ZERO).minus(_166_v6)));
        let _rhs64 = _383_v140;
        let _rhs65 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((((_183_v20).contains(new BigNumber((_181_v19).length))) ? ((_183_v20).get(new BigNumber((_181_v19).length))) : (_162_v2)))).length));
        let _lhs56 = _182_globalState;
        let _lhs57 = _384_v141;
        let _lhs58 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_384_v141).length));
        let _lhs59 = _182_globalState;
        _lhs56.f10 = _rhs63;
        _lhs57[_lhs58] = _rhs64;
        _lhs59.f10 = _rhs65;
        let _385_v142;
        _385_v142 = _module.D12.create_DC31(_166_v6, !(false), _162_v2, new BigNumber((_169_v9).cardinality()));
        let _386_v143;
        _386_v143 = _module.D8.create_DC23((_385_v142).dtor_cf52);
        _386_v143 = _386_v143;
        let _387_v144;
        let _nw78 = new _module.C0();
        _nw78.__ctor(_162_v2);
        _387_v144 = _nw78;
      } else {
        _module.__default.m8(_182_globalState);
        (_182_globalState).f10 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.fm0(_166_v6, _182_globalState), _166_v6));
        (_182_globalState).f7 = !(!_dafny.areEqual(_168_v8, _dafny.Seq.UnicodeFromString("glmboqi")));
        let _388_v145;
        let _nw79 = new _module.C2();
        _nw79.__ctor((_module.D0.create_DC0(_168_v8)).dtor_cf0);
        _388_v145 = _nw79;
        (_388_v145).f33 = _388_v145.f33;
      }
      _335_v102 = _335_v102;
      if (_162_v2) {
        _166_v6 = _166_v6;
        let _389_v146;
        _389_v146 = _dafny.Map.Empty.slice().updateUnsafe(true,_166_v6);
        _389_v146 = (_389_v146).update(((_162_v2) ? (_162_v2) : (_162_v2)), _166_v6);
        let _nw80 = Array((new BigNumber(2)).toNumber());
        _nw80[(_dafny.ZERO).toNumber()] = !(!(_162_v2) || (_162_v2));
        _nw80[(_dafny.ONE).toNumber()] = _162_v2;
        _160_v0 = _nw80;
        let _390_v147;
        let _nw81 = Array((new BigNumber(18)).toNumber()).fill(_module.D8.Default());
        _390_v147 = _nw81;
        let _391_v148;
        _391_v148 = _module.D8.create_DC22(_166_v6, _162_v2, new _dafny.CodePoint('v'.codePointAt(0)));
        let _index78 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_390_v147).length));
        (_390_v147)[_index78] = _391_v148;
        let _index79 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_390_v147).length));
        (_390_v147)[_index79] = _391_v148;
        let _392_v149;
        let _nw82 = Array((new BigNumber(24)).toNumber()).fill([]);
        _392_v149 = _nw82;
        let _393_v150;
        _393_v150 = _392_v149;
        let _394_v151;
        let _nw83 = new _module.C4();
        _nw83.__ctor((_393_v150));
        _394_v151 = _nw83;
      } else {
        _168_v8 = _dafny.Seq.UnicodeFromString("faliulvks");
        let _395_v152;
        let _nw84 = Array((new BigNumber(22)).toNumber());
        _nw84[(_dafny.ZERO).toNumber()] = _335_v102;
        _nw84[(_dafny.ONE).toNumber()] = _335_v102;
        _nw84[(new BigNumber(2)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(3)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(4)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(5)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(6)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(7)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(8)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(9)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(10)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(11)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(12)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(13)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(14)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(15)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(16)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(17)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(18)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(19)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(20)).toNumber()] = _335_v102;
        _nw84[(new BigNumber(21)).toNumber()] = _335_v102;
        _395_v152 = _nw84;
        let _396_v153;
        let _nw85 = new _module.C5();
        _nw85.__ctor(_395_v152, _166_v6);
        _396_v153 = _nw85;
        let _397_v154;
        _397_v154 = _dafny.MultiSet.fromElements(_335_v102, _335_v102, _335_v102);
        let _rhs66 = _396_v153;
        let _rhs67 = (_254_v54).contains(_162_v2);
        let _rhs68 = _335_v102;
        let _rhs69 = _dafny.MultiSet.fromElements(_335_v102, _335_v102, _335_v102);
        let _rhs70 = _335_v102;
        let _lhs60 = _182_globalState;
        _396_v153 = _rhs66;
        _lhs60.f13 = _rhs67;
        _335_v102 = _rhs68;
        _397_v154 = _rhs69;
        _335_v102 = _rhs70;
        let _398_v155;
        let _nw86 = Array((new BigNumber(28)).toNumber()).fill([]);
        _398_v155 = _nw86;
        let _399_v156;
        let _nw87 = new _module.C4();
        _nw87.__ctor(_398_v155);
        _399_v156 = _nw87;
        let _400_v157;
        _400_v157 = _dafny.Set.fromElements((_module.D14.create_DC33(_399_v156)).dtor_cf56, _399_v156, _399_v156);
        let _401_v158;
        _401_v158 = _dafny.Map.Empty.slice().updateUnsafe(_400_v157,(_179_v17)[_module.__default.safeIndex(_396_v153.f30, new BigNumber((_179_v17).length))]);
        _401_v158 = (_401_v158).update(_400_v157, _dafny.Seq.Concat(_168_v8, _168_v8));
        let _402_v159;
        let _init4 = ((_403_v2, _404_v3, _405_v6) => function (_406_i8) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_403_v2,(_404_v3)[_module.__default.safeIndex(_405_v6, new BigNumber((_404_v3).length))])).Merge(_dafny.Map.Empty.slice().updateUnsafe(_403_v2,_403_v2));
        })(_162_v2, _163_v3, _166_v6);
        let _nw88 = Array((new BigNumber(24)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw88.length); _i0_4++) {
          _nw88[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _402_v159 = _nw88;
        let _407_v160;
        _407_v160 = _dafny.Map.Empty.slice().updateUnsafe(_162_v2,_162_v2);
        let _index80 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_402_v159).length));
        (_402_v159)[_index80] = (_407_v160).update(_162_v2, (_396_v153).fm4(_168_v8, _166_v6, _182_globalState));
        let _index81 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_402_v159).length));
        (_402_v159)[_index81] = ((_407_v160).Merge(_407_v160)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_162_v2));
        (_396_v153).f30 = _396_v153.f30;
      }
      let _408_v161;
      _408_v161 = _module.D4.create_DC15(_module.D4.create_DC14(_162_v2));
      (_182_globalState).f10 = (_module.D7.create_DC20(_dafny.Seq.update(_164_v4, _module.__default.safeIndex(_166_v6, new BigNumber((_164_v4).length)), _166_v6), _160_v0, _408_v161, _162_v2, _166_v6)).dtor_cf34;
      if (_162_v2) {
        let _409_v162;
        _409_v162 = _module.D7.create_DC19();
        let _rhs71 = _162_v2;
        let _rhs72 = _module.__default.fm12(_162_v2, _162_v2, _182_globalState);
        let _rhs73 = _module.D7.create_DC19();
        let _lhs61 = _182_globalState;
        let _lhs62 = _182_globalState;
        _lhs61.f13 = _rhs71;
        _lhs62.f13 = _rhs72;
        _409_v162 = _rhs73;
        (_182_globalState).f13 = _162_v2;
        if ((_163_v3)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_163_v3).length))]) {
          (_182_globalState).f10 = ((_dafny.ZERO).minus(_166_v6)).multipliedBy(_166_v6);
          let _410_v163;
          _410_v163 = _dafny.Map.Empty.slice().updateUnsafe(_164_v4,_162_v2);
          _410_v163 = _410_v163;
          _168_v8 = _168_v8;
          _164_v4 = _dafny.Seq.Concat(_164_v4, _dafny.Seq.of(_166_v6));
          let _411_v164;
          _411_v164 = _module.D4.create_DC13(_175_v13);
          let _pat_let_tv5 = _175_v13;
          let _412_v165;
          _412_v165 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let4_0) {
            return function (_413_dt__update__tmp_h5) {
              return function (_pat_let5_0) {
                return function (_414_dt__update_hcf24_h0) {
                  return _module.D4.create_DC13(_414_dt__update_hcf24_h0);
                }(_pat_let5_0);
              }(_pat_let_tv5);
            }(_pat_let4_0);
          }(_411_v164),((_162_v2) ? (_166_v6) : (_166_v6)));
          (_182_globalState).f10 = (((_412_v165).contains(_411_v164)) ? ((_412_v165).get(_411_v164)) : (_166_v6));
        } else {
          let _415_v166;
          _415_v166 = _dafny.Map.Empty.slice().updateUnsafe(!(_162_v2),_162_v2);
          let _416_v167;
          _416_v167 = _dafny.Map.Empty.slice().updateUnsafe((_415_v166).update(false, _162_v2),_166_v6);
          let _rhs74 = (new BigNumber(802)).multipliedBy(_166_v6);
          let _rhs75 = ((_162_v2) ? (_160_v0) : (_160_v0));
          let _rhs76 = _416_v167;
          let _rhs77 = !(!(_183_v20).equals(_dafny.Map.Empty.slice().updateUnsafe(_166_v6,_162_v2)));
          let _rhs78 = _162_v2;
          let _lhs63 = _182_globalState;
          let _lhs64 = _182_globalState;
          let _lhs65 = _182_globalState;
          _lhs63.f10 = _rhs74;
          _160_v0 = _rhs75;
          _416_v167 = _rhs76;
          _lhs64.f22 = _rhs77;
          _lhs65.f7 = _rhs78;
          let _index82 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_160_v0).length));
          (_160_v0)[_index82] = (((_415_v166).contains(!(_162_v2))) ? ((_415_v166).get(!(_162_v2))) : (_162_v2));
          let _index83 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_160_v0).length));
          (_160_v0)[_index83] = _162_v2;
          let _417_v168;
          let _init5 = ((_418_v0) => function (_419_i9) {
            return (_418_v0)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_418_v0).length))];
          })(_160_v0);
          let _nw89 = Array((new BigNumber(25)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw89.length); _i0_5++) {
            _nw89[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _417_v168 = _nw89;
          let _420_v169;
          _420_v169 = _dafny.Map.Empty.slice().updateUnsafe(_417_v168,_335_v102);
          _420_v169 = (_420_v169).update(_417_v168, _335_v102);
          _183_v20 = (_183_v20).update(new BigNumber((_168_v8).length), (_160_v0)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_160_v0).length))]);
          (_182_globalState).f10 = _166_v6;
        }
        let _421_v170;
        let _nw90 = new _module.C2();
        _nw90.__ctor(_168_v8);
        _421_v170 = _nw90;
        let _index84 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_335_v102).length));
        (_335_v102)[_index84] = new BigNumber((_168_v8).length);
        let _index85 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_335_v102).length));
        (_335_v102)[_index85] = new BigNumber((_dafny.Seq.update(_163_v3, _module.__default.safeIndex(new BigNumber(((_181_v19).Intersect(_181_v19)).length), new BigNumber((_163_v3).length)), !(_162_v2))).length);
      } else {
        let _422_v171;
        _422_v171 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-338),_335_v102);
        let _423_v172;
        _423_v172 = _335_v102;
        let _424_v173;
        let _nw91 = Array((new BigNumber(23)).toNumber());
        _nw91[(_dafny.ZERO).toNumber()] = _335_v102;
        _nw91[(_dafny.ONE).toNumber()] = _335_v102;
        _nw91[(new BigNumber(2)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(3)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(4)).toNumber()] = (((_422_v171).contains(_166_v6)) ? ((_422_v171).get(_166_v6)) : (_335_v102));
        _nw91[(new BigNumber(5)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(6)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(7)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(8)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(9)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(10)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(11)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(12)).toNumber()] = (_423_v172);
        _nw91[(new BigNumber(13)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(14)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(15)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(16)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(17)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(18)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(19)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(20)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(21)).toNumber()] = _335_v102;
        _nw91[(new BigNumber(22)).toNumber()] = _335_v102;
        _424_v173 = _nw91;
        let _425_v174;
        let _nw92 = new _module.C5();
        _nw92.__ctor(_424_v173, _166_v6);
        _425_v174 = _nw92;
        (_182_globalState).f13 = _162_v2;
        let _426_v175;
        _426_v175 = _dafny.Set.fromElements((_169_v9).Union(_dafny.MultiSet.fromElements(new BigNumber((_164_v4).length))), _169_v9, (_169_v9).Union(_169_v9));
        let _rhs79 = (_426_v175).Intersect(_426_v175);
        let _rhs80 = _162_v2;
        let _rhs81 = new BigNumber((_module.__default.fm43(((_162_v2) ? (_175_v13) : (_175_v13)), _dafny.Map.Empty.slice().updateUnsafe(_162_v2,_175_v13), _182_globalState)).length);
        let _lhs66 = _182_globalState;
        let _lhs67 = _182_globalState;
        _426_v175 = _rhs79;
        _lhs66.f7 = _rhs80;
        _lhs67.f10 = _rhs81;
        (_182_globalState).f10 = new BigNumber((_168_v8).length);
        let _427_v176;
        _427_v176 = _dafny.Map.Empty.slice().updateUnsafe(_175_v13,_160_v0);
        let _428_v177;
        _428_v177 = _dafny.Map.Empty.slice().updateUnsafe(_425_v174,_162_v2);
        let _429_v178;
        _429_v178 = _dafny.Seq.of(_160_v0, _160_v0);
        let _430_v179;
        _430_v179 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_428_v177).length),_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),(_429_v178)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_429_v178).length))]));
        let _431_v180;
        let _nw93 = Array((new BigNumber(16)).toNumber());
        _nw93[(_dafny.ZERO).toNumber()] = _427_v176;
        _nw93[(_dafny.ONE).toNumber()] = (_427_v176).Merge(_427_v176);
        _nw93[(new BigNumber(2)).toNumber()] = _427_v176;
        _nw93[(new BigNumber(3)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm11(_166_v6, _182_globalState),_160_v0)).Merge(_427_v176);
        _nw93[(new BigNumber(4)).toNumber()] = _427_v176;
        _nw93[(new BigNumber(5)).toNumber()] = (((_430_v179).contains(_425_v174.f30)) ? ((_430_v179).get(_425_v174.f30)) : (_427_v176));
        _nw93[(new BigNumber(6)).toNumber()] = (_427_v176).Merge(_427_v176);
        _nw93[(new BigNumber(7)).toNumber()] = (_427_v176).Merge(_427_v176);
        _nw93[(new BigNumber(8)).toNumber()] = _427_v176;
        _nw93[(new BigNumber(9)).toNumber()] = _427_v176;
        _nw93[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_175_v13,_160_v0);
        _nw93[(new BigNumber(11)).toNumber()] = (_427_v176).update(new _dafny.CodePoint('d'.codePointAt(0)), _160_v0);
        _nw93[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_175_v13,_160_v0);
        _nw93[(new BigNumber(13)).toNumber()] = _427_v176;
        _nw93[(new BigNumber(14)).toNumber()] = (_427_v176).update(_175_v13, _160_v0);
        _nw93[(new BigNumber(15)).toNumber()] = (_427_v176).Merge(_427_v176);
        _431_v180 = _nw93;
        let _index86 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_431_v180).length));
        (_431_v180)[_index86] = _427_v176;
        let _index87 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_431_v180).length));
        (_431_v180)[_index87] = ((_427_v176).Merge((_427_v176).update(_175_v13, _160_v0))).Merge(_427_v176);
      }
      _168_v8 = _dafny.Seq.Concat(_168_v8, _168_v8);
      _module.__default.m8(_182_globalState);
      _module.__default.m8(_182_globalState);
      process.stdout.write(_dafny.toString((_160_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_161_v1).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_162_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_163_v3, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_164_v4, _dafny.Seq.of(new BigNumber(-551), new BigNumber(-551), new BigNumber(-551), new BigNumber(-551), new BigNumber(-551)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_165_v5).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_166_v6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_167_v7)[_dafny.ZERO], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_167_v7)[_dafny.ONE], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_168_v8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v9).equals(_dafny.MultiSet.fromElements(new BigNumber(177)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),_dafny.MultiSet.fromElements(new BigNumber(177))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v11)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(new BigNumber(177)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v11)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(new BigNumber(177)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v11)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(new BigNumber(177)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v11)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(new BigNumber(-551), new BigNumber(985)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_175_v13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_176_v14).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_177_v15).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),false),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_178_v16).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_179_v17, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("wrohtasav")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(14)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(15)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(16)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_180_v18)[new BigNumber(17)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_181_v19).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_182_globalState).f0).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_182_globalState).f1)[_dafny.ZERO], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_182_globalState).f1)[_dafny.ONE], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_182_globalState).f5)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(new BigNumber(177)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_182_globalState).f5)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(new BigNumber(177)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_182_globalState).f5)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(new BigNumber(177)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_182_globalState).f5)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(new BigNumber(-551), new BigNumber(985)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_182_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_globalState).f8).equals(_dafny.MultiSet.fromElements(new BigNumber(177)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_182_globalState.f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState.f11).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),false),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_182_globalState).f12).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_182_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_182_globalState.f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(14)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(15)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(16)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_182_globalState).f21)[new BigNumber(17)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_182_globalState.f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState.f24).equals(_dafny.Set.fromElements(new BigNumber(667)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_182_globalState).f25, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_183_v20).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_184_v21).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_184_v21).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_184_v21).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v21).dtor_cf4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_184_v21).dtor_cf4)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v54).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_321_i4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_332_v100).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-551),_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new BigNumber(-551)).updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),new BigNumber(-551)).updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),new BigNumber(-551)).updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),new BigNumber(-551)).updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),new BigNumber(-551)).updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),new BigNumber(-551)).updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),new BigNumber(-551)).updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),new BigNumber(-551))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_334_v101).dtor_cf44).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v102)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v102)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v102)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_408_v161).dtor_cf26).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3, cf4) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC2(cf5, cf6) {
      let $dt = new D0(1);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC3(cf7) {
      let $dt = new D0(3);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + this.cf0.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, []);
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
    static create_DC5(cf9, cf10, cf11, cf12) {
      let $dt = new D1(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC6(cf13, cf14, cf15, cf16) {
      let $dt = new D1(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC4(cf8) {
      let $dt = new D1(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && this.cf11 === other.cf11 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(_dafny.ZERO, false, false, _dafny.Map.Empty);
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
    static create_DC8(cf18, cf19) {
      let $dt = new D2(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC7(cf17) {
      let $dt = new D2(1);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC9(cf20) {
      let $dt = new D2(2);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf18 === other.cf18 && this.cf19 === other.cf19;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(false, false);
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
    static create_DC11(cf22) {
      let $dt = new D3(0);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC12(cf23) {
      let $dt = new D3(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC10(cf21) {
      let $dt = new D3(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf23 === other.cf23;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC11(_dafny.ZERO);
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
    static create_DC14(cf25) {
      let $dt = new D4(0);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC13(cf24) {
      let $dt = new D4(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC15(cf26) {
      let $dt = new D4(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf25 === other.cf25;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC14(false);
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
    static create_DC16(cf27) {
      let $dt = new D5(0);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27);
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf28) {
      let $dt = new D6(0);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf28 === other.cf28;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
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
    static create_DC19() {
      let $dt = new D7(0);
      return $dt;
    }
    static create_DC20(cf30, cf31, cf32, cf33, cf34) {
      let $dt = new D7(1);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC18(cf29) {
      let $dt = new D7(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19";
      } else if (this.$tag === 1) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf29) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30) && this.cf31 === other.cf31 && _dafny.areEqual(this.cf32, other.cf32) && this.cf33 === other.cf33 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC19();
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
    static create_DC22(cf36, cf37, cf38) {
      let $dt = new D8(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC23(cf39) {
      let $dt = new D8(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC21(cf35) {
      let $dt = new D8(2);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC24(cf40) {
      let $dt = new D8(3);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get is_DC24() { return this.$tag === 3; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf39 === other.cf39;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC22(_dafny.ZERO, false, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC25(cf41) {
      let $dt = new D9(0);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf41, other.cf41);
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
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC26(cf42) {
      let $dt = new D10(0);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf42 === other.cf42;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
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
    static create_DC27(cf43) {
      let $dt = new D11(0);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf43 === other.cf43;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
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
    static create_DC29(cf45, cf46, cf47) {
      let $dt = new D12(0);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC30(cf48, cf49, cf50) {
      let $dt = new D12(1);
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC31(cf51, cf52, cf53, cf54) {
      let $dt = new D12(2);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC28(cf44) {
      let $dt = new D12(3);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
    get is_DC28() { return this.$tag === 3; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC29" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC30" + "(" + this.cf48.toVerbatimString(true) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf45, other.cf45) && _dafny.areEqual(this.cf46, other.cf46) && this.cf47 === other.cf47;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf48, other.cf48) && _dafny.areEqual(this.cf49, other.cf49) && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf51, other.cf51) && this.cf52 === other.cf52 && this.cf53 === other.cf53 && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC29(_dafny.Seq.of(), _dafny.ZERO, false);
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
    static create_DC32(cf55) {
      let $dt = new D13(0);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC32" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf55 === other.cf55;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
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
    static create_DC34(cf57) {
      let $dt = new D14(0);
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC35(cf58) {
      let $dt = new D14(1);
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC36(cf59, cf60, cf61, cf62, cf63) {
      let $dt = new D14(2);
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC33(cf56) {
      let $dt = new D14(3);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get is_DC33() { return this.$tag === 3; }
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
        return "D14.DC34" + "(" + _dafny.toString(this.cf57) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC35" + "(" + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 3) {
        return "D14.DC33" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf57 === other.cf57;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf58 === other.cf58;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf59, other.cf59) && this.cf60 === other.cf60 && this.cf61 === other.cf61 && this.cf62 === other.cf62 && _dafny.areEqual(this.cf63, other.cf63);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf56 === other.cf56;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC34(false);
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
    static create_DC38() {
      let $dt = new D15(0);
      return $dt;
    }
    static create_DC37(cf64) {
      let $dt = new D15(1);
      $dt.cf64 = cf64;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get dtor_cf64() { return this.cf64; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC38";
      } else if (this.$tag === 1) {
        return "D15.DC37" + "(" + _dafny.toString(this.cf64) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf64, other.cf64);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC38();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
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
      this.f6 = [];
      this.f7 = false;
      this.f10 = _dafny.ZERO;
      this.f11 = _dafny.Map.Empty;
      this.f13 = false;
      this.f19 = new _dafny.CodePoint('D'.codePointAt(0));
      this.f22 = false;
      this.f24 = _dafny.Set.Empty;
      this._f0 = _dafny.Map.Empty;
      this._f1 = [];
      this._f2 = false;
      this._f3 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f4 = false;
      this._f5 = [];
      this._f8 = _dafny.MultiSet.Empty;
      this._f9 = _dafny.ZERO;
      this._f12 = _dafny.Map.Empty;
      this._f14 = false;
      this._f15 = false;
      this._f16 = false;
      this._f17 = _dafny.ZERO;
      this._f18 = false;
      this._f20 = false;
      this._f21 = [];
      this._f23 = false;
      this._f25 = _dafny.Seq.of();
      this._f26 = false;
      this._f27 = _dafny.ZERO;
      this._f28 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27, f28) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this).f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this).f22 = f22;
      (_this)._f23 = f23;
      (_this).f24 = f24;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f12() {
      let _this = this;
      return _this._f12;
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
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    get f23() {
      let _this = this;
      return _this._f23;
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
    get f28() {
      let _this = this;
      return _this._f28;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f32 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f32) {
      let _this = this;
      (_this)._f32 = f32;
      return;
    }
    fm13(p0, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(function (_source10) {
        if (_source10.is_DC5) {
          let _432___mcc_h0 = (_source10).cf9;
          let _433___mcc_h1 = (_source10).cf10;
          let _434___mcc_h2 = (_source10).cf11;
          let _435___mcc_h3 = (_source10).cf12;
          let _436_cf12 = _435___mcc_h3;
          let _437_cf11 = _434___mcc_h2;
          let _438_cf10 = _433___mcc_h1;
          let _439_cf9 = _432___mcc_h0;
          return new BigNumber(732);
        } else if (_source10.is_DC6) {
          let _440___mcc_h4 = (_source10).cf13;
          let _441___mcc_h5 = (_source10).cf14;
          let _442___mcc_h6 = (_source10).cf15;
          let _443___mcc_h7 = (_source10).cf16;
          let _444_cf16 = _443___mcc_h7;
          let _445_cf15 = _442___mcc_h6;
          let _446_cf14 = _441___mcc_h5;
          let _447_cf13 = _440___mcc_h4;
          return (_dafny.ZERO).minus(_module.__default.safeModuloInt(_445_cf15, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_444_cf16,_447_cf13)).length)));
        } else {
          let _448___mcc_h8 = (_source10).cf8;
          let _449_cf8 = _448___mcc_h8;
          return _module.__default.safeDivisionInt(new BigNumber(632), new BigNumber((_dafny.Set.fromElements(new BigNumber(339))).length));
        }
      }(_module.D1.create_DC6((_module.D0.create_DC2(new BigNumber(891), (_this).f32)).dtor_cf5, _dafny.Seq.of(_dafny.Seq.of((_this).f32), _dafny.Seq.of(!((_this).f32), (_this).f32)), new BigNumber(124), (_this).f32)));
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(!(false)),!_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(946)), function (_450_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(true, false, false, true, true)).cardinality()))).length),false);
      }), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-535),!(true))));
    };
    fm2(globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(778)), function (_451_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length), _module.__default.safeDivisionInt(new BigNumber(-566), new BigNumber(-96)), (new BigNumber(964)).minus(new BigNumber((_dafny.Seq.UnicodeFromString("v")).length)), (new BigNumber(68)).multipliedBy(new BigNumber(-443)))).length));
    };
    fm21(p0, globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-402))).cardinality()), new BigNumber(-703), new BigNumber((_dafny.Seq.of(false)).length), new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality()))).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("g")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("thqpyy")).length),new BigNumber((_dafny.Seq.of(false, true)).length))).length)));
    };
    m0(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (p1) {
        let _452_v0;
        let _init6 = function (_453_i0) {
          return (_453_i0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ecpxkbg")).length));
        };
        let _nw94 = Array((new BigNumber(13)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw94.length); _i0_6++) {
          _nw94[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _452_v0 = _nw94;
        let _index88 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_452_v0).length));
        (_452_v0)[_index88] = p0;
        let _454_v1;
        _454_v1 = _dafny.MultiSet.fromElements(new BigNumber(-622), p2, p0, p2);
        let _index89 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_452_v0).length));
        (_452_v0)[_index89] = (((_454_v1).contains(p0)) ? ((_454_v1).get(p0)) : (p0));
        let _455_v2;
        let _nw95 = Array((new BigNumber(2)).toNumber());
        _nw95[(_dafny.ZERO).toNumber()] = (p1) === (p1);
        _nw95[(_dafny.ONE).toNumber()] = p1;
        _455_v2 = _nw95;
        let _index90 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_455_v2).length));
        (_455_v2)[_index90] = p3;
        let _456_v3;
        _456_v3 = _dafny.Seq.of(p0, p0);
        let _457_v4;
        _457_v4 = _dafny.Set.fromElements(new BigNumber((_456_v3).length), p2);
        let _458_v5;
        _458_v5 = _dafny.Seq.of(p0, p2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,false)).length), (_dafny.ZERO).minus(new BigNumber((_457_v4).length)));
        let _459_v6;
        _459_v6 = _dafny.Set.fromElements(p1, !(false), p1, p3);
        let _index91 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_452_v0).length));
        let _index92 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_452_v0).length));
        let _index93 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_455_v2).length));
        let _rhs82 = p0;
        let _rhs83 = _module.__default.fm0(p0, globalState);
        let _rhs84 = ((p1) ? (p0) : ((_456_v3)[_module.__default.safeIndex(new BigNumber((_459_v6).length), new BigNumber((_456_v3).length))]));
        let _rhs85 = p1;
        let _lhs68 = globalState;
        let _lhs69 = _452_v0;
        let _lhs70 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_452_v0).length));
        let _lhs71 = _452_v0;
        let _lhs72 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_452_v0).length));
        let _lhs73 = _455_v2;
        let _lhs74 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_455_v2).length));
        _lhs68.f10 = _rhs82;
        _lhs69[_lhs70] = _rhs83;
        _lhs71[_lhs72] = _rhs84;
        _lhs73[_lhs74] = _rhs85;
        let _460_v7;
        _460_v7 = _dafny.Map.Empty.slice().updateUnsafe((_455_v2)[_module.__default.safeIndex(new BigNumber(858), new BigNumber((_455_v2).length))],p3);
        let _461_v8;
        _461_v8 = _dafny.Seq.of((((_460_v7).contains(p3)) ? ((_460_v7).get(p3)) : (true)), !(p1) || (p3), p1);
        let _index94 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_452_v0).length));
        let _rhs86 = (_module.__default.safeDivisionInt(p0, _module.__default.fm0(p0, globalState))).plus(p0);
        let _rhs87 = _dafny.Seq.Concat(_461_v8, _dafny.Seq.Concat(_461_v8, _461_v8));
        let _rhs88 = ((false) ? (_452_v0) : (_452_v0));
        let _lhs75 = _452_v0;
        let _lhs76 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_452_v0).length));
        _lhs75[_lhs76] = _rhs86;
        _461_v8 = _rhs87;
        _452_v0 = _rhs88;
        if ((_this).fm21(p0, globalState)) {
          let _462_v9;
          _462_v9 = new _dafny.CodePoint('w'.codePointAt(0));
          (globalState).f19 = _462_v9;
          (globalState).f13 = true;
          let _463_v10;
          let _nw96 = new _module.C0();
          _nw96.__ctor(p3);
          _463_v10 = _nw96;
          (globalState).f22 = p3;
          let _index95 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_455_v2).length));
          (_455_v2)[_index95] = (_this).fm21((_dafny.ZERO).minus(p2), globalState);
        } else {
          _461_v8 = _dafny.Seq.Concat(_dafny.Seq.of(false, (_this).fm21((_dafny.ZERO).minus((_452_v0)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_452_v0).length))]), globalState), p1), _dafny.Seq.Concat(_module.__default.fm22((_455_v2)[_module.__default.safeIndex(new BigNumber(858), new BigNumber((_455_v2).length))], globalState), _461_v8));
          let _464_v11;
          _464_v11 = _dafny.Seq.UnicodeFromString("xfsnymi");
          let _465_v12;
          _465_v12 = _module.D0.create_DC0(_464_v11);
          (globalState).f10 = new BigNumber(((_465_v12).dtor_cf0).length);
          (globalState).f13 = p1;
          let _nw97 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _452_v0 = _nw97;
          (globalState).f7 = _dafny.Seq.contains(_dafny.Seq.Concat(_461_v8, _461_v8), p1);
        }
        let _index96 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_452_v0).length));
        (_452_v0)[_index96] = ((_dafny.ZERO).minus(p0)).minus(p2);
        (globalState).f13 = p1;
      } else {
        let _466_v13;
        _466_v13 = _dafny.MultiSet.fromElements(p1, p1, p3, p1, p1);
        let _467_v14;
        _467_v14 = _dafny.Seq.of(false, (_this).fm21(p0, globalState), p1, p1);
        let _468_v15;
        _468_v15 = _dafny.Set.fromElements(p1);
        let _469_v16;
        _469_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _470_v17;
        _470_v17 = _dafny.Seq.of(_467_v14);
        let _471_v18;
        _471_v18 = _module.D1.create_DC6(p2, _470_v17, p2, p3);
        let _472_v19;
        _472_v19 = _dafny.Seq.UnicodeFromString("sj");
        let _473_v20;
        let _nw98 = Array((new BigNumber(24)).toNumber());
        _nw98[(_dafny.ZERO).toNumber()] = _466_v13;
        _nw98[(_dafny.ONE).toNumber()] = _466_v13;
        _nw98[(new BigNumber(2)).toNumber()] = _466_v13;
        _nw98[(new BigNumber(3)).toNumber()] = (((_467_v14)[_module.__default.safeIndex(new BigNumber((_468_v15).length), new BigNumber((_467_v14).length))]) ? (_466_v13) : (_dafny.MultiSet.fromElements((_467_v14)[_module.__default.safeIndex(new BigNumber((_module.__default.fm23(_module.D2.create_DC7(_469_v16), p0, (_471_v18).dtor_cf15, globalState)).length), new BigNumber((_467_v14).length))], p3, false)));
        _nw98[(new BigNumber(4)).toNumber()] = _466_v13;
        _nw98[(new BigNumber(5)).toNumber()] = ((p3) ? (_466_v13) : (_466_v13));
        _nw98[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.FromArray(_467_v14);
        _nw98[(new BigNumber(7)).toNumber()] = (_466_v13).Union(_466_v13);
        _nw98[(new BigNumber(8)).toNumber()] = _module.__default.fm24(new BigNumber((_472_v19).length), p3, globalState);
        _nw98[(new BigNumber(9)).toNumber()] = (_466_v13).Union(_466_v13);
        _nw98[(new BigNumber(10)).toNumber()] = _466_v13;
        _nw98[(new BigNumber(11)).toNumber()] = _466_v13;
        _nw98[(new BigNumber(12)).toNumber()] = (_dafny.MultiSet.FromArray(_467_v14)).Difference(_466_v13);
        _nw98[(new BigNumber(13)).toNumber()] = _466_v13;
        _nw98[(new BigNumber(14)).toNumber()] = _466_v13;
        _nw98[(new BigNumber(15)).toNumber()] = _466_v13;
        _nw98[(new BigNumber(16)).toNumber()] = _466_v13;
        _nw98[(new BigNumber(17)).toNumber()] = _466_v13;
        _nw98[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.fromElements(false, p3);
        _nw98[(new BigNumber(19)).toNumber()] = _466_v13;
        _nw98[(new BigNumber(20)).toNumber()] = _466_v13;
        _nw98[(new BigNumber(21)).toNumber()] = (_466_v13).Difference(_466_v13);
        _nw98[(new BigNumber(22)).toNumber()] = _466_v13;
        _nw98[(new BigNumber(23)).toNumber()] = _466_v13;
        _473_v20 = _nw98;
        _473_v20 = _473_v20;
        let _474_v21;
        _474_v21 = _dafny.Seq.of(new BigNumber(228));
        _474_v21 = _dafny.Seq.Concat(_474_v21, _474_v21);
        let _475_v22;
        _475_v22 = new _dafny.CodePoint('n'.codePointAt(0));
        (globalState).f19 = _475_v22;
        if ((p0).isEqualTo(new BigNumber((_474_v21).length))) {
          let _476_v23;
          let _nw99 = Array((new BigNumber(5)).toNumber()).fill(false);
          _476_v23 = _nw99;
          let _index97 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_476_v23).length));
          (_476_v23)[_index97] = false;
          let _index98 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_476_v23).length));
          let _rhs89 = p3;
          let _rhs90 = _475_v22;
          let _rhs91 = (((p1) ? (p1) : (p1))) === (p1);
          let _lhs77 = _476_v23;
          let _lhs78 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_476_v23).length));
          let _lhs79 = globalState;
          _lhs77[_lhs78] = _rhs89;
          _475_v22 = _rhs90;
          _lhs79.f22 = _rhs91;
          let _477_v24;
          _477_v24 = _module.D2.create_DC7(_469_v16);
          let _478_v25;
          _478_v25 = _dafny.Map.Empty.slice().updateUnsafe(_476_v23,_472_v19);
          let _479_v26;
          let _nw100 = Array((new BigNumber(16)).toNumber());
          _nw100[(_dafny.ZERO).toNumber()] = _module.__default.fm23(_477_v24, (_dafny.ZERO).minus(p2), p2, globalState);
          _nw100[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(400)), ((_480_v22) => function (_481_i1) {
            return _480_v22;
          })(_475_v22));
          _nw100[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_472_v19, _472_v19);
          _nw100[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_472_v19, _dafny.Seq.Create(_module.__default.abs(new BigNumber(945)), ((_482_v22) => function (_483_i2) {
            return _482_v22;
          })(_475_v22)));
          _nw100[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_472_v19, _472_v19);
          _nw100[(new BigNumber(5)).toNumber()] = (((_478_v25).contains(_476_v23)) ? ((_478_v25).get(_476_v23)) : (_472_v19));
          _nw100[(new BigNumber(6)).toNumber()] = _472_v19;
          _nw100[(new BigNumber(7)).toNumber()] = _472_v19;
          _nw100[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-190)), ((_484_v19) => function (_485_i3) {
            return (_484_v19)[_module.__default.safeIndex(new BigNumber((_484_v19).length), new BigNumber((_484_v19).length))];
          })(_472_v19));
          _nw100[(new BigNumber(9)).toNumber()] = _472_v19;
          _nw100[(new BigNumber(10)).toNumber()] = _472_v19;
          _nw100[(new BigNumber(11)).toNumber()] = _module.__default.fm23(_477_v24, new BigNumber(-323), p2, globalState);
          _nw100[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_472_v19, _module.__default.safeIndex(new BigNumber((_474_v21).length), new BigNumber((_472_v19).length)), _475_v22);
          _nw100[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("ubj");
          _nw100[(new BigNumber(14)).toNumber()] = _472_v19;
          _nw100[(new BigNumber(15)).toNumber()] = _472_v19;
          _479_v26 = _nw100;
          let _index99 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_479_v26).length));
          (_479_v26)[_index99] = _472_v19;
          let _index100 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_476_v23).length));
          let _index101 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_479_v26).length));
          let _rhs92 = (_this).fm21(new BigNumber(532), globalState);
          let _rhs93 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fye"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uerjb"), _472_v19));
          let _rhs94 = false;
          let _lhs80 = _476_v23;
          let _lhs81 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_476_v23).length));
          let _lhs82 = _479_v26;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_479_v26).length));
          let _lhs84 = globalState;
          _lhs80[_lhs81] = _rhs92;
          _lhs82[_lhs83] = _rhs93;
          _lhs84.f22 = _rhs94;
          (globalState).f10 = p0;
          let _486_v27;
          let _nw101 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
          _486_v27 = _nw101;
          let _487_v28;
          _487_v28 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_476_v23)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_476_v23).length))]);
          let _index102 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_486_v27).length));
          (_486_v27)[_index102] = _487_v28;
          let _488_v29;
          let _nw102 = Array((new BigNumber(24)).toNumber());
          _488_v29 = _nw102;
          let _489_v30;
          let _nw103 = new _module.C0();
          _nw103.__ctor((_476_v23)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_476_v23).length))]);
          _489_v30 = _nw103;
          let _index103 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_488_v29).length));
          (_488_v29)[_index103] = _489_v30;
          let _490_v31;
          let _init7 = ((_491_p0) => function (_492_i4) {
            return _module.__default.safeModuloInt(_492_i4, _491_p0);
          })(p0);
          let _nw104 = Array((new BigNumber(8)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw104.length); _i0_7++) {
            _nw104[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _490_v31 = _nw104;
          let _index104 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_490_v31).length));
          (_490_v31)[_index104] = _module.__default.safeDivisionInt(p2, p0);
          let _493_v32;
          _493_v32 = _module.D3.create_DC10(_dafny.MultiSet.FromArray(_467_v14));
          let _494_v33;
          _494_v33 = _dafny.Map.Empty.slice().updateUnsafe((_479_v26)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_479_v26).length))],_490_v31);
          let _495_v34;
          _495_v34 = _490_v31;
          let _496_v35;
          _496_v35 = _dafny.Seq.of(_490_v31);
          let _497_v37;
          let _nw105 = Array((new BigNumber(14)).toNumber());
          _nw105[(_dafny.ZERO).toNumber()] = _490_v31;
          _nw105[(_dafny.ONE).toNumber()] = _490_v31;
          _nw105[(new BigNumber(2)).toNumber()] = _490_v31;
          _nw105[(new BigNumber(3)).toNumber()] = _490_v31;
          _nw105[(new BigNumber(4)).toNumber()] = _490_v31;
          _nw105[(new BigNumber(5)).toNumber()] = (((_this).fm21(p2, globalState)) ? ((((_494_v33).contains(_472_v19)) ? ((_494_v33).get(_472_v19)) : (_490_v31))) : (_490_v31));
          _nw105[(new BigNumber(6)).toNumber()] = _490_v31;
          _nw105[(new BigNumber(7)).toNumber()] = _490_v31;
          _nw105[(new BigNumber(8)).toNumber()] = _490_v31;
          _nw105[(new BigNumber(9)).toNumber()] = (_495_v34);
          _nw105[(new BigNumber(10)).toNumber()] = (_496_v35)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((function () {
            let _coll18 = new _dafny.Map();
            for (const _compr_18 of (_dafny.Seq.of(new BigNumber(-571))).Elements) {
              let _498_v36 = _compr_18;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(-571)), _498_v36)) {
                _coll18.push([(_498_v36).multipliedBy((_474_v21)[_module.__default.safeIndex(p0, new BigNumber((_474_v21).length))]),p0]);
              }
            }
            return _coll18;
          }()).length)), new BigNumber((_496_v35).length))];
          _nw105[(new BigNumber(11)).toNumber()] = _490_v31;
          _nw105[(new BigNumber(12)).toNumber()] = _490_v31;
          _nw105[(new BigNumber(13)).toNumber()] = _490_v31;
          _497_v37 = _nw105;
          let _index105 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_486_v27).length));
          let _index106 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_488_v29).length));
          let _index107 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_490_v31).length));
          let _rhs95 = function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of _dafny.IntegerRange(new BigNumber(207), new BigNumber(900))) {
              let _499_v38 = _compr_19;
              if (((new BigNumber(207)).isLessThanOrEqualTo(_499_v38)) && ((_499_v38).isLessThan(new BigNumber(900)))) {
                _coll19.push([(_499_v38).plus(p0),(p0).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_474_v21).length))))]);
              }
            }
            return _coll19;
          }();
          let _rhs96 = _489_v30;
          let _rhs97 = (p2).plus(p2);
          let _rhs98 = _493_v32;
          let _rhs99 = (((p3) === (!((_476_v23)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_476_v23).length))]))) ? (_497_v37) : (_497_v37));
          let _lhs85 = _486_v27;
          let _lhs86 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_486_v27).length));
          let _lhs87 = _488_v29;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_488_v29).length));
          let _lhs89 = _490_v31;
          let _lhs90 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_490_v31).length));
          _lhs85[_lhs86] = _rhs95;
          _lhs87[_lhs88] = _rhs96;
          _lhs89[_lhs90] = _rhs97;
          _493_v32 = _rhs98;
          _497_v37 = _rhs99;
          let _500_v39;
          _500_v39 = _dafny.Map.Empty.slice().updateUnsafe((_476_v23)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_476_v23).length))],(_490_v31)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_490_v31).length))]);
          let _501_v40;
          _501_v40 = _dafny.Map.Empty.slice().updateUnsafe((_476_v23)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_476_v23).length))],_500_v39);
          let _index108 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_479_v26).length));
          (_479_v26)[_index108] = _dafny.Seq.update(_dafny.Seq.Concat(_472_v19, _dafny.Seq.Concat((_479_v26)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_479_v26).length))], (_479_v26)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_479_v26).length))])), _module.__default.safeIndex(new BigNumber(((_501_v40).Merge(_501_v40)).length), new BigNumber((_dafny.Seq.Concat(_472_v19, _dafny.Seq.Concat((_479_v26)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_479_v26).length))], (_479_v26)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_479_v26).length))]))).length)), _475_v22);
        } else {
          let _502_v41;
          let _init8 = ((_503_v13, _504_p1) => function (_505_i5) {
            return _module.D1.create_DC4(_dafny.Map.Empty.slice().updateUnsafe(_503_v13,_504_p1));
          })(_466_v13, p1);
          let _nw106 = Array((new BigNumber(16)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw106.length); _i0_8++) {
            _nw106[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _502_v41 = _nw106;
          let _506_v43;
          _506_v43 = _dafny.Map.Empty.slice().updateUnsafe(_466_v13,p0);
          let _507_v44;
          _507_v44 = _module.D1.create_DC4(function () {
  let _coll20 = new _dafny.Map();
  for (const _compr_20 of (_506_v43).Keys.Elements) {
    let _508_v42 = _compr_20;
    if ((_506_v43).contains(_508_v42)) {
      _coll20.push([_508_v42,p3]);
    }
  }
  return _coll20;
}());
          let _index109 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_502_v41).length));
          (_502_v41)[_index109] = _507_v44;
          let _509_v45;
          _509_v45 = _dafny.Map.Empty.slice().updateUnsafe(_466_v13,p3);
          let _pat_let_tv6 = p1;
          let _pat_let_tv7 = p1;
          let _index110 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_502_v41).length));
          (_502_v41)[_index110] = function (_pat_let6_0) {
            return function (_510_dt__update__tmp_h0) {
              return function (_pat_let7_0) {
                return function (_511_dt__update_hcf8_h0) {
                  return _module.D1.create_DC4(_511_dt__update_hcf8_h0);
                }(_pat_let7_0);
              }(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_pat_let_tv6),_pat_let_tv7));
            }(_pat_let6_0);
          }(_module.D1.create_DC4(_509_v45));
          let _512_v46;
          let _init9 = function (_513_i6) {
            return (_513_i6).multipliedBy(new BigNumber(-650));
          };
          let _nw107 = Array((new BigNumber(23)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw107.length); _i0_9++) {
            _nw107[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _512_v46 = _nw107;
          let _index111 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_512_v46).length));
          (_512_v46)[_index111] = p2;
          let _index112 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_512_v46).length));
          let _rhs100 = p2;
          let _rhs101 = _475_v22;
          let _rhs102 = (_dafny.ZERO).minus(p2);
          let _lhs91 = _512_v46;
          let _lhs92 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_512_v46).length));
          let _lhs93 = globalState;
          let _lhs94 = globalState;
          _lhs91[_lhs92] = _rhs100;
          _lhs93.f19 = _rhs101;
          _lhs94.f10 = _rhs102;
          _471_v18 = ((false) ? (_471_v18) : (_471_v18));
          let _514_v47;
          let _init10 = ((_515_p1, _516_v19, _517_p3) => function (_518_i7) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(20),_515_p1),_515_p1)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_516_v19).length),false),_517_p3));
          })(p1, _472_v19, p3);
          let _nw108 = Array((new BigNumber(27)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw108.length); _i0_10++) {
            _nw108[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _514_v47 = _nw108;
          let _519_v48;
          _519_v48 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          let _520_v49;
          _520_v49 = _dafny.Map.Empty.slice().updateUnsafe(_519_v48,(_this).fm21((_512_v46)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_512_v46).length))], globalState));
          let _index113 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_514_v47).length));
          (_514_v47)[_index113] = _520_v49;
          let _index114 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_512_v46).length));
          let _index115 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_514_v47).length));
          let _rhs103 = p0;
          let _rhs104 = _520_v49;
          let _rhs105 = _dafny.Seq.Concat(_dafny.Seq.Concat(_474_v21, _474_v21), _dafny.Seq.of(new BigNumber(208)));
          let _rhs106 = p1;
          let _lhs95 = _512_v46;
          let _lhs96 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_512_v46).length));
          let _lhs97 = _514_v47;
          let _lhs98 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_514_v47).length));
          let _lhs99 = globalState;
          _lhs95[_lhs96] = _rhs103;
          _lhs97[_lhs98] = _rhs104;
          _474_v21 = _rhs105;
          _lhs99.f13 = _rhs106;
          let _521_v50;
          _521_v50 = _module.D0.create_DC2(_module.__default.fm0(p2, globalState), p1);
          let _522_v51;
          let _nw109 = Array((new BigNumber(15)).toNumber());
          _nw109[(_dafny.ZERO).toNumber()] = p3;
          _nw109[(_dafny.ONE).toNumber()] = p3;
          _nw109[(new BigNumber(2)).toNumber()] = p3;
          _nw109[(new BigNumber(3)).toNumber()] = !(p1);
          _nw109[(new BigNumber(4)).toNumber()] = !(!(!(_468_v15).contains(p3)));
          _nw109[(new BigNumber(5)).toNumber()] = (new BigNumber(867)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-93)), ((_523_p3, _524_v22) => function (_525_i8) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_523_p3,_524_v22)).length);
          })(p3, _475_v22))).length));
          _nw109[(new BigNumber(6)).toNumber()] = (p0).isLessThanOrEqualTo(p0);
          _nw109[(new BigNumber(7)).toNumber()] = p1;
          _nw109[(new BigNumber(8)).toNumber()] = p1;
          _nw109[(new BigNumber(9)).toNumber()] = p3;
          _nw109[(new BigNumber(10)).toNumber()] = (!(p3)) || (p1);
          _nw109[(new BigNumber(11)).toNumber()] = p1;
          _nw109[(new BigNumber(12)).toNumber()] = p1;
          _nw109[(new BigNumber(13)).toNumber()] = p1;
          _nw109[(new BigNumber(14)).toNumber()] = (p0).isLessThan((_521_v50).dtor_cf5);
          _522_v51 = _nw109;
          let _index116 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_522_v51).length));
          (_522_v51)[_index116] = p1;
          let _526_v52;
          _526_v52 = _module.D0.create_DC1(p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,p0)).length), p2, _522_v51);
          let _527_v53;
          _527_v53 = _module.D0.create_DC3(_526_v52);
          let _index117 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_522_v51).length));
          let _index118 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_512_v46).length));
          let _index119 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_512_v46).length));
          let _rhs107 = p1;
          let _rhs108 = ((p3) ? (_527_v53) : (_527_v53));
          let _rhs109 = ((p2).multipliedBy((_512_v46)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_512_v46).length))])).multipliedBy((_512_v46)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_512_v46).length))]);
          let _rhs110 = p0;
          let _lhs100 = _522_v51;
          let _lhs101 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_522_v51).length));
          let _lhs102 = _512_v46;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_512_v46).length));
          let _lhs104 = _512_v46;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_512_v46).length));
          _lhs100[_lhs101] = _rhs107;
          _527_v53 = _rhs108;
          _lhs102[_lhs103] = _rhs109;
          _lhs104[_lhs105] = _rhs110;
        }
        let _528_v54;
        let _init11 = ((_529_p3) => function (_530_i9) {
          return _529_p3;
        })(p3);
        let _nw110 = Array((new BigNumber(19)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw110.length); _i0_11++) {
          _nw110[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _528_v54 = _nw110;
        let _531_v55;
        _531_v55 = _module.D0.create_DC1(p2, new BigNumber(933), new BigNumber(-728), _528_v54);
        let _532_v56;
        _532_v56 = _dafny.Seq.of(_531_v55, _531_v55);
        let _533_v57;
        _533_v57 = _dafny.Map.Empty.slice().updateUnsafe(p1,_532_v56);
        let _534_v58;
        _534_v58 = _dafny.Set.fromElements(_dafny.Seq.of(_531_v55, _531_v55, _531_v55), _dafny.Seq.Concat((((_533_v57).contains(p3)) ? ((_533_v57).get(p3)) : (_532_v56)), _532_v56));
        let _535_v59;
        let _nw111 = Array((_dafny.ONE).toNumber()).fill([]);
        _535_v59 = _nw111;
        let _536_v60;
        let _nw112 = Array((new BigNumber(12)).toNumber()).fill(_module.D0.Default());
        _536_v60 = _nw112;
        let _index120 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_535_v59).length));
        (_535_v59)[_index120] = _536_v60;
        let _537_v61;
        let _init12 = ((_538_v22) => function (_539_i10) {
          return _538_v22;
        })(_475_v22);
        let _nw113 = Array((new BigNumber(6)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw113.length); _i0_12++) {
          _nw113[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _537_v61 = _nw113;
        let _index121 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_537_v61).length));
        (_537_v61)[_index121] = new _dafny.CodePoint('e'.codePointAt(0));
        let _index122 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_535_v59).length));
        let _index123 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_537_v61).length));
        let _rhs111 = _534_v58;
        let _rhs112 = _536_v60;
        let _rhs113 = !_dafny.Seq.contains(_dafny.Seq.Concat(_467_v14, _dafny.Seq.update(_467_v14, _module.__default.safeIndex(new BigNumber((_472_v19).length), new BigNumber((_467_v14).length)), p1)), p3);
        let _rhs114 = p2;
        let _rhs115 = _475_v22;
        let _lhs106 = _535_v59;
        let _lhs107 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_535_v59).length));
        let _lhs108 = globalState;
        let _lhs109 = globalState;
        let _lhs110 = _537_v61;
        let _lhs111 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_537_v61).length));
        _534_v58 = _rhs111;
        _lhs106[_lhs107] = _rhs112;
        _lhs108.f13 = _rhs113;
        _lhs109.f10 = _rhs114;
        _lhs110[_lhs111] = _rhs115;
      }
      let _hi1 = _module.__default.safeModuloInt(new BigNumber(61), new BigNumber(131));
      for (let _540_i11 = new BigNumber(461); _540_i11.isLessThan(_hi1); _540_i11 = _540_i11.plus(_dafny.ONE)) {
        let _541_v62;
        _541_v62 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
        _541_v62 = _541_v62;
        let _542_v63;
        _542_v63 = _dafny.Seq.UnicodeFromString("hrq");
        let _543_v64;
        _543_v64 = _dafny.Seq.of(new BigNumber((_542_v63).length));
        let _544_v65;
        _544_v65 = _dafny.Set.fromElements(new BigNumber((_543_v64).length));
        let _545_v66;
        let _nw114 = new _module.C0();
        _nw114.__ctor(!(p1));
        _545_v66 = _nw114;
        let _546_v67;
        _546_v67 = _dafny.Seq.of(_545_v66);
        let _547_v68;
        _547_v68 = new _dafny.CodePoint('v'.codePointAt(0));
        let _548_v69;
        _548_v69 = _dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), _547_v68, _547_v68, _547_v68);
        let _549_v70;
        _549_v70 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_546_v67, _module.__default.safeIndex(new BigNumber((_548_v69).length), new BigNumber((_546_v67).length)), _545_v66), _546_v67);
        let _550_v71;
        _550_v71 = _dafny.MultiSet.fromElements(p2);
        let _551_v72;
        let _nw115 = Array((new BigNumber(24)).toNumber());
        _nw115[(_dafny.ZERO).toNumber()] = p0;
        _nw115[(_dafny.ONE).toNumber()] = p0;
        _nw115[(new BigNumber(2)).toNumber()] = _540_i11;
        _nw115[(new BigNumber(3)).toNumber()] = p0;
        _nw115[(new BigNumber(4)).toNumber()] = (p0).multipliedBy(new BigNumber((_544_v65).length));
        _nw115[(new BigNumber(5)).toNumber()] = p2;
        _nw115[(new BigNumber(6)).toNumber()] = _540_i11;
        _nw115[(new BigNumber(7)).toNumber()] = new BigNumber((_542_v63).length);
        _nw115[(new BigNumber(8)).toNumber()] = _540_i11;
        _nw115[(new BigNumber(9)).toNumber()] = new BigNumber((_542_v63).length);
        _nw115[(new BigNumber(10)).toNumber()] = new BigNumber((_module.__default.fm22(p3, globalState)).length);
        _nw115[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_542_v63).length), _540_i11);
        _nw115[(new BigNumber(12)).toNumber()] = _540_i11;
        _nw115[(new BigNumber(13)).toNumber()] = (_this).fm2(globalState);
        _nw115[(new BigNumber(14)).toNumber()] = (((_549_v70).contains(_546_v67)) ? ((_549_v70).get(_546_v67)) : (_module.__default.fm0(new BigNumber((_550_v71).cardinality()), globalState)));
        _nw115[(new BigNumber(15)).toNumber()] = p2;
        _nw115[(new BigNumber(16)).toNumber()] = p0;
        _nw115[(new BigNumber(17)).toNumber()] = p2;
        _nw115[(new BigNumber(18)).toNumber()] = _540_i11;
        _nw115[(new BigNumber(19)).toNumber()] = p0;
        _nw115[(new BigNumber(20)).toNumber()] = _module.__default.safeModuloInt((_545_v66).fm13(p2, globalState), new BigNumber((_dafny.Seq.of(_547_v68, _547_v68, _547_v68)).length));
        _nw115[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_545_v66).f32,_547_v68)).length);
        _nw115[(new BigNumber(22)).toNumber()] = _module.__default.safeDivisionInt((_545_v66).fm13(p0, globalState), _module.__default.fm0(_540_i11, globalState));
        _nw115[(new BigNumber(23)).toNumber()] = p0;
        _551_v72 = _nw115;
        let _552_v73;
        _552_v73 = _dafny.Set.fromElements((_545_v66).f32);
        let _index124 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_551_v72).length));
        (_551_v72)[_index124] = _module.__default.safeModuloInt(new BigNumber((_552_v73).length), p2);
        let _index125 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_551_v72).length));
        (_551_v72)[_index125] = p2;
        let _553_v74;
        let _init13 = ((_554_v66, _555_i11, _556_v72) => function (_557_i12) {
          return (_555_i11).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_554_v66).f32,_module.D2.create_DC7(_dafny.Map.Empty.slice().updateUnsafe(_555_i11,(_556_v72)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((_556_v72).length))])))).length));
        })(_545_v66, _540_i11, _551_v72);
        let _nw116 = Array((new BigNumber(4)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw116.length); _i0_13++) {
          _nw116[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _553_v74 = _nw116;
        let _index126 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_553_v74).length));
        (_553_v74)[_index126] = p1;
        let _index127 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_551_v72).length));
        let _index128 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_551_v72).length));
        let _index129 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_553_v74).length));
        let _rhs116 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of((_543_v64)[_module.__default.safeIndex(_540_i11, new BigNumber((_543_v64).length))], (_dafny.ZERO).minus(_540_i11), new BigNumber((_542_v63).length))).length), p2));
        let _rhs117 = _540_i11;
        let _rhs118 = (_540_i11).isLessThanOrEqualTo(_540_i11);
        let _lhs112 = _551_v72;
        let _lhs113 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_551_v72).length));
        let _lhs114 = _551_v72;
        let _lhs115 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_551_v72).length));
        let _lhs116 = _553_v74;
        let _lhs117 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_553_v74).length));
        _lhs112[_lhs113] = _rhs116;
        _lhs114[_lhs115] = _rhs117;
        _lhs116[_lhs117] = _rhs118;
        let _558_v75;
        let _nw117 = new _module.C0();
        _nw117.__ctor((_553_v74)[_module.__default.safeIndex(new BigNumber(971), new BigNumber((_553_v74).length))]);
        _558_v75 = _nw117;
      }
      let _559_v76;
      let _nw118 = Array((new BigNumber(17)).toNumber()).fill([]);
      _559_v76 = _nw118;
      let _560_v77;
      _560_v77 = _dafny.Seq.of(p2);
      let _561_v78;
      _561_v78 = _dafny.Seq.of((_this).fm21(p0, globalState));
      let _562_v79;
      _562_v79 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(28)), function (_563_i13) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })).length));
      let _564_v80;
      _564_v80 = _dafny.Seq.UnicodeFromString("cxrrfwi");
      let _565_v81;
      _565_v81 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm2(globalState),_564_v80);
      let _566_v82;
      let _nw119 = Array((new BigNumber(10)).toNumber());
      _nw119[(_dafny.ZERO).toNumber()] = p2;
      _nw119[(_dafny.ONE).toNumber()] = p2;
      _nw119[(new BigNumber(2)).toNumber()] = new BigNumber((_560_v77).length);
      _nw119[(new BigNumber(3)).toNumber()] = new BigNumber(380);
      _nw119[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_561_v78).length))).length);
      _nw119[(new BigNumber(5)).toNumber()] = _module.__default.fm0(p0, globalState);
      _nw119[(new BigNumber(6)).toNumber()] = p2;
      _nw119[(new BigNumber(7)).toNumber()] = new BigNumber(411);
      _nw119[(new BigNumber(8)).toNumber()] = (((_562_v79).contains(p0)) ? ((_562_v79).get(p0)) : (new BigNumber(((((_565_v81).contains(p0)) ? ((_565_v81).get(p0)) : (_564_v80))).length)));
      _nw119[(new BigNumber(9)).toNumber()] = p0;
      _566_v82 = _nw119;
      let _567_v83;
      let _nw120 = Array((new BigNumber(3)).toNumber());
      _nw120[(_dafny.ZERO).toNumber()] = _566_v82;
      _nw120[(_dafny.ONE).toNumber()] = _566_v82;
      _nw120[(new BigNumber(2)).toNumber()] = _566_v82;
      _567_v83 = _nw120;
      let _index130 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_559_v76).length));
      (_559_v76)[_index130] = _567_v83;
      let _568_v84;
      _568_v84 = _dafny.Map.Empty.slice().updateUnsafe(p3,_567_v83);
      let _index131 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_559_v76).length));
      (_559_v76)[_index131] = (((_568_v84).contains(p1)) ? ((_568_v84).get(p1)) : (_567_v83));
      let _569_v85;
      _569_v85 = _dafny.Seq.of(_561_v78, _561_v78);
      let _570_v86;
      _570_v86 = _module.D1.create_DC6(p2, _569_v85, p0, p3);
      let _571_v87;
      _571_v87 = _dafny.Map.Empty.slice().updateUnsafe(true,p3);
      let _572_v88;
      _572_v88 = _dafny.Map.Empty.slice().updateUnsafe((((_571_v87).contains(p3)) ? ((_571_v87).get(p3)) : (p1)),p1);
      let _573_v89;
      _573_v89 = _module.D0.create_DC2(p0, p1);
      let _574_v90;
      let _nw121 = Array((new BigNumber(19)).toNumber());
      _nw121[(_dafny.ZERO).toNumber()] = p3;
      _nw121[(_dafny.ONE).toNumber()] = p1;
      _nw121[(new BigNumber(2)).toNumber()] = p1;
      _nw121[(new BigNumber(3)).toNumber()] = p1;
      _nw121[(new BigNumber(4)).toNumber()] = (_570_v86).dtor_cf16;
      _nw121[(new BigNumber(5)).toNumber()] = p3;
      _nw121[(new BigNumber(6)).toNumber()] = p3;
      _nw121[(new BigNumber(7)).toNumber()] = (((_571_v87).contains((((_572_v88).contains(p3)) ? ((_572_v88).get(p3)) : (p1)))) ? ((_571_v87).get((((_572_v88).contains(p3)) ? ((_572_v88).get(p3)) : (p1)))) : (p3));
      _nw121[(new BigNumber(8)).toNumber()] = p1;
      _nw121[(new BigNumber(9)).toNumber()] = p3;
      _nw121[(new BigNumber(10)).toNumber()] = p1;
      _nw121[(new BigNumber(11)).toNumber()] = (_573_v89).dtor_cf6;
      _nw121[(new BigNumber(12)).toNumber()] = (_561_v78)[_module.__default.safeIndex(p0, new BigNumber((_561_v78).length))];
      _nw121[(new BigNumber(13)).toNumber()] = p1;
      _nw121[(new BigNumber(14)).toNumber()] = p1;
      _nw121[(new BigNumber(15)).toNumber()] = p3;
      _nw121[(new BigNumber(16)).toNumber()] = p3;
      _nw121[(new BigNumber(17)).toNumber()] = !(p1);
      _nw121[(new BigNumber(18)).toNumber()] = false;
      _574_v90 = _nw121;
      let _575_v91;
      _575_v91 = _dafny.Set.fromElements(_574_v90);
      let _576_v92;
      _576_v92 = _dafny.Set.fromElements((_560_v77)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("cqty")).length), new BigNumber((_560_v77).length))]);
      let _577_v93;
      _577_v93 = _dafny.Seq.of(_dafny.Set.fromElements(_574_v90, _574_v90, _574_v90), _575_v91);
      let _rhs119 = new BigNumber((((_576_v92).Union(_dafny.Set.fromElements((_dafny.ZERO).minus(p2)))).Intersect(((p1) ? (_576_v92) : (_576_v92)))).length);
      let _rhs120 = (_577_v93)[_module.__default.safeIndex((p0).multipliedBy(p2), new BigNumber((_577_v93).length))];
      let _rhs121 = p0;
      let _lhs118 = globalState;
      let _lhs119 = globalState;
      _lhs118.f10 = _rhs119;
      _575_v91 = _rhs120;
      _lhs119.f10 = _rhs121;
      (globalState).f13 = p3;
      let _578_v94;
      _578_v94 = new _dafny.CodePoint('y'.codePointAt(0));
      let _579_v95;
      _579_v95 = _module.D4.create_DC13(_578_v94);
      let _pat_let_tv8 = _578_v94;
      let _pat_let_tv9 = _578_v94;
      let _pat_let_tv10 = _578_v94;
      let _rhs122 = p0;
      let _rhs123 = p2;
      let _rhs124 = ((p1) ? (_578_v94) : (_578_v94));
      let _rhs125 = function (_source11) {
        if (_source11.is_DC14) {
          let _580___mcc_h0 = (_source11).cf25;
          let _581_cf25 = _580___mcc_h0;
          return _pat_let_tv8;
        } else if (_source11.is_DC13) {
          let _582___mcc_h1 = (_source11).cf24;
          let _583_cf24 = _582___mcc_h1;
          return _pat_let_tv9;
        } else {
          let _584___mcc_h2 = (_source11).cf26;
          let _585_cf26 = _584___mcc_h2;
          return _pat_let_tv10;
        }
      }(_579_v95);
      let _lhs120 = globalState;
      let _lhs121 = globalState;
      let _lhs122 = globalState;
      let _lhs123 = globalState;
      _lhs120.f10 = _rhs122;
      _lhs121.f10 = _rhs123;
      _lhs122.f19 = _rhs124;
      _lhs123.f19 = _rhs125;
      return;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let r1 = false;
      let _586_v0;
      _586_v0 = _dafny.Seq.UnicodeFromString("btqf");
      _586_v0 = _586_v0;
      if (p2) {
        (globalState).f13 = !(p2);
        let _587_v1;
        _587_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(408),!(p2));
        let _588_v2;
        let _nw122 = Array((new BigNumber(17)).toNumber());
        _nw122[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw122[(_dafny.ONE).toNumber()] = (p1).plus(p1);
        _nw122[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_586_v0, _586_v0)).length);
        _nw122[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((p1).multipliedBy(new BigNumber((_586_v0).length)))));
        _nw122[(new BigNumber(4)).toNumber()] = new BigNumber((p0).length);
        _nw122[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((p1).multipliedBy(p1));
        _nw122[(new BigNumber(6)).toNumber()] = p1;
        _nw122[(new BigNumber(7)).toNumber()] = p1;
        _nw122[(new BigNumber(8)).toNumber()] = p1;
        _nw122[(new BigNumber(9)).toNumber()] = p1;
        _nw122[(new BigNumber(10)).toNumber()] = p1;
        _nw122[(new BigNumber(11)).toNumber()] = p1;
        _nw122[(new BigNumber(12)).toNumber()] = new BigNumber((_587_v1).length);
        _nw122[(new BigNumber(13)).toNumber()] = p1;
        _nw122[(new BigNumber(14)).toNumber()] = p1;
        _nw122[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw122[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_586_v0, _586_v0)).length);
        _588_v2 = _nw122;
        _588_v2 = ((!(p2)) ? (_588_v2) : (_588_v2));
        let _589_v3;
        let _nw123 = Array((new BigNumber(12)).toNumber()).fill([]);
        _589_v3 = _nw123;
        let _590_v4;
        _590_v4 = _dafny.Seq.of(p2, p2, p2);
        let _rhs126 = ((!(p2)) ? (p1) : ((((_590_v4)[_module.__default.safeIndex(p1, new BigNumber((_590_v4).length))]) ? (p1) : (new BigNumber(-576)))));
        let _rhs127 = _dafny.Seq.UnicodeFromString("ocl");
        let _rhs128 = _589_v3;
        let _rhs129 = _module.__default.safeModuloInt(new BigNumber(548), p1);
        let _rhs130 = (_this).fm21((_dafny.ZERO).minus(p1), globalState);
        let _lhs124 = globalState;
        let _lhs125 = globalState;
        let _lhs126 = globalState;
        _lhs124.f10 = _rhs126;
        _586_v0 = _rhs127;
        _589_v3 = _rhs128;
        _lhs125.f10 = _rhs129;
        _lhs126.f13 = _rhs130;
        if (p2) {
          let _591_v5;
          _591_v5 = _dafny.Seq.of(p1, p1);
          _591_v5 = _dafny.Seq.update(_dafny.Seq.Concat(_591_v5, _dafny.Seq.of(p1, (_dafny.ZERO).minus(p1), p1)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_591_v5, _dafny.Seq.of(p1, (_dafny.ZERO).minus(p1), p1))).length)), ((p2) ? (p1) : (p1)));
          (globalState).f7 = p2;
          let _592_v6;
          _592_v6 = _dafny.MultiSet.fromElements(p1, new BigNumber((_586_v0).length));
          let _593_v7;
          _593_v7 = _dafny.Map.Empty.slice().updateUnsafe(_592_v6,_dafny.Seq.update(p0, _module.__default.safeIndex(p1, new BigNumber((p0).length)), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())));
          let _594_v8;
          _594_v8 = _module.D1.create_DC5(p1, !(p2), p2, _593_v7);
          let _595_v9;
          let _nw124 = Array((new BigNumber(10)).toNumber());
          _nw124[(_dafny.ZERO).toNumber()] = _module.__default.fm25(p2, !(p2), p1, p2, globalState);
          _nw124[(_dafny.ONE).toNumber()] = _594_v8;
          _nw124[(new BigNumber(2)).toNumber()] = _594_v8;
          _nw124[(new BigNumber(3)).toNumber()] = _594_v8;
          _nw124[(new BigNumber(4)).toNumber()] = _594_v8;
          _nw124[(new BigNumber(5)).toNumber()] = _594_v8;
          _nw124[(new BigNumber(6)).toNumber()] = _594_v8;
          _nw124[(new BigNumber(7)).toNumber()] = _594_v8;
          _nw124[(new BigNumber(8)).toNumber()] = _module.D1.create_DC5(p1, p2, true, _593_v7);
          _nw124[(new BigNumber(9)).toNumber()] = _594_v8;
          _595_v9 = _nw124;
          let _index132 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_595_v9).length));
          (_595_v9)[_index132] = _594_v8;
          let _index133 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_595_v9).length));
          (_595_v9)[_index133] = _594_v8;
          let _596_v10;
          let _out2;
          _out2 = (_this).m6((new BigNumber((_591_v5).length)).minus(p1), (new BigNumber((_586_v0).length)).minus(p1), (new BigNumber((_590_v4).length)).multipliedBy(p1), _module.__default.safeDivisionInt(p1, p1), globalState);
          _596_v10 = _out2;
          let _597_v11;
          _597_v11 = _dafny.Set.fromElements(true, p2, p2);
          (globalState).f13 = !(!((_597_v11).contains(p2))) || (p2);
        } else {
          let _598_v12;
          _598_v12 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.Concat(_dafny.Seq.update(p0, _module.__default.safeIndex(p1, new BigNumber((p0).length)), p1), p0));
          let _599_v13;
          _599_v13 = _dafny.MultiSet.fromElements(p2);
          let _600_v14;
          _600_v14 = _dafny.Map.Empty.slice().updateUnsafe(_599_v13,true);
          let _601_v15;
          _601_v15 = _module.D1.create_DC4(_600_v14);
          let _602_v16;
          _602_v16 = _dafny.Seq.of(_601_v15, _601_v15);
          let _603_v17;
          _603_v17 = new _dafny.CodePoint('k'.codePointAt(0));
          let _604_v18;
          _604_v18 = _dafny.Map.Empty.slice().updateUnsafe(_588_v2,_603_v17);
          let _605_v19;
          _605_v19 = _dafny.Set.fromElements(p2, p2);
          let _606_v20;
          _606_v20 = _dafny.MultiSet.fromElements(_603_v17);
          let _607_v21;
          _607_v21 = _dafny.Map.Empty.slice().updateUnsafe(_606_v20,p2);
          let _608_v22;
          _608_v22 = _dafny.MultiSet.fromElements(new BigNumber((_602_v16).length), (new BigNumber((_604_v18).length)).multipliedBy(p1), p1, _module.__default.safeDivisionInt(new BigNumber((_605_v19).length), new BigNumber((_607_v21).length)), p1);
          let _609_v23;
          _609_v23 = _module.D0.create_DC2(p1, p2);
          let _610_v24;
          _610_v24 = _dafny.Map.Empty.slice().updateUnsafe((_609_v23).dtor_cf6,p2);
          let _rhs131 = new BigNumber((((p2) ? (_590_v4) : (_dafny.Seq.Concat(_590_v4, _module.__default.fm22(p2, globalState))))).length);
          let _rhs132 = _module.__default.fm26(p1, (((_610_v24).contains(p2)) ? ((_610_v24).get(p2)) : (p2)), (_586_v0)[_module.__default.safeIndex(new BigNumber((_608_v22).cardinality()), new BigNumber((_586_v0).length))], globalState);
          let _rhs133 = _608_v22;
          let _rhs134 = true;
          let _rhs135 = p1;
          let _lhs127 = globalState;
          let _lhs128 = globalState;
          let _lhs129 = globalState;
          _lhs127.f10 = _rhs131;
          _598_v12 = _rhs132;
          _608_v22 = _rhs133;
          _lhs128.f13 = _rhs134;
          _lhs129.f10 = _rhs135;
          let _611_v25;
          _611_v25 = _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.of(_603_v17), _dafny.Seq.update(_586_v0, _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_586_v0).length)), _603_v17)));
          let _612_v26;
          _612_v26 = _dafny.Map.Empty.slice().updateUnsafe(_605_v19,p2);
          let _rhs136 = _module.__default.fm27(_612_v26, globalState);
          let _rhs137 = p2;
          let _lhs130 = globalState;
          _611_v25 = _rhs136;
          _lhs130.f22 = _rhs137;
          let _613_v27;
          let _nw125 = Array((new BigNumber(27)).toNumber());
          _nw125[(_dafny.ZERO).toNumber()] = p2;
          _nw125[(_dafny.ONE).toNumber()] = p2;
          _nw125[(new BigNumber(2)).toNumber()] = p2;
          _nw125[(new BigNumber(3)).toNumber()] = p2;
          _nw125[(new BigNumber(4)).toNumber()] = p2;
          _nw125[(new BigNumber(5)).toNumber()] = p2;
          _nw125[(new BigNumber(6)).toNumber()] = p2;
          _nw125[(new BigNumber(7)).toNumber()] = p2;
          _nw125[(new BigNumber(8)).toNumber()] = p2;
          _nw125[(new BigNumber(9)).toNumber()] = false;
          _nw125[(new BigNumber(10)).toNumber()] = p2;
          _nw125[(new BigNumber(11)).toNumber()] = (_this).fm21(p1, globalState);
          _nw125[(new BigNumber(12)).toNumber()] = p2;
          _nw125[(new BigNumber(13)).toNumber()] = p2;
          _nw125[(new BigNumber(14)).toNumber()] = false;
          _nw125[(new BigNumber(15)).toNumber()] = p2;
          _nw125[(new BigNumber(16)).toNumber()] = p2;
          _nw125[(new BigNumber(17)).toNumber()] = p2;
          _nw125[(new BigNumber(18)).toNumber()] = p2;
          _nw125[(new BigNumber(19)).toNumber()] = !(p2);
          _nw125[(new BigNumber(20)).toNumber()] = p2;
          _nw125[(new BigNumber(21)).toNumber()] = !(false);
          _nw125[(new BigNumber(22)).toNumber()] = p2;
          _nw125[(new BigNumber(23)).toNumber()] = p2;
          _nw125[(new BigNumber(24)).toNumber()] = p2;
          _nw125[(new BigNumber(25)).toNumber()] = (_590_v4)[_module.__default.safeIndex(p1, new BigNumber((_590_v4).length))];
          _nw125[(new BigNumber(26)).toNumber()] = p2;
          _613_v27 = _nw125;
          let _614_v28;
          _614_v28 = _dafny.Seq.of(_613_v27, _613_v27);
          let _615_v29;
          _615_v29 = _dafny.Map.Empty.slice().updateUnsafe((_614_v28)[_module.__default.safeIndex(p1, new BigNumber((_614_v28).length))],new BigNumber((_586_v0).length));
          let _616_v30;
          _616_v30 = _dafny.Map.Empty.slice().updateUnsafe(_615_v29,_dafny.Seq.contains(_586_v0, _603_v17));
          _616_v30 = (_616_v30).update(_615_v29, p2);
          (globalState).f7 = p2;
          let _index134 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_589_v3).length));
          (_589_v3)[_index134] = _613_v27;
          let _617_v31;
          _617_v31 = _module.D4.create_DC14(p2);
          let _index135 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_589_v3).length));
          let _rhs138 = _613_v27;
          let _rhs139 = p2;
          let _rhs140 = _617_v31;
          let _lhs131 = _589_v3;
          let _lhs132 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_589_v3).length));
          let _lhs133 = globalState;
          _lhs131[_lhs132] = _rhs138;
          _lhs133.f22 = _rhs139;
          _617_v31 = _rhs140;
        }
        (globalState).f10 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_590_v4, _module.__default.fm22(p2, globalState)), _590_v4)).length);
      } else {
        let _618_v32;
        _618_v32 = _dafny.Set.fromElements(_586_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(303)), function (_619_i0) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }));
        (globalState).f7 = (_this).fm21(new BigNumber((_618_v32).length), globalState);
        _586_v0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-309)), function (_620_i1) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        });
        let _621_v33;
        _621_v33 = _dafny.Set.fromElements(p1, p1, p1, p1);
        let _622_v34;
        let _nw126 = Array((new BigNumber(19)).toNumber());
        _nw126[(_dafny.ZERO).toNumber()] = false;
        _nw126[(_dafny.ONE).toNumber()] = false;
        _nw126[(new BigNumber(2)).toNumber()] = (_this).fm21(new BigNumber(-719), globalState);
        _nw126[(new BigNumber(3)).toNumber()] = (_621_v33).IsSubsetOf(_621_v33);
        _nw126[(new BigNumber(4)).toNumber()] = p2;
        _nw126[(new BigNumber(5)).toNumber()] = p2;
        _nw126[(new BigNumber(6)).toNumber()] = false;
        _nw126[(new BigNumber(7)).toNumber()] = (new BigNumber((_dafny.Seq.UnicodeFromString("iqstee")).length)).isLessThan(new BigNumber((_module.__default.fm28(p2, globalState)).length));
        _nw126[(new BigNumber(8)).toNumber()] = !(true);
        _nw126[(new BigNumber(9)).toNumber()] = p2;
        _nw126[(new BigNumber(10)).toNumber()] = p2;
        _nw126[(new BigNumber(11)).toNumber()] = (p2) === (p2);
        _nw126[(new BigNumber(12)).toNumber()] = true;
        _nw126[(new BigNumber(13)).toNumber()] = p2;
        _nw126[(new BigNumber(14)).toNumber()] = p2;
        _nw126[(new BigNumber(15)).toNumber()] = !(false);
        _nw126[(new BigNumber(16)).toNumber()] = p2;
        _nw126[(new BigNumber(17)).toNumber()] = p2;
        _nw126[(new BigNumber(18)).toNumber()] = p2;
        _622_v34 = _nw126;
        let _index136 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_622_v34).length));
        (_622_v34)[_index136] = (p1).isLessThanOrEqualTo(p1);
        let _index137 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_622_v34).length));
        (_622_v34)[_index137] = p2;
        (globalState).f13 = (_622_v34)[_module.__default.safeIndex(new BigNumber(839), new BigNumber((_622_v34).length))];
        let _623_v35;
        _623_v35 = _dafny.Set.fromElements((_622_v34)[_module.__default.safeIndex(new BigNumber(839), new BigNumber((_622_v34).length))]);
        let _624_v36;
        _624_v36 = _dafny.Seq.of((_622_v34)[_module.__default.safeIndex(new BigNumber(839), new BigNumber((_622_v34).length))]);
        let _625_v37;
        _625_v37 = _dafny.MultiSet.fromElements(p1);
        let _626_v38;
        _626_v38 = _dafny.Set.fromElements(!(p2), (_623_v35).IsProperSubsetOf(_dafny.Set.fromElements(p2)), ((!(p2)) ? ((_624_v36)[_module.__default.safeIndex(p1, new BigNumber((_624_v36).length))]) : (!(p2))), (_622_v34)[_module.__default.safeIndex(new BigNumber(839), new BigNumber((_622_v34).length))], (_625_v37).IsProperSubsetOf(_625_v37));
        _623_v35 = _626_v38;
      }
      let _627_v39;
      let _nw127 = Array((new BigNumber(13)).toNumber()).fill(false);
      _627_v39 = _nw127;
      let _628_v40;
      _628_v40 = _dafny.Seq.of(p2);
      let _629_v41;
      _629_v41 = _module.D1.create_DC6(p1, _dafny.Seq.of(_628_v40), p1, p2);
      let _rhs141 = _627_v39;
      let _rhs142 = p1;
      let _rhs143 = _module.__default.fm29(globalState);
      let _lhs134 = globalState;
      _627_v39 = _rhs141;
      _lhs134.f10 = _rhs142;
      _629_v41 = _rhs143;
      let _630_v42;
      let _nw128 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
      _630_v42 = _nw128;
      _630_v42 = _630_v42;
      let _631_v43;
      let _nw129 = Array((new BigNumber(27)).toNumber()).fill([]);
      _631_v43 = _nw129;
      let _index138 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_631_v43).length));
      (_631_v43)[_index138] = _627_v39;
      let _index139 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_631_v43).length));
      let _nw130 = Array((new BigNumber(14)).toNumber()).fill(false);
      (_631_v43)[_index139] = _nw130;
      (globalState).f10 = p1;
      let _632_v44;
      _632_v44 = _module.D0.create_DC1(new BigNumber(914), p1, p1, (_631_v43)[_module.__default.safeIndex(new BigNumber(960), new BigNumber((_631_v43).length))]);
      r0 = ((!((p2) === (p2))) ? (_632_v44) : (_632_v44));
      r1 = (_628_v40)[_module.__default.safeIndex(p1, new BigNumber((_628_v40).length))];
      return [r0, r1];
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let _633_v0;
      _633_v0 = _dafny.Seq.of(p2);
      _633_v0 = _633_v0;
      let _634_v1;
      let _nw131 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
      _634_v1 = _nw131;
      let _index140 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length));
      (_634_v1)[_index140] = p2;
      let _index141 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length));
      (_634_v1)[_index141] = p1;
      let _635_v2;
      _635_v2 = false;
      let _index142 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length));
      (_634_v1)[_index142] = ((!(_635_v2)) ? (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(835)), ((_636_p1) => function (_637_i0) {
        return (_dafny.ZERO).minus(_636_p1);
      })(p1))).length)) : (p2));
      let _638_v3;
      _638_v3 = _dafny.Seq.of(_635_v2);
      (globalState).f10 = ((_634_v1)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length))]).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_638_v3, _module.__default.safeIndex(p3, new BigNumber((_638_v3).length)), _635_v2)).length)));
      if ((_635_v2) === (_635_v2)) {
        let _639_v4;
        _639_v4 = _dafny.Seq.UnicodeFromString("ffxpkyq");
        let _index143 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length));
        let _rhs144 = _module.__default.safeModuloInt((_dafny.ZERO).minus((new BigNumber((_dafny.Seq.of(_635_v2)).length)).multipliedBy(p0)), (_634_v1)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length))]);
        let _rhs145 = ((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_633_v0).length), p3))).plus(_module.__default.safeDivisionInt(new BigNumber((_639_v4).length), (_634_v1)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length))]));
        let _rhs146 = p2;
        let _lhs135 = globalState;
        let _lhs136 = globalState;
        let _lhs137 = _634_v1;
        let _lhs138 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length));
        _lhs135.f10 = _rhs144;
        _lhs136.f10 = _rhs145;
        _lhs137[_lhs138] = _rhs146;
        let _640_v5;
        let _nw132 = new _module.C0();
        _nw132.__ctor(!_dafny.areEqual(_639_v4, _639_v4));
        _640_v5 = _nw132;
        let _641_v6;
        _641_v6 = _module.D3.create_DC12((_640_v5).f32);
        let _642_v7;
        _642_v7 = _dafny.MultiSet.fromElements(p0);
        let _643_v8;
        _643_v8 = _module.D1.create_DC5(p1, (_module.D3.create_DC12((_640_v5).f32)).dtor_cf23, !((_641_v6).dtor_cf23), _dafny.Map.Empty.slice().updateUnsafe(_642_v7,_633_v0));
        _643_v8 = _643_v8;
        let _644_v9;
        _644_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-298),_640_v5);
        (globalState).f10 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(new BigNumber((_644_v9).length), (_dafny.ZERO).minus(p3)), (p3).plus(p2));
        let _645_v10;
        _645_v10 = new _dafny.CodePoint('n'.codePointAt(0));
        (globalState).f19 = _645_v10;
      } else {
        let _index144 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length));
        let _rhs147 = (_this).fm21((_dafny.ZERO).minus(p3), globalState);
        let _rhs148 = p1;
        let _rhs149 = !((p0).isLessThan(p3));
        let _lhs139 = globalState;
        let _lhs140 = _634_v1;
        let _lhs141 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length));
        let _lhs142 = globalState;
        _lhs139.f22 = _rhs147;
        _lhs140[_lhs141] = _rhs148;
        _lhs142.f7 = _rhs149;
        let _646_v11;
        let _nw133 = new _module.C0();
        _nw133.__ctor((_this).fm21(p2, globalState));
        _646_v11 = _nw133;
        let _647_v12;
        _647_v12 = _dafny.MultiSet.fromElements((_646_v11).f32);
        let _648_v13;
        _648_v13 = _dafny.Map.Empty.slice().updateUnsafe(!(!((_646_v11).f32)),!(!((_646_v11).f32)));
        (globalState).f10 = (((_646_v11).f32) ? ((((_647_v12).contains(false)) ? ((_647_v12).get(false)) : (new BigNumber((_648_v13).length)))) : ((_634_v1)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length))]));
        let _649_v14;
        let _nw134 = new _module.C0();
        _nw134.__ctor(_635_v2);
        _649_v14 = _nw134;
        let _650_v15;
        _650_v15 = _dafny.Seq.UnicodeFromString("lroq");
        let _651_v16;
        _651_v16 = _dafny.Map.Empty.slice().updateUnsafe((_634_v1)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length))],p3);
        let _652_v17;
        _652_v17 = _module.D2.create_DC7(_651_v16);
        let _653_v18;
        _653_v18 = new _dafny.CodePoint('x'.codePointAt(0));
        let _654_v19;
        _654_v19 = _module.D3.create_DC11(p3);
        let _655_v20;
        _655_v20 = _dafny.Seq.of(_650_v15);
        let _656_v21;
        let _nw135 = Array((new BigNumber(28)).toNumber());
        _nw135[(_dafny.ZERO).toNumber()] = _650_v15;
        _nw135[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(663)), function (_657_i1) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        });
        _nw135[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm23(_652_v17, p2, (_dafny.ZERO).minus((_634_v1)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length))]), globalState), _dafny.Seq.UnicodeFromString("dtnxh"));
        _nw135[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(620)), function (_658_i2) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        });
        _nw135[(new BigNumber(4)).toNumber()] = _650_v15;
        _nw135[(new BigNumber(5)).toNumber()] = _650_v15;
        _nw135[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("vhcdvwyj");
        _nw135[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("chvp");
        _nw135[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_650_v15, _650_v15);
        _nw135[(new BigNumber(9)).toNumber()] = _650_v15;
        _nw135[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(44)), function (_659_i3) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        });
        _nw135[(new BigNumber(11)).toNumber()] = _650_v15;
        _nw135[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_650_v15, _module.__default.safeIndex(new BigNumber(781), new BigNumber((_650_v15).length)), _653_v18);
        _nw135[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_650_v15, _dafny.Seq.UnicodeFromString("xjxlm"));
        _nw135[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kxppprism"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(677)), ((_660_v18) => function (_661_i4) {
          return _660_v18;
        })(_653_v18)));
        _nw135[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("qghd");
        _nw135[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ajvtl"), _650_v15);
        _nw135[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(137)), ((_662_v18) => function (_663_i5) {
          return _662_v18;
        })(_653_v18));
        _nw135[(new BigNumber(18)).toNumber()] = _650_v15;
        _nw135[(new BigNumber(19)).toNumber()] = _module.__default.fm23(_module.D2.create_DC7(_651_v16), new BigNumber(422), (_634_v1)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length))], globalState);
        _nw135[(new BigNumber(20)).toNumber()] = _650_v15;
        _nw135[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(785)), ((_664_v18) => function (_665_i6) {
          return _664_v18;
        })(_653_v18)), _650_v15);
        _nw135[(new BigNumber(22)).toNumber()] = _650_v15;
        _nw135[(new BigNumber(23)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-926)), ((_666_v18) => function (_667_i7) {
          return _666_v18;
        })(_653_v18));
        _nw135[(new BigNumber(24)).toNumber()] = _650_v15;
        _nw135[(new BigNumber(25)).toNumber()] = _dafny.Seq.UnicodeFromString("ytkqvx");
        _nw135[(new BigNumber(26)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_650_v15, _module.__default.safeIndex((_654_v19).dtor_cf22, new BigNumber((_650_v15).length)), _653_v18), _650_v15);
        _nw135[(new BigNumber(27)).toNumber()] = (_655_v20)[_module.__default.safeIndex((_634_v1)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length))], new BigNumber((_655_v20).length))];
        _656_v21 = _nw135;
        let _index145 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_656_v21).length));
        (_656_v21)[_index145] = _650_v15;
        let _668_v22;
        let _nw136 = Array((new BigNumber(28)).toNumber()).fill(_module.D4.Default());
        _668_v22 = _nw136;
        let _669_v23;
        _669_v23 = _module.D4.create_DC13(_653_v18);
        let _670_v24;
        _670_v24 = _module.D4.create_DC15(_669_v23);
        let _index146 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_668_v22).length));
        (_668_v22)[_index146] = _670_v24;
        let _index147 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_656_v21).length));
        let _index148 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length));
        let _index149 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_668_v22).length));
        let _rhs150 = _dafny.Seq.UnicodeFromString("akub");
        let _rhs151 = new BigNumber((_638_v3).length);
        let _rhs152 = _670_v24;
        let _rhs153 = (_646_v11).f32;
        let _lhs143 = _656_v21;
        let _lhs144 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_656_v21).length));
        let _lhs145 = _634_v1;
        let _lhs146 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_634_v1).length));
        let _lhs147 = _668_v22;
        let _lhs148 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_668_v22).length));
        let _lhs149 = globalState;
        _lhs143[_lhs144] = _rhs150;
        _lhs145[_lhs146] = _rhs151;
        _lhs147[_lhs148] = _rhs152;
        _lhs149.f7 = _rhs153;
      }
      let _671_v25;
      let _nw137 = Array((new BigNumber(14)).toNumber()).fill(false);
      _671_v25 = _nw137;
      let _672_v26;
      _672_v26 = _dafny.Seq.UnicodeFromString("cfgaepsi");
      let _rhs154 = _671_v25;
      let _rhs155 = _672_v26;
      let _rhs156 = _633_v0;
      let _rhs157 = p3;
      let _lhs150 = globalState;
      _671_v25 = _rhs154;
      _672_v26 = _rhs155;
      _633_v0 = _rhs156;
      _lhs150.f10 = _rhs157;
      let _673_v27;
      _673_v27 = new _dafny.CodePoint('m'.codePointAt(0));
      r0 = _673_v27;
      return r0;
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let _674_v0;
      _674_v0 = _dafny.Set.fromElements(false);
      let _675_v1;
      _675_v1 = true;
      let _676_v2;
      let _nw138 = Array((new BigNumber(6)).toNumber());
      _nw138[(_dafny.ZERO).toNumber()] = _675_v1;
      _nw138[(_dafny.ONE).toNumber()] = _675_v1;
      _nw138[(new BigNumber(2)).toNumber()] = _675_v1;
      _nw138[(new BigNumber(3)).toNumber()] = _675_v1;
      _nw138[(new BigNumber(4)).toNumber()] = _675_v1;
      _nw138[(new BigNumber(5)).toNumber()] = _675_v1;
      _676_v2 = _nw138;
      let _pat_let_tv11 = _676_v2;
      let _pat_let_tv12 = _676_v2;
      let _pat_let_tv13 = _675_v1;
      let _source12 = function (_pat_let8_0) {
        return function (_677_dt__update__tmp_h0) {
          return function (_pat_let9_0) {
            return function (_678_dt__update_hcf4_h0) {
              return _module.D0.create_DC1((_677_dt__update__tmp_h0).dtor_cf1, (_677_dt__update__tmp_h0).dtor_cf2, (_677_dt__update__tmp_h0).dtor_cf3, _678_dt__update_hcf4_h0);
            }(_pat_let9_0);
          }(((_pat_let_tv13) ? (_pat_let_tv11) : (_pat_let_tv12)));
        }(_pat_let8_0);
      }(_module.D0.create_DC1(new BigNumber((_674_v0).length), new BigNumber(-556), p0, _676_v2));
      if (_source12.is_DC1) {
        let _679___mcc_h0 = (_source12).cf1;
        let _680___mcc_h1 = (_source12).cf2;
        let _681___mcc_h2 = (_source12).cf3;
        let _682___mcc_h3 = (_source12).cf4;
        let _683_cf4 = _682___mcc_h3;
        let _684_cf3 = _681___mcc_h2;
        let _685_cf2 = _680___mcc_h1;
        let _686_cf1 = _679___mcc_h0;
        let _687_v3;
        let _init14 = ((_688_v0) => function (_689_i0) {
          return (_688_v0).Union(_688_v0);
        })(_674_v0);
        let _nw139 = Array((new BigNumber(25)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw139.length); _i0_14++) {
          _nw139[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _687_v3 = _nw139;
        _687_v3 = _687_v3;
        let _690_v4;
        _690_v4 = _dafny.Seq.UnicodeFromString("gteyicgu");
        let _691_v5;
        _691_v5 = _module.D0.create_DC0(_690_v4);
        let _692_v6;
        _692_v6 = _dafny.Set.fromElements(_691_v5);
        _685_cf2 = new BigNumber((_692_v6).length);
        let _693_v7;
        _693_v7 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_690_v4, _690_v4));
        _693_v7 = _dafny.MultiSet.fromElements(_690_v4, _690_v4);
        let _694_v8;
        let _nw140 = Array((new BigNumber(4)).toNumber());
        _nw140[(_dafny.ZERO).toNumber()] = _686_cf1;
        _nw140[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.of(p0, _684_cf3, new BigNumber((_690_v4).length), new BigNumber(21))).length)).plus(_686_cf1));
        _nw140[(new BigNumber(2)).toNumber()] = _685_cf2;
        _nw140[(new BigNumber(3)).toNumber()] = (new BigNumber(500)).plus(_686_cf1);
        _694_v8 = _nw140;
        let _index150 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_694_v8).length));
        (_694_v8)[_index150] = (_dafny.ZERO).minus(p1);
        let _index151 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_694_v8).length));
        (_694_v8)[_index151] = _686_cf1;
      } else if (_source12.is_DC2) {
        let _695___mcc_h4 = (_source12).cf5;
        let _696___mcc_h5 = (_source12).cf6;
        let _697_cf6 = _696___mcc_h5;
        let _698_cf5 = _695___mcc_h4;
        let _699_v9;
        _699_v9 = _dafny.Seq.UnicodeFromString("rpgfswiru");
        (globalState).f13 = !(_dafny.Seq.IsPrefixOf(_699_v9, _699_v9)) || ((_698_cf5).isEqualTo(new BigNumber((_module.__default.fm22(_697_cf6, globalState)).length)));
        let _700_v10;
        let _init15 = ((_701_cf5) => function (_702_i1) {
          return _module.__default.safeModuloInt(_702_i1, _701_cf5);
        })(_698_cf5);
        let _nw141 = Array((new BigNumber(5)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw141.length); _i0_15++) {
          _nw141[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _700_v10 = _nw141;
        let _703_v11;
        _703_v11 = _dafny.Seq.of(_698_cf5);
        let _index152 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_700_v10).length));
        (_700_v10)[_index152] = (_703_v11)[_module.__default.safeIndex((_this).fm2(globalState), new BigNumber((_703_v11).length))];
        let _index153 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_700_v10).length));
        (_700_v10)[_index153] = p1;
        r0 = _700_v10;
        _698_cf5 = _698_cf5;
      } else if (_source12.is_DC0) {
        let _704___mcc_h6 = (_source12).cf0;
        let _705_cf0 = _704___mcc_h6;
        let _index154 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_676_v2).length));
        (_676_v2)[_index154] = !(_675_v1);
        let _706_v12;
        _706_v12 = _module.D0.create_DC2(new BigNumber(527), _675_v1);
        let _pat_let_tv14 = _675_v1;
        let _index155 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_676_v2).length));
        (_676_v2)[_index155] = ((_675_v1) || ((function (_pat_let10_0) {
          return function (_707_dt__update__tmp_h1) {
            return function (_pat_let11_0) {
              return function (_708_dt__update_hcf6_h0) {
                return _module.D0.create_DC2((_707_dt__update__tmp_h1).dtor_cf5, _708_dt__update_hcf6_h0);
              }(_pat_let11_0);
            }(_pat_let_tv14);
          }(_pat_let10_0);
        }(_706_v12)).dtor_cf6)) || (false);
        let _709_v13;
        let _nw142 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _709_v13 = _nw142;
        r0 = _709_v13;
        let _710_v14;
        _710_v14 = _module.D3.create_DC11(p1);
        (globalState).f10 = (_710_v14).dtor_cf22;
        let _711_v15;
        _711_v15 = _dafny.Map.Empty.slice().updateUnsafe(_710_v14,!(_675_v1));
        let _712_v16;
        _712_v16 = _dafny.Seq.of(_711_v15);
        _712_v16 = _712_v16;
      } else {
        let _713___mcc_h7 = (_source12).cf7;
        let _714_cf7 = _713___mcc_h7;
        let _715_v17;
        _715_v17 = new _dafny.CodePoint('a'.codePointAt(0));
        let _716_v18;
        _716_v18 = _dafny.Seq.of(_715_v17, _715_v17, new _dafny.CodePoint('v'.codePointAt(0)), _715_v17, new _dafny.CodePoint('g'.codePointAt(0)));
        _716_v18 = _dafny.Seq.Concat(_716_v18, _dafny.Seq.of(_715_v17, _715_v17, _715_v17));
        let _717_v19;
        _717_v19 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), ((_718_v17) => function (_719_i2) {
          return _718_v17;
        })(_715_v17)));
        _716_v18 = (_717_v19)[_module.__default.safeIndex(new BigNumber((((_675_v1) ? (_674_v0) : (_674_v0))).length), new BigNumber((_717_v19).length))];
        let _720_v20;
        _720_v20 = _module.D0.create_DC0(_716_v18);
        let _pat_let_tv15 = _716_v18;
        let _pat_let_tv16 = _716_v18;
        let _source13 = function (_pat_let12_0) {
          return function (_721_dt__update__tmp_h2) {
            return function (_pat_let13_0) {
              return function (_722_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_722_dt__update_hcf0_h0);
              }(_pat_let13_0);
            }(_dafny.Seq.Concat(_pat_let_tv15, _pat_let_tv16));
          }(_pat_let12_0);
        }(_720_v20);
        if (_source13.is_DC1) {
          let _723___mcc_h8 = (_source13).cf1;
          let _724___mcc_h9 = (_source13).cf2;
          let _725___mcc_h10 = (_source13).cf3;
          let _726___mcc_h11 = (_source13).cf4;
          let _727_cf4 = _726___mcc_h11;
          let _728_cf3 = _725___mcc_h10;
          let _729_cf2 = _724___mcc_h9;
          let _730_cf1 = _723___mcc_h8;
          let _731_v21;
          let _nw143 = Array((new BigNumber(26)).toNumber()).fill([]);
          _731_v21 = _nw143;
          let _732_v22;
          _732_v22 = _dafny.Set.fromElements(_728_cf3, _729_cf2);
          let _rhs158 = ((_675_v1) ? (_dafny.Seq.UnicodeFromString("l")) : (_716_v18));
          let _rhs159 = _675_v1;
          let _rhs160 = _731_v21;
          let _rhs161 = ((_728_cf3).isLessThan(new BigNumber((_732_v22).length))) === (_675_v1);
          let _lhs151 = globalState;
          _716_v18 = _rhs158;
          _675_v1 = _rhs159;
          _731_v21 = _rhs160;
          _lhs151.f7 = _rhs161;
          let _733_v23;
          _733_v23 = _module.D3.create_DC11(p1);
          _733_v23 = _733_v23;
          let _734_v24;
          _734_v24 = _module.D0.create_DC2(p1, _675_v1);
          let _735_v25;
          _735_v25 = _module.D0.create_DC3(_module.D0.create_DC3(_734_v24));
          let _736_v26;
          let _nw144 = Array((new BigNumber(2)).toNumber());
          _nw144[(_dafny.ZERO).toNumber()] = _735_v25;
          _nw144[(_dafny.ONE).toNumber()] = ((_675_v1) ? (_735_v25) : (_735_v25));
          _736_v26 = _nw144;
          let _index156 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_736_v26).length));
          (_736_v26)[_index156] = ((_675_v1) ? (_735_v25) : (_735_v25));
          let _737_v27;
          _737_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_729_cf2,_675_v1),_675_v1);
          let _738_v28;
          _738_v28 = _module.D0.create_DC2(p0, _675_v1);
          let _index157 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_736_v26).length));
          let _rhs162 = (_717_v19)[_module.__default.safeIndex((_dafny.ZERO).minus(_728_cf3), new BigNumber((_717_v19).length))];
          let _rhs163 = _735_v25;
          let _rhs164 = ((_675_v1) ? (_675_v1) : ((_738_v28).dtor_cf6));
          let _rhs165 = _737_v27;
          let _lhs152 = _736_v26;
          let _lhs153 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_736_v26).length));
          let _lhs154 = globalState;
          _716_v18 = _rhs162;
          _lhs152[_lhs153] = _rhs163;
          _lhs154.f13 = _rhs164;
          _737_v27 = _rhs165;
          let _739_v29;
          let _nw145 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _739_v29 = _nw145;
          let _740_v30;
          _740_v30 = _739_v29;
          let _741_v31;
          _741_v31 = _dafny.Seq.of((_740_v30), _739_v29, _739_v29);
          let _rhs166 = (_741_v31)[_module.__default.safeIndex(p1, new BigNumber((_741_v31).length))];
          let _rhs167 = _676_v2;
          let _rhs168 = _675_v1;
          let _rhs169 = _730_cf1;
          r0 = _rhs166;
          _676_v2 = _rhs167;
          r1 = _rhs168;
          _729_cf2 = _rhs169;
        } else if (_source13.is_DC2) {
          let _742___mcc_h12 = (_source13).cf5;
          let _743___mcc_h13 = (_source13).cf6;
          let _744_cf6 = _743___mcc_h13;
          let _745_cf5 = _742___mcc_h12;
          (globalState).f13 = _744_cf6;
          let _index158 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_676_v2).length));
          (_676_v2)[_index158] = !(_dafny.areEqual(_716_v18, _716_v18));
          let _746_v32;
          _746_v32 = _dafny.Map.Empty.slice().updateUnsafe(_675_v1,!(_675_v1));
          let _index159 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_676_v2).length));
          let _rhs170 = _module.__default.safeModuloInt(p1, p1);
          let _rhs171 = _744_cf6;
          let _rhs172 = (((_module.__default.fm30(true, globalState)).equals(_746_v32)) ? (_module.__default.safeDivisionInt(p1, p0)) : ((_dafny.ZERO).minus(_745_cf5)));
          let _lhs155 = globalState;
          let _lhs156 = _676_v2;
          let _lhs157 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_676_v2).length));
          let _lhs158 = globalState;
          _lhs155.f10 = _rhs170;
          _lhs156[_lhs157] = _rhs171;
          _lhs158.f10 = _rhs172;
          let _747_v33;
          _747_v33 = _dafny.Seq.of((_676_v2)[_module.__default.safeIndex(new BigNumber(413), new BigNumber((_676_v2).length))]);
          let _748_v34;
          _748_v34 = _dafny.Seq.of((_747_v33)[_module.__default.safeIndex(p0, new BigNumber((_747_v33).length))], (_676_v2)[_module.__default.safeIndex(new BigNumber(413), new BigNumber((_676_v2).length))]);
          _748_v34 = _747_v33;
          let _749_v35;
          let _nw146 = Array((new BigNumber(19)).toNumber()).fill(false);
          _749_v35 = _nw146;
          let _750_v36;
          _750_v36 = _dafny.Seq.of(_749_v35);
          r1 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_750_v36, _750_v36), _750_v36);
        } else if (_source13.is_DC0) {
          let _751___mcc_h14 = (_source13).cf0;
          let _752_cf0 = _751___mcc_h14;
          let _753_v37;
          _753_v37 = _dafny.MultiSet.fromElements(_675_v1);
          let _754_v38;
          _754_v38 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Map.Empty.slice().updateUnsafe(p0,_675_v1));
          let _755_v39;
          _755_v39 = _dafny.Seq.of(_module.__default.fm0(p1, globalState), (_this).fm2(globalState), new BigNumber((_753_v37).cardinality()), p0, new BigNumber((_754_v38).length));
          let _756_v40;
          let _out3;
          _out3 = (_this).m6(_module.__default.fm0(p0, globalState), (p0).multipliedBy(p1), (_755_v39)[_module.__default.safeIndex(p0, new BigNumber((_755_v39).length))], p0, globalState);
          _756_v40 = _out3;
          (globalState).f22 = _675_v1;
          let _757_v41;
          let _init16 = ((_758_p1) => function (_759_i3) {
            return (_759_i3).multipliedBy(_758_p1);
          })(p1);
          let _nw147 = Array((new BigNumber(24)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw147.length); _i0_16++) {
            _nw147[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _757_v41 = _nw147;
          let _760_v42;
          _760_v42 = _dafny.Map.Empty.slice().updateUnsafe(_675_v1,_757_v41);
          _760_v42 = (_760_v42).update(_675_v1, _757_v41);
          (globalState).f13 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(214)), ((_761_v17) => function (_762_i4) {
            return _761_v17;
          })(_715_v17)), _716_v18), (_720_v20).dtor_cf0);
        } else {
          let _763___mcc_h15 = (_source13).cf7;
          let _764_cf7 = _763___mcc_h15;
          let _765_v43;
          let _nw148 = Array((new BigNumber(8)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _765_v43 = _nw148;
          let _766_v44;
          _766_v44 = _dafny.Seq.of(_765_v43);
          let _rhs173 = new _dafny.CodePoint('p'.codePointAt(0));
          let _rhs174 = _675_v1;
          let _rhs175 = (_766_v44)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_766_v44).length))];
          let _lhs159 = globalState;
          _715_v17 = _rhs173;
          _lhs159.f22 = _rhs174;
          _765_v43 = _rhs175;
          _676_v2 = _676_v2;
          (globalState).f10 = _module.__default.safeDivisionInt(p0, (_dafny.ZERO).minus(p1));
          (globalState).f22 = _675_v1;
        }
        let _767_v45;
        let _init17 = ((_768_v1) => function (_769_i5) {
          return (((_module.D2.create_DC8(_768_v1, _768_v1)).dtor_cf19) ? (_module.D4.create_DC13(new _dafny.CodePoint('i'.codePointAt(0)))) : (_module.D4.create_DC13(new _dafny.CodePoint('o'.codePointAt(0)))));
        })(_675_v1);
        let _nw149 = Array((new BigNumber(4)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw149.length); _i0_17++) {
          _nw149[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _767_v45 = _nw149;
        let _770_v46;
        _770_v46 = _module.D4.create_DC13(_715_v17);
        let _pat_let_tv17 = _715_v17;
        let _index160 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_767_v45).length));
        (_767_v45)[_index160] = function (_pat_let14_0) {
          return function (_771_dt__update__tmp_h3) {
            return function (_pat_let15_0) {
              return function (_772_dt__update_hcf24_h0) {
                return _module.D4.create_DC13(_772_dt__update_hcf24_h0);
              }(_pat_let15_0);
            }(_pat_let_tv17);
          }(_pat_let14_0);
        }(_770_v46);
        let _pat_let_tv18 = _715_v17;
        let _index161 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_767_v45).length));
        (_767_v45)[_index161] = function (_pat_let16_0) {
          return function (_773_dt__update__tmp_h4) {
            return function (_pat_let17_0) {
              return function (_774_dt__update_hcf24_h1) {
                return _module.D4.create_DC13(_774_dt__update_hcf24_h1);
              }(_pat_let17_0);
            }(_pat_let_tv18);
          }(_pat_let16_0);
        }(_770_v46);
      }
      let _775_v47;
      _775_v47 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).fm21(p1, globalState)),_676_v2);
      _775_v47 = (_775_v47).update((((_this).fm21(p0, globalState)) ? (_675_v1) : (_675_v1)), _676_v2);
      let _776_v48;
      _776_v48 = _dafny.Seq.of(_675_v1);
      let _hi2 = p1;
      for (let _777_i6 = new BigNumber((_776_v48).length); _777_i6.isLessThan(_hi2); _777_i6 = _777_i6.plus(_dafny.ONE)) {
        let _index162 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length));
        (_676_v2)[_index162] = true;
        let _778_v49;
        _778_v49 = _dafny.Map.Empty.slice().updateUnsafe(_675_v1,p0);
        let _779_v50;
        _779_v50 = _dafny.Seq.of(_777_i6);
        let _780_v51;
        _780_v51 = _dafny.Seq.of((((_778_v49).contains(!(_675_v1))) ? ((_778_v49).get(!(_675_v1))) : (p1)), _777_i6, (_779_v50)[_module.__default.safeIndex(p0, new BigNumber((_779_v50).length))], _777_i6, _777_i6);
        let _index163 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length));
        (_676_v2)[_index163] = (_777_i6).isLessThan((_780_v51)[_module.__default.safeIndex(_777_i6, new BigNumber((_780_v51).length))]);
        (globalState).f22 = !((_776_v48)[_module.__default.safeIndex(p1, new BigNumber((_776_v48).length))]);
        let _781_v52;
        _781_v52 = _module.D3.create_DC11(p1);
        let _source14 = _781_v52;
        if (_source14.is_DC11) {
          let _782___mcc_h16 = (_source14).cf22;
          let _783_cf22 = _782___mcc_h16;
          let _784_v53;
          _784_v53 = _dafny.Seq.UnicodeFromString("lkekff");
          _784_v53 = _dafny.Seq.Concat(_dafny.Seq.Concat(_784_v53, _784_v53), _784_v53);
          let _785_v54;
          let _nw150 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _785_v54 = _nw150;
          let _index164 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_785_v54).length));
          (_785_v54)[_index164] = new _dafny.CodePoint('e'.codePointAt(0));
          let _786_v55;
          _786_v55 = _dafny.Set.fromElements(_783_cf22);
          let _index165 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_785_v54).length));
          (_785_v54)[_index165] = (_784_v53)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((_786_v55).Intersect(_786_v55)).length)), new BigNumber((_784_v53).length))];
          let _787_v56;
          _787_v56 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))], _675_v1),(_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))]);
          let _788_v57;
          _788_v57 = _module.D1.create_DC4(_787_v56);
          _788_v57 = _788_v57;
          let _789_v58;
          _789_v58 = _dafny.MultiSet.fromElements((_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))], (_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))], _675_v1);
          let _790_v59;
          let _out4;
          _out4 = (_this).m6(_777_i6, _783_cf22, _777_i6, _module.__default.safeModuloInt(_783_cf22, (((_789_v58).contains(_675_v1)) ? ((_789_v58).get(_675_v1)) : ((_dafny.ZERO).minus(new BigNumber(-473))))), globalState);
          _790_v59 = _out4;
        } else if (_source14.is_DC12) {
          let _791___mcc_h17 = (_source14).cf23;
          let _792_cf23 = _791___mcc_h17;
          let _793_v60;
          _793_v60 = _dafny.Seq.UnicodeFromString("bxjrgwaol");
          let _794_v61;
          let _nw151 = Array((new BigNumber(6)).toNumber());
          _nw151[(_dafny.ZERO).toNumber()] = _792_cf23;
          _nw151[(_dafny.ONE).toNumber()] = _675_v1;
          _nw151[(new BigNumber(2)).toNumber()] = _675_v1;
          _nw151[(new BigNumber(3)).toNumber()] = _792_cf23;
          _nw151[(new BigNumber(4)).toNumber()] = (_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))];
          _nw151[(new BigNumber(5)).toNumber()] = _792_cf23;
          _794_v61 = _nw151;
          let _795_v62;
          _795_v62 = _dafny.MultiSet.fromElements(false);
          let _796_v63;
          _796_v63 = _dafny.Map.Empty.slice().updateUnsafe(_781_v52,(_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))]);
          let _797_v64;
          _797_v64 = _dafny.Seq.of(_796_v63);
          let _798_v65;
          _798_v65 = _dafny.Map.Empty.slice().updateUnsafe(_792_cf23,(_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))]);
          let _799_v66;
          _799_v66 = _module.D4.create_DC14(_675_v1);
          let _800_v67;
          _800_v67 = _module.D4.create_DC15(_799_v66);
          let _801_v68;
          _801_v68 = _module.D4.create_DC15(_800_v67);
          let _802_v69;
          _802_v69 = _module.D7.create_DC20(_dafny.Seq.of(new BigNumber(590)), _794_v61, _801_v68, (_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))], p1);
          let _803_v70;
          _803_v70 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(414),(_802_v69).dtor_cf34);
          let _804_v71;
          let _nw152 = Array((new BigNumber(25)).toNumber());
          _nw152[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(new BigNumber(((_module.D7.create_DC18(_dafny.Seq.Create(_module.__default.abs(new BigNumber(933)), ((_805_v48) => function (_806_i7) {
  return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),new BigNumber((_805_v48).length));
})(_776_v48)))).dtor_cf29).length), new BigNumber((_793_v60).length));
          _nw152[(_dafny.ONE).toNumber()] = p1;
          _nw152[(new BigNumber(2)).toNumber()] = new BigNumber(698);
          _nw152[(new BigNumber(3)).toNumber()] = p0;
          _nw152[(new BigNumber(4)).toNumber()] = p0;
          _nw152[(new BigNumber(5)).toNumber()] = p1;
          _nw152[(new BigNumber(6)).toNumber()] = new BigNumber(249);
          _nw152[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), new BigNumber((_776_v48).length)));
          _nw152[(new BigNumber(8)).toNumber()] = new BigNumber(504);
          _nw152[(new BigNumber(9)).toNumber()] = p0;
          _nw152[(new BigNumber(10)).toNumber()] = p0;
          _nw152[(new BigNumber(11)).toNumber()] = (new BigNumber((_dafny.MultiSet.fromElements(_794_v61)).cardinality())).minus(new BigNumber(-328));
          _nw152[(new BigNumber(12)).toNumber()] = new BigNumber((_795_v62).cardinality());
          _nw152[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(_777_i6, (_dafny.ZERO).minus(_777_i6));
          _nw152[(new BigNumber(14)).toNumber()] = new BigNumber(((_797_v64)[_module.__default.safeIndex(new BigNumber((_795_v62).cardinality()), new BigNumber((_797_v64).length))]).length);
          _nw152[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_793_v60, _dafny.Seq.UnicodeFromString("gicvf"))).length);
          _nw152[(new BigNumber(16)).toNumber()] = p0;
          _nw152[(new BigNumber(17)).toNumber()] = new BigNumber(264);
          _nw152[(new BigNumber(18)).toNumber()] = p0;
          _nw152[(new BigNumber(19)).toNumber()] = new BigNumber(-87);
          _nw152[(new BigNumber(20)).toNumber()] = ((_675_v1) ? (_777_i6) : (_777_i6));
          _nw152[(new BigNumber(21)).toNumber()] = (_777_i6).plus(p0);
          _nw152[(new BigNumber(22)).toNumber()] = new BigNumber(487);
          _nw152[(new BigNumber(23)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_798_v65).length), new BigNumber((_803_v70).length));
          _nw152[(new BigNumber(24)).toNumber()] = _777_i6;
          _804_v71 = _nw152;
          let _index166 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_804_v71).length));
          (_804_v71)[_index166] = new BigNumber((_793_v60).length);
          let _index167 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_804_v71).length));
          let _rhs176 = _dafny.Seq.Concat(_776_v48, _776_v48);
          let _rhs177 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(912)), ((_807_i6) => function (_808_i8) {
            return _807_i6;
          })(_777_i6)), _779_v50);
          let _rhs178 = p0;
          let _lhs160 = _804_v71;
          let _lhs161 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_804_v71).length));
          _776_v48 = _rhs176;
          _780_v51 = _rhs177;
          _lhs160[_lhs161] = _rhs178;
          _798_v65 = (_798_v65).update((_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))], !(_792_cf23));
          (globalState).f7 = (_776_v48)[_module.__default.safeIndex(_777_i6, new BigNumber((_776_v48).length))];
          _798_v65 = ((_792_cf23) ? (_798_v65) : ((_798_v65).Merge(_798_v65)));
        } else {
          let _809___mcc_h18 = (_source14).cf21;
          let _810_cf21 = _809___mcc_h18;
          (globalState).f7 = (_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))];
          (globalState).f13 = (_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))];
          let _811_v72;
          _811_v72 = _dafny.Map.Empty.slice().updateUnsafe(_777_i6,_776_v48);
          let _812_v73;
          let _nw153 = new _module.C0();
          _nw153.__ctor((_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))]);
          _812_v73 = _nw153;
          let _813_v74;
          _813_v74 = _dafny.Seq.of(_812_v73);
          let _814_v75;
          _814_v75 = _dafny.Seq.of(_776_v48, (((_811_v72).contains(new BigNumber((_813_v74).length))) ? ((_811_v72).get(new BigNumber((_813_v74).length))) : (_776_v48)));
          let _815_v76;
          _815_v76 = _module.D1.create_DC6(_777_i6, _814_v75, _777_i6, (_812_v73).f32);
          let _816_v77;
          _816_v77 = _dafny.Set.fromElements(_815_v76);
          let _817_v78;
          _817_v78 = _dafny.MultiSet.fromElements(_816_v77);
          let _818_v79;
          _818_v79 = _module.D3.create_DC12((_812_v73).f32);
          let _819_v80;
          _819_v80 = _dafny.Map.Empty.slice().updateUnsafe(_818_v79,_817_v78);
          (globalState).f22 = ((_dafny.MultiSet.fromElements(_816_v77)).Difference((((_819_v80).contains(_818_v79)) ? ((_819_v80).get(_818_v79)) : (_817_v78)))).IsSubsetOf(_817_v78);
          let _820_v81;
          _820_v81 = _dafny.Seq.UnicodeFromString("epv");
          let _821_v82;
          let _nw154 = new _module.C0();
          _nw154.__ctor((p0).isLessThan(new BigNumber((_820_v81).length)));
          _821_v82 = _nw154;
        }
        let _822_v83;
        let _nw155 = new _module.C0();
        _nw155.__ctor(!(_675_v1) || ((_676_v2)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_676_v2).length))]));
        _822_v83 = _nw155;
      }
      let _823_v84;
      _823_v84 = _module.D4.create_DC15(_module.D4.create_DC14(_675_v1));
      let _824_v85;
      _824_v85 = _dafny.Map.Empty.slice().updateUnsafe((_674_v0).IsProperSubsetOf(_674_v0),_823_v84);
      _824_v85 = (_dafny.Map.Empty.slice().updateUnsafe(false,_823_v84)).Merge(_824_v85);
      let _825_v86;
      _825_v86 = _dafny.Seq.of(_676_v2, _676_v2, _676_v2, ((true) ? (_676_v2) : (_676_v2)), _676_v2);
      _676_v2 = (_825_v86)[_module.__default.safeIndex(p0, new BigNumber((_825_v86).length))];
      let _826_v88;
      _826_v88 = _dafny.Seq.of(p1);
      let _827_v89;
      _827_v89 = _dafny.Map.Empty.slice().updateUnsafe(p1,_675_v1);
      (globalState).f13 = (((function () {
        let _coll21 = new _dafny.Map();
        for (const _compr_21 of (_826_v88).Elements) {
          let _828_v87 = _compr_21;
          if (_dafny.Seq.contains(_826_v88, _828_v87)) {
            _coll21.push([(_828_v87).minus(p0),_675_v1]);
          }
        }
        return _coll21;
      }()).equals(_827_v89)) ? (_675_v1) : (_675_v1));
      let _init18 = ((_829_p1) => function (_830_i9) {
        return _module.__default.safeModuloInt(_830_i9, _829_p1);
      })(p1);
      let _nw156 = Array((new BigNumber(13)).toNumber());
      for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw156.length); _i0_18++) {
        _nw156[_i0_18] = _init18(new BigNumber(_i0_18));
      }
      r0 = _nw156;
      r1 = _675_v1;
      return [r0, r1];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this.f33 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor(f33) {
      let _this = this;
      (_this).f33 = f33;
      return;
    }
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((function () {
        let _coll22 = new _dafny.Map();
        for (const _compr_22 of (_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true, !(true)))).Elements) {
          let _831_v0 = _compr_22;
          if ((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true, !(true)))).contains(_831_v0)) {
            _coll22.push([_831_v0,false]);
          }
        }
        return _coll22;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(!(true)),!(false)));
    };
    fm2(globalState) {
      let _this = this;
      return ((new BigNumber(999)).minus(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-668)), new BigNumber(-602))).length))).plus(((true) ? (new BigNumber((_dafny.Set.fromElements(new BigNumber(-570), new BigNumber((_this.f33).length))).length)) : (new BigNumber((_dafny.Set.fromElements(!(false), true, false)).length))));
    };
    fm31(p0, globalState) {
      let _this = this;
      return !(((!(true)) ? (true) : (false))) || (!_dafny.Seq.contains(_this.f33, new _dafny.CodePoint('n'.codePointAt(0))));
    };
    m0(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _832_v0;
      let _nw157 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _832_v0 = _nw157;
      let _833_v1;
      _833_v1 = _dafny.MultiSet.fromElements(p1, p1);
      let _834_v2;
      let _nw158 = Array((new BigNumber(8)).toNumber());
      _nw158[(_dafny.ZERO).toNumber()] = false;
      _nw158[(_dafny.ONE).toNumber()] = p3;
      _nw158[(new BigNumber(2)).toNumber()] = true;
      _nw158[(new BigNumber(3)).toNumber()] = p1;
      _nw158[(new BigNumber(4)).toNumber()] = p3;
      _nw158[(new BigNumber(5)).toNumber()] = p3;
      _nw158[(new BigNumber(6)).toNumber()] = p3;
      _nw158[(new BigNumber(7)).toNumber()] = p1;
      _834_v2 = _nw158;
      let _835_v3;
      _835_v3 = _dafny.MultiSet.fromElements(_834_v2);
      let _836_v4;
      _836_v4 = _dafny.MultiSet.fromElements(new BigNumber((_833_v1).cardinality()), new BigNumber((_this.f33).length), new BigNumber((_835_v3).cardinality()));
      let _837_v5;
      let _nw159 = new _module.C0();
      _nw159.__ctor(p1);
      _837_v5 = _nw159;
      let _838_v6;
      _838_v6 = _dafny.Set.fromElements(_837_v5, _837_v5);
      let _839_v7;
      _839_v7 = _dafny.Map.Empty.slice().updateUnsafe(_836_v4,new BigNumber((_838_v6).length));
      let _840_v8;
      _840_v8 = _dafny.Set.fromElements(p1);
      let _index168 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length));
      (_832_v0)[_index168] = _module.__default.safeDivisionInt(new BigNumber(855), (((_839_v7).contains(_836_v4)) ? ((_839_v7).get(_836_v4)) : (new BigNumber((_840_v8).length))));
      let _841_v9;
      _841_v9 = _module.D3.create_DC10(_833_v1);
      let _842_v10;
      _842_v10 = _dafny.Set.fromElements(_841_v9, _module.D3.create_DC10(_833_v1), _841_v9, _841_v9);
      let _843_v11;
      _843_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm31(_836_v4, globalState),_842_v10);
      let _index169 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length));
      (_832_v0)[_index169] = new BigNumber(((_843_v11).update((_this).fm31(_836_v4, globalState), (_842_v10).Intersect(_842_v10))).length);
      let _844_v12;
      let _nw160 = Array((new BigNumber(2)).toNumber()).fill([]);
      _844_v12 = _nw160;
      _844_v12 = _844_v12;
      let _845_v13;
      let _nw161 = Array((new BigNumber(12)).toNumber()).fill(_dafny.MultiSet.Empty);
      _845_v13 = _nw161;
      let _846_v14;
      _846_v14 = new _dafny.CodePoint('h'.codePointAt(0));
      let _847_v15;
      _847_v15 = _dafny.Seq.of(_module.__default.fm0(p2, globalState));
      let _848_v16;
      _848_v16 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_847_v15, _module.__default.safeIndex(new BigNumber(468), new BigNumber((_847_v15).length)), p2), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_847_v15, _module.__default.safeIndex(new BigNumber(468), new BigNumber((_847_v15).length)), p2)).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_837_v5).f32,(_dafny.ZERO).minus(p2))).length)), _module.__default.safeIndex((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))], new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_847_v15, _module.__default.safeIndex(new BigNumber(468), new BigNumber((_847_v15).length)), p2), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_847_v15, _module.__default.safeIndex(new BigNumber(468), new BigNumber((_847_v15).length)), p2)).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_837_v5).f32,(_dafny.ZERO).minus(p2))).length))).length)), p0), _dafny.Seq.Create(_module.__default.abs(new BigNumber(261)), function (_849_i0) {
        return new BigNumber(-686);
      }));
      let _index170 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_845_v13).length));
      (_845_v13)[_index170] = (_module.__default.fm32(_846_v14, (_837_v5).f32, _846_v14, globalState)).Difference(_848_v16);
      let _850_v17;
      _850_v17 = _dafny.Seq.of(p1);
      let _851_v18;
      _851_v18 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_850_v17).length));
      let _852_v19;
      _852_v19 = _dafny.Seq.of(_851_v18);
      let _853_v20;
      _853_v20 = _dafny.Map.Empty.slice().updateUnsafe(_852_v19,_dafny.MultiSet.fromElements(_847_v15));
      let _index171 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_845_v13).length));
      (_845_v13)[_index171] = (((_853_v20).contains(_852_v19)) ? ((_853_v20).get(_852_v19)) : (_848_v16));
      if (p1) {
        (_this).f33 = _this.f33;
        if (!(p1)) {
          let _854_v21;
          _854_v21 = _dafny.Map.Empty.slice().updateUnsafe(!((_837_v5).f32),(_837_v5).f32);
          _854_v21 = (_854_v21).Merge(_854_v21);
          let _rhs179 = _832_v0;
          let _rhs180 = _module.__default.fm0(p0, globalState);
          let _lhs162 = globalState;
          _832_v0 = _rhs179;
          _lhs162.f10 = _rhs180;
          let _855_v22;
          let _nw162 = new _module.C0();
          _nw162.__ctor((p1) || (p3));
          _855_v22 = _nw162;
          (globalState).f10 = (_dafny.ZERO).minus((_module.__default.safeModuloInt((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))], p2)).multipliedBy(((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))]).multipliedBy(p2)));
          (_this).f33 = _this.f33;
        } else {
          (globalState).f10 = p0;
          (globalState).f10 = _module.__default.safeDivisionInt((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))], new BigNumber((_850_v17).length));
          (globalState).f10 = _module.__default.fm0((_this).fm2(globalState), globalState);
          (globalState).f7 = (((_this).fm31(_836_v4, globalState)) ? (true) : (!(p0).isEqualTo(new BigNumber((_this.f33).length))));
          let _index172 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length));
          (_832_v0)[_index172] = (_this).fm2(globalState);
        }
        let _856_v23;
        _856_v23 = _module.D2.create_DC9(_module.D2.create_DC9(_module.D2.create_DC7(_851_v18)));
        let _857_v24;
        _857_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC9(_856_v23),(_837_v5).f32);
        let _858_v25;
        _858_v25 = _module.D2.create_DC9(_856_v23);
        _857_v24 = (_857_v24).update(_858_v25, (_this).fm31(_dafny.MultiSet.FromArray(_847_v15), globalState));
        let _859_v26;
        _859_v26 = _dafny.Set.fromElements(new BigNumber((_840_v8).length));
        let _860_v27;
        _860_v27 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.safeDivisionInt(new BigNumber(-76), new BigNumber((_859_v26).length)));
        _860_v27 = (_860_v27).update(p1, (_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))]);
        let _861_v28;
        let _nw163 = new _module.C1();
        _nw163.__ctor();
        _861_v28 = _nw163;
      } else {
        let _862_v29;
        _862_v29 = _dafny.Map.Empty.slice().updateUnsafe((p3) || (false),p2);
        (globalState).f10 = (((_862_v29).contains((_837_v5).f32)) ? ((_862_v29).get((_837_v5).f32)) : ((_this).fm2(globalState)));
        let _index173 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_834_v2).length));
        (_834_v2)[_index173] = true;
        let _index174 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_834_v2).length));
        let _rhs181 = !(_851_v18).contains((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))]);
        let _rhs182 = ((p3) ? (new BigNumber((function () {
          let _coll23 = new _dafny.Map();
          for (const _compr_23 of _dafny.IntegerRange(new BigNumber(416), new BigNumber(822))) {
            let _863_v30 = _compr_23;
            if (((new BigNumber(416)).isLessThanOrEqualTo(_863_v30)) && ((_863_v30).isLessThan(new BigNumber(822)))) {
              _coll23.push([_module.__default.safeDivisionInt(_863_v30, p0),p3]);
            }
          }
          return _coll23;
        }()).length)) : (_module.__default.safeModuloInt(p0, new BigNumber((_847_v15).length))));
        let _lhs163 = _834_v2;
        let _lhs164 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_834_v2).length));
        let _lhs165 = globalState;
        _lhs163[_lhs164] = _rhs181;
        _lhs165.f10 = _rhs182;
        (globalState).f13 = p3;
        let _864_v31;
        _864_v31 = _module.D8.create_DC22((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))], (_834_v2)[_module.__default.safeIndex(new BigNumber(168), new BigNumber((_834_v2).length))], _846_v14);
        let _865_v32;
        _865_v32 = _dafny.Map.Empty.slice().updateUnsafe(!((_864_v31).dtor_cf37),(_837_v5).f32);
        let _source15 = _865_v32;
        let _866___mcc_h0 = _source15;
        let _867_cf27 = _866___mcc_h0;
        let _index175 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length));
        let _rhs183 = false;
        let _rhs184 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_847_v15, _847_v15), _847_v15);
        let _rhs185 = p2;
        let _lhs166 = globalState;
        let _lhs167 = globalState;
        let _lhs168 = _832_v0;
        let _lhs169 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length));
        _lhs166.f13 = _rhs183;
        _lhs167.f13 = _rhs184;
        _lhs168[_lhs169] = _rhs185;
        (globalState).f7 = true;
        let _868_v33;
        let _nw164 = Array((new BigNumber(9)).toNumber());
        _nw164[(_dafny.ZERO).toNumber()] = ((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))]).isLessThanOrEqualTo(new BigNumber(675));
        _nw164[(_dafny.ONE).toNumber()] = !((_836_v4).IsDisjointFrom(_836_v4));
        _nw164[(new BigNumber(2)).toNumber()] = (_this).fm31(_836_v4, globalState);
        _nw164[(new BigNumber(3)).toNumber()] = (_837_v5).f32;
        _nw164[(new BigNumber(4)).toNumber()] = (_837_v5).f32;
        _nw164[(new BigNumber(5)).toNumber()] = (_837_v5).f32;
        _nw164[(new BigNumber(6)).toNumber()] = false;
        _nw164[(new BigNumber(7)).toNumber()] = (p0).isLessThan((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))]);
        _nw164[(new BigNumber(8)).toNumber()] = (_837_v5).f32;
        _868_v33 = _nw164;
        let _rhs186 = _868_v33;
        let _rhs187 = p2;
        let _rhs188 = _dafny.Seq.Concat(_this.f33, _this.f33);
        let _rhs189 = (_dafny.ZERO).minus((_module.__default.fm0(p2, globalState)).plus(_module.__default.fm0((_dafny.ZERO).minus(p2), globalState)));
        let _lhs170 = globalState;
        let _lhs171 = _this;
        let _lhs172 = globalState;
        _868_v33 = _rhs186;
        _lhs170.f10 = _rhs187;
        _lhs171.f33 = _rhs188;
        _lhs172.f10 = _rhs189;
        (globalState).f22 = !(!((_837_v5).f32) || ((_this).fm31(_836_v4, globalState)));
        let _869_v34;
        let _nw165 = Array((new BigNumber(10)).toNumber()).fill(false);
        _869_v34 = _nw165;
        let _870_v35;
        _870_v35 = _module.D0.create_DC1(p0, p2, (_dafny.ZERO).minus(p2), _869_v34);
        let _871_v36;
        _871_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_850_v17).length),(_834_v2)[_module.__default.safeIndex(new BigNumber(168), new BigNumber((_834_v2).length))]);
        _870_v35 = _module.D0.create_DC1(new BigNumber((_871_v36).length), (new BigNumber((_this.f33).length)).plus(new BigNumber((_module.__default.fm33((_837_v5).f32, (_837_v5).f32, p3, (_837_v5).f32, globalState)).length)), (((_851_v18).contains((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))])) ? ((_851_v18).get((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))])) : ((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))])), _869_v34);
      }
      if (p1) {
        if (false) {
          (_this).f33 = _dafny.Seq.Concat(_dafny.Seq.Concat(_this.f33, _dafny.Seq.update(_this.f33, _module.__default.safeIndex(p0, new BigNumber((_this.f33).length)), new _dafny.CodePoint('u'.codePointAt(0)))), _this.f33);
          _850_v17 = _dafny.Seq.Concat(_850_v17, _dafny.Seq.Concat(_850_v17, _850_v17));
          (globalState).f13 = ((true) ? ((_837_v5).f32) : ((_837_v5).f32));
          (globalState).f13 = p3;
          (globalState).f10 = (_dafny.ZERO).minus(p0);
        } else {
          let _872_v37;
          _872_v37 = _dafny.Seq.of(_832_v0);
          _872_v37 = _872_v37;
          let _873_v38;
          let _nw166 = new _module.C1();
          _nw166.__ctor();
          _873_v38 = _nw166;
          (globalState).f10 = _module.__default.fm0(p0, globalState);
          (globalState).f22 = (_873_v38).fm21((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))], globalState);
          let _874_v39;
          let _nw167 = Array((new BigNumber(7)).toNumber()).fill(_dafny.MultiSet.Empty);
          _874_v39 = _nw167;
          let _index176 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_874_v39).length));
          (_874_v39)[_index176] = _833_v1;
          let _index177 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_874_v39).length));
          (_874_v39)[_index177] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_850_v17, _dafny.Seq.Concat(_850_v17, _850_v17)));
        }
        let _875_v40;
        let _nw168 = new _module.C0();
        _nw168.__ctor(!((_837_v5).f32) || (!(p3)));
        _875_v40 = _nw168;
        let _876_v41;
        let _nw169 = Array((new BigNumber(27)).toNumber()).fill([]);
        _876_v41 = _nw169;
        _876_v41 = (((p0).isLessThan(new BigNumber((_module.__default.fm34(p1, (_this.f33)[_module.__default.safeIndex(p2, new BigNumber((_this.f33).length))], (_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))], globalState)).length))) ? (_876_v41) : (_876_v41));
        let _877_v42;
        _877_v42 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.update(_850_v17, _module.__default.safeIndex(p0, new BigNumber((_850_v17).length)), p1), _850_v17), _dafny.Seq.of((_837_v5).f32, false, (_875_v40).f32, (_this).fm31(_dafny.MultiSet.fromElements(p2), globalState)));
        let _878_v43;
        _878_v43 = _dafny.Map.Empty.slice().updateUnsafe((_875_v40).f32,_834_v2);
        let _879_v44;
        _879_v44 = _dafny.Set.fromElements((((_878_v43).contains(p3)) ? ((_878_v43).get(p3)) : (_834_v2)));
        let _880_v45;
        _880_v45 = _module.D4.create_DC13(_846_v14);
        let _881_v46;
        _881_v46 = _module.D4.create_DC15(_880_v45);
        let _882_v47;
        _882_v47 = _module.D7.create_DC20(_847_v15, _834_v2, _881_v46, (_837_v5).f32, (_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))]);
        let _index178 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length));
        let _rhs190 = _dafny.Seq.of(_850_v17, _850_v17, _dafny.Seq.of(!(p3)), _850_v17);
        let _rhs191 = _module.__default.safeModuloInt((p0).multipliedBy(p0), p0);
        let _rhs192 = (_879_v44).IsDisjointFrom(_dafny.Set.fromElements(_834_v2, _834_v2, (_882_v47).dtor_cf31, _834_v2));
        let _lhs173 = _832_v0;
        let _lhs174 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length));
        let _lhs175 = globalState;
        _877_v42 = _rhs190;
        _lhs173[_lhs174] = _rhs191;
        _lhs175.f22 = _rhs192;
        (globalState).f7 = (_837_v5).f32;
      } else {
        let _883_v48;
        _883_v48 = _dafny.Map.Empty.slice().updateUnsafe((_837_v5).f32,(_837_v5).f32);
        let _index179 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length));
        (_832_v0)[_index179] = (_dafny.ZERO).minus((p0).plus(new BigNumber(((_883_v48).update(p1, (_837_v5).f32)).length)));
        let _884_v49;
        _884_v49 = _dafny.Map.Empty.slice().updateUnsafe(_this.f33,p1);
        _883_v48 = (_883_v48).update((_837_v5).f32, (((_884_v49).contains(_this.f33)) ? ((_884_v49).get(_this.f33)) : (p1)));
        let _885_v50;
        let _nw170 = new _module.C1();
        _nw170.__ctor();
        _885_v50 = _nw170;
        let _886_v51;
        _886_v51 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_847_v15, _dafny.Seq.of(new BigNumber((_this.f33).length), p0)), _module.__default.safeIndex(new BigNumber((_this.f33).length), new BigNumber((_dafny.Seq.of(_847_v15, _dafny.Seq.of(new BigNumber((_this.f33).length), p0))).length)), _847_v15)).length));
        let _887_v52;
        _887_v52 = _dafny.Map.Empty.slice().updateUnsafe(_886_v51,_module.__default.fm35((_832_v0)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length))], p3, !((_837_v5).f32), globalState));
        _887_v52 = ((_dafny.Seq.IsProperPrefixOf(_this.f33, _dafny.Seq.UnicodeFromString("nos"))) ? ((_887_v52).update(_886_v51, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-70)), ((_888_v14) => function (_889_i1) {
          return _888_v14;
        })(_846_v14)))) : ((_887_v52).Merge(function () {
          let _coll24 = new _dafny.Map();
          for (const _compr_24 of (_module.__default.fm36(_846_v14, _module.D3.create_DC10(_833_v1), p1, globalState)).Elements) {
            let _890_v53 = _compr_24;
            if (_dafny.Seq.contains(_module.__default.fm36(_846_v14, _module.D3.create_DC10(_833_v1), p1, globalState), _890_v53)) {
              _coll24.push([_890_v53,_dafny.Seq.UnicodeFromString("wqaidjcb")]);
            }
          }
          return _coll24;
        }())));
        (globalState).f7 = p1;
      }
      let _index180 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_832_v0).length));
      (_832_v0)[_index180] = new BigNumber(921);
      return;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let r1 = false;
      let _891_v0;
      let _init19 = ((_892_p2, _893_p1) => function (_894_i0) {
        return _module.D7.create_DC18(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_892_p2,_892_p2)).length)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),_893_p1)));
      })(p2, p1);
      let _nw171 = Array((new BigNumber(3)).toNumber());
      for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw171.length); _i0_19++) {
        _nw171[_i0_19] = _init19(new BigNumber(_i0_19));
      }
      _891_v0 = _nw171;
      let _895_v1;
      _895_v1 = new _dafny.CodePoint('k'.codePointAt(0));
      let _896_v2;
      _896_v2 = _dafny.Seq.of(p2, p2, p2, p2, p2);
      let _897_v3;
      _897_v3 = _dafny.Map.Empty.slice().updateUnsafe(_895_v1,new BigNumber((_896_v2).length));
      let _index181 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_891_v0).length));
      (_891_v0)[_index181] = _module.D7.create_DC18(_dafny.Seq.update(_dafny.Seq.of(_897_v3), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.of(_897_v3)).length)), _897_v3));
      let _index182 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_891_v0).length));
      (_891_v0)[_index182] = _module.__default.fm37(p1, p1, p1, globalState);
      (_this).f33 = _dafny.Seq.UnicodeFromString("xcb");
      let _898_v4;
      let _init20 = ((_899_p1) => function (_900_i1) {
        return _module.__default.safeModuloInt(_900_i1, _899_p1);
      })(p1);
      let _nw172 = Array((new BigNumber(16)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw172.length); _i0_20++) {
        _nw172[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _898_v4 = _nw172;
      let _index183 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_898_v4).length));
      (_898_v4)[_index183] = p1;
      let _901_v5;
      _901_v5 = _dafny.Set.fromElements(p1, p1, new BigNumber((_dafny.Seq.of(p2)).length));
      let _902_v6;
      _902_v6 = _dafny.MultiSet.fromElements(_module.__default.fm0(new BigNumber((_901_v5).length), globalState), p1);
      let _903_v7;
      _903_v7 = _module.D8.create_DC22(p1, p2, _895_v1);
      let _904_v8;
      _904_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _905_v10;
      _905_v10 = _module.D1.create_DC6(p1, _dafny.Seq.of(_896_v2, _dafny.Seq.of(p2)), p1, p2);
      let _906_v11;
      let _nw173 = Array((new BigNumber(3)).toNumber());
      _nw173[(_dafny.ZERO).toNumber()] = (new BigNumber((_this.f33).length)).isLessThanOrEqualTo((((_902_v6).contains((_903_v7).dtor_cf36)) ? ((_902_v6).get((_903_v7).dtor_cf36)) : (p1)));
      _nw173[(_dafny.ONE).toNumber()] = (function () {
        let _coll25 = new _dafny.Set();
        for (const _compr_25 of (_904_v8).Keys.Elements) {
          let _907_v9 = _compr_25;
          if ((_904_v8).contains(_907_v9)) {
            _coll25.add((_907_v9).multipliedBy(new BigNumber(-326)));
          }
        }
        return _coll25;
      }()).IsProperSubsetOf(_901_v5);
      _nw173[(new BigNumber(2)).toNumber()] = (_905_v10).dtor_cf16;
      _906_v11 = _nw173;
      let _index184 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_906_v11).length));
      (_906_v11)[_index184] = !(p2);
      let _index185 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_898_v4).length));
      (_898_v4)[_index185] = p1;
      let _index186 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_898_v4).length));
      let _index187 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_906_v11).length));
      let _index188 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_898_v4).length));
      let _rhs193 = new BigNumber((_dafny.Seq.Concat(_896_v2, _dafny.Seq.update(_896_v2, _module.__default.safeIndex((_this).fm2(globalState), new BigNumber((_896_v2).length)), (_this).fm31(_902_v6, globalState)))).length);
      let _rhs194 = !(!((_dafny.Set.fromElements(p2, p2)).IsDisjointFrom(_module.__default.fm38(new BigNumber((p0).length), true, globalState)))) || (p2);
      let _rhs195 = (p1).multipliedBy(_module.__default.safeModuloInt(new BigNumber(-561), p1));
      let _rhs196 = new BigNumber((_902_v6).cardinality());
      let _rhs197 = p1;
      let _lhs176 = _898_v4;
      let _lhs177 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_898_v4).length));
      let _lhs178 = _906_v11;
      let _lhs179 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_906_v11).length));
      let _lhs180 = globalState;
      let _lhs181 = _898_v4;
      let _lhs182 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_898_v4).length));
      let _lhs183 = globalState;
      _lhs176[_lhs177] = _rhs193;
      _lhs178[_lhs179] = _rhs194;
      _lhs180.f10 = _rhs195;
      _lhs181[_lhs182] = _rhs196;
      _lhs183.f10 = _rhs197;
      let _ingredients0 = [];
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_898_v4).length))) {
        let _908_i2 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_908_i2)) && ((_908_i2).isLessThan(new BigNumber((_898_v4).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_898_v4, (_908_i2).toNumber(), _module.__default.safeModuloInt(_908_i2, (_898_v4)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_898_v4).length))])));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      let _909_v12;
      _909_v12 = _module.D0.create_DC0(_this.f33);
      let _910_v13;
      _910_v13 = _dafny.Map.Empty.slice().updateUnsafe((_898_v4)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_898_v4).length))],_this.f33);
      let _911_v14;
      let _nw174 = Array((new BigNumber(20)).toNumber());
      _nw174[(_dafny.ZERO).toNumber()] = (_909_v12).dtor_cf0;
      _nw174[(_dafny.ONE).toNumber()] = (_module.D0.create_DC0(_this.f33)).dtor_cf0;
      _nw174[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("gbvposwab");
      _nw174[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_this.f33, _module.__default.safeIndex(p1, new BigNumber((_this.f33).length)), _895_v1);
      _nw174[(new BigNumber(4)).toNumber()] = _this.f33;
      _nw174[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("iwtorenfq");
      _nw174[(new BigNumber(6)).toNumber()] = _this.f33;
      _nw174[(new BigNumber(7)).toNumber()] = (((_910_v13).contains(p1)) ? ((_910_v13).get(p1)) : (_this.f33));
      _nw174[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("scc");
      _nw174[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qltrd"), _this.f33);
      _nw174[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(30)), ((_912_v1) => function (_913_i3) {
        return _912_v1;
      })(_895_v1));
      _nw174[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("ssiokybl");
      _nw174[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("atla");
      _nw174[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("j");
      _nw174[(new BigNumber(14)).toNumber()] = _dafny.Seq.update(_this.f33, _module.__default.safeIndex(new BigNumber((_904_v8).length), new BigNumber((_this.f33).length)), _895_v1);
      _nw174[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_this.f33, _module.__default.fm35((_898_v4)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_898_v4).length))], !(p2), (_906_v11)[_module.__default.safeIndex(new BigNumber(951), new BigNumber((_906_v11).length))], globalState));
      _nw174[(new BigNumber(16)).toNumber()] = _this.f33;
      _nw174[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("uokfrfjh");
      _nw174[(new BigNumber(18)).toNumber()] = _this.f33;
      _nw174[(new BigNumber(19)).toNumber()] = _module.__default.fm35(new BigNumber(308), p2, false, globalState);
      _911_v14 = _nw174;
      let _914_v15;
      _914_v15 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_906_v11)[_module.__default.safeIndex(new BigNumber(951), new BigNumber((_906_v11).length))]);
      let _index189 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_911_v14).length));
      (_911_v14)[_index189] = _dafny.Seq.update(_this.f33, _module.__default.safeIndex(new BigNumber((_this.f33).length), new BigNumber((_this.f33).length)), (_this.f33)[_module.__default.safeIndex(new BigNumber((_914_v15).length), new BigNumber((_this.f33).length))]);
      let _index190 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_911_v14).length));
      (_911_v14)[_index190] = _this.f33;
      let _915_v16;
      _915_v16 = _dafny.Set.fromElements(p2, (_dafny.MultiSet.fromElements(new BigNumber((_896_v2).length))).IsProperSubsetOf(_902_v6));
      let _916_v17;
      _916_v17 = _dafny.Seq.of(_915_v16);
      let _rhs198 = ((_module.__default.fm38(p1, p2, globalState)).Intersect(_915_v16)).Difference((_916_v17)[_module.__default.safeIndex(p1, new BigNumber((_916_v17).length))]);
      let _rhs199 = _895_v1;
      _915_v16 = _rhs198;
      _895_v1 = _rhs199;
      r0 = _module.D0.create_DC1(p1, _module.__default.safeModuloInt(p1, p1), p1, _906_v11);
      r1 = p2;
      return [r0, r1];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(!(false), true, false),!(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true, !(true)),!(true)));
    };
    fm2(globalState) {
      let _this = this;
      let _source16 = _module.D1.create_DC6(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll26 = new _dafny.Set();
  for (const _compr_26 of _dafny.IntegerRange(new BigNumber(24), new BigNumber(182))) {
    let _917_v0 = _compr_26;
    if (((new BigNumber(24)).isLessThanOrEqualTo(_917_v0)) && ((_917_v0).isLessThan(new BigNumber(182)))) {
      _coll26.add(_module.__default.safeModuloInt(_917_v0, new BigNumber(21)));
    }
  }
  return _coll26;
}()).length))).length), _dafny.Seq.of(_dafny.Seq.of(false, false), _dafny.Seq.of(false)), new BigNumber(423), !(true));
      if (_source16.is_DC5) {
        let _918___mcc_h0 = (_source16).cf9;
        let _919___mcc_h1 = (_source16).cf10;
        let _920___mcc_h2 = (_source16).cf11;
        let _921___mcc_h3 = (_source16).cf12;
        let _922_cf12 = _921___mcc_h3;
        let _923_cf11 = _920___mcc_h2;
        let _924_cf10 = _919___mcc_h1;
        let _925_cf9 = _918___mcc_h0;
        return _module.__default.safeDivisionInt(_925_cf9, _925_cf9);
      } else if (_source16.is_DC6) {
        let _926___mcc_h4 = (_source16).cf13;
        let _927___mcc_h5 = (_source16).cf14;
        let _928___mcc_h6 = (_source16).cf15;
        let _929___mcc_h7 = (_source16).cf16;
        let _930_cf16 = _929___mcc_h7;
        let _931_cf15 = _928___mcc_h6;
        let _932_cf14 = _927___mcc_h5;
        let _933_cf13 = _926___mcc_h4;
        return _933_cf13;
      } else {
        let _934___mcc_h8 = (_source16).cf8;
        let _935_cf8 = _934___mcc_h8;
        return _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("vpngnpve")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(872),new BigNumber(-382))).length));
      }
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt((new BigNumber((_dafny.Seq.UnicodeFromString("yxgl")).length)).multipliedBy(new BigNumber(-516)), (new BigNumber(183)).multipliedBy(new BigNumber(((_module.D0.create_DC0(_dafny.Seq.UnicodeFromString("ulgojwmoa"))).dtor_cf0).length)));
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return _module.D1.create_DC4(((_module.D1.create_DC4(function () {
  let _coll27 = new _dafny.Map();
  for (const _compr_27 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true, !(false), true, false, false),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-510),new BigNumber((function () {
    let _coll28 = new _dafny.Map();
    for (const _compr_28 of _dafny.IntegerRange(new BigNumber(521), new BigNumber(997))) {
      let _936_v1 = _compr_28;
      if (((new BigNumber(521)).isLessThanOrEqualTo(_936_v1)) && ((_936_v1).isLessThan(new BigNumber(997)))) {
        _coll28.push([_module.__default.safeDivisionInt(_936_v1, new BigNumber(433)),new BigNumber((_dafny.Seq.of(false)).length)]);
      }
    }
    return _coll28;
  }()).length))).length))).Keys.Elements) {
    let _937_v0 = _compr_27;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true, !(false), true, false, false),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-510),new BigNumber((function () {
      let _coll29 = new _dafny.Map();
      for (const _compr_29 of _dafny.IntegerRange(new BigNumber(521), new BigNumber(997))) {
        let _936_v1 = _compr_29;
        if (((new BigNumber(521)).isLessThanOrEqualTo(_936_v1)) && ((_936_v1).isLessThan(new BigNumber(997)))) {
          _coll29.push([_module.__default.safeDivisionInt(_936_v1, new BigNumber(433)),new BigNumber((_dafny.Seq.of(false)).length)]);
        }
      }
      return _coll29;
    }()).length))).length))).contains(_937_v0)) {
      _coll27.push([_937_v0,false]);
    }
  }
  return _coll27;
}())).dtor_cf8).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false, false, false, false),false)));
    };
    m0(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _938_v0;
      _938_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(604));
      let _939_v1;
      _939_v1 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_938_v0).length));
      let _940_v2;
      _940_v2 = _dafny.Seq.of(_938_v0);
      (globalState).f22 = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_940_v2, _module.__default.safeIndex(p0, new BigNumber((_940_v2).length)), (_module.D2.create_DC7(_939_v1)).dtor_cf17), _dafny.Seq.Concat(_940_v2, _940_v2));
      let _941_v3;
      let _nw175 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
      _941_v3 = _nw175;
      let _index191 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length));
      (_941_v3)[_index191] = p0;
      let _index192 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length));
      (_941_v3)[_index192] = _module.__default.safeModuloInt(p2, p0);
      let _942_v4;
      _942_v4 = _dafny.Seq.UnicodeFromString("iiynvtexy");
      let _943_v5;
      _943_v5 = _module.D0.create_DC0(_942_v4);
      let _944_v6;
      _944_v6 = new _dafny.CodePoint('j'.codePointAt(0));
      let _945_v7;
      _945_v7 = _dafny.Map.Empty.slice().updateUnsafe(p1,_942_v4);
      let _pat_let_tv19 = _942_v4;
      let _pat_let_tv20 = _942_v4;
      let _pat_let_tv21 = _943_v5;
      let _pat_let_tv22 = _942_v4;
      let _946_v8;
      let _nw176 = Array((new BigNumber(16)).toNumber());
      _nw176[(_dafny.ZERO).toNumber()] = _943_v5;
      _nw176[(_dafny.ONE).toNumber()] = _943_v5;
      _nw176[(new BigNumber(2)).toNumber()] = _module.D0.create_DC0(_dafny.Seq.update(_942_v4, _module.__default.safeIndex((_941_v3)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length))], new BigNumber((_942_v4).length)), _944_v6));
      _nw176[(new BigNumber(3)).toNumber()] = _943_v5;
      _nw176[(new BigNumber(4)).toNumber()] = _943_v5;
      _nw176[(new BigNumber(5)).toNumber()] = function (_pat_let18_0) {
        return function (_949_dt__update__tmp_h1) {
          return function (_pat_let21_0) {
            return function (_950_dt__update_hcf0_h1) {
              return _module.D0.create_DC0(_950_dt__update_hcf0_h1);
            }(_pat_let21_0);
          }(_pat_let_tv20);
        }(_pat_let18_0);
      }(function (_pat_let19_0) {
        return function (_947_dt__update__tmp_h0) {
          return function (_pat_let20_0) {
            return function (_948_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_948_dt__update_hcf0_h0);
            }(_pat_let20_0);
          }(_pat_let_tv19);
        }(_pat_let19_0);
      }(_943_v5));
      _nw176[(new BigNumber(6)).toNumber()] = _module.D0.create_DC0(_module.__default.fm10(_dafny.Map.Empty.slice().updateUnsafe((_941_v3)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length))],p0), _module.__default.fm11((_941_v3)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length))], globalState), p3, globalState));
      _nw176[(new BigNumber(7)).toNumber()] = _943_v5;
      _nw176[(new BigNumber(8)).toNumber()] = _943_v5;
      _nw176[(new BigNumber(9)).toNumber()] = _943_v5;
      _nw176[(new BigNumber(10)).toNumber()] = _943_v5;
      _nw176[(new BigNumber(11)).toNumber()] = function (_pat_let22_0) {
        return function (_951_dt__update__tmp_h2) {
          return function (_pat_let23_0) {
            return function (_952_dt__update_hcf0_h2) {
              return _module.D0.create_DC0(_952_dt__update_hcf0_h2);
            }(_pat_let23_0);
          }((_pat_let_tv21).dtor_cf0);
        }(_pat_let22_0);
      }(_943_v5);
      _nw176[(new BigNumber(12)).toNumber()] = _943_v5;
      _nw176[(new BigNumber(13)).toNumber()] = function (_pat_let24_0) {
        return function (_953_dt__update__tmp_h3) {
          return function (_pat_let25_0) {
            return function (_954_dt__update_hcf0_h3) {
              return _module.D0.create_DC0(_954_dt__update_hcf0_h3);
            }(_pat_let25_0);
          }(_pat_let_tv22);
        }(_pat_let24_0);
      }(_module.D0.create_DC0((((_945_v7).contains(p3)) ? ((_945_v7).get(p3)) : (_942_v4))));
      _nw176[(new BigNumber(14)).toNumber()] = _943_v5;
      _nw176[(new BigNumber(15)).toNumber()] = _943_v5;
      _946_v8 = _nw176;
      let _index193 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_946_v8).length));
      (_946_v8)[_index193] = _943_v5;
      let _955_v9;
      _955_v9 = _dafny.Seq.of(true, p1, !(p1));
      let _pat_let_tv23 = _942_v4;
      let _pat_let_tv24 = p2;
      let _pat_let_tv25 = _942_v4;
      let _pat_let_tv26 = _944_v6;
      let _index194 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_946_v8).length));
      (_946_v8)[_index194] = function (_pat_let26_0) {
        return function (_956_dt__update__tmp_h4) {
          return function (_pat_let27_0) {
            return function (_957_dt__update_hcf0_h4) {
              return _module.D0.create_DC0(_957_dt__update_hcf0_h4);
            }(_pat_let27_0);
          }(_dafny.Seq.update(_dafny.Seq.Concat(_pat_let_tv23, _dafny.Seq.UnicodeFromString("q")), _module.__default.safeIndex(_pat_let_tv24, new BigNumber((_dafny.Seq.Concat(_pat_let_tv25, _dafny.Seq.UnicodeFromString("q"))).length)), _pat_let_tv26));
        }(_pat_let26_0);
      }(((!((_955_v9)[_module.__default.safeIndex(p0, new BigNumber((_955_v9).length))])) ? (_943_v5) : (_module.D0.create_DC0((_943_v5).dtor_cf0))));
      let _958_v10;
      _958_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_955_v9).length),p3);
      let _959_v11;
      _959_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_958_v10).update((_941_v3)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length))], p1)).length),_941_v3);
      _959_v11 = (_959_v11).update(p0, _941_v3);
      if ((p3) === (!(p1))) {
        let _index195 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length));
        (_941_v3)[_index195] = new BigNumber((_module.__default.fm10((_938_v0).Merge(_939_v1), new _dafny.CodePoint('s'.codePointAt(0)), p3, globalState)).length);
        let _960_v12;
        _960_v12 = _dafny.MultiSet.fromElements(p2);
        let _961_v13;
        _961_v13 = _dafny.Map.Empty.slice().updateUnsafe(_960_v12,_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(510)), ((_962_p0) => function (_963_i0) {
          return _962_p0;
        })(p0)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(510)), ((_964_p0) => function (_965_i0) {
          return _964_p0;
        })(p0))).length)), p0));
        let _966_v14;
        _966_v14 = _module.D1.create_DC5(new BigNumber((_955_v9).length), p1, p1, _961_v13);
        let _pat_let_tv27 = p3;
        let _pat_let_tv28 = p2;
        let _pat_let_tv29 = _961_v13;
        let _index196 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length));
        (_941_v3)[_index196] = (_dafny.ZERO).minus((new BigNumber((_958_v10).length)).minus((function (_pat_let28_0) {
          return function (_967_dt__update__tmp_h5) {
            return function (_pat_let29_0) {
              return function (_968_dt__update_hcf11_h0) {
                return function (_pat_let30_0) {
                  return function (_969_dt__update_hcf9_h0) {
                    return function (_pat_let31_0) {
                      return function (_970_dt__update_hcf12_h0) {
                        return _module.D1.create_DC5(_969_dt__update_hcf9_h0, (_967_dt__update__tmp_h5).dtor_cf10, _968_dt__update_hcf11_h0, _970_dt__update_hcf12_h0);
                      }(_pat_let31_0);
                    }(_pat_let_tv29);
                  }(_pat_let30_0);
                }(_pat_let_tv28);
              }(_pat_let29_0);
            }(_pat_let_tv27);
          }(_pat_let28_0);
        }(_966_v14)).dtor_cf9));
        let _index197 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length));
        (_941_v3)[_index197] = p2;
        (globalState).f22 = _module.__default.fm12(p1, p1, globalState);
        let _971_v15;
        _971_v15 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_942_v4),p1);
        let _972_v16;
        let _nw177 = Array((new BigNumber(21)).toNumber());
        _nw177[(_dafny.ZERO).toNumber()] = p1;
        _nw177[(_dafny.ONE).toNumber()] = p1;
        _nw177[(new BigNumber(2)).toNumber()] = (((_971_v15).contains((_946_v8)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_946_v8).length))])) ? ((_971_v15).get((_946_v8)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_946_v8).length))])) : (p1));
        _nw177[(new BigNumber(3)).toNumber()] = p1;
        _nw177[(new BigNumber(4)).toNumber()] = false;
        _nw177[(new BigNumber(5)).toNumber()] = p3;
        _nw177[(new BigNumber(6)).toNumber()] = true;
        _nw177[(new BigNumber(7)).toNumber()] = p3;
        _nw177[(new BigNumber(8)).toNumber()] = false;
        _nw177[(new BigNumber(9)).toNumber()] = p3;
        _nw177[(new BigNumber(10)).toNumber()] = p1;
        _nw177[(new BigNumber(11)).toNumber()] = p3;
        _nw177[(new BigNumber(12)).toNumber()] = p3;
        _nw177[(new BigNumber(13)).toNumber()] = p1;
        _nw177[(new BigNumber(14)).toNumber()] = p1;
        _nw177[(new BigNumber(15)).toNumber()] = p1;
        _nw177[(new BigNumber(16)).toNumber()] = true;
        _nw177[(new BigNumber(17)).toNumber()] = p1;
        _nw177[(new BigNumber(18)).toNumber()] = p3;
        _nw177[(new BigNumber(19)).toNumber()] = p1;
        _nw177[(new BigNumber(20)).toNumber()] = false;
        _972_v16 = _nw177;
        let _973_v17;
        _973_v17 = _dafny.Map.Empty.slice().updateUnsafe(p2,_972_v16);
        let _974_v18;
        _974_v18 = _dafny.Seq.of(_972_v16, _972_v16, _972_v16);
        _973_v17 = (_973_v17).update(p0, (_974_v18)[_module.__default.safeIndex((_this).fm2(globalState), new BigNumber((_974_v18).length))]);
      } else {
        let _975_v19;
        let _nw178 = new _module.C0();
        _nw178.__ctor(p3);
        _975_v19 = _nw178;
        _942_v4 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(133)), ((_976_v6) => function (_977_i1) {
          return _976_v6;
        })(_944_v6));
        (globalState).f10 = new BigNumber(718);
        if (_module.__default.fm12(!(p3), (_975_v19).f32, globalState)) {
          let _978_v20;
          _978_v20 = _dafny.Set.fromElements(p3, p3, (_975_v19).f32, true, p3);
          let _979_v21;
          _979_v21 = _dafny.Seq.of(p0);
          let _index198 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length));
          (_941_v3)[_index198] = ((new BigNumber((_945_v7).length)).multipliedBy(new BigNumber((_978_v20).length))).plus(new BigNumber((_dafny.Seq.Concat(_979_v21, _module.__default.fm14(false, new BigNumber(-248), p2, globalState))).length));
          let _980_v22;
          let _981_v23;
          let _982_v24;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = (_this).m4(_942_v4, new BigNumber((_979_v21).length), globalState);
          _out5 = _outcollector1[0];
          _out6 = _outcollector1[1];
          _out7 = _outcollector1[2];
          _980_v22 = _out5;
          _981_v23 = _out6;
          _982_v24 = _out7;
          (globalState).f22 = (_955_v9)[_module.__default.safeIndex(p0, new BigNumber((_955_v9).length))];
          let _983_v25;
          _983_v25 = _dafny.Map.Empty.slice().updateUnsafe(_941_v3,new _dafny.CodePoint('u'.codePointAt(0)));
          _983_v25 = (_983_v25).update(_941_v3, _944_v6);
          let _984_v26;
          let _nw179 = new _module.C0();
          _nw179.__ctor(p1);
          _984_v26 = _nw179;
        } else {
          let _985_v27;
          let _nw180 = new _module.C0();
          _nw180.__ctor((_dafny.Map.Empty.slice().updateUnsafe(!(p3),_955_v9)).equals(_dafny.Map.Empty.slice().updateUnsafe(!((_955_v9)[_module.__default.safeIndex(new BigNumber(-175), new BigNumber((_955_v9).length))]),_955_v9)));
          _985_v27 = _nw180;
          (globalState).f22 = !((_985_v27).f32);
          (globalState).f13 = p1;
          _958_v10 = (_958_v10).update(p0, !((_975_v19).f32));
          (globalState).f10 = _module.__default.safeDivisionInt((new BigNumber((_958_v10).length)).plus(new BigNumber((_942_v4).length)), ((_dafny.ZERO).minus(p2)).plus((_941_v3)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length))]));
        }
        (globalState).f19 = (((_975_v19).f32) ? (_944_v6) : ((((_975_v19).f32) ? (_944_v6) : (_944_v6))));
      }
      let _986_v28;
      let _nw181 = Array((new BigNumber(14)).toNumber());
      _nw181[(_dafny.ZERO).toNumber()] = _958_v10;
      _nw181[(_dafny.ONE).toNumber()] = (_958_v10).Merge((_958_v10).update((_941_v3)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length))], true));
      _nw181[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),p3);
      _nw181[(new BigNumber(3)).toNumber()] = (_958_v10).Merge(_958_v10);
      _nw181[(new BigNumber(4)).toNumber()] = _958_v10;
      _nw181[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_941_v3)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length))],p1);
      _nw181[(new BigNumber(6)).toNumber()] = _958_v10;
      _nw181[(new BigNumber(7)).toNumber()] = ((p1) ? (_958_v10) : (_dafny.Map.Empty.slice().updateUnsafe((_941_v3)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length))],p1)));
      _nw181[(new BigNumber(8)).toNumber()] = _958_v10;
      _nw181[(new BigNumber(9)).toNumber()] = _958_v10;
      _nw181[(new BigNumber(10)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_941_v3)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length))],!(p1))).Merge(_958_v10);
      _nw181[(new BigNumber(11)).toNumber()] = (_958_v10).Merge((_958_v10).update(p2, p1));
      _nw181[(new BigNumber(12)).toNumber()] = (_958_v10).Merge((_958_v10).update((_941_v3)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_941_v3).length))], !(p3)));
      _nw181[(new BigNumber(13)).toNumber()] = _958_v10;
      _986_v28 = _nw181;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_986_v28).length))) {
        let _987_i2 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_987_i2)) && ((_987_i2).isLessThan(new BigNumber((_986_v28).length))))) {
          (_986_v28)[(_987_i2)] = _958_v10;
        }
      }
      return;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let r1 = false;
      let _988_v0;
      let _nw182 = Array((new BigNumber(11)).toNumber()).fill([]);
      _988_v0 = _nw182;
      let _989_v1;
      let _nw183 = Array((new BigNumber(22)).toNumber()).fill(false);
      _989_v1 = _nw183;
      let _index199 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_988_v0).length));
      (_988_v0)[_index199] = _989_v1;
      let _index200 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_988_v0).length));
      (_988_v0)[_index200] = ((_module.__default.fm12(p2, p2, globalState)) ? (_989_v1) : (_989_v1));
      let _990_v2;
      _990_v2 = _dafny.Seq.UnicodeFromString("nnaijwfws");
      (globalState).f10 = new BigNumber((_990_v2).length);
      let _index201 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_989_v1).length));
      (_989_v1)[_index201] = p2;
      let _index202 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_989_v1).length));
      (_989_v1)[_index202] = (_module.D1.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("ydkxlx")).length), _dafny.Seq.of(_dafny.Seq.of(p2, p2, p2, p2, p2)), p1, !(p2))).dtor_cf16;
      let _991_v3;
      let _nw184 = new _module.C0();
      _nw184.__ctor(true);
      _991_v3 = _nw184;
      let _992_i0;
      _992_i0 = _dafny.ZERO;
      L1: {
        while (p2) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_992_i0)) {
              break L1;
            }
            _992_i0 = (_992_i0).plus(_dafny.ONE);
            (globalState).f13 = (_991_v3).f32;
            let _993_v4;
            let _nw185 = new _module.C0();
            _nw185.__ctor(p2);
            _993_v4 = _nw185;
            let _994_v5;
            _994_v5 = new _dafny.CodePoint('y'.codePointAt(0));
            let _995_v6;
            _995_v6 = _dafny.MultiSet.fromElements((_993_v4).f32, !((_991_v3).f32));
            let _996_v7;
            let _nw186 = new _module.C0();
            _nw186.__ctor(!(!((_module.__default.fm15(_994_v5, p1, _module.__default.fm12((_989_v1)[_module.__default.safeIndex(new BigNumber(582), new BigNumber((_989_v1).length))], p2, globalState), globalState)).dtor_cf21).equals(_995_v6)));
            _996_v7 = _nw186;
            let _997_v8;
            _997_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_994_v5,_module.D0.create_DC2(p1, false))).length));
            let _998_v9;
            _998_v9 = _dafny.Seq.of(p1, ((_dafny.ZERO).minus(p1)).multipliedBy(p1), p1, (_dafny.ZERO).minus((((_997_v8).contains(p1)) ? ((_997_v8).get(p1)) : (p1))), new BigNumber(-191));
            let _999_v10;
            _999_v10 = _dafny.Seq.of((_989_v1)[_module.__default.safeIndex(new BigNumber(582), new BigNumber((_989_v1).length))]);
            let _1000_v11;
            _1000_v11 = _dafny.Set.fromElements((_989_v1)[_module.__default.safeIndex(new BigNumber(582), new BigNumber((_989_v1).length))], (_989_v1)[_module.__default.safeIndex(new BigNumber(582), new BigNumber((_989_v1).length))], p2);
            let _rhs200 = ((p2) ? (p2) : (_dafny.areEqual(_999_v10, _module.__default.fm16(globalState))));
            let _rhs201 = _dafny.Seq.Concat(_998_v9, _dafny.Seq.Concat(_998_v9, _998_v9));
            let _rhs202 = p1;
            let _rhs203 = p2;
            let _rhs204 = new BigNumber(((_1000_v11).Intersect((_dafny.Set.fromElements(true)).Difference(_1000_v11))).length);
            let _lhs184 = globalState;
            let _lhs185 = globalState;
            let _lhs186 = globalState;
            let _lhs187 = globalState;
            _lhs184.f7 = _rhs200;
            _998_v9 = _rhs201;
            _lhs185.f10 = _rhs202;
            _lhs186.f22 = _rhs203;
            _lhs187.f10 = _rhs204;
          }
        }
      }
      let _1001_v12;
      let _nw187 = Array((new BigNumber(9)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1001_v12 = _nw187;
      let _1002_v13;
      _1002_v13 = new _dafny.CodePoint('h'.codePointAt(0));
      let _index203 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1001_v12).length));
      (_1001_v12)[_index203] = _1002_v13;
      let _index204 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1001_v12).length));
      (_1001_v12)[_index204] = _1002_v13;
      let _1003_v14;
      _1003_v14 = _module.D0.create_DC1(p1, p1, (_dafny.ZERO).minus(new BigNumber((_990_v2).length)), (_988_v0)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_988_v0).length))]);
      let _1004_v15;
      _1004_v15 = _dafny.Seq.of(_991_v3);
      let _1005_v16;
      _1005_v16 = _dafny.Set.fromElements(_1004_v15);
      let _pat_let_tv30 = _1005_v16;
      let _pat_let_tv31 = p1;
      let _pat_let_tv32 = p2;
      let _pat_let_tv33 = _989_v1;
      r0 = function (_pat_let32_0) {
        return function (_1006_dt__update__tmp_h0) {
          return function (_pat_let33_0) {
            return function (_1007_dt__update_hcf3_h0) {
              return function (_pat_let34_0) {
                return function (_1008_dt__update_hcf1_h0) {
                  return function (_pat_let35_0) {
                    return function (_1009_dt__update_hcf4_h0) {
                      return _module.D0.create_DC1(_1008_dt__update_hcf1_h0, (_1006_dt__update__tmp_h0).dtor_cf2, _1007_dt__update_hcf3_h0, _1009_dt__update_hcf4_h0);
                    }(_pat_let35_0);
                  }(_pat_let_tv33);
                }(_pat_let34_0);
              }(((_pat_let_tv32) ? ((_dafny.ZERO).minus(new BigNumber((_pat_let_tv30).length))) : (_pat_let_tv31)));
            }(_pat_let33_0);
          }(new BigNumber(-614));
        }(_pat_let32_0);
      }(_1003_v14);
      let _1010_v17;
      _1010_v17 = _dafny.MultiSet.fromElements(p1, p1, p1);
      let _1011_v18;
      _1011_v18 = _dafny.MultiSet.fromElements(_1010_v17);
      r1 = ((_1011_v18).update(_1010_v17, _module.__default.abs(p1))).IsDisjointFrom(_1011_v18);
      return [r0, r1];
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = undefined;
      let _1012_v0;
      _1012_v0 = false;
      let _1013_i0;
      _1013_i0 = _dafny.ZERO;
      L2: {
        while (_1012_v0) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1013_i0)) {
              break L2;
            }
            _1013_i0 = (_1013_i0).plus(_dafny.ONE);
            let _1014_v1;
            let _nw188 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
            _1014_v1 = _nw188;
            let _1015_v2;
            _1015_v2 = _dafny.MultiSet.fromElements(_1014_v1, _1014_v1);
            let _1016_v3;
            _1016_v3 = _dafny.Seq.of(_1012_v0, !(_1015_v2).equals(_1015_v2), _1012_v0, !(((_module.__default.fm12(_1012_v0, _1012_v0, globalState)) ? (_1012_v0) : (_1012_v0))));
            _1016_v3 = _1016_v3;
            let _1017_v4;
            _1017_v4 = _dafny.Seq.of(_1016_v3, _1016_v3, _1016_v3, _module.__default.fm16(globalState));
            let _1018_v5;
            _1018_v5 = _module.D1.create_DC6(p1, _1017_v4, p1, _1012_v0);
            let _index205 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1014_v1).length));
            (_1014_v1)[_index205] = new BigNumber((p0).length);
            let _1019_v6;
            _1019_v6 = _dafny.Seq.of(p1);
            let _1020_v7;
            _1020_v7 = _dafny.Set.fromElements(p1, p1);
            let _1021_v8;
            _1021_v8 = new _dafny.CodePoint('a'.codePointAt(0));
            let _1022_v9;
            _1022_v9 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm17((_dafny.ZERO).minus(p1), _1012_v0, _1021_v8, globalState),_1019_v6);
            let _1023_v10;
            _1023_v10 = _module.D1.create_DC5(p1, !(_1012_v0), _1012_v0, _1022_v9);
            let _index206 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1014_v1).length));
            let _rhs205 = _1012_v0;
            let _rhs206 = !(_1020_v7).contains((_1019_v6)[_module.__default.safeIndex(p1, new BigNumber((_1019_v6).length))]);
            let _rhs207 = _module.__default.fm12(_1012_v0, _1012_v0, globalState);
            let _rhs208 = _1018_v5;
            let _rhs209 = (_1023_v10).dtor_cf9;
            let _lhs188 = globalState;
            let _lhs189 = _1014_v1;
            let _lhs190 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1014_v1).length));
            _lhs188.f7 = _rhs205;
            r0 = _rhs206;
            _1012_v0 = _rhs207;
            _1018_v5 = _rhs208;
            _lhs189[_lhs190] = _rhs209;
            let _1024_v11;
            let _nw189 = Array((new BigNumber(24)).toNumber()).fill(false);
            _1024_v11 = _nw189;
            let _index207 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_1024_v11).length));
            (_1024_v11)[_index207] = true;
            let _index208 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_1024_v11).length));
            (_1024_v11)[_index208] = _1012_v0;
            let _1025_v12;
            _1025_v12 = _module.D3.create_DC11((_1014_v1)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1014_v1).length))]);
            let _1026_v13;
            _1026_v13 = _dafny.Set.fromElements(_1014_v1, _1014_v1);
            let _pat_let_tv34 = _1014_v1;
            let _pat_let_tv35 = _1014_v1;
            let _index209 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1014_v1).length));
            (_1014_v1)[_index209] = (_dafny.ZERO).minus(((function (_pat_let36_0) {
              return function (_1027_dt__update__tmp_h0) {
                return function (_pat_let37_0) {
                  return function (_1028_dt__update_hcf22_h0) {
                    return _module.D3.create_DC11(_1028_dt__update_hcf22_h0);
                  }(_pat_let37_0);
                }((_dafny.ZERO).minus((_pat_let_tv35)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_pat_let_tv34).length))]));
              }(_pat_let36_0);
            }(_1025_v12)).dtor_cf22).plus((p1).minus(new BigNumber((_1026_v13).length))));
          }
        }
      }
      let _hi3 = p1;
      for (let _1029_i1 = p1; _1029_i1.isLessThan(_hi3); _1029_i1 = _1029_i1.plus(_dafny.ONE)) {
        let _1030_v14;
        _1030_v14 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1012_v0);
        let _1031_v15;
        _1031_v15 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_1030_v14).length)));
        let _1032_v16;
        _1032_v16 = _dafny.Seq.of(_1012_v0, _1012_v0);
        let _1033_v17;
        _1033_v17 = _dafny.Set.fromElements(_1012_v0);
        let _1034_v18;
        let _nw190 = Array((new BigNumber(28)).toNumber());
        _nw190[(_dafny.ZERO).toNumber()] = (_1031_v15).IsDisjointFrom(_dafny.MultiSet.fromElements(_module.__default.fm0(p1, globalState), _1029_i1));
        _nw190[(_dafny.ONE).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(2)).toNumber()] = (_module.__default.fm0(p1, globalState)).isLessThan(_1029_i1);
        _nw190[(new BigNumber(3)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(4)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(5)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(6)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(7)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(8)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(9)).toNumber()] = (_1029_i1).isLessThanOrEqualTo(p1);
        _nw190[(new BigNumber(10)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(11)).toNumber()] = (_1029_i1).isLessThanOrEqualTo(p1);
        _nw190[(new BigNumber(12)).toNumber()] = !(_1012_v0);
        _nw190[(new BigNumber(13)).toNumber()] = _dafny.Seq.contains(_1032_v16, _1012_v0);
        _nw190[(new BigNumber(14)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(15)).toNumber()] = _module.__default.fm12(!(_1012_v0), _1012_v0, globalState);
        _nw190[(new BigNumber(16)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(17)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(18)).toNumber()] = !(_1012_v0) || (_1012_v0);
        _nw190[(new BigNumber(19)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(20)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(21)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(22)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(23)).toNumber()] = (_1012_v0) === (_1012_v0);
        _nw190[(new BigNumber(24)).toNumber()] = !(true);
        _nw190[(new BigNumber(25)).toNumber()] = (!(_1012_v0)) || (_1012_v0);
        _nw190[(new BigNumber(26)).toNumber()] = _1012_v0;
        _nw190[(new BigNumber(27)).toNumber()] = (_1033_v17).IsSubsetOf(_1033_v17);
        _1034_v18 = _nw190;
        let _index210 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1034_v18).length));
        (_1034_v18)[_index210] = false;
        let _1035_v19;
        _1035_v19 = _dafny.Set.fromElements(p1);
        let _index211 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1034_v18).length));
        let _rhs210 = _module.__default.fm12(_1012_v0, (_1029_i1).isLessThan(new BigNumber((_1035_v19).length)), globalState);
        let _rhs211 = false;
        let _rhs212 = p1;
        let _lhs191 = _1034_v18;
        let _lhs192 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1034_v18).length));
        let _lhs193 = globalState;
        r0 = _rhs210;
        _lhs191[_lhs192] = _rhs211;
        _lhs193.f10 = _rhs212;
        let _1036_v20;
        let _init21 = ((_1037_v14) => function (_1038_i2) {
          return _1037_v14;
        })(_1030_v14);
        let _nw191 = Array((new BigNumber(15)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw191.length); _i0_21++) {
          _nw191[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _1036_v20 = _nw191;
        let _index212 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1036_v20).length));
        (_1036_v20)[_index212] = _module.__default.fm18(p1, _1012_v0, globalState);
        let _index213 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1036_v20).length));
        (_1036_v20)[_index213] = _1030_v14;
        let _1039_v21;
        let _init22 = ((_1040_p0) => function (_1041_i3) {
          return _1040_p0;
        })(p0);
        let _nw192 = Array((new BigNumber(16)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw192.length); _i0_22++) {
          _nw192[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _1039_v21 = _nw192;
        let _1042_v22;
        _1042_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-100),p0);
        let _index214 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_1039_v21).length));
        (_1039_v21)[_index214] = _dafny.Seq.update((((_1042_v22).contains(_1029_i1)) ? ((_1042_v22).get(_1029_i1)) : (p0)), _module.__default.safeIndex(p1, new BigNumber(((((_1042_v22).contains(_1029_i1)) ? ((_1042_v22).get(_1029_i1)) : (p0))).length)), _module.__default.fm11(p1, globalState));
        let _1043_v23;
        _1043_v23 = new _dafny.CodePoint('t'.codePointAt(0));
        let _index215 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_1039_v21).length));
        (_1039_v21)[_index215] = _dafny.Seq.Concat(_dafny.Seq.of(_1043_v23), p0);
        let _1044_v24;
        let _nw193 = new _module.C0();
        _nw193.__ctor(_module.__default.fm12((_1034_v18)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_1034_v18).length))], (_1034_v18)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_1034_v18).length))], globalState));
        _1044_v24 = _nw193;
      }
      let _hi4 = p1;
      for (let _1045_i4 = new BigNumber((p0).length); _1045_i4.isLessThan(_hi4); _1045_i4 = _1045_i4.plus(_dafny.ONE)) {
        let _1046_v25;
        _1046_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1012_v0,_1012_v0);
        _1046_v25 = (_1046_v25).update(_1012_v0, false);
        let _1047_v26;
        _1047_v26 = new _dafny.CodePoint('q'.codePointAt(0));
        let _1048_v27;
        _1048_v27 = _dafny.MultiSet.fromElements(_1047_v26, _1047_v26, _1047_v26);
        r1 = ((_1012_v0) ? (new BigNumber((_1048_v27).cardinality())) : (new BigNumber(331)));
        let _1049_v28;
        let _nw194 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _1049_v28 = _nw194;
        let _1050_v29;
        _1050_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1049_v28,true);
        let _1051_v30;
        _1051_v30 = _dafny.MultiSet.fromElements(_1012_v0);
        _1050_v29 = (_1050_v29).update(_1049_v28, (_1051_v30).IsDisjointFrom(_1051_v30));
        let _1052_v31;
        _1052_v31 = _dafny.Seq.of(_1012_v0);
        if ((_1052_v31)[_module.__default.safeIndex(_1045_i4, new BigNumber((_1052_v31).length))]) {
          (globalState).f10 = (p1).plus(_module.__default.fm0(_1045_i4, globalState));
          let _1053_v32;
          let _nw195 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
          _1053_v32 = _nw195;
          let _1054_v33;
          _1054_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1053_v32,(p1).multipliedBy(_1045_i4));
          (globalState).f10 = (_dafny.ZERO).minus((((_1054_v33).contains(_1053_v32)) ? ((_1054_v33).get(_1053_v32)) : ((_1045_i4).plus(_1045_i4))));
          let _index216 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1049_v28).length));
          (_1049_v28)[_index216] = p1;
          let _index217 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1049_v28).length));
          (_1049_v28)[_index217] = p1;
          (globalState).f22 = !(_module.__default.fm12(!(_1012_v0), _1012_v0, globalState));
          let _1055_v34;
          let _init23 = ((_1056_v25) => function (_1057_i5) {
            return _1056_v25;
          })(_1046_v25);
          let _nw196 = Array((new BigNumber(27)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw196.length); _i0_23++) {
            _nw196[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _1055_v34 = _nw196;
          let _1058_v35;
          _1058_v35 = _dafny.Seq.of(_1046_v25, _1046_v25, _1046_v25, _1046_v25, _1046_v25);
          let _index218 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1055_v34).length));
          (_1055_v34)[_index218] = (_1058_v35)[_module.__default.safeIndex(p1, new BigNumber((_1058_v35).length))];
          let _index219 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1055_v34).length));
          (_1055_v34)[_index219] = ((_module.__default.fm12(_1012_v0, _1012_v0, globalState)) ? (_1046_v25) : (_1046_v25));
        } else {
          let _1059_v36;
          let _init24 = ((_1060_v0) => function (_1061_i6) {
            return _1060_v0;
          })(_1012_v0);
          let _nw197 = Array((new BigNumber(7)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw197.length); _i0_24++) {
            _nw197[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _1059_v36 = _nw197;
          let _index220 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1059_v36).length));
          (_1059_v36)[_index220] = _1012_v0;
          let _index221 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1059_v36).length));
          (_1059_v36)[_index221] = true;
          _1047_v26 = _1047_v26;
          let _1062_v37;
          _1062_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1047_v26,_1012_v0);
          let _1063_v38;
          _1063_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1045_i4,true);
          let _index222 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1059_v36).length));
          (_1059_v36)[_index222] = (_module.__default.fm12((_1059_v36)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_1059_v36).length))], (((_1062_v37).contains(_1047_v26)) ? ((_1062_v37).get(_1047_v26)) : ((((_1063_v38).contains(p1)) ? ((_1063_v38).get(p1)) : (true)))), globalState)) || (_1012_v0);
          _1012_v0 = true;
          let _1064_v39;
          let _nw198 = Array((new BigNumber(15)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1064_v39 = _nw198;
          let _index223 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_1064_v39).length));
          (_1064_v39)[_index223] = _1047_v26;
          let _1065_v40;
          _1065_v40 = _module.D4.create_DC13((p0)[_module.__default.safeIndex(_1045_i4, new BigNumber((p0).length))]);
          let _index224 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_1064_v39).length));
          (_1064_v39)[_index224] = (_1065_v40).dtor_cf24;
        }
      }
      (globalState).f13 = !(_1012_v0);
      let _1066_v41;
      _1066_v41 = _dafny.Seq.of(false);
      if (((_1066_v41)[_module.__default.safeIndex(p1, new BigNumber((_1066_v41).length))]) === (_module.__default.fm12(_1012_v0, _1012_v0, globalState))) {
        let _1067_v42;
        _1067_v42 = _module.D0.create_DC2(p1, _1012_v0);
        let _1068_v43;
        _1068_v43 = _dafny.MultiSet.fromElements(_1067_v42);
        (globalState).f22 = (_1068_v43).IsProperSubsetOf(_1068_v43);
        (globalState).f10 = ((_1012_v0) ? (p1) : (p1));
        let _1069_v44;
        _1069_v44 = _module.D3.create_DC11((p1).plus(p1));
        let _source17 = _1069_v44;
        if (_source17.is_DC11) {
          let _1070___mcc_h0 = (_source17).cf22;
          let _1071_cf22 = _1070___mcc_h0;
          let _1072_v45;
          _1072_v45 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm11(_1071_cf22, globalState),_1066_v41);
          _1072_v45 = _1072_v45;
          let _1073_v46;
          _1073_v46 = _dafny.Map.Empty.slice().updateUnsafe(p1,false);
          let _1074_v47;
          _1074_v47 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(733)), ((_1075_p1) => function (_1076_i7) {
            return _1075_p1;
          })(p1)), _module.__default.fm14(_1012_v0, _1071_cf22, _module.__default.fm0(_1071_cf22, globalState), globalState)),_1073_v46);
          let _1077_v48;
          _1077_v48 = _dafny.Seq.of(p1);
          _1074_v47 = (_1074_v47).update(_1077_v48, _1073_v46);
          let _1078_v49;
          let _nw199 = new _module.C0();
          _nw199.__ctor(_1012_v0);
          _1078_v49 = _nw199;
          let _1079_v50;
          _1079_v50 = _dafny.Map.Empty.slice().updateUnsafe(((_1012_v0) ? (_module.__default.fm12(_1012_v0, _1012_v0, globalState)) : (_1012_v0)),((_1012_v0) ? (new BigNumber((_dafny.Seq.of(_1078_v49, _1078_v49)).length)) : ((_this).fm8(p1, _1071_cf22, p1, globalState))));
          _1079_v50 = (_1079_v50).update(!(_1012_v0), new BigNumber((_dafny.Seq.Concat(p0, p0)).length));
          let _1080_v51;
          _1080_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1012_v0,(_1078_v49).f32);
          let _1081_v52;
          _1081_v52 = _1080_v51;
          let _1082_v53;
          _1082_v53 = _dafny.MultiSet.fromElements(_1071_cf22, _1071_cf22);
          let _1083_v54;
          let _nw200 = Array((new BigNumber(14)).toNumber());
          _nw200[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(false,_1012_v0);
          _nw200[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1012_v0,(_1078_v49).f32);
          _nw200[(new BigNumber(2)).toNumber()] = _1080_v51;
          _nw200[(new BigNumber(3)).toNumber()] = _1080_v51;
          _nw200[(new BigNumber(4)).toNumber()] = _1080_v51;
          _nw200[(new BigNumber(5)).toNumber()] = _1080_v51;
          _nw200[(new BigNumber(6)).toNumber()] = _1080_v51;
          _nw200[(new BigNumber(7)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_1078_v49).f32,(_1078_v49).f32)).Merge(_1080_v51);
          _nw200[(new BigNumber(8)).toNumber()] = _1080_v51;
          _nw200[(new BigNumber(9)).toNumber()] = _1080_v51;
          _nw200[(new BigNumber(10)).toNumber()] = (_1081_v52);
          _nw200[(new BigNumber(11)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_1078_v49).f32,(_1078_v49).f32)).update((_1078_v49).f32, _1012_v0);
          _nw200[(new BigNumber(12)).toNumber()] = _module.__default.fm19(_1082_v53, p1, globalState);
          _nw200[(new BigNumber(13)).toNumber()] = (_1080_v51).Merge(_1080_v51);
          _1083_v54 = _nw200;
          _1083_v54 = _1083_v54;
        } else if (_source17.is_DC12) {
          let _1084___mcc_h1 = (_source17).cf23;
          let _1085_cf23 = _1084___mcc_h1;
          let _1086_v55;
          let _nw201 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _1086_v55 = _nw201;
          let _index225 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1086_v55).length));
          (_1086_v55)[_index225] = p1;
          let _1087_v56;
          _1087_v56 = _dafny.MultiSet.fromElements(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(702)), function (_1088_i8) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          }));
          let _1089_v57;
          let _nw202 = new _module.C0();
          _nw202.__ctor(false);
          _1089_v57 = _nw202;
          let _index226 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1086_v55).length));
          let _rhs213 = (_1066_v41)[_module.__default.safeIndex(new BigNumber(((_1087_v56).Union(_1087_v56)).cardinality()), new BigNumber((_1066_v41).length))];
          let _rhs214 = (new BigNumber((_dafny.Seq.of(_1089_v57)).length)).plus(p1);
          let _rhs215 = (_dafny.ZERO).minus(p1);
          let _rhs216 = p1;
          let _rhs217 = p1;
          let _lhs194 = globalState;
          let _lhs195 = globalState;
          let _lhs196 = globalState;
          let _lhs197 = _1086_v55;
          let _lhs198 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1086_v55).length));
          _lhs194.f13 = _rhs213;
          _lhs195.f10 = _rhs214;
          r1 = _rhs215;
          _lhs196.f10 = _rhs216;
          _lhs197[_lhs198] = _rhs217;
          (globalState).f10 = p1;
          let _1090_v58;
          _1090_v58 = _dafny.Map.Empty.slice().updateUnsafe(false,(_1086_v55)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1086_v55).length))]);
          let _1091_v59;
          _1091_v59 = _dafny.Map.Empty.slice().updateUnsafe(p1,(((_1090_v58).contains(true)) ? ((_1090_v58).get(true)) : (new BigNumber(891))));
          _1091_v59 = (_1091_v59).update(new BigNumber((p0).length), (_1086_v55)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1086_v55).length))]);
          (globalState).f7 = _1012_v0;
        } else {
          let _1092___mcc_h2 = (_source17).cf21;
          let _1093_cf21 = _1092___mcc_h2;
          let _1094_v60;
          _1094_v60 = _dafny.Set.fromElements(new BigNumber((p0).length));
          let _1095_v61;
          _1095_v61 = _dafny.Seq.of((_dafny.ZERO).minus((_this).fm8(p1, new BigNumber((p0).length), new BigNumber(790), globalState)), new BigNumber(100), new BigNumber((_1094_v60).length));
          let _1096_v62;
          _1096_v62 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_1095_v61),_1095_v61);
          let _1097_v63;
          _1097_v63 = _module.D1.create_DC5(p1, _1012_v0, _1012_v0, _module.__default.fm20(p1, _1012_v0, (_dafny.ZERO).minus(p1), globalState));
          let _1098_v64;
          _1098_v64 = _dafny.Seq.of(_module.D1.create_DC5(new BigNumber((p0).length), _1012_v0, _1012_v0, _1096_v62), _1097_v63, _module.D1.create_DC5(p1, !(_1012_v0), !(_1012_v0), _1096_v62), _1097_v63);
          _1098_v64 = _1098_v64;
          let _1099_v65;
          let _nw203 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Set.Empty);
          _1099_v65 = _nw203;
          let _1100_v66;
          _1100_v66 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus(p1));
          let _index227 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_1099_v65).length));
          (_1099_v65)[_index227] = (_1094_v60).Union(function () {
            let _coll30 = new _dafny.Set();
            for (const _compr_30 of (_1100_v66).Keys.Elements) {
              let _1101_v67 = _compr_30;
              if ((_1100_v66).contains(_1101_v67)) {
                _coll30.add((_1101_v67).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("mbi")).length)));
              }
            }
            return _coll30;
          }());
          let _index228 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_1099_v65).length));
          (_1099_v65)[_index228] = _dafny.Set.fromElements(p1);
          let _1102_v68;
          _1102_v68 = new _dafny.CodePoint('d'.codePointAt(0));
          let _1103_v69;
          _1103_v69 = _module.D4.create_DC13(_1102_v68);
          let _1104_v70;
          _1104_v70 = _dafny.Seq.of(_1103_v69);
          r1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1104_v70, _1104_v70), _dafny.Seq.of(_module.D4.create_DC13(_1102_v68)))).length);
          let _1105_v71;
          _1105_v71 = _dafny.Set.fromElements(_1102_v68, _1102_v68);
          r0 = (_1105_v71).IsSubsetOf(_1105_v71);
        }
        (globalState).f22 = (_1012_v0) || (_1012_v0);
        (globalState).f10 = p1;
      } else {
        (globalState).f7 = _1012_v0;
        let _1106_v72;
        let _nw204 = Array((new BigNumber(8)).toNumber());
        _nw204[(_dafny.ZERO).toNumber()] = _1012_v0;
        _nw204[(_dafny.ONE).toNumber()] = _module.__default.fm12(_1012_v0, _1012_v0, globalState);
        _nw204[(new BigNumber(2)).toNumber()] = _1012_v0;
        _nw204[(new BigNumber(3)).toNumber()] = (_1066_v41)[_module.__default.safeIndex(p1, new BigNumber((_1066_v41).length))];
        _nw204[(new BigNumber(4)).toNumber()] = (((_1066_v41)[_module.__default.safeIndex(_module.__default.fm0(p1, globalState), new BigNumber((_1066_v41).length))]) ? (_1012_v0) : (_1012_v0));
        _nw204[(new BigNumber(5)).toNumber()] = _1012_v0;
        _nw204[(new BigNumber(6)).toNumber()] = (!(_1012_v0)) && (_1012_v0);
        _nw204[(new BigNumber(7)).toNumber()] = _1012_v0;
        _1106_v72 = _nw204;
        let _1107_v73;
        _1107_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1012_v0,_1012_v0);
        let _index229 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_1106_v72).length));
        (_1106_v72)[_index229] = (_1107_v73).equals(_1107_v73);
        let _index230 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_1106_v72).length));
        (_1106_v72)[_index230] = _1012_v0;
        let _1108_v74;
        _1108_v74 = _dafny.Seq.UnicodeFromString("rlvkgkw");
        let _rhs218 = _1108_v74;
        let _rhs219 = _dafny.Seq.Concat(p0, _dafny.Seq.UnicodeFromString("opjwmmjub"));
        _1108_v74 = _rhs218;
        _1108_v74 = _rhs219;
        let _1109_v75;
        let _nw205 = new _module.C0();
        _nw205.__ctor((_1106_v72)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_1106_v72).length))]);
        _1109_v75 = _nw205;
        let _1110_v76;
        let _nw206 = new _module.C1();
        _nw206.__ctor();
        _1110_v76 = _nw206;
        let _1111_v77;
        _1111_v77 = new _dafny.CodePoint('h'.codePointAt(0));
        let _1112_v78;
        _1112_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1110_v76,new BigNumber((_dafny.Seq.Concat(_module.__default.fm10(_dafny.Map.Empty.slice().updateUnsafe(p1,p1), _1111_v77, (_1109_v75).f32, globalState), p0)).length));
        _1112_v78 = (_1112_v78).update(_1110_v76, p1);
      }
      let _1113_v79;
      _1113_v79 = _dafny.MultiSet.fromElements(p1);
      let _1114_v80;
      _1114_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1066_v41).length),_1012_v0);
      let _1115_v81;
      _1115_v81 = _dafny.Map.Empty.slice().updateUnsafe((_1113_v79).update((_dafny.ZERO).minus(new BigNumber((_1114_v80).length)), _module.__default.abs(new BigNumber(-908))),_dafny.Seq.Create(_module.__default.abs(new BigNumber(687)), ((_1116_p1) => function (_1117_i9) {
        return _1116_p1;
      })(p1)));
      let _1118_v82;
      _1118_v82 = _module.D1.create_DC5(new BigNumber((_1066_v41).length), (p1).isLessThanOrEqualTo(p1), _1012_v0, _1115_v81);
      let _source18 = _1118_v82;
      if (_source18.is_DC5) {
        let _1119___mcc_h3 = (_source18).cf9;
        let _1120___mcc_h4 = (_source18).cf10;
        let _1121___mcc_h5 = (_source18).cf11;
        let _1122___mcc_h6 = (_source18).cf12;
        let _1123_cf12 = _1122___mcc_h6;
        let _1124_cf11 = _1121___mcc_h5;
        let _1125_cf10 = _1120___mcc_h4;
        let _1126_cf9 = _1119___mcc_h3;
        if (_1125_cf10) {
          let _1127_v83;
          _1127_v83 = _dafny.MultiSet.fromElements(_1124_cf11, _1125_cf10);
          let _1128_v84;
          _1128_v84 = _dafny.Seq.of(_1127_v83, _1127_v83);
          let _1129_v85;
          let _nw207 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _1129_v85 = _nw207;
          let _rhs220 = p1;
          let _rhs221 = _dafny.Seq.update(_dafny.Seq.Concat(_1128_v84, _1128_v84), _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_1129_v85, _1129_v85, _1129_v85)).cardinality()), new BigNumber((_dafny.Seq.Concat(_1128_v84, _1128_v84)).length)), _dafny.MultiSet.fromElements(_1012_v0, _1012_v0));
          _1126_cf9 = _rhs220;
          _1128_v84 = _rhs221;
          let _1130_v86;
          _1130_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1012_v0,_1124_cf11);
          let _1131_v87;
          _1131_v87 = new _dafny.CodePoint('p'.codePointAt(0));
          let _1132_v88;
          _1132_v88 = _dafny.Seq.of(new BigNumber((_1130_v86).length), _1126_cf9, new BigNumber((_dafny.Set.fromElements(_1131_v87)).length), _1126_cf9);
          let _1133_v89;
          _1133_v89 = _dafny.Seq.of(_1132_v88, _1132_v88);
          let _1134_v90;
          _1134_v90 = _dafny.Map.Empty.slice().updateUnsafe(((_1124_cf11) ? ((_1133_v89)[_module.__default.safeIndex(_1126_cf9, new BigNumber((_1133_v89).length))]) : (_1132_v88)),_1124_cf11);
          _1134_v90 = _1134_v90;
          let _1135_v91;
          _1135_v91 = _dafny.MultiSet.fromElements(_1131_v87, _1131_v87);
          let _1136_v92;
          _1136_v92 = _dafny.Set.fromElements(_1126_cf9, new BigNumber((_1135_v91).cardinality()));
          let _1137_v93;
          _1137_v93 = _module.D0.create_DC2(p1, false);
          let _1138_v94;
          _1138_v94 = _dafny.Map.Empty.slice().updateUnsafe(!(_1125_cf10),p1);
          let _1139_v95;
          let _nw208 = Array((new BigNumber(22)).toNumber());
          _nw208[(_dafny.ZERO).toNumber()] = _1125_cf10;
          _nw208[(_dafny.ONE).toNumber()] = _1124_cf11;
          _nw208[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.fromElements(_module.__default.fm11(new BigNumber((_1066_v41).length), globalState))).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1131_v87));
          _nw208[(new BigNumber(3)).toNumber()] = true;
          _nw208[(new BigNumber(4)).toNumber()] = _1125_cf10;
          _nw208[(new BigNumber(5)).toNumber()] = (new BigNumber((_1136_v92).length)).isLessThan(p1);
          _nw208[(new BigNumber(6)).toNumber()] = _1125_cf10;
          _nw208[(new BigNumber(7)).toNumber()] = (((_1130_v86).contains(false)) ? ((_1130_v86).get(false)) : (_1125_cf10));
          _nw208[(new BigNumber(8)).toNumber()] = _1124_cf11;
          _nw208[(new BigNumber(9)).toNumber()] = (_1137_v93).dtor_cf6;
          _nw208[(new BigNumber(10)).toNumber()] = _1012_v0;
          _nw208[(new BigNumber(11)).toNumber()] = _module.__default.fm12(_1012_v0, _1012_v0, globalState);
          _nw208[(new BigNumber(12)).toNumber()] = _module.__default.fm12(_1012_v0, false, globalState);
          _nw208[(new BigNumber(13)).toNumber()] = _module.__default.fm12(_1124_cf11, !(_1124_cf11), globalState);
          _nw208[(new BigNumber(14)).toNumber()] = _1124_cf11;
          _nw208[(new BigNumber(15)).toNumber()] = _1012_v0;
          _nw208[(new BigNumber(16)).toNumber()] = (_1138_v94).equals(_dafny.Map.Empty.slice().updateUnsafe(_1012_v0,_1126_cf9));
          _nw208[(new BigNumber(17)).toNumber()] = _1012_v0;
          _nw208[(new BigNumber(18)).toNumber()] = !(_1012_v0);
          _nw208[(new BigNumber(19)).toNumber()] = _1012_v0;
          _nw208[(new BigNumber(20)).toNumber()] = !(_1113_v79).equals(_1113_v79);
          _nw208[(new BigNumber(21)).toNumber()] = _1124_cf11;
          _1139_v95 = _nw208;
          let _index231 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1139_v95).length));
          (_1139_v95)[_index231] = _1125_cf10;
          let _index232 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1139_v95).length));
          (_1139_v95)[_index232] = !((_1127_v83).IsDisjointFrom(_1127_v83));
          let _1140_v96;
          _1140_v96 = _dafny.Set.fromElements(_1012_v0, _1012_v0);
          let _1141_v97;
          _1141_v97 = _dafny.Map.Empty.slice().updateUnsafe(_1139_v95,_1140_v96);
          let _1142_v98;
          _1142_v98 = _module.D8.create_DC21(_1066_v41);
          let _1143_v99;
          _1143_v99 = _dafny.Map.Empty.slice().updateUnsafe(_1141_v97,_dafny.Seq.Concat((_1142_v98).dtor_cf35, _1066_v41));
          _1143_v99 = (_1143_v99).update(_1141_v97, _1066_v41);
          (globalState).f24 = (_1136_v92).Intersect(_1136_v92);
        } else {
          let _1144_v100;
          let _nw209 = new _module.C2();
          _nw209.__ctor(_dafny.Seq.Concat(p0, p0));
          _1144_v100 = _nw209;
          let _1145_v101;
          let _nw210 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1145_v101 = _nw210;
          let _1146_v102;
          _1146_v102 = new _dafny.CodePoint('i'.codePointAt(0));
          let _index233 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_1145_v101).length));
          (_1145_v101)[_index233] = _1146_v102;
          let _1147_v103;
          let _nw211 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
          _1147_v103 = _nw211;
          let _index234 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_1145_v101).length));
          let _rhs222 = _module.__default.safeModuloInt(new BigNumber((p0).length), p1);
          let _rhs223 = _1144_v100;
          let _rhs224 = _1146_v102;
          let _rhs225 = _1147_v103;
          let _lhs199 = globalState;
          let _lhs200 = _1145_v101;
          let _lhs201 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_1145_v101).length));
          _lhs199.f10 = _rhs222;
          _1144_v100 = _rhs223;
          _lhs200[_lhs201] = _rhs224;
          _1147_v103 = _rhs225;
          let _1148_v104;
          _1148_v104 = _dafny.Seq.of(_1126_cf9, p1);
          (globalState).f10 = ((_dafny.ZERO).minus((_1148_v104)[_module.__default.safeIndex(_1126_cf9, new BigNumber((_1148_v104).length))])).multipliedBy(_1126_cf9);
          r1 = (_dafny.ZERO).minus(p1);
          (globalState).f10 = _1126_cf9;
          let _1149_v105;
          _1149_v105 = _dafny.Map.Empty.slice().updateUnsafe(_1124_cf11,_1146_v102);
          let _1150_v106;
          let _nw212 = Array((new BigNumber(3)).toNumber());
          _nw212[(_dafny.ZERO).toNumber()] = (_1149_v105).Merge(_1149_v105);
          _nw212[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1124_cf11,(_1145_v101)[_module.__default.safeIndex(new BigNumber(930), new BigNumber((_1145_v101).length))]);
          _nw212[(new BigNumber(2)).toNumber()] = _1149_v105;
          _1150_v106 = _nw212;
          let _index235 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_1150_v106).length));
          (_1150_v106)[_index235] = _1149_v105;
          let _1151_v107;
          _1151_v107 = _dafny.Map.Empty.slice().updateUnsafe(true,_1125_cf10);
          let _index236 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_1150_v106).length));
          (_1150_v106)[_index236] = (_1149_v105).Merge(_dafny.Map.Empty.slice().updateUnsafe((((_1151_v107).contains(_1012_v0)) ? ((_1151_v107).get(_1012_v0)) : (_1125_cf10)),_1146_v102));
        }
        let _1152_v108;
        let _nw213 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _1152_v108 = _nw213;
        let _rhs226 = _1152_v108;
        let _rhs227 = (p1).isLessThanOrEqualTo(_1126_cf9);
        let _rhs228 = p1;
        let _rhs229 = _1125_cf10;
        let _rhs230 = _1126_cf9;
        let _lhs202 = globalState;
        let _lhs203 = globalState;
        let _lhs204 = globalState;
        let _lhs205 = globalState;
        _1152_v108 = _rhs226;
        _lhs202.f13 = _rhs227;
        _lhs203.f10 = _rhs228;
        _lhs204.f13 = _rhs229;
        _lhs205.f10 = _rhs230;
        _1114_v80 = (_1114_v80).update(new BigNumber(235), (_1124_cf11) === (_1125_cf10));
        let _1153_v109;
        _1153_v109 = _1152_v108;
        let _source19 = _1152_v108;
        let _1154___mcc_h12 = _source19;
        let _1155_cf28 = _1154___mcc_h12;
        (globalState).f10 = p1;
        let _1156_v110;
        _1156_v110 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_1152_v108)).length),(_dafny.ZERO).minus(_1126_cf9));
        let _1157_v112;
        _1157_v112 = _dafny.MultiSet.fromElements(_1012_v0, _1124_cf11);
        let _1158_v113;
        _1158_v113 = _dafny.Seq.of(_1157_v112, _dafny.MultiSet.fromElements(true), _dafny.MultiSet.FromArray(_1066_v41));
        _1156_v110 = (_1156_v110).update(_module.__default.safeDivisionInt((_this).fm8(new BigNumber((_dafny.Seq.UnicodeFromString("isvwjc")).length), new BigNumber(-69), _1126_cf9, globalState), p1), new BigNumber((function () {
          let _coll31 = new _dafny.Map();
          for (const _compr_31 of (_1158_v113).Elements) {
            let _1159_v111 = _compr_31;
            if (_dafny.Seq.contains(_1158_v113, _1159_v111)) {
              _coll31.push([_1159_v111,_1126_cf9]);
            }
          }
          return _coll31;
        }()).length));
        let _1160_v114;
        let _init25 = ((_1161_cf11) => function (_1162_i10) {
          return _1161_cf11;
        })(_1124_cf11);
        let _nw214 = Array((new BigNumber(5)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw214.length); _i0_25++) {
          _nw214[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _1160_v114 = _nw214;
        let _1163_v115;
        _1163_v115 = _dafny.Map.Empty.slice().updateUnsafe(_1126_cf9,_1160_v114);
        let _1164_v116;
        _1164_v116 = _dafny.Seq.of(_1126_cf9, _1126_cf9);
        _1163_v115 = (_1163_v115).update(new BigNumber((_1164_v116).length), _1160_v114);
        (globalState).f10 = _1126_cf9;
      } else if (_source18.is_DC6) {
        let _1165___mcc_h7 = (_source18).cf13;
        let _1166___mcc_h8 = (_source18).cf14;
        let _1167___mcc_h9 = (_source18).cf15;
        let _1168___mcc_h10 = (_source18).cf16;
        let _1169_cf16 = _1168___mcc_h10;
        let _1170_cf15 = _1167___mcc_h9;
        let _1171_cf14 = _1166___mcc_h8;
        let _1172_cf13 = _1165___mcc_h7;
        let _1173_v117;
        _1173_v117 = _dafny.Seq.UnicodeFromString("esqyiagdp");
        let _1174_v118;
        _1174_v118 = new _dafny.CodePoint('w'.codePointAt(0));
        let _1175_v119;
        _1175_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1012_v0,_dafny.Seq.update(_1173_v117, _module.__default.safeIndex(_1172_cf13, new BigNumber((_1173_v117).length)), _1174_v118));
        _1173_v117 = _dafny.Seq.update(_1173_v117, _module.__default.safeIndex(new BigNumber((((!(_1012_v0)) ? (_1173_v117) : ((((_1175_v119).contains(_1012_v0)) ? ((_1175_v119).get(_1012_v0)) : (p0))))).length), new BigNumber((_1173_v117).length)), _1174_v118);
        (globalState).f13 = ((_1012_v0) ? (_1169_cf16) : ((_module.__default.fm0(_1172_cf13, globalState)).isLessThan(p1)));
        let _1176_v120;
        _1176_v120 = _dafny.Map.Empty.slice().updateUnsafe(!((_1066_v41)[_module.__default.safeIndex(new BigNumber(271), new BigNumber((_1066_v41).length))]),true);
        let _1177_v121;
        _1177_v121 = _dafny.Set.fromElements(_1169_cf16, _1012_v0);
        (globalState).f13 = (((_1176_v120).contains((_1177_v121).IsSubsetOf(_1177_v121))) ? ((_1176_v120).get((_1177_v121).IsSubsetOf(_1177_v121))) : (_1012_v0));
        let _1178_v122;
        let _nw215 = Array((new BigNumber(8)).toNumber()).fill(false);
        _1178_v122 = _nw215;
        let _1179_v123;
        _1179_v123 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber(824), _1172_cf13),_1178_v122);
        _1179_v123 = (_1179_v123).update(new BigNumber((_1176_v120).length), _1178_v122);
      } else {
        let _1180___mcc_h11 = (_source18).cf8;
        let _1181_cf8 = _1180___mcc_h11;
        let _1182_v124;
        _1182_v124 = _dafny.MultiSet.fromElements(_1012_v0);
        let _1183_v125;
        _1183_v125 = _dafny.Seq.of(p1);
        let _1184_v126;
        _1184_v126 = _dafny.Map.Empty.slice().updateUnsafe(!((_1182_v124).IsProperSubsetOf(_1182_v124)),_1183_v125);
        _1184_v126 = _1184_v126;
        let _1185_v127;
        let _init26 = ((_1186_v0) => function (_1187_i11) {
          return _1186_v0;
        })(_1012_v0);
        let _nw216 = Array((new BigNumber(15)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw216.length); _i0_26++) {
          _nw216[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _1185_v127 = _nw216;
        let _index237 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_1185_v127).length));
        (_1185_v127)[_index237] = ((_1183_v125)[_module.__default.safeIndex(p1, new BigNumber((_1183_v125).length))]).isLessThanOrEqualTo(p1);
        let _index238 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_1185_v127).length));
        (_1185_v127)[_index238] = !_dafny.areEqual(_dafny.Seq.of(_1012_v0, _1012_v0), ((_1012_v0) ? (_dafny.Seq.of(_1012_v0)) : (_1066_v41)));
        (globalState).f13 = (_1185_v127)[_module.__default.safeIndex(new BigNumber(345), new BigNumber((_1185_v127).length))];
        let _1188_v128;
        let _nw217 = Array((new BigNumber(5)).toNumber()).fill([]);
        _1188_v128 = _nw217;
        let _1189_v129;
        _1189_v129 = _dafny.Set.fromElements(true);
        let _1190_v130;
        let _nw218 = Array((new BigNumber(9)).toNumber());
        _nw218[(_dafny.ZERO).toNumber()] = _1113_v79;
        _nw218[(_dafny.ONE).toNumber()] = _1113_v79;
        _nw218[(new BigNumber(2)).toNumber()] = _1113_v79;
        _nw218[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(p1, p1);
        _nw218[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(30)), function (_1191_i12) {
          return new BigNumber(961);
        }));
        _nw218[(new BigNumber(5)).toNumber()] = _1113_v79;
        _nw218[(new BigNumber(6)).toNumber()] = _1113_v79;
        _nw218[(new BigNumber(7)).toNumber()] = _1113_v79;
        _nw218[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_1189_v129).length), p1);
        _1190_v130 = _nw218;
        let _index239 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_1188_v128).length));
        (_1188_v128)[_index239] = _1190_v130;
        let _index240 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_1188_v128).length));
        let _nw219 = Array((new BigNumber(11)).toNumber()).fill(_dafny.MultiSet.Empty);
        (_1188_v128)[_index240] = _nw219;
      }
      let _1192_v131;
      let _nw220 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
      _1192_v131 = _nw220;
      let _1193_v132;
      _1193_v132 = _dafny.MultiSet.fromElements(_1192_v131, _1192_v131);
      r0 = (_1193_v132).IsProperSubsetOf((_1193_v132).Union(_1193_v132));
      r1 = (p1).minus(new BigNumber(367));
      let _1194_v133;
      let _nw221 = new _module.C1();
      _nw221.__ctor();
      _1194_v133 = _nw221;
      r2 = _1194_v133;
      return [r0, r1, r2];
    }
    m5(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      (globalState).f22 = false;
      (globalState).f22 = p3;
      let _1195_v0;
      _1195_v0 = _dafny.Set.fromElements(p0, (_dafny.ZERO).minus(p0), p0);
      let _hi5 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(212)), function (_1196_i1) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })).length);
      for (let _1197_i0 = (new BigNumber((_1195_v0).length)).minus(p0); _1197_i0.isLessThan(_hi5); _1197_i0 = _1197_i0.plus(_dafny.ONE)) {
        let _1198_v1;
        _1198_v1 = _dafny.Seq.UnicodeFromString("rfyol");
        _1198_v1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("arqqntbg"), _1198_v1), _1198_v1);
        let _1199_v2;
        let _nw222 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Set.Empty);
        _1199_v2 = _nw222;
        let _1200_v3;
        _1200_v3 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm12(p3, !(p2), globalState),_1199_v2);
        _1200_v3 = (_1200_v3).update(p3, _1199_v2);
        let _rhs231 = p2;
        let _rhs232 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat(_1198_v1, _1198_v1)).length), _1197_i0);
        let _lhs206 = globalState;
        let _lhs207 = globalState;
        _lhs206.f13 = _rhs231;
        _lhs207.f10 = _rhs232;
        let _1201_v4;
        let _init27 = ((_1202_p3) => function (_1203_i2) {
          return _1202_p3;
        })(p3);
        let _nw223 = Array((new BigNumber(23)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw223.length); _i0_27++) {
          _nw223[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1201_v4 = _nw223;
        let _index241 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_1201_v4).length));
        (_1201_v4)[_index241] = (_1195_v0).IsSubsetOf(_1195_v0);
        let _index242 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_1201_v4).length));
        (_1201_v4)[_index242] = p3;
      }
      let _1204_v5;
      _1204_v5 = new _dafny.CodePoint('s'.codePointAt(0));
      let _1205_v7;
      _1205_v7 = _dafny.MultiSet.fromElements(p3);
      let _1206_v8;
      _1206_v8 = _dafny.Seq.of(_1205_v7);
      let _1207_v9;
      _1207_v9 = _module.D1.create_DC4(function () {
  let _coll32 = new _dafny.Map();
  for (const _compr_32 of (_1206_v8).Elements) {
    let _1208_v6 = _compr_32;
    if (_dafny.Seq.contains(_1206_v8, _1208_v6)) {
      _coll32.push([_1208_v6,p3]);
    }
  }
  return _coll32;
}());
      let _1209_v10;
      _1209_v10 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("bk"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("bk")).length)), _1204_v5),_1207_v9);
      let _1210_v11;
      _1210_v11 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1209_v10);
      let _1211_v12;
      _1211_v12 = _dafny.Seq.UnicodeFromString("rwbrc");
      _1210_v11 = (_1210_v11).update(!(_dafny.Seq.IsPrefixOf(_dafny.Seq.update(_1211_v12, _module.__default.safeIndex(p0, new BigNumber((_1211_v12).length)), _1204_v5), _dafny.Seq.UnicodeFromString("wouynh"))), (_1209_v10).update(_1211_v12, _1207_v9));
      let _1212_i3;
      _1212_i3 = _dafny.ZERO;
      L3: {
        while (p1) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1212_i3)) {
              break L3;
            }
            _1212_i3 = (_1212_i3).plus(_dafny.ONE);
            let _1213_v13;
            _1213_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1211_v12).length));
            let _1214_v14;
            _1214_v14 = _dafny.Seq.of(p1, p1);
            _1213_v13 = _module.__default.fm33((p3) && (true), (p2) === (p2), (_1214_v14)[_module.__default.safeIndex(p0, new BigNumber((_1214_v14).length))], p1, globalState);
            let _1215_v15;
            let _nw224 = new _module.C0();
            _nw224.__ctor(p2);
            _1215_v15 = _nw224;
            let _1216_v16;
            let _nw225 = new _module.C0();
            _nw225.__ctor(!(p0).isEqualTo(p0));
            _1216_v16 = _nw225;
            let _1217_v17;
            _1217_v17 = _dafny.MultiSet.fromElements(p0, new BigNumber((_1211_v12).length));
            let _1218_v18;
            _1218_v18 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_1214_v14, _1214_v14),_1217_v17);
            let _1219_v19;
            let _nw226 = new _module.C2();
            _nw226.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(45)), ((_1220_v5) => function (_1221_i4) {
              return _1220_v5;
            })(_1204_v5)));
            _1219_v19 = _nw226;
            let _1222_v20;
            _1222_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1219_v19,p0);
            let _1223_v21;
            _1223_v21 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_1219_v19.f33).length));
            let _rhs233 = new BigNumber((_1218_v18).length);
            let _rhs234 = (((_1222_v20).contains(_1219_v19)) ? ((_1222_v20).get(_1219_v19)) : (p0));
            let _rhs235 = (_dafny.ZERO).minus(new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus(p0))).Merge(_1223_v21)).update((_1216_v16).f32, p0)).length));
            let _lhs208 = globalState;
            let _lhs209 = globalState;
            let _lhs210 = globalState;
            _lhs208.f10 = _rhs233;
            _lhs209.f10 = _rhs234;
            _lhs210.f10 = _rhs235;
          }
        }
      }
      let _1224_v22;
      let _nw227 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _1224_v22 = _nw227;
      let _index243 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_1224_v22).length));
      (_1224_v22)[_index243] = p0;
      let _index244 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_1224_v22).length));
      (_1224_v22)[_index244] = p0;
      r0 = p0;
      r1 = (_1224_v22)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_1224_v22).length))];
      return [r0, r1];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f31 = [];
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    __ctor(f31) {
      let _this = this;
      (_this)._f31 = f31;
      return;
    }
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (((false) ? (false) : (!(true)))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(!(true)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true, true),true));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),false));
      }
    };
    fm2(globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(((((false) ? (false) : (false))) ? ((new BigNumber(973)).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(631))).length)))) : (new BigNumber((((true) ? (_dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(357)), function (_1225_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })))).length))));
    };
    fm5(p0, p1, p2, globalState) {
      let _this = this;
      if (!(!(_dafny.MultiSet.fromElements(new BigNumber(-735))).equals(_dafny.MultiSet.fromElements(new BigNumber(669))))) {
        return true;
      } else if (!(true)) {
        return false;
      } else {
        return false;
      }
    };
    fm6(globalState) {
      let _this = this;
      return new BigNumber(-665);
    };
    m0(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _rhs236 = p3;
      let _rhs237 = p1;
      let _rhs238 = p2;
      let _lhs211 = globalState;
      let _lhs212 = globalState;
      let _lhs213 = globalState;
      _lhs211.f13 = _rhs236;
      _lhs212.f22 = _rhs237;
      _lhs213.f10 = _rhs238;
      let _1226_i0;
      _1226_i0 = _dafny.ZERO;
      L4: {
        while ((p2).isLessThanOrEqualTo(((p1) ? (new BigNumber(363)) : (p2)))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1226_i0)) {
              break L4;
            }
            _1226_i0 = (_1226_i0).plus(_dafny.ONE);
            let _1227_v0;
            _1227_v0 = _dafny.Seq.UnicodeFromString("gjwrovu");
            let _1228_v1;
            _1228_v1 = new _dafny.CodePoint('o'.codePointAt(0));
            let _1229_v2;
            _1229_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_1227_v0, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(993))).length)), new BigNumber((_1227_v0).length)), _1228_v1)).length),_dafny.MultiSet.fromElements(p2));
            _1229_v2 = (_1229_v2).update(p2, (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(979), p0, p0))).update(p0, _module.__default.abs(p2)));
            let _1230_v3;
            _1230_v3 = _dafny.Set.fromElements(p3, p3);
            let _1231_v4;
            let _nw228 = new _module.C3();
            _nw228.__ctor();
            _1231_v4 = _nw228;
            let _1232_v5;
            _1232_v5 = _dafny.Seq.of(_1231_v4);
            let _1233_v6;
            _1233_v6 = _dafny.MultiSet.fromElements(_1230_v3, _1230_v3);
            if (((_dafny.MultiSet.fromElements(_module.__default.fm7(p1, p3, globalState), _1230_v3, _dafny.Set.fromElements(p1))).update(_dafny.Set.fromElements(p3), _module.__default.abs(new BigNumber((_1232_v5).length)))).IsSubsetOf(_1233_v6)) {
              let _1234_v7;
              _1234_v7 = _module.D2.create_DC8(p3, p1);
              let _1235_v8;
              _1235_v8 = _module.D2.create_DC9(_1234_v7);
              let _1236_v9;
              _1236_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1228_v1,_module.__default.safeDivisionInt(p0, new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p0,_1235_v8)).update(p2, _1235_v8)).length)));
              _1236_v9 = (_1236_v9).update(_1228_v1, p2);
              let _1237_v10;
              let _nw229 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
              _1237_v10 = _nw229;
              let _index245 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_1237_v10).length));
              (_1237_v10)[_index245] = p0;
              let _index246 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_1237_v10).length));
              (_1237_v10)[_index246] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), p0);
              _1237_v10 = _1237_v10;
              let _1238_v11;
              _1238_v11 = _dafny.Seq.of(new BigNumber(742), (_this).fm2(globalState), (_1237_v10)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1237_v10).length))]);
              let _1239_v12;
              _1239_v12 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_1238_v11).length));
              let _1240_v13;
              _1240_v13 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p3,p2));
              (globalState).f7 = !((_dafny.Set.fromElements(_1239_v12)).Difference(_1240_v13)).equals(_dafny.Set.fromElements(_1239_v12, _1239_v12, _1239_v12, _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(208)), _1239_v12));
              let _1241_v14;
              let _nw230 = Array((new BigNumber(7)).toNumber()).fill([]);
              _1241_v14 = _nw230;
              let _1242_v15;
              let _nw231 = Array((new BigNumber(22)).toNumber()).fill(false);
              _1242_v15 = _nw231;
              let _index247 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1241_v14).length));
              (_1241_v14)[_index247] = _1242_v15;
              let _index248 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1241_v14).length));
              (_1241_v14)[_index248] = _1242_v15;
            } else {
              let _1243_v16;
              let _init28 = ((_1244_p1) => function (_1245_i1) {
                return _1244_p1;
              })(p1);
              let _nw232 = Array((new BigNumber(15)).toNumber());
              for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw232.length); _i0_28++) {
                _nw232[_i0_28] = _init28(new BigNumber(_i0_28));
              }
              _1243_v16 = _nw232;
              let _index249 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_1243_v16).length));
              (_1243_v16)[_index249] = p3;
              let _index250 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_1243_v16).length));
              (_1243_v16)[_index250] = true;
              let _1246_v17;
              let _init29 = ((_1247_v0) => function (_1248_i2) {
                return (_1248_i2).plus(new BigNumber((_1247_v0).length));
              })(_1227_v0);
              let _nw233 = Array((new BigNumber(14)).toNumber());
              for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw233.length); _i0_29++) {
                _nw233[_i0_29] = _init29(new BigNumber(_i0_29));
              }
              _1246_v17 = _nw233;
              let _index251 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_1246_v17).length));
              (_1246_v17)[_index251] = p2;
              let _index252 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_1246_v17).length));
              let _rhs239 = p0;
              let _rhs240 = !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("xnwqqkdl"), _1228_v1);
              let _rhs241 = p2;
              let _lhs214 = globalState;
              let _lhs215 = globalState;
              let _lhs216 = _1246_v17;
              let _lhs217 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_1246_v17).length));
              _lhs214.f10 = _rhs239;
              _lhs215.f13 = _rhs240;
              _lhs216[_lhs217] = _rhs241;
              let _1249_v18;
              let _nw234 = new _module.C0();
              _nw234.__ctor(false);
              _1249_v18 = _nw234;
              let _1250_v19;
              _1250_v19 = _dafny.Map.Empty.slice().updateUnsafe(false,_1228_v1);
              let _1251_v20;
              _1251_v20 = _dafny.Seq.of(p1);
              let _1252_v21;
              _1252_v21 = _dafny.Seq.of((_1251_v20)[_module.__default.safeIndex(p2, new BigNumber((_1251_v20).length))]);
              _1250_v19 = (_1250_v19).update(!_dafny.areEqual(_1252_v21, _1252_v21), _1228_v1);
              (globalState).f10 = (_dafny.ZERO).minus(p0);
            }
            let _1253_v23;
            _1253_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1228_v1,p2);
            let _1254_v24;
            _1254_v24 = _dafny.MultiSet.fromElements(p1, p1);
            let _1255_v25;
            _1255_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1254_v24,p3);
            let _1256_v26;
            _1256_v26 = _module.D1.create_DC4(_1255_v25);
            let _pat_let_tv36 = _1255_v25;
            let _1257_v27;
            _1257_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
              let _coll33 = new _dafny.Map();
              for (const _compr_33 of (_1253_v23).Keys.Elements) {
                let _1258_v22 = _compr_33;
                if ((_1253_v23).contains(_1258_v22)) {
                  _coll33.push([_1258_v22,p2]);
                }
              }
              return _coll33;
            }()).length),function (_pat_let38_0) {
              return function (_1259_dt__update__tmp_h0) {
                return function (_pat_let39_0) {
                  return function (_1260_dt__update_hcf8_h0) {
                    return _module.D1.create_DC4(_1260_dt__update_hcf8_h0);
                  }(_pat_let39_0);
                }(_pat_let_tv36);
              }(_pat_let38_0);
            }(_1256_v26));
            let _1261_v28;
            _1261_v28 = _dafny.MultiSet.fromElements(_1257_v27, _dafny.Map.Empty.slice().updateUnsafe(p0,_1256_v26), _1257_v27);
            let _1262_v29;
            _1262_v29 = _dafny.Map.Empty.slice().updateUnsafe(false,(_1261_v28).Union(_1261_v28));
            let _1263_v30;
            _1263_v30 = _dafny.Set.fromElements(_1227_v0);
            let _1264_v31;
            _1264_v31 = _dafny.Seq.of(_1257_v27);
            _1262_v29 = (_1262_v29).update(!((_1263_v30).IsProperSubsetOf(_1263_v30)), _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1264_v31, _1264_v31)));
            if (p3) {
              (globalState).f10 = p2;
              let _1265_v32;
              let _init30 = function (_1266_i3) {
                return true;
              };
              let _nw235 = Array((new BigNumber(8)).toNumber());
              for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw235.length); _i0_30++) {
                _nw235[_i0_30] = _init30(new BigNumber(_i0_30));
              }
              _1265_v32 = _nw235;
              let _index253 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_1265_v32).length));
              (_1265_v32)[_index253] = p3;
              let _index254 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_1265_v32).length));
              (_1265_v32)[_index254] = (p3) && (p3);
              let _1267_v33;
              let _nw236 = new _module.C0();
              _nw236.__ctor(p1);
              _1267_v33 = _nw236;
              (globalState).f10 = new BigNumber(383);
              let _1268_v34;
              _1268_v34 = _dafny.Seq.of(((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).fm2(globalState)))).multipliedBy(p2), new BigNumber((((!((_1265_v32)[_module.__default.safeIndex(new BigNumber(49), new BigNumber((_1265_v32).length))])) ? (_1227_v0) : (_1227_v0))).length), _module.__default.safeModuloInt((_dafny.ZERO).minus(p2), p2));
              _1268_v34 = _1268_v34;
            } else {
              let _1269_v35;
              let _nw237 = new _module.C3();
              _nw237.__ctor();
              _1269_v35 = _nw237;
              let _1270_v36;
              _1270_v36 = _dafny.Seq.of(_1227_v0);
              let _1271_v37;
              _1271_v37 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
              let _1272_v38;
              let _nw238 = Array((new BigNumber(16)).toNumber());
              _nw238[(_dafny.ZERO).toNumber()] = _module.__default.fm10(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1227_v0).length),new BigNumber((_1230_v3).length)), _1228_v1, p1, globalState);
              _nw238[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1227_v0, _dafny.Seq.UnicodeFromString("gdhniefl"));
              _nw238[(new BigNumber(2)).toNumber()] = _1227_v0;
              _nw238[(new BigNumber(3)).toNumber()] = _1227_v0;
              _nw238[(new BigNumber(4)).toNumber()] = _1227_v0;
              _nw238[(new BigNumber(5)).toNumber()] = _1227_v0;
              _nw238[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("arq");
              _nw238[(new BigNumber(7)).toNumber()] = _1227_v0;
              _nw238[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(777)), ((_1273_v1) => function (_1274_i4) {
                return _1273_v1;
              })(_1228_v1)), _1227_v0);
              _nw238[(new BigNumber(9)).toNumber()] = _1227_v0;
              _nw238[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat((_1270_v36)[_module.__default.safeIndex(new BigNumber((_1271_v37).length), new BigNumber((_1270_v36).length))], _1227_v0);
              _nw238[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("pl");
              _nw238[(new BigNumber(12)).toNumber()] = _1227_v0;
              _nw238[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("abm");
              _nw238[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(381)), ((_1275_v1) => function (_1276_i5) {
                return _1275_v1;
              })(_1228_v1)), _1227_v0);
              _nw238[(new BigNumber(15)).toNumber()] = _1227_v0;
              _1272_v38 = _nw238;
              let _index255 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1272_v38).length));
              (_1272_v38)[_index255] = _dafny.Seq.Concat(_1227_v0, _1227_v0);
              let _index256 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1272_v38).length));
              (_1272_v38)[_index256] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(914)), ((_1277_v1) => function (_1278_i6) {
                return _1277_v1;
              })(_1228_v1)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(577)), function (_1279_i7) {
                return new _dafny.CodePoint('x'.codePointAt(0));
              }), _1227_v0));
              (globalState).f10 = p0;
              let _1280_v39;
              _1280_v39 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
              (globalState).f10 = new BigNumber(((_1280_v39).update(new BigNumber(591), p1)).length);
              let _1281_v40;
              let _nw239 = new _module.C1();
              _nw239.__ctor();
              _1281_v40 = _nw239;
            }
          }
        }
      }
      let _hi6 = new BigNumber(-238);
      for (let _1282_i8 = p0; _1282_i8.isLessThan(_hi6); _1282_i8 = _1282_i8.plus(_dafny.ONE)) {
        let _1283_v41;
        _1283_v41 = _dafny.MultiSet.fromElements(p1, p3);
        let _1284_v42;
        _1284_v42 = _module.D3.create_DC10(_1283_v41);
        _1284_v42 = ((p3) ? (_1284_v42) : (_1284_v42));
        let _1285_v43;
        let _nw240 = new _module.C1();
        _nw240.__ctor();
        _1285_v43 = _nw240;
        _1285_v43 = _1285_v43;
        let _1286_v44;
        _1286_v44 = _dafny.Seq.UnicodeFromString("v");
        let _1287_v45;
        _1287_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1286_v44,new BigNumber((_1286_v44).length));
        let _1288_v46;
        _1288_v46 = _dafny.Map.Empty.slice().updateUnsafe(p3,p1);
        let _1289_v47;
        _1289_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1282_i8,p2);
        let _1290_v48;
        _1290_v48 = _dafny.MultiSet.fromElements(_module.D2.create_DC7(_1289_v47));
        let _1291_v49;
        _1291_v49 = _dafny.Seq.of(_1290_v48);
        let _1292_v50;
        _1292_v50 = _dafny.Seq.of(new BigNumber((_1288_v46).length), new BigNumber(((_1291_v49)[_module.__default.safeIndex(p2, new BigNumber((_1291_v49).length))]).cardinality()));
        let _1293_v51;
        let _nw241 = Array((new BigNumber(15)).toNumber());
        _nw241[(_dafny.ZERO).toNumber()] = p0;
        _nw241[(_dafny.ONE).toNumber()] = (((_1287_v45).contains(_1286_v44)) ? ((_1287_v45).get(_1286_v44)) : (p2));
        _nw241[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(682), new BigNumber(-113));
        _nw241[(new BigNumber(3)).toNumber()] = _1282_i8;
        _nw241[(new BigNumber(4)).toNumber()] = _module.__default.fm0(_1282_i8, globalState);
        _nw241[(new BigNumber(5)).toNumber()] = (new BigNumber((_1283_v41).cardinality())).multipliedBy(p2);
        _nw241[(new BigNumber(6)).toNumber()] = new BigNumber(584);
        _nw241[(new BigNumber(7)).toNumber()] = p0;
        _nw241[(new BigNumber(8)).toNumber()] = p0;
        _nw241[(new BigNumber(9)).toNumber()] = (_1282_i8).minus((_dafny.ZERO).minus(new BigNumber(-7)));
        _nw241[(new BigNumber(10)).toNumber()] = p2;
        _nw241[(new BigNumber(11)).toNumber()] = p2;
        _nw241[(new BigNumber(12)).toNumber()] = (_1285_v43).fm2(globalState);
        _nw241[(new BigNumber(13)).toNumber()] = (_1292_v50)[_module.__default.safeIndex(p0, new BigNumber((_1292_v50).length))];
        _nw241[(new BigNumber(14)).toNumber()] = p0;
        _1293_v51 = _nw241;
        let _index257 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1293_v51).length));
        (_1293_v51)[_index257] = ((p1) ? (p2) : (p2));
        let _index258 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1293_v51).length));
        let _rhs242 = (_1282_i8).plus(p0);
        let _rhs243 = p2;
        let _lhs218 = globalState;
        let _lhs219 = _1293_v51;
        let _lhs220 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1293_v51).length));
        _lhs218.f10 = _rhs242;
        _lhs219[_lhs220] = _rhs243;
        if (!(!(p1))) {
          (globalState).f10 = _module.__default.safeDivisionInt((_1293_v51)[_module.__default.safeIndex(new BigNumber(233), new BigNumber((_1293_v51).length))], _1282_i8);
          (globalState).f7 = !(p1);
          let _1294_v52;
          _1294_v52 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1295_v53;
          _1295_v53 = _module.D4.create_DC13(_1294_v52);
          let _1296_v54;
          _1296_v54 = _module.D4.create_DC15(_1295_v53);
          let _1297_v55;
          _1297_v55 = _module.D4.create_DC15(_1295_v53);
          let _1298_v56;
          _1298_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1282_i8,_1297_v55);
          _1298_v56 = _1298_v56;
          let _1299_v57;
          let _nw242 = Array((new BigNumber(26)).toNumber()).fill(false);
          _1299_v57 = _nw242;
          let _1300_v58;
          _1300_v58 = _dafny.Set.fromElements(_1299_v57, _1299_v57, _1299_v57);
          (globalState).f22 = !(_1300_v58).contains(_1299_v57);
          _1286_v44 = _1286_v44;
        } else {
          let _1301_v59;
          _1301_v59 = _dafny.Set.fromElements(_1282_i8);
          let _1302_v60;
          _1302_v60 = _module.D8.create_DC23((_1301_v59).IsProperSubsetOf(_1301_v59));
          _1302_v60 = _1302_v60;
          (globalState).f13 = p1;
          let _1303_v61;
          let _nw243 = new _module.C2();
          _nw243.__ctor(_1286_v44);
          _1303_v61 = _nw243;
          (globalState).f13 = p1;
          let _index259 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1293_v51).length));
          (_1293_v51)[_index259] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(728)), function (_1304_i9) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          })).length);
        }
      }
      let _1305_v62;
      let _nw244 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
      _1305_v62 = _nw244;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1305_v62).length))) {
        let _1306_i10 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1306_i10)) && ((_1306_i10).isLessThan(new BigNumber((_1305_v62).length))))) {
          (_1305_v62)[(_1306_i10)] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(704)), function (_1307_i11) {
            return _module.D4.create_DC13(new _dafny.CodePoint('n'.codePointAt(0)));
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(83)), function (_1308_i12) {
            return _module.D4.create_DC13(new _dafny.CodePoint('q'.codePointAt(0)));
          }));
        }
      }
      if (((!(!(p0).isEqualTo(p2))) ? (!(p3) || (p1)) : (p3))) {
        (globalState).f10 = p2;
        let _1309_v63;
        let _nw245 = new _module.C0();
        _nw245.__ctor(p1);
        _1309_v63 = _nw245;
        let _1310_v64;
        let _init31 = function (_1311_i13) {
          return _dafny.Seq.of(false);
        };
        let _nw246 = Array((new BigNumber(25)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw246.length); _i0_31++) {
          _nw246[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1310_v64 = _nw246;
        let _1312_v65;
        _1312_v65 = _dafny.Seq.of(_1310_v64);
        let _1313_v66;
        _1313_v66 = new _dafny.CodePoint('a'.codePointAt(0));
        let _1314_v67;
        _1314_v67 = _dafny.Map.Empty.slice().updateUnsafe((_1312_v65)[_module.__default.safeIndex(p2, new BigNumber((_1312_v65).length))],_1313_v66);
        _1314_v67 = (_1314_v67).update(_1310_v64, _1313_v66);
        (globalState).f7 = p3;
        let _1315_v68;
        _1315_v68 = _dafny.MultiSet.fromElements(p3, true, p3, p1, p3);
        let _1316_v69;
        _1316_v69 = _dafny.Seq.of((((_1309_v63).f32) ? (_1315_v68) : (_1315_v68)));
        _1316_v69 = _1316_v69;
      } else {
        let _1317_v70;
        _1317_v70 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
        let _1318_v71;
        _1318_v71 = _dafny.MultiSet.fromElements(p3, p1, p1, (((_1317_v70).contains(p2)) ? ((_1317_v70).get(p2)) : (p3)), p3);
        let _1319_v72;
        _1319_v72 = _dafny.Set.fromElements(_1318_v71, _1318_v71, _dafny.MultiSet.fromElements(!(p3), p1));
        _1319_v72 = _1319_v72;
        let _pat_let_tv37 = _1318_v71;
        let _pat_let_tv38 = p1;
        let _pat_let_tv39 = p2;
        let _source20 = function (_pat_let40_0) {
          return function (_1320_dt__update__tmp_h1) {
            return function (_pat_let41_0) {
              return function (_1321_dt__update_hcf21_h0) {
                return _module.D3.create_DC10(_1321_dt__update_hcf21_h0);
              }(_pat_let41_0);
            }((_pat_let_tv37).update(_pat_let_tv38, _module.__default.abs((_dafny.ZERO).minus(_pat_let_tv39))));
          }(_pat_let40_0);
        }(_module.D3.create_DC10(_module.__default.fm24(p2, p1, globalState)));
        if (_source20.is_DC11) {
          let _1322___mcc_h0 = (_source20).cf22;
          let _1323_cf22 = _1322___mcc_h0;
          let _1324_v73;
          let _init32 = ((_1325_p1) => function (_1326_i14) {
            return _1325_p1;
          })(p1);
          let _nw247 = Array((new BigNumber(10)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw247.length); _i0_32++) {
            _nw247[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1324_v73 = _nw247;
          let _1327_v74;
          _1327_v74 = _dafny.Set.fromElements(_1324_v73, _1324_v73);
          let _1328_v75;
          _1328_v75 = _dafny.Seq.of(_1327_v74, _1327_v74, _1327_v74);
          (globalState).f10 = new BigNumber(((_1328_v75)[_module.__default.safeIndex(p2, new BigNumber((_1328_v75).length))]).length);
          let _1329_v76;
          let _nw248 = new _module.C0();
          _nw248.__ctor(p3);
          _1329_v76 = _nw248;
          _1329_v76 = _1329_v76;
          (globalState).f19 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1330_v77;
          let _nw249 = new _module.C1();
          _nw249.__ctor();
          _1330_v77 = _nw249;
          let _1331_v78;
          _1331_v78 = _dafny.Set.fromElements(_1323_cf22);
          let _1332_v79;
          _1332_v79 = _dafny.Seq.of((_1331_v78).equals(_1331_v78));
          let _1333_v80;
          _1333_v80 = _dafny.Seq.of(_1323_cf22, _1323_cf22);
          let _1334_v81;
          _1334_v81 = _dafny.Seq.UnicodeFromString("pwm");
          let _rhs244 = _1330_v77;
          let _rhs245 = (_1333_v80)[_module.__default.safeIndex(p0, new BigNumber((_1333_v80).length))];
          let _rhs246 = !((_this).fm2(globalState)).isEqualTo(new BigNumber((_1334_v81).length));
          let _rhs247 = ((false) ? (_dafny.Seq.Concat(_dafny.Seq.update(_1332_v79, _module.__default.safeIndex(_1323_cf22, new BigNumber((_1332_v79).length)), (_1329_v76).f32), _1332_v79)) : (_1332_v79));
          let _lhs221 = globalState;
          let _lhs222 = globalState;
          _1330_v77 = _rhs244;
          _lhs221.f10 = _rhs245;
          _lhs222.f22 = _rhs246;
          _1332_v79 = _rhs247;
        } else if (_source20.is_DC12) {
          let _1335___mcc_h1 = (_source20).cf23;
          let _1336_cf23 = _1335___mcc_h1;
          let _1337_v82;
          _1337_v82 = new _dafny.CodePoint('j'.codePointAt(0));
          let _1338_v83;
          _1338_v83 = _module.D8.create_DC22(new BigNumber(652), p1, _1337_v82);
          _1336_cf23 = !((_dafny.ZERO).minus((_1338_v83).dtor_cf36)).isEqualTo(p2);
          let _1339_v84;
          _1339_v84 = _dafny.Set.fromElements(new BigNumber(708));
          let _1340_v85;
          _1340_v85 = _dafny.Seq.UnicodeFromString("tmjx");
          let _1341_v86;
          _1341_v86 = _dafny.MultiSet.fromElements(p2, new BigNumber((_1339_v84).length), _module.__default.safeModuloInt(new BigNumber((_module.__default.fm39(globalState)).length), p2), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ywd"), _1340_v85)).length)));
          let _1342_v87;
          _1342_v87 = _dafny.Seq.of(_1340_v85, _dafny.Seq.UnicodeFromString("vwntchfj"));
          let _1343_v88;
          _1343_v88 = _dafny.Seq.of((_this).fm5(p3, _1342_v87, false, globalState), _1336_cf23);
          let _1344_v89;
          _1344_v89 = _dafny.Seq.of(new BigNumber((_1343_v88).length), new BigNumber(184));
          _1341_v86 = (_1341_v86).Union(_dafny.MultiSet.FromArray(_1344_v89));
          (globalState).f10 = _module.__default.safeDivisionInt((p0).minus(p0), new BigNumber(774));
          let _1345_v90;
          let _nw250 = Array((new BigNumber(11)).toNumber());
          _nw250[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(433)), ((_1346_v82) => function (_1347_i15) {
            return _1346_v82;
          })(_1337_v82));
          _nw250[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_1340_v85, _module.__default.safeIndex(p2, new BigNumber((_1340_v85).length)), _1337_v82), _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(465)), ((_1348_v82) => function (_1349_i16) {
            return _1348_v82;
          })(_1337_v82)), _module.__default.safeIndex(new BigNumber((_1340_v85).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(465)), ((_1350_v82) => function (_1351_i16) {
            return _1350_v82;
          })(_1337_v82))).length)), _1337_v82)).length), new BigNumber((_dafny.Seq.update(_1340_v85, _module.__default.safeIndex(p2, new BigNumber((_1340_v85).length)), _1337_v82)).length)), new _dafny.CodePoint('h'.codePointAt(0)));
          _nw250[(new BigNumber(2)).toNumber()] = _1340_v85;
          _nw250[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(14)), ((_1352_v82) => function (_1353_i17) {
            return _1352_v82;
          })(_1337_v82));
          _nw250[(new BigNumber(4)).toNumber()] = _1340_v85;
          _nw250[(new BigNumber(5)).toNumber()] = _1340_v85;
          _nw250[(new BigNumber(6)).toNumber()] = _module.__default.fm10(_module.__default.fm33(p3, p3, true, true, globalState), _1337_v82, p3, globalState);
          _nw250[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("pyctsgo");
          _nw250[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_1340_v85, _module.__default.safeIndex(new BigNumber((_1340_v85).length), new BigNumber((_1340_v85).length)), _1337_v82);
          _nw250[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("qdgk");
          _nw250[(new BigNumber(10)).toNumber()] = _1340_v85;
          _1345_v90 = _nw250;
          let _1354_v91;
          _1354_v91 = _dafny.Map.Empty.slice().updateUnsafe(_1345_v90,_1337_v82);
          _1354_v91 = (_1354_v91).update(_1345_v90, new _dafny.CodePoint('u'.codePointAt(0)));
        } else {
          let _1355___mcc_h2 = (_source20).cf21;
          let _1356_cf21 = _1355___mcc_h2;
          let _1357_v92;
          let _nw251 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
          _1357_v92 = _nw251;
          let _1358_v93;
          _1358_v93 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _1359_v94;
          _1359_v94 = (_1358_v93).update(p1, p1);
          let _index260 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1357_v92).length));
          (_1357_v92)[_index260] = _1359_v94;
          let _index261 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1357_v92).length));
          (_1357_v92)[_index261] = _1358_v93;
          let _1360_v95;
          _1360_v95 = _dafny.Seq.UnicodeFromString("tdhv");
          let _1361_v96;
          let _nw252 = new _module.C2();
          _nw252.__ctor(_1360_v95);
          _1361_v96 = _nw252;
          let _1362_v97;
          _1362_v97 = _module.D8.create_DC21(_dafny.Seq.of(false, p1, !(p3)));
          _1362_v97 = _1362_v97;
          let _1363_v98;
          let _nw253 = new _module.C1();
          _nw253.__ctor();
          _1363_v98 = _nw253;
        }
        if (p3) {
          let _1364_v99;
          _1364_v99 = _dafny.Seq.UnicodeFromString("qvt");
          let _1365_v100;
          let _nw254 = new _module.C2();
          _nw254.__ctor(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rlamqeya"), _1364_v99));
          _1365_v100 = _nw254;
          let _1366_v101;
          let _nw255 = Array((new BigNumber(11)).toNumber()).fill([]);
          _1366_v101 = _nw255;
          let _1367_v102;
          _1367_v102 = _dafny.Seq.of(p1, p3);
          let _1368_v103;
          _1368_v103 = _dafny.Seq.of(_1365_v100.f33);
          let _1369_v104;
          _1369_v104 = _dafny.Seq.of(p0, new BigNumber((_1368_v103).length));
          let _1370_v105;
          _1370_v105 = _dafny.MultiSet.fromElements(p2);
          let _1371_v106;
          _1371_v106 = _dafny.Map.Empty.slice().updateUnsafe((_1367_v102)[_module.__default.safeIndex(p2, new BigNumber((_1367_v102).length))],(_1369_v104)[_module.__default.safeIndex(new BigNumber((_1370_v105).cardinality()), new BigNumber((_1369_v104).length))]);
          let _1372_v107;
          _1372_v107 = _dafny.Map.Empty.slice().updateUnsafe((_1371_v106).equals(_1371_v106),_1366_v101);
          _1366_v101 = (((_1372_v107).contains((_1367_v102)[_module.__default.safeIndex(p2, new BigNumber((_1367_v102).length))])) ? ((_1372_v107).get((_1367_v102)[_module.__default.safeIndex(p2, new BigNumber((_1367_v102).length))])) : (_1366_v101));
          let _1373_v108;
          _1373_v108 = _dafny.Set.fromElements(p1);
          let _1374_v109;
          _1374_v109 = _module.D8.create_DC23(p1);
          let _1375_v110;
          _1375_v110 = _module.D8.create_DC24(_1374_v109);
          let _1376_v111;
          _1376_v111 = new _dafny.CodePoint('r'.codePointAt(0));
          let _rhs248 = (_1373_v108).Difference(_1373_v108);
          let _rhs249 = _1376_v111;
          let _rhs250 = p3;
          let _rhs251 = _1375_v110;
          let _rhs252 = !(!(p1));
          let _lhs223 = globalState;
          let _lhs224 = globalState;
          let _lhs225 = globalState;
          _1373_v108 = _rhs248;
          _lhs223.f19 = _rhs249;
          _lhs224.f13 = _rhs250;
          _1375_v110 = _rhs251;
          _lhs225.f7 = _rhs252;
          (globalState).f19 = _1376_v111;
          let _1377_v112;
          _1377_v112 = _dafny.Map.Empty.slice().updateUnsafe(_1365_v100.f33,p2);
          (globalState).f10 = (_dafny.ZERO).minus((new BigNumber((_1377_v112).length)).minus(p0));
        } else {
          let _1378_v113;
          let _nw256 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _1378_v113 = _nw256;
          let _1379_v114;
          _1379_v114 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _index262 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1378_v113).length));
          (_1378_v113)[_index262] = (new BigNumber((_1379_v114).length)).multipliedBy(p2);
          let _1380_v115;
          _1380_v115 = _1379_v114;
          let _1381_v116;
          _1381_v116 = _dafny.Seq.UnicodeFromString("esrblhn");
          let _index263 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_1378_v113).length));
          (_1378_v113)[_index263] = _module.__default.safeModuloInt(p0, new BigNumber((_1381_v116).length));
          let _index264 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1378_v113).length));
          (_1378_v113)[_index264] = p2;
          let _index265 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1378_v113).length));
          let _index266 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_1378_v113).length));
          let _index267 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1378_v113).length));
          let _rhs253 = ((p0).multipliedBy(p2)).multipliedBy(new BigNumber(373));
          let _rhs254 = (_this).fm5(p3, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ochwdrjm"), _1381_v116), p1, globalState);
          let _rhs255 = _1380_v115;
          let _rhs256 = (_dafny.ZERO).minus(((_this).fm6(globalState)).multipliedBy(new BigNumber(((_1317_v70).Merge(_1317_v70)).length)));
          let _rhs257 = new BigNumber((_1381_v116).length);
          let _lhs226 = _1378_v113;
          let _lhs227 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1378_v113).length));
          let _lhs228 = globalState;
          let _lhs229 = _1378_v113;
          let _lhs230 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_1378_v113).length));
          let _lhs231 = _1378_v113;
          let _lhs232 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1378_v113).length));
          _lhs226[_lhs227] = _rhs253;
          _lhs228.f22 = _rhs254;
          _1380_v115 = _rhs255;
          _lhs229[_lhs230] = _rhs256;
          _lhs231[_lhs232] = _rhs257;
          let _1382_v117;
          let _init33 = ((_1383_p1) => function (_1384_i18) {
            return _1383_p1;
          })(p1);
          let _nw257 = Array((new BigNumber(8)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw257.length); _i0_33++) {
            _nw257[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1382_v117 = _nw257;
          let _1385_v118;
          let _nw258 = Array((new BigNumber(6)).toNumber());
          _nw258[(_dafny.ZERO).toNumber()] = _1382_v117;
          _nw258[(_dafny.ONE).toNumber()] = _1382_v117;
          _nw258[(new BigNumber(2)).toNumber()] = _1382_v117;
          _nw258[(new BigNumber(3)).toNumber()] = _1382_v117;
          _nw258[(new BigNumber(4)).toNumber()] = _1382_v117;
          _nw258[(new BigNumber(5)).toNumber()] = ((p1) ? (_1382_v117) : (_1382_v117));
          _1385_v118 = _nw258;
          let _index268 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_1385_v118).length));
          (_1385_v118)[_index268] = _1382_v117;
          let _1386_v119;
          let _nw259 = new _module.C2();
          _nw259.__ctor(_1381_v116);
          _1386_v119 = _nw259;
          let _1387_v120;
          _1387_v120 = _dafny.Set.fromElements(p1);
          let _1388_v121;
          _1388_v121 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1386_v119);
          let _index269 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_1385_v118).length));
          let _rhs258 = !(_1387_v120).equals(((p3) ? (_1387_v120) : (_1387_v120)));
          let _rhs259 = _1382_v117;
          let _rhs260 = (((_1388_v121).contains(p3)) ? ((_1388_v121).get(p3)) : (_1386_v119));
          let _rhs261 = p3;
          let _lhs233 = globalState;
          let _lhs234 = _1385_v118;
          let _lhs235 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_1385_v118).length));
          let _lhs236 = globalState;
          _lhs233.f13 = _rhs258;
          _lhs234[_lhs235] = _rhs259;
          _1386_v119 = _rhs260;
          _lhs236.f22 = _rhs261;
          (globalState).f10 = _module.__default.safeModuloInt((p2).multipliedBy(p0), p2);
          let _arr0 = (_1385_v118)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_1385_v118).length))];
          let _index270 = _module.__default.safeIndex(new BigNumber(716), new BigNumber(((_1385_v118)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_1385_v118).length))]).length));
          _arr0[_index270] = p3;
          let _arr1 = (_1385_v118)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_1385_v118).length))];
          let _index271 = _module.__default.safeIndex(new BigNumber(716), new BigNumber(((_1385_v118)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_1385_v118).length))]).length));
          _arr1[_index271] = p3;
          (globalState).f10 = p0;
        }
        let _rhs262 = p1;
        let _rhs263 = p1;
        let _lhs237 = globalState;
        let _lhs238 = globalState;
        _lhs237.f13 = _rhs262;
        _lhs238.f13 = _rhs263;
        let _1389_v123;
        _1389_v123 = _dafny.Seq.UnicodeFromString("idmp");
        let _1390_v124;
        _1390_v124 = _dafny.MultiSet.fromElements(_1389_v123);
        (globalState).f10 = new BigNumber((function () {
          let _coll34 = new _dafny.Map();
          for (const _compr_34 of (_1390_v124).Elements) {
            let _1391_v122 = _compr_34;
            if ((_1390_v124).contains(_1391_v122)) {
              _coll34.push([_1391_v122,p1]);
            }
          }
          return _coll34;
        }()).length);
      }
      let _1392_v125;
      let _nw260 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1392_v125 = _nw260;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1392_v125).length))) {
        let _1393_i19 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1393_i19)) && ((_1393_i19).isLessThan(new BigNumber((_1392_v125).length))))) {
          (_1392_v125)[(_1393_i19)] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("roywi"), _dafny.Seq.UnicodeFromString("fllqnoebg")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(919)), function (_1394_i20) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          }));
        }
      }
      return;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let r1 = false;
      r1 = p2;
      let _1395_i0;
      _1395_i0 = _dafny.ZERO;
      L5: {
        while (p2) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1395_i0)) {
              break L5;
            }
            _1395_i0 = (_1395_i0).plus(_dafny.ONE);
            let _1396_v0;
            _1396_v0 = _dafny.Seq.UnicodeFromString("oekpnmql");
            let _1397_v1;
            _1397_v1 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(776));
            let _1398_v2;
            _1398_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.UnicodeFromString("i"));
            let _1399_v3;
            _1399_v3 = _dafny.Set.fromElements(p2, !(p2), p2);
            let _1400_v4;
            _1400_v4 = new _dafny.CodePoint('k'.codePointAt(0));
            let _rhs264 = (((_1398_v2).contains(new BigNumber(-300))) ? ((_1398_v2).get(new BigNumber(-300))) : (_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("phjwav"), _1396_v0), _module.__default.safeIndex(new BigNumber((_1399_v3).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("phjwav"), _1396_v0)).length)), _1400_v4)));
            let _rhs265 = ((p2) ? (_1397_v1) : (_1397_v1));
            _1396_v0 = _rhs264;
            _1397_v1 = _rhs265;
            let _1401_v5;
            _1401_v5 = _dafny.Set.fromElements(_1396_v0, _1396_v0, _1396_v0);
            let _1402_v6;
            _1402_v6 = _dafny.Map.Empty.slice().updateUnsafe((p1).plus((_dafny.ZERO).minus((_this).fm2(globalState))),_1401_v5);
            let _1403_v7;
            _1403_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(990),(_this).fm2(globalState));
            _1402_v6 = (_1402_v6).update(new BigNumber((_module.__default.fm23(_module.D2.create_DC7(_1403_v7), new BigNumber(728), p1, globalState)).length), (_1401_v5).Difference(_1401_v5));
            (globalState).f7 = p2;
            let _1404_v8;
            let _out8;
            _out8 = (_this).m2(globalState);
            _1404_v8 = _out8;
          }
        }
      }
      (globalState).f10 = p1;
      let _1405_v9;
      _1405_v9 = _dafny.MultiSet.fromElements(p2);
      let _1406_v10;
      _1406_v10 = _module.D3.create_DC10(_1405_v9);
      if (((_1406_v10).dtor_cf21).IsDisjointFrom((_1405_v9).Union(_1405_v9))) {
        let _1407_v11;
        let _nw261 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1407_v11 = _nw261;
        let _index272 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_1407_v11).length));
        (_1407_v11)[_index272] = _dafny.Seq.UnicodeFromString("xbbmhsxm");
        let _1408_v12;
        _1408_v12 = _dafny.Seq.UnicodeFromString("oobix");
        let _index273 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_1407_v11).length));
        (_1407_v11)[_index273] = _dafny.Seq.Concat(_1408_v12, _1408_v12);
        let _1409_v14;
        _1409_v14 = _module.D0.create_DC2(p1, !(p2));
        let _1410_v15;
        let _nw262 = Array((new BigNumber(19)).toNumber());
        _nw262[(_dafny.ZERO).toNumber()] = true;
        _nw262[(_dafny.ONE).toNumber()] = p2;
        _nw262[(new BigNumber(2)).toNumber()] = p2;
        _nw262[(new BigNumber(3)).toNumber()] = p2;
        _nw262[(new BigNumber(4)).toNumber()] = false;
        _nw262[(new BigNumber(5)).toNumber()] = p2;
        _nw262[(new BigNumber(6)).toNumber()] = p2;
        _nw262[(new BigNumber(7)).toNumber()] = p2;
        _nw262[(new BigNumber(8)).toNumber()] = false;
        _nw262[(new BigNumber(9)).toNumber()] = p2;
        _nw262[(new BigNumber(10)).toNumber()] = p2;
        _nw262[(new BigNumber(11)).toNumber()] = p2;
        _nw262[(new BigNumber(12)).toNumber()] = p2;
        _nw262[(new BigNumber(13)).toNumber()] = p2;
        _nw262[(new BigNumber(14)).toNumber()] = (_1409_v14).dtor_cf6;
        _nw262[(new BigNumber(15)).toNumber()] = p2;
        _nw262[(new BigNumber(16)).toNumber()] = p2;
        _nw262[(new BigNumber(17)).toNumber()] = p2;
        _nw262[(new BigNumber(18)).toNumber()] = p2;
        _1410_v15 = _nw262;
        let _1411_v16;
        _1411_v16 = _module.D4.create_DC14(p2);
        let _1412_v17;
        _1412_v17 = _module.D4.create_DC15(_1411_v16);
        let _1413_v18;
        _1413_v18 = _module.D7.create_DC20(p0, _1410_v15, _1412_v17, p2, p1);
        let _1414_v19;
        _1414_v19 = _dafny.Seq.of(p2);
        let _1415_v20;
        _1415_v20 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(-839));
        let _1416_v22;
        _1416_v22 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
        let _1417_v23;
        _1417_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(91)), ((_1418_v12) => function (_1419_i1) {
          return new BigNumber((_1418_v12).length);
        })(_1408_v12))).length), new BigNumber((function () {
          let _coll35 = new _dafny.Map();
          for (const _compr_35 of (_1416_v22).Keys.Elements) {
            let _1420_v21 = _compr_35;
            if ((_1416_v22).contains(_1420_v21)) {
              _coll35.push([_module.__default.safeDivisionInt(_1420_v21, new BigNumber(-486)),p2]);
            }
          }
          return _coll35;
        }()).length))).cardinality()),(_this).fm5(p2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(884)), ((_1421_v12) => function (_1422_i2) {
          return _1421_v12;
        })(_1408_v12)), p2, globalState));
        let _1423_v24;
        let _nw263 = Array((new BigNumber(22)).toNumber());
        _nw263[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("axtltt")).length));
        _nw263[(_dafny.ONE).toNumber()] = p1;
        _nw263[(new BigNumber(2)).toNumber()] = p1;
        _nw263[(new BigNumber(3)).toNumber()] = p1;
        _nw263[(new BigNumber(4)).toNumber()] = p1;
        _nw263[(new BigNumber(5)).toNumber()] = p1;
        _nw263[(new BigNumber(6)).toNumber()] = p1;
        _nw263[(new BigNumber(7)).toNumber()] = p1;
        _nw263[(new BigNumber(8)).toNumber()] = p1;
        _nw263[(new BigNumber(9)).toNumber()] = p1;
        _nw263[(new BigNumber(10)).toNumber()] = new BigNumber((function () {
          let _coll36 = new _dafny.Map();
          for (const _compr_36 of ((_1413_v18).dtor_cf30).Elements) {
            let _1424_v13 = _compr_36;
            if (_dafny.Seq.contains((_1413_v18).dtor_cf30, _1424_v13)) {
              _coll36.push([(_1424_v13).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(p1))).cardinality())),p1]);
            }
          }
          return _coll36;
        }()).length);
        _nw263[(new BigNumber(11)).toNumber()] = p1;
        _nw263[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw263[(new BigNumber(13)).toNumber()] = p1;
        _nw263[(new BigNumber(14)).toNumber()] = new BigNumber((_1414_v19).length);
        _nw263[(new BigNumber(15)).toNumber()] = new BigNumber((_1415_v20).length);
        _nw263[(new BigNumber(16)).toNumber()] = p1;
        _nw263[(new BigNumber(17)).toNumber()] = p1;
        _nw263[(new BigNumber(18)).toNumber()] = new BigNumber(218);
        _nw263[(new BigNumber(19)).toNumber()] = new BigNumber((_1417_v23).length);
        _nw263[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Seq.of(p2, !(false))).length);
        _nw263[(new BigNumber(21)).toNumber()] = p1;
        _1423_v24 = _nw263;
        let _1425_v25;
        _1425_v25 = _dafny.MultiSet.fromElements(_1423_v24);
        let _1426_v26;
        _1426_v26 = _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(p1, new BigNumber((_1425_v25).cardinality())), p1, p1, ((p0)[_module.__default.safeIndex(p1, new BigNumber((p0).length))]).minus(p1));
        _1426_v26 = ((p2) ? (_1426_v26) : (_1426_v26));
        let _1427_v27;
        _1427_v27 = _dafny.Set.fromElements(p1);
        (globalState).f24 = _1427_v27;
        (globalState).f7 = (p2) === (p2);
        _1414_v19 = _dafny.Seq.of(p2, p2);
      } else {
        let _1428_v28;
        let _nw264 = Array((new BigNumber(5)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1428_v28 = _nw264;
        let _1429_v29;
        let _nw265 = Array((new BigNumber(6)).toNumber()).fill(false);
        _1429_v29 = _nw265;
        let _index274 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_1428_v28).length));
        (_1428_v28)[_index274] = _dafny.MultiSet.fromElements(_1429_v29);
        let _1430_v30;
        _1430_v30 = _dafny.MultiSet.fromElements(_1429_v29, _1429_v29);
        let _index275 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_1428_v28).length));
        (_1428_v28)[_index275] = _1430_v30;
        (globalState).f19 = new _dafny.CodePoint('b'.codePointAt(0));
        _1429_v29 = _1429_v29;
        let _1431_v31;
        _1431_v31 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1432_v32;
        _1432_v32 = _dafny.Seq.UnicodeFromString("kxbccg");
        let _1433_v33;
        _1433_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1431_v31,_1432_v32);
        (globalState).f10 = _module.__default.safeModuloInt(new BigNumber((_1433_v33).length), p1);
        let _1434_v34;
        let _nw266 = new _module.C1();
        _nw266.__ctor();
        _1434_v34 = _nw266;
      }
      if (p2) {
        let _1435_v35;
        _1435_v35 = _dafny.Seq.UnicodeFromString("g");
        let _1436_v36;
        let _nw267 = new _module.C2();
        _nw267.__ctor(_dafny.Seq.Concat(_1435_v35, _dafny.Seq.UnicodeFromString("enwdan")));
        _1436_v36 = _nw267;
        (globalState).f10 = p1;
        let _1437_v37;
        _1437_v37 = _module.D0.create_DC2(_module.__default.fm0(p1, globalState), false);
        let _pat_let_tv40 = p1;
        if ((function (_pat_let42_0) {
          return function (_1438_dt__update__tmp_h0) {
            return function (_pat_let43_0) {
              return function (_1439_dt__update_hcf5_h0) {
                return _module.D0.create_DC2(_1439_dt__update_hcf5_h0, (_1438_dt__update__tmp_h0).dtor_cf6);
              }(_pat_let43_0);
            }(_pat_let_tv40);
          }(_pat_let42_0);
        }(_1437_v37)).dtor_cf6) {
          (globalState).f13 = p2;
          let _1440_v40;
          _1440_v40 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _1441_v41;
          _1441_v41 = _dafny.Seq.of((((_1440_v40).contains(p2)) ? ((_1440_v40).get(p2)) : (p2)));
          let _1442_v42;
          _1442_v42 = _dafny.Set.fromElements(_1441_v41);
          let _rhs266 = ((!(p2)) ? ((_this).fm2(globalState)) : (new BigNumber((function () {
            let _coll37 = new _dafny.Map();
            for (const _compr_37 of (function () {
              let _coll38 = new _dafny.Map();
              for (const _compr_38 of (_1442_v42).Elements) {
                let _1443_v39 = _compr_38;
                if ((_1442_v42).contains(_1443_v39)) {
                  _coll38.push([_1443_v39,p2]);
                }
              }
              return _coll38;
            }()).Keys.Elements) {
              let _1444_v38 = _compr_37;
              if ((function () {
                let _coll39 = new _dafny.Map();
                for (const _compr_39 of (_1442_v42).Elements) {
                  let _1443_v39 = _compr_39;
                  if ((_1442_v42).contains(_1443_v39)) {
                    _coll39.push([_1443_v39,p2]);
                  }
                }
                return _coll39;
              }()).contains(_1444_v38)) {
                _coll37.push([_1444_v38,new BigNumber((_1405_v9).cardinality())]);
              }
            }
            return _coll37;
          }()).length)));
          let _rhs267 = _1405_v9;
          let _lhs239 = globalState;
          _lhs239.f10 = _rhs266;
          _1405_v9 = _rhs267;
          let _1445_v43;
          _1445_v43 = new _dafny.CodePoint('c'.codePointAt(0));
          let _1446_v44;
          _1446_v44 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
          let _1447_v45;
          _1447_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1445_v43,new BigNumber((_1446_v44).length));
          let _1448_v46;
          _1448_v46 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber((_1447_v45).length), new BigNumber((p0).length)), p1));
          (_1436_v36).m0(new BigNumber(((((_1448_v46).contains(p1)) ? ((_1448_v46).get(p1)) : (p0))).length), p2, p1, !(p2), globalState);
          (globalState).f13 = ((_dafny.ZERO).minus(p1)).isLessThan(p1);
          let _1449_v47;
          _1449_v47 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
          _1449_v47 = (_1449_v47).update((p1).isLessThanOrEqualTo(p1), p1);
        } else {
          (globalState).f13 = p2;
          (globalState).f10 = (p0)[_module.__default.safeIndex(p1, new BigNumber((p0).length))];
          let _1450_v48;
          let _nw268 = Array((new BigNumber(6)).toNumber()).fill(false);
          _1450_v48 = _nw268;
          (globalState).f10 = new BigNumber((_dafny.MultiSet.fromElements(_1450_v48, _1450_v48)).cardinality());
          let _1451_v49;
          _1451_v49 = new _dafny.CodePoint('h'.codePointAt(0));
          (globalState).f19 = _1451_v49;
          let _1452_v50;
          _1452_v50 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1436_v36.f33);
          let _1453_v51;
          _1453_v51 = _dafny.Seq.of(_1436_v36.f33, _dafny.Seq.update((((_1452_v50).contains(p2)) ? ((_1452_v50).get(p2)) : (_1436_v36.f33)), _module.__default.safeIndex(new BigNumber(137), new BigNumber(((((_1452_v50).contains(p2)) ? ((_1452_v50).get(p2)) : (_1436_v36.f33))).length)), _1451_v49), _dafny.Seq.Concat(_1435_v35, _1435_v35));
          _1453_v51 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1453_v51, _1453_v51), _1453_v51);
        }
        (_1436_v36).m0(new BigNumber((_1435_v35).length), p2, (_dafny.ZERO).minus(_module.__default.safeModuloInt(p1, p1)), true, globalState);
        (globalState).f10 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((p1).minus(p1), p1));
      } else {
        (globalState).f13 = p2;
        (globalState).f10 = (new BigNumber((_1405_v9).cardinality())).plus((new BigNumber(146)).plus(p1));
        let _1454_v52;
        _1454_v52 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).fm6(globalState), p1),(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,p0)).length)).multipliedBy(p1));
        _1454_v52 = (_1454_v52).update(p1, new BigNumber(841));
        (globalState).f10 = new BigNumber(622);
        let _1455_v53;
        _1455_v53 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_dafny.Seq.of(p2)).length));
        let _1456_v54;
        _1456_v54 = _dafny.MultiSet.fromElements(new BigNumber(844));
        let _1457_v55;
        _1457_v55 = _dafny.Seq.of(_1456_v54);
        let _1458_v56;
        _1458_v56 = _dafny.Seq.UnicodeFromString("gyebnjvkd");
        let _1459_v57;
        _1459_v57 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
        let _1460_v58;
        let _nw269 = Array((new BigNumber(26)).toNumber());
        _nw269[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p1),p2)).length);
        _nw269[(_dafny.ONE).toNumber()] = p1;
        _nw269[(new BigNumber(2)).toNumber()] = p1;
        _nw269[(new BigNumber(3)).toNumber()] = p1;
        _nw269[(new BigNumber(4)).toNumber()] = p1;
        _nw269[(new BigNumber(5)).toNumber()] = p1;
        _nw269[(new BigNumber(6)).toNumber()] = new BigNumber((_1455_v53).length);
        _nw269[(new BigNumber(7)).toNumber()] = p1;
        _nw269[(new BigNumber(8)).toNumber()] = new BigNumber((_1457_v55).length);
        _nw269[(new BigNumber(9)).toNumber()] = p1;
        _nw269[(new BigNumber(10)).toNumber()] = p1;
        _nw269[(new BigNumber(11)).toNumber()] = p1;
        _nw269[(new BigNumber(12)).toNumber()] = p1;
        _nw269[(new BigNumber(13)).toNumber()] = p1;
        _nw269[(new BigNumber(14)).toNumber()] = (((_1456_v54).contains(p1)) ? ((_1456_v54).get(p1)) : (p1));
        _nw269[(new BigNumber(15)).toNumber()] = new BigNumber(((_1456_v54).update(p1, _module.__default.abs(p1))).cardinality());
        _nw269[(new BigNumber(16)).toNumber()] = new BigNumber((_1458_v56).length);
        _nw269[(new BigNumber(17)).toNumber()] = p1;
        _nw269[(new BigNumber(18)).toNumber()] = p1;
        _nw269[(new BigNumber(19)).toNumber()] = new BigNumber((_1458_v56).length);
        _nw269[(new BigNumber(20)).toNumber()] = p1;
        _nw269[(new BigNumber(21)).toNumber()] = new BigNumber((_1459_v57).length);
        _nw269[(new BigNumber(22)).toNumber()] = p1;
        _nw269[(new BigNumber(23)).toNumber()] = p1;
        _nw269[(new BigNumber(24)).toNumber()] = p1;
        _nw269[(new BigNumber(25)).toNumber()] = p1;
        _1460_v58 = _nw269;
        let _1461_v59;
        _1461_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1460_v58,p2);
        (globalState).f10 = (new BigNumber(447)).plus(new BigNumber((_1461_v59).length));
      }
      let _1462_v60;
      _1462_v60 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_dafny.ZERO).minus(p1), new BigNumber(-889)),(_dafny.ZERO).minus(((p2) ? (p1) : (p1))));
      _1462_v60 = (_1462_v60).update(p1, new BigNumber(565));
      let _1463_v61;
      _1463_v61 = _dafny.Seq.of(p2);
      let _1464_v62;
      _1464_v62 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
      let _1465_v63;
      _1465_v63 = _dafny.MultiSet.fromElements(p1);
      let _1466_v64;
      _1466_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1465_v63,p0);
      let _1467_v65;
      _1467_v65 = _module.D1.create_DC5(p1, p2, (((_1464_v62).contains(false)) ? ((_1464_v62).get(false)) : (!(p2))), _1466_v64);
      let _1468_v66;
      let _nw270 = Array((new BigNumber(17)).toNumber()).fill(false);
      _1468_v66 = _nw270;
      let _1469_v67;
      _1469_v67 = _module.D0.create_DC1(new BigNumber((_dafny.Seq.update(_1463_v61, _module.__default.safeIndex(p1, new BigNumber((_1463_v61).length)), p2)).length), p1, (((_1467_v65).dtor_cf10) ? ((_dafny.ZERO).minus(p1)) : (new BigNumber((_1463_v61).length))), _1468_v66);
      r0 = _1469_v67;
      r1 = (_1465_v63).IsSubsetOf(((true) ? (_1465_v63) : (_1465_v63)));
      return [r0, r1];
    }
    m2(globalState) {
      let _this = this;
      let r0 = [];
      let _1470_v0;
      let _nw271 = Array((new BigNumber(8)).toNumber()).fill(false);
      _1470_v0 = _nw271;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1470_v0).length))) {
        let _1471_i0 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1471_i0)) && ((_1471_i0).isLessThan(new BigNumber((_1470_v0).length))))) {
          (_1470_v0)[(_1471_i0)] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gkgylxbum"), _dafny.Seq.UnicodeFromString("o")), _dafny.Seq.UnicodeFromString("txvt"));
        }
      }
      let _1472_v1;
      _1472_v1 = new BigNumber(109);
      let _1473_v2;
      _1473_v2 = false;
      let _1474_v3;
      _1474_v3 = _module.D0.create_DC2(_1472_v1, _1473_v2);
      let _1475_v4;
      _1475_v4 = _dafny.Seq.of(_1474_v3, _1474_v3, _1474_v3);
      let _1476_v6;
      _1476_v6 = _dafny.MultiSet.fromElements(function () {
        let _coll40 = new _dafny.Map();
        for (const _compr_40 of _dafny.IntegerRange(new BigNumber(-15), new BigNumber(-440))) {
          let _1477_v5 = _compr_40;
          if (((new BigNumber(-15)).isLessThanOrEqualTo(_1477_v5)) && ((_1477_v5).isLessThan(new BigNumber(-440)))) {
            _coll40.push([(_1477_v5).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("lem")).length),_1472_v1)).length)),new BigNumber(518)]);
          }
        }
        return _coll40;
      }());
      let _1478_v8;
      _1478_v8 = _module.D0.create_DC3((_1475_v4)[_module.__default.safeIndex(new BigNumber((function () {
  let _coll41 = new _dafny.Set();
  for (const _compr_41 of (_1476_v6).Elements) {
    let _1479_v7 = _compr_41;
    if ((_1476_v6).contains(_1479_v7)) {
      _coll41.add(_1479_v7);
    }
  }
  return _coll41;
}()).length), new BigNumber((_1475_v4).length))]);
      let _1480_v9;
      _1480_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1472_v1,_1478_v8);
      let _1481_v10;
      _1481_v10 = _dafny.Seq.UnicodeFromString("w");
      let _1482_v11;
      _1482_v11 = _module.D3.create_DC11(_1472_v1);
      let _1483_v12;
      _1483_v12 = _dafny.MultiSet.fromElements(_1472_v1);
      let _1484_v13;
      _1484_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1483_v12,_dafny.Seq.Create(_module.__default.abs(new BigNumber(808)), function (_1485_i1) {
        return new BigNumber(-553);
      }));
      let _1486_v14;
      _1486_v14 = _module.D1.create_DC5(new BigNumber((_dafny.Seq.UnicodeFromString("csvrecknv")).length), _1473_v2, _1473_v2, _1484_v13);
      let _1487_v15;
      _1487_v15 = _module.D1.create_DC5(_1472_v1, false, _1473_v2, (_1486_v14).dtor_cf12);
      let _1488_v16;
      let _init34 = ((_1489_v10) => function (_1490_i2) {
        return _module.__default.safeDivisionInt(_1490_i2, new BigNumber((_1489_v10).length));
      })(_1481_v10);
      let _nw272 = Array((new BigNumber(3)).toNumber());
      for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw272.length); _i0_34++) {
        _nw272[_i0_34] = _init34(new BigNumber(_i0_34));
      }
      _1488_v16 = _nw272;
      let _1491_v17;
      _1491_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1487_v15,_1488_v16);
      let _1492_v18;
      let _nw273 = Array((new BigNumber(12)).toNumber());
      _nw273[(_dafny.ZERO).toNumber()] = new BigNumber(-908);
      _nw273[(_dafny.ONE).toNumber()] = _1472_v1;
      _nw273[(new BigNumber(2)).toNumber()] = new BigNumber((_1480_v9).length);
      _nw273[(new BigNumber(3)).toNumber()] = (new BigNumber((_1481_v10).length)).plus((_dafny.ZERO).minus(_1472_v1));
      _nw273[(new BigNumber(4)).toNumber()] = _1472_v1;
      _nw273[(new BigNumber(5)).toNumber()] = (_1482_v11).dtor_cf22;
      _nw273[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_1472_v1);
      _nw273[(new BigNumber(7)).toNumber()] = _1472_v1;
      _nw273[(new BigNumber(8)).toNumber()] = _1472_v1;
      _nw273[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_1491_v17).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm25(_1473_v2, _1473_v2, _1472_v1, _1473_v2, globalState),_1488_v16))).length));
      _nw273[(new BigNumber(10)).toNumber()] = _1472_v1;
      _nw273[(new BigNumber(11)).toNumber()] = _1472_v1;
      _1492_v18 = _nw273;
      let _index276 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1492_v18).length));
      (_1492_v18)[_index276] = _1472_v1;
      let _index277 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1492_v18).length));
      (_1492_v18)[_index277] = _1472_v1;
      let _1493_v19;
      _1493_v19 = _dafny.Map.Empty.slice().updateUnsafe(((_1492_v18)[_module.__default.safeIndex(new BigNumber(476), new BigNumber((_1492_v18).length))]).isLessThanOrEqualTo(_1472_v1),_dafny.Set.fromElements(_1488_v16, _1488_v16, _1492_v18));
      let _1494_v20;
      _1494_v20 = _dafny.Set.fromElements(_1492_v18);
      _1493_v19 = (_1493_v19).update(_1473_v2, _1494_v20);
      let _1495_v21;
      _1495_v21 = _module.D2.create_DC8(_1473_v2, _1473_v2);
      let _pat_let_tv41 = _1472_v1;
      let _pat_let_tv42 = _1472_v1;
      let _pat_let_tv43 = _1473_v2;
      _1472_v1 = function (_source21) {
        if (_source21.is_DC8) {
          let _1496___mcc_h0 = (_source21).cf18;
          let _1497___mcc_h1 = (_source21).cf19;
          let _1498_cf19 = _1497___mcc_h1;
          let _1499_cf18 = _1496___mcc_h0;
          return _pat_let_tv41;
        } else if (_source21.is_DC7) {
          let _1500___mcc_h2 = (_source21).cf17;
          let _1501_cf17 = _1500___mcc_h2;
          return _pat_let_tv42;
        } else {
          let _1502___mcc_h3 = (_source21).cf20;
          let _1503_cf20 = _1502___mcc_h3;
          return new BigNumber(866);
        }
      }(function (_pat_let44_0) {
        return function (_1504_dt__update__tmp_h0) {
          return function (_pat_let45_0) {
            return function (_1505_dt__update_hcf18_h0) {
              return _module.D2.create_DC8(_1505_dt__update_hcf18_h0, (_1504_dt__update__tmp_h0).dtor_cf19);
            }(_pat_let45_0);
          }(_pat_let_tv43);
        }(_pat_let44_0);
      }(_1495_v21));
      let _index278 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1492_v18).length));
      (_1492_v18)[_index278] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_1492_v18)[_module.__default.safeIndex(new BigNumber(476), new BigNumber((_1492_v18).length))]));
      let _1506_v22;
      let _1507_v23;
      let _1508_v24;
      let _out9;
      let _out10;
      let _out11;
      let _outcollector2 = (_this).m3(((_1473_v2) ? (_1473_v2) : (_1473_v2)), _1470_v0, globalState);
      _out9 = _outcollector2[0];
      _out10 = _outcollector2[1];
      _out11 = _outcollector2[2];
      _1506_v22 = _out9;
      _1507_v23 = _out10;
      _1508_v24 = _out11;
      r0 = _1488_v16;
      return r0;
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = _module.D1.Default();
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let _1509_v0;
      _1509_v0 = _dafny.Seq.UnicodeFromString("ri");
      let _1510_v1;
      _1510_v1 = new BigNumber(212);
      let _1511_v2;
      _1511_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1509_v0,_dafny.Set.fromElements((_dafny.ZERO).minus(_1510_v1), _1510_v1));
      let _1512_v3;
      _1512_v3 = _dafny.Set.fromElements(_1510_v1);
      _1511_v2 = (_1511_v2).update(_1509_v0, _1512_v3);
      (globalState).f13 = p0;
      let _1513_v4;
      _1513_v4 = _dafny.Seq.of(p0);
      let _1514_v5;
      _1514_v5 = _dafny.Seq.of((_1513_v4)[_module.__default.safeIndex(_1510_v1, new BigNumber((_1513_v4).length))]);
      let _1515_v6;
      _1515_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1513_v4).length),_1510_v1);
      _1515_v6 = (_1515_v6).update(_1510_v1, _1510_v1);
      let _1516_v7;
      let _nw274 = new _module.C0();
      _nw274.__ctor(p0);
      _1516_v7 = _nw274;
      _1516_v7 = _1516_v7;
      (globalState).f13 = !_dafny.Seq.contains(_dafny.Seq.Concat(_1509_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(668)), function (_1517_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })), _module.__default.fm11(_1510_v1, globalState));
      let _1518_v8;
      let _nw275 = new _module.C3();
      _nw275.__ctor();
      _1518_v8 = _nw275;
      r0 = _module.__default.fm29(globalState);
      let _1519_v9;
      _1519_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _1520_v10;
      _1520_v10 = _dafny.Set.fromElements(_module.__default.fm30((_1516_v7).f32, globalState), _1519_v9);
      r1 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_module.__default.safeModuloInt(_1510_v1, new BigNumber((_1520_v10).length))), (_this).fm6(globalState));
      r2 = _module.__default.fm35(_1510_v1, p0, (((_1516_v7).f32) ? ((_1516_v7).f32) : (p0)), globalState);
      return [r0, r1, r2];
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this.f30 = _dafny.ZERO;
      this._f29 = [];
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor(f29, f30) {
      let _this = this;
      (_this)._f29 = f29;
      (_this).f30 = f30;
      return;
    }
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_module.D1.create_DC4(function () {
  let _coll42 = new _dafny.Map();
  for (const _compr_42 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false, true, false),_this.f30)).Keys.Elements) {
    let _1521_v0 = _compr_42;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false, true, false),_this.f30)).contains(_1521_v0)) {
      _coll42.push([_1521_v0,true]);
    }
  }
  return _coll42;
}())).dtor_cf8;
    };
    fm2(globalState) {
      let _this = this;
      return _this.f30;
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this.f30).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f30,!(false))).length));
    };
    fm4(p0, p1, globalState) {
      let _this = this;
      return false;
    };
    m0(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (p1) {
        let _1522_v0;
        let _nw276 = Array((new BigNumber(11)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1522_v0 = _nw276;
        let _1523_v1;
        let _init35 = ((_1524_p3) => function (_1525_i0) {
          return _1524_p3;
        })(p3);
        let _nw277 = Array((_dafny.ONE).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw277.length); _i0_35++) {
          _nw277[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1523_v1 = _nw277;
        let _1526_v2;
        _1526_v2 = _dafny.MultiSet.fromElements(_1523_v1);
        let _index279 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_1522_v0).length));
        (_1522_v0)[_index279] = _1526_v2;
        let _index280 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_1522_v0).length));
        let _rhs268 = (_1526_v2).Difference(_1526_v2);
        let _rhs269 = (p1) === (p1);
        let _lhs240 = _1522_v0;
        let _lhs241 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_1522_v0).length));
        let _lhs242 = globalState;
        _lhs240[_lhs241] = _rhs268;
        _lhs242.f22 = _rhs269;
        let _index281 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_1523_v1).length));
        (_1523_v1)[_index281] = p3;
        let _1527_v3;
        _1527_v3 = _dafny.Seq.of(true);
        let _1528_v4;
        _1528_v4 = _dafny.Set.fromElements(p1, p3, (_1527_v3)[_module.__default.safeIndex(_this.f30, new BigNumber((_1527_v3).length))]);
        let _index282 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_1523_v1).length));
        (_1523_v1)[_index282] = ((_dafny.Set.fromElements(false)).Difference(_dafny.Set.fromElements(!(p1), p1, p1, true, p1))).IsProperSubsetOf(_1528_v4);
        let _1529_v5;
        let _init36 = ((_1530_p2, _1531_p0, _1532_p1) => function (_1533_i1) {
          return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(723)), ((_1534_p2) => function (_1535_i2) {
            return _1534_p2;
          })(_1530_p2)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(604)), ((_1536_p0, _1537_p1) => function (_1538_i3) {
            return (_module.D0.create_DC2(_1536_p0, _1537_p1)).dtor_cf5;
          })(_1531_p0, _1532_p1)));
        })(p2, p0, p1);
        let _nw278 = Array((new BigNumber(7)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw278.length); _i0_36++) {
          _nw278[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1529_v5 = _nw278;
        let _1539_v6;
        _1539_v6 = _dafny.MultiSet.fromElements(p0, new BigNumber(259));
        let _rhs270 = _1529_v5;
        let _rhs271 = _1539_v6;
        let _rhs272 = ((_1523_v1)[_module.__default.safeIndex(new BigNumber(417), new BigNumber((_1523_v1).length))]) && (!((_1523_v1)[_module.__default.safeIndex(new BigNumber(417), new BigNumber((_1523_v1).length))]));
        let _lhs243 = globalState;
        _1529_v5 = _rhs270;
        _1539_v6 = _rhs271;
        _lhs243.f7 = _rhs272;
        let _1540_v7;
        _1540_v7 = _dafny.Seq.UnicodeFromString("spjomng");
        let _1541_v8;
        _1541_v8 = _module.D0.create_DC0(_dafny.Seq.Concat(_1540_v7, _1540_v7));
        _1541_v8 = _1541_v8;
        let _1542_v9;
        _1542_v9 = _dafny.Map.Empty.slice().updateUnsafe((_1528_v4).IsProperSubsetOf(_1528_v4),p3);
        let _rhs273 = (((_1542_v9).contains(!(p3))) ? ((_1542_v9).get(!(p3))) : (p1));
        let _rhs274 = _this.f30;
        let _rhs275 = (_dafny.ZERO).minus((_dafny.ZERO).minus((p2).minus(p0)));
        let _rhs276 = _1540_v7;
        let _rhs277 = (_this).fm2(globalState);
        let _lhs244 = globalState;
        let _lhs245 = globalState;
        let _lhs246 = _this;
        let _lhs247 = _this;
        _lhs244.f7 = _rhs273;
        _lhs245.f10 = _rhs274;
        _lhs246.f30 = _rhs275;
        _1540_v7 = _rhs276;
        _lhs247.f30 = _rhs277;
      } else {
        (globalState).f7 = p1;
        let _1543_v10;
        let _nw279 = Array((new BigNumber(8)).toNumber()).fill([]);
        _1543_v10 = _nw279;
        let _1544_v11;
        let _nw280 = new _module.C4();
        _nw280.__ctor(_1543_v10);
        _1544_v11 = _nw280;
        let _1545_v12;
        _1545_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1544_v11,p0);
        let _1546_v13;
        let _nw281 = Array((new BigNumber(3)).toNumber());
        _nw281[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1544_v11,p0);
        _nw281[(_dafny.ONE).toNumber()] = _1545_v12;
        _nw281[(new BigNumber(2)).toNumber()] = _1545_v12;
        _1546_v13 = _nw281;
        let _index283 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1546_v13).length));
        (_1546_v13)[_index283] = _1545_v12;
        let _index284 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1546_v13).length));
        (_1546_v13)[_index284] = (_1545_v12).Merge(_1545_v12);
        let _1547_v14;
        let _init37 = function (_1548_i4) {
          return _dafny.Seq.UnicodeFromString("qj");
        };
        let _nw282 = Array((new BigNumber(23)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw282.length); _i0_37++) {
          _nw282[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1547_v14 = _nw282;
        let _1549_v15;
        _1549_v15 = _dafny.Seq.UnicodeFromString("vkarjbt");
        let _index285 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_1547_v14).length));
        (_1547_v14)[_index285] = _dafny.Seq.Concat(_1549_v15, _1549_v15);
        let _index286 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_1547_v14).length));
        (_1547_v14)[_index286] = _module.__default.fm35(new BigNumber(603), p1, !(_this.f30).isEqualTo(new BigNumber(162)), globalState);
        let _1550_v17;
        _1550_v17 = _dafny.MultiSet.fromElements(function () {
          let _coll43 = new _dafny.Map();
          for (const _compr_43 of _dafny.IntegerRange(new BigNumber(469), new BigNumber(878))) {
            let _1551_v16 = _compr_43;
            if (((new BigNumber(469)).isLessThanOrEqualTo(_1551_v16)) && ((_1551_v16).isLessThan(new BigNumber(878)))) {
              _coll43.push([(_1551_v16).plus(_this.f30),p1]);
            }
          }
          return _coll43;
        }());
        let _1552_v18;
        _1552_v18 = _dafny.Seq.of(_1550_v17);
        let _1553_v19;
        _1553_v19 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f30);
        let _1554_v20;
        let _init38 = ((_1555_p1) => function (_1556_i5) {
          return _1555_p1;
        })(p1);
        let _nw283 = Array((new BigNumber(7)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw283.length); _i0_38++) {
          _nw283[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1554_v20 = _nw283;
        let _1557_v21;
        _1557_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1553_v19,_1554_v20);
        let _1558_v22;
        _1558_v22 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
        let _1559_v23;
        _1559_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1558_v22,p3);
        let _1560_v24;
        _1560_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
        let _1561_v25;
        _1561_v25 = _module.D2.create_DC7(_dafny.Map.Empty.slice().updateUnsafe((((_1553_v19).contains(_this.f30)) ? ((_1553_v19).get(_this.f30)) : (p2)),new BigNumber((_dafny.Seq.UnicodeFromString("h")).length)));
        let _1562_v26;
        _1562_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1561_v25,p0);
        let _1563_v27;
        _1563_v27 = _dafny.Seq.of(_1562_v26, _1562_v26, _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC7(_1553_v19),p0), _1562_v26, _1562_v26);
        let _1564_v28;
        _1564_v28 = _dafny.Map.Empty.slice().updateUnsafe(((p1) ? ((_1552_v18)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1557_v21).length),p0)).length), new BigNumber((_1552_v18).length))]) : ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((((_1558_v22).contains(p1)) ? ((_1558_v22).get(p1)) : (new BigNumber((_1559_v23).length))),p1), _1560_v24)).update(_1560_v24, _module.__default.abs(p2)))),_1563_v27);
        _1564_v28 = (_1564_v28).update(_1550_v17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(11)), ((_1565_v26) => function (_1566_i6) {
          return _1565_v26;
        })(_1562_v26)));
        let _1567_v29;
        _1567_v29 = new _dafny.CodePoint('c'.codePointAt(0));
        let _index287 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_1547_v14).length));
        (_1547_v14)[_index287] = _dafny.Seq.update(_module.__default.fm23(_module.D2.create_DC7(_1553_v19), (_dafny.ZERO).minus((((_1558_v22).contains(p1)) ? ((_1558_v22).get(p1)) : (p2))), (_dafny.ZERO).minus(p0), globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm23(_module.D2.create_DC7(_1553_v19), (_dafny.ZERO).minus((((_1558_v22).contains(p1)) ? ((_1558_v22).get(p1)) : (p2))), (_dafny.ZERO).minus(p0), globalState)).length)), _1567_v29);
      }
      (globalState).f10 = _this.f30;
      let _1568_v30;
      let _nw284 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
      _1568_v30 = _nw284;
      let _1569_v31;
      _1569_v31 = new _dafny.CodePoint('a'.codePointAt(0));
      let _1570_v32;
      _1570_v32 = _dafny.MultiSet.fromElements(_1569_v31, _1569_v31, _1569_v31);
      let _index288 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length));
      (_1568_v30)[_index288] = new BigNumber((_1570_v32).cardinality());
      let _index289 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_1568_v30).length));
      (_1568_v30)[_index289] = _module.__default.safeModuloInt(p2, _this.f30);
      let _1571_v33;
      _1571_v33 = _dafny.Map.Empty.slice().updateUnsafe(!(p3),new BigNumber(554));
      let _1572_v34;
      _1572_v34 = _dafny.MultiSet.fromElements(new BigNumber((_1571_v33).length));
      let _1573_v36;
      _1573_v36 = _module.D3.create_DC11(p0);
      let _1574_v37;
      _1574_v37 = _dafny.Seq.of(new BigNumber((_1572_v34).cardinality()), new BigNumber((function () {
        let _coll44 = new _dafny.Map();
        for (const _compr_44 of (_module.__default.fm14(p1, p2, (_1573_v36).dtor_cf22, globalState)).Elements) {
          let _1575_v35 = _compr_44;
          if (_dafny.Seq.contains(_module.__default.fm14(p1, p2, (_1573_v36).dtor_cf22, globalState), _1575_v35)) {
            _coll44.push([_module.__default.safeModuloInt(_1575_v35, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,_1572_v34)).length)),_1569_v31]);
          }
        }
        return _coll44;
      }()).length), (_this.f30).plus(_module.__default.fm0(p0, globalState)));
      let _index290 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length));
      let _index291 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_1568_v30).length));
      let _rhs278 = (_dafny.ZERO).minus((p0).minus(p2));
      let _rhs279 = p2;
      let _rhs280 = (_1574_v37)[_module.__default.safeIndex((p0).multipliedBy(p0), new BigNumber((_1574_v37).length))];
      let _rhs281 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(p0), new BigNumber(-948)), _1574_v37)).length)).multipliedBy((p2).plus(_this.f30));
      let _rhs282 = new BigNumber((_dafny.Seq.UnicodeFromString("iqxo")).length);
      let _lhs248 = _1568_v30;
      let _lhs249 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length));
      let _lhs250 = _this;
      let _lhs251 = _1568_v30;
      let _lhs252 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_1568_v30).length));
      let _lhs253 = _this;
      let _lhs254 = globalState;
      _lhs248[_lhs249] = _rhs278;
      _lhs250.f30 = _rhs279;
      _lhs251[_lhs252] = _rhs280;
      _lhs253.f30 = _rhs281;
      _lhs254.f10 = _rhs282;
      let _1576_v38;
      _1576_v38 = _module.D0.create_DC2((_1568_v30)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length))], !(p3));
      _1576_v38 = _1576_v38;
      let _1577_v39;
      let _nw285 = new _module.C3();
      _nw285.__ctor();
      _1577_v39 = _nw285;
      let _1578_v40;
      _1578_v40 = _dafny.MultiSet.fromElements(_1577_v39);
      let _1579_v41;
      _1579_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1578_v40).cardinality()),(_1577_v39).fm8(p2, new BigNumber(-44), p0, globalState));
      _1579_v41 = (_1579_v41).update(new BigNumber(744), (_1568_v30)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length))]);
      let _1580_i7;
      _1580_i7 = _dafny.ZERO;
      L6: {
        while (p3) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1580_i7)) {
              break L6;
            }
            _1580_i7 = (_1580_i7).plus(_dafny.ONE);
            if (p3) {
              let _1581_v42;
              let _init39 = ((_1582_p3, _1583_p1) => function (_1584_i8) {
                return ((_1582_p3) ? (_1583_p1) : (_1582_p3));
              })(p3, p1);
              let _nw286 = Array((new BigNumber(11)).toNumber());
              for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw286.length); _i0_39++) {
                _nw286[_i0_39] = _init39(new BigNumber(_i0_39));
              }
              _1581_v42 = _nw286;
              let _index292 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1581_v42).length));
              (_1581_v42)[_index292] = p1;
              let _1585_v43;
              _1585_v43 = _dafny.Seq.UnicodeFromString("fxrsbmfcb");
              let _1586_v44;
              _1586_v44 = _module.D4.create_DC14(p1);
              let _1587_v45;
              _1587_v45 = _module.D4.create_DC15(_1586_v44);
              let _1588_v46;
              _1588_v46 = _module.D7.create_DC20(_dafny.Seq.update(_dafny.Seq.of(_this.f30), _module.__default.safeIndex(new BigNumber((_1579_v41).length), new BigNumber((_dafny.Seq.of(_this.f30)).length)), new BigNumber((_1585_v43).length)), _1581_v42, _1587_v45, p3, _this.f30);
              let _1589_v47;
              _1589_v47 = _module.D0.create_DC0(_1585_v43);
              let _index293 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1581_v42).length));
              let _rhs283 = (_1588_v46).dtor_cf33;
              let _rhs284 = p3;
              let _rhs285 = _dafny.areEqual(_dafny.Seq.Concat(_1585_v43, (_1589_v47).dtor_cf0), _dafny.Seq.Concat(_1585_v43, _dafny.Seq.Create(_module.__default.abs(new BigNumber(936)), ((_1590_v31) => function (_1591_i9) {
                return _1590_v31;
              })(_1569_v31))));
              let _rhs286 = p2;
              let _lhs255 = _1581_v42;
              let _lhs256 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1581_v42).length));
              let _lhs257 = globalState;
              let _lhs258 = globalState;
              let _lhs259 = globalState;
              _lhs255[_lhs256] = _rhs283;
              _lhs257.f7 = _rhs284;
              _lhs258.f13 = _rhs285;
              _lhs259.f10 = _rhs286;
              let _1592_v49;
              _1592_v49 = _dafny.Seq.of((_1581_v42)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1581_v42).length))]);
              let _1593_v50;
              _1593_v50 = _dafny.Seq.of(_1592_v49);
              let _pat_let_tv44 = _1574_v37;
              let _pat_let_tv45 = _1593_v50;
              let _pat_let_tv46 = _1593_v50;
              let _pat_let_tv47 = _1581_v42;
              let _pat_let_tv48 = _1581_v42;
              let _pat_let_tv49 = _1568_v30;
              let _pat_let_tv50 = _1568_v30;
              _1588_v46 = function (_pat_let46_0) {
                return function (_1594_dt__update__tmp_h0) {
                  return function (_pat_let47_0) {
                    return function (_1595_dt__update_hcf30_h0) {
                      return function (_pat_let48_0) {
                        return function (_1597_dt__update_hcf34_h0) {
                          return _module.D7.create_DC20(_1595_dt__update_hcf30_h0, (_1594_dt__update__tmp_h0).dtor_cf31, (_1594_dt__update__tmp_h0).dtor_cf32, (_1594_dt__update__tmp_h0).dtor_cf33, _1597_dt__update_hcf34_h0);
                        }(_pat_let48_0);
                      }(_module.__default.safeDivisionInt(new BigNumber((function () {
                        let _coll45 = new _dafny.Map();
                        for (const _compr_45 of (_dafny.MultiSet.FromArray(_pat_let_tv45)).Elements) {
                          let _1596_v48 = _compr_45;
                          if ((_dafny.MultiSet.FromArray(_pat_let_tv46)).contains(_1596_v48)) {
                            _coll45.push([_1596_v48,(_pat_let_tv48)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_pat_let_tv47).length))]]);
                          }
                        }
                        return _coll45;
                      }()).length), (_pat_let_tv50)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_pat_let_tv49).length))]));
                    }(_pat_let47_0);
                  }(_pat_let_tv44);
                }(_pat_let46_0);
              }(_1588_v46);
              let _1598_v51;
              let _init40 = ((_1599_v49, _1600_v42) => function (_1601_i10) {
                return _dafny.Seq.Concat((_module.D8.create_DC21(_1599_v49)).dtor_cf35, _dafny.Seq.of((_1600_v42)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1600_v42).length))]));
              })(_1592_v49, _1581_v42);
              let _nw287 = Array((new BigNumber(25)).toNumber());
              for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw287.length); _i0_40++) {
                _nw287[_i0_40] = _init40(new BigNumber(_i0_40));
              }
              _1598_v51 = _nw287;
              _1598_v51 = _1598_v51;
              let _1602_v52;
              let _nw288 = Array((new BigNumber(14)).toNumber()).fill(_module.D1.Default());
              _1602_v52 = _nw288;
              let _1603_v53;
              _1603_v53 = _dafny.Map.Empty.slice().updateUnsafe((_1581_v42)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1581_v42).length))],_1602_v52);
              let _1604_v54;
              _1604_v54 = _dafny.Seq.of((((_1603_v53).contains(false)) ? ((_1603_v53).get(false)) : (_1602_v52)), _1602_v52, _1602_v52, _1602_v52);
              let _1605_v55;
              let _nw289 = Array((new BigNumber(6)).toNumber());
              _nw289[(_dafny.ZERO).toNumber()] = _1602_v52;
              _nw289[(_dafny.ONE).toNumber()] = (_1604_v54)[_module.__default.safeIndex((_1574_v37)[_module.__default.safeIndex(_this.f30, new BigNumber((_1574_v37).length))], new BigNumber((_1604_v54).length))];
              _nw289[(new BigNumber(2)).toNumber()] = _1602_v52;
              _nw289[(new BigNumber(3)).toNumber()] = _1602_v52;
              _nw289[(new BigNumber(4)).toNumber()] = _1602_v52;
              _nw289[(new BigNumber(5)).toNumber()] = _1602_v52;
              _1605_v55 = _nw289;
              let _1606_v56;
              let _nw290 = new _module.C4();
              _nw290.__ctor(_1605_v55);
              _1606_v56 = _nw290;
              _1606_v56 = _1606_v56;
              let _1607_v57;
              let _nw291 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _1607_v57 = _nw291;
              let _index294 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_1607_v57).length));
              (_1607_v57)[_index294] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(427)), ((_1608_v43, _1609_v31, _1610_p2) => function (_1611_i11) {
                return (_1608_v43)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1609_v31,_1610_p2)).length), new BigNumber((_1608_v43).length))];
              })(_1585_v43, _1569_v31, p2));
              let _index295 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_1607_v57).length));
              let _rhs287 = _module.__default.fm34(p1, _1569_v31, (((_1581_v42)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1581_v42).length))]) ? ((_1577_v39).fm2(globalState)) : (new BigNumber((_1571_v33).length))), globalState);
              let _rhs288 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-404)), ((_1612_v31) => function (_1613_i12) {
                return _1612_v31;
              })(_1569_v31)), _1585_v43);
              let _lhs260 = _1607_v57;
              let _lhs261 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_1607_v57).length));
              _1592_v49 = _rhs287;
              _lhs260[_lhs261] = _rhs288;
            } else {
              let _1614_v58;
              let _init41 = ((_1615_v36) => function (_1616_i13) {
                return _1615_v36;
              })(_1573_v36);
              let _nw292 = Array((_dafny.ONE).toNumber());
              for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw292.length); _i0_41++) {
                _nw292[_i0_41] = _init41(new BigNumber(_i0_41));
              }
              _1614_v58 = _nw292;
              let _index296 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_1614_v58).length));
              (_1614_v58)[_index296] = _1573_v36;
              let _pat_let_tv51 = _1574_v37;
              let _pat_let_tv52 = _1574_v37;
              let _index297 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_1614_v58).length));
              (_1614_v58)[_index297] = function (_pat_let49_0) {
                return function (_1617_dt__update__tmp_h1) {
                  return function (_pat_let50_0) {
                    return function (_1618_dt__update_hcf22_h0) {
                      return _module.D3.create_DC11(_1618_dt__update_hcf22_h0);
                    }(_pat_let50_0);
                  }((_pat_let_tv51)[_module.__default.safeIndex(new BigNumber(323), new BigNumber((_pat_let_tv52).length))]);
                }(_pat_let49_0);
              }(_module.__default.fm40(new BigNumber((_1572_v34).cardinality()), false, p3, globalState));
              let _index298 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length));
              let _index299 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length));
              let _rhs289 = (_1568_v30)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length))];
              let _rhs290 = p0;
              let _rhs291 = _this.f30;
              let _lhs262 = _1568_v30;
              let _lhs263 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length));
              let _lhs264 = _1568_v30;
              let _lhs265 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length));
              let _lhs266 = globalState;
              _lhs262[_lhs263] = _rhs289;
              _lhs264[_lhs265] = _rhs290;
              _lhs266.f10 = _rhs291;
              (globalState).f7 = p3;
              let _1619_v59;
              let _nw293 = new _module.C1();
              _nw293.__ctor();
              _1619_v59 = _nw293;
              let _1620_v60;
              _1620_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1619_v59,p1);
              let _1621_v61;
              _1621_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1569_v31,p1);
              _1620_v60 = (_1620_v60).update(_1619_v59, (((_1621_v61).contains(_1569_v31)) ? ((_1621_v61).get(_1569_v31)) : (true)));
              let _1622_v62;
              _1622_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(478),_1571_v33);
              let _index300 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length));
              (_1568_v30)[_index300] = (((p1) ? ((_1568_v30)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length))]) : (new BigNumber((_1622_v62).length)))).multipliedBy(p2);
            }
            let _1623_v63;
            _1623_v63 = _dafny.Seq.UnicodeFromString("v");
            let _rhs292 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(46)), ((_1624_v31) => function (_1625_i14) {
              return _1624_v31;
            })(_1569_v31)), _module.__default.safeIndex(_this.f30, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(46)), ((_1626_v31) => function (_1627_i14) {
              return _1626_v31;
            })(_1569_v31))).length)), _1569_v31);
            let _rhs293 = new BigNumber(((_1572_v34).Difference(_dafny.MultiSet.FromArray(_1574_v37))).cardinality());
            let _lhs267 = globalState;
            _1623_v63 = _rhs292;
            _lhs267.f10 = _rhs293;
            if (p3) {
              let _1628_v64;
              let _nw294 = Array((_dafny.ONE).toNumber()).fill(_module.D1.Default());
              _1628_v64 = _nw294;
              let _1629_v65;
              let _nw295 = Array((new BigNumber(10)).toNumber());
              _nw295[(_dafny.ZERO).toNumber()] = _1628_v64;
              _nw295[(_dafny.ONE).toNumber()] = _1628_v64;
              _nw295[(new BigNumber(2)).toNumber()] = _1628_v64;
              _nw295[(new BigNumber(3)).toNumber()] = _1628_v64;
              _nw295[(new BigNumber(4)).toNumber()] = _1628_v64;
              _nw295[(new BigNumber(5)).toNumber()] = _1628_v64;
              _nw295[(new BigNumber(6)).toNumber()] = _1628_v64;
              _nw295[(new BigNumber(7)).toNumber()] = _1628_v64;
              _nw295[(new BigNumber(8)).toNumber()] = _1628_v64;
              _nw295[(new BigNumber(9)).toNumber()] = _1628_v64;
              _1629_v65 = _nw295;
              let _1630_v66;
              let _nw296 = new _module.C4();
              _nw296.__ctor(_1629_v65);
              _1630_v66 = _nw296;
              let _1631_v67;
              _1631_v67 = _dafny.MultiSet.fromElements(_1568_v30, _1568_v30);
              (globalState).f10 = (((_1631_v67).contains(_1568_v30)) ? ((_1631_v67).get(_1568_v30)) : (_this.f30));
              (globalState).f10 = new BigNumber(-41);
              (globalState).f13 = !(((p0).minus(p0)).isLessThan(_this.f30));
              let _1632_v68;
              _1632_v68 = _dafny.Seq.of(p3);
              _1571_v33 = (_1571_v33).update(!(p3) || (p3), (new BigNumber((_1632_v68).length)).minus(_this.f30));
            } else {
              let _1633_v69;
              _1633_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1568_v30,(_this).fm4(_1623_v63, _this.f30, globalState));
              let _1634_v70;
              _1634_v70 = _dafny.Seq.of(_1633_v69, _1633_v69, _1633_v69, _1633_v69, _dafny.Map.Empty.slice().updateUnsafe(_1568_v30,p3));
              let _1635_v71;
              _1635_v71 = _dafny.Seq.of(_1568_v30, _1568_v30);
              _1633_v69 = ((_1634_v70)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(new BigNumber(881))).length), new BigNumber((_1634_v70).length))]).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1635_v71)[_module.__default.safeIndex(_this.f30, new BigNumber((_1635_v71).length))],true));
              let _1636_v72;
              _1636_v72 = _dafny.Seq.of(_module.D4.create_DC14(p1));
              let _1637_v73;
              _1637_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1636_v72,_this.f30);
              let _rhs294 = _1637_v73;
              let _rhs295 = _1623_v63;
              _1637_v73 = _rhs294;
              _1623_v63 = _rhs295;
              let _1638_v74;
              let _nw297 = Array((new BigNumber(22)).toNumber()).fill(_module.D7.Default());
              _1638_v74 = _nw297;
              let _1639_v75;
              _1639_v75 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
              let _1640_v76;
              _1640_v76 = _dafny.Seq.of(_1639_v75);
              let _1641_v77;
              let _nw298 = Array((new BigNumber(26)).toNumber());
              _nw298[(_dafny.ZERO).toNumber()] = p3;
              _nw298[(_dafny.ONE).toNumber()] = true;
              _nw298[(new BigNumber(2)).toNumber()] = !(p1);
              _nw298[(new BigNumber(3)).toNumber()] = p1;
              _nw298[(new BigNumber(4)).toNumber()] = false;
              _nw298[(new BigNumber(5)).toNumber()] = p3;
              _nw298[(new BigNumber(6)).toNumber()] = p1;
              _nw298[(new BigNumber(7)).toNumber()] = !(true);
              _nw298[(new BigNumber(8)).toNumber()] = p3;
              _nw298[(new BigNumber(9)).toNumber()] = p3;
              _nw298[(new BigNumber(10)).toNumber()] = p3;
              _nw298[(new BigNumber(11)).toNumber()] = p3;
              _nw298[(new BigNumber(12)).toNumber()] = (_this).fm4(_1623_v63, new BigNumber((_1640_v76).length), globalState);
              _nw298[(new BigNumber(13)).toNumber()] = p1;
              _nw298[(new BigNumber(14)).toNumber()] = p1;
              _nw298[(new BigNumber(15)).toNumber()] = p1;
              _nw298[(new BigNumber(16)).toNumber()] = p3;
              _nw298[(new BigNumber(17)).toNumber()] = !(false);
              _nw298[(new BigNumber(18)).toNumber()] = false;
              _nw298[(new BigNumber(19)).toNumber()] = p3;
              _nw298[(new BigNumber(20)).toNumber()] = p1;
              _nw298[(new BigNumber(21)).toNumber()] = (_this).fm4(_dafny.Seq.UnicodeFromString("whc"), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1568_v30)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length))],new BigNumber(138))).length), globalState);
              _nw298[(new BigNumber(22)).toNumber()] = p1;
              _nw298[(new BigNumber(23)).toNumber()] = p3;
              _nw298[(new BigNumber(24)).toNumber()] = p3;
              _nw298[(new BigNumber(25)).toNumber()] = p3;
              _1641_v77 = _nw298;
              let _1642_v78;
              _1642_v78 = _module.D4.create_DC14(p3);
              let _1643_v79;
              _1643_v79 = _module.D4.create_DC15(_1642_v78);
              let _1644_v80;
              _1644_v80 = _module.D7.create_DC20(_1574_v37, _1641_v77, _1643_v79, p3, p0);
              let _index301 = _module.__default.safeIndex(new BigNumber(516), new BigNumber((_1638_v74).length));
              (_1638_v74)[_index301] = _1644_v80;
              let _index302 = _module.__default.safeIndex(new BigNumber(516), new BigNumber((_1638_v74).length));
              (_1638_v74)[_index302] = _module.D7.create_DC20(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(312)), function (_1645_i15) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("flxdys")).length);
}), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-914)), ((_1646_p0) => function (_1647_i16) {
  return _1646_p0;
})(p0))), _1641_v77, _1643_v79, p3, p0);
              let _1648_v81;
              _1648_v81 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1623_v63, _module.__default.fm10(_1579_v41, _1569_v31, p3, globalState)),_dafny.Seq.of((_1568_v30)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length))]));
              _1648_v81 = (_1648_v81).update(_1623_v63, _1574_v37);
              let _1649_v82;
              _1649_v82 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(680),_1641_v77);
              _1641_v77 = (((_1649_v82).contains((_1568_v30)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length))])) ? ((_1649_v82).get((_1568_v30)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length))])) : (_1641_v77));
            }
            let _1650_v83;
            let _nw299 = Array((new BigNumber(29)).toNumber()).fill(false);
            _1650_v83 = _nw299;
            let _index303 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1650_v83).length));
            (_1650_v83)[_index303] = ((_1568_v30)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length))]).isLessThanOrEqualTo((_1568_v30)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_1568_v30).length))]);
            let _index304 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1650_v83).length));
            (_1650_v83)[_index304] = (_this).fm4(_1623_v63, new BigNumber((_1623_v63).length), globalState);
          }
        }
      }
      return;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let r1 = false;
      let _1651_v0;
      _1651_v0 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,p1);
      let _hi7 = _module.__default.safeDivisionInt(p1, new BigNumber((_1651_v0).length));
      for (let _1652_i0 = p1; _1652_i0.isLessThan(_hi7); _1652_i0 = _1652_i0.plus(_dafny.ONE)) {
        let _1653_v1;
        _1653_v1 = _dafny.Seq.UnicodeFromString("ctawxugn");
        let _1654_v2;
        _1654_v2 = _dafny.Seq.of(p2, p2);
        let _1655_v3;
        _1655_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm4(_1653_v1, new BigNumber((_dafny.MultiSet.FromArray(_1654_v2)).cardinality()), globalState),new BigNumber(-828));
        let _1656_v4;
        _1656_v4 = _dafny.MultiSet.fromElements(new BigNumber(311));
        let _1657_v5;
        _1657_v5 = _module.D2.create_DC7(_1651_v0);
        let _1658_v6;
        let _nw300 = new _module.C2();
        _nw300.__ctor(_1653_v1);
        _1658_v6 = _nw300;
        let _1659_v7;
        _1659_v7 = _dafny.MultiSet.fromElements(_1658_v6);
        let _1660_v8;
        _1660_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p0).length),_1659_v7);
        let _1661_v9;
        let _nw301 = Array((new BigNumber(27)).toNumber());
        _nw301[(_dafny.ZERO).toNumber()] = !(true);
        _nw301[(_dafny.ONE).toNumber()] = p2;
        _nw301[(new BigNumber(2)).toNumber()] = (_1652_i0).isEqualTo(p1);
        _nw301[(new BigNumber(3)).toNumber()] = !(p2);
        _nw301[(new BigNumber(4)).toNumber()] = p2;
        _nw301[(new BigNumber(5)).toNumber()] = p2;
        _nw301[(new BigNumber(6)).toNumber()] = (_this).fm4(_1653_v1, _this.f30, globalState);
        _nw301[(new BigNumber(7)).toNumber()] = p2;
        _nw301[(new BigNumber(8)).toNumber()] = p2;
        _nw301[(new BigNumber(9)).toNumber()] = p2;
        _nw301[(new BigNumber(10)).toNumber()] = (new BigNumber((_1651_v0).length)).isLessThan(new BigNumber(232));
        _nw301[(new BigNumber(11)).toNumber()] = ((!(p2)) ? (p2) : (p2));
        _nw301[(new BigNumber(12)).toNumber()] = !(p2) || (p2);
        _nw301[(new BigNumber(13)).toNumber()] = p2;
        _nw301[(new BigNumber(14)).toNumber()] = (new BigNumber((_1655_v3).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,p2)).length));
        _nw301[(new BigNumber(15)).toNumber()] = (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_1654_v2).length), new BigNumber(657), p1, new BigNumber(56))).cardinality())).isLessThanOrEqualTo(new BigNumber(801));
        _nw301[(new BigNumber(16)).toNumber()] = (_1656_v4).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1652_i0, (_dafny.ZERO).minus(_1652_i0), new BigNumber((_1653_v1).length), new BigNumber((_module.__default.fm23(_1657_v5, new BigNumber((_dafny.Seq.of(new BigNumber((_1653_v1).length))).length), _this.f30, globalState)).length), new BigNumber((_1654_v2).length)));
        _nw301[(new BigNumber(17)).toNumber()] = p2;
        _nw301[(new BigNumber(18)).toNumber()] = (p1).isLessThan(_this.f30);
        _nw301[(new BigNumber(19)).toNumber()] = p2;
        _nw301[(new BigNumber(20)).toNumber()] = p2;
        _nw301[(new BigNumber(21)).toNumber()] = p2;
        _nw301[(new BigNumber(22)).toNumber()] = p2;
        _nw301[(new BigNumber(23)).toNumber()] = p2;
        _nw301[(new BigNumber(24)).toNumber()] = (_1660_v8).contains(_1652_i0);
        _nw301[(new BigNumber(25)).toNumber()] = (_this.f30).isLessThanOrEqualTo((((_1655_v3).contains(!(!(!(p2))))) ? ((_1655_v3).get(!(!(!(p2))))) : (_1652_i0)));
        _nw301[(new BigNumber(26)).toNumber()] = p2;
        _1661_v9 = _nw301;
        let _rhs296 = (_1652_i0).minus(new BigNumber(376));
        let _rhs297 = _1661_v9;
        let _rhs298 = (_1654_v2)[_module.__default.safeIndex(_1652_i0, new BigNumber((_1654_v2).length))];
        let _lhs268 = globalState;
        let _lhs269 = globalState;
        _lhs268.f10 = _rhs296;
        _1661_v9 = _rhs297;
        _lhs269.f7 = _rhs298;
        let _1662_v10;
        _1662_v10 = new _dafny.CodePoint('b'.codePointAt(0));
        let _rhs299 = _1652_i0;
        let _rhs300 = _1662_v10;
        let _rhs301 = p1;
        let _lhs270 = globalState;
        let _lhs271 = globalState;
        let _lhs272 = _this;
        _lhs270.f10 = _rhs299;
        _lhs271.f19 = _rhs300;
        _lhs272.f30 = _rhs301;
        let _1663_v11;
        let _nw302 = new _module.C3();
        _nw302.__ctor();
        _1663_v11 = _nw302;
        let _1664_v12;
        _1664_v12 = _module.D0.create_DC2(p1, p2);
        let _1665_v13;
        _1665_v13 = _module.D0.create_DC3(_1664_v12);
        let _1666_v14;
        _1666_v14 = _module.D0.create_DC3(_1664_v12);
        let _pat_let_tv53 = _1665_v13;
        let _pat_let_tv54 = _1664_v12;
        let _1667_v15;
        let _nw303 = Array((new BigNumber(20)).toNumber());
        _nw303[(_dafny.ZERO).toNumber()] = _1666_v14;
        _nw303[(_dafny.ONE).toNumber()] = _1666_v14;
        _nw303[(new BigNumber(2)).toNumber()] = _1666_v14;
        _nw303[(new BigNumber(3)).toNumber()] = _1666_v14;
        _nw303[(new BigNumber(4)).toNumber()] = _module.D0.create_DC3(_1665_v13);
        _nw303[(new BigNumber(5)).toNumber()] = function (_pat_let51_0) {
          return function (_1668_dt__update__tmp_h0) {
            return function (_pat_let52_0) {
              return function (_1669_dt__update_hcf7_h0) {
                return _module.D0.create_DC3(_1669_dt__update_hcf7_h0);
              }(_pat_let52_0);
            }(_pat_let_tv53);
          }(_pat_let51_0);
        }(_1666_v14);
        _nw303[(new BigNumber(6)).toNumber()] = function (_pat_let53_0) {
          return function (_1670_dt__update__tmp_h1) {
            return function (_pat_let54_0) {
              return function (_1671_dt__update_hcf7_h1) {
                return _module.D0.create_DC3(_1671_dt__update_hcf7_h1);
              }(_pat_let54_0);
            }(_pat_let_tv54);
          }(_pat_let53_0);
        }(_1666_v14);
        _nw303[(new BigNumber(7)).toNumber()] = _1666_v14;
        _nw303[(new BigNumber(8)).toNumber()] = _1666_v14;
        _nw303[(new BigNumber(9)).toNumber()] = _1666_v14;
        _nw303[(new BigNumber(10)).toNumber()] = _module.D0.create_DC3(_1665_v13);
        _nw303[(new BigNumber(11)).toNumber()] = _1666_v14;
        _nw303[(new BigNumber(12)).toNumber()] = _module.D0.create_DC3(_1664_v12);
        _nw303[(new BigNumber(13)).toNumber()] = _module.D0.create_DC3(_1664_v12);
        _nw303[(new BigNumber(14)).toNumber()] = _1666_v14;
        _nw303[(new BigNumber(15)).toNumber()] = _1666_v14;
        _nw303[(new BigNumber(16)).toNumber()] = _1666_v14;
        _nw303[(new BigNumber(17)).toNumber()] = _1666_v14;
        _nw303[(new BigNumber(18)).toNumber()] = _1666_v14;
        _nw303[(new BigNumber(19)).toNumber()] = _module.D0.create_DC3(_1665_v13);
        _1667_v15 = _nw303;
        let _index305 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1667_v15).length));
        (_1667_v15)[_index305] = _1666_v14;
        let _index306 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1667_v15).length));
        (_1667_v15)[_index306] = _module.D0.create_DC3(_1665_v13);
      }
      let _1672_v17;
      _1672_v17 = _dafny.MultiSet.fromElements(_this.f30);
      let _1673_v18;
      _1673_v18 = _module.D2.create_DC7(function () {
  let _coll46 = new _dafny.Map();
  for (const _compr_46 of (_1672_v17).Elements) {
    let _1674_v16 = _compr_46;
    if ((_1672_v17).contains(_1674_v16)) {
      _coll46.push([(_1674_v16).plus(new BigNumber(645)),p1]);
    }
  }
  return _coll46;
}());
      let _1675_v19;
      _1675_v19 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
      let _1676_v20;
      _1676_v20 = _module.D3.create_DC12((_this).fm4(_module.__default.fm23(_1673_v18, _this.f30, new BigNumber((_1675_v19).length), globalState), p1, globalState));
      _1676_v20 = _1676_v20;
      let _1677_v21;
      let _init42 = ((_1678_p1) => function (_1679_i1) {
        return (_1679_i1).plus(_1678_p1);
      })(p1);
      let _nw304 = Array((new BigNumber(20)).toNumber());
      for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw304.length); _i0_42++) {
        _nw304[_i0_42] = _init42(new BigNumber(_i0_42));
      }
      _1677_v21 = _nw304;
      if ((_dafny.MultiSet.fromElements(_1677_v21, _1677_v21)).IsDisjointFrom(_dafny.MultiSet.fromElements(_1677_v21))) {
        let _1680_v22;
        _1680_v22 = _dafny.MultiSet.fromElements(p2, true, p2, _module.__default.fm12(p2, p2, globalState), p2);
        let _1681_v23;
        _1681_v23 = _module.D1.create_DC4(_dafny.Map.Empty.slice().updateUnsafe(_1680_v22,p2));
        let _source22 = _1681_v23;
        if (_source22.is_DC5) {
          let _1682___mcc_h0 = (_source22).cf9;
          let _1683___mcc_h1 = (_source22).cf10;
          let _1684___mcc_h2 = (_source22).cf11;
          let _1685___mcc_h3 = (_source22).cf12;
          let _1686_cf12 = _1685___mcc_h3;
          let _1687_cf11 = _1684___mcc_h2;
          let _1688_cf10 = _1683___mcc_h1;
          let _1689_cf9 = _1682___mcc_h0;
          _1672_v17 = _1672_v17;
          let _1690_v24;
          _1690_v24 = _module.D8.create_DC22(new BigNumber(98), _1687_cf11, new _dafny.CodePoint('b'.codePointAt(0)));
          let _1691_v25;
          _1691_v25 = new _dafny.CodePoint('o'.codePointAt(0));
          let _1692_v26;
          _1692_v26 = _dafny.Seq.of(_1690_v24, _1690_v24, _1690_v24, _module.D8.create_DC22(_1689_cf9, p2, _1691_v25));
          _1692_v26 = _1692_v26;
          let _1693_v27;
          let _nw305 = Array((new BigNumber(29)).toNumber()).fill([]);
          _1693_v27 = _nw305;
          let _1694_v28;
          let _nw306 = new _module.C4();
          _nw306.__ctor(_1693_v27);
          _1694_v28 = _nw306;
          let _1695_v29;
          _1695_v29 = _dafny.Seq.of(_1687_cf11);
          let _1696_v30;
          _1696_v30 = _dafny.Seq.UnicodeFromString("ra");
          let _1697_v32;
          _1697_v32 = _module.D8.create_DC21(_1695_v29);
          let _1698_v33;
          _1698_v33 = _dafny.Seq.of(_1697_v32);
          let _1699_v34;
          _1699_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1697_v32,new BigNumber((_1696_v30).length));
          let _1700_v35;
          let _nw307 = Array((new BigNumber(5)).toNumber());
          _nw307[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC21(_1695_v29),new BigNumber((_1696_v30).length))).Merge(function () {
            let _coll47 = new _dafny.Map();
            for (const _compr_47 of (_1698_v33).Elements) {
              let _1701_v31 = _compr_47;
              if (_dafny.Seq.contains(_1698_v33, _1701_v31)) {
                _coll47.push([_1701_v31,new BigNumber((_dafny.Set.fromElements(new BigNumber((_1672_v17).cardinality()))).length)]);
              }
            }
            return _coll47;
          }());
          _nw307[(_dafny.ONE).toNumber()] = (_1699_v34).update(_module.D8.create_DC21(_dafny.Seq.update(_1695_v29, _module.__default.safeIndex(_1689_cf9, new BigNumber((_1695_v29).length)), p2)), new BigNumber((_1696_v30).length));
          _nw307[(new BigNumber(2)).toNumber()] = _1699_v34;
          _nw307[(new BigNumber(3)).toNumber()] = _1699_v34;
          _nw307[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1697_v32,_this.f30)).Merge(_1699_v34);
          _1700_v35 = _nw307;
          let _index307 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_1700_v35).length));
          (_1700_v35)[_index307] = _dafny.Map.Empty.slice().updateUnsafe(_1697_v32,_this.f30);
          let _index308 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_1700_v35).length));
          let _rhs302 = _1688_cf10;
          let _rhs303 = _1694_v28;
          let _rhs304 = _1699_v34;
          let _lhs273 = globalState;
          let _lhs274 = _1700_v35;
          let _lhs275 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_1700_v35).length));
          _lhs273.f22 = _rhs302;
          _1694_v28 = _rhs303;
          _lhs274[_lhs275] = _rhs304;
          _1689_cf9 = (_module.__default.safeDivisionInt(p1, _this.f30)).multipliedBy(((_1688_cf10) ? ((_dafny.ZERO).minus(_this.f30)) : (new BigNumber((_1696_v30).length))));
        } else if (_source22.is_DC6) {
          let _1702___mcc_h4 = (_source22).cf13;
          let _1703___mcc_h5 = (_source22).cf14;
          let _1704___mcc_h6 = (_source22).cf15;
          let _1705___mcc_h7 = (_source22).cf16;
          let _1706_cf16 = _1705___mcc_h7;
          let _1707_cf15 = _1704___mcc_h6;
          let _1708_cf14 = _1703___mcc_h5;
          let _1709_cf13 = _1702___mcc_h4;
          let _1710_v36;
          _1710_v36 = _dafny.Seq.UnicodeFromString("rmnqe");
          let _1711_v37;
          _1711_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1706_cf16,_1710_v36);
          let _1712_v38;
          _1712_v38 = _dafny.Set.fromElements(true);
          let _index309 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_1677_v21).length));
          (_1677_v21)[_index309] = (_dafny.ZERO).minus(((_1706_cf16) ? (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_1711_v37).length), _1709_cf13, new BigNumber(287))).cardinality())) : (new BigNumber((_1712_v38).length))));
          let _1713_v39;
          let _nw308 = Array((_dafny.ONE).toNumber());
          _nw308[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
          _1713_v39 = _nw308;
          let _1714_v40;
          _1714_v40 = new _dafny.CodePoint('t'.codePointAt(0));
          let _index310 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_1713_v39).length));
          (_1713_v39)[_index310] = _1714_v40;
          let _1715_v41;
          _1715_v41 = _module.D8.create_DC23(_1706_cf16);
          let _1716_v42;
          let _nw309 = Array((new BigNumber(17)).toNumber());
          _nw309[(_dafny.ZERO).toNumber()] = !(p2);
          _nw309[(_dafny.ONE).toNumber()] = p2;
          _nw309[(new BigNumber(2)).toNumber()] = p2;
          _nw309[(new BigNumber(3)).toNumber()] = p2;
          _nw309[(new BigNumber(4)).toNumber()] = p2;
          _nw309[(new BigNumber(5)).toNumber()] = p2;
          _nw309[(new BigNumber(6)).toNumber()] = p2;
          _nw309[(new BigNumber(7)).toNumber()] = _module.__default.fm12(_1706_cf16, p2, globalState);
          _nw309[(new BigNumber(8)).toNumber()] = p2;
          _nw309[(new BigNumber(9)).toNumber()] = _1706_cf16;
          _nw309[(new BigNumber(10)).toNumber()] = _1706_cf16;
          _nw309[(new BigNumber(11)).toNumber()] = p2;
          _nw309[(new BigNumber(12)).toNumber()] = p2;
          _nw309[(new BigNumber(13)).toNumber()] = p2;
          _nw309[(new BigNumber(14)).toNumber()] = _1706_cf16;
          _nw309[(new BigNumber(15)).toNumber()] = _1706_cf16;
          _nw309[(new BigNumber(16)).toNumber()] = false;
          _1716_v42 = _nw309;
          let _1717_v43;
          _1717_v43 = _module.D4.create_DC13(_1714_v40);
          let _1718_v44;
          _1718_v44 = _module.D4.create_DC15(_1717_v43);
          let _1719_v45;
          _1719_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1672_v17,_dafny.Seq.of(_1709_cf13, _this.f30));
          let _1720_v46;
          _1720_v46 = _module.D1.create_DC5(_1709_cf13, _1706_cf16, _1706_cf16, _1719_v45);
          let _1721_v47;
          _1721_v47 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,_1720_v46);
          let _1722_v49;
          _1722_v49 = _module.D7.create_DC20(p0, _1716_v42, _1718_v44, !(_1706_cf16), (p0)[_module.__default.safeIndex(new BigNumber((function () {
  let _coll48 = new _dafny.Set();
  for (const _compr_48 of (_1721_v47).Keys.Elements) {
    let _1723_v48 = _compr_48;
    if ((_1721_v47).contains(_1723_v48)) {
      _coll48.add(_module.__default.safeModuloInt(_1723_v48, new BigNumber(330)));
    }
  }
  return _coll48;
}()).length), new BigNumber((p0).length))]);
          let _index311 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_1677_v21).length));
          let _index312 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_1713_v39).length));
          let _rhs305 = (_dafny.ZERO).minus((p0)[_module.__default.safeIndex(_1709_cf13, new BigNumber((p0).length))]);
          let _rhs306 = (_1715_v41).dtor_cf39;
          let _rhs307 = (_1722_v49).dtor_cf34;
          let _rhs308 = p1;
          let _rhs309 = (((_module.__default.fm0(_1709_cf13, globalState)).isLessThanOrEqualTo(_1709_cf13)) ? (_1714_v40) : (_1714_v40));
          let _lhs276 = _1677_v21;
          let _lhs277 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_1677_v21).length));
          let _lhs278 = globalState;
          let _lhs279 = globalState;
          let _lhs280 = _1713_v39;
          let _lhs281 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_1713_v39).length));
          _lhs276[_lhs277] = _rhs305;
          _lhs278.f22 = _rhs306;
          _lhs279.f10 = _rhs307;
          _1709_cf13 = _rhs308;
          _lhs280[_lhs281] = _rhs309;
          let _1724_v50;
          _1724_v50 = _dafny.MultiSet.fromElements(_1710_v36, _1710_v36, _dafny.Seq.UnicodeFromString("ybyitirvg"));
          _1724_v50 = _1724_v50;
          let _1725_v51;
          let _nw310 = new _module.C0();
          _nw310.__ctor(p2);
          _1725_v51 = _nw310;
          let _index313 = _module.__default.safeIndex(new BigNumber(631), new BigNumber((_1716_v42).length));
          (_1716_v42)[_index313] = p2;
          let _index314 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_1716_v42).length));
          (_1716_v42)[_index314] = (_1707_cf15).isLessThan((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(919)), function (_1726_i2) {
            return new BigNumber(401);
          })).length)));
          let _1727_v52;
          _1727_v52 = _dafny.Seq.of(_1712_v38, _1712_v38);
          let _index315 = _module.__default.safeIndex(new BigNumber(631), new BigNumber((_1716_v42).length));
          let _index316 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_1716_v42).length));
          let _rhs310 = (_1725_v51).fm13(_1709_cf13, globalState);
          let _rhs311 = !(_1706_cf16);
          let _rhs312 = !(!(!(p2)));
          let _rhs313 = _dafny.Seq.IsPrefixOf(_1727_v52, _1727_v52);
          let _lhs282 = globalState;
          let _lhs283 = _1716_v42;
          let _lhs284 = _module.__default.safeIndex(new BigNumber(631), new BigNumber((_1716_v42).length));
          let _lhs285 = _1716_v42;
          let _lhs286 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_1716_v42).length));
          let _lhs287 = globalState;
          _lhs282.f10 = _rhs310;
          _lhs283[_lhs284] = _rhs311;
          _lhs285[_lhs286] = _rhs312;
          _lhs287.f13 = _rhs313;
        } else {
          let _1728___mcc_h8 = (_source22).cf8;
          let _1729_cf8 = _1728___mcc_h8;
          let _1730_v53;
          _1730_v53 = _dafny.Seq.of(p2, p2, p2);
          let _1731_v54;
          _1731_v54 = _dafny.Seq.of(_1730_v53);
          let _1732_v55;
          _1732_v55 = _module.D1.create_DC6(p1, _1731_v54, p1, p2);
          let _1733_v56;
          _1733_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1732_v55,_1680_v22);
          _1733_v56 = (_1733_v56).update(_1732_v55, _module.__default.fm24(new BigNumber(199), p2, globalState));
          let _1734_v57;
          _1734_v57 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber(728), p1),!(true) || (p2));
          let _1735_v58;
          _1735_v58 = _dafny.Seq.UnicodeFromString("nvdnrsmts");
          _1734_v57 = (_1734_v57).update(p1, !((_this).fm4(_1735_v58, _this.f30, globalState)));
          let _1736_v59;
          _1736_v59 = _module.D8.create_DC21(_1730_v53);
          let _1737_v60;
          _1737_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1736_v59,_1677_v21);
          let _1738_v61;
          _1738_v61 = _1737_v60;
          let _1739_v62;
          _1739_v62 = new _dafny.CodePoint('u'.codePointAt(0));
          let _rhs314 = (_1738_v61);
          let _rhs315 = _1739_v62;
          let _rhs316 = !(p2) || (!(p1).isEqualTo(_this.f30));
          let _lhs288 = globalState;
          let _lhs289 = globalState;
          _1737_v60 = _rhs314;
          _lhs288.f19 = _rhs315;
          _lhs289.f7 = _rhs316;
          let _1740_v63;
          _1740_v63 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,_dafny.Seq.Concat(_1735_v58, _1735_v58));
          _1740_v63 = (_1740_v63).update((_this.f30).minus(p1), _1735_v58);
        }
        (globalState).f10 = (p1).minus(new BigNumber(627));
        let _1741_v64;
        _1741_v64 = _dafny.Seq.UnicodeFromString("vvk");
        let _1742_v65;
        let _nw311 = Array((new BigNumber(27)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1742_v65 = _nw311;
        let _1743_v66;
        _1743_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1741_v64,_1742_v65);
        _1743_v66 = (_dafny.Map.Empty.slice().updateUnsafe(_1741_v64,_1742_v65)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1741_v64,_1742_v65)).Merge(_1743_v66));
        _1742_v65 = _1742_v65;
        let _1744_v67;
        let _nw312 = new _module.C0();
        _nw312.__ctor(!((p1).isLessThanOrEqualTo(_this.f30)));
        _1744_v67 = _nw312;
      } else {
        let _1745_v68;
        let _nw313 = new _module.C3();
        _nw313.__ctor();
        _1745_v68 = _nw313;
        let _1746_v69;
        _1746_v69 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1745_v68);
        let _rhs317 = p2;
        let _rhs318 = _1746_v69;
        let _lhs290 = globalState;
        _lhs290.f13 = _rhs317;
        _1746_v69 = _rhs318;
        (globalState).f7 = p2;
        if (p2) {
          let _1747_v70;
          _1747_v70 = new _dafny.CodePoint('y'.codePointAt(0));
          (globalState).f19 = _1747_v70;
          let _1748_v71;
          _1748_v71 = _dafny.MultiSet.fromElements(p2, !(p2));
          let _1749_v72;
          _1749_v72 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1677_v21);
          let _1750_v73;
          _1750_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1748_v71,(((_1749_v72).contains(!(true))) ? ((_1749_v72).get(!(true))) : (_1677_v21)));
          _1750_v73 = (_1750_v73).update(_1748_v71, _1677_v21);
          (globalState).f10 = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements(p2)).cardinality()), p1))).plus(_this.f30);
          let _1751_v74;
          _1751_v74 = _dafny.Seq.UnicodeFromString("lcy");
          let _1752_v75;
          _1752_v75 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_1751_v74).length));
          (_this).f30 = ((_dafny.ZERO).minus((p1).minus(new BigNumber((_1752_v75).length)))).multipliedBy(_this.f30);
          let _1753_v76;
          let _init43 = function (_1754_i3) {
            return true;
          };
          let _nw314 = Array((new BigNumber(7)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw314.length); _i0_43++) {
            _nw314[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1753_v76 = _nw314;
          let _index317 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_1753_v76).length));
          (_1753_v76)[_index317] = p2;
          let _1755_v77;
          _1755_v77 = _dafny.Seq.of(p2, false);
          let _1756_v78;
          _1756_v78 = _dafny.Set.fromElements(_dafny.Seq.of(p1), p0, p0, p0);
          let _index318 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_1753_v76).length));
          let _rhs319 = p2;
          let _rhs320 = _module.__default.fm22((_dafny.Set.fromElements(p0, p0, _dafny.Seq.of(new BigNumber((p0).length)))).IsSubsetOf(_1756_v78), globalState);
          let _rhs321 = p2;
          let _lhs291 = _1753_v76;
          let _lhs292 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_1753_v76).length));
          let _lhs293 = globalState;
          _lhs291[_lhs292] = _rhs319;
          _1755_v77 = _rhs320;
          _lhs293.f13 = _rhs321;
        } else {
          let _1757_v79;
          _1757_v79 = _dafny.Seq.of(_module.__default.fm12(true, p2, globalState));
          let _1758_v80;
          _1758_v80 = _dafny.Seq.UnicodeFromString("shrji");
          let _1759_v81;
          let _nw315 = new _module.C0();
          _nw315.__ctor((_1757_v79)[_module.__default.safeIndex(new BigNumber((_1758_v80).length), new BigNumber((_1757_v79).length))]);
          _1759_v81 = _nw315;
          let _1760_v82;
          _1760_v82 = _dafny.Seq.of(_1759_v81);
          let _1761_v83;
          let _nw316 = Array((new BigNumber(23)).toNumber());
          _nw316[(_dafny.ZERO).toNumber()] = _1759_v81;
          _nw316[(_dafny.ONE).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(2)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(3)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(4)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(5)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(6)).toNumber()] = (_1760_v82)[_module.__default.safeIndex(_this.f30, new BigNumber((_1760_v82).length))];
          _nw316[(new BigNumber(7)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(8)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(9)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(10)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(11)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(12)).toNumber()] = ((p2) ? (_1759_v81) : (_1759_v81));
          _nw316[(new BigNumber(13)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(14)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(15)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(16)).toNumber()] = (((_1759_v81).f32) ? (_1759_v81) : (_1759_v81));
          _nw316[(new BigNumber(17)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(18)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(19)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(20)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(21)).toNumber()] = _1759_v81;
          _nw316[(new BigNumber(22)).toNumber()] = _1759_v81;
          _1761_v83 = _nw316;
          let _index319 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_1761_v83).length));
          (_1761_v83)[_index319] = _1759_v81;
          let _index320 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_1761_v83).length));
          (_1761_v83)[_index320] = _1759_v81;
          (globalState).f10 = new BigNumber((_1758_v80).length);
          let _1762_v84;
          _1762_v84 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p0).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(345)), function (_1763_i4) {
            return _this.f30;
          }));
          let _1764_v85;
          let _init44 = ((_1765_p2) => function (_1766_i5) {
            return !(_1765_p2);
          })(p2);
          let _nw317 = Array((new BigNumber(22)).toNumber());
          for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw317.length); _i0_44++) {
            _nw317[_i0_44] = _init44(new BigNumber(_i0_44));
          }
          _1764_v85 = _nw317;
          let _1767_v86;
          _1767_v86 = new _dafny.CodePoint('x'.codePointAt(0));
          let _1768_v87;
          _1768_v87 = _module.D4.create_DC13(_1767_v86);
          let _1769_v88;
          _1769_v88 = _module.D4.create_DC15(_1768_v87);
          let _1770_v89;
          _1770_v89 = _module.D7.create_DC20((((_1762_v84).contains(new BigNumber(897))) ? ((_1762_v84).get(new BigNumber(897))) : (p0)), _1764_v85, _1769_v88, p2, (_this).fm2(globalState));
          let _rhs322 = _this.f30;
          let _rhs323 = (((p1).isLessThanOrEqualTo(_this.f30)) ? (_1770_v89) : (_1770_v89));
          let _lhs294 = _this;
          _lhs294.f30 = _rhs322;
          _1770_v89 = _rhs323;
          (globalState).f10 = (_dafny.ZERO).minus(_this.f30);
          let _1771_v90;
          _1771_v90 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
          let _1772_v91;
          _1772_v91 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),(_1759_v81).f32);
          let _1773_v92;
          let _nw318 = Array((new BigNumber(11)).toNumber()).fill(_module.D4.Default());
          _1773_v92 = _nw318;
          let _1774_v93;
          _1774_v93 = _dafny.Seq.of(_1758_v80, _dafny.Seq.UnicodeFromString("e"));
          let _index321 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1773_v92).length));
          (_1773_v92)[_index321] = _module.__default.fm41(new BigNumber((_dafny.Seq.update(_1774_v93, _module.__default.safeIndex(_this.f30, new BigNumber((_1774_v93).length)), _1758_v80)).length), p2, _module.__default.fm24(new BigNumber(93), p2, globalState), _1672_v17, globalState);
          let _1775_v95;
          _1775_v95 = _module.D4.create_DC13(_module.__default.fm11(p1, globalState));
          let _index322 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1773_v92).length));
          let _rhs324 = (function () {
            let _coll49 = new _dafny.Map();
            for (const _compr_49 of (_1651_v0).Keys.Elements) {
              let _1776_v94 = _compr_49;
              if ((_1651_v0).contains(_1776_v94)) {
                _coll49.push([_module.__default.safeDivisionInt(_1776_v94, new BigNumber((_1672_v17).cardinality())),(_1759_v81).f32]);
              }
            }
            return _coll49;
          }()).Merge(_1771_v90);
          let _rhs325 = _1772_v91;
          let _rhs326 = _1775_v95;
          let _lhs295 = _1773_v92;
          let _lhs296 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1773_v92).length));
          _1771_v90 = _rhs324;
          _1772_v91 = _rhs325;
          _lhs295[_lhs296] = _rhs326;
        }
        if (!(p2) || (false)) {
          let _1777_v96;
          _1777_v96 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1778_v97;
          _1778_v97 = _dafny.MultiSet.fromElements(_1777_v96);
          let _rhs327 = ((p2) ? (new _dafny.CodePoint('q'.codePointAt(0))) : (_1777_v96));
          let _rhs328 = p1;
          let _rhs329 = ((p0)[_module.__default.safeIndex(new BigNumber(((_1778_v97).update(_1777_v96, _module.__default.abs(p1))).cardinality()), new BigNumber((p0).length))]).minus(_this.f30);
          let _rhs330 = _this.f30;
          let _lhs297 = globalState;
          let _lhs298 = globalState;
          let _lhs299 = _this;
          let _lhs300 = _this;
          _lhs297.f19 = _rhs327;
          _lhs298.f10 = _rhs328;
          _lhs299.f30 = _rhs329;
          _lhs300.f30 = _rhs330;
          let _1779_v98;
          _1779_v98 = _dafny.Set.fromElements(p2, p2, p2, p2);
          (globalState).f13 = !((_1779_v98).equals(_1779_v98));
          (_this).f30 = (_module.__default.safeDivisionInt(_this.f30, _this.f30)).minus(_this.f30);
          (globalState).f13 = !(!(true)) || (true);
          (globalState).f10 = new BigNumber((_dafny.Seq.Concat(p0, p0)).length);
        } else {
          let _index323 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1677_v21).length));
          (_1677_v21)[_index323] = (_dafny.ZERO).minus(_this.f30);
          let _index324 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1677_v21).length));
          (_1677_v21)[_index324] = (p1).multipliedBy(((_1745_v68).fm2(globalState)).multipliedBy(p1));
          (globalState).f10 = ((!(_1675_v19).equals(_1675_v19)) ? (p1) : (_this.f30));
          let _1780_v99;
          _1780_v99 = _dafny.Seq.of(p2);
          let _1781_v100;
          _1781_v100 = _dafny.MultiSet.fromElements(_1780_v99, _1780_v99);
          let _1782_v101;
          _1782_v101 = _dafny.MultiSet.fromElements(p2);
          (globalState).f10 = ((!(p2)) ? (new BigNumber((_1781_v100).cardinality())) : ((_dafny.ZERO).minus(((_1677_v21)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_1677_v21).length))]).plus(new BigNumber((_dafny.MultiSet.fromElements((_1677_v21)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_1677_v21).length))], new BigNumber((_1782_v101).cardinality()), _this.f30)).cardinality())))));
          let _1783_v102;
          _1783_v102 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1677_v21)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_1677_v21).length))]);
          let _1784_v103;
          _1784_v103 = _dafny.Seq.UnicodeFromString("rbconal");
          _1783_v102 = (_1783_v102).update(p2, new BigNumber((_1784_v103).length));
          (globalState).f13 = p2;
        }
        let _1785_v104;
        _1785_v104 = _dafny.Seq.UnicodeFromString("llhcw");
        let _1786_v105;
        _1786_v105 = _module.D0.create_DC2(((!(p2)) ? ((((_1672_v17).contains(new BigNumber(714))) ? ((_1672_v17).get(new BigNumber(714))) : (_this.f30))) : ((_dafny.ZERO).minus(new BigNumber((_1785_v104).length)))), !(p2));
        _1786_v105 = _module.D0.create_DC2((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_this.f30, (_dafny.ZERO).minus(p1))), p2);
      }
      let _rhs331 = (((_this.f30).isLessThan(new BigNumber((_dafny.Seq.of(p2)).length))) ? (_this.f30) : (_this.f30));
      let _rhs332 = p2;
      let _rhs333 = ((p2) ? ((p0)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((p0).length))]) : (new BigNumber((_dafny.Set.fromElements(p2)).length)));
      let _lhs301 = globalState;
      let _lhs302 = _this;
      _lhs301.f10 = _rhs331;
      r1 = _rhs332;
      _lhs302.f30 = _rhs333;
      if (p2) {
        let _1787_v106;
        let _init45 = ((_1788_p2) => function (_1789_i6) {
          return _1788_p2;
        })(p2);
        let _nw319 = Array((new BigNumber(10)).toNumber());
        for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw319.length); _i0_45++) {
          _nw319[_i0_45] = _init45(new BigNumber(_i0_45));
        }
        _1787_v106 = _nw319;
        let _index325 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1787_v106).length));
        (_1787_v106)[_index325] = p2;
        let _1790_v107;
        let _init46 = function (_1791_i7) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        };
        let _nw320 = Array((_dafny.ONE).toNumber());
        for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw320.length); _i0_46++) {
          _nw320[_i0_46] = _init46(new BigNumber(_i0_46));
        }
        _1790_v107 = _nw320;
        let _1792_v108;
        _1792_v108 = _1790_v107;
        let _1793_v109;
        _1793_v109 = _dafny.Seq.of(_1790_v107);
        let _1794_v110;
        let _nw321 = Array((new BigNumber(17)).toNumber());
        _nw321[(_dafny.ZERO).toNumber()] = _1790_v107;
        _nw321[(_dafny.ONE).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(2)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(3)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(4)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(5)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(6)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(7)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(8)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(9)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(10)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(11)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(12)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(13)).toNumber()] = (_1792_v108);
        _nw321[(new BigNumber(14)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(15)).toNumber()] = _1790_v107;
        _nw321[(new BigNumber(16)).toNumber()] = (_1793_v109)[_module.__default.safeIndex(_this.f30, new BigNumber((_1793_v109).length))];
        _1794_v110 = _nw321;
        let _1795_v111;
        _1795_v111 = _dafny.Seq.of(_1794_v110);
        let _index326 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1787_v106).length));
        let _rhs334 = p2;
        let _rhs335 = p2;
        let _rhs336 = (_1795_v111)[_module.__default.safeIndex(_this.f30, new BigNumber((_1795_v111).length))];
        let _lhs303 = _1787_v106;
        let _lhs304 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1787_v106).length));
        let _lhs305 = globalState;
        let _lhs306 = globalState;
        _lhs303[_lhs304] = _rhs334;
        _lhs305.f22 = _rhs335;
        _lhs306.f6 = _rhs336;
        let _1796_v112;
        _1796_v112 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(_this.f30));
        if ((_1796_v112).IsProperSubsetOf(_1796_v112)) {
          let _1797_v113;
          _1797_v113 = _dafny.Seq.of((_1672_v17).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,p1)).length), p1, p1, _this.f30)), p2);
          _1797_v113 = _module.__default.fm22(!((_1787_v106)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_1787_v106).length))]), globalState);
          let _1798_v114;
          _1798_v114 = _dafny.Seq.UnicodeFromString("bvnrgdjr");
          let _1799_v115;
          _1799_v115 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), function (_1800_i9) {
            return _this.f30;
          }));
          let _1801_v116;
          _1801_v116 = new _dafny.CodePoint('m'.codePointAt(0));
          let _1802_v117;
          _1802_v117 = _module.D0.create_DC0(_1798_v114);
          let _1803_v118;
          let _nw322 = Array((new BigNumber(11)).toNumber());
          _nw322[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_1798_v114, _dafny.Seq.Create(_module.__default.abs(new BigNumber(177)), function (_1804_i8) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          })), _module.__default.safeIndex(new BigNumber((_1799_v115).length), new BigNumber((_dafny.Seq.Concat(_1798_v114, _dafny.Seq.Create(_module.__default.abs(new BigNumber(177)), function (_1805_i8) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          }))).length)), _1801_v116);
          _nw322[(_dafny.ONE).toNumber()] = _1798_v114;
          _nw322[(new BigNumber(2)).toNumber()] = _1798_v114;
          _nw322[(new BigNumber(3)).toNumber()] = _1798_v114;
          _nw322[(new BigNumber(4)).toNumber()] = _1798_v114;
          _nw322[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-829)), ((_1806_v116) => function (_1807_i10) {
            return _1806_v116;
          })(_1801_v116));
          _nw322[(new BigNumber(6)).toNumber()] = _1798_v114;
          _nw322[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), ((_1808_v116) => function (_1809_i11) {
            return _1808_v116;
          })(_1801_v116)), _dafny.Seq.UnicodeFromString("aiqhwkuhp"));
          _nw322[(new BigNumber(8)).toNumber()] = (_1802_v117).dtor_cf0;
          _nw322[(new BigNumber(9)).toNumber()] = _1798_v114;
          _nw322[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1798_v114, _1798_v114);
          _1803_v118 = _nw322;
          let _index327 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1803_v118).length));
          (_1803_v118)[_index327] = _1798_v114;
          let _1810_v119;
          _1810_v119 = _dafny.Map.Empty.slice().updateUnsafe((_1787_v106)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_1787_v106).length))],_dafny.Seq.Create(_module.__default.abs(new BigNumber(-220)), ((_1811_v116) => function (_1812_i12) {
            return _1811_v116;
          })(_1801_v116)));
          let _1813_v120;
          _1813_v120 = _dafny.Map.Empty.slice().updateUnsafe(_1787_v106,p1);
          let _index328 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1803_v118).length));
          (_1803_v118)[_index328] = (((_1810_v119).contains((p2) && ((_1787_v106)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_1787_v106).length))]))) ? ((_1810_v119).get((p2) && ((_1787_v106)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_1787_v106).length))]))) : (_module.__default.fm23(_1673_v18, p1, new BigNumber(((_1813_v120).update(_1787_v106, p1)).length), globalState)));
          let _index329 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1803_v118).length));
          (_1803_v118)[_index329] = _dafny.Seq.update(_dafny.Seq.Concat((_1803_v118)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1803_v118).length))], _1798_v114), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("d")).length), new BigNumber((_dafny.Seq.Concat((_1803_v118)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1803_v118).length))], _1798_v114)).length)), _1801_v116);
          (globalState).f22 = (_1787_v106)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_1787_v106).length))];
          let _1814_v121;
          _1814_v121 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1801_v116,new BigNumber(467)));
          let _1815_v122;
          _1815_v122 = _module.D7.create_DC18(_1814_v121);
          let _pat_let_tv55 = _1814_v121;
          let _1816_v123;
          _1816_v123 = _dafny.Seq.of(_1815_v122, function (_pat_let55_0) {
            return function (_1817_dt__update__tmp_h2) {
              return function (_pat_let56_0) {
                return function (_1818_dt__update_hcf29_h0) {
                  return _module.D7.create_DC18(_1818_dt__update_hcf29_h0);
                }(_pat_let56_0);
              }(_pat_let_tv55);
            }(_pat_let55_0);
          }(_1815_v122));
          (globalState).f13 = (_this).fm4((_1803_v118)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1803_v118).length))], new BigNumber((_1816_v123).length), globalState);
        } else {
          (globalState).f10 = _module.__default.safeDivisionInt(_module.__default.fm0(new BigNumber(-529), globalState), _this.f30);
          let _1819_v124;
          _1819_v124 = _dafny.MultiSet.fromElements(p2);
          let _1820_v125;
          _1820_v125 = _dafny.Seq.of((_1787_v106)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_1787_v106).length))]);
          let _1821_v126;
          _1821_v126 = _dafny.Set.fromElements(_1787_v106, _1787_v106);
          let _rhs337 = _dafny.MultiSet.fromElements(p2, !(_dafny.Set.fromElements(new BigNumber((_1820_v125).length))).equals(_dafny.Set.fromElements(p1)), !((_dafny.Set.fromElements(_1787_v106)).equals(_1821_v126)), (_dafny.MultiSet.fromElements(_this.f30)).IsProperSubsetOf(_dafny.MultiSet.fromElements(p1)));
          let _rhs338 = p2;
          let _rhs339 = p1;
          let _lhs307 = globalState;
          let _lhs308 = _this;
          _1819_v124 = _rhs337;
          _lhs307.f7 = _rhs338;
          _lhs308.f30 = _rhs339;
          let _1822_v127;
          let _nw323 = Array((new BigNumber(19)).toNumber()).fill(_module.D1.Default());
          _1822_v127 = _nw323;
          let _1823_v128;
          _1823_v128 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_1820_v125),p2);
          let _1824_v129;
          _1824_v129 = _module.D1.create_DC4(_1823_v128);
          let _1825_v130;
          _1825_v130 = _dafny.Map.Empty.slice().updateUnsafe(_1824_v129,_1822_v127);
          let _1826_v131;
          _1826_v131 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1822_v127);
          let _1827_v132;
          let _nw324 = Array((new BigNumber(27)).toNumber());
          _nw324[(_dafny.ZERO).toNumber()] = _1822_v127;
          _nw324[(_dafny.ONE).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(2)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(3)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(4)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(5)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(6)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(7)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(8)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(9)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(10)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(11)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(12)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(13)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(14)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(15)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(16)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(17)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(18)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(19)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(20)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(21)).toNumber()] = (((_1825_v130).contains(_1824_v129)) ? ((_1825_v130).get(_1824_v129)) : (_1822_v127));
          _nw324[(new BigNumber(22)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(23)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(24)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(25)).toNumber()] = _1822_v127;
          _nw324[(new BigNumber(26)).toNumber()] = (((_1826_v131).contains(p2)) ? ((_1826_v131).get(p2)) : (_1822_v127));
          _1827_v132 = _nw324;
          let _1828_v133;
          let _nw325 = new _module.C4();
          _nw325.__ctor(_1827_v132);
          _1828_v133 = _nw325;
          let _1829_v134;
          _1829_v134 = _dafny.Map.Empty.slice().updateUnsafe(p2,((p2) ? (_1828_v133) : (_1828_v133)));
          _1829_v134 = (_1829_v134).update((_1787_v106)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_1787_v106).length))], _1828_v133);
          let _1830_v135;
          _1830_v135 = _dafny.Set.fromElements(((_dafny.ZERO).minus(_this.f30)).plus(_this.f30), _this.f30);
          (globalState).f24 = _1830_v135;
          let _1831_v136;
          let _nw326 = new _module.C4();
          _nw326.__ctor((_1828_v133).f31);
          _1831_v136 = _nw326;
        }
        (globalState).f13 = (_module.__default.safeDivisionInt(_this.f30, p1)).isLessThanOrEqualTo((_dafny.ZERO).minus(_this.f30));
        let _index330 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1787_v106).length));
        (_1787_v106)[_index330] = p2;
        (globalState).f13 = p2;
      } else {
        let _1832_v137;
        let _nw327 = Array((new BigNumber(19)).toNumber()).fill([]);
        _1832_v137 = _nw327;
        let _index331 = _module.__default.safeIndex(new BigNumber(77), new BigNumber(((_this).f29).length));
        ((_this).f29)[_index331] = _1677_v21;
        let _index332 = _module.__default.safeIndex(new BigNumber(77), new BigNumber(((_this).f29).length));
        let _rhs340 = _1832_v137;
        let _rhs341 = _this.f30;
        let _rhs342 = _1677_v21;
        let _rhs343 = _this.f30;
        let _lhs309 = globalState;
        let _lhs310 = (_this).f29;
        let _lhs311 = _module.__default.safeIndex(new BigNumber(77), new BigNumber(((_this).f29).length));
        let _lhs312 = globalState;
        _1832_v137 = _rhs340;
        _lhs309.f10 = _rhs341;
        _lhs310[_lhs311] = _rhs342;
        _lhs312.f10 = _rhs343;
        if (p2) {
          let _1833_v138;
          _1833_v138 = _dafny.MultiSet.fromElements(p2);
          (globalState).f10 = (_module.__default.safeDivisionInt(p1, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ekgpmyva")).length)))).plus(new BigNumber(((_1833_v138).Union(_dafny.MultiSet.fromElements(p2))).cardinality()));
          (globalState).f10 = _module.__default.safeModuloInt(p1, _module.__default.safeDivisionInt(_this.f30, _this.f30));
          (globalState).f13 = p2;
          (globalState).f7 = p2;
          (globalState).f22 = false;
        } else {
          (globalState).f13 = p2;
          let _index333 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_1677_v21).length));
          (_1677_v21)[_index333] = p1;
          let _index334 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_1677_v21).length));
          (_1677_v21)[_index334] = _this.f30;
          let _index335 = _module.__default.safeIndex(new BigNumber(77), new BigNumber(((_this).f29).length));
          ((_this).f29)[_index335] = ((_this).f29)[_module.__default.safeIndex(new BigNumber(77), new BigNumber(((_this).f29).length))];
          let _1834_v139;
          _1834_v139 = _dafny.Seq.UnicodeFromString("lhnsen");
          _1834_v139 = _1834_v139;
          let _1835_v140;
          let _nw328 = Array((new BigNumber(10)).toNumber());
          _nw328[(_dafny.ZERO).toNumber()] = (p2) === (p2);
          _nw328[(_dafny.ONE).toNumber()] = p2;
          _nw328[(new BigNumber(2)).toNumber()] = p2;
          _nw328[(new BigNumber(3)).toNumber()] = p2;
          _nw328[(new BigNumber(4)).toNumber()] = p2;
          _nw328[(new BigNumber(5)).toNumber()] = ((p2) ? (p2) : (p2));
          _nw328[(new BigNumber(6)).toNumber()] = p2;
          _nw328[(new BigNumber(7)).toNumber()] = p2;
          _nw328[(new BigNumber(8)).toNumber()] = true;
          _nw328[(new BigNumber(9)).toNumber()] = p2;
          _1835_v140 = _nw328;
          let _index336 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1835_v140).length));
          (_1835_v140)[_index336] = !(p2);
          let _index337 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1835_v140).length));
          (_1835_v140)[_index337] = p2;
        }
        let _1836_v141;
        _1836_v141 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_this.f30),p2);
        let _1837_v142;
        _1837_v142 = _dafny.Map.Empty.slice().updateUnsafe(_1672_v17,p0);
        let _1838_v143;
        _1838_v143 = _module.D1.create_DC5(new BigNumber((_1836_v141).length), true, p2, _1837_v142);
        let _1839_v144;
        _1839_v144 = _dafny.Seq.UnicodeFromString("ipg");
        let _1840_v145;
        let _nw329 = Array((new BigNumber(27)).toNumber());
        _nw329[(_dafny.ZERO).toNumber()] = p2;
        _nw329[(_dafny.ONE).toNumber()] = p2;
        _nw329[(new BigNumber(2)).toNumber()] = p2;
        _nw329[(new BigNumber(3)).toNumber()] = p2;
        _nw329[(new BigNumber(4)).toNumber()] = p2;
        _nw329[(new BigNumber(5)).toNumber()] = p2;
        _nw329[(new BigNumber(6)).toNumber()] = !(p2);
        _nw329[(new BigNumber(7)).toNumber()] = p2;
        _nw329[(new BigNumber(8)).toNumber()] = (_1838_v143).dtor_cf11;
        _nw329[(new BigNumber(9)).toNumber()] = p2;
        _nw329[(new BigNumber(10)).toNumber()] = p2;
        _nw329[(new BigNumber(11)).toNumber()] = p2;
        _nw329[(new BigNumber(12)).toNumber()] = p2;
        _nw329[(new BigNumber(13)).toNumber()] = p2;
        _nw329[(new BigNumber(14)).toNumber()] = p2;
        _nw329[(new BigNumber(15)).toNumber()] = !(true);
        _nw329[(new BigNumber(16)).toNumber()] = p2;
        _nw329[(new BigNumber(17)).toNumber()] = true;
        _nw329[(new BigNumber(18)).toNumber()] = p2;
        _nw329[(new BigNumber(19)).toNumber()] = p2;
        _nw329[(new BigNumber(20)).toNumber()] = p2;
        _nw329[(new BigNumber(21)).toNumber()] = (_this).fm4(_1839_v144, p1, globalState);
        _nw329[(new BigNumber(22)).toNumber()] = p2;
        _nw329[(new BigNumber(23)).toNumber()] = true;
        _nw329[(new BigNumber(24)).toNumber()] = p2;
        _nw329[(new BigNumber(25)).toNumber()] = p2;
        _nw329[(new BigNumber(26)).toNumber()] = p2;
        _1840_v145 = _nw329;
        let _1841_v146;
        _1841_v146 = _module.D4.create_DC14(p2);
        let _1842_v147;
        _1842_v147 = _module.D4.create_DC15(_1841_v146);
        let _1843_v148;
        _1843_v148 = _module.D7.create_DC20(p0, _1840_v145, _module.D4.create_DC15(_1841_v146), !_dafny.Seq.contains(p0, _this.f30), p1);
        let _1844_v149;
        _1844_v149 = _dafny.Seq.of(p0, _dafny.Seq.of(_this.f30, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(937)), function (_1845_i13) {
          return (_dafny.ZERO).minus(_this.f30);
        })).length)), _this.f30), p0, p0, p0);
        let _1846_v150;
        _1846_v150 = _module.D4.create_DC15(_1841_v146);
        _1843_v148 = ((p2) ? (_module.D7.create_DC20(_dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber(((_1844_v149)[_module.__default.safeIndex(_this.f30, new BigNumber((_1844_v149).length))]).length), new BigNumber((p0).length)), p1), _1840_v145, _1846_v150, p2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f30,(_dafny.ZERO).minus(p1))).length))) : (_1843_v148));
        let _1847_v151;
        let _nw330 = new _module.C1();
        _nw330.__ctor();
        _1847_v151 = _nw330;
        if ((new BigNumber((_dafny.Seq.update(p0, _module.__default.safeIndex(_this.f30, new BigNumber((p0).length)), new BigNumber((p0).length))).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f30,(_this).fm4(_1839_v144, _this.f30, globalState))).length)))) {
          let _1848_v152;
          _1848_v152 = _module.D0.create_DC2((_dafny.ZERO).minus(p1), p2);
          let _1849_v153;
          _1849_v153 = _dafny.Seq.of(_1848_v152);
          let _1850_v154;
          _1850_v154 = _dafny.Seq.of(_1849_v153, _1849_v153, _dafny.Seq.update(_1849_v153, _module.__default.safeIndex(_this.f30, new BigNumber((_1849_v153).length)), _1848_v152), _dafny.Seq.Create(_module.__default.abs(new BigNumber(367)), ((_1851_v152) => function (_1852_i14) {
            return _1851_v152;
          })(_1848_v152)), _1849_v153);
          let _1853_v155;
          _1853_v155 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.update(_1850_v154, _module.__default.safeIndex(new BigNumber((_1651_v0).length), new BigNumber((_1850_v154).length)), _1849_v153));
          _1850_v154 = ((p2) ? (_1850_v154) : ((((_1853_v155).contains(false)) ? ((_1853_v155).get(false)) : (_1850_v154))));
          let _1854_v156;
          let _nw331 = new _module.C3();
          _nw331.__ctor();
          _1854_v156 = _nw331;
          (globalState).f10 = p1;
          let _1855_v157;
          _1855_v157 = _dafny.Map.Empty.slice().updateUnsafe((_1847_v151).fm21(p1, globalState),(new BigNumber((_1839_v144).length)).plus(_this.f30));
          (globalState).f10 = new BigNumber((_1855_v157).length);
          let _1856_v158;
          let _nw332 = Array((new BigNumber(22)).toNumber()).fill(_module.D1.Default());
          _1856_v158 = _nw332;
          let _1857_v159;
          _1857_v159 = _1856_v158;
          let _1858_v160;
          _1858_v160 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1856_v158);
          let _1859_v161;
          let _nw333 = Array((new BigNumber(12)).toNumber());
          _nw333[(_dafny.ZERO).toNumber()] = _1856_v158;
          _nw333[(_dafny.ONE).toNumber()] = _1856_v158;
          _nw333[(new BigNumber(2)).toNumber()] = _1856_v158;
          _nw333[(new BigNumber(3)).toNumber()] = _1856_v158;
          _nw333[(new BigNumber(4)).toNumber()] = _1856_v158;
          _nw333[(new BigNumber(5)).toNumber()] = _1856_v158;
          _nw333[(new BigNumber(6)).toNumber()] = _1856_v158;
          _nw333[(new BigNumber(7)).toNumber()] = _1856_v158;
          _nw333[(new BigNumber(8)).toNumber()] = (_1857_v159);
          _nw333[(new BigNumber(9)).toNumber()] = (((_1858_v160).contains(p2)) ? ((_1858_v160).get(p2)) : (_1856_v158));
          _nw333[(new BigNumber(10)).toNumber()] = _1856_v158;
          _nw333[(new BigNumber(11)).toNumber()] = _1856_v158;
          _1859_v161 = _nw333;
          let _1860_v162;
          let _nw334 = new _module.C4();
          _nw334.__ctor(_1859_v161);
          _1860_v162 = _nw334;
          _1860_v162 = _1860_v162;
        } else {
          (globalState).f7 = p2;
          let _1861_v163;
          let _nw335 = new _module.C1();
          _nw335.__ctor();
          _1861_v163 = _nw335;
          (globalState).f22 = p2;
          let _1862_v164;
          _1862_v164 = _dafny.Seq.of(p2);
          let _1863_v165;
          _1863_v165 = new _dafny.CodePoint('y'.codePointAt(0));
          _1672_v17 = (_module.__default.fm17(_this.f30, (_1862_v164)[_module.__default.safeIndex(p1, new BigNumber((_1862_v164).length))], _1863_v165, globalState)).Intersect(_1672_v17);
          let _1864_v166;
          let _nw336 = new _module.C0();
          _nw336.__ctor(((true) ? (p2) : (p2)));
          _1864_v166 = _nw336;
        }
      }
      let _1865_i15;
      _1865_i15 = _dafny.ZERO;
      L7: {
        while (p2) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1865_i15)) {
              break L7;
            }
            _1865_i15 = (_1865_i15).plus(_dafny.ONE);
            (globalState).f22 = (((_1675_v19).contains(p2)) ? ((_1675_v19).get(p2)) : (false));
            (globalState).f22 = p2;
            if (p2) {
              (globalState).f10 = ((p0)[_module.__default.safeIndex(p1, new BigNumber((p0).length))]).minus(_this.f30);
              (globalState).f10 = _module.__default.safeModuloInt(p1, (_dafny.ZERO).minus(_this.f30));
              let _rhs344 = _1677_v21;
              let _rhs345 = _module.__default.fm0((p1).plus(p1), globalState);
              let _lhs313 = globalState;
              _1677_v21 = _rhs344;
              _lhs313.f10 = _rhs345;
              let _1866_v167;
              let _nw337 = new _module.C3();
              _nw337.__ctor();
              _1866_v167 = _nw337;
              let _1867_v168;
              _1867_v168 = _dafny.Seq.UnicodeFromString("kdqjxd");
              let _1868_v169;
              let _nw338 = new _module.C2();
              _nw338.__ctor(_1867_v168);
              _1868_v169 = _nw338;
            } else {
              _1677_v21 = _1677_v21;
              let _1869_v170;
              let _nw339 = Array((new BigNumber(10)).toNumber());
              _nw339[(_dafny.ZERO).toNumber()] = _1677_v21;
              _nw339[(_dafny.ONE).toNumber()] = ((p2) ? (_1677_v21) : (_1677_v21));
              _nw339[(new BigNumber(2)).toNumber()] = _1677_v21;
              _nw339[(new BigNumber(3)).toNumber()] = _1677_v21;
              _nw339[(new BigNumber(4)).toNumber()] = _1677_v21;
              _nw339[(new BigNumber(5)).toNumber()] = ((p2) ? (_1677_v21) : (_1677_v21));
              _nw339[(new BigNumber(6)).toNumber()] = _1677_v21;
              _nw339[(new BigNumber(7)).toNumber()] = _1677_v21;
              _nw339[(new BigNumber(8)).toNumber()] = _1677_v21;
              _nw339[(new BigNumber(9)).toNumber()] = _1677_v21;
              _1869_v170 = _nw339;
              _1869_v170 = _1869_v170;
              let _1870_v171;
              let _nw340 = Array((_dafny.ONE).toNumber()).fill(_dafny.MultiSet.Empty);
              _1870_v171 = _nw340;
              _1870_v171 = _1870_v171;
              let _1871_v172;
              _1871_v172 = new _dafny.CodePoint('o'.codePointAt(0));
              let _1872_v173;
              _1872_v173 = _module.D8.create_DC22(p1, !(p2), _1871_v172);
              let _1873_v174;
              _1873_v174 = _dafny.Set.fromElements(_1872_v173);
              let _1874_v175;
              _1874_v175 = _dafny.Seq.UnicodeFromString("lfkqt");
              _1873_v174 = ((((_this).fm4(_1874_v175, p1, globalState)) ? (_1873_v174) : (_1873_v174))).Union(_1873_v174);
              _1651_v0 = (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_this.f30),p1)).Merge(_1651_v0);
            }
            let _1875_v176;
            let _nw341 = new _module.C1();
            _nw341.__ctor();
            _1875_v176 = _nw341;
            let _1876_v177;
            _1876_v177 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1875_v176);
            if (((_1876_v177).Merge(_1876_v177)).contains(p2)) {
              let _1877_v178;
              _1877_v178 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.UnicodeFromString("q"));
              let _1878_v179;
              _1878_v179 = _dafny.Seq.UnicodeFromString("djlw");
              _1877_v178 = (_1877_v178).update(p2, _dafny.Seq.Concat(_1878_v179, _1878_v179));
              r1 = !(p2);
              let _1879_v180;
              let _nw342 = Array((new BigNumber(7)).toNumber());
              _nw342[(_dafny.ZERO).toNumber()] = _1878_v179;
              _nw342[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1878_v179, _1878_v179);
              _nw342[(new BigNumber(2)).toNumber()] = _1878_v179;
              _nw342[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("vxtcqx");
              _nw342[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1878_v179, _1878_v179);
              _nw342[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1878_v179, _1878_v179);
              _nw342[(new BigNumber(6)).toNumber()] = _module.__default.fm23(_module.D2.create_DC7(_1651_v0), _this.f30, (_dafny.ZERO).minus(_this.f30), globalState);
              _1879_v180 = _nw342;
              let _index338 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_1879_v180).length));
              (_1879_v180)[_index338] = _1878_v179;
              let _index339 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_1879_v180).length));
              (_1879_v180)[_index339] = _1878_v179;
              let _index340 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1677_v21).length));
              (_1677_v21)[_index340] = _module.__default.safeModuloInt(_this.f30, new BigNumber(-430));
              let _1880_v181;
              _1880_v181 = _dafny.Map.Empty.slice().updateUnsafe(_1875_v176,_1677_v21);
              let _1881_v182;
              let _nw343 = Array((new BigNumber(7)).toNumber()).fill(_dafny.MultiSet.Empty);
              _1881_v182 = _nw343;
              let _index341 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_1881_v182).length));
              (_1881_v182)[_index341] = _dafny.MultiSet.FromArray(_dafny.Seq.of(_1878_v179, _1878_v179, (_1879_v180)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_1879_v180).length))]));
              let _1882_v183;
              let _nw344 = new _module.C0();
              _nw344.__ctor(true);
              _1882_v183 = _nw344;
              let _1883_v184;
              _1883_v184 = _dafny.Seq.of((_1882_v183).f32, (_1882_v183).f32);
              let _1884_v185;
              _1884_v185 = _module.D8.create_DC21(_1883_v184);
              let _1885_v186;
              _1885_v186 = _module.D8.create_DC24(_1884_v185);
              let _1886_v187;
              _1886_v187 = _dafny.Map.Empty.slice().updateUnsafe(_1882_v183,_1885_v186);
              let _1887_v188;
              _1887_v188 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
              let _1888_v189;
              _1888_v189 = _dafny.Map.Empty.slice().updateUnsafe((_1886_v187).Merge(_1886_v187),_dafny.MultiSet.fromElements(_module.__default.fm23(_1673_v18, (_dafny.ZERO).minus(new BigNumber((_1887_v188).length)), p1, globalState), (_1879_v180)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_1879_v180).length))], (_1879_v180)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_1879_v180).length))]));
              let _index342 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1677_v21).length));
              let _index343 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_1881_v182).length));
              let _rhs346 = p1;
              let _rhs347 = (_dafny.Set.fromElements(p1)).IsDisjointFrom(_module.__default.fm42((_dafny.ZERO).minus(_this.f30), _dafny.Set.fromElements(p2, !(p2)), p1, globalState));
              let _rhs348 = new BigNumber((((p2) ? ((((_1877_v178).contains(p2)) ? ((_1877_v178).get(p2)) : (_module.__default.fm23(_1673_v18, p1, _this.f30, globalState)))) : ((_1879_v180)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_1879_v180).length))]))).length);
              let _rhs349 = _1880_v181;
              let _rhs350 = (((_1888_v189).contains((_dafny.Map.Empty.slice().updateUnsafe(_1882_v183,_1885_v186)).Merge(_1886_v187))) ? ((_1888_v189).get((_dafny.Map.Empty.slice().updateUnsafe(_1882_v183,_1885_v186)).Merge(_1886_v187))) : ((_dafny.MultiSet.fromElements((_1879_v180)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_1879_v180).length))], (_1879_v180)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_1879_v180).length))])).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(883)), ((_1889_v180) => function (_1890_i16) {
                return (_1889_v180)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_1889_v180).length))];
              })(_1879_v180))))));
              let _lhs314 = globalState;
              let _lhs315 = globalState;
              let _lhs316 = _1677_v21;
              let _lhs317 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1677_v21).length));
              let _lhs318 = _1881_v182;
              let _lhs319 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_1881_v182).length));
              _lhs314.f10 = _rhs346;
              _lhs315.f7 = _rhs347;
              _lhs316[_lhs317] = _rhs348;
              _1880_v181 = _rhs349;
              _lhs318[_lhs319] = _rhs350;
              (globalState).f22 = p2;
            } else {
              (globalState).f13 = p2;
              let _1891_v190;
              _1891_v190 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,p2);
              let _1892_v191;
              _1892_v191 = _module.D2.create_DC8(p2, !((((_1891_v190).contains(p1)) ? ((_1891_v190).get(p1)) : (true))));
              _1892_v191 = _1892_v191;
              let _1893_v192;
              _1893_v192 = _dafny.Seq.of(_module.__default.fm12(p2, p2, globalState), p2);
              let _1894_v193;
              _1894_v193 = _dafny.Seq.of(_1651_v0, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_1893_v192)).cardinality()),_this.f30));
              let _1895_v194;
              _1895_v194 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1894_v193);
              let _1896_v195;
              _1896_v195 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f30);
              _1895_v194 = (_1895_v194).update((((_1896_v195).contains(!(false))) ? ((_1896_v195).get(!(false))) : (new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("nqoaxevgx")).length))).length))), _1894_v193);
              (_1875_v176).m0(new BigNumber(436), _dafny.Seq.contains(p0, p1), _this.f30, p2, globalState);
              (globalState).f10 = new BigNumber(-389);
            }
          }
        }
      }
      let _1897_v196;
      _1897_v196 = _dafny.Seq.of(!(!(p2)), p2, p2);
      let _1898_v197;
      _1898_v197 = _dafny.Seq.UnicodeFromString("htgli");
      let _1899_v198;
      _1899_v198 = _module.D0.create_DC2(new BigNumber((p0).length), p2);
      let _1900_v199;
      let _nw345 = Array((new BigNumber(5)).toNumber()).fill(false);
      _1900_v199 = _nw345;
      let _1901_v200;
      _1901_v200 = _module.D0.create_DC1(new BigNumber((_dafny.Seq.Concat(_1897_v196, _1897_v196)).length), p1, (_this).fm3(new BigNumber((_1898_v197).length), p2, (_this).fm2(globalState), _1899_v198, globalState), _1900_v199);
      r0 = _1901_v200;
      r1 = p2;
      return [r0, r1];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
