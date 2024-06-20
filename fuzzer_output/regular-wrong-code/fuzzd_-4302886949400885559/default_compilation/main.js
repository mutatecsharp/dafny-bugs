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
      return new _dafny.CodePoint('y'.codePointAt(0));
    };
    static fm1(p0, p1, globalState) {
      return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(false), false)).length), new BigNumber((_dafny.MultiSet.fromElements(!(!(!((false)))))).cardinality()), (_dafny.ZERO).minus(new BigNumber(-825))),(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("wirhrrwqa")).length), (_dafny.ZERO).minus(new BigNumber(-305)), new BigNumber(846), new BigNumber((_dafny.Seq.of(!(true))).length))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality()))))).length);
    };
    static fm2(p0, p1, globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(-177)), _dafny.Seq.of(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(true)).length)))).length), new BigNumber((_dafny.Seq.of(false, true, true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(288),!(false))).length)))).IsSubsetOf(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(423)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("kxsyhfc")).length)), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))));
    };
    static fm3(p0, p1, globalState) {
      return _dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)));
    };
    static fm5(p0, p1, p2, p3, globalState) {
      return false;
    };
    static fm6(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(800))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(186))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(331)), function (_0_i0) {
          return new BigNumber(206);
        })).Elements) {
          let _1_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(331)), function (_0_i0) {
            return new BigNumber(206);
          }), _1_v0)) {
            _coll0.add(_module.__default.safeModuloInt(_1_v0, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false)).length))));
          }
        }
        return _coll0;
      }()).length))));
    };
    static fm7(p0, globalState) {
      return (_module.D5.create_DC5(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))))).dtor_cf5;
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(934))).length));
    };
    static fm9(p0, globalState) {
      return _dafny.Seq.of(((true) ? (true) : (!(!(false)))), true, false, !(!(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("d")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).length), new BigNumber(274), new BigNumber(377))).cardinality())).isEqualTo(new BigNumber(-936))), !(false) || (false));
    };
    static m0(p0, p1, p2, globalState) {
      let r0 = _dafny.Map.Empty;
      let _2_v0;
      _2_v0 = new BigNumber(896);
      (globalState).f12 = (_2_v0).isLessThanOrEqualTo(_2_v0);
      let _3_v1;
      _3_v1 = _dafny.Set.fromElements(p2);
      let _4_v2;
      _4_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2_v0);
      let _5_v3;
      _5_v3 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(p1, _4_v2, globalState),p2);
      if (((_3_v1).Intersect(_dafny.Set.fromElements((((_5_v3).contains((_dafny.ZERO).minus(_2_v0))) ? ((_5_v3).get((_dafny.ZERO).minus(_2_v0))) : (p2))))).IsSubsetOf(_3_v1)) {
        let _6_v4;
        _6_v4 = _dafny.Seq.of(p0);
        _6_v4 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false), _6_v4), _6_v4);
        let _7_v5;
        let _nw0 = Array((new BigNumber(12)).toNumber());
        _nw0[(_dafny.ZERO).toNumber()] = p2;
        _nw0[(_dafny.ONE).toNumber()] = p2;
        _nw0[(new BigNumber(2)).toNumber()] = p1;
        _nw0[(new BigNumber(3)).toNumber()] = p2;
        _nw0[(new BigNumber(4)).toNumber()] = p2;
        _nw0[(new BigNumber(5)).toNumber()] = !(!(p0));
        _nw0[(new BigNumber(6)).toNumber()] = !(p2);
        _nw0[(new BigNumber(7)).toNumber()] = p1;
        _nw0[(new BigNumber(8)).toNumber()] = p1;
        _nw0[(new BigNumber(9)).toNumber()] = p0;
        _nw0[(new BigNumber(10)).toNumber()] = !(p0);
        _nw0[(new BigNumber(11)).toNumber()] = p1;
        _7_v5 = _nw0;
        let _8_v6;
        let _nw1 = new _module.C0();
        _nw1.__ctor();
        _8_v6 = _nw1;
        let _9_v7;
        _9_v7 = _dafny.Map.Empty.slice().updateUnsafe(_7_v5,_8_v6);
        _9_v7 = (_9_v7).update(_7_v5, _8_v6);
        _6_v4 = _dafny.Seq.Concat(_6_v4, _6_v4);
        let _10_v8;
        _10_v8 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        let _source0 = _10_v8;
        let _11___mcc_h0 = _source0;
        let _12_cf1 = _11___mcc_h0;
        let _13_v9;
        let _nw2 = new _module.C0();
        _nw2.__ctor();
        _13_v9 = _nw2;
        let _14_v10;
        _14_v10 = _dafny.Map.Empty.slice().updateUnsafe(_8_v6,p2);
        let _15_v11;
        _15_v11 = _dafny.MultiSet.fromElements((_4_v2).Merge(_module.__default.fm6((((_14_v10).contains(_8_v6)) ? ((_14_v10).get(_8_v6)) : (p1)), p0, globalState)));
        (globalState).f5 = (((_15_v11).contains(_4_v2)) ? ((_15_v11).get(_4_v2)) : (_2_v0));
        let _16_v12;
        let _init0 = function (_17_i0) {
          return _dafny.Seq.UnicodeFromString("mm");
        };
        let _nw3 = Array((new BigNumber(27)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw3.length); _i0_0++) {
          _nw3[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _16_v12 = _nw3;
        let _18_v13;
        _18_v13 = _dafny.Seq.UnicodeFromString("igqyhvkw");
        let _19_v14;
        _19_v14 = _18_v13;
        let _index0 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_16_v12).length));
        (_16_v12)[_index0] = _19_v14;
        let _index1 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_16_v12).length));
        (_16_v12)[_index1] = _dafny.Seq.Concat(_18_v13, _18_v13);
        let _20_v15;
        _20_v15 = new _dafny.CodePoint('j'.codePointAt(0));
        let _21_v16;
        _21_v16 = _2_v0;
        let _22_v17;
        _22_v17 = _dafny.MultiSet.fromElements(_13_v9);
        let _23_v18;
        let _nw4 = Array((new BigNumber(10)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = new BigNumber((_22_v17).cardinality());
        _nw4[(_dafny.ONE).toNumber()] = _2_v0;
        _nw4[(new BigNumber(2)).toNumber()] = new BigNumber((_4_v2).length);
        _nw4[(new BigNumber(3)).toNumber()] = new BigNumber((_18_v13).length);
        _nw4[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.of(_2_v0)).length);
        _nw4[(new BigNumber(5)).toNumber()] = new BigNumber((_6_v4).length);
        _nw4[(new BigNumber(6)).toNumber()] = _2_v0;
        _nw4[(new BigNumber(7)).toNumber()] = _2_v0;
        _nw4[(new BigNumber(8)).toNumber()] = new BigNumber((_18_v13).length);
        _nw4[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_18_v13, _module.__default.safeIndex(_2_v0, new BigNumber((_18_v13).length)), _20_v15), _module.__default.safeIndex(_2_v0, new BigNumber((_dafny.Seq.update(_18_v13, _module.__default.safeIndex(_2_v0, new BigNumber((_18_v13).length)), _20_v15)).length)), _20_v15)).length);
        _23_v18 = _nw4;
        let _24_v19;
        _24_v19 = _dafny.MultiSet.fromElements((_21_v16), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(p0),_23_v18)).length), (_2_v0).multipliedBy(_module.__default.fm1(p0, _4_v2, globalState)));
        let _rhs0 = (_2_v0).isLessThan(_module.__default.fm1(p0, _4_v2, globalState));
        let _rhs1 = _20_v15;
        let _rhs2 = _24_v19;
        let _rhs3 = p2;
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        _lhs0.f7 = _rhs0;
        _20_v15 = _rhs1;
        _24_v19 = _rhs2;
        _lhs1.f7 = _rhs3;
        if (!((_2_v0).isLessThan(new BigNumber((_6_v4).length)))) {
          let _index2 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_7_v5).length));
          (_7_v5)[_index2] = p1;
          let _index3 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_7_v5).length));
          (_7_v5)[_index3] = true;
          (globalState).f7 = (_2_v0).isLessThanOrEqualTo(_2_v0);
          (globalState).f12 = p0;
          let _25_v20;
          let _init1 = ((_26_v0) => function (_27_i1) {
            return _module.__default.safeDivisionInt(_27_i1, _26_v0);
          })(_2_v0);
          let _nw5 = Array((new BigNumber(24)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw5.length); _i0_1++) {
            _nw5[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _25_v20 = _nw5;
          let _28_v21;
          _28_v21 = _dafny.Seq.of(_25_v20, _25_v20);
          _28_v21 = _28_v21;
          (globalState).f3 = _2_v0;
        } else {
          (globalState).f12 = (_2_v0).isLessThan((_2_v0).minus((_dafny.ZERO).minus(_2_v0)));
          let _29_v22;
          _29_v22 = _dafny.Seq.UnicodeFromString("aov");
          let _30_v23;
          _30_v23 = _dafny.Map.Empty.slice().updateUnsafe(_8_v6,_29_v22);
          let _31_v24;
          _31_v24 = new BigNumber(-98);
          let _32_v25;
          _32_v25 = _dafny.Map.Empty.slice().updateUnsafe((_31_v24),_29_v22);
          let _33_v26;
          _33_v26 = _dafny.MultiSet.fromElements(_8_v6);
          let _34_v27;
          let _nw6 = Array((new BigNumber(21)).toNumber());
          _nw6[(_dafny.ZERO).toNumber()] = _module.__default.fm1(p2, _4_v2, globalState);
          _nw6[(_dafny.ONE).toNumber()] = new BigNumber(((_30_v23).Merge(_30_v23)).length);
          _nw6[(new BigNumber(2)).toNumber()] = _2_v0;
          _nw6[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_2_v0);
          _nw6[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm1(p1, _4_v2, globalState));
          _nw6[(new BigNumber(5)).toNumber()] = _2_v0;
          _nw6[(new BigNumber(6)).toNumber()] = (_31_v24);
          _nw6[(new BigNumber(7)).toNumber()] = _2_v0;
          _nw6[(new BigNumber(8)).toNumber()] = _2_v0;
          _nw6[(new BigNumber(9)).toNumber()] = new BigNumber((_32_v25).length);
          _nw6[(new BigNumber(10)).toNumber()] = _2_v0;
          _nw6[(new BigNumber(11)).toNumber()] = new BigNumber(((_4_v2).Merge(_4_v2)).length);
          _nw6[(new BigNumber(12)).toNumber()] = _2_v0;
          _nw6[(new BigNumber(13)).toNumber()] = _2_v0;
          _nw6[(new BigNumber(14)).toNumber()] = new BigNumber(597);
          _nw6[(new BigNumber(15)).toNumber()] = _2_v0;
          _nw6[(new BigNumber(16)).toNumber()] = new BigNumber(353);
          _nw6[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_2_v0, (_dafny.ZERO).minus(new BigNumber((_3_v1).length))));
          _nw6[(new BigNumber(18)).toNumber()] = (_2_v0).minus(new BigNumber((_33_v26).cardinality()));
          _nw6[(new BigNumber(19)).toNumber()] = _2_v0;
          _nw6[(new BigNumber(20)).toNumber()] = _module.__default.fm1(p0, _4_v2, globalState);
          _34_v27 = _nw6;
          let _index4 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_34_v27).length));
          (_34_v27)[_index4] = _2_v0;
          let _index5 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_34_v27).length));
          (_34_v27)[_index5] = (_2_v0).multipliedBy(_2_v0);
          let _35_v28;
          _35_v28 = new _dafny.CodePoint('x'.codePointAt(0));
          _35_v28 = _35_v28;
          let _36_v29;
          let _nw7 = new _module.C0();
          _nw7.__ctor();
          _36_v29 = _nw7;
          let _index6 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_7_v5).length));
          (_7_v5)[_index6] = p2;
          let _index7 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_7_v5).length));
          (_7_v5)[_index7] = p0;
        }
      } else {
        let _37_v30;
        let _nw8 = new _module.C0();
        _nw8.__ctor();
        _37_v30 = _nw8;
        let _38_v31;
        _38_v31 = _dafny.Seq.UnicodeFromString("corwclx");
        let _39_v32;
        _39_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2_v0,_module.__default.fm8(true, _2_v0, _38_v31, _2_v0, globalState));
        let _40_v33;
        _40_v33 = _dafny.MultiSet.fromElements(_2_v0);
        _39_v32 = (_39_v32).update(_2_v0, _40_v33);
        let _41_v34;
        _41_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(296),_5_v3);
        _41_v34 = (_41_v34).update(_module.__default.safeDivisionInt(_2_v0, _2_v0), _5_v3);
        (globalState).f12 = p2;
        let _42_v35;
        let _init2 = function (_43_i2) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        };
        let _nw9 = Array((new BigNumber(2)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw9.length); _i0_2++) {
          _nw9[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _42_v35 = _nw9;
        let _44_v36;
        _44_v36 = new _dafny.CodePoint('b'.codePointAt(0));
        let _index8 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_42_v35).length));
        (_42_v35)[_index8] = _44_v36;
        let _index9 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_42_v35).length));
        (_42_v35)[_index9] = _44_v36;
      }
      let _45_i3;
      _45_i3 = _dafny.ZERO;
      L0: {
        while (p0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_45_i3)) {
              break L0;
            }
            _45_i3 = (_45_i3).plus(_dafny.ONE);
            if (p1) {
              (globalState).f3 = _2_v0;
              let _46_v37;
              let _init3 = function (_47_i4) {
                return _dafny.Seq.UnicodeFromString("woh");
              };
              let _nw10 = Array((new BigNumber(5)).toNumber());
              for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw10.length); _i0_3++) {
                _nw10[_i0_3] = _init3(new BigNumber(_i0_3));
              }
              _46_v37 = _nw10;
              let _48_v38;
              _48_v38 = _module.__default.fm3(p0, _2_v0, globalState);
              let _index10 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_46_v37).length));
              (_46_v37)[_index10] = _48_v38;
              let _49_v39;
              _49_v39 = _dafny.Seq.UnicodeFromString("siujnxa");
              let _index11 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_46_v37).length));
              (_46_v37)[_index11] = _49_v39;
              (globalState).f7 = p1;
              let _50_v40;
              let _nw11 = Array((new BigNumber(13)).toNumber());
              _nw11[(_dafny.ZERO).toNumber()] = p1;
              _nw11[(_dafny.ONE).toNumber()] = p1;
              _nw11[(new BigNumber(2)).toNumber()] = p2;
              _nw11[(new BigNumber(3)).toNumber()] = p1;
              _nw11[(new BigNumber(4)).toNumber()] = ((p0) ? (p2) : (p2));
              _nw11[(new BigNumber(5)).toNumber()] = p2;
              _nw11[(new BigNumber(6)).toNumber()] = p1;
              _nw11[(new BigNumber(7)).toNumber()] = p0;
              _nw11[(new BigNumber(8)).toNumber()] = p0;
              _nw11[(new BigNumber(9)).toNumber()] = !(_3_v1).contains(_module.__default.fm2((_dafny.ZERO).minus(_2_v0), p0, globalState));
              _nw11[(new BigNumber(10)).toNumber()] = p2;
              _nw11[(new BigNumber(11)).toNumber()] = p0;
              _nw11[(new BigNumber(12)).toNumber()] = (p0) && (p1);
              _50_v40 = _nw11;
              _50_v40 = _50_v40;
              let _51_v41;
              let _nw12 = new _module.C0();
              _nw12.__ctor();
              _51_v41 = _nw12;
              let _52_v42;
              _52_v42 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
              let _53_v43;
              _53_v43 = _dafny.MultiSet.fromElements(false, true, (((_52_v42).contains(true)) ? ((_52_v42).get(true)) : (p0)));
              let _rhs4 = p0;
              let _rhs5 = _2_v0;
              let _rhs6 = false;
              let _rhs7 = (_53_v43).IsProperSubsetOf(_53_v43);
              let _rhs8 = _51_v41;
              let _lhs2 = globalState;
              let _lhs3 = globalState;
              let _lhs4 = globalState;
              let _lhs5 = globalState;
              _lhs2.f12 = _rhs4;
              _lhs3.f3 = _rhs5;
              _lhs4.f12 = _rhs6;
              _lhs5.f12 = _rhs7;
              _51_v41 = _rhs8;
            } else {
              let _54_v44;
              let _nw13 = new _module.C0();
              _nw13.__ctor();
              _54_v44 = _nw13;
              let _55_v45;
              _55_v45 = _dafny.Seq.UnicodeFromString("sb");
              (globalState).f12 = (new BigNumber((_55_v45).length)).isLessThan(_2_v0);
              let _56_v46;
              _56_v46 = _dafny.MultiSet.fromElements(_2_v0);
              (globalState).f7 = ((_56_v46).Difference(_56_v46)).IsProperSubsetOf(_module.__default.fm8(p0, _2_v0, _dafny.Seq.UnicodeFromString("eqibxy"), _module.__default.fm1(p1, _dafny.Map.Empty.slice().updateUnsafe(false,_2_v0), globalState), globalState));
              let _57_v47;
              let _init4 = ((_58_p1) => function (_59_i5) {
                return _58_p1;
              })(p1);
              let _nw14 = Array((new BigNumber(14)).toNumber());
              for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw14.length); _i0_4++) {
                _nw14[_i0_4] = _init4(new BigNumber(_i0_4));
              }
              _57_v47 = _nw14;
              let _60_v48;
              _60_v48 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
              let _61_v49;
              _61_v49 = _dafny.Seq.of(_60_v48, _60_v48);
              let _62_v50;
              _62_v50 = _dafny.Map.Empty.slice().updateUnsafe(_57_v47,(_61_v49)[_module.__default.safeIndex(_2_v0, new BigNumber((_61_v49).length))]);
              let _63_v51;
              _63_v51 = _dafny.Seq.of(p2);
              let _64_v52;
              _64_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_63_v51).length),_60_v48);
              let _rhs9 = (_dafny.ZERO).minus(_2_v0);
              let _rhs10 = (((_62_v50).contains(_57_v47)) ? ((_62_v50).get(_57_v47)) : ((((_64_v52).contains(new BigNumber(768))) ? ((_64_v52).get(new BigNumber(768))) : (_60_v48))));
              let _lhs6 = globalState;
              _lhs6.f5 = _rhs9;
              r0 = _rhs10;
              let _65_v53;
              let _init5 = ((_66_v45) => function (_67_i6) {
                return _66_v45;
              })(_55_v45);
              let _nw15 = Array((new BigNumber(10)).toNumber());
              for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw15.length); _i0_5++) {
                _nw15[_i0_5] = _init5(new BigNumber(_i0_5));
              }
              _65_v53 = _nw15;
              _65_v53 = _65_v53;
            }
            (globalState).f12 = !(true);
            let _68_v54;
            let _init6 = ((_69_p2) => function (_70_i7) {
              return _69_p2;
            })(p2);
            let _nw16 = Array((new BigNumber(8)).toNumber());
            for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw16.length); _i0_6++) {
              _nw16[_i0_6] = _init6(new BigNumber(_i0_6));
            }
            _68_v54 = _nw16;
            let _init7 = ((_71_p2, _72_p0, _73_v3, _74_p1) => function (_75_i8) {
              return ((_71_p2) ? ((((_73_v3).contains(new BigNumber((_dafny.Seq.of(_72_p0)).length))) ? ((_73_v3).get(new BigNumber((_dafny.Seq.of(_72_p0)).length))) : (_74_p1))) : (_74_p1));
            })(p2, p0, _5_v3, p1);
            let _nw17 = Array((new BigNumber(23)).toNumber());
            for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw17.length); _i0_7++) {
              _nw17[_i0_7] = _init7(new BigNumber(_i0_7));
            }
            _68_v54 = _nw17;
            let _76_v55;
            _76_v55 = _dafny.Map.Empty.slice().updateUnsafe(_68_v54,p0);
            (globalState).f5 = (_2_v0).multipliedBy(new BigNumber(((_76_v55).Merge(_76_v55)).length));
          }
        }
      }
      let _77_v56;
      _77_v56 = new _dafny.CodePoint('w'.codePointAt(0));
      let _78_v57;
      _78_v57 = _dafny.Seq.UnicodeFromString("rwuaf");
      if (!(!(((p0) ? (!_dafny.Seq.contains(_78_v57, _77_v56)) : (p2))))) {
        let _79_v58;
        let _nw18 = Array((new BigNumber(26)).toNumber());
        _79_v58 = _nw18;
        _79_v58 = _79_v58;
        let _80_v59;
        _80_v59 = _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_2_v0))).cardinality()));
        let _81_v60;
        _81_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,_80_v59);
        _81_v60 = (_81_v60).update(p2, _80_v59);
        let _82_v61;
        let _nw19 = new _module.C0();
        _nw19.__ctor();
        _82_v61 = _nw19;
        (globalState).f12 = p0;
        let _83_v62;
        _83_v62 = _dafny.Seq.of(_2_v0, _2_v0, new BigNumber(179), (_dafny.ZERO).minus(_2_v0));
        let _84_v63;
        _84_v63 = _dafny.Map.Empty.slice().updateUnsafe(_83_v62,!(!(!(p0)) || (p2)));
        _84_v63 = (_84_v63).update(_dafny.Seq.Concat(_83_v62, _83_v62), p2);
      } else {
        _78_v57 = _78_v57;
        let _85_v64;
        _85_v64 = _dafny.MultiSet.fromElements(p1);
        let _86_v65;
        _86_v65 = _dafny.Set.fromElements(_2_v0);
        let _87_v66;
        _87_v66 = _dafny.MultiSet.fromElements(new BigNumber((_78_v57).length));
        let _88_v67;
        _88_v67 = _dafny.MultiSet.fromElements(_87_v66, _87_v66);
        let _89_v68;
        _89_v68 = _dafny.Map.Empty.slice().updateUnsafe(_88_v67,_4_v2);
        let _90_v69;
        _90_v69 = _dafny.Seq.of(_2_v0);
        let _91_v70;
        let _nw20 = Array((new BigNumber(19)).toNumber());
        _nw20[(_dafny.ZERO).toNumber()] = _2_v0;
        _nw20[(_dafny.ONE).toNumber()] = _2_v0;
        _nw20[(new BigNumber(2)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(3)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(4)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(5)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(6)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(7)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(8)).toNumber()] = new BigNumber((_85_v64).cardinality());
        _nw20[(new BigNumber(9)).toNumber()] = new BigNumber((_78_v57).length);
        _nw20[(new BigNumber(10)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(11)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_90_v69).length));
        _nw20[(new BigNumber(13)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(14)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(15)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(16)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(17)).toNumber()] = _2_v0;
        _nw20[(new BigNumber(18)).toNumber()] = _2_v0;
        _91_v70 = _nw20;
        let _92_v71;
        _92_v71 = _dafny.Seq.of(_91_v70, _91_v70);
        let _93_v72;
        _93_v72 = _dafny.Seq.of((_92_v71)[_module.__default.safeIndex(_2_v0, new BigNumber((_92_v71).length))]);
        let _94_v73;
        _94_v73 = false;
        let _95_v74;
        _95_v74 = (_94_v73);
        let _96_v75;
        let _nw21 = new _module.C0();
        _nw21.__ctor();
        _96_v75 = _nw21;
        let _97_v76;
        _97_v76 = _dafny.Map.Empty.slice().updateUnsafe(_96_v75,_78_v57);
        let _98_v77;
        let _nw22 = Array((new BigNumber(25)).toNumber());
        _nw22[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.of(p2)).length);
        _nw22[(_dafny.ONE).toNumber()] = new BigNumber((_85_v64).cardinality());
        _nw22[(new BigNumber(2)).toNumber()] = _2_v0;
        _nw22[(new BigNumber(3)).toNumber()] = _2_v0;
        _nw22[(new BigNumber(4)).toNumber()] = new BigNumber(-713);
        _nw22[(new BigNumber(5)).toNumber()] = _module.__default.fm1(p0, _4_v2, globalState);
        _nw22[(new BigNumber(6)).toNumber()] = ((p2) ? (_2_v0) : (new BigNumber(144)));
        _nw22[(new BigNumber(7)).toNumber()] = new BigNumber(((_86_v65).Intersect(_86_v65)).length);
        _nw22[(new BigNumber(8)).toNumber()] = _2_v0;
        _nw22[(new BigNumber(9)).toNumber()] = _2_v0;
        _nw22[(new BigNumber(10)).toNumber()] = _module.__default.fm1(_module.__default.fm2(new BigNumber(661), p0, globalState), (((_89_v68).contains(_88_v67)) ? ((_89_v68).get(_88_v67)) : (_4_v2)), globalState);
        _nw22[(new BigNumber(11)).toNumber()] = _2_v0;
        _nw22[(new BigNumber(12)).toNumber()] = new BigNumber((_93_v72).length);
        _nw22[(new BigNumber(13)).toNumber()] = _2_v0;
        _nw22[(new BigNumber(14)).toNumber()] = (((_94_v73)) ? (_2_v0) : (_2_v0));
        _nw22[(new BigNumber(15)).toNumber()] = _2_v0;
        _nw22[(new BigNumber(16)).toNumber()] = _2_v0;
        _nw22[(new BigNumber(17)).toNumber()] = new BigNumber((_90_v69).length);
        _nw22[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_90_v69, _90_v69)).length);
        _nw22[(new BigNumber(19)).toNumber()] = _2_v0;
        _nw22[(new BigNumber(20)).toNumber()] = new BigNumber((_97_v76).length);
        _nw22[(new BigNumber(21)).toNumber()] = _2_v0;
        _nw22[(new BigNumber(22)).toNumber()] = _2_v0;
        _nw22[(new BigNumber(23)).toNumber()] = new BigNumber((_3_v1).length);
        _nw22[(new BigNumber(24)).toNumber()] = _2_v0;
        _98_v77 = _nw22;
        let _rhs11 = _98_v77;
        let _rhs12 = _77_v56;
        _91_v70 = _rhs11;
        _77_v56 = _rhs12;
        let _99_v78;
        _99_v78 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _100_v79;
        _100_v79 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_2_v0, _2_v0, _2_v0),_module.__default.fm0((((_85_v64).contains(p2)) ? ((_85_v64).get(p2)) : (_2_v0)), p1, _99_v78, globalState));
        _100_v79 = (_100_v79).update((_dafny.Set.fromElements(_2_v0, _2_v0)).Intersect(_dafny.Set.fromElements(_2_v0, _2_v0)), (_78_v57)[_module.__default.safeIndex(_2_v0, new BigNumber((_78_v57).length))]);
        _85_v64 = _85_v64;
        let _101_v80;
        let _nw23 = new _module.C0();
        _nw23.__ctor();
        _101_v80 = _nw23;
      }
      let _102_v81;
      _102_v81 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
      let _103_v82;
      _103_v82 = _102_v81;
      let _pat_let_tv0 = p2;
      let _rhs13 = function (_source1) {
        let _104___mcc_h1 = _source1;
        let _105_cf1 = _104___mcc_h1;
        return _pat_let_tv0;
      }(_103_v82);
      let _rhs14 = p0;
      let _lhs7 = globalState;
      let _lhs8 = globalState;
      _lhs7.f7 = _rhs13;
      _lhs8.f12 = _rhs14;
      let _hi0 = new BigNumber((_dafny.Seq.Concat(_78_v57, _78_v57)).length);
      for (let _106_i9 = new BigNumber((_3_v1).length); _106_i9.isLessThan(_hi0); _106_i9 = _106_i9.plus(_dafny.ONE)) {
        (globalState).f3 = _2_v0;
        let _107_v83;
        _107_v83 = _dafny.Seq.of(p2);
        let _108_v84;
        _108_v84 = p1;
        let _109_v85;
        _109_v85 = _dafny.Map.Empty.slice().updateUnsafe(p1,_107_v83);
        let _110_v86;
        let _nw24 = Array((new BigNumber(25)).toNumber());
        _nw24[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_107_v83, _dafny.Seq.of(p0, p2));
        _nw24[(_dafny.ONE).toNumber()] = _107_v83;
        _nw24[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_107_v83, _107_v83);
        _nw24[(new BigNumber(3)).toNumber()] = _107_v83;
        _nw24[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_107_v83, _107_v83);
        _nw24[(new BigNumber(5)).toNumber()] = _107_v83;
        _nw24[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(p1);
        _nw24[(new BigNumber(7)).toNumber()] = _107_v83;
        _nw24[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_107_v83, _dafny.Seq.of(p0, _module.__default.fm2(_106_i9, p0, globalState)));
        _nw24[(new BigNumber(9)).toNumber()] = _107_v83;
        _nw24[(new BigNumber(10)).toNumber()] = _107_v83;
        _nw24[(new BigNumber(11)).toNumber()] = _107_v83;
        _nw24[(new BigNumber(12)).toNumber()] = _107_v83;
        _nw24[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_107_v83, _module.__default.safeIndex(_106_i9, new BigNumber((_107_v83).length)), p0);
        _nw24[(new BigNumber(14)).toNumber()] = _107_v83;
        _nw24[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_107_v83, _107_v83), _module.__default.safeIndex(_2_v0, new BigNumber((_dafny.Seq.Concat(_107_v83, _107_v83)).length)), !(p0)), _module.__default.safeIndex(_2_v0, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_107_v83, _107_v83), _module.__default.safeIndex(_2_v0, new BigNumber((_dafny.Seq.Concat(_107_v83, _107_v83)).length)), !(p0))).length)), p1);
        _nw24[(new BigNumber(16)).toNumber()] = _module.__default.fm9((_108_v84), globalState);
        _nw24[(new BigNumber(17)).toNumber()] = _107_v83;
        _nw24[(new BigNumber(18)).toNumber()] = _107_v83;
        _nw24[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(p0, p1, !(!(p0)), p2);
        _nw24[(new BigNumber(20)).toNumber()] = _107_v83;
        _nw24[(new BigNumber(21)).toNumber()] = _dafny.Seq.of(p1);
        _nw24[(new BigNumber(22)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_107_v83, _107_v83), _module.__default.safeIndex(_2_v0, new BigNumber((_dafny.Seq.Concat(_107_v83, _107_v83)).length)), p2);
        _nw24[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_107_v83, _107_v83);
        _nw24[(new BigNumber(24)).toNumber()] = (((_109_v85).contains((((_5_v3).contains(_106_i9)) ? ((_5_v3).get(_106_i9)) : (p1)))) ? ((_109_v85).get((((_5_v3).contains(_106_i9)) ? ((_5_v3).get(_106_i9)) : (p1)))) : (_dafny.Seq.of(p0)));
        _110_v86 = _nw24;
        let _index12 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_110_v86).length));
        (_110_v86)[_index12] = _107_v83;
        let _index13 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_110_v86).length));
        let _rhs15 = _dafny.Seq.Concat(_107_v83, _107_v83);
        let _rhs16 = true;
        let _rhs17 = p1;
        let _lhs9 = _110_v86;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_110_v86).length));
        let _lhs11 = globalState;
        let _lhs12 = globalState;
        _lhs9[_lhs10] = _rhs15;
        _lhs11.f12 = _rhs16;
        _lhs12.f7 = _rhs17;
        (globalState).f3 = _106_i9;
        (globalState).f5 = _106_i9;
      }
      r0 = _102_v81;
      return r0;
    }
    static Main(__noArgsParameter) {
      let _111_v0;
      _111_v0 = _dafny.Seq.UnicodeFromString("olqswci");
      let _112_v1;
      _112_v1 = true;
      let _113_v2;
      _113_v2 = new BigNumber(-891);
      let _114_globalState;
      let _nw25 = new _module.GlobalState();
      _nw25.__ctor(false, new BigNumber(115), true, new BigNumber(-895), _111_v0, new BigNumber(403), new BigNumber(751), false, _111_v0, new BigNumber(958), (_dafny.Map.Empty.slice().updateUnsafe(_112_v1,_113_v2)).update(_112_v1, _113_v2), true, false, false, _111_v0);
      _114_globalState = _nw25;
      if (_112_v1) {
        let _115_v3;
        _115_v3 = _dafny.Set.fromElements((_112_v1), _112_v1, _112_v1, _112_v1, _112_v1);
        (_114_globalState).f5 = ((true) ? ((_113_v2).minus(new BigNumber((_115_v3).length))) : (_113_v2));
        let _116_v4;
        _116_v4 = _dafny.MultiSet.fromElements(new BigNumber(-807));
        _112_v1 = (_116_v4).IsProperSubsetOf((_116_v4).update(_113_v2, _module.__default.abs(new BigNumber(837))));
        let _117_v5;
        let _nw26 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _117_v5 = _nw26;
        let _index14 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_117_v5).length));
        (_117_v5)[_index14] = _111_v0;
        let _118_v6;
        _118_v6 = _112_v1;
        let _119_v7;
        _119_v7 = _dafny.Seq.of(_112_v1, _112_v1);
        let _index15 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_117_v5).length));
        let _rhs18 = _dafny.Seq.IsProperPrefixOf(_119_v7, _119_v7);
        let _rhs19 = _113_v2;
        let _rhs20 = _dafny.Seq.Concat(_111_v0, _111_v0);
        let _rhs21 = _118_v6;
        let _rhs22 = _module.__default.safeModuloInt((_113_v2).plus(_113_v2), _module.__default.safeDivisionInt(_113_v2, _113_v2));
        let _lhs13 = _114_globalState;
        let _lhs14 = _114_globalState;
        let _lhs15 = _117_v5;
        let _lhs16 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_117_v5).length));
        let _lhs17 = _114_globalState;
        _lhs13.f12 = _rhs18;
        _lhs14.f3 = _rhs19;
        _lhs15[_lhs16] = _rhs20;
        _118_v6 = _rhs21;
        _lhs17.f5 = _rhs22;
        let _120_v8;
        let _init8 = ((_121_v1) => function (_122_i0) {
          return (_dafny.MultiSet.fromElements(_121_v1)).IsSubsetOf(_dafny.MultiSet.fromElements(_121_v1));
        })(_112_v1);
        let _nw27 = Array((new BigNumber(7)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw27.length); _i0_8++) {
          _nw27[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _120_v8 = _nw27;
        let _index16 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_120_v8).length));
        (_120_v8)[_index16] = _112_v1;
        let _index17 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_120_v8).length));
        (_120_v8)[_index17] = _112_v1;
        let _123_v9;
        let _nw28 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _123_v9 = _nw28;
        let _index18 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_123_v9).length));
        (_123_v9)[_index18] = _113_v2;
        let _index19 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_123_v9).length));
        let _rhs23 = _113_v2;
        let _rhs24 = new BigNumber(670);
        let _lhs18 = _123_v9;
        let _lhs19 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_123_v9).length));
        let _lhs20 = _114_globalState;
        _lhs18[_lhs19] = _rhs23;
        _lhs20.f3 = _rhs24;
      } else {
        let _124_v10;
        let _out0;
        _out0 = _module.__default.m0(true, _112_v1, _112_v1, _114_globalState);
        _124_v10 = _out0;
        let _125_v11;
        let _init9 = ((_126_v1) => function (_127_i1) {
          return _126_v1;
        })(_112_v1);
        let _nw29 = Array((new BigNumber(6)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw29.length); _i0_9++) {
          _nw29[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _125_v11 = _nw29;
        let _128_v12;
        _128_v12 = (_112_v1);
        let _index20 = _module.__default.safeIndex(new BigNumber(38), new BigNumber((_125_v11).length));
        (_125_v11)[_index20] = _128_v12;
        let _index21 = _module.__default.safeIndex(new BigNumber(38), new BigNumber((_125_v11).length));
        (_125_v11)[_index21] = ((_112_v1) ? (_128_v12) : (_128_v12));
        let _129_v13;
        _129_v13 = new _dafny.CodePoint('e'.codePointAt(0));
        let _130_v14;
        _130_v14 = _124_v10;
        _129_v13 = _module.__default.fm0(_113_v2, _112_v1, (_130_v14), _114_globalState);
        if (true) {
          let _131_v15;
          _131_v15 = _dafny.Seq.of(true);
          let _132_v16;
          _132_v16 = _dafny.Set.fromElements(_112_v1);
          let _133_v17;
          _133_v17 = _dafny.Set.fromElements(_module.__default.fm0(new BigNumber((_132_v16).length), _112_v1, _124_v10, _114_globalState), _129_v13);
          let _134_v18;
          _134_v18 = _dafny.Map.Empty.slice().updateUnsafe((_111_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_111_v0).length))],_113_v2);
          let _135_v20;
          _135_v20 = _dafny.Map.Empty.slice().updateUnsafe(!(_112_v1),_113_v2);
          let _136_v21;
          _136_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_112_v1,_112_v1),_113_v2),_113_v2);
          let _137_v22;
          _137_v22 = _dafny.Map.Empty.slice().updateUnsafe(_124_v10,new BigNumber((_131_v15).length));
          let _138_v23;
          _138_v23 = _dafny.Seq.of(_113_v2);
          let _139_v24;
          let _nw30 = Array((new BigNumber(26)).toNumber());
          _nw30[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Concat(_131_v15, _131_v15)).length);
          _nw30[(_dafny.ONE).toNumber()] = _113_v2;
          _nw30[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(_113_v2, _113_v2);
          _nw30[(new BigNumber(3)).toNumber()] = _113_v2;
          _nw30[(new BigNumber(4)).toNumber()] = new BigNumber(737);
          _nw30[(new BigNumber(5)).toNumber()] = new BigNumber((((_112_v1) ? (_133_v17) : (function () {
            let _coll1 = new _dafny.Set();
            for (const _compr_1 of (_134_v18).Keys.Elements) {
              let _140_v19 = _compr_1;
              if ((_134_v18).contains(_140_v19)) {
                _coll1.add(_140_v19);
              }
            }
            return _coll1;
          }()))).length);
          _nw30[(new BigNumber(6)).toNumber()] = _113_v2;
          _nw30[(new BigNumber(7)).toNumber()] = _113_v2;
          _nw30[(new BigNumber(8)).toNumber()] = _113_v2;
          _nw30[(new BigNumber(9)).toNumber()] = _113_v2;
          _nw30[(new BigNumber(10)).toNumber()] = _module.__default.fm1(!(_module.__default.fm2(_113_v2, _112_v1, _114_globalState)), _135_v20, _114_globalState);
          _nw30[(new BigNumber(11)).toNumber()] = _113_v2;
          _nw30[(new BigNumber(12)).toNumber()] = _113_v2;
          _nw30[(new BigNumber(13)).toNumber()] = _113_v2;
          _nw30[(new BigNumber(14)).toNumber()] = (_113_v2).minus(_113_v2);
          _nw30[(new BigNumber(15)).toNumber()] = _113_v2;
          _nw30[(new BigNumber(16)).toNumber()] = _module.__default.fm1(_112_v1, _dafny.Map.Empty.slice().updateUnsafe(_112_v1,_113_v2), _114_globalState);
          _nw30[(new BigNumber(17)).toNumber()] = (((_136_v21).contains(_137_v22)) ? ((_136_v21).get(_137_v22)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(847)), function (_141_i2) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })).length)));
          _nw30[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(_113_v2);
          _nw30[(new BigNumber(19)).toNumber()] = new BigNumber(-851);
          _nw30[(new BigNumber(20)).toNumber()] = _113_v2;
          _nw30[(new BigNumber(21)).toNumber()] = _113_v2;
          _nw30[(new BigNumber(22)).toNumber()] = (_113_v2).multipliedBy(_113_v2);
          _nw30[(new BigNumber(23)).toNumber()] = _113_v2;
          _nw30[(new BigNumber(24)).toNumber()] = ((_112_v1) ? ((_138_v23)[_module.__default.safeIndex(_module.__default.fm1(_112_v1, _135_v20, _114_globalState), new BigNumber((_138_v23).length))]) : ((_dafny.ZERO).minus(_113_v2)));
          _nw30[(new BigNumber(25)).toNumber()] = (_113_v2).plus(_113_v2);
          _139_v24 = _nw30;
          let _index22 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_139_v24).length));
          (_139_v24)[_index22] = _113_v2;
          let _index23 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_139_v24).length));
          (_139_v24)[_index23] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(115)), ((_142_v13) => function (_143_i3) {
            return _142_v13;
          })(_129_v13))).length);
          let _144_v25;
          let _nw31 = Array((new BigNumber(23)).toNumber()).fill(false);
          _144_v25 = _nw31;
          let _145_v26;
          _145_v26 = _dafny.Map.Empty.slice().updateUnsafe((_139_v24)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_139_v24).length))],!(true));
          let _index24 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_144_v25).length));
          (_144_v25)[_index24] = (((_145_v26).contains(_113_v2)) ? ((_145_v26).get(_113_v2)) : (_112_v1));
          let _index25 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_144_v25).length));
          (_144_v25)[_index25] = _112_v1;
          let _rhs25 = (((_144_v25)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_144_v25).length))]) ? (_module.__default.fm3(_112_v1, (_139_v24)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_139_v24).length))], _114_globalState)) : (_dafny.Seq.Concat(_dafny.Seq.update(_111_v0, _module.__default.safeIndex(_113_v2, new BigNumber((_111_v0).length)), _129_v13), _111_v0)));
          let _rhs26 = (_111_v0)[_module.__default.safeIndex(_113_v2, new BigNumber((_111_v0).length))];
          let _rhs27 = (_144_v25)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_144_v25).length))];
          let _rhs28 = _113_v2;
          let _rhs29 = _112_v1;
          let _lhs21 = _114_globalState;
          let _lhs22 = _114_globalState;
          _111_v0 = _rhs25;
          _129_v13 = _rhs26;
          _lhs21.f12 = _rhs27;
          _lhs22.f5 = _rhs28;
          _112_v1 = _rhs29;
          let _146_v27;
          _146_v27 = _dafny.Map.Empty.slice().updateUnsafe(((_144_v25)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_144_v25).length))]) === (_module.__default.fm2(new BigNumber((_111_v0).length), _112_v1, _114_globalState)),_111_v0);
          _146_v27 = _146_v27;
          let _147_v28;
          _147_v28 = _dafny.Seq.of(_dafny.Seq.of((_139_v24)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_139_v24).length))]), _138_v23, _dafny.Seq.of((_139_v24)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_139_v24).length))]), _dafny.Seq.update(_138_v23, _module.__default.safeIndex((_139_v24)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_139_v24).length))], new BigNumber((_138_v23).length)), _113_v2), _138_v23);
          let _148_v29;
          _148_v29 = _dafny.MultiSet.fromElements((_147_v28)[_module.__default.safeIndex((_139_v24)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_139_v24).length))], new BigNumber((_147_v28).length))]);
          let _149_v30;
          let _out1;
          _out1 = _module.__default.m0((_144_v25)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_144_v25).length))], (_144_v25)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_144_v25).length))], !((_148_v29).equals(_148_v29)), _114_globalState);
          _149_v30 = _out1;
        } else {
          let _150_v31;
          let _nw32 = new _module.C0();
          _nw32.__ctor();
          _150_v31 = _nw32;
          _130_v14 = _130_v14;
          let _151_v32;
          let _init10 = function (_152_i4) {
            return (_152_i4).multipliedBy(new BigNumber(-925));
          };
          let _nw33 = Array((new BigNumber(13)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw33.length); _i0_10++) {
            _nw33[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _151_v32 = _nw33;
          let _153_v33;
          _153_v33 = _dafny.Seq.of(_151_v32, _151_v32, _151_v32);
          let _154_v34;
          _154_v34 = _dafny.Map.Empty.slice().updateUnsafe(_124_v10,((_112_v1) ? ((_153_v33)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_111_v0, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_112_v1,_113_v2)).length), new BigNumber((_111_v0).length)), _129_v13)).length), new BigNumber((_153_v33).length))]) : (_151_v32)));
          _154_v34 = ((_154_v34).Merge(_154_v34)).Merge(_154_v34);
          let _155_v35;
          _155_v35 = _dafny.MultiSet.fromElements(_112_v1);
          let _156_v36;
          _156_v36 = _dafny.Map.Empty.slice().updateUnsafe(_113_v2,_module.__default.safeModuloInt(_113_v2, (((_155_v35).contains(false)) ? ((_155_v35).get(false)) : (_113_v2))));
          let _157_v37;
          _157_v37 = _dafny.Map.Empty.slice().updateUnsafe(_112_v1,_113_v2);
          _156_v36 = (_156_v36).update(_module.__default.fm1(_112_v1, _157_v37, _114_globalState), _113_v2);
          let _158_v38;
          let _init11 = ((_159_v0) => function (_160_i5) {
            return _dafny.Seq.Concat(_159_v0, _159_v0);
          })(_111_v0);
          let _nw34 = Array((new BigNumber(22)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw34.length); _i0_11++) {
            _nw34[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _158_v38 = _nw34;
          let _index26 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_158_v38).length));
          (_158_v38)[_index26] = _dafny.Seq.Concat(_111_v0, _dafny.Seq.UnicodeFromString("hydejw"));
          let _index27 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_158_v38).length));
          (_158_v38)[_index27] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(255)), function (_161_i6) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          });
        }
        (_114_globalState).f7 = _112_v1;
      }
      let _162_i7;
      _162_i7 = _dafny.ZERO;
      L1: {
        while (true) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_162_i7)) {
              break L1;
            }
            _162_i7 = (_162_i7).plus(_dafny.ONE);
            let _163_v39;
            _163_v39 = _dafny.Set.fromElements(_112_v1, false, _112_v1, _112_v1);
            let _164_v40;
            _164_v40 = _dafny.MultiSet.fromElements(new BigNumber((_111_v0).length));
            let _165_v41;
            _165_v41 = _dafny.Seq.of(_113_v2);
            let _166_v42;
            let _nw35 = Array((new BigNumber(16)).toNumber());
            _nw35[(_dafny.ZERO).toNumber()] = _112_v1;
            _nw35[(_dafny.ONE).toNumber()] = ((_112_v1) ? (_112_v1) : (_112_v1));
            _nw35[(new BigNumber(2)).toNumber()] = false;
            _nw35[(new BigNumber(3)).toNumber()] = _112_v1;
            _nw35[(new BigNumber(4)).toNumber()] = _112_v1;
            _nw35[(new BigNumber(5)).toNumber()] = _112_v1;
            _nw35[(new BigNumber(6)).toNumber()] = (_113_v2).isLessThan((_dafny.ZERO).minus(new BigNumber((_163_v39).length)));
            _nw35[(new BigNumber(7)).toNumber()] = (_164_v40).IsDisjointFrom((_dafny.MultiSet.fromElements(_113_v2, _113_v2)).update(_113_v2, _module.__default.abs(_113_v2)));
            _nw35[(new BigNumber(8)).toNumber()] = !(_dafny.Seq.IsProperPrefixOf(_165_v41, _165_v41));
            _nw35[(new BigNumber(9)).toNumber()] = _112_v1;
            _nw35[(new BigNumber(10)).toNumber()] = _112_v1;
            _nw35[(new BigNumber(11)).toNumber()] = _112_v1;
            _nw35[(new BigNumber(12)).toNumber()] = _112_v1;
            _nw35[(new BigNumber(13)).toNumber()] = _112_v1;
            _nw35[(new BigNumber(14)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_112_v1,_112_v1)).contains(_112_v1);
            _nw35[(new BigNumber(15)).toNumber()] = _112_v1;
            _166_v42 = _nw35;
            let _167_v43;
            _167_v43 = _dafny.Seq.of(_112_v1);
            let _168_v44;
            let _nw36 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
            _168_v44 = _nw36;
            let _rhs30 = _166_v42;
            let _rhs31 = _167_v43;
            let _rhs32 = _168_v44;
            let _rhs33 = (_113_v2).minus((_113_v2).minus(_113_v2));
            _166_v42 = _rhs30;
            _167_v43 = _rhs31;
            _168_v44 = _rhs32;
            _113_v2 = _rhs33;
            let _index28 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_166_v42).length));
            (_166_v42)[_index28] = _112_v1;
            let _index29 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_166_v42).length));
            (_166_v42)[_index29] = (_164_v40).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_111_v0).length)));
            (_114_globalState).f12 = (_166_v42)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_166_v42).length))];
            let _169_v45;
            _169_v45 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_166_v42)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_166_v42).length))],_module.__default.fm2(_113_v2, _112_v1, _114_globalState))).length);
            let _170_v46;
            _170_v46 = _dafny.Set.fromElements((_169_v45), (_dafny.ZERO).minus(_113_v2), _113_v2);
            let _171_v47;
            _171_v47 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_113_v2),_113_v2);
            _170_v46 = (_170_v46).Difference((_dafny.Set.fromElements(new BigNumber((_171_v47).length), _113_v2)).Difference(_170_v46));
          }
        }
      }
      let _172_v48;
      let _init12 = ((_173_v1) => function (_174_i8) {
        return !(_173_v1);
      })(_112_v1);
      let _nw37 = Array((new BigNumber(11)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw37.length); _i0_12++) {
        _nw37[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _172_v48 = _nw37;
      let _175_v49;
      _175_v49 = _dafny.Map.Empty.slice().updateUnsafe(_172_v48,_112_v1);
      _175_v49 = (_175_v49).update(_172_v48, _112_v1);
      let _hi1 = _113_v2;
      for (let _176_i9 = _113_v2; _176_i9.isLessThan(_hi1); _176_i9 = _176_i9.plus(_dafny.ONE)) {
        let _177_v50;
        let _nw38 = new _module.C0();
        _nw38.__ctor();
        _177_v50 = _nw38;
        let _178_v51;
        _178_v51 = _112_v1;
        let _179_v52;
        _179_v52 = _dafny.Set.fromElements(_112_v1);
        let _rhs34 = _179_v52;
        let _rhs35 = _112_v1;
        let _lhs23 = _114_globalState;
        _179_v52 = _rhs34;
        _lhs23.f7 = _rhs35;
        let _180_v53;
        let _out2;
        _out2 = _module.__default.m0(false, _112_v1, _112_v1, _114_globalState);
        _180_v53 = _out2;
        let _181_v54;
        _181_v54 = _dafny.MultiSet.fromElements(new BigNumber(791));
        let _182_v55;
        _182_v55 = _dafny.Seq.of(_181_v54, _dafny.MultiSet.fromElements(_113_v2, _113_v2), _181_v54);
        let _183_v56;
        _183_v56 = _dafny.Seq.of(_176_i9, new BigNumber(736));
        let _rhs36 = (new BigNumber(858)).minus(_113_v2);
        let _rhs37 = _176_i9;
        let _rhs38 = _112_v1;
        let _rhs39 = _112_v1;
        let _rhs40 = ((_182_v55)[_module.__default.safeIndex(_113_v2, new BigNumber((_182_v55).length))]).update(((_112_v1) ? (_113_v2) : ((_183_v56)[_module.__default.safeIndex(_113_v2, new BigNumber((_183_v56).length))])), _module.__default.abs(_113_v2));
        let _lhs24 = _114_globalState;
        let _lhs25 = _114_globalState;
        _lhs24.f5 = _rhs36;
        _lhs25.f5 = _rhs37;
        _112_v1 = _rhs38;
        _112_v1 = _rhs39;
        _181_v54 = _rhs40;
      }
      let _184_v57;
      _184_v57 = _dafny.Map.Empty.slice().updateUnsafe(_112_v1,_113_v2);
      let _hi2 = (_dafny.ZERO).minus(_module.__default.fm1(_112_v1, _184_v57, _114_globalState));
      for (let _185_i10 = _113_v2; _185_i10.isLessThan(_hi2); _185_i10 = _185_i10.plus(_dafny.ONE)) {
        if (_112_v1) {
          let _186_v58;
          _186_v58 = _dafny.Seq.of(_111_v0);
          let _187_v59;
          let _nw39 = new _module.C0();
          _nw39.__ctor();
          _187_v59 = _nw39;
          let _188_v60;
          _188_v60 = _dafny.Set.fromElements(_187_v59);
          let _189_v61;
          _189_v61 = _dafny.Seq.of(_188_v60, _188_v60);
          let _190_v62;
          _190_v62 = _111_v0;
          let _191_v63;
          _191_v63 = _112_v1;
          let _192_v64;
          _192_v64 = new _dafny.CodePoint('b'.codePointAt(0));
          let _193_v65;
          let _nw40 = Array((new BigNumber(20)).toNumber());
          _nw40[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("apenmlek");
          _nw40[(_dafny.ONE).toNumber()] = _module.__default.fm3(_112_v1, (_dafny.ZERO).minus(_185_i10), _114_globalState);
          _nw40[(new BigNumber(2)).toNumber()] = _111_v0;
          _nw40[(new BigNumber(3)).toNumber()] = _111_v0;
          _nw40[(new BigNumber(4)).toNumber()] = _111_v0;
          _nw40[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("jxraijmpf");
          _nw40[(new BigNumber(6)).toNumber()] = (_186_v58)[_module.__default.safeIndex(new BigNumber(((_189_v61)[_module.__default.safeIndex(_113_v2, new BigNumber((_189_v61).length))]).length), new BigNumber((_186_v58).length))];
          _nw40[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat((_190_v62), _111_v0);
          _nw40[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("vqhbcmux");
          _nw40[(new BigNumber(9)).toNumber()] = _111_v0;
          _nw40[(new BigNumber(10)).toNumber()] = _dafny.Seq.update((_190_v62), _module.__default.safeIndex(_185_i10, new BigNumber(((_190_v62)).length)), new _dafny.CodePoint('o'.codePointAt(0)));
          _nw40[(new BigNumber(11)).toNumber()] = _111_v0;
          _nw40[(new BigNumber(12)).toNumber()] = _111_v0;
          _nw40[(new BigNumber(13)).toNumber()] = _111_v0;
          _nw40[(new BigNumber(14)).toNumber()] = _111_v0;
          _nw40[(new BigNumber(15)).toNumber()] = _module.__default.fm3((_187_v59).fm4(_112_v1, _114_globalState), _185_i10, _114_globalState);
          _nw40[(new BigNumber(16)).toNumber()] = _111_v0;
          _nw40[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_module.__default.fm3(!((_191_v63)), (_dafny.ZERO).minus(_113_v2), _114_globalState), _module.__default.safeIndex(_185_i10, new BigNumber((_module.__default.fm3(!((_191_v63)), (_dafny.ZERO).minus(_113_v2), _114_globalState)).length)), _192_v64);
          _nw40[(new BigNumber(18)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(247)), ((_194_v64) => function (_195_i11) {
            return _194_v64;
          })(_192_v64));
          _nw40[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_111_v0, _111_v0);
          _193_v65 = _nw40;
          let _index30 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_193_v65).length));
          (_193_v65)[_index30] = _dafny.Seq.UnicodeFromString("evqq");
          let _index31 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_193_v65).length));
          (_193_v65)[_index31] = _111_v0;
          (_114_globalState).f7 = _112_v1;
          let _196_v66;
          _196_v66 = _dafny.Map.Empty.slice().updateUnsafe(_113_v2,_187_v59);
          let _197_v67;
          _197_v67 = _dafny.Seq.of(((_112_v1) ? (_187_v59) : ((((_196_v66).contains(_113_v2)) ? ((_196_v66).get(_113_v2)) : (_187_v59)))), _187_v59);
          _187_v59 = (_197_v67)[_module.__default.safeIndex(_113_v2, new BigNumber((_197_v67).length))];
          let _198_v68;
          let _nw41 = new _module.C0();
          _nw41.__ctor();
          _198_v68 = _nw41;
          let _199_v69;
          let _nw42 = Array((new BigNumber(16)).toNumber()).fill([]);
          _199_v69 = _nw42;
          let _200_v70;
          let _nw43 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _200_v70 = _nw43;
          let _index32 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_199_v69).length));
          (_199_v69)[_index32] = _200_v70;
          let _index33 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_199_v69).length));
          (_199_v69)[_index33] = _200_v70;
        } else {
          let _201_v71;
          let _out3;
          _out3 = _module.__default.m0(_112_v1, _112_v1, _112_v1, _114_globalState);
          _201_v71 = _out3;
          (_114_globalState).f7 = _112_v1;
          (_114_globalState).f7 = _dafny.Seq.IsProperPrefixOf(_111_v0, _111_v0);
          let _202_v72;
          let _nw44 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _202_v72 = _nw44;
          let _index34 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_202_v72).length));
          (_202_v72)[_index34] = _dafny.ONE;
          let _index35 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_202_v72).length));
          (_202_v72)[_index35] = (_185_i10).minus(_module.__default.safeDivisionInt(_185_i10, _113_v2));
          (_114_globalState).f5 = _module.__default.fm1(_dafny.areEqual(_111_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(170)), function (_203_i12) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          })), _184_v57, _114_globalState);
        }
        let _204_v73;
        _204_v73 = _dafny.Seq.of(_185_i10, _113_v2);
        let _205_v74;
        _205_v74 = _dafny.Seq.of(_204_v73, _204_v73);
        let _206_v75;
        let _out4;
        _out4 = _module.__default.m0(false, _dafny.areEqual(_205_v74, _dafny.Seq.of(_204_v73, _dafny.Seq.of(_185_i10, _185_i10, _113_v2))), _112_v1, _114_globalState);
        _206_v75 = _out4;
        (_114_globalState).f12 = _112_v1;
        let _207_v76;
        let _nw45 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
        _207_v76 = _nw45;
        let _208_v77;
        let _nw46 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _208_v77 = _nw46;
        let _index36 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_208_v77).length));
        (_208_v77)[_index36] = (_185_i10).multipliedBy(_185_i10);
        let _209_v78;
        _209_v78 = _dafny.MultiSet.fromElements(_112_v1);
        let _210_v79;
        _210_v79 = _dafny.MultiSet.fromElements((((_209_v78).contains(_112_v1)) ? ((_209_v78).get(_112_v1)) : (_113_v2)));
        let _index37 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_208_v77).length));
        (_208_v77)[_index37] = new BigNumber((_210_v79).cardinality());
        let _211_v80;
        let _nw47 = new _module.C0();
        _nw47.__ctor();
        _211_v80 = _nw47;
        let _212_v81;
        _212_v81 = _dafny.Map.Empty.slice().updateUnsafe(_211_v80,_dafny.Seq.UnicodeFromString("pqq"));
        let _index38 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_208_v77).length));
        let _index39 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_208_v77).length));
        let _rhs41 = (_dafny.ZERO).minus(new BigNumber(((((_212_v81).contains(_211_v80)) ? ((_212_v81).get(_211_v80)) : (_111_v0))).length));
        let _rhs42 = _207_v76;
        let _rhs43 = ((_dafny.ZERO).minus(new BigNumber((_111_v0).length))).plus(_113_v2);
        let _rhs44 = new BigNumber((_111_v0).length);
        let _lhs26 = _114_globalState;
        let _lhs27 = _208_v77;
        let _lhs28 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_208_v77).length));
        let _lhs29 = _208_v77;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_208_v77).length));
        _lhs26.f5 = _rhs41;
        _207_v76 = _rhs42;
        _lhs27[_lhs28] = _rhs43;
        _lhs29[_lhs30] = _rhs44;
      }
      if (true) {
        let _213_v82;
        let _nw48 = new _module.C0();
        _nw48.__ctor();
        _213_v82 = _nw48;
        _184_v57 = (_184_v57).update(_112_v1, _module.__default.safeModuloInt(_113_v2, _113_v2));
        let _214_v83;
        _214_v83 = _dafny.MultiSet.fromElements(_172_v48);
        let _215_v84;
        _215_v84 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("uknqeb")).length), _113_v2);
        let _216_v85;
        _216_v85 = _dafny.Map.Empty.slice().updateUnsafe(_215_v84,(_214_v83).update(_172_v48, _module.__default.abs(_113_v2)));
        _214_v83 = (((_216_v85).contains(_215_v84)) ? ((_216_v85).get(_215_v84)) : ((_214_v83).Union(_214_v83)));
        (_114_globalState).f7 = (false) || (_112_v1);
        let _217_v86;
        _217_v86 = _dafny.Map.Empty.slice().updateUnsafe(_113_v2,_113_v2);
        let _218_v87;
        _218_v87 = _dafny.MultiSet.fromElements(_112_v1);
        _217_v86 = (_217_v86).update(new BigNumber(467), new BigNumber((((!(_112_v1)) ? (_218_v87) : (_218_v87))).cardinality()));
      } else {
        let _219_v88;
        _219_v88 = _dafny.MultiSet.fromElements(_112_v1);
        let _220_v89;
        _220_v89 = _dafny.Seq.of(_113_v2);
        let _221_v90;
        _221_v90 = _dafny.Set.fromElements(_172_v48);
        let _222_v91;
        _222_v91 = new _dafny.CodePoint('d'.codePointAt(0));
        let _223_v92;
        _223_v92 = _dafny.Seq.of(_dafny.Set.fromElements(_222_v91));
        let _224_v93;
        _224_v93 = _dafny.Map.Empty.slice().updateUnsafe(_223_v92,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-475)), function (_225_i14) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        }));
        let _226_v94;
        _226_v94 = _dafny.Seq.of(_112_v1);
        let _227_v95;
        let _nw49 = Array((new BigNumber(4)).toNumber());
        _nw49[(_dafny.ZERO).toNumber()] = _113_v2;
        _nw49[(_dafny.ONE).toNumber()] = _113_v2;
        _nw49[(new BigNumber(2)).toNumber()] = new BigNumber((_226_v94).length);
        _nw49[(new BigNumber(3)).toNumber()] = _113_v2;
        _227_v95 = _nw49;
        let _228_v96;
        _228_v96 = _dafny.Map.Empty.slice().updateUnsafe(_227_v95,_112_v1);
        let _229_v97;
        let _nw50 = Array((new BigNumber(26)).toNumber());
        _nw50[(_dafny.ZERO).toNumber()] = _113_v2;
        _nw50[(_dafny.ONE).toNumber()] = _113_v2;
        _nw50[(new BigNumber(2)).toNumber()] = _113_v2;
        _nw50[(new BigNumber(3)).toNumber()] = _113_v2;
        _nw50[(new BigNumber(4)).toNumber()] = _113_v2;
        _nw50[(new BigNumber(5)).toNumber()] = new BigNumber(-621);
        _nw50[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-603)), function (_230_i13) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        })).length), new BigNumber((_219_v88).cardinality()));
        _nw50[(new BigNumber(7)).toNumber()] = _module.__default.fm1(_112_v1, _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(742)), _114_globalState);
        _nw50[(new BigNumber(8)).toNumber()] = _113_v2;
        _nw50[(new BigNumber(9)).toNumber()] = new BigNumber(-310);
        _nw50[(new BigNumber(10)).toNumber()] = _113_v2;
        _nw50[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_112_v1)).length);
        _nw50[(new BigNumber(12)).toNumber()] = new BigNumber((_184_v57).length);
        _nw50[(new BigNumber(13)).toNumber()] = _113_v2;
        _nw50[(new BigNumber(14)).toNumber()] = ((!(_112_v1)) ? (new BigNumber((_220_v89).length)) : ((_220_v89)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("oe")).length)), new BigNumber((_220_v89).length))]));
        _nw50[(new BigNumber(15)).toNumber()] = _module.__default.fm1(_112_v1, _184_v57, _114_globalState);
        _nw50[(new BigNumber(16)).toNumber()] = _113_v2;
        _nw50[(new BigNumber(17)).toNumber()] = _module.__default.fm1(_112_v1, _184_v57, _114_globalState);
        _nw50[(new BigNumber(18)).toNumber()] = _module.__default.safeDivisionInt(_113_v2, _113_v2);
        _nw50[(new BigNumber(19)).toNumber()] = _113_v2;
        _nw50[(new BigNumber(20)).toNumber()] = _module.__default.safeModuloInt(_113_v2, new BigNumber((_221_v90).length));
        _nw50[(new BigNumber(21)).toNumber()] = _113_v2;
        _nw50[(new BigNumber(22)).toNumber()] = new BigNumber(((((_224_v93).contains(_223_v92)) ? ((_224_v93).get(_223_v92)) : (_111_v0))).length);
        _nw50[(new BigNumber(23)).toNumber()] = _113_v2;
        _nw50[(new BigNumber(24)).toNumber()] = _113_v2;
        _nw50[(new BigNumber(25)).toNumber()] = new BigNumber((_228_v96).length);
        _229_v97 = _nw50;
        let _index40 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_227_v95).length));
        (_227_v95)[_index40] = _113_v2;
        let _231_v98;
        _231_v98 = _113_v2;
        let _index41 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_227_v95).length));
        (_227_v95)[_index41] = (_231_v98);
        let _232_v99;
        _232_v99 = _dafny.Set.fromElements(((_112_v1) ? ((_227_v95)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_227_v95).length))]) : ((_227_v95)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_227_v95).length))])));
        let _233_v100;
        _233_v100 = _112_v1;
        let _234_v101;
        let _nw51 = Array((new BigNumber(7)).toNumber());
        _nw51[(_dafny.ZERO).toNumber()] = _module.__default.fm2(_113_v2, _112_v1, _114_globalState);
        _nw51[(_dafny.ONE).toNumber()] = _233_v100;
        _nw51[(new BigNumber(2)).toNumber()] = _233_v100;
        _nw51[(new BigNumber(3)).toNumber()] = _233_v100;
        _nw51[(new BigNumber(4)).toNumber()] = _112_v1;
        _nw51[(new BigNumber(5)).toNumber()] = _module.__default.fm5(false, new BigNumber(549), _112_v1, _113_v2, _114_globalState);
        _nw51[(new BigNumber(6)).toNumber()] = _112_v1;
        _234_v101 = _nw51;
        let _index42 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_234_v101).length));
        (_234_v101)[_index42] = _233_v100;
        let _235_v102;
        let _nw52 = new _module.C0();
        _nw52.__ctor();
        _235_v102 = _nw52;
        let _236_v103;
        _236_v103 = _dafny.Seq.of(_235_v102);
        let _index43 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_234_v101).length));
        let _rhs45 = _232_v99;
        let _rhs46 = !(_112_v1) || (_dafny.areEqual(_dafny.Seq.of(_235_v102), _236_v103));
        let _rhs47 = !(_module.__default.fm6(_112_v1, _112_v1, _114_globalState)).equals(_184_v57);
        let _rhs48 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.of(!(_112_v1)), _226_v94), _226_v94);
        let _lhs31 = _114_globalState;
        let _lhs32 = _234_v101;
        let _lhs33 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_234_v101).length));
        _232_v99 = _rhs45;
        _lhs31.f12 = _rhs46;
        _lhs32[_lhs33] = _rhs47;
        _112_v1 = _rhs48;
        let _237_v104;
        let _out5;
        _out5 = _module.__default.m0(_dafny.Seq.contains(_220_v89, (_227_v95)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_227_v95).length))]), _112_v1, _112_v1, _114_globalState);
        _237_v104 = _out5;
        let _238_v105;
        let _nw53 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Set.Empty);
        _238_v105 = _nw53;
        let _239_v106;
        _239_v106 = _dafny.Set.fromElements(_112_v1);
        let _index44 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_238_v105).length));
        (_238_v105)[_index44] = _239_v106;
        let _index45 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_238_v105).length));
        let _rhs49 = _dafny.MultiSet.fromElements(true);
        let _rhs50 = (_239_v106).Difference(((_112_v1) ? (_239_v106) : (_239_v106)));
        let _rhs51 = _113_v2;
        let _lhs34 = _238_v105;
        let _lhs35 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_238_v105).length));
        let _lhs36 = _114_globalState;
        _219_v88 = _rhs49;
        _lhs34[_lhs35] = _rhs50;
        _lhs36.f3 = _rhs51;
        let _240_v107;
        let _nw54 = new _module.C0();
        _nw54.__ctor();
        _240_v107 = _nw54;
      }
      let _241_v108;
      _241_v108 = _dafny.Map.Empty.slice().updateUnsafe(_112_v1,_112_v1);
      let _242_v109;
      _242_v109 = _dafny.Set.fromElements(_241_v108, _241_v108);
      let _243_v110;
      _243_v110 = new BigNumber((_242_v109).length);
      let _244_v111;
      let _nw55 = Array((new BigNumber(5)).toNumber());
      _nw55[(_dafny.ZERO).toNumber()] = (_243_v110);
      _nw55[(_dafny.ONE).toNumber()] = _module.__default.fm1(_112_v1, _184_v57, _114_globalState);
      _nw55[(new BigNumber(2)).toNumber()] = _113_v2;
      _nw55[(new BigNumber(3)).toNumber()] = _113_v2;
      _nw55[(new BigNumber(4)).toNumber()] = _113_v2;
      _244_v111 = _nw55;
      let _nw56 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _244_v111 = _nw56;
      let _245_v112;
      _245_v112 = new _dafny.CodePoint('j'.codePointAt(0));
      let _246_v113;
      _246_v113 = _111_v0;
      let _247_v114;
      _247_v114 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.contains(_111_v0, _245_v112),_dafny.Seq.Concat((_246_v113), _111_v0));
      let _248_v115;
      _248_v115 = _dafny.Map.Empty.slice().updateUnsafe(_112_v1,_112_v1);
      _247_v114 = _module.__default.fm7(_248_v115, _114_globalState);
      _112_v1 = _112_v1;
      let _249_v116;
      let _out6;
      _out6 = _module.__default.m0(_112_v1, !(_112_v1), _112_v1, _114_globalState);
      _249_v116 = _out6;
      let _pat_let_tv1 = _113_v2;
      let _pat_let_tv2 = _112_v1;
      let _pat_let_tv3 = _112_v1;
      _113_v2 = function (_source2) {
        let _250___mcc_h0 = _source2;
        let _251_cf3 = _250___mcc_h0;
        return (_pat_let_tv1).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv2,_dafny.Seq.of(_pat_let_tv3))).length));
      }(_246_v113);
      if (!(_112_v1)) {
        (_114_globalState).f7 = _112_v1;
        let _252_v117;
        _252_v117 = _dafny.Seq.of(_244_v111);
        _244_v111 = (((_112_v1) && (_112_v1)) ? (_244_v111) : ((_252_v117)[_module.__default.safeIndex(_113_v2, new BigNumber((_252_v117).length))]));
        let _253_v118;
        _253_v118 = _dafny.Map.Empty.slice().updateUnsafe(!(_113_v2).isEqualTo((((_184_v57).contains(_112_v1)) ? ((_184_v57).get(_112_v1)) : (_113_v2))),_244_v111);
        _253_v118 = (_253_v118).update(!(false), _244_v111);
        (_114_globalState).f3 = _113_v2;
        (_114_globalState).f8 = (((((_248_v115).contains(_112_v1)) ? ((_248_v115).get(_112_v1)) : (_112_v1))) ? (_dafny.Seq.Concat(_dafny.Seq.update(_111_v0, _module.__default.safeIndex(_113_v2, new BigNumber((_111_v0).length)), _245_v112), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(272)), ((_254_v112) => function (_255_i15) {
          return _254_v112;
        })(_245_v112)), _module.__default.safeIndex(_113_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(272)), ((_256_v112) => function (_257_i15) {
          return _256_v112;
        })(_245_v112))).length)), _245_v112))) : (_111_v0));
      } else {
        let _258_v119;
        let _nw57 = new _module.C0();
        _nw57.__ctor();
        _258_v119 = _nw57;
        let _259_v120;
        _259_v120 = _dafny.Map.Empty.slice().updateUnsafe(_258_v119,!(_112_v1));
        let _rhs52 = _module.__default.fm2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_112_v1,_258_v119)).length), false, _114_globalState);
        let _rhs53 = _258_v119;
        let _rhs54 = ((_112_v1) ? (_112_v1) : ((((_259_v120).contains(_258_v119)) ? ((_259_v120).get(_258_v119)) : (_112_v1))));
        let _lhs37 = _114_globalState;
        _112_v1 = _rhs52;
        _258_v119 = _rhs53;
        _lhs37.f7 = _rhs54;
        let _260_v121;
        _260_v121 = _112_v1;
        _113_v2 = new BigNumber((((_249_v116).Merge(_248_v115)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_112_v1,(_260_v121)))).length);
        let _261_v122;
        _261_v122 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_258_v119)).length));
        _245_v112 = _module.__default.fm0(new BigNumber((_261_v122).cardinality()), _112_v1, (_249_v116).Merge((_249_v116).update(_112_v1, _112_v1)), _114_globalState);
        (_114_globalState).f7 = _112_v1;
        (_114_globalState).f7 = true;
      }
      let _262_i16;
      _262_i16 = _dafny.ZERO;
      L2: {
        while (_112_v1) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_262_i16)) {
              break L2;
            }
            _262_i16 = (_262_i16).plus(_dafny.ONE);
            _112_v1 = _112_v1;
            let _263_v123;
            let _nw58 = new _module.C0();
            _nw58.__ctor();
            _263_v123 = _nw58;
            _263_v123 = _263_v123;
            let _264_v124;
            let _nw59 = new _module.C0();
            _nw59.__ctor();
            _264_v124 = _nw59;
            let _265_v125;
            let _out7;
            _out7 = _module.__default.m0(_112_v1, _112_v1, (_112_v1) || (_112_v1), _114_globalState);
            _265_v125 = _out7;
          }
        }
      }
      let _hi3 = (_113_v2).minus(_113_v2);
      for (let _266_i17 = _113_v2; _266_i17.isLessThan(_hi3); _266_i17 = _266_i17.plus(_dafny.ONE)) {
        let _source3 = _246_v113;
        let _267___mcc_h1 = _source3;
        let _268_cf3 = _267___mcc_h1;
        (_114_globalState).f5 = _266_i17;
        _184_v57 = (_184_v57).update(!(!(_112_v1) || (_112_v1)), _266_i17);
        let _index46 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_172_v48).length));
        (_172_v48)[_index46] = _module.__default.fm2(_113_v2, _112_v1, _114_globalState);
        let _index47 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_172_v48).length));
        (_172_v48)[_index47] = _112_v1;
        (_114_globalState).f7 = _module.__default.fm2((_113_v2).plus(_266_i17), !(_112_v1), _114_globalState);
        let _269_v126;
        _269_v126 = _dafny.Set.fromElements((_dafny.ZERO).minus(_266_i17));
        (_114_globalState).f12 = !((_269_v126).IsProperSubsetOf(_dafny.Set.fromElements(_113_v2)));
        let _270_v127;
        let _nw60 = new _module.C0();
        _nw60.__ctor();
        _270_v127 = _nw60;
        let _271_v128;
        _271_v128 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(433)), function (_272_i18) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        }),_112_v1);
        let _273_v129;
        _273_v129 = _dafny.Seq.of(_112_v1);
        let _rhs55 = _dafny.Seq.Concat((((((_271_v128).contains(_111_v0)) ? ((_271_v128).get(_111_v0)) : (false))) ? (_111_v0) : (_111_v0)), _111_v0);
        let _rhs56 = (((_184_v57).contains(!(true))) ? ((_184_v57).get(!(true))) : (new BigNumber((_273_v129).length)));
        let _rhs57 = _270_v127;
        let _rhs58 = _dafny.Seq.IsPrefixOf(_111_v0, _dafny.Seq.Concat(_111_v0, _111_v0));
        let _lhs38 = _114_globalState;
        _111_v0 = _rhs55;
        _lhs38.f5 = _rhs56;
        _270_v127 = _rhs57;
        _112_v1 = _rhs58;
        let _274_v130;
        let _init13 = ((_275_v112) => function (_276_i20) {
          return _275_v112;
        })(_245_v112);
        let _nw61 = Array((new BigNumber(27)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw61.length); _i0_13++) {
          _nw61[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _274_v130 = _nw61;
        let _277_v131;
        _277_v131 = _dafny.Map.Empty.slice().updateUnsafe(_274_v130,_112_v1);
        let _index48 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_172_v48).length));
        (_172_v48)[_index48] = (new BigNumber((_277_v131).length)).isLessThan(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-295)), ((_278_i17) => function (_279_i19) {
          return _278_i17;
        })(_266_i17))).length));
        let _index49 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_172_v48).length));
        (_172_v48)[_index49] = !(((_112_v1) ? ((_113_v2).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("dvpiwucvt")).length))) : (false)));
      }
      _112_v1 = !(true) || (_112_v1);
      (_114_globalState).f12 = !(_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_111_v0, _111_v0), _111_v0));
      process.stdout.write((_111_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_112_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_113_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_114_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_114_globalState).f4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_114_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_114_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_114_globalState.f8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_114_globalState).f10).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-891)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_114_globalState.f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_114_globalState).f14).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_162_i7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v48)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v48)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v48)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v48)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v48)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v48)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v48)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v48)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v48)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v48)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v48)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_175_v49).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_184_v57).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_241_v108)).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v109).equals(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v110)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_245_v112));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_246_v113)).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_247_v114).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_248_v115).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_249_v116).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_262_i16));
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
        return other.$tag === 0 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return false;
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
    static create_DC1(cf1) {
      let $dt = new D1(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
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
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC2(cf2) {
      let $dt = new D2(0);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2);
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC3(cf3) {
      let $dt = new D3(0);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC3" + "(" + this.cf3.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf3, other.cf3);
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
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf4) {
      let $dt = new D4(0);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC4" + "(" + _dafny.toString(this.cf4) + ")";
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
      return _dafny.Map.Empty;
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
    static create_DC6() {
      let $dt = new D5(0);
      return $dt;
    }
    static create_DC5(cf5) {
      let $dt = new D5(1);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC7(cf6) {
      let $dt = new D5(2);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC6";
      } else if (this.$tag === 1) {
        return "D5.DC5" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC7" + "(" + _dafny.toString(this.cf6) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC6();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f3 = _dafny.ZERO;
      this.f5 = _dafny.ZERO;
      this.f7 = false;
      this.f8 = _dafny.Seq.UnicodeFromString("");
      this.f12 = false;
      this._f0 = false;
      this._f1 = _dafny.ZERO;
      this._f2 = false;
      this._f4 = _dafny.Seq.UnicodeFromString("");
      this._f6 = _dafny.ZERO;
      this._f9 = _dafny.ZERO;
      this._f10 = _dafny.Map.Empty;
      this._f11 = false;
      this._f13 = false;
      this._f14 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
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
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
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
    fm4(p0, globalState) {
      let _this = this;
      return true;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
