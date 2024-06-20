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
      return true;
    };
    static fm2(p0, globalState) {
      return _dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_module.D3.create_DC8()), _dafny.Seq.of(_module.D3.create_DC8(), _module.D3.create_DC8())))).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("gyruwxl")).length), (new BigNumber(123)).multipliedBy(new BigNumber(758)));
    };
    static fm3(p0, p1, p2, globalState) {
      return new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false)))).length);
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)))).Merge(function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (function () {
            let _coll2 = new _dafny.Map();
            for (const _compr_2 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), function (_0_i0) {
              return new _dafny.CodePoint('e'.codePointAt(0));
            })).Elements) {
              let _1_v1 = _compr_2;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), function (_0_i0) {
                return new _dafny.CodePoint('e'.codePointAt(0));
              }), _1_v1)) {
                _coll2.push([_1_v1,new BigNumber(904)]);
              }
            }
            return _coll2;
          }()).Keys.Elements) {
            let _2_v0 = _compr_1;
            if ((function () {
              let _coll3 = new _dafny.Map();
              for (const _compr_3 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), function (_0_i0) {
                return new _dafny.CodePoint('e'.codePointAt(0));
              })).Elements) {
                let _1_v1 = _compr_3;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), function (_0_i0) {
                  return new _dafny.CodePoint('e'.codePointAt(0));
                }), _1_v1)) {
                  _coll3.push([_1_v1,new BigNumber(904)]);
                }
              }
              return _coll3;
            }()).contains(_2_v0)) {
              _coll1.push([_2_v0,_dafny.Seq.of(new BigNumber(587))]);
            }
          }
          return _coll1;
        }())).Keys.Elements) {
          let _3_v2 = _compr_0;
          if (((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)))).Merge(function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of (function () {
              let _coll5 = new _dafny.Map();
              for (const _compr_5 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), function (_0_i0) {
                return new _dafny.CodePoint('e'.codePointAt(0));
              })).Elements) {
                let _1_v1 = _compr_5;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), function (_0_i0) {
                  return new _dafny.CodePoint('e'.codePointAt(0));
                }), _1_v1)) {
                  _coll5.push([_1_v1,new BigNumber(904)]);
                }
              }
              return _coll5;
            }()).Keys.Elements) {
              let _2_v0 = _compr_4;
              if ((function () {
                let _coll6 = new _dafny.Map();
                for (const _compr_6 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), function (_0_i0) {
                  return new _dafny.CodePoint('e'.codePointAt(0));
                })).Elements) {
                  let _1_v1 = _compr_6;
                  if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), function (_0_i0) {
                    return new _dafny.CodePoint('e'.codePointAt(0));
                  }), _1_v1)) {
                    _coll6.push([_1_v1,new BigNumber(904)]);
                  }
                }
                return _coll6;
              }()).contains(_2_v0)) {
                _coll4.push([_2_v0,_dafny.Seq.of(new BigNumber(587))]);
              }
            }
            return _coll4;
          }())).contains(_3_v2)) {
            _coll0.add(_3_v2);
          }
        }
        return _coll0;
      }();
    };
    static fm5(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(!(false), false)).cardinality())))).Difference(((false) ? (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(844))).length))) : (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(118)), function (_4_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length)))));
    };
    static fm6(globalState) {
      return new _dafny.CodePoint('q'.codePointAt(0));
    };
    static fm7(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(!(true), false, true), _dafny.Seq.of(true));
    };
    static fm8(globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("b"))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("bl"), _dafny.Seq.UnicodeFromString("oj")));
    };
    static Main(__noArgsParameter) {
      let _5_v0;
      let _nw0 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
      _5_v0 = _nw0;
      let _6_v1;
      _6_v1 = true;
      let _7_v2;
      _7_v2 = _dafny.Seq.of(!(_6_v1), _6_v1);
      let _8_v3;
      _8_v3 = new BigNumber(861);
      let _9_v4;
      _9_v4 = _dafny.Set.fromElements(_8_v3, _8_v3);
      let _10_v5;
      _10_v5 = false;
      let _11_v6;
      let _nw1 = Array((new BigNumber(12)).toNumber());
      _nw1[(_dafny.ZERO).toNumber()] = true;
      _nw1[(_dafny.ONE).toNumber()] = (_10_v5);
      _nw1[(new BigNumber(2)).toNumber()] = false;
      _nw1[(new BigNumber(3)).toNumber()] = (_7_v2)[_module.__default.safeIndex(_8_v3, new BigNumber((_7_v2).length))];
      _nw1[(new BigNumber(4)).toNumber()] = _6_v1;
      _nw1[(new BigNumber(5)).toNumber()] = _6_v1;
      _nw1[(new BigNumber(6)).toNumber()] = !(_6_v1);
      _nw1[(new BigNumber(7)).toNumber()] = _6_v1;
      _nw1[(new BigNumber(8)).toNumber()] = _6_v1;
      _nw1[(new BigNumber(9)).toNumber()] = _6_v1;
      _nw1[(new BigNumber(10)).toNumber()] = !(_6_v1);
      _nw1[(new BigNumber(11)).toNumber()] = _6_v1;
      _11_v6 = _nw1;
      let _12_v7;
      _12_v7 = new _dafny.CodePoint('l'.codePointAt(0));
      let _13_v8;
      _13_v8 = _dafny.Map.Empty.slice().updateUnsafe(_8_v3,_12_v7);
      let _14_v9;
      let _nw2 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
      _14_v9 = _nw2;
      let _15_v10;
      _15_v10 = _dafny.Map.Empty.slice().updateUnsafe(_6_v1,_12_v7);
      let _16_v11;
      _16_v11 = _dafny.MultiSet.fromElements(_15_v10, _15_v10);
      let _17_v12;
      _17_v12 = _module.D1.create_DC1(_5_v0);
      let _18_globalState;
      let _nw3 = new _module.GlobalState();
      _nw3.__ctor(_5_v0, _dafny.Seq.Concat(_7_v2, _7_v2), false, _9_v4, _11_v6, _5_v0, _13_v8, true, false, _14_v9, false, new _dafny.CodePoint('s'.codePointAt(0)), _11_v6, new BigNumber(242), true, new BigNumber(442), false, new BigNumber(-101), new BigNumber(92), new BigNumber(999), (_16_v11).Union(_16_v11), false, (_17_v12).dtor_cf1, (_17_v12).dtor_cf1, new BigNumber(279), new BigNumber(491));
      _18_globalState = _nw3;
      if ((new BigNumber((_dafny.Seq.of(_6_v1, _6_v1, _6_v1)).length)).isLessThanOrEqualTo(_module.__default.safeDivisionInt(_8_v3, _8_v3))) {
        _11_v6 = _11_v6;
        let _19_v13;
        _19_v13 = _dafny.Seq.UnicodeFromString("shmbaf");
        let _20_v14;
        _20_v14 = _dafny.MultiSet.fromElements(_19_v13, _19_v13, _19_v13);
        let _21_v15;
        _21_v15 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_dafny.Seq.of(_8_v3), _8_v3, (_dafny.ZERO).minus(new BigNumber((_19_v13).length)), _18_globalState),(_dafny.ZERO).minus(new BigNumber((_9_v4).length)));
        _8_v3 = (((_20_v14).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-845)), ((_24_v7) => function (_25_i0) {
          return _24_v7;
        })(_12_v7)))) ? ((_20_v14).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-845)), ((_22_v7) => function (_23_i0) {
          return _22_v7;
        })(_12_v7)))) : (((_6_v1) ? (new BigNumber((_21_v15).length)) : (new BigNumber(-872)))));
        _6_v1 = !_dafny.Seq.contains(_19_v13, _12_v7);
        if ((_9_v4).IsProperSubsetOf(_9_v4)) {
          (_18_globalState).f15 = _8_v3;
          _6_v1 = _6_v1;
          let _26_v16;
          let _nw4 = new _module.C0();
          _nw4.__ctor(_6_v1);
          _26_v16 = _nw4;
          (_26_v16).f26 = !((_module.__default.safeDivisionInt(_8_v3, new BigNumber(-365))).isLessThan(_8_v3));
          let _27_v17;
          let _28_v18;
          let _out0;
          let _out1;
          let _outcollector0 = (_26_v16).m0(_26_v16.f26, (_dafny.ZERO).minus(new BigNumber((_19_v13).length)), _18_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _27_v17 = _out0;
          _28_v18 = _out1;
        } else {
          let _29_v19;
          let _nw5 = new _module.C0();
          _nw5.__ctor(!(_8_v3).isEqualTo(_8_v3));
          _29_v19 = _nw5;
          let _30_v20;
          _30_v20 = _dafny.Map.Empty.slice().updateUnsafe(_8_v3,(_dafny.ZERO).minus(_8_v3));
          let _31_v21;
          _31_v21 = _module.D1.create_DC2(_8_v3, (_dafny.ZERO).minus(_8_v3), _8_v3, (_30_v20).Merge(_30_v20));
          _31_v21 = _31_v21;
          let _32_v22;
          let _nw6 = Array((new BigNumber(12)).toNumber());
          _nw6[(_dafny.ZERO).toNumber()] = _5_v0;
          _nw6[(_dafny.ONE).toNumber()] = _5_v0;
          _nw6[(new BigNumber(2)).toNumber()] = _5_v0;
          _nw6[(new BigNumber(3)).toNumber()] = _5_v0;
          _nw6[(new BigNumber(4)).toNumber()] = _5_v0;
          _nw6[(new BigNumber(5)).toNumber()] = _5_v0;
          _nw6[(new BigNumber(6)).toNumber()] = _5_v0;
          _nw6[(new BigNumber(7)).toNumber()] = _5_v0;
          _nw6[(new BigNumber(8)).toNumber()] = _5_v0;
          _nw6[(new BigNumber(9)).toNumber()] = _5_v0;
          _nw6[(new BigNumber(10)).toNumber()] = _5_v0;
          _nw6[(new BigNumber(11)).toNumber()] = _5_v0;
          _32_v22 = _nw6;
          let _index0 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_32_v22).length));
          (_32_v22)[_index0] = _5_v0;
          let _index1 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_32_v22).length));
          let _rhs0 = _19_v13;
          let _rhs1 = _8_v3;
          let _rhs2 = _5_v0;
          let _rhs3 = _29_v19.f26;
          let _rhs4 = new BigNumber(850);
          let _lhs0 = _18_globalState;
          let _lhs1 = _32_v22;
          let _lhs2 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_32_v22).length));
          let _lhs3 = _29_v19;
          let _lhs4 = _18_globalState;
          _19_v13 = _rhs0;
          _lhs0.f19 = _rhs1;
          _lhs1[_lhs2] = _rhs2;
          _lhs3.f26 = _rhs3;
          _lhs4.f19 = _rhs4;
          _19_v13 = _19_v13;
          let _33_v23;
          _33_v23 = _module.D1.create_DC3(_8_v3);
          _33_v23 = _33_v23;
        }
        let _34_v24;
        let _nw7 = new _module.C0();
        _nw7.__ctor(_6_v1);
        _34_v24 = _nw7;
        let _35_v25;
        _35_v25 = _dafny.Set.fromElements(_34_v24);
        let _36_v26;
        _36_v26 = _dafny.MultiSet.fromElements(new BigNumber((_35_v25).length), _8_v3, new BigNumber(721), _8_v3);
        _6_v1 = (_8_v3).isLessThanOrEqualTo(new BigNumber((_36_v26).cardinality()));
      } else {
        _5_v0 = _5_v0;
        _12_v7 = _12_v7;
        let _37_v27;
        _37_v27 = _dafny.Seq.UnicodeFromString("u");
        let _38_v28;
        _38_v28 = _dafny.Set.fromElements(_6_v1);
        (_18_globalState).f15 = ((new BigNumber((_37_v27).length)).multipliedBy(_8_v3)).minus(((_6_v1) ? (_8_v3) : (new BigNumber((_38_v28).length))));
        let _39_v29;
        let _init0 = ((_40_v3, _41_v2) => function (_42_i1) {
          return _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(123), _40_v3, _40_v3, _dafny.Map.Empty.slice().updateUnsafe(_40_v3,new BigNumber((_dafny.MultiSet.FromArray(_41_v2)).cardinality()))));
        })(_8_v3, _7_v2);
        let _nw8 = Array((new BigNumber(3)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw8.length); _i0_0++) {
          _nw8[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _39_v29 = _nw8;
        let _43_v30;
        _43_v30 = _module.D2.create_DC4(_39_v29);
        _39_v29 = (_43_v30).dtor_cf7;
        let _44_v31;
        let _nw9 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
        _44_v31 = _nw9;
        let _45_v32;
        _45_v32 = _dafny.Map.Empty.slice().updateUnsafe(_6_v1,_6_v1);
        let _index2 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_44_v31).length));
        (_44_v31)[_index2] = _45_v32;
        let _46_v33;
        _46_v33 = _dafny.MultiSet.fromElements(_37_v27);
        let _index3 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_44_v31).length));
        let _rhs5 = new BigNumber(((_46_v33).Difference(_46_v33)).cardinality());
        let _rhs6 = _37_v27;
        let _rhs7 = _45_v32;
        let _lhs5 = _18_globalState;
        let _lhs6 = _44_v31;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_44_v31).length));
        _lhs5.f19 = _rhs5;
        _37_v27 = _rhs6;
        _lhs6[_lhs7] = _rhs7;
      }
      let _47_v34;
      let _nw10 = new _module.C0();
      _nw10.__ctor(_6_v1);
      _47_v34 = _nw10;
      let _48_i2;
      _48_i2 = _dafny.ZERO;
      L0: {
        while (_47_v34.f26) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_48_i2)) {
              break L0;
            }
            _48_i2 = (_48_i2).plus(_dafny.ONE);
            let _49_v35;
            _49_v35 = _dafny.Map.Empty.slice().updateUnsafe(_6_v1,false);
            let _50_v36;
            _50_v36 = _dafny.Map.Empty.slice().updateUnsafe(_6_v1,_49_v35);
            let _51_v37;
            _51_v37 = _dafny.Map.Empty.slice().updateUnsafe(_50_v36,_6_v1);
            let _52_v38;
            _52_v38 = _dafny.Seq.of(_8_v3);
            (_47_v34).f26 = (((_51_v37).contains(_50_v36)) ? ((_51_v37).get(_50_v36)) : (_module.__default.fm0(_52_v38, _8_v3, _8_v3, _18_globalState)));
            let _53_v39;
            let _nw11 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
            _53_v39 = _nw11;
            let _54_v40;
            _54_v40 = _module.D2.create_DC4(((_6_v1) ? (_53_v39) : (_53_v39)));
            _54_v40 = _54_v40;
            let _55_v41;
            _55_v41 = _dafny.Set.fromElements(_6_v1);
            let _56_v42;
            let _nw12 = Array((new BigNumber(4)).toNumber());
            _nw12[(_dafny.ZERO).toNumber()] = (_55_v41).Intersect(_55_v41);
            _nw12[(_dafny.ONE).toNumber()] = (_dafny.Set.fromElements(_6_v1)).Intersect(_55_v41);
            _nw12[(new BigNumber(2)).toNumber()] = _55_v41;
            _nw12[(new BigNumber(3)).toNumber()] = _55_v41;
            _56_v42 = _nw12;
            let _index4 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_56_v42).length));
            (_56_v42)[_index4] = _55_v41;
            let _index5 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_56_v42).length));
            (_56_v42)[_index5] = (_55_v41).Intersect(_dafny.Set.fromElements(_47_v34.f26, _47_v34.f26));
            if (_6_v1) {
              _49_v35 = (_49_v35).update(_47_v34.f26, _6_v1);
              let _57_v43;
              _57_v43 = _dafny.Map.Empty.slice().updateUnsafe(_11_v6,_8_v3);
              let _58_v44;
              _58_v44 = _dafny.Map.Empty.slice().updateUnsafe(_8_v3,new BigNumber((_52_v38).length));
              _57_v43 = (_57_v43).update(_11_v6, _module.__default.safeModuloInt((((_58_v44).contains(new BigNumber((_52_v38).length))) ? ((_58_v44).get(new BigNumber((_52_v38).length))) : (new BigNumber(999))), _8_v3));
              let _59_v45;
              let _60_v46;
              let _out2;
              let _out3;
              let _outcollector1 = (_47_v34).m0(_47_v34.f26, _8_v3, _18_globalState);
              _out2 = _outcollector1[0];
              _out3 = _outcollector1[1];
              _59_v45 = _out2;
              _60_v46 = _out3;
              let _61_v47;
              let _nw13 = new _module.C0();
              _nw13.__ctor(true);
              _61_v47 = _nw13;
              let _62_v48;
              _62_v48 = _module.D1.create_DC3(new BigNumber(((_9_v4).Intersect(_9_v4)).length));
              let _63_v49;
              let _nw14 = Array((new BigNumber(15)).toNumber()).fill([]);
              _63_v49 = _nw14;
              let _index6 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_63_v49).length));
              (_63_v49)[_index6] = _5_v0;
              let _64_v50;
              _64_v50 = _dafny.Seq.of(_11_v6);
              let _65_v51;
              _65_v51 = _dafny.Seq.UnicodeFromString("qfhvebs");
              let _66_v52;
              _66_v52 = _dafny.Map.Empty.slice().updateUnsafe(_47_v34.f26,_65_v51);
              let _index7 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_63_v49).length));
              let _rhs8 = _61_v47;
              let _rhs9 = (_module.D2.create_DC5(_5_v0, _64_v50, new BigNumber((_66_v52).length))).dtor_cf10;
              let _rhs10 = _module.D1.create_DC3(_8_v3);
              let _rhs11 = _47_v34.f26;
              let _rhs12 = _5_v0;
              let _lhs8 = _47_v34;
              let _lhs9 = _63_v49;
              let _lhs10 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_63_v49).length));
              _61_v47 = _rhs8;
              _8_v3 = _rhs9;
              _62_v48 = _rhs10;
              _lhs8.f26 = _rhs11;
              _lhs9[_lhs10] = _rhs12;
              let _67_v53;
              let _nw15 = Array((new BigNumber(14)).toNumber());
              _nw15[(_dafny.ZERO).toNumber()] = _module.__default.fm6(_18_globalState);
              _nw15[(_dafny.ONE).toNumber()] = _12_v7;
              _nw15[(new BigNumber(2)).toNumber()] = _12_v7;
              _nw15[(new BigNumber(3)).toNumber()] = _12_v7;
              _nw15[(new BigNumber(4)).toNumber()] = _12_v7;
              _nw15[(new BigNumber(5)).toNumber()] = _12_v7;
              _nw15[(new BigNumber(6)).toNumber()] = _12_v7;
              _nw15[(new BigNumber(7)).toNumber()] = _12_v7;
              _nw15[(new BigNumber(8)).toNumber()] = _12_v7;
              _nw15[(new BigNumber(9)).toNumber()] = _12_v7;
              _nw15[(new BigNumber(10)).toNumber()] = _12_v7;
              _nw15[(new BigNumber(11)).toNumber()] = _12_v7;
              _nw15[(new BigNumber(12)).toNumber()] = _12_v7;
              _nw15[(new BigNumber(13)).toNumber()] = _12_v7;
              _67_v53 = _nw15;
              let _index8 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_67_v53).length));
              (_67_v53)[_index8] = _12_v7;
              let _index9 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_67_v53).length));
              (_67_v53)[_index9] = _12_v7;
            } else {
              let _nw16 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
              (_18_globalState).f5 = _nw16;
              (_18_globalState).f15 = _8_v3;
              let _68_v54;
              _68_v54 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC6(),_12_v7);
              let _69_v55;
              _69_v55 = _dafny.Map.Empty.slice().updateUnsafe(_8_v3,_68_v54);
              (_18_globalState).f19 = ((new BigNumber((_69_v55).length)).minus(_8_v3)).multipliedBy(_8_v3);
              _7_v2 = _dafny.Seq.update(_7_v2, _module.__default.safeIndex(_8_v3, new BigNumber((_7_v2).length)), _47_v34.f26);
              _12_v7 = _12_v7;
            }
          }
        }
      }
      let _70_v56;
      let _init1 = function (_71_i3) {
        return _module.D2.create_DC6();
      };
      let _nw17 = Array((new BigNumber(7)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw17.length); _i0_1++) {
        _nw17[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _70_v56 = _nw17;
      let _72_v57;
      let _nw18 = Array((new BigNumber(25)).toNumber());
      _nw18[(_dafny.ZERO).toNumber()] = _70_v56;
      _nw18[(_dafny.ONE).toNumber()] = _70_v56;
      _nw18[(new BigNumber(2)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(3)).toNumber()] = ((_47_v34.f26) ? (_70_v56) : (_70_v56));
      _nw18[(new BigNumber(4)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(5)).toNumber()] = ((!(_47_v34.f26)) ? (_70_v56) : (_70_v56));
      _nw18[(new BigNumber(6)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(7)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(8)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(9)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(10)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(11)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(12)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(13)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(14)).toNumber()] = ((_6_v1) ? (_70_v56) : (_70_v56));
      _nw18[(new BigNumber(15)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(16)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(17)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(18)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(19)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(20)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(21)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(22)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(23)).toNumber()] = _70_v56;
      _nw18[(new BigNumber(24)).toNumber()] = _70_v56;
      _72_v57 = _nw18;
      _72_v57 = ((_47_v34.f26) ? (_72_v57) : (_72_v57));
      let _73_v58;
      _73_v58 = _dafny.Seq.UnicodeFromString("xseuq");
      let _index10 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
      (_5_v0)[_index10] = new BigNumber((_73_v58).length);
      let _index11 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
      (_5_v0)[_index11] = (_dafny.ZERO).minus(_8_v3);
      let _hi0 = _8_v3;
      for (let _74_i4 = new BigNumber(210); _74_i4.isLessThan(_hi0); _74_i4 = _74_i4.plus(_dafny.ONE)) {
        let _75_v59;
        let _76_v60;
        let _out4;
        let _out5;
        let _outcollector2 = (_47_v34).m0(_47_v34.f26, ((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]).multipliedBy(_8_v3), _18_globalState);
        _out4 = _outcollector2[0];
        _out5 = _outcollector2[1];
        _75_v59 = _out4;
        _76_v60 = _out5;
        let _77_v61;
        _77_v61 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),_9_v4);
        _77_v61 = (_77_v61).update(_12_v7, _9_v4);
        let _78_v62;
        _78_v62 = _dafny.Seq.of(_73_v58);
        if (!_dafny.areEqual(_73_v58, _dafny.Seq.Concat((_78_v62)[_module.__default.safeIndex((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))], new BigNumber((_78_v62).length))], _73_v58))) {
          (_18_globalState).f5 = _5_v0;
          let _79_v63;
          let _nw19 = Array((new BigNumber(8)).toNumber()).fill([]);
          _79_v63 = _nw19;
          let _80_v64;
          _80_v64 = _dafny.Map.Empty.slice().updateUnsafe(_79_v63,_5_v0);
          _80_v64 = (_80_v64).update(_79_v63, _5_v0);
          let _81_v65;
          _81_v65 = _dafny.Seq.of(_11_v6);
          let _82_v66;
          _82_v66 = _module.D2.create_DC5(_5_v0, _81_v65, (_dafny.ZERO).minus((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]));
          let _pat_let_tv0 = _5_v0;
          let _83_v67;
          _83_v67 = _dafny.Map.Empty.slice().updateUnsafe(_47_v34.f26,function (_pat_let0_0) {
            return function (_84_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_85_dt__update_hcf10_h0) {
                  return function (_pat_let2_0) {
                    return function (_86_dt__update_hcf8_h0) {
                      return _module.D2.create_DC5(_86_dt__update_hcf8_h0, (_84_dt__update__tmp_h0).dtor_cf9, _85_dt__update_hcf10_h0);
                    }(_pat_let2_0);
                  }(_pat_let_tv0);
                }(_pat_let1_0);
              }(new BigNumber(762));
            }(_pat_let0_0);
          }(_82_v66));
          _83_v67 = _83_v67;
          let _87_v68;
          let _88_v69;
          let _out6;
          let _out7;
          let _outcollector3 = (_47_v34).m0((new BigNumber(826)).isLessThan(_74_i4), new BigNumber(143), _18_globalState);
          _out6 = _outcollector3[0];
          _out7 = _outcollector3[1];
          _87_v68 = _out6;
          _88_v69 = _out7;
          (_47_v34).f26 = _47_v34.f26;
        } else {
          let _89_v70;
          _89_v70 = _dafny.Map.Empty.slice().updateUnsafe((_47_v34.f26) || (_6_v1),_dafny.Seq.IsProperPrefixOf(_73_v58, _73_v58));
          _89_v70 = (_89_v70).update(_47_v34.f26, !(_75_v59) || (_47_v34.f26));
          _6_v1 = !(_6_v1) || (_47_v34.f26);
          let _90_v71;
          let _91_v72;
          let _out8;
          let _out9;
          let _outcollector4 = (_47_v34).m0(_75_v59, new BigNumber((_module.__default.fm7(_18_globalState)).length), _18_globalState);
          _out8 = _outcollector4[0];
          _out9 = _outcollector4[1];
          _90_v71 = _out8;
          _91_v72 = _out9;
          let _92_v73;
          let _nw20 = new _module.C0();
          _nw20.__ctor(_90_v71);
          _92_v73 = _nw20;
          _92_v73 = _47_v34;
          let _93_v74;
          let _94_v75;
          let _out10;
          let _out11;
          let _outcollector5 = (_92_v73).m0(true, _74_i4, _18_globalState);
          _out10 = _outcollector5[0];
          _out11 = _outcollector5[1];
          _93_v74 = _out10;
          _94_v75 = _out11;
        }
        let _95_v76;
        _95_v76 = _dafny.Seq.of(_7_v2, _7_v2);
        let _96_v77;
        _96_v77 = _module.D3.create_DC7(_7_v2);
        let _97_v78;
        _97_v78 = _dafny.MultiSet.fromElements((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]);
        let _98_v79;
        _98_v79 = _dafny.Map.Empty.slice().updateUnsafe(_97_v78,_dafny.Seq.of(_6_v1));
        let _99_v80;
        let _nw21 = Array((new BigNumber(28)).toNumber());
        _nw21[(_dafny.ZERO).toNumber()] = _7_v2;
        _nw21[(_dafny.ONE).toNumber()] = _7_v2;
        _nw21[(new BigNumber(2)).toNumber()] = _7_v2;
        _nw21[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_7_v2, _7_v2);
        _nw21[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_47_v34.f26, _6_v1, true, _75_v59, !(_47_v34.f26));
        _nw21[(new BigNumber(5)).toNumber()] = _7_v2;
        _nw21[(new BigNumber(6)).toNumber()] = ((true) ? (_7_v2) : (_7_v2));
        _nw21[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_7_v2, (_95_v76)[_module.__default.safeIndex(_8_v3, new BigNumber((_95_v76).length))]);
        _nw21[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_7_v2, _7_v2);
        _nw21[(new BigNumber(9)).toNumber()] = _7_v2;
        _nw21[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_47_v34.f26), _dafny.Seq.of(_47_v34.f26));
        _nw21[(new BigNumber(11)).toNumber()] = _7_v2;
        _nw21[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(_6_v1, _47_v34.f26), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(574)), ((_100_v0) => function (_101_i5) {
          return (_100_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_100_v0).length))];
        })(_5_v0))).length), new BigNumber((_dafny.Seq.of(_6_v1, _47_v34.f26)).length)), _47_v34.f26);
        _nw21[(new BigNumber(13)).toNumber()] = (_96_v77).dtor_cf11;
        _nw21[(new BigNumber(14)).toNumber()] = _7_v2;
        _nw21[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_7_v2, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((_47_v34).fm1(_75_v59, _18_globalState)).length)), new BigNumber((_7_v2).length)), _47_v34.f26);
        _nw21[(new BigNumber(16)).toNumber()] = _7_v2;
        _nw21[(new BigNumber(17)).toNumber()] = _7_v2;
        _nw21[(new BigNumber(18)).toNumber()] = ((_6_v1) ? (_7_v2) : (_7_v2));
        _nw21[(new BigNumber(19)).toNumber()] = _7_v2;
        _nw21[(new BigNumber(20)).toNumber()] = _7_v2;
        _nw21[(new BigNumber(21)).toNumber()] = (_95_v76)[_module.__default.safeIndex(new BigNumber(-260), new BigNumber((_95_v76).length))];
        _nw21[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_7_v2, _7_v2);
        _nw21[(new BigNumber(23)).toNumber()] = _7_v2;
        _nw21[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_7_v2, _7_v2);
        _nw21[(new BigNumber(25)).toNumber()] = (((_98_v79).contains(_dafny.MultiSet.fromElements(_8_v3))) ? ((_98_v79).get(_dafny.MultiSet.fromElements(_8_v3))) : (_dafny.Seq.of(_47_v34.f26, _6_v1)));
        _nw21[(new BigNumber(26)).toNumber()] = _7_v2;
        _nw21[(new BigNumber(27)).toNumber()] = _7_v2;
        _99_v80 = _nw21;
        let _index12 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_99_v80).length));
        (_99_v80)[_index12] = _dafny.Seq.Concat(_dafny.Seq.update(_7_v2, _module.__default.safeIndex(_8_v3, new BigNumber((_7_v2).length)), _47_v34.f26), _7_v2);
        let _index13 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_99_v80).length));
        (_99_v80)[_index13] = _dafny.Seq.of(!(false), _6_v1, _47_v34.f26);
      }
      if (_6_v1) {
        if ((_10_v5)) {
          (_47_v34).f26 = _47_v34.f26;
          let _102_v81;
          let _103_v82;
          let _out12;
          let _out13;
          let _outcollector6 = (_47_v34).m0(_47_v34.f26, (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))], _18_globalState);
          _out12 = _outcollector6[0];
          _out13 = _outcollector6[1];
          _102_v81 = _out12;
          _103_v82 = _out13;
          (_18_globalState).f15 = (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))];
          (_18_globalState).f15 = new BigNumber(398);
          (_47_v34).f26 = _102_v81;
        } else {
          let _104_v83;
          let _105_v84;
          let _out14;
          let _out15;
          let _outcollector7 = (_47_v34).m0(_47_v34.f26, _8_v3, _18_globalState);
          _out14 = _outcollector7[0];
          _out15 = _outcollector7[1];
          _104_v83 = _out14;
          _105_v84 = _out15;
          let _index14 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_11_v6).length));
          (_11_v6)[_index14] = !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("aqoumw"), _12_v7);
          let _106_v85;
          let _nw22 = Array((new BigNumber(5)).toNumber());
          _nw22[(_dafny.ZERO).toNumber()] = _12_v7;
          _nw22[(_dafny.ONE).toNumber()] = _12_v7;
          _nw22[(new BigNumber(2)).toNumber()] = _12_v7;
          _nw22[(new BigNumber(3)).toNumber()] = _12_v7;
          _nw22[(new BigNumber(4)).toNumber()] = _12_v7;
          _106_v85 = _nw22;
          let _107_v86;
          _107_v86 = _dafny.Map.Empty.slice().updateUnsafe(_12_v7,_106_v85);
          let _108_v87;
          let _nw23 = Array((new BigNumber(8)).toNumber());
          _nw23[(_dafny.ZERO).toNumber()] = ((_47_v34.f26) ? (_106_v85) : (_106_v85));
          _nw23[(_dafny.ONE).toNumber()] = _106_v85;
          _nw23[(new BigNumber(2)).toNumber()] = _106_v85;
          _nw23[(new BigNumber(3)).toNumber()] = _106_v85;
          _nw23[(new BigNumber(4)).toNumber()] = _106_v85;
          _nw23[(new BigNumber(5)).toNumber()] = (((_107_v86).contains(_12_v7)) ? ((_107_v86).get(_12_v7)) : (_106_v85));
          _nw23[(new BigNumber(6)).toNumber()] = _106_v85;
          _nw23[(new BigNumber(7)).toNumber()] = _106_v85;
          _108_v87 = _nw23;
          let _109_v88;
          _109_v88 = _module.D3.create_DC9((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]);
          let _110_v89;
          _110_v89 = _module.D3.create_DC10(_109_v88);
          let _index15 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_11_v6).length));
          let _index16 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
          let _rhs13 = (_10_v5);
          let _rhs14 = _108_v87;
          let _rhs15 = ((!(_47_v34.f26)) ? (_110_v89) : (_110_v89));
          let _rhs16 = (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))];
          let _lhs11 = _11_v6;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_11_v6).length));
          let _lhs13 = _5_v0;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
          _lhs11[_lhs12] = _rhs13;
          _108_v87 = _rhs14;
          _110_v89 = _rhs15;
          _lhs13[_lhs14] = _rhs16;
          let _index17 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_11_v6).length));
          let _rhs17 = (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))];
          let _rhs18 = _module.__default.fm0(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-879)), ((_111_v3) => function (_112_i6) {
            return _111_v3;
          })(_8_v3)), _module.__default.safeDivisionInt((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))], (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]), (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))], _18_globalState);
          let _lhs15 = _18_globalState;
          let _lhs16 = _11_v6;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_11_v6).length));
          _lhs15.f19 = _rhs17;
          _lhs16[_lhs17] = _rhs18;
          let _index18 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
          (_5_v0)[_index18] = (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))];
          let _113_v90;
          _113_v90 = _dafny.Seq.of(_47_v34, _47_v34);
          let _114_v91;
          _114_v91 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_113_v90).length)));
          let _115_v92;
          _115_v92 = _dafny.Map.Empty.slice().updateUnsafe(_47_v34.f26,(_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]);
          let _116_v93;
          _116_v93 = _dafny.Map.Empty.slice().updateUnsafe(_8_v3,_115_v92);
          let _117_v94;
          _117_v94 = _dafny.MultiSet.fromElements(_module.__default.fm0(_114_v91, (_dafny.ZERO).minus(_8_v3), _8_v3, _18_globalState), !((((_116_v93).contains((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))])) ? ((_116_v93).get((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))])) : (_115_v92))).contains(_47_v34.f26));
          let _118_v95;
          _118_v95 = _dafny.Map.Empty.slice().updateUnsafe((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))],(_11_v6)[_module.__default.safeIndex(new BigNumber(320), new BigNumber((_11_v6).length))]);
          let _119_v96;
          _119_v96 = _dafny.Map.Empty.slice().updateUnsafe(_12_v7,_dafny.MultiSet.fromElements((_11_v6)[_module.__default.safeIndex(new BigNumber(320), new BigNumber((_11_v6).length))], (_11_v6)[_module.__default.safeIndex(new BigNumber(320), new BigNumber((_11_v6).length))]));
          let _index19 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
          let _rhs19 = (_module.__default.fm3(new BigNumber(82), _118_v95, (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))], _18_globalState)).plus(_module.__default.safeDivisionInt(_8_v3, (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]));
          let _rhs20 = _12_v7;
          let _rhs21 = (_11_v6)[_module.__default.safeIndex(new BigNumber(320), new BigNumber((_11_v6).length))];
          let _rhs22 = (((_119_v96).contains(_12_v7)) ? ((_119_v96).get(_12_v7)) : (_117_v94));
          let _lhs18 = _5_v0;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
          let _lhs20 = _47_v34;
          _lhs18[_lhs19] = _rhs19;
          _12_v7 = _rhs20;
          _lhs20.f26 = _rhs21;
          _117_v94 = _rhs22;
        }
        if (_6_v1) {
          let _index20 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
          (_5_v0)[_index20] = _8_v3;
          let _120_v97;
          _120_v97 = _dafny.Seq.of((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]);
          let _121_v98;
          _121_v98 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.Seq.contains(_120_v97, _8_v3),_47_v34.f26);
          _121_v98 = (_121_v98).update(_47_v34.f26, !(_47_v34.f26));
          let _122_v99;
          _122_v99 = _dafny.Set.fromElements(_47_v34.f26);
          _122_v99 = _122_v99;
          let _123_v100;
          let _nw24 = new _module.C0();
          _nw24.__ctor(false);
          _123_v100 = _nw24;
          let _124_v101;
          _124_v101 = _dafny.Seq.of(_123_v100, _123_v100);
          let _125_v102;
          _125_v102 = _module.D4.create_DC11((_124_v101)[_module.__default.safeIndex((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))], new BigNumber((_124_v101).length))]);
          _123_v100 = (_125_v102).dtor_cf14;
          let _126_v103;
          let _nw25 = Array((new BigNumber(19)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _126_v103 = _nw25;
          let _127_v104;
          _127_v104 = _dafny.Map.Empty.slice().updateUnsafe(_5_v0,_126_v103);
          let _128_v105;
          let _nw26 = Array((new BigNumber(14)).toNumber());
          _nw26[(_dafny.ZERO).toNumber()] = _126_v103;
          _nw26[(_dafny.ONE).toNumber()] = _126_v103;
          _nw26[(new BigNumber(2)).toNumber()] = _126_v103;
          _nw26[(new BigNumber(3)).toNumber()] = _126_v103;
          _nw26[(new BigNumber(4)).toNumber()] = _126_v103;
          _nw26[(new BigNumber(5)).toNumber()] = _126_v103;
          _nw26[(new BigNumber(6)).toNumber()] = _126_v103;
          _nw26[(new BigNumber(7)).toNumber()] = _126_v103;
          _nw26[(new BigNumber(8)).toNumber()] = _126_v103;
          _nw26[(new BigNumber(9)).toNumber()] = (((_127_v104).contains(_5_v0)) ? ((_127_v104).get(_5_v0)) : (_126_v103));
          _nw26[(new BigNumber(10)).toNumber()] = _126_v103;
          _nw26[(new BigNumber(11)).toNumber()] = _126_v103;
          _nw26[(new BigNumber(12)).toNumber()] = _126_v103;
          _nw26[(new BigNumber(13)).toNumber()] = _126_v103;
          _128_v105 = _nw26;
          let _129_v106;
          _129_v106 = _dafny.Map.Empty.slice().updateUnsafe(_7_v2,_128_v105);
          _129_v106 = (_129_v106).update(_dafny.Seq.Concat(_7_v2, _7_v2), _128_v105);
        } else {
          let _130_v107;
          let _nw27 = new _module.C0();
          _nw27.__ctor(true);
          _130_v107 = _nw27;
          (_47_v34).f26 = true;
          _12_v7 = _12_v7;
          let _131_v108;
          _131_v108 = _dafny.Map.Empty.slice().updateUnsafe(_130_v107,_12_v7);
          let _132_v109;
          _132_v109 = _dafny.Seq.of(new BigNumber((_131_v108).length), new BigNumber(254));
          let _133_v110;
          _133_v110 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_47_v34.f26,_47_v34.f26)).length)),_47_v34.f26);
          let _134_v111;
          _134_v111 = _dafny.Map.Empty.slice().updateUnsafe(_133_v110,new BigNumber((_73_v58).length));
          let _rhs23 = _132_v109;
          let _rhs24 = (new BigNumber((_134_v111).length)).plus(((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]).minus(new BigNumber((_dafny.MultiSet.fromElements(_130_v107.f26)).cardinality())));
          let _lhs21 = _18_globalState;
          _132_v109 = _rhs23;
          _lhs21.f19 = _rhs24;
          let _135_v112;
          _135_v112 = _dafny.Seq.of(_73_v58);
          let _136_v113;
          _136_v113 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("d"));
          let _137_v114;
          let _nw28 = Array((new BigNumber(7)).toNumber());
          _nw28[(_dafny.ZERO).toNumber()] = (_dafny.MultiSet.fromElements(_73_v58, _73_v58, _dafny.Seq.UnicodeFromString("cbpcgr"))).Intersect(_dafny.MultiSet.FromArray(_135_v112));
          _nw28[(_dafny.ONE).toNumber()] = _136_v113;
          _nw28[(new BigNumber(2)).toNumber()] = (_136_v113).Union(_module.__default.fm8(_18_globalState));
          _nw28[(new BigNumber(3)).toNumber()] = (_dafny.MultiSet.fromElements(_73_v58)).update(_dafny.Seq.UnicodeFromString("ihs"), _module.__default.abs(_8_v3));
          _nw28[(new BigNumber(4)).toNumber()] = _module.__default.fm8(_18_globalState);
          _nw28[(new BigNumber(5)).toNumber()] = _136_v113;
          _nw28[(new BigNumber(6)).toNumber()] = (_dafny.MultiSet.fromElements(_73_v58, _73_v58)).Union(_136_v113);
          _137_v114 = _nw28;
          let _index21 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_137_v114).length));
          (_137_v114)[_index21] = _136_v113;
          let _index22 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_137_v114).length));
          (_137_v114)[_index22] = _136_v113;
        }
        _11_v6 = _11_v6;
        (_18_globalState).f15 = (new BigNumber(906)).plus(new BigNumber(758));
        let _138_v115;
        _138_v115 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC5(_5_v0, _dafny.Seq.of(_11_v6, _11_v6, _11_v6, _11_v6), new BigNumber((_7_v2).length)),_73_v58);
        _138_v115 = _138_v115;
      } else {
        let _139_v116;
        _139_v116 = _dafny.Seq.of(_11_v6, _11_v6, _11_v6, _11_v6);
        let _140_v117;
        _140_v117 = _module.D2.create_DC5(_5_v0, _139_v116, (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]);
        let _141_v118;
        _141_v118 = _dafny.Seq.of(_139_v116);
        _140_v117 = _module.D2.create_DC5(_5_v0, (_141_v118)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_5_v0, _5_v0, _5_v0)).length), new BigNumber((_141_v118).length))], (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]);
        let _142_v119;
        _142_v119 = _dafny.Seq.of(new BigNumber(195));
        let _143_v120;
        _143_v120 = _dafny.Map.Empty.slice().updateUnsafe((_7_v2)[_module.__default.safeIndex((_142_v119)[_module.__default.safeIndex(_8_v3, new BigNumber((_142_v119).length))], new BigNumber((_7_v2).length))],_8_v3);
        let _144_v121;
        _144_v121 = _dafny.MultiSet.fromElements((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]);
        let _145_v122;
        _145_v122 = _dafny.Map.Empty.slice().updateUnsafe((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))],new BigNumber((_144_v121).cardinality()));
        let _146_v123;
        _146_v123 = _dafny.MultiSet.fromElements(_145_v122);
        _143_v120 = (_143_v120).update(_47_v34.f26, new BigNumber((_146_v123).cardinality()));
        let _147_v124;
        _147_v124 = _dafny.Map.Empty.slice().updateUnsafe(!(_6_v1),_47_v34.f26);
        _147_v124 = (_147_v124).update(_47_v34.f26, (_7_v2)[_module.__default.safeIndex(_8_v3, new BigNumber((_7_v2).length))]);
        let _148_v125;
        let _nw29 = new _module.C0();
        _nw29.__ctor((_47_v34.f26) === (_6_v1));
        _148_v125 = _nw29;
        let _149_v126;
        _149_v126 = _dafny.Set.fromElements(true);
        _149_v126 = (_149_v126).Union(_149_v126);
      }
      let _150_v127;
      _150_v127 = _dafny.Map.Empty.slice().updateUnsafe(_5_v0,new BigNumber((_7_v2).length));
      (_47_v34).f26 = !(_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_6_v1)).length), (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))])).isEqualTo(((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]).minus((_dafny.ZERO).minus(new BigNumber((_150_v127).length))));
      (_18_globalState).f19 = (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))];
      let _151_v128;
      _151_v128 = _dafny.MultiSet.fromElements(_47_v34.f26, _6_v1);
      let _152_v129;
      _152_v129 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("stmah"), _73_v58)).length),!((_151_v128).IsProperSubsetOf(_151_v128)));
      _152_v129 = (_152_v129).update(_8_v3, _6_v1);
      let _153_v130;
      let _nw30 = Array((new BigNumber(11)).toNumber()).fill(_dafny.MultiSet.Empty);
      _153_v130 = _nw30;
      let _154_v131;
      _154_v131 = _module.D4.create_DC12(_153_v130, _11_v6, _6_v1);
      let _155_v132;
      _155_v132 = _module.D4.create_DC13(_154_v131);
      _155_v132 = _module.D4.create_DC13(_154_v131);
      let _156_v133;
      let _nw31 = new _module.C0();
      _nw31.__ctor(_6_v1);
      _156_v133 = _nw31;
      let _index23 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
      let _rhs25 = new BigNumber(872);
      let _rhs26 = _dafny.Seq.Concat(_7_v2, _7_v2);
      let _rhs27 = ((_8_v3).multipliedBy(new BigNumber(165))).multipliedBy((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]);
      let _rhs28 = _dafny.Seq.of(_6_v1, _156_v133.f26, _47_v34.f26);
      let _lhs22 = _5_v0;
      let _lhs23 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
      let _lhs24 = _18_globalState;
      _lhs22[_lhs23] = _rhs25;
      _7_v2 = _rhs26;
      _lhs24.f19 = _rhs27;
      _7_v2 = _rhs28;
      if (_47_v34.f26) {
        (_47_v34).f26 = _47_v34.f26;
        let _157_v134;
        let _nw32 = new _module.C0();
        _nw32.__ctor(_47_v34.f26);
        _157_v134 = _nw32;
        let _index24 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
        (_5_v0)[_index24] = _module.__default.safeModuloInt((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))], _8_v3);
        let _158_v135;
        _158_v135 = _dafny.Map.Empty.slice().updateUnsafe(_47_v34.f26,_157_v134.f26);
        let _159_v136;
        let _160_v137;
        let _out16;
        let _out17;
        let _outcollector8 = (_157_v134).m0((_6_v1) === ((_7_v2)[_module.__default.safeIndex(_8_v3, new BigNumber((_7_v2).length))]), new BigNumber((_158_v135).length), _18_globalState);
        _out16 = _outcollector8[0];
        _out17 = _outcollector8[1];
        _159_v136 = _out16;
        _160_v137 = _out17;
        _8_v3 = _8_v3;
      } else {
        let _161_v138;
        let _162_v139;
        let _out18;
        let _out19;
        let _outcollector9 = (_156_v133).m0(_47_v34.f26, _8_v3, _18_globalState);
        _out18 = _outcollector9[0];
        _out19 = _outcollector9[1];
        _161_v138 = _out18;
        _162_v139 = _out19;
        let _index25 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
        (_5_v0)[_index25] = _8_v3;
        _156_v133 = _156_v133;
        let _163_v140;
        _163_v140 = _module.D5.create_DC14(_152_v129);
        let _164_v141;
        _164_v141 = _dafny.MultiSet.fromElements(_73_v58);
        let _165_v142;
        _165_v142 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_73_v58, _module.__default.safeIndex(_module.__default.fm3(_8_v3, (_163_v140).dtor_cf19, _8_v3, _18_globalState), new BigNumber((_73_v58).length)), _12_v7),new BigNumber((_164_v141).cardinality()));
        _165_v142 = (_165_v142).update(_73_v58, (_8_v3).plus(_8_v3));
        if ((!(_47_v34.f26)) === (!(_156_v133.f26))) {
          let _index26 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length));
          (_5_v0)[_index26] = (_dafny.ZERO).minus(_8_v3);
          (_18_globalState).f5 = _5_v0;
          _12_v7 = (((_15_v10).contains(_47_v34.f26)) ? ((_15_v10).get(_47_v34.f26)) : (_12_v7));
          let _166_v143;
          _166_v143 = _dafny.Map.Empty.slice().updateUnsafe(_12_v7,_5_v0);
          let _167_v144;
          let _nw33 = Array((new BigNumber(13)).toNumber());
          _nw33[(_dafny.ZERO).toNumber()] = (((_166_v143).contains(_12_v7)) ? ((_166_v143).get(_12_v7)) : (_5_v0));
          _nw33[(_dafny.ONE).toNumber()] = _5_v0;
          _nw33[(new BigNumber(2)).toNumber()] = _5_v0;
          _nw33[(new BigNumber(3)).toNumber()] = _5_v0;
          _nw33[(new BigNumber(4)).toNumber()] = _5_v0;
          _nw33[(new BigNumber(5)).toNumber()] = _5_v0;
          _nw33[(new BigNumber(6)).toNumber()] = _5_v0;
          _nw33[(new BigNumber(7)).toNumber()] = _5_v0;
          _nw33[(new BigNumber(8)).toNumber()] = _5_v0;
          _nw33[(new BigNumber(9)).toNumber()] = _5_v0;
          _nw33[(new BigNumber(10)).toNumber()] = _5_v0;
          _nw33[(new BigNumber(11)).toNumber()] = _5_v0;
          _nw33[(new BigNumber(12)).toNumber()] = _5_v0;
          _167_v144 = _nw33;
          _167_v144 = _167_v144;
          let _168_v145;
          let _169_v146;
          let _out20;
          let _out21;
          let _outcollector10 = (_47_v34).m0(_47_v34.f26, _module.__default.safeModuloInt(_8_v3, _8_v3), _18_globalState);
          _out20 = _outcollector10[0];
          _out21 = _outcollector10[1];
          _168_v145 = _out20;
          _169_v146 = _out21;
        } else {
          _161_v138 = ((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]).isLessThan((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]);
          let _170_v147;
          _170_v147 = _module.D2.create_DC6();
          _170_v147 = _170_v147;
          (_156_v133).f26 = _47_v34.f26;
          let _171_v148;
          let _172_v149;
          let _out22;
          let _out23;
          let _outcollector11 = (_47_v34).m0(_47_v34.f26, _8_v3, _18_globalState);
          _out22 = _outcollector11[0];
          _out23 = _outcollector11[1];
          _171_v148 = _out22;
          _172_v149 = _out23;
          let _173_v150;
          let _init2 = ((_174_v58) => function (_175_i7) {
            return _174_v58;
          })(_73_v58);
          let _nw34 = Array((new BigNumber(15)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw34.length); _i0_2++) {
            _nw34[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _173_v150 = _nw34;
          let _index27 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((_173_v150).length));
          (_173_v150)[_index27] = _73_v58;
          let _index28 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((_173_v150).length));
          (_173_v150)[_index28] = _dafny.Seq.Concat(((_47_v34.f26) ? (_73_v58) : (_dafny.Seq.UnicodeFromString("faxxuf"))), _73_v58);
        }
      }
      let _ingredients0 = [];
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_14_v9).length))) {
        let _176_i8 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_176_i8)) && ((_176_i8).isLessThan(new BigNumber((_14_v9).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_14_v9, (_176_i8).toNumber(), _dafny.Set.fromElements(((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]).minus((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))]), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(-379)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-899)), function (_177_i9) {
            return new BigNumber(212);
          })), _dafny.Seq.of(_dafny.Seq.of((_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))], _8_v3, (_5_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_5_v0).length))])))).length), new BigNumber(934), (_dafny.ZERO).minus(_8_v3))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      _6_v1 = !(_47_v34.f26);
      process.stdout.write(_dafny.toString((_5_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_6_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_7_v2, _dafny.Seq.of(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_8_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_9_v4).equals(_dafny.Set.fromElements(new BigNumber(861)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_10_v5)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_11_v6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_11_v6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_11_v6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_11_v6)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_11_v6)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_11_v6)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_11_v6)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_11_v6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_11_v6)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_11_v6)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_11_v6)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_11_v6)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_12_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_13_v8).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(861),new _dafny.CodePoint('l'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_14_v9)[_dafny.ZERO]).equals(_dafny.Set.fromElements(_dafny.ZERO, new BigNumber(3), new BigNumber(934), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_14_v9)[_dafny.ONE]).equals(_dafny.Set.fromElements(_dafny.ZERO, new BigNumber(3), new BigNumber(934), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_14_v9)[new BigNumber(2)]).equals(_dafny.Set.fromElements(_dafny.ZERO, new BigNumber(3), new BigNumber(934), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_14_v9)[new BigNumber(3)]).equals(_dafny.Set.fromElements(_dafny.ZERO, new BigNumber(3), new BigNumber(934), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_14_v9)[new BigNumber(4)]).equals(_dafny.Set.fromElements(_dafny.ZERO, new BigNumber(3), new BigNumber(934), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_15_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('l'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_16_v11).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('l'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('l'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_17_v12).dtor_cf1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_18_globalState).f1, _dafny.Seq.of(false, true, false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f3).equals(_dafny.Set.fromElements(new BigNumber(861)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f4)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f4)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f4)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f4)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f4)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f4)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f4)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f4)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f4)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f4)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f4)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState.f5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f6).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(861),new _dafny.CodePoint('l'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_18_globalState).f9)[_dafny.ZERO]).equals(_dafny.Set.fromElements(_dafny.ZERO, new BigNumber(3), new BigNumber(934), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_18_globalState).f9)[_dafny.ONE]).equals(_dafny.Set.fromElements(_dafny.ZERO, new BigNumber(3), new BigNumber(934), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_18_globalState).f9)[new BigNumber(2)]).equals(_dafny.Set.fromElements(_dafny.ZERO, new BigNumber(3), new BigNumber(934), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_18_globalState).f9)[new BigNumber(3)]).equals(_dafny.Set.fromElements(_dafny.ZERO, new BigNumber(3), new BigNumber(934), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_18_globalState).f9)[new BigNumber(4)]).equals(_dafny.Set.fromElements(_dafny.ZERO, new BigNumber(3), new BigNumber(934), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f12)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f12)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f12)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f12)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f12)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f12)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f12)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f12)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f12)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f12)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f12)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f12)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_18_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_18_globalState.f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState.f20).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('l'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('l'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('l'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('l'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f22)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_18_globalState).f23)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_globalState).f25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_47_v34.f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_48_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_73_v58).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_150_v127).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v128).equals(_dafny.MultiSet.fromElements(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_152_v129).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(10),true).updateUnsafe(_dafny.ONE,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v131).dtor_cf16)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v131).dtor_cf16)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v131).dtor_cf16)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v131).dtor_cf16)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v131).dtor_cf16)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v131).dtor_cf16)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v131).dtor_cf16)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v131).dtor_cf16)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v131).dtor_cf16)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v131).dtor_cf16)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v131).dtor_cf16)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v131).dtor_cf16)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_154_v131).dtor_cf17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_155_v132).dtor_cf18).dtor_cf16)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_155_v132).dtor_cf18).dtor_cf16)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_155_v132).dtor_cf18).dtor_cf16)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_155_v132).dtor_cf18).dtor_cf16)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_155_v132).dtor_cf18).dtor_cf16)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_155_v132).dtor_cf18).dtor_cf16)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_155_v132).dtor_cf18).dtor_cf16)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_155_v132).dtor_cf18).dtor_cf16)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_155_v132).dtor_cf18).dtor_cf16)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_155_v132).dtor_cf18).dtor_cf16)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_155_v132).dtor_cf18).dtor_cf16)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_155_v132).dtor_cf18).dtor_cf16)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_155_v132).dtor_cf18).dtor_cf17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_156_v133.f26));
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
    static create_DC2(cf2, cf3, cf4, cf5) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC3(cf6) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(2);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC1() { return this.$tag === 2; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf1 === other.cf1;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.Map.Empty);
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
    static create_DC5(cf8, cf9, cf10) {
      let $dt = new D2(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC6() {
      let $dt = new D2(1);
      return $dt;
    }
    static create_DC4(cf7) {
      let $dt = new D2(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6";
      } else if (this.$tag === 2) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf7 === other.cf7;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC5([], _dafny.Seq.of(), _dafny.ZERO);
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
    static create_DC8() {
      let $dt = new D3(0);
      return $dt;
    }
    static create_DC9(cf12) {
      let $dt = new D3(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC7(cf11) {
      let $dt = new D3(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC10(cf13) {
      let $dt = new D3(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get is_DC10() { return this.$tag === 3; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf13) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8();
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
    static create_DC12(cf15, cf16, cf17) {
      let $dt = new D4(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC11(cf14) {
      let $dt = new D4(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC13(cf18) {
      let $dt = new D4(2);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf15 === other.cf15 && this.cf16 === other.cf16 && this.cf17 === other.cf17;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf14 === other.cf14;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12([], [], false);
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
    static create_DC15(cf20) {
      let $dt = new D5(0);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC14(cf19) {
      let $dt = new D5(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC16(cf21) {
      let $dt = new D5(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC15(null);
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
      this.f5 = [];
      this.f15 = _dafny.ZERO;
      this.f19 = _dafny.ZERO;
      this.f20 = _dafny.MultiSet.Empty;
      this._f0 = [];
      this._f1 = _dafny.Seq.of();
      this._f2 = false;
      this._f3 = _dafny.Set.Empty;
      this._f4 = [];
      this._f6 = _dafny.Map.Empty;
      this._f7 = false;
      this._f8 = false;
      this._f9 = [];
      this._f10 = false;
      this._f11 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f12 = [];
      this._f13 = _dafny.ZERO;
      this._f14 = false;
      this._f16 = false;
      this._f17 = _dafny.ZERO;
      this._f18 = _dafny.ZERO;
      this._f21 = false;
      this._f22 = [];
      this._f23 = [];
      this._f24 = _dafny.ZERO;
      this._f25 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this).f19 = f19;
      (_this).f20 = f20;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
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
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
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
    get f11() {
      let _this = this;
      return _this._f11;
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
    get f21() {
      let _this = this;
      return _this._f21;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f26 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f26) {
      let _this = this;
      (_this).f26 = f26;
      return;
    }
    fm1(p0, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("yim");
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.of();
      let _178_v0;
      _178_v0 = _dafny.Seq.of(p1, p1);
      let _179_v1;
      _179_v1 = p0;
      let _180_v2;
      _180_v2 = _dafny.Seq.of(_178_v0, _178_v0, _module.__default.fm2((_179_v1), globalState));
      let _181_v3;
      _181_v3 = _dafny.Map.Empty.slice().updateUnsafe((p1).isLessThan(p1),(_180_v2)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(532)), function (_182_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length)), new BigNumber((_180_v2).length))]);
      let _183_v4;
      let _nw35 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _183_v4 = _nw35;
      let _184_v5;
      _184_v5 = _dafny.Seq.UnicodeFromString("ylwodeou");
      let _index29 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_183_v4).length));
      (_183_v4)[_index29] = _184_v5;
      let _185_v6;
      _185_v6 = _dafny.Seq.of(_this.f26, !(_this.f26), _this.f26);
      let _index30 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_183_v4).length));
      let _rhs29 = new BigNumber((_dafny.Seq.Concat(_178_v0, _dafny.Seq.Concat(_178_v0, _178_v0))).length);
      let _rhs30 = (_185_v6)[_module.__default.safeIndex(p1, new BigNumber((_185_v6).length))];
      let _rhs31 = (_181_v3).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f26,_178_v0));
      let _rhs32 = _184_v5;
      let _lhs25 = globalState;
      let _lhs26 = _183_v4;
      let _lhs27 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_183_v4).length));
      _lhs25.f15 = _rhs29;
      r0 = _rhs30;
      _181_v3 = _rhs31;
      _lhs26[_lhs27] = _rhs32;
      r0 = (_179_v1);
      let _hi1 = p1;
      for (let _186_i1 = (_dafny.ZERO).minus(p1); _186_i1.isLessThan(_hi1); _186_i1 = _186_i1.plus(_dafny.ONE)) {
        let _187_v7;
        let _nw36 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _187_v7 = _nw36;
        let _index31 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_187_v7).length));
        (_187_v7)[_index31] = _186_i1;
        let _188_v8;
        _188_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _189_v9;
        _189_v9 = _dafny.Map.Empty.slice().updateUnsafe(_186_i1,_this.f26);
        let _190_v10;
        _190_v10 = _dafny.MultiSet.fromElements(_186_i1);
        let _index32 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_187_v7).length));
        let _rhs33 = _187_v7;
        let _rhs34 = !(false);
        let _rhs35 = _module.__default.safeDivisionInt(p1, _module.__default.safeDivisionInt(p1, new BigNumber(((_188_v8).update(true, p1)).length)));
        let _rhs36 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(848)), function (_191_i2) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        });
        let _rhs37 = (((_188_v8).contains((p1).isLessThan(_module.__default.fm3(_186_i1, _189_v9, _186_i1, globalState)))) ? ((_188_v8).get((p1).isLessThan(_module.__default.fm3(_186_i1, _189_v9, _186_i1, globalState)))) : (new BigNumber((_190_v10).cardinality())));
        let _lhs28 = globalState;
        let _lhs29 = _187_v7;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_187_v7).length));
        let _lhs31 = globalState;
        _lhs28.f5 = _rhs33;
        r0 = _rhs34;
        _lhs29[_lhs30] = _rhs35;
        _184_v5 = _rhs36;
        _lhs31.f19 = _rhs37;
        (_this).f26 = _this.f26;
        let _192_v11;
        _192_v11 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (p0) : (p0)),_this.f26);
        _192_v11 = (_192_v11).update(((p0) ? (p0) : (!(_this.f26))), _this.f26);
        let _193_v12;
        _193_v12 = _dafny.Map.Empty.slice().updateUnsafe(_186_i1,_187_v7);
        let _rhs38 = _this.f26;
        let _rhs39 = (_185_v6)[_module.__default.safeIndex((p1).multipliedBy(new BigNumber((_189_v9).length)), new BigNumber((_185_v6).length))];
        let _rhs40 = (_186_i1).multipliedBy(new BigNumber((_193_v12).length));
        let _rhs41 = (_dafny.ZERO).minus(new BigNumber((_178_v0).length));
        let _rhs42 = ((_187_v7)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_187_v7).length))]).isLessThanOrEqualTo((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f26,_this.f26)).length)).multipliedBy((_187_v7)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_187_v7).length))]));
        let _lhs32 = globalState;
        let _lhs33 = globalState;
        let _lhs34 = _this;
        r0 = _rhs38;
        r0 = _rhs39;
        _lhs32.f15 = _rhs40;
        _lhs33.f15 = _rhs41;
        _lhs34.f26 = _rhs42;
      }
      let _194_v13;
      _194_v13 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,p1);
      let _195_v14;
      _195_v14 = _dafny.Set.fromElements(p0, _this.f26, p0);
      let _196_v15;
      _196_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(456)), function (_197_i3) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length),_this.f26);
      _194_v13 = ((_194_v13).Merge(_194_v13)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_178_v0, new BigNumber(693), new BigNumber((_195_v14).length), globalState),_module.__default.fm3(p1, _196_v15, p1, globalState)));
      (globalState).f19 = _module.__default.fm3(new BigNumber((_module.__default.fm4((_dafny.ZERO).minus(p1), _this.f26, _this.f26, new BigNumber((_dafny.MultiSet.fromElements(p1, p1)).cardinality()), globalState)).length), _196_v15, (_dafny.ZERO).minus((p1).plus(p1)), globalState);
      let _198_v16;
      _198_v16 = _dafny.MultiSet.fromElements(_this.f26, false, _this.f26, _this.f26, _this.f26);
      let _199_v17;
      _199_v17 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _200_v18;
      _200_v18 = _module.D1.create_DC2(new BigNumber(((_183_v4)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_183_v4).length))]).length), p1, new BigNumber(-594), _199_v17);
      let _201_v19;
      let _nw37 = Array((new BigNumber(19)).toNumber());
      _nw37[(_dafny.ZERO).toNumber()] = _this.f26;
      _nw37[(_dafny.ONE).toNumber()] = false;
      _nw37[(new BigNumber(2)).toNumber()] = _this.f26;
      _nw37[(new BigNumber(3)).toNumber()] = p0;
      _nw37[(new BigNumber(4)).toNumber()] = false;
      _nw37[(new BigNumber(5)).toNumber()] = true;
      _nw37[(new BigNumber(6)).toNumber()] = p0;
      _nw37[(new BigNumber(7)).toNumber()] = true;
      _nw37[(new BigNumber(8)).toNumber()] = _this.f26;
      _nw37[(new BigNumber(9)).toNumber()] = p0;
      _nw37[(new BigNumber(10)).toNumber()] = _this.f26;
      _nw37[(new BigNumber(11)).toNumber()] = true;
      _nw37[(new BigNumber(12)).toNumber()] = p0;
      _nw37[(new BigNumber(13)).toNumber()] = (p0);
      _nw37[(new BigNumber(14)).toNumber()] = p0;
      _nw37[(new BigNumber(15)).toNumber()] = _this.f26;
      _nw37[(new BigNumber(16)).toNumber()] = _this.f26;
      _nw37[(new BigNumber(17)).toNumber()] = p0;
      _nw37[(new BigNumber(18)).toNumber()] = true;
      _201_v19 = _nw37;
      let _202_v20;
      _202_v20 = _dafny.Map.Empty.slice().updateUnsafe(_200_v18,_201_v19);
      let _203_v21;
      _203_v21 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(p1, _196_v15, new BigNumber((_198_v16).cardinality()), globalState),(((_202_v20).contains(_200_v18)) ? ((_202_v20).get(_200_v18)) : (_201_v19)));
      _203_v21 = _203_v21;
      let _204_v22;
      _204_v22 = _dafny.MultiSet.fromElements(p1);
      let _205_v23;
      _205_v23 = new _dafny.CodePoint('c'.codePointAt(0));
      let _206_v24;
      _206_v24 = _dafny.Map.Empty.slice().updateUnsafe(_204_v22,_205_v23);
      r0 = !((_206_v24).Merge(_206_v24)).equals(_206_v24);
      let _207_v25;
      _207_v25 = _dafny.Seq.of(_204_v22, (_204_v22).Difference(_module.__default.fm5(_module.__default.fm0(_178_v0, p1, p1, globalState), _this.f26, _this.f26, globalState)));
      r1 = _207_v25;
      return [r0, r1];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
